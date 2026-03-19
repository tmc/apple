// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// Encodes machine-learning model inference commands for a single pass.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningCommandEncoder
type MTL4MachineLearningCommandEncoder interface {
	objectivec.IObject
	MTL4CommandEncoder

	// Configures the encoder with a machine learning pipeline state instance.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningCommandEncoder/setPipelineState(_:)
	SetPipelineState(pipelineState MTL4MachineLearningPipelineState)

	// Sets an argument table for the command encoder’s machine learning shader stage.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningCommandEncoder/setArgumentTable(_:)
	SetArgumentTable(argumentTable MTL4ArgumentTable)

	// Dispatches a machine learning network using the current pipeline state and argument table.
	//
	// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningCommandEncoder/dispatchNetwork(intermediatesHeap:)
	DispatchNetworkWithIntermediatesHeap(heap MTLHeap)
}

// MTL4MachineLearningCommandEncoderObject wraps an existing Objective-C object that conforms to the MTL4MachineLearningCommandEncoder protocol.
type MTL4MachineLearningCommandEncoderObject struct {
	objectivec.Object
}
func (o MTL4MachineLearningCommandEncoderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTL4MachineLearningCommandEncoderObjectFromID constructs a [MTL4MachineLearningCommandEncoderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTL4MachineLearningCommandEncoderObjectFromID(id objc.ID) MTL4MachineLearningCommandEncoderObject {
	return MTL4MachineLearningCommandEncoderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Configures the encoder with a machine learning pipeline state instance.
//
// pipelineState: A Machine Learning pipeline state instance.
//
// # Discussion
// 
// The pipeline state instance affects all subsequent Machine Learning
// commands.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningCommandEncoder/setPipelineState(_:)
func (o MTL4MachineLearningCommandEncoderObject) SetPipelineState(pipelineState MTL4MachineLearningPipelineState) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setPipelineState:"), pipelineState)
	}
// Sets an argument table for the command encoder’s machine learning shader
// stage.
//
// argumentTable: An argument table to set on the command encoder’s Machine Learning stage.
//
// # Discussion
// 
// The argument table provides inputs to all subsequent Machine Learning
// dispatches.
//
// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningCommandEncoder/setArgumentTable(_:)
func (o MTL4MachineLearningCommandEncoderObject) SetArgumentTable(argumentTable MTL4ArgumentTable) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setArgumentTable:"), argumentTable)
	}
// Dispatches a machine learning network using the current pipeline state and
// argument table.
//
// heap: A heap that Metal can use to allocate intermediate tensors.
//
// # Discussion
// 
// This method takes a parameter consisting of a [MTLHeap] that Metal can use
// to allocate intermediate tensors. You can query the minimum size Metal
// requires for this heap by calling [IntermediatesHeapSize].
//
// See: https://developer.apple.com/documentation/Metal/MTL4MachineLearningCommandEncoder/dispatchNetwork(intermediatesHeap:)
func (o MTL4MachineLearningCommandEncoderObject) DispatchNetworkWithIntermediatesHeap(heap MTLHeap) {
	
	objc.Send[struct{}](o.ID, objc.Sel("dispatchNetworkWithIntermediatesHeap:"), heap)
	}
// Returns the command buffer that is currently encoding commands.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/commandBuffer
func (o MTL4MachineLearningCommandEncoderObject) CommandBuffer() MTL4CommandBuffer {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("commandBuffer"))
	return MTL4CommandBufferObjectFromID(rv)
	}
// Provides an optional label to assign to the command encoder for debug
// purposes.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/label
func (o MTL4MachineLearningCommandEncoderObject) Label() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
	}
// Declares that all command generation from this encoder is complete.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/endEncoding()
func (o MTL4MachineLearningCommandEncoderObject) EndEncoding() {
	
	objc.Send[struct{}](o.ID, objc.Sel("endEncoding"))
	}
// Inserts a debug string into the frame data to aid debugging.
//
// string: The debug string to insert as a signpost.
//
// # Discussion
// 
// Calling this method doesn’t change any behaviors, but can be useful for
// debugging purposes.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/insertDebugSignpost(_:)
func (o MTL4MachineLearningCommandEncoderObject) InsertDebugSignpost(string_ string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("insertDebugSignpost:"), objc.String(string_))
	}
// Pops the latest debug group string from this encoder’s stack of debug
// groups.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/popDebugGroup()
func (o MTL4MachineLearningCommandEncoderObject) PopDebugGroup() {
	
	objc.Send[struct{}](o.ID, objc.Sel("popDebugGroup"))
	}
// Pushes a string onto this encoder’s stack of debug groups.
//
// string: The debug string to push.
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/pushDebugGroup(_:)
func (o MTL4MachineLearningCommandEncoderObject) PushDebugGroup(string_ string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("pushDebugGroup:"), objc.String(string_))
	}
