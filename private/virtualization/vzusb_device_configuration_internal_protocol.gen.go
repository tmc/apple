// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _VZUSBDeviceConfigurationInternal protocol.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBDeviceConfigurationInternal
type VZUSBDeviceConfigurationInternal interface {
	objectivec.IObject
}

// VZUSBDeviceConfigurationInternalObject wraps an existing Objective-C object that conforms to the VZUSBDeviceConfigurationInternal protocol.
type VZUSBDeviceConfigurationInternalObject struct {
	objectivec.Object
}

func (o VZUSBDeviceConfigurationInternalObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZUSBDeviceConfigurationInternalObjectFromID constructs a [VZUSBDeviceConfigurationInternalObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZUSBDeviceConfigurationInternalObjectFromID(id objc.ID) VZUSBDeviceConfigurationInternalObject {
	return VZUSBDeviceConfigurationInternalObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUSBDeviceConfigurationInternal/isDuplicateConfiguration:
func (o VZUSBDeviceConfigurationInternalObject) IsDuplicateConfiguration(configuration objectivec.IObject) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isDuplicateConfiguration:"), configuration)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUSBDeviceConfigurationInternal/makeUSBDeviceWithVirtualMachine:
func (o VZUSBDeviceConfigurationInternalObject) MakeUSBDeviceWithVirtualMachine(machine objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("makeUSBDeviceWithVirtualMachine:"), machine)
	return objectivec.Object{ID: rv}
}
