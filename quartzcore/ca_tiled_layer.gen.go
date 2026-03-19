// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CATiledLayer] class.
var (
	_CATiledLayerClass     CATiledLayerClass
	_CATiledLayerClassOnce sync.Once
)

func getCATiledLayerClass() CATiledLayerClass {
	_CATiledLayerClassOnce.Do(func() {
		_CATiledLayerClass = CATiledLayerClass{class: objc.GetClass("CATiledLayer")}
	})
	return _CATiledLayerClass
}

// GetCATiledLayerClass returns the class object for CATiledLayer.
func GetCATiledLayerClass() CATiledLayerClass {
	return getCATiledLayerClass()
}

type CATiledLayerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CATiledLayerClass) Alloc() CATiledLayer {
	rv := objc.Send[CATiledLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A layer that provides a way to asynchronously provide tiles of the
// layer’s content, potentially cached at multiple levels of detail.
//
// # Overview
// 
// As more data is required by the renderer, the layer’s [DrawInContext]
// method is called on one or more background threads to supply the drawing
// operations to fill in one tile of data. The clip bounds and current
// transformation matrix (CTM) of the drawing context can be used to determine
// the bounds and resolution of the tile being requested.
// 
// Regions of the layer may be invalidated using the [SetNeedsDisplayInRect]
// method however the update will be asynchronous. While the next display
// update will most likely not contain the updated content, a future update
// will.
//
// # Levels of detail
//
//   - [CATiledLayer.LevelsOfDetail]: The number of levels of detail maintained by this layer.
//   - [CATiledLayer.SetLevelsOfDetail]
//   - [CATiledLayer.LevelsOfDetailBias]: The number of magnified levels of detail for this layer.
//   - [CATiledLayer.SetLevelsOfDetailBias]
//
// # Layer tile size
//
//   - [CATiledLayer.TileSize]: The maximum size of each tile used to create the layer’s content.
//   - [CATiledLayer.SetTileSize]
//
// See: https://developer.apple.com/documentation/QuartzCore/CATiledLayer
type CATiledLayer struct {
	CALayer
}

// CATiledLayerFromID constructs a [CATiledLayer] from an objc.ID.
//
// A layer that provides a way to asynchronously provide tiles of the
// layer’s content, potentially cached at multiple levels of detail.
func CATiledLayerFromID(id objc.ID) CATiledLayer {
	return CATiledLayer{CALayer: CALayerFromID(id)}
}
// NOTE: CATiledLayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CATiledLayer] class.
//
// # Levels of detail
//
//   - [ICATiledLayer.LevelsOfDetail]: The number of levels of detail maintained by this layer.
//   - [ICATiledLayer.SetLevelsOfDetail]
//   - [ICATiledLayer.LevelsOfDetailBias]: The number of magnified levels of detail for this layer.
//   - [ICATiledLayer.SetLevelsOfDetailBias]
//
// # Layer tile size
//
//   - [ICATiledLayer.TileSize]: The maximum size of each tile used to create the layer’s content.
//   - [ICATiledLayer.SetTileSize]
//
// See: https://developer.apple.com/documentation/QuartzCore/CATiledLayer
type ICATiledLayer interface {
	ICALayer

	// Topic: Levels of detail

	// The number of levels of detail maintained by this layer.
	LevelsOfDetail() uintptr
	SetLevelsOfDetail(value uintptr)
	// The number of magnified levels of detail for this layer.
	LevelsOfDetailBias() uintptr
	SetLevelsOfDetailBias(value uintptr)

	// Topic: Layer tile size

	// The maximum size of each tile used to create the layer’s content.
	TileSize() corefoundation.CGSize
	SetTileSize(value corefoundation.CGSize)
}

// Init initializes the instance.
func (t CATiledLayer) Init() CATiledLayer {
	rv := objc.Send[CATiledLayer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t CATiledLayer) Autorelease() CATiledLayer {
	rv := objc.Send[CATiledLayer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewCATiledLayer creates a new CATiledLayer instance.
func NewCATiledLayer() CATiledLayer {
	class := getCATiledLayerClass()
	rv := objc.Send[CATiledLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Override to copy or initialize custom fields of the specified layer.
//
// layer: The layer from which custom fields should be copied.
//
// # Return Value
// 
// A layer instance with any custom instance variables copied from `layer`.
//
// # Discussion
// 
// This initializer is used to create shadow copies of layers, for example,
// for the [PresentationLayer] method. Using this method in any other
// situation will produce undefined behavior. For example, do not use this
// method to initialize a new layer with an existing layer’s content.
// 
// If you are implementing a custom layer subclass, you can override this
// method and use it to copy the values of instance variables into the new
// object. Subclasses should always invoke the superclass implementation.
// 
// This method is the designated initializer for layer objects in the
// presentation layer.
//
// See: https://developer.apple.com/documentation/QuartzCore/CALayer/init(layer:)
func NewTiledLayerWithLayer(layer objectivec.IObject) CATiledLayer {
	instance := getCATiledLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLayer:"), layer)
	return CATiledLayerFromID(rv)
}

// The time, in seconds, that newly added images take to “fade-in” to the
// rendered representation of the tiled layer.
//
// # Discussion
// 
// The default implementation returns 0.25 seconds.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATiledLayer/fadeDuration()
func (_CATiledLayerClass CATiledLayerClass) FadeDuration() float64 {
	rv := objc.Send[float64](objc.ID(_CATiledLayerClass.class), objc.Sel("fadeDuration"))
	return rv
}

// The number of levels of detail maintained by this layer.
//
// # Discussion
// 
// Defaults to 1. Each level of detail is half the resolution of the previous
// level. If too many levels are specified for the current size of the layer,
// then the number of levels is clamped to the maximum value (the bottom most
// level of detail must contain at least a single pixel in each dimension.)
//
// See: https://developer.apple.com/documentation/QuartzCore/CATiledLayer/levelsOfDetail
func (t CATiledLayer) LevelsOfDetail() uintptr {
	rv := objc.Send[uintptr](t.ID, objc.Sel("levelsOfDetail"))
	return rv
}
func (t CATiledLayer) SetLevelsOfDetail(value uintptr) {
	objc.Send[struct{}](t.ID, objc.Sel("setLevelsOfDetail:"), value)
}
// The number of magnified levels of detail for this layer.
//
// # Discussion
// 
// Defaults to 0. Each previous level of detail is twice the resolution of the
// later. For example, specifying a value of 2 means that the layer has two
// extra levels of detail: 2x and 4x.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATiledLayer/levelsOfDetailBias
func (t CATiledLayer) LevelsOfDetailBias() uintptr {
	rv := objc.Send[uintptr](t.ID, objc.Sel("levelsOfDetailBias"))
	return rv
}
func (t CATiledLayer) SetLevelsOfDetailBias(value uintptr) {
	objc.Send[struct{}](t.ID, objc.Sel("setLevelsOfDetailBias:"), value)
}
// The maximum size of each tile used to create the layer’s content.
//
// # Discussion
// 
// Defaults to (256.0, 256.0).
//
// See: https://developer.apple.com/documentation/QuartzCore/CATiledLayer/tileSize
func (t CATiledLayer) TileSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](t.ID, objc.Sel("tileSize"))
	return corefoundation.CGSize(rv)
}
func (t CATiledLayer) SetTileSize(value corefoundation.CGSize) {
	objc.Send[struct{}](t.ID, objc.Sel("setTileSize:"), value)
}

