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
	rv := objc.SendIfResponds[AVVCDuckSettings](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVVCDuckSettings.DuckLevel]
//   - [AVVCDuckSettings.SetDuckLevel]
//   - [AVVCDuckSettings.DuckOverride]
//   - [AVVCDuckSettings.SetDuckOverride]
//   - [AVVCDuckSettings.FadeDuration]
//   - [AVVCDuckSettings.SetFadeDuration]
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
func (a AVVCDuckSettings) Init() AVVCDuckSettings {
	rv := objc.SendIfResponds[AVVCDuckSettings](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVVCDuckSettings) Autorelease() AVVCDuckSettings {
	rv := objc.SendIfResponds[AVVCDuckSettings](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCDuckSettings creates a new AVVCDuckSettings instance.
func NewAVVCDuckSettings() AVVCDuckSettings {
	class := getAVVCDuckSettingsClass()
	rv := objc.SendIfResponds[AVVCDuckSettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (a AVVCDuckSettings) DuckLevel() IAVVCDuckLevel {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("duckLevel"))
	return AVVCDuckLevelFromID(objc.ID(rv))
}
func (a AVVCDuckSettings) SetDuckLevel(value IAVVCDuckLevel) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setDuckLevel:"), value)
}
func (a AVVCDuckSettings) DuckOverride() IAVVCDuckOverride {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("duckOverride"))
	return AVVCDuckOverrideFromID(objc.ID(rv))
}
func (a AVVCDuckSettings) SetDuckOverride(value IAVVCDuckOverride) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setDuckOverride:"), value)
}
func (a AVVCDuckSettings) FadeDuration() IAVVCDuckFadeDuration {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("fadeDuration"))
	return AVVCDuckFadeDurationFromID(objc.ID(rv))
}
func (a AVVCDuckSettings) SetFadeDuration(value IAVVCDuckFadeDuration) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setFadeDuration:"), value)
}
