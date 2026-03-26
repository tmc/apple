// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// registerFunc resolves a framework symbol and registers it as a Go function.
// If the symbol is not found, a warning is printed and the function pointer is left nil.
func registerFunc(fptr any, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: AVFoundation: symbol %s not found\n", name)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "warning: AVFoundation: symbol %s registration skipped: %v\n", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, handle uintptr, name string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: AVFoundation: symbol %s not found\n", name)
		return
	}
	*dst = sym
}

var _aVCaptionDimensionMake func(value float64, units AVCaptionUnitsType) AVCaptionDimension

// AVCaptionDimensionMake creates a caption dimension with a value and unit type.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionDimensionMake
func AVCaptionDimensionMake(value float64, units AVCaptionUnitsType) AVCaptionDimension {
	if _aVCaptionDimensionMake == nil {
		panic("AVFoundation: symbol AVCaptionDimensionMake not loaded")
	}
	return _aVCaptionDimensionMake(value, units)
}

var _aVCaptionPointMake func(x AVCaptionDimension, y AVCaptionDimension) AVCaptionPoint

// AVCaptionPointMake creates a caption point with the specified x and y positions.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionPointMake
func AVCaptionPointMake(x AVCaptionDimension, y AVCaptionDimension) AVCaptionPoint {
	if _aVCaptionPointMake == nil {
		panic("AVFoundation: symbol AVCaptionPointMake not loaded")
	}
	return _aVCaptionPointMake(x, y)
}

var _aVCaptionSizeMake func(width AVCaptionDimension, height AVCaptionDimension) AVCaptionSize

// AVCaptionSizeMake creates a caption size with the specified width and height.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionSizeMake
func AVCaptionSizeMake(width AVCaptionDimension, height AVCaptionDimension) AVCaptionSize {
	if _aVCaptionSizeMake == nil {
		panic("AVFoundation: symbol AVCaptionSizeMake not loaded")
	}
	return _aVCaptionSizeMake(width, height)
}

var _aVCaptureReactionSystemImageNameForType func(reactionType AVCaptureReactionType) foundation.NSString

// AVCaptureReactionSystemImageNameForType returns the name of a system image that displays the recommended iconography for a specified reaction type.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureReactionType/systemImageName
func AVCaptureReactionSystemImageNameForType(reactionType AVCaptureReactionType) foundation.NSString {
	if _aVCaptureReactionSystemImageNameForType == nil {
		panic("AVFoundation: symbol AVCaptureReactionSystemImageNameForType not loaded")
	}
	return _aVCaptureReactionSystemImageNameForType(reactionType)
}

var _aVCaptureTimecodeAdvancedByFrames func(timecode AVCaptureTimecode, framesToAdd int64) AVCaptureTimecode

// AVCaptureTimecodeAdvancedByFrames generates a new timecode by adding a specified number of frames to the given timecode, handling overflow for seconds, minutes, and hours.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/advanced(_:by:)
func AVCaptureTimecodeAdvancedByFrames(timecode AVCaptureTimecode, framesToAdd int64) AVCaptureTimecode {
	if _aVCaptureTimecodeAdvancedByFrames == nil {
		panic("AVFoundation: symbol AVCaptureTimecodeAdvancedByFrames not loaded")
	}
	return _aVCaptureTimecodeAdvancedByFrames(timecode, framesToAdd)
}

var _aVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp func(timecode AVCaptureTimecode, presentationTimeStamp uintptr) objectivec.IObject

// AVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp creates a sample buffer containing Timecode Media Description metadata for integration with a video track.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/createMetadataSampleBuffer(from:associatedWithPresentationTimeStamp:)
func AVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp(timecode AVCaptureTimecode, presentationTimeStamp uintptr) objectivec.IObject {
	if _aVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp == nil {
		panic("AVFoundation: symbol AVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp not loaded")
	}
	return _aVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp(timecode, presentationTimeStamp)
}

var _aVCaptureTimecodeCreateMetadataSampleBufferForDuration func(timecode AVCaptureTimecode, duration uintptr) objectivec.IObject

