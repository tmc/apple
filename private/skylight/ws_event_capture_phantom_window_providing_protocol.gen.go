// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// WSEventCapturePhantomWindowProviding protocol.
type WSEventCapturePhantomWindowProviding interface {
	objectivec.IObject

	// CreatePhantomWindowForConnection protocol.
	CreatePhantomWindowForConnection(connection *CGXConnection) uint32

	// ReleasePhantomWindowID protocol.
	ReleasePhantomWindowID(id uint32)

	// StartGestureLifetimeTrackingForPhantomWindowIDConnectionManager protocol.
	StartGestureLifetimeTrackingForPhantomWindowIDConnectionManager(id uint32, connection *CGXConnection, manager objectivec.IObject)
}

// WSEventCapturePhantomWindowProvidingObject wraps an existing Objective-C object that conforms to the WSEventCapturePhantomWindowProviding protocol.
type WSEventCapturePhantomWindowProvidingObject struct {
	objectivec.Object
}

func (o WSEventCapturePhantomWindowProvidingObject) BaseObject() objectivec.Object {
	return o.Object
}

// WSEventCapturePhantomWindowProvidingObjectFromID constructs a [WSEventCapturePhantomWindowProvidingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func WSEventCapturePhantomWindowProvidingObjectFromID(id objc.ID) WSEventCapturePhantomWindowProvidingObject {
	return WSEventCapturePhantomWindowProvidingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o WSEventCapturePhantomWindowProvidingObject) CreatePhantomWindowForConnection(connection *CGXConnection) uint32 {
	rv := objc.SendIfResponds[uint32](o.ID, objc.Sel("createPhantomWindowForConnection:"), unsafe.Pointer(connection))
	return rv
}
func (o WSEventCapturePhantomWindowProvidingObject) ReleasePhantomWindowID(id uint32) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("releasePhantomWindowID:"), id)
}
func (o WSEventCapturePhantomWindowProvidingObject) StartGestureLifetimeTrackingForPhantomWindowIDConnectionManager(id uint32, connection *CGXConnection, manager objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("startGestureLifetimeTrackingForPhantomWindowID:connection:manager:"), id, unsafe.Pointer(connection), manager)
}
