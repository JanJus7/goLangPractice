package vfs

import (
	"time"
)

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