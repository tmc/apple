// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Methods implemented for read and write operations that deliver data to and from the extension.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/ReadWriteOperations
type FSVolumeReadWriteOperations interface {
	objectivec.IObject

	// Reads the contents of the given file item.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/ReadWriteOperations/read(from:at:length:into:replyHandler:)
	ReadFromFileOffsetLengthIntoBufferReplyHandler(item IFSItem, offset int64, length uintptr, buffer IFSMutableFileDataBuffer, reply size_tErrorHandler)

	// Writes contents to the given file item.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/ReadWriteOperations/write(contents:to:at:replyHandler:)
	WriteContentsToFileAtOffsetReplyHandler(contents foundation.INSData, item IFSItem, offset int64, reply size_tErrorHandler)
}

// FSVolumeReadWriteOperationsObject wraps an existing Objective-C object that conforms to the FSVolumeReadWriteOperations protocol.
type FSVolumeReadWriteOperationsObject struct {
	objectivec.Object
}

func (o FSVolumeReadWriteOperationsObject) BaseObject() objectivec.Object {
	return o.Object
}

// FSVolumeReadWriteOperationsObjectFromID constructs a [FSVolumeReadWriteOperationsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func FSVolumeReadWriteOperationsObjectFromID(id objc.ID) FSVolumeReadWriteOperationsObject {
	return FSVolumeReadWriteOperationsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Reads the contents of the given file item.
//
// item: The item from which to read. FSKit guarantees this item will be of type
// [FSItem.ItemType.file].
//
// offset: The offset in the file from which to start reading.
//
// length: The number of bytes to read.
//
// buffer: A buffer to receive the bytes read from the file.
//
// reply: A block or closure to indicate success or failure. If reading succeeds,
// pass the number of bytes read and a `nil` error. If reading fails, pass the
// number of bytes read prior to the error along with the relevant error. For
// an `async` Swift implementation, there’s no reply handler; simply return
// the byte count or throw an error.
//
// # Discussion
//
// If the number of bytes requested exceeds the number of bytes available
// before the end of the file, then the call copies only those bytes to
// `buffer`. If `offset` points past the last valid byte of the file, don’t
// reply with an error but set `actuallyRead` to `0`.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/ReadWriteOperations/read(from:at:length:into:replyHandler:)
//
// [FSItem.ItemType.file]: https://developer.apple.com/documentation/FSKit/FSItem/ItemType/file
func (o FSVolumeReadWriteOperationsObject) ReadFromFileOffsetLengthIntoBufferReplyHandler(item IFSItem, offset int64, length uintptr, buffer IFSMutableFileDataBuffer, reply size_tErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("readFromFile:offset:length:intoBuffer:replyHandler:"), item, offset, length, buffer, reply)
}

// Writes contents to the given file item.
//
// contents: A buffer containing the data to write to the file.
//
// item: The item to which to write. FSKit guarantees this item will be of type
// [FSItem.ItemType.file].
//
// offset: The offset in the file from which to start writing.
//
// reply: A block or closure to indicate success or failure. If writing succeeds,
// pass the number of bytes written and a `nil` error. If writing fails, pass
// the number of bytes written prior to the error along with the relevant
// error. For an `async` Swift implementation, there’s no reply handler;
// simply return the byte count or throw an error.
//
// # Discussion
//
// FSKit expects this routine to allocate space in the file system to extend
// the file as necessary.
//
// If the volume experiences an out-of-space condition, reply with an error of
// domain [NSPOSIXErrorDomain] and code [ENOSPC].
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/ReadWriteOperations/write(contents:to:at:replyHandler:)
//
// [FSItem.ItemType.file]: https://developer.apple.com/documentation/FSKit/FSItem/ItemType/file
// [NSPOSIXErrorDomain]: https://developer.apple.com/documentation/Foundation/NSPOSIXErrorDomain
func (o FSVolumeReadWriteOperationsObject) WriteContentsToFileAtOffsetReplyHandler(contents foundation.INSData, item IFSItem, offset int64, reply size_tErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("writeContents:toFile:atOffset:replyHandler:"), contents, item, offset, reply)
}
