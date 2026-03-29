// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
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

// A class that represents a USB Extensible Host Controller Interface (XHCI)
// controller in a VM.
//
// # Overview
// 
// Don’t create [VZXHCIController] objects directly. Instead, you create a
// [VZXHCIController] object at runtime though the [VZXHCIController.UsbControllers] property
// of the [VZVirtualMachineConfiguration] object by populating it with
// [VZXHCIControllerConfiguration] objects.
//
// See: https://developer.apple.com/documentation/Virtualization/VZXHCIController
type VZXHCIController struct {
	VZUSBController
}

// VZXHCIControllerFromID constructs a [VZXHCIController] from an objc.ID.
//
// A class that represents a USB Extensible Host Controller Interface (XHCI)
// controller in a VM.
func VZXHCIControllerFromID(id objc.ID) VZXHCIController {
	return VZXHCIController{VZUSBController: VZUSBControllerFromID(id)}
}
// NOTE: VZXHCIController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [VZXHCIController] class.
//
// See: https://developer.apple.com/documentation/Virtualization/VZXHCIController
type IVZXHCIController interface {
	IVZUSBController
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

