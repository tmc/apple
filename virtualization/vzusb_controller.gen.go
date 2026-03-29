// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"context"
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

// A class that represents a USB controller in a VM.
//
// # Overview
// 
// Don’t create a [VZUSBController] directly. You need to first configure
// USB controllers on a [VZVirtualMachineConfiguration] through a subclass of
// [VZUSBControllerConfiguration]. When you create a [VZVirtualMachine] from
// the configuration, the USB controllers are available through the
// [VZUSBController.UsbControllers] property.
// 
// The concrete type of a [VZUSBController] corresponds to the type the
// configuration uses. For example, a [VZXHCIControllerConfiguration] leads to
// a device of type [VZXHCIController].
//
// # Instance properties
//
//   - [VZUSBController.UsbDevices]: The list of attached USB devices for the controller.
//
// # Attaching and detaching devices
//
//   - [VZUSBController.AttachDeviceCompletionHandler]: Attaches a USB device to the controller.
//   - [VZUSBController.DetachDeviceCompletionHandler]: Detaches a USB device from the controller.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBController
type VZUSBController struct {
	objectivec.Object
}

// VZUSBControllerFromID constructs a [VZUSBController] from an objc.ID.
//
// A class that represents a USB controller in a VM.
func VZUSBControllerFromID(id objc.ID) VZUSBController {
	return VZUSBController{objectivec.Object{ID: id}}
}
// NOTE: VZUSBController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZUSBController] class.
//
// # Instance properties
//
//   - [IVZUSBController.UsbDevices]: The list of attached USB devices for the controller.
//
// # Attaching and detaching devices
//
//   - [IVZUSBController.AttachDeviceCompletionHandler]: Attaches a USB device to the controller.
//   - [IVZUSBController.DetachDeviceCompletionHandler]: Detaches a USB device from the controller.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBController
type IVZUSBController interface {
	objectivec.IObject

	// Topic: Instance properties

	// The list of attached USB devices for the controller.
	UsbDevices() []objectivec.IObject

	// Topic: Attaching and detaching devices

	// Attaches a USB device to the controller.
	AttachDeviceCompletionHandler(device VZUSBDevice, completionHandler ErrorHandler)
	// Detaches a USB device from the controller.
	DetachDeviceCompletionHandler(device VZUSBDevice, completionHandler ErrorHandler)

	// The list of runtime USB controller objects.
	UsbControllers() IVZUSBController
	SetUsbControllers(value IVZUSBController)
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

// Attaches a USB device to the controller.
//
// device: The USB device to attach.
//
// completionHandler: A block the framework calls after the device attaches, or on an error. The
// error parameter that passes to the block is `nil` if attaching is
// successful. The framework calls the block on a VM’s queue.
//
// # Discussion
// 
// If the device successfully attaches to the controller, it appears in the
// [UsbDevices] property, the framework sets its [UsbController] property to
// point to the attached USB controller, and the completion handler returns
// `nil`.
// 
// If the device has a previous attachment to the current USB controller, or
// to another USB controller, the attach function fails with
// [ErrorDeviceAlreadyAttached]. If the controller can’t initialize the
// device correctly, the attach function fails with
// [ErrorDeviceInitializationFailure].
// 
// You need to call this method on the virtual machine’s queue.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBController/attach(device:completionHandler:)
func (u VZUSBController) AttachDeviceCompletionHandler(device VZUSBDevice, completionHandler ErrorHandler) {
_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](u.ID, objc.Sel("attachDevice:completionHandler:"), device, _block1)
}
// Detaches a USB device from the controller.
//
// device: The USB device to detach.
//
// completionHandler: A block the framework calls after the device detaches, or on an error. The
// error parameter that passes to the block is `nil` if detaching the device
// is successful. The framework calls the block on a VM’s queue.
//
// # Discussion
// 
// If the device successfully detaches from the controller, it disappears from
// the [UsbDevices] property, the framework sets its [UsbController] property
// to `nil,` and the completion handler returns `nil`.
// 
// If the device doesn’t have an attachment to the controller at the time of
// calling the detach method, the process fails with [ErrorDeviceNotFound].
// 
// You need to call this method on the virtual machine’s queue.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBController/detach(device:completionHandler:)
func (u VZUSBController) DetachDeviceCompletionHandler(device VZUSBDevice, completionHandler ErrorHandler) {
_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](u.ID, objc.Sel("detachDevice:completionHandler:"), device, _block1)
}

// The list of attached USB devices for the controller.
//
// # Discussion
// 
// If a [VZVirtualMachineConfiguration] contains a USB controller
// configuration that contains USB devices, those devices appear in the list
// when you start the virtual machine.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBController/usbDevices
func (u VZUSBController) UsbDevices() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](u.ID, objc.Sel("usbDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
// The list of runtime USB controller objects.
//
// See: https://developer.apple.com/documentation/virtualization/vzvirtualmachine/usbcontrollers
func (u VZUSBController) UsbControllers() IVZUSBController {
	rv := objc.Send[objc.ID](u.ID, objc.Sel("usbControllers"))
	return VZUSBControllerFromID(objc.ID(rv))
}
func (u VZUSBController) SetUsbControllers(value IVZUSBController) {
	objc.Send[struct{}](u.ID, objc.Sel("setUsbControllers:"), value)
}

// AttachDevice is a synchronous wrapper around [VZUSBController.AttachDeviceCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u VZUSBController) AttachDevice(ctx context.Context, device VZUSBDevice) error {
	done := make(chan error, 1)
	u.AttachDeviceCompletionHandler(device, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// DetachDevice is a synchronous wrapper around [VZUSBController.DetachDeviceCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (u VZUSBController) DetachDevice(ctx context.Context, device VZUSBDevice) error {
	done := make(chan error, 1)
	u.DetachDeviceCompletionHandler(device, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

