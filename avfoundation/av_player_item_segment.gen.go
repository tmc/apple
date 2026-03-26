// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerItemSegment] class.
var (
	_AVPlayerItemSegmentClass     AVPlayerItemSegmentClass
	_AVPlayerItemSegmentClassOnce sync.Once
)

func getAVPlayerItemSegmentClass() AVPlayerItemSegmentClass {
	_AVPlayerItemSegmentClassOnce.Do(func() {
		_AVPlayerItemSegmentClass = AVPlayerItemSegmentClass{class: objc.GetClass("AVPlayerItemSegment")}
	})
	return _AVPlayerItemSegmentClass
}

// GetAVPlayerItemSegmentClass returns the class object for AVPlayerItemSegment.
func GetAVPlayerItemSegmentClass() AVPlayerItemSegmentClass {
	return getAVPlayerItemSegmentClass()
}

type AVPlayerItemSegmentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerItemSegmentClass) Alloc() AVPlayerItemSegment {
	rv := objc.Send[AVPlayerItemSegment](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An immutable object that represents a segment of time on the integrated
// timeline.
//
// # Identifying the type
//
//   - [AVPlayerItemSegment.SegmentType]: The type content this segment represents.
//
// # Inspecting the segment
//
//   - [AVPlayerItemSegment.TimeMapping]: The time mapping for this segment.
//   - [AVPlayerItemSegment.StartDate]: The date at which a segment starts.
//   - [AVPlayerItemSegment.InterstitialEvent]: The associated interstitial event for this segment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSegment
type AVPlayerItemSegment struct {
	objectivec.Object
}

// AVPlayerItemSegmentFromID constructs a [AVPlayerItemSegment] from an objc.ID.
//
// An immutable object that represents a segment of time on the integrated
// timeline.
func AVPlayerItemSegmentFromID(id objc.ID) AVPlayerItemSegment {
	return AVPlayerItemSegment{objectivec.Object{ID: id}}
}
// NOTE: AVPlayerItemSegment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerItemSegment] class.
//
// # Identifying the type
//
//   - [IAVPlayerItemSegment.SegmentType]: The type content this segment represents.
//
// # Inspecting the segment
//
//   - [IAVPlayerItemSegment.TimeMapping]: The time mapping for this segment.
//   - [IAVPlayerItemSegment.StartDate]: The date at which a segment starts.
//   - [IAVPlayerItemSegment.InterstitialEvent]: The associated interstitial event for this segment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSegment
type IAVPlayerItemSegment interface {
	objectivec.IObject

	// Topic: Identifying the type

	// The type content this segment represents.
	SegmentType() AVPlayerItemSegmentType

	// Topic: Inspecting the segment

	// The time mapping for this segment.
	TimeMapping() objectivec.IObject
	// The date at which a segment starts.
	StartDate() foundation.INSDate
	// The associated interstitial event for this segment.
	InterstitialEvent() IAVPlayerInterstitialEvent

	// The current date on the integrated timeline when the system created the snapshot.
	CurrentDate() foundation.INSDate
	SetCurrentDate(value foundation.INSDate)
	// The currently playing segment.
	CurrentSegment() IAVPlayerItemSegment
	SetCurrentSegment(value IAVPlayerItemSegment)
	// The current time on the integrated timeline when the system created the snapshot.
	CurrentTime() objectivec.IObject
	SetCurrentTime(value objectivec.IObject)
	// The total duration of the primary item and scheduled interstitial events.
	Duration() objectivec.IObject
	SetDuration(value objectivec.IObject)
	// The time ranges for the segment that have media data is readily available.
	LoadedTimeRanges() []foundation.NSValue
	// The segments for this snapshot.
	Segments() IAVPlayerItemSegment
	SetSegments(value IAVPlayerItemSegment)
}

// Init initializes the instance.
func (p AVPlayerItemSegment) Init() AVPlayerItemSegment {
	rv := objc.Send[AVPlayerItemSegment](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerItemSegment) Autorelease() AVPlayerItemSegment {
	rv := objc.Send[AVPlayerItemSegment](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemSegment creates a new AVPlayerItemSegment instance.
func NewAVPlayerItemSegment() AVPlayerItemSegment {
	class := getAVPlayerItemSegmentClass()
	rv := objc.Send[AVPlayerItemSegment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The type content this segment represents.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSegment/segmentType-swift.property
func (p AVPlayerItemSegment) SegmentType() AVPlayerItemSegmentType {
	rv := objc.Send[AVPlayerItemSegmentType](p.ID, objc.Sel("segmentType"))
	return AVPlayerItemSegmentType(rv)
}
// The time mapping for this segment.
//
// # Discussion
// 
// The time mapping’s source time range represents the start time and
// duration in the segment source’s timeline, either the primary item or
// interstitial event. The target time range represents the start time and
// duration in the integrated timeline. For interstitial events that occupy a
// single point, the target’s duration is [zero].
//
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSegment/timeMapping
func (p AVPlayerItemSegment) TimeMapping() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("timeMapping"))
	return objectivec.Object{ID: rv}
}
// The date at which a segment starts.
//
// # Discussion
// 
// This value is `nil` if the primary item doesn’t contain dates.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSegment/startDate
func (p AVPlayerItemSegment) StartDate() foundation.INSDate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("startDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
// The associated interstitial event for this segment.
//
// # Discussion
// 
// This value is `nil` for segments that represent playback of the primary
// item.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSegment/interstitialEvent
func (p AVPlayerItemSegment) InterstitialEvent() IAVPlayerInterstitialEvent {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("interstitialEvent"))
	return AVPlayerInterstitialEventFromID(objc.ID(rv))
}
// The current date on the integrated timeline when the system created the
// snapshot.
//
// See: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimelinesnapshot/currentdate
func (p AVPlayerItemSegment) CurrentDate() foundation.INSDate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (p AVPlayerItemSegment) SetCurrentDate(value foundation.INSDate) {
	objc.Send[struct{}](p.ID, objc.Sel("setCurrentDate:"), value)
}
// The currently playing segment.
//
// See: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimelinesnapshot/currentsegment
func (p AVPlayerItemSegment) CurrentSegment() IAVPlayerItemSegment {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentSegment"))
	return AVPlayerItemSegmentFromID(objc.ID(rv))
}
func (p AVPlayerItemSegment) SetCurrentSegment(value IAVPlayerItemSegment) {
	objc.Send[struct{}](p.ID, objc.Sel("setCurrentSegment:"), value)
}
// The current time on the integrated timeline when the system created the
// snapshot.
//
// See: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimelinesnapshot/currenttime
func (p AVPlayerItemSegment) CurrentTime() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentTime"))
	return objectivec.Object{ID: rv}
}
func (p AVPlayerItemSegment) SetCurrentTime(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setCurrentTime:"), value)
}
// The total duration of the primary item and scheduled interstitial events.
//
// See: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimelinesnapshot/duration
func (p AVPlayerItemSegment) Duration() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("duration"))
	return objectivec.Object{ID: rv}
}
func (p AVPlayerItemSegment) SetDuration(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setDuration:"), value)
}
// The time ranges for the segment that have media data is readily available.
//
// # Discussion
// 
// The loaded time ranges might be discontinuous.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemSegment/loadedTimeRanges-2p0fl
func (p AVPlayerItemSegment) LoadedTimeRanges() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("loadedTimeRanges"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}
// The segments for this snapshot.
//
// See: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimelinesnapshot/segments
func (p AVPlayerItemSegment) Segments() IAVPlayerItemSegment {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("segments"))
	return AVPlayerItemSegmentFromID(objc.ID(rv))
}
func (p AVPlayerItemSegment) SetSegments(value IAVPlayerItemSegment) {
	objc.Send[struct{}](p.ID, objc.Sel("setSegments:"), value)
}

