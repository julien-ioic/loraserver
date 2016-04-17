// Code generated by go-bindata.
// sources:
// ../../static/index.html
// ../../static/js/app.js
// ../../static/partials/api.html
// ../../static/partials/applications.html
// ../../static/partials/navbar.html
// ../../static/partials/nodes.html
// DO NOT EDIT!

package static

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

var _indexHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x54\x5f\x53\x9b\x4e\x14\x7d\xf7\x53\xf0\xdb\x57\x7f\x40\x10\x63\x62\x07\x9c\x89\x86\xd6\x66\x12\x13\x8d\x25\xea\xdb\x06\x2e\xb0\x64\xd9\x85\xdd\x25\x21\x7e\xfa\x2e\x49\xad\xf6\x9f\xb5\x9d\xe9\x13\x7b\xef\x9e\xfb\xe7\x1c\xe0\x78\xff\x0d\xa7\x17\xb7\xf7\xb3\xc0\xc8\x54\x41\xcf\x0e\xbc\xf6\x61\x50\xcc\x52\x1f\x01\x43\x06\x4b\x4d\x5c\x96\x3e\xa2\x5c\x60\x09\x62\x0d\x02\x9d\x1d\x18\x86\x97\x01\x8e\xdb\x83\x3e\x16\xa0\xb0\x11\x65\x58\x48\x50\x3e\xaa\x55\x62\xf6\xd1\xcb\xab\x4c\xa9\xd2\x84\xaa\x26\x6b\x1f\xdd\x99\x9f\x06\xe6\x05\x2f\x4a\xac\xc8\x92\x02\x32\x22\xce\x14\x30\x5d\xf7\x31\xf0\x21\x4e\xe1\x9b\x4a\x86\x0b\xf0\xd1\x9a\xc0\xa6\xe4\x42\xbd\x00\x6f\x48\xac\x32\x3f\x86\x35\x89\xc0\xdc\x05\xff\x1b\x84\x11\x45\x30\x35\x65\x84\x29\xf8\xce\x53\x23\x45\x14\x85\xb3\x31\xbf\xc1\xc6\x7c\xb7\xbf\x67\xef\x53\xfb\x6b\x4a\xd8\xca\x10\x40\x7d\x24\xd5\x96\x82\xcc\x00\xf4\xa0\x4c\x40\xe2\xa3\x76\x71\xf9\xce\xb6\x0b\xdc\x44\x31\xb3\x96\x9c\x2b\xa9\x04\x2e\xdb\x20\xe2\x85\xfd\x35\x61\xbb\x96\x6b\x9d\xd8\x91\x94\xcf\x39\xab\x20\x1a\x25\x25\xd2\x8b\x29\x48\x05\x51\x5b\x3d\x23\xc3\x6e\xff\xd8\x74\xaa\x7e\x71\x3b\x9a\x0e\xe6\x4d\x3f\x77\x06\xf5\x21\xee\x2e\x86\x21\x9b\x91\x23\xba\x7a\x9f\x6c\x36\xc1\x00\xf7\xb3\xe1\x30\xce\x1f\x68\x39\x86\xb4\xc9\xf2\x70\x12\x38\x49\x9a\x2f\x66\x1f\x8a\xd5\xa3\xec\x69\x25\x04\x97\x92\x0b\x92\x12\xe6\x23\xcc\x38\xdb\x16\xbc\x96\xe8\x5f\x93\x32\x55\x06\x05\xbc\x46\x2d\x19\x2f\x8e\xae\x3a\x0e\x9d\x54\x39\x5e\x9d\xaf\x1a\x97\xda\x93\xd3\x00\x67\xf5\xa6\x9c\x27\x70\xb5\x0e\x4f\xdc\x51\x17\x1e\x99\x5b\x3f\x3c\xe2\xf2\xb6\x53\xf7\x82\x7b\x79\x37\xc9\xaf\xc3\xc3\x4e\xc0\xba\xe2\x77\xd4\x64\x24\x48\xa9\x0c\x29\xa2\x67\x2a\x38\xc7\x8d\x95\x72\x9e\x52\xc0\x25\x91\x3b\x1a\x6d\xce\xa6\x64\x29\xed\xbc\xaa\x41\x6c\x6d\xc7\x72\x1c\xcb\xfd\x12\xed\x18\xe4\xba\xa9\x67\xef\x1b\xbe\xd2\xfd\xad\x42\xe5\xdf\xbf\xfc\xfc\xa7\x02\x75\x8a\xf9\x72\x34\x0c\x2e\xf5\xa7\x9a\x14\xf5\xf9\xf9\xf5\xec\x64\x70\x7c\x2d\x4a\x51\x75\xa7\x61\xb2\x70\x7b\xb3\x9b\x1b\x37\xef\x06\xe3\xaa\x91\xd2\xd9\x86\xd5\x54\x31\x28\xd9\x65\x38\x3b\xc5\xa3\x5e\x33\xff\xb5\x40\x6f\xe0\xf2\xba\x52\xfa\xa7\xaf\x29\x16\x9a\x88\x63\x75\xad\xce\x53\xfc\x27\x62\xfd\xd5\x00\x53\xf0\x5a\xc1\x5b\xc6\xe8\x4a\xed\x46\x3f\xa0\x3c\xfb\xc9\x8e\xbc\x25\x8f\xb7\xfb\x42\xc3\x8b\xc9\xba\xf5\xaf\xd6\x3f\x34\x5c\x47\x7b\xec\x1e\xa2\x6b\x5a\xc3\xfb\x1c\x00\x00\xff\xff\x78\x44\x55\x33\x00\x05\x00\x00")

