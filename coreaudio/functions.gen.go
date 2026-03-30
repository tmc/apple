// Code generated from Apple documentation for CoreAudio. DO NOT EDIT.

package coreaudio

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: CoreAudio: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: CoreAudio: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: CoreAudio: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _audioConvertHostTimeToNanos func(inHostTime uint64) uint64

// AudioConvertHostTimeToNanos.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioConvertHostTimeToNanos(_:)
func AudioConvertHostTimeToNanos(inHostTime uint64) uint64 {
	if _audioConvertHostTimeToNanos == nil {
		panic("CoreAudio: symbol AudioConvertHostTimeToNanos not loaded")
	}
	return _audioConvertHostTimeToNanos(inHostTime)
}

var _audioConvertNanosToHostTime func(inNanos uint64) uint64

// AudioConvertNanosToHostTime.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioConvertNanosToHostTime(_:)
func AudioConvertNanosToHostTime(inNanos uint64) uint64 {
	if _audioConvertNanosToHostTime == nil {
		panic("CoreAudio: symbol AudioConvertNanosToHostTime not loaded")
	}
	return _audioConvertNanosToHostTime(inNanos)
}

var _audioDeviceAddIOProc func(inDevice uint32, inProc AudioDeviceIOProc, inClientData unsafe.Pointer) int32

// AudioDeviceAddIOProc.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceAddIOProc
func AudioDeviceAddIOProc(inDevice uint32, inProc AudioDeviceIOProc, inClientData unsafe.Pointer) int32 {
	if _audioDeviceAddIOProc == nil {
		panic("CoreAudio: symbol AudioDeviceAddIOProc not loaded")
	}
	return _audioDeviceAddIOProc(inDevice, inProc, inClientData)
}

var _audioDeviceAddPropertyListener func(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, inProc AudioDevicePropertyListenerProc, inClientData unsafe.Pointer) int32

// AudioDeviceAddPropertyListener.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceAddPropertyListener
func AudioDeviceAddPropertyListener(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, inProc AudioDevicePropertyListenerProc, inClientData unsafe.Pointer) int32 {
	if _audioDeviceAddPropertyListener == nil {
		panic("CoreAudio: symbol AudioDeviceAddPropertyListener not loaded")
	}
	return _audioDeviceAddPropertyListener(inDevice, inChannel, isInput, inPropertyID, inProc, inClientData)
}

var _audioDeviceCreateIOProcID func(inDevice uint32, inProc AudioDeviceIOProc, inClientData unsafe.Pointer, outIOProcID *AudioDeviceIOProcID) int32

// AudioDeviceCreateIOProcID.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceCreateIOProcID(_:_:_:_:)
func AudioDeviceCreateIOProcID(inDevice uint32, inProc AudioDeviceIOProc, inClientData unsafe.Pointer, outIOProcID *AudioDeviceIOProcID) int32 {
	if _audioDeviceCreateIOProcID == nil {
		panic("CoreAudio: symbol AudioDeviceCreateIOProcID not loaded")
	}
	return _audioDeviceCreateIOProcID(inDevice, inProc, inClientData, outIOProcID)
}

var _audioDeviceCreateIOProcIDWithBlock func(outIOProcID *AudioDeviceIOProcID, inDevice uint32, inDispatchQueue uintptr, inIOBlock AudioDeviceIOBlock) int32

// AudioDeviceCreateIOProcIDWithBlock.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceCreateIOProcIDWithBlock(_:_:_:_:)
func AudioDeviceCreateIOProcIDWithBlock(outIOProcID *AudioDeviceIOProcID, inDevice uint32, inDispatchQueue dispatch.Queue, inIOBlock AudioDeviceIOBlock) int32 {
	if _audioDeviceCreateIOProcIDWithBlock == nil {
		panic("CoreAudio: symbol AudioDeviceCreateIOProcIDWithBlock not loaded")
	}
	return _audioDeviceCreateIOProcIDWithBlock(outIOProcID, inDevice, uintptr(inDispatchQueue.Handle()), inIOBlock)
}

var _audioDeviceDestroyIOProcID func(inDevice uint32, inIOProcID AudioDeviceIOProcID) int32

// AudioDeviceDestroyIOProcID.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceDestroyIOProcID(_:_:)
func AudioDeviceDestroyIOProcID(inDevice uint32, inIOProcID AudioDeviceIOProcID) int32 {
	if _audioDeviceDestroyIOProcID == nil {
		panic("CoreAudio: symbol AudioDeviceDestroyIOProcID not loaded")
	}
	return _audioDeviceDestroyIOProcID(inDevice, inIOProcID)
}

var _audioDeviceGetCurrentTime func(inDevice uint32, outTime uintptr) int32

// AudioDeviceGetCurrentTime.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceGetCurrentTime(_:_:)
func AudioDeviceGetCurrentTime(inDevice uint32, outTime uintptr) int32 {
	if _audioDeviceGetCurrentTime == nil {
		panic("CoreAudio: symbol AudioDeviceGetCurrentTime not loaded")
	}
	return _audioDeviceGetCurrentTime(inDevice, outTime)
}

