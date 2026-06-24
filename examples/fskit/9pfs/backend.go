//go:build darwin

package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"path"
	"sort"
	"strings"

	"9fans.net/go/plan9"
	"9fans.net/go/plan9/client"
	p9 "github.com/hugelgupf/p9/p9"
)

var errUnsupported = errors.ErrUnsupported

type backend interface {
	Close() error
	Stat(name string) (nodeInfo, error)
	ReadDir(name string) ([]nodeInfo, error)
	ReadFile(name string) ([]byte, error)
	// ReadFileAt reads up to len(buf) bytes from name at offset, returning
	// the number of bytes read. A short read or reading past the end of the
	// file returns the bytes available without an error.
	ReadFileAt(name string, offset int64, buf []byte) (int, error)
	WriteFile(name string, offset int64, data []byte) (int, error)
	Create(name string, mode uint32, directory bool) (nodeInfo, error)
	CreateSymlink(name, target string) (nodeInfo, error)
	CreateLink(oldName, newName string) (nodeInfo, error)
	Readlink(name string) (string, error)
	Remove(name string) error
	Rename(oldName, newName string) error
	SetAttr(name string, attr setAttr) (nodeInfo, error)
	GetXattr(name, attr string) ([]byte, error)
	SetXattr(name, attr string, data []byte) error
	ListXattr(name string) ([]string, error)
	RemoveXattr(name, attr string) error
	Preallocate(name string, offset int64, length uint64) (uint64, error)
}

type nodeInfo struct {
	Name     string
	Mode     uint32
	Length   uint64
	Modified uint64
}

type setAttr struct {
	Mode     *uint32
	Size     *uint64
	Accessed *uint64
	Modified *uint64
}

const modeDirectory = uint32(plan9.DMDIR)

type ninePBackend struct {
	fs *client.Fsys
}

func dialBackend(dialect, network, addr, aname string) (backend, error) {
	switch dialect {
	case "9p", "9p2000", "plan9":
		b, err := dial9P(network, addr, aname)
		if err != nil {
			return nil, err
		}
		return errnoBackend{b}, nil
	case "9p2000l", "linux":
		b, err := dial9P2000L(network, addr, aname)
		if err != nil {
			return nil, err
		}
		return errnoBackend{b}, nil
	default:
		return nil, fmt.Errorf("unknown 9p dialect %q", dialect)
	}
}

func dial9P(network, addr, aname string) (*ninePBackend, error) {
	fs, err := mount9P(network, addr, aname)
	if err != nil {
		return nil, err
	}
	return &ninePBackend{fs: fs}, nil
}

func (b *ninePBackend) Close() error {
	return b.fs.Close()
}

func (b *ninePBackend) Stat(name string) (nodeInfo, error) {
	d, err := b.fs.Stat(clean9PPath(name))
	if err != nil {
		return nodeInfo{}, fmt.Errorf("stat 9p %s: %w", name, err)
	}
	return nodeInfoFromPlan9Dir(d), nil
}

func (b *ninePBackend) ReadDir(name string) ([]nodeInfo, error) {
	fid, err := b.fs.Open(clean9PPath(name), plan9.OREAD)
	if err != nil {
		return nil, fmt.Errorf("open 9p dir %s: %w", name, err)
	}
	defer fid.Close()
	entries, err := fid.Dirreadall()
	if err != nil {
		return nil, fmt.Errorf("read 9p dir %s: %w", name, err)
	}
	nodes := make([]nodeInfo, 0, len(entries))
	for _, entry := range entries {
		nodes = append(nodes, nodeInfoFromPlan9Dir(entry))
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Name < nodes[j].Name
	})
	return nodes, nil
}

func (b *ninePBackend) ReadFile(name string) ([]byte, error) {
	fid, err := b.fs.Open(clean9PPath(name), plan9.OREAD)
	if err != nil {
		return nil, fmt.Errorf("open 9p file %s: %w", name, err)
	}
	defer fid.Close()
	data, err := io.ReadAll(fid)
	if err != nil {
		return nil, fmt.Errorf("read 9p file %s: %w", name, err)
	}
	return data, nil
}

