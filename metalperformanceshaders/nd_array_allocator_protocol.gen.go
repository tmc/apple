// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MPSNDArrayAllocator protocol.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAllocator
type MPSNDArrayAllocator interface {
	objectivec.IObject
	foundation.NSCoding
	foundation.NSCopying
	foundation.NSSecureCoding

	// ArrayForCommandBufferArrayDescriptorKernel protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAllocator/array(for:arrayDescriptor:kernel:)
	ArrayForCommandBufferArrayDescriptorKernel(cmdBuf metal.MTLCommandBuffer, descriptor IMPSNDArrayDescriptor, kernel IMPSKernel) IMPSNDArray
}

// MPSNDArrayAllocatorObject wraps an existing Objective-C object that conforms to the MPSNDArrayAllocator protocol.
type MPSNDArrayAllocatorObject struct {
	foundation.NSCodingObject
}

func (o MPSNDArrayAllocatorObject) BaseObject() objectivec.Object {
	return o.NSCodingObject.BaseObject()
}

// MPSNDArrayAllocatorObjectFromID constructs a [MPSNDArrayAllocatorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MPSNDArrayAllocatorObjectFromID(id objc.ID) MPSNDArrayAllocatorObject {
	return MPSNDArrayAllocatorObject{
		NSCodingObject: foundation.NSCodingObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSNDArrayAllocator/array(for:arrayDescriptor:kernel:)
func (o MPSNDArrayAllocatorObject) ArrayForCommandBufferArrayDescriptorKernel(cmdBuf metal.MTLCommandBuffer, descriptor IMPSNDArrayDescriptor, kernel IMPSKernel) IMPSNDArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("arrayForCommandBuffer:arrayDescriptor:kernel:"), cmdBuf, descriptor, kernel)
	return MPSNDArrayFromID(rv)
}
