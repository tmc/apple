// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _VZVirtualMachineConfigurationEncodable protocol.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineConfigurationEncodable
type VZVirtualMachineConfigurationEncodable interface {
	objectivec.IObject
}

// VZVirtualMachineConfigurationEncodableObject wraps an existing Objective-C object that conforms to the VZVirtualMachineConfigurationEncodable protocol.
type VZVirtualMachineConfigurationEncodableObject struct {
	objectivec.Object
}
func (o VZVirtualMachineConfigurationEncodableObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZVirtualMachineConfigurationEncodableObjectFromID constructs a [VZVirtualMachineConfigurationEncodableObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZVirtualMachineConfigurationEncodableObjectFromID(id objc.ID) VZVirtualMachineConfigurationEncodableObject {
	return VZVirtualMachineConfigurationEncodableObject{
		Object: objectivec.ObjectFromID(id),
	}
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZVirtualMachineConfigurationEncodable/encodeWithEncoder:
func (o VZVirtualMachineConfigurationEncodableObject) EncodeWithEncoder(encoder objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return objectivec.Object{ID: rv}
	}

