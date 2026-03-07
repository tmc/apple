// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A container that stores a sequence of GPU commands that you encode into it.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer
type MTLCommandBuffer interface {
	objectivec.IObject

	// Applies a residency set to a command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/useResidencySet(_:)
	UseResidencySet(residencySet MTLResidencySet)

	// Encodes a command into the command buffer that pauses the GPU from running the buffer’s subsequent passes until the event equals or exceeds a value.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/encodeWaitForEvent(_:value:)
	EncodeWaitForEventValue(event MTLEvent, value uint64)

	// Encodes a command that updates an event’s value, which can clear the GPU to run passes from other command buffers waiting for the event.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/encodeSignalEvent(_:value:)
	EncodeSignalEventValue(event MTLEvent, value uint64)

	// Presents a drawable as early as possible.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/present(_:)
	PresentDrawable(drawable MTLDrawable)

	// Presents a drawable at a specific time.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/present(_:atTime:)
	PresentDrawableAtTime(drawable MTLDrawable, presentationTime float64)

	// Presents a drawable after the system presents the previous drawable for an amount of time.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/present(_:afterMinimumDuration:)
	PresentDrawableAfterMinimumDuration(drawable MTLDrawable, duration float64)

	// Registers a completion handler the GPU device calls immediately after it schedules the command buffer to run on the GPU.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/addScheduledHandler(_:)
	AddScheduledHandler(block MTLCommandBufferHandler)

	// Registers a completion handler the GPU device calls immediately after the GPU finishes running the commands in the command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/addCompletedHandler(_:)
	AddCompletedHandler(block MTLCommandBufferHandler)

	// Reserves the next available place for the command buffer in its command queue.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/enqueue()
	Enqueue()

	// Submits the command buffer to run on the GPU.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/commit()
	Commit()

	// Blocks the current thread until the command queue schedules the buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/waitUntilScheduled()
	WaitUntilScheduled()

	// Blocks the current thread until the GPU finishes executing the command buffer and all of its completion handlers.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/waitUntilCompleted()
	WaitUntilCompleted()

	// The command buffer’s current state.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/status
	Status() MTLCommandBufferStatus

	// Creates a ray-tracing acceleration structure command encoder that uses default settings.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/makeAccelerationStructureCommandEncoder()
	AccelerationStructureCommandEncoder() MTLAccelerationStructureCommandEncoder

	// Creates a block information transfer (blit) encoder.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/makeBlitCommandEncoder()
	BlitCommandEncoder() MTLBlitCommandEncoder

	// Creates a block information transfer (blit) encoder from a descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/makeBlitCommandEncoder(descriptor:)
	BlitCommandEncoderWithDescriptor(blitPassDescriptor IMTLBlitPassDescriptor) MTLBlitCommandEncoder

	// Creates a ray-tracing acceleration structure command encoder from a descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/makeAccelerationStructureCommandEncoder(descriptor:)
	AccelerationStructureCommandEncoderWithDescriptor(descriptor IMTLAccelerationStructurePassDescriptor) MTLAccelerationStructureCommandEncoder

	// Creates a compute command encoder that uses default settings.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/makeComputeCommandEncoder()
	ComputeCommandEncoder() MTLComputeCommandEncoder

	// Creates a compute command encoder from a descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/makeComputeCommandEncoder(descriptor:)
	ComputeCommandEncoderWithDescriptor(computePassDescriptor IMTLComputePassDescriptor) MTLComputeCommandEncoder

	// Creates a parallel render command encoder from a descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/makeParallelRenderCommandEncoder(descriptor:)
	ParallelRenderCommandEncoderWithDescriptor(renderPassDescriptor IMTLRenderPassDescriptor) MTLParallelRenderCommandEncoder

	// Creates a compute command encoder with a dispatch type.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/makeComputeCommandEncoder(dispatchType:)
	ComputeCommandEncoderWithDispatchType(dispatchType MTLDispatchType) MTLComputeCommandEncoder

	// Creates a resource state command encoder that uses default settings.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/makeResourceStateCommandEncoder()
	ResourceStateCommandEncoder() MTLResourceStateCommandEncoder

	// Marks the end of a debug group and, if applicable, restores the previous group from a stack.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/popDebugGroup()
	PopDebugGroup()

	// Creates a render command encoder from a descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/makeRenderCommandEncoder(descriptor:)
	RenderCommandEncoderWithDescriptor(renderPassDescriptor IMTLRenderPassDescriptor) MTLRenderCommandEncoder

	// Marks the beginning of a debug group and gives it an identifying label, which temporarily replaces the previous group, if applicable.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/pushDebugGroup(_:)
	PushDebugGroup(string_ string)

	// Creates a resource state command encoder from a descriptor.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/resourceStateCommandEncoder(with:)
	ResourceStateCommandEncoderWithDescriptor(resourceStatePassDescriptor IMTLResourceStatePassDescriptor) MTLResourceStateCommandEncoder

	// Applies multiple residency sets to a command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/useResidencySets:count:
	UseResidencySetsCount(residencySets []MTLResidencySet, count uint)

	// The GPU device that indirectly owns the command buffer because you create it from a command queue the device also owns.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/device
	Device() MTLDevice

	// The command queue that creates the command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/commandQueue
	CommandQueue() MTLCommandQueue

	// A description of an error when the GPU encounters an issue as it runs the command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/error
	Error() foundation.INSError

	// The host time, in seconds, when the GPU finishes execution of the command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/gpuEndTime
	GPUEndTime() float64

	// Settings that determine which information the command buffer records about execution errors, and how it does it.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/errorOptions
	ErrorOptions() MTLCommandBufferErrorOption

	// The messages the command buffer records as the GPU runs its commands.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/logs
	Logs() MTLLogContainer

	// An optional name that can help you identify the command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/label
	Label() string
	SetLabel(value string)

	// The host time, in seconds, when the CPU finishes scheduling the command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/kernelEndTime
	KernelEndTime() float64

	// The host time, in seconds, when the CPU begins to schedule the command buffer.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/kernelStartTime
	KernelStartTime() float64

	// The host time, in seconds, when the GPU starts command buffer execution.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/gpuStartTime
	GPUStartTime() float64

	// A Boolean value that indicates whether the command buffer maintains strong references to the resources it uses.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/retainedReferences
	RetainedReferences() bool
}



