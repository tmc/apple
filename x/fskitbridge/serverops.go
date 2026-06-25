package fskitbridge

import (
	"io"
	"math/bits"
	"syscall"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/fskit"
	"github.com/tmc/apple/objc"
)

func (s *Server) volumeMethods() []objc.MethodDef {
	return []objc.MethodDef{
		{Cmd: objc.Sel("activateWithOptions:replyHandler:"), Fn: s.impVolumeActivate},
		{Cmd: objc.Sel("deactivateWithOptions:replyHandler:"), Fn: s.impVolumeDeactivate},
		{Cmd: objc.Sel("mountWithOptions:replyHandler:"), Fn: s.impVolumeMount},
		{Cmd: objc.Sel("unmountWithReplyHandler:"), Fn: s.impVolumeUnmount},
		{Cmd: objc.Sel("getAttributes:ofItem:replyHandler:"), Fn: s.impVolumeGetAttributes},
		{Cmd: objc.Sel("setAttributes:onItem:replyHandler:"), Fn: s.impVolumeSetAttributes},
		{Cmd: objc.Sel("lookupItemNamed:inDirectory:replyHandler:"), Fn: s.impVolumeLookup},
		{Cmd: objc.Sel("lookupItemNamed:inDirectory:packer:replyHandler:"), Fn: s.impVolumeLookupWithPacker},
		{Cmd: objc.Sel("reclaimItem:replyHandler:"), Fn: s.impVolumeReclaim},
		{Cmd: objc.Sel("enumerateDirectory:startingAtCookie:verifier:providingAttributes:usingPacker:replyHandler:"), Fn: s.impVolumeEnumerate},
		{Cmd: objc.Sel("readSymbolicLink:replyHandler:"), Fn: s.impVolumeReadSymbolicLink},
		{Cmd: objc.Sel("createItemNamed:type:inDirectory:attributes:replyHandler:"), Fn: s.impVolumeCreateItem},
		{Cmd: objc.Sel("createSymbolicLinkNamed:inDirectory:attributes:linkContents:replyHandler:"), Fn: s.impVolumeCreateSymbolicLink},
		{Cmd: objc.Sel("createLinkToItem:named:inDirectory:replyHandler:"), Fn: s.impVolumeCreateLink},
		{Cmd: objc.Sel("removeItem:named:fromDirectory:replyHandler:"), Fn: s.impVolumeRemoveItem},
		{Cmd: objc.Sel("renameItem:inDirectory:named:toNewName:inDirectory:overItem:replyHandler:"), Fn: s.impVolumeRenameItem},
		{Cmd: objc.Sel("readFromFile:offset:length:intoBuffer:replyHandler:"), Fn: s.impVolumeRead},
		{Cmd: objc.Sel("writeContents:toFile:atOffset:replyHandler:"), Fn: s.impVolumeWrite},
		{Cmd: objc.Sel("openItem:withModes:replyHandler:"), Fn: s.impVolumeOpenItem},
		{Cmd: objc.Sel("closeItem:keepingModes:replyHandler:"), Fn: s.impVolumeCloseItem},
		{Cmd: objc.Sel("isOpenCloseInhibited"), Fn: s.impVolumeIsOpenCloseInhibited},
		{Cmd: objc.Sel("setOpenCloseInhibited:"), Fn: s.impVolumeSetBoolNoop},
		{Cmd: objc.Sel("checkAccessToItem:requestedAccess:replyHandler:"), Fn: s.impVolumeCheckAccess},
		{Cmd: objc.Sel("isAccessCheckInhibited"), Fn: s.impVolumeIsAccessCheckInhibited},
		{Cmd: objc.Sel("setAccessCheckInhibited:"), Fn: s.impVolumeSetBoolNoop},
		{Cmd: objc.Sel("preallocateSpaceForItem:atOffset:length:flags:replyHandler:"), Fn: s.impVolumePreallocate},
		{Cmd: objc.Sel("isPreallocateInhibited"), Fn: s.impVolumeIsPreallocateInhibited},
		{Cmd: objc.Sel("setPreallocateInhibited:"), Fn: s.impVolumeSetBoolNoop},
		{Cmd: objc.Sel("getXattrNamed:ofItem:replyHandler:"), Fn: s.impVolumeGetXattr},
		{Cmd: objc.Sel("setXattrNamed:toData:onItem:policy:replyHandler:"), Fn: s.impVolumeSetXattr},
		{Cmd: objc.Sel("listXattrsOfItem:replyHandler:"), Fn: s.impVolumeListXattrs},
		{Cmd: objc.Sel("supportedXattrNamesForItem:"), Fn: s.impVolumeSupportedXattrNames},
		{Cmd: objc.Sel("xattrOperationsInhibited"), Fn: s.impVolumeXattrOperationsInhibited},
		{Cmd: objc.Sel("setXattrOperationsInhibited:"), Fn: s.impVolumeSetBoolNoop},
		{Cmd: objc.Sel("setVolumeName:replyHandler:"), Fn: s.impVolumeSetVolumeName},
		{Cmd: objc.Sel("isVolumeRenameInhibited"), Fn: s.impVolumeIsVolumeRenameInhibited},
		{Cmd: objc.Sel("setVolumeRenameInhibited:"), Fn: s.impVolumeSetBoolNoop},
		{Cmd: objc.Sel("synchronizeWithFlags:replyHandler:"), Fn: s.impVolumeSynchronize},
		{Cmd: objc.Sel("supportedVolumeCapabilities"), Fn: s.impVolumeSupportedCapabilities},
		{Cmd: objc.Sel("requestedMountOptions"), Fn: s.impVolumeRequestedMountOptions},
		{Cmd: objc.Sel("setRequestedMountOptions:"), Fn: s.impVolumeSetRequestedMountOptions},
		{Cmd: objc.Sel("volumeStatistics"), Fn: s.impVolumeStatistics},
		{Cmd: objc.Sel("enableOpenUnlinkEmulation"), Fn: s.impVolumeEnableOpenUnlinkEmulation},
		{Cmd: objc.Sel("setEnableOpenUnlinkEmulation:"), Fn: s.impVolumeSetBoolNoop},
		{Cmd: objc.Sel("maximumLinkCount"), Fn: s.impVolumeMaximumLinkCount},
		{Cmd: objc.Sel("maximumNameLength"), Fn: s.impVolumeMaximumNameLength},
		{Cmd: objc.Sel("maximumFileSize"), Fn: s.impVolumeMaximumFileSize},
		{Cmd: objc.Sel("maximumFileSizeInBits"), Fn: s.impVolumeMaximumFileSizeInBits},
		{Cmd: objc.Sel("maximumXattrSize"), Fn: s.impVolumeMaximumXattrSize},
		{Cmd: objc.Sel("maximumXattrSizeInBits"), Fn: s.impVolumeMaximumXattrSizeInBits},
		{Cmd: objc.Sel("restrictsOwnershipChanges"), Fn: s.impVolumeRestrictsOwnershipChanges},
		{Cmd: objc.Sel("truncatesLongNames"), Fn: s.impVolumeTruncatesLongNames},
		{Cmd: objc.Sel("dealloc"), Fn: s.impVolumeDealloc},
	}
}

