// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A memory pool from which you can suballocate resources.
//
// See: https://developer.apple.com/documentation/Metal/MTLHeap
type MTLHeap interface {
	objectivec.IObject
	MTLAllocation

	// A string that identifies the heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/label
	Label() string

	// Creates a buffer on the heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/makeBuffer(length:options:)
	NewBufferWithLengthOptions(length uint, options MTLResourceOptions) MTLBuffer

	// Creates a buffer at a specified offset on the heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/makeBuffer(length:options:offset:)
	NewBufferWithLengthOptionsOffset(length uint, options MTLResourceOptions, offset uint) MTLBuffer

	// Creates a texture on the heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/makeTexture(descriptor:)
	NewTextureWithDescriptor(descriptor IMTLTextureDescriptor) MTLTexture

	// Creates a texture at a specified offset on the heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/makeTexture(descriptor:offset:)
	NewTextureWithDescriptorOffset(descriptor IMTLTextureDescriptor, offset uint) MTLTexture

	// NewAccelerationStructureWithSize protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/makeAccelerationStructure(size:)
	NewAccelerationStructureWithSize(size uint) MTLAccelerationStructure

	// NewAccelerationStructureWithSizeOffset protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/makeAccelerationStructure(size:offset:)
	NewAccelerationStructureWithSizeOffset(size uint, offset uint) MTLAccelerationStructure

	// NewAccelerationStructureWithDescriptor protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/makeAccelerationStructure(descriptor:)
	NewAccelerationStructureWithDescriptor(descriptor IMTLAccelerationStructureDescriptor) MTLAccelerationStructure

	// NewAccelerationStructureWithDescriptorOffset protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/makeAccelerationStructure(descriptor:offset:)
	NewAccelerationStructureWithDescriptorOffset(descriptor IMTLAccelerationStructureDescriptor, offset uint) MTLAccelerationStructure

	// Sets the purgeable state of the heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/setPurgeableState(_:)
	SetPurgeableState(state MTLPurgeableState) MTLPurgeableState

	// The maximum size of a resource, in bytes, that can be currently allocated from the heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/maxAvailableSize(alignment:)
	MaxAvailableSizeWithAlignment(alignment uint) uint

	// The total size of the heap, in bytes.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/size
	Size() uint

	// The size of all resources currently in the heap, in bytes.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/usedSize
	UsedSize() uint

	// The size, in bytes, of the current heap allocation.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/currentAllocatedSize
	CurrentAllocatedSize() uint

	// The device object that created the heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/device
	Device() MTLDevice

	// The heap’s type.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/type
	Type() MTLHeapType

	// The heap’s storage mode.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/storageMode
	StorageMode() MTLStorageMode

	// The heap’s CPU cache mode.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/cpuCacheMode
	CpuCacheMode() MTLCPUCacheMode

	// The heap’s hazard tracking mode.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/hazardTrackingMode
	HazardTrackingMode() MTLHazardTrackingMode

	// The options for resources created by the heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/resourceOptions
	ResourceOptions() MTLResourceOptions

	// A string that identifies the heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLHeap/label
	SetLabel(value string)
}



// MTLHeapObject wraps an existing Objective-C object that conforms to the MTLHeap protocol.
type MTLHeapObject struct {
	objectivec.Object
}
func (o MTLHeapObject) BaseObject() objectivec.Object {
	return o.Object
}



// MTLHeapObjectFromID constructs a [MTLHeapObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLHeapObjectFromID(id objc.ID) MTLHeapObject {
	return MTLHeapObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// A string that identifies the heap.
//
// See: https://developer.apple.com/documentation/Metal/MTLHeap/label

func (o MTLHeapObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}

// Creates a buffer on the heap.
//
// length: The size, in bytes, of the buffer.
//
// options: Options that describe the properties of the buffer.
//
// # Return Value
// 
// A new buffer object backed by heap memory, or `nil` if the heap memory is
// full.
//
// # Discussion
// 
// You can call the method with the following restrictions:
// 
// - The heap’s type needs to be [HeapTypeAutomatic] - The buffer’s
// storage mode option needs to match the heap’s [StorageMode] property -
// The buffer’s CPU cache mode option needs to match the heap’s
// [CpuCacheMode] property
//
// See: https://developer.apple.com/documentation/Metal/MTLHeap/makeBuffer(length:options:)

func (o MTLHeapObject) NewBufferWithLengthOptions(length uint, options MTLResourceOptions) MTLBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newBufferWithLength:options:"), length, options)
	return MTLBufferObjectFromID(rv)
	}

