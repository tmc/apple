// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSGradient] class.
var (
	_NSGradientClass     NSGradientClass
	_NSGradientClassOnce sync.Once
)

func getNSGradientClass() NSGradientClass {
	_NSGradientClassOnce.Do(func() {
		_NSGradientClass = NSGradientClass{class: objc.GetClass("NSGradient")}
	})
	return _NSGradientClass
}

// GetNSGradientClass returns the class object for NSGradient.
func GetNSGradientClass() NSGradientClass {
	return getNSGradientClass()
}

type NSGradientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSGradientClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSGradientClass) Alloc() NSGradient {
	rv := objc.Send[NSGradient](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that can draw gradient fill colors
//
// # Overview
// 
// This class provides convenience methods for drawing radial or linear
// (axial) gradients for rectangles and [NSBezierPath] objects. It also
// supports primitive methods that let you customize the shape of the gradient
// fill. A gradient consists of two or more color changes over the range of
// the gradient shape. When creating a gradient object, you specify the colors
// and their locations relative to the start and end of the gradient. This
// combination of color and location is known as a . During drawing, the
// [NSGradient] object uses the color stop information to compute color
// changes for you and passes that information to the Quartz shading
// functions.
// 
// Because the [NSGradient] class uses Quartz shadings, drawing is handled by
// computing the colors at a given point mathematically. This technique
// results in smooth gradients regardless of the resolution of the target
// device.
// 
// For more information about gradients and their appearance, see [Gradients]
// in [Quartz 2D Programming Guide].
//
// [Gradients]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/drawingwithquartz2d/dq_shadings/dq_shadings.html#//apple_ref/doc/uid/TP30001066-CH207
// [Quartz 2D Programming Guide]: https://developer.apple.com/library/archive/documentation/GraphicsImaging/Conceptual/drawingwithquartz2d/Introduction/Introduction.html#//apple_ref/doc/uid/TP30001066
//
// # Creating a Gradient
//
//   - [NSGradient.InitWithStartingColorEndingColor]: Initializes a newly allocated gradient object with two colors.
//   - [NSGradient.InitWithColors]: Initializes a newly allocated gradient object with an array of colors.
//   - [NSGradient.InitWithColorsAtLocationsColorSpace]: Initializes a newly allocated gradient object with the specified colors, color locations, and color space.
//   - [NSGradient.InitWithCoder]: Creates a gradient from data in an unarchiver.
//
// # Drawing a Linear Gradient
//
//   - [NSGradient.DrawFromPointToPointOptions]: Draws a linear gradient between the specified start and end points.
//   - [NSGradient.DrawInRectAngle]: Fills the specified rectangle with a linear gradient.
//   - [NSGradient.DrawInBezierPathAngle]: Fills the specified path with a linear gradient.
//
// # Drawing a Radial Gradient
//
//   - [NSGradient.DrawFromCenterRadiusToCenterRadiusOptions]: Draws a radial gradient between the specified circles.
//   - [NSGradient.DrawInRectRelativeCenterPosition]: Draws a radial gradient starting at the center of the specified rectangle.
//   - [NSGradient.DrawInBezierPathRelativeCenterPosition]: Draws a radial gradient starting at the center point of the specified path.
//
// # Getting Gradient Properties
//
//   - [NSGradient.ColorSpace]: The color space of the colors associated with the gradient.
//   - [NSGradient.NumberOfColorStops]: The number of color stops associated with the gradient.
//   - [NSGradient.GetColorLocationAtIndex]: Returns information about the color stop at the specified index in the receiver’s color array.
//   - [NSGradient.InterpolatedColorAtLocation]: Returns the color of the rendered gradient at the specified relative location.
//
// See: https://developer.apple.com/documentation/AppKit/NSGradient
type NSGradient struct {
	objectivec.Object
}

// NSGradientFromID constructs a [NSGradient] from an objc.ID.
//
// An object that can draw gradient fill colors
func NSGradientFromID(id objc.ID) NSGradient {
	return NSGradient{objectivec.Object{ID: id}}
}
// NOTE: NSGradient adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSGradient] class.
//
// # Creating a Gradient
//
//   - [INSGradient.InitWithStartingColorEndingColor]: Initializes a newly allocated gradient object with two colors.
//   - [INSGradient.InitWithColors]: Initializes a newly allocated gradient object with an array of colors.
//   - [INSGradient.InitWithColorsAtLocationsColorSpace]: Initializes a newly allocated gradient object with the specified colors, color locations, and color space.
//   - [INSGradient.InitWithCoder]: Creates a gradient from data in an unarchiver.
//
// # Drawing a Linear Gradient
//
//   - [INSGradient.DrawFromPointToPointOptions]: Draws a linear gradient between the specified start and end points.
//   - [INSGradient.DrawInRectAngle]: Fills the specified rectangle with a linear gradient.
//   - [INSGradient.DrawInBezierPathAngle]: Fills the specified path with a linear gradient.
//
// # Drawing a Radial Gradient
//
//   - [INSGradient.DrawFromCenterRadiusToCenterRadiusOptions]: Draws a radial gradient between the specified circles.
//   - [INSGradient.DrawInRectRelativeCenterPosition]: Draws a radial gradient starting at the center of the specified rectangle.
//   - [INSGradient.DrawInBezierPathRelativeCenterPosition]: Draws a radial gradient starting at the center point of the specified path.
//
// # Getting Gradient Properties
//
//   - [INSGradient.ColorSpace]: The color space of the colors associated with the gradient.
//   - [INSGradient.NumberOfColorStops]: The number of color stops associated with the gradient.
//   - [INSGradient.GetColorLocationAtIndex]: Returns information about the color stop at the specified index in the receiver’s color array.
//   - [INSGradient.InterpolatedColorAtLocation]: Returns the color of the rendered gradient at the specified relative location.
//
// See: https://developer.apple.com/documentation/AppKit/NSGradient
type INSGradient interface {
	objectivec.IObject

	// Topic: Creating a Gradient

	// Initializes a newly allocated gradient object with two colors.
	InitWithStartingColorEndingColor(startingColor INSColor, endingColor INSColor) NSGradient
	// Initializes a newly allocated gradient object with an array of colors.
	InitWithColors(colorArray []NSColor) NSGradient
	// Initializes a newly allocated gradient object with the specified colors, color locations, and color space.
	InitWithColorsAtLocationsColorSpace(colorArray []NSColor, locations unsafe.Pointer, colorSpace INSColorSpace) NSGradient
	// Creates a gradient from data in an unarchiver.
	InitWithCoder(coder foundation.INSCoder) NSGradient

	// Topic: Drawing a Linear Gradient

	// Draws a linear gradient between the specified start and end points.
	DrawFromPointToPointOptions(startingPoint corefoundation.CGPoint, endingPoint corefoundation.CGPoint, options NSGradientDrawingOptions)
	// Fills the specified rectangle with a linear gradient.
	DrawInRectAngle(rect corefoundation.CGRect, angle float64)
	// Fills the specified path with a linear gradient.
	DrawInBezierPathAngle(path INSBezierPath, angle float64)

	// Topic: Drawing a Radial Gradient

	// Draws a radial gradient between the specified circles.
	DrawFromCenterRadiusToCenterRadiusOptions(startCenter corefoundation.CGPoint, startRadius float64, endCenter corefoundation.CGPoint, endRadius float64, options NSGradientDrawingOptions)
	// Draws a radial gradient starting at the center of the specified rectangle.
	DrawInRectRelativeCenterPosition(rect corefoundation.CGRect, relativeCenterPosition corefoundation.CGPoint)
	// Draws a radial gradient starting at the center point of the specified path.
	DrawInBezierPathRelativeCenterPosition(path INSBezierPath, relativeCenterPosition corefoundation.CGPoint)

	// Topic: Getting Gradient Properties

	// The color space of the colors associated with the gradient.
	ColorSpace() INSColorSpace
	// The number of color stops associated with the gradient.
	NumberOfColorStops() int
	// Returns information about the color stop at the specified index in the receiver’s color array.
	GetColorLocationAtIndex(color INSColor, location unsafe.Pointer, index int)
	// Returns the color of the rendered gradient at the specified relative location.
	InterpolatedColorAtLocation(location float64) INSColor

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (g NSGradient) Init() NSGradient {
	rv := objc.Send[NSGradient](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g NSGradient) Autorelease() NSGradient {
	rv := objc.Send[NSGradient](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSGradient creates a new NSGradient instance.
func NewNSGradient() NSGradient {
	class := getNSGradientClass()
	rv := objc.Send[NSGradient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a gradient from data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSGradient/init(coder:)
func NewGradientWithCoder(coder foundation.INSCoder) NSGradient {
	instance := getNSGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSGradientFromID(rv)
}

// Initializes a newly allocated gradient object with an array of colors.
//
// colorArray: An array of [NSColor] objects representing the colors to use to initialize
// the gradient. There must be at least two colors in the array. The first
// color is placed at location 0.0 and the last at location 1.0. If there are
// more than two colors, the additional colors are placed at evenly spaced
// intervals between the first and last colors.
//
// # Return Value
// 
// The initialized [NSGradient] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSGradient/init(colors:)
func NewGradientWithColors(colorArray []NSColor) NSGradient {
	instance := getNSGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithColors:"), objectivec.IObjectSliceToNSArray(colorArray))
	return NSGradientFromID(rv)
}

// Initializes a newly allocated gradient object with the specified colors,
// color locations, and color space.
//
// colorArray: An array of [NSColor] objects representing the colors in the gradient.
//
// locations: An array of [CGFloat] values containing the location for each color in the
// gradient. Each value must be in the range 0.0 to 1.0. There must be the
// same number of locations as are colors in the `colorArray` parameter.
//
// colorSpace: The color space to use for the gradient.
//
// # Return Value
// 
// The initialized [NSGradient] object.
//
// # Discussion
// 
// This method is the designated initializer of [NSGradient]. The colors in
// the `colorArray` parameter are converted to the specified color space if
// they are not already in that color space.
// 
// Typically, at least one color should have a location of 0.0 and one should
// have a location of 1.0. If these locations are not specified, the color at
// the closest color stop is used to fill the gap.
//
// See: https://developer.apple.com/documentation/AppKit/NSGradient/init(colors:atLocations:colorSpace:)
func NewGradientWithColorsAtLocationsColorSpace(colorArray []NSColor, locations unsafe.Pointer, colorSpace INSColorSpace) NSGradient {
	instance := getNSGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithColors:atLocations:colorSpace:"), objectivec.IObjectSliceToNSArray(colorArray), locations, colorSpace)
	return NSGradientFromID(rv)
}

// Initializes a newly allocated gradient object with two colors.
//
// startingColor: The starting color of the gradient. The location of this color is fixed at
// 0.0.
//
// endingColor: The ending color of the gradient. The location of this color is fixed at
// 1.0.
//
// # Return Value
// 
// The initialized [NSGradient] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSGradient/init(starting:ending:)
func NewGradientWithStartingColorEndingColor(startingColor INSColor, endingColor INSColor) NSGradient {
	instance := getNSGradientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithStartingColor:endingColor:"), startingColor, endingColor)
	return NSGradientFromID(rv)
}

// Initializes a newly allocated gradient object with two colors.
//
// startingColor: The starting color of the gradient. The location of this color is fixed at
// 0.0.
//
// endingColor: The ending color of the gradient. The location of this color is fixed at
// 1.0.
//
// # Return Value
// 
// The initialized [NSGradient] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSGradient/init(starting:ending:)
func (g NSGradient) InitWithStartingColorEndingColor(startingColor INSColor, endingColor INSColor) NSGradient {
	rv := objc.Send[NSGradient](g.ID, objc.Sel("initWithStartingColor:endingColor:"), startingColor, endingColor)
	return rv
}
// Initializes a newly allocated gradient object with an array of colors.
//
// colorArray: An array of [NSColor] objects representing the colors to use to initialize
// the gradient. There must be at least two colors in the array. The first
// color is placed at location 0.0 and the last at location 1.0. If there are
// more than two colors, the additional colors are placed at evenly spaced
// intervals between the first and last colors.
//
// # Return Value
// 
// The initialized [NSGradient] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSGradient/init(colors:)
func (g NSGradient) InitWithColors(colorArray []NSColor) NSGradient {
	rv := objc.Send[NSGradient](g.ID, objc.Sel("initWithColors:"), objectivec.IObjectSliceToNSArray(colorArray))
	return rv
}
// Initializes a newly allocated gradient object with the specified colors,
// color locations, and color space.
//
// colorArray: An array of [NSColor] objects representing the colors in the gradient.
//
// locations: An array of [CGFloat] values containing the location for each color in the
// gradient. Each value must be in the range 0.0 to 1.0. There must be the
// same number of locations as are colors in the `colorArray` parameter.
//
// colorSpace: The color space to use for the gradient.
//
// # Return Value
// 
// The initialized [NSGradient] object.
//
// # Discussion
// 
// This method is the designated initializer of [NSGradient]. The colors in
// the `colorArray` parameter are converted to the specified color space if
// they are not already in that color space.
// 
// Typically, at least one color should have a location of 0.0 and one should
// have a location of 1.0. If these locations are not specified, the color at
// the closest color stop is used to fill the gap.
//
// See: https://developer.apple.com/documentation/AppKit/NSGradient/init(colors:atLocations:colorSpace:)
func (g NSGradient) InitWithColorsAtLocationsColorSpace(colorArray []NSColor, locations unsafe.Pointer, colorSpace INSColorSpace) NSGradient {
	rv := objc.Send[NSGradient](g.ID, objc.Sel("initWithColors:atLocations:colorSpace:"), objectivec.IObjectSliceToNSArray(colorArray), locations, colorSpace)
	return rv
}
// Creates a gradient from data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSGradient/init(coder:)
func (g NSGradient) InitWithCoder(coder foundation.INSCoder) NSGradient {
	rv := objc.Send[NSGradient](g.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Draws a linear gradient between the specified start and end points.
//
// startingPoint: The starting point for the gradient, in the local coordinate system. The
// gradient’s first color is drawn at this point.
//
// endingPoint: The end point for the gradient, in the local coordinate system. The
// gradient’s last color is drawn at this point.
//
// options: The gradient options, if any. You can use these options to extend the
// gradient size beyond the start and end points. For more information, see
// `Gradient Drawing Options`.
//
// # Discussion
// 
// This method draws the gradient color changes along the line formed by the
// two points. The gradient fill extends perpendicularly outward from line
// until it reaches the edges of the current clipping region.
// 
// This is a primitive method used by the [NSGradient] class to draw linear
// gradients. Because this method does not perform any clipping of the
// gradient fill pattern, you must ensure that the clipping region is
// configured properly if you intend to invoke this method directly. By
// default, the clipping region is set to the current view or window in which
// drawing occurs.
//
// See: https://developer.apple.com/documentation/AppKit/NSGradient/draw(from:to:options:)
func (g NSGradient) DrawFromPointToPointOptions(startingPoint corefoundation.CGPoint, endingPoint corefoundation.CGPoint, options NSGradientDrawingOptions) {
	objc.Send[objc.ID](g.ID, objc.Sel("drawFromPoint:toPoint:options:"), startingPoint, endingPoint, options)
}
// Fills the specified rectangle with a linear gradient.
//
// rect: The rectangle to fill.
//
// angle: The angle of the linear gradient, specified in degrees. Positive values
// indicate rotation in the counter-clockwise direction relative to the
// horizontal axis.
//
// # Discussion
// 
// This convenience method draws a linear gradient inside the specified
// rectangle. The gradient is drawn so that the start and end colors are
// guaranteed to be visible in opposite corners of the rectangle. The angle of
// rotation determines which corner contains the start color; see the table
// below.
// 
// [Table data omitted]
// 
// The gradient’s color transitions occur along the line formed by the angle
// of rotation. For example, a rotation of 0 degrees results in colors
// changing from left-to-right across the rectangle, while a rotation of 90
// degrees results in colors changing from bottom to top.
// 
// The gradient drawn by this method is clipped to `rect`.
//
// See: https://developer.apple.com/documentation/AppKit/NSGradient/draw(in:angle:)-7sdyh
func (g NSGradient) DrawInRectAngle(rect corefoundation.CGRect, angle float64) {
	objc.Send[objc.ID](g.ID, objc.Sel("drawInRect:angle:"), rect, angle)
}
// Fills the specified path with a linear gradient.
//
// path: The path object to fill.
//
// angle: The angle of the linear gradient, specified in degrees. Positive values
// indicate rotation in the counter-clockwise direction relative to the
// horizontal axis.
//
// # Discussion
// 
// This convenience method behaves in a similar way to the [DrawInRectAngle]
// method, with the path object replacing the rectangle as the clipping
// region. Like the other method, the start and end colors are guaranteed to
// be visible at the farthest ends of the path.
// 
// The gradient formed by this method is clipped to `path`.
//
// See: https://developer.apple.com/documentation/AppKit/NSGradient/draw(in:angle:)-68adz
func (g NSGradient) DrawInBezierPathAngle(path INSBezierPath, angle float64) {
	objc.Send[objc.ID](g.ID, objc.Sel("drawInBezierPath:angle:"), path, angle)
}
// Draws a radial gradient between the specified circles.
//
// startCenter: The center point of the circle that represents the beginning of the
// gradient.
//
// startRadius: The radius of the circle that represents the beginning of the gradient.
//
// endCenter: The center point of the circle that represents the end of the gradient.
//
// endRadius: The radius of the circle that represents the end of the gradient.
//
// options: The gradient options, if any. You can use these options to extend the
// gradient size beyond the start and end circles. For more information, see
// [NSGradient.DrawingOptions].
// //
// [NSGradient.DrawingOptions]: https://developer.apple.com/documentation/AppKit/NSGradient/DrawingOptions
//
// # Discussion
// 
// This method draws a radial gradient pattern starting at the first circle
// and ending at the second circle. The gradient color transitions occur in
// circular bands emanating from the starting circle and ending at the second
// circle.
// 
// This is a primitive method used by the [NSGradient] class to draw radial
// gradients. Because this method does not perform any clipping of the
// gradient fill pattern, you must ensure that the clipping region is
// configured properly if you intend to invoke this method directly. By
// default, the clipping region is set to the current view or window in which
// drawing occurs.
//
// See: https://developer.apple.com/documentation/AppKit/NSGradient/draw(fromCenter:radius:toCenter:radius:options:)
func (g NSGradient) DrawFromCenterRadiusToCenterRadiusOptions(startCenter corefoundation.CGPoint, startRadius float64, endCenter corefoundation.CGPoint, endRadius float64, options NSGradientDrawingOptions) {
	objc.Send[objc.ID](g.ID, objc.Sel("drawFromCenter:radius:toCenter:radius:options:"), startCenter, startRadius, endCenter, endRadius, options)
}
// Draws a radial gradient starting at the center of the specified rectangle.
//
// rect: The rectangle to fill.
//
// relativeCenterPosition: The relative location within the rectangle to use as the center point of
// the gradient’s end circle. Each coordinate must contain a value between
// -1.0 and 1.0. A coordinate value of 0 represents the center of `rect` along
// the given axis. In the default coordinate system, a value of -1.0
// corresponds to the bottom or left edge of the rectangle and a value of 1.0
// corresponds to the top or right edge.
//
// # Discussion
// 
// The center point of the starting circle is the same as the center point of
// `rect`. The radius of the starting circle is 0, resulting in the starting
// circle being just a point.
// 
// The center point of the end circle starts at the center point of `rect` and
// is modified by the value in the `relativeCenterPosition` parameter. For
// example, if `relativeCenterPosition` contains the point (1.0, 1.0), the
// center of the end circle is located in the top-right corner of `rect`. The
// radius of the end circle is set to the smallest value that ensures `rect`
// is covered by the end circle.
// 
// The gradient formed by this method is clipped to `rect`.
//
// See: https://developer.apple.com/documentation/AppKit/NSGradient/draw(in:relativeCenterPosition:)-3a83
func (g NSGradient) DrawInRectRelativeCenterPosition(rect corefoundation.CGRect, relativeCenterPosition corefoundation.CGPoint) {
	objc.Send[objc.ID](g.ID, objc.Sel("drawInRect:relativeCenterPosition:"), rect, relativeCenterPosition)
}
// Draws a radial gradient starting at the center point of the specified path.
//
// path: The path to fill.
//
// relativeCenterPosition: The relative location within the bounding rectangle of `path` to use as the
// center point of the gradient’s end circle. Each coordinate must contain a
// value between -1.0 and 1.0. A coordinate value of 0 represents the center
// of the path’s bounding rectangle along the given axis. In the default
// coordinate system, a value of -1.0 corresponds to the bottom or left edge
// of the bounding rectangle and a value of 1.0 corresponds to the top or
// right edge.
//
// # Discussion
// 
// The center point of the starting circle is the same as the center point of
// `path`. The radius of the starting circle is 0, resulting in the starting
// circle being just a point.
// 
// The center point of the end circle starts at the center point of `path` and
// is modified by the value in the `relativeCenterPosition` parameter. For
// example, if `relativeCenterPosition` contains the point (1.0, 1.0), the
// center of the end circle is located in the top-right corner of the path’s
// bounding rectangle. The radius of the end circle is set to the smallest
// value that ensures `rect` is covered by the end circle.
// 
// The gradient formed by this method is clipped to `path`.
//
// See: https://developer.apple.com/documentation/AppKit/NSGradient/draw(in:relativeCenterPosition:)-502cc
func (g NSGradient) DrawInBezierPathRelativeCenterPosition(path INSBezierPath, relativeCenterPosition corefoundation.CGPoint) {
	objc.Send[objc.ID](g.ID, objc.Sel("drawInBezierPath:relativeCenterPosition:"), path, relativeCenterPosition)
}
// Returns information about the color stop at the specified index in the
// receiver’s color array.
//
// color: On input, a pointer to a color object. On output, the color at the
// specified index in the receiver’s color array. You may specify `nil` if
// you are not interested in this parameter.
//
// location: On input, a pointer to a floating point number. On output, contains the
// location value associated with the color. This value is between 0.0 and
// 1.0. It is used to determine the position of the color relative to the
// start and end points of the gradient. You may specify [NULL] if you are not
// interested in this parameter.
//
// index: The index of the color you want.
//
// # Discussion
// 
// This method returns the color stop information that was used to create the
// receiver. It does not return the interpolated color values at any point
// along the gradient. The location of the gradient’s first color is
// typically 0.0 and the location of the last color is typically 1.0, although
// the locations can vary depending on how the receiver was created.
//
// See: https://developer.apple.com/documentation/AppKit/NSGradient/getColor(_:location:at:)
func (g NSGradient) GetColorLocationAtIndex(color INSColor, location unsafe.Pointer, index int) {
	objc.Send[objc.ID](g.ID, objc.Sel("getColor:location:atIndex:"), color, location, index)
}
// Returns the color of the rendered gradient at the specified relative
// location.
//
// location: The location value for the color you want. This value must be between 0.0
// and 1.0. This value need not correspond to the location of one of the color
// objects used to create the gradient.
//
// # Discussion
// 
// This method does not simply return the color values used to initialize the
// receiver. Instead, it computes the value that would be drawn at the
// specified location.
// 
// The start color of the gradient is always located at 0.0 and the end color
// is always at 1.0.
//
// See: https://developer.apple.com/documentation/AppKit/NSGradient/interpolatedColor(atLocation:)
func (g NSGradient) InterpolatedColorAtLocation(location float64) INSColor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("interpolatedColorAtLocation:"), location)
	return NSColorFromID(rv)
}
func (g NSGradient) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](g.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The color space of the colors associated with the gradient.
//
// # Discussion
// 
// When the receiver is initialized, colors that do not conform to the
// receiver’s color space are converted automatically.
//
// See: https://developer.apple.com/documentation/AppKit/NSGradient/colorSpace
func (g NSGradient) ColorSpace() INSColorSpace {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("colorSpace"))
	return NSColorSpaceFromID(objc.ID(rv))
}
// The number of color stops associated with the gradient.
//
// # Discussion
// 
// Gradients must have at least two color stops: one defining the location of
// the start color and one defining the location of the end color. Gradients
// may have additional color stops located at different transition points in
// between the start and end stops.
//
// See: https://developer.apple.com/documentation/AppKit/NSGradient/numberOfColorStops
func (g NSGradient) NumberOfColorStops() int {
	rv := objc.Send[int](g.ID, objc.Sel("numberOfColorStops"))
	return rv
}

