// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A table of intersection functions that Metal calls to perform ray-tracing intersection tests.
//
// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable
type MTLIntersectionFunctionTable interface {
	objectivec.IObject
	MTLAllocation
	MTLResource

	// Sets an entry in the table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/setFunction(_:index:)
	SetFunctionAtIndex(function MTLFunctionHandle, index uint)

	// Sets a buffer for the intersection functions.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/setBuffer(_:offset:index:)
	SetBufferOffsetAtIndex(buffer MTLBuffer, offset uint, index uint)

	// Sets a visible function table for the intersection functions.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/setVisibleFunctionTable(_:bufferIndex:)
	SetVisibleFunctionTableAtBufferIndex(functionTable MTLVisibleFunctionTable, bufferIndex uint)

	// Sets an entry in the intersection table to point to a system-defined opaque triangle intersection function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/setOpaqueTriangleIntersectionFunction(signature:index:)
	SetOpaqueTriangleIntersectionFunctionWithSignatureAtIndex(signature MTLIntersectionFunctionSignature, index uint)

	// Sets a range of entries in the intersection table to point to a system-defined opaque triangle intersection function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/setOpaqueTriangleIntersectionFunction(signature:range:)
	SetOpaqueTriangleIntersectionFunctionWithSignatureWithRange(signature MTLIntersectionFunctionSignature, range_ foundation.NSRange)

	// GpuResourceID protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/gpuResourceID
	GpuResourceID() MTLResourceID

	// SetOpaqueCurveIntersectionFunctionWithSignatureAtIndex protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/setOpaqueCurveIntersectionFunction(signature:index:)
	SetOpaqueCurveIntersectionFunctionWithSignatureAtIndex(signature MTLIntersectionFunctionSignature, index uint)

	// SetOpaqueCurveIntersectionFunctionWithSignatureWithRange protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/setOpaqueCurveIntersectionFunction(signature:range:)
	SetOpaqueCurveIntersectionFunctionWithSignatureWithRange(signature MTLIntersectionFunctionSignature, range_ foundation.NSRange)

	// Sets a range of buffers for the intersection functions.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/setBuffers:offsets:withRange:
	SetBuffersOffsetsWithRange(buffers []MTLBuffer, offsets uint, range_ foundation.NSRange)

	// Sets a range of entries in the table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/setFunctions:withRange:
	SetFunctionsWithRange(functions []MTLFunctionHandle, range_ foundation.NSRange)

	// Sets a range of visible function tables for the intersection functions.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/setVisibleFunctionTables:withBufferRange:
	SetVisibleFunctionTablesWithBufferRange(functionTables []MTLVisibleFunctionTable, bufferRange foundation.NSRange)
}

