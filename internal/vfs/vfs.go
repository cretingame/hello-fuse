package vfs

import (
	"bytes"
	"io/fs"
	"time"
)

// to be compatible with FileInfo inteface from "io/fs" module
type virtualFileInfo struct {
	name string
	data []byte
}

func (vfi *virtualFileInfo) Name() string       { return vfi.name }
func (vfi *virtualFileInfo) Size() int64        { return int64(len(vfi.data)) }
func (vfi *virtualFileInfo) Mode() fs.FileMode  { return 0444 }
func (vfi *virtualFileInfo) ModTime() time.Time { return time.Unix(0, 0) }
func (vfi *virtualFileInfo) IsDir() bool        { return false }
func (vfi *virtualFileInfo) Sys() any           { return nil }

// to be compatible with File inteface from "io/fs" module
type virtualFile struct {
	// bytes.Reader provides the `Read([]byte) int` function
	*bytes.Reader
	info virtualFileInfo
}

func NewVirtualFile(name string, bs []byte) *virtualFile {
	return &virtualFile{
		Reader: bytes.NewReader(bs),
		info: virtualFileInfo{
			name: name,
			data: bs,
		},
	}
}

func (vf *virtualFile) Stat() (fs.FileInfo, error) {
	return &vf.info, nil
}

func (vf *virtualFile) Close() error {
	return nil
}

type virtualFS struct {
	files map[string]fs.File
}

func NewVirtualFS() *virtualFS {
	return &virtualFS{
		files: map[string]fs.File{},
	}
}

func (vfs *virtualFS) Open(name string) (fs.File, error) {
	_, ok := vfs.files[name]
	if !ok {
		// OPTIM: I should watch `fs.ValidPathname()` the fs module how the errors are handled
		// maybe with `fs.PathError` ?
		return nil, fs.ErrNotExist
	} else {

		return vfs.files[name], nil
	}
}

func (vfs *virtualFS) Add(file fs.File) error {
	name := ""
	info, err := file.Stat()
	if err != nil {
		return err
	}
	name = info.Name()
	vfs.files[name] = file
	return err
}
