// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

import (
	"fmt"
	"os"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/dispatch"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: CoreMediaIO: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: CoreMediaIO: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: CoreMediaIO: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _cMIODeviceProcessAVCCommand func(deviceID CMIODeviceID, ioAVCCommand *CMIODeviceAVCCommand) int32

// CMIODeviceProcessAVCCommand.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceProcessAVCCommand(_:_:)
func CMIODeviceProcessAVCCommand(deviceID CMIODeviceID, ioAVCCommand *CMIODeviceAVCCommand) int32 {
	if _cMIODeviceProcessAVCCommand == nil {
		panic("CoreMediaIO: symbol CMIODeviceProcessAVCCommand not loaded")
	}
	return _cMIODeviceProcessAVCCommand(deviceID, ioAVCCommand)
}

var _cMIODeviceProcessRS422Command func(deviceID CMIODeviceID, ioRS422Command *CMIODeviceRS422Command) int32

// CMIODeviceProcessRS422Command.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceProcessRS422Command(_:_:)
func CMIODeviceProcessRS422Command(deviceID CMIODeviceID, ioRS422Command *CMIODeviceRS422Command) int32 {
	if _cMIODeviceProcessRS422Command == nil {
		panic("CoreMediaIO: symbol CMIODeviceProcessRS422Command not loaded")
	}
	return _cMIODeviceProcessRS422Command(deviceID, ioRS422Command)
}

var _cMIODeviceStartStream func(deviceID CMIODeviceID, streamID CMIOStreamID) int32

// CMIODeviceStartStream.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceStartStream(_:_:)
func CMIODeviceStartStream(deviceID CMIODeviceID, streamID CMIOStreamID) int32 {
	if _cMIODeviceStartStream == nil {
		panic("CoreMediaIO: symbol CMIODeviceStartStream not loaded")
	}
	return _cMIODeviceStartStream(deviceID, streamID)
}

var _cMIODeviceStopStream func(deviceID CMIODeviceID, streamID CMIOStreamID) int32

// CMIODeviceStopStream.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceStopStream(_:_:)
func CMIODeviceStopStream(deviceID CMIODeviceID, streamID CMIOStreamID) int32 {
	if _cMIODeviceStopStream == nil {
		panic("CoreMediaIO: symbol CMIODeviceStopStream not loaded")
	}
	return _cMIODeviceStopStream(deviceID, streamID)
}

var _cMIOObjectAddPropertyListener func(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, listener CMIOObjectPropertyListenerProc, clientData unsafe.Pointer) int32

// CMIOObjectAddPropertyListener.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectAddPropertyListener(_:_:_:_:)
func CMIOObjectAddPropertyListener(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, listener CMIOObjectPropertyListenerProc, clientData unsafe.Pointer) int32 {
	if _cMIOObjectAddPropertyListener == nil {
		panic("CoreMediaIO: symbol CMIOObjectAddPropertyListener not loaded")
	}
	return _cMIOObjectAddPropertyListener(objectID, address, listener, clientData)
}

var _cMIOObjectAddPropertyListenerBlock func(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, dispatchQueue uintptr, listener CMIOObjectPropertyListenerBlock) int32

// CMIOObjectAddPropertyListenerBlock.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectAddPropertyListenerBlock(_:_:_:_:)
func CMIOObjectAddPropertyListenerBlock(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, dispatchQueue dispatch.Queue, listener CMIOObjectPropertyListenerBlock) int32 {
	if _cMIOObjectAddPropertyListenerBlock == nil {
		panic("CoreMediaIO: symbol CMIOObjectAddPropertyListenerBlock not loaded")
	}
	return _cMIOObjectAddPropertyListenerBlock(objectID, address, uintptr(dispatchQueue.Handle()), listener)
}

var _cMIOObjectCreate func(owningPlugIn CMIOHardwarePlugInRef, owningObjectID CMIOObjectID, classID CMIOClassID, objectID *CMIOObjectID) int32

// CMIOObjectCreate.
//
// Deprecated: Deprecated since macOS 12.3.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectCreate
func CMIOObjectCreate(owningPlugIn CMIOHardwarePlugInRef, owningObjectID CMIOObjectID, classID CMIOClassID, objectID *CMIOObjectID) int32 {
	if _cMIOObjectCreate == nil {
		panic("CoreMediaIO: symbol CMIOObjectCreate not loaded")
	}
	return _cMIOObjectCreate(owningPlugIn, owningObjectID, classID, objectID)
}

var _cMIOObjectGetPropertyData func(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, qualifierDataSize uint32, qualifierData unsafe.Pointer, dataSize uint32, dataUsed *uint32, data unsafe.Pointer) int32

// CMIOObjectGetPropertyData.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectGetPropertyData(_:_:_:_:_:_:_:)
func CMIOObjectGetPropertyData(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, qualifierDataSize uint32, qualifierData unsafe.Pointer, dataSize uint32, dataUsed *uint32, data unsafe.Pointer) int32 {
	if _cMIOObjectGetPropertyData == nil {
		panic("CoreMediaIO: symbol CMIOObjectGetPropertyData not loaded")
	}
	return _cMIOObjectGetPropertyData(objectID, address, qualifierDataSize, qualifierData, dataSize, dataUsed, data)
}

var _cMIOObjectGetPropertyDataSize func(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, qualifierDataSize uint32, qualifierData unsafe.Pointer, dataSize *uint32) int32

// CMIOObjectGetPropertyDataSize.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectGetPropertyDataSize(_:_:_:_:_:)
func CMIOObjectGetPropertyDataSize(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, qualifierDataSize uint32, qualifierData unsafe.Pointer, dataSize *uint32) int32 {
	if _cMIOObjectGetPropertyDataSize == nil {
		panic("CoreMediaIO: symbol CMIOObjectGetPropertyDataSize not loaded")
	}
	return _cMIOObjectGetPropertyDataSize(objectID, address, qualifierDataSize, qualifierData, dataSize)
}

var _cMIOObjectHasProperty func(objectID CMIOObjectID, address *CMIOObjectPropertyAddress) bool

// CMIOObjectHasProperty.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectHasProperty(_:_:)
func CMIOObjectHasProperty(objectID CMIOObjectID, address *CMIOObjectPropertyAddress) bool {
	if _cMIOObjectHasProperty == nil {
		panic("CoreMediaIO: symbol CMIOObjectHasProperty not loaded")
	}
	return _cMIOObjectHasProperty(objectID, address)
}

var _cMIOObjectIsPropertySettable func(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, isSettable *bool) int32

// CMIOObjectIsPropertySettable.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectIsPropertySettable(_:_:_:)
func CMIOObjectIsPropertySettable(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, isSettable *bool) int32 {
	if _cMIOObjectIsPropertySettable == nil {
		panic("CoreMediaIO: symbol CMIOObjectIsPropertySettable not loaded")
	}
	return _cMIOObjectIsPropertySettable(objectID, address, isSettable)
}

var _cMIOObjectPropertiesChanged func(owningPlugIn CMIOHardwarePlugInRef, objectID CMIOObjectID, numberAddresses uint32, addresses CMIOObjectPropertyAddress) int32

// CMIOObjectPropertiesChanged.
//
// Deprecated: Deprecated since macOS 12.3.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectPropertiesChanged
func CMIOObjectPropertiesChanged(owningPlugIn CMIOHardwarePlugInRef, objectID CMIOObjectID, numberAddresses uint32, addresses CMIOObjectPropertyAddress) int32 {
	if _cMIOObjectPropertiesChanged == nil {
		panic("CoreMediaIO: symbol CMIOObjectPropertiesChanged not loaded")
	}
	return _cMIOObjectPropertiesChanged(owningPlugIn, objectID, numberAddresses, addresses)
}

var _cMIOObjectRemovePropertyListener func(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, listener CMIOObjectPropertyListenerProc, clientData unsafe.Pointer) int32

// CMIOObjectRemovePropertyListener.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectRemovePropertyListener(_:_:_:_:)
func CMIOObjectRemovePropertyListener(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, listener CMIOObjectPropertyListenerProc, clientData unsafe.Pointer) int32 {
	if _cMIOObjectRemovePropertyListener == nil {
		panic("CoreMediaIO: symbol CMIOObjectRemovePropertyListener not loaded")
	}
	return _cMIOObjectRemovePropertyListener(objectID, address, listener, clientData)
}

var _cMIOObjectRemovePropertyListenerBlock func(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, dispatchQueue uintptr, listener CMIOObjectPropertyListenerBlock) int32

// CMIOObjectRemovePropertyListenerBlock.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectRemovePropertyListenerBlock(_:_:_:_:)
func CMIOObjectRemovePropertyListenerBlock(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, dispatchQueue dispatch.Queue, listener CMIOObjectPropertyListenerBlock) int32 {
	if _cMIOObjectRemovePropertyListenerBlock == nil {
		panic("CoreMediaIO: symbol CMIOObjectRemovePropertyListenerBlock not loaded")
	}
	return _cMIOObjectRemovePropertyListenerBlock(objectID, address, uintptr(dispatchQueue.Handle()), listener)
}

var _cMIOObjectSetPropertyData func(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, qualifierDataSize uint32, qualifierData unsafe.Pointer, dataSize uint32, data unsafe.Pointer) int32

// CMIOObjectSetPropertyData.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectSetPropertyData(_:_:_:_:_:_:)
func CMIOObjectSetPropertyData(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, qualifierDataSize uint32, qualifierData unsafe.Pointer, dataSize uint32, data unsafe.Pointer) int32 {
	if _cMIOObjectSetPropertyData == nil {
		panic("CoreMediaIO: symbol CMIOObjectSetPropertyData not loaded")
	}
	return _cMIOObjectSetPropertyData(objectID, address, qualifierDataSize, qualifierData, dataSize, data)
}

var _cMIOObjectShow func(objectID CMIOObjectID)

// CMIOObjectShow.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectShow(_:)
func CMIOObjectShow(objectID CMIOObjectID) {
	if _cMIOObjectShow == nil {
		panic("CoreMediaIO: symbol CMIOObjectShow not loaded")
	}
	_cMIOObjectShow(objectID)
}

var _cMIOObjectsPublishedAndDied func(owningPlugIn CMIOHardwarePlugInRef, owningObjectID CMIOObjectID, numberPublishedCMIOObjects uint32, publishedCMIOObjects CMIOObjectID, numberDeadCMIOObjects uint32, deadCMIOObjects CMIOObjectID) int32

// CMIOObjectsPublishedAndDied.
//
// Deprecated: Deprecated since macOS 12.3.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectsPublishedAndDied
func CMIOObjectsPublishedAndDied(owningPlugIn CMIOHardwarePlugInRef, owningObjectID CMIOObjectID, numberPublishedCMIOObjects uint32, publishedCMIOObjects CMIOObjectID, numberDeadCMIOObjects uint32, deadCMIOObjects CMIOObjectID) int32 {
	if _cMIOObjectsPublishedAndDied == nil {
		panic("CoreMediaIO: symbol CMIOObjectsPublishedAndDied not loaded")
	}
	return _cMIOObjectsPublishedAndDied(owningPlugIn, owningObjectID, numberPublishedCMIOObjects, publishedCMIOObjects, numberDeadCMIOObjects, deadCMIOObjects)
}

var _cMIOSampleBufferCopyNonRequiredAttachments func(sourceSBuf uintptr, destSBuf uintptr, attachmentMode coremedia.CMAttachmentMode) int32

// CMIOSampleBufferCopyNonRequiredAttachments.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferCopyNonRequiredAttachments
func CMIOSampleBufferCopyNonRequiredAttachments(sourceSBuf uintptr, destSBuf uintptr, attachmentMode coremedia.CMAttachmentMode) int32 {
	if _cMIOSampleBufferCopyNonRequiredAttachments == nil {
		panic("CoreMediaIO: symbol CMIOSampleBufferCopyNonRequiredAttachments not loaded")
	}
	return _cMIOSampleBufferCopyNonRequiredAttachments(sourceSBuf, destSBuf, attachmentMode)
}

var _cMIOSampleBufferCopySampleAttachments func(sourceSBuf uintptr, destSBuf uintptr) int32

// CMIOSampleBufferCopySampleAttachments.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferCopySampleAttachments
func CMIOSampleBufferCopySampleAttachments(sourceSBuf uintptr, destSBuf uintptr) int32 {
	if _cMIOSampleBufferCopySampleAttachments == nil {
		panic("CoreMediaIO: symbol CMIOSampleBufferCopySampleAttachments not loaded")
	}
	return _cMIOSampleBufferCopySampleAttachments(sourceSBuf, destSBuf)
}

var _cMIOSampleBufferCreate func(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, formatDescription uintptr, numSamples uint32, numSampleTimingEntries uint32, sampleTimingArray *coremedia.CMSampleTimingInfo, numSampleSizeEntries uint32, sampleSizeArray *uintptr, sequenceNumber uint64, discontinuityFlags uint32, sBufOut *uintptr) int32

// CMIOSampleBufferCreate.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferCreate
func CMIOSampleBufferCreate(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, formatDescription uintptr, numSamples uint32, numSampleTimingEntries uint32, sampleTimingArray *coremedia.CMSampleTimingInfo, numSampleSizeEntries uint32, sampleSizeArray *uintptr, sequenceNumber uint64, discontinuityFlags uint32, sBufOut *uintptr) int32 {
	if _cMIOSampleBufferCreate == nil {
		panic("CoreMediaIO: symbol CMIOSampleBufferCreate not loaded")
	}
	return _cMIOSampleBufferCreate(allocator, dataBuffer, formatDescription, numSamples, numSampleTimingEntries, sampleTimingArray, numSampleSizeEntries, sampleSizeArray, sequenceNumber, discontinuityFlags, sBufOut)
}

var _cMIOSampleBufferCreateForImageBuffer func(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, formatDescription uintptr, sampleTiming *coremedia.CMSampleTimingInfo, sequenceNumber uint64, discontinuityFlags uint32, sBufOut *uintptr) int32

// CMIOSampleBufferCreateForImageBuffer.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferCreateForImageBuffer
func CMIOSampleBufferCreateForImageBuffer(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, formatDescription uintptr, sampleTiming *coremedia.CMSampleTimingInfo, sequenceNumber uint64, discontinuityFlags uint32, sBufOut *uintptr) int32 {
	if _cMIOSampleBufferCreateForImageBuffer == nil {
		panic("CoreMediaIO: symbol CMIOSampleBufferCreateForImageBuffer not loaded")
	}
	return _cMIOSampleBufferCreateForImageBuffer(allocator, imageBuffer, formatDescription, sampleTiming, sequenceNumber, discontinuityFlags, sBufOut)
}

var _cMIOSampleBufferCreateNoDataMarker func(allocator corefoundation.CFAllocatorRef, noDataEvent uint32, formatDescription uintptr, sequenceNumber uint64, discontinuityFlags uint32, sBufOut *uintptr) int32

// CMIOSampleBufferCreateNoDataMarker.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferCreateNoDataMarker
func CMIOSampleBufferCreateNoDataMarker(allocator corefoundation.CFAllocatorRef, noDataEvent uint32, formatDescription uintptr, sequenceNumber uint64, discontinuityFlags uint32, sBufOut *uintptr) int32 {
	if _cMIOSampleBufferCreateNoDataMarker == nil {
		panic("CoreMediaIO: symbol CMIOSampleBufferCreateNoDataMarker not loaded")
	}
	return _cMIOSampleBufferCreateNoDataMarker(allocator, noDataEvent, formatDescription, sequenceNumber, discontinuityFlags, sBufOut)
}

var _cMIOSampleBufferGetDiscontinuityFlags func(sbuf uintptr) uint32

// CMIOSampleBufferGetDiscontinuityFlags.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferGetDiscontinuityFlags
func CMIOSampleBufferGetDiscontinuityFlags(sbuf uintptr) uint32 {
	if _cMIOSampleBufferGetDiscontinuityFlags == nil {
		panic("CoreMediaIO: symbol CMIOSampleBufferGetDiscontinuityFlags not loaded")
	}
	return _cMIOSampleBufferGetDiscontinuityFlags(sbuf)
}

var _cMIOSampleBufferGetSequenceNumber func(sbuf uintptr) uint64

// CMIOSampleBufferGetSequenceNumber.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferGetSequenceNumber
func CMIOSampleBufferGetSequenceNumber(sbuf uintptr) uint64 {
	if _cMIOSampleBufferGetSequenceNumber == nil {
		panic("CoreMediaIO: symbol CMIOSampleBufferGetSequenceNumber not loaded")
	}
	return _cMIOSampleBufferGetSequenceNumber(sbuf)
}

var _cMIOSampleBufferSetDiscontinuityFlags func(allocator corefoundation.CFAllocatorRef, sbuf uintptr, discontinuityFlags uint32)

// CMIOSampleBufferSetDiscontinuityFlags.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferSetDiscontinuityFlags
func CMIOSampleBufferSetDiscontinuityFlags(allocator corefoundation.CFAllocatorRef, sbuf uintptr, discontinuityFlags uint32) {
	if _cMIOSampleBufferSetDiscontinuityFlags == nil {
		panic("CoreMediaIO: symbol CMIOSampleBufferSetDiscontinuityFlags not loaded")
	}
	_cMIOSampleBufferSetDiscontinuityFlags(allocator, sbuf, discontinuityFlags)
}

var _cMIOSampleBufferSetSequenceNumber func(allocator corefoundation.CFAllocatorRef, sbuf uintptr, sequenceNumber uint64)

// CMIOSampleBufferSetSequenceNumber.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferSetSequenceNumber
func CMIOSampleBufferSetSequenceNumber(allocator corefoundation.CFAllocatorRef, sbuf uintptr, sequenceNumber uint64) {
	if _cMIOSampleBufferSetSequenceNumber == nil {
		panic("CoreMediaIO: symbol CMIOSampleBufferSetSequenceNumber not loaded")
	}
	_cMIOSampleBufferSetSequenceNumber(allocator, sbuf, sequenceNumber)
}

var _cMIOStreamClockConvertHostTimeToDeviceTime func(hostTime uint64, clock corefoundation.CFTypeRef) coremedia.CMTime

// CMIOStreamClockConvertHostTimeToDeviceTime.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamClockConvertHostTimeToDeviceTime(_:_:)
func CMIOStreamClockConvertHostTimeToDeviceTime(hostTime uint64, clock corefoundation.CFTypeRef) coremedia.CMTime {
	if _cMIOStreamClockConvertHostTimeToDeviceTime == nil {
		panic("CoreMediaIO: symbol CMIOStreamClockConvertHostTimeToDeviceTime not loaded")
	}
	return _cMIOStreamClockConvertHostTimeToDeviceTime(hostTime, clock)
}

var _cMIOStreamClockCreate func(allocator corefoundation.CFAllocatorRef, clockName corefoundation.CFStringRef, sourceIdentifier unsafe.Pointer, getTimeCallMinimumInterval coremedia.CMTime, numberOfEventsForRateSmoothing uint32, numberOfAveragesForRateSmoothing uint32, clock *corefoundation.CFTypeRef) int32

// CMIOStreamClockCreate.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamClockCreate(_:_:_:_:_:_:_:)
func CMIOStreamClockCreate(allocator corefoundation.CFAllocatorRef, clockName corefoundation.CFStringRef, sourceIdentifier unsafe.Pointer, getTimeCallMinimumInterval coremedia.CMTime, numberOfEventsForRateSmoothing uint32, numberOfAveragesForRateSmoothing uint32, clock *corefoundation.CFTypeRef) int32 {
	if _cMIOStreamClockCreate == nil {
		panic("CoreMediaIO: symbol CMIOStreamClockCreate not loaded")
	}
	return _cMIOStreamClockCreate(allocator, clockName, sourceIdentifier, getTimeCallMinimumInterval, numberOfEventsForRateSmoothing, numberOfAveragesForRateSmoothing, clock)
}

var _cMIOStreamClockInvalidate func(clock corefoundation.CFTypeRef) int32

// CMIOStreamClockInvalidate.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamClockInvalidate(_:)
func CMIOStreamClockInvalidate(clock corefoundation.CFTypeRef) int32 {
	if _cMIOStreamClockInvalidate == nil {
		panic("CoreMediaIO: symbol CMIOStreamClockInvalidate not loaded")
	}
	return _cMIOStreamClockInvalidate(clock)
}

var _cMIOStreamClockPostTimingEvent func(eventTime coremedia.CMTime, hostTime uint64, resynchronize bool, clock corefoundation.CFTypeRef) int32

// CMIOStreamClockPostTimingEvent.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamClockPostTimingEvent(_:_:_:_:)
func CMIOStreamClockPostTimingEvent(eventTime coremedia.CMTime, hostTime uint64, resynchronize bool, clock corefoundation.CFTypeRef) int32 {
	if _cMIOStreamClockPostTimingEvent == nil {
		panic("CoreMediaIO: symbol CMIOStreamClockPostTimingEvent not loaded")
	}
	return _cMIOStreamClockPostTimingEvent(eventTime, hostTime, resynchronize, clock)
}

var _cMIOStreamCopyBufferQueue func(streamID CMIOStreamID, queueAlteredProc CMIODeviceStreamQueueAlteredProc, queueAlteredRefCon unsafe.Pointer, queue *coremedia.CMSimpleQueueRef) int32

// CMIOStreamCopyBufferQueue.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamCopyBufferQueue(_:_:_:_:)
func CMIOStreamCopyBufferQueue(streamID CMIOStreamID, queueAlteredProc CMIODeviceStreamQueueAlteredProc, queueAlteredRefCon unsafe.Pointer, queue *coremedia.CMSimpleQueueRef) int32 {
	if _cMIOStreamCopyBufferQueue == nil {
		panic("CoreMediaIO: symbol CMIOStreamCopyBufferQueue not loaded")
	}
	return _cMIOStreamCopyBufferQueue(streamID, queueAlteredProc, queueAlteredRefCon, queue)
}

var _cMIOStreamDeckCueTo func(streamID CMIOStreamID, frameNumber uint64, playOnCue bool) int32

// CMIOStreamDeckCueTo.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamDeckCueTo(_:_:_:)
func CMIOStreamDeckCueTo(streamID CMIOStreamID, frameNumber uint64, playOnCue bool) int32 {
	if _cMIOStreamDeckCueTo == nil {
		panic("CoreMediaIO: symbol CMIOStreamDeckCueTo not loaded")
	}
	return _cMIOStreamDeckCueTo(streamID, frameNumber, playOnCue)
}

var _cMIOStreamDeckJog func(streamID CMIOStreamID, speed int32) int32

// CMIOStreamDeckJog.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamDeckJog(_:_:)
func CMIOStreamDeckJog(streamID CMIOStreamID, speed int32) int32 {
	if _cMIOStreamDeckJog == nil {
		panic("CoreMediaIO: symbol CMIOStreamDeckJog not loaded")
	}
	return _cMIOStreamDeckJog(streamID, speed)
}

var _cMIOStreamDeckPlay func(streamID CMIOStreamID) int32

// CMIOStreamDeckPlay.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamDeckPlay(_:)
func CMIOStreamDeckPlay(streamID CMIOStreamID) int32 {
	if _cMIOStreamDeckPlay == nil {
		panic("CoreMediaIO: symbol CMIOStreamDeckPlay not loaded")
	}
	return _cMIOStreamDeckPlay(streamID)
}

var _cMIOStreamDeckStop func(streamID CMIOStreamID) int32

// CMIOStreamDeckStop.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamDeckStop(_:)
func CMIOStreamDeckStop(streamID CMIOStreamID) int32 {
	if _cMIOStreamDeckStop == nil {
		panic("CoreMediaIO: symbol CMIOStreamDeckStop not loaded")
	}
	return _cMIOStreamDeckStop(streamID)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_cMIODeviceProcessAVCCommand, frameworkHandle, "CMIODeviceProcessAVCCommand")
	registerFunc(&_cMIODeviceProcessRS422Command, frameworkHandle, "CMIODeviceProcessRS422Command")
	registerFunc(&_cMIODeviceStartStream, frameworkHandle, "CMIODeviceStartStream")
	registerFunc(&_cMIODeviceStopStream, frameworkHandle, "CMIODeviceStopStream")
	registerFunc(&_cMIOObjectAddPropertyListener, frameworkHandle, "CMIOObjectAddPropertyListener")
	registerFunc(&_cMIOObjectAddPropertyListenerBlock, frameworkHandle, "CMIOObjectAddPropertyListenerBlock")
	registerFunc(&_cMIOObjectCreate, frameworkHandle, "CMIOObjectCreate")
	registerFunc(&_cMIOObjectGetPropertyData, frameworkHandle, "CMIOObjectGetPropertyData")
	registerFunc(&_cMIOObjectGetPropertyDataSize, frameworkHandle, "CMIOObjectGetPropertyDataSize")
	registerFunc(&_cMIOObjectHasProperty, frameworkHandle, "CMIOObjectHasProperty")
	registerFunc(&_cMIOObjectIsPropertySettable, frameworkHandle, "CMIOObjectIsPropertySettable")
	registerFunc(&_cMIOObjectPropertiesChanged, frameworkHandle, "CMIOObjectPropertiesChanged")
	registerFunc(&_cMIOObjectRemovePropertyListener, frameworkHandle, "CMIOObjectRemovePropertyListener")
	registerFunc(&_cMIOObjectRemovePropertyListenerBlock, frameworkHandle, "CMIOObjectRemovePropertyListenerBlock")
	registerFunc(&_cMIOObjectSetPropertyData, frameworkHandle, "CMIOObjectSetPropertyData")
	registerFunc(&_cMIOObjectShow, frameworkHandle, "CMIOObjectShow")
	registerFunc(&_cMIOObjectsPublishedAndDied, frameworkHandle, "CMIOObjectsPublishedAndDied")
	registerFunc(&_cMIOSampleBufferCopyNonRequiredAttachments, frameworkHandle, "CMIOSampleBufferCopyNonRequiredAttachments")
	registerFunc(&_cMIOSampleBufferCopySampleAttachments, frameworkHandle, "CMIOSampleBufferCopySampleAttachments")
	registerFunc(&_cMIOSampleBufferCreate, frameworkHandle, "CMIOSampleBufferCreate")
	registerFunc(&_cMIOSampleBufferCreateForImageBuffer, frameworkHandle, "CMIOSampleBufferCreateForImageBuffer")
	registerFunc(&_cMIOSampleBufferCreateNoDataMarker, frameworkHandle, "CMIOSampleBufferCreateNoDataMarker")
	registerFunc(&_cMIOSampleBufferGetDiscontinuityFlags, frameworkHandle, "CMIOSampleBufferGetDiscontinuityFlags")
	registerFunc(&_cMIOSampleBufferGetSequenceNumber, frameworkHandle, "CMIOSampleBufferGetSequenceNumber")
	registerFunc(&_cMIOSampleBufferSetDiscontinuityFlags, frameworkHandle, "CMIOSampleBufferSetDiscontinuityFlags")
	registerFunc(&_cMIOSampleBufferSetSequenceNumber, frameworkHandle, "CMIOSampleBufferSetSequenceNumber")
	registerFunc(&_cMIOStreamClockConvertHostTimeToDeviceTime, frameworkHandle, "CMIOStreamClockConvertHostTimeToDeviceTime")
	registerFunc(&_cMIOStreamClockCreate, frameworkHandle, "CMIOStreamClockCreate")
	registerFunc(&_cMIOStreamClockInvalidate, frameworkHandle, "CMIOStreamClockInvalidate")
	registerFunc(&_cMIOStreamClockPostTimingEvent, frameworkHandle, "CMIOStreamClockPostTimingEvent")
	registerFunc(&_cMIOStreamCopyBufferQueue, frameworkHandle, "CMIOStreamCopyBufferQueue")
	registerFunc(&_cMIOStreamDeckCueTo, frameworkHandle, "CMIOStreamDeckCueTo")
	registerFunc(&_cMIOStreamDeckJog, frameworkHandle, "CMIOStreamDeckJog")
	registerFunc(&_cMIOStreamDeckPlay, frameworkHandle, "CMIOStreamDeckPlay")
	registerFunc(&_cMIOStreamDeckStop, frameworkHandle, "CMIOStreamDeckStop")
}
