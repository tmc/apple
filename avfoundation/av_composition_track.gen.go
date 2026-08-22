// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCompositionTrack] class.
var (
	_AVCompositionTrackClass     AVCompositionTrackClass
	_AVCompositionTrackClassOnce sync.Once
)

func getAVCompositionTrackClass() AVCompositionTrackClass {
	_AVCompositionTrackClassOnce.Do(func() {
		_AVCompositionTrackClass = AVCompositionTrackClass{class: objc.GetClass("AVCompositionTrack")}
	})
	return _AVCompositionTrackClass
}

// GetAVCompositionTrackClass returns the class object for AVCompositionTrack.
func GetAVCompositionTrackClass() AVCompositionTrackClass {
	return getAVCompositionTrackClass()
}

type AVCompositionTrackClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCompositionTrackClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCompositionTrackClass) Alloc() AVCompositionTrack {
	rv := objc.Send[AVCompositionTrack](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A track in a composition that presents media of a uniform type.
//
// # Overview
//
// This object provides an immutable composition track. The framework also
// provides a mutable subclass, [AVMutableCompositionTrack].
//
// # Accessing track information
//
//   - [AVCompositionTrack.IsPlayable]: A Boolean value that indicates whether the track is playable in the current environment.
//   - [AVCompositionTrack.IsDecodable]: A Boolean value that indicates whether the track is decodable in the current environment.
//   - [AVCompositionTrack.IsEnabled]: A Boolean value that indicates whether the track’s container enables it.
//   - [AVCompositionTrack.IsSelfContained]: A Boolean value that indicates whether this track references sample data only within its container file.
//   - [AVCompositionTrack.TotalSampleDataLength]: The total number of bytes of sample data the track requires.
//   - [AVCompositionTrack.HasMediaCharacteristic]: Returns a Boolean value that indicates whether the track references media with the specified media characteristic.
//
// # Accessing temporal information
//
//   - [AVCompositionTrack.TimeRange]: The time range of the track within the overall timeline of the asset.
//   - [AVCompositionTrack.NaturalTimeScale]: The natural time scale of the media that a track references.
//   - [AVCompositionTrack.EstimatedDataRate]: The estimated data rate, in bits per second, of the media that the track references.
//   - [AVCompositionTrack.SamplePresentationTimeForTrackTime]: Maps the specified track time through the appropriate time mapping and returns the resulting sample presentation time.
//
// # Accessing language support
//
//   - [AVCompositionTrack.LanguageCode]: The language code of the track.
//   - [AVCompositionTrack.ExtendedLanguageTag]: The language tag of the track.
//
// # Managing format descriptions
//
//   - [AVCompositionTrack.FormatDescriptions]: The format descriptions of the media samples that a track references.
//   - [AVCompositionTrack.FormatDescriptionReplacements]: The replacement format descriptions.
//
// # Accessing visual characteristics
//
//   - [AVCompositionTrack.NaturalSize]: The natural dimensions of the media data that the track references.
//   - [AVCompositionTrack.PreferredTransform]: The track’s transform preference to apply to its visual content during presentation or processing.
//
// # Accessing audible characteristics
//
//   - [AVCompositionTrack.PreferredVolume]: The track’s volume preference for playing its audible media.
//   - [AVCompositionTrack.HasAudioSampleDependencies]: A Boolean value that indicates whether the track has sample dependencies.
//
// # Accessing frame-based characteristics
//
//   - [AVCompositionTrack.NominalFrameRate]: The frame rate of the track, in frames per second.
//   - [AVCompositionTrack.MinFrameDuration]: The minimum duration of the track’s frames.
//   - [AVCompositionTrack.RequiresFrameReordering]: A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
//
// # Accessing metadata
//
//   - [AVCompositionTrack.Metadata]: An array of metadata items for all metadata identifiers that have a value.
//   - [AVCompositionTrack.CommonMetadata]: An array of metadata items for all common metadata keys that have a value.
//   - [AVCompositionTrack.AvailableMetadataFormats]: An array of metadata formats available for the track.
//   - [AVCompositionTrack.MetadataForFormat]: Returns metadata items that a track contains for the specified format.
//
// # Accessing track segments
//
//   - [AVCompositionTrack.Segments]: The time mappings from the track’s media samples to its timeline.
//   - [AVCompositionTrack.SegmentForTrackTime]: Returns a segment whose target time range contains, or is closest to, the specified track time.
//
// # Accessing track associations
//
//   - [AVCompositionTrack.AvailableTrackAssociationTypes]: An array of association types that the track uses to associate with other tracks.
//   - [AVCompositionTrack.AssociatedTracksOfType]: Returns an array of associated tracks that have the specified association type.
//
// # Determining sample cursor support
//
//   - [AVCompositionTrack.CanProvideSampleCursors]: A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack
type AVCompositionTrack struct {
	AVAssetTrack
}

// AVCompositionTrackFromID constructs a [AVCompositionTrack] from an objc.ID.
//
// A track in a composition that presents media of a uniform type.
func AVCompositionTrackFromID(id objc.ID) AVCompositionTrack {
	return AVCompositionTrack{AVAssetTrack: AVAssetTrackFromID(id)}
}

// NOTE: AVCompositionTrack adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCompositionTrack] class.
//
// # Accessing track information
//
//   - [IAVCompositionTrack.IsPlayable]: A Boolean value that indicates whether the track is playable in the current environment.
//   - [IAVCompositionTrack.IsDecodable]: A Boolean value that indicates whether the track is decodable in the current environment.
//   - [IAVCompositionTrack.IsEnabled]: A Boolean value that indicates whether the track’s container enables it.
//   - [IAVCompositionTrack.IsSelfContained]: A Boolean value that indicates whether this track references sample data only within its container file.
//   - [IAVCompositionTrack.TotalSampleDataLength]: The total number of bytes of sample data the track requires.
//   - [IAVCompositionTrack.HasMediaCharacteristic]: Returns a Boolean value that indicates whether the track references media with the specified media characteristic.
//
// # Accessing temporal information
//
//   - [IAVCompositionTrack.TimeRange]: The time range of the track within the overall timeline of the asset.
//   - [IAVCompositionTrack.NaturalTimeScale]: The natural time scale of the media that a track references.
//   - [IAVCompositionTrack.EstimatedDataRate]: The estimated data rate, in bits per second, of the media that the track references.
//   - [IAVCompositionTrack.SamplePresentationTimeForTrackTime]: Maps the specified track time through the appropriate time mapping and returns the resulting sample presentation time.
//
// # Accessing language support
//
//   - [IAVCompositionTrack.LanguageCode]: The language code of the track.
//   - [IAVCompositionTrack.ExtendedLanguageTag]: The language tag of the track.
//
// # Managing format descriptions
//
//   - [IAVCompositionTrack.FormatDescriptions]: The format descriptions of the media samples that a track references.
//   - [IAVCompositionTrack.FormatDescriptionReplacements]: The replacement format descriptions.
//
// # Accessing visual characteristics
//
//   - [IAVCompositionTrack.NaturalSize]: The natural dimensions of the media data that the track references.
//   - [IAVCompositionTrack.PreferredTransform]: The track’s transform preference to apply to its visual content during presentation or processing.
//
// # Accessing audible characteristics
//
//   - [IAVCompositionTrack.PreferredVolume]: The track’s volume preference for playing its audible media.
//   - [IAVCompositionTrack.HasAudioSampleDependencies]: A Boolean value that indicates whether the track has sample dependencies.
//
// # Accessing frame-based characteristics
//
//   - [IAVCompositionTrack.NominalFrameRate]: The frame rate of the track, in frames per second.
//   - [IAVCompositionTrack.MinFrameDuration]: The minimum duration of the track’s frames.
//   - [IAVCompositionTrack.RequiresFrameReordering]: A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
//
// # Accessing metadata
//
//   - [IAVCompositionTrack.Metadata]: An array of metadata items for all metadata identifiers that have a value.
//   - [IAVCompositionTrack.CommonMetadata]: An array of metadata items for all common metadata keys that have a value.
//   - [IAVCompositionTrack.AvailableMetadataFormats]: An array of metadata formats available for the track.
//   - [IAVCompositionTrack.MetadataForFormat]: Returns metadata items that a track contains for the specified format.
//
// # Accessing track segments
//
//   - [IAVCompositionTrack.Segments]: The time mappings from the track’s media samples to its timeline.
//   - [IAVCompositionTrack.SegmentForTrackTime]: Returns a segment whose target time range contains, or is closest to, the specified track time.
//
// # Accessing track associations
//
//   - [IAVCompositionTrack.AvailableTrackAssociationTypes]: An array of association types that the track uses to associate with other tracks.
//   - [IAVCompositionTrack.AssociatedTracksOfType]: Returns an array of associated tracks that have the specified association type.
//
// # Determining sample cursor support
//
//   - [IAVCompositionTrack.CanProvideSampleCursors]: A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack
type IAVCompositionTrack interface {
	IAVAssetTrack

	// Topic: Accessing track information

	// A Boolean value that indicates whether the track is playable in the current environment.
	IsPlayable() bool
	// A Boolean value that indicates whether the track is decodable in the current environment.
	IsDecodable() bool
	// A Boolean value that indicates whether the track’s container enables it.
	IsEnabled() bool
	// A Boolean value that indicates whether this track references sample data only within its container file.
	IsSelfContained() bool
	// The total number of bytes of sample data the track requires.
	TotalSampleDataLength() int64
	// Returns a Boolean value that indicates whether the track references media with the specified media characteristic.
	HasMediaCharacteristic(mediaCharacteristic AVMediaCharacteristic) bool

	// Topic: Accessing temporal information

	// The time range of the track within the overall timeline of the asset.
	TimeRange() coremedia.CMTimeRange
	// The natural time scale of the media that a track references.
	NaturalTimeScale() coremedia.CMTimeScale
	// The estimated data rate, in bits per second, of the media that the track references.
	EstimatedDataRate() kernel.Float
	// Maps the specified track time through the appropriate time mapping and returns the resulting sample presentation time.
	SamplePresentationTimeForTrackTime(trackTime coremedia.CMTime) coremedia.CMTime

	// Topic: Accessing language support

	// The language code of the track.
	LanguageCode() string
	// The language tag of the track.
	ExtendedLanguageTag() string

	// Topic: Managing format descriptions

	// The format descriptions of the media samples that a track references.
	FormatDescriptions() []objectivec.IObject
	// The replacement format descriptions.
	FormatDescriptionReplacements() []AVCompositionTrackFormatDescriptionReplacement

	// Topic: Accessing visual characteristics

	// The natural dimensions of the media data that the track references.
	NaturalSize() corefoundation.CGSize
	// The track’s transform preference to apply to its visual content during presentation or processing.
	PreferredTransform() corefoundation.CGAffineTransform

	// Topic: Accessing audible characteristics

	// The track’s volume preference for playing its audible media.
	PreferredVolume() kernel.Float
	// A Boolean value that indicates whether the track has sample dependencies.
	HasAudioSampleDependencies() bool

	// Topic: Accessing frame-based characteristics

	// The frame rate of the track, in frames per second.
	NominalFrameRate() kernel.Float
	// The minimum duration of the track’s frames.
	MinFrameDuration() coremedia.CMTime
	// A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
	RequiresFrameReordering() bool

	// Topic: Accessing metadata

	// An array of metadata items for all metadata identifiers that have a value.
	Metadata() []AVMetadataItem
	// An array of metadata items for all common metadata keys that have a value.
	CommonMetadata() []AVMetadataItem
	// An array of metadata formats available for the track.
	AvailableMetadataFormats() []AVMetadataFormat
	// Returns metadata items that a track contains for the specified format.
	MetadataForFormat(format AVMetadataFormat) []AVMetadataItem

	// Topic: Accessing track segments

	// The time mappings from the track’s media samples to its timeline.
	Segments() []AVCompositionTrackSegment
	// Returns a segment whose target time range contains, or is closest to, the specified track time.
	SegmentForTrackTime(trackTime coremedia.CMTime) IAVCompositionTrackSegment

	// Topic: Accessing track associations

	// An array of association types that the track uses to associate with other tracks.
	AvailableTrackAssociationTypes() []objectivec.IObject
	// Returns an array of associated tracks that have the specified association type.
	AssociatedTracksOfType(trackAssociationType AVTrackAssociationType) []AVAssetTrack

	// Topic: Determining sample cursor support

	// A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
	CanProvideSampleCursors() bool
}

// Init initializes the instance.
func (c AVCompositionTrack) Init() AVCompositionTrack {
	rv := objc.Send[AVCompositionTrack](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCompositionTrack) Autorelease() AVCompositionTrack {
	rv := objc.Send[AVCompositionTrack](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCompositionTrack creates a new AVCompositionTrack instance.
func NewAVCompositionTrack() AVCompositionTrack {
	class := getAVCompositionTrackClass()
	rv := objc.Send[AVCompositionTrack](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a Boolean value that indicates whether the track references media
// with the specified media characteristic.
//
// mediaCharacteristic: The media characteristic of interest.
//
// # Return Value
//
// true if the track references media with the specified characteristic;
// otherwise, false.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/hasMediaCharacteristic(_:)
func (c AVCompositionTrack) HasMediaCharacteristic(mediaCharacteristic AVMediaCharacteristic) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasMediaCharacteristic:"), objc.String(string(mediaCharacteristic)))
	return rv
}

// Maps the specified track time through the appropriate time mapping and
// returns the resulting sample presentation time.
//
// trackTime: The track time for which to request the sample presentation time.
//
// # Return Value
//
// The sample presentation time corresponding to the specified time; otherwise
// [invalid] if the time is out of range.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/samplePresentationTime(forTrackTime:)
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
func (c AVCompositionTrack) SamplePresentationTimeForTrackTime(trackTime coremedia.CMTime) coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](c.ID, objc.Sel("samplePresentationTimeForTrackTime:"), trackTime)
	return coremedia.CMTime(rv)
}

// Returns metadata items that a track contains for the specified format.
//
// format: The format of the metadata items to retrieve.
//
// # Return Value
//
// An array of metadata items matching the specified format, or an empty array
// if none are found.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/metadata(forFormat:)
func (c AVCompositionTrack) MetadataForFormat(format AVMetadataFormat) []AVMetadataItem {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("metadataForFormat:"), objc.String(string(format)))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}

// Returns a segment whose target time range contains, or is closest to, the
// specified track time.
//
// trackTime: The track time of the segment to return.
//
// # Return Value
//
// The [AVCompositionTrackSegment] associated with the track time.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/segment(forTrackTime:)
func (c AVCompositionTrack) SegmentForTrackTime(trackTime coremedia.CMTime) IAVCompositionTrackSegment {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("segmentForTrackTime:"), trackTime)
	return AVCompositionTrackSegmentFromID(rv)
}

// Returns an array of associated tracks that have the specified association
// type.
//
// trackAssociationType: The requested track association type.
//
// # Return Value
//
// An array of tracks matching the specified track association type, or an
// empty array if none are found.
//
// # Discussion
//
// Apple discourages using this method in iOS 15, tvOS 15, macOS 12, and
// watchOS 8 or later. Load associated tracks asynchronously using
// [AVAssetTrack.LoadAssociatedTracksOfTypeCompletionHandler] instead.
//
// You can call this method without blocking the current thread after you’ve
// loaded the [availableTrackAssociationTypes] property.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/associatedTracks(ofType:)
//
// [availableTrackAssociationTypes]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/availableTrackAssociationTypes
func (c AVCompositionTrack) AssociatedTracksOfType(trackAssociationType AVTrackAssociationType) []AVAssetTrack {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("associatedTracksOfType:"), objc.String(string(trackAssociationType)))
	return objc.ConvertSlice(rv, func(id objc.ID) AVAssetTrack {
		return AVAssetTrackFromID(id)
	})
}

// A Boolean value that indicates whether the track is playable in the current
// environment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isPlayable
func (c AVCompositionTrack) IsPlayable() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isPlayable"))
	return rv
}

