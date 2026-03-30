// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMediaPresentationSelector] class.
var (
	_AVMediaPresentationSelectorClass     AVMediaPresentationSelectorClass
	_AVMediaPresentationSelectorClassOnce sync.Once
)

func getAVMediaPresentationSelectorClass() AVMediaPresentationSelectorClass {
	_AVMediaPresentationSelectorClassOnce.Do(func() {
		_AVMediaPresentationSelectorClass = AVMediaPresentationSelectorClass{class: objc.GetClass("AVMediaPresentationSelector")}
	})
	return _AVMediaPresentationSelectorClass
}

// GetAVMediaPresentationSelectorClass returns the class object for AVMediaPresentationSelector.
func GetAVMediaPresentationSelectorClass() AVMediaPresentationSelectorClass {
	return getAVMediaPresentationSelectorClass()
}

type AVMediaPresentationSelectorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMediaPresentationSelectorClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMediaPresentationSelectorClass) Alloc() AVMediaPresentationSelector {
	rv := objc.Send[AVMediaPresentationSelector](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// For content that has been authored with the express intent of offering an
// alternative selection interface for AVMediaSelectionOptions,
// AVMediaPresentationSelector represents a collection of mutually exclusive
// settings.
//
// # Overview
//
// Subclasses of this type that are used from Swift must fulfill the
// requirements of a Sendable type.
//
// # Instance Properties
//
//   - [AVMediaPresentationSelector.Identifier]: Provides the authored identifier for the selector.
//   - [AVMediaPresentationSelector.Settings]: Provides selectable mutually exclusive settings for the selector.
//
// # Instance Methods
//
//   - [AVMediaPresentationSelector.DisplayNameForLocaleIdentifier]: Returns the display name for the selector that best matches the specified locale identifier.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaPresentationSelector
type AVMediaPresentationSelector struct {
	objectivec.Object
}

// AVMediaPresentationSelectorFromID constructs a [AVMediaPresentationSelector] from an objc.ID.
//
// For content that has been authored with the express intent of offering an
// alternative selection interface for AVMediaSelectionOptions,
// AVMediaPresentationSelector represents a collection of mutually exclusive
// settings.
func AVMediaPresentationSelectorFromID(id objc.ID) AVMediaPresentationSelector {
	return AVMediaPresentationSelector{objectivec.Object{ID: id}}
}

// NOTE: AVMediaPresentationSelector adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMediaPresentationSelector] class.
//
// # Instance Properties
//
//   - [IAVMediaPresentationSelector.Identifier]: Provides the authored identifier for the selector.
//   - [IAVMediaPresentationSelector.Settings]: Provides selectable mutually exclusive settings for the selector.
//
// # Instance Methods
//
//   - [IAVMediaPresentationSelector.DisplayNameForLocaleIdentifier]: Returns the display name for the selector that best matches the specified locale identifier.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaPresentationSelector
type IAVMediaPresentationSelector interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Provides the authored identifier for the selector.
	Identifier() string
	// Provides selectable mutually exclusive settings for the selector.
	Settings() []AVMediaPresentationSetting

	// Topic: Instance Methods

	// Returns the display name for the selector that best matches the specified locale identifier.
	DisplayNameForLocaleIdentifier(localeIdentifier string) string
}

// Init initializes the instance.
func (m AVMediaPresentationSelector) Init() AVMediaPresentationSelector {
	rv := objc.Send[AVMediaPresentationSelector](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMediaPresentationSelector) Autorelease() AVMediaPresentationSelector {
	rv := objc.Send[AVMediaPresentationSelector](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMediaPresentationSelector creates a new AVMediaPresentationSelector instance.
func NewAVMediaPresentationSelector() AVMediaPresentationSelector {
	class := getAVMediaPresentationSelectorClass()
	rv := objc.Send[AVMediaPresentationSelector](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the display name for the selector that best matches the specified
// locale identifier.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaPresentationSelector/displayName(forLocaleIdentifier:)
func (m AVMediaPresentationSelector) DisplayNameForLocaleIdentifier(localeIdentifier string) string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("displayNameForLocaleIdentifier:"), objc.String(localeIdentifier))
	return foundation.NSStringFromID(rv).String()
}

// Provides the authored identifier for the selector.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaPresentationSelector/identifier
func (m AVMediaPresentationSelector) Identifier() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("identifier"))
	return foundation.NSStringFromID(rv).String()
}

// Provides selectable mutually exclusive settings for the selector.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaPresentationSelector/settings
func (m AVMediaPresentationSelector) Settings() []AVMediaPresentationSetting {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("settings"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVMediaPresentationSetting {
		return AVMediaPresentationSettingFromID(id)
	})
}
