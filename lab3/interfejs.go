package main

import (
	"errors"
	"time"
)

// Iterfejs definiujący obiekt w systemie plików
type FileSystemItem interface {
	Name() string
	Path() string
	Size() int64
	CreatedAt() time.Time
	ModifiedAt() time.Time
}

// Interfejs definiujący obiekty które mogą być odczttywane
type Readable interface {
	Read(p []byte) (n int, err error)
}

// Interfejs definiujący obiekty w których można dokonywać zapisu
type Writable interface {
	Write(p []byte) (n int, err error)
}

// Katalog definiuje pliki i podkatalogi
type Directory interface {
	FileSystemItem
	AddItem(item FileSystemItem) error
	RemoveItem(name string) error
	Items() []FileSystemItem
}

// Przykładowe komunikaty błędów, które można użyć
var (
	ErrItemExists       = errors.New("item already exists")
	ErrItemNotFound     = errors.New("item not found")
	ErrNotImplemented   = errors.New("operation not implemented")
	ErrPermissionDenied = errors.New("permission denied")
	ErrNotDirectory     = errors.New("not a directory")
	ErrIsDirectory      = errors.New("is a directory")
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

type DirectoryImpl struct {
	name       string
	parentPath string
	created    time.Time
	modified time.Time
	items      []FileSystemItem
}

func NewDirectory(name string, parentPath string) *DirectoryImpl {
	return &DirectoryImpl{
		name:       name,
		parentPath: parentPath,
		created:    time.Now(),
		modified: time.Now(),
		items:      []FileSystemItem{},
	}
}

func (f *DirectoryImpl) Name() string {
	return f.name
}

func (f *DirectoryImpl) Path() string {
	return f.parentPath + "/" + f.name
}

func (f *DirectoryImpl) Size() int64 {
	return int64(len(f.items))
}

func (f *DirectoryImpl) CreatedAt() time.Time {
	return f.created
}

func (f *DirectoryImpl) ModifiedAt() time.Time {
	return f.modified
}

func (f *DirectoryImpl) AddItem(item FileSystemItem) error {
	for _, existingItem := range f.items {
		if existingItem.Name() == item.Name() {
			return ErrItemExists
		}
	}
	f.items = append(f.items, item)
	f.modified = time.Now()
	return nil
}

func (f *DirectoryImpl) RemoveItem(name string) error {
	for i, item := range f.items {
		if item.Name() == name {
			f.items = append(f.items[:i], f.items[i+1:]...)
			f.modified = time.Now()
			return nil
		}
	}
	return ErrItemNotFound
}

func (f *DirectoryImpl) Items() []FileSystemItem {
	return f.items
}

type SymLink struct {
	name       string
	parentPath string
	target FileSystemItem
	created    time.Time
	modified time.Time
}

func NewSymLink(name string, parentPath string, target FileSystemItem) *SymLink {
	return &SymLink{
		name:       name,
		parentPath: parentPath,
		target: target,
		created:    time.Now(),
		modified: time.Now(),
	}
}

func (s *SymLink) Name() string {
	return s.name
}

func (s *SymLink) Path() string {
	return s.parentPath + "/" + s.name
}

func (s *SymLink) Size() int64 {
	return s.target.Size()
}

func (s *SymLink) CreatedAt() time.Time {
	return s.created
}

func (s *SymLink) ModifiedAt() time.Time {
	return s.modified
}

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