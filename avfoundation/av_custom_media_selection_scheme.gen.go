// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCustomMediaSelectionScheme] class.
var (
	_AVCustomMediaSelectionSchemeClass     AVCustomMediaSelectionSchemeClass
	_AVCustomMediaSelectionSchemeClassOnce sync.Once
)

func getAVCustomMediaSelectionSchemeClass() AVCustomMediaSelectionSchemeClass {
	_AVCustomMediaSelectionSchemeClassOnce.Do(func() {
		_AVCustomMediaSelectionSchemeClass = AVCustomMediaSelectionSchemeClass{class: objc.GetClass("AVCustomMediaSelectionScheme")}
	})
	return _AVCustomMediaSelectionSchemeClass
}

// GetAVCustomMediaSelectionSchemeClass returns the class object for AVCustomMediaSelectionScheme.
func GetAVCustomMediaSelectionSchemeClass() AVCustomMediaSelectionSchemeClass {
	return getAVCustomMediaSelectionSchemeClass()
}

type AVCustomMediaSelectionSchemeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCustomMediaSelectionSchemeClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCustomMediaSelectionSchemeClass) Alloc() AVCustomMediaSelectionScheme {
	rv := objc.Send[AVCustomMediaSelectionScheme](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// For content that has been authored with the express intent of offering an
// alternative selection interface for AVMediaSelectionOptions,
// AVCustomMediaSelectionScheme provides a collection of custom settings for
// controlling the presentation of the media.
//
// # Overview
//
// Each selectable setting is associated with a media characteristic that one
// or more of the AVMediaSelectionOptions in the AVMediaSelectionGroup
// possesses. By selecting a setting in a user interface based on an
// AVCustomMediaSelectionScheme, users are essentially indicating a preference
// for the media characteristic of the selected setting. Selection of a
// specific AVMediaSelectionOption in the AVMediaSelectionGroup is then
// derived from the user’s indicated preferences. Subclasses of this type
// that are used from Swift must fulfill the requirements of a Sendable type.
//
// # Inspecting the scheme
//
//   - [AVCustomMediaSelectionScheme.AvailableLanguages]: Provides available language choices.
//   - [AVCustomMediaSelectionScheme.Selectors]: Provides custom settings.
//   - [AVCustomMediaSelectionScheme.ShouldOfferLanguageSelection]: Indicates whether an alternative selection interface should provide a menu of language choices.
//   - [AVCustomMediaSelectionScheme.MediaPresentationSettingsForSelectorComplementaryToLanguageSettings]: Provides an array of media presentation settings that can be effective at the same time as the specified language and settings for other selectors of the receiver.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCustomMediaSelectionScheme
type AVCustomMediaSelectionScheme struct {
	objectivec.Object
}

// AVCustomMediaSelectionSchemeFromID constructs a [AVCustomMediaSelectionScheme] from an objc.ID.
//
// For content that has been authored with the express intent of offering an
// alternative selection interface for AVMediaSelectionOptions,
// AVCustomMediaSelectionScheme provides a collection of custom settings for
// controlling the presentation of the media.
func AVCustomMediaSelectionSchemeFromID(id objc.ID) AVCustomMediaSelectionScheme {
	return AVCustomMediaSelectionScheme{objectivec.Object{ID: id}}
}

// NOTE: AVCustomMediaSelectionScheme adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCustomMediaSelectionScheme] class.
//
// # Inspecting the scheme
//
//   - [IAVCustomMediaSelectionScheme.AvailableLanguages]: Provides available language choices.
//   - [IAVCustomMediaSelectionScheme.Selectors]: Provides custom settings.
//   - [IAVCustomMediaSelectionScheme.ShouldOfferLanguageSelection]: Indicates whether an alternative selection interface should provide a menu of language choices.
//   - [IAVCustomMediaSelectionScheme.MediaPresentationSettingsForSelectorComplementaryToLanguageSettings]: Provides an array of media presentation settings that can be effective at the same time as the specified language and settings for other selectors of the receiver.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCustomMediaSelectionScheme
type IAVCustomMediaSelectionScheme interface {
	objectivec.IObject

	// Topic: Inspecting the scheme

	// Provides available language choices.
	AvailableLanguages() []string
	// Provides custom settings.
	Selectors() []AVMediaPresentationSelector
	// Indicates whether an alternative selection interface should provide a menu of language choices.
	ShouldOfferLanguageSelection() bool
	// Provides an array of media presentation settings that can be effective at the same time as the specified language and settings for other selectors of the receiver.
	MediaPresentationSettingsForSelectorComplementaryToLanguageSettings(selector IAVMediaPresentationSelector, language string, settings []AVMediaPresentationSetting) []AVMediaPresentationSetting
}

// Init initializes the instance.
func (c AVCustomMediaSelectionScheme) Init() AVCustomMediaSelectionScheme {
	rv := objc.Send[AVCustomMediaSelectionScheme](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCustomMediaSelectionScheme) Autorelease() AVCustomMediaSelectionScheme {
	rv := objc.Send[AVCustomMediaSelectionScheme](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCustomMediaSelectionScheme creates a new AVCustomMediaSelectionScheme instance.
func NewAVCustomMediaSelectionScheme() AVCustomMediaSelectionScheme {
	class := getAVCustomMediaSelectionSchemeClass()
	rv := objc.Send[AVCustomMediaSelectionScheme](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Provides an array of media presentation settings that can be effective at
// the same time as the specified language and settings for other selectors of
// the receiver.
//
// selector: The AVMediaPresentationSelector for which complementary settings are
// requested.
//
// language: A BCP 47 language tag chosen among the availableLanguages of the receiver.
// If no language setting pertains, can be nil.
//
// settings: A collection of AVMediaPresentationSettings provided by selectors of the
// receiver other than the specified selector. Because no two
// AVMediaPresentationSettings of the same AVMediaPresentationSelector are
// complementary, an empty array will be returned if you specify more than one
// setting for any selector.
//
// # Discussion
//
// If the content is authored to provide a collection of
// AVMediaSelectionOptions that include one or more with all of the
// combinations of media characteristics of the specified
// AVMediaPresentationSettings together with all of the settings of the
// specified AVMediaPresentationSelector, this method will return all of the
// settings for that selector. However, if one or more of the available
// combinations are not possessed by any of the AVMediaSelectionOptions, it
// will return fewer.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCustomMediaSelectionScheme/mediaPresentationSettings(for:complementaryToLanguage:settings:)
func (c AVCustomMediaSelectionScheme) MediaPresentationSettingsForSelectorComplementaryToLanguageSettings(selector IAVMediaPresentationSelector, language string, settings []AVMediaPresentationSetting) []AVMediaPresentationSetting {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("mediaPresentationSettingsForSelector:complementaryToLanguage:settings:"), selector, objc.String(language), objectivec.IObjectSliceToNSArray(settings))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMediaPresentationSetting {
		return AVMediaPresentationSettingFromID(id)
	})
}

// Provides available language choices.
//
// # Discussion
//
// Each string in the array is intended to be interpreted as a BCP 47 language
// tag.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCustomMediaSelectionScheme/availableLanguages
func (c AVCustomMediaSelectionScheme) AvailableLanguages() []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("availableLanguages"))
	return objc.ConvertSliceToStrings(rv)
}

// Provides custom settings.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCustomMediaSelectionScheme/selectors
func (c AVCustomMediaSelectionScheme) Selectors() []AVMediaPresentationSelector {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("selectors"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMediaPresentationSelector {
		return AVMediaPresentationSelectorFromID(id)
	})
}

// Indicates whether an alternative selection interface should provide a menu
// of language choices.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCustomMediaSelectionScheme/shouldOfferLanguageSelection
func (c AVCustomMediaSelectionScheme) ShouldOfferLanguageSelection() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("shouldOfferLanguageSelection"))
	return rv
}
