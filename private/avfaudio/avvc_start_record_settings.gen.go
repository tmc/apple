// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVCStartRecordSettings] class.
var (
	_AVVCStartRecordSettingsClass     AVVCStartRecordSettingsClass
	_AVVCStartRecordSettingsClassOnce sync.Once
)

func getAVVCStartRecordSettingsClass() AVVCStartRecordSettingsClass {
	_AVVCStartRecordSettingsClassOnce.Do(func() {
		_AVVCStartRecordSettingsClass = AVVCStartRecordSettingsClass{class: objc.GetClass("AVVCStartRecordSettings")}
	})
	return _AVVCStartRecordSettingsClass
}

// GetAVVCStartRecordSettingsClass returns the class object for AVVCStartRecordSettings.
func GetAVVCStartRecordSettingsClass() AVVCStartRecordSettingsClass {
	return getAVVCStartRecordSettingsClass()
}

type AVVCStartRecordSettingsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVCStartRecordSettingsClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVCStartRecordSettingsClass) Alloc() AVVCStartRecordSettings {
	rv := objc.Send[AVVCStartRecordSettings](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVVCStartRecordSettings.SkipAlert]
//   - [AVVCStartRecordSettings.SetSkipAlert]
//   - [AVVCStartRecordSettings.StartAlert]
//   - [AVVCStartRecordSettings.SetStartAlert]
//   - [AVVCStartRecordSettings.StartAnchorPoint]
//   - [AVVCStartRecordSettings.SetStartAnchorPoint]
//   - [AVVCStartRecordSettings.StartHostTime]
//   - [AVVCStartRecordSettings.SetStartHostTime]
//   - [AVVCStartRecordSettings.StopAlert]
//   - [AVVCStartRecordSettings.SetStopAlert]
//   - [AVVCStartRecordSettings.StopOnErrorAlert]
//   - [AVVCStartRecordSettings.SetStopOnErrorAlert]
//   - [AVVCStartRecordSettings.StreamID]
//   - [AVVCStartRecordSettings.SetStreamID]
//   - [AVVCStartRecordSettings.InitWithStreamIDAtStartHostTime]
// See: https://developer.apple.com/documentation/AVFAudio/AVVCStartRecordSettings
type AVVCStartRecordSettings struct {
	objectivec.Object
}

// AVVCStartRecordSettingsFromID constructs a [AVVCStartRecordSettings] from an objc.ID.
func AVVCStartRecordSettingsFromID(id objc.ID) AVVCStartRecordSettings {
	return AVVCStartRecordSettings{objectivec.Object{ID: id}}
}
// Ensure AVVCStartRecordSettings implements IAVVCStartRecordSettings.
var _ IAVVCStartRecordSettings = AVVCStartRecordSettings{}

// An interface definition for the [AVVCStartRecordSettings] class.
//
// # Methods
//
//   - [IAVVCStartRecordSettings.SkipAlert]
//   - [IAVVCStartRecordSettings.SetSkipAlert]
//   - [IAVVCStartRecordSettings.StartAlert]
//   - [IAVVCStartRecordSettings.SetStartAlert]
//   - [IAVVCStartRecordSettings.StartAnchorPoint]
//   - [IAVVCStartRecordSettings.SetStartAnchorPoint]
//   - [IAVVCStartRecordSettings.StartHostTime]
//   - [IAVVCStartRecordSettings.SetStartHostTime]
//   - [IAVVCStartRecordSettings.StopAlert]
//   - [IAVVCStartRecordSettings.SetStopAlert]
//   - [IAVVCStartRecordSettings.StopOnErrorAlert]
//   - [IAVVCStartRecordSettings.SetStopOnErrorAlert]
//   - [IAVVCStartRecordSettings.StreamID]
//   - [IAVVCStartRecordSettings.SetStreamID]
//   - [IAVVCStartRecordSettings.InitWithStreamIDAtStartHostTime]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCStartRecordSettings
type IAVVCStartRecordSettings interface {
	objectivec.IObject

	// Topic: Methods

	SkipAlert() bool
	SetSkipAlert(value bool)
	StartAlert() int64
	SetStartAlert(value int64)
	StartAnchorPoint() uint32
	SetStartAnchorPoint(value uint32)
	StartHostTime() uint64
	SetStartHostTime(value uint64)
	StopAlert() int64
	SetStopAlert(value int64)
	StopOnErrorAlert() int64
	SetStopOnErrorAlert(value int64)
	StreamID() uint64
	SetStreamID(value uint64)
	InitWithStreamIDAtStartHostTime(id uint64, time uint64) AVVCStartRecordSettings
}

// Init initializes the instance.
func (v AVVCStartRecordSettings) Init() AVVCStartRecordSettings {
	rv := objc.Send[AVVCStartRecordSettings](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVCStartRecordSettings) Autorelease() AVVCStartRecordSettings {
	rv := objc.Send[AVVCStartRecordSettings](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCStartRecordSettings creates a new AVVCStartRecordSettings instance.
func NewAVVCStartRecordSettings() AVVCStartRecordSettings {
	class := getAVVCStartRecordSettingsClass()
	rv := objc.Send[AVVCStartRecordSettings](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCStartRecordSettings/initWithStreamID:atStartHostTime:
func NewVCStartRecordSettingsWithStreamIDAtStartHostTime(id uint64, time uint64) AVVCStartRecordSettings {
	instance := getAVVCStartRecordSettingsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStreamID:atStartHostTime:"), id, time)
	return AVVCStartRecordSettingsFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCStartRecordSettings/initWithStreamID:atStartHostTime:
func (v AVVCStartRecordSettings) InitWithStreamIDAtStartHostTime(id uint64, time uint64) AVVCStartRecordSettings {
	rv := objc.Send[AVVCStartRecordSettings](v.ID, objc.Sel("initWithStreamID:atStartHostTime:"), id, time)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCStartRecordSettings/skipAlert
func (v AVVCStartRecordSettings) SkipAlert() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("skipAlert"))
	return rv
}
func (v AVVCStartRecordSettings) SetSkipAlert(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setSkipAlert:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCStartRecordSettings/startAlert
func (v AVVCStartRecordSettings) StartAlert() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("startAlert"))
	return rv
}
func (v AVVCStartRecordSettings) SetStartAlert(value int64) {
	objc.Send[struct{}](v.ID, objc.Sel("setStartAlert:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCStartRecordSettings/startAnchorPoint
func (v AVVCStartRecordSettings) StartAnchorPoint() uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("startAnchorPoint"))
	return rv
}
func (v AVVCStartRecordSettings) SetStartAnchorPoint(value uint32) {
	objc.Send[struct{}](v.ID, objc.Sel("setStartAnchorPoint:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCStartRecordSettings/startHostTime
func (v AVVCStartRecordSettings) StartHostTime() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("startHostTime"))
	return rv
}
func (v AVVCStartRecordSettings) SetStartHostTime(value uint64) {
	objc.Send[struct{}](v.ID, objc.Sel("setStartHostTime:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCStartRecordSettings/stopAlert
func (v AVVCStartRecordSettings) StopAlert() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("stopAlert"))
	return rv
}
func (v AVVCStartRecordSettings) SetStopAlert(value int64) {
	objc.Send[struct{}](v.ID, objc.Sel("setStopAlert:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCStartRecordSettings/stopOnErrorAlert
func (v AVVCStartRecordSettings) StopOnErrorAlert() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("stopOnErrorAlert"))
	return rv
}
func (v AVVCStartRecordSettings) SetStopOnErrorAlert(value int64) {
	objc.Send[struct{}](v.ID, objc.Sel("setStopOnErrorAlert:"), value)
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCStartRecordSettings/streamID
func (v AVVCStartRecordSettings) StreamID() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("streamID"))
	return rv
}
func (v AVVCStartRecordSettings) SetStreamID(value uint64) {
	objc.Send[struct{}](v.ID, objc.Sel("setStreamID:"), value)
}

