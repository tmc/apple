// Code generated from Apple documentation for MetalPerformanceShaders. DO NOT EDIT.

package metalperformanceshaders

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/metal"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// MPSImageAllocator protocol.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAllocator
type MPSImageAllocator interface {
	objectivec.IObject
	foundation.NSCoding
	foundation.NSSecureCoding

	// ImageForCommandBufferImageDescriptorKernel protocol.
	//
	// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAllocator/image(for:imageDescriptor:kernel:)
	ImageForCommandBufferImageDescriptorKernel(cmdBuf metal.MTLCommandBuffer, descriptor IMPSImageDescriptor, kernel IMPSKernel) IMPSImage
}

// MPSImageAllocatorObject wraps an existing Objective-C object that conforms to the MPSImageAllocator protocol.
type MPSImageAllocatorObject struct {
	foundation.NSCodingObject
}

func (o MPSImageAllocatorObject) BaseObject() objectivec.Object {
	return o.NSCodingObject.BaseObject()
}

// MPSImageAllocatorObjectFromID constructs a [MPSImageAllocatorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MPSImageAllocatorObjectFromID(id objc.ID) MPSImageAllocatorObject {
	return MPSImageAllocatorObject{
		NSCodingObject: foundation.NSCodingObjectFromID(id),
	}
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAllocator/image(for:imageDescriptor:kernel:)
func (o MPSImageAllocatorObject) ImageForCommandBufferImageDescriptorKernel(cmdBuf metal.MTLCommandBuffer, descriptor IMPSImageDescriptor, kernel IMPSKernel) IMPSImage {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("imageForCommandBuffer:imageDescriptor:kernel:"), cmdBuf, descriptor, kernel)
	return MPSImageFromID(rv)
}

// See: https://developer.apple.com/documentation/MetalPerformanceShaders/MPSImageAllocator/imageBatch(for:imageDescriptor:kernel:count:)
func (o MPSImageAllocatorObject) ImageBatchForCommandBufferImageDescriptorKernelCount(cmdBuf metal.MTLCommandBuffer, descriptor IMPSImageDescriptor, kernel IMPSKernel, count uint) MPSImageBatch {
	rv := objc.Send[MPSImageBatch](o.ID, objc.Sel("imageBatchForCommandBuffer:imageDescriptor:kernel:count:"), cmdBuf, descriptor, kernel, count)
	return rv
}
