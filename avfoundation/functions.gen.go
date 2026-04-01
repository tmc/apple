// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"fmt"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
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
		return fmt.Sprintf("AVFoundation: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("AVFoundation: symbol %s unavailable on this system", e.symbol)
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
		return fmt.Errorf("AVFoundation: symbol %s unavailable because the framework could not be loaded", name)
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
			*errDst = fmt.Errorf("AVFoundation: register symbol %s: %v", name, r)
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

var _aVCaptionDimensionMake func(value float64, units AVCaptionUnitsType) AVCaptionDimension
var _aVCaptionDimensionMakeErr error

func tryAVCaptionDimensionMake(value float64, units AVCaptionUnitsType) (AVCaptionDimension, error) {
	if _aVCaptionDimensionMake == nil {
		return AVCaptionDimension{}, symbolCallError("AVCaptionDimensionMake", "12.0", _aVCaptionDimensionMakeErr)
	}
	return _aVCaptionDimensionMake(value, units), nil
}

// AVCaptionDimensionMake creates a caption dimension with a value and unit type.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionDimensionMake
func AVCaptionDimensionMake(value float64, units AVCaptionUnitsType) AVCaptionDimension {
	result, callErr := tryAVCaptionDimensionMake(value, units)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aVCaptionPointMake func(x AVCaptionDimension, y AVCaptionDimension) AVCaptionPoint
var _aVCaptionPointMakeErr error

func tryAVCaptionPointMake(x AVCaptionDimension, y AVCaptionDimension) (AVCaptionPoint, error) {
	if _aVCaptionPointMake == nil {
		return AVCaptionPoint{}, symbolCallError("AVCaptionPointMake", "12.0", _aVCaptionPointMakeErr)
	}
	return _aVCaptionPointMake(x, y), nil
}

// AVCaptionPointMake creates a caption point with the specified x and y positions.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionPointMake
func AVCaptionPointMake(x AVCaptionDimension, y AVCaptionDimension) AVCaptionPoint {
	result, callErr := tryAVCaptionPointMake(x, y)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aVCaptionSizeMake func(width AVCaptionDimension, height AVCaptionDimension) AVCaptionSize
var _aVCaptionSizeMakeErr error

func tryAVCaptionSizeMake(width AVCaptionDimension, height AVCaptionDimension) (AVCaptionSize, error) {
	if _aVCaptionSizeMake == nil {
		return AVCaptionSize{}, symbolCallError("AVCaptionSizeMake", "12.0", _aVCaptionSizeMakeErr)
	}
	return _aVCaptionSizeMake(width, height), nil
}

// AVCaptionSizeMake creates a caption size with the specified width and height.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionSizeMake
func AVCaptionSizeMake(width AVCaptionDimension, height AVCaptionDimension) AVCaptionSize {
	result, callErr := tryAVCaptionSizeMake(width, height)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aVCaptureReactionSystemImageNameForType func(reactionType AVCaptureReactionType) foundation.NSString
var _aVCaptureReactionSystemImageNameForTypeErr error

func tryAVCaptureReactionSystemImageNameForType(reactionType AVCaptureReactionType) (foundation.NSString, error) {
	if _aVCaptureReactionSystemImageNameForType == nil {
		return foundation.NSString{}, symbolCallError("AVCaptureReactionSystemImageNameForType", "14.0", _aVCaptureReactionSystemImageNameForTypeErr)
	}
	return _aVCaptureReactionSystemImageNameForType(reactionType), nil
}

// AVCaptureReactionSystemImageNameForType returns the name of a system image that displays the recommended iconography for a specified reaction type.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureReactionType/systemImageName
func AVCaptureReactionSystemImageNameForType(reactionType AVCaptureReactionType) foundation.NSString {
	result, callErr := tryAVCaptureReactionSystemImageNameForType(reactionType)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aVCaptureTimecodeAdvancedByFrames func(timecode AVCaptureTimecode, framesToAdd int64) AVCaptureTimecode
var _aVCaptureTimecodeAdvancedByFramesErr error

func tryAVCaptureTimecodeAdvancedByFrames(timecode AVCaptureTimecode, framesToAdd int64) (AVCaptureTimecode, error) {
	if _aVCaptureTimecodeAdvancedByFrames == nil {
		return AVCaptureTimecode{}, symbolCallError("AVCaptureTimecodeAdvancedByFrames", "26.0", _aVCaptureTimecodeAdvancedByFramesErr)
	}
	return _aVCaptureTimecodeAdvancedByFrames(timecode, framesToAdd), nil
}

// AVCaptureTimecodeAdvancedByFrames generates a new timecode by adding a specified number of frames to the given timecode, handling overflow for seconds, minutes, and hours.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/advanced(_:by:)
func AVCaptureTimecodeAdvancedByFrames(timecode AVCaptureTimecode, framesToAdd int64) AVCaptureTimecode {
	result, callErr := tryAVCaptureTimecodeAdvancedByFrames(timecode, framesToAdd)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp func(timecode AVCaptureTimecode, presentationTimeStamp coremedia.CMTime) uintptr
var _aVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStampErr error

func tryAVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp(timecode AVCaptureTimecode, presentationTimeStamp coremedia.CMTime) (uintptr, error) {
	if _aVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp == nil {
		return 0, symbolCallError("AVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp", "26.0", _aVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStampErr)
	}
	return _aVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp(timecode, presentationTimeStamp), nil
}

// AVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp creates a sample buffer containing Timecode Media Description metadata for integration with a video track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/createMetadataSampleBuffer(from:associatedWithPresentationTimeStamp:)
func AVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp(timecode AVCaptureTimecode, presentationTimeStamp coremedia.CMTime) uintptr {
	result, callErr := tryAVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp(timecode, presentationTimeStamp)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aVCaptureTimecodeCreateMetadataSampleBufferForDuration func(timecode AVCaptureTimecode, duration coremedia.CMTime) uintptr
var _aVCaptureTimecodeCreateMetadataSampleBufferForDurationErr error

func tryAVCaptureTimecodeCreateMetadataSampleBufferForDuration(timecode AVCaptureTimecode, duration coremedia.CMTime) (uintptr, error) {
	if _aVCaptureTimecodeCreateMetadataSampleBufferForDuration == nil {
		return 0, symbolCallError("AVCaptureTimecodeCreateMetadataSampleBufferForDuration", "26.0", _aVCaptureTimecodeCreateMetadataSampleBufferForDurationErr)
	}
	return _aVCaptureTimecodeCreateMetadataSampleBufferForDuration(timecode, duration), nil
}

// AVCaptureTimecodeCreateMetadataSampleBufferForDuration creates a sample buffer containing Timecode Media Description metadata for a specified duration.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/createMetadataSampleBuffer(from:forDuration:)
func AVCaptureTimecodeCreateMetadataSampleBufferForDuration(timecode AVCaptureTimecode, duration coremedia.CMTime) uintptr {
	result, callErr := tryAVCaptureTimecodeCreateMetadataSampleBufferForDuration(timecode, duration)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aVMakeRectWithAspectRatioInsideRect func(aspectRatio corefoundation.CGSize, boundingRect corefoundation.CGRect) corefoundation.CGRect
var _aVMakeRectWithAspectRatioInsideRectErr error

func tryAVMakeRectWithAspectRatioInsideRect(aspectRatio corefoundation.CGSize, boundingRect corefoundation.CGRect) (corefoundation.CGRect, error) {
	if _aVMakeRectWithAspectRatioInsideRect == nil {
		return corefoundation.CGRect{}, symbolCallError("AVMakeRectWithAspectRatioInsideRect", "10.7", _aVMakeRectWithAspectRatioInsideRectErr)
	}
	return _aVMakeRectWithAspectRatioInsideRect(aspectRatio, boundingRect), nil
}

// AVMakeRectWithAspectRatioInsideRect returns a scaled rectangle that maintains the specified aspect ratio within a bounding rectangle.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMakeRect(aspectRatio:insideRect:)
func AVMakeRectWithAspectRatioInsideRect(aspectRatio corefoundation.CGSize, boundingRect corefoundation.CGRect) corefoundation.CGRect {
	result, callErr := tryAVMakeRectWithAspectRatioInsideRect(aspectRatio, boundingRect)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _aVSampleBufferAttachContentKey func(sbuf uintptr, contentKey AVContentKey, outError *foundation.NSError) bool
var _aVSampleBufferAttachContentKeyErr error

func tryAVSampleBufferAttachContentKey(sbuf uintptr, contentKey AVContentKey, outError *foundation.NSError) (bool, error) {
	if _aVSampleBufferAttachContentKey == nil {
		return false, symbolCallError("AVSampleBufferAttachContentKey", "10.10", _aVSampleBufferAttachContentKeyErr)
	}
	return _aVSampleBufferAttachContentKey(sbuf, contentKey, outError), nil
}

// AVSampleBufferAttachContentKey attaches a content key to a sample buffer for the purpose of content decryption.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAttachContentKey(_:_:_:)
func AVSampleBufferAttachContentKey(sbuf uintptr, contentKey AVContentKey, outError *foundation.NSError) bool {
	result, callErr := tryAVSampleBufferAttachContentKey(sbuf, contentKey, outError)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _cMTagCollectionCreateWithVideoOutputPreset func(allocator corefoundation.CFAllocatorRef, preset CMTagCollectionVideoOutputPreset, newCollectionOut *coremedia.CMTagCollectionRef) int32
var _cMTagCollectionCreateWithVideoOutputPresetErr error

func tryCMTagCollectionCreateWithVideoOutputPreset(allocator corefoundation.CFAllocatorRef, preset CMTagCollectionVideoOutputPreset, newCollectionOut *coremedia.CMTagCollectionRef) (int32, error) {
	if _cMTagCollectionCreateWithVideoOutputPreset == nil {
		return 0, symbolCallError("CMTagCollectionCreateWithVideoOutputPreset", "14.2", _cMTagCollectionCreateWithVideoOutputPresetErr)
	}
	return _cMTagCollectionCreateWithVideoOutputPreset(allocator, preset, newCollectionOut), nil
}

// CMTagCollectionCreateWithVideoOutputPreset creates a collection with the required tags to describe the specified video output requirements.
//
// See: https://developer.apple.com/documentation/AVFoundation/CMTagCollectionCreateWithVideoOutputPreset
func CMTagCollectionCreateWithVideoOutputPreset(allocator corefoundation.CFAllocatorRef, preset CMTagCollectionVideoOutputPreset, newCollectionOut *coremedia.CMTagCollectionRef) int32 {
	result, callErr := tryCMTagCollectionCreateWithVideoOutputPreset(allocator, preset, newCollectionOut)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_aVCaptionDimensionMake, &_aVCaptionDimensionMakeErr, frameworkHandle, "AVCaptionDimensionMake", "12.0")
	registerFunc(&_aVCaptionPointMake, &_aVCaptionPointMakeErr, frameworkHandle, "AVCaptionPointMake", "12.0")
	registerFunc(&_aVCaptionSizeMake, &_aVCaptionSizeMakeErr, frameworkHandle, "AVCaptionSizeMake", "12.0")
	registerFunc(&_aVCaptureReactionSystemImageNameForType, &_aVCaptureReactionSystemImageNameForTypeErr, frameworkHandle, "AVCaptureReactionSystemImageNameForType", "14.0")
	registerFunc(&_aVCaptureTimecodeAdvancedByFrames, &_aVCaptureTimecodeAdvancedByFramesErr, frameworkHandle, "AVCaptureTimecodeAdvancedByFrames", "26.0")
	registerFunc(&_aVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp, &_aVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStampErr, frameworkHandle, "AVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp", "26.0")
	registerFunc(&_aVCaptureTimecodeCreateMetadataSampleBufferForDuration, &_aVCaptureTimecodeCreateMetadataSampleBufferForDurationErr, frameworkHandle, "AVCaptureTimecodeCreateMetadataSampleBufferForDuration", "26.0")
	registerFunc(&_aVMakeRectWithAspectRatioInsideRect, &_aVMakeRectWithAspectRatioInsideRectErr, frameworkHandle, "AVMakeRectWithAspectRatioInsideRect", "10.7")
	registerFunc(&_aVSampleBufferAttachContentKey, &_aVSampleBufferAttachContentKeyErr, frameworkHandle, "AVSampleBufferAttachContentKey", "10.10")
	registerFunc(&_cMTagCollectionCreateWithVideoOutputPreset, &_cMTagCollectionCreateWithVideoOutputPresetErr, frameworkHandle, "CMTagCollectionCreateWithVideoOutputPreset", "14.2")
}
