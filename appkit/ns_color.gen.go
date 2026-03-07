// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSColor] class.
var (
	_NSColorClass     NSColorClass
	_NSColorClassOnce sync.Once
)

func getNSColorClass() NSColorClass {
	_NSColorClassOnce.Do(func() {
		_NSColorClass = NSColorClass{class: objc.GetClass("NSColor")}
	})
	return _NSColorClass
}

// GetNSColorClass returns the class object for NSColor.
func GetNSColorClass() NSColorClass {
	return getNSColorClass()
}

type NSColorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSColorClass) Alloc() NSColor {
	rv := objc.Send[NSColor](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object that stores color data and sometimes opacity (alpha value).
//
// # Overview
// 
// Many methods in AppKit require you to specify color data using an [NSColor]
// object; when drawing you use them to set the current fill and stroke
// colors. Color objects are immutable and thread-safe. You can create color
// objects in many ways:
// 
// - Load colors from an asset catalog. Colors created from assets can adapt
// automatically to system appearance changes. - Use the semantic colors for
// custom UI elements, so that they match the appearance of other AppKit
// views; see [UI element colors]. - Use the adaptable system colors, such as
// [systemBlue], when you want a specific tint that looks correct in both
// light and dark environments. - Create a color object from another object,
// such as a Core Graphics representation of a color, or a Core Image color. -
// Create a color from an [NSImage] object, and paint a repeating pattern
// instead of using a solid color. - Create a color by applying a transform to
// another [NSColor] object. For example, you might perform a blend operation
// between two colors, or you might create a color that represents the same
// color, but in a different color space. - Create custom colors using raw
// component values, and a variety of color spaces, when you need to represent
// user-specified colors.
// 
// For user-specified colors, you can also display a color panel and let the
// user specify the color. For information about color panels, see
// [NSColorPanel].
// 
// # Color and color spaces
// 
// A color object is typically represented internally as a Core Graphics color
// ([NSColor.CGColor]) in a Core Graphics color space ([CGColorSpace]). Colors can
// also be created in extended color spaces:
// 
// - [ExtendedSRGBColorSpace]
// - [ExtendedGenericGamma22GrayColorSpace]
// 
// When you need to worry about color spaces, use extended color spaces as
// working color spaces. When you need to worry about representing that color
// as closely as possible in a specific color space, convert the color from
// the extended color space into the target color space.
// 
// When working in an extended color space, color values are not clamped to
// fit inside the color gamut, meaning that component values may be less than
// `0.0` or greater than `1.0`. When displayed on an sRGB display, such colors
// are outside the gamut and won’t render accurately. However, extended
// color spaces are useful as working color spaces when you want a pixel
// format and representation that other color spaces can be easily converted
// into. For example, a color in the Display P3 color space can convert to an
// extended sRGB format, even if it isn’t within the sRGB color gamut. While
// some of the converted color’s values are outside of the 0-1.0 range, the
// color renders correctly when viewed on a device with a P3 display gamut.
// 
// It is a programmer error to access color components of a color space that
// the [NSColor] object does not support. For example, you cannot access the
// [NSColor.RedComponent] property and [NSColor.GetRedGreenBlueAlpha] method on a color that
// uses the CMYK color space. Further, the [NSColor.GetComponents] method and
// [NSColor.NumberOfComponents] property work only in color spaces that have
// individual components. As such, they return the components of color objects
// as individual floating-point values regardless of whether they’re based
// on [NSColorSpace] objects or named color spaces. However, older
// component-fetching methods such as [NSColor.GetRedGreenBlueAlpha] are effective
// only on color objects based on named color spaces.
// 
// If you have a color object in an unknown color space and you want to
// extract its components, convert the color object to a known color space and
// then use the component accessor methods of that color space.
// 
// For design guidance, see Human Interface Guidelines > [Color].
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
// [NSColor.CGColor]: https://developer.apple.com/documentation/CoreGraphics/CGColor
// [Color]: https://developer.apple.com/design/human-interface-guidelines/color/
// [UI element colors]: https://developer.apple.com/documentation/AppKit/ui-element-colors
// [systemBlue]: https://developer.apple.com/documentation/AppKit/NSColor/systemBlue
//
// # Applying specific appearances to colors
//
//   - [NSColor.ColorWithSystemEffect]: Returns a new color object that represents the current color modified to include the specified visual effect.
//
// # Transforming existing color objects
//
//   - [NSColor.ColorUsingColorSpace]: Creates a new color object representing the color of the current color object in the specified color space.
//   - [NSColor.BlendedColorWithFractionOfColor]: Creates a new color object whose component values are a weighted sum of the current color object and the specified color object’s.
//   - [NSColor.ColorWithAlphaComponent]: Creates a new color object that has the same color space and component values as the current color object, but the specified alpha component.
//   - [NSColor.HighlightWithLevel]: Creates a new color object that represents a blend between the current color and the highlight color.
//   - [NSColor.ShadowWithLevel]: Creates a new color object that represents a blend between the current color and the shadow color.
//
// # Copying and pasting color Information
//
//   - [NSColor.WriteToPasteboard]: Writes the color object’s data to the specified pasteboard.
//
// # Retrieving component values from color objects
//
//   - [NSColor.GetCyanMagentaYellowBlackAlpha]: Returns the color object’s CMYK and opacity values.
//   - [NSColor.GetHueSaturationBrightnessAlpha]: Returns the color object’s HSB component and opacity values in the respective arguments.
//   - [NSColor.GetRedGreenBlueAlpha]: Returns the color object’s RGB component and opacity values in the respective arguments.
//   - [NSColor.GetWhiteAlpha]: Returns the grayscale and alpha values of the color.
//   - [NSColor.NumberOfComponents]: The number of components in the color.
//   - [NSColor.GetComponents]: Returns the components of the color as an array.
//
// # Retrieving individual components
//
//   - [NSColor.AlphaComponent]: The alpha (opacity) component value of the color.
//   - [NSColor.WhiteComponent]: The white component value of the color.
//   - [NSColor.RedComponent]: The red component value of the color.
//   - [NSColor.GreenComponent]: The green component value of the color.
//   - [NSColor.BlueComponent]: The blue component value of the color.
//   - [NSColor.CyanComponent]: The cyan component value of the color.
//   - [NSColor.MagentaComponent]: The magenta component value of the color.
//   - [NSColor.YellowComponent]: The yellow component value of the color.
//   - [NSColor.BlackComponent]: The black component value of the color.
//   - [NSColor.HueComponent]: The hue component value of the color.
//   - [NSColor.SaturationComponent]: The saturation component value of the color.
//   - [NSColor.BrightnessComponent]: The brightness component value of the color.
//   - [NSColor.CatalogNameComponent]: The catalog containing the color’s name.
//   - [NSColor.LocalizedCatalogNameComponent]: The localized version of the catalog name containing the color.
//   - [NSColor.ColorNameComponent]: The name of the color.
//   - [NSColor.LocalizedColorNameComponent]: The localized version of the color name.
//
// # Working with the color space
//
//   - [NSColor.Type]: The type of the color object.
//   - [NSColor.ColorUsingType]: Returns a version of the color object that is compatible with the specified color type.
//   - [NSColor.ColorSpace]: The color space associated with the color.
//
// # Supporting high dynamic range (HDR) colors
//
//   - [NSColor.LinearExposure]: For HDR colors, the linear brightness multiplier that was applied when generating the color. Colors created with an exposure by NSColor create CGColors that are tagged with a contentHeadroom value. While CGColors created without a contentHeadroom tag will return 0 from CGColorGetHeadroom, NSColors generated in a similar fashion return a linearExposure of 1.0.
//   - [NSColor.StandardDynamicRangeColor]: In some cases it is useful to recover the color that was base the SDR color that was exposed to generate an HDR color. If a color’s `linearExposure` is > 1, then this will return the base SDR color. If the color is not an HDR color, this will return `self`.
//   - [NSColor.ColorByApplyingContentHeadroom]: Reinterpret the color by applying a new `contentHeadroom` without changing the color components. Changing the `contentHeadroom` redefines the color relative to a different peak white, changing its behavior under tone mapping and the result of calling `standardDynamicRangeColor`. The new color will have a `contentHeadroom` >= 1.0. If called on a color with a color space that does not support extended range, or does not have an equivalent extended range counterpart, this will return `self`.
//
// # Retrieving core graphics color information
//
//   - [NSColor.CGColor]: The Core Graphics color object corresponding to the color.
//
// # Drawing with colors
//
//   - [NSColor.DrawSwatchInRect]: Draws the current color in the specified rectangle.
//   - [NSColor.Set]: Sets the color of subsequent drawing to the color that the color object represents.
//   - [NSColor.SetFill]: Sets the fill color of subsequent drawing to the color object’s color.
//   - [NSColor.SetStroke]: Sets the stroke color of subsequent drawing to the color object’s color.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor
type NSColor struct {
	objectivec.Object
}

// NSColorFromID constructs a [NSColor] from an objc.ID.
//
// An object that stores color data and sometimes opacity (alpha value).
func NSColorFromID(id objc.ID) NSColor {
	return NSColor{objectivec.Object{id}}
}
// NOTE: NSColor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSColor] class.
//
// # Applying specific appearances to colors
//
//   - [INSColor.ColorWithSystemEffect]: Returns a new color object that represents the current color modified to include the specified visual effect.
//
// # Transforming existing color objects
//
//   - [INSColor.ColorUsingColorSpace]: Creates a new color object representing the color of the current color object in the specified color space.
//   - [INSColor.BlendedColorWithFractionOfColor]: Creates a new color object whose component values are a weighted sum of the current color object and the specified color object’s.
//   - [INSColor.ColorWithAlphaComponent]: Creates a new color object that has the same color space and component values as the current color object, but the specified alpha component.
//   - [INSColor.HighlightWithLevel]: Creates a new color object that represents a blend between the current color and the highlight color.
//   - [INSColor.ShadowWithLevel]: Creates a new color object that represents a blend between the current color and the shadow color.
//
// # Copying and pasting color Information
//
//   - [INSColor.WriteToPasteboard]: Writes the color object’s data to the specified pasteboard.
//
// # Retrieving component values from color objects
//
//   - [INSColor.GetCyanMagentaYellowBlackAlpha]: Returns the color object’s CMYK and opacity values.
//   - [INSColor.GetHueSaturationBrightnessAlpha]: Returns the color object’s HSB component and opacity values in the respective arguments.
//   - [INSColor.GetRedGreenBlueAlpha]: Returns the color object’s RGB component and opacity values in the respective arguments.
//   - [INSColor.GetWhiteAlpha]: Returns the grayscale and alpha values of the color.
//   - [INSColor.NumberOfComponents]: The number of components in the color.
//   - [INSColor.GetComponents]: Returns the components of the color as an array.
//
// # Retrieving individual components
//
//   - [INSColor.AlphaComponent]: The alpha (opacity) component value of the color.
//   - [INSColor.WhiteComponent]: The white component value of the color.
//   - [INSColor.RedComponent]: The red component value of the color.
//   - [INSColor.GreenComponent]: The green component value of the color.
//   - [INSColor.BlueComponent]: The blue component value of the color.
//   - [INSColor.CyanComponent]: The cyan component value of the color.
//   - [INSColor.MagentaComponent]: The magenta component value of the color.
//   - [INSColor.YellowComponent]: The yellow component value of the color.
//   - [INSColor.BlackComponent]: The black component value of the color.
//   - [INSColor.HueComponent]: The hue component value of the color.
//   - [INSColor.SaturationComponent]: The saturation component value of the color.
//   - [INSColor.BrightnessComponent]: The brightness component value of the color.
//   - [INSColor.CatalogNameComponent]: The catalog containing the color’s name.
//   - [INSColor.LocalizedCatalogNameComponent]: The localized version of the catalog name containing the color.
//   - [INSColor.ColorNameComponent]: The name of the color.
//   - [INSColor.LocalizedColorNameComponent]: The localized version of the color name.
//
// # Working with the color space
//
//   - [INSColor.Type]: The type of the color object.
//   - [INSColor.ColorUsingType]: Returns a version of the color object that is compatible with the specified color type.
//   - [INSColor.ColorSpace]: The color space associated with the color.
//
// # Supporting high dynamic range (HDR) colors
//
//   - [INSColor.LinearExposure]: For HDR colors, the linear brightness multiplier that was applied when generating the color. Colors created with an exposure by NSColor create CGColors that are tagged with a contentHeadroom value. While CGColors created without a contentHeadroom tag will return 0 from CGColorGetHeadroom, NSColors generated in a similar fashion return a linearExposure of 1.0.
//   - [INSColor.StandardDynamicRangeColor]: In some cases it is useful to recover the color that was base the SDR color that was exposed to generate an HDR color. If a color’s `linearExposure` is > 1, then this will return the base SDR color. If the color is not an HDR color, this will return `self`.
//   - [INSColor.ColorByApplyingContentHeadroom]: Reinterpret the color by applying a new `contentHeadroom` without changing the color components. Changing the `contentHeadroom` redefines the color relative to a different peak white, changing its behavior under tone mapping and the result of calling `standardDynamicRangeColor`. The new color will have a `contentHeadroom` >= 1.0. If called on a color with a color space that does not support extended range, or does not have an equivalent extended range counterpart, this will return `self`.
//
// # Retrieving core graphics color information
//
//   - [INSColor.CGColor]: The Core Graphics color object corresponding to the color.
//
// # Drawing with colors
//
//   - [INSColor.DrawSwatchInRect]: Draws the current color in the specified rectangle.
//   - [INSColor.Set]: Sets the color of subsequent drawing to the color that the color object represents.
//   - [INSColor.SetFill]: Sets the fill color of subsequent drawing to the color object’s color.
//   - [INSColor.SetStroke]: Sets the stroke color of subsequent drawing to the color object’s color.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor
type INSColor interface {
	objectivec.IObject
	NSAccessibilityColor
	NSPasteboardWriting

	// Topic: Applying specific appearances to colors

	// Returns a new color object that represents the current color modified to include the specified visual effect.
	ColorWithSystemEffect(systemEffect NSColorSystemEffect) INSColor

	// Topic: Transforming existing color objects

	// Creates a new color object representing the color of the current color object in the specified color space.
	ColorUsingColorSpace(space INSColorSpace) INSColor
	// Creates a new color object whose component values are a weighted sum of the current color object and the specified color object’s.
	BlendedColorWithFractionOfColor(fraction float64, color INSColor) INSColor
	// Creates a new color object that has the same color space and component values as the current color object, but the specified alpha component.
	ColorWithAlphaComponent(alpha float64) INSColor
	// Creates a new color object that represents a blend between the current color and the highlight color.
	HighlightWithLevel(val float64) INSColor
	// Creates a new color object that represents a blend between the current color and the shadow color.
	ShadowWithLevel(val float64) INSColor

	// Topic: Copying and pasting color Information

	// Writes the color object’s data to the specified pasteboard.
	WriteToPasteboard(pasteBoard INSPasteboard)

	// Topic: Retrieving component values from color objects

	// Returns the color object’s CMYK and opacity values.
	GetCyanMagentaYellowBlackAlpha(cyan float64, magenta float64, yellow float64, black float64, alpha float64)
	// Returns the color object’s HSB component and opacity values in the respective arguments.
	GetHueSaturationBrightnessAlpha(hue float64, saturation float64, brightness float64, alpha float64)
	// Returns the color object’s RGB component and opacity values in the respective arguments.
	GetRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64)
	// Returns the grayscale and alpha values of the color.
	GetWhiteAlpha(white float64, alpha float64)
	// The number of components in the color.
	NumberOfComponents() int
	// Returns the components of the color as an array.
	GetComponents(components float64)

	// Topic: Retrieving individual components

	// The alpha (opacity) component value of the color.
	AlphaComponent() float64
	// The white component value of the color.
	WhiteComponent() float64
	// The red component value of the color.
	RedComponent() float64
	// The green component value of the color.
	GreenComponent() float64
	// The blue component value of the color.
	BlueComponent() float64
	// The cyan component value of the color.
	CyanComponent() float64
	// The magenta component value of the color.
	MagentaComponent() float64
	// The yellow component value of the color.
	YellowComponent() float64
	// The black component value of the color.
	BlackComponent() float64
	// The hue component value of the color.
	HueComponent() float64
	// The saturation component value of the color.
	SaturationComponent() float64
	// The brightness component value of the color.
	BrightnessComponent() float64
	// The catalog containing the color’s name.
	CatalogNameComponent() NSColorListName
	// The localized version of the catalog name containing the color.
	LocalizedCatalogNameComponent() string
	// The name of the color.
	ColorNameComponent() string
	// The localized version of the color name.
	LocalizedColorNameComponent() string

	// Topic: Working with the color space

	// The type of the color object.
	Type() NSColorType
	// Returns a version of the color object that is compatible with the specified color type.
	ColorUsingType(type_ NSColorType) INSColor
	// The color space associated with the color.
	ColorSpace() INSColorSpace

	// Topic: Supporting high dynamic range (HDR) colors

	// For HDR colors, the linear brightness multiplier that was applied when generating the color. Colors created with an exposure by NSColor create CGColors that are tagged with a contentHeadroom value. While CGColors created without a contentHeadroom tag will return 0 from CGColorGetHeadroom, NSColors generated in a similar fashion return a linearExposure of 1.0.
	LinearExposure() float64
	// In some cases it is useful to recover the color that was base the SDR color that was exposed to generate an HDR color. If a color’s `linearExposure` is > 1, then this will return the base SDR color. If the color is not an HDR color, this will return `self`.
	StandardDynamicRangeColor() INSColor
	// Reinterpret the color by applying a new `contentHeadroom` without changing the color components. Changing the `contentHeadroom` redefines the color relative to a different peak white, changing its behavior under tone mapping and the result of calling `standardDynamicRangeColor`. The new color will have a `contentHeadroom` >= 1.0. If called on a color with a color space that does not support extended range, or does not have an equivalent extended range counterpart, this will return `self`.
	ColorByApplyingContentHeadroom(contentHeadroom float64) INSColor

	// Topic: Retrieving core graphics color information

	// The Core Graphics color object corresponding to the color.
	CGColor() coregraphics.CGColorRef

	// Topic: Drawing with colors

	// Draws the current color in the specified rectangle.
	DrawSwatchInRect(rect corefoundation.CGRect)
	// Sets the color of subsequent drawing to the color that the color object represents.
	Set()
	// Sets the fill color of subsequent drawing to the color object’s color.
	SetFill()
	// Sets the stroke color of subsequent drawing to the color object’s color.
	SetStroke()

	// The pattern image used to paint the target area.
	PatternImage() INSImage
	SetPatternImage(value INSImage)
	// Creates a color object from data in an unarchiver.
	InitWithCoder(coder foundation.INSCoder) NSColor
	// Initializes an instance with a property list object and a type string.
	InitWithPasteboardPropertyListOfType(propertyList objectivec.IObject, type_ NSPasteboardType) NSColor
	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (c NSColor) Init() NSColor {
	rv := objc.Send[NSColor](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSColor) Autorelease() NSColor {
	rv := objc.Send[NSColor](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSColor creates a new NSColor instance.
func NewNSColor() NSColor {
	class := getNSColorClass()
	rv := objc.Send[NSColor](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a color object from color data currently on the pasteboard.
//
// pasteBoard: The pasteboard from which to return the color.
//
// # Return Value
// 
// The color currently on the pasteboard or `nil` if `pasteBoard` doesn’t
// contain color data. The returned color’s alpha component is set to 1.0 if
// [IgnoresAlpha] returns [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(from:)
func NewColorFromPasteboard(pasteBoard INSPasteboard) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorFromPasteboard:"), pasteBoard)
	return NSColorFromID(rv)
}


// Creates a color object from the provided name, which corresponds to a color
// in the default asset catalog of the app’s main bundle.
//
// name: The name of the color in the asset catalog.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(named:)
func NewColorNamed(name string) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorNamed:"), objc.String(name))
	return NSColorFromID(rv)
}


// Creates a color object from the provided name, which corresponds to a color
// in the default asset catalog of the specified bundle.
//
// name: The name of the color to find.
//
// bundle: The app bundle.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(named:bundle:)
func NewColorNamedBundle(name string, bundle *foundation.NSBundle) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorNamed:bundle:"), objc.String(name), bundle)
	return NSColorFromID(rv)
}


// Creates a color object using the specified Core Graphics color.
//
// cgColor: The Core Graphics color reference.
//
// # Return Value
// 
// An [NSColor] instance.
//
// # Discussion
// 
// This method may return `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(cgColor:)
func NewColorWithCGColor(cgColor coregraphics.CGColorRef) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithCGColor:"), cgColor)
	return NSColorFromID(rv)
}


// Creates a color object from the specified Core Image color.
//
// color: The Core Image color to convert.
//
// # Return Value
// 
// The [NSColor] object corresponding to the specified Core Image color.
//
// # Discussion
// 
// The method raises if the color space and components associated with `color`
// are `nil` or invalid.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(CIColor:)
// color is a [coreimage.CIColor].
func NewColorWithCIColor(color objectivec.IObject) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithCIColor:"), color)
	return NSColorFromID(rv)
}


// Creates a color object using the given opacity and HSB color space
// components.
//
// hue: The hue component of the color object in the HSB color space.
//
// saturation: The saturation component of the color object in the HSB color space.
//
// brightness: The brightness (or value) component of the color object in the HSB color
// space.
//
// alpha: The opacity value of the color object,
//
// # Return Value
// 
// The color object.
//
// # Discussion
// 
// Values below 0.0 are interpreted as 0.0, and values above 1.0 are
// interpreted as 1.0.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(calibratedHue:saturation:brightness:alpha:)
func NewColorWithCalibratedHueSaturationBrightnessAlpha(hue float64, saturation float64, brightness float64, alpha float64) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithCalibratedHue:saturation:brightness:alpha:"), hue, saturation, brightness, alpha)
	return NSColorFromID(rv)
}


// Creates a color object using the given opacity and RGB components.
//
// red: The red component of the color object.
//
// green: The green component of the color object.
//
// blue: The blue component of the color object.
//
// alpha: The opacity value of the color object.
//
// # Return Value
// 
// The color object.
//
// # Discussion
// 
// Values below 0.0 are interpreted as 0.0, and values above 1.0 are
// interpreted as 1.0.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(calibratedRed:green:blue:alpha:)
func NewColorWithCalibratedRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithCalibratedRed:green:blue:alpha:"), red, green, blue, alpha)
	return NSColorFromID(rv)
}


