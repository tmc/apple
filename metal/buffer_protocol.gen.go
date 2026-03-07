// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A resource that stores data in a format defined by your app.
//
// See: https://developer.apple.com/documentation/Metal/MTLBuffer
type MTLBuffer interface {
	objectivec.IObject
	MTLAllocation
	MTLResource

	// Creates a texture that shares its storage with the buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBuffer/makeTexture(descriptor:offset:bytesPerRow:)
	NewTextureWithDescriptorOffsetBytesPerRow(descriptor IMTLTextureDescriptor, offset uint, bytesPerRow uint) MTLTexture

	// Gets the system address of the buffer’s storage allocation.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBuffer/contents()
	Contents()

	// Removes all debug marker strings from the buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBuffer/removeAllDebugMarkers()
	RemoveAllDebugMarkers()

	// The logical size of the buffer, in bytes.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBuffer/length
	Length() uint

	// Creates a remote view of the buffer for another GPU in the same peer group.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBuffer/makeRemoteBufferView(_:)
	NewRemoteBufferViewForDevice(device MTLDevice) MTLBuffer

	// The buffer on another GPU that the buffer was created from, if any.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBuffer/remoteStorageBuffer
	RemoteStorageBuffer() MTLBuffer

	// GpuAddress protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBuffer/gpuAddress
	GpuAddress() MTLGPUAddress

	// SparseBufferTier protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBuffer/sparseBufferTier
	SparseBufferTier() MTLBufferSparseTier

	// Creates a tensor that shares storage with this buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBuffer/makeTensor(descriptor:offset:)
	NewTensorWithDescriptorOffsetError(descriptor IMTLTensorDescriptor, offset uint) (MTLTensor, error)

	// Adds a debug marker string to a specific buffer range.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBuffer/addDebugMarker:range:
	AddDebugMarkerRange(marker string, range_ foundation.NSRange)

	// Informs the GPU that the CPU has modified a section of the buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLBuffer/didModifyRange:
	DidModifyRange(range_ foundation.NSRange)
}



// MTLBufferObject wraps an existing Objective-C object that conforms to the MTLBuffer protocol.
type MTLBufferObject struct {
	objectivec.Object
}
func (o MTLBufferObject) BaseObject() objectivec.Object {
	return o.Object
}



// MTLBufferObjectFromID constructs a [MTLBufferObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLBufferObjectFromID(id objc.ID) MTLBufferObject {
	return MTLBufferObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Creates a texture that shares its storage with the buffer.
//
// descriptor: The descriptor that contains the properties of the texture.
//
// offset: The offset, in bytes, from the base address for the first row of texture
// data.
//
// bytesPerRow: The stride, in bytes, from one row of texture data to the next.
//
// # Return Value
// 
// A new texture that shares the buffer’s underlying storage.
//
// # Discussion
// 
// This method creates a new [MTLTexture] instance that uses the same data as
// the buffer’s. Modifying the buffer also modifies the new texture because
// they share the same underlying memory.
// 
// The texture’s resource data is coherent between multiple render passes.
// However, that data may not be coherent within a single render pass due to
// caching at runtime. For example, a texture you create from the method may
// not be able to immediately reflect changes to the underlying buffer that
// come from a render or kernel function.
// 
// If this buffer’s [StorageMode] is [StorageModeManaged], and a render or
// kernel function modifies it, the CPU can access the new values through a
// texture after calling the [SynchronizeResource] method. CPU memory
// operations are only coherent between command buffer boundaries. GPU
// barriers guard its memory operations to buffers and textures so that each
// operation finishes running before the next one begins.
// 
// You can create multiple, nonoverlapping textures that use the same buffer;
// however, the GPU serializes memory operations to those textures.
// 
// To create a linear texture, you need to:
// 
// - Align the `offset` and `bytesPerRow` parameters to the value that the
// [MinimumLinearTextureAlignmentForPixelFormat] method returns. - Set the
// `bytesPerRow` parameter to a value greater than or equal to the number of
// bytes in one row of pixels — the product of the row’s width, in pixels,
// and the size of one pixel, in bytes.
// 
// Additionally, creating a linear texture from this method adds the following
// restrictions for the `descriptor` parameter’s properties:
// 
// [Table data omitted]
// 
// Samplers can use any [MTLSamplerAddressMode] to sample linear textures from
// this method on any device that supports the [GPUFamilyApple2] feature
// family or later.
//
// [MTLSamplerAddressMode]: https://developer.apple.com/documentation/Metal/MTLSamplerAddressMode
//
// See: https://developer.apple.com/documentation/Metal/MTLBuffer/makeTexture(descriptor:offset:bytesPerRow:)

