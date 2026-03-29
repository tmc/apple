// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CAGradientLayer] class.
var (
	_CAGradientLayerClass     CAGradientLayerClass
	_CAGradientLayerClassOnce sync.Once
)

func getCAGradientLayerClass() CAGradientLayerClass {
	_CAGradientLayerClassOnce.Do(func() {
		_CAGradientLayerClass = CAGradientLayerClass{class: objc.GetClass("CAGradientLayer")}
	})
	return _CAGradientLayerClass
}

// GetCAGradientLayerClass returns the class object for CAGradientLayer.
func GetCAGradientLayerClass() CAGradientLayerClass {
	return getCAGradientLayerClass()
}

type CAGradientLayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CAGradientLayerClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CAGradientLayerClass) Alloc() CAGradientLayer {
	rv := objc.Send[CAGradientLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A layer that draws a color gradient over its background color, filling the
// shape of the layer.
//
// # Overview
// 
// You use a gradient layer to create a color gradient containing an arbitrary
// number of colors. By default, the colors are spread uniformly across the
// layer, but you can optionally specify locations for control over the color
// positions through the gradient.
// 
// The following code shows how to create a gradient layer containing four
// colors that are evenly distributed through the gradient. Rotating the layer
// by 90° ([CAGradientLayer.Pi] ⁄ `2` radians) gives a horizontal gradient.
// 
// The following figure shows the appearance of the gradient layer.
// 
// [media-2825193]
//
// # Gradient Style Properties
//
//   - [CAGradientLayer.Colors]: An array of [CGColorRef] objects defining the color of each gradient stop. Animatable.
//   - [CAGradientLayer.SetColors]
//   - [CAGradientLayer.Locations]: An optional array of NSNumber objects defining the location of each gradient stop. Animatable.
//   - [CAGradientLayer.SetLocations]
//   - [CAGradientLayer.EndPoint]: The end point of the gradient when drawn in the layer’s coordinate space. Animatable.
//   - [CAGradientLayer.SetEndPoint]
//   - [CAGradientLayer.StartPoint]: The start point of the gradient when drawn in the layer’s coordinate space. Animatable.
//   - [CAGradientLayer.SetStartPoint]
//   - [CAGradientLayer.Type]: Style of gradient drawn by the layer.
//   - [CAGradientLayer.SetType]
//
// See: https://developer.apple.com/documentation/QuartzCore/CAGradientLayer
type CAGradientLayer struct {
	CALayer
}

// CAGradientLayerFromID constructs a [CAGradientLayer] from an objc.ID.
//
// A layer that draws a color gradient over its background color, filling the
// shape of the layer.
func CAGradientLayerFromID(id objc.ID) CAGradientLayer {
	return CAGradientLayer{CALayer: CALayerFromID(id)}
}
// NOTE: CAGradientLayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CAGradientLayer] class.
//
// # Gradient Style Properties
//
//   - [ICAGradientLayer.Colors]: An array of [CGColorRef] objects defining the color of each gradient stop. Animatable.
//   - [ICAGradientLayer.SetColors]
//   - [ICAGradientLayer.Locations]: An optional array of NSNumber objects defining the location of each gradient stop. Animatable.
//   - [ICAGradientLayer.SetLocations]
//   - [ICAGradientLayer.EndPoint]: The end point of the gradient when drawn in the layer’s coordinate space. Animatable.
//   - [ICAGradientLayer.SetEndPoint]
//   - [ICAGradientLayer.StartPoint]: The start point of the gradient when drawn in the layer’s coordinate space. Animatable.
//   - [ICAGradientLayer.SetStartPoint]
//   - [ICAGradientLayer.Type]: Style of gradient drawn by the layer.
//   - [ICAGradientLayer.SetType]
//
// See: https://developer.apple.com/documentation/QuartzCore/CAGradientLayer
type ICAGradientLayer interface {
	ICALayer

	// Topic: Gradient Style Properties

	// An array of [CGColorRef] objects defining the color of each gradient stop. Animatable.
	Colors() foundation.INSArray
	SetColors(value foundation.INSArray)
	// An optional array of NSNumber objects defining the location of each gradient stop. Animatable.
	Locations() []foundation.NSNumber
	SetLocations(value []foundation.NSNumber)
	// The end point of the gradient when drawn in the layer’s coordinate space. Animatable.
	EndPoint() corefoundation.CGPoint
	SetEndPoint(value corefoundation.CGPoint)
	// The start point of the gradient when drawn in the layer’s coordinate space. Animatable.
	StartPoint() corefoundation.CGPoint
	SetStartPoint(value corefoundation.CGPoint)
	// Style of gradient drawn by the layer.
	Type() CAGradientLayerType
	SetType(value CAGradientLayerType)

	// The mathematical constant pi (π), approximately equal to 3.14159.
	Pi() objectivec.IObject
	SetPi(value objectivec.IObject)
}

// Init initializes the instance.
func (g CAGradientLayer) Init() CAGradientLayer {
	rv := objc.Send[CAGradientLayer](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g CAGradientLayer) Autorelease() CAGradientLayer {
	rv := objc.Send[CAGradientLayer](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewCAGradientLayer creates a new CAGradientLayer instance.
func NewCAGradientLayer() CAGradientLayer {
	class := getCAGradientLayerClass()
	rv := objc.Send[CAGradientLayer](objc.ID(class.class), objc.Sel("new"))
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
func NewGradientLayerWithLayer(layer objectivec.IObject) CAGradientLayer {
	instance := getCAGradientLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLayer:"), layer)
	return CAGradientLayerFromID(rv)
}

// An array of [CGColorRef] objects defining the color of each gradient stop.
// Animatable.
//
// # Discussion
// 
// Defaults to `nil`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAGradientLayer/colors
func (g CAGradientLayer) Colors() foundation.INSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("colors"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (g CAGradientLayer) SetColors(value foundation.INSArray) {
	objc.Send[struct{}](g.ID, objc.Sel("setColors:"), value)
}
// An optional array of NSNumber objects defining the location of each
// gradient stop. Animatable.
//
// # Discussion
// 
// The gradient stops are specified as values between `0` and `1`. The values
// must be monotonically increasing. If `nil`, the stops are spread uniformly
// across the range. Defaults to `nil`.
// 
// When rendered, the colors are mapped to the output color space before being
// interpolated.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAGradientLayer/locations
func (g CAGradientLayer) Locations() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("locations"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (g CAGradientLayer) SetLocations(value []foundation.NSNumber) {
	objc.Send[struct{}](g.ID, objc.Sel("setLocations:"), objectivec.IObjectSliceToNSArray(value))
}
// The end point of the gradient when drawn in the layer’s coordinate space.
// Animatable.
//
// # Discussion
// 
// The end point corresponds to the last stop of the gradient. The point is
// defined in the unit coordinate space and is then mapped to the layer’s
// bounds rectangle when drawn.
// 
// Default value is `(0.5,1.0)`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAGradientLayer/endPoint
func (g CAGradientLayer) EndPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](g.ID, objc.Sel("endPoint"))
	return corefoundation.CGPoint(rv)
}
func (g CAGradientLayer) SetEndPoint(value corefoundation.CGPoint) {
	objc.Send[struct{}](g.ID, objc.Sel("setEndPoint:"), value)
}
// The start point of the gradient when drawn in the layer’s coordinate
// space. Animatable.
//
// # Discussion
// 
// The start point corresponds to the first stop of the gradient. The point is
// defined in the unit coordinate space and is then mapped to the layer’s
// bounds rectangle when drawn.
// 
// Default value is `(0.5,0.0)`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAGradientLayer/startPoint
func (g CAGradientLayer) StartPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](g.ID, objc.Sel("startPoint"))
	return corefoundation.CGPoint(rv)
}
func (g CAGradientLayer) SetStartPoint(value corefoundation.CGPoint) {
	objc.Send[struct{}](g.ID, objc.Sel("setStartPoint:"), value)
}
// Style of gradient drawn by the layer.
//
// # Discussion
// 
// Defaults to [axial].
//
// [axial]: https://developer.apple.com/documentation/QuartzCore/CAGradientLayerType/axial
//
// See: https://developer.apple.com/documentation/QuartzCore/CAGradientLayer/type
func (g CAGradientLayer) Type() CAGradientLayerType {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("type"))
	return CAGradientLayerType(foundation.NSStringFromID(rv).String())
}
func (g CAGradientLayer) SetType(value CAGradientLayerType) {
	objc.Send[struct{}](g.ID, objc.Sel("setType:"), objc.String(string(value)))
}
// The mathematical constant pi (π), approximately equal to 3.14159.
//
// See: https://developer.apple.com/documentation/Swift/FloatingPoint/pi
func (g CAGradientLayer) Pi() objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("pi"))
	return objectivec.Object{ID: rv}
}
func (g CAGradientLayer) SetPi(value objectivec.IObject) {
	objc.Send[struct{}](g.ID, objc.Sel("setPi:"), value)
}