func (b *ninePBackend) ReadFileAt(name string, offset int64, buf []byte) (int, error) {
	fid, err := b.fs.Open(clean9PPath(name), plan9.OREAD)
	if err != nil {
		return 0, fmt.Errorf("open 9p file %s: %w", name, err)
	}
	defer fid.Close()
	n, err := fid.ReadAt(buf, offset)
	if err != nil && err != io.EOF {
		return n, fmt.Errorf("read 9p file %s: %w", name, err)
	}
	return n, nil
}

func (b *ninePBackend) WriteFile(name string, offset int64, data []byte) (int, error) {
	fid, err := b.fs.Open(clean9PPath(name), plan9.OWRITE)
	if err != nil {
		return 0, fmt.Errorf("open 9p file for write %s: %w", name, err)
	}
	defer fid.Close()
	n, err := fid.WriteAt(data, offset)
	if err != nil {
		return n, fmt.Errorf("write 9p file %s: %w", name, err)
	}
	return n, nil
}

func (b *ninePBackend) Create(name string, mode uint32, directory bool) (nodeInfo, error) {
	perm := plan9.Perm(mode & 0777)
	if directory {
		perm |= plan9.DMDIR
	}
	fid, err := b.fs.Create(clean9PPath(name), plan9.ORDWR, perm)
	if err != nil {
		return nodeInfo{}, fmt.Errorf("create 9p %s: %w", name, err)
	}
	_ = fid.Close()
	return b.Stat(name)
}

func (b *ninePBackend) CreateSymlink(name, target string) (nodeInfo, error) {
	return nodeInfo{}, errUnsupported
}

func (b *ninePBackend) CreateLink(oldName, newName string) (nodeInfo, error) {
	return nodeInfo{}, errUnsupported
}

func (b *ninePBackend) Readlink(name string) (string, error) {
	return "", errUnsupported
}

func (b *ninePBackend) Remove(name string) error {
	if err := b.fs.Remove(clean9PPath(name)); err != nil {
		return fmt.Errorf("remove 9p %s: %w", name, err)
	}
	return nil
}

func (b *ninePBackend) Rename(oldName, newName string) error {
	d := new(plan9.Dir)
	d.Null()
	d.Name = path.Base(clean9PPath(newName))
	if err := b.fs.Wstat(clean9PPath(oldName), d); err != nil {
		return fmt.Errorf("rename 9p %s to %s: %w", oldName, newName, err)
	}
	return nil
}

func (b *ninePBackend) SetAttr(name string, attr setAttr) (nodeInfo, error) {
	d := new(plan9.Dir)
	d.Null()
	if attr.Mode != nil {
		d.Mode = plan9.Perm(*attr.Mode)
	}
	if attr.Size != nil {
		d.Length = *attr.Size
	}
	if attr.Accessed != nil {
		d.Atime = uint32(*attr.Accessed)
	}
	if attr.Modified != nil {
		d.Mtime = uint32(*attr.Modified)
	}
	if err := b.fs.Wstat(clean9PPath(name), d); err != nil {
		return nodeInfo{}, fmt.Errorf("set 9p attrs %s: %w", name, err)
	}
	return b.Stat(name)
}

func (b *ninePBackend) GetXattr(name, attr string) ([]byte, error) {
	return nil, errUnsupported
}

func (b *ninePBackend) SetXattr(name, attr string, data []byte) error {
	return errUnsupported
}

func (b *ninePBackend) ListXattr(name string) ([]string, error) {
	return nil, errUnsupported
}

func (b *ninePBackend) RemoveXattr(name, attr string) error {
	return errUnsupported
}

func (b *ninePBackend) Preallocate(name string, offset int64, length uint64) (uint64, error) {
	return preallocateGrow(b, name, offset, length)
}

// Classic 9P2000 has no symbolic or hard links.
func (b *ninePBackend) supportsSymlinks() bool  { return false }
func (b *ninePBackend) supportsHardLinks() bool { return false }

type p9LBackend struct {
	client *p9.Client
	root   p9.File
}

func dial9P2000L(network, addr, aname string) (*p9LBackend, error) {
	conn, err := net.Dial(network, addr)
	if err != nil {
		return nil, fmt.Errorf("dial 9p2000.l %s %s: %w", network, addr, err)
	}
	client, err := p9.NewClient(conn)
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("negotiate 9p2000.l %s %s: %w", network, addr, err)
	}
	root, err := client.Attach(aname)
	if err != nil {
		client.Close()
		return nil, fmt.Errorf("attach 9p2000.l aname %q: %w", aname, err)
	}
	return &p9LBackend{client: client, root: root}, nil
}

