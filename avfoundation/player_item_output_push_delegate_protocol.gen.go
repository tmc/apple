// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// A protocol that defines the methods to implement to respond to changes in the media data sequence.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutputPushDelegate
type AVPlayerItemOutputPushDelegate interface {
	objectivec.IObject
}

// AVPlayerItemOutputPushDelegateObject wraps an existing Objective-C object that conforms to the AVPlayerItemOutputPushDelegate protocol.
type AVPlayerItemOutputPushDelegateObject struct {
	objectivec.Object
}

func (o AVPlayerItemOutputPushDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVPlayerItemOutputPushDelegateObjectFromID constructs a [AVPlayerItemOutputPushDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVPlayerItemOutputPushDelegateObjectFromID(id objc.ID) AVPlayerItemOutputPushDelegateObject {
	return AVPlayerItemOutputPushDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the delegate that the output is starting a new sequence of media
// data.
//
// output: The [AVPlayerItemOutput] object.
//
// # Discussion
//
// This method is invoked after any seeking and change in playback direction.
// If you are maintaining any queued future media data, you may want to
// discard those objects after receiving this message.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemOutputPushDelegate/outputSequenceWasFlushed(_:)
func (o AVPlayerItemOutputPushDelegateObject) OutputSequenceWasFlushed(output IAVPlayerItemOutput) {
	objc.Send[struct{}](o.ID, objc.Sel("outputSequenceWasFlushed:"), output)
}

// AVPlayerItemOutputPushDelegateConfig holds optional typed callbacks for [AVPlayerItemOutputPushDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutputpushdelegate
type AVPlayerItemOutputPushDelegateConfig struct {

	// Flushing sequence state
	// OutputSequenceWasFlushed — Tells the delegate that the output is starting a new sequence of media data.
	OutputSequenceWasFlushed func(output AVPlayerItemOutput)
}

// NewAVPlayerItemOutputPushDelegate creates an Objective-C object implementing the [AVPlayerItemOutputPushDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AVPlayerItemOutputPushDelegateObject] satisfies the [AVPlayerItemOutputPushDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avplayeritemoutputpushdelegate
func NewAVPlayerItemOutputPushDelegate(config AVPlayerItemOutputPushDelegateConfig) AVPlayerItemOutputPushDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAVPlayerItemOutputPushDelegate_%d", n)

	var methods []objc.MethodDef

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
	proto := objc.GetProtocol("AVPlayerItemOutputPushDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAVPlayerItemOutputPushDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AVPlayerItemOutputPushDelegateObjectFromID(instance)
}
