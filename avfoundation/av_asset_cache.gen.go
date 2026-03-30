// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVAssetCache] class.
var (
	_AVAssetCacheClass     AVAssetCacheClass
	_AVAssetCacheClassOnce sync.Once
)

func getAVAssetCacheClass() AVAssetCacheClass {
	_AVAssetCacheClassOnce.Do(func() {
		_AVAssetCacheClass = AVAssetCacheClass{class: objc.GetClass("AVAssetCache")}
	})
	return _AVAssetCacheClass
}

// GetAVAssetCacheClass returns the class object for AVAssetCache.
func GetAVAssetCacheClass() AVAssetCacheClass {
	return getAVAssetCacheClass()
}

type AVAssetCacheClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVAssetCacheClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVAssetCacheClass) Alloc() AVAssetCache {
	rv := objc.Send[AVAssetCache](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that you use to inspect locally cached media data.
//
// # Overview
//
// You can download HTTP Live Streaming assets to an iOS device using the
// [AVAssetDownloadURLSession] and [AVAssetDownloadTask] classes.
//
// # Inspecting the cached media
//
//   - [AVAssetCache.PlayableOffline]: A Boolean value that indicates whether the asset is playable without an internet connection.
//   - [AVAssetCache.MediaSelectionOptionsInMediaSelectionGroup]: Returns an array of locally cached media selection options that are available for offline use.
//   - [AVAssetCache.MediaPresentationLanguagesForMediaSelectionGroup]: Returns an array of extended language tags for languages that can be selected for offline operations via use of the AVMediaSelectionGroup’s AVCustomMediaSelectionScheme.
//   - [AVAssetCache.MediaPresentationSettingsForMediaSelectionGroup]: For each AVMediaPresentationSelector defined by the AVCustomMediaSelectionScheme of an AVMediaSelectionGroup, returns the AVMediaPresentationSettings that can be satisfied for offline operations, e.g. playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetCache
type AVAssetCache struct {
	objectivec.Object
}

// AVAssetCacheFromID constructs a [AVAssetCache] from an objc.ID.
//
// An object that you use to inspect locally cached media data.
func AVAssetCacheFromID(id objc.ID) AVAssetCache {
	return AVAssetCache{objectivec.Object{ID: id}}
}

// NOTE: AVAssetCache adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVAssetCache] class.
//
// # Inspecting the cached media
//
//   - [IAVAssetCache.PlayableOffline]: A Boolean value that indicates whether the asset is playable without an internet connection.
//   - [IAVAssetCache.MediaSelectionOptionsInMediaSelectionGroup]: Returns an array of locally cached media selection options that are available for offline use.
//   - [IAVAssetCache.MediaPresentationLanguagesForMediaSelectionGroup]: Returns an array of extended language tags for languages that can be selected for offline operations via use of the AVMediaSelectionGroup’s AVCustomMediaSelectionScheme.
//   - [IAVAssetCache.MediaPresentationSettingsForMediaSelectionGroup]: For each AVMediaPresentationSelector defined by the AVCustomMediaSelectionScheme of an AVMediaSelectionGroup, returns the AVMediaPresentationSettings that can be satisfied for offline operations, e.g. playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetCache
type IAVAssetCache interface {
	objectivec.IObject

	// Topic: Inspecting the cached media

	// A Boolean value that indicates whether the asset is playable without an internet connection.
	PlayableOffline() bool
	// Returns an array of locally cached media selection options that are available for offline use.
	MediaSelectionOptionsInMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) []AVMediaSelectionOption
	// Returns an array of extended language tags for languages that can be selected for offline operations via use of the AVMediaSelectionGroup’s AVCustomMediaSelectionScheme.
	MediaPresentationLanguagesForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) []string
	// For each AVMediaPresentationSelector defined by the AVCustomMediaSelectionScheme of an AVMediaSelectionGroup, returns the AVMediaPresentationSettings that can be satisfied for offline operations, e.g. playback.
	MediaPresentationSettingsForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) foundation.INSDictionary
}

// Init initializes the instance.
func (a AVAssetCache) Init() AVAssetCache {
	rv := objc.Send[AVAssetCache](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVAssetCache) Autorelease() AVAssetCache {
	rv := objc.Send[AVAssetCache](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVAssetCache creates a new AVAssetCache instance.
func NewAVAssetCache() AVAssetCache {
	class := getAVAssetCacheClass()
	rv := objc.Send[AVAssetCache](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an array of locally cached media selection options that are
// available for offline use.
//
// mediaSelectionGroup: The containing media selection group.
//
// # Return Value
//
// The array of media selection options, or an empty array if none are
// available.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetCache/mediaSelectionOptions(in:)
func (a AVAssetCache) MediaSelectionOptionsInMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) []AVMediaSelectionOption {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("mediaSelectionOptionsInMediaSelectionGroup:"), mediaSelectionGroup)
	return objc.ConvertSlice(rv, func(id objc.ID) AVMediaSelectionOption {
		return AVMediaSelectionOptionFromID(id)
	})
}

// Returns an array of extended language tags for languages that can be
// selected for offline operations via use of the AVMediaSelectionGroup’s
// AVCustomMediaSelectionScheme.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetCache/mediaPresentationLanguages(for:)
func (a AVAssetCache) MediaPresentationLanguagesForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) []string {
	rv := objc.Send[[]objc.ID](a.ID, objc.Sel("mediaPresentationLanguagesForMediaSelectionGroup:"), mediaSelectionGroup)
	return objc.ConvertSliceToStrings(rv)
}

// For each AVMediaPresentationSelector defined by the
// AVCustomMediaSelectionScheme of an AVMediaSelectionGroup, returns the
// AVMediaPresentationSettings that can be satisfied for offline operations,
// e.g. playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetCache/mediaPresentationSettings(for:)
func (a AVAssetCache) MediaPresentationSettingsForMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("mediaPresentationSettingsForMediaSelectionGroup:"), mediaSelectionGroup)
	return foundation.NSDictionaryFromID(rv)
}

// A Boolean value that indicates whether the asset is playable without an
// internet connection.
//
// # Discussion
//
// Check the value of this property to determine the asset’s suitability for
// playback before presenting or attempting to play it.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAssetCache/isPlayableOffline
func (a AVAssetCache) PlayableOffline() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isPlayableOffline"))
	return rv
}
