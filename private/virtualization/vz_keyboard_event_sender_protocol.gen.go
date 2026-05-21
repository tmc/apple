// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _VZKeyboardEventSender protocol.
type VZKeyboardEventSender interface {
	objectivec.IObject

	// SendKeyboardEventsKeyboardID protocol.
	SendKeyboardEventsKeyboardID(events VZOpaqueKeyboardEvents, id uint32)
}

// VZKeyboardEventSenderObject wraps an existing Objective-C object that conforms to the VZKeyboardEventSender protocol.
type VZKeyboardEventSenderObject struct {
	objectivec.Object
}

func (o VZKeyboardEventSenderObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZKeyboardEventSenderObjectFromID constructs a [VZKeyboardEventSenderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZKeyboardEventSenderObjectFromID(id objc.ID) VZKeyboardEventSenderObject {
	return VZKeyboardEventSenderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o VZKeyboardEventSenderObject) SendKeyboardEventsKeyboardID(events VZOpaqueKeyboardEvents, id uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("sendKeyboardEvents:keyboardID:"), events.UnsafePointer(), id)
}
