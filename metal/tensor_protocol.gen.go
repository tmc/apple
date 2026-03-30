// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A resource representing a multi-dimensional array that you can use with machine learning workloads.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensor
type MTLTensor interface {
	objectivec.IObject
	MTLAllocation
	MTLResource

	// A buffer instance this tensor shares its storage with or nil if this tensor does not wrap an underlying buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTensor/buffer
	Buffer() MTLBuffer

	// An offset, in bytes, into the buffer instance this tensor shares its storage with, or zero if this tensor does not wrap an underlying buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTensor/bufferOffset
	BufferOffset() uint

	// An underlying data format of this tensor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTensor/dataType
	DataType() MTLTensorDataType

	// An array of sizes, in elements, one for each dimension of this tensor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTensor/dimensions
	Dimensions() IMTLTensorExtents

	// A handle that represents the GPU resource, which you can store in an argument buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTensor/gpuResourceID
	GpuResourceID() MTLResourceID

	// An array of strides, in elements, one for each dimension of this tensor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTensor/strides
	Strides() IMTLTensorExtents

	// A set of contexts in which you can use this tensor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTensor/usage
	Usage() MTLTensorUsage

	// Copies the data corresponding to a slice of this tensor into a pointer you provide.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTensor/getBytes(_:strides:sliceOrigin:sliceDimensions:)
	GetBytesStridesFromSliceOriginSliceDimensions(bytes unsafe.Pointer, strides IMTLTensorExtents, sliceOrigin IMTLTensorExtents, sliceDimensions IMTLTensorExtents)

	// Replaces the contents of a slice of this tensor with data you provide.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLTensor/replace(sliceOrigin:sliceDimensions:withBytes:strides:)
	ReplaceSliceOriginSliceDimensionsWithBytesStrides(sliceOrigin IMTLTensorExtents, sliceDimensions IMTLTensorExtents, bytes unsafe.Pointer, strides IMTLTensorExtents)
}

// MTLTensorObject wraps an existing Objective-C object that conforms to the MTLTensor protocol.
type MTLTensorObject struct {
	objectivec.Object
}

func (o MTLTensorObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLTensorObjectFromID constructs a [MTLTensorObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLTensorObjectFromID(id objc.ID) MTLTensorObject {
	return MTLTensorObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// A buffer instance this tensor shares its storage with or nil if this tensor
// does not wrap an underlying buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensor/buffer
func (o MTLTensorObject) Buffer() MTLBuffer {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("buffer"))
	return MTLBufferObjectFromID(rv)
}

// An offset, in bytes, into the buffer instance this tensor shares its
// storage with, or zero if this tensor does not wrap an underlying buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensor/bufferOffset
func (o MTLTensorObject) BufferOffset() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("bufferOffset"))
	return rv
}

// An underlying data format of this tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensor/dataType
func (o MTLTensorObject) DataType() MTLTensorDataType {
	rv := objc.Send[MTLTensorDataType](o.ID, objc.Sel("dataType"))
	return rv
}

// An array of sizes, in elements, one for each dimension of this tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensor/dimensions
func (o MTLTensorObject) Dimensions() IMTLTensorExtents {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("dimensions"))
	return MTLTensorExtentsFromID(rv)
}

// A handle that represents the GPU resource, which you can store in an
// argument buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensor/gpuResourceID
func (o MTLTensorObject) GpuResourceID() MTLResourceID {
	rv := objc.Send[MTLResourceID](o.ID, objc.Sel("gpuResourceID"))
	return rv
}

// An array of strides, in elements, one for each dimension of this tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensor/strides
func (o MTLTensorObject) Strides() IMTLTensorExtents {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("strides"))
	return MTLTensorExtentsFromID(rv)
}

// A set of contexts in which you can use this tensor.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensor/usage
func (o MTLTensorObject) Usage() MTLTensorUsage {
	rv := objc.Send[MTLTensorUsage](o.ID, objc.Sel("usage"))
	return rv
}

// Copies the data corresponding to a slice of this tensor into a pointer you
// provide.
//
// bytes: A pointer to bytes of data that this method copies into the slice you
// specify with `sliceOrigin` and `sliceDimensions`.
//
// strides: An array of strides, in elements, that describes the layout of the data in
// `bytes`. You are responsible for ensuring `strides` meets the following
// requirements:
//
// - Elements of `strides`are in monotonically non-decreasing order. - For any
// `i` larger than zero, `strides[i]` is greater than or equal to
// `strides[i-1] * dimensions[i-1]`.
//
// sliceOrigin: An array of offsets, in elements, to the first element of the slice that
// this method reads data from.
//
// sliceDimensions: An array of sizes, in elements, of the slice this method reads data from.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensor/getBytes(_:strides:sliceOrigin:sliceDimensions:)
func (o MTLTensorObject) GetBytesStridesFromSliceOriginSliceDimensions(bytes unsafe.Pointer, strides IMTLTensorExtents, sliceOrigin IMTLTensorExtents, sliceDimensions IMTLTensorExtents) {
	objc.Send[struct{}](o.ID, objc.Sel("getBytes:strides:fromSliceOrigin:sliceDimensions:"), bytes, strides, sliceOrigin, sliceDimensions)
}

