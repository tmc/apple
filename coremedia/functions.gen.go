// Code generated from Apple documentation for CoreMedia. DO NOT EDIT.

package coremedia

import (
	"fmt"
	"os"
	"unsafe"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objectivec"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: CoreMedia: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: CoreMedia: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: CoreMedia: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _cMAudioDeviceClockCreate func(allocator corefoundation.CFAllocatorRef, deviceUID corefoundation.CFStringRef, clockOut *uintptr) int32

// CMAudioDeviceClockCreate creates a clock that tracks playback through a Core Audio device with the specified unique identifier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioDeviceClockCreate(allocator:deviceUID:clockOut:)
func CMAudioDeviceClockCreate(allocator corefoundation.CFAllocatorRef, deviceUID corefoundation.CFStringRef, clockOut *uintptr) int32 {
	if _cMAudioDeviceClockCreate == nil {
		panic("CoreMedia: symbol CMAudioDeviceClockCreate not loaded")
	}
	return _cMAudioDeviceClockCreate(allocator, deviceUID, clockOut)
}

var _cMAudioDeviceClockCreateFromAudioDeviceID func(allocator corefoundation.CFAllocatorRef, deviceID uint32, clockOut *uintptr) int32

// CMAudioDeviceClockCreateFromAudioDeviceID creates a clock that tracks playback through a Core Audio device with the specified identifier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioDeviceClockCreateFromAudioDeviceID(allocator:deviceID:clockOut:)
func CMAudioDeviceClockCreateFromAudioDeviceID(allocator corefoundation.CFAllocatorRef, deviceID uint32, clockOut *uintptr) int32 {
	if _cMAudioDeviceClockCreateFromAudioDeviceID == nil {
		panic("CoreMedia: symbol CMAudioDeviceClockCreateFromAudioDeviceID not loaded")
	}
	return _cMAudioDeviceClockCreateFromAudioDeviceID(allocator, deviceID, clockOut)
}

var _cMAudioDeviceClockGetAudioDevice func(clock uintptr, deviceUIDOut *corefoundation.CFStringRef, deviceIDOut *uint32, trackingDefaultDeviceOut *bool) int32

// CMAudioDeviceClockGetAudioDevice returns the Core Audio device the clock is tracking.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioDeviceClockGetAudioDevice(_:deviceUIDOut:deviceIDOut:trackingDefaultDeviceOut:)
func CMAudioDeviceClockGetAudioDevice(clock uintptr, deviceUIDOut *corefoundation.CFStringRef, deviceIDOut *uint32, trackingDefaultDeviceOut *bool) int32 {
	if _cMAudioDeviceClockGetAudioDevice == nil {
		panic("CoreMedia: symbol CMAudioDeviceClockGetAudioDevice not loaded")
	}
	return _cMAudioDeviceClockGetAudioDevice(clock, deviceUIDOut, deviceIDOut, trackingDefaultDeviceOut)
}

var _cMAudioDeviceClockSetAudioDeviceID func(clock uintptr, deviceID uint32) int32

// CMAudioDeviceClockSetAudioDeviceID changes the Core Audio device the clock is tracking by specifying a new device identifier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioDeviceClockSetAudioDeviceID(_:deviceID:)
func CMAudioDeviceClockSetAudioDeviceID(clock uintptr, deviceID uint32) int32 {
	if _cMAudioDeviceClockSetAudioDeviceID == nil {
		panic("CoreMedia: symbol CMAudioDeviceClockSetAudioDeviceID not loaded")
	}
	return _cMAudioDeviceClockSetAudioDeviceID(clock, deviceID)
}

var _cMAudioDeviceClockSetAudioDeviceUID func(clock uintptr, deviceUID corefoundation.CFStringRef) int32

// CMAudioDeviceClockSetAudioDeviceUID changes the Core Audio device the clock is tracking by specifying a new device unique identifier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioDeviceClockSetAudioDeviceUID(_:deviceUID:)
func CMAudioDeviceClockSetAudioDeviceUID(clock uintptr, deviceUID corefoundation.CFStringRef) int32 {
	if _cMAudioDeviceClockSetAudioDeviceUID == nil {
		panic("CoreMedia: symbol CMAudioDeviceClockSetAudioDeviceUID not loaded")
	}
	return _cMAudioDeviceClockSetAudioDeviceUID(clock, deviceUID)
}

var _cMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, audioFormatDescription uintptr, flavor CMSoundDescriptionFlavor, blockBufferOut *uintptr) int32

// CMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer copies the contents of an audio format description to a buffer in big-endian byte ordering.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer(allocator:audioFormatDescription:flavor:blockBufferOut:)
func CMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, audioFormatDescription uintptr, flavor CMSoundDescriptionFlavor, blockBufferOut *uintptr) int32 {
	if _cMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer == nil {
		panic("CoreMedia: symbol CMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer not loaded")
	}
	return _cMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer(allocator, audioFormatDescription, flavor, blockBufferOut)
}

var _cMAudioFormatDescriptionCreate func(allocator corefoundation.CFAllocatorRef, asbd uintptr, layoutSize uintptr, layout uintptr, magicCookieSize uintptr, magicCookie unsafe.Pointer, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *uintptr) int32

// CMAudioFormatDescriptionCreate creates a format description for an audio media stream.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionCreate(allocator:asbd:layoutSize:layout:magicCookieSize:magicCookie:extensions:formatDescriptionOut:)
func CMAudioFormatDescriptionCreate(allocator corefoundation.CFAllocatorRef, asbd uintptr, layoutSize uintptr, layout uintptr, magicCookieSize uintptr, magicCookie unsafe.Pointer, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *uintptr) int32 {
	if _cMAudioFormatDescriptionCreate == nil {
		panic("CoreMedia: symbol CMAudioFormatDescriptionCreate not loaded")
	}
	return _cMAudioFormatDescriptionCreate(allocator, asbd, layoutSize, layout, magicCookieSize, magicCookie, extensions, formatDescriptionOut)
}

var _cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, soundDescriptionBlockBuffer uintptr, flavor CMSoundDescriptionFlavor, formatDescriptionOut *uintptr) int32

// CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer creates an audio format description from a big-endian sound description data structure in a buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer(allocator:bigEndianSoundDescriptionBlockBuffer:flavor:formatDescriptionOut:)
func CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, soundDescriptionBlockBuffer uintptr, flavor CMSoundDescriptionFlavor, formatDescriptionOut *uintptr) int32 {
	if _cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer == nil {
		panic("CoreMedia: symbol CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer not loaded")
	}
	return _cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer(allocator, soundDescriptionBlockBuffer, flavor, formatDescriptionOut)
}

var _cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData func(allocator corefoundation.CFAllocatorRef, soundDescriptionData *uint8, size uintptr, flavor CMSoundDescriptionFlavor, formatDescriptionOut *uintptr) int32

// CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData creates an audio format description from a big-endian sound description data structure.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData(allocator:bigEndianSoundDescriptionData:size:flavor:formatDescriptionOut:)
func CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData(allocator corefoundation.CFAllocatorRef, soundDescriptionData *uint8, size uintptr, flavor CMSoundDescriptionFlavor, formatDescriptionOut *uintptr) int32 {
	if _cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData == nil {
		panic("CoreMedia: symbol CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData not loaded")
	}
	return _cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData(allocator, soundDescriptionData, size, flavor, formatDescriptionOut)
}

var _cMAudioFormatDescriptionCreateSummary func(allocator corefoundation.CFAllocatorRef, formatDescriptionArray corefoundation.CFArrayRef, flags uint32, formatDescriptionOut *uintptr) int32

// CMAudioFormatDescriptionCreateSummary creates a summary audio format description from an array of descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionCreateSummary(allocator:formatDescriptionArray:flags:formatDescriptionOut:)
func CMAudioFormatDescriptionCreateSummary(allocator corefoundation.CFAllocatorRef, formatDescriptionArray corefoundation.CFArrayRef, flags uint32, formatDescriptionOut *uintptr) int32 {
	if _cMAudioFormatDescriptionCreateSummary == nil {
		panic("CoreMedia: symbol CMAudioFormatDescriptionCreateSummary not loaded")
	}
	return _cMAudioFormatDescriptionCreateSummary(allocator, formatDescriptionArray, flags, formatDescriptionOut)
}

var _cMAudioFormatDescriptionEqual func(formatDescription uintptr, otherFormatDescription uintptr, equalityMask CMAudioFormatDescriptionMask, equalityMaskOut *CMAudioFormatDescriptionMask) bool

// CMAudioFormatDescriptionEqual returns a Boolean value that indicates whether the two audio format descriptions are equal.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionEqual(_:otherFormatDescription:equalityMask:equalityMaskOut:)
func CMAudioFormatDescriptionEqual(formatDescription uintptr, otherFormatDescription uintptr, equalityMask CMAudioFormatDescriptionMask, equalityMaskOut *CMAudioFormatDescriptionMask) bool {
	if _cMAudioFormatDescriptionEqual == nil {
		panic("CoreMedia: symbol CMAudioFormatDescriptionEqual not loaded")
	}
	return _cMAudioFormatDescriptionEqual(formatDescription, otherFormatDescription, equalityMask, equalityMaskOut)
}

var _cMAudioFormatDescriptionGetChannelLayout func(desc uintptr, sizeOut *uintptr) *objectivec.IObject

// CMAudioFormatDescriptionGetChannelLayout returns a read-only pointer to, and the size of, the audio channel layout inside an audio format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionGetChannelLayout(_:sizeOut:)
func CMAudioFormatDescriptionGetChannelLayout(desc uintptr, sizeOut *uintptr) *objectivec.IObject {
	if _cMAudioFormatDescriptionGetChannelLayout == nil {
		panic("CoreMedia: symbol CMAudioFormatDescriptionGetChannelLayout not loaded")
	}
	return _cMAudioFormatDescriptionGetChannelLayout(desc, sizeOut)
}

var _cMAudioFormatDescriptionGetFormatList func(desc uintptr, sizeOut *uintptr) *objectivec.IObject

// CMAudioFormatDescriptionGetFormatList returns a read-only pointer to, and size of, the array of audio format list item structures in an audio format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionGetFormatList(_:sizeOut:)
func CMAudioFormatDescriptionGetFormatList(desc uintptr, sizeOut *uintptr) *objectivec.IObject {
	if _cMAudioFormatDescriptionGetFormatList == nil {
		panic("CoreMedia: symbol CMAudioFormatDescriptionGetFormatList not loaded")
	}
	return _cMAudioFormatDescriptionGetFormatList(desc, sizeOut)
}

var _cMAudioFormatDescriptionGetMagicCookie func(desc uintptr, sizeOut *uintptr) unsafe.Pointer

// CMAudioFormatDescriptionGetMagicCookie returns a read-only pointer to, and size of, the magic cookie in an audio format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionGetMagicCookie(_:sizeOut:)
func CMAudioFormatDescriptionGetMagicCookie(desc uintptr, sizeOut *uintptr) unsafe.Pointer {
	if _cMAudioFormatDescriptionGetMagicCookie == nil {
		panic("CoreMedia: symbol CMAudioFormatDescriptionGetMagicCookie not loaded")
	}
	return _cMAudioFormatDescriptionGetMagicCookie(desc, sizeOut)
}

var _cMAudioFormatDescriptionGetMostCompatibleFormat func(desc uintptr) *objectivec.IObject

// CMAudioFormatDescriptionGetMostCompatibleFormat returns a read-only pointer to the appropriate audio format list item in an audio format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionGetMostCompatibleFormat(_:)
func CMAudioFormatDescriptionGetMostCompatibleFormat(desc uintptr) *objectivec.IObject {
	if _cMAudioFormatDescriptionGetMostCompatibleFormat == nil {
		panic("CoreMedia: symbol CMAudioFormatDescriptionGetMostCompatibleFormat not loaded")
	}
	return _cMAudioFormatDescriptionGetMostCompatibleFormat(desc)
}

var _cMAudioFormatDescriptionGetRichestDecodableFormat func(desc uintptr) *objectivec.IObject

// CMAudioFormatDescriptionGetRichestDecodableFormat returns a read-only pointer to the appropriate audio format list item in an audio format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionGetRichestDecodableFormat(_:)
func CMAudioFormatDescriptionGetRichestDecodableFormat(desc uintptr) *objectivec.IObject {
	if _cMAudioFormatDescriptionGetRichestDecodableFormat == nil {
		panic("CoreMedia: symbol CMAudioFormatDescriptionGetRichestDecodableFormat not loaded")
	}
	return _cMAudioFormatDescriptionGetRichestDecodableFormat(desc)
}

var _cMAudioFormatDescriptionGetStreamBasicDescription func(desc uintptr) *objectivec.IObject

// CMAudioFormatDescriptionGetStreamBasicDescription returns a read-only pointer to the audio stream description in an audio format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioFormatDescriptionGetStreamBasicDescription(_:)
func CMAudioFormatDescriptionGetStreamBasicDescription(desc uintptr) *objectivec.IObject {
	if _cMAudioFormatDescriptionGetStreamBasicDescription == nil {
		panic("CoreMedia: symbol CMAudioFormatDescriptionGetStreamBasicDescription not loaded")
	}
	return _cMAudioFormatDescriptionGetStreamBasicDescription(desc)
}

var _cMAudioSampleBufferCreateReadyWithPacketDescriptions func(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, formatDescription uintptr, numSamples int, presentationTimeStamp CMTime, packetDescriptions uintptr, sampleBufferOut *uintptr) int32

// CMAudioSampleBufferCreateReadyWithPacketDescriptions creates a sample buffer with packet descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioSampleBufferCreateReadyWithPacketDescriptions(allocator:dataBuffer:formatDescription:sampleCount:presentationTimeStamp:packetDescriptions:sampleBufferOut:)
func CMAudioSampleBufferCreateReadyWithPacketDescriptions(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, formatDescription uintptr, numSamples int, presentationTimeStamp CMTime, packetDescriptions uintptr, sampleBufferOut *uintptr) int32 {
	if _cMAudioSampleBufferCreateReadyWithPacketDescriptions == nil {
		panic("CoreMedia: symbol CMAudioSampleBufferCreateReadyWithPacketDescriptions not loaded")
	}
	return _cMAudioSampleBufferCreateReadyWithPacketDescriptions(allocator, dataBuffer, formatDescription, numSamples, presentationTimeStamp, packetDescriptions, sampleBufferOut)
}

var _cMAudioSampleBufferCreateWithPacketDescriptions func(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, dataReady bool, makeDataReadyCallback CMSampleBufferMakeDataReadyCallback, makeDataReadyRefcon unsafe.Pointer, formatDescription uintptr, numSamples int, presentationTimeStamp CMTime, packetDescriptions uintptr, sampleBufferOut *uintptr) int32

// CMAudioSampleBufferCreateWithPacketDescriptions creates a sample buffer with packet descriptions and a callback to make the data ready for use.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioSampleBufferCreateWithPacketDescriptions(allocator:dataBuffer:dataReady:makeDataReadyCallback:refcon:formatDescription:sampleCount:presentationTimeStamp:packetDescriptions:sampleBufferOut:)
func CMAudioSampleBufferCreateWithPacketDescriptions(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, dataReady bool, makeDataReadyCallback CMSampleBufferMakeDataReadyCallback, makeDataReadyRefcon unsafe.Pointer, formatDescription uintptr, numSamples int, presentationTimeStamp CMTime, packetDescriptions uintptr, sampleBufferOut *uintptr) int32 {
	if _cMAudioSampleBufferCreateWithPacketDescriptions == nil {
		panic("CoreMedia: symbol CMAudioSampleBufferCreateWithPacketDescriptions not loaded")
	}
	return _cMAudioSampleBufferCreateWithPacketDescriptions(allocator, dataBuffer, dataReady, makeDataReadyCallback, makeDataReadyRefcon, formatDescription, numSamples, presentationTimeStamp, packetDescriptions, sampleBufferOut)
}

var _cMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler func(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, dataReady bool, formatDescription uintptr, numSamples int, presentationTimeStamp CMTime, packetDescriptions uintptr, sampleBufferOut *uintptr, makeDataReadyHandler CMSampleBufferMakeDataReadyHandler) int32

// CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler creates a sample buffer with packet descriptions and a handler to make the data ready for use.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler(_:_:_:_:_:_:_:_:_:)
func CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, dataReady bool, formatDescription uintptr, numSamples int, presentationTimeStamp CMTime, packetDescriptions uintptr, sampleBufferOut *uintptr, makeDataReadyHandler CMSampleBufferMakeDataReadyHandler) int32 {
	if _cMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler == nil {
		panic("CoreMedia: symbol CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler not loaded")
	}
	return _cMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler(allocator, dataBuffer, dataReady, formatDescription, numSamples, presentationTimeStamp, packetDescriptions, sampleBufferOut, makeDataReadyHandler)
}

var _cMBlockBufferAccessDataBytes func(theBuffer uintptr, offset uintptr, length uintptr, temporaryBlock unsafe.Pointer, returnedPointerOut string) int32

// CMBlockBufferAccessDataBytes accesses potentially noncontiguous data in a block buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferAccessDataBytes(_:atOffset:length:temporaryBlock:returnedPointerOut:)
func CMBlockBufferAccessDataBytes(theBuffer uintptr, offset uintptr, length uintptr, temporaryBlock unsafe.Pointer, returnedPointerOut string) int32 {
	if _cMBlockBufferAccessDataBytes == nil {
		panic("CoreMedia: symbol CMBlockBufferAccessDataBytes not loaded")
	}
	return _cMBlockBufferAccessDataBytes(theBuffer, offset, length, temporaryBlock, returnedPointerOut)
}

var _cMBlockBufferAppendBufferReference func(theBuffer uintptr, targetBBuf uintptr, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags) int32

// CMBlockBufferAppendBufferReference adds a reference to an existing block buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferAppendBufferReference(_:targetBBuf:offsetToData:dataLength:flags:)
func CMBlockBufferAppendBufferReference(theBuffer uintptr, targetBBuf uintptr, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags) int32 {
	if _cMBlockBufferAppendBufferReference == nil {
		panic("CoreMedia: symbol CMBlockBufferAppendBufferReference not loaded")
	}
	return _cMBlockBufferAppendBufferReference(theBuffer, targetBBuf, offsetToData, dataLength, flags)
}

var _cMBlockBufferAppendMemoryBlock func(theBuffer uintptr, memoryBlock unsafe.Pointer, blockLength uintptr, blockAllocator corefoundation.CFAllocatorRef, customBlockSource *CMBlockBufferCustomBlockSource, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags) int32

// CMBlockBufferAppendMemoryBlock adds a memory block to an existing block buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferAppendMemoryBlock(_:memoryBlock:length:blockAllocator:customBlockSource:offsetToData:dataLength:flags:)
func CMBlockBufferAppendMemoryBlock(theBuffer uintptr, memoryBlock unsafe.Pointer, blockLength uintptr, blockAllocator corefoundation.CFAllocatorRef, customBlockSource *CMBlockBufferCustomBlockSource, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags) int32 {
	if _cMBlockBufferAppendMemoryBlock == nil {
		panic("CoreMedia: symbol CMBlockBufferAppendMemoryBlock not loaded")
	}
	return _cMBlockBufferAppendMemoryBlock(theBuffer, memoryBlock, blockLength, blockAllocator, customBlockSource, offsetToData, dataLength, flags)
}

var _cMBlockBufferAssureBlockMemory func(theBuffer uintptr) int32

// CMBlockBufferAssureBlockMemory assures that the system allocates memory for all memory blocks in a block buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferAssureBlockMemory(_:)
func CMBlockBufferAssureBlockMemory(theBuffer uintptr) int32 {
	if _cMBlockBufferAssureBlockMemory == nil {
		panic("CoreMedia: symbol CMBlockBufferAssureBlockMemory not loaded")
	}
	return _cMBlockBufferAssureBlockMemory(theBuffer)
}

var _cMBlockBufferCopyDataBytes func(theSourceBuffer uintptr, offsetToData uintptr, dataLength uintptr, destination unsafe.Pointer) int32

// CMBlockBufferCopyDataBytes copies bytes from a block buffer into a provided memory area.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCopyDataBytes(_:atOffset:dataLength:destination:)
func CMBlockBufferCopyDataBytes(theSourceBuffer uintptr, offsetToData uintptr, dataLength uintptr, destination unsafe.Pointer) int32 {
	if _cMBlockBufferCopyDataBytes == nil {
		panic("CoreMedia: symbol CMBlockBufferCopyDataBytes not loaded")
	}
	return _cMBlockBufferCopyDataBytes(theSourceBuffer, offsetToData, dataLength, destination)
}

var _cMBlockBufferCreateContiguous func(structureAllocator corefoundation.CFAllocatorRef, sourceBuffer uintptr, blockAllocator corefoundation.CFAllocatorRef, customBlockSource *CMBlockBufferCustomBlockSource, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags, blockBufferOut *uintptr) int32

// CMBlockBufferCreateContiguous creates a block buffer that contains a contiguous copy of, or reference to, the data specified by the parameters.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCreateContiguous(allocator:sourceBuffer:blockAllocator:customBlockSource:offsetToData:dataLength:flags:blockBufferOut:)
func CMBlockBufferCreateContiguous(structureAllocator corefoundation.CFAllocatorRef, sourceBuffer uintptr, blockAllocator corefoundation.CFAllocatorRef, customBlockSource *CMBlockBufferCustomBlockSource, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags, blockBufferOut *uintptr) int32 {
	if _cMBlockBufferCreateContiguous == nil {
		panic("CoreMedia: symbol CMBlockBufferCreateContiguous not loaded")
	}
	return _cMBlockBufferCreateContiguous(structureAllocator, sourceBuffer, blockAllocator, customBlockSource, offsetToData, dataLength, flags, blockBufferOut)
}

var _cMBlockBufferCreateEmpty func(structureAllocator corefoundation.CFAllocatorRef, subBlockCapacity uint32, flags CMBlockBufferFlags, blockBufferOut *uintptr) int32

// CMBlockBufferCreateEmpty creates an empty block buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCreateEmpty(allocator:capacity:flags:blockBufferOut:)
func CMBlockBufferCreateEmpty(structureAllocator corefoundation.CFAllocatorRef, subBlockCapacity uint32, flags CMBlockBufferFlags, blockBufferOut *uintptr) int32 {
	if _cMBlockBufferCreateEmpty == nil {
		panic("CoreMedia: symbol CMBlockBufferCreateEmpty not loaded")
	}
	return _cMBlockBufferCreateEmpty(structureAllocator, subBlockCapacity, flags, blockBufferOut)
}

var _cMBlockBufferCreateWithBufferReference func(structureAllocator corefoundation.CFAllocatorRef, bufferReference uintptr, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags, blockBufferOut *uintptr) int32

// CMBlockBufferCreateWithBufferReference creates a block buffer that refers to another block buffer object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCreateWithBufferReference(allocator:referenceBuffer:offsetToData:dataLength:flags:blockBufferOut:)
func CMBlockBufferCreateWithBufferReference(structureAllocator corefoundation.CFAllocatorRef, bufferReference uintptr, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags, blockBufferOut *uintptr) int32 {
	if _cMBlockBufferCreateWithBufferReference == nil {
		panic("CoreMedia: symbol CMBlockBufferCreateWithBufferReference not loaded")
	}
	return _cMBlockBufferCreateWithBufferReference(structureAllocator, bufferReference, offsetToData, dataLength, flags, blockBufferOut)
}

var _cMBlockBufferCreateWithMemoryBlock func(structureAllocator corefoundation.CFAllocatorRef, memoryBlock unsafe.Pointer, blockLength uintptr, blockAllocator corefoundation.CFAllocatorRef, customBlockSource *CMBlockBufferCustomBlockSource, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags, blockBufferOut *uintptr) int32

// CMBlockBufferCreateWithMemoryBlock creates a block buffer that’s backed by a memory block.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferCreateWithMemoryBlock(allocator:memoryBlock:blockLength:blockAllocator:customBlockSource:offsetToData:dataLength:flags:blockBufferOut:)
func CMBlockBufferCreateWithMemoryBlock(structureAllocator corefoundation.CFAllocatorRef, memoryBlock unsafe.Pointer, blockLength uintptr, blockAllocator corefoundation.CFAllocatorRef, customBlockSource *CMBlockBufferCustomBlockSource, offsetToData uintptr, dataLength uintptr, flags CMBlockBufferFlags, blockBufferOut *uintptr) int32 {
	if _cMBlockBufferCreateWithMemoryBlock == nil {
		panic("CoreMedia: symbol CMBlockBufferCreateWithMemoryBlock not loaded")
	}
	return _cMBlockBufferCreateWithMemoryBlock(structureAllocator, memoryBlock, blockLength, blockAllocator, customBlockSource, offsetToData, dataLength, flags, blockBufferOut)
}

var _cMBlockBufferFillDataBytes func(fillByte int8, destinationBuffer uintptr, offsetIntoDestination uintptr, dataLength uintptr) int32

// CMBlockBufferFillDataBytes fills the destination buffer with the specified data byte.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferFillDataBytes(with:blockBuffer:offsetIntoDestination:dataLength:)
func CMBlockBufferFillDataBytes(fillByte int8, destinationBuffer uintptr, offsetIntoDestination uintptr, dataLength uintptr) int32 {
	if _cMBlockBufferFillDataBytes == nil {
		panic("CoreMedia: symbol CMBlockBufferFillDataBytes not loaded")
	}
	return _cMBlockBufferFillDataBytes(fillByte, destinationBuffer, offsetIntoDestination, dataLength)
}

var _cMBlockBufferGetDataLength func(theBuffer uintptr) uintptr

// CMBlockBufferGetDataLength returns the total length of data that’s accessible by a block buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferGetDataLength(_:)
func CMBlockBufferGetDataLength(theBuffer uintptr) uintptr {
	if _cMBlockBufferGetDataLength == nil {
		panic("CoreMedia: symbol CMBlockBufferGetDataLength not loaded")
	}
	return _cMBlockBufferGetDataLength(theBuffer)
}

var _cMBlockBufferGetDataPointer func(theBuffer uintptr, offset uintptr, lengthAtOffsetOut *uintptr, totalLengthOut *uintptr, dataPointerOut string) int32

// CMBlockBufferGetDataPointer gains access to the data represented by a block buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferGetDataPointer(_:atOffset:lengthAtOffsetOut:totalLengthOut:dataPointerOut:)
func CMBlockBufferGetDataPointer(theBuffer uintptr, offset uintptr, lengthAtOffsetOut *uintptr, totalLengthOut *uintptr, dataPointerOut string) int32 {
	if _cMBlockBufferGetDataPointer == nil {
		panic("CoreMedia: symbol CMBlockBufferGetDataPointer not loaded")
	}
	return _cMBlockBufferGetDataPointer(theBuffer, offset, lengthAtOffsetOut, totalLengthOut, dataPointerOut)
}

var _cMBlockBufferGetTypeID func() uint

// CMBlockBufferGetTypeID returns the type identifier for block buffer objects.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferGetTypeID()
func CMBlockBufferGetTypeID() uint {
	if _cMBlockBufferGetTypeID == nil {
		panic("CoreMedia: symbol CMBlockBufferGetTypeID not loaded")
	}
	return _cMBlockBufferGetTypeID()
}

var _cMBlockBufferIsEmpty func(theBuffer uintptr) bool

// CMBlockBufferIsEmpty returns a Boolean value that indicates whether the buffer is empty.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferIsEmpty(_:)
func CMBlockBufferIsEmpty(theBuffer uintptr) bool {
	if _cMBlockBufferIsEmpty == nil {
		panic("CoreMedia: symbol CMBlockBufferIsEmpty not loaded")
	}
	return _cMBlockBufferIsEmpty(theBuffer)
}

var _cMBlockBufferIsRangeContiguous func(theBuffer uintptr, offset uintptr, length uintptr) bool

// CMBlockBufferIsRangeContiguous returns a Boolean value that indicates whether the specified range within a block buffer is contiguous.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferIsRangeContiguous(_:atOffset:length:)
func CMBlockBufferIsRangeContiguous(theBuffer uintptr, offset uintptr, length uintptr) bool {
	if _cMBlockBufferIsRangeContiguous == nil {
		panic("CoreMedia: symbol CMBlockBufferIsRangeContiguous not loaded")
	}
	return _cMBlockBufferIsRangeContiguous(theBuffer, offset, length)
}

var _cMBlockBufferReplaceDataBytes func(sourceBytes unsafe.Pointer, destinationBuffer uintptr, offsetIntoDestination uintptr, dataLength uintptr) int32

// CMBlockBufferReplaceDataBytes copies bytes from a given memory block into a block buffer replacing bytes in the underlying data blocks.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBlockBufferReplaceDataBytes(with:blockBuffer:offsetIntoDestination:dataLength:)
func CMBlockBufferReplaceDataBytes(sourceBytes unsafe.Pointer, destinationBuffer uintptr, offsetIntoDestination uintptr, dataLength uintptr) int32 {
	if _cMBlockBufferReplaceDataBytes == nil {
		panic("CoreMedia: symbol CMBlockBufferReplaceDataBytes not loaded")
	}
	return _cMBlockBufferReplaceDataBytes(sourceBytes, destinationBuffer, offsetIntoDestination, dataLength)
}

var _cMBufferQueueCallForEachBuffer func(queue uintptr, refcon uintptr) int32

// CMBufferQueueCallForEachBuffer calls a function for every buffer in a queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueCallForEachBuffer(_:callback:refcon:)
func CMBufferQueueCallForEachBuffer(queue uintptr, refcon uintptr) int32 {
	if _cMBufferQueueCallForEachBuffer == nil {
		panic("CoreMedia: symbol CMBufferQueueCallForEachBuffer not loaded")
	}
	return _cMBufferQueueCallForEachBuffer(queue, refcon)
}

var _cMBufferQueueContainsEndOfData func(queue uintptr) bool

// CMBufferQueueContainsEndOfData returns a Boolean value that indicates whether a buffer queue has its end-of-data marker set.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueContainsEndOfData(_:)
func CMBufferQueueContainsEndOfData(queue uintptr) bool {
	if _cMBufferQueueContainsEndOfData == nil {
		panic("CoreMedia: symbol CMBufferQueueContainsEndOfData not loaded")
	}
	return _cMBufferQueueContainsEndOfData(queue)
}

var _cMBufferQueueCopyHead func(queue uintptr) CMBufferRef

// CMBufferQueueCopyHead.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueCopyHead(_:)
func CMBufferQueueCopyHead(queue uintptr) CMBufferRef {
	if _cMBufferQueueCopyHead == nil {
		panic("CoreMedia: symbol CMBufferQueueCopyHead not loaded")
	}
	return _cMBufferQueueCopyHead(queue)
}

var _cMBufferQueueCreate func(allocator corefoundation.CFAllocatorRef, capacity int, callbacks *CMBufferCallbacks, queueOut *uintptr) int32

// CMBufferQueueCreate creates a buffer queue with callbacks to inspect buffers.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueCreate(allocator:capacity:callbacks:queueOut:)
func CMBufferQueueCreate(allocator corefoundation.CFAllocatorRef, capacity int, callbacks *CMBufferCallbacks, queueOut *uintptr) int32 {
	if _cMBufferQueueCreate == nil {
		panic("CoreMedia: symbol CMBufferQueueCreate not loaded")
	}
	return _cMBufferQueueCreate(allocator, capacity, callbacks, queueOut)
}

var _cMBufferQueueCreateWithHandlers func(allocator corefoundation.CFAllocatorRef, capacity int, handlers *CMBufferHandlers, queueOut *uintptr) int32

// CMBufferQueueCreateWithHandlers creates a buffer queue with handlers to inspect buffers.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueCreateWithHandlers(_:_:_:_:)
func CMBufferQueueCreateWithHandlers(allocator corefoundation.CFAllocatorRef, capacity int, handlers *CMBufferHandlers, queueOut *uintptr) int32 {
	if _cMBufferQueueCreateWithHandlers == nil {
		panic("CoreMedia: symbol CMBufferQueueCreateWithHandlers not loaded")
	}
	return _cMBufferQueueCreateWithHandlers(allocator, capacity, handlers, queueOut)
}

var _cMBufferQueueDequeueAndRetain func(queue uintptr) CMBufferRef

// CMBufferQueueDequeueAndRetain dequeues a buffer from a queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueDequeue(_:)
func CMBufferQueueDequeueAndRetain(queue uintptr) CMBufferRef {
	if _cMBufferQueueDequeueAndRetain == nil {
		panic("CoreMedia: symbol CMBufferQueueDequeueAndRetain not loaded")
	}
	return _cMBufferQueueDequeueAndRetain(queue)
}

var _cMBufferQueueDequeueIfDataReadyAndRetain func(queue uintptr) CMBufferRef

// CMBufferQueueDequeueIfDataReadyAndRetain dequeues a buffer from a queue, if it’s ready.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueDequeueIfDataReady(_:)
func CMBufferQueueDequeueIfDataReadyAndRetain(queue uintptr) CMBufferRef {
	if _cMBufferQueueDequeueIfDataReadyAndRetain == nil {
		panic("CoreMedia: symbol CMBufferQueueDequeueIfDataReadyAndRetain not loaded")
	}
	return _cMBufferQueueDequeueIfDataReadyAndRetain(queue)
}

var _cMBufferQueueEnqueue func(queue uintptr, buf CMBufferRef) int32

// CMBufferQueueEnqueue enqueues a buffer onto a queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueEnqueue(_:buffer:)
func CMBufferQueueEnqueue(queue uintptr, buf CMBufferRef) int32 {
	if _cMBufferQueueEnqueue == nil {
		panic("CoreMedia: symbol CMBufferQueueEnqueue not loaded")
	}
	return _cMBufferQueueEnqueue(queue, buf)
}

var _cMBufferQueueGetBufferCount func(queue uintptr) int

// CMBufferQueueGetBufferCount gets the number of buffers in the queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetBufferCount(_:)
func CMBufferQueueGetBufferCount(queue uintptr) int {
	if _cMBufferQueueGetBufferCount == nil {
		panic("CoreMedia: symbol CMBufferQueueGetBufferCount not loaded")
	}
	return _cMBufferQueueGetBufferCount(queue)
}

var _cMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS func() *CMBufferCallbacks

// CMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS returns a pointer to a structure that contains callbacks to sort sample buffers by output presentation timestamp.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS()
func CMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS() *CMBufferCallbacks {
	if _cMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS == nil {
		panic("CoreMedia: symbol CMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS not loaded")
	}
	return _cMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS()
}

var _cMBufferQueueGetCallbacksForUnsortedSampleBuffers func() *CMBufferCallbacks

// CMBufferQueueGetCallbacksForUnsortedSampleBuffers returns a pointer to a callback structure for unsorted sample buffers.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetCallbacksForUnsortedSampleBuffers()
func CMBufferQueueGetCallbacksForUnsortedSampleBuffers() *CMBufferCallbacks {
	if _cMBufferQueueGetCallbacksForUnsortedSampleBuffers == nil {
		panic("CoreMedia: symbol CMBufferQueueGetCallbacksForUnsortedSampleBuffers not loaded")
	}
	return _cMBufferQueueGetCallbacksForUnsortedSampleBuffers()
}

var _cMBufferQueueGetDuration func(queue uintptr) CMTime

// CMBufferQueueGetDuration gets the duration of a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetDuration(_:)
func CMBufferQueueGetDuration(queue uintptr) CMTime {
	if _cMBufferQueueGetDuration == nil {
		panic("CoreMedia: symbol CMBufferQueueGetDuration not loaded")
	}
	return _cMBufferQueueGetDuration(queue)
}

var _cMBufferQueueGetEndPresentationTimeStamp func(queue uintptr) CMTime

// CMBufferQueueGetEndPresentationTimeStamp gets the greatest end presentation timestamp of a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetEndPresentationTimeStamp(_:)
func CMBufferQueueGetEndPresentationTimeStamp(queue uintptr) CMTime {
	if _cMBufferQueueGetEndPresentationTimeStamp == nil {
		panic("CoreMedia: symbol CMBufferQueueGetEndPresentationTimeStamp not loaded")
	}
	return _cMBufferQueueGetEndPresentationTimeStamp(queue)
}

var _cMBufferQueueGetFirstDecodeTimeStamp func(queue uintptr) CMTime

