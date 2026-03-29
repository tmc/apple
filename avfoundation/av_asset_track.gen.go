// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetTrack] class.
var (
	_AVAssetTrackClass     AVAssetTrackClass
	_AVAssetTrackClassOnce sync.Once
)

func getAVAssetTrackClass() AVAssetTrackClass {
	_AVAssetTrackClassOnce.Do(func() {
		_AVAssetTrackClass = AVAssetTrackClass{class: objc.GetClass("AVAssetTrack")}
	})
	return _AVAssetTrackClass
}

// GetAVAssetTrackClass returns the class object for AVAssetTrack.
func GetAVAssetTrackClass() AVAssetTrackClass {
	return getAVAssetTrackClass()
}

type AVAssetTrackClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetTrackClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetTrackClass) Alloc() AVAssetTrack {
	rv := objc.Send[AVAssetTrack](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that models a track of media that an asset contains.
//
// # Overview
// 
// An asset contains one or more tracks of media that the framework models
// using the [AVAssetTrack] class. A track object holds the uniformly typed
// media that an asset provides such as audio, video, or closed captions.
// 
// A track, like its containing [AVAsset], doesn’t load all of its media
// upon creation. Instead, it defers loading its data until you perform an
// operation that requires it. Because loading the data can take time, an
// asset track adopts the [AVAsynchronousKeyValueLoading] protocol so you can
// load its property values asynchronously by calling the [load(_:isolation:)]
// method.
//
// [load(_:isolation:)]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousKeyValueLoading/load(_:isolation:)
//
// # Identifying an asset track
//
//   - [AVAssetTrack.TrackID]: The persistent unique identifier for this track.
//   - [AVAssetTrack.MediaType]: The type of media that a track presents.
//   - [AVAssetTrack.Asset]: The asset object that contains this track.
//
// # Loading track information
//
//   - [AVAssetTrack.MediaCharacteristics]: The media characteristics for the track.
//   - [AVAssetTrack.SetMediaCharacteristics]
//
// # Loading track segments
//
//   - [AVAssetTrack.LoadSegmentForTrackTimeCompletionHandler]: Loads a segment with a target time range that contains, or is closest to, the specified track time.
//   - [AVAssetTrack.LoadSamplePresentationTimeForTrackTimeCompletionHandler]: Loads a sample presentation time that maps to the specified track time.
//
// # Creating sample cursors
//
//   - [AVAssetTrack.MakeSampleCursorWithPresentationTimeStamp]: Creates a sample cursor and positions it at or near the specified presentation timestamp.
//   - [AVAssetTrack.MakeSampleCursorAtFirstSampleInDecodeOrder]: Creates a sample cursor and positions it at the track’s first media sample in decode order.
//   - [AVAssetTrack.MakeSampleCursorAtLastSampleInDecodeOrder]: Creates a sample cursor and positions it at the track’s last media sample in decode order.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack
type AVAssetTrack struct {
	objectivec.Object
}

// AVAssetTrackFromID constructs a [AVAssetTrack] from an objc.ID.
//
// An object that models a track of media that an asset contains.
func AVAssetTrackFromID(id objc.ID) AVAssetTrack {
	return AVAssetTrack{objectivec.Object{ID: id}}
}
// NOTE: AVAssetTrack adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetTrack] class.
//
// # Identifying an asset track
//
//   - [IAVAssetTrack.TrackID]: The persistent unique identifier for this track.
//   - [IAVAssetTrack.MediaType]: The type of media that a track presents.
//   - [IAVAssetTrack.Asset]: The asset object that contains this track.
//
// # Loading track information
//
//   - [IAVAssetTrack.MediaCharacteristics]: The media characteristics for the track.
//   - [IAVAssetTrack.SetMediaCharacteristics]
//
// # Loading track segments
//
//   - [IAVAssetTrack.LoadSegmentForTrackTimeCompletionHandler]: Loads a segment with a target time range that contains, or is closest to, the specified track time.
//   - [IAVAssetTrack.LoadSamplePresentationTimeForTrackTimeCompletionHandler]: Loads a sample presentation time that maps to the specified track time.
//
// # Creating sample cursors
//
//   - [IAVAssetTrack.MakeSampleCursorWithPresentationTimeStamp]: Creates a sample cursor and positions it at or near the specified presentation timestamp.
//   - [IAVAssetTrack.MakeSampleCursorAtFirstSampleInDecodeOrder]: Creates a sample cursor and positions it at the track’s first media sample in decode order.
//   - [IAVAssetTrack.MakeSampleCursorAtLastSampleInDecodeOrder]: Creates a sample cursor and positions it at the track’s last media sample in decode order.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack
type IAVAssetTrack interface {
	objectivec.IObject
	AVAsynchronousKeyValueLoading

	// Topic: Identifying an asset track

	// The persistent unique identifier for this track.
	TrackID() int32
	// The type of media that a track presents.
	MediaType() AVMediaType
	// The asset object that contains this track.
	Asset() IAVAsset

	// Topic: Loading track information

	// The media characteristics for the track.
	MediaCharacteristics() AVMediaCharacteristic
	SetMediaCharacteristics(value AVMediaCharacteristic)

	// Topic: Loading track segments

	// Loads a segment with a target time range that contains, or is closest to, the specified track time.
	LoadSegmentForTrackTimeCompletionHandler(trackTime uintptr, completionHandler AVAssetTrackSegmentErrorHandler)
	// Loads a sample presentation time that maps to the specified track time.
	LoadSamplePresentationTimeForTrackTimeCompletionHandler(trackTime uintptr, completionHandler CMTimeErrorHandler)

	// Topic: Creating sample cursors

	// Creates a sample cursor and positions it at or near the specified presentation timestamp.
	MakeSampleCursorWithPresentationTimeStamp(presentationTimeStamp uintptr) IAVSampleCursor
	// Creates a sample cursor and positions it at the track’s first media sample in decode order.
	MakeSampleCursorAtFirstSampleInDecodeOrder() IAVSampleCursor
	// Creates a sample cursor and positions it at the track’s last media sample in decode order.
	MakeSampleCursorAtLastSampleInDecodeOrder() IAVSampleCursor
}

// Init initializes the instance.
func (a AVAssetTrack) Init() AVAssetTrack {
	rv := objc.Send[AVAssetTrack](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetTrack) Autorelease() AVAssetTrack {
	rv := objc.Send[AVAssetTrack](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetTrack creates a new AVAssetTrack instance.
func NewAVAssetTrack() AVAssetTrack {
	class := getAVAssetTrackClass()
	rv := objc.Send[AVAssetTrack](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Loads a segment with a target time range that contains, or is closest to,
// the specified track time.
//
// trackTime: The track time of the segment to load.
//
// completionHandler: A callback that the system invokes after it finishes the loading request.
// It passes the completion handler the following parameters:
// 
// segment: The loaded track segment, or `nil` if an error occurs. error: An
// error object if the request fails; otherwise, `nil`.
//
// # Discussion
// 
// If the specified track time doesn’t map to a sample presentation time,
// the system returns the segment with the closest matching time.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/loadSegment(forTrackTime:completionHandler:)
func (a AVAssetTrack) LoadSegmentForTrackTimeCompletionHandler(trackTime uintptr, completionHandler AVAssetTrackSegmentErrorHandler) {
_block1, _ := NewAVAssetTrackSegmentErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("loadSegmentForTrackTime:completionHandler:"), trackTime, _block1)
}
// Loads a sample presentation time that maps to the specified track time.
//
// trackTime: The track time of the presentation time to load.
//
// completionHandler: A callback that the system invokes after it finishes the loading request.
// It passes the completion handler the following parameters:
// 
// time: A [CMTime] value, which is [invalid] if the track time is out of
// range or if an error occurs. error: An error object if the request fails;
// otherwise, `nil`.
// //
// [CMTime]: https://developer.apple.com/documentation/CoreMedia/CMTime
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/loadSamplePresentationTime(forTrackTime:completionHandler:)
func (a AVAssetTrack) LoadSamplePresentationTimeForTrackTimeCompletionHandler(trackTime uintptr, completionHandler CMTimeErrorHandler) {
_block1, _ := NewCMTimeErrorBlock(completionHandler)
	objc.Send[objc.ID](a.ID, objc.Sel("loadSamplePresentationTimeForTrackTime:completionHandler:"), trackTime, _block1)
}
// Creates a sample cursor and positions it at or near the specified
// presentation timestamp.
//
// presentationTimeStamp: The initial presentation timestamp of the sample cursor.
//
// # Return Value
// 
// An instance of [AVSampleCursor].
//
// # Discussion
// 
// If the track’s [Asset] property value for
// [providesPreciseDurationAndTiming] is [true], the sample cursor is
// accurately positioned at the track’slast media sample with a presentation
// timestamp less than or equal to the desired timestamp, or, if there are no
// such samples, the first sample in presentation order.
// 
// If the track’s [Asset] property value for
// [providesPreciseDurationAndTiming] is [false], and it’s prohibitively
// expensive to locate the precise sample at the desired timestamp, the sample
// cursor may be approximately positioned.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [providesPreciseDurationAndTiming]: https://developer.apple.com/documentation/AVFoundation/AVAsset/providesPreciseDurationAndTiming
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/makeSampleCursor(presentationTimeStamp:)
func (a AVAssetTrack) MakeSampleCursorWithPresentationTimeStamp(presentationTimeStamp uintptr) IAVSampleCursor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("makeSampleCursorWithPresentationTimeStamp:"), presentationTimeStamp)
	return AVSampleCursorFromID(rv)
}
// Creates a sample cursor and positions it at the track’s first media
// sample in decode order.
//
// # Return Value
// 
// An instance of [AVSampleCursor].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/makeSampleCursorAtFirstSampleInDecodeOrder()
func (a AVAssetTrack) MakeSampleCursorAtFirstSampleInDecodeOrder() IAVSampleCursor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("makeSampleCursorAtFirstSampleInDecodeOrder"))
	return AVSampleCursorFromID(rv)
}
// Creates a sample cursor and positions it at the track’s last media sample
// in decode order.
//
// # Return Value
// 
// An instance of [AVSampleCursor].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/makeSampleCursorAtLastSampleInDecodeOrder()
func (a AVAssetTrack) MakeSampleCursorAtLastSampleInDecodeOrder() IAVSampleCursor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("makeSampleCursorAtLastSampleInDecodeOrder"))
	return AVSampleCursorFromID(rv)
}

