// Code generated from Apple documentation. DO NOT EDIT.

package coreaudio

import (
	"unsafe"
	"github.com/tmc/apple/objectivec"
)

// See: https://developer.apple.com/documentation/CoreAudio/AudioClassID
type AudioClassID = uint32

// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceID
type AudioDeviceID = uint32

// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceIOBlock
type AudioDeviceIOBlock = func(objectivec.IObject, objectivec.IObject, objectivec.IObject, objectivec.IObject, objectivec.IObject)

// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceIOProc
type AudioDeviceIOProc = func(uint, uintptr, uintptr, uintptr, uintptr, uintptr, unsafe.Pointer) int

// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceIOProcID
type AudioDeviceIOProcID = string

// See: https://developer.apple.com/documentation/CoreAudio/AudioDevicePropertyID
type AudioDevicePropertyID = uint32

// See: https://developer.apple.com/documentation/CoreAudio/AudioDevicePropertyListenerProc
type AudioDevicePropertyListenerProc = func(uint, uint, uint8, uint, unsafe.Pointer) int

// See: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInDevicePropertyChangedProc
type AudioDriverPlugInDevicePropertyChangedProc = func(uint, uint, uint8, uint) int

// See: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInStreamPropertyChangedProc
type AudioDriverPlugInStreamPropertyChangedProc = func(uint, uint, uint, uint) int

// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwarePropertyID
type AudioHardwarePropertyID = uint32

// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwarePropertyListenerProc
type AudioHardwarePropertyListenerProc = func(uint, unsafe.Pointer) int

// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectID
type AudioObjectID = uint32

// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectPropertyElement
type AudioObjectPropertyElement = uint32

// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectPropertyListenerBlock
type AudioObjectPropertyListenerBlock = func(uint32, *AudioObjectPropertyAddress)

// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectPropertyListenerProc
type AudioObjectPropertyListenerProc = func(uint, uint, uintptr, unsafe.Pointer) int

// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectPropertyScope
type AudioObjectPropertyScope = uint32

// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectPropertySelector
type AudioObjectPropertySelector = uint32

// AudioServerPlugInCustomPropertyDataType is the set of data types the Host knows how to marshal between the server and the client.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInCustomPropertyDataType
type AudioServerPlugInCustomPropertyDataType = uint32

// See: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInDriverRef
type AudioServerPlugInDriverRef = *AudioServerPlugInDriverInterface

// See: https://developer.apple.com/documentation/CoreAudio/AudioServerPlugInHostRef
type AudioServerPlugInHostRef = *AudioServerPlugInHostInterface

// See: https://developer.apple.com/documentation/CoreAudio/AudioStreamID
type AudioStreamID = uint32

// See: https://developer.apple.com/documentation/CoreAudio/AudioStreamPropertyListenerProc
type AudioStreamPropertyListenerProc = func(uint, uint, uint, unsafe.Pointer) int

