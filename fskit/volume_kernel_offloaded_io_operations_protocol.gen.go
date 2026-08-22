// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Methods and properties implemented by volumes that use kernel-offloaded I/O to achieve higher file transfer performance.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolumeKernelOffloadedIOOperations
type FSVolumeKernelOffloadedIOOperations interface {
	objectivec.IObject

	// Maps a file’s disk space into extents, allowing the kernel to perform I/O with that space.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolumeKernelOffloadedIOOperations/blockmapFile(_:offset:length:flags:operationID:packer:replyHandler:)
	BlockmapFileOffsetLengthFlagsOperationIDPackerReplyHandler(file IFSItem, offset int64, length uintptr, flags FSBlockmapFlags, operationID FSOperationID, packer IFSExtentPacker, reply ErrorHandler)

	// Completes an I/O operation for a given file.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolumeKernelOffloadedIOOperations/completeIO(for:offset:length:status:flags:operationID:replyHandler:)
	CompleteIOForFileOffsetLengthStatusFlagsOperationIDReplyHandler(file IFSItem, offset int64, length uintptr, status foundation.NSError, flags FSCompleteIOFlags, operationID FSOperationID, reply ErrorHandler)

	// Creates a new file item and map its disk space.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolumeKernelOffloadedIOOperations/createFile(name:in:attributes:packer:replyHandler:)
	CreateFileNamedInDirectoryAttributesPackerReplyHandler(name IFSFileName, directory IFSItem, attributes IFSItemSetAttributesRequest, packer IFSExtentPacker, reply FSItemFSFileNameErrorHandler)

	// Looks up an item within a directory and maps its disk space.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolumeKernelOffloadedIOOperations/lookupItem(name:in:packer:replyHandler:)
	LookupItemNamedInDirectoryPackerReplyHandler(name IFSFileName, directory IFSItem, packer IFSExtentPacker, reply FSItemFSFileNameErrorHandler)
}

// FSVolumeKernelOffloadedIOOperationsObject wraps an existing Objective-C object that conforms to the FSVolumeKernelOffloadedIOOperations protocol.
type FSVolumeKernelOffloadedIOOperationsObject struct {
	objectivec.Object
}

func (o FSVolumeKernelOffloadedIOOperationsObject) BaseObject() objectivec.Object {
	return o.Object
}

// FSVolumeKernelOffloadedIOOperationsObjectFromID constructs a [FSVolumeKernelOffloadedIOOperationsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func FSVolumeKernelOffloadedIOOperationsObjectFromID(id objc.ID) FSVolumeKernelOffloadedIOOperationsObject {
	return FSVolumeKernelOffloadedIOOperationsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Maps a file’s disk space into extents, allowing the kernel to perform I/O
// with that space.
//
// file: The file for which to map disk space.
//
// offset: The starting logical offset of the range to be mapped (in bytes).
//
// length: The length of the range to be mapped (in bytes).
//
// flags: Flags that affect the behavior of the blockmap operation.
//
// operationID: A unique identifier of the blockmap call. Any value other than `0`
// (Objective-C) or [unspecified] (Swift) indicates the beginning of an I/O
// operation. A value of `0` or [unspecified] indicates the kernel maps the
// file without performing I/O. In this case, FSKit doesn’t perform a
// corresponding call to
// [CompleteIOForFileOffsetLengthStatusFlagsOperationIDReplyHandler].
//
// packer: An extent packer you use to pack the requested range of the file’s
// allocated disk space. FSKit sends all of the packed extents to the kernel
// when it invokes `reply`.
//
// reply: A block or closure to indicate success or failure. If mapping fails, pass
// an error as the one parameter to the reply handler. If mapping succeeds,
// pass `nil`. For an `async` Swift implementation, there’s no reply
// handler; simply throw an error or return normally.
//
// # Discussion
//
// FSKit calls this method when the kernel needs to get a mapping of
// logical-to-physical offsets of the file’s data. This call may occur as
// part of an I/O operation on the file, or just to get the mapping as part of
// a `fcntl(F_LOG2PHYS)` system call. In the case of an I/O operation on the
// file, `operationID` has a nonzero value; a future call to
// [CompleteIOForFileOffsetLengthStatusFlagsOperationIDReplyHandler] uses the
// same `operationID` to indicate which operation it completes. In the case of
// a `fcntl(F_LOG2PHYS)` system call, the `operationID` parameter is `0`
// (Objective-C) or [unspecified] (Swift). In both cases the kernel retains
// the mapping, and it may perform I/O to this range (or a part of it) at any
// time.
//
// If satisfying a blockmap request requires more extents than `packer` can
// handle, FSKit makes additional calls to this method with the same operation
// ID to collect the remainder.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolumeKernelOffloadedIOOperations/blockmapFile(_:offset:length:flags:operationID:packer:replyHandler:)
//
// [unspecified]: https://developer.apple.com/documentation/FSKit/FSOperationID/unspecified
//
// [unspecified]: https://developer.apple.com/documentation/FSKit/FSOperationID/unspecified
func (o FSVolumeKernelOffloadedIOOperationsObject) BlockmapFileOffsetLengthFlagsOperationIDPackerReplyHandler(file IFSItem, offset int64, length uintptr, flags FSBlockmapFlags, operationID FSOperationID, packer IFSExtentPacker, reply ErrorHandler) {
	_block6, _cleanup6 := NewErrorBlock(reply)
	defer _cleanup6()
	objc.Send[struct{}](o.ID, objc.Sel("blockmapFile:offset:length:flags:operationID:packer:replyHandler:"), file, offset, length, flags, operationID, packer, objc.ID(_block6))
}