// A Boolean value that indicates whether the track is decodable in the
// current environment.
//
// # Discussion
//
// When this property is true, the system can decode the track, even if
// decoding may be too slow for real-time playback.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isDecodable
func (c AVCompositionTrack) IsDecodable() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDecodable"))
	return rv
}

// A Boolean value that indicates whether the track’s container enables it.
//
// # Discussion
//
// For file-based media, you can change its [AVPlayerItemTrack.Enabled]
// presentation state using [AVPlayerItemTrack].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isEnabled
func (c AVCompositionTrack) IsEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isEnabled"))
	return rv
}

// A Boolean value that indicates whether this track references sample data
// only within its container file.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isSelfContained
func (c AVCompositionTrack) IsSelfContained() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isSelfContained"))
	return rv
}

// The total number of bytes of sample data the track requires.
//
// # Discussion
//
// The value may be `0` if the framework can’t determine the total sample
// data length.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/totalSampleDataLength
func (c AVCompositionTrack) TotalSampleDataLength() int64 {
	rv := objc.Send[int64](c.ID, objc.Sel("totalSampleDataLength"))
	return rv
}

// The time range of the track within the overall timeline of the asset.
//
// # Discussion
//
// If the start of the time range is greater than [zero], the track doesn’t
// initially have media data to present. This condition may occur when the
// media delays an audio track to align the start of audio with a specific
// video frame. You can test for this as the example below shows:
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/timeRange
//
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
func (c AVCompositionTrack) TimeRange() coremedia.CMTimeRange {
	rv := objc.Send[coremedia.CMTimeRange](c.ID, objc.Sel("timeRange"))
	return coremedia.CMTimeRange(rv)
}

