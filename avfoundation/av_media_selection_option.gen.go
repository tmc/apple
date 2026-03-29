// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMediaSelectionOption] class.
var (
	_AVMediaSelectionOptionClass     AVMediaSelectionOptionClass
	_AVMediaSelectionOptionClassOnce sync.Once
)

func getAVMediaSelectionOptionClass() AVMediaSelectionOptionClass {
	_AVMediaSelectionOptionClassOnce.Do(func() {
		_AVMediaSelectionOptionClass = AVMediaSelectionOptionClass{class: objc.GetClass("AVMediaSelectionOption")}
	})
	return _AVMediaSelectionOptionClass
}

// GetAVMediaSelectionOptionClass returns the class object for AVMediaSelectionOption.
func GetAVMediaSelectionOptionClass() AVMediaSelectionOptionClass {
	return getAVMediaSelectionOptionClass()
}

type AVMediaSelectionOptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMediaSelectionOptionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMediaSelectionOptionClass) Alloc() AVMediaSelectionOption {
	rv := objc.Send[AVMediaSelectionOption](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a specific option for the presentation of media
// within a group of options.
//
// # Accessing media information
//
//   - [AVMediaSelectionOption.MediaType]: The media type of the media data.
//   - [AVMediaSelectionOption.MediaSubTypes]: The media sub-types of the media data associated with the option.
//   - [AVMediaSelectionOption.HasMediaCharacteristic]: Returns a Boolean value that indicates whether the receiver has media with the given media characteristic.
//
// # Managing metadata
//
//   - [AVMediaSelectionOption.CommonMetadata]: An array of metadata items for each common metadata key for which a value is available.
//   - [AVMediaSelectionOption.AvailableMetadataFormats]: The metadata formats that contain metadata associated with the option.
//   - [AVMediaSelectionOption.MetadataForFormat]: Returns an array of metadata items—one for each metadata item in the container of a given format.
//
// # Determining playability
//
//   - [AVMediaSelectionOption.Playable]: A Boolean value that indicates whether the media selection option is playable.
//
// # Getting the language and locale settings
//
//   - [AVMediaSelectionOption.DisplayName]: A string suitable for display using the current system locale.
//   - [AVMediaSelectionOption.DisplayNameWithLocale]: Returns a string suitable for display using the specified locale.
//   - [AVMediaSelectionOption.Locale]: The locale for which the media option was authored.
//   - [AVMediaSelectionOption.ExtendedLanguageTag]: The IETF BCP 47 language tag associated with the option
//
// # Getting the associated media selection option
//
//   - [AVMediaSelectionOption.AssociatedMediaSelectionOptionInMediaSelectionGroup]: Returns a media selection option associated with the receiver in a given group.
//
// # Creating a Now Playing language option
//
//   - [AVMediaSelectionOption.MakeNowPlayingInfoLanguageOption]: Creates a language option for a media selection option.
//
// # Creating a property list representation
//
//   - [AVMediaSelectionOption.PropertyList]: Returns a serializable property list that’s sufficient to identify the option within its group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption
type AVMediaSelectionOption struct {
	objectivec.Object
}

// AVMediaSelectionOptionFromID constructs a [AVMediaSelectionOption] from an objc.ID.
//
// An object that represents a specific option for the presentation of media
// within a group of options.
func AVMediaSelectionOptionFromID(id objc.ID) AVMediaSelectionOption {
	return AVMediaSelectionOption{objectivec.Object{ID: id}}
}
// NOTE: AVMediaSelectionOption adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMediaSelectionOption] class.
//
// # Accessing media information
//
//   - [IAVMediaSelectionOption.MediaType]: The media type of the media data.
//   - [IAVMediaSelectionOption.MediaSubTypes]: The media sub-types of the media data associated with the option.
//   - [IAVMediaSelectionOption.HasMediaCharacteristic]: Returns a Boolean value that indicates whether the receiver has media with the given media characteristic.
//
// # Managing metadata
//
//   - [IAVMediaSelectionOption.CommonMetadata]: An array of metadata items for each common metadata key for which a value is available.
//   - [IAVMediaSelectionOption.AvailableMetadataFormats]: The metadata formats that contain metadata associated with the option.
//   - [IAVMediaSelectionOption.MetadataForFormat]: Returns an array of metadata items—one for each metadata item in the container of a given format.
//
// # Determining playability
//
//   - [IAVMediaSelectionOption.Playable]: A Boolean value that indicates whether the media selection option is playable.
//
// # Getting the language and locale settings
//
//   - [IAVMediaSelectionOption.DisplayName]: A string suitable for display using the current system locale.
//   - [IAVMediaSelectionOption.DisplayNameWithLocale]: Returns a string suitable for display using the specified locale.
//   - [IAVMediaSelectionOption.Locale]: The locale for which the media option was authored.
//   - [IAVMediaSelectionOption.ExtendedLanguageTag]: The IETF BCP 47 language tag associated with the option
//
// # Getting the associated media selection option
//
//   - [IAVMediaSelectionOption.AssociatedMediaSelectionOptionInMediaSelectionGroup]: Returns a media selection option associated with the receiver in a given group.
//
// # Creating a Now Playing language option
//
//   - [IAVMediaSelectionOption.MakeNowPlayingInfoLanguageOption]: Creates a language option for a media selection option.
//
// # Creating a property list representation
//
//   - [IAVMediaSelectionOption.PropertyList]: Returns a serializable property list that’s sufficient to identify the option within its group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption
type IAVMediaSelectionOption interface {
	objectivec.IObject

	// Topic: Accessing media information

	// The media type of the media data.
	MediaType() AVMediaType
	// The media sub-types of the media data associated with the option.
	MediaSubTypes() []foundation.NSNumber
	// Returns a Boolean value that indicates whether the receiver has media with the given media characteristic.
	HasMediaCharacteristic(mediaCharacteristic AVMediaCharacteristic) bool

	// Topic: Managing metadata

	// An array of metadata items for each common metadata key for which a value is available.
	CommonMetadata() []AVMetadataItem
	// The metadata formats that contain metadata associated with the option.
	AvailableMetadataFormats() []string
	// Returns an array of metadata items—one for each metadata item in the container of a given format.
	MetadataForFormat(format string) []AVMetadataItem

	// Topic: Determining playability

	// A Boolean value that indicates whether the media selection option is playable.
	Playable() bool

	// Topic: Getting the language and locale settings

	// A string suitable for display using the current system locale.
	DisplayName() string
	// Returns a string suitable for display using the specified locale.
	DisplayNameWithLocale(locale foundation.NSLocale) string
	// The locale for which the media option was authored.
	Locale() foundation.NSLocale
	// The IETF BCP 47 language tag associated with the option
	ExtendedLanguageTag() string

	// Topic: Getting the associated media selection option

	// Returns a media selection option associated with the receiver in a given group.
	AssociatedMediaSelectionOptionInMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) IAVMediaSelectionOption

	// Topic: Creating a Now Playing language option

	// Creates a language option for a media selection option.
	MakeNowPlayingInfoLanguageOption() objectivec.IObject

	// Topic: Creating a property list representation

	// Returns a serializable property list that’s sufficient to identify the option within its group.
	PropertyList() objectivec.IObject
}

// Init initializes the instance.
func (m AVMediaSelectionOption) Init() AVMediaSelectionOption {
	rv := objc.Send[AVMediaSelectionOption](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMediaSelectionOption) Autorelease() AVMediaSelectionOption {
	rv := objc.Send[AVMediaSelectionOption](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMediaSelectionOption creates a new AVMediaSelectionOption instance.
func NewAVMediaSelectionOption() AVMediaSelectionOption {
	class := getAVMediaSelectionOptionClass()
	rv := objc.Send[AVMediaSelectionOption](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a Boolean value that indicates whether the receiver has media with
// the given media characteristic.
//
// mediaCharacteristic: The media characteristic of interest, for example, [visual], [audible], or
// [legible].
// //
// [audible]: https://developer.apple.com/documentation/AVFoundation/AVMediaCharacteristic/audible
// [legible]: https://developer.apple.com/documentation/AVFoundation/AVMediaCharacteristic/legible
// [visual]: https://developer.apple.com/documentation/AVFoundation/AVMediaCharacteristic/visual
//
// # Return Value
// 
// [true] if the media selection option has media with mediaCharacteristic,
// otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/hasMediaCharacteristic(_:)
func (m AVMediaSelectionOption) HasMediaCharacteristic(mediaCharacteristic AVMediaCharacteristic) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasMediaCharacteristic:"), objc.String(string(mediaCharacteristic)))
	return rv
}
// Returns an array of metadata items—one for each metadata item in the
// container of a given format.
//
// format: The metadata format for which items are requested.
//
// # Return Value
// 
// An array of [AVMetadataItem] objects, one for each metadata item in the
// container of format, or `nil` if there is no metadata of the specified
// format.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/metadata(forFormat:)
func (m AVMediaSelectionOption) MetadataForFormat(format string) []AVMetadataItem {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("metadataForFormat:"), objc.String(format))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}
// Returns a string suitable for display using the specified locale.
//
// locale: The locale to use in generating the display name.
//
// # Return Value
// 
// A string containing the localized display name.
//
// # Discussion
// 
// The string takes into account this option’s common metadata, media
// characteristics and locale properties in addition to the provided locale to
// formulate a string intended for display
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/displayName(with:)
func (m AVMediaSelectionOption) DisplayNameWithLocale(locale foundation.NSLocale) string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("displayNameWithLocale:"), locale)
	return foundation.NSStringFromID(rv).String()
}
// Returns a media selection option associated with the receiver in a given
// group.
//
// mediaSelectionGroup: A media selection group in which an associated option is to be sought.
//
// # Return Value
// 
// A media selection option associated with the receiver in
// `mediaSelectionGroup`, or `nil` if none were found.
//
// # Discussion
// 
// Audible media selection options often have associated legible media
// selection options; in particular, audible options are typically associated
// with forced-only subtitle options with the same locale. See
// [containsOnlyForcedSubtitles] for a discussion of forced-only subtitles.
//
// [containsOnlyForcedSubtitles]: https://developer.apple.com/documentation/AVFoundation/AVMediaCharacteristic/containsOnlyForcedSubtitles
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/associatedMediaSelectionOption(in:)
func (m AVMediaSelectionOption) AssociatedMediaSelectionOptionInMediaSelectionGroup(mediaSelectionGroup IAVMediaSelectionGroup) IAVMediaSelectionOption {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("associatedMediaSelectionOptionInMediaSelectionGroup:"), mediaSelectionGroup)
	return AVMediaSelectionOptionFromID(rv)
}
// Creates a language option for a media selection option.
//
// # Return Value
// 
// A new language option, or `nil` if the media selection option isn’t an
// audible or legible option.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/makeNowPlayingInfoLanguageOption()
func (m AVMediaSelectionOption) MakeNowPlayingInfoLanguageOption() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("makeNowPlayingInfoLanguageOption"))
	return objectivec.Object{ID: rv}
}
// Returns a serializable property list that’s sufficient to identify the
// option within its group.
//
// # Return Value
// 
// A serializable property list that you can use to obtain an instance of
// [AVMediaSelectionOption] representing the same option as the receiver using
// [MediaSelectionOptionWithPropertyList].
//
// # Discussion
// 
// You can serialize the returned property list using
// [PropertyListSerialization].
//
// [PropertyListSerialization]: https://developer.apple.com/documentation/Foundation/PropertyListSerialization
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/propertyList()
func (m AVMediaSelectionOption) PropertyList() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("propertyList"))
	return objectivec.Object{ID: rv}
}

