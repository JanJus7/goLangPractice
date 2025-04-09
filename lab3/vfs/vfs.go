package vfs

type VirtualFileSystem struct {
	root Directory
}

func NewVirtualFileSystem(root string) *VirtualFileSystem {
	return &VirtualFileSystem{
		root: NewDirectory(root, ""),
	}
}

func (vfs *VirtualFileSystem) Root() Directory {
	return vfs.root
}

func (vfs *VirtualFileSystem) CreateFile(name string, parentPath string) (*File, error) {
	file := NewFile(name, parentPath)
	err := vfs.root.AddItem(file)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (vfs *VirtualFileSystem) CreateDirectory(name string, parentPath string) (*DirectoryImpl, error) {
	dir := NewDirectory(name, parentPath)
	err := vfs.root.AddItem(dir)
	if err != nil {
		return nil, err
	}
	return dir, nil
}

func (vfs *VirtualFileSystem) CreateSymLink(name string, parentPath string, target FileSystemItem) (*SymLink, error) {
	symLink := NewSymLink(name, parentPath, target)
	err := vfs.root.AddItem(symLink)
	if err != nil {
		return nil, err
	}
	return symLink, nil
}
func (vfs *VirtualFileSystem) CreateReadonlyFile(name string, parentPath string) (*ReadonlyFile, error) {
	readonlyFile := NewReadonlyFile(name, parentPath)
	err := vfs.root.AddItem(readonlyFile)
	if err != nil {
		return nil, err
	}
	return readonlyFile, nil
}

func (vfs *VirtualFileSystem) RemoveItem(name string) error {
	return vfs.root.RemoveItem(name)
}

func (vfs *VirtualFileSystem) ListItems() []FileSystemItem {
	return vfs.root.Items()
}

func (vfs *VirtualFileSystem) FindItem(name string) (FileSystemItem, error) {
	for _, item := range vfs.root.Items() {
		if item.Name() == name {
			return item, nil
		}
	}
	return nil, ErrItemNotFound
}

func (vfs *VirtualFileSystem) Open(name string) (Readable, error) {
	item, err := vfs.FindItem(name)
	if err != nil {
		return nil, err
	}
	if readable, ok := item.(Readable); ok {
		return readable, nil
	}
	return nil, ErrNotImplemented
}

func (vfs *VirtualFileSystem) Write(name string, data []byte) (int, error) {
	item, err := vfs.FindItem(name)
	if err != nil {
		return 0, err
	}
	if writable, ok := item.(Writable); ok {
		return writable.Write(data)
	}
	return 0, ErrNotImplemented
}