// Code generated by go-bindata.
// sources:
// etc/build/Makefile
// etc/build/make/bootstrap.mk
// etc/build/make/rules.mk
// DO NOT EDIT!

package build

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

var _makefile = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x84\x90\x41\x4f\xdc\x40\x0c\x85\xef\xf9\x15\x4f\x62\x0f\xbb\x12\xc9\x4a\xbd\x83\x84\x44\xc5\x81\xb6\x17\xda\x7b\x9c\x8c\x43\xa6\x99\x8c\x53\xdb\xb3\x82\x7f\x5f\x25\x94\x82\xaa\x8a\x3d\xda\xef\xe9\xbd\xcf\xbe\xc0\xf7\x31\x1a\xa2\xc1\x47\x06\x67\xd7\xe7\x45\x62\x76\xc8\xb0\x6d\xba\x2e\x0d\x36\x62\xa6\x89\xf1\x70\x7b\x7f\xb9\x2d\x1f\x6e\xef\x61\xa3\x94\x14\xd0\x31\xcc\x45\x39\x80\xbc\xba\x40\x6b\x61\x4a\xd2\x93\x47\xc9\xed\x25\x1e\x39\xb3\x92\x73\x00\x3f\x71\x5f\x3c\xe6\x47\xb4\x2f\x91\xb5\x85\x09\x8b\xf2\x42\xca\x75\x57\x62\x0a\x6d\x83\xaf\x94\x9f\x11\xe2\x30\xb0\x72\x5e\xf3\x4e\xa4\x06\x32\xb4\x5f\x6e\xbe\xdd\xfd\xb8\xb9\xfb\xdc\x82\x94\xa1\xfc\xab\x44\xe5\xb0\xe2\x44\x5b\x5d\x91\xba\xc4\xb6\x89\xfc\xe4\x4a\xfd\xd6\x35\xa8\xcc\x2b\xf1\x4a\x36\x53\x8e\x03\x9b\x37\x2e\x73\x6a\x41\xbe\x9d\xa2\x22\x0e\xca\x61\x1b\x16\x95\x9f\xdc\x3b\x5c\x29\x5b\xa2\x2d\x22\x3a\x5c\x40\x2f\x1f\x88\xb9\x4f\x25\xac\x71\xe4\x6f\x89\x2d\x8a\xfd\x7b\xd9\x9b\xe6\x22\xa9\xa9\xaa\x77\x8f\xc1\x15\x1a\x0b\x53\xf5\xea\xc1\x15\x76\xfb\x77\xfa\xe1\xb8\x96\x1d\xff\x02\xcf\x53\xb5\xdb\xdb\xc8\x29\xe1\x3f\x0d\xa8\xeb\x45\x63\xf6\x9a\xf3\x09\x9f\xae\x8f\x81\x4f\xc7\x5c\x52\xc2\x35\x76\xfb\x57\xd3\xe1\x50\xfd\x61\x3f\x5b\xf5\x91\xaf\x13\x71\x73\xa5\xe5\x9c\x51\x4b\x62\x6b\xe6\xe9\x77\x00\x00\x00\xff\xff\xde\x41\x22\x7f\x5f\x02\x00\x00")

func makefileBytes() ([]byte, error) {
	return bindataRead(
		_makefile,
		"Makefile",
	)
}

