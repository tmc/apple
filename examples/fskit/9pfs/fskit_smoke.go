//go:build darwin

package main

import (
	"errors"
	"fmt"
	"io/fs"
	"sort"
	"syscall"
	"unsafe"

	"9fans.net/go/plan9"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/fskit"
	"github.com/tmc/apple/objc"
)

type memBackend struct {
	files map[string]*memNode
}

type memNode struct {
	info     nodeInfo
	data     []byte
	modified uint64
	xattrs   map[string][]byte
}

func newSmokeBackend() *memBackend {
	return &memBackend{files: map[string]*memNode{
		"": {
			info: nodeInfo{Name: "", Mode: modeDirectory | 0555},
		},
		"README": {
			info: nodeInfo{Name: "README", Mode: 0444, Length: uint64(len("hello from 9p\n"))},
			data: []byte("hello from 9p\n"),
		},
	}}
}

func (b *memBackend) Close() error { return nil }

func (b *memBackend) Stat(name string) (nodeInfo, error) {
	node, ok := b.files[clean9PPath(name)]
	if !ok {
		return nodeInfo{}, fs.ErrNotExist
	}
	return node.info, nil
}

func (b *memBackend) ReadDir(name string) ([]nodeInfo, error) {
	if clean9PPath(name) != "" {
		return nil, fs.ErrInvalid
	}
	var entries []nodeInfo
	for name, node := range b.files {
		if name == "" {
			continue
		}
		entries = append(entries, node.info)
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Name < entries[j].Name
	})
	return entries, nil
}

func (b *memBackend) ReadFile(name string) ([]byte, error) {
	node, ok := b.files[clean9PPath(name)]
	if !ok {
		return nil, fs.ErrNotExist
	}
	return append([]byte(nil), node.data...), nil
}

func (b *memBackend) ReadFileAt(name string, offset int64, buf []byte) (int, error) {
	node, ok := b.files[clean9PPath(name)]
	if !ok {
		return 0, fs.ErrNotExist
	}
	if offset < 0 {
		return 0, fs.ErrInvalid
	}
	if offset >= int64(len(node.data)) {
		return 0, nil
	}
	return copy(buf, node.data[offset:]), nil
}

func (b *memBackend) WriteFile(name string, offset int64, data []byte) (int, error) {
	node, ok := b.files[clean9PPath(name)]
	if !ok {
		return 0, fs.ErrNotExist
	}
	if offset < 0 {
		return 0, fs.ErrInvalid
	}
	end := int(offset) + len(data)
	if end > len(node.data) {
		next := make([]byte, end)
		copy(next, node.data)
		node.data = next
	}
	copy(node.data[offset:], data)
	node.info.Length = uint64(len(node.data))
	return len(data), nil
}

func (b *memBackend) Create(name string, mode uint32, directory bool) (nodeInfo, error) {
	name = clean9PPath(name)
	if _, ok := b.files[name]; ok {
		return nodeInfo{}, fs.ErrExist
	}
	info := nodeInfo{Name: pathBase(name), Mode: mode, Length: 0}
	if directory {
		info.Mode |= modeDirectory
	}
	b.files[name] = &memNode{info: info}
	return info, nil
}

func (b *memBackend) CreateSymlink(name, target string) (nodeInfo, error) {
	name = clean9PPath(name)
	if _, ok := b.files[name]; ok {
		return nodeInfo{}, fs.ErrExist
	}
	info := nodeInfo{Name: pathBase(name), Mode: uint32(plan9.DMSYMLINK | 0777), Length: uint64(len(target))}
	b.files[name] = &memNode{info: info, data: []byte(target)}
	return info, nil
}

func (b *memBackend) CreateLink(oldName, newName string) (nodeInfo, error) {
	oldName = clean9PPath(oldName)
	newName = clean9PPath(newName)
	node, ok := b.files[oldName]
	if !ok {
		return nodeInfo{}, fs.ErrNotExist
	}
	info := node.info
	info.Name = pathBase(newName)
	b.files[newName] = &memNode{info: info, data: node.data}
	return info, nil
}

