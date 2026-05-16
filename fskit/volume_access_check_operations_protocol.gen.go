// Code generated from Apple documentation for FSKit. DO NOT EDIT.

package fskit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Methods and properties implemented by volumes that want to enforce access check operations.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/AccessCheckOperations
type FSVolumeAccessCheckOperations interface {
	objectivec.IObject

	// Checks whether the file system allows access to the given item.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/AccessCheckOperations/checkAccess(to:requestedAccess:replyHandler:)
	CheckAccessToItemRequestedAccessReplyHandler(theItem IFSItem, access FSAccessMask, reply BoolErrorHandler)

	// A Boolean value that instructs FSKit not to call this protocol’s methods, even if the volume conforms to it.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/AccessCheckOperations/isAccessCheckInhibited
	IsAccessCheckInhibited() bool

	// A Boolean value that instructs FSKit not to call this protocol’s methods, even if the volume conforms to it.
	//
	// See: https://developer.apple.com/documentation/FSKit/FSVolume/AccessCheckOperations/isAccessCheckInhibited
	SetAccessCheckInhibited(value bool)
}

// FSVolumeAccessCheckOperationsObject wraps an existing Objective-C object that conforms to the FSVolumeAccessCheckOperations protocol.
type FSVolumeAccessCheckOperationsObject struct {
	objectivec.Object
}

func (o FSVolumeAccessCheckOperationsObject) BaseObject() objectivec.Object {
	return o.Object
}

// FSVolumeAccessCheckOperationsObjectFromID constructs a [FSVolumeAccessCheckOperationsObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func FSVolumeAccessCheckOperationsObjectFromID(id objc.ID) FSVolumeAccessCheckOperationsObject {
	return FSVolumeAccessCheckOperationsObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Checks whether the file system allows access to the given item.
//
// theItem: The item for which to check access.
//
// access: A mask indicating a set of access types for which to check.
//
// reply: A block or closure to indicate success or failure. If the access check
// succeeds, pass a Boolean value to indicate whether the file system grants
// access, followed by a `nil` error. If the access check fails, pass the
// relevant error as the second parameter; FSKit ignores the Boolean parameter
// in this case. For an `async` Swift implementation, there’s no reply
// handler; simply return the [Bool] or throw an error.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/AccessCheckOperations/checkAccess(to:requestedAccess:replyHandler:)
func (o FSVolumeAccessCheckOperationsObject) CheckAccessToItemRequestedAccessReplyHandler(theItem IFSItem, access FSAccessMask, reply BoolErrorHandler) {
	objc.Send[struct{}](o.ID, objc.Sel("checkAccessToItem:requestedAccess:replyHandler:"), theItem, access, reply)
}

// A Boolean value that instructs FSKit not to call this protocol’s methods,
// even if the volume conforms to it.
//
// See: https://developer.apple.com/documentation/FSKit/FSVolume/AccessCheckOperations/isAccessCheckInhibited
func (o FSVolumeAccessCheckOperationsObject) IsAccessCheckInhibited() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessCheckInhibited"))
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
// See: https://developer.apple.com/documentation/FSKit/FSVolume/AccessCheckOperations/isAccessCheckInhibited
func (o FSVolumeAccessCheckOperationsObject) SetAccessCheckInhibited(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessCheckInhibited:"), value)
}
