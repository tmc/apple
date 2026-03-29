// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSampleBufferAudioRenderer] class.
var (
	_AVSampleBufferAudioRendererClass     AVSampleBufferAudioRendererClass
	_AVSampleBufferAudioRendererClassOnce sync.Once
)

func getAVSampleBufferAudioRendererClass() AVSampleBufferAudioRendererClass {
	_AVSampleBufferAudioRendererClassOnce.Do(func() {
		_AVSampleBufferAudioRendererClass = AVSampleBufferAudioRendererClass{class: objc.GetClass("AVSampleBufferAudioRenderer")}
	})
	return _AVSampleBufferAudioRendererClass
}

// GetAVSampleBufferAudioRendererClass returns the class object for AVSampleBufferAudioRenderer.
func GetAVSampleBufferAudioRendererClass() AVSampleBufferAudioRendererClass {
	return getAVSampleBufferAudioRendererClass()
}

type AVSampleBufferAudioRendererClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSampleBufferAudioRendererClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSampleBufferAudioRendererClass) Alloc() AVSampleBufferAudioRenderer {
	rv := objc.Send[AVSampleBufferAudioRenderer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object used to decompress audio and play compressed or uncompressed
// audio.
//
// # Overview
// 
// You must add an instance of this class to an
// [AVSampleBufferRenderSynchronizer] before queuing the first sample buffer.
//
// # Determining rendering status
//
//   - [AVSampleBufferAudioRenderer.Status]: The status of the audio renderer.
//
// # Removing queued buffers
//
//   - [AVSampleBufferAudioRenderer.FlushFromSourceTimeCompletionHandler]: Flushes queued sample buffers with presentation time stamps later than or equal to the specified time.
//   - [AVSampleBufferAudioRenderer.AVSampleBufferAudioRendererFlushTimeKey]: The key that indicates the presentation timestamp of the first queued sample that was flushed.
//
// # Configuring time and pitch
//
//   - [AVSampleBufferAudioRenderer.AudioTimePitchAlgorithm]: The processing algorithm used to manage audio pitch at different rates.
//   - [AVSampleBufferAudioRenderer.SetAudioTimePitchAlgorithm]
//
// # Configuring audio spatialization
//
//   - [AVSampleBufferAudioRenderer.AllowedAudioSpatializationFormats]: The source audio channel layouts the audio renderer supports for spatialization.
//   - [AVSampleBufferAudioRenderer.SetAllowedAudioSpatializationFormats]
//
// # Managing audio output
//
//   - [AVSampleBufferAudioRenderer.Volume]: The current audio volume for the audio renderer.
//   - [AVSampleBufferAudioRenderer.SetVolume]
//   - [AVSampleBufferAudioRenderer.Muted]: A Boolean value that indicates whether audio for the renderer is in a muted state.
//   - [AVSampleBufferAudioRenderer.SetMuted]
//   - [AVSampleBufferAudioRenderer.AudioOutputDeviceUniqueID]: The unique identifier of the output device used to play audio.
//   - [AVSampleBufferAudioRenderer.SetAudioOutputDeviceUniqueID]
//
// # Responding to errors
//
//   - [AVSampleBufferAudioRenderer.Error]: The error that caused the renderer to no longer render sample buffers.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer
type AVSampleBufferAudioRenderer struct {
	objectivec.Object
}

// AVSampleBufferAudioRendererFromID constructs a [AVSampleBufferAudioRenderer] from an objc.ID.
//
// An object used to decompress audio and play compressed or uncompressed
// audio.
func AVSampleBufferAudioRendererFromID(id objc.ID) AVSampleBufferAudioRenderer {
	return AVSampleBufferAudioRenderer{objectivec.Object{ID: id}}
}
// NOTE: AVSampleBufferAudioRenderer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVSampleBufferAudioRenderer] class.
//
// # Determining rendering status
//
//   - [IAVSampleBufferAudioRenderer.Status]: The status of the audio renderer.
//
// # Removing queued buffers
//
//   - [IAVSampleBufferAudioRenderer.FlushFromSourceTimeCompletionHandler]: Flushes queued sample buffers with presentation time stamps later than or equal to the specified time.
//   - [IAVSampleBufferAudioRenderer.AVSampleBufferAudioRendererFlushTimeKey]: The key that indicates the presentation timestamp of the first queued sample that was flushed.
//
// # Configuring time and pitch
//
//   - [IAVSampleBufferAudioRenderer.AudioTimePitchAlgorithm]: The processing algorithm used to manage audio pitch at different rates.
//   - [IAVSampleBufferAudioRenderer.SetAudioTimePitchAlgorithm]
//
// # Configuring audio spatialization
//
//   - [IAVSampleBufferAudioRenderer.AllowedAudioSpatializationFormats]: The source audio channel layouts the audio renderer supports for spatialization.
//   - [IAVSampleBufferAudioRenderer.SetAllowedAudioSpatializationFormats]
//
// # Managing audio output
//
//   - [IAVSampleBufferAudioRenderer.Volume]: The current audio volume for the audio renderer.
//   - [IAVSampleBufferAudioRenderer.SetVolume]
//   - [IAVSampleBufferAudioRenderer.Muted]: A Boolean value that indicates whether audio for the renderer is in a muted state.
//   - [IAVSampleBufferAudioRenderer.SetMuted]
//   - [IAVSampleBufferAudioRenderer.AudioOutputDeviceUniqueID]: The unique identifier of the output device used to play audio.
//   - [IAVSampleBufferAudioRenderer.SetAudioOutputDeviceUniqueID]
//
// # Responding to errors
//
//   - [IAVSampleBufferAudioRenderer.Error]: The error that caused the renderer to no longer render sample buffers.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer
type IAVSampleBufferAudioRenderer interface {
	objectivec.IObject
	AVQueuedSampleBufferRendering

	// Topic: Determining rendering status

	// The status of the audio renderer.
	Status() AVQueuedSampleBufferRenderingStatus

	// Topic: Removing queued buffers

	// Flushes queued sample buffers with presentation time stamps later than or equal to the specified time.
	FlushFromSourceTimeCompletionHandler(time uintptr, completionHandler BoolHandler)
	// The key that indicates the presentation timestamp of the first queued sample that was flushed.
	AVSampleBufferAudioRendererFlushTimeKey() string

	// Topic: Configuring time and pitch

	// The processing algorithm used to manage audio pitch at different rates.
	AudioTimePitchAlgorithm() AVAudioTimePitchAlgorithm
	SetAudioTimePitchAlgorithm(value AVAudioTimePitchAlgorithm)

	// Topic: Configuring audio spatialization

	// The source audio channel layouts the audio renderer supports for spatialization.
	AllowedAudioSpatializationFormats() AVAudioSpatializationFormats
	SetAllowedAudioSpatializationFormats(value AVAudioSpatializationFormats)

	// Topic: Managing audio output

	// The current audio volume for the audio renderer.
	Volume() float32
	SetVolume(value float32)
	// A Boolean value that indicates whether audio for the renderer is in a muted state.
	Muted() bool
	SetMuted(value bool)
	// The unique identifier of the output device used to play audio.
	AudioOutputDeviceUniqueID() string
	SetAudioOutputDeviceUniqueID(value string)

	// Topic: Responding to errors

	// The error that caused the renderer to no longer render sample buffers.
	Error() foundation.INSError
}

// Init initializes the instance.
func (s AVSampleBufferAudioRenderer) Init() AVSampleBufferAudioRenderer {
	rv := objc.Send[AVSampleBufferAudioRenderer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSampleBufferAudioRenderer) Autorelease() AVSampleBufferAudioRenderer {
	rv := objc.Send[AVSampleBufferAudioRenderer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSampleBufferAudioRenderer creates a new AVSampleBufferAudioRenderer instance.
func NewAVSampleBufferAudioRenderer() AVSampleBufferAudioRenderer {
	class := getAVSampleBufferAudioRendererClass()
	rv := objc.Send[AVSampleBufferAudioRenderer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Flushes queued sample buffers with presentation time stamps later than or
// equal to the specified time.
//
// time: The time used to flush all later sample buffers.
//
// completionHandler: The block to invoke when the flush operation has either been completed or
// been interrupted. The block takes one argument:
// 
// flushSucceeded: A Boolean value indicating whether the sample buffers were
// flushed.
//
// # Discussion
// 
// This method can be used to replace media data scheduled to be rendered in
// the future, without interrupting playback. One example of this is when the
// data that has already been enqueued is from a sequence of two songs and the
// second song is swapped for a new song. In this case, this method would be
// called with the timestamp of the first sample buffer from the second song.
// After the completion handler is executed with a [YES] parameter, media data
// may again be enqueued with time stamps at the specified time.
// 
// If [NO] is provided to the completion handler, the flush did not succeed
// and the set of enqueued sample buffers remains unchanged. A flush can fail
// because the source time was too close to (or earlier than) the current time
// or because the current configuration of the receiver does not support
// flushing at a particular time. In these cases, the caller can choose to
// flush all enqueued media data by invoking [Flush].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/flush(fromSourceTime:completionHandler:)
func (s AVSampleBufferAudioRenderer) FlushFromSourceTimeCompletionHandler(time uintptr, completionHandler BoolHandler) {
_block1, _ := NewBoolBlock(completionHandler)
	objc.Send[objc.ID](s.ID, objc.Sel("flushFromSourceTime:completionHandler:"), time, _block1)
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
func (s AVSampleBufferAudioRenderer) EnqueueSampleBuffer(sampleBuffer uintptr) {
	objc.Send[objc.ID](s.ID, objc.Sel("enqueueSampleBuffer:"), sampleBuffer)
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
func (s AVSampleBufferAudioRenderer) Flush() {
	objc.Send[objc.ID](s.ID, objc.Sel("flush"))
}
// A Boolean value that indicates whether the receiver is able to accept more
// sample buffers.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/isReadyForMoreMediaData
func (s AVSampleBufferAudioRenderer) IsReadyForMoreMediaData() bool {
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
func (s AVSampleBufferAudioRenderer) RequestMediaDataWhenReadyOnQueueUsingBlock(queue dispatch.Queue, block VoidHandler) {
_block1, _ := NewVoidBlock(block)
	objc.Send[objc.ID](s.ID, objc.Sel("requestMediaDataWhenReadyOnQueue:usingBlock:"), uintptr(queue.Handle()), _block1)
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
func (s AVSampleBufferAudioRenderer) StopRequestingMediaData() {
	objc.Send[objc.ID](s.ID, objc.Sel("stopRequestingMediaData"))
}

// The status of the audio renderer.
//
// # Discussion
// 
// A renderer begins with a status of
// [QueuedSampleBufferRenderingStatusUnknown]. As you add sample buffers to
// the queue for rendering, the renderer transitions to either
// [QueuedSampleBufferRenderingStatusRendering] or
// [QueuedSampleBufferRenderingStatusFailed].
// 
// If the status is `AVQueuedSampleBufferRenderingStatus.Failed()`, check the
// value of the renderer’s error property for information on the error
// encountered. This property is key value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/status
func (s AVSampleBufferAudioRenderer) Status() AVQueuedSampleBufferRenderingStatus {
	rv := objc.Send[AVQueuedSampleBufferRenderingStatus](s.ID, objc.Sel("status"))
	return AVQueuedSampleBufferRenderingStatus(rv)
}
// The key that indicates the presentation timestamp of the first queued
// sample that was flushed.
//
// See: https://developer.apple.com/documentation/avfoundation/avsamplebufferaudiorendererflushtimekey
func (s AVSampleBufferAudioRenderer) AVSampleBufferAudioRendererFlushTimeKey() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("AVSampleBufferAudioRendererFlushTimeKey"))
	return foundation.NSStringFromID(rv).String()
}
// The processing algorithm used to manage audio pitch at different rates.
//
// # Discussion
// 
// The default value on iOS is [lowQualityZeroLatency]; on macOS, the default
// is [timeDomain]. The device automatically mutes audio when [Timebase] is
// not supported by [AVAudioTimePitchAlgorithm]. Modifying this property while
// [Timebase] is not `0.0` may cause the rate to briefly change to `0.0`.
//
// [lowQualityZeroLatency]: https://developer.apple.com/documentation/AVFoundation/AVAudioTimePitchAlgorithm/lowQualityZeroLatency
// [timeDomain]: https://developer.apple.com/documentation/AVFoundation/AVAudioTimePitchAlgorithm/timeDomain
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/audioTimePitchAlgorithm
func (s AVSampleBufferAudioRenderer) AudioTimePitchAlgorithm() AVAudioTimePitchAlgorithm {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("audioTimePitchAlgorithm"))
	return AVAudioTimePitchAlgorithm(foundation.NSStringFromID(rv).String())
}
func (s AVSampleBufferAudioRenderer) SetAudioTimePitchAlgorithm(value AVAudioTimePitchAlgorithm) {
	objc.Send[struct{}](s.ID, objc.Sel("setAudioTimePitchAlgorithm:"), objc.String(string(value)))
}
// The source audio channel layouts the audio renderer supports for
// spatialization.
//
// # Discussion
// 
// The default property value is [AudioSpatializationFormatMultichannel],
// which tells the player to spatialize any decodable multichannel layout.
// Setting the value to [AudioSpatializationFormatMonoStereoAndMultichannel]
// tells the player to spatialize any decodable mono, stereo, or multichannel
// layout. When this property value is
// [AudioSpatializationFormatMonoAndStereo] the player attempts to spatialize
// content tagged with a stereo channel layout (two-channel content with no
// layout specified as well as mono).
// 
// This property isn’t key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/allowedAudioSpatializationFormats
func (s AVSampleBufferAudioRenderer) AllowedAudioSpatializationFormats() AVAudioSpatializationFormats {
	rv := objc.Send[AVAudioSpatializationFormats](s.ID, objc.Sel("allowedAudioSpatializationFormats"))
	return AVAudioSpatializationFormats(rv)
}
func (s AVSampleBufferAudioRenderer) SetAllowedAudioSpatializationFormats(value AVAudioSpatializationFormats) {
	objc.Send[struct{}](s.ID, objc.Sel("setAllowedAudioSpatializationFormats:"), value)
}
// The current audio volume for the audio renderer.
//
// # Discussion
// 
// Use this property for frequent vloume changes; for example, a volume knob
// or fader. A value of `0.0` silences all audio while a value of `1.0` plays
// all audio at full volume.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/volume
func (s AVSampleBufferAudioRenderer) Volume() float32 {
	rv := objc.Send[float32](s.ID, objc.Sel("volume"))
	return rv
}
func (s AVSampleBufferAudioRenderer) SetVolume(value float32) {
	objc.Send[struct{}](s.ID, objc.Sel("setVolume:"), value)
}
// A Boolean value that indicates whether audio for the renderer is in a muted
// state.
//
// # Discussion
// 
// This property only affects muting the renderer instance and not the device.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/isMuted
func (s AVSampleBufferAudioRenderer) Muted() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isMuted"))
	return rv
}
func (s AVSampleBufferAudioRenderer) SetMuted(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setMuted:"), value)
}
// The unique identifier of the output device used to play audio.
//
// # Discussion
// 
// The default value of this property is `nil`, which indicates the use of the
// default audio device. Otherwise, set the value to an [NSString] containing
// the unique identifier of the Core Audio output device to use for audio
// output. [kAudioDevicePropertyDeviceUID] is a suitable source of audio
// output device unique IDs.
// 
// Modifying this property while the timebase’s rate isn’t `0.0` may cause
// the rate to briefly change to `0.0`.
// 
// On macOS, you can use the audio device clock as the
// [AVSampleBufferRenderSynchronizer] and all attached
// [AVQueuedSampleBufferRendering] timebase clocks. If you modify the
// `audioOutputDeviceUniqueID`, the clocks of all these timebases may also
// change.
// 
// If you attach multiple renderers with different values for
// `audioOutputDeviceUniqueID` to the same buffer renderer synchronizer, audio
// may not stay in sync during playback. To avoid this, ensure that all
// synchronized sample buffer renderers are using the same audio output
// device.
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [kAudioDevicePropertyDeviceUID]: https://developer.apple.com/documentation/CoreAudio/kAudioDevicePropertyDeviceUID
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/audioOutputDeviceUniqueID
func (s AVSampleBufferAudioRenderer) AudioOutputDeviceUniqueID() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("audioOutputDeviceUniqueID"))
	return foundation.NSStringFromID(rv).String()
}
func (s AVSampleBufferAudioRenderer) SetAudioOutputDeviceUniqueID(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setAudioOutputDeviceUniqueID:"), objc.String(value))
}
// The error that caused the renderer to no longer render sample buffers.
//
// # Discussion
// 
// The value of this property is nil unless the value of [Status] is
// [QueuedSampleBufferRenderingStatusFailed].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAudioRenderer/error
func (s AVSampleBufferAudioRenderer) Error() foundation.INSError {
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
func (s AVSampleBufferAudioRenderer) HasSufficientMediaDataForReliablePlaybackStart() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("hasSufficientMediaDataForReliablePlaybackStart"))
	return rv
}
// The timebase for a renderer.
//
// # Discussion
// 
// The timebase governs how time stamps are interpreted by the renderer.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVQueuedSampleBufferRendering/timebase
func (s AVSampleBufferAudioRenderer) Timebase() uintptr {
	rv := objc.Send[uintptr](s.ID, objc.Sel("timebase"))
	return rv
}

			// Protocol methods for AVQueuedSampleBufferRendering
			

// FlushFromSourceTime is a synchronous wrapper around [AVSampleBufferAudioRenderer.FlushFromSourceTimeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s AVSampleBufferAudioRenderer) FlushFromSourceTime(ctx context.Context, time uintptr) (bool, error) {
	done := make(chan bool, 1)
	s.FlushFromSourceTimeCompletionHandler(time, func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// RequestMediaDataWhenReadyOnQueueUsingBlockSync is a synchronous wrapper around [AVSampleBufferAudioRenderer.RequestMediaDataWhenReadyOnQueueUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (s AVSampleBufferAudioRenderer) RequestMediaDataWhenReadyOnQueueUsingBlockSync(ctx context.Context, queue dispatch.Queue) error {
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

