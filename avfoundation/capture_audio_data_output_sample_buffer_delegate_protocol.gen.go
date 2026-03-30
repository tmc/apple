// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// Methods for receiving audio sample data from an audio capture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioDataOutputSampleBufferDelegate
type AVCaptureAudioDataOutputSampleBufferDelegate interface {
	objectivec.IObject
}

// AVCaptureAudioDataOutputSampleBufferDelegateObject wraps an existing Objective-C object that conforms to the AVCaptureAudioDataOutputSampleBufferDelegate protocol.
type AVCaptureAudioDataOutputSampleBufferDelegateObject struct {
	objectivec.Object
}

func (o AVCaptureAudioDataOutputSampleBufferDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVCaptureAudioDataOutputSampleBufferDelegateObjectFromID constructs a [AVCaptureAudioDataOutputSampleBufferDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVCaptureAudioDataOutputSampleBufferDelegateObjectFromID(id objc.ID) AVCaptureAudioDataOutputSampleBufferDelegateObject {
	return AVCaptureAudioDataOutputSampleBufferDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Notifies the delegate that a sample buffer was written.
//
// output: The capture output object.
//
// sampleBuffer: The sample buffer that was output.
//
// connection: The connection.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureAudioDataOutputSampleBufferDelegate/captureOutput(_:didOutput:from:)
func (o AVCaptureAudioDataOutputSampleBufferDelegateObject) CaptureOutputDidOutputSampleBufferFromConnection(output IAVCaptureOutput, sampleBuffer uintptr, connection IAVCaptureConnection) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:didOutputSampleBuffer:fromConnection:"), output, sampleBuffer, connection)
}

// AVCaptureAudioDataOutputSampleBufferDelegateConfig holds optional typed callbacks for [AVCaptureAudioDataOutputSampleBufferDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutputsamplebufferdelegate
type AVCaptureAudioDataOutputSampleBufferDelegateConfig struct {

	// Other Methods
	// CaptureOutputDidOutputSampleBufferFromConnection — Notifies the delegate that a sample buffer was written.
	CaptureOutputDidOutputSampleBufferFromConnection func(output AVCaptureOutput, sampleBuffer uintptr, connection AVCaptureConnection)
}

// NewAVCaptureAudioDataOutputSampleBufferDelegate creates an Objective-C object implementing the [AVCaptureAudioDataOutputSampleBufferDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AVCaptureAudioDataOutputSampleBufferDelegateObject] satisfies the [AVCaptureAudioDataOutputSampleBufferDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avcaptureaudiodataoutputsamplebufferdelegate
func NewAVCaptureAudioDataOutputSampleBufferDelegate(config AVCaptureAudioDataOutputSampleBufferDelegateConfig) AVCaptureAudioDataOutputSampleBufferDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAVCaptureAudioDataOutputSampleBufferDelegate_%d", n)

	var methods []objc.MethodDef

	if config.CaptureOutputDidOutputSampleBufferFromConnection != nil {
		fn := config.CaptureOutputDidOutputSampleBufferFromConnection
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("captureOutput:didOutputSampleBuffer:fromConnection:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outputID objc.ID, sampleBuffer uintptr, connectionID objc.ID) {
				output := AVCaptureOutputFromID(outputID)
				connection := AVCaptureConnectionFromID(connectionID)
				fn(output, sampleBuffer, connection)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("AVCaptureAudioDataOutputSampleBufferDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAVCaptureAudioDataOutputSampleBufferDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AVCaptureAudioDataOutputSampleBufferDelegateObjectFromID(instance)
}
