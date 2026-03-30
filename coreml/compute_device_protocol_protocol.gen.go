// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An interface that represents a compute device type.
//
// See: https://developer.apple.com/documentation/CoreML/MLComputeDeviceProtocol
type MLComputeDeviceProtocol interface {
	objectivec.IObject
}

// MLComputeDeviceProtocolObject wraps an existing Objective-C object that conforms to the MLComputeDeviceProtocol protocol.
type MLComputeDeviceProtocolObject struct {
	objectivec.Object
}

func (o MLComputeDeviceProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// MLComputeDeviceProtocolObjectFromID constructs a [MLComputeDeviceProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MLComputeDeviceProtocolObjectFromID(id objc.ID) MLComputeDeviceProtocolObject {
	return MLComputeDeviceProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}