// MTLCommandBufferObject wraps an existing Objective-C object that conforms to the MTLCommandBuffer protocol.
type MTLCommandBufferObject struct {
	objectivec.Object
}
func (o MTLCommandBufferObject) BaseObject() objectivec.Object {
	return o.Object
}



// MTLCommandBufferObjectFromID constructs a [MTLCommandBufferObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLCommandBufferObjectFromID(id objc.ID) MTLCommandBufferObject {
	return MTLCommandBufferObject{
		Object: objectivec.ObjectFromID(id),
	}
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
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/useResidencySet(_:)

func (o MTLCommandBufferObject) UseResidencySet(residencySet MTLResidencySet) {
	
	objc.Send[struct{}](o.ID, objc.Sel("useResidencySet:"), residencySet)
	}

// Encodes a command into the command buffer that pauses the GPU from running
// the buffer’s subsequent passes until the event equals or exceeds a value.
//
// event: An [MTLEvent] instance the GPU driver waits for between passes as it runs
// the command buffer.
// 
// If `event` is an [MTLSharedEvent] instance, a command buffer from any GPU
// device can signal this command buffer. Otherwise, only command buffers from
// the same GPU device can signal this command buffer.
//
// value: The event’s smallest value that allows the GPU to continue running the
// remaining passes in the command buffer.
//
// # Discussion
// 
// This method prevents the GPU from starting the next pass in the command
// buffer until another command buffer signals `event` (see
// [EncodeSignalEventValue].
// 
// A command buffer can instruct the GPU to wait for an event only between
// passes, not within a pass. If a command buffer has an active encoder,
// finish using the encoder, call its [EndEncoding] method, and then call this
// method before creating another encoder.
// 
// When the GPU device reaches the wait command that this method encodes into
// the buffer, it checks the event’s current value. If the event’s value
// — which increases monotonically — is less than the `value` parameter,
// the GPU waits before running the next pass in the buffer. The GPU starts
// the next pass when the event signals a value that’s equal to or greater
// than the `value` parameter (see [EncodeSignalEventValue]). However, If the
// event’s value is already greater than or equal to the `value` parameter,
// the GPU immediately starts the next pass without waiting.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/encodeWaitForEvent(_:value:)