// CMBufferQueueGetFirstDecodeTimeStamp gets the decode timestamp of the first buffer in a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetFirstDecodeTimeStamp(_:)
func CMBufferQueueGetFirstDecodeTimeStamp(queue uintptr) CMTime {
	if _cMBufferQueueGetFirstDecodeTimeStamp == nil {
		panic("CoreMedia: symbol CMBufferQueueGetFirstDecodeTimeStamp not loaded")
	}
	return _cMBufferQueueGetFirstDecodeTimeStamp(queue)
}

var _cMBufferQueueGetFirstPresentationTimeStamp func(queue uintptr) CMTime

// CMBufferQueueGetFirstPresentationTimeStamp gets the presentation timestamp of the first buffer in a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetFirstPresentationTimeStamp(_:)
func CMBufferQueueGetFirstPresentationTimeStamp(queue uintptr) CMTime {
	if _cMBufferQueueGetFirstPresentationTimeStamp == nil {
		panic("CoreMedia: symbol CMBufferQueueGetFirstPresentationTimeStamp not loaded")
	}
	return _cMBufferQueueGetFirstPresentationTimeStamp(queue)
}

var _cMBufferQueueGetHead func(queue uintptr) CMBufferRef

// CMBufferQueueGetHead retrieves the next buffer from a queue, but doesn’t remove it.
//
// Deprecated: Deprecated since macOS 15.0.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetHead(_:)
func CMBufferQueueGetHead(queue uintptr) CMBufferRef {
	if _cMBufferQueueGetHead == nil {
		panic("CoreMedia: symbol CMBufferQueueGetHead not loaded")
	}
	return _cMBufferQueueGetHead(queue)
}

var _cMBufferQueueGetMaxPresentationTimeStamp func(queue uintptr) CMTime

// CMBufferQueueGetMaxPresentationTimeStamp gets the greatest presentation timestamp of a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetMaxPresentationTimeStamp(_:)
func CMBufferQueueGetMaxPresentationTimeStamp(queue uintptr) CMTime {
	if _cMBufferQueueGetMaxPresentationTimeStamp == nil {
		panic("CoreMedia: symbol CMBufferQueueGetMaxPresentationTimeStamp not loaded")
	}
	return _cMBufferQueueGetMaxPresentationTimeStamp(queue)
}

var _cMBufferQueueGetMinDecodeTimeStamp func(queue uintptr) CMTime

// CMBufferQueueGetMinDecodeTimeStamp gets the earliest decode timestamp of a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetMinDecodeTimeStamp(_:)
func CMBufferQueueGetMinDecodeTimeStamp(queue uintptr) CMTime {
	if _cMBufferQueueGetMinDecodeTimeStamp == nil {
		panic("CoreMedia: symbol CMBufferQueueGetMinDecodeTimeStamp not loaded")
	}
	return _cMBufferQueueGetMinDecodeTimeStamp(queue)
}

var _cMBufferQueueGetMinPresentationTimeStamp func(queue uintptr) CMTime

// CMBufferQueueGetMinPresentationTimeStamp gets the earliest presentation timestamp of a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetMinPresentationTimeStamp(_:)
func CMBufferQueueGetMinPresentationTimeStamp(queue uintptr) CMTime {
	if _cMBufferQueueGetMinPresentationTimeStamp == nil {
		panic("CoreMedia: symbol CMBufferQueueGetMinPresentationTimeStamp not loaded")
	}
	return _cMBufferQueueGetMinPresentationTimeStamp(queue)
}

var _cMBufferQueueGetTotalSize func(queue uintptr) uintptr

// CMBufferQueueGetTotalSize gets the total size of all sample buffers of a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetTotalSize(_:)
func CMBufferQueueGetTotalSize(queue uintptr) uintptr {
	if _cMBufferQueueGetTotalSize == nil {
		panic("CoreMedia: symbol CMBufferQueueGetTotalSize not loaded")
	}
	return _cMBufferQueueGetTotalSize(queue)
}

var _cMBufferQueueGetTypeID func() uint

// CMBufferQueueGetTypeID returns the type identifier of buffer queue objects.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueGetTypeID()
func CMBufferQueueGetTypeID() uint {
	if _cMBufferQueueGetTypeID == nil {
		panic("CoreMedia: symbol CMBufferQueueGetTypeID not loaded")
	}
	return _cMBufferQueueGetTypeID()
}

var _cMBufferQueueInstallTrigger func(queue uintptr, callback CMBufferQueueTriggerCallback, refcon uintptr, condition CMBufferQueueTriggerCondition, time CMTime, triggerTokenOut *CMBufferQueueTriggerToken) int32

// CMBufferQueueInstallTrigger installs a trigger with a callback on a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueInstallTrigger(_:callback:refcon:condition:time:triggerTokenOut:)
func CMBufferQueueInstallTrigger(queue uintptr, callback CMBufferQueueTriggerCallback, refcon uintptr, condition CMBufferQueueTriggerCondition, time CMTime, triggerTokenOut *CMBufferQueueTriggerToken) int32 {
	if _cMBufferQueueInstallTrigger == nil {
		panic("CoreMedia: symbol CMBufferQueueInstallTrigger not loaded")
	}
	return _cMBufferQueueInstallTrigger(queue, callback, refcon, condition, time, triggerTokenOut)
}

var _cMBufferQueueInstallTriggerHandler func(queue uintptr, condition CMBufferQueueTriggerCondition, time CMTime, triggerTokenOut *CMBufferQueueTriggerToken, handler CMBufferQueueTriggerHandler) int32

// CMBufferQueueInstallTriggerHandler installs a trigger with a handler on a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueInstallTriggerHandler(_:_:_:_:_:)
func CMBufferQueueInstallTriggerHandler(queue uintptr, condition CMBufferQueueTriggerCondition, time CMTime, triggerTokenOut *CMBufferQueueTriggerToken, handler CMBufferQueueTriggerHandler) int32 {
	if _cMBufferQueueInstallTriggerHandler == nil {
		panic("CoreMedia: symbol CMBufferQueueInstallTriggerHandler not loaded")
	}
	return _cMBufferQueueInstallTriggerHandler(queue, condition, time, triggerTokenOut, handler)
}

var _cMBufferQueueInstallTriggerHandlerWithIntegerThreshold func(queue uintptr, condition CMBufferQueueTriggerCondition, threshold int, triggerTokenOut *CMBufferQueueTriggerToken, handler CMBufferQueueTriggerHandler) int32

// CMBufferQueueInstallTriggerHandlerWithIntegerThreshold installs a trigger with a handler and threshold on a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueInstallTriggerHandlerWithIntegerThreshold(_:_:_:_:_:)
func CMBufferQueueInstallTriggerHandlerWithIntegerThreshold(queue uintptr, condition CMBufferQueueTriggerCondition, threshold int, triggerTokenOut *CMBufferQueueTriggerToken, handler CMBufferQueueTriggerHandler) int32 {
	if _cMBufferQueueInstallTriggerHandlerWithIntegerThreshold == nil {
		panic("CoreMedia: symbol CMBufferQueueInstallTriggerHandlerWithIntegerThreshold not loaded")
	}
	return _cMBufferQueueInstallTriggerHandlerWithIntegerThreshold(queue, condition, threshold, triggerTokenOut, handler)
}

var _cMBufferQueueInstallTriggerWithIntegerThreshold func(queue uintptr, callback CMBufferQueueTriggerCallback, refcon uintptr, condition CMBufferQueueTriggerCondition, threshold int, triggerTokenOut *CMBufferQueueTriggerToken) int32

// CMBufferQueueInstallTriggerWithIntegerThreshold installs a trigger with a callback and threshold on a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueInstallTriggerWithIntegerThreshold(_:callback:refcon:condition:threshold:triggerTokenOut:)
func CMBufferQueueInstallTriggerWithIntegerThreshold(queue uintptr, callback CMBufferQueueTriggerCallback, refcon uintptr, condition CMBufferQueueTriggerCondition, threshold int, triggerTokenOut *CMBufferQueueTriggerToken) int32 {
	if _cMBufferQueueInstallTriggerWithIntegerThreshold == nil {
		panic("CoreMedia: symbol CMBufferQueueInstallTriggerWithIntegerThreshold not loaded")
	}
	return _cMBufferQueueInstallTriggerWithIntegerThreshold(queue, callback, refcon, condition, threshold, triggerTokenOut)
}

var _cMBufferQueueIsAtEndOfData func(queue uintptr) bool

// CMBufferQueueIsAtEndOfData returns a Boolean value that indicates whether a buffer queue has its end-of-data marker set, and is now empty.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueIsAtEndOfData(_:)
func CMBufferQueueIsAtEndOfData(queue uintptr) bool {
	if _cMBufferQueueIsAtEndOfData == nil {
		panic("CoreMedia: symbol CMBufferQueueIsAtEndOfData not loaded")
	}
	return _cMBufferQueueIsAtEndOfData(queue)
}

var _cMBufferQueueIsEmpty func(queue uintptr) bool

// CMBufferQueueIsEmpty returns a Boolean value that indicates whether a buffer queue is empty.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueIsEmpty(_:)
func CMBufferQueueIsEmpty(queue uintptr) bool {
	if _cMBufferQueueIsEmpty == nil {
		panic("CoreMedia: symbol CMBufferQueueIsEmpty not loaded")
	}
	return _cMBufferQueueIsEmpty(queue)
}

var _cMBufferQueueMarkEndOfData func(queue uintptr) int32

// CMBufferQueueMarkEndOfData sets a marker to indicate this queue doesn’t allow enqueuing new buffers.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueMarkEndOfData(_:)
func CMBufferQueueMarkEndOfData(queue uintptr) int32 {
	if _cMBufferQueueMarkEndOfData == nil {
		panic("CoreMedia: symbol CMBufferQueueMarkEndOfData not loaded")
	}
	return _cMBufferQueueMarkEndOfData(queue)
}

var _cMBufferQueueRemoveTrigger func(queue uintptr, triggerToken CMBufferQueueTriggerToken) int32

// CMBufferQueueRemoveTrigger removes a previously installed trigger from a buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueRemoveTrigger(_:triggerToken:)
func CMBufferQueueRemoveTrigger(queue uintptr, triggerToken CMBufferQueueTriggerToken) int32 {
	if _cMBufferQueueRemoveTrigger == nil {
		panic("CoreMedia: symbol CMBufferQueueRemoveTrigger not loaded")
	}
	return _cMBufferQueueRemoveTrigger(queue, triggerToken)
}

var _cMBufferQueueReset func(queue uintptr) int32

// CMBufferQueueReset resets a buffer queue, which allows it to enqueue new buffers.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueReset(_:)
func CMBufferQueueReset(queue uintptr) int32 {
	if _cMBufferQueueReset == nil {
		panic("CoreMedia: symbol CMBufferQueueReset not loaded")
	}
	return _cMBufferQueueReset(queue)
}

var _cMBufferQueueResetWithCallback func(queue uintptr, refcon uintptr) int32

// CMBufferQueueResetWithCallback a callback that invokes a function for every buffer in a queue and then resets the queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueResetWithCallback(_:callback:refcon:)
func CMBufferQueueResetWithCallback(queue uintptr, refcon uintptr) int32 {
	if _cMBufferQueueResetWithCallback == nil {
		panic("CoreMedia: symbol CMBufferQueueResetWithCallback not loaded")
	}
	return _cMBufferQueueResetWithCallback(queue, refcon)
}

var _cMBufferQueueSetValidationCallback func(queue uintptr, callback CMBufferValidationCallback, refcon uintptr) int32

// CMBufferQueueSetValidationCallback a validation callback for the queue to call before enqueuing buffers.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueSetValidationCallback(_:callback:refcon:)
func CMBufferQueueSetValidationCallback(queue uintptr, callback CMBufferValidationCallback, refcon uintptr) int32 {
	if _cMBufferQueueSetValidationCallback == nil {
		panic("CoreMedia: symbol CMBufferQueueSetValidationCallback not loaded")
	}
	return _cMBufferQueueSetValidationCallback(queue, callback, refcon)
}

var _cMBufferQueueSetValidationHandler func(queue uintptr, handler CMBufferValidationHandler) int32

// CMBufferQueueSetValidationHandler a validation handler for the queue to call before enqueuing buffers.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueSetValidationHandler(_:_:)
func CMBufferQueueSetValidationHandler(queue uintptr, handler CMBufferValidationHandler) int32 {
	if _cMBufferQueueSetValidationHandler == nil {
		panic("CoreMedia: symbol CMBufferQueueSetValidationHandler not loaded")
	}
	return _cMBufferQueueSetValidationHandler(queue, handler)
}

var _cMBufferQueueTestTrigger func(queue uintptr, triggerToken CMBufferQueueTriggerToken) bool

// CMBufferQueueTestTrigger tests whether the trigger condition is true for the specified buffer queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMBufferQueueTestTrigger(_:triggerToken:)
func CMBufferQueueTestTrigger(queue uintptr, triggerToken CMBufferQueueTriggerToken) bool {
	if _cMBufferQueueTestTrigger == nil {
		panic("CoreMedia: symbol CMBufferQueueTestTrigger not loaded")
	}
	return _cMBufferQueueTestTrigger(queue, triggerToken)
}

var _cMClockConvertHostTimeToSystemUnits func(hostTime CMTime) uint64

// CMClockConvertHostTimeToSystemUnits converts a host time from a core media time structure to the host time’s native units.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClockConvertHostTimeToSystemUnits(_:)
func CMClockConvertHostTimeToSystemUnits(hostTime CMTime) uint64 {
	if _cMClockConvertHostTimeToSystemUnits == nil {
		panic("CoreMedia: symbol CMClockConvertHostTimeToSystemUnits not loaded")
	}
	return _cMClockConvertHostTimeToSystemUnits(hostTime)
}

var _cMClockGetAnchorTime func(clock uintptr, clockTimeOut *CMTime, referenceClockTimeOut *CMTime) int32

// CMClockGetAnchorTime returns the current time from a clock and the matching time from the clock’s reference clock.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClockGetAnchorTime(_:clockTimeOut:referenceClockTimeOut:)
func CMClockGetAnchorTime(clock uintptr, clockTimeOut *CMTime, referenceClockTimeOut *CMTime) int32 {
	if _cMClockGetAnchorTime == nil {
		panic("CoreMedia: symbol CMClockGetAnchorTime not loaded")
	}
	return _cMClockGetAnchorTime(clock, clockTimeOut, referenceClockTimeOut)
}

var _cMClockGetHostTimeClock func() uintptr

// CMClockGetHostTimeClock returns a reference to the singleton clock that reflects the host time.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClockGetHostTimeClock()
func CMClockGetHostTimeClock() uintptr {
	if _cMClockGetHostTimeClock == nil {
		panic("CoreMedia: symbol CMClockGetHostTimeClock not loaded")
	}
	return _cMClockGetHostTimeClock()
}

var _cMClockGetTime func(clock uintptr) CMTime

// CMClockGetTime returns the current time from a clock.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClockGetTime(_:)
func CMClockGetTime(clock uintptr) CMTime {
	if _cMClockGetTime == nil {
		panic("CoreMedia: symbol CMClockGetTime not loaded")
	}
	return _cMClockGetTime(clock)
}

var _cMClockGetTypeID func() uint

// CMClockGetTypeID returns the core foundation type identifier of a clock type.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClockGetTypeID()
func CMClockGetTypeID() uint {
	if _cMClockGetTypeID == nil {
		panic("CoreMedia: symbol CMClockGetTypeID not loaded")
	}
	return _cMClockGetTypeID()
}

var _cMClockInvalidate func(clock uintptr)

// CMClockInvalidate stops the clock.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClockInvalidate(_:)
func CMClockInvalidate(clock uintptr) {
	if _cMClockInvalidate == nil {
		panic("CoreMedia: symbol CMClockInvalidate not loaded")
	}
	_cMClockInvalidate(clock)
}

var _cMClockMakeHostTimeFromSystemUnits func(hostTime uint64) CMTime

// CMClockMakeHostTimeFromSystemUnits converts a host time from native units to a core media time structure.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClockMakeHostTimeFromSystemUnits(_:)
func CMClockMakeHostTimeFromSystemUnits(hostTime uint64) CMTime {
	if _cMClockMakeHostTimeFromSystemUnits == nil {
		panic("CoreMedia: symbol CMClockMakeHostTimeFromSystemUnits not loaded")
	}
	return _cMClockMakeHostTimeFromSystemUnits(hostTime)
}

var _cMClockMightDrift func(clock uintptr, otherClock uintptr) bool

// CMClockMightDrift returns a Boolean value that indicates whether it’s possible for two clocks to drift relative to each other.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClockMightDrift(_:otherClock:)
func CMClockMightDrift(clock uintptr, otherClock uintptr) bool {
	if _cMClockMightDrift == nil {
		panic("CoreMedia: symbol CMClockMightDrift not loaded")
	}
	return _cMClockMightDrift(clock, otherClock)
}

var _cMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, closedCaptionFormatDescription CMClosedCaptionFormatDescriptionRef, flavor CMClosedCaptionDescriptionFlavor, blockBufferOut *uintptr) int32

// CMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer copies the contents of a closed caption format description to a buffer in big-endian byte order.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer(allocator:closedCaptionFormatDescription:flavor:blockBufferOut:)
func CMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, closedCaptionFormatDescription CMClosedCaptionFormatDescriptionRef, flavor CMClosedCaptionDescriptionFlavor, blockBufferOut *uintptr) int32 {
	if _cMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer == nil {
		panic("CoreMedia: symbol CMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer not loaded")
	}
	return _cMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer(allocator, closedCaptionFormatDescription, flavor, blockBufferOut)
}

var _cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, closedCaptionDescriptionBlockBuffer uintptr, flavor CMClosedCaptionDescriptionFlavor, formatDescriptionOut *CMClosedCaptionFormatDescriptionRef) int32

// CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer creates a closed caption format description from a big-endian closed caption description structure in a buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer(allocator:bigEndianClosedCaptionDescriptionBlockBuffer:flavor:formatDescriptionOut:)
func CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, closedCaptionDescriptionBlockBuffer uintptr, flavor CMClosedCaptionDescriptionFlavor, formatDescriptionOut *CMClosedCaptionFormatDescriptionRef) int32 {
	if _cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer == nil {
		panic("CoreMedia: symbol CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer not loaded")
	}
	return _cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer(allocator, closedCaptionDescriptionBlockBuffer, flavor, formatDescriptionOut)
}

var _cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData func(allocator corefoundation.CFAllocatorRef, closedCaptionDescriptionData *uint8, size uintptr, flavor CMClosedCaptionDescriptionFlavor, formatDescriptionOut *CMClosedCaptionFormatDescriptionRef) int32

// CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData creates a closed caption format description from a big-endian closed caption description structure.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData(allocator:bigEndianClosedCaptionDescriptionData:size:flavor:formatDescriptionOut:)
func CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData(allocator corefoundation.CFAllocatorRef, closedCaptionDescriptionData *uint8, size uintptr, flavor CMClosedCaptionDescriptionFlavor, formatDescriptionOut *CMClosedCaptionFormatDescriptionRef) int32 {
	if _cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData == nil {
		panic("CoreMedia: symbol CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData not loaded")
	}
	return _cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData(allocator, closedCaptionDescriptionData, size, flavor, formatDescriptionOut)
}

var _cMCopyDictionaryOfAttachments func(allocator corefoundation.CFAllocatorRef, target CMAttachmentBearerRef, attachmentMode CMAttachmentMode) corefoundation.CFDictionaryRef

// CMCopyDictionaryOfAttachments returns a dictionary of all attachments for an attachment bearer object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMCopyDictionaryOfAttachments(allocator:target:attachmentMode:)
func CMCopyDictionaryOfAttachments(allocator corefoundation.CFAllocatorRef, target CMAttachmentBearerRef, attachmentMode CMAttachmentMode) corefoundation.CFDictionaryRef {
	if _cMCopyDictionaryOfAttachments == nil {
		panic("CoreMedia: symbol CMCopyDictionaryOfAttachments not loaded")
	}
	return _cMCopyDictionaryOfAttachments(allocator, target, attachmentMode)
}

var _cMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout func(soundDescriptionBlockBuffer uintptr, flavor CMSoundDescriptionFlavor) bool

// CMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout returns a Boolean value that indicates whether the sample tables need to use the legacy constant bit-rate encoding layout.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout(_:flavor:)
func CMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout(soundDescriptionBlockBuffer uintptr, flavor CMSoundDescriptionFlavor) bool {
	if _cMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout == nil {
		panic("CoreMedia: symbol CMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout not loaded")
	}
	return _cMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout(soundDescriptionBlockBuffer, flavor)
}

var _cMFormatDescriptionCreate func(allocator corefoundation.CFAllocatorRef, mediaType uint32, mediaSubType uint32, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *uintptr) int32

// CMFormatDescriptionCreate creates a format description for general use.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionCreate(allocator:mediaType:mediaSubType:extensions:formatDescriptionOut:)
func CMFormatDescriptionCreate(allocator corefoundation.CFAllocatorRef, mediaType uint32, mediaSubType uint32, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *uintptr) int32 {
	if _cMFormatDescriptionCreate == nil {
		panic("CoreMedia: symbol CMFormatDescriptionCreate not loaded")
	}
	return _cMFormatDescriptionCreate(allocator, mediaType, mediaSubType, extensions, formatDescriptionOut)
}

var _cMFormatDescriptionEqual func(formatDescription uintptr, otherFormatDescription uintptr) bool

// CMFormatDescriptionEqual returns a Boolean value that indicates whether two format descriptions are equal.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionEqual(_:otherFormatDescription:)
func CMFormatDescriptionEqual(formatDescription uintptr, otherFormatDescription uintptr) bool {
	if _cMFormatDescriptionEqual == nil {
		panic("CoreMedia: symbol CMFormatDescriptionEqual not loaded")
	}
	return _cMFormatDescriptionEqual(formatDescription, otherFormatDescription)
}

var _cMFormatDescriptionEqualIgnoringExtensionKeys func(formatDescription uintptr, otherFormatDescription uintptr, formatDescriptionExtensionKeysToIgnore corefoundation.CFTypeRef, sampleDescriptionExtensionAtomKeysToIgnore corefoundation.CFTypeRef) bool

// CMFormatDescriptionEqualIgnoringExtensionKeys returns a Boolean value that indicates whether two format descriptions are equal, ignoring differences in the extension keys you specify.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionEqualIgnoringExtensionKeys(_:otherFormatDescription:extensionKeysToIgnore:sampleDescriptionExtensionAtomKeysToIgnore:)
func CMFormatDescriptionEqualIgnoringExtensionKeys(formatDescription uintptr, otherFormatDescription uintptr, formatDescriptionExtensionKeysToIgnore corefoundation.CFTypeRef, sampleDescriptionExtensionAtomKeysToIgnore corefoundation.CFTypeRef) bool {
	if _cMFormatDescriptionEqualIgnoringExtensionKeys == nil {
		panic("CoreMedia: symbol CMFormatDescriptionEqualIgnoringExtensionKeys not loaded")
	}
	return _cMFormatDescriptionEqualIgnoringExtensionKeys(formatDescription, otherFormatDescription, formatDescriptionExtensionKeysToIgnore, sampleDescriptionExtensionAtomKeysToIgnore)
}

var _cMFormatDescriptionGetExtension func(desc uintptr, extensionKey corefoundation.CFStringRef) corefoundation.CFPropertyListRef

// CMFormatDescriptionGetExtension returns an extension from the format description by using an extension key.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetExtension(_:extensionKey:)
func CMFormatDescriptionGetExtension(desc uintptr, extensionKey corefoundation.CFStringRef) corefoundation.CFPropertyListRef {
	if _cMFormatDescriptionGetExtension == nil {
		panic("CoreMedia: symbol CMFormatDescriptionGetExtension not loaded")
	}
	return _cMFormatDescriptionGetExtension(desc, extensionKey)
}

var _cMFormatDescriptionGetExtensions func(desc uintptr) corefoundation.CFDictionaryRef

// CMFormatDescriptionGetExtensions returns all of the extensions for a format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetExtensions(_:)
func CMFormatDescriptionGetExtensions(desc uintptr) corefoundation.CFDictionaryRef {
	if _cMFormatDescriptionGetExtensions == nil {
		panic("CoreMedia: symbol CMFormatDescriptionGetExtensions not loaded")
	}
	return _cMFormatDescriptionGetExtensions(desc)
}

var _cMFormatDescriptionGetMediaSubType func(desc uintptr) uint32

// CMFormatDescriptionGetMediaSubType returns the media subtype of a format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetMediaSubType(_:)
func CMFormatDescriptionGetMediaSubType(desc uintptr) uint32 {
	if _cMFormatDescriptionGetMediaSubType == nil {
		panic("CoreMedia: symbol CMFormatDescriptionGetMediaSubType not loaded")
	}
	return _cMFormatDescriptionGetMediaSubType(desc)
}

var _cMFormatDescriptionGetMediaType func(desc uintptr) uint32

// CMFormatDescriptionGetMediaType returns the media type of a format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetMediaType(_:)
func CMFormatDescriptionGetMediaType(desc uintptr) uint32 {
	if _cMFormatDescriptionGetMediaType == nil {
		panic("CoreMedia: symbol CMFormatDescriptionGetMediaType not loaded")
	}
	return _cMFormatDescriptionGetMediaType(desc)
}

var _cMFormatDescriptionGetTypeID func() uint

// CMFormatDescriptionGetTypeID returns the Core Foundation type identifier that identifies format description objects.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMFormatDescriptionGetTypeID()
func CMFormatDescriptionGetTypeID() uint {
	if _cMFormatDescriptionGetTypeID == nil {
		panic("CoreMedia: symbol CMFormatDescriptionGetTypeID not loaded")
	}
	return _cMFormatDescriptionGetTypeID()
}

var _cMGetAttachment func(target CMAttachmentBearerRef, key corefoundation.CFStringRef, attachmentModeOut *CMAttachmentMode) corefoundation.CFTypeRef

// CMGetAttachment returns an attachment from an attachment bearer object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMGetAttachment(_:key:attachmentModeOut:)
func CMGetAttachment(target CMAttachmentBearerRef, key corefoundation.CFStringRef, attachmentModeOut *CMAttachmentMode) corefoundation.CFTypeRef {
	if _cMGetAttachment == nil {
		panic("CoreMedia: symbol CMGetAttachment not loaded")
	}
	return _cMGetAttachment(target, key, attachmentModeOut)
}

var _cMMemoryPoolCreate func(options corefoundation.CFDictionaryRef) uintptr

// CMMemoryPoolCreate creates a memory pool.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolCreate(options:)
func CMMemoryPoolCreate(options corefoundation.CFDictionaryRef) uintptr {
	if _cMMemoryPoolCreate == nil {
		panic("CoreMedia: symbol CMMemoryPoolCreate not loaded")
	}
	return _cMMemoryPoolCreate(options)
}

var _cMMemoryPoolFlush func(pool uintptr)

// CMMemoryPoolFlush deallocates all memory the pool holds.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolFlush(_:)
func CMMemoryPoolFlush(pool uintptr) {
	if _cMMemoryPoolFlush == nil {
		panic("CoreMedia: symbol CMMemoryPoolFlush not loaded")
	}
	_cMMemoryPoolFlush(pool)
}

var _cMMemoryPoolGetAllocator func(pool uintptr) corefoundation.CFAllocatorRef

// CMMemoryPoolGetAllocator returns the allocator for the memory pool.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolGetAllocator(_:)
func CMMemoryPoolGetAllocator(pool uintptr) corefoundation.CFAllocatorRef {
	if _cMMemoryPoolGetAllocator == nil {
		panic("CoreMedia: symbol CMMemoryPoolGetAllocator not loaded")
	}
	return _cMMemoryPoolGetAllocator(pool)
}

var _cMMemoryPoolGetTypeID func() uint

// CMMemoryPoolGetTypeID returns the type identifier of memory pool objects.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolGetTypeID()
func CMMemoryPoolGetTypeID() uint {
	if _cMMemoryPoolGetTypeID == nil {
		panic("CoreMedia: symbol CMMemoryPoolGetTypeID not loaded")
	}
	return _cMMemoryPoolGetTypeID()
}

var _cMMemoryPoolInvalidate func(pool uintptr)

// CMMemoryPoolInvalidate invalidates the memory pool, which causes its allocator to stop recycling memory.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMemoryPoolInvalidate(_:)
func CMMemoryPoolInvalidate(pool uintptr) {
	if _cMMemoryPoolInvalidate == nil {
		panic("CoreMedia: symbol CMMemoryPoolInvalidate not loaded")
	}
	_cMMemoryPoolInvalidate(pool)
}

var _cMMetadataCreateIdentifierForKeyAndKeySpace func(allocator corefoundation.CFAllocatorRef, key corefoundation.CFTypeRef, keySpace corefoundation.CFStringRef, identifierOut *corefoundation.CFStringRef) int32

// CMMetadataCreateIdentifierForKeyAndKeySpace creates a URL-like string identifier that represents a key or keyspace tuple.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataCreateIdentifierForKeyAndKeySpace(allocator:key:keySpace:identifierOut:)
func CMMetadataCreateIdentifierForKeyAndKeySpace(allocator corefoundation.CFAllocatorRef, key corefoundation.CFTypeRef, keySpace corefoundation.CFStringRef, identifierOut *corefoundation.CFStringRef) int32 {
	if _cMMetadataCreateIdentifierForKeyAndKeySpace == nil {
		panic("CoreMedia: symbol CMMetadataCreateIdentifierForKeyAndKeySpace not loaded")
	}
	return _cMMetadataCreateIdentifierForKeyAndKeySpace(allocator, key, keySpace, identifierOut)
}

var _cMMetadataCreateKeyFromIdentifier func(allocator corefoundation.CFAllocatorRef, identifier corefoundation.CFStringRef, keyOut *corefoundation.CFTypeRef) int32

// CMMetadataCreateKeyFromIdentifier creates a copy of the key by using an identifier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataCreateKeyFromIdentifier(allocator:identifier:keyOut:)
func CMMetadataCreateKeyFromIdentifier(allocator corefoundation.CFAllocatorRef, identifier corefoundation.CFStringRef, keyOut *corefoundation.CFTypeRef) int32 {
	if _cMMetadataCreateKeyFromIdentifier == nil {
		panic("CoreMedia: symbol CMMetadataCreateKeyFromIdentifier not loaded")
	}
	return _cMMetadataCreateKeyFromIdentifier(allocator, identifier, keyOut)
}

var _cMMetadataCreateKeyFromIdentifierAsCFData func(allocator corefoundation.CFAllocatorRef, identifier corefoundation.CFStringRef, keyOut *corefoundation.CFDataRef) int32

// CMMetadataCreateKeyFromIdentifierAsCFData creates a copy of the key by using an identifier, and results in a core foundation data object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataCreateKeyFromIdentifierAsCFData(allocator:identifier:keyOut:)
func CMMetadataCreateKeyFromIdentifierAsCFData(allocator corefoundation.CFAllocatorRef, identifier corefoundation.CFStringRef, keyOut *corefoundation.CFDataRef) int32 {
	if _cMMetadataCreateKeyFromIdentifierAsCFData == nil {
		panic("CoreMedia: symbol CMMetadataCreateKeyFromIdentifierAsCFData not loaded")
	}
	return _cMMetadataCreateKeyFromIdentifierAsCFData(allocator, identifier, keyOut)
}

var _cMMetadataCreateKeySpaceFromIdentifier func(allocator corefoundation.CFAllocatorRef, identifier corefoundation.CFStringRef, keySpaceOut *corefoundation.CFStringRef) int32

// CMMetadataCreateKeySpaceFromIdentifier creates a copy of the keyspace by using an identifier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataCreateKeySpaceFromIdentifier(allocator:identifier:keySpaceOut:)
func CMMetadataCreateKeySpaceFromIdentifier(allocator corefoundation.CFAllocatorRef, identifier corefoundation.CFStringRef, keySpaceOut *corefoundation.CFStringRef) int32 {
	if _cMMetadataCreateKeySpaceFromIdentifier == nil {
		panic("CoreMedia: symbol CMMetadataCreateKeySpaceFromIdentifier not loaded")
	}
	return _cMMetadataCreateKeySpaceFromIdentifier(allocator, identifier, keySpaceOut)
}

var _cMMetadataDataTypeRegistryDataTypeConformsToDataType func(dataType corefoundation.CFStringRef, conformsToDataType corefoundation.CFStringRef) bool

// CMMetadataDataTypeRegistryDataTypeConformsToDataType returns a Boolean value that indicates whether a data type conforms to another data type.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryDataTypeConformsToDataType(_:conformsTo:)
func CMMetadataDataTypeRegistryDataTypeConformsToDataType(dataType corefoundation.CFStringRef, conformsToDataType corefoundation.CFStringRef) bool {
	if _cMMetadataDataTypeRegistryDataTypeConformsToDataType == nil {
		panic("CoreMedia: symbol CMMetadataDataTypeRegistryDataTypeConformsToDataType not loaded")
	}
	return _cMMetadataDataTypeRegistryDataTypeConformsToDataType(dataType, conformsToDataType)
}

var _cMMetadataDataTypeRegistryDataTypeIsBaseDataType func(dataType corefoundation.CFStringRef) bool

// CMMetadataDataTypeRegistryDataTypeIsBaseDataType returns a Boolean value that indicates whether a data type identifier represents a base data type.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryDataTypeIsBaseDataType(_:)
func CMMetadataDataTypeRegistryDataTypeIsBaseDataType(dataType corefoundation.CFStringRef) bool {
	if _cMMetadataDataTypeRegistryDataTypeIsBaseDataType == nil {
		panic("CoreMedia: symbol CMMetadataDataTypeRegistryDataTypeIsBaseDataType not loaded")
	}
	return _cMMetadataDataTypeRegistryDataTypeIsBaseDataType(dataType)
}

var _cMMetadataDataTypeRegistryDataTypeIsRegistered func(dataType corefoundation.CFStringRef) bool

// CMMetadataDataTypeRegistryDataTypeIsRegistered returns a Boolean value that indicates the registration status of a data type identifier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryDataTypeIsRegistered(_:)
func CMMetadataDataTypeRegistryDataTypeIsRegistered(dataType corefoundation.CFStringRef) bool {
	if _cMMetadataDataTypeRegistryDataTypeIsRegistered == nil {
		panic("CoreMedia: symbol CMMetadataDataTypeRegistryDataTypeIsRegistered not loaded")
	}
	return _cMMetadataDataTypeRegistryDataTypeIsRegistered(dataType)
}

var _cMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType func(dataType corefoundation.CFStringRef) corefoundation.CFStringRef

// CMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType returns the base data type identifier that a data type conforms to.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType(_:)
func CMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType(dataType corefoundation.CFStringRef) corefoundation.CFStringRef {
	if _cMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType == nil {
		panic("CoreMedia: symbol CMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType not loaded")
	}
	return _cMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType(dataType)
}

var _cMMetadataDataTypeRegistryGetBaseDataTypes func() corefoundation.CFArrayRef

// CMMetadataDataTypeRegistryGetBaseDataTypes returns an array of base data type identifiers.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryGetBaseDataTypes()
func CMMetadataDataTypeRegistryGetBaseDataTypes() corefoundation.CFArrayRef {
	if _cMMetadataDataTypeRegistryGetBaseDataTypes == nil {
		panic("CoreMedia: symbol CMMetadataDataTypeRegistryGetBaseDataTypes not loaded")
	}
	return _cMMetadataDataTypeRegistryGetBaseDataTypes()
}

var _cMMetadataDataTypeRegistryGetConformingDataTypes func(dataType corefoundation.CFStringRef) corefoundation.CFArrayRef

// CMMetadataDataTypeRegistryGetConformingDataTypes returns the conforming data types for the data type, if any.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryGetConformingDataTypes(_:)
func CMMetadataDataTypeRegistryGetConformingDataTypes(dataType corefoundation.CFStringRef) corefoundation.CFArrayRef {
	if _cMMetadataDataTypeRegistryGetConformingDataTypes == nil {
		panic("CoreMedia: symbol CMMetadataDataTypeRegistryGetConformingDataTypes not loaded")
	}
	return _cMMetadataDataTypeRegistryGetConformingDataTypes(dataType)
}

