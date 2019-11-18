// Code generated by go-bindata. DO NOT EDIT.
// sources:
// metadata.tmpl (685B)

package main

import (
	"bytes"
	"compress/gzip"
	"crypto/sha256"
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
		return nil, fmt.Errorf("read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes  []byte
	info   os.FileInfo
	digest [sha256.Size]byte
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

var _metadataTmpl = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x7c\x91\xcd\x8e\x9b\x40\x10\x84\xef\xf3\x14\x15\xe4\x03\x48\x84\xb9\x27\xf2\x21\xb2\xa3\x28\x87\xc4\x87\x58\xb9\xac\xf6\xd0\x86\x06\x21\xe6\xc7\x1a\x86\x59\xed\xe2\x79\xf7\x15\xf8\x67\x2d\xad\x6d\x6e\xf4\x74\x7d\xaa\xaa\x96\x12\x2b\x5b\x31\x1a\x36\xec\xc8\x73\x85\xdd\x2b\x1a\x7b\xf9\x47\x6b\x3c\x3b\x43\x4a\x96\xba\x92\x9a\x3d\x55\xe4\xe9\x3b\xd6\x1b\xfc\xdd\x6c\xf1\x73\xfd\x7b\x5b\x88\x3d\x95\x1d\x35\x8c\xf3\xb3\x10\xad\xde\x5b\xe7\x91\x0a\x00\x48\x7c\xab\x39\x11\x99\x10\x52\xe2\x87\x52\xa0\x40\xad\xa2\x9d\xfa\x50\x14\xa2\xb4\xa6\x9f\x04\xe3\xf8\x15\x8e\x4c\xc3\x58\x74\x39\x16\x01\xdf\x96\x28\xd6\xe4\x09\x31\xce\xb4\x71\xc4\xa2\xc3\x01\x25\x69\x56\x2b\xea\x19\x31\x62\x89\xe4\x38\x8f\x31\x99\x11\x6c\xaa\x49\x90\x89\xc7\x40\x29\xf1\x8b\xfd\x4d\xe4\x4b\xab\x14\x1a\xf6\x38\x83\x11\x48\x0d\x8c\xda\x59\x7d\xe5\xbb\x1e\x4c\x89\x54\xe3\xcf\x69\x92\xdd\x03\xa6\x19\xd2\x69\x1e\x10\x63\x8e\x9d\xb5\x2a\x03\xc6\x39\x52\xc8\x61\xbb\xc9\x97\x7e\xba\xa5\x7c\x9e\x97\xda\x1a\x5f\x6c\x77\x52\x4c\x9f\x63\x3f\x38\x83\x23\xf3\x80\x37\x76\xf6\xff\xec\x70\xe2\xd7\xa4\x7a\x9e\x57\x8f\xad\x9d\x96\x43\x71\xf1\x90\xe5\xf0\x6e\x60\x11\xe7\xb3\xfc\x7b\xd4\x42\xff\xb9\x85\xd6\x78\xfb\xb0\x85\x3b\xc0\x34\xe0\x62\xe0\x1c\xff\x4e\x6a\x2c\x11\x44\xbc\xbe\xe6\x7b\x00\x00\x00\xff\xff\x1d\x4b\x64\x99\xad\x02\x00\x00")

func metadataTmplBytes() ([]byte, error) {
	return bindataRead(
		_metadataTmpl,
		"metadata.tmpl",
	)
}

func metadataTmpl() (*asset, error) {
	bytes, err := metadataTmplBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "metadata.tmpl", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info, digest: [32]uint8{0xb, 0x5f, 0xb7, 0x6e, 0x71, 0xeb, 0x77, 0x36, 0xa4, 0x3c, 0x69, 0xd0, 0xc, 0xe4, 0x17, 0x37, 0xfe, 0x32, 0x33, 0xe, 0x3c, 0x99, 0x3b, 0x96, 0xda, 0x1d, 0xf8, 0xc2, 0xfc, 0xfe, 0x6f, 0x3d}}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetString returns the asset contents as a string (instead of a []byte).
func AssetString(name string) (string, error) {
	data, err := Asset(name)
	return string(data), err
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

// MustAssetString is like AssetString but panics when Asset would return an
// error. It simplifies safe initialization of global variables.
func MustAssetString(name string) string {
	return string(MustAsset(name))
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetDigest returns the digest of the file with the given name. It returns an
// error if the asset could not be found or the digest could not be loaded.
func AssetDigest(name string) ([sha256.Size]byte, error) {
	canonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[canonicalName]; ok {
		a, err := f()
		if err != nil {
			return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s can't read by error: %v", name, err)
		}
		return a.digest, nil
	}
	return [sha256.Size]byte{}, fmt.Errorf("AssetDigest %s not found", name)
}

// Digests returns a map of all known files and their checksums.
func Digests() (map[string][sha256.Size]byte, error) {
	mp := make(map[string][sha256.Size]byte, len(_bindata))
	for name := range _bindata {
		a, err := _bindata[name]()
		if err != nil {
			return nil, err
		}
		mp[name] = a.digest
	}
	return mp, nil
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
	"metadata.tmpl": metadataTmpl,
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
// then AssetDir("data") would return []string{"foo.txt", "img"},
// AssetDir("data/img") would return []string{"a.png", "b.png"},
// AssetDir("foo.txt") and AssetDir("notexist") would return an error, and
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		canonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(canonicalName, "/")
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
	"metadata.tmpl": &bintree{metadataTmpl, map[string]*bintree{}},
}}

// RestoreAsset restores an asset under the given directory.
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
	return os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
}

// RestoreAssets restores an asset under the given directory recursively.
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
	canonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(canonicalName, "/")...)...)
}