// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _VZVirtualMachineConfigurationEncodable protocol.
type VZVirtualMachineConfigurationEncodable interface {
	objectivec.IObject

	// EncodeWithEncoder protocol.
	EncodeWithEncoder(encoder objectivec.IObject) unsafe.Pointer
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

func (o VZVirtualMachineConfigurationEncodableObject) EncodeWithEncoder(encoder objectivec.IObject) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](o.ID, objc.Sel("encodeWithEncoder:"), encoder)
	return rv
}
