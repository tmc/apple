// Code generated from Apple documentation for SpriteKit. DO NOT EDIT.

package spritekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/appkit"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [SKShapeNode] class.
var (
	_SKShapeNodeClass     SKShapeNodeClass
	_SKShapeNodeClassOnce sync.Once
)

func getSKShapeNodeClass() SKShapeNodeClass {
	_SKShapeNodeClassOnce.Do(func() {
		_SKShapeNodeClass = SKShapeNodeClass{class: objc.GetClass("SKShapeNode")}
	})
	return _SKShapeNodeClass
}

// GetSKShapeNodeClass returns the class object for SKShapeNode.
func GetSKShapeNodeClass() SKShapeNodeClass {
	return getSKShapeNodeClass()
}

type SKShapeNodeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SKShapeNodeClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SKShapeNodeClass) Alloc() SKShapeNode {
	rv := objc.Send[SKShapeNode](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// A mathematical shape that can be stroked or filled.
//
// # Overview
//
// [SKShapeNode] allows you to create onscreen graphical elements from
// mathematical points, lines, and curves. The advantage this has over
// rasterized graphics, such as those displayed by textures, is that shapes
// are rasterized dynamically at runtime to produce crisp detail and smoother
// edges.
//
// # Creating a Shape from a Path
//
//   - [SKShapeNode.Path]: The path that defines the shape.
//   - [SKShapeNode.SetPath]
//
// # Filling a Shape
//
//   - [SKShapeNode.FillColor]: The color used to fill the shape.
//   - [SKShapeNode.SetFillColor]
//   - [SKShapeNode.FillTexture]: The texture used to fill the shape.
//   - [SKShapeNode.SetFillTexture]
//
// # Stroking a Shape
//
//   - [SKShapeNode.LineWidth]: The width used to stroke the path.
//   - [SKShapeNode.SetLineWidth]
//   - [SKShapeNode.StrokeColor]: The color used to stroke the shape.
//   - [SKShapeNode.SetStrokeColor]
//   - [SKShapeNode.StrokeTexture]: The texture used to stroke the shape.
//   - [SKShapeNode.SetStrokeTexture]
//   - [SKShapeNode.GlowWidth]: A glow that extends outward from the stroked line.
//   - [SKShapeNode.SetGlowWidth]
//   - [SKShapeNode.LineCap]: The style used to render the endpoints of the stroked portion of the shape node.
//   - [SKShapeNode.SetLineCap]
//   - [SKShapeNode.LineJoin]: The junction type used when the stroked portion of the shape node is rendered.
//   - [SKShapeNode.SetLineJoin]
//   - [SKShapeNode.MiterLimit]: The miter limit to use when the line is stroked using a miter join style.
//   - [SKShapeNode.SetMiterLimit]
//   - [SKShapeNode.IsAntialiased]: A Boolean value that determines whether the stroked path is smoothed when drawn.
//   - [SKShapeNode.SetAntialiased]
//
// # Configuring Alpha Blending
//
//   - [SKShapeNode.BlendMode]: The blend mode used to blend the shape into the parent’s framebuffer.
//   - [SKShapeNode.SetBlendMode]
//
// # Controlling or Animating Sroke Length
//
//   - [SKShapeNode.LineLength]: The length of the line defined by the shape node.
//
// # Customizing Stroking or Fill Drawing
//
//   - [SKShapeNode.StrokeShader]: A custom shader used to determine the color of the stroked portion of the shape node.
//   - [SKShapeNode.SetStrokeShader]
//   - [SKShapeNode.FillShader]: A custom shader used to determine the color of the filled portion of the shape node.
//   - [SKShapeNode.SetFillShader]
//   - [SKShapeNode.AttributeValues]: The values of each attribute associated with the node’s attached shader.
//   - [SKShapeNode.SetAttributeValues]
//   - [SKShapeNode.SetValueForAttributeNamed]: Sets an attribute value for an attached shader.
//   - [SKShapeNode.ValueForAttributeNamed]: The value of a shader attribute.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode
type SKShapeNode struct {
	SKNode
}

// SKShapeNodeFromID constructs a [SKShapeNode] from an objc.ID.
//
// A mathematical shape that can be stroked or filled.
func SKShapeNodeFromID(id objc.ID) SKShapeNode {
	return SKShapeNode{SKNode: SKNodeFromID(id)}
}

// NOTE: SKShapeNode adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SKShapeNode] class.
//
// # Creating a Shape from a Path
//
//   - [ISKShapeNode.Path]: The path that defines the shape.
//   - [ISKShapeNode.SetPath]
//
// # Filling a Shape
//
//   - [ISKShapeNode.FillColor]: The color used to fill the shape.
//   - [ISKShapeNode.SetFillColor]
//   - [ISKShapeNode.FillTexture]: The texture used to fill the shape.
//   - [ISKShapeNode.SetFillTexture]
//
// # Stroking a Shape
//
//   - [ISKShapeNode.LineWidth]: The width used to stroke the path.
//   - [ISKShapeNode.SetLineWidth]
//   - [ISKShapeNode.StrokeColor]: The color used to stroke the shape.
//   - [ISKShapeNode.SetStrokeColor]
//   - [ISKShapeNode.StrokeTexture]: The texture used to stroke the shape.
//   - [ISKShapeNode.SetStrokeTexture]
//   - [ISKShapeNode.GlowWidth]: A glow that extends outward from the stroked line.
//   - [ISKShapeNode.SetGlowWidth]
//   - [ISKShapeNode.LineCap]: The style used to render the endpoints of the stroked portion of the shape node.
//   - [ISKShapeNode.SetLineCap]
//   - [ISKShapeNode.LineJoin]: The junction type used when the stroked portion of the shape node is rendered.
//   - [ISKShapeNode.SetLineJoin]
//   - [ISKShapeNode.MiterLimit]: The miter limit to use when the line is stroked using a miter join style.
//   - [ISKShapeNode.SetMiterLimit]
//   - [ISKShapeNode.IsAntialiased]: A Boolean value that determines whether the stroked path is smoothed when drawn.
//   - [ISKShapeNode.SetAntialiased]
//
// # Configuring Alpha Blending
//
//   - [ISKShapeNode.BlendMode]: The blend mode used to blend the shape into the parent’s framebuffer.
//   - [ISKShapeNode.SetBlendMode]
//
// # Controlling or Animating Sroke Length
//
//   - [ISKShapeNode.LineLength]: The length of the line defined by the shape node.
//
// # Customizing Stroking or Fill Drawing
//
//   - [ISKShapeNode.StrokeShader]: A custom shader used to determine the color of the stroked portion of the shape node.
//   - [ISKShapeNode.SetStrokeShader]
//   - [ISKShapeNode.FillShader]: A custom shader used to determine the color of the filled portion of the shape node.
//   - [ISKShapeNode.SetFillShader]
//   - [ISKShapeNode.AttributeValues]: The values of each attribute associated with the node’s attached shader.
//   - [ISKShapeNode.SetAttributeValues]
//   - [ISKShapeNode.SetValueForAttributeNamed]: Sets an attribute value for an attached shader.
//   - [ISKShapeNode.ValueForAttributeNamed]: The value of a shader attribute.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode
type ISKShapeNode interface {
	ISKNode

	// Topic: Creating a Shape from a Path

	// The path that defines the shape.
	Path() coregraphics.CGPathRef
	SetPath(value coregraphics.CGPathRef)

	// Topic: Filling a Shape

	// The color used to fill the shape.
	FillColor() appkit.NSColor
	SetFillColor(value appkit.NSColor)
	// The texture used to fill the shape.
	FillTexture() ISKTexture
	SetFillTexture(value ISKTexture)

	// Topic: Stroking a Shape

	// The width used to stroke the path.
	LineWidth() float64
	SetLineWidth(value float64)
	// The color used to stroke the shape.
	StrokeColor() appkit.NSColor
	SetStrokeColor(value appkit.NSColor)
	// The texture used to stroke the shape.
	StrokeTexture() ISKTexture
	SetStrokeTexture(value ISKTexture)
	// A glow that extends outward from the stroked line.
	GlowWidth() float64
	SetGlowWidth(value float64)
	// The style used to render the endpoints of the stroked portion of the shape node.
	LineCap() int32
	SetLineCap(value int32)
	// The junction type used when the stroked portion of the shape node is rendered.
	LineJoin() int32
	SetLineJoin(value int32)
	// The miter limit to use when the line is stroked using a miter join style.
	MiterLimit() float64
	SetMiterLimit(value float64)
	// A Boolean value that determines whether the stroked path is smoothed when drawn.
	IsAntialiased() bool
	SetAntialiased(value bool)

	// Topic: Configuring Alpha Blending

	// The blend mode used to blend the shape into the parent’s framebuffer.
	BlendMode() SKBlendMode
	SetBlendMode(value SKBlendMode)

	// Topic: Controlling or Animating Sroke Length

	// The length of the line defined by the shape node.
	LineLength() float64

	// Topic: Customizing Stroking or Fill Drawing

	// A custom shader used to determine the color of the stroked portion of the shape node.
	StrokeShader() ISKShader
	SetStrokeShader(value ISKShader)
	// A custom shader used to determine the color of the filled portion of the shape node.
	FillShader() ISKShader
	SetFillShader(value ISKShader)
	// The values of each attribute associated with the node’s attached shader.
	AttributeValues() foundation.INSDictionary
	SetAttributeValues(value foundation.INSDictionary)
	// Sets an attribute value for an attached shader.
	SetValueForAttributeNamed(value ISKAttributeValue, key string)
	// The value of a shader attribute.
	ValueForAttributeNamed(key string) ISKAttributeValue
}

// Init initializes the instance.
func (s SKShapeNode) Init() SKShapeNode {
	rv := objc.Send[SKShapeNode](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SKShapeNode) Autorelease() SKShapeNode {
	rv := objc.Send[SKShapeNode](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSKShapeNode creates a new SKShapeNode instance.
func NewSKShapeNode() SKShapeNode {
	class := getSKShapeNodeClass()
	rv := objc.Send[SKShapeNode](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a shape node with a circular path centered on the node’s origin.
//
// radius: The radius of the circle.
//
// # Return Value
//
// A new shape node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/init(circleOfRadius:)
func NewShapeNodeWithCircleOfRadius(radius float64) SKShapeNode {
	rv := objc.Send[objc.ID](objc.ID(getSKShapeNodeClass().class), objc.Sel("shapeNodeWithCircleOfRadius:"), radius)
	return SKShapeNodeFromID(rv)
}

// Called when a node is initialized from an .sks file.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(coder:)
func NewShapeNodeWithCoder(aDecoder foundation.INSCoder) SKShapeNode {
	instance := getSKShapeNodeClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	return SKShapeNodeFromID(rv)
}

// Creates a shape node with an elliptical path that fills the specified
// rectangle.
//
// rect: A rectangle, relative to the node’s origin.
//
// # Return Value
//
// A new shape node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/init(ellipseIn:)
func NewShapeNodeWithEllipseInRect(rect corefoundation.CGRect) SKShapeNode {
	rv := objc.Send[objc.ID](objc.ID(getSKShapeNodeClass().class), objc.Sel("shapeNodeWithEllipseInRect:"), rect)
	return SKShapeNodeFromID(rv)
}

// Creates a shape node with an elliptical path centered on the node’s
// origin.
//
// size: The height and width of the ellipse.
//
// # Return Value
//
// A new shape node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/init(ellipseOf:)
func NewShapeNodeWithEllipseOfSize(size corefoundation.CGSize) SKShapeNode {
	rv := objc.Send[objc.ID](objc.ID(getSKShapeNodeClass().class), objc.Sel("shapeNodeWithEllipseOfSize:"), size)
	return SKShapeNodeFromID(rv)
}

// Creates a new node by loading an archive file from the game’s main
// bundle.
//
// filename: The name of the file, without a file extension. The file must be in the
// app’s main bundle and have a `XCUIElementTypeSks` filename extension.
//
// # Return Value
//
// The unarchived node object.
//
// # Discussion
//
// If you call this method on a subclass of the [SKScene] class and the object
// in the archive is an [SKScene] object, the returned object is initialized
// as if it is a member of the subclass. You use this behavior to create scene
// layouts in the Xcode Editor and provide custom behaviors in your subclass.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:)
func NewShapeNodeWithFileNamed(filename string) SKShapeNode {
	rv := objc.Send[objc.ID](objc.ID(getSKShapeNodeClass().class), objc.Sel("nodeWithFileNamed:"), objc.String(filename))
	return SKShapeNodeFromID(rv)
}

// See: https://developer.apple.com/documentation/SpriteKit/SKNode/init(fileNamed:securelyWithClasses:)
func NewShapeNodeWithFileNamedSecurelyWithClassesAndError(filename string, classes foundation.INSSet) (SKShapeNode, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(getSKShapeNodeClass().class), objc.Sel("nodeWithFileNamed:securelyWithClasses:andError:"), objc.String(filename), classes, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return SKShapeNode{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return SKShapeNode{}, objc.ErrInitFailed
	}
	return SKShapeNodeFromID(rv), nil
}

// Creates a shape node from a Core Graphics path.
//
// path: The Core Graphics path to use. The path is relative to the node’s origin.
//
// # Return Value
//
// A new shape node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/init(path:)
func NewShapeNodeWithPath(path coregraphics.CGPathRef) SKShapeNode {
	rv := objc.Send[objc.ID](objc.ID(getSKShapeNodeClass().class), objc.Sel("shapeNodeWithPath:"), path)
	return SKShapeNodeFromID(rv)
}

// Creates a shape node from a Core Graphics path, centered around its
// position.
//
// path: The Core Graphics path to use.
//
// centered: If true, the path is translated so that the center of the path’s bounding
// box is at the node’s origin; otherwise the path is relative to the
// node’s origin.
//
// # Return Value
//
// A new shape node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/init(path:centered:)
func NewShapeNodeWithPathCentered(path coregraphics.CGPathRef, centered bool) SKShapeNode {
	rv := objc.Send[objc.ID](objc.ID(getSKShapeNodeClass().class), objc.Sel("shapeNodeWithPath:centered:"), path, centered)
	return SKShapeNodeFromID(rv)
}

// Creates a shape node from a series of points.
//
// points: An array of Core Graphics points. The points are relative to the node’s
// origin.
//
// numPoints: The number of points in the array.
//
// # Return Value
//
// A new shape node. The node is created with a path that starts at the first
// point in the array, joining each adjacent pair of points with a line
// segment.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/init(points:count:)
func NewShapeNodeWithPointsCount(points []corefoundation.CGPoint, numPoints uintptr) SKShapeNode {
	rv := objc.Send[objc.ID](objc.ID(getSKShapeNodeClass().class), objc.Sel("shapeNodeWithPoints:count:"), objc.CArray(points), numPoints)
	return SKShapeNodeFromID(rv)
}

// Creates a shape node with a rectangular path.
//
// rect: A rectangle, relative to the node’s origin.
//
// # Return Value
//
// A new shape node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/init(rect:)
func NewShapeNodeWithRect(rect corefoundation.CGRect) SKShapeNode {
	rv := objc.Send[objc.ID](objc.ID(getSKShapeNodeClass().class), objc.Sel("shapeNodeWithRect:"), rect)
	return SKShapeNodeFromID(rv)
}

// Creates a shape with a rectangular path that has rounded corners.
//
// rect: A rectangle, relative to the node’s origin.
//
// cornerRadius: The radius of the rounded corners. The radius should not be a negative
// number. The value should be no larger than half of the rectangle’s width
// or height, whichever is smaller.
//
// # Return Value
//
// A new shape node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/init(rect:cornerRadius:)
func NewShapeNodeWithRectCornerRadius(rect corefoundation.CGRect, cornerRadius float64) SKShapeNode {
	rv := objc.Send[objc.ID](objc.ID(getSKShapeNodeClass().class), objc.Sel("shapeNodeWithRect:cornerRadius:"), rect, cornerRadius)
	return SKShapeNodeFromID(rv)
}

// Creates a shape node with a rectangular path centered on the node’s
// origin.
//
// size: The size of the rectangle.
//
// # Return Value
//
// A new shape node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/init(rectOf:)
func NewShapeNodeWithRectOfSize(size corefoundation.CGSize) SKShapeNode {
	rv := objc.Send[objc.ID](objc.ID(getSKShapeNodeClass().class), objc.Sel("shapeNodeWithRectOfSize:"), size)
	return SKShapeNodeFromID(rv)
}

// Creates a shape with a rectangular path that has rounded corners centered
// on the node’s position.
//
// size: The size of the rectangle.
//
// cornerRadius: The radius of the rounded corners. The radius should not be a negative
// number. The value should be no larger than half of the rectangle’s width
// or height, whichever is smaller.
//
// # Return Value
//
// A new shape node.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/init(rectOf:cornerRadius:)
func NewShapeNodeWithRectOfSizeCornerRadius(size corefoundation.CGSize, cornerRadius float64) SKShapeNode {
	rv := objc.Send[objc.ID](objc.ID(getSKShapeNodeClass().class), objc.Sel("shapeNodeWithRectOfSize:cornerRadius:"), size, cornerRadius)
	return SKShapeNodeFromID(rv)
}

// Creates a shape node from a series of spline points.
//
// points: An array of Core Graphics points.
//
// numPoints: The number of points in the array.
//
// # Return Value
//
// A new shape node is created. The node is created with a path that starts at
// the first point in the array, joining each pair of points with a quadratic
// curve. The control points are calculated automatically based on previous
// points in the array.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/init(splinePoints:count:)
func NewShapeNodeWithSplinePointsCount(points []corefoundation.CGPoint, numPoints uintptr) SKShapeNode {
	rv := objc.Send[objc.ID](objc.ID(getSKShapeNodeClass().class), objc.Sel("shapeNodeWithSplinePoints:count:"), objc.CArray(points), numPoints)
	return SKShapeNodeFromID(rv)
}

// Sets an attribute value for an attached shader.
//
// value: An attribute value object containing the scalar or vector value to set in
// the attached shader.
//
// key: The attribute name.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/setValue(_:forAttribute:)
func (s SKShapeNode) SetValueForAttributeNamed(value ISKAttributeValue, key string) {
	objc.Send[objc.ID](s.ID, objc.Sel("setValue:forAttributeNamed:"), value, objc.String(key))
}

// The value of a shader attribute.
//
// key: The attribute name.
//
// # Return Value
//
// An attribute value object containing the scalar or vector value or `nil` if
// no such attribute exists.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/value(forAttributeNamed:)
func (s SKShapeNode) ValueForAttributeNamed(key string) ISKAttributeValue {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("valueForAttributeNamed:"), objc.String(key))
	return SKAttributeValueFromID(rv)
}

// The path that defines the shape.
//
// # Discussion
//
// The path is defined in the node’s coordinate space.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/path
func (s SKShapeNode) Path() coregraphics.CGPathRef {
	rv := objc.Send[coregraphics.CGPathRef](s.ID, objc.Sel("path"))
	return coregraphics.CGPathRef(rv)
}
func (s SKShapeNode) SetPath(value coregraphics.CGPathRef) {
	objc.Send[struct{}](s.ID, objc.Sel("setPath:"), value)
}

// The color used to fill the shape.
//
// # Discussion
//
// The default fill color is `[SKColor clearColor]`, which means the shape is
// not filled.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/fillColor
func (s SKShapeNode) FillColor() appkit.NSColor {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("fillColor"))
	return appkit.NSColorFromID(objc.ID(rv))
}
func (s SKShapeNode) SetFillColor(value appkit.NSColor) {
	objc.Send[struct{}](s.ID, objc.Sel("setFillColor:"), value)
}

// The texture used to fill the shape.
//
// # Discussion
//
// The default value is `nil`. If a fill texture is specified, the shape node
// is rendered using that texture blended with the [SKShapeNode.FillColor].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/fillTexture
func (s SKShapeNode) FillTexture() ISKTexture {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("fillTexture"))
	return SKTextureFromID(objc.ID(rv))
}
func (s SKShapeNode) SetFillTexture(value ISKTexture) {
	objc.Send[struct{}](s.ID, objc.Sel("setFillTexture:"), value)
}

// The width used to stroke the path.
//
// # Discussion
//
// A line width larger than `2.0` may cause rendering artifacts in the final
// rendered image. The default value is `1.0`.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/lineWidth
func (s SKShapeNode) LineWidth() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("lineWidth"))
	return rv
}
func (s SKShapeNode) SetLineWidth(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setLineWidth:"), value)
}

// The color used to stroke the shape.
//
// # Discussion
//
// The default stroke color is `[SKColor whiteColor]`. If you do not want to
// stroke the shape, use `[SKColor clearColor].`
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/strokeColor
func (s SKShapeNode) StrokeColor() appkit.NSColor {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("strokeColor"))
	return appkit.NSColorFromID(objc.ID(rv))
}
func (s SKShapeNode) SetStrokeColor(value appkit.NSColor) {
	objc.Send[struct{}](s.ID, objc.Sel("setStrokeColor:"), value)
}

// The texture used to stroke the shape.
//
// # Discussion
//
// The default value is `nil`. If a stroke texture is specified, the
// [SKShapeNode.StrokeColor] property is ignored and the stroked portion of
// the shape node is rendered using the texture instead.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/strokeTexture
func (s SKShapeNode) StrokeTexture() ISKTexture {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("strokeTexture"))
	return SKTextureFromID(objc.ID(rv))
}
func (s SKShapeNode) SetStrokeTexture(value ISKTexture) {
	objc.Send[struct{}](s.ID, objc.Sel("setStrokeTexture:"), value)
}