var _cMMetadataDataTypeRegistryGetDataTypeDescription func(dataType corefoundation.CFStringRef) corefoundation.CFStringRef

// CMMetadataDataTypeRegistryGetDataTypeDescription returns the data type description if it exists.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryGetDataTypeDescription(_:)
func CMMetadataDataTypeRegistryGetDataTypeDescription(dataType corefoundation.CFStringRef) corefoundation.CFStringRef {
	if _cMMetadataDataTypeRegistryGetDataTypeDescription == nil {
		panic("CoreMedia: symbol CMMetadataDataTypeRegistryGetDataTypeDescription not loaded")
	}
	return _cMMetadataDataTypeRegistryGetDataTypeDescription(dataType)
}

var _cMMetadataDataTypeRegistryRegisterDataType func(dataType corefoundation.CFStringRef, description corefoundation.CFStringRef, conformingDataTypes corefoundation.CFArrayRef) int32

// CMMetadataDataTypeRegistryRegisterDataType register a data type with the data type registry.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataDataTypeRegistryRegisterDataType(_:description:conformingDataTypes:)
func CMMetadataDataTypeRegistryRegisterDataType(dataType corefoundation.CFStringRef, description corefoundation.CFStringRef, conformingDataTypes corefoundation.CFArrayRef) int32 {
	if _cMMetadataDataTypeRegistryRegisterDataType == nil {
		panic("CoreMedia: symbol CMMetadataDataTypeRegistryRegisterDataType not loaded")
	}
	return _cMMetadataDataTypeRegistryRegisterDataType(dataType, description, conformingDataTypes)
}

var _cMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, metadataFormatDescription CMMetadataFormatDescriptionRef, flavor CMMetadataDescriptionFlavor, blockBufferOut *uintptr) int32

// CMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer copies the contents of a metadata format description to a buffer in big-endian byte order.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer(allocator:metadataFormatDescription:flavor:blockBufferOut:)
func CMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, metadataFormatDescription CMMetadataFormatDescriptionRef, flavor CMMetadataDescriptionFlavor, blockBufferOut *uintptr) int32 {
	if _cMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer == nil {
		panic("CoreMedia: symbol CMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer not loaded")
	}
	return _cMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer(allocator, metadataFormatDescription, flavor, blockBufferOut)
}

