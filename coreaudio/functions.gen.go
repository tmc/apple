// Code generated from Apple documentation for CoreAudio. DO NOT EDIT.

package coreaudio

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objc"
)

type unavailableSymbolError struct {
	symbol     string
	introduced string
	cause      error
}

func (e *unavailableSymbolError) Error() string {
	if e == nil {
		return ""
	}
	if e.introduced != "" {
		return fmt.Sprintf("CoreAudio: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("CoreAudio: symbol %s unavailable on this system", e.symbol)
}

func (e *unavailableSymbolError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

func missingSymbolError(name, introduced string, cause error) error {
	return &unavailableSymbolError{
		symbol:     name,
		introduced: introduced,
		cause:      cause,
	}
}

func symbolCallError(name, introduced string, err error) error {
	if err != nil {
		return err
	}
	if frameworkHandle == 0 {
		return fmt.Errorf("CoreAudio: symbol %s unavailable because the framework could not be loaded", name)
	}
	return missingSymbolError(name, introduced, nil)
}

// registerFunc resolves a framework symbol and registers it as a Go function.
func registerFunc(fptr any, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			*errDst = fmt.Errorf("CoreAudio: register symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	*errDst = nil
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	*dst = sym
	*errDst = nil
}

var _audioConvertHostTimeToNanos func(inHostTime uint64) uint64
var _audioConvertHostTimeToNanosErr error

func tryAudioConvertHostTimeToNanos(inHostTime uint64) (uint64, error) {
	if _audioConvertHostTimeToNanos == nil {
		return 0, symbolCallError("AudioConvertHostTimeToNanos", "10.0", _audioConvertHostTimeToNanosErr)
	}
	return _audioConvertHostTimeToNanos(inHostTime), nil
}

// AudioConvertHostTimeToNanos.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioConvertHostTimeToNanos(_:)
func AudioConvertHostTimeToNanos(inHostTime uint64) uint64 {
	result, callErr := tryAudioConvertHostTimeToNanos(inHostTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioConvertNanosToHostTime func(inNanos uint64) uint64
var _audioConvertNanosToHostTimeErr error

func tryAudioConvertNanosToHostTime(inNanos uint64) (uint64, error) {
	if _audioConvertNanosToHostTime == nil {
		return 0, symbolCallError("AudioConvertNanosToHostTime", "10.0", _audioConvertNanosToHostTimeErr)
	}
	return _audioConvertNanosToHostTime(inNanos), nil
}

// AudioConvertNanosToHostTime.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioConvertNanosToHostTime(_:)
func AudioConvertNanosToHostTime(inNanos uint64) uint64 {
	result, callErr := tryAudioConvertNanosToHostTime(inNanos)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDeviceAddIOProc func(inDevice uint32, inProc AudioDeviceIOProc, inClientData unsafe.Pointer) int32
var _audioDeviceAddIOProcErr error

func tryAudioDeviceAddIOProc(inDevice uint32, inProc AudioDeviceIOProc, inClientData unsafe.Pointer) (int32, error) {
	if _audioDeviceAddIOProc == nil {
		return 0, symbolCallError("AudioDeviceAddIOProc", "10.0", _audioDeviceAddIOProcErr)
	}
	return _audioDeviceAddIOProc(inDevice, inProc, inClientData), nil
}

// AudioDeviceAddIOProc.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceAddIOProc
func AudioDeviceAddIOProc(inDevice uint32, inProc AudioDeviceIOProc, inClientData unsafe.Pointer) int32 {
	result, callErr := tryAudioDeviceAddIOProc(inDevice, inProc, inClientData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDeviceAddPropertyListener func(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, inProc AudioDevicePropertyListenerProc, inClientData unsafe.Pointer) int32
var _audioDeviceAddPropertyListenerErr error

func tryAudioDeviceAddPropertyListener(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, inProc AudioDevicePropertyListenerProc, inClientData unsafe.Pointer) (int32, error) {
	if _audioDeviceAddPropertyListener == nil {
		return 0, symbolCallError("AudioDeviceAddPropertyListener", "10.0", _audioDeviceAddPropertyListenerErr)
	}
	return _audioDeviceAddPropertyListener(inDevice, inChannel, isInput, inPropertyID, inProc, inClientData), nil
}

// AudioDeviceAddPropertyListener.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceAddPropertyListener
func AudioDeviceAddPropertyListener(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, inProc AudioDevicePropertyListenerProc, inClientData unsafe.Pointer) int32 {
	result, callErr := tryAudioDeviceAddPropertyListener(inDevice, inChannel, isInput, inPropertyID, inProc, inClientData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDeviceCreateIOProcID func(inDevice uint32, inProc AudioDeviceIOProc, inClientData unsafe.Pointer, outIOProcID *AudioDeviceIOProcID) int32
var _audioDeviceCreateIOProcIDErr error

func tryAudioDeviceCreateIOProcID(inDevice uint32, inProc AudioDeviceIOProc, inClientData unsafe.Pointer, outIOProcID *AudioDeviceIOProcID) (int32, error) {
	if _audioDeviceCreateIOProcID == nil {
		return 0, symbolCallError("AudioDeviceCreateIOProcID", "10.5", _audioDeviceCreateIOProcIDErr)
	}
	return _audioDeviceCreateIOProcID(inDevice, inProc, inClientData, outIOProcID), nil
}

// AudioDeviceCreateIOProcID.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceCreateIOProcID(_:_:_:_:)
func AudioDeviceCreateIOProcID(inDevice uint32, inProc AudioDeviceIOProc, inClientData unsafe.Pointer, outIOProcID *AudioDeviceIOProcID) int32 {
	result, callErr := tryAudioDeviceCreateIOProcID(inDevice, inProc, inClientData, outIOProcID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDeviceCreateIOProcIDWithBlock func(outIOProcID *AudioDeviceIOProcID, inDevice uint32, inDispatchQueue uintptr, inIOBlock unsafe.Pointer) int32
var _audioDeviceCreateIOProcIDWithBlockErr error

func tryAudioDeviceCreateIOProcIDWithBlock(outIOProcID *AudioDeviceIOProcID, inDevice uint32, inDispatchQueue dispatch.Queue, inIOBlock AudioDeviceIOBlock) (int32, error) {
	if _audioDeviceCreateIOProcIDWithBlock == nil {
		return 0, symbolCallError("AudioDeviceCreateIOProcIDWithBlock", "10.7", _audioDeviceCreateIOProcIDWithBlockErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 unsafe.Pointer, arg1 unsafe.Pointer, arg2 unsafe.Pointer, arg3 unsafe.Pointer, arg4 unsafe.Pointer) {
		inIOBlock(arg0, arg1, arg2, arg3, arg4)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _audioDeviceCreateIOProcIDWithBlock(outIOProcID, inDevice, uintptr(inDispatchQueue.Handle()), _block0), nil
}

// AudioDeviceCreateIOProcIDWithBlock.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceCreateIOProcIDWithBlock(_:_:_:_:)
func AudioDeviceCreateIOProcIDWithBlock(outIOProcID *AudioDeviceIOProcID, inDevice uint32, inDispatchQueue dispatch.Queue, inIOBlock AudioDeviceIOBlock) int32 {
	result, callErr := tryAudioDeviceCreateIOProcIDWithBlock(outIOProcID, inDevice, inDispatchQueue, inIOBlock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDeviceDestroyIOProcID func(inDevice uint32, inIOProcID AudioDeviceIOProcID) int32
var _audioDeviceDestroyIOProcIDErr error

func tryAudioDeviceDestroyIOProcID(inDevice uint32, inIOProcID AudioDeviceIOProcID) (int32, error) {
	if _audioDeviceDestroyIOProcID == nil {
		return 0, symbolCallError("AudioDeviceDestroyIOProcID", "10.5", _audioDeviceDestroyIOProcIDErr)
	}
	return _audioDeviceDestroyIOProcID(inDevice, inIOProcID), nil
}

// AudioDeviceDestroyIOProcID.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceDestroyIOProcID(_:_:)
func AudioDeviceDestroyIOProcID(inDevice uint32, inIOProcID AudioDeviceIOProcID) int32 {
	result, callErr := tryAudioDeviceDestroyIOProcID(inDevice, inIOProcID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDeviceGetCurrentTime func(inDevice uint32, outTime unsafe.Pointer) int32
var _audioDeviceGetCurrentTimeErr error

func tryAudioDeviceGetCurrentTime(inDevice uint32, outTime unsafe.Pointer) (int32, error) {
	if _audioDeviceGetCurrentTime == nil {
		return 0, symbolCallError("AudioDeviceGetCurrentTime", "10.0", _audioDeviceGetCurrentTimeErr)
	}
	return _audioDeviceGetCurrentTime(inDevice, outTime), nil
}

// AudioDeviceGetCurrentTime.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceGetCurrentTime(_:_:)
func AudioDeviceGetCurrentTime(inDevice uint32, outTime unsafe.Pointer) int32 {
	result, callErr := tryAudioDeviceGetCurrentTime(inDevice, outTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDeviceGetNearestStartTime func(inDevice uint32, ioRequestedStartTime unsafe.Pointer, inFlags uint32) int32
var _audioDeviceGetNearestStartTimeErr error

func tryAudioDeviceGetNearestStartTime(inDevice uint32, ioRequestedStartTime unsafe.Pointer, inFlags uint32) (int32, error) {
	if _audioDeviceGetNearestStartTime == nil {
		return 0, symbolCallError("AudioDeviceGetNearestStartTime", "10.3", _audioDeviceGetNearestStartTimeErr)
	}
	return _audioDeviceGetNearestStartTime(inDevice, ioRequestedStartTime, inFlags), nil
}

// AudioDeviceGetNearestStartTime.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceGetNearestStartTime(_:_:_:)
func AudioDeviceGetNearestStartTime(inDevice uint32, ioRequestedStartTime unsafe.Pointer, inFlags uint32) int32 {
	result, callErr := tryAudioDeviceGetNearestStartTime(inDevice, ioRequestedStartTime, inFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDeviceGetProperty func(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32
var _audioDeviceGetPropertyErr error

func tryAudioDeviceGetProperty(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) (int32, error) {
	if _audioDeviceGetProperty == nil {
		return 0, symbolCallError("AudioDeviceGetProperty", "10.0", _audioDeviceGetPropertyErr)
	}
	return _audioDeviceGetProperty(inDevice, inChannel, isInput, inPropertyID, ioPropertyDataSize, outPropertyData), nil
}

// AudioDeviceGetProperty.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceGetProperty
func AudioDeviceGetProperty(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioDeviceGetProperty(inDevice, inChannel, isInput, inPropertyID, ioPropertyDataSize, outPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDeviceGetPropertyInfo func(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, outSize *uint32, outWritable *bool) int32
var _audioDeviceGetPropertyInfoErr error

func tryAudioDeviceGetPropertyInfo(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, outSize *uint32, outWritable *bool) (int32, error) {
	if _audioDeviceGetPropertyInfo == nil {
		return 0, symbolCallError("AudioDeviceGetPropertyInfo", "10.0", _audioDeviceGetPropertyInfoErr)
	}
	return _audioDeviceGetPropertyInfo(inDevice, inChannel, isInput, inPropertyID, outSize, outWritable), nil
}

// AudioDeviceGetPropertyInfo.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceGetPropertyInfo
func AudioDeviceGetPropertyInfo(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, outSize *uint32, outWritable *bool) int32 {
	result, callErr := tryAudioDeviceGetPropertyInfo(inDevice, inChannel, isInput, inPropertyID, outSize, outWritable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDeviceRead func(inDevice uint32, inStartTime unsafe.Pointer, outData unsafe.Pointer) int32
var _audioDeviceReadErr error

func tryAudioDeviceRead(inDevice uint32, inStartTime unsafe.Pointer, outData unsafe.Pointer) (int32, error) {
	if _audioDeviceRead == nil {
		return 0, symbolCallError("AudioDeviceRead", "10.1", _audioDeviceReadErr)
	}
	return _audioDeviceRead(inDevice, inStartTime, outData), nil
}

// AudioDeviceRead.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceRead
func AudioDeviceRead(inDevice uint32, inStartTime unsafe.Pointer, outData unsafe.Pointer) int32 {
	result, callErr := tryAudioDeviceRead(inDevice, inStartTime, outData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDeviceRemoveIOProc func(inDevice uint32, inProc AudioDeviceIOProc) int32
var _audioDeviceRemoveIOProcErr error

func tryAudioDeviceRemoveIOProc(inDevice uint32, inProc AudioDeviceIOProc) (int32, error) {
	if _audioDeviceRemoveIOProc == nil {
		return 0, symbolCallError("AudioDeviceRemoveIOProc", "10.0", _audioDeviceRemoveIOProcErr)
	}
	return _audioDeviceRemoveIOProc(inDevice, inProc), nil
}

// AudioDeviceRemoveIOProc.
//
// Deprecated: Deprecated since macOS 10.5.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceRemoveIOProc
func AudioDeviceRemoveIOProc(inDevice uint32, inProc AudioDeviceIOProc) int32 {
	result, callErr := tryAudioDeviceRemoveIOProc(inDevice, inProc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDeviceRemovePropertyListener func(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, inProc AudioDevicePropertyListenerProc) int32
var _audioDeviceRemovePropertyListenerErr error

func tryAudioDeviceRemovePropertyListener(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, inProc AudioDevicePropertyListenerProc) (int32, error) {
	if _audioDeviceRemovePropertyListener == nil {
		return 0, symbolCallError("AudioDeviceRemovePropertyListener", "10.0", _audioDeviceRemovePropertyListenerErr)
	}
	return _audioDeviceRemovePropertyListener(inDevice, inChannel, isInput, inPropertyID, inProc), nil
}

// AudioDeviceRemovePropertyListener.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceRemovePropertyListener
func AudioDeviceRemovePropertyListener(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, inProc AudioDevicePropertyListenerProc) int32 {
	result, callErr := tryAudioDeviceRemovePropertyListener(inDevice, inChannel, isInput, inPropertyID, inProc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDeviceSetProperty func(inDevice uint32, inWhen unsafe.Pointer, inChannel uint32, isInput bool, inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32
var _audioDeviceSetPropertyErr error

func tryAudioDeviceSetProperty(inDevice uint32, inWhen unsafe.Pointer, inChannel uint32, isInput bool, inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) (int32, error) {
	if _audioDeviceSetProperty == nil {
		return 0, symbolCallError("AudioDeviceSetProperty", "10.0", _audioDeviceSetPropertyErr)
	}
	return _audioDeviceSetProperty(inDevice, inWhen, inChannel, isInput, inPropertyID, inPropertyDataSize, inPropertyData), nil
}

// AudioDeviceSetProperty.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceSetProperty
func AudioDeviceSetProperty(inDevice uint32, inWhen unsafe.Pointer, inChannel uint32, isInput bool, inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioDeviceSetProperty(inDevice, inWhen, inChannel, isInput, inPropertyID, inPropertyDataSize, inPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDeviceStart func(inDevice uint32, inProcID AudioDeviceIOProcID) int32
var _audioDeviceStartErr error

func tryAudioDeviceStart(inDevice uint32, inProcID AudioDeviceIOProcID) (int32, error) {
	if _audioDeviceStart == nil {
		return 0, symbolCallError("AudioDeviceStart", "10.0", _audioDeviceStartErr)
	}
	return _audioDeviceStart(inDevice, inProcID), nil
}

// AudioDeviceStart.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceStart(_:_:)
func AudioDeviceStart(inDevice uint32, inProcID AudioDeviceIOProcID) int32 {
	result, callErr := tryAudioDeviceStart(inDevice, inProcID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDeviceStartAtTime func(inDevice uint32, inProcID AudioDeviceIOProcID, ioRequestedStartTime unsafe.Pointer, inFlags uint32) int32
var _audioDeviceStartAtTimeErr error

func tryAudioDeviceStartAtTime(inDevice uint32, inProcID AudioDeviceIOProcID, ioRequestedStartTime unsafe.Pointer, inFlags uint32) (int32, error) {
	if _audioDeviceStartAtTime == nil {
		return 0, symbolCallError("AudioDeviceStartAtTime", "10.3", _audioDeviceStartAtTimeErr)
	}
	return _audioDeviceStartAtTime(inDevice, inProcID, ioRequestedStartTime, inFlags), nil
}

// AudioDeviceStartAtTime.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceStartAtTime(_:_:_:_:)
func AudioDeviceStartAtTime(inDevice uint32, inProcID AudioDeviceIOProcID, ioRequestedStartTime unsafe.Pointer, inFlags uint32) int32 {
	result, callErr := tryAudioDeviceStartAtTime(inDevice, inProcID, ioRequestedStartTime, inFlags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDeviceStop func(inDevice uint32, inProcID AudioDeviceIOProcID) int32
var _audioDeviceStopErr error

func tryAudioDeviceStop(inDevice uint32, inProcID AudioDeviceIOProcID) (int32, error) {
	if _audioDeviceStop == nil {
		return 0, symbolCallError("AudioDeviceStop", "10.0", _audioDeviceStopErr)
	}
	return _audioDeviceStop(inDevice, inProcID), nil
}

// AudioDeviceStop.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceStop(_:_:)
func AudioDeviceStop(inDevice uint32, inProcID AudioDeviceIOProcID) int32 {
	result, callErr := tryAudioDeviceStop(inDevice, inProcID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDeviceTranslateTime func(inDevice uint32, inTime unsafe.Pointer, outTime unsafe.Pointer) int32
var _audioDeviceTranslateTimeErr error

func tryAudioDeviceTranslateTime(inDevice uint32, inTime unsafe.Pointer, outTime unsafe.Pointer) (int32, error) {
	if _audioDeviceTranslateTime == nil {
		return 0, symbolCallError("AudioDeviceTranslateTime", "10.0", _audioDeviceTranslateTimeErr)
	}
	return _audioDeviceTranslateTime(inDevice, inTime, outTime), nil
}

// AudioDeviceTranslateTime.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDeviceTranslateTime(_:_:_:)
func AudioDeviceTranslateTime(inDevice uint32, inTime unsafe.Pointer, outTime unsafe.Pointer) int32 {
	result, callErr := tryAudioDeviceTranslateTime(inDevice, inTime, outTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDriverPlugInClose func(inDevice uint32) int32
var _audioDriverPlugInCloseErr error

func tryAudioDriverPlugInClose(inDevice uint32) (int32, error) {
	if _audioDriverPlugInClose == nil {
		return 0, symbolCallError("AudioDriverPlugInClose", "", _audioDriverPlugInCloseErr)
	}
	return _audioDriverPlugInClose(inDevice), nil
}

// AudioDriverPlugInClose.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInClose
func AudioDriverPlugInClose(inDevice uint32) int32 {
	result, callErr := tryAudioDriverPlugInClose(inDevice)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDriverPlugInDeviceGetProperty func(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32
var _audioDriverPlugInDeviceGetPropertyErr error

func tryAudioDriverPlugInDeviceGetProperty(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) (int32, error) {
	if _audioDriverPlugInDeviceGetProperty == nil {
		return 0, symbolCallError("AudioDriverPlugInDeviceGetProperty", "", _audioDriverPlugInDeviceGetPropertyErr)
	}
	return _audioDriverPlugInDeviceGetProperty(inDevice, inChannel, isInput, inPropertyID, ioPropertyDataSize, outPropertyData), nil
}

// AudioDriverPlugInDeviceGetProperty.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInDeviceGetProperty
func AudioDriverPlugInDeviceGetProperty(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioDriverPlugInDeviceGetProperty(inDevice, inChannel, isInput, inPropertyID, ioPropertyDataSize, outPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDriverPlugInDeviceGetPropertyInfo func(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, outSize *uint32, outWritable *bool) int32
var _audioDriverPlugInDeviceGetPropertyInfoErr error

func tryAudioDriverPlugInDeviceGetPropertyInfo(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, outSize *uint32, outWritable *bool) (int32, error) {
	if _audioDriverPlugInDeviceGetPropertyInfo == nil {
		return 0, symbolCallError("AudioDriverPlugInDeviceGetPropertyInfo", "", _audioDriverPlugInDeviceGetPropertyInfoErr)
	}
	return _audioDriverPlugInDeviceGetPropertyInfo(inDevice, inChannel, isInput, inPropertyID, outSize, outWritable), nil
}

// AudioDriverPlugInDeviceGetPropertyInfo.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInDeviceGetPropertyInfo
func AudioDriverPlugInDeviceGetPropertyInfo(inDevice uint32, inChannel uint32, isInput bool, inPropertyID uint32, outSize *uint32, outWritable *bool) int32 {
	result, callErr := tryAudioDriverPlugInDeviceGetPropertyInfo(inDevice, inChannel, isInput, inPropertyID, outSize, outWritable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDriverPlugInDeviceSetProperty func(inDevice uint32, inWhen unsafe.Pointer, inChannel uint32, isInput bool, inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32
var _audioDriverPlugInDeviceSetPropertyErr error

func tryAudioDriverPlugInDeviceSetProperty(inDevice uint32, inWhen unsafe.Pointer, inChannel uint32, isInput bool, inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) (int32, error) {
	if _audioDriverPlugInDeviceSetProperty == nil {
		return 0, symbolCallError("AudioDriverPlugInDeviceSetProperty", "", _audioDriverPlugInDeviceSetPropertyErr)
	}
	return _audioDriverPlugInDeviceSetProperty(inDevice, inWhen, inChannel, isInput, inPropertyID, inPropertyDataSize, inPropertyData), nil
}

// AudioDriverPlugInDeviceSetProperty.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInDeviceSetProperty
func AudioDriverPlugInDeviceSetProperty(inDevice uint32, inWhen unsafe.Pointer, inChannel uint32, isInput bool, inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioDriverPlugInDeviceSetProperty(inDevice, inWhen, inChannel, isInput, inPropertyID, inPropertyDataSize, inPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDriverPlugInOpen func(inHostInfo *AudioDriverPlugInHostInfo) int32
var _audioDriverPlugInOpenErr error

func tryAudioDriverPlugInOpen(inHostInfo *AudioDriverPlugInHostInfo) (int32, error) {
	if _audioDriverPlugInOpen == nil {
		return 0, symbolCallError("AudioDriverPlugInOpen", "", _audioDriverPlugInOpenErr)
	}
	return _audioDriverPlugInOpen(inHostInfo), nil
}

// AudioDriverPlugInOpen.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInOpen
func AudioDriverPlugInOpen(inHostInfo *AudioDriverPlugInHostInfo) int32 {
	result, callErr := tryAudioDriverPlugInOpen(inHostInfo)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDriverPlugInStreamGetProperty func(inDevice uint32, inIOAudioStream uintptr, inChannel uint32, inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32
var _audioDriverPlugInStreamGetPropertyErr error

func tryAudioDriverPlugInStreamGetProperty(inDevice uint32, inIOAudioStream uintptr, inChannel uint32, inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) (int32, error) {
	if _audioDriverPlugInStreamGetProperty == nil {
		return 0, symbolCallError("AudioDriverPlugInStreamGetProperty", "", _audioDriverPlugInStreamGetPropertyErr)
	}
	return _audioDriverPlugInStreamGetProperty(inDevice, inIOAudioStream, inChannel, inPropertyID, ioPropertyDataSize, outPropertyData), nil
}

// AudioDriverPlugInStreamGetProperty.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInStreamGetProperty
func AudioDriverPlugInStreamGetProperty(inDevice uint32, inIOAudioStream uintptr, inChannel uint32, inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioDriverPlugInStreamGetProperty(inDevice, inIOAudioStream, inChannel, inPropertyID, ioPropertyDataSize, outPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDriverPlugInStreamGetPropertyInfo func(inDevice uint32, inIOAudioStream uintptr, inChannel uint32, inPropertyID uint32, outSize *uint32, outWritable *bool) int32
var _audioDriverPlugInStreamGetPropertyInfoErr error

func tryAudioDriverPlugInStreamGetPropertyInfo(inDevice uint32, inIOAudioStream uintptr, inChannel uint32, inPropertyID uint32, outSize *uint32, outWritable *bool) (int32, error) {
	if _audioDriverPlugInStreamGetPropertyInfo == nil {
		return 0, symbolCallError("AudioDriverPlugInStreamGetPropertyInfo", "", _audioDriverPlugInStreamGetPropertyInfoErr)
	}
	return _audioDriverPlugInStreamGetPropertyInfo(inDevice, inIOAudioStream, inChannel, inPropertyID, outSize, outWritable), nil
}

// AudioDriverPlugInStreamGetPropertyInfo.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInStreamGetPropertyInfo
func AudioDriverPlugInStreamGetPropertyInfo(inDevice uint32, inIOAudioStream uintptr, inChannel uint32, inPropertyID uint32, outSize *uint32, outWritable *bool) int32 {
	result, callErr := tryAudioDriverPlugInStreamGetPropertyInfo(inDevice, inIOAudioStream, inChannel, inPropertyID, outSize, outWritable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioDriverPlugInStreamSetProperty func(inDevice uint32, inIOAudioStream uintptr, inWhen unsafe.Pointer, inChannel uint32, inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32
var _audioDriverPlugInStreamSetPropertyErr error

func tryAudioDriverPlugInStreamSetProperty(inDevice uint32, inIOAudioStream uintptr, inWhen unsafe.Pointer, inChannel uint32, inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) (int32, error) {
	if _audioDriverPlugInStreamSetProperty == nil {
		return 0, symbolCallError("AudioDriverPlugInStreamSetProperty", "", _audioDriverPlugInStreamSetPropertyErr)
	}
	return _audioDriverPlugInStreamSetProperty(inDevice, inIOAudioStream, inWhen, inChannel, inPropertyID, inPropertyDataSize, inPropertyData), nil
}

// AudioDriverPlugInStreamSetProperty.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioDriverPlugInStreamSetProperty
func AudioDriverPlugInStreamSetProperty(inDevice uint32, inIOAudioStream uintptr, inWhen unsafe.Pointer, inChannel uint32, inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioDriverPlugInStreamSetProperty(inDevice, inIOAudioStream, inWhen, inChannel, inPropertyID, inPropertyDataSize, inPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioGetCurrentHostTime func() uint64
var _audioGetCurrentHostTimeErr error

func tryAudioGetCurrentHostTime() (uint64, error) {
	if _audioGetCurrentHostTime == nil {
		return 0, symbolCallError("AudioGetCurrentHostTime", "10.0", _audioGetCurrentHostTimeErr)
	}
	return _audioGetCurrentHostTime(), nil
}

// AudioGetCurrentHostTime.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioGetCurrentHostTime()
func AudioGetCurrentHostTime() uint64 {
	result, callErr := tryAudioGetCurrentHostTime()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioGetHostClockFrequency func() float64
var _audioGetHostClockFrequencyErr error

func tryAudioGetHostClockFrequency() (float64, error) {
	if _audioGetHostClockFrequency == nil {
		return 0.0, symbolCallError("AudioGetHostClockFrequency", "10.0", _audioGetHostClockFrequencyErr)
	}
	return _audioGetHostClockFrequency(), nil
}

// AudioGetHostClockFrequency.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioGetHostClockFrequency()
func AudioGetHostClockFrequency() float64 {
	result, callErr := tryAudioGetHostClockFrequency()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioGetHostClockMinimumTimeDelta func() uint32
var _audioGetHostClockMinimumTimeDeltaErr error

func tryAudioGetHostClockMinimumTimeDelta() (uint32, error) {
	if _audioGetHostClockMinimumTimeDelta == nil {
		return 0, symbolCallError("AudioGetHostClockMinimumTimeDelta", "10.0", _audioGetHostClockMinimumTimeDeltaErr)
	}
	return _audioGetHostClockMinimumTimeDelta(), nil
}

// AudioGetHostClockMinimumTimeDelta.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioGetHostClockMinimumTimeDelta()
func AudioGetHostClockMinimumTimeDelta() uint32 {
	result, callErr := tryAudioGetHostClockMinimumTimeDelta()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioHardwareAddPropertyListener func(inPropertyID uint32, inProc AudioHardwarePropertyListenerProc, inClientData unsafe.Pointer) int32
var _audioHardwareAddPropertyListenerErr error

func tryAudioHardwareAddPropertyListener(inPropertyID uint32, inProc AudioHardwarePropertyListenerProc, inClientData unsafe.Pointer) (int32, error) {
	if _audioHardwareAddPropertyListener == nil {
		return 0, symbolCallError("AudioHardwareAddPropertyListener", "10.0", _audioHardwareAddPropertyListenerErr)
	}
	return _audioHardwareAddPropertyListener(inPropertyID, inProc, inClientData), nil
}

// AudioHardwareAddPropertyListener.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareAddPropertyListener
func AudioHardwareAddPropertyListener(inPropertyID uint32, inProc AudioHardwarePropertyListenerProc, inClientData unsafe.Pointer) int32 {
	result, callErr := tryAudioHardwareAddPropertyListener(inPropertyID, inProc, inClientData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioHardwareAddRunLoopSource func(inRunLoopSource corefoundation.CFRunLoopSourceRef) int32
var _audioHardwareAddRunLoopSourceErr error

func tryAudioHardwareAddRunLoopSource(inRunLoopSource corefoundation.CFRunLoopSourceRef) (int32, error) {
	if _audioHardwareAddRunLoopSource == nil {
		return 0, symbolCallError("AudioHardwareAddRunLoopSource", "10.3", _audioHardwareAddRunLoopSourceErr)
	}
	return _audioHardwareAddRunLoopSource(inRunLoopSource), nil
}

// AudioHardwareAddRunLoopSource.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareAddRunLoopSource
func AudioHardwareAddRunLoopSource(inRunLoopSource corefoundation.CFRunLoopSourceRef) int32 {
	result, callErr := tryAudioHardwareAddRunLoopSource(inRunLoopSource)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioHardwareCreateAggregateDevice func(inDescription corefoundation.CFDictionaryRef, outDeviceID *uint32) int32
var _audioHardwareCreateAggregateDeviceErr error

func tryAudioHardwareCreateAggregateDevice(inDescription corefoundation.CFDictionaryRef, outDeviceID *uint32) (int32, error) {
	if _audioHardwareCreateAggregateDevice == nil {
		return 0, symbolCallError("AudioHardwareCreateAggregateDevice", "10.9", _audioHardwareCreateAggregateDeviceErr)
	}
	return _audioHardwareCreateAggregateDevice(inDescription, outDeviceID), nil
}

// AudioHardwareCreateAggregateDevice.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareCreateAggregateDevice(_:_:)
func AudioHardwareCreateAggregateDevice(inDescription corefoundation.CFDictionaryRef, outDeviceID *uint32) int32 {
	result, callErr := tryAudioHardwareCreateAggregateDevice(inDescription, outDeviceID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioHardwareCreateProcessTap func(inDescription *CATapDescription, outTapID *uint32) int32
var _audioHardwareCreateProcessTapErr error

func tryAudioHardwareCreateProcessTap(inDescription *CATapDescription, outTapID *uint32) (int32, error) {
	if _audioHardwareCreateProcessTap == nil {
		return 0, symbolCallError("AudioHardwareCreateProcessTap", "14.2", _audioHardwareCreateProcessTapErr)
	}
	return _audioHardwareCreateProcessTap(inDescription, outTapID), nil
}

// AudioHardwareCreateProcessTap.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareCreateProcessTap(_:_:)
func AudioHardwareCreateProcessTap(inDescription *CATapDescription, outTapID *uint32) int32 {
	result, callErr := tryAudioHardwareCreateProcessTap(inDescription, outTapID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioHardwareDestroyAggregateDevice func(inDeviceID uint32) int32
var _audioHardwareDestroyAggregateDeviceErr error

func tryAudioHardwareDestroyAggregateDevice(inDeviceID uint32) (int32, error) {
	if _audioHardwareDestroyAggregateDevice == nil {
		return 0, symbolCallError("AudioHardwareDestroyAggregateDevice", "10.9", _audioHardwareDestroyAggregateDeviceErr)
	}
	return _audioHardwareDestroyAggregateDevice(inDeviceID), nil
}

// AudioHardwareDestroyAggregateDevice.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareDestroyAggregateDevice(_:)
func AudioHardwareDestroyAggregateDevice(inDeviceID uint32) int32 {
	result, callErr := tryAudioHardwareDestroyAggregateDevice(inDeviceID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioHardwareDestroyProcessTap func(inTapID uint32) int32
var _audioHardwareDestroyProcessTapErr error

func tryAudioHardwareDestroyProcessTap(inTapID uint32) (int32, error) {
	if _audioHardwareDestroyProcessTap == nil {
		return 0, symbolCallError("AudioHardwareDestroyProcessTap", "14.2", _audioHardwareDestroyProcessTapErr)
	}
	return _audioHardwareDestroyProcessTap(inTapID), nil
}

// AudioHardwareDestroyProcessTap.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareDestroyProcessTap(_:)
func AudioHardwareDestroyProcessTap(inTapID uint32) int32 {
	result, callErr := tryAudioHardwareDestroyProcessTap(inTapID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioHardwareGetProperty func(inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32
var _audioHardwareGetPropertyErr error

func tryAudioHardwareGetProperty(inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) (int32, error) {
	if _audioHardwareGetProperty == nil {
		return 0, symbolCallError("AudioHardwareGetProperty", "10.0", _audioHardwareGetPropertyErr)
	}
	return _audioHardwareGetProperty(inPropertyID, ioPropertyDataSize, outPropertyData), nil
}

// AudioHardwareGetProperty.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareGetProperty
func AudioHardwareGetProperty(inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioHardwareGetProperty(inPropertyID, ioPropertyDataSize, outPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioHardwareGetPropertyInfo func(inPropertyID uint32, outSize *uint32, outWritable *bool) int32
var _audioHardwareGetPropertyInfoErr error

func tryAudioHardwareGetPropertyInfo(inPropertyID uint32, outSize *uint32, outWritable *bool) (int32, error) {
	if _audioHardwareGetPropertyInfo == nil {
		return 0, symbolCallError("AudioHardwareGetPropertyInfo", "10.0", _audioHardwareGetPropertyInfoErr)
	}
	return _audioHardwareGetPropertyInfo(inPropertyID, outSize, outWritable), nil
}

// AudioHardwareGetPropertyInfo.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareGetPropertyInfo
func AudioHardwareGetPropertyInfo(inPropertyID uint32, outSize *uint32, outWritable *bool) int32 {
	result, callErr := tryAudioHardwareGetPropertyInfo(inPropertyID, outSize, outWritable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioHardwareRemovePropertyListener func(inPropertyID uint32, inProc AudioHardwarePropertyListenerProc) int32
var _audioHardwareRemovePropertyListenerErr error

func tryAudioHardwareRemovePropertyListener(inPropertyID uint32, inProc AudioHardwarePropertyListenerProc) (int32, error) {
	if _audioHardwareRemovePropertyListener == nil {
		return 0, symbolCallError("AudioHardwareRemovePropertyListener", "10.0", _audioHardwareRemovePropertyListenerErr)
	}
	return _audioHardwareRemovePropertyListener(inPropertyID, inProc), nil
}

// AudioHardwareRemovePropertyListener.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareRemovePropertyListener
func AudioHardwareRemovePropertyListener(inPropertyID uint32, inProc AudioHardwarePropertyListenerProc) int32 {
	result, callErr := tryAudioHardwareRemovePropertyListener(inPropertyID, inProc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioHardwareRemoveRunLoopSource func(inRunLoopSource corefoundation.CFRunLoopSourceRef) int32
var _audioHardwareRemoveRunLoopSourceErr error

func tryAudioHardwareRemoveRunLoopSource(inRunLoopSource corefoundation.CFRunLoopSourceRef) (int32, error) {
	if _audioHardwareRemoveRunLoopSource == nil {
		return 0, symbolCallError("AudioHardwareRemoveRunLoopSource", "10.3", _audioHardwareRemoveRunLoopSourceErr)
	}
	return _audioHardwareRemoveRunLoopSource(inRunLoopSource), nil
}

// AudioHardwareRemoveRunLoopSource.
//
// Deprecated: Deprecated since macOS 10.7.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareRemoveRunLoopSource
func AudioHardwareRemoveRunLoopSource(inRunLoopSource corefoundation.CFRunLoopSourceRef) int32 {
	result, callErr := tryAudioHardwareRemoveRunLoopSource(inRunLoopSource)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioHardwareSetProperty func(inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32
var _audioHardwareSetPropertyErr error

func tryAudioHardwareSetProperty(inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) (int32, error) {
	if _audioHardwareSetProperty == nil {
		return 0, symbolCallError("AudioHardwareSetProperty", "10.0", _audioHardwareSetPropertyErr)
	}
	return _audioHardwareSetProperty(inPropertyID, inPropertyDataSize, inPropertyData), nil
}

// AudioHardwareSetProperty.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareSetProperty
func AudioHardwareSetProperty(inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioHardwareSetProperty(inPropertyID, inPropertyDataSize, inPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioHardwareUnload func() int32
var _audioHardwareUnloadErr error

func tryAudioHardwareUnload() (int32, error) {
	if _audioHardwareUnload == nil {
		return 0, symbolCallError("AudioHardwareUnload", "10.1", _audioHardwareUnloadErr)
	}
	return _audioHardwareUnload(), nil
}

// AudioHardwareUnload.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioHardwareUnload()
func AudioHardwareUnload() int32 {
	result, callErr := tryAudioHardwareUnload()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioObjectAddPropertyListener func(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inListener AudioObjectPropertyListenerProc, inClientData unsafe.Pointer) int32
var _audioObjectAddPropertyListenerErr error

func tryAudioObjectAddPropertyListener(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inListener AudioObjectPropertyListenerProc, inClientData unsafe.Pointer) (int32, error) {
	if _audioObjectAddPropertyListener == nil {
		return 0, symbolCallError("AudioObjectAddPropertyListener", "10.4", _audioObjectAddPropertyListenerErr)
	}
	return _audioObjectAddPropertyListener(inObjectID, inAddress, inListener, inClientData), nil
}

// AudioObjectAddPropertyListener.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectAddPropertyListener(_:_:_:_:)
func AudioObjectAddPropertyListener(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inListener AudioObjectPropertyListenerProc, inClientData unsafe.Pointer) int32 {
	result, callErr := tryAudioObjectAddPropertyListener(inObjectID, inAddress, inListener, inClientData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioObjectAddPropertyListenerBlock func(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inDispatchQueue uintptr, inListener unsafe.Pointer) int32
var _audioObjectAddPropertyListenerBlockErr error

func tryAudioObjectAddPropertyListenerBlock(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inDispatchQueue dispatch.Queue, inListener AudioObjectPropertyListenerBlock) (int32, error) {
	if _audioObjectAddPropertyListenerBlock == nil {
		return 0, symbolCallError("AudioObjectAddPropertyListenerBlock", "10.7", _audioObjectAddPropertyListenerBlockErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 uint32, arg1 *AudioObjectPropertyAddress) { inListener(arg0, arg1) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _audioObjectAddPropertyListenerBlock(inObjectID, inAddress, uintptr(inDispatchQueue.Handle()), _block0), nil
}

// AudioObjectAddPropertyListenerBlock.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectAddPropertyListenerBlock(_:_:_:_:)
func AudioObjectAddPropertyListenerBlock(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inDispatchQueue dispatch.Queue, inListener AudioObjectPropertyListenerBlock) int32 {
	result, callErr := tryAudioObjectAddPropertyListenerBlock(inObjectID, inAddress, inDispatchQueue, inListener)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioObjectGetPropertyData func(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inQualifierDataSize uint32, inQualifierData unsafe.Pointer, ioDataSize *uint32, outData unsafe.Pointer) int32
var _audioObjectGetPropertyDataErr error

func tryAudioObjectGetPropertyData(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inQualifierDataSize uint32, inQualifierData unsafe.Pointer, ioDataSize *uint32, outData unsafe.Pointer) (int32, error) {
	if _audioObjectGetPropertyData == nil {
		return 0, symbolCallError("AudioObjectGetPropertyData", "10.4", _audioObjectGetPropertyDataErr)
	}
	return _audioObjectGetPropertyData(inObjectID, inAddress, inQualifierDataSize, inQualifierData, ioDataSize, outData), nil
}

// AudioObjectGetPropertyData.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectGetPropertyData(_:_:_:_:_:_:)
func AudioObjectGetPropertyData(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inQualifierDataSize uint32, inQualifierData unsafe.Pointer, ioDataSize *uint32, outData unsafe.Pointer) int32 {
	result, callErr := tryAudioObjectGetPropertyData(inObjectID, inAddress, inQualifierDataSize, inQualifierData, ioDataSize, outData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioObjectGetPropertyDataSize func(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inQualifierDataSize uint32, inQualifierData unsafe.Pointer, outDataSize *uint32) int32
var _audioObjectGetPropertyDataSizeErr error

func tryAudioObjectGetPropertyDataSize(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inQualifierDataSize uint32, inQualifierData unsafe.Pointer, outDataSize *uint32) (int32, error) {
	if _audioObjectGetPropertyDataSize == nil {
		return 0, symbolCallError("AudioObjectGetPropertyDataSize", "10.4", _audioObjectGetPropertyDataSizeErr)
	}
	return _audioObjectGetPropertyDataSize(inObjectID, inAddress, inQualifierDataSize, inQualifierData, outDataSize), nil
}

// AudioObjectGetPropertyDataSize.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectGetPropertyDataSize(_:_:_:_:_:)
func AudioObjectGetPropertyDataSize(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inQualifierDataSize uint32, inQualifierData unsafe.Pointer, outDataSize *uint32) int32 {
	result, callErr := tryAudioObjectGetPropertyDataSize(inObjectID, inAddress, inQualifierDataSize, inQualifierData, outDataSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioObjectHasProperty func(inObjectID uint32, inAddress *AudioObjectPropertyAddress) bool
var _audioObjectHasPropertyErr error

func tryAudioObjectHasProperty(inObjectID uint32, inAddress *AudioObjectPropertyAddress) (bool, error) {
	if _audioObjectHasProperty == nil {
		return false, symbolCallError("AudioObjectHasProperty", "10.4", _audioObjectHasPropertyErr)
	}
	return _audioObjectHasProperty(inObjectID, inAddress), nil
}

// AudioObjectHasProperty.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectHasProperty(_:_:)
func AudioObjectHasProperty(inObjectID uint32, inAddress *AudioObjectPropertyAddress) bool {
	result, callErr := tryAudioObjectHasProperty(inObjectID, inAddress)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioObjectIsPropertySettable func(inObjectID uint32, inAddress *AudioObjectPropertyAddress, outIsSettable *bool) int32
var _audioObjectIsPropertySettableErr error

func tryAudioObjectIsPropertySettable(inObjectID uint32, inAddress *AudioObjectPropertyAddress, outIsSettable *bool) (int32, error) {
	if _audioObjectIsPropertySettable == nil {
		return 0, symbolCallError("AudioObjectIsPropertySettable", "10.4", _audioObjectIsPropertySettableErr)
	}
	return _audioObjectIsPropertySettable(inObjectID, inAddress, outIsSettable), nil
}

// AudioObjectIsPropertySettable.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectIsPropertySettable(_:_:_:)
func AudioObjectIsPropertySettable(inObjectID uint32, inAddress *AudioObjectPropertyAddress, outIsSettable *bool) int32 {
	result, callErr := tryAudioObjectIsPropertySettable(inObjectID, inAddress, outIsSettable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioObjectRemovePropertyListener func(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inListener AudioObjectPropertyListenerProc, inClientData unsafe.Pointer) int32
var _audioObjectRemovePropertyListenerErr error

func tryAudioObjectRemovePropertyListener(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inListener AudioObjectPropertyListenerProc, inClientData unsafe.Pointer) (int32, error) {
	if _audioObjectRemovePropertyListener == nil {
		return 0, symbolCallError("AudioObjectRemovePropertyListener", "10.4", _audioObjectRemovePropertyListenerErr)
	}
	return _audioObjectRemovePropertyListener(inObjectID, inAddress, inListener, inClientData), nil
}

// AudioObjectRemovePropertyListener.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectRemovePropertyListener(_:_:_:_:)
func AudioObjectRemovePropertyListener(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inListener AudioObjectPropertyListenerProc, inClientData unsafe.Pointer) int32 {
	result, callErr := tryAudioObjectRemovePropertyListener(inObjectID, inAddress, inListener, inClientData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioObjectRemovePropertyListenerBlock func(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inDispatchQueue uintptr, inListener unsafe.Pointer) int32
var _audioObjectRemovePropertyListenerBlockErr error

func tryAudioObjectRemovePropertyListenerBlock(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inDispatchQueue dispatch.Queue, inListener AudioObjectPropertyListenerBlock) (int32, error) {
	if _audioObjectRemovePropertyListenerBlock == nil {
		return 0, symbolCallError("AudioObjectRemovePropertyListenerBlock", "10.7", _audioObjectRemovePropertyListenerBlockErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 uint32, arg1 *AudioObjectPropertyAddress) { inListener(arg0, arg1) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _audioObjectRemovePropertyListenerBlock(inObjectID, inAddress, uintptr(inDispatchQueue.Handle()), _block0), nil
}

// AudioObjectRemovePropertyListenerBlock.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectRemovePropertyListenerBlock(_:_:_:_:)
func AudioObjectRemovePropertyListenerBlock(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inDispatchQueue dispatch.Queue, inListener AudioObjectPropertyListenerBlock) int32 {
	result, callErr := tryAudioObjectRemovePropertyListenerBlock(inObjectID, inAddress, inDispatchQueue, inListener)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioObjectSetPropertyData func(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inQualifierDataSize uint32, inQualifierData unsafe.Pointer, inDataSize uint32, inData unsafe.Pointer) int32
var _audioObjectSetPropertyDataErr error

func tryAudioObjectSetPropertyData(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inQualifierDataSize uint32, inQualifierData unsafe.Pointer, inDataSize uint32, inData unsafe.Pointer) (int32, error) {
	if _audioObjectSetPropertyData == nil {
		return 0, symbolCallError("AudioObjectSetPropertyData", "10.4", _audioObjectSetPropertyDataErr)
	}
	return _audioObjectSetPropertyData(inObjectID, inAddress, inQualifierDataSize, inQualifierData, inDataSize, inData), nil
}

// AudioObjectSetPropertyData.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectSetPropertyData(_:_:_:_:_:_:)
func AudioObjectSetPropertyData(inObjectID uint32, inAddress *AudioObjectPropertyAddress, inQualifierDataSize uint32, inQualifierData unsafe.Pointer, inDataSize uint32, inData unsafe.Pointer) int32 {
	result, callErr := tryAudioObjectSetPropertyData(inObjectID, inAddress, inQualifierDataSize, inQualifierData, inDataSize, inData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioObjectShow func(inObjectID uint32)
var _audioObjectShowErr error

func tryAudioObjectShow(inObjectID uint32) error {
	if _audioObjectShow == nil {
		return symbolCallError("AudioObjectShow", "10.4", _audioObjectShowErr)
	}
	_audioObjectShow(inObjectID)
	return nil
}

// AudioObjectShow.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioObjectShow(_:)
func AudioObjectShow(inObjectID uint32) {
	if callErr := tryAudioObjectShow(inObjectID); callErr != nil {
		panic(callErr)
	}
}

var _audioStreamAddPropertyListener func(inStream uint32, inChannel uint32, inPropertyID uint32, inProc AudioStreamPropertyListenerProc, inClientData unsafe.Pointer) int32
var _audioStreamAddPropertyListenerErr error

func tryAudioStreamAddPropertyListener(inStream uint32, inChannel uint32, inPropertyID uint32, inProc AudioStreamPropertyListenerProc, inClientData unsafe.Pointer) (int32, error) {
	if _audioStreamAddPropertyListener == nil {
		return 0, symbolCallError("AudioStreamAddPropertyListener", "10.1", _audioStreamAddPropertyListenerErr)
	}
	return _audioStreamAddPropertyListener(inStream, inChannel, inPropertyID, inProc, inClientData), nil
}

// AudioStreamAddPropertyListener.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioStreamAddPropertyListener
func AudioStreamAddPropertyListener(inStream uint32, inChannel uint32, inPropertyID uint32, inProc AudioStreamPropertyListenerProc, inClientData unsafe.Pointer) int32 {
	result, callErr := tryAudioStreamAddPropertyListener(inStream, inChannel, inPropertyID, inProc, inClientData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioStreamGetProperty func(inStream uint32, inChannel uint32, inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32
var _audioStreamGetPropertyErr error

func tryAudioStreamGetProperty(inStream uint32, inChannel uint32, inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) (int32, error) {
	if _audioStreamGetProperty == nil {
		return 0, symbolCallError("AudioStreamGetProperty", "10.1", _audioStreamGetPropertyErr)
	}
	return _audioStreamGetProperty(inStream, inChannel, inPropertyID, ioPropertyDataSize, outPropertyData), nil
}

// AudioStreamGetProperty.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioStreamGetProperty
func AudioStreamGetProperty(inStream uint32, inChannel uint32, inPropertyID uint32, ioPropertyDataSize *uint32, outPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioStreamGetProperty(inStream, inChannel, inPropertyID, ioPropertyDataSize, outPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioStreamGetPropertyInfo func(inStream uint32, inChannel uint32, inPropertyID uint32, outSize *uint32, outWritable *bool) int32
var _audioStreamGetPropertyInfoErr error

func tryAudioStreamGetPropertyInfo(inStream uint32, inChannel uint32, inPropertyID uint32, outSize *uint32, outWritable *bool) (int32, error) {
	if _audioStreamGetPropertyInfo == nil {
		return 0, symbolCallError("AudioStreamGetPropertyInfo", "10.1", _audioStreamGetPropertyInfoErr)
	}
	return _audioStreamGetPropertyInfo(inStream, inChannel, inPropertyID, outSize, outWritable), nil
}

// AudioStreamGetPropertyInfo.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioStreamGetPropertyInfo
func AudioStreamGetPropertyInfo(inStream uint32, inChannel uint32, inPropertyID uint32, outSize *uint32, outWritable *bool) int32 {
	result, callErr := tryAudioStreamGetPropertyInfo(inStream, inChannel, inPropertyID, outSize, outWritable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioStreamRemovePropertyListener func(inStream uint32, inChannel uint32, inPropertyID uint32, inProc AudioStreamPropertyListenerProc) int32
var _audioStreamRemovePropertyListenerErr error

func tryAudioStreamRemovePropertyListener(inStream uint32, inChannel uint32, inPropertyID uint32, inProc AudioStreamPropertyListenerProc) (int32, error) {
	if _audioStreamRemovePropertyListener == nil {
		return 0, symbolCallError("AudioStreamRemovePropertyListener", "10.1", _audioStreamRemovePropertyListenerErr)
	}
	return _audioStreamRemovePropertyListener(inStream, inChannel, inPropertyID, inProc), nil
}

// AudioStreamRemovePropertyListener.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioStreamRemovePropertyListener
func AudioStreamRemovePropertyListener(inStream uint32, inChannel uint32, inPropertyID uint32, inProc AudioStreamPropertyListenerProc) int32 {
	result, callErr := tryAudioStreamRemovePropertyListener(inStream, inChannel, inPropertyID, inProc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _audioStreamSetProperty func(inStream uint32, inWhen unsafe.Pointer, inChannel uint32, inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32
var _audioStreamSetPropertyErr error

func tryAudioStreamSetProperty(inStream uint32, inWhen unsafe.Pointer, inChannel uint32, inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) (int32, error) {
	if _audioStreamSetProperty == nil {
		return 0, symbolCallError("AudioStreamSetProperty", "10.1", _audioStreamSetPropertyErr)
	}
	return _audioStreamSetProperty(inStream, inWhen, inChannel, inPropertyID, inPropertyDataSize, inPropertyData), nil
}

// AudioStreamSetProperty.
//
// Deprecated: Deprecated since macOS 10.6.
//
// See: https://developer.apple.com/documentation/CoreAudio/AudioStreamSetProperty
func AudioStreamSetProperty(inStream uint32, inWhen unsafe.Pointer, inChannel uint32, inPropertyID uint32, inPropertyDataSize uint32, inPropertyData unsafe.Pointer) int32 {
	result, callErr := tryAudioStreamSetProperty(inStream, inWhen, inChannel, inPropertyID, inPropertyDataSize, inPropertyData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_audioConvertHostTimeToNanos, &_audioConvertHostTimeToNanosErr, frameworkHandle, "AudioConvertHostTimeToNanos", "10.0")
	registerFunc(&_audioConvertNanosToHostTime, &_audioConvertNanosToHostTimeErr, frameworkHandle, "AudioConvertNanosToHostTime", "10.0")
	registerFunc(&_audioDeviceAddIOProc, &_audioDeviceAddIOProcErr, frameworkHandle, "AudioDeviceAddIOProc", "10.0")
	registerFunc(&_audioDeviceAddPropertyListener, &_audioDeviceAddPropertyListenerErr, frameworkHandle, "AudioDeviceAddPropertyListener", "10.0")
	registerFunc(&_audioDeviceCreateIOProcID, &_audioDeviceCreateIOProcIDErr, frameworkHandle, "AudioDeviceCreateIOProcID", "10.5")
	registerFunc(&_audioDeviceCreateIOProcIDWithBlock, &_audioDeviceCreateIOProcIDWithBlockErr, frameworkHandle, "AudioDeviceCreateIOProcIDWithBlock", "10.7")
	registerFunc(&_audioDeviceDestroyIOProcID, &_audioDeviceDestroyIOProcIDErr, frameworkHandle, "AudioDeviceDestroyIOProcID", "10.5")
	registerFunc(&_audioDeviceGetCurrentTime, &_audioDeviceGetCurrentTimeErr, frameworkHandle, "AudioDeviceGetCurrentTime", "10.0")
	registerFunc(&_audioDeviceGetNearestStartTime, &_audioDeviceGetNearestStartTimeErr, frameworkHandle, "AudioDeviceGetNearestStartTime", "10.3")
	registerFunc(&_audioDeviceGetProperty, &_audioDeviceGetPropertyErr, frameworkHandle, "AudioDeviceGetProperty", "10.0")
	registerFunc(&_audioDeviceGetPropertyInfo, &_audioDeviceGetPropertyInfoErr, frameworkHandle, "AudioDeviceGetPropertyInfo", "10.0")
	registerFunc(&_audioDeviceRead, &_audioDeviceReadErr, frameworkHandle, "AudioDeviceRead", "10.1")
	registerFunc(&_audioDeviceRemoveIOProc, &_audioDeviceRemoveIOProcErr, frameworkHandle, "AudioDeviceRemoveIOProc", "10.0")
	registerFunc(&_audioDeviceRemovePropertyListener, &_audioDeviceRemovePropertyListenerErr, frameworkHandle, "AudioDeviceRemovePropertyListener", "10.0")
	registerFunc(&_audioDeviceSetProperty, &_audioDeviceSetPropertyErr, frameworkHandle, "AudioDeviceSetProperty", "10.0")
	registerFunc(&_audioDeviceStart, &_audioDeviceStartErr, frameworkHandle, "AudioDeviceStart", "10.0")
	registerFunc(&_audioDeviceStartAtTime, &_audioDeviceStartAtTimeErr, frameworkHandle, "AudioDeviceStartAtTime", "10.3")
	registerFunc(&_audioDeviceStop, &_audioDeviceStopErr, frameworkHandle, "AudioDeviceStop", "10.0")
	registerFunc(&_audioDeviceTranslateTime, &_audioDeviceTranslateTimeErr, frameworkHandle, "AudioDeviceTranslateTime", "10.0")
	registerFunc(&_audioDriverPlugInClose, &_audioDriverPlugInCloseErr, frameworkHandle, "AudioDriverPlugInClose", "")
	registerFunc(&_audioDriverPlugInDeviceGetProperty, &_audioDriverPlugInDeviceGetPropertyErr, frameworkHandle, "AudioDriverPlugInDeviceGetProperty", "")
	registerFunc(&_audioDriverPlugInDeviceGetPropertyInfo, &_audioDriverPlugInDeviceGetPropertyInfoErr, frameworkHandle, "AudioDriverPlugInDeviceGetPropertyInfo", "")
	registerFunc(&_audioDriverPlugInDeviceSetProperty, &_audioDriverPlugInDeviceSetPropertyErr, frameworkHandle, "AudioDriverPlugInDeviceSetProperty", "")
	registerFunc(&_audioDriverPlugInOpen, &_audioDriverPlugInOpenErr, frameworkHandle, "AudioDriverPlugInOpen", "")
	registerFunc(&_audioDriverPlugInStreamGetProperty, &_audioDriverPlugInStreamGetPropertyErr, frameworkHandle, "AudioDriverPlugInStreamGetProperty", "")
	registerFunc(&_audioDriverPlugInStreamGetPropertyInfo, &_audioDriverPlugInStreamGetPropertyInfoErr, frameworkHandle, "AudioDriverPlugInStreamGetPropertyInfo", "")
	registerFunc(&_audioDriverPlugInStreamSetProperty, &_audioDriverPlugInStreamSetPropertyErr, frameworkHandle, "AudioDriverPlugInStreamSetProperty", "")
	registerFunc(&_audioGetCurrentHostTime, &_audioGetCurrentHostTimeErr, frameworkHandle, "AudioGetCurrentHostTime", "10.0")
	registerFunc(&_audioGetHostClockFrequency, &_audioGetHostClockFrequencyErr, frameworkHandle, "AudioGetHostClockFrequency", "10.0")
	registerFunc(&_audioGetHostClockMinimumTimeDelta, &_audioGetHostClockMinimumTimeDeltaErr, frameworkHandle, "AudioGetHostClockMinimumTimeDelta", "10.0")
	registerFunc(&_audioHardwareAddPropertyListener, &_audioHardwareAddPropertyListenerErr, frameworkHandle, "AudioHardwareAddPropertyListener", "10.0")
	registerFunc(&_audioHardwareAddRunLoopSource, &_audioHardwareAddRunLoopSourceErr, frameworkHandle, "AudioHardwareAddRunLoopSource", "10.3")
	registerFunc(&_audioHardwareCreateAggregateDevice, &_audioHardwareCreateAggregateDeviceErr, frameworkHandle, "AudioHardwareCreateAggregateDevice", "10.9")
	registerFunc(&_audioHardwareCreateProcessTap, &_audioHardwareCreateProcessTapErr, frameworkHandle, "AudioHardwareCreateProcessTap", "14.2")
	registerFunc(&_audioHardwareDestroyAggregateDevice, &_audioHardwareDestroyAggregateDeviceErr, frameworkHandle, "AudioHardwareDestroyAggregateDevice", "10.9")
	registerFunc(&_audioHardwareDestroyProcessTap, &_audioHardwareDestroyProcessTapErr, frameworkHandle, "AudioHardwareDestroyProcessTap", "14.2")
	registerFunc(&_audioHardwareGetProperty, &_audioHardwareGetPropertyErr, frameworkHandle, "AudioHardwareGetProperty", "10.0")
	registerFunc(&_audioHardwareGetPropertyInfo, &_audioHardwareGetPropertyInfoErr, frameworkHandle, "AudioHardwareGetPropertyInfo", "10.0")
	registerFunc(&_audioHardwareRemovePropertyListener, &_audioHardwareRemovePropertyListenerErr, frameworkHandle, "AudioHardwareRemovePropertyListener", "10.0")
	registerFunc(&_audioHardwareRemoveRunLoopSource, &_audioHardwareRemoveRunLoopSourceErr, frameworkHandle, "AudioHardwareRemoveRunLoopSource", "10.3")
	registerFunc(&_audioHardwareSetProperty, &_audioHardwareSetPropertyErr, frameworkHandle, "AudioHardwareSetProperty", "10.0")
	registerFunc(&_audioHardwareUnload, &_audioHardwareUnloadErr, frameworkHandle, "AudioHardwareUnload", "10.1")
	registerFunc(&_audioObjectAddPropertyListener, &_audioObjectAddPropertyListenerErr, frameworkHandle, "AudioObjectAddPropertyListener", "10.4")
	registerFunc(&_audioObjectAddPropertyListenerBlock, &_audioObjectAddPropertyListenerBlockErr, frameworkHandle, "AudioObjectAddPropertyListenerBlock", "10.7")
	registerFunc(&_audioObjectGetPropertyData, &_audioObjectGetPropertyDataErr, frameworkHandle, "AudioObjectGetPropertyData", "10.4")
	registerFunc(&_audioObjectGetPropertyDataSize, &_audioObjectGetPropertyDataSizeErr, frameworkHandle, "AudioObjectGetPropertyDataSize", "10.4")
	registerFunc(&_audioObjectHasProperty, &_audioObjectHasPropertyErr, frameworkHandle, "AudioObjectHasProperty", "10.4")
	registerFunc(&_audioObjectIsPropertySettable, &_audioObjectIsPropertySettableErr, frameworkHandle, "AudioObjectIsPropertySettable", "10.4")
	registerFunc(&_audioObjectRemovePropertyListener, &_audioObjectRemovePropertyListenerErr, frameworkHandle, "AudioObjectRemovePropertyListener", "10.4")
	registerFunc(&_audioObjectRemovePropertyListenerBlock, &_audioObjectRemovePropertyListenerBlockErr, frameworkHandle, "AudioObjectRemovePropertyListenerBlock", "10.7")
	registerFunc(&_audioObjectSetPropertyData, &_audioObjectSetPropertyDataErr, frameworkHandle, "AudioObjectSetPropertyData", "10.4")
	registerFunc(&_audioObjectShow, &_audioObjectShowErr, frameworkHandle, "AudioObjectShow", "10.4")
	registerFunc(&_audioStreamAddPropertyListener, &_audioStreamAddPropertyListenerErr, frameworkHandle, "AudioStreamAddPropertyListener", "10.1")
	registerFunc(&_audioStreamGetProperty, &_audioStreamGetPropertyErr, frameworkHandle, "AudioStreamGetProperty", "10.1")
	registerFunc(&_audioStreamGetPropertyInfo, &_audioStreamGetPropertyInfoErr, frameworkHandle, "AudioStreamGetPropertyInfo", "10.1")
	registerFunc(&_audioStreamRemovePropertyListener, &_audioStreamRemovePropertyListenerErr, frameworkHandle, "AudioStreamRemovePropertyListener", "10.1")
	registerFunc(&_audioStreamSetProperty, &_audioStreamSetPropertyErr, frameworkHandle, "AudioStreamSetProperty", "10.1")
}