func (b *memBackend) Readlink(name string) (string, error) {
	node, ok := b.files[clean9PPath(name)]
	if !ok {
		return "", fs.ErrNotExist
	}
	if node.info.Mode&uint32(plan9.DMSYMLINK) == 0 {
		return "", fs.ErrInvalid
	}
	return string(node.data), nil
}

func (b *memBackend) Remove(name string) error {
	delete(b.files, clean9PPath(name))
	return nil
}

func (b *memBackend) Rename(oldName, newName string) error {
	oldName = clean9PPath(oldName)
	newName = clean9PPath(newName)
	node, ok := b.files[oldName]
	if !ok {
		return fs.ErrNotExist
	}
	delete(b.files, oldName)
	node.info.Name = pathBase(newName)
	b.files[newName] = node
	return nil
}

func (b *memBackend) SetAttr(name string, attr setAttr) (nodeInfo, error) {
	node, ok := b.files[clean9PPath(name)]
	if !ok {
		return nodeInfo{}, fs.ErrNotExist
	}
	if attr.Mode != nil {
		node.info.Mode = *attr.Mode
	}
	if attr.Size != nil {
		size := int(*attr.Size)
		if size < len(node.data) {
			node.data = node.data[:size]
		} else if size > len(node.data) {
			next := make([]byte, size)
			copy(next, node.data)
			node.data = next
		}
		node.info.Length = uint64(size)
	}
	if attr.Modified != nil {
		node.modified = *attr.Modified
		node.info.Modified = *attr.Modified
	}
	return node.info, nil
}

func (b *memBackend) GetXattr(name, attr string) ([]byte, error) {
	node, ok := b.files[clean9PPath(name)]
	if !ok {
		return nil, fs.ErrNotExist
	}
	data, ok := node.xattrs[attr]
	if !ok {
		return nil, fs.ErrNotExist
	}
	return append([]byte(nil), data...), nil
}

func (b *memBackend) SetXattr(name, attr string, data []byte) error {
	node, ok := b.files[clean9PPath(name)]
	if !ok {
		return fs.ErrNotExist
	}
	if node.xattrs == nil {
		node.xattrs = make(map[string][]byte)
	}
	node.xattrs[attr] = append([]byte(nil), data...)
	return nil
}

func (b *memBackend) ListXattr(name string) ([]string, error) {
	node, ok := b.files[clean9PPath(name)]
	if !ok {
		return nil, fs.ErrNotExist
	}
	var names []string
	for name := range node.xattrs {
		names = append(names, name)
	}
	sort.Strings(names)
	return names, nil
}

func (b *memBackend) RemoveXattr(name, attr string) error {
	node, ok := b.files[clean9PPath(name)]
	if !ok {
		return fs.ErrNotExist
	}
	delete(node.xattrs, attr)
	return nil
}

func (b *memBackend) Preallocate(name string, offset int64, length uint64) (uint64, error) {
	return preallocateGrow(b, name, offset, length)
}

// The in-memory backend supports symbolic and hard links.
func (b *memBackend) supportsSymlinks() bool  { return true }
func (b *memBackend) supportsHardLinks() bool { return true }

func pathBase(name string) string {
	for i := len(name) - 1; i >= 0; i-- {
		if name[i] == '/' {
			return name[i+1:]
		}
	}
	return name
}

func bytesFromNSData(id objc.ID) []byte {
	if id == 0 {
		return nil
	}
	data := foundation.NSDataFromID(id)
	n := data.Length()
	if n == 0 {
		return nil
	}
	src := data.Bytes()
	out := make([]byte, n)
	copy(out, unsafe.Slice((*byte)(src), n))
	return out
}

func fskitSmoke() error {
	var err error
	objc.AutoreleasePool(func() {
		err = fskitSmokeInPool()
	})
	return err
}

