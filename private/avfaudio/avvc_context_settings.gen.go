// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[AVVCContextSettings](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVVCContextSettings.ActivationDeviceUID]
//   - [AVVCContextSettings.SetActivationDeviceUID]
//   - [AVVCContextSettings.ActivationMode]
//   - [AVVCContextSettings.SetActivationMode]
//   - [AVVCContextSettings.AnnounceCallsEnabled]
//   - [AVVCContextSettings.SetAnnounceCallsEnabled]
//   - [AVVCContextSettings.InitWithModeDeviceUID]
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
func (a AVVCContextSettings) Init() AVVCContextSettings {
	rv := objc.SendIfResponds[AVVCContextSettings](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVVCContextSettings) Autorelease() AVVCContextSettings {
	rv := objc.SendIfResponds[AVVCContextSettings](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCContextSettings creates a new AVVCContextSettings instance.
func NewAVVCContextSettings() AVVCContextSettings {
	class := getAVVCContextSettingsClass()
	rv := objc.SendIfResponds[AVVCContextSettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVCContextSettingsWithModeDeviceUID(mode int64, uid objectivec.IObject) AVVCContextSettings {
	instance := getAVVCContextSettingsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithMode:deviceUID:"), mode, uid)
	return AVVCContextSettingsFromID(rv)
}

func (a AVVCContextSettings) InitWithModeDeviceUID(mode int64, uid objectivec.IObject) AVVCContextSettings {
	rv := objc.SendIfResponds[AVVCContextSettings](a.ID, objc.Sel("initWithMode:deviceUID:"), mode, uid)
	return rv
}

func (a AVVCContextSettings) ActivationDeviceUID() string {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("activationDeviceUID"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVVCContextSettings) SetActivationDeviceUID(value string) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setActivationDeviceUID:"), objc.String(value))
}
func (a AVVCContextSettings) ActivationMode() int64 {
	rv := objc.SendIfResponds[int64](a.ID, objc.Sel("activationMode"))
	return rv
}
func (a AVVCContextSettings) SetActivationMode(value int64) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setActivationMode:"), value)
}
func (a AVVCContextSettings) AnnounceCallsEnabled() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("announceCallsEnabled"))
	return rv
}
func (a AVVCContextSettings) SetAnnounceCallsEnabled(value bool) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setAnnounceCallsEnabled:"), value)
}