func (b *p9LBackend) Close() error {
	_ = b.root.Close()
	return b.client.Close()
}

func (b *p9LBackend) Stat(name string) (nodeInfo, error) {
	file, err := b.walk(clean9PPath(name))
	if err != nil {
		return nodeInfo{}, err
	}
	defer file.Close()
	_, _, attr, err := file.GetAttr(p9.AttrMaskAll)
	if err != nil {
		return nodeInfo{}, fmt.Errorf("stat 9p2000.l %s: %w", name, err)
	}
	return nodeInfo{Name: path.Base(clean9PPath(name)), Mode: uint32(attr.Mode), Length: attr.Size, Modified: attr.MTimeSeconds}, nil
}

func (b *p9LBackend) ReadDir(name string) ([]nodeInfo, error) {
	file, err := b.walk(clean9PPath(name))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	if _, _, err := file.Open(p9.ReadOnly); err != nil {
		return nil, fmt.Errorf("open 9p2000.l dir %s: %w", name, err)
	}
	var nodes []nodeInfo
	var offset uint64
	for {
		dirents, err := file.Readdir(offset, 64<<10)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("read 9p2000.l dir %s: %w", name, err)
		}
		if len(dirents) == 0 {
			break
		}
		for _, dirent := range dirents {
			child := child9PPath(name, dirent.Name)
			info, err := b.Stat(child)
			if err != nil {
				info = nodeInfo{Name: dirent.Name, Mode: modeFromP9QID(dirent.Type), Length: 0}
			}
			info.Name = dirent.Name
			nodes = append(nodes, info)
			offset = dirent.Offset
		}
	}
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Name < nodes[j].Name
	})
	return nodes, nil
}

func (b *p9LBackend) ReadFile(name string) ([]byte, error) {
	file, err := b.walk(clean9PPath(name))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	if _, _, err := file.Open(p9.ReadOnly); err != nil {
		return nil, fmt.Errorf("open 9p2000.l file %s: %w", name, err)
	}
	var data []byte
	buf := make([]byte, 64<<10)
	var offset int64
	for {
		n, err := file.ReadAt(buf, offset)
		if n > 0 {
			data = append(data, buf[:n]...)
			offset += int64(n)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("read 9p2000.l file %s: %w", name, err)
		}
		if n == 0 {
			break
		}
	}
	return data, nil
}

func (b *p9LBackend) ReadFileAt(name string, offset int64, buf []byte) (int, error) {
	file, err := b.walk(clean9PPath(name))
	if err != nil {
		return 0, err
	}
	defer file.Close()
	if _, _, err := file.Open(p9.ReadOnly); err != nil {
		return 0, fmt.Errorf("open 9p2000.l file %s: %w", name, err)
	}
	n, err := file.ReadAt(buf, offset)
	if err != nil && err != io.EOF {
		return n, fmt.Errorf("read 9p2000.l file %s: %w", name, err)
	}
	return n, nil
}

func (b *p9LBackend) WriteFile(name string, offset int64, data []byte) (int, error) {
	file, err := b.walk(clean9PPath(name))
	if err != nil {
		return 0, err
	}
	defer file.Close()
	if _, _, err := file.Open(p9.ReadWrite); err != nil {
		return 0, fmt.Errorf("open 9p2000.l file for write %s: %w", name, err)
	}
	n, err := file.WriteAt(data, offset)
	if err != nil {
		return n, fmt.Errorf("write 9p2000.l file %s: %w", name, err)
	}
	return n, nil
}

