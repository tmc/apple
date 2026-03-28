// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _VZKeyboardEventSender protocol.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZKeyboardEventSender
type VZKeyboardEventSender interface {
	objectivec.IObject

	// SendKeyboardEventsKeyboardID protocol.
	//
	// See: https://developer.apple.com/documentation/Virtualization/_VZKeyboardEventSender/sendKeyboardEvents:keyboardID:
	SendKeyboardEventsKeyboardID(events unsafe.Pointer, id uint32)
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

//
// See: https://developer.apple.com/documentation/Virtualization/_VZKeyboardEventSender/sendKeyboardEvents:keyboardID:
func (o VZKeyboardEventSenderObject) SendKeyboardEventsKeyboardID(events unsafe.Pointer, id uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("sendKeyboardEvents:keyboardID:"), events, id)
	}

