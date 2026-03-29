// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSImageSymbolConfiguration] class.
var (
	_NSImageSymbolConfigurationClass     NSImageSymbolConfigurationClass
	_NSImageSymbolConfigurationClassOnce sync.Once
)

func getNSImageSymbolConfigurationClass() NSImageSymbolConfigurationClass {
	_NSImageSymbolConfigurationClassOnce.Do(func() {
		_NSImageSymbolConfigurationClass = NSImageSymbolConfigurationClass{class: objc.GetClass("NSImageSymbolConfiguration")}
	})
	return _NSImageSymbolConfigurationClass
}

// GetNSImageSymbolConfigurationClass returns the class object for NSImageSymbolConfiguration.
func GetNSImageSymbolConfigurationClass() NSImageSymbolConfigurationClass {
	return getNSImageSymbolConfigurationClass()
}

type NSImageSymbolConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSImageSymbolConfigurationClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSImageSymbolConfigurationClass) Alloc() NSImageSymbolConfiguration {
	rv := objc.Send[NSImageSymbolConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that contains the specific font, style, and weight attributes to
// apply to a symbol image.
//
// # Overview
// 
// Symbol image configuration objects include details such as the point size,
// scale, text style, and weight to apply to your symbol image. The system
// uses these details to determine which variant of the image to use and how
// to scale or style the image.
// 
// [NSImageSymbolConfiguration] objects are immutable after you create them.
// If you use the [NSImageSymbolConfiguration.ConfigurationByApplyingConfiguration] method on the object,
// the new image attributes replace any previous attributes you supplied.
// After creating a symbol configuration object, assign it to the
// [NSImageSymbolConfiguration.SymbolConfiguration] property of the [NSImageView] object you use to
// display the image. If you draw the image directly, use the
// [ImageWithSymbolConfiguration] method to create a new image that contains
// the new attributes.
// 
// For design guidance, see [Human Interface Guidelines].
//
// [Human Interface Guidelines]: https://developer.apple.com/design/human-interface-guidelines/sf-symbols/overview/
//
// # Applying a Configuration
//
//   - [NSImageSymbolConfiguration.ConfigurationByApplyingConfiguration]: Creates a configuration object by applying the values from the configuration you specify.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class
type NSImageSymbolConfiguration struct {
	objectivec.Object
}

// NSImageSymbolConfigurationFromID constructs a [NSImageSymbolConfiguration] from an objc.ID.
//
// An object that contains the specific font, style, and weight attributes to
// apply to a symbol image.
func NSImageSymbolConfigurationFromID(id objc.ID) NSImageSymbolConfiguration {
	return NSImageSymbolConfiguration{objectivec.Object{ID: id}}
}
// NOTE: NSImageSymbolConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSImageSymbolConfiguration] class.
//
// # Applying a Configuration
//
//   - [INSImageSymbolConfiguration.ConfigurationByApplyingConfiguration]: Creates a configuration object by applying the values from the configuration you specify.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class
type INSImageSymbolConfiguration interface {
	objectivec.IObject

	// Topic: Applying a Configuration

	// Creates a configuration object by applying the values from the configuration you specify.
	ConfigurationByApplyingConfiguration(configuration INSImageSymbolConfiguration) INSImageSymbolConfiguration

	SymbolConfiguration() INSImageSymbolConfiguration
	SetSymbolConfiguration(value INSImageSymbolConfiguration)
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (i NSImageSymbolConfiguration) Init() NSImageSymbolConfiguration {
	rv := objc.Send[NSImageSymbolConfiguration](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i NSImageSymbolConfiguration) Autorelease() NSImageSymbolConfiguration {
	rv := objc.Send[NSImageSymbolConfiguration](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSImageSymbolConfiguration creates a new NSImageSymbolConfiguration instance.
func NewNSImageSymbolConfiguration() NSImageSymbolConfiguration {
	class := getNSImageSymbolConfigurationClass()
	rv := objc.Send[NSImageSymbolConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Create a configuration with a specific color rendering mode.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(colorRenderingMode:)
func NewImageSymbolConfigurationWithColorRenderingMode(mode NSImageSymbolColorRenderingMode) NSImageSymbolConfiguration {
	rv := objc.Send[objc.ID](objc.ID(getNSImageSymbolConfigurationClass().class), objc.Sel("configurationWithColorRenderingMode:"), mode)
	return NSImageSymbolConfigurationFromID(rv)
}

// Creates a hierarchical color configuration using the color you specify.
//
// hierarchicalColor: The primary color for the symbol.
//
// # Return Value
// 
// A new configuration object that creates a palette from a primary color.
//
// # Discussion
// 
// This method creates a color scheme based on a single color. The system
// reduces the intensity of the base color to create the secondary and
// tertiary colors.
// 
// When combining this with another configuration, the last configuration
// overrides existing values.
// 
// If the symbol doesn’t have a palette variant, this color configuration
// doesn’t have an effect, so the symbol uses the monochrome (templated)
// symbol.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(hierarchicalColor:)
func NewImageSymbolConfigurationWithHierarchicalColor(hierarchicalColor INSColor) NSImageSymbolConfiguration {
	rv := objc.Send[objc.ID](objc.ID(getNSImageSymbolConfigurationClass().class), objc.Sel("configurationWithHierarchicalColor:"), hierarchicalColor)
	return NSImageSymbolConfigurationFromID(rv)
}

// Creates a color configuration by specifying a palette of colors.
//
// paletteColors: The colors to apply to the symbol.
//
// # Return Value
// 
// A new configuration object that prefers its palette variant.
//
// # Discussion
// 
// The system applies the colors sequentially per layer — the first color
// for the first layer, and the second color for the second layer. This is
// independent of the hierarchy level of the layer.
// 
// When you combine this with another configuration to create a palette, the
// last configuration overrides any existing color configuration.
// 
// If the symbol doesn’t have a palette variant, this color configuration
// doesn’t have an effect, so the symbol uses the monochrome (templated)
// symbol.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(paletteColors:)
func NewImageSymbolConfigurationWithPaletteColors(paletteColors []NSColor) NSImageSymbolConfiguration {
	rv := objc.Send[objc.ID](objc.ID(getNSImageSymbolConfigurationClass().class), objc.Sel("configurationWithPaletteColors:"), objectivec.IObjectSliceToNSArray(paletteColors))
	return NSImageSymbolConfigurationFromID(rv)
}

// Creates a symbol configuration with the specified point size and font
// weight.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(pointSize:weight:)
func NewImageSymbolConfigurationWithPointSizeWeight(pointSize float64, weight NSFontWeight) NSImageSymbolConfiguration {
	rv := objc.Send[objc.ID](objc.ID(getNSImageSymbolConfigurationClass().class), objc.Sel("configurationWithPointSize:weight:"), pointSize, weight)
	return NSImageSymbolConfigurationFromID(rv)
}

// Creates a symbol configuration with the specified point size, font weight,
// and symbol scale.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(pointSize:weight:scale:)
func NewImageSymbolConfigurationWithPointSizeWeightScale(pointSize float64, weight NSFontWeight, scale NSImageSymbolScale) NSImageSymbolConfiguration {
	rv := objc.Send[objc.ID](objc.ID(getNSImageSymbolConfigurationClass().class), objc.Sel("configurationWithPointSize:weight:scale:"), pointSize, weight, scale)
	return NSImageSymbolConfigurationFromID(rv)
}

// Creates a symbol configuration using the scale you specify.
//
// scale: The symbol scale.
//
// # Return Value
// 
// A new configuration object.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(scale:)
func NewImageSymbolConfigurationWithScale(scale NSImageSymbolScale) NSImageSymbolConfiguration {
	rv := objc.Send[objc.ID](objc.ID(getNSImageSymbolConfigurationClass().class), objc.Sel("configurationWithScale:"), scale)
	return NSImageSymbolConfigurationFromID(rv)
}

// Creates a symbol configuration with the specified text style.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(textStyle:)
func NewImageSymbolConfigurationWithTextStyle(style NSFontTextStyle) NSImageSymbolConfiguration {
	rv := objc.Send[objc.ID](objc.ID(getNSImageSymbolConfigurationClass().class), objc.Sel("configurationWithTextStyle:"), objc.String(string(style)))
	return NSImageSymbolConfigurationFromID(rv)
}

// Creates a symbol configuration with the specified text style and symbol
// scale.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(textStyle:scale:)
func NewImageSymbolConfigurationWithTextStyleScale(style NSFontTextStyle, scale NSImageSymbolScale) NSImageSymbolConfiguration {
	rv := objc.Send[objc.ID](objc.ID(getNSImageSymbolConfigurationClass().class), objc.Sel("configurationWithTextStyle:scale:"), objc.String(string(style)), scale)
	return NSImageSymbolConfigurationFromID(rv)
}

// Create a configuration with a specified variable value mode.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/init(variableValueMode:)
func NewImageSymbolConfigurationWithVariableValueMode(variableValueMode NSImageSymbolVariableValueMode) NSImageSymbolConfiguration {
	rv := objc.Send[objc.ID](objc.ID(getNSImageSymbolConfigurationClass().class), objc.Sel("configurationWithVariableValueMode:"), variableValueMode)
	return NSImageSymbolConfigurationFromID(rv)
}

// Creates a configuration object by applying the values from the
// configuration you specify.
//
// configuration: The configuration details to apply.
//
// # Return Value
// 
// A new configuration object that prioritizes the values from the
// configuration you specify.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/applying(_:)
func (i NSImageSymbolConfiguration) ConfigurationByApplyingConfiguration(configuration INSImageSymbolConfiguration) INSImageSymbolConfiguration {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("configurationByApplyingConfiguration:"), configuration)
	return NSImageSymbolConfigurationFromID(rv)
}
func (i NSImageSymbolConfiguration) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](i.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates a configuration that specifies that the symbol should prefer its
// multicolor variant if one exists.
//
// # Return Value
// 
// A new configuration object that prefers its multicolor variant.
//
// # Discussion
// 
// You can combine this configuration with one of the palette-based
// configurations. In that case, the symbol uses the multicolor variant if one
// exists; otherwise the symbol uses the palette version.
// 
// If the symbol supports neither, the symbol uses the monochrome (templated)
// symbol.
//
// See: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/preferringMulticolor()
func (_NSImageSymbolConfigurationClass NSImageSymbolConfigurationClass) ConfigurationPreferringMulticolor() NSImageSymbolConfiguration {
	rv := objc.Send[objc.ID](objc.ID(_NSImageSymbolConfigurationClass.class), objc.Sel("configurationPreferringMulticolor"))
	return NSImageSymbolConfigurationFromID(rv)
}
// See: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/preferringHierarchical()
func (_NSImageSymbolConfigurationClass NSImageSymbolConfigurationClass) ConfigurationPreferringHierarchical() NSImageSymbolConfiguration {
	rv := objc.Send[objc.ID](objc.ID(_NSImageSymbolConfigurationClass.class), objc.Sel("configurationPreferringHierarchical"))
	return NSImageSymbolConfigurationFromID(rv)
}
// See: https://developer.apple.com/documentation/AppKit/NSImage/SymbolConfiguration-swift.class/preferringMonochrome()
func (_NSImageSymbolConfigurationClass NSImageSymbolConfigurationClass) ConfigurationPreferringMonochrome() NSImageSymbolConfiguration {
	rv := objc.Send[objc.ID](objc.ID(_NSImageSymbolConfigurationClass.class), objc.Sel("configurationPreferringMonochrome"))
	return NSImageSymbolConfigurationFromID(rv)
}

// See: https://developer.apple.com/documentation/appkit/nsimageview/symbolconfiguration
func (i NSImageSymbolConfiguration) SymbolConfiguration() INSImageSymbolConfiguration {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("symbolConfiguration"))
	return NSImageSymbolConfigurationFromID(objc.ID(rv))
}
func (i NSImageSymbolConfiguration) SetSymbolConfiguration(value INSImageSymbolConfiguration) {
	objc.Send[struct{}](i.ID, objc.Sel("setSymbolConfiguration:"), value)
}

