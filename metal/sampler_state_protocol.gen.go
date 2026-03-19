// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// An instance that defines how a texture should be sampled.
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerState
type MTLSamplerState interface {
	objectivec.IObject

	// The device object that created the sampler.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLSamplerState/device
	Device() MTLDevice

	// A string that identifies the sampler.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLSamplerState/label
	Label() string

	// GpuResourceID protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLSamplerState/gpuResourceID
	GpuResourceID() MTLResourceID
}

// MTLSamplerStateObject wraps an existing Objective-C object that conforms to the MTLSamplerState protocol.
type MTLSamplerStateObject struct {
	objectivec.Object
}
func (o MTLSamplerStateObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLSamplerStateObjectFromID constructs a [MTLSamplerStateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLSamplerStateObjectFromID(id objc.ID) MTLSamplerStateObject {
	return MTLSamplerStateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The device object that created the sampler.
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerState/device
func (o MTLSamplerStateObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}
// A string that identifies the sampler.
//
// See: https://developer.apple.com/documentation/Metal/MTLSamplerState/label
func (o MTLSamplerStateObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}
// See: https://developer.apple.com/documentation/Metal/MTLSamplerState/gpuResourceID
func (o MTLSamplerStateObject) GpuResourceID() MTLResourceID {
	
	rv := objc.Send[MTLResourceID](o.ID, objc.Sel("gpuResourceID"))
	return rv
	}

