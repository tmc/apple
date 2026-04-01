// Code generated from Apple documentation for CoreMediaIO. DO NOT EDIT.

package coremediaio

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/corevideo"
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
		return fmt.Sprintf("CoreMediaIO: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("CoreMediaIO: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("CoreMediaIO: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("CoreMediaIO: register symbol %s: %v", name, r)
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

var _cMIODeviceProcessAVCCommand func(deviceID CMIODeviceID, ioAVCCommand *CMIODeviceAVCCommand) int32
var _cMIODeviceProcessAVCCommandErr error

func tryCMIODeviceProcessAVCCommand(deviceID CMIODeviceID, ioAVCCommand *CMIODeviceAVCCommand) (int32, error) {
	if _cMIODeviceProcessAVCCommand == nil {
		return 0, symbolCallError("CMIODeviceProcessAVCCommand", "10.7", _cMIODeviceProcessAVCCommandErr)
	}
	return _cMIODeviceProcessAVCCommand(deviceID, ioAVCCommand), nil
}

// CMIODeviceProcessAVCCommand.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceProcessAVCCommand(_:_:)
func CMIODeviceProcessAVCCommand(deviceID CMIODeviceID, ioAVCCommand *CMIODeviceAVCCommand) int32 {
	result, callErr := tryCMIODeviceProcessAVCCommand(deviceID, ioAVCCommand)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIODeviceProcessRS422Command func(deviceID CMIODeviceID, ioRS422Command *CMIODeviceRS422Command) int32
var _cMIODeviceProcessRS422CommandErr error

func tryCMIODeviceProcessRS422Command(deviceID CMIODeviceID, ioRS422Command *CMIODeviceRS422Command) (int32, error) {
	if _cMIODeviceProcessRS422Command == nil {
		return 0, symbolCallError("CMIODeviceProcessRS422Command", "10.7", _cMIODeviceProcessRS422CommandErr)
	}
	return _cMIODeviceProcessRS422Command(deviceID, ioRS422Command), nil
}

// CMIODeviceProcessRS422Command.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceProcessRS422Command(_:_:)
func CMIODeviceProcessRS422Command(deviceID CMIODeviceID, ioRS422Command *CMIODeviceRS422Command) int32 {
	result, callErr := tryCMIODeviceProcessRS422Command(deviceID, ioRS422Command)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIODeviceStartStream func(deviceID CMIODeviceID, streamID CMIOStreamID) int32
var _cMIODeviceStartStreamErr error

func tryCMIODeviceStartStream(deviceID CMIODeviceID, streamID CMIOStreamID) (int32, error) {
	if _cMIODeviceStartStream == nil {
		return 0, symbolCallError("CMIODeviceStartStream", "10.7", _cMIODeviceStartStreamErr)
	}
	return _cMIODeviceStartStream(deviceID, streamID), nil
}

// CMIODeviceStartStream.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceStartStream(_:_:)
func CMIODeviceStartStream(deviceID CMIODeviceID, streamID CMIOStreamID) int32 {
	result, callErr := tryCMIODeviceStartStream(deviceID, streamID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIODeviceStopStream func(deviceID CMIODeviceID, streamID CMIOStreamID) int32
var _cMIODeviceStopStreamErr error

func tryCMIODeviceStopStream(deviceID CMIODeviceID, streamID CMIOStreamID) (int32, error) {
	if _cMIODeviceStopStream == nil {
		return 0, symbolCallError("CMIODeviceStopStream", "10.7", _cMIODeviceStopStreamErr)
	}
	return _cMIODeviceStopStream(deviceID, streamID), nil
}

// CMIODeviceStopStream.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIODeviceStopStream(_:_:)
func CMIODeviceStopStream(deviceID CMIODeviceID, streamID CMIOStreamID) int32 {
	result, callErr := tryCMIODeviceStopStream(deviceID, streamID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOObjectAddPropertyListener func(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, listener CMIOObjectPropertyListenerProc, clientData unsafe.Pointer) int32
var _cMIOObjectAddPropertyListenerErr error

func tryCMIOObjectAddPropertyListener(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, listener CMIOObjectPropertyListenerProc, clientData unsafe.Pointer) (int32, error) {
	if _cMIOObjectAddPropertyListener == nil {
		return 0, symbolCallError("CMIOObjectAddPropertyListener", "10.7", _cMIOObjectAddPropertyListenerErr)
	}
	return _cMIOObjectAddPropertyListener(objectID, address, listener, clientData), nil
}

// CMIOObjectAddPropertyListener.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectAddPropertyListener(_:_:_:_:)
func CMIOObjectAddPropertyListener(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, listener CMIOObjectPropertyListenerProc, clientData unsafe.Pointer) int32 {
	result, callErr := tryCMIOObjectAddPropertyListener(objectID, address, listener, clientData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOObjectAddPropertyListenerBlock func(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, dispatchQueue uintptr, listener unsafe.Pointer) int32
var _cMIOObjectAddPropertyListenerBlockErr error

func tryCMIOObjectAddPropertyListenerBlock(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, dispatchQueue dispatch.Queue, listener CMIOObjectPropertyListenerBlock) (int32, error) {
	if _cMIOObjectAddPropertyListenerBlock == nil {
		return 0, symbolCallError("CMIOObjectAddPropertyListenerBlock", "10.8", _cMIOObjectAddPropertyListenerBlockErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 uint32, arg1 *CMIOObjectPropertyAddress) { listener(arg0, arg1) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _cMIOObjectAddPropertyListenerBlock(objectID, address, uintptr(dispatchQueue.Handle()), _block0), nil
}

// CMIOObjectAddPropertyListenerBlock.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectAddPropertyListenerBlock(_:_:_:_:)
func CMIOObjectAddPropertyListenerBlock(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, dispatchQueue dispatch.Queue, listener CMIOObjectPropertyListenerBlock) int32 {
	result, callErr := tryCMIOObjectAddPropertyListenerBlock(objectID, address, dispatchQueue, listener)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOObjectCreate func(owningPlugIn CMIOHardwarePlugInRef, owningObjectID CMIOObjectID, classID CMIOClassID, objectID *CMIOObjectID) int32
var _cMIOObjectCreateErr error

func tryCMIOObjectCreate(owningPlugIn CMIOHardwarePlugInRef, owningObjectID CMIOObjectID, classID CMIOClassID, objectID *CMIOObjectID) (int32, error) {
	if _cMIOObjectCreate == nil {
		return 0, symbolCallError("CMIOObjectCreate", "10.7", _cMIOObjectCreateErr)
	}
	return _cMIOObjectCreate(owningPlugIn, owningObjectID, classID, objectID), nil
}

// CMIOObjectCreate.
//
// Deprecated: Deprecated since macOS 12.3.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectCreate
func CMIOObjectCreate(owningPlugIn CMIOHardwarePlugInRef, owningObjectID CMIOObjectID, classID CMIOClassID, objectID *CMIOObjectID) int32 {
	result, callErr := tryCMIOObjectCreate(owningPlugIn, owningObjectID, classID, objectID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOObjectGetPropertyData func(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, qualifierDataSize uint32, qualifierData unsafe.Pointer, dataSize uint32, dataUsed *uint32, data unsafe.Pointer) int32
var _cMIOObjectGetPropertyDataErr error

func tryCMIOObjectGetPropertyData(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, qualifierDataSize uint32, qualifierData unsafe.Pointer, dataSize uint32, dataUsed *uint32, data unsafe.Pointer) (int32, error) {
	if _cMIOObjectGetPropertyData == nil {
		return 0, symbolCallError("CMIOObjectGetPropertyData", "10.7", _cMIOObjectGetPropertyDataErr)
	}
	return _cMIOObjectGetPropertyData(objectID, address, qualifierDataSize, qualifierData, dataSize, dataUsed, data), nil
}

// CMIOObjectGetPropertyData.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectGetPropertyData(_:_:_:_:_:_:_:)
func CMIOObjectGetPropertyData(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, qualifierDataSize uint32, qualifierData unsafe.Pointer, dataSize uint32, dataUsed *uint32, data unsafe.Pointer) int32 {
	result, callErr := tryCMIOObjectGetPropertyData(objectID, address, qualifierDataSize, qualifierData, dataSize, dataUsed, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOObjectGetPropertyDataSize func(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, qualifierDataSize uint32, qualifierData unsafe.Pointer, dataSize *uint32) int32
var _cMIOObjectGetPropertyDataSizeErr error

func tryCMIOObjectGetPropertyDataSize(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, qualifierDataSize uint32, qualifierData unsafe.Pointer, dataSize *uint32) (int32, error) {
	if _cMIOObjectGetPropertyDataSize == nil {
		return 0, symbolCallError("CMIOObjectGetPropertyDataSize", "10.7", _cMIOObjectGetPropertyDataSizeErr)
	}
	return _cMIOObjectGetPropertyDataSize(objectID, address, qualifierDataSize, qualifierData, dataSize), nil
}

// CMIOObjectGetPropertyDataSize.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectGetPropertyDataSize(_:_:_:_:_:)
func CMIOObjectGetPropertyDataSize(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, qualifierDataSize uint32, qualifierData unsafe.Pointer, dataSize *uint32) int32 {
	result, callErr := tryCMIOObjectGetPropertyDataSize(objectID, address, qualifierDataSize, qualifierData, dataSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOObjectHasProperty func(objectID CMIOObjectID, address *CMIOObjectPropertyAddress) bool
var _cMIOObjectHasPropertyErr error

func tryCMIOObjectHasProperty(objectID CMIOObjectID, address *CMIOObjectPropertyAddress) (bool, error) {
	if _cMIOObjectHasProperty == nil {
		return false, symbolCallError("CMIOObjectHasProperty", "10.7", _cMIOObjectHasPropertyErr)
	}
	return _cMIOObjectHasProperty(objectID, address), nil
}

// CMIOObjectHasProperty.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectHasProperty(_:_:)
func CMIOObjectHasProperty(objectID CMIOObjectID, address *CMIOObjectPropertyAddress) bool {
	result, callErr := tryCMIOObjectHasProperty(objectID, address)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOObjectIsPropertySettable func(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, isSettable *bool) int32
var _cMIOObjectIsPropertySettableErr error

func tryCMIOObjectIsPropertySettable(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, isSettable *bool) (int32, error) {
	if _cMIOObjectIsPropertySettable == nil {
		return 0, symbolCallError("CMIOObjectIsPropertySettable", "10.7", _cMIOObjectIsPropertySettableErr)
	}
	return _cMIOObjectIsPropertySettable(objectID, address, isSettable), nil
}

// CMIOObjectIsPropertySettable.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectIsPropertySettable(_:_:_:)
func CMIOObjectIsPropertySettable(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, isSettable *bool) int32 {
	result, callErr := tryCMIOObjectIsPropertySettable(objectID, address, isSettable)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOObjectPropertiesChanged func(owningPlugIn CMIOHardwarePlugInRef, objectID CMIOObjectID, numberAddresses uint32, addresses CMIOObjectPropertyAddress) int32
var _cMIOObjectPropertiesChangedErr error

func tryCMIOObjectPropertiesChanged(owningPlugIn CMIOHardwarePlugInRef, objectID CMIOObjectID, numberAddresses uint32, addresses CMIOObjectPropertyAddress) (int32, error) {
	if _cMIOObjectPropertiesChanged == nil {
		return 0, symbolCallError("CMIOObjectPropertiesChanged", "10.7", _cMIOObjectPropertiesChangedErr)
	}
	return _cMIOObjectPropertiesChanged(owningPlugIn, objectID, numberAddresses, addresses), nil
}

// CMIOObjectPropertiesChanged.
//
// Deprecated: Deprecated since macOS 12.3.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectPropertiesChanged
func CMIOObjectPropertiesChanged(owningPlugIn CMIOHardwarePlugInRef, objectID CMIOObjectID, numberAddresses uint32, addresses CMIOObjectPropertyAddress) int32 {
	result, callErr := tryCMIOObjectPropertiesChanged(owningPlugIn, objectID, numberAddresses, addresses)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOObjectRemovePropertyListener func(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, listener CMIOObjectPropertyListenerProc, clientData unsafe.Pointer) int32
var _cMIOObjectRemovePropertyListenerErr error

func tryCMIOObjectRemovePropertyListener(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, listener CMIOObjectPropertyListenerProc, clientData unsafe.Pointer) (int32, error) {
	if _cMIOObjectRemovePropertyListener == nil {
		return 0, symbolCallError("CMIOObjectRemovePropertyListener", "10.7", _cMIOObjectRemovePropertyListenerErr)
	}
	return _cMIOObjectRemovePropertyListener(objectID, address, listener, clientData), nil
}

// CMIOObjectRemovePropertyListener.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectRemovePropertyListener(_:_:_:_:)
func CMIOObjectRemovePropertyListener(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, listener CMIOObjectPropertyListenerProc, clientData unsafe.Pointer) int32 {
	result, callErr := tryCMIOObjectRemovePropertyListener(objectID, address, listener, clientData)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOObjectRemovePropertyListenerBlock func(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, dispatchQueue uintptr, listener unsafe.Pointer) int32
var _cMIOObjectRemovePropertyListenerBlockErr error

func tryCMIOObjectRemovePropertyListenerBlock(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, dispatchQueue dispatch.Queue, listener CMIOObjectPropertyListenerBlock) (int32, error) {
	if _cMIOObjectRemovePropertyListenerBlock == nil {
		return 0, symbolCallError("CMIOObjectRemovePropertyListenerBlock", "10.8", _cMIOObjectRemovePropertyListenerBlockErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, arg0 uint32, arg1 *CMIOObjectPropertyAddress) { listener(arg0, arg1) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _cMIOObjectRemovePropertyListenerBlock(objectID, address, uintptr(dispatchQueue.Handle()), _block0), nil
}

// CMIOObjectRemovePropertyListenerBlock.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectRemovePropertyListenerBlock(_:_:_:_:)
func CMIOObjectRemovePropertyListenerBlock(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, dispatchQueue dispatch.Queue, listener CMIOObjectPropertyListenerBlock) int32 {
	result, callErr := tryCMIOObjectRemovePropertyListenerBlock(objectID, address, dispatchQueue, listener)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOObjectSetPropertyData func(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, qualifierDataSize uint32, qualifierData unsafe.Pointer, dataSize uint32, data unsafe.Pointer) int32
var _cMIOObjectSetPropertyDataErr error

func tryCMIOObjectSetPropertyData(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, qualifierDataSize uint32, qualifierData unsafe.Pointer, dataSize uint32, data unsafe.Pointer) (int32, error) {
	if _cMIOObjectSetPropertyData == nil {
		return 0, symbolCallError("CMIOObjectSetPropertyData", "10.7", _cMIOObjectSetPropertyDataErr)
	}
	return _cMIOObjectSetPropertyData(objectID, address, qualifierDataSize, qualifierData, dataSize, data), nil
}

// CMIOObjectSetPropertyData.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectSetPropertyData(_:_:_:_:_:_:)
func CMIOObjectSetPropertyData(objectID CMIOObjectID, address *CMIOObjectPropertyAddress, qualifierDataSize uint32, qualifierData unsafe.Pointer, dataSize uint32, data unsafe.Pointer) int32 {
	result, callErr := tryCMIOObjectSetPropertyData(objectID, address, qualifierDataSize, qualifierData, dataSize, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOObjectShow func(objectID CMIOObjectID)
var _cMIOObjectShowErr error

func tryCMIOObjectShow(objectID CMIOObjectID) error {
	if _cMIOObjectShow == nil {
		return symbolCallError("CMIOObjectShow", "10.7", _cMIOObjectShowErr)
	}
	_cMIOObjectShow(objectID)
	return nil
}

// CMIOObjectShow.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectShow(_:)
func CMIOObjectShow(objectID CMIOObjectID) {
	if callErr := tryCMIOObjectShow(objectID); callErr != nil {
		panic(callErr)
	}
}

var _cMIOObjectsPublishedAndDied func(owningPlugIn CMIOHardwarePlugInRef, owningObjectID CMIOObjectID, numberPublishedCMIOObjects uint32, publishedCMIOObjects CMIOObjectID, numberDeadCMIOObjects uint32, deadCMIOObjects CMIOObjectID) int32
var _cMIOObjectsPublishedAndDiedErr error

func tryCMIOObjectsPublishedAndDied(owningPlugIn CMIOHardwarePlugInRef, owningObjectID CMIOObjectID, numberPublishedCMIOObjects uint32, publishedCMIOObjects CMIOObjectID, numberDeadCMIOObjects uint32, deadCMIOObjects CMIOObjectID) (int32, error) {
	if _cMIOObjectsPublishedAndDied == nil {
		return 0, symbolCallError("CMIOObjectsPublishedAndDied", "10.7", _cMIOObjectsPublishedAndDiedErr)
	}
	return _cMIOObjectsPublishedAndDied(owningPlugIn, owningObjectID, numberPublishedCMIOObjects, publishedCMIOObjects, numberDeadCMIOObjects, deadCMIOObjects), nil
}

// CMIOObjectsPublishedAndDied.
//
// Deprecated: Deprecated since macOS 12.3.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOObjectsPublishedAndDied
func CMIOObjectsPublishedAndDied(owningPlugIn CMIOHardwarePlugInRef, owningObjectID CMIOObjectID, numberPublishedCMIOObjects uint32, publishedCMIOObjects CMIOObjectID, numberDeadCMIOObjects uint32, deadCMIOObjects CMIOObjectID) int32 {
	result, callErr := tryCMIOObjectsPublishedAndDied(owningPlugIn, owningObjectID, numberPublishedCMIOObjects, publishedCMIOObjects, numberDeadCMIOObjects, deadCMIOObjects)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOSampleBufferCopyNonRequiredAttachments func(sourceSBuf uintptr, destSBuf uintptr, attachmentMode coremedia.CMAttachmentMode) int32
var _cMIOSampleBufferCopyNonRequiredAttachmentsErr error

func tryCMIOSampleBufferCopyNonRequiredAttachments(sourceSBuf uintptr, destSBuf uintptr, attachmentMode coremedia.CMAttachmentMode) (int32, error) {
	if _cMIOSampleBufferCopyNonRequiredAttachments == nil {
		return 0, symbolCallError("CMIOSampleBufferCopyNonRequiredAttachments", "10.7", _cMIOSampleBufferCopyNonRequiredAttachmentsErr)
	}
	return _cMIOSampleBufferCopyNonRequiredAttachments(sourceSBuf, destSBuf, attachmentMode), nil
}

// CMIOSampleBufferCopyNonRequiredAttachments.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferCopyNonRequiredAttachments
func CMIOSampleBufferCopyNonRequiredAttachments(sourceSBuf uintptr, destSBuf uintptr, attachmentMode coremedia.CMAttachmentMode) int32 {
	result, callErr := tryCMIOSampleBufferCopyNonRequiredAttachments(sourceSBuf, destSBuf, attachmentMode)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOSampleBufferCopySampleAttachments func(sourceSBuf uintptr, destSBuf uintptr) int32
var _cMIOSampleBufferCopySampleAttachmentsErr error

func tryCMIOSampleBufferCopySampleAttachments(sourceSBuf uintptr, destSBuf uintptr) (int32, error) {
	if _cMIOSampleBufferCopySampleAttachments == nil {
		return 0, symbolCallError("CMIOSampleBufferCopySampleAttachments", "10.7", _cMIOSampleBufferCopySampleAttachmentsErr)
	}
	return _cMIOSampleBufferCopySampleAttachments(sourceSBuf, destSBuf), nil
}

// CMIOSampleBufferCopySampleAttachments.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferCopySampleAttachments
func CMIOSampleBufferCopySampleAttachments(sourceSBuf uintptr, destSBuf uintptr) int32 {
	result, callErr := tryCMIOSampleBufferCopySampleAttachments(sourceSBuf, destSBuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOSampleBufferCreate func(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, formatDescription uintptr, numSamples uint32, numSampleTimingEntries uint32, sampleTimingArray *coremedia.CMSampleTimingInfo, numSampleSizeEntries uint32, sampleSizeArray *uintptr, sequenceNumber uint64, discontinuityFlags uint32, sBufOut *uintptr) int32
var _cMIOSampleBufferCreateErr error

func tryCMIOSampleBufferCreate(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, formatDescription uintptr, numSamples uint32, numSampleTimingEntries uint32, sampleTimingArray *coremedia.CMSampleTimingInfo, numSampleSizeEntries uint32, sampleSizeArray *uintptr, sequenceNumber uint64, discontinuityFlags uint32, sBufOut *uintptr) (int32, error) {
	if _cMIOSampleBufferCreate == nil {
		return 0, symbolCallError("CMIOSampleBufferCreate", "10.7", _cMIOSampleBufferCreateErr)
	}
	return _cMIOSampleBufferCreate(allocator, dataBuffer, formatDescription, numSamples, numSampleTimingEntries, sampleTimingArray, numSampleSizeEntries, sampleSizeArray, sequenceNumber, discontinuityFlags, sBufOut), nil
}

// CMIOSampleBufferCreate.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferCreate
func CMIOSampleBufferCreate(allocator corefoundation.CFAllocatorRef, dataBuffer uintptr, formatDescription uintptr, numSamples uint32, numSampleTimingEntries uint32, sampleTimingArray *coremedia.CMSampleTimingInfo, numSampleSizeEntries uint32, sampleSizeArray *uintptr, sequenceNumber uint64, discontinuityFlags uint32, sBufOut *uintptr) int32 {
	result, callErr := tryCMIOSampleBufferCreate(allocator, dataBuffer, formatDescription, numSamples, numSampleTimingEntries, sampleTimingArray, numSampleSizeEntries, sampleSizeArray, sequenceNumber, discontinuityFlags, sBufOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOSampleBufferCreateForImageBuffer func(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, formatDescription uintptr, sampleTiming *coremedia.CMSampleTimingInfo, sequenceNumber uint64, discontinuityFlags uint32, sBufOut *uintptr) int32
var _cMIOSampleBufferCreateForImageBufferErr error

func tryCMIOSampleBufferCreateForImageBuffer(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, formatDescription uintptr, sampleTiming *coremedia.CMSampleTimingInfo, sequenceNumber uint64, discontinuityFlags uint32, sBufOut *uintptr) (int32, error) {
	if _cMIOSampleBufferCreateForImageBuffer == nil {
		return 0, symbolCallError("CMIOSampleBufferCreateForImageBuffer", "10.7", _cMIOSampleBufferCreateForImageBufferErr)
	}
	return _cMIOSampleBufferCreateForImageBuffer(allocator, imageBuffer, formatDescription, sampleTiming, sequenceNumber, discontinuityFlags, sBufOut), nil
}

// CMIOSampleBufferCreateForImageBuffer.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferCreateForImageBuffer
func CMIOSampleBufferCreateForImageBuffer(allocator corefoundation.CFAllocatorRef, imageBuffer corevideo.CVImageBufferRef, formatDescription uintptr, sampleTiming *coremedia.CMSampleTimingInfo, sequenceNumber uint64, discontinuityFlags uint32, sBufOut *uintptr) int32 {
	result, callErr := tryCMIOSampleBufferCreateForImageBuffer(allocator, imageBuffer, formatDescription, sampleTiming, sequenceNumber, discontinuityFlags, sBufOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOSampleBufferCreateNoDataMarker func(allocator corefoundation.CFAllocatorRef, noDataEvent uint32, formatDescription uintptr, sequenceNumber uint64, discontinuityFlags uint32, sBufOut *uintptr) int32
var _cMIOSampleBufferCreateNoDataMarkerErr error

func tryCMIOSampleBufferCreateNoDataMarker(allocator corefoundation.CFAllocatorRef, noDataEvent uint32, formatDescription uintptr, sequenceNumber uint64, discontinuityFlags uint32, sBufOut *uintptr) (int32, error) {
	if _cMIOSampleBufferCreateNoDataMarker == nil {
		return 0, symbolCallError("CMIOSampleBufferCreateNoDataMarker", "10.7", _cMIOSampleBufferCreateNoDataMarkerErr)
	}
	return _cMIOSampleBufferCreateNoDataMarker(allocator, noDataEvent, formatDescription, sequenceNumber, discontinuityFlags, sBufOut), nil
}

// CMIOSampleBufferCreateNoDataMarker.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferCreateNoDataMarker
func CMIOSampleBufferCreateNoDataMarker(allocator corefoundation.CFAllocatorRef, noDataEvent uint32, formatDescription uintptr, sequenceNumber uint64, discontinuityFlags uint32, sBufOut *uintptr) int32 {
	result, callErr := tryCMIOSampleBufferCreateNoDataMarker(allocator, noDataEvent, formatDescription, sequenceNumber, discontinuityFlags, sBufOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOSampleBufferGetDiscontinuityFlags func(sbuf uintptr) uint32
var _cMIOSampleBufferGetDiscontinuityFlagsErr error

func tryCMIOSampleBufferGetDiscontinuityFlags(sbuf uintptr) (uint32, error) {
	if _cMIOSampleBufferGetDiscontinuityFlags == nil {
		return 0, symbolCallError("CMIOSampleBufferGetDiscontinuityFlags", "10.7", _cMIOSampleBufferGetDiscontinuityFlagsErr)
	}
	return _cMIOSampleBufferGetDiscontinuityFlags(sbuf), nil
}

// CMIOSampleBufferGetDiscontinuityFlags.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferGetDiscontinuityFlags
func CMIOSampleBufferGetDiscontinuityFlags(sbuf uintptr) uint32 {
	result, callErr := tryCMIOSampleBufferGetDiscontinuityFlags(sbuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOSampleBufferGetSequenceNumber func(sbuf uintptr) uint64
var _cMIOSampleBufferGetSequenceNumberErr error

func tryCMIOSampleBufferGetSequenceNumber(sbuf uintptr) (uint64, error) {
	if _cMIOSampleBufferGetSequenceNumber == nil {
		return 0, symbolCallError("CMIOSampleBufferGetSequenceNumber", "10.7", _cMIOSampleBufferGetSequenceNumberErr)
	}
	return _cMIOSampleBufferGetSequenceNumber(sbuf), nil
}

// CMIOSampleBufferGetSequenceNumber.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferGetSequenceNumber
func CMIOSampleBufferGetSequenceNumber(sbuf uintptr) uint64 {
	result, callErr := tryCMIOSampleBufferGetSequenceNumber(sbuf)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOSampleBufferSetDiscontinuityFlags func(allocator corefoundation.CFAllocatorRef, sbuf uintptr, discontinuityFlags uint32)
var _cMIOSampleBufferSetDiscontinuityFlagsErr error

func tryCMIOSampleBufferSetDiscontinuityFlags(allocator corefoundation.CFAllocatorRef, sbuf uintptr, discontinuityFlags uint32) error {
	if _cMIOSampleBufferSetDiscontinuityFlags == nil {
		return symbolCallError("CMIOSampleBufferSetDiscontinuityFlags", "10.7", _cMIOSampleBufferSetDiscontinuityFlagsErr)
	}
	_cMIOSampleBufferSetDiscontinuityFlags(allocator, sbuf, discontinuityFlags)
	return nil
}

// CMIOSampleBufferSetDiscontinuityFlags.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferSetDiscontinuityFlags
func CMIOSampleBufferSetDiscontinuityFlags(allocator corefoundation.CFAllocatorRef, sbuf uintptr, discontinuityFlags uint32) {
	if callErr := tryCMIOSampleBufferSetDiscontinuityFlags(allocator, sbuf, discontinuityFlags); callErr != nil {
		panic(callErr)
	}
}

var _cMIOSampleBufferSetSequenceNumber func(allocator corefoundation.CFAllocatorRef, sbuf uintptr, sequenceNumber uint64)
var _cMIOSampleBufferSetSequenceNumberErr error

func tryCMIOSampleBufferSetSequenceNumber(allocator corefoundation.CFAllocatorRef, sbuf uintptr, sequenceNumber uint64) error {
	if _cMIOSampleBufferSetSequenceNumber == nil {
		return symbolCallError("CMIOSampleBufferSetSequenceNumber", "10.7", _cMIOSampleBufferSetSequenceNumberErr)
	}
	_cMIOSampleBufferSetSequenceNumber(allocator, sbuf, sequenceNumber)
	return nil
}

// CMIOSampleBufferSetSequenceNumber.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOSampleBufferSetSequenceNumber
func CMIOSampleBufferSetSequenceNumber(allocator corefoundation.CFAllocatorRef, sbuf uintptr, sequenceNumber uint64) {
	if callErr := tryCMIOSampleBufferSetSequenceNumber(allocator, sbuf, sequenceNumber); callErr != nil {
		panic(callErr)
	}
}

var _cMIOStreamClockConvertHostTimeToDeviceTime func(hostTime uint64, clock corefoundation.CFTypeRef) coremedia.CMTime
var _cMIOStreamClockConvertHostTimeToDeviceTimeErr error

func tryCMIOStreamClockConvertHostTimeToDeviceTime(hostTime uint64, clock corefoundation.CFTypeRef) (coremedia.CMTime, error) {
	if _cMIOStreamClockConvertHostTimeToDeviceTime == nil {
		return coremedia.CMTime{}, symbolCallError("CMIOStreamClockConvertHostTimeToDeviceTime", "10.7", _cMIOStreamClockConvertHostTimeToDeviceTimeErr)
	}
	return _cMIOStreamClockConvertHostTimeToDeviceTime(hostTime, clock), nil
}

// CMIOStreamClockConvertHostTimeToDeviceTime.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamClockConvertHostTimeToDeviceTime(_:_:)
func CMIOStreamClockConvertHostTimeToDeviceTime(hostTime uint64, clock corefoundation.CFTypeRef) coremedia.CMTime {
	result, callErr := tryCMIOStreamClockConvertHostTimeToDeviceTime(hostTime, clock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOStreamClockCreate func(allocator corefoundation.CFAllocatorRef, clockName corefoundation.CFStringRef, sourceIdentifier unsafe.Pointer, getTimeCallMinimumInterval coremedia.CMTime, numberOfEventsForRateSmoothing uint32, numberOfAveragesForRateSmoothing uint32, clock *corefoundation.CFTypeRef) int32
var _cMIOStreamClockCreateErr error

func tryCMIOStreamClockCreate(allocator corefoundation.CFAllocatorRef, clockName corefoundation.CFStringRef, sourceIdentifier unsafe.Pointer, getTimeCallMinimumInterval coremedia.CMTime, numberOfEventsForRateSmoothing uint32, numberOfAveragesForRateSmoothing uint32, clock *corefoundation.CFTypeRef) (int32, error) {
	if _cMIOStreamClockCreate == nil {
		return 0, symbolCallError("CMIOStreamClockCreate", "10.7", _cMIOStreamClockCreateErr)
	}
	return _cMIOStreamClockCreate(allocator, clockName, sourceIdentifier, getTimeCallMinimumInterval, numberOfEventsForRateSmoothing, numberOfAveragesForRateSmoothing, clock), nil
}

// CMIOStreamClockCreate.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamClockCreate(_:_:_:_:_:_:_:)
func CMIOStreamClockCreate(allocator corefoundation.CFAllocatorRef, clockName corefoundation.CFStringRef, sourceIdentifier unsafe.Pointer, getTimeCallMinimumInterval coremedia.CMTime, numberOfEventsForRateSmoothing uint32, numberOfAveragesForRateSmoothing uint32, clock *corefoundation.CFTypeRef) int32 {
	result, callErr := tryCMIOStreamClockCreate(allocator, clockName, sourceIdentifier, getTimeCallMinimumInterval, numberOfEventsForRateSmoothing, numberOfAveragesForRateSmoothing, clock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOStreamClockInvalidate func(clock corefoundation.CFTypeRef) int32
var _cMIOStreamClockInvalidateErr error

func tryCMIOStreamClockInvalidate(clock corefoundation.CFTypeRef) (int32, error) {
	if _cMIOStreamClockInvalidate == nil {
		return 0, symbolCallError("CMIOStreamClockInvalidate", "10.7", _cMIOStreamClockInvalidateErr)
	}
	return _cMIOStreamClockInvalidate(clock), nil
}

// CMIOStreamClockInvalidate.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamClockInvalidate(_:)
func CMIOStreamClockInvalidate(clock corefoundation.CFTypeRef) int32 {
	result, callErr := tryCMIOStreamClockInvalidate(clock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOStreamClockPostTimingEvent func(eventTime coremedia.CMTime, hostTime uint64, resynchronize bool, clock corefoundation.CFTypeRef) int32
var _cMIOStreamClockPostTimingEventErr error

func tryCMIOStreamClockPostTimingEvent(eventTime coremedia.CMTime, hostTime uint64, resynchronize bool, clock corefoundation.CFTypeRef) (int32, error) {
	if _cMIOStreamClockPostTimingEvent == nil {
		return 0, symbolCallError("CMIOStreamClockPostTimingEvent", "10.7", _cMIOStreamClockPostTimingEventErr)
	}
	return _cMIOStreamClockPostTimingEvent(eventTime, hostTime, resynchronize, clock), nil
}

// CMIOStreamClockPostTimingEvent.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamClockPostTimingEvent(_:_:_:_:)
func CMIOStreamClockPostTimingEvent(eventTime coremedia.CMTime, hostTime uint64, resynchronize bool, clock corefoundation.CFTypeRef) int32 {
	result, callErr := tryCMIOStreamClockPostTimingEvent(eventTime, hostTime, resynchronize, clock)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOStreamCopyBufferQueue func(streamID CMIOStreamID, queueAlteredProc CMIODeviceStreamQueueAlteredProc, queueAlteredRefCon unsafe.Pointer, queue *coremedia.CMSimpleQueueRef) int32
var _cMIOStreamCopyBufferQueueErr error

func tryCMIOStreamCopyBufferQueue(streamID CMIOStreamID, queueAlteredProc CMIODeviceStreamQueueAlteredProc, queueAlteredRefCon unsafe.Pointer, queue *coremedia.CMSimpleQueueRef) (int32, error) {
	if _cMIOStreamCopyBufferQueue == nil {
		return 0, symbolCallError("CMIOStreamCopyBufferQueue", "10.7", _cMIOStreamCopyBufferQueueErr)
	}
	return _cMIOStreamCopyBufferQueue(streamID, queueAlteredProc, queueAlteredRefCon, queue), nil
}

// CMIOStreamCopyBufferQueue.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamCopyBufferQueue(_:_:_:_:)
func CMIOStreamCopyBufferQueue(streamID CMIOStreamID, queueAlteredProc CMIODeviceStreamQueueAlteredProc, queueAlteredRefCon unsafe.Pointer, queue *coremedia.CMSimpleQueueRef) int32 {
	result, callErr := tryCMIOStreamCopyBufferQueue(streamID, queueAlteredProc, queueAlteredRefCon, queue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOStreamDeckCueTo func(streamID CMIOStreamID, frameNumber uint64, playOnCue bool) int32
var _cMIOStreamDeckCueToErr error

func tryCMIOStreamDeckCueTo(streamID CMIOStreamID, frameNumber uint64, playOnCue bool) (int32, error) {
	if _cMIOStreamDeckCueTo == nil {
		return 0, symbolCallError("CMIOStreamDeckCueTo", "10.7", _cMIOStreamDeckCueToErr)
	}
	return _cMIOStreamDeckCueTo(streamID, frameNumber, playOnCue), nil
}

// CMIOStreamDeckCueTo.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamDeckCueTo(_:_:_:)
func CMIOStreamDeckCueTo(streamID CMIOStreamID, frameNumber uint64, playOnCue bool) int32 {
	result, callErr := tryCMIOStreamDeckCueTo(streamID, frameNumber, playOnCue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOStreamDeckJog func(streamID CMIOStreamID, speed int32) int32
var _cMIOStreamDeckJogErr error

func tryCMIOStreamDeckJog(streamID CMIOStreamID, speed int32) (int32, error) {
	if _cMIOStreamDeckJog == nil {
		return 0, symbolCallError("CMIOStreamDeckJog", "10.7", _cMIOStreamDeckJogErr)
	}
	return _cMIOStreamDeckJog(streamID, speed), nil
}

// CMIOStreamDeckJog.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamDeckJog(_:_:)
func CMIOStreamDeckJog(streamID CMIOStreamID, speed int32) int32 {
	result, callErr := tryCMIOStreamDeckJog(streamID, speed)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOStreamDeckPlay func(streamID CMIOStreamID) int32
var _cMIOStreamDeckPlayErr error

func tryCMIOStreamDeckPlay(streamID CMIOStreamID) (int32, error) {
	if _cMIOStreamDeckPlay == nil {
		return 0, symbolCallError("CMIOStreamDeckPlay", "10.7", _cMIOStreamDeckPlayErr)
	}
	return _cMIOStreamDeckPlay(streamID), nil
}

// CMIOStreamDeckPlay.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamDeckPlay(_:)
func CMIOStreamDeckPlay(streamID CMIOStreamID) int32 {
	result, callErr := tryCMIOStreamDeckPlay(streamID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMIOStreamDeckStop func(streamID CMIOStreamID) int32
var _cMIOStreamDeckStopErr error

func tryCMIOStreamDeckStop(streamID CMIOStreamID) (int32, error) {
	if _cMIOStreamDeckStop == nil {
		return 0, symbolCallError("CMIOStreamDeckStop", "10.7", _cMIOStreamDeckStopErr)
	}
	return _cMIOStreamDeckStop(streamID), nil
}

// CMIOStreamDeckStop.
//
// See: https://developer.apple.com/documentation/CoreMediaIO/CMIOStreamDeckStop(_:)
func CMIOStreamDeckStop(streamID CMIOStreamID) int32 {
	result, callErr := tryCMIOStreamDeckStop(streamID)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_cMIODeviceProcessAVCCommand, &_cMIODeviceProcessAVCCommandErr, frameworkHandle, "CMIODeviceProcessAVCCommand", "10.7")
	registerFunc(&_cMIODeviceProcessRS422Command, &_cMIODeviceProcessRS422CommandErr, frameworkHandle, "CMIODeviceProcessRS422Command", "10.7")
	registerFunc(&_cMIODeviceStartStream, &_cMIODeviceStartStreamErr, frameworkHandle, "CMIODeviceStartStream", "10.7")
	registerFunc(&_cMIODeviceStopStream, &_cMIODeviceStopStreamErr, frameworkHandle, "CMIODeviceStopStream", "10.7")
	registerFunc(&_cMIOObjectAddPropertyListener, &_cMIOObjectAddPropertyListenerErr, frameworkHandle, "CMIOObjectAddPropertyListener", "10.7")
	registerFunc(&_cMIOObjectAddPropertyListenerBlock, &_cMIOObjectAddPropertyListenerBlockErr, frameworkHandle, "CMIOObjectAddPropertyListenerBlock", "10.8")
	registerFunc(&_cMIOObjectCreate, &_cMIOObjectCreateErr, frameworkHandle, "CMIOObjectCreate", "10.7")
	registerFunc(&_cMIOObjectGetPropertyData, &_cMIOObjectGetPropertyDataErr, frameworkHandle, "CMIOObjectGetPropertyData", "10.7")
	registerFunc(&_cMIOObjectGetPropertyDataSize, &_cMIOObjectGetPropertyDataSizeErr, frameworkHandle, "CMIOObjectGetPropertyDataSize", "10.7")
	registerFunc(&_cMIOObjectHasProperty, &_cMIOObjectHasPropertyErr, frameworkHandle, "CMIOObjectHasProperty", "10.7")
	registerFunc(&_cMIOObjectIsPropertySettable, &_cMIOObjectIsPropertySettableErr, frameworkHandle, "CMIOObjectIsPropertySettable", "10.7")
	registerFunc(&_cMIOObjectPropertiesChanged, &_cMIOObjectPropertiesChangedErr, frameworkHandle, "CMIOObjectPropertiesChanged", "10.7")
	registerFunc(&_cMIOObjectRemovePropertyListener, &_cMIOObjectRemovePropertyListenerErr, frameworkHandle, "CMIOObjectRemovePropertyListener", "10.7")
	registerFunc(&_cMIOObjectRemovePropertyListenerBlock, &_cMIOObjectRemovePropertyListenerBlockErr, frameworkHandle, "CMIOObjectRemovePropertyListenerBlock", "10.8")
	registerFunc(&_cMIOObjectSetPropertyData, &_cMIOObjectSetPropertyDataErr, frameworkHandle, "CMIOObjectSetPropertyData", "10.7")
	registerFunc(&_cMIOObjectShow, &_cMIOObjectShowErr, frameworkHandle, "CMIOObjectShow", "10.7")
	registerFunc(&_cMIOObjectsPublishedAndDied, &_cMIOObjectsPublishedAndDiedErr, frameworkHandle, "CMIOObjectsPublishedAndDied", "10.7")
	registerFunc(&_cMIOSampleBufferCopyNonRequiredAttachments, &_cMIOSampleBufferCopyNonRequiredAttachmentsErr, frameworkHandle, "CMIOSampleBufferCopyNonRequiredAttachments", "10.7")
	registerFunc(&_cMIOSampleBufferCopySampleAttachments, &_cMIOSampleBufferCopySampleAttachmentsErr, frameworkHandle, "CMIOSampleBufferCopySampleAttachments", "10.7")
	registerFunc(&_cMIOSampleBufferCreate, &_cMIOSampleBufferCreateErr, frameworkHandle, "CMIOSampleBufferCreate", "10.7")
	registerFunc(&_cMIOSampleBufferCreateForImageBuffer, &_cMIOSampleBufferCreateForImageBufferErr, frameworkHandle, "CMIOSampleBufferCreateForImageBuffer", "10.7")
	registerFunc(&_cMIOSampleBufferCreateNoDataMarker, &_cMIOSampleBufferCreateNoDataMarkerErr, frameworkHandle, "CMIOSampleBufferCreateNoDataMarker", "10.7")
	registerFunc(&_cMIOSampleBufferGetDiscontinuityFlags, &_cMIOSampleBufferGetDiscontinuityFlagsErr, frameworkHandle, "CMIOSampleBufferGetDiscontinuityFlags", "10.7")
	registerFunc(&_cMIOSampleBufferGetSequenceNumber, &_cMIOSampleBufferGetSequenceNumberErr, frameworkHandle, "CMIOSampleBufferGetSequenceNumber", "10.7")
	registerFunc(&_cMIOSampleBufferSetDiscontinuityFlags, &_cMIOSampleBufferSetDiscontinuityFlagsErr, frameworkHandle, "CMIOSampleBufferSetDiscontinuityFlags", "10.7")
	registerFunc(&_cMIOSampleBufferSetSequenceNumber, &_cMIOSampleBufferSetSequenceNumberErr, frameworkHandle, "CMIOSampleBufferSetSequenceNumber", "10.7")
	registerFunc(&_cMIOStreamClockConvertHostTimeToDeviceTime, &_cMIOStreamClockConvertHostTimeToDeviceTimeErr, frameworkHandle, "CMIOStreamClockConvertHostTimeToDeviceTime", "10.7")
	registerFunc(&_cMIOStreamClockCreate, &_cMIOStreamClockCreateErr, frameworkHandle, "CMIOStreamClockCreate", "10.7")
	registerFunc(&_cMIOStreamClockInvalidate, &_cMIOStreamClockInvalidateErr, frameworkHandle, "CMIOStreamClockInvalidate", "10.7")
	registerFunc(&_cMIOStreamClockPostTimingEvent, &_cMIOStreamClockPostTimingEventErr, frameworkHandle, "CMIOStreamClockPostTimingEvent", "10.7")
	registerFunc(&_cMIOStreamCopyBufferQueue, &_cMIOStreamCopyBufferQueueErr, frameworkHandle, "CMIOStreamCopyBufferQueue", "10.7")
	registerFunc(&_cMIOStreamDeckCueTo, &_cMIOStreamDeckCueToErr, frameworkHandle, "CMIOStreamDeckCueTo", "10.7")
	registerFunc(&_cMIOStreamDeckJog, &_cMIOStreamDeckJogErr, frameworkHandle, "CMIOStreamDeckJog", "10.7")
	registerFunc(&_cMIOStreamDeckPlay, &_cMIOStreamDeckPlayErr, frameworkHandle, "CMIOStreamDeckPlay", "10.7")
	registerFunc(&_cMIOStreamDeckStop, &_cMIOStreamDeckStopErr, frameworkHandle, "CMIOStreamDeckStop", "10.7")
}
