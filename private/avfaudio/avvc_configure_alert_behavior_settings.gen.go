// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVCConfigureAlertBehaviorSettings] class.
var (
	_AVVCConfigureAlertBehaviorSettingsClass     AVVCConfigureAlertBehaviorSettingsClass
	_AVVCConfigureAlertBehaviorSettingsClassOnce sync.Once
)

func getAVVCConfigureAlertBehaviorSettingsClass() AVVCConfigureAlertBehaviorSettingsClass {
	_AVVCConfigureAlertBehaviorSettingsClassOnce.Do(func() {
		_AVVCConfigureAlertBehaviorSettingsClass = AVVCConfigureAlertBehaviorSettingsClass{class: objc.GetClass("AVVCConfigureAlertBehaviorSettings")}
	})
	return _AVVCConfigureAlertBehaviorSettingsClass
}

// GetAVVCConfigureAlertBehaviorSettingsClass returns the class object for AVVCConfigureAlertBehaviorSettings.
func GetAVVCConfigureAlertBehaviorSettingsClass() AVVCConfigureAlertBehaviorSettingsClass {
	return getAVVCConfigureAlertBehaviorSettingsClass()
}

type AVVCConfigureAlertBehaviorSettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVCConfigureAlertBehaviorSettingsClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVCConfigureAlertBehaviorSettingsClass) Alloc() AVVCConfigureAlertBehaviorSettings {
	rv := objc.SendIfResponds[AVVCConfigureAlertBehaviorSettings](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVVCConfigureAlertBehaviorSettings.StartAlert]
//   - [AVVCConfigureAlertBehaviorSettings.SetStartAlert]
//   - [AVVCConfigureAlertBehaviorSettings.StopAlert]
//   - [AVVCConfigureAlertBehaviorSettings.SetStopAlert]
//   - [AVVCConfigureAlertBehaviorSettings.StopOnErrorAlert]
//   - [AVVCConfigureAlertBehaviorSettings.SetStopOnErrorAlert]
//   - [AVVCConfigureAlertBehaviorSettings.StreamID]
//   - [AVVCConfigureAlertBehaviorSettings.SetStreamID]
//   - [AVVCConfigureAlertBehaviorSettings.InitWithStreamID]
type AVVCConfigureAlertBehaviorSettings struct {
	objectivec.Object
}

// AVVCConfigureAlertBehaviorSettingsFromID constructs a [AVVCConfigureAlertBehaviorSettings] from an objc.ID.
func AVVCConfigureAlertBehaviorSettingsFromID(id objc.ID) AVVCConfigureAlertBehaviorSettings {
	return AVVCConfigureAlertBehaviorSettings{objectivec.Object{ID: id}}
}

// Ensure AVVCConfigureAlertBehaviorSettings implements IAVVCConfigureAlertBehaviorSettings.
var _ IAVVCConfigureAlertBehaviorSettings = AVVCConfigureAlertBehaviorSettings{}

// An interface definition for the [AVVCConfigureAlertBehaviorSettings] class.
//
// # Methods
//
//   - [IAVVCConfigureAlertBehaviorSettings.StartAlert]
//   - [IAVVCConfigureAlertBehaviorSettings.SetStartAlert]
//   - [IAVVCConfigureAlertBehaviorSettings.StopAlert]
//   - [IAVVCConfigureAlertBehaviorSettings.SetStopAlert]
//   - [IAVVCConfigureAlertBehaviorSettings.StopOnErrorAlert]
//   - [IAVVCConfigureAlertBehaviorSettings.SetStopOnErrorAlert]
//   - [IAVVCConfigureAlertBehaviorSettings.StreamID]
//   - [IAVVCConfigureAlertBehaviorSettings.SetStreamID]
//   - [IAVVCConfigureAlertBehaviorSettings.InitWithStreamID]
type IAVVCConfigureAlertBehaviorSettings interface {
	objectivec.IObject

	// Topic: Methods

	StartAlert() int64
	SetStartAlert(value int64)
	StopAlert() int64
	SetStopAlert(value int64)
	StopOnErrorAlert() int64
	SetStopOnErrorAlert(value int64)
	StreamID() uint64
	SetStreamID(value uint64)
	InitWithStreamID(id uint64) AVVCConfigureAlertBehaviorSettings
}

// Init initializes the instance.
func (a AVVCConfigureAlertBehaviorSettings) Init() AVVCConfigureAlertBehaviorSettings {
	rv := objc.SendIfResponds[AVVCConfigureAlertBehaviorSettings](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVVCConfigureAlertBehaviorSettings) Autorelease() AVVCConfigureAlertBehaviorSettings {
	rv := objc.SendIfResponds[AVVCConfigureAlertBehaviorSettings](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCConfigureAlertBehaviorSettings creates a new AVVCConfigureAlertBehaviorSettings instance.
func NewAVVCConfigureAlertBehaviorSettings() AVVCConfigureAlertBehaviorSettings {
	class := getAVVCConfigureAlertBehaviorSettingsClass()
	rv := objc.SendIfResponds[AVVCConfigureAlertBehaviorSettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVCConfigureAlertBehaviorSettingsWithStreamID(id uint64) AVVCConfigureAlertBehaviorSettings {
	instance := getAVVCConfigureAlertBehaviorSettingsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithStreamID:"), id)
	return AVVCConfigureAlertBehaviorSettingsFromID(rv)
}

func (a AVVCConfigureAlertBehaviorSettings) InitWithStreamID(id uint64) AVVCConfigureAlertBehaviorSettings {
	rv := objc.SendIfResponds[AVVCConfigureAlertBehaviorSettings](a.ID, objc.Sel("initWithStreamID:"), id)
	return rv
}

func (a AVVCConfigureAlertBehaviorSettings) StartAlert() int64 {
	rv := objc.SendIfResponds[int64](a.ID, objc.Sel("startAlert"))
	return rv
}
func (a AVVCConfigureAlertBehaviorSettings) SetStartAlert(value int64) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setStartAlert:"), value)
}
func (a AVVCConfigureAlertBehaviorSettings) StopAlert() int64 {
	rv := objc.SendIfResponds[int64](a.ID, objc.Sel("stopAlert"))
	return rv
}
func (a AVVCConfigureAlertBehaviorSettings) SetStopAlert(value int64) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setStopAlert:"), value)
}
func (a AVVCConfigureAlertBehaviorSettings) StopOnErrorAlert() int64 {
	rv := objc.SendIfResponds[int64](a.ID, objc.Sel("stopOnErrorAlert"))
	return rv
}
func (a AVVCConfigureAlertBehaviorSettings) SetStopOnErrorAlert(value int64) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setStopOnErrorAlert:"), value)
}
func (a AVVCConfigureAlertBehaviorSettings) StreamID() uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("streamID"))
	return rv
}
func (a AVVCConfigureAlertBehaviorSettings) SetStreamID(value uint64) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setStreamID:"), value)
}
