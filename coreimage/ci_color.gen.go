// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [CIColor] class.
var (
	_CIColorClass     CIColorClass
	_CIColorClassOnce sync.Once
)

func getCIColorClass() CIColorClass {
	_CIColorClassOnce.Do(func() {
		_CIColorClass = CIColorClass{class: objc.GetClass("CIColor")}
	})
	return _CIColorClass
}

// GetCIColorClass returns the class object for CIColor.
func GetCIColorClass() CIColorClass {
	return getCIColorClass()
}

type CIColorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (cc CIColorClass) Class() objc.Class {
	return cc.class
}

// Alloc allocates memory for a new instance of the class.
func (cc CIColorClass) Alloc() CIColor {
	rv := objc.Send[CIColor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// The Core Image class that defines a color object.
//
// # Overview
// 
// Use [CIColor] instances in conjunction with other Core Image classes, such
// as [CIFilter] and [CIKernel]. Many of the built-in Core Image filters have
// one or more [CIColor] inputs that you can set to affect the filter’s
// behavior.
// 
// # Color Model
// 
// A color is defined as a N-dimensional model where each dimension’s color
// component is represented by intensity values. A color component may also be
// referred to as a color channel. An RGB color model, for example, is
// three-dimensional and the red, green, and blue component intensities define
// each unique color.
// 
// # Color Space
// 
// A color is also defined by a color space that locates the axes of
// N-dimensional model within the greater volume of human perceivable colors.
// Core Image uses [CGColorSpace] instances to specify a variety of different
// color spaces such as sRGB, P3, BT.2020, etc. The [CGColorSpace] also
// defines if the color space is coded linearly or in a non-linear perceptual
// curve. (For more information on [CGColorSpace] see [CGColorSpace])
// 
// # Color Range
// 
// Standard dynamic range (SDR) color color component values range from `0.0`
// to `1.0`, with `0.0` representing an 0% of that component and `1.0`
// representing 100%. In contrast, high dynamic range (HDR) color values can
// be less than `0.0` (for more saturation) or greater than `1.0` (for more
// brightness).
// 
// # Color Opacity
// 
// [CIColor] instances also have an alpha component, which represents the
// opacity of the color, with 0.0 meaning completely transparent and 1.0
// meaning completely opaque. If a color does not have an explicit alpha
// component, Core Image assumes that the alpha component equals 1.0. With
// [CIColor] that color components values are not premultiplied. So for
// example, a semi-transparent pure red [CIColor] is represented by RGB
// `1.0,0.0,0.0` and A `0.5`. In contrast color components values in [CIImage]
// buffers or read in [CIKernel] samplers are premultiplied by default.
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
//
// # Initializing Color Objects
//
//   - [CIColor.InitWithCGColor]: Create a Core Image color object with a Core Graphics color object.
//   - [CIColor.InitWithColor]
//   - [CIColor.InitWithRedGreenBlueAlpha]: Initialize a Core Image color object in the sRGB color space with the specified red, green, blue, and alpha component values.
//   - [CIColor.InitWithRedGreenBlueColorSpace]: Initialize a Core Image color object with the specified red, green, and blue component values as measured in the specified color space.
//   - [CIColor.InitWithRedGreenBlueAlphaColorSpace]: Initialize a Core Image color object with the specified red, green, and blue component values as measured in the specified color space.
//
// # Getting Color Components
//
//   - [CIColor.ColorSpace]: Returns the [CGColorSpace] associated with the color
//   - [CIColor.Components]: Return a pointer to an array of [CGFloat] values including alpha.
//   - [CIColor.NumberOfComponents]: Returns the color components of the color including alpha.
//   - [CIColor.Red]: Returns the unpremultiplied red component of the color.
//   - [CIColor.Green]: Returns the unpremultiplied green component of the color.
//   - [CIColor.Blue]: Returns the unpremultiplied blue component of the color.
//   - [CIColor.Alpha]: Returns the alpha value of the color.
//   - [CIColor.StringRepresentation]: Returns a formatted string with the unpremultiplied color and alpha components of the color.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor
type CIColor struct {
	objectivec.Object
}

// CIColorFromID constructs a [CIColor] from an objc.ID.
//
// The Core Image class that defines a color object.
func CIColorFromID(id objc.ID) CIColor {
	return CIColor{objectivec.Object{ID: id}}
}
// NOTE: CIColor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [CIColor] class.
//
// # Initializing Color Objects
//
//   - [ICIColor.InitWithCGColor]: Create a Core Image color object with a Core Graphics color object.
//   - [ICIColor.InitWithColor]
//   - [ICIColor.InitWithRedGreenBlueAlpha]: Initialize a Core Image color object in the sRGB color space with the specified red, green, blue, and alpha component values.
//   - [ICIColor.InitWithRedGreenBlueColorSpace]: Initialize a Core Image color object with the specified red, green, and blue component values as measured in the specified color space.
//   - [ICIColor.InitWithRedGreenBlueAlphaColorSpace]: Initialize a Core Image color object with the specified red, green, and blue component values as measured in the specified color space.
//
// # Getting Color Components
//
//   - [ICIColor.ColorSpace]: Returns the [CGColorSpace] associated with the color
//   - [ICIColor.Components]: Return a pointer to an array of [CGFloat] values including alpha.
//   - [ICIColor.NumberOfComponents]: Returns the color components of the color including alpha.
//   - [ICIColor.Red]: Returns the unpremultiplied red component of the color.
//   - [ICIColor.Green]: Returns the unpremultiplied green component of the color.
//   - [ICIColor.Blue]: Returns the unpremultiplied blue component of the color.
//   - [ICIColor.Alpha]: Returns the alpha value of the color.
//   - [ICIColor.StringRepresentation]: Returns a formatted string with the unpremultiplied color and alpha components of the color.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor
type ICIColor interface {
	objectivec.IObject

	// Topic: Initializing Color Objects

	// Create a Core Image color object with a Core Graphics color object.
	InitWithCGColor(color coregraphics.CGColorRef) CIColor
	InitWithColor(color objectivec.IObject) CIColor
	// Initialize a Core Image color object in the sRGB color space with the specified red, green, blue, and alpha component values.
	InitWithRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) CIColor
	// Initialize a Core Image color object with the specified red, green, and blue component values as measured in the specified color space.
	InitWithRedGreenBlueColorSpace(red float64, green float64, blue float64, colorSpace coregraphics.CGColorSpaceRef) CIColor
	// Initialize a Core Image color object with the specified red, green, and blue component values as measured in the specified color space.
	InitWithRedGreenBlueAlphaColorSpace(red float64, green float64, blue float64, alpha float64, colorSpace coregraphics.CGColorSpaceRef) CIColor

	// Topic: Getting Color Components

	// Returns the [CGColorSpace] associated with the color
	ColorSpace() coregraphics.CGColorSpaceRef
	// Return a pointer to an array of [CGFloat] values including alpha.
	Components() unsafe.Pointer
	// Returns the color components of the color including alpha.
	NumberOfComponents() uintptr
	// Returns the unpremultiplied red component of the color.
	Red() float64
	// Returns the unpremultiplied green component of the color.
	Green() float64
	// Returns the unpremultiplied blue component of the color.
	Blue() float64
	// Returns the alpha value of the color.
	Alpha() float64
	// Returns a formatted string with the unpremultiplied color and alpha components of the color.
	StringRepresentation() string

	// Initialize a Core Image color object in the sRGB color space with the specified red, green, and blue component values.
	InitWithRedGreenBlue(red float64, green float64, blue float64) CIColor
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (c CIColor) Init() CIColor {
	rv := objc.Send[CIColor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c CIColor) Autorelease() CIColor {
	rv := objc.Send[CIColor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewCIColor creates a new CIColor instance.
func NewCIColor() CIColor {
	class := getCIColorClass()
	rv := objc.Send[CIColor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Create a Core Image color object with a Core Graphics color object.
//
// # Return Value
// 
// An initialized [CIColor] instance.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/init(cgColor:)
func NewColorWithCGColor(color coregraphics.CGColorRef) CIColor {
	instance := getCIColorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCGColor:"), color)
	return CIColorFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/init(color:)
func NewColorWithColor(color objectivec.IObject) CIColor {
	instance := getCIColorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithColor:"), color)
	return CIColorFromID(rv)
}

// Initialize a Core Image color object in the sRGB color space with the
// specified red, green, and blue component values.
//
// red: The color’s unpremultiplied red component value between 0 and 1.
//
// green: The color’s unpremultiplied green component value between 0 and 1.
//
// blue: The color’s unpremultiplied blue component value between 0 and 1.
//
// # Return Value
// 
// An initialized [CIColor] instance.
//
// # Discussion
// 
// On macOS before 10.10, the CIColor’s color space will be Generic RGB.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/initWithRed:green:blue:
func NewColorWithRedGreenBlue(red float64, green float64, blue float64) CIColor {
	instance := getCIColorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRed:green:blue:"), red, green, blue)
	return CIColorFromID(rv)
}

// Initialize a Core Image color object in the sRGB color space with the
// specified red, green, blue, and alpha component values.
//
// red: The color’s unpremultiplied red component value between 0 and 1.
//
// green: The color’s unpremultiplied green component value between 0 and 1.
//
// blue: The color’s unpremultiplied blue component value between 0 and 1.
//
// alpha: The color’s alpha (opacity) value between 0 and 1.
//
// # Return Value
// 
// An initialized [CIColor] instance.
//
// # Discussion
// 
// On macOS before 10.10, the CIColor’s color space will be Generic RGB.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/init(red:green:blue:alpha:)
func NewColorWithRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) CIColor {
	instance := getCIColorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRed:green:blue:alpha:"), red, green, blue, alpha)
	return CIColorFromID(rv)
}

// Initialize a Core Image color object with the specified red, green, and
// blue component values as measured in the specified color space.
//
// red: The color’s unpremultiplied red component value.
//
// green: The color’s unpremultiplied green component value.
//
// blue: The color’s unpremultiplied blue component value.
//
// alpha: The color’s alpha (opacity) value between 0 and 1.
//
// colorSpace: The color’s [CGColorSpace] which must have `kCGColorSpaceModelRGB`.
//
// # Return Value
// 
// An initialized [CIColor] instance.
//
// # Discussion
// 
// This will return null if the [CGColorSpace] is not `kCGColorSpaceModelRGB`.
// The RGB values can be outside the `0...1` range if the [CGColorSpace] is
// unclamped.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/init(red:green:blue:alpha:colorSpace:)
func NewColorWithRedGreenBlueAlphaColorSpace(red float64, green float64, blue float64, alpha float64, colorSpace coregraphics.CGColorSpaceRef) CIColor {
	instance := getCIColorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRed:green:blue:alpha:colorSpace:"), red, green, blue, alpha, colorSpace)
	return CIColorFromID(rv)
}

// Initialize a Core Image color object with the specified red, green, and
// blue component values as measured in the specified color space.
//
// red: The color’s unpremultiplied red component value.
//
// green: The color’s unpremultiplied green component value.
//
// blue: The color’s unpremultiplied blue component value.
//
// colorSpace: The color’s [CGColorSpace] which must have `kCGColorSpaceModelRGB`.
//
// # Return Value
// 
// An initialized [CIColor] instance.
//
// # Discussion
// 
// This will return null if the [CGColorSpace] is not `kCGColorSpaceModelRGB`.
// The RGB values can be outside the `0...1` range if the [CGColorSpace] is
// unclamped.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/init(red:green:blue:colorSpace:)
func NewColorWithRedGreenBlueColorSpace(red float64, green float64, blue float64, colorSpace coregraphics.CGColorSpaceRef) CIColor {
	instance := getCIColorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRed:green:blue:colorSpace:"), red, green, blue, colorSpace)
	return CIColorFromID(rv)
}

// Create a Core Image color object in the sRGB color space using a string
// containing the RGBA color component values.
//
// representation: A string that contains color and alpha float values. For example, the
// string: `"0.5 0.7 0.3 1.0"` indicates an RGB color whose components are 50%
// red, 70% green, 30% blue, and 100% opaque. If the string contains only 3
// float values, the alpha component will be `1.0` If the string contains no
// float values, then `/CIColor/clearColor` will be returned.
//
// # Return Value
// 
// An autoreleased [CIColor] instance.
//
// # Discussion
// 
// On macOS before 10.10, the CIColor’s color space will be Generic RGB.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/init(string:)
func NewColorWithString(representation string) CIColor {
	rv := objc.Send[objc.ID](objc.ID(getCIColorClass().class), objc.Sel("colorWithString:"), objc.String(representation))
	return CIColorFromID(rv)
}

// Create a Core Image color object with a Core Graphics color object.
//
// # Return Value
// 
// An initialized [CIColor] instance.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/init(cgColor:)
func (c CIColor) InitWithCGColor(color coregraphics.CGColorRef) CIColor {
	rv := objc.Send[CIColor](c.ID, objc.Sel("initWithCGColor:"), color)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/init(color:)
func (c CIColor) InitWithColor(color objectivec.IObject) CIColor {
	rv := objc.Send[CIColor](c.ID, objc.Sel("initWithColor:"), color)
	return rv
}
// Initialize a Core Image color object in the sRGB color space with the
// specified red, green, blue, and alpha component values.
//
// red: The color’s unpremultiplied red component value between 0 and 1.
//
// green: The color’s unpremultiplied green component value between 0 and 1.
//
// blue: The color’s unpremultiplied blue component value between 0 and 1.
//
// alpha: The color’s alpha (opacity) value between 0 and 1.
//
// # Return Value
// 
// An initialized [CIColor] instance.
//
// # Discussion
// 
// On macOS before 10.10, the CIColor’s color space will be Generic RGB.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/init(red:green:blue:alpha:)
func (c CIColor) InitWithRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) CIColor {
	rv := objc.Send[CIColor](c.ID, objc.Sel("initWithRed:green:blue:alpha:"), red, green, blue, alpha)
	return rv
}
// Initialize a Core Image color object with the specified red, green, and
// blue component values as measured in the specified color space.
//
// red: The color’s unpremultiplied red component value.
//
// green: The color’s unpremultiplied green component value.
//
// blue: The color’s unpremultiplied blue component value.
//
// colorSpace: The color’s [CGColorSpace] which must have `kCGColorSpaceModelRGB`.
//
// # Return Value
// 
// An initialized [CIColor] instance.
//
// # Discussion
// 
// This will return null if the [CGColorSpace] is not `kCGColorSpaceModelRGB`.
// The RGB values can be outside the `0...1` range if the [CGColorSpace] is
// unclamped.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/init(red:green:blue:colorSpace:)
func (c CIColor) InitWithRedGreenBlueColorSpace(red float64, green float64, blue float64, colorSpace coregraphics.CGColorSpaceRef) CIColor {
	rv := objc.Send[CIColor](c.ID, objc.Sel("initWithRed:green:blue:colorSpace:"), red, green, blue, colorSpace)
	return rv
}
// Initialize a Core Image color object with the specified red, green, and
// blue component values as measured in the specified color space.
//
// red: The color’s unpremultiplied red component value.
//
// green: The color’s unpremultiplied green component value.
//
// blue: The color’s unpremultiplied blue component value.
//
// alpha: The color’s alpha (opacity) value between 0 and 1.
//
// colorSpace: The color’s [CGColorSpace] which must have `kCGColorSpaceModelRGB`.
//
// # Return Value
// 
// An initialized [CIColor] instance.
//
// # Discussion
// 
// This will return null if the [CGColorSpace] is not `kCGColorSpaceModelRGB`.
// The RGB values can be outside the `0...1` range if the [CGColorSpace] is
// unclamped.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/init(red:green:blue:alpha:colorSpace:)
func (c CIColor) InitWithRedGreenBlueAlphaColorSpace(red float64, green float64, blue float64, alpha float64, colorSpace coregraphics.CGColorSpaceRef) CIColor {
	rv := objc.Send[CIColor](c.ID, objc.Sel("initWithRed:green:blue:alpha:colorSpace:"), red, green, blue, alpha, colorSpace)
	return rv
}
// Initialize a Core Image color object in the sRGB color space with the
// specified red, green, and blue component values.
//
// red: The color’s unpremultiplied red component value between 0 and 1.
//
// green: The color’s unpremultiplied green component value between 0 and 1.
//
// blue: The color’s unpremultiplied blue component value between 0 and 1.
//
// # Return Value
// 
// An initialized [CIColor] instance.
//
// # Discussion
// 
// On macOS before 10.10, the CIColor’s color space will be Generic RGB.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/initWithRed:green:blue:
func (c CIColor) InitWithRedGreenBlue(red float64, green float64, blue float64) CIColor {
	rv := objc.Send[CIColor](c.ID, objc.Sel("initWithRed:green:blue:"), red, green, blue)
	return rv
}
func (c CIColor) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Create a Core Image color object in the sRGB color space with the specified
// red, green, and blue component values.
//
// red: The color’s unpremultiplied red component value between 0 and 1.
//
// green: The color’s unpremultiplied green component value between 0 and 1.
//
// blue: The color’s unpremultiplied blue component value between 0 and 1.
//
// # Return Value
// 
// An autoreleased [CIColor] instance.
//
// # Discussion
// 
// On macOS before 10.10, the CIColor’s color space will be Generic RGB.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/init(red:green:blue:)
func (_CIColorClass CIColorClass) ColorWithRedGreenBlue(red float64, green float64, blue float64) CIColor {
	rv := objc.Send[objc.ID](objc.ID(_CIColorClass.class), objc.Sel("colorWithRed:green:blue:"), red, green, blue)
	return CIColorFromID(rv)
}
// Create a Core Image color object with a Core Graphics color object.
//
// # Return Value
// 
// An autoreleased [CIColor] instance.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/colorWithCGColor:
func (_CIColorClass CIColorClass) ColorWithCGColor(color coregraphics.CGColorRef) CIColor {
	rv := objc.Send[objc.ID](objc.ID(_CIColorClass.class), objc.Sel("colorWithCGColor:"), color)
	return CIColorFromID(rv)
}
// Create a Core Image color object in the sRGB color space with the specified
// red, green, blue, and alpha component values.
//
// red: The color’s unpremultiplied red component value between 0 and 1.
//
// green: The color’s unpremultiplied green component value between 0 and 1.
//
// blue: The color’s unpremultiplied blue component value between 0 and 1.
//
// alpha: The color’s alpha (opacity) value between 0 and 1.
//
// # Return Value
// 
// An autoreleased [CIColor] instance.
//
// # Discussion
// 
// On macOS before 10.10, the CIColor’s color space will be Generic RGB.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/colorWithRed:green:blue:alpha:
func (_CIColorClass CIColorClass) ColorWithRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) CIColor {
	rv := objc.Send[objc.ID](objc.ID(_CIColorClass.class), objc.Sel("colorWithRed:green:blue:alpha:"), red, green, blue, alpha)
	return CIColorFromID(rv)
}
// Create a Core Image color object with the specified red, green, blue, and
// alpha component values as measured in the specified color space.
//
// red: The color’s unpremultiplied red component value.
//
// green: The color’s unpremultiplied green component value.
//
// blue: The color’s unpremultiplied blue component value.
//
// alpha: The color’s alpha (opacity) value between 0 and 1.
//
// colorSpace: The color’s [CGColorSpace] which must have `kCGColorSpaceModelRGB`.
//
// # Return Value
// 
// An autoreleased [CIColor] instance.
//
// # Discussion
// 
// This will return `null` if the [CGColorSpace] is not
// `kCGColorSpaceModelRGB`.
// 
// The RGB values can be outside the `0...1` range if the [CGColorSpace] is
// unclamped.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/colorWithRed:green:blue:alpha:colorSpace:
func (_CIColorClass CIColorClass) ColorWithRedGreenBlueAlphaColorSpace(red float64, green float64, blue float64, alpha float64, colorSpace coregraphics.CGColorSpaceRef) CIColor {
	rv := objc.Send[objc.ID](objc.ID(_CIColorClass.class), objc.Sel("colorWithRed:green:blue:alpha:colorSpace:"), red, green, blue, alpha, colorSpace)
	return CIColorFromID(rv)
}
// Create a Core Image color object with the specified red, green, and blue
// component values as measured in the specified color space.
//
// red: The color’s unpremultiplied red component value.
//
// green: The color’s unpremultiplied green component value.
//
// blue: The color’s unpremultiplied blue component value.
//
// colorSpace: The color’s [CGColorSpace] which must have `kCGColorSpaceModelRGB`.
//
// # Return Value
// 
// An autoreleased [CIColor] instance.
//
// # Discussion
// 
// This will return `null` if the [CGColorSpace] is not
// `kCGColorSpaceModelRGB`.
// 
// The RGB values can be outside the `0...1` range if the [CGColorSpace] is
// unclamped.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/colorWithRed:green:blue:colorSpace:
func (_CIColorClass CIColorClass) ColorWithRedGreenBlueColorSpace(red float64, green float64, blue float64, colorSpace coregraphics.CGColorSpaceRef) CIColor {
	rv := objc.Send[objc.ID](objc.ID(_CIColorClass.class), objc.Sel("colorWithRed:green:blue:colorSpace:"), red, green, blue, colorSpace)
	return CIColorFromID(rv)
}

// Returns the [CGColorSpace] associated with the color
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/colorSpace
func (c CIColor) ColorSpace() coregraphics.CGColorSpaceRef {
	rv := objc.Send[coregraphics.CGColorSpaceRef](c.ID, objc.Sel("colorSpace"))
	return coregraphics.CGColorSpaceRef(rv)
}
// Return a pointer to an array of [CGFloat] values including alpha.
//
// # Discussion
// 
// Typically this array will contain `4` [CGFloat] values for red, green,
// blue, and alpha. If the [CIColor] was initialized with a [CGColor] then
// returned pointer will be the same as calling `CGColorGetComponents()`
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/components
func (c CIColor) Components() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c.ID, objc.Sel("components"))
	return rv
}
// Returns the color components of the color including alpha.
//
// # Discussion
// 
// This number includes the alpha component if the color contains one.
// 
// Typically this number will be `4` for red, green, blue, and alpha. If the
// [CIColor] was initialized with a [CGColor] then the number will be the same
// as calling `CGColorGetNumberOfComponents()`
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/numberOfComponents
func (c CIColor) NumberOfComponents() uintptr {
	rv := objc.Send[uintptr](c.ID, objc.Sel("numberOfComponents"))
	return rv
}
// Returns the unpremultiplied red component of the color.
//
// # Discussion
// 
// If the [CIColor] was initialized with a [CGColor] in a non-RGB
// [CGColorSpace] then it will be converted to sRGB to get the red component.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/red-swift.property
func (c CIColor) Red() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("red"))
	return rv
}
// Returns the unpremultiplied green component of the color.
//
// # Discussion
// 
// If the [CIColor] was initialized with a [CGColor] in a non-RGB
// [CGColorSpace] then it will be converted to sRGB to get the green
// component.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/green-swift.property
func (c CIColor) Green() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("green"))
	return rv
}
// Returns the unpremultiplied blue component of the color.
//
// # Discussion
// 
// If the [CIColor] was initialized with a [CGColor] in a non-RGB
// [CGColorSpace] then it will be converted to sRGB to get the green
// component.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/blue-swift.property
func (c CIColor) Blue() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("blue"))
	return rv
}
// Returns the alpha value of the color.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/alpha
func (c CIColor) Alpha() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("alpha"))
	return rv
}
// Returns a formatted string with the unpremultiplied color and alpha
// components of the color.
//
// # Discussion
// 
// The string representation always has four components: red, green, blue, and
// alpha.
// 
// Some example string representations of colors:
// 
// [Table data omitted]
// 
// To create a [CIColor] instance from a string representation, use the
// [ColorWithString] method.
// 
// If the [CIColor] was initialized with a [CGColor] in a non-RGB
// [CGColorSpace] then it will be converted to sRGB to get the red, green, and
// blue components.
// 
// This property is not KVO-safe because it returns a new [NSString] instance
// each time. The value of the [NSString] will be the same each time it is
// called.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/stringRepresentation
func (c CIColor) StringRepresentation() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("stringRepresentation"))
	return foundation.NSStringFromID(rv).String()
}

