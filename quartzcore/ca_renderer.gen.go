// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/corevideo"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CARenderer] class.
var (
	_CARendererClass     CARendererClass
	_CARendererClassOnce sync.Once
)

func getCARendererClass() CARendererClass {
	_CARendererClassOnce.Do(func() {
		_CARendererClass = CARendererClass{class: objc.GetClass("CARenderer")}
	})
	return _CARendererClass
}

// GetCARendererClass returns the class object for CARenderer.
func GetCARendererClass() CARendererClass {
	return getCARendererClass()
}

type CARendererClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CARendererClass) Alloc() CARenderer {
	rv := objc.Send[CARenderer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A layer that allows an application to render a layer tree into a Core
// OpenGL context.
//
// # Overview
// 
// For real-time output you should use an instance of [NSView] to host the
// layer-tree.
//
// [NSView]: https://developer.apple.com/documentation/AppKit/NSView
//
// # Getting the Rendered Layer
//
//   - [CARenderer.Layer]: The root layer of the layer-tree the receiver should render.
//   - [CARenderer.SetLayer]
//
// # Determining Layer Bounds
//
//   - [CARenderer.Bounds]: The bounds of the receiver.
//   - [CARenderer.SetBounds]
//
// # Rendering a Frame
//
//   - [CARenderer.BeginFrameAtTimeTimeStamp]: Begin rendering a frame at the specified time.
//   - [CARenderer.UpdateBounds]: Returns the bounds of the update region that contains all pixels that will be rendered by the current frame.
//   - [CARenderer.AddUpdateRect]: Adds the rectangle to the update region of the current frame.
//   - [CARenderer.Render]: Render the update region of the current frame to the target context.
//   - [CARenderer.NextFrameTime]: Returns the time at which the next update should happen.
//   - [CARenderer.EndFrame]: Release any data associated with the current frame.
//
// # Instance Methods
//
//   - [CARenderer.SetDestination]
//
// See: https://developer.apple.com/documentation/QuartzCore/CARenderer
type CARenderer struct {
	objectivec.Object
}

// CARendererFromID constructs a [CARenderer] from an objc.ID.
//
// A layer that allows an application to render a layer tree into a Core
// OpenGL context.
func CARendererFromID(id objc.ID) CARenderer {
	return CARenderer{objectivec.Object{ID: id}}
}
// NOTE: CARenderer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CARenderer] class.
//
// # Getting the Rendered Layer
//
//   - [ICARenderer.Layer]: The root layer of the layer-tree the receiver should render.
//   - [ICARenderer.SetLayer]
//
// # Determining Layer Bounds
//
//   - [ICARenderer.Bounds]: The bounds of the receiver.
//   - [ICARenderer.SetBounds]
//
// # Rendering a Frame
//
//   - [ICARenderer.BeginFrameAtTimeTimeStamp]: Begin rendering a frame at the specified time.
//   - [ICARenderer.UpdateBounds]: Returns the bounds of the update region that contains all pixels that will be rendered by the current frame.
//   - [ICARenderer.AddUpdateRect]: Adds the rectangle to the update region of the current frame.
//   - [ICARenderer.Render]: Render the update region of the current frame to the target context.
//   - [ICARenderer.NextFrameTime]: Returns the time at which the next update should happen.
//   - [ICARenderer.EndFrame]: Release any data associated with the current frame.
//
// # Instance Methods
//
//   - [ICARenderer.SetDestination]
//
// See: https://developer.apple.com/documentation/QuartzCore/CARenderer
type ICARenderer interface {
	objectivec.IObject

	// Topic: Getting the Rendered Layer

	// The root layer of the layer-tree the receiver should render.
	Layer() ICALayer
	SetLayer(value ICALayer)

	// Topic: Determining Layer Bounds

	// The bounds of the receiver.
	Bounds() corefoundation.CGRect
	SetBounds(value corefoundation.CGRect)

	// Topic: Rendering a Frame

	// Begin rendering a frame at the specified time.
	BeginFrameAtTimeTimeStamp(t float64, ts *corevideo.CVTimeStamp)
	// Returns the bounds of the update region that contains all pixels that will be rendered by the current frame.
	UpdateBounds() corefoundation.CGRect
	// Adds the rectangle to the update region of the current frame.
	AddUpdateRect(r corefoundation.CGRect)
	// Render the update region of the current frame to the target context.
	Render()
	// Returns the time at which the next update should happen.
	NextFrameTime() float64
	// Release any data associated with the current frame.
	EndFrame()

	// Topic: Instance Methods

	SetDestination(tex objectivec.IObject)
}

// Init initializes the instance.
func (r CARenderer) Init() CARenderer {
	rv := objc.Send[CARenderer](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r CARenderer) Autorelease() CARenderer {
	rv := objc.Send[CARenderer](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewCARenderer creates a new CARenderer instance.
func NewCARenderer() CARenderer {
	class := getCARendererClass()
	rv := objc.Send[CARenderer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a layer renderer from a Metal texture.
//
// See: https://developer.apple.com/documentation/QuartzCore/CARenderer/init(mtlTexture:options:)
func NewRendererWithMTLTextureOptions(tex objectivec.IObject, dict foundation.INSDictionary) CARenderer {
	rv := objc.Send[objc.ID](objc.ID(getCARendererClass().class), objc.Sel("rendererWithMTLTexture:options:"), tex, dict)
	return CARendererFromID(rv)
}

// Begin rendering a frame at the specified time.
//
// t: The layer time.
//
// ts: The display timestamp associated with timeInterval. Can be null.
//
// See: https://developer.apple.com/documentation/QuartzCore/CARenderer/beginFrame(atTime:timeStamp:)
func (r CARenderer) BeginFrameAtTimeTimeStamp(t float64, ts *corevideo.CVTimeStamp) {
	objc.Send[objc.ID](r.ID, objc.Sel("beginFrameAtTime:timeStamp:"), t, ts)
}
// Returns the bounds of the update region that contains all pixels that will
// be rendered by the current frame.
//
// # Return Value
// 
// The bounds of the update region..
//
// # Discussion
// 
// Initially `updateBounds` will include all differences between the current
// frame and the previously rendered frame.
//
// See: https://developer.apple.com/documentation/QuartzCore/CARenderer/updateBounds()
func (r CARenderer) UpdateBounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](r.ID, objc.Sel("updateBounds"))
	return corefoundation.CGRect(rv)
}
// Adds the rectangle to the update region of the current frame.
//
// r: A rectangle defining the region to be added to the update region.
//
// See: https://developer.apple.com/documentation/QuartzCore/CARenderer/addUpdate(_:)
func (r_ CARenderer) AddUpdateRect(r corefoundation.CGRect) {
	objc.Send[objc.ID](r_.ID, objc.Sel("addUpdateRect:"), r)
}
// Render the update region of the current frame to the target context.
//
// See: https://developer.apple.com/documentation/QuartzCore/CARenderer/render()
func (r CARenderer) Render() {
	objc.Send[objc.ID](r.ID, objc.Sel("render"))
}
// Returns the time at which the next update should happen.
//
// # Return Value
// 
// The time at which the next update should happen.
//
// # Discussion
// 
// If infinite, no update needs to be scheduled yet. If `nextFrameTime` is the
// current frame time, a continuous animation is running and an update should
// be scheduled after an appropriate delay.
//
// See: https://developer.apple.com/documentation/QuartzCore/CARenderer/nextFrameTime()
func (r CARenderer) NextFrameTime() float64 {
	rv := objc.Send[float64](r.ID, objc.Sel("nextFrameTime"))
	return rv
}
// Release any data associated with the current frame.
//
// See: https://developer.apple.com/documentation/QuartzCore/CARenderer/endFrame()
func (r CARenderer) EndFrame() {
	objc.Send[objc.ID](r.ID, objc.Sel("endFrame"))
}
//
// See: https://developer.apple.com/documentation/QuartzCore/CARenderer/setDestination(_:)
func (r CARenderer) SetDestination(tex objectivec.IObject) {
	objc.Send[objc.ID](r.ID, objc.Sel("setDestination:"), tex)
}

// The root layer of the layer-tree the receiver should render.
//
// See: https://developer.apple.com/documentation/QuartzCore/CARenderer/layer
func (r CARenderer) Layer() ICALayer {
	rv := objc.Send[objc.ID](r.ID, objc.Sel("layer"))
	return CALayerFromID(objc.ID(rv))
}
func (r CARenderer) SetLayer(value ICALayer) {
	objc.Send[struct{}](r.ID, objc.Sel("setLayer:"), value)
}
// The bounds of the receiver.
//
// See: https://developer.apple.com/documentation/QuartzCore/CARenderer/bounds
func (r CARenderer) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](r.ID, objc.Sel("bounds"))
	return corefoundation.CGRect(rv)
}
func (r CARenderer) SetBounds(value corefoundation.CGRect) {
	objc.Send[struct{}](r.ID, objc.Sel("setBounds:"), value)
}

