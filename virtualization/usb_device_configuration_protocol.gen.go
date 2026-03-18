// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The protocol for configuring USB devices.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBDeviceConfiguration
type VZUSBDeviceConfiguration interface {
	objectivec.IObject

	// The device’s unique identifier.
	//
	// See: https://developer.apple.com/documentation/Virtualization/VZUSBDeviceConfiguration/uuid
	Uuid() foundation.NSUUID

	// The device’s unique identifier.
	//
	// See: https://developer.apple.com/documentation/Virtualization/VZUSBDeviceConfiguration/uuid
	SetUuid(value foundation.NSUUID)
}

// VZUSBDeviceConfigurationObject wraps an existing Objective-C object that conforms to the VZUSBDeviceConfiguration protocol.
type VZUSBDeviceConfigurationObject struct {
	objectivec.Object
}
func (o VZUSBDeviceConfigurationObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZUSBDeviceConfigurationObjectFromID constructs a [VZUSBDeviceConfigurationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZUSBDeviceConfigurationObjectFromID(id objc.ID) VZUSBDeviceConfigurationObject {
	return VZUSBDeviceConfigurationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The device’s unique identifier.
//
// See: https://developer.apple.com/documentation/Virtualization/VZUSBDeviceConfiguration/uuid

func (o VZUSBDeviceConfigurationObject) Uuid() foundation.NSUUID {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("uuid"))
	return foundation.NSUUIDFromID(rv)
	}

func (o VZUSBDeviceConfigurationObject) SetUuid(value foundation.NSUUID) {
	objc.Send[struct{}](o.ID, objc.Sel("setUuid:"), value)
}

