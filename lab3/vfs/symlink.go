package vfs

import (
	"time"
)

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