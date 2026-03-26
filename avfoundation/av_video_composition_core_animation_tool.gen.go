// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/quartzcore"
)

// The class instance for the [AVVideoCompositionCoreAnimationTool] class.
var (
	_AVVideoCompositionCoreAnimationToolClass     AVVideoCompositionCoreAnimationToolClass
	_AVVideoCompositionCoreAnimationToolClassOnce sync.Once
)

func getAVVideoCompositionCoreAnimationToolClass() AVVideoCompositionCoreAnimationToolClass {
	_AVVideoCompositionCoreAnimationToolClassOnce.Do(func() {
		_AVVideoCompositionCoreAnimationToolClass = AVVideoCompositionCoreAnimationToolClass{class: objc.GetClass("AVVideoCompositionCoreAnimationTool")}
	})
	return _AVVideoCompositionCoreAnimationToolClass
}

// GetAVVideoCompositionCoreAnimationToolClass returns the class object for AVVideoCompositionCoreAnimationTool.
func GetAVVideoCompositionCoreAnimationToolClass() AVVideoCompositionCoreAnimationToolClass {
	return getAVVideoCompositionCoreAnimationToolClass()
}

type AVVideoCompositionCoreAnimationToolClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVVideoCompositionCoreAnimationToolClass) Alloc() AVVideoCompositionCoreAnimationTool {
	rv := objc.Send[AVVideoCompositionCoreAnimationTool](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object used to incorporate Core Animation into a video composition.
//
// # Overview
// 
// Any animations will be interpreted on the video’s timeline, not
// real-time, so you should:
// 
// - Set animations’ [AVVideoCompositionCoreAnimationTool.BeginTime] property to
// [AVVideoCompositionCoreAnimationTool.AVCoreAnimationBeginTimeAtZero] rather than `0` (which CoreAnimation
// replaces with [CACurrentMediaTime()]); - Set [AVVideoCompositionCoreAnimationTool.IsRemovedOnCompletion] to
// [false] on animations so they are not automatically removed; - Avoid using
// layers that are associated with [UIView] objects.
//
// [AVVideoCompositionCoreAnimationTool.AVCoreAnimationBeginTimeAtZero]: https://developer.apple.com/documentation/AVFoundation/AVCoreAnimationBeginTimeAtZero
// [CACurrentMediaTime()]: https://developer.apple.com/documentation/QuartzCore/CACurrentMediaTime()
// [UIView]: https://developer.apple.com/documentation/UIKit/UIView
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionCoreAnimationTool
type AVVideoCompositionCoreAnimationTool struct {
	objectivec.Object
}

// AVVideoCompositionCoreAnimationToolFromID constructs a [AVVideoCompositionCoreAnimationTool] from an objc.ID.
//
// An object used to incorporate Core Animation into a video composition.
func AVVideoCompositionCoreAnimationToolFromID(id objc.ID) AVVideoCompositionCoreAnimationTool {
	return AVVideoCompositionCoreAnimationTool{objectivec.Object{ID: id}}
}
// NOTE: AVVideoCompositionCoreAnimationTool adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVVideoCompositionCoreAnimationTool] class.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionCoreAnimationTool
type IAVVideoCompositionCoreAnimationTool interface {
	objectivec.IObject

	// A value that sets an animation begin time to 
	AVCoreAnimationBeginTimeAtZero() float64
	// Specifies the begin time of the receiver in relation to its parent object, if applicable.
	BeginTime() float64
	SetBeginTime(value float64)
	// Determines if the animation is removed from the target layer’s animations upon completion.
	IsRemovedOnCompletion() bool
	SetIsRemovedOnCompletion(value bool)
}

// Init initializes the instance.
func (v AVVideoCompositionCoreAnimationTool) Init() AVVideoCompositionCoreAnimationTool {
	rv := objc.Send[AVVideoCompositionCoreAnimationTool](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v AVVideoCompositionCoreAnimationTool) Autorelease() AVVideoCompositionCoreAnimationTool {
	rv := objc.Send[AVVideoCompositionCoreAnimationTool](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVVideoCompositionCoreAnimationTool creates a new AVVideoCompositionCoreAnimationTool instance.
func NewAVVideoCompositionCoreAnimationTool() AVVideoCompositionCoreAnimationTool {
	class := getAVVideoCompositionCoreAnimationToolClass()
	rv := objc.Send[AVVideoCompositionCoreAnimationTool](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Adds a Core Animation layer to the video composition.
//
// layer: The Core Animation layer to add.
//
// trackID: A track ID to identify the track.
// 
// `trackID` should not match any real trackID in the source.
//
// # Return Value
// 
// A new Core Animation tool for the layer.
//
// # Discussion
// 
// You use this method to include a Core Animation layer as an individual
// track input in video composition.
// 
// Video composition instructions should reference `trackID` where the
// rendered animation should be included.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionCoreAnimationTool/init(additionalLayer:asTrackID:)
func NewVideoCompositionCoreAnimationToolWithAdditionalLayerAsTrackID(layer quartzcore.CALayer, trackID int32) AVVideoCompositionCoreAnimationTool {
	rv := objc.Send[objc.ID](objc.ID(getAVVideoCompositionCoreAnimationToolClass().class), objc.Sel("videoCompositionCoreAnimationToolWithAdditionalLayer:asTrackID:"), layer, trackID)
	return AVVideoCompositionCoreAnimationToolFromID(rv)
}

// Composes the composited video frames with the Core Animation layer.
//
// videoLayers: An array containing the video layers
//
// animationLayer: The animation layer.
//
// # Return Value
// 
// A new [AVVideoCompositionCoreAnimationTool] instance with the composited
// video frames and the rendered animation layer.
//
// # Discussion
// 
// Duplicates the composited video frames in each videoLayer and renders
// animationLayer to produce the final frame. The `videoLayers` should be in
// `animationLayer`’s sublayer tree.
// 
// The `animationLayer` should not come from, or be added to, another layer
// tree.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVVideoCompositionCoreAnimationTool/init(postProcessingAsVideoLayers:in:)
func NewVideoCompositionCoreAnimationToolWithPostProcessingAsVideoLayersInLayer(videoLayers []quartzcore.CALayer, animationLayer quartzcore.CALayer) AVVideoCompositionCoreAnimationTool {
	rv := objc.Send[objc.ID](objc.ID(getAVVideoCompositionCoreAnimationToolClass().class), objc.Sel("videoCompositionCoreAnimationToolWithPostProcessingAsVideoLayers:inLayer:"), objectivec.IObjectSliceToNSArray(videoLayers), animationLayer)
	return AVVideoCompositionCoreAnimationToolFromID(rv)
}

// A value that sets an animation begin time to
//
// See: https://developer.apple.com/documentation/avfoundation/avcoreanimationbegintimeatzero
func (v AVVideoCompositionCoreAnimationTool) AVCoreAnimationBeginTimeAtZero() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("AVCoreAnimationBeginTimeAtZero"))
	return rv
}
// Specifies the begin time of the receiver in relation to its parent object,
// if applicable.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAMediaTiming/beginTime
func (v AVVideoCompositionCoreAnimationTool) BeginTime() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("beginTime"))
	return rv
}
func (v AVVideoCompositionCoreAnimationTool) SetBeginTime(value float64) {
	objc.Send[struct{}](v.ID, objc.Sel("setBeginTime:"), value)
}
// Determines if the animation is removed from the target layer’s animations
// upon completion.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAAnimation/isRemovedOnCompletion
func (v AVVideoCompositionCoreAnimationTool) IsRemovedOnCompletion() bool {
	rv := objc.Send[bool](v.ID, objc.Sel("isRemovedOnCompletion"))
	return rv
}
func (v AVVideoCompositionCoreAnimationTool) SetIsRemovedOnCompletion(value bool) {
	objc.Send[struct{}](v.ID, objc.Sel("setIsRemovedOnCompletion:"), value)
}

