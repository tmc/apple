package ext4

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

// nodeKind classifies a source tree entry.
type nodeKind int

const (
	kindDir nodeKind = iota
	kindReg
	kindSymlink
)

// node is an entry in the in-memory source tree.
//
// Hardlinks collapse to a single node referenced by multiple dir entries; ino
// (the assigned ext4 inode number) is filled in during layout.
type node struct {
	name     string // base name (empty for root)
	kind     nodeKind
	mode     uint16 // raw ext4 mode bits (type + permission)
	uid, gid uint32
	size     int64 // file size, or symlink target length
	atime    time.Time
	mtime    time.Time
	ctime    time.Time

	link   string  // symlink target
	path   string  // absolute source path (regular files: read at write time)
	shared *shared // non-nil when this node is a hardlinked regular file

	children []*node // directories, sorted by name

	ino uint32 // assigned during layout
}

// shared tracks a set of hardlinked regular files that share one inode.
type shared struct {
	path  string // source path to read content from
	links int    // number of directory entries referring to the inode
	ino   uint32 // assigned during layout
}

// walk reads the directory tree rooted at dir into a node tree.
func walk(dir string) (*node, error) {
	hard := make(map[hardKey]*shared)
	root, err := walkNode(dir, "", hard)
	if err != nil {
		return nil, err
	}
	if root == nil || root.kind != kindDir {
		return nil, fmt.Errorf("rootfs %s is not a directory", dir)
	}
	return root, nil
}

type hardKey struct {
	dev uint64
	ino uint64
}

func walkNode(path, name string, hard map[hardKey]*shared) (*node, error) {
	fi, err := os.Lstat(path)
	if err != nil {
		return nil, fmt.Errorf("lstat %s: %w", path, err)
	}
	st, haveStat := statOf(fi)

	// ctime is set to mtime: a freshly materialized filesystem has no
	// meaningful inode-change history, and the host ctime is not preservable
	// (it tracks metadata operations and is not deterministic across builds).
	atime := fi.ModTime()
	if haveStat {
		atime = st.atime
	}
	n := &node{
		name:  name,
		mode:  modeBits(fi),
		atime: atime,
		mtime: fi.ModTime(),
		ctime: fi.ModTime(),
	}
	if haveStat {
		n.uid, n.gid = st.uid, st.gid
	}

	switch {
	case fi.IsDir():
		n.kind = kindDir
		entries, err := os.ReadDir(path)
		if err != nil {
			return nil, fmt.Errorf("readdir %s: %w", path, err)
		}
		names := make([]string, 0, len(entries))
		for _, e := range entries {
			names = append(names, e.Name())
		}
		sort.Strings(names)
		for _, childName := range names {
			child, err := walkNode(filepath.Join(path, childName), childName, hard)
			if err != nil {
				return nil, err
			}
			if child != nil {
				n.children = append(n.children, child)
			}
		}

	case fi.Mode()&os.ModeSymlink != 0:
		target, err := os.Readlink(path)
		if err != nil {
			return nil, fmt.Errorf("readlink %s: %w", path, err)
		}
		n.kind = kindSymlink
		n.link = target
		n.size = int64(len(target))

	case fi.Mode().IsRegular():
		n.kind = kindReg
		n.size = fi.Size()
		n.path = path
		if haveStat && st.nlink > 1 {
			key := hardKey{dev: st.dev, ino: st.ino}
			sh, ok := hard[key]
			if !ok {
				sh = &shared{path: path}
				hard[key] = sh
			}
			sh.links++
			n.shared = sh
		}

	default:
		// Device nodes, FIFOs, sockets: skip defensively.
		return nil, nil
	}
	return n, nil
}

// modeBits returns the raw ext4 i_mode value (file-type bits | permissions).
func modeBits(fi os.FileInfo) uint16 {
	perm := uint16(fi.Mode().Perm())
	// setuid/setgid/sticky
	if fi.Mode()&os.ModeSetuid != 0 {
		perm |= 0o4000
	}
	if fi.Mode()&os.ModeSetgid != 0 {
		perm |= 0o2000
	}
	if fi.Mode()&os.ModeSticky != 0 {
		perm |= 0o1000
	}
	switch {
	case fi.IsDir():
		return sIFDIR | perm
	case fi.Mode()&os.ModeSymlink != 0:
		return sIFLNK | perm
	default:
		return sIFREG | perm
	}
}