func (o MTLCommandBufferObject) EncodeWaitForEventValue(event MTLEvent, value uint64) {
	
	objc.Send[struct{}](o.ID, objc.Sel("encodeWaitForEvent:value:"), event, value)
	}

// Encodes a command that updates an event’s value, which can clear the GPU
// to run passes from other command buffers waiting for the event.
//
// event: An [MTLEvent] instance the GPU driver signals between passes as it runs the
// command buffer.
// 
// If `event` is an [MTLSharedEvent] instance, the update:
// 
// - Signals any command buffers waiting for the shared event, including those
// on other GPU devices - Invokes any notification handlers waiting for the
// shared event (see [NotifyListenerAtValueBlock])
// 
// Otherwise, the method can signal only command buffers from the same GPU
// device.
//
// value: A value that’s greater than or equal to the event’s current value;
// otherwise, the command has no effect.
//
// # Discussion
// 
// The method can unblock one or more command buffers that are waiting for
// `event`, including those in other command queues (see
// [EncodeWaitForEventValue]).
// 
// A command buffer can signal an event only between passes, not within a
// pass. If a command buffer has an active encoder, finish using the encoder,
// call its [EndEncoding] method, and then call this method before creating
// another encoder.
// 
// When the GPU device reaches the signal command that this method encodes,
// Metal updates the event after the GPU finishes the buffer’s prior
// commands. Updating the event’s value can signal any command buffer
// that’s waiting for a value equal to or less than the `value` parameter.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/encodeSignalEvent(_:value:)

func (o MTLCommandBufferObject) EncodeSignalEventValue(event MTLEvent, value uint64) {
	
	objc.Send[struct{}](o.ID, objc.Sel("encodeSignalEvent:value:"), event, value)
	}

// Presents a drawable as early as possible.
//
// drawable: An [MTLDrawable] instance that contains a texture the system can show on a
// display.
//
// # Discussion
// 
// This convenience method calls the drawable’s [Present] method after the
// command queue schedules the command buffer for execution. The command
// buffer does this by adding a completion handler by calling its own
// [AddScheduledHandler] method for you.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/present(_:)

func (o MTLCommandBufferObject) PresentDrawable(drawable MTLDrawable) {
	
	objc.Send[struct{}](o.ID, objc.Sel("presentDrawable:"), drawable)
	}

// Presents a drawable at a specific time.
//
// drawable: An [MTLDrawable] instance that contains a texture the system can show on a
// display.
//
// presentationTime: The Mach absolute time, in seconds, that you want to present the drawable.
//
// # Discussion
// 
// This convenience method calls the drawable’s [PresentAtTime] method after
// the command queue schedules the command buffer for execution. The command
// buffer does this by adding a completion handler by calling its own
// [AddScheduledHandler] method for you.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/present(_:atTime:)

func (o MTLCommandBufferObject) PresentDrawableAtTime(drawable MTLDrawable, presentationTime float64) {
	
	objc.Send[struct{}](o.ID, objc.Sel("presentDrawable:atTime:"), drawable, presentationTime)
	}

// Presents a drawable after the system presents the previous drawable for an
// amount of time.
//
// drawable: An [MTLDrawable] instance that contains a texture the system can show on a
// display.
//
// duration: The shortest display time you want the system to give to the previous
// drawable before presenting this one.
//
// # Discussion
// 
// This convenience method calls the drawable’s
// [PresentAfterMinimumDuration] method after the command queue schedules the
// command buffer for execution. The command buffer does this by adding a
// completion handler by calling its own [AddScheduledHandler] method for you.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/present(_:afterMinimumDuration:)

func (o MTLCommandBufferObject) PresentDrawableAfterMinimumDuration(drawable MTLDrawable, duration float64) {
	
	objc.Send[struct{}](o.ID, objc.Sel("presentDrawable:afterMinimumDuration:"), drawable, duration)
	}