// Returns a singleton Core Image color instance in the sRGB color space with
// RGB values `0,0,0` and alpha value `1`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/black
func (_CIColorClass CIColorClass) BlackColor() CIColor {
	rv := objc.Send[objc.ID](objc.ID(_CIColorClass.class), objc.Sel("blackColor"))
	return CIColorFromID(objc.ID(rv))
}
// Returns a singleton Core Image color instance in the sRGB color space with
// RGB values `0,0,1` and alpha value `1`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/blue-swift.type.property
func (_CIColorClass CIColorClass) BlueColor() CIColor {
	rv := objc.Send[objc.ID](objc.ID(_CIColorClass.class), objc.Sel("blueColor"))
	return CIColorFromID(objc.ID(rv))
}
// Returns a singleton Core Image color instance in the sRGB color space with
// RGB values `0,0,0` and alpha value `0`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/clear
func (_CIColorClass CIColorClass) ClearColor() CIColor {
	rv := objc.Send[objc.ID](objc.ID(_CIColorClass.class), objc.Sel("clearColor"))
	return CIColorFromID(objc.ID(rv))
}
// Returns a singleton Core Image color instance in the sRGB color space with
// RGB values `0,1,1` and alpha value `1`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/cyan
func (_CIColorClass CIColorClass) CyanColor() CIColor {
	rv := objc.Send[objc.ID](objc.ID(_CIColorClass.class), objc.Sel("cyanColor"))
	return CIColorFromID(objc.ID(rv))
}
// Returns a singleton Core Image color instance in the sRGB color space with
// RGB values `0.5,0.5,0.5` and alpha value `1`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/gray
func (_CIColorClass CIColorClass) GrayColor() CIColor {
	rv := objc.Send[objc.ID](objc.ID(_CIColorClass.class), objc.Sel("grayColor"))
	return CIColorFromID(objc.ID(rv))
}
// Returns a singleton Core Image color instance in the sRGB color space with
// RGB values `0,1,0` and alpha value `1`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/green-swift.type.property
func (_CIColorClass CIColorClass) GreenColor() CIColor {
	rv := objc.Send[objc.ID](objc.ID(_CIColorClass.class), objc.Sel("greenColor"))
	return CIColorFromID(objc.ID(rv))
}
// Returns a singleton Core Image color instance in the sRGB color space with
// RGB values `1,0,1` and alpha value `1`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/magenta
func (_CIColorClass CIColorClass) MagentaColor() CIColor {
	rv := objc.Send[objc.ID](objc.ID(_CIColorClass.class), objc.Sel("magentaColor"))
	return CIColorFromID(objc.ID(rv))
}
// Returns a singleton Core Image color instance in the sRGB color space with
// RGB values `1,0,0` and alpha value `1`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/red-swift.type.property
func (_CIColorClass CIColorClass) RedColor() CIColor {
	rv := objc.Send[objc.ID](objc.ID(_CIColorClass.class), objc.Sel("redColor"))
	return CIColorFromID(objc.ID(rv))
}
// Returns a singleton Core Image color instance in the sRGB color space with
// RGB values `1,1,1` and alpha value `1`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/white
func (_CIColorClass CIColorClass) WhiteColor() CIColor {
	rv := objc.Send[objc.ID](objc.ID(_CIColorClass.class), objc.Sel("whiteColor"))
	return CIColorFromID(objc.ID(rv))
}
// Returns a singleton Core Image color instance in the sRGB color space with
// RGB values `1,1,0` and alpha value `1`.
//
// See: https://developer.apple.com/documentation/CoreImage/CIColor/yellow
func (_CIColorClass CIColorClass) YellowColor() CIColor {
	rv := objc.Send[objc.ID](objc.ID(_CIColorClass.class), objc.Sel("yellowColor"))
	return CIColorFromID(objc.ID(rv))
}

