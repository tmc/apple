// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVFragmentedAssetTrack] class.
var (
	_AVFragmentedAssetTrackClass     AVFragmentedAssetTrackClass
	_AVFragmentedAssetTrackClassOnce sync.Once
)

func getAVFragmentedAssetTrackClass() AVFragmentedAssetTrackClass {
	_AVFragmentedAssetTrackClassOnce.Do(func() {
		_AVFragmentedAssetTrackClass = AVFragmentedAssetTrackClass{class: objc.GetClass("AVFragmentedAssetTrack")}
	})
	return _AVFragmentedAssetTrackClass
}

// GetAVFragmentedAssetTrackClass returns the class object for AVFragmentedAssetTrack.
func GetAVFragmentedAssetTrackClass() AVFragmentedAssetTrackClass {
	return getAVFragmentedAssetTrackClass()
}

type AVFragmentedAssetTrackClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVFragmentedAssetTrackClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVFragmentedAssetTrackClass) Alloc() AVFragmentedAssetTrack {
	rv := objc.Send[AVFragmentedAssetTrack](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that provides the track-level interface to inspect a fragmented
// asset’s media tracks.
//
// # Overview
//
// This class subclasses [AVAssetTrack]. It has no methods or properties of
// its own.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetTrack
type AVFragmentedAssetTrack struct {
	AVAssetTrack
}

// AVFragmentedAssetTrackFromID constructs a [AVFragmentedAssetTrack] from an objc.ID.
//
// An object that provides the track-level interface to inspect a fragmented
// asset’s media tracks.
func AVFragmentedAssetTrackFromID(id objc.ID) AVFragmentedAssetTrack {
	return AVFragmentedAssetTrack{AVAssetTrack: AVAssetTrackFromID(id)}
}

// NOTE: AVFragmentedAssetTrack adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVFragmentedAssetTrack] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVFragmentedAssetTrack
type IAVFragmentedAssetTrack interface {
	IAVAssetTrack
}

// Init initializes the instance.
func (f AVFragmentedAssetTrack) Init() AVFragmentedAssetTrack {
	rv := objc.Send[AVFragmentedAssetTrack](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f AVFragmentedAssetTrack) Autorelease() AVFragmentedAssetTrack {
	rv := objc.Send[AVFragmentedAssetTrack](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVFragmentedAssetTrack creates a new AVFragmentedAssetTrack instance.
func NewAVFragmentedAssetTrack() AVFragmentedAssetTrack {
	class := getAVFragmentedAssetTrackClass()
	rv := objc.Send[AVFragmentedAssetTrack](objc.ID(class.class), objc.Sel("new"))
	return rv
}