// A glow that extends outward from the stroked line.
//
// # Discussion
//
// The default value is `0.0`, which means no glow is added. The glow color is
// determined by [SKShapeNode.StrokeColor].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/glowWidth
func (s SKShapeNode) GlowWidth() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("glowWidth"))
	return rv
}
func (s SKShapeNode) SetGlowWidth(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setGlowWidth:"), value)
}

// The style used to render the endpoints of the stroked portion of the shape
// node.
//
// # Discussion
//
// The default value is [CGLineCap.butt]. See [CGLineCap].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/lineCap
//
// [CGLineCap.butt]: https://developer.apple.com/documentation/CoreGraphics/CGLineCap/butt
// [CGLineCap]: https://developer.apple.com/documentation/CoreGraphics/CGLineCap
func (s SKShapeNode) LineCap() int32 {
	rv := objc.Send[int32](s.ID, objc.Sel("lineCap"))
	return rv
}
func (s SKShapeNode) SetLineCap(value int32) {
	objc.Send[struct{}](s.ID, objc.Sel("setLineCap:"), value)
}

// The junction type used when the stroked portion of the shape node is
// rendered.
//
// # Discussion
//
// The default value is [CGLineJoin.bevel]. See [CGLineJoin].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/lineJoin
//
// [CGLineJoin.bevel]: https://developer.apple.com/documentation/CoreGraphics/CGLineJoin/bevel
// [CGLineJoin]: https://developer.apple.com/documentation/CoreGraphics/CGLineJoin
func (s SKShapeNode) LineJoin() int32 {
	rv := objc.Send[int32](s.ID, objc.Sel("lineJoin"))
	return rv
}
func (s SKShapeNode) SetLineJoin(value int32) {
	objc.Send[struct{}](s.ID, objc.Sel("setLineJoin:"), value)
}

