// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTintConfiguration] class.
var (
	_NSTintConfigurationClass     NSTintConfigurationClass
	_NSTintConfigurationClassOnce sync.Once
)

func getNSTintConfigurationClass() NSTintConfigurationClass {
	_NSTintConfigurationClassOnce.Do(func() {
		_NSTintConfigurationClass = NSTintConfigurationClass{class: objc.GetClass("NSTintConfiguration")}
	})
	return _NSTintConfigurationClass
}

// GetNSTintConfigurationClass returns the class object for NSTintConfiguration.
func GetNSTintConfigurationClass() NSTintConfigurationClass {
	return getNSTintConfigurationClass()
}

type NSTintConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTintConfigurationClass) Alloc() NSTintConfiguration {
	rv := objc.Send[NSTintConfiguration](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An object that gives you the ability to choose from system-provided tinting
// behaviors.
//
// # Changing an App’s Appearance
//
//   - [NSTintConfiguration.AdaptsToUserAccentColor]: A Boolean value that indicates whether the tint configuration alters its effect based on the user’s preferred accent color choice.
//
// # Setting the Tint Color
//
//   - [NSTintConfiguration.BaseTintColor]: The color the system supplies when you create a tint configuration.
//   - [NSTintConfiguration.EquivalentContentTintColor]: A color object that matches the effective content tint.
//
// See: https://developer.apple.com/documentation/AppKit/NSTintConfiguration
type NSTintConfiguration struct {
	objectivec.Object
}

// NSTintConfigurationFromID constructs a [NSTintConfiguration] from an objc.ID.
//
// An object that gives you the ability to choose from system-provided tinting
// behaviors.
func NSTintConfigurationFromID(id objc.ID) NSTintConfiguration {
	return NSTintConfiguration{objectivec.Object{id}}
}
// NOTE: NSTintConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSTintConfiguration] class.
//
// # Changing an App’s Appearance
//
//   - [INSTintConfiguration.AdaptsToUserAccentColor]: A Boolean value that indicates whether the tint configuration alters its effect based on the user’s preferred accent color choice.
//
// # Setting the Tint Color
//
//   - [INSTintConfiguration.BaseTintColor]: The color the system supplies when you create a tint configuration.
//   - [INSTintConfiguration.EquivalentContentTintColor]: A color object that matches the effective content tint.
//
// See: https://developer.apple.com/documentation/AppKit/NSTintConfiguration
type INSTintConfiguration interface {
	objectivec.IObject

	// Topic: Changing an App’s Appearance

	// A Boolean value that indicates whether the tint configuration alters its effect based on the user’s preferred accent color choice.
	AdaptsToUserAccentColor() bool

	// Topic: Setting the Tint Color

	// The color the system supplies when you create a tint configuration.
	BaseTintColor() INSColor
	// A color object that matches the effective content tint.
	EquivalentContentTintColor() INSColor

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (t NSTintConfiguration) Init() NSTintConfiguration {
	rv := objc.Send[NSTintConfiguration](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTintConfiguration) Autorelease() NSTintConfiguration {
	rv := objc.Send[NSTintConfiguration](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTintConfiguration creates a new NSTintConfiguration instance.
func NewNSTintConfiguration() NSTintConfiguration {
	class := getNSTintConfigurationClass()
	rv := objc.Send[NSTintConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}






// Creates a new tint configuration using a specific color value.
//
// color: The color used regardless of the system accent color.
//
// See: https://developer.apple.com/documentation/AppKit/NSTintConfiguration/init(fixedColor:)
func NewTintConfigurationWithFixedColor(color INSColor) NSTintConfiguration {
	rv := objc.Send[objc.ID](objc.ID(getNSTintConfigurationClass().class), objc.Sel("tintConfigurationWithFixedColor:"), color)
	return NSTintConfigurationFromID(rv)
}


// Creates a new tint configuration for the system to use when the app’s
// preferred accent color is in use.
//
// color: The color used when the system accent color is [Multicolor].
//
// # Discussion
// 
// Use this tint configuration for custom colors designed to match
// app-specific accent colors, but doesn’t look appropriate matched with a
// user-selected color. The tint configuation only uses the preferred color
// when the system accent color is [Multicolor]. If the system accent color is
// any other color, the tint configuration defers to the accent color.
//
// See: https://developer.apple.com/documentation/AppKit/NSTintConfiguration/init(preferredColor:)
func NewTintConfigurationWithPreferredColor(color INSColor) NSTintConfiguration {
	rv := objc.Send[objc.ID](objc.ID(getNSTintConfigurationClass().class), objc.Sel("tintConfigurationWithPreferredColor:"), color)
	return NSTintConfigurationFromID(rv)
}






func (t NSTintConfiguration) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](t.ID, objc.Sel("encodeWithCoder:"), coder)
}












// A Boolean value that indicates whether the tint configuration alters its
// effect based on the user’s preferred accent color choice.
//
// # Discussion
// 
// When this property is [YES], the tint configuration alters it effect based
// on the user’s preferred accent color. Otherwise, the tint configuration
// produces a constant effect regardless of the accent color preference.
//
// See: https://developer.apple.com/documentation/AppKit/NSTintConfiguration/adaptsToUserAccentColor
func (t NSTintConfiguration) AdaptsToUserAccentColor() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("adaptsToUserAccentColor"))
	return rv
}



// The color the system supplies when you create a tint configuration.
//
// # Discussion
// 
// This property is `nil` if the tint configuration wasn’t created using an
// [NSColor] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSTintConfiguration/baseTintColor
func (t NSTintConfiguration) BaseTintColor() INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("baseTintColor"))
	return NSColorFromID(objc.ID(rv))
}



// A color object that matches the effective content tint.
//
// # Discussion
// 
// The value of this property is an [NSColor] that matches the represented
// content tint. This property is `nil` if the system can’t represent the
// content tint as an [NSColor].
//
// See: https://developer.apple.com/documentation/AppKit/NSTintConfiguration/equivalentContentTintColor
func (t NSTintConfiguration) EquivalentContentTintColor() INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("equivalentContentTintColor"))
	return NSColorFromID(objc.ID(rv))
}







// The system tints the content using the system default value for its
// context.
//
// See: https://developer.apple.com/documentation/AppKit/NSTintConfiguration/default
func (_NSTintConfigurationClass NSTintConfigurationClass) DefaultTintConfiguration() NSTintConfiguration {
	rv := objc.Send[objc.ID](objc.ID(_NSTintConfigurationClass.class), objc.Sel("defaultTintConfiguration"))
	return NSTintConfigurationFromID(objc.ID(rv))
}



// The content always displays in monochrome.
//
// # Discussion
// 
// Content marked as monochrome remains monochrome regardless of the system
// accent color.
//
// See: https://developer.apple.com/documentation/AppKit/NSTintConfiguration/monochrome
func (_NSTintConfigurationClass NSTintConfigurationClass) MonochromeTintConfiguration() NSTintConfiguration {
	rv := objc.Send[objc.ID](objc.ID(_NSTintConfigurationClass.class), objc.Sel("monochromeTintConfiguration"))
	return NSTintConfigurationFromID(objc.ID(rv))
}






















