// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVComposition] class.
var (
	_AVCompositionClass     AVCompositionClass
	_AVCompositionClassOnce sync.Once
)

func getAVCompositionClass() AVCompositionClass {
	_AVCompositionClassOnce.Do(func() {
		_AVCompositionClass = AVCompositionClass{class: objc.GetClass("AVComposition")}
	})
	return _AVCompositionClass
}

// GetAVCompositionClass returns the class object for AVComposition.
func GetAVCompositionClass() AVCompositionClass {
	return getAVCompositionClass()
}

type AVCompositionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCompositionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCompositionClass) Alloc() AVComposition {
	rv := objc.Send[AVComposition](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that combines and arranges media from multiple assets into a
// single composite asset that you can play or process.
//
// # Overview
// 
// A composition is a container for one or more tracks of media. Its tracks
// are instances of [AVCompositionTrack] that present media of a uniform type
// like audio or video. A track itself is a container for one or more segments
// of media, which are instances of [AVCompositionTrackSegment], a type that
// represents a region of media in the source track.
//
// # Accessing tracks
//
//   - [AVComposition.Tracks]: The tracks that a composition contains.
//   - [AVComposition.TrackWithTrackID]: Returns a track that contains the specified identifier.
//   - [AVComposition.TracksWithMediaType]: Returns tracks that contain media of a specified type.
//   - [AVComposition.TracksWithMediaCharacteristic]: Returns tracks that contain media of a specified characteristic.
//   - [AVComposition.UnusedTrackID]: Returns an identifier that no other tracks in the asset use.
//
// # Accessing track groups
//
//   - [AVComposition.TrackGroups]: The track groups an asset contains.
//   - [AVComposition.SetTrackGroups]
//
// # Accessing duration and timing
//
//   - [AVComposition.Duration]: A time value that indicates the asset’s duration.
//   - [AVComposition.SetDuration]
//   - [AVComposition.ProvidesPreciseDurationAndTiming]: A Boolean value that indicates whether the asset provides precise duration and timing.
//   - [AVComposition.SetProvidesPreciseDurationAndTiming]
//   - [AVComposition.MinimumTimeOffsetFromLive]: A time value that indicates how closely playback follows the latest live stream content.
//   - [AVComposition.SetMinimumTimeOffsetFromLive]
//
// # Accessing metadata
//
//   - [AVComposition.Metadata]: An array of metadata items for all metadata identifiers for which a value is available.
//   - [AVComposition.SetMetadata]
//   - [AVComposition.CommonMetadata]: The metadata items an asset contains for common metadata identifiers that provide a value.
//   - [AVComposition.SetCommonMetadata]
//   - [AVComposition.AvailableMetadataFormats]: The metadata formats this asset contains.
//   - [AVComposition.SetAvailableMetadataFormats]
//   - [AVComposition.MetadataForFormat]: Returns an array of metadata items from the container with the specified format.
//   - [AVComposition.CreationDate]: A metadata item that indicates the asset’s creation date.
//   - [AVComposition.SetCreationDate]
//   - [AVComposition.Lyrics]: The lyrics of the asset in a language suitable for the current locale.
//   - [AVComposition.SetLyrics]
//
// # Determining suitability
//
//   - [AVComposition.IsPlayable]: A Boolean value that indicates whether the asset has playable content.
//   - [AVComposition.SetIsPlayable]
//   - [AVComposition.IsReadable]: A Boolean value that indicates whether you can extract the asset’s media data using an asset reader.
//   - [AVComposition.SetIsReadable]
//   - [AVComposition.IsExportable]: A Boolean value that indicates whether you can export this asset using an export session.
//   - [AVComposition.SetIsExportable]
//   - [AVComposition.IsComposable]: A Boolean value that indicates whether you can use the asset as a segment of a composition track.
//   - [AVComposition.SetIsComposable]
//   - [AVComposition.IsCompatibleWithAirPlayVideo]: A Boolean value that indicates whether the asset is compatible with AirPlay Video.
//   - [AVComposition.SetIsCompatibleWithAirPlayVideo]
//
// # Inspecting preferences
//
//   - [AVComposition.PreferredRate]: The asset’s rate preference for playing its media.
//   - [AVComposition.SetPreferredRate]
//   - [AVComposition.PreferredVolume]: The asset’s volume preference for playing its audible media.
//   - [AVComposition.SetPreferredVolume]
//   - [AVComposition.PreferredTransform]: The asset’s transform preference to apply to its visual content during presentation or processing.
//   - [AVComposition.SetPreferredTransform]
//   - [AVComposition.PreferredMediaSelection]: The default media selections for this asset’s media selection groups.
//   - [AVComposition.SetPreferredMediaSelection]
//
// # Accessing media selections
//
//   - [AVComposition.AllMediaSelections]: The array of available media selections for this asset.
//   - [AVComposition.SetAllMediaSelections]
//   - [AVComposition.AvailableMediaCharacteristicsWithMediaSelectionOptions]: An array of media characteristics for which a media selection option is available.
//   - [AVComposition.SetAvailableMediaCharacteristicsWithMediaSelectionOptions]
//   - [AVComposition.MediaSelectionGroupForMediaCharacteristic]: Returns a media selection group that contains one or more options with the specified media characteristic.
//
// # Accessing chapter metadata
//
//   - [AVComposition.AvailableChapterLocales]: The locales of the asset’s chapter metadata.
//   - [AVComposition.SetAvailableChapterLocales]
//   - [AVComposition.ChapterMetadataGroupsBestMatchingPreferredLanguages]: Returns an array of chapters with a locale that best matches the list of preferred languages.
//   - [AVComposition.ChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeys]: Returns an array of chapters that contain the specified title locale and common keys.
//
// # Accessing initialization options
//
//   - [AVComposition.URLAssetInitializationOptions]: The options you used to create a composition.
//
// # Determining content protections
//
//   - [AVComposition.HasProtectedContent]: A Boolean value that indicates whether the asset contains protected content.
//   - [AVComposition.SetHasProtectedContent]
//
// # Determining fragment support
//
//   - [AVComposition.CanContainFragments]: A Boolean value that indicates whether you can extend the asset by fragments.
//   - [AVComposition.SetCanContainFragments]
//   - [AVComposition.ContainsFragments]: A Boolean value that indicates whether at least one movie fragment extends the asset.
//   - [AVComposition.SetContainsFragments]
//   - [AVComposition.OverallDurationHint]: The total duration of fragments that currently exist, or may exist in the future.
//   - [AVComposition.SetOverallDurationHint]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition
type AVComposition struct {
	AVAsset
}

// AVCompositionFromID constructs a [AVComposition] from an objc.ID.
//
// An object that combines and arranges media from multiple assets into a
// single composite asset that you can play or process.
func AVCompositionFromID(id objc.ID) AVComposition {
	return AVComposition{AVAsset: AVAssetFromID(id)}
}
// NOTE: AVComposition adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVComposition] class.
//
// # Accessing tracks
//
//   - [IAVComposition.Tracks]: The tracks that a composition contains.
//   - [IAVComposition.TrackWithTrackID]: Returns a track that contains the specified identifier.
//   - [IAVComposition.TracksWithMediaType]: Returns tracks that contain media of a specified type.
//   - [IAVComposition.TracksWithMediaCharacteristic]: Returns tracks that contain media of a specified characteristic.
//   - [IAVComposition.UnusedTrackID]: Returns an identifier that no other tracks in the asset use.
//
// # Accessing track groups
//
//   - [IAVComposition.TrackGroups]: The track groups an asset contains.
//   - [IAVComposition.SetTrackGroups]
//
// # Accessing duration and timing
//
//   - [IAVComposition.Duration]: A time value that indicates the asset’s duration.
//   - [IAVComposition.SetDuration]
//   - [IAVComposition.ProvidesPreciseDurationAndTiming]: A Boolean value that indicates whether the asset provides precise duration and timing.
//   - [IAVComposition.SetProvidesPreciseDurationAndTiming]
//   - [IAVComposition.MinimumTimeOffsetFromLive]: A time value that indicates how closely playback follows the latest live stream content.
//   - [IAVComposition.SetMinimumTimeOffsetFromLive]
//
// # Accessing metadata
//
//   - [IAVComposition.Metadata]: An array of metadata items for all metadata identifiers for which a value is available.
//   - [IAVComposition.SetMetadata]
//   - [IAVComposition.CommonMetadata]: The metadata items an asset contains for common metadata identifiers that provide a value.
//   - [IAVComposition.SetCommonMetadata]
//   - [IAVComposition.AvailableMetadataFormats]: The metadata formats this asset contains.
//   - [IAVComposition.SetAvailableMetadataFormats]
//   - [IAVComposition.MetadataForFormat]: Returns an array of metadata items from the container with the specified format.
//   - [IAVComposition.CreationDate]: A metadata item that indicates the asset’s creation date.
//   - [IAVComposition.SetCreationDate]
//   - [IAVComposition.Lyrics]: The lyrics of the asset in a language suitable for the current locale.
//   - [IAVComposition.SetLyrics]
//
// # Determining suitability
//
//   - [IAVComposition.IsPlayable]: A Boolean value that indicates whether the asset has playable content.
//   - [IAVComposition.SetIsPlayable]
//   - [IAVComposition.IsReadable]: A Boolean value that indicates whether you can extract the asset’s media data using an asset reader.
//   - [IAVComposition.SetIsReadable]
//   - [IAVComposition.IsExportable]: A Boolean value that indicates whether you can export this asset using an export session.
//   - [IAVComposition.SetIsExportable]
//   - [IAVComposition.IsComposable]: A Boolean value that indicates whether you can use the asset as a segment of a composition track.
//   - [IAVComposition.SetIsComposable]
//   - [IAVComposition.IsCompatibleWithAirPlayVideo]: A Boolean value that indicates whether the asset is compatible with AirPlay Video.
//   - [IAVComposition.SetIsCompatibleWithAirPlayVideo]
//
// # Inspecting preferences
//
//   - [IAVComposition.PreferredRate]: The asset’s rate preference for playing its media.
//   - [IAVComposition.SetPreferredRate]
//   - [IAVComposition.PreferredVolume]: The asset’s volume preference for playing its audible media.
//   - [IAVComposition.SetPreferredVolume]
//   - [IAVComposition.PreferredTransform]: The asset’s transform preference to apply to its visual content during presentation or processing.
//   - [IAVComposition.SetPreferredTransform]
//   - [IAVComposition.PreferredMediaSelection]: The default media selections for this asset’s media selection groups.
//   - [IAVComposition.SetPreferredMediaSelection]
//
// # Accessing media selections
//
//   - [IAVComposition.AllMediaSelections]: The array of available media selections for this asset.
//   - [IAVComposition.SetAllMediaSelections]
//   - [IAVComposition.AvailableMediaCharacteristicsWithMediaSelectionOptions]: An array of media characteristics for which a media selection option is available.
//   - [IAVComposition.SetAvailableMediaCharacteristicsWithMediaSelectionOptions]
//   - [IAVComposition.MediaSelectionGroupForMediaCharacteristic]: Returns a media selection group that contains one or more options with the specified media characteristic.
//
// # Accessing chapter metadata
//
//   - [IAVComposition.AvailableChapterLocales]: The locales of the asset’s chapter metadata.
//   - [IAVComposition.SetAvailableChapterLocales]
//   - [IAVComposition.ChapterMetadataGroupsBestMatchingPreferredLanguages]: Returns an array of chapters with a locale that best matches the list of preferred languages.
//   - [IAVComposition.ChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeys]: Returns an array of chapters that contain the specified title locale and common keys.
//
// # Accessing initialization options
//
//   - [IAVComposition.URLAssetInitializationOptions]: The options you used to create a composition.
//
// # Determining content protections
//
//   - [IAVComposition.HasProtectedContent]: A Boolean value that indicates whether the asset contains protected content.
//   - [IAVComposition.SetHasProtectedContent]
//
// # Determining fragment support
//
//   - [IAVComposition.CanContainFragments]: A Boolean value that indicates whether you can extend the asset by fragments.
//   - [IAVComposition.SetCanContainFragments]
//   - [IAVComposition.ContainsFragments]: A Boolean value that indicates whether at least one movie fragment extends the asset.
//   - [IAVComposition.SetContainsFragments]
//   - [IAVComposition.OverallDurationHint]: The total duration of fragments that currently exist, or may exist in the future.
//   - [IAVComposition.SetOverallDurationHint]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition
type IAVComposition interface {
	IAVAsset

	// Topic: Accessing tracks

	// The tracks that a composition contains.
	Tracks() []AVCompositionTrack
	// Returns a track that contains the specified identifier.
	TrackWithTrackID(trackID int32) IAVCompositionTrack
	// Returns tracks that contain media of a specified type.
	TracksWithMediaType(mediaType AVMediaType) []AVCompositionTrack
	// Returns tracks that contain media of a specified characteristic.
	TracksWithMediaCharacteristic(mediaCharacteristic AVMediaCharacteristic) []AVCompositionTrack
	// Returns an identifier that no other tracks in the asset use.
	UnusedTrackID() int32

	// Topic: Accessing track groups

	// The track groups an asset contains.
	TrackGroups() IAVAssetTrackGroup
	SetTrackGroups(value IAVAssetTrackGroup)

	// Topic: Accessing duration and timing

	// A time value that indicates the asset’s duration.
	Duration() coremedia.CMTime
	SetDuration(value coremedia.CMTime)
	// A Boolean value that indicates whether the asset provides precise duration and timing.
	ProvidesPreciseDurationAndTiming() bool
	SetProvidesPreciseDurationAndTiming(value bool)
	// A time value that indicates how closely playback follows the latest live stream content.
	MinimumTimeOffsetFromLive() coremedia.CMTime
	SetMinimumTimeOffsetFromLive(value coremedia.CMTime)

	// Topic: Accessing metadata

	// An array of metadata items for all metadata identifiers for which a value is available.
	Metadata() IAVMetadataItem
	SetMetadata(value IAVMetadataItem)
	// The metadata items an asset contains for common metadata identifiers that provide a value.
	CommonMetadata() IAVMetadataItem
	SetCommonMetadata(value IAVMetadataItem)
	// The metadata formats this asset contains.
	AvailableMetadataFormats() AVMetadataFormat
	SetAvailableMetadataFormats(value AVMetadataFormat)
	// Returns an array of metadata items from the container with the specified format.
	MetadataForFormat(format AVMetadataFormat) []AVMetadataItem
	// A metadata item that indicates the asset’s creation date.
	CreationDate() IAVMetadataItem
	SetCreationDate(value IAVMetadataItem)
	// The lyrics of the asset in a language suitable for the current locale.
	Lyrics() string
	SetLyrics(value string)

	// Topic: Determining suitability

	// A Boolean value that indicates whether the asset has playable content.
	IsPlayable() bool
	SetIsPlayable(value bool)
	// A Boolean value that indicates whether you can extract the asset’s media data using an asset reader.
	IsReadable() bool
	SetIsReadable(value bool)
	// A Boolean value that indicates whether you can export this asset using an export session.
	IsExportable() bool
	SetIsExportable(value bool)
	// A Boolean value that indicates whether you can use the asset as a segment of a composition track.
	IsComposable() bool
	SetIsComposable(value bool)
	// A Boolean value that indicates whether the asset is compatible with AirPlay Video.
	IsCompatibleWithAirPlayVideo() bool
	SetIsCompatibleWithAirPlayVideo(value bool)

	// Topic: Inspecting preferences

	// The asset’s rate preference for playing its media.
	PreferredRate() float32
	SetPreferredRate(value float32)
	// The asset’s volume preference for playing its audible media.
	PreferredVolume() float32
	SetPreferredVolume(value float32)
	// The asset’s transform preference to apply to its visual content during presentation or processing.
	PreferredTransform() corefoundation.CGAffineTransform
	SetPreferredTransform(value corefoundation.CGAffineTransform)
	// The default media selections for this asset’s media selection groups.
	PreferredMediaSelection() IAVMediaSelection
	SetPreferredMediaSelection(value IAVMediaSelection)

	// Topic: Accessing media selections

	// The array of available media selections for this asset.
	AllMediaSelections() IAVMediaSelection
	SetAllMediaSelections(value IAVMediaSelection)
	// An array of media characteristics for which a media selection option is available.
	AvailableMediaCharacteristicsWithMediaSelectionOptions() AVMediaCharacteristic
	SetAvailableMediaCharacteristicsWithMediaSelectionOptions(value AVMediaCharacteristic)
	// Returns a media selection group that contains one or more options with the specified media characteristic.
	MediaSelectionGroupForMediaCharacteristic(mediaCharacteristic AVMediaCharacteristic) IAVMediaSelectionGroup

	// Topic: Accessing chapter metadata

	// The locales of the asset’s chapter metadata.
	AvailableChapterLocales() objectivec.IObject
	SetAvailableChapterLocales(value objectivec.IObject)
	// Returns an array of chapters with a locale that best matches the list of preferred languages.
	ChapterMetadataGroupsBestMatchingPreferredLanguages(preferredLanguages []string) []AVTimedMetadataGroup
	// Returns an array of chapters that contain the specified title locale and common keys.
	ChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeys(locale foundation.NSLocale, commonKeys []string) []AVTimedMetadataGroup

	// Topic: Accessing initialization options

	// The options you used to create a composition.
	URLAssetInitializationOptions() foundation.INSDictionary

	// Topic: Determining content protections

	// A Boolean value that indicates whether the asset contains protected content.
	HasProtectedContent() bool
	SetHasProtectedContent(value bool)

	// Topic: Determining fragment support

	// A Boolean value that indicates whether you can extend the asset by fragments.
	CanContainFragments() bool
	SetCanContainFragments(value bool)
	// A Boolean value that indicates whether at least one movie fragment extends the asset.
	ContainsFragments() bool
	SetContainsFragments(value bool)
	// The total duration of fragments that currently exist, or may exist in the future.
	OverallDurationHint() coremedia.CMTime
	SetOverallDurationHint(value coremedia.CMTime)
}

// Init initializes the instance.
func (c AVComposition) Init() AVComposition {
	rv := objc.Send[AVComposition](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVComposition) Autorelease() AVComposition {
	rv := objc.Send[AVComposition](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVComposition creates a new AVComposition instance.
func NewAVComposition() AVComposition {
	class := getAVCompositionClass()
	rv := objc.Send[AVComposition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an asset that models the media at the specified URL.
//
// URL: A URL to a local, remote, or HTTP Live Streaming media resource.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/init(url:)
func NewCompositionAssetWithURL(URL foundation.INSURL) AVComposition {
	rv := objc.Send[objc.ID](objc.ID(getAVCompositionClass().class), objc.Sel("assetWithURL:"), URL)
	return AVCompositionFromID(rv)
}

// Returns a track that contains the specified identifier.
//
// trackID: The identifier of the track to retrieve.
//
// # Return Value
// 
// A composition track, or `nil` if no track with the identifier exists.
//
// # Discussion
// 
// Apple discourages using this method in iOS 15, tvOS 15, macOS 12, and
// watchOS 8 or later. Load a track asynchronously using
// [LoadTrackWithTrackIDCompletionHandler] instead.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/track(withTrackID:)
func (c AVComposition) TrackWithTrackID(trackID int32) IAVCompositionTrack {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("trackWithTrackID:"), trackID)
	return AVCompositionTrackFromID(rv)
}
// Returns tracks that contain media of a specified type.
//
// mediaType: The media type of the tracks to return.
//
// # Return Value
// 
// An array of tracks, which is empty if no tracks with the media type exist.
//
// # Discussion
// 
// Apple discourages using this method in iOS 15, tvOS 15, macOS 12, and
// watchOS 8 or later. Load tracks asynchronously using
// [LoadTracksWithMediaTypeCompletionHandler] instead.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/tracks(withMediaType:)
func (c AVComposition) TracksWithMediaType(mediaType AVMediaType) []AVCompositionTrack {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("tracksWithMediaType:"), objc.String(string(mediaType)))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCompositionTrack {
		return AVCompositionTrackFromID(id)
	})
}
// Returns tracks that contain media of a specified characteristic.
//
// mediaCharacteristic: The media characteristic of the tracks to return.
//
// # Return Value
// 
// An array of tracks, which is empty if no tracks with the media
// characteristic exist.
//
// # Discussion
// 
// Apple discourages using this method in iOS 15, tvOS 15, macOS 12, and
// watchOS 8 or later. Load tracks asynchronously using
// [LoadTracksWithMediaCharacteristicCompletionHandler] instead.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/tracks(withMediaCharacteristic:)
func (c AVComposition) TracksWithMediaCharacteristic(mediaCharacteristic AVMediaCharacteristic) []AVCompositionTrack {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("tracksWithMediaCharacteristic:"), objc.String(string(mediaCharacteristic)))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCompositionTrack {
		return AVCompositionTrackFromID(id)
	})
}
// Returns an identifier that no other tracks in the asset use.
//
// # Return Value
// 
// An unused [CMPersistentTrackID] value.
//
// [CMPersistentTrackID]: https://developer.apple.com/documentation/CoreMedia/CMPersistentTrackID
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/unusedTrackID()
func (c AVComposition) UnusedTrackID() int32 {
	rv := objc.Send[int32](c.ID, objc.Sel("unusedTrackID"))
	return rv
}
// Returns an array of metadata items from the container with the specified
// format.
//
// format: The metadata format for which you want items.
//
// # Return Value
// 
// An array of [AVMetadataItem] objects, one for each metadata item in the
// container of the specified format, or an empty array if there is no
// metadata for the specified format.
//
// # Discussion
// 
// You can filter the array to the specific items of interest using the class
// methods provided by [AVMetadataItem], like
// [MetadataItemsFromArrayFilteredByIdentifier] or
// [MetadataItemsFromArrayWithLocale].
// 
// You can call this method without blocking the current thread after you’ve
// asynchronously loaded the [availableMetadataFormats] property.
//
// [availableMetadataFormats]: https://developer.apple.com/documentation/AVFoundation/AVAsset/availableMetadataFormats
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/metadata(forFormat:)
func (c AVComposition) MetadataForFormat(format AVMetadataFormat) []AVMetadataItem {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("metadataForFormat:"), objc.String(string(format)))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}
// Returns a media selection group that contains one or more options with the
// specified media characteristic.
//
// mediaCharacteristic: A media characteristic for which to obtain the available media selection
// options.
// 
// Only [audible], [visual], and [legible] are currently supported.
// 
// - Pass [audible] to return the group of available options for audio media
// in various languages and for various purposes, such as descriptive audio. -
// Pass [legible] to return the group of available options for subtitles in
// various languages and for various purposes. - Pass [visual] to return the
// group of available options for video media.
// //
// [audible]: https://developer.apple.com/documentation/AVFoundation/AVMediaCharacteristic/audible
// [legible]: https://developer.apple.com/documentation/AVFoundation/AVMediaCharacteristic/legible
// [visual]: https://developer.apple.com/documentation/AVFoundation/AVMediaCharacteristic/visual
//
// # Return Value
// 
// An [AVMediaSelectionGroup] that contains one or more options with the
// specified media characteristic, or `nil` if none could be found.
//
// # Discussion
// 
// Use the filtering methods [AVMediaSelectionGroup] defines to filter the
// group’s options according to playability, locale, and additional media
// characteristics.
// 
// You can call this method without blocking the current thread after you’ve
// asynchronously loaded the
// [availableMediaCharacteristicsWithMediaSelectionOptions] property.
//
// [availableMediaCharacteristicsWithMediaSelectionOptions]: https://developer.apple.com/documentation/AVFoundation/AVAsset/availableMediaCharacteristicsWithMediaSelectionOptions
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/mediaSelectionGroup(forMediaCharacteristic:)
func (c AVComposition) MediaSelectionGroupForMediaCharacteristic(mediaCharacteristic AVMediaCharacteristic) IAVMediaSelectionGroup {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("mediaSelectionGroupForMediaCharacteristic:"), objc.String(string(mediaCharacteristic)))
	return AVMediaSelectionGroupFromID(rv)
}
// Returns an array of chapters with a locale that best matches the list of
// preferred languages.
//
// preferredLanguages: An array of BCP 47 language identifier strings. The order of the
// identifiers in the array reflects the preferred language order, with the
// most preferred language being first in the array. Typically, you pass the
// user’s preferred languages by retrieving this array from the
// [preferredLanguages] class method of [NSLocale].
// //
// [NSLocale]: https://developer.apple.com/documentation/Foundation/NSLocale
// [preferredLanguages]: https://developer.apple.com/documentation/Foundation/NSLocale/preferredLanguages
//
// # Return Value
// 
// An array of [AVTimedMetadataGroup] objects.
//
// # Discussion
// 
// Each object in the returned array contains an [AVMetadataItem] object
// representing the chapter title. The time range property of the
// [AVTimedMetadataGroup] object is equal to the time range of the chapter
// title item.
// 
// The metadata group contains all chapter metadata, including items with the
// common key [commonKeyArtwork], if such items are present. The system adds
// an [AVMetadataItem] with the specified common key to an existing
// [AVTimedMetadataGroup] object if the time range (timestamp and duration) of
// the metadata item and the metadata group overlap. The locale of such items
// don’t need to match the locale of the chapter titles.
// 
// You can use the
// [MetadataItemsFromArrayFilteredAndSortedAccordingToPreferredLanguages]
// method to further filter the metadata items in each group. You can also
// filter the returned items based on locale using the
// [MetadataItemsFromArrayWithLocale] method.
// 
// This method is callable without blocking the current thread after you’ve
// asynchronously loaded the [availableChapterLocales] property.
//
// [availableChapterLocales]: https://developer.apple.com/documentation/AVFoundation/AVAsset/availableChapterLocales
// [commonKeyArtwork]: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyArtwork
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/chapterMetadataGroups(bestMatchingPreferredLanguages:)
func (c AVComposition) ChapterMetadataGroupsBestMatchingPreferredLanguages(preferredLanguages []string) []AVTimedMetadataGroup {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("chapterMetadataGroupsBestMatchingPreferredLanguages:"), objectivec.StringSliceToNSArray(preferredLanguages))
	return objc.ConvertSlice(rv, func(id objc.ID) AVTimedMetadataGroup {
		return AVTimedMetadataGroupFromID(id)
	})
}
// Returns an array of chapters that contain the specified title locale and
// common keys.
//
// locale: The locale of the metadata items carrying chapter titles.
//
// commonKeys: An array of common keys of [AVMetadataItem] to include in the returned
// array. The framework currently only supports the [commonKeyArtwork] key.
// //
// [commonKeyArtwork]: https://developer.apple.com/documentation/AVFoundation/AVMetadataKey/commonKeyArtwork
//
// # Return Value
// 
// An array of [AVTimedMetadataGroup] objects.
//
// # Discussion
// 
// A metadata group contains an [AVMetadataItem] object that represents the
// chapter title, and a time range equal to the time range of the chapter
// title item.
// 
// The system adds a metadata item with the specified common key to an
// existing [AVTimedMetadataGroup] object if the time range (timestamp and
// duration) of the metadata item and the metadata group overlap.
// 
// The locale of items that don’t contain chapter titles doesn’t need to
// match the specified locale parameter. You can filter the returned items
// based on locale using [MetadataItemsFromArrayWithLocale].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/chapterMetadataGroups(withTitleLocale:containingItemsWithCommonKeys:)
func (c AVComposition) ChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeys(locale foundation.NSLocale, commonKeys []string) []AVTimedMetadataGroup {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("chapterMetadataGroupsWithTitleLocale:containingItemsWithCommonKeys:"), locale, objectivec.StringSliceToNSArray(commonKeys))
	return objc.ConvertSlice(rv, func(id objc.ID) AVTimedMetadataGroup {
		return AVTimedMetadataGroupFromID(id)
	})
}

