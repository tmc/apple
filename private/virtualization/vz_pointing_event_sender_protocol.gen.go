// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _VZPointingEventSender protocol.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZPointingEventSender
type VZPointingEventSender interface {
	objectivec.IObject

	// SendDigitizerEventsPointingDeviceIndex protocol.
	//
	// See: https://developer.apple.com/documentation/Virtualization/_VZPointingEventSender/sendDigitizerEvents:pointingDeviceIndex:
	SendDigitizerEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)

	// SendMagnifyEventsPointingDeviceIndex protocol.
	//
	// See: https://developer.apple.com/documentation/Virtualization/_VZPointingEventSender/sendMagnifyEvents:pointingDeviceIndex:
	SendMagnifyEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)

	// SendMouseEventsPointingDeviceIndex protocol.
	//
	// See: https://developer.apple.com/documentation/Virtualization/_VZPointingEventSender/sendMouseEvents:pointingDeviceIndex:
	SendMouseEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)

	// SendQuickLookEventsPointingDeviceIndex protocol.
	//
	// See: https://developer.apple.com/documentation/Virtualization/_VZPointingEventSender/sendQuickLookEvents:pointingDeviceIndex:
	SendQuickLookEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)

	// SendRotationEventsPointingDeviceIndex protocol.
	//
	// See: https://developer.apple.com/documentation/Virtualization/_VZPointingEventSender/sendRotationEvents:pointingDeviceIndex:
	SendRotationEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)

	// SendScrollWheelEventsPointingDeviceIndex protocol.
	//
	// See: https://developer.apple.com/documentation/Virtualization/_VZPointingEventSender/sendScrollWheelEvents:pointingDeviceIndex:
	SendScrollWheelEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)

	// SendSmartMagnifyEventsPointingDeviceIndex protocol.
	//
	// See: https://developer.apple.com/documentation/Virtualization/_VZPointingEventSender/sendSmartMagnifyEvents:pointingDeviceIndex:
	SendSmartMagnifyEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)
}

// VZPointingEventSenderObject wraps an existing Objective-C object that conforms to the VZPointingEventSender protocol.
type VZPointingEventSenderObject struct {
	objectivec.Object
}

func (o VZPointingEventSenderObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZPointingEventSenderObjectFromID constructs a [VZPointingEventSenderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZPointingEventSenderObjectFromID(id objc.ID) VZPointingEventSenderObject {
	return VZPointingEventSenderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPointingEventSender/sendDigitizerEvents:pointingDeviceIndex:
func (o VZPointingEventSenderObject) SendDigitizerEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("sendDigitizerEvents:pointingDeviceIndex:"), events, index)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPointingEventSender/sendMagnifyEvents:pointingDeviceIndex:
func (o VZPointingEventSenderObject) SendMagnifyEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("sendMagnifyEvents:pointingDeviceIndex:"), events, index)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPointingEventSender/sendMouseEvents:pointingDeviceIndex:
func (o VZPointingEventSenderObject) SendMouseEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("sendMouseEvents:pointingDeviceIndex:"), events, index)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPointingEventSender/sendPointerNSEvent:pointingDeviceIndex:
func (o VZPointingEventSenderObject) SendPointerNSEventPointingDeviceIndex(nSEvent objectivec.IObject, index uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("sendPointerNSEvent:pointingDeviceIndex:"), nSEvent, index)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPointingEventSender/sendQuickLookEvents:pointingDeviceIndex:
func (o VZPointingEventSenderObject) SendQuickLookEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("sendQuickLookEvents:pointingDeviceIndex:"), events, index)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPointingEventSender/sendRotationEvents:pointingDeviceIndex:
func (o VZPointingEventSenderObject) SendRotationEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("sendRotationEvents:pointingDeviceIndex:"), events, index)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPointingEventSender/sendScrollWheelEvents:pointingDeviceIndex:
func (o VZPointingEventSenderObject) SendScrollWheelEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("sendScrollWheelEvents:pointingDeviceIndex:"), events, index)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZPointingEventSender/sendSmartMagnifyEvents:pointingDeviceIndex:
func (o VZPointingEventSenderObject) SendSmartMagnifyEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("sendSmartMagnifyEvents:pointingDeviceIndex:"), events, index)
}