// The miter limit to use when the line is stroked using a miter join style.
//
// # Discussion
//
// If the line join style is set to [CGLineJoin.miter], SpriteKit uses the
// miter limit to determine whether the lines should be joined with a bevel
// instead of a miter. SpriteKit divides the length of the miter by the line
// width. If the result is greater than the miter limit, SpriteKit converts
// the style to a bevel.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/miterLimit
//
// [CGLineJoin.miter]: https://developer.apple.com/documentation/CoreGraphics/CGLineJoin/miter
func (s SKShapeNode) MiterLimit() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("miterLimit"))
	return rv
}
func (s SKShapeNode) SetMiterLimit(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMiterLimit:"), value)
}

// A Boolean value that determines whether the stroked path is smoothed when
// drawn.
//
// # Discussion
//
// The default value is true.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/isAntialiased
func (s SKShapeNode) IsAntialiased() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isAntialiased"))
	return rv
}
func (s SKShapeNode) SetAntialiased(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setAntialiased:"), value)
}

// The blend mode used to blend the shape into the parent’s framebuffer.
//
// # Discussion
//
// The default value is [SKBlendMode.alpha].
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/blendMode
//
// [SKBlendMode.alpha]: https://developer.apple.com/documentation/SpriteKit/SKBlendMode/alpha
func (s SKShapeNode) BlendMode() SKBlendMode {
	rv := objc.Send[SKBlendMode](s.ID, objc.Sel("blendMode"))
	return SKBlendMode(rv)
}
func (s SKShapeNode) SetBlendMode(value SKBlendMode) {
	objc.Send[struct{}](s.ID, objc.Sel("setBlendMode:"), value)
}