// The tracks that a composition contains.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/tracks
func (c AVComposition) Tracks() []AVCompositionTrack {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("tracks"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCompositionTrack {
		return AVCompositionTrackFromID(id)
	})
}
// The track groups an asset contains.
//
// # Discussion
// 
// This value is an empty array if the composition has no track groups.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/trackGroups
func (c AVComposition) TrackGroups() IAVAssetTrackGroup {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("trackGroups"))
	return AVAssetTrackGroupFromID(objc.ID(rv))
}
func (c AVComposition) SetTrackGroups(value IAVAssetTrackGroup) {
	objc.Send[struct{}](c.ID, objc.Sel("setTrackGroups:"), value)
}
// A time value that indicates the asset’s duration.
//
// # Discussion
// 
// If you initialized the composition’s assets by passing the
// [AVURLAssetPreferPreciseDurationAndTimingKey] initialization option, this
// property value provides precise duration; otherwise, it provides a
// best-available estimate. You can determine the value’s accuracy by
// querying the asset’s [providesPreciseDurationAndTiming] property.
//
// [AVURLAssetPreferPreciseDurationAndTimingKey]: https://developer.apple.com/documentation/AVFoundation/AVURLAssetPreferPreciseDurationAndTimingKey
// [providesPreciseDurationAndTiming]: https://developer.apple.com/documentation/AVFoundation/AVAsset/providesPreciseDurationAndTiming
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/duration
func (c AVComposition) Duration() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](c.ID, objc.Sel("duration"))
	return coremedia.CMTime(rv)
}
func (c AVComposition) SetDuration(value coremedia.CMTime) {
	objc.Send[struct{}](c.ID, objc.Sel("setDuration:"), value)
}
// A Boolean value that indicates whether the asset provides precise duration
// and timing.
//
// # Discussion
// 
// This property value is [true] if you initialized the asset with the
// [AVURLAssetPreferPreciseDurationAndTimingKey] initialization option,
// otherwise it’s [false].
//
// [AVURLAssetPreferPreciseDurationAndTimingKey]: https://developer.apple.com/documentation/AVFoundation/AVURLAssetPreferPreciseDurationAndTimingKey
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/providesPreciseDurationAndTiming
func (c AVComposition) ProvidesPreciseDurationAndTiming() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("providesPreciseDurationAndTiming"))
	return rv
}
func (c AVComposition) SetProvidesPreciseDurationAndTiming(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setProvidesPreciseDurationAndTiming:"), value)
}
// A time value that indicates how closely playback follows the latest live
// stream content.
//
// # Discussion
// 
// This property value is only valid when working with live streaming content.
// For non-live assets, this property value is [invalid].
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/minimumTimeOffsetFromLive
func (c AVComposition) MinimumTimeOffsetFromLive() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](c.ID, objc.Sel("minimumTimeOffsetFromLive"))
	return coremedia.CMTime(rv)
}
func (c AVComposition) SetMinimumTimeOffsetFromLive(value coremedia.CMTime) {
	objc.Send[struct{}](c.ID, objc.Sel("setMinimumTimeOffsetFromLive:"), value)
}
// An array of metadata items for all metadata identifiers for which a value
// is available.
//
// # Discussion
// 
// You can filter the metadata items by language using the
// [MetadataItemsFromArrayFilteredAndSortedAccordingToPreferredLanguages]
// method, or by identifier with the
// [MetadataItemsFromArrayFilteredByIdentifier] method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/metadata
func (c AVComposition) Metadata() IAVMetadataItem {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("metadata"))
	return AVMetadataItemFromID(objc.ID(rv))
}
func (c AVComposition) SetMetadata(value IAVMetadataItem) {
	objc.Send[struct{}](c.ID, objc.Sel("setMetadata:"), value)
}
// The metadata items an asset contains for common metadata identifiers that
// provide a value.
//
// # Discussion
// 
// This property value is an array of metadata items, one for each metadata
// key from the common key space for which the asset has an available value.
// You can use the various class methods provided by [AVMetadataItem], such as
// [MetadataItemsFromArrayFilteredByIdentifier] or
// [MetadataItemsFromArrayWithLocale] to filter the array to the specific
// items of interest.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/commonMetadata
func (c AVComposition) CommonMetadata() IAVMetadataItem {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("commonMetadata"))
	return AVMetadataItemFromID(objc.ID(rv))
}
func (c AVComposition) SetCommonMetadata(value IAVMetadataItem) {
	objc.Send[struct{}](c.ID, objc.Sel("setCommonMetadata:"), value)
}
// The metadata formats this asset contains.
//
// # Discussion
// 
// Metadata formats may include ID3, iTunes metadata, and so on.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/availableMetadataFormats
func (c AVComposition) AvailableMetadataFormats() AVMetadataFormat {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("availableMetadataFormats"))
	return AVMetadataFormat(foundation.NSStringFromID(rv).String())
}
func (c AVComposition) SetAvailableMetadataFormats(value AVMetadataFormat) {
	objc.Send[struct{}](c.ID, objc.Sel("setAvailableMetadataFormats:"), objc.String(string(value)))
}
// A metadata item that indicates the asset’s creation date.
//
// # Discussion
// 
// If the asset contains metadata that the framework can convert to an
// [NSDate], you can retrieve it from the metadata item using its [dateValue]
// property. Otherwise, you retrieve it as a string by using the metadata
// item’s [stringValue] property.
// 
// This property value is `nil` if no creation date metadata exists.
//
// [NSDate]: https://developer.apple.com/documentation/Foundation/NSDate
// [dateValue]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/dateValue
// [stringValue]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem/stringValue
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/creationDate
func (c AVComposition) CreationDate() IAVMetadataItem {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("creationDate"))
	return AVMetadataItemFromID(objc.ID(rv))
}
func (c AVComposition) SetCreationDate(value IAVMetadataItem) {
	objc.Send[struct{}](c.ID, objc.Sel("setCreationDate:"), value)
}
// The lyrics of the asset in a language suitable for the current locale.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/lyrics
func (c AVComposition) Lyrics() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("lyrics"))
	return foundation.NSStringFromID(rv).String()
}
func (c AVComposition) SetLyrics(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setLyrics:"), objc.String(value))
}
// A Boolean value that indicates whether the asset has playable content.
//
// # Discussion
// 
// This property value is [true] if you can use the composition to create an
// [AVPlayerItem].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/isPlayable
func (c AVComposition) IsPlayable() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isPlayable"))
	return rv
}
func (c AVComposition) SetIsPlayable(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setPlayable:"), value)
}
// A Boolean value that indicates whether you can extract the asset’s media
// data using an asset reader.
//
// # Discussion
// 
// This property value is [true] if you can use [AVAssetReader] to extract the
// composition’s media data.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/isReadable
func (c AVComposition) IsReadable() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isReadable"))
	return rv
}
func (c AVComposition) SetIsReadable(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setReadable:"), value)
}
// A Boolean value that indicates whether you can export this asset using an
// export session.
//
// # Discussion
// 
// This property value is [true] if you can export the composition using
// [AVAssetExportSession].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/isExportable
func (c AVComposition) IsExportable() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isExportable"))
	return rv
}
func (c AVComposition) SetIsExportable(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setExportable:"), value)
}
// A Boolean value that indicates whether you can use the asset as a segment
// of a composition track.
//
// # Discussion
// 
// This property value is [true] if you can use the composition as a segment
// within an [AVCompositionTrack] object.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/isComposable
func (c AVComposition) IsComposable() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isComposable"))
	return rv
}
func (c AVComposition) SetIsComposable(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setComposable:"), value)
}
// A Boolean value that indicates whether the asset is compatible with AirPlay
// Video.
//
// # Discussion
// 
// This property value is [true] if you can play this composition’s content
// to an external AirPlay device, like an Apple TV.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/isCompatibleWithAirPlayVideo
func (c AVComposition) IsCompatibleWithAirPlayVideo() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isCompatibleWithAirPlayVideo"))
	return rv
}
func (c AVComposition) SetIsCompatibleWithAirPlayVideo(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setCompatibleWithAirPlayVideo:"), value)
}
// The asset’s rate preference for playing its media.
//
// # Discussion
// 
// This value is typically, but not always, 1.0.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/preferredRate
func (c AVComposition) PreferredRate() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("preferredRate"))
	return rv
}
func (c AVComposition) SetPreferredRate(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setPreferredRate:"), value)
}
// The asset’s volume preference for playing its audible media.
//
// # Discussion
// 
// This value is typically, but not always, 1.0.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/preferredVolume
func (c AVComposition) PreferredVolume() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("preferredVolume"))
	return rv
}
func (c AVComposition) SetPreferredVolume(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setPreferredVolume:"), value)
}
// The asset’s transform preference to apply to its visual content during
// presentation or processing.
//
// # Discussion
// 
// The value is typically, but not always, the identity transform.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/preferredTransform
func (c AVComposition) PreferredTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](c.ID, objc.Sel("preferredTransform"))
	return corefoundation.CGAffineTransform(rv)
}
func (c AVComposition) SetPreferredTransform(value corefoundation.CGAffineTransform) {
	objc.Send[struct{}](c.ID, objc.Sel("setPreferredTransform:"), value)
}
// The default media selections for this asset’s media selection groups.
//
// # Discussion
// 
// Provides an instance of [AVMediaSelection] with the default selections for
// each of the assets media selection groups.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/preferredMediaSelection
func (c AVComposition) PreferredMediaSelection() IAVMediaSelection {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("preferredMediaSelection"))
	return AVMediaSelectionFromID(objc.ID(rv))
}
func (c AVComposition) SetPreferredMediaSelection(value IAVMediaSelection) {
	objc.Send[struct{}](c.ID, objc.Sel("setPreferredMediaSelection:"), value)
}
// The array of available media selections for this asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/allMediaSelections
func (c AVComposition) AllMediaSelections() IAVMediaSelection {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("allMediaSelections"))
	return AVMediaSelectionFromID(objc.ID(rv))
}
func (c AVComposition) SetAllMediaSelections(value IAVMediaSelection) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllMediaSelections:"), value)
}
// An array of media characteristics for which a media selection option is
// available.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/availableMediaCharacteristicsWithMediaSelectionOptions
func (c AVComposition) AvailableMediaCharacteristicsWithMediaSelectionOptions() AVMediaCharacteristic {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("availableMediaCharacteristicsWithMediaSelectionOptions"))
	return AVMediaCharacteristic(foundation.NSStringFromID(rv).String())
}
func (c AVComposition) SetAvailableMediaCharacteristicsWithMediaSelectionOptions(value AVMediaCharacteristic) {
	objc.Send[struct{}](c.ID, objc.Sel("setAvailableMediaCharacteristicsWithMediaSelectionOptions:"), objc.String(string(value)))
}
// The locales of the asset’s chapter metadata.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/availableChapterLocales
func (c AVComposition) AvailableChapterLocales() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("availableChapterLocales"))
	return objectivec.Object{ID: rv}
}
func (c AVComposition) SetAvailableChapterLocales(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setAvailableChapterLocales:"), value)
}
// The options you used to create a composition.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/urlAssetInitializationOptions
func (c AVComposition) URLAssetInitializationOptions() foundation.INSDictionary {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("URLAssetInitializationOptions"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the asset contains protected
// content.
//
// # Discussion
// 
// Assets that contain protected content may not be playable without
// successful authorization, even if the value of its [isPlayable] property is
// [true].
//
// [isPlayable]: https://developer.apple.com/documentation/AVFoundation/AVAsset/isPlayable
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/hasProtectedContent
func (c AVComposition) HasProtectedContent() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasProtectedContent"))
	return rv
}
func (c AVComposition) SetHasProtectedContent(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setHasProtectedContent:"), value)
}
// A Boolean value that indicates whether you can extend the asset by
// fragments.
//
// # Discussion
// 
// For QuickTime movie files and MPEG-4 files, the value is [true] if an
// `mvex` box is present in the `moov` box. For those types, the `mvex` box
// signals the possible presence of later `moof` boxes.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/canContainFragments
func (c AVComposition) CanContainFragments() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("canContainFragments"))
	return rv
}
func (c AVComposition) SetCanContainFragments(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setCanContainFragments:"), value)
}
// A Boolean value that indicates whether at least one movie fragment extends
// the asset.
//
// # Discussion
// 
// For QuickTime movie files and MPEG-4 files, the value is [true] if
// [canContainFragments] is [true] and at least one `moof` box is present
// after the `moov` box.
//
// [canContainFragments]: https://developer.apple.com/documentation/AVFoundation/AVAsset/canContainFragments
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/containsFragments
func (c AVComposition) ContainsFragments() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("containsFragments"))
	return rv
}
func (c AVComposition) SetContainsFragments(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setContainsFragments:"), value)
}
// The total duration of fragments that currently exist, or may exist in the
// future.
//
// # Discussion
// 
// For QuickTime movie files and MPEG-4 files, the asset retrieves this value
// from the `mehd` box of the `mvex` box, if present. If no total fragment
// duration hint is available, the value of this property is [invalid].
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVComposition/overallDurationHint
func (c AVComposition) OverallDurationHint() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](c.ID, objc.Sel("overallDurationHint"))
	return coremedia.CMTime(rv)
}
func (c AVComposition) SetOverallDurationHint(value coremedia.CMTime) {
	objc.Send[struct{}](c.ID, objc.Sel("setOverallDurationHint:"), value)
}

