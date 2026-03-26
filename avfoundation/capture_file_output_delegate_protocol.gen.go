// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"fmt"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)
var _ = fmt.Sprintf

// Methods for monitoring or controlling the output of a media file capture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutputDelegate
type AVCaptureFileOutputDelegate interface {
	objectivec.IObject

	// Allows a client to opt in to frame accurate recording in [fileOutput(_:didOutputSampleBuffer:from:)](<doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureFileOutputDelegate/fileOutput(_:didOutputSampleBuffer:from:)>).
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutputDelegate/fileOutputShouldProvideSampleAccurateRecordingStart(_:)
	CaptureOutputShouldProvideSampleAccurateRecordingStart(output IAVCaptureOutput) bool
}

// AVCaptureFileOutputDelegateObject wraps an existing Objective-C object that conforms to the AVCaptureFileOutputDelegate protocol.
type AVCaptureFileOutputDelegateObject struct {
	objectivec.Object
}
func (o AVCaptureFileOutputDelegateObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVCaptureFileOutputDelegateObjectFromID constructs a [AVCaptureFileOutputDelegateObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVCaptureFileOutputDelegateObjectFromID(id objc.ID) AVCaptureFileOutputDelegateObject {
	return AVCaptureFileOutputDelegateObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Allows a client to opt in to frame accurate recording in
// [CaptureOutputDidOutputSampleBufferFromConnection].
//
// output: The capture file output instance that is associated with the delegate.
//
// # Return Value
// 
// [true] if frame accurate recording is required; otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// In apps linked before OS X Mountain Lion, delegates that implement the
// [CaptureOutputDidOutputSampleBufferFromConnection] method can ensure that
// starting and stopping a recording is frame accurate by calling
// [StartRecordingToOutputFileURLRecordingDelegate] or [StopRecording] from
// within the callback. Frame accurate recording requires the capture output
// to apply outputSettings when the session starts running, so it is ready to
// start and/or stop recording on any given frame boundary. Applying
// compression settings for the entire length of the session has power,
// thermal, and CPU implications.
// 
// In apps linked on or after OS X Mountain Lion, delegates must implement
// this method to indicate whether frame accurate recording is required. The
// capture file output calls this method only once when the delegate is added
// and never again. If your delegate returns [false], the capture file output
// applies compression settings only when
// [StartRecordingToOutputFileURLRecordingDelegate] is called and disables
// these settings once the recording stops.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutputDelegate/fileOutputShouldProvideSampleAccurateRecordingStart(_:)
func (o AVCaptureFileOutputDelegateObject) CaptureOutputShouldProvideSampleAccurateRecordingStart(output IAVCaptureOutput) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("captureOutputShouldProvideSampleAccurateRecordingStart:"), output)
	return rv
	}
// Gives the delegate the opportunity to inspect samples as they are received
// by the output and start and stop recording at exact times.
//
// output: The capture file output that is receiving the media data.
//
// sampleBuffer: A [CMSampleBuffer] object containing the sample data and additional
// information about the sample, such as its format and presentation time.
//
// connection: The [AVCaptureConnection] object attached to the file output from which the
// sample data was received.
//
// # Discussion
// 
// This method is called whenever the file output receives a single sample
// buffer (a single video frame or audio buffer, for example) from the given
// connection. This gives delegates an opportunity to start and stop recording
// or change output files at an exact sample boundary. If called from within
// this method, the file output’s
// [StartRecordingToOutputFileURLRecordingDelegate] and [ResumeRecording]
// methods are guaranteed to include the received sample buffer in the new
// file, whereas calls to [StopRecording] and [PauseRecording] are guaranteed
// to include all samples leading up to those in the current sample buffer in
// the existing file.
// 
// You can gather information particular to the samples by inspecting the
// [CMSampleBuffer] object. Sample buffers always contain a single frame of
// video if called from this method but may also contain multiple samples of
// audio. For B-frame video formats, samples are always delivered in
// presentation order.
// 
// If you need to reference the [CMSampleBuffer] object outside of the scope
// of this method, you must retain it and then release it when you have
// finished with it.
// 
// To maintain optimal performance, some sample buffers directly reference
// pools of memory that may need to be reused by the device system and other
// capture inputs. This is frequently the case for uncompressed device native
// capture where memory blocks are copied as little as possible. If multiple
// sample buffers reference such pools of memory for too long, inputs will no
// longer be able to copy new samples into memory and those samples will be
// dropped. If your application is causing samples to be dropped by retaining
// the provided [CMSampleBuffer] objects for too long, but it needs access to
// the sample data for a long period of time, consider copying the data into a
// new buffer and then calling [CFRelease] on the sample buffer (if it was
// previously retained) so that the memory it references can be reused.
// 
// You should not assume that this method will be called on a specific thread.
// In addition, this method is called frequently, so it must be efficient to
// prevent capture performance problems.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutputDelegate/fileOutput(_:didOutputSampleBuffer:from:)
func (o AVCaptureFileOutputDelegateObject) CaptureOutputDidOutputSampleBufferFromConnection(output IAVCaptureFileOutput, sampleBuffer objectivec.IObject, connection IAVCaptureConnection) {
	objc.Send[struct{}](o.ID, objc.Sel("captureOutput:didOutputSampleBuffer:fromConnection:"), output, sampleBuffer, connection)
	}

// AVCaptureFileOutputDelegateConfig holds optional typed callbacks for [AVCaptureFileOutputDelegate] methods.
// Set non-nil fields to register the corresponding Objective-C delegate method.
// Methods with nil callbacks are not registered, so [NSObject.RespondsToSelector]
// returns false for them — matching the Objective-C delegate pattern exactly.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputdelegate
type AVCaptureFileOutputDelegateConfig struct {

	// Other Methods
	// CaptureOutputShouldProvideSampleAccurateRecordingStart — Allows a client to opt in to frame accurate recording in [fileOutput(_:didOutputSampleBuffer:from:)](<doc://com.apple.avfoundation/documentation/AVFoundation/AVCaptureFileOutputDelegate/fileOutput(_:didOutputSampleBuffer:from:)>).
	CaptureOutputShouldProvideSampleAccurateRecordingStart func(output AVCaptureOutput) bool
}

// NewAVCaptureFileOutputDelegate creates an Objective-C object implementing the [AVCaptureFileOutputDelegate] protocol.
//
// Each call registers a unique Objective-C class containing only the methods
// set in config. This means [NSObject.RespondsToSelector] works correctly
// for optional delegate methods — only non-nil callbacks are registered.
//
// The returned [AVCaptureFileOutputDelegateObject] satisfies the [AVCaptureFileOutputDelegate] interface
// and can be passed directly to SetDelegate and similar methods.
//
// See [Apple Documentation] for protocol details.
//
// [Apple Documentation]: https://developer.apple.com/documentation/avfoundation/avcapturefileoutputdelegate
func NewAVCaptureFileOutputDelegate(config AVCaptureFileOutputDelegateConfig) AVCaptureFileOutputDelegateObject {
	n := delegateClassCounter.Add(1)
	className := fmt.Sprintf("GoAVCaptureFileOutputDelegate_%d", n)

	var methods []objc.MethodDef

	if config.CaptureOutputShouldProvideSampleAccurateRecordingStart != nil {
		fn := config.CaptureOutputShouldProvideSampleAccurateRecordingStart
		methods = append(methods, objc.MethodDef{
			Cmd: objc.RegisterName("captureOutputShouldProvideSampleAccurateRecordingStart:"),
			Fn: func(self objc.ID, _cmd objc.SEL, outputID objc.ID) bool {
				output := AVCaptureOutputFromID(outputID)
				return fn(output)
			},
		})
	}

	nsObjectClass := objc.GetClass("NSObject")
	proto := objc.GetProtocol("AVCaptureFileOutputDelegate")

	var protocols []*objc.Protocol
	if proto != nil {
		protocols = append(protocols, proto)
	}

	cls, err := objc.RegisterClass(className, nsObjectClass, protocols, nil, methods)
	if err != nil {
		panic(fmt.Sprintf("NewAVCaptureFileOutputDelegate: RegisterClass %s: %v", className, err))
	}

	instance := objc.ID(cls).Send(objc.RegisterName("alloc")).Send(objc.RegisterName("init"))
	return AVCaptureFileOutputDelegateObjectFromID(instance)
}