// Encodes a command that instructs the GPU to update a fence after one or
// more stages, which can unblock other passes waiting for the fence.
//
// fence: A fence the pass updates after the stages in `afterEncoderStages` complete.
//
// afterEncoderStages: The encoder stages that need to complete before the pass updates `fence`.
//
// # Discussion
// 
// You can synchronize memory operations of a pass that access resources with
// an [MTLFence]. This method instructs the pass to update `fence` after the
// stages you pass to the `afterEncoderStages` run all their memory store
// operations to the resources it accesses. The fence indicates when other
// passes can access those resources without a race condition.
// 
// For more information about synchronization with fences, see:
// 
// - [Resource synchronization]
// - [Synchronizing passes with a fence]
// 
// # Reuse a fence by waiting first and updating second
// 
// When encoding a pass that reuses a fence, wait for other passes to update
// the fence before repurposing that fence to notify subsequent passes with an
// update:
// 
// - Call the [WaitForFenceBeforeEncoderStages] method before encoding
// commands that need to wait for other passes. - Call the
// [UpdateFenceAfterEncoderStages] method after encoding commands that later
// passes depend on.
// 
// The GPU driver evaluates the fences that apply to the pass and the commands
// that depend on those fences when your app commits the enclosing
// [MTLCommandBuffer].
// 
// To synchronize different stages within a single pass, create an because a
// fence can only synchronize memory operations between different passes. For
// more information, see [Synchronizing stages within a pass].
//
// [Resource synchronization]: https://developer.apple.com/documentation/Metal/resource-synchronization
// [Synchronizing passes with a fence]: https://developer.apple.com/documentation/Metal/synchronizing-passes-with-a-fence
// [Synchronizing stages within a pass]: https://developer.apple.com/documentation/Metal/synchronizing-stages-within-a-pass
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/updateFence(_:afterEncoderStages:)
func (o MTL4MachineLearningCommandEncoderObject) UpdateFenceAfterEncoderStages(fence MTLFence, afterEncoderStages MTLStages) {
	
	objc.Send[struct{}](o.ID, objc.Sel("updateFence:afterEncoderStages:"), fence, afterEncoderStages)
	}
// Encodes a command that instructs the GPU to pause before starting one or
// more stages of the pass until a pass updates a fence.
//
// fence: A fence that the pass waits for before running the stages you pass to
// `beforeEncoderStages`.
//
// beforeEncoderStages: The encoder stages that need to wait for another pass to update `fence`
// before they run.
//
// # Discussion
// 
// You can synchronize memory operations of a pass that access resources with
// an [MTLFence]. This method instructs the GPU to wait until another pass
// updates `fence` before running the stages you pass to the
// `beforeEncoderStages` parameter. The fence indicates when the pass can
// access those resources without a race condition.
// 
// For more information about synchronization with fences, see:
// 
// - [Resource synchronization]
// - [Synchronizing passes with a fence]
// 
// # Reuse a fence by waiting first and updating second
// 
// When encoding a pass that reuses a fence, wait for other passes to update
// the fence before repurposing that fence to notify subsequent passes with an
// update:
// 
// - Call the [WaitForFenceBeforeEncoderStages] method before encoding
// commands that need to wait for other passes. - Call the
// [UpdateFenceAfterEncoderStages] method after encoding commands that later
// passes depend on.
// 
// The GPU driver evaluates the fences that apply to the pass and the commands
// that depend on those fences when your app commits the enclosing
// [MTLCommandBuffer].
// 
// To synchronize different stages within a single pass, create an because a
// fence can only synchronize memory operations between different passes. For
// more information, see [Synchronizing stages within a pass].
//
// [Resource synchronization]: https://developer.apple.com/documentation/Metal/resource-synchronization
// [Synchronizing passes with a fence]: https://developer.apple.com/documentation/Metal/synchronizing-passes-with-a-fence
// [Synchronizing stages within a pass]: https://developer.apple.com/documentation/Metal/synchronizing-stages-within-a-pass
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/waitForFence(_:beforeEncoderStages:)
func (o MTL4MachineLearningCommandEncoderObject) WaitForFenceBeforeEncoderStages(fence MTLFence, beforeEncoderStages MTLStages) {
	
	objc.Send[struct{}](o.ID, objc.Sel("waitForFence:beforeEncoderStages:"), fence, beforeEncoderStages)
	}