// Creates a color object using the given opacity and grayscale values.
//
// white: The grayscale value of the color object.
//
// alpha: The opacity value of the color object.
//
// # Return Value
// 
// The color object.
//
// # Discussion
// 
// Values below 0.0 are interpreted as 0.0, and values above 1.0 are
// interpreted as 1.0.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(calibratedWhite:alpha:)
func NewColorWithCalibratedWhiteAlpha(white float64, alpha float64) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithCalibratedWhite:alpha:"), white, alpha)
	return NSColorFromID(rv)
}


// Creates a color object using the specified asset catalog and color names.
//
// listName: The name of the asset catalog in which to find the specified color; this
// may be a standard color catalog.
//
// colorName: The name of the color. Note that the color must be defined in the named
// color space to retrieve it with this method.
//
// # Return Value
// 
// The color object.
//
// # Discussion
// 
// This method searches the app’s main bundle for an asset catalog with the
// specified name. It then creates a color object based on the asset whose
// name matches the value in `colorName`.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(catalogName:colorName:)
func NewColorWithCatalogNameColorName(listName string, colorName string) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithCatalogName:colorName:"), objc.String(listName), objc.String(colorName))
	return NSColorFromID(rv)
}


// Creates a color object from data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(coder:)
func NewColorWithCoder(coder foundation.INSCoder) NSColor {
	instance := getNSColorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSColorFromID(rv)
}