// Completes an I/O operation for a given file.
//
// file: The file for which the I/O operation completed.
//
// offset: The starting logical offset at which I/O started.
//
// length: The length of the I/O range (in bytes).
//
// status: Any error that occurred during the operation. If no error occurred, this
// parameter is `nil`.
//
// flags: Flags that affect the behavior of the complete I/O operation.
//
// operationID: A unique identifier of the blockmap call. Any value other than `0`
// (Objective-C) or [unspecified] (Swift) corresponds to a previous call to
// [BlockmapFileOffsetLengthFlagsOperationIDPackerReplyHandler] with the same
// `operationID`.
//
// reply: A block or closure to indicate success or failure. If completing I/O fails,
// pass an error as the one parameter to the reply handler. If completing I/O
// succeeds, pass `nil`. For an `async` Swift implementation, there’s no
// reply handler; simply throw an error or return normally.
//
// # Discussion
//
// Implement this method by updating a file’s metadata, such as its size and
// modification time.
//
// FSKit may call this method without an earlier call to
// [BlockmapFileOffsetLengthFlagsOperationIDPackerReplyHandler]. In this case,
// the `operationID` is `0` (Objective-C) or [unspecified] (Swift).
//
// See: https://developer.apple.com/documentation/FSKit/FSVolumeKernelOffloadedIOOperations/completeIO(for:offset:length:status:flags:operationID:replyHandler:)
//
// [unspecified]: https://developer.apple.com/documentation/FSKit/FSOperationID/unspecified
//
// [unspecified]: https://developer.apple.com/documentation/FSKit/FSOperationID/unspecified
func (o FSVolumeKernelOffloadedIOOperationsObject) CompleteIOForFileOffsetLengthStatusFlagsOperationIDReplyHandler(file IFSItem, offset int64, length uintptr, status foundation.NSError, flags FSCompleteIOFlags, operationID FSOperationID, reply ErrorHandler) {
	_block6, _cleanup6 := NewErrorBlock(reply)
	defer _cleanup6()
	objc.Send[struct{}](o.ID, objc.Sel("completeIOForFile:offset:length:status:flags:operationID:replyHandler:"), file, offset, length, status, flags, operationID, objc.ID(_block6))
}

