// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

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
func (o AVCaptureAudioDataOutputSampleBufferDelegateObject) CaptureOutputDidOutputSampleBufferFromConnection(output IAVCaptureOutput, sampleBuffer objectivec.IObject, connection IAVCaptureConnection) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:didOutputSampleBuffer:fromConnection:"), output, sampleBuffer, connection)
	}

