PROTOC := protoc
GO := go
HAVE_PROTOC := $(shell which $(PROTOC))
CGO_ENABLED := 1
HOST_OS := linux

export SHIFTLEFT_DEBUG ?= 0
GO_BUILD_ARGS := $(shell test "$(SHIFTLEFT_DEBUG)" -eq 1 && echo -race)

GO_PKGS := $(shell go list ./... | grep -v /vendor)
GO_TEST := GOOS=$(HOST_OS) GOARCH=$(HOST_ARCH) $(GO) test

# If you can use docker without being root, you can do "make SUDO="
SUDO=$(shell docker info >/dev/null 2>&1 || echo "sudo -E")
INSIDE_CONTAINER ?= false
DOCKER_BUILDER_IMAGE := shiftleftsecurity/builder
DOCKER_BUILDER_FILE := builder.Dockerfile
VENDOR_DIR := vendor
BUILD_DIR := build
BIN_DIR := $(BUILD_DIR)/bin
SLAGENT_BIN := slagent
SLAGENT_MAIN := cli/main.go
GENERATOR_DIR := generator
# The only proto definitions are in generator for now
PROTO_DIR := $(GENERATOR_DIR)
METAGENERATOR_DIR := metagenerator
BPF_DIR := bpf

TRACER_DIR := $(realpath tracer)
BATTERY_DIR := $(realpath battery)
CONFIG_FILE := $(realpath config.json)

PROTO_SOURCES := $(wildcard $(PROTO_DIR)/*.proto)
PROTO_TARGETS := $(patsubst $(PROTO_DIR)/%.proto,$(PROTO_DIR)/%.pb.go,$(PROTO_SOURCES))

.PHONY: all
all: protogen handlers pregen slagent

#
# Docker builder image

.PHONY: build-docker-image delete-docker-image
build-docker-image:
	$(SUDO) docker build -t $(DOCKER_BUILDER_IMAGE) -f $(DOCKER_BUILDER_FILE) .

delete-docker-image:
	$(SUDO) docker rmi -f $(DOCKER_BUILDER_IMAGE)

#
# Main slagent target

ifeq ($(INSIDE_CONTAINER),false)

$(SLAGENT_BIN): build-docker-image
	$(SUDO) docker run --rm \
		-v $(PWD):/go/src/github.com/ShiftLeftSecurity/traceleft \
		-e INSIDE_CONTAINER=true \
		--workdir=/go/src/github.com/ShiftLeftSecurity/traceleft \
		$(DOCKER_BUILDER_IMAGE) \
		make $@
	sudo chown -R $(UID):$(UID) $(BIN_DIR)/$@

else

$(SLAGENT_BIN):
	CGO_ENABLED=$(CGO_ENABLED) \
	$(GO) build $(GO_BUILD_ARGS) \
	-o $(BIN_DIR)/$@ $(SLAGENT_MAIN)

endif

#
# Proto targets

PROTOC_CMD = \
	@if [ -n "$(HAVE_PROTOC)" ]; then \
		$(PROTOC) -I $(dir $<) --go_out=$(dir $<) $< ;\
	else \
		echo "Warning! protoc not found." \
			"Should be ok if not modifying protos." ;\
	fi

protogen: $(PROTO_TARGETS)

$(PROTO_DIR)/%.pb.go: $(PROTO_DIR)/%.proto
	$(PROTOC_CMD)

.PHONY: pretest
pretest: vet lint fmt

.PHONY: test
test: pretest
	$(GO_TEST) -v $(GO_PKGS) > $(BUILD_DIR)/report.txt
	cat $(BUILD_DIR)/report.txt

.PHONY: pregen
pregen: bpf battery

#
# Metagenerator target

.PHONY: metagen
metagen:
	$(MAKE) -C $(METAGENERATOR_DIR) \
		TRACER_DIR=$(TRACER_DIR) \
		BATTERY_DIR=$(BATTERY_DIR) \
		CONFIG_FILE=$(CONFIG_FILE)

#
# BPF common trace target

.PHONY: bpf
bpf: build-docker-image
	$(MAKE) -C $(BPF_DIR)

#
# BPF battery targets

.PHONY: battery
battery: build-docker-image
	$(MAKE) -C $(BATTERY_DIR)

#
# Go pretest targets

FIND_GO_FILES := find . -type f -name '*.go' \
	! -path './vendor/*' \
	! -path './.git/*' \
	! -name '*.pb.go' \
	! -name '*-generated.go'

.PHONY: vet
vet:
	$(FIND_GO_FILES) -exec go tool vet {} \;

.PHONY: lint
lint:
	$(FIND_GO_FILES) -exec golint {} \;

.PHONY: fmt
fmt:
	$(FIND_GO_FILES) -exec gofmt -s -w {} \;

#
# Misc targets


traceleft-handlers:
	$(GO) build $(GO_BUILD_ARGS) -o $(BIN_DIR)/$@ github.com/ShiftLeftSecurity/traceleft/cli/traceleft-handlers

.PHONY: handlers
handlers: traceleft-handlers
	$(BIN_DIR)/traceleft-handlers gen-handler

.PHONY: clean
clean:
	rm -rf $(BIN_DIR) \
		$(BUILD_DIR) \
		$(BATTERY_DIR)/handle_syscall_*.c \
		$(BATTERY_DIR)/out \
		$(BPF_DIR)/out

