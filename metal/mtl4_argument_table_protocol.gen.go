// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// Provides a mechanism to manage and provide resource bindings for buffers, textures, sampler states and other Metal resources.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTable
type MTL4ArgumentTable interface {
	objectivec.IObject

	// The device from which you created this argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTable/device
	Device() MTLDevice

	// Assigns an optional label with this argument table for debugging purposes.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTable/label
	Label() string

	// Binds a GPU address to a buffer binding slot, providing a dynamic vertex stride.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTable/setAddress(_:attributeStride:index:)
	SetAddressAttributeStrideAtIndex(gpuAddress MTLGPUAddress, stride uint, bindingIndex uint)

	// Binds a GPU address to a buffer binding slot.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTable/setAddress(_:index:)
	SetAddressAtIndex(gpuAddress MTLGPUAddress, bindingIndex uint)

	// Binds a resource to a buffer binding slot.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTable/setResource(_:bufferIndex:)
	SetResourceAtBufferIndex(resourceID MTLResourceID, bindingIndex uint)

	// Binds a sampler state to a sampler state binding slot.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTable/setSamplerState(_:index:)
	SetSamplerStateAtIndex(resourceID MTLResourceID, bindingIndex uint)

	// Binds a texture to a texture binding slot.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTable/setTexture(_:index:)
	SetTextureAtIndex(resourceID MTLResourceID, bindingIndex uint)
}

// MTL4ArgumentTableObject wraps an existing Objective-C object that conforms to the MTL4ArgumentTable protocol.
type MTL4ArgumentTableObject struct {
	objectivec.Object
}
func (o MTL4ArgumentTableObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTL4ArgumentTableObjectFromID constructs a [MTL4ArgumentTableObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTL4ArgumentTableObjectFromID(id objc.ID) MTL4ArgumentTableObject {
	return MTL4ArgumentTableObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The device from which you created this argument table.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTable/device
func (o MTL4ArgumentTableObject) Device() MTLDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}
// Assigns an optional label with this argument table for debugging purposes.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTable/label
func (o MTL4ArgumentTableObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}
// Binds a GPU address to a buffer binding slot, providing a dynamic vertex
// stride.
//
// gpuAddress: The GPU address of a [MTLBuffer] to set.
//
// stride: The stride between attributes in the buffer.
//
// bindingIndex: A valid binding index in the buffer binding range. It is an error for this
// value to match or exceed the value of property [MaxBufferBindCount] on the
// descriptor from which you created this argument table.
//
// # Discussion
// 
// This method requires that the value of property [SupportAttributeStrides]
// on the descriptor from which you created this argument table is true.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTable/setAddress(_:attributeStride:index:)
func (o MTL4ArgumentTableObject) SetAddressAttributeStrideAtIndex(gpuAddress MTLGPUAddress, stride uint, bindingIndex uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setAddress:attributeStride:atIndex:"), gpuAddress, stride, bindingIndex)
	}
// Binds a GPU address to a buffer binding slot.
//
// gpuAddress: The GPU address of a [MTLBuffer] to set.
//
// bindingIndex: A valid binding index in the buffer binding range. It is an error for this
// value to match or exceed the value of property [MaxBufferBindCount] on the
// descriptor from which you created this argument table.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTable/setAddress(_:index:)
func (o MTL4ArgumentTableObject) SetAddressAtIndex(gpuAddress MTLGPUAddress, bindingIndex uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setAddress:atIndex:"), gpuAddress, bindingIndex)
	}
// Binds a resource to a buffer binding slot.
//
// resourceID: The [MTLResourceID] of the Metal resource to bind.
// //
// [MTLResourceID]: https://developer.apple.com/documentation/Metal/MTLResourceID
//
// bindingIndex: A valid binding index in the buffer binding range. It is an error for this
// value to match or exceed the value of property [MaxBufferBindCount] on the
// descriptor from which you created this argument table.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTable/setResource(_:bufferIndex:)
func (o MTL4ArgumentTableObject) SetResourceAtBufferIndex(resourceID MTLResourceID, bindingIndex uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setResource:atBufferIndex:"), resourceID, bindingIndex)
	}
// Binds a sampler state to a sampler state binding slot.
//
// resourceID: The [MTLResourceID] of the [MTLSamplerState] instance to bind.
// //
// [MTLResourceID]: https://developer.apple.com/documentation/Metal/MTLResourceID
//
// bindingIndex: A valid binding index in the sampler binding range. It is an error for this
// value to match or exceed the value of property [MaxSamplerStateBindCount]
// on the descriptor from which you created this argument table.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTable/setSamplerState(_:index:)
func (o MTL4ArgumentTableObject) SetSamplerStateAtIndex(resourceID MTLResourceID, bindingIndex uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setSamplerState:atIndex:"), resourceID, bindingIndex)
	}
// Binds a texture to a texture binding slot.
//
// resourceID: The [MTLResourceID] of the [MTLTexture] instance to bind.
// //
// [MTLResourceID]: https://developer.apple.com/documentation/Metal/MTLResourceID
//
// bindingIndex: A valid binding index in the texture binding range. It is an error for this
// value to match or exceed the value of property [MaxTextureBindCount] on the
// descriptor from which you created this argument table.
//
// See: https://developer.apple.com/documentation/Metal/MTL4ArgumentTable/setTexture(_:index:)
func (o MTL4ArgumentTableObject) SetTextureAtIndex(resourceID MTLResourceID, bindingIndex uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setTexture:atIndex:"), resourceID, bindingIndex)
	}

