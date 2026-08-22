// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// Defines common delegate methods for objects participating in sample buffer output.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSampleBufferOutputDelegate
type AVPlayerItemSampleBufferOutputDelegate interface {
	objectivec.IObject

	// Invoked when the output becomes ready to deliver a sample buffer.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSampleBufferOutputDelegate/outputMediaDataAvailable:
	OutputMediaDataAvailable(output IAVPlayerItemSampleBufferOutput)

	// Invoked when the output is commencing a new sequence.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSampleBufferOutputDelegate/outputSequenceWasRestarted:
	OutputSequenceWasRestarted(output IAVPlayerItemSampleBufferOutput)
}

// AVPlayerItemSampleBufferOutputDelegateObject wraps an existing Objective-C object that conforms to the AVPlayerItemSampleBufferOutputDelegate protocol.
type AVPlayerItemSampleBufferOutputDelegateObject struct {
	objectivec.Object
}

func (o AVPlayerItemSampleBufferOutputDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVPlayerItemSampleBufferOutputDelegateObjectFromID constructs a [AVPlayerItemSampleBufferOutputDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVPlayerItemSampleBufferOutputDelegateObjectFromID(id objc.ID) AVPlayerItemSampleBufferOutputDelegateObject {
	return AVPlayerItemSampleBufferOutputDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Invoked when the output becomes ready to deliver a sample buffer.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSampleBufferOutputDelegate/outputMediaDataAvailable:
func (o AVPlayerItemSampleBufferOutputDelegateObject) OutputMediaDataAvailable(output IAVPlayerItemSampleBufferOutput) {
	objc.Send[struct{}](o.ID, objc.Sel("outputMediaDataAvailable:"), output)
}

// Invoked when the output is commencing a new sequence.
//
// # Discussion
//
// This method is invoked after seeks and changes in playback direction. If
// you are maintaining any queued future samples previously copied, it may be
// appropriate to discard these upon receiving this message.
//
// Note that delivery of this message may race with calls to
// `-copyNextSampleBuffer`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSampleBufferOutputDelegate/outputSequenceWasRestarted:
func (o AVPlayerItemSampleBufferOutputDelegateObject) OutputSequenceWasRestarted(output IAVPlayerItemSampleBufferOutput) {
	objc.Send[struct{}](o.ID, objc.Sel("outputSequenceWasRestarted:"), output)
}

// AVPlayerItemSampleBufferOutputDelegateConfig holds optional typed callbacks for [AVPlayerItemSampleBufferOutputDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avplayeritemsamplebufferoutputdelegate
type AVPlayerItemSampleBufferOutputDelegateConfig struct {

	// Other Methods
	// OutputMediaDataAvailable — Invoked when the output becomes ready to deliver a sample buffer.
	OutputMediaDataAvailable func(output AVPlayerItemSampleBufferOutput)
	// OutputSequenceWasRestarted — Invoked when the output is commencing a new sequence.
	OutputSequenceWasRestarted func(output AVPlayerItemSampleBufferOutput)
}

// NewAVPlayerItemSampleBufferOutputDelegate creates an Objective-C object implementing the [AVPlayerItemSampleBufferOutputDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AVPlayerItemSampleBufferOutputDelegateObject] satisfies the [AVPlayerItemSampleBufferOutputDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avplayeritemsamplebufferoutputdelegate
func NewAVPlayerItemSampleBufferOutputDelegate(config AVPlayerItemSampleBufferOutputDelegateConfig) AVPlayerItemSampleBufferOutputDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAVPlayerItemSampleBufferOutputDelegate_%d", n)

	var methods []objc.MethodDef

	if config.OutputMediaDataAvailable != nil {
		fn := config.OutputMediaDataAvailable
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outputMediaDataAvailable:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outputID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("AVPlayerItemSampleBufferOutputDelegate", "outputMediaDataAvailable:")
					}
				}()
				output := AVPlayerItemSampleBufferOutputFromID(outputID)
				fn(output)
				_delegateDone = true
			},
		})
	}

	if config.OutputSequenceWasRestarted != nil {
		fn := config.OutputSequenceWasRestarted
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("outputSequenceWasRestarted:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outputID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("AVPlayerItemSampleBufferOutputDelegate", "outputSequenceWasRestarted:")
					}
				}()
				output := AVPlayerItemSampleBufferOutputFromID(outputID)
				fn(output)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("AVPlayerItemSampleBufferOutputDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAVPlayerItemSampleBufferOutputDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AVPlayerItemSampleBufferOutputDelegateObjectFromID(instance)
}