var _cMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions func(allocator corefoundation.CFAllocatorRef, sourceDescription CMMetadataFormatDescriptionRef, otherSourceDescription CMMetadataFormatDescriptionRef, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32

// CMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions creates a metadata format description object by merging with another description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions(allocator:sourceDescription:otherSourceDescription:formatDescriptionOut:)
func CMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions(allocator corefoundation.CFAllocatorRef, sourceDescription CMMetadataFormatDescriptionRef, otherSourceDescription CMMetadataFormatDescriptionRef, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32 {
	if _cMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions == nil {
		panic("CoreMedia: symbol CMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions not loaded")
	}
	return _cMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions(allocator, sourceDescription, otherSourceDescription, formatDescriptionOut)
}

var _cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, metadataDescriptionBlockBuffer uintptr, flavor CMMetadataDescriptionFlavor, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32

// CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer creates a metadata format description from a big-endian metadata description structure inside a buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer(allocator:bigEndianMetadataDescriptionBlockBuffer:flavor:formatDescriptionOut:)
func CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, metadataDescriptionBlockBuffer uintptr, flavor CMMetadataDescriptionFlavor, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32 {
	if _cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer == nil {
		panic("CoreMedia: symbol CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer not loaded")
	}
	return _cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer(allocator, metadataDescriptionBlockBuffer, flavor, formatDescriptionOut)
}

var _cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData func(allocator corefoundation.CFAllocatorRef, metadataDescriptionData *uint8, size uintptr, flavor CMMetadataDescriptionFlavor, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32

// CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData creates a metadata format description from a big-endian metadata description structure.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData(allocator:bigEndianMetadataDescriptionData:size:flavor:formatDescriptionOut:)
func CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData(allocator corefoundation.CFAllocatorRef, metadataDescriptionData *uint8, size uintptr, flavor CMMetadataDescriptionFlavor, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32 {
	if _cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData == nil {
		panic("CoreMedia: symbol CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData not loaded")
	}
	return _cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData(allocator, metadataDescriptionData, size, flavor, formatDescriptionOut)
}

var _cMMetadataFormatDescriptionCreateWithKeys func(allocator corefoundation.CFAllocatorRef, metadataType CMMetadataFormatType, keys corefoundation.CFArrayRef, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32

// CMMetadataFormatDescriptionCreateWithKeys creates a metadata format description with the metadata keys you specify.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCreateWithKeys(allocator:metadataType:keys:formatDescriptionOut:)
func CMMetadataFormatDescriptionCreateWithKeys(allocator corefoundation.CFAllocatorRef, metadataType CMMetadataFormatType, keys corefoundation.CFArrayRef, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32 {
	if _cMMetadataFormatDescriptionCreateWithKeys == nil {
		panic("CoreMedia: symbol CMMetadataFormatDescriptionCreateWithKeys not loaded")
	}
	return _cMMetadataFormatDescriptionCreateWithKeys(allocator, metadataType, keys, formatDescriptionOut)
}

var _cMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications func(allocator corefoundation.CFAllocatorRef, sourceDescription CMMetadataFormatDescriptionRef, metadataSpecifications corefoundation.CFArrayRef, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32

// CMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications creates a metadata format description by extending an existing description with the values you specify.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications(allocator:sourceDescription:metadataSpecifications:formatDescriptionOut:)
func CMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications(allocator corefoundation.CFAllocatorRef, sourceDescription CMMetadataFormatDescriptionRef, metadataSpecifications corefoundation.CFArrayRef, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32 {
	if _cMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications == nil {
		panic("CoreMedia: symbol CMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications not loaded")
	}
	return _cMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications(allocator, sourceDescription, metadataSpecifications, formatDescriptionOut)
}

var _cMMetadataFormatDescriptionCreateWithMetadataSpecifications func(allocator corefoundation.CFAllocatorRef, metadataType CMMetadataFormatType, metadataSpecifications corefoundation.CFArrayRef, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32

// CMMetadataFormatDescriptionCreateWithMetadataSpecifications creates a metadata format description with the specifications you specify.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionCreateWithMetadataSpecifications(allocator:metadataType:metadataSpecifications:formatDescriptionOut:)
func CMMetadataFormatDescriptionCreateWithMetadataSpecifications(allocator corefoundation.CFAllocatorRef, metadataType CMMetadataFormatType, metadataSpecifications corefoundation.CFArrayRef, formatDescriptionOut *CMMetadataFormatDescriptionRef) int32 {
	if _cMMetadataFormatDescriptionCreateWithMetadataSpecifications == nil {
		panic("CoreMedia: symbol CMMetadataFormatDescriptionCreateWithMetadataSpecifications not loaded")
	}
	return _cMMetadataFormatDescriptionCreateWithMetadataSpecifications(allocator, metadataType, metadataSpecifications, formatDescriptionOut)
}

var _cMMetadataFormatDescriptionGetIdentifiers func(desc CMMetadataFormatDescriptionRef) corefoundation.CFArrayRef

// CMMetadataFormatDescriptionGetIdentifiers returns an array of metadata identifiers from a metadata format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionGetIdentifiers(_:)
func CMMetadataFormatDescriptionGetIdentifiers(desc CMMetadataFormatDescriptionRef) corefoundation.CFArrayRef {
	if _cMMetadataFormatDescriptionGetIdentifiers == nil {
		panic("CoreMedia: symbol CMMetadataFormatDescriptionGetIdentifiers not loaded")
	}
	return _cMMetadataFormatDescriptionGetIdentifiers(desc)
}

var _cMMetadataFormatDescriptionGetKeyWithLocalID func(desc CMMetadataFormatDescriptionRef, localKeyID uint32) corefoundation.CFDictionaryRef

// CMMetadataFormatDescriptionGetKeyWithLocalID returns the key for the local identifier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMetadataFormatDescriptionGetKeyWithLocalID(_:localKeyID:)
func CMMetadataFormatDescriptionGetKeyWithLocalID(desc CMMetadataFormatDescriptionRef, localKeyID uint32) corefoundation.CFDictionaryRef {
	if _cMMetadataFormatDescriptionGetKeyWithLocalID == nil {
		panic("CoreMedia: symbol CMMetadataFormatDescriptionGetKeyWithLocalID not loaded")
	}
	return _cMMetadataFormatDescriptionGetKeyWithLocalID(desc, localKeyID)
}

var _cMMuxedFormatDescriptionCreate func(allocator corefoundation.CFAllocatorRef, muxType CMMuxedStreamType, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *CMMuxedFormatDescriptionRef) int32

// CMMuxedFormatDescriptionCreate creates a format description for a muxed media stream.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMMuxedFormatDescriptionCreate(allocator:muxType:extensions:formatDescriptionOut:)
func CMMuxedFormatDescriptionCreate(allocator corefoundation.CFAllocatorRef, muxType CMMuxedStreamType, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *CMMuxedFormatDescriptionRef) int32 {
	if _cMMuxedFormatDescriptionCreate == nil {
		panic("CoreMedia: symbol CMMuxedFormatDescriptionCreate not loaded")
	}
	return _cMMuxedFormatDescriptionCreate(allocator, muxType, extensions, formatDescriptionOut)
}

var _cMPropagateAttachments func(source CMAttachmentBearerRef, destination CMAttachmentBearerRef)

// CMPropagateAttachments copies all propagable attachments from one attachment bearer object to another.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMPropagateAttachments(_:destination:)
func CMPropagateAttachments(source CMAttachmentBearerRef, destination CMAttachmentBearerRef) {
	if _cMPropagateAttachments == nil {
		panic("CoreMedia: symbol CMPropagateAttachments not loaded")
	}
	_cMPropagateAttachments(source, destination)
}

var _cMRemoveAllAttachments func(target CMAttachmentBearerRef)

// CMRemoveAllAttachments removes all attachments from an attachment bearer object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMRemoveAllAttachments(_:)
func CMRemoveAllAttachments(target CMAttachmentBearerRef) {
	if _cMRemoveAllAttachments == nil {
		panic("CoreMedia: symbol CMRemoveAllAttachments not loaded")
	}
	_cMRemoveAllAttachments(target)
}

var _cMRemoveAttachment func(target CMAttachmentBearerRef, key corefoundation.CFStringRef)

// CMRemoveAttachment removes a specific attachment from an attachment bearer object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMRemoveAttachment(_:key:)
func CMRemoveAttachment(target CMAttachmentBearerRef, key corefoundation.CFStringRef) {
	if _cMRemoveAttachment == nil {
		panic("CoreMedia: symbol CMRemoveAttachment not loaded")
	}
	_cMRemoveAttachment(target, key)
}

var _cMSampleBufferCallBlockForEachSample func(sbuf uintptr) int32

// CMSampleBufferCallBlockForEachSample calls a block for every individual sample in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCallBlockForEachSample(_:_:)
func CMSampleBufferCallBlockForEachSample(sbuf uintptr) int32 {
	if _cMSampleBufferCallBlockForEachSample == nil {
		panic("CoreMedia: symbol CMSampleBufferCallBlockForEachSample not loaded")
	}
	return _cMSampleBufferCallBlockForEachSample(sbuf)
}

var _cMSampleBufferCallForEachSample func(sbuf uintptr, refcon uintptr) int32

// CMSampleBufferCallForEachSample calls a function for every individual sample in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCallForEachSample(_:callback:refcon:)
func CMSampleBufferCallForEachSample(sbuf uintptr, refcon uintptr) int32 {
	if _cMSampleBufferCallForEachSample == nil {
		panic("CoreMedia: symbol CMSampleBufferCallForEachSample not loaded")
	}
	return _cMSampleBufferCallForEachSample(sbuf, refcon)
}

var _cMSampleBufferCopyPCMDataIntoAudioBufferList func(sbuf uintptr, frameOffset int32, numFrames int32, bufferList uintptr) int32

// CMSampleBufferCopyPCMDataIntoAudioBufferList copies PCM audio data from a sample buffer into an audio buffer list.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCopyPCMDataIntoAudioBufferList(_:at:frameCount:into:)
func CMSampleBufferCopyPCMDataIntoAudioBufferList(sbuf uintptr, frameOffset int32, numFrames int32, bufferList uintptr) int32 {
	if _cMSampleBufferCopyPCMDataIntoAudioBufferList == nil {
		panic("CoreMedia: symbol CMSampleBufferCopyPCMDataIntoAudioBufferList not loaded")
	}
	return _cMSampleBufferCopyPCMDataIntoAudioBufferList(sbuf, frameOffset, numFrames, bufferList)
}

var _cMSampleBufferCopySampleBufferForRange func(allocator corefoundation.CFAllocatorRef, sbuf uintptr, sampleRange corefoundation.CFRange, sampleBufferOut *uintptr) int32

// CMSampleBufferCopySampleBufferForRange creates a sample buffer that contains a range of samples from an existing sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCopySampleBufferForRange(allocator:sampleBuffer:sampleRange:sampleBufferOut:)
func CMSampleBufferCopySampleBufferForRange(allocator corefoundation.CFAllocatorRef, sbuf uintptr, sampleRange corefoundation.CFRange, sampleBufferOut *uintptr) int32 {
	if _cMSampleBufferCopySampleBufferForRange == nil {
		panic("CoreMedia: symbol CMSampleBufferCopySampleBufferForRange not loaded")
	}
	return _cMSampleBufferCopySampleBufferForRange(allocator, sbuf, sampleRange, sampleBufferOut)
}

var _cMSampleBufferCreate func(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, dataReady bool, makeDataReadyCallback CMSampleBufferMakeDataReadyCallback, makeDataReadyRefcon unsafe.Pointer, formatDescription uintptr, numSamples int, numSampleTimingEntries int, sampleTimingArray *CMSampleTimingInfo, numSampleSizeEntries int, sampleSizeArray *uintptr, sampleBufferOut *uintptr) int32

// CMSampleBufferCreate creates a sample buffer with a callback to make the data ready for use.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreate(allocator:dataBuffer:dataReady:makeDataReadyCallback:refcon:formatDescription:sampleCount:sampleTimingEntryCount:sampleTimingArray:sampleSizeEntryCount:sampleSizeArray:sampleBufferOut:)
func CMSampleBufferCreate(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, dataReady bool, makeDataReadyCallback CMSampleBufferMakeDataReadyCallback, makeDataReadyRefcon unsafe.Pointer, formatDescription uintptr, numSamples int, numSampleTimingEntries int, sampleTimingArray *CMSampleTimingInfo, numSampleSizeEntries int, sampleSizeArray *uintptr, sampleBufferOut *uintptr) int32 {
	if _cMSampleBufferCreate == nil {
		panic("CoreMedia: symbol CMSampleBufferCreate not loaded")
	}
	return _cMSampleBufferCreate(allocator, dataBuffer, dataReady, makeDataReadyCallback, makeDataReadyRefcon, formatDescription, numSamples, numSampleTimingEntries, sampleTimingArray, numSampleSizeEntries, sampleSizeArray, sampleBufferOut)
}

var _cMSampleBufferCreateCopy func(allocator corefoundation.CFAllocatorRef, sbuf uintptr, sampleBufferOut *uintptr) int32

// CMSampleBufferCreateCopy creates a copy of a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateCopy(allocator:sampleBuffer:sampleBufferOut:)
func CMSampleBufferCreateCopy(allocator corefoundation.CFAllocatorRef, sbuf uintptr, sampleBufferOut *uintptr) int32 {
	if _cMSampleBufferCreateCopy == nil {
		panic("CoreMedia: symbol CMSampleBufferCreateCopy not loaded")
	}
	return _cMSampleBufferCreateCopy(allocator, sbuf, sampleBufferOut)
}

var _cMSampleBufferCreateCopyWithNewTiming func(allocator corefoundation.CFAllocatorRef, originalSBuf uintptr, numSampleTimingEntries int, sampleTimingArray *CMSampleTimingInfo, sampleBufferOut *uintptr) int32

// CMSampleBufferCreateCopyWithNewTiming creates a copy of a sample buffer with new timing information.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateCopyWithNewTiming(allocator:sampleBuffer:sampleTimingEntryCount:sampleTimingArray:sampleBufferOut:)
func CMSampleBufferCreateCopyWithNewTiming(allocator corefoundation.CFAllocatorRef, originalSBuf uintptr, numSampleTimingEntries int, sampleTimingArray *CMSampleTimingInfo, sampleBufferOut *uintptr) int32 {
	if _cMSampleBufferCreateCopyWithNewTiming == nil {
		panic("CoreMedia: symbol CMSampleBufferCreateCopyWithNewTiming not loaded")
	}
	return _cMSampleBufferCreateCopyWithNewTiming(allocator, originalSBuf, numSampleTimingEntries, sampleTimingArray, sampleBufferOut)
}

var _cMSampleBufferCreateForImageBuffer func(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, dataReady bool, makeDataReadyCallback CMSampleBufferMakeDataReadyCallback, makeDataReadyRefcon unsafe.Pointer, formatDescription uintptr, sampleTiming *CMSampleTimingInfo, sampleBufferOut *uintptr) int32

// CMSampleBufferCreateForImageBuffer creates a sample buffer with an image buffer and a callback to make the data ready for use.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateForImageBuffer(allocator:imageBuffer:dataReady:makeDataReadyCallback:refcon:formatDescription:sampleTiming:sampleBufferOut:)
func CMSampleBufferCreateForImageBuffer(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, dataReady bool, makeDataReadyCallback CMSampleBufferMakeDataReadyCallback, makeDataReadyRefcon unsafe.Pointer, formatDescription uintptr, sampleTiming *CMSampleTimingInfo, sampleBufferOut *uintptr) int32 {
	if _cMSampleBufferCreateForImageBuffer == nil {
		panic("CoreMedia: symbol CMSampleBufferCreateForImageBuffer not loaded")
	}
	return _cMSampleBufferCreateForImageBuffer(allocator, imageBuffer, dataReady, makeDataReadyCallback, makeDataReadyRefcon, formatDescription, sampleTiming, sampleBufferOut)
}

var _cMSampleBufferCreateForImageBufferWithMakeDataReadyHandler func(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, dataReady bool, formatDescription uintptr, sampleTiming *CMSampleTimingInfo, sampleBufferOut *uintptr, makeDataReadyHandler CMSampleBufferMakeDataReadyHandler) int32

// CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler creates a sample buffer with an image buffer and a handler to make the data ready for use.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler(_:_:_:_:_:_:_:)
func CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, dataReady bool, formatDescription uintptr, sampleTiming *CMSampleTimingInfo, sampleBufferOut *uintptr, makeDataReadyHandler CMSampleBufferMakeDataReadyHandler) int32 {
	if _cMSampleBufferCreateForImageBufferWithMakeDataReadyHandler == nil {
		panic("CoreMedia: symbol CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler not loaded")
	}
	return _cMSampleBufferCreateForImageBufferWithMakeDataReadyHandler(allocator, imageBuffer, dataReady, formatDescription, sampleTiming, sampleBufferOut, makeDataReadyHandler)
}

var _cMSampleBufferCreateForTaggedBufferGroup func(allocator corefoundation.CFAllocatorRef, taggedBufferGroup CMTaggedBufferGroupRef, sbufPTS CMTime, sbufDuration CMTime, formatDescription CMTaggedBufferGroupFormatDescriptionRef, sBufOut *uintptr) int32

// CMSampleBufferCreateForTaggedBufferGroup creates a new sample buffer from a tagged buffer group.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateForTaggedBufferGroup
func CMSampleBufferCreateForTaggedBufferGroup(allocator corefoundation.CFAllocatorRef, taggedBufferGroup CMTaggedBufferGroupRef, sbufPTS CMTime, sbufDuration CMTime, formatDescription CMTaggedBufferGroupFormatDescriptionRef, sBufOut *uintptr) int32 {
	if _cMSampleBufferCreateForTaggedBufferGroup == nil {
		panic("CoreMedia: symbol CMSampleBufferCreateForTaggedBufferGroup not loaded")
	}
	return _cMSampleBufferCreateForTaggedBufferGroup(allocator, taggedBufferGroup, sbufPTS, sbufDuration, formatDescription, sBufOut)
}

var _cMSampleBufferCreateReady func(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, formatDescription uintptr, numSamples int, numSampleTimingEntries int, sampleTimingArray *CMSampleTimingInfo, numSampleSizeEntries int, sampleSizeArray *uintptr, sampleBufferOut *uintptr) int32

// CMSampleBufferCreateReady creates a sample buffer with media data.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateReady(allocator:dataBuffer:formatDescription:sampleCount:sampleTimingEntryCount:sampleTimingArray:sampleSizeEntryCount:sampleSizeArray:sampleBufferOut:)
func CMSampleBufferCreateReady(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, formatDescription uintptr, numSamples int, numSampleTimingEntries int, sampleTimingArray *CMSampleTimingInfo, numSampleSizeEntries int, sampleSizeArray *uintptr, sampleBufferOut *uintptr) int32 {
	if _cMSampleBufferCreateReady == nil {
		panic("CoreMedia: symbol CMSampleBufferCreateReady not loaded")
	}
	return _cMSampleBufferCreateReady(allocator, dataBuffer, formatDescription, numSamples, numSampleTimingEntries, sampleTimingArray, numSampleSizeEntries, sampleSizeArray, sampleBufferOut)
}

var _cMSampleBufferCreateReadyWithImageBuffer func(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, formatDescription uintptr, sampleTiming *CMSampleTimingInfo, sampleBufferOut *uintptr) int32

// CMSampleBufferCreateReadyWithImageBuffer creates a sample buffer with image data.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateReadyWithImageBuffer(allocator:imageBuffer:formatDescription:sampleTiming:sampleBufferOut:)
func CMSampleBufferCreateReadyWithImageBuffer(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, formatDescription uintptr, sampleTiming *CMSampleTimingInfo, sampleBufferOut *uintptr) int32 {
	if _cMSampleBufferCreateReadyWithImageBuffer == nil {
		panic("CoreMedia: symbol CMSampleBufferCreateReadyWithImageBuffer not loaded")
	}
	return _cMSampleBufferCreateReadyWithImageBuffer(allocator, imageBuffer, formatDescription, sampleTiming, sampleBufferOut)
}

var _cMSampleBufferCreateWithMakeDataReadyHandler func(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, dataReady bool, formatDescription uintptr, numSamples int, numSampleTimingEntries int, sampleTimingArray *CMSampleTimingInfo, numSampleSizeEntries int, sampleSizeArray *uintptr, sampleBufferOut *uintptr, makeDataReadyHandler CMSampleBufferMakeDataReadyHandler) int32

// CMSampleBufferCreateWithMakeDataReadyHandler creates a sample buffer with a handler to make the data ready for use.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferCreateWithMakeDataReadyHandler(_:_:_:_:_:_:_:_:_:_:_:)
func CMSampleBufferCreateWithMakeDataReadyHandler(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, dataReady bool, formatDescription uintptr, numSamples int, numSampleTimingEntries int, sampleTimingArray *CMSampleTimingInfo, numSampleSizeEntries int, sampleSizeArray *uintptr, sampleBufferOut *uintptr, makeDataReadyHandler CMSampleBufferMakeDataReadyHandler) int32 {
	if _cMSampleBufferCreateWithMakeDataReadyHandler == nil {
		panic("CoreMedia: symbol CMSampleBufferCreateWithMakeDataReadyHandler not loaded")
	}
	return _cMSampleBufferCreateWithMakeDataReadyHandler(allocator, dataBuffer, dataReady, formatDescription, numSamples, numSampleTimingEntries, sampleTimingArray, numSampleSizeEntries, sampleSizeArray, sampleBufferOut, makeDataReadyHandler)
}

var _cMSampleBufferDataIsReady func(sbuf uintptr) bool

// CMSampleBufferDataIsReady returns a Boolean value that indicates whether the sample buffer’s data is ready for use.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferDataIsReady(_:)
func CMSampleBufferDataIsReady(sbuf uintptr) bool {
	if _cMSampleBufferDataIsReady == nil {
		panic("CoreMedia: symbol CMSampleBufferDataIsReady not loaded")
	}
	return _cMSampleBufferDataIsReady(sbuf)
}

var _cMSampleBufferGetAudioBufferListWithRetainedBlockBuffer func(sbuf uintptr, bufferListSizeNeededOut *uintptr, bufferListOut uintptr, bufferListSize uintptr, blockBufferStructureAllocator corefoundation.CFAllocatorRef, blockBufferBlockAllocator corefoundation.CFAllocatorRef, flags uint32, blockBufferOut *uintptr) int32

// CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer returns an audio buffer list that contains the media data.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer(_:bufferListSizeNeededOut:bufferListOut:bufferListSize:blockBufferAllocator:blockBufferMemoryAllocator:flags:blockBufferOut:)
func CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer(sbuf uintptr, bufferListSizeNeededOut *uintptr, bufferListOut uintptr, bufferListSize uintptr, blockBufferStructureAllocator corefoundation.CFAllocatorRef, blockBufferBlockAllocator corefoundation.CFAllocatorRef, flags uint32, blockBufferOut *uintptr) int32 {
	if _cMSampleBufferGetAudioBufferListWithRetainedBlockBuffer == nil {
		panic("CoreMedia: symbol CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer not loaded")
	}
	return _cMSampleBufferGetAudioBufferListWithRetainedBlockBuffer(sbuf, bufferListSizeNeededOut, bufferListOut, bufferListSize, blockBufferStructureAllocator, blockBufferBlockAllocator, flags, blockBufferOut)
}

var _cMSampleBufferGetAudioStreamPacketDescriptions func(sbuf uintptr, packetDescriptionsSize uintptr, packetDescriptionsOut uintptr, packetDescriptionsSizeNeededOut *uintptr) int32

// CMSampleBufferGetAudioStreamPacketDescriptions creates an array of audio stream packet descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetAudioStreamPacketDescriptions(_:allocatedSize:packetDescriptionsOut:packetDescriptionsSizeNeededOut:)
func CMSampleBufferGetAudioStreamPacketDescriptions(sbuf uintptr, packetDescriptionsSize uintptr, packetDescriptionsOut uintptr, packetDescriptionsSizeNeededOut *uintptr) int32 {
	if _cMSampleBufferGetAudioStreamPacketDescriptions == nil {
		panic("CoreMedia: symbol CMSampleBufferGetAudioStreamPacketDescriptions not loaded")
	}
	return _cMSampleBufferGetAudioStreamPacketDescriptions(sbuf, packetDescriptionsSize, packetDescriptionsOut, packetDescriptionsSizeNeededOut)
}

var _cMSampleBufferGetAudioStreamPacketDescriptionsPtr func(sbuf uintptr, packetDescriptionsPointerOut uintptr, packetDescriptionsSizeOut *uintptr) int32

// CMSampleBufferGetAudioStreamPacketDescriptionsPtr returns a pointer to a constant array of audio stream packet descriptions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetAudioStreamPacketDescriptionsPtr(_:packetDescriptionsPointerOut:sizeOut:)
func CMSampleBufferGetAudioStreamPacketDescriptionsPtr(sbuf uintptr, packetDescriptionsPointerOut uintptr, packetDescriptionsSizeOut *uintptr) int32 {
	if _cMSampleBufferGetAudioStreamPacketDescriptionsPtr == nil {
		panic("CoreMedia: symbol CMSampleBufferGetAudioStreamPacketDescriptionsPtr not loaded")
	}
	return _cMSampleBufferGetAudioStreamPacketDescriptionsPtr(sbuf, packetDescriptionsPointerOut, packetDescriptionsSizeOut)
}

var _cMSampleBufferGetDataBuffer func(sbuf uintptr) uintptr

// CMSampleBufferGetDataBuffer returns a block buffer that contains the media data.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetDataBuffer(_:)
func CMSampleBufferGetDataBuffer(sbuf uintptr) uintptr {
	if _cMSampleBufferGetDataBuffer == nil {
		panic("CoreMedia: symbol CMSampleBufferGetDataBuffer not loaded")
	}
	return _cMSampleBufferGetDataBuffer(sbuf)
}

var _cMSampleBufferGetDecodeTimeStamp func(sbuf uintptr) CMTime

// CMSampleBufferGetDecodeTimeStamp returns the decode timestamp that’s the earliest numerically of all the samples in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetDecodeTimeStamp(_:)
func CMSampleBufferGetDecodeTimeStamp(sbuf uintptr) CMTime {
	if _cMSampleBufferGetDecodeTimeStamp == nil {
		panic("CoreMedia: symbol CMSampleBufferGetDecodeTimeStamp not loaded")
	}
	return _cMSampleBufferGetDecodeTimeStamp(sbuf)
}

var _cMSampleBufferGetDuration func(sbuf uintptr) CMTime

// CMSampleBufferGetDuration returns the total duration of a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetDuration(_:)
func CMSampleBufferGetDuration(sbuf uintptr) CMTime {
	if _cMSampleBufferGetDuration == nil {
		panic("CoreMedia: symbol CMSampleBufferGetDuration not loaded")
	}
	return _cMSampleBufferGetDuration(sbuf)
}

var _cMSampleBufferGetFormatDescription func(sbuf uintptr) uintptr

// CMSampleBufferGetFormatDescription returns the format description of the samples in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetFormatDescription(_:)
func CMSampleBufferGetFormatDescription(sbuf uintptr) uintptr {
	if _cMSampleBufferGetFormatDescription == nil {
		panic("CoreMedia: symbol CMSampleBufferGetFormatDescription not loaded")
	}
	return _cMSampleBufferGetFormatDescription(sbuf)
}

var _cMSampleBufferGetImageBuffer func(sbuf uintptr) corevideo.CVImageBufferRef

// CMSampleBufferGetImageBuffer returns an image buffer that contains the media data.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetImageBuffer(_:)
func CMSampleBufferGetImageBuffer(sbuf uintptr) corevideo.CVImageBufferRef {
	if _cMSampleBufferGetImageBuffer == nil {
		panic("CoreMedia: symbol CMSampleBufferGetImageBuffer not loaded")
	}
	return _cMSampleBufferGetImageBuffer(sbuf)
}

var _cMSampleBufferGetNumSamples func(sbuf uintptr) int

// CMSampleBufferGetNumSamples returns the number of media samples in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetNumSamples(_:)
func CMSampleBufferGetNumSamples(sbuf uintptr) int {
	if _cMSampleBufferGetNumSamples == nil {
		panic("CoreMedia: symbol CMSampleBufferGetNumSamples not loaded")
	}
	return _cMSampleBufferGetNumSamples(sbuf)
}

var _cMSampleBufferGetOutputDecodeTimeStamp func(sbuf uintptr) CMTime

// CMSampleBufferGetOutputDecodeTimeStamp returns the output decode timestamp of a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetOutputDecodeTimeStamp(_:)
func CMSampleBufferGetOutputDecodeTimeStamp(sbuf uintptr) CMTime {
	if _cMSampleBufferGetOutputDecodeTimeStamp == nil {
		panic("CoreMedia: symbol CMSampleBufferGetOutputDecodeTimeStamp not loaded")
	}
	return _cMSampleBufferGetOutputDecodeTimeStamp(sbuf)
}

var _cMSampleBufferGetOutputDuration func(sbuf uintptr) CMTime

// CMSampleBufferGetOutputDuration returns the output duration of a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetOutputDuration(_:)
func CMSampleBufferGetOutputDuration(sbuf uintptr) CMTime {
	if _cMSampleBufferGetOutputDuration == nil {
		panic("CoreMedia: symbol CMSampleBufferGetOutputDuration not loaded")
	}
	return _cMSampleBufferGetOutputDuration(sbuf)
}

var _cMSampleBufferGetOutputPresentationTimeStamp func(sbuf uintptr) CMTime

// CMSampleBufferGetOutputPresentationTimeStamp returns the output presentation timestamp of a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetOutputPresentationTimeStamp(_:)
func CMSampleBufferGetOutputPresentationTimeStamp(sbuf uintptr) CMTime {
	if _cMSampleBufferGetOutputPresentationTimeStamp == nil {
		panic("CoreMedia: symbol CMSampleBufferGetOutputPresentationTimeStamp not loaded")
	}
	return _cMSampleBufferGetOutputPresentationTimeStamp(sbuf)
}

var _cMSampleBufferGetOutputSampleTimingInfoArray func(sbuf uintptr, timingArrayEntries int, timingArrayOut *CMSampleTimingInfo, timingArrayEntriesNeededOut *int) int32

// CMSampleBufferGetOutputSampleTimingInfoArray retrieves an array of output timing information structures that represents each sample in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetOutputSampleTimingInfoArray(_:entryCount:arrayToFill:entriesNeededOut:)
func CMSampleBufferGetOutputSampleTimingInfoArray(sbuf uintptr, timingArrayEntries int, timingArrayOut *CMSampleTimingInfo, timingArrayEntriesNeededOut *int) int32 {
	if _cMSampleBufferGetOutputSampleTimingInfoArray == nil {
		panic("CoreMedia: symbol CMSampleBufferGetOutputSampleTimingInfoArray not loaded")
	}
	return _cMSampleBufferGetOutputSampleTimingInfoArray(sbuf, timingArrayEntries, timingArrayOut, timingArrayEntriesNeededOut)
}

var _cMSampleBufferGetPresentationTimeStamp func(sbuf uintptr) CMTime

// CMSampleBufferGetPresentationTimeStamp returns the presentation timestamp that’s the earliest numerically of all the samples in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetPresentationTimeStamp(_:)
func CMSampleBufferGetPresentationTimeStamp(sbuf uintptr) CMTime {
	if _cMSampleBufferGetPresentationTimeStamp == nil {
		panic("CoreMedia: symbol CMSampleBufferGetPresentationTimeStamp not loaded")
	}
	return _cMSampleBufferGetPresentationTimeStamp(sbuf)
}

var _cMSampleBufferGetSampleAttachmentsArray func(sbuf uintptr, createIfNecessary bool) corefoundation.CFArrayRef

// CMSampleBufferGetSampleAttachmentsArray retrieves an array of sample attachment dictionaries that represents each sample in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetSampleAttachmentsArray(_:createIfNecessary:)
func CMSampleBufferGetSampleAttachmentsArray(sbuf uintptr, createIfNecessary bool) corefoundation.CFArrayRef {
	if _cMSampleBufferGetSampleAttachmentsArray == nil {
		panic("CoreMedia: symbol CMSampleBufferGetSampleAttachmentsArray not loaded")
	}
	return _cMSampleBufferGetSampleAttachmentsArray(sbuf, createIfNecessary)
}

var _cMSampleBufferGetSampleSize func(sbuf uintptr, sampleIndex int) uintptr

// CMSampleBufferGetSampleSize returns the size in bytes of a specified sample in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetSampleSize(_:at:)
func CMSampleBufferGetSampleSize(sbuf uintptr, sampleIndex int) uintptr {
	if _cMSampleBufferGetSampleSize == nil {
		panic("CoreMedia: symbol CMSampleBufferGetSampleSize not loaded")
	}
	return _cMSampleBufferGetSampleSize(sbuf, sampleIndex)
}

var _cMSampleBufferGetSampleSizeArray func(sbuf uintptr, sizeArrayEntries int, sizeArrayOut *uintptr, sizeArrayEntriesNeededOut *int) int32

// CMSampleBufferGetSampleSizeArray retrieves an array of sample sizes that represents each sample in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetSampleSizeArray(_:entryCount:arrayToFill:entriesNeededOut:)
func CMSampleBufferGetSampleSizeArray(sbuf uintptr, sizeArrayEntries int, sizeArrayOut *uintptr, sizeArrayEntriesNeededOut *int) int32 {
	if _cMSampleBufferGetSampleSizeArray == nil {
		panic("CoreMedia: symbol CMSampleBufferGetSampleSizeArray not loaded")
	}
	return _cMSampleBufferGetSampleSizeArray(sbuf, sizeArrayEntries, sizeArrayOut, sizeArrayEntriesNeededOut)
}

var _cMSampleBufferGetSampleTimingInfo func(sbuf uintptr, sampleIndex int, timingInfoOut *CMSampleTimingInfo) int32

// CMSampleBufferGetSampleTimingInfo retrieves a timing information structure that describes a specified sample in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetSampleTimingInfo(_:at:timingInfoOut:)
func CMSampleBufferGetSampleTimingInfo(sbuf uintptr, sampleIndex int, timingInfoOut *CMSampleTimingInfo) int32 {
	if _cMSampleBufferGetSampleTimingInfo == nil {
		panic("CoreMedia: symbol CMSampleBufferGetSampleTimingInfo not loaded")
	}
	return _cMSampleBufferGetSampleTimingInfo(sbuf, sampleIndex, timingInfoOut)
}

var _cMSampleBufferGetSampleTimingInfoArray func(sbuf uintptr, numSampleTimingEntries int, timingArrayOut *CMSampleTimingInfo, timingArrayEntriesNeededOut *int) int32

// CMSampleBufferGetSampleTimingInfoArray retrieves an array of sample timing information structures that represents each sample in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetSampleTimingInfoArray(_:entryCount:arrayToFill:entriesNeededOut:)
func CMSampleBufferGetSampleTimingInfoArray(sbuf uintptr, numSampleTimingEntries int, timingArrayOut *CMSampleTimingInfo, timingArrayEntriesNeededOut *int) int32 {
	if _cMSampleBufferGetSampleTimingInfoArray == nil {
		panic("CoreMedia: symbol CMSampleBufferGetSampleTimingInfoArray not loaded")
	}
	return _cMSampleBufferGetSampleTimingInfoArray(sbuf, numSampleTimingEntries, timingArrayOut, timingArrayEntriesNeededOut)
}

var _cMSampleBufferGetTaggedBufferGroup func(sbuf uintptr) CMTaggedBufferGroupRef

// CMSampleBufferGetTaggedBufferGroup gets the tagged buffer group of a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetTaggedBufferGroup
func CMSampleBufferGetTaggedBufferGroup(sbuf uintptr) CMTaggedBufferGroupRef {
	if _cMSampleBufferGetTaggedBufferGroup == nil {
		panic("CoreMedia: symbol CMSampleBufferGetTaggedBufferGroup not loaded")
	}
	return _cMSampleBufferGetTaggedBufferGroup(sbuf)
}

var _cMSampleBufferGetTotalSampleSize func(sbuf uintptr) uintptr

// CMSampleBufferGetTotalSampleSize returns the total size in bytes of sample data in a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetTotalSampleSize(_:)
func CMSampleBufferGetTotalSampleSize(sbuf uintptr) uintptr {
	if _cMSampleBufferGetTotalSampleSize == nil {
		panic("CoreMedia: symbol CMSampleBufferGetTotalSampleSize not loaded")
	}
	return _cMSampleBufferGetTotalSampleSize(sbuf)
}

var _cMSampleBufferGetTypeID func() uint

// CMSampleBufferGetTypeID returns the type identifier of sample buffer objects.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferGetTypeID()
func CMSampleBufferGetTypeID() uint {
	if _cMSampleBufferGetTypeID == nil {
		panic("CoreMedia: symbol CMSampleBufferGetTypeID not loaded")
	}
	return _cMSampleBufferGetTypeID()
}

var _cMSampleBufferHasDataFailed func(sbuf uintptr, statusOut *int32) bool

// CMSampleBufferHasDataFailed returns a Boolean value that indicates whether the sample buffer’s data loading request failed.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferHasDataFailed(_:statusOut:)
func CMSampleBufferHasDataFailed(sbuf uintptr, statusOut *int32) bool {
	if _cMSampleBufferHasDataFailed == nil {
		panic("CoreMedia: symbol CMSampleBufferHasDataFailed not loaded")
	}
	return _cMSampleBufferHasDataFailed(sbuf, statusOut)
}

var _cMSampleBufferInvalidate func(sbuf uintptr) int32

// CMSampleBufferInvalidate invalidates a sample buffer by calling its invalidation callback.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferInvalidate(_:)
func CMSampleBufferInvalidate(sbuf uintptr) int32 {
	if _cMSampleBufferInvalidate == nil {
		panic("CoreMedia: symbol CMSampleBufferInvalidate not loaded")
	}
	return _cMSampleBufferInvalidate(sbuf)
}

var _cMSampleBufferIsValid func(sbuf uintptr) bool

// CMSampleBufferIsValid returns a Boolean value that indicates whether a sample buffer is valid.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferIsValid(_:)
func CMSampleBufferIsValid(sbuf uintptr) bool {
	if _cMSampleBufferIsValid == nil {
		panic("CoreMedia: symbol CMSampleBufferIsValid not loaded")
	}
	return _cMSampleBufferIsValid(sbuf)
}

var _cMSampleBufferMakeDataReady func(sbuf uintptr) int32

// CMSampleBufferMakeDataReady makes the sample buffer’s data ready for use by invoking its callback to load the data.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferMakeDataReady(_:)
func CMSampleBufferMakeDataReady(sbuf uintptr) int32 {
	if _cMSampleBufferMakeDataReady == nil {
		panic("CoreMedia: symbol CMSampleBufferMakeDataReady not loaded")
	}
	return _cMSampleBufferMakeDataReady(sbuf)
}

var _cMSampleBufferSetDataBuffer func(sbuf uintptr, dataBuffer uintptr) int32

// CMSampleBufferSetDataBuffer sets a block buffer of media data on a sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetDataBuffer(_:newValue:)
func CMSampleBufferSetDataBuffer(sbuf uintptr, dataBuffer uintptr) int32 {
	if _cMSampleBufferSetDataBuffer == nil {
		panic("CoreMedia: symbol CMSampleBufferSetDataBuffer not loaded")
	}
	return _cMSampleBufferSetDataBuffer(sbuf, dataBuffer)
}

var _cMSampleBufferSetDataBufferFromAudioBufferList func(sbuf uintptr, blockBufferStructureAllocator corefoundation.CFAllocatorRef, blockBufferBlockAllocator corefoundation.CFAllocatorRef, flags uint32, bufferList uintptr) int32

// CMSampleBufferSetDataBufferFromAudioBufferList creates a block buffer that contains a copy of the data from an audio buffer list.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetDataBufferFromAudioBufferList(_:blockBufferAllocator:blockBufferMemoryAllocator:flags:bufferList:)
func CMSampleBufferSetDataBufferFromAudioBufferList(sbuf uintptr, blockBufferStructureAllocator corefoundation.CFAllocatorRef, blockBufferBlockAllocator corefoundation.CFAllocatorRef, flags uint32, bufferList uintptr) int32 {
	if _cMSampleBufferSetDataBufferFromAudioBufferList == nil {
		panic("CoreMedia: symbol CMSampleBufferSetDataBufferFromAudioBufferList not loaded")
	}
	return _cMSampleBufferSetDataBufferFromAudioBufferList(sbuf, blockBufferStructureAllocator, blockBufferBlockAllocator, flags, bufferList)
}

var _cMSampleBufferSetDataFailed func(sbuf uintptr, status int32) int32

// CMSampleBufferSetDataFailed marks the sample buffer’s data as failed to indicate that it won’t become ready.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetDataFailed(_:status:)
func CMSampleBufferSetDataFailed(sbuf uintptr, status int32) int32 {
	if _cMSampleBufferSetDataFailed == nil {
		panic("CoreMedia: symbol CMSampleBufferSetDataFailed not loaded")
	}
	return _cMSampleBufferSetDataFailed(sbuf, status)
}

var _cMSampleBufferSetDataReady func(sbuf uintptr) int32

// CMSampleBufferSetDataReady marks a sample buffer’s data as ready for use.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetDataReady(_:)
func CMSampleBufferSetDataReady(sbuf uintptr) int32 {
	if _cMSampleBufferSetDataReady == nil {
		panic("CoreMedia: symbol CMSampleBufferSetDataReady not loaded")
	}
	return _cMSampleBufferSetDataReady(sbuf)
}

var _cMSampleBufferSetInvalidateCallback func(sbuf uintptr, invalidateCallback CMSampleBufferInvalidateCallback, invalidateRefCon uint64) int32

// CMSampleBufferSetInvalidateCallback sets the sample buffer’s invalidation callback.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetInvalidateCallback(_:callback:refcon:)
func CMSampleBufferSetInvalidateCallback(sbuf uintptr, invalidateCallback CMSampleBufferInvalidateCallback, invalidateRefCon uint64) int32 {
	if _cMSampleBufferSetInvalidateCallback == nil {
		panic("CoreMedia: symbol CMSampleBufferSetInvalidateCallback not loaded")
	}
	return _cMSampleBufferSetInvalidateCallback(sbuf, invalidateCallback, invalidateRefCon)
}

var _cMSampleBufferSetInvalidateHandler func(sbuf uintptr, invalidateHandler CMSampleBufferInvalidateHandler) int32

// CMSampleBufferSetInvalidateHandler sets the sample buffer’s invalidation handler.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetInvalidateHandler(_:invalidateHandler:)
func CMSampleBufferSetInvalidateHandler(sbuf uintptr, invalidateHandler CMSampleBufferInvalidateHandler) int32 {
	if _cMSampleBufferSetInvalidateHandler == nil {
		panic("CoreMedia: symbol CMSampleBufferSetInvalidateHandler not loaded")
	}
	return _cMSampleBufferSetInvalidateHandler(sbuf, invalidateHandler)
}

var _cMSampleBufferSetOutputPresentationTimeStamp func(sbuf uintptr, outputPresentationTimeStamp CMTime) int32

// CMSampleBufferSetOutputPresentationTimeStamp sets an output presentation timestamp to use in place of a calculated value.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferSetOutputPresentationTimeStamp(_:newValue:)
func CMSampleBufferSetOutputPresentationTimeStamp(sbuf uintptr, outputPresentationTimeStamp CMTime) int32 {
	if _cMSampleBufferSetOutputPresentationTimeStamp == nil {
		panic("CoreMedia: symbol CMSampleBufferSetOutputPresentationTimeStamp not loaded")
	}
	return _cMSampleBufferSetOutputPresentationTimeStamp(sbuf, outputPresentationTimeStamp)
}

var _cMSampleBufferTrackDataReadiness func(sbuf uintptr, sampleBufferToTrack uintptr) int32

// CMSampleBufferTrackDataReadiness associates a sample buffer’s data readiness with that of another sample buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSampleBufferTrackDataReadiness(_:sampleBufferToTrack:)
func CMSampleBufferTrackDataReadiness(sbuf uintptr, sampleBufferToTrack uintptr) int32 {
	if _cMSampleBufferTrackDataReadiness == nil {
		panic("CoreMedia: symbol CMSampleBufferTrackDataReadiness not loaded")
	}
	return _cMSampleBufferTrackDataReadiness(sbuf, sampleBufferToTrack)
}

var _cMSetAttachment func(target CMAttachmentBearerRef, key corefoundation.CFStringRef, value corefoundation.CFTypeRef, attachmentMode CMAttachmentMode)

// CMSetAttachment sets or adds an attachment to an attachment bearer object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSetAttachment(_:key:value:attachmentMode:)
func CMSetAttachment(target CMAttachmentBearerRef, key corefoundation.CFStringRef, value corefoundation.CFTypeRef, attachmentMode CMAttachmentMode) {
	if _cMSetAttachment == nil {
		panic("CoreMedia: symbol CMSetAttachment not loaded")
	}
	_cMSetAttachment(target, key, value, attachmentMode)
}

var _cMSetAttachments func(target CMAttachmentBearerRef, theAttachments corefoundation.CFDictionaryRef, attachmentMode CMAttachmentMode)

// CMSetAttachments sets a dictionary of attachments on an attachment bearer object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSetAttachments(_:attachments:attachmentMode:)
func CMSetAttachments(target CMAttachmentBearerRef, theAttachments corefoundation.CFDictionaryRef, attachmentMode CMAttachmentMode) {
	if _cMSetAttachments == nil {
		panic("CoreMedia: symbol CMSetAttachments not loaded")
	}
	_cMSetAttachments(target, theAttachments, attachmentMode)
}

var _cMSimpleQueueCreate func(allocator corefoundation.CFAllocatorRef, capacity int32, queueOut *CMSimpleQueueRef) int32

// CMSimpleQueueCreate creates a queue that has the specified capacity.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueCreate(allocator:capacity:queueOut:)
func CMSimpleQueueCreate(allocator corefoundation.CFAllocatorRef, capacity int32, queueOut *CMSimpleQueueRef) int32 {
	if _cMSimpleQueueCreate == nil {
		panic("CoreMedia: symbol CMSimpleQueueCreate not loaded")
	}
	return _cMSimpleQueueCreate(allocator, capacity, queueOut)
}

var _cMSimpleQueueDequeue func(queue CMSimpleQueueRef) unsafe.Pointer

// CMSimpleQueueDequeue dequeues an element from the queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueDequeue(_:)
func CMSimpleQueueDequeue(queue CMSimpleQueueRef) unsafe.Pointer {
	if _cMSimpleQueueDequeue == nil {
		panic("CoreMedia: symbol CMSimpleQueueDequeue not loaded")
	}
	return _cMSimpleQueueDequeue(queue)
}

var _cMSimpleQueueEnqueue func(queue CMSimpleQueueRef, element unsafe.Pointer) int32

// CMSimpleQueueEnqueue enqueues an element in the queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueEnqueue(_:element:)
func CMSimpleQueueEnqueue(queue CMSimpleQueueRef, element unsafe.Pointer) int32 {
	if _cMSimpleQueueEnqueue == nil {
		panic("CoreMedia: symbol CMSimpleQueueEnqueue not loaded")
	}
	return _cMSimpleQueueEnqueue(queue, element)
}

var _cMSimpleQueueGetCapacity func(queue CMSimpleQueueRef) int32

// CMSimpleQueueGetCapacity returns the number of elements that the queue can hold.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueGetCapacity(_:)
func CMSimpleQueueGetCapacity(queue CMSimpleQueueRef) int32 {
	if _cMSimpleQueueGetCapacity == nil {
		panic("CoreMedia: symbol CMSimpleQueueGetCapacity not loaded")
	}
	return _cMSimpleQueueGetCapacity(queue)
}

var _cMSimpleQueueGetCount func(queue CMSimpleQueueRef) int32

// CMSimpleQueueGetCount returns the number of elements currently in the queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueGetCount(_:)
func CMSimpleQueueGetCount(queue CMSimpleQueueRef) int32 {
	if _cMSimpleQueueGetCount == nil {
		panic("CoreMedia: symbol CMSimpleQueueGetCount not loaded")
	}
	return _cMSimpleQueueGetCount(queue)
}

var _cMSimpleQueueGetHead func(queue CMSimpleQueueRef) unsafe.Pointer

// CMSimpleQueueGetHead returns the element at the head of the queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueGetHead(_:)
func CMSimpleQueueGetHead(queue CMSimpleQueueRef) unsafe.Pointer {
	if _cMSimpleQueueGetHead == nil {
		panic("CoreMedia: symbol CMSimpleQueueGetHead not loaded")
	}
	return _cMSimpleQueueGetHead(queue)
}

var _cMSimpleQueueGetTypeID func() uint

// CMSimpleQueueGetTypeID returns the type identifier of sample buffer objects.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueGetTypeID()
func CMSimpleQueueGetTypeID() uint {
	if _cMSimpleQueueGetTypeID == nil {
		panic("CoreMedia: symbol CMSimpleQueueGetTypeID not loaded")
	}
	return _cMSimpleQueueGetTypeID()
}

var _cMSimpleQueueReset func(queue CMSimpleQueueRef) int32

// CMSimpleQueueReset resets the queue.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSimpleQueueReset(_:)
func CMSimpleQueueReset(queue CMSimpleQueueRef) int32 {
	if _cMSimpleQueueReset == nil {
		panic("CoreMedia: symbol CMSimpleQueueReset not loaded")
	}
	return _cMSimpleQueueReset(queue)
}

var _cMSwapBigEndianClosedCaptionDescriptionToHost func(closedCaptionDescriptionData *uint8, closedCaptionDescriptionSize uintptr) int32

// CMSwapBigEndianClosedCaptionDescriptionToHost converts a closed caption description structure from big-endian to host-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapBigEndianClosedCaptionDescriptionToHost(_:_:)
func CMSwapBigEndianClosedCaptionDescriptionToHost(closedCaptionDescriptionData *uint8, closedCaptionDescriptionSize uintptr) int32 {
	if _cMSwapBigEndianClosedCaptionDescriptionToHost == nil {
		panic("CoreMedia: symbol CMSwapBigEndianClosedCaptionDescriptionToHost not loaded")
	}
	return _cMSwapBigEndianClosedCaptionDescriptionToHost(closedCaptionDescriptionData, closedCaptionDescriptionSize)
}

var _cMSwapBigEndianImageDescriptionToHost func(imageDescriptionData *uint8, imageDescriptionSize uintptr) int32

// CMSwapBigEndianImageDescriptionToHost converts an image description data structure from big-endian to host-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapBigEndianImageDescriptionToHost(_:_:)
func CMSwapBigEndianImageDescriptionToHost(imageDescriptionData *uint8, imageDescriptionSize uintptr) int32 {
	if _cMSwapBigEndianImageDescriptionToHost == nil {
		panic("CoreMedia: symbol CMSwapBigEndianImageDescriptionToHost not loaded")
	}
	return _cMSwapBigEndianImageDescriptionToHost(imageDescriptionData, imageDescriptionSize)
}

var _cMSwapBigEndianMetadataDescriptionToHost func(metadataDescriptionData *uint8, metadataDescriptionSize uintptr) int32

// CMSwapBigEndianMetadataDescriptionToHost converts a metadata description data structure from big-endian to host-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapBigEndianMetadataDescriptionToHost(_:_:)
func CMSwapBigEndianMetadataDescriptionToHost(metadataDescriptionData *uint8, metadataDescriptionSize uintptr) int32 {
	if _cMSwapBigEndianMetadataDescriptionToHost == nil {
		panic("CoreMedia: symbol CMSwapBigEndianMetadataDescriptionToHost not loaded")
	}
	return _cMSwapBigEndianMetadataDescriptionToHost(metadataDescriptionData, metadataDescriptionSize)
}

var _cMSwapBigEndianSoundDescriptionToHost func(soundDescriptionData *uint8, soundDescriptionSize uintptr) int32

// CMSwapBigEndianSoundDescriptionToHost converts a sound description data structure from big-endian to host-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapBigEndianSoundDescriptionToHost(_:_:)
func CMSwapBigEndianSoundDescriptionToHost(soundDescriptionData *uint8, soundDescriptionSize uintptr) int32 {
	if _cMSwapBigEndianSoundDescriptionToHost == nil {
		panic("CoreMedia: symbol CMSwapBigEndianSoundDescriptionToHost not loaded")
	}
	return _cMSwapBigEndianSoundDescriptionToHost(soundDescriptionData, soundDescriptionSize)
}

var _cMSwapBigEndianTextDescriptionToHost func(textDescriptionData *uint8, textDescriptionSize uintptr) int32

// CMSwapBigEndianTextDescriptionToHost converts a text description structure from big-endian to host-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapBigEndianTextDescriptionToHost(_:_:)
func CMSwapBigEndianTextDescriptionToHost(textDescriptionData *uint8, textDescriptionSize uintptr) int32 {
	if _cMSwapBigEndianTextDescriptionToHost == nil {
		panic("CoreMedia: symbol CMSwapBigEndianTextDescriptionToHost not loaded")
	}
	return _cMSwapBigEndianTextDescriptionToHost(textDescriptionData, textDescriptionSize)
}

var _cMSwapBigEndianTimeCodeDescriptionToHost func(timeCodeDescriptionData *uint8, timeCodeDescriptionSize uintptr) int32

// CMSwapBigEndianTimeCodeDescriptionToHost converts a time code description data structure from big-endian to host-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapBigEndianTimeCodeDescriptionToHost(_:_:)
func CMSwapBigEndianTimeCodeDescriptionToHost(timeCodeDescriptionData *uint8, timeCodeDescriptionSize uintptr) int32 {
	if _cMSwapBigEndianTimeCodeDescriptionToHost == nil {
		panic("CoreMedia: symbol CMSwapBigEndianTimeCodeDescriptionToHost not loaded")
	}
	return _cMSwapBigEndianTimeCodeDescriptionToHost(timeCodeDescriptionData, timeCodeDescriptionSize)
}

var _cMSwapHostEndianClosedCaptionDescriptionToBig func(closedCaptionDescriptionData *uint8, closedCaptionDescriptionSize uintptr) int32

// CMSwapHostEndianClosedCaptionDescriptionToBig converts a closed caption description structure from host-endian to big-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapHostEndianClosedCaptionDescriptionToBig(_:_:)
func CMSwapHostEndianClosedCaptionDescriptionToBig(closedCaptionDescriptionData *uint8, closedCaptionDescriptionSize uintptr) int32 {
	if _cMSwapHostEndianClosedCaptionDescriptionToBig == nil {
		panic("CoreMedia: symbol CMSwapHostEndianClosedCaptionDescriptionToBig not loaded")
	}
	return _cMSwapHostEndianClosedCaptionDescriptionToBig(closedCaptionDescriptionData, closedCaptionDescriptionSize)
}

var _cMSwapHostEndianImageDescriptionToBig func(imageDescriptionData *uint8, imageDescriptionSize uintptr) int32

// CMSwapHostEndianImageDescriptionToBig converts an image description data structure from host-endian to big-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapHostEndianImageDescriptionToBig(_:_:)
func CMSwapHostEndianImageDescriptionToBig(imageDescriptionData *uint8, imageDescriptionSize uintptr) int32 {
	if _cMSwapHostEndianImageDescriptionToBig == nil {
		panic("CoreMedia: symbol CMSwapHostEndianImageDescriptionToBig not loaded")
	}
	return _cMSwapHostEndianImageDescriptionToBig(imageDescriptionData, imageDescriptionSize)
}

var _cMSwapHostEndianMetadataDescriptionToBig func(metadataDescriptionData *uint8, metadataDescriptionSize uintptr) int32

// CMSwapHostEndianMetadataDescriptionToBig converts a metadata description data structure from host-endian to big-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapHostEndianMetadataDescriptionToBig(_:_:)
func CMSwapHostEndianMetadataDescriptionToBig(metadataDescriptionData *uint8, metadataDescriptionSize uintptr) int32 {
	if _cMSwapHostEndianMetadataDescriptionToBig == nil {
		panic("CoreMedia: symbol CMSwapHostEndianMetadataDescriptionToBig not loaded")
	}
	return _cMSwapHostEndianMetadataDescriptionToBig(metadataDescriptionData, metadataDescriptionSize)
}

var _cMSwapHostEndianSoundDescriptionToBig func(soundDescriptionData *uint8, soundDescriptionSize uintptr) int32

// CMSwapHostEndianSoundDescriptionToBig converts a sound description data structure from host-endian to big-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapHostEndianSoundDescriptionToBig(_:_:)
func CMSwapHostEndianSoundDescriptionToBig(soundDescriptionData *uint8, soundDescriptionSize uintptr) int32 {
	if _cMSwapHostEndianSoundDescriptionToBig == nil {
		panic("CoreMedia: symbol CMSwapHostEndianSoundDescriptionToBig not loaded")
	}
	return _cMSwapHostEndianSoundDescriptionToBig(soundDescriptionData, soundDescriptionSize)
}

var _cMSwapHostEndianTextDescriptionToBig func(textDescriptionData *uint8, textDescriptionSize uintptr) int32

// CMSwapHostEndianTextDescriptionToBig converts a text description structure from host-endian to big-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapHostEndianTextDescriptionToBig(_:_:)
func CMSwapHostEndianTextDescriptionToBig(textDescriptionData *uint8, textDescriptionSize uintptr) int32 {
	if _cMSwapHostEndianTextDescriptionToBig == nil {
		panic("CoreMedia: symbol CMSwapHostEndianTextDescriptionToBig not loaded")
	}
	return _cMSwapHostEndianTextDescriptionToBig(textDescriptionData, textDescriptionSize)
}

var _cMSwapHostEndianTimeCodeDescriptionToBig func(timeCodeDescriptionData *uint8, timeCodeDescriptionSize uintptr) int32

// CMSwapHostEndianTimeCodeDescriptionToBig converts a time code description data structure from host-endian to big-endian, in place.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSwapHostEndianTimeCodeDescriptionToBig(_:_:)
func CMSwapHostEndianTimeCodeDescriptionToBig(timeCodeDescriptionData *uint8, timeCodeDescriptionSize uintptr) int32 {
	if _cMSwapHostEndianTimeCodeDescriptionToBig == nil {
		panic("CoreMedia: symbol CMSwapHostEndianTimeCodeDescriptionToBig not loaded")
	}
	return _cMSwapHostEndianTimeCodeDescriptionToBig(timeCodeDescriptionData, timeCodeDescriptionSize)
}

var _cMSyncConvertTime func(time CMTime, fromClockOrTimebase CMClockOrTimebaseRef, toClockOrTimebase CMClockOrTimebaseRef) CMTime

// CMSyncConvertTime converts a time from one timebase or clock to another timebase or clock.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSyncConvertTime(_:from:to:)
func CMSyncConvertTime(time CMTime, fromClockOrTimebase CMClockOrTimebaseRef, toClockOrTimebase CMClockOrTimebaseRef) CMTime {
	if _cMSyncConvertTime == nil {
		panic("CoreMedia: symbol CMSyncConvertTime not loaded")
	}
	return _cMSyncConvertTime(time, fromClockOrTimebase, toClockOrTimebase)
}

var _cMSyncGetRelativeRate func(ofClockOrTimebase CMClockOrTimebaseRef, relativeToClockOrTimebase CMClockOrTimebaseRef) float64

// CMSyncGetRelativeRate returns the relative rate of one timebase or clock relative to another timebase or clock.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSyncGetRelativeRate(_:relativeTo:)
func CMSyncGetRelativeRate(ofClockOrTimebase CMClockOrTimebaseRef, relativeToClockOrTimebase CMClockOrTimebaseRef) float64 {
	if _cMSyncGetRelativeRate == nil {
		panic("CoreMedia: symbol CMSyncGetRelativeRate not loaded")
	}
	return _cMSyncGetRelativeRate(ofClockOrTimebase, relativeToClockOrTimebase)
}

var _cMSyncGetRelativeRateAndAnchorTime func(ofClockOrTimebase CMClockOrTimebaseRef, relativeToClockOrTimebase CMClockOrTimebaseRef, outRelativeRate *float64, outOfClockOrTimebaseAnchorTime *CMTime, outRelativeToClockOrTimebaseAnchorTime *CMTime) int32

// CMSyncGetRelativeRateAndAnchorTime returns the relative rate of one timebase or clock relative to another timebase or clock and the times of each timebase or clock at which the relative rate went into effect.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSyncGetRelativeRateAndAnchorTime(_:relativeTo:relativeRateOut:anchorTimeOut:relativeToAnchorTimeOut:)
func CMSyncGetRelativeRateAndAnchorTime(ofClockOrTimebase CMClockOrTimebaseRef, relativeToClockOrTimebase CMClockOrTimebaseRef, outRelativeRate *float64, outOfClockOrTimebaseAnchorTime *CMTime, outRelativeToClockOrTimebaseAnchorTime *CMTime) int32 {
	if _cMSyncGetRelativeRateAndAnchorTime == nil {
		panic("CoreMedia: symbol CMSyncGetRelativeRateAndAnchorTime not loaded")
	}
	return _cMSyncGetRelativeRateAndAnchorTime(ofClockOrTimebase, relativeToClockOrTimebase, outRelativeRate, outOfClockOrTimebaseAnchorTime, outRelativeToClockOrTimebaseAnchorTime)
}

var _cMSyncGetTime func(clockOrTimebase CMClockOrTimebaseRef) CMTime

// CMSyncGetTime returns the time from a clock or timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSyncGetTime(_:)
func CMSyncGetTime(clockOrTimebase CMClockOrTimebaseRef) CMTime {
	if _cMSyncGetTime == nil {
		panic("CoreMedia: symbol CMSyncGetTime not loaded")
	}
	return _cMSyncGetTime(clockOrTimebase)
}

var _cMSyncMightDrift func(clockOrTimebase1 CMClockOrTimebaseRef, clockOrTimebase2 CMClockOrTimebaseRef) bool

// CMSyncMightDrift returns a Boolean value that indicates whether it’s possible for one timebase or clock to drift relative to the other.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMSyncMightDrift(_:_:)
func CMSyncMightDrift(clockOrTimebase1 CMClockOrTimebaseRef, clockOrTimebase2 CMClockOrTimebaseRef) bool {
	if _cMSyncMightDrift == nil {
		panic("CoreMedia: symbol CMSyncMightDrift not loaded")
	}
	return _cMSyncMightDrift(clockOrTimebase1, clockOrTimebase2)
}

var _cMTagCollectionAddTag func(tagCollection CMMutableTagCollectionRef, tagToAdd CMTag) int32

// CMTagCollectionAddTag adds a new tag to an existing collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionAddTag
func CMTagCollectionAddTag(tagCollection CMMutableTagCollectionRef, tagToAdd CMTag) int32 {
	if _cMTagCollectionAddTag == nil {
		panic("CoreMedia: symbol CMTagCollectionAddTag not loaded")
	}
	return _cMTagCollectionAddTag(tagCollection, tagToAdd)
}

var _cMTagCollectionAddTagsFromArray func(tagCollection CMMutableTagCollectionRef, tags *CMTag, tagCount int) int32

// CMTagCollectionAddTagsFromArray adds the tags contained in a C-style array to a tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionAddTagsFromArray
func CMTagCollectionAddTagsFromArray(tagCollection CMMutableTagCollectionRef, tags *CMTag, tagCount int) int32 {
	if _cMTagCollectionAddTagsFromArray == nil {
		panic("CoreMedia: symbol CMTagCollectionAddTagsFromArray not loaded")
	}
	return _cMTagCollectionAddTagsFromArray(tagCollection, tags, tagCount)
}

var _cMTagCollectionAddTagsFromCollection func(tagCollection CMMutableTagCollectionRef, collectionWithTagsToAdd CMTagCollectionRef) int32

// CMTagCollectionAddTagsFromCollection add the tags contained in one collection to another.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionAddTagsFromCollection
func CMTagCollectionAddTagsFromCollection(tagCollection CMMutableTagCollectionRef, collectionWithTagsToAdd CMTagCollectionRef) int32 {
	if _cMTagCollectionAddTagsFromCollection == nil {
		panic("CoreMedia: symbol CMTagCollectionAddTagsFromCollection not loaded")
	}
	return _cMTagCollectionAddTagsFromCollection(tagCollection, collectionWithTagsToAdd)
}

var _cMTagCollectionApply func(tagCollection CMTagCollectionRef, applier CMTagCollectionApplierFunction, context unsafe.Pointer)

// CMTagCollectionApply applies a function to all tags in a collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionApply
func CMTagCollectionApply(tagCollection CMTagCollectionRef, applier CMTagCollectionApplierFunction, context unsafe.Pointer) {
	if _cMTagCollectionApply == nil {
		panic("CoreMedia: symbol CMTagCollectionApply not loaded")
	}
	_cMTagCollectionApply(tagCollection, applier, context)
}

var _cMTagCollectionApplyUntil func(tagCollection CMTagCollectionRef, applier CMTagCollectionTagFilterFunction, context unsafe.Pointer) CMTag

// CMTagCollectionApplyUntil applies a Boolean function to tags in a collection, stopping when it returns true.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionApplyUntil
func CMTagCollectionApplyUntil(tagCollection CMTagCollectionRef, applier CMTagCollectionTagFilterFunction, context unsafe.Pointer) CMTag {
	if _cMTagCollectionApplyUntil == nil {
		panic("CoreMedia: symbol CMTagCollectionApplyUntil not loaded")
	}
	return _cMTagCollectionApplyUntil(tagCollection, applier, context)
}

var _cMTagCollectionContainsCategory func(tagCollection CMTagCollectionRef, category CMTagCategory) bool

// CMTagCollectionContainsCategory determines if a tag collection contains tags for a given category.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionContainsCategory
func CMTagCollectionContainsCategory(tagCollection CMTagCollectionRef, category CMTagCategory) bool {
	if _cMTagCollectionContainsCategory == nil {
		panic("CoreMedia: symbol CMTagCollectionContainsCategory not loaded")
	}
	return _cMTagCollectionContainsCategory(tagCollection, category)
}

var _cMTagCollectionContainsSpecifiedTags func(tagCollection CMTagCollectionRef, containedTags *CMTag, containedTagCount int) bool

// CMTagCollectionContainsSpecifiedTags determines if a tag collection contains a subset of tags.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionContainsSpecifiedTags
func CMTagCollectionContainsSpecifiedTags(tagCollection CMTagCollectionRef, containedTags *CMTag, containedTagCount int) bool {
	if _cMTagCollectionContainsSpecifiedTags == nil {
		panic("CoreMedia: symbol CMTagCollectionContainsSpecifiedTags not loaded")
	}
	return _cMTagCollectionContainsSpecifiedTags(tagCollection, containedTags, containedTagCount)
}

var _cMTagCollectionContainsTag func(tagCollection CMTagCollectionRef, tag CMTag) bool

// CMTagCollectionContainsTag determines if a tag collection contains a specific tag.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionContainsTag
func CMTagCollectionContainsTag(tagCollection CMTagCollectionRef, tag CMTag) bool {
	if _cMTagCollectionContainsTag == nil {
		panic("CoreMedia: symbol CMTagCollectionContainsTag not loaded")
	}
	return _cMTagCollectionContainsTag(tagCollection, tag)
}

var _cMTagCollectionContainsTagsOfCollection func(tagCollection CMTagCollectionRef, containedTagCollection CMTagCollectionRef) bool

// CMTagCollectionContainsTagsOfCollection determines if one collection of tags contains every tag from another collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionContainsTagsOfCollection
func CMTagCollectionContainsTagsOfCollection(tagCollection CMTagCollectionRef, containedTagCollection CMTagCollectionRef) bool {
	if _cMTagCollectionContainsTagsOfCollection == nil {
		panic("CoreMedia: symbol CMTagCollectionContainsTagsOfCollection not loaded")
	}
	return _cMTagCollectionContainsTagsOfCollection(tagCollection, containedTagCollection)
}

var _cMTagCollectionCopyAsData func(tagCollection CMTagCollectionRef, allocator corefoundation.CFAllocatorRef) corefoundation.CFDataRef

// CMTagCollectionCopyAsData creates a new Core Foundation data instance from a tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCopyAsData
func CMTagCollectionCopyAsData(tagCollection CMTagCollectionRef, allocator corefoundation.CFAllocatorRef) corefoundation.CFDataRef {
	if _cMTagCollectionCopyAsData == nil {
		panic("CoreMedia: symbol CMTagCollectionCopyAsData not loaded")
	}
	return _cMTagCollectionCopyAsData(tagCollection, allocator)
}

var _cMTagCollectionCopyAsDictionary func(tagCollection CMTagCollectionRef, allocator corefoundation.CFAllocatorRef) corefoundation.CFDictionaryRef

// CMTagCollectionCopyAsDictionary creates a new Core Foundation dictionary from a tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCopyAsDictionary
func CMTagCollectionCopyAsDictionary(tagCollection CMTagCollectionRef, allocator corefoundation.CFAllocatorRef) corefoundation.CFDictionaryRef {
	if _cMTagCollectionCopyAsDictionary == nil {
		panic("CoreMedia: symbol CMTagCollectionCopyAsDictionary not loaded")
	}
	return _cMTagCollectionCopyAsDictionary(tagCollection, allocator)
}

var _cMTagCollectionCopyDescription func(allocator corefoundation.CFAllocatorRef, tagCollection CMTagCollectionRef) corefoundation.CFStringRef

// CMTagCollectionCopyDescription retrieves a copy of the tag collection’s description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCopyDescription
func CMTagCollectionCopyDescription(allocator corefoundation.CFAllocatorRef, tagCollection CMTagCollectionRef) corefoundation.CFStringRef {
	if _cMTagCollectionCopyDescription == nil {
		panic("CoreMedia: symbol CMTagCollectionCopyDescription not loaded")
	}
	return _cMTagCollectionCopyDescription(allocator, tagCollection)
}

var _cMTagCollectionCopyTagsOfCategories func(allocator corefoundation.CFAllocatorRef, tagCollection CMTagCollectionRef, categories *CMTagCategory, categoriesCount int, collectionWithTagsOfCategories *CMTagCollectionRef) int32

// CMTagCollectionCopyTagsOfCategories creates a new tag collection from an existing collection, copying all tags which match a list of categories.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCopyTagsOfCategories
func CMTagCollectionCopyTagsOfCategories(allocator corefoundation.CFAllocatorRef, tagCollection CMTagCollectionRef, categories *CMTagCategory, categoriesCount int, collectionWithTagsOfCategories *CMTagCollectionRef) int32 {
	if _cMTagCollectionCopyTagsOfCategories == nil {
		panic("CoreMedia: symbol CMTagCollectionCopyTagsOfCategories not loaded")
	}
	return _cMTagCollectionCopyTagsOfCategories(allocator, tagCollection, categories, categoriesCount, collectionWithTagsOfCategories)
}

var _cMTagCollectionCountTagsWithFilterFunction func(tagCollection CMTagCollectionRef, filterApplier CMTagCollectionTagFilterFunction, context unsafe.Pointer) int

// CMTagCollectionCountTagsWithFilterFunction counts the number of tags in a collection matching an evaluation function.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCountTagsWithFilterFunction
func CMTagCollectionCountTagsWithFilterFunction(tagCollection CMTagCollectionRef, filterApplier CMTagCollectionTagFilterFunction, context unsafe.Pointer) int {
	if _cMTagCollectionCountTagsWithFilterFunction == nil {
		panic("CoreMedia: symbol CMTagCollectionCountTagsWithFilterFunction not loaded")
	}
	return _cMTagCollectionCountTagsWithFilterFunction(tagCollection, filterApplier, context)
}

var _cMTagCollectionCreate func(allocator corefoundation.CFAllocatorRef, tags *CMTag, tagCount int, newCollectionOut *CMTagCollectionRef) int32

// CMTagCollectionCreate creates a new tag collection from an existing C-style array of tags.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreate
func CMTagCollectionCreate(allocator corefoundation.CFAllocatorRef, tags *CMTag, tagCount int, newCollectionOut *CMTagCollectionRef) int32 {
	if _cMTagCollectionCreate == nil {
		panic("CoreMedia: symbol CMTagCollectionCreate not loaded")
	}
	return _cMTagCollectionCreate(allocator, tags, tagCount, newCollectionOut)
}

var _cMTagCollectionCreateCopy func(tagCollection CMTagCollectionRef, allocator corefoundation.CFAllocatorRef, newCollectionCopyOut *CMTagCollectionRef) int32

// CMTagCollectionCreateCopy creates a copy of a tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateCopy
func CMTagCollectionCreateCopy(tagCollection CMTagCollectionRef, allocator corefoundation.CFAllocatorRef, newCollectionCopyOut *CMTagCollectionRef) int32 {
	if _cMTagCollectionCreateCopy == nil {
		panic("CoreMedia: symbol CMTagCollectionCreateCopy not loaded")
	}
	return _cMTagCollectionCreateCopy(tagCollection, allocator, newCollectionCopyOut)
}

var _cMTagCollectionCreateDifference func(tagCollectionMinuend CMTagCollectionRef, tagCollectionSubtrahend CMTagCollectionRef, tagCollectionOut *CMTagCollectionRef) int32

// CMTagCollectionCreateDifference creates a new tag collection with the difference of two existing collections.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateDifference
func CMTagCollectionCreateDifference(tagCollectionMinuend CMTagCollectionRef, tagCollectionSubtrahend CMTagCollectionRef, tagCollectionOut *CMTagCollectionRef) int32 {
	if _cMTagCollectionCreateDifference == nil {
		panic("CoreMedia: symbol CMTagCollectionCreateDifference not loaded")
	}
	return _cMTagCollectionCreateDifference(tagCollectionMinuend, tagCollectionSubtrahend, tagCollectionOut)
}

var _cMTagCollectionCreateExclusiveOr func(tagCollection1 CMTagCollectionRef, tagCollection2 CMTagCollectionRef, tagCollectionOut *CMTagCollectionRef) int32

// CMTagCollectionCreateExclusiveOr creates a new tag collection from two existing tag collections, copying elements which are in one collection or the other, but not both.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateExclusiveOr
func CMTagCollectionCreateExclusiveOr(tagCollection1 CMTagCollectionRef, tagCollection2 CMTagCollectionRef, tagCollectionOut *CMTagCollectionRef) int32 {
	if _cMTagCollectionCreateExclusiveOr == nil {
		panic("CoreMedia: symbol CMTagCollectionCreateExclusiveOr not loaded")
	}
	return _cMTagCollectionCreateExclusiveOr(tagCollection1, tagCollection2, tagCollectionOut)
}

var _cMTagCollectionCreateFromData func(data corefoundation.CFDataRef, allocator corefoundation.CFAllocatorRef, newCollectionOut *CMTagCollectionRef) int32

// CMTagCollectionCreateFromData creates a new tag collection from an existing Core Foundation data instance.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateFromData
func CMTagCollectionCreateFromData(data corefoundation.CFDataRef, allocator corefoundation.CFAllocatorRef, newCollectionOut *CMTagCollectionRef) int32 {
	if _cMTagCollectionCreateFromData == nil {
		panic("CoreMedia: symbol CMTagCollectionCreateFromData not loaded")
	}
	return _cMTagCollectionCreateFromData(data, allocator, newCollectionOut)
}

var _cMTagCollectionCreateFromDictionary func(dict corefoundation.CFDictionaryRef, allocator corefoundation.CFAllocatorRef, newCollectionOut *CMTagCollectionRef) int32

// CMTagCollectionCreateFromDictionary creates a new tag collection from an existing Core Foundation dictionary.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateFromDictionary
func CMTagCollectionCreateFromDictionary(dict corefoundation.CFDictionaryRef, allocator corefoundation.CFAllocatorRef, newCollectionOut *CMTagCollectionRef) int32 {
	if _cMTagCollectionCreateFromDictionary == nil {
		panic("CoreMedia: symbol CMTagCollectionCreateFromDictionary not loaded")
	}
	return _cMTagCollectionCreateFromDictionary(dict, allocator, newCollectionOut)
}

var _cMTagCollectionCreateIntersection func(tagCollection1 CMTagCollectionRef, tagCollection2 CMTagCollectionRef, tagCollectionOut *CMTagCollectionRef) int32

// CMTagCollectionCreateIntersection creates a new tag collection containing only the tags from two existing collections which match.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateIntersection
func CMTagCollectionCreateIntersection(tagCollection1 CMTagCollectionRef, tagCollection2 CMTagCollectionRef, tagCollectionOut *CMTagCollectionRef) int32 {
	if _cMTagCollectionCreateIntersection == nil {
		panic("CoreMedia: symbol CMTagCollectionCreateIntersection not loaded")
	}
	return _cMTagCollectionCreateIntersection(tagCollection1, tagCollection2, tagCollectionOut)
}

var _cMTagCollectionCreateMutable func(allocator corefoundation.CFAllocatorRef, capacity int, newMutableCollectionOut *CMMutableTagCollectionRef) int32

// CMTagCollectionCreateMutable creates a new, mutable, empty tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateMutable
func CMTagCollectionCreateMutable(allocator corefoundation.CFAllocatorRef, capacity int, newMutableCollectionOut *CMMutableTagCollectionRef) int32 {
	if _cMTagCollectionCreateMutable == nil {
		panic("CoreMedia: symbol CMTagCollectionCreateMutable not loaded")
	}
	return _cMTagCollectionCreateMutable(allocator, capacity, newMutableCollectionOut)
}

var _cMTagCollectionCreateMutableCopy func(tagCollection CMTagCollectionRef, allocator corefoundation.CFAllocatorRef, newMutableCollectionCopyOut *CMMutableTagCollectionRef) int32

// CMTagCollectionCreateMutableCopy creates a new mutable copy from an existing tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateMutableCopy
func CMTagCollectionCreateMutableCopy(tagCollection CMTagCollectionRef, allocator corefoundation.CFAllocatorRef, newMutableCollectionCopyOut *CMMutableTagCollectionRef) int32 {
	if _cMTagCollectionCreateMutableCopy == nil {
		panic("CoreMedia: symbol CMTagCollectionCreateMutableCopy not loaded")
	}
	return _cMTagCollectionCreateMutableCopy(tagCollection, allocator, newMutableCollectionCopyOut)
}

var _cMTagCollectionCreateUnion func(tagCollection1 CMTagCollectionRef, tagCollection2 CMTagCollectionRef, tagCollectionOut *CMTagCollectionRef) int32

// CMTagCollectionCreateUnion creates a new tag collection containing all tags from two collections without duplicates.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionCreateUnion
func CMTagCollectionCreateUnion(tagCollection1 CMTagCollectionRef, tagCollection2 CMTagCollectionRef, tagCollectionOut *CMTagCollectionRef) int32 {
	if _cMTagCollectionCreateUnion == nil {
		panic("CoreMedia: symbol CMTagCollectionCreateUnion not loaded")
	}
	return _cMTagCollectionCreateUnion(tagCollection1, tagCollection2, tagCollectionOut)
}

var _cMTagCollectionGetCount func(tagCollection CMTagCollectionRef) int

// CMTagCollectionGetCount gets the number of tags in a tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionGetCount
func CMTagCollectionGetCount(tagCollection CMTagCollectionRef) int {
	if _cMTagCollectionGetCount == nil {
		panic("CoreMedia: symbol CMTagCollectionGetCount not loaded")
	}
	return _cMTagCollectionGetCount(tagCollection)
}

var _cMTagCollectionGetCountOfCategory func(tagCollection CMTagCollectionRef, category CMTagCategory) int

// CMTagCollectionGetCountOfCategory retrieves the number of tags in the collection matching a given category.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionGetCountOfCategory
func CMTagCollectionGetCountOfCategory(tagCollection CMTagCollectionRef, category CMTagCategory) int {
	if _cMTagCollectionGetCountOfCategory == nil {
		panic("CoreMedia: symbol CMTagCollectionGetCountOfCategory not loaded")
	}
	return _cMTagCollectionGetCountOfCategory(tagCollection, category)
}

var _cMTagCollectionGetTags func(tagCollection CMTagCollectionRef, tagBuffer *CMTag, tagBufferCount int, numberOfTagsCopied *int) int32

// CMTagCollectionGetTags retrieves an arbitrary number of tags from the collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionGetTags
func CMTagCollectionGetTags(tagCollection CMTagCollectionRef, tagBuffer *CMTag, tagBufferCount int, numberOfTagsCopied *int) int32 {
	if _cMTagCollectionGetTags == nil {
		panic("CoreMedia: symbol CMTagCollectionGetTags not loaded")
	}
	return _cMTagCollectionGetTags(tagCollection, tagBuffer, tagBufferCount, numberOfTagsCopied)
}

var _cMTagCollectionGetTagsWithCategory func(tagCollection CMTagCollectionRef, category CMTagCategory, tagBuffer *CMTag, tagBufferCount int, numberOfTagsCopied *int) int32

// CMTagCollectionGetTagsWithCategory retrieves a C-style array of tags with a given category from a tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionGetTagsWithCategory
func CMTagCollectionGetTagsWithCategory(tagCollection CMTagCollectionRef, category CMTagCategory, tagBuffer *CMTag, tagBufferCount int, numberOfTagsCopied *int) int32 {
	if _cMTagCollectionGetTagsWithCategory == nil {
		panic("CoreMedia: symbol CMTagCollectionGetTagsWithCategory not loaded")
	}
	return _cMTagCollectionGetTagsWithCategory(tagCollection, category, tagBuffer, tagBufferCount, numberOfTagsCopied)
}

var _cMTagCollectionGetTagsWithFilterFunction func(tagCollection CMTagCollectionRef, tagBuffer *CMTag, tagBufferCount int, numberOfTagsCopied *int, filter CMTagCollectionTagFilterFunction, context unsafe.Pointer) int32

// CMTagCollectionGetTagsWithFilterFunction gets all tags in a collection matching an evaluation function.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionGetTagsWithFilterFunction
func CMTagCollectionGetTagsWithFilterFunction(tagCollection CMTagCollectionRef, tagBuffer *CMTag, tagBufferCount int, numberOfTagsCopied *int, filter CMTagCollectionTagFilterFunction, context unsafe.Pointer) int32 {
	if _cMTagCollectionGetTagsWithFilterFunction == nil {
		panic("CoreMedia: symbol CMTagCollectionGetTagsWithFilterFunction not loaded")
	}
	return _cMTagCollectionGetTagsWithFilterFunction(tagCollection, tagBuffer, tagBufferCount, numberOfTagsCopied, filter, context)
}

var _cMTagCollectionGetTypeID func() uint

// CMTagCollectionGetTypeID retrieves the internal type ID for tag collections.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionGetTypeID
func CMTagCollectionGetTypeID() uint {
	if _cMTagCollectionGetTypeID == nil {
		panic("CoreMedia: symbol CMTagCollectionGetTypeID not loaded")
	}
	return _cMTagCollectionGetTypeID()
}

var _cMTagCollectionIsEmpty func(tagCollection CMTagCollectionRef) bool

// CMTagCollectionIsEmpty checks if a tag collection has no elements.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionIsEmpty
func CMTagCollectionIsEmpty(tagCollection CMTagCollectionRef) bool {
	if _cMTagCollectionIsEmpty == nil {
		panic("CoreMedia: symbol CMTagCollectionIsEmpty not loaded")
	}
	return _cMTagCollectionIsEmpty(tagCollection)
}

var _cMTagCollectionRemoveAllTags func(tagCollection CMMutableTagCollectionRef) int32

// CMTagCollectionRemoveAllTags removes all tags from a collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionRemoveAllTags
func CMTagCollectionRemoveAllTags(tagCollection CMMutableTagCollectionRef) int32 {
	if _cMTagCollectionRemoveAllTags == nil {
		panic("CoreMedia: symbol CMTagCollectionRemoveAllTags not loaded")
	}
	return _cMTagCollectionRemoveAllTags(tagCollection)
}

var _cMTagCollectionRemoveAllTagsOfCategory func(tagCollection CMMutableTagCollectionRef, category CMTagCategory) int32

// CMTagCollectionRemoveAllTagsOfCategory removes all tags of a given category from a collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionRemoveAllTagsOfCategory
func CMTagCollectionRemoveAllTagsOfCategory(tagCollection CMMutableTagCollectionRef, category CMTagCategory) int32 {
	if _cMTagCollectionRemoveAllTagsOfCategory == nil {
		panic("CoreMedia: symbol CMTagCollectionRemoveAllTagsOfCategory not loaded")
	}
	return _cMTagCollectionRemoveAllTagsOfCategory(tagCollection, category)
}

var _cMTagCollectionRemoveTag func(tagCollection CMMutableTagCollectionRef, tagToRemove CMTag) int32

// CMTagCollectionRemoveTag removes a specific tag from a collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCollectionRemoveTag
func CMTagCollectionRemoveTag(tagCollection CMMutableTagCollectionRef, tagToRemove CMTag) int32 {
	if _cMTagCollectionRemoveTag == nil {
		panic("CoreMedia: symbol CMTagCollectionRemoveTag not loaded")
	}
	return _cMTagCollectionRemoveTag(tagCollection, tagToRemove)
}

var _cMTagCompare func(tag1 CMTag, tag2 CMTag) corefoundation.CFComparisonResult

// CMTagCompare compares two tags in terms of partial equality.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCompare
func CMTagCompare(tag1 CMTag, tag2 CMTag) corefoundation.CFComparisonResult {
	if _cMTagCompare == nil {
		panic("CoreMedia: symbol CMTagCompare not loaded")
	}
	return _cMTagCompare(tag1, tag2)
}

var _cMTagCopyAsDictionary func(tag CMTag, allocator corefoundation.CFAllocatorRef) corefoundation.CFDictionaryRef

// CMTagCopyAsDictionary copies an existing tag to a new dictionary object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCopyAsDictionary
func CMTagCopyAsDictionary(tag CMTag, allocator corefoundation.CFAllocatorRef) corefoundation.CFDictionaryRef {
	if _cMTagCopyAsDictionary == nil {
		panic("CoreMedia: symbol CMTagCopyAsDictionary not loaded")
	}
	return _cMTagCopyAsDictionary(tag, allocator)
}

var _cMTagCopyDescription func(allocator corefoundation.CFAllocatorRef, tag CMTag) corefoundation.CFStringRef

// CMTagCopyDescription copies the description of a tag to a new string.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagCopyDescription
func CMTagCopyDescription(allocator corefoundation.CFAllocatorRef, tag CMTag) corefoundation.CFStringRef {
	if _cMTagCopyDescription == nil {
		panic("CoreMedia: symbol CMTagCopyDescription not loaded")
	}
	return _cMTagCopyDescription(allocator, tag)
}

var _cMTagEqualToTag func(tag1 CMTag, tag2 CMTag) bool

// CMTagEqualToTag compares two tags for strict equality.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagEqualToTag
func CMTagEqualToTag(tag1 CMTag, tag2 CMTag) bool {
	if _cMTagEqualToTag == nil {
		panic("CoreMedia: symbol CMTagEqualToTag not loaded")
	}
	return _cMTagEqualToTag(tag1, tag2)
}

var _cMTagGetFlagsValue func(tag CMTag) uint64

// CMTagGetFlagsValue retrieves a tag’s value as a 64-bit field flag.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagGetFlagsValue
func CMTagGetFlagsValue(tag CMTag) uint64 {
	if _cMTagGetFlagsValue == nil {
		panic("CoreMedia: symbol CMTagGetFlagsValue not loaded")
	}
	return _cMTagGetFlagsValue(tag)
}

var _cMTagGetFloat64Value func(tag CMTag) float64

// CMTagGetFloat64Value retrieves a tag’s value as a 64-bit floating point number.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagGetFloat64Value
func CMTagGetFloat64Value(tag CMTag) float64 {
	if _cMTagGetFloat64Value == nil {
		panic("CoreMedia: symbol CMTagGetFloat64Value not loaded")
	}
	return _cMTagGetFloat64Value(tag)
}

var _cMTagGetOSTypeValue func(tag CMTag) uint32

// CMTagGetOSTypeValue retrieves a tag’s value for use by the operating system.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagGetOSTypeValue
func CMTagGetOSTypeValue(tag CMTag) uint32 {
	if _cMTagGetOSTypeValue == nil {
		panic("CoreMedia: symbol CMTagGetOSTypeValue not loaded")
	}
	return _cMTagGetOSTypeValue(tag)
}

var _cMTagGetSInt64Value func(tag CMTag) int64

// CMTagGetSInt64Value retrieves a tag’s value as a signed 64-bit integer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagGetSInt64Value
func CMTagGetSInt64Value(tag CMTag) int64 {
	if _cMTagGetSInt64Value == nil {
		panic("CoreMedia: symbol CMTagGetSInt64Value not loaded")
	}
	return _cMTagGetSInt64Value(tag)
}

var _cMTagGetValueDataType func(tag CMTag) CMTagDataType

// CMTagGetValueDataType retrieves the data type of a tag.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagGetValueDataType
func CMTagGetValueDataType(tag CMTag) CMTagDataType {
	if _cMTagGetValueDataType == nil {
		panic("CoreMedia: symbol CMTagGetValueDataType not loaded")
	}
	return _cMTagGetValueDataType(tag)
}

var _cMTagHasFlagsValue func(tag CMTag) bool

// CMTagHasFlagsValue whether a given tag contains a value for a 64-bit flag field.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagHasFlagsValue
func CMTagHasFlagsValue(tag CMTag) bool {
	if _cMTagHasFlagsValue == nil {
		panic("CoreMedia: symbol CMTagHasFlagsValue not loaded")
	}
	return _cMTagHasFlagsValue(tag)
}

var _cMTagHasFloat64Value func(tag CMTag) bool

// CMTagHasFloat64Value whether a given tag contains a value for a 64-bit floating point number.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagHasFloat64Value
func CMTagHasFloat64Value(tag CMTag) bool {
	if _cMTagHasFloat64Value == nil {
		panic("CoreMedia: symbol CMTagHasFloat64Value not loaded")
	}
	return _cMTagHasFloat64Value(tag)
}

var _cMTagHasOSTypeValue func(tag CMTag) bool

// CMTagHasOSTypeValue whether a given tag contains a value for use by the operating system.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagHasOSTypeValue
func CMTagHasOSTypeValue(tag CMTag) bool {
	if _cMTagHasOSTypeValue == nil {
		panic("CoreMedia: symbol CMTagHasOSTypeValue not loaded")
	}
	return _cMTagHasOSTypeValue(tag)
}

var _cMTagHasSInt64Value func(tag CMTag) bool

// CMTagHasSInt64Value whether a given tag contains a value for a signed 64-bit integer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagHasSInt64Value
func CMTagHasSInt64Value(tag CMTag) bool {
	if _cMTagHasSInt64Value == nil {
		panic("CoreMedia: symbol CMTagHasSInt64Value not loaded")
	}
	return _cMTagHasSInt64Value(tag)
}

var _cMTagHash func(tag CMTag) uint

// CMTagHash generates a hash identifier for a tag.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagHash
func CMTagHash(tag CMTag) uint {
	if _cMTagHash == nil {
		panic("CoreMedia: symbol CMTagHash not loaded")
	}
	return _cMTagHash(tag)
}

var _cMTagMakeFromDictionary func(dict corefoundation.CFDictionaryRef) CMTag

// CMTagMakeFromDictionary create a new tag from a dictionary object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagMakeFromDictionary
func CMTagMakeFromDictionary(dict corefoundation.CFDictionaryRef) CMTag {
	if _cMTagMakeFromDictionary == nil {
		panic("CoreMedia: symbol CMTagMakeFromDictionary not loaded")
	}
	return _cMTagMakeFromDictionary(dict)
}

var _cMTagMakeWithFlagsValue func(category CMTagCategory, flagsForTag uint64) CMTag

// CMTagMakeWithFlagsValue creates a new tag with a given category and a value interpreted as a 64-bit flag field.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagMakeWithFlagsValue
func CMTagMakeWithFlagsValue(category CMTagCategory, flagsForTag uint64) CMTag {
	if _cMTagMakeWithFlagsValue == nil {
		panic("CoreMedia: symbol CMTagMakeWithFlagsValue not loaded")
	}
	return _cMTagMakeWithFlagsValue(category, flagsForTag)
}

var _cMTagMakeWithFloat64Value func(category CMTagCategory, value float64) CMTag

// CMTagMakeWithFloat64Value creates a new tag with a given category and a 64-bit floating point value.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagMakeWithFloat64Value
func CMTagMakeWithFloat64Value(category CMTagCategory, value float64) CMTag {
	if _cMTagMakeWithFloat64Value == nil {
		panic("CoreMedia: symbol CMTagMakeWithFloat64Value not loaded")
	}
	return _cMTagMakeWithFloat64Value(category, value)
}

var _cMTagMakeWithOSTypeValue func(category CMTagCategory, value uint32) CMTag

// CMTagMakeWithOSTypeValue creates a new tag with a given category and a 64-bit value for use by the framework.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagMakeWithOSTypeValue
func CMTagMakeWithOSTypeValue(category CMTagCategory, value uint32) CMTag {
	if _cMTagMakeWithOSTypeValue == nil {
		panic("CoreMedia: symbol CMTagMakeWithOSTypeValue not loaded")
	}
	return _cMTagMakeWithOSTypeValue(category, value)
}

var _cMTagMakeWithSInt64Value func(category CMTagCategory, value int64) CMTag

// CMTagMakeWithSInt64Value creates a new tag with a given category and a 64-bit signed integer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTagMakeWithSInt64Value
func CMTagMakeWithSInt64Value(category CMTagCategory, value int64) CMTag {
	if _cMTagMakeWithSInt64Value == nil {
		panic("CoreMedia: symbol CMTagMakeWithSInt64Value not loaded")
	}
	return _cMTagMakeWithSInt64Value(category, value)
}

var _cMTaggedBufferGroupCreate func(allocator corefoundation.CFAllocatorRef, tagCollections corefoundation.CFArrayRef, buffers corefoundation.CFArrayRef, groupOut *CMTaggedBufferGroupRef) int32

// CMTaggedBufferGroupCreate creates a new tagged buffer group from a pair of buffers and the tags to associate with them.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupCreate
func CMTaggedBufferGroupCreate(allocator corefoundation.CFAllocatorRef, tagCollections corefoundation.CFArrayRef, buffers corefoundation.CFArrayRef, groupOut *CMTaggedBufferGroupRef) int32 {
	if _cMTaggedBufferGroupCreate == nil {
		panic("CoreMedia: symbol CMTaggedBufferGroupCreate not loaded")
	}
	return _cMTaggedBufferGroupCreate(allocator, tagCollections, buffers, groupOut)
}

var _cMTaggedBufferGroupCreateCombined func(allocator corefoundation.CFAllocatorRef, taggedBufferGroups corefoundation.CFArrayRef, groupOut *CMTaggedBufferGroupRef) int32

// CMTaggedBufferGroupCreateCombined creates a new tagged buffer group from an array of existing tagged buffer groups.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupCreateCombined
func CMTaggedBufferGroupCreateCombined(allocator corefoundation.CFAllocatorRef, taggedBufferGroups corefoundation.CFArrayRef, groupOut *CMTaggedBufferGroupRef) int32 {
	if _cMTaggedBufferGroupCreateCombined == nil {
		panic("CoreMedia: symbol CMTaggedBufferGroupCreateCombined not loaded")
	}
	return _cMTaggedBufferGroupCreateCombined(allocator, taggedBufferGroups, groupOut)
}

var _cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup func(allocator corefoundation.CFAllocatorRef, taggedBufferGroup CMTaggedBufferGroupRef, formatDescriptionOut *CMTaggedBufferGroupFormatDescriptionRef) int32

// CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup creates a new format description for a tagged buffer group.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup
func CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup(allocator corefoundation.CFAllocatorRef, taggedBufferGroup CMTaggedBufferGroupRef, formatDescriptionOut *CMTaggedBufferGroupFormatDescriptionRef) int32 {
	if _cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup == nil {
		panic("CoreMedia: symbol CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup not loaded")
	}
	return _cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup(allocator, taggedBufferGroup, formatDescriptionOut)
}

var _cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions func(allocator corefoundation.CFAllocatorRef, taggedBufferGroup CMTaggedBufferGroupRef, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *CMTaggedBufferGroupFormatDescriptionRef) int32

// CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions
func CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions(allocator corefoundation.CFAllocatorRef, taggedBufferGroup CMTaggedBufferGroupRef, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *CMTaggedBufferGroupFormatDescriptionRef) int32 {
	if _cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions == nil {
		panic("CoreMedia: symbol CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions not loaded")
	}
	return _cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions(allocator, taggedBufferGroup, extensions, formatDescriptionOut)
}

var _cMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup func(desc CMTaggedBufferGroupFormatDescriptionRef, taggedBufferGroup CMTaggedBufferGroupRef) bool

// CMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup checks to see if a tagged buffer group’s format matches an existing format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup
func CMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup(desc CMTaggedBufferGroupFormatDescriptionRef, taggedBufferGroup CMTaggedBufferGroupRef) bool {
	if _cMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup == nil {
		panic("CoreMedia: symbol CMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup not loaded")
	}
	return _cMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup(desc, taggedBufferGroup)
}

var _cMTaggedBufferGroupGetCMSampleBufferAtIndex func(group CMTaggedBufferGroupRef, index int) uintptr

// CMTaggedBufferGroupGetCMSampleBufferAtIndex gets the sample buffer at a given index in the buffer group.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCMSampleBufferAtIndex
func CMTaggedBufferGroupGetCMSampleBufferAtIndex(group CMTaggedBufferGroupRef, index int) uintptr {
	if _cMTaggedBufferGroupGetCMSampleBufferAtIndex == nil {
		panic("CoreMedia: symbol CMTaggedBufferGroupGetCMSampleBufferAtIndex not loaded")
	}
	return _cMTaggedBufferGroupGetCMSampleBufferAtIndex(group, index)
}

var _cMTaggedBufferGroupGetCMSampleBufferForTag func(group CMTaggedBufferGroupRef, tag CMTag, indexOut *int) uintptr

// CMTaggedBufferGroupGetCMSampleBufferForTag gets the single sample buffer in a group which contains a given tag, if present.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCMSampleBufferForTag
func CMTaggedBufferGroupGetCMSampleBufferForTag(group CMTaggedBufferGroupRef, tag CMTag, indexOut *int) uintptr {
	if _cMTaggedBufferGroupGetCMSampleBufferForTag == nil {
		panic("CoreMedia: symbol CMTaggedBufferGroupGetCMSampleBufferForTag not loaded")
	}
	return _cMTaggedBufferGroupGetCMSampleBufferForTag(group, tag, indexOut)
}

var _cMTaggedBufferGroupGetCMSampleBufferForTagCollection func(group CMTaggedBufferGroupRef, tagCollection CMTagCollectionRef, indexOut *int) uintptr

// CMTaggedBufferGroupGetCMSampleBufferForTagCollection gets the single sample buffer in a group which contains a given tag collection, if present.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCMSampleBufferForTagCollection
func CMTaggedBufferGroupGetCMSampleBufferForTagCollection(group CMTaggedBufferGroupRef, tagCollection CMTagCollectionRef, indexOut *int) uintptr {
	if _cMTaggedBufferGroupGetCMSampleBufferForTagCollection == nil {
		panic("CoreMedia: symbol CMTaggedBufferGroupGetCMSampleBufferForTagCollection not loaded")
	}
	return _cMTaggedBufferGroupGetCMSampleBufferForTagCollection(group, tagCollection, indexOut)
}

var _cMTaggedBufferGroupGetCVPixelBufferAtIndex func(group CMTaggedBufferGroupRef, index int) corevideo.CVPixelBufferRef

// CMTaggedBufferGroupGetCVPixelBufferAtIndex gets the pixel buffer at a given index in the buffer group.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCVPixelBufferAtIndex
func CMTaggedBufferGroupGetCVPixelBufferAtIndex(group CMTaggedBufferGroupRef, index int) corevideo.CVPixelBufferRef {
	if _cMTaggedBufferGroupGetCVPixelBufferAtIndex == nil {
		panic("CoreMedia: symbol CMTaggedBufferGroupGetCVPixelBufferAtIndex not loaded")
	}
	return _cMTaggedBufferGroupGetCVPixelBufferAtIndex(group, index)
}

var _cMTaggedBufferGroupGetCVPixelBufferForTag func(group CMTaggedBufferGroupRef, tag CMTag, indexOut *int) corevideo.CVPixelBufferRef

// CMTaggedBufferGroupGetCVPixelBufferForTag gets the single pixel buffer in a group which contains a given tag, if present.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCVPixelBufferForTag
func CMTaggedBufferGroupGetCVPixelBufferForTag(group CMTaggedBufferGroupRef, tag CMTag, indexOut *int) corevideo.CVPixelBufferRef {
	if _cMTaggedBufferGroupGetCVPixelBufferForTag == nil {
		panic("CoreMedia: symbol CMTaggedBufferGroupGetCVPixelBufferForTag not loaded")
	}
	return _cMTaggedBufferGroupGetCVPixelBufferForTag(group, tag, indexOut)
}

var _cMTaggedBufferGroupGetCVPixelBufferForTagCollection func(group CMTaggedBufferGroupRef, tagCollection CMTagCollectionRef, indexOut *int) corevideo.CVPixelBufferRef

// CMTaggedBufferGroupGetCVPixelBufferForTagCollection gets the single pixel buffer in a group which contains a given tag collection, if present.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCVPixelBufferForTagCollection
func CMTaggedBufferGroupGetCVPixelBufferForTagCollection(group CMTaggedBufferGroupRef, tagCollection CMTagCollectionRef, indexOut *int) corevideo.CVPixelBufferRef {
	if _cMTaggedBufferGroupGetCVPixelBufferForTagCollection == nil {
		panic("CoreMedia: symbol CMTaggedBufferGroupGetCVPixelBufferForTagCollection not loaded")
	}
	return _cMTaggedBufferGroupGetCVPixelBufferForTagCollection(group, tagCollection, indexOut)
}

var _cMTaggedBufferGroupGetCount func(group CMTaggedBufferGroupRef) int

// CMTaggedBufferGroupGetCount gets the number of buffers contained within a tagged buffer group.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetCount
func CMTaggedBufferGroupGetCount(group CMTaggedBufferGroupRef) int {
	if _cMTaggedBufferGroupGetCount == nil {
		panic("CoreMedia: symbol CMTaggedBufferGroupGetCount not loaded")
	}
	return _cMTaggedBufferGroupGetCount(group)
}

var _cMTaggedBufferGroupGetNumberOfMatchesForTagCollection func(group CMTaggedBufferGroupRef, tagCollection CMTagCollectionRef) int

// CMTaggedBufferGroupGetNumberOfMatchesForTagCollection gets the number of buffers in the group associated with a given tag collection.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetNumberOfMatchesForTagCollection
func CMTaggedBufferGroupGetNumberOfMatchesForTagCollection(group CMTaggedBufferGroupRef, tagCollection CMTagCollectionRef) int {
	if _cMTaggedBufferGroupGetNumberOfMatchesForTagCollection == nil {
		panic("CoreMedia: symbol CMTaggedBufferGroupGetNumberOfMatchesForTagCollection not loaded")
	}
	return _cMTaggedBufferGroupGetNumberOfMatchesForTagCollection(group, tagCollection)
}

var _cMTaggedBufferGroupGetTagCollectionAtIndex func(group CMTaggedBufferGroupRef, index int) CMTagCollectionRef

// CMTaggedBufferGroupGetTagCollectionAtIndex gets the collection of tags for a buffer at a given index in the group.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetTagCollectionAtIndex
func CMTaggedBufferGroupGetTagCollectionAtIndex(group CMTaggedBufferGroupRef, index int) CMTagCollectionRef {
	if _cMTaggedBufferGroupGetTagCollectionAtIndex == nil {
		panic("CoreMedia: symbol CMTaggedBufferGroupGetTagCollectionAtIndex not loaded")
	}
	return _cMTaggedBufferGroupGetTagCollectionAtIndex(group, index)
}

var _cMTaggedBufferGroupGetTypeID func() uint

// CMTaggedBufferGroupGetTypeID gets the internal type ID for a tagged buffer group.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTaggedBufferGroupGetTypeID
func CMTaggedBufferGroupGetTypeID() uint {
	if _cMTaggedBufferGroupGetTypeID == nil {
		panic("CoreMedia: symbol CMTaggedBufferGroupGetTypeID not loaded")
	}
	return _cMTaggedBufferGroupGetTypeID()
}

var _cMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, textFormatDescription CMTextFormatDescriptionRef, flavor CMTextDescriptionFlavor, blockBufferOut *uintptr) int32

// CMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer copies the contents of a text format description to a buffer in big-endian byte order.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer(allocator:textFormatDescription:flavor:blockBufferOut:)
func CMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, textFormatDescription CMTextFormatDescriptionRef, flavor CMTextDescriptionFlavor, blockBufferOut *uintptr) int32 {
	if _cMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer == nil {
		panic("CoreMedia: symbol CMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer not loaded")
	}
	return _cMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer(allocator, textFormatDescription, flavor, blockBufferOut)
}

var _cMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, textDescriptionBlockBuffer uintptr, flavor CMTextDescriptionFlavor, mediaType uint32, formatDescriptionOut *CMTextFormatDescriptionRef) int32

// CMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer creates a text format description from a big-endian text description structure inside a buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer(allocator:bigEndianTextDescriptionBlockBuffer:flavor:mediaType:formatDescriptionOut:)
func CMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, textDescriptionBlockBuffer uintptr, flavor CMTextDescriptionFlavor, mediaType uint32, formatDescriptionOut *CMTextFormatDescriptionRef) int32 {
	if _cMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer == nil {
		panic("CoreMedia: symbol CMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer not loaded")
	}
	return _cMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer(allocator, textDescriptionBlockBuffer, flavor, mediaType, formatDescriptionOut)
}

var _cMTextFormatDescriptionCreateFromBigEndianTextDescriptionData func(allocator corefoundation.CFAllocatorRef, textDescriptionData *uint8, size uintptr, flavor CMTextDescriptionFlavor, mediaType uint32, formatDescriptionOut *CMTextFormatDescriptionRef) int32

// CMTextFormatDescriptionCreateFromBigEndianTextDescriptionData creates a text format description from a big-endian text description structure.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionCreateFromBigEndianTextDescriptionData(allocator:bigEndianTextDescriptionData:size:flavor:mediaType:formatDescriptionOut:)
func CMTextFormatDescriptionCreateFromBigEndianTextDescriptionData(allocator corefoundation.CFAllocatorRef, textDescriptionData *uint8, size uintptr, flavor CMTextDescriptionFlavor, mediaType uint32, formatDescriptionOut *CMTextFormatDescriptionRef) int32 {
	if _cMTextFormatDescriptionCreateFromBigEndianTextDescriptionData == nil {
		panic("CoreMedia: symbol CMTextFormatDescriptionCreateFromBigEndianTextDescriptionData not loaded")
	}
	return _cMTextFormatDescriptionCreateFromBigEndianTextDescriptionData(allocator, textDescriptionData, size, flavor, mediaType, formatDescriptionOut)
}

var _cMTextFormatDescriptionGetDefaultStyle func(desc uintptr, localFontIDOut *uint16, boldOut *bool, italicOut *bool, underlineOut *bool, fontSizeOut *float64, colorComponentsOut float64) int32

// CMTextFormatDescriptionGetDefaultStyle returns the default text style.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionGetDefaultStyle(_:localFontIDOut:boldOut:italicOut:underlineOut:fontSizeOut:colorComponentsOut:)
func CMTextFormatDescriptionGetDefaultStyle(desc uintptr, localFontIDOut *uint16, boldOut *bool, italicOut *bool, underlineOut *bool, fontSizeOut *float64, colorComponentsOut float64) int32 {
	if _cMTextFormatDescriptionGetDefaultStyle == nil {
		panic("CoreMedia: symbol CMTextFormatDescriptionGetDefaultStyle not loaded")
	}
	return _cMTextFormatDescriptionGetDefaultStyle(desc, localFontIDOut, boldOut, italicOut, underlineOut, fontSizeOut, colorComponentsOut)
}

var _cMTextFormatDescriptionGetDefaultTextBox func(desc uintptr, originIsAtTopLeft bool, heightOfTextTrack float64, defaultTextBoxOut *corefoundation.CGRect) int32

// CMTextFormatDescriptionGetDefaultTextBox returns the default text box.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionGetDefaultTextBox(_:originIsAtTopLeft:heightOfTextTrack:defaultTextBoxOut:)
func CMTextFormatDescriptionGetDefaultTextBox(desc uintptr, originIsAtTopLeft bool, heightOfTextTrack float64, defaultTextBoxOut *corefoundation.CGRect) int32 {
	if _cMTextFormatDescriptionGetDefaultTextBox == nil {
		panic("CoreMedia: symbol CMTextFormatDescriptionGetDefaultTextBox not loaded")
	}
	return _cMTextFormatDescriptionGetDefaultTextBox(desc, originIsAtTopLeft, heightOfTextTrack, defaultTextBoxOut)
}

var _cMTextFormatDescriptionGetDisplayFlags func(desc uintptr, displayFlagsOut *CMTextDisplayFlags) int32

// CMTextFormatDescriptionGetDisplayFlags returns the display flags.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionGetDisplayFlags(_:displayFlagsOut:)
func CMTextFormatDescriptionGetDisplayFlags(desc uintptr, displayFlagsOut *CMTextDisplayFlags) int32 {
	if _cMTextFormatDescriptionGetDisplayFlags == nil {
		panic("CoreMedia: symbol CMTextFormatDescriptionGetDisplayFlags not loaded")
	}
	return _cMTextFormatDescriptionGetDisplayFlags(desc, displayFlagsOut)
}

var _cMTextFormatDescriptionGetFontName func(desc uintptr, localFontID uint16, fontNameOut *corefoundation.CFStringRef) int32

// CMTextFormatDescriptionGetFontName returns a font name for a local font identifier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionGetFontName(_:localFontID:fontNameOut:)
func CMTextFormatDescriptionGetFontName(desc uintptr, localFontID uint16, fontNameOut *corefoundation.CFStringRef) int32 {
	if _cMTextFormatDescriptionGetFontName == nil {
		panic("CoreMedia: symbol CMTextFormatDescriptionGetFontName not loaded")
	}
	return _cMTextFormatDescriptionGetFontName(desc, localFontID, fontNameOut)
}

var _cMTextFormatDescriptionGetJustification func(desc uintptr, horizontaJustificationlOut *CMTextJustificationValue, verticalJustificationOut *CMTextJustificationValue) int32

// CMTextFormatDescriptionGetJustification returns the horizontal and vertical justification.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTextFormatDescriptionGetJustification(_:horizontalOut:verticalOut:)
func CMTextFormatDescriptionGetJustification(desc uintptr, horizontaJustificationlOut *CMTextJustificationValue, verticalJustificationOut *CMTextJustificationValue) int32 {
	if _cMTextFormatDescriptionGetJustification == nil {
		panic("CoreMedia: symbol CMTextFormatDescriptionGetJustification not loaded")
	}
	return _cMTextFormatDescriptionGetJustification(desc, horizontaJustificationlOut, verticalJustificationOut)
}

var _cMTimeAbsoluteValue func(time CMTime) CMTime

// CMTimeAbsoluteValue returns the absolute value of a time.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeAbsoluteValue(_:)
func CMTimeAbsoluteValue(time CMTime) CMTime {
	if _cMTimeAbsoluteValue == nil {
		panic("CoreMedia: symbol CMTimeAbsoluteValue not loaded")
	}
	return _cMTimeAbsoluteValue(time)
}

var _cMTimeAdd func(lhs CMTime, rhs CMTime) CMTime

// CMTimeAdd returns the sum of two times.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeAdd(_:_:)
func CMTimeAdd(lhs CMTime, rhs CMTime) CMTime {
	if _cMTimeAdd == nil {
		panic("CoreMedia: symbol CMTimeAdd not loaded")
	}
	return _cMTimeAdd(lhs, rhs)
}

var _cMTimeClampToRange func(time CMTime, range_ CMTimeRange) CMTime

// CMTimeClampToRange returns the nearest time value inside the time range.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeClampToRange(_:range:)
func CMTimeClampToRange(time CMTime, range_ CMTimeRange) CMTime {
	if _cMTimeClampToRange == nil {
		panic("CoreMedia: symbol CMTimeClampToRange not loaded")
	}
	return _cMTimeClampToRange(time, range_)
}

var _cMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, timeCodeFormatDescription CMTimeCodeFormatDescriptionRef, flavor CMTimeCodeDescriptionFlavor, blockBufferOut *uintptr) int32

// CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer copies the contents of a time code format description to a buffer in big-endian byte order.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer(allocator:timeCodeFormatDescription:flavor:blockBufferOut:)
func CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, timeCodeFormatDescription CMTimeCodeFormatDescriptionRef, flavor CMTimeCodeDescriptionFlavor, blockBufferOut *uintptr) int32 {
	if _cMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer == nil {
		panic("CoreMedia: symbol CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer not loaded")
	}
	return _cMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer(allocator, timeCodeFormatDescription, flavor, blockBufferOut)
}

var _cMTimeCodeFormatDescriptionCreate func(allocator corefoundation.CFAllocatorRef, timeCodeFormatType CMTimeCodeFormatType, frameDuration CMTime, frameQuanta uint32, flags uint32, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *CMTimeCodeFormatDescriptionRef) int32

// CMTimeCodeFormatDescriptionCreate creates a format description for time code media.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionCreate(allocator:timeCodeFormatType:frameDuration:frameQuanta:flags:extensions:formatDescriptionOut:)
func CMTimeCodeFormatDescriptionCreate(allocator corefoundation.CFAllocatorRef, timeCodeFormatType CMTimeCodeFormatType, frameDuration CMTime, frameQuanta uint32, flags uint32, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *CMTimeCodeFormatDescriptionRef) int32 {
	if _cMTimeCodeFormatDescriptionCreate == nil {
		panic("CoreMedia: symbol CMTimeCodeFormatDescriptionCreate not loaded")
	}
	return _cMTimeCodeFormatDescriptionCreate(allocator, timeCodeFormatType, frameDuration, frameQuanta, flags, extensions, formatDescriptionOut)
}

var _cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, timeCodeDescriptionBlockBuffer uintptr, flavor CMTimeCodeDescriptionFlavor, formatDescriptionOut *CMTimeCodeFormatDescriptionRef) int32

// CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer creates a time code format description from a big-endian time code description data structure in a buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer(allocator:bigEndianTimeCodeDescriptionBlockBuffer:flavor:formatDescriptionOut:)
func CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, timeCodeDescriptionBlockBuffer uintptr, flavor CMTimeCodeDescriptionFlavor, formatDescriptionOut *CMTimeCodeFormatDescriptionRef) int32 {
	if _cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer == nil {
		panic("CoreMedia: symbol CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer not loaded")
	}
	return _cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer(allocator, timeCodeDescriptionBlockBuffer, flavor, formatDescriptionOut)
}

