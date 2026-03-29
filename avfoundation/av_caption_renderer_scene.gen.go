// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptionRendererScene] class.
var (
	_AVCaptionRendererSceneClass     AVCaptionRendererSceneClass
	_AVCaptionRendererSceneClassOnce sync.Once
)

func getAVCaptionRendererSceneClass() AVCaptionRendererSceneClass {
	_AVCaptionRendererSceneClassOnce.Do(func() {
		_AVCaptionRendererSceneClass = AVCaptionRendererSceneClass{class: objc.GetClass("AVCaptionRendererScene")}
	})
	return _AVCaptionRendererSceneClass
}

// GetAVCaptionRendererSceneClass returns the class object for AVCaptionRendererScene.
func GetAVCaptionRendererSceneClass() AVCaptionRendererSceneClass {
	return getAVCaptionRendererSceneClass()
}

type AVCaptionRendererSceneClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptionRendererSceneClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptionRendererSceneClass) Alloc() AVCaptionRendererScene {
	rv := objc.Send[AVCaptionRendererScene](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that holds a time range and an associated state which indicates
// when the renderer draws output.
//
// # Overview
// 
// To render a scene, the object considers state like the existence of
// captions and regions, their temporal overlaps, and whether captions use
// animation effects. Your app can request time ranges where visual
// differences exist and use these time ranges to optimize drawing
// performance, like drawing once per scene. Alternatively, it can ignore
// scenes, and instead call [RenderInContextForTime] repeatedly, but this may
// have additional performance impact.
//
// # Inspecting the scene
//
//   - [AVCaptionRendererScene.TimeRange]: The time range during which the system doesn’t modify the scene.
//   - [AVCaptionRendererScene.HasActiveCaptions]: A Boolean value that indicates whether the scene contains one or more active captions.
//   - [AVCaptionRendererScene.NeedsPeriodicRefresh]: A Boolean value that indicates whether the scene requires redrawing while your app progresses through the content.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/Scene
type AVCaptionRendererScene struct {
	objectivec.Object
}

// AVCaptionRendererSceneFromID constructs a [AVCaptionRendererScene] from an objc.ID.
//
// An object that holds a time range and an associated state which indicates
// when the renderer draws output.
func AVCaptionRendererSceneFromID(id objc.ID) AVCaptionRendererScene {
	return AVCaptionRendererScene{objectivec.Object{ID: id}}
}
// NOTE: AVCaptionRendererScene adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptionRendererScene] class.
//
// # Inspecting the scene
//
//   - [IAVCaptionRendererScene.TimeRange]: The time range during which the system doesn’t modify the scene.
//   - [IAVCaptionRendererScene.HasActiveCaptions]: A Boolean value that indicates whether the scene contains one or more active captions.
//   - [IAVCaptionRendererScene.NeedsPeriodicRefresh]: A Boolean value that indicates whether the scene requires redrawing while your app progresses through the content.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/Scene
type IAVCaptionRendererScene interface {
	objectivec.IObject

	// Topic: Inspecting the scene

	// The time range during which the system doesn’t modify the scene.
	TimeRange() coremedia.CMTimeRange
	// A Boolean value that indicates whether the scene contains one or more active captions.
	HasActiveCaptions() bool
	// A Boolean value that indicates whether the scene requires redrawing while your app progresses through the content.
	NeedsPeriodicRefresh() bool
}

// Init initializes the instance.
func (c AVCaptionRendererScene) Init() AVCaptionRendererScene {
	rv := objc.Send[AVCaptionRendererScene](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptionRendererScene) Autorelease() AVCaptionRendererScene {
	rv := objc.Send[AVCaptionRendererScene](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptionRendererScene creates a new AVCaptionRendererScene instance.
func NewAVCaptionRendererScene() AVCaptionRendererScene {
	class := getAVCaptionRendererSceneClass()
	rv := objc.Send[AVCaptionRendererScene](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The time range during which the system doesn’t modify the scene.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/Scene/timeRange
func (c AVCaptionRendererScene) TimeRange() coremedia.CMTimeRange {
	rv := objc.Send[coremedia.CMTimeRange](c.ID, objc.Sel("timeRange"))
	return coremedia.CMTimeRange(rv)
}
// A Boolean value that indicates whether the scene contains one or more
// active captions.
//
// # Discussion
// 
// Knowing when the renderer has active captions can be useful to scrub to
// times where captions are present, or skip scenes where no captions exist.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/Scene/hasActiveCaptions
func (c AVCaptionRendererScene) HasActiveCaptions() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasActiveCaptions"))
	return rv
}
// A Boolean value that indicates whether the scene requires redrawing while
// your app progresses through the content.
//
// # Discussion
// 
// If your app isn’t progressing through the content, a single render at the
// current time is enough.
// 
// Choose a refresh rate appropriate for your app. For example, an app may
// choose rates that match rates of associated video frames or other timing
// appropriate for the client.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/Scene/needsPeriodicRefresh
func (c AVCaptionRendererScene) NeedsPeriodicRefresh() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("needsPeriodicRefresh"))
	return rv
}

