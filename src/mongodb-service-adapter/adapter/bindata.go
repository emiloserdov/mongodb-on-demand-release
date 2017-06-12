// Code generated by go-bindata.
// sources:
// om_cluster_docs/replica_set.json
// om_cluster_docs/sharded_set.json
// om_cluster_docs/standalone.json
// DO NOT EDIT!

package adapter

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

var _replica_setJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x57\x4b\x53\xdb\x3a\x14\xde\xe7\x57\x78\xd4\x2d\x49\x4c\x78\x5c\xc8\xaa\x14\x16\x5d\x5c\x2e\x1d\x68\xcb\xa2\x65\x32\xb2\x75\x6c\xeb\x22\x4b\x1e\x49\x06\x82\xc7\xff\xfd\x8e\x64\xc7\xb1\xfc\x80\x74\xa6\x5c\xca\x2e\x47\x47\xe7\xf1\xe9\xfb\x8e\x0f\xc5\xc4\xf3\x3c\x0f\x89\x4c\x53\xc1\x15\x5a\x7a\x95\xc1\x1a\x89\x78\xe4\x4c\x60\xf2\x09\x2b\x40\x4b\x0f\xcd\x1f\xb0\x9c\x33\x1a\xcc\x53\xc1\x63\x41\x82\x69\x9a\xaa\x29\xce\xb5\x48\xb1\xb9\x8d\xf6\x86\xaf\xde\x52\x4e\xc4\xa3\x89\x8d\xce\x97\x3f\x7f\x8e\x5c\xb6\x77\xcb\x2a\x04\xb2\x3e\x17\xc1\x77\x90\xaa\x2e\xeb\x47\xab\xae\x20\xa7\x8c\x58\x63\x63\xdb\xfc\x15\x3d\x4b\x75\x83\x6a\xe3\x7f\x7c\xb8\x37\x7c\x1e\x31\xfc\x20\xa4\xa9\x10\x8d\x78\xc4\x54\xd7\xd5\x18\xaf\xc3\xc5\xe1\x69\xb8\x4f\x16\xc1\xd1\xe9\xe9\x29\x04\x41\xb4\x1f\x91\xe8\x20\x08\x7d\xf0\xe1\x20\x38\x88\xa2\xa3\xd0\xc7\x38\x5a\x8c\x45\x4b\xf1\xd3\x95\x6a\xc5\x1b\xf5\xa3\x7c\x37\x3f\x41\x72\x06\x16\x92\xbb\x11\x97\x8c\x61\x1d\x09\x99\x9a\x30\x42\x3d\x8d\x45\xca\x25\x33\x1e\x89\xd6\x99\x5a\xce\xe7\x11\x56\x9a\xb0\x59\xfd\x66\x33\x21\xe3\x39\xa3\x3c\x7f\x6a\x28\x60\x7f\x4d\x9f\x4e\x8e\x57\xc7\x87\x53\x9c\xe2\x67\xc1\xa7\x07\xb3\xc5\xec\xaf\x99\x8e\x9f\xc7\x92\x3c\x52\xbe\xf0\xfd\x93\x8c\xe5\xa6\xe4\x08\x33\x05\xe3\x9e\xdf\xcf\xaf\x81\x50\xa5\x2f\x18\x7b\x09\x82\x96\xe7\x55\x43\xe6\x51\x34\x5a\xde\xdf\xe4\xae\x71\xdb\x0f\xd1\x73\x2e\x1d\x4b\x2b\x2f\xe2\x38\xb5\xf2\xb1\xb0\xd4\x3c\xaf\xcf\x51\x80\xc3\xfb\x3c\x1b\xe1\x79\x22\x94\xde\x5c\x2e\x0a\x2e\x08\xa8\xd9\x0f\xff\xae\x2c\xdb\x4a\x63\x22\xfe\x82\x75\xb2\xd5\xa7\x88\x47\xf4\x39\xaf\x92\x4d\x71\x0c\x5c\xcf\x98\x88\x3b\x61\xae\x85\xc6\x1a\x9c\x11\x60\x8f\x14\x7d\x86\xaf\x89\x04\x95\x08\x46\x2e\x3f\xa1\xa5\xb7\xef\xfb\xbe\x0b\x17\xd2\x34\xdd\x3a\x7d\x96\xa6\x95\xc5\xe1\xc4\x05\xc7\x74\xdd\xe8\x9b\x6a\x21\x29\x8f\xff\xa7\xd6\xb7\x09\xdf\xb7\x7d\xeb\x9c\x49\x11\x82\x52\x95\x5e\x8b\xe2\x03\xe0\x30\xf1\x6c\x8f\xe5\x86\x45\x2d\x28\xb0\x8c\xd5\x62\x75\xdc\x2f\x8c\x83\xee\x19\xab\xf8\x42\x9a\x93\xc5\x89\xef\xfb\xce\x69\xd9\xa9\x5a\x42\xc6\x68\x58\xcd\xdf\xc1\x48\xc6\xe1\x06\xf4\x3f\xf5\x4b\x64\x61\xb4\x32\x26\xf4\x62\x54\xa5\x85\xc4\x71\x1f\x49\x7b\x48\x02\xe7\xc9\x1e\x42\x9c\xcd\xcd\x05\x68\x5e\x8e\x60\x8d\x5f\x49\xb0\x56\x1a\xd2\xbf\x45\x3c\x92\x02\x94\xa6\x7c\xd3\x15\x8a\x28\x83\x01\x71\xa3\xac\x5f\xc7\x5a\xb5\x38\xb4\x32\x2f\xb2\xa9\xca\xf2\xc5\x2d\x6a\x32\x50\x5e\x87\xb7\x3a\xa1\xaa\x47\xd9\xb7\xe4\xd9\xc0\xd4\x19\xaa\xa2\x26\xe0\xd7\x75\x66\x5d\xaa\x1e\xdb\x0e\x0f\xdb\x49\x57\x4d\xad\xd6\x19\xce\x75\x72\x13\x26\x90\xe2\xed\x3c\x3c\xaa\x08\x5e\x14\x1f\x68\xe4\x7d\x64\x58\xe9\xb2\x2c\x0a\x60\x0a\xca\x72\xaf\x28\xe6\x34\xaa\x99\x5d\x14\x73\xc3\x76\xfb\x6b\x23\x87\x9a\x85\x37\xa0\xbb\x53\x60\x45\x89\x43\xbb\x56\x15\x29\xa4\x01\xc8\xee\x12\x30\x2c\x26\xaf\xb7\x18\xd4\x91\x8b\xe2\x23\xe5\x04\x9e\xca\x2e\xc3\xcc\x33\x3a\xe0\xb9\x39\xb6\x4d\xee\x21\x2c\x03\xaa\x41\x5e\x71\xb6\x46\x4b\x4f\xcb\x1c\xf6\x50\x26\xa9\x90\x54\x1b\x83\xef\x14\x51\x43\xd2\x60\xd2\x86\xa3\xfa\xbb\x73\x47\x85\x14\xee\x67\x1d\xa9\x04\x4b\x42\x79\x5c\xdb\x26\xcd\x93\xb8\xab\x9b\x19\x7d\xdf\x14\xd8\x95\x66\x7c\x4b\x33\xd6\x2f\x8f\xa4\xea\xd4\xcb\xb0\x52\x8f\x42\x12\xcf\x25\x0b\x81\x8c\x89\x75\x0a\x5c\x9f\xe5\x3a\xb9\x84\x30\xc1\x9c\xaa\xb4\xbf\x7f\xa1\x9b\xf3\xeb\xb3\xcb\xe9\xcd\xe7\xb3\xe9\xfe\x16\xb0\xf6\x97\xf0\x1e\xd6\x75\xae\x7b\x58\x77\xd2\xdc\xc3\xda\x2a\xd5\x51\xe4\xbf\x22\x50\x8e\x1c\x43\xc1\x23\x5a\x2b\x74\x25\xd2\x99\x89\xd8\xae\x95\x2a\x1c\x30\x20\xfd\xad\x02\xe5\x0a\xa4\xba\x00\x06\xda\x1e\xb7\x37\x83\xea\xec\x16\xf3\xfa\xc8\xe9\x6a\x70\x86\x99\x2a\x31\x49\x29\x1f\x1a\x2c\xcd\x9b\x0d\x2e\x14\xc3\x1b\xea\x0e\x71\x9d\xf8\xc6\x31\x64\xb9\xd2\x20\x2f\xab\x4f\x5b\x7f\x23\xf1\x7a\x5b\x89\xe7\xbe\x47\x13\x31\x6f\x31\xa5\xfb\xa5\x1c\xea\x90\x72\xaa\x87\x79\xf3\xd2\xdc\xfe\xe3\x91\x3c\xb3\xfe\xc3\x38\x0e\x47\xf9\x7d\x25\x48\xc0\xe4\x8c\xaf\x2f\xb0\xc6\x81\xf9\x77\xeb\x7d\xaa\x30\x44\xb0\x28\xbc\x59\x29\x4c\x84\x98\xed\x0a\xc8\xad\xa4\xfa\xbd\xa0\x78\x2d\xff\x2f\x0b\xab\xbd\x7d\xbf\xa1\xa8\x9c\x3e\xbd\x5e\x9e\xa6\xa4\x51\x87\x17\x44\x37\x82\xee\x2e\xd0\xee\xaa\xb5\xa1\x77\xfd\x1d\x69\x5f\xd7\xd7\x5b\x65\xde\x51\x53\xbf\x9e\xfe\x45\x29\xed\xc4\xe3\xb7\x04\x7b\x34\x67\xd7\xd4\xd7\x4d\x47\x0a\x36\xe7\x6a\x5c\x10\x83\xab\x86\x59\x6d\x9c\x7d\xc5\x04\xeb\x2d\x28\xe5\xa4\x9c\xfc\x17\x00\x00\xff\xff\xf4\x31\xf6\xbb\x02\x13\x00\x00")