// Creates a new file item and map its disk space.
//
// name: The new file’s name.
//
// directory: The directory in which to create the file.
//
// attributes: Attributes to apply to the new file.
//
// packer: An extent packer you use to pack the file’s allocated disk space.
//
// reply: A block or closure to indicate success or failure. If creation succeeds,
// pass the newly created [FSItem] and its [FSFileName], along with a `nil`
// error. If creation fails, pass the relevant error as the third parameter;
// FSKit ignores any [FSItem] or [FSFileName] in this case. For an `async`
// Swift implementation, there’s no reply handler; instead, return a tuple
// of the [FSItem] and its [FSFileName] or throw an error.
//
// # Discussion
//
// This method allows the module to opportunistically supply extents, avoiding
// future calls to
// [BlockmapFileOffsetLengthFlagsOperationIDPackerReplyHandler]. Only perform
// this technique opportunistically. In particular, don’t perform additional
// I/O to fetch extent data.
//
// Packing extents in this method requires that `attributes` defines a size
// greater than 0.
//
// An implementation that doesn’t supply the extents can ignore the packer
// and call the corresponding method in the [FSVolumeOperations] protocol,
// [CreateItemNamedTypeInDirectoryAttributesReplyHandler].
//
// See: https://developer.apple.com/documentation/FSKit/FSVolumeKernelOffloadedIOOperations/createFile(name:in:attributes:packer:replyHandler:)
func (o FSVolumeKernelOffloadedIOOperationsObject) CreateFileNamedInDirectoryAttributesPackerReplyHandler(name IFSFileName, directory IFSItem, attributes IFSItemSetAttributesRequest, packer IFSExtentPacker, reply FSItemFSFileNameErrorHandler) {
	_block4, _cleanup4 := NewFSItemFSFileNameErrorBlock(reply)
	defer _cleanup4()
	objc.Send[struct{}](o.ID, objc.Sel("createFileNamed:inDirectory:attributes:packer:replyHandler:"), name, directory, attributes, packer, objc.ID(_block4))
}

// Looks up an item within a directory and maps its disk space.
//
// name: The name of the file to look up.
//
// directory: The directory in which to look up the file.
//
// packer: An extent packer you use to pack the file’s allocated disk space.
//
// reply: A block or closure to indicate success or failure. If lookup succeeds, pass
// the found [FSItem] and its [FSFileName], along with a `nil` error. If
// lookup fails, pass the relevant error as the third parameter; FSKit ignores
// any [FSItem] or [FSFileName] in this case. For an `async` Swift
// implementation, there’s no reply handler; instead, return a tuple of the
// [FSItem] and its [FSFileName] or throw an error.
//
// # Discussion
//
// This method allows the module to opportunistically supply extents, avoiding
// future calls to
// [BlockmapFileOffsetLengthFlagsOperationIDPackerReplyHandler]. Only perform
// this technique opportunistically. In particular, don’t perform additional
// I/O to fetch extent data.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolumeKernelOffloadedIOOperations/lookupItem(name:in:packer:replyHandler:)
func (o FSVolumeKernelOffloadedIOOperationsObject) LookupItemNamedInDirectoryPackerReplyHandler(name IFSFileName, directory IFSItem, packer IFSExtentPacker, reply FSItemFSFileNameErrorHandler) {
	_block3, _cleanup3 := NewFSItemFSFileNameErrorBlock(reply)
	defer _cleanup3()
	objc.Send[struct{}](o.ID, objc.Sel("lookupItemNamed:inDirectory:packer:replyHandler:"), name, directory, packer, objc.ID(_block3))
}

// Preallocates and maps disk space for the given file.
//
// file: The item for which to preallocate space.
//
// offset: The offset from which to allocate.
//
// length: The length of the space in bytes.
//
// flags: Flags that affect the preallocation behavior.
//
// packer: An extent packer you use to pack the file’s preallocated disk space.
//
// reply: A block or closure to indicate success or failure. If preallocation
// succeeds, pass the amount of bytes allocated and a nil error. If
// preallocation fails, pass the relevant error as the second parameter; FSKit
// ignores any byte count in this case. For an `async` Swift implementation,
// there’s no reply handler; simply return the allocated byte count or throw
// an error.
//
// # Discussion
//
// This method allows the module to opportunistically supply extents, avoiding
// future calls to
// [BlockmapFileOffsetLengthFlagsOperationIDPackerReplyHandler].
//
// See: https://developer.apple.com/documentation/FSKit/FSVolumeKernelOffloadedIOOperations/preallocateSpace(for:at:length:flags:packer:replyHandler:)
func (o FSVolumeKernelOffloadedIOOperationsObject) PreallocateSpaceForFileAtOffsetLengthFlagsPackerReplyHandler(file IFSItem, offset int64, length uintptr, flags FSPreallocateFlags, packer IFSExtentPacker, reply size_tErrorHandler) {
	_block5, _cleanup5 := Newsize_tErrorBlock(reply)
	defer _cleanup5()
	objc.Send[struct{}](o.ID, objc.Sel("preallocateSpaceForFile:atOffset:length:flags:packer:replyHandler:"), file, offset, length, flags, packer, objc.ID(_block5))
}
