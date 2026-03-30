// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/coreimage"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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

// Class returns the underlying Objective-C class pointer.
func (nc NSColorClass) Class() objc.Class {
	return nc.class
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
// [NSColor.SystemBlueColor], when you want a specific tint that looks correct in both
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
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
// [NSColor.CGColor]: https://developer.apple.com/documentation/CoreGraphics/CGColor
// [Color]: https://developer.apple.com/design/human-interface-guidelines/color/
// [UI element colors]: https://developer.apple.com/documentation/AppKit/ui-element-colors
type NSColor struct {
	objectivec.Object
}

// NSColorFromID constructs a [NSColor] from an objc.ID.
//
// An object that stores color data and sometimes opacity (alpha value).
func NSColorFromID(id objc.ID) NSColor {
	return NSColor{objectivec.Object{ID: id}}
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
	GetCyanMagentaYellowBlackAlpha(cyan unsafe.Pointer, magenta unsafe.Pointer, yellow unsafe.Pointer, black unsafe.Pointer, alpha unsafe.Pointer)
	// Returns the color object’s HSB component and opacity values in the respective arguments.
	GetHueSaturationBrightnessAlpha(hue unsafe.Pointer, saturation unsafe.Pointer, brightness unsafe.Pointer, alpha unsafe.Pointer)
	// Returns the color object’s RGB component and opacity values in the respective arguments.
	GetRedGreenBlueAlpha(red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, alpha unsafe.Pointer)
	// Returns the grayscale and alpha values of the color.
	GetWhiteAlpha(white unsafe.Pointer, alpha unsafe.Pointer)
	// The number of components in the color.
	NumberOfComponents() int
	// Returns the components of the color as an array.
	GetComponents(components unsafe.Pointer)

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
	ColorNameComponent() NSColorName
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
// [IgnoresAlpha] returns true.
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
func NewColorNamedBundle(name string, bundle foundation.NSBundle) NSColor {
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
func NewColorWithCIColor(color coreimage.CIColor) NSColor {
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(hue:saturation:brightness:alpha:)
//
// [UIColor]: https://developer.apple.com/documentation/UIKit/UIColor
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
// than [NSPasteboardReadingAsData], the `propertyList` may be any other
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
// [NSPasteboardReadingAsKeyedArchive] reading option, then instances are
// initialized using [init(coder:)] instead of this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/init(pasteboardPropertyList:ofType:)
//
// [init(coder:)]: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(red:green:blue:alpha:)
//
// [UIColor]: https://developer.apple.com/documentation/UIKit/UIColor
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(white:alpha:)
//
// [UIColor]: https://developer.apple.com/documentation/UIKit/UIColor
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
// [NSColorSystemEffectPressed] for the `systemEffect` parameter yields the
// color to use when you want your view to appear as if it had been pressed.
// This method takes the current appearance into account, returning an
// appropriately modified color for both light and dark appearances.
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
// The highlight color is provided by the [HighlightColor] property. Call this
// method when you want to brighten the current color for use in highlights.
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
// The shadow color is provided by the [ShadowColor] property. Call this
// method when you want to darken the current color for use in shadows.
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/getCyan(_:magenta:yellow:black:alpha:)
//
// [deviceCMYK]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceCMYK
func (c NSColor) GetCyanMagentaYellowBlackAlpha(cyan unsafe.Pointer, magenta unsafe.Pointer, yellow unsafe.Pointer, black unsafe.Pointer, alpha unsafe.Pointer) {
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/getHue(_:saturation:brightness:alpha:)
//
// [calibratedRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedRGB
// [deviceRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceRGB
func (c NSColor) GetHueSaturationBrightnessAlpha(hue unsafe.Pointer, saturation unsafe.Pointer, brightness unsafe.Pointer, alpha unsafe.Pointer) {
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/getRed(_:green:blue:alpha:)
//
// [calibratedRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedRGB
// [deviceRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceRGB
func (c NSColor) GetRedGreenBlueAlpha(red unsafe.Pointer, green unsafe.Pointer, blue unsafe.Pointer, alpha unsafe.Pointer) {
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/getWhite(_:alpha:)
//
// [NSCalibratedBlackColorSpace]: https://developer.apple.com/documentation/AppKit/NSCalibratedBlackColorSpace
// [NSDeviceBlackColorSpace]: https://developer.apple.com/documentation/AppKit/NSDeviceBlackColorSpace
// [calibratedWhite]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedWhite
// [deviceWhite]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceWhite
func (c NSColor) GetWhiteAlpha(white unsafe.Pointer, alpha unsafe.Pointer) {
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
func (c NSColor) GetComponents(components unsafe.Pointer) {
	objc.Send[objc.ID](c.ID, objc.Sel("getComponents:"), components)
}

// Returns a version of the color object that is compatible with the specified
// color type.
//
// type: The type of color object that you want. For example, if you want a color
// object containing RGB components, specify [NSColorTypeComponentBased].
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
// object of type [NSColorTypeComponentBased]. For some types of colors,
// conversions to a compatible color type may be possible.
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
// than [NSPasteboardReadingAsData], the `propertyList` may be any other
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
// [NSPasteboardReadingAsKeyedArchive] reading option, then instances are
// initialized using [init(coder:)] instead of this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/init(pasteboardPropertyList:ofType:)
//
// [init(coder:)]: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
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
// -writingOptionsForType:pasteboard: and return [NSPasteboardWritingPromised]
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
// # Discussion
//
// Do not perform other pasteboard operations in the method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardWriting/writingOptions(forType:pasteboard:)
//
// [Pasteboard Writing Options]: https://developer.apple.com/documentation/AppKit/pasteboard-writing-options
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
// provider with [CurrentDrawingAppearance]. The provider can use the
// appearance to return another color to use for drawing. For example, if you
// create a color matching [SystemYellowColor] for Light appearance, you’ll
// get a color matching [SystemYellowColor] for Dark appearance when using
// Dark Mode. As often as possible, use the given appearance.
//
// The `colorName` is equal to the identity of the color and should be
// universally unique; if `nil`, the system generates a unique name.
//
// The name and the Light (aqua) appearance are encoded as part of this color.
// If a color is already registered with the same name as the decoded color,
// the decoded color joins the other color and becomes dynamic again.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/init(name:dynamicProvider:)
func (_NSColorClass NSColorClass) ColorWithNameDynamicProvider(colorName string, dynamicProvider AppearanceHandler) NSColor {
	_block1, _ := NewAppearanceBlock(dynamicProvider)
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
// # Discussion
//
// Do not perform other pasteboard operations in this method implementation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPasteboardReading/readingOptions(forType:pasteboard:)
//
// [NSPasteboard.ReadingOptions]: https://developer.apple.com/documentation/AppKit/NSPasteboard/ReadingOptions
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/whiteComponent
//
// [NSCalibratedBlackColorSpace]: https://developer.apple.com/documentation/AppKit/NSCalibratedBlackColorSpace
// [NSDeviceBlackColorSpace]: https://developer.apple.com/documentation/AppKit/NSDeviceBlackColorSpace
// [calibratedWhite]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedWhite
// [deviceWhite]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceWhite
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/redComponent
//
// [calibratedRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedRGB
// [deviceRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceRGB
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/greenComponent
//
// [calibratedRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedRGB
// [deviceRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceRGB
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/blueComponent
//
// [calibratedRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedRGB
// [deviceRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceRGB
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/cyanComponent
//
// [deviceCMYK]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceCMYK
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/magentaComponent
//
// [deviceCMYK]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceCMYK
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/yellowComponent
//
// [deviceCMYK]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceCMYK
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/blackComponent
//
// [deviceCMYK]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceCMYK
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/hueComponent
//
// [calibratedRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedRGB
// [deviceRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceRGB
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/saturationComponent
//
// [calibratedRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedRGB
// [deviceRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceRGB
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/brightnessComponent
//
// [calibratedRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedRGB
// [deviceRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceRGB
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/catalogNameComponent
//
// [named]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/named
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/localizedCatalogNameComponent
//
// [named]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/named
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/colorNameComponent
//
// [named]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/named
func (c NSColor) ColorNameComponent() NSColorName {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("colorNameComponent"))
	return NSColorName(foundation.NSStringFromID(rv).String())
}

// The localized version of the color name.
//
// # Discussion
//
// Use the value in this property when displaying the color name in your user
// interface. Access this property only for colors in the [named] color space.
// Accessing it for other color types raises an exception.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/localizedColorNameComponent
//
// [named]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/named
func (c NSColor) LocalizedColorNameComponent() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("localizedColorNameComponent"))
	return foundation.NSStringFromID(rv).String()
}

// The type of the color object.
//
// # Discussion
//
// A color object’s type determines which of its methods and properties you
// may access. For example, if the type is [NSColorTypePattern], you may
// safely access the [PatternImage] property.
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
// See: https://developer.apple.com/documentation/AppKit/NSColor/colorSpace
//
// [NSCalibratedBlackColorSpace]: https://developer.apple.com/documentation/AppKit/NSCalibratedBlackColorSpace
// [NSDeviceBlackColorSpace]: https://developer.apple.com/documentation/AppKit/NSDeviceBlackColorSpace
// [calibratedRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedRGB
// [calibratedWhite]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/calibratedWhite
// [custom]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/custom
// [deviceCMYK]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceCMYK
// [deviceRGB]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceRGB
// [deviceWhite]: https://developer.apple.com/documentation/AppKit/NSColorSpaceName/deviceWhite
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

// Returns a localized description of the color for use in accessibility
// attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityColor/accessibilityName
func (c NSColor) AccessibilityName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("accessibilityName"))
	return foundation.NSStringFromID(rv).String()
}

// The pattern image used to paint the target area.
//
// # Discussion
//
// This property contains the image (instead of the solid color) to use during
// drawing. Pattern images are tiled as needed to fill the rendered area.
//
// Do not access this property unless you created the color object using a
// pattern image. Accessing this property for colors created using other types
// of color information raises an exception.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/patternImage
func (c NSColor) PatternImage() INSImage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("patternImage"))
	return NSImageFromID(objc.ID(rv))
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
// true if the app doesn’t support alpha; otherwise, false.
//
// # Discussion
//
// The system consults this global value when the app imports alpha (for
// instance, through color dragging). By default this property is false;
// meaning the system supports the alpha component for colors globally. To
// ignore alpha for an app, invoke the `setIgnoresAlpha` method with a
// parameter of true. This value also determines whether the color panel has
// an opacity slider.
//
// This method provides a global approach for removing alpha, which might not
// always be appropriate. Apps that need to disable alpha can use more
// fine-grained APIs for individual controls, such as [ShowsAlpha] and
// [SupportsAlpha].
//
// In macOS 13 and earlier, the default value is true. This property is
// deprecated as of macOS 14.
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

// The color to use for text in a selected control.
//
// # Discussion
//
// This color is the table and list view equivalent to the
// [SelectedControlTextColor] color, which is used for controls in other
// views.
//
// For more information, see [NSColor].
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/alternateSelectedControlTextColor
func (_NSColorClass NSColorClass) AlternateSelectedControlTextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("alternateSelectedControlTextColor"))
	return NSColorFromID(objc.ID(rv))
}

// The colors to use for alternating content, typically found in table views
// and collection views.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/alternatingContentBackgroundColors
func (_NSColorClass NSColorClass) AlternatingContentBackgroundColors() []NSColor {
	rv := objc.Send[[]objc.ID](objc.ID(_NSColorClass.class), objc.Sel("alternatingContentBackgroundColors"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSColor {
		return NSColorFromID(id)
	})
}

// Returns a color object whose grayscale value is `0.0` and whose alpha value
// is `1.0`.
//
// # Return Value
//
// The [NSColor] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/black
func (_NSColorClass NSColorClass) BlackColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("blackColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object whose RGB value is `0.0`, `0.0`, `1.0` and whose
// alpha value is `1.0`.
//
// # Return Value
//
// The [NSColor] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/blue
func (_NSColorClass NSColorClass) BlueColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("blueColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object whose RGB value is `0.6`, `0.4`, `0.2` and whose
// alpha value is `1.0`.
//
// # Return Value
//
// The [NSColor] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/brown
func (_NSColorClass NSColorClass) BrownColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("brownColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object whose grayscale and alpha values are both `0.0`.
//
// # Return Value
//
// The [NSColor] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/clear
func (_NSColorClass NSColorClass) ClearColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("clearColor"))
	return NSColorFromID(objc.ID(rv))
}

// The user’s current accent color preference.
//
// # Discussion
//
// People set the accent color in the Appearance pane of System Settings.
// Don’t make assumptions about the color space associated with this color.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/controlAccentColor
func (_NSColorClass NSColorClass) ControlAccentColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("controlAccentColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for the background of large controls, such as scroll views
// or table views.
//
// # Return Value
//
// Do not use this color for drawing. Instead, use an [NSVisualEffectView]
// with the appropriate background material.
//
// # Discussion
//
// When applied to an [NSBox] object, this color supports Desktop Tinting in
// Dark Mode. With Desktop Tinting, the system modifies this color dynamically
// by incorporating some of the color from the underlying desktop image. The
// system does not apply this dynamic tinting effect to other types of views.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/controlBackgroundColor
func (_NSColorClass NSColorClass) ControlBackgroundColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("controlBackgroundColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for the flat surfaces of a control.
//
// # Return Value
//
// The system color used for the flat surfaces of a control. By default, the
// control color is a pattern color that will draw the ruled lines for the
// window background, which is the same as returned by
// [WindowBackgroundColor].
//
// # Discussion
//
// If you use [ControlColor] assuming that it is a solid, you may have an
// incorrect appearance. You should use [LightGrayColor] in its place.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/controlColor
func (_NSColorClass NSColorClass) ControlColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("controlColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for text on enabled controls.
//
// # Return Value
//
// The color used for text on enabled controls. For more information, see
// [NSColor].
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/controlTextColor
func (_NSColorClass NSColorClass) ControlTextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("controlTextColor"))
	return NSColorFromID(objc.ID(rv))
}

// The current system control tint color.
//
// # Return Value
//
// The current system control tint.
//
// # Discussion
//
// An application can register for the
// [CurrentControlTintDidChangeNotification] notification to be notified of
// changes to the system control tint.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/currentControlTint
func (_NSColorClass NSColorClass) CurrentControlTint() NSControlTint {
	rv := objc.Send[NSControlTint](objc.ID(_NSColorClass.class), objc.Sel("currentControlTint"))
	return NSControlTint(rv)
}

// Returns a color object whose RGB value is `0.0`, `1.0`, `1.0` and whose
// alpha value is `1.0`.
//
// # Return Value
//
// The [NSColor] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/cyan
func (_NSColorClass NSColorClass) CyanColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("cyanColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object whose grayscale value is `1/3` and whose alpha value
// is `1.0`.
//
// # Return Value
//
// The [NSColor] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/darkGray
func (_NSColorClass NSColorClass) DarkGrayColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("darkGrayColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for text on disabled controls.
//
// # Return Value
//
// The color used for text on disabled controls. For more information, see
// [NSColor].
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/disabledControlTextColor
func (_NSColorClass NSColorClass) DisabledControlTextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("disabledControlTextColor"))
	return NSColorFromID(objc.ID(rv))
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

// The highlight color to use for the bubble that shows inline search result
// values.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/findHighlightColor
func (_NSColorClass NSColorClass) FindHighlightColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("findHighlightColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object whose grayscale value is `0.5` and whose alpha value
// is `1.0`.
//
// # Return Value
//
// The [NSColor] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/gray
func (_NSColorClass NSColorClass) GrayColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("grayColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object whose RGB value is `0.0`, `1.0`, `0.0` and whose
// alpha value is `1.0`.
//
// # Return Value
//
// The [NSColor] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/green
func (_NSColorClass NSColorClass) GreenColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("greenColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for the optional gridlines, such as those in a table view.
//
// # Return Value
//
// The system color used for gridlines. For more information, see [NSColor].
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/gridColor
func (_NSColorClass NSColorClass) GridColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("gridColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for text in header cells in table views and outline views.
//
// # Return Value
//
// The system color used for text in header cells in table and outline views.
// For more information, see [NSColor].
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/headerTextColor
func (_NSColorClass NSColorClass) HeaderTextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("headerTextColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use as a virtual light source on the screen.
//
// # Return Value
//
// The system color for the virtual light source on the screen.
//
// # Discussion
//
// This method is invoked by the [HighlightWithLevel] method. For more
// information, see [NSColor].
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/highlightColor
func (_NSColorClass NSColorClass) HighlightColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("highlightColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for the keyboard focus ring around controls.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/keyboardFocusIndicatorColor
func (_NSColorClass NSColorClass) KeyboardFocusIndicatorColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("keyboardFocusIndicatorColor"))
	return NSColorFromID(objc.ID(rv))
}

// The primary color to use for text labels.
//
// # Discussion
//
// Use this color in the most important text labels of your user interface.
// You can also use it for other types of primary app content.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/labelColor
func (_NSColorClass NSColorClass) LabelColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("labelColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object whose grayscale value is `2/3` and whose alpha value
// is `1.0`.
//
// # Return Value
//
// The [NSColor] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/lightGray
func (_NSColorClass NSColorClass) LightGrayColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("lightGrayColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for links.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/linkColor
func (_NSColorClass NSColorClass) LinkColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("linkColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object whose RGB value is `1.0`, `0.0`, `1.0` and whose
// alpha value is `1.0`.
//
// # Return Value
//
// The [NSColor] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/magenta
func (_NSColorClass NSColorClass) MagentaColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("magentaColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object whose RGB value is `1.0`, `0.5`, `0.0` and whose
// alpha value is `1.0`.
//
// # Return Value
//
// The [NSColor] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/orange
func (_NSColorClass NSColorClass) OrangeColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("orangeColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for placeholder text in controls or text views.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/placeholderTextColor
func (_NSColorClass NSColorClass) PlaceholderTextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("placeholderTextColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object whose RGB value is `0.5`, `0.0`, `0.5` and whose
// alpha value is `1.0`.
//
// # Return Value
//
// The [NSColor] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/purple
func (_NSColorClass NSColorClass) PurpleColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("purpleColor"))
	return NSColorFromID(objc.ID(rv))
}

// The quaternary color to use for text labels and separators.
//
// # Discussion
//
// Use this color for the least important text in your interface and for
// separators between text items. For example, you would use this color for
// secondary text that is disabled. You can also use it for other types of
// quaternary app content.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/quaternaryLabelColor
func (_NSColorClass NSColorClass) QuaternaryLabelColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("quaternaryLabelColor"))
	return NSColorFromID(objc.ID(rv))
}

// A color appropriate for filling large areas, such as a group box or tab
// pane.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/quaternarySystemFill
func (_NSColorClass NSColorClass) QuaternarySystemFillColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("quaternarySystemFillColor"))
	return NSColorFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AppKit/NSColor/quinaryLabel
func (_NSColorClass NSColorClass) QuinaryLabelColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("quinaryLabelColor"))
	return NSColorFromID(objc.ID(rv))
}

// A color appropriate for filling large areas that require subtle emphasis,
// such as content of a form.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/quinarySystemFill
func (_NSColorClass NSColorClass) QuinarySystemFillColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("quinarySystemFillColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object whose RGB value is `1.0`, `0.0`, `0.0` and whose
// alpha value is `1.0`.
//
// # Return Value
//
// The [NSColor] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/red
func (_NSColorClass NSColorClass) RedColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("redColor"))
	return NSColorFromID(objc.ID(rv))
}

// The patterned color to use for the background of a scrubber control.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/scrubberTexturedBackground
func (_NSColorClass NSColorClass) ScrubberTexturedBackgroundColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("scrubberTexturedBackgroundColor"))
	return NSColorFromID(objc.ID(rv))
}

// The secondary color to use for text labels.
//
// # Discussion
//
// Use this color in text fields that contain less important text in your user
// interface. For example, you might use this in labels that display subheads
// or additional information. You can also use it for other types of secondary
// app content.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/secondaryLabelColor
func (_NSColorClass NSColorClass) SecondaryLabelColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("secondaryLabelColor"))
	return NSColorFromID(objc.ID(rv))
}

// A color appropriate for filling small-size shapes, such as the backing of a
// progress indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/secondarySystemFill
func (_NSColorClass NSColorClass) SecondarySystemFillColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("secondarySystemFillColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for the background of selected and emphasized content.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/selectedContentBackgroundColor
func (_NSColorClass NSColorClass) SelectedContentBackgroundColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("selectedContentBackgroundColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for the face of a selected control—that is, a control
// that has been clicked or is being dragged.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/selectedControlColor
func (_NSColorClass NSColorClass) SelectedControlColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("selectedControlColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for text in a selected control—that is, a control being
// clicked or dragged.
//
// # Discussion
//
// For more information, see [NSColor].
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/selectedControlTextColor
func (_NSColorClass NSColorClass) SelectedControlTextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("selectedControlTextColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for the text in menu items.
//
// # Return Value
//
// The system color used for text in selected menu items. For more
// information, see [NSColor].
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/selectedMenuItemTextColor
func (_NSColorClass NSColorClass) SelectedMenuItemTextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("selectedMenuItemTextColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for the background of selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/selectedTextBackgroundColor
func (_NSColorClass NSColorClass) SelectedTextBackgroundColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("selectedTextBackgroundColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/selectedTextColor
func (_NSColorClass NSColorClass) SelectedTextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("selectedTextColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for separators between different sections of content.
//
// # Discussion
//
// Do not use this color for split view dividers or window chrome dividers.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/separatorColor
func (_NSColorClass NSColorClass) SeparatorColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("separatorColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for virtual shadows cast by raised objects on the screen.
//
// # Return Value
//
// The system color for the virtual shadows case by raised objects on the
// screen.
//
// # Discussion
//
// This method is invoked by [ShadowWithLevel]. For more information, see
// [NSColor].
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/shadowColor
func (_NSColorClass NSColorClass) ShadowColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("shadowColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object for blue that automatically adapts to vibrancy and
// accessibility settings.
//
// # Discussion
//
// Use this color when you want a tint color that looks great on both light
// and dark backgrounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/systemBlue
func (_NSColorClass NSColorClass) SystemBlueColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemBlueColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object for brown that automatically adapts to vibrancy and
// accessibility settings.
//
// # Discussion
//
// Use this color when you want a tint color that looks great on both light
// and dark backgrounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/systemBrown
func (_NSColorClass NSColorClass) SystemBrownColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemBrownColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object for cyan that automatically adapts to vibrancy and
// accessibility settings.
//
// # Discussion
//
// Use this color when you want a tint color that looks great on both light
// and dark backgrounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/systemCyan
func (_NSColorClass NSColorClass) SystemCyanColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemCyanColor"))
	return NSColorFromID(objc.ID(rv))
}

// A color appropriate for filling thin shapes, such as the track of a slider.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/systemFill
func (_NSColorClass NSColorClass) SystemFillColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemFillColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object for gray that automatically adapts to vibrancy and
// accessibility settings.
//
// # Discussion
//
// Use this color when you want a tint color that looks great on both light
// and dark backgrounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/systemGray
func (_NSColorClass NSColorClass) SystemGrayColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemGrayColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object for green that automatically adapts to vibrancy and
// accessibility settings.
//
// # Discussion
//
// Use this color when you want a tint color that looks great on both light
// and dark backgrounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/systemGreen
func (_NSColorClass NSColorClass) SystemGreenColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemGreenColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object for indigo that automatically adapts to vibrancy and
// accessibility settings.
//
// # Discussion
//
// Use this color when you want a tint color that looks great on both light
// and dark backgrounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/systemIndigo
func (_NSColorClass NSColorClass) SystemIndigoColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemIndigoColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object for mint that automatically adapts to vibrancy and
// accessibility settings.
//
// # Discussion
//
// Use this color when you want a tint color that looks great on both light
// and dark backgrounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/systemMint
func (_NSColorClass NSColorClass) SystemMintColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemMintColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object for orange that automatically adapts to vibrancy and
// accessibility settings.
//
// # Discussion
//
// Use this color when you want a tint color that looks great on both light
// and dark backgrounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/systemOrange
func (_NSColorClass NSColorClass) SystemOrangeColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemOrangeColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object for pink that automatically adapts to vibrancy and
// accessibility settings.
//
// # Discussion
//
// Use this color when you want a tint color that looks great on both light
// and dark backgrounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/systemPink
func (_NSColorClass NSColorClass) SystemPinkColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemPinkColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object for purple that automatically adapts to vibrancy and
// accessibility settings.
//
// # Discussion
//
// Use this color when you want a tint color that looks great on both light
// and dark backgrounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/systemPurple
func (_NSColorClass NSColorClass) SystemPurpleColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemPurpleColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object for red that automatically adapts to vibrancy and
// accessibility settings.
//
// # Discussion
//
// Use this color when you want a tint color that looks great on both light
// and dark backgrounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/systemRed
func (_NSColorClass NSColorClass) SystemRedColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemRedColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object for teal that automatically adapts to vibrancy and
// accessibility settings.
//
// # Discussion
//
// Use this color when you want a tint color that looks great on both light
// and dark backgrounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/systemTeal
func (_NSColorClass NSColorClass) SystemTealColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemTealColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object for yellow that automatically adapts to vibrancy and
// accessibility settings.
//
// # Discussion
//
// Use this color when you want a tint color that looks great on both light
// and dark backgrounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/systemYellow
func (_NSColorClass NSColorClass) SystemYellowColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("systemYellowColor"))
	return NSColorFromID(objc.ID(rv))
}

// The tertiary color to use for text labels.
//
// # Discussion
//
// Use this color for disabled text and for other less important text in your
// interface. You can also use it for other types of tertiary app content.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/tertiaryLabelColor
func (_NSColorClass NSColorClass) TertiaryLabelColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("tertiaryLabelColor"))
	return NSColorFromID(objc.ID(rv))
}

// A color appropriate for filling medium-size shapes, such as the backing of
// a switch.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/tertiarySystemFill
func (_NSColorClass NSColorClass) TertiarySystemFillColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("tertiarySystemFillColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for the background area behind text.
//
// # Discussion
//
// When text is selected, its background color changes to the return value of
// [SelectedTextBackgroundColor].
//
// When applied to an [NSBox] object, this color supports Desktop Tinting in
// Dark Mode. With Desktop Tinting, the system modifies this color dynamically
// by incorporating some of the color from the underlying desktop image. The
// system does not apply this dynamic tinting effect to other types of views.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/textBackgroundColor
func (_NSColorClass NSColorClass) TextBackgroundColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("textBackgroundColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for text.
//
// # Discussion
//
// When text is selected, its color changes to the return value of
// [SelectedTextColor].
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/textColor
func (_NSColorClass NSColorClass) TextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("textColor"))
	return NSColorFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AppKit/NSColor/textInsertionPointColor
func (_NSColorClass NSColorClass) TextInsertionPointColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("textInsertionPointColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use in the area beneath your window’s views.
//
// # Return Value
//
// Use this color to fill the backdrop underneath your app’s main content.
//
// # Discussion
//
// When applied to an [NSBox] object, this color supports Desktop Tinting in
// Dark Mode. With Desktop Tinting, the system modifies this color dynamically
// by incorporating some of the color from the underlying desktop image. The
// system does not apply this dynamic tinting effect to other types of views.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/underPageBackgroundColor
func (_NSColorClass NSColorClass) UnderPageBackgroundColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("underPageBackgroundColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for selected and unemphasized content.
//
// # Discussion
//
// Use this color when the window containing the content is not key, or when
// the view containing the content does not have key focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/unemphasizedSelectedContentBackgroundColor
func (_NSColorClass NSColorClass) UnemphasizedSelectedContentBackgroundColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("unemphasizedSelectedContentBackgroundColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for the text background in an unemphasized context.
//
// # Discussion
//
// Use this color when the window containing the text is not key, or when the
// view containing the text does not have key focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/unemphasizedSelectedTextBackgroundColor
func (_NSColorClass NSColorClass) UnemphasizedSelectedTextBackgroundColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("unemphasizedSelectedTextBackgroundColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for selected text in an unemphasized context.
//
// # Discussion
//
// Use this color when the window containing the text is not key, or when the
// view containing the text does not have key focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/unemphasizedSelectedTextColor
func (_NSColorClass NSColorClass) UnemphasizedSelectedTextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("unemphasizedSelectedTextColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object whose grayscale and alpha values are both `1.0`.
//
// # Return Value
//
// The [NSColor] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/white
func (_NSColorClass NSColorClass) WhiteColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("whiteColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for the window background.
//
// # Return Value
//
// The window background color.
//
// # Discussion
//
// When applied to an [NSBox] object, this color supports Desktop Tinting in
// Dark Mode. With Desktop Tinting, the system modifies this color dynamically
// by incorporating some of the color from the underlying desktop image. The
// system does not apply this dynamic tinting effect to other types of views.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/windowBackgroundColor
func (_NSColorClass NSColorClass) WindowBackgroundColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("windowBackgroundColor"))
	return NSColorFromID(objc.ID(rv))
}

// The color to use for text in a window’s frame.
//
// # Return Value
//
// The color used for text in window frames. For more information, see
// [NSColor].
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/windowFrameTextColor
func (_NSColorClass NSColorClass) WindowFrameTextColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("windowFrameTextColor"))
	return NSColorFromID(objc.ID(rv))
}

// Returns a color object whose RGB value is `1.0`, `1.0`, `0.0` and whose
// alpha value is `1.0`.
//
// # Return Value
//
// The [NSColor] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSColor/yellow
func (_NSColorClass NSColorClass) YellowColor() NSColor {
	rv := objc.Send[objc.ID](objc.ID(_NSColorClass.class), objc.Sel("yellowColor"))
	return NSColorFromID(objc.ID(rv))
}

// Protocol methods for NSAccessibilityColor

// Protocol methods for NSPasteboardWriting

// ColorWithNameDynamicProviderSync is a synchronous wrapper around [NSColor.ColorWithNameDynamicProvider].
// It blocks until the completion handler fires or the context is cancelled.
func (cc NSColorClass) ColorWithNameDynamicProviderSync(ctx context.Context, colorName NSColorName) (*NSAppearance, error) {
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
