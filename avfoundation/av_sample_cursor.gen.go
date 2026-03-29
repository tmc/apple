// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVSampleCursor] class.
var (
	_AVSampleCursorClass     AVSampleCursorClass
	_AVSampleCursorClassOnce sync.Once
)

func getAVSampleCursorClass() AVSampleCursorClass {
	_AVSampleCursorClassOnce.Do(func() {
		_AVSampleCursorClass = AVSampleCursorClass{class: objc.GetClass("AVSampleCursor")}
	})
	return _AVSampleCursorClass
}

// GetAVSampleCursorClass returns the class object for AVSampleCursor.
func GetAVSampleCursorClass() AVSampleCursorClass {
	return getAVSampleCursorClass()
}

type AVSampleCursorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVSampleCursorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVSampleCursorClass) Alloc() AVSampleCursor {
	rv := objc.Send[AVSampleCursor](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that provides information about the media sample at the
// cursor’s current position.
//
// # Overview
// 
// You position a sample cursor at a specific media sample in a sequence of
// samples contained in a higher-level object, like an [AVAssetTrack]. You can
// move it to a new position in that sequence either backwards or forwards,
// either in decode order or in presentation order. You can also request
// moving it according to a count of samples or a delta in time.
// 
// Use a sample cursor to get information about the media sample such as its
// duration, timestamps, dependency information, and so on. You can also use
// them to synchronously to perform I/O in order to load media data of one or
// more media samples into memory.
//
// # Navigating samples
//
//   - [AVSampleCursor.StepByDecodeTimeWasPinned]: Moves the cursor by a given delta time on the decode timeline.
//   - [AVSampleCursor.StepByPresentationTimeWasPinned]: Moves the cursor by a given delta time on the presentation timeline.
//   - [AVSampleCursor.StepInDecodeOrderByCount]: Moves the cursor a given number of samples in decode order.
//   - [AVSampleCursor.StepInPresentationOrderByCount]: Moves the cursor a given number of samples in presentation order.
//
// # Getting timestamps
//
//   - [AVSampleCursor.DecodeTimeStamp]: The decode timestamp of the sample at the current position of the cursor.
//   - [AVSampleCursor.PresentationTimeStamp]: The presentation timestamp of the sample at the current position of the cursor.
//
// # Getting sample information
//
//   - [AVSampleCursor.CurrentChunkInfo]: A value that provides information about the chunk of samples to which the current sample belongs.
//   - [AVSampleCursor.CurrentChunkStorageRange]: The sample range in the storage container to load together with the current sample as a chunk.
//   - [AVSampleCursor.CurrentChunkStorageURL]: The URL of the storage container of the current sample and other samples to load in the same operation as a chunk.
//   - [AVSampleCursor.CurrentSampleDependencyInfo]: The dependency information that describes relationships between a media sample and other media samples in the same sample sequence.
//   - [AVSampleCursor.CurrentSampleDuration]: The decode duration of the sample at the cursor’s current position.
//   - [AVSampleCursor.CurrentSampleIndexInChunk]: The index of the current sample within the chunk to which it belongs.
//   - [AVSampleCursor.CurrentSampleStorageRange]: The offset and length of the current sample in the current chunk storage URL.
//   - [AVSampleCursor.CurrentSampleSyncInfo]: The synchronization information for the current sample for consideration when resynchronizing a decoder.
//   - [AVSampleCursor.CopyCurrentSampleFormatDescription]: Returns the format description of the sample at the cursor’s current position.
//   - [AVSampleCursor.CurrentSampleAudioDependencyInfo]: The independent decodability information for the audio sample.
//   - [AVSampleCursor.CurrentSampleDependencyAttachments]: A dictionary of dependency-related sample buffer attachments.
//
// # Accessing samples
//
//   - [AVSampleCursor.SamplesWithEarlierDecodeTimeStampsMayHaveLaterPresentationTimeStampsThanCursor]: Determines whether a sample earlier in decode order can have a presentation timestamp later than that of the specified sample cursor.
//   - [AVSampleCursor.SamplesWithLaterDecodeTimeStampsMayHaveEarlierPresentationTimeStampsThanCursor]: Determines whether a sample later in decode order can have a presentation timestamp earlier than that of the specified sample cursor.
//   - [AVSampleCursor.SamplesRequiredForDecoderRefresh]: The number of samples prior to the current sample, in decode order, the decoder requires to achieve a coherent output at the current decode time.
//
// # Comparing sample cursors
//
//   - [AVSampleCursor.ComparePositionInDecodeOrderWithPositionOfCursor]: Compares the relative positions of two sample cursors and returns their relative positions.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor
type AVSampleCursor struct {
	objectivec.Object
}

// AVSampleCursorFromID constructs a [AVSampleCursor] from an objc.ID.
//
// An object that provides information about the media sample at the
// cursor’s current position.
func AVSampleCursorFromID(id objc.ID) AVSampleCursor {
	return AVSampleCursor{objectivec.Object{ID: id}}
}
// NOTE: AVSampleCursor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVSampleCursor] class.
//
// # Navigating samples
//
//   - [IAVSampleCursor.StepByDecodeTimeWasPinned]: Moves the cursor by a given delta time on the decode timeline.
//   - [IAVSampleCursor.StepByPresentationTimeWasPinned]: Moves the cursor by a given delta time on the presentation timeline.
//   - [IAVSampleCursor.StepInDecodeOrderByCount]: Moves the cursor a given number of samples in decode order.
//   - [IAVSampleCursor.StepInPresentationOrderByCount]: Moves the cursor a given number of samples in presentation order.
//
// # Getting timestamps
//
//   - [IAVSampleCursor.DecodeTimeStamp]: The decode timestamp of the sample at the current position of the cursor.
//   - [IAVSampleCursor.PresentationTimeStamp]: The presentation timestamp of the sample at the current position of the cursor.
//
// # Getting sample information
//
//   - [IAVSampleCursor.CurrentChunkInfo]: A value that provides information about the chunk of samples to which the current sample belongs.
//   - [IAVSampleCursor.CurrentChunkStorageRange]: The sample range in the storage container to load together with the current sample as a chunk.
//   - [IAVSampleCursor.CurrentChunkStorageURL]: The URL of the storage container of the current sample and other samples to load in the same operation as a chunk.
//   - [IAVSampleCursor.CurrentSampleDependencyInfo]: The dependency information that describes relationships between a media sample and other media samples in the same sample sequence.
//   - [IAVSampleCursor.CurrentSampleDuration]: The decode duration of the sample at the cursor’s current position.
//   - [IAVSampleCursor.CurrentSampleIndexInChunk]: The index of the current sample within the chunk to which it belongs.
//   - [IAVSampleCursor.CurrentSampleStorageRange]: The offset and length of the current sample in the current chunk storage URL.
//   - [IAVSampleCursor.CurrentSampleSyncInfo]: The synchronization information for the current sample for consideration when resynchronizing a decoder.
//   - [IAVSampleCursor.CopyCurrentSampleFormatDescription]: Returns the format description of the sample at the cursor’s current position.
//   - [IAVSampleCursor.CurrentSampleAudioDependencyInfo]: The independent decodability information for the audio sample.
//   - [IAVSampleCursor.CurrentSampleDependencyAttachments]: A dictionary of dependency-related sample buffer attachments.
//
// # Accessing samples
//
//   - [IAVSampleCursor.SamplesWithEarlierDecodeTimeStampsMayHaveLaterPresentationTimeStampsThanCursor]: Determines whether a sample earlier in decode order can have a presentation timestamp later than that of the specified sample cursor.
//   - [IAVSampleCursor.SamplesWithLaterDecodeTimeStampsMayHaveEarlierPresentationTimeStampsThanCursor]: Determines whether a sample later in decode order can have a presentation timestamp earlier than that of the specified sample cursor.
//   - [IAVSampleCursor.SamplesRequiredForDecoderRefresh]: The number of samples prior to the current sample, in decode order, the decoder requires to achieve a coherent output at the current decode time.
//
// # Comparing sample cursors
//
//   - [IAVSampleCursor.ComparePositionInDecodeOrderWithPositionOfCursor]: Compares the relative positions of two sample cursors and returns their relative positions.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor
type IAVSampleCursor interface {
	objectivec.IObject

	// Topic: Navigating samples

	// Moves the cursor by a given delta time on the decode timeline.
	StepByDecodeTimeWasPinned(deltaDecodeTime coremedia.CMTime, outWasPinned unsafe.Pointer) coremedia.CMTime
	// Moves the cursor by a given delta time on the presentation timeline.
	StepByPresentationTimeWasPinned(deltaPresentationTime coremedia.CMTime, outWasPinned unsafe.Pointer) coremedia.CMTime
	// Moves the cursor a given number of samples in decode order.
	StepInDecodeOrderByCount(stepCount int64) int64
	// Moves the cursor a given number of samples in presentation order.
	StepInPresentationOrderByCount(stepCount int64) int64

	// Topic: Getting timestamps

	// The decode timestamp of the sample at the current position of the cursor.
	DecodeTimeStamp() coremedia.CMTime
	// The presentation timestamp of the sample at the current position of the cursor.
	PresentationTimeStamp() coremedia.CMTime

	// Topic: Getting sample information

	// A value that provides information about the chunk of samples to which the current sample belongs.
	CurrentChunkInfo() AVSampleCursorChunkInfo
	// The sample range in the storage container to load together with the current sample as a chunk.
	CurrentChunkStorageRange() AVSampleCursorStorageRange
	// The URL of the storage container of the current sample and other samples to load in the same operation as a chunk.
	CurrentChunkStorageURL() foundation.INSURL
	// The dependency information that describes relationships between a media sample and other media samples in the same sample sequence.
	CurrentSampleDependencyInfo() AVSampleCursorDependencyInfo
	// The decode duration of the sample at the cursor’s current position.
	CurrentSampleDuration() coremedia.CMTime
	// The index of the current sample within the chunk to which it belongs.
	CurrentSampleIndexInChunk() int64
	// The offset and length of the current sample in the current chunk storage URL.
	CurrentSampleStorageRange() AVSampleCursorStorageRange
	// The synchronization information for the current sample for consideration when resynchronizing a decoder.
	CurrentSampleSyncInfo() AVSampleCursorSyncInfo
	// Returns the format description of the sample at the cursor’s current position.
	CopyCurrentSampleFormatDescription() uintptr
	// The independent decodability information for the audio sample.
	CurrentSampleAudioDependencyInfo() AVSampleCursorAudioDependencyInfo
	// A dictionary of dependency-related sample buffer attachments.
	CurrentSampleDependencyAttachments() foundation.INSDictionary

	// Topic: Accessing samples

	// Determines whether a sample earlier in decode order can have a presentation timestamp later than that of the specified sample cursor.
	SamplesWithEarlierDecodeTimeStampsMayHaveLaterPresentationTimeStampsThanCursor(cursor IAVSampleCursor) bool
	// Determines whether a sample later in decode order can have a presentation timestamp earlier than that of the specified sample cursor.
	SamplesWithLaterDecodeTimeStampsMayHaveEarlierPresentationTimeStampsThanCursor(cursor IAVSampleCursor) bool
	// The number of samples prior to the current sample, in decode order, the decoder requires to achieve a coherent output at the current decode time.
	SamplesRequiredForDecoderRefresh() int

	// Topic: Comparing sample cursors

	// Compares the relative positions of two sample cursors and returns their relative positions.
	ComparePositionInDecodeOrderWithPositionOfCursor(cursor IAVSampleCursor) foundation.NSComparisonResult
}

// Init initializes the instance.
func (s AVSampleCursor) Init() AVSampleCursor {
	rv := objc.Send[AVSampleCursor](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s AVSampleCursor) Autorelease() AVSampleCursor {
	rv := objc.Send[AVSampleCursor](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVSampleCursor creates a new AVSampleCursor instance.
func NewAVSampleCursor() AVSampleCursor {
	class := getAVSampleCursorClass()
	rv := objc.Send[AVSampleCursor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Moves the cursor by a given delta time on the decode timeline.
//
// deltaDecodeTime: The amount of time to move in the decode timeline.
//
// outWasPinned: The system sets the value of this pointer to [true] if the cursor reaches
// the beginning or the end of the sample sequence before it reaches the
// requested time. You may specify `nil` if you’re not interested in this
// information.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The amount of time the cursor was moved along the decode timeline. Because
// sample cursors snap to sample boundaries when stepped, this value may not
// be equal to the specified time delta even if the cursor wasn’t pinned.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/step(byDecodeTime:wasPinned:)
func (s AVSampleCursor) StepByDecodeTimeWasPinned(deltaDecodeTime coremedia.CMTime, outWasPinned unsafe.Pointer) coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](s.ID, objc.Sel("stepByDecodeTime:wasPinned:"), deltaDecodeTime, outWasPinned)
	return coremedia.CMTime(rv)
}
// Moves the cursor by a given delta time on the presentation timeline.
//
// deltaPresentationTime: The amount of time to move in the presentation timeline.
//
// outWasPinned: The system sets the value of this pointer to [true] if the cursor reaches
// the beginning or the end of the sample sequence before it reaches the
// requested time. You may specify `nil` if you’re not interested in this
// information.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The amount of time the cursor was moved along the presentation timeline.
// Because sample cursors snap to sample boundaries when stepped, this value
// may not be equal to the specified time delta even if the cursor was not
// pinned.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/step(byPresentationTime:wasPinned:)
func (s AVSampleCursor) StepByPresentationTimeWasPinned(deltaPresentationTime coremedia.CMTime, outWasPinned unsafe.Pointer) coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](s.ID, objc.Sel("stepByPresentationTime:wasPinned:"), deltaPresentationTime, outWasPinned)
	return coremedia.CMTime(rv)
}
// Moves the cursor a given number of samples in decode order.
//
// stepCount: The number of samples to move across. If positive, step forward this many
// samples. If negative, step backward this many samples.
//
// # Return Value
// 
// The number of samples the cursor traversed. If the cursor reaches the
// beginning or the end of the sample sequence before the requested number of
// samples was traversed, the absolute value of the result will be less than
// the absolute value of the specified step count.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/stepInDecodeOrder(byCount:)
func (s AVSampleCursor) StepInDecodeOrderByCount(stepCount int64) int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("stepInDecodeOrderByCount:"), stepCount)
	return rv
}
// Moves the cursor a given number of samples in presentation order.
//
// stepCount: The number of samples to move across. If positive, step forward this many
// samples. If negative, step backward this many samples.
//
// # Return Value
// 
// The number of samples the cursor traversed. If the cursor reaches the
// beginning or the end of the sample sequence before the requested number of
// samples was traversed, the absolute value of the result will be less than
// the absolute value of the specified step count.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/stepInPresentationOrder(byCount:)
func (s AVSampleCursor) StepInPresentationOrderByCount(stepCount int64) int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("stepInPresentationOrderByCount:"), stepCount)
	return rv
}
// Returns the format description of the sample at the cursor’s current
// position.
//
// # Return Value
// 
// The current sample’s format description.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/copyCurrentSampleFormatDescription()
func (s AVSampleCursor) CopyCurrentSampleFormatDescription() uintptr {
	rv := objc.Send[uintptr](s.ID, objc.Sel("copyCurrentSampleFormatDescription"))
	return rv
}
// Determines whether a sample earlier in decode order can have a presentation
// timestamp later than that of the specified sample cursor.
//
// cursor: An instance of [AVSampleCursor] with which to test the sample reordering
// boundary.
//
// # Return Value
// 
// [true] if it’s possible for any sample earlier in decode order than the
// sample at the position of the receiver can have a presentation timestamp
// later than that of the specified sample cursor; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Undefined results occur if this cursor and the passed in cursor reference
// different sequences of samples, such as when they’re created by different
// instances of [AVAssetTrack].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/maySamplesWithEarlierDecodeTimeStampsHavePresentationTimeStamps(laterThan:)
func (s AVSampleCursor) SamplesWithEarlierDecodeTimeStampsMayHaveLaterPresentationTimeStampsThanCursor(cursor IAVSampleCursor) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("samplesWithEarlierDecodeTimeStampsMayHaveLaterPresentationTimeStampsThanCursor:"), cursor)
	return rv
}
// Determines whether a sample later in decode order can have a presentation
// timestamp earlier than that of the specified sample cursor.
//
// cursor: An instance of [AVSampleCursor] with which to test the sample reordering
// boundary.
//
// # Return Value
// 
// [true] if it’s possible for any sample later in decode order than the
// sample at the position of the receiver can have a presentation timestamp
// earlier than that of the specified sample cursor; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Undefined results occur if this cursor and the passed in cursor reference
// different sequences of samples, such as when they’re created by different
// instances of [AVAssetTrack].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/maySamplesWithLaterDecodeTimeStampsHavePresentationTimeStamps(earlierThan:)
func (s AVSampleCursor) SamplesWithLaterDecodeTimeStampsMayHaveEarlierPresentationTimeStampsThanCursor(cursor IAVSampleCursor) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("samplesWithLaterDecodeTimeStampsMayHaveEarlierPresentationTimeStampsThanCursor:"), cursor)
	return rv
}
// Compares the relative positions of two sample cursors and returns their
// relative positions.
//
// cursor: An instance of [AVSampleCursor] with which to compare positions.
//
// # Return Value
// 
// Returns a comparison result that indicates of this cursor points at a
// sample before, the same as, or after the sample pointed to by the specified
// cursor.
//
// # Discussion
// 
// Undefined results occur if this cursor and the passed in cursor reference
// different sequences of samples, such as when they’re created by different
// instances of [AVAssetTrack].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/comparePositionInDecodeOrder(withPositionOf:)
func (s AVSampleCursor) ComparePositionInDecodeOrderWithPositionOfCursor(cursor IAVSampleCursor) foundation.NSComparisonResult {
	rv := objc.Send[foundation.NSComparisonResult](s.ID, objc.Sel("comparePositionInDecodeOrderWithPositionOfCursor:"), cursor)
	return foundation.NSComparisonResult(rv)
}