func (s *Server) impVolumeActivate(self objc.ID, _ objc.SEL, options, reply objc.ID) {
	volume, ok := s.volumeFor(self)
	if !ok {
		s.replyErr("activate", s.reply.ObjectError(reply, 0, POSIXError(syscall.EINVAL)))
		return
	}
	s.replyErr("activate", s.reply.ObjectError(reply, volume.root, 0))
}

func (s *Server) impVolumeDeactivate(self objc.ID, _ objc.SEL, options, reply objc.ID) {
	s.replyErr("deactivate", s.reply.Error(reply, 0))
}

func (s *Server) impVolumeMount(self objc.ID, _ objc.SEL, options, reply objc.ID) {
	s.replyErr("mount", s.reply.Error(reply, 0))
}

func (s *Server) impVolumeUnmount(self objc.ID, _ objc.SEL, reply objc.ID) {
	s.replyErr("unmount", s.reply.Void(reply))
}

func (s *Server) impVolumeGetAttributes(self objc.ID, _ objc.SEL, desired, item, reply objc.ID) {
	volume, ok := s.volumeFor(self)
	it, itemOK := s.itemFor(item)
	if !ok || !itemOK {
		s.replyErr("getAttributes", s.reply.ObjectError(reply, 0, POSIXError(syscall.EINVAL)))
		return
	}
	attrs, err := volume.impl.Attributes(it)
	if err != nil {
		s.replyErr("getAttributes", s.reply.ObjectError(reply, 0, s.errorFor(err)))
		return
	}
	s.replyErr("getAttributes", s.reply.ObjectError(reply, attrs.GetID(), 0))
}

