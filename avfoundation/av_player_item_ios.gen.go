// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
//go:build ios
// +build ios

package avfoundation

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Determines whether this item is subject to parental restrictions, and, if
// so, prompts the user to enter the restrictions passcode.
//
// completion: A callback the system invokes after it makes a determination of parental
// restrictions.
//
// `isAuthorized`: A Boolean value that indicates whether the system
// authorizes the app to play an item. `error`: An optional error that
// contains error details if the system encountered an error.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/requestPlaybackRestrictionsAuthorization(_:)
func (p AVPlayerItem) RequestPlaybackRestrictionsAuthorization(completion BoolErrorHandler) {
	_block0, _cleanup0 := NewBoolErrorBlock(completion)
	defer _cleanup0()
	objc.Send[objc.ID](p.ID, objc.Sel("requestPlaybackRestrictionsAuthorization:"), _block0)
}

// Cancels a pending authorization request and dismisses the passcode entry,
// if displayed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/cancelPlaybackRestrictionsAuthorizationRequest()
func (p AVPlayerItem) CancelPlaybackRestrictionsAuthorizationRequest() {
	objc.Send[objc.ID](p.ID, objc.Sel("cancelPlaybackRestrictionsAuthorizationRequest"))
}

// An array of additional metadata for the player item to supplement or
// replace an asset’s embedded metadata.
//
// # Discussion
//
// [AVPlayerViewController] supports displaying the following metadata
// identifiers:
//
// - [commonKeyTitle] - [iTunesMetadataTrackSubTitle] -
// [commonIdentifierArtwork] - [commonKeyDescription] -
// [iTunesMetadataKeyContentRating] - [quickTimeMetadataKeyGenre]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/externalMetadata
//
// [AVPlayerViewController]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController
// [commonIdentifierArtwork]: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/commonIdentifierArtwork
// [commonKeyDescription]: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyDescription
// [commonKeyTitle]: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyTitle
// [iTunesMetadataKeyContentRating]: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/iTunesMetadataKeyContentRating
// [iTunesMetadataTrackSubTitle]: https://developer.apple.com/documentation/AVFoundation/AVMetadataIdentifier/iTunesMetadataTrackSubTitle
// [quickTimeMetadataKeyGenre]: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/quickTimeMetadataKeyGenre
func (p AVPlayerItem) ExternalMetadata() []AVMetadataItem {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("externalMetadata"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}
func (p AVPlayerItem) SetExternalMetadata(value []AVMetadataItem) {
	objc.Send[struct{}](p.ID, objc.Sel("setExternalMetadata:"), objectivec.IObjectSliceToNSArray(value))
}

// A Boolean value that indicates whether the player translates interstitial
// events to interstitial time ranges.
//
// # Discussion
//
// Enable this property value to support using interstitial events, or set it
// to false to perform your own interstitial management.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/translatesPlayerInterstitialEvents
func (p AVPlayerItem) TranslatesPlayerInterstitialEvents() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("translatesPlayerInterstitialEvents"))
	return rv
}
func (p AVPlayerItem) SetTranslatesPlayerInterstitialEvents(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setTranslatesPlayerInterstitialEvents:"), value)
}

// An array of time ranges that identify interstitial content.
//
// # Discussion
//
// Interstitial content is material that’s unrelated to a player item’s
// primary content, such as advertisements and legal notices. If you use
// [AVPlayerViewController] to present an item that contains interstitial time
// ranges, the user interface marks those time ranges differently on the
// playback timeline. A player view controller can also call your app when it
// begins and ends playing interstitial content. You can use these events to
// customize playback behavior, such as preventing viewers from skipping
// required content.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/interstitialTimeRanges
//
// [AVPlayerViewController]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController
func (p AVPlayerItem) InterstitialTimeRanges() []objc.ID {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("interstitialTimeRanges"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVInterstitialTimeRange {
		return AVInterstitialTimeRangeFromID(id)
	})
}

