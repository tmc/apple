// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

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
// [CMSampleBuffer]: https://developer.apple.com/documentation/CoreMedia/CMSampleBuffer
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutputSampleBufferDelegate/captureOutput(_:didOutput:from:)
func (o AVCaptureVideoDataOutputSampleBufferDelegateObject) CaptureOutputDidOutputSampleBufferFromConnection(output IAVCaptureOutput, sampleBuffer objectivec.IObject, connection IAVCaptureConnection) {
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
// [kCMSampleBufferAttachmentKey_DroppedFrameReason]: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_DroppedFrameReason
// [kCMSampleBufferDroppedFrameReason_Discontinuity]: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferDroppedFrameReason_Discontinuity
// [kCMSampleBufferDroppedFrameReason_FrameWasLate]: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferDroppedFrameReason_FrameWasLate
// [kCMSampleBufferDroppedFrameReason_OutOfBuffers]: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferDroppedFrameReason_OutOfBuffers
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutputSampleBufferDelegate/captureOutput(_:didDrop:from:)
func (o AVCaptureVideoDataOutputSampleBufferDelegateObject) CaptureOutputDidDropSampleBufferFromConnection(output IAVCaptureOutput, sampleBuffer objectivec.IObject, connection IAVCaptureConnection) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:didDropSampleBuffer:fromConnection:"), output, sampleBuffer, connection)
	}

