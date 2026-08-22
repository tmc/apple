// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Methods that all volumes implement to provide required capabilities.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations
type FSVolumeOperations interface {
	objectivec.IObject

	// Activates the volume using the specified options.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/activate(options:replyHandler:)
	ActivateWithOptionsReplyHandler(options IFSTaskOptions, reply FSItemErrorHandler)

	// Tears down a previously initialized volume instance.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/deactivate(options:replyHandler:)
	DeactivateWithOptionsReplyHandler(options FSDeactivateOptions, reply ErrorHandler)

	// Mounts this volume, using the specified options.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/mount(options:replyHandler:)
	MountWithOptionsReplyHandler(options IFSTaskOptions, reply ErrorHandler)

	// Unmounts this volume.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/unmount(replyHandler:)
	UnmountWithReplyHandler(reply VoidHandler)

	// Creates a new file or directory item.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/createItem(named:type:inDirectory:attributes:replyHandler:)
	CreateItemNamedTypeInDirectoryAttributesReplyHandler(name IFSFileName, type_ FSItemType, directory IFSItem, newAttributes IFSItemSetAttributesRequest, reply FSItemFSFileNameErrorHandler)

	// Looks up an item within a directory.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/lookupItem(named:inDirectory:replyHandler:)
	LookupItemNamedInDirectoryReplyHandler(name IFSFileName, directory IFSItem, reply FSItemFSFileNameErrorHandler)

	// Removes an existing item from a given directory.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/removeItem(_:named:fromDirectory:replyHandler:)
	RemoveItemNamedFromDirectoryReplyHandler(item IFSItem, name IFSFileName, directory IFSItem, reply ErrorHandler)

	// Renames an item from one path in the file system to another.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/renameItem(_:inDirectory:named:to:inDirectory:overItem:replyHandler:)
	RenameItemInDirectoryNamedToNewNameInDirectoryOverItemReplyHandler(item IFSItem, sourceDirectory IFSItem, sourceName IFSFileName, destinationName IFSFileName, destinationDirectory IFSItem, overItem IFSItem, reply FSFileNameErrorHandler)

	// Reclaims an item, releasing any resources allocated for the item.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/reclaimItem(_:replyHandler:)
	ReclaimItemReplyHandler(item IFSItem, reply ErrorHandler)

	// Creates a new hard link.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/createLink(to:named:inDirectory:replyHandler:)
	CreateLinkToItemNamedInDirectoryReplyHandler(item IFSItem, name IFSFileName, directory IFSItem, reply FSFileNameErrorHandler)

	// Creates a new symbolic link.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/createSymbolicLink(named:inDirectory:attributes:linkContents:replyHandler:)
	CreateSymbolicLinkNamedInDirectoryAttributesLinkContentsReplyHandler(name IFSFileName, directory IFSItem, newAttributes IFSItemSetAttributesRequest, contents IFSFileName, reply FSItemFSFileNameErrorHandler)

	// Reads a symbolic link.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/readSymbolicLink(_:replyHandler:)
	ReadSymbolicLinkReplyHandler(item IFSItem, reply FSFileNameErrorHandler)

	// Fetches attributes for the given item.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/getAttributes(_:of:replyHandler:)
	GetAttributesOfItemReplyHandler(desiredAttributes IFSItemGetAttributesRequest, item IFSItem, reply FSItemAttributesErrorHandler)

	// Sets the given attributes on an item.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/setAttributes(_:on:replyHandler:)
	SetAttributesOnItemReplyHandler(newAttributes IFSItemSetAttributesRequest, item IFSItem, reply FSItemAttributesErrorHandler)

	// Enumerates the contents of the given directory.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/enumerateDirectory(_:startingAt:verifier:attributes:packer:replyHandler:)
	EnumerateDirectoryStartingAtCookieVerifierProvidingAttributesUsingPackerReplyHandler(directory IFSItem, cookie FSDirectoryCookie, verifier FSDirectoryVerifier, attributes IFSItemGetAttributesRequest, packer IFSDirectoryEntryPacker, reply FSDirectoryVerifierErrorHandler)

	// Synchronizes the volume with its underlying resource.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/synchronize(flags:replyHandler:)
	SynchronizeWithFlagsReplyHandler(flags FSSyncFlags, reply ErrorHandler)

	// A property that provides the supported capabilities of the volume.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/supportedVolumeCapabilities
	SupportedVolumeCapabilities() IFSVolumeSupportedCapabilities

	// A property that provides up-to-date statistics of the volume.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/volumeStatistics
	VolumeStatistics() IFSStatFSResult

	// A property that allows the file system to request for specific mount options from FSKit.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/requestedMountOptions
	RequestedMountOptions() FSMountOptions
	SetRequestedMountOptions(value FSMountOptions)

	// A property that allows the file system to use open-unlink emulation.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/enableOpenUnlinkEmulation
	EnableOpenUnlinkEmulation() bool
	SetEnableOpenUnlinkEmulation(value bool)
}

