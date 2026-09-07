// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// VZUSBDeviceConfiguration protocol.
type VZUSBDeviceConfiguration interface {
	objectivec.IObject

	// SetUuid protocol.
	SetUuid(uuid objectivec.IObject)

	// Uuid protocol.
	Uuid() objectivec.IObject
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

func (o VZUSBDeviceConfigurationObject) SetUuid(uuid objectivec.IObject) {
	objc.SendIfResponds[struct{}](o.ID, objc.Sel("setUuid:"), uuid)
}
func (o VZUSBDeviceConfigurationObject) Uuid() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("uuid"))
	return objectivec.Object{ID: rv}
}
