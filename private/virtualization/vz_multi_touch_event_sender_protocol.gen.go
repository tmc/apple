// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _VZMultiTouchEventSender protocol.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMultiTouchEventSender
type VZMultiTouchEventSender interface {
	objectivec.IObject

	// SendMultiTouchEventsMultiTouchDeviceIndex protocol.
	//
	// See: https://developer.apple.com/documentation/Virtualization/_VZMultiTouchEventSender/sendMultiTouchEvents:multiTouchDeviceIndex:
	SendMultiTouchEventsMultiTouchDeviceIndex(events unsafe.Pointer, index uint32)
}

// VZMultiTouchEventSenderObject wraps an existing Objective-C object that conforms to the VZMultiTouchEventSender protocol.
type VZMultiTouchEventSenderObject struct {
	objectivec.Object
}
func (o VZMultiTouchEventSenderObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZMultiTouchEventSenderObjectFromID constructs a [VZMultiTouchEventSenderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZMultiTouchEventSenderObjectFromID(id objc.ID) VZMultiTouchEventSenderObject {
	return VZMultiTouchEventSenderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZMultiTouchEventSender/sendMultiTouchEvents:multiTouchDeviceIndex:
func (o VZMultiTouchEventSenderObject) SendMultiTouchEventsMultiTouchDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("sendMultiTouchEvents:multiTouchDeviceIndex:"), events, index)
	}