// FSVolumeOperationsObject wraps an existing Objective-C object that conforms to the FSVolumeOperations protocol.
type FSVolumeOperationsObject struct {
	objectivec.Object
}

func (o FSVolumeOperationsObject) BaseObject() objectivec.Object {
	return o.Object
}

// FSVolumeOperationsObjectFromID constructs a [FSVolumeOperationsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func FSVolumeOperationsObjectFromID(id objc.ID) FSVolumeOperationsObject {
	return FSVolumeOperationsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Activates the volume using the specified options.
//
// options: Options to apply to the activation. These can include security-scoped file
// paths. There are no defined options currently.
//
// reply: A block or closure to indicate success or failure. If activation succeeds,
// pass the root [FSItem] and a `nil` error. If activation fails, pass the
// relevant error as the second parameter; FSKit ignores any [FSItem] in this
// case. In Swift, `reply` takes only the [FSItem] as the parameter; you
// signal any error with a `throw`. For an `async` Swift implementation,
// there’s no reply handler; simply return the [FSItem] or throw an error.
//
// # Discussion
//
// When FSKit calls this method, allocate any in-memory state required to
// represent the file system. Also allocate an [FSItem] for the root directory
// of the file system, and pass it to the reply block. FSKit caches this root
// item for the lifetime of the volume, and uses it as a starting point for
// all file look-ups.
//
// Volume activation occurs prior to any call to mount the volume.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/activate(options:replyHandler:)
func (o FSVolumeOperationsObject) ActivateWithOptionsReplyHandler(options IFSTaskOptions, reply FSItemErrorHandler) {
	_block1, _cleanup1 := NewFSItemErrorBlock(reply)
	defer _cleanup1()
	objc.Send[struct{}](o.ID, objc.Sel("activateWithOptions:replyHandler:"), options, objc.ID(_block1))
}

// Tears down a previously initialized volume instance.
//
// options: Options to apply to the deactivation.
//
// reply: A block or closure to indicate success or failure. If activation fails,
// pass an error as the one parameter to the reply handler. If activation
// succeeds, pass `nil`. For an `async` Swift implementation, there’s no
// reply handler; simply throw an error or return normally.
//
// # Discussion
//
// Set up your implementation to release any resources allocated for the
// volume instance. By the time you receive this callback, FSKit has already
// performed a reclaim call to release all other file nodes associated with
// this file system instance.
//
// Avoid performing any I/O in this method. Prior to calling this method,
// FSKit has already issued a sync call to perform any cleanup-related I/O.
//
// FSKit unmounts any mounted volume with a call to [UnmountWithReplyHandler]
// prior to the deactivate callback.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/deactivate(options:replyHandler:)
func (o FSVolumeOperationsObject) DeactivateWithOptionsReplyHandler(options FSDeactivateOptions, reply ErrorHandler) {
	_block1, _cleanup1 := NewErrorBlock(reply)
	defer _cleanup1()
	objc.Send[struct{}](o.ID, objc.Sel("deactivateWithOptions:replyHandler:"), options, objc.ID(_block1))
}