// The persistent unique identifier for this track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/trackID
func (a AVAssetTrack) TrackID() int32 {
	rv := objc.Send[int32](a.ID, objc.Sel("trackID"))
	return rv
}
// The type of media that a track presents.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/mediaType
func (a AVAssetTrack) MediaType() AVMediaType {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("mediaType"))
	return AVMediaType(foundation.NSStringFromID(rv).String())
}
// The asset object that contains this track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/asset
func (a AVAssetTrack) Asset() IAVAsset {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("asset"))
	return AVAssetFromID(objc.ID(rv))
}
// The media characteristics for the track.
//
// See: https://developer.apple.com/documentation/avfoundation/avpartialasyncproperty/mediacharacteristics
func (a AVAssetTrack) MediaCharacteristics() AVMediaCharacteristic {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("mediaCharacteristics"))
	return AVMediaCharacteristic(foundation.NSStringFromID(rv).String())
}
func (a AVAssetTrack) SetMediaCharacteristics(value AVMediaCharacteristic) {
	objc.Send[struct{}](a.ID, objc.Sel("setMediaCharacteristics:"), objc.String(string(value)))
}

			// Protocol methods for AVAsynchronousKeyValueLoading
			

// LoadSegmentForTrackTime is a synchronous wrapper around [AVAssetTrack.LoadSegmentForTrackTimeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAssetTrack) LoadSegmentForTrackTime(ctx context.Context, trackTime uintptr) (*AVAssetTrackSegment, error) {
	type result struct {
		val *AVAssetTrackSegment
		err error
	}
	done := make(chan result, 1)
	a.LoadSegmentForTrackTimeCompletionHandler(trackTime, func(val *AVAssetTrackSegment, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// LoadSamplePresentationTimeForTrackTime is a synchronous wrapper around [AVAssetTrack.LoadSamplePresentationTimeForTrackTimeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a AVAssetTrack) LoadSamplePresentationTimeForTrackTime(ctx context.Context, trackTime uintptr) (uintptr, error) {
	type result struct {
		val uintptr
		err error
	}
	done := make(chan result, 1)
	a.LoadSamplePresentationTimeForTrackTimeCompletionHandler(trackTime, func(val uintptr, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return 0, ctx.Err()
	}
}

