// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/quartzcore"
)

// The class instance for the [AVSampleBufferDisplayLayer] class.
var (
	_AVSampleBufferDisplayLayerClass     AVSampleBufferDisplayLayerClass
	_AVSampleBufferDisplayLayerClassOnce sync.Once
)

func getAVSampleBufferDisplayLayerClass() AVSampleBufferDisplayLayerClass {
	_AVSampleBufferDisplayLayerClassOnce.Do(func() {
		_AVSampleBufferDisplayLayerClass = AVSampleBufferDisplayLayerClass{class: objc.GetClass("AVSampleBufferDisplayLayer")}
	})
	return _AVSampleBufferDisplayLayerClass
}

// GetAVSampleBufferDisplayLayerClass returns the class object for AVSampleBufferDisplayLayer.
func GetAVSampleBufferDisplayLayerClass() AVSampleBufferDisplayLayerClass {
	return getAVSampleBufferDisplayLayerClass()
}

type AVSampleBufferDisplayLayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSampleBufferDisplayLayerClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSampleBufferDisplayLayerClass) Alloc() AVSampleBufferDisplayLayer {
	rv := objc.Send[AVSampleBufferDisplayLayer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that displays compressed or uncompressed video frames.
//
// # Accessing the video renderer
//
//   - [AVSampleBufferDisplayLayer.SampleBufferRenderer]: An object that enqueues video sample buffers for rendering.
//
// # Configuring the layer
//
//   - [AVSampleBufferDisplayLayer.ReadyForDisplay]: A Boolean value that indicates whether the first video frame is ready for display.
//   - [AVSampleBufferDisplayLayer.ControlTimebase]: A timebase that determines how the layer interprets timestamps.
//   - [AVSampleBufferDisplayLayer.SetControlTimebase]
//   - [AVSampleBufferDisplayLayer.VideoGravity]: A value that indicates how the layer displays video within its bounds.
//   - [AVSampleBufferDisplayLayer.SetVideoGravity]
//
// # Protecting content
//
//   - [AVSampleBufferDisplayLayer.PreventsCapture]: A Boolean value that indicates whether the layer protects against screen capture.
//   - [AVSampleBufferDisplayLayer.SetPreventsCapture]
//   - [AVSampleBufferDisplayLayer.OutputObscuredDueToInsufficientExternalProtection]: A Boolean value that indicates whether the system obscures decoded output due to insufficient external protection on the current device.
//
// # Preventing backgrounding
//
//   - [AVSampleBufferDisplayLayer.PreventsDisplaySleepDuringVideoPlayback]: A Boolean value that indicates whether the layer prevents the system from sleeping during video playback.
//   - [AVSampleBufferDisplayLayer.SetPreventsDisplaySleepDuringVideoPlayback]
//
// # Handling errors
//
//   - [AVSampleBufferDisplayLayer.AVSampleBufferDisplayLayerFailedToDecode]: A notification the system posts when a sample buffer display layer fails to decode.
//   - [AVSampleBufferDisplayLayer.AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey]: The key for the corresponding error.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer
type AVSampleBufferDisplayLayer struct {
	quartzcore.CALayer
}

// AVSampleBufferDisplayLayerFromID constructs a [AVSampleBufferDisplayLayer] from an objc.ID.
//
// An object that displays compressed or uncompressed video frames.
func AVSampleBufferDisplayLayerFromID(id objc.ID) AVSampleBufferDisplayLayer {
	return AVSampleBufferDisplayLayer{CALayer: quartzcore.CALayerFromID(id)}
}

// NOTE: AVSampleBufferDisplayLayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVSampleBufferDisplayLayer] class.
//
// # Accessing the video renderer
//
//   - [IAVSampleBufferDisplayLayer.SampleBufferRenderer]: An object that enqueues video sample buffers for rendering.
//
// # Configuring the layer
//
//   - [IAVSampleBufferDisplayLayer.ReadyForDisplay]: A Boolean value that indicates whether the first video frame is ready for display.
//   - [IAVSampleBufferDisplayLayer.ControlTimebase]: A timebase that determines how the layer interprets timestamps.
//   - [IAVSampleBufferDisplayLayer.SetControlTimebase]
//   - [IAVSampleBufferDisplayLayer.VideoGravity]: A value that indicates how the layer displays video within its bounds.
//   - [IAVSampleBufferDisplayLayer.SetVideoGravity]
//
// # Protecting content
//
//   - [IAVSampleBufferDisplayLayer.PreventsCapture]: A Boolean value that indicates whether the layer protects against screen capture.
//   - [IAVSampleBufferDisplayLayer.SetPreventsCapture]
//   - [IAVSampleBufferDisplayLayer.OutputObscuredDueToInsufficientExternalProtection]: A Boolean value that indicates whether the system obscures decoded output due to insufficient external protection on the current device.
//
// # Preventing backgrounding
//
//   - [IAVSampleBufferDisplayLayer.PreventsDisplaySleepDuringVideoPlayback]: A Boolean value that indicates whether the layer prevents the system from sleeping during video playback.
//   - [IAVSampleBufferDisplayLayer.SetPreventsDisplaySleepDuringVideoPlayback]
//
// # Handling errors
//
//   - [IAVSampleBufferDisplayLayer.AVSampleBufferDisplayLayerFailedToDecode]: A notification the system posts when a sample buffer display layer fails to decode.
//   - [IAVSampleBufferDisplayLayer.AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey]: The key for the corresponding error.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer
type IAVSampleBufferDisplayLayer interface {
	quartzcore.ICALayer
	AVQueuedSampleBufferRendering

	// Topic: Accessing the video renderer

	// An object that enqueues video sample buffers for rendering.
	SampleBufferRenderer() IAVSampleBufferVideoRenderer

	// Topic: Configuring the layer

	// A Boolean value that indicates whether the first video frame is ready for display.
	ReadyForDisplay() bool
	// A timebase that determines how the layer interprets timestamps.
	ControlTimebase() uintptr
	SetControlTimebase(value uintptr)
	// A value that indicates how the layer displays video within its bounds.
	VideoGravity() AVLayerVideoGravity
	SetVideoGravity(value AVLayerVideoGravity)

	// Topic: Protecting content

	// A Boolean value that indicates whether the layer protects against screen capture.
	PreventsCapture() bool
	SetPreventsCapture(value bool)
	// A Boolean value that indicates whether the system obscures decoded output due to insufficient external protection on the current device.
	OutputObscuredDueToInsufficientExternalProtection() bool

	// Topic: Preventing backgrounding

	// A Boolean value that indicates whether the layer prevents the system from sleeping during video playback.
	PreventsDisplaySleepDuringVideoPlayback() bool
	SetPreventsDisplaySleepDuringVideoPlayback(value bool)

	// Topic: Handling errors

	// A notification the system posts when a sample buffer display layer fails to decode.
	AVSampleBufferDisplayLayerFailedToDecode() foundation.NSString
	// The key for the corresponding error.
	AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey() string

	// The error that caused the failure.
	Error() foundation.INSError
	// A Boolean value that indicates whether the layer needs to flush its state to continue decoding frames.
	RequiresFlushToResumeDecoding() bool
	// The ability of the display layer to be used for enqueuing sample buffers.
	Status() AVQueuedSampleBufferRenderingStatus
}

// Init initializes the instance.
func (s AVSampleBufferDisplayLayer) Init() AVSampleBufferDisplayLayer {
	rv := objc.Send[AVSampleBufferDisplayLayer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSampleBufferDisplayLayer) Autorelease() AVSampleBufferDisplayLayer {
	rv := objc.Send[AVSampleBufferDisplayLayer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSampleBufferDisplayLayer creates a new AVSampleBufferDisplayLayer instance.
func NewAVSampleBufferDisplayLayer() AVSampleBufferDisplayLayer {
	class := getAVSampleBufferDisplayLayerClass()
	rv := objc.Send[AVSampleBufferDisplayLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether the receiver is able to accept more
// sample buffers.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/isReadyForMoreMediaData
func (s AVSampleBufferDisplayLayer) IsReadyForMoreMediaData() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isReadyForMoreMediaData"))
	return rv
}

// An object that enqueues video sample buffers for rendering.
//
// # Discussion
//
// This object allows you to safely enqueue sample buffers from a background
// thread.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/sampleBufferRenderer
func (s AVSampleBufferDisplayLayer) SampleBufferRenderer() IAVSampleBufferVideoRenderer {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("sampleBufferRenderer"))
	return AVSampleBufferVideoRendererFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the first video frame is ready for
// display.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/isReadyForDisplay
func (s AVSampleBufferDisplayLayer) ReadyForDisplay() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isReadyForDisplay"))
	return rv
}

// A timebase that determines how the layer interprets timestamps.
//
// # Discussion
//
// By default, this property is `nil`, which indicates the layer interprets
// timestamps according the host time clock (`mach_absolute_time` with the
// appropriate timescale conversion; this is the same as Core Animation’s
// [CACurrentMediaTime()]). Without a control timebase, it isn’t possible to
// change when the layer displays frames after enqueuing them.
//
// Setting a valid time base enables you to control the timing of frame
// display by setting the rate and time of the control timebase.
//
// If you’re synchronizing video to audio, you should use a timebase whose
// host clock is a [CMClock] for the appropriate audio device to prevent
// drift. See [CMAudioClock] for more information.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/controlTimebase
//
// [CACurrentMediaTime()]: https://developer.apple.com/documentation/QuartzCore/CACurrentMediaTime()
// [CMAudioClock]: https://developer.apple.com/documentation/CoreMedia/cmaudioclock-api
// [CMClock]: https://developer.apple.com/documentation/CoreMedia/CMClock
func (s AVSampleBufferDisplayLayer) ControlTimebase() uintptr {
	rv := objc.Send[uintptr](s.ID, objc.Sel("controlTimebase"))
	return rv
}
func (s AVSampleBufferDisplayLayer) SetControlTimebase(value uintptr) {
	objc.Send[struct{}](s.ID, objc.Sel("setControlTimebase:"), value)
}

// A value that indicates how the layer displays video within its bounds.
//
// # Discussion
//
// [AVLayerVideoGravity] defines the supported video gravities. The default
// value is [resizeAspect].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/videoGravity
//
// [resizeAspect]: https://developer.apple.com/documentation/AVFoundation/AVLayerVideoGravity/resizeAspect
func (s AVSampleBufferDisplayLayer) VideoGravity() AVLayerVideoGravity {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("videoGravity"))
	return AVLayerVideoGravity(foundation.NSStringFromID(rv).String())
}
func (s AVSampleBufferDisplayLayer) SetVideoGravity(value AVLayerVideoGravity) {
	objc.Send[struct{}](s.ID, objc.Sel("setVideoGravity:"), objc.String(string(value)))
}

// A Boolean value that indicates whether the layer protects against screen
// capture.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/preventsCapture
func (s AVSampleBufferDisplayLayer) PreventsCapture() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("preventsCapture"))
	return rv
}
func (s AVSampleBufferDisplayLayer) SetPreventsCapture(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setPreventsCapture:"), value)
}

// A Boolean value that indicates whether the system obscures decoded output
// due to insufficient external protection on the current device.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/isOutputObscuredDueToInsufficientExternalProtection
func (s AVSampleBufferDisplayLayer) OutputObscuredDueToInsufficientExternalProtection() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("outputObscuredDueToInsufficientExternalProtection"))
	return rv
}

