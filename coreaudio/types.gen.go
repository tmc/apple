// Code generated from Apple documentation for CoreAudio. DO NOT EDIT.

package coreaudio

import (
	"unsafe"

	"github.com/tmc/apple/coreaudiotypes"
	"github.com/tmc/apple/corefoundation"
)

// C struct types

// AudioDriverPlugInHostInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInHostInfo
type AudioDriverPlugInHostInfo struct {
	MDeviceID                  AudioDeviceID
	MIOAudioDevice             uint32
	MIOAudioEngine             uint32
	MDevicePropertyChangedProc AudioDriverPlugInDevicePropertyChangedProc
	MStreamPropertyChangedProc AudioDriverPlugInStreamPropertyChangedProc
}

// AudioHardwareIOProcStreamUsage - This structure describes which streams a given AudioDeviceIOProc will use. It is used in conjunction with kAudioDevicePropertyIOProcStreamUsage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioHardwareIOProcStreamUsage
type AudioHardwareIOProcStreamUsage struct {
	MIOProc        unsafe.Pointer
	MNumberStreams uint32
	MStreamIsOn    [1]uint32
}

// AudioObjectPropertyAddress - An AudioObjectPropertyAddress collects the three parts that identify a specific property together in a struct for easy transmission.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioObjectPropertyAddress
type AudioObjectPropertyAddress struct {
	MSelector AudioObjectPropertySelector
	MScope    AudioObjectPropertyScope
	MElement  AudioObjectPropertyElement
}

// AudioServerPlugInClientInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInClientInfo
type AudioServerPlugInClientInfo struct {
	MClientID       uint32
	MProcessID      int32
	MIsNativeEndian bool
	MBundleID       corefoundation.CFStringRef
}

// AudioServerPlugInCustomPropertyInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInCustomPropertyInfo
type AudioServerPlugInCustomPropertyInfo struct {
	MSelector          AudioObjectPropertySelector
	MPropertyDataType  uint32
	MQualifierDataType uint32
}

// AudioServerPlugInDriverInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInDriverInterface
type AudioServerPlugInDriverInterface struct {
	_reserved                        unsafe.Pointer
	QueryInterface                   func(unsafe.Pointer, corefoundation.CFUUIDBytes, unsafe.Pointer) int32
	AddRef                           func(unsafe.Pointer) uint32
	Release                          func(unsafe.Pointer) uint32
	Initialize                       func(uintptr, uintptr) int32
	CreateDevice                     func(uintptr, uintptr, uintptr, *uint32) int32
	DestroyDevice                    func(uintptr, uint32) int32
	AddDeviceClient                  func(uintptr, uint32, uintptr) int32
	RemoveDeviceClient               func(uintptr, uint32, uintptr) int32
	PerformDeviceConfigurationChange func(uintptr, uint32, uint64, unsafe.Pointer) int32
	AbortDeviceConfigurationChange   func(uintptr, uint32, uint64, unsafe.Pointer) int32
	HasProperty                      func(uintptr, uint32, int32, uintptr) uint8
	IsPropertySettable               func(uintptr, uint32, int32, uintptr, *byte) int32
	GetPropertyDataSize              func(uintptr, uint32, int32, uintptr, uint32, unsafe.Pointer, *uint32) int32
	GetPropertyData                  func(uintptr, uint32, int32, uintptr, uint32, unsafe.Pointer, uint32, *uint32, unsafe.Pointer) int32
	SetPropertyData                  func(uintptr, uint32, int32, uintptr, uint32, unsafe.Pointer, uint32, unsafe.Pointer) int32
	StartIO                          func(uintptr, uint32, uint32) int32
	StopIO                           func(uintptr, uint32, uint32) int32
	GetZeroTimeStamp                 func(uintptr, uint32, uint32, []float64, *uint64, *uint64) int32
	WillDoIOOperation                func(uintptr, uint32, uint32, uint32, *byte, *byte) int32
	BeginIOOperation                 func(uintptr, uint32, uint32, uint32, uint32, uintptr) int32
	DoIOOperation                    func(uintptr, uint32, uint32, uint32, uint32, uint32, uintptr, unsafe.Pointer, unsafe.Pointer) int32
	EndIOOperation                   func(uintptr, uint32, uint32, uint32, uint32, uintptr) int32
}

// AudioServerPlugInHostInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInHostInterface
type AudioServerPlugInHostInterface struct {
	PropertiesChanged                func(uintptr, uint32, uint32, uintptr) int32
	CopyFromStorage                  func(uintptr, uintptr, unsafe.Pointer) int32
	WriteToStorage                   func(uintptr, uintptr, unsafe.Pointer) int32
	DeleteFromStorage                func(uintptr, uintptr) int32
	RequestDeviceConfigurationChange func(uintptr, uint32, uint64, unsafe.Pointer) int32
}

// AudioServerPlugInIOCycleInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInIOCycleInfo
type AudioServerPlugInIOCycleInfo struct {
	MIOCycleCounter           uint64
	MNominalIOBufferFrameSize uint32
	MCurrentTime              coreaudiotypes.AudioTimeStamp
	MInputTime                coreaudiotypes.AudioTimeStamp
	MOutputTime               coreaudiotypes.AudioTimeStamp
	MMainHostTicksPerFrame    float64
	MDeviceHostTicksPerFrame  float64
}

// AudioStreamRangedDescription - This structure allows a specific sample rate range to be associated with an AudioStreamBasicDescription that specifies its sample rate as kAudioStreamAnyRate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioStreamRangedDescription
type AudioStreamRangedDescription struct {
	MFormat          coreaudiotypes.AudioStreamBasicDescription
	MSampleRateRange coreaudiotypes.AudioValueRange
}
