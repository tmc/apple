// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// VZUSBDevice protocol.
type VZUSBDevice interface {
	objectivec.IObject
}

// VZUSBDeviceObject wraps an existing Objective-C object that conforms to the VZUSBDevice protocol.
type VZUSBDeviceObject struct {
	objectivec.Object
}

func (o VZUSBDeviceObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZUSBDeviceObjectFromID constructs a [VZUSBDeviceObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZUSBDeviceObjectFromID(id objc.ID) VZUSBDeviceObject {
	return VZUSBDeviceObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o VZUSBDeviceObject) UsbController() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("usbController"))
	return objectivec.Object{ID: rv}
}
func (o VZUSBDeviceObject) Uuid() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("uuid"))
	return objectivec.Object{ID: rv}
}