// Creates a color object from the specified components of the given color
// space.
//
// space: An [NSColorSpace] object representing a color space. The colorspace should
// be component-based. The method raises if this is `nil` or a color space
// that cannot be used with [NSColor] objects.
//
// components: An array of the components in the specified color space to use to create
// the [NSColor] object. The order of these components is determined by the
// color-space profile, with the alpha component always last. (If you want the
// created color to be opaque, specify 1.0 for the alpha component.)
//
// numberOfComponents: The number of components in the `components` array. This should match the
// number dictated by the specified color space plus one for alpha. This
// method raises an exception if they do not match.
//
// # Return Value
// 
// The color object.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(colorSpace:components:count:)
func NewColorWithColorSpaceComponentsCount(space INSColorSpace, components []float64, numberOfComponents int) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithColorSpace:components:count:"), space, objc.CArray(components), numberOfComponents)
	return NSColorFromID(rv)
}


// Creates a color object with the specified color space, hue, saturation,
// brightness, and alpha channel values.
//
// space: An [NSColorSpace] object representing a color space. An exception is raised
// if the color model of the provided color space is not RGB.
//
// hue: The hue (color) component, expressed as a floating-point value in the range
// 0–1.0.
//
// saturation: The color saturation component, expressed as a floating-point value in the
// range 0–1.0.
//
// brightness: The brightness component, expressed as a floating-point value in the range
// 0–1.0.
//
// alpha: The alpha (opacity), expressed as a floating-point value in the range 0
// (transparent) to 1.0 (opaque).
//
// # Return Value
// 
// The color object.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(colorSpace:hue:saturation:brightness:alpha:)
func NewColorWithColorSpaceHueSaturationBrightnessAlpha(space INSColorSpace, hue float64, saturation float64, brightness float64, alpha float64) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithColorSpace:hue:saturation:brightness:alpha:"), space, hue, saturation, brightness, alpha)
	return NSColorFromID(rv)
}


// Creates a color object using the given opacity value and CMYK components.
//
// cyan: The cyan component of the color object.
//
// magenta: The magenta component of the color object.
//
// yellow: The yellow component of the color object.
//
// black: The black component of the color object.
//
// alpha: The opacity value of the color object.
//
// # Return Value
// 
// The color object.
//
// # Discussion
// 
// Values below 0.0 are interpreted as 0.0, and values above 1.0 are
// interpreted as 1.0.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(deviceCyan:magenta:yellow:black:alpha:)
func NewColorWithDeviceCyanMagentaYellowBlackAlpha(cyan float64, magenta float64, yellow float64, black float64, alpha float64) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithDeviceCyan:magenta:yellow:black:alpha:"), cyan, magenta, yellow, black, alpha)
	return NSColorFromID(rv)
}


// Creates a color object using the given opacity value and HSB color space
// components.
//
// hue: The hue component of the color object.
//
// saturation: The saturation component of the color object.
//
// brightness: The brightness component of the color object.
//
// alpha: The opacity value of the color object.
//
// # Return Value
// 
// The color object.
//
// # Discussion
// 
// Values below 0.0 are interpreted as 0.0, and values above 1.0 are
// interpreted as 1.0.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(deviceHue:saturation:brightness:alpha:)
func NewColorWithDeviceHueSaturationBrightnessAlpha(hue float64, saturation float64, brightness float64, alpha float64) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithDeviceHue:saturation:brightness:alpha:"), hue, saturation, brightness, alpha)
	return NSColorFromID(rv)
}


// Creates a color object using the given opacity value and RGB components.
//
// red: The red component of the color object.
//
// green: The green component of the color object.
//
// blue: The blue component of the color object.
//
// alpha: The opacity value of the color object.
//
// # Return Value
// 
// The color object.
//
// # Discussion
// 
// Values below 0.0 are interpreted as 0.0, and values above 1.0 are
// interpreted as 1.0.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(deviceRed:green:blue:alpha:)
func NewColorWithDeviceRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithDeviceRed:green:blue:alpha:"), red, green, blue, alpha)
	return NSColorFromID(rv)
}


// Creates a color object using the given opacity and grayscale values.
//
// white: The grayscale value of the color object.
//
// alpha: The opacity value of the color object.
//
// # Return Value
// 
// The color object.
//
// # Discussion
// 
// Values below 0.0 are interpreted as 0.0, and values above 1.0 are
// interpreted as 1.0.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(deviceWhite:alpha:)
func NewColorWithDeviceWhiteAlpha(white float64, alpha float64) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithDeviceWhite:alpha:"), white, alpha)
	return NSColorFromID(rv)
}


// Creates a color object from the specified components in the Display P3
// color space.
//
// red: The red component of the color object.
//
// green: The green component of the color object.
//
// blue: The blue component of the color object.
//
// alpha: The opacity value of the color object.
//
// # Return Value
// 
// The color object.
//
// # Discussion
// 
// Values below 0.0 are interpreted as 0.0, and values above 1.0 are
// interpreted as 1.0.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(displayP3Red:green:blue:alpha:)
func NewColorWithDisplayP3RedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithDisplayP3Red:green:blue:alpha:"), red, green, blue, alpha)
	return NSColorFromID(rv)
}


// Returns a color object with the specified white and alpha values in the
// GenericGamma22 colorspace.
//
// white: The white value of the color object.
//
// alpha: The opacity value of the color object.
//
// # Return Value
// 
// The color object.
//
// # Discussion
// 
// Values below 0.0 are interpreted as 0.0, and values above 1.0 are
// interpreted as 1.0.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(genericGamma22White:alpha:)
func NewColorWithGenericGamma22WhiteAlpha(white float64, alpha float64) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithGenericGamma22White:alpha:"), white, alpha)
	return NSColorFromID(rv)
}


// Creates a color object with the specified hue, saturation, brightness, and
// alpha channel values.
//
// hue: The hue (color) component. If the value is outside of the range `0–1.0`,
// the extended range color space is used.
//
// saturation: The color saturation component. If the value is outside of the range
// `0–1.0`, the extended range color space is used.
//
// brightness: The brightness component. If the value is outside of the range `0–1.0`,
// the extended range color space is used.
//
// alpha: The alpha (opacity), specified as a value from `0-1.0`. Alpha values below
// `0` are interpreted as `0.0`, and values above `1.0` are interpreted as
// `1.0`.
//
// # Return Value
// 
// The color object.
//
// # Discussion
// 
// This method accepts extended color component values. If the component
// values are outside of the `0-1.0` range, the method creates a color in the
// extended range color space. This method is provided for easier reuse of
// code that uses [UIColor] in iOS.
// 
// Where possible, it is preferable to specify the colorspace explicitly using
// the [ColorWithSRGBRedGreenBlueAlpha] or [ColorWithGenericGamma22WhiteAlpha]
// method.
//
// [UIColor]: https://developer.apple.com/documentation/UIKit/UIColor
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(hue:saturation:brightness:alpha:)
func NewColorWithHueSaturationBrightnessAlpha(hue float64, saturation float64, brightness float64, alpha float64) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithHue:saturation:brightness:alpha:"), hue, saturation, brightness, alpha)
	return NSColorFromID(rv)
}


// Initializes an instance with a property list object and a type string.
//
// propertyList: A property list containing data to initialize the receiver.
// 
// By default, the property list object is an instance of [NSData]. If you
// implement [ReadingOptionsForTypePasteboard] and specify an option other
// than [PasteboardReadingAsData], the `propertyList` may be any other
// property list object.
//
// type: A UTI supported by the receiver for reading (one of the types returned by
// [ReadableTypesForPasteboard]).
//
// # Return Value
// 
// An object initialized using the data in `propertyList`.
//
// # Discussion
// 
// This method is considered optional because, if [ReadableTypesForPasteboard]
// returns just a single type, and that type uses the
// [PasteboardReadingAsKeyedArchive] reading option, then instances are
// initialized using [init(coder:)] instead of this method.
//
// [init(coder:)]: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/init(pasteboardPropertyList:ofType:)
func NewColorWithPasteboardPropertyListOfType(propertyList objectivec.IObject, type_ NSPasteboardType) NSColor {
	instance := getNSColorClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPasteboardPropertyList:ofType:"), propertyList, objc.String(string(type_)))
	return NSColorFromID(rv)
}


// Creates a color object that uses the specified image pattern to paint the
// target area.
//
// image: The image to use as the pattern for the color object. The image is tiled
// starting at the bottom of the window. The image is not scaled.
//
// # Return Value
// 
// The [NSColor] object. This color object is autoreleased.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(patternImage:)
func NewColorWithPatternImage(image INSImage) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithPatternImage:"), image)
	return NSColorFromID(rv)
}


// Creates a color object with the specified red, green, blue, and alpha
// channel values.
//
// red: The red channel value. If the value is outside of the range `0–1.0`, the
// extended sRGB color space is used.
//
// green: The green channel value. If the value is outside of the range `0–1.0`,
// the extended sRGB color space is used.
//
// blue: The blue channel value. If the value is outside of the range `0–1.0`, the
// extended sRGB color space is used.
//
// alpha: The alpha (opacity), specified as a value from `0-1.0`. Alpha values below
// `0` are interpreted as `0.0`, and values above `1.0` are interpreted as
// `1.0`.
//
// # Discussion
// 
// This method accepts extended color component values. If the red, green,
// blue, or alpha values are outside of the `0-1.0` range, the method creates
// a color in the extended range color space. This method is provided for
// easier reuse of code that uses [UIColor] in iOS.
// 
// Where possible, it is preferable to specify the colorspace explicitly using
// the [ColorWithSRGBRedGreenBlueAlpha] or [ColorWithGenericGamma22WhiteAlpha]
// method.
//
// [UIColor]: https://developer.apple.com/documentation/UIKit/UIColor
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(red:green:blue:alpha:)
func NewColorWithRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithRed:green:blue:alpha:"), red, green, blue, alpha)
	return NSColorFromID(rv)
}


