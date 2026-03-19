// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CATransformLayer] class.
var (
	_CATransformLayerClass     CATransformLayerClass
	_CATransformLayerClassOnce sync.Once
)

func getCATransformLayerClass() CATransformLayerClass {
	_CATransformLayerClassOnce.Do(func() {
		_CATransformLayerClass = CATransformLayerClass{class: objc.GetClass("CATransformLayer")}
	})
	return _CATransformLayerClass
}

// GetCATransformLayerClass returns the class object for CATransformLayer.
func GetCATransformLayerClass() CATransformLayerClass {
	return getCATransformLayerClass()
}

type CATransformLayerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CATransformLayerClass) Alloc() CATransformLayer {
	rv := objc.Send[CATransformLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// Objects used to create true 3D layer hierarchies, rather than the flattened
// hierarchy rendering model used by other layer types.
//
// # Overview
// 
// Unlike normal layers, transform layers do not flatten their sublayers into
// the plane at `Z=0`. Due to this, they do not support many of the features
// of the [CALayer] class compositing model:
// 
// - Only the sublayers of a transform layer are rendered. The [CALayer]
// properties that are rendered by a layer are ignored, including:
// `backgroundColor`, `contents`, border style properties, stroke style
// properties, etc. - The properties that assume 2D image processing are also
// ignored, including: `filters`, `backgroundFilters`, `compositingFilter`,
// `mask`, `masksToBounds`, and shadow style properties. - The `opacity`
// property is applied to each sublayer individually, the transform layer does
// not form a compositing group. - The [HitTest] method should never be called
// on a transform layer as they do not have a 2D coordinate space into which
// the point can be mapped.
// 
// # Example: Displaying layers in 3D
// 
// Because [CATransformLayer] creates true 3D layer hierarchies, you can
// display otherwise hidden layers when applying 3D transforms.
// 
// The following code shows three layers with different colors but identical
// sizes added at the same position to `layer`. The blue layer is visible
// because it has the highest [CATransformLayer.ZPosition]. Defining the layer’s transform
// rotates the viewpoint in 3D space and, because `layer` is a
// [CATransformLayer], all three layers are visible as illustrated below.
// 
// [media-2826921]
// 
// However, if `layer` is created as a [CALayer], the green and red layers,
// being hidden by the blue layer, are not rendered as illustrated in the
// following figure.
// 
// [media-2826922]
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransformLayer
type CATransformLayer struct {
	CALayer
}

// CATransformLayerFromID constructs a [CATransformLayer] from an objc.ID.
//
// Objects used to create true 3D layer hierarchies, rather than the flattened
// hierarchy rendering model used by other layer types.
func CATransformLayerFromID(id objc.ID) CATransformLayer {
	return CATransformLayer{CALayer: CALayerFromID(id)}
}
// NOTE: CATransformLayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CATransformLayer] class.
//
// See: https://developer.apple.com/documentation/QuartzCore/CATransformLayer
type ICATransformLayer interface {
	ICALayer
}

// Init initializes the instance.
func (t CATransformLayer) Init() CATransformLayer {
	rv := objc.Send[CATransformLayer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t CATransformLayer) Autorelease() CATransformLayer {
	rv := objc.Send[CATransformLayer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewCATransformLayer creates a new CATransformLayer instance.
func NewCATransformLayer() CATransformLayer {
	class := getCATransformLayerClass()
	rv := objc.Send[CATransformLayer](objc.ID(class.class), objc.Sel("new"))
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
func NewTransformLayerWithLayer(layer objectivec.IObject) CATransformLayer {
	instance := getCATransformLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLayer:"), layer)
	return CATransformLayerFromID(rv)
}