func (s *Server) impVolumeSetAttributes(self objc.ID, _ objc.SEL, attrs, item, reply objc.ID) {
	volume, ok := s.volumeFor(self)
	it, itemOK := s.itemFor(item)
	if !ok || !itemOK {
		s.replyErr("setAttributes", s.reply.ObjectError(reply, 0, POSIXError(syscall.EINVAL)))
		return
	}
	mutable, ok := volume.impl.(MutableVolume)
	if !ok {
		s.replyErr("setAttributes", s.reply.ObjectError(reply, 0, POSIXError(syscall.EROFS)))
		return
	}
	request := fskit.FSItemSetAttributesRequestFromID(attrs)
	var set SetAttributes
	if request.IsValid(fskit.FSItemAttributeMode) {
		mode := request.Mode()
		set.Mode = &mode
		consume(request, fskit.FSItemAttributeMode)
	}
	if request.IsValid(fskit.FSItemAttributeSize) {
		size := request.Size()
		set.Size = &size
		consume(request, fskit.FSItemAttributeSize)
	}
	if request.IsValid(fskit.FSItemAttributeAccessTime) {
		ts := request.AccessTime()
		set.AccessTime = &ts
		consume(request, fskit.FSItemAttributeAccessTime)
	}
	if request.IsValid(fskit.FSItemAttributeModifyTime) {
		ts := request.ModifyTime()
		set.ModifyTime = &ts
		consume(request, fskit.FSItemAttributeModifyTime)
	}
	if err := mutable.SetAttributes(it, set); err != nil {
		s.replyErr("setAttributes", s.reply.ObjectError(reply, 0, s.errorFor(err)))
		return
	}
	updated, err := volume.impl.Attributes(it)
	if err != nil {
		s.replyErr("setAttributes", s.reply.ObjectError(reply, 0, s.errorFor(err)))
		return
	}
	s.replyErr("setAttributes", s.reply.ObjectError(reply, updated.GetID(), 0))
}

func consume(request fskit.FSItemSetAttributesRequest, attr fskit.FSItemAttribute) {
	request.SetConsumedAttributes(request.ConsumedAttributes() | attr)
}

func (s *Server) impVolumeLookup(self objc.ID, _ objc.SEL, name, directory, reply objc.ID) {
	volume, ok := s.volumeFor(self)
	dir, dirOK := s.itemFor(directory)
	if !ok || !dirOK {
		s.replyErr("lookup", s.reply.ItemNameError(reply, 0, 0, POSIXError(syscall.EINVAL)))
		return
	}
	child, err := volume.impl.Lookup(dir, fileNameString(name))
	if err != nil {
		s.replyErr("lookup", s.reply.ItemNameError(reply, 0, 0, s.errorFor(err)))
		return
	}
	s.replyErr("lookup", s.reply.ItemNameError(reply, s.newItem(volume.impl, child), name, 0))
}

func (s *Server) impVolumeLookupWithPacker(self objc.ID, _ objc.SEL, name, directory, packer, reply objc.ID) {
	s.impVolumeLookup(self, 0, name, directory, reply)
}

func (s *Server) impVolumeReclaim(self objc.ID, _ objc.SEL, item, reply objc.ID) {
	if volume, ok := s.volumeFor(self); ok {
		if it, ok := s.itemFor(item); ok {
			volume.impl.Reclaim(it)
		}
	}
	s.replyErr("reclaim", s.reply.Error(reply, 0))
}

func (s *Server) impVolumeEnumerate(self objc.ID, _ objc.SEL, directory objc.ID, cookie, verifier uint64, attrs, packer, reply objc.ID) {
	volume, ok := s.volumeFor(self)
	dir, dirOK := s.itemFor(directory)
	if !ok || !dirOK {
		s.replyErr("enumerate", s.reply.VerifierError(reply, verifier, POSIXError(syscall.EINVAL)))
		return
	}
	entries, err := volume.impl.ReadDir(dir)
	if err != nil {
		s.replyErr("enumerate", s.reply.VerifierError(reply, verifier, s.errorFor(err)))
		return
	}
	p := fskit.FSDirectoryEntryPackerFromID(packer)
	for i := int(cookie); i < len(entries); i++ {
		entry := entries[i]
		name := fskit.NewFileNameWithString(entry.Name)
		if !p.PackEntryWithNameItemTypeItemIDNextCookieAttributes(name, entry.Type, entry.Attributes.FileID(), uint64(i+1), entry.Attributes) {
			break
		}
	}
	s.replyErr("enumerate", s.reply.VerifierError(reply, verifier, 0))
}