// Registers a completion handler the GPU device calls immediately after it
// schedules the command buffer to run on the GPU.
//
// block: A Swift closure or an Objective-C block that Metal calls after it schedules
// the command buffer to run on the GPU.
//
// # Discussion
// 
// You can register one or more scheduling completion handlers for the same
// command buffer. The GPU device’s driver (on the CPU) calls the completion
// handlers after it finishes scheduling the command buffer to run on the GPU.
// 
// The GPU device schedules each command buffer — along with tasks from
// other command buffers — after it identifies the command buffer’s
// dependencies. At that time, the GPU device sets the command buffer’s
// status to [CommandBufferStatusScheduled] and calls your completion handler.
// 
// You can use the command buffer’s [kernelEndTime] and [kernelStartTime]
// properties to calculate how much time the CPU spends scheduling the command
// buffer.
//
// [kernelEndTime]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/kernelEndTime
// [kernelStartTime]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/kernelStartTime
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/addScheduledHandler(_:)

func (o MTLCommandBufferObject) AddScheduledHandler(block MTLCommandBufferHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("addScheduledHandler:"), block)
	}

// Registers a completion handler the GPU device calls immediately after the
// GPU finishes running the commands in the command buffer.
//
// block: A Swift closure or an Objective-C block that Metal calls after the GPU
// finishes running the commands in the command buffer.
//
// # Discussion
// 
// You can register one or more completion handlers for the same command
// buffer. The GPU device’s driver (on the CPU) calls the completion
// handlers after the GPU finishes executing the command buffer.
// 
// For example, you can use the command buffer’s [gpuEndTime] and
// [gpuStartTime] properties to calculate how much time the GPU spends running
// the command buffer.
// 
// The completion handler is also a good place to check the [Status] property
// to determine whether the GPU successfully completes the buffer’s
// commands. If the status is equal to [CommandBufferStatusError], you can
// investigate further by checking the [error] and log properties for more
// details about the issue. See [Command buffer debugging] for more methods
// and properties that can help you isolate the issue.
//
// [Command buffer debugging]: https://developer.apple.com/documentation/Metal/command-buffer-debugging
// [error]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/error
// [gpuEndTime]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/gpuEndTime
// [gpuStartTime]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/gpuStartTime
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/addCompletedHandler(_:)

func (o MTLCommandBufferObject) AddCompletedHandler(block MTLCommandBufferHandler) {
	
	objc.Send[struct{}](o.ID, objc.Sel("addCompletedHandler:"), block)
	}

// Reserves the next available place for the command buffer in its command
// queue.
//
// # Discussion
// 
// The [Enqueue] method adds the command buffer to the [MTLCommandQueue]
// instance that owns it, but doesn’t commit the command buffer to run on
// the GPU. You can call the command buffer’s [Commit] method at a later
// time when it’s ready to run on the GPU. You can call a command buffer’s
// [Enqueue] method any time before you call [Commit], including before,
// after, or as you encode commands to it.
// 
// Enqueuing your command buffers first gives you the flexibility to arrange
// their relative order of execution before encoding commands to any of them.
// This approach lets you potentially encode each command buffer on a thread,
// in parallel, instead of encoding them one by one on a single thread. The
// order in which each worker thread finishes encoding and commits its command
// buffer doesn’t matter when you enqueue them in order before committing.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/enqueue()

func (o MTLCommandBufferObject) Enqueue() {
	
	objc.Send[struct{}](o.ID, objc.Sel("enqueue"))
	}

// Submits the command buffer to run on the GPU.
//
// # Discussion
// 
// The [Commit] method sends the command buffer to the [MTLCommandQueue]
// instance that owns it, which then schedules it to run on the GPU. If your
// app calls [Commit] for a command buffer that isn’t enqueued, the method
// effectively calls [Enqueue] for you.
// 
// The [Commit] method has several restrictions, including:
// 
// - You can commit a command buffer to its command queue only one time. - You
// can only commit a command buffer when it doesn’t have an active encoder
// (see [MTLCommandBuffer] and [MTLCommandEncoder]). - You can’t encode
// additional commands to a command buffer after you commit it. - You can’t
// call the [AddScheduledHandler] or [AddCompletedHandler] methods after you
// commit a command buffer.
// 
// The GPU starts the command buffer after it starts any command buffers that
// are ahead of it in the same command queue.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/commit()

