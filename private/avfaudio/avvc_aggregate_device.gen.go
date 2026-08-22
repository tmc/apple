// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[AVVCAggregateDevice](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVVCAggregateDevice.BuildAggregateDevice]
//   - [AVVCAggregateDevice.CreateDictionaryForDeviceEnableTap]
//   - [AVVCAggregateDevice.DestroyAggregateDevice]
//   - [AVVCAggregateDevice.GetBuiltinSpeakerDevice]
//   - [AVVCAggregateDevice.AggregateDeviceID]
//   - [AVVCAggregateDevice.AggregateDeviceUID]
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
type IAVVCAggregateDevice interface {
	objectivec.IObject

	// Topic: Methods

	BuildAggregateDevice() int
	CreateDictionaryForDeviceEnableTap(device uint32, tap bool) corefoundation.CFDictionaryRef
	DestroyAggregateDevice() int
	GetBuiltinSpeakerDevice() uint32
	AggregateDeviceID() uint32
	AggregateDeviceUID() string
}

// Init initializes the instance.
func (a AVVCAggregateDevice) Init() AVVCAggregateDevice {
	rv := objc.SendIfResponds[AVVCAggregateDevice](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVVCAggregateDevice) Autorelease() AVVCAggregateDevice {
	rv := objc.SendIfResponds[AVVCAggregateDevice](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCAggregateDevice creates a new AVVCAggregateDevice instance.
func NewAVVCAggregateDevice() AVVCAggregateDevice {
	class := getAVVCAggregateDeviceClass()
	rv := objc.SendIfResponds[AVVCAggregateDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (a AVVCAggregateDevice) BuildAggregateDevice() int {
	rv := objc.SendIfResponds[int](a.ID, objc.Sel("BuildAggregateDevice"))
	return rv
}
func (a AVVCAggregateDevice) CreateDictionaryForDeviceEnableTap(device uint32, tap bool) corefoundation.CFDictionaryRef {
	rv := objc.SendIfResponds[corefoundation.CFDictionaryRef](a.ID, objc.Sel("CreateDictionaryForDevice:enableTap:"), device, tap)
	return corefoundation.CFDictionaryRef(rv)
}
func (a AVVCAggregateDevice) DestroyAggregateDevice() int {
	rv := objc.SendIfResponds[int](a.ID, objc.Sel("DestroyAggregateDevice"))
	return rv
}
func (a AVVCAggregateDevice) GetBuiltinSpeakerDevice() uint32 {
	rv := objc.SendIfResponds[uint32](a.ID, objc.Sel("GetBuiltinSpeakerDevice"))
	return rv
}

func (_AVVCAggregateDeviceClass AVVCAggregateDeviceClass) GetAOPDeviceID(id bool) uint32 {
	rv := objc.SendIfResponds[uint32](objc.ID(_AVVCAggregateDeviceClass.class), objc.Sel("GetAOPDeviceID:"), id)
	return rv
}
func (_AVVCAggregateDeviceClass AVVCAggregateDeviceClass) IsAOPDevicePresent() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_AVVCAggregateDeviceClass.class), objc.Sel("IsAOPDevicePresent"))
	return rv
}
func (_AVVCAggregateDeviceClass AVVCAggregateDeviceClass) SharedAggregateDevice() AVVCAggregateDevice {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_AVVCAggregateDeviceClass.class), objc.Sel("sharedAggregateDevice"))
	return AVVCAggregateDeviceFromID(rv)
}

func (a AVVCAggregateDevice) AggregateDeviceID() uint32 {
	rv := objc.SendIfResponds[uint32](a.ID, objc.Sel("aggregateDeviceID"))
	return rv
}
func (a AVVCAggregateDevice) AggregateDeviceUID() string {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("aggregateDeviceUID"))
	return foundation.NSStringFromID(rv).String()
}