// The length of the line defined by the shape node.
//
// # Discussion
//
// This property takes effect only when the shape has a stroke. The valid
// range is between `[0.1]` where one indicates that the shape is fully
// stroked, and zero indicates that the shape is not stroked at all. By
// interpolating this value over time (for example, in your scene’s
// [SKScene.Update] callback), you can animate the shape as if it were drawn
// in real time.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/lineLength
func (s SKShapeNode) LineLength() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("lineLength"))
	return rv
}

// A custom shader used to determine the color of the stroked portion of the
// shape node.
//
// # Discussion
//
// The default value is `nil`. If a `strokeShader` is specified, when the
// shape node is drawn, the shader is used to determine the output colors for
// any part of the shape node that’s stroked. SpriteKit implements many
// stroke features using a default shader, such as:
//
// - [lineCap]
// - [SKShapeNode.GlowWidth]
// - [SKShapeNode.StrokeColor]
//
// If you supply a custom value for `strokeShader`, your custom shader
// overrides the default shader which neutralizes the default features. It is
// the responsibility of your custom `strokeShader` to implement any of the
// features your shape requires.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/strokeShader
//
// [lineCap]: https://developer.apple.com/documentation/QuartzCore/CAShapeLayer/lineCap
func (s SKShapeNode) StrokeShader() ISKShader {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("strokeShader"))
	return SKShaderFromID(objc.ID(rv))
}
func (s SKShapeNode) SetStrokeShader(value ISKShader) {
	objc.Send[struct{}](s.ID, objc.Sel("setStrokeShader:"), value)
}