// The natural time scale of the media that a track references.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/naturalTimeScale
func (c AVCompositionTrack) NaturalTimeScale() coremedia.CMTimeScale {
	rv := objc.Send[coremedia.CMTimeScale](c.ID, objc.Sel("naturalTimeScale"))
	return coremedia.CMTimeScale(rv)
}

// The estimated data rate, in bits per second, of the media that the track
// references.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/estimatedDataRate
func (c AVCompositionTrack) EstimatedDataRate() kernel.Float {
	rv := objc.Send[kernel.Float](c.ID, objc.Sel("estimatedDataRate"))
	return kernel.Float(rv)
}

// The language code of the track.
//
// # Discussion
//
// The value is an ISO 639-2/T language code, or `nil` if the track doesn’t
// specify a language code.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/languageCode
func (c AVCompositionTrack) LanguageCode() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("languageCode"))
	return foundation.NSStringFromID(rv).String()
}

// The language tag of the track.
//
// # Discussion
//
// The value is a [BCP-47] language tag, or `nil` if the track doesn’t
// specify a language tag.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/extendedLanguageTag
//
// [BCP-47]: https://tools.ietf.org/html/bcp47
func (c AVCompositionTrack) ExtendedLanguageTag() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("extendedLanguageTag"))
	return foundation.NSStringFromID(rv).String()
}