// AVCaptureTimecodeCreateMetadataSampleBufferForDuration creates a sample buffer containing Timecode Media Description metadata for a specified duration.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureTimecode/createMetadataSampleBuffer(from:forDuration:)
func AVCaptureTimecodeCreateMetadataSampleBufferForDuration(timecode AVCaptureTimecode, duration uintptr) objectivec.IObject {
	if _aVCaptureTimecodeCreateMetadataSampleBufferForDuration == nil {
		panic("AVFoundation: symbol AVCaptureTimecodeCreateMetadataSampleBufferForDuration not loaded")
	}
	return _aVCaptureTimecodeCreateMetadataSampleBufferForDuration(timecode, duration)
}

var _aVMakeRectWithAspectRatioInsideRect func(aspectRatio corefoundation.CGSize, boundingRect corefoundation.CGRect) corefoundation.CGRect

// AVMakeRectWithAspectRatioInsideRect returns a scaled rectangle that maintains the specified aspect ratio within a bounding rectangle.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMakeRect(aspectRatio:insideRect:)
func AVMakeRectWithAspectRatioInsideRect(aspectRatio corefoundation.CGSize, boundingRect corefoundation.CGRect) corefoundation.CGRect {
	if _aVMakeRectWithAspectRatioInsideRect == nil {
		panic("AVFoundation: symbol AVMakeRectWithAspectRatioInsideRect not loaded")
	}
	return _aVMakeRectWithAspectRatioInsideRect(aspectRatio, boundingRect)
}

var _aVSampleBufferAttachContentKey func(sbuf uintptr, contentKey AVContentKey, outError *foundation.NSError) bool

// AVSampleBufferAttachContentKey attaches a content key to a sample buffer for the purpose of content decryption.
//
// Deprecated: Deprecated since macOS 11.0.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVSampleBufferAttachContentKey(_:_:_:)
func AVSampleBufferAttachContentKey(sbuf uintptr, contentKey AVContentKey, outError *foundation.NSError) bool {
	if _aVSampleBufferAttachContentKey == nil {
		panic("AVFoundation: symbol AVSampleBufferAttachContentKey not loaded")
	}
	return _aVSampleBufferAttachContentKey(sbuf, contentKey, outError)
}

var _cMTagCollectionCreateWithVideoOutputPreset func(allocator corefoundation.CFAllocatorRef, preset CMTagCollectionVideoOutputPreset, newCollectionOut uintptr) int32

// CMTagCollectionCreateWithVideoOutputPreset creates a collection with the required tags to describe the specified video output requirements.
//
// See: https://developer.apple.com/documentation/AVFoundation/CMTagCollectionCreateWithVideoOutputPreset
func CMTagCollectionCreateWithVideoOutputPreset(allocator corefoundation.CFAllocatorRef, preset CMTagCollectionVideoOutputPreset, newCollectionOut uintptr) int32 {
	if _cMTagCollectionCreateWithVideoOutputPreset == nil {
		panic("AVFoundation: symbol CMTagCollectionCreateWithVideoOutputPreset not loaded")
	}
	return _cMTagCollectionCreateWithVideoOutputPreset(allocator, preset, newCollectionOut)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
		registerFunc(&_aVCaptionDimensionMake, frameworkHandle, "AVCaptionDimensionMake")
		registerFunc(&_aVCaptionPointMake, frameworkHandle, "AVCaptionPointMake")
		registerFunc(&_aVCaptionSizeMake, frameworkHandle, "AVCaptionSizeMake")
		registerFunc(&_aVCaptureReactionSystemImageNameForType, frameworkHandle, "AVCaptureReactionSystemImageNameForType")
		registerFunc(&_aVCaptureTimecodeAdvancedByFrames, frameworkHandle, "AVCaptureTimecodeAdvancedByFrames")
		registerFunc(&_aVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp, frameworkHandle, "AVCaptureTimecodeCreateMetadataSampleBufferAssociatedWithPresentationTimeStamp")
		registerFunc(&_aVCaptureTimecodeCreateMetadataSampleBufferForDuration, frameworkHandle, "AVCaptureTimecodeCreateMetadataSampleBufferForDuration")
		registerFunc(&_aVMakeRectWithAspectRatioInsideRect, frameworkHandle, "AVMakeRectWithAspectRatioInsideRect")
		registerFunc(&_aVSampleBufferAttachContentKey, frameworkHandle, "AVSampleBufferAttachContentKey")
		registerFunc(&_cMTagCollectionCreateWithVideoOutputPreset, frameworkHandle, "CMTagCollectionCreateWithVideoOutputPreset")
	}

