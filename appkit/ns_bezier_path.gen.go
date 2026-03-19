// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSBezierPath] class.
var (
	_NSBezierPathClass     NSBezierPathClass
	_NSBezierPathClassOnce sync.Once
)

func getNSBezierPathClass() NSBezierPathClass {
	_NSBezierPathClassOnce.Do(func() {
		_NSBezierPathClass = NSBezierPathClass{class: objc.GetClass("NSBezierPath")}
	})
	return _NSBezierPathClass
}

// GetNSBezierPathClass returns the class object for NSBezierPath.
func GetNSBezierPathClass() NSBezierPathClass {
	return getNSBezierPathClass()
}

type NSBezierPathClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSBezierPathClass) Alloc() NSBezierPath {
	rv := objc.Send[NSBezierPath](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that can create paths using PostScript-style commands.
//
// # Overview
// 
// Paths consist of straight and curved line segments joined together. Paths
// can form recognizable shapes such as rectangles, ovals, arcs, and glyphs;
// they can also form complex polygons using either straight or curved line
// segments. A single path can be closed by connecting its two endpoints, or
// it can be left open.
// 
// An [NSBezierPath] object can contain multiple disconnected paths, whether
// they are closed or open. Each of these paths is referred to as a subpath.
// The subpaths of a Bézier path object must be manipulated as a group. The
// only way to manipulate subpaths individually is to create separate
// [NSBezierPath] objects for each.
// 
// For a given [NSBezierPath] object, you can stroke the path’s outline or
// fill the region occupied by the path. You can also use the path as a
// clipping region for views or other regions. Using methods of
// [NSBezierPath], you can also perform hit detection on the filled or stroked
// path. Hit detection is needed to implement interactive graphics, as in
// rubber banding and dragging operations.
// 
// The current graphics context is automatically saved and restored for all
// drawing operations involving Bézier path objects, so your application does
// not need to worry about the graphics settings changing across invocations.
//
// # Creating a Bézier Path
//
//   - [NSBezierPath.BezierPathByFlatteningPath]: A flattened version of the path object.
//   - [NSBezierPath.BezierPathByReversingPath]: A path containing the reversed contents of the current path object.
//
// # Constructing a Path
//
//   - [NSBezierPath.MoveToPoint]: Moves the path’s current point to the specified location.
//   - [NSBezierPath.LineToPoint]: Appends a straight line to the path.
//   - [NSBezierPath.CurveToPointControlPoint1ControlPoint2]: Adds a Bezier cubic curve to the path.
//   - [NSBezierPath.CurveToPointControlPoint]
//   - [NSBezierPath.ClosePath]: Closes the most recently added subpath.
//   - [NSBezierPath.RelativeMoveToPoint]: Moves the path’s current point to a new point whose location is the specified distance from the current point.
//   - [NSBezierPath.RelativeLineToPoint]: Appends a straight line segment to the path starting at the current point and moving towards the specified point, relative to the current location.
//   - [NSBezierPath.RelativeCurveToPointControlPoint1ControlPoint2]: Adds a Bezier cubic curve to the path from the current point to a new location, which is specified as a relative distance from the current point.
//   - [NSBezierPath.RelativeCurveToPointControlPoint]
//
// # Appending Common Shapes to a Path
//
//   - [NSBezierPath.AppendBezierPath]: Appends the contents of the specified path object to the path.
//   - [NSBezierPath.AppendBezierPathWithPointsCount]: Appends a series of line segments to the path.
//   - [NSBezierPath.AppendBezierPathWithOvalInRect]: Appends an oval path to the path, inscribing the oval in the specified rectangle.
//   - [NSBezierPath.AppendBezierPathWithArcFromPointToPointRadius]: Appends an arc to the path.
//   - [NSBezierPath.AppendBezierPathWithArcWithCenterRadiusStartAngleEndAngle]: Appends an arc of a circle to the path.
//   - [NSBezierPath.AppendBezierPathWithArcWithCenterRadiusStartAngleEndAngleClockwise]: Appends an arc of a circle to the path.
//   - [NSBezierPath.AppendBezierPathWithRect]: Appends a rectangular path to the path.
//   - [NSBezierPath.AppendBezierPathWithRoundedRectXRadiusYRadius]: Appends a rounded rectangular path to the path.
//   - [NSBezierPath.AppendBezierPathWithCGGlyphInFont]: Appends an outline of the specified glyph to the path.
//   - [NSBezierPath.AppendBezierPathWithCGGlyphsCountInFont]: Appends the outlines of the specified glyphs to the path.
//
// # Accessing a Path’s Attributes
//
//   - [NSBezierPath.WindingRule]: The winding rule used to fill the path.
//   - [NSBezierPath.SetWindingRule]
//   - [NSBezierPath.LineCapStyle]: The line cap style for the path.
//   - [NSBezierPath.SetLineCapStyle]
//   - [NSBezierPath.LineJoinStyle]: The line join style for the path.
//   - [NSBezierPath.SetLineJoinStyle]
//   - [NSBezierPath.LineWidth]: The width of stroked path lines.
//   - [NSBezierPath.SetLineWidth]
//   - [NSBezierPath.MiterLimit]: The limit at which miter joins are converted to bevel joins.
//   - [NSBezierPath.SetMiterLimit]
//   - [NSBezierPath.Flatness]: The accuracy with which curves are rendered.
//   - [NSBezierPath.SetFlatness]
//   - [NSBezierPath.GetLineDashCountPhase]: Returns the line-stroking pattern for the receiver.
//   - [NSBezierPath.SetLineDashCountPhase]: Sets the line-stroking pattern for the path.
//
// # Drawing a Path
//
//   - [NSBezierPath.Stroke]: Draws a line along the path using the current stroke color and drawing attributes.
//   - [NSBezierPath.Fill]: Paints the region enclosed by the path.
//
// # Specifying a Clipping Path
//
//   - [NSBezierPath.AddClip]: Intersects the area enclosed by the path with the clipping path of the current graphics context and makes the resulting shape the current clipping path.
//   - [NSBezierPath.SetClip]: Replaces the clipping path of the current graphics context with the area inside the path.
//
// # Performing Hit-Testing
//
//   - [NSBezierPath.ContainsPoint]: Returns a Boolean value that indicates whether the path contains the specified point.
//
// # Querying a Path
//
//   - [NSBezierPath.Bounds]: The bounding box of the path.
//   - [NSBezierPath.ControlPointBounds]: The bounding box of the path, including any control points.
//   - [NSBezierPath.CurrentPoint]: The current point (the trailing point or ending point in the most recently added segment).
//   - [NSBezierPath.Empty]: A Boolean value that indicates whether the path is empty.
//
// # Applying Transformations
//
//   - [NSBezierPath.TransformUsingAffineTransform]: Transforms all points in the path using the specified transform.
//
// # Accessing Elements of a Path
//
//   - [NSBezierPath.CGPath]
//   - [NSBezierPath.SetCGPath]
//   - [NSBezierPath.ElementCount]: The total number of path elements in the path.
//   - [NSBezierPath.ElementAtIndex]: Returns the type of path element at the specified index.
//   - [NSBezierPath.ElementAtIndexAssociatedPoints]: Gets the element type and (and optionally) the associated points for the path element at the specified index.
//   - [NSBezierPath.RemoveAllPoints]: Removes all path elements from the path, effectively clearing the path.
//   - [NSBezierPath.SetAssociatedPointsAtIndex]: Changes the points associated with the specified path element.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath
type NSBezierPath struct {
	objectivec.Object
}

// NSBezierPathFromID constructs a [NSBezierPath] from an objc.ID.
//
// An object that can create paths using PostScript-style commands.
func NSBezierPathFromID(id objc.ID) NSBezierPath {
	return NSBezierPath{objectivec.Object{ID: id}}
}
// NOTE: NSBezierPath adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSBezierPath] class.
//
// # Creating a Bézier Path
//
//   - [INSBezierPath.BezierPathByFlatteningPath]: A flattened version of the path object.
//   - [INSBezierPath.BezierPathByReversingPath]: A path containing the reversed contents of the current path object.
//
// # Constructing a Path
//
//   - [INSBezierPath.MoveToPoint]: Moves the path’s current point to the specified location.
//   - [INSBezierPath.LineToPoint]: Appends a straight line to the path.
//   - [INSBezierPath.CurveToPointControlPoint1ControlPoint2]: Adds a Bezier cubic curve to the path.
//   - [INSBezierPath.CurveToPointControlPoint]
//   - [INSBezierPath.ClosePath]: Closes the most recently added subpath.
//   - [INSBezierPath.RelativeMoveToPoint]: Moves the path’s current point to a new point whose location is the specified distance from the current point.
//   - [INSBezierPath.RelativeLineToPoint]: Appends a straight line segment to the path starting at the current point and moving towards the specified point, relative to the current location.
//   - [INSBezierPath.RelativeCurveToPointControlPoint1ControlPoint2]: Adds a Bezier cubic curve to the path from the current point to a new location, which is specified as a relative distance from the current point.
//   - [INSBezierPath.RelativeCurveToPointControlPoint]
//
// # Appending Common Shapes to a Path
//
//   - [INSBezierPath.AppendBezierPath]: Appends the contents of the specified path object to the path.
//   - [INSBezierPath.AppendBezierPathWithPointsCount]: Appends a series of line segments to the path.
//   - [INSBezierPath.AppendBezierPathWithOvalInRect]: Appends an oval path to the path, inscribing the oval in the specified rectangle.
//   - [INSBezierPath.AppendBezierPathWithArcFromPointToPointRadius]: Appends an arc to the path.
//   - [INSBezierPath.AppendBezierPathWithArcWithCenterRadiusStartAngleEndAngle]: Appends an arc of a circle to the path.
//   - [INSBezierPath.AppendBezierPathWithArcWithCenterRadiusStartAngleEndAngleClockwise]: Appends an arc of a circle to the path.
//   - [INSBezierPath.AppendBezierPathWithRect]: Appends a rectangular path to the path.
//   - [INSBezierPath.AppendBezierPathWithRoundedRectXRadiusYRadius]: Appends a rounded rectangular path to the path.
//   - [INSBezierPath.AppendBezierPathWithCGGlyphInFont]: Appends an outline of the specified glyph to the path.
//   - [INSBezierPath.AppendBezierPathWithCGGlyphsCountInFont]: Appends the outlines of the specified glyphs to the path.
//
// # Accessing a Path’s Attributes
//
//   - [INSBezierPath.WindingRule]: The winding rule used to fill the path.
//   - [INSBezierPath.SetWindingRule]
//   - [INSBezierPath.LineCapStyle]: The line cap style for the path.
//   - [INSBezierPath.SetLineCapStyle]
//   - [INSBezierPath.LineJoinStyle]: The line join style for the path.
//   - [INSBezierPath.SetLineJoinStyle]
//   - [INSBezierPath.LineWidth]: The width of stroked path lines.
//   - [INSBezierPath.SetLineWidth]
//   - [INSBezierPath.MiterLimit]: The limit at which miter joins are converted to bevel joins.
//   - [INSBezierPath.SetMiterLimit]
//   - [INSBezierPath.Flatness]: The accuracy with which curves are rendered.
//   - [INSBezierPath.SetFlatness]
//   - [INSBezierPath.GetLineDashCountPhase]: Returns the line-stroking pattern for the receiver.
//   - [INSBezierPath.SetLineDashCountPhase]: Sets the line-stroking pattern for the path.
//
// # Drawing a Path
//
//   - [INSBezierPath.Stroke]: Draws a line along the path using the current stroke color and drawing attributes.
//   - [INSBezierPath.Fill]: Paints the region enclosed by the path.
//
// # Specifying a Clipping Path
//
//   - [INSBezierPath.AddClip]: Intersects the area enclosed by the path with the clipping path of the current graphics context and makes the resulting shape the current clipping path.
//   - [INSBezierPath.SetClip]: Replaces the clipping path of the current graphics context with the area inside the path.
//
// # Performing Hit-Testing
//
//   - [INSBezierPath.ContainsPoint]: Returns a Boolean value that indicates whether the path contains the specified point.
//
// # Querying a Path
//
//   - [INSBezierPath.Bounds]: The bounding box of the path.
//   - [INSBezierPath.ControlPointBounds]: The bounding box of the path, including any control points.
//   - [INSBezierPath.CurrentPoint]: The current point (the trailing point or ending point in the most recently added segment).
//   - [INSBezierPath.Empty]: A Boolean value that indicates whether the path is empty.
//
// # Applying Transformations
//
//   - [INSBezierPath.TransformUsingAffineTransform]: Transforms all points in the path using the specified transform.
//
// # Accessing Elements of a Path
//
//   - [INSBezierPath.CGPath]
//   - [INSBezierPath.SetCGPath]
//   - [INSBezierPath.ElementCount]: The total number of path elements in the path.
//   - [INSBezierPath.ElementAtIndex]: Returns the type of path element at the specified index.
//   - [INSBezierPath.ElementAtIndexAssociatedPoints]: Gets the element type and (and optionally) the associated points for the path element at the specified index.
//   - [INSBezierPath.RemoveAllPoints]: Removes all path elements from the path, effectively clearing the path.
//   - [INSBezierPath.SetAssociatedPointsAtIndex]: Changes the points associated with the specified path element.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath
type INSBezierPath interface {
	objectivec.IObject

	// Topic: Creating a Bézier Path

	// A flattened version of the path object.
	BezierPathByFlatteningPath() INSBezierPath
	// A path containing the reversed contents of the current path object.
	BezierPathByReversingPath() INSBezierPath

	// Topic: Constructing a Path

	// Moves the path’s current point to the specified location.
	MoveToPoint(point corefoundation.CGPoint)
	// Appends a straight line to the path.
	LineToPoint(point corefoundation.CGPoint)
	// Adds a Bezier cubic curve to the path.
	CurveToPointControlPoint1ControlPoint2(endPoint corefoundation.CGPoint, controlPoint1 corefoundation.CGPoint, controlPoint2 corefoundation.CGPoint)
	CurveToPointControlPoint(endPoint corefoundation.CGPoint, controlPoint corefoundation.CGPoint)
	// Closes the most recently added subpath.
	ClosePath()
	// Moves the path’s current point to a new point whose location is the specified distance from the current point.
	RelativeMoveToPoint(point corefoundation.CGPoint)
	// Appends a straight line segment to the path starting at the current point and moving towards the specified point, relative to the current location.
	RelativeLineToPoint(point corefoundation.CGPoint)
	// Adds a Bezier cubic curve to the path from the current point to a new location, which is specified as a relative distance from the current point.
	RelativeCurveToPointControlPoint1ControlPoint2(endPoint corefoundation.CGPoint, controlPoint1 corefoundation.CGPoint, controlPoint2 corefoundation.CGPoint)
	RelativeCurveToPointControlPoint(endPoint corefoundation.CGPoint, controlPoint corefoundation.CGPoint)

	// Topic: Appending Common Shapes to a Path

	// Appends the contents of the specified path object to the path.
	AppendBezierPath(path INSBezierPath)
	// Appends a series of line segments to the path.
	AppendBezierPathWithPointsCount(points []foundation.NSPoint, count int)
	// Appends an oval path to the path, inscribing the oval in the specified rectangle.
	AppendBezierPathWithOvalInRect(rect corefoundation.CGRect)
	// Appends an arc to the path.
	AppendBezierPathWithArcFromPointToPointRadius(point1 corefoundation.CGPoint, point2 corefoundation.CGPoint, radius float64)
	// Appends an arc of a circle to the path.
	AppendBezierPathWithArcWithCenterRadiusStartAngleEndAngle(center corefoundation.CGPoint, radius float64, startAngle float64, endAngle float64)
	// Appends an arc of a circle to the path.
	AppendBezierPathWithArcWithCenterRadiusStartAngleEndAngleClockwise(center corefoundation.CGPoint, radius float64, startAngle float64, endAngle float64, clockwise bool)
	// Appends a rectangular path to the path.
	AppendBezierPathWithRect(rect corefoundation.CGRect)
	// Appends a rounded rectangular path to the path.
	AppendBezierPathWithRoundedRectXRadiusYRadius(rect corefoundation.CGRect, xRadius float64, yRadius float64)
	// Appends an outline of the specified glyph to the path.
	AppendBezierPathWithCGGlyphInFont(glyph coregraphics.CGFontIndex, font NSFont)
	// Appends the outlines of the specified glyphs to the path.
	AppendBezierPathWithCGGlyphsCountInFont(glyphs []coregraphics.CGFontIndex, count int, font NSFont)

	// Topic: Accessing a Path’s Attributes

	// The winding rule used to fill the path.
	WindingRule() NSWindingRule
	SetWindingRule(value NSWindingRule)
	// The line cap style for the path.
	LineCapStyle() NSLineCapStyle
	SetLineCapStyle(value NSLineCapStyle)
	// The line join style for the path.
	LineJoinStyle() NSLineJoinStyle
	SetLineJoinStyle(value NSLineJoinStyle)
	// The width of stroked path lines.
	LineWidth() float64
	SetLineWidth(value float64)
	// The limit at which miter joins are converted to bevel joins.
	MiterLimit() float64
	SetMiterLimit(value float64)
	// The accuracy with which curves are rendered.
	Flatness() float64
	SetFlatness(value float64)
	// Returns the line-stroking pattern for the receiver.
	GetLineDashCountPhase(pattern []float64, count unsafe.Pointer, phase unsafe.Pointer)
	// Sets the line-stroking pattern for the path.
	SetLineDashCountPhase(pattern []float64, count int, phase float64)

	// Topic: Drawing a Path

	// Draws a line along the path using the current stroke color and drawing attributes.
	Stroke()
	// Paints the region enclosed by the path.
	Fill()

	// Topic: Specifying a Clipping Path

	// Intersects the area enclosed by the path with the clipping path of the current graphics context and makes the resulting shape the current clipping path.
	AddClip()
	// Replaces the clipping path of the current graphics context with the area inside the path.
	SetClip()

	// Topic: Performing Hit-Testing

	// Returns a Boolean value that indicates whether the path contains the specified point.
	ContainsPoint(point corefoundation.CGPoint) bool

	// Topic: Querying a Path

	// The bounding box of the path.
	Bounds() corefoundation.CGRect
	// The bounding box of the path, including any control points.
	ControlPointBounds() corefoundation.CGRect
	// The current point (the trailing point or ending point in the most recently added segment).
	CurrentPoint() corefoundation.CGPoint
	// A Boolean value that indicates whether the path is empty.
	Empty() bool

	// Topic: Applying Transformations

	// Transforms all points in the path using the specified transform.
	TransformUsingAffineTransform(transform foundation.NSAffineTransform)

	// Topic: Accessing Elements of a Path

	CGPath() coregraphics.CGPathRef
	SetCGPath(value coregraphics.CGPathRef)
	// The total number of path elements in the path.
	ElementCount() int
	// Returns the type of path element at the specified index.
	ElementAtIndex(index int) NSBezierPathElement
	// Gets the element type and (and optionally) the associated points for the path element at the specified index.
	ElementAtIndexAssociatedPoints(index int, points foundation.NSPoint) NSBezierPathElement
	// Removes all path elements from the path, effectively clearing the path.
	RemoveAllPoints()
	// Changes the points associated with the specified path element.
	SetAssociatedPointsAtIndex(points foundation.NSPoint, index int)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (b NSBezierPath) Init() NSBezierPath {
	rv := objc.Send[NSBezierPath](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b NSBezierPath) Autorelease() NSBezierPath {
	rv := objc.Send[NSBezierPath](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSBezierPath creates a new NSBezierPath instance.
func NewNSBezierPath() NSBezierPath {
	class := getNSBezierPathClass()
	rv := objc.Send[NSBezierPath](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/init(cgPath:)
func NewBezierPathWithCGPath(cgPath coregraphics.CGPathRef) NSBezierPath {
	rv := objc.Send[objc.ID](objc.ID(getNSBezierPathClass().class), objc.Sel("bezierPathWithCGPath:"), cgPath)
	return NSBezierPathFromID(rv)
}

// Creates and returns a new Bézier path object initialized with an oval path
// inscribed in the specified rectangle.
//
// rect: The rectangle in which to inscribe an oval.
//
// # Return Value
// 
// An [NSBezierPath] new path object with the oval path.
//
// # Discussion
// 
// If the `aRect` parameter specifies a square, the inscribed path is a
// circle. The path is constructed by starting in the lower-right quadrant of
// the rectangle and adding arc segments counterclockwise to complete the
// oval.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/init(ovalIn:)
func NewBezierPathWithOvalInRect(rect corefoundation.CGRect) NSBezierPath {
	rv := objc.Send[objc.ID](objc.ID(getNSBezierPathClass().class), objc.Sel("bezierPathWithOvalInRect:"), rect)
	return NSBezierPathFromID(rv)
}

// Creates and returns a new Bézier path object initialized with a
// rectangular path.
//
// rect: The rectangle describing the path to create.
//
// # Return Value
// 
// A new path object with the rectangular path.
//
// # Discussion
// 
// The path is constructed by starting at the origin of `aRect` and adding
// line segments in a counterclockwise direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/init(rect:)
func NewBezierPathWithRect(rect corefoundation.CGRect) NSBezierPath {
	rv := objc.Send[objc.ID](objc.ID(getNSBezierPathClass().class), objc.Sel("bezierPathWithRect:"), rect)
	return NSBezierPathFromID(rv)
}

// Creates and returns a new Bézier path object initialized with a rounded
// rectangular path.
//
// rect: The rectangle that defines the basic shape of the path.
//
// xRadius: The radius of each corner oval along the x-axis. Values larger than half
// the rectangle’s width are clamped to half the width.
//
// yRadius: The radius of each corner oval along the y-axis. Values larger than half
// the rectangle’s height are clamped to half the height.
//
// # Return Value
// 
// A new path object with the rounded rectangular path.
//
// # Discussion
// 
// The path is constructed in a counter-clockwise direction, starting at the
// top-left corner of the rectangle. If either one of the radius parameters
// contains the value `0.0`, the returned path is a plain rectangle without
// rounded corners.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/init(roundedRect:xRadius:yRadius:)
func NewBezierPathWithRoundedRectXRadiusYRadius(rect corefoundation.CGRect, xRadius float64, yRadius float64) NSBezierPath {
	rv := objc.Send[objc.ID](objc.ID(getNSBezierPathClass().class), objc.Sel("bezierPathWithRoundedRect:xRadius:yRadius:"), rect, xRadius, yRadius)
	return NSBezierPathFromID(rv)
}

// Moves the path’s current point to the specified location.
//
// point: A point in the current coordinate system.
//
// # Discussion
// 
// This method implicitly closes the current subpath (if any) and sets the
// current point to the value in `aPoint`. When closing the previous subpath,
// this method does not cause a line to be created from the first and last
// points in the subpath.
// 
// For many path operations, you must invoke this method before issuing any
// commands that cause a line or curve segment to be drawn.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/move(to:)
func (b NSBezierPath) MoveToPoint(point corefoundation.CGPoint) {
	objc.Send[objc.ID](b.ID, objc.Sel("moveToPoint:"), point)
}
// Appends a straight line to the path.
//
// point: The destination point of the line segment, specified in the current
// coordinate system.
//
// # Discussion
// 
// This method creates a straight line segment starting at the current point
// and ending at the point specified by the `aPoint` parameter. The current
// point is the last point in the receiver’s most recently added segment.
// 
// You must set the path’s current point (using the [MoveToPoint] method or
// through the creation of a preceding line or curve segment) before you
// invoke this method. If the path is empty, this method raises an
// [genericException] exception.
//
// [genericException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/genericException
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/line(to:)
func (b NSBezierPath) LineToPoint(point corefoundation.CGPoint) {
	objc.Send[objc.ID](b.ID, objc.Sel("lineToPoint:"), point)
}
// Adds a Bezier cubic curve to the path.
//
// endPoint: The destination point of the curve segment, specified in the current
// coordinate system
//
// controlPoint1: The point that determines the shape of the curve near the current point.
//
// controlPoint2: The point that determines the shape of the curve near the destination
// point.
//
// # Discussion
// 
// You must set the path’s current point (using the [MoveToPoint] method or
// through the creation of a preceding line or curve segment) before you
// invoke this method. If the path is empty, this method raises an
// [genericException] exception.
//
// [genericException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/genericException
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/curve(to:controlPoint1:controlPoint2:)
func (b NSBezierPath) CurveToPointControlPoint1ControlPoint2(endPoint corefoundation.CGPoint, controlPoint1 corefoundation.CGPoint, controlPoint2 corefoundation.CGPoint) {
	objc.Send[objc.ID](b.ID, objc.Sel("curveToPoint:controlPoint1:controlPoint2:"), endPoint, controlPoint1, controlPoint2)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/curve(to:controlPoint:)
func (b NSBezierPath) CurveToPointControlPoint(endPoint corefoundation.CGPoint, controlPoint corefoundation.CGPoint) {
	objc.Send[objc.ID](b.ID, objc.Sel("curveToPoint:controlPoint:"), endPoint, controlPoint)
}
// Closes the most recently added subpath.
//
// # Discussion
// 
// This method closes the current subpath by creating a line segment between
// the first and last points in the subpath. This method subsequently updates
// the current point to the end of the newly created line segment, which is
// also the first point in the now closed subpath.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/close()
func (b NSBezierPath) ClosePath() {
	objc.Send[objc.ID](b.ID, objc.Sel("closePath"))
}
// Moves the path’s current point to a new point whose location is the
// specified distance from the current point.
//
// point: A point whose coordinates are interpreted as a relative offset from the
// current point.
//
// # Discussion
// 
// This method implicitly closes the current subpath (if any) and updates the
// location of the current point. For example, if the current point is (1, 1)
// and `aPoint` contains the value (1, 2), the previous subpath would be
// closed and the current point would become (2, 3). When closing the previous
// subpath, this method does not cause a line to be created from the first and
// last points in the subpath.
// 
// You must set the path’s current point (using the [MoveToPoint] method or
// through the creation of a preceding line or curve segment) before you
// invoke this method. If the path is empty, this method raises an
// [genericException] exception.
//
// [genericException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/genericException
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/relativeMove(to:)
func (b NSBezierPath) RelativeMoveToPoint(point corefoundation.CGPoint) {
	objc.Send[objc.ID](b.ID, objc.Sel("relativeMoveToPoint:"), point)
}
// Appends a straight line segment to the path starting at the current point
// and moving towards the specified point, relative to the current location.
//
// point: A point whose coordinates are interpreted as a relative offset from the
// current point.
//
// # Discussion
// 
// The destination point is relative to the current point. For example, if the
// current point is (1, 1) and `aPoint` contains the value (1, 2), a line
// segment is created between the points (1, 1) and (2, 3).
// 
// You must set the path’s current point (using the [MoveToPoint] method or
// through the creation of a preceding line or curve segment) before you
// invoke this method. If the path is empty, this method raises an
// [genericException] exception.
//
// [genericException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/genericException
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/relativeLine(to:)
func (b NSBezierPath) RelativeLineToPoint(point corefoundation.CGPoint) {
	objc.Send[objc.ID](b.ID, objc.Sel("relativeLineToPoint:"), point)
}
// Adds a Bezier cubic curve to the path from the current point to a new
// location, which is specified as a relative distance from the current point.
//
// endPoint: The destination point of the curve segment, interpreted as a relative
// offset from the current point.
//
// controlPoint1: The point that determines the shape of the curve near the current point,
// interpreted as a relative offset from the current point.
//
// controlPoint2: The point that determines the shape of the curve near the destination
// point, interpreted as a relative offset from the current point.
//
// # Discussion
// 
// You must set the path’s current point (using the [MoveToPoint] method or
// through the creation of a preceding line or curve segment) before you
// invoke this method. If the path is empty, this method raises an
// [genericException] exception.
//
// [genericException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/genericException
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/relativeCurve(to:controlPoint1:controlPoint2:)
func (b NSBezierPath) RelativeCurveToPointControlPoint1ControlPoint2(endPoint corefoundation.CGPoint, controlPoint1 corefoundation.CGPoint, controlPoint2 corefoundation.CGPoint) {
	objc.Send[objc.ID](b.ID, objc.Sel("relativeCurveToPoint:controlPoint1:controlPoint2:"), endPoint, controlPoint1, controlPoint2)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/relativeCurve(to:controlPoint:)
func (b NSBezierPath) RelativeCurveToPointControlPoint(endPoint corefoundation.CGPoint, controlPoint corefoundation.CGPoint) {
	objc.Send[objc.ID](b.ID, objc.Sel("relativeCurveToPoint:controlPoint:"), endPoint, controlPoint)
}
// Appends the contents of the specified path object to the path.
//
// path: The path to add to the receiver.
//
// # Discussion
// 
// This method adds the commands used to create `aPath` to the end of the
// receiver’s path. This method does not explicitly try to connect the
// subpaths in the two objects, although the operations in `aPath` may still
// cause that effect.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/append(_:)
func (b NSBezierPath) AppendBezierPath(path INSBezierPath) {
	objc.Send[objc.ID](b.ID, objc.Sel("appendBezierPath:"), path)
}
// Appends a series of line segments to the path.
//
// points: A C-style array of [NSPoint] data types, each of which contains the end
// point of the next line segment.
//
// count: The number of points in the `points` parameter.
//
// # Discussion
// 
// This method interprets the points as a set of connected line segments. If
// the current path contains an open subpath, a line is created from the last
// point in that subpath to the first point in the points array. If the
// current path is empty, the first point in the points array is used to set
// the starting point of the line segments. Subsequent line segments are added
// using the remaining points in the array.
// 
// This method does not close the path that is created. If you wish to create
// a closed path, you must do so by explicitly invoking the receiver’s
// [ClosePath] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/appendPoints(_:count:)
func (b NSBezierPath) AppendBezierPathWithPointsCount(points []foundation.NSPoint, count int) {
	objc.Send[objc.ID](b.ID, objc.Sel("appendBezierPathWithPoints:count:"), objc.CArray(points), count)
}
// Appends an oval path to the path, inscribing the oval in the specified
// rectangle.
//
// rect: The rectangle in which to inscribe the oval.
//
// # Discussion
// 
// Before adding the oval, this method moves the current point, which
// implicitly closes the current subpath. If the `aRect` parameter specifies a
// square, the inscribed path is a circle. The path is constructed by starting
// in the lower-right quadrant of the rectangle and adding arc segments
// counterclockwise to complete the oval.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/appendOval(in:)
func (b NSBezierPath) AppendBezierPathWithOvalInRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](b.ID, objc.Sel("appendBezierPathWithOvalInRect:"), rect)
}
// Appends an arc to the path.
//
// point1: The middle point of the angle.
//
// point2: The end point of the angle.
//
// radius: The radius of the circle inscribed in the angle.
//
// # Discussion
// 
// The created arc is defined by a circle inscribed inside the angle specified
// by three points: the current point, the `fromPoint` parameter, and the
// `toPoint` parameter (in that order). The arc itself lies on the perimeter
// of the circle, whose radius is specified by the `radius` parameter. The arc
// is drawn between the two points of the circle that are tangent to the two
// legs of the angle.
// 
// The arc usually does not contain the points in the `fromPoint` and
// `toPoint` parameters. If the starting point of the arc does not coincide
// with the current point, a line is drawn between the two points. The
// starting point of the arc lies on the line defined by the current point and
// the `fromPoint` parameter.
// 
// You must set the path’s current point (using the [MoveToPoint] method or
// through the creation of a preceding line or curve segment) before you
// invoke this method. If the path is empty, this method raises an
// [genericException] exception.
// 
// Depending on the length of the arc, this method may add multiple connected
// curve segments to the path.
//
// [genericException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/genericException
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/appendArc(from:to:radius:)
func (b NSBezierPath) AppendBezierPathWithArcFromPointToPointRadius(point1 corefoundation.CGPoint, point2 corefoundation.CGPoint, radius float64) {
	objc.Send[objc.ID](b.ID, objc.Sel("appendBezierPathWithArcFromPoint:toPoint:radius:"), point1, point2, radius)
}
// Appends an arc of a circle to the path.
//
// center: Specifies the center point of the circle used to define the arc.
//
// radius: Specifies the radius of the circle used to define the arc.
//
// startAngle: Specifies the starting angle of the arc, measured in degrees
// counterclockwise from the x-axis.
//
// endAngle: Specifies the end angle of the arc, measured in degrees counterclockwise
// from the x-axis.
//
// # Discussion
// 
// The created arc lies on the perimeter of the circle, between the angles
// specified by the `startAngle` and `endAngle` parameters. The arc is drawn
// in a counterclockwise direction. If the receiver’s path is empty, this
// method sets the current point to the beginning of the arc before adding the
// arc segment. If the receiver’s path is not empty, a line is drawn from
// the current point to the starting point of the arc.
// 
// Depending on the length of the arc, this method may add multiple connected
// curve segments to the path.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/appendArc(withCenter:radius:startAngle:endAngle:)
func (b NSBezierPath) AppendBezierPathWithArcWithCenterRadiusStartAngleEndAngle(center corefoundation.CGPoint, radius float64, startAngle float64, endAngle float64) {
	objc.Send[objc.ID](b.ID, objc.Sel("appendBezierPathWithArcWithCenter:radius:startAngle:endAngle:"), center, radius, startAngle, endAngle)
}
// Appends an arc of a circle to the path.
//
// center: Specifies the center point of the circle used to define the arc.
//
// radius: Specifies the radius of the circle used to define the arc.
//
// startAngle: Specifies the starting angle of the arc, measured in degrees
// counterclockwise from the x-axis.
//
// endAngle: Specifies the end angle of the arc, measured in degrees counterclockwise
// from the x-axis.
//
// clockwise: [true] if you want the arc to be drawn in a clockwise direction; otherwise
// [false] to draw the arc in a counterclockwise direction.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The created arc lies on the perimeter of the circle, between the angles
// specified by the `startAngle` and `endAngle` parameters. The arc is drawn
// in the direction indicated by the `clockwise` parameter. If the
// receiver’s path is empty, this method sets the current point to the
// beginning of the arc before adding the arc segment. If the receiver’s
// path is not empty, a line is drawn from the current point to the starting
// point of the arc.
// 
// Depending on the length of the arc, this method may add multiple connected
// curve segments to the path.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/appendArc(withCenter:radius:startAngle:endAngle:clockwise:)
func (b NSBezierPath) AppendBezierPathWithArcWithCenterRadiusStartAngleEndAngleClockwise(center corefoundation.CGPoint, radius float64, startAngle float64, endAngle float64, clockwise bool) {
	objc.Send[objc.ID](b.ID, objc.Sel("appendBezierPathWithArcWithCenter:radius:startAngle:endAngle:clockwise:"), center, radius, startAngle, endAngle, clockwise)
}
// Appends a rectangular path to the path.
//
// rect: The rectangle describing the path to create.
//
// # Discussion
// 
// Before adding the rectangle, this method moves the current point to the
// origin of the rectangle, which implicitly closes the current subpath (if
// any). The path is constructed by starting at the origin of `aRect` and
// adding line segments in a counterclockwise direction. The final segment is
// added using a [ClosePath] message.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/appendRect(_:)
func (b NSBezierPath) AppendBezierPathWithRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](b.ID, objc.Sel("appendBezierPathWithRect:"), rect)
}
// Appends a rounded rectangular path to the path.
//
// rect: The rectangle that defines the basic shape of the path.
//
// xRadius: The radius of each corner oval along the x-axis. Values larger than half
// the rectangle’s width are clamped to half the width.
//
// yRadius: The radius of each corner oval along the y-axis. Values larger than half
// the rectangle’s height are clamped to half the height.
//
// # Discussion
// 
// The path is constructed in a counter-clockwise direction, starting at the
// top-left corner of the rectangle. If either one of the radius parameters
// contains the value `0.0`, the returned path is a plain rectangle without
// rounded corners.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/appendRoundedRect(_:xRadius:yRadius:)
func (b NSBezierPath) AppendBezierPathWithRoundedRectXRadiusYRadius(rect corefoundation.CGRect, xRadius float64, yRadius float64) {
	objc.Send[objc.ID](b.ID, objc.Sel("appendBezierPathWithRoundedRect:xRadius:yRadius:"), rect, xRadius, yRadius)
}
// Appends an outline of the specified glyph to the path.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/append(withCGGlyph:in:)
func (b NSBezierPath) AppendBezierPathWithCGGlyphInFont(glyph coregraphics.CGFontIndex, font NSFont) {
	objc.Send[objc.ID](b.ID, objc.Sel("appendBezierPathWithCGGlyph:inFont:"), glyph, font)
}
// Appends the outlines of the specified glyphs to the path.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/append(withCGGlyphs:count:in:)
func (b NSBezierPath) AppendBezierPathWithCGGlyphsCountInFont(glyphs []coregraphics.CGFontIndex, count int, font NSFont) {
	objc.Send[objc.ID](b.ID, objc.Sel("appendBezierPathWithCGGlyphs:count:inFont:"), objc.CArray(glyphs), count, font)
}
// Returns the line-stroking pattern for the receiver.
//
// pattern: On input, a C-style array of floating point values, or `nil` if you do not
// want the pattern values. On output, this array contains the lengths
// (measured in points) of the line segments and gaps in the pattern. The
// values in the array alternate, starting with the first line segment length,
// followed by the first gap length, followed by the second line segment
// length, and so on.
//
// count: On input, a pointer to an integer or `nil` if you do not want the number of
// pattern entries. On output, the number of entries written to `pattern`.
//
// phase: On input, a pointer to a floating point value or `nil` if you do not want
// the phase. On output, this value contains the offset at which to start
// drawing the pattern, measured in points along the dashed-line pattern. For
// example, a phase of 6 in the pattern 5-2-3-2 would cause drawing to begin
// in the middle of the first gap.
//
// # Discussion
// 
// The array in the `pattern` parameter must be large enough to hold all of
// the returned values in the pattern. If you are not sure how many values
// there might be, you can call this method twice. The first time you call it,
// do not pass a value for `pattern` but use the returned value in `count` to
// allocate an array of floating-point numbers that you can then pass in the
// second time.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/getLineDash(_:count:phase:)
func (b NSBezierPath) GetLineDashCountPhase(pattern []float64, count unsafe.Pointer, phase unsafe.Pointer) {
	objc.Send[objc.ID](b.ID, objc.Sel("getLineDash:count:phase:"), objc.CArray(pattern), count, phase)
}
// Sets the line-stroking pattern for the path.
//
// pattern: A C-style array of floating point values that contains the lengths
// (measured in points) of the line segments and gaps in the pattern. The
// values in the array alternate, starting with the first line segment length,
// followed by the first gap length, followed by the second line segment
// length, and so on
//
// count: The number of values in `pattern`.
//
// phase: The offset at which to start drawing the pattern, measured in points along
// the dashed-line pattern. For example, a phase of 6 in the pattern 5-2-3-2
// would cause drawing to begin in the middle of the first gap
//
// # Discussion
// 
// For example, to produce a supermarket coupon type of dashed line:
// 
// In the above example, if you set `phase` to 6.0, the line dash would begin
// exactly six units into `pattern`, which would start the pattern in the
// middle of the first gap.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/setLineDash(_:count:phase:)
func (b NSBezierPath) SetLineDashCountPhase(pattern []float64, count int, phase float64) {
	objc.Send[objc.ID](b.ID, objc.Sel("setLineDash:count:phase:"), objc.CArray(pattern), count, phase)
}
// Draws a line along the path using the current stroke color and drawing
// attributes.
//
// # Discussion
// 
// The drawn line is centered on the path with its sides parallel to the path
// segment. This method uses the current drawing attributes associated with
// the receiver. If a particular attribute is not set for the receiver, this
// method uses the corresponding default attribute.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/stroke()
func (b NSBezierPath) Stroke() {
	objc.Send[objc.ID](b.ID, objc.Sel("stroke"))
}
// Paints the region enclosed by the path.
//
// # Discussion
// 
// This method fills the path using the current fill color and the
// receiver’s current winding rule. If the path contains any open subpaths,
// this method implicitly closes them before painting the fill region.
// 
// The painted region includes the pixels right up to, but not including, the
// path line itself. For paths with large line widths, this can result in
// overlap between the fill region and the stroked path (which is itself
// centered on the path line).
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/fill()
func (b NSBezierPath) Fill() {
	objc.Send[objc.ID](b.ID, objc.Sel("fill"))
}
// Intersects the area enclosed by the path with the clipping path of the
// current graphics context and makes the resulting shape the current clipping
// path.
//
// # Discussion
// 
// This method uses the current winding rule to determine the clipping shape
// of the receiver. This method does not affect the receiver’s path.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/addClip()
func (b NSBezierPath) AddClip() {
	objc.Send[objc.ID](b.ID, objc.Sel("addClip"))
}
// Replaces the clipping path of the current graphics context with the area
// inside the path.
//
// # Discussion
// 
// You should avoid using this method as a way of adjusting the clipping path,
// as it may expand the clipping path beyond the bounds set by the enclosing
// view. If you do use this method, be sure to save the graphics state prior
// to modifying the clipping path and restore the graphics state when you are
// done.
// 
// This method uses the current winding rule to determine the clipping shape
// of the receiver. This method does not affect the receiver’s path.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/setClip()
func (b NSBezierPath) SetClip() {
	objc.Send[objc.ID](b.ID, objc.Sel("setClip"))
}
// Returns a Boolean value that indicates whether the path contains the
// specified point.
//
// point: The point to test against the path, specified in the path object’s
// coordinate system.
//
// # Return Value
// 
// [true] if the path’s enclosed area contains the specified point;
// otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method checks the point against the path itself and the area it
// encloses. When determining hits in the enclosed area, this method uses the
// non-zero winding rule ([NSNonZeroWindingRule]). It does not take into
// account the line width used to stroke the path.
//
// [NSNonZeroWindingRule]: https://developer.apple.com/documentation/AppKit/NSNonZeroWindingRule
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/contains(_:)
func (b NSBezierPath) ContainsPoint(point corefoundation.CGPoint) bool {
	rv := objc.Send[bool](b.ID, objc.Sel("containsPoint:"), point)
	return rv
}
// Transforms all points in the path using the specified transform.
//
// transform: The transform to apply to the path.
//
// # Discussion
// 
// This method applies the transform to the path’s points immediately. The
// following code translates a line from 0,0 to 100,100 to a line from 10,10
// to 110,110.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/transform(using:)
func (b NSBezierPath) TransformUsingAffineTransform(transform foundation.NSAffineTransform) {
	objc.Send[objc.ID](b.ID, objc.Sel("transformUsingAffineTransform:"), transform)
}
// Returns the type of path element at the specified index.
//
// index: The index of the desired path element.
//
// # Return Value
// 
// The type of the path element. For a list of constants, see
// [NSBezierPath.ElementType].
//
// [NSBezierPath.ElementType]: https://developer.apple.com/documentation/AppKit/NSBezierPath/ElementType
//
// # Discussion
// 
// Path elements describe the commands used to define a path and include basic
// commands such as moving to a specific point, creating a line segment,
// creating a curve, or closing the path. The elements are stored in the order
// of their execution.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/element(at:)
func (b NSBezierPath) ElementAtIndex(index int) NSBezierPathElement {
	rv := objc.Send[NSBezierPathElement](b.ID, objc.Sel("elementAtIndex:"), index)
	return NSBezierPathElement(rv)
}
// Gets the element type and (and optionally) the associated points for the
// path element at the specified index.
//
// index: The index of the desired path element.
//
// points: On input, a C-style array containing up to three [NSPoint] data types, or
// [NULL] if you do not want the points. On output, the data points associated
// with the specified path element.
//
// # Return Value
// 
// The type of the path element. For a list of constants, see
// [NSBezierPath.ElementType].
//
// [NSBezierPath.ElementType]: https://developer.apple.com/documentation/AppKit/NSBezierPath/ElementType
//
// # Discussion
// 
// If you specify a value for the points parameter, your array must be large
// enough to hold the number of points for the given path element. Move, close
// path, and line segment commands return one point. Curve operations return
// three points.
// 
// For curve operations, the order of the points is controlPoint1
// (`points`[0]), controlPoint2 (`points`[1]), endPoint (`points`[2]).
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/element(at:associatedPoints:)
func (b NSBezierPath) ElementAtIndexAssociatedPoints(index int, points foundation.NSPoint) NSBezierPathElement {
	rv := objc.Send[NSBezierPathElement](b.ID, objc.Sel("elementAtIndex:associatedPoints:"), index, points)
	return NSBezierPathElement(rv)
}
// Removes all path elements from the path, effectively clearing the path.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/removeAllPoints()
func (b NSBezierPath) RemoveAllPoints() {
	objc.Send[objc.ID](b.ID, objc.Sel("removeAllPoints"))
}
// Changes the points associated with the specified path element.
//
// points: A C-style array containing up to three [NSPoint] data types. This parameter
// must contain the correct number of points for the path element at the
// specified index. Move, close path, and line segment commands require one
// point. Curve operations require three points.
//
// index: The index of the path element you want to modify.
//
// # Discussion
// 
// You can use this method to change the points associated with a path quickly
// and without recreating the path. You cannot use this method to change the
// type of the path element.
// 
// The following example shows you how you would modify the point associated
// with a line path element. The path created by this example results in a
// path with two elements. The first path element specifies a move to point
// (0, 0) while the second creates a line to point (100, 100). It then changes
// the line to go only to the point (50,50) using this method:
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/setAssociatedPoints(_:at:)
func (b NSBezierPath) SetAssociatedPointsAtIndex(points foundation.NSPoint, index int) {
	objc.Send[objc.ID](b.ID, objc.Sel("setAssociatedPoints:atIndex:"), points, index)
}
func (b NSBezierPath) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](b.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Fills the specified rectangular path with the current fill color.
//
// rect: A rectangle in the current coordinate system.
//
// # Discussion
// 
// This method fills the specified region immediately. This method uses the
// compositing operation returned by the `compositingOperation` method of
// [NSGraphicsContext].
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/fill(_:)
func (_NSBezierPathClass NSBezierPathClass) FillRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](objc.ID(_NSBezierPathClass.class), objc.Sel("fillRect:"), rect)
}
// Strokes the path of the specified rectangle using the current stroke color
// and the default drawing attributes.
//
// rect: A rectangle in the current coordinate system.
//
// # Discussion
// 
// The path is drawn beginning at the rectangle’s origin and proceeding in a
// counterclockwise direction. This method strokes the specified path
// immediately.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/stroke(_:)
func (_NSBezierPathClass NSBezierPathClass) StrokeRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](objc.ID(_NSBezierPathClass.class), objc.Sel("strokeRect:"), rect)
}
// Strokes a line between two points using the current stroke color and the
// default drawing attributes.
//
// point1: The starting point of the line.
//
// point2: The ending point of the line.
//
// # Discussion
// 
// This method strokes the specified path immediately.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/strokeLine(from:to:)
func (_NSBezierPathClass NSBezierPathClass) StrokeLineFromPointToPoint(point1 corefoundation.CGPoint, point2 corefoundation.CGPoint) {
	objc.Send[objc.ID](objc.ID(_NSBezierPathClass.class), objc.Sel("strokeLineFromPoint:toPoint:"), point1, point2)
}
// Draws a set of packed glyphs at the specified point in the current
// coordinate system.
//
// packedGlyphs: A C-style array containing one or more [CGGlyph] data types terminated by a
// [NULL] character.
//
// point: The starting point at which to draw the glyphs.
//
// # Discussion
// 
// This method draws the glyphs immediately.
// 
// You should avoid using this method directly. Instead, use the
// [appendGlyph(_:in:)] and [appendGlyphs(_:count:in:)] methods to create a
// path with one or more glyphs.
//
// [appendGlyph(_:in:)]: https://developer.apple.com/documentation/AppKit/NSBezierPath/appendGlyph(_:in:)
// [appendGlyphs(_:count:in:)]: https://developer.apple.com/documentation/AppKit/NSBezierPath/appendGlyphs(_:count:in:)
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/drawPackedGlyphs(_:at:)
func (_NSBezierPathClass NSBezierPathClass) DrawPackedGlyphsAtPoint(packedGlyphs []byte, point corefoundation.CGPoint) {
	objc.Send[objc.ID](objc.ID(_NSBezierPathClass.class), objc.Sel("drawPackedGlyphs:atPoint:"), unsafe.Pointer(unsafe.SliceData(packedGlyphs)), point)
}
// Intersects the specified rectangle with the clipping path of the current
// graphics context and makes the resulting shape the current clipping path.
//
// rect: The rectangle to intersect with the current clipping path.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/clip(_:)
func (_NSBezierPathClass NSBezierPathClass) ClipRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](objc.ID(_NSBezierPathClass.class), objc.Sel("clipRect:"), rect)
}

