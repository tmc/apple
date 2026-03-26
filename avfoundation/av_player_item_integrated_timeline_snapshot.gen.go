// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVPlayerItemIntegratedTimelineSnapshot] class.
var (
	_AVPlayerItemIntegratedTimelineSnapshotClass     AVPlayerItemIntegratedTimelineSnapshotClass
	_AVPlayerItemIntegratedTimelineSnapshotClassOnce sync.Once
)

func getAVPlayerItemIntegratedTimelineSnapshotClass() AVPlayerItemIntegratedTimelineSnapshotClass {
	_AVPlayerItemIntegratedTimelineSnapshotClassOnce.Do(func() {
		_AVPlayerItemIntegratedTimelineSnapshotClass = AVPlayerItemIntegratedTimelineSnapshotClass{class: objc.GetClass("AVPlayerItemIntegratedTimelineSnapshot")}
	})
	return _AVPlayerItemIntegratedTimelineSnapshotClass
}

// GetAVPlayerItemIntegratedTimelineSnapshotClass returns the class object for AVPlayerItemIntegratedTimelineSnapshot.
func GetAVPlayerItemIntegratedTimelineSnapshotClass() AVPlayerItemIntegratedTimelineSnapshotClass {
	return getAVPlayerItemIntegratedTimelineSnapshotClass()
}

