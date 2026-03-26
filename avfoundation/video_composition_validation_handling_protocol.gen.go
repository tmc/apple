// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Methods you can implement to indicate whether validation of a video composition should continue after specific errors are found.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionValidationHandling
type AVVideoCompositionValidationHandling interface {
	objectivec.IObject
}

// AVVideoCompositionValidationHandlingObject wraps an existing Objective-C object that conforms to the AVVideoCompositionValidationHandling protocol.
type AVVideoCompositionValidationHandlingObject struct {
	objectivec.Object
}
func (o AVVideoCompositionValidationHandlingObject) BaseObject() objectivec.Object {
	return o.Object
}

// AVVideoCompositionValidationHandlingObjectFromID constructs a [AVVideoCompositionValidationHandlingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func AVVideoCompositionValidationHandlingObjectFromID(id objc.ID) AVVideoCompositionValidationHandlingObject {
	return AVVideoCompositionValidationHandlingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Reports that a key that has an invalid value.
//
// videoComposition: The video composition being validated.
//
// key: The key being validated.
//
// # Return Value
// 
// [true] if the video composition should continue validation in order to
// report additional problems that may exist, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionValidationHandling/videoComposition(_:shouldContinueValidatingAfterFindingInvalidValueForKey:)
func (o AVVideoCompositionValidationHandlingObject) VideoCompositionShouldContinueValidatingAfterFindingInvalidValueForKey(videoComposition IAVVideoComposition, key string) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("videoComposition:shouldContinueValidatingAfterFindingInvalidValueForKey:"), videoComposition, objc.String(key))
	return rv
	}
// Reports a time range that has no corresponding video composition
// instruction.
//
// videoComposition: The video composition being validated.
//
// timeRange: The time range that has no corresponding video composition instruction.
//
// # Return Value
// 
// [true] if the video composition should continue validation in order to
// report additional problems that may exist, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionValidationHandling/videoComposition(_:shouldContinueValidatingAfterFindingEmptyTimeRange:)
func (o AVVideoCompositionValidationHandlingObject) VideoCompositionShouldContinueValidatingAfterFindingEmptyTimeRange(videoComposition IAVVideoComposition, timeRange objectivec.IObject) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("videoComposition:shouldContinueValidatingAfterFindingEmptyTimeRange:"), videoComposition, timeRange)
	return rv
	}
// Reports a video composition instruction with a time range that is invalid.
//
// videoComposition: The video composition being validated.
//
// videoCompositionInstruction: The video composition instruction.
//
// # Return Value
// 
// [true] if the video composition should continue validation in order to
// report additional problems that may exist, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// A time range is considered invalid when it overlaps with the time range of
// a prior instruction, or that contains times earlier than the time range of
// a prior instruction
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionValidationHandling/videoComposition(_:shouldContinueValidatingAfterFindingInvalidTimeRangeIn:)
func (o AVVideoCompositionValidationHandlingObject) VideoCompositionShouldContinueValidatingAfterFindingInvalidTimeRangeInInstruction(videoComposition IAVVideoComposition, videoCompositionInstruction AVVideoCompositionInstruction) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("videoComposition:shouldContinueValidatingAfterFindingInvalidTimeRangeInInstruction:"), videoComposition, videoCompositionInstruction)
	return rv
	}
// Reports a video composition layer instruction that does not correspond to
// the track ID used for the composition’s animation or to a track of the
// asset.
//
// videoComposition: The video composition being validated.
//
// videoCompositionInstruction: The video composition instruction.
//
// layerInstruction: The layer instruction.
//
// asset: The underlying asset.
//
// # Return Value
// 
// [true] if the video composition should continue validation in order to
// report additional problems that may exist, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The asset track is specified in the
// [IsValidForAssetTimeRangeValidationDelegate] method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionValidationHandling/videoComposition(_:shouldContinueValidatingAfterFindingInvalidTrackIDIn:layerInstruction:asset:)
func (o AVVideoCompositionValidationHandlingObject) VideoCompositionShouldContinueValidatingAfterFindingInvalidTrackIDInInstructionLayerInstructionAsset(videoComposition IAVVideoComposition, videoCompositionInstruction AVVideoCompositionInstruction, layerInstruction IAVVideoCompositionLayerInstruction, asset IAVAsset) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("videoComposition:shouldContinueValidatingAfterFindingInvalidTrackIDInInstruction:layerInstruction:asset:"), videoComposition, videoCompositionInstruction, layerInstruction, asset)
	return rv
	}