// The format descriptions of the media samples that a track references.
//
// # Discussion
//
// The array contains [CMFormatDescription] objects that indicate the format
// of media samples the track references.
//
// Asset tracks typically present uniform media (for example, media that uses
// the same encoding settings) and contain a single format description.
// However, in some cases, an asset track may contain multiple format
// descriptions. For example, an H.264-encoded video track may have some
// segments that use the Main profile and others that use the High profile.
// Also, an individual [AVCompositionTrack], which subclasses [AVAssetTrack],
// may contain audio or video segments using different codecs.
//
// You can use [CMFormatDescription] to access low-level details about the
// media the track references. For example, you can retrieve the details of
// track’s media type and subtype as the code below shows:
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/formatDescriptions
//
// [CMFormatDescription]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescription
func (c AVCompositionTrack) FormatDescriptions() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("formatDescriptions"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// The replacement format descriptions.
//
// # Discussion
//
// The property’s values specify an original and a replacement format
// description, as set in a previous call to
// [AVMutableCompositionTrack.ReplaceFormatDescriptionWithFormatDescription].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/formatDescriptionReplacements
func (c AVCompositionTrack) FormatDescriptionReplacements() []AVCompositionTrackFormatDescriptionReplacement {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("formatDescriptionReplacements"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCompositionTrackFormatDescriptionReplacement {
		return AVCompositionTrackFormatDescriptionReplacementFromID(id)
	})
}

// The natural dimensions of the media data that the track references.
//
// # Discussion
//
// For visual tracks, like video or subtitle tracks, this property value is
// the natural size of the media. For nonvisual tracks, like audio or chapter
// tracks, the value is [zero].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/naturalSize
//
// [zero]: https://developer.apple.com/documentation/CoreFoundation/CGSize/zero
func (c AVCompositionTrack) NaturalSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("naturalSize"))
	return corefoundation.CGSize(rv)
}

