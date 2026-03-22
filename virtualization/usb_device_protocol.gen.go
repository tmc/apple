// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A protocol that represents a USB device in a VM.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBDevice
type VZUSBDevice interface {
	objectivec.IObject

	// The USB controller that has an attachment to the device.
	//
	// See: https://developer.apple.com/documentation/Virtualization/VZUSBDevice/usbController
	UsbController() IVZUSBController

	// The device’s unique identifier.
	//
	// See: https://developer.apple.com/documentation/Virtualization/VZUSBDevice/uuid
	Uuid() foundation.NSUUID
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

// The USB controller that has an attachment to the device.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBDevice/usbController
func (o VZUSBDeviceObject) UsbController() IVZUSBController {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("usbController"))
	return VZUSBControllerFromID(rv)
	}
// The device’s unique identifier.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBDevice/uuid
func (o VZUSBDeviceObject) Uuid() foundation.NSUUID {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("uuid"))
	return foundation.NSUUIDFromID(rv)
	}