// A flattened version of the path object.
//
// # Discussion
// 
// Flattening a path converts all curved line segments into straight line
// approximations. The granularity of the approximations is controlled by the
// path’s current flatness value, which is set using [DefaultFlatness].
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/flattened
func (b NSBezierPath) BezierPathByFlatteningPath() INSBezierPath {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("bezierPathByFlatteningPath"))
	return NSBezierPathFromID(objc.ID(rv))
}
// A path containing the reversed contents of the current path object.
//
// # Discussion
// 
// The reversed path does not necessarily have a different appearance when
// rendered. Instead, it changes the direction in which path segments are
// drawn. For example, reversing the path of a rectangle (whose line segments
// are normally drawn starting at the origin and proceeding in a
// counterclockwise direction) causes its line segments to be drawn in a
// clockwise direction instead. Drawing a reversed path could affect the
// appearance of a filled pattern, depending on the pattern and the fill rule
// in use.
// 
// The path in this property is created by reversing each whole or partial
// subpath in the path object individually.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/reversed
func (b NSBezierPath) BezierPathByReversingPath() INSBezierPath {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("bezierPathByReversingPath"))
	return NSBezierPathFromID(objc.ID(rv))
}
// The winding rule used to fill the path.
//
// # Discussion
// 
// This value may be either [NSNonZeroWindingRule] or [NSEvenOddWindingRule].
// This value overrides the default value returned by the [DefaultWindingRule]
// method.
// 
// For more information on how winding rules affect the appearance of filled
// paths, see [Cocoa Drawing Guide].
//
// [Cocoa Drawing Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/CocoaDrawingGuide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40003290
// [NSEvenOddWindingRule]: https://developer.apple.com/documentation/AppKit/NSEvenOddWindingRule
// [NSNonZeroWindingRule]: https://developer.apple.com/documentation/AppKit/NSNonZeroWindingRule
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/windingRule-swift.property
func (b NSBezierPath) WindingRule() NSWindingRule {
	rv := objc.Send[NSWindingRule](b.ID, objc.Sel("windingRule"))
	return NSWindingRule(rv)
}
func (b NSBezierPath) SetWindingRule(value NSWindingRule) {
	objc.Send[struct{}](b.ID, objc.Sel("setWindingRule:"), value)
}
// The line cap style for the path.
//
// # Discussion
// 
// The line cap style specifies the shape of the endpoints on an open path
// when stroked. The default value of this property is the value returned by
// the [DefaultLineCapStyle] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/lineCapStyle-swift.property
func (b NSBezierPath) LineCapStyle() NSLineCapStyle {
	rv := objc.Send[NSLineCapStyle](b.ID, objc.Sel("lineCapStyle"))
	return NSLineCapStyle(rv)
}
func (b NSBezierPath) SetLineCapStyle(value NSLineCapStyle) {
	objc.Send[struct{}](b.ID, objc.Sel("setLineCapStyle:"), value)
}
// The line join style for the path.
//
// # Discussion
// 
// The line join style specifies the shape of the joints between connected
// segments of a stroked path. The default value of this property is the value
// returned by the [DefaultLineJoinStyle] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/lineJoinStyle-swift.property
func (b NSBezierPath) LineJoinStyle() NSLineJoinStyle {
	rv := objc.Send[NSLineJoinStyle](b.ID, objc.Sel("lineJoinStyle"))
	return NSLineJoinStyle(rv)
}
func (b NSBezierPath) SetLineJoinStyle(value NSLineJoinStyle) {
	objc.Send[struct{}](b.ID, objc.Sel("setLineJoinStyle:"), value)
}
// The width of stroked path lines.
//
// # Discussion
// 
// The line width defines the thickness of the receiver’s stroked path. A
// width of 0 is interpreted as the thinnest line that can be rendered on a
// particular device. The actual rendered line width may vary from the
// specified width by as much as 2 device pixels, depending on the position of
// the line with respect to the pixel grid and the current anti-aliasing
// settings. The width of the line may also be affected by scaling factors
// specified in the current transformation matrix of the active graphics
// context.
// 
// The default value of this property is the value returned by the
// [DefaultLineWidth] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/lineWidth
func (b NSBezierPath) LineWidth() float64 {
	rv := objc.Send[float64](b.ID, objc.Sel("lineWidth"))
	return rv
}
func (b NSBezierPath) SetLineWidth(value float64) {
	objc.Send[struct{}](b.ID, objc.Sel("setLineWidth:"), value)
}
// The limit at which miter joins are converted to bevel joins.
//
// # Discussion
// 
// The miter limit helps you avoid spikes at the junction of two line segments
// connected by a miter join ([NSMiterLineJoinStyle]). If the ratio of the
// miter length—the diagonal length of the miter join—to the line
// thickness exceeds the miter limit, the joint is converted to a bevel join.
// 
// The default value of this property is the value returned by the
// [DefaultMiterLimit] method.
//
// [NSMiterLineJoinStyle]: https://developer.apple.com/documentation/AppKit/NSMiterLineJoinStyle
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/miterLimit
func (b NSBezierPath) MiterLimit() float64 {
	rv := objc.Send[float64](b.ID, objc.Sel("miterLimit"))
	return rv
}
func (b NSBezierPath) SetMiterLimit(value float64) {
	objc.Send[struct{}](b.ID, objc.Sel("setMiterLimit:"), value)
}
// The accuracy with which curves are rendered.
//
// # Discussion
// 
// The flatness value specifies the accuracy (or smoothness) with which curves
// are rendered. It is also the maximum error tolerance (measured in pixels)
// for rendering curves, where smaller numbers give smoother curves at the
// expense of more computation. The exact interpretation may vary slightly on
// different rendering devices.
// 
// The default value of this property is the value returned by the
// [DefaultFlatness] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/flatness
func (b NSBezierPath) Flatness() float64 {
	rv := objc.Send[float64](b.ID, objc.Sel("flatness"))
	return rv
}
func (b NSBezierPath) SetFlatness(value float64) {
	objc.Send[struct{}](b.ID, objc.Sel("setFlatness:"), value)
}
// The bounding box of the path.
//
// # Discussion
// 
// This property contains the rectangle that encloses the path of the
// receiver. If the path contains curve segments, the bounding box encloses
// the curve but may not enclose the control points used to calculate the
// curve.
// 
// If the path is empty, accessing this property raises [genericException].
//
// [genericException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/genericException
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/bounds
func (b NSBezierPath) Bounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](b.ID, objc.Sel("bounds"))
	return corefoundation.CGRect(rv)
}
// The bounding box of the path, including any control points.
//
// # Discussion
// 
// This property contains the rectangle that encloses the receiver’s path.
// If the path contains curve segments, the bounding box encloses the control
// points of the curves as well as the curves themselves.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/controlPointBounds
func (b NSBezierPath) ControlPointBounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](b.ID, objc.Sel("controlPointBounds"))
	return corefoundation.CGRect(rv)
}
// The current point (the trailing point or ending point in the most recently
// added segment).
//
// # Discussion
// 
// This property contains the point from which the next drawn line or curve
// segment begins. If the path is empty, accessing this property raises
// [genericException].
//
// [genericException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/genericException
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/currentPoint
func (b NSBezierPath) CurrentPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](b.ID, objc.Sel("currentPoint"))
	return corefoundation.CGPoint(rv)
}
// A Boolean value that indicates whether the path is empty.
//
// # Discussion
// 
// The value of this property is [true] if the path contains no elements, or
// [false] if it contains at least one element.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/isEmpty
func (b NSBezierPath) Empty() bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isEmpty"))
	return rv
}
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/cgPath
func (b NSBezierPath) CGPath() coregraphics.CGPathRef {
	rv := objc.Send[coregraphics.CGPathRef](b.ID, objc.Sel("CGPath"))
	return coregraphics.CGPathRef(rv)
}
func (b NSBezierPath) SetCGPath(value coregraphics.CGPathRef) {
	objc.Send[struct{}](b.ID, objc.Sel("setCGPath:"), value)
}
// The total number of path elements in the path.
//
// # Discussion
// 
// Each element type corresponds to one of the operations described in Path
// Elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/elementCount
func (b NSBezierPath) ElementCount() int {
	rv := objc.Send[int](b.ID, objc.Sel("elementCount"))
	return rv
}