// The current now playing information for the player item.
//
// # Discussion
//
// Setting this value to `nil` clears the player item’s now playing
// information.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/nowPlayingInfo
func (p AVPlayerItem) NowPlayingInfo() foundation.INSDictionary {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("nowPlayingInfo"))
	return foundation.NSDictionaryFromID(rv)
}
func (p AVPlayerItem) SetNowPlayingInfo(value foundation.INSDictionary) {
	objc.Send[struct{}](p.ID, objc.Sel("setNowPlayingInfo:"), value)
}

// The time marker groups that provide ways to navigate the player item’s
// content.
//
// # Discussion
//
// A navigation marker group provides a set of time markers for navigating
// playback. The most common form of navigation marker group is a chapter
// list; however, you can also provide other sets of markers to allow a user
// to jump to significant events in the presentation. For example, a “Goals
// Scored” marker group might summarize key moments in a recorded sporting
// event. When you present a player item containing marker groups with the
// [AVPlayerViewController] class, the user interface provides options for
// navigating each group.
//
// To provide a chapter list, use the first item in the
// [NavigationMarkerGroups] array and set its title property to `nil`. To
// provide additional or alternate means of navigating content, use a unique
// title value for each navigation marker group in the array.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/navigationMarkerGroups
//
// [AVPlayerViewController]: https://developer.apple.com/documentation/AVKit/AVPlayerViewController
func (p AVPlayerItem) NavigationMarkerGroups() []objc.ID {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("navigationMarkerGroups"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVNavigationMarkersGroup {
		return AVNavigationMarkersGroupFromID(id)
	})
}
func (p AVPlayerItem) SetNavigationMarkerGroups(value []objc.ID) {
	objc.Send[struct{}](p.ID, objc.Sel("setNavigationMarkerGroups:"), objectivec.IObjectSliceToNSArray(value))
}

// The item proposed to follow the current content.
//
// # Discussion
//
// You can use this property to associate content proposed to follow the
// current content.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/nextContentProposal
func (p AVPlayerItem) NextContentProposal() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p.ID, objc.Sel("nextContentProposal"))
	return rv
}
func (p AVPlayerItem) SetNextContentProposal(value unsafe.Pointer) {
	objc.Send[struct{}](p.ID, objc.Sel("setNextContentProposal:"), value)
}

// # Discussion
//
// An array of BCP 47 language codes that supplements the list of subtitle
// options that will be presented to the user.
//
// This list is strictly for non-standard, application-rendered subtitles,
// that cannot be handled by AVFoundation. Most clients should not need to set
// this property. The application should implement the
// playerViewController:didSelectExternalSubtitleOptionLanguage: method of its
// AVPlayerViewControllerDelegate to be notified when one of these languages
// has been chosen by the user.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/externalSubtitleOptionLanguages
func (p AVPlayerItem) ExternalSubtitleOptionLanguages() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("externalSubtitleOptionLanguages"))
	return objc.ConvertSliceToStrings(rv)
}
func (p AVPlayerItem) SetExternalSubtitleOptionLanguages(value []string) {
	objc.Send[struct{}](p.ID, objc.Sel("setExternalSubtitleOptionLanguages:"), objectivec.StringSliceToNSArray(value))
}

// # Discussion
//
// Specifies BCP 47 language code of the external subtitle option language
// marked in the user interface.
//
// If anything other than an external subtitle option is selected (including
// “Off”), then this property should be set to an empty string. If the
// value is not an empty string, it should be an element of the
// externalSubtitleOptionLanguages array.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVPlayerItem/selectedExternalSubtitleOptionLanguage
func (p AVPlayerItem) SelectedExternalSubtitleOptionLanguage() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("selectedExternalSubtitleOptionLanguage"))
	return foundation.NSStringFromID(rv).String()
}
func (p AVPlayerItem) SetSelectedExternalSubtitleOptionLanguage(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setSelectedExternalSubtitleOptionLanguage:"), objc.String(value))
}