var _cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData func(allocator corefoundation.CFAllocatorRef, timeCodeDescriptionData *uint8, size uintptr, flavor CMTimeCodeDescriptionFlavor, formatDescriptionOut *CMTimeCodeFormatDescriptionRef) int32

// CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData creates a time code format description from a big-endian time code description structure.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData(allocator:bigEndianTimeCodeDescriptionData:size:flavor:formatDescriptionOut:)
func CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData(allocator corefoundation.CFAllocatorRef, timeCodeDescriptionData *uint8, size uintptr, flavor CMTimeCodeDescriptionFlavor, formatDescriptionOut *CMTimeCodeFormatDescriptionRef) int32 {
	if _cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData == nil {
		panic("CoreMedia: symbol CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData not loaded")
	}
	return _cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData(allocator, timeCodeDescriptionData, size, flavor, formatDescriptionOut)
}

var _cMTimeCodeFormatDescriptionGetFrameDuration func(timeCodeFormatDescription CMTimeCodeFormatDescriptionRef) CMTime

// CMTimeCodeFormatDescriptionGetFrameDuration returns the duration of each frame.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionGetFrameDuration(_:)
func CMTimeCodeFormatDescriptionGetFrameDuration(timeCodeFormatDescription CMTimeCodeFormatDescriptionRef) CMTime {
	if _cMTimeCodeFormatDescriptionGetFrameDuration == nil {
		panic("CoreMedia: symbol CMTimeCodeFormatDescriptionGetFrameDuration not loaded")
	}
	return _cMTimeCodeFormatDescriptionGetFrameDuration(timeCodeFormatDescription)
}