// Generates an HDR color in the extended sRGB colorspace by applying an
// exposure to the SDR color defined by the red, green, and blue components.
// The `red`, `green`, and `blue` components have a nominal range of [0..1],
// `exposure` is a value >= 0. To produce an HDR color, we process the given
// color in a linear color space, multiplying component values by
// `2^exposure`. The produced color will have a `contentHeadroom` equal to the
// linearized exposure value. Each whole value of exposure produces a color
// that is twice as bright.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(red:green:blue:alpha:exposure:)
func NewColorWithRedGreenBlueAlphaExposure(red float64, green float64, blue float64, alpha float64, exposure float64) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithRed:green:blue:alpha:exposure:"), red, green, blue, alpha, exposure)
	return NSColorFromID(rv)
}


// Generates an HDR color in the extended sRGB colorspace by applying an
// exposure to the SDR color defined by the red, green, and blue components.
// The `red`, `green`, and `blue` components have a nominal range of [0..1],
// `linearExposure` is a value >= 1. To produce an HDR color, we process the
// given color in a linear color space, multiplying component values by
// `linearExposure `. The produced color will have a `contentHeadroom` equal
// to `linearExposure`. Each doubling of `linearExposure` produces a color
// that is twice as bright.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(red:green:blue:alpha:linearExposure:)
func NewColorWithRedGreenBlueAlphaLinearExposure(red float64, green float64, blue float64, alpha float64, linearExposure float64) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithRed:green:blue:alpha:linearExposure:"), red, green, blue, alpha, linearExposure)
	return NSColorFromID(rv)
}


// Creates a color object from the specified components in the sRGB
// colorspace.
//
// red: The red component of the color object.
//
// green: The green component of the color object.
//
// blue: The blue component of the color object.
//
// alpha: The opacity value of the color object.
//
// # Return Value
// 
// The color object.
//
// # Discussion
// 
// Values below 0.0 are interpreted as 0.0, and values above 1.0 are
// interpreted as 1.0.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(srgbRed:green:blue:alpha:)
func NewColorWithSRGBRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithSRGBRed:green:blue:alpha:"), red, green, blue, alpha)
	return NSColorFromID(rv)
}


// Creates a color object with the specified brightness and alpha channel
// values.
//
// white: The brightness. If the value is outside of the range `0–1.0`, the
// extended sRGB color space is used.
//
// alpha: The alpha (opacity), expressed as a floating-point value in the range `0`
// (transparent) to `1.0` (opaque). Alpha values below `0` are interpreted as
// `0.0`, and values above `1.0` are interpreted as `1.0`.
//
// # Discussion
// 
// This method accepts extended color component values. If the alpha or white
// values are outside of the `0-1.0` range, the method creates a color in the
// extended range or [ExtendedGenericGamma22GrayColorSpace] color space that
// is compatible with the sRGB colorspace. This method is provided for easier
// reuse of code that uses [UIColor] in iOS.
// 
// Where possible, it is preferable to specify the colorspace explicitly using
// the [ColorWithSRGBRedGreenBlueAlpha] or [ColorWithGenericGamma22WhiteAlpha]
// method.
//
// [UIColor]: https://developer.apple.com/documentation/UIKit/UIColor
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(white:alpha:)
func NewColorWithWhiteAlpha(white float64, alpha float64) NSColor {
	rv := objc.Send[objc.ID](objc.ID(getNSColorClass().class), objc.Sel("colorWithWhite:alpha:"), white, alpha)
	return NSColorFromID(rv)
}







// Returns a new color object that represents the current color modified to
// include the specified visual effect.
//
// systemEffect: The visual effect you want to apply to a view or control.
//
// # Discussion
// 
// Instead of defining separate colors for user interactions with a view, use
// this method to retrieve the appropriate color for use with those
// interactions. This method blends the current color with an appropriate
// modifier and returns the results. For example, specifying
// [NSColor.SystemEffect.pressed] for the `systemEffect` parameter yields the
// color to use when you want your view to appear as if it had been pressed.
// This method takes the current appearance into account, returning an
// appropriately modified color for both light and dark appearances.
//
// [NSColor.SystemEffect.pressed]: https://developer.apple.com/documentation/AppKit/NSColor/SystemEffect/pressed
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/withSystemEffect(_:)
func (c NSColor) ColorWithSystemEffect(systemEffect NSColorSystemEffect) INSColor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("colorWithSystemEffect:"), systemEffect)
	return NSColorFromID(rv)
}

// Creates a new color object representing the color of the current color
// object in the specified color space.
//
// space: The color space of the new [NSColor] object.
//
// # Return Value
// 
// The new [NSColor] object. This method converts the receiver’s color to an
// equivalent one in the new color space. Although the new color might have
// different component values, it looks the same as the original. Returns
// `nil` if conversion is not possible.
//
// # Discussion
// 
// If the receiver’s color space is the same as that specified in `space`,
// this method returns the same [NSColor] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/usingColorSpace(_:)
func (c NSColor) ColorUsingColorSpace(space INSColorSpace) INSColor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("colorUsingColorSpace:"), space)
	return NSColorFromID(rv)
}

// Creates a new color object whose component values are a weighted sum of the
// current color object and the specified color object’s.
//
// fraction: The amount of the color to blend with the receiver’s color. The method
// converts `color` and a copy of the receiver to RGB, and then sets each
// component of the returned color to `fraction` of `color`‘s value plus 1
// – `fraction` of the receiver’s.
//
// color: The color to blend with the receiver’s color.
//
// # Return Value
// 
// The resulting color object or `nil` if the colors can’t be converted.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/blended(withFraction:of:)
func (c NSColor) BlendedColorWithFractionOfColor(fraction float64, color INSColor) INSColor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("blendedColorWithFraction:ofColor:"), fraction, color)
	return NSColorFromID(rv)
}

// Creates a new color object that has the same color space and component
// values as the current color object, but the specified alpha component.
//
// alpha: The opacity value of the new [NSColor] object.
//
// # Return Value
// 
// The new [NSColor] object. If the receiver’s color space doesn’t include
// an alpha component, the receiver is returned.
//
// # Discussion
// 
// A subclass with explicit opacity components should override this method to
// return a color with the specified alpha.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/withAlphaComponent(_:)
func (c NSColor) ColorWithAlphaComponent(alpha float64) INSColor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("colorWithAlphaComponent:"), alpha)
	return NSColorFromID(rv)
}

// Creates a new color object that represents a blend between the current
// color and the highlight color.
//
// val: The amount of the highlight color that is blended with the receiver’s
// color. This should be a number from `0.0` through `1.0`. A `highlightLevel`
// below `0.0` is interpreted as `0.0`; a `highlightLevel` above `1.0` is
// interpreted as `1.0`.
//
// # Return Value
// 
// The new [NSColor] object. Returns `nil` if the colors can’t be converted.
//
// # Discussion
// 
// The highlight color is provided by the [highlightColor] property. Call this
// method when you want to brighten the current color for use in highlights.
//
// [highlightColor]: https://developer.apple.com/documentation/AppKit/NSColor/highlightColor
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/highlight(withLevel:)
func (c NSColor) HighlightWithLevel(val float64) INSColor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("highlightWithLevel:"), val)
	return NSColorFromID(rv)
}

// Creates a new color object that represents a blend between the current
// color and the shadow color.
//
// val: The amount of the shadow color used for the blend. This should be a number
// from `0.0` through `1.0`. A `shadowLevel` below `0.0` is interpreted as
// `0.0`; a `shadowLevel` above `1.0` is interpreted as `1.0`.
//
// # Return Value
// 
// The new [NSColor] object. Returns `nil` if the colors can’t be converted.
//
// # Discussion
// 
// The shadow color is provided by the [shadowColor] property. Call this
// method when you want to darken the current color for use in shadows.
//
// [shadowColor]: https://developer.apple.com/documentation/AppKit/NSColor/shadowColor
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/shadow(withLevel:)
func (c NSColor) ShadowWithLevel(val float64) INSColor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("shadowWithLevel:"), val)
	return NSColorFromID(rv)
}

// Writes the color object’s data to the specified pasteboard.
//
// pasteBoard: The pasteboard to which to write the receiver’s color data. If this
// pasteboard doesn’t support color data, the method does nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/write(to:)
func (c NSColor) WriteToPasteboard(pasteBoard INSPasteboard) {
	objc.Send[objc.ID](c.ID, objc.Sel("writeToPasteboard:"), pasteBoard)
}

// Returns the color object’s CMYK and opacity values.
//
// cyan: Upon return, contains the cyan component of the color object.
//
// magenta: Upon return, contains the magenta component of the color object.
//
// yellow: Upon return, contains the yellow component of the color object.
//
// black: Upon return, contains the black component of the color object.
//
// alpha: Upon return, contains opacity value of the color object.
//
// # Discussion
// 
// If [NULL] is passed in as an argument, the method doesn’t set that value.
// This method works only with objects representing colors in the
// [deviceCMYK]. Sending it to other objects raises an exception.
//
// [deviceCMYK]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceCMYK
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/getCyan(_:magenta:yellow:black:alpha:)
func (c NSColor) GetCyanMagentaYellowBlackAlpha(cyan float64, magenta float64, yellow float64, black float64, alpha float64) {
	objc.Send[objc.ID](c.ID, objc.Sel("getCyan:magenta:yellow:black:alpha:"), cyan, magenta, yellow, black, alpha)
}

// Returns the color object’s HSB component and opacity values in the
// respective arguments.
//
// hue: Upon return, contains the hue component of the color object.
//
// saturation: Upon return, contains the saturation component of the color object.
//
// brightness: Upon return, contains the brightness component of the color object.
//
// alpha: Upon return, contains the opacity value of the color object.
//
// # Discussion
// 
// If [NULL] is passed in as an argument, the method doesn’t set that value.
// This method works only with objects representing colors in the
// [calibratedRGB] or [deviceRGB] color space. Sending it to other objects
// raises an exception.
//
// [calibratedRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedRGB
// [deviceRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceRGB
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/getHue(_:saturation:brightness:alpha:)
func (c NSColor) GetHueSaturationBrightnessAlpha(hue float64, saturation float64, brightness float64, alpha float64) {
	objc.Send[objc.ID](c.ID, objc.Sel("getHue:saturation:brightness:alpha:"), hue, saturation, brightness, alpha)
}

// Returns the color object’s RGB component and opacity values in the
// respective arguments.
//
// red: Upon return, contains the red component of the color object.
//
// green: Upon return, contains the green component of the color object.
//
// blue: Upon return, contains the blue component of the color object.
//
// alpha: Upon return, contains the opacity value of the color object.
//
// # Discussion
// 
// If [NULL] is passed in as an argument, the method doesn’t set that value.
// This method works only with objects representing colors in the
// [calibratedRGB] or [deviceRGB] color space. Sending it to other objects
// raises an exception.
//
// [calibratedRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedRGB
// [deviceRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceRGB
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/getRed(_:green:blue:alpha:)
func (c NSColor) GetRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) {
	objc.Send[objc.ID](c.ID, objc.Sel("getRed:green:blue:alpha:"), red, green, blue, alpha)
}

