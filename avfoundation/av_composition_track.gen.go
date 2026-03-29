// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
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
//   - [AVCompositionTrack.SetIsPlayable]
//   - [AVCompositionTrack.IsDecodable]: A Boolean value that indicates whether the track is decodable in the current environment.
//   - [AVCompositionTrack.SetIsDecodable]
//   - [AVCompositionTrack.IsEnabled]: A Boolean value that indicates whether the track’s container enables it.
//   - [AVCompositionTrack.SetIsEnabled]
//   - [AVCompositionTrack.IsSelfContained]: A Boolean value that indicates whether this track references sample data only within its container file.
//   - [AVCompositionTrack.SetIsSelfContained]
//   - [AVCompositionTrack.TotalSampleDataLength]: The total number of bytes of sample data the track requires.
//   - [AVCompositionTrack.SetTotalSampleDataLength]
//   - [AVCompositionTrack.HasMediaCharacteristic]: Returns a Boolean value that indicates whether the track references media with the specified media characteristic.
//
// # Accessing temporal information
//
//   - [AVCompositionTrack.TimeRange]: The time range of the track within the overall timeline of the asset.
//   - [AVCompositionTrack.SetTimeRange]
//   - [AVCompositionTrack.NaturalTimeScale]: The natural time scale of the media that a track references.
//   - [AVCompositionTrack.SetNaturalTimeScale]
//   - [AVCompositionTrack.EstimatedDataRate]: The estimated data rate, in bits per second, of the media that the track references.
//   - [AVCompositionTrack.SetEstimatedDataRate]
//   - [AVCompositionTrack.SamplePresentationTimeForTrackTime]: Maps the specified track time through the appropriate time mapping and returns the resulting sample presentation time.
//
// # Accessing language support
//
//   - [AVCompositionTrack.LanguageCode]: The language code of the track.
//   - [AVCompositionTrack.SetLanguageCode]
//   - [AVCompositionTrack.ExtendedLanguageTag]: The language tag of the track.
//   - [AVCompositionTrack.SetExtendedLanguageTag]
//
// # Managing format descriptions
//
//   - [AVCompositionTrack.FormatDescriptions]: The format descriptions of the media samples that a track references.
//   - [AVCompositionTrack.SetFormatDescriptions]
//   - [AVCompositionTrack.FormatDescriptionReplacements]: The replacement format descriptions.
//
// # Accessing visual characteristics
//
//   - [AVCompositionTrack.NaturalSize]: The natural dimensions of the media data that the track references.
//   - [AVCompositionTrack.SetNaturalSize]
//   - [AVCompositionTrack.PreferredTransform]: The track’s transform preference to apply to its visual content during presentation or processing.
//   - [AVCompositionTrack.SetPreferredTransform]
//
// # Accessing audible characteristics
//
//   - [AVCompositionTrack.PreferredVolume]: The track’s volume preference for playing its audible media.
//   - [AVCompositionTrack.SetPreferredVolume]
//   - [AVCompositionTrack.HasAudioSampleDependencies]: A Boolean value that indicates whether the track has sample dependencies.
//   - [AVCompositionTrack.SetHasAudioSampleDependencies]
//
// # Accessing frame-based characteristics
//
//   - [AVCompositionTrack.NominalFrameRate]: The frame rate of the track, in frames per second.
//   - [AVCompositionTrack.SetNominalFrameRate]
//   - [AVCompositionTrack.MinFrameDuration]: The minimum duration of the track’s frames.
//   - [AVCompositionTrack.SetMinFrameDuration]
//   - [AVCompositionTrack.RequiresFrameReordering]: A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
//   - [AVCompositionTrack.SetRequiresFrameReordering]
//
// # Accessing metadata
//
//   - [AVCompositionTrack.Metadata]: An array of metadata items for all metadata identifiers that have a value.
//   - [AVCompositionTrack.SetMetadata]
//   - [AVCompositionTrack.CommonMetadata]: An array of metadata items for all common metadata keys that have a value.
//   - [AVCompositionTrack.SetCommonMetadata]
//   - [AVCompositionTrack.AvailableMetadataFormats]: An array of metadata formats available for the track.
//   - [AVCompositionTrack.SetAvailableMetadataFormats]
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
//   - [AVCompositionTrack.SetAvailableTrackAssociationTypes]
//   - [AVCompositionTrack.AssociatedTracksOfType]: Returns an array of associated tracks that have the specified association type.
//
// # Determining sample cursor support
//
//   - [AVCompositionTrack.CanProvideSampleCursors]: A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
//   - [AVCompositionTrack.SetCanProvideSampleCursors]
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
//   - [IAVCompositionTrack.SetIsPlayable]
//   - [IAVCompositionTrack.IsDecodable]: A Boolean value that indicates whether the track is decodable in the current environment.
//   - [IAVCompositionTrack.SetIsDecodable]
//   - [IAVCompositionTrack.IsEnabled]: A Boolean value that indicates whether the track’s container enables it.
//   - [IAVCompositionTrack.SetIsEnabled]
//   - [IAVCompositionTrack.IsSelfContained]: A Boolean value that indicates whether this track references sample data only within its container file.
//   - [IAVCompositionTrack.SetIsSelfContained]
//   - [IAVCompositionTrack.TotalSampleDataLength]: The total number of bytes of sample data the track requires.
//   - [IAVCompositionTrack.SetTotalSampleDataLength]
//   - [IAVCompositionTrack.HasMediaCharacteristic]: Returns a Boolean value that indicates whether the track references media with the specified media characteristic.
//
// # Accessing temporal information
//
//   - [IAVCompositionTrack.TimeRange]: The time range of the track within the overall timeline of the asset.
//   - [IAVCompositionTrack.SetTimeRange]
//   - [IAVCompositionTrack.NaturalTimeScale]: The natural time scale of the media that a track references.
//   - [IAVCompositionTrack.SetNaturalTimeScale]
//   - [IAVCompositionTrack.EstimatedDataRate]: The estimated data rate, in bits per second, of the media that the track references.
//   - [IAVCompositionTrack.SetEstimatedDataRate]
//   - [IAVCompositionTrack.SamplePresentationTimeForTrackTime]: Maps the specified track time through the appropriate time mapping and returns the resulting sample presentation time.
//
// # Accessing language support
//
//   - [IAVCompositionTrack.LanguageCode]: The language code of the track.
//   - [IAVCompositionTrack.SetLanguageCode]
//   - [IAVCompositionTrack.ExtendedLanguageTag]: The language tag of the track.
//   - [IAVCompositionTrack.SetExtendedLanguageTag]
//
// # Managing format descriptions
//
//   - [IAVCompositionTrack.FormatDescriptions]: The format descriptions of the media samples that a track references.
//   - [IAVCompositionTrack.SetFormatDescriptions]
//   - [IAVCompositionTrack.FormatDescriptionReplacements]: The replacement format descriptions.
//
// # Accessing visual characteristics
//
//   - [IAVCompositionTrack.NaturalSize]: The natural dimensions of the media data that the track references.
//   - [IAVCompositionTrack.SetNaturalSize]
//   - [IAVCompositionTrack.PreferredTransform]: The track’s transform preference to apply to its visual content during presentation or processing.
//   - [IAVCompositionTrack.SetPreferredTransform]
//
// # Accessing audible characteristics
//
//   - [IAVCompositionTrack.PreferredVolume]: The track’s volume preference for playing its audible media.
//   - [IAVCompositionTrack.SetPreferredVolume]
//   - [IAVCompositionTrack.HasAudioSampleDependencies]: A Boolean value that indicates whether the track has sample dependencies.
//   - [IAVCompositionTrack.SetHasAudioSampleDependencies]
//
// # Accessing frame-based characteristics
//
//   - [IAVCompositionTrack.NominalFrameRate]: The frame rate of the track, in frames per second.
//   - [IAVCompositionTrack.SetNominalFrameRate]
//   - [IAVCompositionTrack.MinFrameDuration]: The minimum duration of the track’s frames.
//   - [IAVCompositionTrack.SetMinFrameDuration]
//   - [IAVCompositionTrack.RequiresFrameReordering]: A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
//   - [IAVCompositionTrack.SetRequiresFrameReordering]
//
// # Accessing metadata
//
//   - [IAVCompositionTrack.Metadata]: An array of metadata items for all metadata identifiers that have a value.
//   - [IAVCompositionTrack.SetMetadata]
//   - [IAVCompositionTrack.CommonMetadata]: An array of metadata items for all common metadata keys that have a value.
//   - [IAVCompositionTrack.SetCommonMetadata]
//   - [IAVCompositionTrack.AvailableMetadataFormats]: An array of metadata formats available for the track.
//   - [IAVCompositionTrack.SetAvailableMetadataFormats]
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
//   - [IAVCompositionTrack.SetAvailableTrackAssociationTypes]
//   - [IAVCompositionTrack.AssociatedTracksOfType]: Returns an array of associated tracks that have the specified association type.
//
// # Determining sample cursor support
//
//   - [IAVCompositionTrack.CanProvideSampleCursors]: A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
//   - [IAVCompositionTrack.SetCanProvideSampleCursors]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack
type IAVCompositionTrack interface {
	IAVAssetTrack

	// Topic: Accessing track information

	// A Boolean value that indicates whether the track is playable in the current environment.
	IsPlayable() bool
	SetIsPlayable(value bool)
	// A Boolean value that indicates whether the track is decodable in the current environment.
	IsDecodable() bool
	SetIsDecodable(value bool)
	// A Boolean value that indicates whether the track’s container enables it.
	IsEnabled() bool
	SetIsEnabled(value bool)
	// A Boolean value that indicates whether this track references sample data only within its container file.
	IsSelfContained() bool
	SetIsSelfContained(value bool)
	// The total number of bytes of sample data the track requires.
	TotalSampleDataLength() objectivec.IObject
	SetTotalSampleDataLength(value objectivec.IObject)
	// Returns a Boolean value that indicates whether the track references media with the specified media characteristic.
	HasMediaCharacteristic(mediaCharacteristic AVMediaCharacteristic) bool

	// Topic: Accessing temporal information

	// The time range of the track within the overall timeline of the asset.
	TimeRange() uintptr
	SetTimeRange(value uintptr)
	// The natural time scale of the media that a track references.
	NaturalTimeScale() int32
	SetNaturalTimeScale(value int32)
	// The estimated data rate, in bits per second, of the media that the track references.
	EstimatedDataRate() float32
	SetEstimatedDataRate(value float32)
	// Maps the specified track time through the appropriate time mapping and returns the resulting sample presentation time.
	SamplePresentationTimeForTrackTime(trackTime uintptr) uintptr

	// Topic: Accessing language support

	// The language code of the track.
	LanguageCode() string
	SetLanguageCode(value string)
	// The language tag of the track.
	ExtendedLanguageTag() string
	SetExtendedLanguageTag(value string)

	// Topic: Managing format descriptions

	// The format descriptions of the media samples that a track references.
	FormatDescriptions() objectivec.IObject
	SetFormatDescriptions(value objectivec.IObject)
	// The replacement format descriptions.
	FormatDescriptionReplacements() []AVCompositionTrackFormatDescriptionReplacement

	// Topic: Accessing visual characteristics

	// The natural dimensions of the media data that the track references.
	NaturalSize() corefoundation.CGSize
	SetNaturalSize(value corefoundation.CGSize)
	// The track’s transform preference to apply to its visual content during presentation or processing.
	PreferredTransform() corefoundation.CGAffineTransform
	SetPreferredTransform(value corefoundation.CGAffineTransform)

	// Topic: Accessing audible characteristics

	// The track’s volume preference for playing its audible media.
	PreferredVolume() float32
	SetPreferredVolume(value float32)
	// A Boolean value that indicates whether the track has sample dependencies.
	HasAudioSampleDependencies() bool
	SetHasAudioSampleDependencies(value bool)

	// Topic: Accessing frame-based characteristics

	// The frame rate of the track, in frames per second.
	NominalFrameRate() float32
	SetNominalFrameRate(value float32)
	// The minimum duration of the track’s frames.
	MinFrameDuration() uintptr
	SetMinFrameDuration(value uintptr)
	// A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
	RequiresFrameReordering() bool
	SetRequiresFrameReordering(value bool)

	// Topic: Accessing metadata

	// An array of metadata items for all metadata identifiers that have a value.
	Metadata() IAVMetadataItem
	SetMetadata(value IAVMetadataItem)
	// An array of metadata items for all common metadata keys that have a value.
	CommonMetadata() IAVMetadataItem
	SetCommonMetadata(value IAVMetadataItem)
	// An array of metadata formats available for the track.
	AvailableMetadataFormats() AVMetadataFormat
	SetAvailableMetadataFormats(value AVMetadataFormat)
	// Returns metadata items that a track contains for the specified format.
	MetadataForFormat(format AVMetadataFormat) []AVMetadataItem

	// Topic: Accessing track segments

	// The time mappings from the track’s media samples to its timeline.
	Segments() []AVCompositionTrackSegment
	// Returns a segment whose target time range contains, or is closest to, the specified track time.
	SegmentForTrackTime(trackTime uintptr) IAVCompositionTrackSegment

	// Topic: Accessing track associations

	// An array of association types that the track uses to associate with other tracks.
	AvailableTrackAssociationTypes() objc.ID
	SetAvailableTrackAssociationTypes(value objc.ID)
	// Returns an array of associated tracks that have the specified association type.
	AssociatedTracksOfType(trackAssociationType AVTrackAssociationType) []AVAssetTrack

	// Topic: Determining sample cursor support

	// A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
	CanProvideSampleCursors() bool
	SetCanProvideSampleCursors(value bool)
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
// [true] if the track references media with the specified characteristic;
// otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/samplePresentationTime(forTrackTime:)
func (c AVCompositionTrack) SamplePresentationTimeForTrackTime(trackTime uintptr) uintptr {
	rv := objc.Send[uintptr](c.ID, objc.Sel("samplePresentationTimeForTrackTime:"), trackTime)
	return rv
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
func (c AVCompositionTrack) SegmentForTrackTime(trackTime uintptr) IAVCompositionTrackSegment {
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
// [LoadAssociatedTracksOfTypeCompletionHandler] instead.
// 
// You can call this method without blocking the current thread after you’ve
// loaded the [availableTrackAssociationTypes] property.
//
// [availableTrackAssociationTypes]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/availableTrackAssociationTypes
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/associatedTracks(ofType:)
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
func (c AVCompositionTrack) SetIsPlayable(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setPlayable:"), value)
}
// A Boolean value that indicates whether the track is decodable in the
// current environment.
//
// # Discussion
// 
// When this property is [true], the system can decode the track, even if
// decoding may be too slow for real-time playback.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isDecodable
func (c AVCompositionTrack) IsDecodable() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isDecodable"))
	return rv
}
func (c AVCompositionTrack) SetIsDecodable(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setDecodable:"), value)
}
// A Boolean value that indicates whether the track’s container enables it.
//
// # Discussion
// 
// For file-based media, you can change its [Enabled] presentation state using
// [AVPlayerItemTrack].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isEnabled
func (c AVCompositionTrack) IsEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isEnabled"))
	return rv
}
func (c AVCompositionTrack) SetIsEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setEnabled:"), value)
}
// A Boolean value that indicates whether this track references sample data
// only within its container file.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/isSelfContained
func (c AVCompositionTrack) IsSelfContained() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isSelfContained"))
	return rv
}
func (c AVCompositionTrack) SetIsSelfContained(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setSelfContained:"), value)
}
// The total number of bytes of sample data the track requires.
//
// # Discussion
// 
// The value may be `0` if the framework can’t determine the total sample
// data length.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/totalSampleDataLength
func (c AVCompositionTrack) TotalSampleDataLength() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("totalSampleDataLength"))
	return objectivec.Object{ID: rv}
}
func (c AVCompositionTrack) SetTotalSampleDataLength(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setTotalSampleDataLength:"), value)
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
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/timeRange
func (c AVCompositionTrack) TimeRange() uintptr {
	rv := objc.Send[uintptr](c.ID, objc.Sel("timeRange"))
	return rv
}
func (c AVCompositionTrack) SetTimeRange(value uintptr) {
	objc.Send[struct{}](c.ID, objc.Sel("setTimeRange:"), value)
}
// The natural time scale of the media that a track references.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/naturalTimeScale
func (c AVCompositionTrack) NaturalTimeScale() int32 {
	rv := objc.Send[int32](c.ID, objc.Sel("naturalTimeScale"))
	return rv
}
func (c AVCompositionTrack) SetNaturalTimeScale(value int32) {
	objc.Send[struct{}](c.ID, objc.Sel("setNaturalTimeScale:"), value)
}
// The estimated data rate, in bits per second, of the media that the track
// references.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/estimatedDataRate
func (c AVCompositionTrack) EstimatedDataRate() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("estimatedDataRate"))
	return rv
}
func (c AVCompositionTrack) SetEstimatedDataRate(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setEstimatedDataRate:"), value)
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
func (c AVCompositionTrack) SetLanguageCode(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setLanguageCode:"), objc.String(value))
}
// The language tag of the track.
//
// # Discussion
// 
// The value is a [BCP-47] language tag, or `nil` if the track doesn’t
// specify a language tag.
//
// [BCP-47]: https://tools.ietf.org/html/bcp47
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/extendedLanguageTag
func (c AVCompositionTrack) ExtendedLanguageTag() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("extendedLanguageTag"))
	return foundation.NSStringFromID(rv).String()
}
func (c AVCompositionTrack) SetExtendedLanguageTag(value string) {
	objc.Send[struct{}](c.ID, objc.Sel("setExtendedLanguageTag:"), objc.String(value))
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
// [CMFormatDescription]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescription
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/formatDescriptions
func (c AVCompositionTrack) FormatDescriptions() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("formatDescriptions"))
	return objectivec.Object{ID: rv}
}
func (c AVCompositionTrack) SetFormatDescriptions(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setFormatDescriptions:"), value)
}
// The replacement format descriptions.
//
// # Discussion
// 
// The property’s values specify an original and a replacement format
// description, as set in a previous call to
// [ReplaceFormatDescriptionWithFormatDescription].
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
// [zero]: https://developer.apple.com/documentation/CoreFoundation/CGSize/zero
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/naturalSize
func (c AVCompositionTrack) NaturalSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("naturalSize"))
	return corefoundation.CGSize(rv)
}
func (c AVCompositionTrack) SetNaturalSize(value corefoundation.CGSize) {
	objc.Send[struct{}](c.ID, objc.Sel("setNaturalSize:"), value)
}
// The track’s transform preference to apply to its visual content during
// presentation or processing.
//
// # Discussion
// 
// The value of this property is typically, but not always,
// [CGAffineTransformIdentity].
//
// [CGAffineTransformIdentity]: https://developer.apple.com/documentation/CoreGraphics/CGAffineTransformIdentity
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/preferredTransform
func (c AVCompositionTrack) PreferredTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](c.ID, objc.Sel("preferredTransform"))
	return corefoundation.CGAffineTransform(rv)
}
func (c AVCompositionTrack) SetPreferredTransform(value corefoundation.CGAffineTransform) {
	objc.Send[struct{}](c.ID, objc.Sel("setPreferredTransform:"), value)
}
// The track’s volume preference for playing its audible media.
//
// # Discussion
// 
// The preferred volume for an audio track is typically, but not always,
// `1.0`. For non-audible tracks, the value is `0.0`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/preferredVolume
func (c AVCompositionTrack) PreferredVolume() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("preferredVolume"))
	return rv
}
func (c AVCompositionTrack) SetPreferredVolume(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setPreferredVolume:"), value)
}
// A Boolean value that indicates whether the track has sample dependencies.
//
// # Discussion
// 
// The value is always [false] for nonaudible media.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/hasAudioSampleDependencies
func (c AVCompositionTrack) HasAudioSampleDependencies() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasAudioSampleDependencies"))
	return rv
}
func (c AVCompositionTrack) SetHasAudioSampleDependencies(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setHasAudioSampleDependencies:"), value)
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
func (c AVCompositionTrack) NominalFrameRate() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("nominalFrameRate"))
	return rv
}
func (c AVCompositionTrack) SetNominalFrameRate(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setNominalFrameRate:"), value)
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
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/minFrameDuration
func (c AVCompositionTrack) MinFrameDuration() uintptr {
	rv := objc.Send[uintptr](c.ID, objc.Sel("minFrameDuration"))
	return rv
}
func (c AVCompositionTrack) SetMinFrameDuration(value uintptr) {
	objc.Send[struct{}](c.ID, objc.Sel("setMinFrameDuration:"), value)
}
// A Boolean value that indicates whether samples in the track may have
// different presentation and decode timestamps.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/requiresFrameReordering
func (c AVCompositionTrack) RequiresFrameReordering() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("requiresFrameReordering"))
	return rv
}
func (c AVCompositionTrack) SetRequiresFrameReordering(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setRequiresFrameReordering:"), value)
}
// An array of metadata items for all metadata identifiers that have a value.
//
// # Discussion
// 
// You can filter the array of metadata items according to language using the
// [MetadataItemsFromArrayFilteredAndSortedAccordingToPreferredLanguages]
// method. Filter the results by identifier using the
// [MetadataItemsFromArrayFilteredByIdentifier] method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/metadata
func (c AVCompositionTrack) Metadata() IAVMetadataItem {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("metadata"))
	return AVMetadataItemFromID(objc.ID(rv))
}
func (c AVCompositionTrack) SetMetadata(value IAVMetadataItem) {
	objc.Send[struct{}](c.ID, objc.Sel("setMetadata:"), value)
}
// An array of metadata items for all common metadata keys that have a value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/commonMetadata
func (c AVCompositionTrack) CommonMetadata() IAVMetadataItem {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("commonMetadata"))
	return AVMetadataItemFromID(objc.ID(rv))
}
func (c AVCompositionTrack) SetCommonMetadata(value IAVMetadataItem) {
	objc.Send[struct{}](c.ID, objc.Sel("setCommonMetadata:"), value)
}
// An array of metadata formats available for the track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/availableMetadataFormats
func (c AVCompositionTrack) AvailableMetadataFormats() AVMetadataFormat {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("availableMetadataFormats"))
	return AVMetadataFormat(foundation.NSStringFromID(rv).String())
}
func (c AVCompositionTrack) SetAvailableMetadataFormats(value AVMetadataFormat) {
	objc.Send[struct{}](c.ID, objc.Sel("setAvailableMetadataFormats:"), objc.String(string(value)))
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
func (c AVCompositionTrack) AvailableTrackAssociationTypes() objc.ID {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("availableTrackAssociationTypes"))
	return rv
}
func (c AVCompositionTrack) SetAvailableTrackAssociationTypes(value objc.ID) {
	objc.Send[struct{}](c.ID, objc.Sel("setAvailableTrackAssociationTypes:"), value)
}
// A Boolean value that indicates whether the track can provide instances of
// sample cursors to traverse its media samples and discover information.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCompositionTrack/canProvideSampleCursors
func (c AVCompositionTrack) CanProvideSampleCursors() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("canProvideSampleCursors"))
	return rv
}
func (c AVCompositionTrack) SetCanProvideSampleCursors(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setCanProvideSampleCursors:"), value)
}