func fskitSmokeInPool() error {
	backend := newSmokeBackend()
	server, err := ensureServer(0, &ninepFileSystem{backend: backend})
	if err != nil {
		return err
	}
	fs := server.NewFileSystem()
	defer objc.Send[struct{}](fs, objc.Sel("release"))

	var callbackErr error
	var probeName string
	probeReply := objc.NewBlock(func(_ objc.Block, result objc.ID, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("probe returned error %d", errID)
			return
		}
		probeName = fskit.FSProbeResultFromID(result).Name()
	})
	defer probeReply.Release()
	objc.Send[struct{}](fs, objc.Sel("probeResource:replyHandler:"), objc.ID(0), objc.ID(probeReply))
	if callbackErr != nil {
		return callbackErr
	}
	if probeName != "9pfs" {
		return fmt.Errorf("probe name = %q, want 9pfs", probeName)
	}

	var volume objc.ID
	loadReply := objc.NewBlock(func(_ objc.Block, result objc.ID, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("load returned error %d", errID)
			return
		}
		volume = result
	})
	defer loadReply.Release()
	objc.Send[struct{}](fs, objc.Sel("loadResource:options:replyHandler:"), objc.ID(0), objc.ID(0), objc.ID(loadReply))
	if callbackErr != nil {
		return callbackErr
	}
	if volume == 0 {
		return errors.New("load returned nil volume")
	}

	var root objc.ID
	activateReply := objc.NewBlock(func(_ objc.Block, item objc.ID, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("activate returned error %d", errID)
			return
		}
		root = item
	})
	defer activateReply.Release()
	objc.Send[struct{}](volume, objc.Sel("activateWithOptions:replyHandler:"), objc.ID(0), objc.ID(activateReply))
	if callbackErr != nil {
		return callbackErr
	}
	if root == 0 {
		return errors.New("activate returned nil root")
	}

	var readme objc.ID
	name := fskit.NewFileNameWithString("README")
	lookupReply := objc.NewBlock(func(_ objc.Block, item objc.ID, actualName objc.ID, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("lookup returned error %d", errID)
			return
		}
		readme = item
	})
	defer lookupReply.Release()
	objc.Send[struct{}](volume, objc.Sel("lookupItemNamed:inDirectory:replyHandler:"), name.GetID(), root, objc.ID(lookupReply))
	if callbackErr != nil {
		return callbackErr
	}
	if readme == 0 {
		return errors.New("lookup returned nil README")
	}

	var size uint64
	attrsReply := objc.NewBlock(func(_ objc.Block, attrs objc.ID, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("get attributes returned error %d", errID)
			return
		}
		size = fskit.FSItemAttributesFromID(attrs).Size()
	})
	defer attrsReply.Release()
	objc.Send[struct{}](volume, objc.Sel("getAttributes:ofItem:replyHandler:"), objc.ID(0), readme, objc.ID(attrsReply))
	if callbackErr != nil {
		return callbackErr
	}
	if size != uint64(len("hello from 9p\n")) {
		return fmt.Errorf("README size = %d, want %d", size, len("hello from 9p\n"))
	}

	var symlink objc.ID
	linkName := fskit.NewFileNameWithString("readme-link")
	linkTarget := fskit.NewFileNameWithString("README")
	createSymlinkReply := objc.NewBlock(func(_ objc.Block, item objc.ID, actualName objc.ID, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("create symlink returned error %d", errID)
			return
		}
		symlink = item
		if fskit.FSFileNameFromID(actualName).String() != "readme-link" {
			callbackErr = fmt.Errorf("create symlink returned unexpected name")
		}
	})
	defer createSymlinkReply.Release()
	objc.Send[struct{}](volume, objc.Sel("createSymbolicLinkNamed:inDirectory:attributes:linkContents:replyHandler:"), linkName.GetID(), root, objc.ID(0), linkTarget.GetID(), objc.ID(createSymlinkReply))
	if callbackErr != nil {
		return callbackErr
	}
	if symlink == 0 {
		return errors.New("create symlink returned nil item")
	}

	var symlinkTarget string
	readSymlinkReply := objc.NewBlock(func(_ objc.Block, target objc.ID, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("read symlink returned error %d", errID)
			return
		}
		symlinkTarget = fskit.FSFileNameFromID(target).String()
	})
	defer readSymlinkReply.Release()
	objc.Send[struct{}](volume, objc.Sel("readSymbolicLink:replyHandler:"), symlink, objc.ID(readSymlinkReply))
	if callbackErr != nil {
		return callbackErr
	}
	if symlinkTarget != "README" {
		return fmt.Errorf("symlink target = %q, want README", symlinkTarget)
	}

	hardlinkName := fskit.NewFileNameWithString("README.hard")
	createLinkReply := objc.NewBlock(func(_ objc.Block, actualName objc.ID, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("create hardlink returned error %d", errID)
			return
		}
		if fskit.FSFileNameFromID(actualName).String() != "README.hard" {
			callbackErr = fmt.Errorf("create hardlink returned unexpected name")
		}
	})
	defer createLinkReply.Release()
	objc.Send[struct{}](volume, objc.Sel("createLinkToItem:named:inDirectory:replyHandler:"), readme, hardlinkName.GetID(), root, objc.ID(createLinkReply))
	if callbackErr != nil {
		return callbackErr
	}

	var created objc.ID
	createName := fskit.NewFileNameWithString("created.txt")
	createReply := objc.NewBlock(func(_ objc.Block, item objc.ID, actualName objc.ID, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("create returned error %d", errID)
			return
		}
		created = item
	})
	defer createReply.Release()
	objc.Send[struct{}](volume, objc.Sel("createItemNamed:type:inDirectory:attributes:replyHandler:"), createName.GetID(), fskit.FSItemTypeFile, root, objc.ID(0), objc.ID(createReply))
	if callbackErr != nil {
		return callbackErr
	}
	if created == 0 {
		return errors.New("create returned nil item")
	}

	payload := []byte("created through fskit\n")
	writeReply := objc.NewBlock(func(_ objc.Block, n uintptr, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("write returned error %d", errID)
			return
		}
		if n != uintptr(len(payload)) {
			callbackErr = fmt.Errorf("write returned %d, want %d", n, len(payload))
		}
	})
	defer writeReply.Release()
	data := foundation.NewDataWithBytesLength(payload)
	objc.Send[struct{}](volume, objc.Sel("writeContents:toFile:atOffset:replyHandler:"), data.GetID(), created, int64(0), objc.ID(writeReply))
	if callbackErr != nil {
		return callbackErr
	}

	readBuffer := foundation.NewMutableDataWithLength(uint(len(payload)))
	readReply := objc.NewBlock(func(_ objc.Block, n uintptr, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("read returned error %d", errID)
			return
		}
		if n != uintptr(len(payload)) {
			callbackErr = fmt.Errorf("read returned %d, want %d", n, len(payload))
		}
	})
	defer readReply.Release()
	objc.Send[struct{}](volume, objc.Sel("readFromFile:offset:length:intoBuffer:replyHandler:"), created, int64(0), uintptr(len(payload)), readBuffer.GetID(), objc.ID(readReply))
	if callbackErr != nil {
		return callbackErr
	}
	if got := string(bytesFromNSData(readBuffer.GetID())); got != string(payload) {
		return fmt.Errorf("read returned %q, want %q", got, payload)
	}

	xattrName := fskit.NewFileNameWithString("user.codex")
	xattrValue := []byte("xattr through fskit")
	xattrData := foundation.NewDataWithBytesLength(xattrValue)
	setXattrReply := objc.NewBlock(func(_ objc.Block, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("set xattr returned error %d", errID)
		}
	})
	defer setXattrReply.Release()
	objc.Send[struct{}](volume, objc.Sel("setXattrNamed:toData:onItem:policy:replyHandler:"), xattrName.GetID(), xattrData.GetID(), created, fskit.FSSetXattrPolicyAlwaysSet, objc.ID(setXattrReply))
	if callbackErr != nil {
		return callbackErr
	}
	objc.Send[struct{}](volume, objc.Sel("setXattrNamed:toData:onItem:policy:replyHandler:"), xattrName.GetID(), xattrData.GetID(), created, fskit.FSSetXattrPolicyMustCreate, objc.ID(setXattrReply))
	if callbackErr == nil {
		return errors.New("must-create xattr replaced an existing value")
	}
	callbackErr = nil
	objc.Send[struct{}](volume, objc.Sel("setXattrNamed:toData:onItem:policy:replyHandler:"), xattrName.GetID(), xattrData.GetID(), created, fskit.FSSetXattrPolicyMustReplace, objc.ID(setXattrReply))
	if callbackErr != nil {
		return callbackErr
	}

	var listedXattr string
	listXattrsReply := objc.NewBlock(func(_ objc.Block, arrayID objc.ID, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("list xattrs returned error %d", errID)
			return
		}
		array := foundation.NSArrayFromID(arrayID)
		if array.Count() != 1 {
			callbackErr = fmt.Errorf("list xattrs count = %d, want 1", array.Count())
			return
		}
		listedXattr = fskit.FSFileNameFromID(array.ObjectAtIndex(0).GetID()).String()
	})
	defer listXattrsReply.Release()
	objc.Send[struct{}](volume, objc.Sel("listXattrsOfItem:replyHandler:"), created, objc.ID(listXattrsReply))
	if callbackErr != nil {
		return callbackErr
	}
	if listedXattr != "user.codex" {
		return fmt.Errorf("listed xattr = %q, want user.codex", listedXattr)
	}

	var gotXattr []byte
	getXattrReply := objc.NewBlock(func(_ objc.Block, dataID objc.ID, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("get xattr returned error %d", errID)
			return
		}
		gotXattr = bytesFromNSData(dataID)
	})
	defer getXattrReply.Release()
	objc.Send[struct{}](volume, objc.Sel("getXattrNamed:ofItem:replyHandler:"), xattrName.GetID(), created, objc.ID(getXattrReply))
	if callbackErr != nil {
		return callbackErr
	}
	if string(gotXattr) != string(xattrValue) {
		return fmt.Errorf("xattr value = %q, want %q", gotXattr, xattrValue)
	}
	removeXattrReply := objc.NewBlock(func(_ objc.Block, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("remove xattr returned error %d", errID)
		}
	})
	defer removeXattrReply.Release()
	objc.Send[struct{}](volume, objc.Sel("setXattrNamed:toData:onItem:policy:replyHandler:"), xattrName.GetID(), objc.ID(0), created, fskit.FSSetXattrPolicyDelete, objc.ID(removeXattrReply))
	if callbackErr != nil {
		return callbackErr
	}
	objc.Send[struct{}](volume, objc.Sel("setXattrNamed:toData:onItem:policy:replyHandler:"), xattrName.GetID(), xattrData.GetID(), created, fskit.FSSetXattrPolicyMustReplace, objc.ID(setXattrReply))
	if callbackErr == nil {
		return errors.New("must-replace xattr created a missing value")
	}
	callbackErr = nil
	emptyXattrsReply := objc.NewBlock(func(_ objc.Block, arrayID objc.ID, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("empty list xattrs returned error %d", errID)
			return
		}
		array := foundation.NSArrayFromID(arrayID)
		if array.Count() != 0 {
			callbackErr = fmt.Errorf("empty list xattrs count = %d, want 0", array.Count())
		}
	})
	defer emptyXattrsReply.Release()
	objc.Send[struct{}](volume, objc.Sel("listXattrsOfItem:replyHandler:"), created, objc.ID(emptyXattrsReply))
	if callbackErr != nil {
		return callbackErr
	}

	truncated := uint64(len("created"))
	setAttrs := fskit.NewFSItemSetAttributesRequest()
	setAttrs.SetSize(truncated)
	setAttrs.SetConsumedAttributes(0)
	setAttrsReply := objc.NewBlock(func(_ objc.Block, attrs objc.ID, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("set attributes returned error %d", errID)
			return
		}
		if got := fskit.FSItemAttributesFromID(attrs).Size(); got != truncated {
			callbackErr = fmt.Errorf("set attributes size = %d, want %d", got, truncated)
		}
	})
	defer setAttrsReply.Release()
	objc.Send[struct{}](volume, objc.Sel("setAttributes:onItem:replyHandler:"), setAttrs.GetID(), created, objc.ID(setAttrsReply))
	if callbackErr != nil {
		return callbackErr
	}
	if !setAttrs.WasAttributeConsumed(fskit.FSItemAttributeSize) {
		return errors.New("set attributes did not consume size")
	}

	modified := uint64(1577934245)
	setTimeAttrs := fskit.NewFSItemSetAttributesRequest()
	setTimeAttrs.SetModifyTime(syscall.Timespec{Sec: int64(modified)})
	setTimeAttrs.SetConsumedAttributes(0)
	setTimeAttrsReply := objc.NewBlock(func(_ objc.Block, attrs objc.ID, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("set mtime returned error %d", errID)
		}
	})
	defer setTimeAttrsReply.Release()
	objc.Send[struct{}](volume, objc.Sel("setAttributes:onItem:replyHandler:"), setTimeAttrs.GetID(), created, objc.ID(setTimeAttrsReply))
	if callbackErr != nil {
		return callbackErr
	}
	if !setTimeAttrs.WasAttributeConsumed(fskit.FSItemAttributeModifyTime) {
		return errors.New("set attributes did not consume modify time")
	}
	if got := backend.files["created.txt"].modified; got != modified {
		return fmt.Errorf("mtime = %d, want %d", got, modified)
	}

	preallocated := uint64(32)
	preallocateReply := objc.NewBlock(func(_ objc.Block, n uintptr, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("preallocate returned error %d", errID)
			return
		}
		if uint64(n) != preallocated {
			callbackErr = fmt.Errorf("preallocate length = %d, want %d", n, preallocated)
		}
	})
	defer preallocateReply.Release()
	objc.Send[struct{}](volume, objc.Sel("preallocateSpaceForItem:atOffset:length:flags:replyHandler:"), created, int64(0), uintptr(preallocated), fskit.FSPreallocateFlags(0), objc.ID(preallocateReply))
	if callbackErr != nil {
		return callbackErr
	}

	renamed := fskit.NewFileNameWithString("renamed.txt")
	renameReply := objc.NewBlock(func(_ objc.Block, actualName objc.ID, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("rename returned error %d", errID)
			return
		}
		if fskit.FSFileNameFromID(actualName).String() != "renamed.txt" {
			callbackErr = fmt.Errorf("rename returned unexpected name")
		}
	})
	defer renameReply.Release()
	objc.Send[struct{}](volume, objc.Sel("renameItem:inDirectory:named:toNewName:inDirectory:overItem:replyHandler:"), created, root, createName.GetID(), renamed.GetID(), root, objc.ID(0), objc.ID(renameReply))
	if callbackErr != nil {
		return callbackErr
	}

	removeReply := objc.NewBlock(func(_ objc.Block, errID objc.ID) {
		if errID != 0 {
			callbackErr = fmt.Errorf("remove returned error %d", errID)
		}
	})
	defer removeReply.Release()
	objc.Send[struct{}](volume, objc.Sel("removeItem:named:fromDirectory:replyHandler:"), created, renamed.GetID(), root, objc.ID(removeReply))
	if callbackErr != nil {
		return callbackErr
	}

	fmt.Println("9pfs: fskit smoke ok")
	fmt.Println("9pfs: mapped 9p lookup, stat, symlink, hardlink, create, read, write, xattr, setattr, mtime, rename, and remove through FSKit callbacks")
	return nil
}
