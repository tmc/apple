// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZVirtualMachineAccessor] class.
var (
	_VZVirtualMachineAccessorClass     VZVirtualMachineAccessorClass
	_VZVirtualMachineAccessorClassOnce sync.Once
)

func getVZVirtualMachineAccessorClass() VZVirtualMachineAccessorClass {
	_VZVirtualMachineAccessorClassOnce.Do(func() {
		_VZVirtualMachineAccessorClass = VZVirtualMachineAccessorClass{class: objc.GetClass("_VZVirtualMachineAccessor")}
	})
	return _VZVirtualMachineAccessorClass
}

// GetVZVirtualMachineAccessorClass returns the class object for _VZVirtualMachineAccessor.
func GetVZVirtualMachineAccessorClass() VZVirtualMachineAccessorClass {
	return getVZVirtualMachineAccessorClass()
}

type VZVirtualMachineAccessorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZVirtualMachineAccessorClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZVirtualMachineAccessorClass) Alloc() VZVirtualMachineAccessor {
	rv := objc.Send[VZVirtualMachineAccessor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVirtualMachineAccessor._hidEventMonitor]
//   - [VZVirtualMachineAccessor._processHIDReportsForDeviceDeviceType]
//   - [VZVirtualMachineAccessor._shouldSendHIDReports]
//   - [VZVirtualMachineAccessor.SendDigitizerEventsPointingDeviceIndex]
//   - [VZVirtualMachineAccessor.SendIOHIDEventsHidDeviceIndex]
//   - [VZVirtualMachineAccessor.SendKeyboardEventsKeyboardID]
//   - [VZVirtualMachineAccessor.SendMagnifyEventsPointingDeviceIndex]
//   - [VZVirtualMachineAccessor.SendMouseEventsPointingDeviceIndex]
//   - [VZVirtualMachineAccessor.SendMultiTouchEventsMultiTouchDeviceIndex]
//   - [VZVirtualMachineAccessor.SendPointerNSEventPointingDeviceIndex]
//   - [VZVirtualMachineAccessor.SendQuickLookEventsPointingDeviceIndex]
//   - [VZVirtualMachineAccessor.SendRotationEventsPointingDeviceIndex]
//   - [VZVirtualMachineAccessor.SendScrollWheelEventsPointingDeviceIndex]
//   - [VZVirtualMachineAccessor.SendSmartMagnifyEventsPointingDeviceIndex]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineAccessor
type VZVirtualMachineAccessor struct {
	objectivec.Object
}

// VZVirtualMachineAccessorFromID constructs a [VZVirtualMachineAccessor] from an objc.ID.
func VZVirtualMachineAccessorFromID(id objc.ID) VZVirtualMachineAccessor {
	return VZVirtualMachineAccessor{objectivec.Object{ID: id}}
}

// Ensure VZVirtualMachineAccessor implements IVZVirtualMachineAccessor.
var _ IVZVirtualMachineAccessor = VZVirtualMachineAccessor{}

// An interface definition for the [VZVirtualMachineAccessor] class.
//
// # Methods
//
//   - [IVZVirtualMachineAccessor._hidEventMonitor]
//   - [IVZVirtualMachineAccessor._processHIDReportsForDeviceDeviceType]
//   - [IVZVirtualMachineAccessor._shouldSendHIDReports]
//   - [IVZVirtualMachineAccessor.SendDigitizerEventsPointingDeviceIndex]
//   - [IVZVirtualMachineAccessor.SendIOHIDEventsHidDeviceIndex]
//   - [IVZVirtualMachineAccessor.SendKeyboardEventsKeyboardID]
//   - [IVZVirtualMachineAccessor.SendMagnifyEventsPointingDeviceIndex]
//   - [IVZVirtualMachineAccessor.SendMouseEventsPointingDeviceIndex]
//   - [IVZVirtualMachineAccessor.SendMultiTouchEventsMultiTouchDeviceIndex]
//   - [IVZVirtualMachineAccessor.SendPointerNSEventPointingDeviceIndex]
//   - [IVZVirtualMachineAccessor.SendQuickLookEventsPointingDeviceIndex]
//   - [IVZVirtualMachineAccessor.SendRotationEventsPointingDeviceIndex]
//   - [IVZVirtualMachineAccessor.SendScrollWheelEventsPointingDeviceIndex]
//   - [IVZVirtualMachineAccessor.SendSmartMagnifyEventsPointingDeviceIndex]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineAccessor
type IVZVirtualMachineAccessor interface {
	objectivec.IObject

	// Topic: Methods

	_hidEventMonitor() IVZHIDEventMonitor
	_processHIDReportsForDeviceDeviceType(hIDReports VZOpaqueHIDReports, device uint32, type_ int32)
	_shouldSendHIDReports() bool
	SendDigitizerEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)
	SendIOHIDEventsHidDeviceIndex(iOHIDEvents VZOpaqueIOHIDEvents, index uint32)
	SendKeyboardEventsKeyboardID(events VZOpaqueKeyboardEvents, id uint32)
	SendMagnifyEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)
	SendMouseEventsPointingDeviceIndex(events VZOpaqueMouseEvents, index uint32)
	SendMultiTouchEventsMultiTouchDeviceIndex(events unsafe.Pointer, index uint32)
	SendPointerNSEventPointingDeviceIndex(nSEvent objectivec.IObject, index uint32)
	SendQuickLookEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)
	SendRotationEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)
	SendScrollWheelEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)
	SendSmartMagnifyEventsPointingDeviceIndex(events unsafe.Pointer, index uint32)
}