// Returns the default winding rule used to fill all paths.
//
// # Return Value
// 
// The current default winding rule or [NSNonZeroWindingRule] if no default
// rule has been set. This value may be either [NSNonZeroWindingRule] or
// [NSEvenOddWindingRule].
//
// [NSEvenOddWindingRule]: https://developer.apple.com/documentation/AppKit/NSEvenOddWindingRule
// [NSNonZeroWindingRule]: https://developer.apple.com/documentation/AppKit/NSNonZeroWindingRule
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultWindingRule
func (_NSBezierPathClass NSBezierPathClass) DefaultWindingRule() NSWindingRule {
	rv := objc.Send[NSWindingRule](objc.ID(_NSBezierPathClass.class), objc.Sel("defaultWindingRule"))
	return NSWindingRule(rv)
}
func (_NSBezierPathClass NSBezierPathClass) SetDefaultWindingRule(value NSWindingRule) {
	objc.Send[struct{}](objc.ID(_NSBezierPathClass.class), objc.Sel("setDefaultWindingRule:"), value)
}
// Returns the default line cap style for all paths.
//
// # Return Value
// 
// The default line cap style or [NSButtLineCapStyle] if no other style has
// been set. For a list of values, see Constants.
// 
// # Discussion
// 
// The default line cap style can be overridden for individual paths by
// setting a custom style for that path using the [NSBezierPath] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultLineCapStyle
func (_NSBezierPathClass NSBezierPathClass) DefaultLineCapStyle() NSLineCapStyle {
	rv := objc.Send[NSLineCapStyle](objc.ID(_NSBezierPathClass.class), objc.Sel("defaultLineCapStyle"))
	return NSLineCapStyle(rv)
}
func (_NSBezierPathClass NSBezierPathClass) SetDefaultLineCapStyle(value NSLineCapStyle) {
	objc.Send[struct{}](objc.ID(_NSBezierPathClass.class), objc.Sel("setDefaultLineCapStyle:"), value)
}
// Returns the default line join style for all paths.
//
// # Return Value
// 
// The default line join style or [NSMiterLineJoinStyle] if no other value has
// been set. For a list of values, see Constants.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultLineJoinStyle
func (_NSBezierPathClass NSBezierPathClass) DefaultLineJoinStyle() NSLineJoinStyle {
	rv := objc.Send[NSLineJoinStyle](objc.ID(_NSBezierPathClass.class), objc.Sel("defaultLineJoinStyle"))
	return NSLineJoinStyle(rv)
}
func (_NSBezierPathClass NSBezierPathClass) SetDefaultLineJoinStyle(value NSLineJoinStyle) {
	objc.Send[struct{}](objc.ID(_NSBezierPathClass.class), objc.Sel("setDefaultLineJoinStyle:"), value)
}
// Returns the default line width for the all paths.
//
// # Return Value
// 
// The default line width, measured in points in the user coordinate space, or
// 1.0 if no other value has been set.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultLineWidth
func (_NSBezierPathClass NSBezierPathClass) DefaultLineWidth() float64 {
	rv := objc.Send[float64](objc.ID(_NSBezierPathClass.class), objc.Sel("defaultLineWidth"))
	return rv
}
func (_NSBezierPathClass NSBezierPathClass) SetDefaultLineWidth(value float64) {
	objc.Send[struct{}](objc.ID(_NSBezierPathClass.class), objc.Sel("setDefaultLineWidth:"), value)
}
// Returns the default miter limit for all paths.
//
// # Return Value
// 
// The default miter limit for all paths, or 10.0 if no other value has been
// set.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultMiterLimit
func (_NSBezierPathClass NSBezierPathClass) DefaultMiterLimit() float64 {
	rv := objc.Send[float64](objc.ID(_NSBezierPathClass.class), objc.Sel("defaultMiterLimit"))
	return rv
}
func (_NSBezierPathClass NSBezierPathClass) SetDefaultMiterLimit(value float64) {
	objc.Send[struct{}](objc.ID(_NSBezierPathClass.class), objc.Sel("setDefaultMiterLimit:"), value)
}
// The default flatness value for all paths.
//
// # Return Value
// 
// The default value for determining the smoothness of curved paths, or 0.6 if
// no other value has been set.
//
// See: https://developer.apple.com/documentation/AppKit/NSBezierPath/defaultFlatness
func (_NSBezierPathClass NSBezierPathClass) DefaultFlatness() float64 {
	rv := objc.Send[float64](objc.ID(_NSBezierPathClass.class), objc.Sel("defaultFlatness"))
	return rv
}
func (_NSBezierPathClass NSBezierPathClass) SetDefaultFlatness(value float64) {
	objc.Send[struct{}](objc.ID(_NSBezierPathClass.class), objc.Sel("setDefaultFlatness:"), value)
}

