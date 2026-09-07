// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _VZIOHIDEventSender protocol.
type VZIOHIDEventSender interface {
	objectivec.IObject

	// SendIOHIDEventsHidDeviceIndex protocol.
	SendIOHIDEventsHidDeviceIndex(iOHIDEvents VZOpaqueIOHIDEvents, index uint32)
}

// VZIOHIDEventSenderObject wraps an existing Objective-C object that conforms to the VZIOHIDEventSender protocol.
type VZIOHIDEventSenderObject struct {
	objectivec.Object
}

func (o VZIOHIDEventSenderObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZIOHIDEventSenderObjectFromID constructs a [VZIOHIDEventSenderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZIOHIDEventSenderObjectFromID(id objc.ID) VZIOHIDEventSenderObject {
	return VZIOHIDEventSenderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o VZIOHIDEventSenderObject) SendIOHIDEventsHidDeviceIndex(iOHIDEvents VZOpaqueIOHIDEvents, index uint32) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("sendIOHIDEvents:hidDeviceIndex:"), iOHIDEvents.UnsafePointer(), index)
}