func (o MTLBufferObject) NewTextureWithDescriptorOffsetBytesPerRow(descriptor IMTLTextureDescriptor, offset uint, bytesPerRow uint) MTLTexture {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newTextureWithDescriptor:offset:bytesPerRow:"), descriptor, offset, bytesPerRow)
	return MTLTextureObjectFromID(rv)
	}

// Gets the system address of the buffer’s storage allocation.
//
// # Return Value
// 
// A pointer to the shared copy of the buffer data, or [NULL] for buffers
// allocated with a private resource storage mode ([StorageModePrivate]).
//
// # Discussion
// 
// Private resources aren’t CPU-accessible.
//
// See: https://developer.apple.com/documentation/Metal/MTLBuffer/contents()

func (o MTLBufferObject) Contents() {
	
	objc.Send[struct{}](o.ID, objc.Sel("contents"))
	}

// Removes all debug marker strings from the buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLBuffer/removeAllDebugMarkers()

func (o MTLBufferObject) RemoveAllDebugMarkers() {
	
	objc.Send[struct{}](o.ID, objc.Sel("removeAllDebugMarkers"))
	}

// The logical size of the buffer, in bytes.
//
// See: https://developer.apple.com/documentation/Metal/MTLBuffer/length

func (o MTLBufferObject) Length() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("length"))
	return rv
	}

// Creates a remote view of the buffer for another GPU in the same peer group.
//
// # Discussion
// 
// The device instance that this buffer belongs to and the device you pass to
// the method both need to have the same nonzero peer group identifier
// ([peerGroupID]). This buffer needs to use the private storage mode
// ([StorageModePrivate]).
// 
// A remote view doesn’t allocate any storage for the new buffer; it
// references the memory allocated for the original buffer. You can use remote
// views only as a source for copy commands encoded by an
// [MTLBlitCommandEncoder]. For more information, see [Transferring data
// between connected GPUs].
//
// [Transferring data between connected GPUs]: https://developer.apple.com/documentation/Metal/transferring-data-between-connected-gpus
// [peerGroupID]: https://developer.apple.com/documentation/Metal/MTLDevice/peerGroupID
//
// See: https://developer.apple.com/documentation/Metal/MTLBuffer/makeRemoteBufferView(_:)

func (o MTLBufferObject) NewRemoteBufferViewForDevice(device MTLDevice) MTLBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newRemoteBufferViewForDevice:"), device)
	return MTLBufferObjectFromID(rv)
	}

// The buffer on another GPU that the buffer was created from, if any.
//
// See: https://developer.apple.com/documentation/Metal/MTLBuffer/remoteStorageBuffer

func (o MTLBufferObject) RemoteStorageBuffer() MTLBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("remoteStorageBuffer"))
	return MTLBufferObjectFromID(rv)
	}

// See: https://developer.apple.com/documentation/Metal/MTLBuffer/gpuAddress

func (o MTLBufferObject) GpuAddress() MTLGPUAddress {
	
	rv := objc.Send[MTLGPUAddress](o.ID, objc.Sel("gpuAddress"))
	return rv
	}

// See: https://developer.apple.com/documentation/Metal/MTLBuffer/sparseBufferTier

func (o MTLBufferObject) SparseBufferTier() MTLBufferSparseTier {
	
	rv := objc.Send[MTLBufferSparseTier](o.ID, objc.Sel("sparseBufferTier"))
	return rv
	}