func (o MTLCommandBufferObject) Commit() {
	
	objc.Send[struct{}](o.ID, objc.Sel("commit"))
	}

// Blocks the current thread until the command queue schedules the buffer.
//
// # Discussion
// 
// This method returns after the following events:
// 
// - The command queue (see [Status] and [CommandBufferStatusScheduled]) the
// command buffer to run on the GPU. - The command buffer invokes all the
// completion handlers your app submits with [AddScheduledHandler].
// 
// Use the [WaitUntilCompleted] method to check for completion of the
// scheduled work.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/waitUntilScheduled()

func (o MTLCommandBufferObject) WaitUntilScheduled() {
	
	objc.Send[struct{}](o.ID, objc.Sel("waitUntilScheduled"))
	}

// Blocks the current thread until the GPU finishes executing the command
// buffer and all of its completion handlers.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/waitUntilCompleted()

func (o MTLCommandBufferObject) WaitUntilCompleted() {
	
	objc.Send[struct{}](o.ID, objc.Sel("waitUntilCompleted"))
	}

// The command buffer’s current state.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/status

func (o MTLCommandBufferObject) Status() MTLCommandBufferStatus {
	
	rv := objc.Send[MTLCommandBufferStatus](o.ID, objc.Sel("status"))
	return rv
	}

// Creates a ray-tracing acceleration structure command encoder that uses
// default settings.
//
// # Discussion
// 
// Use an [MTLAccelerationStructureCommandEncoder] instance’s methods to set
// up a single ray-tracing pass.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/makeAccelerationStructureCommandEncoder()

func (o MTLCommandBufferObject) AccelerationStructureCommandEncoder() MTLAccelerationStructureCommandEncoder {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accelerationStructureCommandEncoder"))
	return MTLAccelerationStructureCommandEncoderObjectFromID(rv)
	}

// Creates a block information transfer (blit) encoder.
//
// # Discussion
// 
// Use an [MTLBlitCommandEncoder] instance’s methods to create a block
// information transfer (blit) pass that quickly copies memory between a GPU
// device’s resources.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/makeBlitCommandEncoder()

func (o MTLCommandBufferObject) BlitCommandEncoder() MTLBlitCommandEncoder {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("blitCommandEncoder"))
	return MTLBlitCommandEncoderObjectFromID(rv)
	}

// Creates a block information transfer (blit) encoder from a descriptor.
//
// blitPassDescriptor: An [MTLBlitPassDescriptor] instance that configures the
// [MTLBlitCommandEncoder] the method returns.
//
// # Discussion
// 
// Use an [MTLBlitCommandEncoder] instance’s methods to create a block
// information transfer (blit) pass that quickly copies memory between a GPU
// device’s resources.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/makeBlitCommandEncoder(descriptor:)

func (o MTLCommandBufferObject) BlitCommandEncoderWithDescriptor(blitPassDescriptor IMTLBlitPassDescriptor) MTLBlitCommandEncoder {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("blitCommandEncoderWithDescriptor:"), blitPassDescriptor)
	return MTLBlitCommandEncoderObjectFromID(rv)
	}

// Creates a ray-tracing acceleration structure command encoder from a
// descriptor.
//
// descriptor: An [MTLAccelerationStructurePassDescriptor] instance that configures the
// [MTLAccelerationStructureCommandEncoder] the method returns.
//
// # Discussion
// 
// Use an [MTLAccelerationStructureCommandEncoder] instance’s methods to set
// up a single ray-tracing pass.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/makeAccelerationStructureCommandEncoder(descriptor:)

func (o MTLCommandBufferObject) AccelerationStructureCommandEncoderWithDescriptor(descriptor IMTLAccelerationStructurePassDescriptor) MTLAccelerationStructureCommandEncoder {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accelerationStructureCommandEncoderWithDescriptor:"), descriptor)
	return MTLAccelerationStructureCommandEncoderObjectFromID(rv)
	}

// Creates a compute command encoder that uses default settings.
//
// # Discussion
// 
// Use an [MTLComputeCommandEncoder] instance’s methods to set up a single
// compute pass. The encoder this method returns dispatches its compute
// commands serially (see [DispatchTypeSerial]). To create a compute command
// encoder that dispatches commands concurrently (see
// [DispatchTypeConcurrent]), use the [ComputeCommandEncoderWithDispatchType]
// or [ComputeCommandEncoderWithDescriptor] method.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/makeComputeCommandEncoder()

