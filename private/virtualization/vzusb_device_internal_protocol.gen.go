// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _VZUSBDeviceInternal protocol.
type VZUSBDeviceInternal interface {
	objectivec.IObject

	// Configuration protocol.
	Configuration() objectivec.IObject

	// IsPointingDevice protocol.
	IsPointingDevice() bool

	// SetUsbController protocol.
	SetUsbController(controller objectivec.IObject)

	// SetVirtualMachine protocol.
	SetVirtualMachine(machine objectivec.IObject)

	// UsbController protocol.
	UsbController() objectivec.IObject

	// VirtualMachine protocol.
	VirtualMachine() objectivec.IObject
}

// VZUSBDeviceInternalObject wraps an existing Objective-C object that conforms to the VZUSBDeviceInternal protocol.
type VZUSBDeviceInternalObject struct {
	objectivec.Object
}

func (o VZUSBDeviceInternalObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZUSBDeviceInternalObjectFromID constructs a [VZUSBDeviceInternalObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZUSBDeviceInternalObjectFromID(id objc.ID) VZUSBDeviceInternalObject {
	return VZUSBDeviceInternalObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o VZUSBDeviceInternalObject) Configuration() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("configuration"))
	return objectivec.Object{ID: rv}
}
func (o VZUSBDeviceInternalObject) IsPointingDevice() bool {
	rv := objc.SendIfResponds[bool](o.ID, objc.Sel("isPointingDevice"))
	return rv
}
func (o VZUSBDeviceInternalObject) SetUsbController(controller objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setUsbController:"), controller)
}
func (o VZUSBDeviceInternalObject) SetVirtualMachine(machine objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setVirtualMachine:"), machine)
}
func (o VZUSBDeviceInternalObject) UsbController() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("usbController"))
	return objectivec.Object{ID: rv}
}
func (o VZUSBDeviceInternalObject) VirtualMachine() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("virtualMachine"))
	return objectivec.Object{ID: rv}
}
