// Code generated by go-bindata.
// sources:
// ../dist/trace_events.bpf
// DO NOT EDIT!

package probe

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _trace_eventsBpf = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x9d\x0b\x6c\x9c\xc5\xb5\xc7\x67\x6d\xc7\x5e\xaf\xf3\x70\x48\x42\x9c\x97\xe2\x10\x42\x1c\xf2\xf2\xda\x0e\x71\x08\x21\xce\xdb\x70\xcd\x4d\xc8\xd3\x49\x9c\x6c\x36\xeb\x75\xfc\x58\xdb\xcb\xee\x82\xed\x1b\xae\x6e\xb8\x05\xe4\x0a\x21\xa2\x06\x2a\xab\x2a\x12\x21\x54\xb8\x55\x84\x42\x2b\x88\x51\x69\x1d\x55\x48\x75\xab\x14\x19\x89\xaa\xae\x90\x8a\xd5\x22\x61\xa4\x3e\xa8\x8a\xc0\x52\x03\x5b\xcd\xf9\x66\xbd\x9f\xe7\x9c\xb3\xdf\xa8\xaa\x54\xa4\xce\x48\x30\x3e\xbf\x39\x67\xfe\xf3\x9d\xef\x9b\x87\x77\xd7\x9b\xff\xdb\xd3\xb0\x37\xcf\xe7\x13\x99\xe2\x13\x5f\x88\xac\x95\x2d\xb5\x9d\xd9\x9f\xeb\xd4\xff\x7b\x85\x4f\x0c\xdf\xee\xb0\xa7\x84\x10\xb3\x85\x10\xe7\x03\x93\x69\x69\xf7\x08\x21\xca\x85\x10\x91\xc0\x67\x60\xf7\x85\x5b\xc1\xaf\xef\x6c\x1c\xea\xf3\x77\x8f\x03\x3f\xbf\x78\x02\xea\xe1\x97\x9d\x7e\x8a\xf2\x84\xf8\x2c\x9d\x4e\x97\x69\x83\x78\x0a\xc6\x26\xc4\x02\x11\x00\x3b\xac\xda\x17\xf8\xfc\x42\x8f\x9f\x4c\xa7\xd3\xc3\x57\x94\x9d\x2f\xc4\x38\xd1\xdf\xf5\x82\x6c\xbf\x79\xd2\x56\xfc\x05\x55\x73\xd7\x25\x6d\xa9\x38\x5c\x24\x48\x5d\x6e\xdc\xc3\xc5\x6a\xbc\xc5\x27\x9c\xeb\x6f\xaf\x74\xae\x7b\x50\xc5\xfb\xe9\xf8\xe1\xff\xcf\xf6\x93\x2f\xc7\x59\x28\x54\xfc\xa7\x4e\xfe\xda\x55\xfe\x2e\xab\x7e\x7c\x42\x4c\xa4\xd3\xe9\xeb\x79\x42\x94\x4a\x7e\x29\x1b\x5f\x00\xfe\x2a\xef\xed\x63\xaa\x1e\x85\x3a\xd2\x7e\xd3\xb9\x4f\x57\x9c\x7e\x65\xde\xca\xb5\x7e\xc7\x55\xbf\x7e\x57\x7f\x7d\x57\xc6\xa7\xfc\xeb\x34\xff\x31\x0f\xff\x0a\xcd\x7f\x94\xf4\x1f\x9d\xf2\x2f\xd5\xfc\x6f\x2a\xff\x02\x97\x7f\xae\xfc\xc1\x7d\xf0\x67\xed\x19\x70\x5f\x6b\xa1\xff\xeb\x99\xe7\xc9\xef\xfc\xd0\xf7\x84\xca\xf3\xe2\x3a\x68\x0f\x5f\xbe\xa5\x9e\xd7\x7a\x95\xb7\x03\x53\x71\xb1\x98\xe4\x67\xc0\x96\xfa\x2d\xcd\xa1\x36\xf9\x73\x57\x32\x15\x3e\xbf\xb8\x51\xf9\x5f\x54\x75\xbf\xaa\x2f\xa8\xba\x57\xd5\x71\x55\xb7\xaa\xfa\x25\x55\x0f\x38\xf7\xe5\xf2\x84\xd2\x1f\x54\xb6\xba\x7f\x8b\xaf\xa9\xf1\x39\xf7\x6f\xf8\x36\x67\xdc\x3d\x79\x42\xcc\x87\x79\x36\x04\xfc\x98\x4f\x88\x74\x7a\x8e\x38\xbf\xf8\x06\xd8\xe7\x8a\x9c\xfb\xdb\xa3\xea\xe1\x57\x54\x5e\x0b\x84\xa8\x4d\xa7\xd3\x7d\x97\x9d\xe7\xb3\x2c\x4f\xcb\x67\xd2\xa9\xaf\xcf\x10\xe2\xa2\xca\xe3\x22\x83\xf9\x93\x5d\x0f\x6e\xa1\xf9\x7e\x2b\xc7\xbc\x39\x92\x99\xef\x57\x54\x5c\xd8\x69\xd7\xc7\x25\xfd\x67\xe6\x68\xbf\x9e\x3f\xdd\xcf\x8e\xd7\x8e\xd7\x8e\xd7\x8e\xd7\x8e\xd7\x8e\xd7\x8e\xd7\x8e\xd7\x8e\xd7\x8e\xd7\x8e\xd7\x8e\xd7\x8e\xd7\x8e\xd7\x8e\xd7\x8e\xd7\x8e\xd7\x8e\xd7\x8e\xd7\x8e\xd7\x8e\xd7\x8e\xf7\xeb\x38\xde\x02\xd7\x7f\x42\xbd\x97\x25\xbb\x48\x25\xc2\x91\x68\x2c\xda\x92\x12\xff\x99\xc5\xa7\xd2\x01\x6f\xf3\x5c\x70\x98\xcd\x4b\x36\x2f\xbe\xa9\xf7\x6f\x6d\x11\xea\x3d\xd5\xa9\x79\x74\xe1\xdf\x3d\x9a\xaf\x4f\xb1\x79\xa1\x8b\xcd\x0b\x5d\x6c\x5e\xe8\x62\xf3\x42\x17\x9b\x17\xba\xd8\xbc\xd0\xc5\xe6\x85\x2e\x36\x2f\x74\xb1\x79\xa1\x8b\xcd\x0b\x5d\x6c\x5e\xe8\x62\xf3\x42\x17\x9b\x17\xba\xd8\xbc\xd0\xc5\xe6\x85\x2e\x36\x2f\x74\xb1\x79\xa1\x8b\xcd\x0b\x5d\x6c\x5e\xe8\x62\xf3\x42\x17\x9b\x17\xba\xd8\xbc\xd0\xc5\xe6\x85\x2e\x3e\x9b\x17\xb2\xf8\xd4\x7b\xb0\xa5\x22\xfb\xfe\xb4\x2d\xce\xdf\x06\x76\x25\x53\xe1\x58\xec\x9f\x8b\xb7\xf3\x90\x2e\x36\x2f\x74\xb1\x79\xa1\x8b\xcd\x0b\x5d\x6c\x5e\xe8\x62\xf3\x42\x17\x9b\x17\xba\xd8\xbc\xd0\xc5\xe6\x85\x2e\x36\x2f\x74\xd9\x77\xa0\x41\x7c\x95\x4e\xa7\x4b\x95\xed\xfb\x9f\x83\xc2\xff\x78\x89\x6f\x89\xfa\xdd\xa2\x4c\xf1\x31\x55\x4b\xb6\x4e\xfd\x5c\x97\x9f\x65\x99\xcf\x3c\x5e\x73\xf9\x1d\x26\x58\x2b\xc1\x1e\x27\xd8\x45\x82\xbd\x42\xb0\x21\x82\xdd\x24\xd8\x38\xc1\x3e\x23\x98\xfa\x7a\x92\x69\x6c\x29\xc1\x2a\x09\xb6\x9b\x60\x8d\x04\x8b\x11\xec\x02\xc1\x9e\x27\xd8\x20\xc1\xde\x26\xd8\x28\xc1\x3e\x22\xd8\x24\xc1\x66\xe6\x61\x56\x4e\xb0\x1a\x82\xd5\x13\xac\x89\x60\x71\x82\x3d\x49\xb0\x01\x82\x5d\x25\xd8\x0d\x8d\xcd\x17\x42\xbc\xaf\xb1\x5c\xa5\xb3\x44\xc0\x2a\x31\xa2\xf1\xd4\x4c\xf9\xff\x85\xe2\x25\x8d\x37\x01\x5f\x84\xf8\x6e\xe0\x4b\x10\x5f\x05\x7c\x19\xe2\x85\xc0\x97\x23\xfe\x31\x8c\x67\x05\xe2\x23\xc0\x57\x22\xfe\x3a\xf0\x55\x88\x7f\x07\xf8\x6a\xc4\xbf\x01\x7c\x0d\xe2\x2d\xc0\x67\x08\xbf\xf6\x79\xf9\x90\xe2\x95\xf9\xd3\x79\x0c\xc6\xbf\x16\xf5\x73\x18\xf8\x7a\xc4\xef\x03\xbe\x11\xf1\x65\xc0\x83\x88\x7f\x05\xba\xd5\x88\x7f\x08\x7c\x13\xe2\x37\x80\x6f\x46\xfc\xfb\xc0\xb7\x20\xfe\x2d\xe0\x5b\x11\x3f\x0f\x7c\x1b\xe2\xeb\x61\x9c\x45\x88\x37\x03\xdf\x8e\x78\x03\xf0\x1d\x88\xd7\x00\xdf\x85\xf8\x02\xe0\x7b\x10\xff\x1c\xc6\xb3\x0f\xf1\xdf\x00\x7f\x00\xf1\xd9\xd0\x4f\x31\xe2\x7f\x01\xff\x12\xc4\x47\x81\xcf\x42\xfc\x2d\xe0\x73\x10\x7f\x19\xf8\x5c\xc4\x9f\x01\x3e\x0f\xf1\x14\xf0\x05\x88\x0b\x98\x75\xf9\x04\x95\x7c\x06\xc3\x8b\x18\x5e\xcc\xf0\x12\x86\xcf\x62\xf8\x1c\x86\xcf\x65\xf8\x3c\x86\x2f\x60\xf8\x42\x86\x2f\x62\xf8\x12\x86\x2f\x63\xf8\x72\x86\xaf\x60\xf8\x4a\x86\xaf\x62\xf8\x6a\x86\xaf\x61\xf8\x5a\x86\xaf\x67\xf8\x46\x86\x07\x19\x5e\xcd\xf0\x4d\x0c\xdf\xcc\xf0\x2d\x0c\xdf\xca\xf0\x6d\x0c\xdf\xce\xf0\x1d\x0c\xdf\xc5\xf0\x3d\x0c\xdf\xc7\xf0\x07\x10\x2b\x0d\xc8\xfd\x31\x89\xf8\x1b\x45\x92\xe3\xbf\xd6\xb9\x5a\x20\xf9\x7f\x21\x3e\x04\xfc\x1c\xe2\x8f\xf8\x25\x6f\x41\x3c\x5a\x28\xf9\x59\xc4\x4f\xc2\xf7\xe6\x45\x10\x7f\x7d\x86\xe4\xcd\x88\x3f\xe3\x93\x3c\x8a\x78\x0d\xf4\x7f\x00\xf1\x3b\xa0\xff\x87\x11\x6f\x00\xff\xd3\x88\xdf\x0f\xfe\x21\x3c\x7e\x18\xcf\x19\xc4\xab\x61\x3c\x61\xc4\x07\xc0\xff\x04\xe2\x5d\xe0\x7f\x12\xf1\x07\xc1\xbf\x09\xf1\xf9\xe0\x7f\x0a\xf1\x9b\x90\xff\x4e\xc4\x9f\x86\xf3\x4f\x17\xe2\x17\xa0\xff\xa3\x88\x1f\x84\xfe\x8f\x21\x5e\x09\xfe\x8d\x88\x7f\x0e\xfd\x1f\x47\xfc\x17\xe0\xff\xdf\x88\xbf\x0a\xfd\xef\x47\xfc\x7f\x21\xff\x07\x11\xef\x81\xfc\x1f\x42\xbc\x0c\xfc\x1f\x41\x3c\x0f\xfc\x13\x88\x4f\xc2\x78\xba\x11\xff\x1d\x8c\x27\x8e\xf8\x1c\xf0\x6f\x43\xfc\x3d\xb8\xde\x76\xc4\xff\x0c\xf9\xef\x40\xfc\x2a\xf8\xe3\x37\x1b\xc6\xa1\xff\xc3\x88\xff\x0c\xc6\x73\x04\xf1\xbb\x61\x1e\xe1\x7d\xf0\x45\x98\xbf\x78\xfd\xdc\x09\xf3\x17\xaf\xf3\xef\x16\x4b\x8e\xf7\xa3\x3a\xe8\x07\xaf\xab\x4d\x70\x5d\x78\xfd\xfc\x12\xf2\x8f\xf7\x85\xcf\xe1\x7b\x22\xf1\xfa\xbf\x06\xf2\x80\xd7\xbd\x1f\x43\x3f\x78\xbf\xbb\x04\xfd\xe0\x7d\xed\xd7\x70\x5d\x78\xff\x8d\xc3\x75\xe1\x75\x2f\x0f\x38\x5e\x3f\xef\x84\xfe\xf1\x3a\xff\x1a\x3c\x3f\x78\x3d\x2f\x29\x91\x1c\x9f\x2b\x9e\x81\xfb\x82\xcf\x0f\xf3\x81\xe3\xf3\xc9\x53\x90\x67\xbc\xaf\xbd\x08\xd7\x85\xf7\xdf\x37\x61\xfc\xf8\x9c\xb0\x1a\xfa\xc1\xfb\xdd\x0e\xb8\x5f\x78\x5f\xfb\x00\xae\x17\xef\xbf\x4f\x82\x2e\xde\xc7\x8b\xe0\x7e\xe1\xfd\xe8\x0a\xdc\x2f\x7c\x0e\xe9\x83\xfe\xf1\x79\xe3\xa7\xd0\x3f\x3e\x17\x1d\x86\xeb\xc2\xfb\xd1\xef\x21\x6f\x78\x5f\xcb\x87\xfe\xf1\xfe\xfb\x2c\xdc\x2f\xbc\xcf\xbe\x08\xfd\xe0\xf3\xdb\x17\x90\x37\x7c\x0e\xfc\x08\x78\x03\xe2\xcf\xc3\xf5\x3e\x84\x78\x66\x25\x94\xbf\x1e\xcb\x55\x6e\xd0\x65\x4b\xef\x5a\xcd\xee\xd5\xec\x21\x5f\xd6\x7e\x41\x08\xe1\xcf\xcf\xda\x72\x14\x95\x2e\xff\x6f\x0b\x31\xb5\xd3\x48\x3b\xea\x7a\x89\x2a\x63\xe7\xf2\x6f\xd1\xfc\x5b\x3c\xfc\xbb\x35\xff\x6e\x0f\xff\xb8\xe6\x1f\xf7\xf0\x8f\x68\xfe\x11\x0f\xff\x66\xcd\xbf\xd9\xc3\xff\x90\xe6\x7f\xc8\xc3\xff\xb0\xe6\x7f\xd8\xc3\x3f\xa4\xf9\x87\x3c\xfc\xcf\x68\xfe\x67\x3c\xfc\xc3\x9a\x7f\xd8\xc3\xff\xac\xe6\x7f\xd6\xc3\xff\x61\xcd\xff\x61\x0f\xff\x83\x9a\xff\x41\x0f\xff\xe3\x9a\xff\x71\x0f\xff\x13\x9a\xff\x09\x0f\xff\x93\x9a\xff\x49\x0f\xff\x26\xcd\xbf\xc9\xc3\x7f\xbf\xe6\xbf\xdf\xc3\xff\x80\xe6\x7f\xc0\xc3\xff\x88\xe6\x7f\xc4\xc3\xff\xa8\xe6\x7f\xd4\xc3\xff\x98\xe6\x7f\xcc\xc3\xbf\x51\xf3\x6f\xf4\xf0\xef\xd0\xfc\x3b\x3c\xfc\x63\x9a\x7f\xcc\xc3\xbf\x53\xf3\xef\xf4\xf0\xef\xd2\xfc\xbb\x3c\xfc\x4f\x69\xfe\xa7\x3c\xfc\x4f\x6b\xfe\xa7\x3d\xfc\xdb\x34\xff\x36\x0f\xff\x76\xcd\xbf\xdd\xc3\xff\x9c\xe6\x7f\xce\xc3\xbf\x55\xf3\x97\xf6\x52\x65\x07\x84\x10\x77\xb9\xe2\x03\xf0\x6a\xa0\x10\xbb\x5d\x76\x85\x6b\xbf\x0b\xa8\xb3\x5c\xcc\x65\xdf\xed\xea\x3f\x00\xaf\xf6\x09\xf1\xbc\xcb\x5e\xe7\xda\x1f\x03\xf0\xaa\x9f\x10\x6f\xbb\xec\x0d\x42\x88\x51\x97\x2d\x4f\xa1\x1f\xb9\x6c\x39\xb6\x49\x97\x2d\x4f\x3d\x33\x7d\x59\xbb\x4a\x08\x51\xee\xb2\xe5\x69\xb5\xc6\x65\xd7\x08\x21\xea\x5d\xb6\x3c\x1d\x35\xb9\xec\x7b\xe4\x9e\xe5\xb2\xe5\xa9\xf0\x49\x97\x2d\xf7\xf2\x01\x97\x2d\x4f\x21\x57\x5d\xf6\xbd\x42\x88\x1b\x2e\x5b\x9e\x1e\xdf\x77\xd9\xf7\x09\x21\x26\x5c\xb6\x3c\xad\xdc\x72\xd9\xf7\xcb\xb3\x56\x5e\xd6\x96\xa7\xe4\x3b\x5d\x76\x9d\x1c\x83\xcb\x96\xa7\xb2\x06\x97\xbd\x53\xde\x6f\x97\x2d\x4f\xbb\x29\x97\x2d\xef\x65\xbf\xcb\x96\xa7\xab\xef\xba\xec\xbd\x42\x88\x6b\x2e\x5b\x9e\xa2\xdf\x71\xd9\xf5\x42\x88\x31\x97\x0d\xa7\xb6\x0d\xa9\x68\x6f\x4a\x74\x24\xa2\xa9\x78\xa2\xfb\x6c\x34\x14\x6a\x0d\x77\x35\xc7\xa2\xa1\xb6\xae\x68\x2a\x14\x49\x76\x84\xc2\x91\x48\x34\x9e\x12\x1d\xb9\x9b\x37\x24\xa2\xb1\xa9\x4e\x36\x92\xad\x74\x53\x67\x38\x9e\xdc\x48\x77\x1a\x8a\x27\xba\xcf\x25\x43\x89\xe8\x74\xaf\x54\x24\x1e\x7a\xec\x9e\x50\xa4\xbb\xab\x2b\x1a\xc9\xe9\x54\xe3\xe1\xd4\xd9\xd1\xdc\x96\x08\x73\xad\x2d\x91\xd6\xee\x9e\xae\x5c\xcd\x9d\xdd\xcd\x6c\x33\xf4\x9d\xb3\x67\xa6\x31\x57\x5b\x77\x3c\xca\x35\xf5\x24\xda\x52\xd1\x1c\xc9\x48\x46\x53\xa1\x64\x2a\x9c\xd3\x27\x12\xeb\x4e\x72\xed\xb9\xda\x9c\x54\xb0\xd7\xc3\xb7\x25\xa2\x61\x77\x13\x7e\x0a\xa7\xdf\x6b\xfd\x21\xd4\x5a\xa7\x3f\x83\x54\x23\xd9\xc2\xa8\xd6\xe4\x54\xad\xc9\xa5\x5a\xc3\xaa\xd6\xe4\x50\x55\x4f\xa3\x2e\x97\xc1\xd3\x75\x0e\xf5\x1d\x9a\xde\x82\x31\x56\xc8\x3c\xd1\xba\xc4\x14\xc7\x1a\xd3\x9b\x08\x4e\xaa\xc0\xc4\x20\x54\x1c\x4e\xaa\xb8\x9a\x08\x0e\x0f\x4d\xf4\xb1\x68\x57\x2a\xe9\xfc\xdc\xd2\xac\x3e\xfe\x19\x0a\x27\xce\x25\xbd\x57\x11\xcf\x15\xc4\x73\xf5\xe0\x57\x8e\x1c\xab\x46\x8e\x15\x83\x5b\x2d\xd8\x95\x82\x5b\x25\x98\x15\x82\x5b\x1d\xbc\x56\x86\x5c\xab\x02\xb7\x22\xb0\xab\x01\xb7\x12\x30\xab\x80\x83\x1f\xed\x4a\x25\xc2\x91\x8e\x68\x73\x28\xde\xd6\x9c\xe4\x66\x0a\x39\x4f\xb8\x59\x42\xce\x11\x7a\x1a\xb0\x93\x80\x9e\x02\x7a\x27\x5c\x1f\x64\x17\xaa\x07\x6a\x0a\xf5\x74\x53\x0d\x0e\x0f\x3d\x16\x4d\x24\xdb\x48\x07\x79\xff\xf5\x21\x01\xc3\x23\xca\x62\x8d\xb9\x7a\x75\x7d\xcc\xba\x03\xa3\xe9\x7d\xea\x0d\x88\x3a\x13\xb7\x2d\x16\x0d\x39\x33\x19\x6e\x6f\x28\xd5\x1d\xea\x09\xa7\x22\xad\x53\x02\xc9\xbe\xa4\xf3\xbc\x12\x97\xe7\x70\x7c\x2d\x2e\xae\x43\x7a\x79\x9f\x7a\xe8\xa9\xd5\x3d\xdb\x88\x17\x77\xad\x8d\x6a\xa0\x15\x61\xba\x50\x6a\x4e\x03\x56\x72\x71\x1d\x12\x0f\x06\xd5\x3b\xd5\x33\x3c\x76\x7a\xcf\x59\x18\x8a\xb5\x45\xa2\x5d\xaa\x75\x43\xb4\x35\xd4\x92\x08\x77\x52\x82\xce\x2c\xa7\xd7\x78\x76\x85\xa7\xd7\x77\xf2\x39\x27\x3a\xe7\xfa\x26\xbb\x76\x20\x3c\x72\x72\x69\x49\x84\x3b\x43\x6d\xcd\xa1\x78\x34\x21\x9f\x3a\xb7\xa2\x7c\xdc\xe4\x0a\x34\xed\xf9\x03\x80\x95\xb2\x58\x67\xc9\x54\x22\x15\x3e\x2b\x36\x24\xfb\x3a\xa1\x4e\xc8\x85\x3e\x2c\x1a\x76\xee\x0c\x86\x6a\x9c\xaa\x5a\x56\x95\x4e\xb5\x25\x54\x25\xab\x2a\x55\x07\x55\x5d\xab\xb0\xaa\x83\xaa\xde\xac\xb0\xaa\x83\xaa\xbe\x47\x61\x55\x07\x55\xbd\xc9\xa9\xaa\x55\x5d\xa5\xea\xa0\xaa\x6b\x54\xb3\xaa\xab\x54\x1d\xcc\x70\x55\xa9\xba\x4a\xd5\xc1\x8c\xad\x9a\x55\x5d\xa5\xea\x60\x86\x07\x15\x57\x75\x50\xd5\xd5\x95\x8a\xab\x3a\x28\xeb\x7f\x45\xd9\x57\x22\xc8\x4f\x15\x0c\xa9\x5f\x9c\x1f\x9f\x39\x9d\xfb\x04\xb6\xe5\x7f\x85\x1a\xe7\xbe\xdb\x4f\xff\x13\x9c\x13\xfe\xdc\xf1\x63\x1a\xf7\x6b\xf6\x51\xbf\x20\x3f\xcd\x30\xa4\xde\xba\x2b\x57\x76\x8f\xfa\x6c\x63\x26\x3e\xf3\x39\xbe\x1d\x8c\x7e\xe6\x7b\x35\xeb\xf2\x73\xeb\x6f\x63\xf4\xc7\x8e\x4d\xbf\x8e\x1e\xf5\xef\x7c\xe9\xfa\x9b\x4b\x68\xfd\x46\x95\x28\xfd\x33\x59\xba\x7e\x75\x09\xad\x5f\xa6\x5e\xa8\xa8\x74\xe9\x17\x11\xfa\xe5\x8c\xbe\x7f\x86\x99\xfe\x52\x46\xbf\x9e\xd0\x2f\x26\xf4\xdf\x63\xf2\x3f\x64\xa8\xff\x2b\x26\xff\xbd\x84\x7e\x09\xa1\xff\x16\xa3\xdf\x5a\x68\xa6\xff\x06\xa3\x3f\x48\xe8\xcf\x22\xf4\x0b\xb8\xe7\xaf\xc8\x4c\x5f\x70\xcf\x1f\xa1\x3f\x87\xd0\xff\x53\x11\xad\x3f\x62\xa8\xff\x49\x11\xad\xef\x3f\x8e\xf5\xe7\x12\xfa\xe9\x62\x5a\xbf\xd7\x6f\xa6\xff\xf7\x62\x5a\xbf\x96\xd0\x9f\x47\xe8\x7f\xcc\xe8\x57\x14\x9b\xe9\xff\x81\xd1\x6f\x25\xf4\x17\x10\xfa\x7f\x2c\xa4\xf5\xc7\x0c\xf5\x27\x0a\x69\xfd\x01\x42\x7f\x21\xa1\xff\x3e\xa3\xdf\x1f\x30\xd3\x1f\x65\xf4\x47\x08\xfd\x45\x84\xfe\x48\x3e\xad\x5f\x5b\x62\xa6\xff\x4e\x3e\xad\x3f\x49\xe8\x2f\x21\xf4\x5f\x63\xf4\x27\x0c\xf5\x7f\xc0\xe8\x57\x9c\xc0\xfa\xcb\x08\xfd\x24\x33\xff\x06\x66\x9a\xe9\x77\x33\xf3\xaf\x91\xd0\x5f\x4e\xe8\x9f\x64\xf4\xeb\x67\x99\xe9\x1f\x63\xf4\xfb\x09\xfd\x15\x84\x7e\x25\xa3\x3f\x69\xa8\xbf\x8e\xd1\x1f\x22\xf4\x57\x12\xfa\x0b\x19\xfd\xc1\xd9\x66\xfa\xf3\x18\xfd\x09\x42\x7f\x15\xa1\xbf\xbe\x80\xd9\xff\xe7\x98\xe9\xaf\x29\x60\xf6\xff\x93\x58\x7f\x35\xa1\x7f\x1b\xa3\xef\x2f\x35\xd3\x9f\xcd\xe8\xd7\x13\xfa\x6b\x08\xfd\x77\x03\xcc\xfe\x6f\xa8\xff\xcb\x00\xb3\xff\x13\xfa\x6b\x09\xfd\x37\x19\xfd\xd6\xb9\x66\xfa\x3f\x64\xf4\x07\x09\xfd\xf5\xd4\xfc\x67\xf4\xcb\x6e\x33\xd3\xef\x66\xf4\xc7\x08\xfd\x8d\x84\xfe\x09\x46\x7f\xc4\x50\xff\x28\xa3\xef\x6f\xc2\xfa\x41\x42\xff\x12\xf3\xfc\xf5\xce\x33\xd3\x7f\x8e\x79\xfe\x6a\x09\xfd\x6a\x42\xbf\x97\xd1\xaf\x98\x6f\xa6\x9f\x62\xf4\x5b\x09\xfd\x4d\x84\x7e\x84\xd9\x7f\xc6\x0c\xf5\x43\xcc\xfe\x33\x40\xe8\x6f\x26\xf4\xf7\x32\xfa\xfd\x0b\xcc\xf4\x77\x32\xfa\x23\x84\xfe\x16\xea\xfc\x99\xc7\xec\xff\xb7\x9b\xe9\x7f\x92\xc7\xec\xff\x84\xfe\x56\x42\x7f\x94\xd1\x9f\x30\xd4\xbf\xc9\xe8\x57\x9c\xc2\xfa\xdb\x08\xfd\x76\xe6\xf7\xe7\x81\x85\x66\xfa\x2d\x82\xd9\xff\x09\xfd\xed\x84\x7e\x03\xa3\x5f\x5f\x66\xa6\x5f\xcf\xe8\xf7\x13\xfa\x3b\x08\xfd\xed\xcc\xf9\x7b\xd2\x50\x7f\x2b\x73\xfe\x1e\x22\xf4\x77\x11\xfa\x2b\x19\xfd\xc1\x45\x66\xfa\xcb\x19\xfd\x09\x42\x7f\x0f\xa1\xff\x3d\x46\xbf\x71\xb1\x99\xfe\x65\x46\xbf\xec\x34\xd6\xdf\x47\xe8\x3f\xcb\xe8\xfb\x97\x98\xe9\x7f\x93\xd1\xaf\x27\xf4\x1f\x20\xf4\x5f\x55\xeb\xaf\xfe\x1a\xd4\x90\xd2\xd7\xff\x7d\x13\xfd\xf5\xa3\x0f\x03\x74\xfc\xe8\x52\xb3\xf8\xe7\x0a\xe9\xf8\x89\x65\x66\xf1\xef\xcc\xa0\xe3\x33\x1f\x39\xf6\x8a\x7f\xd9\x47\xc7\x97\x95\x9b\xc5\x6f\x64\xc6\x5f\xb9\xc2\x2c\x7e\x59\x1e\x1d\x5f\x7f\x87\x59\x7c\x2f\xa3\x7f\x66\xa5\x59\x7c\x82\xd1\xef\xbd\xd3\x2c\xfe\x03\x26\xff\x17\x57\x99\xc5\xff\x84\xc9\xff\xe0\x5d\x66\xf1\xe7\x19\xfd\x1b\xab\xcd\xe2\x1f\x62\xf4\xc7\x2a\xcc\xe2\xd7\x32\xfa\x9f\xae\x31\x8b\xff\xab\xa0\xe3\xfd\x6b\xcd\xe2\x2f\x31\xfa\xe5\xeb\xcc\xe2\xdb\x98\xeb\xaf\x5d\x6f\x16\xbf\x87\xd1\x3f\xb0\xc1\x2c\x7e\x0e\xa3\xdf\xba\xd1\x2c\x7e\x1f\xf3\xfc\x5f\xa8\x34\x8b\xdf\xc2\x3c\xff\x03\x41\xb3\xf8\x33\xcc\xf5\x5f\xab\x32\x8b\xdf\xc0\x5c\xff\x48\xb5\xa1\x3e\x73\xfd\xe3\x35\x66\xf1\x47\x99\xeb\x9f\xdc\x64\x16\x7f\x95\xb9\xfe\xd2\xcd\x66\xf1\x4f\x33\xd7\x5f\x51\x6b\x16\x1f\xf3\xd3\xf1\x75\x5b\xcc\xe2\x7f\xc4\xec\x7f\x8d\xf7\x9a\xc5\x1f\x53\xaf\xdf\x6b\x7f\x86\x2c\xe2\xea\x4f\x37\xf4\xf7\x61\xf4\xf7\x6f\x02\xdc\xf3\xb3\xd5\x4c\xff\x26\xb3\x7e\x8c\xdc\x67\x16\xff\x09\x73\xfd\xe3\xdb\xcc\xe2\x5f\x65\xf4\x27\xef\x37\x8b\xff\x39\xa3\x5f\x5a\x67\x16\xff\x04\xa3\x5f\xb1\xc3\x2c\xfe\x6f\x4c\xfe\xeb\x76\x9a\xc5\xff\x96\x79\x7e\x1b\x77\x99\xc5\xcf\x63\xe6\x6f\x7c\xb7\x59\xfc\x97\x8c\x7e\xff\x1e\xb3\xf8\xb9\xcc\xf9\xed\xa5\xbd\xb4\xbf\xfe\xfc\xbe\x59\x44\xc7\xbf\xc2\xc4\xa3\xf3\x4f\x80\x9e\x3f\x83\x2a\x7e\x42\x6b\xd0\xcf\xbf\x8b\xb8\xd7\x9f\xd4\xf9\xb7\x4e\xc5\xcb\xf3\xef\xa3\xc4\xf9\x77\x7f\x09\xd6\x96\xe5\xa2\xfa\x93\xa0\x0a\xf5\x3a\xa4\x1c\xe3\x83\xae\xf8\xcc\xf7\x8d\xfc\x23\x00\x00\xff\xff\x87\xbf\x98\xdb\x38\x8b\x00\x00")

func trace_eventsBpfBytes() ([]byte, error) {
	return bindataRead(
		_trace_eventsBpf,
		"trace_events.bpf",
	)
}

func trace_eventsBpf() (*asset, error) {
	bytes, err := trace_eventsBpfBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "trace_events.bpf", size: 35640, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"trace_events.bpf": trace_eventsBpf,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}
var _bintree = &bintree{nil, map[string]*bintree{
	"trace_events.bpf": &bintree{trace_eventsBpf, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}

