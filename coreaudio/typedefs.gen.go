// Code generated from Apple documentation. DO NOT EDIT.

package coreaudio

import (
	"unsafe"

	"github.com/tmc/apple/coreaudiotypes"
)

// See: https://developer.apple.com/documentation/CoreAudio/AudioClassID
type AudioClassID = uint32

// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceID
type AudioDeviceID = uint32

// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceIOBlock
type AudioDeviceIOBlock = func(inNow *coreaudiotypes.AudioTimeStamp, inInputData *coreaudiotypes.AudioBufferList, inInputTime *coreaudiotypes.AudioTimeStamp, outOutputData *coreaudiotypes.AudioBufferList, inOutputTime *coreaudiotypes.AudioTimeStamp)

// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceIOProc
type AudioDeviceIOProc = func(inDevice uint32, inNow uintptr, inInputData uintptr, inInputTime uintptr, outOutputData uintptr, inOutputTime uintptr, inClientData unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceIOProcID
type AudioDeviceIOProcID = unsafe.Pointer

// See: https://developer.apple.com/documentation/CoreAudio/AudioDevicePropertyID
type AudioDevicePropertyID = uint32

// See: https://developer.apple.com/documentation/CoreAudio/AudioDevicePropertyListenerProc
type AudioDevicePropertyListenerProc = func(inDevice uint32, inChannel uint32, isInput uint8, inPropertyID uint32, inClientData unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInDevicePropertyChangedProc
type AudioDriverPlugInDevicePropertyChangedProc = func(inDevice uint32, inChannel uint32, isInput uint8, inPropertyID uint32) int32

// See: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInStreamPropertyChangedProc
type AudioDriverPlugInStreamPropertyChangedProc = func(inDevice uint32, inIOAudioStream uint32, inChannel uint32, inPropertyID uint32) int32

// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwarePropertyID
type AudioHardwarePropertyID = uint32

// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwarePropertyListenerProc
type AudioHardwarePropertyListenerProc = func(inPropertyID uint32, inClientData unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectID
type AudioObjectID = uint32

// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectPropertyElement
type AudioObjectPropertyElement = uint32

// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectPropertyListenerBlock
type AudioObjectPropertyListenerBlock = func(inNumberAddresses uint32, inAddresses *AudioObjectPropertyAddress)

// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectPropertyListenerProc
type AudioObjectPropertyListenerProc = func(inObjectID uint32, inNumberAddresses uint32, inAddresses uintptr, inClientData unsafe.Pointer) int32

// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectPropertyScope
type AudioObjectPropertyScope = uint32

// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectPropertySelector
type AudioObjectPropertySelector = uint32

// AudioServerPlugInCustomPropertyDataType is the set of data types the Host knows how to marshal between the server and the client.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInCustomPropertyDataType
type AudioServerPlugInCustomPropertyDataType = uint32

// See: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInDriverRef
type AudioServerPlugInDriverRef = **AudioServerPlugInDriverInterface

// See: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInHostRef
type AudioServerPlugInHostRef = *AudioServerPlugInHostInterface

// See: https://developer.apple.com/documentation/CoreAudio/AudioStreamID
type AudioStreamID = uint32

// See: https://developer.apple.com/documentation/CoreAudio/AudioStreamPropertyListenerProc
type AudioStreamPropertyListenerProc = func(inStream uint32, inChannel uint32, inPropertyID uint32, inClientData unsafe.Pointer) int32