// Replaces the contents of a slice of this tensor with data you provide.
//
// sliceOrigin: An array of offsets, in elements, to the first element of the slice that
// this method writes data to.
//
// sliceDimensions: An array of sizes, in elements, of the slice this method writes data to.
//
// bytes: A pointer to bytes of data that this method copies into the slice you
// specify with `sliceOrigin` and `sliceDimensions`.
//
// strides: An array of strides, in elements, that describes the layout of the data in
// `bytes`. You are responsible for ensuring `strides` meets the following
// requirements:
//
// - Elements of `strides`are in monotonically non-decreasing order. - For any
// `i` larger than zero, `strides[i]` is greater than or equal to
// `strides[i-1] * dimensions[i-1]`.
//
// See: https://developer.apple.com/documentation/Metal/MTLTensor/replace(sliceOrigin:sliceDimensions:withBytes:strides:)
func (o MTLTensorObject) ReplaceSliceOriginSliceDimensionsWithBytesStrides(sliceOrigin IMTLTensorExtents, sliceDimensions IMTLTensorExtents, bytes unsafe.Pointer, strides IMTLTensorExtents) {
	objc.Send[struct{}](o.ID, objc.Sel("replaceSliceOrigin:sliceDimensions:withBytes:strides:"), sliceOrigin, sliceDimensions, bytes, strides)
}

// The amount of memory, in byes, a resource consumes, such as for a buffer,
// texture, or heap.
//
// See: https://developer.apple.com/documentation/Metal/MTLAllocation/allocatedSize
func (o MTLTensorObject) AllocatedSize() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("allocatedSize"))
	return rv
}

// The device object that created the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/device
func (o MTLTensorObject) Device() MTLDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
}

// A string that identifies the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/label
func (o MTLTensorObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

// The CPU cache mode that defines the CPU mapping of the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/cpuCacheMode
func (o MTLTensorObject) CpuCacheMode() MTLCPUCacheMode {
	rv := objc.Send[MTLCPUCacheMode](o.ID, objc.Sel("cpuCacheMode"))
	return rv
}

// The location and access permissions of the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/storageMode
func (o MTLTensorObject) StorageMode() MTLStorageMode {
	rv := objc.Send[MTLStorageMode](o.ID, objc.Sel("storageMode"))
	return rv
}

// A mode that determines whether Metal tracks and synchronizes resource
// access.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/hazardTrackingMode
func (o MTLTensorObject) HazardTrackingMode() MTLHazardTrackingMode {
	rv := objc.Send[MTLHazardTrackingMode](o.ID, objc.Sel("hazardTrackingMode"))
	return rv
}

// The storage mode, CPU cache mode, and hazard tracking mode of the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/resourceOptions
func (o MTLTensorObject) ResourceOptions() MTLResourceOptions {
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
// If `state` is [MTLPurgeableStateKeepCurrent], the method returns the
// current purgeable state without changing it.
//
// If `state` is [MTLPurgeableStateNonVolatile], the resource is marked to
// inform the caller that the data should not be discarded.
//
// If `state` is [MTLPurgeableState.empty], the resource is marked as data
// that can be discarded, because the caller no longer needs the contents of
// the resource.
//
// If `state` is [MTLPurgeableStateVolatile], the resource is marked as data
// that can be discarded, even if the caller may need the resource.
// [MTLResource] objects can be made purgeable, even if the caller may need
// the resource, where the implementation can reclaim the underlying storage
// at any time without informing the app. Purgeable resources may enable an
// app to keep larger caches of idle memory that may be useful again in the
// future without the risk of preventing the allocation of more important
// memory.
//
// When you use purgeable memory, make sure the block of memory is locked
// before you access it. This locking mechanism ensures that auto-removal
// policies don’t discard the data while you are accessing it. Similarly,
// the locking mechanism ensures that the virtual memory system has not
// already discarded the data.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/setPurgeableState(_:)
//
// [MTLPurgeableState.empty]: https://developer.apple.com/documentation/Metal/MTLPurgeableState/empty
func (o MTLTensorObject) SetPurgeableState(state MTLPurgeableState) MTLPurgeableState {
	rv := objc.Send[MTLPurgeableState](o.ID, objc.Sel("setPurgeableState:"), state)
	return rv
}

// The distance, in bytes, from the beginning of the heap to the first byte of
// the resource, if you allocated the resource on a heap.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/heapOffset
func (o MTLTensorObject) HeapOffset() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("heapOffset"))
	return rv
}

// The heap on which the resource is allocated, if any.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/heap
func (o MTLTensorObject) Heap() MTLHeap {
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
// allocation type [MTLHeapTypePlacement] and manage memory yourself instead.
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
func (o MTLTensorObject) MakeAliasable() {
	objc.Send[struct{}](o.ID, objc.Sel("makeAliasable"))
}

// A Boolean value that indicates whether future heap resource allocations may
// alias against the resource’s memory.
//
// # Return Value
//
// The default value is false. The value is true only if the [MakeAliasable]
// method was previously called on this resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/isAliasable()
func (o MTLTensorObject) IsAliasable() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAliasable"))
	return rv
}

// See: https://developer.apple.com/documentation/Metal/MTLResource/setOwnerWithIdentity:
func (o MTLTensorObject) SetOwnerWithIdentity(task_id_token kernel.Task_id_token_t) int32 {
	rv := objc.Send[int32](o.ID, objc.Sel("setOwnerWithIdentity:"), task_id_token)
	return rv
}

// A string that identifies the resource.
//
// # Discussion
//
// Object and command labels are useful identifiers at runtime or when
// profiling and debugging your app using any Metal tool. See [Naming
// resources and commands].
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/label
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
func (o MTLTensorObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}