// Mounts this volume, using the specified options.
//
// options: Options to apply to the mount. These can include security-scoped file
// paths. There are no defined options currently.
//
// reply: A block or closure to indicate success or failure. If mounting fails, pass
// an error as the one parameter to the reply handler. If mounting succeeds,
// pass `nil`. For an `async` Swift implementation, there’s no reply
// handler; simply return normally.
//
// # Discussion
//
// FSKit calls this method as a signal that some process is trying to mount
// this volume. Your file system receives a call to
// [ActivateWithOptionsReplyHandler] prior to receiving any mount calls.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/mount(options:replyHandler:)
func (o FSVolumeOperationsObject) MountWithOptionsReplyHandler(options IFSTaskOptions, reply ErrorHandler) {
	_block1, _cleanup1 := NewErrorBlock(reply)
	defer _cleanup1()
	objc.Send[struct{}](o.ID, objc.Sel("mountWithOptions:replyHandler:"), options, objc.ID(_block1))
}

// Unmounts this volume.
//
// reply: A block or closure to indicate success or failure. If unmounting fails,
// pass an error as the one parameter to the reply handler. If unmounting
// succeeds, pass `nil`. For an `async` Swift implementation, there’s no
// reply handler; simply return normally.
//
// # Discussion
//
// Clear and flush all cached state in your implementation of this method.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/unmount(replyHandler:)
func (o FSVolumeOperationsObject) UnmountWithReplyHandler(reply VoidHandler) {
	_block0, _cleanup0 := NewVoidBlock(reply)
	defer _cleanup0()
	objc.Send[struct{}](o.ID, objc.Sel("unmountWithReplyHandler:"), objc.ID(_block0))
}

// Creates a new file or directory item.
//
// name: The new item’s name.
//
// type: The new item’s type. Valid values are [FSItemTypeFile] or
// [FSItemTypeDirectory].
//
// directory: The directory in which to create the item.
//
// newAttributes: Attributes to apply to the new item.
//
// reply: A block or closure to indicate success or failure. If creation succeeds,
// pass the newly-created [FSItem] and its [FSFileName], along with a `nil`
// error. If creation fails, pass the relevant error as the third parameter;
// FSKit ignores any [FSItem] or [FSFileName] in this case. For an `async`
// Swift implementation, there’s no reply handler; simply return a tuple of
// the [FSItem] and its [FSFileName] or throw an error.
//
// # Discussion
//
// If an item named `name` already exists in the directory indicated by
// `directory`, complete the request with an error with a domain of
// [NSPOSIXErrorDomain] and a code of [EEXIST].
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/createItem(named:type:inDirectory:attributes:replyHandler:)
//
// [NSPOSIXErrorDomain]: https://developer.apple.com/documentation/Foundation/NSPOSIXErrorDomain
func (o FSVolumeOperationsObject) CreateItemNamedTypeInDirectoryAttributesReplyHandler(name IFSFileName, type_ FSItemType, directory IFSItem, newAttributes IFSItemSetAttributesRequest, reply FSItemFSFileNameErrorHandler) {
	_block4, _cleanup4 := NewFSItemFSFileNameErrorBlock(reply)
	defer _cleanup4()
	objc.Send[struct{}](o.ID, objc.Sel("createItemNamed:type:inDirectory:attributes:replyHandler:"), name, type_, directory, newAttributes, objc.ID(_block4))
}