func makefile() (*asset, error) {
	bytes, err := makefileBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "Makefile", size: 607, mode: os.FileMode(420), modTime: time.Unix(1487887719, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _makeBootstrapMk = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x6c\x92\x41\x6f\xd3\x40\x10\x85\xcf\xdd\x5f\xf1\xd4\xe6\xe0\x95\x6c\x10\x27\xa4\x4a\x15\x72\x1b\x13\x2c\xda\x04\xb9\xa9\xc5\xcd\xb5\xb3\xe3\x64\x84\xb3\x1b\x76\xd7\xae\x90\xf8\xf1\x68\x93\x38\x29\x85\xeb\xce\xec\xf7\xde\x9b\x99\x2b\x0c\xb5\xe5\xba\xe9\xc8\xa1\x21\xd6\x6b\xb0\x5e\x75\xbd\x22\x85\xd6\x9a\x2d\xfc\x86\xf0\xbc\xad\x35\xb7\xe4\xfc\xbb\xed\x8f\x67\x71\x9f\xce\x67\x4f\xe9\x2c\xc3\xa7\x1b\x51\x3c\xcd\x97\xf9\x43\x56\xcd\xd3\x65\x5e\x66\x55\x99\x15\x8f\xf9\x62\x1e\x2a\xe2\x0a\x66\xe7\xd9\xe8\xba\x3b\x2b\x88\x69\x91\x97\x59\x51\x4d\xb3\xb2\xfa\x56\x64\x9f\xf3\xef\xb8\xbe\x81\xa2\x61\x2c\x9c\x01\x98\x44\xff\x34\xcb\x64\x12\xb9\x0d\x75\x1d\xd6\xec\x61\x69\x48\x76\xb5\x75\x84\x2f\x59\x3a\xc5\x6f\xac\x7a\x8f\x64\xf5\x21\xf9\x28\x4f\xbe\x66\x8b\x13\x32\x10\xd7\x06\x03\x59\xc7\x46\x4b\x21\xa6\x8b\xbb\xaf\x59\x51\xe5\x0f\x87\x2c\x68\x9a\xae\x75\x9b\xf7\x93\x68\x0c\x28\x13\x65\x79\x20\x3b\x76\xde\x3e\xe5\xf7\xd3\x73\xff\x24\x7a\x4d\x90\x49\xd3\x73\xa7\x42\x6e\x45\x2d\x6b\x52\x68\x68\x53\x0f\x6c\x7a\x8b\xd6\x58\xec\xcb\x0e\xac\x1d\x2b\x82\xb7\xf5\xc0\x2e\x59\xb1\xe0\x56\xd3\x4f\x44\x93\xc8\x58\x5e\xb3\xc6\x5d\x2e\x63\xf4\xfa\x08\x91\x02\x00\xae\xc0\x2d\x5e\x08\xb5\xa5\x11\x70\x97\xc7\x21\x4b\x63\x1c\x81\x1d\x48\x87\x01\x2b\x34\xbf\x82\x7c\xdd\x77\x5e\x5c\x94\x59\x71\xbb\x78\xcc\xc2\x88\xbd\xed\x49\x90\x56\xdc\x06\x83\xdc\x62\x59\xa4\x65\xfe\x58\x2d\xd3\xd9\xc9\xee\x9b\x15\xb0\x83\x19\xc8\x5a\x56\xa4\xce\x26\xcf\xff\x64\x8c\x83\xb9\x37\xff\xae\xc3\x64\x5e\xb5\xfd\x25\x7b\xcc\xa0\x8d\x07\x6b\xf8\x7a\x1d\xef\x0f\x6c\xd7\xbb\x4d\x10\x54\xec\xf6\x31\x04\xb7\x07\xb9\x96\xad\xf3\x2f\xc6\x2a\x4c\x22\xd7\x37\xce\x23\x89\x11\x9f\x4e\xe3\xa8\x29\xa5\x8c\xff\x7b\x2f\x52\x5c\x04\xf4\x88\xc5\x0d\x2e\xf7\x52\xa7\x87\xb0\x18\x45\x03\x75\x66\xb7\x25\xed\xc7\xeb\x70\x97\x07\xd7\x7f\x02\x00\x00\xff\xff\xff\x6c\xae\x38\x1d\x03\x00\x00")

func makeBootstrapMkBytes() ([]byte, error) {
	return bindataRead(
		_makeBootstrapMk,
		"make/bootstrap.mk",
	)
}

func makeBootstrapMk() (*asset, error) {
	bytes, err := makeBootstrapMkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "make/bootstrap.mk", size: 797, mode: os.FileMode(420), modTime: time.Unix(1487887753, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _makeRulesMk = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x94\x94\x4d\x6f\xda\x4c\x10\xc7\xcf\xec\xa7\x18\x25\x3e\x78\xf5\x3c\x0e\x8d\x72\x03\xa1\x86\xc4\x6e\x6a\x35\x40\x64\x30\x97\xa4\x42\xc6\x3b\x0e\xab\xd8\x6b\xb2\xbb\x26\xe9\xb7\xaf\xd6\xf6\x62\x12\xa8\xaa\xde\x60\xe6\x37\x33\xff\x79\xf1\x9e\x03\x2b\xd3\x17\x94\x20\x2b\xa1\x79\x81\x90\x96\x45\x91\x08\xa6\x88\x3f\xbb\xfd\x11\x44\xab\xdb\x89\x0f\x5f\x47\x2d\x65\x8d\x37\x71\x78\x5f\x9b\x1d\xb7\xc3\x28\xac\x2b\x9e\x33\xcb\x44\xf1\xf4\x98\x90\x95\x00\xcf\x93\x85\x85\x1e\xe2\xf9\xf7\x63\x6a\x5b\xa9\x8d\x25\xbe\x85\xf7\x81\x21\xfc\x5a\x40\xc6\x73\x3c\xf4\x74\x4a\x3a\xff\x45\x23\xa3\xf6\xac\x96\xb3\xfb\x78\x12\xac\x16\xe3\xe8\x2e\x58\x18\xae\x5f\x6e\x75\x9f\x49\xbe\x43\xd9\x57\x32\xed\x7f\xe4\x1e\xc6\x8b\x56\x8f\xda\x60\x9e\xc3\xf6\x8d\x51\x42\xce\x9b\xce\x00\xc5\x8e\xcb\xb2\x40\xa1\x61\x97\x48\x9e\xac\x73\x54\x6d\x7c\x3c\x0f\x22\x13\xb8\x5e\xe7\x99\xda\x58\x63\xe8\x1f\x26\xe3\x0c\xbc\x0a\x1c\xd7\xb0\x94\xb6\xcc\x38\xba\x9b\x1b\xe8\x89\xf4\x3c\xaf\x2e\xe3\x25\xf2\x19\xba\xac\x23\xc7\xed\xfe\xd0\x3f\x70\xa1\xdf\x61\xa1\x7f\x44\x45\xf1\x74\x11\x4e\x82\xd5\x74\xbc\x08\x97\xc1\x6a\x19\x44\xf3\x70\x36\x1d\x39\xee\x69\x87\xd5\xd6\x5a\xdb\x13\xd8\xef\x28\x8a\xa7\x4d\x05\xd3\xcc\x81\xb6\xc1\x91\x84\xdd\x1e\x38\x98\x6f\xc7\x7d\x58\x4e\x1d\xb1\xaf\xd1\x00\xe1\x64\x7c\x17\xd4\x0b\xe0\x19\x2c\x83\xe8\x66\x36\x0f\x80\x2b\xa8\x84\x42\x6d\x2f\xb7\xd9\x0d\x57\x80\xef\x98\x56\x1a\x19\x70\x01\xaf\x15\xd7\x08\x45\xc9\x90\xf0\x0c\x5f\xc1\x75\xdc\x52\xf2\x67\x2e\x6c\x1a\xfa\x3f\x54\x82\x61\xc6\x05\x32\x4a\x7a\x07\xcb\xf8\x6f\x04\xde\x2b\x41\xc1\x78\x66\x2a\x6b\x94\x05\xa4\x65\x5e\x4a\x05\x1b\xcc\xb7\x28\x15\x49\x53\x89\x6c\x04\x4f\xf8\xf8\x65\x78\x75\x59\x90\x34\xfd\x85\x79\x5e\xbe\x59\xd3\x55\x51\x23\x0a\x75\x63\x29\x08\x49\xf2\x7c\xd0\x7e\x20\xe4\x64\x97\x83\xcf\xdd\x53\xf0\xb2\x6e\xe8\xdd\xb5\xd3\xfd\x50\x8d\x5c\x0a\x9e\x06\xe7\x1a\xe0\x82\x10\x8d\x4a\x0f\xe0\xf4\x0c\x7b\x36\xa8\x5b\x2a\x85\x22\x79\x41\x30\x51\x9e\x48\x34\xdf\x21\x21\xb5\xc2\x7f\xcf\xd1\x9c\x5a\x93\x64\x78\xbc\xc8\xe3\x56\x4e\x35\x61\xfd\xed\x3c\x1c\xd7\x8f\xc2\x65\x10\xed\xaf\xd2\xb4\x68\x9e\x06\x3b\xc7\xde\x35\xcf\xe0\x11\x1c\xd7\x18\x19\x57\xe6\x7b\x64\x14\x7e\x0e\x41\x6f\x50\x18\x15\x3d\x4c\x37\x25\x78\x08\x67\x8e\x5b\xef\x8c\x7e\x82\x1b\xb3\x42\x4d\xcf\x86\x4d\xc0\x3b\xd7\x70\x59\xff\xce\x38\x39\xe8\xc3\xbc\x54\xf4\xef\x1a\x09\x21\xbf\x03\x00\x00\xff\xff\xc8\x3b\xb0\xc6\x53\x05\x00\x00")

func makeRulesMkBytes() ([]byte, error) {
	return bindataRead(
		_makeRulesMk,
		"make/rules.mk",
	)
}

func makeRulesMk() (*asset, error) {
	bytes, err := makeRulesMkBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "make/rules.mk", size: 1363, mode: os.FileMode(420), modTime: time.Unix(1487887831, 0)}
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
	"Makefile": makefile,
	"make/bootstrap.mk": makeBootstrapMk,
	"make/rules.mk": makeRulesMk,
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
	"Makefile": &bintree{makefile, map[string]*bintree{}},
	"make": &bintree{nil, map[string]*bintree{
		"bootstrap.mk": &bintree{makeBootstrapMk, map[string]*bintree{}},
		"rules.mk": &bintree{makeRulesMk, map[string]*bintree{}},
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

