// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVCContextSettings] class.
var (
	_AVVCContextSettingsClass     AVVCContextSettingsClass
	_AVVCContextSettingsClassOnce sync.Once
)

func getAVVCContextSettingsClass() AVVCContextSettingsClass {
	_AVVCContextSettingsClassOnce.Do(func() {
		_AVVCContextSettingsClass = AVVCContextSettingsClass{class: objc.GetClass("AVVCContextSettings")}
	})
	return _AVVCContextSettingsClass
}

// GetAVVCContextSettingsClass returns the class object for AVVCContextSettings.
func GetAVVCContextSettingsClass() AVVCContextSettingsClass {
	return getAVVCContextSettingsClass()
}

type AVVCContextSettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVCContextSettingsClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVCContextSettingsClass) Alloc() AVVCContextSettings {
	rv := objc.Send[AVVCContextSettings](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVVCContextSettings.ActivationDeviceUID]
//   - [AVVCContextSettings.SetActivationDeviceUID]
//   - [AVVCContextSettings.ActivationMode]
//   - [AVVCContextSettings.SetActivationMode]
//   - [AVVCContextSettings.AnnounceCallsEnabled]
//   - [AVVCContextSettings.SetAnnounceCallsEnabled]
//   - [AVVCContextSettings.InitWithModeDeviceUID]
// See: https://developer.apple.com/documentation/AVFAudio/AVVCContextSettings
type AVVCContextSettings struct {
	objectivec.Object
}

// AVVCContextSettingsFromID constructs a [AVVCContextSettings] from an objc.ID.
func AVVCContextSettingsFromID(id objc.ID) AVVCContextSettings {
	return AVVCContextSettings{objectivec.Object{ID: id}}
}
// Ensure AVVCContextSettings implements IAVVCContextSettings.
var _ IAVVCContextSettings = AVVCContextSettings{}

// An interface definition for the [AVVCContextSettings] class.
//
// # Methods
//
//   - [IAVVCContextSettings.ActivationDeviceUID]
//   - [IAVVCContextSettings.SetActivationDeviceUID]
//   - [IAVVCContextSettings.ActivationMode]
//   - [IAVVCContextSettings.SetActivationMode]
//   - [IAVVCContextSettings.AnnounceCallsEnabled]
//   - [IAVVCContextSettings.SetAnnounceCallsEnabled]
//   - [IAVVCContextSettings.InitWithModeDeviceUID]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCContextSettings
type IAVVCContextSettings interface {
	objectivec.IObject

	// Topic: Methods

	ActivationDeviceUID() string
	SetActivationDeviceUID(value string)
	ActivationMode() int64
	SetActivationMode(value int64)
	AnnounceCallsEnabled() bool
	SetAnnounceCallsEnabled(value bool)
	InitWithModeDeviceUID(mode int64, uid objectivec.IObject) AVVCContextSettings
}

// Init initializes the instance.
func (v AVVCContextSettings) Init() AVVCContextSettings {
	rv := objc.Send[AVVCContextSettings](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVCContextSettings) Autorelease() AVVCContextSettings {
	rv := objc.Send[AVVCContextSettings](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCContextSettings creates a new AVVCContextSettings instance.
func NewAVVCContextSettings() AVVCContextSettings {
	class := getAVVCContextSettingsClass()
	rv := objc.Send[AVVCContextSettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCContextSettings/initWithMode:deviceUID:
func NewVCContextSettingsWithModeDeviceUID(mode int64, uid objectivec.IObject) AVVCContextSettings {
	instance := getAVVCContextSettingsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMode:deviceUID:"), mode, uid)
	return AVVCContextSettingsFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCContextSettings/initWithMode:deviceUID:
func (v AVVCContextSettings) InitWithModeDeviceUID(mode int64, uid objectivec.IObject) AVVCContextSettings {
	rv := objc.Send[AVVCContextSettings](v.ID, objc.Sel("initWithMode:deviceUID:"), mode, uid)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCContextSettings/activationDeviceUID
func (v AVVCContextSettings) ActivationDeviceUID() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("activationDeviceUID"))
	return foundation.NSStringFromID(rv).String()
}
func (v AVVCContextSettings) SetActivationDeviceUID(value string) {
	objc.Send[struct{}](v.ID, objc.Sel("setActivationDeviceUID:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCContextSettings/activationMode
func (v AVVCContextSettings) ActivationMode() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("activationMode"))
	return rv
}
func (v AVVCContextSettings) SetActivationMode(value int64) {
	objc.Send[struct{}](v.ID, objc.Sel("setActivationMode:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCContextSettings/announceCallsEnabled
func (v AVVCContextSettings) AnnounceCallsEnabled() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("announceCallsEnabled"))
	return rv
}
func (v AVVCContextSettings) SetAnnounceCallsEnabled(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setAnnounceCallsEnabled:"), value)
}