var _cMTimeCodeFormatDescriptionGetFrameQuanta func(timeCodeFormatDescription CMTimeCodeFormatDescriptionRef) uint32

// CMTimeCodeFormatDescriptionGetFrameQuanta returns the frames per second for a time code, or frames per tick in counter mode.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionGetFrameQuanta(_:)
func CMTimeCodeFormatDescriptionGetFrameQuanta(timeCodeFormatDescription CMTimeCodeFormatDescriptionRef) uint32 {
	if _cMTimeCodeFormatDescriptionGetFrameQuanta == nil {
		panic("CoreMedia: symbol CMTimeCodeFormatDescriptionGetFrameQuanta not loaded")
	}
	return _cMTimeCodeFormatDescriptionGetFrameQuanta(timeCodeFormatDescription)
}

var _cMTimeCodeFormatDescriptionGetTimeCodeFlags func(desc CMTimeCodeFormatDescriptionRef) uint32

// CMTimeCodeFormatDescriptionGetTimeCodeFlags returns time code flags.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCodeFormatDescriptionGetTimeCodeFlags(_:)
func CMTimeCodeFormatDescriptionGetTimeCodeFlags(desc CMTimeCodeFormatDescriptionRef) uint32 {
	if _cMTimeCodeFormatDescriptionGetTimeCodeFlags == nil {
		panic("CoreMedia: symbol CMTimeCodeFormatDescriptionGetTimeCodeFlags not loaded")
	}
	return _cMTimeCodeFormatDescriptionGetTimeCodeFlags(desc)
}

var _cMTimeCompare func(time1 CMTime, time2 CMTime) int32

// CMTimeCompare returns the numerical relationship of two times.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCompare(_:_:)
func CMTimeCompare(time1 CMTime, time2 CMTime) int32 {
	if _cMTimeCompare == nil {
		panic("CoreMedia: symbol CMTimeCompare not loaded")
	}
	return _cMTimeCompare(time1, time2)
}

var _cMTimeConvertScale func(time CMTime, newTimescale int32, method CMTimeRoundingMethod) CMTime

// CMTimeConvertScale converts the source time to a new timescale using the specified rounding method.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeConvertScale(_:timescale:method:)
func CMTimeConvertScale(time CMTime, newTimescale int32, method CMTimeRoundingMethod) CMTime {
	if _cMTimeConvertScale == nil {
		panic("CoreMedia: symbol CMTimeConvertScale not loaded")
	}
	return _cMTimeConvertScale(time, newTimescale, method)
}

var _cMTimeCopyAsDictionary func(time CMTime, allocator corefoundation.CFAllocatorRef) corefoundation.CFDictionaryRef

// CMTimeCopyAsDictionary creates a dictionary representation of the time.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCopyAsDictionary(_:allocator:)
func CMTimeCopyAsDictionary(time CMTime, allocator corefoundation.CFAllocatorRef) corefoundation.CFDictionaryRef {
	if _cMTimeCopyAsDictionary == nil {
		panic("CoreMedia: symbol CMTimeCopyAsDictionary not loaded")
	}
	return _cMTimeCopyAsDictionary(time, allocator)
}

var _cMTimeCopyDescription func(allocator corefoundation.CFAllocatorRef, time CMTime) corefoundation.CFStringRef

// CMTimeCopyDescription creates a string representation of the time.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeCopyDescription(allocator:time:)
func CMTimeCopyDescription(allocator corefoundation.CFAllocatorRef, time CMTime) corefoundation.CFStringRef {
	if _cMTimeCopyDescription == nil {
		panic("CoreMedia: symbol CMTimeCopyDescription not loaded")
	}
	return _cMTimeCopyDescription(allocator, time)
}

var _cMTimeFoldIntoRange func(time CMTime, foldRange CMTimeRange) CMTime

// CMTimeFoldIntoRange folds a time into a time range.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeFoldIntoRange(_:foldRange:)
func CMTimeFoldIntoRange(time CMTime, foldRange CMTimeRange) CMTime {
	if _cMTimeFoldIntoRange == nil {
		panic("CoreMedia: symbol CMTimeFoldIntoRange not loaded")
	}
	return _cMTimeFoldIntoRange(time, foldRange)
}

var _cMTimeGetSeconds func(time CMTime) float64

// CMTimeGetSeconds returns a representation of the time in seconds.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeGetSeconds(_:)
func CMTimeGetSeconds(time CMTime) float64 {
	if _cMTimeGetSeconds == nil {
		panic("CoreMedia: symbol CMTimeGetSeconds not loaded")
	}
	return _cMTimeGetSeconds(time)
}

var _cMTimeMake func(value int64, timescale int32) CMTime

// CMTimeMake creates a time with a value and timescale.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMake(value:timescale:)
func CMTimeMake(value int64, timescale int32) CMTime {
	if _cMTimeMake == nil {
		panic("CoreMedia: symbol CMTimeMake not loaded")
	}
	return _cMTimeMake(value, timescale)
}

var _cMTimeMakeFromDictionary func(dictionaryRepresentation corefoundation.CFDictionaryRef) CMTime

// CMTimeMakeFromDictionary creates a time from a dictionary representation of its fields.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMakeFromDictionary(_:)
func CMTimeMakeFromDictionary(dictionaryRepresentation corefoundation.CFDictionaryRef) CMTime {
	if _cMTimeMakeFromDictionary == nil {
		panic("CoreMedia: symbol CMTimeMakeFromDictionary not loaded")
	}
	return _cMTimeMakeFromDictionary(dictionaryRepresentation)
}

var _cMTimeMakeWithEpoch func(value int64, timescale int32, epoch int64) CMTime

// CMTimeMakeWithEpoch creates a time with a value, timescale, and epoch.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMakeWithEpoch(value:timescale:epoch:)
func CMTimeMakeWithEpoch(value int64, timescale int32, epoch int64) CMTime {
	if _cMTimeMakeWithEpoch == nil {
		panic("CoreMedia: symbol CMTimeMakeWithEpoch not loaded")
	}
	return _cMTimeMakeWithEpoch(value, timescale, epoch)
}

var _cMTimeMakeWithSeconds func(seconds float64, preferredTimescale int32) CMTime

// CMTimeMakeWithSeconds creates a time that represents a number of seconds in a preferred timescale.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMakeWithSeconds(_:preferredTimescale:)
func CMTimeMakeWithSeconds(seconds float64, preferredTimescale int32) CMTime {
	if _cMTimeMakeWithSeconds == nil {
		panic("CoreMedia: symbol CMTimeMakeWithSeconds not loaded")
	}
	return _cMTimeMakeWithSeconds(seconds, preferredTimescale)
}

var _cMTimeMapDurationFromRangeToRange func(dur CMTime, fromRange CMTimeRange, toRange CMTimeRange) CMTime

// CMTimeMapDurationFromRangeToRange translates a duration through a mapping from two time ranges.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMapDurationFromRangeToRange(_:fromRange:toRange:)
func CMTimeMapDurationFromRangeToRange(dur CMTime, fromRange CMTimeRange, toRange CMTimeRange) CMTime {
	if _cMTimeMapDurationFromRangeToRange == nil {
		panic("CoreMedia: symbol CMTimeMapDurationFromRangeToRange not loaded")
	}
	return _cMTimeMapDurationFromRangeToRange(dur, fromRange, toRange)
}

var _cMTimeMapTimeFromRangeToRange func(t CMTime, fromRange CMTimeRange, toRange CMTimeRange) CMTime

// CMTimeMapTimeFromRangeToRange translates a time through a mapping from two time ranges.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMapTimeFromRangeToRange(_:fromRange:toRange:)
func CMTimeMapTimeFromRangeToRange(t CMTime, fromRange CMTimeRange, toRange CMTimeRange) CMTime {
	if _cMTimeMapTimeFromRangeToRange == nil {
		panic("CoreMedia: symbol CMTimeMapTimeFromRangeToRange not loaded")
	}
	return _cMTimeMapTimeFromRangeToRange(t, fromRange, toRange)
}

var _cMTimeMappingCopyAsDictionary func(mapping CMTimeMapping, allocator corefoundation.CFAllocatorRef) corefoundation.CFDictionaryRef

// CMTimeMappingCopyAsDictionary returns a dictionary representation of a time mapping.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingCopyAsDictionary(_:allocator:)
func CMTimeMappingCopyAsDictionary(mapping CMTimeMapping, allocator corefoundation.CFAllocatorRef) corefoundation.CFDictionaryRef {
	if _cMTimeMappingCopyAsDictionary == nil {
		panic("CoreMedia: symbol CMTimeMappingCopyAsDictionary not loaded")
	}
	return _cMTimeMappingCopyAsDictionary(mapping, allocator)
}

var _cMTimeMappingCopyDescription func(allocator corefoundation.CFAllocatorRef, mapping CMTimeMapping) corefoundation.CFStringRef

// CMTimeMappingCopyDescription copies a string description of a time mapping.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingCopyDescription(allocator:mapping:)
func CMTimeMappingCopyDescription(allocator corefoundation.CFAllocatorRef, mapping CMTimeMapping) corefoundation.CFStringRef {
	if _cMTimeMappingCopyDescription == nil {
		panic("CoreMedia: symbol CMTimeMappingCopyDescription not loaded")
	}
	return _cMTimeMappingCopyDescription(allocator, mapping)
}

var _cMTimeMappingMake func(source CMTimeRange, target CMTimeRange) CMTimeMapping

// CMTimeMappingMake creates a time mapping with a source and target time range.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingMake(source:target:)
func CMTimeMappingMake(source CMTimeRange, target CMTimeRange) CMTimeMapping {
	if _cMTimeMappingMake == nil {
		panic("CoreMedia: symbol CMTimeMappingMake not loaded")
	}
	return _cMTimeMappingMake(source, target)
}

var _cMTimeMappingMakeEmpty func(target CMTimeRange) CMTimeMapping

// CMTimeMappingMakeEmpty creates a valid time mapping with an empty source.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingMakeEmpty(target:)
func CMTimeMappingMakeEmpty(target CMTimeRange) CMTimeMapping {
	if _cMTimeMappingMakeEmpty == nil {
		panic("CoreMedia: symbol CMTimeMappingMakeEmpty not loaded")
	}
	return _cMTimeMappingMakeEmpty(target)
}

var _cMTimeMappingMakeFromDictionary func(dictionaryRepresentation corefoundation.CFDictionaryRef) CMTimeMapping

// CMTimeMappingMakeFromDictionary creates a time mapping from a dictionary representation.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingMakeFromDictionary(_:)
func CMTimeMappingMakeFromDictionary(dictionaryRepresentation corefoundation.CFDictionaryRef) CMTimeMapping {
	if _cMTimeMappingMakeFromDictionary == nil {
		panic("CoreMedia: symbol CMTimeMappingMakeFromDictionary not loaded")
	}
	return _cMTimeMappingMakeFromDictionary(dictionaryRepresentation)
}

var _cMTimeMappingShow func(mapping CMTimeMapping)

// CMTimeMappingShow prints a description of a time mapping to standard output.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMappingShow(_:)
func CMTimeMappingShow(mapping CMTimeMapping) {
	if _cMTimeMappingShow == nil {
		panic("CoreMedia: symbol CMTimeMappingShow not loaded")
	}
	_cMTimeMappingShow(mapping)
}

var _cMTimeMaximum func(time1 CMTime, time2 CMTime) CMTime

// CMTimeMaximum returns the greater of two time values.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMaximum(_:_:)
func CMTimeMaximum(time1 CMTime, time2 CMTime) CMTime {
	if _cMTimeMaximum == nil {
		panic("CoreMedia: symbol CMTimeMaximum not loaded")
	}
	return _cMTimeMaximum(time1, time2)
}

var _cMTimeMinimum func(time1 CMTime, time2 CMTime) CMTime

// CMTimeMinimum returns the lesser of two time values.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMinimum(_:_:)
func CMTimeMinimum(time1 CMTime, time2 CMTime) CMTime {
	if _cMTimeMinimum == nil {
		panic("CoreMedia: symbol CMTimeMinimum not loaded")
	}
	return _cMTimeMinimum(time1, time2)
}

var _cMTimeMultiply func(time CMTime, multiplier int32) CMTime

// CMTimeMultiply returns the result of multiplying a time by an integer multiplier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMultiply(_:multiplier:)
func CMTimeMultiply(time CMTime, multiplier int32) CMTime {
	if _cMTimeMultiply == nil {
		panic("CoreMedia: symbol CMTimeMultiply not loaded")
	}
	return _cMTimeMultiply(time, multiplier)
}

var _cMTimeMultiplyByFloat64 func(time CMTime, multiplier float64) CMTime

// CMTimeMultiplyByFloat64 returns the result of multiplying a time by a floating-point multiplier.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMultiplyByFloat64(_:multiplier:)
func CMTimeMultiplyByFloat64(time CMTime, multiplier float64) CMTime {
	if _cMTimeMultiplyByFloat64 == nil {
		panic("CoreMedia: symbol CMTimeMultiplyByFloat64 not loaded")
	}
	return _cMTimeMultiplyByFloat64(time, multiplier)
}

var _cMTimeMultiplyByRatio func(time CMTime, multiplier int32, divisor int32) CMTime

// CMTimeMultiplyByRatio returns the result of multiplying a time by an integer multiplier, and then dividing the result by the divisor.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeMultiplyByRatio(_:multiplier:divisor:)
func CMTimeMultiplyByRatio(time CMTime, multiplier int32, divisor int32) CMTime {
	if _cMTimeMultiplyByRatio == nil {
		panic("CoreMedia: symbol CMTimeMultiplyByRatio not loaded")
	}
	return _cMTimeMultiplyByRatio(time, multiplier, divisor)
}

var _cMTimeRangeContainsTime func(range_ CMTimeRange, time CMTime) bool

// CMTimeRangeContainsTime returns a Boolean value that indicates whether a time range contains a time.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeContainsTime(_:time:)
func CMTimeRangeContainsTime(range_ CMTimeRange, time CMTime) bool {
	if _cMTimeRangeContainsTime == nil {
		panic("CoreMedia: symbol CMTimeRangeContainsTime not loaded")
	}
	return _cMTimeRangeContainsTime(range_, time)
}

var _cMTimeRangeContainsTimeRange func(range_ CMTimeRange, otherRange CMTimeRange) bool

// CMTimeRangeContainsTimeRange returns a Boolean value that indicates whether a time range contains another time range.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeContainsTimeRange(_:otherRange:)
func CMTimeRangeContainsTimeRange(range_ CMTimeRange, otherRange CMTimeRange) bool {
	if _cMTimeRangeContainsTimeRange == nil {
		panic("CoreMedia: symbol CMTimeRangeContainsTimeRange not loaded")
	}
	return _cMTimeRangeContainsTimeRange(range_, otherRange)
}

var _cMTimeRangeCopyAsDictionary func(range_ CMTimeRange, allocator corefoundation.CFAllocatorRef) corefoundation.CFDictionaryRef

// CMTimeRangeCopyAsDictionary returns a dictionary representation of a time range.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeCopyAsDictionary(_:allocator:)
func CMTimeRangeCopyAsDictionary(range_ CMTimeRange, allocator corefoundation.CFAllocatorRef) corefoundation.CFDictionaryRef {
	if _cMTimeRangeCopyAsDictionary == nil {
		panic("CoreMedia: symbol CMTimeRangeCopyAsDictionary not loaded")
	}
	return _cMTimeRangeCopyAsDictionary(range_, allocator)
}

var _cMTimeRangeCopyDescription func(allocator corefoundation.CFAllocatorRef, range_ CMTimeRange) corefoundation.CFStringRef

// CMTimeRangeCopyDescription returns a string with a description of a time range.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeCopyDescription(allocator:range:)
func CMTimeRangeCopyDescription(allocator corefoundation.CFAllocatorRef, range_ CMTimeRange) corefoundation.CFStringRef {
	if _cMTimeRangeCopyDescription == nil {
		panic("CoreMedia: symbol CMTimeRangeCopyDescription not loaded")
	}
	return _cMTimeRangeCopyDescription(allocator, range_)
}

var _cMTimeRangeEqual func(range1 CMTimeRange, range2 CMTimeRange) bool

// CMTimeRangeEqual returns a Boolean value that indicates whether two time ranges are equal.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeEqual(_:_:)
func CMTimeRangeEqual(range1 CMTimeRange, range2 CMTimeRange) bool {
	if _cMTimeRangeEqual == nil {
		panic("CoreMedia: symbol CMTimeRangeEqual not loaded")
	}
	return _cMTimeRangeEqual(range1, range2)
}

var _cMTimeRangeFromTimeToTime func(start CMTime, end CMTime) CMTimeRange

// CMTimeRangeFromTimeToTime creates a valid time range from a start and end time.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeFromTimeToTime(start:end:)
func CMTimeRangeFromTimeToTime(start CMTime, end CMTime) CMTimeRange {
	if _cMTimeRangeFromTimeToTime == nil {
		panic("CoreMedia: symbol CMTimeRangeFromTimeToTime not loaded")
	}
	return _cMTimeRangeFromTimeToTime(start, end)
}

var _cMTimeRangeGetEnd func(range_ CMTimeRange) CMTime

// CMTimeRangeGetEnd returns a time value that represents the end of a time range.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeGetEnd(_:)
func CMTimeRangeGetEnd(range_ CMTimeRange) CMTime {
	if _cMTimeRangeGetEnd == nil {
		panic("CoreMedia: symbol CMTimeRangeGetEnd not loaded")
	}
	return _cMTimeRangeGetEnd(range_)
}

var _cMTimeRangeGetIntersection func(range_ CMTimeRange, otherRange CMTimeRange) CMTimeRange

// CMTimeRangeGetIntersection returns a new time range with the time elements that are common between the input.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeGetIntersection(_:otherRange:)
func CMTimeRangeGetIntersection(range_ CMTimeRange, otherRange CMTimeRange) CMTimeRange {
	if _cMTimeRangeGetIntersection == nil {
		panic("CoreMedia: symbol CMTimeRangeGetIntersection not loaded")
	}
	return _cMTimeRangeGetIntersection(range_, otherRange)
}

var _cMTimeRangeGetUnion func(range_ CMTimeRange, otherRange CMTimeRange) CMTimeRange

// CMTimeRangeGetUnion returns a new time range with the time elements of the input.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeGetUnion(_:otherRange:)
func CMTimeRangeGetUnion(range_ CMTimeRange, otherRange CMTimeRange) CMTimeRange {
	if _cMTimeRangeGetUnion == nil {
		panic("CoreMedia: symbol CMTimeRangeGetUnion not loaded")
	}
	return _cMTimeRangeGetUnion(range_, otherRange)
}

var _cMTimeRangeMake func(start CMTime, duration CMTime) CMTimeRange

// CMTimeRangeMake creates a valid time range with a start time and duration.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeMake(start:duration:)
func CMTimeRangeMake(start CMTime, duration CMTime) CMTimeRange {
	if _cMTimeRangeMake == nil {
		panic("CoreMedia: symbol CMTimeRangeMake not loaded")
	}
	return _cMTimeRangeMake(start, duration)
}

var _cMTimeRangeMakeFromDictionary func(dictionaryRepresentation corefoundation.CFDictionaryRef) CMTimeRange

// CMTimeRangeMakeFromDictionary creates a time range from a dictionary representation of its fields.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeMakeFromDictionary(_:)
func CMTimeRangeMakeFromDictionary(dictionaryRepresentation corefoundation.CFDictionaryRef) CMTimeRange {
	if _cMTimeRangeMakeFromDictionary == nil {
		panic("CoreMedia: symbol CMTimeRangeMakeFromDictionary not loaded")
	}
	return _cMTimeRangeMakeFromDictionary(dictionaryRepresentation)
}

var _cMTimeRangeShow func(range_ CMTimeRange)

// CMTimeRangeShow prints a description of the time range to standard error.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeRangeShow(_:)
func CMTimeRangeShow(range_ CMTimeRange) {
	if _cMTimeRangeShow == nil {
		panic("CoreMedia: symbol CMTimeRangeShow not loaded")
	}
	_cMTimeRangeShow(range_)
}

var _cMTimeShow func(time CMTime)

// CMTimeShow prints a description of the time to the console.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeShow(_:)
func CMTimeShow(time CMTime) {
	if _cMTimeShow == nil {
		panic("CoreMedia: symbol CMTimeShow not loaded")
	}
	_cMTimeShow(time)
}

var _cMTimeSubtract func(lhs CMTime, rhs CMTime) CMTime

// CMTimeSubtract returns the difference between two times.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimeSubtract(_:_:)
func CMTimeSubtract(lhs CMTime, rhs CMTime) CMTime {
	if _cMTimeSubtract == nil {
		panic("CoreMedia: symbol CMTimeSubtract not loaded")
	}
	return _cMTimeSubtract(lhs, rhs)
}

var _cMTimebaseAddTimer func(timebase uintptr, timer corefoundation.CFRunLoopTimerRef, runloop corefoundation.CFRunLoopRef) int32

// CMTimebaseAddTimer adds the timer to the list of timers the timebase manages.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseAddTimer(_:timer:runloop:)
func CMTimebaseAddTimer(timebase uintptr, timer corefoundation.CFRunLoopTimerRef, runloop corefoundation.CFRunLoopRef) int32 {
	if _cMTimebaseAddTimer == nil {
		panic("CoreMedia: symbol CMTimebaseAddTimer not loaded")
	}
	return _cMTimebaseAddTimer(timebase, timer, runloop)
}

var _cMTimebaseAddTimerDispatchSource func(timebase uintptr, timerSource uintptr) int32

// CMTimebaseAddTimerDispatchSource adds the timer dispatch source to the list of timers the timebase manages.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseAddTimerDispatchSource(_:timerSource:)
func CMTimebaseAddTimerDispatchSource(timebase uintptr, timerSource dispatch.Source) int32 {
	if _cMTimebaseAddTimerDispatchSource == nil {
		panic("CoreMedia: symbol CMTimebaseAddTimerDispatchSource not loaded")
	}
	return _cMTimebaseAddTimerDispatchSource(timebase, uintptr(timerSource.Handle()))
}

var _cMTimebaseCopySource func(timebase uintptr) CMClockOrTimebaseRef

// CMTimebaseCopySource returns the immediate source — either a clock or timebase — of a timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopySource(_:)
func CMTimebaseCopySource(timebase uintptr) CMClockOrTimebaseRef {
	if _cMTimebaseCopySource == nil {
		panic("CoreMedia: symbol CMTimebaseCopySource not loaded")
	}
	return _cMTimebaseCopySource(timebase)
}

var _cMTimebaseCopySourceClock func(timebase uintptr) uintptr

// CMTimebaseCopySourceClock returns the immediate source clock of a timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopySourceClock(_:)
func CMTimebaseCopySourceClock(timebase uintptr) uintptr {
	if _cMTimebaseCopySourceClock == nil {
		panic("CoreMedia: symbol CMTimebaseCopySourceClock not loaded")
	}
	return _cMTimebaseCopySourceClock(timebase)
}

var _cMTimebaseCopySourceTimebase func(timebase uintptr) uintptr

// CMTimebaseCopySourceTimebase returns the immediate source timebase of a timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopySourceTimebase(_:)
func CMTimebaseCopySourceTimebase(timebase uintptr) uintptr {
	if _cMTimebaseCopySourceTimebase == nil {
		panic("CoreMedia: symbol CMTimebaseCopySourceTimebase not loaded")
	}
	return _cMTimebaseCopySourceTimebase(timebase)
}

var _cMTimebaseCopyUltimateSourceClock func(timebase uintptr) uintptr

// CMTimebaseCopyUltimateSourceClock returns the source clock that’s the source of all of a timebase’s source timebases.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCopyUltimateSourceClock(_:)
func CMTimebaseCopyUltimateSourceClock(timebase uintptr) uintptr {
	if _cMTimebaseCopyUltimateSourceClock == nil {
		panic("CoreMedia: symbol CMTimebaseCopyUltimateSourceClock not loaded")
	}
	return _cMTimebaseCopyUltimateSourceClock(timebase)
}

var _cMTimebaseCreateWithSourceClock func(allocator corefoundation.CFAllocatorRef, sourceClock uintptr, timebaseOut *uintptr) int32

// CMTimebaseCreateWithSourceClock creates a timebase by using a source clock.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCreateWithSourceClock(allocator:sourceClock:timebaseOut:)
func CMTimebaseCreateWithSourceClock(allocator corefoundation.CFAllocatorRef, sourceClock uintptr, timebaseOut *uintptr) int32 {
	if _cMTimebaseCreateWithSourceClock == nil {
		panic("CoreMedia: symbol CMTimebaseCreateWithSourceClock not loaded")
	}
	return _cMTimebaseCreateWithSourceClock(allocator, sourceClock, timebaseOut)
}

var _cMTimebaseCreateWithSourceTimebase func(allocator corefoundation.CFAllocatorRef, sourceTimebase uintptr, timebaseOut *uintptr) int32

// CMTimebaseCreateWithSourceTimebase creates a timebase by using a source timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseCreateWithSourceTimebase(allocator:sourceTimebase:timebaseOut:)
func CMTimebaseCreateWithSourceTimebase(allocator corefoundation.CFAllocatorRef, sourceTimebase uintptr, timebaseOut *uintptr) int32 {
	if _cMTimebaseCreateWithSourceTimebase == nil {
		panic("CoreMedia: symbol CMTimebaseCreateWithSourceTimebase not loaded")
	}
	return _cMTimebaseCreateWithSourceTimebase(allocator, sourceTimebase, timebaseOut)
}

var _cMTimebaseGetEffectiveRate func(timebase uintptr) float64

// CMTimebaseGetEffectiveRate returns the effective rate of a timebase, which combines its rate with the rates of all its host timebases.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetEffectiveRate(_:)
func CMTimebaseGetEffectiveRate(timebase uintptr) float64 {
	if _cMTimebaseGetEffectiveRate == nil {
		panic("CoreMedia: symbol CMTimebaseGetEffectiveRate not loaded")
	}
	return _cMTimebaseGetEffectiveRate(timebase)
}

