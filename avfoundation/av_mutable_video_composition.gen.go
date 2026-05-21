// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [AVMutableVideoComposition] class.
var (
	_AVMutableVideoCompositionClass     AVMutableVideoCompositionClass
	_AVMutableVideoCompositionClassOnce sync.Once
)

func getAVMutableVideoCompositionClass() AVMutableVideoCompositionClass {
	_AVMutableVideoCompositionClassOnce.Do(func() {
		_AVMutableVideoCompositionClass = AVMutableVideoCompositionClass{class: objc.GetClass("AVMutableVideoComposition")}
	})
	return _AVMutableVideoCompositionClass
}

// GetAVMutableVideoCompositionClass returns the class object for AVMutableVideoComposition.
func GetAVMutableVideoCompositionClass() AVMutableVideoCompositionClass {
	return getAVMutableVideoCompositionClass()
}

type AVMutableVideoCompositionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVMutableVideoCompositionClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVMutableVideoCompositionClass) Alloc() AVMutableVideoComposition {
	rv := objc.Send[AVMutableVideoComposition](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// A mutable video composition subclass.
//
// # Overview
//
// If you use the built-in video compositor, the instructions a video
// composition contain can specify a spatial transformation, an opacity value,
// and a cropping rectangle for each video source. This values can vary over
// time by applying linear ramping functions.
//
// You can create a custom video compositor by implementing the
// [AVVideoCompositing] protocol. The system provides the custom video
// compositor with pixel buffers for each of its video sources during
// playback, and can perform arbitrary graphical operations on them to produce
// visual output.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition
type AVMutableVideoComposition struct {
	AVVideoComposition
}

// AVMutableVideoCompositionFromID constructs a [AVMutableVideoComposition] from an objc.ID.
//
// A mutable video composition subclass.
func AVMutableVideoCompositionFromID(id objc.ID) AVMutableVideoComposition {
	return AVMutableVideoComposition{AVVideoComposition: AVVideoCompositionFromID(id)}
}

// NOTE: AVMutableVideoComposition adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVMutableVideoComposition] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition
type IAVMutableVideoComposition interface {
	IAVVideoComposition
}

// Init initializes the instance.
func (m AVMutableVideoComposition) Init() AVMutableVideoComposition {
	rv := objc.Send[AVMutableVideoComposition](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m AVMutableVideoComposition) Autorelease() AVMutableVideoComposition {
	rv := objc.Send[AVMutableVideoComposition](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVMutableVideoComposition creates a new AVMutableVideoComposition instance.
func NewAVMutableVideoComposition() AVMutableVideoComposition {
	class := getAVMutableVideoCompositionClass()
	rv := objc.Send[AVMutableVideoComposition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a mutable video composition with the specified asset properties.
//
// asset: An instance of [AVAsset]. Ensure that the duration and tracks properties of
// the asset are already loaded before invoking this method.
//
// # Discussion
//
// The returned [AVMutableVideoComposition] has instructions that respect the
// spatial properties and time ranges of the specified asset’s video tracks.
//
// It also has the following values for its properties:
//
// - A value for [AVMutableVideoComposition.FrameDuration] short enough to
// accommodate the greatest [nominalFrameRate] among the asset’s video
// tracks. If the [nominalFrameRate] of all of the asset’s video tracks is
// 0, a default frame rate of 30fps is used. - If the specified asset is an
// instance of [AVComposition], the [AVMutableVideoComposition.RenderSize] is
// set to the [AVComposition.NaturalSize] of the [AVComposition]; otherwise
// the [AVMutableVideoComposition.RenderSize] will be set to a value that
// encompasses all of the asset’s video tracks. - A
// [AVMutableVideoComposition.RenderScale] of 1.0. - The
// [AVMutableVideoComposition.AnimationTool] property set to `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/init(propertiesOf:)
//
// [nominalFrameRate]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/nominalFrameRate
func NewMutableVideoCompositionWithPropertiesOfAsset(asset IAVAsset) AVMutableVideoComposition {
	rv := objc.Send[objc.ID](objc.ID(getAVMutableVideoCompositionClass().class), objc.Sel("videoCompositionWithPropertiesOfAsset:"), asset)
	return AVMutableVideoCompositionFromID(rv)
}

// Creates a mutable video composition with the specified asset properties and
// a prototype video composition instruction.
//
// asset: The asset for which to create a video composition. Load the asset’s
// [duration] and [tracks] properties before invoking this method.
//
// prototypeInstruction: A video composition instruction to use as a prototype.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/init(propertiesOf:prototypeInstruction:)
//
// [duration]: https://developer.apple.com/documentation/AVFoundation/AVAsset/duration
// [tracks]: https://developer.apple.com/documentation/AVFoundation/AVAsset/tracks
func NewMutableVideoCompositionWithPropertiesOfAssetPrototypeInstruction(asset IAVAsset, prototypeInstruction AVVideoCompositionInstruction) AVMutableVideoComposition {
	rv := objc.Send[objc.ID](objc.ID(getAVMutableVideoCompositionClass().class), objc.Sel("videoCompositionWithPropertiesOfAsset:prototypeInstruction:"), asset, prototypeInstruction)
	return AVMutableVideoCompositionFromID(rv)
}

// Returns a new video composition that’s configured to apply Core Image
// filters to each video frame of the specified asset.
//
// asset: The asset whose configuration matches the intended use of the video
// composition.
//
// applier: A block that AVFoundation calls when processing each video frame.
//
// The block takes a single parameter and has no return value:
//
// request: An [AVAsynchronousCIImageFilteringRequest] object representing the
// frame to be processed.
//
// completionHandler: A block the system calls when it finishes creating the new video
// composition.
//
// # Discussion
//
// The composition calls the specified handler one time for each frame to
// display (or processed for export) from the asset’s first enabled video
// track. In that block, you access the video frame and return a filtered
// result using the provided [AVAsynchronousCIImageFilteringRequest] object.
// Use that object’s [sourceImage] property to get the video frame in the
// form of a [CIImage] object you can apply filters to. Pass the result of
// your filters to the `request` object’s [finish(with:context:)] method.
// (If your filter rendering fails, call the `request` object’s
// [finish(with:)] method if you can’t apply filters).
//
// Creating a composition with this method sets values for the following
// properties:
//
// - The value of the [AVVideoComposition.FrameDuration] property accommodates
// the [nominalFrameRate] value for the asset’s first enabled video track.
// If the nominal frame rate is zero, AVFoundation uses a default frame rate
// of 30 fps. - The [AVVideoComposition.RenderSize] property value a size that
// encompasses the asset’s first enabled video track, respecting the
// track’s [preferredTransform] property. - The
// [AVVideoComposition.RenderScale] property value is `1.0`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/videoComposition(with:applyingCIFiltersWithHandler:completionHandler:)
//
// [AVAsynchronousCIImageFilteringRequest]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousCIImageFilteringRequest
// [CIImage]: https://developer.apple.com/documentation/CoreImage/CIImage
// [finish(with:)]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousCIImageFilteringRequest/finish(with:)
// [finish(with:context:)]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousCIImageFilteringRequest/finish(with:context:)
// [nominalFrameRate]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/nominalFrameRate
// [preferredTransform]: https://developer.apple.com/documentation/AVFoundation/AVAssetTrack/preferredTransform
// [sourceImage]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousCIImageFilteringRequest/sourceImage
//
// [AVAsynchronousCIImageFilteringRequest]: https://developer.apple.com/documentation/AVFoundation/AVAsynchronousCIImageFilteringRequest
func (_AVMutableVideoCompositionClass AVMutableVideoCompositionClass) VideoCompositionWithAssetApplyingCIFiltersWithHandlerCompletionHandler(asset IAVAsset, applier AVAsynchronousCIImageFilteringRequestHandler, completionHandler AVMutableVideoCompositionErrorHandler) {
	_block1, _ := NewAVAsynchronousCIImageFilteringRequestBlock(applier)
	_block2, _ := NewAVMutableVideoCompositionErrorBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_AVMutableVideoCompositionClass.class), objc.Sel("videoCompositionWithAsset:applyingCIFiltersWithHandler:completionHandler:"), asset, _block1, _block2)
}

// Creates a new mutable video composition.
//
// # Return Value
//
// A newly created and initialized instance of [AVMutableVideoComposition].
//
// # Discussion
//
// The returned [AVMutableVideoComposition] has the following properties:
//
// - A [AVMutableVideoComposition.FrameDuration] of [zero]. - A
// [AVMutableVideoComposition.RenderSize] of `{0.0, 0.0}`. - A `nil` array of
// [AVMutableVideoComposition.Instructions]. - The
// [AVMutableVideoComposition.AnimationTool] property set to `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/videoComposition
//
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
func (_AVMutableVideoCompositionClass AVMutableVideoCompositionClass) VideoComposition() AVMutableVideoComposition {
	rv := objc.Send[objc.ID](objc.ID(_AVMutableVideoCompositionClass.class), objc.Sel("videoComposition"))
	return AVMutableVideoCompositionFromID(rv)
}

// VideoCompositionWithAssetApplyingCIFiltersWithHandlerSync is a synchronous wrapper around [AVMutableVideoComposition.VideoCompositionWithAssetApplyingCIFiltersWithHandlerCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (mc AVMutableVideoCompositionClass) VideoCompositionWithAssetApplyingCIFiltersWithHandlerSync(ctx context.Context, asset IAVAsset, applier AVAsynchronousCIImageFilteringRequestHandler) (*AVMutableVideoComposition, error) {
	type result struct {
		val *AVMutableVideoComposition
		err error
	}
	done := make(chan result, 1)
	mc.VideoCompositionWithAssetApplyingCIFiltersWithHandlerCompletionHandler(asset, applier, func(val *AVMutableVideoComposition, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