func (s *Server) impVolumeReadSymbolicLink(self objc.ID, _ objc.SEL, item, reply objc.ID) {
	volume, ok := s.volumeFor(self)
	it, itemOK := s.itemFor(item)
	if !ok || !itemOK {
		s.replyErr("readlink", s.reply.ObjectError(reply, 0, POSIXError(syscall.EINVAL)))
		return
	}
	symlinks, ok := volume.impl.(SymlinkVolume)
	if !ok {
		s.replyErr("readlink", s.reply.ObjectError(reply, 0, POSIXError(syscall.ENOTSUP)))
		return
	}
	target, err := symlinks.Readlink(it)
	if err != nil {
		s.replyErr("readlink", s.reply.ObjectError(reply, 0, s.errorFor(err)))
		return
	}
	s.replyErr("readlink", s.reply.ObjectError(reply, fskit.NewFileNameWithString(target).GetID(), 0))
}

func (s *Server) impVolumeCreateItem(self objc.ID, _ objc.SEL, name objc.ID, typ fskit.FSItemType, directory, attrs, reply objc.ID) {
	volume, ok := s.volumeFor(self)
	dir, dirOK := s.itemFor(directory)
	if !ok || !dirOK {
		s.replyErr("create", s.reply.ItemNameError(reply, 0, 0, POSIXError(syscall.EINVAL)))
		return
	}
	mutable, ok := volume.impl.(MutableVolume)
	if !ok {
		s.replyErr("create", s.reply.ItemNameError(reply, 0, 0, POSIXError(syscall.EROFS)))
		return
	}
	mode := uint32(0666)
	if attrs != 0 {
		request := fskit.FSItemSetAttributesRequestFromID(attrs)
		if request.IsValid(fskit.FSItemAttributeMode) {
			mode = request.Mode()
			consume(request, fskit.FSItemAttributeMode)
		}
	}
	child, err := mutable.Create(dir, fileNameString(name), typ, mode)
	if err != nil {
		s.replyErr("create", s.reply.ItemNameError(reply, 0, 0, s.errorFor(err)))
		return
	}
	s.replyErr("create", s.reply.ItemNameError(reply, s.newItem(volume.impl, child), name, 0))
}

func (s *Server) impVolumeCreateSymbolicLink(self objc.ID, _ objc.SEL, name, directory, attrs, contents, reply objc.ID) {
	volume, ok := s.volumeFor(self)
	dir, dirOK := s.itemFor(directory)
	if !ok || !dirOK {
		s.replyErr("symlink", s.reply.ItemNameError(reply, 0, 0, POSIXError(syscall.EINVAL)))
		return
	}
	symlinks, ok := volume.impl.(SymlinkVolume)
	if !ok {
		s.replyErr("symlink", s.reply.ItemNameError(reply, 0, 0, POSIXError(syscall.ENOTSUP)))
		return
	}
	child, err := symlinks.Symlink(dir, fileNameString(name), fileNameString(contents))
	if err != nil {
		s.replyErr("symlink", s.reply.ItemNameError(reply, 0, 0, s.errorFor(err)))
		return
	}
	s.replyErr("symlink", s.reply.ItemNameError(reply, s.newItem(volume.impl, child), name, 0))
}

func (s *Server) impVolumeCreateLink(self objc.ID, _ objc.SEL, item, name, directory, reply objc.ID) {
	volume, ok := s.volumeFor(self)
	it, itemOK := s.itemFor(item)
	dir, dirOK := s.itemFor(directory)
	if !ok || !itemOK || !dirOK {
		s.replyErr("link", s.reply.ObjectError(reply, 0, POSIXError(syscall.EINVAL)))
		return
	}
	links, ok := volume.impl.(LinkVolume)
	if !ok {
		s.replyErr("link", s.reply.ObjectError(reply, 0, POSIXError(syscall.ENOTSUP)))
		return
	}
	if err := links.Link(it, dir, fileNameString(name)); err != nil {
		s.replyErr("link", s.reply.ObjectError(reply, 0, s.errorFor(err)))
		return
	}
	s.replyErr("link", s.reply.ObjectError(reply, name, 0))
}

