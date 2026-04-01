// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSColorPickerTouchBarItem] class.
var (
	_NSColorPickerTouchBarItemClass     NSColorPickerTouchBarItemClass
	_NSColorPickerTouchBarItemClassOnce sync.Once
)

func getNSColorPickerTouchBarItemClass() NSColorPickerTouchBarItemClass {
	_NSColorPickerTouchBarItemClassOnce.Do(func() {
		_NSColorPickerTouchBarItemClass = NSColorPickerTouchBarItemClass{class: objc.GetClass("NSColorPickerTouchBarItem")}
	})
	return _NSColorPickerTouchBarItemClass
}

// GetNSColorPickerTouchBarItemClass returns the class object for NSColorPickerTouchBarItem.
func GetNSColorPickerTouchBarItemClass() NSColorPickerTouchBarItemClass {
	return getNSColorPickerTouchBarItemClass()
}

type NSColorPickerTouchBarItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSColorPickerTouchBarItemClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSColorPickerTouchBarItemClass) Alloc() NSColorPickerTouchBarItem {
	rv := objc.Send[NSColorPickerTouchBarItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A bar item that provides a system-defined color picker.
//
// # Overview
//
// For design guidance, see [Human Interface Guidelines].
//
// # Configuring the color picker
//
//   - [NSColorPickerTouchBarItem.ColorList]: The list of colors displayed in the color picker.
//   - [NSColorPickerTouchBarItem.SetColorList]
//   - [NSColorPickerTouchBarItem.AllowedColorSpaces]: Controls the color spaces that the color picker can produce.
//   - [NSColorPickerTouchBarItem.SetAllowedColorSpaces]
//   - [NSColorPickerTouchBarItem.ShowsAlpha]: A Boolean value that controls whether the color picker allows picking of colors with alpha values other than `1.0`.
//   - [NSColorPickerTouchBarItem.SetShowsAlpha]
//   - [NSColorPickerTouchBarItem.Enabled]: A Boolean value that determines whether the color picker is enabled.
//   - [NSColorPickerTouchBarItem.SetEnabled]
//
// # Obtaining the selected color
//
//   - [NSColorPickerTouchBarItem.Color]: The picker’s currently selected color.
//   - [NSColorPickerTouchBarItem.SetColor]
//   - [NSColorPickerTouchBarItem.Target]: An object that is notified when a user interacts with the color picker.
//   - [NSColorPickerTouchBarItem.SetTarget]
//   - [NSColorPickerTouchBarItem.Action]: The selector on the target object that is invoked when a user interacts with the color picker.
//   - [NSColorPickerTouchBarItem.SetAction]
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem
//
// [Human Interface Guidelines]: https://developer.apple.com/design/human-interface-guidelines/macos/touch-bar/touch-bar-controls-and-views/#color-pickers
type NSColorPickerTouchBarItem struct {
	NSTouchBarItem
}

// NSColorPickerTouchBarItemFromID constructs a [NSColorPickerTouchBarItem] from an objc.ID.
//
// A bar item that provides a system-defined color picker.
func NSColorPickerTouchBarItemFromID(id objc.ID) NSColorPickerTouchBarItem {
	return NSColorPickerTouchBarItem{NSTouchBarItem: NSTouchBarItemFromID(id)}
}

// NOTE: NSColorPickerTouchBarItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSColorPickerTouchBarItem] class.
//
// # Configuring the color picker
//
//   - [INSColorPickerTouchBarItem.ColorList]: The list of colors displayed in the color picker.
//   - [INSColorPickerTouchBarItem.SetColorList]
//   - [INSColorPickerTouchBarItem.AllowedColorSpaces]: Controls the color spaces that the color picker can produce.
//   - [INSColorPickerTouchBarItem.SetAllowedColorSpaces]
//   - [INSColorPickerTouchBarItem.ShowsAlpha]: A Boolean value that controls whether the color picker allows picking of colors with alpha values other than `1.0`.
//   - [INSColorPickerTouchBarItem.SetShowsAlpha]
//   - [INSColorPickerTouchBarItem.Enabled]: A Boolean value that determines whether the color picker is enabled.
//   - [INSColorPickerTouchBarItem.SetEnabled]
//
// # Obtaining the selected color
//
//   - [INSColorPickerTouchBarItem.Color]: The picker’s currently selected color.
//   - [INSColorPickerTouchBarItem.SetColor]
//   - [INSColorPickerTouchBarItem.Target]: An object that is notified when a user interacts with the color picker.
//   - [INSColorPickerTouchBarItem.SetTarget]
//   - [INSColorPickerTouchBarItem.Action]: The selector on the target object that is invoked when a user interacts with the color picker.
//   - [INSColorPickerTouchBarItem.SetAction]
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem
type INSColorPickerTouchBarItem interface {
	INSTouchBarItem

	// Topic: Configuring the color picker

	// The list of colors displayed in the color picker.
	ColorList() INSColorList
	SetColorList(value INSColorList)
	// Controls the color spaces that the color picker can produce.
	AllowedColorSpaces() []NSColorSpace
	SetAllowedColorSpaces(value []NSColorSpace)
	// A Boolean value that controls whether the color picker allows picking of colors with alpha values other than `1.0`.
	ShowsAlpha() bool
	SetShowsAlpha(value bool)
	// A Boolean value that determines whether the color picker is enabled.
	Enabled() bool
	SetEnabled(value bool)

	// Topic: Obtaining the selected color

	// The picker’s currently selected color.
	Color() INSColor
	SetColor(value INSColor)
	// An object that is notified when a user interacts with the color picker.
	Target() objectivec.IObject
	SetTarget(value objectivec.IObject)
	// The selector on the target object that is invoked when a user interacts with the color picker.
	Action() objc.SEL
	SetAction(value objc.SEL)
}

// Init initializes the instance.
func (c NSColorPickerTouchBarItem) Init() NSColorPickerTouchBarItem {
	rv := objc.Send[NSColorPickerTouchBarItem](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSColorPickerTouchBarItem) Autorelease() NSColorPickerTouchBarItem {
	rv := objc.Send[NSColorPickerTouchBarItem](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSColorPickerTouchBarItem creates a new NSColorPickerTouchBarItem instance.
func NewNSColorPickerTouchBarItem() NSColorPickerTouchBarItem {
	class := getNSColorPickerTouchBarItemClass()
	rv := objc.Send[NSColorPickerTouchBarItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes and returns a new item from a storyboard or nib file.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(coder:)
func NewColorPickerTouchBarItemWithCoder(coder foundation.INSCoder) NSColorPickerTouchBarItem {
	instance := getNSColorPickerTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSColorPickerTouchBarItemFromID(rv)
}

// Creates a new item with the specified identifier.
//
// # Discussion
//
// The designated initializer. The identifier must be globally unique for
// every item, except for space items.
//
// See: https://developer.apple.com/documentation/AppKit/NSTouchBarItem/init(identifier:)
func NewColorPickerTouchBarItemWithIdentifier(identifier NSTouchBarItemIdentifier) NSColorPickerTouchBarItem {
	instance := getNSColorPickerTouchBarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(string(identifier)))
	return NSColorPickerTouchBarItemFromID(rv)
}

// Creates a bar item with the standard color picker icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/colorPicker(withIdentifier:)
func (_NSColorPickerTouchBarItemClass NSColorPickerTouchBarItemClass) ColorPickerWithIdentifier(identifier NSTouchBarItemIdentifier) NSColorPickerTouchBarItem {
	rv := objc.Send[objc.ID](objc.ID(_NSColorPickerTouchBarItemClass.class), objc.Sel("colorPickerWithIdentifier:"), objc.String(string(identifier)))
	return NSColorPickerTouchBarItemFromID(rv)
}

// Creates a bar item with the standard text color picker icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/textColorPicker(withIdentifier:)
func (_NSColorPickerTouchBarItemClass NSColorPickerTouchBarItemClass) TextColorPickerWithIdentifier(identifier NSTouchBarItemIdentifier) NSColorPickerTouchBarItem {
	rv := objc.Send[objc.ID](objc.ID(_NSColorPickerTouchBarItemClass.class), objc.Sel("textColorPickerWithIdentifier:"), objc.String(string(identifier)))
	return NSColorPickerTouchBarItemFromID(rv)
}

// Creates a bar item with the standard stroke color picker icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/strokeColorPicker(withIdentifier:)
func (_NSColorPickerTouchBarItemClass NSColorPickerTouchBarItemClass) StrokeColorPickerWithIdentifier(identifier NSTouchBarItemIdentifier) NSColorPickerTouchBarItem {
	rv := objc.Send[objc.ID](objc.ID(_NSColorPickerTouchBarItemClass.class), objc.Sel("strokeColorPickerWithIdentifier:"), objc.String(string(identifier)))
	return NSColorPickerTouchBarItemFromID(rv)
}

// Creates a color picker bar item using the supplied image as its icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/colorPicker(withIdentifier:buttonImage:)
func (_NSColorPickerTouchBarItemClass NSColorPickerTouchBarItemClass) ColorPickerWithIdentifierButtonImage(identifier NSTouchBarItemIdentifier, image INSImage) NSColorPickerTouchBarItem {
	rv := objc.Send[objc.ID](objc.ID(_NSColorPickerTouchBarItemClass.class), objc.Sel("colorPickerWithIdentifier:buttonImage:"), objc.String(string(identifier)), image)
	return NSColorPickerTouchBarItemFromID(rv)
}

// The list of colors displayed in the color picker.
//
// # Discussion
//
// Defaults to the standard system color list.
//
// Note that setting a custom color list disables the additional tints and
// shades that appear when the user sustains a touch.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/colorList
func (c NSColorPickerTouchBarItem) ColorList() INSColorList {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("colorList"))
	return NSColorListFromID(objc.ID(rv))
}
func (c NSColorPickerTouchBarItem) SetColorList(value INSColorList) {
	objc.Send[struct{}](c.ID, objc.Sel("setColorList:"), value)
}

// Controls the color spaces that the color picker can produce.
//
// # Discussion
//
// If a selected color is outside the allowed color spaces, the picker
// converts it to the first color space in the array.
//
// Set the value of this property to `nil` to allow all color spaces. An empty
// array is an invalid value.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/allowedColorSpaces
func (c NSColorPickerTouchBarItem) AllowedColorSpaces() []NSColorSpace {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("allowedColorSpaces"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSColorSpace {
		return NSColorSpaceFromID(id)
	})
}
func (c NSColorPickerTouchBarItem) SetAllowedColorSpaces(value []NSColorSpace) {
	objc.Send[struct{}](c.ID, objc.Sel("setAllowedColorSpaces:"), objectivec.IObjectSliceToNSArray(value))
}

// A Boolean value that controls whether the color picker allows picking of
// colors with alpha values other than `1.0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/showsAlpha
func (c NSColorPickerTouchBarItem) ShowsAlpha() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("showsAlpha"))
	return rv
}
func (c NSColorPickerTouchBarItem) SetShowsAlpha(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setShowsAlpha:"), value)
}

// A Boolean value that determines whether the color picker is enabled.
//
// # Discussion
//
// If the picker is currently displayed as a popover, and you set the value of
// this property to false, the picker is dismissed.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/isEnabled
func (c NSColorPickerTouchBarItem) Enabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isEnabled"))
	return rv
}
func (c NSColorPickerTouchBarItem) SetEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setEnabled:"), value)
}

// The picker’s currently selected color.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/color
func (c NSColorPickerTouchBarItem) Color() INSColor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("color"))
	return NSColorFromID(objc.ID(rv))
}
func (c NSColorPickerTouchBarItem) SetColor(value INSColor) {
	objc.Send[struct{}](c.ID, objc.Sel("setColor:"), value)
}

// An object that is notified when a user interacts with the color picker.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/target
func (c NSColorPickerTouchBarItem) Target() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("target"))
	return objectivec.Object{ID: rv}
}
func (c NSColorPickerTouchBarItem) SetTarget(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setTarget:"), value)
}

// The selector on the target object that is invoked when a user interacts
// with the color picker.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPickerTouchBarItem/action
func (c NSColorPickerTouchBarItem) Action() objc.SEL {
	rv := objc.Send[objc.SEL](c.ID, objc.Sel("action"))
	return rv
}
func (c NSColorPickerTouchBarItem) SetAction(value objc.SEL) {
	objc.Send[struct{}](c.ID, objc.Sel("setAction:"), value)
}
