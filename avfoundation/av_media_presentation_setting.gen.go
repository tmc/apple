// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVMediaPresentationSetting] class.
var (
	_AVMediaPresentationSettingClass     AVMediaPresentationSettingClass
	_AVMediaPresentationSettingClassOnce sync.Once
)

func getAVMediaPresentationSettingClass() AVMediaPresentationSettingClass {
	_AVMediaPresentationSettingClassOnce.Do(func() {
		_AVMediaPresentationSettingClass = AVMediaPresentationSettingClass{class: objc.GetClass("AVMediaPresentationSetting")}
	})
	return _AVMediaPresentationSettingClass
}

// GetAVMediaPresentationSettingClass returns the class object for AVMediaPresentationSetting.
func GetAVMediaPresentationSettingClass() AVMediaPresentationSettingClass {
	return getAVMediaPresentationSettingClass()
}

type AVMediaPresentationSettingClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMediaPresentationSettingClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMediaPresentationSettingClass) Alloc() AVMediaPresentationSetting {
	rv := objc.Send[AVMediaPresentationSetting](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// For content that has been authored with the express intent of offering an
// alternative selection interface for AVMediaSelectionOptions,
// AVMediaPresentationSetting represents a selectable setting for controlling
// the presentation of the media.
//
// # Overview
// 
// Each selectable setting is associated with a media characteristic that one
// or more of the AVMediaSelectionOptions in the AVMediaSelectionGroup
// possesses. By selecting a setting in a user interface that offers
// AVMediaPresentationSettings, users are essentially indicating a preference
// for the media characteristic of the selected setting. Subclasses of this
// type that are used from Swift must fulfill the requirements of a Sendable
// type.
//
// # Instance Properties
//
//   - [AVMediaPresentationSetting.MediaCharacteristic]: Provides the media characteristic that corresponds to the selectable setting.
//
// # Instance Methods
//
//   - [AVMediaPresentationSetting.DisplayNameForLocaleIdentifier]: Returns the display name for the selectable setting that best matches the specified locale identifier.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaPresentationSetting
type AVMediaPresentationSetting struct {
	objectivec.Object
}

// AVMediaPresentationSettingFromID constructs a [AVMediaPresentationSetting] from an objc.ID.
//
// For content that has been authored with the express intent of offering an
// alternative selection interface for AVMediaSelectionOptions,
// AVMediaPresentationSetting represents a selectable setting for controlling
// the presentation of the media.
func AVMediaPresentationSettingFromID(id objc.ID) AVMediaPresentationSetting {
	return AVMediaPresentationSetting{objectivec.Object{ID: id}}
}
// NOTE: AVMediaPresentationSetting adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMediaPresentationSetting] class.
//
// # Instance Properties
//
//   - [IAVMediaPresentationSetting.MediaCharacteristic]: Provides the media characteristic that corresponds to the selectable setting.
//
// # Instance Methods
//
//   - [IAVMediaPresentationSetting.DisplayNameForLocaleIdentifier]: Returns the display name for the selectable setting that best matches the specified locale identifier.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaPresentationSetting
type IAVMediaPresentationSetting interface {
	objectivec.IObject

	// Topic: Instance Properties

	// Provides the media characteristic that corresponds to the selectable setting.
	MediaCharacteristic() AVMediaCharacteristic

	// Topic: Instance Methods

	// Returns the display name for the selectable setting that best matches the specified locale identifier.
	DisplayNameForLocaleIdentifier(localeIdentifier string) string
}

// Init initializes the instance.
func (m AVMediaPresentationSetting) Init() AVMediaPresentationSetting {
	rv := objc.Send[AVMediaPresentationSetting](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMediaPresentationSetting) Autorelease() AVMediaPresentationSetting {
	rv := objc.Send[AVMediaPresentationSetting](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMediaPresentationSetting creates a new AVMediaPresentationSetting instance.
func NewAVMediaPresentationSetting() AVMediaPresentationSetting {
	class := getAVMediaPresentationSettingClass()
	rv := objc.Send[AVMediaPresentationSetting](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the display name for the selectable setting that best matches the
// specified locale identifier.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaPresentationSetting/displayName(forLocaleIdentifier:)
func (m AVMediaPresentationSetting) DisplayNameForLocaleIdentifier(localeIdentifier string) string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("displayNameForLocaleIdentifier:"), objc.String(localeIdentifier))
	return foundation.NSStringFromID(rv).String()
}

// Provides the media characteristic that corresponds to the selectable
// setting.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMediaPresentationSetting/mediaCharacteristic
func (m AVMediaPresentationSetting) MediaCharacteristic() AVMediaCharacteristic {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mediaCharacteristic"))
	return AVMediaCharacteristic(foundation.NSStringFromID(rv).String())
}

