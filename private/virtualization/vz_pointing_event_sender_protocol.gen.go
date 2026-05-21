// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _VZPointingEventSender protocol.
type VZPointingEventSender interface {
	objectivec.IObject

	// SendDigitizerEventsPointingDeviceIndex protocol.
	SendDigitizerEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)

	// SendMagnifyEventsPointingDeviceIndex protocol.
	SendMagnifyEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)

	// SendMouseEventsPointingDeviceIndex protocol.
	SendMouseEventsPointingDeviceIndex(events VZOpaqueMouseEvents, index uint32)

	// SendQuickLookEventsPointingDeviceIndex protocol.
	SendQuickLookEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)

	// SendRotationEventsPointingDeviceIndex protocol.
	SendRotationEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)

	// SendScrollWheelEventsPointingDeviceIndex protocol.
	SendScrollWheelEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)

	// SendSmartMagnifyEventsPointingDeviceIndex protocol.
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

func (o VZPointingEventSenderObject) SendDigitizerEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("sendDigitizerEvents:pointingDeviceIndex:"), events, index)
}
func (o VZPointingEventSenderObject) SendMagnifyEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("sendMagnifyEvents:pointingDeviceIndex:"), events, index)
}
func (o VZPointingEventSenderObject) SendMouseEventsPointingDeviceIndex(events VZOpaqueMouseEvents, index uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("sendMouseEvents:pointingDeviceIndex:"), events.UnsafePointer(), index)
}
func (o VZPointingEventSenderObject) SendPointerNSEventPointingDeviceIndex(nSEvent objectivec.IObject, index uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("sendPointerNSEvent:pointingDeviceIndex:"), nSEvent, index)
}
func (o VZPointingEventSenderObject) SendQuickLookEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("sendQuickLookEvents:pointingDeviceIndex:"), events, index)
}
func (o VZPointingEventSenderObject) SendRotationEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("sendRotationEvents:pointingDeviceIndex:"), events, index)
}
func (o VZPointingEventSenderObject) SendScrollWheelEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("sendScrollWheelEvents:pointingDeviceIndex:"), events, index)
}
func (o VZPointingEventSenderObject) SendSmartMagnifyEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[struct{}](o.ID, objc.Sel("sendSmartMagnifyEvents:pointingDeviceIndex:"), events, index)
}
