// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MLComputeDeviceRegistryProtocol protocol.
type MLComputeDeviceRegistryProtocol interface {
	objectivec.IObject

	// RegisteredComputeDevices protocol.
	RegisteredComputeDevices() objectivec.IObject

	// SharedRegistry protocol.
	SharedRegistry() objectivec.IObject
}

// MLComputeDeviceRegistryProtocolObject wraps an existing Objective-C object that conforms to the MLComputeDeviceRegistryProtocol protocol.
type MLComputeDeviceRegistryProtocolObject struct {
	objectivec.Object
}

func (o MLComputeDeviceRegistryProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLComputeDeviceRegistryProtocolObjectFromID constructs a [MLComputeDeviceRegistryProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLComputeDeviceRegistryProtocolObjectFromID(id objc.ID) MLComputeDeviceRegistryProtocolObject {
	return MLComputeDeviceRegistryProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o MLComputeDeviceRegistryProtocolObject) RegisteredComputeDevices() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("registeredComputeDevices"))
	return objectivec.Object{ID: rv}
}
func (o MLComputeDeviceRegistryProtocolObject) SharedRegistry() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](o.ID, objc.Sel("sharedRegistry"))
	return objectivec.Object{ID: rv}
}