// A custom shader used to determine the color of the filled portion of the
// shape node.
//
// # Discussion
//
// The default value is `nil`. If a `fillShader` is specified, when the shape
// node is drawn, the shader is used to determine the output colors for any
// part of the shape node that’s fillled. SpriteKit implements many fill
// features using a default shader, such as:
//
// - Fill color. - Animations on [SKNode.Alpha]. - Light cast by
// [SKLightNode].
//
// If you supply a custom value for `fillShader`, your custom shader overrides
// the default shader which neutralizes the default features. It is the
// responsibility of your custom fillShader to implement any of the features
// your shape requires.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/fillShader
func (s SKShapeNode) FillShader() ISKShader {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("fillShader"))
	return SKShaderFromID(objc.ID(rv))
}
func (s SKShapeNode) SetFillShader(value ISKShader) {
	objc.Send[struct{}](s.ID, objc.Sel("setFillShader:"), value)
}

// The values of each attribute associated with the node’s attached shader.
//
// # Discussion
//
// All nodes have their own copy of an attribute value and therefore the
// attribute values can be different across the same [SKShader]. If instead
// you need all nodes to share the same value, use [SKUniform]. Uniforms can
// change values every frame, but uniforms cannot vary per-node like
// attributes can.
//
// See: https://developer.apple.com/documentation/SpriteKit/SKShapeNode/attributeValues
func (s SKShapeNode) AttributeValues() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("attributeValues"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (s SKShapeNode) SetAttributeValues(value foundation.INSDictionary) {
	objc.Send[struct{}](s.ID, objc.Sel("setAttributeValues:"), value)
}