// Returns the grayscale and alpha values of the color.
//
// white: Upon return, contains the grayscale value of the color object.
//
// alpha: Upon return, contains the opacity value of the color object.
//
// # Discussion
// 
// If [NULL] is passed in as an argument, the method doesn’t set that value.
// This method works only with objects representing colors in the
// [calibratedWhite], [NSCalibratedBlackColorSpace],
// [NSDeviceBlackColorSpace], or [deviceWhite] color space. Sending it to
// other objects raises an exception.
//
// [NSCalibratedBlackColorSpace]: https://developer.apple.com/documentation/AppKit/NSCalibratedBlackColorSpace
// [NSDeviceBlackColorSpace]: https://developer.apple.com/documentation/AppKit/NSDeviceBlackColorSpace
// [calibratedWhite]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedWhite
// [deviceWhite]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceWhite
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/getWhite(_:alpha:)
func (c NSColor) GetWhiteAlpha(white float64, alpha float64) {
	objc.Send[objc.ID](c.ID, objc.Sel("getWhite:alpha:"), white, alpha)
}

// Returns the components of the color as an array.
//
// components: An array containing the components of the color object as `float` values.
//
// # Discussion
// 
// You can invoke this method on [NSColor] objects created from custom color
// spaces to get the individual floating-point components, including alpha.
// Raises an exception if the receiver doesn’t have floating-point
// components. To find out how many components are in the `components` array,
// use the [NumberOfComponents] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/getComponents(_:)
func (c NSColor) GetComponents(components float64) {
	objc.Send[objc.ID](c.ID, objc.Sel("getComponents:"), components)
}

// Returns a version of the color object that is compatible with the specified
// color type.
//
// type: The type of color object that you want. For example, if you want a color
// object containing RGB components, specify
// [NSColor.ColorType.componentBased].
// //
// [NSColor.ColorType.componentBased]: https://developer.apple.com/documentation/AppKit/NSColor/ColorType/componentBased
//
// # Return Value
// 
// A compatible color object, or `nil` if a compatible color object is not
// available.
//
// # Discussion
// 
// Before accessing the details of an [NSColor] object, use this method to
// ensure that you have an object capable of returning those details. For
// example, before you access the component values, make sure you have a color
// object of type [NSColor.ColorType.componentBased]. For some types of
// colors, conversions to a compatible color type may be possible.
//
// [NSColor.ColorType.componentBased]: https://developer.apple.com/documentation/AppKit/NSColor/ColorType/componentBased
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/usingType(_:)
func (c NSColor) ColorUsingType(type_ NSColorType) INSColor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("colorUsingType:"), type_)
	return NSColorFromID(rv)
}

// Reinterpret the color by applying a new `contentHeadroom` without changing
// the color components. Changing the `contentHeadroom` redefines the color
// relative to a different peak white, changing its behavior under tone
// mapping and the result of calling `standardDynamicRangeColor`. The new
// color will have a `contentHeadroom` >= 1.0. If called on a color with a
// color space that does not support extended range, or does not have an
// equivalent extended range counterpart, this will return `self`.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/applyingContentHeadroom(_:)
func (c NSColor) ColorByApplyingContentHeadroom(contentHeadroom float64) INSColor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("colorByApplyingContentHeadroom:"), contentHeadroom)
	return NSColorFromID(rv)
}

// Draws the current color in the specified rectangle.
//
// rect: The rectangle in which to draw the color.
//
// # Discussion
// 
// Subclasses adorn the rectangle in some manner to indicate the type of
// color. This method is invoked by color wells, swatches, and other user
// interface objects that need to display colors.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/drawSwatch(in:)
func (c NSColor) DrawSwatchInRect(rect corefoundation.CGRect) {
	objc.Send[objc.ID](c.ID, objc.Sel("drawSwatchInRect:"), rect)
}

// Sets the color of subsequent drawing to the color that the color object
// represents.
//
// # Discussion
// 
// This method should be implemented in subclasses.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/set()
func (c NSColor) Set() {
	objc.Send[objc.ID](c.ID, objc.Sel("set"))
}

// Sets the fill color of subsequent drawing to the color object’s color.
//
// # Discussion
// 
// This method should be implemented in subclasses.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/setFill()
func (c NSColor) SetFill() {
	objc.Send[objc.ID](c.ID, objc.Sel("setFill"))
}

// Sets the stroke color of subsequent drawing to the color object’s color.
//
// # Discussion
// 
// This method should be implemented in subclasses.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/setStroke()
func (c NSColor) SetStroke() {
	objc.Send[objc.ID](c.ID, objc.Sel("setStroke"))
}

// Creates a color object from data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(coder:)
func (c NSColor) InitWithCoder(coder foundation.INSCoder) NSColor {
	rv := objc.Send[NSColor](c.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Initializes an instance with a property list object and a type string.
//
// propertyList: A property list containing data to initialize the receiver.
// 
// By default, the property list object is an instance of [NSData]. If you
// implement [ReadingOptionsForTypePasteboard] and specify an option other
// than [PasteboardReadingAsData], the `propertyList` may be any other
// property list object.
//
// type: A UTI supported by the receiver for reading (one of the types returned by
// [ReadableTypesForPasteboard]).
//
// # Return Value
// 
// An object initialized using the data in `propertyList`.
//
// # Discussion
// 
// This method is considered optional because, if [ReadableTypesForPasteboard]
// returns just a single type, and that type uses the
// [PasteboardReadingAsKeyedArchive] reading option, then instances are
// initialized using [init(coder:)] instead of this method.
//
// [init(coder:)]: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/init(pasteboardPropertyList:ofType:)
func (c NSColor) InitWithPasteboardPropertyListOfType(propertyList objectivec.IObject, type_ NSPasteboardType) NSColor {
	rv := objc.Send[NSColor](c.ID, objc.Sel("initWithPasteboardPropertyList:ofType:"), propertyList, objc.String(string(type_)))
	return rv
}

// Returns a property list object to represent the receiver on a pasteboard as
// an object of a specified type.
//
// type: One of the types the receiver supports for writing (one of the UTIs
// returned by its implementation of [WritableTypesForPasteboard]).
//
// # Return Value
// 
// A property list object to represent the receiver on a pasteboard as an
// object of type `type`.
//
// # Discussion
// 
// The returned value will commonly be the [NSData] object for the specified
// data type. However, if this method returns either a string, or any other
// property-list type, the pasteboard will automatically convert these items
// to the correct data format required for the pasteboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/pasteboardPropertyList(forType:)
func (c NSColor) PasteboardPropertyListForType(type_ NSPasteboardType) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("pasteboardPropertyListForType:"), objc.String(string(type_)))
	return objectivec.Object{ID: rv}
}

// Returns an array of UTI strings of data types the receiver can write to a
// given pasteboard.
//
// pasteboard: A pasteboard.
// 
// You can use this argument to provide different options based on the
// pasteboard name, if you need to.
//
// # Return Value
// 
// An array of UTI strings of data types the receiver can write to
// `pasteboard`.
//
// # Discussion
// 
// By default, data for the first returned type is put onto the pasteboard
// immediately, with the remaining types being promised.
// 
// To change the default behavior, implement
// -writingOptionsForType:pasteboard: and return [PasteboardWritingPromised]
// to lazily provide data for types, return no option to provide the data for
// that type immediately. Use the pasteboard argument to provide different
// types based on the pasteboard name, if desired. Do not perform other
// pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/writableTypes(for:)
func (c NSColor) WritableTypesForPasteboard(pasteboard INSPasteboard) []string {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("writableTypesForPasteboard:"), pasteboard)
	return objc.ConvertSliceToStrings(rv)
}

// Returns options for writing data of a specified type to a given pasteboard.
//
// type: One of the types the receiver supports for writing (one of the UTIs
// returned by its implementation of [WritableTypesForPasteboard]).
//
// pasteboard: A pasteboard.
// 
// You can use this argument to provide different options based on the
// pasteboard name, if you need to.
//
// # Return Value
// 
// Options for writing data of type type to `pasteboard`. Return `0` for no
// options, or a value given in [Pasteboard Writing Options].
//
// [Pasteboard Writing Options]: https://developer.apple.com/documentation/AppKit/pasteboard-writing-options
//
// # Discussion
// 
// Do not perform other pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/writingOptions(forType:pasteboard:)
func (c NSColor) WritingOptionsForTypePasteboard(type_ NSPasteboardType, pasteboard INSPasteboard) NSPasteboardWritingOptions {
	rv := objc.Send[NSPasteboardWritingOptions](c.ID, objc.Sel("writingOptionsForType:pasteboard:"), objc.String(string(type_)), pasteboard)
	return NSPasteboardWritingOptions(rv)
}
func (c NSColor) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](c.ID, objc.Sel("encodeWithCoder:"), coder)
}





// Creates a dynamic catalog color with a provider that’s used to resolve
// the exact color value, calculated on first use.
//
// # Discussion
// 
// When methods on a color need color component values, AppKit calls the
// provider with [currentDrawingAppearance]. The provider can use the
// appearance to return another color to use for drawing. For example, if you
// create a color matching [systemYellow] for Light appearance, you’ll get a
// color matching [systemYellow] for Dark appearance when using Dark Mode. As
// often as possible, use the given appearance.
// 
// The `colorName` is equal to the identity of the color and should be
// universally unique; if `nil`, the system generates a unique name.
// 
// The name and the Light (aqua) appearance are encoded as part of this color.
// If a color is already registered with the same name as the decoded color,
// the decoded color joins the other color and becomes dynamic again.
//
// [currentDrawingAppearance]: https://developer.apple.com/documentation/AppKit/NSAppearance/currentDrawingAppearance
// [systemYellow]: https://developer.apple.com/documentation/AppKit/NSColor/systemYellow
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(name:dynamicProvider:)
func (_NSColorClass NSColorClass) ColorWithNameDynamicProvider(colorName string, dynamicProvider AppearanceHandler) NSColor {
		_block1, _cleanup1 := NewAppearanceBlock(dynamicProvider)
	defer _cleanup1()
		rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("colorWithName:dynamicProvider:"), objc.String(colorName), _block1)
	return NSColorFromID(rv)
}

// Returns an array of uniform type identifier strings of data types the
// receiver can read from the pasteboard and initialize from.
//
// pasteboard: A pasteboard. You can use the pasteboard argument to provide different
// types based on the pasteboard name, if you need to.
//
// # Return Value
// 
// An array of uniform type identifier strings of data types instances that
// the receiver can read from the pasteboard and initialize from.
//
// # Discussion
// 
// By default, the system provides the data for a type to
// [InitWithPasteboardPropertyListOfType] as an instance of [NSData]. If you
// implement [ReadingOptionsForTypePasteboard] and specify a different option,
// the system converts the [NSData] object for a type to an [NSString] object
// or any other property list object.
// 
// # Special Considerations
// 
// Don’t perform other pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/readableTypes(for:)
func (_NSColorClass NSColorClass) ReadableTypesForPasteboard(pasteboard INSPasteboard) []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSColorClass.class), objc.Sel("readableTypesForPasteboard:"), pasteboard)
	return objc.ConvertSliceToStrings(rv)
}

// Returns options for reading data of a specified type from a given
// pasteboard.
//
// type: A UTI supported by instances of the receiver for reading (one of the types
// returned by [ReadableTypesForPasteboard]).
//
// pasteboard: A pasteboard.
// 
// You can use the pasteboard argument to provide return different based on
// the pasteboard name, should you need to do so.
//
// # Return Value
// 
// Options for reading data of `type` from `pasteboard`. For a list of valid
// values, see [NSPasteboard.ReadingOptions].
//
// [NSPasteboard.ReadingOptions]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions
//
// # Discussion
// 
// Do not perform other pasteboard operations in this method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/readingOptions(forType:pasteboard:)
func (_NSColorClass NSColorClass) ReadingOptionsForTypePasteboard(type_ NSPasteboardType, pasteboard INSPasteboard) NSPasteboardReadingOptions {
	rv := objc.Send[NSPasteboardReadingOptions](objc.ID(_NSColorClass.class), objc.Sel("readingOptionsForType:pasteboard:"), objc.String(string(type_)), pasteboard)
	return NSPasteboardReadingOptions(rv)
}