// A Boolean value that indicates whether the layer prevents the system from
// sleeping during video playback.
//
// # Discussion
//
// Setting this property to false doesn’t force the display to sleep; it
// only stops preventing display sleep. Other apps or frameworks within your
// app may still be preventing display sleep for various reasons.
//
// The default value is true in iOS, tvOS, and Mac Catalyst. The default value
// in macOS is false.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/preventsDisplaySleepDuringVideoPlayback
func (s AVSampleBufferDisplayLayer) PreventsDisplaySleepDuringVideoPlayback() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("preventsDisplaySleepDuringVideoPlayback"))
	return rv
}
func (s AVSampleBufferDisplayLayer) SetPreventsDisplaySleepDuringVideoPlayback(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setPreventsDisplaySleepDuringVideoPlayback:"), value)
}

// A notification the system posts when a sample buffer display layer fails to
// decode.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/AVSampleBufferDisplayLayerFailedToDecode
func (s AVSampleBufferDisplayLayer) AVSampleBufferDisplayLayerFailedToDecode() foundation.NSString {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("AVSampleBufferDisplayLayerFailedToDecode"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// The key for the corresponding error.
//
// See: https://developer.apple.com/documentation/avfoundation/avsamplebufferdisplaylayerfailedtodecodenotificationerrorkey
func (s AVSampleBufferDisplayLayer) AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("AVSampleBufferDisplayLayerFailedToDecodeNotificationErrorKey"))
	return foundation.NSStringFromID(rv).String()
}

