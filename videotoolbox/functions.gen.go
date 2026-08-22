// Code generated from Apple documentation for VideoToolbox. DO NOT EDIT.

package videotoolbox

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/corevideo"
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
		return fmt.Sprintf("VideoToolbox: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("VideoToolbox: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("VideoToolbox: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("VideoToolbox: register symbol %s: %v", name, r)
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

var _vTCompressionSessionBeginPass func(session VTCompressionSessionRef, beginPassFlags VTCompressionSessionOptionFlags, reserved *uint32) int32
var _vTCompressionSessionBeginPassErr error

func tryVTCompressionSessionBeginPass(session VTCompressionSessionRef, beginPassFlags VTCompressionSessionOptionFlags, reserved *uint32) (int32, error) {
	if _vTCompressionSessionBeginPass == nil {
		return 0, symbolCallError("VTCompressionSessionBeginPass", "10.10", _vTCompressionSessionBeginPassErr)
	}
	return _vTCompressionSessionBeginPass(session, beginPassFlags, reserved), nil
}

// VTCompressionSessionBeginPass marks the start of a specific compression pass.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCompressionSessionBeginPass(_:flags:_:)
func VTCompressionSessionBeginPass(session VTCompressionSessionRef, beginPassFlags VTCompressionSessionOptionFlags, reserved *uint32) int32 {
	result, callErr := tryVTCompressionSessionBeginPass(session, beginPassFlags, reserved)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTCompressionSessionCompleteFrames func(session VTCompressionSessionRef, completeUntilPresentationTimeStamp coremedia.CMTime) int32
var _vTCompressionSessionCompleteFramesErr error

func tryVTCompressionSessionCompleteFrames(session VTCompressionSessionRef, completeUntilPresentationTimeStamp coremedia.CMTime) (int32, error) {
	if _vTCompressionSessionCompleteFrames == nil {
		return 0, symbolCallError("VTCompressionSessionCompleteFrames", "10.8", _vTCompressionSessionCompleteFramesErr)
	}
	return _vTCompressionSessionCompleteFrames(session, completeUntilPresentationTimeStamp), nil
}

// VTCompressionSessionCompleteFrames forces the compression session to complete the encoding of frames.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCompressionSessionCompleteFrames(_:untilPresentationTimeStamp:)
func VTCompressionSessionCompleteFrames(session VTCompressionSessionRef, completeUntilPresentationTimeStamp coremedia.CMTime) int32 {
	result, callErr := tryVTCompressionSessionCompleteFrames(session, completeUntilPresentationTimeStamp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTCompressionSessionCreate func(allocator corefoundation.CFAllocatorRef, width int32, height int32, codecType uint32, encoderSpecification corefoundation.CFDictionaryRef, sourceImageBufferAttributes corefoundation.CFDictionaryRef, compressedDataAllocator corefoundation.CFAllocatorRef, outputCallback VTCompressionOutputCallback, outputCallbackRefCon unsafe.Pointer, compressionSessionOut *VTCompressionSessionRef) int32
var _vTCompressionSessionCreateErr error

func tryVTCompressionSessionCreate(allocator corefoundation.CFAllocatorRef, width int32, height int32, codecType uint32, encoderSpecification corefoundation.CFDictionaryRef, sourceImageBufferAttributes corefoundation.CFDictionaryRef, compressedDataAllocator corefoundation.CFAllocatorRef, outputCallback VTCompressionOutputCallback, outputCallbackRefCon unsafe.Pointer, compressionSessionOut *VTCompressionSessionRef) (int32, error) {
	if _vTCompressionSessionCreate == nil {
		return 0, symbolCallError("VTCompressionSessionCreate", "10.8", _vTCompressionSessionCreateErr)
	}
	return _vTCompressionSessionCreate(allocator, width, height, codecType, encoderSpecification, sourceImageBufferAttributes, compressedDataAllocator, outputCallback, outputCallbackRefCon, compressionSessionOut), nil
}

// VTCompressionSessionCreate creates an object that compresses video frames.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCompressionSessionCreate(allocator:width:height:codecType:encoderSpecification:imageBufferAttributes:compressedDataAllocator:outputCallback:refcon:compressionSessionOut:)
func VTCompressionSessionCreate(allocator corefoundation.CFAllocatorRef, width int32, height int32, codecType uint32, encoderSpecification corefoundation.CFDictionaryRef, sourceImageBufferAttributes corefoundation.CFDictionaryRef, compressedDataAllocator corefoundation.CFAllocatorRef, outputCallback VTCompressionOutputCallback, outputCallbackRefCon unsafe.Pointer, compressionSessionOut *VTCompressionSessionRef) int32 {
	result, callErr := tryVTCompressionSessionCreate(allocator, width, height, codecType, encoderSpecification, sourceImageBufferAttributes, compressedDataAllocator, outputCallback, outputCallbackRefCon, compressionSessionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTCompressionSessionEncodeFrame func(session VTCompressionSessionRef, imageBuffer corevideo.CVImageBufferRef, presentationTimeStamp coremedia.CMTime, duration coremedia.CMTime, frameProperties corefoundation.CFDictionaryRef, sourceFrameRefcon unsafe.Pointer, infoFlagsOut *VTEncodeInfoFlags) int32
var _vTCompressionSessionEncodeFrameErr error

func tryVTCompressionSessionEncodeFrame(session VTCompressionSessionRef, imageBuffer corevideo.CVImageBufferRef, presentationTimeStamp coremedia.CMTime, duration coremedia.CMTime, frameProperties corefoundation.CFDictionaryRef, sourceFrameRefcon unsafe.Pointer, infoFlagsOut *VTEncodeInfoFlags) (int32, error) {
	if _vTCompressionSessionEncodeFrame == nil {
		return 0, symbolCallError("VTCompressionSessionEncodeFrame", "10.8", _vTCompressionSessionEncodeFrameErr)
	}
	return _vTCompressionSessionEncodeFrame(session, imageBuffer, presentationTimeStamp, duration, frameProperties, sourceFrameRefcon, infoFlagsOut), nil
}

// VTCompressionSessionEncodeFrame presents frames to the compression session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCompressionSessionEncodeFrame(_:imageBuffer:presentationTimeStamp:duration:frameProperties:sourceFrameRefcon:infoFlagsOut:)
func VTCompressionSessionEncodeFrame(session VTCompressionSessionRef, imageBuffer corevideo.CVImageBufferRef, presentationTimeStamp coremedia.CMTime, duration coremedia.CMTime, frameProperties corefoundation.CFDictionaryRef, sourceFrameRefcon unsafe.Pointer, infoFlagsOut *VTEncodeInfoFlags) int32 {
	result, callErr := tryVTCompressionSessionEncodeFrame(session, imageBuffer, presentationTimeStamp, duration, frameProperties, sourceFrameRefcon, infoFlagsOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTCompressionSessionEncodeFrameWithOutputHandler func(session VTCompressionSessionRef, imageBuffer corevideo.CVImageBufferRef, presentationTimeStamp coremedia.CMTime, duration coremedia.CMTime, frameProperties corefoundation.CFDictionaryRef, infoFlagsOut *VTEncodeInfoFlags, outputHandler unsafe.Pointer) int32
var _vTCompressionSessionEncodeFrameWithOutputHandlerErr error

func tryVTCompressionSessionEncodeFrameWithOutputHandler(session VTCompressionSessionRef, imageBuffer corevideo.CVImageBufferRef, presentationTimeStamp coremedia.CMTime, duration coremedia.CMTime, frameProperties corefoundation.CFDictionaryRef, infoFlagsOut *VTEncodeInfoFlags, outputHandler VTCompressionOutputHandler) (int32, error) {
	if _vTCompressionSessionEncodeFrameWithOutputHandler == nil {
		return 0, symbolCallError("VTCompressionSessionEncodeFrameWithOutputHandler", "10.11", _vTCompressionSessionEncodeFrameWithOutputHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 int32, blockArg1 uint, blockArg2 unsafe.Pointer) {
		outputHandler(blockArg0, blockArg1, blockArg2)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _vTCompressionSessionEncodeFrameWithOutputHandler(session, imageBuffer, presentationTimeStamp, duration, frameProperties, infoFlagsOut, _block0), nil
}

// VTCompressionSessionEncodeFrameWithOutputHandler presents frames to the compression session and invokes the output callback when compression is complete.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCompressionSessionEncodeFrame(_:imageBuffer:presentationTimeStamp:duration:frameProperties:infoFlagsOut:outputHandler:)
func VTCompressionSessionEncodeFrameWithOutputHandler(session VTCompressionSessionRef, imageBuffer corevideo.CVImageBufferRef, presentationTimeStamp coremedia.CMTime, duration coremedia.CMTime, frameProperties corefoundation.CFDictionaryRef, infoFlagsOut *VTEncodeInfoFlags, outputHandler VTCompressionOutputHandler) int32 {
	result, callErr := tryVTCompressionSessionEncodeFrameWithOutputHandler(session, imageBuffer, presentationTimeStamp, duration, frameProperties, infoFlagsOut, outputHandler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTCompressionSessionEncodeMultiImageFrame func(session VTCompressionSessionRef, taggedBufferGroup coremedia.CMTaggedBufferGroupRef, presentationTimeStamp coremedia.CMTime, duration coremedia.CMTime, frameProperties corefoundation.CFDictionaryRef, sourceFrameRefcon unsafe.Pointer, infoFlagsOut *VTEncodeInfoFlags) int32
var _vTCompressionSessionEncodeMultiImageFrameErr error

func tryVTCompressionSessionEncodeMultiImageFrame(session VTCompressionSessionRef, taggedBufferGroup coremedia.CMTaggedBufferGroupRef, presentationTimeStamp coremedia.CMTime, duration coremedia.CMTime, frameProperties corefoundation.CFDictionaryRef, sourceFrameRefcon unsafe.Pointer, infoFlagsOut *VTEncodeInfoFlags) (int32, error) {
	if _vTCompressionSessionEncodeMultiImageFrame == nil {
		return 0, symbolCallError("VTCompressionSessionEncodeMultiImageFrame", "14.0", _vTCompressionSessionEncodeMultiImageFrameErr)
	}
	return _vTCompressionSessionEncodeMultiImageFrame(session, taggedBufferGroup, presentationTimeStamp, duration, frameProperties, sourceFrameRefcon, infoFlagsOut), nil
}

// VTCompressionSessionEncodeMultiImageFrame passes a multi-image frame to a compression session for encoding.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCompressionSessionEncodeMultiImageFrame
func VTCompressionSessionEncodeMultiImageFrame(session VTCompressionSessionRef, taggedBufferGroup coremedia.CMTaggedBufferGroupRef, presentationTimeStamp coremedia.CMTime, duration coremedia.CMTime, frameProperties corefoundation.CFDictionaryRef, sourceFrameRefcon unsafe.Pointer, infoFlagsOut *VTEncodeInfoFlags) int32 {
	result, callErr := tryVTCompressionSessionEncodeMultiImageFrame(session, taggedBufferGroup, presentationTimeStamp, duration, frameProperties, sourceFrameRefcon, infoFlagsOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTCompressionSessionEncodeMultiImageFrameWithOutputHandler func(session VTCompressionSessionRef, taggedBufferGroup coremedia.CMTaggedBufferGroupRef, presentationTimeStamp coremedia.CMTime, duration coremedia.CMTime, frameProperties corefoundation.CFDictionaryRef, infoFlagsOut *VTEncodeInfoFlags, outputHandler unsafe.Pointer) int32
var _vTCompressionSessionEncodeMultiImageFrameWithOutputHandlerErr error

func tryVTCompressionSessionEncodeMultiImageFrameWithOutputHandler(session VTCompressionSessionRef, taggedBufferGroup coremedia.CMTaggedBufferGroupRef, presentationTimeStamp coremedia.CMTime, duration coremedia.CMTime, frameProperties corefoundation.CFDictionaryRef, infoFlagsOut *VTEncodeInfoFlags, outputHandler VTCompressionOutputHandler) (int32, error) {
	if _vTCompressionSessionEncodeMultiImageFrameWithOutputHandler == nil {
		return 0, symbolCallError("VTCompressionSessionEncodeMultiImageFrameWithOutputHandler", "14.0", _vTCompressionSessionEncodeMultiImageFrameWithOutputHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 int32, blockArg1 uint, blockArg2 unsafe.Pointer) {
		outputHandler(blockArg0, blockArg1, blockArg2)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _vTCompressionSessionEncodeMultiImageFrameWithOutputHandler(session, taggedBufferGroup, presentationTimeStamp, duration, frameProperties, infoFlagsOut, _block0), nil
}

// VTCompressionSessionEncodeMultiImageFrameWithOutputHandler passes a multi-image frame to a compression session for encoding and provides a callback to handle the output.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCompressionSessionEncodeMultiImageFrameWithOutputHandler
func VTCompressionSessionEncodeMultiImageFrameWithOutputHandler(session VTCompressionSessionRef, taggedBufferGroup coremedia.CMTaggedBufferGroupRef, presentationTimeStamp coremedia.CMTime, duration coremedia.CMTime, frameProperties corefoundation.CFDictionaryRef, infoFlagsOut *VTEncodeInfoFlags, outputHandler VTCompressionOutputHandler) int32 {
	result, callErr := tryVTCompressionSessionEncodeMultiImageFrameWithOutputHandler(session, taggedBufferGroup, presentationTimeStamp, duration, frameProperties, infoFlagsOut, outputHandler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTCompressionSessionEndPass func(session VTCompressionSessionRef, furtherPassesRequestedOut *bool, reserved *uint32) int32
var _vTCompressionSessionEndPassErr error

func tryVTCompressionSessionEndPass(session VTCompressionSessionRef, furtherPassesRequestedOut *bool, reserved *uint32) (int32, error) {
	if _vTCompressionSessionEndPass == nil {
		return 0, symbolCallError("VTCompressionSessionEndPass", "10.10", _vTCompressionSessionEndPassErr)
	}
	return _vTCompressionSessionEndPass(session, furtherPassesRequestedOut, reserved), nil
}

// VTCompressionSessionEndPass marks the end of a compression pass.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCompressionSessionEndPass(_:furtherPassesRequestedOut:_:)
func VTCompressionSessionEndPass(session VTCompressionSessionRef, furtherPassesRequestedOut *bool, reserved *uint32) int32 {
	result, callErr := tryVTCompressionSessionEndPass(session, furtherPassesRequestedOut, reserved)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTCompressionSessionGetPixelBufferPool func(session VTCompressionSessionRef) corevideo.CVPixelBufferPoolRef
var _vTCompressionSessionGetPixelBufferPoolErr error

func tryVTCompressionSessionGetPixelBufferPool(session VTCompressionSessionRef) (corevideo.CVPixelBufferPoolRef, error) {
	if _vTCompressionSessionGetPixelBufferPool == nil {
		return *new(corevideo.CVPixelBufferPoolRef), symbolCallError("VTCompressionSessionGetPixelBufferPool", "10.8", _vTCompressionSessionGetPixelBufferPoolErr)
	}
	return _vTCompressionSessionGetPixelBufferPool(session), nil
}

// VTCompressionSessionGetPixelBufferPool returns a pool that provides ideal source pixel buffers for a compression session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCompressionSessionGetPixelBufferPool(_:)
func VTCompressionSessionGetPixelBufferPool(session VTCompressionSessionRef) corevideo.CVPixelBufferPoolRef {
	result, callErr := tryVTCompressionSessionGetPixelBufferPool(session)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTCompressionSessionGetTimeRangesForNextPass func(session VTCompressionSessionRef, timeRangeCountOut *int, timeRangeArrayOut *objc.ID) int32
var _vTCompressionSessionGetTimeRangesForNextPassErr error

func tryVTCompressionSessionGetTimeRangesForNextPass(session VTCompressionSessionRef, timeRangeCountOut *int, timeRangeArrayOut *objc.ID) (int32, error) {
	if _vTCompressionSessionGetTimeRangesForNextPass == nil {
		return 0, symbolCallError("VTCompressionSessionGetTimeRangesForNextPass", "10.10", _vTCompressionSessionGetTimeRangesForNextPassErr)
	}
	return _vTCompressionSessionGetTimeRangesForNextPass(session, timeRangeCountOut, timeRangeArrayOut), nil
}

// VTCompressionSessionGetTimeRangesForNextPass retrieves the time ranges for the next pass.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCompressionSessionGetTimeRangesForNextPass(_:timeRangeCountOut:timeRangeArrayOut:)
func VTCompressionSessionGetTimeRangesForNextPass(session VTCompressionSessionRef, timeRangeCountOut *int, timeRangeArrayOut *objc.ID) int32 {
	result, callErr := tryVTCompressionSessionGetTimeRangesForNextPass(session, timeRangeCountOut, timeRangeArrayOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTCompressionSessionGetTypeID func() uint
var _vTCompressionSessionGetTypeIDErr error

func tryVTCompressionSessionGetTypeID() (uint, error) {
	if _vTCompressionSessionGetTypeID == nil {
		return 0, symbolCallError("VTCompressionSessionGetTypeID", "10.8", _vTCompressionSessionGetTypeIDErr)
	}
	return _vTCompressionSessionGetTypeID(), nil
}

// VTCompressionSessionGetTypeID retrieves the Core Foundation type identifier for the compression session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCompressionSessionGetTypeID()
func VTCompressionSessionGetTypeID() uint {
	result, callErr := tryVTCompressionSessionGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTCompressionSessionInvalidate func(session VTCompressionSessionRef)
var _vTCompressionSessionInvalidateErr error

func tryVTCompressionSessionInvalidate(session VTCompressionSessionRef) error {
	if _vTCompressionSessionInvalidate == nil {
		return symbolCallError("VTCompressionSessionInvalidate", "10.8", _vTCompressionSessionInvalidateErr)
	}
	_vTCompressionSessionInvalidate(session)
	return nil
}

// VTCompressionSessionInvalidate tears down a compression session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCompressionSessionInvalidate(_:)
func VTCompressionSessionInvalidate(session VTCompressionSessionRef) {
	if callErr := tryVTCompressionSessionInvalidate(session); callErr != nil {
		panic(callErr)
	}
}

var _vTCompressionSessionPrepareToEncodeFrames func(session VTCompressionSessionRef) int32
var _vTCompressionSessionPrepareToEncodeFramesErr error

func tryVTCompressionSessionPrepareToEncodeFrames(session VTCompressionSessionRef) (int32, error) {
	if _vTCompressionSessionPrepareToEncodeFrames == nil {
		return 0, symbolCallError("VTCompressionSessionPrepareToEncodeFrames", "10.9", _vTCompressionSessionPrepareToEncodeFramesErr)
	}
	return _vTCompressionSessionPrepareToEncodeFrames(session), nil
}

// VTCompressionSessionPrepareToEncodeFrames enables the encoder to perform any necessary resource allocation before the encoder begins encoding frames (optional).
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCompressionSessionPrepareToEncodeFrames(_:)
func VTCompressionSessionPrepareToEncodeFrames(session VTCompressionSessionRef) int32 {
	result, callErr := tryVTCompressionSessionPrepareToEncodeFrames(session)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTCopyRAWProcessorExtensionProperties func(formatDesc uintptr, mediaExtensionPropertiesOut *corefoundation.CFDictionaryRef) int32
var _vTCopyRAWProcessorExtensionPropertiesErr error

func tryVTCopyRAWProcessorExtensionProperties(formatDesc uintptr, mediaExtensionPropertiesOut *corefoundation.CFDictionaryRef) (int32, error) {
	if _vTCopyRAWProcessorExtensionProperties == nil {
		return 0, symbolCallError("VTCopyRAWProcessorExtensionProperties", "15.0", _vTCopyRAWProcessorExtensionPropertiesErr)
	}
	return _vTCopyRAWProcessorExtensionProperties(formatDesc, mediaExtensionPropertiesOut), nil
}

// VTCopyRAWProcessorExtensionProperties returns information about the Media Extension RAW processor supporting the specified format.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCopyRAWProcessorExtensionProperties
func VTCopyRAWProcessorExtensionProperties(formatDesc uintptr, mediaExtensionPropertiesOut *corefoundation.CFDictionaryRef) int32 {
	result, callErr := tryVTCopyRAWProcessorExtensionProperties(formatDesc, mediaExtensionPropertiesOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTCopySupportedPropertyDictionaryForEncoder func(width int32, height int32, codecType uint32, encoderSpecification corefoundation.CFDictionaryRef, encoderIDOut *corefoundation.CFStringRef, supportedPropertiesOut *corefoundation.CFDictionaryRef) int32
var _vTCopySupportedPropertyDictionaryForEncoderErr error

func tryVTCopySupportedPropertyDictionaryForEncoder(width int32, height int32, codecType uint32, encoderSpecification corefoundation.CFDictionaryRef, encoderIDOut *corefoundation.CFStringRef, supportedPropertiesOut *corefoundation.CFDictionaryRef) (int32, error) {
	if _vTCopySupportedPropertyDictionaryForEncoder == nil {
		return 0, symbolCallError("VTCopySupportedPropertyDictionaryForEncoder", "10.13", _vTCopySupportedPropertyDictionaryForEncoderErr)
	}
	return _vTCopySupportedPropertyDictionaryForEncoder(width, height, codecType, encoderSpecification, encoderIDOut, supportedPropertiesOut), nil
}

// VTCopySupportedPropertyDictionaryForEncoder builds a list of supported properties and encoder ID for an encoder.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCopySupportedPropertyDictionaryForEncoder(width:height:codecType:encoderSpecification:encoderIDOut:supportedPropertiesOut:)
func VTCopySupportedPropertyDictionaryForEncoder(width int32, height int32, codecType uint32, encoderSpecification corefoundation.CFDictionaryRef, encoderIDOut *corefoundation.CFStringRef, supportedPropertiesOut *corefoundation.CFDictionaryRef) int32 {
	result, callErr := tryVTCopySupportedPropertyDictionaryForEncoder(width, height, codecType, encoderSpecification, encoderIDOut, supportedPropertiesOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTCopyVideoDecoderExtensionProperties func(formatDesc uintptr, mediaExtensionPropertiesOut *corefoundation.CFDictionaryRef) int32
var _vTCopyVideoDecoderExtensionPropertiesErr error

func tryVTCopyVideoDecoderExtensionProperties(formatDesc uintptr, mediaExtensionPropertiesOut *corefoundation.CFDictionaryRef) (int32, error) {
	if _vTCopyVideoDecoderExtensionProperties == nil {
		return 0, symbolCallError("VTCopyVideoDecoderExtensionProperties", "15.0", _vTCopyVideoDecoderExtensionPropertiesErr)
	}
	return _vTCopyVideoDecoderExtensionProperties(formatDesc, mediaExtensionPropertiesOut), nil
}

// VTCopyVideoDecoderExtensionProperties returns information about the Media Extension video decoder required to decode the specified format.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCopyVideoDecoderExtensionProperties
func VTCopyVideoDecoderExtensionProperties(formatDesc uintptr, mediaExtensionPropertiesOut *corefoundation.CFDictionaryRef) int32 {
	result, callErr := tryVTCopyVideoDecoderExtensionProperties(formatDesc, mediaExtensionPropertiesOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTCopyVideoEncoderList func(options corefoundation.CFDictionaryRef, listOfVideoEncodersOut *corefoundation.CFArrayRef) int32
var _vTCopyVideoEncoderListErr error

func tryVTCopyVideoEncoderList(options corefoundation.CFDictionaryRef, listOfVideoEncodersOut *corefoundation.CFArrayRef) (int32, error) {
	if _vTCopyVideoEncoderList == nil {
		return 0, symbolCallError("VTCopyVideoEncoderList", "10.8", _vTCopyVideoEncoderListErr)
	}
	return _vTCopyVideoEncoderList(options, listOfVideoEncodersOut), nil
}

// VTCopyVideoEncoderList builds a list of available video encoders.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCopyVideoEncoderList(_:_:)
func VTCopyVideoEncoderList(options corefoundation.CFDictionaryRef, listOfVideoEncodersOut *corefoundation.CFArrayRef) int32 {
	result, callErr := tryVTCopyVideoEncoderList(options, listOfVideoEncodersOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTCreateCGImageFromCVPixelBuffer func(pixelBuffer corevideo.CVPixelBufferRef, options corefoundation.CFDictionaryRef, imageOut *coregraphics.CGImageRef) int32
var _vTCreateCGImageFromCVPixelBufferErr error

func tryVTCreateCGImageFromCVPixelBuffer(pixelBuffer corevideo.CVPixelBufferRef, options corefoundation.CFDictionaryRef, imageOut *coregraphics.CGImageRef) (int32, error) {
	if _vTCreateCGImageFromCVPixelBuffer == nil {
		return 0, symbolCallError("VTCreateCGImageFromCVPixelBuffer", "10.11", _vTCreateCGImageFromCVPixelBufferErr)
	}
	return _vTCreateCGImageFromCVPixelBuffer(pixelBuffer, options, imageOut), nil
}

// VTCreateCGImageFromCVPixelBuffer creates a Core Graphics bitmap image or image mask using the provided pixel buffer.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTCreateCGImageFromCVPixelBuffer(_:options:imageOut:)
func VTCreateCGImageFromCVPixelBuffer(pixelBuffer corevideo.CVPixelBufferRef, options corefoundation.CFDictionaryRef, imageOut *coregraphics.CGImageRef) int32 {
	result, callErr := tryVTCreateCGImageFromCVPixelBuffer(pixelBuffer, options, imageOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTDecompressionSessionCanAcceptFormatDescription func(session VTDecompressionSessionRef, newFormatDesc uintptr) bool
var _vTDecompressionSessionCanAcceptFormatDescriptionErr error

func tryVTDecompressionSessionCanAcceptFormatDescription(session VTDecompressionSessionRef, newFormatDesc uintptr) (bool, error) {
	if _vTDecompressionSessionCanAcceptFormatDescription == nil {
		return false, symbolCallError("VTDecompressionSessionCanAcceptFormatDescription", "10.8", _vTDecompressionSessionCanAcceptFormatDescriptionErr)
	}
	return _vTDecompressionSessionCanAcceptFormatDescription(session, newFormatDesc), nil
}

// VTDecompressionSessionCanAcceptFormatDescription indicates whether the session can decode frames with the given format description.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSessionCanAcceptFormatDescription(_:formatDescription:)
func VTDecompressionSessionCanAcceptFormatDescription(session VTDecompressionSessionRef, newFormatDesc uintptr) bool {
	result, callErr := tryVTDecompressionSessionCanAcceptFormatDescription(session, newFormatDesc)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTDecompressionSessionCopyBlackPixelBuffer func(session VTDecompressionSessionRef, pixelBufferOut *corevideo.CVImageBufferRef) int32
var _vTDecompressionSessionCopyBlackPixelBufferErr error

func tryVTDecompressionSessionCopyBlackPixelBuffer(session VTDecompressionSessionRef, pixelBufferOut *corevideo.CVImageBufferRef) (int32, error) {
	if _vTDecompressionSessionCopyBlackPixelBuffer == nil {
		return 0, symbolCallError("VTDecompressionSessionCopyBlackPixelBuffer", "10.8", _vTDecompressionSessionCopyBlackPixelBufferErr)
	}
	return _vTDecompressionSessionCopyBlackPixelBuffer(session, pixelBufferOut), nil
}

// VTDecompressionSessionCopyBlackPixelBuffer copies a black pixel buffer from the decompression session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSessionCopyBlackPixelBuffer(_:pixelBufferOut:)
func VTDecompressionSessionCopyBlackPixelBuffer(session VTDecompressionSessionRef, pixelBufferOut *corevideo.CVImageBufferRef) int32 {
	result, callErr := tryVTDecompressionSessionCopyBlackPixelBuffer(session, pixelBufferOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTDecompressionSessionCreate func(allocator corefoundation.CFAllocatorRef, videoFormatDescription uintptr, videoDecoderSpecification corefoundation.CFDictionaryRef, destinationImageBufferAttributes corefoundation.CFDictionaryRef, outputCallback *VTDecompressionOutputCallbackRecord, decompressionSessionOut *VTDecompressionSessionRef) int32
var _vTDecompressionSessionCreateErr error

func tryVTDecompressionSessionCreate(allocator corefoundation.CFAllocatorRef, videoFormatDescription uintptr, videoDecoderSpecification corefoundation.CFDictionaryRef, destinationImageBufferAttributes corefoundation.CFDictionaryRef, outputCallback *VTDecompressionOutputCallbackRecord, decompressionSessionOut *VTDecompressionSessionRef) (int32, error) {
	if _vTDecompressionSessionCreate == nil {
		return 0, symbolCallError("VTDecompressionSessionCreate", "10.8", _vTDecompressionSessionCreateErr)
	}
	return _vTDecompressionSessionCreate(allocator, videoFormatDescription, videoDecoderSpecification, destinationImageBufferAttributes, outputCallback, decompressionSessionOut), nil
}

// VTDecompressionSessionCreate creates a session for decompressing video frames.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSessionCreate(allocator:formatDescription:decoderSpecification:imageBufferAttributes:outputCallback:decompressionSessionOut:)
func VTDecompressionSessionCreate(allocator corefoundation.CFAllocatorRef, videoFormatDescription uintptr, videoDecoderSpecification corefoundation.CFDictionaryRef, destinationImageBufferAttributes corefoundation.CFDictionaryRef, outputCallback *VTDecompressionOutputCallbackRecord, decompressionSessionOut *VTDecompressionSessionRef) int32 {
	result, callErr := tryVTDecompressionSessionCreate(allocator, videoFormatDescription, videoDecoderSpecification, destinationImageBufferAttributes, outputCallback, decompressionSessionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTDecompressionSessionDecodeFrame func(session VTDecompressionSessionRef, sampleBuffer uintptr, decodeFlags VTDecodeFrameFlags, sourceFrameRefCon unsafe.Pointer, infoFlagsOut *VTDecodeInfoFlags) int32
var _vTDecompressionSessionDecodeFrameErr error

func tryVTDecompressionSessionDecodeFrame(session VTDecompressionSessionRef, sampleBuffer uintptr, decodeFlags VTDecodeFrameFlags, sourceFrameRefCon unsafe.Pointer, infoFlagsOut *VTDecodeInfoFlags) (int32, error) {
	if _vTDecompressionSessionDecodeFrame == nil {
		return 0, symbolCallError("VTDecompressionSessionDecodeFrame", "10.8", _vTDecompressionSessionDecodeFrameErr)
	}
	return _vTDecompressionSessionDecodeFrame(session, sampleBuffer, decodeFlags, sourceFrameRefCon, infoFlagsOut), nil
}

// VTDecompressionSessionDecodeFrame decompresses a video frame.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSessionDecodeFrame(_:sampleBuffer:flags:frameRefcon:infoFlagsOut:)
func VTDecompressionSessionDecodeFrame(session VTDecompressionSessionRef, sampleBuffer uintptr, decodeFlags VTDecodeFrameFlags, sourceFrameRefCon unsafe.Pointer, infoFlagsOut *VTDecodeInfoFlags) int32 {
	result, callErr := tryVTDecompressionSessionDecodeFrame(session, sampleBuffer, decodeFlags, sourceFrameRefCon, infoFlagsOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTDecompressionSessionDecodeFrameWithMultiImageCapableOutputHandler func(session VTDecompressionSessionRef, sampleBuffer uintptr, decodeFlags VTDecodeFrameFlags, infoFlagsOut *VTDecodeInfoFlags, multiImageCapableOutputHandler unsafe.Pointer) int32
var _vTDecompressionSessionDecodeFrameWithMultiImageCapableOutputHandlerErr error

func tryVTDecompressionSessionDecodeFrameWithMultiImageCapableOutputHandler(session VTDecompressionSessionRef, sampleBuffer uintptr, decodeFlags VTDecodeFrameFlags, infoFlagsOut *VTDecodeInfoFlags, multiImageCapableOutputHandler VTDecompressionMultiImageCapableOutputHandler) (int32, error) {
	if _vTDecompressionSessionDecodeFrameWithMultiImageCapableOutputHandler == nil {
		return 0, symbolCallError("VTDecompressionSessionDecodeFrameWithMultiImageCapableOutputHandler", "14.0", _vTDecompressionSessionDecodeFrameWithMultiImageCapableOutputHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 int32, blockArg1 uint, blockArg2 corevideo.CVImageBufferRef, blockArg3 *uintptr, blockArg4 coremedia.CMTime, blockArg5 coremedia.CMTime) {
		multiImageCapableOutputHandler(blockArg0, blockArg1, blockArg2, blockArg3, blockArg4, blockArg5)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _vTDecompressionSessionDecodeFrameWithMultiImageCapableOutputHandler(session, sampleBuffer, decodeFlags, infoFlagsOut, _block0), nil
}

// VTDecompressionSessionDecodeFrameWithMultiImageCapableOutputHandler decompresses a multi-image frame and calls the specified output handler upon completion.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSessionDecodeFrameWithMultiImageCapableOutputHandler
func VTDecompressionSessionDecodeFrameWithMultiImageCapableOutputHandler(session VTDecompressionSessionRef, sampleBuffer uintptr, decodeFlags VTDecodeFrameFlags, infoFlagsOut *VTDecodeInfoFlags, multiImageCapableOutputHandler VTDecompressionMultiImageCapableOutputHandler) int32 {
	result, callErr := tryVTDecompressionSessionDecodeFrameWithMultiImageCapableOutputHandler(session, sampleBuffer, decodeFlags, infoFlagsOut, multiImageCapableOutputHandler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTDecompressionSessionDecodeFrameWithOptions func(session VTDecompressionSessionRef, sampleBuffer uintptr, decodeFlags VTDecodeFrameFlags, frameOptions corefoundation.CFDictionaryRef, sourceFrameRefCon unsafe.Pointer, infoFlagsOut *VTDecodeInfoFlags) int32
var _vTDecompressionSessionDecodeFrameWithOptionsErr error

func tryVTDecompressionSessionDecodeFrameWithOptions(session VTDecompressionSessionRef, sampleBuffer uintptr, decodeFlags VTDecodeFrameFlags, frameOptions corefoundation.CFDictionaryRef, sourceFrameRefCon unsafe.Pointer, infoFlagsOut *VTDecodeInfoFlags) (int32, error) {
	if _vTDecompressionSessionDecodeFrameWithOptions == nil {
		return 0, symbolCallError("VTDecompressionSessionDecodeFrameWithOptions", "15.0", _vTDecompressionSessionDecodeFrameWithOptionsErr)
	}
	return _vTDecompressionSessionDecodeFrameWithOptions(session, sampleBuffer, decodeFlags, frameOptions, sourceFrameRefCon, infoFlagsOut), nil
}

// VTDecompressionSessionDecodeFrameWithOptions.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSessionDecodeFrame(_:sampleBuffer:flags:frameOptions:frameRefcon:infoFlagsOut:)
func VTDecompressionSessionDecodeFrameWithOptions(session VTDecompressionSessionRef, sampleBuffer uintptr, decodeFlags VTDecodeFrameFlags, frameOptions corefoundation.CFDictionaryRef, sourceFrameRefCon unsafe.Pointer, infoFlagsOut *VTDecodeInfoFlags) int32 {
	result, callErr := tryVTDecompressionSessionDecodeFrameWithOptions(session, sampleBuffer, decodeFlags, frameOptions, sourceFrameRefCon, infoFlagsOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTDecompressionSessionDecodeFrameWithOptionsAndOutputHandler func(session VTDecompressionSessionRef, sampleBuffer uintptr, decodeFlags VTDecodeFrameFlags, frameOptions corefoundation.CFDictionaryRef, infoFlagsOut *VTDecodeInfoFlags, outputHandler unsafe.Pointer) int32
var _vTDecompressionSessionDecodeFrameWithOptionsAndOutputHandlerErr error

func tryVTDecompressionSessionDecodeFrameWithOptionsAndOutputHandler(session VTDecompressionSessionRef, sampleBuffer uintptr, decodeFlags VTDecodeFrameFlags, frameOptions corefoundation.CFDictionaryRef, infoFlagsOut *VTDecodeInfoFlags, outputHandler VTDecompressionOutputHandler) (int32, error) {
	if _vTDecompressionSessionDecodeFrameWithOptionsAndOutputHandler == nil {
		return 0, symbolCallError("VTDecompressionSessionDecodeFrameWithOptionsAndOutputHandler", "15.0", _vTDecompressionSessionDecodeFrameWithOptionsAndOutputHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 int32, blockArg1 uint, blockArg2 corevideo.CVImageBufferRef, blockArg3 coremedia.CMTime, blockArg4 coremedia.CMTime) {
		outputHandler(blockArg0, blockArg1, blockArg2, blockArg3, blockArg4)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _vTDecompressionSessionDecodeFrameWithOptionsAndOutputHandler(session, sampleBuffer, decodeFlags, frameOptions, infoFlagsOut, _block0), nil
}

// VTDecompressionSessionDecodeFrameWithOptionsAndOutputHandler.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSessionDecodeFrame(_:sampleBuffer:flags:frameOptions:infoFlagsOut:outputHandler:)
func VTDecompressionSessionDecodeFrameWithOptionsAndOutputHandler(session VTDecompressionSessionRef, sampleBuffer uintptr, decodeFlags VTDecodeFrameFlags, frameOptions corefoundation.CFDictionaryRef, infoFlagsOut *VTDecodeInfoFlags, outputHandler VTDecompressionOutputHandler) int32 {
	result, callErr := tryVTDecompressionSessionDecodeFrameWithOptionsAndOutputHandler(session, sampleBuffer, decodeFlags, frameOptions, infoFlagsOut, outputHandler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTDecompressionSessionDecodeFrameWithOutputHandler func(session VTDecompressionSessionRef, sampleBuffer uintptr, decodeFlags VTDecodeFrameFlags, infoFlagsOut *VTDecodeInfoFlags, outputHandler unsafe.Pointer) int32
var _vTDecompressionSessionDecodeFrameWithOutputHandlerErr error

func tryVTDecompressionSessionDecodeFrameWithOutputHandler(session VTDecompressionSessionRef, sampleBuffer uintptr, decodeFlags VTDecodeFrameFlags, infoFlagsOut *VTDecodeInfoFlags, outputHandler VTDecompressionOutputHandler) (int32, error) {
	if _vTDecompressionSessionDecodeFrameWithOutputHandler == nil {
		return 0, symbolCallError("VTDecompressionSessionDecodeFrameWithOutputHandler", "10.11", _vTDecompressionSessionDecodeFrameWithOutputHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 int32, blockArg1 uint, blockArg2 corevideo.CVImageBufferRef, blockArg3 coremedia.CMTime, blockArg4 coremedia.CMTime) {
		outputHandler(blockArg0, blockArg1, blockArg2, blockArg3, blockArg4)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _vTDecompressionSessionDecodeFrameWithOutputHandler(session, sampleBuffer, decodeFlags, infoFlagsOut, _block0), nil
}

// VTDecompressionSessionDecodeFrameWithOutputHandler decompresses a video frame and invokes the output callback when the decompression completes.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSessionDecodeFrame(_:sampleBuffer:flags:infoFlagsOut:outputHandler:)
func VTDecompressionSessionDecodeFrameWithOutputHandler(session VTDecompressionSessionRef, sampleBuffer uintptr, decodeFlags VTDecodeFrameFlags, infoFlagsOut *VTDecodeInfoFlags, outputHandler VTDecompressionOutputHandler) int32 {
	result, callErr := tryVTDecompressionSessionDecodeFrameWithOutputHandler(session, sampleBuffer, decodeFlags, infoFlagsOut, outputHandler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTDecompressionSessionFinishDelayedFrames func(session VTDecompressionSessionRef) int32
var _vTDecompressionSessionFinishDelayedFramesErr error

func tryVTDecompressionSessionFinishDelayedFrames(session VTDecompressionSessionRef) (int32, error) {
	if _vTDecompressionSessionFinishDelayedFrames == nil {
		return 0, symbolCallError("VTDecompressionSessionFinishDelayedFrames", "10.8", _vTDecompressionSessionFinishDelayedFramesErr)
	}
	return _vTDecompressionSessionFinishDelayedFrames(session), nil
}

// VTDecompressionSessionFinishDelayedFrames directs the decompression session to emit all delayed frames.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSessionFinishDelayedFrames(_:)
func VTDecompressionSessionFinishDelayedFrames(session VTDecompressionSessionRef) int32 {
	result, callErr := tryVTDecompressionSessionFinishDelayedFrames(session)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTDecompressionSessionGetTypeID func() uint
var _vTDecompressionSessionGetTypeIDErr error

func tryVTDecompressionSessionGetTypeID() (uint, error) {
	if _vTDecompressionSessionGetTypeID == nil {
		return 0, symbolCallError("VTDecompressionSessionGetTypeID", "10.8", _vTDecompressionSessionGetTypeIDErr)
	}
	return _vTDecompressionSessionGetTypeID(), nil
}

// VTDecompressionSessionGetTypeID returns the Core Foundation type identifier for the decompression session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSessionGetTypeID()
func VTDecompressionSessionGetTypeID() uint {
	result, callErr := tryVTDecompressionSessionGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTDecompressionSessionInvalidate func(session VTDecompressionSessionRef)
var _vTDecompressionSessionInvalidateErr error

func tryVTDecompressionSessionInvalidate(session VTDecompressionSessionRef) error {
	if _vTDecompressionSessionInvalidate == nil {
		return symbolCallError("VTDecompressionSessionInvalidate", "10.8", _vTDecompressionSessionInvalidateErr)
	}
	_vTDecompressionSessionInvalidate(session)
	return nil
}

// VTDecompressionSessionInvalidate tears down a decompression session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSessionInvalidate(_:)
func VTDecompressionSessionInvalidate(session VTDecompressionSessionRef) {
	if callErr := tryVTDecompressionSessionInvalidate(session); callErr != nil {
		panic(callErr)
	}
}

var _vTDecompressionSessionSetMultiImageCallback func(decompressionSession VTDecompressionSessionRef, outputMultiImageCallback VTDecompressionOutputMultiImageCallback, outputMultiImageRefcon unsafe.Pointer) int32
var _vTDecompressionSessionSetMultiImageCallbackErr error

func tryVTDecompressionSessionSetMultiImageCallback(decompressionSession VTDecompressionSessionRef, outputMultiImageCallback VTDecompressionOutputMultiImageCallback, outputMultiImageRefcon unsafe.Pointer) (int32, error) {
	if _vTDecompressionSessionSetMultiImageCallback == nil {
		return 0, symbolCallError("VTDecompressionSessionSetMultiImageCallback", "14.0", _vTDecompressionSessionSetMultiImageCallbackErr)
	}
	return _vTDecompressionSessionSetMultiImageCallback(decompressionSession, outputMultiImageCallback, outputMultiImageRefcon), nil
}

// VTDecompressionSessionSetMultiImageCallback provides a callback capable of receiving multiple images for individual frame decoding requests.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSessionSetMultiImageCallback
func VTDecompressionSessionSetMultiImageCallback(decompressionSession VTDecompressionSessionRef, outputMultiImageCallback VTDecompressionOutputMultiImageCallback, outputMultiImageRefcon unsafe.Pointer) int32 {
	result, callErr := tryVTDecompressionSessionSetMultiImageCallback(decompressionSession, outputMultiImageCallback, outputMultiImageRefcon)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTDecompressionSessionWaitForAsynchronousFrames func(session VTDecompressionSessionRef) int32
var _vTDecompressionSessionWaitForAsynchronousFramesErr error

func tryVTDecompressionSessionWaitForAsynchronousFrames(session VTDecompressionSessionRef) (int32, error) {
	if _vTDecompressionSessionWaitForAsynchronousFrames == nil {
		return 0, symbolCallError("VTDecompressionSessionWaitForAsynchronousFrames", "10.8", _vTDecompressionSessionWaitForAsynchronousFramesErr)
	}
	return _vTDecompressionSessionWaitForAsynchronousFrames(session), nil
}

// VTDecompressionSessionWaitForAsynchronousFrames waits for any and all outstanding asynchronous and delayed frames to complete, then returns.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTDecompressionSessionWaitForAsynchronousFrames(_:)
func VTDecompressionSessionWaitForAsynchronousFrames(session VTDecompressionSessionRef) int32 {
	result, callErr := tryVTDecompressionSessionWaitForAsynchronousFrames(session)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTFrameSiloAddSampleBuffer func(silo VTFrameSiloRef, sampleBuffer uintptr) int32
var _vTFrameSiloAddSampleBufferErr error

func tryVTFrameSiloAddSampleBuffer(silo VTFrameSiloRef, sampleBuffer uintptr) (int32, error) {
	if _vTFrameSiloAddSampleBuffer == nil {
		return 0, symbolCallError("VTFrameSiloAddSampleBuffer", "10.10", _vTFrameSiloAddSampleBufferErr)
	}
	return _vTFrameSiloAddSampleBuffer(silo, sampleBuffer), nil
}

// VTFrameSiloAddSampleBuffer adds a sample buffer to a frame silo object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameSiloAddSampleBuffer(_:sampleBuffer:)
func VTFrameSiloAddSampleBuffer(silo VTFrameSiloRef, sampleBuffer uintptr) int32 {
	result, callErr := tryVTFrameSiloAddSampleBuffer(silo, sampleBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTFrameSiloCallBlockForEachSampleBuffer func(silo VTFrameSiloRef, timeRange coremedia.CMTimeRange, handler func(uintptr) int32) int32
var _vTFrameSiloCallBlockForEachSampleBufferErr error

func tryVTFrameSiloCallBlockForEachSampleBuffer(silo VTFrameSiloRef, timeRange coremedia.CMTimeRange, handler func(uintptr) int32) (int32, error) {
	if _vTFrameSiloCallBlockForEachSampleBuffer == nil {
		return 0, symbolCallError("VTFrameSiloCallBlockForEachSampleBuffer", "10.10", _vTFrameSiloCallBlockForEachSampleBufferErr)
	}
	return _vTFrameSiloCallBlockForEachSampleBuffer(silo, timeRange, handler), nil
}

// VTFrameSiloCallBlockForEachSampleBuffer retrieves sample buffers from a frame silo object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameSiloCallBlockForEachSampleBuffer(_:in:handler:)
func VTFrameSiloCallBlockForEachSampleBuffer(silo VTFrameSiloRef, timeRange coremedia.CMTimeRange, handler func(uintptr) int32) int32 {
	result, callErr := tryVTFrameSiloCallBlockForEachSampleBuffer(silo, timeRange, handler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTFrameSiloCallFunctionForEachSampleBuffer func(silo VTFrameSiloRef, timeRange coremedia.CMTimeRange, refcon uintptr, callback func(unsafe.Pointer, uintptr) int32) int32
var _vTFrameSiloCallFunctionForEachSampleBufferErr error

func tryVTFrameSiloCallFunctionForEachSampleBuffer(silo VTFrameSiloRef, timeRange coremedia.CMTimeRange, refcon uintptr, callback func(unsafe.Pointer, uintptr) int32) (int32, error) {
	if _vTFrameSiloCallFunctionForEachSampleBuffer == nil {
		return 0, symbolCallError("VTFrameSiloCallFunctionForEachSampleBuffer", "10.10", _vTFrameSiloCallFunctionForEachSampleBufferErr)
	}
	return _vTFrameSiloCallFunctionForEachSampleBuffer(silo, timeRange, refcon, callback), nil
}

// VTFrameSiloCallFunctionForEachSampleBuffer retrieves sample buffers from a frame silo object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameSiloCallFunctionForEachSampleBuffer(_:in:refcon:callback:)
func VTFrameSiloCallFunctionForEachSampleBuffer(silo VTFrameSiloRef, timeRange coremedia.CMTimeRange, refcon uintptr, callback func(unsafe.Pointer, uintptr) int32) int32 {
	result, callErr := tryVTFrameSiloCallFunctionForEachSampleBuffer(silo, timeRange, refcon, callback)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTFrameSiloCreate func(allocator corefoundation.CFAllocatorRef, fileURL corefoundation.CFURLRef, timeRange coremedia.CMTimeRange, options corefoundation.CFDictionaryRef, frameSiloOut *VTFrameSiloRef) int32
var _vTFrameSiloCreateErr error

func tryVTFrameSiloCreate(allocator corefoundation.CFAllocatorRef, fileURL corefoundation.CFURLRef, timeRange coremedia.CMTimeRange, options corefoundation.CFDictionaryRef, frameSiloOut *VTFrameSiloRef) (int32, error) {
	if _vTFrameSiloCreate == nil {
		return 0, symbolCallError("VTFrameSiloCreate", "10.10", _vTFrameSiloCreateErr)
	}
	return _vTFrameSiloCreate(allocator, fileURL, timeRange, options, frameSiloOut), nil
}

// VTFrameSiloCreate creates a frame silo object using a temporary file.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameSiloCreate(allocator:fileURL:timeRange:options:frameSiloOut:)
func VTFrameSiloCreate(allocator corefoundation.CFAllocatorRef, fileURL corefoundation.CFURLRef, timeRange coremedia.CMTimeRange, options corefoundation.CFDictionaryRef, frameSiloOut *VTFrameSiloRef) int32 {
	result, callErr := tryVTFrameSiloCreate(allocator, fileURL, timeRange, options, frameSiloOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTFrameSiloGetProgressOfCurrentPass func(silo VTFrameSiloRef, progressOut *float32) int32
var _vTFrameSiloGetProgressOfCurrentPassErr error

func tryVTFrameSiloGetProgressOfCurrentPass(silo VTFrameSiloRef, progressOut *float32) (int32, error) {
	if _vTFrameSiloGetProgressOfCurrentPass == nil {
		return 0, symbolCallError("VTFrameSiloGetProgressOfCurrentPass", "10.10", _vTFrameSiloGetProgressOfCurrentPassErr)
	}
	return _vTFrameSiloGetProgressOfCurrentPass(silo, progressOut), nil
}

// VTFrameSiloGetProgressOfCurrentPass gets the progress of the current pass.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameSiloGetProgressOfCurrentPass(_:progressOut:)
func VTFrameSiloGetProgressOfCurrentPass(silo VTFrameSiloRef, progressOut *float32) int32 {
	result, callErr := tryVTFrameSiloGetProgressOfCurrentPass(silo, progressOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTFrameSiloGetTypeID func() uint
var _vTFrameSiloGetTypeIDErr error

func tryVTFrameSiloGetTypeID() (uint, error) {
	if _vTFrameSiloGetTypeID == nil {
		return 0, symbolCallError("VTFrameSiloGetTypeID", "10.10", _vTFrameSiloGetTypeIDErr)
	}
	return _vTFrameSiloGetTypeID(), nil
}

// VTFrameSiloGetTypeID retrieves the Core Foundation type identifier for the frame silo object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameSiloGetTypeID()
func VTFrameSiloGetTypeID() uint {
	result, callErr := tryVTFrameSiloGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTFrameSiloSetTimeRangesForNextPass func(silo VTFrameSiloRef, timeRangeCount int, timeRangeArray *coremedia.CMTimeRange) int32
var _vTFrameSiloSetTimeRangesForNextPassErr error

func tryVTFrameSiloSetTimeRangesForNextPass(silo VTFrameSiloRef, timeRangeCount int, timeRangeArray *coremedia.CMTimeRange) (int32, error) {
	if _vTFrameSiloSetTimeRangesForNextPass == nil {
		return 0, symbolCallError("VTFrameSiloSetTimeRangesForNextPass", "10.10", _vTFrameSiloSetTimeRangesForNextPassErr)
	}
	return _vTFrameSiloSetTimeRangesForNextPass(silo, timeRangeCount, timeRangeArray), nil
}

// VTFrameSiloSetTimeRangesForNextPass begins a new pass of samples to be added to a frame silo object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTFrameSiloSetTimeRangesForNextPass(_:timeRangeCount:timeRangeArray:)
func VTFrameSiloSetTimeRangesForNextPass(silo VTFrameSiloRef, timeRangeCount int, timeRangeArray *coremedia.CMTimeRange) int32 {
	result, callErr := tryVTFrameSiloSetTimeRangesForNextPass(silo, timeRangeCount, timeRangeArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTHDRPerFrameMetadataGenerationSessionAttachMetadata func(hdrPerFrameMetadataGenerationSession VTHDRPerFrameMetadataGenerationSessionRef, pixelBuffer corevideo.CVPixelBufferRef, sceneChange bool) int32
var _vTHDRPerFrameMetadataGenerationSessionAttachMetadataErr error

func tryVTHDRPerFrameMetadataGenerationSessionAttachMetadata(hdrPerFrameMetadataGenerationSession VTHDRPerFrameMetadataGenerationSessionRef, pixelBuffer corevideo.CVPixelBufferRef, sceneChange bool) (int32, error) {
	if _vTHDRPerFrameMetadataGenerationSessionAttachMetadata == nil {
		return 0, symbolCallError("VTHDRPerFrameMetadataGenerationSessionAttachMetadata", "15.0", _vTHDRPerFrameMetadataGenerationSessionAttachMetadataErr)
	}
	return _vTHDRPerFrameMetadataGenerationSessionAttachMetadata(hdrPerFrameMetadataGenerationSession, pixelBuffer, sceneChange), nil
}

// VTHDRPerFrameMetadataGenerationSessionAttachMetadata attaches per-frame metadata to the pixel buffer and the backing IOSurface.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTHDRPerFrameMetadataGenerationSessionAttachMetadata
func VTHDRPerFrameMetadataGenerationSessionAttachMetadata(hdrPerFrameMetadataGenerationSession VTHDRPerFrameMetadataGenerationSessionRef, pixelBuffer corevideo.CVPixelBufferRef, sceneChange bool) int32 {
	result, callErr := tryVTHDRPerFrameMetadataGenerationSessionAttachMetadata(hdrPerFrameMetadataGenerationSession, pixelBuffer, sceneChange)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTHDRPerFrameMetadataGenerationSessionCreate func(allocator corefoundation.CFAllocatorRef, framesPerSecond float32, options corefoundation.CFDictionaryRef, hdrPerFrameMetadataGenerationSessionOut *VTHDRPerFrameMetadataGenerationSessionRef) int32
var _vTHDRPerFrameMetadataGenerationSessionCreateErr error

func tryVTHDRPerFrameMetadataGenerationSessionCreate(allocator corefoundation.CFAllocatorRef, framesPerSecond float32, options corefoundation.CFDictionaryRef, hdrPerFrameMetadataGenerationSessionOut *VTHDRPerFrameMetadataGenerationSessionRef) (int32, error) {
	if _vTHDRPerFrameMetadataGenerationSessionCreate == nil {
		return 0, symbolCallError("VTHDRPerFrameMetadataGenerationSessionCreate", "15.0", _vTHDRPerFrameMetadataGenerationSessionCreateErr)
	}
	return _vTHDRPerFrameMetadataGenerationSessionCreate(allocator, framesPerSecond, options, hdrPerFrameMetadataGenerationSessionOut), nil
}

// VTHDRPerFrameMetadataGenerationSessionCreate creates a metadata generation session object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTHDRPerFrameMetadataGenerationSessionCreate
func VTHDRPerFrameMetadataGenerationSessionCreate(allocator corefoundation.CFAllocatorRef, framesPerSecond float32, options corefoundation.CFDictionaryRef, hdrPerFrameMetadataGenerationSessionOut *VTHDRPerFrameMetadataGenerationSessionRef) int32 {
	result, callErr := tryVTHDRPerFrameMetadataGenerationSessionCreate(allocator, framesPerSecond, options, hdrPerFrameMetadataGenerationSessionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTHDRPerFrameMetadataGenerationSessionGetTypeID func() uint
var _vTHDRPerFrameMetadataGenerationSessionGetTypeIDErr error

func tryVTHDRPerFrameMetadataGenerationSessionGetTypeID() (uint, error) {
	if _vTHDRPerFrameMetadataGenerationSessionGetTypeID == nil {
		return 0, symbolCallError("VTHDRPerFrameMetadataGenerationSessionGetTypeID", "15.0", _vTHDRPerFrameMetadataGenerationSessionGetTypeIDErr)
	}
	return _vTHDRPerFrameMetadataGenerationSessionGetTypeID(), nil
}

// VTHDRPerFrameMetadataGenerationSessionGetTypeID retrieves the Core Foundation type identifier for the session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTHDRPerFrameMetadataGenerationSessionGetTypeID
func VTHDRPerFrameMetadataGenerationSessionGetTypeID() uint {
	result, callErr := tryVTHDRPerFrameMetadataGenerationSessionGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTIsHardwareDecodeSupported func(codecType uint32) bool
var _vTIsHardwareDecodeSupportedErr error

func tryVTIsHardwareDecodeSupported(codecType uint32) (bool, error) {
	if _vTIsHardwareDecodeSupported == nil {
		return false, symbolCallError("VTIsHardwareDecodeSupported", "10.13", _vTIsHardwareDecodeSupportedErr)
	}
	return _vTIsHardwareDecodeSupported(codecType), nil
}

// VTIsHardwareDecodeSupported returns a Boolean value that indicates whether the current system supports hardware decode for the specified codec.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTIsHardwareDecodeSupported(_:)
func VTIsHardwareDecodeSupported(codecType uint32) bool {
	result, callErr := tryVTIsHardwareDecodeSupported(codecType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTIsStereoMVHEVCDecodeSupported func() bool
var _vTIsStereoMVHEVCDecodeSupportedErr error

func tryVTIsStereoMVHEVCDecodeSupported() (bool, error) {
	if _vTIsStereoMVHEVCDecodeSupported == nil {
		return false, symbolCallError("VTIsStereoMVHEVCDecodeSupported", "14.0", _vTIsStereoMVHEVCDecodeSupportedErr)
	}
	return _vTIsStereoMVHEVCDecodeSupported(), nil
}

// VTIsStereoMVHEVCDecodeSupported returns a Boolean value that indicates whether the system supports MV-HEVC decoding.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTIsStereoMVHEVCDecodeSupported()
func VTIsStereoMVHEVCDecodeSupported() bool {
	result, callErr := tryVTIsStereoMVHEVCDecodeSupported()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTIsStereoMVHEVCEncodeSupported func() bool
var _vTIsStereoMVHEVCEncodeSupportedErr error

func tryVTIsStereoMVHEVCEncodeSupported() (bool, error) {
	if _vTIsStereoMVHEVCEncodeSupported == nil {
		return false, symbolCallError("VTIsStereoMVHEVCEncodeSupported", "14.0", _vTIsStereoMVHEVCEncodeSupportedErr)
	}
	return _vTIsStereoMVHEVCEncodeSupported(), nil
}

// VTIsStereoMVHEVCEncodeSupported returns a Boolean value that indicates whether the system supports MV-HEVC encoding.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTIsStereoMVHEVCEncodeSupported()
func VTIsStereoMVHEVCEncodeSupported() bool {
	result, callErr := tryVTIsStereoMVHEVCEncodeSupported()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTMotionEstimationSessionCompleteFrames func(session VTMotionEstimationSessionRef) int32
var _vTMotionEstimationSessionCompleteFramesErr error

func tryVTMotionEstimationSessionCompleteFrames(session VTMotionEstimationSessionRef) (int32, error) {
	if _vTMotionEstimationSessionCompleteFrames == nil {
		return 0, symbolCallError("VTMotionEstimationSessionCompleteFrames", "26.0", _vTMotionEstimationSessionCompleteFramesErr)
	}
	return _vTMotionEstimationSessionCompleteFrames(session), nil
}

// VTMotionEstimationSessionCompleteFrames directs the motion-estimation session to emit all pending frames and waits for completion.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionEstimationSessionCompleteFrames
func VTMotionEstimationSessionCompleteFrames(session VTMotionEstimationSessionRef) int32 {
	result, callErr := tryVTMotionEstimationSessionCompleteFrames(session)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTMotionEstimationSessionCopySourcePixelBufferAttributes func(motionEstimationSession VTMotionEstimationSessionRef, attributesOut *corefoundation.CFDictionaryRef) int32
var _vTMotionEstimationSessionCopySourcePixelBufferAttributesErr error

func tryVTMotionEstimationSessionCopySourcePixelBufferAttributes(motionEstimationSession VTMotionEstimationSessionRef, attributesOut *corefoundation.CFDictionaryRef) (int32, error) {
	if _vTMotionEstimationSessionCopySourcePixelBufferAttributes == nil {
		return 0, symbolCallError("VTMotionEstimationSessionCopySourcePixelBufferAttributes", "26.0", _vTMotionEstimationSessionCopySourcePixelBufferAttributesErr)
	}
	return _vTMotionEstimationSessionCopySourcePixelBufferAttributes(motionEstimationSession, attributesOut), nil
}

// VTMotionEstimationSessionCopySourcePixelBufferAttributes copies the attributes for source pixel buffers expected by motion-estimation session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionEstimationSessionCopySourcePixelBufferAttributes
func VTMotionEstimationSessionCopySourcePixelBufferAttributes(motionEstimationSession VTMotionEstimationSessionRef, attributesOut *corefoundation.CFDictionaryRef) int32 {
	result, callErr := tryVTMotionEstimationSessionCopySourcePixelBufferAttributes(motionEstimationSession, attributesOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTMotionEstimationSessionCreate func(allocator corefoundation.CFAllocatorRef, motionVectorProcessorSelectionOptions corefoundation.CFDictionaryRef, width uint32, height uint32, motionEstimationSessionOut *VTMotionEstimationSessionRef) int32
var _vTMotionEstimationSessionCreateErr error

func tryVTMotionEstimationSessionCreate(allocator corefoundation.CFAllocatorRef, motionVectorProcessorSelectionOptions corefoundation.CFDictionaryRef, width uint32, height uint32, motionEstimationSessionOut *VTMotionEstimationSessionRef) (int32, error) {
	if _vTMotionEstimationSessionCreate == nil {
		return 0, symbolCallError("VTMotionEstimationSessionCreate", "26.0", _vTMotionEstimationSessionCreateErr)
	}
	return _vTMotionEstimationSessionCreate(allocator, motionVectorProcessorSelectionOptions, width, height, motionEstimationSessionOut), nil
}

// VTMotionEstimationSessionCreate creates a session you use to generate a pixel buffer of motion vectors from two pixel buffers.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionEstimationSessionCreate
func VTMotionEstimationSessionCreate(allocator corefoundation.CFAllocatorRef, motionVectorProcessorSelectionOptions corefoundation.CFDictionaryRef, width uint32, height uint32, motionEstimationSessionOut *VTMotionEstimationSessionRef) int32 {
	result, callErr := tryVTMotionEstimationSessionCreate(allocator, motionVectorProcessorSelectionOptions, width, height, motionEstimationSessionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTMotionEstimationSessionEstimateMotionVectors func(session VTMotionEstimationSessionRef, referenceImage corevideo.CVPixelBufferRef, currentImage corevideo.CVPixelBufferRef, motionEstimationFrameFlags VTMotionEstimationFrameFlags, additionalFrameOptions corefoundation.CFDictionaryRef, outputHandler unsafe.Pointer) int32
var _vTMotionEstimationSessionEstimateMotionVectorsErr error

func tryVTMotionEstimationSessionEstimateMotionVectors(session VTMotionEstimationSessionRef, referenceImage corevideo.CVPixelBufferRef, currentImage corevideo.CVPixelBufferRef, motionEstimationFrameFlags VTMotionEstimationFrameFlags, additionalFrameOptions corefoundation.CFDictionaryRef, outputHandler VTMotionEstimationOutputHandler) (int32, error) {
	if _vTMotionEstimationSessionEstimateMotionVectors == nil {
		return 0, symbolCallError("VTMotionEstimationSessionEstimateMotionVectors", "26.0", _vTMotionEstimationSessionEstimateMotionVectorsErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 int32, blockArg1 uint, blockArg2 corefoundation.CFDictionaryRef, blockArg3 corevideo.CVImageBufferRef) {
		outputHandler(blockArg0, blockArg1, blockArg2, blockArg3)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _vTMotionEstimationSessionEstimateMotionVectors(session, referenceImage, currentImage, motionEstimationFrameFlags, additionalFrameOptions, _block0), nil
}

// VTMotionEstimationSessionEstimateMotionVectors creates a new pixel buffer that contains motion vectors between the input pixel buffers.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionEstimationSessionEstimateMotionVectors
func VTMotionEstimationSessionEstimateMotionVectors(session VTMotionEstimationSessionRef, referenceImage corevideo.CVPixelBufferRef, currentImage corevideo.CVPixelBufferRef, motionEstimationFrameFlags VTMotionEstimationFrameFlags, additionalFrameOptions corefoundation.CFDictionaryRef, outputHandler VTMotionEstimationOutputHandler) int32 {
	result, callErr := tryVTMotionEstimationSessionEstimateMotionVectors(session, referenceImage, currentImage, motionEstimationFrameFlags, additionalFrameOptions, outputHandler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTMotionEstimationSessionGetTypeID func() uint
var _vTMotionEstimationSessionGetTypeIDErr error

func tryVTMotionEstimationSessionGetTypeID() (uint, error) {
	if _vTMotionEstimationSessionGetTypeID == nil {
		return 0, symbolCallError("VTMotionEstimationSessionGetTypeID", "26.0", _vTMotionEstimationSessionGetTypeIDErr)
	}
	return _vTMotionEstimationSessionGetTypeID(), nil
}

// VTMotionEstimationSessionGetTypeID get the CoreFoundation type identifier for motion-estimation session type.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionEstimationSessionGetTypeID
func VTMotionEstimationSessionGetTypeID() uint {
	result, callErr := tryVTMotionEstimationSessionGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTMotionEstimationSessionInvalidate func(session VTMotionEstimationSessionRef)
var _vTMotionEstimationSessionInvalidateErr error

func tryVTMotionEstimationSessionInvalidate(session VTMotionEstimationSessionRef) error {
	if _vTMotionEstimationSessionInvalidate == nil {
		return symbolCallError("VTMotionEstimationSessionInvalidate", "26.0", _vTMotionEstimationSessionInvalidateErr)
	}
	_vTMotionEstimationSessionInvalidate(session)
	return nil
}

// VTMotionEstimationSessionInvalidate tears down a motion-estimation session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMotionEstimationSessionInvalidate
func VTMotionEstimationSessionInvalidate(session VTMotionEstimationSessionRef) {
	if callErr := tryVTMotionEstimationSessionInvalidate(session); callErr != nil {
		panic(callErr)
	}
}

var _vTMultiPassStorageClose func(multiPassStorage VTMultiPassStorageRef) int32
var _vTMultiPassStorageCloseErr error

func tryVTMultiPassStorageClose(multiPassStorage VTMultiPassStorageRef) (int32, error) {
	if _vTMultiPassStorageClose == nil {
		return 0, symbolCallError("VTMultiPassStorageClose", "10.10", _vTMultiPassStorageCloseErr)
	}
	return _vTMultiPassStorageClose(multiPassStorage), nil
}

// VTMultiPassStorageClose ensures that any pending data is written to the multipass storage file and closes the file.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMultiPassStorageClose(_:)
func VTMultiPassStorageClose(multiPassStorage VTMultiPassStorageRef) int32 {
	result, callErr := tryVTMultiPassStorageClose(multiPassStorage)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTMultiPassStorageCreate func(allocator corefoundation.CFAllocatorRef, fileURL corefoundation.CFURLRef, timeRange coremedia.CMTimeRange, options corefoundation.CFDictionaryRef, multiPassStorageOut *VTMultiPassStorageRef) int32
var _vTMultiPassStorageCreateErr error

func tryVTMultiPassStorageCreate(allocator corefoundation.CFAllocatorRef, fileURL corefoundation.CFURLRef, timeRange coremedia.CMTimeRange, options corefoundation.CFDictionaryRef, multiPassStorageOut *VTMultiPassStorageRef) (int32, error) {
	if _vTMultiPassStorageCreate == nil {
		return 0, symbolCallError("VTMultiPassStorageCreate", "10.10", _vTMultiPassStorageCreateErr)
	}
	return _vTMultiPassStorageCreate(allocator, fileURL, timeRange, options, multiPassStorageOut), nil
}

// VTMultiPassStorageCreate creates a multipass storage object using a temporary file.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMultiPassStorageCreate(allocator:fileURL:timeRange:options:multiPassStorageOut:)
func VTMultiPassStorageCreate(allocator corefoundation.CFAllocatorRef, fileURL corefoundation.CFURLRef, timeRange coremedia.CMTimeRange, options corefoundation.CFDictionaryRef, multiPassStorageOut *VTMultiPassStorageRef) int32 {
	result, callErr := tryVTMultiPassStorageCreate(allocator, fileURL, timeRange, options, multiPassStorageOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTMultiPassStorageGetTypeID func() uint
var _vTMultiPassStorageGetTypeIDErr error

func tryVTMultiPassStorageGetTypeID() (uint, error) {
	if _vTMultiPassStorageGetTypeID == nil {
		return 0, symbolCallError("VTMultiPassStorageGetTypeID", "10.10", _vTMultiPassStorageGetTypeIDErr)
	}
	return _vTMultiPassStorageGetTypeID(), nil
}

// VTMultiPassStorageGetTypeID retrieves the Core Foundation type identifier for the multipass storage object.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTMultiPassStorageGetTypeID()
func VTMultiPassStorageGetTypeID() uint {
	result, callErr := tryVTMultiPassStorageGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTPixelRotationSessionCreate func(allocator corefoundation.CFAllocatorRef, pixelRotationSessionOut *VTPixelRotationSessionRef) int32
var _vTPixelRotationSessionCreateErr error

func tryVTPixelRotationSessionCreate(allocator corefoundation.CFAllocatorRef, pixelRotationSessionOut *VTPixelRotationSessionRef) (int32, error) {
	if _vTPixelRotationSessionCreate == nil {
		return 0, symbolCallError("VTPixelRotationSessionCreate", "13.0", _vTPixelRotationSessionCreateErr)
	}
	return _vTPixelRotationSessionCreate(allocator, pixelRotationSessionOut), nil
}

// VTPixelRotationSessionCreate creates a session to rotate images between pixel buffers.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTPixelRotationSessionCreate(_:_:)
func VTPixelRotationSessionCreate(allocator corefoundation.CFAllocatorRef, pixelRotationSessionOut *VTPixelRotationSessionRef) int32 {
	result, callErr := tryVTPixelRotationSessionCreate(allocator, pixelRotationSessionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTPixelRotationSessionGetTypeID func() uint
var _vTPixelRotationSessionGetTypeIDErr error

func tryVTPixelRotationSessionGetTypeID() (uint, error) {
	if _vTPixelRotationSessionGetTypeID == nil {
		return 0, symbolCallError("VTPixelRotationSessionGetTypeID", "13.0", _vTPixelRotationSessionGetTypeIDErr)
	}
	return _vTPixelRotationSessionGetTypeID(), nil
}

// VTPixelRotationSessionGetTypeID returns the Core Foundation type identifier for the rotation session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTPixelRotationSessionGetTypeID()
func VTPixelRotationSessionGetTypeID() uint {
	result, callErr := tryVTPixelRotationSessionGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTPixelRotationSessionInvalidate func(session VTPixelRotationSessionRef)
var _vTPixelRotationSessionInvalidateErr error

func tryVTPixelRotationSessionInvalidate(session VTPixelRotationSessionRef) error {
	if _vTPixelRotationSessionInvalidate == nil {
		return symbolCallError("VTPixelRotationSessionInvalidate", "13.0", _vTPixelRotationSessionInvalidateErr)
	}
	_vTPixelRotationSessionInvalidate(session)
	return nil
}

// VTPixelRotationSessionInvalidate tears down a pixel rotation session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTPixelRotationSessionInvalidate(_:)
func VTPixelRotationSessionInvalidate(session VTPixelRotationSessionRef) {
	if callErr := tryVTPixelRotationSessionInvalidate(session); callErr != nil {
		panic(callErr)
	}
}

var _vTPixelRotationSessionRotateImage func(session VTPixelRotationSessionRef, sourceBuffer corevideo.CVPixelBufferRef, destinationBuffer corevideo.CVPixelBufferRef) int32
var _vTPixelRotationSessionRotateImageErr error

func tryVTPixelRotationSessionRotateImage(session VTPixelRotationSessionRef, sourceBuffer corevideo.CVPixelBufferRef, destinationBuffer corevideo.CVPixelBufferRef) (int32, error) {
	if _vTPixelRotationSessionRotateImage == nil {
		return 0, symbolCallError("VTPixelRotationSessionRotateImage", "13.0", _vTPixelRotationSessionRotateImageErr)
	}
	return _vTPixelRotationSessionRotateImage(session, sourceBuffer, destinationBuffer), nil
}

// VTPixelRotationSessionRotateImage rotates a source pixel buffer and writes the output to the destination pixel buffer.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTPixelRotationSessionRotateImage(_:_:_:)
func VTPixelRotationSessionRotateImage(session VTPixelRotationSessionRef, sourceBuffer corevideo.CVPixelBufferRef, destinationBuffer corevideo.CVPixelBufferRef) int32 {
	result, callErr := tryVTPixelRotationSessionRotateImage(session, sourceBuffer, destinationBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTPixelTransferSessionCreate func(allocator corefoundation.CFAllocatorRef, pixelTransferSessionOut *VTPixelTransferSessionRef) int32
var _vTPixelTransferSessionCreateErr error

func tryVTPixelTransferSessionCreate(allocator corefoundation.CFAllocatorRef, pixelTransferSessionOut *VTPixelTransferSessionRef) (int32, error) {
	if _vTPixelTransferSessionCreate == nil {
		return 0, symbolCallError("VTPixelTransferSessionCreate", "10.8", _vTPixelTransferSessionCreateErr)
	}
	return _vTPixelTransferSessionCreate(allocator, pixelTransferSessionOut), nil
}

// VTPixelTransferSessionCreate creates a session for transferring images between Core Video image buffers that hold pixels in main memory.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTPixelTransferSessionCreate(allocator:pixelTransferSessionOut:)
func VTPixelTransferSessionCreate(allocator corefoundation.CFAllocatorRef, pixelTransferSessionOut *VTPixelTransferSessionRef) int32 {
	result, callErr := tryVTPixelTransferSessionCreate(allocator, pixelTransferSessionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTPixelTransferSessionGetTypeID func() uint
var _vTPixelTransferSessionGetTypeIDErr error

func tryVTPixelTransferSessionGetTypeID() (uint, error) {
	if _vTPixelTransferSessionGetTypeID == nil {
		return 0, symbolCallError("VTPixelTransferSessionGetTypeID", "10.8", _vTPixelTransferSessionGetTypeIDErr)
	}
	return _vTPixelTransferSessionGetTypeID(), nil
}

// VTPixelTransferSessionGetTypeID retrieves the Core Foundation type identifier for the pixel transfer session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTPixelTransferSessionGetTypeID()
func VTPixelTransferSessionGetTypeID() uint {
	result, callErr := tryVTPixelTransferSessionGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTPixelTransferSessionInvalidate func(session VTPixelTransferSessionRef)
var _vTPixelTransferSessionInvalidateErr error

func tryVTPixelTransferSessionInvalidate(session VTPixelTransferSessionRef) error {
	if _vTPixelTransferSessionInvalidate == nil {
		return symbolCallError("VTPixelTransferSessionInvalidate", "10.8", _vTPixelTransferSessionInvalidateErr)
	}
	_vTPixelTransferSessionInvalidate(session)
	return nil
}

// VTPixelTransferSessionInvalidate tears down a pixel transfer session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTPixelTransferSessionInvalidate(_:)
func VTPixelTransferSessionInvalidate(session VTPixelTransferSessionRef) {
	if callErr := tryVTPixelTransferSessionInvalidate(session); callErr != nil {
		panic(callErr)
	}
}

var _vTPixelTransferSessionTransferImage func(session VTPixelTransferSessionRef, sourceBuffer corevideo.CVPixelBufferRef, destinationBuffer corevideo.CVPixelBufferRef) int32
var _vTPixelTransferSessionTransferImageErr error

func tryVTPixelTransferSessionTransferImage(session VTPixelTransferSessionRef, sourceBuffer corevideo.CVPixelBufferRef, destinationBuffer corevideo.CVPixelBufferRef) (int32, error) {
	if _vTPixelTransferSessionTransferImage == nil {
		return 0, symbolCallError("VTPixelTransferSessionTransferImage", "10.8", _vTPixelTransferSessionTransferImageErr)
	}
	return _vTPixelTransferSessionTransferImage(session, sourceBuffer, destinationBuffer), nil
}

// VTPixelTransferSessionTransferImage copies and/or converts an image from one pixel buffer to another.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTPixelTransferSessionTransferImage(_:from:to:)
func VTPixelTransferSessionTransferImage(session VTPixelTransferSessionRef, sourceBuffer corevideo.CVPixelBufferRef, destinationBuffer corevideo.CVPixelBufferRef) int32 {
	result, callErr := tryVTPixelTransferSessionTransferImage(session, sourceBuffer, destinationBuffer)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTRAWProcessingSessionCompleteFrames func(session VTRAWProcessingSessionRef) int32
var _vTRAWProcessingSessionCompleteFramesErr error

func tryVTRAWProcessingSessionCompleteFrames(session VTRAWProcessingSessionRef) (int32, error) {
	if _vTRAWProcessingSessionCompleteFrames == nil {
		return 0, symbolCallError("VTRAWProcessingSessionCompleteFrames", "15.0", _vTRAWProcessingSessionCompleteFramesErr)
	}
	return _vTRAWProcessingSessionCompleteFrames(session), nil
}

// VTRAWProcessingSessionCompleteFrames forces the RAW Processor to complete processing frames.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTRAWProcessingSessionCompleteFrames
func VTRAWProcessingSessionCompleteFrames(session VTRAWProcessingSessionRef) int32 {
	result, callErr := tryVTRAWProcessingSessionCompleteFrames(session)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTRAWProcessingSessionCopyProcessingParameters func(session VTRAWProcessingSessionRef, outParameterArray *corefoundation.CFArrayRef) int32
var _vTRAWProcessingSessionCopyProcessingParametersErr error

func tryVTRAWProcessingSessionCopyProcessingParameters(session VTRAWProcessingSessionRef, outParameterArray *corefoundation.CFArrayRef) (int32, error) {
	if _vTRAWProcessingSessionCopyProcessingParameters == nil {
		return 0, symbolCallError("VTRAWProcessingSessionCopyProcessingParameters", "15.0", _vTRAWProcessingSessionCopyProcessingParametersErr)
	}
	return _vTRAWProcessingSessionCopyProcessingParameters(session, outParameterArray), nil
}

// VTRAWProcessingSessionCopyProcessingParameters copies an array of dictionaries describing the parameters provided by the RAW Processor for frame processing.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTRAWProcessingSessionCopyProcessingParameters
func VTRAWProcessingSessionCopyProcessingParameters(session VTRAWProcessingSessionRef, outParameterArray *corefoundation.CFArrayRef) int32 {
	result, callErr := tryVTRAWProcessingSessionCopyProcessingParameters(session, outParameterArray)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTRAWProcessingSessionCreate func(allocator corefoundation.CFAllocatorRef, formatDescription uintptr, outputPixelBufferAttributes corefoundation.CFDictionaryRef, processingSessionOptions corefoundation.CFDictionaryRef, processingSessionOut *VTRAWProcessingSessionRef) int32
var _vTRAWProcessingSessionCreateErr error

func tryVTRAWProcessingSessionCreate(allocator corefoundation.CFAllocatorRef, formatDescription uintptr, outputPixelBufferAttributes corefoundation.CFDictionaryRef, processingSessionOptions corefoundation.CFDictionaryRef, processingSessionOut *VTRAWProcessingSessionRef) (int32, error) {
	if _vTRAWProcessingSessionCreate == nil {
		return 0, symbolCallError("VTRAWProcessingSessionCreate", "15.0", _vTRAWProcessingSessionCreateErr)
	}
	return _vTRAWProcessingSessionCreate(allocator, formatDescription, outputPixelBufferAttributes, processingSessionOptions, processingSessionOut), nil
}

// VTRAWProcessingSessionCreate creates a RAW video frame processing session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTRAWProcessingSessionCreate
func VTRAWProcessingSessionCreate(allocator corefoundation.CFAllocatorRef, formatDescription uintptr, outputPixelBufferAttributes corefoundation.CFDictionaryRef, processingSessionOptions corefoundation.CFDictionaryRef, processingSessionOut *VTRAWProcessingSessionRef) int32 {
	result, callErr := tryVTRAWProcessingSessionCreate(allocator, formatDescription, outputPixelBufferAttributes, processingSessionOptions, processingSessionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTRAWProcessingSessionGetTypeID func() uint
var _vTRAWProcessingSessionGetTypeIDErr error

func tryVTRAWProcessingSessionGetTypeID() (uint, error) {
	if _vTRAWProcessingSessionGetTypeID == nil {
		return 0, symbolCallError("VTRAWProcessingSessionGetTypeID", "15.0", _vTRAWProcessingSessionGetTypeIDErr)
	}
	return _vTRAWProcessingSessionGetTypeID(), nil
}

// VTRAWProcessingSessionGetTypeID returns the type identifier for a RAW processing session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTRAWProcessingSessionGetTypeID
func VTRAWProcessingSessionGetTypeID() uint {
	result, callErr := tryVTRAWProcessingSessionGetTypeID()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTRAWProcessingSessionInvalidate func(session VTRAWProcessingSessionRef)
var _vTRAWProcessingSessionInvalidateErr error

func tryVTRAWProcessingSessionInvalidate(session VTRAWProcessingSessionRef) error {
	if _vTRAWProcessingSessionInvalidate == nil {
		return symbolCallError("VTRAWProcessingSessionInvalidate", "15.0", _vTRAWProcessingSessionInvalidateErr)
	}
	_vTRAWProcessingSessionInvalidate(session)
	return nil
}

// VTRAWProcessingSessionInvalidate tears down a RAW processing session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTRAWProcessingSessionInvalidate
func VTRAWProcessingSessionInvalidate(session VTRAWProcessingSessionRef) {
	if callErr := tryVTRAWProcessingSessionInvalidate(session); callErr != nil {
		panic(callErr)
	}
}

var _vTRAWProcessingSessionProcessFrame func(session VTRAWProcessingSessionRef, inputPixelBuffer corevideo.CVPixelBufferRef, frameOptions corefoundation.CFDictionaryRef, outputHandler unsafe.Pointer) int32
var _vTRAWProcessingSessionProcessFrameErr error

func tryVTRAWProcessingSessionProcessFrame(session VTRAWProcessingSessionRef, inputPixelBuffer corevideo.CVPixelBufferRef, frameOptions corefoundation.CFDictionaryRef, outputHandler VTRAWProcessingOutputHandler) (int32, error) {
	if _vTRAWProcessingSessionProcessFrame == nil {
		return 0, symbolCallError("VTRAWProcessingSessionProcessFrame", "15.0", _vTRAWProcessingSessionProcessFrameErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 int32, blockArg1 corevideo.CVImageBufferRef) {
		outputHandler(blockArg0, blockArg1)
	})
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _vTRAWProcessingSessionProcessFrame(session, inputPixelBuffer, frameOptions, _block0), nil
}

// VTRAWProcessingSessionProcessFrame submits RAW frames for format-specific processing using sequence and frame level parameters.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTRAWProcessingSessionProcessFrame
func VTRAWProcessingSessionProcessFrame(session VTRAWProcessingSessionRef, inputPixelBuffer corevideo.CVPixelBufferRef, frameOptions corefoundation.CFDictionaryRef, outputHandler VTRAWProcessingOutputHandler) int32 {
	result, callErr := tryVTRAWProcessingSessionProcessFrame(session, inputPixelBuffer, frameOptions, outputHandler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTRAWProcessingSessionSetParameterChangedHander func(session VTRAWProcessingSessionRef, parameterChangeHandler unsafe.Pointer) int32
var _vTRAWProcessingSessionSetParameterChangedHanderErr error

func tryVTRAWProcessingSessionSetParameterChangedHander(session VTRAWProcessingSessionRef, parameterChangeHandler VTRAWProcessingParameterChangeHandler) (int32, error) {
	if _vTRAWProcessingSessionSetParameterChangedHander == nil {
		return 0, symbolCallError("VTRAWProcessingSessionSetParameterChangedHander", "15.0", _vTRAWProcessingSessionSetParameterChangedHanderErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 corefoundation.CFArrayRef) { parameterChangeHandler(blockArg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _vTRAWProcessingSessionSetParameterChangedHander(session, _block0), nil
}

// VTRAWProcessingSessionSetParameterChangedHander provides a block which will be called when the session changes the set of processing parameters.
//
// Deprecated: Deprecated since macOS 26.0.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTRAWProcessingSessionSetParameterChangedHander
func VTRAWProcessingSessionSetParameterChangedHander(session VTRAWProcessingSessionRef, parameterChangeHandler VTRAWProcessingParameterChangeHandler) int32 {
	result, callErr := tryVTRAWProcessingSessionSetParameterChangedHander(session, parameterChangeHandler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTRAWProcessingSessionSetParameterChangedHandler func(session VTRAWProcessingSessionRef, parameterChangeHandler unsafe.Pointer) int32
var _vTRAWProcessingSessionSetParameterChangedHandlerErr error

func tryVTRAWProcessingSessionSetParameterChangedHandler(session VTRAWProcessingSessionRef, parameterChangeHandler VTRAWProcessingParameterChangeHandler) (int32, error) {
	if _vTRAWProcessingSessionSetParameterChangedHandler == nil {
		return 0, symbolCallError("VTRAWProcessingSessionSetParameterChangedHandler", "26.0", _vTRAWProcessingSessionSetParameterChangedHandlerErr)
	}
	_block0Value := objc.NewBlock(func(_ objc.Block, blockArg0 corefoundation.CFArrayRef) { parameterChangeHandler(blockArg0) })
	defer _block0Value.Release()
	_block0 := unsafe.Pointer(_block0Value)
	return _vTRAWProcessingSessionSetParameterChangedHandler(session, _block0), nil
}

// VTRAWProcessingSessionSetParameterChangedHandler.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTRAWProcessingSessionSetParameterChangedHandler
func VTRAWProcessingSessionSetParameterChangedHandler(session VTRAWProcessingSessionRef, parameterChangeHandler VTRAWProcessingParameterChangeHandler) int32 {
	result, callErr := tryVTRAWProcessingSessionSetParameterChangedHandler(session, parameterChangeHandler)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTRAWProcessingSessionSetProcessingParameters func(session VTRAWProcessingSessionRef, processingParameters corefoundation.CFDictionaryRef) int32
var _vTRAWProcessingSessionSetProcessingParametersErr error

func tryVTRAWProcessingSessionSetProcessingParameters(session VTRAWProcessingSessionRef, processingParameters corefoundation.CFDictionaryRef) (int32, error) {
	if _vTRAWProcessingSessionSetProcessingParameters == nil {
		return 0, symbolCallError("VTRAWProcessingSessionSetProcessingParameters", "15.0", _vTRAWProcessingSessionSetProcessingParametersErr)
	}
	return _vTRAWProcessingSessionSetProcessingParameters(session, processingParameters), nil
}

// VTRAWProcessingSessionSetProcessingParameters sets a collection of RAW Processing parameters.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTRAWProcessingSessionSetProcessingParameters
func VTRAWProcessingSessionSetProcessingParameters(session VTRAWProcessingSessionRef, processingParameters corefoundation.CFDictionaryRef) int32 {
	result, callErr := tryVTRAWProcessingSessionSetProcessingParameters(session, processingParameters)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTRegisterProfessionalVideoWorkflowVideoDecoders func()
var _vTRegisterProfessionalVideoWorkflowVideoDecodersErr error

func tryVTRegisterProfessionalVideoWorkflowVideoDecoders() error {
	if _vTRegisterProfessionalVideoWorkflowVideoDecoders == nil {
		return symbolCallError("VTRegisterProfessionalVideoWorkflowVideoDecoders", "10.9", _vTRegisterProfessionalVideoWorkflowVideoDecodersErr)
	}
	_vTRegisterProfessionalVideoWorkflowVideoDecoders()
	return nil
}

// VTRegisterProfessionalVideoWorkflowVideoDecoders loads decoders appropriate for the client’s professional video workflows.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTRegisterProfessionalVideoWorkflowVideoDecoders()
func VTRegisterProfessionalVideoWorkflowVideoDecoders() {
	if callErr := tryVTRegisterProfessionalVideoWorkflowVideoDecoders(); callErr != nil {
		panic(callErr)
	}
}

var _vTRegisterProfessionalVideoWorkflowVideoEncoders func()
var _vTRegisterProfessionalVideoWorkflowVideoEncodersErr error

func tryVTRegisterProfessionalVideoWorkflowVideoEncoders() error {
	if _vTRegisterProfessionalVideoWorkflowVideoEncoders == nil {
		return symbolCallError("VTRegisterProfessionalVideoWorkflowVideoEncoders", "10.10", _vTRegisterProfessionalVideoWorkflowVideoEncodersErr)
	}
	_vTRegisterProfessionalVideoWorkflowVideoEncoders()
	return nil
}

// VTRegisterProfessionalVideoWorkflowVideoEncoders loads encoders appropriate for the client’s professional video workflows.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTRegisterProfessionalVideoWorkflowVideoEncoders()
func VTRegisterProfessionalVideoWorkflowVideoEncoders() {
	if callErr := tryVTRegisterProfessionalVideoWorkflowVideoEncoders(); callErr != nil {
		panic(callErr)
	}
}

var _vTRegisterSupplementalVideoDecoderIfAvailable func(codecType uint32)
var _vTRegisterSupplementalVideoDecoderIfAvailableErr error

func tryVTRegisterSupplementalVideoDecoderIfAvailable(codecType uint32) error {
	if _vTRegisterSupplementalVideoDecoderIfAvailable == nil {
		return symbolCallError("VTRegisterSupplementalVideoDecoderIfAvailable", "11.0", _vTRegisterSupplementalVideoDecoderIfAvailableErr)
	}
	_vTRegisterSupplementalVideoDecoderIfAvailable(codecType)
	return nil
}

// VTRegisterSupplementalVideoDecoderIfAvailable registers a video decoder for the specified codec type, if one exists on the current system.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTRegisterSupplementalVideoDecoderIfAvailable(_:)
func VTRegisterSupplementalVideoDecoderIfAvailable(codecType uint32) {
	if callErr := tryVTRegisterSupplementalVideoDecoderIfAvailable(codecType); callErr != nil {
		panic(callErr)
	}
}

var _vTSessionCopyProperty func(session VTSessionRef, propertyKey corefoundation.CFStringRef, allocator corefoundation.CFAllocatorRef, propertyValueOut unsafe.Pointer) int32
var _vTSessionCopyPropertyErr error

func tryVTSessionCopyProperty(session VTSessionRef, propertyKey corefoundation.CFStringRef, allocator corefoundation.CFAllocatorRef, propertyValueOut unsafe.Pointer) (int32, error) {
	if _vTSessionCopyProperty == nil {
		return 0, symbolCallError("VTSessionCopyProperty", "10.8", _vTSessionCopyPropertyErr)
	}
	return _vTSessionCopyProperty(session, propertyKey, allocator, propertyValueOut), nil
}

// VTSessionCopyProperty retrieves a property on a Video Toolbox session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSessionCopyProperty(_:key:allocator:valueOut:)
func VTSessionCopyProperty(session VTSessionRef, propertyKey corefoundation.CFStringRef, allocator corefoundation.CFAllocatorRef, propertyValueOut unsafe.Pointer) int32 {
	result, callErr := tryVTSessionCopyProperty(session, propertyKey, allocator, propertyValueOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTSessionCopySerializableProperties func(session VTSessionRef, allocator corefoundation.CFAllocatorRef, dictionaryOut *corefoundation.CFDictionaryRef) int32
var _vTSessionCopySerializablePropertiesErr error

func tryVTSessionCopySerializableProperties(session VTSessionRef, allocator corefoundation.CFAllocatorRef, dictionaryOut *corefoundation.CFDictionaryRef) (int32, error) {
	if _vTSessionCopySerializableProperties == nil {
		return 0, symbolCallError("VTSessionCopySerializableProperties", "10.8", _vTSessionCopySerializablePropertiesErr)
	}
	return _vTSessionCopySerializableProperties(session, allocator, dictionaryOut), nil
}

// VTSessionCopySerializableProperties retrieves the set of serializable property keys and their current values.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSessionCopySerializableProperties(_:allocator:dictionaryOut:)
func VTSessionCopySerializableProperties(session VTSessionRef, allocator corefoundation.CFAllocatorRef, dictionaryOut *corefoundation.CFDictionaryRef) int32 {
	result, callErr := tryVTSessionCopySerializableProperties(session, allocator, dictionaryOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTSessionCopySupportedPropertyDictionary func(session VTSessionRef, supportedPropertyDictionaryOut *corefoundation.CFDictionaryRef) int32
var _vTSessionCopySupportedPropertyDictionaryErr error

func tryVTSessionCopySupportedPropertyDictionary(session VTSessionRef, supportedPropertyDictionaryOut *corefoundation.CFDictionaryRef) (int32, error) {
	if _vTSessionCopySupportedPropertyDictionary == nil {
		return 0, symbolCallError("VTSessionCopySupportedPropertyDictionary", "10.8", _vTSessionCopySupportedPropertyDictionaryErr)
	}
	return _vTSessionCopySupportedPropertyDictionary(session, supportedPropertyDictionaryOut), nil
}

// VTSessionCopySupportedPropertyDictionary retrieves a dictionary enumerating all the supported properties of a video toolbox session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSessionCopySupportedPropertyDictionary(_:supportedPropertyDictionaryOut:)
func VTSessionCopySupportedPropertyDictionary(session VTSessionRef, supportedPropertyDictionaryOut *corefoundation.CFDictionaryRef) int32 {
	result, callErr := tryVTSessionCopySupportedPropertyDictionary(session, supportedPropertyDictionaryOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTSessionSetProperties func(session VTSessionRef, propertyDictionary corefoundation.CFDictionaryRef) int32
var _vTSessionSetPropertiesErr error

func tryVTSessionSetProperties(session VTSessionRef, propertyDictionary corefoundation.CFDictionaryRef) (int32, error) {
	if _vTSessionSetProperties == nil {
		return 0, symbolCallError("VTSessionSetProperties", "10.8", _vTSessionSetPropertiesErr)
	}
	return _vTSessionSetProperties(session, propertyDictionary), nil
}

// VTSessionSetProperties sets multiple properties at once.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSessionSetProperties(_:propertyDictionary:)
func VTSessionSetProperties(session VTSessionRef, propertyDictionary corefoundation.CFDictionaryRef) int32 {
	result, callErr := tryVTSessionSetProperties(session, propertyDictionary)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _vTSessionSetProperty func(session VTSessionRef, propertyKey corefoundation.CFStringRef, propertyValue corefoundation.CFTypeRef) int32
var _vTSessionSetPropertyErr error

func tryVTSessionSetProperty(session VTSessionRef, propertyKey corefoundation.CFStringRef, propertyValue corefoundation.CFTypeRef) (int32, error) {
	if _vTSessionSetProperty == nil {
		return 0, symbolCallError("VTSessionSetProperty", "10.8", _vTSessionSetPropertyErr)
	}
	return _vTSessionSetProperty(session, propertyKey, propertyValue), nil
}

// VTSessionSetProperty sets a property on a VideoToolbox session.
//
// See: https://developer.apple.com/documentation/VideoToolbox/VTSessionSetProperty(_:key:value:)
func VTSessionSetProperty(session VTSessionRef, propertyKey corefoundation.CFStringRef, propertyValue corefoundation.CFTypeRef) int32 {
	result, callErr := tryVTSessionSetProperty(session, propertyKey, propertyValue)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_vTCompressionSessionBeginPass, &_vTCompressionSessionBeginPassErr, frameworkHandle, "VTCompressionSessionBeginPass", "10.10")
	registerFunc(&_vTCompressionSessionCompleteFrames, &_vTCompressionSessionCompleteFramesErr, frameworkHandle, "VTCompressionSessionCompleteFrames", "10.8")
	registerFunc(&_vTCompressionSessionCreate, &_vTCompressionSessionCreateErr, frameworkHandle, "VTCompressionSessionCreate", "10.8")
	registerFunc(&_vTCompressionSessionEncodeFrame, &_vTCompressionSessionEncodeFrameErr, frameworkHandle, "VTCompressionSessionEncodeFrame", "10.8")
	registerFunc(&_vTCompressionSessionEncodeFrameWithOutputHandler, &_vTCompressionSessionEncodeFrameWithOutputHandlerErr, frameworkHandle, "VTCompressionSessionEncodeFrameWithOutputHandler", "10.11")
	registerFunc(&_vTCompressionSessionEncodeMultiImageFrame, &_vTCompressionSessionEncodeMultiImageFrameErr, frameworkHandle, "VTCompressionSessionEncodeMultiImageFrame", "14.0")
	registerFunc(&_vTCompressionSessionEncodeMultiImageFrameWithOutputHandler, &_vTCompressionSessionEncodeMultiImageFrameWithOutputHandlerErr, frameworkHandle, "VTCompressionSessionEncodeMultiImageFrameWithOutputHandler", "14.0")
	registerFunc(&_vTCompressionSessionEndPass, &_vTCompressionSessionEndPassErr, frameworkHandle, "VTCompressionSessionEndPass", "10.10")
	registerFunc(&_vTCompressionSessionGetPixelBufferPool, &_vTCompressionSessionGetPixelBufferPoolErr, frameworkHandle, "VTCompressionSessionGetPixelBufferPool", "10.8")
	registerFunc(&_vTCompressionSessionGetTimeRangesForNextPass, &_vTCompressionSessionGetTimeRangesForNextPassErr, frameworkHandle, "VTCompressionSessionGetTimeRangesForNextPass", "10.10")
	registerFunc(&_vTCompressionSessionGetTypeID, &_vTCompressionSessionGetTypeIDErr, frameworkHandle, "VTCompressionSessionGetTypeID", "10.8")
	registerFunc(&_vTCompressionSessionInvalidate, &_vTCompressionSessionInvalidateErr, frameworkHandle, "VTCompressionSessionInvalidate", "10.8")
	registerFunc(&_vTCompressionSessionPrepareToEncodeFrames, &_vTCompressionSessionPrepareToEncodeFramesErr, frameworkHandle, "VTCompressionSessionPrepareToEncodeFrames", "10.9")
	registerFunc(&_vTCopyRAWProcessorExtensionProperties, &_vTCopyRAWProcessorExtensionPropertiesErr, frameworkHandle, "VTCopyRAWProcessorExtensionProperties", "15.0")
	registerFunc(&_vTCopySupportedPropertyDictionaryForEncoder, &_vTCopySupportedPropertyDictionaryForEncoderErr, frameworkHandle, "VTCopySupportedPropertyDictionaryForEncoder", "10.13")
	registerFunc(&_vTCopyVideoDecoderExtensionProperties, &_vTCopyVideoDecoderExtensionPropertiesErr, frameworkHandle, "VTCopyVideoDecoderExtensionProperties", "15.0")
	registerFunc(&_vTCopyVideoEncoderList, &_vTCopyVideoEncoderListErr, frameworkHandle, "VTCopyVideoEncoderList", "10.8")
	registerFunc(&_vTCreateCGImageFromCVPixelBuffer, &_vTCreateCGImageFromCVPixelBufferErr, frameworkHandle, "VTCreateCGImageFromCVPixelBuffer", "10.11")
	registerFunc(&_vTDecompressionSessionCanAcceptFormatDescription, &_vTDecompressionSessionCanAcceptFormatDescriptionErr, frameworkHandle, "VTDecompressionSessionCanAcceptFormatDescription", "10.8")
	registerFunc(&_vTDecompressionSessionCopyBlackPixelBuffer, &_vTDecompressionSessionCopyBlackPixelBufferErr, frameworkHandle, "VTDecompressionSessionCopyBlackPixelBuffer", "10.8")
	registerFunc(&_vTDecompressionSessionCreate, &_vTDecompressionSessionCreateErr, frameworkHandle, "VTDecompressionSessionCreate", "10.8")
	registerFunc(&_vTDecompressionSessionDecodeFrame, &_vTDecompressionSessionDecodeFrameErr, frameworkHandle, "VTDecompressionSessionDecodeFrame", "10.8")
	registerFunc(&_vTDecompressionSessionDecodeFrameWithMultiImageCapableOutputHandler, &_vTDecompressionSessionDecodeFrameWithMultiImageCapableOutputHandlerErr, frameworkHandle, "VTDecompressionSessionDecodeFrameWithMultiImageCapableOutputHandler", "14.0")
	registerFunc(&_vTDecompressionSessionDecodeFrameWithOptions, &_vTDecompressionSessionDecodeFrameWithOptionsErr, frameworkHandle, "VTDecompressionSessionDecodeFrameWithOptions", "15.0")
	registerFunc(&_vTDecompressionSessionDecodeFrameWithOptionsAndOutputHandler, &_vTDecompressionSessionDecodeFrameWithOptionsAndOutputHandlerErr, frameworkHandle, "VTDecompressionSessionDecodeFrameWithOptionsAndOutputHandler", "15.0")
	registerFunc(&_vTDecompressionSessionDecodeFrameWithOutputHandler, &_vTDecompressionSessionDecodeFrameWithOutputHandlerErr, frameworkHandle, "VTDecompressionSessionDecodeFrameWithOutputHandler", "10.11")
	registerFunc(&_vTDecompressionSessionFinishDelayedFrames, &_vTDecompressionSessionFinishDelayedFramesErr, frameworkHandle, "VTDecompressionSessionFinishDelayedFrames", "10.8")
	registerFunc(&_vTDecompressionSessionGetTypeID, &_vTDecompressionSessionGetTypeIDErr, frameworkHandle, "VTDecompressionSessionGetTypeID", "10.8")
	registerFunc(&_vTDecompressionSessionInvalidate, &_vTDecompressionSessionInvalidateErr, frameworkHandle, "VTDecompressionSessionInvalidate", "10.8")
	registerFunc(&_vTDecompressionSessionSetMultiImageCallback, &_vTDecompressionSessionSetMultiImageCallbackErr, frameworkHandle, "VTDecompressionSessionSetMultiImageCallback", "14.0")
	registerFunc(&_vTDecompressionSessionWaitForAsynchronousFrames, &_vTDecompressionSessionWaitForAsynchronousFramesErr, frameworkHandle, "VTDecompressionSessionWaitForAsynchronousFrames", "10.8")
	registerFunc(&_vTFrameSiloAddSampleBuffer, &_vTFrameSiloAddSampleBufferErr, frameworkHandle, "VTFrameSiloAddSampleBuffer", "10.10")
	registerFunc(&_vTFrameSiloCallBlockForEachSampleBuffer, &_vTFrameSiloCallBlockForEachSampleBufferErr, frameworkHandle, "VTFrameSiloCallBlockForEachSampleBuffer", "10.10")
	registerFunc(&_vTFrameSiloCallFunctionForEachSampleBuffer, &_vTFrameSiloCallFunctionForEachSampleBufferErr, frameworkHandle, "VTFrameSiloCallFunctionForEachSampleBuffer", "10.10")
	registerFunc(&_vTFrameSiloCreate, &_vTFrameSiloCreateErr, frameworkHandle, "VTFrameSiloCreate", "10.10")
	registerFunc(&_vTFrameSiloGetProgressOfCurrentPass, &_vTFrameSiloGetProgressOfCurrentPassErr, frameworkHandle, "VTFrameSiloGetProgressOfCurrentPass", "10.10")
	registerFunc(&_vTFrameSiloGetTypeID, &_vTFrameSiloGetTypeIDErr, frameworkHandle, "VTFrameSiloGetTypeID", "10.10")
	registerFunc(&_vTFrameSiloSetTimeRangesForNextPass, &_vTFrameSiloSetTimeRangesForNextPassErr, frameworkHandle, "VTFrameSiloSetTimeRangesForNextPass", "10.10")
	registerFunc(&_vTHDRPerFrameMetadataGenerationSessionAttachMetadata, &_vTHDRPerFrameMetadataGenerationSessionAttachMetadataErr, frameworkHandle, "VTHDRPerFrameMetadataGenerationSessionAttachMetadata", "15.0")
	registerFunc(&_vTHDRPerFrameMetadataGenerationSessionCreate, &_vTHDRPerFrameMetadataGenerationSessionCreateErr, frameworkHandle, "VTHDRPerFrameMetadataGenerationSessionCreate", "15.0")
	registerFunc(&_vTHDRPerFrameMetadataGenerationSessionGetTypeID, &_vTHDRPerFrameMetadataGenerationSessionGetTypeIDErr, frameworkHandle, "VTHDRPerFrameMetadataGenerationSessionGetTypeID", "15.0")
	registerFunc(&_vTIsHardwareDecodeSupported, &_vTIsHardwareDecodeSupportedErr, frameworkHandle, "VTIsHardwareDecodeSupported", "10.13")
	registerFunc(&_vTIsStereoMVHEVCDecodeSupported, &_vTIsStereoMVHEVCDecodeSupportedErr, frameworkHandle, "VTIsStereoMVHEVCDecodeSupported", "14.0")
	registerFunc(&_vTIsStereoMVHEVCEncodeSupported, &_vTIsStereoMVHEVCEncodeSupportedErr, frameworkHandle, "VTIsStereoMVHEVCEncodeSupported", "14.0")
	registerFunc(&_vTMotionEstimationSessionCompleteFrames, &_vTMotionEstimationSessionCompleteFramesErr, frameworkHandle, "VTMotionEstimationSessionCompleteFrames", "26.0")
	registerFunc(&_vTMotionEstimationSessionCopySourcePixelBufferAttributes, &_vTMotionEstimationSessionCopySourcePixelBufferAttributesErr, frameworkHandle, "VTMotionEstimationSessionCopySourcePixelBufferAttributes", "26.0")
	registerFunc(&_vTMotionEstimationSessionCreate, &_vTMotionEstimationSessionCreateErr, frameworkHandle, "VTMotionEstimationSessionCreate", "26.0")
	registerFunc(&_vTMotionEstimationSessionEstimateMotionVectors, &_vTMotionEstimationSessionEstimateMotionVectorsErr, frameworkHandle, "VTMotionEstimationSessionEstimateMotionVectors", "26.0")
	registerFunc(&_vTMotionEstimationSessionGetTypeID, &_vTMotionEstimationSessionGetTypeIDErr, frameworkHandle, "VTMotionEstimationSessionGetTypeID", "26.0")
	registerFunc(&_vTMotionEstimationSessionInvalidate, &_vTMotionEstimationSessionInvalidateErr, frameworkHandle, "VTMotionEstimationSessionInvalidate", "26.0")
	registerFunc(&_vTMultiPassStorageClose, &_vTMultiPassStorageCloseErr, frameworkHandle, "VTMultiPassStorageClose", "10.10")
	registerFunc(&_vTMultiPassStorageCreate, &_vTMultiPassStorageCreateErr, frameworkHandle, "VTMultiPassStorageCreate", "10.10")
	registerFunc(&_vTMultiPassStorageGetTypeID, &_vTMultiPassStorageGetTypeIDErr, frameworkHandle, "VTMultiPassStorageGetTypeID", "10.10")
	registerFunc(&_vTPixelRotationSessionCreate, &_vTPixelRotationSessionCreateErr, frameworkHandle, "VTPixelRotationSessionCreate", "13.0")
	registerFunc(&_vTPixelRotationSessionGetTypeID, &_vTPixelRotationSessionGetTypeIDErr, frameworkHandle, "VTPixelRotationSessionGetTypeID", "13.0")
	registerFunc(&_vTPixelRotationSessionInvalidate, &_vTPixelRotationSessionInvalidateErr, frameworkHandle, "VTPixelRotationSessionInvalidate", "13.0")
	registerFunc(&_vTPixelRotationSessionRotateImage, &_vTPixelRotationSessionRotateImageErr, frameworkHandle, "VTPixelRotationSessionRotateImage", "13.0")
	registerFunc(&_vTPixelTransferSessionCreate, &_vTPixelTransferSessionCreateErr, frameworkHandle, "VTPixelTransferSessionCreate", "10.8")
	registerFunc(&_vTPixelTransferSessionGetTypeID, &_vTPixelTransferSessionGetTypeIDErr, frameworkHandle, "VTPixelTransferSessionGetTypeID", "10.8")
	registerFunc(&_vTPixelTransferSessionInvalidate, &_vTPixelTransferSessionInvalidateErr, frameworkHandle, "VTPixelTransferSessionInvalidate", "10.8")
	registerFunc(&_vTPixelTransferSessionTransferImage, &_vTPixelTransferSessionTransferImageErr, frameworkHandle, "VTPixelTransferSessionTransferImage", "10.8")
	registerFunc(&_vTRAWProcessingSessionCompleteFrames, &_vTRAWProcessingSessionCompleteFramesErr, frameworkHandle, "VTRAWProcessingSessionCompleteFrames", "15.0")
	registerFunc(&_vTRAWProcessingSessionCopyProcessingParameters, &_vTRAWProcessingSessionCopyProcessingParametersErr, frameworkHandle, "VTRAWProcessingSessionCopyProcessingParameters", "15.0")
	registerFunc(&_vTRAWProcessingSessionCreate, &_vTRAWProcessingSessionCreateErr, frameworkHandle, "VTRAWProcessingSessionCreate", "15.0")
	registerFunc(&_vTRAWProcessingSessionGetTypeID, &_vTRAWProcessingSessionGetTypeIDErr, frameworkHandle, "VTRAWProcessingSessionGetTypeID", "15.0")
	registerFunc(&_vTRAWProcessingSessionInvalidate, &_vTRAWProcessingSessionInvalidateErr, frameworkHandle, "VTRAWProcessingSessionInvalidate", "15.0")
	registerFunc(&_vTRAWProcessingSessionProcessFrame, &_vTRAWProcessingSessionProcessFrameErr, frameworkHandle, "VTRAWProcessingSessionProcessFrame", "15.0")
	registerFunc(&_vTRAWProcessingSessionSetParameterChangedHander, &_vTRAWProcessingSessionSetParameterChangedHanderErr, frameworkHandle, "VTRAWProcessingSessionSetParameterChangedHander", "15.0")
	registerFunc(&_vTRAWProcessingSessionSetParameterChangedHandler, &_vTRAWProcessingSessionSetParameterChangedHandlerErr, frameworkHandle, "VTRAWProcessingSessionSetParameterChangedHandler", "26.0")
	registerFunc(&_vTRAWProcessingSessionSetProcessingParameters, &_vTRAWProcessingSessionSetProcessingParametersErr, frameworkHandle, "VTRAWProcessingSessionSetProcessingParameters", "15.0")
	registerFunc(&_vTRegisterProfessionalVideoWorkflowVideoDecoders, &_vTRegisterProfessionalVideoWorkflowVideoDecodersErr, frameworkHandle, "VTRegisterProfessionalVideoWorkflowVideoDecoders", "10.9")
	registerFunc(&_vTRegisterProfessionalVideoWorkflowVideoEncoders, &_vTRegisterProfessionalVideoWorkflowVideoEncodersErr, frameworkHandle, "VTRegisterProfessionalVideoWorkflowVideoEncoders", "10.10")
	registerFunc(&_vTRegisterSupplementalVideoDecoderIfAvailable, &_vTRegisterSupplementalVideoDecoderIfAvailableErr, frameworkHandle, "VTRegisterSupplementalVideoDecoderIfAvailable", "11.0")
	registerFunc(&_vTSessionCopyProperty, &_vTSessionCopyPropertyErr, frameworkHandle, "VTSessionCopyProperty", "10.8")
	registerFunc(&_vTSessionCopySerializableProperties, &_vTSessionCopySerializablePropertiesErr, frameworkHandle, "VTSessionCopySerializableProperties", "10.8")
	registerFunc(&_vTSessionCopySupportedPropertyDictionary, &_vTSessionCopySupportedPropertyDictionaryErr, frameworkHandle, "VTSessionCopySupportedPropertyDictionary", "10.8")
	registerFunc(&_vTSessionSetProperties, &_vTSessionSetPropertiesErr, frameworkHandle, "VTSessionSetProperties", "10.8")
	registerFunc(&_vTSessionSetProperty, &_vTSessionSetPropertyErr, frameworkHandle, "VTSessionSetProperty", "10.8")
}
