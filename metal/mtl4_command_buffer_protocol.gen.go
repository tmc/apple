// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// Records a sequence of GPU commands.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer
type MTL4CommandBuffer interface {
	objectivec.IObject

	// Returns the GPU device that this command buffer belongs to.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/device
	Device() MTLDevice

	// Assigns an optional label with this command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/label
	Label() string

	// Prepares a command buffer for encoding.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/beginCommandBuffer(allocator:)
	BeginCommandBufferWithAllocator(allocator MTL4CommandAllocator)

	// Prepares a command buffer for encoding with additional options.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/beginCommandBuffer(allocator:options:)
	BeginCommandBufferWithAllocatorOptions(allocator MTL4CommandAllocator, options IMTL4CommandBufferOptions)

	// Closes a command buffer to prepare it for submission to a command queue.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/endCommandBuffer()
	EndCommandBuffer()

	// Creates a compute command encoder.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/makeComputeCommandEncoder()
	ComputeCommandEncoder() MTL4ComputeCommandEncoder

	// Creates a machine learning command encoder.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/makeMachineLearningCommandEncoder()
	MachineLearningCommandEncoder() MTL4MachineLearningCommandEncoder

	// Creates a render command encoder from a render pass descriptor with additional options.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/makeRenderCommandEncoder(descriptor:options:)
	RenderCommandEncoderWithDescriptorOptions(descriptor IMTL4RenderPassDescriptor, options MTL4RenderEncoderOptions) MTL4RenderCommandEncoder

	// Pops the latest string from the stack of debug groups for this command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/popDebugGroup()
	PopDebugGroup()

	// Pushes a string onto a stack of debug groups for this command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/pushDebugGroup(_:)
	PushDebugGroup(string_ string)

	// Applies a residency set to a command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/useResidencySet(_:)
	UseResidencySet(residencySet MTLResidencySet)

	// Writes a GPU timestamp into the given counter heap.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/writeTimestamp(counterHeap:index:)
	WriteTimestampIntoHeapAtIndex(counterHeap MTL4CounterHeap, index uint)

	// Creates a render command encoder from a render pass descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/renderCommandEncoderWithDescriptor:
	RenderCommandEncoderWithDescriptor(descriptor IMTL4RenderPassDescriptor) MTL4RenderCommandEncoder

	// Encodes a command that resolves an opaque counter heap into a buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/resolveCounterHeap:withRange:intoBuffer:waitFence:updateFence:
	ResolveCounterHeapWithRangeIntoBufferWaitFenceUpdateFence(counterHeap MTL4CounterHeap, range_ foundation.NSRange, bufferRange MTL4BufferRange, fenceToWait MTLFence, fenceToUpdate MTLFence)

	// Applies multiple residency sets to a command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/useResidencySets:count:
	UseResidencySetsCount(residencySets []MTLResidencySet, count uint)

	// Assigns an optional label with this command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/label
	SetLabel(value string)
}