var _audioDeviceGetNearestStartTime func(inDevice uint32, ioRequestedStartTime uintptr, inFlags uint32) int32

// AudioDeviceGetNearestStartTime.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceGetNearestStartTime(_:_:_:)
func AudioDeviceGetNearestStartTime(inDevice uint32, ioRequestedStartTime uintptr, inFlags uint32) int32 {
	if _audioDeviceGetNearestStartTime == nil {
		panic("CoreAudio: symbol AudioDeviceGetNearestStartTime not loaded")
	}
	return _audioDeviceGetNearestStartTime(inDevice, ioRequestedStartTime, inFlags)
}

var _audioDeviceGetProperty func(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32

// AudioDeviceGetProperty.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceGetProperty
func AudioDeviceGetProperty(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	if _audioDeviceGetProperty == nil {
		panic("CoreAudio: symbol AudioDeviceGetProperty not loaded")
	}
	return _audioDeviceGetProperty(inDevice, inChannel, isInput, inPropertyID, ioPropertyDataSize, outPropertyData)
}

var _audioDeviceGetPropertyInfo func(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, outSize *uint32, outWritable *bool) int32

// AudioDeviceGetPropertyInfo.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceGetPropertyInfo
func AudioDeviceGetPropertyInfo(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, outSize *uint32, outWritable *bool) int32 {
	if _audioDeviceGetPropertyInfo == nil {
		panic("CoreAudio: symbol AudioDeviceGetPropertyInfo not loaded")
	}
	return _audioDeviceGetPropertyInfo(inDevice, inChannel, isInput, inPropertyID, outSize, outWritable)
}

var _audioDeviceRead func(inDevice uint32, inStartTime uintptr, outData uintptr) int32

// AudioDeviceRead.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceRead
func AudioDeviceRead(inDevice uint32, inStartTime uintptr, outData uintptr) int32 {
	if _audioDeviceRead == nil {
		panic("CoreAudio: symbol AudioDeviceRead not loaded")
	}
	return _audioDeviceRead(inDevice, inStartTime, outData)
}

var _audioDeviceRemoveIOProc func(inDevice uint32, inProc AudioDeviceIOProc) int32

// AudioDeviceRemoveIOProc.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceRemoveIOProc
func AudioDeviceRemoveIOProc(inDevice uint32, inProc AudioDeviceIOProc) int32 {
	if _audioDeviceRemoveIOProc == nil {
		panic("CoreAudio: symbol AudioDeviceRemoveIOProc not loaded")
	}
	return _audioDeviceRemoveIOProc(inDevice, inProc)
}

var _audioDeviceRemovePropertyListener func(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, inProc AudioDevicePropertyListenerProc) int32

// AudioDeviceRemovePropertyListener.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceRemovePropertyListener
func AudioDeviceRemovePropertyListener(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, inProc AudioDevicePropertyListenerProc) int32 {
	if _audioDeviceRemovePropertyListener == nil {
		panic("CoreAudio: symbol AudioDeviceRemovePropertyListener not loaded")
	}
	return _audioDeviceRemovePropertyListener(inDevice, inChannel, isInput, inPropertyID, inProc)
}

var _audioDeviceSetProperty func(inDevice uint32, inWhen uintptr, inChannel uint32, isInput bool, inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32

// AudioDeviceSetProperty.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceSetProperty
func AudioDeviceSetProperty(inDevice uint32, inWhen uintptr, inChannel uint32, isInput bool, inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32 {
	if _audioDeviceSetProperty == nil {
		panic("CoreAudio: symbol AudioDeviceSetProperty not loaded")
	}
	return _audioDeviceSetProperty(inDevice, inWhen, inChannel, isInput, inPropertyID, inPropertyDataSize, inPropertyData)
}

var _audioDeviceStart func(inDevice uint32, inProcID AudioDeviceIOProcID) int32

// AudioDeviceStart.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceStart(_:_:)
func AudioDeviceStart(inDevice uint32, inProcID AudioDeviceIOProcID) int32 {
	if _audioDeviceStart == nil {
		panic("CoreAudio: symbol AudioDeviceStart not loaded")
	}
	return _audioDeviceStart(inDevice, inProcID)
}

var _audioDeviceStartAtTime func(inDevice uint32, inProcID AudioDeviceIOProcID, ioRequestedStartTime uintptr, inFlags uint32) int32

// AudioDeviceStartAtTime.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceStartAtTime(_:_:_:_:)
func AudioDeviceStartAtTime(inDevice uint32, inProcID AudioDeviceIOProcID, ioRequestedStartTime uintptr, inFlags uint32) int32 {
	if _audioDeviceStartAtTime == nil {
		panic("CoreAudio: symbol AudioDeviceStartAtTime not loaded")
	}
	return _audioDeviceStartAtTime(inDevice, inProcID, ioRequestedStartTime, inFlags)
}

var _audioDeviceStop func(inDevice uint32, inProcID AudioDeviceIOProcID) int32

// AudioDeviceStop.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceStop(_:_:)
func AudioDeviceStop(inDevice uint32, inProcID AudioDeviceIOProcID) int32 {
	if _audioDeviceStop == nil {
		panic("CoreAudio: symbol AudioDeviceStop not loaded")
	}
	return _audioDeviceStop(inDevice, inProcID)
}

var _audioDeviceTranslateTime func(inDevice uint32, inTime uintptr, outTime uintptr) int32

// AudioDeviceTranslateTime.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceTranslateTime(_:_:_:)
func AudioDeviceTranslateTime(inDevice uint32, inTime uintptr, outTime uintptr) int32 {
	if _audioDeviceTranslateTime == nil {
		panic("CoreAudio: symbol AudioDeviceTranslateTime not loaded")
	}
	return _audioDeviceTranslateTime(inDevice, inTime, outTime)
}

var _audioDriverPlugInClose func(inDevice uint32) int32

// AudioDriverPlugInClose.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInClose
func AudioDriverPlugInClose(inDevice uint32) int32 {
	if _audioDriverPlugInClose == nil {
		panic("CoreAudio: symbol AudioDriverPlugInClose not loaded")
	}
	return _audioDriverPlugInClose(inDevice)
}

var _audioDriverPlugInDeviceGetProperty func(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32

// AudioDriverPlugInDeviceGetProperty.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInDeviceGetProperty
func AudioDriverPlugInDeviceGetProperty(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	if _audioDriverPlugInDeviceGetProperty == nil {
		panic("CoreAudio: symbol AudioDriverPlugInDeviceGetProperty not loaded")
	}
	return _audioDriverPlugInDeviceGetProperty(inDevice, inChannel, isInput, inPropertyID, ioPropertyDataSize, outPropertyData)
}

var _audioDriverPlugInDeviceGetPropertyInfo func(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, outSize *uint32, outWritable *bool) int32

// AudioDriverPlugInDeviceGetPropertyInfo.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInDeviceGetPropertyInfo
func AudioDriverPlugInDeviceGetPropertyInfo(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, outSize *uint32, outWritable *bool) int32 {
	if _audioDriverPlugInDeviceGetPropertyInfo == nil {
		panic("CoreAudio: symbol AudioDriverPlugInDeviceGetPropertyInfo not loaded")
	}
	return _audioDriverPlugInDeviceGetPropertyInfo(inDevice, inChannel, isInput, inPropertyID, outSize, outWritable)
}

var _audioDriverPlugInDeviceSetProperty func(inDevice uint32, inWhen uintptr, inChannel uint32, isInput bool, inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32

// AudioDriverPlugInDeviceSetProperty.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInDeviceSetProperty
func AudioDriverPlugInDeviceSetProperty(inDevice uint32, inWhen uintptr, inChannel uint32, isInput bool, inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32 {
	if _audioDriverPlugInDeviceSetProperty == nil {
		panic("CoreAudio: symbol AudioDriverPlugInDeviceSetProperty not loaded")
	}
	return _audioDriverPlugInDeviceSetProperty(inDevice, inWhen, inChannel, isInput, inPropertyID, inPropertyDataSize, inPropertyData)
}

var _audioDriverPlugInOpen func(inHostInfo *AudioDriverPlugInHostInfo) int32

// AudioDriverPlugInOpen.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInOpen
func AudioDriverPlugInOpen(inHostInfo *AudioDriverPlugInHostInfo) int32 {
	if _audioDriverPlugInOpen == nil {
		panic("CoreAudio: symbol AudioDriverPlugInOpen not loaded")
	}
	return _audioDriverPlugInOpen(inHostInfo)
}

var _audioDriverPlugInStreamGetProperty func(inDevice uint32, inIOAudioStream uintptr, inChannel uint32, inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32

// AudioDriverPlugInStreamGetProperty.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInStreamGetProperty
func AudioDriverPlugInStreamGetProperty(inDevice uint32, inIOAudioStream uintptr, inChannel uint32, inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	if _audioDriverPlugInStreamGetProperty == nil {
		panic("CoreAudio: symbol AudioDriverPlugInStreamGetProperty not loaded")
	}
	return _audioDriverPlugInStreamGetProperty(inDevice, inIOAudioStream, inChannel, inPropertyID, ioPropertyDataSize, outPropertyData)
}

var _audioDriverPlugInStreamGetPropertyInfo func(inDevice uint32, inIOAudioStream uintptr, inChannel uint32, inPropertyID uint32, outSize *uint32, outWritable *bool) int32

// AudioDriverPlugInStreamGetPropertyInfo.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInStreamGetPropertyInfo
func AudioDriverPlugInStreamGetPropertyInfo(inDevice uint32, inIOAudioStream uintptr, inChannel uint32, inPropertyID uint32, outSize *uint32, outWritable *bool) int32 {
	if _audioDriverPlugInStreamGetPropertyInfo == nil {
		panic("CoreAudio: symbol AudioDriverPlugInStreamGetPropertyInfo not loaded")
	}
	return _audioDriverPlugInStreamGetPropertyInfo(inDevice, inIOAudioStream, inChannel, inPropertyID, outSize, outWritable)
}

var _audioDriverPlugInStreamSetProperty func(inDevice uint32, inIOAudioStream uintptr, inWhen uintptr, inChannel uint32, inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32

// AudioDriverPlugInStreamSetProperty.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInStreamSetProperty
func AudioDriverPlugInStreamSetProperty(inDevice uint32, inIOAudioStream uintptr, inWhen uintptr, inChannel uint32, inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32 {
	if _audioDriverPlugInStreamSetProperty == nil {
		panic("CoreAudio: symbol AudioDriverPlugInStreamSetProperty not loaded")
	}
	return _audioDriverPlugInStreamSetProperty(inDevice, inIOAudioStream, inWhen, inChannel, inPropertyID, inPropertyDataSize, inPropertyData)
}

var _audioGetCurrentHostTime func() uint64

// AudioGetCurrentHostTime.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioGetCurrentHostTime()
func AudioGetCurrentHostTime() uint64 {
	if _audioGetCurrentHostTime == nil {
		panic("CoreAudio: symbol AudioGetCurrentHostTime not loaded")
	}
	return _audioGetCurrentHostTime()
}

var _audioGetHostClockFrequency func() float64

// AudioGetHostClockFrequency.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioGetHostClockFrequency()
func AudioGetHostClockFrequency() float64 {
	if _audioGetHostClockFrequency == nil {
		panic("CoreAudio: symbol AudioGetHostClockFrequency not loaded")
	}
	return _audioGetHostClockFrequency()
}

var _audioGetHostClockMinimumTimeDelta func() uint32

// AudioGetHostClockMinimumTimeDelta.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioGetHostClockMinimumTimeDelta()
func AudioGetHostClockMinimumTimeDelta() uint32 {
	if _audioGetHostClockMinimumTimeDelta == nil {
		panic("CoreAudio: symbol AudioGetHostClockMinimumTimeDelta not loaded")
	}
	return _audioGetHostClockMinimumTimeDelta()
}

var _audioHardwareAddPropertyListener func(inPropertyID uint32, inProc AudioHardwarePropertyListenerProc, inClientData unsafe.Pointer) int32

// AudioHardwareAddPropertyListener.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareAddPropertyListener
func AudioHardwareAddPropertyListener(inPropertyID uint32, inProc AudioHardwarePropertyListenerProc, inClientData unsafe.Pointer) int32 {
	if _audioHardwareAddPropertyListener == nil {
		panic("CoreAudio: symbol AudioHardwareAddPropertyListener not loaded")
	}
	return _audioHardwareAddPropertyListener(inPropertyID, inProc, inClientData)
}

var _audioHardwareAddRunLoopSource func(inRunLoopSource corefoundation.CFRunLoopSourceRef) int32

// AudioHardwareAddRunLoopSource.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareAddRunLoopSource
func AudioHardwareAddRunLoopSource(inRunLoopSource corefoundation.CFRunLoopSourceRef) int32 {
	if _audioHardwareAddRunLoopSource == nil {
		panic("CoreAudio: symbol AudioHardwareAddRunLoopSource not loaded")
	}
	return _audioHardwareAddRunLoopSource(inRunLoopSource)
}

var _audioHardwareCreateAggregateDevice func(inDescription corefoundation.CFDictionaryRef, outDeviceID *uint32) int32

// AudioHardwareCreateAggregateDevice.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareCreateAggregateDevice(_:_:)
func AudioHardwareCreateAggregateDevice(inDescription corefoundation.CFDictionaryRef, outDeviceID *uint32) int32 {
	if _audioHardwareCreateAggregateDevice == nil {
		panic("CoreAudio: symbol AudioHardwareCreateAggregateDevice not loaded")
	}
	return _audioHardwareCreateAggregateDevice(inDescription, outDeviceID)
}

var _audioHardwareCreateProcessTap func(inDescription *CATapDescription, outTapID *uint32) int32

// AudioHardwareCreateProcessTap.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareCreateProcessTap(_:_:)
func AudioHardwareCreateProcessTap(inDescription *CATapDescription, outTapID *uint32) int32 {
	if _audioHardwareCreateProcessTap == nil {
		panic("CoreAudio: symbol AudioHardwareCreateProcessTap not loaded")
	}
	return _audioHardwareCreateProcessTap(inDescription, outTapID)
}

var _audioHardwareDestroyAggregateDevice func(inDeviceID uint32) int32

// AudioHardwareDestroyAggregateDevice.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareDestroyAggregateDevice(_:)
func AudioHardwareDestroyAggregateDevice(inDeviceID uint32) int32 {
	if _audioHardwareDestroyAggregateDevice == nil {
		panic("CoreAudio: symbol AudioHardwareDestroyAggregateDevice not loaded")
	}
	return _audioHardwareDestroyAggregateDevice(inDeviceID)
}

var _audioHardwareDestroyProcessTap func(inTapID uint32) int32

// AudioHardwareDestroyProcessTap.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareDestroyProcessTap(_:)
func AudioHardwareDestroyProcessTap(inTapID uint32) int32 {
	if _audioHardwareDestroyProcessTap == nil {
		panic("CoreAudio: symbol AudioHardwareDestroyProcessTap not loaded")
	}
	return _audioHardwareDestroyProcessTap(inTapID)
}

var _audioHardwareGetProperty func(inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32

// AudioHardwareGetProperty.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareGetProperty
func AudioHardwareGetProperty(inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	if _audioHardwareGetProperty == nil {
		panic("CoreAudio: symbol AudioHardwareGetProperty not loaded")
	}
	return _audioHardwareGetProperty(inPropertyID, ioPropertyDataSize, outPropertyData)
}

var _audioHardwareGetPropertyInfo func(inPropertyID uint32, outSize *uint32, outWritable *bool) int32

// AudioHardwareGetPropertyInfo.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareGetPropertyInfo
func AudioHardwareGetPropertyInfo(inPropertyID uint32, outSize *uint32, outWritable *bool) int32 {
	if _audioHardwareGetPropertyInfo == nil {
		panic("CoreAudio: symbol AudioHardwareGetPropertyInfo not loaded")
	}
	return _audioHardwareGetPropertyInfo(inPropertyID, outSize, outWritable)
}

var _audioHardwareRemovePropertyListener func(inPropertyID uint32, inProc AudioHardwarePropertyListenerProc) int32

// AudioHardwareRemovePropertyListener.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareRemovePropertyListener
func AudioHardwareRemovePropertyListener(inPropertyID uint32, inProc AudioHardwarePropertyListenerProc) int32 {
	if _audioHardwareRemovePropertyListener == nil {
		panic("CoreAudio: symbol AudioHardwareRemovePropertyListener not loaded")
	}
	return _audioHardwareRemovePropertyListener(inPropertyID, inProc)
}

var _audioHardwareRemoveRunLoopSource func(inRunLoopSource corefoundation.CFRunLoopSourceRef) int32

// AudioHardwareRemoveRunLoopSource.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareRemoveRunLoopSource
func AudioHardwareRemoveRunLoopSource(inRunLoopSource corefoundation.CFRunLoopSourceRef) int32 {
	if _audioHardwareRemoveRunLoopSource == nil {
		panic("CoreAudio: symbol AudioHardwareRemoveRunLoopSource not loaded")
	}
	return _audioHardwareRemoveRunLoopSource(inRunLoopSource)
}

var _audioHardwareSetProperty func(inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32

// AudioHardwareSetProperty.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareSetProperty
func AudioHardwareSetProperty(inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32 {
	if _audioHardwareSetProperty == nil {
		panic("CoreAudio: symbol AudioHardwareSetProperty not loaded")
	}
	return _audioHardwareSetProperty(inPropertyID, inPropertyDataSize, inPropertyData)
}

var _audioHardwareUnload func() int32

// AudioHardwareUnload.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareUnload()
func AudioHardwareUnload() int32 {
	if _audioHardwareUnload == nil {
		panic("CoreAudio: symbol AudioHardwareUnload not loaded")
	}
	return _audioHardwareUnload()
}

var _audioObjectAddPropertyListener func(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inListener AudioObjectPropertyListenerProc, inClientData unsafe.Pointer) int32

// AudioObjectAddPropertyListener.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectAddPropertyListener(_:_:_:_:)
func AudioObjectAddPropertyListener(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inListener AudioObjectPropertyListenerProc, inClientData unsafe.Pointer) int32 {
	if _audioObjectAddPropertyListener == nil {
		panic("CoreAudio: symbol AudioObjectAddPropertyListener not loaded")
	}
	return _audioObjectAddPropertyListener(inObjectID, inAddress, inListener, inClientData)
}

var _audioObjectAddPropertyListenerBlock func(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inDispatchQueue uintptr, inListener AudioObjectPropertyListenerBlock) int32

// AudioObjectAddPropertyListenerBlock.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectAddPropertyListenerBlock(_:_:_:_:)
func AudioObjectAddPropertyListenerBlock(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inDispatchQueue dispatch.Queue, inListener AudioObjectPropertyListenerBlock) int32 {
	if _audioObjectAddPropertyListenerBlock == nil {
		panic("CoreAudio: symbol AudioObjectAddPropertyListenerBlock not loaded")
	}
	return _audioObjectAddPropertyListenerBlock(inObjectID, inAddress, uintptr(inDispatchQueue.Handle()), inListener)
}

var _audioObjectGetPropertyData func(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inQualifierDataSize uint32, inQualifierData unsafe.Pointer, ioDataSize *uint32, outData unsafe.Pointer) int32

// AudioObjectGetPropertyData.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectGetPropertyData(_:_:_:_:_:_:)
func AudioObjectGetPropertyData(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inQualifierDataSize uint32, inQualifierData unsafe.Pointer, ioDataSize *uint32, outData unsafe.Pointer) int32 {
	if _audioObjectGetPropertyData == nil {
		panic("CoreAudio: symbol AudioObjectGetPropertyData not loaded")
	}
	return _audioObjectGetPropertyData(inObjectID, inAddress, inQualifierDataSize, inQualifierData, ioDataSize, outData)
}

var _audioObjectGetPropertyDataSize func(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inQualifierDataSize uint32, inQualifierData unsafe.Pointer, outDataSize *uint32) int32

// AudioObjectGetPropertyDataSize.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectGetPropertyDataSize(_:_:_:_:_:)
func AudioObjectGetPropertyDataSize(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inQualifierDataSize uint32, inQualifierData unsafe.Pointer, outDataSize *uint32) int32 {
	if _audioObjectGetPropertyDataSize == nil {
		panic("CoreAudio: symbol AudioObjectGetPropertyDataSize not loaded")
	}
	return _audioObjectGetPropertyDataSize(inObjectID, inAddress, inQualifierDataSize, inQualifierData, outDataSize)
}

var _audioObjectHasProperty func(inObjectID uint32, inAddress *AudioObjectPropertyAddress) bool

// AudioObjectHasProperty.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectHasProperty(_:_:)
func AudioObjectHasProperty(inObjectID uint32, inAddress *AudioObjectPropertyAddress) bool {
	if _audioObjectHasProperty == nil {
		panic("CoreAudio: symbol AudioObjectHasProperty not loaded")
	}
	return _audioObjectHasProperty(inObjectID, inAddress)
}

var _audioObjectIsPropertySettable func(inObjectID uint32, inAddress *AudioObjectPropertyAddress, outIsSettable *bool) int32

// AudioObjectIsPropertySettable.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectIsPropertySettable(_:_:_:)
func AudioObjectIsPropertySettable(inObjectID uint32, inAddress *AudioObjectPropertyAddress, outIsSettable *bool) int32 {
	if _audioObjectIsPropertySettable == nil {
		panic("CoreAudio: symbol AudioObjectIsPropertySettable not loaded")
	}
	return _audioObjectIsPropertySettable(inObjectID, inAddress, outIsSettable)
}

var _audioObjectRemovePropertyListener func(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inListener AudioObjectPropertyListenerProc, inClientData unsafe.Pointer) int32

// AudioObjectRemovePropertyListener.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectRemovePropertyListener(_:_:_:_:)
func AudioObjectRemovePropertyListener(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inListener AudioObjectPropertyListenerProc, inClientData unsafe.Pointer) int32 {
	if _audioObjectRemovePropertyListener == nil {
		panic("CoreAudio: symbol AudioObjectRemovePropertyListener not loaded")
	}
	return _audioObjectRemovePropertyListener(inObjectID, inAddress, inListener, inClientData)
}

var _audioObjectRemovePropertyListenerBlock func(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inDispatchQueue uintptr, inListener AudioObjectPropertyListenerBlock) int32

// AudioObjectRemovePropertyListenerBlock.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectRemovePropertyListenerBlock(_:_:_:_:)
func AudioObjectRemovePropertyListenerBlock(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inDispatchQueue dispatch.Queue, inListener AudioObjectPropertyListenerBlock) int32 {
	if _audioObjectRemovePropertyListenerBlock == nil {
		panic("CoreAudio: symbol AudioObjectRemovePropertyListenerBlock not loaded")
	}
	return _audioObjectRemovePropertyListenerBlock(inObjectID, inAddress, uintptr(inDispatchQueue.Handle()), inListener)
}

var _audioObjectSetPropertyData func(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inQualifierDataSize uint32, inQualifierData unsafe.Pointer, inDataSize uint32, inData unsafe.Pointer) int32

// AudioObjectSetPropertyData.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectSetPropertyData(_:_:_:_:_:_:)
func AudioObjectSetPropertyData(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inQualifierDataSize uint32, inQualifierData unsafe.Pointer, inDataSize uint32, inData unsafe.Pointer) int32 {
	if _audioObjectSetPropertyData == nil {
		panic("CoreAudio: symbol AudioObjectSetPropertyData not loaded")
	}
	return _audioObjectSetPropertyData(inObjectID, inAddress, inQualifierDataSize, inQualifierData, inDataSize, inData)
}

var _audioObjectShow func(inObjectID uint32)

// AudioObjectShow.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectShow(_:)
func AudioObjectShow(inObjectID uint32) {
	if _audioObjectShow == nil {
		panic("CoreAudio: symbol AudioObjectShow not loaded")
	}
	_audioObjectShow(inObjectID)
}

var _audioStreamAddPropertyListener func(inStream uint32, inChannel uint32, inPropertyID uint32, inProc AudioStreamPropertyListenerProc, inClientData unsafe.Pointer) int32

// AudioStreamAddPropertyListener.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioStreamAddPropertyListener
func AudioStreamAddPropertyListener(inStream uint32, inChannel uint32, inPropertyID uint32, inProc AudioStreamPropertyListenerProc, inClientData unsafe.Pointer) int32 {
	if _audioStreamAddPropertyListener == nil {
		panic("CoreAudio: symbol AudioStreamAddPropertyListener not loaded")
	}
	return _audioStreamAddPropertyListener(inStream, inChannel, inPropertyID, inProc, inClientData)
}

var _audioStreamGetProperty func(inStream uint32, inChannel uint32, inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32

// AudioStreamGetProperty.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioStreamGetProperty
func AudioStreamGetProperty(inStream uint32, inChannel uint32, inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	if _audioStreamGetProperty == nil {
		panic("CoreAudio: symbol AudioStreamGetProperty not loaded")
	}
	return _audioStreamGetProperty(inStream, inChannel, inPropertyID, ioPropertyDataSize, outPropertyData)
}

var _audioStreamGetPropertyInfo func(inStream uint32, inChannel uint32, inPropertyID uint32, outSize *uint32, outWritable *bool) int32

// AudioStreamGetPropertyInfo.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioStreamGetPropertyInfo
func AudioStreamGetPropertyInfo(inStream uint32, inChannel uint32, inPropertyID uint32, outSize *uint32, outWritable *bool) int32 {
	if _audioStreamGetPropertyInfo == nil {
		panic("CoreAudio: symbol AudioStreamGetPropertyInfo not loaded")
	}
	return _audioStreamGetPropertyInfo(inStream, inChannel, inPropertyID, outSize, outWritable)
}

var _audioStreamRemovePropertyListener func(inStream uint32, inChannel uint32, inPropertyID uint32, inProc AudioStreamPropertyListenerProc) int32

// AudioStreamRemovePropertyListener.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioStreamRemovePropertyListener
func AudioStreamRemovePropertyListener(inStream uint32, inChannel uint32, inPropertyID uint32, inProc AudioStreamPropertyListenerProc) int32 {
	if _audioStreamRemovePropertyListener == nil {
		panic("CoreAudio: symbol AudioStreamRemovePropertyListener not loaded")
	}
	return _audioStreamRemovePropertyListener(inStream, inChannel, inPropertyID, inProc)
}

var _audioStreamSetProperty func(inStream uint32, inWhen uintptr, inChannel uint32, inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32

// AudioStreamSetProperty.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioStreamSetProperty
func AudioStreamSetProperty(inStream uint32, inWhen uintptr, inChannel uint32, inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32 {
	if _audioStreamSetProperty == nil {
		panic("CoreAudio: symbol AudioStreamSetProperty not loaded")
	}
	return _audioStreamSetProperty(inStream, inWhen, inChannel, inPropertyID, inPropertyDataSize, inPropertyData)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_audioConvertHostTimeToNanos, frameworkHandle, "AudioConvertHostTimeToNanos")
	registerFunc(&_audioConvertNanosToHostTime, frameworkHandle, "AudioConvertNanosToHostTime")
	registerFunc(&_audioDeviceAddIOProc, frameworkHandle, "AudioDeviceAddIOProc")
	registerFunc(&_audioDeviceAddPropertyListener, frameworkHandle, "AudioDeviceAddPropertyListener")
	registerFunc(&_audioDeviceCreateIOProcID, frameworkHandle, "AudioDeviceCreateIOProcID")
	registerFunc(&_audioDeviceCreateIOProcIDWithBlock, frameworkHandle, "AudioDeviceCreateIOProcIDWithBlock")
	registerFunc(&_audioDeviceDestroyIOProcID, frameworkHandle, "AudioDeviceDestroyIOProcID")
	registerFunc(&_audioDeviceGetCurrentTime, frameworkHandle, "AudioDeviceGetCurrentTime")
	registerFunc(&_audioDeviceGetNearestStartTime, frameworkHandle, "AudioDeviceGetNearestStartTime")
	registerFunc(&_audioDeviceGetProperty, frameworkHandle, "AudioDeviceGetProperty")
	registerFunc(&_audioDeviceGetPropertyInfo, frameworkHandle, "AudioDeviceGetPropertyInfo")
	registerFunc(&_audioDeviceRead, frameworkHandle, "AudioDeviceRead")
	registerFunc(&_audioDeviceRemoveIOProc, frameworkHandle, "AudioDeviceRemoveIOProc")
	registerFunc(&_audioDeviceRemovePropertyListener, frameworkHandle, "AudioDeviceRemovePropertyListener")
	registerFunc(&_audioDeviceSetProperty, frameworkHandle, "AudioDeviceSetProperty")
	registerFunc(&_audioDeviceStart, frameworkHandle, "AudioDeviceStart")
	registerFunc(&_audioDeviceStartAtTime, frameworkHandle, "AudioDeviceStartAtTime")
	registerFunc(&_audioDeviceStop, frameworkHandle, "AudioDeviceStop")
	registerFunc(&_audioDeviceTranslateTime, frameworkHandle, "AudioDeviceTranslateTime")
	registerFunc(&_audioDriverPlugInClose, frameworkHandle, "AudioDriverPlugInClose")
	registerFunc(&_audioDriverPlugInDeviceGetProperty, frameworkHandle, "AudioDriverPlugInDeviceGetProperty")
	registerFunc(&_audioDriverPlugInDeviceGetPropertyInfo, frameworkHandle, "AudioDriverPlugInDeviceGetPropertyInfo")
	registerFunc(&_audioDriverPlugInDeviceSetProperty, frameworkHandle, "AudioDriverPlugInDeviceSetProperty")
	registerFunc(&_audioDriverPlugInOpen, frameworkHandle, "AudioDriverPlugInOpen")
	registerFunc(&_audioDriverPlugInStreamGetProperty, frameworkHandle, "AudioDriverPlugInStreamGetProperty")
	registerFunc(&_audioDriverPlugInStreamGetPropertyInfo, frameworkHandle, "AudioDriverPlugInStreamGetPropertyInfo")
	registerFunc(&_audioDriverPlugInStreamSetProperty, frameworkHandle, "AudioDriverPlugInStreamSetProperty")
	registerFunc(&_audioGetCurrentHostTime, frameworkHandle, "AudioGetCurrentHostTime")
	registerFunc(&_audioGetHostClockFrequency, frameworkHandle, "AudioGetHostClockFrequency")
	registerFunc(&_audioGetHostClockMinimumTimeDelta, frameworkHandle, "AudioGetHostClockMinimumTimeDelta")
	registerFunc(&_audioHardwareAddPropertyListener, frameworkHandle, "AudioHardwareAddPropertyListener")
	registerFunc(&_audioHardwareAddRunLoopSource, frameworkHandle, "AudioHardwareAddRunLoopSource")
	registerFunc(&_audioHardwareCreateAggregateDevice, frameworkHandle, "AudioHardwareCreateAggregateDevice")
	registerFunc(&_audioHardwareCreateProcessTap, frameworkHandle, "AudioHardwareCreateProcessTap")
	registerFunc(&_audioHardwareDestroyAggregateDevice, frameworkHandle, "AudioHardwareDestroyAggregateDevice")
	registerFunc(&_audioHardwareDestroyProcessTap, frameworkHandle, "AudioHardwareDestroyProcessTap")
	registerFunc(&_audioHardwareGetProperty, frameworkHandle, "AudioHardwareGetProperty")
	registerFunc(&_audioHardwareGetPropertyInfo, frameworkHandle, "AudioHardwareGetPropertyInfo")
	registerFunc(&_audioHardwareRemovePropertyListener, frameworkHandle, "AudioHardwareRemovePropertyListener")
	registerFunc(&_audioHardwareRemoveRunLoopSource, frameworkHandle, "AudioHardwareRemoveRunLoopSource")
	registerFunc(&_audioHardwareSetProperty, frameworkHandle, "AudioHardwareSetProperty")
	registerFunc(&_audioHardwareUnload, frameworkHandle, "AudioHardwareUnload")
	registerFunc(&_audioObjectAddPropertyListener, frameworkHandle, "AudioObjectAddPropertyListener")
	registerFunc(&_audioObjectAddPropertyListenerBlock, frameworkHandle, "AudioObjectAddPropertyListenerBlock")
	registerFunc(&_audioObjectGetPropertyData, frameworkHandle, "AudioObjectGetPropertyData")
	registerFunc(&_audioObjectGetPropertyDataSize, frameworkHandle, "AudioObjectGetPropertyDataSize")
	registerFunc(&_audioObjectHasProperty, frameworkHandle, "AudioObjectHasProperty")
	registerFunc(&_audioObjectIsPropertySettable, frameworkHandle, "AudioObjectIsPropertySettable")
	registerFunc(&_audioObjectRemovePropertyListener, frameworkHandle, "AudioObjectRemovePropertyListener")
	registerFunc(&_audioObjectRemovePropertyListenerBlock, frameworkHandle, "AudioObjectRemovePropertyListenerBlock")
	registerFunc(&_audioObjectSetPropertyData, frameworkHandle, "AudioObjectSetPropertyData")
	registerFunc(&_audioObjectShow, frameworkHandle, "AudioObjectShow")
	registerFunc(&_audioStreamAddPropertyListener, frameworkHandle, "AudioStreamAddPropertyListener")
	registerFunc(&_audioStreamGetProperty, frameworkHandle, "AudioStreamGetProperty")
	registerFunc(&_audioStreamGetPropertyInfo, frameworkHandle, "AudioStreamGetPropertyInfo")
	registerFunc(&_audioStreamRemovePropertyListener, frameworkHandle, "AudioStreamRemovePropertyListener")
	registerFunc(&_audioStreamSetProperty, frameworkHandle, "AudioStreamSetProperty")
}