// The media type of the media data.
//
// # Discussion
// 
// The value of the property might be, for example, [audio] or [subtitle].
//
// [audio]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/audio
// [subtitle]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/subtitle
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/mediaType
func (m AVMediaSelectionOption) MediaType() AVMediaType {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mediaType"))
	return AVMediaType(foundation.NSStringFromID(rv).String())
}
// The media sub-types of the media data associated with the option.
//
// # Discussion
// 
// The value is an array of [NSNumber] objects carrying four character codes
// (of type FourCharCode) as defined in `CoreAudioTypes.H()` for audio media
// and in `CMFormatDescription.H()` for video media.
// 
// Also see [CMFormatDescriptionGetMediaSubType(_:)] for more information
// about media subtypes.
//
// [CMFormatDescriptionGetMediaSubType(_:)]: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetMediaSubType(_:)
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/mediaSubTypes
func (m AVMediaSelectionOption) MediaSubTypes() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("mediaSubTypes"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
// An array of metadata items for each common metadata key for which a value
// is available.
//
// # Discussion
// 
// You can filter the array of [AVMetadataItem] objects according to locale
// using [MetadataItemsFromArrayWithLocale], key using
// [MetadataItemsFromArrayWithKeyKeySpace], or language using
// [MetadataItemsFromArrayFilteredAndSortedAccordingToPreferredLanguages].
// 
// Clients that are filtering media selection options by language should be
// prepared to handle cases in which the [ExtendedLanguageTag] property value
// is `nil`. Further, they should be prepared to handle cases in which an
// `extendedLanguageTag` is present but indicates that the language is
// “undetermined” (a language value of @“und”, as defined in ISO
// 639-2).
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/commonMetadata
func (m AVMediaSelectionOption) CommonMetadata() []AVMetadataItem {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("commonMetadata"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMetadataItem {
		return AVMetadataItemFromID(id)
	})
}
// The metadata formats that contain metadata associated with the option.
//
// # Discussion
// 
// The array contains [NSString] objects, each representing a metadata format
// that contains metadata associated with the option (for example, ID3, iTunes
// metadata, and so on).
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/availableMetadataFormats
func (m AVMediaSelectionOption) AvailableMetadataFormats() []string {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("availableMetadataFormats"))
	return objc.ConvertSliceToStrings(rv)
}
// A Boolean value that indicates whether the media selection option is
// playable.
//
// # Discussion
// 
// If the media data associated with the option cannot be decoded or otherwise
// rendered, the value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/isPlayable
func (m AVMediaSelectionOption) Playable() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isPlayable"))
	return rv
}
// A string suitable for display using the current system locale.
//
// # Discussion
// 
// The string takes into account this option’s common metadata, media
// characteristics, and locale properties in addition to the provided locale
// to formulate a string intended for display
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/displayName
func (m AVMediaSelectionOption) DisplayName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("displayName"))
	return foundation.NSStringFromID(rv).String()
}
// The locale for which the media option was authored.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/locale
func (m AVMediaSelectionOption) Locale() foundation.NSLocale {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("locale"))
	return foundation.NSLocaleFromID(objc.ID(rv))
}
// The IETF BCP 47 language tag associated with the option
//
// # Discussion
// 
// This property may be `nil` indicating that the underlying media presented
// when the option is selected carries no language information. This can occur
// with media formats for which language information is optional, such as HTTP
// Live Streaming playlists, or that do not accommodate language information
// in machine-readable form.
// 
// Clients that are filtering media selection options by language should be
// prepared to handle cases in which this value is `nil`. Further, they should
// be prepared to handle cases in which an `extendedLanguageTag` is present
// but indicates that the language is “undetermined” (a language value of
// @“und”, as defined in ISO 639-2).
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionOption/extendedLanguageTag
func (m AVMediaSelectionOption) ExtendedLanguageTag() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("extendedLanguageTag"))
	return foundation.NSStringFromID(rv).String()
}

