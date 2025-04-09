package vfs

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