func indexHtmlBytes() ([]byte, error) {
	return bindataRead(
		_indexHtml,
		"index.html",
	)
}

func indexHtml() (*asset, error) {
	bytes, err := indexHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "index.html", size: 1280, mode: os.FileMode(420), modTime: time.Unix(1457866204, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _jsAppJs = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x58\x5f\x6f\xdb\x36\x10\x7f\xf7\xa7\xe0\xbc\x00\x94\x31\x57\x5e\x1f\xe7\x20\x18\xba\x64\x28\x0a\x6c\x43\xd1\x36\x4f\x41\x1e\x38\x91\xb6\x35\xd0\xa2\x40\x52\x09\x8c\xc0\xdf\x7d\x77\xd4\x3f\x4a\xa6\x6d\x29\xe9\xda\x6e\x98\x5e\x12\x91\x77\xc7\xbb\xdf\xdd\xfd\x8e\xd6\x03\xd3\x44\x2a\xcd\x8c\xd0\x0f\x42\x93\x2b\xc2\xb2\x75\x21\x99\x8e\xb7\x8a\x17\x52\x44\xb4\xdd\xa4\x73\x72\x37\x21\xf0\xd0\x6c\xfd\x41\x15\x56\xd0\xf9\xa4\x7c\x6f\x65\xae\x55\x66\xb5\x92\x52\x68\x43\xdd\xde\xfd\xec\x72\x32\x69\xf7\xe3\x44\x65\xab\x74\x1d\xdd\x4d\x2f\x72\xad\x1e\x52\x2e\xa6\x73\xb2\x2a\xb2\xc4\xa6\x2a\x8b\xea\xb5\x19\x79\x72\xca\x5a\xd8\x42\x67\xa4\x5e\x8e\xb9\x48\xc0\x92\x55\x3a\xa2\x17\x1b\x6b\x73\xf4\x88\x5e\x70\x21\xc5\x9a\xa1\x3b\x9e\xa5\x7a\xb1\x36\x85\x4f\xb3\x18\xeb\x3c\x81\x50\x1b\xe9\xad\xb0\x1b\xc5\xe7\x24\x67\x9a\xc1\xff\xe0\xbc\xaf\x86\x0f\xc2\xc4\x99\x65\xa0\xf5\x34\x2d\xc5\xa7\x4b\x52\xeb\x4d\x9d\xa2\x81\x95\xbb\xd6\xc4\x3d\xac\xa7\x7c\x4a\x96\xe4\xf5\xfe\xb2\x63\xac\x0e\xab\x71\x27\x57\xc6\x46\x74\x01\x5e\x41\x0c\x78\xcc\x9c\x3c\xd1\x8d\x60\x1c\x61\x5c\x3e\x51\x44\x55\x64\xf6\xd5\xa7\x5d\x2e\xe8\x92\x50\x96\xe7\x32\x4d\x18\xfa\xbe\xf8\xcb\xa8\x8c\xee\xf7\xb3\xf6\x08\xef\xb4\xfe\x49\x9e\xd0\x7d\xa5\xb1\x3f\x96\x21\x7a\xa1\x31\xc9\xef\x4b\xec\x21\xfb\x4e\xbc\x45\xb8\xb3\xdb\x81\xb9\xb3\x13\x77\x42\x7f\xdc\x88\x0c\x22\xf5\x02\x30\x10\x72\x17\x6b\x7c\xac\xd8\xe6\x12\x1c\xbe\xd5\x12\xe2\x05\x50\x6d\xca\xa4\xe9\xe8\xc5\x1b\xbb\x95\x95\x57\xfe\x93\x34\x25\x08\x9a\x6f\x5a\x85\xdf\x52\x63\xaf\xad\x96\xb4\xa3\xb1\x9f\x9d\x75\x70\xb1\xf4\xde\xbe\x49\x6f\x33\xc5\xc5\x28\x1c\x9d\xc2\x20\x97\xfe\x00\xc9\xf1\xbe\x2c\x96\xf8\xe7\x1b\xf1\x88\xe5\xe9\xb8\xac\xa5\xc3\x92\xf5\xfe\xdd\x79\x17\x94\xdd\x08\xfd\x98\x1a\x11\x1d\x9e\xaf\x05\x4f\xb5\x48\xec\x27\x05\xd6\xba\x2d\xd1\xb3\xd9\x6b\xdb\x49\x97\xb4\x3d\xce\x3d\xc9\xdf\x3e\x37\x03\x71\x3a\x4b\x8b\x05\xe1\xa9\x01\x18\x76\x04\x39\x91\xab\xc4\x4c\x82\x1a\x71\x1b\x7b\xd4\x84\xee\xe8\xd7\x24\x2a\xc7\x54\xd7\x94\xdc\xa1\x09\x12\x95\xfb\x73\xe2\x76\x3b\x34\xe1\x36\xe2\x9c\xad\x05\xb8\x4d\x31\x4b\x6d\x9c\x4e\x3a\x5e\x8b\x9a\x16\x67\xb1\x29\x92\x44\x18\x13\x35\x04\x84\x3c\xd9\xa7\xe9\xca\x26\xd8\x42\x28\x50\xc2\x83\xae\x43\x78\x10\xf8\x96\x65\x78\xb6\x0f\xfc\x90\xd8\x03\x3d\x1a\xc4\x81\xd4\x04\xea\x66\x43\xfb\x7e\x1a\xa1\x39\xf1\xb5\xea\xb7\x93\xb8\x79\x75\xd3\x07\x10\xa0\xeb\x78\x1c\xbf\x15\x16\xbd\xc6\x7e\xa0\x32\xdd\xa6\x16\xc6\xc9\x4f\xf0\x80\x77\x6a\xb5\x32\x02\xdf\x7f\xdc\x0f\x44\xbb\x83\x78\x5e\x23\x1e\x6b\x61\x0a\x69\xbb\xc0\xf7\xbd\x4f\xb4\x80\xce\xf3\x67\x30\x58\xe8\xdb\x4f\x57\xb8\x4a\xae\xae\x48\x56\x48\x19\x3c\x3d\xa2\xdf\x97\xa6\x7e\x57\x9c\x49\x28\x93\x2d\xfe\x8d\x66\x31\x58\xa4\x9b\x94\x73\x91\xc5\x7f\x9a\x72\xd5\xbf\x1f\x84\x8c\x39\x83\x0e\x6e\x08\x41\x2a\xc6\xa3\xd9\xe5\x81\xd0\xbe\xb7\xb6\x27\x42\x1a\x11\x72\x2d\x9c\x80\x6b\xe7\x2d\xb8\x82\xf1\x0e\x87\xb9\x84\x83\xb8\xfd\x58\x68\xad\xf4\x29\x58\xce\xc1\x83\xc8\x08\x1a\x88\xce\x45\x14\x06\xa6\xcc\x5b\x75\x34\x69\xfd\x18\x00\x91\x7f\x33\xe9\x17\x42\x91\xf3\xf3\x85\x70\x04\xcb\x5b\xa7\x3b\x1a\xcb\x71\x38\x22\x86\x40\xd4\x76\x20\x82\x87\xe8\x0d\x44\xae\x77\x83\xeb\xe3\x84\x37\xb8\x01\x0d\x43\x22\x77\x7d\xd3\x5b\x40\x4a\x0b\xb2\x53\x05\x31\x45\xf5\xcf\x23\xcb\x2c\xb1\x8a\x54\xa6\x28\xf9\x01\x81\xc3\xd6\xfd\xf5\xf6\x1d\xbc\xd0\x9f\xe9\x2c\xd8\x64\x61\xf4\x6f\x9c\x99\x12\xfd\xca\xc8\xcb\x0a\xfa\xbb\xf3\x05\xcd\x80\x82\xad\xa7\x33\xae\x84\x47\xf7\x76\x30\x21\xc0\x4a\x3e\x45\xc7\x1e\x01\x0f\xac\x5b\x20\x61\xda\xe5\xf9\x8e\x91\xe7\x90\xef\x31\xee\x6d\x24\x83\x55\x3c\x84\x26\xc9\x06\xe6\x85\xd2\xbb\x78\xad\xa2\x57\xaf\x67\x97\x87\x28\x1d\xfe\xf4\xe8\x8f\x58\x77\xb9\x1b\x30\x5b\x3b\x57\xbb\x2f\x3b\x54\xdb\x4c\xa1\x13\x9f\x77\x4e\x56\x69\x72\x28\xbc\x6c\x48\xa2\x89\x50\xd3\xe3\xfa\x3f\x33\x27\xbf\xe0\x98\x74\xc0\x37\xf3\xd1\x85\xfa\xff\x80\x0c\x0c\xc8\x50\x11\xf4\x61\x6c\x46\xe3\x48\x18\xff\x23\xb3\xf1\x58\x9f\x8c\x1a\x8e\x68\x04\x6c\x3f\x0c\x9e\x8e\x0e\xf9\x66\x2c\x7a\xea\xe3\xea\xf8\xdf\x39\x16\xab\x5c\x18\x88\x12\x09\x78\x5c\xbd\x7e\x2c\xb5\x90\x76\x7f\xd9\xdd\x38\xcc\x9e\x8d\x60\xcd\xb6\x47\xa9\xb6\x7e\x10\x67\x4f\xf8\x4c\xa9\x7b\x56\x8f\x27\xa2\x74\x76\xe9\x7b\x7e\xf8\x15\xa1\x7e\xca\x3b\x53\x25\x5c\xbe\x1c\x17\x5e\x5d\x67\xf6\x36\x87\xa1\x73\x5a\xe4\x46\x3d\x66\x20\x14\x4e\xff\xa0\x8e\x84\x96\xae\x72\xf8\xd5\x7e\x52\x9d\xee\xf6\x92\x10\x3f\x06\xea\xec\xe0\xab\xed\x91\x2a\x6b\xc9\xd1\x0c\xaf\xab\xf1\xcc\x18\x84\xf1\xb3\x90\x23\x19\x83\xd7\x5a\xd8\x0f\x2c\xe3\x6a\x0b\x7d\xf5\x86\x73\xfd\x2c\xc8\xde\xf6\x8c\x20\x78\x18\xfe\xf0\xb6\xcc\x0c\x36\x44\xe5\xc0\xe9\xbb\xea\xcb\x86\x42\xff\x7e\x3e\x64\x5c\x06\x6e\xe4\x23\x07\xa7\x77\xc7\x3b\x17\x1f\x0c\x21\xa3\xa4\x88\xa5\x5a\x97\xe6\xbe\xf6\x75\xdd\xf3\xbf\xfe\xaa\x54\x7e\x51\x6e\x6f\xf3\x7f\x07\x00\x00\xff\xff\xdc\xbb\xf2\xfd\x27\x1a\x00\x00")

func jsAppJsBytes() ([]byte, error) {
	return bindataRead(
		_jsAppJs,
		"js/app.js",
	)
}

func jsAppJs() (*asset, error) {
	bytes, err := jsAppJsBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "js/app.js", size: 6695, mode: os.FileMode(420), modTime: time.Unix(1458469975, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _partialsApiHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xbc\x54\xcb\x6e\xdb\x30\x10\xbc\xfb\x2b\x58\xe5\xe0\x04\xa8\x44\xa0\xe9\x29\xa0\x04\x14\x3d\xb5\x87\xd4\x48\xfb\x03\x6b\x69\x23\x12\xa5\x48\x75\x49\x19\x30\x02\xff\x7b\x49\xc9\xb2\xa5\x58\x70\x1f\x87\x5c\x4c\x7a\x39\x3b\x3b\x9c\xa1\x2d\x4c\x9d\x2a\x53\xea\xae\x42\xe6\xa8\xcc\x93\x75\x0b\xe4\x15\x68\xc7\x0d\xec\xb6\x40\x99\xf4\x8d\x5e\x27\x85\xe0\x67\x64\xb1\x5a\x89\x4a\xed\x58\xa9\xc1\xb9\x3c\x29\xad\xf1\xa0\x0c\x52\x52\xac\x18\x13\xf2\x43\xf1\xf5\xfb\xb7\xc7\xf4\x69\xf3\x99\x7d\xda\x7c\x11\x3c\x14\x62\x3d\x76\x04\x0e\xc2\x16\xc1\xe7\xc9\x2d\xb4\xea\x11\x1a\x7c\xcf\xc2\xe6\x8e\x29\x13\x57\xd7\x53\x0c\x24\x2f\x2f\xec\x08\x61\x87\x03\x6b\xd0\x4b\x5b\xb9\x91\x6d\x81\x6f\x40\x0c\x94\xc3\x7e\x64\xcd\x8e\xcd\x47\xf2\x48\xff\xb1\x10\xc0\x24\xe1\x73\x9e\x48\xef\x5b\xf7\xc0\x79\x6d\x2b\x5b\x66\x96\x6a\x5e\x2b\x2f\xbb\x6d\x56\xda\x86\x6f\xc9\x96\x00\xc4\xb5\x25\x70\x48\x3b\xa4\x9b\x99\xae\x70\xc1\x2c\x14\xce\xb3\x43\x2d\x99\x4b\xbf\x38\x17\x1c\x82\x9d\x41\xc2\xa8\x66\xe2\x65\x0b\x06\x35\xeb\x3f\xd3\x0a\x9f\xa1\xd3\xfe\xa4\x7a\x01\x99\x4a\x84\x4a\x99\x7a\x82\x89\xb7\xbb\x9f\x83\xbc\xf2\x1a\x93\xe2\x09\x7f\x75\xe8\xfc\x03\x13\x2e\xd4\x4f\x90\x4e\xeb\x94\x54\x2d\x7d\xaf\x7b\x50\x9a\x01\xd5\x5d\x83\xc6\xff\xd8\xb7\x38\x46\x30\xb4\xc5\x67\x10\x5c\x7b\x85\xdb\xfc\xac\x37\xe0\x25\x7b\x97\xb3\x75\x78\x2d\xb7\x57\xdc\xbd\x1c\x32\x36\x1f\x0e\x37\x57\x15\x2c\x09\x3c\xf7\x46\x5b\xef\x04\x8f\x22\x8b\xd3\x22\xef\x27\xee\xf1\x60\xdf\x35\x33\xb7\xb6\xda\xcf\x9d\x6c\x09\x17\x66\xc6\xe7\xdd\x0f\x8c\xc7\xcb\xf4\xf3\x2f\x6f\x16\xb0\x6b\xad\x71\xf8\x57\x09\x87\xdf\x8d\xde\xff\x29\xde\x1e\xf4\x5f\xd9\xce\x3a\x67\xc1\xbe\x1e\x7c\x21\xea\x2d\x23\xed\x07\xfe\x5b\x9e\xd3\xad\xa4\xfe\xaf\x6d\xa8\x0c\xcb\xef\x00\x00\x00\xff\xff\x13\x14\x53\x6f\x52\x05\x00\x00")

func partialsApiHtmlBytes() ([]byte, error) {
	return bindataRead(
		_partialsApiHtml,
		"partials/api.html",
	)
}

func partialsApiHtml() (*asset, error) {
	bytes, err := partialsApiHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "partials/api.html", size: 1362, mode: os.FileMode(420), modTime: time.Unix(1458470541, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _partialsApplicationsHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x56\x4f\x6f\x9c\x3e\x10\xbd\xef\xa7\xb0\xfc\x93\x7e\x49\x0f\x2c\x6a\x95\x53\x0b\x48\x51\x9a\x43\xa5\xb6\x97\xaa\x1f\x60\xc0\xde\xc5\xaa\xb1\x91\x31\xab\xac\xa2\x7c\xf7\x8e\x6d\x60\x0d\x4b\x92\xe6\xd0\xaa\xaa\x72\x48\xf0\x9f\xc7\x78\xe6\xf9\xcd\x5b\x32\xb5\x4f\x84\xaa\x64\xcf\x38\xe9\x4c\x95\xd3\x8b\x16\x8c\x15\x20\xbb\x54\xc1\xa1\x04\xb3\xad\x6d\x23\x2f\x68\x91\xa5\x27\x64\xb1\xd9\x64\x4c\x1c\x48\x25\xa1\xeb\x72\x5a\x69\x65\x41\x28\x6e\x68\xb1\x21\x24\xab\xdf\x15\xd7\x6d\x2b\x45\x05\x56\x68\xd5\x65\x29\x2e\xe0\x3a\x21\x6e\xd3\x42\x29\xf9\xf8\x62\x98\xf8\xff\x49\xad\x0f\x43\x00\x87\xaa\x39\xb0\x30\xf6\x33\x17\xf0\xf6\xfb\xa7\x2c\xc5\x61\xb4\xfa\x15\x1a\xbe\x58\x9b\x42\xf3\x3b\x9b\x18\xb1\xaf\x2d\x2d\xae\xab\x21\x91\x11\xe9\x46\xd3\x01\x99\x2d\x35\x3b\x9e\x42\x18\x82\x85\x1a\xde\x72\xb0\x39\x85\xb6\x25\x42\x11\x7c\x74\x74\x84\x38\x10\x2b\xee\xef\xdd\xea\x16\x7c\x62\xe4\xe1\x01\x63\xb2\x55\x84\xc2\x24\xd7\xf6\xd7\x32\x9d\xf6\x11\x01\x23\xa0\xb4\x8a\xe0\x5f\xd2\x35\xfe\xc1\xf8\x0e\x7a\x69\x29\xa9\x0d\xdf\xe5\x34\xfd\x2f\x85\x88\xed\x74\x99\x17\x2d\x6e\x99\xb0\x59\x0a\xb3\xe0\x65\x6f\xad\x56\x8f\x9d\x00\x6a\x8f\x77\xe1\x78\xa8\x30\xf0\x8f\x9c\x32\x2e\xb9\xe5\x97\x18\xf4\x0d\x2d\x3e\xfa\xc9\x2c\x62\x5c\x1c\x8e\xcd\x44\xf3\x48\x2d\x0e\xdd\x2d\xa3\x70\x70\x1c\x49\xc7\x1d\xb7\x37\xba\x6f\x29\x31\x5a\xf2\x9c\x0e\x13\x30\x02\x12\x09\x25\x97\x39\xdd\x6e\xb7\xa3\x2e\x86\xb4\xed\xb1\x45\x68\x98\xd0\x65\x11\x13\x3f\xa7\xf4\x2b\x83\x97\xc9\x2f\x31\xf7\x1b\x3f\x22\x11\x63\x59\x1a\xe2\x84\x2c\x31\xb5\x62\x33\x3c\x66\x1a\x6f\x34\x03\x49\x76\xc0\x38\x25\x82\xe5\x94\x23\xa9\x5f\xdc\x1a\x75\xfa\x15\x8a\xf1\xbb\x9c\x26\x6f\xc7\x32\x18\x76\x90\xde\xc7\x75\x48\xce\xca\x23\xc6\x39\xfa\xb7\x3e\xbb\xa5\xd0\x2e\xcb\x43\x92\xf1\xdd\x21\x92\xae\xfa\x86\xab\x51\x1c\xe7\x70\xd7\x7d\xa7\xfd\x35\x84\xd3\xfa\xd4\x5b\xcf\xf2\x58\x49\xdd\x61\x91\x0c\x2c\x60\x2a\x5d\x23\xa6\x40\xf3\x6b\xb9\xf1\xb8\x22\xeb\x5a\x50\x61\xa3\x16\x8c\x71\x85\x8a\x36\x3d\x6e\xfc\x6f\x45\xc3\xbb\x0f\x59\xea\x00\x45\x4c\xf3\x90\x44\x7d\x35\x4f\xd3\x0a\x2b\x07\x76\xe7\x2c\x39\x01\x93\xf3\x7e\xab\xaf\x4e\x92\xf3\x17\xf6\x58\xfd\x4e\x84\x71\xf5\x3b\x6d\x9a\x59\x37\x44\x6f\xb8\xbd\x41\x91\x31\x04\x41\xbe\x6c\x82\xfb\x39\x75\x2d\x4d\x63\xf7\xc3\xab\x0a\xb4\x50\xef\x49\xef\xb3\xd4\xcf\x16\x11\x84\x6a\x7b\x3b\x90\xee\xba\x9e\xce\x0e\x1d\xe2\x04\x02\xc2\x09\xa8\x60\xac\xc0\x91\x3d\x1a\xc9\xdc\x22\xe2\xb2\xdd\x34\x2e\xec\x19\x4e\x76\x5a\xdb\xb9\x26\x22\x48\xdb\x4b\xac\x87\xef\x30\x59\x67\x4e\x91\x1f\x08\x74\x1c\x6e\x8c\xc6\x37\xf1\x42\xfc\xc8\xdf\xc5\x3c\x91\x17\xb5\xe9\x9a\xd0\x0a\x2f\xae\x15\xcd\xfc\x4a\xe4\xd6\x88\x06\xcc\x31\x36\x80\xbe\x65\x30\xf9\xd7\x37\x38\xe0\x2f\x50\xed\x8a\xea\x96\x47\x44\x95\x4c\xc3\x17\xb8\x42\x30\x9a\x57\x5f\xf8\x33\xbe\xb0\x66\xe6\x7f\xa7\x29\x84\xcf\x97\xdf\x63\x0b\xc1\x11\x9f\x32\x86\x57\x87\x3b\x83\xfc\x73\x0e\x37\x7c\xe2\x04\x87\x7b\xfa\x2b\x67\xc1\xdd\xba\xcf\xfd\x0c\x00\x00\xff\xff\xce\x40\x8a\x4e\x18\x0c\x00\x00")

func partialsApplicationsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_partialsApplicationsHtml,
		"partials/applications.html",
	)
}

func partialsApplicationsHtml() (*asset, error) {
	bytes, err := partialsApplicationsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "partials/applications.html", size: 3096, mode: os.FileMode(420), modTime: time.Unix(1457866204, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _partialsNavbarHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xb4\x51\xc1\x4e\xc3\x30\x0c\xbd\xf3\x15\x91\x77\xd8\xa9\xca\x1d\xb5\x95\x38\x22\x21\x84\x80\x1f\xf0\x1a\xb7\xb3\x14\x25\x55\x9a\x4e\x4c\x88\x7f\xc7\x5d\xd3\x30\x56\x10\x70\xe0\x94\xf8\xf9\xd9\xcf\xcf\x2e\x1d\x1e\x54\x63\x71\x18\x2a\x90\xef\x0e\x83\x9a\x9f\xc2\x50\x8b\xa3\x8d\x50\x5f\x29\x55\x1a\xce\xac\xc6\xbb\x88\xec\x28\x14\xad\x1d\xd9\x9c\xf2\x9f\x19\xa9\xc1\x9e\xd0\x50\x48\x79\x61\xec\xc6\x18\xbd\x53\xf1\xd8\x53\x05\x73\x00\x17\x25\xd1\x77\x9d\x25\xd5\x78\x6b\xb1\x1f\xc8\x80\x32\x18\x31\xc1\x93\xf4\x8c\x2f\x30\x86\x8e\x62\x05\x9b\x54\xfd\x91\xc6\xc0\x58\xd0\x4b\x8f\xce\x90\xa9\xa0\x45\x2b\x68\x1e\x64\x10\x7c\x11\x1e\x42\xe1\x9d\x3d\x42\xfd\x3c\x4b\x4b\x2b\xee\x30\xb2\x77\xa5\x9e\x78\x5f\x16\xb1\xec\xa0\x10\x45\xa8\xff\x8d\xa4\xe7\x05\xe5\x18\x2f\x36\xb5\x0b\xe2\x0d\xd4\x3e\x50\x5b\x81\x86\xfa\xce\x3f\xa2\x7a\xa2\x70\xa0\x50\x6a\x4c\x37\xd1\x72\x94\xf5\x79\x96\x35\xa9\xd5\xda\xd8\xe4\xfe\x19\xcc\x13\x8c\xf6\x6c\x84\xa5\x56\x9e\xcc\x10\x8e\x65\xe5\xba\x22\xd1\x5e\xb1\x89\x7c\xa0\x6b\xd5\x63\x47\xaa\xaa\xd4\x16\xfb\xde\x72\x73\xda\xee\xb0\x7d\x13\xcf\xb8\x18\xd8\xe8\xf3\x1c\xd4\x37\x67\xd1\xe4\xa7\xd4\x96\x7f\xaf\xe3\xbc\xa1\x95\xc0\x09\x84\xfa\x7e\x7a\x2e\x5b\x96\x7a\xb4\x3f\x1a\x5d\xbe\x81\xbb\x7d\xfc\x93\x6b\x5e\x9b\x65\xf1\xf8\x70\xfb\xfd\x1c\xf9\x76\xe9\x53\x6a\x51\xaf\xdf\x03\x00\x00\xff\xff\x41\x3c\x84\x3a\xae\x03\x00\x00")

func partialsNavbarHtmlBytes() ([]byte, error) {
	return bindataRead(
		_partialsNavbarHtml,
		"partials/navbar.html",
	)
}

func partialsNavbarHtml() (*asset, error) {
	bytes, err := partialsNavbarHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "partials/navbar.html", size: 942, mode: os.FileMode(420), modTime: time.Unix(1458333988, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _partialsNodesHtml = []byte("\x1f\x8b\x08\x00\x00\x09\x6e\x88\x00\xff\xec\x58\x5f\x6f\xdb\x36\x10\x7f\xcf\xa7\x20\x58\x60\x4d\x1f\x6c\x61\x43\x9f\x36\x49\x40\x96\x74\xc0\xb0\xad\x18\x66\xf4\x03\xd0\x22\x6d\x11\x95\x48\x81\xa4\xdc\x18\x45\xbf\xfb\x8e\xff\x6c\xca\x91\x9d\xb8\xd1\x12\x60\xf0\x43\xe2\xe3\xf1\x74\x77\xbf\xe3\xdd\xcf\x94\x73\xb1\x9e\x71\x51\x35\x3d\x65\x48\xab\xaa\xc0\x6f\x3b\xa2\x0c\x27\x8d\xce\x04\xd9\x2c\x89\x9a\xd7\xa6\x6d\xde\xe2\x32\xcf\xf6\x96\xe5\xd5\x55\x4e\xf9\x06\x55\x0d\xd1\xba\xc0\x95\x14\x86\x70\xc1\x14\x2e\xaf\x10\xca\xeb\x9f\xca\x8f\x92\x32\x9d\x67\x20\x59\x85\x21\xcb\x86\x45\x63\xbf\x70\xff\x67\xb5\xdc\x84\x87\xac\x55\xcd\x08\xf5\xb2\x5b\x95\x77\x6c\xf3\xe1\xd3\xef\x79\x06\x62\xa2\xbd\xe9\xba\x71\xed\x1f\x6c\x7b\xa0\xdd\x85\x64\xf7\x66\xa6\xf8\xba\x36\xb8\xbc\xa9\x0c\x97\x42\xef\x2d\xad\xb4\x0b\x9c\x9b\xa5\xa4\xdb\xbd\x0b\x85\x00\xb4\x62\x1d\x23\xa6\xc0\x02\x40\x21\x2e\x90\xfd\xd4\x38\x1a\x59\x33\x5a\x7e\xfd\xea\xd4\x73\xea\x72\x46\xdf\xbe\x81\x5b\x3a\x6e\x42\x1c\x80\xc7\x4c\x00\xcd\x98\xc9\x18\xa4\xdd\x3e\x58\x90\x68\xb0\x34\x02\xc1\xdf\x4c\xb7\xee\x83\xb2\x15\xe9\x1b\x83\x2d\x9e\xaa\xe1\xd5\xe7\x02\x6b\xa6\x35\x94\xe2\xda\x46\x7c\x87\xcb\x85\x5f\xa2\x0c\xdd\xfc\xfa\x77\x9e\x91\xf3\xdc\xd6\x8a\xad\x0a\x9c\xbd\xc9\x5c\x71\xb2\x07\xe5\xc0\x88\x12\x43\x66\x46\xae\xd7\x0d\x2b\x70\x2b\x29\x69\xa2\x8e\xa8\x35\x83\xf2\xbe\x69\xb7\x7f\x39\x75\xf9\x81\x72\x73\x98\xc1\xb2\x37\x06\xb2\x3b\x92\x06\x11\x6b\x68\xa4\x04\x1c\x65\x0d\x33\x2c\x62\xbb\x73\xab\x81\xcb\xb4\xb2\x20\xab\x5d\x33\xc4\x06\x00\xd1\xf6\x28\xb4\x3a\xc8\x49\xb3\xdb\x78\x6b\x25\xfb\x0e\x23\x25\x2d\x96\xb0\x20\x8a\x93\x59\x43\x96\xac\x29\xf0\x7c\x3e\x8f\x5d\x1d\xf2\x36\xdb\x0e\x4c\xfd\x02\x1f\xa2\x18\x39\x9c\x4a\x41\xcb\xb1\x6b\xc8\xfd\xd6\x49\xae\x9c\x79\xe6\x1d\xf8\xf4\x20\xa7\xf2\x2a\x7c\x0c\xc6\xd1\x15\x17\xad\x08\x65\x18\x71\x5a\x60\x06\xe5\xf4\x95\xb5\x63\xc7\x05\x65\xf7\x05\x9e\xfd\x18\xf3\xa7\x30\xec\x72\x9d\x02\x68\x18\x5d\x6e\xc1\x8f\x3f\x8f\x3f\xad\xca\x4f\xf6\x61\x90\x59\x7c\x36\x78\x92\x55\xdf\x32\x11\x5b\xf2\xa1\xb9\x25\x8a\xfd\xfe\x98\x85\x1d\xc5\x1d\x25\x3c\x5a\xc0\xaa\x91\x9a\x85\x36\xa2\x5c\xb7\x7c\xe7\x68\x78\x1e\xb7\xce\xae\xcc\x75\x47\x84\xdf\xa8\x39\xa5\x4c\xc0\x1c\xa9\x1e\x36\x7e\x30\xbc\x65\xfa\x97\x3c\xb3\x06\x65\x5a\xe6\x90\x44\xfd\x7e\x98\xa6\xe1\xa6\x09\xd5\x1d\x56\xc9\xb6\x2e\x1a\x99\xf4\xfa\xfd\xbe\xd9\xdc\x89\x1d\x2b\x80\x6d\xbf\x14\xfe\x4a\xaa\x76\x30\x08\xc9\x13\x76\x2f\xf4\x62\x6a\x02\x46\x0e\x37\x82\xfd\x02\xfb\x24\x70\xca\xd5\x70\x5a\xbe\x32\x38\xb0\xe9\xcf\x79\xe6\xd6\x07\x5e\xb8\xe8\x7a\x13\x2a\x6f\x09\x07\x0f\x02\x07\x4f\xbe\x0a\x31\x0a\x74\x30\xe0\xb0\x35\x4f\x2a\x30\xa4\xa8\x14\xff\xf7\x22\x02\x76\x3c\x81\x08\x76\x27\x41\xe4\xa2\x3c\x44\x64\xd5\x27\x10\xe5\x59\x7a\x66\x8f\x1c\xf7\x4a\x4a\x93\xf4\x7b\x6a\xd0\xf5\x0d\x80\x62\x2b\xc8\xd7\x92\x7d\xc2\x71\x1c\xa8\x96\x29\x25\xe1\x39\x68\x35\x27\xb9\x26\x1b\xa6\x71\x16\xf3\x8c\x8d\x50\xe9\xc6\x66\x64\x1a\x9e\xe2\xb9\x53\xbc\x25\x6a\x9b\x72\x5a\xdf\x51\xb2\xe7\xe4\x05\xd9\xc0\x9d\xa0\xb6\xa8\xf4\x61\x8c\x04\xca\x4e\x3c\x83\xf0\x3c\x79\x5e\x28\xef\x65\x28\x6f\xf0\x05\xf5\x7a\x3c\xe7\xaf\x1a\xc7\x58\xc1\xdf\x25\x9f\xcf\x0a\x31\xca\x01\x2b\x04\xf5\xe4\x3c\x77\x61\xee\xef\x43\x74\x61\xee\x49\x99\x3b\xdc\x46\xff\x7b\xe6\x0e\xef\x24\x17\xea\x7e\x19\xea\xb6\xaf\xe9\x28\xd4\xdc\x4e\xa8\xbb\xb9\xea\xf4\x25\xf6\x55\xf9\xfc\x86\x52\x75\x82\xd0\xed\xf6\x11\xb6\x38\x63\x82\x9c\x7c\xaf\x91\x1b\x5a\xff\x46\x9d\x74\x3e\xbc\x99\xfe\x43\x04\x95\x6d\x88\x77\x2d\x34\x4c\xc0\x9a\x09\xa6\x88\x19\x99\xb0\x73\xb9\x6a\x87\x32\x21\x2b\x77\x00\x4e\x3b\x39\xfb\x2e\x4e\xd3\xef\x62\x2a\xfe\x5d\x1c\x12\xb0\x9e\x47\xed\xb4\x98\xc4\x97\xcf\xa7\x30\x7d\xf4\xdb\xcf\xc7\xb4\x8b\x33\xc0\x14\xb5\xd3\x62\x5a\xdd\x0a\xf3\xa9\x3b\x06\xe9\x37\xd8\x45\xd7\x7d\xf7\xee\x09\xa0\x44\xdf\x2e\x99\x3a\x01\x2b\x86\x1a\xa0\x0a\xca\xe9\x41\xdd\xc9\x2f\xe2\x34\x2c\x0a\x16\x93\x01\xf3\xe1\x1e\x40\x73\xea\x69\xc1\xbd\xda\xdd\x33\x72\x35\x7c\xab\x70\x6d\x7f\xae\xa2\xff\x93\x2b\xa8\x8e\x17\xd0\xa7\x00\xbb\xdc\xdb\x4e\xbe\x71\x2f\xe2\x0f\xbd\xfa\xd9\x97\xb7\x7f\x03\x00\x00\xff\xff\x1c\xfe\xd3\x2a\x2e\x18\x00\x00")

func partialsNodesHtmlBytes() ([]byte, error) {
	return bindataRead(
		_partialsNodesHtml,
		"partials/nodes.html",
	)
}

func partialsNodesHtml() (*asset, error) {
	bytes, err := partialsNodesHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "partials/nodes.html", size: 6190, mode: os.FileMode(420), modTime: time.Unix(1458903022, 0)}
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
	"index.html": indexHtml,
	"js/app.js": jsAppJs,
	"partials/api.html": partialsApiHtml,
	"partials/applications.html": partialsApplicationsHtml,
	"partials/navbar.html": partialsNavbarHtml,
	"partials/nodes.html": partialsNodesHtml,
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
	"index.html": &bintree{indexHtml, map[string]*bintree{}},
	"js": &bintree{nil, map[string]*bintree{
		"app.js": &bintree{jsAppJs, map[string]*bintree{}},
	}},
	"partials": &bintree{nil, map[string]*bintree{
		"api.html": &bintree{partialsApiHtml, map[string]*bintree{}},
		"applications.html": &bintree{partialsApplicationsHtml, map[string]*bintree{}},
		"navbar.html": &bintree{partialsNavbarHtml, map[string]*bintree{}},
		"nodes.html": &bintree{partialsNodesHtml, map[string]*bintree{}},
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