// Looks up an item within a directory.
//
// name: The name of the item to look up.
//
// directory: The directory in which to look up the item.
//
// reply: A block or closure to indicate success or failure. If lookup succeeds, pass
// the found [FSItem] and its [FSFileName] (as saved within the file system),
// along with a `nil` error. If lookup fails, pass the relevant error as the
// third parameter; any [FSItem] or [FSFileName] are ignored in this case. For
// an `async` Swift implementation, there’s no reply handler; simply return
// the [FSItem] and [FSFileName] as a tuple or throw an error.
//
// # Discussion
//
// If no item matching `name` exists in the directory indicated by
// `directory`, complete the request with an error with a domain of
// [NSPOSIXErrorDomain] and a code of [ENOENT].
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/lookupItem(named:inDirectory:replyHandler:)
//
// [NSPOSIXErrorDomain]: https://developer.apple.com/documentation/Foundation/NSPOSIXErrorDomain
func (o FSVolumeOperationsObject) LookupItemNamedInDirectoryReplyHandler(name IFSFileName, directory IFSItem, reply FSItemFSFileNameErrorHandler) {
	_block2, _cleanup2 := NewFSItemFSFileNameErrorBlock(reply)
	defer _cleanup2()
	objc.Send[struct{}](o.ID, objc.Sel("lookupItemNamed:inDirectory:replyHandler:"), name, directory, objc.ID(_block2))
}

// Removes an existing item from a given directory.
//
// item: The item to remove.
//
// name: The name of the item to remove.
//
// directory: The directory from which to remove the item.
//
// reply: A block or closure to indicate success or failure. If removal fails, pass
// an error as the one parameter to the reply handler. If removal succeeds,
// pass `nil`. For an `async` Swift implementation, there’s no reply
// handler; simply throw an error or return normally.
//
// # Discussion
//
// Don’t actually remove the item object itself in your implementation;
// instead, only remove the given item name from the given directory. Remove
// and deallocate the item in [ReclaimItemReplyHandler].
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/removeItem(_:named:fromDirectory:replyHandler:)
func (o FSVolumeOperationsObject) RemoveItemNamedFromDirectoryReplyHandler(item IFSItem, name IFSFileName, directory IFSItem, reply ErrorHandler) {
	_block3, _cleanup3 := NewErrorBlock(reply)
	defer _cleanup3()
	objc.Send[struct{}](o.ID, objc.Sel("removeItem:named:fromDirectory:replyHandler:"), item, name, directory, objc.ID(_block3))
}

// Renames an item from one path in the file system to another.
//
// item: The file system object being renamed.
//
// sourceDirectory: The directory that currently contains the item to rename.
//
// sourceName: The name of the item within the source directory.
//
// destinationName: The new name of the item as it appears in `destinationDirectory`.
//
// destinationDirectory: The directory to contain the renamed object, which may be the same as
// `sourceDirectory`.
//
// overItem: The file system object if the destination exists, as discovered in a prior
// lookup. If this parameter is non-`nil`, mark `overItem` as deleted, so the
// file system can free its allocated space on the next call to
// [ReclaimItemReplyHandler]. After doing so, ensure the operation finishes
// without errors.
//
// reply: A block or closure to indicate success or failure. If renaming succeeds,
// pass the [FSFileName] as it exists within `destinationDirectory` and a
// `nil` error. If renaming fails, pass the relevant error as the second
// parameter; FSKit ignores any [FSFileName] in this case. For an `async`
// Swift implementation, there’s no reply handler; simply return the
// [FSFileName] or throw an error.
//
// # Discussion
//
// Implement renaming along the lines of this algorithm:
//
// - If `item` is a file: - - If the destination file exists: - - Remove the
// destination file. - If the source and destination directories are the same:
// - - Rewrite the name in the existing directory. - Else: - - Write the new
// entry in the destination directory. - Clear the old directory entry. - If
// `item` is a directory: - - If the destination directory exists: - - If the
// destination directory isn’t empty: - - Fail the operation with an error
// of [NSPOSIXErrorDomain] and a code of [ENOTEMPTY]. - Else: - - Remove the
// destination directory. - If the source and destination directories are the
// same: - - Rewrite the name in the existing directory. - Else: - - If the
// destination is a child of the source directory: - - Fail the operation with
// an error. - Else: - - Write the new entry in the destination directory. -
// Update `"."` and `".."` in the moved directory. - Clear the old directory
// entry.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/renameItem(_:inDirectory:named:to:inDirectory:overItem:replyHandler:)
//
// [NSPOSIXErrorDomain]: https://developer.apple.com/documentation/Foundation/NSPOSIXErrorDomain
func (o FSVolumeOperationsObject) RenameItemInDirectoryNamedToNewNameInDirectoryOverItemReplyHandler(item IFSItem, sourceDirectory IFSItem, sourceName IFSFileName, destinationName IFSFileName, destinationDirectory IFSItem, overItem IFSItem, reply FSFileNameErrorHandler) {
	_block6, _cleanup6 := NewFSFileNameErrorBlock(reply)
	defer _cleanup6()
	objc.Send[struct{}](o.ID, objc.Sel("renameItem:inDirectory:named:toNewName:inDirectory:overItem:replyHandler:"), item, sourceDirectory, sourceName, destinationName, destinationDirectory, overItem, objc.ID(_block6))
}

