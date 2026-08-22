// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVCPrepareRecordSettings] class.
var (
	_AVVCPrepareRecordSettingsClass     AVVCPrepareRecordSettingsClass
	_AVVCPrepareRecordSettingsClassOnce sync.Once
)

func getAVVCPrepareRecordSettingsClass() AVVCPrepareRecordSettingsClass {
	_AVVCPrepareRecordSettingsClassOnce.Do(func() {
		_AVVCPrepareRecordSettingsClass = AVVCPrepareRecordSettingsClass{class: objc.GetClass("AVVCPrepareRecordSettings")}
	})
	return _AVVCPrepareRecordSettingsClass
}

// GetAVVCPrepareRecordSettingsClass returns the class object for AVVCPrepareRecordSettings.
func GetAVVCPrepareRecordSettingsClass() AVVCPrepareRecordSettingsClass {
	return getAVVCPrepareRecordSettingsClass()
}

type AVVCPrepareRecordSettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVCPrepareRecordSettingsClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVCPrepareRecordSettingsClass) Alloc() AVVCPrepareRecordSettings {
	rv := objc.SendIfResponds[AVVCPrepareRecordSettings](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVVCPrepareRecordSettings.AvAudioSettings]
//   - [AVVCPrepareRecordSettings.SetAvAudioSettings]
//   - [AVVCPrepareRecordSettings.DeviceBufferFrameSize]
//   - [AVVCPrepareRecordSettings.SetDeviceBufferFrameSize]
//   - [AVVCPrepareRecordSettings.MeteringEnabled]
//   - [AVVCPrepareRecordSettings.SetMeteringEnabled]
//   - [AVVCPrepareRecordSettings.RecordBufferDuration]
//   - [AVVCPrepareRecordSettings.SetRecordBufferDuration]
//   - [AVVCPrepareRecordSettings.StreamID]
//   - [AVVCPrepareRecordSettings.SetStreamID]
//   - [AVVCPrepareRecordSettings.InitWithStreamIDSettingsBufferDuration]
type AVVCPrepareRecordSettings struct {
	objectivec.Object
}

// AVVCPrepareRecordSettingsFromID constructs a [AVVCPrepareRecordSettings] from an objc.ID.
func AVVCPrepareRecordSettingsFromID(id objc.ID) AVVCPrepareRecordSettings {
	return AVVCPrepareRecordSettings{objectivec.Object{ID: id}}
}

// Ensure AVVCPrepareRecordSettings implements IAVVCPrepareRecordSettings.
var _ IAVVCPrepareRecordSettings = AVVCPrepareRecordSettings{}

// An interface definition for the [AVVCPrepareRecordSettings] class.
//
// # Methods
//
//   - [IAVVCPrepareRecordSettings.AvAudioSettings]
//   - [IAVVCPrepareRecordSettings.SetAvAudioSettings]
//   - [IAVVCPrepareRecordSettings.DeviceBufferFrameSize]
//   - [IAVVCPrepareRecordSettings.SetDeviceBufferFrameSize]
//   - [IAVVCPrepareRecordSettings.MeteringEnabled]
//   - [IAVVCPrepareRecordSettings.SetMeteringEnabled]
//   - [IAVVCPrepareRecordSettings.RecordBufferDuration]
//   - [IAVVCPrepareRecordSettings.SetRecordBufferDuration]
//   - [IAVVCPrepareRecordSettings.StreamID]
//   - [IAVVCPrepareRecordSettings.SetStreamID]
//   - [IAVVCPrepareRecordSettings.InitWithStreamIDSettingsBufferDuration]
type IAVVCPrepareRecordSettings interface {
	objectivec.IObject

	// Topic: Methods

	AvAudioSettings() foundation.INSDictionary
	SetAvAudioSettings(value foundation.INSDictionary)
	DeviceBufferFrameSize() uint32
	SetDeviceBufferFrameSize(value uint32)
	MeteringEnabled() bool
	SetMeteringEnabled(value bool)
	RecordBufferDuration() float64
	SetRecordBufferDuration(value float64)
	StreamID() uint64
	SetStreamID(value uint64)
	InitWithStreamIDSettingsBufferDuration(id uint64, settings objectivec.IObject, duration float64) AVVCPrepareRecordSettings
}

// Init initializes the instance.
func (a AVVCPrepareRecordSettings) Init() AVVCPrepareRecordSettings {
	rv := objc.SendIfResponds[AVVCPrepareRecordSettings](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVVCPrepareRecordSettings) Autorelease() AVVCPrepareRecordSettings {
	rv := objc.SendIfResponds[AVVCPrepareRecordSettings](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCPrepareRecordSettings creates a new AVVCPrepareRecordSettings instance.
func NewAVVCPrepareRecordSettings() AVVCPrepareRecordSettings {
	class := getAVVCPrepareRecordSettingsClass()
	rv := objc.SendIfResponds[AVVCPrepareRecordSettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVCPrepareRecordSettingsWithStreamIDSettingsBufferDuration(id uint64, settings objectivec.IObject, duration float64) AVVCPrepareRecordSettings {
	instance := getAVVCPrepareRecordSettingsClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithStreamID:settings:bufferDuration:"), id, settings, duration)
	return AVVCPrepareRecordSettingsFromID(rv)
}

func (a AVVCPrepareRecordSettings) InitWithStreamIDSettingsBufferDuration(id uint64, settings objectivec.IObject, duration float64) AVVCPrepareRecordSettings {
	rv := objc.SendIfResponds[AVVCPrepareRecordSettings](a.ID, objc.Sel("initWithStreamID:settings:bufferDuration:"), id, settings, duration)
	return rv
}

func (a AVVCPrepareRecordSettings) AvAudioSettings() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("avAudioSettings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (a AVVCPrepareRecordSettings) SetAvAudioSettings(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setAvAudioSettings:"), value)
}
func (a AVVCPrepareRecordSettings) DeviceBufferFrameSize() uint32 {
	rv := objc.SendIfResponds[uint32](a.ID, objc.Sel("deviceBufferFrameSize"))
	return rv
}
func (a AVVCPrepareRecordSettings) SetDeviceBufferFrameSize(value uint32) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setDeviceBufferFrameSize:"), value)
}
func (a AVVCPrepareRecordSettings) MeteringEnabled() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("meteringEnabled"))
	return rv
}
func (a AVVCPrepareRecordSettings) SetMeteringEnabled(value bool) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setMeteringEnabled:"), value)
}
func (a AVVCPrepareRecordSettings) RecordBufferDuration() float64 {
	rv := objc.SendIfResponds[float64](a.ID, objc.Sel("recordBufferDuration"))
	return rv
}
func (a AVVCPrepareRecordSettings) SetRecordBufferDuration(value float64) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setRecordBufferDuration:"), value)
}
func (a AVVCPrepareRecordSettings) StreamID() uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("streamID"))
	return rv
}
func (a AVVCPrepareRecordSettings) SetStreamID(value uint64) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setStreamID:"), value)
}
