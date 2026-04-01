// Code generated from Apple documentation for CoreMedia. DO NOT EDIT.

package coremedia

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
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
		return fmt.Sprintf("CoreMedia: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("CoreMedia: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("CoreMedia: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("CoreMedia: register symbol %s: %v", name, r)
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

var _cMAudioDeviceClockCreate func(allocator corefoundation.CFAllocatorRef, deviceUID corefoundation.CFStringRef, clockOut *uintptr) int32
var _cMAudioDeviceClockCreateErr error

func tryCMAudioDeviceClockCreate(allocator corefoundation.CFAllocatorRef, deviceUID corefoundation.CFStringRef, clockOut *uintptr) (int32, error) {
	if _cMAudioDeviceClockCreate == nil {
		return 0, symbolCallError("CMAudioDeviceClockCreate", "10.8", _cMAudioDeviceClockCreateErr)
	}
	return _cMAudioDeviceClockCreate(allocator, deviceUID, clockOut), nil
}

// CMAudioDeviceClockCreate creates a clock that tracks playback through a Core Audio device with the specified unique identifier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioDeviceClockCreate(allocator:deviceUID:clockOut:)
func CMAudioDeviceClockCreate(allocator corefoundation.CFAllocatorRef, deviceUID corefoundation.CFStringRef, clockOut *uintptr) int32 {
	result, callErr := tryCMAudioDeviceClockCreate(allocator, deviceUID, clockOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMAudioDeviceClockCreateFromAudioDeviceID func(allocator corefoundation.CFAllocatorRef, deviceID uint32, clockOut *uintptr) int32
var _cMAudioDeviceClockCreateFromAudioDeviceIDErr error

func tryCMAudioDeviceClockCreateFromAudioDeviceID(allocator corefoundation.CFAllocatorRef, deviceID uint32, clockOut *uintptr) (int32, error) {
	if _cMAudioDeviceClockCreateFromAudioDeviceID == nil {
		return 0, symbolCallError("CMAudioDeviceClockCreateFromAudioDeviceID", "10.8", _cMAudioDeviceClockCreateFromAudioDeviceIDErr)
	}
	return _cMAudioDeviceClockCreateFromAudioDeviceID(allocator, deviceID, clockOut), nil
}

// CMAudioDeviceClockCreateFromAudioDeviceID creates a clock that tracks playback through a Core Audio device with the specified identifier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioDeviceClockCreateFromAudioDeviceID(allocator:deviceID:clockOut:)
func CMAudioDeviceClockCreateFromAudioDeviceID(allocator corefoundation.CFAllocatorRef, deviceID uint32, clockOut *uintptr) int32 {
	result, callErr := tryCMAudioDeviceClockCreateFromAudioDeviceID(allocator, deviceID, clockOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMAudioDeviceClockGetAudioDevice func(clock uintptr, deviceUIDOut *corefoundation.CFStringRef, deviceIDOut *uint32, trackingDefaultDeviceOut *bool) int32
var _cMAudioDeviceClockGetAudioDeviceErr error

func tryCMAudioDeviceClockGetAudioDevice(clock uintptr, deviceUIDOut *corefoundation.CFStringRef, deviceIDOut *uint32, trackingDefaultDeviceOut *bool) (int32, error) {
	if _cMAudioDeviceClockGetAudioDevice == nil {
		return 0, symbolCallError("CMAudioDeviceClockGetAudioDevice", "10.8", _cMAudioDeviceClockGetAudioDeviceErr)
	}
	return _cMAudioDeviceClockGetAudioDevice(clock, deviceUIDOut, deviceIDOut, trackingDefaultDeviceOut), nil
}

// CMAudioDeviceClockGetAudioDevice returns the Core Audio device the clock is tracking.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioDeviceClockGetAudioDevice(_:deviceUIDOut:deviceIDOut:trackingDefaultDeviceOut:)
func CMAudioDeviceClockGetAudioDevice(clock uintptr, deviceUIDOut *corefoundation.CFStringRef, deviceIDOut *uint32, trackingDefaultDeviceOut *bool) int32 {
	result, callErr := tryCMAudioDeviceClockGetAudioDevice(clock, deviceUIDOut, deviceIDOut, trackingDefaultDeviceOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMAudioDeviceClockSetAudioDeviceID func(clock uintptr, deviceID uint32) int32
var _cMAudioDeviceClockSetAudioDeviceIDErr error

func tryCMAudioDeviceClockSetAudioDeviceID(clock uintptr, deviceID uint32) (int32, error) {
	if _cMAudioDeviceClockSetAudioDeviceID == nil {
		return 0, symbolCallError("CMAudioDeviceClockSetAudioDeviceID", "10.8", _cMAudioDeviceClockSetAudioDeviceIDErr)
	}
	return _cMAudioDeviceClockSetAudioDeviceID(clock, deviceID), nil
}

// CMAudioDeviceClockSetAudioDeviceID changes the Core Audio device the clock is tracking by specifying a new device identifier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioDeviceClockSetAudioDeviceID(_:deviceID:)
func CMAudioDeviceClockSetAudioDeviceID(clock uintptr, deviceID uint32) int32 {
	result, callErr := tryCMAudioDeviceClockSetAudioDeviceID(clock, deviceID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMAudioDeviceClockSetAudioDeviceUID func(clock uintptr, deviceUID corefoundation.CFStringRef) int32
var _cMAudioDeviceClockSetAudioDeviceUIDErr error

func tryCMAudioDeviceClockSetAudioDeviceUID(clock uintptr, deviceUID corefoundation.CFStringRef) (int32, error) {
	if _cMAudioDeviceClockSetAudioDeviceUID == nil {
		return 0, symbolCallError("CMAudioDeviceClockSetAudioDeviceUID", "10.8", _cMAudioDeviceClockSetAudioDeviceUIDErr)
	}
	return _cMAudioDeviceClockSetAudioDeviceUID(clock, deviceUID), nil
}

// CMAudioDeviceClockSetAudioDeviceUID changes the Core Audio device the clock is tracking by specifying a new device unique identifier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioDeviceClockSetAudioDeviceUID(_:deviceUID:)
func CMAudioDeviceClockSetAudioDeviceUID(clock uintptr, deviceUID corefoundation.CFStringRef) int32 {
	result, callErr := tryCMAudioDeviceClockSetAudioDeviceUID(clock, deviceUID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, audioFormatDescription uintptr, flavor CMSoundDescriptionFlavor, blockBufferOut *uintptr) int32
var _cMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBufferErr error

func tryCMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, audioFormatDescription uintptr, flavor CMSoundDescriptionFlavor, blockBufferOut *uintptr) (int32, error) {
	if _cMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer == nil {
		return 0, symbolCallError("CMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer", "10.10", _cMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBufferErr)
	}
	return _cMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer(allocator, audioFormatDescription, flavor, blockBufferOut), nil
}

// CMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer copies the contents of an audio format description to a buffer in big-endian byte ordering.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer(allocator:audioFormatDescription:flavor:blockBufferOut:)
func CMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, audioFormatDescription uintptr, flavor CMSoundDescriptionFlavor, blockBufferOut *uintptr) int32 {
	result, callErr := tryCMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer(allocator, audioFormatDescription, flavor, blockBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMAudioFormatDescriptionCreate func(allocator corefoundation.CFAllocatorRef, asbd unsafe.Pointer, layoutSize uintptr, layout unsafe.Pointer, magicCookieSize uintptr, magicCookie unsafe.Pointer, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *uintptr) int32
var _cMAudioFormatDescriptionCreateErr error

func tryCMAudioFormatDescriptionCreate(allocator corefoundation.CFAllocatorRef, asbd unsafe.Pointer, layoutSize uintptr, layout unsafe.Pointer, magicCookieSize uintptr, magicCookie unsafe.Pointer, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *uintptr) (int32, error) {
	if _cMAudioFormatDescriptionCreate == nil {
		return 0, symbolCallError("CMAudioFormatDescriptionCreate", "10.7", _cMAudioFormatDescriptionCreateErr)
	}
	return _cMAudioFormatDescriptionCreate(allocator, asbd, layoutSize, layout, magicCookieSize, magicCookie, extensions, formatDescriptionOut), nil
}

// CMAudioFormatDescriptionCreate creates a format description for an audio media stream.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionCreate(allocator:asbd:layoutSize:layout:magicCookieSize:magicCookie:extensions:formatDescriptionOut:)
func CMAudioFormatDescriptionCreate(allocator corefoundation.CFAllocatorRef, asbd unsafe.Pointer, layoutSize uintptr, layout unsafe.Pointer, magicCookieSize uintptr, magicCookie unsafe.Pointer, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *uintptr) int32 {
	result, callErr := tryCMAudioFormatDescriptionCreate(allocator, asbd, layoutSize, layout, magicCookieSize, magicCookie, extensions, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, soundDescriptionBlockBuffer uintptr, flavor CMSoundDescriptionFlavor, formatDescriptionOut *uintptr) int32
var _cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBufferErr error

func tryCMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, soundDescriptionBlockBuffer uintptr, flavor CMSoundDescriptionFlavor, formatDescriptionOut *uintptr) (int32, error) {
	if _cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer == nil {
		return 0, symbolCallError("CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer", "10.10", _cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBufferErr)
	}
	return _cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer(allocator, soundDescriptionBlockBuffer, flavor, formatDescriptionOut), nil
}

// CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer creates an audio format description from a big-endian sound description data structure in a buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer(allocator:bigEndianSoundDescriptionBlockBuffer:flavor:formatDescriptionOut:)
func CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, soundDescriptionBlockBuffer uintptr, flavor CMSoundDescriptionFlavor, formatDescriptionOut *uintptr) int32 {
	result, callErr := tryCMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer(allocator, soundDescriptionBlockBuffer, flavor, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData func(allocator corefoundation.CFAllocatorRef, soundDescriptionData *uint8, size uintptr, flavor CMSoundDescriptionFlavor, formatDescriptionOut *uintptr) int32
var _cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionDataErr error

func tryCMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData(allocator corefoundation.CFAllocatorRef, soundDescriptionData *uint8, size uintptr, flavor CMSoundDescriptionFlavor, formatDescriptionOut *uintptr) (int32, error) {
	if _cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData == nil {
		return 0, symbolCallError("CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData", "10.10", _cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionDataErr)
	}
	return _cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData(allocator, soundDescriptionData, size, flavor, formatDescriptionOut), nil
}

// CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData creates an audio format description from a big-endian sound description data structure.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData(allocator:bigEndianSoundDescriptionData:size:flavor:formatDescriptionOut:)
func CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData(allocator corefoundation.CFAllocatorRef, soundDescriptionData *uint8, size uintptr, flavor CMSoundDescriptionFlavor, formatDescriptionOut *uintptr) int32 {
	result, callErr := tryCMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData(allocator, soundDescriptionData, size, flavor, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMAudioFormatDescriptionCreateSummary func(allocator corefoundation.CFAllocatorRef, formatDescriptionArray corefoundation.CFArrayRef, flags uint32, formatDescriptionOut *uintptr) int32
var _cMAudioFormatDescriptionCreateSummaryErr error

func tryCMAudioFormatDescriptionCreateSummary(allocator corefoundation.CFAllocatorRef, formatDescriptionArray corefoundation.CFArrayRef, flags uint32, formatDescriptionOut *uintptr) (int32, error) {
	if _cMAudioFormatDescriptionCreateSummary == nil {
		return 0, symbolCallError("CMAudioFormatDescriptionCreateSummary", "10.7", _cMAudioFormatDescriptionCreateSummaryErr)
	}
	return _cMAudioFormatDescriptionCreateSummary(allocator, formatDescriptionArray, flags, formatDescriptionOut), nil
}

// CMAudioFormatDescriptionCreateSummary creates a summary audio format description from an array of descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionCreateSummary(allocator:formatDescriptionArray:flags:formatDescriptionOut:)
func CMAudioFormatDescriptionCreateSummary(allocator corefoundation.CFAllocatorRef, formatDescriptionArray corefoundation.CFArrayRef, flags uint32, formatDescriptionOut *uintptr) int32 {
	result, callErr := tryCMAudioFormatDescriptionCreateSummary(allocator, formatDescriptionArray, flags, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMAudioFormatDescriptionEqual func(formatDescription uintptr, otherFormatDescription uintptr, equalityMask CMAudioFormatDescriptionMask, equalityMaskOut *CMAudioFormatDescriptionMask) bool
var _cMAudioFormatDescriptionEqualErr error

func tryCMAudioFormatDescriptionEqual(formatDescription uintptr, otherFormatDescription uintptr, equalityMask CMAudioFormatDescriptionMask, equalityMaskOut *CMAudioFormatDescriptionMask) (bool, error) {
	if _cMAudioFormatDescriptionEqual == nil {
		return false, symbolCallError("CMAudioFormatDescriptionEqual", "10.7", _cMAudioFormatDescriptionEqualErr)
	}
	return _cMAudioFormatDescriptionEqual(formatDescription, otherFormatDescription, equalityMask, equalityMaskOut), nil
}

// CMAudioFormatDescriptionEqual returns a Boolean value that indicates whether the two audio format descriptions are equal.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionEqual(_:otherFormatDescription:equalityMask:equalityMaskOut:)
func CMAudioFormatDescriptionEqual(formatDescription uintptr, otherFormatDescription uintptr, equalityMask CMAudioFormatDescriptionMask, equalityMaskOut *CMAudioFormatDescriptionMask) bool {
	result, callErr := tryCMAudioFormatDescriptionEqual(formatDescription, otherFormatDescription, equalityMask, equalityMaskOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMAudioFormatDescriptionGetChannelLayout func(desc uintptr, sizeOut *uintptr) unsafe.Pointer
var _cMAudioFormatDescriptionGetChannelLayoutErr error

func tryCMAudioFormatDescriptionGetChannelLayout(desc uintptr, sizeOut *uintptr) (unsafe.Pointer, error) {
	if _cMAudioFormatDescriptionGetChannelLayout == nil {
		return nil, symbolCallError("CMAudioFormatDescriptionGetChannelLayout", "10.7", _cMAudioFormatDescriptionGetChannelLayoutErr)
	}
	return _cMAudioFormatDescriptionGetChannelLayout(desc, sizeOut), nil
}

// CMAudioFormatDescriptionGetChannelLayout returns a read-only pointer to, and the size of, the audio channel layout inside an audio format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionGetChannelLayout(_:sizeOut:)
func CMAudioFormatDescriptionGetChannelLayout(desc uintptr, sizeOut *uintptr) unsafe.Pointer {
	result, callErr := tryCMAudioFormatDescriptionGetChannelLayout(desc, sizeOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMAudioFormatDescriptionGetFormatList func(desc uintptr, sizeOut *uintptr) unsafe.Pointer
var _cMAudioFormatDescriptionGetFormatListErr error

func tryCMAudioFormatDescriptionGetFormatList(desc uintptr, sizeOut *uintptr) (unsafe.Pointer, error) {
	if _cMAudioFormatDescriptionGetFormatList == nil {
		return nil, symbolCallError("CMAudioFormatDescriptionGetFormatList", "10.7", _cMAudioFormatDescriptionGetFormatListErr)
	}
	return _cMAudioFormatDescriptionGetFormatList(desc, sizeOut), nil
}

// CMAudioFormatDescriptionGetFormatList returns a read-only pointer to, and size of, the array of audio format list item structures in an audio format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionGetFormatList(_:sizeOut:)
func CMAudioFormatDescriptionGetFormatList(desc uintptr, sizeOut *uintptr) unsafe.Pointer {
	result, callErr := tryCMAudioFormatDescriptionGetFormatList(desc, sizeOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMAudioFormatDescriptionGetMagicCookie func(desc uintptr, sizeOut *uintptr) unsafe.Pointer
var _cMAudioFormatDescriptionGetMagicCookieErr error

func tryCMAudioFormatDescriptionGetMagicCookie(desc uintptr, sizeOut *uintptr) (unsafe.Pointer, error) {
	if _cMAudioFormatDescriptionGetMagicCookie == nil {
		return nil, symbolCallError("CMAudioFormatDescriptionGetMagicCookie", "10.7", _cMAudioFormatDescriptionGetMagicCookieErr)
	}
	return _cMAudioFormatDescriptionGetMagicCookie(desc, sizeOut), nil
}

// CMAudioFormatDescriptionGetMagicCookie returns a read-only pointer to, and size of, the magic cookie in an audio format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionGetMagicCookie(_:sizeOut:)
func CMAudioFormatDescriptionGetMagicCookie(desc uintptr, sizeOut *uintptr) unsafe.Pointer {
	result, callErr := tryCMAudioFormatDescriptionGetMagicCookie(desc, sizeOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMAudioFormatDescriptionGetMostCompatibleFormat func(desc uintptr) unsafe.Pointer
var _cMAudioFormatDescriptionGetMostCompatibleFormatErr error

func tryCMAudioFormatDescriptionGetMostCompatibleFormat(desc uintptr) (unsafe.Pointer, error) {
	if _cMAudioFormatDescriptionGetMostCompatibleFormat == nil {
		return nil, symbolCallError("CMAudioFormatDescriptionGetMostCompatibleFormat", "10.7", _cMAudioFormatDescriptionGetMostCompatibleFormatErr)
	}
	return _cMAudioFormatDescriptionGetMostCompatibleFormat(desc), nil
}

// CMAudioFormatDescriptionGetMostCompatibleFormat returns a read-only pointer to the appropriate audio format list item in an audio format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionGetMostCompatibleFormat(_:)
func CMAudioFormatDescriptionGetMostCompatibleFormat(desc uintptr) unsafe.Pointer {
	result, callErr := tryCMAudioFormatDescriptionGetMostCompatibleFormat(desc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMAudioFormatDescriptionGetRichestDecodableFormat func(desc uintptr) unsafe.Pointer
var _cMAudioFormatDescriptionGetRichestDecodableFormatErr error

func tryCMAudioFormatDescriptionGetRichestDecodableFormat(desc uintptr) (unsafe.Pointer, error) {
	if _cMAudioFormatDescriptionGetRichestDecodableFormat == nil {
		return nil, symbolCallError("CMAudioFormatDescriptionGetRichestDecodableFormat", "10.7", _cMAudioFormatDescriptionGetRichestDecodableFormatErr)
	}
	return _cMAudioFormatDescriptionGetRichestDecodableFormat(desc), nil
}

// CMAudioFormatDescriptionGetRichestDecodableFormat returns a read-only pointer to the appropriate audio format list item in an audio format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionGetRichestDecodableFormat(_:)
func CMAudioFormatDescriptionGetRichestDecodableFormat(desc uintptr) unsafe.Pointer {
	result, callErr := tryCMAudioFormatDescriptionGetRichestDecodableFormat(desc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMAudioFormatDescriptionGetStreamBasicDescription func(desc uintptr) unsafe.Pointer
var _cMAudioFormatDescriptionGetStreamBasicDescriptionErr error

func tryCMAudioFormatDescriptionGetStreamBasicDescription(desc uintptr) (unsafe.Pointer, error) {
	if _cMAudioFormatDescriptionGetStreamBasicDescription == nil {
		return nil, symbolCallError("CMAudioFormatDescriptionGetStreamBasicDescription", "10.7", _cMAudioFormatDescriptionGetStreamBasicDescriptionErr)
	}
	return _cMAudioFormatDescriptionGetStreamBasicDescription(desc), nil
}

// CMAudioFormatDescriptionGetStreamBasicDescription returns a read-only pointer to the audio stream description in an audio format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionGetStreamBasicDescription(_:)
func CMAudioFormatDescriptionGetStreamBasicDescription(desc uintptr) unsafe.Pointer {
	result, callErr := tryCMAudioFormatDescriptionGetStreamBasicDescription(desc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMAudioSampleBufferCreateReadyWithPacketDescriptions func(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, formatDescription uintptr, numSamples int, presentationTimeStamp CMTime, packetDescriptions unsafe.Pointer, sampleBufferOut *uintptr) int32
var _cMAudioSampleBufferCreateReadyWithPacketDescriptionsErr error

func tryCMAudioSampleBufferCreateReadyWithPacketDescriptions(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, formatDescription uintptr, numSamples int, presentationTimeStamp CMTime, packetDescriptions unsafe.Pointer, sampleBufferOut *uintptr) (int32, error) {
	if _cMAudioSampleBufferCreateReadyWithPacketDescriptions == nil {
		return 0, symbolCallError("CMAudioSampleBufferCreateReadyWithPacketDescriptions", "10.10", _cMAudioSampleBufferCreateReadyWithPacketDescriptionsErr)
	}
	return _cMAudioSampleBufferCreateReadyWithPacketDescriptions(allocator, dataBuffer, formatDescription, numSamples, presentationTimeStamp, packetDescriptions, sampleBufferOut), nil
}

// CMAudioSampleBufferCreateReadyWithPacketDescriptions creates a sample buffer with packet descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioSampleBufferCreateReadyWithPacketDescriptions(allocator:dataBuffer:formatDescription:sampleCount:presentationTimeStamp:packetDescriptions:sampleBufferOut:)
func CMAudioSampleBufferCreateReadyWithPacketDescriptions(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, formatDescription uintptr, numSamples int, presentationTimeStamp CMTime, packetDescriptions unsafe.Pointer, sampleBufferOut *uintptr) int32 {
	result, callErr := tryCMAudioSampleBufferCreateReadyWithPacketDescriptions(allocator, dataBuffer, formatDescription, numSamples, presentationTimeStamp, packetDescriptions, sampleBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMAudioSampleBufferCreateWithPacketDescriptions func(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, dataReady bool, makeDataReadyCallback CMSampleBufferMakeDataReadyCallback, makeDataReadyRefcon unsafe.Pointer, formatDescription uintptr, numSamples int, presentationTimeStamp CMTime, packetDescriptions unsafe.Pointer, sampleBufferOut *uintptr) int32
var _cMAudioSampleBufferCreateWithPacketDescriptionsErr error

func tryCMAudioSampleBufferCreateWithPacketDescriptions(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, dataReady bool, makeDataReadyCallback CMSampleBufferMakeDataReadyCallback, makeDataReadyRefcon unsafe.Pointer, formatDescription uintptr, numSamples int, presentationTimeStamp CMTime, packetDescriptions unsafe.Pointer, sampleBufferOut *uintptr) (int32, error) {
	if _cMAudioSampleBufferCreateWithPacketDescriptions == nil {
		return 0, symbolCallError("CMAudioSampleBufferCreateWithPacketDescriptions", "10.7", _cMAudioSampleBufferCreateWithPacketDescriptionsErr)
	}
	return _cMAudioSampleBufferCreateWithPacketDescriptions(allocator, dataBuffer, dataReady, makeDataReadyCallback, makeDataReadyRefcon, formatDescription, numSamples, presentationTimeStamp, packetDescriptions, sampleBufferOut), nil
}

// CMAudioSampleBufferCreateWithPacketDescriptions creates a sample buffer with packet descriptions and a callback to make the data ready for use.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioSampleBufferCreateWithPacketDescriptions(allocator:dataBuffer:dataReady:makeDataReadyCallback:refcon:formatDescription:sampleCount:presentationTimeStamp:packetDescriptions:sampleBufferOut:)
func CMAudioSampleBufferCreateWithPacketDescriptions(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, dataReady bool, makeDataReadyCallback CMSampleBufferMakeDataReadyCallback, makeDataReadyRefcon unsafe.Pointer, formatDescription uintptr, numSamples int, presentationTimeStamp CMTime, packetDescriptions unsafe.Pointer, sampleBufferOut *uintptr) int32 {
	result, callErr := tryCMAudioSampleBufferCreateWithPacketDescriptions(allocator, dataBuffer, dataReady, makeDataReadyCallback, makeDataReadyRefcon, formatDescription, numSamples, presentationTimeStamp, packetDescriptions, sampleBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler func(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, dataReady bool, formatDescription uintptr, numSamples int, presentationTimeStamp CMTime, packetDescriptions unsafe.Pointer, sampleBufferOut *uintptr, makeDataReadyHandler unsafe.Pointer) int32
var _cMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandlerErr error

func tryCMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, dataReady bool, formatDescription uintptr, numSamples int, presentationTimeStamp CMTime, packetDescriptions unsafe.Pointer, sampleBufferOut *uintptr, makeDataReadyHandler CMSampleBufferMakeDataReadyHandler) (int32, error) {
	if _cMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler == nil {
		return 0, symbolCallError("CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler", "10.14.4", _cMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.IObject) int { return makeDataReadyHandler(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _cMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler(allocator, dataBuffer, dataReady, formatDescription, numSamples, presentationTimeStamp, packetDescriptions, sampleBufferOut, _block0), nil
}

// CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler creates a sample buffer with packet descriptions and a handler to make the data ready for use.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler(_:_:_:_:_:_:_:_:_:)
func CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, dataReady bool, formatDescription uintptr, numSamples int, presentationTimeStamp CMTime, packetDescriptions unsafe.Pointer, sampleBufferOut *uintptr, makeDataReadyHandler CMSampleBufferMakeDataReadyHandler) int32 {
	result, callErr := tryCMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler(allocator, dataBuffer, dataReady, formatDescription, numSamples, presentationTimeStamp, packetDescriptions, sampleBufferOut, makeDataReadyHandler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBlockBufferAccessDataBytes func(theBuffer uintptr, offset uintptr, length uintptr, temporaryBlock unsafe.Pointer, returnedPointerOut string) int32
var _cMBlockBufferAccessDataBytesErr error

func tryCMBlockBufferAccessDataBytes(theBuffer uintptr, offset uintptr, length uintptr, temporaryBlock unsafe.Pointer, returnedPointerOut string) (int32, error) {
	if _cMBlockBufferAccessDataBytes == nil {
		return 0, symbolCallError("CMBlockBufferAccessDataBytes", "10.7", _cMBlockBufferAccessDataBytesErr)
	}
	return _cMBlockBufferAccessDataBytes(theBuffer, offset, length, temporaryBlock, returnedPointerOut), nil
}

// CMBlockBufferAccessDataBytes accesses potentially noncontiguous data in a block buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferAccessDataBytes(_:atOffset:length:temporaryBlock:returnedPointerOut:)
func CMBlockBufferAccessDataBytes(theBuffer uintptr, offset uintptr, length uintptr, temporaryBlock unsafe.Pointer, returnedPointerOut string) int32 {
	result, callErr := tryCMBlockBufferAccessDataBytes(theBuffer, offset, length, temporaryBlock, returnedPointerOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBlockBufferAppendBufferReference func(theBuffer uintptr, targetBBuf uintptr, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags) int32
var _cMBlockBufferAppendBufferReferenceErr error

func tryCMBlockBufferAppendBufferReference(theBuffer uintptr, targetBBuf uintptr, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags) (int32, error) {
	if _cMBlockBufferAppendBufferReference == nil {
		return 0, symbolCallError("CMBlockBufferAppendBufferReference", "10.7", _cMBlockBufferAppendBufferReferenceErr)
	}
	return _cMBlockBufferAppendBufferReference(theBuffer, targetBBuf, offsetToData, dataLength, flags), nil
}

// CMBlockBufferAppendBufferReference adds a reference to an existing block buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferAppendBufferReference(_:targetBBuf:offsetToData:dataLength:flags:)
func CMBlockBufferAppendBufferReference(theBuffer uintptr, targetBBuf uintptr, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags) int32 {
	result, callErr := tryCMBlockBufferAppendBufferReference(theBuffer, targetBBuf, offsetToData, dataLength, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBlockBufferAppendMemoryBlock func(theBuffer uintptr, memoryBlock unsafe.Pointer, blockLength uintptr, blockAllocator corefoundation.CFAllocatorRef, customBlockSource *CMBlockBufferCustomBlockSource, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags) int32
var _cMBlockBufferAppendMemoryBlockErr error

func tryCMBlockBufferAppendMemoryBlock(theBuffer uintptr, memoryBlock unsafe.Pointer, blockLength uintptr, blockAllocator corefoundation.CFAllocatorRef, customBlockSource *CMBlockBufferCustomBlockSource, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags) (int32, error) {
	if _cMBlockBufferAppendMemoryBlock == nil {
		return 0, symbolCallError("CMBlockBufferAppendMemoryBlock", "10.7", _cMBlockBufferAppendMemoryBlockErr)
	}
	return _cMBlockBufferAppendMemoryBlock(theBuffer, memoryBlock, blockLength, blockAllocator, customBlockSource, offsetToData, dataLength, flags), nil
}

// CMBlockBufferAppendMemoryBlock adds a memory block to an existing block buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferAppendMemoryBlock(_:memoryBlock:length:blockAllocator:customBlockSource:offsetToData:dataLength:flags:)
func CMBlockBufferAppendMemoryBlock(theBuffer uintptr, memoryBlock unsafe.Pointer, blockLength uintptr, blockAllocator corefoundation.CFAllocatorRef, customBlockSource *CMBlockBufferCustomBlockSource, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags) int32 {
	result, callErr := tryCMBlockBufferAppendMemoryBlock(theBuffer, memoryBlock, blockLength, blockAllocator, customBlockSource, offsetToData, dataLength, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBlockBufferAssureBlockMemory func(theBuffer uintptr) int32
var _cMBlockBufferAssureBlockMemoryErr error

func tryCMBlockBufferAssureBlockMemory(theBuffer uintptr) (int32, error) {
	if _cMBlockBufferAssureBlockMemory == nil {
		return 0, symbolCallError("CMBlockBufferAssureBlockMemory", "10.7", _cMBlockBufferAssureBlockMemoryErr)
	}
	return _cMBlockBufferAssureBlockMemory(theBuffer), nil
}

// CMBlockBufferAssureBlockMemory assures that the system allocates memory for all memory blocks in a block buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferAssureBlockMemory(_:)
func CMBlockBufferAssureBlockMemory(theBuffer uintptr) int32 {
	result, callErr := tryCMBlockBufferAssureBlockMemory(theBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBlockBufferCopyDataBytes func(theSourceBuffer uintptr, offsetToData uintptr, dataLength uintptr, destination unsafe.Pointer) int32
var _cMBlockBufferCopyDataBytesErr error

func tryCMBlockBufferCopyDataBytes(theSourceBuffer uintptr, offsetToData uintptr, dataLength uintptr, destination unsafe.Pointer) (int32, error) {
	if _cMBlockBufferCopyDataBytes == nil {
		return 0, symbolCallError("CMBlockBufferCopyDataBytes", "10.7", _cMBlockBufferCopyDataBytesErr)
	}
	return _cMBlockBufferCopyDataBytes(theSourceBuffer, offsetToData, dataLength, destination), nil
}

// CMBlockBufferCopyDataBytes copies bytes from a block buffer into a provided memory area.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCopyDataBytes(_:atOffset:dataLength:destination:)
func CMBlockBufferCopyDataBytes(theSourceBuffer uintptr, offsetToData uintptr, dataLength uintptr, destination unsafe.Pointer) int32 {
	result, callErr := tryCMBlockBufferCopyDataBytes(theSourceBuffer, offsetToData, dataLength, destination)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBlockBufferCreateContiguous func(structureAllocator corefoundation.CFAllocatorRef, sourceBuffer uintptr, blockAllocator corefoundation.CFAllocatorRef, customBlockSource *CMBlockBufferCustomBlockSource, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags, blockBufferOut *uintptr) int32
var _cMBlockBufferCreateContiguousErr error

func tryCMBlockBufferCreateContiguous(structureAllocator corefoundation.CFAllocatorRef, sourceBuffer uintptr, blockAllocator corefoundation.CFAllocatorRef, customBlockSource *CMBlockBufferCustomBlockSource, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags, blockBufferOut *uintptr) (int32, error) {
	if _cMBlockBufferCreateContiguous == nil {
		return 0, symbolCallError("CMBlockBufferCreateContiguous", "10.7", _cMBlockBufferCreateContiguousErr)
	}
	return _cMBlockBufferCreateContiguous(structureAllocator, sourceBuffer, blockAllocator, customBlockSource, offsetToData, dataLength, flags, blockBufferOut), nil
}

// CMBlockBufferCreateContiguous creates a block buffer that contains a contiguous copy of, or reference to, the data specified by the parameters.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCreateContiguous(allocator:sourceBuffer:blockAllocator:customBlockSource:offsetToData:dataLength:flags:blockBufferOut:)
func CMBlockBufferCreateContiguous(structureAllocator corefoundation.CFAllocatorRef, sourceBuffer uintptr, blockAllocator corefoundation.CFAllocatorRef, customBlockSource *CMBlockBufferCustomBlockSource, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags, blockBufferOut *uintptr) int32 {
	result, callErr := tryCMBlockBufferCreateContiguous(structureAllocator, sourceBuffer, blockAllocator, customBlockSource, offsetToData, dataLength, flags, blockBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBlockBufferCreateEmpty func(structureAllocator corefoundation.CFAllocatorRef, subBlockCapacity uint32, flags CMBlockBufferFlags, blockBufferOut *uintptr) int32
var _cMBlockBufferCreateEmptyErr error

func tryCMBlockBufferCreateEmpty(structureAllocator corefoundation.CFAllocatorRef, subBlockCapacity uint32, flags CMBlockBufferFlags, blockBufferOut *uintptr) (int32, error) {
	if _cMBlockBufferCreateEmpty == nil {
		return 0, symbolCallError("CMBlockBufferCreateEmpty", "10.7", _cMBlockBufferCreateEmptyErr)
	}
	return _cMBlockBufferCreateEmpty(structureAllocator, subBlockCapacity, flags, blockBufferOut), nil
}

// CMBlockBufferCreateEmpty creates an empty block buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCreateEmpty(allocator:capacity:flags:blockBufferOut:)
func CMBlockBufferCreateEmpty(structureAllocator corefoundation.CFAllocatorRef, subBlockCapacity uint32, flags CMBlockBufferFlags, blockBufferOut *uintptr) int32 {
	result, callErr := tryCMBlockBufferCreateEmpty(structureAllocator, subBlockCapacity, flags, blockBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBlockBufferCreateWithBufferReference func(structureAllocator corefoundation.CFAllocatorRef, bufferReference uintptr, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags, blockBufferOut *uintptr) int32
var _cMBlockBufferCreateWithBufferReferenceErr error

func tryCMBlockBufferCreateWithBufferReference(structureAllocator corefoundation.CFAllocatorRef, bufferReference uintptr, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags, blockBufferOut *uintptr) (int32, error) {
	if _cMBlockBufferCreateWithBufferReference == nil {
		return 0, symbolCallError("CMBlockBufferCreateWithBufferReference", "10.7", _cMBlockBufferCreateWithBufferReferenceErr)
	}
	return _cMBlockBufferCreateWithBufferReference(structureAllocator, bufferReference, offsetToData, dataLength, flags, blockBufferOut), nil
}

// CMBlockBufferCreateWithBufferReference creates a block buffer that refers to another block buffer object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCreateWithBufferReference(allocator:referenceBuffer:offsetToData:dataLength:flags:blockBufferOut:)
func CMBlockBufferCreateWithBufferReference(structureAllocator corefoundation.CFAllocatorRef, bufferReference uintptr, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags, blockBufferOut *uintptr) int32 {
	result, callErr := tryCMBlockBufferCreateWithBufferReference(structureAllocator, bufferReference, offsetToData, dataLength, flags, blockBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBlockBufferCreateWithMemoryBlock func(structureAllocator corefoundation.CFAllocatorRef, memoryBlock unsafe.Pointer, blockLength uintptr, blockAllocator corefoundation.CFAllocatorRef, customBlockSource *CMBlockBufferCustomBlockSource, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags, blockBufferOut *uintptr) int32
var _cMBlockBufferCreateWithMemoryBlockErr error

func tryCMBlockBufferCreateWithMemoryBlock(structureAllocator corefoundation.CFAllocatorRef, memoryBlock unsafe.Pointer, blockLength uintptr, blockAllocator corefoundation.CFAllocatorRef, customBlockSource *CMBlockBufferCustomBlockSource, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags, blockBufferOut *uintptr) (int32, error) {
	if _cMBlockBufferCreateWithMemoryBlock == nil {
		return 0, symbolCallError("CMBlockBufferCreateWithMemoryBlock", "10.7", _cMBlockBufferCreateWithMemoryBlockErr)
	}
	return _cMBlockBufferCreateWithMemoryBlock(structureAllocator, memoryBlock, blockLength, blockAllocator, customBlockSource, offsetToData, dataLength, flags, blockBufferOut), nil
}

// CMBlockBufferCreateWithMemoryBlock creates a block buffer that’s backed by a memory block.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCreateWithMemoryBlock(allocator:memoryBlock:blockLength:blockAllocator:customBlockSource:offsetToData:dataLength:flags:blockBufferOut:)
func CMBlockBufferCreateWithMemoryBlock(structureAllocator corefoundation.CFAllocatorRef, memoryBlock unsafe.Pointer, blockLength uintptr, blockAllocator corefoundation.CFAllocatorRef, customBlockSource *CMBlockBufferCustomBlockSource, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags, blockBufferOut *uintptr) int32 {
	result, callErr := tryCMBlockBufferCreateWithMemoryBlock(structureAllocator, memoryBlock, blockLength, blockAllocator, customBlockSource, offsetToData, dataLength, flags, blockBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBlockBufferFillDataBytes func(fillByte int8, destinationBuffer uintptr, offsetIntoDestination uintptr, dataLength uintptr) int32
var _cMBlockBufferFillDataBytesErr error

func tryCMBlockBufferFillDataBytes(fillByte int8, destinationBuffer uintptr, offsetIntoDestination uintptr, dataLength uintptr) (int32, error) {
	if _cMBlockBufferFillDataBytes == nil {
		return 0, symbolCallError("CMBlockBufferFillDataBytes", "10.7", _cMBlockBufferFillDataBytesErr)
	}
	return _cMBlockBufferFillDataBytes(fillByte, destinationBuffer, offsetIntoDestination, dataLength), nil
}

// CMBlockBufferFillDataBytes fills the destination buffer with the specified data byte.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferFillDataBytes(with:blockBuffer:offsetIntoDestination:dataLength:)
func CMBlockBufferFillDataBytes(fillByte int8, destinationBuffer uintptr, offsetIntoDestination uintptr, dataLength uintptr) int32 {
	result, callErr := tryCMBlockBufferFillDataBytes(fillByte, destinationBuffer, offsetIntoDestination, dataLength)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBlockBufferGetDataLength func(theBuffer uintptr) uintptr
var _cMBlockBufferGetDataLengthErr error

func tryCMBlockBufferGetDataLength(theBuffer uintptr) (uintptr, error) {
	if _cMBlockBufferGetDataLength == nil {
		return 0, symbolCallError("CMBlockBufferGetDataLength", "10.7", _cMBlockBufferGetDataLengthErr)
	}
	return _cMBlockBufferGetDataLength(theBuffer), nil
}

// CMBlockBufferGetDataLength returns the total length of data that’s accessible by a block buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferGetDataLength(_:)
func CMBlockBufferGetDataLength(theBuffer uintptr) uintptr {
	result, callErr := tryCMBlockBufferGetDataLength(theBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBlockBufferGetDataPointer func(theBuffer uintptr, offset uintptr, lengthAtOffsetOut *uintptr, totalLengthOut *uintptr, dataPointerOut string) int32
var _cMBlockBufferGetDataPointerErr error

func tryCMBlockBufferGetDataPointer(theBuffer uintptr, offset uintptr, lengthAtOffsetOut *uintptr, totalLengthOut *uintptr, dataPointerOut string) (int32, error) {
	if _cMBlockBufferGetDataPointer == nil {
		return 0, symbolCallError("CMBlockBufferGetDataPointer", "10.7", _cMBlockBufferGetDataPointerErr)
	}
	return _cMBlockBufferGetDataPointer(theBuffer, offset, lengthAtOffsetOut, totalLengthOut, dataPointerOut), nil
}

// CMBlockBufferGetDataPointer gains access to the data represented by a block buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferGetDataPointer(_:atOffset:lengthAtOffsetOut:totalLengthOut:dataPointerOut:)
func CMBlockBufferGetDataPointer(theBuffer uintptr, offset uintptr, lengthAtOffsetOut *uintptr, totalLengthOut *uintptr, dataPointerOut string) int32 {
	result, callErr := tryCMBlockBufferGetDataPointer(theBuffer, offset, lengthAtOffsetOut, totalLengthOut, dataPointerOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBlockBufferGetTypeID func() uint
var _cMBlockBufferGetTypeIDErr error

func tryCMBlockBufferGetTypeID() (uint, error) {
	if _cMBlockBufferGetTypeID == nil {
		return 0, symbolCallError("CMBlockBufferGetTypeID", "10.7", _cMBlockBufferGetTypeIDErr)
	}
	return _cMBlockBufferGetTypeID(), nil
}

// CMBlockBufferGetTypeID returns the type identifier for block buffer objects.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferGetTypeID()
func CMBlockBufferGetTypeID() uint {
	result, callErr := tryCMBlockBufferGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBlockBufferIsEmpty func(theBuffer uintptr) bool
var _cMBlockBufferIsEmptyErr error

func tryCMBlockBufferIsEmpty(theBuffer uintptr) (bool, error) {
	if _cMBlockBufferIsEmpty == nil {
		return false, symbolCallError("CMBlockBufferIsEmpty", "10.7", _cMBlockBufferIsEmptyErr)
	}
	return _cMBlockBufferIsEmpty(theBuffer), nil
}

// CMBlockBufferIsEmpty returns a Boolean value that indicates whether the buffer is empty.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferIsEmpty(_:)
func CMBlockBufferIsEmpty(theBuffer uintptr) bool {
	result, callErr := tryCMBlockBufferIsEmpty(theBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBlockBufferIsRangeContiguous func(theBuffer uintptr, offset uintptr, length uintptr) bool
var _cMBlockBufferIsRangeContiguousErr error

func tryCMBlockBufferIsRangeContiguous(theBuffer uintptr, offset uintptr, length uintptr) (bool, error) {
	if _cMBlockBufferIsRangeContiguous == nil {
		return false, symbolCallError("CMBlockBufferIsRangeContiguous", "10.7", _cMBlockBufferIsRangeContiguousErr)
	}
	return _cMBlockBufferIsRangeContiguous(theBuffer, offset, length), nil
}

// CMBlockBufferIsRangeContiguous returns a Boolean value that indicates whether the specified range within a block buffer is contiguous.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferIsRangeContiguous(_:atOffset:length:)
func CMBlockBufferIsRangeContiguous(theBuffer uintptr, offset uintptr, length uintptr) bool {
	result, callErr := tryCMBlockBufferIsRangeContiguous(theBuffer, offset, length)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBlockBufferReplaceDataBytes func(sourceBytes unsafe.Pointer, destinationBuffer uintptr, offsetIntoDestination uintptr, dataLength uintptr) int32
var _cMBlockBufferReplaceDataBytesErr error

func tryCMBlockBufferReplaceDataBytes(sourceBytes unsafe.Pointer, destinationBuffer uintptr, offsetIntoDestination uintptr, dataLength uintptr) (int32, error) {
	if _cMBlockBufferReplaceDataBytes == nil {
		return 0, symbolCallError("CMBlockBufferReplaceDataBytes", "10.7", _cMBlockBufferReplaceDataBytesErr)
	}
	return _cMBlockBufferReplaceDataBytes(sourceBytes, destinationBuffer, offsetIntoDestination, dataLength), nil
}

// CMBlockBufferReplaceDataBytes copies bytes from a given memory block into a block buffer replacing bytes in the underlying data blocks.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferReplaceDataBytes(with:blockBuffer:offsetIntoDestination:dataLength:)
func CMBlockBufferReplaceDataBytes(sourceBytes unsafe.Pointer, destinationBuffer uintptr, offsetIntoDestination uintptr, dataLength uintptr) int32 {
	result, callErr := tryCMBlockBufferReplaceDataBytes(sourceBytes, destinationBuffer, offsetIntoDestination, dataLength)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueCallForEachBuffer func(queue uintptr, refcon uintptr) int32
var _cMBufferQueueCallForEachBufferErr error

func tryCMBufferQueueCallForEachBuffer(queue uintptr, refcon uintptr) (int32, error) {
	if _cMBufferQueueCallForEachBuffer == nil {
		return 0, symbolCallError("CMBufferQueueCallForEachBuffer", "10.7", _cMBufferQueueCallForEachBufferErr)
	}
	return _cMBufferQueueCallForEachBuffer(queue, refcon), nil
}

// CMBufferQueueCallForEachBuffer calls a function for every buffer in a queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueCallForEachBuffer(_:callback:refcon:)
func CMBufferQueueCallForEachBuffer(queue uintptr, refcon uintptr) int32 {
	result, callErr := tryCMBufferQueueCallForEachBuffer(queue, refcon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueContainsEndOfData func(queue uintptr) bool
var _cMBufferQueueContainsEndOfDataErr error

func tryCMBufferQueueContainsEndOfData(queue uintptr) (bool, error) {
	if _cMBufferQueueContainsEndOfData == nil {
		return false, symbolCallError("CMBufferQueueContainsEndOfData", "10.7", _cMBufferQueueContainsEndOfDataErr)
	}
	return _cMBufferQueueContainsEndOfData(queue), nil
}

// CMBufferQueueContainsEndOfData returns a Boolean value that indicates whether a buffer queue has its end-of-data marker set.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueContainsEndOfData(_:)
func CMBufferQueueContainsEndOfData(queue uintptr) bool {
	result, callErr := tryCMBufferQueueContainsEndOfData(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueCopyHead func(queue uintptr) CMBufferRef
var _cMBufferQueueCopyHeadErr error

func tryCMBufferQueueCopyHead(queue uintptr) (CMBufferRef, error) {
	if _cMBufferQueueCopyHead == nil {
		return 0, symbolCallError("CMBufferQueueCopyHead", "14.0", _cMBufferQueueCopyHeadErr)
	}
	return _cMBufferQueueCopyHead(queue), nil
}

// CMBufferQueueCopyHead.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueCopyHead(_:)
func CMBufferQueueCopyHead(queue uintptr) CMBufferRef {
	result, callErr := tryCMBufferQueueCopyHead(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueCreate func(allocator corefoundation.CFAllocatorRef, capacity int, callbacks *CMBufferCallbacks, queueOut *uintptr) int32
var _cMBufferQueueCreateErr error

func tryCMBufferQueueCreate(allocator corefoundation.CFAllocatorRef, capacity int, callbacks *CMBufferCallbacks, queueOut *uintptr) (int32, error) {
	if _cMBufferQueueCreate == nil {
		return 0, symbolCallError("CMBufferQueueCreate", "10.7", _cMBufferQueueCreateErr)
	}
	return _cMBufferQueueCreate(allocator, capacity, callbacks, queueOut), nil
}

// CMBufferQueueCreate creates a buffer queue with callbacks to inspect buffers.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueCreate(allocator:capacity:callbacks:queueOut:)
func CMBufferQueueCreate(allocator corefoundation.CFAllocatorRef, capacity int, callbacks *CMBufferCallbacks, queueOut *uintptr) int32 {
	result, callErr := tryCMBufferQueueCreate(allocator, capacity, callbacks, queueOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueCreateWithHandlers func(allocator corefoundation.CFAllocatorRef, capacity int, handlers *CMBufferHandlers, queueOut *uintptr) int32
var _cMBufferQueueCreateWithHandlersErr error

func tryCMBufferQueueCreateWithHandlers(allocator corefoundation.CFAllocatorRef, capacity int, handlers *CMBufferHandlers, queueOut *uintptr) (int32, error) {
	if _cMBufferQueueCreateWithHandlers == nil {
		return 0, symbolCallError("CMBufferQueueCreateWithHandlers", "10.14.4", _cMBufferQueueCreateWithHandlersErr)
	}
	return _cMBufferQueueCreateWithHandlers(allocator, capacity, handlers, queueOut), nil
}

// CMBufferQueueCreateWithHandlers creates a buffer queue with handlers to inspect buffers.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueCreateWithHandlers(_:_:_:_:)
func CMBufferQueueCreateWithHandlers(allocator corefoundation.CFAllocatorRef, capacity int, handlers *CMBufferHandlers, queueOut *uintptr) int32 {
	result, callErr := tryCMBufferQueueCreateWithHandlers(allocator, capacity, handlers, queueOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueDequeueAndRetain func(queue uintptr) CMBufferRef
var _cMBufferQueueDequeueAndRetainErr error

func tryCMBufferQueueDequeueAndRetain(queue uintptr) (CMBufferRef, error) {
	if _cMBufferQueueDequeueAndRetain == nil {
		return 0, symbolCallError("CMBufferQueueDequeueAndRetain", "10.7", _cMBufferQueueDequeueAndRetainErr)
	}
	return _cMBufferQueueDequeueAndRetain(queue), nil
}

// CMBufferQueueDequeueAndRetain dequeues a buffer from a queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueDequeue(_:)
func CMBufferQueueDequeueAndRetain(queue uintptr) CMBufferRef {
	result, callErr := tryCMBufferQueueDequeueAndRetain(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueDequeueIfDataReadyAndRetain func(queue uintptr) CMBufferRef
var _cMBufferQueueDequeueIfDataReadyAndRetainErr error

func tryCMBufferQueueDequeueIfDataReadyAndRetain(queue uintptr) (CMBufferRef, error) {
	if _cMBufferQueueDequeueIfDataReadyAndRetain == nil {
		return 0, symbolCallError("CMBufferQueueDequeueIfDataReadyAndRetain", "10.7", _cMBufferQueueDequeueIfDataReadyAndRetainErr)
	}
	return _cMBufferQueueDequeueIfDataReadyAndRetain(queue), nil
}

// CMBufferQueueDequeueIfDataReadyAndRetain dequeues a buffer from a queue, if it’s ready.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueDequeueIfDataReady(_:)
func CMBufferQueueDequeueIfDataReadyAndRetain(queue uintptr) CMBufferRef {
	result, callErr := tryCMBufferQueueDequeueIfDataReadyAndRetain(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueEnqueue func(queue uintptr, buf CMBufferRef) int32
var _cMBufferQueueEnqueueErr error

func tryCMBufferQueueEnqueue(queue uintptr, buf CMBufferRef) (int32, error) {
	if _cMBufferQueueEnqueue == nil {
		return 0, symbolCallError("CMBufferQueueEnqueue", "10.7", _cMBufferQueueEnqueueErr)
	}
	return _cMBufferQueueEnqueue(queue, buf), nil
}

// CMBufferQueueEnqueue enqueues a buffer onto a queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueEnqueue(_:buffer:)
func CMBufferQueueEnqueue(queue uintptr, buf CMBufferRef) int32 {
	result, callErr := tryCMBufferQueueEnqueue(queue, buf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueGetBufferCount func(queue uintptr) int
var _cMBufferQueueGetBufferCountErr error

func tryCMBufferQueueGetBufferCount(queue uintptr) (int, error) {
	if _cMBufferQueueGetBufferCount == nil {
		return 0, symbolCallError("CMBufferQueueGetBufferCount", "10.7", _cMBufferQueueGetBufferCountErr)
	}
	return _cMBufferQueueGetBufferCount(queue), nil
}

// CMBufferQueueGetBufferCount gets the number of buffers in the queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetBufferCount(_:)
func CMBufferQueueGetBufferCount(queue uintptr) int {
	result, callErr := tryCMBufferQueueGetBufferCount(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS func() *CMBufferCallbacks
var _cMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTSErr error

func tryCMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS() (*CMBufferCallbacks, error) {
	if _cMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS == nil {
		return nil, symbolCallError("CMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS", "10.7", _cMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTSErr)
	}
	return _cMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS(), nil
}

// CMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS returns a pointer to a structure that contains callbacks to sort sample buffers by output presentation timestamp.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS()
func CMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS() *CMBufferCallbacks {
	result, callErr := tryCMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueGetCallbacksForUnsortedSampleBuffers func() *CMBufferCallbacks
var _cMBufferQueueGetCallbacksForUnsortedSampleBuffersErr error

func tryCMBufferQueueGetCallbacksForUnsortedSampleBuffers() (*CMBufferCallbacks, error) {
	if _cMBufferQueueGetCallbacksForUnsortedSampleBuffers == nil {
		return nil, symbolCallError("CMBufferQueueGetCallbacksForUnsortedSampleBuffers", "10.7", _cMBufferQueueGetCallbacksForUnsortedSampleBuffersErr)
	}
	return _cMBufferQueueGetCallbacksForUnsortedSampleBuffers(), nil
}

// CMBufferQueueGetCallbacksForUnsortedSampleBuffers returns a pointer to a callback structure for unsorted sample buffers.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetCallbacksForUnsortedSampleBuffers()
func CMBufferQueueGetCallbacksForUnsortedSampleBuffers() *CMBufferCallbacks {
	result, callErr := tryCMBufferQueueGetCallbacksForUnsortedSampleBuffers()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueGetDuration func(queue uintptr) CMTime
var _cMBufferQueueGetDurationErr error

func tryCMBufferQueueGetDuration(queue uintptr) (CMTime, error) {
	if _cMBufferQueueGetDuration == nil {
		return CMTime{}, symbolCallError("CMBufferQueueGetDuration", "10.7", _cMBufferQueueGetDurationErr)
	}
	return _cMBufferQueueGetDuration(queue), nil
}

// CMBufferQueueGetDuration gets the duration of a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetDuration(_:)
func CMBufferQueueGetDuration(queue uintptr) CMTime {
	result, callErr := tryCMBufferQueueGetDuration(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueGetEndPresentationTimeStamp func(queue uintptr) CMTime
var _cMBufferQueueGetEndPresentationTimeStampErr error

func tryCMBufferQueueGetEndPresentationTimeStamp(queue uintptr) (CMTime, error) {
	if _cMBufferQueueGetEndPresentationTimeStamp == nil {
		return CMTime{}, symbolCallError("CMBufferQueueGetEndPresentationTimeStamp", "10.7", _cMBufferQueueGetEndPresentationTimeStampErr)
	}
	return _cMBufferQueueGetEndPresentationTimeStamp(queue), nil
}

// CMBufferQueueGetEndPresentationTimeStamp gets the greatest end presentation timestamp of a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetEndPresentationTimeStamp(_:)
func CMBufferQueueGetEndPresentationTimeStamp(queue uintptr) CMTime {
	result, callErr := tryCMBufferQueueGetEndPresentationTimeStamp(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueGetFirstDecodeTimeStamp func(queue uintptr) CMTime
var _cMBufferQueueGetFirstDecodeTimeStampErr error

func tryCMBufferQueueGetFirstDecodeTimeStamp(queue uintptr) (CMTime, error) {
	if _cMBufferQueueGetFirstDecodeTimeStamp == nil {
		return CMTime{}, symbolCallError("CMBufferQueueGetFirstDecodeTimeStamp", "10.7", _cMBufferQueueGetFirstDecodeTimeStampErr)
	}
	return _cMBufferQueueGetFirstDecodeTimeStamp(queue), nil
}

// CMBufferQueueGetFirstDecodeTimeStamp gets the decode timestamp of the first buffer in a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetFirstDecodeTimeStamp(_:)
func CMBufferQueueGetFirstDecodeTimeStamp(queue uintptr) CMTime {
	result, callErr := tryCMBufferQueueGetFirstDecodeTimeStamp(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueGetFirstPresentationTimeStamp func(queue uintptr) CMTime
var _cMBufferQueueGetFirstPresentationTimeStampErr error

func tryCMBufferQueueGetFirstPresentationTimeStamp(queue uintptr) (CMTime, error) {
	if _cMBufferQueueGetFirstPresentationTimeStamp == nil {
		return CMTime{}, symbolCallError("CMBufferQueueGetFirstPresentationTimeStamp", "10.7", _cMBufferQueueGetFirstPresentationTimeStampErr)
	}
	return _cMBufferQueueGetFirstPresentationTimeStamp(queue), nil
}

// CMBufferQueueGetFirstPresentationTimeStamp gets the presentation timestamp of the first buffer in a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetFirstPresentationTimeStamp(_:)
func CMBufferQueueGetFirstPresentationTimeStamp(queue uintptr) CMTime {
	result, callErr := tryCMBufferQueueGetFirstPresentationTimeStamp(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueGetHead func(queue uintptr) CMBufferRef
var _cMBufferQueueGetHeadErr error

func tryCMBufferQueueGetHead(queue uintptr) (CMBufferRef, error) {
	if _cMBufferQueueGetHead == nil {
		return 0, symbolCallError("CMBufferQueueGetHead", "10.7", _cMBufferQueueGetHeadErr)
	}
	return _cMBufferQueueGetHead(queue), nil
}

// CMBufferQueueGetHead retrieves the next buffer from a queue, but doesn’t remove it.
//
// Deprecated: Deprecated since macOS 15.0.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetHead(_:)
func CMBufferQueueGetHead(queue uintptr) CMBufferRef {
	result, callErr := tryCMBufferQueueGetHead(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueGetMaxPresentationTimeStamp func(queue uintptr) CMTime
var _cMBufferQueueGetMaxPresentationTimeStampErr error

func tryCMBufferQueueGetMaxPresentationTimeStamp(queue uintptr) (CMTime, error) {
	if _cMBufferQueueGetMaxPresentationTimeStamp == nil {
		return CMTime{}, symbolCallError("CMBufferQueueGetMaxPresentationTimeStamp", "10.7", _cMBufferQueueGetMaxPresentationTimeStampErr)
	}
	return _cMBufferQueueGetMaxPresentationTimeStamp(queue), nil
}

// CMBufferQueueGetMaxPresentationTimeStamp gets the greatest presentation timestamp of a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetMaxPresentationTimeStamp(_:)
func CMBufferQueueGetMaxPresentationTimeStamp(queue uintptr) CMTime {
	result, callErr := tryCMBufferQueueGetMaxPresentationTimeStamp(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueGetMinDecodeTimeStamp func(queue uintptr) CMTime
var _cMBufferQueueGetMinDecodeTimeStampErr error

func tryCMBufferQueueGetMinDecodeTimeStamp(queue uintptr) (CMTime, error) {
	if _cMBufferQueueGetMinDecodeTimeStamp == nil {
		return CMTime{}, symbolCallError("CMBufferQueueGetMinDecodeTimeStamp", "10.7", _cMBufferQueueGetMinDecodeTimeStampErr)
	}
	return _cMBufferQueueGetMinDecodeTimeStamp(queue), nil
}

// CMBufferQueueGetMinDecodeTimeStamp gets the earliest decode timestamp of a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetMinDecodeTimeStamp(_:)
func CMBufferQueueGetMinDecodeTimeStamp(queue uintptr) CMTime {
	result, callErr := tryCMBufferQueueGetMinDecodeTimeStamp(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueGetMinPresentationTimeStamp func(queue uintptr) CMTime
var _cMBufferQueueGetMinPresentationTimeStampErr error

func tryCMBufferQueueGetMinPresentationTimeStamp(queue uintptr) (CMTime, error) {
	if _cMBufferQueueGetMinPresentationTimeStamp == nil {
		return CMTime{}, symbolCallError("CMBufferQueueGetMinPresentationTimeStamp", "10.7", _cMBufferQueueGetMinPresentationTimeStampErr)
	}
	return _cMBufferQueueGetMinPresentationTimeStamp(queue), nil
}

// CMBufferQueueGetMinPresentationTimeStamp gets the earliest presentation timestamp of a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetMinPresentationTimeStamp(_:)
func CMBufferQueueGetMinPresentationTimeStamp(queue uintptr) CMTime {
	result, callErr := tryCMBufferQueueGetMinPresentationTimeStamp(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueGetTotalSize func(queue uintptr) uintptr
var _cMBufferQueueGetTotalSizeErr error

func tryCMBufferQueueGetTotalSize(queue uintptr) (uintptr, error) {
	if _cMBufferQueueGetTotalSize == nil {
		return 0, symbolCallError("CMBufferQueueGetTotalSize", "10.10", _cMBufferQueueGetTotalSizeErr)
	}
	return _cMBufferQueueGetTotalSize(queue), nil
}

// CMBufferQueueGetTotalSize gets the total size of all sample buffers of a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetTotalSize(_:)
func CMBufferQueueGetTotalSize(queue uintptr) uintptr {
	result, callErr := tryCMBufferQueueGetTotalSize(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueGetTypeID func() uint
var _cMBufferQueueGetTypeIDErr error

func tryCMBufferQueueGetTypeID() (uint, error) {
	if _cMBufferQueueGetTypeID == nil {
		return 0, symbolCallError("CMBufferQueueGetTypeID", "10.7", _cMBufferQueueGetTypeIDErr)
	}
	return _cMBufferQueueGetTypeID(), nil
}

// CMBufferQueueGetTypeID returns the type identifier of buffer queue objects.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetTypeID()
func CMBufferQueueGetTypeID() uint {
	result, callErr := tryCMBufferQueueGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueInstallTrigger func(queue uintptr, callback CMBufferQueueTriggerCallback, refcon uintptr, condition CMBufferQueueTriggerCondition, time CMTime, triggerTokenOut *CMBufferQueueTriggerToken) int32
var _cMBufferQueueInstallTriggerErr error

func tryCMBufferQueueInstallTrigger(queue uintptr, callback CMBufferQueueTriggerCallback, refcon uintptr, condition CMBufferQueueTriggerCondition, time CMTime, triggerTokenOut *CMBufferQueueTriggerToken) (int32, error) {
	if _cMBufferQueueInstallTrigger == nil {
		return 0, symbolCallError("CMBufferQueueInstallTrigger", "10.7", _cMBufferQueueInstallTriggerErr)
	}
	return _cMBufferQueueInstallTrigger(queue, callback, refcon, condition, time, triggerTokenOut), nil
}

// CMBufferQueueInstallTrigger installs a trigger with a callback on a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueInstallTrigger(_:callback:refcon:condition:time:triggerTokenOut:)
func CMBufferQueueInstallTrigger(queue uintptr, callback CMBufferQueueTriggerCallback, refcon uintptr, condition CMBufferQueueTriggerCondition, time CMTime, triggerTokenOut *CMBufferQueueTriggerToken) int32 {
	result, callErr := tryCMBufferQueueInstallTrigger(queue, callback, refcon, condition, time, triggerTokenOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueInstallTriggerHandler func(queue uintptr, condition CMBufferQueueTriggerCondition, time CMTime, triggerTokenOut *CMBufferQueueTriggerToken, handler unsafe.Pointer) int32
var _cMBufferQueueInstallTriggerHandlerErr error

func tryCMBufferQueueInstallTriggerHandler(queue uintptr, condition CMBufferQueueTriggerCondition, time CMTime, triggerTokenOut *CMBufferQueueTriggerToken, handler CMBufferQueueTriggerHandler) (int32, error) {
	if _cMBufferQueueInstallTriggerHandler == nil {
		return 0, symbolCallError("CMBufferQueueInstallTriggerHandler", "10.14.4", _cMBufferQueueInstallTriggerHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.IObject) { handler(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _cMBufferQueueInstallTriggerHandler(queue, condition, time, triggerTokenOut, _block0), nil
}

// CMBufferQueueInstallTriggerHandler installs a trigger with a handler on a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueInstallTriggerHandler(_:_:_:_:_:)
func CMBufferQueueInstallTriggerHandler(queue uintptr, condition CMBufferQueueTriggerCondition, time CMTime, triggerTokenOut *CMBufferQueueTriggerToken, handler CMBufferQueueTriggerHandler) int32 {
	result, callErr := tryCMBufferQueueInstallTriggerHandler(queue, condition, time, triggerTokenOut, handler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueInstallTriggerHandlerWithIntegerThreshold func(queue uintptr, condition CMBufferQueueTriggerCondition, threshold int, triggerTokenOut *CMBufferQueueTriggerToken, handler unsafe.Pointer) int32
var _cMBufferQueueInstallTriggerHandlerWithIntegerThresholdErr error

func tryCMBufferQueueInstallTriggerHandlerWithIntegerThreshold(queue uintptr, condition CMBufferQueueTriggerCondition, threshold int, triggerTokenOut *CMBufferQueueTriggerToken, handler CMBufferQueueTriggerHandler) (int32, error) {
	if _cMBufferQueueInstallTriggerHandlerWithIntegerThreshold == nil {
		return 0, symbolCallError("CMBufferQueueInstallTriggerHandlerWithIntegerThreshold", "10.14.4", _cMBufferQueueInstallTriggerHandlerWithIntegerThresholdErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.IObject) { handler(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _cMBufferQueueInstallTriggerHandlerWithIntegerThreshold(queue, condition, threshold, triggerTokenOut, _block0), nil
}

// CMBufferQueueInstallTriggerHandlerWithIntegerThreshold installs a trigger with a handler and threshold on a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueInstallTriggerHandlerWithIntegerThreshold(_:_:_:_:_:)
func CMBufferQueueInstallTriggerHandlerWithIntegerThreshold(queue uintptr, condition CMBufferQueueTriggerCondition, threshold int, triggerTokenOut *CMBufferQueueTriggerToken, handler CMBufferQueueTriggerHandler) int32 {
	result, callErr := tryCMBufferQueueInstallTriggerHandlerWithIntegerThreshold(queue, condition, threshold, triggerTokenOut, handler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueInstallTriggerWithIntegerThreshold func(queue uintptr, callback CMBufferQueueTriggerCallback, refcon uintptr, condition CMBufferQueueTriggerCondition, threshold int, triggerTokenOut *CMBufferQueueTriggerToken) int32
var _cMBufferQueueInstallTriggerWithIntegerThresholdErr error

func tryCMBufferQueueInstallTriggerWithIntegerThreshold(queue uintptr, callback CMBufferQueueTriggerCallback, refcon uintptr, condition CMBufferQueueTriggerCondition, threshold int, triggerTokenOut *CMBufferQueueTriggerToken) (int32, error) {
	if _cMBufferQueueInstallTriggerWithIntegerThreshold == nil {
		return 0, symbolCallError("CMBufferQueueInstallTriggerWithIntegerThreshold", "10.7", _cMBufferQueueInstallTriggerWithIntegerThresholdErr)
	}
	return _cMBufferQueueInstallTriggerWithIntegerThreshold(queue, callback, refcon, condition, threshold, triggerTokenOut), nil
}

// CMBufferQueueInstallTriggerWithIntegerThreshold installs a trigger with a callback and threshold on a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueInstallTriggerWithIntegerThreshold(_:callback:refcon:condition:threshold:triggerTokenOut:)
func CMBufferQueueInstallTriggerWithIntegerThreshold(queue uintptr, callback CMBufferQueueTriggerCallback, refcon uintptr, condition CMBufferQueueTriggerCondition, threshold int, triggerTokenOut *CMBufferQueueTriggerToken) int32 {
	result, callErr := tryCMBufferQueueInstallTriggerWithIntegerThreshold(queue, callback, refcon, condition, threshold, triggerTokenOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueIsAtEndOfData func(queue uintptr) bool
var _cMBufferQueueIsAtEndOfDataErr error

func tryCMBufferQueueIsAtEndOfData(queue uintptr) (bool, error) {
	if _cMBufferQueueIsAtEndOfData == nil {
		return false, symbolCallError("CMBufferQueueIsAtEndOfData", "10.7", _cMBufferQueueIsAtEndOfDataErr)
	}
	return _cMBufferQueueIsAtEndOfData(queue), nil
}

// CMBufferQueueIsAtEndOfData returns a Boolean value that indicates whether a buffer queue has its end-of-data marker set, and is now empty.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueIsAtEndOfData(_:)
func CMBufferQueueIsAtEndOfData(queue uintptr) bool {
	result, callErr := tryCMBufferQueueIsAtEndOfData(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueIsEmpty func(queue uintptr) bool
var _cMBufferQueueIsEmptyErr error

func tryCMBufferQueueIsEmpty(queue uintptr) (bool, error) {
	if _cMBufferQueueIsEmpty == nil {
		return false, symbolCallError("CMBufferQueueIsEmpty", "10.7", _cMBufferQueueIsEmptyErr)
	}
	return _cMBufferQueueIsEmpty(queue), nil
}

// CMBufferQueueIsEmpty returns a Boolean value that indicates whether a buffer queue is empty.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueIsEmpty(_:)
func CMBufferQueueIsEmpty(queue uintptr) bool {
	result, callErr := tryCMBufferQueueIsEmpty(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueMarkEndOfData func(queue uintptr) int32
var _cMBufferQueueMarkEndOfDataErr error

func tryCMBufferQueueMarkEndOfData(queue uintptr) (int32, error) {
	if _cMBufferQueueMarkEndOfData == nil {
		return 0, symbolCallError("CMBufferQueueMarkEndOfData", "10.7", _cMBufferQueueMarkEndOfDataErr)
	}
	return _cMBufferQueueMarkEndOfData(queue), nil
}

// CMBufferQueueMarkEndOfData sets a marker to indicate this queue doesn’t allow enqueuing new buffers.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueMarkEndOfData(_:)
func CMBufferQueueMarkEndOfData(queue uintptr) int32 {
	result, callErr := tryCMBufferQueueMarkEndOfData(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueRemoveTrigger func(queue uintptr, triggerToken CMBufferQueueTriggerToken) int32
var _cMBufferQueueRemoveTriggerErr error

func tryCMBufferQueueRemoveTrigger(queue uintptr, triggerToken CMBufferQueueTriggerToken) (int32, error) {
	if _cMBufferQueueRemoveTrigger == nil {
		return 0, symbolCallError("CMBufferQueueRemoveTrigger", "10.7", _cMBufferQueueRemoveTriggerErr)
	}
	return _cMBufferQueueRemoveTrigger(queue, triggerToken), nil
}

// CMBufferQueueRemoveTrigger removes a previously installed trigger from a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueRemoveTrigger(_:triggerToken:)
func CMBufferQueueRemoveTrigger(queue uintptr, triggerToken CMBufferQueueTriggerToken) int32 {
	result, callErr := tryCMBufferQueueRemoveTrigger(queue, triggerToken)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueReset func(queue uintptr) int32
var _cMBufferQueueResetErr error

func tryCMBufferQueueReset(queue uintptr) (int32, error) {
	if _cMBufferQueueReset == nil {
		return 0, symbolCallError("CMBufferQueueReset", "10.7", _cMBufferQueueResetErr)
	}
	return _cMBufferQueueReset(queue), nil
}

// CMBufferQueueReset resets a buffer queue, which allows it to enqueue new buffers.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueReset(_:)
func CMBufferQueueReset(queue uintptr) int32 {
	result, callErr := tryCMBufferQueueReset(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueResetWithCallback func(queue uintptr, refcon uintptr) int32
var _cMBufferQueueResetWithCallbackErr error

func tryCMBufferQueueResetWithCallback(queue uintptr, refcon uintptr) (int32, error) {
	if _cMBufferQueueResetWithCallback == nil {
		return 0, symbolCallError("CMBufferQueueResetWithCallback", "10.7", _cMBufferQueueResetWithCallbackErr)
	}
	return _cMBufferQueueResetWithCallback(queue, refcon), nil
}

// CMBufferQueueResetWithCallback a callback that invokes a function for every buffer in a queue and then resets the queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueResetWithCallback(_:callback:refcon:)
func CMBufferQueueResetWithCallback(queue uintptr, refcon uintptr) int32 {
	result, callErr := tryCMBufferQueueResetWithCallback(queue, refcon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueSetValidationCallback func(queue uintptr, callback CMBufferValidationCallback, refcon uintptr) int32
var _cMBufferQueueSetValidationCallbackErr error

func tryCMBufferQueueSetValidationCallback(queue uintptr, callback CMBufferValidationCallback, refcon uintptr) (int32, error) {
	if _cMBufferQueueSetValidationCallback == nil {
		return 0, symbolCallError("CMBufferQueueSetValidationCallback", "10.7", _cMBufferQueueSetValidationCallbackErr)
	}
	return _cMBufferQueueSetValidationCallback(queue, callback, refcon), nil
}

// CMBufferQueueSetValidationCallback a validation callback for the queue to call before enqueuing buffers.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueSetValidationCallback(_:callback:refcon:)
func CMBufferQueueSetValidationCallback(queue uintptr, callback CMBufferValidationCallback, refcon uintptr) int32 {
	result, callErr := tryCMBufferQueueSetValidationCallback(queue, callback, refcon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueSetValidationHandler func(queue uintptr, handler unsafe.Pointer) int32
var _cMBufferQueueSetValidationHandlerErr error

func tryCMBufferQueueSetValidationHandler(queue uintptr, handler CMBufferValidationHandler) (int32, error) {
	if _cMBufferQueueSetValidationHandler == nil {
		return 0, symbolCallError("CMBufferQueueSetValidationHandler", "10.14.4", _cMBufferQueueSetValidationHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.IObject, arg1 unsafe.Pointer) int { return handler(arg0, arg1) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _cMBufferQueueSetValidationHandler(queue, _block0), nil
}

// CMBufferQueueSetValidationHandler a validation handler for the queue to call before enqueuing buffers.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueSetValidationHandler(_:_:)
func CMBufferQueueSetValidationHandler(queue uintptr, handler CMBufferValidationHandler) int32 {
	result, callErr := tryCMBufferQueueSetValidationHandler(queue, handler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMBufferQueueTestTrigger func(queue uintptr, triggerToken CMBufferQueueTriggerToken) bool
var _cMBufferQueueTestTriggerErr error

func tryCMBufferQueueTestTrigger(queue uintptr, triggerToken CMBufferQueueTriggerToken) (bool, error) {
	if _cMBufferQueueTestTrigger == nil {
		return false, symbolCallError("CMBufferQueueTestTrigger", "10.7", _cMBufferQueueTestTriggerErr)
	}
	return _cMBufferQueueTestTrigger(queue, triggerToken), nil
}

// CMBufferQueueTestTrigger tests whether the trigger condition is true for the specified buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueTestTrigger(_:triggerToken:)
func CMBufferQueueTestTrigger(queue uintptr, triggerToken CMBufferQueueTriggerToken) bool {
	result, callErr := tryCMBufferQueueTestTrigger(queue, triggerToken)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMClockConvertHostTimeToSystemUnits func(hostTime CMTime) uint64
var _cMClockConvertHostTimeToSystemUnitsErr error

func tryCMClockConvertHostTimeToSystemUnits(hostTime CMTime) (uint64, error) {
	if _cMClockConvertHostTimeToSystemUnits == nil {
		return 0, symbolCallError("CMClockConvertHostTimeToSystemUnits", "10.8", _cMClockConvertHostTimeToSystemUnitsErr)
	}
	return _cMClockConvertHostTimeToSystemUnits(hostTime), nil
}

// CMClockConvertHostTimeToSystemUnits converts a host time from a core media time structure to the host time’s native units.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClockConvertHostTimeToSystemUnits(_:)
func CMClockConvertHostTimeToSystemUnits(hostTime CMTime) uint64 {
	result, callErr := tryCMClockConvertHostTimeToSystemUnits(hostTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMClockGetAnchorTime func(clock uintptr, clockTimeOut *CMTime, referenceClockTimeOut *CMTime) int32
var _cMClockGetAnchorTimeErr error

func tryCMClockGetAnchorTime(clock uintptr, clockTimeOut *CMTime, referenceClockTimeOut *CMTime) (int32, error) {
	if _cMClockGetAnchorTime == nil {
		return 0, symbolCallError("CMClockGetAnchorTime", "10.8", _cMClockGetAnchorTimeErr)
	}
	return _cMClockGetAnchorTime(clock, clockTimeOut, referenceClockTimeOut), nil
}

// CMClockGetAnchorTime returns the current time from a clock and the matching time from the clock’s reference clock.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClockGetAnchorTime(_:clockTimeOut:referenceClockTimeOut:)
func CMClockGetAnchorTime(clock uintptr, clockTimeOut *CMTime, referenceClockTimeOut *CMTime) int32 {
	result, callErr := tryCMClockGetAnchorTime(clock, clockTimeOut, referenceClockTimeOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMClockGetHostTimeClock func() uintptr
var _cMClockGetHostTimeClockErr error

func tryCMClockGetHostTimeClock() (uintptr, error) {
	if _cMClockGetHostTimeClock == nil {
		return 0, symbolCallError("CMClockGetHostTimeClock", "10.8", _cMClockGetHostTimeClockErr)
	}
	return _cMClockGetHostTimeClock(), nil
}

// CMClockGetHostTimeClock returns a reference to the singleton clock that reflects the host time.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClockGetHostTimeClock()
func CMClockGetHostTimeClock() uintptr {
	result, callErr := tryCMClockGetHostTimeClock()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMClockGetTime func(clock uintptr) CMTime
var _cMClockGetTimeErr error

func tryCMClockGetTime(clock uintptr) (CMTime, error) {
	if _cMClockGetTime == nil {
		return CMTime{}, symbolCallError("CMClockGetTime", "10.8", _cMClockGetTimeErr)
	}
	return _cMClockGetTime(clock), nil
}

// CMClockGetTime returns the current time from a clock.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClockGetTime(_:)
func CMClockGetTime(clock uintptr) CMTime {
	result, callErr := tryCMClockGetTime(clock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMClockGetTypeID func() uint
var _cMClockGetTypeIDErr error

func tryCMClockGetTypeID() (uint, error) {
	if _cMClockGetTypeID == nil {
		return 0, symbolCallError("CMClockGetTypeID", "10.8", _cMClockGetTypeIDErr)
	}
	return _cMClockGetTypeID(), nil
}

// CMClockGetTypeID returns the core foundation type identifier of a clock type.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClockGetTypeID()
func CMClockGetTypeID() uint {
	result, callErr := tryCMClockGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMClockInvalidate func(clock uintptr)
var _cMClockInvalidateErr error

func tryCMClockInvalidate(clock uintptr) error {
	if _cMClockInvalidate == nil {
		return symbolCallError("CMClockInvalidate", "10.8", _cMClockInvalidateErr)
	}
	_cMClockInvalidate(clock)
	return nil
}

// CMClockInvalidate stops the clock.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClockInvalidate(_:)
func CMClockInvalidate(clock uintptr) {
	if callErr := tryCMClockInvalidate(clock); callErr != nil {
		panic(callErr)
	}
}

var _cMClockMakeHostTimeFromSystemUnits func(hostTime uint64) CMTime
var _cMClockMakeHostTimeFromSystemUnitsErr error

func tryCMClockMakeHostTimeFromSystemUnits(hostTime uint64) (CMTime, error) {
	if _cMClockMakeHostTimeFromSystemUnits == nil {
		return CMTime{}, symbolCallError("CMClockMakeHostTimeFromSystemUnits", "10.8", _cMClockMakeHostTimeFromSystemUnitsErr)
	}
	return _cMClockMakeHostTimeFromSystemUnits(hostTime), nil
}

// CMClockMakeHostTimeFromSystemUnits converts a host time from native units to a core media time structure.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClockMakeHostTimeFromSystemUnits(_:)
func CMClockMakeHostTimeFromSystemUnits(hostTime uint64) CMTime {
	result, callErr := tryCMClockMakeHostTimeFromSystemUnits(hostTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMClockMightDrift func(clock uintptr, otherClock uintptr) bool
var _cMClockMightDriftErr error

func tryCMClockMightDrift(clock uintptr, otherClock uintptr) (bool, error) {
	if _cMClockMightDrift == nil {
		return false, symbolCallError("CMClockMightDrift", "10.8", _cMClockMightDriftErr)
	}
	return _cMClockMightDrift(clock, otherClock), nil
}

// CMClockMightDrift returns a Boolean value that indicates whether it’s possible for two clocks to drift relative to each other.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClockMightDrift(_:otherClock:)
func CMClockMightDrift(clock uintptr, otherClock uintptr) bool {
	result, callErr := tryCMClockMightDrift(clock, otherClock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, closedCaptionFormatDescription CMClosedCaptionFormatDescriptionRef, flavor CMClosedCaptionDescriptionFlavor, blockBufferOut *uintptr) int32
var _cMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBufferErr error

func tryCMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, closedCaptionFormatDescription CMClosedCaptionFormatDescriptionRef, flavor CMClosedCaptionDescriptionFlavor, blockBufferOut *uintptr) (int32, error) {
	if _cMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer == nil {
		return 0, symbolCallError("CMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer", "10.10", _cMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBufferErr)
	}
	return _cMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer(allocator, closedCaptionFormatDescription, flavor, blockBufferOut), nil
}

// CMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer copies the contents of a closed caption format description to a buffer in big-endian byte order.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer(allocator:closedCaptionFormatDescription:flavor:blockBufferOut:)
func CMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, closedCaptionFormatDescription CMClosedCaptionFormatDescriptionRef, flavor CMClosedCaptionDescriptionFlavor, blockBufferOut *uintptr) int32 {
	result, callErr := tryCMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer(allocator, closedCaptionFormatDescription, flavor, blockBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, closedCaptionDescriptionBlockBuffer uintptr, flavor CMClosedCaptionDescriptionFlavor, formatDescriptionOut *CMClosedCaptionFormatDescriptionRef) int32
var _cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBufferErr error

func tryCMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, closedCaptionDescriptionBlockBuffer uintptr, flavor CMClosedCaptionDescriptionFlavor, formatDescriptionOut *CMClosedCaptionFormatDescriptionRef) (int32, error) {
	if _cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer == nil {
		return 0, symbolCallError("CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer", "10.10", _cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBufferErr)
	}
	return _cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer(allocator, closedCaptionDescriptionBlockBuffer, flavor, formatDescriptionOut), nil
}

// CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer creates a closed caption format description from a big-endian closed caption description structure in a buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer(allocator:bigEndianClosedCaptionDescriptionBlockBuffer:flavor:formatDescriptionOut:)
func CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, closedCaptionDescriptionBlockBuffer uintptr, flavor CMClosedCaptionDescriptionFlavor, formatDescriptionOut *CMClosedCaptionFormatDescriptionRef) int32 {
	result, callErr := tryCMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer(allocator, closedCaptionDescriptionBlockBuffer, flavor, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData func(allocator corefoundation.CFAllocatorRef, closedCaptionDescriptionData *uint8, size uintptr, flavor CMClosedCaptionDescriptionFlavor, formatDescriptionOut *CMClosedCaptionFormatDescriptionRef) int32
var _cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionDataErr error

func tryCMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData(allocator corefoundation.CFAllocatorRef, closedCaptionDescriptionData *uint8, size uintptr, flavor CMClosedCaptionDescriptionFlavor, formatDescriptionOut *CMClosedCaptionFormatDescriptionRef) (int32, error) {
	if _cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData == nil {
		return 0, symbolCallError("CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData", "10.10", _cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionDataErr)
	}
	return _cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData(allocator, closedCaptionDescriptionData, size, flavor, formatDescriptionOut), nil
}

// CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData creates a closed caption format description from a big-endian closed caption description structure.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData(allocator:bigEndianClosedCaptionDescriptionData:size:flavor:formatDescriptionOut:)
func CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData(allocator corefoundation.CFAllocatorRef, closedCaptionDescriptionData *uint8, size uintptr, flavor CMClosedCaptionDescriptionFlavor, formatDescriptionOut *CMClosedCaptionFormatDescriptionRef) int32 {
	result, callErr := tryCMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData(allocator, closedCaptionDescriptionData, size, flavor, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMCopyDictionaryOfAttachments func(allocator corefoundation.CFAllocatorRef, target CMAttachmentBearerRef, attachmentMode CMAttachmentMode) corefoundation.CFDictionaryRef
var _cMCopyDictionaryOfAttachmentsErr error

func tryCMCopyDictionaryOfAttachments(allocator corefoundation.CFAllocatorRef, target CMAttachmentBearerRef, attachmentMode CMAttachmentMode) (corefoundation.CFDictionaryRef, error) {
	if _cMCopyDictionaryOfAttachments == nil {
		return 0, symbolCallError("CMCopyDictionaryOfAttachments", "10.7", _cMCopyDictionaryOfAttachmentsErr)
	}
	return _cMCopyDictionaryOfAttachments(allocator, target, attachmentMode), nil
}

// CMCopyDictionaryOfAttachments returns a dictionary of all attachments for an attachment bearer object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMCopyDictionaryOfAttachments(allocator:target:attachmentMode:)
func CMCopyDictionaryOfAttachments(allocator corefoundation.CFAllocatorRef, target CMAttachmentBearerRef, attachmentMode CMAttachmentMode) corefoundation.CFDictionaryRef {
	result, callErr := tryCMCopyDictionaryOfAttachments(allocator, target, attachmentMode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout func(soundDescriptionBlockBuffer uintptr, flavor CMSoundDescriptionFlavor) bool
var _cMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayoutErr error

func tryCMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout(soundDescriptionBlockBuffer uintptr, flavor CMSoundDescriptionFlavor) (bool, error) {
	if _cMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout == nil {
		return false, symbolCallError("CMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout", "10.10", _cMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayoutErr)
	}
	return _cMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout(soundDescriptionBlockBuffer, flavor), nil
}

// CMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout returns a Boolean value that indicates whether the sample tables need to use the legacy constant bit-rate encoding layout.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout(_:flavor:)
func CMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout(soundDescriptionBlockBuffer uintptr, flavor CMSoundDescriptionFlavor) bool {
	result, callErr := tryCMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout(soundDescriptionBlockBuffer, flavor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMFormatDescriptionCreate func(allocator corefoundation.CFAllocatorRef, mediaType uint32, mediaSubType uint32, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *uintptr) int32
var _cMFormatDescriptionCreateErr error

func tryCMFormatDescriptionCreate(allocator corefoundation.CFAllocatorRef, mediaType uint32, mediaSubType uint32, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *uintptr) (int32, error) {
	if _cMFormatDescriptionCreate == nil {
		return 0, symbolCallError("CMFormatDescriptionCreate", "10.7", _cMFormatDescriptionCreateErr)
	}
	return _cMFormatDescriptionCreate(allocator, mediaType, mediaSubType, extensions, formatDescriptionOut), nil
}

// CMFormatDescriptionCreate creates a format description for general use.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionCreate(allocator:mediaType:mediaSubType:extensions:formatDescriptionOut:)
func CMFormatDescriptionCreate(allocator corefoundation.CFAllocatorRef, mediaType uint32, mediaSubType uint32, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *uintptr) int32 {
	result, callErr := tryCMFormatDescriptionCreate(allocator, mediaType, mediaSubType, extensions, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMFormatDescriptionEqual func(formatDescription uintptr, otherFormatDescription uintptr) bool
var _cMFormatDescriptionEqualErr error

func tryCMFormatDescriptionEqual(formatDescription uintptr, otherFormatDescription uintptr) (bool, error) {
	if _cMFormatDescriptionEqual == nil {
		return false, symbolCallError("CMFormatDescriptionEqual", "10.7", _cMFormatDescriptionEqualErr)
	}
	return _cMFormatDescriptionEqual(formatDescription, otherFormatDescription), nil
}

// CMFormatDescriptionEqual returns a Boolean value that indicates whether two format descriptions are equal.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionEqual(_:otherFormatDescription:)
func CMFormatDescriptionEqual(formatDescription uintptr, otherFormatDescription uintptr) bool {
	result, callErr := tryCMFormatDescriptionEqual(formatDescription, otherFormatDescription)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMFormatDescriptionEqualIgnoringExtensionKeys func(formatDescription uintptr, otherFormatDescription uintptr, formatDescriptionExtensionKeysToIgnore corefoundation.CFTypeRef, sampleDescriptionExtensionAtomKeysToIgnore corefoundation.CFTypeRef) bool
var _cMFormatDescriptionEqualIgnoringExtensionKeysErr error

func tryCMFormatDescriptionEqualIgnoringExtensionKeys(formatDescription uintptr, otherFormatDescription uintptr, formatDescriptionExtensionKeysToIgnore corefoundation.CFTypeRef, sampleDescriptionExtensionAtomKeysToIgnore corefoundation.CFTypeRef) (bool, error) {
	if _cMFormatDescriptionEqualIgnoringExtensionKeys == nil {
		return false, symbolCallError("CMFormatDescriptionEqualIgnoringExtensionKeys", "10.7", _cMFormatDescriptionEqualIgnoringExtensionKeysErr)
	}
	return _cMFormatDescriptionEqualIgnoringExtensionKeys(formatDescription, otherFormatDescription, formatDescriptionExtensionKeysToIgnore, sampleDescriptionExtensionAtomKeysToIgnore), nil
}

// CMFormatDescriptionEqualIgnoringExtensionKeys returns a Boolean value that indicates whether two format descriptions are equal, ignoring differences in the extension keys you specify.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionEqualIgnoringExtensionKeys(_:otherFormatDescription:extensionKeysToIgnore:sampleDescriptionExtensionAtomKeysToIgnore:)
func CMFormatDescriptionEqualIgnoringExtensionKeys(formatDescription uintptr, otherFormatDescription uintptr, formatDescriptionExtensionKeysToIgnore corefoundation.CFTypeRef, sampleDescriptionExtensionAtomKeysToIgnore corefoundation.CFTypeRef) bool {
	result, callErr := tryCMFormatDescriptionEqualIgnoringExtensionKeys(formatDescription, otherFormatDescription, formatDescriptionExtensionKeysToIgnore, sampleDescriptionExtensionAtomKeysToIgnore)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMFormatDescriptionGetExtension func(desc uintptr, extensionKey corefoundation.CFStringRef) corefoundation.CFPropertyListRef
var _cMFormatDescriptionGetExtensionErr error

func tryCMFormatDescriptionGetExtension(desc uintptr, extensionKey corefoundation.CFStringRef) (corefoundation.CFPropertyListRef, error) {
	if _cMFormatDescriptionGetExtension == nil {
		return 0, symbolCallError("CMFormatDescriptionGetExtension", "10.7", _cMFormatDescriptionGetExtensionErr)
	}
	return _cMFormatDescriptionGetExtension(desc, extensionKey), nil
}

// CMFormatDescriptionGetExtension returns an extension from the format description by using an extension key.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetExtension(_:extensionKey:)
func CMFormatDescriptionGetExtension(desc uintptr, extensionKey corefoundation.CFStringRef) corefoundation.CFPropertyListRef {
	result, callErr := tryCMFormatDescriptionGetExtension(desc, extensionKey)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMFormatDescriptionGetExtensions func(desc uintptr) corefoundation.CFDictionaryRef
var _cMFormatDescriptionGetExtensionsErr error

func tryCMFormatDescriptionGetExtensions(desc uintptr) (corefoundation.CFDictionaryRef, error) {
	if _cMFormatDescriptionGetExtensions == nil {
		return 0, symbolCallError("CMFormatDescriptionGetExtensions", "10.7", _cMFormatDescriptionGetExtensionsErr)
	}
	return _cMFormatDescriptionGetExtensions(desc), nil
}

// CMFormatDescriptionGetExtensions returns all of the extensions for a format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetExtensions(_:)
func CMFormatDescriptionGetExtensions(desc uintptr) corefoundation.CFDictionaryRef {
	result, callErr := tryCMFormatDescriptionGetExtensions(desc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMFormatDescriptionGetMediaSubType func(desc uintptr) uint32
var _cMFormatDescriptionGetMediaSubTypeErr error

func tryCMFormatDescriptionGetMediaSubType(desc uintptr) (uint32, error) {
	if _cMFormatDescriptionGetMediaSubType == nil {
		return 0, symbolCallError("CMFormatDescriptionGetMediaSubType", "10.7", _cMFormatDescriptionGetMediaSubTypeErr)
	}
	return _cMFormatDescriptionGetMediaSubType(desc), nil
}

// CMFormatDescriptionGetMediaSubType returns the media subtype of a format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetMediaSubType(_:)
func CMFormatDescriptionGetMediaSubType(desc uintptr) uint32 {
	result, callErr := tryCMFormatDescriptionGetMediaSubType(desc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMFormatDescriptionGetMediaType func(desc uintptr) uint32
var _cMFormatDescriptionGetMediaTypeErr error

func tryCMFormatDescriptionGetMediaType(desc uintptr) (uint32, error) {
	if _cMFormatDescriptionGetMediaType == nil {
		return 0, symbolCallError("CMFormatDescriptionGetMediaType", "10.7", _cMFormatDescriptionGetMediaTypeErr)
	}
	return _cMFormatDescriptionGetMediaType(desc), nil
}

// CMFormatDescriptionGetMediaType returns the media type of a format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetMediaType(_:)
func CMFormatDescriptionGetMediaType(desc uintptr) uint32 {
	result, callErr := tryCMFormatDescriptionGetMediaType(desc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMFormatDescriptionGetTypeID func() uint
var _cMFormatDescriptionGetTypeIDErr error

func tryCMFormatDescriptionGetTypeID() (uint, error) {
	if _cMFormatDescriptionGetTypeID == nil {
		return 0, symbolCallError("CMFormatDescriptionGetTypeID", "10.7", _cMFormatDescriptionGetTypeIDErr)
	}
	return _cMFormatDescriptionGetTypeID(), nil
}

// CMFormatDescriptionGetTypeID returns the Core Foundation type identifier that identifies format description objects.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetTypeID()
func CMFormatDescriptionGetTypeID() uint {
	result, callErr := tryCMFormatDescriptionGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMGetAttachment func(target CMAttachmentBearerRef, key corefoundation.CFStringRef, attachmentModeOut *CMAttachmentMode) corefoundation.CFTypeRef
var _cMGetAttachmentErr error

func tryCMGetAttachment(target CMAttachmentBearerRef, key corefoundation.CFStringRef, attachmentModeOut *CMAttachmentMode) (corefoundation.CFTypeRef, error) {
	if _cMGetAttachment == nil {
		return 0, symbolCallError("CMGetAttachment", "10.7", _cMGetAttachmentErr)
	}
	return _cMGetAttachment(target, key, attachmentModeOut), nil
}

// CMGetAttachment returns an attachment from an attachment bearer object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMGetAttachment(_:key:attachmentModeOut:)
func CMGetAttachment(target CMAttachmentBearerRef, key corefoundation.CFStringRef, attachmentModeOut *CMAttachmentMode) corefoundation.CFTypeRef {
	result, callErr := tryCMGetAttachment(target, key, attachmentModeOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMemoryPoolCreate func(options corefoundation.CFDictionaryRef) uintptr
var _cMMemoryPoolCreateErr error

func tryCMMemoryPoolCreate(options corefoundation.CFDictionaryRef) (uintptr, error) {
	if _cMMemoryPoolCreate == nil {
		return 0, symbolCallError("CMMemoryPoolCreate", "10.8", _cMMemoryPoolCreateErr)
	}
	return _cMMemoryPoolCreate(options), nil
}

// CMMemoryPoolCreate creates a memory pool.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolCreate(options:)
func CMMemoryPoolCreate(options corefoundation.CFDictionaryRef) uintptr {
	result, callErr := tryCMMemoryPoolCreate(options)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMemoryPoolFlush func(pool uintptr)
var _cMMemoryPoolFlushErr error

func tryCMMemoryPoolFlush(pool uintptr) error {
	if _cMMemoryPoolFlush == nil {
		return symbolCallError("CMMemoryPoolFlush", "10.8", _cMMemoryPoolFlushErr)
	}
	_cMMemoryPoolFlush(pool)
	return nil
}

// CMMemoryPoolFlush deallocates all memory the pool holds.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolFlush(_:)
func CMMemoryPoolFlush(pool uintptr) {
	if callErr := tryCMMemoryPoolFlush(pool); callErr != nil {
		panic(callErr)
	}
}

var _cMMemoryPoolGetAllocator func(pool uintptr) corefoundation.CFAllocatorRef
var _cMMemoryPoolGetAllocatorErr error

func tryCMMemoryPoolGetAllocator(pool uintptr) (corefoundation.CFAllocatorRef, error) {
	if _cMMemoryPoolGetAllocator == nil {
		return 0, symbolCallError("CMMemoryPoolGetAllocator", "10.8", _cMMemoryPoolGetAllocatorErr)
	}
	return _cMMemoryPoolGetAllocator(pool), nil
}

// CMMemoryPoolGetAllocator returns the allocator for the memory pool.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolGetAllocator(_:)
func CMMemoryPoolGetAllocator(pool uintptr) corefoundation.CFAllocatorRef {
	result, callErr := tryCMMemoryPoolGetAllocator(pool)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMemoryPoolGetTypeID func() uint
var _cMMemoryPoolGetTypeIDErr error

func tryCMMemoryPoolGetTypeID() (uint, error) {
	if _cMMemoryPoolGetTypeID == nil {
		return 0, symbolCallError("CMMemoryPoolGetTypeID", "10.8", _cMMemoryPoolGetTypeIDErr)
	}
	return _cMMemoryPoolGetTypeID(), nil
}

// CMMemoryPoolGetTypeID returns the type identifier of memory pool objects.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolGetTypeID()
func CMMemoryPoolGetTypeID() uint {
	result, callErr := tryCMMemoryPoolGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMemoryPoolInvalidate func(pool uintptr)
var _cMMemoryPoolInvalidateErr error

func tryCMMemoryPoolInvalidate(pool uintptr) error {
	if _cMMemoryPoolInvalidate == nil {
		return symbolCallError("CMMemoryPoolInvalidate", "10.8", _cMMemoryPoolInvalidateErr)
	}
	_cMMemoryPoolInvalidate(pool)
	return nil
}

// CMMemoryPoolInvalidate invalidates the memory pool, which causes its allocator to stop recycling memory.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolInvalidate(_:)
func CMMemoryPoolInvalidate(pool uintptr) {
	if callErr := tryCMMemoryPoolInvalidate(pool); callErr != nil {
		panic(callErr)
	}
}

var _cMMetadataCreateIdentifierForKeyAndKeySpace func(allocator corefoundation.CFAllocatorRef, key corefoundation.CFTypeRef, keySpace corefoundation.CFStringRef, identifierOut *corefoundation.CFStringRef) int32
var _cMMetadataCreateIdentifierForKeyAndKeySpaceErr error

func tryCMMetadataCreateIdentifierForKeyAndKeySpace(allocator corefoundation.CFAllocatorRef, key corefoundation.CFTypeRef, keySpace corefoundation.CFStringRef, identifierOut *corefoundation.CFStringRef) (int32, error) {
	if _cMMetadataCreateIdentifierForKeyAndKeySpace == nil {
		return 0, symbolCallError("CMMetadataCreateIdentifierForKeyAndKeySpace", "10.10", _cMMetadataCreateIdentifierForKeyAndKeySpaceErr)
	}
	return _cMMetadataCreateIdentifierForKeyAndKeySpace(allocator, key, keySpace, identifierOut), nil
}

// CMMetadataCreateIdentifierForKeyAndKeySpace creates a URL-like string identifier that represents a key or keyspace tuple.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataCreateIdentifierForKeyAndKeySpace(allocator:key:keySpace:identifierOut:)
func CMMetadataCreateIdentifierForKeyAndKeySpace(allocator corefoundation.CFAllocatorRef, key corefoundation.CFTypeRef, keySpace corefoundation.CFStringRef, identifierOut *corefoundation.CFStringRef) int32 {
	result, callErr := tryCMMetadataCreateIdentifierForKeyAndKeySpace(allocator, key, keySpace, identifierOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMetadataCreateKeyFromIdentifier func(allocator corefoundation.CFAllocatorRef, identifier corefoundation.CFStringRef, keyOut *corefoundation.CFTypeRef) int32
var _cMMetadataCreateKeyFromIdentifierErr error

func tryCMMetadataCreateKeyFromIdentifier(allocator corefoundation.CFAllocatorRef, identifier corefoundation.CFStringRef, keyOut *corefoundation.CFTypeRef) (int32, error) {
	if _cMMetadataCreateKeyFromIdentifier == nil {
		return 0, symbolCallError("CMMetadataCreateKeyFromIdentifier", "10.10", _cMMetadataCreateKeyFromIdentifierErr)
	}
	return _cMMetadataCreateKeyFromIdentifier(allocator, identifier, keyOut), nil
}

// CMMetadataCreateKeyFromIdentifier creates a copy of the key by using an identifier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataCreateKeyFromIdentifier(allocator:identifier:keyOut:)
func CMMetadataCreateKeyFromIdentifier(allocator corefoundation.CFAllocatorRef, identifier corefoundation.CFStringRef, keyOut *corefoundation.CFTypeRef) int32 {
	result, callErr := tryCMMetadataCreateKeyFromIdentifier(allocator, identifier, keyOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMetadataCreateKeyFromIdentifierAsCFData func(allocator corefoundation.CFAllocatorRef, identifier corefoundation.CFStringRef, keyOut *corefoundation.CFDataRef) int32
var _cMMetadataCreateKeyFromIdentifierAsCFDataErr error

func tryCMMetadataCreateKeyFromIdentifierAsCFData(allocator corefoundation.CFAllocatorRef, identifier corefoundation.CFStringRef, keyOut *corefoundation.CFDataRef) (int32, error) {
	if _cMMetadataCreateKeyFromIdentifierAsCFData == nil {
		return 0, symbolCallError("CMMetadataCreateKeyFromIdentifierAsCFData", "10.10", _cMMetadataCreateKeyFromIdentifierAsCFDataErr)
	}
	return _cMMetadataCreateKeyFromIdentifierAsCFData(allocator, identifier, keyOut), nil
}

// CMMetadataCreateKeyFromIdentifierAsCFData creates a copy of the key by using an identifier, and results in a core foundation data object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataCreateKeyFromIdentifierAsCFData(allocator:identifier:keyOut:)
func CMMetadataCreateKeyFromIdentifierAsCFData(allocator corefoundation.CFAllocatorRef, identifier corefoundation.CFStringRef, keyOut *corefoundation.CFDataRef) int32 {
	result, callErr := tryCMMetadataCreateKeyFromIdentifierAsCFData(allocator, identifier, keyOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMetadataCreateKeySpaceFromIdentifier func(allocator corefoundation.CFAllocatorRef, identifier corefoundation.CFStringRef, keySpaceOut *corefoundation.CFStringRef) int32
var _cMMetadataCreateKeySpaceFromIdentifierErr error

func tryCMMetadataCreateKeySpaceFromIdentifier(allocator corefoundation.CFAllocatorRef, identifier corefoundation.CFStringRef, keySpaceOut *corefoundation.CFStringRef) (int32, error) {
	if _cMMetadataCreateKeySpaceFromIdentifier == nil {
		return 0, symbolCallError("CMMetadataCreateKeySpaceFromIdentifier", "10.10", _cMMetadataCreateKeySpaceFromIdentifierErr)
	}
	return _cMMetadataCreateKeySpaceFromIdentifier(allocator, identifier, keySpaceOut), nil
}

// CMMetadataCreateKeySpaceFromIdentifier creates a copy of the keyspace by using an identifier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataCreateKeySpaceFromIdentifier(allocator:identifier:keySpaceOut:)
func CMMetadataCreateKeySpaceFromIdentifier(allocator corefoundation.CFAllocatorRef, identifier corefoundation.CFStringRef, keySpaceOut *corefoundation.CFStringRef) int32 {
	result, callErr := tryCMMetadataCreateKeySpaceFromIdentifier(allocator, identifier, keySpaceOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMetadataDataTypeRegistryDataTypeConformsToDataType func(dataType corefoundation.CFStringRef, conformsToDataType corefoundation.CFStringRef) bool
var _cMMetadataDataTypeRegistryDataTypeConformsToDataTypeErr error

func tryCMMetadataDataTypeRegistryDataTypeConformsToDataType(dataType corefoundation.CFStringRef, conformsToDataType corefoundation.CFStringRef) (bool, error) {
	if _cMMetadataDataTypeRegistryDataTypeConformsToDataType == nil {
		return false, symbolCallError("CMMetadataDataTypeRegistryDataTypeConformsToDataType", "10.10", _cMMetadataDataTypeRegistryDataTypeConformsToDataTypeErr)
	}
	return _cMMetadataDataTypeRegistryDataTypeConformsToDataType(dataType, conformsToDataType), nil
}

// CMMetadataDataTypeRegistryDataTypeConformsToDataType returns a Boolean value that indicates whether a data type conforms to another data type.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryDataTypeConformsToDataType(_:conformsTo:)
func CMMetadataDataTypeRegistryDataTypeConformsToDataType(dataType corefoundation.CFStringRef, conformsToDataType corefoundation.CFStringRef) bool {
	result, callErr := tryCMMetadataDataTypeRegistryDataTypeConformsToDataType(dataType, conformsToDataType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMetadataDataTypeRegistryDataTypeIsBaseDataType func(dataType corefoundation.CFStringRef) bool
var _cMMetadataDataTypeRegistryDataTypeIsBaseDataTypeErr error

func tryCMMetadataDataTypeRegistryDataTypeIsBaseDataType(dataType corefoundation.CFStringRef) (bool, error) {
	if _cMMetadataDataTypeRegistryDataTypeIsBaseDataType == nil {
		return false, symbolCallError("CMMetadataDataTypeRegistryDataTypeIsBaseDataType", "10.10", _cMMetadataDataTypeRegistryDataTypeIsBaseDataTypeErr)
	}
	return _cMMetadataDataTypeRegistryDataTypeIsBaseDataType(dataType), nil
}

// CMMetadataDataTypeRegistryDataTypeIsBaseDataType returns a Boolean value that indicates whether a data type identifier represents a base data type.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryDataTypeIsBaseDataType(_:)
func CMMetadataDataTypeRegistryDataTypeIsBaseDataType(dataType corefoundation.CFStringRef) bool {
	result, callErr := tryCMMetadataDataTypeRegistryDataTypeIsBaseDataType(dataType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMetadataDataTypeRegistryDataTypeIsRegistered func(dataType corefoundation.CFStringRef) bool
var _cMMetadataDataTypeRegistryDataTypeIsRegisteredErr error

func tryCMMetadataDataTypeRegistryDataTypeIsRegistered(dataType corefoundation.CFStringRef) (bool, error) {
	if _cMMetadataDataTypeRegistryDataTypeIsRegistered == nil {
		return false, symbolCallError("CMMetadataDataTypeRegistryDataTypeIsRegistered", "10.10", _cMMetadataDataTypeRegistryDataTypeIsRegisteredErr)
	}
	return _cMMetadataDataTypeRegistryDataTypeIsRegistered(dataType), nil
}

// CMMetadataDataTypeRegistryDataTypeIsRegistered returns a Boolean value that indicates the registration status of a data type identifier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryDataTypeIsRegistered(_:)
func CMMetadataDataTypeRegistryDataTypeIsRegistered(dataType corefoundation.CFStringRef) bool {
	result, callErr := tryCMMetadataDataTypeRegistryDataTypeIsRegistered(dataType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType func(dataType corefoundation.CFStringRef) corefoundation.CFStringRef
var _cMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataTypeErr error

func tryCMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType(dataType corefoundation.CFStringRef) (corefoundation.CFStringRef, error) {
	if _cMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType == nil {
		return 0, symbolCallError("CMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType", "10.10", _cMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataTypeErr)
	}
	return _cMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType(dataType), nil
}

// CMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType returns the base data type identifier that a data type conforms to.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType(_:)
func CMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType(dataType corefoundation.CFStringRef) corefoundation.CFStringRef {
	result, callErr := tryCMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType(dataType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMetadataDataTypeRegistryGetBaseDataTypes func() corefoundation.CFArrayRef
var _cMMetadataDataTypeRegistryGetBaseDataTypesErr error

func tryCMMetadataDataTypeRegistryGetBaseDataTypes() (corefoundation.CFArrayRef, error) {
	if _cMMetadataDataTypeRegistryGetBaseDataTypes == nil {
		return 0, symbolCallError("CMMetadataDataTypeRegistryGetBaseDataTypes", "10.10", _cMMetadataDataTypeRegistryGetBaseDataTypesErr)
	}
	return _cMMetadataDataTypeRegistryGetBaseDataTypes(), nil
}

// CMMetadataDataTypeRegistryGetBaseDataTypes returns an array of base data type identifiers.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryGetBaseDataTypes()
func CMMetadataDataTypeRegistryGetBaseDataTypes() corefoundation.CFArrayRef {
	result, callErr := tryCMMetadataDataTypeRegistryGetBaseDataTypes()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMetadataDataTypeRegistryGetConformingDataTypes func(dataType corefoundation.CFStringRef) corefoundation.CFArrayRef
var _cMMetadataDataTypeRegistryGetConformingDataTypesErr error

func tryCMMetadataDataTypeRegistryGetConformingDataTypes(dataType corefoundation.CFStringRef) (corefoundation.CFArrayRef, error) {
	if _cMMetadataDataTypeRegistryGetConformingDataTypes == nil {
		return 0, symbolCallError("CMMetadataDataTypeRegistryGetConformingDataTypes", "10.10", _cMMetadataDataTypeRegistryGetConformingDataTypesErr)
	}
	return _cMMetadataDataTypeRegistryGetConformingDataTypes(dataType), nil
}

// CMMetadataDataTypeRegistryGetConformingDataTypes returns the conforming data types for the data type, if any.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryGetConformingDataTypes(_:)
func CMMetadataDataTypeRegistryGetConformingDataTypes(dataType corefoundation.CFStringRef) corefoundation.CFArrayRef {
	result, callErr := tryCMMetadataDataTypeRegistryGetConformingDataTypes(dataType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMetadataDataTypeRegistryGetDataTypeDescription func(dataType corefoundation.CFStringRef) corefoundation.CFStringRef
var _cMMetadataDataTypeRegistryGetDataTypeDescriptionErr error

func tryCMMetadataDataTypeRegistryGetDataTypeDescription(dataType corefoundation.CFStringRef) (corefoundation.CFStringRef, error) {
	if _cMMetadataDataTypeRegistryGetDataTypeDescription == nil {
		return 0, symbolCallError("CMMetadataDataTypeRegistryGetDataTypeDescription", "10.10", _cMMetadataDataTypeRegistryGetDataTypeDescriptionErr)
	}
	return _cMMetadataDataTypeRegistryGetDataTypeDescription(dataType), nil
}

// CMMetadataDataTypeRegistryGetDataTypeDescription returns the data type description if it exists.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryGetDataTypeDescription(_:)
func CMMetadataDataTypeRegistryGetDataTypeDescription(dataType corefoundation.CFStringRef) corefoundation.CFStringRef {
	result, callErr := tryCMMetadataDataTypeRegistryGetDataTypeDescription(dataType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMetadataDataTypeRegistryRegisterDataType func(dataType corefoundation.CFStringRef, description corefoundation.CFStringRef, conformingDataTypes corefoundation.CFArrayRef) int32
var _cMMetadataDataTypeRegistryRegisterDataTypeErr error

func tryCMMetadataDataTypeRegistryRegisterDataType(dataType corefoundation.CFStringRef, description corefoundation.CFStringRef, conformingDataTypes corefoundation.CFArrayRef) (int32, error) {
	if _cMMetadataDataTypeRegistryRegisterDataType == nil {
		return 0, symbolCallError("CMMetadataDataTypeRegistryRegisterDataType", "10.10", _cMMetadataDataTypeRegistryRegisterDataTypeErr)
	}
	return _cMMetadataDataTypeRegistryRegisterDataType(dataType, description, conformingDataTypes), nil
}

// CMMetadataDataTypeRegistryRegisterDataType register a data type with the data type registry.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryRegisterDataType(_:description:conformingDataTypes:)
func CMMetadataDataTypeRegistryRegisterDataType(dataType corefoundation.CFStringRef, description corefoundation.CFStringRef, conformingDataTypes corefoundation.CFArrayRef) int32 {
	result, callErr := tryCMMetadataDataTypeRegistryRegisterDataType(dataType, description, conformingDataTypes)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, metadataFormatDescription CMMetadataFormatDescriptionRef, flavor CMMetadataDescriptionFlavor, blockBufferOut *uintptr) int32
var _cMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBufferErr error

func tryCMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, metadataFormatDescription CMMetadataFormatDescriptionRef, flavor CMMetadataDescriptionFlavor, blockBufferOut *uintptr) (int32, error) {
	if _cMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer == nil {
		return 0, symbolCallError("CMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer", "10.10", _cMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBufferErr)
	}
	return _cMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer(allocator, metadataFormatDescription, flavor, blockBufferOut), nil
}

// CMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer copies the contents of a metadata format description to a buffer in big-endian byte order.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer(allocator:metadataFormatDescription:flavor:blockBufferOut:)
func CMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, metadataFormatDescription CMMetadataFormatDescriptionRef, flavor CMMetadataDescriptionFlavor, blockBufferOut *uintptr) int32 {
	result, callErr := tryCMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer(allocator, metadataFormatDescription, flavor, blockBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions func(allocator corefoundation.CFAllocatorRef, sourceDescription CMMetadataFormatDescriptionRef, otherSourceDescription CMMetadataFormatDescriptionRef, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32
var _cMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptionsErr error

func tryCMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions(allocator corefoundation.CFAllocatorRef, sourceDescription CMMetadataFormatDescriptionRef, otherSourceDescription CMMetadataFormatDescriptionRef, formatDescriptionOut *CMMetadataFormatDescriptionRef) (int32, error) {
	if _cMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions == nil {
		return 0, symbolCallError("CMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions", "10.10", _cMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptionsErr)
	}
	return _cMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions(allocator, sourceDescription, otherSourceDescription, formatDescriptionOut), nil
}

// CMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions creates a metadata format description object by merging with another description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions(allocator:sourceDescription:otherSourceDescription:formatDescriptionOut:)
func CMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions(allocator corefoundation.CFAllocatorRef, sourceDescription CMMetadataFormatDescriptionRef, otherSourceDescription CMMetadataFormatDescriptionRef, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32 {
	result, callErr := tryCMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions(allocator, sourceDescription, otherSourceDescription, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, metadataDescriptionBlockBuffer uintptr, flavor CMMetadataDescriptionFlavor, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32
var _cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBufferErr error

func tryCMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, metadataDescriptionBlockBuffer uintptr, flavor CMMetadataDescriptionFlavor, formatDescriptionOut *CMMetadataFormatDescriptionRef) (int32, error) {
	if _cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer == nil {
		return 0, symbolCallError("CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer", "10.10", _cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBufferErr)
	}
	return _cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer(allocator, metadataDescriptionBlockBuffer, flavor, formatDescriptionOut), nil
}

// CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer creates a metadata format description from a big-endian metadata description structure inside a buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer(allocator:bigEndianMetadataDescriptionBlockBuffer:flavor:formatDescriptionOut:)
func CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, metadataDescriptionBlockBuffer uintptr, flavor CMMetadataDescriptionFlavor, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32 {
	result, callErr := tryCMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer(allocator, metadataDescriptionBlockBuffer, flavor, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData func(allocator corefoundation.CFAllocatorRef, metadataDescriptionData *uint8, size uintptr, flavor CMMetadataDescriptionFlavor, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32
var _cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionDataErr error

func tryCMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData(allocator corefoundation.CFAllocatorRef, metadataDescriptionData *uint8, size uintptr, flavor CMMetadataDescriptionFlavor, formatDescriptionOut *CMMetadataFormatDescriptionRef) (int32, error) {
	if _cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData == nil {
		return 0, symbolCallError("CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData", "10.10", _cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionDataErr)
	}
	return _cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData(allocator, metadataDescriptionData, size, flavor, formatDescriptionOut), nil
}

// CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData creates a metadata format description from a big-endian metadata description structure.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData(allocator:bigEndianMetadataDescriptionData:size:flavor:formatDescriptionOut:)
func CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData(allocator corefoundation.CFAllocatorRef, metadataDescriptionData *uint8, size uintptr, flavor CMMetadataDescriptionFlavor, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32 {
	result, callErr := tryCMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData(allocator, metadataDescriptionData, size, flavor, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMetadataFormatDescriptionCreateWithKeys func(allocator corefoundation.CFAllocatorRef, metadataType CMMetadataFormatType, keys corefoundation.CFArrayRef, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32
var _cMMetadataFormatDescriptionCreateWithKeysErr error

func tryCMMetadataFormatDescriptionCreateWithKeys(allocator corefoundation.CFAllocatorRef, metadataType CMMetadataFormatType, keys corefoundation.CFArrayRef, formatDescriptionOut *CMMetadataFormatDescriptionRef) (int32, error) {
	if _cMMetadataFormatDescriptionCreateWithKeys == nil {
		return 0, symbolCallError("CMMetadataFormatDescriptionCreateWithKeys", "10.7", _cMMetadataFormatDescriptionCreateWithKeysErr)
	}
	return _cMMetadataFormatDescriptionCreateWithKeys(allocator, metadataType, keys, formatDescriptionOut), nil
}

// CMMetadataFormatDescriptionCreateWithKeys creates a metadata format description with the metadata keys you specify.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCreateWithKeys(allocator:metadataType:keys:formatDescriptionOut:)
func CMMetadataFormatDescriptionCreateWithKeys(allocator corefoundation.CFAllocatorRef, metadataType CMMetadataFormatType, keys corefoundation.CFArrayRef, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32 {
	result, callErr := tryCMMetadataFormatDescriptionCreateWithKeys(allocator, metadataType, keys, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications func(allocator corefoundation.CFAllocatorRef, sourceDescription CMMetadataFormatDescriptionRef, metadataSpecifications corefoundation.CFArrayRef, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32
var _cMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecificationsErr error

func tryCMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications(allocator corefoundation.CFAllocatorRef, sourceDescription CMMetadataFormatDescriptionRef, metadataSpecifications corefoundation.CFArrayRef, formatDescriptionOut *CMMetadataFormatDescriptionRef) (int32, error) {
	if _cMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications == nil {
		return 0, symbolCallError("CMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications", "10.10", _cMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecificationsErr)
	}
	return _cMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications(allocator, sourceDescription, metadataSpecifications, formatDescriptionOut), nil
}

// CMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications creates a metadata format description by extending an existing description with the values you specify.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications(allocator:sourceDescription:metadataSpecifications:formatDescriptionOut:)
func CMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications(allocator corefoundation.CFAllocatorRef, sourceDescription CMMetadataFormatDescriptionRef, metadataSpecifications corefoundation.CFArrayRef, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32 {
	result, callErr := tryCMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications(allocator, sourceDescription, metadataSpecifications, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMetadataFormatDescriptionCreateWithMetadataSpecifications func(allocator corefoundation.CFAllocatorRef, metadataType CMMetadataFormatType, metadataSpecifications corefoundation.CFArrayRef, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32
var _cMMetadataFormatDescriptionCreateWithMetadataSpecificationsErr error

func tryCMMetadataFormatDescriptionCreateWithMetadataSpecifications(allocator corefoundation.CFAllocatorRef, metadataType CMMetadataFormatType, metadataSpecifications corefoundation.CFArrayRef, formatDescriptionOut *CMMetadataFormatDescriptionRef) (int32, error) {
	if _cMMetadataFormatDescriptionCreateWithMetadataSpecifications == nil {
		return 0, symbolCallError("CMMetadataFormatDescriptionCreateWithMetadataSpecifications", "10.10", _cMMetadataFormatDescriptionCreateWithMetadataSpecificationsErr)
	}
	return _cMMetadataFormatDescriptionCreateWithMetadataSpecifications(allocator, metadataType, metadataSpecifications, formatDescriptionOut), nil
}

// CMMetadataFormatDescriptionCreateWithMetadataSpecifications creates a metadata format description with the specifications you specify.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCreateWithMetadataSpecifications(allocator:metadataType:metadataSpecifications:formatDescriptionOut:)
func CMMetadataFormatDescriptionCreateWithMetadataSpecifications(allocator corefoundation.CFAllocatorRef, metadataType CMMetadataFormatType, metadataSpecifications corefoundation.CFArrayRef, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32 {
	result, callErr := tryCMMetadataFormatDescriptionCreateWithMetadataSpecifications(allocator, metadataType, metadataSpecifications, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMetadataFormatDescriptionGetIdentifiers func(desc CMMetadataFormatDescriptionRef) corefoundation.CFArrayRef
var _cMMetadataFormatDescriptionGetIdentifiersErr error

func tryCMMetadataFormatDescriptionGetIdentifiers(desc CMMetadataFormatDescriptionRef) (corefoundation.CFArrayRef, error) {
	if _cMMetadataFormatDescriptionGetIdentifiers == nil {
		return 0, symbolCallError("CMMetadataFormatDescriptionGetIdentifiers", "10.10", _cMMetadataFormatDescriptionGetIdentifiersErr)
	}
	return _cMMetadataFormatDescriptionGetIdentifiers(desc), nil
}

// CMMetadataFormatDescriptionGetIdentifiers returns an array of metadata identifiers from a metadata format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionGetIdentifiers(_:)
func CMMetadataFormatDescriptionGetIdentifiers(desc CMMetadataFormatDescriptionRef) corefoundation.CFArrayRef {
	result, callErr := tryCMMetadataFormatDescriptionGetIdentifiers(desc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMetadataFormatDescriptionGetKeyWithLocalID func(desc CMMetadataFormatDescriptionRef, localKeyID uint32) corefoundation.CFDictionaryRef
var _cMMetadataFormatDescriptionGetKeyWithLocalIDErr error

func tryCMMetadataFormatDescriptionGetKeyWithLocalID(desc CMMetadataFormatDescriptionRef, localKeyID uint32) (corefoundation.CFDictionaryRef, error) {
	if _cMMetadataFormatDescriptionGetKeyWithLocalID == nil {
		return 0, symbolCallError("CMMetadataFormatDescriptionGetKeyWithLocalID", "10.7", _cMMetadataFormatDescriptionGetKeyWithLocalIDErr)
	}
	return _cMMetadataFormatDescriptionGetKeyWithLocalID(desc, localKeyID), nil
}

// CMMetadataFormatDescriptionGetKeyWithLocalID returns the key for the local identifier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionGetKeyWithLocalID(_:localKeyID:)
func CMMetadataFormatDescriptionGetKeyWithLocalID(desc CMMetadataFormatDescriptionRef, localKeyID uint32) corefoundation.CFDictionaryRef {
	result, callErr := tryCMMetadataFormatDescriptionGetKeyWithLocalID(desc, localKeyID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMMuxedFormatDescriptionCreate func(allocator corefoundation.CFAllocatorRef, muxType CMMuxedStreamType, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *CMMuxedFormatDescriptionRef) int32
var _cMMuxedFormatDescriptionCreateErr error

func tryCMMuxedFormatDescriptionCreate(allocator corefoundation.CFAllocatorRef, muxType CMMuxedStreamType, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *CMMuxedFormatDescriptionRef) (int32, error) {
	if _cMMuxedFormatDescriptionCreate == nil {
		return 0, symbolCallError("CMMuxedFormatDescriptionCreate", "10.7", _cMMuxedFormatDescriptionCreateErr)
	}
	return _cMMuxedFormatDescriptionCreate(allocator, muxType, extensions, formatDescriptionOut), nil
}

// CMMuxedFormatDescriptionCreate creates a format description for a muxed media stream.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMuxedFormatDescriptionCreate(allocator:muxType:extensions:formatDescriptionOut:)
func CMMuxedFormatDescriptionCreate(allocator corefoundation.CFAllocatorRef, muxType CMMuxedStreamType, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *CMMuxedFormatDescriptionRef) int32 {
	result, callErr := tryCMMuxedFormatDescriptionCreate(allocator, muxType, extensions, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMPropagateAttachments func(source CMAttachmentBearerRef, destination CMAttachmentBearerRef)
var _cMPropagateAttachmentsErr error

func tryCMPropagateAttachments(source CMAttachmentBearerRef, destination CMAttachmentBearerRef) error {
	if _cMPropagateAttachments == nil {
		return symbolCallError("CMPropagateAttachments", "10.7", _cMPropagateAttachmentsErr)
	}
	_cMPropagateAttachments(source, destination)
	return nil
}

// CMPropagateAttachments copies all propagable attachments from one attachment bearer object to another.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMPropagateAttachments(_:destination:)
func CMPropagateAttachments(source CMAttachmentBearerRef, destination CMAttachmentBearerRef) {
	if callErr := tryCMPropagateAttachments(source, destination); callErr != nil {
		panic(callErr)
	}
}

var _cMRemoveAllAttachments func(target CMAttachmentBearerRef)
var _cMRemoveAllAttachmentsErr error

func tryCMRemoveAllAttachments(target CMAttachmentBearerRef) error {
	if _cMRemoveAllAttachments == nil {
		return symbolCallError("CMRemoveAllAttachments", "10.7", _cMRemoveAllAttachmentsErr)
	}
	_cMRemoveAllAttachments(target)
	return nil
}

// CMRemoveAllAttachments removes all attachments from an attachment bearer object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMRemoveAllAttachments(_:)
func CMRemoveAllAttachments(target CMAttachmentBearerRef) {
	if callErr := tryCMRemoveAllAttachments(target); callErr != nil {
		panic(callErr)
	}
}

var _cMRemoveAttachment func(target CMAttachmentBearerRef, key corefoundation.CFStringRef)
var _cMRemoveAttachmentErr error

func tryCMRemoveAttachment(target CMAttachmentBearerRef, key corefoundation.CFStringRef) error {
	if _cMRemoveAttachment == nil {
		return symbolCallError("CMRemoveAttachment", "10.7", _cMRemoveAttachmentErr)
	}
	_cMRemoveAttachment(target, key)
	return nil
}

// CMRemoveAttachment removes a specific attachment from an attachment bearer object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMRemoveAttachment(_:key:)
func CMRemoveAttachment(target CMAttachmentBearerRef, key corefoundation.CFStringRef) {
	if callErr := tryCMRemoveAttachment(target, key); callErr != nil {
		panic(callErr)
	}
}

var _cMSampleBufferCallBlockForEachSample func(sbuf uintptr) int32
var _cMSampleBufferCallBlockForEachSampleErr error

func tryCMSampleBufferCallBlockForEachSample(sbuf uintptr) (int32, error) {
	if _cMSampleBufferCallBlockForEachSample == nil {
		return 0, symbolCallError("CMSampleBufferCallBlockForEachSample", "10.10", _cMSampleBufferCallBlockForEachSampleErr)
	}
	return _cMSampleBufferCallBlockForEachSample(sbuf), nil
}

// CMSampleBufferCallBlockForEachSample calls a block for every individual sample in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCallBlockForEachSample(_:_:)
func CMSampleBufferCallBlockForEachSample(sbuf uintptr) int32 {
	result, callErr := tryCMSampleBufferCallBlockForEachSample(sbuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferCallForEachSample func(sbuf uintptr, refcon uintptr) int32
var _cMSampleBufferCallForEachSampleErr error

func tryCMSampleBufferCallForEachSample(sbuf uintptr, refcon uintptr) (int32, error) {
	if _cMSampleBufferCallForEachSample == nil {
		return 0, symbolCallError("CMSampleBufferCallForEachSample", "10.7", _cMSampleBufferCallForEachSampleErr)
	}
	return _cMSampleBufferCallForEachSample(sbuf, refcon), nil
}

// CMSampleBufferCallForEachSample calls a function for every individual sample in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCallForEachSample(_:callback:refcon:)
func CMSampleBufferCallForEachSample(sbuf uintptr, refcon uintptr) int32 {
	result, callErr := tryCMSampleBufferCallForEachSample(sbuf, refcon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferCopyPCMDataIntoAudioBufferList func(sbuf uintptr, frameOffset int32, numFrames int32, bufferList unsafe.Pointer) int32
var _cMSampleBufferCopyPCMDataIntoAudioBufferListErr error

func tryCMSampleBufferCopyPCMDataIntoAudioBufferList(sbuf uintptr, frameOffset int32, numFrames int32, bufferList unsafe.Pointer) (int32, error) {
	if _cMSampleBufferCopyPCMDataIntoAudioBufferList == nil {
		return 0, symbolCallError("CMSampleBufferCopyPCMDataIntoAudioBufferList", "10.9", _cMSampleBufferCopyPCMDataIntoAudioBufferListErr)
	}
	return _cMSampleBufferCopyPCMDataIntoAudioBufferList(sbuf, frameOffset, numFrames, bufferList), nil
}

// CMSampleBufferCopyPCMDataIntoAudioBufferList copies PCM audio data from a sample buffer into an audio buffer list.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCopyPCMDataIntoAudioBufferList(_:at:frameCount:into:)
func CMSampleBufferCopyPCMDataIntoAudioBufferList(sbuf uintptr, frameOffset int32, numFrames int32, bufferList unsafe.Pointer) int32 {
	result, callErr := tryCMSampleBufferCopyPCMDataIntoAudioBufferList(sbuf, frameOffset, numFrames, bufferList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferCopySampleBufferForRange func(allocator corefoundation.CFAllocatorRef, sbuf uintptr, sampleRange corefoundation.CFRange, sampleBufferOut *uintptr) int32
var _cMSampleBufferCopySampleBufferForRangeErr error

func tryCMSampleBufferCopySampleBufferForRange(allocator corefoundation.CFAllocatorRef, sbuf uintptr, sampleRange corefoundation.CFRange, sampleBufferOut *uintptr) (int32, error) {
	if _cMSampleBufferCopySampleBufferForRange == nil {
		return 0, symbolCallError("CMSampleBufferCopySampleBufferForRange", "10.7", _cMSampleBufferCopySampleBufferForRangeErr)
	}
	return _cMSampleBufferCopySampleBufferForRange(allocator, sbuf, sampleRange, sampleBufferOut), nil
}

// CMSampleBufferCopySampleBufferForRange creates a sample buffer that contains a range of samples from an existing sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCopySampleBufferForRange(allocator:sampleBuffer:sampleRange:sampleBufferOut:)
func CMSampleBufferCopySampleBufferForRange(allocator corefoundation.CFAllocatorRef, sbuf uintptr, sampleRange corefoundation.CFRange, sampleBufferOut *uintptr) int32 {
	result, callErr := tryCMSampleBufferCopySampleBufferForRange(allocator, sbuf, sampleRange, sampleBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferCreate func(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, dataReady bool, makeDataReadyCallback CMSampleBufferMakeDataReadyCallback, makeDataReadyRefcon unsafe.Pointer, formatDescription uintptr, numSamples int, numSampleTimingEntries int, sampleTimingArray *CMSampleTimingInfo, numSampleSizeEntries int, sampleSizeArray *uintptr, sampleBufferOut *uintptr) int32
var _cMSampleBufferCreateErr error

func tryCMSampleBufferCreate(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, dataReady bool, makeDataReadyCallback CMSampleBufferMakeDataReadyCallback, makeDataReadyRefcon unsafe.Pointer, formatDescription uintptr, numSamples int, numSampleTimingEntries int, sampleTimingArray *CMSampleTimingInfo, numSampleSizeEntries int, sampleSizeArray *uintptr, sampleBufferOut *uintptr) (int32, error) {
	if _cMSampleBufferCreate == nil {
		return 0, symbolCallError("CMSampleBufferCreate", "10.7", _cMSampleBufferCreateErr)
	}
	return _cMSampleBufferCreate(allocator, dataBuffer, dataReady, makeDataReadyCallback, makeDataReadyRefcon, formatDescription, numSamples, numSampleTimingEntries, sampleTimingArray, numSampleSizeEntries, sampleSizeArray, sampleBufferOut), nil
}

// CMSampleBufferCreate creates a sample buffer with a callback to make the data ready for use.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreate(allocator:dataBuffer:dataReady:makeDataReadyCallback:refcon:formatDescription:sampleCount:sampleTimingEntryCount:sampleTimingArray:sampleSizeEntryCount:sampleSizeArray:sampleBufferOut:)
func CMSampleBufferCreate(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, dataReady bool, makeDataReadyCallback CMSampleBufferMakeDataReadyCallback, makeDataReadyRefcon unsafe.Pointer, formatDescription uintptr, numSamples int, numSampleTimingEntries int, sampleTimingArray *CMSampleTimingInfo, numSampleSizeEntries int, sampleSizeArray *uintptr, sampleBufferOut *uintptr) int32 {
	result, callErr := tryCMSampleBufferCreate(allocator, dataBuffer, dataReady, makeDataReadyCallback, makeDataReadyRefcon, formatDescription, numSamples, numSampleTimingEntries, sampleTimingArray, numSampleSizeEntries, sampleSizeArray, sampleBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferCreateCopy func(allocator corefoundation.CFAllocatorRef, sbuf uintptr, sampleBufferOut *uintptr) int32
var _cMSampleBufferCreateCopyErr error

func tryCMSampleBufferCreateCopy(allocator corefoundation.CFAllocatorRef, sbuf uintptr, sampleBufferOut *uintptr) (int32, error) {
	if _cMSampleBufferCreateCopy == nil {
		return 0, symbolCallError("CMSampleBufferCreateCopy", "10.7", _cMSampleBufferCreateCopyErr)
	}
	return _cMSampleBufferCreateCopy(allocator, sbuf, sampleBufferOut), nil
}

// CMSampleBufferCreateCopy creates a copy of a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateCopy(allocator:sampleBuffer:sampleBufferOut:)
func CMSampleBufferCreateCopy(allocator corefoundation.CFAllocatorRef, sbuf uintptr, sampleBufferOut *uintptr) int32 {
	result, callErr := tryCMSampleBufferCreateCopy(allocator, sbuf, sampleBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferCreateCopyWithNewTiming func(allocator corefoundation.CFAllocatorRef, originalSBuf uintptr, numSampleTimingEntries int, sampleTimingArray *CMSampleTimingInfo, sampleBufferOut *uintptr) int32
var _cMSampleBufferCreateCopyWithNewTimingErr error

func tryCMSampleBufferCreateCopyWithNewTiming(allocator corefoundation.CFAllocatorRef, originalSBuf uintptr, numSampleTimingEntries int, sampleTimingArray *CMSampleTimingInfo, sampleBufferOut *uintptr) (int32, error) {
	if _cMSampleBufferCreateCopyWithNewTiming == nil {
		return 0, symbolCallError("CMSampleBufferCreateCopyWithNewTiming", "10.7", _cMSampleBufferCreateCopyWithNewTimingErr)
	}
	return _cMSampleBufferCreateCopyWithNewTiming(allocator, originalSBuf, numSampleTimingEntries, sampleTimingArray, sampleBufferOut), nil
}

// CMSampleBufferCreateCopyWithNewTiming creates a copy of a sample buffer with new timing information.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateCopyWithNewTiming(allocator:sampleBuffer:sampleTimingEntryCount:sampleTimingArray:sampleBufferOut:)
func CMSampleBufferCreateCopyWithNewTiming(allocator corefoundation.CFAllocatorRef, originalSBuf uintptr, numSampleTimingEntries int, sampleTimingArray *CMSampleTimingInfo, sampleBufferOut *uintptr) int32 {
	result, callErr := tryCMSampleBufferCreateCopyWithNewTiming(allocator, originalSBuf, numSampleTimingEntries, sampleTimingArray, sampleBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferCreateForImageBuffer func(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, dataReady bool, makeDataReadyCallback CMSampleBufferMakeDataReadyCallback, makeDataReadyRefcon unsafe.Pointer, formatDescription uintptr, sampleTiming *CMSampleTimingInfo, sampleBufferOut *uintptr) int32
var _cMSampleBufferCreateForImageBufferErr error

func tryCMSampleBufferCreateForImageBuffer(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, dataReady bool, makeDataReadyCallback CMSampleBufferMakeDataReadyCallback, makeDataReadyRefcon unsafe.Pointer, formatDescription uintptr, sampleTiming *CMSampleTimingInfo, sampleBufferOut *uintptr) (int32, error) {
	if _cMSampleBufferCreateForImageBuffer == nil {
		return 0, symbolCallError("CMSampleBufferCreateForImageBuffer", "10.7", _cMSampleBufferCreateForImageBufferErr)
	}
	return _cMSampleBufferCreateForImageBuffer(allocator, imageBuffer, dataReady, makeDataReadyCallback, makeDataReadyRefcon, formatDescription, sampleTiming, sampleBufferOut), nil
}

// CMSampleBufferCreateForImageBuffer creates a sample buffer with an image buffer and a callback to make the data ready for use.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateForImageBuffer(allocator:imageBuffer:dataReady:makeDataReadyCallback:refcon:formatDescription:sampleTiming:sampleBufferOut:)
func CMSampleBufferCreateForImageBuffer(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, dataReady bool, makeDataReadyCallback CMSampleBufferMakeDataReadyCallback, makeDataReadyRefcon unsafe.Pointer, formatDescription uintptr, sampleTiming *CMSampleTimingInfo, sampleBufferOut *uintptr) int32 {
	result, callErr := tryCMSampleBufferCreateForImageBuffer(allocator, imageBuffer, dataReady, makeDataReadyCallback, makeDataReadyRefcon, formatDescription, sampleTiming, sampleBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferCreateForImageBufferWithMakeDataReadyHandler func(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, dataReady bool, formatDescription uintptr, sampleTiming *CMSampleTimingInfo, sampleBufferOut *uintptr, makeDataReadyHandler unsafe.Pointer) int32
var _cMSampleBufferCreateForImageBufferWithMakeDataReadyHandlerErr error

func tryCMSampleBufferCreateForImageBufferWithMakeDataReadyHandler(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, dataReady bool, formatDescription uintptr, sampleTiming *CMSampleTimingInfo, sampleBufferOut *uintptr, makeDataReadyHandler CMSampleBufferMakeDataReadyHandler) (int32, error) {
	if _cMSampleBufferCreateForImageBufferWithMakeDataReadyHandler == nil {
		return 0, symbolCallError("CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler", "10.14.4", _cMSampleBufferCreateForImageBufferWithMakeDataReadyHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.IObject) int { return makeDataReadyHandler(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _cMSampleBufferCreateForImageBufferWithMakeDataReadyHandler(allocator, imageBuffer, dataReady, formatDescription, sampleTiming, sampleBufferOut, _block0), nil
}

// CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler creates a sample buffer with an image buffer and a handler to make the data ready for use.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler(_:_:_:_:_:_:_:)
func CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, dataReady bool, formatDescription uintptr, sampleTiming *CMSampleTimingInfo, sampleBufferOut *uintptr, makeDataReadyHandler CMSampleBufferMakeDataReadyHandler) int32 {
	result, callErr := tryCMSampleBufferCreateForImageBufferWithMakeDataReadyHandler(allocator, imageBuffer, dataReady, formatDescription, sampleTiming, sampleBufferOut, makeDataReadyHandler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferCreateForTaggedBufferGroup func(allocator corefoundation.CFAllocatorRef, taggedBufferGroup CMTaggedBufferGroupRef, sbufPTS CMTime, sbufDuration CMTime, formatDescription CMTaggedBufferGroupFormatDescriptionRef, sBufOut *uintptr) int32
var _cMSampleBufferCreateForTaggedBufferGroupErr error

func tryCMSampleBufferCreateForTaggedBufferGroup(allocator corefoundation.CFAllocatorRef, taggedBufferGroup CMTaggedBufferGroupRef, sbufPTS CMTime, sbufDuration CMTime, formatDescription CMTaggedBufferGroupFormatDescriptionRef, sBufOut *uintptr) (int32, error) {
	if _cMSampleBufferCreateForTaggedBufferGroup == nil {
		return 0, symbolCallError("CMSampleBufferCreateForTaggedBufferGroup", "14.0", _cMSampleBufferCreateForTaggedBufferGroupErr)
	}
	return _cMSampleBufferCreateForTaggedBufferGroup(allocator, taggedBufferGroup, sbufPTS, sbufDuration, formatDescription, sBufOut), nil
}

// CMSampleBufferCreateForTaggedBufferGroup creates a new sample buffer from a tagged buffer group.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateForTaggedBufferGroup
func CMSampleBufferCreateForTaggedBufferGroup(allocator corefoundation.CFAllocatorRef, taggedBufferGroup CMTaggedBufferGroupRef, sbufPTS CMTime, sbufDuration CMTime, formatDescription CMTaggedBufferGroupFormatDescriptionRef, sBufOut *uintptr) int32 {
	result, callErr := tryCMSampleBufferCreateForTaggedBufferGroup(allocator, taggedBufferGroup, sbufPTS, sbufDuration, formatDescription, sBufOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferCreateReady func(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, formatDescription uintptr, numSamples int, numSampleTimingEntries int, sampleTimingArray *CMSampleTimingInfo, numSampleSizeEntries int, sampleSizeArray *uintptr, sampleBufferOut *uintptr) int32
var _cMSampleBufferCreateReadyErr error

func tryCMSampleBufferCreateReady(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, formatDescription uintptr, numSamples int, numSampleTimingEntries int, sampleTimingArray *CMSampleTimingInfo, numSampleSizeEntries int, sampleSizeArray *uintptr, sampleBufferOut *uintptr) (int32, error) {
	if _cMSampleBufferCreateReady == nil {
		return 0, symbolCallError("CMSampleBufferCreateReady", "10.10", _cMSampleBufferCreateReadyErr)
	}
	return _cMSampleBufferCreateReady(allocator, dataBuffer, formatDescription, numSamples, numSampleTimingEntries, sampleTimingArray, numSampleSizeEntries, sampleSizeArray, sampleBufferOut), nil
}

// CMSampleBufferCreateReady creates a sample buffer with media data.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateReady(allocator:dataBuffer:formatDescription:sampleCount:sampleTimingEntryCount:sampleTimingArray:sampleSizeEntryCount:sampleSizeArray:sampleBufferOut:)
func CMSampleBufferCreateReady(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, formatDescription uintptr, numSamples int, numSampleTimingEntries int, sampleTimingArray *CMSampleTimingInfo, numSampleSizeEntries int, sampleSizeArray *uintptr, sampleBufferOut *uintptr) int32 {
	result, callErr := tryCMSampleBufferCreateReady(allocator, dataBuffer, formatDescription, numSamples, numSampleTimingEntries, sampleTimingArray, numSampleSizeEntries, sampleSizeArray, sampleBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferCreateReadyWithImageBuffer func(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, formatDescription uintptr, sampleTiming *CMSampleTimingInfo, sampleBufferOut *uintptr) int32
var _cMSampleBufferCreateReadyWithImageBufferErr error

func tryCMSampleBufferCreateReadyWithImageBuffer(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, formatDescription uintptr, sampleTiming *CMSampleTimingInfo, sampleBufferOut *uintptr) (int32, error) {
	if _cMSampleBufferCreateReadyWithImageBuffer == nil {
		return 0, symbolCallError("CMSampleBufferCreateReadyWithImageBuffer", "10.10", _cMSampleBufferCreateReadyWithImageBufferErr)
	}
	return _cMSampleBufferCreateReadyWithImageBuffer(allocator, imageBuffer, formatDescription, sampleTiming, sampleBufferOut), nil
}

// CMSampleBufferCreateReadyWithImageBuffer creates a sample buffer with image data.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateReadyWithImageBuffer(allocator:imageBuffer:formatDescription:sampleTiming:sampleBufferOut:)
func CMSampleBufferCreateReadyWithImageBuffer(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, formatDescription uintptr, sampleTiming *CMSampleTimingInfo, sampleBufferOut *uintptr) int32 {
	result, callErr := tryCMSampleBufferCreateReadyWithImageBuffer(allocator, imageBuffer, formatDescription, sampleTiming, sampleBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferCreateWithMakeDataReadyHandler func(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, dataReady bool, formatDescription uintptr, numSamples int, numSampleTimingEntries int, sampleTimingArray *CMSampleTimingInfo, numSampleSizeEntries int, sampleSizeArray *uintptr, sampleBufferOut *uintptr, makeDataReadyHandler unsafe.Pointer) int32
var _cMSampleBufferCreateWithMakeDataReadyHandlerErr error

func tryCMSampleBufferCreateWithMakeDataReadyHandler(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, dataReady bool, formatDescription uintptr, numSamples int, numSampleTimingEntries int, sampleTimingArray *CMSampleTimingInfo, numSampleSizeEntries int, sampleSizeArray *uintptr, sampleBufferOut *uintptr, makeDataReadyHandler CMSampleBufferMakeDataReadyHandler) (int32, error) {
	if _cMSampleBufferCreateWithMakeDataReadyHandler == nil {
		return 0, symbolCallError("CMSampleBufferCreateWithMakeDataReadyHandler", "10.14.4", _cMSampleBufferCreateWithMakeDataReadyHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.IObject) int { return makeDataReadyHandler(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _cMSampleBufferCreateWithMakeDataReadyHandler(allocator, dataBuffer, dataReady, formatDescription, numSamples, numSampleTimingEntries, sampleTimingArray, numSampleSizeEntries, sampleSizeArray, sampleBufferOut, _block0), nil
}

// CMSampleBufferCreateWithMakeDataReadyHandler creates a sample buffer with a handler to make the data ready for use.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateWithMakeDataReadyHandler(_:_:_:_:_:_:_:_:_:_:_:)
func CMSampleBufferCreateWithMakeDataReadyHandler(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, dataReady bool, formatDescription uintptr, numSamples int, numSampleTimingEntries int, sampleTimingArray *CMSampleTimingInfo, numSampleSizeEntries int, sampleSizeArray *uintptr, sampleBufferOut *uintptr, makeDataReadyHandler CMSampleBufferMakeDataReadyHandler) int32 {
	result, callErr := tryCMSampleBufferCreateWithMakeDataReadyHandler(allocator, dataBuffer, dataReady, formatDescription, numSamples, numSampleTimingEntries, sampleTimingArray, numSampleSizeEntries, sampleSizeArray, sampleBufferOut, makeDataReadyHandler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferDataIsReady func(sbuf uintptr) bool
var _cMSampleBufferDataIsReadyErr error

func tryCMSampleBufferDataIsReady(sbuf uintptr) (bool, error) {
	if _cMSampleBufferDataIsReady == nil {
		return false, symbolCallError("CMSampleBufferDataIsReady", "10.7", _cMSampleBufferDataIsReadyErr)
	}
	return _cMSampleBufferDataIsReady(sbuf), nil
}

// CMSampleBufferDataIsReady returns a Boolean value that indicates whether the sample buffer’s data is ready for use.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferDataIsReady(_:)
func CMSampleBufferDataIsReady(sbuf uintptr) bool {
	result, callErr := tryCMSampleBufferDataIsReady(sbuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetAudioBufferListWithRetainedBlockBuffer func(sbuf uintptr, bufferListSizeNeededOut *uintptr, bufferListOut unsafe.Pointer, bufferListSize uintptr, blockBufferStructureAllocator corefoundation.CFAllocatorRef, blockBufferBlockAllocator corefoundation.CFAllocatorRef, flags uint32, blockBufferOut *uintptr) int32
var _cMSampleBufferGetAudioBufferListWithRetainedBlockBufferErr error

func tryCMSampleBufferGetAudioBufferListWithRetainedBlockBuffer(sbuf uintptr, bufferListSizeNeededOut *uintptr, bufferListOut unsafe.Pointer, bufferListSize uintptr, blockBufferStructureAllocator corefoundation.CFAllocatorRef, blockBufferBlockAllocator corefoundation.CFAllocatorRef, flags uint32, blockBufferOut *uintptr) (int32, error) {
	if _cMSampleBufferGetAudioBufferListWithRetainedBlockBuffer == nil {
		return 0, symbolCallError("CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer", "10.7", _cMSampleBufferGetAudioBufferListWithRetainedBlockBufferErr)
	}
	return _cMSampleBufferGetAudioBufferListWithRetainedBlockBuffer(sbuf, bufferListSizeNeededOut, bufferListOut, bufferListSize, blockBufferStructureAllocator, blockBufferBlockAllocator, flags, blockBufferOut), nil
}

// CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer returns an audio buffer list that contains the media data.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer(_:bufferListSizeNeededOut:bufferListOut:bufferListSize:blockBufferAllocator:blockBufferMemoryAllocator:flags:blockBufferOut:)
func CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer(sbuf uintptr, bufferListSizeNeededOut *uintptr, bufferListOut unsafe.Pointer, bufferListSize uintptr, blockBufferStructureAllocator corefoundation.CFAllocatorRef, blockBufferBlockAllocator corefoundation.CFAllocatorRef, flags uint32, blockBufferOut *uintptr) int32 {
	result, callErr := tryCMSampleBufferGetAudioBufferListWithRetainedBlockBuffer(sbuf, bufferListSizeNeededOut, bufferListOut, bufferListSize, blockBufferStructureAllocator, blockBufferBlockAllocator, flags, blockBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetAudioStreamPacketDescriptions func(sbuf uintptr, packetDescriptionsSize uintptr, packetDescriptionsOut unsafe.Pointer, packetDescriptionsSizeNeededOut *uintptr) int32
var _cMSampleBufferGetAudioStreamPacketDescriptionsErr error

func tryCMSampleBufferGetAudioStreamPacketDescriptions(sbuf uintptr, packetDescriptionsSize uintptr, packetDescriptionsOut unsafe.Pointer, packetDescriptionsSizeNeededOut *uintptr) (int32, error) {
	if _cMSampleBufferGetAudioStreamPacketDescriptions == nil {
		return 0, symbolCallError("CMSampleBufferGetAudioStreamPacketDescriptions", "10.7", _cMSampleBufferGetAudioStreamPacketDescriptionsErr)
	}
	return _cMSampleBufferGetAudioStreamPacketDescriptions(sbuf, packetDescriptionsSize, packetDescriptionsOut, packetDescriptionsSizeNeededOut), nil
}

// CMSampleBufferGetAudioStreamPacketDescriptions creates an array of audio stream packet descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetAudioStreamPacketDescriptions(_:allocatedSize:packetDescriptionsOut:packetDescriptionsSizeNeededOut:)
func CMSampleBufferGetAudioStreamPacketDescriptions(sbuf uintptr, packetDescriptionsSize uintptr, packetDescriptionsOut unsafe.Pointer, packetDescriptionsSizeNeededOut *uintptr) int32 {
	result, callErr := tryCMSampleBufferGetAudioStreamPacketDescriptions(sbuf, packetDescriptionsSize, packetDescriptionsOut, packetDescriptionsSizeNeededOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetAudioStreamPacketDescriptionsPtr func(sbuf uintptr, packetDescriptionsPointerOut unsafe.Pointer, packetDescriptionsSizeOut *uintptr) int32
var _cMSampleBufferGetAudioStreamPacketDescriptionsPtrErr error

func tryCMSampleBufferGetAudioStreamPacketDescriptionsPtr(sbuf uintptr, packetDescriptionsPointerOut unsafe.Pointer, packetDescriptionsSizeOut *uintptr) (int32, error) {
	if _cMSampleBufferGetAudioStreamPacketDescriptionsPtr == nil {
		return 0, symbolCallError("CMSampleBufferGetAudioStreamPacketDescriptionsPtr", "10.7", _cMSampleBufferGetAudioStreamPacketDescriptionsPtrErr)
	}
	return _cMSampleBufferGetAudioStreamPacketDescriptionsPtr(sbuf, packetDescriptionsPointerOut, packetDescriptionsSizeOut), nil
}

// CMSampleBufferGetAudioStreamPacketDescriptionsPtr returns a pointer to a constant array of audio stream packet descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetAudioStreamPacketDescriptionsPtr(_:packetDescriptionsPointerOut:sizeOut:)
func CMSampleBufferGetAudioStreamPacketDescriptionsPtr(sbuf uintptr, packetDescriptionsPointerOut unsafe.Pointer, packetDescriptionsSizeOut *uintptr) int32 {
	result, callErr := tryCMSampleBufferGetAudioStreamPacketDescriptionsPtr(sbuf, packetDescriptionsPointerOut, packetDescriptionsSizeOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetDataBuffer func(sbuf uintptr) uintptr
var _cMSampleBufferGetDataBufferErr error

func tryCMSampleBufferGetDataBuffer(sbuf uintptr) (uintptr, error) {
	if _cMSampleBufferGetDataBuffer == nil {
		return 0, symbolCallError("CMSampleBufferGetDataBuffer", "10.7", _cMSampleBufferGetDataBufferErr)
	}
	return _cMSampleBufferGetDataBuffer(sbuf), nil
}

// CMSampleBufferGetDataBuffer returns a block buffer that contains the media data.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetDataBuffer(_:)
func CMSampleBufferGetDataBuffer(sbuf uintptr) uintptr {
	result, callErr := tryCMSampleBufferGetDataBuffer(sbuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetDecodeTimeStamp func(sbuf uintptr) CMTime
var _cMSampleBufferGetDecodeTimeStampErr error

func tryCMSampleBufferGetDecodeTimeStamp(sbuf uintptr) (CMTime, error) {
	if _cMSampleBufferGetDecodeTimeStamp == nil {
		return CMTime{}, symbolCallError("CMSampleBufferGetDecodeTimeStamp", "10.7", _cMSampleBufferGetDecodeTimeStampErr)
	}
	return _cMSampleBufferGetDecodeTimeStamp(sbuf), nil
}

// CMSampleBufferGetDecodeTimeStamp returns the decode timestamp that’s the earliest numerically of all the samples in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetDecodeTimeStamp(_:)
func CMSampleBufferGetDecodeTimeStamp(sbuf uintptr) CMTime {
	result, callErr := tryCMSampleBufferGetDecodeTimeStamp(sbuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetDuration func(sbuf uintptr) CMTime
var _cMSampleBufferGetDurationErr error

func tryCMSampleBufferGetDuration(sbuf uintptr) (CMTime, error) {
	if _cMSampleBufferGetDuration == nil {
		return CMTime{}, symbolCallError("CMSampleBufferGetDuration", "10.7", _cMSampleBufferGetDurationErr)
	}
	return _cMSampleBufferGetDuration(sbuf), nil
}

// CMSampleBufferGetDuration returns the total duration of a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetDuration(_:)
func CMSampleBufferGetDuration(sbuf uintptr) CMTime {
	result, callErr := tryCMSampleBufferGetDuration(sbuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetFormatDescription func(sbuf uintptr) uintptr
var _cMSampleBufferGetFormatDescriptionErr error

func tryCMSampleBufferGetFormatDescription(sbuf uintptr) (uintptr, error) {
	if _cMSampleBufferGetFormatDescription == nil {
		return 0, symbolCallError("CMSampleBufferGetFormatDescription", "10.7", _cMSampleBufferGetFormatDescriptionErr)
	}
	return _cMSampleBufferGetFormatDescription(sbuf), nil
}

// CMSampleBufferGetFormatDescription returns the format description of the samples in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetFormatDescription(_:)
func CMSampleBufferGetFormatDescription(sbuf uintptr) uintptr {
	result, callErr := tryCMSampleBufferGetFormatDescription(sbuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetImageBuffer func(sbuf uintptr) corevideo.CVImageBufferRef
var _cMSampleBufferGetImageBufferErr error

func tryCMSampleBufferGetImageBuffer(sbuf uintptr) (corevideo.CVImageBufferRef, error) {
	if _cMSampleBufferGetImageBuffer == nil {
		return 0, symbolCallError("CMSampleBufferGetImageBuffer", "10.7", _cMSampleBufferGetImageBufferErr)
	}
	return _cMSampleBufferGetImageBuffer(sbuf), nil
}

// CMSampleBufferGetImageBuffer returns an image buffer that contains the media data.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetImageBuffer(_:)
func CMSampleBufferGetImageBuffer(sbuf uintptr) corevideo.CVImageBufferRef {
	result, callErr := tryCMSampleBufferGetImageBuffer(sbuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetNumSamples func(sbuf uintptr) int
var _cMSampleBufferGetNumSamplesErr error

func tryCMSampleBufferGetNumSamples(sbuf uintptr) (int, error) {
	if _cMSampleBufferGetNumSamples == nil {
		return 0, symbolCallError("CMSampleBufferGetNumSamples", "10.7", _cMSampleBufferGetNumSamplesErr)
	}
	return _cMSampleBufferGetNumSamples(sbuf), nil
}

// CMSampleBufferGetNumSamples returns the number of media samples in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetNumSamples(_:)
func CMSampleBufferGetNumSamples(sbuf uintptr) int {
	result, callErr := tryCMSampleBufferGetNumSamples(sbuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetOutputDecodeTimeStamp func(sbuf uintptr) CMTime
var _cMSampleBufferGetOutputDecodeTimeStampErr error

func tryCMSampleBufferGetOutputDecodeTimeStamp(sbuf uintptr) (CMTime, error) {
	if _cMSampleBufferGetOutputDecodeTimeStamp == nil {
		return CMTime{}, symbolCallError("CMSampleBufferGetOutputDecodeTimeStamp", "10.7", _cMSampleBufferGetOutputDecodeTimeStampErr)
	}
	return _cMSampleBufferGetOutputDecodeTimeStamp(sbuf), nil
}

// CMSampleBufferGetOutputDecodeTimeStamp returns the output decode timestamp of a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetOutputDecodeTimeStamp(_:)
func CMSampleBufferGetOutputDecodeTimeStamp(sbuf uintptr) CMTime {
	result, callErr := tryCMSampleBufferGetOutputDecodeTimeStamp(sbuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetOutputDuration func(sbuf uintptr) CMTime
var _cMSampleBufferGetOutputDurationErr error

func tryCMSampleBufferGetOutputDuration(sbuf uintptr) (CMTime, error) {
	if _cMSampleBufferGetOutputDuration == nil {
		return CMTime{}, symbolCallError("CMSampleBufferGetOutputDuration", "10.7", _cMSampleBufferGetOutputDurationErr)
	}
	return _cMSampleBufferGetOutputDuration(sbuf), nil
}

// CMSampleBufferGetOutputDuration returns the output duration of a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetOutputDuration(_:)
func CMSampleBufferGetOutputDuration(sbuf uintptr) CMTime {
	result, callErr := tryCMSampleBufferGetOutputDuration(sbuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetOutputPresentationTimeStamp func(sbuf uintptr) CMTime
var _cMSampleBufferGetOutputPresentationTimeStampErr error

func tryCMSampleBufferGetOutputPresentationTimeStamp(sbuf uintptr) (CMTime, error) {
	if _cMSampleBufferGetOutputPresentationTimeStamp == nil {
		return CMTime{}, symbolCallError("CMSampleBufferGetOutputPresentationTimeStamp", "10.7", _cMSampleBufferGetOutputPresentationTimeStampErr)
	}
	return _cMSampleBufferGetOutputPresentationTimeStamp(sbuf), nil
}

// CMSampleBufferGetOutputPresentationTimeStamp returns the output presentation timestamp of a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetOutputPresentationTimeStamp(_:)
func CMSampleBufferGetOutputPresentationTimeStamp(sbuf uintptr) CMTime {
	result, callErr := tryCMSampleBufferGetOutputPresentationTimeStamp(sbuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetOutputSampleTimingInfoArray func(sbuf uintptr, timingArrayEntries int, timingArrayOut *CMSampleTimingInfo, timingArrayEntriesNeededOut *int) int32
var _cMSampleBufferGetOutputSampleTimingInfoArrayErr error

func tryCMSampleBufferGetOutputSampleTimingInfoArray(sbuf uintptr, timingArrayEntries int, timingArrayOut *CMSampleTimingInfo, timingArrayEntriesNeededOut *int) (int32, error) {
	if _cMSampleBufferGetOutputSampleTimingInfoArray == nil {
		return 0, symbolCallError("CMSampleBufferGetOutputSampleTimingInfoArray", "10.7", _cMSampleBufferGetOutputSampleTimingInfoArrayErr)
	}
	return _cMSampleBufferGetOutputSampleTimingInfoArray(sbuf, timingArrayEntries, timingArrayOut, timingArrayEntriesNeededOut), nil
}

// CMSampleBufferGetOutputSampleTimingInfoArray retrieves an array of output timing information structures that represents each sample in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetOutputSampleTimingInfoArray(_:entryCount:arrayToFill:entriesNeededOut:)
func CMSampleBufferGetOutputSampleTimingInfoArray(sbuf uintptr, timingArrayEntries int, timingArrayOut *CMSampleTimingInfo, timingArrayEntriesNeededOut *int) int32 {
	result, callErr := tryCMSampleBufferGetOutputSampleTimingInfoArray(sbuf, timingArrayEntries, timingArrayOut, timingArrayEntriesNeededOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetPresentationTimeStamp func(sbuf uintptr) CMTime
var _cMSampleBufferGetPresentationTimeStampErr error

func tryCMSampleBufferGetPresentationTimeStamp(sbuf uintptr) (CMTime, error) {
	if _cMSampleBufferGetPresentationTimeStamp == nil {
		return CMTime{}, symbolCallError("CMSampleBufferGetPresentationTimeStamp", "10.7", _cMSampleBufferGetPresentationTimeStampErr)
	}
	return _cMSampleBufferGetPresentationTimeStamp(sbuf), nil
}

// CMSampleBufferGetPresentationTimeStamp returns the presentation timestamp that’s the earliest numerically of all the samples in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetPresentationTimeStamp(_:)
func CMSampleBufferGetPresentationTimeStamp(sbuf uintptr) CMTime {
	result, callErr := tryCMSampleBufferGetPresentationTimeStamp(sbuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetSampleAttachmentsArray func(sbuf uintptr, createIfNecessary bool) corefoundation.CFArrayRef
var _cMSampleBufferGetSampleAttachmentsArrayErr error

func tryCMSampleBufferGetSampleAttachmentsArray(sbuf uintptr, createIfNecessary bool) (corefoundation.CFArrayRef, error) {
	if _cMSampleBufferGetSampleAttachmentsArray == nil {
		return 0, symbolCallError("CMSampleBufferGetSampleAttachmentsArray", "10.7", _cMSampleBufferGetSampleAttachmentsArrayErr)
	}
	return _cMSampleBufferGetSampleAttachmentsArray(sbuf, createIfNecessary), nil
}

// CMSampleBufferGetSampleAttachmentsArray retrieves an array of sample attachment dictionaries that represents each sample in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetSampleAttachmentsArray(_:createIfNecessary:)
func CMSampleBufferGetSampleAttachmentsArray(sbuf uintptr, createIfNecessary bool) corefoundation.CFArrayRef {
	result, callErr := tryCMSampleBufferGetSampleAttachmentsArray(sbuf, createIfNecessary)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetSampleSize func(sbuf uintptr, sampleIndex int) uintptr
var _cMSampleBufferGetSampleSizeErr error

func tryCMSampleBufferGetSampleSize(sbuf uintptr, sampleIndex int) (uintptr, error) {
	if _cMSampleBufferGetSampleSize == nil {
		return 0, symbolCallError("CMSampleBufferGetSampleSize", "10.7", _cMSampleBufferGetSampleSizeErr)
	}
	return _cMSampleBufferGetSampleSize(sbuf, sampleIndex), nil
}

// CMSampleBufferGetSampleSize returns the size in bytes of a specified sample in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetSampleSize(_:at:)
func CMSampleBufferGetSampleSize(sbuf uintptr, sampleIndex int) uintptr {
	result, callErr := tryCMSampleBufferGetSampleSize(sbuf, sampleIndex)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetSampleSizeArray func(sbuf uintptr, sizeArrayEntries int, sizeArrayOut *uintptr, sizeArrayEntriesNeededOut *int) int32
var _cMSampleBufferGetSampleSizeArrayErr error

func tryCMSampleBufferGetSampleSizeArray(sbuf uintptr, sizeArrayEntries int, sizeArrayOut *uintptr, sizeArrayEntriesNeededOut *int) (int32, error) {
	if _cMSampleBufferGetSampleSizeArray == nil {
		return 0, symbolCallError("CMSampleBufferGetSampleSizeArray", "10.7", _cMSampleBufferGetSampleSizeArrayErr)
	}
	return _cMSampleBufferGetSampleSizeArray(sbuf, sizeArrayEntries, sizeArrayOut, sizeArrayEntriesNeededOut), nil
}

// CMSampleBufferGetSampleSizeArray retrieves an array of sample sizes that represents each sample in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetSampleSizeArray(_:entryCount:arrayToFill:entriesNeededOut:)
func CMSampleBufferGetSampleSizeArray(sbuf uintptr, sizeArrayEntries int, sizeArrayOut *uintptr, sizeArrayEntriesNeededOut *int) int32 {
	result, callErr := tryCMSampleBufferGetSampleSizeArray(sbuf, sizeArrayEntries, sizeArrayOut, sizeArrayEntriesNeededOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetSampleTimingInfo func(sbuf uintptr, sampleIndex int, timingInfoOut *CMSampleTimingInfo) int32
var _cMSampleBufferGetSampleTimingInfoErr error

func tryCMSampleBufferGetSampleTimingInfo(sbuf uintptr, sampleIndex int, timingInfoOut *CMSampleTimingInfo) (int32, error) {
	if _cMSampleBufferGetSampleTimingInfo == nil {
		return 0, symbolCallError("CMSampleBufferGetSampleTimingInfo", "10.7", _cMSampleBufferGetSampleTimingInfoErr)
	}
	return _cMSampleBufferGetSampleTimingInfo(sbuf, sampleIndex, timingInfoOut), nil
}

// CMSampleBufferGetSampleTimingInfo retrieves a timing information structure that describes a specified sample in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetSampleTimingInfo(_:at:timingInfoOut:)
func CMSampleBufferGetSampleTimingInfo(sbuf uintptr, sampleIndex int, timingInfoOut *CMSampleTimingInfo) int32 {
	result, callErr := tryCMSampleBufferGetSampleTimingInfo(sbuf, sampleIndex, timingInfoOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetSampleTimingInfoArray func(sbuf uintptr, numSampleTimingEntries int, timingArrayOut *CMSampleTimingInfo, timingArrayEntriesNeededOut *int) int32
var _cMSampleBufferGetSampleTimingInfoArrayErr error

func tryCMSampleBufferGetSampleTimingInfoArray(sbuf uintptr, numSampleTimingEntries int, timingArrayOut *CMSampleTimingInfo, timingArrayEntriesNeededOut *int) (int32, error) {
	if _cMSampleBufferGetSampleTimingInfoArray == nil {
		return 0, symbolCallError("CMSampleBufferGetSampleTimingInfoArray", "10.7", _cMSampleBufferGetSampleTimingInfoArrayErr)
	}
	return _cMSampleBufferGetSampleTimingInfoArray(sbuf, numSampleTimingEntries, timingArrayOut, timingArrayEntriesNeededOut), nil
}

// CMSampleBufferGetSampleTimingInfoArray retrieves an array of sample timing information structures that represents each sample in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetSampleTimingInfoArray(_:entryCount:arrayToFill:entriesNeededOut:)
func CMSampleBufferGetSampleTimingInfoArray(sbuf uintptr, numSampleTimingEntries int, timingArrayOut *CMSampleTimingInfo, timingArrayEntriesNeededOut *int) int32 {
	result, callErr := tryCMSampleBufferGetSampleTimingInfoArray(sbuf, numSampleTimingEntries, timingArrayOut, timingArrayEntriesNeededOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetTaggedBufferGroup func(sbuf uintptr) CMTaggedBufferGroupRef
var _cMSampleBufferGetTaggedBufferGroupErr error

func tryCMSampleBufferGetTaggedBufferGroup(sbuf uintptr) (CMTaggedBufferGroupRef, error) {
	if _cMSampleBufferGetTaggedBufferGroup == nil {
		return 0, symbolCallError("CMSampleBufferGetTaggedBufferGroup", "14.0", _cMSampleBufferGetTaggedBufferGroupErr)
	}
	return _cMSampleBufferGetTaggedBufferGroup(sbuf), nil
}

// CMSampleBufferGetTaggedBufferGroup gets the tagged buffer group of a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetTaggedBufferGroup
func CMSampleBufferGetTaggedBufferGroup(sbuf uintptr) CMTaggedBufferGroupRef {
	result, callErr := tryCMSampleBufferGetTaggedBufferGroup(sbuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetTotalSampleSize func(sbuf uintptr) uintptr
var _cMSampleBufferGetTotalSampleSizeErr error

func tryCMSampleBufferGetTotalSampleSize(sbuf uintptr) (uintptr, error) {
	if _cMSampleBufferGetTotalSampleSize == nil {
		return 0, symbolCallError("CMSampleBufferGetTotalSampleSize", "10.7", _cMSampleBufferGetTotalSampleSizeErr)
	}
	return _cMSampleBufferGetTotalSampleSize(sbuf), nil
}

// CMSampleBufferGetTotalSampleSize returns the total size in bytes of sample data in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetTotalSampleSize(_:)
func CMSampleBufferGetTotalSampleSize(sbuf uintptr) uintptr {
	result, callErr := tryCMSampleBufferGetTotalSampleSize(sbuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferGetTypeID func() uint
var _cMSampleBufferGetTypeIDErr error

func tryCMSampleBufferGetTypeID() (uint, error) {
	if _cMSampleBufferGetTypeID == nil {
		return 0, symbolCallError("CMSampleBufferGetTypeID", "10.7", _cMSampleBufferGetTypeIDErr)
	}
	return _cMSampleBufferGetTypeID(), nil
}

// CMSampleBufferGetTypeID returns the type identifier of sample buffer objects.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetTypeID()
func CMSampleBufferGetTypeID() uint {
	result, callErr := tryCMSampleBufferGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferHasDataFailed func(sbuf uintptr, statusOut *int32) bool
var _cMSampleBufferHasDataFailedErr error

func tryCMSampleBufferHasDataFailed(sbuf uintptr, statusOut *int32) (bool, error) {
	if _cMSampleBufferHasDataFailed == nil {
		return false, symbolCallError("CMSampleBufferHasDataFailed", "10.10", _cMSampleBufferHasDataFailedErr)
	}
	return _cMSampleBufferHasDataFailed(sbuf, statusOut), nil
}

// CMSampleBufferHasDataFailed returns a Boolean value that indicates whether the sample buffer’s data loading request failed.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferHasDataFailed(_:statusOut:)
func CMSampleBufferHasDataFailed(sbuf uintptr, statusOut *int32) bool {
	result, callErr := tryCMSampleBufferHasDataFailed(sbuf, statusOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferInvalidate func(sbuf uintptr) int32
var _cMSampleBufferInvalidateErr error

func tryCMSampleBufferInvalidate(sbuf uintptr) (int32, error) {
	if _cMSampleBufferInvalidate == nil {
		return 0, symbolCallError("CMSampleBufferInvalidate", "10.7", _cMSampleBufferInvalidateErr)
	}
	return _cMSampleBufferInvalidate(sbuf), nil
}

// CMSampleBufferInvalidate invalidates a sample buffer by calling its invalidation callback.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferInvalidate(_:)
func CMSampleBufferInvalidate(sbuf uintptr) int32 {
	result, callErr := tryCMSampleBufferInvalidate(sbuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferIsValid func(sbuf uintptr) bool
var _cMSampleBufferIsValidErr error

func tryCMSampleBufferIsValid(sbuf uintptr) (bool, error) {
	if _cMSampleBufferIsValid == nil {
		return false, symbolCallError("CMSampleBufferIsValid", "10.7", _cMSampleBufferIsValidErr)
	}
	return _cMSampleBufferIsValid(sbuf), nil
}

// CMSampleBufferIsValid returns a Boolean value that indicates whether a sample buffer is valid.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferIsValid(_:)
func CMSampleBufferIsValid(sbuf uintptr) bool {
	result, callErr := tryCMSampleBufferIsValid(sbuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferMakeDataReady func(sbuf uintptr) int32
var _cMSampleBufferMakeDataReadyErr error

func tryCMSampleBufferMakeDataReady(sbuf uintptr) (int32, error) {
	if _cMSampleBufferMakeDataReady == nil {
		return 0, symbolCallError("CMSampleBufferMakeDataReady", "10.7", _cMSampleBufferMakeDataReadyErr)
	}
	return _cMSampleBufferMakeDataReady(sbuf), nil
}

// CMSampleBufferMakeDataReady makes the sample buffer’s data ready for use by invoking its callback to load the data.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferMakeDataReady(_:)
func CMSampleBufferMakeDataReady(sbuf uintptr) int32 {
	result, callErr := tryCMSampleBufferMakeDataReady(sbuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferSetDataBuffer func(sbuf uintptr, dataBuffer uintptr) int32
var _cMSampleBufferSetDataBufferErr error

func tryCMSampleBufferSetDataBuffer(sbuf uintptr, dataBuffer uintptr) (int32, error) {
	if _cMSampleBufferSetDataBuffer == nil {
		return 0, symbolCallError("CMSampleBufferSetDataBuffer", "10.7", _cMSampleBufferSetDataBufferErr)
	}
	return _cMSampleBufferSetDataBuffer(sbuf, dataBuffer), nil
}

// CMSampleBufferSetDataBuffer sets a block buffer of media data on a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetDataBuffer(_:newValue:)
func CMSampleBufferSetDataBuffer(sbuf uintptr, dataBuffer uintptr) int32 {
	result, callErr := tryCMSampleBufferSetDataBuffer(sbuf, dataBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferSetDataBufferFromAudioBufferList func(sbuf uintptr, blockBufferStructureAllocator corefoundation.CFAllocatorRef, blockBufferBlockAllocator corefoundation.CFAllocatorRef, flags uint32, bufferList unsafe.Pointer) int32
var _cMSampleBufferSetDataBufferFromAudioBufferListErr error

func tryCMSampleBufferSetDataBufferFromAudioBufferList(sbuf uintptr, blockBufferStructureAllocator corefoundation.CFAllocatorRef, blockBufferBlockAllocator corefoundation.CFAllocatorRef, flags uint32, bufferList unsafe.Pointer) (int32, error) {
	if _cMSampleBufferSetDataBufferFromAudioBufferList == nil {
		return 0, symbolCallError("CMSampleBufferSetDataBufferFromAudioBufferList", "10.7", _cMSampleBufferSetDataBufferFromAudioBufferListErr)
	}
	return _cMSampleBufferSetDataBufferFromAudioBufferList(sbuf, blockBufferStructureAllocator, blockBufferBlockAllocator, flags, bufferList), nil
}

// CMSampleBufferSetDataBufferFromAudioBufferList creates a block buffer that contains a copy of the data from an audio buffer list.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetDataBufferFromAudioBufferList(_:blockBufferAllocator:blockBufferMemoryAllocator:flags:bufferList:)
func CMSampleBufferSetDataBufferFromAudioBufferList(sbuf uintptr, blockBufferStructureAllocator corefoundation.CFAllocatorRef, blockBufferBlockAllocator corefoundation.CFAllocatorRef, flags uint32, bufferList unsafe.Pointer) int32 {
	result, callErr := tryCMSampleBufferSetDataBufferFromAudioBufferList(sbuf, blockBufferStructureAllocator, blockBufferBlockAllocator, flags, bufferList)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferSetDataFailed func(sbuf uintptr, status int32) int32
var _cMSampleBufferSetDataFailedErr error

func tryCMSampleBufferSetDataFailed(sbuf uintptr, status int32) (int32, error) {
	if _cMSampleBufferSetDataFailed == nil {
		return 0, symbolCallError("CMSampleBufferSetDataFailed", "10.10", _cMSampleBufferSetDataFailedErr)
	}
	return _cMSampleBufferSetDataFailed(sbuf, status), nil
}

// CMSampleBufferSetDataFailed marks the sample buffer’s data as failed to indicate that it won’t become ready.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetDataFailed(_:status:)
func CMSampleBufferSetDataFailed(sbuf uintptr, status int32) int32 {
	result, callErr := tryCMSampleBufferSetDataFailed(sbuf, status)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferSetDataReady func(sbuf uintptr) int32
var _cMSampleBufferSetDataReadyErr error

func tryCMSampleBufferSetDataReady(sbuf uintptr) (int32, error) {
	if _cMSampleBufferSetDataReady == nil {
		return 0, symbolCallError("CMSampleBufferSetDataReady", "10.7", _cMSampleBufferSetDataReadyErr)
	}
	return _cMSampleBufferSetDataReady(sbuf), nil
}

// CMSampleBufferSetDataReady marks a sample buffer’s data as ready for use.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetDataReady(_:)
func CMSampleBufferSetDataReady(sbuf uintptr) int32 {
	result, callErr := tryCMSampleBufferSetDataReady(sbuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferSetInvalidateCallback func(sbuf uintptr, invalidateCallback CMSampleBufferInvalidateCallback, invalidateRefCon uint64) int32
var _cMSampleBufferSetInvalidateCallbackErr error

func tryCMSampleBufferSetInvalidateCallback(sbuf uintptr, invalidateCallback CMSampleBufferInvalidateCallback, invalidateRefCon uint64) (int32, error) {
	if _cMSampleBufferSetInvalidateCallback == nil {
		return 0, symbolCallError("CMSampleBufferSetInvalidateCallback", "10.7", _cMSampleBufferSetInvalidateCallbackErr)
	}
	return _cMSampleBufferSetInvalidateCallback(sbuf, invalidateCallback, invalidateRefCon), nil
}

// CMSampleBufferSetInvalidateCallback sets the sample buffer’s invalidation callback.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetInvalidateCallback(_:callback:refcon:)
func CMSampleBufferSetInvalidateCallback(sbuf uintptr, invalidateCallback CMSampleBufferInvalidateCallback, invalidateRefCon uint64) int32 {
	result, callErr := tryCMSampleBufferSetInvalidateCallback(sbuf, invalidateCallback, invalidateRefCon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferSetInvalidateHandler func(sbuf uintptr, invalidateHandler unsafe.Pointer) int32
var _cMSampleBufferSetInvalidateHandlerErr error

func tryCMSampleBufferSetInvalidateHandler(sbuf uintptr, invalidateHandler CMSampleBufferInvalidateHandler) (int32, error) {
	if _cMSampleBufferSetInvalidateHandler == nil {
		return 0, symbolCallError("CMSampleBufferSetInvalidateHandler", "10.10", _cMSampleBufferSetInvalidateHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 objectivec.IObject) { invalidateHandler(arg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _cMSampleBufferSetInvalidateHandler(sbuf, _block0), nil
}

// CMSampleBufferSetInvalidateHandler sets the sample buffer’s invalidation handler.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetInvalidateHandler(_:invalidateHandler:)
func CMSampleBufferSetInvalidateHandler(sbuf uintptr, invalidateHandler CMSampleBufferInvalidateHandler) int32 {
	result, callErr := tryCMSampleBufferSetInvalidateHandler(sbuf, invalidateHandler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferSetOutputPresentationTimeStamp func(sbuf uintptr, outputPresentationTimeStamp CMTime) int32
var _cMSampleBufferSetOutputPresentationTimeStampErr error

func tryCMSampleBufferSetOutputPresentationTimeStamp(sbuf uintptr, outputPresentationTimeStamp CMTime) (int32, error) {
	if _cMSampleBufferSetOutputPresentationTimeStamp == nil {
		return 0, symbolCallError("CMSampleBufferSetOutputPresentationTimeStamp", "10.7", _cMSampleBufferSetOutputPresentationTimeStampErr)
	}
	return _cMSampleBufferSetOutputPresentationTimeStamp(sbuf, outputPresentationTimeStamp), nil
}

// CMSampleBufferSetOutputPresentationTimeStamp sets an output presentation timestamp to use in place of a calculated value.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetOutputPresentationTimeStamp(_:newValue:)
func CMSampleBufferSetOutputPresentationTimeStamp(sbuf uintptr, outputPresentationTimeStamp CMTime) int32 {
	result, callErr := tryCMSampleBufferSetOutputPresentationTimeStamp(sbuf, outputPresentationTimeStamp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSampleBufferTrackDataReadiness func(sbuf uintptr, sampleBufferToTrack uintptr) int32
var _cMSampleBufferTrackDataReadinessErr error

func tryCMSampleBufferTrackDataReadiness(sbuf uintptr, sampleBufferToTrack uintptr) (int32, error) {
	if _cMSampleBufferTrackDataReadiness == nil {
		return 0, symbolCallError("CMSampleBufferTrackDataReadiness", "10.7", _cMSampleBufferTrackDataReadinessErr)
	}
	return _cMSampleBufferTrackDataReadiness(sbuf, sampleBufferToTrack), nil
}

// CMSampleBufferTrackDataReadiness associates a sample buffer’s data readiness with that of another sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferTrackDataReadiness(_:sampleBufferToTrack:)
func CMSampleBufferTrackDataReadiness(sbuf uintptr, sampleBufferToTrack uintptr) int32 {
	result, callErr := tryCMSampleBufferTrackDataReadiness(sbuf, sampleBufferToTrack)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSetAttachment func(target CMAttachmentBearerRef, key corefoundation.CFStringRef, value corefoundation.CFTypeRef, attachmentMode CMAttachmentMode)
var _cMSetAttachmentErr error

func tryCMSetAttachment(target CMAttachmentBearerRef, key corefoundation.CFStringRef, value corefoundation.CFTypeRef, attachmentMode CMAttachmentMode) error {
	if _cMSetAttachment == nil {
		return symbolCallError("CMSetAttachment", "10.7", _cMSetAttachmentErr)
	}
	_cMSetAttachment(target, key, value, attachmentMode)
	return nil
}

// CMSetAttachment sets or adds an attachment to an attachment bearer object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSetAttachment(_:key:value:attachmentMode:)
func CMSetAttachment(target CMAttachmentBearerRef, key corefoundation.CFStringRef, value corefoundation.CFTypeRef, attachmentMode CMAttachmentMode) {
	if callErr := tryCMSetAttachment(target, key, value, attachmentMode); callErr != nil {
		panic(callErr)
	}
}

var _cMSetAttachments func(target CMAttachmentBearerRef, theAttachments corefoundation.CFDictionaryRef, attachmentMode CMAttachmentMode)
var _cMSetAttachmentsErr error

func tryCMSetAttachments(target CMAttachmentBearerRef, theAttachments corefoundation.CFDictionaryRef, attachmentMode CMAttachmentMode) error {
	if _cMSetAttachments == nil {
		return symbolCallError("CMSetAttachments", "10.7", _cMSetAttachmentsErr)
	}
	_cMSetAttachments(target, theAttachments, attachmentMode)
	return nil
}

// CMSetAttachments sets a dictionary of attachments on an attachment bearer object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSetAttachments(_:attachments:attachmentMode:)
func CMSetAttachments(target CMAttachmentBearerRef, theAttachments corefoundation.CFDictionaryRef, attachmentMode CMAttachmentMode) {
	if callErr := tryCMSetAttachments(target, theAttachments, attachmentMode); callErr != nil {
		panic(callErr)
	}
}

var _cMSimpleQueueCreate func(allocator corefoundation.CFAllocatorRef, capacity int32, queueOut *CMSimpleQueueRef) int32
var _cMSimpleQueueCreateErr error

func tryCMSimpleQueueCreate(allocator corefoundation.CFAllocatorRef, capacity int32, queueOut *CMSimpleQueueRef) (int32, error) {
	if _cMSimpleQueueCreate == nil {
		return 0, symbolCallError("CMSimpleQueueCreate", "10.7", _cMSimpleQueueCreateErr)
	}
	return _cMSimpleQueueCreate(allocator, capacity, queueOut), nil
}

// CMSimpleQueueCreate creates a queue that has the specified capacity.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueCreate(allocator:capacity:queueOut:)
func CMSimpleQueueCreate(allocator corefoundation.CFAllocatorRef, capacity int32, queueOut *CMSimpleQueueRef) int32 {
	result, callErr := tryCMSimpleQueueCreate(allocator, capacity, queueOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSimpleQueueDequeue func(queue CMSimpleQueueRef) unsafe.Pointer
var _cMSimpleQueueDequeueErr error

func tryCMSimpleQueueDequeue(queue CMSimpleQueueRef) (unsafe.Pointer, error) {
	if _cMSimpleQueueDequeue == nil {
		return nil, symbolCallError("CMSimpleQueueDequeue", "10.7", _cMSimpleQueueDequeueErr)
	}
	return _cMSimpleQueueDequeue(queue), nil
}

// CMSimpleQueueDequeue dequeues an element from the queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueDequeue(_:)
func CMSimpleQueueDequeue(queue CMSimpleQueueRef) unsafe.Pointer {
	result, callErr := tryCMSimpleQueueDequeue(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSimpleQueueEnqueue func(queue CMSimpleQueueRef, element unsafe.Pointer) int32
var _cMSimpleQueueEnqueueErr error

func tryCMSimpleQueueEnqueue(queue CMSimpleQueueRef, element unsafe.Pointer) (int32, error) {
	if _cMSimpleQueueEnqueue == nil {
		return 0, symbolCallError("CMSimpleQueueEnqueue", "10.7", _cMSimpleQueueEnqueueErr)
	}
	return _cMSimpleQueueEnqueue(queue, element), nil
}

// CMSimpleQueueEnqueue enqueues an element in the queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueEnqueue(_:element:)
func CMSimpleQueueEnqueue(queue CMSimpleQueueRef, element unsafe.Pointer) int32 {
	result, callErr := tryCMSimpleQueueEnqueue(queue, element)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSimpleQueueGetCapacity func(queue CMSimpleQueueRef) int32
var _cMSimpleQueueGetCapacityErr error

func tryCMSimpleQueueGetCapacity(queue CMSimpleQueueRef) (int32, error) {
	if _cMSimpleQueueGetCapacity == nil {
		return 0, symbolCallError("CMSimpleQueueGetCapacity", "10.7", _cMSimpleQueueGetCapacityErr)
	}
	return _cMSimpleQueueGetCapacity(queue), nil
}

// CMSimpleQueueGetCapacity returns the number of elements that the queue can hold.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueGetCapacity(_:)
func CMSimpleQueueGetCapacity(queue CMSimpleQueueRef) int32 {
	result, callErr := tryCMSimpleQueueGetCapacity(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSimpleQueueGetCount func(queue CMSimpleQueueRef) int32
var _cMSimpleQueueGetCountErr error

func tryCMSimpleQueueGetCount(queue CMSimpleQueueRef) (int32, error) {
	if _cMSimpleQueueGetCount == nil {
		return 0, symbolCallError("CMSimpleQueueGetCount", "10.7", _cMSimpleQueueGetCountErr)
	}
	return _cMSimpleQueueGetCount(queue), nil
}

// CMSimpleQueueGetCount returns the number of elements currently in the queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueGetCount(_:)
func CMSimpleQueueGetCount(queue CMSimpleQueueRef) int32 {
	result, callErr := tryCMSimpleQueueGetCount(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSimpleQueueGetHead func(queue CMSimpleQueueRef) unsafe.Pointer
var _cMSimpleQueueGetHeadErr error

func tryCMSimpleQueueGetHead(queue CMSimpleQueueRef) (unsafe.Pointer, error) {
	if _cMSimpleQueueGetHead == nil {
		return nil, symbolCallError("CMSimpleQueueGetHead", "10.7", _cMSimpleQueueGetHeadErr)
	}
	return _cMSimpleQueueGetHead(queue), nil
}

// CMSimpleQueueGetHead returns the element at the head of the queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueGetHead(_:)
func CMSimpleQueueGetHead(queue CMSimpleQueueRef) unsafe.Pointer {
	result, callErr := tryCMSimpleQueueGetHead(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSimpleQueueGetTypeID func() uint
var _cMSimpleQueueGetTypeIDErr error

func tryCMSimpleQueueGetTypeID() (uint, error) {
	if _cMSimpleQueueGetTypeID == nil {
		return 0, symbolCallError("CMSimpleQueueGetTypeID", "10.7", _cMSimpleQueueGetTypeIDErr)
	}
	return _cMSimpleQueueGetTypeID(), nil
}

// CMSimpleQueueGetTypeID returns the type identifier of sample buffer objects.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueGetTypeID()
func CMSimpleQueueGetTypeID() uint {
	result, callErr := tryCMSimpleQueueGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSimpleQueueReset func(queue CMSimpleQueueRef) int32
var _cMSimpleQueueResetErr error

func tryCMSimpleQueueReset(queue CMSimpleQueueRef) (int32, error) {
	if _cMSimpleQueueReset == nil {
		return 0, symbolCallError("CMSimpleQueueReset", "10.7", _cMSimpleQueueResetErr)
	}
	return _cMSimpleQueueReset(queue), nil
}

// CMSimpleQueueReset resets the queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueReset(_:)
func CMSimpleQueueReset(queue CMSimpleQueueRef) int32 {
	result, callErr := tryCMSimpleQueueReset(queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSwapBigEndianClosedCaptionDescriptionToHost func(closedCaptionDescriptionData *uint8, closedCaptionDescriptionSize uintptr) int32
var _cMSwapBigEndianClosedCaptionDescriptionToHostErr error

func tryCMSwapBigEndianClosedCaptionDescriptionToHost(closedCaptionDescriptionData *uint8, closedCaptionDescriptionSize uintptr) (int32, error) {
	if _cMSwapBigEndianClosedCaptionDescriptionToHost == nil {
		return 0, symbolCallError("CMSwapBigEndianClosedCaptionDescriptionToHost", "10.10", _cMSwapBigEndianClosedCaptionDescriptionToHostErr)
	}
	return _cMSwapBigEndianClosedCaptionDescriptionToHost(closedCaptionDescriptionData, closedCaptionDescriptionSize), nil
}

// CMSwapBigEndianClosedCaptionDescriptionToHost converts a closed caption description structure from big-endian to host-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapBigEndianClosedCaptionDescriptionToHost(_:_:)
func CMSwapBigEndianClosedCaptionDescriptionToHost(closedCaptionDescriptionData *uint8, closedCaptionDescriptionSize uintptr) int32 {
	result, callErr := tryCMSwapBigEndianClosedCaptionDescriptionToHost(closedCaptionDescriptionData, closedCaptionDescriptionSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSwapBigEndianImageDescriptionToHost func(imageDescriptionData *uint8, imageDescriptionSize uintptr) int32
var _cMSwapBigEndianImageDescriptionToHostErr error

func tryCMSwapBigEndianImageDescriptionToHost(imageDescriptionData *uint8, imageDescriptionSize uintptr) (int32, error) {
	if _cMSwapBigEndianImageDescriptionToHost == nil {
		return 0, symbolCallError("CMSwapBigEndianImageDescriptionToHost", "10.10", _cMSwapBigEndianImageDescriptionToHostErr)
	}
	return _cMSwapBigEndianImageDescriptionToHost(imageDescriptionData, imageDescriptionSize), nil
}

// CMSwapBigEndianImageDescriptionToHost converts an image description data structure from big-endian to host-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapBigEndianImageDescriptionToHost(_:_:)
func CMSwapBigEndianImageDescriptionToHost(imageDescriptionData *uint8, imageDescriptionSize uintptr) int32 {
	result, callErr := tryCMSwapBigEndianImageDescriptionToHost(imageDescriptionData, imageDescriptionSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSwapBigEndianMetadataDescriptionToHost func(metadataDescriptionData *uint8, metadataDescriptionSize uintptr) int32
var _cMSwapBigEndianMetadataDescriptionToHostErr error

func tryCMSwapBigEndianMetadataDescriptionToHost(metadataDescriptionData *uint8, metadataDescriptionSize uintptr) (int32, error) {
	if _cMSwapBigEndianMetadataDescriptionToHost == nil {
		return 0, symbolCallError("CMSwapBigEndianMetadataDescriptionToHost", "10.10", _cMSwapBigEndianMetadataDescriptionToHostErr)
	}
	return _cMSwapBigEndianMetadataDescriptionToHost(metadataDescriptionData, metadataDescriptionSize), nil
}

// CMSwapBigEndianMetadataDescriptionToHost converts a metadata description data structure from big-endian to host-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapBigEndianMetadataDescriptionToHost(_:_:)
func CMSwapBigEndianMetadataDescriptionToHost(metadataDescriptionData *uint8, metadataDescriptionSize uintptr) int32 {
	result, callErr := tryCMSwapBigEndianMetadataDescriptionToHost(metadataDescriptionData, metadataDescriptionSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSwapBigEndianSoundDescriptionToHost func(soundDescriptionData *uint8, soundDescriptionSize uintptr) int32
var _cMSwapBigEndianSoundDescriptionToHostErr error

func tryCMSwapBigEndianSoundDescriptionToHost(soundDescriptionData *uint8, soundDescriptionSize uintptr) (int32, error) {
	if _cMSwapBigEndianSoundDescriptionToHost == nil {
		return 0, symbolCallError("CMSwapBigEndianSoundDescriptionToHost", "10.10", _cMSwapBigEndianSoundDescriptionToHostErr)
	}
	return _cMSwapBigEndianSoundDescriptionToHost(soundDescriptionData, soundDescriptionSize), nil
}

// CMSwapBigEndianSoundDescriptionToHost converts a sound description data structure from big-endian to host-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapBigEndianSoundDescriptionToHost(_:_:)
func CMSwapBigEndianSoundDescriptionToHost(soundDescriptionData *uint8, soundDescriptionSize uintptr) int32 {
	result, callErr := tryCMSwapBigEndianSoundDescriptionToHost(soundDescriptionData, soundDescriptionSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSwapBigEndianTextDescriptionToHost func(textDescriptionData *uint8, textDescriptionSize uintptr) int32
var _cMSwapBigEndianTextDescriptionToHostErr error

func tryCMSwapBigEndianTextDescriptionToHost(textDescriptionData *uint8, textDescriptionSize uintptr) (int32, error) {
	if _cMSwapBigEndianTextDescriptionToHost == nil {
		return 0, symbolCallError("CMSwapBigEndianTextDescriptionToHost", "10.10", _cMSwapBigEndianTextDescriptionToHostErr)
	}
	return _cMSwapBigEndianTextDescriptionToHost(textDescriptionData, textDescriptionSize), nil
}

// CMSwapBigEndianTextDescriptionToHost converts a text description structure from big-endian to host-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapBigEndianTextDescriptionToHost(_:_:)
func CMSwapBigEndianTextDescriptionToHost(textDescriptionData *uint8, textDescriptionSize uintptr) int32 {
	result, callErr := tryCMSwapBigEndianTextDescriptionToHost(textDescriptionData, textDescriptionSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSwapBigEndianTimeCodeDescriptionToHost func(timeCodeDescriptionData *uint8, timeCodeDescriptionSize uintptr) int32
var _cMSwapBigEndianTimeCodeDescriptionToHostErr error

func tryCMSwapBigEndianTimeCodeDescriptionToHost(timeCodeDescriptionData *uint8, timeCodeDescriptionSize uintptr) (int32, error) {
	if _cMSwapBigEndianTimeCodeDescriptionToHost == nil {
		return 0, symbolCallError("CMSwapBigEndianTimeCodeDescriptionToHost", "10.10", _cMSwapBigEndianTimeCodeDescriptionToHostErr)
	}
	return _cMSwapBigEndianTimeCodeDescriptionToHost(timeCodeDescriptionData, timeCodeDescriptionSize), nil
}

// CMSwapBigEndianTimeCodeDescriptionToHost converts a time code description data structure from big-endian to host-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapBigEndianTimeCodeDescriptionToHost(_:_:)
func CMSwapBigEndianTimeCodeDescriptionToHost(timeCodeDescriptionData *uint8, timeCodeDescriptionSize uintptr) int32 {
	result, callErr := tryCMSwapBigEndianTimeCodeDescriptionToHost(timeCodeDescriptionData, timeCodeDescriptionSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSwapHostEndianClosedCaptionDescriptionToBig func(closedCaptionDescriptionData *uint8, closedCaptionDescriptionSize uintptr) int32
var _cMSwapHostEndianClosedCaptionDescriptionToBigErr error

func tryCMSwapHostEndianClosedCaptionDescriptionToBig(closedCaptionDescriptionData *uint8, closedCaptionDescriptionSize uintptr) (int32, error) {
	if _cMSwapHostEndianClosedCaptionDescriptionToBig == nil {
		return 0, symbolCallError("CMSwapHostEndianClosedCaptionDescriptionToBig", "10.10", _cMSwapHostEndianClosedCaptionDescriptionToBigErr)
	}
	return _cMSwapHostEndianClosedCaptionDescriptionToBig(closedCaptionDescriptionData, closedCaptionDescriptionSize), nil
}

// CMSwapHostEndianClosedCaptionDescriptionToBig converts a closed caption description structure from host-endian to big-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapHostEndianClosedCaptionDescriptionToBig(_:_:)
func CMSwapHostEndianClosedCaptionDescriptionToBig(closedCaptionDescriptionData *uint8, closedCaptionDescriptionSize uintptr) int32 {
	result, callErr := tryCMSwapHostEndianClosedCaptionDescriptionToBig(closedCaptionDescriptionData, closedCaptionDescriptionSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSwapHostEndianImageDescriptionToBig func(imageDescriptionData *uint8, imageDescriptionSize uintptr) int32
var _cMSwapHostEndianImageDescriptionToBigErr error

func tryCMSwapHostEndianImageDescriptionToBig(imageDescriptionData *uint8, imageDescriptionSize uintptr) (int32, error) {
	if _cMSwapHostEndianImageDescriptionToBig == nil {
		return 0, symbolCallError("CMSwapHostEndianImageDescriptionToBig", "10.10", _cMSwapHostEndianImageDescriptionToBigErr)
	}
	return _cMSwapHostEndianImageDescriptionToBig(imageDescriptionData, imageDescriptionSize), nil
}

// CMSwapHostEndianImageDescriptionToBig converts an image description data structure from host-endian to big-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapHostEndianImageDescriptionToBig(_:_:)
func CMSwapHostEndianImageDescriptionToBig(imageDescriptionData *uint8, imageDescriptionSize uintptr) int32 {
	result, callErr := tryCMSwapHostEndianImageDescriptionToBig(imageDescriptionData, imageDescriptionSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSwapHostEndianMetadataDescriptionToBig func(metadataDescriptionData *uint8, metadataDescriptionSize uintptr) int32
var _cMSwapHostEndianMetadataDescriptionToBigErr error

func tryCMSwapHostEndianMetadataDescriptionToBig(metadataDescriptionData *uint8, metadataDescriptionSize uintptr) (int32, error) {
	if _cMSwapHostEndianMetadataDescriptionToBig == nil {
		return 0, symbolCallError("CMSwapHostEndianMetadataDescriptionToBig", "10.10", _cMSwapHostEndianMetadataDescriptionToBigErr)
	}
	return _cMSwapHostEndianMetadataDescriptionToBig(metadataDescriptionData, metadataDescriptionSize), nil
}

// CMSwapHostEndianMetadataDescriptionToBig converts a metadata description data structure from host-endian to big-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapHostEndianMetadataDescriptionToBig(_:_:)
func CMSwapHostEndianMetadataDescriptionToBig(metadataDescriptionData *uint8, metadataDescriptionSize uintptr) int32 {
	result, callErr := tryCMSwapHostEndianMetadataDescriptionToBig(metadataDescriptionData, metadataDescriptionSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSwapHostEndianSoundDescriptionToBig func(soundDescriptionData *uint8, soundDescriptionSize uintptr) int32
var _cMSwapHostEndianSoundDescriptionToBigErr error

func tryCMSwapHostEndianSoundDescriptionToBig(soundDescriptionData *uint8, soundDescriptionSize uintptr) (int32, error) {
	if _cMSwapHostEndianSoundDescriptionToBig == nil {
		return 0, symbolCallError("CMSwapHostEndianSoundDescriptionToBig", "10.10", _cMSwapHostEndianSoundDescriptionToBigErr)
	}
	return _cMSwapHostEndianSoundDescriptionToBig(soundDescriptionData, soundDescriptionSize), nil
}

// CMSwapHostEndianSoundDescriptionToBig converts a sound description data structure from host-endian to big-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapHostEndianSoundDescriptionToBig(_:_:)
func CMSwapHostEndianSoundDescriptionToBig(soundDescriptionData *uint8, soundDescriptionSize uintptr) int32 {
	result, callErr := tryCMSwapHostEndianSoundDescriptionToBig(soundDescriptionData, soundDescriptionSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSwapHostEndianTextDescriptionToBig func(textDescriptionData *uint8, textDescriptionSize uintptr) int32
var _cMSwapHostEndianTextDescriptionToBigErr error

func tryCMSwapHostEndianTextDescriptionToBig(textDescriptionData *uint8, textDescriptionSize uintptr) (int32, error) {
	if _cMSwapHostEndianTextDescriptionToBig == nil {
		return 0, symbolCallError("CMSwapHostEndianTextDescriptionToBig", "10.10", _cMSwapHostEndianTextDescriptionToBigErr)
	}
	return _cMSwapHostEndianTextDescriptionToBig(textDescriptionData, textDescriptionSize), nil
}

// CMSwapHostEndianTextDescriptionToBig converts a text description structure from host-endian to big-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapHostEndianTextDescriptionToBig(_:_:)
func CMSwapHostEndianTextDescriptionToBig(textDescriptionData *uint8, textDescriptionSize uintptr) int32 {
	result, callErr := tryCMSwapHostEndianTextDescriptionToBig(textDescriptionData, textDescriptionSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSwapHostEndianTimeCodeDescriptionToBig func(timeCodeDescriptionData *uint8, timeCodeDescriptionSize uintptr) int32
var _cMSwapHostEndianTimeCodeDescriptionToBigErr error

func tryCMSwapHostEndianTimeCodeDescriptionToBig(timeCodeDescriptionData *uint8, timeCodeDescriptionSize uintptr) (int32, error) {
	if _cMSwapHostEndianTimeCodeDescriptionToBig == nil {
		return 0, symbolCallError("CMSwapHostEndianTimeCodeDescriptionToBig", "10.10", _cMSwapHostEndianTimeCodeDescriptionToBigErr)
	}
	return _cMSwapHostEndianTimeCodeDescriptionToBig(timeCodeDescriptionData, timeCodeDescriptionSize), nil
}

// CMSwapHostEndianTimeCodeDescriptionToBig converts a time code description data structure from host-endian to big-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapHostEndianTimeCodeDescriptionToBig(_:_:)
func CMSwapHostEndianTimeCodeDescriptionToBig(timeCodeDescriptionData *uint8, timeCodeDescriptionSize uintptr) int32 {
	result, callErr := tryCMSwapHostEndianTimeCodeDescriptionToBig(timeCodeDescriptionData, timeCodeDescriptionSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSyncConvertTime func(time CMTime, fromClockOrTimebase CMClockOrTimebaseRef, toClockOrTimebase CMClockOrTimebaseRef) CMTime
var _cMSyncConvertTimeErr error

func tryCMSyncConvertTime(time CMTime, fromClockOrTimebase CMClockOrTimebaseRef, toClockOrTimebase CMClockOrTimebaseRef) (CMTime, error) {
	if _cMSyncConvertTime == nil {
		return CMTime{}, symbolCallError("CMSyncConvertTime", "10.8", _cMSyncConvertTimeErr)
	}
	return _cMSyncConvertTime(time, fromClockOrTimebase, toClockOrTimebase), nil
}

// CMSyncConvertTime converts a time from one timebase or clock to another timebase or clock.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSyncConvertTime(_:from:to:)
func CMSyncConvertTime(time CMTime, fromClockOrTimebase CMClockOrTimebaseRef, toClockOrTimebase CMClockOrTimebaseRef) CMTime {
	result, callErr := tryCMSyncConvertTime(time, fromClockOrTimebase, toClockOrTimebase)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSyncGetRelativeRate func(ofClockOrTimebase CMClockOrTimebaseRef, relativeToClockOrTimebase CMClockOrTimebaseRef) float64
var _cMSyncGetRelativeRateErr error

func tryCMSyncGetRelativeRate(ofClockOrTimebase CMClockOrTimebaseRef, relativeToClockOrTimebase CMClockOrTimebaseRef) (float64, error) {
	if _cMSyncGetRelativeRate == nil {
		return 0.0, symbolCallError("CMSyncGetRelativeRate", "10.8", _cMSyncGetRelativeRateErr)
	}
	return _cMSyncGetRelativeRate(ofClockOrTimebase, relativeToClockOrTimebase), nil
}

// CMSyncGetRelativeRate returns the relative rate of one timebase or clock relative to another timebase or clock.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSyncGetRelativeRate(_:relativeTo:)
func CMSyncGetRelativeRate(ofClockOrTimebase CMClockOrTimebaseRef, relativeToClockOrTimebase CMClockOrTimebaseRef) float64 {
	result, callErr := tryCMSyncGetRelativeRate(ofClockOrTimebase, relativeToClockOrTimebase)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSyncGetRelativeRateAndAnchorTime func(ofClockOrTimebase CMClockOrTimebaseRef, relativeToClockOrTimebase CMClockOrTimebaseRef, outRelativeRate *float64, outOfClockOrTimebaseAnchorTime *CMTime, outRelativeToClockOrTimebaseAnchorTime *CMTime) int32
var _cMSyncGetRelativeRateAndAnchorTimeErr error

func tryCMSyncGetRelativeRateAndAnchorTime(ofClockOrTimebase CMClockOrTimebaseRef, relativeToClockOrTimebase CMClockOrTimebaseRef, outRelativeRate *float64, outOfClockOrTimebaseAnchorTime *CMTime, outRelativeToClockOrTimebaseAnchorTime *CMTime) (int32, error) {
	if _cMSyncGetRelativeRateAndAnchorTime == nil {
		return 0, symbolCallError("CMSyncGetRelativeRateAndAnchorTime", "10.8", _cMSyncGetRelativeRateAndAnchorTimeErr)
	}
	return _cMSyncGetRelativeRateAndAnchorTime(ofClockOrTimebase, relativeToClockOrTimebase, outRelativeRate, outOfClockOrTimebaseAnchorTime, outRelativeToClockOrTimebaseAnchorTime), nil
}

// CMSyncGetRelativeRateAndAnchorTime returns the relative rate of one timebase or clock relative to another timebase or clock and the times of each timebase or clock at which the relative rate went into effect.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSyncGetRelativeRateAndAnchorTime(_:relativeTo:relativeRateOut:anchorTimeOut:relativeToAnchorTimeOut:)
func CMSyncGetRelativeRateAndAnchorTime(ofClockOrTimebase CMClockOrTimebaseRef, relativeToClockOrTimebase CMClockOrTimebaseRef, outRelativeRate *float64, outOfClockOrTimebaseAnchorTime *CMTime, outRelativeToClockOrTimebaseAnchorTime *CMTime) int32 {
	result, callErr := tryCMSyncGetRelativeRateAndAnchorTime(ofClockOrTimebase, relativeToClockOrTimebase, outRelativeRate, outOfClockOrTimebaseAnchorTime, outRelativeToClockOrTimebaseAnchorTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSyncGetTime func(clockOrTimebase CMClockOrTimebaseRef) CMTime
var _cMSyncGetTimeErr error

func tryCMSyncGetTime(clockOrTimebase CMClockOrTimebaseRef) (CMTime, error) {
	if _cMSyncGetTime == nil {
		return CMTime{}, symbolCallError("CMSyncGetTime", "10.8", _cMSyncGetTimeErr)
	}
	return _cMSyncGetTime(clockOrTimebase), nil
}

// CMSyncGetTime returns the time from a clock or timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSyncGetTime(_:)
func CMSyncGetTime(clockOrTimebase CMClockOrTimebaseRef) CMTime {
	result, callErr := tryCMSyncGetTime(clockOrTimebase)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMSyncMightDrift func(clockOrTimebase1 CMClockOrTimebaseRef, clockOrTimebase2 CMClockOrTimebaseRef) bool
var _cMSyncMightDriftErr error

func tryCMSyncMightDrift(clockOrTimebase1 CMClockOrTimebaseRef, clockOrTimebase2 CMClockOrTimebaseRef) (bool, error) {
	if _cMSyncMightDrift == nil {
		return false, symbolCallError("CMSyncMightDrift", "10.8", _cMSyncMightDriftErr)
	}
	return _cMSyncMightDrift(clockOrTimebase1, clockOrTimebase2), nil
}

// CMSyncMightDrift returns a Boolean value that indicates whether it’s possible for one timebase or clock to drift relative to the other.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSyncMightDrift(_:_:)
func CMSyncMightDrift(clockOrTimebase1 CMClockOrTimebaseRef, clockOrTimebase2 CMClockOrTimebaseRef) bool {
	result, callErr := tryCMSyncMightDrift(clockOrTimebase1, clockOrTimebase2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionAddTag func(tagCollection CMMutableTagCollectionRef, tagToAdd CMTag) int32
var _cMTagCollectionAddTagErr error

func tryCMTagCollectionAddTag(tagCollection CMMutableTagCollectionRef, tagToAdd CMTag) (int32, error) {
	if _cMTagCollectionAddTag == nil {
		return 0, symbolCallError("CMTagCollectionAddTag", "14.0", _cMTagCollectionAddTagErr)
	}
	return _cMTagCollectionAddTag(tagCollection, tagToAdd), nil
}

// CMTagCollectionAddTag adds a new tag to an existing collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionAddTag
func CMTagCollectionAddTag(tagCollection CMMutableTagCollectionRef, tagToAdd CMTag) int32 {
	result, callErr := tryCMTagCollectionAddTag(tagCollection, tagToAdd)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionAddTagsFromArray func(tagCollection CMMutableTagCollectionRef, tags *CMTag, tagCount int) int32
var _cMTagCollectionAddTagsFromArrayErr error

func tryCMTagCollectionAddTagsFromArray(tagCollection CMMutableTagCollectionRef, tags *CMTag, tagCount int) (int32, error) {
	if _cMTagCollectionAddTagsFromArray == nil {
		return 0, symbolCallError("CMTagCollectionAddTagsFromArray", "14.0", _cMTagCollectionAddTagsFromArrayErr)
	}
	return _cMTagCollectionAddTagsFromArray(tagCollection, tags, tagCount), nil
}

// CMTagCollectionAddTagsFromArray adds the tags contained in a C-style array to a tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionAddTagsFromArray
func CMTagCollectionAddTagsFromArray(tagCollection CMMutableTagCollectionRef, tags *CMTag, tagCount int) int32 {
	result, callErr := tryCMTagCollectionAddTagsFromArray(tagCollection, tags, tagCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionAddTagsFromCollection func(tagCollection CMMutableTagCollectionRef, collectionWithTagsToAdd CMTagCollectionRef) int32
var _cMTagCollectionAddTagsFromCollectionErr error

func tryCMTagCollectionAddTagsFromCollection(tagCollection CMMutableTagCollectionRef, collectionWithTagsToAdd CMTagCollectionRef) (int32, error) {
	if _cMTagCollectionAddTagsFromCollection == nil {
		return 0, symbolCallError("CMTagCollectionAddTagsFromCollection", "14.0", _cMTagCollectionAddTagsFromCollectionErr)
	}
	return _cMTagCollectionAddTagsFromCollection(tagCollection, collectionWithTagsToAdd), nil
}

// CMTagCollectionAddTagsFromCollection add the tags contained in one collection to another.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionAddTagsFromCollection
func CMTagCollectionAddTagsFromCollection(tagCollection CMMutableTagCollectionRef, collectionWithTagsToAdd CMTagCollectionRef) int32 {
	result, callErr := tryCMTagCollectionAddTagsFromCollection(tagCollection, collectionWithTagsToAdd)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionApply func(tagCollection CMTagCollectionRef, applier CMTagCollectionApplierFunction, context unsafe.Pointer)
var _cMTagCollectionApplyErr error

func tryCMTagCollectionApply(tagCollection CMTagCollectionRef, applier CMTagCollectionApplierFunction, context unsafe.Pointer) error {
	if _cMTagCollectionApply == nil {
		return symbolCallError("CMTagCollectionApply", "14.0", _cMTagCollectionApplyErr)
	}
	_cMTagCollectionApply(tagCollection, applier, context)
	return nil
}

// CMTagCollectionApply applies a function to all tags in a collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionApply
func CMTagCollectionApply(tagCollection CMTagCollectionRef, applier CMTagCollectionApplierFunction, context unsafe.Pointer) {
	if callErr := tryCMTagCollectionApply(tagCollection, applier, context); callErr != nil {
		panic(callErr)
	}
}

var _cMTagCollectionApplyUntil func(tagCollection CMTagCollectionRef, applier CMTagCollectionTagFilterFunction, context unsafe.Pointer) CMTag
var _cMTagCollectionApplyUntilErr error

func tryCMTagCollectionApplyUntil(tagCollection CMTagCollectionRef, applier CMTagCollectionTagFilterFunction, context unsafe.Pointer) (CMTag, error) {
	if _cMTagCollectionApplyUntil == nil {
		return CMTag{}, symbolCallError("CMTagCollectionApplyUntil", "14.0", _cMTagCollectionApplyUntilErr)
	}
	return _cMTagCollectionApplyUntil(tagCollection, applier, context), nil
}

// CMTagCollectionApplyUntil applies a Boolean function to tags in a collection, stopping when it returns true.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionApplyUntil
func CMTagCollectionApplyUntil(tagCollection CMTagCollectionRef, applier CMTagCollectionTagFilterFunction, context unsafe.Pointer) CMTag {
	result, callErr := tryCMTagCollectionApplyUntil(tagCollection, applier, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionContainsCategory func(tagCollection CMTagCollectionRef, category CMTagCategory) bool
var _cMTagCollectionContainsCategoryErr error

func tryCMTagCollectionContainsCategory(tagCollection CMTagCollectionRef, category CMTagCategory) (bool, error) {
	if _cMTagCollectionContainsCategory == nil {
		return false, symbolCallError("CMTagCollectionContainsCategory", "14.0", _cMTagCollectionContainsCategoryErr)
	}
	return _cMTagCollectionContainsCategory(tagCollection, category), nil
}

// CMTagCollectionContainsCategory determines if a tag collection contains tags for a given category.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionContainsCategory
func CMTagCollectionContainsCategory(tagCollection CMTagCollectionRef, category CMTagCategory) bool {
	result, callErr := tryCMTagCollectionContainsCategory(tagCollection, category)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionContainsSpecifiedTags func(tagCollection CMTagCollectionRef, containedTags *CMTag, containedTagCount int) bool
var _cMTagCollectionContainsSpecifiedTagsErr error

func tryCMTagCollectionContainsSpecifiedTags(tagCollection CMTagCollectionRef, containedTags *CMTag, containedTagCount int) (bool, error) {
	if _cMTagCollectionContainsSpecifiedTags == nil {
		return false, symbolCallError("CMTagCollectionContainsSpecifiedTags", "14.0", _cMTagCollectionContainsSpecifiedTagsErr)
	}
	return _cMTagCollectionContainsSpecifiedTags(tagCollection, containedTags, containedTagCount), nil
}

// CMTagCollectionContainsSpecifiedTags determines if a tag collection contains a subset of tags.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionContainsSpecifiedTags
func CMTagCollectionContainsSpecifiedTags(tagCollection CMTagCollectionRef, containedTags *CMTag, containedTagCount int) bool {
	result, callErr := tryCMTagCollectionContainsSpecifiedTags(tagCollection, containedTags, containedTagCount)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionContainsTag func(tagCollection CMTagCollectionRef, tag CMTag) bool
var _cMTagCollectionContainsTagErr error

func tryCMTagCollectionContainsTag(tagCollection CMTagCollectionRef, tag CMTag) (bool, error) {
	if _cMTagCollectionContainsTag == nil {
		return false, symbolCallError("CMTagCollectionContainsTag", "14.0", _cMTagCollectionContainsTagErr)
	}
	return _cMTagCollectionContainsTag(tagCollection, tag), nil
}

// CMTagCollectionContainsTag determines if a tag collection contains a specific tag.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionContainsTag
func CMTagCollectionContainsTag(tagCollection CMTagCollectionRef, tag CMTag) bool {
	result, callErr := tryCMTagCollectionContainsTag(tagCollection, tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionContainsTagsOfCollection func(tagCollection CMTagCollectionRef, containedTagCollection CMTagCollectionRef) bool
var _cMTagCollectionContainsTagsOfCollectionErr error

func tryCMTagCollectionContainsTagsOfCollection(tagCollection CMTagCollectionRef, containedTagCollection CMTagCollectionRef) (bool, error) {
	if _cMTagCollectionContainsTagsOfCollection == nil {
		return false, symbolCallError("CMTagCollectionContainsTagsOfCollection", "14.0", _cMTagCollectionContainsTagsOfCollectionErr)
	}
	return _cMTagCollectionContainsTagsOfCollection(tagCollection, containedTagCollection), nil
}

// CMTagCollectionContainsTagsOfCollection determines if one collection of tags contains every tag from another collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionContainsTagsOfCollection
func CMTagCollectionContainsTagsOfCollection(tagCollection CMTagCollectionRef, containedTagCollection CMTagCollectionRef) bool {
	result, callErr := tryCMTagCollectionContainsTagsOfCollection(tagCollection, containedTagCollection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionCopyAsData func(tagCollection CMTagCollectionRef, allocator corefoundation.CFAllocatorRef) corefoundation.CFDataRef
var _cMTagCollectionCopyAsDataErr error

func tryCMTagCollectionCopyAsData(tagCollection CMTagCollectionRef, allocator corefoundation.CFAllocatorRef) (corefoundation.CFDataRef, error) {
	if _cMTagCollectionCopyAsData == nil {
		return 0, symbolCallError("CMTagCollectionCopyAsData", "14.0", _cMTagCollectionCopyAsDataErr)
	}
	return _cMTagCollectionCopyAsData(tagCollection, allocator), nil
}

// CMTagCollectionCopyAsData creates a new Core Foundation data instance from a tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCopyAsData
func CMTagCollectionCopyAsData(tagCollection CMTagCollectionRef, allocator corefoundation.CFAllocatorRef) corefoundation.CFDataRef {
	result, callErr := tryCMTagCollectionCopyAsData(tagCollection, allocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionCopyAsDictionary func(tagCollection CMTagCollectionRef, allocator corefoundation.CFAllocatorRef) corefoundation.CFDictionaryRef
var _cMTagCollectionCopyAsDictionaryErr error

func tryCMTagCollectionCopyAsDictionary(tagCollection CMTagCollectionRef, allocator corefoundation.CFAllocatorRef) (corefoundation.CFDictionaryRef, error) {
	if _cMTagCollectionCopyAsDictionary == nil {
		return 0, symbolCallError("CMTagCollectionCopyAsDictionary", "14.0", _cMTagCollectionCopyAsDictionaryErr)
	}
	return _cMTagCollectionCopyAsDictionary(tagCollection, allocator), nil
}

// CMTagCollectionCopyAsDictionary creates a new Core Foundation dictionary from a tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCopyAsDictionary
func CMTagCollectionCopyAsDictionary(tagCollection CMTagCollectionRef, allocator corefoundation.CFAllocatorRef) corefoundation.CFDictionaryRef {
	result, callErr := tryCMTagCollectionCopyAsDictionary(tagCollection, allocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionCopyDescription func(allocator corefoundation.CFAllocatorRef, tagCollection CMTagCollectionRef) corefoundation.CFStringRef
var _cMTagCollectionCopyDescriptionErr error

func tryCMTagCollectionCopyDescription(allocator corefoundation.CFAllocatorRef, tagCollection CMTagCollectionRef) (corefoundation.CFStringRef, error) {
	if _cMTagCollectionCopyDescription == nil {
		return 0, symbolCallError("CMTagCollectionCopyDescription", "14.0", _cMTagCollectionCopyDescriptionErr)
	}
	return _cMTagCollectionCopyDescription(allocator, tagCollection), nil
}

// CMTagCollectionCopyDescription retrieves a copy of the tag collection’s description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCopyDescription
func CMTagCollectionCopyDescription(allocator corefoundation.CFAllocatorRef, tagCollection CMTagCollectionRef) corefoundation.CFStringRef {
	result, callErr := tryCMTagCollectionCopyDescription(allocator, tagCollection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionCopyTagsOfCategories func(allocator corefoundation.CFAllocatorRef, tagCollection CMTagCollectionRef, categories *CMTagCategory, categoriesCount int, collectionWithTagsOfCategories *CMTagCollectionRef) int32
var _cMTagCollectionCopyTagsOfCategoriesErr error

func tryCMTagCollectionCopyTagsOfCategories(allocator corefoundation.CFAllocatorRef, tagCollection CMTagCollectionRef, categories *CMTagCategory, categoriesCount int, collectionWithTagsOfCategories *CMTagCollectionRef) (int32, error) {
	if _cMTagCollectionCopyTagsOfCategories == nil {
		return 0, symbolCallError("CMTagCollectionCopyTagsOfCategories", "14.0", _cMTagCollectionCopyTagsOfCategoriesErr)
	}
	return _cMTagCollectionCopyTagsOfCategories(allocator, tagCollection, categories, categoriesCount, collectionWithTagsOfCategories), nil
}

// CMTagCollectionCopyTagsOfCategories creates a new tag collection from an existing collection, copying all tags which match a list of categories.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCopyTagsOfCategories
func CMTagCollectionCopyTagsOfCategories(allocator corefoundation.CFAllocatorRef, tagCollection CMTagCollectionRef, categories *CMTagCategory, categoriesCount int, collectionWithTagsOfCategories *CMTagCollectionRef) int32 {
	result, callErr := tryCMTagCollectionCopyTagsOfCategories(allocator, tagCollection, categories, categoriesCount, collectionWithTagsOfCategories)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionCountTagsWithFilterFunction func(tagCollection CMTagCollectionRef, filterApplier CMTagCollectionTagFilterFunction, context unsafe.Pointer) int
var _cMTagCollectionCountTagsWithFilterFunctionErr error

func tryCMTagCollectionCountTagsWithFilterFunction(tagCollection CMTagCollectionRef, filterApplier CMTagCollectionTagFilterFunction, context unsafe.Pointer) (int, error) {
	if _cMTagCollectionCountTagsWithFilterFunction == nil {
		return 0, symbolCallError("CMTagCollectionCountTagsWithFilterFunction", "14.0", _cMTagCollectionCountTagsWithFilterFunctionErr)
	}
	return _cMTagCollectionCountTagsWithFilterFunction(tagCollection, filterApplier, context), nil
}

// CMTagCollectionCountTagsWithFilterFunction counts the number of tags in a collection matching an evaluation function.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCountTagsWithFilterFunction
func CMTagCollectionCountTagsWithFilterFunction(tagCollection CMTagCollectionRef, filterApplier CMTagCollectionTagFilterFunction, context unsafe.Pointer) int {
	result, callErr := tryCMTagCollectionCountTagsWithFilterFunction(tagCollection, filterApplier, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionCreate func(allocator corefoundation.CFAllocatorRef, tags *CMTag, tagCount int, newCollectionOut *CMTagCollectionRef) int32
var _cMTagCollectionCreateErr error

func tryCMTagCollectionCreate(allocator corefoundation.CFAllocatorRef, tags *CMTag, tagCount int, newCollectionOut *CMTagCollectionRef) (int32, error) {
	if _cMTagCollectionCreate == nil {
		return 0, symbolCallError("CMTagCollectionCreate", "14.0", _cMTagCollectionCreateErr)
	}
	return _cMTagCollectionCreate(allocator, tags, tagCount, newCollectionOut), nil
}

// CMTagCollectionCreate creates a new tag collection from an existing C-style array of tags.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreate
func CMTagCollectionCreate(allocator corefoundation.CFAllocatorRef, tags *CMTag, tagCount int, newCollectionOut *CMTagCollectionRef) int32 {
	result, callErr := tryCMTagCollectionCreate(allocator, tags, tagCount, newCollectionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionCreateCopy func(tagCollection CMTagCollectionRef, allocator corefoundation.CFAllocatorRef, newCollectionCopyOut *CMTagCollectionRef) int32
var _cMTagCollectionCreateCopyErr error

func tryCMTagCollectionCreateCopy(tagCollection CMTagCollectionRef, allocator corefoundation.CFAllocatorRef, newCollectionCopyOut *CMTagCollectionRef) (int32, error) {
	if _cMTagCollectionCreateCopy == nil {
		return 0, symbolCallError("CMTagCollectionCreateCopy", "14.0", _cMTagCollectionCreateCopyErr)
	}
	return _cMTagCollectionCreateCopy(tagCollection, allocator, newCollectionCopyOut), nil
}

// CMTagCollectionCreateCopy creates a copy of a tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateCopy
func CMTagCollectionCreateCopy(tagCollection CMTagCollectionRef, allocator corefoundation.CFAllocatorRef, newCollectionCopyOut *CMTagCollectionRef) int32 {
	result, callErr := tryCMTagCollectionCreateCopy(tagCollection, allocator, newCollectionCopyOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionCreateDifference func(tagCollectionMinuend CMTagCollectionRef, tagCollectionSubtrahend CMTagCollectionRef, tagCollectionOut *CMTagCollectionRef) int32
var _cMTagCollectionCreateDifferenceErr error

func tryCMTagCollectionCreateDifference(tagCollectionMinuend CMTagCollectionRef, tagCollectionSubtrahend CMTagCollectionRef, tagCollectionOut *CMTagCollectionRef) (int32, error) {
	if _cMTagCollectionCreateDifference == nil {
		return 0, symbolCallError("CMTagCollectionCreateDifference", "14.0", _cMTagCollectionCreateDifferenceErr)
	}
	return _cMTagCollectionCreateDifference(tagCollectionMinuend, tagCollectionSubtrahend, tagCollectionOut), nil
}

// CMTagCollectionCreateDifference creates a new tag collection with the difference of two existing collections.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateDifference
func CMTagCollectionCreateDifference(tagCollectionMinuend CMTagCollectionRef, tagCollectionSubtrahend CMTagCollectionRef, tagCollectionOut *CMTagCollectionRef) int32 {
	result, callErr := tryCMTagCollectionCreateDifference(tagCollectionMinuend, tagCollectionSubtrahend, tagCollectionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionCreateExclusiveOr func(tagCollection1 CMTagCollectionRef, tagCollection2 CMTagCollectionRef, tagCollectionOut *CMTagCollectionRef) int32
var _cMTagCollectionCreateExclusiveOrErr error

func tryCMTagCollectionCreateExclusiveOr(tagCollection1 CMTagCollectionRef, tagCollection2 CMTagCollectionRef, tagCollectionOut *CMTagCollectionRef) (int32, error) {
	if _cMTagCollectionCreateExclusiveOr == nil {
		return 0, symbolCallError("CMTagCollectionCreateExclusiveOr", "14.0", _cMTagCollectionCreateExclusiveOrErr)
	}
	return _cMTagCollectionCreateExclusiveOr(tagCollection1, tagCollection2, tagCollectionOut), nil
}

// CMTagCollectionCreateExclusiveOr creates a new tag collection from two existing tag collections, copying elements which are in one collection or the other, but not both.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateExclusiveOr
func CMTagCollectionCreateExclusiveOr(tagCollection1 CMTagCollectionRef, tagCollection2 CMTagCollectionRef, tagCollectionOut *CMTagCollectionRef) int32 {
	result, callErr := tryCMTagCollectionCreateExclusiveOr(tagCollection1, tagCollection2, tagCollectionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionCreateFromData func(data corefoundation.CFDataRef, allocator corefoundation.CFAllocatorRef, newCollectionOut *CMTagCollectionRef) int32
var _cMTagCollectionCreateFromDataErr error

func tryCMTagCollectionCreateFromData(data corefoundation.CFDataRef, allocator corefoundation.CFAllocatorRef, newCollectionOut *CMTagCollectionRef) (int32, error) {
	if _cMTagCollectionCreateFromData == nil {
		return 0, symbolCallError("CMTagCollectionCreateFromData", "14.0", _cMTagCollectionCreateFromDataErr)
	}
	return _cMTagCollectionCreateFromData(data, allocator, newCollectionOut), nil
}

// CMTagCollectionCreateFromData creates a new tag collection from an existing Core Foundation data instance.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateFromData
func CMTagCollectionCreateFromData(data corefoundation.CFDataRef, allocator corefoundation.CFAllocatorRef, newCollectionOut *CMTagCollectionRef) int32 {
	result, callErr := tryCMTagCollectionCreateFromData(data, allocator, newCollectionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionCreateFromDictionary func(dict corefoundation.CFDictionaryRef, allocator corefoundation.CFAllocatorRef, newCollectionOut *CMTagCollectionRef) int32
var _cMTagCollectionCreateFromDictionaryErr error

func tryCMTagCollectionCreateFromDictionary(dict corefoundation.CFDictionaryRef, allocator corefoundation.CFAllocatorRef, newCollectionOut *CMTagCollectionRef) (int32, error) {
	if _cMTagCollectionCreateFromDictionary == nil {
		return 0, symbolCallError("CMTagCollectionCreateFromDictionary", "14.0", _cMTagCollectionCreateFromDictionaryErr)
	}
	return _cMTagCollectionCreateFromDictionary(dict, allocator, newCollectionOut), nil
}

// CMTagCollectionCreateFromDictionary creates a new tag collection from an existing Core Foundation dictionary.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateFromDictionary
func CMTagCollectionCreateFromDictionary(dict corefoundation.CFDictionaryRef, allocator corefoundation.CFAllocatorRef, newCollectionOut *CMTagCollectionRef) int32 {
	result, callErr := tryCMTagCollectionCreateFromDictionary(dict, allocator, newCollectionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionCreateIntersection func(tagCollection1 CMTagCollectionRef, tagCollection2 CMTagCollectionRef, tagCollectionOut *CMTagCollectionRef) int32
var _cMTagCollectionCreateIntersectionErr error

func tryCMTagCollectionCreateIntersection(tagCollection1 CMTagCollectionRef, tagCollection2 CMTagCollectionRef, tagCollectionOut *CMTagCollectionRef) (int32, error) {
	if _cMTagCollectionCreateIntersection == nil {
		return 0, symbolCallError("CMTagCollectionCreateIntersection", "14.0", _cMTagCollectionCreateIntersectionErr)
	}
	return _cMTagCollectionCreateIntersection(tagCollection1, tagCollection2, tagCollectionOut), nil
}

// CMTagCollectionCreateIntersection creates a new tag collection containing only the tags from two existing collections which match.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateIntersection
func CMTagCollectionCreateIntersection(tagCollection1 CMTagCollectionRef, tagCollection2 CMTagCollectionRef, tagCollectionOut *CMTagCollectionRef) int32 {
	result, callErr := tryCMTagCollectionCreateIntersection(tagCollection1, tagCollection2, tagCollectionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionCreateMutable func(allocator corefoundation.CFAllocatorRef, capacity int, newMutableCollectionOut *CMMutableTagCollectionRef) int32
var _cMTagCollectionCreateMutableErr error

func tryCMTagCollectionCreateMutable(allocator corefoundation.CFAllocatorRef, capacity int, newMutableCollectionOut *CMMutableTagCollectionRef) (int32, error) {
	if _cMTagCollectionCreateMutable == nil {
		return 0, symbolCallError("CMTagCollectionCreateMutable", "14.0", _cMTagCollectionCreateMutableErr)
	}
	return _cMTagCollectionCreateMutable(allocator, capacity, newMutableCollectionOut), nil
}

// CMTagCollectionCreateMutable creates a new, mutable, empty tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateMutable
func CMTagCollectionCreateMutable(allocator corefoundation.CFAllocatorRef, capacity int, newMutableCollectionOut *CMMutableTagCollectionRef) int32 {
	result, callErr := tryCMTagCollectionCreateMutable(allocator, capacity, newMutableCollectionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionCreateMutableCopy func(tagCollection CMTagCollectionRef, allocator corefoundation.CFAllocatorRef, newMutableCollectionCopyOut *CMMutableTagCollectionRef) int32
var _cMTagCollectionCreateMutableCopyErr error

func tryCMTagCollectionCreateMutableCopy(tagCollection CMTagCollectionRef, allocator corefoundation.CFAllocatorRef, newMutableCollectionCopyOut *CMMutableTagCollectionRef) (int32, error) {
	if _cMTagCollectionCreateMutableCopy == nil {
		return 0, symbolCallError("CMTagCollectionCreateMutableCopy", "14.0", _cMTagCollectionCreateMutableCopyErr)
	}
	return _cMTagCollectionCreateMutableCopy(tagCollection, allocator, newMutableCollectionCopyOut), nil
}

// CMTagCollectionCreateMutableCopy creates a new mutable copy from an existing tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateMutableCopy
func CMTagCollectionCreateMutableCopy(tagCollection CMTagCollectionRef, allocator corefoundation.CFAllocatorRef, newMutableCollectionCopyOut *CMMutableTagCollectionRef) int32 {
	result, callErr := tryCMTagCollectionCreateMutableCopy(tagCollection, allocator, newMutableCollectionCopyOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionCreateUnion func(tagCollection1 CMTagCollectionRef, tagCollection2 CMTagCollectionRef, tagCollectionOut *CMTagCollectionRef) int32
var _cMTagCollectionCreateUnionErr error

func tryCMTagCollectionCreateUnion(tagCollection1 CMTagCollectionRef, tagCollection2 CMTagCollectionRef, tagCollectionOut *CMTagCollectionRef) (int32, error) {
	if _cMTagCollectionCreateUnion == nil {
		return 0, symbolCallError("CMTagCollectionCreateUnion", "14.0", _cMTagCollectionCreateUnionErr)
	}
	return _cMTagCollectionCreateUnion(tagCollection1, tagCollection2, tagCollectionOut), nil
}

// CMTagCollectionCreateUnion creates a new tag collection containing all tags from two collections without duplicates.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateUnion
func CMTagCollectionCreateUnion(tagCollection1 CMTagCollectionRef, tagCollection2 CMTagCollectionRef, tagCollectionOut *CMTagCollectionRef) int32 {
	result, callErr := tryCMTagCollectionCreateUnion(tagCollection1, tagCollection2, tagCollectionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionGetCount func(tagCollection CMTagCollectionRef) int
var _cMTagCollectionGetCountErr error

func tryCMTagCollectionGetCount(tagCollection CMTagCollectionRef) (int, error) {
	if _cMTagCollectionGetCount == nil {
		return 0, symbolCallError("CMTagCollectionGetCount", "14.0", _cMTagCollectionGetCountErr)
	}
	return _cMTagCollectionGetCount(tagCollection), nil
}

// CMTagCollectionGetCount gets the number of tags in a tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionGetCount
func CMTagCollectionGetCount(tagCollection CMTagCollectionRef) int {
	result, callErr := tryCMTagCollectionGetCount(tagCollection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionGetCountOfCategory func(tagCollection CMTagCollectionRef, category CMTagCategory) int
var _cMTagCollectionGetCountOfCategoryErr error

func tryCMTagCollectionGetCountOfCategory(tagCollection CMTagCollectionRef, category CMTagCategory) (int, error) {
	if _cMTagCollectionGetCountOfCategory == nil {
		return 0, symbolCallError("CMTagCollectionGetCountOfCategory", "14.0", _cMTagCollectionGetCountOfCategoryErr)
	}
	return _cMTagCollectionGetCountOfCategory(tagCollection, category), nil
}

// CMTagCollectionGetCountOfCategory retrieves the number of tags in the collection matching a given category.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionGetCountOfCategory
func CMTagCollectionGetCountOfCategory(tagCollection CMTagCollectionRef, category CMTagCategory) int {
	result, callErr := tryCMTagCollectionGetCountOfCategory(tagCollection, category)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionGetTags func(tagCollection CMTagCollectionRef, tagBuffer *CMTag, tagBufferCount int, numberOfTagsCopied *int) int32
var _cMTagCollectionGetTagsErr error

func tryCMTagCollectionGetTags(tagCollection CMTagCollectionRef, tagBuffer *CMTag, tagBufferCount int, numberOfTagsCopied *int) (int32, error) {
	if _cMTagCollectionGetTags == nil {
		return 0, symbolCallError("CMTagCollectionGetTags", "14.0", _cMTagCollectionGetTagsErr)
	}
	return _cMTagCollectionGetTags(tagCollection, tagBuffer, tagBufferCount, numberOfTagsCopied), nil
}

// CMTagCollectionGetTags retrieves an arbitrary number of tags from the collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionGetTags
func CMTagCollectionGetTags(tagCollection CMTagCollectionRef, tagBuffer *CMTag, tagBufferCount int, numberOfTagsCopied *int) int32 {
	result, callErr := tryCMTagCollectionGetTags(tagCollection, tagBuffer, tagBufferCount, numberOfTagsCopied)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionGetTagsWithCategory func(tagCollection CMTagCollectionRef, category CMTagCategory, tagBuffer *CMTag, tagBufferCount int, numberOfTagsCopied *int) int32
var _cMTagCollectionGetTagsWithCategoryErr error

func tryCMTagCollectionGetTagsWithCategory(tagCollection CMTagCollectionRef, category CMTagCategory, tagBuffer *CMTag, tagBufferCount int, numberOfTagsCopied *int) (int32, error) {
	if _cMTagCollectionGetTagsWithCategory == nil {
		return 0, symbolCallError("CMTagCollectionGetTagsWithCategory", "14.0", _cMTagCollectionGetTagsWithCategoryErr)
	}
	return _cMTagCollectionGetTagsWithCategory(tagCollection, category, tagBuffer, tagBufferCount, numberOfTagsCopied), nil
}

// CMTagCollectionGetTagsWithCategory retrieves a C-style array of tags with a given category from a tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionGetTagsWithCategory
func CMTagCollectionGetTagsWithCategory(tagCollection CMTagCollectionRef, category CMTagCategory, tagBuffer *CMTag, tagBufferCount int, numberOfTagsCopied *int) int32 {
	result, callErr := tryCMTagCollectionGetTagsWithCategory(tagCollection, category, tagBuffer, tagBufferCount, numberOfTagsCopied)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionGetTagsWithFilterFunction func(tagCollection CMTagCollectionRef, tagBuffer *CMTag, tagBufferCount int, numberOfTagsCopied *int, filter CMTagCollectionTagFilterFunction, context unsafe.Pointer) int32
var _cMTagCollectionGetTagsWithFilterFunctionErr error

func tryCMTagCollectionGetTagsWithFilterFunction(tagCollection CMTagCollectionRef, tagBuffer *CMTag, tagBufferCount int, numberOfTagsCopied *int, filter CMTagCollectionTagFilterFunction, context unsafe.Pointer) (int32, error) {
	if _cMTagCollectionGetTagsWithFilterFunction == nil {
		return 0, symbolCallError("CMTagCollectionGetTagsWithFilterFunction", "14.0", _cMTagCollectionGetTagsWithFilterFunctionErr)
	}
	return _cMTagCollectionGetTagsWithFilterFunction(tagCollection, tagBuffer, tagBufferCount, numberOfTagsCopied, filter, context), nil
}

// CMTagCollectionGetTagsWithFilterFunction gets all tags in a collection matching an evaluation function.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionGetTagsWithFilterFunction
func CMTagCollectionGetTagsWithFilterFunction(tagCollection CMTagCollectionRef, tagBuffer *CMTag, tagBufferCount int, numberOfTagsCopied *int, filter CMTagCollectionTagFilterFunction, context unsafe.Pointer) int32 {
	result, callErr := tryCMTagCollectionGetTagsWithFilterFunction(tagCollection, tagBuffer, tagBufferCount, numberOfTagsCopied, filter, context)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionGetTypeID func() uint
var _cMTagCollectionGetTypeIDErr error

func tryCMTagCollectionGetTypeID() (uint, error) {
	if _cMTagCollectionGetTypeID == nil {
		return 0, symbolCallError("CMTagCollectionGetTypeID", "14.0", _cMTagCollectionGetTypeIDErr)
	}
	return _cMTagCollectionGetTypeID(), nil
}

// CMTagCollectionGetTypeID retrieves the internal type ID for tag collections.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionGetTypeID
func CMTagCollectionGetTypeID() uint {
	result, callErr := tryCMTagCollectionGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionIsEmpty func(tagCollection CMTagCollectionRef) bool
var _cMTagCollectionIsEmptyErr error

func tryCMTagCollectionIsEmpty(tagCollection CMTagCollectionRef) (bool, error) {
	if _cMTagCollectionIsEmpty == nil {
		return false, symbolCallError("CMTagCollectionIsEmpty", "14.0", _cMTagCollectionIsEmptyErr)
	}
	return _cMTagCollectionIsEmpty(tagCollection), nil
}

// CMTagCollectionIsEmpty checks if a tag collection has no elements.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionIsEmpty
func CMTagCollectionIsEmpty(tagCollection CMTagCollectionRef) bool {
	result, callErr := tryCMTagCollectionIsEmpty(tagCollection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionRemoveAllTags func(tagCollection CMMutableTagCollectionRef) int32
var _cMTagCollectionRemoveAllTagsErr error

func tryCMTagCollectionRemoveAllTags(tagCollection CMMutableTagCollectionRef) (int32, error) {
	if _cMTagCollectionRemoveAllTags == nil {
		return 0, symbolCallError("CMTagCollectionRemoveAllTags", "14.0", _cMTagCollectionRemoveAllTagsErr)
	}
	return _cMTagCollectionRemoveAllTags(tagCollection), nil
}

// CMTagCollectionRemoveAllTags removes all tags from a collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionRemoveAllTags
func CMTagCollectionRemoveAllTags(tagCollection CMMutableTagCollectionRef) int32 {
	result, callErr := tryCMTagCollectionRemoveAllTags(tagCollection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionRemoveAllTagsOfCategory func(tagCollection CMMutableTagCollectionRef, category CMTagCategory) int32
var _cMTagCollectionRemoveAllTagsOfCategoryErr error

func tryCMTagCollectionRemoveAllTagsOfCategory(tagCollection CMMutableTagCollectionRef, category CMTagCategory) (int32, error) {
	if _cMTagCollectionRemoveAllTagsOfCategory == nil {
		return 0, symbolCallError("CMTagCollectionRemoveAllTagsOfCategory", "14.0", _cMTagCollectionRemoveAllTagsOfCategoryErr)
	}
	return _cMTagCollectionRemoveAllTagsOfCategory(tagCollection, category), nil
}

// CMTagCollectionRemoveAllTagsOfCategory removes all tags of a given category from a collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionRemoveAllTagsOfCategory
func CMTagCollectionRemoveAllTagsOfCategory(tagCollection CMMutableTagCollectionRef, category CMTagCategory) int32 {
	result, callErr := tryCMTagCollectionRemoveAllTagsOfCategory(tagCollection, category)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionRemoveTag func(tagCollection CMMutableTagCollectionRef, tagToRemove CMTag) int32
var _cMTagCollectionRemoveTagErr error

func tryCMTagCollectionRemoveTag(tagCollection CMMutableTagCollectionRef, tagToRemove CMTag) (int32, error) {
	if _cMTagCollectionRemoveTag == nil {
		return 0, symbolCallError("CMTagCollectionRemoveTag", "14.0", _cMTagCollectionRemoveTagErr)
	}
	return _cMTagCollectionRemoveTag(tagCollection, tagToRemove), nil
}

// CMTagCollectionRemoveTag removes a specific tag from a collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionRemoveTag
func CMTagCollectionRemoveTag(tagCollection CMMutableTagCollectionRef, tagToRemove CMTag) int32 {
	result, callErr := tryCMTagCollectionRemoveTag(tagCollection, tagToRemove)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCompare func(tag1 CMTag, tag2 CMTag) corefoundation.CFComparisonResult
var _cMTagCompareErr error

func tryCMTagCompare(tag1 CMTag, tag2 CMTag) (corefoundation.CFComparisonResult, error) {
	if _cMTagCompare == nil {
		return *new(corefoundation.CFComparisonResult), symbolCallError("CMTagCompare", "14.0", _cMTagCompareErr)
	}
	return _cMTagCompare(tag1, tag2), nil
}

// CMTagCompare compares two tags in terms of partial equality.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCompare
func CMTagCompare(tag1 CMTag, tag2 CMTag) corefoundation.CFComparisonResult {
	result, callErr := tryCMTagCompare(tag1, tag2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCopyAsDictionary func(tag CMTag, allocator corefoundation.CFAllocatorRef) corefoundation.CFDictionaryRef
var _cMTagCopyAsDictionaryErr error

func tryCMTagCopyAsDictionary(tag CMTag, allocator corefoundation.CFAllocatorRef) (corefoundation.CFDictionaryRef, error) {
	if _cMTagCopyAsDictionary == nil {
		return 0, symbolCallError("CMTagCopyAsDictionary", "14.0", _cMTagCopyAsDictionaryErr)
	}
	return _cMTagCopyAsDictionary(tag, allocator), nil
}

// CMTagCopyAsDictionary copies an existing tag to a new dictionary object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCopyAsDictionary
func CMTagCopyAsDictionary(tag CMTag, allocator corefoundation.CFAllocatorRef) corefoundation.CFDictionaryRef {
	result, callErr := tryCMTagCopyAsDictionary(tag, allocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCopyDescription func(allocator corefoundation.CFAllocatorRef, tag CMTag) corefoundation.CFStringRef
var _cMTagCopyDescriptionErr error

func tryCMTagCopyDescription(allocator corefoundation.CFAllocatorRef, tag CMTag) (corefoundation.CFStringRef, error) {
	if _cMTagCopyDescription == nil {
		return 0, symbolCallError("CMTagCopyDescription", "14.0", _cMTagCopyDescriptionErr)
	}
	return _cMTagCopyDescription(allocator, tag), nil
}

// CMTagCopyDescription copies the description of a tag to a new string.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCopyDescription
func CMTagCopyDescription(allocator corefoundation.CFAllocatorRef, tag CMTag) corefoundation.CFStringRef {
	result, callErr := tryCMTagCopyDescription(allocator, tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagEqualToTag func(tag1 CMTag, tag2 CMTag) bool
var _cMTagEqualToTagErr error

func tryCMTagEqualToTag(tag1 CMTag, tag2 CMTag) (bool, error) {
	if _cMTagEqualToTag == nil {
		return false, symbolCallError("CMTagEqualToTag", "14.0", _cMTagEqualToTagErr)
	}
	return _cMTagEqualToTag(tag1, tag2), nil
}

// CMTagEqualToTag compares two tags for strict equality.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagEqualToTag
func CMTagEqualToTag(tag1 CMTag, tag2 CMTag) bool {
	result, callErr := tryCMTagEqualToTag(tag1, tag2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagGetFlagsValue func(tag CMTag) uint64
var _cMTagGetFlagsValueErr error

func tryCMTagGetFlagsValue(tag CMTag) (uint64, error) {
	if _cMTagGetFlagsValue == nil {
		return 0, symbolCallError("CMTagGetFlagsValue", "14.0", _cMTagGetFlagsValueErr)
	}
	return _cMTagGetFlagsValue(tag), nil
}

// CMTagGetFlagsValue retrieves a tag’s value as a 64-bit field flag.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagGetFlagsValue
func CMTagGetFlagsValue(tag CMTag) uint64 {
	result, callErr := tryCMTagGetFlagsValue(tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagGetFloat64Value func(tag CMTag) float64
var _cMTagGetFloat64ValueErr error

func tryCMTagGetFloat64Value(tag CMTag) (float64, error) {
	if _cMTagGetFloat64Value == nil {
		return 0.0, symbolCallError("CMTagGetFloat64Value", "14.0", _cMTagGetFloat64ValueErr)
	}
	return _cMTagGetFloat64Value(tag), nil
}

// CMTagGetFloat64Value retrieves a tag’s value as a 64-bit floating point number.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagGetFloat64Value
func CMTagGetFloat64Value(tag CMTag) float64 {
	result, callErr := tryCMTagGetFloat64Value(tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagGetOSTypeValue func(tag CMTag) uint32
var _cMTagGetOSTypeValueErr error

func tryCMTagGetOSTypeValue(tag CMTag) (uint32, error) {
	if _cMTagGetOSTypeValue == nil {
		return 0, symbolCallError("CMTagGetOSTypeValue", "14.0", _cMTagGetOSTypeValueErr)
	}
	return _cMTagGetOSTypeValue(tag), nil
}

// CMTagGetOSTypeValue retrieves a tag’s value for use by the operating system.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagGetOSTypeValue
func CMTagGetOSTypeValue(tag CMTag) uint32 {
	result, callErr := tryCMTagGetOSTypeValue(tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagGetSInt64Value func(tag CMTag) int64
var _cMTagGetSInt64ValueErr error

func tryCMTagGetSInt64Value(tag CMTag) (int64, error) {
	if _cMTagGetSInt64Value == nil {
		return 0, symbolCallError("CMTagGetSInt64Value", "14.0", _cMTagGetSInt64ValueErr)
	}
	return _cMTagGetSInt64Value(tag), nil
}

// CMTagGetSInt64Value retrieves a tag’s value as a signed 64-bit integer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagGetSInt64Value
func CMTagGetSInt64Value(tag CMTag) int64 {
	result, callErr := tryCMTagGetSInt64Value(tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagGetValueDataType func(tag CMTag) CMTagDataType
var _cMTagGetValueDataTypeErr error

func tryCMTagGetValueDataType(tag CMTag) (CMTagDataType, error) {
	if _cMTagGetValueDataType == nil {
		return *new(CMTagDataType), symbolCallError("CMTagGetValueDataType", "14.0", _cMTagGetValueDataTypeErr)
	}
	return _cMTagGetValueDataType(tag), nil
}

// CMTagGetValueDataType retrieves the data type of a tag.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagGetValueDataType
func CMTagGetValueDataType(tag CMTag) CMTagDataType {
	result, callErr := tryCMTagGetValueDataType(tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagHasFlagsValue func(tag CMTag) bool
var _cMTagHasFlagsValueErr error

func tryCMTagHasFlagsValue(tag CMTag) (bool, error) {
	if _cMTagHasFlagsValue == nil {
		return false, symbolCallError("CMTagHasFlagsValue", "14.0", _cMTagHasFlagsValueErr)
	}
	return _cMTagHasFlagsValue(tag), nil
}

// CMTagHasFlagsValue whether a given tag contains a value for a 64-bit flag field.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagHasFlagsValue
func CMTagHasFlagsValue(tag CMTag) bool {
	result, callErr := tryCMTagHasFlagsValue(tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagHasFloat64Value func(tag CMTag) bool
var _cMTagHasFloat64ValueErr error

func tryCMTagHasFloat64Value(tag CMTag) (bool, error) {
	if _cMTagHasFloat64Value == nil {
		return false, symbolCallError("CMTagHasFloat64Value", "14.0", _cMTagHasFloat64ValueErr)
	}
	return _cMTagHasFloat64Value(tag), nil
}

// CMTagHasFloat64Value whether a given tag contains a value for a 64-bit floating point number.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagHasFloat64Value
func CMTagHasFloat64Value(tag CMTag) bool {
	result, callErr := tryCMTagHasFloat64Value(tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagHasOSTypeValue func(tag CMTag) bool
var _cMTagHasOSTypeValueErr error

func tryCMTagHasOSTypeValue(tag CMTag) (bool, error) {
	if _cMTagHasOSTypeValue == nil {
		return false, symbolCallError("CMTagHasOSTypeValue", "14.0", _cMTagHasOSTypeValueErr)
	}
	return _cMTagHasOSTypeValue(tag), nil
}

// CMTagHasOSTypeValue whether a given tag contains a value for use by the operating system.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagHasOSTypeValue
func CMTagHasOSTypeValue(tag CMTag) bool {
	result, callErr := tryCMTagHasOSTypeValue(tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagHasSInt64Value func(tag CMTag) bool
var _cMTagHasSInt64ValueErr error

func tryCMTagHasSInt64Value(tag CMTag) (bool, error) {
	if _cMTagHasSInt64Value == nil {
		return false, symbolCallError("CMTagHasSInt64Value", "14.0", _cMTagHasSInt64ValueErr)
	}
	return _cMTagHasSInt64Value(tag), nil
}

// CMTagHasSInt64Value whether a given tag contains a value for a signed 64-bit integer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagHasSInt64Value
func CMTagHasSInt64Value(tag CMTag) bool {
	result, callErr := tryCMTagHasSInt64Value(tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagHash func(tag CMTag) uint
var _cMTagHashErr error

func tryCMTagHash(tag CMTag) (uint, error) {
	if _cMTagHash == nil {
		return 0, symbolCallError("CMTagHash", "14.0", _cMTagHashErr)
	}
	return _cMTagHash(tag), nil
}

// CMTagHash generates a hash identifier for a tag.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagHash
func CMTagHash(tag CMTag) uint {
	result, callErr := tryCMTagHash(tag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagMakeFromDictionary func(dict corefoundation.CFDictionaryRef) CMTag
var _cMTagMakeFromDictionaryErr error

func tryCMTagMakeFromDictionary(dict corefoundation.CFDictionaryRef) (CMTag, error) {
	if _cMTagMakeFromDictionary == nil {
		return CMTag{}, symbolCallError("CMTagMakeFromDictionary", "14.0", _cMTagMakeFromDictionaryErr)
	}
	return _cMTagMakeFromDictionary(dict), nil
}

// CMTagMakeFromDictionary create a new tag from a dictionary object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagMakeFromDictionary
func CMTagMakeFromDictionary(dict corefoundation.CFDictionaryRef) CMTag {
	result, callErr := tryCMTagMakeFromDictionary(dict)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagMakeWithFlagsValue func(category CMTagCategory, flagsForTag uint64) CMTag
var _cMTagMakeWithFlagsValueErr error

func tryCMTagMakeWithFlagsValue(category CMTagCategory, flagsForTag uint64) (CMTag, error) {
	if _cMTagMakeWithFlagsValue == nil {
		return CMTag{}, symbolCallError("CMTagMakeWithFlagsValue", "14.0", _cMTagMakeWithFlagsValueErr)
	}
	return _cMTagMakeWithFlagsValue(category, flagsForTag), nil
}

// CMTagMakeWithFlagsValue creates a new tag with a given category and a value interpreted as a 64-bit flag field.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagMakeWithFlagsValue
func CMTagMakeWithFlagsValue(category CMTagCategory, flagsForTag uint64) CMTag {
	result, callErr := tryCMTagMakeWithFlagsValue(category, flagsForTag)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagMakeWithFloat64Value func(category CMTagCategory, value float64) CMTag
var _cMTagMakeWithFloat64ValueErr error

func tryCMTagMakeWithFloat64Value(category CMTagCategory, value float64) (CMTag, error) {
	if _cMTagMakeWithFloat64Value == nil {
		return CMTag{}, symbolCallError("CMTagMakeWithFloat64Value", "14.0", _cMTagMakeWithFloat64ValueErr)
	}
	return _cMTagMakeWithFloat64Value(category, value), nil
}

// CMTagMakeWithFloat64Value creates a new tag with a given category and a 64-bit floating point value.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagMakeWithFloat64Value
func CMTagMakeWithFloat64Value(category CMTagCategory, value float64) CMTag {
	result, callErr := tryCMTagMakeWithFloat64Value(category, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagMakeWithOSTypeValue func(category CMTagCategory, value uint32) CMTag
var _cMTagMakeWithOSTypeValueErr error

func tryCMTagMakeWithOSTypeValue(category CMTagCategory, value uint32) (CMTag, error) {
	if _cMTagMakeWithOSTypeValue == nil {
		return CMTag{}, symbolCallError("CMTagMakeWithOSTypeValue", "14.0", _cMTagMakeWithOSTypeValueErr)
	}
	return _cMTagMakeWithOSTypeValue(category, value), nil
}

// CMTagMakeWithOSTypeValue creates a new tag with a given category and a 64-bit value for use by the framework.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagMakeWithOSTypeValue
func CMTagMakeWithOSTypeValue(category CMTagCategory, value uint32) CMTag {
	result, callErr := tryCMTagMakeWithOSTypeValue(category, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagMakeWithSInt64Value func(category CMTagCategory, value int64) CMTag
var _cMTagMakeWithSInt64ValueErr error

func tryCMTagMakeWithSInt64Value(category CMTagCategory, value int64) (CMTag, error) {
	if _cMTagMakeWithSInt64Value == nil {
		return CMTag{}, symbolCallError("CMTagMakeWithSInt64Value", "14.0", _cMTagMakeWithSInt64ValueErr)
	}
	return _cMTagMakeWithSInt64Value(category, value), nil
}

// CMTagMakeWithSInt64Value creates a new tag with a given category and a 64-bit signed integer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagMakeWithSInt64Value
func CMTagMakeWithSInt64Value(category CMTagCategory, value int64) CMTag {
	result, callErr := tryCMTagMakeWithSInt64Value(category, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTaggedBufferGroupCreate func(allocator corefoundation.CFAllocatorRef, tagCollections corefoundation.CFArrayRef, buffers corefoundation.CFArrayRef, groupOut *CMTaggedBufferGroupRef) int32
var _cMTaggedBufferGroupCreateErr error

func tryCMTaggedBufferGroupCreate(allocator corefoundation.CFAllocatorRef, tagCollections corefoundation.CFArrayRef, buffers corefoundation.CFArrayRef, groupOut *CMTaggedBufferGroupRef) (int32, error) {
	if _cMTaggedBufferGroupCreate == nil {
		return 0, symbolCallError("CMTaggedBufferGroupCreate", "14.0", _cMTaggedBufferGroupCreateErr)
	}
	return _cMTaggedBufferGroupCreate(allocator, tagCollections, buffers, groupOut), nil
}

// CMTaggedBufferGroupCreate creates a new tagged buffer group from a pair of buffers and the tags to associate with them.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupCreate
func CMTaggedBufferGroupCreate(allocator corefoundation.CFAllocatorRef, tagCollections corefoundation.CFArrayRef, buffers corefoundation.CFArrayRef, groupOut *CMTaggedBufferGroupRef) int32 {
	result, callErr := tryCMTaggedBufferGroupCreate(allocator, tagCollections, buffers, groupOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTaggedBufferGroupCreateCombined func(allocator corefoundation.CFAllocatorRef, taggedBufferGroups corefoundation.CFArrayRef, groupOut *CMTaggedBufferGroupRef) int32
var _cMTaggedBufferGroupCreateCombinedErr error

func tryCMTaggedBufferGroupCreateCombined(allocator corefoundation.CFAllocatorRef, taggedBufferGroups corefoundation.CFArrayRef, groupOut *CMTaggedBufferGroupRef) (int32, error) {
	if _cMTaggedBufferGroupCreateCombined == nil {
		return 0, symbolCallError("CMTaggedBufferGroupCreateCombined", "14.0", _cMTaggedBufferGroupCreateCombinedErr)
	}
	return _cMTaggedBufferGroupCreateCombined(allocator, taggedBufferGroups, groupOut), nil
}

// CMTaggedBufferGroupCreateCombined creates a new tagged buffer group from an array of existing tagged buffer groups.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupCreateCombined
func CMTaggedBufferGroupCreateCombined(allocator corefoundation.CFAllocatorRef, taggedBufferGroups corefoundation.CFArrayRef, groupOut *CMTaggedBufferGroupRef) int32 {
	result, callErr := tryCMTaggedBufferGroupCreateCombined(allocator, taggedBufferGroups, groupOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup func(allocator corefoundation.CFAllocatorRef, taggedBufferGroup CMTaggedBufferGroupRef, formatDescriptionOut *CMTaggedBufferGroupFormatDescriptionRef) int32
var _cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupErr error

func tryCMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup(allocator corefoundation.CFAllocatorRef, taggedBufferGroup CMTaggedBufferGroupRef, formatDescriptionOut *CMTaggedBufferGroupFormatDescriptionRef) (int32, error) {
	if _cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup == nil {
		return 0, symbolCallError("CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup", "14.0", _cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupErr)
	}
	return _cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup(allocator, taggedBufferGroup, formatDescriptionOut), nil
}

// CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup creates a new format description for a tagged buffer group.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup
func CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup(allocator corefoundation.CFAllocatorRef, taggedBufferGroup CMTaggedBufferGroupRef, formatDescriptionOut *CMTaggedBufferGroupFormatDescriptionRef) int32 {
	result, callErr := tryCMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup(allocator, taggedBufferGroup, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions func(allocator corefoundation.CFAllocatorRef, taggedBufferGroup CMTaggedBufferGroupRef, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *CMTaggedBufferGroupFormatDescriptionRef) int32
var _cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensionsErr error

func tryCMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions(allocator corefoundation.CFAllocatorRef, taggedBufferGroup CMTaggedBufferGroupRef, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *CMTaggedBufferGroupFormatDescriptionRef) (int32, error) {
	if _cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions == nil {
		return 0, symbolCallError("CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions", "26.0", _cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensionsErr)
	}
	return _cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions(allocator, taggedBufferGroup, extensions, formatDescriptionOut), nil
}

// CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions
func CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions(allocator corefoundation.CFAllocatorRef, taggedBufferGroup CMTaggedBufferGroupRef, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *CMTaggedBufferGroupFormatDescriptionRef) int32 {
	result, callErr := tryCMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions(allocator, taggedBufferGroup, extensions, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup func(desc CMTaggedBufferGroupFormatDescriptionRef, taggedBufferGroup CMTaggedBufferGroupRef) bool
var _cMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroupErr error

func tryCMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup(desc CMTaggedBufferGroupFormatDescriptionRef, taggedBufferGroup CMTaggedBufferGroupRef) (bool, error) {
	if _cMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup == nil {
		return false, symbolCallError("CMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup", "14.0", _cMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroupErr)
	}
	return _cMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup(desc, taggedBufferGroup), nil
}

// CMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup checks to see if a tagged buffer group’s format matches an existing format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup
func CMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup(desc CMTaggedBufferGroupFormatDescriptionRef, taggedBufferGroup CMTaggedBufferGroupRef) bool {
	result, callErr := tryCMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup(desc, taggedBufferGroup)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTaggedBufferGroupGetCMSampleBufferAtIndex func(group CMTaggedBufferGroupRef, index int) uintptr
var _cMTaggedBufferGroupGetCMSampleBufferAtIndexErr error

func tryCMTaggedBufferGroupGetCMSampleBufferAtIndex(group CMTaggedBufferGroupRef, index int) (uintptr, error) {
	if _cMTaggedBufferGroupGetCMSampleBufferAtIndex == nil {
		return 0, symbolCallError("CMTaggedBufferGroupGetCMSampleBufferAtIndex", "14.0", _cMTaggedBufferGroupGetCMSampleBufferAtIndexErr)
	}
	return _cMTaggedBufferGroupGetCMSampleBufferAtIndex(group, index), nil
}

// CMTaggedBufferGroupGetCMSampleBufferAtIndex gets the sample buffer at a given index in the buffer group.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCMSampleBufferAtIndex
func CMTaggedBufferGroupGetCMSampleBufferAtIndex(group CMTaggedBufferGroupRef, index int) uintptr {
	result, callErr := tryCMTaggedBufferGroupGetCMSampleBufferAtIndex(group, index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTaggedBufferGroupGetCMSampleBufferForTag func(group CMTaggedBufferGroupRef, tag CMTag, indexOut *int) uintptr
var _cMTaggedBufferGroupGetCMSampleBufferForTagErr error

func tryCMTaggedBufferGroupGetCMSampleBufferForTag(group CMTaggedBufferGroupRef, tag CMTag, indexOut *int) (uintptr, error) {
	if _cMTaggedBufferGroupGetCMSampleBufferForTag == nil {
		return 0, symbolCallError("CMTaggedBufferGroupGetCMSampleBufferForTag", "14.0", _cMTaggedBufferGroupGetCMSampleBufferForTagErr)
	}
	return _cMTaggedBufferGroupGetCMSampleBufferForTag(group, tag, indexOut), nil
}

// CMTaggedBufferGroupGetCMSampleBufferForTag gets the single sample buffer in a group which contains a given tag, if present.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCMSampleBufferForTag
func CMTaggedBufferGroupGetCMSampleBufferForTag(group CMTaggedBufferGroupRef, tag CMTag, indexOut *int) uintptr {
	result, callErr := tryCMTaggedBufferGroupGetCMSampleBufferForTag(group, tag, indexOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTaggedBufferGroupGetCMSampleBufferForTagCollection func(group CMTaggedBufferGroupRef, tagCollection CMTagCollectionRef, indexOut *int) uintptr
var _cMTaggedBufferGroupGetCMSampleBufferForTagCollectionErr error

func tryCMTaggedBufferGroupGetCMSampleBufferForTagCollection(group CMTaggedBufferGroupRef, tagCollection CMTagCollectionRef, indexOut *int) (uintptr, error) {
	if _cMTaggedBufferGroupGetCMSampleBufferForTagCollection == nil {
		return 0, symbolCallError("CMTaggedBufferGroupGetCMSampleBufferForTagCollection", "14.0", _cMTaggedBufferGroupGetCMSampleBufferForTagCollectionErr)
	}
	return _cMTaggedBufferGroupGetCMSampleBufferForTagCollection(group, tagCollection, indexOut), nil
}

// CMTaggedBufferGroupGetCMSampleBufferForTagCollection gets the single sample buffer in a group which contains a given tag collection, if present.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCMSampleBufferForTagCollection
func CMTaggedBufferGroupGetCMSampleBufferForTagCollection(group CMTaggedBufferGroupRef, tagCollection CMTagCollectionRef, indexOut *int) uintptr {
	result, callErr := tryCMTaggedBufferGroupGetCMSampleBufferForTagCollection(group, tagCollection, indexOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTaggedBufferGroupGetCVPixelBufferAtIndex func(group CMTaggedBufferGroupRef, index int) corevideo.CVPixelBufferRef
var _cMTaggedBufferGroupGetCVPixelBufferAtIndexErr error

func tryCMTaggedBufferGroupGetCVPixelBufferAtIndex(group CMTaggedBufferGroupRef, index int) (corevideo.CVPixelBufferRef, error) {
	if _cMTaggedBufferGroupGetCVPixelBufferAtIndex == nil {
		return 0, symbolCallError("CMTaggedBufferGroupGetCVPixelBufferAtIndex", "14.0", _cMTaggedBufferGroupGetCVPixelBufferAtIndexErr)
	}
	return _cMTaggedBufferGroupGetCVPixelBufferAtIndex(group, index), nil
}

// CMTaggedBufferGroupGetCVPixelBufferAtIndex gets the pixel buffer at a given index in the buffer group.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCVPixelBufferAtIndex
func CMTaggedBufferGroupGetCVPixelBufferAtIndex(group CMTaggedBufferGroupRef, index int) corevideo.CVPixelBufferRef {
	result, callErr := tryCMTaggedBufferGroupGetCVPixelBufferAtIndex(group, index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTaggedBufferGroupGetCVPixelBufferForTag func(group CMTaggedBufferGroupRef, tag CMTag, indexOut *int) corevideo.CVPixelBufferRef
var _cMTaggedBufferGroupGetCVPixelBufferForTagErr error

func tryCMTaggedBufferGroupGetCVPixelBufferForTag(group CMTaggedBufferGroupRef, tag CMTag, indexOut *int) (corevideo.CVPixelBufferRef, error) {
	if _cMTaggedBufferGroupGetCVPixelBufferForTag == nil {
		return 0, symbolCallError("CMTaggedBufferGroupGetCVPixelBufferForTag", "14.0", _cMTaggedBufferGroupGetCVPixelBufferForTagErr)
	}
	return _cMTaggedBufferGroupGetCVPixelBufferForTag(group, tag, indexOut), nil
}

// CMTaggedBufferGroupGetCVPixelBufferForTag gets the single pixel buffer in a group which contains a given tag, if present.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCVPixelBufferForTag
func CMTaggedBufferGroupGetCVPixelBufferForTag(group CMTaggedBufferGroupRef, tag CMTag, indexOut *int) corevideo.CVPixelBufferRef {
	result, callErr := tryCMTaggedBufferGroupGetCVPixelBufferForTag(group, tag, indexOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTaggedBufferGroupGetCVPixelBufferForTagCollection func(group CMTaggedBufferGroupRef, tagCollection CMTagCollectionRef, indexOut *int) corevideo.CVPixelBufferRef
var _cMTaggedBufferGroupGetCVPixelBufferForTagCollectionErr error

func tryCMTaggedBufferGroupGetCVPixelBufferForTagCollection(group CMTaggedBufferGroupRef, tagCollection CMTagCollectionRef, indexOut *int) (corevideo.CVPixelBufferRef, error) {
	if _cMTaggedBufferGroupGetCVPixelBufferForTagCollection == nil {
		return 0, symbolCallError("CMTaggedBufferGroupGetCVPixelBufferForTagCollection", "14.0", _cMTaggedBufferGroupGetCVPixelBufferForTagCollectionErr)
	}
	return _cMTaggedBufferGroupGetCVPixelBufferForTagCollection(group, tagCollection, indexOut), nil
}

// CMTaggedBufferGroupGetCVPixelBufferForTagCollection gets the single pixel buffer in a group which contains a given tag collection, if present.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCVPixelBufferForTagCollection
func CMTaggedBufferGroupGetCVPixelBufferForTagCollection(group CMTaggedBufferGroupRef, tagCollection CMTagCollectionRef, indexOut *int) corevideo.CVPixelBufferRef {
	result, callErr := tryCMTaggedBufferGroupGetCVPixelBufferForTagCollection(group, tagCollection, indexOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTaggedBufferGroupGetCount func(group CMTaggedBufferGroupRef) int
var _cMTaggedBufferGroupGetCountErr error

func tryCMTaggedBufferGroupGetCount(group CMTaggedBufferGroupRef) (int, error) {
	if _cMTaggedBufferGroupGetCount == nil {
		return 0, symbolCallError("CMTaggedBufferGroupGetCount", "14.0", _cMTaggedBufferGroupGetCountErr)
	}
	return _cMTaggedBufferGroupGetCount(group), nil
}

// CMTaggedBufferGroupGetCount gets the number of buffers contained within a tagged buffer group.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCount
func CMTaggedBufferGroupGetCount(group CMTaggedBufferGroupRef) int {
	result, callErr := tryCMTaggedBufferGroupGetCount(group)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTaggedBufferGroupGetNumberOfMatchesForTagCollection func(group CMTaggedBufferGroupRef, tagCollection CMTagCollectionRef) int
var _cMTaggedBufferGroupGetNumberOfMatchesForTagCollectionErr error

func tryCMTaggedBufferGroupGetNumberOfMatchesForTagCollection(group CMTaggedBufferGroupRef, tagCollection CMTagCollectionRef) (int, error) {
	if _cMTaggedBufferGroupGetNumberOfMatchesForTagCollection == nil {
		return 0, symbolCallError("CMTaggedBufferGroupGetNumberOfMatchesForTagCollection", "14.0", _cMTaggedBufferGroupGetNumberOfMatchesForTagCollectionErr)
	}
	return _cMTaggedBufferGroupGetNumberOfMatchesForTagCollection(group, tagCollection), nil
}

// CMTaggedBufferGroupGetNumberOfMatchesForTagCollection gets the number of buffers in the group associated with a given tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetNumberOfMatchesForTagCollection
func CMTaggedBufferGroupGetNumberOfMatchesForTagCollection(group CMTaggedBufferGroupRef, tagCollection CMTagCollectionRef) int {
	result, callErr := tryCMTaggedBufferGroupGetNumberOfMatchesForTagCollection(group, tagCollection)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTaggedBufferGroupGetTagCollectionAtIndex func(group CMTaggedBufferGroupRef, index int) CMTagCollectionRef
var _cMTaggedBufferGroupGetTagCollectionAtIndexErr error

func tryCMTaggedBufferGroupGetTagCollectionAtIndex(group CMTaggedBufferGroupRef, index int) (CMTagCollectionRef, error) {
	if _cMTaggedBufferGroupGetTagCollectionAtIndex == nil {
		return 0, symbolCallError("CMTaggedBufferGroupGetTagCollectionAtIndex", "14.0", _cMTaggedBufferGroupGetTagCollectionAtIndexErr)
	}
	return _cMTaggedBufferGroupGetTagCollectionAtIndex(group, index), nil
}

// CMTaggedBufferGroupGetTagCollectionAtIndex gets the collection of tags for a buffer at a given index in the group.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetTagCollectionAtIndex
func CMTaggedBufferGroupGetTagCollectionAtIndex(group CMTaggedBufferGroupRef, index int) CMTagCollectionRef {
	result, callErr := tryCMTaggedBufferGroupGetTagCollectionAtIndex(group, index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTaggedBufferGroupGetTypeID func() uint
var _cMTaggedBufferGroupGetTypeIDErr error

func tryCMTaggedBufferGroupGetTypeID() (uint, error) {
	if _cMTaggedBufferGroupGetTypeID == nil {
		return 0, symbolCallError("CMTaggedBufferGroupGetTypeID", "14.0", _cMTaggedBufferGroupGetTypeIDErr)
	}
	return _cMTaggedBufferGroupGetTypeID(), nil
}

// CMTaggedBufferGroupGetTypeID gets the internal type ID for a tagged buffer group.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetTypeID
func CMTaggedBufferGroupGetTypeID() uint {
	result, callErr := tryCMTaggedBufferGroupGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, textFormatDescription CMTextFormatDescriptionRef, flavor CMTextDescriptionFlavor, blockBufferOut *uintptr) int32
var _cMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBufferErr error

func tryCMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, textFormatDescription CMTextFormatDescriptionRef, flavor CMTextDescriptionFlavor, blockBufferOut *uintptr) (int32, error) {
	if _cMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer == nil {
		return 0, symbolCallError("CMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer", "10.10", _cMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBufferErr)
	}
	return _cMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer(allocator, textFormatDescription, flavor, blockBufferOut), nil
}

// CMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer copies the contents of a text format description to a buffer in big-endian byte order.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer(allocator:textFormatDescription:flavor:blockBufferOut:)
func CMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, textFormatDescription CMTextFormatDescriptionRef, flavor CMTextDescriptionFlavor, blockBufferOut *uintptr) int32 {
	result, callErr := tryCMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer(allocator, textFormatDescription, flavor, blockBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, textDescriptionBlockBuffer uintptr, flavor CMTextDescriptionFlavor, mediaType uint32, formatDescriptionOut *CMTextFormatDescriptionRef) int32
var _cMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBufferErr error

func tryCMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, textDescriptionBlockBuffer uintptr, flavor CMTextDescriptionFlavor, mediaType uint32, formatDescriptionOut *CMTextFormatDescriptionRef) (int32, error) {
	if _cMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer == nil {
		return 0, symbolCallError("CMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer", "10.10", _cMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBufferErr)
	}
	return _cMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer(allocator, textDescriptionBlockBuffer, flavor, mediaType, formatDescriptionOut), nil
}

// CMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer creates a text format description from a big-endian text description structure inside a buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer(allocator:bigEndianTextDescriptionBlockBuffer:flavor:mediaType:formatDescriptionOut:)
func CMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, textDescriptionBlockBuffer uintptr, flavor CMTextDescriptionFlavor, mediaType uint32, formatDescriptionOut *CMTextFormatDescriptionRef) int32 {
	result, callErr := tryCMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer(allocator, textDescriptionBlockBuffer, flavor, mediaType, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTextFormatDescriptionCreateFromBigEndianTextDescriptionData func(allocator corefoundation.CFAllocatorRef, textDescriptionData *uint8, size uintptr, flavor CMTextDescriptionFlavor, mediaType uint32, formatDescriptionOut *CMTextFormatDescriptionRef) int32
var _cMTextFormatDescriptionCreateFromBigEndianTextDescriptionDataErr error

func tryCMTextFormatDescriptionCreateFromBigEndianTextDescriptionData(allocator corefoundation.CFAllocatorRef, textDescriptionData *uint8, size uintptr, flavor CMTextDescriptionFlavor, mediaType uint32, formatDescriptionOut *CMTextFormatDescriptionRef) (int32, error) {
	if _cMTextFormatDescriptionCreateFromBigEndianTextDescriptionData == nil {
		return 0, symbolCallError("CMTextFormatDescriptionCreateFromBigEndianTextDescriptionData", "10.10", _cMTextFormatDescriptionCreateFromBigEndianTextDescriptionDataErr)
	}
	return _cMTextFormatDescriptionCreateFromBigEndianTextDescriptionData(allocator, textDescriptionData, size, flavor, mediaType, formatDescriptionOut), nil
}

// CMTextFormatDescriptionCreateFromBigEndianTextDescriptionData creates a text format description from a big-endian text description structure.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionCreateFromBigEndianTextDescriptionData(allocator:bigEndianTextDescriptionData:size:flavor:mediaType:formatDescriptionOut:)
func CMTextFormatDescriptionCreateFromBigEndianTextDescriptionData(allocator corefoundation.CFAllocatorRef, textDescriptionData *uint8, size uintptr, flavor CMTextDescriptionFlavor, mediaType uint32, formatDescriptionOut *CMTextFormatDescriptionRef) int32 {
	result, callErr := tryCMTextFormatDescriptionCreateFromBigEndianTextDescriptionData(allocator, textDescriptionData, size, flavor, mediaType, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTextFormatDescriptionGetDefaultStyle func(desc uintptr, localFontIDOut *uint16, boldOut *bool, italicOut *bool, underlineOut *bool, fontSizeOut *float64, colorComponentsOut float64) int32
var _cMTextFormatDescriptionGetDefaultStyleErr error

func tryCMTextFormatDescriptionGetDefaultStyle(desc uintptr, localFontIDOut *uint16, boldOut *bool, italicOut *bool, underlineOut *bool, fontSizeOut *float64, colorComponentsOut float64) (int32, error) {
	if _cMTextFormatDescriptionGetDefaultStyle == nil {
		return 0, symbolCallError("CMTextFormatDescriptionGetDefaultStyle", "10.7", _cMTextFormatDescriptionGetDefaultStyleErr)
	}
	return _cMTextFormatDescriptionGetDefaultStyle(desc, localFontIDOut, boldOut, italicOut, underlineOut, fontSizeOut, colorComponentsOut), nil
}

// CMTextFormatDescriptionGetDefaultStyle returns the default text style.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionGetDefaultStyle(_:localFontIDOut:boldOut:italicOut:underlineOut:fontSizeOut:colorComponentsOut:)
func CMTextFormatDescriptionGetDefaultStyle(desc uintptr, localFontIDOut *uint16, boldOut *bool, italicOut *bool, underlineOut *bool, fontSizeOut *float64, colorComponentsOut float64) int32 {
	result, callErr := tryCMTextFormatDescriptionGetDefaultStyle(desc, localFontIDOut, boldOut, italicOut, underlineOut, fontSizeOut, colorComponentsOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTextFormatDescriptionGetDefaultTextBox func(desc uintptr, originIsAtTopLeft bool, heightOfTextTrack float64, defaultTextBoxOut *corefoundation.CGRect) int32
var _cMTextFormatDescriptionGetDefaultTextBoxErr error

func tryCMTextFormatDescriptionGetDefaultTextBox(desc uintptr, originIsAtTopLeft bool, heightOfTextTrack float64, defaultTextBoxOut *corefoundation.CGRect) (int32, error) {
	if _cMTextFormatDescriptionGetDefaultTextBox == nil {
		return 0, symbolCallError("CMTextFormatDescriptionGetDefaultTextBox", "10.7", _cMTextFormatDescriptionGetDefaultTextBoxErr)
	}
	return _cMTextFormatDescriptionGetDefaultTextBox(desc, originIsAtTopLeft, heightOfTextTrack, defaultTextBoxOut), nil
}

// CMTextFormatDescriptionGetDefaultTextBox returns the default text box.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionGetDefaultTextBox(_:originIsAtTopLeft:heightOfTextTrack:defaultTextBoxOut:)
func CMTextFormatDescriptionGetDefaultTextBox(desc uintptr, originIsAtTopLeft bool, heightOfTextTrack float64, defaultTextBoxOut *corefoundation.CGRect) int32 {
	result, callErr := tryCMTextFormatDescriptionGetDefaultTextBox(desc, originIsAtTopLeft, heightOfTextTrack, defaultTextBoxOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTextFormatDescriptionGetDisplayFlags func(desc uintptr, displayFlagsOut *CMTextDisplayFlags) int32
var _cMTextFormatDescriptionGetDisplayFlagsErr error

func tryCMTextFormatDescriptionGetDisplayFlags(desc uintptr, displayFlagsOut *CMTextDisplayFlags) (int32, error) {
	if _cMTextFormatDescriptionGetDisplayFlags == nil {
		return 0, symbolCallError("CMTextFormatDescriptionGetDisplayFlags", "10.7", _cMTextFormatDescriptionGetDisplayFlagsErr)
	}
	return _cMTextFormatDescriptionGetDisplayFlags(desc, displayFlagsOut), nil
}

// CMTextFormatDescriptionGetDisplayFlags returns the display flags.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionGetDisplayFlags(_:displayFlagsOut:)
func CMTextFormatDescriptionGetDisplayFlags(desc uintptr, displayFlagsOut *CMTextDisplayFlags) int32 {
	result, callErr := tryCMTextFormatDescriptionGetDisplayFlags(desc, displayFlagsOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTextFormatDescriptionGetFontName func(desc uintptr, localFontID uint16, fontNameOut *corefoundation.CFStringRef) int32
var _cMTextFormatDescriptionGetFontNameErr error

func tryCMTextFormatDescriptionGetFontName(desc uintptr, localFontID uint16, fontNameOut *corefoundation.CFStringRef) (int32, error) {
	if _cMTextFormatDescriptionGetFontName == nil {
		return 0, symbolCallError("CMTextFormatDescriptionGetFontName", "10.7", _cMTextFormatDescriptionGetFontNameErr)
	}
	return _cMTextFormatDescriptionGetFontName(desc, localFontID, fontNameOut), nil
}

// CMTextFormatDescriptionGetFontName returns a font name for a local font identifier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionGetFontName(_:localFontID:fontNameOut:)
func CMTextFormatDescriptionGetFontName(desc uintptr, localFontID uint16, fontNameOut *corefoundation.CFStringRef) int32 {
	result, callErr := tryCMTextFormatDescriptionGetFontName(desc, localFontID, fontNameOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTextFormatDescriptionGetJustification func(desc uintptr, horizontaJustificationlOut *CMTextJustificationValue, verticalJustificationOut *CMTextJustificationValue) int32
var _cMTextFormatDescriptionGetJustificationErr error

func tryCMTextFormatDescriptionGetJustification(desc uintptr, horizontaJustificationlOut *CMTextJustificationValue, verticalJustificationOut *CMTextJustificationValue) (int32, error) {
	if _cMTextFormatDescriptionGetJustification == nil {
		return 0, symbolCallError("CMTextFormatDescriptionGetJustification", "10.7", _cMTextFormatDescriptionGetJustificationErr)
	}
	return _cMTextFormatDescriptionGetJustification(desc, horizontaJustificationlOut, verticalJustificationOut), nil
}

// CMTextFormatDescriptionGetJustification returns the horizontal and vertical justification.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionGetJustification(_:horizontalOut:verticalOut:)
func CMTextFormatDescriptionGetJustification(desc uintptr, horizontaJustificationlOut *CMTextJustificationValue, verticalJustificationOut *CMTextJustificationValue) int32 {
	result, callErr := tryCMTextFormatDescriptionGetJustification(desc, horizontaJustificationlOut, verticalJustificationOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeAbsoluteValue func(time CMTime) CMTime
var _cMTimeAbsoluteValueErr error

func tryCMTimeAbsoluteValue(time CMTime) (CMTime, error) {
	if _cMTimeAbsoluteValue == nil {
		return CMTime{}, symbolCallError("CMTimeAbsoluteValue", "10.7", _cMTimeAbsoluteValueErr)
	}
	return _cMTimeAbsoluteValue(time), nil
}

// CMTimeAbsoluteValue returns the absolute value of a time.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeAbsoluteValue(_:)
func CMTimeAbsoluteValue(time CMTime) CMTime {
	result, callErr := tryCMTimeAbsoluteValue(time)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeAdd func(lhs CMTime, rhs CMTime) CMTime
var _cMTimeAddErr error

func tryCMTimeAdd(lhs CMTime, rhs CMTime) (CMTime, error) {
	if _cMTimeAdd == nil {
		return CMTime{}, symbolCallError("CMTimeAdd", "10.7", _cMTimeAddErr)
	}
	return _cMTimeAdd(lhs, rhs), nil
}

// CMTimeAdd returns the sum of two times.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeAdd(_:_:)
func CMTimeAdd(lhs CMTime, rhs CMTime) CMTime {
	result, callErr := tryCMTimeAdd(lhs, rhs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeClampToRange func(time CMTime, range_ CMTimeRange) CMTime
var _cMTimeClampToRangeErr error

func tryCMTimeClampToRange(time CMTime, range_ CMTimeRange) (CMTime, error) {
	if _cMTimeClampToRange == nil {
		return CMTime{}, symbolCallError("CMTimeClampToRange", "10.7", _cMTimeClampToRangeErr)
	}
	return _cMTimeClampToRange(time, range_), nil
}

// CMTimeClampToRange returns the nearest time value inside the time range.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeClampToRange(_:range:)
func CMTimeClampToRange(time CMTime, range_ CMTimeRange) CMTime {
	result, callErr := tryCMTimeClampToRange(time, range_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, timeCodeFormatDescription CMTimeCodeFormatDescriptionRef, flavor CMTimeCodeDescriptionFlavor, blockBufferOut *uintptr) int32
var _cMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBufferErr error

func tryCMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, timeCodeFormatDescription CMTimeCodeFormatDescriptionRef, flavor CMTimeCodeDescriptionFlavor, blockBufferOut *uintptr) (int32, error) {
	if _cMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer == nil {
		return 0, symbolCallError("CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer", "10.10", _cMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBufferErr)
	}
	return _cMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer(allocator, timeCodeFormatDescription, flavor, blockBufferOut), nil
}

// CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer copies the contents of a time code format description to a buffer in big-endian byte order.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer(allocator:timeCodeFormatDescription:flavor:blockBufferOut:)
func CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, timeCodeFormatDescription CMTimeCodeFormatDescriptionRef, flavor CMTimeCodeDescriptionFlavor, blockBufferOut *uintptr) int32 {
	result, callErr := tryCMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer(allocator, timeCodeFormatDescription, flavor, blockBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeCodeFormatDescriptionCreate func(allocator corefoundation.CFAllocatorRef, timeCodeFormatType CMTimeCodeFormatType, frameDuration CMTime, frameQuanta uint32, flags uint32, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *CMTimeCodeFormatDescriptionRef) int32
var _cMTimeCodeFormatDescriptionCreateErr error

func tryCMTimeCodeFormatDescriptionCreate(allocator corefoundation.CFAllocatorRef, timeCodeFormatType CMTimeCodeFormatType, frameDuration CMTime, frameQuanta uint32, flags uint32, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *CMTimeCodeFormatDescriptionRef) (int32, error) {
	if _cMTimeCodeFormatDescriptionCreate == nil {
		return 0, symbolCallError("CMTimeCodeFormatDescriptionCreate", "10.7", _cMTimeCodeFormatDescriptionCreateErr)
	}
	return _cMTimeCodeFormatDescriptionCreate(allocator, timeCodeFormatType, frameDuration, frameQuanta, flags, extensions, formatDescriptionOut), nil
}

// CMTimeCodeFormatDescriptionCreate creates a format description for time code media.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionCreate(allocator:timeCodeFormatType:frameDuration:frameQuanta:flags:extensions:formatDescriptionOut:)
func CMTimeCodeFormatDescriptionCreate(allocator corefoundation.CFAllocatorRef, timeCodeFormatType CMTimeCodeFormatType, frameDuration CMTime, frameQuanta uint32, flags uint32, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *CMTimeCodeFormatDescriptionRef) int32 {
	result, callErr := tryCMTimeCodeFormatDescriptionCreate(allocator, timeCodeFormatType, frameDuration, frameQuanta, flags, extensions, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, timeCodeDescriptionBlockBuffer uintptr, flavor CMTimeCodeDescriptionFlavor, formatDescriptionOut *CMTimeCodeFormatDescriptionRef) int32
var _cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBufferErr error

func tryCMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, timeCodeDescriptionBlockBuffer uintptr, flavor CMTimeCodeDescriptionFlavor, formatDescriptionOut *CMTimeCodeFormatDescriptionRef) (int32, error) {
	if _cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer == nil {
		return 0, symbolCallError("CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer", "10.10", _cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBufferErr)
	}
	return _cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer(allocator, timeCodeDescriptionBlockBuffer, flavor, formatDescriptionOut), nil
}

// CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer creates a time code format description from a big-endian time code description data structure in a buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer(allocator:bigEndianTimeCodeDescriptionBlockBuffer:flavor:formatDescriptionOut:)
func CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, timeCodeDescriptionBlockBuffer uintptr, flavor CMTimeCodeDescriptionFlavor, formatDescriptionOut *CMTimeCodeFormatDescriptionRef) int32 {
	result, callErr := tryCMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer(allocator, timeCodeDescriptionBlockBuffer, flavor, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData func(allocator corefoundation.CFAllocatorRef, timeCodeDescriptionData *uint8, size uintptr, flavor CMTimeCodeDescriptionFlavor, formatDescriptionOut *CMTimeCodeFormatDescriptionRef) int32
var _cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionDataErr error

func tryCMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData(allocator corefoundation.CFAllocatorRef, timeCodeDescriptionData *uint8, size uintptr, flavor CMTimeCodeDescriptionFlavor, formatDescriptionOut *CMTimeCodeFormatDescriptionRef) (int32, error) {
	if _cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData == nil {
		return 0, symbolCallError("CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData", "10.10", _cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionDataErr)
	}
	return _cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData(allocator, timeCodeDescriptionData, size, flavor, formatDescriptionOut), nil
}

// CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData creates a time code format description from a big-endian time code description structure.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData(allocator:bigEndianTimeCodeDescriptionData:size:flavor:formatDescriptionOut:)
func CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData(allocator corefoundation.CFAllocatorRef, timeCodeDescriptionData *uint8, size uintptr, flavor CMTimeCodeDescriptionFlavor, formatDescriptionOut *CMTimeCodeFormatDescriptionRef) int32 {
	result, callErr := tryCMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData(allocator, timeCodeDescriptionData, size, flavor, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeCodeFormatDescriptionGetFrameDuration func(timeCodeFormatDescription CMTimeCodeFormatDescriptionRef) CMTime
var _cMTimeCodeFormatDescriptionGetFrameDurationErr error

func tryCMTimeCodeFormatDescriptionGetFrameDuration(timeCodeFormatDescription CMTimeCodeFormatDescriptionRef) (CMTime, error) {
	if _cMTimeCodeFormatDescriptionGetFrameDuration == nil {
		return CMTime{}, symbolCallError("CMTimeCodeFormatDescriptionGetFrameDuration", "10.7", _cMTimeCodeFormatDescriptionGetFrameDurationErr)
	}
	return _cMTimeCodeFormatDescriptionGetFrameDuration(timeCodeFormatDescription), nil
}

// CMTimeCodeFormatDescriptionGetFrameDuration returns the duration of each frame.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionGetFrameDuration(_:)
func CMTimeCodeFormatDescriptionGetFrameDuration(timeCodeFormatDescription CMTimeCodeFormatDescriptionRef) CMTime {
	result, callErr := tryCMTimeCodeFormatDescriptionGetFrameDuration(timeCodeFormatDescription)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeCodeFormatDescriptionGetFrameQuanta func(timeCodeFormatDescription CMTimeCodeFormatDescriptionRef) uint32
var _cMTimeCodeFormatDescriptionGetFrameQuantaErr error

func tryCMTimeCodeFormatDescriptionGetFrameQuanta(timeCodeFormatDescription CMTimeCodeFormatDescriptionRef) (uint32, error) {
	if _cMTimeCodeFormatDescriptionGetFrameQuanta == nil {
		return 0, symbolCallError("CMTimeCodeFormatDescriptionGetFrameQuanta", "10.7", _cMTimeCodeFormatDescriptionGetFrameQuantaErr)
	}
	return _cMTimeCodeFormatDescriptionGetFrameQuanta(timeCodeFormatDescription), nil
}

// CMTimeCodeFormatDescriptionGetFrameQuanta returns the frames per second for a time code, or frames per tick in counter mode.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionGetFrameQuanta(_:)
func CMTimeCodeFormatDescriptionGetFrameQuanta(timeCodeFormatDescription CMTimeCodeFormatDescriptionRef) uint32 {
	result, callErr := tryCMTimeCodeFormatDescriptionGetFrameQuanta(timeCodeFormatDescription)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeCodeFormatDescriptionGetTimeCodeFlags func(desc CMTimeCodeFormatDescriptionRef) uint32
var _cMTimeCodeFormatDescriptionGetTimeCodeFlagsErr error

func tryCMTimeCodeFormatDescriptionGetTimeCodeFlags(desc CMTimeCodeFormatDescriptionRef) (uint32, error) {
	if _cMTimeCodeFormatDescriptionGetTimeCodeFlags == nil {
		return 0, symbolCallError("CMTimeCodeFormatDescriptionGetTimeCodeFlags", "10.7", _cMTimeCodeFormatDescriptionGetTimeCodeFlagsErr)
	}
	return _cMTimeCodeFormatDescriptionGetTimeCodeFlags(desc), nil
}

// CMTimeCodeFormatDescriptionGetTimeCodeFlags returns time code flags.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionGetTimeCodeFlags(_:)
func CMTimeCodeFormatDescriptionGetTimeCodeFlags(desc CMTimeCodeFormatDescriptionRef) uint32 {
	result, callErr := tryCMTimeCodeFormatDescriptionGetTimeCodeFlags(desc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeCompare func(time1 CMTime, time2 CMTime) int32
var _cMTimeCompareErr error

func tryCMTimeCompare(time1 CMTime, time2 CMTime) (int32, error) {
	if _cMTimeCompare == nil {
		return 0, symbolCallError("CMTimeCompare", "10.7", _cMTimeCompareErr)
	}
	return _cMTimeCompare(time1, time2), nil
}

// CMTimeCompare returns the numerical relationship of two times.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCompare(_:_:)
func CMTimeCompare(time1 CMTime, time2 CMTime) int32 {
	result, callErr := tryCMTimeCompare(time1, time2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeConvertScale func(time CMTime, newTimescale int32, method CMTimeRoundingMethod) CMTime
var _cMTimeConvertScaleErr error

func tryCMTimeConvertScale(time CMTime, newTimescale int32, method CMTimeRoundingMethod) (CMTime, error) {
	if _cMTimeConvertScale == nil {
		return CMTime{}, symbolCallError("CMTimeConvertScale", "10.7", _cMTimeConvertScaleErr)
	}
	return _cMTimeConvertScale(time, newTimescale, method), nil
}

// CMTimeConvertScale converts the source time to a new timescale using the specified rounding method.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeConvertScale(_:timescale:method:)
func CMTimeConvertScale(time CMTime, newTimescale int32, method CMTimeRoundingMethod) CMTime {
	result, callErr := tryCMTimeConvertScale(time, newTimescale, method)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeCopyAsDictionary func(time CMTime, allocator corefoundation.CFAllocatorRef) corefoundation.CFDictionaryRef
var _cMTimeCopyAsDictionaryErr error

func tryCMTimeCopyAsDictionary(time CMTime, allocator corefoundation.CFAllocatorRef) (corefoundation.CFDictionaryRef, error) {
	if _cMTimeCopyAsDictionary == nil {
		return 0, symbolCallError("CMTimeCopyAsDictionary", "10.7", _cMTimeCopyAsDictionaryErr)
	}
	return _cMTimeCopyAsDictionary(time, allocator), nil
}

// CMTimeCopyAsDictionary creates a dictionary representation of the time.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCopyAsDictionary(_:allocator:)
func CMTimeCopyAsDictionary(time CMTime, allocator corefoundation.CFAllocatorRef) corefoundation.CFDictionaryRef {
	result, callErr := tryCMTimeCopyAsDictionary(time, allocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeCopyDescription func(allocator corefoundation.CFAllocatorRef, time CMTime) corefoundation.CFStringRef
var _cMTimeCopyDescriptionErr error

func tryCMTimeCopyDescription(allocator corefoundation.CFAllocatorRef, time CMTime) (corefoundation.CFStringRef, error) {
	if _cMTimeCopyDescription == nil {
		return 0, symbolCallError("CMTimeCopyDescription", "10.7", _cMTimeCopyDescriptionErr)
	}
	return _cMTimeCopyDescription(allocator, time), nil
}

// CMTimeCopyDescription creates a string representation of the time.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCopyDescription(allocator:time:)
func CMTimeCopyDescription(allocator corefoundation.CFAllocatorRef, time CMTime) corefoundation.CFStringRef {
	result, callErr := tryCMTimeCopyDescription(allocator, time)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeFoldIntoRange func(time CMTime, foldRange CMTimeRange) CMTime
var _cMTimeFoldIntoRangeErr error

func tryCMTimeFoldIntoRange(time CMTime, foldRange CMTimeRange) (CMTime, error) {
	if _cMTimeFoldIntoRange == nil {
		return CMTime{}, symbolCallError("CMTimeFoldIntoRange", "10.14", _cMTimeFoldIntoRangeErr)
	}
	return _cMTimeFoldIntoRange(time, foldRange), nil
}

// CMTimeFoldIntoRange folds a time into a time range.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeFoldIntoRange(_:foldRange:)
func CMTimeFoldIntoRange(time CMTime, foldRange CMTimeRange) CMTime {
	result, callErr := tryCMTimeFoldIntoRange(time, foldRange)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeGetSeconds func(time CMTime) float64
var _cMTimeGetSecondsErr error

func tryCMTimeGetSeconds(time CMTime) (float64, error) {
	if _cMTimeGetSeconds == nil {
		return 0.0, symbolCallError("CMTimeGetSeconds", "10.7", _cMTimeGetSecondsErr)
	}
	return _cMTimeGetSeconds(time), nil
}

// CMTimeGetSeconds returns a representation of the time in seconds.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeGetSeconds(_:)
func CMTimeGetSeconds(time CMTime) float64 {
	result, callErr := tryCMTimeGetSeconds(time)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeMake func(value int64, timescale int32) CMTime
var _cMTimeMakeErr error

func tryCMTimeMake(value int64, timescale int32) (CMTime, error) {
	if _cMTimeMake == nil {
		return CMTime{}, symbolCallError("CMTimeMake", "10.7", _cMTimeMakeErr)
	}
	return _cMTimeMake(value, timescale), nil
}

// CMTimeMake creates a time with a value and timescale.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMake(value:timescale:)
func CMTimeMake(value int64, timescale int32) CMTime {
	result, callErr := tryCMTimeMake(value, timescale)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeMakeFromDictionary func(dictionaryRepresentation corefoundation.CFDictionaryRef) CMTime
var _cMTimeMakeFromDictionaryErr error

func tryCMTimeMakeFromDictionary(dictionaryRepresentation corefoundation.CFDictionaryRef) (CMTime, error) {
	if _cMTimeMakeFromDictionary == nil {
		return CMTime{}, symbolCallError("CMTimeMakeFromDictionary", "10.7", _cMTimeMakeFromDictionaryErr)
	}
	return _cMTimeMakeFromDictionary(dictionaryRepresentation), nil
}

// CMTimeMakeFromDictionary creates a time from a dictionary representation of its fields.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMakeFromDictionary(_:)
func CMTimeMakeFromDictionary(dictionaryRepresentation corefoundation.CFDictionaryRef) CMTime {
	result, callErr := tryCMTimeMakeFromDictionary(dictionaryRepresentation)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeMakeWithEpoch func(value int64, timescale int32, epoch int64) CMTime
var _cMTimeMakeWithEpochErr error

func tryCMTimeMakeWithEpoch(value int64, timescale int32, epoch int64) (CMTime, error) {
	if _cMTimeMakeWithEpoch == nil {
		return CMTime{}, symbolCallError("CMTimeMakeWithEpoch", "10.7", _cMTimeMakeWithEpochErr)
	}
	return _cMTimeMakeWithEpoch(value, timescale, epoch), nil
}

// CMTimeMakeWithEpoch creates a time with a value, timescale, and epoch.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMakeWithEpoch(value:timescale:epoch:)
func CMTimeMakeWithEpoch(value int64, timescale int32, epoch int64) CMTime {
	result, callErr := tryCMTimeMakeWithEpoch(value, timescale, epoch)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeMakeWithSeconds func(seconds float64, preferredTimescale int32) CMTime
var _cMTimeMakeWithSecondsErr error

func tryCMTimeMakeWithSeconds(seconds float64, preferredTimescale int32) (CMTime, error) {
	if _cMTimeMakeWithSeconds == nil {
		return CMTime{}, symbolCallError("CMTimeMakeWithSeconds", "10.7", _cMTimeMakeWithSecondsErr)
	}
	return _cMTimeMakeWithSeconds(seconds, preferredTimescale), nil
}

// CMTimeMakeWithSeconds creates a time that represents a number of seconds in a preferred timescale.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMakeWithSeconds(_:preferredTimescale:)
func CMTimeMakeWithSeconds(seconds float64, preferredTimescale int32) CMTime {
	result, callErr := tryCMTimeMakeWithSeconds(seconds, preferredTimescale)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeMapDurationFromRangeToRange func(dur CMTime, fromRange CMTimeRange, toRange CMTimeRange) CMTime
var _cMTimeMapDurationFromRangeToRangeErr error

func tryCMTimeMapDurationFromRangeToRange(dur CMTime, fromRange CMTimeRange, toRange CMTimeRange) (CMTime, error) {
	if _cMTimeMapDurationFromRangeToRange == nil {
		return CMTime{}, symbolCallError("CMTimeMapDurationFromRangeToRange", "10.7", _cMTimeMapDurationFromRangeToRangeErr)
	}
	return _cMTimeMapDurationFromRangeToRange(dur, fromRange, toRange), nil
}

// CMTimeMapDurationFromRangeToRange translates a duration through a mapping from two time ranges.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMapDurationFromRangeToRange(_:fromRange:toRange:)
func CMTimeMapDurationFromRangeToRange(dur CMTime, fromRange CMTimeRange, toRange CMTimeRange) CMTime {
	result, callErr := tryCMTimeMapDurationFromRangeToRange(dur, fromRange, toRange)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeMapTimeFromRangeToRange func(t CMTime, fromRange CMTimeRange, toRange CMTimeRange) CMTime
var _cMTimeMapTimeFromRangeToRangeErr error

func tryCMTimeMapTimeFromRangeToRange(t CMTime, fromRange CMTimeRange, toRange CMTimeRange) (CMTime, error) {
	if _cMTimeMapTimeFromRangeToRange == nil {
		return CMTime{}, symbolCallError("CMTimeMapTimeFromRangeToRange", "10.7", _cMTimeMapTimeFromRangeToRangeErr)
	}
	return _cMTimeMapTimeFromRangeToRange(t, fromRange, toRange), nil
}

// CMTimeMapTimeFromRangeToRange translates a time through a mapping from two time ranges.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMapTimeFromRangeToRange(_:fromRange:toRange:)
func CMTimeMapTimeFromRangeToRange(t CMTime, fromRange CMTimeRange, toRange CMTimeRange) CMTime {
	result, callErr := tryCMTimeMapTimeFromRangeToRange(t, fromRange, toRange)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeMappingCopyAsDictionary func(mapping CMTimeMapping, allocator corefoundation.CFAllocatorRef) corefoundation.CFDictionaryRef
var _cMTimeMappingCopyAsDictionaryErr error

func tryCMTimeMappingCopyAsDictionary(mapping CMTimeMapping, allocator corefoundation.CFAllocatorRef) (corefoundation.CFDictionaryRef, error) {
	if _cMTimeMappingCopyAsDictionary == nil {
		return 0, symbolCallError("CMTimeMappingCopyAsDictionary", "10.11", _cMTimeMappingCopyAsDictionaryErr)
	}
	return _cMTimeMappingCopyAsDictionary(mapping, allocator), nil
}

// CMTimeMappingCopyAsDictionary returns a dictionary representation of a time mapping.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingCopyAsDictionary(_:allocator:)
func CMTimeMappingCopyAsDictionary(mapping CMTimeMapping, allocator corefoundation.CFAllocatorRef) corefoundation.CFDictionaryRef {
	result, callErr := tryCMTimeMappingCopyAsDictionary(mapping, allocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeMappingCopyDescription func(allocator corefoundation.CFAllocatorRef, mapping CMTimeMapping) corefoundation.CFStringRef
var _cMTimeMappingCopyDescriptionErr error

func tryCMTimeMappingCopyDescription(allocator corefoundation.CFAllocatorRef, mapping CMTimeMapping) (corefoundation.CFStringRef, error) {
	if _cMTimeMappingCopyDescription == nil {
		return 0, symbolCallError("CMTimeMappingCopyDescription", "10.11", _cMTimeMappingCopyDescriptionErr)
	}
	return _cMTimeMappingCopyDescription(allocator, mapping), nil
}

// CMTimeMappingCopyDescription copies a string description of a time mapping.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingCopyDescription(allocator:mapping:)
func CMTimeMappingCopyDescription(allocator corefoundation.CFAllocatorRef, mapping CMTimeMapping) corefoundation.CFStringRef {
	result, callErr := tryCMTimeMappingCopyDescription(allocator, mapping)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeMappingMake func(source CMTimeRange, target CMTimeRange) CMTimeMapping
var _cMTimeMappingMakeErr error

func tryCMTimeMappingMake(source CMTimeRange, target CMTimeRange) (CMTimeMapping, error) {
	if _cMTimeMappingMake == nil {
		return CMTimeMapping{}, symbolCallError("CMTimeMappingMake", "10.11", _cMTimeMappingMakeErr)
	}
	return _cMTimeMappingMake(source, target), nil
}

// CMTimeMappingMake creates a time mapping with a source and target time range.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingMake(source:target:)
func CMTimeMappingMake(source CMTimeRange, target CMTimeRange) CMTimeMapping {
	result, callErr := tryCMTimeMappingMake(source, target)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeMappingMakeEmpty func(target CMTimeRange) CMTimeMapping
var _cMTimeMappingMakeEmptyErr error

func tryCMTimeMappingMakeEmpty(target CMTimeRange) (CMTimeMapping, error) {
	if _cMTimeMappingMakeEmpty == nil {
		return CMTimeMapping{}, symbolCallError("CMTimeMappingMakeEmpty", "10.11", _cMTimeMappingMakeEmptyErr)
	}
	return _cMTimeMappingMakeEmpty(target), nil
}

// CMTimeMappingMakeEmpty creates a valid time mapping with an empty source.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingMakeEmpty(target:)
func CMTimeMappingMakeEmpty(target CMTimeRange) CMTimeMapping {
	result, callErr := tryCMTimeMappingMakeEmpty(target)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeMappingMakeFromDictionary func(dictionaryRepresentation corefoundation.CFDictionaryRef) CMTimeMapping
var _cMTimeMappingMakeFromDictionaryErr error

func tryCMTimeMappingMakeFromDictionary(dictionaryRepresentation corefoundation.CFDictionaryRef) (CMTimeMapping, error) {
	if _cMTimeMappingMakeFromDictionary == nil {
		return CMTimeMapping{}, symbolCallError("CMTimeMappingMakeFromDictionary", "10.11", _cMTimeMappingMakeFromDictionaryErr)
	}
	return _cMTimeMappingMakeFromDictionary(dictionaryRepresentation), nil
}

// CMTimeMappingMakeFromDictionary creates a time mapping from a dictionary representation.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingMakeFromDictionary(_:)
func CMTimeMappingMakeFromDictionary(dictionaryRepresentation corefoundation.CFDictionaryRef) CMTimeMapping {
	result, callErr := tryCMTimeMappingMakeFromDictionary(dictionaryRepresentation)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeMappingShow func(mapping CMTimeMapping)
var _cMTimeMappingShowErr error

func tryCMTimeMappingShow(mapping CMTimeMapping) error {
	if _cMTimeMappingShow == nil {
		return symbolCallError("CMTimeMappingShow", "10.11", _cMTimeMappingShowErr)
	}
	_cMTimeMappingShow(mapping)
	return nil
}

// CMTimeMappingShow prints a description of a time mapping to standard output.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingShow(_:)
func CMTimeMappingShow(mapping CMTimeMapping) {
	if callErr := tryCMTimeMappingShow(mapping); callErr != nil {
		panic(callErr)
	}
}

var _cMTimeMaximum func(time1 CMTime, time2 CMTime) CMTime
var _cMTimeMaximumErr error

func tryCMTimeMaximum(time1 CMTime, time2 CMTime) (CMTime, error) {
	if _cMTimeMaximum == nil {
		return CMTime{}, symbolCallError("CMTimeMaximum", "10.7", _cMTimeMaximumErr)
	}
	return _cMTimeMaximum(time1, time2), nil
}

// CMTimeMaximum returns the greater of two time values.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMaximum(_:_:)
func CMTimeMaximum(time1 CMTime, time2 CMTime) CMTime {
	result, callErr := tryCMTimeMaximum(time1, time2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeMinimum func(time1 CMTime, time2 CMTime) CMTime
var _cMTimeMinimumErr error

func tryCMTimeMinimum(time1 CMTime, time2 CMTime) (CMTime, error) {
	if _cMTimeMinimum == nil {
		return CMTime{}, symbolCallError("CMTimeMinimum", "10.7", _cMTimeMinimumErr)
	}
	return _cMTimeMinimum(time1, time2), nil
}

// CMTimeMinimum returns the lesser of two time values.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMinimum(_:_:)
func CMTimeMinimum(time1 CMTime, time2 CMTime) CMTime {
	result, callErr := tryCMTimeMinimum(time1, time2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeMultiply func(time CMTime, multiplier int32) CMTime
var _cMTimeMultiplyErr error

func tryCMTimeMultiply(time CMTime, multiplier int32) (CMTime, error) {
	if _cMTimeMultiply == nil {
		return CMTime{}, symbolCallError("CMTimeMultiply", "10.7", _cMTimeMultiplyErr)
	}
	return _cMTimeMultiply(time, multiplier), nil
}

// CMTimeMultiply returns the result of multiplying a time by an integer multiplier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMultiply(_:multiplier:)
func CMTimeMultiply(time CMTime, multiplier int32) CMTime {
	result, callErr := tryCMTimeMultiply(time, multiplier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeMultiplyByFloat64 func(time CMTime, multiplier float64) CMTime
var _cMTimeMultiplyByFloat64Err error

func tryCMTimeMultiplyByFloat64(time CMTime, multiplier float64) (CMTime, error) {
	if _cMTimeMultiplyByFloat64 == nil {
		return CMTime{}, symbolCallError("CMTimeMultiplyByFloat64", "10.7", _cMTimeMultiplyByFloat64Err)
	}
	return _cMTimeMultiplyByFloat64(time, multiplier), nil
}

// CMTimeMultiplyByFloat64 returns the result of multiplying a time by a floating-point multiplier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMultiplyByFloat64(_:multiplier:)
func CMTimeMultiplyByFloat64(time CMTime, multiplier float64) CMTime {
	result, callErr := tryCMTimeMultiplyByFloat64(time, multiplier)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeMultiplyByRatio func(time CMTime, multiplier int32, divisor int32) CMTime
var _cMTimeMultiplyByRatioErr error

func tryCMTimeMultiplyByRatio(time CMTime, multiplier int32, divisor int32) (CMTime, error) {
	if _cMTimeMultiplyByRatio == nil {
		return CMTime{}, symbolCallError("CMTimeMultiplyByRatio", "10.10", _cMTimeMultiplyByRatioErr)
	}
	return _cMTimeMultiplyByRatio(time, multiplier, divisor), nil
}

// CMTimeMultiplyByRatio returns the result of multiplying a time by an integer multiplier, and then dividing the result by the divisor.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMultiplyByRatio(_:multiplier:divisor:)
func CMTimeMultiplyByRatio(time CMTime, multiplier int32, divisor int32) CMTime {
	result, callErr := tryCMTimeMultiplyByRatio(time, multiplier, divisor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeRangeContainsTime func(range_ CMTimeRange, time CMTime) bool
var _cMTimeRangeContainsTimeErr error

func tryCMTimeRangeContainsTime(range_ CMTimeRange, time CMTime) (bool, error) {
	if _cMTimeRangeContainsTime == nil {
		return false, symbolCallError("CMTimeRangeContainsTime", "10.7", _cMTimeRangeContainsTimeErr)
	}
	return _cMTimeRangeContainsTime(range_, time), nil
}

// CMTimeRangeContainsTime returns a Boolean value that indicates whether a time range contains a time.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeContainsTime(_:time:)
func CMTimeRangeContainsTime(range_ CMTimeRange, time CMTime) bool {
	result, callErr := tryCMTimeRangeContainsTime(range_, time)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeRangeContainsTimeRange func(range_ CMTimeRange, otherRange CMTimeRange) bool
var _cMTimeRangeContainsTimeRangeErr error

func tryCMTimeRangeContainsTimeRange(range_ CMTimeRange, otherRange CMTimeRange) (bool, error) {
	if _cMTimeRangeContainsTimeRange == nil {
		return false, symbolCallError("CMTimeRangeContainsTimeRange", "10.7", _cMTimeRangeContainsTimeRangeErr)
	}
	return _cMTimeRangeContainsTimeRange(range_, otherRange), nil
}

// CMTimeRangeContainsTimeRange returns a Boolean value that indicates whether a time range contains another time range.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeContainsTimeRange(_:otherRange:)
func CMTimeRangeContainsTimeRange(range_ CMTimeRange, otherRange CMTimeRange) bool {
	result, callErr := tryCMTimeRangeContainsTimeRange(range_, otherRange)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeRangeCopyAsDictionary func(range_ CMTimeRange, allocator corefoundation.CFAllocatorRef) corefoundation.CFDictionaryRef
var _cMTimeRangeCopyAsDictionaryErr error

func tryCMTimeRangeCopyAsDictionary(range_ CMTimeRange, allocator corefoundation.CFAllocatorRef) (corefoundation.CFDictionaryRef, error) {
	if _cMTimeRangeCopyAsDictionary == nil {
		return 0, symbolCallError("CMTimeRangeCopyAsDictionary", "10.7", _cMTimeRangeCopyAsDictionaryErr)
	}
	return _cMTimeRangeCopyAsDictionary(range_, allocator), nil
}

// CMTimeRangeCopyAsDictionary returns a dictionary representation of a time range.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeCopyAsDictionary(_:allocator:)
func CMTimeRangeCopyAsDictionary(range_ CMTimeRange, allocator corefoundation.CFAllocatorRef) corefoundation.CFDictionaryRef {
	result, callErr := tryCMTimeRangeCopyAsDictionary(range_, allocator)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeRangeCopyDescription func(allocator corefoundation.CFAllocatorRef, range_ CMTimeRange) corefoundation.CFStringRef
var _cMTimeRangeCopyDescriptionErr error

func tryCMTimeRangeCopyDescription(allocator corefoundation.CFAllocatorRef, range_ CMTimeRange) (corefoundation.CFStringRef, error) {
	if _cMTimeRangeCopyDescription == nil {
		return 0, symbolCallError("CMTimeRangeCopyDescription", "10.7", _cMTimeRangeCopyDescriptionErr)
	}
	return _cMTimeRangeCopyDescription(allocator, range_), nil
}

// CMTimeRangeCopyDescription returns a string with a description of a time range.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeCopyDescription(allocator:range:)
func CMTimeRangeCopyDescription(allocator corefoundation.CFAllocatorRef, range_ CMTimeRange) corefoundation.CFStringRef {
	result, callErr := tryCMTimeRangeCopyDescription(allocator, range_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeRangeEqual func(range1 CMTimeRange, range2 CMTimeRange) bool
var _cMTimeRangeEqualErr error

func tryCMTimeRangeEqual(range1 CMTimeRange, range2 CMTimeRange) (bool, error) {
	if _cMTimeRangeEqual == nil {
		return false, symbolCallError("CMTimeRangeEqual", "10.7", _cMTimeRangeEqualErr)
	}
	return _cMTimeRangeEqual(range1, range2), nil
}

// CMTimeRangeEqual returns a Boolean value that indicates whether two time ranges are equal.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeEqual(_:_:)
func CMTimeRangeEqual(range1 CMTimeRange, range2 CMTimeRange) bool {
	result, callErr := tryCMTimeRangeEqual(range1, range2)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeRangeFromTimeToTime func(start CMTime, end CMTime) CMTimeRange
var _cMTimeRangeFromTimeToTimeErr error

func tryCMTimeRangeFromTimeToTime(start CMTime, end CMTime) (CMTimeRange, error) {
	if _cMTimeRangeFromTimeToTime == nil {
		return CMTimeRange{}, symbolCallError("CMTimeRangeFromTimeToTime", "10.7", _cMTimeRangeFromTimeToTimeErr)
	}
	return _cMTimeRangeFromTimeToTime(start, end), nil
}

// CMTimeRangeFromTimeToTime creates a valid time range from a start and end time.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeFromTimeToTime(start:end:)
func CMTimeRangeFromTimeToTime(start CMTime, end CMTime) CMTimeRange {
	result, callErr := tryCMTimeRangeFromTimeToTime(start, end)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeRangeGetEnd func(range_ CMTimeRange) CMTime
var _cMTimeRangeGetEndErr error

func tryCMTimeRangeGetEnd(range_ CMTimeRange) (CMTime, error) {
	if _cMTimeRangeGetEnd == nil {
		return CMTime{}, symbolCallError("CMTimeRangeGetEnd", "10.7", _cMTimeRangeGetEndErr)
	}
	return _cMTimeRangeGetEnd(range_), nil
}

// CMTimeRangeGetEnd returns a time value that represents the end of a time range.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeGetEnd(_:)
func CMTimeRangeGetEnd(range_ CMTimeRange) CMTime {
	result, callErr := tryCMTimeRangeGetEnd(range_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeRangeGetIntersection func(range_ CMTimeRange, otherRange CMTimeRange) CMTimeRange
var _cMTimeRangeGetIntersectionErr error

func tryCMTimeRangeGetIntersection(range_ CMTimeRange, otherRange CMTimeRange) (CMTimeRange, error) {
	if _cMTimeRangeGetIntersection == nil {
		return CMTimeRange{}, symbolCallError("CMTimeRangeGetIntersection", "10.7", _cMTimeRangeGetIntersectionErr)
	}
	return _cMTimeRangeGetIntersection(range_, otherRange), nil
}

// CMTimeRangeGetIntersection returns a new time range with the time elements that are common between the input.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeGetIntersection(_:otherRange:)
func CMTimeRangeGetIntersection(range_ CMTimeRange, otherRange CMTimeRange) CMTimeRange {
	result, callErr := tryCMTimeRangeGetIntersection(range_, otherRange)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeRangeGetUnion func(range_ CMTimeRange, otherRange CMTimeRange) CMTimeRange
var _cMTimeRangeGetUnionErr error

func tryCMTimeRangeGetUnion(range_ CMTimeRange, otherRange CMTimeRange) (CMTimeRange, error) {
	if _cMTimeRangeGetUnion == nil {
		return CMTimeRange{}, symbolCallError("CMTimeRangeGetUnion", "10.7", _cMTimeRangeGetUnionErr)
	}
	return _cMTimeRangeGetUnion(range_, otherRange), nil
}

// CMTimeRangeGetUnion returns a new time range with the time elements of the input.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeGetUnion(_:otherRange:)
func CMTimeRangeGetUnion(range_ CMTimeRange, otherRange CMTimeRange) CMTimeRange {
	result, callErr := tryCMTimeRangeGetUnion(range_, otherRange)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeRangeMake func(start CMTime, duration CMTime) CMTimeRange
var _cMTimeRangeMakeErr error

func tryCMTimeRangeMake(start CMTime, duration CMTime) (CMTimeRange, error) {
	if _cMTimeRangeMake == nil {
		return CMTimeRange{}, symbolCallError("CMTimeRangeMake", "10.7", _cMTimeRangeMakeErr)
	}
	return _cMTimeRangeMake(start, duration), nil
}

// CMTimeRangeMake creates a valid time range with a start time and duration.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeMake(start:duration:)
func CMTimeRangeMake(start CMTime, duration CMTime) CMTimeRange {
	result, callErr := tryCMTimeRangeMake(start, duration)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeRangeMakeFromDictionary func(dictionaryRepresentation corefoundation.CFDictionaryRef) CMTimeRange
var _cMTimeRangeMakeFromDictionaryErr error

func tryCMTimeRangeMakeFromDictionary(dictionaryRepresentation corefoundation.CFDictionaryRef) (CMTimeRange, error) {
	if _cMTimeRangeMakeFromDictionary == nil {
		return CMTimeRange{}, symbolCallError("CMTimeRangeMakeFromDictionary", "10.7", _cMTimeRangeMakeFromDictionaryErr)
	}
	return _cMTimeRangeMakeFromDictionary(dictionaryRepresentation), nil
}

// CMTimeRangeMakeFromDictionary creates a time range from a dictionary representation of its fields.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeMakeFromDictionary(_:)
func CMTimeRangeMakeFromDictionary(dictionaryRepresentation corefoundation.CFDictionaryRef) CMTimeRange {
	result, callErr := tryCMTimeRangeMakeFromDictionary(dictionaryRepresentation)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimeRangeShow func(range_ CMTimeRange)
var _cMTimeRangeShowErr error

func tryCMTimeRangeShow(range_ CMTimeRange) error {
	if _cMTimeRangeShow == nil {
		return symbolCallError("CMTimeRangeShow", "10.7", _cMTimeRangeShowErr)
	}
	_cMTimeRangeShow(range_)
	return nil
}

// CMTimeRangeShow prints a description of the time range to standard error.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeShow(_:)
func CMTimeRangeShow(range_ CMTimeRange) {
	if callErr := tryCMTimeRangeShow(range_); callErr != nil {
		panic(callErr)
	}
}

var _cMTimeShow func(time CMTime)
var _cMTimeShowErr error

func tryCMTimeShow(time CMTime) error {
	if _cMTimeShow == nil {
		return symbolCallError("CMTimeShow", "10.7", _cMTimeShowErr)
	}
	_cMTimeShow(time)
	return nil
}

// CMTimeShow prints a description of the time to the console.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeShow(_:)
func CMTimeShow(time CMTime) {
	if callErr := tryCMTimeShow(time); callErr != nil {
		panic(callErr)
	}
}

var _cMTimeSubtract func(lhs CMTime, rhs CMTime) CMTime
var _cMTimeSubtractErr error

func tryCMTimeSubtract(lhs CMTime, rhs CMTime) (CMTime, error) {
	if _cMTimeSubtract == nil {
		return CMTime{}, symbolCallError("CMTimeSubtract", "10.7", _cMTimeSubtractErr)
	}
	return _cMTimeSubtract(lhs, rhs), nil
}

// CMTimeSubtract returns the difference between two times.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeSubtract(_:_:)
func CMTimeSubtract(lhs CMTime, rhs CMTime) CMTime {
	result, callErr := tryCMTimeSubtract(lhs, rhs)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseAddTimer func(timebase uintptr, timer corefoundation.CFRunLoopTimerRef, runloop corefoundation.CFRunLoopRef) int32
var _cMTimebaseAddTimerErr error

func tryCMTimebaseAddTimer(timebase uintptr, timer corefoundation.CFRunLoopTimerRef, runloop corefoundation.CFRunLoopRef) (int32, error) {
	if _cMTimebaseAddTimer == nil {
		return 0, symbolCallError("CMTimebaseAddTimer", "10.8", _cMTimebaseAddTimerErr)
	}
	return _cMTimebaseAddTimer(timebase, timer, runloop), nil
}

// CMTimebaseAddTimer adds the timer to the list of timers the timebase manages.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseAddTimer(_:timer:runloop:)
func CMTimebaseAddTimer(timebase uintptr, timer corefoundation.CFRunLoopTimerRef, runloop corefoundation.CFRunLoopRef) int32 {
	result, callErr := tryCMTimebaseAddTimer(timebase, timer, runloop)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseAddTimerDispatchSource func(timebase uintptr, timerSource uintptr) int32
var _cMTimebaseAddTimerDispatchSourceErr error

func tryCMTimebaseAddTimerDispatchSource(timebase uintptr, timerSource dispatch.Source) (int32, error) {
	if _cMTimebaseAddTimerDispatchSource == nil {
		return 0, symbolCallError("CMTimebaseAddTimerDispatchSource", "10.8", _cMTimebaseAddTimerDispatchSourceErr)
	}
	return _cMTimebaseAddTimerDispatchSource(timebase, uintptr(timerSource.Handle())), nil
}

// CMTimebaseAddTimerDispatchSource adds the timer dispatch source to the list of timers the timebase manages.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseAddTimerDispatchSource(_:timerSource:)
func CMTimebaseAddTimerDispatchSource(timebase uintptr, timerSource dispatch.Source) int32 {
	result, callErr := tryCMTimebaseAddTimerDispatchSource(timebase, timerSource)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseCopySource func(timebase uintptr) CMClockOrTimebaseRef
var _cMTimebaseCopySourceErr error

func tryCMTimebaseCopySource(timebase uintptr) (CMClockOrTimebaseRef, error) {
	if _cMTimebaseCopySource == nil {
		return 0, symbolCallError("CMTimebaseCopySource", "10.11", _cMTimebaseCopySourceErr)
	}
	return _cMTimebaseCopySource(timebase), nil
}

// CMTimebaseCopySource returns the immediate source — either a clock or timebase — of a timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopySource(_:)
func CMTimebaseCopySource(timebase uintptr) CMClockOrTimebaseRef {
	result, callErr := tryCMTimebaseCopySource(timebase)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseCopySourceClock func(timebase uintptr) uintptr
var _cMTimebaseCopySourceClockErr error

func tryCMTimebaseCopySourceClock(timebase uintptr) (uintptr, error) {
	if _cMTimebaseCopySourceClock == nil {
		return 0, symbolCallError("CMTimebaseCopySourceClock", "10.11", _cMTimebaseCopySourceClockErr)
	}
	return _cMTimebaseCopySourceClock(timebase), nil
}

// CMTimebaseCopySourceClock returns the immediate source clock of a timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopySourceClock(_:)
func CMTimebaseCopySourceClock(timebase uintptr) uintptr {
	result, callErr := tryCMTimebaseCopySourceClock(timebase)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseCopySourceTimebase func(timebase uintptr) uintptr
var _cMTimebaseCopySourceTimebaseErr error

func tryCMTimebaseCopySourceTimebase(timebase uintptr) (uintptr, error) {
	if _cMTimebaseCopySourceTimebase == nil {
		return 0, symbolCallError("CMTimebaseCopySourceTimebase", "10.11", _cMTimebaseCopySourceTimebaseErr)
	}
	return _cMTimebaseCopySourceTimebase(timebase), nil
}

// CMTimebaseCopySourceTimebase returns the immediate source timebase of a timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopySourceTimebase(_:)
func CMTimebaseCopySourceTimebase(timebase uintptr) uintptr {
	result, callErr := tryCMTimebaseCopySourceTimebase(timebase)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseCopyUltimateSourceClock func(timebase uintptr) uintptr
var _cMTimebaseCopyUltimateSourceClockErr error

func tryCMTimebaseCopyUltimateSourceClock(timebase uintptr) (uintptr, error) {
	if _cMTimebaseCopyUltimateSourceClock == nil {
		return 0, symbolCallError("CMTimebaseCopyUltimateSourceClock", "10.11", _cMTimebaseCopyUltimateSourceClockErr)
	}
	return _cMTimebaseCopyUltimateSourceClock(timebase), nil
}

// CMTimebaseCopyUltimateSourceClock returns the source clock that’s the source of all of a timebase’s source timebases.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopyUltimateSourceClock(_:)
func CMTimebaseCopyUltimateSourceClock(timebase uintptr) uintptr {
	result, callErr := tryCMTimebaseCopyUltimateSourceClock(timebase)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseCreateWithSourceClock func(allocator corefoundation.CFAllocatorRef, sourceClock uintptr, timebaseOut *uintptr) int32
var _cMTimebaseCreateWithSourceClockErr error

func tryCMTimebaseCreateWithSourceClock(allocator corefoundation.CFAllocatorRef, sourceClock uintptr, timebaseOut *uintptr) (int32, error) {
	if _cMTimebaseCreateWithSourceClock == nil {
		return 0, symbolCallError("CMTimebaseCreateWithSourceClock", "10.8", _cMTimebaseCreateWithSourceClockErr)
	}
	return _cMTimebaseCreateWithSourceClock(allocator, sourceClock, timebaseOut), nil
}

// CMTimebaseCreateWithSourceClock creates a timebase by using a source clock.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCreateWithSourceClock(allocator:sourceClock:timebaseOut:)
func CMTimebaseCreateWithSourceClock(allocator corefoundation.CFAllocatorRef, sourceClock uintptr, timebaseOut *uintptr) int32 {
	result, callErr := tryCMTimebaseCreateWithSourceClock(allocator, sourceClock, timebaseOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseCreateWithSourceTimebase func(allocator corefoundation.CFAllocatorRef, sourceTimebase uintptr, timebaseOut *uintptr) int32
var _cMTimebaseCreateWithSourceTimebaseErr error

func tryCMTimebaseCreateWithSourceTimebase(allocator corefoundation.CFAllocatorRef, sourceTimebase uintptr, timebaseOut *uintptr) (int32, error) {
	if _cMTimebaseCreateWithSourceTimebase == nil {
		return 0, symbolCallError("CMTimebaseCreateWithSourceTimebase", "10.8", _cMTimebaseCreateWithSourceTimebaseErr)
	}
	return _cMTimebaseCreateWithSourceTimebase(allocator, sourceTimebase, timebaseOut), nil
}

// CMTimebaseCreateWithSourceTimebase creates a timebase by using a source timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCreateWithSourceTimebase(allocator:sourceTimebase:timebaseOut:)
func CMTimebaseCreateWithSourceTimebase(allocator corefoundation.CFAllocatorRef, sourceTimebase uintptr, timebaseOut *uintptr) int32 {
	result, callErr := tryCMTimebaseCreateWithSourceTimebase(allocator, sourceTimebase, timebaseOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseGetEffectiveRate func(timebase uintptr) float64
var _cMTimebaseGetEffectiveRateErr error

func tryCMTimebaseGetEffectiveRate(timebase uintptr) (float64, error) {
	if _cMTimebaseGetEffectiveRate == nil {
		return 0.0, symbolCallError("CMTimebaseGetEffectiveRate", "10.8", _cMTimebaseGetEffectiveRateErr)
	}
	return _cMTimebaseGetEffectiveRate(timebase), nil
}

// CMTimebaseGetEffectiveRate returns the effective rate of a timebase, which combines its rate with the rates of all its host timebases.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetEffectiveRate(_:)
func CMTimebaseGetEffectiveRate(timebase uintptr) float64 {
	result, callErr := tryCMTimebaseGetEffectiveRate(timebase)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseGetRate func(timebase uintptr) float64
var _cMTimebaseGetRateErr error

func tryCMTimebaseGetRate(timebase uintptr) (float64, error) {
	if _cMTimebaseGetRate == nil {
		return 0.0, symbolCallError("CMTimebaseGetRate", "10.8", _cMTimebaseGetRateErr)
	}
	return _cMTimebaseGetRate(timebase), nil
}

// CMTimebaseGetRate returns the current rate of a timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetRate(_:)
func CMTimebaseGetRate(timebase uintptr) float64 {
	result, callErr := tryCMTimebaseGetRate(timebase)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseGetTime func(timebase uintptr) CMTime
var _cMTimebaseGetTimeErr error

func tryCMTimebaseGetTime(timebase uintptr) (CMTime, error) {
	if _cMTimebaseGetTime == nil {
		return CMTime{}, symbolCallError("CMTimebaseGetTime", "10.8", _cMTimebaseGetTimeErr)
	}
	return _cMTimebaseGetTime(timebase), nil
}

// CMTimebaseGetTime returns the current time from a timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetTime(_:)
func CMTimebaseGetTime(timebase uintptr) CMTime {
	result, callErr := tryCMTimebaseGetTime(timebase)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseGetTimeAndRate func(timebase uintptr, timeOut *CMTime, rateOut *float64) int32
var _cMTimebaseGetTimeAndRateErr error

func tryCMTimebaseGetTimeAndRate(timebase uintptr, timeOut *CMTime, rateOut *float64) (int32, error) {
	if _cMTimebaseGetTimeAndRate == nil {
		return 0, symbolCallError("CMTimebaseGetTimeAndRate", "10.8", _cMTimebaseGetTimeAndRateErr)
	}
	return _cMTimebaseGetTimeAndRate(timebase, timeOut, rateOut), nil
}

// CMTimebaseGetTimeAndRate returns the current time and rate of a timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetTimeAndRate(_:timeOut:rateOut:)
func CMTimebaseGetTimeAndRate(timebase uintptr, timeOut *CMTime, rateOut *float64) int32 {
	result, callErr := tryCMTimebaseGetTimeAndRate(timebase, timeOut, rateOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseGetTimeWithTimeScale func(timebase uintptr, timescale int32, method CMTimeRoundingMethod) CMTime
var _cMTimebaseGetTimeWithTimeScaleErr error

func tryCMTimebaseGetTimeWithTimeScale(timebase uintptr, timescale int32, method CMTimeRoundingMethod) (CMTime, error) {
	if _cMTimebaseGetTimeWithTimeScale == nil {
		return CMTime{}, symbolCallError("CMTimebaseGetTimeWithTimeScale", "10.8", _cMTimebaseGetTimeWithTimeScaleErr)
	}
	return _cMTimebaseGetTimeWithTimeScale(timebase, timescale, method), nil
}

// CMTimebaseGetTimeWithTimeScale returns the current time from a timebase in the specified timescale.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetTimeWithTimeScale(_:timescale:method:)
func CMTimebaseGetTimeWithTimeScale(timebase uintptr, timescale int32, method CMTimeRoundingMethod) CMTime {
	result, callErr := tryCMTimebaseGetTimeWithTimeScale(timebase, timescale, method)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseGetTypeID func() uint
var _cMTimebaseGetTypeIDErr error

func tryCMTimebaseGetTypeID() (uint, error) {
	if _cMTimebaseGetTypeID == nil {
		return 0, symbolCallError("CMTimebaseGetTypeID", "10.8", _cMTimebaseGetTypeIDErr)
	}
	return _cMTimebaseGetTypeID(), nil
}

// CMTimebaseGetTypeID returns the Core Foundation type identifier that identifies a timebase object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetTypeID()
func CMTimebaseGetTypeID() uint {
	result, callErr := tryCMTimebaseGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseNotificationBarrier func(timebase uintptr) int32
var _cMTimebaseNotificationBarrierErr error

func tryCMTimebaseNotificationBarrier(timebase uintptr) (int32, error) {
	if _cMTimebaseNotificationBarrier == nil {
		return 0, symbolCallError("CMTimebaseNotificationBarrier", "10.8", _cMTimebaseNotificationBarrierErr)
	}
	return _cMTimebaseNotificationBarrier(timebase), nil
}

// CMTimebaseNotificationBarrier requests that the timebase wait until it isn’t posting notifications.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseNotificationBarrier(_:)
func CMTimebaseNotificationBarrier(timebase uintptr) int32 {
	result, callErr := tryCMTimebaseNotificationBarrier(timebase)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseRemoveTimer func(timebase uintptr, timer corefoundation.CFRunLoopTimerRef) int32
var _cMTimebaseRemoveTimerErr error

func tryCMTimebaseRemoveTimer(timebase uintptr, timer corefoundation.CFRunLoopTimerRef) (int32, error) {
	if _cMTimebaseRemoveTimer == nil {
		return 0, symbolCallError("CMTimebaseRemoveTimer", "10.8", _cMTimebaseRemoveTimerErr)
	}
	return _cMTimebaseRemoveTimer(timebase, timer), nil
}

// CMTimebaseRemoveTimer removes the timer from the list of timers the timebase manages.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseRemoveTimer(_:timer:)
func CMTimebaseRemoveTimer(timebase uintptr, timer corefoundation.CFRunLoopTimerRef) int32 {
	result, callErr := tryCMTimebaseRemoveTimer(timebase, timer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseRemoveTimerDispatchSource func(timebase uintptr, timerSource uintptr) int32
var _cMTimebaseRemoveTimerDispatchSourceErr error

func tryCMTimebaseRemoveTimerDispatchSource(timebase uintptr, timerSource dispatch.Source) (int32, error) {
	if _cMTimebaseRemoveTimerDispatchSource == nil {
		return 0, symbolCallError("CMTimebaseRemoveTimerDispatchSource", "10.8", _cMTimebaseRemoveTimerDispatchSourceErr)
	}
	return _cMTimebaseRemoveTimerDispatchSource(timebase, uintptr(timerSource.Handle())), nil
}

// CMTimebaseRemoveTimerDispatchSource removes the timer dispatch source from the list of timers the timebase manages.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseRemoveTimerDispatchSource(_:timerSource:)
func CMTimebaseRemoveTimerDispatchSource(timebase uintptr, timerSource dispatch.Source) int32 {
	result, callErr := tryCMTimebaseRemoveTimerDispatchSource(timebase, timerSource)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseSetAnchorTime func(timebase uintptr, timebaseTime CMTime, immediateSourceTime CMTime) int32
var _cMTimebaseSetAnchorTimeErr error

func tryCMTimebaseSetAnchorTime(timebase uintptr, timebaseTime CMTime, immediateSourceTime CMTime) (int32, error) {
	if _cMTimebaseSetAnchorTime == nil {
		return 0, symbolCallError("CMTimebaseSetAnchorTime", "10.8", _cMTimebaseSetAnchorTimeErr)
	}
	return _cMTimebaseSetAnchorTime(timebase, timebaseTime, immediateSourceTime), nil
}

// CMTimebaseSetAnchorTime sets the time of a timebase at a particular host time.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetAnchorTime(_:timebaseTime:immediateSourceTime:)
func CMTimebaseSetAnchorTime(timebase uintptr, timebaseTime CMTime, immediateSourceTime CMTime) int32 {
	result, callErr := tryCMTimebaseSetAnchorTime(timebase, timebaseTime, immediateSourceTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseSetRate func(timebase uintptr, rate float64) int32
var _cMTimebaseSetRateErr error

func tryCMTimebaseSetRate(timebase uintptr, rate float64) (int32, error) {
	if _cMTimebaseSetRate == nil {
		return 0, symbolCallError("CMTimebaseSetRate", "10.8", _cMTimebaseSetRateErr)
	}
	return _cMTimebaseSetRate(timebase, rate), nil
}

// CMTimebaseSetRate sets the rate of a timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetRate(_:rate:)
func CMTimebaseSetRate(timebase uintptr, rate float64) int32 {
	result, callErr := tryCMTimebaseSetRate(timebase, rate)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseSetRateAndAnchorTime func(timebase uintptr, rate float64, timebaseTime CMTime, immediateSourceTime CMTime) int32
var _cMTimebaseSetRateAndAnchorTimeErr error

func tryCMTimebaseSetRateAndAnchorTime(timebase uintptr, rate float64, timebaseTime CMTime, immediateSourceTime CMTime) (int32, error) {
	if _cMTimebaseSetRateAndAnchorTime == nil {
		return 0, symbolCallError("CMTimebaseSetRateAndAnchorTime", "10.8", _cMTimebaseSetRateAndAnchorTimeErr)
	}
	return _cMTimebaseSetRateAndAnchorTime(timebase, rate, timebaseTime, immediateSourceTime), nil
}

// CMTimebaseSetRateAndAnchorTime sets the time of a timebase at a particular host time, and changes the rate at exactly that time.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetRateAndAnchorTime(_:rate:anchorTime:immediateSourceTime:)
func CMTimebaseSetRateAndAnchorTime(timebase uintptr, rate float64, timebaseTime CMTime, immediateSourceTime CMTime) int32 {
	result, callErr := tryCMTimebaseSetRateAndAnchorTime(timebase, rate, timebaseTime, immediateSourceTime)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseSetSourceClock func(timebase uintptr, newSourceClock uintptr) int32
var _cMTimebaseSetSourceClockErr error

func tryCMTimebaseSetSourceClock(timebase uintptr, newSourceClock uintptr) (int32, error) {
	if _cMTimebaseSetSourceClock == nil {
		return 0, symbolCallError("CMTimebaseSetSourceClock", "10.8", _cMTimebaseSetSourceClockErr)
	}
	return _cMTimebaseSetSourceClock(timebase, newSourceClock), nil
}

// CMTimebaseSetSourceClock sets the source clock of a timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetSourceClock(_:_:)
func CMTimebaseSetSourceClock(timebase uintptr, newSourceClock uintptr) int32 {
	result, callErr := tryCMTimebaseSetSourceClock(timebase, newSourceClock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseSetSourceTimebase func(timebase uintptr, newSourceTimebase uintptr) int32
var _cMTimebaseSetSourceTimebaseErr error

func tryCMTimebaseSetSourceTimebase(timebase uintptr, newSourceTimebase uintptr) (int32, error) {
	if _cMTimebaseSetSourceTimebase == nil {
		return 0, symbolCallError("CMTimebaseSetSourceTimebase", "10.8", _cMTimebaseSetSourceTimebaseErr)
	}
	return _cMTimebaseSetSourceTimebase(timebase, newSourceTimebase), nil
}

// CMTimebaseSetSourceTimebase sets the source timebase of a timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetSourceTimebase(_:_:)
func CMTimebaseSetSourceTimebase(timebase uintptr, newSourceTimebase uintptr) int32 {
	result, callErr := tryCMTimebaseSetSourceTimebase(timebase, newSourceTimebase)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseSetTime func(timebase uintptr, time CMTime) int32
var _cMTimebaseSetTimeErr error

func tryCMTimebaseSetTime(timebase uintptr, time CMTime) (int32, error) {
	if _cMTimebaseSetTime == nil {
		return 0, symbolCallError("CMTimebaseSetTime", "10.8", _cMTimebaseSetTimeErr)
	}
	return _cMTimebaseSetTime(timebase, time), nil
}

// CMTimebaseSetTime sets the current time of a timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTime(_:time:)
func CMTimebaseSetTime(timebase uintptr, time CMTime) int32 {
	result, callErr := tryCMTimebaseSetTime(timebase, time)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseSetTimerDispatchSourceNextFireTime func(timebase uintptr, timerSource uintptr, fireTime CMTime, flags uint32) int32
var _cMTimebaseSetTimerDispatchSourceNextFireTimeErr error

func tryCMTimebaseSetTimerDispatchSourceNextFireTime(timebase uintptr, timerSource dispatch.Source, fireTime CMTime, flags uint32) (int32, error) {
	if _cMTimebaseSetTimerDispatchSourceNextFireTime == nil {
		return 0, symbolCallError("CMTimebaseSetTimerDispatchSourceNextFireTime", "10.8", _cMTimebaseSetTimerDispatchSourceNextFireTimeErr)
	}
	return _cMTimebaseSetTimerDispatchSourceNextFireTime(timebase, uintptr(timerSource.Handle()), fireTime, flags), nil
}

// CMTimebaseSetTimerDispatchSourceNextFireTime sets the time on the timebase’s timeline at which the timer dispatch source should fire next.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTimerDispatchSourceNextFireTime(_:timerSource:fireTime:flags:)
func CMTimebaseSetTimerDispatchSourceNextFireTime(timebase uintptr, timerSource dispatch.Source, fireTime CMTime, flags uint32) int32 {
	result, callErr := tryCMTimebaseSetTimerDispatchSourceNextFireTime(timebase, timerSource, fireTime, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseSetTimerDispatchSourceToFireImmediately func(timebase uintptr, timerSource uintptr) int32
var _cMTimebaseSetTimerDispatchSourceToFireImmediatelyErr error

func tryCMTimebaseSetTimerDispatchSourceToFireImmediately(timebase uintptr, timerSource dispatch.Source) (int32, error) {
	if _cMTimebaseSetTimerDispatchSourceToFireImmediately == nil {
		return 0, symbolCallError("CMTimebaseSetTimerDispatchSourceToFireImmediately", "10.8", _cMTimebaseSetTimerDispatchSourceToFireImmediatelyErr)
	}
	return _cMTimebaseSetTimerDispatchSourceToFireImmediately(timebase, uintptr(timerSource.Handle())), nil
}

// CMTimebaseSetTimerDispatchSourceToFireImmediately sets the timer dispatch source to fire immediately once, overriding any previous timer call.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTimerDispatchSourceToFireImmediately(_:timerSource:)
func CMTimebaseSetTimerDispatchSourceToFireImmediately(timebase uintptr, timerSource dispatch.Source) int32 {
	result, callErr := tryCMTimebaseSetTimerDispatchSourceToFireImmediately(timebase, timerSource)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseSetTimerNextFireTime func(timebase uintptr, timer corefoundation.CFRunLoopTimerRef, fireTime CMTime, flags uint32) int32
var _cMTimebaseSetTimerNextFireTimeErr error

func tryCMTimebaseSetTimerNextFireTime(timebase uintptr, timer corefoundation.CFRunLoopTimerRef, fireTime CMTime, flags uint32) (int32, error) {
	if _cMTimebaseSetTimerNextFireTime == nil {
		return 0, symbolCallError("CMTimebaseSetTimerNextFireTime", "10.8", _cMTimebaseSetTimerNextFireTimeErr)
	}
	return _cMTimebaseSetTimerNextFireTime(timebase, timer, fireTime, flags), nil
}

// CMTimebaseSetTimerNextFireTime sets the time on the timebase’s timeline at which the timer should fire next.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTimerNextFireTime(_:timer:fireTime:flags:)
func CMTimebaseSetTimerNextFireTime(timebase uintptr, timer corefoundation.CFRunLoopTimerRef, fireTime CMTime, flags uint32) int32 {
	result, callErr := tryCMTimebaseSetTimerNextFireTime(timebase, timer, fireTime, flags)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTimebaseSetTimerToFireImmediately func(timebase uintptr, timer corefoundation.CFRunLoopTimerRef) int32
var _cMTimebaseSetTimerToFireImmediatelyErr error

func tryCMTimebaseSetTimerToFireImmediately(timebase uintptr, timer corefoundation.CFRunLoopTimerRef) (int32, error) {
	if _cMTimebaseSetTimerToFireImmediately == nil {
		return 0, symbolCallError("CMTimebaseSetTimerToFireImmediately", "10.8", _cMTimebaseSetTimerToFireImmediatelyErr)
	}
	return _cMTimebaseSetTimerToFireImmediately(timebase, timer), nil
}

// CMTimebaseSetTimerToFireImmediately sets the timer to fire immediately once, overriding any previous timer calls.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTimerToFireImmediately(_:timer:)
func CMTimebaseSetTimerToFireImmediately(timebase uintptr, timer corefoundation.CFRunLoopTimerRef) int32 {
	result, callErr := tryCMTimebaseSetTimerToFireImmediately(timebase, timer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, videoFormatDescription uintptr, stringEncoding uint32, flavor CMImageDescriptionFlavor, blockBufferOut *uintptr) int32
var _cMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBufferErr error

func tryCMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, videoFormatDescription uintptr, stringEncoding uint32, flavor CMImageDescriptionFlavor, blockBufferOut *uintptr) (int32, error) {
	if _cMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer == nil {
		return 0, symbolCallError("CMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer", "10.10", _cMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBufferErr)
	}
	return _cMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer(allocator, videoFormatDescription, stringEncoding, flavor, blockBufferOut), nil
}

// CMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer copies the contents of a video format description to a buffer in big-endian byte ordering.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer(allocator:videoFormatDescription:stringEncoding:flavor:blockBufferOut:)
func CMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, videoFormatDescription uintptr, stringEncoding uint32, flavor CMImageDescriptionFlavor, blockBufferOut *uintptr) int32 {
	result, callErr := tryCMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer(allocator, videoFormatDescription, stringEncoding, flavor, blockBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMVideoFormatDescriptionCopyTagCollectionArray func(formatDescription uintptr, tagCollectionsOut *corefoundation.CFArrayRef) int32
var _cMVideoFormatDescriptionCopyTagCollectionArrayErr error

func tryCMVideoFormatDescriptionCopyTagCollectionArray(formatDescription uintptr, tagCollectionsOut *corefoundation.CFArrayRef) (int32, error) {
	if _cMVideoFormatDescriptionCopyTagCollectionArray == nil {
		return 0, symbolCallError("CMVideoFormatDescriptionCopyTagCollectionArray", "14.0", _cMVideoFormatDescriptionCopyTagCollectionArrayErr)
	}
	return _cMVideoFormatDescriptionCopyTagCollectionArray(formatDescription, tagCollectionsOut), nil
}

// CMVideoFormatDescriptionCopyTagCollectionArray copies the multi-image encoding properties as an array of CMTagCollections.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCopyTagCollectionArray
func CMVideoFormatDescriptionCopyTagCollectionArray(formatDescription uintptr, tagCollectionsOut *corefoundation.CFArrayRef) int32 {
	result, callErr := tryCMVideoFormatDescriptionCopyTagCollectionArray(formatDescription, tagCollectionsOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMVideoFormatDescriptionCreate func(allocator corefoundation.CFAllocatorRef, codecType CMVideoCodecType, width int32, height int32, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *uintptr) int32
var _cMVideoFormatDescriptionCreateErr error

func tryCMVideoFormatDescriptionCreate(allocator corefoundation.CFAllocatorRef, codecType CMVideoCodecType, width int32, height int32, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *uintptr) (int32, error) {
	if _cMVideoFormatDescriptionCreate == nil {
		return 0, symbolCallError("CMVideoFormatDescriptionCreate", "10.7", _cMVideoFormatDescriptionCreateErr)
	}
	return _cMVideoFormatDescriptionCreate(allocator, codecType, width, height, extensions, formatDescriptionOut), nil
}

// CMVideoFormatDescriptionCreate creates a format description for a video media stream.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCreate(allocator:codecType:width:height:extensions:formatDescriptionOut:)
func CMVideoFormatDescriptionCreate(allocator corefoundation.CFAllocatorRef, codecType CMVideoCodecType, width int32, height int32, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *uintptr) int32 {
	result, callErr := tryCMVideoFormatDescriptionCreate(allocator, codecType, width, height, extensions, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMVideoFormatDescriptionCreateForImageBuffer func(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, formatDescriptionOut *uintptr) int32
var _cMVideoFormatDescriptionCreateForImageBufferErr error

func tryCMVideoFormatDescriptionCreateForImageBuffer(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, formatDescriptionOut *uintptr) (int32, error) {
	if _cMVideoFormatDescriptionCreateForImageBuffer == nil {
		return 0, symbolCallError("CMVideoFormatDescriptionCreateForImageBuffer", "10.7", _cMVideoFormatDescriptionCreateForImageBufferErr)
	}
	return _cMVideoFormatDescriptionCreateForImageBuffer(allocator, imageBuffer, formatDescriptionOut), nil
}

// CMVideoFormatDescriptionCreateForImageBuffer creates a format description for a video media stream by using an image buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCreateForImageBuffer(allocator:imageBuffer:formatDescriptionOut:)
func CMVideoFormatDescriptionCreateForImageBuffer(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, formatDescriptionOut *uintptr) int32 {
	result, callErr := tryCMVideoFormatDescriptionCreateForImageBuffer(allocator, imageBuffer, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, imageDescriptionBlockBuffer uintptr, stringEncoding uint32, flavor CMImageDescriptionFlavor, formatDescriptionOut *uintptr) int32
var _cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBufferErr error

func tryCMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, imageDescriptionBlockBuffer uintptr, stringEncoding uint32, flavor CMImageDescriptionFlavor, formatDescriptionOut *uintptr) (int32, error) {
	if _cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer == nil {
		return 0, symbolCallError("CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer", "10.10", _cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBufferErr)
	}
	return _cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer(allocator, imageDescriptionBlockBuffer, stringEncoding, flavor, formatDescriptionOut), nil
}

// CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer creates a video format description from a big-endian image description inside a buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer(allocator:bigEndianImageDescriptionBlockBuffer:stringEncoding:flavor:formatDescriptionOut:)
func CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, imageDescriptionBlockBuffer uintptr, stringEncoding uint32, flavor CMImageDescriptionFlavor, formatDescriptionOut *uintptr) int32 {
	result, callErr := tryCMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer(allocator, imageDescriptionBlockBuffer, stringEncoding, flavor, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData func(allocator corefoundation.CFAllocatorRef, imageDescriptionData *uint8, size uintptr, stringEncoding uint32, flavor CMImageDescriptionFlavor, formatDescriptionOut *uintptr) int32
var _cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionDataErr error

func tryCMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData(allocator corefoundation.CFAllocatorRef, imageDescriptionData *uint8, size uintptr, stringEncoding uint32, flavor CMImageDescriptionFlavor, formatDescriptionOut *uintptr) (int32, error) {
	if _cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData == nil {
		return 0, symbolCallError("CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData", "10.10", _cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionDataErr)
	}
	return _cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData(allocator, imageDescriptionData, size, stringEncoding, flavor, formatDescriptionOut), nil
}

// CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData creates a video format description from a big-endian image description structure.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData(allocator:bigEndianImageDescriptionData:size:stringEncoding:flavor:formatDescriptionOut:)
func CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData(allocator corefoundation.CFAllocatorRef, imageDescriptionData *uint8, size uintptr, stringEncoding uint32, flavor CMImageDescriptionFlavor, formatDescriptionOut *uintptr) int32 {
	result, callErr := tryCMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData(allocator, imageDescriptionData, size, stringEncoding, flavor, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMVideoFormatDescriptionCreateFromH264ParameterSets func(allocator corefoundation.CFAllocatorRef, parameterSetCount uintptr, parameterSetPointers *uint8, parameterSetSizes *uintptr, NALUnitHeaderLength int, formatDescriptionOut *uintptr) int32
var _cMVideoFormatDescriptionCreateFromH264ParameterSetsErr error

func tryCMVideoFormatDescriptionCreateFromH264ParameterSets(allocator corefoundation.CFAllocatorRef, parameterSetCount uintptr, parameterSetPointers *uint8, parameterSetSizes *uintptr, NALUnitHeaderLength int, formatDescriptionOut *uintptr) (int32, error) {
	if _cMVideoFormatDescriptionCreateFromH264ParameterSets == nil {
		return 0, symbolCallError("CMVideoFormatDescriptionCreateFromH264ParameterSets", "10.9", _cMVideoFormatDescriptionCreateFromH264ParameterSetsErr)
	}
	return _cMVideoFormatDescriptionCreateFromH264ParameterSets(allocator, parameterSetCount, parameterSetPointers, parameterSetSizes, NALUnitHeaderLength, formatDescriptionOut), nil
}

// CMVideoFormatDescriptionCreateFromH264ParameterSets creates a format description for a video media stream that the parameter set describes.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCreateFromH264ParameterSets(allocator:parameterSetCount:parameterSetPointers:parameterSetSizes:nalUnitHeaderLength:formatDescriptionOut:)
func CMVideoFormatDescriptionCreateFromH264ParameterSets(allocator corefoundation.CFAllocatorRef, parameterSetCount uintptr, parameterSetPointers *uint8, parameterSetSizes *uintptr, NALUnitHeaderLength int, formatDescriptionOut *uintptr) int32 {
	result, callErr := tryCMVideoFormatDescriptionCreateFromH264ParameterSets(allocator, parameterSetCount, parameterSetPointers, parameterSetSizes, NALUnitHeaderLength, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMVideoFormatDescriptionCreateFromHEVCParameterSets func(allocator corefoundation.CFAllocatorRef, parameterSetCount uintptr, parameterSetPointers *uint8, parameterSetSizes *uintptr, NALUnitHeaderLength int, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *uintptr) int32
var _cMVideoFormatDescriptionCreateFromHEVCParameterSetsErr error

func tryCMVideoFormatDescriptionCreateFromHEVCParameterSets(allocator corefoundation.CFAllocatorRef, parameterSetCount uintptr, parameterSetPointers *uint8, parameterSetSizes *uintptr, NALUnitHeaderLength int, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *uintptr) (int32, error) {
	if _cMVideoFormatDescriptionCreateFromHEVCParameterSets == nil {
		return 0, symbolCallError("CMVideoFormatDescriptionCreateFromHEVCParameterSets", "10.13", _cMVideoFormatDescriptionCreateFromHEVCParameterSetsErr)
	}
	return _cMVideoFormatDescriptionCreateFromHEVCParameterSets(allocator, parameterSetCount, parameterSetPointers, parameterSetSizes, NALUnitHeaderLength, extensions, formatDescriptionOut), nil
}

// CMVideoFormatDescriptionCreateFromHEVCParameterSets creates a format description for a video media stream using HEVC (H.265) parameter set NAL units.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCreateFromHEVCParameterSets(allocator:parameterSetCount:parameterSetPointers:parameterSetSizes:nalUnitHeaderLength:extensions:formatDescriptionOut:)
func CMVideoFormatDescriptionCreateFromHEVCParameterSets(allocator corefoundation.CFAllocatorRef, parameterSetCount uintptr, parameterSetPointers *uint8, parameterSetSizes *uintptr, NALUnitHeaderLength int, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *uintptr) int32 {
	result, callErr := tryCMVideoFormatDescriptionCreateFromHEVCParameterSets(allocator, parameterSetCount, parameterSetPointers, parameterSetSizes, NALUnitHeaderLength, extensions, formatDescriptionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMVideoFormatDescriptionGetCleanAperture func(videoDesc uintptr, originIsAtTopLeft bool) corefoundation.CGRect
var _cMVideoFormatDescriptionGetCleanApertureErr error

func tryCMVideoFormatDescriptionGetCleanAperture(videoDesc uintptr, originIsAtTopLeft bool) (corefoundation.CGRect, error) {
	if _cMVideoFormatDescriptionGetCleanAperture == nil {
		return corefoundation.CGRect{}, symbolCallError("CMVideoFormatDescriptionGetCleanAperture", "10.7", _cMVideoFormatDescriptionGetCleanApertureErr)
	}
	return _cMVideoFormatDescriptionGetCleanAperture(videoDesc, originIsAtTopLeft), nil
}

// CMVideoFormatDescriptionGetCleanAperture returns a rectangle that defines the portion of the encoded pixel dimensions that represent the image data that’s valid for displaying.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetCleanAperture(_:originIsAtTopLeft:)
func CMVideoFormatDescriptionGetCleanAperture(videoDesc uintptr, originIsAtTopLeft bool) corefoundation.CGRect {
	result, callErr := tryCMVideoFormatDescriptionGetCleanAperture(videoDesc, originIsAtTopLeft)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMVideoFormatDescriptionGetDimensions func(videoDesc uintptr) CMVideoDimensions
var _cMVideoFormatDescriptionGetDimensionsErr error

func tryCMVideoFormatDescriptionGetDimensions(videoDesc uintptr) (CMVideoDimensions, error) {
	if _cMVideoFormatDescriptionGetDimensions == nil {
		return CMVideoDimensions{}, symbolCallError("CMVideoFormatDescriptionGetDimensions", "10.7", _cMVideoFormatDescriptionGetDimensionsErr)
	}
	return _cMVideoFormatDescriptionGetDimensions(videoDesc), nil
}

// CMVideoFormatDescriptionGetDimensions returns the video dimensions, in encoded pixels.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetDimensions(_:)
func CMVideoFormatDescriptionGetDimensions(videoDesc uintptr) CMVideoDimensions {
	result, callErr := tryCMVideoFormatDescriptionGetDimensions(videoDesc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers func() corefoundation.CFArrayRef
var _cMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffersErr error

func tryCMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers() (corefoundation.CFArrayRef, error) {
	if _cMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers == nil {
		return 0, symbolCallError("CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers", "10.7", _cMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffersErr)
	}
	return _cMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers(), nil
}

// CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers returns an array of keys that you use for video format description extensions, image buffer attachments, and attributes.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers()
func CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers() corefoundation.CFArrayRef {
	result, callErr := tryCMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMVideoFormatDescriptionGetH264ParameterSetAtIndex func(videoDesc uintptr, parameterSetIndex uintptr, parameterSetPointerOut *uint8, parameterSetSizeOut *uintptr, parameterSetCountOut *uintptr, NALUnitHeaderLengthOut *int) int32
var _cMVideoFormatDescriptionGetH264ParameterSetAtIndexErr error

func tryCMVideoFormatDescriptionGetH264ParameterSetAtIndex(videoDesc uintptr, parameterSetIndex uintptr, parameterSetPointerOut *uint8, parameterSetSizeOut *uintptr, parameterSetCountOut *uintptr, NALUnitHeaderLengthOut []int) (int32, error) {
	if _cMVideoFormatDescriptionGetH264ParameterSetAtIndex == nil {
		return 0, symbolCallError("CMVideoFormatDescriptionGetH264ParameterSetAtIndex", "10.9", _cMVideoFormatDescriptionGetH264ParameterSetAtIndexErr)
	}
	return _cMVideoFormatDescriptionGetH264ParameterSetAtIndex(videoDesc, parameterSetIndex, parameterSetPointerOut, parameterSetSizeOut, parameterSetCountOut, unsafe.SliceData(NALUnitHeaderLengthOut)), nil
}

// CMVideoFormatDescriptionGetH264ParameterSetAtIndex returns a parameter set that an H.264 format description contains.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetH264ParameterSetAtIndex(_:parameterSetIndex:parameterSetPointerOut:parameterSetSizeOut:parameterSetCountOut:nalUnitHeaderLengthOut:)
func CMVideoFormatDescriptionGetH264ParameterSetAtIndex(videoDesc uintptr, parameterSetIndex uintptr, parameterSetPointerOut *uint8, parameterSetSizeOut *uintptr, parameterSetCountOut *uintptr, NALUnitHeaderLengthOut []int) int32 {
	result, callErr := tryCMVideoFormatDescriptionGetH264ParameterSetAtIndex(videoDesc, parameterSetIndex, parameterSetPointerOut, parameterSetSizeOut, parameterSetCountOut, NALUnitHeaderLengthOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMVideoFormatDescriptionGetHEVCParameterSetAtIndex func(videoDesc uintptr, parameterSetIndex uintptr, parameterSetPointerOut *uint8, parameterSetSizeOut *uintptr, parameterSetCountOut *uintptr, NALUnitHeaderLengthOut *int) int32
var _cMVideoFormatDescriptionGetHEVCParameterSetAtIndexErr error

func tryCMVideoFormatDescriptionGetHEVCParameterSetAtIndex(videoDesc uintptr, parameterSetIndex uintptr, parameterSetPointerOut *uint8, parameterSetSizeOut *uintptr, parameterSetCountOut *uintptr, NALUnitHeaderLengthOut []int) (int32, error) {
	if _cMVideoFormatDescriptionGetHEVCParameterSetAtIndex == nil {
		return 0, symbolCallError("CMVideoFormatDescriptionGetHEVCParameterSetAtIndex", "10.13", _cMVideoFormatDescriptionGetHEVCParameterSetAtIndexErr)
	}
	return _cMVideoFormatDescriptionGetHEVCParameterSetAtIndex(videoDesc, parameterSetIndex, parameterSetPointerOut, parameterSetSizeOut, parameterSetCountOut, unsafe.SliceData(NALUnitHeaderLengthOut)), nil
}

// CMVideoFormatDescriptionGetHEVCParameterSetAtIndex returns a parameter set contained in an HEVC (H.265) format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetHEVCParameterSetAtIndex(_:parameterSetIndex:parameterSetPointerOut:parameterSetSizeOut:parameterSetCountOut:nalUnitHeaderLengthOut:)
func CMVideoFormatDescriptionGetHEVCParameterSetAtIndex(videoDesc uintptr, parameterSetIndex uintptr, parameterSetPointerOut *uint8, parameterSetSizeOut *uintptr, parameterSetCountOut *uintptr, NALUnitHeaderLengthOut []int) int32 {
	result, callErr := tryCMVideoFormatDescriptionGetHEVCParameterSetAtIndex(videoDesc, parameterSetIndex, parameterSetPointerOut, parameterSetSizeOut, parameterSetCountOut, NALUnitHeaderLengthOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMVideoFormatDescriptionGetPresentationDimensions func(videoDesc uintptr, usePixelAspectRatio bool, useCleanAperture bool) corefoundation.CGSize
var _cMVideoFormatDescriptionGetPresentationDimensionsErr error

func tryCMVideoFormatDescriptionGetPresentationDimensions(videoDesc uintptr, usePixelAspectRatio bool, useCleanAperture bool) (corefoundation.CGSize, error) {
	if _cMVideoFormatDescriptionGetPresentationDimensions == nil {
		return corefoundation.CGSize{}, symbolCallError("CMVideoFormatDescriptionGetPresentationDimensions", "10.7", _cMVideoFormatDescriptionGetPresentationDimensionsErr)
	}
	return _cMVideoFormatDescriptionGetPresentationDimensions(videoDesc, usePixelAspectRatio, useCleanAperture), nil
}

// CMVideoFormatDescriptionGetPresentationDimensions returns the dimensions after taking the pixel aspect ratio and clean aperture into account.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetPresentationDimensions(_:usePixelAspectRatio:useCleanAperture:)
func CMVideoFormatDescriptionGetPresentationDimensions(videoDesc uintptr, usePixelAspectRatio bool, useCleanAperture bool) corefoundation.CGSize {
	result, callErr := tryCMVideoFormatDescriptionGetPresentationDimensions(videoDesc, usePixelAspectRatio, useCleanAperture)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMVideoFormatDescriptionMatchesImageBuffer func(desc uintptr, imageBuffer corevideo.CVImageBufferRef) bool
var _cMVideoFormatDescriptionMatchesImageBufferErr error

func tryCMVideoFormatDescriptionMatchesImageBuffer(desc uintptr, imageBuffer corevideo.CVImageBufferRef) (bool, error) {
	if _cMVideoFormatDescriptionMatchesImageBuffer == nil {
		return false, symbolCallError("CMVideoFormatDescriptionMatchesImageBuffer", "10.7", _cMVideoFormatDescriptionMatchesImageBufferErr)
	}
	return _cMVideoFormatDescriptionMatchesImageBuffer(desc, imageBuffer), nil
}

// CMVideoFormatDescriptionMatchesImageBuffer returns a Boolean value that indicates whether a format description matches an image buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionMatchesImageBuffer(_:imageBuffer:)
func CMVideoFormatDescriptionMatchesImageBuffer(desc uintptr, imageBuffer corevideo.CVImageBufferRef) bool {
	result, callErr := tryCMVideoFormatDescriptionMatchesImageBuffer(desc, imageBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_cMAudioDeviceClockCreate, &_cMAudioDeviceClockCreateErr, frameworkHandle, "CMAudioDeviceClockCreate", "10.8")
	registerFunc(&_cMAudioDeviceClockCreateFromAudioDeviceID, &_cMAudioDeviceClockCreateFromAudioDeviceIDErr, frameworkHandle, "CMAudioDeviceClockCreateFromAudioDeviceID", "10.8")
	registerFunc(&_cMAudioDeviceClockGetAudioDevice, &_cMAudioDeviceClockGetAudioDeviceErr, frameworkHandle, "CMAudioDeviceClockGetAudioDevice", "10.8")
	registerFunc(&_cMAudioDeviceClockSetAudioDeviceID, &_cMAudioDeviceClockSetAudioDeviceIDErr, frameworkHandle, "CMAudioDeviceClockSetAudioDeviceID", "10.8")
	registerFunc(&_cMAudioDeviceClockSetAudioDeviceUID, &_cMAudioDeviceClockSetAudioDeviceUIDErr, frameworkHandle, "CMAudioDeviceClockSetAudioDeviceUID", "10.8")
	registerFunc(&_cMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer, &_cMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBufferErr, frameworkHandle, "CMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer", "10.10")
	registerFunc(&_cMAudioFormatDescriptionCreate, &_cMAudioFormatDescriptionCreateErr, frameworkHandle, "CMAudioFormatDescriptionCreate", "10.7")
	registerFunc(&_cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer, &_cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBufferErr, frameworkHandle, "CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer", "10.10")
	registerFunc(&_cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData, &_cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionDataErr, frameworkHandle, "CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData", "10.10")
	registerFunc(&_cMAudioFormatDescriptionCreateSummary, &_cMAudioFormatDescriptionCreateSummaryErr, frameworkHandle, "CMAudioFormatDescriptionCreateSummary", "10.7")
	registerFunc(&_cMAudioFormatDescriptionEqual, &_cMAudioFormatDescriptionEqualErr, frameworkHandle, "CMAudioFormatDescriptionEqual", "10.7")
	registerFunc(&_cMAudioFormatDescriptionGetChannelLayout, &_cMAudioFormatDescriptionGetChannelLayoutErr, frameworkHandle, "CMAudioFormatDescriptionGetChannelLayout", "10.7")
	registerFunc(&_cMAudioFormatDescriptionGetFormatList, &_cMAudioFormatDescriptionGetFormatListErr, frameworkHandle, "CMAudioFormatDescriptionGetFormatList", "10.7")
	registerFunc(&_cMAudioFormatDescriptionGetMagicCookie, &_cMAudioFormatDescriptionGetMagicCookieErr, frameworkHandle, "CMAudioFormatDescriptionGetMagicCookie", "10.7")
	registerFunc(&_cMAudioFormatDescriptionGetMostCompatibleFormat, &_cMAudioFormatDescriptionGetMostCompatibleFormatErr, frameworkHandle, "CMAudioFormatDescriptionGetMostCompatibleFormat", "10.7")
	registerFunc(&_cMAudioFormatDescriptionGetRichestDecodableFormat, &_cMAudioFormatDescriptionGetRichestDecodableFormatErr, frameworkHandle, "CMAudioFormatDescriptionGetRichestDecodableFormat", "10.7")
	registerFunc(&_cMAudioFormatDescriptionGetStreamBasicDescription, &_cMAudioFormatDescriptionGetStreamBasicDescriptionErr, frameworkHandle, "CMAudioFormatDescriptionGetStreamBasicDescription", "10.7")
	registerFunc(&_cMAudioSampleBufferCreateReadyWithPacketDescriptions, &_cMAudioSampleBufferCreateReadyWithPacketDescriptionsErr, frameworkHandle, "CMAudioSampleBufferCreateReadyWithPacketDescriptions", "10.10")
	registerFunc(&_cMAudioSampleBufferCreateWithPacketDescriptions, &_cMAudioSampleBufferCreateWithPacketDescriptionsErr, frameworkHandle, "CMAudioSampleBufferCreateWithPacketDescriptions", "10.7")
	registerFunc(&_cMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler, &_cMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandlerErr, frameworkHandle, "CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler", "10.14.4")
	registerFunc(&_cMBlockBufferAccessDataBytes, &_cMBlockBufferAccessDataBytesErr, frameworkHandle, "CMBlockBufferAccessDataBytes", "10.7")
	registerFunc(&_cMBlockBufferAppendBufferReference, &_cMBlockBufferAppendBufferReferenceErr, frameworkHandle, "CMBlockBufferAppendBufferReference", "10.7")
	registerFunc(&_cMBlockBufferAppendMemoryBlock, &_cMBlockBufferAppendMemoryBlockErr, frameworkHandle, "CMBlockBufferAppendMemoryBlock", "10.7")
	registerFunc(&_cMBlockBufferAssureBlockMemory, &_cMBlockBufferAssureBlockMemoryErr, frameworkHandle, "CMBlockBufferAssureBlockMemory", "10.7")
	registerFunc(&_cMBlockBufferCopyDataBytes, &_cMBlockBufferCopyDataBytesErr, frameworkHandle, "CMBlockBufferCopyDataBytes", "10.7")
	registerFunc(&_cMBlockBufferCreateContiguous, &_cMBlockBufferCreateContiguousErr, frameworkHandle, "CMBlockBufferCreateContiguous", "10.7")
	registerFunc(&_cMBlockBufferCreateEmpty, &_cMBlockBufferCreateEmptyErr, frameworkHandle, "CMBlockBufferCreateEmpty", "10.7")
	registerFunc(&_cMBlockBufferCreateWithBufferReference, &_cMBlockBufferCreateWithBufferReferenceErr, frameworkHandle, "CMBlockBufferCreateWithBufferReference", "10.7")
	registerFunc(&_cMBlockBufferCreateWithMemoryBlock, &_cMBlockBufferCreateWithMemoryBlockErr, frameworkHandle, "CMBlockBufferCreateWithMemoryBlock", "10.7")
	registerFunc(&_cMBlockBufferFillDataBytes, &_cMBlockBufferFillDataBytesErr, frameworkHandle, "CMBlockBufferFillDataBytes", "10.7")
	registerFunc(&_cMBlockBufferGetDataLength, &_cMBlockBufferGetDataLengthErr, frameworkHandle, "CMBlockBufferGetDataLength", "10.7")
	registerFunc(&_cMBlockBufferGetDataPointer, &_cMBlockBufferGetDataPointerErr, frameworkHandle, "CMBlockBufferGetDataPointer", "10.7")
	registerFunc(&_cMBlockBufferGetTypeID, &_cMBlockBufferGetTypeIDErr, frameworkHandle, "CMBlockBufferGetTypeID", "10.7")
	registerFunc(&_cMBlockBufferIsEmpty, &_cMBlockBufferIsEmptyErr, frameworkHandle, "CMBlockBufferIsEmpty", "10.7")
	registerFunc(&_cMBlockBufferIsRangeContiguous, &_cMBlockBufferIsRangeContiguousErr, frameworkHandle, "CMBlockBufferIsRangeContiguous", "10.7")
	registerFunc(&_cMBlockBufferReplaceDataBytes, &_cMBlockBufferReplaceDataBytesErr, frameworkHandle, "CMBlockBufferReplaceDataBytes", "10.7")
	registerFunc(&_cMBufferQueueCallForEachBuffer, &_cMBufferQueueCallForEachBufferErr, frameworkHandle, "CMBufferQueueCallForEachBuffer", "10.7")
	registerFunc(&_cMBufferQueueContainsEndOfData, &_cMBufferQueueContainsEndOfDataErr, frameworkHandle, "CMBufferQueueContainsEndOfData", "10.7")
	registerFunc(&_cMBufferQueueCopyHead, &_cMBufferQueueCopyHeadErr, frameworkHandle, "CMBufferQueueCopyHead", "14.0")
	registerFunc(&_cMBufferQueueCreate, &_cMBufferQueueCreateErr, frameworkHandle, "CMBufferQueueCreate", "10.7")
	registerFunc(&_cMBufferQueueCreateWithHandlers, &_cMBufferQueueCreateWithHandlersErr, frameworkHandle, "CMBufferQueueCreateWithHandlers", "10.14.4")
	registerFunc(&_cMBufferQueueDequeueAndRetain, &_cMBufferQueueDequeueAndRetainErr, frameworkHandle, "CMBufferQueueDequeueAndRetain", "10.7")
	registerFunc(&_cMBufferQueueDequeueIfDataReadyAndRetain, &_cMBufferQueueDequeueIfDataReadyAndRetainErr, frameworkHandle, "CMBufferQueueDequeueIfDataReadyAndRetain", "10.7")
	registerFunc(&_cMBufferQueueEnqueue, &_cMBufferQueueEnqueueErr, frameworkHandle, "CMBufferQueueEnqueue", "10.7")
	registerFunc(&_cMBufferQueueGetBufferCount, &_cMBufferQueueGetBufferCountErr, frameworkHandle, "CMBufferQueueGetBufferCount", "10.7")
	registerFunc(&_cMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS, &_cMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTSErr, frameworkHandle, "CMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS", "10.7")
	registerFunc(&_cMBufferQueueGetCallbacksForUnsortedSampleBuffers, &_cMBufferQueueGetCallbacksForUnsortedSampleBuffersErr, frameworkHandle, "CMBufferQueueGetCallbacksForUnsortedSampleBuffers", "10.7")
	registerFunc(&_cMBufferQueueGetDuration, &_cMBufferQueueGetDurationErr, frameworkHandle, "CMBufferQueueGetDuration", "10.7")
	registerFunc(&_cMBufferQueueGetEndPresentationTimeStamp, &_cMBufferQueueGetEndPresentationTimeStampErr, frameworkHandle, "CMBufferQueueGetEndPresentationTimeStamp", "10.7")
	registerFunc(&_cMBufferQueueGetFirstDecodeTimeStamp, &_cMBufferQueueGetFirstDecodeTimeStampErr, frameworkHandle, "CMBufferQueueGetFirstDecodeTimeStamp", "10.7")
	registerFunc(&_cMBufferQueueGetFirstPresentationTimeStamp, &_cMBufferQueueGetFirstPresentationTimeStampErr, frameworkHandle, "CMBufferQueueGetFirstPresentationTimeStamp", "10.7")
	registerFunc(&_cMBufferQueueGetHead, &_cMBufferQueueGetHeadErr, frameworkHandle, "CMBufferQueueGetHead", "10.7")
	registerFunc(&_cMBufferQueueGetMaxPresentationTimeStamp, &_cMBufferQueueGetMaxPresentationTimeStampErr, frameworkHandle, "CMBufferQueueGetMaxPresentationTimeStamp", "10.7")
	registerFunc(&_cMBufferQueueGetMinDecodeTimeStamp, &_cMBufferQueueGetMinDecodeTimeStampErr, frameworkHandle, "CMBufferQueueGetMinDecodeTimeStamp", "10.7")
	registerFunc(&_cMBufferQueueGetMinPresentationTimeStamp, &_cMBufferQueueGetMinPresentationTimeStampErr, frameworkHandle, "CMBufferQueueGetMinPresentationTimeStamp", "10.7")
	registerFunc(&_cMBufferQueueGetTotalSize, &_cMBufferQueueGetTotalSizeErr, frameworkHandle, "CMBufferQueueGetTotalSize", "10.10")
	registerFunc(&_cMBufferQueueGetTypeID, &_cMBufferQueueGetTypeIDErr, frameworkHandle, "CMBufferQueueGetTypeID", "10.7")
	registerFunc(&_cMBufferQueueInstallTrigger, &_cMBufferQueueInstallTriggerErr, frameworkHandle, "CMBufferQueueInstallTrigger", "10.7")
	registerFunc(&_cMBufferQueueInstallTriggerHandler, &_cMBufferQueueInstallTriggerHandlerErr, frameworkHandle, "CMBufferQueueInstallTriggerHandler", "10.14.4")
	registerFunc(&_cMBufferQueueInstallTriggerHandlerWithIntegerThreshold, &_cMBufferQueueInstallTriggerHandlerWithIntegerThresholdErr, frameworkHandle, "CMBufferQueueInstallTriggerHandlerWithIntegerThreshold", "10.14.4")
	registerFunc(&_cMBufferQueueInstallTriggerWithIntegerThreshold, &_cMBufferQueueInstallTriggerWithIntegerThresholdErr, frameworkHandle, "CMBufferQueueInstallTriggerWithIntegerThreshold", "10.7")
	registerFunc(&_cMBufferQueueIsAtEndOfData, &_cMBufferQueueIsAtEndOfDataErr, frameworkHandle, "CMBufferQueueIsAtEndOfData", "10.7")
	registerFunc(&_cMBufferQueueIsEmpty, &_cMBufferQueueIsEmptyErr, frameworkHandle, "CMBufferQueueIsEmpty", "10.7")
	registerFunc(&_cMBufferQueueMarkEndOfData, &_cMBufferQueueMarkEndOfDataErr, frameworkHandle, "CMBufferQueueMarkEndOfData", "10.7")
	registerFunc(&_cMBufferQueueRemoveTrigger, &_cMBufferQueueRemoveTriggerErr, frameworkHandle, "CMBufferQueueRemoveTrigger", "10.7")
	registerFunc(&_cMBufferQueueReset, &_cMBufferQueueResetErr, frameworkHandle, "CMBufferQueueReset", "10.7")
	registerFunc(&_cMBufferQueueResetWithCallback, &_cMBufferQueueResetWithCallbackErr, frameworkHandle, "CMBufferQueueResetWithCallback", "10.7")
	registerFunc(&_cMBufferQueueSetValidationCallback, &_cMBufferQueueSetValidationCallbackErr, frameworkHandle, "CMBufferQueueSetValidationCallback", "10.7")
	registerFunc(&_cMBufferQueueSetValidationHandler, &_cMBufferQueueSetValidationHandlerErr, frameworkHandle, "CMBufferQueueSetValidationHandler", "10.14.4")
	registerFunc(&_cMBufferQueueTestTrigger, &_cMBufferQueueTestTriggerErr, frameworkHandle, "CMBufferQueueTestTrigger", "10.7")
	registerFunc(&_cMClockConvertHostTimeToSystemUnits, &_cMClockConvertHostTimeToSystemUnitsErr, frameworkHandle, "CMClockConvertHostTimeToSystemUnits", "10.8")
	registerFunc(&_cMClockGetAnchorTime, &_cMClockGetAnchorTimeErr, frameworkHandle, "CMClockGetAnchorTime", "10.8")
	registerFunc(&_cMClockGetHostTimeClock, &_cMClockGetHostTimeClockErr, frameworkHandle, "CMClockGetHostTimeClock", "10.8")
	registerFunc(&_cMClockGetTime, &_cMClockGetTimeErr, frameworkHandle, "CMClockGetTime", "10.8")
	registerFunc(&_cMClockGetTypeID, &_cMClockGetTypeIDErr, frameworkHandle, "CMClockGetTypeID", "10.8")
	registerFunc(&_cMClockInvalidate, &_cMClockInvalidateErr, frameworkHandle, "CMClockInvalidate", "10.8")
	registerFunc(&_cMClockMakeHostTimeFromSystemUnits, &_cMClockMakeHostTimeFromSystemUnitsErr, frameworkHandle, "CMClockMakeHostTimeFromSystemUnits", "10.8")
	registerFunc(&_cMClockMightDrift, &_cMClockMightDriftErr, frameworkHandle, "CMClockMightDrift", "10.8")
	registerFunc(&_cMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer, &_cMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBufferErr, frameworkHandle, "CMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer", "10.10")
	registerFunc(&_cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer, &_cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBufferErr, frameworkHandle, "CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer", "10.10")
	registerFunc(&_cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData, &_cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionDataErr, frameworkHandle, "CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData", "10.10")
	registerFunc(&_cMCopyDictionaryOfAttachments, &_cMCopyDictionaryOfAttachmentsErr, frameworkHandle, "CMCopyDictionaryOfAttachments", "10.7")
	registerFunc(&_cMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout, &_cMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayoutErr, frameworkHandle, "CMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout", "10.10")
	registerFunc(&_cMFormatDescriptionCreate, &_cMFormatDescriptionCreateErr, frameworkHandle, "CMFormatDescriptionCreate", "10.7")
	registerFunc(&_cMFormatDescriptionEqual, &_cMFormatDescriptionEqualErr, frameworkHandle, "CMFormatDescriptionEqual", "10.7")
	registerFunc(&_cMFormatDescriptionEqualIgnoringExtensionKeys, &_cMFormatDescriptionEqualIgnoringExtensionKeysErr, frameworkHandle, "CMFormatDescriptionEqualIgnoringExtensionKeys", "10.7")
	registerFunc(&_cMFormatDescriptionGetExtension, &_cMFormatDescriptionGetExtensionErr, frameworkHandle, "CMFormatDescriptionGetExtension", "10.7")
	registerFunc(&_cMFormatDescriptionGetExtensions, &_cMFormatDescriptionGetExtensionsErr, frameworkHandle, "CMFormatDescriptionGetExtensions", "10.7")
	registerFunc(&_cMFormatDescriptionGetMediaSubType, &_cMFormatDescriptionGetMediaSubTypeErr, frameworkHandle, "CMFormatDescriptionGetMediaSubType", "10.7")
	registerFunc(&_cMFormatDescriptionGetMediaType, &_cMFormatDescriptionGetMediaTypeErr, frameworkHandle, "CMFormatDescriptionGetMediaType", "10.7")
	registerFunc(&_cMFormatDescriptionGetTypeID, &_cMFormatDescriptionGetTypeIDErr, frameworkHandle, "CMFormatDescriptionGetTypeID", "10.7")
	registerFunc(&_cMGetAttachment, &_cMGetAttachmentErr, frameworkHandle, "CMGetAttachment", "10.7")
	registerFunc(&_cMMemoryPoolCreate, &_cMMemoryPoolCreateErr, frameworkHandle, "CMMemoryPoolCreate", "10.8")
	registerFunc(&_cMMemoryPoolFlush, &_cMMemoryPoolFlushErr, frameworkHandle, "CMMemoryPoolFlush", "10.8")
	registerFunc(&_cMMemoryPoolGetAllocator, &_cMMemoryPoolGetAllocatorErr, frameworkHandle, "CMMemoryPoolGetAllocator", "10.8")
	registerFunc(&_cMMemoryPoolGetTypeID, &_cMMemoryPoolGetTypeIDErr, frameworkHandle, "CMMemoryPoolGetTypeID", "10.8")
	registerFunc(&_cMMemoryPoolInvalidate, &_cMMemoryPoolInvalidateErr, frameworkHandle, "CMMemoryPoolInvalidate", "10.8")
	registerFunc(&_cMMetadataCreateIdentifierForKeyAndKeySpace, &_cMMetadataCreateIdentifierForKeyAndKeySpaceErr, frameworkHandle, "CMMetadataCreateIdentifierForKeyAndKeySpace", "10.10")
	registerFunc(&_cMMetadataCreateKeyFromIdentifier, &_cMMetadataCreateKeyFromIdentifierErr, frameworkHandle, "CMMetadataCreateKeyFromIdentifier", "10.10")
	registerFunc(&_cMMetadataCreateKeyFromIdentifierAsCFData, &_cMMetadataCreateKeyFromIdentifierAsCFDataErr, frameworkHandle, "CMMetadataCreateKeyFromIdentifierAsCFData", "10.10")
	registerFunc(&_cMMetadataCreateKeySpaceFromIdentifier, &_cMMetadataCreateKeySpaceFromIdentifierErr, frameworkHandle, "CMMetadataCreateKeySpaceFromIdentifier", "10.10")
	registerFunc(&_cMMetadataDataTypeRegistryDataTypeConformsToDataType, &_cMMetadataDataTypeRegistryDataTypeConformsToDataTypeErr, frameworkHandle, "CMMetadataDataTypeRegistryDataTypeConformsToDataType", "10.10")
	registerFunc(&_cMMetadataDataTypeRegistryDataTypeIsBaseDataType, &_cMMetadataDataTypeRegistryDataTypeIsBaseDataTypeErr, frameworkHandle, "CMMetadataDataTypeRegistryDataTypeIsBaseDataType", "10.10")
	registerFunc(&_cMMetadataDataTypeRegistryDataTypeIsRegistered, &_cMMetadataDataTypeRegistryDataTypeIsRegisteredErr, frameworkHandle, "CMMetadataDataTypeRegistryDataTypeIsRegistered", "10.10")
	registerFunc(&_cMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType, &_cMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataTypeErr, frameworkHandle, "CMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType", "10.10")
	registerFunc(&_cMMetadataDataTypeRegistryGetBaseDataTypes, &_cMMetadataDataTypeRegistryGetBaseDataTypesErr, frameworkHandle, "CMMetadataDataTypeRegistryGetBaseDataTypes", "10.10")
	registerFunc(&_cMMetadataDataTypeRegistryGetConformingDataTypes, &_cMMetadataDataTypeRegistryGetConformingDataTypesErr, frameworkHandle, "CMMetadataDataTypeRegistryGetConformingDataTypes", "10.10")
	registerFunc(&_cMMetadataDataTypeRegistryGetDataTypeDescription, &_cMMetadataDataTypeRegistryGetDataTypeDescriptionErr, frameworkHandle, "CMMetadataDataTypeRegistryGetDataTypeDescription", "10.10")
	registerFunc(&_cMMetadataDataTypeRegistryRegisterDataType, &_cMMetadataDataTypeRegistryRegisterDataTypeErr, frameworkHandle, "CMMetadataDataTypeRegistryRegisterDataType", "10.10")
	registerFunc(&_cMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer, &_cMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBufferErr, frameworkHandle, "CMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer", "10.10")
	registerFunc(&_cMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions, &_cMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptionsErr, frameworkHandle, "CMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions", "10.10")
	registerFunc(&_cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer, &_cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBufferErr, frameworkHandle, "CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer", "10.10")
	registerFunc(&_cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData, &_cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionDataErr, frameworkHandle, "CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData", "10.10")
	registerFunc(&_cMMetadataFormatDescriptionCreateWithKeys, &_cMMetadataFormatDescriptionCreateWithKeysErr, frameworkHandle, "CMMetadataFormatDescriptionCreateWithKeys", "10.7")
	registerFunc(&_cMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications, &_cMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecificationsErr, frameworkHandle, "CMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications", "10.10")
	registerFunc(&_cMMetadataFormatDescriptionCreateWithMetadataSpecifications, &_cMMetadataFormatDescriptionCreateWithMetadataSpecificationsErr, frameworkHandle, "CMMetadataFormatDescriptionCreateWithMetadataSpecifications", "10.10")
	registerFunc(&_cMMetadataFormatDescriptionGetIdentifiers, &_cMMetadataFormatDescriptionGetIdentifiersErr, frameworkHandle, "CMMetadataFormatDescriptionGetIdentifiers", "10.10")
	registerFunc(&_cMMetadataFormatDescriptionGetKeyWithLocalID, &_cMMetadataFormatDescriptionGetKeyWithLocalIDErr, frameworkHandle, "CMMetadataFormatDescriptionGetKeyWithLocalID", "10.7")
	registerFunc(&_cMMuxedFormatDescriptionCreate, &_cMMuxedFormatDescriptionCreateErr, frameworkHandle, "CMMuxedFormatDescriptionCreate", "10.7")
	registerFunc(&_cMPropagateAttachments, &_cMPropagateAttachmentsErr, frameworkHandle, "CMPropagateAttachments", "10.7")
	registerFunc(&_cMRemoveAllAttachments, &_cMRemoveAllAttachmentsErr, frameworkHandle, "CMRemoveAllAttachments", "10.7")
	registerFunc(&_cMRemoveAttachment, &_cMRemoveAttachmentErr, frameworkHandle, "CMRemoveAttachment", "10.7")
	registerFunc(&_cMSampleBufferCallBlockForEachSample, &_cMSampleBufferCallBlockForEachSampleErr, frameworkHandle, "CMSampleBufferCallBlockForEachSample", "10.10")
	registerFunc(&_cMSampleBufferCallForEachSample, &_cMSampleBufferCallForEachSampleErr, frameworkHandle, "CMSampleBufferCallForEachSample", "10.7")
	registerFunc(&_cMSampleBufferCopyPCMDataIntoAudioBufferList, &_cMSampleBufferCopyPCMDataIntoAudioBufferListErr, frameworkHandle, "CMSampleBufferCopyPCMDataIntoAudioBufferList", "10.9")
	registerFunc(&_cMSampleBufferCopySampleBufferForRange, &_cMSampleBufferCopySampleBufferForRangeErr, frameworkHandle, "CMSampleBufferCopySampleBufferForRange", "10.7")
	registerFunc(&_cMSampleBufferCreate, &_cMSampleBufferCreateErr, frameworkHandle, "CMSampleBufferCreate", "10.7")
	registerFunc(&_cMSampleBufferCreateCopy, &_cMSampleBufferCreateCopyErr, frameworkHandle, "CMSampleBufferCreateCopy", "10.7")
	registerFunc(&_cMSampleBufferCreateCopyWithNewTiming, &_cMSampleBufferCreateCopyWithNewTimingErr, frameworkHandle, "CMSampleBufferCreateCopyWithNewTiming", "10.7")
	registerFunc(&_cMSampleBufferCreateForImageBuffer, &_cMSampleBufferCreateForImageBufferErr, frameworkHandle, "CMSampleBufferCreateForImageBuffer", "10.7")
	registerFunc(&_cMSampleBufferCreateForImageBufferWithMakeDataReadyHandler, &_cMSampleBufferCreateForImageBufferWithMakeDataReadyHandlerErr, frameworkHandle, "CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler", "10.14.4")
	registerFunc(&_cMSampleBufferCreateForTaggedBufferGroup, &_cMSampleBufferCreateForTaggedBufferGroupErr, frameworkHandle, "CMSampleBufferCreateForTaggedBufferGroup", "14.0")
	registerFunc(&_cMSampleBufferCreateReady, &_cMSampleBufferCreateReadyErr, frameworkHandle, "CMSampleBufferCreateReady", "10.10")
	registerFunc(&_cMSampleBufferCreateReadyWithImageBuffer, &_cMSampleBufferCreateReadyWithImageBufferErr, frameworkHandle, "CMSampleBufferCreateReadyWithImageBuffer", "10.10")
	registerFunc(&_cMSampleBufferCreateWithMakeDataReadyHandler, &_cMSampleBufferCreateWithMakeDataReadyHandlerErr, frameworkHandle, "CMSampleBufferCreateWithMakeDataReadyHandler", "10.14.4")
	registerFunc(&_cMSampleBufferDataIsReady, &_cMSampleBufferDataIsReadyErr, frameworkHandle, "CMSampleBufferDataIsReady", "10.7")
	registerFunc(&_cMSampleBufferGetAudioBufferListWithRetainedBlockBuffer, &_cMSampleBufferGetAudioBufferListWithRetainedBlockBufferErr, frameworkHandle, "CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer", "10.7")
	registerFunc(&_cMSampleBufferGetAudioStreamPacketDescriptions, &_cMSampleBufferGetAudioStreamPacketDescriptionsErr, frameworkHandle, "CMSampleBufferGetAudioStreamPacketDescriptions", "10.7")
	registerFunc(&_cMSampleBufferGetAudioStreamPacketDescriptionsPtr, &_cMSampleBufferGetAudioStreamPacketDescriptionsPtrErr, frameworkHandle, "CMSampleBufferGetAudioStreamPacketDescriptionsPtr", "10.7")
	registerFunc(&_cMSampleBufferGetDataBuffer, &_cMSampleBufferGetDataBufferErr, frameworkHandle, "CMSampleBufferGetDataBuffer", "10.7")
	registerFunc(&_cMSampleBufferGetDecodeTimeStamp, &_cMSampleBufferGetDecodeTimeStampErr, frameworkHandle, "CMSampleBufferGetDecodeTimeStamp", "10.7")
	registerFunc(&_cMSampleBufferGetDuration, &_cMSampleBufferGetDurationErr, frameworkHandle, "CMSampleBufferGetDuration", "10.7")
	registerFunc(&_cMSampleBufferGetFormatDescription, &_cMSampleBufferGetFormatDescriptionErr, frameworkHandle, "CMSampleBufferGetFormatDescription", "10.7")
	registerFunc(&_cMSampleBufferGetImageBuffer, &_cMSampleBufferGetImageBufferErr, frameworkHandle, "CMSampleBufferGetImageBuffer", "10.7")
	registerFunc(&_cMSampleBufferGetNumSamples, &_cMSampleBufferGetNumSamplesErr, frameworkHandle, "CMSampleBufferGetNumSamples", "10.7")
	registerFunc(&_cMSampleBufferGetOutputDecodeTimeStamp, &_cMSampleBufferGetOutputDecodeTimeStampErr, frameworkHandle, "CMSampleBufferGetOutputDecodeTimeStamp", "10.7")
	registerFunc(&_cMSampleBufferGetOutputDuration, &_cMSampleBufferGetOutputDurationErr, frameworkHandle, "CMSampleBufferGetOutputDuration", "10.7")
	registerFunc(&_cMSampleBufferGetOutputPresentationTimeStamp, &_cMSampleBufferGetOutputPresentationTimeStampErr, frameworkHandle, "CMSampleBufferGetOutputPresentationTimeStamp", "10.7")
	registerFunc(&_cMSampleBufferGetOutputSampleTimingInfoArray, &_cMSampleBufferGetOutputSampleTimingInfoArrayErr, frameworkHandle, "CMSampleBufferGetOutputSampleTimingInfoArray", "10.7")
	registerFunc(&_cMSampleBufferGetPresentationTimeStamp, &_cMSampleBufferGetPresentationTimeStampErr, frameworkHandle, "CMSampleBufferGetPresentationTimeStamp", "10.7")
	registerFunc(&_cMSampleBufferGetSampleAttachmentsArray, &_cMSampleBufferGetSampleAttachmentsArrayErr, frameworkHandle, "CMSampleBufferGetSampleAttachmentsArray", "10.7")
	registerFunc(&_cMSampleBufferGetSampleSize, &_cMSampleBufferGetSampleSizeErr, frameworkHandle, "CMSampleBufferGetSampleSize", "10.7")
	registerFunc(&_cMSampleBufferGetSampleSizeArray, &_cMSampleBufferGetSampleSizeArrayErr, frameworkHandle, "CMSampleBufferGetSampleSizeArray", "10.7")
	registerFunc(&_cMSampleBufferGetSampleTimingInfo, &_cMSampleBufferGetSampleTimingInfoErr, frameworkHandle, "CMSampleBufferGetSampleTimingInfo", "10.7")
	registerFunc(&_cMSampleBufferGetSampleTimingInfoArray, &_cMSampleBufferGetSampleTimingInfoArrayErr, frameworkHandle, "CMSampleBufferGetSampleTimingInfoArray", "10.7")
	registerFunc(&_cMSampleBufferGetTaggedBufferGroup, &_cMSampleBufferGetTaggedBufferGroupErr, frameworkHandle, "CMSampleBufferGetTaggedBufferGroup", "14.0")
	registerFunc(&_cMSampleBufferGetTotalSampleSize, &_cMSampleBufferGetTotalSampleSizeErr, frameworkHandle, "CMSampleBufferGetTotalSampleSize", "10.7")
	registerFunc(&_cMSampleBufferGetTypeID, &_cMSampleBufferGetTypeIDErr, frameworkHandle, "CMSampleBufferGetTypeID", "10.7")
	registerFunc(&_cMSampleBufferHasDataFailed, &_cMSampleBufferHasDataFailedErr, frameworkHandle, "CMSampleBufferHasDataFailed", "10.10")
	registerFunc(&_cMSampleBufferInvalidate, &_cMSampleBufferInvalidateErr, frameworkHandle, "CMSampleBufferInvalidate", "10.7")
	registerFunc(&_cMSampleBufferIsValid, &_cMSampleBufferIsValidErr, frameworkHandle, "CMSampleBufferIsValid", "10.7")
	registerFunc(&_cMSampleBufferMakeDataReady, &_cMSampleBufferMakeDataReadyErr, frameworkHandle, "CMSampleBufferMakeDataReady", "10.7")
	registerFunc(&_cMSampleBufferSetDataBuffer, &_cMSampleBufferSetDataBufferErr, frameworkHandle, "CMSampleBufferSetDataBuffer", "10.7")
	registerFunc(&_cMSampleBufferSetDataBufferFromAudioBufferList, &_cMSampleBufferSetDataBufferFromAudioBufferListErr, frameworkHandle, "CMSampleBufferSetDataBufferFromAudioBufferList", "10.7")
	registerFunc(&_cMSampleBufferSetDataFailed, &_cMSampleBufferSetDataFailedErr, frameworkHandle, "CMSampleBufferSetDataFailed", "10.10")
	registerFunc(&_cMSampleBufferSetDataReady, &_cMSampleBufferSetDataReadyErr, frameworkHandle, "CMSampleBufferSetDataReady", "10.7")
	registerFunc(&_cMSampleBufferSetInvalidateCallback, &_cMSampleBufferSetInvalidateCallbackErr, frameworkHandle, "CMSampleBufferSetInvalidateCallback", "10.7")
	registerFunc(&_cMSampleBufferSetInvalidateHandler, &_cMSampleBufferSetInvalidateHandlerErr, frameworkHandle, "CMSampleBufferSetInvalidateHandler", "10.10")
	registerFunc(&_cMSampleBufferSetOutputPresentationTimeStamp, &_cMSampleBufferSetOutputPresentationTimeStampErr, frameworkHandle, "CMSampleBufferSetOutputPresentationTimeStamp", "10.7")
	registerFunc(&_cMSampleBufferTrackDataReadiness, &_cMSampleBufferTrackDataReadinessErr, frameworkHandle, "CMSampleBufferTrackDataReadiness", "10.7")
	registerFunc(&_cMSetAttachment, &_cMSetAttachmentErr, frameworkHandle, "CMSetAttachment", "10.7")
	registerFunc(&_cMSetAttachments, &_cMSetAttachmentsErr, frameworkHandle, "CMSetAttachments", "10.7")
	registerFunc(&_cMSimpleQueueCreate, &_cMSimpleQueueCreateErr, frameworkHandle, "CMSimpleQueueCreate", "10.7")
	registerFunc(&_cMSimpleQueueDequeue, &_cMSimpleQueueDequeueErr, frameworkHandle, "CMSimpleQueueDequeue", "10.7")
	registerFunc(&_cMSimpleQueueEnqueue, &_cMSimpleQueueEnqueueErr, frameworkHandle, "CMSimpleQueueEnqueue", "10.7")
	registerFunc(&_cMSimpleQueueGetCapacity, &_cMSimpleQueueGetCapacityErr, frameworkHandle, "CMSimpleQueueGetCapacity", "10.7")
	registerFunc(&_cMSimpleQueueGetCount, &_cMSimpleQueueGetCountErr, frameworkHandle, "CMSimpleQueueGetCount", "10.7")
	registerFunc(&_cMSimpleQueueGetHead, &_cMSimpleQueueGetHeadErr, frameworkHandle, "CMSimpleQueueGetHead", "10.7")
	registerFunc(&_cMSimpleQueueGetTypeID, &_cMSimpleQueueGetTypeIDErr, frameworkHandle, "CMSimpleQueueGetTypeID", "10.7")
	registerFunc(&_cMSimpleQueueReset, &_cMSimpleQueueResetErr, frameworkHandle, "CMSimpleQueueReset", "10.7")
	registerFunc(&_cMSwapBigEndianClosedCaptionDescriptionToHost, &_cMSwapBigEndianClosedCaptionDescriptionToHostErr, frameworkHandle, "CMSwapBigEndianClosedCaptionDescriptionToHost", "10.10")
	registerFunc(&_cMSwapBigEndianImageDescriptionToHost, &_cMSwapBigEndianImageDescriptionToHostErr, frameworkHandle, "CMSwapBigEndianImageDescriptionToHost", "10.10")
	registerFunc(&_cMSwapBigEndianMetadataDescriptionToHost, &_cMSwapBigEndianMetadataDescriptionToHostErr, frameworkHandle, "CMSwapBigEndianMetadataDescriptionToHost", "10.10")
	registerFunc(&_cMSwapBigEndianSoundDescriptionToHost, &_cMSwapBigEndianSoundDescriptionToHostErr, frameworkHandle, "CMSwapBigEndianSoundDescriptionToHost", "10.10")
	registerFunc(&_cMSwapBigEndianTextDescriptionToHost, &_cMSwapBigEndianTextDescriptionToHostErr, frameworkHandle, "CMSwapBigEndianTextDescriptionToHost", "10.10")
	registerFunc(&_cMSwapBigEndianTimeCodeDescriptionToHost, &_cMSwapBigEndianTimeCodeDescriptionToHostErr, frameworkHandle, "CMSwapBigEndianTimeCodeDescriptionToHost", "10.10")
	registerFunc(&_cMSwapHostEndianClosedCaptionDescriptionToBig, &_cMSwapHostEndianClosedCaptionDescriptionToBigErr, frameworkHandle, "CMSwapHostEndianClosedCaptionDescriptionToBig", "10.10")
	registerFunc(&_cMSwapHostEndianImageDescriptionToBig, &_cMSwapHostEndianImageDescriptionToBigErr, frameworkHandle, "CMSwapHostEndianImageDescriptionToBig", "10.10")
	registerFunc(&_cMSwapHostEndianMetadataDescriptionToBig, &_cMSwapHostEndianMetadataDescriptionToBigErr, frameworkHandle, "CMSwapHostEndianMetadataDescriptionToBig", "10.10")
	registerFunc(&_cMSwapHostEndianSoundDescriptionToBig, &_cMSwapHostEndianSoundDescriptionToBigErr, frameworkHandle, "CMSwapHostEndianSoundDescriptionToBig", "10.10")
	registerFunc(&_cMSwapHostEndianTextDescriptionToBig, &_cMSwapHostEndianTextDescriptionToBigErr, frameworkHandle, "CMSwapHostEndianTextDescriptionToBig", "10.10")
	registerFunc(&_cMSwapHostEndianTimeCodeDescriptionToBig, &_cMSwapHostEndianTimeCodeDescriptionToBigErr, frameworkHandle, "CMSwapHostEndianTimeCodeDescriptionToBig", "10.10")
	registerFunc(&_cMSyncConvertTime, &_cMSyncConvertTimeErr, frameworkHandle, "CMSyncConvertTime", "10.8")
	registerFunc(&_cMSyncGetRelativeRate, &_cMSyncGetRelativeRateErr, frameworkHandle, "CMSyncGetRelativeRate", "10.8")
	registerFunc(&_cMSyncGetRelativeRateAndAnchorTime, &_cMSyncGetRelativeRateAndAnchorTimeErr, frameworkHandle, "CMSyncGetRelativeRateAndAnchorTime", "10.8")
	registerFunc(&_cMSyncGetTime, &_cMSyncGetTimeErr, frameworkHandle, "CMSyncGetTime", "10.8")
	registerFunc(&_cMSyncMightDrift, &_cMSyncMightDriftErr, frameworkHandle, "CMSyncMightDrift", "10.8")
	registerFunc(&_cMTagCollectionAddTag, &_cMTagCollectionAddTagErr, frameworkHandle, "CMTagCollectionAddTag", "14.0")
	registerFunc(&_cMTagCollectionAddTagsFromArray, &_cMTagCollectionAddTagsFromArrayErr, frameworkHandle, "CMTagCollectionAddTagsFromArray", "14.0")
	registerFunc(&_cMTagCollectionAddTagsFromCollection, &_cMTagCollectionAddTagsFromCollectionErr, frameworkHandle, "CMTagCollectionAddTagsFromCollection", "14.0")
	registerFunc(&_cMTagCollectionApply, &_cMTagCollectionApplyErr, frameworkHandle, "CMTagCollectionApply", "14.0")
	registerFunc(&_cMTagCollectionApplyUntil, &_cMTagCollectionApplyUntilErr, frameworkHandle, "CMTagCollectionApplyUntil", "14.0")
	registerFunc(&_cMTagCollectionContainsCategory, &_cMTagCollectionContainsCategoryErr, frameworkHandle, "CMTagCollectionContainsCategory", "14.0")
	registerFunc(&_cMTagCollectionContainsSpecifiedTags, &_cMTagCollectionContainsSpecifiedTagsErr, frameworkHandle, "CMTagCollectionContainsSpecifiedTags", "14.0")
	registerFunc(&_cMTagCollectionContainsTag, &_cMTagCollectionContainsTagErr, frameworkHandle, "CMTagCollectionContainsTag", "14.0")
	registerFunc(&_cMTagCollectionContainsTagsOfCollection, &_cMTagCollectionContainsTagsOfCollectionErr, frameworkHandle, "CMTagCollectionContainsTagsOfCollection", "14.0")
	registerFunc(&_cMTagCollectionCopyAsData, &_cMTagCollectionCopyAsDataErr, frameworkHandle, "CMTagCollectionCopyAsData", "14.0")
	registerFunc(&_cMTagCollectionCopyAsDictionary, &_cMTagCollectionCopyAsDictionaryErr, frameworkHandle, "CMTagCollectionCopyAsDictionary", "14.0")
	registerFunc(&_cMTagCollectionCopyDescription, &_cMTagCollectionCopyDescriptionErr, frameworkHandle, "CMTagCollectionCopyDescription", "14.0")
	registerFunc(&_cMTagCollectionCopyTagsOfCategories, &_cMTagCollectionCopyTagsOfCategoriesErr, frameworkHandle, "CMTagCollectionCopyTagsOfCategories", "14.0")
	registerFunc(&_cMTagCollectionCountTagsWithFilterFunction, &_cMTagCollectionCountTagsWithFilterFunctionErr, frameworkHandle, "CMTagCollectionCountTagsWithFilterFunction", "14.0")
	registerFunc(&_cMTagCollectionCreate, &_cMTagCollectionCreateErr, frameworkHandle, "CMTagCollectionCreate", "14.0")
	registerFunc(&_cMTagCollectionCreateCopy, &_cMTagCollectionCreateCopyErr, frameworkHandle, "CMTagCollectionCreateCopy", "14.0")
	registerFunc(&_cMTagCollectionCreateDifference, &_cMTagCollectionCreateDifferenceErr, frameworkHandle, "CMTagCollectionCreateDifference", "14.0")
	registerFunc(&_cMTagCollectionCreateExclusiveOr, &_cMTagCollectionCreateExclusiveOrErr, frameworkHandle, "CMTagCollectionCreateExclusiveOr", "14.0")
	registerFunc(&_cMTagCollectionCreateFromData, &_cMTagCollectionCreateFromDataErr, frameworkHandle, "CMTagCollectionCreateFromData", "14.0")
	registerFunc(&_cMTagCollectionCreateFromDictionary, &_cMTagCollectionCreateFromDictionaryErr, frameworkHandle, "CMTagCollectionCreateFromDictionary", "14.0")
	registerFunc(&_cMTagCollectionCreateIntersection, &_cMTagCollectionCreateIntersectionErr, frameworkHandle, "CMTagCollectionCreateIntersection", "14.0")
	registerFunc(&_cMTagCollectionCreateMutable, &_cMTagCollectionCreateMutableErr, frameworkHandle, "CMTagCollectionCreateMutable", "14.0")
	registerFunc(&_cMTagCollectionCreateMutableCopy, &_cMTagCollectionCreateMutableCopyErr, frameworkHandle, "CMTagCollectionCreateMutableCopy", "14.0")
	registerFunc(&_cMTagCollectionCreateUnion, &_cMTagCollectionCreateUnionErr, frameworkHandle, "CMTagCollectionCreateUnion", "14.0")
	registerFunc(&_cMTagCollectionGetCount, &_cMTagCollectionGetCountErr, frameworkHandle, "CMTagCollectionGetCount", "14.0")
	registerFunc(&_cMTagCollectionGetCountOfCategory, &_cMTagCollectionGetCountOfCategoryErr, frameworkHandle, "CMTagCollectionGetCountOfCategory", "14.0")
	registerFunc(&_cMTagCollectionGetTags, &_cMTagCollectionGetTagsErr, frameworkHandle, "CMTagCollectionGetTags", "14.0")
	registerFunc(&_cMTagCollectionGetTagsWithCategory, &_cMTagCollectionGetTagsWithCategoryErr, frameworkHandle, "CMTagCollectionGetTagsWithCategory", "14.0")
	registerFunc(&_cMTagCollectionGetTagsWithFilterFunction, &_cMTagCollectionGetTagsWithFilterFunctionErr, frameworkHandle, "CMTagCollectionGetTagsWithFilterFunction", "14.0")
	registerFunc(&_cMTagCollectionGetTypeID, &_cMTagCollectionGetTypeIDErr, frameworkHandle, "CMTagCollectionGetTypeID", "14.0")
	registerFunc(&_cMTagCollectionIsEmpty, &_cMTagCollectionIsEmptyErr, frameworkHandle, "CMTagCollectionIsEmpty", "14.0")
	registerFunc(&_cMTagCollectionRemoveAllTags, &_cMTagCollectionRemoveAllTagsErr, frameworkHandle, "CMTagCollectionRemoveAllTags", "14.0")
	registerFunc(&_cMTagCollectionRemoveAllTagsOfCategory, &_cMTagCollectionRemoveAllTagsOfCategoryErr, frameworkHandle, "CMTagCollectionRemoveAllTagsOfCategory", "14.0")
	registerFunc(&_cMTagCollectionRemoveTag, &_cMTagCollectionRemoveTagErr, frameworkHandle, "CMTagCollectionRemoveTag", "14.0")
	registerFunc(&_cMTagCompare, &_cMTagCompareErr, frameworkHandle, "CMTagCompare", "14.0")
	registerFunc(&_cMTagCopyAsDictionary, &_cMTagCopyAsDictionaryErr, frameworkHandle, "CMTagCopyAsDictionary", "14.0")
	registerFunc(&_cMTagCopyDescription, &_cMTagCopyDescriptionErr, frameworkHandle, "CMTagCopyDescription", "14.0")
	registerFunc(&_cMTagEqualToTag, &_cMTagEqualToTagErr, frameworkHandle, "CMTagEqualToTag", "14.0")
	registerFunc(&_cMTagGetFlagsValue, &_cMTagGetFlagsValueErr, frameworkHandle, "CMTagGetFlagsValue", "14.0")
	registerFunc(&_cMTagGetFloat64Value, &_cMTagGetFloat64ValueErr, frameworkHandle, "CMTagGetFloat64Value", "14.0")
	registerFunc(&_cMTagGetOSTypeValue, &_cMTagGetOSTypeValueErr, frameworkHandle, "CMTagGetOSTypeValue", "14.0")
	registerFunc(&_cMTagGetSInt64Value, &_cMTagGetSInt64ValueErr, frameworkHandle, "CMTagGetSInt64Value", "14.0")
	registerFunc(&_cMTagGetValueDataType, &_cMTagGetValueDataTypeErr, frameworkHandle, "CMTagGetValueDataType", "14.0")
	registerFunc(&_cMTagHasFlagsValue, &_cMTagHasFlagsValueErr, frameworkHandle, "CMTagHasFlagsValue", "14.0")
	registerFunc(&_cMTagHasFloat64Value, &_cMTagHasFloat64ValueErr, frameworkHandle, "CMTagHasFloat64Value", "14.0")
	registerFunc(&_cMTagHasOSTypeValue, &_cMTagHasOSTypeValueErr, frameworkHandle, "CMTagHasOSTypeValue", "14.0")
	registerFunc(&_cMTagHasSInt64Value, &_cMTagHasSInt64ValueErr, frameworkHandle, "CMTagHasSInt64Value", "14.0")
	registerFunc(&_cMTagHash, &_cMTagHashErr, frameworkHandle, "CMTagHash", "14.0")
	registerFunc(&_cMTagMakeFromDictionary, &_cMTagMakeFromDictionaryErr, frameworkHandle, "CMTagMakeFromDictionary", "14.0")
	registerFunc(&_cMTagMakeWithFlagsValue, &_cMTagMakeWithFlagsValueErr, frameworkHandle, "CMTagMakeWithFlagsValue", "14.0")
	registerFunc(&_cMTagMakeWithFloat64Value, &_cMTagMakeWithFloat64ValueErr, frameworkHandle, "CMTagMakeWithFloat64Value", "14.0")
	registerFunc(&_cMTagMakeWithOSTypeValue, &_cMTagMakeWithOSTypeValueErr, frameworkHandle, "CMTagMakeWithOSTypeValue", "14.0")
	registerFunc(&_cMTagMakeWithSInt64Value, &_cMTagMakeWithSInt64ValueErr, frameworkHandle, "CMTagMakeWithSInt64Value", "14.0")
	registerFunc(&_cMTaggedBufferGroupCreate, &_cMTaggedBufferGroupCreateErr, frameworkHandle, "CMTaggedBufferGroupCreate", "14.0")
	registerFunc(&_cMTaggedBufferGroupCreateCombined, &_cMTaggedBufferGroupCreateCombinedErr, frameworkHandle, "CMTaggedBufferGroupCreateCombined", "14.0")
	registerFunc(&_cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup, &_cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupErr, frameworkHandle, "CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup", "14.0")
	registerFunc(&_cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions, &_cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensionsErr, frameworkHandle, "CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions", "26.0")
	registerFunc(&_cMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup, &_cMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroupErr, frameworkHandle, "CMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup", "14.0")
	registerFunc(&_cMTaggedBufferGroupGetCMSampleBufferAtIndex, &_cMTaggedBufferGroupGetCMSampleBufferAtIndexErr, frameworkHandle, "CMTaggedBufferGroupGetCMSampleBufferAtIndex", "14.0")
	registerFunc(&_cMTaggedBufferGroupGetCMSampleBufferForTag, &_cMTaggedBufferGroupGetCMSampleBufferForTagErr, frameworkHandle, "CMTaggedBufferGroupGetCMSampleBufferForTag", "14.0")
	registerFunc(&_cMTaggedBufferGroupGetCMSampleBufferForTagCollection, &_cMTaggedBufferGroupGetCMSampleBufferForTagCollectionErr, frameworkHandle, "CMTaggedBufferGroupGetCMSampleBufferForTagCollection", "14.0")
	registerFunc(&_cMTaggedBufferGroupGetCVPixelBufferAtIndex, &_cMTaggedBufferGroupGetCVPixelBufferAtIndexErr, frameworkHandle, "CMTaggedBufferGroupGetCVPixelBufferAtIndex", "14.0")
	registerFunc(&_cMTaggedBufferGroupGetCVPixelBufferForTag, &_cMTaggedBufferGroupGetCVPixelBufferForTagErr, frameworkHandle, "CMTaggedBufferGroupGetCVPixelBufferForTag", "14.0")
	registerFunc(&_cMTaggedBufferGroupGetCVPixelBufferForTagCollection, &_cMTaggedBufferGroupGetCVPixelBufferForTagCollectionErr, frameworkHandle, "CMTaggedBufferGroupGetCVPixelBufferForTagCollection", "14.0")
	registerFunc(&_cMTaggedBufferGroupGetCount, &_cMTaggedBufferGroupGetCountErr, frameworkHandle, "CMTaggedBufferGroupGetCount", "14.0")
	registerFunc(&_cMTaggedBufferGroupGetNumberOfMatchesForTagCollection, &_cMTaggedBufferGroupGetNumberOfMatchesForTagCollectionErr, frameworkHandle, "CMTaggedBufferGroupGetNumberOfMatchesForTagCollection", "14.0")
	registerFunc(&_cMTaggedBufferGroupGetTagCollectionAtIndex, &_cMTaggedBufferGroupGetTagCollectionAtIndexErr, frameworkHandle, "CMTaggedBufferGroupGetTagCollectionAtIndex", "14.0")
	registerFunc(&_cMTaggedBufferGroupGetTypeID, &_cMTaggedBufferGroupGetTypeIDErr, frameworkHandle, "CMTaggedBufferGroupGetTypeID", "14.0")
	registerFunc(&_cMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer, &_cMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBufferErr, frameworkHandle, "CMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer", "10.10")
	registerFunc(&_cMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer, &_cMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBufferErr, frameworkHandle, "CMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer", "10.10")
	registerFunc(&_cMTextFormatDescriptionCreateFromBigEndianTextDescriptionData, &_cMTextFormatDescriptionCreateFromBigEndianTextDescriptionDataErr, frameworkHandle, "CMTextFormatDescriptionCreateFromBigEndianTextDescriptionData", "10.10")
	registerFunc(&_cMTextFormatDescriptionGetDefaultStyle, &_cMTextFormatDescriptionGetDefaultStyleErr, frameworkHandle, "CMTextFormatDescriptionGetDefaultStyle", "10.7")
	registerFunc(&_cMTextFormatDescriptionGetDefaultTextBox, &_cMTextFormatDescriptionGetDefaultTextBoxErr, frameworkHandle, "CMTextFormatDescriptionGetDefaultTextBox", "10.7")
	registerFunc(&_cMTextFormatDescriptionGetDisplayFlags, &_cMTextFormatDescriptionGetDisplayFlagsErr, frameworkHandle, "CMTextFormatDescriptionGetDisplayFlags", "10.7")
	registerFunc(&_cMTextFormatDescriptionGetFontName, &_cMTextFormatDescriptionGetFontNameErr, frameworkHandle, "CMTextFormatDescriptionGetFontName", "10.7")
	registerFunc(&_cMTextFormatDescriptionGetJustification, &_cMTextFormatDescriptionGetJustificationErr, frameworkHandle, "CMTextFormatDescriptionGetJustification", "10.7")
	registerFunc(&_cMTimeAbsoluteValue, &_cMTimeAbsoluteValueErr, frameworkHandle, "CMTimeAbsoluteValue", "10.7")
	registerFunc(&_cMTimeAdd, &_cMTimeAddErr, frameworkHandle, "CMTimeAdd", "10.7")
	registerFunc(&_cMTimeClampToRange, &_cMTimeClampToRangeErr, frameworkHandle, "CMTimeClampToRange", "10.7")
	registerFunc(&_cMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer, &_cMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBufferErr, frameworkHandle, "CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer", "10.10")
	registerFunc(&_cMTimeCodeFormatDescriptionCreate, &_cMTimeCodeFormatDescriptionCreateErr, frameworkHandle, "CMTimeCodeFormatDescriptionCreate", "10.7")
	registerFunc(&_cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer, &_cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBufferErr, frameworkHandle, "CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer", "10.10")
	registerFunc(&_cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData, &_cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionDataErr, frameworkHandle, "CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData", "10.10")
	registerFunc(&_cMTimeCodeFormatDescriptionGetFrameDuration, &_cMTimeCodeFormatDescriptionGetFrameDurationErr, frameworkHandle, "CMTimeCodeFormatDescriptionGetFrameDuration", "10.7")
	registerFunc(&_cMTimeCodeFormatDescriptionGetFrameQuanta, &_cMTimeCodeFormatDescriptionGetFrameQuantaErr, frameworkHandle, "CMTimeCodeFormatDescriptionGetFrameQuanta", "10.7")
	registerFunc(&_cMTimeCodeFormatDescriptionGetTimeCodeFlags, &_cMTimeCodeFormatDescriptionGetTimeCodeFlagsErr, frameworkHandle, "CMTimeCodeFormatDescriptionGetTimeCodeFlags", "10.7")
	registerFunc(&_cMTimeCompare, &_cMTimeCompareErr, frameworkHandle, "CMTimeCompare", "10.7")
	registerFunc(&_cMTimeConvertScale, &_cMTimeConvertScaleErr, frameworkHandle, "CMTimeConvertScale", "10.7")
	registerFunc(&_cMTimeCopyAsDictionary, &_cMTimeCopyAsDictionaryErr, frameworkHandle, "CMTimeCopyAsDictionary", "10.7")
	registerFunc(&_cMTimeCopyDescription, &_cMTimeCopyDescriptionErr, frameworkHandle, "CMTimeCopyDescription", "10.7")
	registerFunc(&_cMTimeFoldIntoRange, &_cMTimeFoldIntoRangeErr, frameworkHandle, "CMTimeFoldIntoRange", "10.14")
	registerFunc(&_cMTimeGetSeconds, &_cMTimeGetSecondsErr, frameworkHandle, "CMTimeGetSeconds", "10.7")
	registerFunc(&_cMTimeMake, &_cMTimeMakeErr, frameworkHandle, "CMTimeMake", "10.7")
	registerFunc(&_cMTimeMakeFromDictionary, &_cMTimeMakeFromDictionaryErr, frameworkHandle, "CMTimeMakeFromDictionary", "10.7")
	registerFunc(&_cMTimeMakeWithEpoch, &_cMTimeMakeWithEpochErr, frameworkHandle, "CMTimeMakeWithEpoch", "10.7")
	registerFunc(&_cMTimeMakeWithSeconds, &_cMTimeMakeWithSecondsErr, frameworkHandle, "CMTimeMakeWithSeconds", "10.7")
	registerFunc(&_cMTimeMapDurationFromRangeToRange, &_cMTimeMapDurationFromRangeToRangeErr, frameworkHandle, "CMTimeMapDurationFromRangeToRange", "10.7")
	registerFunc(&_cMTimeMapTimeFromRangeToRange, &_cMTimeMapTimeFromRangeToRangeErr, frameworkHandle, "CMTimeMapTimeFromRangeToRange", "10.7")
	registerFunc(&_cMTimeMappingCopyAsDictionary, &_cMTimeMappingCopyAsDictionaryErr, frameworkHandle, "CMTimeMappingCopyAsDictionary", "10.11")
	registerFunc(&_cMTimeMappingCopyDescription, &_cMTimeMappingCopyDescriptionErr, frameworkHandle, "CMTimeMappingCopyDescription", "10.11")
	registerFunc(&_cMTimeMappingMake, &_cMTimeMappingMakeErr, frameworkHandle, "CMTimeMappingMake", "10.11")
	registerFunc(&_cMTimeMappingMakeEmpty, &_cMTimeMappingMakeEmptyErr, frameworkHandle, "CMTimeMappingMakeEmpty", "10.11")
	registerFunc(&_cMTimeMappingMakeFromDictionary, &_cMTimeMappingMakeFromDictionaryErr, frameworkHandle, "CMTimeMappingMakeFromDictionary", "10.11")
	registerFunc(&_cMTimeMappingShow, &_cMTimeMappingShowErr, frameworkHandle, "CMTimeMappingShow", "10.11")
	registerFunc(&_cMTimeMaximum, &_cMTimeMaximumErr, frameworkHandle, "CMTimeMaximum", "10.7")
	registerFunc(&_cMTimeMinimum, &_cMTimeMinimumErr, frameworkHandle, "CMTimeMinimum", "10.7")
	registerFunc(&_cMTimeMultiply, &_cMTimeMultiplyErr, frameworkHandle, "CMTimeMultiply", "10.7")
	registerFunc(&_cMTimeMultiplyByFloat64, &_cMTimeMultiplyByFloat64Err, frameworkHandle, "CMTimeMultiplyByFloat64", "10.7")
	registerFunc(&_cMTimeMultiplyByRatio, &_cMTimeMultiplyByRatioErr, frameworkHandle, "CMTimeMultiplyByRatio", "10.10")
	registerFunc(&_cMTimeRangeContainsTime, &_cMTimeRangeContainsTimeErr, frameworkHandle, "CMTimeRangeContainsTime", "10.7")
	registerFunc(&_cMTimeRangeContainsTimeRange, &_cMTimeRangeContainsTimeRangeErr, frameworkHandle, "CMTimeRangeContainsTimeRange", "10.7")
	registerFunc(&_cMTimeRangeCopyAsDictionary, &_cMTimeRangeCopyAsDictionaryErr, frameworkHandle, "CMTimeRangeCopyAsDictionary", "10.7")
	registerFunc(&_cMTimeRangeCopyDescription, &_cMTimeRangeCopyDescriptionErr, frameworkHandle, "CMTimeRangeCopyDescription", "10.7")
	registerFunc(&_cMTimeRangeEqual, &_cMTimeRangeEqualErr, frameworkHandle, "CMTimeRangeEqual", "10.7")
	registerFunc(&_cMTimeRangeFromTimeToTime, &_cMTimeRangeFromTimeToTimeErr, frameworkHandle, "CMTimeRangeFromTimeToTime", "10.7")
	registerFunc(&_cMTimeRangeGetEnd, &_cMTimeRangeGetEndErr, frameworkHandle, "CMTimeRangeGetEnd", "10.7")
	registerFunc(&_cMTimeRangeGetIntersection, &_cMTimeRangeGetIntersectionErr, frameworkHandle, "CMTimeRangeGetIntersection", "10.7")
	registerFunc(&_cMTimeRangeGetUnion, &_cMTimeRangeGetUnionErr, frameworkHandle, "CMTimeRangeGetUnion", "10.7")
	registerFunc(&_cMTimeRangeMake, &_cMTimeRangeMakeErr, frameworkHandle, "CMTimeRangeMake", "10.7")
	registerFunc(&_cMTimeRangeMakeFromDictionary, &_cMTimeRangeMakeFromDictionaryErr, frameworkHandle, "CMTimeRangeMakeFromDictionary", "10.7")
	registerFunc(&_cMTimeRangeShow, &_cMTimeRangeShowErr, frameworkHandle, "CMTimeRangeShow", "10.7")
	registerFunc(&_cMTimeShow, &_cMTimeShowErr, frameworkHandle, "CMTimeShow", "10.7")
	registerFunc(&_cMTimeSubtract, &_cMTimeSubtractErr, frameworkHandle, "CMTimeSubtract", "10.7")
	registerFunc(&_cMTimebaseAddTimer, &_cMTimebaseAddTimerErr, frameworkHandle, "CMTimebaseAddTimer", "10.8")
	registerFunc(&_cMTimebaseAddTimerDispatchSource, &_cMTimebaseAddTimerDispatchSourceErr, frameworkHandle, "CMTimebaseAddTimerDispatchSource", "10.8")
	registerFunc(&_cMTimebaseCopySource, &_cMTimebaseCopySourceErr, frameworkHandle, "CMTimebaseCopySource", "10.11")
	registerFunc(&_cMTimebaseCopySourceClock, &_cMTimebaseCopySourceClockErr, frameworkHandle, "CMTimebaseCopySourceClock", "10.11")
	registerFunc(&_cMTimebaseCopySourceTimebase, &_cMTimebaseCopySourceTimebaseErr, frameworkHandle, "CMTimebaseCopySourceTimebase", "10.11")
	registerFunc(&_cMTimebaseCopyUltimateSourceClock, &_cMTimebaseCopyUltimateSourceClockErr, frameworkHandle, "CMTimebaseCopyUltimateSourceClock", "10.11")
	registerFunc(&_cMTimebaseCreateWithSourceClock, &_cMTimebaseCreateWithSourceClockErr, frameworkHandle, "CMTimebaseCreateWithSourceClock", "10.8")
	registerFunc(&_cMTimebaseCreateWithSourceTimebase, &_cMTimebaseCreateWithSourceTimebaseErr, frameworkHandle, "CMTimebaseCreateWithSourceTimebase", "10.8")
	registerFunc(&_cMTimebaseGetEffectiveRate, &_cMTimebaseGetEffectiveRateErr, frameworkHandle, "CMTimebaseGetEffectiveRate", "10.8")
	registerFunc(&_cMTimebaseGetRate, &_cMTimebaseGetRateErr, frameworkHandle, "CMTimebaseGetRate", "10.8")
	registerFunc(&_cMTimebaseGetTime, &_cMTimebaseGetTimeErr, frameworkHandle, "CMTimebaseGetTime", "10.8")
	registerFunc(&_cMTimebaseGetTimeAndRate, &_cMTimebaseGetTimeAndRateErr, frameworkHandle, "CMTimebaseGetTimeAndRate", "10.8")
	registerFunc(&_cMTimebaseGetTimeWithTimeScale, &_cMTimebaseGetTimeWithTimeScaleErr, frameworkHandle, "CMTimebaseGetTimeWithTimeScale", "10.8")
	registerFunc(&_cMTimebaseGetTypeID, &_cMTimebaseGetTypeIDErr, frameworkHandle, "CMTimebaseGetTypeID", "10.8")
	registerFunc(&_cMTimebaseNotificationBarrier, &_cMTimebaseNotificationBarrierErr, frameworkHandle, "CMTimebaseNotificationBarrier", "10.8")
	registerFunc(&_cMTimebaseRemoveTimer, &_cMTimebaseRemoveTimerErr, frameworkHandle, "CMTimebaseRemoveTimer", "10.8")
	registerFunc(&_cMTimebaseRemoveTimerDispatchSource, &_cMTimebaseRemoveTimerDispatchSourceErr, frameworkHandle, "CMTimebaseRemoveTimerDispatchSource", "10.8")
	registerFunc(&_cMTimebaseSetAnchorTime, &_cMTimebaseSetAnchorTimeErr, frameworkHandle, "CMTimebaseSetAnchorTime", "10.8")
	registerFunc(&_cMTimebaseSetRate, &_cMTimebaseSetRateErr, frameworkHandle, "CMTimebaseSetRate", "10.8")
	registerFunc(&_cMTimebaseSetRateAndAnchorTime, &_cMTimebaseSetRateAndAnchorTimeErr, frameworkHandle, "CMTimebaseSetRateAndAnchorTime", "10.8")
	registerFunc(&_cMTimebaseSetSourceClock, &_cMTimebaseSetSourceClockErr, frameworkHandle, "CMTimebaseSetSourceClock", "10.8")
	registerFunc(&_cMTimebaseSetSourceTimebase, &_cMTimebaseSetSourceTimebaseErr, frameworkHandle, "CMTimebaseSetSourceTimebase", "10.8")
	registerFunc(&_cMTimebaseSetTime, &_cMTimebaseSetTimeErr, frameworkHandle, "CMTimebaseSetTime", "10.8")
	registerFunc(&_cMTimebaseSetTimerDispatchSourceNextFireTime, &_cMTimebaseSetTimerDispatchSourceNextFireTimeErr, frameworkHandle, "CMTimebaseSetTimerDispatchSourceNextFireTime", "10.8")
	registerFunc(&_cMTimebaseSetTimerDispatchSourceToFireImmediately, &_cMTimebaseSetTimerDispatchSourceToFireImmediatelyErr, frameworkHandle, "CMTimebaseSetTimerDispatchSourceToFireImmediately", "10.8")
	registerFunc(&_cMTimebaseSetTimerNextFireTime, &_cMTimebaseSetTimerNextFireTimeErr, frameworkHandle, "CMTimebaseSetTimerNextFireTime", "10.8")
	registerFunc(&_cMTimebaseSetTimerToFireImmediately, &_cMTimebaseSetTimerToFireImmediatelyErr, frameworkHandle, "CMTimebaseSetTimerToFireImmediately", "10.8")
	registerFunc(&_cMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer, &_cMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBufferErr, frameworkHandle, "CMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer", "10.10")
	registerFunc(&_cMVideoFormatDescriptionCopyTagCollectionArray, &_cMVideoFormatDescriptionCopyTagCollectionArrayErr, frameworkHandle, "CMVideoFormatDescriptionCopyTagCollectionArray", "14.0")
	registerFunc(&_cMVideoFormatDescriptionCreate, &_cMVideoFormatDescriptionCreateErr, frameworkHandle, "CMVideoFormatDescriptionCreate", "10.7")
	registerFunc(&_cMVideoFormatDescriptionCreateForImageBuffer, &_cMVideoFormatDescriptionCreateForImageBufferErr, frameworkHandle, "CMVideoFormatDescriptionCreateForImageBuffer", "10.7")
	registerFunc(&_cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer, &_cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBufferErr, frameworkHandle, "CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer", "10.10")
	registerFunc(&_cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData, &_cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionDataErr, frameworkHandle, "CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData", "10.10")
	registerFunc(&_cMVideoFormatDescriptionCreateFromH264ParameterSets, &_cMVideoFormatDescriptionCreateFromH264ParameterSetsErr, frameworkHandle, "CMVideoFormatDescriptionCreateFromH264ParameterSets", "10.9")
	registerFunc(&_cMVideoFormatDescriptionCreateFromHEVCParameterSets, &_cMVideoFormatDescriptionCreateFromHEVCParameterSetsErr, frameworkHandle, "CMVideoFormatDescriptionCreateFromHEVCParameterSets", "10.13")
	registerFunc(&_cMVideoFormatDescriptionGetCleanAperture, &_cMVideoFormatDescriptionGetCleanApertureErr, frameworkHandle, "CMVideoFormatDescriptionGetCleanAperture", "10.7")
	registerFunc(&_cMVideoFormatDescriptionGetDimensions, &_cMVideoFormatDescriptionGetDimensionsErr, frameworkHandle, "CMVideoFormatDescriptionGetDimensions", "10.7")
	registerFunc(&_cMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers, &_cMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffersErr, frameworkHandle, "CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers", "10.7")
	registerFunc(&_cMVideoFormatDescriptionGetH264ParameterSetAtIndex, &_cMVideoFormatDescriptionGetH264ParameterSetAtIndexErr, frameworkHandle, "CMVideoFormatDescriptionGetH264ParameterSetAtIndex", "10.9")
	registerFunc(&_cMVideoFormatDescriptionGetHEVCParameterSetAtIndex, &_cMVideoFormatDescriptionGetHEVCParameterSetAtIndexErr, frameworkHandle, "CMVideoFormatDescriptionGetHEVCParameterSetAtIndex", "10.13")
	registerFunc(&_cMVideoFormatDescriptionGetPresentationDimensions, &_cMVideoFormatDescriptionGetPresentationDimensionsErr, frameworkHandle, "CMVideoFormatDescriptionGetPresentationDimensions", "10.7")
	registerFunc(&_cMVideoFormatDescriptionMatchesImageBuffer, &_cMVideoFormatDescriptionMatchesImageBufferErr, frameworkHandle, "CMVideoFormatDescriptionMatchesImageBuffer", "10.7")
}
