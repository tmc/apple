// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSampleBufferVideoRenderer] class.
var (
	_AVSampleBufferVideoRendererClass     AVSampleBufferVideoRendererClass
	_AVSampleBufferVideoRendererClassOnce sync.Once
)

func getAVSampleBufferVideoRendererClass() AVSampleBufferVideoRendererClass {
	_AVSampleBufferVideoRendererClassOnce.Do(func() {
		_AVSampleBufferVideoRendererClass = AVSampleBufferVideoRendererClass{class: objc.GetClass("AVSampleBufferVideoRenderer")}
	})
	return _AVSampleBufferVideoRendererClass
}

// GetAVSampleBufferVideoRendererClass returns the class object for AVSampleBufferVideoRenderer.
func GetAVSampleBufferVideoRendererClass() AVSampleBufferVideoRendererClass {
	return getAVSampleBufferVideoRendererClass()
}

type AVSampleBufferVideoRendererClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSampleBufferVideoRendererClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSampleBufferVideoRendererClass) Alloc() AVSampleBufferVideoRenderer {
	rv := objc.Send[AVSampleBufferVideoRenderer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that enqueues video sample buffers for rendering.
//
// # Flushing the renderer
//
//   - [AVSampleBufferVideoRenderer.RequiresFlushToResumeDecoding]: A Boolean value that Indicates whether the renderer requires flushing to continue decoding frames.
//   - [AVSampleBufferVideoRenderer.FlushWithRemovalOfDisplayedImageCompletionHandler]: Tells the video renderer to discard pending enqueued sample buffers.
//
// # Setting presentation time expectations
//
//   - [AVSampleBufferVideoRenderer.PresentationTimeExpectation]
//   - [AVSampleBufferVideoRenderer.SetPresentationTimeExpectation]
//
// # Inspecting the status
//
//   - [AVSampleBufferVideoRenderer.Status]: A status value that indicates whether this object can enqueue and render sample buffers.
//   - [AVSampleBufferVideoRenderer.Error]: An object the describes the error that caused the rendering failure.
//
// # Accessing the pixel buffer
//
//   - [AVSampleBufferVideoRenderer.CopyDisplayedPixelBuffer]
//
// # Capturing performance metrics
//
//   - [AVSampleBufferVideoRenderer.LoadVideoPerformanceMetricsWithCompletionHandler]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer
type AVSampleBufferVideoRenderer struct {
	objectivec.Object
}

// AVSampleBufferVideoRendererFromID constructs a [AVSampleBufferVideoRenderer] from an objc.ID.
//
// An object that enqueues video sample buffers for rendering.
func AVSampleBufferVideoRendererFromID(id objc.ID) AVSampleBufferVideoRenderer {
	return AVSampleBufferVideoRenderer{objectivec.Object{ID: id}}
}
// NOTE: AVSampleBufferVideoRenderer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVSampleBufferVideoRenderer] class.
//
// # Flushing the renderer
//
//   - [IAVSampleBufferVideoRenderer.RequiresFlushToResumeDecoding]: A Boolean value that Indicates whether the renderer requires flushing to continue decoding frames.
//   - [IAVSampleBufferVideoRenderer.FlushWithRemovalOfDisplayedImageCompletionHandler]: Tells the video renderer to discard pending enqueued sample buffers.
//
// # Setting presentation time expectations
//
//   - [IAVSampleBufferVideoRenderer.PresentationTimeExpectation]
//   - [IAVSampleBufferVideoRenderer.SetPresentationTimeExpectation]
//
// # Inspecting the status
//
//   - [IAVSampleBufferVideoRenderer.Status]: A status value that indicates whether this object can enqueue and render sample buffers.
//   - [IAVSampleBufferVideoRenderer.Error]: An object the describes the error that caused the rendering failure.
//
// # Accessing the pixel buffer
//
//   - [IAVSampleBufferVideoRenderer.CopyDisplayedPixelBuffer]
//
// # Capturing performance metrics
//
//   - [IAVSampleBufferVideoRenderer.LoadVideoPerformanceMetricsWithCompletionHandler]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer
type IAVSampleBufferVideoRenderer interface {
	objectivec.IObject
	AVQueuedSampleBufferRendering

	// Topic: Flushing the renderer

	// A Boolean value that Indicates whether the renderer requires flushing to continue decoding frames.
	RequiresFlushToResumeDecoding() bool
	// Tells the video renderer to discard pending enqueued sample buffers.
	FlushWithRemovalOfDisplayedImageCompletionHandler(removeDisplayedImage bool, handler VoidHandler)

	// Topic: Setting presentation time expectations

	PresentationTimeExpectation() objectivec.IObject
	SetPresentationTimeExpectation(value objectivec.IObject)

	// Topic: Inspecting the status

	// A status value that indicates whether this object can enqueue and render sample buffers.
	Status() AVQueuedSampleBufferRenderingStatus
	// An object the describes the error that caused the rendering failure.
	Error() foundation.INSError

	// Topic: Accessing the pixel buffer

	CopyDisplayedPixelBuffer() corevideo.CVImageBufferRef

	// Topic: Capturing performance metrics

	LoadVideoPerformanceMetricsWithCompletionHandler(completionHandler AVVideoPerformanceMetricsHandler)

	RecommendedPixelBufferAttributes() foundation.INSDictionary
	ExpectMinimumUpcomingSampleBufferPresentationTime(minimumUpcomingPresentationTime coremedia.CMTime)
	ExpectMonotonicallyIncreasingUpcomingSampleBufferPresentationTimes()
	ResetUpcomingSampleBufferPresentationTimeExpectations()
}

// Init initializes the instance.
func (s AVSampleBufferVideoRenderer) Init() AVSampleBufferVideoRenderer {
	rv := objc.Send[AVSampleBufferVideoRenderer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSampleBufferVideoRenderer) Autorelease() AVSampleBufferVideoRenderer {
	rv := objc.Send[AVSampleBufferVideoRenderer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSampleBufferVideoRenderer creates a new AVSampleBufferVideoRenderer instance.
func NewAVSampleBufferVideoRenderer() AVSampleBufferVideoRenderer {
	class := getAVSampleBufferVideoRendererClass()
	rv := objc.Send[AVSampleBufferVideoRenderer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Tells the video renderer to discard pending enqueued sample buffers.
//
// removeDisplayedImage: A Boolean value that indicates whether to remove the display image.
//
// handler: A completion handler the system invokes when the flush completes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/flush(removingDisplayedImage:completionHandler:)
func (s AVSampleBufferVideoRenderer) FlushWithRemovalOfDisplayedImageCompletionHandler(removeDisplayedImage bool, handler VoidHandler) {
_block1, _ := NewVoidBlock(handler)
	objc.Send[objc.ID](s.ID, objc.Sel("flushWithRemovalOfDisplayedImage:completionHandler:"), removeDisplayedImage, _block1)
}
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/displayedPixelBuffer()
func (s AVSampleBufferVideoRenderer) CopyDisplayedPixelBuffer() corevideo.CVImageBufferRef {
	rv := objc.Send[corevideo.CVImageBufferRef](s.ID, objc.Sel("copyDisplayedPixelBuffer"))
	return corevideo.CVImageBufferRef(rv)
}
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/loadVideoPerformanceMetrics(completionHandler:)
func (s AVSampleBufferVideoRenderer) LoadVideoPerformanceMetricsWithCompletionHandler(completionHandler AVVideoPerformanceMetricsHandler) {
_block0, _ := NewAVVideoPerformanceMetricsBlock(completionHandler)
	objc.Send[objc.ID](s.ID, objc.Sel("loadVideoPerformanceMetricsWithCompletionHandler:"), _block0)
}
// Sends a sample buffer to the queue for rendering.
//
// sampleBuffer: The sample buffer to be enqueued.
//
// # Discussion
// 
// For video data, the sample buffer is processed according to the attachments
// it contains. If it has a true value for its
// [kCMSampleAttachmentKey_DoNotDisplay] attachment, the frame is decoded but
// not displayed. If it has a `true` value for its
// [kCMSampleAttachmentKey_DisplayImmediately] attachment, the frame is
// displayed as soon as possible, regardless of its presentation timestamp.
// Otherwise, the frame is displayed according to its presentation timestamp,
// relative to the timebase.
// 
// To schedule the removal of previous images at a specific timestamp, enqueue
// a marker sample buffer that doesn’t contain any samples, with the
// [kCMSampleBufferAttachmentKey_EmptyMedia] attachment set to
// [kCFBooleanTrue].
//
// [kCFBooleanTrue]: https://developer.apple.com/documentation/CoreFoundation/kCFBooleanTrue
// [kCMSampleAttachmentKey_DisplayImmediately]: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_DisplayImmediately
// [kCMSampleAttachmentKey_DoNotDisplay]: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_DoNotDisplay
// [kCMSampleBufferAttachmentKey_EmptyMedia]: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_EmptyMedia
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/enqueue(_:)
func (s AVSampleBufferVideoRenderer) EnqueueSampleBuffer(sampleBuffer uintptr) {
	objc.Send[objc.ID](s.ID, objc.Sel("enqueueSampleBuffer:"), sampleBuffer)
}
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/expectMinimumUpcomingSampleBufferPresentationTime:
func (s AVSampleBufferVideoRenderer) ExpectMinimumUpcomingSampleBufferPresentationTime(minimumUpcomingPresentationTime coremedia.CMTime) {
	objc.Send[objc.ID](s.ID, objc.Sel("expectMinimumUpcomingSampleBufferPresentationTime:"), minimumUpcomingPresentationTime)
}
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/expectMonotonicallyIncreasingUpcomingSampleBufferPresentationTimes
func (s AVSampleBufferVideoRenderer) ExpectMonotonicallyIncreasingUpcomingSampleBufferPresentationTimes() {
	objc.Send[objc.ID](s.ID, objc.Sel("expectMonotonicallyIncreasingUpcomingSampleBufferPresentationTimes"))
}
// Discards all pending enqueued sample buffers.
//
// # Discussion
// 
// It is not possible to determine which sample buffers have been decoded for
// video. The next frame passed to [EnqueueSampleBuffer] should be an IDR
// frame (also known as a key frame or sync sample).
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/flush()
func (s AVSampleBufferVideoRenderer) Flush() {
	objc.Send[objc.ID](s.ID, objc.Sel("flush"))
}
// A Boolean value that indicates whether the receiver is able to accept more
// sample buffers.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/isReadyForMoreMediaData
func (s AVSampleBufferVideoRenderer) IsReadyForMoreMediaData() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isReadyForMoreMediaData"))
	return rv
}
// Tells the target to invoke a client-supplied block in order to gather
// sample buffers for playback.
//
// queue: The dispatch queue.
//
// block: A block that enqueues sample buffers until the receiver is no longer ready
// or there is no more data to supply.
//
// # Discussion
// 
// When this method is called multiple times, only the last call is
// implemented. Pair each call to [RequestMediaDataWhenReadyOnQueueUsingBlock]
// with a corresponding call to [StopRequestingMediaData]. Releasing the
// [AVQueuedSampleBufferRendering] object without a call to
// `stopRequestingMediaData` results in undefined behavior.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/requestMediaDataWhenReady(on:using:)
func (s AVSampleBufferVideoRenderer) RequestMediaDataWhenReadyOnQueueUsingBlock(queue dispatch.Queue, block VoidHandler) {
_block1, _ := NewVoidBlock(block)
	objc.Send[objc.ID](s.ID, objc.Sel("requestMediaDataWhenReadyOnQueue:usingBlock:"), uintptr(queue.Handle()), _block1)
}
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/resetUpcomingSampleBufferPresentationTimeExpectations
func (s AVSampleBufferVideoRenderer) ResetUpcomingSampleBufferPresentationTimeExpectations() {
	objc.Send[objc.ID](s.ID, objc.Sel("resetUpcomingSampleBufferPresentationTimeExpectations"))
}
// Cancels any current [RequestMediaDataWhenReadyOnQueueUsingBlock] call.
//
// # Discussion
// 
// Always pair a call to [RequestMediaDataWhenReadyOnQueueUsingBlock] with
// this method. You can call this method from inside or outside of the
// requesting method’s block parameter.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/stopRequestingMediaData()
func (s AVSampleBufferVideoRenderer) StopRequestingMediaData() {
	objc.Send[objc.ID](s.ID, objc.Sel("stopRequestingMediaData"))
}

// A Boolean value that Indicates whether the renderer requires flushing to
// continue decoding frames.
//
// # Discussion
// 
// When your app enters a state where using a video decoder resources is not
// permissible, the value of this property changes to [true] along with the
// video renderer’s status changing to
// [QueuedSampleBufferRenderingStatusFailed]. To resume rendering sample
// buffers, you must first reset the video renderer by calling [Flush] or
// [FlushWithRemovalOfDisplayedImageCompletionHandler].
// 
// This property is not key-value observable. Instead, track changes to this
// property by observing notifications of type
// [requiresFlushToResumeDecodingDidChangeNotification].
//
// [requiresFlushToResumeDecodingDidChangeNotification]: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/requiresFlushToResumeDecodingDidChangeNotification
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/requiresFlushToResumeDecoding
func (s AVSampleBufferVideoRenderer) RequiresFlushToResumeDecoding() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("requiresFlushToResumeDecoding"))
	return rv
}
// See: https://developer.apple.com/documentation/avfoundation/avsamplebuffervideorenderer/presentationtimeexpectation-swift.property
func (s AVSampleBufferVideoRenderer) PresentationTimeExpectation() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("presentationTimeExpectation"))
	return objectivec.Object{ID: rv}
}
func (s AVSampleBufferVideoRenderer) SetPresentationTimeExpectation(value objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("setPresentationTimeExpectation:"), value)
}
// A status value that indicates whether this object can enqueue and render
// sample buffers.
//
// # Discussion
// 
// If the status is [QueuedSampleBufferRenderingStatusFailed], check the value
// of the [Error] property to determine the failure. To resume rendering
// sample buffers after a failure, you must first reset the status to
// [QueuedSampleBufferRenderingStatusUnknown], which you do by invoking
// [Flush] on the video renderer.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/status
func (s AVSampleBufferVideoRenderer) Status() AVQueuedSampleBufferRenderingStatus {
	rv := objc.Send[AVQueuedSampleBufferRenderingStatus](s.ID, objc.Sel("status"))
	return AVQueuedSampleBufferRenderingStatus(rv)
}
// An object the describes the error that caused the rendering failure.
//
// # Discussion
// 
// This value is `nil` by default. It only contains a valid error object when
// the [Status] value is [QueuedSampleBufferRenderingStatusFailed].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/error
func (s AVSampleBufferVideoRenderer) Error() foundation.INSError {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("error"))
	return foundation.NSErrorFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the enqued media meets the required
// preroll level for reliable playback.
//
// # Discussion
// 
// Starting playback when this property is [false] may prevent smooth playback
// following an immediate start.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/hasSufficientMediaDataForReliablePlaybackStart
func (s AVSampleBufferVideoRenderer) HasSufficientMediaDataForReliablePlaybackStart() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasSufficientMediaDataForReliablePlaybackStart"))
	return rv
}
//
// # Discussion
// 
// Recommended pixel buffer attributes for optimal performance when using
// CMSampleBuffers containing CVPixelBuffers.
// 
// The returned dictionary does not contain all of the attributes needed for
// creating pixel buffers. Use
// `CVPixelBufferCreateResolvedAttributesDictionary()` to reconcile these
// attributes with the pixel buffer creation attributes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferVideoRenderer/recommendedPixelBufferAttributes-6326f
func (s AVSampleBufferVideoRenderer) RecommendedPixelBufferAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("recommendedPixelBufferAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// The timebase for a renderer.
//
// # Discussion
// 
// The timebase governs how time stamps are interpreted by the renderer.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/timebase
func (s AVSampleBufferVideoRenderer) Timebase() uintptr {
	rv := objc.Send[uintptr](s.ID, objc.Sel("timebase"))
	return rv
}

			// Protocol methods for AVQueuedSampleBufferRendering
			

// FlushWithRemovalOfDisplayedImage is a synchronous wrapper around [AVSampleBufferVideoRenderer.FlushWithRemovalOfDisplayedImageCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s AVSampleBufferVideoRenderer) FlushWithRemovalOfDisplayedImage(ctx context.Context, removeDisplayedImage bool) error {
	done := make(chan struct{}, 1)
	s.FlushWithRemovalOfDisplayedImageCompletionHandler(removeDisplayedImage, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// LoadVideoPerformanceMetrics is a synchronous wrapper around [AVSampleBufferVideoRenderer.LoadVideoPerformanceMetricsWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s AVSampleBufferVideoRenderer) LoadVideoPerformanceMetrics(ctx context.Context) (*AVVideoPerformanceMetrics, error) {
	done := make(chan *AVVideoPerformanceMetrics, 1)
	s.LoadVideoPerformanceMetricsWithCompletionHandler(func(val *AVVideoPerformanceMetrics) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// RequestMediaDataWhenReadyOnQueueUsingBlockSync is a synchronous wrapper around [AVSampleBufferVideoRenderer.RequestMediaDataWhenReadyOnQueueUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s AVSampleBufferVideoRenderer) RequestMediaDataWhenReadyOnQueueUsingBlockSync(ctx context.Context, queue dispatch.Queue) error {
	done := make(chan struct{}, 1)
	s.RequestMediaDataWhenReadyOnQueueUsingBlock(queue, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