// MTLIntersectionFunctionTableObject wraps an existing Objective-C object that conforms to the MTLIntersectionFunctionTable protocol.
type MTLIntersectionFunctionTableObject struct {
	objectivec.Object
}
func (o MTLIntersectionFunctionTableObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLIntersectionFunctionTableObjectFromID constructs a [MTLIntersectionFunctionTableObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLIntersectionFunctionTableObjectFromID(id objc.ID) MTLIntersectionFunctionTableObject {
	return MTLIntersectionFunctionTableObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Sets an entry in the table.
//
// function: A function handle for the intersection function.
//
// index: The index of the table entry to change.
//
// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/setFunction(_:index:)
func (o MTLIntersectionFunctionTableObject) SetFunctionAtIndex(function MTLFunctionHandle, index uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setFunction:atIndex:"), function, index)
	}
// Sets a buffer for the intersection functions.
//
// buffer: The [MTLBuffer] object to set in the argument table.
//
// offset: Where the data begins, in bytes, from the start of the buffer.
//
// index: An index in the function table’s buffer argument table.
//
// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/setBuffer(_:offset:index:)
func (o MTLIntersectionFunctionTableObject) SetBufferOffsetAtIndex(buffer MTLBuffer, offset uint, index uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setBuffer:offset:atIndex:"), buffer, offset, index)
	}
// Sets a visible function table for the intersection functions.
//
// functionTable: A visible function table.
//
// bufferIndex: An index in the function table’s buffer argument table.
//
// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/setVisibleFunctionTable(_:bufferIndex:)
func (o MTLIntersectionFunctionTableObject) SetVisibleFunctionTableAtBufferIndex(functionTable MTLVisibleFunctionTable, bufferIndex uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setVisibleFunctionTable:atBufferIndex:"), functionTable, bufferIndex)
	}
// Sets an entry in the intersection table to point to a system-defined opaque
// triangle intersection function.
//
// signature: The signature of the function.
//
// index: The index in the table to change.
//
// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/setOpaqueTriangleIntersectionFunction(signature:index:)
func (o MTLIntersectionFunctionTableObject) SetOpaqueTriangleIntersectionFunctionWithSignatureAtIndex(signature MTLIntersectionFunctionSignature, index uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setOpaqueTriangleIntersectionFunctionWithSignature:atIndex:"), signature, index)
	}
// Sets a range of entries in the intersection table to point to a
// system-defined opaque triangle intersection function.
//
// signature: The signature of the function.
//
// range: A range of indices to change in the table.
//
// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/setOpaqueTriangleIntersectionFunction(signature:range:)
func (o MTLIntersectionFunctionTableObject) SetOpaqueTriangleIntersectionFunctionWithSignatureWithRange(signature MTLIntersectionFunctionSignature, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setOpaqueTriangleIntersectionFunctionWithSignature:withRange:"), signature, range_)
	}
// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/gpuResourceID
func (o MTLIntersectionFunctionTableObject) GpuResourceID() MTLResourceID {
	rv := objc.Send[MTLResourceID](o.ID, objc.Sel("gpuResourceID"))
	return rv
	}
//
// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/setOpaqueCurveIntersectionFunction(signature:index:)
func (o MTLIntersectionFunctionTableObject) SetOpaqueCurveIntersectionFunctionWithSignatureAtIndex(signature MTLIntersectionFunctionSignature, index uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setOpaqueCurveIntersectionFunctionWithSignature:atIndex:"), signature, index)
	}
//
// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/setOpaqueCurveIntersectionFunction(signature:range:)
func (o MTLIntersectionFunctionTableObject) SetOpaqueCurveIntersectionFunctionWithSignatureWithRange(signature MTLIntersectionFunctionSignature, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setOpaqueCurveIntersectionFunctionWithSignature:withRange:"), signature, range_)
	}
// Sets a range of buffers for the intersection functions.
//
// buffers: A pointer to an array of buffers.
//
// offsets: A pointer to an array of offsets, measured in bytes, from the start of each
// buffer.
//
// range: A range of indices in the function table’s buffer argument table.
//
// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/setBuffers:offsets:withRange:
func (o MTLIntersectionFunctionTableObject) SetBuffersOffsetsWithRange(buffers []MTLBuffer, offsets uint, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setBuffers:offsets:withRange:"), buffers, offsets, range_)
	}
// Sets a range of entries in the table.
//
// functions: A pointer to an array of function handles.
//
// range: A range of indices to change in the table.
//
// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/setFunctions:withRange:
func (o MTLIntersectionFunctionTableObject) SetFunctionsWithRange(functions []MTLFunctionHandle, range_ foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setFunctions:withRange:"), functions, range_)
	}
// Sets a range of visible function tables for the intersection functions.
//
// functionTables: The function tables to insert.
//
// bufferRange: A range of indices in the function table’s buffer argument table.
//
// See: https://developer.apple.com/documentation/Metal/MTLIntersectionFunctionTable/setVisibleFunctionTables:withBufferRange:
func (o MTLIntersectionFunctionTableObject) SetVisibleFunctionTablesWithBufferRange(functionTables []MTLVisibleFunctionTable, bufferRange foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setVisibleFunctionTables:withBufferRange:"), functionTables, bufferRange)
	}
// The amount of memory, in byes, a resource consumes, such as for a buffer,
// texture, or heap.
//
// See: https://developer.apple.com/documentation/Metal/MTLAllocation/allocatedSize
func (o MTLIntersectionFunctionTableObject) AllocatedSize() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("allocatedSize"))
	return rv
	}
// The device object that created the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/device
func (o MTLIntersectionFunctionTableObject) Device() MTLDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}
// A string that identifies the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/label
func (o MTLIntersectionFunctionTableObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}
// The CPU cache mode that defines the CPU mapping of the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/cpuCacheMode
func (o MTLIntersectionFunctionTableObject) CpuCacheMode() MTLCPUCacheMode {
	rv := objc.Send[MTLCPUCacheMode](o.ID, objc.Sel("cpuCacheMode"))
	return rv
	}
// The location and access permissions of the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/storageMode
func (o MTLIntersectionFunctionTableObject) StorageMode() MTLStorageMode {
	rv := objc.Send[MTLStorageMode](o.ID, objc.Sel("storageMode"))
	return rv
	}
// A mode that determines whether Metal tracks and synchronizes resource
// access.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/hazardTrackingMode
func (o MTLIntersectionFunctionTableObject) HazardTrackingMode() MTLHazardTrackingMode {
	rv := objc.Send[MTLHazardTrackingMode](o.ID, objc.Sel("hazardTrackingMode"))
	return rv
	}
// The storage mode, CPU cache mode, and hazard tracking mode of the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/resourceOptions
func (o MTLIntersectionFunctionTableObject) ResourceOptions() MTLResourceOptions {
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
func (o MTLIntersectionFunctionTableObject) SetPurgeableState(state MTLPurgeableState) MTLPurgeableState {
	rv := objc.Send[MTLPurgeableState](o.ID, objc.Sel("setPurgeableState:"), state)
	return rv
	}
// The distance, in bytes, from the beginning of the heap to the first byte of
// the resource, if you allocated the resource on a heap.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/heapOffset
func (o MTLIntersectionFunctionTableObject) HeapOffset() uint {
	rv := objc.Send[uint](o.ID, objc.Sel("heapOffset"))
	return rv
	}
// The heap on which the resource is allocated, if any.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/heap
func (o MTLIntersectionFunctionTableObject) Heap() MTLHeap {
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
func (o MTLIntersectionFunctionTableObject) MakeAliasable() {
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
func (o MTLIntersectionFunctionTableObject) IsAliasable() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAliasable"))
	return rv
	}
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/setOwnerWithIdentity:
func (o MTLIntersectionFunctionTableObject) SetOwnerWithIdentity(task_id_token objectivec.IObject) int32 {
	rv := objc.Send[int32](o.ID, objc.Sel("setOwnerWithIdentity:"), task_id_token)
	return rv
	}

func (o MTLIntersectionFunctionTableObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

