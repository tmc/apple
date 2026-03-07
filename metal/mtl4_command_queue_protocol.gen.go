// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// An abstraction representing a command queue that you use commit and synchronize command buffers and to perform other GPU operations.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue
type MTL4CommandQueue interface {
	objectivec.IObject

	// Returns the GPU device that the command queue belongs to.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/device
	Device() MTLDevice

	// Obtains this queue’s optional label for debugging purposes.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/label
	Label() string

	// Applies a residency set to a queue, which Metal applies to the queue’s command buffers as you commit them.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/addResidencySet(_:)
	AddResidencySet(residencySet MTLResidencySet)

	// Removes a residency set from a command queue’s list, which means Metal doesn’t apply it to the queue’s command buffers as you commit them.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/removeResidencySet(_:)
	RemoveResidencySet(residencySet MTLResidencySet)

	// Schedules a signal operation on the command queue to indicate when rendering to a Metal drawable is complete.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/signalDrawable(_:)
	SignalDrawable(drawable MTLDrawable)

	// Schedules an operation to signal a GPU event with a specific value after all GPU work prior to this point is complete.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/signalEvent(_:value:)
	SignalEventValue(event MTLEvent, value uint64)

	// Schedules a wait operation on the command queue to ensure the display is no longer using a specific Metal drawable.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/waitForDrawable(_:)
	WaitForDrawable(drawable MTLDrawable)

	// Schedules an operation to wait for a GPU event of a specific value before continuing to execute any future GPU work.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/waitForEvent(_:value:)
	WaitForEventValue(event MTLEvent, value uint64)

	// Applies multiple residency sets to a queue, which Metal applies to the queue’s command buffers as you commit them.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/addResidencySets:count:
	AddResidencySetsCount(residencySets []MTLResidencySet, count uint)

	// Enqueues an array of command buffers for execution.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/commit:count:
	CommitCount(commandBuffers []MTL4CommandBuffer, count uint)

	// Copies multiple offsets within a source placement sparse buffer to a destination placement sparse buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/copyBufferMappingsFromBuffer:toBuffer:operations:count:
	CopyBufferMappingsFromBufferToBufferOperationsCount(sourceBuffer MTLBuffer, destinationBuffer MTLBuffer, operations []MTL4CopySparseBufferMappingOperation, count uint)

	// Enqueues an array of command buffer instances for execution with a set of options.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/commit:count:options:
	CommitCountOptions(commandBuffers []MTL4CommandBuffer, count uint, options IMTL4CommitOptions)

	// Copies multiple regions within a source placement sparse texture to a destination placement sparse texture.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/copyTextureMappingsFromTexture:toTexture:operations:count:
	CopyTextureMappingsFromTextureToTextureOperationsCount(sourceTexture MTLTexture, destinationTexture MTLTexture, operations []MTL4CopySparseTextureMappingOperation, count uint)

	// Removes multiple residency sets from a command queue’s list, which means Metal doesn’t apply them to the queue’s command buffers as you commit them.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/removeResidencySets:count:
	RemoveResidencySetsCount(residencySets []MTLResidencySet, count uint)

	// Updates multiple regions within a placement sparse buffer to alias specific tiles from a Metal heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/updateBufferMappings:heap:operations:count:
	UpdateBufferMappingsHeapOperationsCount(buffer MTLBuffer, heap MTLHeap, operations []MTL4UpdateSparseBufferMappingOperation, count uint)

	// Updates multiple regions within a placement sparse texture to alias specific tiles of a Metal heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/updateTextureMappings:heap:operations:count:
	UpdateTextureMappingsHeapOperationsCount(texture MTLTexture, heap MTLHeap, operations []MTL4UpdateSparseTextureMappingOperation, count uint)
}



// MTL4CommandQueueObject wraps an existing Objective-C object that conforms to the MTL4CommandQueue protocol.
type MTL4CommandQueueObject struct {
	objectivec.Object
}
func (o MTL4CommandQueueObject) BaseObject() objectivec.Object {
	return o.Object
}



// MTL4CommandQueueObjectFromID constructs a [MTL4CommandQueueObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTL4CommandQueueObjectFromID(id objc.ID) MTL4CommandQueueObject {
	return MTL4CommandQueueObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Returns the GPU device that the command queue belongs to.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/device

func (o MTL4CommandQueueObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}

// Obtains this queue’s optional label for debugging purposes.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/label

func (o MTL4CommandQueueObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}