// Creates a tensor that shares storage with this buffer.
//
// descriptor: A description of the properties for the new tensor.
//
// offset: Offset into the buffer at which the data of the tensor begins.
//
// # Discussion
// 
// If the descriptor specifies [MTLTensorUsageMachineLearning] usage, you need
// to observe the following restrictions:
// 
// - pass in `0` for the `offset` parameter - set the element stride the
// descriptor to `1` - ensure that number of bytes per row is a multiple of
// `64` - for dimensions greater than `2`, make sure `strides[dim] =
// strides[dim -1] * dimensions[dim - 1]`
//
// See: https://developer.apple.com/documentation/Metal/MTLBuffer/makeTensor(descriptor:offset:)

func (o MTLBufferObject) NewTensorWithDescriptorOffsetError(descriptor IMTLTensorDescriptor, offset uint) (MTLTensor, error) {
	
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("newTensorWithDescriptor:offset:error:"), descriptor, offset)
	if err != nil {
		return nil, err
	}
	return MTLTensorObjectFromID(rv), nil
	}

// Adds a debug marker string to a specific buffer range.
//
// marker: A string that identifies the marked buffer range.
//
// range: The range of bytes that you want to identify.
//
// See: https://developer.apple.com/documentation/Metal/MTLBuffer/addDebugMarker:range:

func (o MTLBufferObject) AddDebugMarkerRange(marker string, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("addDebugMarker:range:"), objc.String(marker), range_)
	}

// Informs the GPU that the CPU has modified a section of the buffer.
//
// range: The range of bytes that were modified.
//
// # Discussion
// 
// If you write information to a buffer created with the [StorageModeManaged]
// storage mode, you need to call this method to inform the GPU that the
// information has changed. If you execute GPU commands that read from the
// modified sections without calling this method first, the behavior is
// undefined.
//
// See: https://developer.apple.com/documentation/Metal/MTLBuffer/didModifyRange:

func (o MTLBufferObject) DidModifyRange(range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("didModifyRange:"), range_)
	}

// The amount of memory, in byes, a resource consumes, such as for a buffer,
// texture, or heap.
//
// See: https://developer.apple.com/documentation/Metal/MTLAllocation/allocatedSize

func (o MTLBufferObject) AllocatedSize() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("allocatedSize"))
	return rv
	}

// The device object that created the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/device

func (o MTLBufferObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}

// A string that identifies the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/label

func (o MTLBufferObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}

// The CPU cache mode that defines the CPU mapping of the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/cpuCacheMode

func (o MTLBufferObject) CpuCacheMode() MTLCPUCacheMode {
	
	rv := objc.Send[MTLCPUCacheMode](o.ID, objc.Sel("cpuCacheMode"))
	return rv
	}

// The location and access permissions of the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/storageMode

func (o MTLBufferObject) StorageMode() MTLStorageMode {
	
	rv := objc.Send[MTLStorageMode](o.ID, objc.Sel("storageMode"))
	return rv
	}

// A mode that determines whether Metal tracks and synchronizes resource
// access.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/hazardTrackingMode

func (o MTLBufferObject) HazardTrackingMode() MTLHazardTrackingMode {
	
	rv := objc.Send[MTLHazardTrackingMode](o.ID, objc.Sel("hazardTrackingMode"))
	return rv
	}

// The storage mode, CPU cache mode, and hazard tracking mode of the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/resourceOptions

func (o MTLBufferObject) ResourceOptions() MTLResourceOptions {
	
	rv := objc.Send[MTLResourceOptions](o.ID, objc.Sel("resourceOptions"))
	return rv
	}

