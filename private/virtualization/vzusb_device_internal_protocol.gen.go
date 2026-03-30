// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _VZUSBDeviceInternal protocol.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBDeviceInternal
type VZUSBDeviceInternal interface {
	objectivec.IObject

	// IsPointingDevice protocol.
	//
	// See: https://developer.apple.com/documentation/Virtualization/_VZUSBDeviceInternal/isPointingDevice
	IsPointingDevice() bool
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

// See: https://developer.apple.com/documentation/Virtualization/_VZUSBDeviceInternal/configuration
func (o VZUSBDeviceInternalObject) Configuration() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("configuration"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUSBDeviceInternal/isPointingDevice
func (o VZUSBDeviceInternalObject) IsPointingDevice() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isPointingDevice"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUSBDeviceInternal/setUsbController:
func (o VZUSBDeviceInternalObject) SetUsbController(controller objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setUsbController:"), controller)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUSBDeviceInternal/setVirtualMachine:
func (o VZUSBDeviceInternalObject) SetVirtualMachine(machine objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setVirtualMachine:"), machine)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUSBDeviceInternal/usbController
func (o VZUSBDeviceInternalObject) UsbController() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("usbController"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUSBDeviceInternal/virtualMachine
func (o VZUSBDeviceInternalObject) VirtualMachine() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("virtualMachine"))
	return objectivec.Object{ID: rv}
}