// Applies a residency set to a queue, which Metal applies to the queue’s
// command buffers as you commit them.
//
// residencySet: A residency set that contains resource allocations, such as [MTLBuffer],
// [MTLTexture], and [MTLHeap] instances.
//
// # Discussion
// 
// Each command queue can maintain a list of up to 32 different residency
// sets. See [Simplifying GPU resource management with residency sets] and
// [MTLResidencySet] for more information.
//
// [Simplifying GPU resource management with residency sets]: https://developer.apple.com/documentation/Metal/simplifying-gpu-resource-management-with-residency-sets
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/addResidencySet(_:)

func (o MTL4CommandQueueObject) AddResidencySet(residencySet MTLResidencySet) {
	
	objc.Send[struct{}](o.ID, objc.Sel("addResidencySet:"), residencySet)
	}

// Removes a residency set from a command queue’s list, which means Metal
// doesn’t apply it to the queue’s command buffers as you commit them.
//
// residencySet: A residency set that contains resource allocations, such as [MTLBuffer],
// [MTLTexture], and [MTLHeap] instances.
//
// # Discussion
// 
// The method doesn’t remove the residency set from command buffers the
// queue owns with an [Status] property that’s equal to
// [CommandBufferStatusCommitted] or [CommandBufferStatusScheduled].
// 
// See [Simplifying GPU resource management with residency sets] and
// [MTLResidencySet] for more information.
//
// [Simplifying GPU resource management with residency sets]: https://developer.apple.com/documentation/Metal/simplifying-gpu-resource-management-with-residency-sets
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/removeResidencySet(_:)

func (o MTL4CommandQueueObject) RemoveResidencySet(residencySet MTLResidencySet) {
	
	objc.Send[struct{}](o.ID, objc.Sel("removeResidencySet:"), residencySet)
	}

// Schedules a signal operation on the command queue to indicate when
// rendering to a Metal drawable is complete.
//
// drawable: [MTLDrawable] instance to signal.
//
// # Discussion
// 
// Signaling when rendering to a [MTLDrawable] instance is complete indicates
// that it’s safe to present it to the display.
// 
// You are responsible for calling this method after committing all command
// buffers that contain commands targeting this drawable, and before calling
// [Present], [PresentAtTime], or [PresentAfterMinimumDuration].
// 
// Metal doesn’t guarantee that command buffers you commit to the command
// queue after calling this method execute before presentation.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/signalDrawable(_:)

func (o MTL4CommandQueueObject) SignalDrawable(drawable MTLDrawable) {
	
	objc.Send[struct{}](o.ID, objc.Sel("signalDrawable:"), drawable)
	}

// Schedules an operation to signal a GPU event with a specific value after
// all GPU work prior to this point is complete.
//
// event: [MTLEvent] to signal.
//
// value: The value to signal the [MTLEvent] with.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/signalEvent(_:value:)

func (o MTL4CommandQueueObject) SignalEventValue(event MTLEvent, value uint64) {
	
	objc.Send[struct{}](o.ID, objc.Sel("signalEvent:value:"), event, value)
	}

// Schedules a wait operation on the command queue to ensure the display is no
// longer using a specific Metal drawable.
//
// drawable: [MTLDrawable] instance to signal.
//
// # Discussion
// 
// Use this method to ensure the display is no longer using a [MTLDrawable]
// instance before executing any subsequent commands.
// 
// This method returns immediately and doesn’t perform any synchronization
// on the current thread. You are responsible for calling this method before
// committing any command buffers containing commands that target this
// drawable.
// 
// Call this method multiple times if you commit your command buffers to
// multiple command queues.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/waitForDrawable(_:)

func (o MTL4CommandQueueObject) WaitForDrawable(drawable MTLDrawable) {
	
	objc.Send[struct{}](o.ID, objc.Sel("waitForDrawable:"), drawable)
	}

// Schedules an operation to wait for a GPU event of a specific value before
// continuing to execute any future GPU work.
//
// event: [MTLEvent] to wait on.
//
// value: The specific value to wait for.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/waitForEvent(_:value:)

func (o MTL4CommandQueueObject) WaitForEventValue(event MTLEvent, value uint64) {
	
	objc.Send[struct{}](o.ID, objc.Sel("waitForEvent:value:"), event, value)
	}

// Applies multiple residency sets to a queue, which Metal applies to the
// queue’s command buffers as you commit them.
//
// residencySets: A C array of residency sets, each of which contains resource allocations,
// such as [MTLBuffer], [MTLTexture], and [MTLHeap] instances.
//
// count: The number of elements in `residencySets`.
//
// # Discussion
// 
// Each command queue can maintain a list of up to 32 different residency
// sets. See [Simplifying GPU resource management with residency sets] and
// [MTLResidencySet] for more information.
//
// [Simplifying GPU resource management with residency sets]: https://developer.apple.com/documentation/Metal/simplifying-gpu-resource-management-with-residency-sets
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/addResidencySets:count:

func (o MTL4CommandQueueObject) AddResidencySetsCount(residencySets []MTLResidencySet, count uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("addResidencySets:count:"), objc.CArray(residencySets), count)
	}

// Enqueues an array of command buffers for execution.
//
// commandBuffers: An array of [MTL4CommandBuffer].
//
// count: The number of [MTL4CommandBuffer] instances in the `commandBuffers` array.
//
// # Discussion
// 
// The order in which you sort the command buffers in the array is meaningful,
// especially when it contains suspending/resuming render passes. A
// suspending/resuming render pass is a render pass you create by calling
// [RenderCommandEncoderWithDescriptorOptions], and provide
// [MTL4RenderEncoderOptionSuspending] or [MTL4RenderEncoderOptionResuming]
// for the `options` parameter.
// 
// If your command buffers contain suspend/resume render passes, ensure that
// the first command buffer only suspends, and the last one only resumes.
// Additionally, make sure that all intermediate command buffers are both
// suspending and resuming.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/commit:count:

func (o MTL4CommandQueueObject) CommitCount(commandBuffers []MTL4CommandBuffer, count uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("commit:count:"), objc.CArray(commandBuffers), count)
	}

// Copies multiple offsets within a source placement sparse buffer to a
// destination placement sparse buffer.
//
// sourceBuffer: The source placement sparse [MTLBuffer].
//
// destinationBuffer: The destination placement sparse [MTLBuffer].
//
// operations: An array of [MTL4CopySparseBufferMappingOperation] instances to perform.
// //
// [MTL4CopySparseBufferMappingOperation]: https://developer.apple.com/documentation/Metal/MTL4CopySparseBufferMappingOperation
//
// count: Number of operations to perform.
//
// # Discussion
// 
// You are responsible for ensuring the source destination sparse buffers have
// the same `placementSparsePageSize` when you create them via
// [NewBufferWithLengthOptionsPlacementSparsePageSize].
// 
// Additionally, you are responsible for ensuring both the source and
// destination sparse buffers don’t use the same aliased tiles at the same
// time.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/copyBufferMappingsFromBuffer:toBuffer:operations:count:

func (o MTL4CommandQueueObject) CopyBufferMappingsFromBufferToBufferOperationsCount(sourceBuffer MTLBuffer, destinationBuffer MTLBuffer, operations []MTL4CopySparseBufferMappingOperation, count uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("copyBufferMappingsFromBuffer:toBuffer:operations:count:"), sourceBuffer, destinationBuffer, objc.CArray(operations), count)
	}

// Enqueues an array of command buffer instances for execution with a set of
// options.
//
// commandBuffers: An array of [MTL4CommandBuffer].
//
// count: The number of [MTL4CommandBuffer] instances in the `commandBuffers` array.
//
// options: An instance of [MTL4CommitOptions] that configures the commit operation.
//
// # Discussion
// 
// Provide an [MTL4CommitOptions] instance to configure the commit operation.
// 
// The order in which you sort the command buffers in the array is meaningful,
// especially when it contains suspending/resuming render passes. A
// suspending/resuming render pass is a render pass you create by calling
// [RenderCommandEncoderWithDescriptorOptions], and provide
// [MTL4RenderEncoderOptionSuspending] or [MTL4RenderEncoderOptionResuming]
// for the `options` parameter.
// 
// If your command buffers contain suspend/resume render passes, ensure that
// the first command buffer only suspends, and the last one only resumes.
// Additionally, make sure that all intermediate command buffers are both
// suspending and resuming.
// 
// When you commit work from multiple threads, modifying and reusing the same
// options instance, you are responsible for externally synchronizing access
// to it.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/commit:count:options:

func (o MTL4CommandQueueObject) CommitCountOptions(commandBuffers []MTL4CommandBuffer, count uint, options IMTL4CommitOptions) {
	
	objc.Send[struct{}](o.ID, objc.Sel("commit:count:options:"), objc.CArray(commandBuffers), count, options)
	}