// Reclaims an item, releasing any resources allocated for the item.
//
// item: The item to reclaim.
//
// reply: A block or closure to indicate success or failure. If removal fails, pass
// an error as the one parameter to the reply handler. If removal succeeds,
// pass `nil`. For an `async` Swift implementation, there’s no reply
// handler; simply throw an error or return normally.
//
// # Discussion
//
// FSKit guarantees that for every [FSItem] returned by the volume, a
// corresponding reclaim operation occurs after the upper layers no longer
// reference that item.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/reclaimItem(_:replyHandler:)
func (o FSVolumeOperationsObject) ReclaimItemReplyHandler(item IFSItem, reply ErrorHandler) {
	_block1, _cleanup1 := NewErrorBlock(reply)
	defer _cleanup1()
	objc.Send[struct{}](o.ID, objc.Sel("reclaimItem:replyHandler:"), item, objc.ID(_block1))
}

// Creates a new hard link.
//
// item: The existing item to which to link.
//
// name: The name for the new link.
//
// directory: The directory in which to create the link.
//
// reply: A block or closure to indicate success or failure. If creation succeeds,
// pass an [FSFileName] of the newly-created link and a `nil` error. If
// creation fails, pass the relevant error as the second parameter; FSKit
// ignores any [FSFileName] in this case. For an `async` Swift implementation,
// there’s no reply handler; simply return the [FSFileName] or throw an
// error.
//
// # Discussion
//
// If creating the link fails, complete the request with an error with a
// domain of [NSPOSIXErrorDomain] and the following error codes:
//
// - [EEXIST] if there’s already an item named `name` in the directory. -
// [EMLINK] if creating the link would exceed the maximum number of hard links
// supported on `item`. - [ENOTSUP] if the file system doesn’t support
// creating hard links to the type of file system object that `item`
// represents.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/createLink(to:named:inDirectory:replyHandler:)
//
// [NSPOSIXErrorDomain]: https://developer.apple.com/documentation/Foundation/NSPOSIXErrorDomain
func (o FSVolumeOperationsObject) CreateLinkToItemNamedInDirectoryReplyHandler(item IFSItem, name IFSFileName, directory IFSItem, reply FSFileNameErrorHandler) {
	_block3, _cleanup3 := NewFSFileNameErrorBlock(reply)
	defer _cleanup3()
	objc.Send[struct{}](o.ID, objc.Sel("createLinkToItem:named:inDirectory:replyHandler:"), item, name, directory, objc.ID(_block3))
}