var _cMTimebaseGetRate func(timebase uintptr) float64

// CMTimebaseGetRate returns the current rate of a timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetRate(_:)
func CMTimebaseGetRate(timebase uintptr) float64 {
	if _cMTimebaseGetRate == nil {
		panic("CoreMedia: symbol CMTimebaseGetRate not loaded")
	}
	return _cMTimebaseGetRate(timebase)
}

var _cMTimebaseGetTime func(timebase uintptr) CMTime

// CMTimebaseGetTime returns the current time from a timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetTime(_:)
func CMTimebaseGetTime(timebase uintptr) CMTime {
	if _cMTimebaseGetTime == nil {
		panic("CoreMedia: symbol CMTimebaseGetTime not loaded")
	}
	return _cMTimebaseGetTime(timebase)
}

var _cMTimebaseGetTimeAndRate func(timebase uintptr, timeOut *CMTime, rateOut *float64) int32

// CMTimebaseGetTimeAndRate returns the current time and rate of a timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetTimeAndRate(_:timeOut:rateOut:)
func CMTimebaseGetTimeAndRate(timebase uintptr, timeOut *CMTime, rateOut *float64) int32 {
	if _cMTimebaseGetTimeAndRate == nil {
		panic("CoreMedia: symbol CMTimebaseGetTimeAndRate not loaded")
	}
	return _cMTimebaseGetTimeAndRate(timebase, timeOut, rateOut)
}

var _cMTimebaseGetTimeWithTimeScale func(timebase uintptr, timescale int32, method CMTimeRoundingMethod) CMTime

// CMTimebaseGetTimeWithTimeScale returns the current time from a timebase in the specified timescale.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetTimeWithTimeScale(_:timescale:method:)
func CMTimebaseGetTimeWithTimeScale(timebase uintptr, timescale int32, method CMTimeRoundingMethod) CMTime {
	if _cMTimebaseGetTimeWithTimeScale == nil {
		panic("CoreMedia: symbol CMTimebaseGetTimeWithTimeScale not loaded")
	}
	return _cMTimebaseGetTimeWithTimeScale(timebase, timescale, method)
}

var _cMTimebaseGetTypeID func() uint

// CMTimebaseGetTypeID returns the Core Foundation type identifier that identifies a timebase object.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseGetTypeID()
func CMTimebaseGetTypeID() uint {
	if _cMTimebaseGetTypeID == nil {
		panic("CoreMedia: symbol CMTimebaseGetTypeID not loaded")
	}
	return _cMTimebaseGetTypeID()
}

var _cMTimebaseNotificationBarrier func(timebase uintptr) int32

// CMTimebaseNotificationBarrier requests that the timebase wait until it isn’t posting notifications.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseNotificationBarrier(_:)
func CMTimebaseNotificationBarrier(timebase uintptr) int32 {
	if _cMTimebaseNotificationBarrier == nil {
		panic("CoreMedia: symbol CMTimebaseNotificationBarrier not loaded")
	}
	return _cMTimebaseNotificationBarrier(timebase)
}

var _cMTimebaseRemoveTimer func(timebase uintptr, timer corefoundation.CFRunLoopTimerRef) int32

// CMTimebaseRemoveTimer removes the timer from the list of timers the timebase manages.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseRemoveTimer(_:timer:)
func CMTimebaseRemoveTimer(timebase uintptr, timer corefoundation.CFRunLoopTimerRef) int32 {
	if _cMTimebaseRemoveTimer == nil {
		panic("CoreMedia: symbol CMTimebaseRemoveTimer not loaded")
	}
	return _cMTimebaseRemoveTimer(timebase, timer)
}

var _cMTimebaseRemoveTimerDispatchSource func(timebase uintptr, timerSource uintptr) int32

// CMTimebaseRemoveTimerDispatchSource removes the timer dispatch source from the list of timers the timebase manages.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseRemoveTimerDispatchSource(_:timerSource:)
func CMTimebaseRemoveTimerDispatchSource(timebase uintptr, timerSource dispatch.Source) int32 {
	if _cMTimebaseRemoveTimerDispatchSource == nil {
		panic("CoreMedia: symbol CMTimebaseRemoveTimerDispatchSource not loaded")
	}
	return _cMTimebaseRemoveTimerDispatchSource(timebase, uintptr(timerSource.Handle()))
}

var _cMTimebaseSetAnchorTime func(timebase uintptr, timebaseTime CMTime, immediateSourceTime CMTime) int32

// CMTimebaseSetAnchorTime sets the time of a timebase at a particular host time.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetAnchorTime(_:timebaseTime:immediateSourceTime:)
func CMTimebaseSetAnchorTime(timebase uintptr, timebaseTime CMTime, immediateSourceTime CMTime) int32 {
	if _cMTimebaseSetAnchorTime == nil {
		panic("CoreMedia: symbol CMTimebaseSetAnchorTime not loaded")
	}
	return _cMTimebaseSetAnchorTime(timebase, timebaseTime, immediateSourceTime)
}

var _cMTimebaseSetRate func(timebase uintptr, rate float64) int32

// CMTimebaseSetRate sets the rate of a timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetRate(_:rate:)
func CMTimebaseSetRate(timebase uintptr, rate float64) int32 {
	if _cMTimebaseSetRate == nil {
		panic("CoreMedia: symbol CMTimebaseSetRate not loaded")
	}
	return _cMTimebaseSetRate(timebase, rate)
}

var _cMTimebaseSetRateAndAnchorTime func(timebase uintptr, rate float64, timebaseTime CMTime, immediateSourceTime CMTime) int32

// CMTimebaseSetRateAndAnchorTime sets the time of a timebase at a particular host time, and changes the rate at exactly that time.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetRateAndAnchorTime(_:rate:anchorTime:immediateSourceTime:)
func CMTimebaseSetRateAndAnchorTime(timebase uintptr, rate float64, timebaseTime CMTime, immediateSourceTime CMTime) int32 {
	if _cMTimebaseSetRateAndAnchorTime == nil {
		panic("CoreMedia: symbol CMTimebaseSetRateAndAnchorTime not loaded")
	}
	return _cMTimebaseSetRateAndAnchorTime(timebase, rate, timebaseTime, immediateSourceTime)
}

var _cMTimebaseSetSourceClock func(timebase uintptr, newSourceClock uintptr) int32

// CMTimebaseSetSourceClock sets the source clock of a timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetSourceClock(_:_:)
func CMTimebaseSetSourceClock(timebase uintptr, newSourceClock uintptr) int32 {
	if _cMTimebaseSetSourceClock == nil {
		panic("CoreMedia: symbol CMTimebaseSetSourceClock not loaded")
	}
	return _cMTimebaseSetSourceClock(timebase, newSourceClock)
}

var _cMTimebaseSetSourceTimebase func(timebase uintptr, newSourceTimebase uintptr) int32

// CMTimebaseSetSourceTimebase sets the source timebase of a timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetSourceTimebase(_:_:)
func CMTimebaseSetSourceTimebase(timebase uintptr, newSourceTimebase uintptr) int32 {
	if _cMTimebaseSetSourceTimebase == nil {
		panic("CoreMedia: symbol CMTimebaseSetSourceTimebase not loaded")
	}
	return _cMTimebaseSetSourceTimebase(timebase, newSourceTimebase)
}

var _cMTimebaseSetTime func(timebase uintptr, time CMTime) int32

// CMTimebaseSetTime sets the current time of a timebase.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTime(_:time:)
func CMTimebaseSetTime(timebase uintptr, time CMTime) int32 {
	if _cMTimebaseSetTime == nil {
		panic("CoreMedia: symbol CMTimebaseSetTime not loaded")
	}
	return _cMTimebaseSetTime(timebase, time)
}

var _cMTimebaseSetTimerDispatchSourceNextFireTime func(timebase uintptr, timerSource uintptr, fireTime CMTime, flags uint32) int32

// CMTimebaseSetTimerDispatchSourceNextFireTime sets the time on the timebase’s timeline at which the timer dispatch source should fire next.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTimerDispatchSourceNextFireTime(_:timerSource:fireTime:flags:)
func CMTimebaseSetTimerDispatchSourceNextFireTime(timebase uintptr, timerSource dispatch.Source, fireTime CMTime, flags uint32) int32 {
	if _cMTimebaseSetTimerDispatchSourceNextFireTime == nil {
		panic("CoreMedia: symbol CMTimebaseSetTimerDispatchSourceNextFireTime not loaded")
	}
	return _cMTimebaseSetTimerDispatchSourceNextFireTime(timebase, uintptr(timerSource.Handle()), fireTime, flags)
}

var _cMTimebaseSetTimerDispatchSourceToFireImmediately func(timebase uintptr, timerSource uintptr) int32

// CMTimebaseSetTimerDispatchSourceToFireImmediately sets the timer dispatch source to fire immediately once, overriding any previous timer call.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTimerDispatchSourceToFireImmediately(_:timerSource:)
func CMTimebaseSetTimerDispatchSourceToFireImmediately(timebase uintptr, timerSource dispatch.Source) int32 {
	if _cMTimebaseSetTimerDispatchSourceToFireImmediately == nil {
		panic("CoreMedia: symbol CMTimebaseSetTimerDispatchSourceToFireImmediately not loaded")
	}
	return _cMTimebaseSetTimerDispatchSourceToFireImmediately(timebase, uintptr(timerSource.Handle()))
}

var _cMTimebaseSetTimerNextFireTime func(timebase uintptr, timer corefoundation.CFRunLoopTimerRef, fireTime CMTime, flags uint32) int32

// CMTimebaseSetTimerNextFireTime sets the time on the timebase’s timeline at which the timer should fire next.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTimerNextFireTime(_:timer:fireTime:flags:)
func CMTimebaseSetTimerNextFireTime(timebase uintptr, timer corefoundation.CFRunLoopTimerRef, fireTime CMTime, flags uint32) int32 {
	if _cMTimebaseSetTimerNextFireTime == nil {
		panic("CoreMedia: symbol CMTimebaseSetTimerNextFireTime not loaded")
	}
	return _cMTimebaseSetTimerNextFireTime(timebase, timer, fireTime, flags)
}

var _cMTimebaseSetTimerToFireImmediately func(timebase uintptr, timer corefoundation.CFRunLoopTimerRef) int32

// CMTimebaseSetTimerToFireImmediately sets the timer to fire immediately once, overriding any previous timer calls.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMTimebaseSetTimerToFireImmediately(_:timer:)
func CMTimebaseSetTimerToFireImmediately(timebase uintptr, timer corefoundation.CFRunLoopTimerRef) int32 {
	if _cMTimebaseSetTimerToFireImmediately == nil {
		panic("CoreMedia: symbol CMTimebaseSetTimerToFireImmediately not loaded")
	}
	return _cMTimebaseSetTimerToFireImmediately(timebase, timer)
}

var _cMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, videoFormatDescription uintptr, stringEncoding uint32, flavor CMImageDescriptionFlavor, blockBufferOut *uintptr) int32

// CMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer copies the contents of a video format description to a buffer in big-endian byte ordering.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer(allocator:videoFormatDescription:stringEncoding:flavor:blockBufferOut:)
func CMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, videoFormatDescription uintptr, stringEncoding uint32, flavor CMImageDescriptionFlavor, blockBufferOut *uintptr) int32 {
	if _cMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer == nil {
		panic("CoreMedia: symbol CMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer not loaded")
	}
	return _cMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer(allocator, videoFormatDescription, stringEncoding, flavor, blockBufferOut)
}

var _cMVideoFormatDescriptionCopyTagCollectionArray func(formatDescription uintptr, tagCollectionsOut *corefoundation.CFArrayRef) int32

// CMVideoFormatDescriptionCopyTagCollectionArray copies the multi-image encoding properties as an array of CMTagCollections.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCopyTagCollectionArray
func CMVideoFormatDescriptionCopyTagCollectionArray(formatDescription uintptr, tagCollectionsOut *corefoundation.CFArrayRef) int32 {
	if _cMVideoFormatDescriptionCopyTagCollectionArray == nil {
		panic("CoreMedia: symbol CMVideoFormatDescriptionCopyTagCollectionArray not loaded")
	}
	return _cMVideoFormatDescriptionCopyTagCollectionArray(formatDescription, tagCollectionsOut)
}

var _cMVideoFormatDescriptionCreate func(allocator corefoundation.CFAllocatorRef, codecType CMVideoCodecType, width int32, height int32, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *uintptr) int32

// CMVideoFormatDescriptionCreate creates a format description for a video media stream.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCreate(allocator:codecType:width:height:extensions:formatDescriptionOut:)
func CMVideoFormatDescriptionCreate(allocator corefoundation.CFAllocatorRef, codecType CMVideoCodecType, width int32, height int32, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *uintptr) int32 {
	if _cMVideoFormatDescriptionCreate == nil {
		panic("CoreMedia: symbol CMVideoFormatDescriptionCreate not loaded")
	}
	return _cMVideoFormatDescriptionCreate(allocator, codecType, width, height, extensions, formatDescriptionOut)
}

var _cMVideoFormatDescriptionCreateForImageBuffer func(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, formatDescriptionOut *uintptr) int32

// CMVideoFormatDescriptionCreateForImageBuffer creates a format description for a video media stream by using an image buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCreateForImageBuffer(allocator:imageBuffer:formatDescriptionOut:)
func CMVideoFormatDescriptionCreateForImageBuffer(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, formatDescriptionOut *uintptr) int32 {
	if _cMVideoFormatDescriptionCreateForImageBuffer == nil {
		panic("CoreMedia: symbol CMVideoFormatDescriptionCreateForImageBuffer not loaded")
	}
	return _cMVideoFormatDescriptionCreateForImageBuffer(allocator, imageBuffer, formatDescriptionOut)
}

var _cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer func(allocator corefoundation.CFAllocatorRef, imageDescriptionBlockBuffer uintptr, stringEncoding uint32, flavor CMImageDescriptionFlavor, formatDescriptionOut *uintptr) int32

// CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer creates a video format description from a big-endian image description inside a buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer(allocator:bigEndianImageDescriptionBlockBuffer:stringEncoding:flavor:formatDescriptionOut:)
func CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer(allocator corefoundation.CFAllocatorRef, imageDescriptionBlockBuffer uintptr, stringEncoding uint32, flavor CMImageDescriptionFlavor, formatDescriptionOut *uintptr) int32 {
	if _cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer == nil {
		panic("CoreMedia: symbol CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer not loaded")
	}
	return _cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer(allocator, imageDescriptionBlockBuffer, stringEncoding, flavor, formatDescriptionOut)
}

var _cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData func(allocator corefoundation.CFAllocatorRef, imageDescriptionData *uint8, size uintptr, stringEncoding uint32, flavor CMImageDescriptionFlavor, formatDescriptionOut *uintptr) int32

// CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData creates a video format description from a big-endian image description structure.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData(allocator:bigEndianImageDescriptionData:size:stringEncoding:flavor:formatDescriptionOut:)
func CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData(allocator corefoundation.CFAllocatorRef, imageDescriptionData *uint8, size uintptr, stringEncoding uint32, flavor CMImageDescriptionFlavor, formatDescriptionOut *uintptr) int32 {
	if _cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData == nil {
		panic("CoreMedia: symbol CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData not loaded")
	}
	return _cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData(allocator, imageDescriptionData, size, stringEncoding, flavor, formatDescriptionOut)
}

var _cMVideoFormatDescriptionCreateFromH264ParameterSets func(allocator corefoundation.CFAllocatorRef, parameterSetCount uintptr, parameterSetPointers *uint8, parameterSetSizes *uintptr, NALUnitHeaderLength int, formatDescriptionOut *uintptr) int32

// CMVideoFormatDescriptionCreateFromH264ParameterSets creates a format description for a video media stream that the parameter set describes.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCreateFromH264ParameterSets(allocator:parameterSetCount:parameterSetPointers:parameterSetSizes:nalUnitHeaderLength:formatDescriptionOut:)
func CMVideoFormatDescriptionCreateFromH264ParameterSets(allocator corefoundation.CFAllocatorRef, parameterSetCount uintptr, parameterSetPointers *uint8, parameterSetSizes *uintptr, NALUnitHeaderLength int, formatDescriptionOut *uintptr) int32 {
	if _cMVideoFormatDescriptionCreateFromH264ParameterSets == nil {
		panic("CoreMedia: symbol CMVideoFormatDescriptionCreateFromH264ParameterSets not loaded")
	}
	return _cMVideoFormatDescriptionCreateFromH264ParameterSets(allocator, parameterSetCount, parameterSetPointers, parameterSetSizes, NALUnitHeaderLength, formatDescriptionOut)
}

var _cMVideoFormatDescriptionCreateFromHEVCParameterSets func(allocator corefoundation.CFAllocatorRef, parameterSetCount uintptr, parameterSetPointers *uint8, parameterSetSizes *uintptr, NALUnitHeaderLength int, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *uintptr) int32

// CMVideoFormatDescriptionCreateFromHEVCParameterSets creates a format description for a video media stream using HEVC (H.265) parameter set NAL units.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionCreateFromHEVCParameterSets(allocator:parameterSetCount:parameterSetPointers:parameterSetSizes:nalUnitHeaderLength:extensions:formatDescriptionOut:)
func CMVideoFormatDescriptionCreateFromHEVCParameterSets(allocator corefoundation.CFAllocatorRef, parameterSetCount uintptr, parameterSetPointers *uint8, parameterSetSizes *uintptr, NALUnitHeaderLength int, extensions corefoundation.CFDictionaryRef, formatDescriptionOut *uintptr) int32 {
	if _cMVideoFormatDescriptionCreateFromHEVCParameterSets == nil {
		panic("CoreMedia: symbol CMVideoFormatDescriptionCreateFromHEVCParameterSets not loaded")
	}
	return _cMVideoFormatDescriptionCreateFromHEVCParameterSets(allocator, parameterSetCount, parameterSetPointers, parameterSetSizes, NALUnitHeaderLength, extensions, formatDescriptionOut)
}

var _cMVideoFormatDescriptionGetCleanAperture func(videoDesc uintptr, originIsAtTopLeft bool) corefoundation.CGRect

// CMVideoFormatDescriptionGetCleanAperture returns a rectangle that defines the portion of the encoded pixel dimensions that represent the image data that’s valid for displaying.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetCleanAperture(_:originIsAtTopLeft:)
func CMVideoFormatDescriptionGetCleanAperture(videoDesc uintptr, originIsAtTopLeft bool) corefoundation.CGRect {
	if _cMVideoFormatDescriptionGetCleanAperture == nil {
		panic("CoreMedia: symbol CMVideoFormatDescriptionGetCleanAperture not loaded")
	}
	return _cMVideoFormatDescriptionGetCleanAperture(videoDesc, originIsAtTopLeft)
}

var _cMVideoFormatDescriptionGetDimensions func(videoDesc uintptr) CMVideoDimensions

// CMVideoFormatDescriptionGetDimensions returns the video dimensions, in encoded pixels.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetDimensions(_:)
func CMVideoFormatDescriptionGetDimensions(videoDesc uintptr) CMVideoDimensions {
	if _cMVideoFormatDescriptionGetDimensions == nil {
		panic("CoreMedia: symbol CMVideoFormatDescriptionGetDimensions not loaded")
	}
	return _cMVideoFormatDescriptionGetDimensions(videoDesc)
}

var _cMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers func() corefoundation.CFArrayRef

// CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers returns an array of keys that you use for video format description extensions, image buffer attachments, and attributes.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers()
func CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers() corefoundation.CFArrayRef {
	if _cMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers == nil {
		panic("CoreMedia: symbol CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers not loaded")
	}
	return _cMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers()
}

var _cMVideoFormatDescriptionGetH264ParameterSetAtIndex func(videoDesc uintptr, parameterSetIndex uintptr, parameterSetPointerOut *uint8, parameterSetSizeOut *uintptr, parameterSetCountOut *uintptr, NALUnitHeaderLengthOut *int) int32

// CMVideoFormatDescriptionGetH264ParameterSetAtIndex returns a parameter set that an H.264 format description contains.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetH264ParameterSetAtIndex(_:parameterSetIndex:parameterSetPointerOut:parameterSetSizeOut:parameterSetCountOut:nalUnitHeaderLengthOut:)
func CMVideoFormatDescriptionGetH264ParameterSetAtIndex(videoDesc uintptr, parameterSetIndex uintptr, parameterSetPointerOut *uint8, parameterSetSizeOut *uintptr, parameterSetCountOut *uintptr, NALUnitHeaderLengthOut []int) int32 {
	if _cMVideoFormatDescriptionGetH264ParameterSetAtIndex == nil {
		panic("CoreMedia: symbol CMVideoFormatDescriptionGetH264ParameterSetAtIndex not loaded")
	}
	return _cMVideoFormatDescriptionGetH264ParameterSetAtIndex(videoDesc, parameterSetIndex, parameterSetPointerOut, parameterSetSizeOut, parameterSetCountOut, unsafe.SliceData(NALUnitHeaderLengthOut))
}

var _cMVideoFormatDescriptionGetHEVCParameterSetAtIndex func(videoDesc uintptr, parameterSetIndex uintptr, parameterSetPointerOut *uint8, parameterSetSizeOut *uintptr, parameterSetCountOut *uintptr, NALUnitHeaderLengthOut *int) int32

// CMVideoFormatDescriptionGetHEVCParameterSetAtIndex returns a parameter set contained in an HEVC (H.265) format description.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetHEVCParameterSetAtIndex(_:parameterSetIndex:parameterSetPointerOut:parameterSetSizeOut:parameterSetCountOut:nalUnitHeaderLengthOut:)
func CMVideoFormatDescriptionGetHEVCParameterSetAtIndex(videoDesc uintptr, parameterSetIndex uintptr, parameterSetPointerOut *uint8, parameterSetSizeOut *uintptr, parameterSetCountOut *uintptr, NALUnitHeaderLengthOut []int) int32 {
	if _cMVideoFormatDescriptionGetHEVCParameterSetAtIndex == nil {
		panic("CoreMedia: symbol CMVideoFormatDescriptionGetHEVCParameterSetAtIndex not loaded")
	}
	return _cMVideoFormatDescriptionGetHEVCParameterSetAtIndex(videoDesc, parameterSetIndex, parameterSetPointerOut, parameterSetSizeOut, parameterSetCountOut, unsafe.SliceData(NALUnitHeaderLengthOut))
}

var _cMVideoFormatDescriptionGetPresentationDimensions func(videoDesc uintptr, usePixelAspectRatio bool, useCleanAperture bool) corefoundation.CGSize

// CMVideoFormatDescriptionGetPresentationDimensions returns the dimensions after taking the pixel aspect ratio and clean aperture into account.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionGetPresentationDimensions(_:usePixelAspectRatio:useCleanAperture:)
func CMVideoFormatDescriptionGetPresentationDimensions(videoDesc uintptr, usePixelAspectRatio bool, useCleanAperture bool) corefoundation.CGSize {
	if _cMVideoFormatDescriptionGetPresentationDimensions == nil {
		panic("CoreMedia: symbol CMVideoFormatDescriptionGetPresentationDimensions not loaded")
	}
	return _cMVideoFormatDescriptionGetPresentationDimensions(videoDesc, usePixelAspectRatio, useCleanAperture)
}

var _cMVideoFormatDescriptionMatchesImageBuffer func(desc uintptr, imageBuffer corevideo.CVImageBufferRef) bool