// The number of components in the color.
//
// # Discussion
// 
// This property reflects the number of floating-point component values in the
// color and includes the alpha component. If the color object does not have
// any floating-point component values, accessing this property raises an
// exception.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/numberOfComponents
func (c NSColor) NumberOfComponents() int {
	rv := objc.Send[int](c.ID, objc.Sel("numberOfComponents"))
	return rv
}



// The alpha (opacity) component value of the color.
//
// # Discussion
// 
// The value of this property is between `0.0` and `1.0`, where `0.0`
// represents fully transparent and `1.0` represents fully opaque. If the
// color has no alpha component, the value of this property is `1.0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/alphaComponent
func (c NSColor) AlphaComponent() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("alphaComponent"))
	return rv
}



// The white component value of the color.
//
// # Discussion
// 
// Access this property only for colors in the [calibratedWhite],
// [NSCalibratedBlackColorSpace], [NSDeviceBlackColorSpace], or [deviceWhite]
// color spaces. Accessing it for other color types raises an exception.
//
// [NSCalibratedBlackColorSpace]: https://developer.apple.com/documentation/AppKit/NSCalibratedBlackColorSpace
// [NSDeviceBlackColorSpace]: https://developer.apple.com/documentation/AppKit/NSDeviceBlackColorSpace
// [calibratedWhite]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedWhite
// [deviceWhite]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceWhite
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/whiteComponent
func (c NSColor) WhiteComponent() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("whiteComponent"))
	return rv
}



// The red component value of the color.
//
// # Discussion
// 
// Access this property only for colors in the [calibratedRGB] or [deviceRGB]
// color spaces. Accessing it for other color types raises an exception.
//
// [calibratedRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedRGB
// [deviceRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceRGB
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/redComponent
func (c NSColor) RedComponent() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("redComponent"))
	return rv
}



// The green component value of the color.
//
// # Discussion
// 
// Access this property only for colors in the [calibratedRGB] or [deviceRGB]
// color spaces. Accessing it for other color types raises an exception.
//
// [calibratedRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedRGB
// [deviceRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceRGB
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/greenComponent
func (c NSColor) GreenComponent() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("greenComponent"))
	return rv
}



// The blue component value of the color.
//
// # Discussion
// 
// Access this property only for colors in the [calibratedRGB] or [deviceRGB]
// color spaces. Accessing it for other color types raises an exception.
//
// [calibratedRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedRGB
// [deviceRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceRGB
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/blueComponent
func (c NSColor) BlueComponent() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("blueComponent"))
	return rv
}



// The cyan component value of the color.
//
// # Discussion
// 
// Access this property only for colors in the [deviceCMYK] color space.
// Accessing it for other color types raises an exception.
//
// [deviceCMYK]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceCMYK
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/cyanComponent
func (c NSColor) CyanComponent() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("cyanComponent"))
	return rv
}



// The magenta component value of the color.
//
// # Discussion
// 
// Access this property only for colors in the [deviceCMYK] color space.
// Accessing it for other color types raises an exception.
//
// [deviceCMYK]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceCMYK
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/magentaComponent
func (c NSColor) MagentaComponent() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("magentaComponent"))
	return rv
}



// The yellow component value of the color.
//
// # Discussion
// 
// Access this property only for colors in the [deviceCMYK] color space.
// Accessing it for other color types raises an exception.
//
// [deviceCMYK]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceCMYK
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/yellowComponent
func (c NSColor) YellowComponent() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("yellowComponent"))
	return rv
}



// The black component value of the color.
//
// # Discussion
// 
// Access this property only for colors in the [deviceCMYK] color space, which
// have a black component. Accessing it for other color types raises an
// exception.
//
// [deviceCMYK]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceCMYK
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/blackComponent
func (c NSColor) BlackComponent() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("blackComponent"))
	return rv
}



// The hue component value of the color.
//
// # Discussion
// 
// Access this property only for colors in the [calibratedRGB] or [deviceRGB]
// color spaces. RGB values are converted to HSB values as needed. Accessing
// it for other color types raises an exception.
//
// [calibratedRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedRGB
// [deviceRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceRGB
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/hueComponent
func (c NSColor) HueComponent() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("hueComponent"))
	return rv
}



// The saturation component value of the color.
//
// # Discussion
// 
// Access this property only for colors in the [calibratedRGB] or [deviceRGB]
// color spaces. RGB values are converted to HSB values as needed. Accessing
// it for other color types raises an exception.
//
// [calibratedRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedRGB
// [deviceRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceRGB
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/saturationComponent
func (c NSColor) SaturationComponent() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("saturationComponent"))
	return rv
}



// The brightness component value of the color.
//
// # Discussion
// 
// Access this property only for colors in the [calibratedRGB] or [deviceRGB]
// color spaces. RGB values are converted to HSB values as needed. Accessing
// it for other color types raises an exception.
//
// [calibratedRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedRGB
// [deviceRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceRGB
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/brightnessComponent
func (c NSColor) BrightnessComponent() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("brightnessComponent"))
	return rv
}



// The catalog containing the color’s name.
//
// # Discussion
// 
// Access this property only for colors in the [named] color space. Accessing
// it for other color types raises an exception.
//
// [named]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/named
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/catalogNameComponent
func (c NSColor) CatalogNameComponent() NSColorListName {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("catalogNameComponent"))
	return NSColorListName(foundation.NSStringFromID(rv).String())
}



// The localized version of the catalog name containing the color.
//
// # Discussion
// 
// Use the value in this property when displaying the catalog name in your
// user interface. Access this property only for colors in the [named] color
// space. Accessing it for other color types raises an exception.
//
// [named]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/named
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/localizedCatalogNameComponent
func (c NSColor) LocalizedCatalogNameComponent() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("localizedCatalogNameComponent"))
	return foundation.NSStringFromID(rv).String()
}



// The name of the color.
//
// # Discussion
// 
// Access this property only for colors in the [named] color space. Accessing
// it for other color types raises an exception.
//
// [named]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/named
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/colorNameComponent
func (c NSColor) ColorNameComponent() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("colorNameComponent"))
	return foundation.NSStringFromID(rv).String()
}



// The localized version of the color name.
//
// # Discussion
// 
// Use the value in this property when displaying the color name in your user
// interface. Access this property only for colors in the [named] color space.
// Accessing it for other color types raises an exception.
//
// [named]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/named
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/localizedColorNameComponent
func (c NSColor) LocalizedColorNameComponent() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("localizedColorNameComponent"))
	return foundation.NSStringFromID(rv).String()
}



// The type of the color object.
//
// # Discussion
// 
// A color object’s type determines which of its methods and properties you
// may access. For example, if the type is [NSColor.ColorType.pattern], you
// may safely access the [patternImage] property.
//
// [NSColor.ColorType.pattern]: https://developer.apple.com/documentation/AppKit/NSColor/ColorType/pattern
// [patternImage]: https://developer.apple.com/documentation/AppKit/NSColor/patternImage
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/type
func (c NSColor) Type() NSColorType {
	rv := objc.Send[NSColorType](c.ID, objc.Sel("type"))
	return NSColorType(rv)
}



// The color space associated with the color.
//
// # Discussion
// 
// Access this property only for colors that have an associated color
// space—specifically, colors not created by name or using a pattern image.
// Accessing it for other color types raises an exception. If you are unsure
// about a color object, convert it to an equivalent [NSColorSpace]-based
// object before calling this method.
// 
// It is safe to access this property for color objects created with the color
// space names [calibratedWhite], [NSCalibratedBlackColorSpace],
// [calibratedRGB], [deviceWhite], [NSDeviceBlackColorSpace], [deviceRGB],
// [deviceCMYK], or [custom]—or with the [NSColorSpace] class methods
// corresponding to these names.
//
// [NSCalibratedBlackColorSpace]: https://developer.apple.com/documentation/AppKit/NSCalibratedBlackColorSpace
// [NSDeviceBlackColorSpace]: https://developer.apple.com/documentation/AppKit/NSDeviceBlackColorSpace
// [calibratedRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedRGB
// [calibratedWhite]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedWhite
// [custom]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/custom
// [deviceCMYK]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceCMYK
// [deviceRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceRGB
// [deviceWhite]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceWhite
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/colorSpace
func (c NSColor) ColorSpace() INSColorSpace {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("colorSpace"))
	return NSColorSpaceFromID(objc.ID(rv))
}



// For HDR colors, the linear brightness multiplier that was applied when
// generating the color. Colors created with an exposure by NSColor create
// CGColors that are tagged with a contentHeadroom value. While CGColors
// created without a contentHeadroom tag will return 0 from
// CGColorGetHeadroom, NSColors generated in a similar fashion return a
// linearExposure of 1.0.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/linearExposure
func (c NSColor) LinearExposure() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("linearExposure"))
	return rv
}



// In some cases it is useful to recover the color that was base the SDR color
// that was exposed to generate an HDR color. If a color’s `linearExposure`
// is > 1, then this will return the base SDR color. If the color is not an
// HDR color, this will return `self`.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/standardDynamicRange
func (c NSColor) StandardDynamicRangeColor() INSColor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("standardDynamicRangeColor"))
	return NSColorFromID(objc.ID(rv))
}



// The Core Graphics color object corresponding to the color.
//
// # Discussion
// 
// This property always contains a valid color, even though the value may be
// an approximation in some cases. There is no guaranteed round-trip color
// fidelity.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/cgColor
func (c NSColor) CGColor() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](c.ID, objc.Sel("CGColor"))
	return coregraphics.CGColorRef(rv)
}



// The pattern image used to paint the target area.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/patternimage
func (c NSColor) PatternImage() INSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("patternImage"))
	return NSImageFromID(objc.ID(rv))
}
func (c NSColor) SetPatternImage(value INSImage) {
	objc.Send[struct{}](c.ID, objc.Sel("setPatternImage:"), value)
}



// Returns a localized description of the color for use in accessibility
// attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityColor/accessibilityName
func (c NSColor) AccessibilityName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("accessibilityName"))
	return foundation.NSStringFromID(rv).String()
}







// Sent when the system colors have changed, such as through a system control
// panel interface.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/systemcolorsdidchangenotification
func (_NSColorClass NSColorClass) SystemColorsDidChangeNotification() foundation.NSString {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("NSSystemColorsDidChangeNotification"))
	return foundation.NSStringFromID(objc.ID(rv))
}



// A Boolean value that indicates whether the app supports alpha.
//
// # Return Value
// 
// [true] if the app doesn’t support alpha; otherwise, [false].
// 
// # Discussion
// 
// The system consults this global value when the app imports alpha (for
// instance, through color dragging). By default this property is [false];
// meaning the system supports the alpha component for colors globally. To
// ignore alpha for an app, invoke the `setIgnoresAlpha` method with a
// parameter of [true]. This value also determines whether the color panel has
// an opacity slider.
// 
// This method provides a global approach for removing alpha, which might not
// always be appropriate. Apps that need to disable alpha can use more
// fine-grained APIs for individual controls, such as [ShowsAlpha] and
// [SupportsAlpha].
// 
// In macOS 13 and earlier, the default value is [true]. This property is
// deprecated as of macOS 14.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/ignoresAlpha
func (_NSColorClass NSColorClass) IgnoresAlpha() bool {
	rv := objc.Send[bool](objc.ID(_NSColorClass.class), objc.Sel("ignoresAlpha"))
	return rv
}
func (_NSColorClass NSColorClass) SetIgnoresAlpha(value bool) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setIgnoresAlpha:"), value)
}



