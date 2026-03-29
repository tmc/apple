// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVCAggregateDevice] class.
var (
	_AVVCAggregateDeviceClass     AVVCAggregateDeviceClass
	_AVVCAggregateDeviceClassOnce sync.Once
)

func getAVVCAggregateDeviceClass() AVVCAggregateDeviceClass {
	_AVVCAggregateDeviceClassOnce.Do(func() {
		_AVVCAggregateDeviceClass = AVVCAggregateDeviceClass{class: objc.GetClass("AVVCAggregateDevice")}
	})
	return _AVVCAggregateDeviceClass
}

// GetAVVCAggregateDeviceClass returns the class object for AVVCAggregateDevice.
func GetAVVCAggregateDeviceClass() AVVCAggregateDeviceClass {
	return getAVVCAggregateDeviceClass()
}

type AVVCAggregateDeviceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVCAggregateDeviceClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVCAggregateDeviceClass) Alloc() AVVCAggregateDevice {
	rv := objc.Send[AVVCAggregateDevice](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [AVVCAggregateDevice.BuildAggregateDevice]
//   - [AVVCAggregateDevice.CreateDictionaryForDeviceEnableTap]
//   - [AVVCAggregateDevice.DestroyAggregateDevice]
//   - [AVVCAggregateDevice.GetBuiltinSpeakerDevice]
//   - [AVVCAggregateDevice.AggregateDeviceID]
//   - [AVVCAggregateDevice.AggregateDeviceUID]
// See: https://developer.apple.com/documentation/AVFAudio/AVVCAggregateDevice
type AVVCAggregateDevice struct {
	objectivec.Object
}

// AVVCAggregateDeviceFromID constructs a [AVVCAggregateDevice] from an objc.ID.
func AVVCAggregateDeviceFromID(id objc.ID) AVVCAggregateDevice {
	return AVVCAggregateDevice{objectivec.Object{ID: id}}
}
// Ensure AVVCAggregateDevice implements IAVVCAggregateDevice.
var _ IAVVCAggregateDevice = AVVCAggregateDevice{}

// An interface definition for the [AVVCAggregateDevice] class.
//
// # Methods
//
//   - [IAVVCAggregateDevice.BuildAggregateDevice]
//   - [IAVVCAggregateDevice.CreateDictionaryForDeviceEnableTap]
//   - [IAVVCAggregateDevice.DestroyAggregateDevice]
//   - [IAVVCAggregateDevice.GetBuiltinSpeakerDevice]
//   - [IAVVCAggregateDevice.AggregateDeviceID]
//   - [IAVVCAggregateDevice.AggregateDeviceUID]
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCAggregateDevice
type IAVVCAggregateDevice interface {
	objectivec.IObject

	// Topic: Methods

	BuildAggregateDevice() int
	CreateDictionaryForDeviceEnableTap(device uint32, tap bool) objectivec.IObject
	DestroyAggregateDevice() int
	GetBuiltinSpeakerDevice() uint32
	AggregateDeviceID() uint32
	AggregateDeviceUID() string
}

// Init initializes the instance.
func (v AVVCAggregateDevice) Init() AVVCAggregateDevice {
	rv := objc.Send[AVVCAggregateDevice](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVCAggregateDevice) Autorelease() AVVCAggregateDevice {
	rv := objc.Send[AVVCAggregateDevice](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCAggregateDevice creates a new AVVCAggregateDevice instance.
func NewAVVCAggregateDevice() AVVCAggregateDevice {
	class := getAVVCAggregateDeviceClass()
	rv := objc.Send[AVVCAggregateDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAggregateDevice/BuildAggregateDevice
func (v AVVCAggregateDevice) BuildAggregateDevice() int {
	rv := objc.Send[int](v.ID, objc.Sel("BuildAggregateDevice"))
	return rv
}
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCAggregateDevice/CreateDictionaryForDevice:enableTap:
func (v AVVCAggregateDevice) CreateDictionaryForDeviceEnableTap(device uint32, tap bool) objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("CreateDictionaryForDevice:enableTap:"), device, tap)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCAggregateDevice/DestroyAggregateDevice
func (v AVVCAggregateDevice) DestroyAggregateDevice() int {
	rv := objc.Send[int](v.ID, objc.Sel("DestroyAggregateDevice"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCAggregateDevice/GetBuiltinSpeakerDevice
func (v AVVCAggregateDevice) GetBuiltinSpeakerDevice() uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("GetBuiltinSpeakerDevice"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCAggregateDevice/GetAOPDeviceID:
func (_AVVCAggregateDeviceClass AVVCAggregateDeviceClass) GetAOPDeviceID(id bool) uint32 {
	rv := objc.Send[uint32](objc.ID(_AVVCAggregateDeviceClass.class), objc.Sel("GetAOPDeviceID:"), id)
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCAggregateDevice/IsAOPDevicePresent
func (_AVVCAggregateDeviceClass AVVCAggregateDeviceClass) IsAOPDevicePresent() bool {
	rv := objc.Send[bool](objc.ID(_AVVCAggregateDeviceClass.class), objc.Sel("IsAOPDevicePresent"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCAggregateDevice/sharedAggregateDevice
func (_AVVCAggregateDeviceClass AVVCAggregateDeviceClass) SharedAggregateDevice() AVVCAggregateDevice {
	rv := objc.Send[objc.ID](objc.ID(_AVVCAggregateDeviceClass.class), objc.Sel("sharedAggregateDevice"))
	return AVVCAggregateDeviceFromID(rv)
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCAggregateDevice/aggregateDeviceID
func (v AVVCAggregateDevice) AggregateDeviceID() uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("aggregateDeviceID"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCAggregateDevice/aggregateDeviceUID
func (v AVVCAggregateDevice) AggregateDeviceUID() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("aggregateDeviceUID"))
	return foundation.NSStringFromID(rv).String()
}

