// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An instance that splits up a single render pass so that it can be simultaneously encoded from multiple threads.
//
// See: https://developer.apple.com/documentation/Metal/MTLParallelRenderCommandEncoder
type MTLParallelRenderCommandEncoder interface {
	objectivec.IObject
	MTLCommandEncoder

	// Create an object that encodes commands that perform graphics rendering operations and may be assigned to a different thread.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLParallelRenderCommandEncoder/makeRenderCommandEncoder()
	RenderCommandEncoder() MTLRenderCommandEncoder

	// Specifies a known store action to replace the initial [MTLStoreAction.unknown](<doc://com.apple.metal/documentation/Metal/MTLStoreAction/unknown>) value specified for a given color attachment.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLParallelRenderCommandEncoder/setColorStoreAction(_:index:)
	SetColorStoreActionAtIndex(storeAction MTLStoreAction, colorAttachmentIndex uint)

	// Specifies known store action options for a given color attachment.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLParallelRenderCommandEncoder/setColorStoreActionOptions(_:index:)
	SetColorStoreActionOptionsAtIndex(storeActionOptions MTLStoreActionOptions, colorAttachmentIndex uint)

	// Specifies a known store action to replace the initial [MTLStoreAction.unknown](<doc://com.apple.metal/documentation/Metal/MTLStoreAction/unknown>) value specified for a given depth attachment.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLParallelRenderCommandEncoder/setDepthStoreAction(_:)
	SetDepthStoreAction(storeAction MTLStoreAction)

	// Specifies known store action options for a given depth attachment.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLParallelRenderCommandEncoder/setDepthStoreActionOptions(_:)
	SetDepthStoreActionOptions(storeActionOptions MTLStoreActionOptions)

	// Specifies a known store action to replace the initial [MTLStoreAction.unknown](<doc://com.apple.metal/documentation/Metal/MTLStoreAction/unknown>) value specified for a given stencil attachment.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLParallelRenderCommandEncoder/setStencilStoreAction(_:)
	SetStencilStoreAction(storeAction MTLStoreAction)

	// Specifies known store action options for a given stencil attachment.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLParallelRenderCommandEncoder/setStencilStoreActionOptions(_:)
	SetStencilStoreActionOptions(storeActionOptions MTLStoreActionOptions)
}

// MTLParallelRenderCommandEncoderObject wraps an existing Objective-C object that conforms to the MTLParallelRenderCommandEncoder protocol.
type MTLParallelRenderCommandEncoderObject struct {
	objectivec.Object
}

func (o MTLParallelRenderCommandEncoderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLParallelRenderCommandEncoderObjectFromID constructs a [MTLParallelRenderCommandEncoderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLParallelRenderCommandEncoderObjectFromID(id objc.ID) MTLParallelRenderCommandEncoderObject {
	return MTLParallelRenderCommandEncoderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Create an object that encodes commands that perform graphics rendering
// operations and may be assigned to a different thread.
//
// # Return Value
//
// # A graphics rendering command encoder object
//
// # Discussion
//
// The rendering commands encoded by [MTLRenderCommandEncoder] objects are
// executed in the order in which the [MTLRenderCommandEncoder] objects are
// created, not in the order they are ended.
//
// See: https://developer.apple.com/documentation/Metal/MTLParallelRenderCommandEncoder/makeRenderCommandEncoder()
func (o MTLParallelRenderCommandEncoderObject) RenderCommandEncoder() MTLRenderCommandEncoder {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("renderCommandEncoder"))
	return MTLRenderCommandEncoderObjectFromID(rv)
}

// Specifies a known store action to replace the initial
// [MTLStoreActionUnknown] value specified for a given color attachment.
//
// storeAction: The desired store action for the color attachment. This value can’t be
// [MTLStoreActionUnknown].
//
// colorAttachmentIndex: The index of the color attachment.
//
// # Discussion
//
// If the store action for the given color attachment was set to
// [MTLStoreActionUnknown] when the parallel render command encoder was
// created, you need to call this method to specify another store action
// before you call the [EndEncoding] method.
//
// See: https://developer.apple.com/documentation/Metal/MTLParallelRenderCommandEncoder/setColorStoreAction(_:index:)
func (o MTLParallelRenderCommandEncoderObject) SetColorStoreActionAtIndex(storeAction MTLStoreAction, colorAttachmentIndex uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setColorStoreAction:atIndex:"), storeAction, colorAttachmentIndex)
}

// Specifies known store action options for a given color attachment.
//
// storeActionOptions: The additional store action options for the color attachment.
//
// colorAttachmentIndex: The index of the color attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLParallelRenderCommandEncoder/setColorStoreActionOptions(_:index:)
func (o MTLParallelRenderCommandEncoderObject) SetColorStoreActionOptionsAtIndex(storeActionOptions MTLStoreActionOptions, colorAttachmentIndex uint) {
	objc.Send[struct{}](o.ID, objc.Sel("setColorStoreActionOptions:atIndex:"), storeActionOptions, colorAttachmentIndex)
}

// Specifies a known store action to replace the initial
// [MTLStoreActionUnknown] value specified for a given depth attachment.
//
// storeAction: The desired store action for the depth attachment. This value can’t be
// [MTLStoreActionUnknown].
//
// # Discussion
//
// If the store action for the given depth attachment was set to
// [MTLStoreActionUnknown] when the parallel render command encoder was
// created, you need to call this method to specify another store action
// before you call the [EndEncoding] method.
//
// See: https://developer.apple.com/documentation/Metal/MTLParallelRenderCommandEncoder/setDepthStoreAction(_:)
func (o MTLParallelRenderCommandEncoderObject) SetDepthStoreAction(storeAction MTLStoreAction) {
	objc.Send[struct{}](o.ID, objc.Sel("setDepthStoreAction:"), storeAction)
}

// Specifies known store action options for a given depth attachment.
//
// storeActionOptions: The additional store action options for the depth attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLParallelRenderCommandEncoder/setDepthStoreActionOptions(_:)
func (o MTLParallelRenderCommandEncoderObject) SetDepthStoreActionOptions(storeActionOptions MTLStoreActionOptions) {
	objc.Send[struct{}](o.ID, objc.Sel("setDepthStoreActionOptions:"), storeActionOptions)
}

// Specifies a known store action to replace the initial
// [MTLStoreActionUnknown] value specified for a given stencil attachment.
//
// storeAction: The desired store action for the stencil attachment. This value can’t be
// [MTLStoreActionUnknown].
//
// # Discussion
//
// If the store action for the given stencil attachment was set to
// [MTLStoreActionUnknown] when the parallel render command encoder was
// created, you need to call this method to specify another store action
// before you call the [EndEncoding] method.
//
// See: https://developer.apple.com/documentation/Metal/MTLParallelRenderCommandEncoder/setStencilStoreAction(_:)
func (o MTLParallelRenderCommandEncoderObject) SetStencilStoreAction(storeAction MTLStoreAction) {
	objc.Send[struct{}](o.ID, objc.Sel("setStencilStoreAction:"), storeAction)
}

// Specifies known store action options for a given stencil attachment.
//
// storeActionOptions: The additional store action options for the stencil attachment.
//
// See: https://developer.apple.com/documentation/Metal/MTLParallelRenderCommandEncoder/setStencilStoreActionOptions(_:)
func (o MTLParallelRenderCommandEncoderObject) SetStencilStoreActionOptions(storeActionOptions MTLStoreActionOptions) {
	objc.Send[struct{}](o.ID, objc.Sel("setStencilStoreActionOptions:"), storeActionOptions)
}

// Declares that all command generation from the encoder is completed.
//
// # Discussion
//
// After `endEncoding` is called, the command encoder has no further use. You
// cannot encode any other commands with this encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/endEncoding()
func (o MTLParallelRenderCommandEncoderObject) EndEncoding() {
	objc.Send[struct{}](o.ID, objc.Sel("endEncoding"))
}

// Inserts a debug string into the captured frame data.
//
// # Discussion
//
// For more information, see [Naming resources and commands].
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/insertDebugSignpost(_:)
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
func (o MTLParallelRenderCommandEncoderObject) InsertDebugSignpost(string_ string) {
	objc.Send[struct{}](o.ID, objc.Sel("insertDebugSignpost:"), objc.String(string_))
}

// Pushes a specific string onto a stack of debug group strings for the
// command encoder.
//
// # Discussion
//
// For more information, see [Naming resources and commands].
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/pushDebugGroup(_:)
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
func (o MTLParallelRenderCommandEncoderObject) PushDebugGroup(string_ string) {
	objc.Send[struct{}](o.ID, objc.Sel("pushDebugGroup:"), objc.String(string_))
}

// Pops the latest string off of a stack of debug group strings for the
// command encoder.
//
// # Discussion
//
// For more information, see [Naming resources and commands].
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/popDebugGroup()
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
func (o MTLParallelRenderCommandEncoderObject) PopDebugGroup() {
	objc.Send[struct{}](o.ID, objc.Sel("popDebugGroup"))
}

// The Metal device from which the command encoder was created.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/device
func (o MTLParallelRenderCommandEncoderObject) Device() MTLDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
}

// A string that labels the command encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/label
func (o MTLParallelRenderCommandEncoderObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

// Encodes a consumer barrier on work you commit to the same command queue.
//
// afterQueueStages: [MTLStages] mask that represents the stages of work to wait for. This
// argument applies to work corresponding to these stages you encode in prior
// command encoders, and not for the current encoder.
//
// beforeStages: [MTLStages] mask that represents the stages of work that wait. This
// argument applies to work you encode in the current command encoder.
//
// # Discussion
//
// Encode a barrier that guarantees that any subsequent work you encode in the
// current command encoder that corresponds to the `beforeStages` stages
// doesn’t proceed until Metal completes all work prior to the current
// command encoder corresponding to the `afterQueueStages` stages, completes.
//
// Metal can reorder the exact point where it applies the barrier, so use this
// method for synchronizing between different passes.
//
// If you need to synchronize work within a pass that you encode with an
// instance of a subclass of [MTLCommandEncoder], use memory barriers instead.
// For subclasses of [MTL4CommandEncoder], use encoder barriers.
//
// You can specify `afterQueueStages` and `beforeStages` that contain
// [MTLStages] unrelated to the current command encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/barrier(afterQueueStages:beforeStages:)
//
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
//
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
// [MTLStages]: https://developer.apple.com/documentation/Metal/MTLStages
func (o MTLParallelRenderCommandEncoderObject) BarrierAfterQueueStagesBeforeStages(afterQueueStages MTLStages, beforeStages MTLStages) {
	objc.Send[struct{}](o.ID, objc.Sel("barrierAfterQueueStages:beforeStages:"), afterQueueStages, beforeStages)
}

// A string that labels the command encoder.
//
// # Discussion
//
// Object and command labels are useful identifiers at runtime or when
// profiling and debugging your app using any Metal tool. See [Naming
// resources and commands].
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/label
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
func (o MTLParallelRenderCommandEncoderObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}
