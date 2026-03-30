// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZUSBController] class.
var (
	_VZUSBControllerClass     VZUSBControllerClass
	_VZUSBControllerClassOnce sync.Once
)

func getVZUSBControllerClass() VZUSBControllerClass {
	_VZUSBControllerClassOnce.Do(func() {
		_VZUSBControllerClass = VZUSBControllerClass{class: objc.GetClass("VZUSBController")}
	})
	return _VZUSBControllerClass
}

// GetVZUSBControllerClass returns the class object for VZUSBController.
func GetVZUSBControllerClass() VZUSBControllerClass {
	return getVZUSBControllerClass()
}

type VZUSBControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZUSBControllerClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZUSBControllerClass) Alloc() VZUSBController {
	rv := objc.Send[VZUSBController](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZUSBController._capturePassthroughDevicesWithCompletionHandler]
//   - [VZUSBController._initWithVirtualMachineUsbControllerIndexUsbDevices]
//   - [VZUSBController._releasePassthroughDevices]
//   - [VZUSBController.Delegate]
//   - [VZUSBController.SetDelegate]
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBController
type VZUSBController struct {
	objectivec.Object
}

// VZUSBControllerFromID constructs a [VZUSBController] from an objc.ID.
func VZUSBControllerFromID(id objc.ID) VZUSBController {
	return VZUSBController{objectivec.Object{ID: id}}
}

// Ensure VZUSBController implements IVZUSBController.
var _ IVZUSBController = VZUSBController{}

// An interface definition for the [VZUSBController] class.
//
// # Methods
//
//   - [IVZUSBController._capturePassthroughDevicesWithCompletionHandler]
//   - [IVZUSBController._initWithVirtualMachineUsbControllerIndexUsbDevices]
//   - [IVZUSBController._releasePassthroughDevices]
//   - [IVZUSBController.Delegate]
//   - [IVZUSBController.SetDelegate]
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBController
type IVZUSBController interface {
	objectivec.IObject

	// Topic: Methods

	_capturePassthroughDevicesWithCompletionHandler(handler ErrorHandler)
	_initWithVirtualMachineUsbControllerIndexUsbDevices(machine objectivec.IObject, index uint64, devices objectivec.IObject) objectivec.IObject
	_releasePassthroughDevices()
	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)
}

// Init initializes the instance.
func (u VZUSBController) Init() VZUSBController {
	rv := objc.Send[VZUSBController](u.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (u VZUSBController) Autorelease() VZUSBController {
	rv := objc.Send[VZUSBController](u.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBController creates a new VZUSBController instance.
func NewVZUSBController() VZUSBController {
	class := getVZUSBControllerClass()
	rv := objc.Send[VZUSBController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBController/_capturePassthroughDevicesWithCompletionHandler:
func (u VZUSBController) _capturePassthroughDevicesWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](u.ID, objc.Sel("_capturePassthroughDevicesWithCompletionHandler:"), _block0)
}

// CapturePassthroughDevicesWithCompletionHandler is an exported wrapper for the private method _capturePassthroughDevicesWithCompletionHandler.
func (u VZUSBController) CapturePassthroughDevicesWithCompletionHandler(handler ErrorHandler) {
	u._capturePassthroughDevicesWithCompletionHandler(handler)
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBController/_initWithVirtualMachine:usbControllerIndex:usbDevices:
func (u VZUSBController) _initWithVirtualMachineUsbControllerIndexUsbDevices(machine objectivec.IObject, index uint64, devices objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("_initWithVirtualMachine:usbControllerIndex:usbDevices:"), machine, index, devices)
	return objectivec.Object{ID: rv}
}

// InitWithVirtualMachineUsbControllerIndexUsbDevices is an exported wrapper for the private method _initWithVirtualMachineUsbControllerIndexUsbDevices.
func (u VZUSBController) InitWithVirtualMachineUsbControllerIndexUsbDevices(machine objectivec.IObject, index uint64, devices objectivec.IObject) objectivec.IObject {
	return u._initWithVirtualMachineUsbControllerIndexUsbDevices(machine, index, devices)
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBController/_releasePassthroughDevices
func (u VZUSBController) _releasePassthroughDevices() {
	objc.Send[objc.ID](u.ID, objc.Sel("_releasePassthroughDevices"))
}

// ReleasePassthroughDevices is an exported wrapper for the private method _releasePassthroughDevices.
func (u VZUSBController) ReleasePassthroughDevices() {
	u._releasePassthroughDevices()
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBController/delegate
func (u VZUSBController) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (u VZUSBController) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](u.ID, objc.Sel("setDelegate:"), value)
}
