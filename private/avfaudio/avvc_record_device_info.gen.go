// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVVCRecordDeviceInfo] class.
var (
	_AVVCRecordDeviceInfoClass     AVVCRecordDeviceInfoClass
	_AVVCRecordDeviceInfoClassOnce sync.Once
)

func getAVVCRecordDeviceInfoClass() AVVCRecordDeviceInfoClass {
	_AVVCRecordDeviceInfoClassOnce.Do(func() {
		_AVVCRecordDeviceInfoClass = AVVCRecordDeviceInfoClass{class: objc.GetClass("AVVCRecordDeviceInfo")}
	})
	return _AVVCRecordDeviceInfoClass
}

// GetAVVCRecordDeviceInfoClass returns the class object for AVVCRecordDeviceInfo.
func GetAVVCRecordDeviceInfoClass() AVVCRecordDeviceInfoClass {
	return getAVVCRecordDeviceInfoClass()
}

type AVVCRecordDeviceInfoClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVVCRecordDeviceInfoClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVCRecordDeviceInfoClass) Alloc() AVVCRecordDeviceInfo {
	rv := objc.Send[AVVCRecordDeviceInfo](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [AVVCRecordDeviceInfo.HalDeviceUID]
//   - [AVVCRecordDeviceInfo.IsRemoteDevice]
//   - [AVVCRecordDeviceInfo.IsUpsamplingSourceAudio]
//   - [AVVCRecordDeviceInfo.RecordRoute]
//   - [AVVCRecordDeviceInfo.RemoteDeviceCategory]
//   - [AVVCRecordDeviceInfo.RemoteDeviceUID]
//   - [AVVCRecordDeviceInfo.RemoteDeviceUIDString]
//   - [AVVCRecordDeviceInfo.RemoteProductIdentifier]
//   - [AVVCRecordDeviceInfo.InitWithRecordingEngine]
type AVVCRecordDeviceInfo struct {
	objectivec.Object
}

// AVVCRecordDeviceInfoFromID constructs a [AVVCRecordDeviceInfo] from an objc.ID.
func AVVCRecordDeviceInfoFromID(id objc.ID) AVVCRecordDeviceInfo {
	return AVVCRecordDeviceInfo{objectivec.Object{ID: id}}
}

// Ensure AVVCRecordDeviceInfo implements IAVVCRecordDeviceInfo.
var _ IAVVCRecordDeviceInfo = AVVCRecordDeviceInfo{}

// An interface definition for the [AVVCRecordDeviceInfo] class.
//
// # Methods
//
//   - [IAVVCRecordDeviceInfo.HalDeviceUID]
//   - [IAVVCRecordDeviceInfo.IsRemoteDevice]
//   - [IAVVCRecordDeviceInfo.IsUpsamplingSourceAudio]
//   - [IAVVCRecordDeviceInfo.RecordRoute]
//   - [IAVVCRecordDeviceInfo.RemoteDeviceCategory]
//   - [IAVVCRecordDeviceInfo.RemoteDeviceUID]
//   - [IAVVCRecordDeviceInfo.RemoteDeviceUIDString]
//   - [IAVVCRecordDeviceInfo.RemoteProductIdentifier]
//   - [IAVVCRecordDeviceInfo.InitWithRecordingEngine]
type IAVVCRecordDeviceInfo interface {
	objectivec.IObject

	// Topic: Methods

	HalDeviceUID() string
	IsRemoteDevice() bool
	IsUpsamplingSourceAudio() bool
	RecordRoute() string
	RemoteDeviceCategory() uint32
	RemoteDeviceUID() foundation.NSUUID
	RemoteDeviceUIDString() string
	RemoteProductIdentifier() string
	InitWithRecordingEngine(engine unsafe.Pointer) AVVCRecordDeviceInfo
}

// Init initializes the instance.
func (a AVVCRecordDeviceInfo) Init() AVVCRecordDeviceInfo {
	rv := objc.Send[AVVCRecordDeviceInfo](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a AVVCRecordDeviceInfo) Autorelease() AVVCRecordDeviceInfo {
	rv := objc.Send[AVVCRecordDeviceInfo](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCRecordDeviceInfo creates a new AVVCRecordDeviceInfo instance.
func NewAVVCRecordDeviceInfo() AVVCRecordDeviceInfo {
	class := getAVVCRecordDeviceInfoClass()
	rv := objc.Send[AVVCRecordDeviceInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVCRecordDeviceInfoWithRecordingEngine(engine unsafe.Pointer) AVVCRecordDeviceInfo {
	instance := getAVVCRecordDeviceInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRecordingEngine:"), engine)
	return AVVCRecordDeviceInfoFromID(rv)
}

func (a AVVCRecordDeviceInfo) InitWithRecordingEngine(engine unsafe.Pointer) AVVCRecordDeviceInfo {
	rv := objc.Send[AVVCRecordDeviceInfo](a.ID, objc.Sel("initWithRecordingEngine:"), engine)
	return rv
}

func (a AVVCRecordDeviceInfo) HalDeviceUID() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("halDeviceUID"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVVCRecordDeviceInfo) IsRemoteDevice() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isRemoteDevice"))
	return rv
}
func (a AVVCRecordDeviceInfo) IsUpsamplingSourceAudio() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isUpsamplingSourceAudio"))
	return rv
}
func (a AVVCRecordDeviceInfo) RecordRoute() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("recordRoute"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVVCRecordDeviceInfo) RemoteDeviceCategory() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("remoteDeviceCategory"))
	return rv
}
func (a AVVCRecordDeviceInfo) RemoteDeviceUID() foundation.NSUUID {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("remoteDeviceUID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (a AVVCRecordDeviceInfo) RemoteDeviceUIDString() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("remoteDeviceUIDString"))
	return foundation.NSStringFromID(rv).String()
}
func (a AVVCRecordDeviceInfo) RemoteProductIdentifier() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("remoteProductIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