func (o MTLCommandBufferObject) ComputeCommandEncoder() MTLComputeCommandEncoder {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("computeCommandEncoder"))
	return MTLComputeCommandEncoderObjectFromID(rv)
	}

// Creates a compute command encoder from a descriptor.
//
// computePassDescriptor: An [MTLComputePassDescriptor] instance that configures the
// [MTLComputeCommandEncoder] the method returns.
//
// # Discussion
// 
// Use an [MTLComputeCommandEncoder] instance’s methods to set up a single
// compute pass.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/makeComputeCommandEncoder(descriptor:)

func (o MTLCommandBufferObject) ComputeCommandEncoderWithDescriptor(computePassDescriptor IMTLComputePassDescriptor) MTLComputeCommandEncoder {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("computeCommandEncoderWithDescriptor:"), computePassDescriptor)
	return MTLComputeCommandEncoderObjectFromID(rv)
	}

// Creates a parallel render command encoder from a descriptor.
//
// renderPassDescriptor: An [MTLRenderPassDescriptor] instance that configures the
// [MTLParallelRenderCommandEncoder] the method returns.
//
// # Discussion
// 
// An [MTLParallelRenderCommandEncoder] instance can create multiple,
// independent render command encoders that contribute to the same render pass
// on different threads.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/makeParallelRenderCommandEncoder(descriptor:)

func (o MTLCommandBufferObject) ParallelRenderCommandEncoderWithDescriptor(renderPassDescriptor IMTLRenderPassDescriptor) MTLParallelRenderCommandEncoder {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("parallelRenderCommandEncoderWithDescriptor:"), renderPassDescriptor)
	return MTLParallelRenderCommandEncoderObjectFromID(rv)
	}

// Creates a compute command encoder with a dispatch type.
//
// dispatchType: An [MTLDispatchType] instance that indicates whether the compute pass the
// encoder creates runs commands serially or concurrently.
// //
// [MTLDispatchType]: https://developer.apple.com/documentation/Metal/MTLDispatchType
//
// # Discussion
// 
// Use an [MTLComputeCommandEncoder] instance’s methods to set up a single
// compute pass.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/makeComputeCommandEncoder(dispatchType:)

func (o MTLCommandBufferObject) ComputeCommandEncoderWithDispatchType(dispatchType MTLDispatchType) MTLComputeCommandEncoder {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("computeCommandEncoderWithDispatchType:"), dispatchType)
	return MTLComputeCommandEncoderObjectFromID(rv)
	}

// Creates a resource state command encoder that uses default settings.
//
// # Discussion
// 
// Use an [MTLResourceStateCommandEncoder] instance’s methods to create a
// pass that updates the state of one or more sparse textures.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/makeResourceStateCommandEncoder()

func (o MTLCommandBufferObject) ResourceStateCommandEncoder() MTLResourceStateCommandEncoder {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("resourceStateCommandEncoder"))
	return MTLResourceStateCommandEncoderObjectFromID(rv)
	}

// Marks the end of a debug group and, if applicable, restores the previous
// group from a stack.
//
// # Discussion
// 
// Use [PushDebugGroup] to group commands within the command buffer, which
// adds a new group to a stack, effectively nesting a group within any
// previous group. Call [PopDebugGroup] to mark the end of a group of commands
// within the command buffer, and restore the previous group, if applicable.
// You can inspect the group and the commands it contains when viewing the
// contents of a frame capture with Metal Debugger.
// 
// Labels can help you profile and debug your app at runtime with Metal
// Debugger and other tools. See [Naming resources and commands] for more
// information about using labels and other debugging techniques.
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/popDebugGroup()

func (o MTLCommandBufferObject) PopDebugGroup() {
	
	objc.Send[struct{}](o.ID, objc.Sel("popDebugGroup"))
	}

// Creates a render command encoder from a descriptor.
//
// renderPassDescriptor: An [MTLRenderPassDescriptor] instance that configures the
// [MTLRenderCommandEncoder] the method returns.
//
// # Discussion
// 
// Use an [MTLRenderCommandEncoder] instance’s methods to set up a single
// graphics-rendering pass.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/makeRenderCommandEncoder(descriptor:)

