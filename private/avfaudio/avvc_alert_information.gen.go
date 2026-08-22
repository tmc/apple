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
	rv := objc.SendIfResponds[AVVCAlertInformation](objc.ID(ac.class), objc.Sel("alloc"))
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
func (a AVVCAlertInformation) Init() AVVCAlertInformation {
	rv := objc.SendIfResponds[AVVCAlertInformation](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVVCAlertInformation) Autorelease() AVVCAlertInformation {
	rv := objc.SendIfResponds[AVVCAlertInformation](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCAlertInformation creates a new AVVCAlertInformation instance.
func NewAVVCAlertInformation() AVVCAlertInformation {
	class := getAVVCAlertInformationClass()
	rv := objc.SendIfResponds[AVVCAlertInformation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVCAlertInformationWithAlertTypeModeEndTime(type_ int, mode int64, time uint64) AVVCAlertInformation {
	instance := getAVVCAlertInformationClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithAlertType:mode:endTime:"), type_, mode, time)
	return AVVCAlertInformationFromID(rv)
}

func (a AVVCAlertInformation) InitWithAlertTypeModeEndTime(type_ int, mode int64, time uint64) AVVCAlertInformation {
	rv := objc.SendIfResponds[AVVCAlertInformation](a.ID, objc.Sel("initWithAlertType:mode:endTime:"), type_, mode, time)
	return rv
}

func (a AVVCAlertInformation) AlertEndTime() uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("alertEndTime"))
	return rv
}
func (a AVVCAlertInformation) SetAlertEndTime(value uint64) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setAlertEndTime:"), value)
}
func (a AVVCAlertInformation) Mode() int64 {
	rv := objc.SendIfResponds[int64](a.ID, objc.Sel("mode"))
	return rv
}
func (a AVVCAlertInformation) SetMode(value int64) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setMode:"), value)
}
func (a AVVCAlertInformation) Type() int {
	rv := objc.SendIfResponds[int](a.ID, objc.Sel("type"))
	return rv
}
func (a AVVCAlertInformation) SetType(value int) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setType:"), value)
}
