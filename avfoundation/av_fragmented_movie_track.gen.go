// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [AVFragmentedMovieTrack] class.
var (
	_AVFragmentedMovieTrackClass     AVFragmentedMovieTrackClass
	_AVFragmentedMovieTrackClassOnce sync.Once
)

func getAVFragmentedMovieTrackClass() AVFragmentedMovieTrackClass {
	_AVFragmentedMovieTrackClassOnce.Do(func() {
		_AVFragmentedMovieTrackClass = AVFragmentedMovieTrackClass{class: objc.GetClass("AVFragmentedMovieTrack")}
	})
	return _AVFragmentedMovieTrackClass
}

// GetAVFragmentedMovieTrackClass returns the class object for AVFragmentedMovieTrack.
func GetAVFragmentedMovieTrackClass() AVFragmentedMovieTrackClass {
	return getAVFragmentedMovieTrackClass()
}

type AVFragmentedMovieTrackClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVFragmentedMovieTrackClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVFragmentedMovieTrackClass) Alloc() AVFragmentedMovieTrack {
	rv := objc.Send[AVFragmentedMovieTrack](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a track in a fragmented movie.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieTrack
type AVFragmentedMovieTrack struct {
	AVMovieTrack
}

// AVFragmentedMovieTrackFromID constructs a [AVFragmentedMovieTrack] from an objc.ID.
//
// An object that represents a track in a fragmented movie.
func AVFragmentedMovieTrackFromID(id objc.ID) AVFragmentedMovieTrack {
	return AVFragmentedMovieTrack{AVMovieTrack: AVMovieTrackFromID(id)}
}
// NOTE: AVFragmentedMovieTrack adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVFragmentedMovieTrack] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedMovieTrack
type IAVFragmentedMovieTrack interface {
	IAVMovieTrack
}

// Init initializes the instance.
func (f AVFragmentedMovieTrack) Init() AVFragmentedMovieTrack {
	rv := objc.Send[AVFragmentedMovieTrack](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f AVFragmentedMovieTrack) Autorelease() AVFragmentedMovieTrack {
	rv := objc.Send[AVFragmentedMovieTrack](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVFragmentedMovieTrack creates a new AVFragmentedMovieTrack instance.
func NewAVFragmentedMovieTrack() AVFragmentedMovieTrack {
	class := getAVFragmentedMovieTrackClass()
	rv := objc.Send[AVFragmentedMovieTrack](objc.ID(class.class), objc.Sel("new"))
	return rv
}

