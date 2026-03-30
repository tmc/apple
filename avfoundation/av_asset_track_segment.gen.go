// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetTrackSegment] class.
var (
	_AVAssetTrackSegmentClass     AVAssetTrackSegmentClass
	_AVAssetTrackSegmentClassOnce sync.Once
)

func getAVAssetTrackSegmentClass() AVAssetTrackSegmentClass {
	_AVAssetTrackSegmentClassOnce.Do(func() {
		_AVAssetTrackSegmentClass = AVAssetTrackSegmentClass{class: objc.GetClass("AVAssetTrackSegment")}
	})
	return _AVAssetTrackSegmentClass
}

// GetAVAssetTrackSegmentClass returns the class object for AVAssetTrackSegment.
func GetAVAssetTrackSegmentClass() AVAssetTrackSegmentClass {
	return getAVAssetTrackSegmentClass()
}

type AVAssetTrackSegmentClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetTrackSegmentClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetTrackSegmentClass) Alloc() AVAssetTrackSegment {
	rv := objc.Send[AVAssetTrackSegment](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a time range segment of an asset track.
//
// # Retrieving segment information
//
//   - [AVAssetTrackSegment.TimeMapping]: The time range of the track that this segment presents.
//   - [AVAssetTrackSegment.Empty]: A Boolean value that indicates whether the segment is empty.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackSegment
type AVAssetTrackSegment struct {
	objectivec.Object
}

// AVAssetTrackSegmentFromID constructs a [AVAssetTrackSegment] from an objc.ID.
//
// An object that represents a time range segment of an asset track.
func AVAssetTrackSegmentFromID(id objc.ID) AVAssetTrackSegment {
	return AVAssetTrackSegment{objectivec.Object{ID: id}}
}

// NOTE: AVAssetTrackSegment adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetTrackSegment] class.
//
// # Retrieving segment information
//
//   - [IAVAssetTrackSegment.TimeMapping]: The time range of the track that this segment presents.
//   - [IAVAssetTrackSegment.Empty]: A Boolean value that indicates whether the segment is empty.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackSegment
type IAVAssetTrackSegment interface {
	objectivec.IObject

	// Topic: Retrieving segment information

	// The time range of the track that this segment presents.
	TimeMapping() coremedia.CMTimeMapping
	// A Boolean value that indicates whether the segment is empty.
	Empty() bool
}

// Init initializes the instance.
func (a AVAssetTrackSegment) Init() AVAssetTrackSegment {
	rv := objc.Send[AVAssetTrackSegment](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetTrackSegment) Autorelease() AVAssetTrackSegment {
	rv := objc.Send[AVAssetTrackSegment](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetTrackSegment creates a new AVAssetTrackSegment instance.
func NewAVAssetTrackSegment() AVAssetTrackSegment {
	class := getAVAssetTrackSegmentClass()
	rv := objc.Send[AVAssetTrackSegment](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The time range of the track that this segment presents.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackSegment/timeMapping
func (a AVAssetTrackSegment) TimeMapping() coremedia.CMTimeMapping {
	rv := objc.Send[coremedia.CMTimeMapping](a.ID, objc.Sel("timeMapping"))
	return coremedia.CMTimeMapping(rv)
}

// A Boolean value that indicates whether the segment is empty.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackSegment/isEmpty
func (a AVAssetTrackSegment) Empty() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isEmpty"))
	return rv
}