// MTL4CommandBufferObject wraps an existing Objective-C object that conforms to the MTL4CommandBuffer protocol.
type MTL4CommandBufferObject struct {
	objectivec.Object
}
func (o MTL4CommandBufferObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTL4CommandBufferObjectFromID constructs a [MTL4CommandBufferObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTL4CommandBufferObjectFromID(id objc.ID) MTL4CommandBufferObject {
	return MTL4CommandBufferObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the GPU device that this command buffer belongs to.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/device
func (o MTL4CommandBufferObject) Device() MTLDevice {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
	}
// Assigns an optional label with this command buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/label
func (o MTL4CommandBufferObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}
// Prepares a command buffer for encoding.
//
// allocator: [MTL4CommandAllocator] to attach to.
//
// # Discussion
// 
// Attaches the command buffer to the specified [MTL4CommandAllocator] and
// declares that the application is ready to encode commands into the command
// buffer.
// 
// Command allocators only service a single command buffer at a time. If you
// need to issue multiple calls to this method simultaneously, for example, in
// a multi-threaded command encoding scenario, create multiple instances of
// [MTLCommandAllocator] and use one for each call.
// 
// You can safely reuse command allocators after ending the command buffer
// using it by calling [EndCommandBuffer].
// 
// After calling this method, any prior calls to [UseResidencySet] and
// [UseResidencySetsCount] on this command buffer instance no longer apply.
// Make sure to call these methods again to signal your residency requirements
// to Metal.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/beginCommandBuffer(allocator:)
func (o MTL4CommandBufferObject) BeginCommandBufferWithAllocator(allocator MTL4CommandAllocator) {
	
	objc.Send[struct{}](o.ID, objc.Sel("beginCommandBufferWithAllocator:"), allocator)
	}
// Prepares a command buffer for encoding with additional options.
//
// allocator: [MTL4CommandAllocator] to attach to.
//
// options: [MTL4CommandBufferOptions] to configure the command buffer.
//
// # Discussion
// 
// Attaches the command buffer to the specified [MTL4CommandAllocator] and
// declares that the application is ready to encode commands into the command
// buffer.
// 
// Command allocators only service a single command buffer at a time. If you
// need to issue multiple calls to this method simultaneously, for example, in
// a multi-threaded command encoding scenario, create multiple instances of
// [MTLCommandAllocator] and use one for each call.
// 
// You can safely reuse command allocators after ending the command buffer
// using it by calling [EndCommandBuffer].
// 
// After calling this method, any prior calls to [UseResidencySet] and
// [UseResidencySetsCount] on this command buffer instance no longer apply.
// Make sure to call these methods again to signal your residency requirements
// to Metal.
// 
// The options you provide configure the command buffer only until the command
// buffer ends, in the next call to [EndCommandBuffer].
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/beginCommandBuffer(allocator:options:)
func (o MTL4CommandBufferObject) BeginCommandBufferWithAllocatorOptions(allocator MTL4CommandAllocator, options IMTL4CommandBufferOptions) {
	
	objc.Send[struct{}](o.ID, objc.Sel("beginCommandBufferWithAllocator:options:"), allocator, options)
	}
// Closes a command buffer to prepare it for submission to a command queue.
//
// # Discussion
// 
// Explicitly ending the command buffer allows you to reuse the
// [MTL4CommandAllocator] to start servicing other command buffers. It is an
// error to call `commit` on a command buffer previously recording before
// calling this method.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/endCommandBuffer()
func (o MTL4CommandBufferObject) EndCommandBuffer() {
	
	objc.Send[struct{}](o.ID, objc.Sel("endCommandBuffer"))
	}
// Creates a compute command encoder.
//
// # Return Value
// 
// The created [MTL4ComputeCommandEncoder] instance, or `nil` if the function
// fails.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/makeComputeCommandEncoder()
func (o MTL4CommandBufferObject) ComputeCommandEncoder() MTL4ComputeCommandEncoder {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("computeCommandEncoder"))
	return MTL4ComputeCommandEncoderObjectFromID(rv)
	}
// Creates a machine learning command encoder.
//
// # Return Value
// 
// The created [MTL4MachineLearningCommandEncoder] instance , or `nil` if the
// function fails.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/makeMachineLearningCommandEncoder()
func (o MTL4CommandBufferObject) MachineLearningCommandEncoder() MTL4MachineLearningCommandEncoder {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("machineLearningCommandEncoder"))
	return MTL4MachineLearningCommandEncoderObjectFromID(rv)
	}
// Creates a render command encoder from a render pass descriptor with
// additional options.
//
// descriptor: Descriptor for the render pass.
//
// options: [MTL4RenderEncoderOptions] instance that provide render pass options.
// //
// [MTL4RenderEncoderOptions]: https://developer.apple.com/documentation/Metal/MTL4RenderEncoderOptions
//
// # Return Value
// 
// The created [MTL4RenderCommandEncoder] instance, or `nil` if the function
// fails.
//
// # Discussion
// 
// This method creates a render command encoder to encode a render pass,
// whilst providing you the option to define some render pass characteristics
// via an instance of [MTL4RenderEncoderOptions].
// 
// Use these options to configure suspending/resuming render command encoders,
// which allow you to encode render passes from multiple threads
// simultaneously.
//
// [MTL4RenderEncoderOptions]: https://developer.apple.com/documentation/Metal/MTL4RenderEncoderOptions
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/makeRenderCommandEncoder(descriptor:options:)
func (o MTL4CommandBufferObject) RenderCommandEncoderWithDescriptorOptions(descriptor IMTL4RenderPassDescriptor, options MTL4RenderEncoderOptions) MTL4RenderCommandEncoder {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("renderCommandEncoderWithDescriptor:options:"), descriptor, options)
	return MTL4RenderCommandEncoderObjectFromID(rv)
	}
// Pops the latest string from the stack of debug groups for this command
// buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/popDebugGroup()
func (o MTL4CommandBufferObject) PopDebugGroup() {
	
	objc.Send[struct{}](o.ID, objc.Sel("popDebugGroup"))
	}
// Pushes a string onto a stack of debug groups for this command buffer.
//
// string: The string to push.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/pushDebugGroup(_:)
func (o MTL4CommandBufferObject) PushDebugGroup(string_ string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("pushDebugGroup:"), objc.String(string_))
	}