type AVPlayerItemIntegratedTimelineSnapshotClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVPlayerItemIntegratedTimelineSnapshotClass) Alloc() AVPlayerItemIntegratedTimelineSnapshot {
	rv := objc.Send[AVPlayerItemIntegratedTimelineSnapshot](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An immutable representation of inspectable details of an integrated
// timeline object.
//
// # Overview
// 
// A snapshot doesn’t reflect the new timeline state as playback progresses.
// You can request a new snapshot instance from an
// [AVPlayerItemIntegratedTimeline] that reflect the latest timeline state.
//
// # Inspecting the snapshot
//
//   - [AVPlayerItemIntegratedTimelineSnapshot.Duration]: The total duration of the primary item and scheduled interstitial events.
//   - [AVPlayerItemIntegratedTimelineSnapshot.CurrentSegment]: The currently playing segment.
//   - [AVPlayerItemIntegratedTimelineSnapshot.Segments]: The segments for this snapshot.
//   - [AVPlayerItemIntegratedTimelineSnapshot.CurrentTime]: The current time on the integrated timeline when the system created the snapshot.
//   - [AVPlayerItemIntegratedTimelineSnapshot.CurrentDate]: The current date on the integrated timeline when the system created the snapshot.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot
type AVPlayerItemIntegratedTimelineSnapshot struct {
	objectivec.Object
}

// AVPlayerItemIntegratedTimelineSnapshotFromID constructs a [AVPlayerItemIntegratedTimelineSnapshot] from an objc.ID.
//
// An immutable representation of inspectable details of an integrated
// timeline object.
func AVPlayerItemIntegratedTimelineSnapshotFromID(id objc.ID) AVPlayerItemIntegratedTimelineSnapshot {
	return AVPlayerItemIntegratedTimelineSnapshot{objectivec.Object{ID: id}}
}
// NOTE: AVPlayerItemIntegratedTimelineSnapshot adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVPlayerItemIntegratedTimelineSnapshot] class.
//
// # Inspecting the snapshot
//
//   - [IAVPlayerItemIntegratedTimelineSnapshot.Duration]: The total duration of the primary item and scheduled interstitial events.
//   - [IAVPlayerItemIntegratedTimelineSnapshot.CurrentSegment]: The currently playing segment.
//   - [IAVPlayerItemIntegratedTimelineSnapshot.Segments]: The segments for this snapshot.
//   - [IAVPlayerItemIntegratedTimelineSnapshot.CurrentTime]: The current time on the integrated timeline when the system created the snapshot.
//   - [IAVPlayerItemIntegratedTimelineSnapshot.CurrentDate]: The current date on the integrated timeline when the system created the snapshot.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot
type IAVPlayerItemIntegratedTimelineSnapshot interface {
	objectivec.IObject

	// Topic: Inspecting the snapshot

	// The total duration of the primary item and scheduled interstitial events.
	Duration() objectivec.IObject
	// The currently playing segment.
	CurrentSegment() IAVPlayerItemSegment
	// The segments for this snapshot.
	Segments() []AVPlayerItemSegment
	// The current time on the integrated timeline when the system created the snapshot.
	CurrentTime() objectivec.IObject
	// The current date on the integrated timeline when the system created the snapshot.
	CurrentDate() foundation.INSDate

	// An immutable representation of the timeline state at time of request.
	CurrentSnapshot() IAVPlayerItemIntegratedTimelineSnapshot
	SetCurrentSnapshot(value IAVPlayerItemIntegratedTimelineSnapshot)
	MapTimeToSegmentAtSegmentOffset(time objectivec.IObject, timeSegmentOut *AVPlayerItemSegment, segmentOffsetOut objectivec.IObject)
}

// Init initializes the instance.
func (p AVPlayerItemIntegratedTimelineSnapshot) Init() AVPlayerItemIntegratedTimelineSnapshot {
	rv := objc.Send[AVPlayerItemIntegratedTimelineSnapshot](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p AVPlayerItemIntegratedTimelineSnapshot) Autorelease() AVPlayerItemIntegratedTimelineSnapshot {
	rv := objc.Send[AVPlayerItemIntegratedTimelineSnapshot](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVPlayerItemIntegratedTimelineSnapshot creates a new AVPlayerItemIntegratedTimelineSnapshot instance.
func NewAVPlayerItemIntegratedTimelineSnapshot() AVPlayerItemIntegratedTimelineSnapshot {
	class := getAVPlayerItemIntegratedTimelineSnapshotClass()
	rv := objc.Send[AVPlayerItemIntegratedTimelineSnapshot](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// time is a [coremedia.CMTime].
//
// segmentOffsetOut is a [coremedia.CMTime].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/mapTime:toSegment:atSegmentOffset:
func (p AVPlayerItemIntegratedTimelineSnapshot) MapTimeToSegmentAtSegmentOffset(time objectivec.IObject, timeSegmentOut *AVPlayerItemSegment, segmentOffsetOut objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("mapTime:toSegment:atSegmentOffset:"), time, timeSegmentOut, segmentOffsetOut)
}

// The total duration of the primary item and scheduled interstitial events.
//
// # Discussion
// 
// The duration property takes into account the interstitial event’s
// [PlayoutLimit] and [ResumptionOffset] values.
// 
// Before loading the duration of the primary item, the value of this property
// is [invalid]. For livestreams, this value is [indefinite].
//
// [indefinite]: https://developer.apple.com/documentation/CoreMedia/CMTime/indefinite
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/duration
func (p AVPlayerItemIntegratedTimelineSnapshot) Duration() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("duration"))
	return objectivec.Object{ID: rv}
}
// The currently playing segment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/currentSegment
func (p AVPlayerItemIntegratedTimelineSnapshot) CurrentSegment() IAVPlayerItemSegment {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentSegment"))
	return AVPlayerItemSegmentFromID(objc.ID(rv))
}
// The segments for this snapshot.
//
// # Discussion
// 
// The system presents segments in chronological order, contiguous from the
// previous element, and non-overlapping.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/segments
func (p AVPlayerItemIntegratedTimelineSnapshot) Segments() []AVPlayerItemSegment {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("segments"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVPlayerItemSegment {
		return AVPlayerItemSegmentFromID(id)
	})
}
// The current time on the integrated timeline when the system created the
// snapshot.
//
// # Discussion
// 
// This value doesn’t change as time progresses.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/currentTime
func (p AVPlayerItemIntegratedTimelineSnapshot) CurrentTime() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentTime"))
	return objectivec.Object{ID: rv}
}
// The current date on the integrated timeline when the system created the
// snapshot.
//
// # Discussion
// 
// This value is `nil` if playback doesn’t map to a date.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItemIntegratedTimelineSnapshot/currentDate
func (p AVPlayerItemIntegratedTimelineSnapshot) CurrentDate() foundation.INSDate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
// An immutable representation of the timeline state at time of request.
//
// See: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimeline/currentsnapshot
func (p AVPlayerItemIntegratedTimelineSnapshot) CurrentSnapshot() IAVPlayerItemIntegratedTimelineSnapshot {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("currentSnapshot"))
	return AVPlayerItemIntegratedTimelineSnapshotFromID(objc.ID(rv))
}
func (p AVPlayerItemIntegratedTimelineSnapshot) SetCurrentSnapshot(value IAVPlayerItemIntegratedTimelineSnapshot) {
	objc.Send[struct{}](p.ID, objc.Sel("setCurrentSnapshot:"), value)
}

// A notification the system posts when the snapshot objects provided by this
// timeline become out of sync with the current timeline state.
//
// See: https://developer.apple.com/documentation/avfoundation/avplayeritemintegratedtimeline/snapshotsoutofsyncnotification
func (_AVPlayerItemIntegratedTimelineSnapshotClass AVPlayerItemIntegratedTimelineSnapshotClass) SnapshotsOutOfSyncNotification() foundation.NSString {
	rv := objc.Send[objc.ID](objc.ID(_AVPlayerItemIntegratedTimelineSnapshotClass.class), objc.Sel("AVPlayerIntegratedTimelineSnapshotsOutOfSyncNotification"))
	return foundation.NSStringFromID(objc.ID(rv))
}