// Creates a new symbolic link.
//
// name: The new item’s name.
//
// directory: The directory in which to create the item.
//
// newAttributes: Attributes to apply to the new item.
//
// contents: The contents of the new symbolic link.
//
// reply: A block or closure to indicate success or failure. If creation succeeds,
// pass the newly-created [FSItem] and a `nil` error. If creation fails, pass
// the relevant error as the second parameter; FSKit ignores any [FSItem] in
// this case. For an `async` Swift implementation, there’s no reply handler;
// simply return the [FSItem] or throw an error.
//
// # Discussion
//
// If an item named `name` already exists in the directory indicated by
// `directory`, complete the request with an error with a domain of
// [NSPOSIXErrorDomain] and a code of [EEXIST].
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/createSymbolicLink(named:inDirectory:attributes:linkContents:replyHandler:)
//
// [NSPOSIXErrorDomain]: https://developer.apple.com/documentation/Foundation/NSPOSIXErrorDomain
func (o FSVolumeOperationsObject) CreateSymbolicLinkNamedInDirectoryAttributesLinkContentsReplyHandler(name IFSFileName, directory IFSItem, newAttributes IFSItemSetAttributesRequest, contents IFSFileName, reply FSItemFSFileNameErrorHandler) {
	_block4, _cleanup4 := NewFSItemFSFileNameErrorBlock(reply)
	defer _cleanup4()
	objc.Send[struct{}](o.ID, objc.Sel("createSymbolicLinkNamed:inDirectory:attributes:linkContents:replyHandler:"), name, directory, newAttributes, contents, objc.ID(_block4))
}

// Reads a symbolic link.
//
// item: The symbolic link to read from. FSKit guarantees this item is of type
// [FSItemTypeSymlink].
//
// reply: A block or closure to indicate success or failure. If reading succeeds,
// pass the link’s contents as an [FSFileName] and a `nil` error. If reading
// fails, pass the relevant error as the second parameter; FSKit ignores any
// [FSFileName] in this case. For an `async` Swift implementation, there’s
// no reply handler; simply return the [FSFileName] or throw an error.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/readSymbolicLink(_:replyHandler:)
func (o FSVolumeOperationsObject) ReadSymbolicLinkReplyHandler(item IFSItem, reply FSFileNameErrorHandler) {
	_block1, _cleanup1 := NewFSFileNameErrorBlock(reply)
	defer _cleanup1()
	objc.Send[struct{}](o.ID, objc.Sel("readSymbolicLink:replyHandler:"), item, objc.ID(_block1))
}

// Fetches attributes for the given item.
//
// desiredAttributes: A requested set of attributes to get. The implementation inspects the
// request’s [FSItemGetAttributesRequest.WantedAttributes] to determine
// which attributes to populate.
//
// item: The item to get attributes for.
//
// reply: A block or closure to indicate success or failure. If getting attributes
// succeeds, pass an [FSItemAttributes] with the requested attributes
// populated and a `nil` error. If getting attributes fails, pass the relevant
// error as the second parameter; FSKit ignores any [FSItemAttributes] in this
// case. For an `async` Swift implementation, there’s no reply handler;
// simply return the [FSItemAttributes] or throw an error.
//
// # Discussion
//
// For file systems that don’t support hard links, set
// [FSItemAttributes.LinkCount] to `1` for regular files and symbolic links.
//
// If the item’s `bsdFlags` contain the `UF_COMPRESSED` flag, your file
// system returns the uncompressed size of the file.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/getAttributes(_:of:replyHandler:)
func (o FSVolumeOperationsObject) GetAttributesOfItemReplyHandler(desiredAttributes IFSItemGetAttributesRequest, item IFSItem, reply FSItemAttributesErrorHandler) {
	_block2, _cleanup2 := NewFSItemAttributesErrorBlock(reply)
	defer _cleanup2()
	objc.Send[struct{}](o.ID, objc.Sel("getAttributes:ofItem:replyHandler:"), desiredAttributes, item, objc.ID(_block2))
}

