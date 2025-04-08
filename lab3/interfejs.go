package main
import (
	"time"
	"errors"
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
	ModifiedAt time.Time
	size       int64
}

func NewFile(name string, parentPath string) *File {
	return &File{
		name:       name,
		parentPath: parentPath,
		created:    time.Now(),
		ModifiedAt: time.Now(),
		size:       0,
	}
}

type DirectoryImpl struct {
	name       string
	parentPath string
	created	time.Time
	ModifiedAt time.Time
	items      []FileSystemItem
}

func NewDirectory(name string, parentPath string) *DirectoryImpl {
	return &DirectoryImpl{
		name:       name,
		parentPath: parentPath,
		created:    time.Now(),
		ModifiedAt: time.Now(),
		items:      []FileSystemItem{},
	}
}

type SymLink struct {
	name       string
	parentPath string
	targetPath string
	created    time.Time
	ModifiedAt time.Time
}

func NewSymLink(name string, parentPath string, targetPath string) *SymLink {
	return &SymLink{
		name:       name,
		parentPath: parentPath,
		targetPath: targetPath,
		created:    time.Now(),
		ModifiedAt: time.Now(),
	}
}

type ReadonlyFile struct {
	name       string
	parentPath string
	created    time.Time
	ModifiedAt time.Time
	size       int64
}

func NewReadonlyFile(name string, parentPath string) *ReadonlyFile {
	return &ReadonlyFile{
		name:       name,
		parentPath: parentPath,
		created:    time.Now(),
		ModifiedAt: time.Now(),
		size:       0,
	}
}



