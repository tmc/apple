// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetTrackGroup] class.
var (
	_AVAssetTrackGroupClass     AVAssetTrackGroupClass
	_AVAssetTrackGroupClassOnce sync.Once
)

func getAVAssetTrackGroupClass() AVAssetTrackGroupClass {
	_AVAssetTrackGroupClassOnce.Do(func() {
		_AVAssetTrackGroupClass = AVAssetTrackGroupClass{class: objc.GetClass("AVAssetTrackGroup")}
	})
	return _AVAssetTrackGroupClass
}

// GetAVAssetTrackGroupClass returns the class object for AVAssetTrackGroup.
func GetAVAssetTrackGroupClass() AVAssetTrackGroupClass {
	return getAVAssetTrackGroupClass()
}

type AVAssetTrackGroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetTrackGroupClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetTrackGroupClass) Alloc() AVAssetTrackGroup {
	rv := objc.Send[AVAssetTrackGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A group of related tracks in an asset.
//
// # Overview
// 
// A track group describes a group of related alternative tracks, only one of
// which should play at a time. Groups of alternative tracks typically contain
// variations of the same content, like subtitles in multiple translations.
// 
// You can inspect an asset’s track groups by loading the value of its
// [trackGroups] property.
//
// [trackGroups]: https://developer.apple.com/documentation/AVFoundation/AVPartialAsyncProperty/trackGroups
//
// # Getting track ID values
//
//   - [AVAssetTrackGroup.TrackIDs]: The IDs of the tracks in the group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackGroup
type AVAssetTrackGroup struct {
	objectivec.Object
}

// AVAssetTrackGroupFromID constructs a [AVAssetTrackGroup] from an objc.ID.
//
// A group of related tracks in an asset.
func AVAssetTrackGroupFromID(id objc.ID) AVAssetTrackGroup {
	return AVAssetTrackGroup{objectivec.Object{ID: id}}
}
// NOTE: AVAssetTrackGroup adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetTrackGroup] class.
//
// # Getting track ID values
//
//   - [IAVAssetTrackGroup.TrackIDs]: The IDs of the tracks in the group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackGroup
type IAVAssetTrackGroup interface {
	objectivec.IObject

	// Topic: Getting track ID values

	// The IDs of the tracks in the group.
	TrackIDs() []foundation.NSNumber

	// The track groups an asset contains.
	TrackGroups() IAVAssetTrackGroup
	SetTrackGroups(value IAVAssetTrackGroup)
}

// Init initializes the instance.
func (a AVAssetTrackGroup) Init() AVAssetTrackGroup {
	rv := objc.Send[AVAssetTrackGroup](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetTrackGroup) Autorelease() AVAssetTrackGroup {
	rv := objc.Send[AVAssetTrackGroup](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetTrackGroup creates a new AVAssetTrackGroup instance.
func NewAVAssetTrackGroup() AVAssetTrackGroup {
	class := getAVAssetTrackGroupClass()
	rv := objc.Send[AVAssetTrackGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The IDs of the tracks in the group.
//
// # Discussion
// 
// The value of this property is an array of [NSNumber] instances used as
// [CMPersistentTrackID] values, one for each track in the group.
//
// [CMPersistentTrackID]: https://developer.apple.com/documentation/CoreMedia/CMPersistentTrackID
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetTrackGroup/trackIDs
func (a AVAssetTrackGroup) TrackIDs() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("trackIDs"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
// The track groups an asset contains.
//
// See: https://developer.apple.com/documentation/avfoundation/avpartialasyncproperty/trackgroups
func (a AVAssetTrackGroup) TrackGroups() IAVAssetTrackGroup {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("trackGroups"))
	return AVAssetTrackGroupFromID(objc.ID(rv))
}
func (a AVAssetTrackGroup) SetTrackGroups(value IAVAssetTrackGroup) {
	objc.Send[struct{}](a.ID, objc.Sel("setTrackGroups:"), value)
}

