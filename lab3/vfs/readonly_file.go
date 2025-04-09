package vfs

import (
	"errors"
	"time"
)

type ReadonlyFile struct {
	name       string
	parentPath string
	created    time.Time
	modified time.Time
	data       []byte
}

func NewReadonlyFile(name string, parentPath string) *ReadonlyFile {
	return &ReadonlyFile{
		name:       name,
		parentPath: parentPath,
		created:    time.Now(),
		modified: time.Now(),
		data:       []byte{},
	}
}

func (f *ReadonlyFile) Name() string {
	return f.name
}

func (f *ReadonlyFile) Path() string {
	return f.parentPath + "/" + f.name
}

func (f *ReadonlyFile) Size() int64 {
	return int64(len(f.data))
}	

func (f *ReadonlyFile) CreatedAt() time.Time {
	return f.created
}

func (f *ReadonlyFile) ModifiedAt() time.Time {
	return f.modified
}

func (f *ReadonlyFile) Read(p []byte) (n int, err error) {
	if len(f.data) == 0 {
		return 0, errors.New("no data to read")
	}
	n = copy(p, f.data)
	return n, nil
}

func (f *ReadonlyFile) Write(p []byte) (int, error) {
	return 0, ErrPermissionDenied
}