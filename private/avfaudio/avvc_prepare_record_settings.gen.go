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
	rv := objc.Send[AVVCPrepareRecordSettings](objc.ID(ac.class), objc.Sel("alloc"))
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
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCPrepareRecordSettings
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
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCPrepareRecordSettings
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
func (v AVVCPrepareRecordSettings) Init() AVVCPrepareRecordSettings {
	rv := objc.Send[AVVCPrepareRecordSettings](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVCPrepareRecordSettings) Autorelease() AVVCPrepareRecordSettings {
	rv := objc.Send[AVVCPrepareRecordSettings](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCPrepareRecordSettings creates a new AVVCPrepareRecordSettings instance.
func NewAVVCPrepareRecordSettings() AVVCPrepareRecordSettings {
	class := getAVVCPrepareRecordSettingsClass()
	rv := objc.Send[AVVCPrepareRecordSettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCPrepareRecordSettings/initWithStreamID:settings:bufferDuration:
func NewVCPrepareRecordSettingsWithStreamIDSettingsBufferDuration(id uint64, settings objectivec.IObject, duration float64) AVVCPrepareRecordSettings {
	instance := getAVVCPrepareRecordSettingsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStreamID:settings:bufferDuration:"), id, settings, duration)
	return AVVCPrepareRecordSettingsFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCPrepareRecordSettings/initWithStreamID:settings:bufferDuration:
func (v AVVCPrepareRecordSettings) InitWithStreamIDSettingsBufferDuration(id uint64, settings objectivec.IObject, duration float64) AVVCPrepareRecordSettings {
	rv := objc.Send[AVVCPrepareRecordSettings](v.ID, objc.Sel("initWithStreamID:settings:bufferDuration:"), id, settings, duration)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCPrepareRecordSettings/avAudioSettings
func (v AVVCPrepareRecordSettings) AvAudioSettings() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("avAudioSettings"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (v AVVCPrepareRecordSettings) SetAvAudioSettings(value foundation.INSDictionary) {
	objc.Send[struct{}](v.ID, objc.Sel("setAvAudioSettings:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCPrepareRecordSettings/deviceBufferFrameSize
func (v AVVCPrepareRecordSettings) DeviceBufferFrameSize() uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("deviceBufferFrameSize"))
	return rv
}
func (v AVVCPrepareRecordSettings) SetDeviceBufferFrameSize(value uint32) {
	objc.Send[struct{}](v.ID, objc.Sel("setDeviceBufferFrameSize:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCPrepareRecordSettings/meteringEnabled
func (v AVVCPrepareRecordSettings) MeteringEnabled() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("meteringEnabled"))
	return rv
}
func (v AVVCPrepareRecordSettings) SetMeteringEnabled(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setMeteringEnabled:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCPrepareRecordSettings/recordBufferDuration
func (v AVVCPrepareRecordSettings) RecordBufferDuration() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("recordBufferDuration"))
	return rv
}
func (v AVVCPrepareRecordSettings) SetRecordBufferDuration(value float64) {
	objc.Send[struct{}](v.ID, objc.Sel("setRecordBufferDuration:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCPrepareRecordSettings/streamID
func (v AVVCPrepareRecordSettings) StreamID() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("streamID"))
	return rv
}
func (v AVVCPrepareRecordSettings) SetStreamID(value uint64) {
	objc.Send[struct{}](v.ID, objc.Sel("setStreamID:"), value)
}
