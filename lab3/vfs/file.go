package vfs

import (
	"errors"
	"time"
)

type File struct {
	name       string
	parentPath string
	created    time.Time
	modified time.Time
	data       []byte
}

func NewFile(name string, parentPath string) *File {
	return &File{
		name:       name,
		parentPath: parentPath,
		created:    time.Now(),
		modified: time.Now(),
		data:       []byte{},
	}
}

func (f *File) Name() string {
	return f.name
}

func (f *File) Path() string {
	return f.parentPath + "/" + f.name
}

func (f *File) Size() int64 {
	return int64(len(f.data))
}

func (f *File) CreatedAt() time.Time {
	return f.created
}

func (f *File) ModifiedAt() time.Time {
	return f.modified
}

func (f *File) Read(p []byte) (n int, err error) {
	if len(f.data) == 0 {
		return 0, errors.New("no data to read")
	}
	n = copy(p, f.data)
	return n, nil
}

func (f *File) Write(p []byte) (n int, err error) {
	f.data = append(f.data, p...)
	f.modified = time.Now()
	return len(p), nil
}
