// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"fmt"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// Protocol for receiving broadcast video output events and data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureBroadcastVideoOutputDelegate
type AVCaptureBroadcastVideoOutputDelegate interface {
	objectivec.IObject
}

// AVCaptureBroadcastVideoOutputDelegateObject wraps an existing Objective-C object that conforms to the AVCaptureBroadcastVideoOutputDelegate protocol.
type AVCaptureBroadcastVideoOutputDelegateObject struct {
	objectivec.Object
}

func (o AVCaptureBroadcastVideoOutputDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVCaptureBroadcastVideoOutputDelegateObjectFromID constructs a [AVCaptureBroadcastVideoOutputDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVCaptureBroadcastVideoOutputDelegateObjectFromID(id objc.ID) AVCaptureBroadcastVideoOutputDelegateObject {
	return AVCaptureBroadcastVideoOutputDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Called when a video frame is dropped during broadcast video output
// processing.
//
// output: The [AVCaptureBroadcastVideoOutput] instance that dropped the video frame.
//
// presentationTimeStamp: The presentation timestamp (PTS) of the dropped video frame.
//
// connection: The [AVCaptureConnection] associated with the dropped video frame.
//
// # Discussion
//
// This method is called whenever the broadcast video output system needs to
// drop a video frame due to performance constraints, destination issues,
// buffer overruns, or encoding failures.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureBroadcastVideoOutputDelegate/broadcastVideoOutput(_:didDropVideoFrameWithPresentationTimeStamp:from:)
func (o AVCaptureBroadcastVideoOutputDelegateObject) BroadcastVideoOutputDidDropVideoFrameWithPresentationTimeStampFromConnection(output IAVCaptureBroadcastVideoOutput, presentationTimeStamp coremedia.CMTime, connection IAVCaptureConnection) {
	objc.Send[struct{}](o.ID, objc.Sel("broadcastVideoOutput:didDropVideoFrameWithPresentationTimeStamp:fromConnection:"), output, presentationTimeStamp, connection)
}

// AVCaptureBroadcastVideoOutputDelegateConfig holds optional typed callbacks for [AVCaptureBroadcastVideoOutputDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avcapturebroadcastvideooutputdelegate
type AVCaptureBroadcastVideoOutputDelegateConfig struct {

	// Other Methods
	// BroadcastVideoOutputDidDropVideoFrameWithPresentationTimeStampFromConnection — Called when a video frame is dropped during broadcast video output processing.
	BroadcastVideoOutputDidDropVideoFrameWithPresentationTimeStampFromConnection func(output AVCaptureBroadcastVideoOutput, presentationTimeStamp coremedia.CMTime, connection AVCaptureConnection)
}

// NewAVCaptureBroadcastVideoOutputDelegate creates an Objective-C object implementing the [AVCaptureBroadcastVideoOutputDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AVCaptureBroadcastVideoOutputDelegateObject] satisfies the [AVCaptureBroadcastVideoOutputDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avcapturebroadcastvideooutputdelegate
func NewAVCaptureBroadcastVideoOutputDelegate(config AVCaptureBroadcastVideoOutputDelegateConfig) AVCaptureBroadcastVideoOutputDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAVCaptureBroadcastVideoOutputDelegate_%d", n)

	var methods []objc.MethodDef

	if config.BroadcastVideoOutputDidDropVideoFrameWithPresentationTimeStampFromConnection != nil {
		fn := config.BroadcastVideoOutputDidDropVideoFrameWithPresentationTimeStampFromConnection
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("broadcastVideoOutput:didDropVideoFrameWithPresentationTimeStamp:fromConnection:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outputID objc.ID, presentationTimeStamp coremedia.CMTime, connectionID objc.ID) {
				// Names which delegate was running if a panic unwinds out of
				// it. The frames between here and the Objective-C caller are
				// runtime and purego dispatch, so without this the traceback
				// never says which selector dispatched. Deliberately no
				// recover: see [objc.NoteDelegatePanic].
				_delegateDone := false
				defer func() {
					if !_delegateDone {
						objc.NoteDelegatePanic("AVCaptureBroadcastVideoOutputDelegate", "broadcastVideoOutput:didDropVideoFrameWithPresentationTimeStamp:fromConnection:")
					}
				}()
				output := AVCaptureBroadcastVideoOutputFromID(outputID)
				connection := AVCaptureConnectionFromID(connectionID)
				fn(output, presentationTimeStamp, connection)
				_delegateDone = true
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("AVCaptureBroadcastVideoOutputDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAVCaptureBroadcastVideoOutputDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AVCaptureBroadcastVideoOutputDelegateObjectFromID(instance)
}
