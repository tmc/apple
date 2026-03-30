// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVCAlertInformation] class.
var (
	_AVVCAlertInformationClass     AVVCAlertInformationClass
	_AVVCAlertInformationClassOnce sync.Once
)

func getAVVCAlertInformationClass() AVVCAlertInformationClass {
	_AVVCAlertInformationClassOnce.Do(func() {
		_AVVCAlertInformationClass = AVVCAlertInformationClass{class: objc.GetClass("AVVCAlertInformation")}
	})
	return _AVVCAlertInformationClass
}

// GetAVVCAlertInformationClass returns the class object for AVVCAlertInformation.
func GetAVVCAlertInformationClass() AVVCAlertInformationClass {
	return getAVVCAlertInformationClass()
}

type AVVCAlertInformationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVCAlertInformationClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVCAlertInformationClass) Alloc() AVVCAlertInformation {
	rv := objc.Send[AVVCAlertInformation](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVVCAlertInformation.AlertEndTime]
//   - [AVVCAlertInformation.SetAlertEndTime]
//   - [AVVCAlertInformation.Mode]
//   - [AVVCAlertInformation.SetMode]
//   - [AVVCAlertInformation.Type]
//   - [AVVCAlertInformation.SetType]
//   - [AVVCAlertInformation.InitWithAlertTypeModeEndTime]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCAlertInformation
type AVVCAlertInformation struct {
	objectivec.Object
}

// AVVCAlertInformationFromID constructs a [AVVCAlertInformation] from an objc.ID.
func AVVCAlertInformationFromID(id objc.ID) AVVCAlertInformation {
	return AVVCAlertInformation{objectivec.Object{ID: id}}
}

// Ensure AVVCAlertInformation implements IAVVCAlertInformation.
var _ IAVVCAlertInformation = AVVCAlertInformation{}

// An interface definition for the [AVVCAlertInformation] class.
//
// # Methods
//
//   - [IAVVCAlertInformation.AlertEndTime]
//   - [IAVVCAlertInformation.SetAlertEndTime]
//   - [IAVVCAlertInformation.Mode]
//   - [IAVVCAlertInformation.SetMode]
//   - [IAVVCAlertInformation.Type]
//   - [IAVVCAlertInformation.SetType]
//   - [IAVVCAlertInformation.InitWithAlertTypeModeEndTime]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCAlertInformation
type IAVVCAlertInformation interface {
	objectivec.IObject

	// Topic: Methods

	AlertEndTime() uint64
	SetAlertEndTime(value uint64)
	Mode() int64
	SetMode(value int64)
	Type() int
	SetType(value int)
	InitWithAlertTypeModeEndTime(type_ int, mode int64, time uint64) AVVCAlertInformation
}

// Init initializes the instance.
func (v AVVCAlertInformation) Init() AVVCAlertInformation {
	rv := objc.Send[AVVCAlertInformation](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVCAlertInformation) Autorelease() AVVCAlertInformation {
	rv := objc.Send[AVVCAlertInformation](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCAlertInformation creates a new AVVCAlertInformation instance.
func NewAVVCAlertInformation() AVVCAlertInformation {
	class := getAVVCAlertInformationClass()
	rv := objc.Send[AVVCAlertInformation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAlertInformation/initWithAlertType:mode:endTime:
func NewVCAlertInformationWithAlertTypeModeEndTime(type_ int, mode int64, time uint64) AVVCAlertInformation {
	instance := getAVVCAlertInformationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithAlertType:mode:endTime:"), type_, mode, time)
	return AVVCAlertInformationFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAlertInformation/initWithAlertType:mode:endTime:
func (v AVVCAlertInformation) InitWithAlertTypeModeEndTime(type_ int, mode int64, time uint64) AVVCAlertInformation {
	rv := objc.Send[AVVCAlertInformation](v.ID, objc.Sel("initWithAlertType:mode:endTime:"), type_, mode, time)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAlertInformation/alertEndTime
func (v AVVCAlertInformation) AlertEndTime() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("alertEndTime"))
	return rv
}
func (v AVVCAlertInformation) SetAlertEndTime(value uint64) {
	objc.Send[struct{}](v.ID, objc.Sel("setAlertEndTime:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAlertInformation/mode
func (v AVVCAlertInformation) Mode() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("mode"))
	return rv
}
func (v AVVCAlertInformation) SetMode(value int64) {
	objc.Send[struct{}](v.ID, objc.Sel("setMode:"), value)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAlertInformation/type
func (v AVVCAlertInformation) Type() int {
	rv := objc.Send[int](v.ID, objc.Sel("type"))
	return rv
}
func (v AVVCAlertInformation) SetType(value int) {
	objc.Send[struct{}](v.ID, objc.Sel("setType:"), value)
}