func (s *Server) impVolumeRemoveItem(self objc.ID, _ objc.SEL, item, name, directory, reply objc.ID) {
	volume, ok := s.volumeFor(self)
	it, itemOK := s.itemFor(item)
	dir, dirOK := s.itemFor(directory)
	if !ok || !itemOK || !dirOK {
		s.replyErr("remove", s.reply.Error(reply, POSIXError(syscall.EINVAL)))
		return
	}
	mutable, ok := volume.impl.(MutableVolume)
	if !ok {
		s.replyErr("remove", s.reply.Error(reply, POSIXError(syscall.EROFS)))
		return
	}
	if err := mutable.Remove(dir, fileNameString(name), it); err != nil {
		s.replyErr("remove", s.reply.Error(reply, s.errorFor(err)))
		return
	}
	s.replyErr("remove", s.reply.Error(reply, 0))
}

func (s *Server) impVolumeRenameItem(self objc.ID, _ objc.SEL, item, sourceDirectory, sourceName, destinationName, destinationDirectory, overItem, reply objc.ID) {
	volume, ok := s.volumeFor(self)
	it, itemOK := s.itemFor(item)
	srcDir, srcOK := s.itemFor(sourceDirectory)
	dstDir, dstOK := s.itemFor(destinationDirectory)
	if !ok || !itemOK || !srcOK || !dstOK {
		s.replyErr("rename", s.reply.ObjectError(reply, 0, POSIXError(syscall.EINVAL)))
		return
	}
	mutable, ok := volume.impl.(MutableVolume)
	if !ok {
		s.replyErr("rename", s.reply.ObjectError(reply, 0, POSIXError(syscall.EROFS)))
		return
	}
	var over Item
	if overItem != 0 {
		over, _ = s.itemFor(overItem)
	}
	err := mutable.Rename(it, srcDir, fileNameString(sourceName), dstDir, fileNameString(destinationName), over)
	if err != nil {
		s.replyErr("rename", s.reply.ObjectError(reply, 0, s.errorFor(err)))
		return
	}
	if over != nil {
		volume.impl.Reclaim(over)
		s.items.Delete(overItem)
	}
	s.replyErr("rename", s.reply.ObjectError(reply, destinationName, 0))
}

func (s *Server) impVolumeRead(self objc.ID, _ objc.SEL, item objc.ID, offset int64, length uintptr, buffer, reply objc.ID) {
	volume, ok := s.volumeFor(self)
	it, itemOK := s.itemFor(item)
	if !ok || !itemOK {
		s.replyErr("read", s.reply.SizeError(reply, 0, POSIXError(syscall.EINVAL)))
		return
	}
	var buf []byte
	if length > 0 {
		dst := objc.Send[unsafe.Pointer](buffer, objc.Sel("mutableBytes"))
		if dst == nil {
			s.replyErr("read", s.reply.SizeError(reply, 0, POSIXError(syscall.EINVAL)))
			return
		}
		buf = unsafe.Slice((*byte)(dst), length)
	}
	n, err := volume.impl.Read(it, offset, buf)
	if err != nil && err != io.EOF {
		s.replyErr("read", s.reply.SizeError(reply, uintptr(n), s.errorFor(err)))
		return
	}
	s.replyErr("read", s.reply.SizeError(reply, uintptr(n), 0))
}

func (s *Server) impVolumeWrite(self objc.ID, _ objc.SEL, contents, item objc.ID, offset int64, reply objc.ID) {
	volume, ok := s.volumeFor(self)
	it, itemOK := s.itemFor(item)
	if !ok || !itemOK {
		s.replyErr("write", s.reply.SizeError(reply, 0, POSIXError(syscall.EINVAL)))
		return
	}
	mutable, ok := volume.impl.(MutableVolume)
	if !ok {
		s.replyErr("write", s.reply.SizeError(reply, 0, POSIXError(syscall.EROFS)))
		return
	}
	n, err := mutable.Write(it, offset, bytesFromNSData(contents))
	s.replyErr("write", s.reply.SizeError(reply, uintptr(n), s.errorFor(err)))
}

func (s *Server) impVolumeOpenItem(self objc.ID, _ objc.SEL, item objc.ID, modes fskit.FSVolumeOpenModes, reply objc.ID) {
	volume, ok := s.volumeFor(self)
	it, itemOK := s.itemFor(item)
	if !ok || !itemOK {
		s.replyErr("open", s.reply.Error(reply, POSIXError(syscall.EINVAL)))
		return
	}
	if openClose, ok := volume.impl.(OpenCloseVolume); ok {
		if err := openClose.Open(it, modes); err != nil {
			s.replyErr("open", s.reply.Error(reply, s.errorFor(err)))
			return
		}
	}
	s.replyErr("open", s.reply.Error(reply, 0))
}

