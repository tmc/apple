// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMutableMovie] class.
var (
	_AVMutableMovieClass     AVMutableMovieClass
	_AVMutableMovieClassOnce sync.Once
)

func getAVMutableMovieClass() AVMutableMovieClass {
	_AVMutableMovieClassOnce.Do(func() {
		_AVMutableMovieClass = AVMutableMovieClass{class: objc.GetClass("AVMutableMovie")}
	})
	return _AVMutableMovieClass
}

// GetAVMutableMovieClass returns the class object for AVMutableMovie.
func GetAVMutableMovieClass() AVMutableMovieClass {
	return getAVMutableMovieClass()
}

type AVMutableMovieClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMutableMovieClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMutableMovieClass) Alloc() AVMutableMovie {
	rv := objc.Send[AVMutableMovie](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A mutable object that represents an audiovisual container that conforms to
// the QuickTime movie file format or a related format like MPEG-4.
//
// # Overview
// 
// This class is a mutable subclass of [AVMovie] that provides methods that
// support movie editing. For example, you can use a mutable movie to copy
// media data from one track and paste it into another. You can also use this
// object to create track references from one track to another (for example,
// to set one track as a chapter track of another track). To perform editing
// operations on individual tracks, use the associated classes [AVMovieTrack]
// and [AVMutableMovieTrack].
// 
// You use movie objects only when operating on format-specific features of a
// QuickTime or ISO base media file. You typically don’t use these classes
// to open and play QuickTime movie files or ISO base media files. Instead,
// you use [AVURLAsset] and [AVPlayerItem].
// 
// When performing media insertions, a movie interleaves media data from
// tracks in the source asset to optimize the movie file for playback.
// However, performing a series of media insertions may result in a movie file
// that’s not optimally interleaved. You can optimize a movie file for
// playback by exporting it with an [AVAssetExportSession] object using the
// export preset [AVMutableMovie.AVAssetExportPresetPassthrough], and setting the
// [AVMutableMovie.ShouldOptimizeForNetworkUse] property value to [true].
//
// [AVMutableMovie.AVAssetExportPresetPassthrough]: https://developer.apple.com/documentation/AVFoundation/AVAssetExportPresetPassthrough
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Creating a movie
//
//   - [AVMutableMovie.InitWithURLOptionsError]: Creates a mutable movie object from a movie header stored in a QuickTime movie file of ISO base media file.
//   - [AVMutableMovie.InitWithDataOptionsError]: Creates a mutable movie object from a movie stored in a data object.
//   - [AVMutableMovie.InitWithSettingsFromMovieOptionsError]: Creates a mutable movie object without tracks.
//
// # Configuring a movie
//
//   - [AVMutableMovie.Modified]: A Boolean value that indicates whether the movie is in a modified state.
//   - [AVMutableMovie.SetModified]
//   - [AVMutableMovie.Timescale]: The time scale of the movie.
//   - [AVMutableMovie.SetTimescale]
//   - [AVMutableMovie.InterleavingPeriod]: A time period indicating the duration for interleaving runs of samples for each track.
//   - [AVMutableMovie.SetInterleavingPeriod]
//
// # Accessing tracks
//
//   - [AVMutableMovie.TrackWithTrackID]: Retrieves a track in the movie that contains the specified identifier.
//   - [AVMutableMovie.TracksWithMediaType]: Retrieves tracks in the movie that present media of the specified type.
//   - [AVMutableMovie.TracksWithMediaCharacteristic]: Retrieve tracks in the movie that present media of the specified characteristic.
//   - [AVMutableMovie.UnusedTrackID]: Returns an identifier that no other tracks in the asset use.
//
// # Accessing track groups
//
//   - [AVMutableMovie.TrackGroups]: The track groups an asset contains.
//   - [AVMutableMovie.SetTrackGroups]
//
// # Managing tracks
//
//   - [AVMutableMovie.MutableTrackCompatibleWithTrack]: Provides a reference to a track from a mutable movie into which you can insert any time range.
//   - [AVMutableMovie.AddMutableTrackWithMediaTypeCopySettingsFromTrackOptions]: Adds an empty track to the target movie.
//   - [AVMutableMovie.AddMutableTracksCopyingSettingsFromTracksOptions]: Adds one or more empty tracks to the target movie and copies the track settings from the source tracks.
//   - [AVMutableMovie.RemoveTrack]: Removes the specified track from the target movie.
//
// # Managing time ranges
//
//   - [AVMutableMovie.InsertEmptyTimeRange]: Adds an empty time range to a movie.
//   - [AVMutableMovie.InsertTimeRangeOfAssetAtTimeCopySampleDataError]: Inserts all of the tracks in a specified time range of an asset into a movie.
//   - [AVMutableMovie.ScaleTimeRangeToDuration]: Changes the duration of a time range in a movie.
//   - [AVMutableMovie.RemoveTimeRange]: Removes the specified time range from a movie.
//
// # Accessing duration and timing
//
//   - [AVMutableMovie.Duration]: A time value that indicates the asset’s duration.
//   - [AVMutableMovie.SetDuration]
//   - [AVMutableMovie.ProvidesPreciseDurationAndTiming]: A Boolean value that indicates whether the asset provides precise duration and timing.
//   - [AVMutableMovie.SetProvidesPreciseDurationAndTiming]
//   - [AVMutableMovie.MinimumTimeOffsetFromLive]: A time value that indicates how closely playback follows the latest live stream content.
//   - [AVMutableMovie.SetMinimumTimeOffsetFromLive]
//
// # Accessing metadata
//
//   - [AVMutableMovie.Metadata]: An array of metadata items for all metadata identifiers for which a value is available.
//   - [AVMutableMovie.SetMetadata]
//   - [AVMutableMovie.CommonMetadata]: The metadata items an asset contains for common metadata identifiers that provide a value.
//   - [AVMutableMovie.SetCommonMetadata]
//   - [AVMutableMovie.AvailableMetadataFormats]: The metadata formats this asset contains.
//   - [AVMutableMovie.SetAvailableMetadataFormats]
//   - [AVMutableMovie.MetadataForFormat]: Returns an array of metadata items from the container with the specified format.
//   - [AVMutableMovie.CreationDate]: A metadata item that indicates the asset’s creation date.
//   - [AVMutableMovie.SetCreationDate]
//   - [AVMutableMovie.Lyrics]: The lyrics of the asset in a language suitable for the current locale.
//   - [AVMutableMovie.SetLyrics]
//
// # Determining suitability
//
//   - [AVMutableMovie.IsPlayable]: A Boolean value that indicates whether the asset has playable content.
//   - [AVMutableMovie.SetIsPlayable]
//   - [AVMutableMovie.IsReadable]: A Boolean value that indicates whether you can extract the asset’s media data using an asset reader.
//   - [AVMutableMovie.SetIsReadable]
//   - [AVMutableMovie.IsExportable]: A Boolean value that indicates whether you can export this asset using an export session.
//   - [AVMutableMovie.SetIsExportable]
//   - [AVMutableMovie.IsComposable]: A Boolean value that indicates whether you can use the asset as a segment of a composition track.
//   - [AVMutableMovie.SetIsComposable]
//   - [AVMutableMovie.IsCompatibleWithAirPlayVideo]: A Boolean value that indicates whether the asset is compatible with AirPlay Video.
//   - [AVMutableMovie.SetIsCompatibleWithAirPlayVideo]
//
// # Inspecting preferences
//
//   - [AVMutableMovie.PreferredRate]: The asset’s rate preference for playing its media.
//   - [AVMutableMovie.SetPreferredRate]
//   - [AVMutableMovie.PreferredVolume]: The asset’s volume preference for playing its audible media.
//   - [AVMutableMovie.SetPreferredVolume]
//   - [AVMutableMovie.PreferredTransform]: The asset’s transform preference to apply to its visual content during presentation or processing.
//   - [AVMutableMovie.SetPreferredTransform]
//   - [AVMutableMovie.PreferredMediaSelection]: The default media selections for this asset’s media selection groups.
//   - [AVMutableMovie.SetPreferredMediaSelection]
//
// # Accessing media selections
//
//   - [AVMutableMovie.AllMediaSelections]: The array of available media selections for this asset.
//   - [AVMutableMovie.SetAllMediaSelections]
//   - [AVMutableMovie.AvailableMediaCharacteristicsWithMediaSelectionOptions]: An array of media characteristics for which a media selection option is available.
//   - [AVMutableMovie.SetAvailableMediaCharacteristicsWithMediaSelectionOptions]
//   - [AVMutableMovie.MediaSelectionGroupForMediaCharacteristic]: Returns a media selection group that contains one or more options with the specified media characteristic.
//
// # Accessing chapter metadata
//
//   - [AVMutableMovie.AvailableChapterLocales]: The locales of the asset’s chapter metadata.
//   - [AVMutableMovie.SetAvailableChapterLocales]
//   - [AVMutableMovie.ChapterMetadataGroupsBestMatchingPreferredLanguages]: Returns an array of chapters with a locale that best matches the list of preferred languages.
//   - [AVMutableMovie.ChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeys]: Returns an array of chapters that contain the specified title locale and common keys.
//
// # Determining content protections
//
//   - [AVMutableMovie.HasProtectedContent]: A Boolean value that indicates whether the asset contains protected content.
//   - [AVMutableMovie.SetHasProtectedContent]
//
// # Determining fragment support
//
//   - [AVMutableMovie.CanContainFragments]: A Boolean value that indicates whether you can extend the asset by fragments.
//   - [AVMutableMovie.SetCanContainFragments]
//   - [AVMutableMovie.ContainsFragments]: A Boolean value that indicates whether at least one movie fragment extends the asset.
//   - [AVMutableMovie.SetContainsFragments]
//   - [AVMutableMovie.OverallDurationHint]: The total duration of fragments that currently exist, or may exist in the future.
//   - [AVMutableMovie.SetOverallDurationHint]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie
type AVMutableMovie struct {
	AVMovie
}

// AVMutableMovieFromID constructs a [AVMutableMovie] from an objc.ID.
//
// A mutable object that represents an audiovisual container that conforms to
// the QuickTime movie file format or a related format like MPEG-4.
func AVMutableMovieFromID(id objc.ID) AVMutableMovie {
	return AVMutableMovie{AVMovie: AVMovieFromID(id)}
}
// NOTE: AVMutableMovie adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMutableMovie] class.
//
// # Creating a movie
//
//   - [IAVMutableMovie.InitWithURLOptionsError]: Creates a mutable movie object from a movie header stored in a QuickTime movie file of ISO base media file.
//   - [IAVMutableMovie.InitWithDataOptionsError]: Creates a mutable movie object from a movie stored in a data object.
//   - [IAVMutableMovie.InitWithSettingsFromMovieOptionsError]: Creates a mutable movie object without tracks.
//
// # Configuring a movie
//
//   - [IAVMutableMovie.Modified]: A Boolean value that indicates whether the movie is in a modified state.
//   - [IAVMutableMovie.SetModified]
//   - [IAVMutableMovie.Timescale]: The time scale of the movie.
//   - [IAVMutableMovie.SetTimescale]
//   - [IAVMutableMovie.InterleavingPeriod]: A time period indicating the duration for interleaving runs of samples for each track.
//   - [IAVMutableMovie.SetInterleavingPeriod]
//
// # Accessing tracks
//
//   - [IAVMutableMovie.TrackWithTrackID]: Retrieves a track in the movie that contains the specified identifier.
//   - [IAVMutableMovie.TracksWithMediaType]: Retrieves tracks in the movie that present media of the specified type.
//   - [IAVMutableMovie.TracksWithMediaCharacteristic]: Retrieve tracks in the movie that present media of the specified characteristic.
//   - [IAVMutableMovie.UnusedTrackID]: Returns an identifier that no other tracks in the asset use.
//
// # Accessing track groups
//
//   - [IAVMutableMovie.TrackGroups]: The track groups an asset contains.
//   - [IAVMutableMovie.SetTrackGroups]
//
// # Managing tracks
//
//   - [IAVMutableMovie.MutableTrackCompatibleWithTrack]: Provides a reference to a track from a mutable movie into which you can insert any time range.
//   - [IAVMutableMovie.AddMutableTrackWithMediaTypeCopySettingsFromTrackOptions]: Adds an empty track to the target movie.
//   - [IAVMutableMovie.AddMutableTracksCopyingSettingsFromTracksOptions]: Adds one or more empty tracks to the target movie and copies the track settings from the source tracks.
//   - [IAVMutableMovie.RemoveTrack]: Removes the specified track from the target movie.
//
// # Managing time ranges
//
//   - [IAVMutableMovie.InsertEmptyTimeRange]: Adds an empty time range to a movie.
//   - [IAVMutableMovie.InsertTimeRangeOfAssetAtTimeCopySampleDataError]: Inserts all of the tracks in a specified time range of an asset into a movie.
//   - [IAVMutableMovie.ScaleTimeRangeToDuration]: Changes the duration of a time range in a movie.
//   - [IAVMutableMovie.RemoveTimeRange]: Removes the specified time range from a movie.
//
// # Accessing duration and timing
//
//   - [IAVMutableMovie.Duration]: A time value that indicates the asset’s duration.
//   - [IAVMutableMovie.SetDuration]
//   - [IAVMutableMovie.ProvidesPreciseDurationAndTiming]: A Boolean value that indicates whether the asset provides precise duration and timing.
//   - [IAVMutableMovie.SetProvidesPreciseDurationAndTiming]
//   - [IAVMutableMovie.MinimumTimeOffsetFromLive]: A time value that indicates how closely playback follows the latest live stream content.
//   - [IAVMutableMovie.SetMinimumTimeOffsetFromLive]
//
// # Accessing metadata
//
//   - [IAVMutableMovie.Metadata]: An array of metadata items for all metadata identifiers for which a value is available.
//   - [IAVMutableMovie.SetMetadata]
//   - [IAVMutableMovie.CommonMetadata]: The metadata items an asset contains for common metadata identifiers that provide a value.
//   - [IAVMutableMovie.SetCommonMetadata]
//   - [IAVMutableMovie.AvailableMetadataFormats]: The metadata formats this asset contains.
//   - [IAVMutableMovie.SetAvailableMetadataFormats]
//   - [IAVMutableMovie.MetadataForFormat]: Returns an array of metadata items from the container with the specified format.
//   - [IAVMutableMovie.CreationDate]: A metadata item that indicates the asset’s creation date.
//   - [IAVMutableMovie.SetCreationDate]
//   - [IAVMutableMovie.Lyrics]: The lyrics of the asset in a language suitable for the current locale.
//   - [IAVMutableMovie.SetLyrics]
//
// # Determining suitability
//
//   - [IAVMutableMovie.IsPlayable]: A Boolean value that indicates whether the asset has playable content.
//   - [IAVMutableMovie.SetIsPlayable]
//   - [IAVMutableMovie.IsReadable]: A Boolean value that indicates whether you can extract the asset’s media data using an asset reader.
//   - [IAVMutableMovie.SetIsReadable]
//   - [IAVMutableMovie.IsExportable]: A Boolean value that indicates whether you can export this asset using an export session.
//   - [IAVMutableMovie.SetIsExportable]
//   - [IAVMutableMovie.IsComposable]: A Boolean value that indicates whether you can use the asset as a segment of a composition track.
//   - [IAVMutableMovie.SetIsComposable]
//   - [IAVMutableMovie.IsCompatibleWithAirPlayVideo]: A Boolean value that indicates whether the asset is compatible with AirPlay Video.
//   - [IAVMutableMovie.SetIsCompatibleWithAirPlayVideo]
//
// # Inspecting preferences
//
//   - [IAVMutableMovie.PreferredRate]: The asset’s rate preference for playing its media.
//   - [IAVMutableMovie.SetPreferredRate]
//   - [IAVMutableMovie.PreferredVolume]: The asset’s volume preference for playing its audible media.
//   - [IAVMutableMovie.SetPreferredVolume]
//   - [IAVMutableMovie.PreferredTransform]: The asset’s transform preference to apply to its visual content during presentation or processing.
//   - [IAVMutableMovie.SetPreferredTransform]
//   - [IAVMutableMovie.PreferredMediaSelection]: The default media selections for this asset’s media selection groups.
//   - [IAVMutableMovie.SetPreferredMediaSelection]
//
// # Accessing media selections
//
//   - [IAVMutableMovie.AllMediaSelections]: The array of available media selections for this asset.
//   - [IAVMutableMovie.SetAllMediaSelections]
//   - [IAVMutableMovie.AvailableMediaCharacteristicsWithMediaSelectionOptions]: An array of media characteristics for which a media selection option is available.
//   - [IAVMutableMovie.SetAvailableMediaCharacteristicsWithMediaSelectionOptions]
//   - [IAVMutableMovie.MediaSelectionGroupForMediaCharacteristic]: Returns a media selection group that contains one or more options with the specified media characteristic.
//
// # Accessing chapter metadata
//
//   - [IAVMutableMovie.AvailableChapterLocales]: The locales of the asset’s chapter metadata.
//   - [IAVMutableMovie.SetAvailableChapterLocales]
//   - [IAVMutableMovie.ChapterMetadataGroupsBestMatchingPreferredLanguages]: Returns an array of chapters with a locale that best matches the list of preferred languages.
//   - [IAVMutableMovie.ChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeys]: Returns an array of chapters that contain the specified title locale and common keys.
//
// # Determining content protections
//
//   - [IAVMutableMovie.HasProtectedContent]: A Boolean value that indicates whether the asset contains protected content.
//   - [IAVMutableMovie.SetHasProtectedContent]
//
// # Determining fragment support
//
//   - [IAVMutableMovie.CanContainFragments]: A Boolean value that indicates whether you can extend the asset by fragments.
//   - [IAVMutableMovie.SetCanContainFragments]
//   - [IAVMutableMovie.ContainsFragments]: A Boolean value that indicates whether at least one movie fragment extends the asset.
//   - [IAVMutableMovie.SetContainsFragments]
//   - [IAVMutableMovie.OverallDurationHint]: The total duration of fragments that currently exist, or may exist in the future.
//   - [IAVMutableMovie.SetOverallDurationHint]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie
type IAVMutableMovie interface {
	IAVMovie

	// Topic: Creating a movie

	// Creates a mutable movie object from a movie header stored in a QuickTime movie file of ISO base media file.
	InitWithURLOptionsError(URL foundation.INSURL, options foundation.INSDictionary) (AVMutableMovie, error)
	// Creates a mutable movie object from a movie stored in a data object.
	InitWithDataOptionsError(data foundation.INSData, options foundation.INSDictionary) (AVMutableMovie, error)
	// Creates a mutable movie object without tracks.
	InitWithSettingsFromMovieOptionsError(movie IAVMovie, options foundation.INSDictionary) (AVMutableMovie, error)

	// Topic: Configuring a movie

	// A Boolean value that indicates whether the movie is in a modified state.
	Modified() bool
	SetModified(value bool)
	// The time scale of the movie.
	Timescale() int32
	SetTimescale(value int32)
	// A time period indicating the duration for interleaving runs of samples for each track.
	InterleavingPeriod() coremedia.CMTime
	SetInterleavingPeriod(value coremedia.CMTime)

	// Topic: Accessing tracks

	// Retrieves a track in the movie that contains the specified identifier.
	TrackWithTrackID(trackID int32) IAVMutableMovieTrack
	// Retrieves tracks in the movie that present media of the specified type.
	TracksWithMediaType(mediaType AVMediaType) []AVMutableMovieTrack
	// Retrieve tracks in the movie that present media of the specified characteristic.
	TracksWithMediaCharacteristic(mediaCharacteristic AVMediaCharacteristic) []AVMutableMovieTrack
	// Returns an identifier that no other tracks in the asset use.
	UnusedTrackID() int32

	// Topic: Accessing track groups

	// The track groups an asset contains.
	TrackGroups() IAVAssetTrackGroup
	SetTrackGroups(value IAVAssetTrackGroup)

	// Topic: Managing tracks

	// Provides a reference to a track from a mutable movie into which you can insert any time range.
	MutableTrackCompatibleWithTrack(track IAVAssetTrack) IAVMutableMovieTrack
	// Adds an empty track to the target movie.
	AddMutableTrackWithMediaTypeCopySettingsFromTrackOptions(mediaType AVMediaType, track IAVAssetTrack, options foundation.INSDictionary) IAVMutableMovieTrack
	// Adds one or more empty tracks to the target movie and copies the track settings from the source tracks.
	AddMutableTracksCopyingSettingsFromTracksOptions(existingTracks []AVAssetTrack, options foundation.INSDictionary) []AVMutableMovieTrack
	// Removes the specified track from the target movie.
	RemoveTrack(track IAVMovieTrack)

	// Topic: Managing time ranges

	// Adds an empty time range to a movie.
	InsertEmptyTimeRange(timeRange coremedia.CMTimeRange)
	// Inserts all of the tracks in a specified time range of an asset into a movie.
	InsertTimeRangeOfAssetAtTimeCopySampleDataError(timeRange coremedia.CMTimeRange, asset IAVAsset, startTime coremedia.CMTime, copySampleData bool) (bool, error)
	// Changes the duration of a time range in a movie.
	ScaleTimeRangeToDuration(timeRange coremedia.CMTimeRange, duration coremedia.CMTime)
	// Removes the specified time range from a movie.
	RemoveTimeRange(timeRange coremedia.CMTimeRange)

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
	Metadata() []AVMetadataItem
	SetMetadata(value []AVMetadataItem)
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

	// A preset to export the asset in its current format, unless otherwise prohibited.
	AVAssetExportPresetPassthrough() string
	// A Boolean value that indicates whether to optimize the movie for network use.
	ShouldOptimizeForNetworkUse() bool
	SetShouldOptimizeForNetworkUse(value bool)
}

// Init initializes the instance.
func (m AVMutableMovie) Init() AVMutableMovie {
	rv := objc.Send[AVMutableMovie](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMutableMovie) Autorelease() AVMutableMovie {
	rv := objc.Send[AVMutableMovie](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMutableMovie creates a new AVMutableMovie instance.
func NewAVMutableMovie() AVMutableMovie {
	class := getAVMutableMovieClass()
	rv := objc.Send[AVMutableMovie](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an asset that models the media at the specified URL.
//
// URL: A URL to a local, remote, or HTTP Live Streaming media resource.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVAsset/init(url:)
func NewMutableMovieAssetWithURL(URL foundation.INSURL) AVMutableMovie {
	rv := objc.Send[objc.ID](objc.ID(getAVMutableMovieClass().class), objc.Sel("assetWithURL:"), URL)
	return AVMutableMovieFromID(rv)
}

// Creates a movie object from a movie file’s data.
//
// data: A data object that contains a movie header.
//
// options: A dictionary of options to use to initialize the movie.
//
// # Discussion
// 
// Use this method to create movies from movie headers that aren’t stored in
// files, which can include movies that the pasteboard contains.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovie/init(data:options:)
func NewMutableMovieWithDataOptions(data foundation.INSData, options foundation.INSDictionary) AVMutableMovie {
	instance := getAVMutableMovieClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:options:"), data, options)
	return AVMutableMovieFromID(rv)
}

// Creates a mutable movie object from a movie stored in a data object.
//
// data: An [NSData] object that contains a movie header.
//
// options: A dictionary that contains key for specifying the movie object
// initialization. Currently, no keys are defined.
//
// # Return Value
// 
// An [AVMutableMovie] object.
//
// # Discussion
// 
// On initialization, the [DefaultMediaDataStorage] property and any
// associated [MediaDataStorage] properties are set to `nil`. To create an
// [AVMutableMovie] from a file and then append sample buffers to any of its
// tracks, you must first set one of these properties to indicate where the
// sample data should be written.
// 
// Use this method to create movies from movie headers that are not stored in
// files, which can include movies on the pasteboard.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/init(data:options:error:)
func NewMutableMovieWithDataOptionsError(data foundation.INSData, options foundation.INSDictionary) (AVMutableMovie, error) {
	var errorPtr objc.ID
	instance := getAVMutableMovieClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithData:options:error:"), data, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVMutableMovie{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVMutableMovieFromID(rv), nil
}

// Creates a mutable movie object without tracks.
//
// movie: A movie object that contains settings from an existing movie.
//
// options: A dictionary that contains key for specifying the movie object
// initialization. Currently, no keys are defined.
//
// # Return Value
// 
// An [AVMutableMovie] object.
//
// # Discussion
// 
// On initialization, the [DefaultMediaDataStorage] property and any
// associated [MediaDataStorage] properties are set to `nil`. To create an
// [AVMutableMovie] from a file and then append sample buffers to any of its
// tracks, you must first set one of these properties to indicate where the
// sample data should be written.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/init(settingsFrom:options:)
func NewMutableMovieWithSettingsFromMovieOptionsError(movie IAVMovie, options foundation.INSDictionary) (AVMutableMovie, error) {
	var errorPtr objc.ID
	instance := getAVMutableMovieClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSettingsFromMovie:options:error:"), movie, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVMutableMovie{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVMutableMovieFromID(rv), nil
}

// Creates a movie object from a movie header stored in a QuickTime movie file
// of ISO base media file.
//
// URL: A URL that points to a file containing a movie header.
//
// options: A dictionary of options to use to initialize the movie.
//
// # Discussion
// 
// Upon creation, the values of the [DefaultMediaDataStorage] property and any
// associated [MediaDataStorage] properties are `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMovie/init(url:options:)
func NewMutableMovieWithURLOptions(URL foundation.INSURL, options foundation.INSDictionary) AVMutableMovie {
	instance := getAVMutableMovieClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:options:"), URL, options)
	return AVMutableMovieFromID(rv)
}

// Creates a mutable movie object from a movie header stored in a QuickTime
// movie file of ISO base media file.
//
// URL: The URL that points to a file containing a movie header.
//
// options: A dictionary that contains key for specifying the movie object
// initialization. Currently, no keys are defined.
//
// # Return Value
// 
// An [AVMutableMovie] object.
//
// # Discussion
// 
// On initialization, the [DefaultMediaDataStorage] property and any
// associated [MediaDataStorage] properties are set to `nil`. To create an
// [AVMutableMovie] from a file and then append sample buffers to any of its
// tracks, you must first set one of these properties to indicate where the
// sample data should be written.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/init(url:options:error:)
func NewMutableMovieWithURLOptionsError(URL foundation.INSURL, options foundation.INSDictionary) (AVMutableMovie, error) {
	var errorPtr objc.ID
	instance := getAVMutableMovieClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithURL:options:error:"), URL, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVMutableMovie{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVMutableMovieFromID(rv), nil
}

// Creates a mutable movie object from a movie header stored in a QuickTime
// movie file of ISO base media file.
//
// URL: The URL that points to a file containing a movie header.
//
// options: A dictionary that contains key for specifying the movie object
// initialization. Currently, no keys are defined.
//
// # Return Value
// 
// An [AVMutableMovie] object.
//
// # Discussion
// 
// On initialization, the [DefaultMediaDataStorage] property and any
// associated [MediaDataStorage] properties are set to `nil`. To create an
// [AVMutableMovie] from a file and then append sample buffers to any of its
// tracks, you must first set one of these properties to indicate where the
// sample data should be written.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/init(url:options:error:)
func (m AVMutableMovie) InitWithURLOptionsError(URL foundation.INSURL, options foundation.INSDictionary) (AVMutableMovie, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithURL:options:error:"), URL, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVMutableMovie{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVMutableMovieFromID(rv), nil

}
// Creates a mutable movie object from a movie stored in a data object.
//
// data: An [NSData] object that contains a movie header.
//
// options: A dictionary that contains key for specifying the movie object
// initialization. Currently, no keys are defined.
//
// # Return Value
// 
// An [AVMutableMovie] object.
//
// # Discussion
// 
// On initialization, the [DefaultMediaDataStorage] property and any
// associated [MediaDataStorage] properties are set to `nil`. To create an
// [AVMutableMovie] from a file and then append sample buffers to any of its
// tracks, you must first set one of these properties to indicate where the
// sample data should be written.
// 
// Use this method to create movies from movie headers that are not stored in
// files, which can include movies on the pasteboard.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/init(data:options:error:)
func (m AVMutableMovie) InitWithDataOptionsError(data foundation.INSData, options foundation.INSDictionary) (AVMutableMovie, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithData:options:error:"), data, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVMutableMovie{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVMutableMovieFromID(rv), nil

}
// Creates a mutable movie object without tracks.
//
// movie: A movie object that contains settings from an existing movie.
//
// options: A dictionary that contains key for specifying the movie object
// initialization. Currently, no keys are defined.
//
// # Return Value
// 
// An [AVMutableMovie] object.
//
// # Discussion
// 
// On initialization, the [DefaultMediaDataStorage] property and any
// associated [MediaDataStorage] properties are set to `nil`. To create an
// [AVMutableMovie] from a file and then append sample buffers to any of its
// tracks, you must first set one of these properties to indicate where the
// sample data should be written.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/init(settingsFrom:options:)
func (m AVMutableMovie) InitWithSettingsFromMovieOptionsError(movie IAVMovie, options foundation.INSDictionary) (AVMutableMovie, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithSettingsFromMovie:options:error:"), movie, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVMutableMovie{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVMutableMovieFromID(rv), nil

}
// Retrieves a track in the movie that contains the specified identifier.
//
// trackID: The track identifier for the requested track.
//
// # Return Value
// 
// A movie track, or `nil` if there is no track with the identifier.
//
// # Discussion
// 
// Apple discourages using this method in iOS 15, tvOS 15, macOS 12, and
// watchOS 8 or later. Load a track asynchronously using
// [LoadTrackWithTrackIDCompletionHandler] instead.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/track(withTrackID:)
func (m AVMutableMovie) TrackWithTrackID(trackID int32) IAVMutableMovieTrack {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("trackWithTrackID:"), trackID)
	return AVMutableMovieTrackFromID(rv)
}
// Retrieves tracks in the movie that present media of the specified type.
//
// mediaType: The media type of the tracks to return.
//
// # Return Value
// 
// An array of tracks, which is empty if there are no tracks with the media
// type.
//
// # Discussion
// 
// Apple discourages using this method in iOS 15, tvOS 15, macOS 12, and
// watchOS 8 or later. Load tracks asynchronously using
// [LoadTracksWithMediaTypeCompletionHandler] instead.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/tracks(withMediaType:)
func (m AVMutableMovie) TracksWithMediaType(mediaType AVMediaType) []AVMutableMovieTrack {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("tracksWithMediaType:"), objc.String(string(mediaType)))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMutableMovieTrack {
		return AVMutableMovieTrackFromID(id)
	})
}
// Retrieve tracks in the movie that present media of the specified
// characteristic.
//
// mediaCharacteristic: The media characteristic of the tracks to return.
//
// # Return Value
// 
// An array of tracks, which is empty if there are no tracks with the media
// characteristic.
//
// # Discussion
// 
// Apple discourages using this method in iOS 15, tvOS 15, macOS 12, and
// watchOS 8 or later. Load tracks asynchronously using
// [LoadTracksWithMediaCharacteristicCompletionHandler] instead.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/tracks(withMediaCharacteristic:)
func (m AVMutableMovie) TracksWithMediaCharacteristic(mediaCharacteristic AVMediaCharacteristic) []AVMutableMovieTrack {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("tracksWithMediaCharacteristic:"), objc.String(string(mediaCharacteristic)))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMutableMovieTrack {
		return AVMutableMovieTrackFromID(id)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/unusedTrackID()
func (m AVMutableMovie) UnusedTrackID() int32 {
	rv := objc.Send[int32](m.ID, objc.Sel("unusedTrackID"))
	return rv
}
// Provides a reference to a track from a mutable movie into which you can
// insert any time range.
//
// track: The [AVAssetTrack] containing the desired time range.
//
// # Return Value
// 
// An [AVMutableMovieTrack] object that can accommodate the time range
// insertion. Returns nil when no track is available.
//
// # Discussion
// 
// Keep the number of tracks in a movie to a minimum, corresponding to the
// number of tracks for which media data must be presented in parallel. If
// media data of the same type is presented serially, even from multiple
// assets, a single track of that media type should be used. This method can
// help the client to identify an existing target track for an insertion.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/mutableTrack(compatibleWith:)
func (m AVMutableMovie) MutableTrackCompatibleWithTrack(track IAVAssetTrack) IAVMutableMovieTrack {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mutableTrackCompatibleWithTrack:"), track)
	return AVMutableMovieTrackFromID(rv)
}
// Adds an empty track to the target movie.
//
// mediaType: The media type for the new track.
//
// track: An [AVAssetTrack] containing the desired track settings to be transferred.
// Set to `nil` to create a track with default settings.
//
// options: A dictionary that contains key for specifying the movie object
// initialization. Currently, no keys are defined.
//
// # Return Value
// 
// An [AVMutableMovieTrack] object.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/addMutableTrack(withMediaType:copySettingsFrom:options:)
func (m AVMutableMovie) AddMutableTrackWithMediaTypeCopySettingsFromTrackOptions(mediaType AVMediaType, track IAVAssetTrack, options foundation.INSDictionary) IAVMutableMovieTrack {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addMutableTrackWithMediaType:copySettingsFromTrack:options:"), objc.String(string(mediaType)), track, options)
	return AVMutableMovieTrackFromID(rv)
}
// Adds one or more empty tracks to the target movie and copies the track
// settings from the source tracks.
//
// existingTracks: An array of asset tracks to be added.
//
// options: A dictionary that contains key for specifying the movie object
// initialization. Currently, no keys are defined.
//
// # Return Value
// 
// An array of [AVMutableMovieTrack] objects. The index of a track in this
// array is the same as the index of its source track in the `existingTracks`
// array.
//
// # Discussion
// 
// Properties involving pairs of tracks,such as track references, are copied
// from the source tracks to the target tracks.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/addMutableTracksCopyingSettings(from:options:)
func (m AVMutableMovie) AddMutableTracksCopyingSettingsFromTracksOptions(existingTracks []AVAssetTrack, options foundation.INSDictionary) []AVMutableMovieTrack {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("addMutableTracksCopyingSettingsFromTracks:options:"), objectivec.IObjectSliceToNSArray(existingTracks), options)
	return objc.ConvertSlice(rv, func(id objc.ID) AVMutableMovieTrack {
		return AVMutableMovieTrackFromID(id)
	})
}
// Removes the specified track from the target movie.
//
// track: The movie to be removed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/removeTrack(_:)
func (m AVMutableMovie) RemoveTrack(track IAVMovieTrack) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeTrack:"), track)
}
// Adds an empty time range to a movie.
//
// timeRange: The time range to be made empty.
//
// # Discussion
// 
// You can’t add empty time ranges to the end of a movie.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/insertEmptyTimeRange(_:)
func (m AVMutableMovie) InsertEmptyTimeRange(timeRange coremedia.CMTimeRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("insertEmptyTimeRange:"), timeRange)
}
// Inserts all of the tracks in a specified time range of an asset into a
// movie.
//
// timeRange: The time range of the asset to be inserted.
//
// asset: An [AVAsset] object indicating the source of the inserted media. This value
// can’t be `nil`.
//
// startTime: The time in the target movie at which the media is to be inserted.
//
// copySampleData: A Boolean value that indicates whether sample data is to be copied from the
// source to the destination during edits.
//
// # Discussion
// 
// This method may add new tracks to the target movie to ensure that all
// tracks of the asset are represented in the inserted time range.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/insertTimeRange(_:of:at:copySampleData:)
func (m AVMutableMovie) InsertTimeRangeOfAssetAtTimeCopySampleDataError(timeRange coremedia.CMTimeRange, asset IAVAsset, startTime coremedia.CMTime, copySampleData bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("insertTimeRange:ofAsset:atTime:copySampleData:error:"), timeRange, asset, startTime, copySampleData, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("insertTimeRange:ofAsset:atTime:copySampleData:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Changes the duration of a time range in a movie.
//
// timeRange: The time range to be changed.
//
// duration: The new duration for the time range.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/scale(_:toDuration:)
func (m AVMutableMovie) ScaleTimeRangeToDuration(timeRange coremedia.CMTimeRange, duration coremedia.CMTime) {
	objc.Send[objc.ID](m.ID, objc.Sel("scaleTimeRange:toDuration:"), timeRange, duration)
}
// Removes the specified time range from a movie.
//
// timeRange: The time range to be removed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/removeTimeRange(_:)
func (m AVMutableMovie) RemoveTimeRange(timeRange coremedia.CMTimeRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeTimeRange:"), timeRange)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/metadata(forFormat:)
func (m AVMutableMovie) MetadataForFormat(format AVMetadataFormat) []AVMetadataItem {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("metadataForFormat:"), objc.String(string(format)))
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/mediaSelectionGroup(forMediaCharacteristic:)
func (m AVMutableMovie) MediaSelectionGroupForMediaCharacteristic(mediaCharacteristic AVMediaCharacteristic) IAVMediaSelectionGroup {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mediaSelectionGroupForMediaCharacteristic:"), objc.String(string(mediaCharacteristic)))
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/chapterMetadataGroups(bestMatchingPreferredLanguages:)
func (m AVMutableMovie) ChapterMetadataGroupsBestMatchingPreferredLanguages(preferredLanguages []string) []AVTimedMetadataGroup {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("chapterMetadataGroupsBestMatchingPreferredLanguages:"), objectivec.StringSliceToNSArray(preferredLanguages))
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/chapterMetadataGroups(withTitleLocale:containingItemsWithCommonKeys:)
func (m AVMutableMovie) ChapterMetadataGroupsWithTitleLocaleContainingItemsWithCommonKeys(locale foundation.NSLocale, commonKeys []string) []AVTimedMetadataGroup {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("chapterMetadataGroupsWithTitleLocale:containingItemsWithCommonKeys:"), locale, objectivec.StringSliceToNSArray(commonKeys))
	return objc.ConvertSlice(rv, func(id objc.ID) AVTimedMetadataGroup {
		return AVTimedMetadataGroupFromID(id)
	})
}

// Returns a new mutable movie object from a movie stored in a data object.
//
// data: An [NSData] object that contains a movie header.
//
// options: A dictionary that contains key for specifying the movie object
// initialization. Currently, no keys are defined.
//
// outError: A description of the error that occurred. Default value is `nil`.
//
// # Return Value
// 
// An [AVMutableMovie] object.
//
// # Discussion
// 
// On initialization, the [DefaultMediaDataStorage] property and any
// associated [MediaDataStorage] properties are set to `nil`. To create an
// [AVMutableMovie] from a file and then append sample buffers to any of its
// tracks, you must first set one of these properties to indicate where the
// sample data should be written.
// 
// Use this method to create movies from movie headers that are not stored in
// files, which can include movies on the pasteboard.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/movieWithData:options:error:
func (_AVMutableMovieClass AVMutableMovieClass) MovieWithDataOptionsError(data foundation.INSData, options foundation.INSDictionary) (AVMutableMovie, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_AVMutableMovieClass.class), objc.Sel("movieWithData:options:error:"), data, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVMutableMovie{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVMutableMovieFromID(rv), nil

}
// Returns a new mutable movie object without tracks.
//
// movie: An [AVMovie] object containing settings from an existing movie.
//
// options: A dictionary that contains key for specifying the movie object
// initialization. Currently, no keys are defined.
//
// outError: A description of the error that occurred. Default value is `nil`.
//
// # Return Value
// 
// An [AVMutableMovie] object.
//
// # Discussion
// 
// On initialization, the [DefaultMediaDataStorage] property and any
// associated [MediaDataStorage] properties are set to `nil`. To create an
// [AVMutableMovie] from a file and then append sample buffers to any of its
// tracks, you must first set one of these properties to indicate where the
// sample data should be written.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/movieWithSettingsFromMovie:options:error:
func (_AVMutableMovieClass AVMutableMovieClass) MovieWithSettingsFromMovieOptionsError(movie IAVMovie, options foundation.INSDictionary) (AVMutableMovie, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_AVMutableMovieClass.class), objc.Sel("movieWithSettingsFromMovie:options:error:"), movie, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVMutableMovie{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVMutableMovieFromID(rv), nil

}
// Returns a new mutable movie object from a movie header stored in a
// QuickTime movie file of ISO base media file.
//
// URL: The URL that points to a file containing a movie header.
//
// options: A dictionary that contains key for specifying the movie object
// initialization. Currently, no keys are defined.
//
// outError: A description of the error that occurred. Default value is `nil`.
//
// # Return Value
// 
// An [AVMutableMovie] object.
//
// # Discussion
// 
// On initialization, the [DefaultMediaDataStorage] property and any
// associated [MediaDataStorage] properties are set to `nil`. To create an
// [AVMutableMovie] from a file and then append sample buffers to any of its
// tracks, you must first set one of these properties to indicate where the
// sample data should be written.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/movieWithURL:options:error:
func (_AVMutableMovieClass AVMutableMovieClass) MovieWithURLOptionsError(URL foundation.INSURL, options foundation.INSDictionary) (AVMutableMovie, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_AVMutableMovieClass.class), objc.Sel("movieWithURL:options:error:"), URL, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVMutableMovie{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVMutableMovieFromID(rv), nil

}

// A Boolean value that indicates whether the movie is in a modified state.
//
// # Discussion
// 
// The value is true if you’ve modified the movie since you created it,
// saved it, or had its modified state cleared.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/isModified
func (m AVMutableMovie) Modified() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isModified"))
	return rv
}
func (m AVMutableMovie) SetModified(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setModified:"), value)
}
// The time scale of the movie.
//
// # Discussion
// 
// The default movie time scale is `600`. In certain cases, you may want to
// set this to a different value. For example, a movie that contains a single
// audio track should set the movie time scale to the media time scale of that
// track. Set the value of this property on a new empty movie before you
// perform any edits on it.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/timescale
func (m AVMutableMovie) Timescale() int32 {
	rv := objc.Send[int32](m.ID, objc.Sel("timescale"))
	return rv
}
func (m AVMutableMovie) SetTimescale(value int32) {
	objc.Send[struct{}](m.ID, objc.Sel("setTimescale:"), value)
}
// A time period indicating the duration for interleaving runs of samples for
// each track.
//
// # Discussion
// 
// Default value is `0.5` seconds.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/interleavingPeriod
func (m AVMutableMovie) InterleavingPeriod() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](m.ID, objc.Sel("interleavingPeriod"))
	return coremedia.CMTime(rv)
}
func (m AVMutableMovie) SetInterleavingPeriod(value coremedia.CMTime) {
	objc.Send[struct{}](m.ID, objc.Sel("setInterleavingPeriod:"), value)
}
// The track groups an asset contains.
//
// # Discussion
// 
// This value is an empty array if the composition has no track groups.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/trackGroups
func (m AVMutableMovie) TrackGroups() IAVAssetTrackGroup {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("trackGroups"))
	return AVAssetTrackGroupFromID(objc.ID(rv))
}
func (m AVMutableMovie) SetTrackGroups(value IAVAssetTrackGroup) {
	objc.Send[struct{}](m.ID, objc.Sel("setTrackGroups:"), value)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/duration
func (m AVMutableMovie) Duration() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](m.ID, objc.Sel("duration"))
	return coremedia.CMTime(rv)
}
func (m AVMutableMovie) SetDuration(value coremedia.CMTime) {
	objc.Send[struct{}](m.ID, objc.Sel("setDuration:"), value)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/providesPreciseDurationAndTiming
func (m AVMutableMovie) ProvidesPreciseDurationAndTiming() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("providesPreciseDurationAndTiming"))
	return rv
}
func (m AVMutableMovie) SetProvidesPreciseDurationAndTiming(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setProvidesPreciseDurationAndTiming:"), value)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/minimumTimeOffsetFromLive
func (m AVMutableMovie) MinimumTimeOffsetFromLive() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](m.ID, objc.Sel("minimumTimeOffsetFromLive"))
	return coremedia.CMTime(rv)
}
func (m AVMutableMovie) SetMinimumTimeOffsetFromLive(value coremedia.CMTime) {
	objc.Send[struct{}](m.ID, objc.Sel("setMinimumTimeOffsetFromLive:"), value)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/metadata
func (m AVMutableMovie) Metadata() []AVMetadataItem {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("metadata"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}
func (m AVMutableMovie) SetMetadata(value []AVMetadataItem) {
	objc.Send[struct{}](m.ID, objc.Sel("setMetadata:"), objectivec.IObjectSliceToNSArray(value))
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/commonMetadata
func (m AVMutableMovie) CommonMetadata() IAVMetadataItem {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("commonMetadata"))
	return AVMetadataItemFromID(objc.ID(rv))
}
func (m AVMutableMovie) SetCommonMetadata(value IAVMetadataItem) {
	objc.Send[struct{}](m.ID, objc.Sel("setCommonMetadata:"), value)
}
// The metadata formats this asset contains.
//
// # Discussion
// 
// Metadata formats may include ID3, iTunes metadata, and so on.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/availableMetadataFormats
func (m AVMutableMovie) AvailableMetadataFormats() AVMetadataFormat {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("availableMetadataFormats"))
	return AVMetadataFormat(foundation.NSStringFromID(rv).String())
}
func (m AVMutableMovie) SetAvailableMetadataFormats(value AVMetadataFormat) {
	objc.Send[struct{}](m.ID, objc.Sel("setAvailableMetadataFormats:"), objc.String(string(value)))
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/creationDate
func (m AVMutableMovie) CreationDate() IAVMetadataItem {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("creationDate"))
	return AVMetadataItemFromID(objc.ID(rv))
}
func (m AVMutableMovie) SetCreationDate(value IAVMetadataItem) {
	objc.Send[struct{}](m.ID, objc.Sel("setCreationDate:"), value)
}
// The lyrics of the asset in a language suitable for the current locale.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/lyrics
func (m AVMutableMovie) Lyrics() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("lyrics"))
	return foundation.NSStringFromID(rv).String()
}
func (m AVMutableMovie) SetLyrics(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setLyrics:"), objc.String(value))
}
// A Boolean value that indicates whether the asset has playable content.
//
// # Discussion
// 
// This property value is [true] if you can use the movie to create an
// [AVPlayerItem].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/isPlayable
func (m AVMutableMovie) IsPlayable() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isPlayable"))
	return rv
}
func (m AVMutableMovie) SetIsPlayable(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setPlayable:"), value)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/isReadable
func (m AVMutableMovie) IsReadable() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isReadable"))
	return rv
}
func (m AVMutableMovie) SetIsReadable(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setReadable:"), value)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/isExportable
func (m AVMutableMovie) IsExportable() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isExportable"))
	return rv
}
func (m AVMutableMovie) SetIsExportable(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setExportable:"), value)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/isComposable
func (m AVMutableMovie) IsComposable() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isComposable"))
	return rv
}
func (m AVMutableMovie) SetIsComposable(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setComposable:"), value)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/isCompatibleWithAirPlayVideo
func (m AVMutableMovie) IsCompatibleWithAirPlayVideo() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isCompatibleWithAirPlayVideo"))
	return rv
}
func (m AVMutableMovie) SetIsCompatibleWithAirPlayVideo(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setCompatibleWithAirPlayVideo:"), value)
}
// The asset’s rate preference for playing its media.
//
// # Discussion
// 
// This value is typically, but not always, 1.0.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/preferredRate
func (m AVMutableMovie) PreferredRate() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("preferredRate"))
	return rv
}
func (m AVMutableMovie) SetPreferredRate(value float32) {
	objc.Send[struct{}](m.ID, objc.Sel("setPreferredRate:"), value)
}
// The asset’s volume preference for playing its audible media.
//
// # Discussion
// 
// This value is typically, but not always, 1.0.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/preferredVolume
func (m AVMutableMovie) PreferredVolume() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("preferredVolume"))
	return rv
}
func (m AVMutableMovie) SetPreferredVolume(value float32) {
	objc.Send[struct{}](m.ID, objc.Sel("setPreferredVolume:"), value)
}
// The asset’s transform preference to apply to its visual content during
// presentation or processing.
//
// # Discussion
// 
// The value is typically, but not always, the identity transform.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/preferredTransform
func (m AVMutableMovie) PreferredTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](m.ID, objc.Sel("preferredTransform"))
	return corefoundation.CGAffineTransform(rv)
}
func (m AVMutableMovie) SetPreferredTransform(value corefoundation.CGAffineTransform) {
	objc.Send[struct{}](m.ID, objc.Sel("setPreferredTransform:"), value)
}
// The default media selections for this asset’s media selection groups.
//
// # Discussion
// 
// Provides an instance of [AVMediaSelection] with the default selections for
// each of the assets media selection groups.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/preferredMediaSelection
func (m AVMutableMovie) PreferredMediaSelection() IAVMediaSelection {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("preferredMediaSelection"))
	return AVMediaSelectionFromID(objc.ID(rv))
}
func (m AVMutableMovie) SetPreferredMediaSelection(value IAVMediaSelection) {
	objc.Send[struct{}](m.ID, objc.Sel("setPreferredMediaSelection:"), value)
}
// The array of available media selections for this asset.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/allMediaSelections
func (m AVMutableMovie) AllMediaSelections() IAVMediaSelection {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("allMediaSelections"))
	return AVMediaSelectionFromID(objc.ID(rv))
}
func (m AVMutableMovie) SetAllMediaSelections(value IAVMediaSelection) {
	objc.Send[struct{}](m.ID, objc.Sel("setAllMediaSelections:"), value)
}
// An array of media characteristics for which a media selection option is
// available.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/availableMediaCharacteristicsWithMediaSelectionOptions
func (m AVMutableMovie) AvailableMediaCharacteristicsWithMediaSelectionOptions() AVMediaCharacteristic {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("availableMediaCharacteristicsWithMediaSelectionOptions"))
	return AVMediaCharacteristic(foundation.NSStringFromID(rv).String())
}
func (m AVMutableMovie) SetAvailableMediaCharacteristicsWithMediaSelectionOptions(value AVMediaCharacteristic) {
	objc.Send[struct{}](m.ID, objc.Sel("setAvailableMediaCharacteristicsWithMediaSelectionOptions:"), objc.String(string(value)))
}
// The locales of the asset’s chapter metadata.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/availableChapterLocales
func (m AVMutableMovie) AvailableChapterLocales() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("availableChapterLocales"))
	return objectivec.Object{ID: rv}
}
func (m AVMutableMovie) SetAvailableChapterLocales(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setAvailableChapterLocales:"), value)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/hasProtectedContent
func (m AVMutableMovie) HasProtectedContent() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasProtectedContent"))
	return rv
}
func (m AVMutableMovie) SetHasProtectedContent(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setHasProtectedContent:"), value)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/canContainFragments
func (m AVMutableMovie) CanContainFragments() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("canContainFragments"))
	return rv
}
func (m AVMutableMovie) SetCanContainFragments(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setCanContainFragments:"), value)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/containsFragments
func (m AVMutableMovie) ContainsFragments() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("containsFragments"))
	return rv
}
func (m AVMutableMovie) SetContainsFragments(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setContainsFragments:"), value)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovie/overallDurationHint
func (m AVMutableMovie) OverallDurationHint() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](m.ID, objc.Sel("overallDurationHint"))
	return coremedia.CMTime(rv)
}
func (m AVMutableMovie) SetOverallDurationHint(value coremedia.CMTime) {
	objc.Send[struct{}](m.ID, objc.Sel("setOverallDurationHint:"), value)
}
// A preset to export the asset in its current format, unless otherwise
// prohibited.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetexportpresetpassthrough
func (m AVMutableMovie) AVAssetExportPresetPassthrough() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("AVAssetExportPresetPassthrough"))
	return foundation.NSStringFromID(rv).String()
}
// A Boolean value that indicates whether to optimize the movie for network
// use.
//
// See: https://developer.apple.com/documentation/avfoundation/avassetexportsession/shouldoptimizefornetworkuse
func (m AVMutableMovie) ShouldOptimizeForNetworkUse() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("shouldOptimizeForNetworkUse"))
	return rv
}
func (m AVMutableMovie) SetShouldOptimizeForNetworkUse(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setShouldOptimizeForNetworkUse:"), value)
}

