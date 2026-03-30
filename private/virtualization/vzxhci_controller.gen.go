// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZXHCIController] class.
var (
	_VZXHCIControllerClass     VZXHCIControllerClass
	_VZXHCIControllerClassOnce sync.Once
)

func getVZXHCIControllerClass() VZXHCIControllerClass {
	_VZXHCIControllerClassOnce.Do(func() {
		_VZXHCIControllerClass = VZXHCIControllerClass{class: objc.GetClass("VZXHCIController")}
	})
	return _VZXHCIControllerClass
}

// GetVZXHCIControllerClass returns the class object for VZXHCIController.
func GetVZXHCIControllerClass() VZXHCIControllerClass {
	return getVZXHCIControllerClass()
}

type VZXHCIControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZXHCIControllerClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZXHCIControllerClass) Alloc() VZXHCIController {
	rv := objc.Send[VZXHCIController](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZXHCIController.AttachDeviceCompletionHandler]
//   - [VZXHCIController.DetachDeviceCompletionHandler]
//
// See: https://developer.apple.com/documentation/Virtualization/VZXHCIController
type VZXHCIController struct {
	VZUSBController
}

// VZXHCIControllerFromID constructs a [VZXHCIController] from an objc.ID.
func VZXHCIControllerFromID(id objc.ID) VZXHCIController {
	return VZXHCIController{VZUSBController: VZUSBControllerFromID(id)}
}

// Ensure VZXHCIController implements IVZXHCIController.
var _ IVZXHCIController = VZXHCIController{}

// An interface definition for the [VZXHCIController] class.
//
// # Methods
//
//   - [IVZXHCIController.AttachDeviceCompletionHandler]
//   - [IVZXHCIController.DetachDeviceCompletionHandler]
//
// See: https://developer.apple.com/documentation/Virtualization/VZXHCIController
type IVZXHCIController interface {
	IVZUSBController

	// Topic: Methods

	AttachDeviceCompletionHandler(device objectivec.IObject, handler ErrorHandler)
	DetachDeviceCompletionHandler(device objectivec.IObject, handler ErrorHandler)
}

// Init initializes the instance.
func (x VZXHCIController) Init() VZXHCIController {
	rv := objc.Send[VZXHCIController](x.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (x VZXHCIController) Autorelease() VZXHCIController {
	rv := objc.Send[VZXHCIController](x.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZXHCIController creates a new VZXHCIController instance.
func NewVZXHCIController() VZXHCIController {
	class := getVZXHCIControllerClass()
	rv := objc.Send[VZXHCIController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/VZXHCIController/attachDevice:completionHandler:
func (x VZXHCIController) AttachDeviceCompletionHandler(device objectivec.IObject, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](x.ID, objc.Sel("attachDevice:completionHandler:"), device, _block1)
}

// See: https://developer.apple.com/documentation/Virtualization/VZXHCIController/detachDevice:completionHandler:
func (x VZXHCIController) DetachDeviceCompletionHandler(device objectivec.IObject, handler ErrorHandler) {
	_block1, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](x.ID, objc.Sel("detachDevice:completionHandler:"), device, _block1)
}

// AttachDevice is a synchronous wrapper around [VZXHCIController.AttachDeviceCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (x VZXHCIController) AttachDevice(ctx context.Context, device objectivec.IObject) error {
	done := make(chan error, 1)
	x.AttachDeviceCompletionHandler(device, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// DetachDevice is a synchronous wrapper around [VZXHCIController.DetachDeviceCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (x VZXHCIController) DetachDevice(ctx context.Context, device objectivec.IObject) error {
	done := make(chan error, 1)
	x.DetachDeviceCompletionHandler(device, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