func (s *Server) impVolumeCloseItem(self objc.ID, _ objc.SEL, item objc.ID, modes fskit.FSVolumeOpenModes, reply objc.ID) {
	volume, ok := s.volumeFor(self)
	it, itemOK := s.itemFor(item)
	if !ok || !itemOK {
		s.replyErr("close", s.reply.Error(reply, POSIXError(syscall.EINVAL)))
		return
	}
	if openClose, ok := volume.impl.(OpenCloseVolume); ok {
		if err := openClose.Close(it, modes); err != nil {
			s.replyErr("close", s.reply.Error(reply, s.errorFor(err)))
			return
		}
	}
	s.replyErr("close", s.reply.Error(reply, 0))
}

func (s *Server) impVolumeIsOpenCloseInhibited(self objc.ID, _ objc.SEL) bool {
	volume, ok := s.volumeFor(self)
	if !ok {
		return true
	}
	_, openClose := volume.impl.(OpenCloseVolume)
	return !openClose
}

func (s *Server) impVolumeSetBoolNoop(self objc.ID, _ objc.SEL, value bool) {}

func (s *Server) impVolumeCheckAccess(self objc.ID, _ objc.SEL, item objc.ID, access fskit.FSAccessMask, reply objc.ID) {
	volume, ok := s.volumeFor(self)
	it, itemOK := s.itemFor(item)
	if !ok || !itemOK {
		s.replyErr("checkAccess", s.reply.BoolError(reply, false, POSIXError(syscall.EINVAL)))
		return
	}
	if checks, ok := volume.impl.(AccessCheckVolume); ok {
		allowed, err := checks.CheckAccess(it, access)
		if err != nil {
			s.replyErr("checkAccess", s.reply.BoolError(reply, false, s.errorFor(err)))
			return
		}
		s.replyErr("checkAccess", s.reply.BoolError(reply, allowed, 0))
		return
	}
	s.replyErr("checkAccess", s.reply.BoolError(reply, true, 0))
}

func (s *Server) impVolumeIsAccessCheckInhibited(self objc.ID, _ objc.SEL) bool {
	volume, ok := s.volumeFor(self)
	if !ok {
		return true
	}
	_, checks := volume.impl.(AccessCheckVolume)
	return !checks
}

func (s *Server) impVolumePreallocate(self objc.ID, _ objc.SEL, item objc.ID, offset int64, length uintptr, flags fskit.FSPreallocateFlags, reply objc.ID) {
	volume, ok := s.volumeFor(self)
	it, itemOK := s.itemFor(item)
	if !ok || !itemOK {
		s.replyErr("preallocate", s.reply.SizeError(reply, 0, POSIXError(syscall.EINVAL)))
		return
	}
	prealloc, ok := volume.impl.(PreallocateVolume)
	if !ok {
		s.replyErr("preallocate", s.reply.SizeError(reply, 0, POSIXError(syscall.ENOTSUP)))
		return
	}
	n, err := prealloc.Preallocate(it, offset, uint64(length))
	if err != nil {
		s.replyErr("preallocate", s.reply.SizeError(reply, 0, s.errorFor(err)))
		return
	}
	s.replyErr("preallocate", s.reply.SizeError(reply, uintptr(n), 0))
}

func (s *Server) impVolumeIsPreallocateInhibited(self objc.ID, _ objc.SEL) bool {
	volume, ok := s.volumeFor(self)
	if !ok {
		return true
	}
	_, prealloc := volume.impl.(PreallocateVolume)
	return !prealloc
}

func (s *Server) impVolumeGetXattr(self objc.ID, _ objc.SEL, name, item, reply objc.ID) {
	volume, ok := s.volumeFor(self)
	it, itemOK := s.itemFor(item)
	if !ok || !itemOK {
		s.replyErr("getXattr", s.reply.ObjectError(reply, 0, POSIXError(syscall.EINVAL)))
		return
	}
	xattrs, ok := volume.impl.(XattrVolume)
	if !ok {
		s.replyErr("getXattr", s.reply.ObjectError(reply, 0, POSIXError(syscall.ENOTSUP)))
		return
	}
	data, err := xattrs.GetXattr(it, fileNameString(name))
	if err != nil {
		s.replyErr("getXattr", s.reply.ObjectError(reply, 0, s.xattrErrorFor(err)))
		return
	}
	s.replyErr("getXattr", s.reply.ObjectError(reply, foundation.NewDataWithBytesLength(data).GetID(), 0))
}

