// Code generated by go-bindata.
// sources:
// data/dictionary.json
// DO NOT EDIT!

package data

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

var _dataDictionaryJson = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x5d\x6d\x6f\xdb\xd6\x15\xfe\x9e\x5f\xc1\x19\xfb\xd0\x02\x8e\xdc\xf7\xa6\x46\x31\x2c\x68\x00\x23\xc5\x26\x17\x96\xfb\x61\x1d\xda\x84\x92\x68\x92\x16\x79\xc9\x92\x94\x5f\x3a\x18\xc8\x32\x64\xee\x96\x75\x1b\x90\x66\x6e\xba\x2c\x5e\x67\xd7\xd6\xdc\xd4\x0b\x34\xc0\x91\xbf\xb4\x7f\xc5\xc8\x2f\x29\x49\xf9\x45\xb2\xf3\x5c\x3e\xb4\xae\x6d\x45\x10\x10\x20\xb6\x74\xcf\xc3\x73\xcf\xdb\x3d\xf7\x9c\x7b\xe9\xdf\x5d\xd2\xb4\x11\xd7\x88\xf4\x91\x71\x2d\xf9\x39\xfe\x6d\xc6\x76\x8c\x1b\x42\x77\x8d\xf8\xa3\x91\xaa\x5d\x89\x6c\x4f\xe8\xc1\x62\x61\x36\xf4\xc4\xc8\x68\xc7\x98\xaa\x11\x56\x02\xdb\x4f\xbe\x4f\x86\x7e\x10\x78\x73\x76\xfc\x99\x76\x75\x2e\xf0\xb4\x6b\x87\x84\x9a\x2d\xb4\xf7\x4b\x93\xc5\x82\x76\xb5\xaa\xfb\x91\x51\xd5\x66\x02\xcf\xd5\xf4\x78\xd4\x65\xc7\x2e\xc7\xb8\x07\xa8\xbe\x5e\xa9\xe9\x66\xfa\x5c\xd3\x8e\xac\x7a\xb9\x50\xf1\xdc\xb1\x50\xaf\x3a\xb6\x33\x66\x7a\x97\x53\x12\xdf\xf2\x84\x11\xd9\x95\x03\x22\xc7\xae\x18\x22\x4c\x89\x7e\x7d\x7d\x5a\xfb\xd5\xfe\xaf\xfb\xdf\x86\x5e\x3d\xa8\xa4\x5f\x5a\x51\xe4\x87\xe3\x63\x63\x81\x3e\x5f\xe8\x40\x8f\xbc\xc0\x9e\xd1\xa3\xb1\xd9\x30\x61\xfb\x83\x7d\xf0\x31\x57\x0f\x23\x23\x18\x0b\x83\xca\xd8\x73\x18\x8d\x69\x4c\x5b\xe8\xce\x8d\x8a\x57\x4d\xc1\xbb\xa9\x0f\x86\xd9\xc2\x8e\xec\x78\x54\xd5\x98\x33\x1c\xcf\x37\x82\x64\xe8\x54\xf2\x38\xad\xa8\x97\x6d\xed\xdd\xc8\x2b\xa4\x4f\xff\xa5\xe9\xea\xb6\x93\xf0\xf3\x8b\x03\xda\x8a\xe7\x2f\xc6\x4f\xb1\xa2\x84\xe6\xbd\x83\x5f\xb4\x97\xde\x7b\x59\x9b\x74\xed\x4a\xe0\x09\x47\x2f\x6b\x2f\x25\xb3\x8a\x27\x35\x3f\x3f\x5f\xf0\x0e\x3f\x4e\x80\x5e\x8e\xc5\xed\x38\xda\x54\x42\x15\x6a\x53\x46\x68\x04\x73\x46\xb5\x70\x00\xaf\xb7\x55\x71\xa3\xbc\x98\xe0\x97\x52\x11\x6b\x53\xd6\x62\x64\xc5\xaa\x79\xb7\x2d\xf2\x42\x90\xfc\xee\x3e\x87\x39\x43\xc4\x13\xb7\x85\x99\xd0\xd6\xa3\x99\xcb\x57\x46\xe2\xcf\x97\x92\x2f\x47\xaa\x7a\xa7\x2d\xcd\x79\xf3\x86\x93\x8c\xd2\x0d\xdb\xab\x1f\xcd\x4d\xc4\xb6\xa4\x8b\x74\x6e\xe5\x4a\x75\xc6\xb4\x66\x6b\x8e\x2b\xfc\x4f\x83\x30\x9a\x9b\x5f\x58\xfc\xec\x70\xa4\x1e\x1a\x37\xc2\x58\xa1\xb1\x20\xe7\x52\x49\x7b\x76\xbd\x6a\xce\x8a\x78\xe0\xd1\x28\x51\x77\xcb\x6d\xe1\xbe\xf2\xea\x6b\xaf\xbf\xf1\xe6\x5b\x6f\x5f\x79\xe7\xc8\xa6\xa2\x58\x91\x22\x8c\xbf\xfd\x6d\xfa\x89\xb6\xcf\xdb\xbe\x1d\x8b\x6a\xca\x84\xe5\xec\x13\xa4\x1f\x07\x86\xef\xe8\x6d\xb3\xd9\xdb\xf8\x7e\x6f\xf3\x8b\xbd\x8d\xe6\xc8\xfe\xd7\x4b\xa3\x18\xc6\x0f\x2d\x08\xb3\xb5\xb7\xb1\xc3\x60\x94\xab\x18\xe3\x71\xca\xca\x26\x05\x33\x9b\x81\xf2\x90\x63\x26\x03\x65\x83\x42\x29\x67\xa0\x3c\xa6\x50\xb0\x8a\x1e\xf3\x2a\x2a\x63\xe9\x7e\xcf\xd0\xcf\x29\x31\x14\x89\x40\x28\x26\x7a\x9b\x43\xa5\x38\x01\x01\xbe\x4e\x27\xf1\x88\x82\xa9\x60\x59\xb6\x61\xfe\xc9\xc1\x64\xa0\x7c\x4d\xa1\x60\x5e\x38\x2e\x30\x13\x0c\x79\xd5\x12\x10\x60\x33\x9d\x45\x83\x83\x71\x33\x60\xb6\x29\x18\x13\x4b\x63\x23\x85\xf9\x8a\x82\x91\x84\xa2\x0d\x3e\x14\x55\x25\x3e\xd7\x86\xa1\xac\xb6\x8a\xad\x3e\x0f\x0a\x16\xf0\x06\x2f\xe0\x6b\xd7\x20\xca\x37\x29\xca\x37\x14\x0a\x96\xcb\x7f\x38\x6b\xc1\xc6\xc2\x99\x49\x86\x2c\x56\x38\x2b\xc9\x40\xa1\xd6\x08\x89\x40\x29\x26\x30\x0f\x0c\x79\xa1\x50\x00\x00\xc9\x37\x0c\xc0\x4d\x44\xcf\x3d\x1e\xb1\xff\xed\x7a\xfc\x8f\x82\xc0\x08\x0c\xb9\x29\x89\x5f\x5f\xf1\xf1\x6b\x42\x0d\x8c\x29\xb1\xec\x15\x3e\xf0\x4c\xa8\x81\x31\x8b\x19\x28\x6b\x14\x33\x4a\x50\x4c\x2c\xdf\x95\x1c\x6a\x52\x82\x62\xe2\x50\xba\xc2\x87\xd2\x09\x25\x28\x26\x4e\xcb\x56\xf8\xb4\x6c\x42\x09\x8a\x89\xc3\xea\x43\x3e\xbb\x9a\xc0\x39\x5a\x1e\x14\x25\xbc\x98\x4a\x78\x91\x24\x25\x54\x3a\x32\xd1\x23\x3d\x16\x05\xb5\xce\x61\x19\x50\xe4\x16\xf6\xff\x5d\xde\xff\x25\xf1\x75\x97\xf7\x5c\x49\x96\xb9\xcb\xfb\x9c\x64\xcf\xbc\xcb\x7b\x0b\x56\xea\x2e\x43\x3e\x3b\x8b\x01\xda\xb6\xf9\x2f\x0a\x46\xb2\x25\xca\x61\xe2\x12\x66\x38\x36\xf0\xae\xfd\x21\xbf\x6b\x97\x80\x30\xe4\xef\xf7\x46\x5e\xab\x49\x4c\xfd\x7e\x3a\x8b\xa7\xbc\xc1\xd7\x4a\x4a\xd1\x6a\x12\xd3\xef\x44\xa3\x1c\x20\xe6\x4d\x25\xda\x82\xca\x89\x2e\xa8\xe4\xac\x86\xcd\x7a\x1f\x8c\x14\x97\x12\x18\x49\x91\xed\x3e\x59\x64\xab\x2d\x28\xe1\xa4\x96\x81\x72\x9f\x42\x99\xce\x40\x59\xa5\x50\xa2\x0c\x14\x6a\x1b\x50\xc3\x21\xfd\x3e\x1f\xd2\x6b\x61\x06\x4a\x8b\x42\xc1\x6a\xfe\x07\x45\x8f\x99\x60\xc8\x1d\x49\xd5\xa2\xc9\xd7\x1b\x1c\xc9\x1e\xa4\xc9\xef\x41\x1c\x89\x34\x9a\xa4\x40\x1c\x49\x02\xd6\x24\x73\x28\xc7\x97\x62\x7c\xc7\xcd\x25\x43\x22\x9c\x7e\x70\x3a\xd7\xe4\x8b\x17\x0e\x76\xbe\x26\xef\x7c\x0e\xae\x5e\x34\xf9\x72\x90\xe3\x67\xa0\x6c\x51\x28\xb8\x4c\x96\xc7\x6c\xf1\xf2\xd1\xe4\x17\x0e\x07\x87\x93\x26\x1f\x4e\x1c\x5c\x2d\x6f\xf2\xed\x03\x09\x2b\x0c\xb9\x1b\x61\xc3\xdf\x4e\xb9\xf8\x96\x82\x91\xf8\x4f\x1b\x86\x72\x21\x57\x12\x9c\xb6\x79\x2d\xbb\x3e\x96\xca\x76\xd2\xb2\xa2\x75\xe4\xe2\x7d\xc9\x36\xbf\x2f\x71\xb1\x07\x6c\xf3\x1e\xe0\x62\x0f\xc8\x23\x1b\xec\x01\xdb\xbc\x07\xb8\x32\x01\xf3\xd2\xc5\x1e\xb0\xcd\x7b\x80\x3b\xa3\xc4\xf4\x30\x08\x43\xfe\x0a\x22\xdf\xa4\x0a\xb5\xaf\x42\x72\x6a\x19\x7d\x0d\x92\x53\xf6\xf9\x3a\x24\xff\x2f\x43\xfe\x06\x24\xa7\xec\xfa\x4d\x48\x4e\x29\xee\x2d\x48\x4e\x59\xcf\xdb\x90\x9c\x72\xa7\x2b\x90\x9c\x32\x9b\x77\x20\xf9\xff\x18\xf2\xa2\x29\xdb\x7c\x3c\x68\x67\x1d\x7c\xe2\x1f\xc3\x49\xf2\xb1\xfc\x70\x13\x92\xee\xe5\x23\xbe\x93\x5a\x34\x25\x19\xde\x03\xbe\xef\x17\x4f\x2f\x0b\x87\x4a\x38\x8b\x13\x92\x7a\xc8\x23\xbe\x38\x23\xcc\xc9\x0f\x33\xf8\x59\xd9\xdb\xfc\x0b\x09\x75\x3d\x1b\xea\x73\x52\x4a\x78\x17\x79\x1a\x1b\xc8\x10\x15\xd5\x87\x16\x6a\x2c\x49\xa8\x52\x9c\x2a\x7b\xcc\x12\x34\x07\x93\xa5\x2f\xd2\xf9\xf1\xce\xe3\x01\xbf\xf3\x28\x4a\xda\x27\x0f\xf8\x04\x23\xf6\xb1\x0c\x4d\x51\x45\x3b\x21\xd9\xa6\x36\xf8\x6d\xaa\x98\xce\x82\xf9\x37\x35\x29\x09\xcc\x1a\x0f\x23\x24\x39\x7b\x83\xcf\xd9\x85\x9a\x68\x28\x4c\x4f\x85\xdd\x08\x53\xcf\x0c\x60\x1b\x3f\x70\x48\x76\x36\xd2\x8f\x1c\x12\x11\x54\x6f\x71\x48\xf5\x6c\xa4\xdf\x73\x48\xc4\x9a\x71\x9b\x43\x32\xb2\x91\x96\xc9\xd5\x27\x1b\xe9\x2e\xe5\x21\x92\xb3\x23\x6d\x0f\xa1\x4e\x90\x08\x49\x2d\xb5\x41\xd6\x52\x8b\x66\x20\x99\xd4\xc6\x13\x2e\x8a\x61\x8c\x47\x24\x86\x90\xf0\x71\x9b\xc5\x50\x12\x4c\x8b\x92\x05\x82\x93\x06\xe6\x82\x0b\x58\x2a\x96\x4b\x81\x67\xc1\xb9\x0d\x2e\x09\x34\xf8\x92\x40\x11\x37\x45\xd6\xf8\x76\x48\x11\xf3\xb2\xc6\xf3\x22\xf0\x8a\xdd\xc8\xb1\x62\x63\x94\x35\x1e\x45\xe0\x13\x4f\x0d\xfe\xd4\x95\xc0\xe5\xcf\x06\x5f\xfe\x2c\x62\x94\x35\x1e\x45\xe0\x22\x6a\x83\x2f\xa2\x16\x31\xca\x1a\x8f\x22\x70\x37\xa5\xc1\x77\x53\x04\xee\x83\x34\xf8\x3e\x88\x50\xb3\x31\xc0\xac\x50\x62\xc5\x52\x65\xc8\x27\xaf\xa3\xf3\x71\xe4\x66\x6b\xf2\x43\x0c\x40\x6d\xfc\x26\x31\x3d\xb5\xe0\xe2\x6d\x63\x3c\x81\xce\x6f\xea\x8e\xd1\x79\xd9\xa0\x1b\x2e\x1d\xe2\xea\x51\xc5\x3a\x31\xe8\xe4\xc0\x74\x70\xb4\xe8\xa7\x8f\xf1\x03\x63\xc6\xee\xdc\xbf\x1c\x8e\x08\x2b\x5e\x7b\xc8\xcf\x8e\xee\x55\x1c\x1b\xb5\xd4\xf5\xfb\xc7\xdd\x20\xdd\xea\xfc\x7b\x27\xed\xd2\xe8\xf9\xcf\xc2\xaf\x8b\x4a\x54\xd7\xd3\x5b\x45\xaa\xa6\x71\xf8\xf3\xc7\x94\xa9\x49\x2c\x6d\x80\x34\xfd\xe5\x60\x68\xfa\xcb\x1e\x34\x2d\x09\x09\x03\xa4\xe8\x7b\x83\xa1\xe8\x7b\xa7\x57\xb4\x2f\x39\xaa\xf6\x1d\xdf\x86\xf1\x71\x9e\xb3\xc5\xe7\x39\x3e\xce\x2d\xb6\xf8\xdc\xc2\xc7\x0b\xfa\x16\x9f\xd1\xfa\xb8\xe1\xb6\xc5\x37\xdc\x24\x5d\xc4\x1c\x2d\x44\x1f\x67\x4b\x5b\x7c\xb6\x24\xe9\xaf\x52\x5d\x92\x19\x25\x96\x82\x3b\x6d\x14\x13\x12\xad\x30\xe4\x41\x60\xe3\x6c\xe7\x0f\x24\x82\x04\x60\x80\x62\xe3\xdd\xc1\x88\x8d\x77\x4f\x1f\x1b\x83\xe0\x23\x68\x6d\x4f\xe2\x7f\xcf\x6e\x7d\x91\xda\x3d\xd5\x67\x0b\x82\x45\x85\x60\x12\xc6\x8e\x80\xfa\xc3\x18\xb1\x2d\x1e\xa7\xea\xd1\xe8\x8d\x05\xbd\x12\x3d\x6f\xc8\x9c\xee\xd4\xd3\x21\xc1\xc5\xb3\xb0\x78\xf1\x2c\xcc\x5f\x3c\x0b\x0b\xa7\x77\xe8\xc4\xb4\x9f\x74\x7b\x4a\x37\x00\xe5\xd9\x12\x5f\x1c\xfa\xcf\xf1\x21\x43\xff\xd1\x86\xfe\x93\x02\x1d\x2d\x66\xd2\xb5\xac\x4f\x3c\x47\x92\x06\xe5\xd0\x58\x58\x9f\xc9\x7c\x50\xfb\x65\x1e\x67\xfc\x90\x33\xf6\x4f\x86\x05\x35\x29\x59\x62\x7e\x17\x9d\x5c\xf6\x75\x54\xcd\xe9\xcc\x47\x0d\xbc\x6e\x6a\xc6\x93\xa7\x60\x2b\x6b\xf3\x21\xdf\x79\x9f\x82\x5b\xcb\x4d\xea\x34\xcc\x14\x66\x82\x4a\x87\x71\x34\xea\x93\x50\xd4\xd7\xe6\x36\x5c\xc4\x2f\x84\x85\x93\x8b\xf8\xb9\xb3\xf0\x51\x8f\xa1\xe7\xf4\x71\x27\xb4\x24\xc7\xf1\x76\xf8\xe3\x78\x25\x4b\x72\x20\xea\x29\x7f\x20\xaa\x64\x49\xca\x63\x4f\xf9\x33\xe0\x25\x35\xd3\x0a\x25\xd7\xcd\x5a\x1d\x07\x28\xa9\x92\x5b\x28\x39\xa5\xd5\xe2\x4f\x69\x85\x92\x33\x63\x2d\xfe\xcc\x58\x28\x11\x74\x8b\x17\x74\x6c\x40\x19\x82\xa6\x9a\xbf\xa1\x85\x6b\xce\x3b\x7c\xcd\x39\x94\xdc\x3e\xdf\xe1\x8b\xce\xa1\xe4\xd6\xec\x0e\x7f\x02\x22\x94\x14\xf5\x77\x78\xbb\x29\x59\xf8\xb4\xcc\x53\xfe\xb4\x4c\xec\xa2\x19\x30\x54\x73\x40\x76\xdb\x39\xc7\xf5\xdf\xd8\xd1\x33\x60\xa8\x42\x72\xc9\xc2\x85\xec\x3c\xe1\x42\xa2\xf0\x1c\xd7\x90\x43\x49\x9f\xa1\xed\x53\x39\xba\x0d\x21\xd6\x7a\x2b\x87\xd6\xd5\xb8\x27\x36\x9d\x16\x6f\x3a\x21\xf6\xf1\x56\x0e\x1f\xc7\x2e\xde\xca\xe1\xe2\xd8\xfc\x5a\xbc\xf9\x85\xd8\xfa\xf2\xc4\x50\x6c\x7c\xad\x1c\xc6\x97\x65\x7b\x9c\xd5\x49\xd6\xcd\xf3\xb3\xb7\x92\x92\xd5\xa0\xa4\x64\x31\x28\x29\x59\x0b\x4a\x8a\x96\x02\x1c\xa6\x28\xed\x62\x13\xa1\x1e\xdf\x9b\x6d\x78\x1e\x6e\xfa\x51\x27\xa6\x3d\x78\x3e\x3d\xa6\xef\x8f\x1d\xe6\xf9\x15\xbb\xb2\xb7\x15\x37\x7b\x28\x0e\xfd\xe9\xa2\x2b\x43\xb2\x1a\xd7\x0b\x2c\xc9\xc3\x9f\x99\xdd\x19\xf4\x17\xca\x59\x70\x73\xf2\x0e\xdf\xe2\x44\x1e\xd7\x27\xee\xd6\x73\xc5\x57\xcd\x06\xde\x7b\x91\x0f\x3f\x9d\x87\x0c\xcf\x56\x84\x77\x2e\x5a\x84\x6a\x0a\xf2\x77\x4e\x1f\x29\x22\x49\x51\x60\x9d\x2f\x0a\x44\x37\xf1\x0a\xfd\x57\x06\x60\x1a\x6f\x14\x56\xf9\x8d\xc2\x34\x4e\xb9\x56\xf9\x94\x4b\x52\x91\xa2\x6a\x51\x11\x4e\x1f\xd7\xf9\xf4\x31\xc2\x73\x59\xe7\xe7\x22\x51\x2f\xa7\x58\x9c\x50\xaf\xf3\x09\xb5\x44\xb9\x14\x13\x98\x07\x86\x5c\xcf\x5c\xd0\xc8\x6b\x85\x57\x55\x01\xe9\xd0\x5b\x48\x46\x7a\xa4\xc7\xf7\x2c\x7f\xe8\x93\xf5\x79\x60\xb2\xb8\x3f\x5e\xf4\x12\xa3\x66\x63\xa1\x26\xdd\xd1\x5f\x68\x5d\x6e\xae\x76\xb8\xd7\x09\xee\xcf\x49\x9d\x83\x2e\xe4\x63\x0e\x93\x2b\x99\xc1\x67\x83\xb9\xcb\xde\xf8\xd2\xf8\x8f\x7d\x12\x16\x07\xa5\x4a\xb0\x7c\xd1\x6e\x34\x30\xeb\xcb\xf2\xe9\xdd\x45\x72\x73\x91\x7a\xa3\x01\xbe\x37\x78\x6b\xe8\x2e\x07\x43\x94\x28\xf9\xf3\xa1\xbb\x9c\x89\x24\x73\xb9\x4b\xbd\xc7\x22\x34\x7c\xfd\xc7\xb0\x06\x7d\x34\x64\x58\x83\xee\x63\x49\xe6\x72\x17\xc9\xb5\x76\xea\x25\x1b\xf8\xae\xf2\xed\xa1\xbb\x1c\x0c\x51\xa2\xe4\x3f\x0f\xdd\xe5\x4c\x24\x99\xcb\x5d\x0c\xa3\xc7\x6c\xcc\x80\xaf\x72\x1a\xa6\x63\x47\x43\x86\xe9\x58\x1f\x4b\x32\x9f\xc3\x60\x7f\xa1\xde\x58\x86\xdd\x65\x79\xe8\x2e\x07\x43\x94\x28\xf9\x6f\x43\x77\x39\x13\x49\xe6\x72\x97\xcf\x60\x6d\x8b\x6a\xe7\xc3\xce\x07\x7d\x20\x00\x5e\xd9\xec\xa7\xcb\x9a\xe7\x57\x2c\x57\xd2\xf9\xdd\x5c\xed\x67\xd7\xe2\xed\x7a\xf9\xf8\x4c\x72\x99\xf6\x6f\xa0\x65\x51\xdd\xc5\x4f\xa1\x67\x50\x07\x65\xe7\x21\xf9\xbd\x3e\x31\xea\xf3\x0b\x7f\xcf\x3f\x11\xc2\xdb\xc1\xbd\x3e\xb0\x68\x05\xef\xde\xe8\xfa\xbb\x05\xdd\xa4\x8c\x3d\x67\xfd\x05\xa7\xd6\x0b\x60\x57\x39\xd6\xb4\x13\x7f\x3a\xa9\x1b\x80\x91\xd8\x38\xca\x05\xc7\x19\x0f\x1e\x87\xf2\xa6\xde\x28\xf3\x09\x7a\xf8\x27\x14\x35\x7c\x38\x55\x54\x1c\x1d\xc5\xcb\xea\xb3\x5b\xd4\xfb\x03\x11\xc2\x28\x43\xfc\x73\xf8\xf8\xff\x33\xe4\x99\xe7\x14\xd3\xff\x13\x13\x58\xba\xb4\x74\xe9\xa7\x00\x00\x00\xff\xff\x38\x12\x33\x60\xd3\x83\x00\x00")

func dataDictionaryJsonBytes() ([]byte, error) {
	return bindataRead(
		_dataDictionaryJson,
		"data/dictionary.json",
	)
}

func dataDictionaryJson() (*asset, error) {
	bytes, err := dataDictionaryJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "data/dictionary.json", size: 33747, mode: os.FileMode(436), modTime: time.Unix(1469636056, 0)}
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
	"data/dictionary.json": dataDictionaryJson,
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
	"data": &bintree{nil, map[string]*bintree{
		"dictionary.json": &bintree{dataDictionaryJson, map[string]*bintree{}},
	}},
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