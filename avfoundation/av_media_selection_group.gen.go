// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMediaSelectionGroup] class.
var (
	_AVMediaSelectionGroupClass     AVMediaSelectionGroupClass
	_AVMediaSelectionGroupClassOnce sync.Once
)

func getAVMediaSelectionGroupClass() AVMediaSelectionGroupClass {
	_AVMediaSelectionGroupClassOnce.Do(func() {
		_AVMediaSelectionGroupClass = AVMediaSelectionGroupClass{class: objc.GetClass("AVMediaSelectionGroup")}
	})
	return _AVMediaSelectionGroupClass
}

// GetAVMediaSelectionGroupClass returns the class object for AVMediaSelectionGroup.
func GetAVMediaSelectionGroupClass() AVMediaSelectionGroupClass {
	return getAVMediaSelectionGroupClass()
}

type AVMediaSelectionGroupClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMediaSelectionGroupClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMediaSelectionGroupClass) Alloc() AVMediaSelectionGroup {
	rv := objc.Send[AVMediaSelectionGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a collection of mutually exclusive options for
// the presentation of media within an asset.
//
// # Accessing media selection options
//
//   - [AVMediaSelectionGroup.Options]: A collection of mutually exclusive media selection options
//   - [AVMediaSelectionGroup.MediaSelectionOptionWithPropertyList]: Returns the media selection options that match the given property list.
//   - [AVMediaSelectionGroup.DefaultOption]: The default option in the group.
//
// # Configuring empty selection behavior
//
//   - [AVMediaSelectionGroup.AllowsEmptySelection]: A Boolean value that indicates whether it’s possible to present none of the options in the group when an associated player item is played.
//
// # Filtering selection options
//
//   - [AVMediaSelectionGroup.CustomMediaSelectionScheme]: For content that has been authored with the express intent of offering an alternative selection interface for AVMediaSelectionOptions, AVCustomMediaSelectionScheme provides a collection of custom settings for controlling the presentation of the media.
//
// # Creating a Now Playing language option group
//
//   - [AVMediaSelectionGroup.MakeNowPlayingInfoLanguageOptionGroup]: Creates a language option group from the media selection group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup
type AVMediaSelectionGroup struct {
	objectivec.Object
}

// AVMediaSelectionGroupFromID constructs a [AVMediaSelectionGroup] from an objc.ID.
//
// An object that represents a collection of mutually exclusive options for
// the presentation of media within an asset.
func AVMediaSelectionGroupFromID(id objc.ID) AVMediaSelectionGroup {
	return AVMediaSelectionGroup{objectivec.Object{ID: id}}
}

// NOTE: AVMediaSelectionGroup adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMediaSelectionGroup] class.
//
// # Accessing media selection options
//
//   - [IAVMediaSelectionGroup.Options]: A collection of mutually exclusive media selection options
//   - [IAVMediaSelectionGroup.MediaSelectionOptionWithPropertyList]: Returns the media selection options that match the given property list.
//   - [IAVMediaSelectionGroup.DefaultOption]: The default option in the group.
//
// # Configuring empty selection behavior
//
//   - [IAVMediaSelectionGroup.AllowsEmptySelection]: A Boolean value that indicates whether it’s possible to present none of the options in the group when an associated player item is played.
//
// # Filtering selection options
//
//   - [IAVMediaSelectionGroup.CustomMediaSelectionScheme]: For content that has been authored with the express intent of offering an alternative selection interface for AVMediaSelectionOptions, AVCustomMediaSelectionScheme provides a collection of custom settings for controlling the presentation of the media.
//
// # Creating a Now Playing language option group
//
//   - [IAVMediaSelectionGroup.MakeNowPlayingInfoLanguageOptionGroup]: Creates a language option group from the media selection group.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup
type IAVMediaSelectionGroup interface {
	objectivec.IObject

	// Topic: Accessing media selection options

	// A collection of mutually exclusive media selection options
	Options() []AVMediaSelectionOption
	// Returns the media selection options that match the given property list.
	MediaSelectionOptionWithPropertyList(plist objectivec.IObject) IAVMediaSelectionOption
	// The default option in the group.
	DefaultOption() IAVMediaSelectionOption

	// Topic: Configuring empty selection behavior

	// A Boolean value that indicates whether it’s possible to present none of the options in the group when an associated player item is played.
	AllowsEmptySelection() bool

	// Topic: Filtering selection options

	// For content that has been authored with the express intent of offering an alternative selection interface for AVMediaSelectionOptions, AVCustomMediaSelectionScheme provides a collection of custom settings for controlling the presentation of the media.
	CustomMediaSelectionScheme() IAVCustomMediaSelectionScheme

	// Topic: Creating a Now Playing language option group

	// Creates a language option group from the media selection group.
	MakeNowPlayingInfoLanguageOptionGroup() objectivec.IObject
}

// Init initializes the instance.
func (m AVMediaSelectionGroup) Init() AVMediaSelectionGroup {
	rv := objc.Send[AVMediaSelectionGroup](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMediaSelectionGroup) Autorelease() AVMediaSelectionGroup {
	rv := objc.Send[AVMediaSelectionGroup](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMediaSelectionGroup creates a new AVMediaSelectionGroup instance.
func NewAVMediaSelectionGroup() AVMediaSelectionGroup {
	class := getAVMediaSelectionGroupClass()
	rv := objc.Send[AVMediaSelectionGroup](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the media selection options that match the given property list.
//
// plist: A property list previously obtained from an option in the group using
// [PropertyList] ([AVMediaSelectionOption]).
//
// # Return Value
//
// An [AVMediaSelectionOption] object containing the properites passed by
// `plist`. Returns `nil` when no match is found.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/mediaSelectionOption(withPropertyList:)
func (m AVMediaSelectionGroup) MediaSelectionOptionWithPropertyList(plist objectivec.IObject) IAVMediaSelectionOption {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mediaSelectionOptionWithPropertyList:"), plist)
	return AVMediaSelectionOptionFromID(rv)
}

// Creates a language option group from the media selection group.
//
// # Return Value
//
// The new language option group.
//
// # Discussion
//
// Any option from [AVMediaSelectionOption] in the [AVMediaSelectionGroup] not
// representing an audible or legible selection option is ignored.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/makeNowPlayingInfoLanguageOptionGroup()
func (m AVMediaSelectionGroup) MakeNowPlayingInfoLanguageOptionGroup() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("makeNowPlayingInfoLanguageOptionGroup"))
	return objectivec.Object{ID: rv}
}

// Returns an array containing the media selection options from a given array
// that are playable.
//
// mediaSelectionOptions: An array of [AVMediaSelectionOption] objects to be filtered by playability.
//
// # Return Value
//
// An array containing the media selection options from `array` that are
// playable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/playableMediaSelectionOptions(from:)
func (_AVMediaSelectionGroupClass AVMediaSelectionGroupClass) PlayableMediaSelectionOptionsFromArray(mediaSelectionOptions []AVMediaSelectionOption) []AVMediaSelectionOption {
	rv := objc.Send[[]objc.ID](objc.ID(_AVMediaSelectionGroupClass.class), objc.Sel("playableMediaSelectionOptionsFromArray:"), objectivec.IObjectSliceToNSArray(mediaSelectionOptions))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMediaSelectionOption {
		return AVMediaSelectionOptionFromID(id)
	})
}

// Returns an array containing the media selection options from a given array
// that match the specified locale.
//
// mediaSelectionOptions: An array of [AVMediaSelectionOption] objects to be filtered.
//
// locale: The locale that must be matched for a media selection option to be copied
// to the output array.
//
// # Return Value
//
// An array containing the media selection options from `array` that match the
// `locale`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/mediaSelectionOptions(from:with:)
func (_AVMediaSelectionGroupClass AVMediaSelectionGroupClass) MediaSelectionOptionsFromArrayWithLocale(mediaSelectionOptions []AVMediaSelectionOption, locale foundation.NSLocale) []AVMediaSelectionOption {
	rv := objc.Send[[]objc.ID](objc.ID(_AVMediaSelectionGroupClass.class), objc.Sel("mediaSelectionOptionsFromArray:withLocale:"), objectivec.IObjectSliceToNSArray(mediaSelectionOptions), locale)
	return objc.ConvertSlice(rv, func(id objc.ID) AVMediaSelectionOption {
		return AVMediaSelectionOptionFromID(id)
	})
}

// Returns an array containing the media selection options from a given array
// that match given media characteristics.
//
// mediaSelectionOptions: An array of [AVMediaSelectionOption] objects to be filtered.
//
// mediaCharacteristics: The media characteristics that must be matched for a media selection option
// to be present in the output array.
//
// # Return Value
//
// An array containing the media selection options from `array` that match
// `mediaCharacteristics`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/mediaSelectionOptions(from:withMediaCharacteristics:)
func (_AVMediaSelectionGroupClass AVMediaSelectionGroupClass) MediaSelectionOptionsFromArrayWithMediaCharacteristics(mediaSelectionOptions []AVMediaSelectionOption, mediaCharacteristics []string) []AVMediaSelectionOption {
	rv := objc.Send[[]objc.ID](objc.ID(_AVMediaSelectionGroupClass.class), objc.Sel("mediaSelectionOptionsFromArray:withMediaCharacteristics:"), objectivec.IObjectSliceToNSArray(mediaSelectionOptions), objectivec.StringSliceToNSArray(mediaCharacteristics))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMediaSelectionOption {
		return AVMediaSelectionOptionFromID(id)
	})
}

// Returns an array containing the media selection options from a given array
// that do not match given media characteristics.
//
// mediaSelectionOptions: An array of [AVMediaSelectionOption] objects to be filtered.
//
// mediaCharacteristics: The media characteristics that must not be present for a media selection
// option to be present in the output array.
//
// # Return Value
//
// An array containing the media selection options from `array` that lack the
// media characteristics in `mediaCharacteristics`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/mediaSelectionOptions(from:withoutMediaCharacteristics:)
func (_AVMediaSelectionGroupClass AVMediaSelectionGroupClass) MediaSelectionOptionsFromArrayWithoutMediaCharacteristics(mediaSelectionOptions []AVMediaSelectionOption, mediaCharacteristics []string) []AVMediaSelectionOption {
	rv := objc.Send[[]objc.ID](objc.ID(_AVMediaSelectionGroupClass.class), objc.Sel("mediaSelectionOptionsFromArray:withoutMediaCharacteristics:"), objectivec.IObjectSliceToNSArray(mediaSelectionOptions), objectivec.StringSliceToNSArray(mediaCharacteristics))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMediaSelectionOption {
		return AVMediaSelectionOptionFromID(id)
	})
}

// Returns an array of media selection options, filtering them according to
// whether their locales match one of the specified languages.
//
// mediaSelectionOptions: An array of [AVMediaSelectionOption] objects to be filtered and sorted.
//
// preferredLanguages: An array of [NSString] objects, each of which contains a canonicalized IETF
// BCP 47 language identifier. The strings should be sorted in order of
// preference, with the string corresponding to the most preferred language as
// the first element in the array. Typically, you retrieve these strings using
// the [preferredLanguages] class method of the [NSLocale] class.
//
// # Return Value
//
// An array of [AVMediaSelectionOption] objects that match one of the
// languages in the `preferredLanguages` parameter. The objects in this array
// are sorted based on the language each one matches, with objects matching
// the most preferred language first in the array.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/mediaSelectionOptions(from:filteredAndSortedAccordingToPreferredLanguages:)
//
// [NSLocale]: https://developer.apple.com/documentation/Foundation/NSLocale
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [preferredLanguages]: https://developer.apple.com/documentation/Foundation/NSLocale/preferredLanguages
func (_AVMediaSelectionGroupClass AVMediaSelectionGroupClass) MediaSelectionOptionsFromArrayFilteredAndSortedAccordingToPreferredLanguages(mediaSelectionOptions []AVMediaSelectionOption, preferredLanguages []string) []AVMediaSelectionOption {
	rv := objc.Send[[]objc.ID](objc.ID(_AVMediaSelectionGroupClass.class), objc.Sel("mediaSelectionOptionsFromArray:filteredAndSortedAccordingToPreferredLanguages:"), objectivec.IObjectSliceToNSArray(mediaSelectionOptions), objectivec.StringSliceToNSArray(preferredLanguages))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMediaSelectionOption {
		return AVMediaSelectionOptionFromID(id)
	})
}

// A collection of mutually exclusive media selection options
//
// # Discussion
//
// The value of the property is an array of [AVMediaSelectionOption] objects.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/options
func (m AVMediaSelectionGroup) Options() []AVMediaSelectionOption {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("options"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMediaSelectionOption {
		return AVMediaSelectionOptionFromID(id)
	})
}

// The default option in the group.
//
// # Discussion
//
// The default option is intended for use in the absence of a specific
// end-user selection or preference. Can be `nil`, indicating that without a
// specific end-user selection or preference, no option in the group is
// intended to be selected.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/defaultOption
func (m AVMediaSelectionGroup) DefaultOption() IAVMediaSelectionOption {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("defaultOption"))
	return AVMediaSelectionOptionFromID(objc.ID(rv))
}

// A Boolean value that indicates whether it’s possible to present none of
// the options in the group when an associated player item is played.
//
// # Discussion
//
// If the value of this property is true, you can deselect all of the
// available media options in the group by passing `nil` as the specified
// [AVMediaSelectionOption] object to
// [SelectMediaOptionInMediaSelectionGroup].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/allowsEmptySelection
func (m AVMediaSelectionGroup) AllowsEmptySelection() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("allowsEmptySelection"))
	return rv
}

// For content that has been authored with the express intent of offering an
// alternative selection interface for AVMediaSelectionOptions,
// AVCustomMediaSelectionScheme provides a collection of custom settings for
// controlling the presentation of the media.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaSelectionGroup/customMediaSelectionScheme
func (m AVMediaSelectionGroup) CustomMediaSelectionScheme() IAVCustomMediaSelectionScheme {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("customMediaSelectionScheme"))
	return AVCustomMediaSelectionSchemeFromID(objc.ID(rv))
}