// Sets the given attributes on an item.
//
// newAttributes: A request containing the attributes to set.
//
// item: The item on which to set the attributes.
//
// reply: A block or closure to indicate success or failure. If setting attributes
// succeeds, pass an [FSItemAttributes] with the item’s updated attributes
// and a `nil` error. If setting attributes fails, pass the relevant error as
// the second parameter; FSKit ignores any [FSItemAttributes] in this case.
// For an `async` Swift implementation, there’s no reply handler; simply
// return the [FSItemAttributes] or throw an error.
//
// # Discussion
//
// Several attributes are considered “read-only”, and an attempt to set
// these attributes results in an error with the code [EINVAL].
//
// A request may set [FSItemAttributes.Size] beyond the end of the file. If
// the underlying file system doesn’t support sparse files, allocate space
// to fill the new file size. Either fill this space with zeroes, or configure
// it to read as zeroes.
//
// If a request sets the file size below the current end-of-file, truncate the
// file and return any unused space to the file system as free space.
//
// Ignore attempts to set the size of directories or symbolic links; don’t
// produce an error.
//
// If the caller attepts to sest an attribute not supported by the on-disk
// file system format, don’t produce an error. The upper layers of the
// framework will detect this situation.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/setAttributes(_:on:replyHandler:)
func (o FSVolumeOperationsObject) SetAttributesOnItemReplyHandler(newAttributes IFSItemSetAttributesRequest, item IFSItem, reply FSItemAttributesErrorHandler) {
	_block2, _cleanup2 := NewFSItemAttributesErrorBlock(reply)
	defer _cleanup2()
	objc.Send[struct{}](o.ID, objc.Sel("setAttributes:onItem:replyHandler:"), newAttributes, item, objc.ID(_block2))
}

// Enumerates the contents of the given directory.
//
// directory: The item to enumerate. FSKit guarantees this item is of type
// [FSItemTypeDirectory].
//
// cookie: A value that indicates the location within the directory from which to
// enumerate. Your implementation defines the semantics of the cookie values;
// they’re opaque to FSKit. The first call to the enumerate method passes
// [initial] for this parameter. Subsequent calls pass whatever cookie value
// you previously passed to the packer’s `nextCookie` parmeter.
//
// verifier: A tool to detect whether the directory contents changed since the last call
// to `enumerateDirectory`. Your implementation defines the semantics of the
// verifier values; they’re opaque to FSKit. The first call to the enumerate
// method passes [initial] for this parameter. Subsequent calls pass whatever
// cookie value you previously passed to the packer’s `currentVerifier`
// parmeter.
//
// attributes: The desired attributes to provide, or `nil` if the caller doesn’t require
// attributes.
//
// packer: An object that your implementation uses to enumerate directory items,
// packing one item per callback to `enumerateDirectory`.
//
// reply: A block or closure to indicate success or failure. If enumeration succeeds,
// pass the current verifier and a `nil` error. If enumeration fails, pass the
// relevant error as the second parameter; FSKit ignores any verifier in this
// case. For an `async` Swift implementation, there’s no reply handler;
// simply return the current verifier or throw an error.
//
// # Discussion
//
// This method uses the
// [FSDirectoryEntryPacker.PackEntryWithNameItemTypeItemIDNextCookieAttributes]
// method of the `packer` parameter to deliver the enumerated items to the
// caller. The general flow of an enumeration implementation follows these
// steps:
//
// - Enumeration starts with a call to `enumerateDirectory` using the initial
// next-cookie and verifier values [initial] and [initial], respectively. -
// The implementation uses `packer` to pack the initial set of directory
// entries. Packing also sets a `nextCookie` to use on the next call. - The
// implementation replies with a new verifier value, a nonzero value that
// reflects the directory’s current version. - On the next call the
// implementation packs the next set of entries, starting with the item
// indicated by `cookie`. If `cookie` doesn’t resolve to a valid directory
// entry, complete the request with an error of domain [NSPOSIXErrorDomain]
// and code [FSErrorInvalidDirectoryCookie].
//
// When packing, make sure to use acceptable directory entry names and
// unambiguous input to all file operations that take names without additional
// normalization, such as`lookupName`.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/enumerateDirectory(_:startingAt:verifier:attributes:packer:replyHandler:)
//
// [initial]: https://developer.apple.com/documentation/FSKit/FSDirectoryCookie/initial
// [NSPOSIXErrorDomain]: https://developer.apple.com/documentation/Foundation/NSPOSIXErrorDomain
//
// [initial]: https://developer.apple.com/documentation/FSKit/FSDirectoryVerifier/initial
// [initial]: https://developer.apple.com/documentation/FSKit/FSDirectoryVerifier/initial
func (o FSVolumeOperationsObject) EnumerateDirectoryStartingAtCookieVerifierProvidingAttributesUsingPackerReplyHandler(directory IFSItem, cookie FSDirectoryCookie, verifier FSDirectoryVerifier, attributes IFSItemGetAttributesRequest, packer IFSDirectoryEntryPacker, reply FSDirectoryVerifierErrorHandler) {
	_block5, _cleanup5 := NewFSDirectoryVerifierErrorBlock(reply)
	defer _cleanup5()
	objc.Send[struct{}](o.ID, objc.Sel("enumerateDirectory:startingAtCookie:verifier:providingAttributes:usingPacker:replyHandler:"), directory, cookie, verifier, attributes, packer, objc.ID(_block5))
}