// Creates a buffer at a specified offset on the heap.
//
// length: The size of the buffer, in bytes.
//
// options: Options that describe the properties of the buffer.
//
// offset: The distance, in bytes, to place the buffer relative to the start of the
// heap.
//
// # Return Value
// 
// A new buffer, or `nil` if the heap is not a placement heap.
//
// # Discussion
// 
// You can call the method with the following restrictions:
// 
// - The heap’s type needs to be [HeapTypePlacement] - The buffer’s
// storage mode option needs to match the heap’s [StorageMode] property -
// The buffer’s CPU cache mode option needs to match the heap’s
// [CpuCacheMode] property
// 
// Use the [HeapBufferSizeAndAlignWithLengthOptions] method to determine the
// required size and alignment. If you don’t align the buffer correctly or
// it extends past the end of the heap, the behavior is undefined.
//
// See: https://developer.apple.com/documentation/Metal/MTLHeap/makeBuffer(length:options:offset:)

func (o MTLHeapObject) NewBufferWithLengthOptionsOffset(length uint, options MTLResourceOptions, offset uint) MTLBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newBufferWithLength:options:offset:"), length, options, offset)
	return MTLBufferObjectFromID(rv)
	}

// Creates a texture on the heap.
//
// descriptor: A descriptor object that describes the properties of the texture.
//
// # Return Value
// 
// A new texture object backed by heap memory, or `nil` if the heap memory is
// full.
//
// # Discussion
// 
// You can call the method with the following restrictions:
// 
// - The heap’s type needs to be [HeapTypeAutomatic] - The texture’s CPU
// cache mode option needs to match the heap’s [CpuCacheMode] property - The
// texture’s storage mode option needs to be [StorageModeMemoryless], or
// match the heap’s [StorageMode] property
//
// See: https://developer.apple.com/documentation/Metal/MTLHeap/makeTexture(descriptor:)

func (o MTLHeapObject) NewTextureWithDescriptor(descriptor IMTLTextureDescriptor) MTLTexture {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newTextureWithDescriptor:"), descriptor)
	return MTLTextureObjectFromID(rv)
	}

// Creates a texture at a specified offset on the heap.
//
// descriptor: A descriptor object that describes the properties of the texture.
//
// offset: The distance, in bytes, to place the texture relative to the start of the
// heap.
//
// # Return Value
// 
// A new texture, or `nil` if the heap is not a placement heap.
//
// # Discussion
// 
// You can call the method with the following restrictions:
// 
// - The heap’s type needs to be [HeapTypePlacement] - The texture’s CPU
// cache mode option needs to match the heap’s [CpuCacheMode] property - The
// texture’s storage mode option needs to be [StorageModeMemoryless], or
// match the heap’s [StorageMode] property
// 
// Use the [HeapBufferSizeAndAlignWithLengthOptions] to determine the correct
// size and alignment.
//
// See: https://developer.apple.com/documentation/Metal/MTLHeap/makeTexture(descriptor:offset:)

func (o MTLHeapObject) NewTextureWithDescriptorOffset(descriptor IMTLTextureDescriptor, offset uint) MTLTexture {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newTextureWithDescriptor:offset:"), descriptor, offset)
	return MTLTextureObjectFromID(rv)
	}

//
// See: https://developer.apple.com/documentation/Metal/MTLHeap/makeAccelerationStructure(size:)

func (o MTLHeapObject) NewAccelerationStructureWithSize(size uint) MTLAccelerationStructure {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newAccelerationStructureWithSize:"), size)
	return MTLAccelerationStructureObjectFromID(rv)
	}

//
// See: https://developer.apple.com/documentation/Metal/MTLHeap/makeAccelerationStructure(size:offset:)

func (o MTLHeapObject) NewAccelerationStructureWithSizeOffset(size uint, offset uint) MTLAccelerationStructure {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newAccelerationStructureWithSize:offset:"), size, offset)
	return MTLAccelerationStructureObjectFromID(rv)
	}

//
// See: https://developer.apple.com/documentation/Metal/MTLHeap/makeAccelerationStructure(descriptor:)

func (o MTLHeapObject) NewAccelerationStructureWithDescriptor(descriptor IMTLAccelerationStructureDescriptor) MTLAccelerationStructure {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newAccelerationStructureWithDescriptor:"), descriptor)
	return MTLAccelerationStructureObjectFromID(rv)
	}

//
// See: https://developer.apple.com/documentation/Metal/MTLHeap/makeAccelerationStructure(descriptor:offset:)

func (o MTLHeapObject) NewAccelerationStructureWithDescriptorOffset(descriptor IMTLAccelerationStructureDescriptor, offset uint) MTLAccelerationStructure {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newAccelerationStructureWithDescriptor:offset:"), descriptor, offset)
	return MTLAccelerationStructureObjectFromID(rv)
	}