// The track’s transform preference to apply to its visual content during
// presentation or processing.
//
// # Discussion
//
// The value of this property is typically, but not always,
// [CGAffineTransformIdentity].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/preferredTransform
//
// [CGAffineTransformIdentity]: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformIdentity
func (c AVCompositionTrack) PreferredTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](c.ID, objc.Sel("preferredTransform"))
	return corefoundation.CGAffineTransform(rv)
}

// The track’s volume preference for playing its audible media.
//
// # Discussion
//
// The preferred volume for an audio track is typically, but not always,
// `1.0`. For non-audible tracks, the value is `0.0`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/preferredVolume
func (c AVCompositionTrack) PreferredVolume() kernel.Float {
	rv := objc.Send[kernel.Float](c.ID, objc.Sel("preferredVolume"))
	return kernel.Float(rv)
}

// A Boolean value that indicates whether the track has sample dependencies.
//
// # Discussion
//
// The value is always false for nonaudible media.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/hasAudioSampleDependencies
func (c AVCompositionTrack) HasAudioSampleDependencies() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasAudioSampleDependencies"))
	return rv
}

// The frame rate of the track, in frames per second.
//
// # Discussion
//
// The nominal frame rate indicates the number of frames per second for tracks
// that contain a full frame per media sample. For field-based (interlaced)
// video tracks, the value of this property indicates the field rate, not the
// frame rate.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/nominalFrameRate
func (c AVCompositionTrack) NominalFrameRate() kernel.Float {
	rv := objc.Send[kernel.Float](c.ID, objc.Sel("nominalFrameRate"))
	return kernel.Float(rv)
}

