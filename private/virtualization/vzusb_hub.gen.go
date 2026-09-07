// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZUSBHub] class.
var (
	_VZUSBHubClass     VZUSBHubClass
	_VZUSBHubClassOnce sync.Once
)

func getVZUSBHubClass() VZUSBHubClass {
	_VZUSBHubClassOnce.Do(func() {
		_VZUSBHubClass = VZUSBHubClass{class: objc.GetClass("_VZUSBHub")}
	})
	return _VZUSBHubClass
}

// GetVZUSBHubClass returns the class object for _VZUSBHub.
func GetVZUSBHubClass() VZUSBHubClass {
	return getVZUSBHubClass()
}

type VZUSBHubClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZUSBHubClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZUSBHubClass) Alloc() VZUSBHub {
	rv := objc.SendIfResponds[VZUSBHub](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZUSBHub.Configuration]
//   - [VZUSBHub.IsPointingDevice]
//   - [VZUSBHub.UsbController]
//   - [VZUSBHub.SetUsbController]
//   - [VZUSBHub.UsbDeviceDidConnectOnPort]
//   - [VZUSBHub.UsbDeviceDidDisconnectOnPort]
//   - [VZUSBHub.Uuid]
//   - [VZUSBHub.VirtualMachine]
//   - [VZUSBHub.SetVirtualMachine]
//   - [VZUSBHub.DebugDescription]
//   - [VZUSBHub.Description]
//   - [VZUSBHub.Hash]
//   - [VZUSBHub.Superclass]
type VZUSBHub struct {
	objectivec.Object
}

// VZUSBHubFromID constructs a [VZUSBHub] from an objc.ID.
func VZUSBHubFromID(id objc.ID) VZUSBHub {
	return VZUSBHub{objectivec.Object{ID: id}}
}

// Ensure VZUSBHub implements IVZUSBHub.
var _ IVZUSBHub = VZUSBHub{}

// An interface definition for the [VZUSBHub] class.
//
// # Methods
//
//   - [IVZUSBHub.Configuration]
//   - [IVZUSBHub.IsPointingDevice]
//   - [IVZUSBHub.UsbController]
//   - [IVZUSBHub.SetUsbController]
//   - [IVZUSBHub.UsbDeviceDidConnectOnPort]
//   - [IVZUSBHub.UsbDeviceDidDisconnectOnPort]
//   - [IVZUSBHub.Uuid]
//   - [IVZUSBHub.VirtualMachine]
//   - [IVZUSBHub.SetVirtualMachine]
//   - [IVZUSBHub.DebugDescription]
//   - [IVZUSBHub.Description]
//   - [IVZUSBHub.Hash]
//   - [IVZUSBHub.Superclass]
type IVZUSBHub interface {
	objectivec.IObject

	// Topic: Methods

	Configuration() unsafe.Pointer
	IsPointingDevice() bool
	UsbController() IVZUSBController
	SetUsbController(value IVZUSBController)
	UsbDeviceDidConnectOnPort(device objectivec.IObject, port uint32)
	UsbDeviceDidDisconnectOnPort(port uint32)
	Uuid() foundation.NSUUID
	VirtualMachine() IVZVirtualMachine
	SetVirtualMachine(value IVZVirtualMachine)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (v VZUSBHub) Init() VZUSBHub {
	rv := objc.SendIfResponds[VZUSBHub](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZUSBHub) Autorelease() VZUSBHub {
	rv := objc.SendIfResponds[VZUSBHub](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZUSBHub creates a new VZUSBHub instance.
func NewVZUSBHub() VZUSBHub {
	class := getVZUSBHubClass()
	rv := objc.SendIfResponds[VZUSBHub](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZUSBHub) UsbDeviceDidConnectOnPort(device objectivec.IObject, port uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("usbDevice:didConnectOnPort:"), device, port)
}
func (v VZUSBHub) UsbDeviceDidDisconnectOnPort(port uint32) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("usbDeviceDidDisconnectOnPort:"), port)
}

func (v VZUSBHub) Configuration() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](v.ID, objc.Sel("configuration"))
	return rv
}
func (v VZUSBHub) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZUSBHub) Description() string {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (v VZUSBHub) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](v.ID, objc.Sel("hash"))
	return rv
}
func (v VZUSBHub) IsPointingDevice() bool {
	rv := objc.SendIfResponds[bool](v.ID, objc.Sel("isPointingDevice"))
	return rv
}
func (v VZUSBHub) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](v.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (v VZUSBHub) UsbController() IVZUSBController {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("usbController"))
	return VZUSBControllerFromID(objc.ID(rv))
}
func (v VZUSBHub) SetUsbController(value IVZUSBController) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setUsbController:"), value)
}
func (v VZUSBHub) Uuid() foundation.NSUUID {
	rv := objc.SendIfResponds[foundation.NSUUID](v.ID, objc.Sel("uuid"))
	return foundation.NSUUID(rv)
}
func (v VZUSBHub) VirtualMachine() IVZVirtualMachine {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("virtualMachine"))
	return VZVirtualMachineFromID(objc.ID(rv))
}
func (v VZUSBHub) SetVirtualMachine(value IVZVirtualMachine) {
	objc.SendIfResponds[struct{}](v.ID, objc.Sel("setVirtualMachine:"), value)
}