// The decode timestamp of the sample at the current position of the cursor.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/decodeTimeStamp
func (s AVSampleCursor) DecodeTimeStamp() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](s.ID, objc.Sel("decodeTimeStamp"))
	return coremedia.CMTime(rv)
}
// The presentation timestamp of the sample at the current position of the
// cursor.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/presentationTimeStamp
func (s AVSampleCursor) PresentationTimeStamp() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](s.ID, objc.Sel("presentationTimeStamp"))
	return coremedia.CMTime(rv)
}
// A value that provides information about the chunk of samples to which the
// current sample belongs.
//
// # Discussion
// 
// If the media format that defines the sequence of samples doesn’t indicate
// chunking of samples, the cursor considers each sample to belong to a chunk
// of one sample only.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/currentChunkInfo
func (s AVSampleCursor) CurrentChunkInfo() AVSampleCursorChunkInfo {
	rv := objc.Send[AVSampleCursorChunkInfo](s.ID, objc.Sel("currentChunkInfo"))
	return AVSampleCursorChunkInfo(rv)
}
// The sample range in the storage container to load together with the current
// sample as a chunk.
//
// # Discussion
// 
// If the current chunk isn’t stored contiguously in its storage container,
// the property value’s [offset] is `-1`. In such cases use
// [AVSampleBufferGenerator] to obtain the sample data.
//
// [offset]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursorStorageRange/offset
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/currentChunkStorageRange
func (s AVSampleCursor) CurrentChunkStorageRange() AVSampleCursorStorageRange {
	rv := objc.Send[AVSampleCursorStorageRange](s.ID, objc.Sel("currentChunkStorageRange"))
	return AVSampleCursorStorageRange(rv)
}
// The URL of the storage container of the current sample and other samples to
// load in the same operation as a chunk.
//
// # Discussion
// 
// When this property is `nil`, the storage location of the chunk is the URL
// of the sample cursor’s track’s asset, if it has one.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/currentChunkStorageURL
func (s AVSampleCursor) CurrentChunkStorageURL() foundation.INSURL {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("currentChunkStorageURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// The dependency information that describes relationships between a media
// sample and other media samples in the same sample sequence.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/currentSampleDependencyInfo
func (s AVSampleCursor) CurrentSampleDependencyInfo() AVSampleCursorDependencyInfo {
	rv := objc.Send[AVSampleCursorDependencyInfo](s.ID, objc.Sel("currentSampleDependencyInfo"))
	return AVSampleCursorDependencyInfo(rv)
}
// The decode duration of the sample at the cursor’s current position.
//
// # Discussion
// 
// If the cursor needs to move past its current position to determine the
// decode duration of the current sample, the value of this property is
// [indefinite]. This condition can occur with streaming formats such as
// MPEG-2 transport streams.
//
// [indefinite]: https://developer.apple.com/documentation/CoreMedia/CMTime/indefinite
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/currentSampleDuration
func (s AVSampleCursor) CurrentSampleDuration() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](s.ID, objc.Sel("currentSampleDuration"))
	return coremedia.CMTime(rv)
}
// The index of the current sample within the chunk to which it belongs.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/currentSampleIndexInChunk
func (s AVSampleCursor) CurrentSampleIndexInChunk() int64 {
	rv := objc.Send[int64](s.ID, objc.Sel("currentSampleIndexInChunk"))
	return rv
}
// The offset and length of the current sample in the current chunk storage
// URL.
//
// # Discussion
// 
// If the current chunk isn’t stored contiguously in its storage container,
// the property value’s [offset] is `-1`. In such cases use
// [AVSampleBufferGenerator] to obtain the sample data.
//
// [offset]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursorStorageRange/offset
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/currentSampleStorageRange
func (s AVSampleCursor) CurrentSampleStorageRange() AVSampleCursorStorageRange {
	rv := objc.Send[AVSampleCursorStorageRange](s.ID, objc.Sel("currentSampleStorageRange"))
	return AVSampleCursorStorageRange(rv)
}
// The synchronization information for the current sample for consideration
// when resynchronizing a decoder.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/currentSampleSyncInfo
func (s AVSampleCursor) CurrentSampleSyncInfo() AVSampleCursorSyncInfo {
	rv := objc.Send[AVSampleCursorSyncInfo](s.ID, objc.Sel("currentSampleSyncInfo"))
	return AVSampleCursorSyncInfo(rv)
}
// The independent decodability information for the audio sample.
//
// # Discussion
// 
// To position a sample cursor at the first sample that the audio decoder
// requires for a full refresh, move it back from the current sample until you
// find a sample that meets the following criteria:
// 
// - The value of its [audioSampleIsIndependentlyDecodable] property is
// [true]. - The value of its [audioSamplePacketRefreshCount] property is
// greater than or equal to the number of steps back you’ve taken.
// 
// You don’t need to reposition the cursorif the current sample is
// independently decodable with an [audioSamplePacketRefreshCount] of `0`.
//
// [audioSampleIsIndependentlyDecodable]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursorAudioDependencyInfo/audioSampleIsIndependentlyDecodable
// [audioSamplePacketRefreshCount]: https://developer.apple.com/documentation/AVFoundation/AVSampleCursorAudioDependencyInfo/audioSamplePacketRefreshCount
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/currentSampleAudioDependencyInfo
func (s AVSampleCursor) CurrentSampleAudioDependencyInfo() AVSampleCursorAudioDependencyInfo {
	rv := objc.Send[AVSampleCursorAudioDependencyInfo](s.ID, objc.Sel("currentSampleAudioDependencyInfo"))
	return AVSampleCursorAudioDependencyInfo(rv)
}
// A dictionary of dependency-related sample buffer attachments.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/currentSampleDependencyAttachments
func (s AVSampleCursor) CurrentSampleDependencyAttachments() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("currentSampleDependencyAttachments"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// The number of samples prior to the current sample, in decode order, the
// decoder requires to achieve a coherent output at the current decode time.
//
// # Discussion
// 
// This property value is 0 when the decoder doesn’t require samples for
// refresh or when the track doesn’t contain this information.
// 
// Some sample sequences don’t indicate sample dependencies and instead
// indicate to decode a specific sample with all available accuracy. The
// system must decode samples in decode order before decoding the specific
// sample.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleCursor/samplesRequiredForDecoderRefresh
func (s AVSampleCursor) SamplesRequiredForDecoderRefresh() int {
	rv := objc.Send[int](s.ID, objc.Sel("samplesRequiredForDecoderRefresh"))
	return rv
}