// Synchronizes the volume with its underlying resource.
//
// flags: Timing flags, as defined in `mount.H().` These flags let the file system
// know whether to run the operation in a blocking or nonblocking fashion.
//
// reply: A block or closure to indicate success or failure. If synchronization
// fails, pass an error as the one parameter to the reply handler. If
// synchronization succeeds, pass `nil`. For an `async` Swift implementation,
// there’s no reply handler; simply throw an error or return normally.
//
// # Discussion
//
// After calling this method, FSKit assumes that the volume has sent all
// pending I/O or metadata to its resource.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/synchronize(flags:replyHandler:)
func (o FSVolumeOperationsObject) SynchronizeWithFlagsReplyHandler(flags FSSyncFlags, reply ErrorHandler) {
	_block1, _cleanup1 := NewErrorBlock(reply)
	defer _cleanup1()
	objc.Send[struct{}](o.ID, objc.Sel("synchronizeWithFlags:replyHandler:"), flags, objc.ID(_block1))
}

// A property that provides the supported capabilities of the volume.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/supportedVolumeCapabilities
func (o FSVolumeOperationsObject) SupportedVolumeCapabilities() IFSVolumeSupportedCapabilities {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("supportedVolumeCapabilities"))
	return FSVolumeSupportedCapabilitiesFromID(rv)
}

// A property that provides up-to-date statistics of the volume.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/volumeStatistics
func (o FSVolumeOperationsObject) VolumeStatistics() IFSStatFSResult {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("volumeStatistics"))
	return FSStatFSResultFromID(rv)
}

// A property that allows the file system to request for specific mount
// options from FSKit.
//
// # Discussion
//
// FSKit reads this value after the volume replies to the
// [MountWithOptionsReplyHandler] call. Changing the returned value during the
// runtime of the volume has no effect.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/requestedMountOptions
func (o FSVolumeOperationsObject) RequestedMountOptions() FSMountOptions {
	rv := objc.Send[FSMountOptions](o.ID, objc.Sel("requestedMountOptions"))
	return FSMountOptions(rv)
}

func (o FSVolumeOperationsObject) SetRequestedMountOptions(value FSMountOptions) {
	objc.Send[struct{}](o.ID, objc.Sel("setRequestedMountOptions:"), value)
}

// A property that allows the file system to use open-unlink emulation.
//
// # Discussion
//
// Open-unlink functionality refers to a file system’s ability to support an
// open file being fully unlinked from the file system namespace. If a file
// system doesn’t support this functionality, FSKit can emulate it instead;
// this is called “open-unlink emulation”.
//
// Implement this property to return `true` (Swift) or [YES] (Objective-C) to
// allow FSKit to perform open-unlink emulation. If you don’t implement this
// property at all, FSKit doesn’t perform open-unlink emulation for this
// volume.
//
// FSKit reads this value after the file system replies to the `loadResource`
// message. Changing the returned value during the runtime of the volume has
// no effect.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/Operations/enableOpenUnlinkEmulation
func (o FSVolumeOperationsObject) EnableOpenUnlinkEmulation() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("enableOpenUnlinkEmulation"))
	return bool(rv)
}

func (o FSVolumeOperationsObject) SetEnableOpenUnlinkEmulation(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setEnableOpenUnlinkEmulation:"), value)
}
