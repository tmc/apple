// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Methods and properties implemented by volumes that want to offer preallocation functions.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/PreallocateOperations
type FSVolumePreallocateOperations interface {
	objectivec.IObject

	// Prealocates disk space for the given item.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/PreallocateOperations/preallocateSpace(for:at:length:flags:replyHandler:)
	PreallocateSpaceForItemAtOffsetLengthFlagsReplyHandler(item IFSItem, offset int64, length uintptr, flags FSPreallocateFlags, reply size_tErrorHandler)

	// A Boolean value that instructs FSKit not to call this protocol’s methods, even if the volume conforms to it.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/PreallocateOperations/isPreallocateInhibited
	IsPreallocateInhibited() bool

	// A Boolean value that instructs FSKit not to call this protocol’s methods, even if the volume conforms to it.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/PreallocateOperations/isPreallocateInhibited
	SetPreallocateInhibited(value bool)
}

// FSVolumePreallocateOperationsObject wraps an existing Objective-C object that conforms to the FSVolumePreallocateOperations protocol.
type FSVolumePreallocateOperationsObject struct {
	objectivec.Object
}

func (o FSVolumePreallocateOperationsObject) BaseObject() objectivec.Object {
	return o.Object
}

// FSVolumePreallocateOperationsObjectFromID constructs a [FSVolumePreallocateOperationsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func FSVolumePreallocateOperationsObjectFromID(id objc.ID) FSVolumePreallocateOperationsObject {
	return FSVolumePreallocateOperationsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Prealocates disk space for the given item.
//
// item: The item for which to preallocate space.
//
// offset: The offset from which to allocate.
//
// length: The length of the space in bytes.
//
// flags: Flags that affect the preallocation behavior.
//
// reply: A block or closure to indicate success or failure. If preallocation
// succeeds, pass the amount of bytes allocated and a `nil` error. If
// preallocation fails, pass the relevant error as the second parameter; FSKit
// ignores any byte count in this case. For an `async` Swift implementation,
// there’s no reply handler; simply return the allocated byte count or throw
// an error.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/PreallocateOperations/preallocateSpace(for:at:length:flags:replyHandler:)
func (o FSVolumePreallocateOperationsObject) PreallocateSpaceForItemAtOffsetLengthFlagsReplyHandler(item IFSItem, offset int64, length uintptr, flags FSPreallocateFlags, reply size_tErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("preallocateSpaceForItem:atOffset:length:flags:replyHandler:"), item, offset, length, flags, reply)
}

// A Boolean value that instructs FSKit not to call this protocol’s methods,
// even if the volume conforms to it.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/PreallocateOperations/isPreallocateInhibited
func (o FSVolumePreallocateOperationsObject) IsPreallocateInhibited() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isPreallocateInhibited"))
	return rv
}

// A Boolean value that instructs FSKit not to call this protocol’s methods,
// even if the volume conforms to it.
//
// # Discussion
//
// FSKit reads this value after the file system replies to the `loadResource`
// message. Changing the returned value during the runtime of the volume has
// no effect.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/PreallocateOperations/isPreallocateInhibited
func (o FSVolumePreallocateOperationsObject) SetPreallocateInhibited(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setPreallocateInhibited:"), value)
}