// Sets the purgeable state of the heap.
//
// state: The desired purgeable state of the heap.
//
// # Return Value
// 
// The previous purgeable state of the heap.
//
// # Discussion
// 
// The heap purgeability state refers to its whole backing memory and affects
// all resources in the heap. Heaps can be marked purgeable but its resources
// cannot; the heap’s resources always reflect the heap’s purgeability
// state.
// 
// Refer to the [MTLPurgeableState] and [SetPurgeableState] reference for
// further information.
//
// [MTLPurgeableState]: https://developer.apple.com/documentation/Metal/MTLPurgeableState
//
// See: https://developer.apple.com/documentation/Metal/MTLHeap/setPurgeableState(_:)

func (o MTLHeapObject) SetPurgeableState(state MTLPurgeableState) MTLPurgeableState {
	
	rv := objc.Send[MTLPurgeableState](o.ID, objc.Sel("setPurgeableState:"), state)
	return rv
	}

// The maximum size of a resource, in bytes, that can be currently allocated
// from the heap.
//
// alignment: The alignment of the resource, in bytes. This value needs to be a power of
// two.
//
// # Return Value
// 
// The maximum size for the resource, in bytes.
//
// # Discussion
// 
// This method measures fragmentation within the heap. You can use the
// [HeapBufferSizeAndAlignWithLengthOptions] and
// [HeapTextureSizeAndAlignWithDescriptor] methods to help you determine the
// correct alignment for the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLHeap/maxAvailableSize(alignment:)

func (o MTLHeapObject) MaxAvailableSizeWithAlignment(alignment uint) uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("maxAvailableSizeWithAlignment:"), alignment)
	return rv
	}

// The total size of the heap, in bytes.
//
// See: https://developer.apple.com/documentation/Metal/MTLHeap/size

func (o MTLHeapObject) Size() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("size"))
	return rv
	}

// The size of all resources currently in the heap, in bytes.
//
// See: https://developer.apple.com/documentation/Metal/MTLHeap/usedSize

func (o MTLHeapObject) UsedSize() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("usedSize"))
	return rv
	}

// The size, in bytes, of the current heap allocation.
//
// See: https://developer.apple.com/documentation/Metal/MTLHeap/currentAllocatedSize

func (o MTLHeapObject) CurrentAllocatedSize() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("currentAllocatedSize"))
	return rv
	}

// The device object that created the heap.
//
// See: https://developer.apple.com/documentation/Metal/MTLHeap/device

func (o MTLHeapObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}

// The heap’s type.
//
// See: https://developer.apple.com/documentation/Metal/MTLHeap/type

func (o MTLHeapObject) Type() MTLHeapType {
	
	rv := objc.Send[MTLHeapType](o.ID, objc.Sel("type"))
	return rv
	}

// The heap’s storage mode.
//
// See: https://developer.apple.com/documentation/Metal/MTLHeap/storageMode

func (o MTLHeapObject) StorageMode() MTLStorageMode {
	
	rv := objc.Send[MTLStorageMode](o.ID, objc.Sel("storageMode"))
	return rv
	}

// The heap’s CPU cache mode.
//
// See: https://developer.apple.com/documentation/Metal/MTLHeap/cpuCacheMode

func (o MTLHeapObject) CpuCacheMode() MTLCPUCacheMode {
	
	rv := objc.Send[MTLCPUCacheMode](o.ID, objc.Sel("cpuCacheMode"))
	return rv
	}

// The heap’s hazard tracking mode.
//
// See: https://developer.apple.com/documentation/Metal/MTLHeap/hazardTrackingMode

func (o MTLHeapObject) HazardTrackingMode() MTLHazardTrackingMode {
	
	rv := objc.Send[MTLHazardTrackingMode](o.ID, objc.Sel("hazardTrackingMode"))
	return rv
	}

// The options for resources created by the heap.
//
// See: https://developer.apple.com/documentation/Metal/MTLHeap/resourceOptions

func (o MTLHeapObject) ResourceOptions() MTLResourceOptions {
	
	rv := objc.Send[MTLResourceOptions](o.ID, objc.Sel("resourceOptions"))
	return rv
	}

// The amount of memory, in byes, a resource consumes, such as for a buffer,
// texture, or heap.
//
// See: https://developer.apple.com/documentation/Metal/MTLAllocation/allocatedSize

func (o MTLHeapObject) AllocatedSize() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("allocatedSize"))
	return rv
	}




func (o MTLHeapObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

























