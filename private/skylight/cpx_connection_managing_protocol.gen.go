// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// CPXConnectionManaging protocol.
type CPXConnectionManaging interface {
	objectivec.IObject

	// ConnectionForID protocol.
	ConnectionForID(id uint32) *CGXConnection

	// PidForConnection protocol.
	PidForConnection(connection *CGXConnection) int
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

func (o CPXConnectionManagingObject) ConnectionForID(id uint32) *CGXConnection {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("connectionForID:"), id)
	return (*CGXConnection)(rv)
}
func (o CPXConnectionManagingObject) PidForConnection(connection *CGXConnection) int {
	rv := objc.SendIfResponds[int](o.ID, objc.Sel("pidForConnection:"), unsafe.Pointer(connection))
	return rv
}