// The error that caused the failure.
//
// # Discussion
//
// The value of this property is an [NSError] that describes what caused the
// display layer to no longer be able to enqueue sample buffers. If the status
// is not [AVQueuedSampleBufferRenderingStatusFailed], the value of this
// property is `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/error
func (s AVSampleBufferDisplayLayer) Error() foundation.INSError {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the enqueued media data meets the
// renderer’s preroll level.
//
// # Discussion
//
// Apple discourages the use of this symbol in iOS 17, tvOS 17, and macOS 14
// and later. Use [HasSufficientMediaDataForReliablePlaybackStart] on the
// [SampleBufferRenderer] instead.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/hasSufficientMediaDataForReliablePlaybackStart
func (s AVSampleBufferDisplayLayer) HasSufficientMediaDataForReliablePlaybackStart() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasSufficientMediaDataForReliablePlaybackStart"))
	return rv
}

// A Boolean value that indicates whether the layer needs to flush its state
// to continue decoding frames.
//
// # Discussion
//
// Apple discourages the use of this symbol in iOS 17, tvOS 17, and macOS 14
// and later. Use [RequiresFlushToResumeDecoding] on the
// [SampleBufferRenderer] instead.
//
// When an app enters a state where use of video decoder resources isn’t
// permissible, the value of this property changes to true and the display
// layer’s status changes to a [AVQueuedSampleBufferRenderingStatusFailed]
// state.
//
// To resume rendering sample buffers using the display layer after this
// property’s value is true, apps must first reset the display layer’s
// status to [AVQueuedSampleBufferRenderingStatusUnknown], which you do by
// calling the layer’s [Flush] method.
//
// This property isn’t key-value observable. Instead, observe changes to
// this property value by observing notifications of type
// [AVSampleBufferDisplayLayerRequiresFlushToResumeDecodingDidChangeNotification].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/requiresFlushToResumeDecoding
//
// [AVSampleBufferDisplayLayerRequiresFlushToResumeDecodingDidChangeNotification]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayerRequiresFlushToResumeDecodingDidChangeNotification
func (s AVSampleBufferDisplayLayer) RequiresFlushToResumeDecoding() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("requiresFlushToResumeDecoding"))
	return rv
}

