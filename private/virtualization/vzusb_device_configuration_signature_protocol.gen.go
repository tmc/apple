// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// _VZUSBDeviceConfigurationSignature protocol.
//
// See: https://developer.apple.com/documentation/Virtualization/_VZUSBDeviceConfigurationSignature
type VZUSBDeviceConfigurationSignature interface {
	objectivec.IObject
}

// VZUSBDeviceConfigurationSignatureObject wraps an existing Objective-C object that conforms to the VZUSBDeviceConfigurationSignature protocol.
type VZUSBDeviceConfigurationSignatureObject struct {
	objectivec.Object
}
func (o VZUSBDeviceConfigurationSignatureObject) BaseObject() objectivec.Object {
	return o.Object
}

// VZUSBDeviceConfigurationSignatureObjectFromID constructs a [VZUSBDeviceConfigurationSignatureObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func VZUSBDeviceConfigurationSignatureObjectFromID(id objc.ID) VZUSBDeviceConfigurationSignatureObject {
	return VZUSBDeviceConfigurationSignatureObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/Virtualization/_VZUSBDeviceConfigurationSignature/signature
func (o VZUSBDeviceConfigurationSignatureObject) Signature() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("signature"))
	return objectivec.Object{ID: rv}
	}

