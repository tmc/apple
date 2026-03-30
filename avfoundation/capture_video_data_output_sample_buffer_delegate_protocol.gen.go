// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"fmt"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

var _ = fmt.Sprintf

// Methods for receiving sample buffers from, and monitoring the status of, a video data output.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutputSampleBufferDelegate
type AVCaptureVideoDataOutputSampleBufferDelegate interface {
	objectivec.IObject
}

// AVCaptureVideoDataOutputSampleBufferDelegateObject wraps an existing Objective-C object that conforms to the AVCaptureVideoDataOutputSampleBufferDelegate protocol.
type AVCaptureVideoDataOutputSampleBufferDelegateObject struct {
	objectivec.Object
}

func (o AVCaptureVideoDataOutputSampleBufferDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVCaptureVideoDataOutputSampleBufferDelegateObjectFromID constructs a [AVCaptureVideoDataOutputSampleBufferDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVCaptureVideoDataOutputSampleBufferDelegateObjectFromID(id objc.ID) AVCaptureVideoDataOutputSampleBufferDelegateObject {
	return AVCaptureVideoDataOutputSampleBufferDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Notifies the delegate that a new video frame was written.
//
// output: The capture output object.
//
// sampleBuffer: A [CMSampleBuffer] object containing the video frame data and additional
// information about the frame, such as its format and presentation time.
//
// connection: The connection from which the video was received.
//
// # Discussion
//
// Delegates receive this message whenever the output captures and outputs a
// new video frame, decoding or re-encoding it as specified by its
// `videoSettings` property. Delegates can use the provided video frame in
// conjunction with other APIs for further processing.
//
// This method is called on the dispatch queue specified by the output’s
// [SampleBufferCallbackQueue] property. It is called periodically, so it must
// be efficient to prevent capture performance problems, including dropped
// frames.
//
// If you need to reference the [CMSampleBuffer] object outside of the scope
// of this method, you must [CFRetain] it and then [CFRelease] it when you are
// finished with it.
//
// To maintain optimal performance, some sample buffers directly reference
// pools of memory that may need to be reused by the device system and other
// capture inputs. This is frequently the case for uncompressed device native
// capture where memory blocks are copied as little as possible. If multiple
// sample buffers reference such pools of memory for too long, inputs will no
// longer be able to copy new samples into memory and those samples will be
// dropped.
//
// If your application is causing samples to be dropped by retaining the
// provided [CMSampleBuffer] objects for too long, but it needs access to the
// sample data for a long period of time, consider copying the data into a new
// buffer and then releasing the sample buffer (if it was previously retained)
// so that the memory it references can be reused.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutputSampleBufferDelegate/captureOutput(_:didOutput:from:)
//
// [CMSampleBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer
func (o AVCaptureVideoDataOutputSampleBufferDelegateObject) CaptureOutputDidOutputSampleBufferFromConnection(output IAVCaptureOutput, sampleBuffer uintptr, connection IAVCaptureConnection) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:didOutputSampleBuffer:fromConnection:"), output, sampleBuffer, connection)
}

// Notifies the delegate that a video frame was discarded.
//
// output: The capture output object.
//
// sampleBuffer: A [CMSampleBuffer] object containing information about the dropped frame,
// such as its format and presentation time.
//
// This sample buffer contains none of the original video data.
//
// connection: The connection from which the video was received.
//
// # Discussion
//
// Delegates receive this message whenever a late video frame is dropped. This
// method is called once for each dropped frame. It is called on the dispatch
// queue specified by the output’s [SampleBufferCallbackQueue] property.
//
// The `sampleBuffer` will contain a
// [kCMSampleBufferAttachmentKey_DroppedFrameReason] attachment that details
// why the frame was dropped. The frame may be dropped because it was late
// ([kCMSampleBufferDroppedFrameReason_FrameWasLate]), typically caused by the
// client’s processing taking too long. It can also be dropped because the
// module providing frames is out of buffers
// ([kCMSampleBufferDroppedFrameReason_OutOfBuffers]). Frames can also be
// dropped if the module providing sample buffers has experienced a
// discontinuity ([kCMSampleBufferDroppedFrameReason_Discontinuity]) and an
// unknown number of frames have been lost. This condition is typically caused
// by the system being too busy.
//
// Because this method is called on the same dispatch queue that is
// responsible for outputting video frames, it must be efficient to prevent
// further capture performance problems, such as additional dropped video
// frames.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutputSampleBufferDelegate/captureOutput(_:didDrop:from:)
//
// [kCMSampleBufferAttachmentKey_DroppedFrameReason]: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_DroppedFrameReason
// [kCMSampleBufferDroppedFrameReason_Discontinuity]: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferDroppedFrameReason_Discontinuity
// [kCMSampleBufferDroppedFrameReason_FrameWasLate]: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferDroppedFrameReason_FrameWasLate
// [kCMSampleBufferDroppedFrameReason_OutOfBuffers]: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferDroppedFrameReason_OutOfBuffers
func (o AVCaptureVideoDataOutputSampleBufferDelegateObject) CaptureOutputDidDropSampleBufferFromConnection(output IAVCaptureOutput, sampleBuffer uintptr, connection IAVCaptureConnection) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:didDropSampleBuffer:fromConnection:"), output, sampleBuffer, connection)
}

// AVCaptureVideoDataOutputSampleBufferDelegateConfig holds optional typed callbacks for [AVCaptureVideoDataOutputSampleBufferDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutputsamplebufferdelegate
type AVCaptureVideoDataOutputSampleBufferDelegateConfig struct {

	// Other Methods
	// CaptureOutputDidOutputSampleBufferFromConnection — Notifies the delegate that a new video frame was written.
	CaptureOutputDidOutputSampleBufferFromConnection func(output AVCaptureOutput, sampleBuffer uintptr, connection AVCaptureConnection)
	// CaptureOutputDidDropSampleBufferFromConnection — Notifies the delegate that a video frame was discarded.
	CaptureOutputDidDropSampleBufferFromConnection func(output AVCaptureOutput, sampleBuffer uintptr, connection AVCaptureConnection)
}

// NewAVCaptureVideoDataOutputSampleBufferDelegate creates an Objective-C object implementing the [AVCaptureVideoDataOutputSampleBufferDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AVCaptureVideoDataOutputSampleBufferDelegateObject] satisfies the [AVCaptureVideoDataOutputSampleBufferDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avcapturevideodataoutputsamplebufferdelegate
func NewAVCaptureVideoDataOutputSampleBufferDelegate(config AVCaptureVideoDataOutputSampleBufferDelegateConfig) AVCaptureVideoDataOutputSampleBufferDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAVCaptureVideoDataOutputSampleBufferDelegate_%d", n)

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

	if config.CaptureOutputDidDropSampleBufferFromConnection != nil {
		fn := config.CaptureOutputDidDropSampleBufferFromConnection
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("captureOutput:didDropSampleBuffer:fromConnection:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outputID objc.ID, sampleBuffer uintptr, connectionID objc.ID) {
				output := AVCaptureOutputFromID(outputID)
				connection := AVCaptureConnectionFromID(connectionID)
				fn(output, sampleBuffer, connection)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("AVCaptureVideoDataOutputSampleBufferDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAVCaptureVideoDataOutputSampleBufferDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AVCaptureVideoDataOutputSampleBufferDelegateObjectFromID(instance)
}