// Encodes an intra-pass barrier.
//
// afterEncoderStages: [MTLStages] mask that represents the stages of work to wait for. This
// argument only applies to subsequent work you encode in the current command
// encoder.
// //
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
//
// beforeEncoderStages: [MTLStages] mask that represents the stages of work that wait. This
// argument only applies to work you encode in the current command encoder
// prior to this barrier.
// //
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
//
// visibilityOptions: [MTL4VisibilityOptions] of the barrier, controlling cache flush behavior.
// //
// [MTL4VisibilityOptions]: https://developer.apple.com/documentation/Metal/MTL4VisibilityOptions
//
// # Discussion
// 
// Encode a barrier that guarantees that any subsequent work you encode in the
// , corresponding to `beforeEncoderStages`, doesn’t begin until all prior
// commands in this command encoder, corresponding to `afterEncoderStages`,
// completes.
// 
// When calling this method, it’s your responsibility to ensure parameters
// `afterEncoderStages` and `beforeEncoderStages` contain a combination of
// [MTLStages] for which this encoder can encode commands. For example, for a
// [MTL4ComputeCommandEncoder] instance, you can provide any combination of
// [StageDispatch], [StageBlit] and [StageAccelerationStructure].
//
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/barrierAfterEncoderStages:beforeEncoderStages:visibilityOptions:
func (o MTL4MachineLearningCommandEncoderObject) BarrierAfterEncoderStagesBeforeEncoderStagesVisibilityOptions(afterEncoderStages MTLStages, beforeEncoderStages MTLStages, visibilityOptions MTL4VisibilityOptions) {
	
	objc.Send[struct{}](o.ID, objc.Sel("barrierAfterEncoderStages:beforeEncoderStages:visibilityOptions:"), afterEncoderStages, beforeEncoderStages, visibilityOptions)
	}
// Encodes a consumer barrier on work you commit to the same command queue.
//
// afterQueueStages: [MTLStages] mask that represents the stages of work to wait for. This
// argument applies to work corresponding to these stages you encode in prior
// command encoders, and not for the current encoder.
// //
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
//
// beforeStages: [MTLStages] mask that represents the stages of work that wait. This
// argument applies to work you encode in the current command encoder.
// //
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
//
// visibilityOptions: [MTL4VisibilityOptions] of the barrier.
// //
// [MTL4VisibilityOptions]: https://developer.apple.com/documentation/Metal/MTL4VisibilityOptions
//
// # Discussion
// 
// Encode a barrier that guarantees that any subsequent work you encode in the
// current command encoder that corresponds to the `beforeStages` stages
// doesn’t proceed until Metal completes all work prior to the current
// command encoder corresponding to the `afterQueueStages` stages, completes.
// 
// Metal can reorder the exact point where it applies the barrier, so encode
// the barrier as close to the command that consumes the resource as possible.
// Don’t use this method for synchronizing resource access within the same
// pass.
// 
// If you need to synchronize work within a pass that you encode with an
// instance of a subclass of [MTLCommandEncoder], use memory barriers instead.
// For subclasses of [MTL4CommandEncoder], use encoder barriers.
// 
// You can specify `afterQueueStages` and `beforeStages` that contain
// [MTLStages] unrelated to the current command encoder.
//
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/barrierAfterQueueStages:beforeStages:visibilityOptions:
func (o MTL4MachineLearningCommandEncoderObject) BarrierAfterQueueStagesBeforeStagesVisibilityOptions(afterQueueStages MTLStages, beforeStages MTLStages, visibilityOptions MTL4VisibilityOptions) {
	
	objc.Send[struct{}](o.ID, objc.Sel("barrierAfterQueueStages:beforeStages:visibilityOptions:"), afterQueueStages, beforeStages, visibilityOptions)
	}
// Encodes a producer barrier on work committed to the same command queue.
//
// afterStages: [MTLStages] mask that represents the stages of work to wait for. This
// argument applies to work corresponding to these stages you encode in the
// current command encoder prior to this barrier command.
// //
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
//
// beforeQueueStages: [MTLStages] mask that represents the stages of work that need to wait. This
// argument applies to subsequent encoders and not to work in the current
// command encoder.
// //
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
//
// visibilityOptions: [MTL4VisibilityOptions] of the barrier, controlling cache flush behavior.
// //
// [MTL4VisibilityOptions]: https://developer.apple.com/documentation/Metal/MTL4VisibilityOptions
//
// # Discussion
// 
// This method encodes a barrier that guarantees that any work you encode
// using , corresponding to `beforeQueueStages`, don’t begin until all
// commands you previously encode in the current encoder (and prior encoders),
// corresponding to `afterStages`, complete.
// 
// When calling this method, you can pass any [MTLStages] to parameters
// `afterStages` and `beforeQueueStages`, even stages that don’t relate to
// the current or prior command encoders.
//
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
//
// See: https://developer.apple.com/documentation/Metal/MTL4CommandEncoder/barrierAfterStages:beforeQueueStages:visibilityOptions:
func (o MTL4MachineLearningCommandEncoderObject) BarrierAfterStagesBeforeQueueStagesVisibilityOptions(afterStages MTLStages, beforeQueueStages MTLStages, visibilityOptions MTL4VisibilityOptions) {
	
	objc.Send[struct{}](o.ID, objc.Sel("barrierAfterStages:beforeQueueStages:visibilityOptions:"), afterStages, beforeQueueStages, visibilityOptions)
	}

func (o MTL4MachineLearningCommandEncoderObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}