// Init initializes the instance.
func (v VZVirtualMachineAccessor) Init() VZVirtualMachineAccessor {
	rv := objc.Send[VZVirtualMachineAccessor](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtualMachineAccessor) Autorelease() VZVirtualMachineAccessor {
	rv := objc.Send[VZVirtualMachineAccessor](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachineAccessor creates a new VZVirtualMachineAccessor instance.
func NewVZVirtualMachineAccessor() VZVirtualMachineAccessor {
	class := getVZVirtualMachineAccessorClass()
	rv := objc.Send[VZVirtualMachineAccessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineAccessor/_processHIDReports:forDevice:deviceType:
func (v VZVirtualMachineAccessor) _processHIDReportsForDeviceDeviceType(hIDReports VZOpaqueHIDReports, device uint32, type_ int32) {
	objc.Send[objc.ID](v.ID, objc.Sel("_processHIDReports:forDevice:deviceType:"), hIDReports.UnsafePointer(), device, type_)
}

// ProcessHIDReportsForDeviceDeviceType is an exported wrapper for the private method _processHIDReportsForDeviceDeviceType.
func (v VZVirtualMachineAccessor) ProcessHIDReportsForDeviceDeviceType(hIDReports VZOpaqueHIDReports, device uint32, type_ int32) {
	v._processHIDReportsForDeviceDeviceType(hIDReports, device, type_)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineAccessor/_shouldSendHIDReports
func (v VZVirtualMachineAccessor) _shouldSendHIDReports() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("_shouldSendHIDReports"))
	return rv
}

// ShouldSendHIDReports is an exported wrapper for the private method _shouldSendHIDReports.
func (v VZVirtualMachineAccessor) ShouldSendHIDReports() bool {
	return v._shouldSendHIDReports()
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineAccessor/sendDigitizerEvents:pointingDeviceIndex:
func (v VZVirtualMachineAccessor) SendDigitizerEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendDigitizerEvents:pointingDeviceIndex:"), events, index)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineAccessor/sendIOHIDEvents:hidDeviceIndex:
func (v VZVirtualMachineAccessor) SendIOHIDEventsHidDeviceIndex(iOHIDEvents VZOpaqueIOHIDEvents, index uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendIOHIDEvents:hidDeviceIndex:"), iOHIDEvents.UnsafePointer(), index)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineAccessor/sendKeyboardEvents:keyboardID:
func (v VZVirtualMachineAccessor) SendKeyboardEventsKeyboardID(events VZOpaqueKeyboardEvents, id uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendKeyboardEvents:keyboardID:"), events.UnsafePointer(), id)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineAccessor/sendMagnifyEvents:pointingDeviceIndex:
func (v VZVirtualMachineAccessor) SendMagnifyEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendMagnifyEvents:pointingDeviceIndex:"), events, index)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineAccessor/sendMouseEvents:pointingDeviceIndex:
func (v VZVirtualMachineAccessor) SendMouseEventsPointingDeviceIndex(events VZOpaqueMouseEvents, index uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendMouseEvents:pointingDeviceIndex:"), events.UnsafePointer(), index)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineAccessor/sendMultiTouchEvents:multiTouchDeviceIndex:
func (v VZVirtualMachineAccessor) SendMultiTouchEventsMultiTouchDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendMultiTouchEvents:multiTouchDeviceIndex:"), events, index)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineAccessor/sendPointerNSEvent:pointingDeviceIndex:
func (v VZVirtualMachineAccessor) SendPointerNSEventPointingDeviceIndex(nSEvent objectivec.IObject, index uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendPointerNSEvent:pointingDeviceIndex:"), nSEvent, index)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineAccessor/sendQuickLookEvents:pointingDeviceIndex:
func (v VZVirtualMachineAccessor) SendQuickLookEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendQuickLookEvents:pointingDeviceIndex:"), events, index)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineAccessor/sendRotationEvents:pointingDeviceIndex:
func (v VZVirtualMachineAccessor) SendRotationEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendRotationEvents:pointingDeviceIndex:"), events, index)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineAccessor/sendScrollWheelEvents:pointingDeviceIndex:
func (v VZVirtualMachineAccessor) SendScrollWheelEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendScrollWheelEvents:pointingDeviceIndex:"), events, index)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineAccessor/sendSmartMagnifyEvents:pointingDeviceIndex:
func (v VZVirtualMachineAccessor) SendSmartMagnifyEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.Send[objc.ID](v.ID, objc.Sel("sendSmartMagnifyEvents:pointingDeviceIndex:"), events, index)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineAccessor/_hidEventMonitor
func (v VZVirtualMachineAccessor) _hidEventMonitor() IVZHIDEventMonitor {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_hidEventMonitor"))
	return VZHIDEventMonitorFromID(objc.ID(rv))
}