func (s *Server) impVolumeSetXattr(self objc.ID, _ objc.SEL, name, value, item objc.ID, policy fskit.FSSetXattrPolicy, reply objc.ID) {
	volume, ok := s.volumeFor(self)
	it, itemOK := s.itemFor(item)
	if !ok || !itemOK {
		s.replyErr("setXattr", s.reply.Error(reply, POSIXError(syscall.EINVAL)))
		return
	}
	xattrs, ok := volume.impl.(XattrVolume)
	if !ok {
		s.replyErr("setXattr", s.reply.Error(reply, POSIXError(syscall.ENOTSUP)))
		return
	}
	attr := fileNameString(name)
	var errno syscall.Errno
	switch policy {
	case fskit.FSSetXattrPolicyDelete:
		errno = xattrErrnoOf(xattrs.RemoveXattr(it, attr))
	case fskit.FSSetXattrPolicyAlwaysSet:
		if value == 0 {
			errno = syscall.EINVAL
			break
		}
		errno = xattrErrnoOf(xattrs.SetXattr(it, attr, bytesFromNSData(value)))
	case fskit.FSSetXattrPolicyMustCreate:
		if value == 0 {
			errno = syscall.EINVAL
			break
		}
		if _, err := xattrs.GetXattr(it, attr); err == nil {
			errno = syscall.EEXIST
			break
		}
		errno = xattrErrnoOf(xattrs.SetXattr(it, attr, bytesFromNSData(value)))
	case fskit.FSSetXattrPolicyMustReplace:
		if value == 0 {
			errno = syscall.EINVAL
			break
		}
		if _, err := xattrs.GetXattr(it, attr); err != nil {
			errno = ENOATTR
			break
		}
		errno = xattrErrnoOf(xattrs.SetXattr(it, attr, bytesFromNSData(value)))
	default:
		errno = syscall.ENOTSUP
	}
	if errno != 0 {
		s.replyErr("setXattr", s.reply.Error(reply, POSIXError(errno)))
		return
	}
	s.replyErr("setXattr", s.reply.Error(reply, 0))
}

// xattrErrnoOf maps an xattr operation result to an errno, with 0 for nil.
func xattrErrnoOf(err error) syscall.Errno {
	if err == nil {
		return 0
	}
	return xattrErrnoFor(err)
}

func (s *Server) impVolumeListXattrs(self objc.ID, _ objc.SEL, item, reply objc.ID) {
	array, errID := s.xattrNameArray(self, item)
	if errID != 0 {
		s.replyErr("listXattrs", s.reply.ObjectError(reply, 0, errID))
		return
	}
	s.replyErr("listXattrs", s.reply.ObjectError(reply, array, 0))
}

func (s *Server) impVolumeSupportedXattrNames(self objc.ID, _ objc.SEL, item objc.ID) objc.ID {
	// supportedXattrNamesForItem: returns an array with no error channel;
	// on failure report no names, but log it rather than discard silently.
	array, errID := s.xattrNameArray(self, item)
	if errID != 0 {
		s.logf("supported xattr names failed")
		return 0
	}
	return array
}

func (s *Server) xattrNameArray(self, item objc.ID) (objc.ID, objc.ID) {
	volume, ok := s.volumeFor(self)
	it, itemOK := s.itemFor(item)
	if !ok || !itemOK {
		return 0, POSIXError(syscall.EINVAL)
	}
	xattrs, ok := volume.impl.(XattrVolume)
	if !ok {
		return 0, POSIXError(syscall.ENOTSUP)
	}
	names, err := xattrs.ListXattr(it)
	if err != nil {
		return 0, s.xattrErrorFor(err)
	}
	array := foundation.NewNSMutableArray()
	for _, name := range names {
		array.AddObject(fskit.NewFileNameWithString(name))
	}
	return array.GetID(), 0
}

func (s *Server) impVolumeXattrOperationsInhibited(self objc.ID, _ objc.SEL) bool {
	volume, ok := s.volumeFor(self)
	if !ok {
		return true
	}
	_, xattrs := volume.impl.(XattrVolume)
	return !xattrs
}

func (s *Server) impVolumeSetVolumeName(self objc.ID, _ objc.SEL, name, reply objc.ID) {
	s.replyErr("setVolumeName", s.reply.ObjectError(reply, 0, POSIXError(syscall.ENOTSUP)))
}

