// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCompositionTrackSegment] class.
var (
	_AVCompositionTrackSegmentClass     AVCompositionTrackSegmentClass
	_AVCompositionTrackSegmentClassOnce sync.Once
)

func getAVCompositionTrackSegmentClass() AVCompositionTrackSegmentClass {
	_AVCompositionTrackSegmentClassOnce.Do(func() {
		_AVCompositionTrackSegmentClass = AVCompositionTrackSegmentClass{class: objc.GetClass("AVCompositionTrackSegment")}
	})
	return _AVCompositionTrackSegmentClass
}

// GetAVCompositionTrackSegmentClass returns the class object for AVCompositionTrackSegment.
func GetAVCompositionTrackSegmentClass() AVCompositionTrackSegmentClass {
	return getAVCompositionTrackSegmentClass()
}

type AVCompositionTrackSegmentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCompositionTrackSegmentClass) Alloc() AVCompositionTrackSegment {
	rv := objc.Send[AVCompositionTrackSegment](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A track segment that maps a time from the source media track to the
// composition track.
//
// # Overview
// 
// You typically use this class to save a low-level representation of a
// composition.
//
// # Creating a segment
//
//   - [AVCompositionTrackSegment.InitWithTimeRange]: Creates an object that presents an empty composition track segment.
//   - [AVCompositionTrackSegment.InitWithURLTrackIDSourceTimeRangeTargetTimeRange]: Creates an object that presents a segment of a media file that the specified URL references.
//
// # Accessing segment properties
//
//   - [AVCompositionTrackSegment.SourceURL]: A URL of the container file whose media this track segment presents.
//   - [AVCompositionTrackSegment.SourceTrackID]: An identifier of a track in the container file whose media this track segment presents.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackSegment
type AVCompositionTrackSegment struct {
	AVAssetTrackSegment
}

// AVCompositionTrackSegmentFromID constructs a [AVCompositionTrackSegment] from an objc.ID.
//
// A track segment that maps a time from the source media track to the
// composition track.
func AVCompositionTrackSegmentFromID(id objc.ID) AVCompositionTrackSegment {
	return AVCompositionTrackSegment{AVAssetTrackSegment: AVAssetTrackSegmentFromID(id)}
}
// NOTE: AVCompositionTrackSegment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCompositionTrackSegment] class.
//
// # Creating a segment
//
//   - [IAVCompositionTrackSegment.InitWithTimeRange]: Creates an object that presents an empty composition track segment.
//   - [IAVCompositionTrackSegment.InitWithURLTrackIDSourceTimeRangeTargetTimeRange]: Creates an object that presents a segment of a media file that the specified URL references.
//
// # Accessing segment properties
//
//   - [IAVCompositionTrackSegment.SourceURL]: A URL of the container file whose media this track segment presents.
//   - [IAVCompositionTrackSegment.SourceTrackID]: An identifier of a track in the container file whose media this track segment presents.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackSegment
type IAVCompositionTrackSegment interface {
	IAVAssetTrackSegment

	// Topic: Creating a segment

	// Creates an object that presents an empty composition track segment.
	InitWithTimeRange(timeRange objectivec.IObject) AVCompositionTrackSegment
	// Creates an object that presents a segment of a media file that the specified URL references.
	InitWithURLTrackIDSourceTimeRangeTargetTimeRange(URL foundation.INSURL, trackID int32, sourceTimeRange objectivec.IObject, targetTimeRange objectivec.IObject) AVCompositionTrackSegment

	// Topic: Accessing segment properties

	// A URL of the container file whose media this track segment presents.
	SourceURL() foundation.INSURL
	// An identifier of a track in the container file whose media this track segment presents.
	SourceTrackID() int32
}

// Init initializes the instance.
func (c AVCompositionTrackSegment) Init() AVCompositionTrackSegment {
	rv := objc.Send[AVCompositionTrackSegment](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCompositionTrackSegment) Autorelease() AVCompositionTrackSegment {
	rv := objc.Send[AVCompositionTrackSegment](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCompositionTrackSegment creates a new AVCompositionTrackSegment instance.
func NewAVCompositionTrackSegment() AVCompositionTrackSegment {
	class := getAVCompositionTrackSegmentClass()
	rv := objc.Send[AVCompositionTrackSegment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an object that presents an empty composition track segment.
//
// timeRange: The time range of the empty track segment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackSegment/init(timeRange:)
// timeRange is a [coremedia.CMTimeRange].
func NewCompositionTrackSegmentWithTimeRange(timeRange objectivec.IObject) AVCompositionTrackSegment {
	instance := getAVCompositionTrackSegmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTimeRange:"), timeRange)
	return AVCompositionTrackSegmentFromID(rv)
}

// Creates an object that presents a segment of a media file that the
// specified URL references.
//
// URL: A URL of the source media file.
//
// trackID: The identifier of the track whose media this segment presents.
//
// sourceTimeRange: The time range of the track whose media this segment presents.
//
// targetTimeRange: The time range of the composition track to present the segment’s media.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackSegment/init(url:trackID:sourceTimeRange:targetTimeRange:)
// sourceTimeRange is a [coremedia.CMTimeRange].
// targetTimeRange is a [coremedia.CMTimeRange].
func NewCompositionTrackSegmentWithURLTrackIDSourceTimeRangeTargetTimeRange(URL foundation.INSURL, trackID int32, sourceTimeRange objectivec.IObject, targetTimeRange objectivec.IObject) AVCompositionTrackSegment {
	instance := getAVCompositionTrackSegmentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:trackID:sourceTimeRange:targetTimeRange:"), URL, trackID, sourceTimeRange, targetTimeRange)
	return AVCompositionTrackSegmentFromID(rv)
}

// Creates an object that presents an empty composition track segment.
//
// timeRange: The time range of the empty track segment.
//
// timeRange is a [coremedia.CMTimeRange].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackSegment/init(timeRange:)
func (c AVCompositionTrackSegment) InitWithTimeRange(timeRange objectivec.IObject) AVCompositionTrackSegment {
	rv := objc.Send[AVCompositionTrackSegment](c.ID, objc.Sel("initWithTimeRange:"), timeRange)
	return rv
}
// Creates an object that presents a segment of a media file that the
// specified URL references.
//
// URL: A URL of the source media file.
//
// trackID: The identifier of the track whose media this segment presents.
//
// sourceTimeRange: The time range of the track whose media this segment presents.
//
// targetTimeRange: The time range of the composition track to present the segment’s media.
//
// sourceTimeRange is a [coremedia.CMTimeRange].
//
// targetTimeRange is a [coremedia.CMTimeRange].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackSegment/init(url:trackID:sourceTimeRange:targetTimeRange:)
func (c AVCompositionTrackSegment) InitWithURLTrackIDSourceTimeRangeTargetTimeRange(URL foundation.INSURL, trackID int32, sourceTimeRange objectivec.IObject, targetTimeRange objectivec.IObject) AVCompositionTrackSegment {
	rv := objc.Send[AVCompositionTrackSegment](c.ID, objc.Sel("initWithURL:trackID:sourceTimeRange:targetTimeRange:"), URL, trackID, sourceTimeRange, targetTimeRange)
	return rv
}

// Returns a new object that presents an empty composition track segment.
//
// timeRange: The time range of the empty composition track segment.
//
// timeRange is a [coremedia.CMTimeRange].
//
// # Return Value
// 
// A new composition track segment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackSegment/compositionTrackSegmentWithTimeRange:
func (_AVCompositionTrackSegmentClass AVCompositionTrackSegmentClass) CompositionTrackSegmentWithTimeRange(timeRange objectivec.IObject) AVCompositionTrackSegment {
	rv := objc.Send[objc.ID](objc.ID(_AVCompositionTrackSegmentClass.class), objc.Sel("compositionTrackSegmentWithTimeRange:"), timeRange)
	return AVCompositionTrackSegmentFromID(rv)
}
// Returns a new an object that presents a segment of a media file that the
// specified URL references.
//
// URL: A URL of the source media file.
//
// trackID: The identifier of the track whose media this segment presents.
//
// sourceTimeRange: The time range of the track whose media this segment presents.
//
// targetTimeRange: The time range of the composition track to present the segment’s media.
//
// sourceTimeRange is a [coremedia.CMTimeRange].
//
// targetTimeRange is a [coremedia.CMTimeRange].
//
// # Return Value
// 
// A new composition track segment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackSegment/compositionTrackSegmentWithURL:trackID:sourceTimeRange:targetTimeRange:
func (_AVCompositionTrackSegmentClass AVCompositionTrackSegmentClass) CompositionTrackSegmentWithURLTrackIDSourceTimeRangeTargetTimeRange(URL foundation.INSURL, trackID int32, sourceTimeRange objectivec.IObject, targetTimeRange objectivec.IObject) AVCompositionTrackSegment {
	rv := objc.Send[objc.ID](objc.ID(_AVCompositionTrackSegmentClass.class), objc.Sel("compositionTrackSegmentWithURL:trackID:sourceTimeRange:targetTimeRange:"), URL, trackID, sourceTimeRange, targetTimeRange)
	return AVCompositionTrackSegmentFromID(rv)
}

// A URL of the container file whose media this track segment presents.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackSegment/sourceURL
func (c AVCompositionTrackSegment) SourceURL() foundation.INSURL {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("sourceURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// An identifier of a track in the container file whose media this track
// segment presents.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrackSegment/sourceTrackID
func (c AVCompositionTrackSegment) SourceTrackID() int32 {
	rv := objc.Send[int32](c.ID, objc.Sel("sourceTrackID"))
	return rv
}

