// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A table of shader functions visible to your app that you can pass into compute commands to customize the behavior of a shader.
//
// See: https://developer.apple.com/documentation/Metal/MTLVisibleFunctionTable
type MTLVisibleFunctionTable interface {
	objectivec.IObject
	MTLAllocation
	MTLResource

	// Sets a table entry to point to a callable function.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLVisibleFunctionTable/setFunction(_:index:)
	SetFunctionAtIndex(function MTLFunctionHandle, index uint)

	// GpuResourceID protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLVisibleFunctionTable/gpuResourceID
	GpuResourceID() MTLResourceID

	// Sets a range of table entries to point to an array of callable functions.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLVisibleFunctionTable/setFunctions:withRange:
	SetFunctionsWithRange(functions []MTLFunctionHandle, range_ foundation.NSRange)
}

// MTLVisibleFunctionTableObject wraps an existing Objective-C object that conforms to the MTLVisibleFunctionTable protocol.
type MTLVisibleFunctionTableObject struct {
	objectivec.Object
}
func (o MTLVisibleFunctionTableObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLVisibleFunctionTableObjectFromID constructs a [MTLVisibleFunctionTableObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLVisibleFunctionTableObjectFromID(id objc.ID) MTLVisibleFunctionTableObject {
	return MTLVisibleFunctionTableObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Sets a table entry to point to a callable function.
//
// function: A function handle for the function to be called.
//
// index: The index of the table entry to change.
//
// See: https://developer.apple.com/documentation/Metal/MTLVisibleFunctionTable/setFunction(_:index:)
func (o MTLVisibleFunctionTableObject) SetFunctionAtIndex(function MTLFunctionHandle, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setFunction:atIndex:"), function, index)
	}
// See: https://developer.apple.com/documentation/Metal/MTLVisibleFunctionTable/gpuResourceID
func (o MTLVisibleFunctionTableObject) GpuResourceID() MTLResourceID {
	
	rv := objc.Send[MTLResourceID](o.ID, objc.Sel("gpuResourceID"))
	return rv
	}
// Sets a range of table entries to point to an array of callable functions.
//
// functions: An array of function handles for the functions to be called.
//
// range: A range of indices to change in the table.
//
// See: https://developer.apple.com/documentation/Metal/MTLVisibleFunctionTable/setFunctions:withRange:
func (o MTLVisibleFunctionTableObject) SetFunctionsWithRange(functions []MTLFunctionHandle, range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setFunctions:withRange:"), functions, range_)
	}
// The amount of memory, in byes, a resource consumes, such as for a buffer,
// texture, or heap.
//
// See: https://developer.apple.com/documentation/Metal/MTLAllocation/allocatedSize
func (o MTLVisibleFunctionTableObject) AllocatedSize() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("allocatedSize"))
	return rv
	}
// The device object that created the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/device
func (o MTLVisibleFunctionTableObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}
// A string that identifies the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/label
func (o MTLVisibleFunctionTableObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}
// The CPU cache mode that defines the CPU mapping of the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/cpuCacheMode
func (o MTLVisibleFunctionTableObject) CpuCacheMode() MTLCPUCacheMode {
	
	rv := objc.Send[MTLCPUCacheMode](o.ID, objc.Sel("cpuCacheMode"))
	return rv
	}
// The location and access permissions of the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/storageMode
func (o MTLVisibleFunctionTableObject) StorageMode() MTLStorageMode {
	
	rv := objc.Send[MTLStorageMode](o.ID, objc.Sel("storageMode"))
	return rv
	}
// A mode that determines whether Metal tracks and synchronizes resource
// access.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/hazardTrackingMode
func (o MTLVisibleFunctionTableObject) HazardTrackingMode() MTLHazardTrackingMode {
	
	rv := objc.Send[MTLHazardTrackingMode](o.ID, objc.Sel("hazardTrackingMode"))
	return rv
	}
// The storage mode, CPU cache mode, and hazard tracking mode of the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/resourceOptions
func (o MTLVisibleFunctionTableObject) ResourceOptions() MTLResourceOptions {
	
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
func (o MTLVisibleFunctionTableObject) SetPurgeableState(state MTLPurgeableState) MTLPurgeableState {
	
	rv := objc.Send[MTLPurgeableState](o.ID, objc.Sel("setPurgeableState:"), state)
	return rv
	}
// The distance, in bytes, from the beginning of the heap to the first byte of
// the resource, if you allocated the resource on a heap.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/heapOffset
func (o MTLVisibleFunctionTableObject) HeapOffset() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("heapOffset"))
	return rv
	}
// The heap on which the resource is allocated, if any.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/heap
func (o MTLVisibleFunctionTableObject) Heap() MTLHeap {
	
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
func (o MTLVisibleFunctionTableObject) MakeAliasable() {
	
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
func (o MTLVisibleFunctionTableObject) IsAliasable() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAliasable"))
	return rv
	}
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/setOwnerWithIdentity:
func (o MTLVisibleFunctionTableObject) SetOwnerWithIdentity(task_id_token objectivec.IObject) int32 {
	
	rv := objc.Send[int32](o.ID, objc.Sel("setOwnerWithIdentity:"), task_id_token)
	return rv
	}

func (o MTLVisibleFunctionTableObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