func (o MTLCommandBufferObject) RenderCommandEncoderWithDescriptor(renderPassDescriptor IMTLRenderPassDescriptor) MTLRenderCommandEncoder {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("renderCommandEncoderWithDescriptor:"), renderPassDescriptor)
	return MTLRenderCommandEncoderObjectFromID(rv)
	}

// Marks the beginning of a debug group and gives it an identifying label,
// which temporarily replaces the previous group, if applicable.
//
// string: A name for the debug group.
//
// # Discussion
// 
// Use [PushDebugGroup] to group commands within the command buffer, which
// adds a new group to a stack, effectively nesting a group within any
// previous group. Call [PopDebugGroup] to mark the end of a group of commands
// within the command buffer, and restore the previous group, if applicable.
// You can inspect the group and the commands it contains when viewing the
// contents of a frame capture with Metal Debugger.
// 
// Labels can help you profile and debug your app at runtime with Metal
// Debugger and other tools. See [Naming resources and commands] for more
// information about using labels and other debugging techniques.
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/pushDebugGroup(_:)

func (o MTLCommandBufferObject) PushDebugGroup(string_ string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("pushDebugGroup:"), objc.String(string_))
	}

// Creates a resource state command encoder from a descriptor.
//
// resourceStatePassDescriptor: An [MTLResourceStatePassDescriptor] instance that configures the
// [MTLResourceStateCommandEncoder] the method returns.
//
// # Discussion
// 
// Use an [MTLResourceStateCommandEncoder] instance’s methods to create a
// pass that updates the state of one or more sparse textures.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/resourceStateCommandEncoder(with:)

func (o MTLCommandBufferObject) ResourceStateCommandEncoderWithDescriptor(resourceStatePassDescriptor IMTLResourceStatePassDescriptor) MTLResourceStateCommandEncoder {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("resourceStateCommandEncoderWithDescriptor:"), resourceStatePassDescriptor)
	return MTLResourceStateCommandEncoderObjectFromID(rv)
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
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/useResidencySets:count:

func (o MTLCommandBufferObject) UseResidencySetsCount(residencySets []MTLResidencySet, count uint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("useResidencySets:count:"), objc.CArray(residencySets), count)
	}







// The GPU device that indirectly owns the command buffer because you create
// it from a command queue the device also owns.
//
// # Discussion
// 
// The command buffer can only work with other instances that [device]
// creates, directly or indirectly, such as buffers and textures.
//
// [device]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/device
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/device
func (o MTLCommandBufferObject) Device() MTLDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
}




// The command queue that creates the command buffer.
//
// # Discussion
// 
// Each command buffer can only submit its commands to the queue that creates
// it.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/commandQueue
func (o MTLCommandBufferObject) CommandQueue() MTLCommandQueue {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("commandQueue"))
	return MTLCommandQueueObjectFromID(rv)
}




// A description of an error when the GPU encounters an issue as it runs the
// command buffer.
//
// # Discussion
// 
// You typically check this property during development to get more
// information about a runtime issue. The property remains `nil` unless the
// GPU can’t successfully run the command buffer.
// 
// An error’s [userInfo] dictionary property contains additional information
// if the command buffer’s [errorOptions] property includes
// [CommandBufferErrorOptionEncoderExecutionStatus]. You can retrieve an
// [MTLCommandBufferEncoderInfo] instance from the dictionary by accessing it
// with [MTLCommandBufferEncoderInfoErrorKey].
//
// [MTLCommandBufferEncoderInfoErrorKey]: https://developer.apple.com/documentation/Metal/MTLCommandBufferEncoderInfoErrorKey
// [errorOptions]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/errorOptions
// [userInfo]: https://developer.apple.com/documentation/Foundation/NSError/userInfo
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/error
func (o MTLCommandBufferObject) Error() foundation.INSError {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(rv)
}