// The minimum duration of the track’s frames.
//
// # Discussion
//
// A track’s minimum frame duration is the reciprocal of its maximum frame
// rate. For example, a video track with a maximum frame rate of 30 frames per
// second has a minimum frame duration of 1/30, or 0.033 seconds.
//
// The value of this property is [invalid] if the track can’t calculate its
// minimum frame duration, or if it’s unknown.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/minFrameDuration
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
func (c AVCompositionTrack) MinFrameDuration() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](c.ID, objc.Sel("minFrameDuration"))
	return coremedia.CMTime(rv)
}

// A Boolean value that indicates whether samples in the track may have
// different presentation and decode timestamps.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/requiresFrameReordering
func (c AVCompositionTrack) RequiresFrameReordering() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("requiresFrameReordering"))
	return rv
}

// An array of metadata items for all metadata identifiers that have a value.
//
// # Discussion
//
// You can filter the array of metadata items according to language using the
// [AVMetadataItemClass.MetadataItemsFromArrayFilteredAndSortedAccordingToPreferredLanguages]
// method. Filter the results by identifier using the
// [AVMetadataItemClass.MetadataItemsFromArrayFilteredByIdentifier] method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/metadata
func (c AVCompositionTrack) Metadata() []AVMetadataItem {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("metadata"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}

// An array of metadata items for all common metadata keys that have a value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/commonMetadata
func (c AVCompositionTrack) CommonMetadata() []AVMetadataItem {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("commonMetadata"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}

// An array of metadata formats available for the track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/availableMetadataFormats
func (c AVCompositionTrack) AvailableMetadataFormats() []AVMetadataFormat {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("availableMetadataFormats"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataFormat {
		return AVMetadataFormat(foundation.NSStringFromID(id).String())
	})
}

// The time mappings from the track’s media samples to its timeline.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/segments
func (c AVCompositionTrack) Segments() []AVCompositionTrackSegment {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("segments"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCompositionTrackSegment {
		return AVCompositionTrackSegmentFromID(id)
	})
}

// An array of association types that the track uses to associate with other
// tracks.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/availableTrackAssociationTypes
func (c AVCompositionTrack) AvailableTrackAssociationTypes() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("availableTrackAssociationTypes"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// A Boolean value that indicates whether the track can provide instances of
// sample cursors to traverse its media samples and discover information.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/canProvideSampleCursors
func (c AVCompositionTrack) CanProvideSampleCursors() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("canProvideSampleCursors"))
	return rv
}
