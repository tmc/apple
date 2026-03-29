// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMutableComposition] class.
var (
	_AVMutableCompositionClass     AVMutableCompositionClass
	_AVMutableCompositionClassOnce sync.Once
)

func getAVMutableCompositionClass() AVMutableCompositionClass {
	_AVMutableCompositionClassOnce.Do(func() {
		_AVMutableCompositionClass = AVMutableCompositionClass{class: objc.GetClass("AVMutableComposition")}
	})
	return _AVMutableCompositionClass
}

// GetAVMutableCompositionClass returns the class object for AVMutableComposition.
func GetAVMutableCompositionClass() AVMutableCompositionClass {
	return getAVMutableCompositionClass()
}

type AVMutableCompositionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMutableCompositionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMutableCompositionClass) Alloc() AVMutableComposition {
	rv := objc.Send[AVMutableComposition](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that you use to create a new composition from existing assets.
//
// # Overview
// 
// Use this object to add and remove composition tracks, and add, remove, and
// scale their time ranges. You can make an immutable snapshot of a mutable
// composition for playback and inspection as follows:
//
// # Managing composition tracks
//
//   - [AVMutableComposition.MutableTrackCompatibleWithTrack]: Returns a composition track into which you can insert any time range of the specified asset track.
//   - [AVMutableComposition.AddMutableTrackWithMediaTypePreferredTrackID]: Adds an empty track to a composition.
//   - [AVMutableComposition.RemoveTrack]: Removes a specified track from the composition.
//
// # Managing time ranges
//
//   - [AVMutableComposition.RemoveTimeRange]: Removes a specified time range from all tracks of the composition.
//   - [AVMutableComposition.ScaleTimeRangeToDuration]: Changes the duration of all tracks in a given time range.
//   - [AVMutableComposition.InsertEmptyTimeRange]: Adds or extends an empty time range within all tracks of the composition.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition
type AVMutableComposition struct {
	AVComposition
}

// AVMutableCompositionFromID constructs a [AVMutableComposition] from an objc.ID.
//
// An object that you use to create a new composition from existing assets.
func AVMutableCompositionFromID(id objc.ID) AVMutableComposition {
	return AVMutableComposition{AVComposition: AVCompositionFromID(id)}
}
// NOTE: AVMutableComposition adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMutableComposition] class.
//
// # Managing composition tracks
//
//   - [IAVMutableComposition.MutableTrackCompatibleWithTrack]: Returns a composition track into which you can insert any time range of the specified asset track.
//   - [IAVMutableComposition.AddMutableTrackWithMediaTypePreferredTrackID]: Adds an empty track to a composition.
//   - [IAVMutableComposition.RemoveTrack]: Removes a specified track from the composition.
//
// # Managing time ranges
//
//   - [IAVMutableComposition.RemoveTimeRange]: Removes a specified time range from all tracks of the composition.
//   - [IAVMutableComposition.ScaleTimeRangeToDuration]: Changes the duration of all tracks in a given time range.
//   - [IAVMutableComposition.InsertEmptyTimeRange]: Adds or extends an empty time range within all tracks of the composition.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition
type IAVMutableComposition interface {
	IAVComposition

	// Topic: Managing composition tracks

	// Returns a composition track into which you can insert any time range of the specified asset track.
	MutableTrackCompatibleWithTrack(track IAVAssetTrack) IAVMutableCompositionTrack
	// Adds an empty track to a composition.
	AddMutableTrackWithMediaTypePreferredTrackID(mediaType AVMediaType, preferredTrackID int32) IAVMutableCompositionTrack
	// Removes a specified track from the composition.
	RemoveTrack(track IAVCompositionTrack)

	// Topic: Managing time ranges

	// Removes a specified time range from all tracks of the composition.
	RemoveTimeRange(timeRange uintptr)
	// Changes the duration of all tracks in a given time range.
	ScaleTimeRangeToDuration(timeRange uintptr, duration uintptr)
	// Adds or extends an empty time range within all tracks of the composition.
	InsertEmptyTimeRange(timeRange uintptr)

	// Adds a group of empty tracks associated with a cinematic asset to a mutable composition.
	AddTracksForCinematicAssetInfoPreferredStartingTrackID(assetInfo objectivec.IObject, preferredStartingTrackID int32) objectivec.IObject
}

// Init initializes the instance.
func (m AVMutableComposition) Init() AVMutableComposition {
	rv := objc.Send[AVMutableComposition](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMutableComposition) Autorelease() AVMutableComposition {
	rv := objc.Send[AVMutableComposition](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMutableComposition creates a new AVMutableComposition instance.
func NewAVMutableComposition() AVMutableComposition {
	class := getAVMutableCompositionClass()
	rv := objc.Send[AVMutableComposition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an asset that models the media at the specified URL.
//
// URL: A URL to a local, remote, or HTTP Live Streaming media resource.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/init(url:)
func NewMutableCompositionAssetWithURL(URL foundation.INSURL) AVMutableComposition {
	rv := objc.Send[objc.ID](objc.ID(getAVMutableCompositionClass().class), objc.Sel("assetWithURL:"), URL)
	return AVMutableCompositionFromID(rv)
}

// Creates a mutable composition that uses the specified initialization
// options.
//
// URLAssetInitializationOptions: The initialization options to use to create the composition.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/init(urlAssetInitializationOptions:)
func NewMutableCompositionWithURLAssetInitializationOptions(URLAssetInitializationOptions foundation.INSDictionary) AVMutableComposition {
	rv := objc.Send[objc.ID](objc.ID(getAVMutableCompositionClass().class), objc.Sel("compositionWithURLAssetInitializationOptions:"), URLAssetInitializationOptions)
	return AVMutableCompositionFromID(rv)
}

// Returns a composition track into which you can insert any time range of the
// specified asset track.
//
// track: The asset track to find a composition track for.
//
// # Return Value
// 
// A mutable composition track, of `nil` if a compatible track isn’t
// available.
//
// # Discussion
// 
// To optimize performance, limit the number of tracks to only what you need
// to present media data in parallel. To present media data of the same type
// serially, even from multiple assets, use a single track of that media type.
// You use this method to identify a suitable existing target track for an
// insertion.
// 
// If there’s no compatible track available, you can create a new track of
// the same media type as `track` using
// [AddMutableTrackWithMediaTypePreferredTrackID].
// 
// This method is the counterpart to [compatibleTrack(for:)] on [AVAsset].
//
// [compatibleTrack(for:)]: https://developer.apple.com/documentation/AVFoundation/AVURLAsset/compatibleTrack(for:)
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/mutableTrack(compatibleWith:)
func (m AVMutableComposition) MutableTrackCompatibleWithTrack(track IAVAssetTrack) IAVMutableCompositionTrack {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mutableTrackCompatibleWithTrack:"), track)
	return AVMutableCompositionTrackFromID(rv)
}
// Adds an empty track to a composition.
//
// mediaType: The media type of the new track.
//
// preferredTrackID: The preferred track ID for the new track. The system generates a unique ID
// if the value you specify isn’t available. If you don’t need to specify
// a preferred track ID, pass [kCMPersistentTrackID_Invalid], and the system
// generates an appropriate identifier.
// //
// [kCMPersistentTrackID_Invalid]: https://developer.apple.com/documentation/CoreMedia/kCMPersistentTrackID_Invalid
//
// # Return Value
// 
// A new mutable composition track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/addMutableTrack(withMediaType:preferredTrackID:)
func (m AVMutableComposition) AddMutableTrackWithMediaTypePreferredTrackID(mediaType AVMediaType, preferredTrackID int32) IAVMutableCompositionTrack {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addMutableTrackWithMediaType:preferredTrackID:"), objc.String(string(mediaType)), preferredTrackID)
	return AVMutableCompositionTrackFromID(rv)
}
// Removes a specified track from the composition.
//
// track: The track to remove.
//
// # Discussion
// 
// When you remove a track, the system sets its composition value to nil.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/removeTrack(_:)
func (m AVMutableComposition) RemoveTrack(track IAVCompositionTrack) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeTrack:"), track)
}
// Removes a specified time range from all tracks of the composition.
//
// timeRange: The time range to remove.
//
// # Discussion
// 
// After removing, existing content after the time range moves forward in the
// composition timeline.
// 
// Removing a time range doesn’t remove any existing tracks from the
// composition, even if removing it results in an empty track. Instead, it
// removes or truncates track segments that intersect with the time range.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/removeTimeRange(_:)
func (m AVMutableComposition) RemoveTimeRange(timeRange uintptr) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeTimeRange:"), timeRange)
}
// Changes the duration of all tracks in a given time range.
//
// timeRange: The time range of the composition to scale.
//
// duration: The new time range duration.
//
// # Discussion
// 
// A composition presents each track segment affected by the scaling operation
// at a rate equal to `source.Duration() / target.Duration()` of its resulting
// time mapping.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/scaleTimeRange(_:toDuration:)
func (m AVMutableComposition) ScaleTimeRangeToDuration(timeRange uintptr, duration uintptr) {
	objc.Send[objc.ID](m.ID, objc.Sel("scaleTimeRange:toDuration:"), timeRange, duration)
}
// Adds or extends an empty time range within all tracks of the composition.
//
// timeRange: The empty time range to insert.
//
// # Discussion
// 
// Inserting an empty time range pushes out existing content by the time
// range’s duration. Use this method to reserve a time range in the
// composition for a subsequently created track to present its media.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/insertEmptyTimeRange(_:)
func (m AVMutableComposition) InsertEmptyTimeRange(timeRange uintptr) {
	objc.Send[objc.ID](m.ID, objc.Sel("insertEmptyTimeRange:"), timeRange)
}
// Adds a group of empty tracks associated with a cinematic asset to a mutable
// composition.
//
// assetInfo is a [cinematic.CNAssetInfo].
//
// # Return Value
// 
// Information about the composition tracks added to the mutable composition.
// Be sure to call insertTimeRange on the result to specify at least one time
// range of cinematic asset you’d like in the composition.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/addTracksForCinematicAssetInfo:preferredStartingTrackID:
// assetInfo is a [cinematic.CNAssetInfo].
func (m AVMutableComposition) AddTracksForCinematicAssetInfoPreferredStartingTrackID(assetInfo objectivec.IObject, preferredStartingTrackID int32) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addTracksForCinematicAssetInfo:preferredStartingTrackID:"), assetInfo, preferredStartingTrackID)
	return objectivec.Object{ID: rv}
}

// Returns a new mutable composition.
//
// # Return Value
// 
// A mutable composition.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableComposition/composition
func (_AVMutableCompositionClass AVMutableCompositionClass) Composition() AVMutableComposition {
	rv := objc.Send[objc.ID](objc.ID(_AVMutableCompositionClass.class), objc.Sel("composition"))
	return AVMutableCompositionFromID(rv)
}