// Specifies or queries the resource’s purgeable state.
//
// state: The desired purgeable state of a resource.
//
// # Return Value
// 
// The prior purgeable state of the resource.
//
// # Discussion
// 
// If `state` is [PurgeableStateKeepCurrent], the method returns the current
// purgeable state without changing it.
// 
// If `state` is [PurgeableStateNonVolatile], the resource is marked to inform
// the caller that the data should not be discarded.
// 
// If `state` is [PurgeableStateEmpty], the resource is marked as data that
// can be discarded, because the caller no longer needs the contents of the
// resource.
// 
// If `state` is [PurgeableStateVolatile], the resource is marked as data that
// can be discarded, even if the caller may need the resource. [MTLResource]
// objects can be made purgeable, even if the caller may need the resource,
// where the implementation can reclaim the underlying storage at any time
// without informing the app. Purgeable resources may enable an app to keep
// larger caches of idle memory that may be useful again in the future without
// the risk of preventing the allocation of more important memory.
// 
// When you use purgeable memory, make sure the block of memory is locked
// before you access it. This locking mechanism ensures that auto-removal
// policies don’t discard the data while you are accessing it. Similarly,
// the locking mechanism ensures that the virtual memory system has not
// already discarded the data.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/setPurgeableState(_:)

func (o MTLBufferObject) SetPurgeableState(state MTLPurgeableState) MTLPurgeableState {
	
	rv := objc.Send[MTLPurgeableState](o.ID, objc.Sel("setPurgeableState:"), state)
	return rv
	}

// The distance, in bytes, from the beginning of the heap to the first byte of
// the resource, if you allocated the resource on a heap.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/heapOffset

func (o MTLBufferObject) HeapOffset() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("heapOffset"))
	return rv
	}

// The heap on which the resource is allocated, if any.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/heap

func (o MTLBufferObject) Heap() MTLHeap {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("heap"))
	return MTLHeapObjectFromID(rv)
	}

// Allows future heap resource allocations to alias against the resource’s
// memory, reusing it.
//
// # Discussion
// 
// Resource instances marked as aliased have backing memory available for use
// in new allocations to the heap. One common use case is to make a single
// large resource aliasable for reuse of memory by smaller and more frequent
// resource allocations. For situations where you need fine-grained control
// over your memory management, you might want to use a heap with the
// allocation type [HeapTypePlacement] and manage memory yourself instead.
// 
// Aliased resources can’t be un-aliased or moved. If you use an aliased
// resource instance to read or write data, it results in undefined behavior.
// 
// When working with resources possibly backed by aliased memory, you should
// take great care that the system doesn’t access resources from multiple
// aliases concurrently. Use an [MTLEvent] or [MTLFence] instance to protect
// access to resources that you’ve either already aliased or intend to
// alias.
// 
// The general process to reuse memory from aliased resources is:
// 
// - Allocate an [MTLHeap] instance to hold your task’s resources, using the
// [NewHeapWithDescriptor] method. Your heap should be big enough to store the
// maximum amount of concurrently loaded data you expect. - Allocate your
// resource(s) using a heap method that returns an [MTLResource] instance. -
// Perform your stage on the GPU, and when it completes, mark the resource
// allocation(s) as aliasable by calling this method. - For each successive
// stage of your overall pass, repeat steps 2 and 3. Ensure that the prior
// stage fully completes before making any new resources on an aliasable heap
// through an event or fence.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/makeAliasable()

func (o MTLBufferObject) MakeAliasable() {
	
	objc.Send[struct{}](o.ID, objc.Sel("makeAliasable"))
	}

// A Boolean value that indicates whether future heap resource allocations may
// alias against the resource’s memory.
//
// # Return Value
// 
// The default value is [false]. The value is [true] only if the
// [MakeAliasable] method was previously called on this resource.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/isAliasable()

func (o MTLBufferObject) IsAliasable() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAliasable"))
	return rv
	}

//
// See: https://developer.apple.com/documentation/Metal/MTLResource/setOwnerWithIdentity:

func (o MTLBufferObject) SetOwnerWithIdentity(task_id_token objectivec.IObject) int32 {
	
	rv := objc.Send[int32](o.ID, objc.Sel("setOwnerWithIdentity:"), task_id_token)
	return rv
	}
















func (o MTLBufferObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

















