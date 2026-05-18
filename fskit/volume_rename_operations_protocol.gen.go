// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Methods and properties implemented by volumes that support renaming the volume.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/RenameOperations
type FSVolumeRenameOperations interface {
	objectivec.IObject

	// Sets a new name for the volume.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/RenameOperations/setVolumeName(_:replyHandler:)
	SetVolumeNameReplyHandler(name IFSFileName, reply FSFileNameErrorHandler)

	// A Boolean value that instructs FSKit not to call this protocol’s methods, even if the volume conforms to it.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/RenameOperations/isVolumeRenameInhibited
	IsVolumeRenameInhibited() bool

	// A Boolean value that instructs FSKit not to call this protocol’s methods, even if the volume conforms to it.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/RenameOperations/isVolumeRenameInhibited
	VolumeRenameInhibited() bool
	SetVolumeRenameInhibited(value bool)
}

// FSVolumeRenameOperationsObject wraps an existing Objective-C object that conforms to the FSVolumeRenameOperations protocol.
type FSVolumeRenameOperationsObject struct {
	objectivec.Object
}

func (o FSVolumeRenameOperationsObject) BaseObject() objectivec.Object {
	return o.Object
}

// FSVolumeRenameOperationsObjectFromID constructs a [FSVolumeRenameOperationsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func FSVolumeRenameOperationsObjectFromID(id objc.ID) FSVolumeRenameOperationsObject {
	return FSVolumeRenameOperationsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Sets a new name for the volume.
//
// name: The new volume name.
//
// reply: A block or closure to indicate success or failure. If renaming succeeds,
// pass an [FSFileName] of the new volume name and a `nil` error. If renaming
// fails, pass the relevant error as the second parameter; FSKit ignores any
// [FSFileName] in this case. For an `async` Swift implementation, there’s
// no reply handler; simply return the [FSFileName] or throw an error.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/RenameOperations/setVolumeName(_:replyHandler:)
func (o FSVolumeRenameOperationsObject) SetVolumeNameReplyHandler(name IFSFileName, reply FSFileNameErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("setVolumeName:replyHandler:"), name, reply)
}

// A Boolean value that instructs FSKit not to call this protocol’s methods,
// even if the volume conforms to it.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/RenameOperations/isVolumeRenameInhibited
func (o FSVolumeRenameOperationsObject) IsVolumeRenameInhibited() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isVolumeRenameInhibited"))
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
// See: https://developer.apple.com/documentation/FSKit/FSVolume/RenameOperations/isVolumeRenameInhibited
func (o FSVolumeRenameOperationsObject) VolumeRenameInhibited() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isVolumeRenameInhibited"))
	return bool(rv)
}

func (o FSVolumeRenameOperationsObject) SetVolumeRenameInhibited(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setVolumeRenameInhibited:"), value)
}
