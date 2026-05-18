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
func (v VZUSBController) Init() VZUSBController {
	rv := objc.Send[VZUSBController](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZUSBController) Autorelease() VZUSBController {
	rv := objc.Send[VZUSBController](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBController creates a new VZUSBController instance.
func NewVZUSBController() VZUSBController {
	class := getVZUSBControllerClass()
	rv := objc.Send[VZUSBController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBController/_capturePassthroughDevicesWithCompletionHandler:
func (v VZUSBController) _capturePassthroughDevicesWithCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](v.ID, objc.Sel("_capturePassthroughDevicesWithCompletionHandler:"), _block0)
}

// CapturePassthroughDevicesWithCompletionHandler is an exported wrapper for the private method _capturePassthroughDevicesWithCompletionHandler.
func (v VZUSBController) CapturePassthroughDevicesWithCompletionHandler(handler ErrorHandler) error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_capturePassthroughDevicesWithCompletionHandler:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_capturePassthroughDevicesWithCompletionHandler:"}
		return err
	}
	v._capturePassthroughDevicesWithCompletionHandler(handler)
	return nil
}

// CanCapturePassthroughDevicesWithCompletionHandler reports whether the receiver responds to the private selector _capturePassthroughDevicesWithCompletionHandler:.
func (v VZUSBController) CanCapturePassthroughDevicesWithCompletionHandler() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_capturePassthroughDevicesWithCompletionHandler:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBController/_initWithVirtualMachine:usbControllerIndex:usbDevices:
func (v VZUSBController) _initWithVirtualMachineUsbControllerIndexUsbDevices(machine objectivec.IObject, index uint64, devices objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("_initWithVirtualMachine:usbControllerIndex:usbDevices:"), machine, index, devices)
	return objectivec.Object{ID: rv}
}

// InitWithVirtualMachineUsbControllerIndexUsbDevices is an exported wrapper for the private method _initWithVirtualMachineUsbControllerIndexUsbDevices.
func (v VZUSBController) InitWithVirtualMachineUsbControllerIndexUsbDevices(machine objectivec.IObject, index uint64, devices objectivec.IObject) (objectivec.IObject, error) {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_initWithVirtualMachine:usbControllerIndex:usbDevices:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_initWithVirtualMachine:usbControllerIndex:usbDevices:"}
		return nil, err
	}
	return v._initWithVirtualMachineUsbControllerIndexUsbDevices(machine, index, devices), nil
}

// CanInitWithVirtualMachineUsbControllerIndexUsbDevices reports whether the receiver responds to the private selector _initWithVirtualMachine:usbControllerIndex:usbDevices:.
func (v VZUSBController) CanInitWithVirtualMachineUsbControllerIndexUsbDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_initWithVirtualMachine:usbControllerIndex:usbDevices:"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBController/_releasePassthroughDevices
func (v VZUSBController) _releasePassthroughDevices() {
	objc.Send[objc.ID](v.ID, objc.Sel("_releasePassthroughDevices"))
}

// ReleasePassthroughDevices is an exported wrapper for the private method _releasePassthroughDevices.
func (v VZUSBController) ReleasePassthroughDevices() error {
	if !objc.RespondsToSelector(v.ID, objc.Sel("_releasePassthroughDevices")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_releasePassthroughDevices"}
		return err
	}
	v._releasePassthroughDevices()
	return nil
}

// CanReleasePassthroughDevices reports whether the receiver responds to the private selector _releasePassthroughDevices.
func (v VZUSBController) CanReleasePassthroughDevices() bool {
	return objc.RespondsToSelector(v.ID, objc.Sel("_releasePassthroughDevices"))
}

// See: https://developer.apple.com/documentation/Virtualization/VZUSBController/delegate
func (v VZUSBController) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}
func (v VZUSBController) SetDelegate(value objectivec.IObject) {
	objc.Send[struct{}](v.ID, objc.Sel("setDelegate:"), value)
}
