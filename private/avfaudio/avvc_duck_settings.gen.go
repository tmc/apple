// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVCDuckSettings] class.
var (
	_AVVCDuckSettingsClass     AVVCDuckSettingsClass
	_AVVCDuckSettingsClassOnce sync.Once
)

func getAVVCDuckSettingsClass() AVVCDuckSettingsClass {
	_AVVCDuckSettingsClassOnce.Do(func() {
		_AVVCDuckSettingsClass = AVVCDuckSettingsClass{class: objc.GetClass("AVVCDuckSettings")}
	})
	return _AVVCDuckSettingsClass
}

// GetAVVCDuckSettingsClass returns the class object for AVVCDuckSettings.
func GetAVVCDuckSettingsClass() AVVCDuckSettingsClass {
	return getAVVCDuckSettingsClass()
}

type AVVCDuckSettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVCDuckSettingsClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVCDuckSettingsClass) Alloc() AVVCDuckSettings {
	rv := objc.Send[AVVCDuckSettings](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVVCDuckSettings.DuckLevel]
//   - [AVVCDuckSettings.SetDuckLevel]
//   - [AVVCDuckSettings.DuckOverride]
//   - [AVVCDuckSettings.SetDuckOverride]
//   - [AVVCDuckSettings.FadeDuration]
//   - [AVVCDuckSettings.SetFadeDuration]
// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckSettings
type AVVCDuckSettings struct {
	objectivec.Object
}

// AVVCDuckSettingsFromID constructs a [AVVCDuckSettings] from an objc.ID.
func AVVCDuckSettingsFromID(id objc.ID) AVVCDuckSettings {
	return AVVCDuckSettings{objectivec.Object{ID: id}}
}
// Ensure AVVCDuckSettings implements IAVVCDuckSettings.
var _ IAVVCDuckSettings = AVVCDuckSettings{}

// An interface definition for the [AVVCDuckSettings] class.
//
// # Methods
//
//   - [IAVVCDuckSettings.DuckLevel]
//   - [IAVVCDuckSettings.SetDuckLevel]
//   - [IAVVCDuckSettings.DuckOverride]
//   - [IAVVCDuckSettings.SetDuckOverride]
//   - [IAVVCDuckSettings.FadeDuration]
//   - [IAVVCDuckSettings.SetFadeDuration]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckSettings
type IAVVCDuckSettings interface {
	objectivec.IObject

	// Topic: Methods

	DuckLevel() IAVVCDuckLevel
	SetDuckLevel(value IAVVCDuckLevel)
	DuckOverride() IAVVCDuckOverride
	SetDuckOverride(value IAVVCDuckOverride)
	FadeDuration() IAVVCDuckFadeDuration
	SetFadeDuration(value IAVVCDuckFadeDuration)
}

// Init initializes the instance.
func (v AVVCDuckSettings) Init() AVVCDuckSettings {
	rv := objc.Send[AVVCDuckSettings](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVCDuckSettings) Autorelease() AVVCDuckSettings {
	rv := objc.Send[AVVCDuckSettings](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCDuckSettings creates a new AVVCDuckSettings instance.
func NewAVVCDuckSettings() AVVCDuckSettings {
	class := getAVVCDuckSettingsClass()
	rv := objc.Send[AVVCDuckSettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckSettings/duckLevel
func (v AVVCDuckSettings) DuckLevel() IAVVCDuckLevel {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("duckLevel"))
	return AVVCDuckLevelFromID(objc.ID(rv))
}
func (v AVVCDuckSettings) SetDuckLevel(value IAVVCDuckLevel) {
	objc.Send[struct{}](v.ID, objc.Sel("setDuckLevel:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckSettings/duckOverride
func (v AVVCDuckSettings) DuckOverride() IAVVCDuckOverride {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("duckOverride"))
	return AVVCDuckOverrideFromID(objc.ID(rv))
}
func (v AVVCDuckSettings) SetDuckOverride(value IAVVCDuckOverride) {
	objc.Send[struct{}](v.ID, objc.Sel("setDuckOverride:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCDuckSettings/fadeDuration
func (v AVVCDuckSettings) FadeDuration() IAVVCDuckFadeDuration {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("fadeDuration"))
	return AVVCDuckFadeDurationFromID(objc.ID(rv))
}
func (v AVVCDuckSettings) SetFadeDuration(value IAVVCDuckFadeDuration) {
	objc.Send[struct{}](v.ID, objc.Sel("setFadeDuration:"), value)
}