// Sent after the user changes control tint preference.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/currentcontroltintdidchangenotification
func (_NSColorClass NSColorClass) CurrentControlTintDidChangeNotification() foundation.NSString {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("NSControlTintDidChangeNotification"))
	return foundation.NSStringFromID(objc.ID(rv))
}



// Returns a color object for blue that automatically adapts to vibrancy and
// accessibility settings.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/systemblue
func (_NSColorClass NSColorClass) SystemBlue() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemBlueColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSystemBlue(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSystemBlueColor:"), value)
}



// A color space object that represents an extended gray color space with a
// gamma value of 2.2.
//
// See: https://developer.apple.com/documentation/appkit/nscolorspace/extendedgenericgamma22gray
func (_NSColorClass NSColorClass) ExtendedGenericGamma22Gray() NSColorSpace {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("extendedGenericGamma22GrayColorSpace"))
	return NSColorSpaceFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetExtendedGenericGamma22Gray(value NSColorSpace) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setExtendedGenericGamma22GrayColorSpace:"), value)
}



// A color space object that represents an extended sRGB color space.
//
// See: https://developer.apple.com/documentation/appkit/nscolorspace/extendedsrgb
func (_NSColorClass NSColorClass) ExtendedSRGB() NSColorSpace {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("extendedSRGBColorSpace"))
	return NSColorSpaceFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetExtendedSRGB(value NSColorSpace) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setExtendedSRGBColorSpace:"), value)
}



// The system color used for the face of a selected control in a list or
// table.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/alternateselectedcontrolcolor
func (_NSColorClass NSColorClass) AlternateSelectedControlColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("alternateSelectedControlColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetAlternateSelectedControlColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setAlternateSelectedControlColor:"), value)
}



// The color to use for text in a selected control.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/alternateselectedcontroltextcolor
func (_NSColorClass NSColorClass) AlternateSelectedControlTextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("alternateSelectedControlTextColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetAlternateSelectedControlTextColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setAlternateSelectedControlTextColor:"), value)
}



// The colors to use for alternating content, typically found in table views
// and collection views.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/alternatingcontentbackgroundcolors
func (_NSColorClass NSColorClass) AlternatingContentBackgroundColors() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("alternatingContentBackgroundColors"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetAlternatingContentBackgroundColors(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setAlternatingContentBackgroundColors:"), value)
}



// The user’s current accent color preference.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/controlaccentcolor
func (_NSColorClass NSColorClass) ControlAccentColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("controlAccentColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetControlAccentColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setControlAccentColor:"), value)
}



// An array containing the system specified background colors for alternating
// rows in tables and lists.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/controlalternatingrowbackgroundcolors
func (_NSColorClass NSColorClass) ControlAlternatingRowBackgroundColors() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("controlAlternatingRowBackgroundColors"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetControlAlternatingRowBackgroundColors(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setControlAlternatingRowBackgroundColors:"), value)
}



// The color to use for the background of large controls, such as scroll views
// or table views.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/controlbackgroundcolor
func (_NSColorClass NSColorClass) ControlBackgroundColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("controlBackgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetControlBackgroundColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setControlBackgroundColor:"), value)
}



// The color to use for the flat surfaces of a control.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/controlcolor
func (_NSColorClass NSColorClass) ControlColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("controlColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetControlColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setControlColor:"), value)
}



// The system color used for the dark edge of the shadow dropped from
// controls.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/controldarkshadowcolor
func (_NSColorClass NSColorClass) ControlDarkShadowColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("controlDarkShadowColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetControlDarkShadowColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setControlDarkShadowColor:"), value)
}



// The system color used for the highlighted bezels of controls.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/controlhighlightcolor
func (_NSColorClass NSColorClass) ControlHighlightColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("controlHighlightColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetControlHighlightColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setControlHighlightColor:"), value)
}



// The system color used for light highlights in controls.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/controllighthighlightcolor
func (_NSColorClass NSColorClass) ControlLightHighlightColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("controlLightHighlightColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetControlLightHighlightColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setControlLightHighlightColor:"), value)
}



// The system color used for the shadows dropped from controls.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/controlshadowcolor
func (_NSColorClass NSColorClass) ControlShadowColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("controlShadowColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetControlShadowColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setControlShadowColor:"), value)
}



// The color to use for text on enabled controls.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/controltextcolor
func (_NSColorClass NSColorClass) ControlTextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("controlTextColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetControlTextColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setControlTextColor:"), value)
}



// The current system control tint color.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/currentcontroltint
func (_NSColorClass NSColorClass) CurrentControlTint() NSControlTint {
	rv := objc.Send[NSControlTint](objc.ID(_NSColorClass.class), objc.Sel("currentControlTint"))
	return NSControlTint(rv)
}
func (_NSColorClass NSColorClass) SetCurrentControlTint(value NSControlTint) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setCurrentControlTint:"), value)
}



// The color to use for text on disabled controls.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/disabledcontroltextcolor
func (_NSColorClass NSColorClass) DisabledControlTextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("disabledControlTextColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetDisabledControlTextColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setDisabledControlTextColor:"), value)
}



// The highlight color to use for the bubble that shows inline search result
// values.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/findhighlightcolor
func (_NSColorClass NSColorClass) FindHighlightColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("findHighlightColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetFindHighlightColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setFindHighlightColor:"), value)
}



// The color to use for the optional gridlines, such as those in a table view.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/gridcolor
func (_NSColorClass NSColorClass) GridColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("gridColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetGridColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setGridColor:"), value)
}



// The system color used as the background color for header cells in table
// views and outline views.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/headercolor
func (_NSColorClass NSColorClass) HeaderColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("headerColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetHeaderColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setHeaderColor:"), value)
}



// The color to use for text in header cells in table views and outline views.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/headertextcolor
func (_NSColorClass NSColorClass) HeaderTextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("headerTextColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetHeaderTextColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setHeaderTextColor:"), value)
}



// The color to use as a virtual light source on the screen.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/highlightcolor
func (_NSColorClass NSColorClass) HighlightColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("highlightColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetHighlightColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setHighlightColor:"), value)
}



// The color to use for the keyboard focus ring around controls.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/keyboardfocusindicatorcolor
func (_NSColorClass NSColorClass) KeyboardFocusIndicatorColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("keyboardFocusIndicatorColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetKeyboardFocusIndicatorColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setKeyboardFocusIndicatorColor:"), value)
}



// The system color used for the flat surface of a slider knob that hasn’t
// been selected.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/knobcolor
func (_NSColorClass NSColorClass) KnobColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("knobColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetKnobColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setKnobColor:"), value)
}



// The primary color to use for text labels.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/labelcolor
func (_NSColorClass NSColorClass) LabelColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("labelColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetLabelColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setLabelColor:"), value)
}



// The color to use for links.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/linkcolor
func (_NSColorClass NSColorClass) LinkColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("linkColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetLinkColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setLinkColor:"), value)
}



// The color to use for placeholder text in controls or text views.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/placeholdertextcolor
func (_NSColorClass NSColorClass) PlaceholderTextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("placeholderTextColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetPlaceholderTextColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setPlaceholderTextColor:"), value)
}



// The quaternary color to use for text labels and separators.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/quaternarylabelcolor
func (_NSColorClass NSColorClass) QuaternaryLabelColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("quaternaryLabelColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetQuaternaryLabelColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setQuaternaryLabelColor:"), value)
}



// See: https://developer.apple.com/documentation/appkit/nscolor/quaternarysystemfill
func (_NSColorClass NSColorClass) QuaternarySystemFill() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("quaternarySystemFillColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetQuaternarySystemFill(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setQuaternarySystemFillColor:"), value)
}



// See: https://developer.apple.com/documentation/appkit/nscolor/quinarylabel
func (_NSColorClass NSColorClass) QuinaryLabel() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("quinaryLabelColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetQuinaryLabel(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setQuinaryLabelColor:"), value)
}



// See: https://developer.apple.com/documentation/appkit/nscolor/quinarysystemfill
func (_NSColorClass NSColorClass) QuinarySystemFill() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("quinarySystemFillColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetQuinarySystemFill(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setQuinarySystemFillColor:"), value)
}



// The system color used for scroll “bars”—that is, for the groove in
// which a scroller’s knob moves
//
// See: https://developer.apple.com/documentation/appkit/nscolor/scrollbarcolor
func (_NSColorClass NSColorClass) ScrollBarColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("scrollBarColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetScrollBarColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setScrollBarColor:"), value)
}



// The patterned color to use for the background of a scrubber control.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/scrubbertexturedbackground
func (_NSColorClass NSColorClass) ScrubberTexturedBackground() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("scrubberTexturedBackgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetScrubberTexturedBackground(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setScrubberTexturedBackgroundColor:"), value)
}



// The secondary color to use for text labels.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/secondarylabelcolor
func (_NSColorClass NSColorClass) SecondaryLabelColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("secondaryLabelColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSecondaryLabelColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSecondaryLabelColor:"), value)
}



// The color used for selected controls in non-key views.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/secondaryselectedcontrolcolor
func (_NSColorClass NSColorClass) SecondarySelectedControlColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("secondarySelectedControlColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSecondarySelectedControlColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSecondarySelectedControlColor:"), value)
}



// See: https://developer.apple.com/documentation/appkit/nscolor/secondarysystemfill
func (_NSColorClass NSColorClass) SecondarySystemFill() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("secondarySystemFillColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSecondarySystemFill(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSecondarySystemFillColor:"), value)
}



// The color to use for the background of selected and emphasized content.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/selectedcontentbackgroundcolor
func (_NSColorClass NSColorClass) SelectedContentBackgroundColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("selectedContentBackgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSelectedContentBackgroundColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSelectedContentBackgroundColor:"), value)
}



// The color to use for the face of a selected control—that is, a control
// that has been clicked or is being dragged.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/selectedcontrolcolor
func (_NSColorClass NSColorClass) SelectedControlColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("selectedControlColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSelectedControlColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSelectedControlColor:"), value)
}



// The color to use for text in a selected control—that is, a control being
// clicked or dragged.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/selectedcontroltextcolor
func (_NSColorClass NSColorClass) SelectedControlTextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("selectedControlTextColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSelectedControlTextColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSelectedControlTextColor:"), value)
}



// The system color used for the slider knob when it is selected.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/selectedknobcolor
func (_NSColorClass NSColorClass) SelectedKnobColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("selectedKnobColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSelectedKnobColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSelectedKnobColor:"), value)
}



// The color to use for the face of selected menu items.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/selectedmenuitemcolor
func (_NSColorClass NSColorClass) SelectedMenuItemColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("selectedMenuItemColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSelectedMenuItemColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSelectedMenuItemColor:"), value)
}



