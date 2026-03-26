// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMutableMovieTrack] class.
var (
	_AVMutableMovieTrackClass     AVMutableMovieTrackClass
	_AVMutableMovieTrackClassOnce sync.Once
)

func getAVMutableMovieTrackClass() AVMutableMovieTrackClass {
	_AVMutableMovieTrackClassOnce.Do(func() {
		_AVMutableMovieTrackClass = AVMutableMovieTrackClass{class: objc.GetClass("AVMutableMovieTrack")}
	})
	return _AVMutableMovieTrackClass
}

// GetAVMutableMovieTrackClass returns the class object for AVMutableMovieTrack.
func GetAVMutableMovieTrackClass() AVMutableMovieTrackClass {
	return getAVMutableMovieTrackClass()
}

type AVMutableMovieTrackClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMutableMovieTrackClass) Alloc() AVMutableMovieTrack {
	rv := objc.Send[AVMutableMovieTrack](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A mutable track that conforms to the QuickTime or ISO base media file
// format.
//
// # Managing time ranges
//
//   - [AVMutableMovieTrack.InsertTimeRangeOfTrackAtTimeCopySampleDataError]: Inserts a portion of an asset track into the target movie.
//   - [AVMutableMovieTrack.InsertEmptyTimeRange]: Adds an empty time range to a track.
//   - [AVMutableMovieTrack.RemoveTimeRange]: Removes the specified time range from a track.
//   - [AVMutableMovieTrack.ScaleTimeRangeToDuration]: Changes the duration of a time range in a track.
//
// # Appending sample data
//
//   - [AVMutableMovieTrack.InsertMediaTimeRangeIntoTimeRange]: Inserts a reference to a media time range into a track.
//
// # Accessing media chunks
//
//   - [AVMutableMovieTrack.PreferredMediaChunkAlignment]: The boundary for media chunk alignment for file types that support media chunk alignment.
//   - [AVMutableMovieTrack.SetPreferredMediaChunkAlignment]
//   - [AVMutableMovieTrack.PreferredMediaChunkDuration]: The maximum duration to use for each chunk of sample data written to the file for file types that support media chunk duration.
//   - [AVMutableMovieTrack.SetPreferredMediaChunkDuration]
//   - [AVMutableMovieTrack.PreferredMediaChunkSize]: The maximum size to use for each chunk of sample data written to the file for file types that support media chunk duration.
//   - [AVMutableMovieTrack.SetPreferredMediaChunkSize]
//
// # Changing format descriptions
//
//   - [AVMutableMovieTrack.FormatDescriptions]: The format descriptions of the media samples that a track references.
//   - [AVMutableMovieTrack.SetFormatDescriptions]
//   - [AVMutableMovieTrack.ReplaceFormatDescriptionWithFormatDescription]: Replaces the track’s format description with a new format description.
//
// # Configuring track information
//
//   - [AVMutableMovieTrack.Modified]: A Boolean value that indicates whether a track is in a modified state.
//   - [AVMutableMovieTrack.SetModified]
//   - [AVMutableMovieTrack.SampleReferenceBaseURL]: The base URL for sample references.
//   - [AVMutableMovieTrack.SetSampleReferenceBaseURL]
//
// # Accessing track information
//
//   - [AVMutableMovieTrack.IsPlayable]: A Boolean value that indicates whether the track is playable in the current environment.
//   - [AVMutableMovieTrack.SetIsPlayable]
//   - [AVMutableMovieTrack.IsDecodable]: A Boolean value that indicates whether the track is decodable in the current environment.
//   - [AVMutableMovieTrack.SetIsDecodable]
//   - [AVMutableMovieTrack.Enabled]: A Boolean value that indicates whether the track’s container enables it.
//   - [AVMutableMovieTrack.SetEnabled]
//   - [AVMutableMovieTrack.IsSelfContained]: A Boolean value that indicates whether this track references sample data only within its container file.
//   - [AVMutableMovieTrack.SetIsSelfContained]
//   - [AVMutableMovieTrack.HasProtectedContent]: A Boolean value that indicates whether a track contains protected content.
//   - [AVMutableMovieTrack.TotalSampleDataLength]: The total number of bytes of sample data the track requires.
//   - [AVMutableMovieTrack.SetTotalSampleDataLength]
//   - [AVMutableMovieTrack.HasMediaCharacteristic]: Returns a Boolean value that indicates whether the track references media with the specified media characteristic.
//
// # Accessing temporal information
//
//   - [AVMutableMovieTrack.TimeRange]: The time range of the track within the overall timeline of the asset.
//   - [AVMutableMovieTrack.SetTimeRange]
//   - [AVMutableMovieTrack.Timescale]: The time scale for tracks that contain the `moov` atom.
//   - [AVMutableMovieTrack.SetTimescale]
//   - [AVMutableMovieTrack.NaturalTimeScale]: The natural time scale of the media that a track references.
//   - [AVMutableMovieTrack.SetNaturalTimeScale]
//   - [AVMutableMovieTrack.EstimatedDataRate]: The estimated data rate, in bits per second, of the media that the track references.
//   - [AVMutableMovieTrack.SetEstimatedDataRate]
//   - [AVMutableMovieTrack.SamplePresentationTimeForTrackTime]: Maps the specified track time through the appropriate time mapping and returns the resulting sample presentation time.
//
// # Accessing language support
//
//   - [AVMutableMovieTrack.LanguageCode]: The language code of the track.
//   - [AVMutableMovieTrack.SetLanguageCode]
//   - [AVMutableMovieTrack.ExtendedLanguageTag]: The language tag of the track.
//   - [AVMutableMovieTrack.SetExtendedLanguageTag]
//
// # Accessing visual characteristics
//
//   - [AVMutableMovieTrack.NaturalSize]: The dimensions used to display the visual media data for the track.
//   - [AVMutableMovieTrack.SetNaturalSize]
//   - [AVMutableMovieTrack.PreferredTransform]: The transform performed on the visual media data of the track for display purposes.
//   - [AVMutableMovieTrack.SetPreferredTransform]
//   - [AVMutableMovieTrack.Layer]: The layer level for the visual media of the track.
//   - [AVMutableMovieTrack.SetLayer]
//   - [AVMutableMovieTrack.CleanApertureDimensions]: The clean aperture dimension of the track.
//   - [AVMutableMovieTrack.SetCleanApertureDimensions]
//   - [AVMutableMovieTrack.ProductionApertureDimensions]: The production aperture dimensions of the track.
//   - [AVMutableMovieTrack.SetProductionApertureDimensions]
//   - [AVMutableMovieTrack.EncodedPixelsDimensions]: The encoded pixels dimensions of the track.
//   - [AVMutableMovieTrack.SetEncodedPixelsDimensions]
//
// # Accessing audible characteristics
//
//   - [AVMutableMovieTrack.PreferredVolume]: The preferred volume for the audible medata data of the track.
//   - [AVMutableMovieTrack.SetPreferredVolume]
//   - [AVMutableMovieTrack.HasAudioSampleDependencies]: A Boolean value that indicates whether the track has sample dependencies.
//   - [AVMutableMovieTrack.SetHasAudioSampleDependencies]
//
// # Accessing frame-based characteristics
//
//   - [AVMutableMovieTrack.NominalFrameRate]: The frame rate of the track, in frames per second.
//   - [AVMutableMovieTrack.SetNominalFrameRate]
//   - [AVMutableMovieTrack.MinFrameDuration]: The minimum duration of the track’s frames.
//   - [AVMutableMovieTrack.SetMinFrameDuration]
//   - [AVMutableMovieTrack.RequiresFrameReordering]: A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
//   - [AVMutableMovieTrack.SetRequiresFrameReordering]
//
// # Accessing metadata
//
//   - [AVMutableMovieTrack.Metadata]: An array of metadata stored by the track.
//   - [AVMutableMovieTrack.SetMetadata]
//   - [AVMutableMovieTrack.CommonMetadata]: An array of metadata items for all common metadata keys that have a value.
//   - [AVMutableMovieTrack.SetCommonMetadata]
//   - [AVMutableMovieTrack.AvailableMetadataFormats]: An array of metadata formats available for the track.
//   - [AVMutableMovieTrack.SetAvailableMetadataFormats]
//   - [AVMutableMovieTrack.MetadataForFormat]: Returns metadata items that a track contains for the specified format.
//
// # Accessing track segments
//
//   - [AVMutableMovieTrack.Segments]: The time mappings from the track’s media samples to its timeline.
//   - [AVMutableMovieTrack.SetSegments]
//   - [AVMutableMovieTrack.SegmentForTrackTime]: Returns a segment whose target time range contains, or is closest to, the specified track time.
//
// # Managing track associations
//
//   - [AVMutableMovieTrack.AvailableTrackAssociationTypes]: An array of association types that the track uses to associate with other tracks.
//   - [AVMutableMovieTrack.SetAvailableTrackAssociationTypes]
//   - [AVMutableMovieTrack.AssociatedTracksOfType]: Returns an array of associated tracks that have the specified association type.
//   - [AVMutableMovieTrack.AddTrackAssociationToTrackType]: Creates a specific type of track association between two tracks.
//   - [AVMutableMovieTrack.RemoveTrackAssociationToTrackType]: Removes a specific type of track association between two tracks.
//
// # Determining sample cursor support
//
//   - [AVMutableMovieTrack.CanProvideSampleCursors]: A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
//   - [AVMutableMovieTrack.SetCanProvideSampleCursors]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack
type AVMutableMovieTrack struct {
	AVMovieTrack
}

// AVMutableMovieTrackFromID constructs a [AVMutableMovieTrack] from an objc.ID.
//
// A mutable track that conforms to the QuickTime or ISO base media file
// format.
func AVMutableMovieTrackFromID(id objc.ID) AVMutableMovieTrack {
	return AVMutableMovieTrack{AVMovieTrack: AVMovieTrackFromID(id)}
}
// NOTE: AVMutableMovieTrack adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMutableMovieTrack] class.
//
// # Managing time ranges
//
//   - [IAVMutableMovieTrack.InsertTimeRangeOfTrackAtTimeCopySampleDataError]: Inserts a portion of an asset track into the target movie.
//   - [IAVMutableMovieTrack.InsertEmptyTimeRange]: Adds an empty time range to a track.
//   - [IAVMutableMovieTrack.RemoveTimeRange]: Removes the specified time range from a track.
//   - [IAVMutableMovieTrack.ScaleTimeRangeToDuration]: Changes the duration of a time range in a track.
//
// # Appending sample data
//
//   - [IAVMutableMovieTrack.InsertMediaTimeRangeIntoTimeRange]: Inserts a reference to a media time range into a track.
//
// # Accessing media chunks
//
//   - [IAVMutableMovieTrack.PreferredMediaChunkAlignment]: The boundary for media chunk alignment for file types that support media chunk alignment.
//   - [IAVMutableMovieTrack.SetPreferredMediaChunkAlignment]
//   - [IAVMutableMovieTrack.PreferredMediaChunkDuration]: The maximum duration to use for each chunk of sample data written to the file for file types that support media chunk duration.
//   - [IAVMutableMovieTrack.SetPreferredMediaChunkDuration]
//   - [IAVMutableMovieTrack.PreferredMediaChunkSize]: The maximum size to use for each chunk of sample data written to the file for file types that support media chunk duration.
//   - [IAVMutableMovieTrack.SetPreferredMediaChunkSize]
//
// # Changing format descriptions
//
//   - [IAVMutableMovieTrack.FormatDescriptions]: The format descriptions of the media samples that a track references.
//   - [IAVMutableMovieTrack.SetFormatDescriptions]
//   - [IAVMutableMovieTrack.ReplaceFormatDescriptionWithFormatDescription]: Replaces the track’s format description with a new format description.
//
// # Configuring track information
//
//   - [IAVMutableMovieTrack.Modified]: A Boolean value that indicates whether a track is in a modified state.
//   - [IAVMutableMovieTrack.SetModified]
//   - [IAVMutableMovieTrack.SampleReferenceBaseURL]: The base URL for sample references.
//   - [IAVMutableMovieTrack.SetSampleReferenceBaseURL]
//
// # Accessing track information
//
//   - [IAVMutableMovieTrack.IsPlayable]: A Boolean value that indicates whether the track is playable in the current environment.
//   - [IAVMutableMovieTrack.SetIsPlayable]
//   - [IAVMutableMovieTrack.IsDecodable]: A Boolean value that indicates whether the track is decodable in the current environment.
//   - [IAVMutableMovieTrack.SetIsDecodable]
//   - [IAVMutableMovieTrack.Enabled]: A Boolean value that indicates whether the track’s container enables it.
//   - [IAVMutableMovieTrack.SetEnabled]
//   - [IAVMutableMovieTrack.IsSelfContained]: A Boolean value that indicates whether this track references sample data only within its container file.
//   - [IAVMutableMovieTrack.SetIsSelfContained]
//   - [IAVMutableMovieTrack.HasProtectedContent]: A Boolean value that indicates whether a track contains protected content.
//   - [IAVMutableMovieTrack.TotalSampleDataLength]: The total number of bytes of sample data the track requires.
//   - [IAVMutableMovieTrack.SetTotalSampleDataLength]
//   - [IAVMutableMovieTrack.HasMediaCharacteristic]: Returns a Boolean value that indicates whether the track references media with the specified media characteristic.
//
// # Accessing temporal information
//
//   - [IAVMutableMovieTrack.TimeRange]: The time range of the track within the overall timeline of the asset.
//   - [IAVMutableMovieTrack.SetTimeRange]
//   - [IAVMutableMovieTrack.Timescale]: The time scale for tracks that contain the `moov` atom.
//   - [IAVMutableMovieTrack.SetTimescale]
//   - [IAVMutableMovieTrack.NaturalTimeScale]: The natural time scale of the media that a track references.
//   - [IAVMutableMovieTrack.SetNaturalTimeScale]
//   - [IAVMutableMovieTrack.EstimatedDataRate]: The estimated data rate, in bits per second, of the media that the track references.
//   - [IAVMutableMovieTrack.SetEstimatedDataRate]
//   - [IAVMutableMovieTrack.SamplePresentationTimeForTrackTime]: Maps the specified track time through the appropriate time mapping and returns the resulting sample presentation time.
//
// # Accessing language support
//
//   - [IAVMutableMovieTrack.LanguageCode]: The language code of the track.
//   - [IAVMutableMovieTrack.SetLanguageCode]
//   - [IAVMutableMovieTrack.ExtendedLanguageTag]: The language tag of the track.
//   - [IAVMutableMovieTrack.SetExtendedLanguageTag]
//
// # Accessing visual characteristics
//
//   - [IAVMutableMovieTrack.NaturalSize]: The dimensions used to display the visual media data for the track.
//   - [IAVMutableMovieTrack.SetNaturalSize]
//   - [IAVMutableMovieTrack.PreferredTransform]: The transform performed on the visual media data of the track for display purposes.
//   - [IAVMutableMovieTrack.SetPreferredTransform]
//   - [IAVMutableMovieTrack.Layer]: The layer level for the visual media of the track.
//   - [IAVMutableMovieTrack.SetLayer]
//   - [IAVMutableMovieTrack.CleanApertureDimensions]: The clean aperture dimension of the track.
//   - [IAVMutableMovieTrack.SetCleanApertureDimensions]
//   - [IAVMutableMovieTrack.ProductionApertureDimensions]: The production aperture dimensions of the track.
//   - [IAVMutableMovieTrack.SetProductionApertureDimensions]
//   - [IAVMutableMovieTrack.EncodedPixelsDimensions]: The encoded pixels dimensions of the track.
//   - [IAVMutableMovieTrack.SetEncodedPixelsDimensions]
//
// # Accessing audible characteristics
//
//   - [IAVMutableMovieTrack.PreferredVolume]: The preferred volume for the audible medata data of the track.
//   - [IAVMutableMovieTrack.SetPreferredVolume]
//   - [IAVMutableMovieTrack.HasAudioSampleDependencies]: A Boolean value that indicates whether the track has sample dependencies.
//   - [IAVMutableMovieTrack.SetHasAudioSampleDependencies]
//
// # Accessing frame-based characteristics
//
//   - [IAVMutableMovieTrack.NominalFrameRate]: The frame rate of the track, in frames per second.
//   - [IAVMutableMovieTrack.SetNominalFrameRate]
//   - [IAVMutableMovieTrack.MinFrameDuration]: The minimum duration of the track’s frames.
//   - [IAVMutableMovieTrack.SetMinFrameDuration]
//   - [IAVMutableMovieTrack.RequiresFrameReordering]: A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
//   - [IAVMutableMovieTrack.SetRequiresFrameReordering]
//
// # Accessing metadata
//
//   - [IAVMutableMovieTrack.Metadata]: An array of metadata stored by the track.
//   - [IAVMutableMovieTrack.SetMetadata]
//   - [IAVMutableMovieTrack.CommonMetadata]: An array of metadata items for all common metadata keys that have a value.
//   - [IAVMutableMovieTrack.SetCommonMetadata]
//   - [IAVMutableMovieTrack.AvailableMetadataFormats]: An array of metadata formats available for the track.
//   - [IAVMutableMovieTrack.SetAvailableMetadataFormats]
//   - [IAVMutableMovieTrack.MetadataForFormat]: Returns metadata items that a track contains for the specified format.
//
// # Accessing track segments
//
//   - [IAVMutableMovieTrack.Segments]: The time mappings from the track’s media samples to its timeline.
//   - [IAVMutableMovieTrack.SetSegments]
//   - [IAVMutableMovieTrack.SegmentForTrackTime]: Returns a segment whose target time range contains, or is closest to, the specified track time.
//
// # Managing track associations
//
//   - [IAVMutableMovieTrack.AvailableTrackAssociationTypes]: An array of association types that the track uses to associate with other tracks.
//   - [IAVMutableMovieTrack.SetAvailableTrackAssociationTypes]
//   - [IAVMutableMovieTrack.AssociatedTracksOfType]: Returns an array of associated tracks that have the specified association type.
//   - [IAVMutableMovieTrack.AddTrackAssociationToTrackType]: Creates a specific type of track association between two tracks.
//   - [IAVMutableMovieTrack.RemoveTrackAssociationToTrackType]: Removes a specific type of track association between two tracks.
//
// # Determining sample cursor support
//
//   - [IAVMutableMovieTrack.CanProvideSampleCursors]: A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
//   - [IAVMutableMovieTrack.SetCanProvideSampleCursors]
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack
type IAVMutableMovieTrack interface {
	IAVMovieTrack

	// Topic: Managing time ranges

	// Inserts a portion of an asset track into the target movie.
	InsertTimeRangeOfTrackAtTimeCopySampleDataError(timeRange objectivec.IObject, track IAVAssetTrack, startTime objectivec.IObject, copySampleData bool) (bool, error)
	// Adds an empty time range to a track.
	InsertEmptyTimeRange(timeRange objectivec.IObject)
	// Removes the specified time range from a track.
	RemoveTimeRange(timeRange objectivec.IObject)
	// Changes the duration of a time range in a track.
	ScaleTimeRangeToDuration(timeRange objectivec.IObject, duration objectivec.IObject)

	// Topic: Appending sample data

	// Inserts a reference to a media time range into a track.
	InsertMediaTimeRangeIntoTimeRange(mediaTimeRange objectivec.IObject, trackTimeRange objectivec.IObject) bool

	// Topic: Accessing media chunks

	// The boundary for media chunk alignment for file types that support media chunk alignment.
	PreferredMediaChunkAlignment() int
	SetPreferredMediaChunkAlignment(value int)
	// The maximum duration to use for each chunk of sample data written to the file for file types that support media chunk duration.
	PreferredMediaChunkDuration() objectivec.IObject
	SetPreferredMediaChunkDuration(value objectivec.IObject)
	// The maximum size to use for each chunk of sample data written to the file for file types that support media chunk duration.
	PreferredMediaChunkSize() int
	SetPreferredMediaChunkSize(value int)

	// Topic: Changing format descriptions

	// The format descriptions of the media samples that a track references.
	FormatDescriptions() objectivec.IObject
	SetFormatDescriptions(value objectivec.IObject)
	// Replaces the track’s format description with a new format description.
	ReplaceFormatDescriptionWithFormatDescription(formatDescription objectivec.IObject, newFormatDescription objectivec.IObject)

	// Topic: Configuring track information

	// A Boolean value that indicates whether a track is in a modified state.
	Modified() bool
	SetModified(value bool)
	// The base URL for sample references.
	SampleReferenceBaseURL() foundation.INSURL
	SetSampleReferenceBaseURL(value foundation.INSURL)

	// Topic: Accessing track information

	// A Boolean value that indicates whether the track is playable in the current environment.
	IsPlayable() bool
	SetIsPlayable(value bool)
	// A Boolean value that indicates whether the track is decodable in the current environment.
	IsDecodable() bool
	SetIsDecodable(value bool)
	// A Boolean value that indicates whether the track’s container enables it.
	Enabled() bool
	SetEnabled(value bool)
	// A Boolean value that indicates whether this track references sample data only within its container file.
	IsSelfContained() bool
	SetIsSelfContained(value bool)
	// A Boolean value that indicates whether a track contains protected content.
	HasProtectedContent() bool
	// The total number of bytes of sample data the track requires.
	TotalSampleDataLength() objectivec.IObject
	SetTotalSampleDataLength(value objectivec.IObject)
	// Returns a Boolean value that indicates whether the track references media with the specified media characteristic.
	HasMediaCharacteristic(mediaCharacteristic AVMediaCharacteristic) bool

	// Topic: Accessing temporal information

	// The time range of the track within the overall timeline of the asset.
	TimeRange() objectivec.IObject
	SetTimeRange(value objectivec.IObject)
	// The time scale for tracks that contain the `moov` atom.
	Timescale() int32
	SetTimescale(value int32)
	// The natural time scale of the media that a track references.
	NaturalTimeScale() int32
	SetNaturalTimeScale(value int32)
	// The estimated data rate, in bits per second, of the media that the track references.
	EstimatedDataRate() float32
	SetEstimatedDataRate(value float32)
	// Maps the specified track time through the appropriate time mapping and returns the resulting sample presentation time.
	SamplePresentationTimeForTrackTime(trackTime objectivec.IObject) objectivec.IObject

	// Topic: Accessing language support

	// The language code of the track.
	LanguageCode() string
	SetLanguageCode(value string)
	// The language tag of the track.
	ExtendedLanguageTag() string
	SetExtendedLanguageTag(value string)

	// Topic: Accessing visual characteristics

	// The dimensions used to display the visual media data for the track.
	NaturalSize() corefoundation.CGSize
	SetNaturalSize(value corefoundation.CGSize)
	// The transform performed on the visual media data of the track for display purposes.
	PreferredTransform() corefoundation.CGAffineTransform
	SetPreferredTransform(value corefoundation.CGAffineTransform)
	// The layer level for the visual media of the track.
	Layer() int
	SetLayer(value int)
	// The clean aperture dimension of the track.
	CleanApertureDimensions() corefoundation.CGSize
	SetCleanApertureDimensions(value corefoundation.CGSize)
	// The production aperture dimensions of the track.
	ProductionApertureDimensions() corefoundation.CGSize
	SetProductionApertureDimensions(value corefoundation.CGSize)
	// The encoded pixels dimensions of the track.
	EncodedPixelsDimensions() corefoundation.CGSize
	SetEncodedPixelsDimensions(value corefoundation.CGSize)

	// Topic: Accessing audible characteristics

	// The preferred volume for the audible medata data of the track.
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
	MinFrameDuration() objectivec.IObject
	SetMinFrameDuration(value objectivec.IObject)
	// A Boolean value that indicates whether samples in the track may have different presentation and decode timestamps.
	RequiresFrameReordering() bool
	SetRequiresFrameReordering(value bool)

	// Topic: Accessing metadata

	// An array of metadata stored by the track.
	Metadata() []AVMetadataItem
	SetMetadata(value []AVMetadataItem)
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
	Segments() IAVAssetTrackSegment
	SetSegments(value IAVAssetTrackSegment)
	// Returns a segment whose target time range contains, or is closest to, the specified track time.
	SegmentForTrackTime(trackTime objectivec.IObject) IAVAssetTrackSegment

	// Topic: Managing track associations

	// An array of association types that the track uses to associate with other tracks.
	AvailableTrackAssociationTypes() objc.ID
	SetAvailableTrackAssociationTypes(value objc.ID)
	// Returns an array of associated tracks that have the specified association type.
	AssociatedTracksOfType(trackAssociationType AVTrackAssociationType) []AVAssetTrack
	// Creates a specific type of track association between two tracks.
	AddTrackAssociationToTrackType(movieTrack IAVMovieTrack, trackAssociationType AVTrackAssociationType)
	// Removes a specific type of track association between two tracks.
	RemoveTrackAssociationToTrackType(movieTrack IAVMovieTrack, trackAssociationType AVTrackAssociationType)

	// Topic: Determining sample cursor support

	// A Boolean value that indicates whether the track can provide instances of sample cursors to traverse its media samples and discover information.
	CanProvideSampleCursors() bool
	SetCanProvideSampleCursors(value bool)
}

// Init initializes the instance.
func (m AVMutableMovieTrack) Init() AVMutableMovieTrack {
	rv := objc.Send[AVMutableMovieTrack](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMutableMovieTrack) Autorelease() AVMutableMovieTrack {
	rv := objc.Send[AVMutableMovieTrack](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMutableMovieTrack creates a new AVMutableMovieTrack instance.
func NewAVMutableMovieTrack() AVMutableMovieTrack {
	class := getAVMutableMovieTrackClass()
	rv := objc.Send[AVMutableMovieTrack](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Inserts a portion of an asset track into the target movie.
//
// timeRange: The time range of the track to insert.
//
// track: An [AVAssetTrack] object indicating the source of the inserted media. This
// value can’t be `nil`.
//
// startTime: The time in the target track at which the media is to be inserted.
//
// copySampleData: A Boolean value that indicates whether sample data is to be copied from the
// source to the destination during edits. If [YES], the sample data is
// written to the location specified by the track property `mediaDataStorage`
// if non-nil, or else by the movie property `defaultMediaDataStorage` if
// non-nil; if both are nil, the method fails and returns [NO]. If [NO],
// sample data isn’t written and sample references to the samples in their
// original container are added as necessary.
//
// timeRange is a [coremedia.CMTimeRange].
//
// startTime is a [coremedia.CMTime].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/insertTimeRange(_:of:at:copySampleData:)
func (m AVMutableMovieTrack) InsertTimeRangeOfTrackAtTimeCopySampleDataError(timeRange objectivec.IObject, track IAVAssetTrack, startTime objectivec.IObject, copySampleData bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("insertTimeRange:ofTrack:atTime:copySampleData:error:"), timeRange, track, startTime, copySampleData, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("insertTimeRange:ofTrack:atTime:copySampleData:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Adds an empty time range to a track.
//
// timeRange: A time range to insert.
//
// timeRange is a [coremedia.CMTimeRange].
//
// # Discussion
// 
// You can’t add empty time ranges to the end of a track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/insertEmptyTimeRange(_:)
func (m AVMutableMovieTrack) InsertEmptyTimeRange(timeRange objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("insertEmptyTimeRange:"), timeRange)
}
// Removes the specified time range from a track.
//
// timeRange: The time range to remove.
//
// timeRange is a [coremedia.CMTimeRange].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/removeTimeRange(_:)
func (m AVMutableMovieTrack) RemoveTimeRange(timeRange objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeTimeRange:"), timeRange)
}
// Changes the duration of a time range in a track.
//
// timeRange: The time range to change.
//
// duration: The new duration for the time range.
//
// timeRange is a [coremedia.CMTimeRange].
//
// duration is a [coremedia.CMTime].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/scaleTimeRange(_:toDuration:)
func (m AVMutableMovieTrack) ScaleTimeRangeToDuration(timeRange objectivec.IObject, duration objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("scaleTimeRange:toDuration:"), timeRange, duration)
}
// Inserts a reference to a media time range into a track.
//
// mediaTimeRange: The presentation time range of the media to be inserted.
//
// trackTimeRange: The time range of the track into which the media is to be inserted.
//
// mediaTimeRange is a [coremedia.CMTimeRange].
//
// trackTimeRange is a [coremedia.CMTimeRange].
//
// # Return Value
// 
// A Boolean value that indicates whether the insertion was successful.
//
// # Discussion
// 
// Use this method after appending samples or sample references to a track’s
// media. To specify that the media time range be played at its natural rate,
// pass `mediaTimeRange.Duration() == trackTimeRange.Duration()`; otherwise,
// the ratio between these is used to determine the playback rate. Pass
// [invalid] for `trackTimeRange.Start()` to indicate that the segment should
// be appended to the end of the track.
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/insertMediaTimeRange(_:into:)
func (m AVMutableMovieTrack) InsertMediaTimeRangeIntoTimeRange(mediaTimeRange objectivec.IObject, trackTimeRange objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("insertMediaTimeRange:intoTimeRange:"), mediaTimeRange, trackTimeRange)
	return rv
}
// Replaces the track’s format description with a new format description.
//
// formatDescription: The [CMFormatDescription] object to be replaced.
// //
// [CMFormatDescription]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescription
//
// newFormatDescription: The [CMFormatDescription] object to replacing the specified format
// description.
// //
// [CMFormatDescription]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescription
//
// formatDescription is a [coremedia.CMFormatDescriptionRef].
//
// newFormatDescription is a [coremedia.CMFormatDescriptionRef].
//
// # Discussion
// 
// Use this method to change a track’s format descriptions, such as adding
// format description extensions to a format description or changing the audio
// channel layout of an audio track. Format description can have extensions of
// type [kCMFormatDescriptionExtension_VerbatimSampleDescription] and
// [kCMFormatDescriptionExtension_VerbatimISOSampleEntry]. If you modify a
// copy of a format description, delete those extensions from the copy or your
// changes might be ignored.
//
// [kCMFormatDescriptionExtension_VerbatimISOSampleEntry]: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_VerbatimISOSampleEntry
// [kCMFormatDescriptionExtension_VerbatimSampleDescription]: https://developer.apple.com/documentation/CoreMedia/kCMFormatDescriptionExtension_VerbatimSampleDescription
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/replaceFormatDescription(_:with:)
func (m AVMutableMovieTrack) ReplaceFormatDescriptionWithFormatDescription(formatDescription objectivec.IObject, newFormatDescription objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("replaceFormatDescription:withFormatDescription:"), formatDescription, newFormatDescription)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/hasMediaCharacteristic(_:)
func (m AVMutableMovieTrack) HasMediaCharacteristic(mediaCharacteristic AVMediaCharacteristic) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasMediaCharacteristic:"), objc.String(string(mediaCharacteristic)))
	return rv
}
// Maps the specified track time through the appropriate time mapping and
// returns the resulting sample presentation time.
//
// trackTime: The track time for which to request the sample presentation time.
//
// trackTime is a [coremedia.CMTime].
//
// # Return Value
// 
// The sample presentation time corresponding to the specified time; otherwise
// [invalid] if the time is out of range.
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/samplePresentationTime(forTrackTime:)
func (m AVMutableMovieTrack) SamplePresentationTimeForTrackTime(trackTime objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("samplePresentationTimeForTrackTime:"), trackTime)
	return objectivec.Object{ID: rv}
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/metadata(forFormat:)
func (m AVMutableMovieTrack) MetadataForFormat(format AVMetadataFormat) []AVMetadataItem {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("metadataForFormat:"), objc.String(string(format)))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}
// Returns a segment whose target time range contains, or is closest to, the
// specified track time.
//
// trackTime: The track time of the segment to return.
//
// trackTime is a [coremedia.CMTime].
//
// # Return Value
// 
// The [AVCompositionTrackSegment] associated with the track time.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/segment(forTrackTime:)
func (m AVMutableMovieTrack) SegmentForTrackTime(trackTime objectivec.IObject) IAVAssetTrackSegment {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("segmentForTrackTime:"), trackTime)
	return AVAssetTrackSegmentFromID(rv)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/associatedTracks(ofType:)
func (m AVMutableMovieTrack) AssociatedTracksOfType(trackAssociationType AVTrackAssociationType) []AVAssetTrack {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("associatedTracksOfType:"), objc.String(string(trackAssociationType)))
	return objc.ConvertSlice(rv, func(id objc.ID) AVAssetTrack {
		return AVAssetTrackFromID(id)
	})
}
// Creates a specific type of track association between two tracks.
//
// movieTrack: The AVMovieTrack object to be associated with the receiver.
//
// trackAssociationType: The type of track association to add between the receiver and the specified
// movie track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/addTrackAssociation(to:type:)
func (m AVMutableMovieTrack) AddTrackAssociationToTrackType(movieTrack IAVMovieTrack, trackAssociationType AVTrackAssociationType) {
	objc.Send[objc.ID](m.ID, objc.Sel("addTrackAssociationToTrack:type:"), movieTrack, objc.String(string(trackAssociationType)))
}
// Removes a specific type of track association between two tracks.
//
// movieTrack: The AVMovieTrack object that is associated with the receiver.
//
// trackAssociationType: The type of track association to remove between the receiver and the
// specified movie track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/removeTrackAssociation(to:type:)
func (m AVMutableMovieTrack) RemoveTrackAssociationToTrackType(movieTrack IAVMovieTrack, trackAssociationType AVTrackAssociationType) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeTrackAssociationToTrack:type:"), movieTrack, objc.String(string(trackAssociationType)))
}

// The boundary for media chunk alignment for file types that support media
// chunk alignment.
//
// # Discussion
// 
// The default value is `0`, which indicates to use no padding should to
// achieve chunk alignment. Setting a negative chunk alignment value causes an
// error.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/preferredMediaChunkAlignment
func (m AVMutableMovieTrack) PreferredMediaChunkAlignment() int {
	rv := objc.Send[int](m.ID, objc.Sel("preferredMediaChunkAlignment"))
	return rv
}
func (m AVMutableMovieTrack) SetPreferredMediaChunkAlignment(value int) {
	objc.Send[struct{}](m.ID, objc.Sel("setPreferredMediaChunkAlignment:"), value)
}
// The maximum duration to use for each chunk of sample data written to the
// file for file types that support media chunk duration.
//
// # Discussion
// 
// The total duration of the samples in a chunk can be no greater than the
// preferred chunk duration, or the duration of a single sample if the single
// sample’s duration is greater than the preferred chunk duration. The
// default media chunk duration is `1.0` second. Setting a negative or
// non-numeric value for the chunk duration will cause an error.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/preferredMediaChunkDuration
func (m AVMutableMovieTrack) PreferredMediaChunkDuration() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("preferredMediaChunkDuration"))
	return objectivec.Object{ID: rv}
}
func (m AVMutableMovieTrack) SetPreferredMediaChunkDuration(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setPreferredMediaChunkDuration:"), value)
}
// The maximum size to use for each chunk of sample data written to the file
// for file types that support media chunk duration.
//
// # Discussion
// 
// The total size of the samples in a chunk can be no greater than the
// preferred chunk size, or the size of a single sample if the single
// sample’s size is greater than the preferred chunk size. The default media
// chunk duration is `1024 * 1024` bytes. Setting a negative value for the
// chunk duration will cause an error.
// 
// A larger chunk size can result in fewer reads from the storage container,
// at the potential expense of a larger memory footprint.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/preferredMediaChunkSize
func (m AVMutableMovieTrack) PreferredMediaChunkSize() int {
	rv := objc.Send[int](m.ID, objc.Sel("preferredMediaChunkSize"))
	return rv
}
func (m AVMutableMovieTrack) SetPreferredMediaChunkSize(value int) {
	objc.Send[struct{}](m.ID, objc.Sel("setPreferredMediaChunkSize:"), value)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/formatDescriptions
func (m AVMutableMovieTrack) FormatDescriptions() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("formatDescriptions"))
	return objectivec.Object{ID: rv}
}
func (m AVMutableMovieTrack) SetFormatDescriptions(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setFormatDescriptions:"), value)
}
// A Boolean value that indicates whether a track is in a modified state.
//
// # Discussion
// 
// This property is [YES] when the [AVMutableMovieTrack] object has been
// modified since it was created, was last written, or had its modified state
// cleared.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/isModified
func (m AVMutableMovieTrack) Modified() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isModified"))
	return rv
}
func (m AVMutableMovieTrack) SetModified(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setModified:"), value)
}
// The base URL for sample references.
//
// # Discussion
// 
// When this property is an absolute URL, the sample locations written to the
// file when appending sample references to this track are relative to this
// URL. The URL must point to a location contained by any common parent
// directory of the locations that will be referenced. For example, setting
// the this property to `///Users/johnappleseed/Movies/` and appending sample
// buffers that refer to `///Users/johnappleseed/Movies/data/movie1.Mov()`
// will cause the sample reference `data/movie1.Mov()` to be written to the
// movie file.
// 
// If this property can’t be resolved as an absolute URL or if it points to
// a location that isn’t contained by any common parent directory of the
// locations referenced, the unmodified location is written. The default value
// is `nil`, indicating that the unmodified location will be written.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/sampleReferenceBaseURL
func (m AVMutableMovieTrack) SampleReferenceBaseURL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("sampleReferenceBaseURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (m AVMutableMovieTrack) SetSampleReferenceBaseURL(value foundation.INSURL) {
	objc.Send[struct{}](m.ID, objc.Sel("setSampleReferenceBaseURL:"), value)
}
// A Boolean value that indicates whether the track is playable in the current
// environment.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/isPlayable
func (m AVMutableMovieTrack) IsPlayable() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isPlayable"))
	return rv
}
func (m AVMutableMovieTrack) SetIsPlayable(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setPlayable:"), value)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/isDecodable
func (m AVMutableMovieTrack) IsDecodable() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isDecodable"))
	return rv
}
func (m AVMutableMovieTrack) SetIsDecodable(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setDecodable:"), value)
}
// A Boolean value that indicates whether the track’s container enables it.
//
// # Discussion
// 
// For file-based media, you can change its [Enabled] presentation state using
// [AVPlayerItemTrack].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/isEnabled
func (m AVMutableMovieTrack) Enabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isEnabled"))
	return rv
}
func (m AVMutableMovieTrack) SetEnabled(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setEnabled:"), value)
}
// A Boolean value that indicates whether this track references sample data
// only within its container file.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/isSelfContained
func (m AVMutableMovieTrack) IsSelfContained() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isSelfContained"))
	return rv
}
func (m AVMutableMovieTrack) SetIsSelfContained(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setSelfContained:"), value)
}
// A Boolean value that indicates whether a track contains protected content.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/hasProtectedContent
func (m AVMutableMovieTrack) HasProtectedContent() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasProtectedContent"))
	return rv
}
// The total number of bytes of sample data the track requires.
//
// # Discussion
// 
// The value may be `0` if the framework can’t determine the total sample
// data length.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/totalSampleDataLength
func (m AVMutableMovieTrack) TotalSampleDataLength() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("totalSampleDataLength"))
	return objectivec.Object{ID: rv}
}
func (m AVMutableMovieTrack) SetTotalSampleDataLength(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setTotalSampleDataLength:"), value)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/timeRange
func (m AVMutableMovieTrack) TimeRange() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("timeRange"))
	return objectivec.Object{ID: rv}
}
func (m AVMutableMovieTrack) SetTimeRange(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setTimeRange:"), value)
}
// The time scale for tracks that contain the `moov` atom.
//
// # Discussion
// 
// The default media time is `0`. Set this property on any new, empty tracks
// before any edits are performed on the track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/timescale
func (m AVMutableMovieTrack) Timescale() int32 {
	rv := objc.Send[int32](m.ID, objc.Sel("timescale"))
	return rv
}
func (m AVMutableMovieTrack) SetTimescale(value int32) {
	objc.Send[struct{}](m.ID, objc.Sel("setTimescale:"), value)
}
// The natural time scale of the media that a track references.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/naturalTimeScale
func (m AVMutableMovieTrack) NaturalTimeScale() int32 {
	rv := objc.Send[int32](m.ID, objc.Sel("naturalTimeScale"))
	return rv
}
func (m AVMutableMovieTrack) SetNaturalTimeScale(value int32) {
	objc.Send[struct{}](m.ID, objc.Sel("setNaturalTimeScale:"), value)
}
// The estimated data rate, in bits per second, of the media that the track
// references.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/estimatedDataRate
func (m AVMutableMovieTrack) EstimatedDataRate() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("estimatedDataRate"))
	return rv
}
func (m AVMutableMovieTrack) SetEstimatedDataRate(value float32) {
	objc.Send[struct{}](m.ID, objc.Sel("setEstimatedDataRate:"), value)
}
// The language code of the track.
//
// # Discussion
// 
// The value is an ISO 639-2/T language code, or `nil` if the track doesn’t
// specify a language code.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/languageCode
func (m AVMutableMovieTrack) LanguageCode() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("languageCode"))
	return foundation.NSStringFromID(rv).String()
}
func (m AVMutableMovieTrack) SetLanguageCode(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setLanguageCode:"), objc.String(value))
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/extendedLanguageTag
func (m AVMutableMovieTrack) ExtendedLanguageTag() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("extendedLanguageTag"))
	return foundation.NSStringFromID(rv).String()
}
func (m AVMutableMovieTrack) SetExtendedLanguageTag(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setExtendedLanguageTag:"), objc.String(value))
}
// The dimensions used to display the visual media data for the track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/naturalSize
func (m AVMutableMovieTrack) NaturalSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m.ID, objc.Sel("naturalSize"))
	return corefoundation.CGSize(rv)
}
func (m AVMutableMovieTrack) SetNaturalSize(value corefoundation.CGSize) {
	objc.Send[struct{}](m.ID, objc.Sel("setNaturalSize:"), value)
}
// The transform performed on the visual media data of the track for display
// purposes.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/preferredTransform
func (m AVMutableMovieTrack) PreferredTransform() corefoundation.CGAffineTransform {
	rv := objc.Send[corefoundation.CGAffineTransform](m.ID, objc.Sel("preferredTransform"))
	return corefoundation.CGAffineTransform(rv)
}
func (m AVMutableMovieTrack) SetPreferredTransform(value corefoundation.CGAffineTransform) {
	objc.Send[struct{}](m.ID, objc.Sel("setPreferredTransform:"), value)
}
// The layer level for the visual media of the track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/layer
func (m AVMutableMovieTrack) Layer() int {
	rv := objc.Send[int](m.ID, objc.Sel("layer"))
	return rv
}
func (m AVMutableMovieTrack) SetLayer(value int) {
	objc.Send[struct{}](m.ID, objc.Sel("setLayer:"), value)
}
// The clean aperture dimension of the track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/cleanApertureDimensions
func (m AVMutableMovieTrack) CleanApertureDimensions() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m.ID, objc.Sel("cleanApertureDimensions"))
	return corefoundation.CGSize(rv)
}
func (m AVMutableMovieTrack) SetCleanApertureDimensions(value corefoundation.CGSize) {
	objc.Send[struct{}](m.ID, objc.Sel("setCleanApertureDimensions:"), value)
}
// The production aperture dimensions of the track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/productionApertureDimensions
func (m AVMutableMovieTrack) ProductionApertureDimensions() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m.ID, objc.Sel("productionApertureDimensions"))
	return corefoundation.CGSize(rv)
}
func (m AVMutableMovieTrack) SetProductionApertureDimensions(value corefoundation.CGSize) {
	objc.Send[struct{}](m.ID, objc.Sel("setProductionApertureDimensions:"), value)
}
// The encoded pixels dimensions of the track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/encodedPixelsDimensions
func (m AVMutableMovieTrack) EncodedPixelsDimensions() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m.ID, objc.Sel("encodedPixelsDimensions"))
	return corefoundation.CGSize(rv)
}
func (m AVMutableMovieTrack) SetEncodedPixelsDimensions(value corefoundation.CGSize) {
	objc.Send[struct{}](m.ID, objc.Sel("setEncodedPixelsDimensions:"), value)
}
// The preferred volume for the audible medata data of the track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/preferredVolume
func (m AVMutableMovieTrack) PreferredVolume() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("preferredVolume"))
	return rv
}
func (m AVMutableMovieTrack) SetPreferredVolume(value float32) {
	objc.Send[struct{}](m.ID, objc.Sel("setPreferredVolume:"), value)
}
// A Boolean value that indicates whether the track has sample dependencies.
//
// # Discussion
// 
// The value is always [false] for nonaudible media.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/hasAudioSampleDependencies
func (m AVMutableMovieTrack) HasAudioSampleDependencies() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasAudioSampleDependencies"))
	return rv
}
func (m AVMutableMovieTrack) SetHasAudioSampleDependencies(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setHasAudioSampleDependencies:"), value)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/nominalFrameRate
func (m AVMutableMovieTrack) NominalFrameRate() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("nominalFrameRate"))
	return rv
}
func (m AVMutableMovieTrack) SetNominalFrameRate(value float32) {
	objc.Send[struct{}](m.ID, objc.Sel("setNominalFrameRate:"), value)
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
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/minFrameDuration
func (m AVMutableMovieTrack) MinFrameDuration() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("minFrameDuration"))
	return objectivec.Object{ID: rv}
}
func (m AVMutableMovieTrack) SetMinFrameDuration(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setMinFrameDuration:"), value)
}
// A Boolean value that indicates whether samples in the track may have
// different presentation and decode timestamps.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/requiresFrameReordering
func (m AVMutableMovieTrack) RequiresFrameReordering() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("requiresFrameReordering"))
	return rv
}
func (m AVMutableMovieTrack) SetRequiresFrameReordering(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setRequiresFrameReordering:"), value)
}
// An array of metadata stored by the track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/metadata
func (m AVMutableMovieTrack) Metadata() []AVMetadataItem {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("metadata"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}
func (m AVMutableMovieTrack) SetMetadata(value []AVMetadataItem) {
	objc.Send[struct{}](m.ID, objc.Sel("setMetadata:"), objectivec.IObjectSliceToNSArray(value))
}
// An array of metadata items for all common metadata keys that have a value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/commonMetadata
func (m AVMutableMovieTrack) CommonMetadata() IAVMetadataItem {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("commonMetadata"))
	return AVMetadataItemFromID(objc.ID(rv))
}
func (m AVMutableMovieTrack) SetCommonMetadata(value IAVMetadataItem) {
	objc.Send[struct{}](m.ID, objc.Sel("setCommonMetadata:"), value)
}
// An array of metadata formats available for the track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/availableMetadataFormats
func (m AVMutableMovieTrack) AvailableMetadataFormats() AVMetadataFormat {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("availableMetadataFormats"))
	return AVMetadataFormat(foundation.NSStringFromID(rv).String())
}
func (m AVMutableMovieTrack) SetAvailableMetadataFormats(value AVMetadataFormat) {
	objc.Send[struct{}](m.ID, objc.Sel("setAvailableMetadataFormats:"), objc.String(string(value)))
}
// The time mappings from the track’s media samples to its timeline.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/segments
func (m AVMutableMovieTrack) Segments() IAVAssetTrackSegment {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("segments"))
	return AVAssetTrackSegmentFromID(objc.ID(rv))
}
func (m AVMutableMovieTrack) SetSegments(value IAVAssetTrackSegment) {
	objc.Send[struct{}](m.ID, objc.Sel("setSegments:"), value)
}
// An array of association types that the track uses to associate with other
// tracks.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/availableTrackAssociationTypes
func (m AVMutableMovieTrack) AvailableTrackAssociationTypes() objc.ID {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("availableTrackAssociationTypes"))
	return rv
}
func (m AVMutableMovieTrack) SetAvailableTrackAssociationTypes(value objc.ID) {
	objc.Send[struct{}](m.ID, objc.Sel("setAvailableTrackAssociationTypes:"), value)
}
// A Boolean value that indicates whether the track can provide instances of
// sample cursors to traverse its media samples and discover information.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableMovieTrack/canProvideSampleCursors
func (m AVMutableMovieTrack) CanProvideSampleCursors() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("canProvideSampleCursors"))
	return rv
}
func (m AVMutableMovieTrack) SetCanProvideSampleCursors(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setCanProvideSampleCursors:"), value)
}

