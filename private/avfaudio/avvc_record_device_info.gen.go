// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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

//
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
// See: https://developer.apple.com/documentation/AVFAudio/AVVCRecordDeviceInfo
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
//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCRecordDeviceInfo
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
	InitWithRecordingEngine(engine objectivec.IObject) AVVCRecordDeviceInfo
}

// Init initializes the instance.
func (v AVVCRecordDeviceInfo) Init() AVVCRecordDeviceInfo {
	rv := objc.Send[AVVCRecordDeviceInfo](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVCRecordDeviceInfo) Autorelease() AVVCRecordDeviceInfo {
	rv := objc.Send[AVVCRecordDeviceInfo](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVCRecordDeviceInfo creates a new AVVCRecordDeviceInfo instance.
func NewAVVCRecordDeviceInfo() AVVCRecordDeviceInfo {
	class := getAVVCRecordDeviceInfoClass()
	rv := objc.Send[AVVCRecordDeviceInfo](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCRecordDeviceInfo/initWithRecordingEngine:
func NewVCRecordDeviceInfoWithRecordingEngine(engine objectivec.IObject) AVVCRecordDeviceInfo {
	instance := getAVVCRecordDeviceInfoClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRecordingEngine:"), engine)
	return AVVCRecordDeviceInfoFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AVFAudio/AVVCRecordDeviceInfo/initWithRecordingEngine:
func (v AVVCRecordDeviceInfo) InitWithRecordingEngine(engine objectivec.IObject) AVVCRecordDeviceInfo {
	rv := objc.Send[AVVCRecordDeviceInfo](v.ID, objc.Sel("initWithRecordingEngine:"), engine)
	return rv
}

// See: https://developer.apple.com/documentation/AVFAudio/AVVCRecordDeviceInfo/halDeviceUID
func (v AVVCRecordDeviceInfo) HalDeviceUID() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("halDeviceUID"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCRecordDeviceInfo/isRemoteDevice
func (v AVVCRecordDeviceInfo) IsRemoteDevice() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isRemoteDevice"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCRecordDeviceInfo/isUpsamplingSourceAudio
func (v AVVCRecordDeviceInfo) IsUpsamplingSourceAudio() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isUpsamplingSourceAudio"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCRecordDeviceInfo/recordRoute
func (v AVVCRecordDeviceInfo) RecordRoute() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("recordRoute"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCRecordDeviceInfo/remoteDeviceCategory
func (v AVVCRecordDeviceInfo) RemoteDeviceCategory() uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("remoteDeviceCategory"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCRecordDeviceInfo/remoteDeviceUID
func (v AVVCRecordDeviceInfo) RemoteDeviceUID() foundation.NSUUID {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("remoteDeviceUID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCRecordDeviceInfo/remoteDeviceUIDString
func (v AVVCRecordDeviceInfo) RemoteDeviceUIDString() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("remoteDeviceUIDString"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/AVFAudio/AVVCRecordDeviceInfo/remoteProductIdentifier
func (v AVVCRecordDeviceInfo) RemoteProductIdentifier() string {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("remoteProductIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