// The host time, in seconds, when the GPU finishes execution of the command
// buffer.
//
// # Discussion
// 
// You can calculate how much time the GPU spends running a command buffer by
// subtracting [gpuStartTime] from this value. Both values are relative to
// system mach time.
// 
// The GPU start and end times remain `0.0` until the GPU finishes running the
// command buffer. Check this value after the [WaitUntilCompleted] method
// returns, or within a completion handler passed to the [AddCompletedHandler]
// method.
//
// [gpuStartTime]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/gpuStartTime
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/gpuEndTime
func (o MTLCommandBufferObject) GPUEndTime() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("GPUEndTime"))
	return float64(rv)
}




// Settings that determine which information the command buffer records about
// execution errors, and how it does it.
//
// # Discussion
// 
// The property reflects the [ErrorOptions] property of the
// [MTLCommandBufferDescriptor] instance at the time you create the command
// buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/errorOptions
func (o MTLCommandBufferObject) ErrorOptions() MTLCommandBufferErrorOption {
	rv := objc.Send[MTLCommandBufferErrorOption](o.ID, objc.Sel("errorOptions"))
	return MTLCommandBufferErrorOption(rv)
}




// The messages the command buffer records as the GPU runs its commands.
//
// # Discussion
// 
// The value of this property is valid only after the command buffer finishes
// executing.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/logs
func (o MTLCommandBufferObject) Logs() MTLLogContainer {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("logs"))
	return MTLLogContainerObjectFromID(rv)
}




// An optional name that can help you identify the command buffer.
//
// # Discussion
// 
// Set labels to help you quickly identify a command buffer at runtime in the
// Metal debugging and profiling tools. See [Naming resources and commands].
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/label
func (o MTLCommandBufferObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

func (o MTLCommandBufferObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}



// The host time, in seconds, when the CPU finishes scheduling the command
// buffer.
//
// # Discussion
// 
// You can calculate how much time the kernel spends scheduling a command
// buffer by subtracting [kernelStartTime] from this value.
// 
// The kernel start and end times remain `0.0` until the GPU driver (on the
// CPU) schedules the command buffer to run on the GPU. Apps typically use
// these values after the [WaitUntilScheduled] method returns, or within a
// completion handler (see [AddScheduledHandler] and [AddCompletedHandler]).
//
// [kernelStartTime]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/kernelStartTime
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/kernelEndTime
func (o MTLCommandBufferObject) KernelEndTime() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("kernelEndTime"))
	return float64(rv)
}




// The host time, in seconds, when the CPU begins to schedule the command
// buffer.
//
// # Discussion
// 
// You can calculate how much time the kernel spends scheduling a command
// buffer by subtracting this value from [kernelEndTime].
// 
// The kernel start and end times remain `0.0` until the GPU driver (on the
// CPU) schedules the command buffer to run on the GPU. Apps typically use
// these values after the [WaitUntilScheduled] method returns, or within a
// completion handler (see [AddScheduledHandler] and [AddCompletedHandler]).
//
// [kernelEndTime]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/kernelEndTime
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/kernelStartTime
func (o MTLCommandBufferObject) KernelStartTime() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("kernelStartTime"))
	return float64(rv)
}




// The host time, in seconds, when the GPU starts command buffer execution.
//
// # Discussion
// 
// You can calculate how much time the GPU spends running a command buffer by
// subtracting this value from [gpuEndTime]. Both values are relative to
// system mach time.
// 
// The GPU start and end times remain `0.0` until the GPU finishes running the
// command buffer. Check this value after the [WaitUntilCompleted] method
// returns, or within a completion handler passed to the [AddCompletedHandler]
// method.
//
// [gpuEndTime]: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/gpuEndTime
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/gpuStartTime
func (o MTLCommandBufferObject) GPUStartTime() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("GPUStartTime"))
	return float64(rv)
}




// A Boolean value that indicates whether the command buffer maintains strong
// references to the resources it uses.
//
// # Discussion
// 
// You can configure this property when you create a command buffer by setting
// [RetainedReferences] of an [MTLCommandBufferDescriptor] instance and
// calling the [CommandBufferWithDescriptor] method. The [CommandBuffer]
// method sets this property to [true], and
// [CommandBufferWithUnretainedReferences] sets it to [false].
// 
// If [false], your app is responsible for maintaining strong references to
// all the resources the command buffer relies on until it completes.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandBuffer/retainedReferences
func (o MTLCommandBufferObject) RetainedReferences() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("retainedReferences"))
	return bool(rv)
}