// Applies a residency set to a command buffer.
//
// residencySet: A residency set that contains resource allocations, such as [MTLBuffer],
// [MTLTexture], and [MTLHeap] instances.
//
// # Discussion
// 
// Each command buffer can maintain a list of up to 32 different residency
// sets. See [Simplifying GPU resource management with residency sets] and
// [MTLResidencySet] for more information.
//
// [Simplifying GPU resource management with residency sets]: https://developer.apple.com/documentation/Metal/simplifying-gpu-resource-management-with-residency-sets
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/useResidencySet(_:)
func (o MTL4CommandBufferObject) UseResidencySet(residencySet MTLResidencySet) {
	
	objc.Send[struct{}](o.ID, objc.Sel("useResidencySet:"), residencySet)
	}
// Writes a GPU timestamp into the given counter heap.
//
// counterHeap: [MTL4CounterHeap] to write the timestamp into.
//
// index: The index within the [MTL4CounterHeap] that Metal writes the timestamp to.
//
// # Discussion
// 
// This method captures a timestamp after work prior to this command in the
// command buffer is complete. Work after this call may or may not have
// started.
// 
// You are responsible for ensuring the `counterHeap` is of type
// [MTL4CounterHeapTypeTimestamp].
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/writeTimestamp(counterHeap:index:)
func (o MTL4CommandBufferObject) WriteTimestampIntoHeapAtIndex(counterHeap MTL4CounterHeap, index uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("writeTimestampIntoHeap:atIndex:"), counterHeap, index)
	}
// Creates a render command encoder from a render pass descriptor.
//
// descriptor: Descriptor for the render pass.
//
// # Return Value
// 
// The created [MTL4RenderCommandEncoder] instance, or `nil` if the function
// failed.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/renderCommandEncoderWithDescriptor:
func (o MTL4CommandBufferObject) RenderCommandEncoderWithDescriptor(descriptor IMTL4RenderPassDescriptor) MTL4RenderCommandEncoder {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("renderCommandEncoderWithDescriptor:"), descriptor)
	return MTL4RenderCommandEncoderObjectFromID(rv)
	}
// Encodes a command that resolves an opaque counter heap into a buffer.
//
// counterHeap: A heap the command resolves.
//
// range: A range of index values within the heap the command resolves.
//
// bufferRange: The buffer the command saves the data it resolves into.
//
// fenceToWait: A fence the GPU waits for before starting, if applicable; otherwise `nil`.
//
// fenceToUpdate: A fence the system updates after the command finishes resolving the data;
// otherwise `nil`.
//
// # Discussion
// 
// The command this method encodes converts the data within `counterHeap` into
// a common format and stores it into the `bufferRange` parameter.
// 
// The command places each entry in the counter heap within `range`
// sequentially, starting at `alignedOffset`. Each entry needs to be a fixed
// size that you can query by calling the [SizeOfCounterHeapEntry] method.
// 
// This command runs during the [MTLStageBlit] stage of the GPU timeline.
// Barrier against this stage to ensure the data is present in the resolve
// buffer parameter before you access it.
// 
// Similarly, your app needs to synchronize any GPU accesses to `bufferRange`
// after the command completes with barrier.
// 
// If your app needs to access `bufferRange` from the CPU, signal an
// [MTLSharedEvent] to notify the CPU when it’s ready. Alternatively, you
// can resolve the heap’s data from the CPU by calling the heap’s
// [ResolveCounterRange] method.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/resolveCounterHeap:withRange:intoBuffer:waitFence:updateFence:
func (o MTL4CommandBufferObject) ResolveCounterHeapWithRangeIntoBufferWaitFenceUpdateFence(counterHeap MTL4CounterHeap, range_ foundation.NSRange, bufferRange MTL4BufferRange, fenceToWait MTLFence, fenceToUpdate MTLFence) {
	
	objc.Send[struct{}](o.ID, objc.Sel("resolveCounterHeap:withRange:intoBuffer:waitFence:updateFence:"), counterHeap, range_, bufferRange, fenceToWait, fenceToUpdate)
	}
// Applies multiple residency sets to a command buffer.
//
// residencySets: A C array of residency sets, each of which contains resource allocations,
// such as [MTLBuffer], [MTLTexture], and [MTLHeap] instances.
//
// count: The number of elements in `residencySets`.
//
// # Discussion
// 
// Each command buffer can maintain a list of up to 32 different residency
// sets. See [Simplifying GPU resource management with residency sets] and
// [MTLResidencySet] for more information.
//
// [Simplifying GPU resource management with residency sets]: https://developer.apple.com/documentation/Metal/simplifying-gpu-resource-management-with-residency-sets
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandBuffer/useResidencySets:count:
func (o MTL4CommandBufferObject) UseResidencySetsCount(residencySets []MTLResidencySet, count uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("useResidencySets:count:"), objc.CArray(residencySets), count)
	}

func (o MTL4CommandBufferObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

