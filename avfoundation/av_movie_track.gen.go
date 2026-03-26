// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMovieTrack] class.
var (
	_AVMovieTrackClass     AVMovieTrackClass
	_AVMovieTrackClassOnce sync.Once
)

func getAVMovieTrackClass() AVMovieTrackClass {
	_AVMovieTrackClassOnce.Do(func() {
		_AVMovieTrackClass = AVMovieTrackClass{class: objc.GetClass("AVMovieTrack")}
	})
	return _AVMovieTrackClass
}

// GetAVMovieTrackClass returns the class object for AVMovieTrack.
func GetAVMovieTrackClass() AVMovieTrackClass {
	return getAVMovieTrackClass()
}

type AVMovieTrackClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMovieTrackClass) Alloc() AVMovieTrack {
	rv := objc.Send[AVMovieTrack](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A track in a movie that conforms to the QuickTime or ISO base media file
// format.
//
// # Retrieving track information
//
//   - [AVMovieTrack.AlternateGroupID]: A value that identifies the track as a member of a particular alternate group.
//   - [AVMovieTrack.MediaDataStorage]: The storage container for media data added to a track.
//   - [AVMovieTrack.MediaDecodeTimeRange]: A range of decode times for the track’s media.
//   - [AVMovieTrack.MediaPresentationTimeRange]: A range of presentation times for the track’s media.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovieTrack
type AVMovieTrack struct {
	AVAssetTrack
}

// AVMovieTrackFromID constructs a [AVMovieTrack] from an objc.ID.
//
// A track in a movie that conforms to the QuickTime or ISO base media file
// format.
func AVMovieTrackFromID(id objc.ID) AVMovieTrack {
	return AVMovieTrack{AVAssetTrack: AVAssetTrackFromID(id)}
}
// NOTE: AVMovieTrack adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMovieTrack] class.
//
// # Retrieving track information
//
//   - [IAVMovieTrack.AlternateGroupID]: A value that identifies the track as a member of a particular alternate group.
//   - [IAVMovieTrack.MediaDataStorage]: The storage container for media data added to a track.
//   - [IAVMovieTrack.MediaDecodeTimeRange]: A range of decode times for the track’s media.
//   - [IAVMovieTrack.MediaPresentationTimeRange]: A range of presentation times for the track’s media.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovieTrack
type IAVMovieTrack interface {
	IAVAssetTrack

	// Topic: Retrieving track information

	// A value that identifies the track as a member of a particular alternate group.
	AlternateGroupID() int
	// The storage container for media data added to a track.
	MediaDataStorage() IAVMediaDataStorage
	// A range of decode times for the track’s media.
	MediaDecodeTimeRange() objectivec.IObject
	// A range of presentation times for the track’s media.
	MediaPresentationTimeRange() objectivec.IObject
}

// Init initializes the instance.
func (m AVMovieTrack) Init() AVMovieTrack {
	rv := objc.Send[AVMovieTrack](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMovieTrack) Autorelease() AVMovieTrack {
	rv := objc.Send[AVMovieTrack](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMovieTrack creates a new AVMovieTrack instance.
func NewAVMovieTrack() AVMovieTrack {
	class := getAVMovieTrackClass()
	rv := objc.Send[AVMovieTrack](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A value that identifies the track as a member of a particular alternate
// group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovieTrack/alternateGroupID
func (m AVMovieTrack) AlternateGroupID() int {
	rv := objc.Send[int](m.ID, objc.Sel("alternateGroupID"))
	return rv
}
// The storage container for media data added to a track.
//
// # Discussion
// 
// The value of this property is an [AVMediaDataStorage] object that indicates
// the location to which the system writes media data when it’s inserted or
// appended.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovieTrack/mediaDataStorage
func (m AVMovieTrack) MediaDataStorage() IAVMediaDataStorage {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mediaDataStorage"))
	return AVMediaDataStorageFromID(objc.ID(rv))
}
// A range of decode times for the track’s media.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovieTrack/mediaDecodeTimeRange
func (m AVMovieTrack) MediaDecodeTimeRange() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mediaDecodeTimeRange"))
	return objectivec.Object{ID: rv}
}
// A range of presentation times for the track’s media.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovieTrack/mediaPresentationTimeRange
func (m AVMovieTrack) MediaPresentationTimeRange() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mediaPresentationTimeRange"))
	return objectivec.Object{ID: rv}
}

