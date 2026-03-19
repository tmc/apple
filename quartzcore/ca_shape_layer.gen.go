// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CAShapeLayer] class.
var (
	_CAShapeLayerClass     CAShapeLayerClass
	_CAShapeLayerClassOnce sync.Once
)

func getCAShapeLayerClass() CAShapeLayerClass {
	_CAShapeLayerClassOnce.Do(func() {
		_CAShapeLayerClass = CAShapeLayerClass{class: objc.GetClass("CAShapeLayer")}
	})
	return _CAShapeLayerClass
}

// GetCAShapeLayerClass returns the class object for CAShapeLayer.
func GetCAShapeLayerClass() CAShapeLayerClass {
	return getCAShapeLayerClass()
}

type CAShapeLayerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (cc CAShapeLayerClass) Alloc() CAShapeLayer {
	rv := objc.Send[CAShapeLayer](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// A layer that draws a cubic Bezier spline in its coordinate space.
//
// # Overview
// 
// The shape is composited between the layer’s contents and its first
// sublayer.
// 
// The shape will be drawn antialiased, and whenever possible it will be
// mapped into screen space before being rasterized to preserve resolution
// independence. However, certain kinds of image processing operations, such
// as CoreImage filters, applied to the layer or its ancestors may force
// rasterization in a local coordinate space.
// 
// The following code shows how you can build complex, composite paths and
// display them using a shape layer. In this example, a series of
// progressively transformed ellipses form a simple flower shape. The shape
// layer that displays the path has its [CAShapeLayer.FillRule] set to [evenOdd] which
// stops the overlapping “petals” from filling with the yellow
// [CAShapeLayer.FillColor].
// 
// The following figure shows the resulting shape layer.
// 
// [media-2825196]
//
// [evenOdd]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayerFillRule/evenOdd
//
// # Specifying the Shape Path
//
//   - [CAShapeLayer.Path]: The path defining the shape to be rendered. Animatable.
//   - [CAShapeLayer.SetPath]
//
// # Accessing Shape Style Properties
//
//   - [CAShapeLayer.FillColor]: The color used to fill the shape’s path. Animatable.
//   - [CAShapeLayer.SetFillColor]
//   - [CAShapeLayer.FillRule]: The fill rule used when filling the shape’s path.
//   - [CAShapeLayer.SetFillRule]
//   - [CAShapeLayer.LineCap]: Specifies the line cap style for the shape’s path.
//   - [CAShapeLayer.SetLineCap]
//   - [CAShapeLayer.LineDashPattern]: The dash pattern applied to the shape’s path when stroked.
//   - [CAShapeLayer.SetLineDashPattern]
//   - [CAShapeLayer.LineDashPhase]: The dash phase applied to the shape’s path when stroked. Animatable.
//   - [CAShapeLayer.SetLineDashPhase]
//   - [CAShapeLayer.LineJoin]: Specifies the line join style for the shape’s path.
//   - [CAShapeLayer.SetLineJoin]
//   - [CAShapeLayer.LineWidth]: Specifies the line width of the shape’s path. Animatable.
//   - [CAShapeLayer.SetLineWidth]
//   - [CAShapeLayer.MiterLimit]: The miter limit used when stroking the shape’s path. Animatable.
//   - [CAShapeLayer.SetMiterLimit]
//   - [CAShapeLayer.StrokeColor]: The color used to stroke the shape’s path. Animatable.
//   - [CAShapeLayer.SetStrokeColor]
//   - [CAShapeLayer.StrokeStart]: The relative location at which to begin stroking the path. Animatable.
//   - [CAShapeLayer.SetStrokeStart]
//   - [CAShapeLayer.StrokeEnd]: The relative location at which to stop stroking the path. Animatable.
//   - [CAShapeLayer.SetStrokeEnd]
//
// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer
type CAShapeLayer struct {
	CALayer
}

// CAShapeLayerFromID constructs a [CAShapeLayer] from an objc.ID.
//
// A layer that draws a cubic Bezier spline in its coordinate space.
func CAShapeLayerFromID(id objc.ID) CAShapeLayer {
	return CAShapeLayer{CALayer: CALayerFromID(id)}
}
// NOTE: CAShapeLayer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CAShapeLayer] class.
//
// # Specifying the Shape Path
//
//   - [ICAShapeLayer.Path]: The path defining the shape to be rendered. Animatable.
//   - [ICAShapeLayer.SetPath]
//
// # Accessing Shape Style Properties
//
//   - [ICAShapeLayer.FillColor]: The color used to fill the shape’s path. Animatable.
//   - [ICAShapeLayer.SetFillColor]
//   - [ICAShapeLayer.FillRule]: The fill rule used when filling the shape’s path.
//   - [ICAShapeLayer.SetFillRule]
//   - [ICAShapeLayer.LineCap]: Specifies the line cap style for the shape’s path.
//   - [ICAShapeLayer.SetLineCap]
//   - [ICAShapeLayer.LineDashPattern]: The dash pattern applied to the shape’s path when stroked.
//   - [ICAShapeLayer.SetLineDashPattern]
//   - [ICAShapeLayer.LineDashPhase]: The dash phase applied to the shape’s path when stroked. Animatable.
//   - [ICAShapeLayer.SetLineDashPhase]
//   - [ICAShapeLayer.LineJoin]: Specifies the line join style for the shape’s path.
//   - [ICAShapeLayer.SetLineJoin]
//   - [ICAShapeLayer.LineWidth]: Specifies the line width of the shape’s path. Animatable.
//   - [ICAShapeLayer.SetLineWidth]
//   - [ICAShapeLayer.MiterLimit]: The miter limit used when stroking the shape’s path. Animatable.
//   - [ICAShapeLayer.SetMiterLimit]
//   - [ICAShapeLayer.StrokeColor]: The color used to stroke the shape’s path. Animatable.
//   - [ICAShapeLayer.SetStrokeColor]
//   - [ICAShapeLayer.StrokeStart]: The relative location at which to begin stroking the path. Animatable.
//   - [ICAShapeLayer.SetStrokeStart]
//   - [ICAShapeLayer.StrokeEnd]: The relative location at which to stop stroking the path. Animatable.
//   - [ICAShapeLayer.SetStrokeEnd]
//
// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer
type ICAShapeLayer interface {
	ICALayer

	// Topic: Specifying the Shape Path

	// The path defining the shape to be rendered. Animatable.
	Path() coregraphics.CGPathRef
	SetPath(value coregraphics.CGPathRef)

	// Topic: Accessing Shape Style Properties

	// The color used to fill the shape’s path. Animatable.
	FillColor() coregraphics.CGColorRef
	SetFillColor(value coregraphics.CGColorRef)
	// The fill rule used when filling the shape’s path.
	FillRule() CAShapeLayerFillRule
	SetFillRule(value CAShapeLayerFillRule)
	// Specifies the line cap style for the shape’s path.
	LineCap() CAShapeLayerLineCap
	SetLineCap(value CAShapeLayerLineCap)
	// The dash pattern applied to the shape’s path when stroked.
	LineDashPattern() []foundation.NSNumber
	SetLineDashPattern(value []foundation.NSNumber)
	// The dash phase applied to the shape’s path when stroked. Animatable.
	LineDashPhase() float64
	SetLineDashPhase(value float64)
	// Specifies the line join style for the shape’s path.
	LineJoin() CAShapeLayerLineJoin
	SetLineJoin(value CAShapeLayerLineJoin)
	// Specifies the line width of the shape’s path. Animatable.
	LineWidth() float64
	SetLineWidth(value float64)
	// The miter limit used when stroking the shape’s path. Animatable.
	MiterLimit() float64
	SetMiterLimit(value float64)
	// The color used to stroke the shape’s path. Animatable.
	StrokeColor() coregraphics.CGColorRef
	SetStrokeColor(value coregraphics.CGColorRef)
	// The relative location at which to begin stroking the path. Animatable.
	StrokeStart() float64
	SetStrokeStart(value float64)
	// The relative location at which to stop stroking the path. Animatable.
	StrokeEnd() float64
	SetStrokeEnd(value float64)

	// Specifies the even-odd winding rule. Count the total number of path crossings. If the number of crossings is even, the point is outside the path. If the number of crossings is odd, the point is inside the path and the region containing it should be filled.
	EvenOdd() CAShapeLayerFillRule
}

// Init initializes the instance.
func (s CAShapeLayer) Init() CAShapeLayer {
	rv := objc.Send[CAShapeLayer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s CAShapeLayer) Autorelease() CAShapeLayer {
	rv := objc.Send[CAShapeLayer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewCAShapeLayer creates a new CAShapeLayer instance.
func NewCAShapeLayer() CAShapeLayer {
	class := getCAShapeLayerClass()
	rv := objc.Send[CAShapeLayer](objc.ID(class.class), objc.Sel("new"))
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
func NewShapeLayerWithLayer(layer objectivec.IObject) CAShapeLayer {
	instance := getCAShapeLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLayer:"), layer)
	return CAShapeLayerFromID(rv)
}

// The path defining the shape to be rendered. Animatable.
//
// # Discussion
// 
// Unlike most animatable properties, [Path] (as with all [CGPath] animatable
// properties) does not support implicit animation.
// 
// The path object may be animated using any of the concrete subclasses of
// [CAPropertyAnimation]. Paths will interpolate as a linear blend of the
// “on-line” points; “off-line” points may be interpolated
// non-linearly (e.g. to preserve continuity of the curve’s derivative). If
// the two paths have a different number of control points or segments the
// results are undefined. If the path extends outside the layer bounds it will
// not automatically be clipped to the layer, only if the normal layer masking
// rules cause that.
// 
// The default value of this property is `nil`.
//
// [CGPath]: https://developer.apple.com/documentation/CoreGraphics/CGPath
//
// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/path
func (s CAShapeLayer) Path() coregraphics.CGPathRef {
	rv := objc.Send[coregraphics.CGPathRef](s.ID, objc.Sel("path"))
	return coregraphics.CGPathRef(rv)
}
func (s CAShapeLayer) SetPath(value coregraphics.CGPathRef) {
	objc.Send[struct{}](s.ID, objc.Sel("setPath:"), value)
}
// The color used to fill the shape’s path. Animatable.
//
// # Discussion
// 
// Setting `fillColor` to `nil` results in no fill being rendered.
// 
// Default is opaque black.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/fillColor
func (s CAShapeLayer) FillColor() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](s.ID, objc.Sel("fillColor"))
	return coregraphics.CGColorRef(rv)
}
func (s CAShapeLayer) SetFillColor(value coregraphics.CGColorRef) {
	objc.Send[struct{}](s.ID, objc.Sel("setFillColor:"), value)
}
// The fill rule used when filling the shape’s path.
//
// # Discussion
// 
// The possible values are shown in [Shape Fill Mode Values]. The default is
// [nonZero]. See [Winding Rules] for examples of the two fill rules.
//
// [Shape Fill Mode Values]: https://developer.apple.com/documentation/QuartzCore/shape-fill-mode-values
// [Winding Rules]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CocoaDrawingGuide/Paths/Paths.html#//apple_ref/doc/uid/TP40003290-CH206-BAJIJJGD
// [nonZero]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayerFillRule/nonZero
//
// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/fillRule
func (s CAShapeLayer) FillRule() CAShapeLayerFillRule {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("fillRule"))
	return CAShapeLayerFillRule(foundation.NSStringFromID(rv).String())
}
func (s CAShapeLayer) SetFillRule(value CAShapeLayerFillRule) {
	objc.Send[struct{}](s.ID, objc.Sel("setFillRule:"), objc.String(string(value)))
}
// Specifies the line cap style for the shape’s path.
//
// # Discussion
// 
// The line cap style specifies the shape of the endpoints of an open path
// when stroked. The supported values are described in [Line Cap Values]. The
// following figure shows the appearance of the available line cap styles.
// 
// [media-1965770]
// 
// The default is [butt].
//
// [Line Cap Values]: https://developer.apple.com/documentation/QuartzCore/line-cap-values
// [butt]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayerLineCap/butt
//
// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineCap
func (s CAShapeLayer) LineCap() CAShapeLayerLineCap {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("lineCap"))
	return CAShapeLayerLineCap(foundation.NSStringFromID(rv).String())
}
func (s CAShapeLayer) SetLineCap(value CAShapeLayerLineCap) {
	objc.Send[struct{}](s.ID, objc.Sel("setLineCap:"), objc.String(string(value)))
}
// The dash pattern applied to the shape’s path when stroked.
//
// # Discussion
// 
// The dash pattern is specified as an array of [NSNumber] objects that
// specify the lengths of the painted segments and unpainted segments,
// respectively, of the dash pattern.
// 
// For example, passing an array with the values `[2,3]` sets a dash pattern
// that alternates between a 2-user-space-unit-long painted segment and a
// 3-user-space-unit-long unpainted segment. Passing the values `[10,5,5,5]`
// sets the pattern to a 10-unit painted segment, a 5-unit unpainted segment,
// a 5-unit painted segment, and a 5-unit unpainted segment.
// 
// Default is `nil`, a solid line.
// 
// The following code shows how how you can create three shape layers using
// the dash patterns described above. Each shape layer contains a simple path
// that describes a horizontal line.
// 
// The following figure shows three shape layers created with the code above.
// The top solid line has a `nil` [LineDashPattern], the middle has `[2,3]`
// and the bottom has `[10,5,5,5]`.
// 
// [media-2825198]
//
// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineDashPattern
func (s CAShapeLayer) LineDashPattern() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("lineDashPattern"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}
func (s CAShapeLayer) SetLineDashPattern(value []foundation.NSNumber) {
	objc.Send[struct{}](s.ID, objc.Sel("setLineDashPattern:"), objectivec.IObjectSliceToNSArray(value))
}
// The dash phase applied to the shape’s path when stroked. Animatable.
//
// # Discussion
// 
// Line dash phase specifies how far into the dash pattern the line starts.
// 
// Default is `0`.
// 
// The following code shows how you can create a “marching ant” effect by
// adding an animation to a shape layer that animates its [LineDashPhase] from
// `0` to the sum of the segment lengths of its [LineDashPattern].
//
// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineDashPhase
func (s CAShapeLayer) LineDashPhase() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("lineDashPhase"))
	return rv
}
func (s CAShapeLayer) SetLineDashPhase(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setLineDashPhase:"), value)
}
// Specifies the line join style for the shape’s path.
//
// # Discussion
// 
// The line join style specifies the shape of the joints between connected
// segments of a stroked path. The supported values are described in [Line
// Join Values]. The following figure shows the appearance of the available
// line join styles.
// 
// [media-1965771]
// 
// The default is [miter].
//
// [Line Join Values]: https://developer.apple.com/documentation/QuartzCore/line-join-values
// [miter]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayerLineJoin/miter
//
// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineJoin
func (s CAShapeLayer) LineJoin() CAShapeLayerLineJoin {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("lineJoin"))
	return CAShapeLayerLineJoin(foundation.NSStringFromID(rv).String())
}
func (s CAShapeLayer) SetLineJoin(value CAShapeLayerLineJoin) {
	objc.Send[struct{}](s.ID, objc.Sel("setLineJoin:"), objc.String(string(value)))
}
// Specifies the line width of the shape’s path. Animatable.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineWidth
func (s CAShapeLayer) LineWidth() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("lineWidth"))
	return rv
}
func (s CAShapeLayer) SetLineWidth(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setLineWidth:"), value)
}
// The miter limit used when stroking the shape’s path. Animatable.
//
// # Discussion
// 
// If the current line join style is set to [miter] (see [LineJoin]), the
// miter limit determines whether the lines should be joined with a bevel
// instead of a miter. The length of the miter is divided by the line width.
// If the result is greater than the miter limit, the path is drawn with a
// bevel.
// 
// Default is `10.0`.
//
// [miter]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayerLineJoin/miter
//
// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/miterLimit
func (s CAShapeLayer) MiterLimit() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("miterLimit"))
	return rv
}
func (s CAShapeLayer) SetMiterLimit(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMiterLimit:"), value)
}
// The color used to stroke the shape’s path. Animatable.
//
// # Discussion
// 
// Setting `strokeColor` to `nil` results in no stroke being rendered.
// 
// Default is `nil`.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/strokeColor
func (s CAShapeLayer) StrokeColor() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](s.ID, objc.Sel("strokeColor"))
	return coregraphics.CGColorRef(rv)
}
func (s CAShapeLayer) SetStrokeColor(value coregraphics.CGColorRef) {
	objc.Send[struct{}](s.ID, objc.Sel("setStrokeColor:"), value)
}
// The relative location at which to begin stroking the path. Animatable.
//
// # Discussion
// 
// The value of this property must be in the range 0.0 to 1.0. The default
// value of this property is 0.0.
// 
// Combined with the [StrokeEnd] property, this property defines the subregion
// of the path to stroke. The value in this property indicates the relative
// point along the path at which to begin stroking while the [StrokeEnd]
// property defines the end point. A value of 0.0 represents the beginning of
// the path while a value of 1.0 represents the end of the path. Values in
// between are interpreted linearly along the path length.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/strokeStart
func (s CAShapeLayer) StrokeStart() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("strokeStart"))
	return rv
}
func (s CAShapeLayer) SetStrokeStart(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setStrokeStart:"), value)
}
// The relative location at which to stop stroking the path. Animatable.
//
// # Discussion
// 
// The value of this property must be in the range 0.0 to 1.0. The default
// value of this property is 1.0.
// 
// Combined with the [StrokeStart] property, this property defines the
// subregion of the path to stroke. The value in this property indicates the
// relative point along the path at which to finish stroking while the
// [StrokeStart] property defines the starting point. A value of 0.0
// represents the beginning of the path while a value of 1.0 represents the
// end of the path. Values in between are interpreted linearly along the path
// length.
//
// See: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/strokeEnd
func (s CAShapeLayer) StrokeEnd() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("strokeEnd"))
	return rv
}
func (s CAShapeLayer) SetStrokeEnd(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setStrokeEnd:"), value)
}
// Specifies the even-odd winding rule. Count the total number of path
// crossings. If the number of crossings is even, the point is outside the
// path. If the number of crossings is odd, the point is inside the path and
// the region containing it should be filled.
//
// See: https://developer.apple.com/documentation/quartzcore/cashapelayerfillrule/evenodd
func (s CAShapeLayer) EvenOdd() CAShapeLayerFillRule {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("kCAFillRuleEvenOdd"))
	return CAShapeLayerFillRule(foundation.NSStringFromID(rv).String())
}

