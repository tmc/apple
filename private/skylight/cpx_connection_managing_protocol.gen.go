// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXConnectionManaging protocol.
//
// See: https://developer.apple.com/documentation/SkyLight/CPXConnectionManaging
type CPXConnectionManaging interface {
	objectivec.IObject

	// ConnectionForID protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXConnectionManaging/connectionForID:
	ConnectionForID(id uint32) unsafe.Pointer

	// PidForConnection protocol.
	//
	// See: https://developer.apple.com/documentation/SkyLight/CPXConnectionManaging/pidForConnection:
	PidForConnection(connection unsafe.Pointer) int
}

// CPXConnectionManagingObject wraps an existing Objective-C object that conforms to the CPXConnectionManaging protocol.
type CPXConnectionManagingObject struct {
	objectivec.Object
}

func (o CPXConnectionManagingObject) BaseObject() objectivec.Object {
	return o.Object
}

// CPXConnectionManagingObjectFromID constructs a [CPXConnectionManagingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func CPXConnectionManagingObjectFromID(id objc.ID) CPXConnectionManagingObject {
	return CPXConnectionManagingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/SkyLight/CPXConnectionManaging/connectionForID:
func (o CPXConnectionManagingObject) ConnectionForID(id uint32) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("connectionForID:"), id)
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/CPXConnectionManaging/pidForConnection:
func (o CPXConnectionManagingObject) PidForConnection(connection unsafe.Pointer) int {
	rv := objc.Send[int](o.ID, objc.Sel("pidForConnection:"), connection)
	return rv
}
