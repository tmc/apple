// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A command buffer containing reusable commands, encoded either on the CPU or GPU.
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBuffer
type MTLIndirectCommandBuffer interface {
	objectivec.IObject
	MTLAllocation
	MTLResource

	// The number of commands contained in the indirect command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBuffer/size
	Size() uint

	// Gets the render command at the given index.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBuffer/indirectRenderCommandAt(_:)
	IndirectRenderCommandAtIndex(commandIndex uint) MTLIndirectRenderCommand

	// Gets the compute command at the given index.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBuffer/indirectComputeCommandAt(_:)
	IndirectComputeCommandAtIndex(commandIndex uint) MTLIndirectComputeCommand

	// GpuResourceID protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBuffer/gpuResourceID
	GpuResourceID() MTLResourceID

	// Resets a range of commands to their default state.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBuffer/resetWithRange:
	ResetWithRange(range_ foundation.NSRange)
}

// MTLIndirectCommandBufferObject wraps an existing Objective-C object that conforms to the MTLIndirectCommandBuffer protocol.
type MTLIndirectCommandBufferObject struct {
	objectivec.Object
}
func (o MTLIndirectCommandBufferObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLIndirectCommandBufferObjectFromID constructs a [MTLIndirectCommandBufferObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLIndirectCommandBufferObjectFromID(id objc.ID) MTLIndirectCommandBufferObject {
	return MTLIndirectCommandBufferObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The number of commands contained in the indirect command buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBuffer/size
func (o MTLIndirectCommandBufferObject) Size() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("size"))
	return rv
	}
// Gets the render command at the given index.
//
// commandIndex: The index of the command to retrieve.
//
// # Discussion
// 
// Call this method only if the indirect command buffer contains rendering
// commands.
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBuffer/indirectRenderCommandAt(_:)
func (o MTLIndirectCommandBufferObject) IndirectRenderCommandAtIndex(commandIndex uint) MTLIndirectRenderCommand {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("indirectRenderCommandAtIndex:"), commandIndex)
	return MTLIndirectRenderCommandObjectFromID(rv)
	}
// Gets the compute command at the given index.
//
// commandIndex: The index of the command to retrieve.
//
// # Discussion
// 
// Call this method only if the indirect command buffer contains compute
// commands.
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBuffer/indirectComputeCommandAt(_:)
func (o MTLIndirectCommandBufferObject) IndirectComputeCommandAtIndex(commandIndex uint) MTLIndirectComputeCommand {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("indirectComputeCommandAtIndex:"), commandIndex)
	return MTLIndirectComputeCommandObjectFromID(rv)
	}
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBuffer/gpuResourceID
func (o MTLIndirectCommandBufferObject) GpuResourceID() MTLResourceID {
	
	rv := objc.Send[MTLResourceID](o.ID, objc.Sel("gpuResourceID"))
	return rv
	}
// Resets a range of commands to their default state.
//
// range: The range of commands to reset. The range needs to fit inside the indirect
// command buffer’s extents.
//
// See: https://developer.apple.com/documentation/Metal/MTLIndirectCommandBuffer/resetWithRange:
func (o MTLIndirectCommandBufferObject) ResetWithRange(range_ foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("resetWithRange:"), range_)
	}
// The amount of memory, in byes, a resource consumes, such as for a buffer,
// texture, or heap.
//
// See: https://developer.apple.com/documentation/Metal/MTLAllocation/allocatedSize
func (o MTLIndirectCommandBufferObject) AllocatedSize() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("allocatedSize"))
	return rv
	}
// The device object that created the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/device
func (o MTLIndirectCommandBufferObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}
// A string that identifies the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/label
func (o MTLIndirectCommandBufferObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}
// The CPU cache mode that defines the CPU mapping of the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/cpuCacheMode
func (o MTLIndirectCommandBufferObject) CpuCacheMode() MTLCPUCacheMode {
	
	rv := objc.Send[MTLCPUCacheMode](o.ID, objc.Sel("cpuCacheMode"))
	return rv
	}
// The location and access permissions of the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/storageMode
func (o MTLIndirectCommandBufferObject) StorageMode() MTLStorageMode {
	
	rv := objc.Send[MTLStorageMode](o.ID, objc.Sel("storageMode"))
	return rv
	}
// A mode that determines whether Metal tracks and synchronizes resource
// access.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/hazardTrackingMode
func (o MTLIndirectCommandBufferObject) HazardTrackingMode() MTLHazardTrackingMode {
	
	rv := objc.Send[MTLHazardTrackingMode](o.ID, objc.Sel("hazardTrackingMode"))
	return rv
	}
// The storage mode, CPU cache mode, and hazard tracking mode of the resource.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/resourceOptions
func (o MTLIndirectCommandBufferObject) ResourceOptions() MTLResourceOptions {
	
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
func (o MTLIndirectCommandBufferObject) SetPurgeableState(state MTLPurgeableState) MTLPurgeableState {
	
	rv := objc.Send[MTLPurgeableState](o.ID, objc.Sel("setPurgeableState:"), state)
	return rv
	}
// The distance, in bytes, from the beginning of the heap to the first byte of
// the resource, if you allocated the resource on a heap.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/heapOffset
func (o MTLIndirectCommandBufferObject) HeapOffset() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("heapOffset"))
	return rv
	}
// The heap on which the resource is allocated, if any.
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/heap
func (o MTLIndirectCommandBufferObject) Heap() MTLHeap {
	
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
func (o MTLIndirectCommandBufferObject) MakeAliasable() {
	
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
func (o MTLIndirectCommandBufferObject) IsAliasable() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAliasable"))
	return rv
	}
//
// See: https://developer.apple.com/documentation/Metal/MTLResource/setOwnerWithIdentity:
func (o MTLIndirectCommandBufferObject) SetOwnerWithIdentity(task_id_token objectivec.IObject) int32 {
	
	rv := objc.Send[int32](o.ID, objc.Sel("setOwnerWithIdentity:"), task_id_token)
	return rv
	}

func (o MTLIndirectCommandBufferObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