// The color to use for the text in menu items.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/selectedmenuitemtextcolor
func (_NSColorClass NSColorClass) SelectedMenuItemTextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("selectedMenuItemTextColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSelectedMenuItemTextColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSelectedMenuItemTextColor:"), value)
}



// The color to use for the background of selected text.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/selectedtextbackgroundcolor
func (_NSColorClass NSColorClass) SelectedTextBackgroundColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("selectedTextBackgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSelectedTextBackgroundColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSelectedTextBackgroundColor:"), value)
}



// The color to use for selected text.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/selectedtextcolor
func (_NSColorClass NSColorClass) SelectedTextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("selectedTextColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSelectedTextColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSelectedTextColor:"), value)
}



// The color to use for separators between different sections of content.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/separatorcolor
func (_NSColorClass NSColorClass) SeparatorColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("separatorColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSeparatorColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSeparatorColor:"), value)
}



// The color to use for virtual shadows cast by raised objects on the screen.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/shadowcolor
func (_NSColorClass NSColorClass) ShadowColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("shadowColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetShadowColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setShadowColor:"), value)
}



// See: https://developer.apple.com/documentation/appkit/nscolor/systemfill
func (_NSColorClass NSColorClass) SystemFill() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemFillColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSystemFill(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSystemFillColor:"), value)
}



// The tertiary color to use for text labels.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/tertiarylabelcolor
func (_NSColorClass NSColorClass) TertiaryLabelColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("tertiaryLabelColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetTertiaryLabelColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setTertiaryLabelColor:"), value)
}



// See: https://developer.apple.com/documentation/appkit/nscolor/tertiarysystemfill
func (_NSColorClass NSColorClass) TertiarySystemFill() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("tertiarySystemFillColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetTertiarySystemFill(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setTertiarySystemFillColor:"), value)
}



// The color to use for the background area behind text.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/textbackgroundcolor
func (_NSColorClass NSColorClass) TextBackgroundColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("textBackgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetTextBackgroundColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setTextBackgroundColor:"), value)
}



// The color to use for text.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/textcolor
func (_NSColorClass NSColorClass) TextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("textColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetTextColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setTextColor:"), value)
}



// See: https://developer.apple.com/documentation/appkit/nscolor/textinsertionpointcolor
func (_NSColorClass NSColorClass) TextInsertionPointColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("textInsertionPointColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetTextInsertionPointColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setTextInsertionPointColor:"), value)
}



// The color to use in the area beneath your window’s views.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/underpagebackgroundcolor
func (_NSColorClass NSColorClass) UnderPageBackgroundColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("underPageBackgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetUnderPageBackgroundColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setUnderPageBackgroundColor:"), value)
}



// The color to use for selected and unemphasized content.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/unemphasizedselectedcontentbackgroundcolor
func (_NSColorClass NSColorClass) UnemphasizedSelectedContentBackgroundColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("unemphasizedSelectedContentBackgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetUnemphasizedSelectedContentBackgroundColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setUnemphasizedSelectedContentBackgroundColor:"), value)
}



// The color to use for the text background in an unemphasized context.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/unemphasizedselectedtextbackgroundcolor
func (_NSColorClass NSColorClass) UnemphasizedSelectedTextBackgroundColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("unemphasizedSelectedTextBackgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetUnemphasizedSelectedTextBackgroundColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setUnemphasizedSelectedTextBackgroundColor:"), value)
}



// The color to use for selected text in an unemphasized context.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/unemphasizedselectedtextcolor
func (_NSColorClass NSColorClass) UnemphasizedSelectedTextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("unemphasizedSelectedTextColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetUnemphasizedSelectedTextColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setUnemphasizedSelectedTextColor:"), value)
}



// The color to use for the window background.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/windowbackgroundcolor
func (_NSColorClass NSColorClass) WindowBackgroundColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("windowBackgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetWindowBackgroundColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setWindowBackgroundColor:"), value)
}



// The system color used for window frames, except for their text.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/windowframecolor
func (_NSColorClass NSColorClass) WindowFrameColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("windowFrameColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetWindowFrameColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setWindowFrameColor:"), value)
}



// The color to use for text in a window’s frame.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/windowframetextcolor
func (_NSColorClass NSColorClass) WindowFrameTextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("windowFrameTextColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetWindowFrameTextColor(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setWindowFrameTextColor:"), value)
}



// Returns a color object whose grayscale value is
//
// See: https://developer.apple.com/documentation/appkit/nscolor/black
func (_NSColorClass NSColorClass) Black() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("blackColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetBlack(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setBlackColor:"), value)
}



// Returns a color object whose RGB value is
//
// See: https://developer.apple.com/documentation/appkit/nscolor/blue
func (_NSColorClass NSColorClass) Blue() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("blueColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetBlue(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setBlueColor:"), value)
}



// Returns a color object whose RGB value is
//
// See: https://developer.apple.com/documentation/appkit/nscolor/brown
func (_NSColorClass NSColorClass) Brown() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("brownColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetBrown(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setBrownColor:"), value)
}



// Returns a color object whose grayscale and alpha values are both
//
// See: https://developer.apple.com/documentation/appkit/nscolor/clear
func (_NSColorClass NSColorClass) Clear() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("clearColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetClear(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setClearColor:"), value)
}



// Returns a color object whose RGB value is
//
// See: https://developer.apple.com/documentation/appkit/nscolor/cyan
func (_NSColorClass NSColorClass) Cyan() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("cyanColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetCyan(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setCyanColor:"), value)
}



// Returns a color object whose grayscale value is
//
// See: https://developer.apple.com/documentation/appkit/nscolor/darkgray
func (_NSColorClass NSColorClass) DarkGray() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("darkGrayColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetDarkGray(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setDarkGrayColor:"), value)
}



// Returns a color object whose grayscale value is
//
// See: https://developer.apple.com/documentation/appkit/nscolor/gray
func (_NSColorClass NSColorClass) Gray() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("grayColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetGray(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setGrayColor:"), value)
}



// Returns a color object whose RGB value is
//
// See: https://developer.apple.com/documentation/appkit/nscolor/green
func (_NSColorClass NSColorClass) Green() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("greenColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetGreen(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setGreenColor:"), value)
}



// Returns a color object whose grayscale value is
//
// See: https://developer.apple.com/documentation/appkit/nscolor/lightgray
func (_NSColorClass NSColorClass) LightGray() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("lightGrayColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetLightGray(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setLightGrayColor:"), value)
}



// Returns a color object whose RGB value is
//
// See: https://developer.apple.com/documentation/appkit/nscolor/magenta
func (_NSColorClass NSColorClass) Magenta() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("magentaColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetMagenta(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setMagentaColor:"), value)
}



// Returns a color object whose RGB value is
//
// See: https://developer.apple.com/documentation/appkit/nscolor/orange
func (_NSColorClass NSColorClass) Orange() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("orangeColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetOrange(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setOrangeColor:"), value)
}



// Returns a color object whose RGB value is
//
// See: https://developer.apple.com/documentation/appkit/nscolor/purple
func (_NSColorClass NSColorClass) Purple() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("purpleColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetPurple(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setPurpleColor:"), value)
}



// Returns a color object whose RGB value is
//
// See: https://developer.apple.com/documentation/appkit/nscolor/red
func (_NSColorClass NSColorClass) Red() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("redColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetRed(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setRedColor:"), value)
}



// Returns a color object for brown that automatically adapts to vibrancy and
// accessibility settings.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/systembrown
func (_NSColorClass NSColorClass) SystemBrown() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemBrownColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSystemBrown(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSystemBrownColor:"), value)
}



// Returns a color object for cyan that automatically adapts to vibrancy and
// accessibility settings.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/systemcyan
func (_NSColorClass NSColorClass) SystemCyan() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemCyanColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSystemCyan(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSystemCyanColor:"), value)
}



// Returns a color object for gray that automatically adapts to vibrancy and
// accessibility settings.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/systemgray
func (_NSColorClass NSColorClass) SystemGray() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemGrayColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSystemGray(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSystemGrayColor:"), value)
}



// Returns a color object for green that automatically adapts to vibrancy and
// accessibility settings.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/systemgreen
func (_NSColorClass NSColorClass) SystemGreen() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemGreenColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSystemGreen(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSystemGreenColor:"), value)
}



// Returns a color object for indigo that automatically adapts to vibrancy and
// accessibility settings.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/systemindigo
func (_NSColorClass NSColorClass) SystemIndigo() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemIndigoColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSystemIndigo(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSystemIndigoColor:"), value)
}



// Returns a color object for mint that automatically adapts to vibrancy and
// accessibility settings.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/systemmint
func (_NSColorClass NSColorClass) SystemMint() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemMintColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSystemMint(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSystemMintColor:"), value)
}



// Returns a color object for orange that automatically adapts to vibrancy and
// accessibility settings.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/systemorange
func (_NSColorClass NSColorClass) SystemOrange() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemOrangeColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSystemOrange(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSystemOrangeColor:"), value)
}



// Returns a color object for pink that automatically adapts to vibrancy and
// accessibility settings.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/systempink
func (_NSColorClass NSColorClass) SystemPink() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemPinkColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSystemPink(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSystemPinkColor:"), value)
}



// Returns a color object for purple that automatically adapts to vibrancy and
// accessibility settings.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/systempurple
func (_NSColorClass NSColorClass) SystemPurple() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemPurpleColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSystemPurple(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSystemPurpleColor:"), value)
}



// Returns a color object for red that automatically adapts to vibrancy and
// accessibility settings.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/systemred
func (_NSColorClass NSColorClass) SystemRed() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemRedColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSystemRed(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSystemRedColor:"), value)
}



// Returns a color object for teal that automatically adapts to vibrancy and
// accessibility settings.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/systemteal
func (_NSColorClass NSColorClass) SystemTeal() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemTealColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSystemTeal(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSystemTealColor:"), value)
}



// Returns a color object for yellow that automatically adapts to vibrancy and
// accessibility settings.
//
// See: https://developer.apple.com/documentation/appkit/nscolor/systemyellow
func (_NSColorClass NSColorClass) SystemYellow() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemYellowColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetSystemYellow(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setSystemYellowColor:"), value)
}



// Returns a color object whose grayscale and alpha values are both
//
// See: https://developer.apple.com/documentation/appkit/nscolor/white
func (_NSColorClass NSColorClass) White() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("whiteColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetWhite(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setWhiteColor:"), value)
}



// Returns a color object whose RGB value is
//
// See: https://developer.apple.com/documentation/appkit/nscolor/yellow
func (_NSColorClass NSColorClass) Yellow() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("yellowColor"))
	return NSColorFromID(objc.ID(rv))
}
func (_NSColorClass NSColorClass) SetYellow(value NSColor) {
	objc.Send[struct{}](objc.ID(_NSColorClass.class), objc.Sel("setYellowColor:"), value)
}










			// Protocol methods for NSAccessibilityColor
			








			// Protocol methods for NSPasteboardWriting
			












// ColorWithNameDynamicProviderSync is a synchronous wrapper around [NSColor.ColorWithNameDynamicProvider].
// It blocks until the completion handler fires or the context is cancelled.
func (cc NSColorClass) ColorWithNameDynamicProviderSync(ctx context.Context, colorName string) (*NSAppearance, error) {
	done := make(chan *NSAppearance, 1)
	cc.ColorWithNameDynamicProvider(colorName, func(val *NSAppearance) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}