func (b *p9LBackend) Create(name string, mode uint32, directory bool) (nodeInfo, error) {
	parentName := path.Dir(clean9PPath(name))
	base := path.Base(clean9PPath(name))
	if parentName == "." {
		parentName = ""
	}
	parent, err := b.walk(parentName)
	if err != nil {
		return nodeInfo{}, err
	}
	defer parent.Close()
	if directory {
		if _, err := parent.Mkdir(base, p9.FileMode(mode&0777), p9.NoUID, p9.NoGID); err != nil {
			return nodeInfo{}, fmt.Errorf("mkdir 9p2000.l %s: %w", name, err)
		}
		return b.Stat(name)
	}
	file, _, _, err := parent.Create(base, p9.ReadWrite, p9.FileMode(mode&0777), p9.NoUID, p9.NoGID)
	if err != nil {
		return nodeInfo{}, fmt.Errorf("create 9p2000.l %s: %w", name, err)
	}
	_ = file.Close()
	return b.Stat(name)
}

func (b *p9LBackend) CreateSymlink(name, target string) (nodeInfo, error) {
	parentName := path.Dir(clean9PPath(name))
	base := path.Base(clean9PPath(name))
	if parentName == "." {
		parentName = ""
	}
	parent, err := b.walk(parentName)
	if err != nil {
		return nodeInfo{}, err
	}
	defer parent.Close()
	if _, err := parent.Symlink(target, base, p9.NoUID, p9.NoGID); err != nil {
		return nodeInfo{}, fmt.Errorf("symlink 9p2000.l %s to %s: %w", name, target, err)
	}
	return b.Stat(name)
}

func (b *p9LBackend) CreateLink(oldName, newName string) (nodeInfo, error) {
	oldFile, err := b.walk(clean9PPath(oldName))
	if err != nil {
		return nodeInfo{}, err
	}
	defer oldFile.Close()
	parentName := path.Dir(clean9PPath(newName))
	base := path.Base(clean9PPath(newName))
	if parentName == "." {
		parentName = ""
	}
	parent, err := b.walk(parentName)
	if err != nil {
		return nodeInfo{}, err
	}
	defer parent.Close()
	if err := parent.Link(oldFile, base); err != nil {
		return nodeInfo{}, fmt.Errorf("link 9p2000.l %s to %s: %w", oldName, newName, err)
	}
	return b.Stat(newName)
}

func (b *p9LBackend) Readlink(name string) (string, error) {
	file, err := b.walk(clean9PPath(name))
	if err != nil {
		return "", err
	}
	defer file.Close()
	target, err := file.Readlink()
	if err != nil {
		return "", fmt.Errorf("readlink 9p2000.l %s: %w", name, err)
	}
	return target, nil
}

func (b *p9LBackend) Remove(name string) error {
	parentName := path.Dir(clean9PPath(name))
	base := path.Base(clean9PPath(name))
	if parentName == "." {
		parentName = ""
	}
	parent, err := b.walk(parentName)
	if err != nil {
		return err
	}
	defer parent.Close()
	if err := parent.UnlinkAt(base, 0); err != nil {
		return fmt.Errorf("remove 9p2000.l %s: %w", name, err)
	}
	return nil
}

func (b *p9LBackend) Rename(oldName, newName string) error {
	oldParentName := path.Dir(clean9PPath(oldName))
	oldBase := path.Base(clean9PPath(oldName))
	newParentName := path.Dir(clean9PPath(newName))
	newBase := path.Base(clean9PPath(newName))
	if oldParentName == "." {
		oldParentName = ""
	}
	if newParentName == "." {
		newParentName = ""
	}
	oldParent, err := b.walk(oldParentName)
	if err != nil {
		return err
	}
	defer oldParent.Close()
	newParent, err := b.walk(newParentName)
	if err != nil {
		return err
	}
	defer newParent.Close()
	if err := oldParent.RenameAt(oldBase, newParent, newBase); err != nil {
		return fmt.Errorf("rename 9p2000.l %s to %s: %w", oldName, newName, err)
	}
	return nil
}

func (b *p9LBackend) SetAttr(name string, attr setAttr) (nodeInfo, error) {
	file, err := b.walk(clean9PPath(name))
	if err != nil {
		return nodeInfo{}, err
	}
	defer file.Close()
	var mask p9.SetAttrMask
	var set p9.SetAttr
	if attr.Mode != nil {
		mask.Permissions = true
		set.Permissions = p9.FileMode(*attr.Mode & 0777)
	}
	if attr.Size != nil {
		mask.Size = true
		set.Size = *attr.Size
	}
	if attr.Accessed != nil {
		mask.ATime = true
		set.ATimeSeconds = *attr.Accessed
	}
	if attr.Modified != nil {
		mask.MTime = true
		set.MTimeSeconds = *attr.Modified
	}
	if mask.Empty() {
		return b.Stat(name)
	}
	if err := file.SetAttr(mask, set); err != nil {
		return nodeInfo{}, fmt.Errorf("set 9p2000.l attrs %s: %w", name, err)
	}
	return b.Stat(name)
}