// Copies multiple regions within a source placement sparse texture to a
// destination placement sparse texture.
//
// sourceTexture: The source placement sparse [MTLTexture].
//
// destinationTexture: The destination placement sparse [MTLTexture].
//
// operations: An array of [MTL4CopySparseTextureMappingOperation] instances to perform.
// //
// [MTL4CopySparseTextureMappingOperation]: https://developer.apple.com/documentation/Metal/MTL4CopySparseTextureMappingOperation
//
// count: Number of operations to perform.
//
// # Discussion
// 
// You are responsible for ensuring the source and destination textures have
// the same [PlacementSparsePageSize].
// 
// Additionally, you are responsible for ensuring that the source and
// destination textures don’t use the same aliased tiles at the same time.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/copyTextureMappingsFromTexture:toTexture:operations:count:

func (o MTL4CommandQueueObject) CopyTextureMappingsFromTextureToTextureOperationsCount(sourceTexture MTLTexture, destinationTexture MTLTexture, operations []MTL4CopySparseTextureMappingOperation, count uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("copyTextureMappingsFromTexture:toTexture:operations:count:"), sourceTexture, destinationTexture, objc.CArray(operations), count)
	}

// Removes multiple residency sets from a command queue’s list, which means
// Metal doesn’t apply them to the queue’s command buffers as you commit
// them.
//
// residencySets: A C array of residency sets, each of which contains resource allocations,
// such as [MTLBuffer], [MTLTexture], and [MTLHeap] instances.
//
// count: The number of elements in `residencySets`.
//
// # Discussion
// 
// The method doesn’t remove the residency sets from command buffers the
// queue owns with an [Status] property that’s equal to
// [CommandBufferStatusCommitted] or [CommandBufferStatusScheduled].
// 
// See [Simplifying GPU resource management with residency sets] and
// [MTLResidencySet] for more information.
//
// [Simplifying GPU resource management with residency sets]: https://developer.apple.com/documentation/Metal/simplifying-gpu-resource-management-with-residency-sets
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/removeResidencySets:count:

func (o MTL4CommandQueueObject) RemoveResidencySetsCount(residencySets []MTLResidencySet, count uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("removeResidencySets:count:"), objc.CArray(residencySets), count)
	}

// Updates multiple regions within a placement sparse buffer to alias specific
// tiles from a Metal heap.
//
// buffer: A placement sparse [MTLBuffer].
//
// heap: An [MTLHeap] you allocate with type [HeapTypePlacement].
//
// operations: An array of [MTL4UpdateSparseBufferMappingOperation] instances to perform.
// //
// [MTL4UpdateSparseBufferMappingOperation]: https://developer.apple.com/documentation/Metal/MTL4UpdateSparseBufferMappingOperation
//
// count: Number of operations to perform.
//
// # Discussion
// 
// You can provide a `nil` parameter to the `heap` argument only when you
// perform unmap operations. Otherwise, you are responsible for ensuring
// parameter `heap` references an [MTLHeap] that has a
// [MaxCompatiblePlacementSparsePageSize] of at least the buffer’s
// `placementSparsePageSize` you assign when creating the sparse buffer via
// [NewBufferWithLengthOptionsPlacementSparsePageSize].
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/updateBufferMappings:heap:operations:count:

func (o MTL4CommandQueueObject) UpdateBufferMappingsHeapOperationsCount(buffer MTLBuffer, heap MTLHeap, operations []MTL4UpdateSparseBufferMappingOperation, count uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("updateBufferMappings:heap:operations:count:"), buffer, heap, objc.CArray(operations), count)
	}

// Updates multiple regions within a placement sparse texture to alias
// specific tiles of a Metal heap.
//
// texture: A placement sparse [MTLTexture].
//
// heap: [MTLHeap] you allocate with type [HeapTypePlacement].
//
// operations: An array of [MTL4UpdateSparseTextureMappingOperation] instances to perform.
// //
// [MTL4UpdateSparseTextureMappingOperation]: https://developer.apple.com/documentation/Metal/MTL4UpdateSparseTextureMappingOperation
//
// count: Number of operations to perform.
//
// # Discussion
// 
// You can provide a `nil` parameter to the `heap` argument only if when you
// perform unmap operations. Otherwise, you are responsible for ensuring the
// heap is non-nil and has a [MaxCompatiblePlacementSparsePageSize] of at
// least the texture’s [PlacementSparsePageSize].
// 
// When performing a sparse mapping update, you are responsible for issuing a
// barrier against stage [MTLStageResourceState].
// 
// You can determine the sparse texture tier by calling
// `MTLTexture/sparseTextureTier`.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandQueue/updateTextureMappings:heap:operations:count:

func (o MTL4CommandQueueObject) UpdateTextureMappingsHeapOperationsCount(texture MTLTexture, heap MTLHeap, operations []MTL4UpdateSparseTextureMappingOperation, count uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("updateTextureMappings:heap:operations:count:"), texture, heap, objc.CArray(operations), count)
	}











