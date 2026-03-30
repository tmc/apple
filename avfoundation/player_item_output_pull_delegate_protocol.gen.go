// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// Methods you can implement to respond to pixel buffer changes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutputPullDelegate
type AVPlayerItemOutputPullDelegate interface {
	objectivec.IObject
}

// AVPlayerItemOutputPullDelegateObject wraps an existing Objective-C object that conforms to the AVPlayerItemOutputPullDelegate protocol.
type AVPlayerItemOutputPullDelegateObject struct {
	objectivec.Object
}

func (o AVPlayerItemOutputPullDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVPlayerItemOutputPullDelegateObjectFromID constructs a [AVPlayerItemOutputPullDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVPlayerItemOutputPullDelegateObjectFromID(id objc.ID) AVPlayerItemOutputPullDelegateObject {
	return AVPlayerItemOutputPullDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that new samples are about to arrive.
//
// sender: The output object that sent the message.
//
// # Discussion
//
// You can use this method to prepare for any new sample data. This method is
// called at some point after a call to your video output object’s “
// method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutputPullDelegate/outputMediaDataWillChange(_:)
func (o AVPlayerItemOutputPullDelegateObject) OutputMediaDataWillChange(sender IAVPlayerItemOutput) {
	objc.Send[struct{}](o.ID, objc.Sel("outputMediaDataWillChange:"), sender)
}

// Tells the delegate that a new sample sequence is commencing.
//
// output: The output object that sent the message.
//
// # Discussion
//
// This method is called after any attempt to seek or change the playback
// direction of the item’s content. If you are maintaining any queued future
// samples, you can use your implementation of this method to discard those
// samples.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutputPullDelegate/outputSequenceWasFlushed(_:)
func (o AVPlayerItemOutputPullDelegateObject) OutputSequenceWasFlushed(output IAVPlayerItemOutput) {
	objc.Send[struct{}](o.ID, objc.Sel("outputSequenceWasFlushed:"), output)
}

// AVPlayerItemOutputPullDelegateConfig holds optional typed callbacks for [AVPlayerItemOutputPullDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutputpulldelegate
type AVPlayerItemOutputPullDelegateConfig struct {

	// Responding to pixel buffer changes
	// OutputMediaDataWillChange — Tells the delegate that new samples are about to arrive.
	OutputMediaDataWillChange func(sender AVPlayerItemOutput)
	// OutputSequenceWasFlushed — Tells the delegate that a new sample sequence is commencing.
	OutputSequenceWasFlushed func(output AVPlayerItemOutput)
}

// NewAVPlayerItemOutputPullDelegate creates an Objective-C object implementing the [AVPlayerItemOutputPullDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AVPlayerItemOutputPullDelegateObject] satisfies the [AVPlayerItemOutputPullDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutputpulldelegate
func NewAVPlayerItemOutputPullDelegate(config AVPlayerItemOutputPullDelegateConfig) AVPlayerItemOutputPullDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAVPlayerItemOutputPullDelegate_%d", n)

	var methods []objc.MethodDef

	if config.OutputMediaDataWillChange != nil {
		fn := config.OutputMediaDataWillChange
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outputMediaDataWillChange:"),
			Fn: func(self objc.ID, _cmd objc.SEL, senderID objc.ID) {
				sender := AVPlayerItemOutputFromID(senderID)
				fn(sender)
			},
		})
	}

	if config.OutputSequenceWasFlushed != nil {
		fn := config.OutputSequenceWasFlushed
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outputSequenceWasFlushed:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outputID objc.ID) {
				output := AVPlayerItemOutputFromID(outputID)
				fn(output)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("AVPlayerItemOutputPullDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAVPlayerItemOutputPullDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AVPlayerItemOutputPullDelegateObjectFromID(instance)
}