// CMVideoFormatDescriptionMatchesImageBuffer returns a Boolean value that indicates whether a format description matches an image buffer.
//
// See: https://developer.apple.com/documentation/CoreMedia/CMVideoFormatDescriptionMatchesImageBuffer(_:imageBuffer:)
func CMVideoFormatDescriptionMatchesImageBuffer(desc uintptr, imageBuffer corevideo.CVImageBufferRef) bool {
	if _cMVideoFormatDescriptionMatchesImageBuffer == nil {
		panic("CoreMedia: symbol CMVideoFormatDescriptionMatchesImageBuffer not loaded")
	}
	return _cMVideoFormatDescriptionMatchesImageBuffer(desc, imageBuffer)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
		registerFunc(&_cMAudioDeviceClockCreate, frameworkHandle, "CMAudioDeviceClockCreate")
		registerFunc(&_cMAudioDeviceClockCreateFromAudioDeviceID, frameworkHandle, "CMAudioDeviceClockCreateFromAudioDeviceID")
		registerFunc(&_cMAudioDeviceClockGetAudioDevice, frameworkHandle, "CMAudioDeviceClockGetAudioDevice")
		registerFunc(&_cMAudioDeviceClockSetAudioDeviceID, frameworkHandle, "CMAudioDeviceClockSetAudioDeviceID")
		registerFunc(&_cMAudioDeviceClockSetAudioDeviceUID, frameworkHandle, "CMAudioDeviceClockSetAudioDeviceUID")
		registerFunc(&_cMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer, frameworkHandle, "CMAudioFormatDescriptionCopyAsBigEndianSoundDescriptionBlockBuffer")
		registerFunc(&_cMAudioFormatDescriptionCreate, frameworkHandle, "CMAudioFormatDescriptionCreate")
		registerFunc(&_cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer, frameworkHandle, "CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionBlockBuffer")
		registerFunc(&_cMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData, frameworkHandle, "CMAudioFormatDescriptionCreateFromBigEndianSoundDescriptionData")
		registerFunc(&_cMAudioFormatDescriptionCreateSummary, frameworkHandle, "CMAudioFormatDescriptionCreateSummary")
		registerFunc(&_cMAudioFormatDescriptionEqual, frameworkHandle, "CMAudioFormatDescriptionEqual")
		registerFunc(&_cMAudioFormatDescriptionGetChannelLayout, frameworkHandle, "CMAudioFormatDescriptionGetChannelLayout")
		registerFunc(&_cMAudioFormatDescriptionGetFormatList, frameworkHandle, "CMAudioFormatDescriptionGetFormatList")
		registerFunc(&_cMAudioFormatDescriptionGetMagicCookie, frameworkHandle, "CMAudioFormatDescriptionGetMagicCookie")
		registerFunc(&_cMAudioFormatDescriptionGetMostCompatibleFormat, frameworkHandle, "CMAudioFormatDescriptionGetMostCompatibleFormat")
		registerFunc(&_cMAudioFormatDescriptionGetRichestDecodableFormat, frameworkHandle, "CMAudioFormatDescriptionGetRichestDecodableFormat")
		registerFunc(&_cMAudioFormatDescriptionGetStreamBasicDescription, frameworkHandle, "CMAudioFormatDescriptionGetStreamBasicDescription")
		registerFunc(&_cMAudioSampleBufferCreateReadyWithPacketDescriptions, frameworkHandle, "CMAudioSampleBufferCreateReadyWithPacketDescriptions")
		registerFunc(&_cMAudioSampleBufferCreateWithPacketDescriptions, frameworkHandle, "CMAudioSampleBufferCreateWithPacketDescriptions")
		registerFunc(&_cMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler, frameworkHandle, "CMAudioSampleBufferCreateWithPacketDescriptionsAndMakeDataReadyHandler")
		registerFunc(&_cMBlockBufferAccessDataBytes, frameworkHandle, "CMBlockBufferAccessDataBytes")
		registerFunc(&_cMBlockBufferAppendBufferReference, frameworkHandle, "CMBlockBufferAppendBufferReference")
		registerFunc(&_cMBlockBufferAppendMemoryBlock, frameworkHandle, "CMBlockBufferAppendMemoryBlock")
		registerFunc(&_cMBlockBufferAssureBlockMemory, frameworkHandle, "CMBlockBufferAssureBlockMemory")
		registerFunc(&_cMBlockBufferCopyDataBytes, frameworkHandle, "CMBlockBufferCopyDataBytes")
		registerFunc(&_cMBlockBufferCreateContiguous, frameworkHandle, "CMBlockBufferCreateContiguous")
		registerFunc(&_cMBlockBufferCreateEmpty, frameworkHandle, "CMBlockBufferCreateEmpty")
		registerFunc(&_cMBlockBufferCreateWithBufferReference, frameworkHandle, "CMBlockBufferCreateWithBufferReference")
		registerFunc(&_cMBlockBufferCreateWithMemoryBlock, frameworkHandle, "CMBlockBufferCreateWithMemoryBlock")
		registerFunc(&_cMBlockBufferFillDataBytes, frameworkHandle, "CMBlockBufferFillDataBytes")
		registerFunc(&_cMBlockBufferGetDataLength, frameworkHandle, "CMBlockBufferGetDataLength")
		registerFunc(&_cMBlockBufferGetDataPointer, frameworkHandle, "CMBlockBufferGetDataPointer")
		registerFunc(&_cMBlockBufferGetTypeID, frameworkHandle, "CMBlockBufferGetTypeID")
		registerFunc(&_cMBlockBufferIsEmpty, frameworkHandle, "CMBlockBufferIsEmpty")
		registerFunc(&_cMBlockBufferIsRangeContiguous, frameworkHandle, "CMBlockBufferIsRangeContiguous")
		registerFunc(&_cMBlockBufferReplaceDataBytes, frameworkHandle, "CMBlockBufferReplaceDataBytes")
		registerFunc(&_cMBufferQueueCallForEachBuffer, frameworkHandle, "CMBufferQueueCallForEachBuffer")
		registerFunc(&_cMBufferQueueContainsEndOfData, frameworkHandle, "CMBufferQueueContainsEndOfData")
		registerFunc(&_cMBufferQueueCopyHead, frameworkHandle, "CMBufferQueueCopyHead")
		registerFunc(&_cMBufferQueueCreate, frameworkHandle, "CMBufferQueueCreate")
		registerFunc(&_cMBufferQueueCreateWithHandlers, frameworkHandle, "CMBufferQueueCreateWithHandlers")
		registerFunc(&_cMBufferQueueDequeueAndRetain, frameworkHandle, "CMBufferQueueDequeueAndRetain")
		registerFunc(&_cMBufferQueueDequeueIfDataReadyAndRetain, frameworkHandle, "CMBufferQueueDequeueIfDataReadyAndRetain")
		registerFunc(&_cMBufferQueueEnqueue, frameworkHandle, "CMBufferQueueEnqueue")
		registerFunc(&_cMBufferQueueGetBufferCount, frameworkHandle, "CMBufferQueueGetBufferCount")
		registerFunc(&_cMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS, frameworkHandle, "CMBufferQueueGetCallbacksForSampleBuffersSortedByOutputPTS")
		registerFunc(&_cMBufferQueueGetCallbacksForUnsortedSampleBuffers, frameworkHandle, "CMBufferQueueGetCallbacksForUnsortedSampleBuffers")
		registerFunc(&_cMBufferQueueGetDuration, frameworkHandle, "CMBufferQueueGetDuration")
		registerFunc(&_cMBufferQueueGetEndPresentationTimeStamp, frameworkHandle, "CMBufferQueueGetEndPresentationTimeStamp")
		registerFunc(&_cMBufferQueueGetFirstDecodeTimeStamp, frameworkHandle, "CMBufferQueueGetFirstDecodeTimeStamp")
		registerFunc(&_cMBufferQueueGetFirstPresentationTimeStamp, frameworkHandle, "CMBufferQueueGetFirstPresentationTimeStamp")
		registerFunc(&_cMBufferQueueGetHead, frameworkHandle, "CMBufferQueueGetHead")
		registerFunc(&_cMBufferQueueGetMaxPresentationTimeStamp, frameworkHandle, "CMBufferQueueGetMaxPresentationTimeStamp")
		registerFunc(&_cMBufferQueueGetMinDecodeTimeStamp, frameworkHandle, "CMBufferQueueGetMinDecodeTimeStamp")
		registerFunc(&_cMBufferQueueGetMinPresentationTimeStamp, frameworkHandle, "CMBufferQueueGetMinPresentationTimeStamp")
		registerFunc(&_cMBufferQueueGetTotalSize, frameworkHandle, "CMBufferQueueGetTotalSize")
		registerFunc(&_cMBufferQueueGetTypeID, frameworkHandle, "CMBufferQueueGetTypeID")
		registerFunc(&_cMBufferQueueInstallTrigger, frameworkHandle, "CMBufferQueueInstallTrigger")
		registerFunc(&_cMBufferQueueInstallTriggerHandler, frameworkHandle, "CMBufferQueueInstallTriggerHandler")
		registerFunc(&_cMBufferQueueInstallTriggerHandlerWithIntegerThreshold, frameworkHandle, "CMBufferQueueInstallTriggerHandlerWithIntegerThreshold")
		registerFunc(&_cMBufferQueueInstallTriggerWithIntegerThreshold, frameworkHandle, "CMBufferQueueInstallTriggerWithIntegerThreshold")
		registerFunc(&_cMBufferQueueIsAtEndOfData, frameworkHandle, "CMBufferQueueIsAtEndOfData")
		registerFunc(&_cMBufferQueueIsEmpty, frameworkHandle, "CMBufferQueueIsEmpty")
		registerFunc(&_cMBufferQueueMarkEndOfData, frameworkHandle, "CMBufferQueueMarkEndOfData")
		registerFunc(&_cMBufferQueueRemoveTrigger, frameworkHandle, "CMBufferQueueRemoveTrigger")
		registerFunc(&_cMBufferQueueReset, frameworkHandle, "CMBufferQueueReset")
		registerFunc(&_cMBufferQueueResetWithCallback, frameworkHandle, "CMBufferQueueResetWithCallback")
		registerFunc(&_cMBufferQueueSetValidationCallback, frameworkHandle, "CMBufferQueueSetValidationCallback")
		registerFunc(&_cMBufferQueueSetValidationHandler, frameworkHandle, "CMBufferQueueSetValidationHandler")
		registerFunc(&_cMBufferQueueTestTrigger, frameworkHandle, "CMBufferQueueTestTrigger")
		registerFunc(&_cMClockConvertHostTimeToSystemUnits, frameworkHandle, "CMClockConvertHostTimeToSystemUnits")
		registerFunc(&_cMClockGetAnchorTime, frameworkHandle, "CMClockGetAnchorTime")
		registerFunc(&_cMClockGetHostTimeClock, frameworkHandle, "CMClockGetHostTimeClock")
		registerFunc(&_cMClockGetTime, frameworkHandle, "CMClockGetTime")
		registerFunc(&_cMClockGetTypeID, frameworkHandle, "CMClockGetTypeID")
		registerFunc(&_cMClockInvalidate, frameworkHandle, "CMClockInvalidate")
		registerFunc(&_cMClockMakeHostTimeFromSystemUnits, frameworkHandle, "CMClockMakeHostTimeFromSystemUnits")
		registerFunc(&_cMClockMightDrift, frameworkHandle, "CMClockMightDrift")
		registerFunc(&_cMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer, frameworkHandle, "CMClosedCaptionFormatDescriptionCopyAsBigEndianClosedCaptionDescriptionBlockBuffer")
		registerFunc(&_cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer, frameworkHandle, "CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionBlockBuffer")
		registerFunc(&_cMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData, frameworkHandle, "CMClosedCaptionFormatDescriptionCreateFromBigEndianClosedCaptionDescriptionData")
		registerFunc(&_cMCopyDictionaryOfAttachments, frameworkHandle, "CMCopyDictionaryOfAttachments")
		registerFunc(&_cMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout, frameworkHandle, "CMDoesBigEndianSoundDescriptionRequireLegacyCBRSampleTableLayout")
		registerFunc(&_cMFormatDescriptionCreate, frameworkHandle, "CMFormatDescriptionCreate")
		registerFunc(&_cMFormatDescriptionEqual, frameworkHandle, "CMFormatDescriptionEqual")
		registerFunc(&_cMFormatDescriptionEqualIgnoringExtensionKeys, frameworkHandle, "CMFormatDescriptionEqualIgnoringExtensionKeys")
		registerFunc(&_cMFormatDescriptionGetExtension, frameworkHandle, "CMFormatDescriptionGetExtension")
		registerFunc(&_cMFormatDescriptionGetExtensions, frameworkHandle, "CMFormatDescriptionGetExtensions")
		registerFunc(&_cMFormatDescriptionGetMediaSubType, frameworkHandle, "CMFormatDescriptionGetMediaSubType")
		registerFunc(&_cMFormatDescriptionGetMediaType, frameworkHandle, "CMFormatDescriptionGetMediaType")
		registerFunc(&_cMFormatDescriptionGetTypeID, frameworkHandle, "CMFormatDescriptionGetTypeID")
		registerFunc(&_cMGetAttachment, frameworkHandle, "CMGetAttachment")
		registerFunc(&_cMMemoryPoolCreate, frameworkHandle, "CMMemoryPoolCreate")
		registerFunc(&_cMMemoryPoolFlush, frameworkHandle, "CMMemoryPoolFlush")
		registerFunc(&_cMMemoryPoolGetAllocator, frameworkHandle, "CMMemoryPoolGetAllocator")
		registerFunc(&_cMMemoryPoolGetTypeID, frameworkHandle, "CMMemoryPoolGetTypeID")
		registerFunc(&_cMMemoryPoolInvalidate, frameworkHandle, "CMMemoryPoolInvalidate")
		registerFunc(&_cMMetadataCreateIdentifierForKeyAndKeySpace, frameworkHandle, "CMMetadataCreateIdentifierForKeyAndKeySpace")
		registerFunc(&_cMMetadataCreateKeyFromIdentifier, frameworkHandle, "CMMetadataCreateKeyFromIdentifier")
		registerFunc(&_cMMetadataCreateKeyFromIdentifierAsCFData, frameworkHandle, "CMMetadataCreateKeyFromIdentifierAsCFData")
		registerFunc(&_cMMetadataCreateKeySpaceFromIdentifier, frameworkHandle, "CMMetadataCreateKeySpaceFromIdentifier")
		registerFunc(&_cMMetadataDataTypeRegistryDataTypeConformsToDataType, frameworkHandle, "CMMetadataDataTypeRegistryDataTypeConformsToDataType")
		registerFunc(&_cMMetadataDataTypeRegistryDataTypeIsBaseDataType, frameworkHandle, "CMMetadataDataTypeRegistryDataTypeIsBaseDataType")
		registerFunc(&_cMMetadataDataTypeRegistryDataTypeIsRegistered, frameworkHandle, "CMMetadataDataTypeRegistryDataTypeIsRegistered")
		registerFunc(&_cMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType, frameworkHandle, "CMMetadataDataTypeRegistryGetBaseDataTypeForConformingDataType")
		registerFunc(&_cMMetadataDataTypeRegistryGetBaseDataTypes, frameworkHandle, "CMMetadataDataTypeRegistryGetBaseDataTypes")
		registerFunc(&_cMMetadataDataTypeRegistryGetConformingDataTypes, frameworkHandle, "CMMetadataDataTypeRegistryGetConformingDataTypes")
		registerFunc(&_cMMetadataDataTypeRegistryGetDataTypeDescription, frameworkHandle, "CMMetadataDataTypeRegistryGetDataTypeDescription")
		registerFunc(&_cMMetadataDataTypeRegistryRegisterDataType, frameworkHandle, "CMMetadataDataTypeRegistryRegisterDataType")
		registerFunc(&_cMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer, frameworkHandle, "CMMetadataFormatDescriptionCopyAsBigEndianMetadataDescriptionBlockBuffer")
		registerFunc(&_cMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions, frameworkHandle, "CMMetadataFormatDescriptionCreateByMergingMetadataFormatDescriptions")
		registerFunc(&_cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer, frameworkHandle, "CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionBlockBuffer")
		registerFunc(&_cMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData, frameworkHandle, "CMMetadataFormatDescriptionCreateFromBigEndianMetadataDescriptionData")
		registerFunc(&_cMMetadataFormatDescriptionCreateWithKeys, frameworkHandle, "CMMetadataFormatDescriptionCreateWithKeys")
		registerFunc(&_cMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications, frameworkHandle, "CMMetadataFormatDescriptionCreateWithMetadataFormatDescriptionAndMetadataSpecifications")
		registerFunc(&_cMMetadataFormatDescriptionCreateWithMetadataSpecifications, frameworkHandle, "CMMetadataFormatDescriptionCreateWithMetadataSpecifications")
		registerFunc(&_cMMetadataFormatDescriptionGetIdentifiers, frameworkHandle, "CMMetadataFormatDescriptionGetIdentifiers")
		registerFunc(&_cMMetadataFormatDescriptionGetKeyWithLocalID, frameworkHandle, "CMMetadataFormatDescriptionGetKeyWithLocalID")
		registerFunc(&_cMMuxedFormatDescriptionCreate, frameworkHandle, "CMMuxedFormatDescriptionCreate")
		registerFunc(&_cMPropagateAttachments, frameworkHandle, "CMPropagateAttachments")
		registerFunc(&_cMRemoveAllAttachments, frameworkHandle, "CMRemoveAllAttachments")
		registerFunc(&_cMRemoveAttachment, frameworkHandle, "CMRemoveAttachment")
		registerFunc(&_cMSampleBufferCallBlockForEachSample, frameworkHandle, "CMSampleBufferCallBlockForEachSample")
		registerFunc(&_cMSampleBufferCallForEachSample, frameworkHandle, "CMSampleBufferCallForEachSample")
		registerFunc(&_cMSampleBufferCopyPCMDataIntoAudioBufferList, frameworkHandle, "CMSampleBufferCopyPCMDataIntoAudioBufferList")
		registerFunc(&_cMSampleBufferCopySampleBufferForRange, frameworkHandle, "CMSampleBufferCopySampleBufferForRange")
		registerFunc(&_cMSampleBufferCreate, frameworkHandle, "CMSampleBufferCreate")
		registerFunc(&_cMSampleBufferCreateCopy, frameworkHandle, "CMSampleBufferCreateCopy")
		registerFunc(&_cMSampleBufferCreateCopyWithNewTiming, frameworkHandle, "CMSampleBufferCreateCopyWithNewTiming")
		registerFunc(&_cMSampleBufferCreateForImageBuffer, frameworkHandle, "CMSampleBufferCreateForImageBuffer")
		registerFunc(&_cMSampleBufferCreateForImageBufferWithMakeDataReadyHandler, frameworkHandle, "CMSampleBufferCreateForImageBufferWithMakeDataReadyHandler")
		registerFunc(&_cMSampleBufferCreateForTaggedBufferGroup, frameworkHandle, "CMSampleBufferCreateForTaggedBufferGroup")
		registerFunc(&_cMSampleBufferCreateReady, frameworkHandle, "CMSampleBufferCreateReady")
		registerFunc(&_cMSampleBufferCreateReadyWithImageBuffer, frameworkHandle, "CMSampleBufferCreateReadyWithImageBuffer")
		registerFunc(&_cMSampleBufferCreateWithMakeDataReadyHandler, frameworkHandle, "CMSampleBufferCreateWithMakeDataReadyHandler")
		registerFunc(&_cMSampleBufferDataIsReady, frameworkHandle, "CMSampleBufferDataIsReady")
		registerFunc(&_cMSampleBufferGetAudioBufferListWithRetainedBlockBuffer, frameworkHandle, "CMSampleBufferGetAudioBufferListWithRetainedBlockBuffer")
		registerFunc(&_cMSampleBufferGetAudioStreamPacketDescriptions, frameworkHandle, "CMSampleBufferGetAudioStreamPacketDescriptions")
		registerFunc(&_cMSampleBufferGetAudioStreamPacketDescriptionsPtr, frameworkHandle, "CMSampleBufferGetAudioStreamPacketDescriptionsPtr")
		registerFunc(&_cMSampleBufferGetDataBuffer, frameworkHandle, "CMSampleBufferGetDataBuffer")
		registerFunc(&_cMSampleBufferGetDecodeTimeStamp, frameworkHandle, "CMSampleBufferGetDecodeTimeStamp")
		registerFunc(&_cMSampleBufferGetDuration, frameworkHandle, "CMSampleBufferGetDuration")
		registerFunc(&_cMSampleBufferGetFormatDescription, frameworkHandle, "CMSampleBufferGetFormatDescription")
		registerFunc(&_cMSampleBufferGetImageBuffer, frameworkHandle, "CMSampleBufferGetImageBuffer")
		registerFunc(&_cMSampleBufferGetNumSamples, frameworkHandle, "CMSampleBufferGetNumSamples")
		registerFunc(&_cMSampleBufferGetOutputDecodeTimeStamp, frameworkHandle, "CMSampleBufferGetOutputDecodeTimeStamp")
		registerFunc(&_cMSampleBufferGetOutputDuration, frameworkHandle, "CMSampleBufferGetOutputDuration")
		registerFunc(&_cMSampleBufferGetOutputPresentationTimeStamp, frameworkHandle, "CMSampleBufferGetOutputPresentationTimeStamp")
		registerFunc(&_cMSampleBufferGetOutputSampleTimingInfoArray, frameworkHandle, "CMSampleBufferGetOutputSampleTimingInfoArray")
		registerFunc(&_cMSampleBufferGetPresentationTimeStamp, frameworkHandle, "CMSampleBufferGetPresentationTimeStamp")
		registerFunc(&_cMSampleBufferGetSampleAttachmentsArray, frameworkHandle, "CMSampleBufferGetSampleAttachmentsArray")
		registerFunc(&_cMSampleBufferGetSampleSize, frameworkHandle, "CMSampleBufferGetSampleSize")
		registerFunc(&_cMSampleBufferGetSampleSizeArray, frameworkHandle, "CMSampleBufferGetSampleSizeArray")
		registerFunc(&_cMSampleBufferGetSampleTimingInfo, frameworkHandle, "CMSampleBufferGetSampleTimingInfo")
		registerFunc(&_cMSampleBufferGetSampleTimingInfoArray, frameworkHandle, "CMSampleBufferGetSampleTimingInfoArray")
		registerFunc(&_cMSampleBufferGetTaggedBufferGroup, frameworkHandle, "CMSampleBufferGetTaggedBufferGroup")
		registerFunc(&_cMSampleBufferGetTotalSampleSize, frameworkHandle, "CMSampleBufferGetTotalSampleSize")
		registerFunc(&_cMSampleBufferGetTypeID, frameworkHandle, "CMSampleBufferGetTypeID")
		registerFunc(&_cMSampleBufferHasDataFailed, frameworkHandle, "CMSampleBufferHasDataFailed")
		registerFunc(&_cMSampleBufferInvalidate, frameworkHandle, "CMSampleBufferInvalidate")
		registerFunc(&_cMSampleBufferIsValid, frameworkHandle, "CMSampleBufferIsValid")
		registerFunc(&_cMSampleBufferMakeDataReady, frameworkHandle, "CMSampleBufferMakeDataReady")
		registerFunc(&_cMSampleBufferSetDataBuffer, frameworkHandle, "CMSampleBufferSetDataBuffer")
		registerFunc(&_cMSampleBufferSetDataBufferFromAudioBufferList, frameworkHandle, "CMSampleBufferSetDataBufferFromAudioBufferList")
		registerFunc(&_cMSampleBufferSetDataFailed, frameworkHandle, "CMSampleBufferSetDataFailed")
		registerFunc(&_cMSampleBufferSetDataReady, frameworkHandle, "CMSampleBufferSetDataReady")
		registerFunc(&_cMSampleBufferSetInvalidateCallback, frameworkHandle, "CMSampleBufferSetInvalidateCallback")
		registerFunc(&_cMSampleBufferSetInvalidateHandler, frameworkHandle, "CMSampleBufferSetInvalidateHandler")
		registerFunc(&_cMSampleBufferSetOutputPresentationTimeStamp, frameworkHandle, "CMSampleBufferSetOutputPresentationTimeStamp")
		registerFunc(&_cMSampleBufferTrackDataReadiness, frameworkHandle, "CMSampleBufferTrackDataReadiness")
		registerFunc(&_cMSetAttachment, frameworkHandle, "CMSetAttachment")
		registerFunc(&_cMSetAttachments, frameworkHandle, "CMSetAttachments")
		registerFunc(&_cMSimpleQueueCreate, frameworkHandle, "CMSimpleQueueCreate")
		registerFunc(&_cMSimpleQueueDequeue, frameworkHandle, "CMSimpleQueueDequeue")
		registerFunc(&_cMSimpleQueueEnqueue, frameworkHandle, "CMSimpleQueueEnqueue")
		registerFunc(&_cMSimpleQueueGetCapacity, frameworkHandle, "CMSimpleQueueGetCapacity")
		registerFunc(&_cMSimpleQueueGetCount, frameworkHandle, "CMSimpleQueueGetCount")
		registerFunc(&_cMSimpleQueueGetHead, frameworkHandle, "CMSimpleQueueGetHead")
		registerFunc(&_cMSimpleQueueGetTypeID, frameworkHandle, "CMSimpleQueueGetTypeID")
		registerFunc(&_cMSimpleQueueReset, frameworkHandle, "CMSimpleQueueReset")
		registerFunc(&_cMSwapBigEndianClosedCaptionDescriptionToHost, frameworkHandle, "CMSwapBigEndianClosedCaptionDescriptionToHost")
		registerFunc(&_cMSwapBigEndianImageDescriptionToHost, frameworkHandle, "CMSwapBigEndianImageDescriptionToHost")
		registerFunc(&_cMSwapBigEndianMetadataDescriptionToHost, frameworkHandle, "CMSwapBigEndianMetadataDescriptionToHost")
		registerFunc(&_cMSwapBigEndianSoundDescriptionToHost, frameworkHandle, "CMSwapBigEndianSoundDescriptionToHost")
		registerFunc(&_cMSwapBigEndianTextDescriptionToHost, frameworkHandle, "CMSwapBigEndianTextDescriptionToHost")
		registerFunc(&_cMSwapBigEndianTimeCodeDescriptionToHost, frameworkHandle, "CMSwapBigEndianTimeCodeDescriptionToHost")
		registerFunc(&_cMSwapHostEndianClosedCaptionDescriptionToBig, frameworkHandle, "CMSwapHostEndianClosedCaptionDescriptionToBig")
		registerFunc(&_cMSwapHostEndianImageDescriptionToBig, frameworkHandle, "CMSwapHostEndianImageDescriptionToBig")
		registerFunc(&_cMSwapHostEndianMetadataDescriptionToBig, frameworkHandle, "CMSwapHostEndianMetadataDescriptionToBig")
		registerFunc(&_cMSwapHostEndianSoundDescriptionToBig, frameworkHandle, "CMSwapHostEndianSoundDescriptionToBig")
		registerFunc(&_cMSwapHostEndianTextDescriptionToBig, frameworkHandle, "CMSwapHostEndianTextDescriptionToBig")
		registerFunc(&_cMSwapHostEndianTimeCodeDescriptionToBig, frameworkHandle, "CMSwapHostEndianTimeCodeDescriptionToBig")
		registerFunc(&_cMSyncConvertTime, frameworkHandle, "CMSyncConvertTime")
		registerFunc(&_cMSyncGetRelativeRate, frameworkHandle, "CMSyncGetRelativeRate")
		registerFunc(&_cMSyncGetRelativeRateAndAnchorTime, frameworkHandle, "CMSyncGetRelativeRateAndAnchorTime")
		registerFunc(&_cMSyncGetTime, frameworkHandle, "CMSyncGetTime")
		registerFunc(&_cMSyncMightDrift, frameworkHandle, "CMSyncMightDrift")
		registerFunc(&_cMTagCollectionAddTag, frameworkHandle, "CMTagCollectionAddTag")
		registerFunc(&_cMTagCollectionAddTagsFromArray, frameworkHandle, "CMTagCollectionAddTagsFromArray")
		registerFunc(&_cMTagCollectionAddTagsFromCollection, frameworkHandle, "CMTagCollectionAddTagsFromCollection")
		registerFunc(&_cMTagCollectionApply, frameworkHandle, "CMTagCollectionApply")
		registerFunc(&_cMTagCollectionApplyUntil, frameworkHandle, "CMTagCollectionApplyUntil")
		registerFunc(&_cMTagCollectionContainsCategory, frameworkHandle, "CMTagCollectionContainsCategory")
		registerFunc(&_cMTagCollectionContainsSpecifiedTags, frameworkHandle, "CMTagCollectionContainsSpecifiedTags")
		registerFunc(&_cMTagCollectionContainsTag, frameworkHandle, "CMTagCollectionContainsTag")
		registerFunc(&_cMTagCollectionContainsTagsOfCollection, frameworkHandle, "CMTagCollectionContainsTagsOfCollection")
		registerFunc(&_cMTagCollectionCopyAsData, frameworkHandle, "CMTagCollectionCopyAsData")
		registerFunc(&_cMTagCollectionCopyAsDictionary, frameworkHandle, "CMTagCollectionCopyAsDictionary")
		registerFunc(&_cMTagCollectionCopyDescription, frameworkHandle, "CMTagCollectionCopyDescription")
		registerFunc(&_cMTagCollectionCopyTagsOfCategories, frameworkHandle, "CMTagCollectionCopyTagsOfCategories")
		registerFunc(&_cMTagCollectionCountTagsWithFilterFunction, frameworkHandle, "CMTagCollectionCountTagsWithFilterFunction")
		registerFunc(&_cMTagCollectionCreate, frameworkHandle, "CMTagCollectionCreate")
		registerFunc(&_cMTagCollectionCreateCopy, frameworkHandle, "CMTagCollectionCreateCopy")
		registerFunc(&_cMTagCollectionCreateDifference, frameworkHandle, "CMTagCollectionCreateDifference")
		registerFunc(&_cMTagCollectionCreateExclusiveOr, frameworkHandle, "CMTagCollectionCreateExclusiveOr")
		registerFunc(&_cMTagCollectionCreateFromData, frameworkHandle, "CMTagCollectionCreateFromData")
		registerFunc(&_cMTagCollectionCreateFromDictionary, frameworkHandle, "CMTagCollectionCreateFromDictionary")
		registerFunc(&_cMTagCollectionCreateIntersection, frameworkHandle, "CMTagCollectionCreateIntersection")
		registerFunc(&_cMTagCollectionCreateMutable, frameworkHandle, "CMTagCollectionCreateMutable")
		registerFunc(&_cMTagCollectionCreateMutableCopy, frameworkHandle, "CMTagCollectionCreateMutableCopy")
		registerFunc(&_cMTagCollectionCreateUnion, frameworkHandle, "CMTagCollectionCreateUnion")
		registerFunc(&_cMTagCollectionGetCount, frameworkHandle, "CMTagCollectionGetCount")
		registerFunc(&_cMTagCollectionGetCountOfCategory, frameworkHandle, "CMTagCollectionGetCountOfCategory")
		registerFunc(&_cMTagCollectionGetTags, frameworkHandle, "CMTagCollectionGetTags")
		registerFunc(&_cMTagCollectionGetTagsWithCategory, frameworkHandle, "CMTagCollectionGetTagsWithCategory")
		registerFunc(&_cMTagCollectionGetTagsWithFilterFunction, frameworkHandle, "CMTagCollectionGetTagsWithFilterFunction")
		registerFunc(&_cMTagCollectionGetTypeID, frameworkHandle, "CMTagCollectionGetTypeID")
		registerFunc(&_cMTagCollectionIsEmpty, frameworkHandle, "CMTagCollectionIsEmpty")
		registerFunc(&_cMTagCollectionRemoveAllTags, frameworkHandle, "CMTagCollectionRemoveAllTags")
		registerFunc(&_cMTagCollectionRemoveAllTagsOfCategory, frameworkHandle, "CMTagCollectionRemoveAllTagsOfCategory")
		registerFunc(&_cMTagCollectionRemoveTag, frameworkHandle, "CMTagCollectionRemoveTag")
		registerFunc(&_cMTagCompare, frameworkHandle, "CMTagCompare")
		registerFunc(&_cMTagCopyAsDictionary, frameworkHandle, "CMTagCopyAsDictionary")
		registerFunc(&_cMTagCopyDescription, frameworkHandle, "CMTagCopyDescription")
		registerFunc(&_cMTagEqualToTag, frameworkHandle, "CMTagEqualToTag")
		registerFunc(&_cMTagGetFlagsValue, frameworkHandle, "CMTagGetFlagsValue")
		registerFunc(&_cMTagGetFloat64Value, frameworkHandle, "CMTagGetFloat64Value")
		registerFunc(&_cMTagGetOSTypeValue, frameworkHandle, "CMTagGetOSTypeValue")
		registerFunc(&_cMTagGetSInt64Value, frameworkHandle, "CMTagGetSInt64Value")
		registerFunc(&_cMTagGetValueDataType, frameworkHandle, "CMTagGetValueDataType")
		registerFunc(&_cMTagHasFlagsValue, frameworkHandle, "CMTagHasFlagsValue")
		registerFunc(&_cMTagHasFloat64Value, frameworkHandle, "CMTagHasFloat64Value")
		registerFunc(&_cMTagHasOSTypeValue, frameworkHandle, "CMTagHasOSTypeValue")
		registerFunc(&_cMTagHasSInt64Value, frameworkHandle, "CMTagHasSInt64Value")
		registerFunc(&_cMTagHash, frameworkHandle, "CMTagHash")
		registerFunc(&_cMTagMakeFromDictionary, frameworkHandle, "CMTagMakeFromDictionary")
		registerFunc(&_cMTagMakeWithFlagsValue, frameworkHandle, "CMTagMakeWithFlagsValue")
		registerFunc(&_cMTagMakeWithFloat64Value, frameworkHandle, "CMTagMakeWithFloat64Value")
		registerFunc(&_cMTagMakeWithOSTypeValue, frameworkHandle, "CMTagMakeWithOSTypeValue")
		registerFunc(&_cMTagMakeWithSInt64Value, frameworkHandle, "CMTagMakeWithSInt64Value")
		registerFunc(&_cMTaggedBufferGroupCreate, frameworkHandle, "CMTaggedBufferGroupCreate")
		registerFunc(&_cMTaggedBufferGroupCreateCombined, frameworkHandle, "CMTaggedBufferGroupCreateCombined")
		registerFunc(&_cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup, frameworkHandle, "CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroup")
		registerFunc(&_cMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions, frameworkHandle, "CMTaggedBufferGroupFormatDescriptionCreateForTaggedBufferGroupWithExtensions")
		registerFunc(&_cMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup, frameworkHandle, "CMTaggedBufferGroupFormatDescriptionMatchesTaggedBufferGroup")
		registerFunc(&_cMTaggedBufferGroupGetCMSampleBufferAtIndex, frameworkHandle, "CMTaggedBufferGroupGetCMSampleBufferAtIndex")
		registerFunc(&_cMTaggedBufferGroupGetCMSampleBufferForTag, frameworkHandle, "CMTaggedBufferGroupGetCMSampleBufferForTag")
		registerFunc(&_cMTaggedBufferGroupGetCMSampleBufferForTagCollection, frameworkHandle, "CMTaggedBufferGroupGetCMSampleBufferForTagCollection")
		registerFunc(&_cMTaggedBufferGroupGetCVPixelBufferAtIndex, frameworkHandle, "CMTaggedBufferGroupGetCVPixelBufferAtIndex")
		registerFunc(&_cMTaggedBufferGroupGetCVPixelBufferForTag, frameworkHandle, "CMTaggedBufferGroupGetCVPixelBufferForTag")
		registerFunc(&_cMTaggedBufferGroupGetCVPixelBufferForTagCollection, frameworkHandle, "CMTaggedBufferGroupGetCVPixelBufferForTagCollection")
		registerFunc(&_cMTaggedBufferGroupGetCount, frameworkHandle, "CMTaggedBufferGroupGetCount")
		registerFunc(&_cMTaggedBufferGroupGetNumberOfMatchesForTagCollection, frameworkHandle, "CMTaggedBufferGroupGetNumberOfMatchesForTagCollection")
		registerFunc(&_cMTaggedBufferGroupGetTagCollectionAtIndex, frameworkHandle, "CMTaggedBufferGroupGetTagCollectionAtIndex")
		registerFunc(&_cMTaggedBufferGroupGetTypeID, frameworkHandle, "CMTaggedBufferGroupGetTypeID")
		registerFunc(&_cMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer, frameworkHandle, "CMTextFormatDescriptionCopyAsBigEndianTextDescriptionBlockBuffer")
		registerFunc(&_cMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer, frameworkHandle, "CMTextFormatDescriptionCreateFromBigEndianTextDescriptionBlockBuffer")
		registerFunc(&_cMTextFormatDescriptionCreateFromBigEndianTextDescriptionData, frameworkHandle, "CMTextFormatDescriptionCreateFromBigEndianTextDescriptionData")
		registerFunc(&_cMTextFormatDescriptionGetDefaultStyle, frameworkHandle, "CMTextFormatDescriptionGetDefaultStyle")
		registerFunc(&_cMTextFormatDescriptionGetDefaultTextBox, frameworkHandle, "CMTextFormatDescriptionGetDefaultTextBox")
		registerFunc(&_cMTextFormatDescriptionGetDisplayFlags, frameworkHandle, "CMTextFormatDescriptionGetDisplayFlags")
		registerFunc(&_cMTextFormatDescriptionGetFontName, frameworkHandle, "CMTextFormatDescriptionGetFontName")
		registerFunc(&_cMTextFormatDescriptionGetJustification, frameworkHandle, "CMTextFormatDescriptionGetJustification")
		registerFunc(&_cMTimeAbsoluteValue, frameworkHandle, "CMTimeAbsoluteValue")
		registerFunc(&_cMTimeAdd, frameworkHandle, "CMTimeAdd")
		registerFunc(&_cMTimeClampToRange, frameworkHandle, "CMTimeClampToRange")
		registerFunc(&_cMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer, frameworkHandle, "CMTimeCodeFormatDescriptionCopyAsBigEndianTimeCodeDescriptionBlockBuffer")
		registerFunc(&_cMTimeCodeFormatDescriptionCreate, frameworkHandle, "CMTimeCodeFormatDescriptionCreate")
		registerFunc(&_cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer, frameworkHandle, "CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionBlockBuffer")
		registerFunc(&_cMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData, frameworkHandle, "CMTimeCodeFormatDescriptionCreateFromBigEndianTimeCodeDescriptionData")
		registerFunc(&_cMTimeCodeFormatDescriptionGetFrameDuration, frameworkHandle, "CMTimeCodeFormatDescriptionGetFrameDuration")
		registerFunc(&_cMTimeCodeFormatDescriptionGetFrameQuanta, frameworkHandle, "CMTimeCodeFormatDescriptionGetFrameQuanta")
		registerFunc(&_cMTimeCodeFormatDescriptionGetTimeCodeFlags, frameworkHandle, "CMTimeCodeFormatDescriptionGetTimeCodeFlags")
		registerFunc(&_cMTimeCompare, frameworkHandle, "CMTimeCompare")
		registerFunc(&_cMTimeConvertScale, frameworkHandle, "CMTimeConvertScale")
		registerFunc(&_cMTimeCopyAsDictionary, frameworkHandle, "CMTimeCopyAsDictionary")
		registerFunc(&_cMTimeCopyDescription, frameworkHandle, "CMTimeCopyDescription")
		registerFunc(&_cMTimeFoldIntoRange, frameworkHandle, "CMTimeFoldIntoRange")
		registerFunc(&_cMTimeGetSeconds, frameworkHandle, "CMTimeGetSeconds")
		registerFunc(&_cMTimeMake, frameworkHandle, "CMTimeMake")
		registerFunc(&_cMTimeMakeFromDictionary, frameworkHandle, "CMTimeMakeFromDictionary")
		registerFunc(&_cMTimeMakeWithEpoch, frameworkHandle, "CMTimeMakeWithEpoch")
		registerFunc(&_cMTimeMakeWithSeconds, frameworkHandle, "CMTimeMakeWithSeconds")
		registerFunc(&_cMTimeMapDurationFromRangeToRange, frameworkHandle, "CMTimeMapDurationFromRangeToRange")
		registerFunc(&_cMTimeMapTimeFromRangeToRange, frameworkHandle, "CMTimeMapTimeFromRangeToRange")
		registerFunc(&_cMTimeMappingCopyAsDictionary, frameworkHandle, "CMTimeMappingCopyAsDictionary")
		registerFunc(&_cMTimeMappingCopyDescription, frameworkHandle, "CMTimeMappingCopyDescription")
		registerFunc(&_cMTimeMappingMake, frameworkHandle, "CMTimeMappingMake")
		registerFunc(&_cMTimeMappingMakeEmpty, frameworkHandle, "CMTimeMappingMakeEmpty")
		registerFunc(&_cMTimeMappingMakeFromDictionary, frameworkHandle, "CMTimeMappingMakeFromDictionary")
		registerFunc(&_cMTimeMappingShow, frameworkHandle, "CMTimeMappingShow")
		registerFunc(&_cMTimeMaximum, frameworkHandle, "CMTimeMaximum")
		registerFunc(&_cMTimeMinimum, frameworkHandle, "CMTimeMinimum")
		registerFunc(&_cMTimeMultiply, frameworkHandle, "CMTimeMultiply")
		registerFunc(&_cMTimeMultiplyByFloat64, frameworkHandle, "CMTimeMultiplyByFloat64")
		registerFunc(&_cMTimeMultiplyByRatio, frameworkHandle, "CMTimeMultiplyByRatio")
		registerFunc(&_cMTimeRangeContainsTime, frameworkHandle, "CMTimeRangeContainsTime")
		registerFunc(&_cMTimeRangeContainsTimeRange, frameworkHandle, "CMTimeRangeContainsTimeRange")
		registerFunc(&_cMTimeRangeCopyAsDictionary, frameworkHandle, "CMTimeRangeCopyAsDictionary")
		registerFunc(&_cMTimeRangeCopyDescription, frameworkHandle, "CMTimeRangeCopyDescription")
		registerFunc(&_cMTimeRangeEqual, frameworkHandle, "CMTimeRangeEqual")
		registerFunc(&_cMTimeRangeFromTimeToTime, frameworkHandle, "CMTimeRangeFromTimeToTime")
		registerFunc(&_cMTimeRangeGetEnd, frameworkHandle, "CMTimeRangeGetEnd")
		registerFunc(&_cMTimeRangeGetIntersection, frameworkHandle, "CMTimeRangeGetIntersection")
		registerFunc(&_cMTimeRangeGetUnion, frameworkHandle, "CMTimeRangeGetUnion")
		registerFunc(&_cMTimeRangeMake, frameworkHandle, "CMTimeRangeMake")
		registerFunc(&_cMTimeRangeMakeFromDictionary, frameworkHandle, "CMTimeRangeMakeFromDictionary")
		registerFunc(&_cMTimeRangeShow, frameworkHandle, "CMTimeRangeShow")
		registerFunc(&_cMTimeShow, frameworkHandle, "CMTimeShow")
		registerFunc(&_cMTimeSubtract, frameworkHandle, "CMTimeSubtract")
		registerFunc(&_cMTimebaseAddTimer, frameworkHandle, "CMTimebaseAddTimer")
		registerFunc(&_cMTimebaseAddTimerDispatchSource, frameworkHandle, "CMTimebaseAddTimerDispatchSource")
		registerFunc(&_cMTimebaseCopySource, frameworkHandle, "CMTimebaseCopySource")
		registerFunc(&_cMTimebaseCopySourceClock, frameworkHandle, "CMTimebaseCopySourceClock")
		registerFunc(&_cMTimebaseCopySourceTimebase, frameworkHandle, "CMTimebaseCopySourceTimebase")
		registerFunc(&_cMTimebaseCopyUltimateSourceClock, frameworkHandle, "CMTimebaseCopyUltimateSourceClock")
		registerFunc(&_cMTimebaseCreateWithSourceClock, frameworkHandle, "CMTimebaseCreateWithSourceClock")
		registerFunc(&_cMTimebaseCreateWithSourceTimebase, frameworkHandle, "CMTimebaseCreateWithSourceTimebase")
		registerFunc(&_cMTimebaseGetEffectiveRate, frameworkHandle, "CMTimebaseGetEffectiveRate")
		registerFunc(&_cMTimebaseGetRate, frameworkHandle, "CMTimebaseGetRate")
		registerFunc(&_cMTimebaseGetTime, frameworkHandle, "CMTimebaseGetTime")
		registerFunc(&_cMTimebaseGetTimeAndRate, frameworkHandle, "CMTimebaseGetTimeAndRate")
		registerFunc(&_cMTimebaseGetTimeWithTimeScale, frameworkHandle, "CMTimebaseGetTimeWithTimeScale")
		registerFunc(&_cMTimebaseGetTypeID, frameworkHandle, "CMTimebaseGetTypeID")
		registerFunc(&_cMTimebaseNotificationBarrier, frameworkHandle, "CMTimebaseNotificationBarrier")
		registerFunc(&_cMTimebaseRemoveTimer, frameworkHandle, "CMTimebaseRemoveTimer")
		registerFunc(&_cMTimebaseRemoveTimerDispatchSource, frameworkHandle, "CMTimebaseRemoveTimerDispatchSource")
		registerFunc(&_cMTimebaseSetAnchorTime, frameworkHandle, "CMTimebaseSetAnchorTime")
		registerFunc(&_cMTimebaseSetRate, frameworkHandle, "CMTimebaseSetRate")
		registerFunc(&_cMTimebaseSetRateAndAnchorTime, frameworkHandle, "CMTimebaseSetRateAndAnchorTime")
		registerFunc(&_cMTimebaseSetSourceClock, frameworkHandle, "CMTimebaseSetSourceClock")
		registerFunc(&_cMTimebaseSetSourceTimebase, frameworkHandle, "CMTimebaseSetSourceTimebase")
		registerFunc(&_cMTimebaseSetTime, frameworkHandle, "CMTimebaseSetTime")
		registerFunc(&_cMTimebaseSetTimerDispatchSourceNextFireTime, frameworkHandle, "CMTimebaseSetTimerDispatchSourceNextFireTime")
		registerFunc(&_cMTimebaseSetTimerDispatchSourceToFireImmediately, frameworkHandle, "CMTimebaseSetTimerDispatchSourceToFireImmediately")
		registerFunc(&_cMTimebaseSetTimerNextFireTime, frameworkHandle, "CMTimebaseSetTimerNextFireTime")
		registerFunc(&_cMTimebaseSetTimerToFireImmediately, frameworkHandle, "CMTimebaseSetTimerToFireImmediately")
		registerFunc(&_cMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer, frameworkHandle, "CMVideoFormatDescriptionCopyAsBigEndianImageDescriptionBlockBuffer")
		registerFunc(&_cMVideoFormatDescriptionCopyTagCollectionArray, frameworkHandle, "CMVideoFormatDescriptionCopyTagCollectionArray")
		registerFunc(&_cMVideoFormatDescriptionCreate, frameworkHandle, "CMVideoFormatDescriptionCreate")
		registerFunc(&_cMVideoFormatDescriptionCreateForImageBuffer, frameworkHandle, "CMVideoFormatDescriptionCreateForImageBuffer")
		registerFunc(&_cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer, frameworkHandle, "CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionBlockBuffer")
		registerFunc(&_cMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData, frameworkHandle, "CMVideoFormatDescriptionCreateFromBigEndianImageDescriptionData")
		registerFunc(&_cMVideoFormatDescriptionCreateFromH264ParameterSets, frameworkHandle, "CMVideoFormatDescriptionCreateFromH264ParameterSets")
		registerFunc(&_cMVideoFormatDescriptionCreateFromHEVCParameterSets, frameworkHandle, "CMVideoFormatDescriptionCreateFromHEVCParameterSets")
		registerFunc(&_cMVideoFormatDescriptionGetCleanAperture, frameworkHandle, "CMVideoFormatDescriptionGetCleanAperture")
		registerFunc(&_cMVideoFormatDescriptionGetDimensions, frameworkHandle, "CMVideoFormatDescriptionGetDimensions")
		registerFunc(&_cMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers, frameworkHandle, "CMVideoFormatDescriptionGetExtensionKeysCommonWithImageBuffers")
		registerFunc(&_cMVideoFormatDescriptionGetH264ParameterSetAtIndex, frameworkHandle, "CMVideoFormatDescriptionGetH264ParameterSetAtIndex")
		registerFunc(&_cMVideoFormatDescriptionGetHEVCParameterSetAtIndex, frameworkHandle, "CMVideoFormatDescriptionGetHEVCParameterSetAtIndex")
		registerFunc(&_cMVideoFormatDescriptionGetPresentationDimensions, frameworkHandle, "CMVideoFormatDescriptionGetPresentationDimensions")
		registerFunc(&_cMVideoFormatDescriptionMatchesImageBuffer, frameworkHandle, "CMVideoFormatDescriptionMatchesImageBuffer")
	}