func (s *Server) impVolumeIsVolumeRenameInhibited(self objc.ID, _ objc.SEL) bool { return true }

func (s *Server) impVolumeSynchronize(self objc.ID, _ objc.SEL, flags uint64, reply objc.ID) {
	s.replyErr("synchronize", s.reply.Error(reply, 0))
}

func (s *Server) impVolumeSupportedCapabilities(self objc.ID, _ objc.SEL) objc.ID {
	volume, ok := s.volumeFor(self)
	if !ok {
		return 0
	}
	if caps, ok := volume.impl.(CapabilitiesVolume); ok {
		return caps.SupportedCapabilities().GetID()
	}
	capabilities := fskit.NewFSVolumeSupportedCapabilities()
	capabilities.SetSupports64BitObjectIDs(true)
	capabilities.SetSupportsHiddenFiles(true)
	if _, ok := volume.impl.(SymlinkVolume); ok {
		capabilities.SetSupportsSymbolicLinks(true)
	}
	if _, ok := volume.impl.(LinkVolume); ok {
		capabilities.SetSupportsHardLinks(true)
	}
	capabilities.SetCaseFormat(fskit.FSVolumeCaseFormatSensitive)
	return capabilities.GetID()
}

func (s *Server) impVolumeRequestedMountOptions(self objc.ID, _ objc.SEL) fskit.FSMountOptions {
	return 0
}

func (s *Server) impVolumeSetRequestedMountOptions(self objc.ID, _ objc.SEL, value fskit.FSMountOptions) {
}

func (s *Server) impVolumeStatistics(self objc.ID, _ objc.SEL) objc.ID {
	volume, ok := s.volumeFor(self)
	if !ok {
		return 0
	}
	if stats, ok := volume.impl.(StatisticsVolume); ok {
		return stats.Statistics().GetID()
	}
	result := fskit.NewStatFSResultWithFileSystemTypeName(volume.impl.VolumeName())
	result.SetBlockSize(4096)
	result.SetIoSize(4096)
	result.SetTotalBlocks(1)
	result.SetAvailableBlocks(0)
	result.SetFreeBlocks(0)
	result.SetUsedBlocks(1)
	result.SetTotalFiles(1)
	result.SetFreeFiles(0)
	return result.GetID()
}

func (s *Server) pathConf(self objc.ID) PathConf {
	if volume, ok := s.volumeFor(self); ok {
		if conf, ok := volume.impl.(PathConfVolume); ok {
			return conf.PathConf()
		}
	}
	return DefaultPathConf()
}

func (s *Server) impVolumeEnableOpenUnlinkEmulation(self objc.ID, _ objc.SEL) bool {
	return s.pathConf(self).OpenUnlinkEmulation
}

func (s *Server) impVolumeMaximumLinkCount(self objc.ID, _ objc.SEL) int {
	return s.pathConf(self).MaximumLinkCount
}

func (s *Server) impVolumeMaximumNameLength(self objc.ID, _ objc.SEL) int {
	return s.pathConf(self).MaximumNameLength
}

func (s *Server) impVolumeMaximumFileSize(self objc.ID, _ objc.SEL) uint64 {
	return s.pathConf(self).MaximumFileSize
}

func (s *Server) impVolumeMaximumFileSizeInBits(self objc.ID, _ objc.SEL) int {
	return bits.Len64(s.pathConf(self).MaximumFileSize)
}

func (s *Server) impVolumeMaximumXattrSize(self objc.ID, _ objc.SEL) int {
	return s.pathConf(self).MaximumXattrSize
}

func (s *Server) impVolumeMaximumXattrSizeInBits(self objc.ID, _ objc.SEL) int {
	return s.pathConf(self).MaximumXattrSize * 8
}

func (s *Server) impVolumeRestrictsOwnershipChanges(self objc.ID, _ objc.SEL) bool {
	return s.pathConf(self).RestrictsOwnershipChanges
}

func (s *Server) impVolumeTruncatesLongNames(self objc.ID, _ objc.SEL) bool {
	return s.pathConf(self).TruncatesLongNames
}

func (s *Server) impVolumeDealloc(self objc.ID, _ objc.SEL) {
	s.volumes.Delete(self)
}

func (s *Server) impItemDealloc(self objc.ID, _ objc.SEL) {
	if v, ok := s.items.LoadAndDelete(self); ok {
		entry := v.(*serverItem)
		entry.vol.Reclaim(entry.item)
	}
}
