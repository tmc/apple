// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptionRenderer] class.
var (
	_AVCaptionRendererClass     AVCaptionRendererClass
	_AVCaptionRendererClassOnce sync.Once
)

func getAVCaptionRendererClass() AVCaptionRendererClass {
	_AVCaptionRendererClassOnce.Do(func() {
		_AVCaptionRendererClass = AVCaptionRendererClass{class: objc.GetClass("AVCaptionRenderer")}
	})
	return _AVCaptionRendererClass
}

// GetAVCaptionRendererClass returns the class object for AVCaptionRenderer.
func GetAVCaptionRendererClass() AVCaptionRendererClass {
	return getAVCaptionRendererClass()
}

type AVCaptionRendererClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptionRendererClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptionRendererClass) Alloc() AVCaptionRenderer {
	rv := objc.Send[AVCaptionRenderer](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that renders captions for display at a particular time.
//
// # Overview
// 
// This object renders a caption scene for a given time from a collection of
// captions. If there aren’t any captions to display at the specified time,
// the renderer draws an empty flood fill with a zero alpha or a color.
//
// # Configuring the renderer
//
//   - [AVCaptionRenderer.Captions]: The captions to render.
//   - [AVCaptionRenderer.SetCaptions]
//   - [AVCaptionRenderer.Bounds]: The drawing bounds of caption scenes.
//   - [AVCaptionRenderer.SetBounds]
//
// # Determining scene changes
//
//   - [AVCaptionRenderer.CaptionSceneChangesInRange]: Determine render time ranges within an enclosing time range to account for visual changes among captions.
//
// # Rendering a caption
//
//   - [AVCaptionRenderer.RenderInContextForTime]: Draw the captions for the time you specify.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer
type AVCaptionRenderer struct {
	objectivec.Object
}

// AVCaptionRendererFromID constructs a [AVCaptionRenderer] from an objc.ID.
//
// An object that renders captions for display at a particular time.
func AVCaptionRendererFromID(id objc.ID) AVCaptionRenderer {
	return AVCaptionRenderer{objectivec.Object{ID: id}}
}
// NOTE: AVCaptionRenderer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptionRenderer] class.
//
// # Configuring the renderer
//
//   - [IAVCaptionRenderer.Captions]: The captions to render.
//   - [IAVCaptionRenderer.SetCaptions]
//   - [IAVCaptionRenderer.Bounds]: The drawing bounds of caption scenes.
//   - [IAVCaptionRenderer.SetBounds]
//
// # Determining scene changes
//
//   - [IAVCaptionRenderer.CaptionSceneChangesInRange]: Determine render time ranges within an enclosing time range to account for visual changes among captions.
//
// # Rendering a caption
//
//   - [IAVCaptionRenderer.RenderInContextForTime]: Draw the captions for the time you specify.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer
type IAVCaptionRenderer interface {
	objectivec.IObject

	// Topic: Configuring the renderer

	// The captions to render.
	Captions() []AVCaption
	SetCaptions(value []AVCaption)
	// The drawing bounds of caption scenes.
	Bounds() corefoundation.CGRect
	SetBounds(value corefoundation.CGRect)

	// Topic: Determining scene changes

	// Determine render time ranges within an enclosing time range to account for visual changes among captions.
	CaptionSceneChangesInRange(consideredTimeRange uintptr) []AVCaptionRendererScene

	// Topic: Rendering a caption

	// Draw the captions for the time you specify.
	RenderInContextForTime(ctx coregraphics.CGContextRef, time uintptr)
}

// Init initializes the instance.
func (c AVCaptionRenderer) Init() AVCaptionRenderer {
	rv := objc.Send[AVCaptionRenderer](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptionRenderer) Autorelease() AVCaptionRenderer {
	rv := objc.Send[AVCaptionRenderer](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptionRenderer creates a new AVCaptionRenderer instance.
func NewAVCaptionRenderer() AVCaptionRenderer {
	class := getAVCaptionRendererClass()
	rv := objc.Send[AVCaptionRenderer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Determine render time ranges within an enclosing time range to account for
// visual changes among captions.
//
// consideredTimeRange: The time range to consider for rendering.
//
// # Return Value
// 
// An array of render scenes for the time range, or an empty array if there
// are none.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/captionSceneChanges(in:)
func (c AVCaptionRenderer) CaptionSceneChangesInRange(consideredTimeRange uintptr) []AVCaptionRendererScene {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("captionSceneChangesInRange:"), consideredTimeRange)
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaptionRendererScene {
		return AVCaptionRendererSceneFromID(id)
	})
}
// Draw the captions for the time you specify.
//
// ctx: The drawing content.
//
// time: The time value for which the system draws the captions.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/render(in:for:)
func (c AVCaptionRenderer) RenderInContextForTime(ctx coregraphics.CGContextRef, time uintptr) {
	objc.Send[objc.ID](c.ID, objc.Sel("renderInContext:forTime:"), ctx, time)
}

// The captions to render.
//
// # Discussion
// 
// This property value is an empty array if there are no captions for the
// system to render.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/captions
func (c AVCaptionRenderer) Captions() []AVCaption {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("captions"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaption {
		return AVCaptionFromID(id)
	})
}
func (c AVCaptionRenderer) SetCaptions(value []AVCaption) {
	objc.Send[struct{}](c.ID, objc.Sel("setCaptions:"), objectivec.IObjectSliceToNSArray(value))
}
// The drawing bounds of caption scenes.
//
// # Discussion
// 
// Set this property value before drawing. The renderer uses the value in each
// call to [RenderInContextForTime], until you change it to a new value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptionRenderer/bounds
func (c AVCaptionRenderer) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("bounds"))
	return corefoundation.CGRect(rv)
}
func (c AVCaptionRenderer) SetBounds(value corefoundation.CGRect) {
	objc.Send[struct{}](c.ID, objc.Sel("setBounds:"), value)
}

