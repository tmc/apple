// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An encoder that writes GPU commands into a command buffer.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder
type MTLCommandEncoder interface {
	objectivec.IObject

	// Declares that all command generation from the encoder is completed.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/endEncoding()
	EndEncoding()

	// Inserts a debug string into the captured frame data.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/insertDebugSignpost(_:)
	InsertDebugSignpost(string_ string)

	// Pushes a specific string onto a stack of debug group strings for the command encoder.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/pushDebugGroup(_:)
	PushDebugGroup(string_ string)

	// Pops the latest string off of a stack of debug group strings for the command encoder.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/popDebugGroup()
	PopDebugGroup()

	// The Metal device from which the command encoder was created.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/device
	Device() MTLDevice

	// A string that labels the command encoder.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/label
	Label() string

	// Encodes a consumer barrier on work you commit to the same command queue.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/barrier(afterQueueStages:beforeStages:)
	BarrierAfterQueueStagesBeforeStages(afterQueueStages MTLStages, beforeStages MTLStages)

	// A string that labels the command encoder.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/label
	SetLabel(value string)
}

// MTLCommandEncoderObject wraps an existing Objective-C object that conforms to the MTLCommandEncoder protocol.
type MTLCommandEncoderObject struct {
	objectivec.Object
}

func (o MTLCommandEncoderObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLCommandEncoderObjectFromID constructs a [MTLCommandEncoderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLCommandEncoderObjectFromID(id objc.ID) MTLCommandEncoderObject {
	return MTLCommandEncoderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Declares that all command generation from the encoder is completed.
//
// # Discussion
//
// After `endEncoding` is called, the command encoder has no further use. You
// cannot encode any other commands with this encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/endEncoding()
func (o MTLCommandEncoderObject) EndEncoding() {
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
func (o MTLCommandEncoderObject) InsertDebugSignpost(string_ string) {
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
func (o MTLCommandEncoderObject) PushDebugGroup(string_ string) {
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
func (o MTLCommandEncoderObject) PopDebugGroup() {
	objc.Send[struct{}](o.ID, objc.Sel("popDebugGroup"))
}

// The Metal device from which the command encoder was created.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/device
func (o MTLCommandEncoderObject) Device() MTLDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
}

// A string that labels the command encoder.
//
// See: https://developer.apple.com/documentation/Metal/MTLCommandEncoder/label
func (o MTLCommandEncoderObject) Label() string {
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
func (o MTLCommandEncoderObject) BarrierAfterQueueStagesBeforeStages(afterQueueStages MTLStages, beforeStages MTLStages) {
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
func (o MTLCommandEncoderObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}