func replica_setJsonBytes() ([]byte, error) {
	return bindataRead(
		_replica_setJson,
		"replica_set.json",
	)
}

func replica_setJson() (*asset, error) {
	bytes, err := replica_setJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "replica_set.json", size: 4866, mode: os.FileMode(420), modTime: time.Unix(1497272935, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _sharded_setJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xd4\x58\x4b\x6f\xdb\x38\x17\xdd\x17\xe8\x7f\x10\xd8\xcd\xf7\x01\xf1\x33\x69\xa6\xf1\xaa\x6d\x02\xcc\x0c\x30\x69\x8a\xa4\x8f\x45\x1b\x04\x94\x78\x25\x71\x42\x91\x02\x49\x3b\x71\x05\xfd\xf7\x01\x29\xd9\xa2\x1e\xb4\xd3\x76\x26\x48\xbd\x33\x75\xee\x43\xf7\x9e\x7b\x48\xb1\x78\xfe\x2c\x08\x02\x24\x72\x4d\x05\x57\x68\x11\x54\x0b\xf6\x87\x88\xb8\xe3\x4c\x60\xf2\x16\x2b\x40\x8b\x00\x4d\x56\x58\x4e\x18\x0d\x27\x99\xe0\x89\x20\xe1\x28\xcb\xd4\x08\x2f\xb5\xc8\xb0\x31\x47\x07\x1e\xdb\xcf\x94\x13\x71\x67\xbc\xa3\xd3\xc5\xd7\xaf\x1e\xeb\xca\xb8\xac\x9d\x20\x8b\x3a\x0b\x3f\x81\x54\x75\x6a\x5f\xdc\xdc\xc2\x25\x65\xc4\xae\x36\x8b\x9b\x5f\xd1\x5f\xaa\x6c\xa8\x36\x16\xc7\x47\x07\x1e\x40\xcc\xf0\x4a\x48\x93\x27\xf2\x41\x12\xaa\xeb\x94\x0c\xec\x68\x7e\x74\x12\xcd\xc8\x3c\x7c\x79\x72\x72\x02\x61\x18\xcf\x62\x12\x1f\x86\xd1\x14\xa6\x70\x18\x1e\xc6\xf1\xcb\x68\x8a\x71\x3c\xf7\xba\xcb\xf0\xfd\x85\x72\x1c\xfa\x81\x94\x3f\x10\x28\xc8\x92\x81\xad\xcc\xb5\x0f\x93\x33\xac\x63\x21\x33\xe3\x48\xa8\x7b\xaf\xaf\xa5\x64\x06\x92\x6a\x9d\xab\xc5\x64\x12\x63\xa5\x09\x1b\xd7\xfd\x1b\x0b\x99\x4c\x18\xe5\xcb\xfb\x2d\x1f\xec\xbf\xd1\xfd\xab\xe3\x9b\xe3\xa3\x11\xce\xf0\x37\xc1\x47\x87\xe3\xf9\xf8\xb7\xb1\x4e\xbe\x79\xa3\xdc\x51\x3e\x9f\x4e\x5f\xe5\x6c\x69\xb2\x8e\x31\x53\xb0\x03\xfa\xe9\xf4\x12\x08\x55\xfa\x8c\xb1\x9d\x75\x70\xa0\x17\x5b\x76\xfb\x4b\xe2\xc0\x3f\xca\x07\x7b\x76\x1b\xd2\x47\x97\xed\x25\x37\x36\xe2\x38\xb3\x33\x65\xcb\xb3\xe1\xfe\x06\x81\x42\x1c\xdd\x2e\x73\x97\xfb\xd5\x03\x03\x68\xe6\x83\x6a\x21\x29\x4f\xfa\x30\x77\x50\x52\xa1\xf4\x26\x58\x51\x70\x41\x40\x8d\xbf\x4c\xaf\xcb\xb2\x35\xae\x4c\x24\xef\xb1\x4e\x9b\x29\x17\x89\x67\xca\x27\x4d\xe0\x11\x4e\x80\xeb\x31\x13\x49\xd7\xd7\xa5\xd0\x58\x43\x5b\x4e\xec\x33\x45\xbf\xc1\x87\x54\x82\x4a\x05\x23\xe7\x6f\xd1\x22\x98\x4d\xa7\xd3\x4e\xa5\x91\xa6\x59\x83\xfa\x43\x9a\x37\x9b\x1f\x35\x98\xba\xae\x65\x53\x13\x6b\x95\x4b\x11\x81\x52\x15\xf7\x8b\xe2\x05\xe0\x28\x0d\xec\x0b\x97\xdb\x4e\xb8\x95\xc1\x32\x51\xf3\x9b\xe3\x81\x2c\x39\xe8\xfe\x6a\x15\x43\x48\xf3\x68\xfe\x6a\x3a\x9d\xb6\x1f\x97\x9d\x97\x28\x8a\x17\x34\x0e\xfe\x47\xd5\x9f\xfc\x2a\xc5\x92\x04\xaf\x29\x27\x70\xff\xff\xb2\xc3\x0a\x24\x21\x67\x34\xaa\x34\x70\x38\xa8\x41\x5c\x81\x7e\x57\x77\x51\x19\x77\x37\x45\x41\xe8\xaa\x76\x1a\x1c\x06\x65\x89\xf6\xe5\x33\xa1\x71\x37\xf8\x36\xc9\x53\xc1\x63\x9a\x78\x73\xb4\x21\x29\x4f\x3c\x09\x46\x6c\xa9\x34\xc8\x4b\xc1\x6c\x82\x91\x75\xa6\x56\xf2\x27\x52\x4a\xb1\xba\xd2\x42\xe2\x04\xfc\x49\x55\xcf\x3d\x39\x91\xb0\x45\xe8\x55\x84\xf3\x89\xb1\x80\x2d\xaf\x09\xd6\xf8\x47\x32\x44\x6a\xad\x34\x64\x7f\x09\x5f\x39\x08\x28\x4d\xf9\xa6\xa3\x28\xa6\x0c\x86\xc4\x04\xe5\xfd\x04\xd7\xca\x19\xbd\x1b\x43\xde\x4d\xba\x76\xcc\x3a\xd9\x3a\x23\x71\xe0\x1d\x79\x9d\x52\xd5\x9f\xf6\xff\x78\x42\x87\xa4\x6e\x30\x95\x7a\x68\x3f\xac\xf3\x1a\xe3\x2c\xd4\x9d\x6f\x5b\x38\xfc\x30\xa4\x04\x72\x5a\x91\x6f\x88\x26\x1b\x62\x6e\xa7\x06\xc8\x68\xb3\xd4\x72\xd9\xe9\x32\x5a\x35\xba\x5e\x09\xb4\x9b\x32\x5e\xea\xf4\x2a\x4a\x21\xc3\x8d\xfc\xbf\xac\x25\xa9\xca\xee\x35\xc3\x4a\x97\x65\x51\x00\x53\x50\x96\x07\xad\x00\x45\x31\x31\xca\x54\x76\xe5\xab\xd6\x81\x2b\xd0\x2d\x01\xcb\xb1\xd4\xd4\x50\x09\xc8\xbb\x8d\x96\x75\x48\x3a\x40\xad\x1b\x4a\x5c\xa9\x18\x2a\xe4\x16\x9b\x41\x16\x82\x6c\x05\xad\x1a\x35\xbc\xff\x79\x4e\x57\x4e\xd8\x26\x9e\x67\x0b\xad\xca\x28\x43\xaa\x41\x5e\x70\xb6\xde\xb3\xed\x5b\x78\x4a\x09\x01\xfe\x20\xa4\x50\xda\xc7\xb7\x1e\x38\x97\x54\x48\xaa\x4d\x0a\xb3\x5d\x38\xc5\xf0\x0a\xce\x80\x61\x83\xec\x8e\x43\x0b\xb9\x12\xda\x6e\x41\xb3\x61\xcc\x43\x28\xd2\xfd\x75\x28\xe3\xfe\xae\xf7\x0a\xd8\x90\xe5\x5e\xca\x1c\xee\x63\xca\x8f\x72\x63\x67\xe9\x1e\x81\x12\xf5\x01\xe8\xe4\xfa\xc9\xf1\xc2\x63\xbb\xbf\xa4\x3b\xf3\x7b\xbc\x92\xce\xba\x87\xca\x9e\xc9\x2f\x54\xd3\xf9\x13\xa9\xe9\xec\xe9\xd5\x74\xbf\x06\x75\x36\x36\xe7\xf0\xe8\xc8\xc6\x90\x04\x59\xa4\x57\x5e\x76\xed\x88\x43\x99\xee\xea\xb2\xc6\xc9\xce\xef\xc1\xe0\x3b\x37\xd1\xad\x8d\x55\xc7\xbe\xc9\xe3\xec\x05\x43\x9a\xcd\xdd\xaf\x86\xe1\xf3\xcf\x16\x5b\x1d\xdc\xaf\x40\xae\xec\x99\x69\xb0\x3a\x2d\xd0\x65\x75\x66\xd9\xb7\x6f\x44\x82\x31\x88\x9a\x8f\x70\x2f\x5f\x6a\xc6\x98\x13\x16\x5a\x34\xfd\x43\x84\x2a\x1c\x32\x20\x68\xd1\x99\x21\x03\x15\xef\xef\x48\x35\x35\x41\x8e\x95\xba\x13\x92\x04\xad\x2e\x59\xd0\x47\x65\xde\x09\x79\xaf\xac\x10\x81\x9c\x89\x75\x06\x5c\xbf\x59\xea\xf4\x1c\xa2\x14\x73\xaa\xb2\x3e\x1d\xd1\xf9\xc5\xbb\xdf\x2f\xce\xde\x8e\x4e\x2f\x9b\xbe\x3a\xa5\x42\xb7\xb0\x46\x0b\x93\xce\x2d\xac\x3b\x99\xdc\xc2\xda\x7e\x0e\x2c\x9c\x43\xff\xdf\x22\x54\xad\x13\x7f\x55\xe1\x6a\xe9\x46\x64\x63\xe3\xcf\xf1\xb1\x54\x20\xd5\x67\xcc\xb5\xa9\x47\x2b\xb5\x2e\xe3\x11\x09\xd1\x02\x61\x92\x51\xde\xeb\x0b\xa2\x9c\x6a\x5b\x39\x7f\xe1\x6a\xa4\x14\xf6\x42\xe9\xc1\x37\x6d\xbb\xc2\x3a\x2e\xd1\x62\x73\x3e\x3f\xaf\xae\x14\xf6\x5f\xa1\xf4\x29\x6e\xab\x51\xf7\xb5\x7b\x33\xd1\xf2\xd7\xde\x0b\x7e\xd5\x4a\xbd\xb1\xd0\x81\x3a\x0d\x58\xff\x6c\x48\x09\x98\xbc\xe1\xeb\x33\xac\x71\x88\x15\x3c\x52\x54\xd3\x4f\xfb\x96\xff\x6e\x68\x26\x22\xcc\x1e\xf0\xc2\x9f\x25\xd5\x8f\xf5\xaa\x3b\xe3\x7d\x0f\xf1\xab\xeb\xc2\x21\xd2\x0f\xea\x93\x15\x90\x33\x60\x50\x29\xc8\x75\x47\x29\x5b\xf2\x67\x84\xb5\xa7\x77\xe5\xf3\x67\xe5\xf3\x67\xff\x04\x00\x00\xff\xff\xe2\xc1\x12\xf9\x3c\x18\x00\x00")

func sharded_setJsonBytes() ([]byte, error) {
	return bindataRead(
		_sharded_setJson,
		"sharded_set.json",
	)
}

func sharded_setJson() (*asset, error) {
	bytes, err := sharded_setJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "sharded_set.json", size: 6204, mode: os.FileMode(420), modTime: time.Unix(1497272935, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _standaloneJson = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\xcc\x56\x4d\x6f\xdb\x38\x10\xbd\xfb\x57\x08\x3c\xd7\xb6\xe2\x38\xd9\xc4\x37\x37\x39\xf4\xb0\x41\x8b\xa4\x6d\x0e\x6d\x61\x8c\xc4\x91\xc4\x0d\x45\x0a\x24\x15\xc7\x09\xfc\xdf\x17\xa4\x14\x5b\xb2\x44\xd7\x05\xe2\xcd\xfa\xe6\x99\xa7\xf9\xe2\x7b\x43\xbe\x0c\x82\x20\x08\x88\x2c\x0c\x93\x42\x93\x59\x50\x19\x9c\x91\xca\xa5\xe0\x12\xe8\x47\xd0\x48\x66\x01\x19\x3f\x82\x1a\x73\x16\x8d\x73\x29\x52\x49\xa3\x61\x9e\xeb\x21\x94\x46\xe6\x60\xbf\x26\x1f\xfa\x3f\xbd\x67\x82\xca\xa5\x8d\x4d\xae\x66\x3f\x7f\x7a\x3e\x76\xdf\xae\xab\x10\xc4\x61\xae\xa3\xef\xa8\x74\x5d\xd6\x8f\x46\x5d\x51\xc9\x38\x75\xc6\x8d\xed\xf5\xf7\xd2\xb1\x54\x5f\x30\x63\xf1\xe7\xd3\x0f\xfd\xfe\x84\xc3\xa3\x54\xb6\x42\xe2\x41\xa4\xcc\xd4\xd5\x58\xd4\x74\x32\xbd\x8c\x4f\xe8\x24\x3a\xbb\xbc\xbc\xc4\x28\x4a\x4e\x12\x9a\x9c\x46\x71\x88\x21\x9e\x46\xa7\x49\x72\x16\x87\x00\xc9\xc4\x17\x2d\x87\xa7\xcf\xba\x11\xcf\x8b\x63\xe2\x30\x9c\xa4\x25\x47\x37\x92\x5f\x1e\x48\xc1\xc1\x24\x52\xe5\x36\x8c\xd4\x4f\xbe\x48\xa5\xe2\x16\x91\x19\x53\xe8\xd9\x78\x9c\x80\x36\x94\x8f\xea\x33\x1b\x49\x95\x8e\x39\x13\xe5\xd3\x86\x02\xee\xdf\xf0\xe9\xe2\x7c\x71\x3e\x1d\x42\x0e\xcf\x52\x0c\x4f\x47\x93\xd1\x5f\x23\x93\x3e\xfb\x92\x2c\x99\x98\x84\xe1\x45\xc1\x4b\x5b\x72\x02\x5c\xa3\x1f\xf9\xfd\xea\x16\x29\xd3\xe6\x9a\xf3\x7d\x23\x68\x20\x3f\x6f\xc8\xec\x9d\x46\x03\xfd\x4d\x1d\x1a\xb7\x79\x10\x1d\xf0\xba\x65\x69\xe4\x25\x02\x72\x27\x1f\x37\x96\x9a\xe7\xb5\x9f\x44\x10\x3f\x94\x85\x87\xe7\x99\xd4\xe6\xf5\xe3\x97\x17\x21\x29\xea\xd1\x8f\xf0\xd7\x7a\xdd\x54\x1a\x97\xe9\x17\x30\xd9\x56\x9f\x32\xf5\xe8\x73\x5c\x25\x1b\x42\x8a\xc2\x8c\xb8\x4c\x77\xc2\xdc\x4a\x03\x06\x5b\x2b\xc0\xb9\x34\x7b\xc6\xaf\x99\x42\x9d\x49\x4e\x6f\x3e\x92\x59\x70\x12\x86\x61\x7b\x5c\xc4\xb0\x7c\x0b\xfa\xa4\x6c\x2b\x93\xe9\xa0\x3d\x1c\xdb\xf5\x46\xdf\xcc\x48\xc5\x44\xfa\x1f\xb5\xbe\x4d\xf8\xbe\xed\x3b\x70\xa1\x64\x8c\x5a\xe3\x6e\xd3\xa0\x52\x3d\x59\x9c\x77\x4b\x10\x68\x3a\xc6\x2a\x92\x54\xd6\x33\xb9\x08\xc3\xb0\xe5\x5d\xef\xd4\xa7\x8d\x54\x90\x76\xbb\x73\x4e\x1a\xb5\xc6\xf8\x18\x43\x31\xb6\x1f\xe0\x66\x9a\x14\x0c\x90\xfd\x09\x56\xda\x60\xfe\xb7\x4c\x3d\x29\x50\x1b\x26\xaa\x4d\x3f\x0b\x48\xc2\x38\xf6\x08\x8e\x14\xdd\x3a\x56\xba\x71\xae\x0b\xcb\x84\xd7\xaa\xdc\x19\xb6\x8b\x1a\xf4\x94\x77\x30\x97\x8e\x49\x80\x9e\x75\xe0\x2d\xa5\xa6\xc7\xd7\x55\xe1\x70\x55\xb7\x4d\xc0\xe3\x76\x0f\x55\x3b\xa5\xe1\x83\xd2\x64\x77\x71\x86\x39\x6c\xb7\xd5\x59\x9b\x7e\x0a\x0b\xce\x62\xb8\x43\xd3\x5c\x91\x44\xc9\xf6\x0d\x42\x74\x06\x8a\x32\x91\xd6\xb6\xc1\x26\x7e\xfb\x95\x60\x55\xf6\x4d\xa3\xbb\x3d\xfd\x0f\x02\x6b\xfd\xb2\xa4\x55\xe3\x41\x01\x5a\x2f\xa5\xa2\x41\xbb\x73\x8a\x05\x97\xab\x1c\x85\x99\x97\x26\xbb\xc1\x38\x03\xc1\x74\xde\xbd\xea\xc9\xdd\xd5\xed\xfc\x66\x78\xf7\x69\x3e\x3c\xd9\x52\xa0\xb9\x74\x1f\x70\x55\xe7\x7a\xc0\xd5\x4e\x9a\x07\x5c\x39\x02\xb6\x88\xf6\x8f\x8c\x74\x8b\x65\xb1\x14\x09\xab\x89\xb7\x90\xf9\xc8\x46\x6c\xd6\xca\x34\x44\x1c\x69\xf7\x02\x23\xa5\x46\xa5\xaf\x91\xa3\x71\xee\xe6\x25\x54\xf9\xee\x41\xd4\xae\x56\x57\xbd\xd2\xb4\x55\x02\xcd\x99\xe8\xd3\xcb\xe6\xcc\x7a\xef\xae\xfe\xc7\xd0\x01\x71\x5b\xf1\x2d\x30\xe6\xa5\x36\xa8\x6e\xaa\x2d\xda\xbd\xfc\x82\xce\x05\x18\xb4\xcf\x63\x13\xb1\x6c\x30\x65\x77\x29\xf7\x75\xc8\x04\x33\xfd\xbc\xd9\xb7\x8e\xfe\xf7\x93\x9c\x3b\x7c\xff\x1c\xfb\xa3\xbc\x5d\x09\x0a\x81\xce\xc5\xea\x1a\x0c\x44\xf6\x65\xff\x3e\x55\x58\x22\xb8\x29\x1c\xad\x14\x2e\x63\xe0\x87\x0e\xe4\x5e\x31\xf3\x5e\xa3\xf8\x5d\xfe\x3f\x16\x56\xf3\xa1\x77\x44\x51\xb5\xfa\x0c\x3a\x79\x36\x25\x79\x01\x7b\x44\xe7\x99\xee\x21\xa3\x3d\x54\x6b\x7d\xe7\xfa\x16\x69\x7f\xaf\xaf\x63\x65\x3e\x50\x53\x7f\x9e\x7e\xaf\x94\x0e\xe2\xf1\x31\x87\xed\xcd\xb9\x6b\xea\xea\x66\x47\x0a\x2e\xe7\xc2\x2f\x88\xde\xa7\x86\x7d\xda\xb4\xde\x2b\x36\x58\xe7\x81\xb2\x1e\xac\x07\xff\x06\x00\x00\xff\xff\x44\xd4\x65\x75\x6d\x11\x00\x00")

func standaloneJsonBytes() ([]byte, error) {
	return bindataRead(
		_standaloneJson,
		"standalone.json",
	)
}

func standaloneJson() (*asset, error) {
	bytes, err := standaloneJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "standalone.json", size: 4461, mode: os.FileMode(420), modTime: time.Unix(1497272935, 0)}
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
	"replica_set.json": replica_setJson,
	"sharded_set.json": sharded_setJson,
	"standalone.json": standaloneJson,
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
	"replica_set.json": &bintree{replica_setJson, map[string]*bintree{}},
	"sharded_set.json": &bintree{sharded_setJson, map[string]*bintree{}},
	"standalone.json": &bintree{standaloneJson, map[string]*bintree{}},
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

