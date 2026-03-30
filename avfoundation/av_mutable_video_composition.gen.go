// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
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
// - A value for [FrameDuration] short enough to accommodate the greatest
// [nominalFrameRate] among the asset’s video tracks. If the
// [nominalFrameRate] of all of the asset’s video tracks is 0, a default
// frame rate of 30fps is used. - If the specified asset is an instance of
// [AVComposition], the [RenderSize] is set to the [NaturalSize] of the
// [AVComposition]; otherwise the [RenderSize] will be set to a value that
// encompasses all of the asset’s video tracks. - A [RenderScale] of 1.0. -
// The [AnimationTool] property set to `nil`.
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
// - A [FrameDuration] of [zero]. - A [RenderSize] of `{0.0, 0.0}`. - A `nil`
// array of [Instructions]. - The [AnimationTool] property set to `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVMutableVideoComposition/videoComposition
//
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
func (_AVMutableVideoCompositionClass AVMutableVideoCompositionClass) VideoComposition() AVMutableVideoComposition {
	rv := objc.Send[objc.ID](objc.ID(_AVMutableVideoCompositionClass.class), objc.Sel("videoComposition"))
	return AVMutableVideoCompositionFromID(rv)
}
