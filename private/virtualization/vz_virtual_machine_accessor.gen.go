// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
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
	rv := objc.SendIfResponds[VZVirtualMachineAccessor](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZVirtualMachineAccessor._hidEventMonitor]
//   - [VZVirtualMachineAccessor._processHIDReportsForDeviceDeviceType]
//   - [VZVirtualMachineAccessor._shouldSendHIDReports]
//   - [VZVirtualMachineAccessor.GraphicsDevices]
//   - [VZVirtualMachineAccessor.Queue]
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
//   - [VZVirtualMachineAccessor.InitWithAccessorEndpoint]
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
//   - [IVZVirtualMachineAccessor.GraphicsDevices]
//   - [IVZVirtualMachineAccessor.Queue]
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
//   - [IVZVirtualMachineAccessor.InitWithAccessorEndpoint]
type IVZVirtualMachineAccessor interface {
	objectivec.IObject

	// Topic: Methods

	_hidEventMonitor() IVZHIDEventMonitor
	_processHIDReportsForDeviceDeviceType(hIDReports VZOpaqueHIDReports, device uint32, type_ int32)
	_shouldSendHIDReports() bool
	GraphicsDevices() foundation.INSArray
	Queue() objectivec.Object
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
	InitWithAccessorEndpoint(endpoint objectivec.IObject) VZVirtualMachineAccessor
}

// Init initializes the instance.
func (v VZVirtualMachineAccessor) Init() VZVirtualMachineAccessor {
	rv := objc.SendIfResponds[VZVirtualMachineAccessor](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZVirtualMachineAccessor) Autorelease() VZVirtualMachineAccessor {
	rv := objc.SendIfResponds[VZVirtualMachineAccessor](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZVirtualMachineAccessor creates a new VZVirtualMachineAccessor instance.
func NewVZVirtualMachineAccessor() VZVirtualMachineAccessor {
	class := getVZVirtualMachineAccessorClass()
	rv := objc.SendIfResponds[VZVirtualMachineAccessor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZVirtualMachineAccessorWithAccessorEndpoint(endpoint objectivec.IObject) VZVirtualMachineAccessor {
	instance := getVZVirtualMachineAccessorClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithAccessorEndpoint:"), endpoint)
	return VZVirtualMachineAccessorFromID(rv)
}

func (v VZVirtualMachineAccessor) _processHIDReportsForDeviceDeviceType(hIDReports VZOpaqueHIDReports, device uint32, type_ int32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_processHIDReports:forDevice:deviceType:"), hIDReports.UnsafePointer(), device, type_)
}

// ProcessHIDReportsForDeviceDeviceType is an exported wrapper for the private method _processHIDReportsForDeviceDeviceType.
func (v VZVirtualMachineAccessor) ProcessHIDReportsForDeviceDeviceType(hIDReports VZOpaqueHIDReports, device uint32, type_ int32) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_processHIDReports:forDevice:deviceType:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_processHIDReports:forDevice:deviceType:"}
		return err
	}
	v._processHIDReportsForDeviceDeviceType(hIDReports, device, type_)
	return nil
}

// CanProcessHIDReportsForDeviceDeviceType reports whether the receiver responds to the private selector _processHIDReports:forDevice:deviceType:.
func (v VZVirtualMachineAccessor) CanProcessHIDReportsForDeviceDeviceType() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_processHIDReports:forDevice:deviceType:"))
}
func (v VZVirtualMachineAccessor) _shouldSendHIDReports() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("_shouldSendHIDReports"))
	return rv
}

// ShouldSendHIDReports is an exported wrapper for the private method _shouldSendHIDReports.
func (v VZVirtualMachineAccessor) ShouldSendHIDReports() (bool, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_shouldSendHIDReports")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_shouldSendHIDReports"}
		return false, err
	}
	return v._shouldSendHIDReports(), nil
}

// CanShouldSendHIDReports reports whether the receiver responds to the private selector _shouldSendHIDReports.
func (v VZVirtualMachineAccessor) CanShouldSendHIDReports() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_shouldSendHIDReports"))
}
func (v VZVirtualMachineAccessor) SendDigitizerEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendDigitizerEvents:pointingDeviceIndex:"), events, index)
}
func (v VZVirtualMachineAccessor) SendIOHIDEventsHidDeviceIndex(iOHIDEvents VZOpaqueIOHIDEvents, index uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendIOHIDEvents:hidDeviceIndex:"), iOHIDEvents.UnsafePointer(), index)
}
func (v VZVirtualMachineAccessor) SendKeyboardEventsKeyboardID(events VZOpaqueKeyboardEvents, id uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendKeyboardEvents:keyboardID:"), events.UnsafePointer(), id)
}
func (v VZVirtualMachineAccessor) SendMagnifyEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendMagnifyEvents:pointingDeviceIndex:"), events, index)
}
func (v VZVirtualMachineAccessor) SendMouseEventsPointingDeviceIndex(events VZOpaqueMouseEvents, index uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendMouseEvents:pointingDeviceIndex:"), events.UnsafePointer(), index)
}
func (v VZVirtualMachineAccessor) SendMultiTouchEventsMultiTouchDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendMultiTouchEvents:multiTouchDeviceIndex:"), events, index)
}
func (v VZVirtualMachineAccessor) SendPointerNSEventPointingDeviceIndex(nSEvent objectivec.IObject, index uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendPointerNSEvent:pointingDeviceIndex:"), nSEvent, index)
}
func (v VZVirtualMachineAccessor) SendQuickLookEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendQuickLookEvents:pointingDeviceIndex:"), events, index)
}
func (v VZVirtualMachineAccessor) SendRotationEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendRotationEvents:pointingDeviceIndex:"), events, index)
}
func (v VZVirtualMachineAccessor) SendScrollWheelEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendScrollWheelEvents:pointingDeviceIndex:"), events, index)
}
func (v VZVirtualMachineAccessor) SendSmartMagnifyEventsPointingDeviceIndex(events unsafe.Pointer, index uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("sendSmartMagnifyEvents:pointingDeviceIndex:"), events, index)
}
func (v VZVirtualMachineAccessor) InitWithAccessorEndpoint(endpoint objectivec.IObject) VZVirtualMachineAccessor {
	rv := objc.SendIfResponds[VZVirtualMachineAccessor](v.ID, objc.Sel("initWithAccessorEndpoint:"), endpoint)
	return rv
}

func (v VZVirtualMachineAccessor) _hidEventMonitor() IVZHIDEventMonitor {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_hidEventMonitor"))
	return VZHIDEventMonitorFromID(objc.ID(rv))
}

// CanHidEventMonitor reports whether the receiver responds to the private selector _hidEventMonitor.
func (v VZVirtualMachineAccessor) CanHidEventMonitor() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_hidEventMonitor"))
}

// HidEventMonitor is an exported wrapper for the private property _hidEventMonitor.
func (v VZVirtualMachineAccessor) HidEventMonitor() (IVZHIDEventMonitor, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_hidEventMonitor")) {
		return nil, &objc.UnrecognizedSelectorError{Selector: "_hidEventMonitor"}
	}
	return v._hidEventMonitor(), nil
}
func (v VZVirtualMachineAccessor) GraphicsDevices() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("graphicsDevices"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (v VZVirtualMachineAccessor) Queue() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("queue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
