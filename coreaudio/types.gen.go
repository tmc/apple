// Code generated from Apple documentation for CoreAudio. DO NOT EDIT.

package coreaudio

import (
	"unsafe"

	"github.com/tmc/apple/corefoundation"
)

// C struct types

// AudioDriverPlugInHostInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInHostInfo
type AudioDriverPlugInHostInfo struct {
	MDeviceID                  uint32
	MDevicePropertyChangedProc AudioDriverPlugInDevicePropertyChangedProc
	MIOAudioDevice             uintptr
	MIOAudioEngine             uintptr
	MStreamPropertyChangedProc AudioDriverPlugInStreamPropertyChangedProc
}

// AudioHardwareIOProcStreamUsage - This structure describes which streams a given AudioDeviceIOProc will use. It is used in conjunction with kAudioDevicePropertyIOProcStreamUsage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioHardwareIOProcStreamUsage
type AudioHardwareIOProcStreamUsage struct {
	MIOProc        unsafe.Pointer
	MNumberStreams uint32
	MStreamIsOn    uint32
}

// AudioObjectPropertyAddress - An AudioObjectPropertyAddress collects the three parts that identify a specific property together in a struct for easy transmission.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioObjectPropertyAddress
type AudioObjectPropertyAddress struct {
	MSelector uint32
	MScope    uint32
	MElement  uint32
}

// AudioServerPlugInClientInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInClientInfo
type AudioServerPlugInClientInfo struct {
	MBundleID       corefoundation.CFStringRef
	MClientID       uint32
	MIsNativeEndian bool
	MProcessID      int32
}

// AudioServerPlugInCustomPropertyInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInCustomPropertyInfo
type AudioServerPlugInCustomPropertyInfo struct {
	MPropertyDataType  uint32
	MQualifierDataType uint32
	MSelector          uint32
}

// AudioServerPlugInDriverInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInDriverInterface
type AudioServerPlugInDriverInterface struct {
	AbortDeviceConfigurationChange   func(uintptr, uint, uint64, unsafe.Pointer) int
	AddDeviceClient                  func(uintptr, uint, uintptr) int
	AddRef                           func(unsafe.Pointer) uint
	BeginIOOperation                 func(uintptr, uint, uint, uint, uint, uintptr) int
	CreateDevice                     func(uintptr, uintptr, uintptr, *uint) int
	DestroyDevice                    func(uintptr, uint) int
	DoIOOperation                    func(uintptr, uint, uint, uint, uint, uint, uintptr, unsafe.Pointer, unsafe.Pointer) int
	EndIOOperation                   func(uintptr, uint, uint, uint, uint, uintptr) int
	GetPropertyData                  func(uintptr, uint, int, uintptr, uint, unsafe.Pointer, uint, *uint, unsafe.Pointer) int
	GetPropertyDataSize              func(uintptr, uint, int, uintptr, uint, unsafe.Pointer, *uint) int
	GetZeroTimeStamp                 func(uintptr, uint, uint, []float64, *uint64, *uint64) int
	HasProperty                      func(uintptr, uint, int, uintptr) uint8
	Initialize                       func(uintptr, uintptr) int
	IsPropertySettable               func(uintptr, uint, int, uintptr, *byte) int
	PerformDeviceConfigurationChange func(uintptr, uint, uint64, unsafe.Pointer) int
	QueryInterface                   func(unsafe.Pointer, corefoundation.CFUUIDBytes, unsafe.Pointer) int
	Release                          func(unsafe.Pointer) uint
	RemoveDeviceClient               func(uintptr, uint, uintptr) int
	SetPropertyData                  func(uintptr, uint, int, uintptr, uint, unsafe.Pointer, uint, unsafe.Pointer) int
	StartIO                          func(uintptr, uint, uint) int
	StopIO                           func(uintptr, uint, uint) int
	WillDoIOOperation                func(uintptr, uint, uint, uint, *byte, *byte) int
}

// AudioServerPlugInHostInterface
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInHostInterface
type AudioServerPlugInHostInterface struct {
	CopyFromStorage                  func(uintptr, uintptr, unsafe.Pointer) int
	DeleteFromStorage                func(uintptr, uintptr) int
	PropertiesChanged                func(uintptr, uint, uint, uintptr) int
	RequestDeviceConfigurationChange func(uintptr, uint, uint64, unsafe.Pointer) int
	WriteToStorage                   func(uintptr, uintptr, unsafe.Pointer) int
}

// AudioServerPlugInIOCycleInfo
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInIOCycleInfo
type AudioServerPlugInIOCycleInfo struct {
	MCurrentTime              unsafe.Pointer
	MDeviceHostTicksPerFrame  float64
	MIOCycleCounter           uint64
	MInputTime                unsafe.Pointer
	MMainHostTicksPerFrame    float64
	MMasterHostTicksPerFrame  float64
	MNominalIOBufferFrameSize uint32
	MOutputTime               unsafe.Pointer
}

// AudioStreamRangedDescription - This structure allows a specific sample rate range to be associated with an AudioStreamBasicDescription that specifies its sample rate as kAudioStreamAnyRate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreAudio/AudioStreamRangedDescription
type AudioStreamRangedDescription struct {
	MFormat          unsafe.Pointer
	MSampleRateRange unsafe.Pointer
}
