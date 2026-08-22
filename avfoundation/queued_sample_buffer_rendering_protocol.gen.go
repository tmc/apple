// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Methods you can implement to enqueue sample buffers for presentation.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering
type AVQueuedSampleBufferRendering interface {
	objectivec.IObject

	// Sends a sample buffer to the queue for rendering.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/enqueue(_:)
	EnqueueSampleBuffer(sampleBuffer coremedia.CMSampleBufferRef)

	// Tells the target to invoke a client-supplied block in order to gather sample buffers for playback.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/requestMediaDataWhenReady(on:using:)
	RequestMediaDataWhenReadyOnQueueUsingBlock(queue dispatch.Queue, block VoidHandler)

	// Cancels any current [requestMediaDataWhenReady(on:using:)](<https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/requestMediaDataWhenReady(on:using:)>) call.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/stopRequestingMediaData()
	StopRequestingMediaData()

	// Discards all pending enqueued sample buffers.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/flush()
	Flush()

	// A Boolean value that indicates whether the receiver is able to accept more sample buffers.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/isReadyForMoreMediaData
	IsReadyForMoreMediaData() bool

	// A Boolean value that indicates whether the enqued media meets the required preroll level for reliable playback.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/hasSufficientMediaDataForReliablePlaybackStart
	HasSufficientMediaDataForReliablePlaybackStart() bool

	// The timebase for a renderer.
	//
	// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/timebase
	Timebase() coremedia.CMTimebaseRef
}

// AVQueuedSampleBufferRenderingObject wraps an existing Objective-C object that conforms to the AVQueuedSampleBufferRendering protocol.
type AVQueuedSampleBufferRenderingObject struct {
	objectivec.Object
}

func (o AVQueuedSampleBufferRenderingObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVQueuedSampleBufferRenderingObjectFromID constructs a [AVQueuedSampleBufferRenderingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVQueuedSampleBufferRenderingObjectFromID(id objc.ID) AVQueuedSampleBufferRenderingObject {
	return AVQueuedSampleBufferRenderingObject{
		Object: objectivec.ObjectFromID(id),
	}
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
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/enqueue(_:)
//
// [kCFBooleanTrue]: https://developer.apple.com/documentation/CoreFoundation/kCFBooleanTrue
// [kCMSampleAttachmentKey_DisplayImmediately]: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_DisplayImmediately
// [kCMSampleAttachmentKey_DoNotDisplay]: https://developer.apple.com/documentation/CoreMedia/kCMSampleAttachmentKey_DoNotDisplay
// [kCMSampleBufferAttachmentKey_EmptyMedia]: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_EmptyMedia
func (o AVQueuedSampleBufferRenderingObject) EnqueueSampleBuffer(sampleBuffer coremedia.CMSampleBufferRef) {
	objc.Send[struct{}](o.ID, objc.Sel("enqueueSampleBuffer:"), sampleBuffer)
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
func (o AVQueuedSampleBufferRenderingObject) RequestMediaDataWhenReadyOnQueueUsingBlock(queue dispatch.Queue, block VoidHandler) {
	_block1, _cleanup1 := NewVoidBlock(block)
	defer _cleanup1()
	objc.Send[struct{}](o.ID, objc.Sel("requestMediaDataWhenReadyOnQueue:usingBlock:"), uintptr(queue.Handle()), objc.ID(_block1))
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
func (o AVQueuedSampleBufferRenderingObject) StopRequestingMediaData() {
	objc.Send[struct{}](o.ID, objc.Sel("stopRequestingMediaData"))
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
func (o AVQueuedSampleBufferRenderingObject) Flush() {
	objc.Send[struct{}](o.ID, objc.Sel("flush"))
}

// A Boolean value that indicates whether the receiver is able to accept more
// sample buffers.
//
// # Discussion
//
// An object conforming to [AVQueuedSampleBufferRendering] keeps track of the
// occupancy levels of its internal queues for the benefit of clients that
// enqueue sample buffers from non-real-time sources, for example, clients
// that can supply sample buffers faster than they are consumed, and so need
// to decide when to hold back. Clients enqueueing sample buffers from
// non-real-time sources may hold off from generating or obtaining more sample
// buffers to enqueue when the value of `readyForMoreMediaData` is [NO]. It is
// safe to call [EnqueueSampleBuffer] when `readyForMoreMediaData` is [NO],
// but don’t enqueue sample buffers without bound.
//
// To help with control of the non-real-time supply of sample buffers, clients
// can call [RequestMediaDataWhenReadyOnQueueUsingBlock] in order to specify a
// block that the receiver should invoke whenever it’s ready for sample
// buffers to be appended.
//
// The value of `readyForMoreMediaData` often changes` from [NO] to [YES]
// asynchronously, as previously supplied sample buffers are decoded and
// rendered.
//
// This property is not key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/isReadyForMoreMediaData
func (o AVQueuedSampleBufferRenderingObject) IsReadyForMoreMediaData() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isReadyForMoreMediaData"))
	return bool(rv)
}

// A Boolean value that indicates whether the enqued media meets the required
// preroll level for reliable playback.
//
// # Discussion
//
// Starting playback when this property is false may prevent smooth playback
// following an immediate start.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/hasSufficientMediaDataForReliablePlaybackStart
func (o AVQueuedSampleBufferRenderingObject) HasSufficientMediaDataForReliablePlaybackStart() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("hasSufficientMediaDataForReliablePlaybackStart"))
	return bool(rv)
}

// The timebase for a renderer.
//
// # Discussion
//
// The timebase governs how time stamps are interpreted by the renderer.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/timebase
func (o AVQueuedSampleBufferRenderingObject) Timebase() coremedia.CMTimebaseRef {
	rv := objc.Send[coremedia.CMTimebaseRef](o.ID, objc.Sel("timebase"))
	return coremedia.CMTimebaseRef(rv)
}