func (b *p9LBackend) GetXattr(name, attr string) ([]byte, error) {
	file, err := b.walk(clean9PPath(name))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	data, err := file.GetXattr(attr)
	if err != nil {
		return nil, fmt.Errorf("get 9p2000.l xattr %s %s: %w", name, attr, err)
	}
	return data, nil
}

func (b *p9LBackend) SetXattr(name, attr string, data []byte) error {
	file, err := b.walk(clean9PPath(name))
	if err != nil {
		return err
	}
	defer file.Close()
	if err := file.SetXattr(attr, data, 0); err != nil {
		return fmt.Errorf("set 9p2000.l xattr %s %s: %w", name, attr, err)
	}
	return nil
}

func (b *p9LBackend) ListXattr(name string) ([]string, error) {
	file, err := b.walk(clean9PPath(name))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	names, err := file.ListXattrs()
	if err != nil {
		return nil, fmt.Errorf("list 9p2000.l xattrs %s: %w", name, err)
	}
	return names, nil
}

func (b *p9LBackend) RemoveXattr(name, attr string) error {
	file, err := b.walk(clean9PPath(name))
	if err != nil {
		return err
	}
	defer file.Close()
	if err := file.RemoveXattr(attr); err != nil {
		return fmt.Errorf("remove 9p2000.l xattr %s %s: %w", name, attr, err)
	}
	return nil
}

func (b *p9LBackend) Preallocate(name string, offset int64, length uint64) (uint64, error) {
	return preallocateGrow(b, name, offset, length)
}

// 9P2000.L has both symbolic links and hard links.
func (b *p9LBackend) supportsSymlinks() bool  { return true }
func (b *p9LBackend) supportsHardLinks() bool { return true }

// preallocateGrow emulates preallocation on backends whose only size control
// is truncation. It grows the file to offset+length when needed and never
// shrinks an already larger file.
//
// It runs against the raw backend, so any error it returns is untranslated;
// the caller (errnoBackend.Preallocate) re-wraps the result with backendError,
// which applies the errno translation. Any errno inspection added here would
// need to account for that.
func preallocateGrow(b backend, name string, offset int64, length uint64) (uint64, error) {
	info, err := b.Stat(name)
	if err != nil {
		return 0, err
	}
	size := uint64(offset) + length
	if info.Length >= size {
		return length, nil
	}
	if _, err := b.SetAttr(name, setAttr{Size: &size}); err != nil {
		return 0, err
	}
	return length, nil
}

func (b *p9LBackend) walk(name string) (p9.File, error) {
	if name == "" {
		_, file, err := b.root.Walk(nil)
		if err != nil {
			return nil, fmt.Errorf("walk 9p2000.l root: %w", err)
		}
		return file, nil
	}
	parts := strings.Split(clean9PPath(name), "/")
	_, file, err := b.root.Walk(parts)
	if err != nil {
		return nil, fmt.Errorf("walk 9p2000.l %s: %w", name, err)
	}
	return file, nil
}

func nodeInfoFromPlan9Dir(d *plan9.Dir) nodeInfo {
	return nodeInfo{Name: d.Name, Mode: uint32(d.Mode), Length: d.Length, Modified: uint64(d.Mtime)}
}

func modeFromP9QID(qtype p9.QIDType) uint32 {
	if qtype&p9.TypeDir != 0 {
		return uint32(plan9.DMDIR | 0555)
	}
	return 0444
}

func clean9PPath(name string) string {
	name = strings.TrimPrefix(path.Clean("/"+name), "/")
	if name == "." {
		return ""
	}
	return name
}

func child9PPath(parent, child string) string {
	if clean9PPath(parent) == "" {
		return child
	}
	return path.Join(parent, child)
}

func modeString(mode uint32) string {
	if mode&uint32(plan9.DMDIR) != 0 || p9.FileMode(mode).IsDir() {
		return "d"
	}
	return "-"
}