// The ability of the display layer to be used for enqueuing sample buffers.
//
// # Discussion
//
// Apple discourages the use of this symbol in iOS 17, tvOS 17, and macOS 14
// and later. Use [Status] on [SampleBufferRenderer] instead.
//
// The value of this property is an [AVQueuedSampleBufferRenderingStatus] that
// indicates whether the receiver can be used for enqueuing sample buffers.
//
// When the value of this property is
// [AVQueuedSampleBufferRenderingStatusFailed], the receiver can no longer be
// used and a new instance needs to be created in its place. When this
// happens, clients can check the value of the [Error] property to determine
// the failure.
//
// This property supports key-value observing.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/status
//
// [AVQueuedSampleBufferRenderingStatus]: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRenderingStatus
func (s AVSampleBufferDisplayLayer) Status() AVQueuedSampleBufferRenderingStatus {
	rv := objc.Send[AVQueuedSampleBufferRenderingStatus](s.ID, objc.Sel("status"))
	return AVQueuedSampleBufferRenderingStatus(rv)
}

// The renderer’s timebase, which determines how the layer interprets time
// stamps.
//
// # Discussion
//
// Apple discourages the use of this symbol in iOS 17, tvOS 17, and macOS 14
// and later. Use [Timebase] on the [SampleBufferRenderer] instead.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferDisplayLayer/timebase
func (s AVSampleBufferDisplayLayer) Timebase() uintptr {
	rv := objc.Send[uintptr](s.ID, objc.Sel("timebase"))
	return rv
}

// Protocol methods for AVQueuedSampleBufferRendering
