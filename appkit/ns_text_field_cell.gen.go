// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSTextFieldCell] class.
var (
	_NSTextFieldCellClass     NSTextFieldCellClass
	_NSTextFieldCellClassOnce sync.Once
)

func getNSTextFieldCellClass() NSTextFieldCellClass {
	_NSTextFieldCellClassOnce.Do(func() {
		_NSTextFieldCellClass = NSTextFieldCellClass{class: objc.GetClass("NSTextFieldCell")}
	})
	return _NSTextFieldCellClass
}

// GetNSTextFieldCellClass returns the class object for NSTextFieldCell.
func GetNSTextFieldCellClass() NSTextFieldCellClass {
	return getNSTextFieldCellClass()
}

type NSTextFieldCellClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTextFieldCellClass) Alloc() NSTextFieldCell {
	rv := objc.Send[NSTextFieldCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that enhances the text display capabilities of a cell.
//
// # Overview
// 
// The [NSTextFieldCell] class adds to the text display capabilities of the
// [NSCell] class by allowing you to set the color of both the text and its
// background. You can also specify whether the cell draws its background at
// all.
// 
// All of the methods declared by this class are also declared by the
// [NSTextField] class, which uses [NSTextFieldCell] objects to draw and edit
// text. The [NSTextField] cover methods call the corresponding
// [NSTextFieldCell] methods.
// 
// Placeholder strings, set using the [NSTextFieldCell.PlaceholderString] or
// [NSTextFieldCell.PlaceholderAttributedString] property, appear in the text field cell if
// the actual string is `nil` or an empty string. They’re drawn in gray on
// the cell and aren’t archived in the “pre-10.2” nib format.
// 
// # Designated Initializers
// 
// When subclassing [NSTextFieldCell] you must implement the designated
// initializers [NSTextFieldCell.InitWithCoder] and [NSTextFieldCell.InitTextCell].
//
// # Setting the Text Color
//
//   - [NSTextFieldCell.TextColor]: The color to use to draw the cell’s text.
//   - [NSTextFieldCell.SetTextColor]
//
// # Setting the Bezel Style
//
//   - [NSTextFieldCell.BezelStyle]: The bezel style to use when drawing the text field.
//   - [NSTextFieldCell.SetBezelStyle]
//
// # Controlling the Background
//
//   - [NSTextFieldCell.BackgroundColor]: The color of the cell’s background.
//   - [NSTextFieldCell.SetBackgroundColor]
//   - [NSTextFieldCell.DrawsBackground]: A Boolean value that indicates whether the cell draws its background color.
//   - [NSTextFieldCell.SetDrawsBackground]
//
// # Managing Placeholder Strings
//
//   - [NSTextFieldCell.PlaceholderString]: The placeholder text for the cell, specified as a plain text string.
//   - [NSTextFieldCell.SetPlaceholderString]
//   - [NSTextFieldCell.PlaceholderAttributedString]: The placeholder text for the cell, specified as an attributed string.
//   - [NSTextFieldCell.SetPlaceholderAttributedString]
//
// # Accessing Input Source Locales
//
//   - [NSTextFieldCell.AllowedInputSourceLocales]: An array of locale identifiers that represent the allowed input sources when the text field has the keyboard focus.
//   - [NSTextFieldCell.SetAllowedInputSourceLocales]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFieldCell
type NSTextFieldCell struct {
	NSActionCell
}

// NSTextFieldCellFromID constructs a [NSTextFieldCell] from an objc.ID.
//
// An object that enhances the text display capabilities of a cell.
func NSTextFieldCellFromID(id objc.ID) NSTextFieldCell {
	return NSTextFieldCell{NSActionCell: NSActionCellFromID(id)}
}
// NOTE: NSTextFieldCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTextFieldCell] class.
//
// # Setting the Text Color
//
//   - [INSTextFieldCell.TextColor]: The color to use to draw the cell’s text.
//   - [INSTextFieldCell.SetTextColor]
//
// # Setting the Bezel Style
//
//   - [INSTextFieldCell.BezelStyle]: The bezel style to use when drawing the text field.
//   - [INSTextFieldCell.SetBezelStyle]
//
// # Controlling the Background
//
//   - [INSTextFieldCell.BackgroundColor]: The color of the cell’s background.
//   - [INSTextFieldCell.SetBackgroundColor]
//   - [INSTextFieldCell.DrawsBackground]: A Boolean value that indicates whether the cell draws its background color.
//   - [INSTextFieldCell.SetDrawsBackground]
//
// # Managing Placeholder Strings
//
//   - [INSTextFieldCell.PlaceholderString]: The placeholder text for the cell, specified as a plain text string.
//   - [INSTextFieldCell.SetPlaceholderString]
//   - [INSTextFieldCell.PlaceholderAttributedString]: The placeholder text for the cell, specified as an attributed string.
//   - [INSTextFieldCell.SetPlaceholderAttributedString]
//
// # Accessing Input Source Locales
//
//   - [INSTextFieldCell.AllowedInputSourceLocales]: An array of locale identifiers that represent the allowed input sources when the text field has the keyboard focus.
//   - [INSTextFieldCell.SetAllowedInputSourceLocales]
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFieldCell
type INSTextFieldCell interface {
	INSActionCell

	// Topic: Setting the Text Color

	// The color to use to draw the cell’s text.
	TextColor() INSColor
	SetTextColor(value INSColor)

	// Topic: Setting the Bezel Style

	// The bezel style to use when drawing the text field.
	BezelStyle() NSTextFieldBezelStyle
	SetBezelStyle(value NSTextFieldBezelStyle)

	// Topic: Controlling the Background

	// The color of the cell’s background.
	BackgroundColor() INSColor
	SetBackgroundColor(value INSColor)
	// A Boolean value that indicates whether the cell draws its background color.
	DrawsBackground() bool
	SetDrawsBackground(value bool)

	// Topic: Managing Placeholder Strings

	// The placeholder text for the cell, specified as a plain text string.
	PlaceholderString() string
	SetPlaceholderString(value string)
	// The placeholder text for the cell, specified as an attributed string.
	PlaceholderAttributedString() foundation.NSAttributedString
	SetPlaceholderAttributedString(value foundation.NSAttributedString)

	// Topic: Accessing Input Source Locales

	// An array of locale identifiers that represent the allowed input sources when the text field has the keyboard focus.
	AllowedInputSourceLocales() []string
	SetAllowedInputSourceLocales(value []string)
}

// Init initializes the instance.
func (t NSTextFieldCell) Init() NSTextFieldCell {
	rv := objc.Send[NSTextFieldCell](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTextFieldCell) Autorelease() NSTextFieldCell {
	rv := objc.Send[NSTextFieldCell](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTextFieldCell creates a new NSTextFieldCell instance.
func NewNSTextFieldCell() NSTextFieldCell {
	class := getNSTextFieldCellClass()
	rv := objc.Send[NSTextFieldCell](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an [NSCell] object initialized with the specified image and set to
// have the cell’s default menu.
//
// image: The image to use for the cell. If this parameter is `nil`, no image is set.
//
// # Return Value
// 
// An initialized [NSCell] object, or `nil` if the cell could not be
// initialized.
//
// # Discussion
// 
// This is one of four designated initializers you must implement when
// subclassing. See [NSCell] for the complete list.
//
// See: https://developer.apple.com/documentation/AppKit/NSCell/init(imageCell:)
func NewTextFieldCellImageCell(image INSImage) NSTextFieldCell {
	instance := getNSTextFieldCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSTextFieldCellFromID(rv)
}

// Initializes a text field cell that displays the specified string.
//
// string: The string that the text field cell displays.
//
// # Return Value
// 
// A text field cell that displays a string.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/init(textCell:)
func NewTextFieldCellTextCell(string_ string) NSTextFieldCell {
	instance := getNSTextFieldCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSTextFieldCellFromID(rv)
}

// Initializes a text field cell from data in the provided unarchiver.
//
// coder: An unarchiver object.
//
// # Return Value
// 
// A text field cell that displays a string.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/init(coder:)
func NewTextFieldCellWithCoder(coder foundation.INSCoder) NSTextFieldCell {
	instance := getNSTextFieldCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTextFieldCellFromID(rv)
}

// The color to use to draw the cell’s text.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/textColor
func (t NSTextFieldCell) TextColor() INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("textColor"))
	return NSColorFromID(objc.ID(rv))
}
func (t NSTextFieldCell) SetTextColor(value INSColor) {
	objc.Send[struct{}](t.ID, objc.Sel("setTextColor:"), value)
}

// The bezel style to use when drawing the text field.
//
// # Discussion
// 
// To set the bezel style, you must have already set the the text field’s
// [Bezeled] method with an argument of [true]. For a list of bezel styles,
// see [NSTextField.BezelStyle].
//
// [NSTextField.BezelStyle]: https://developer.apple.com/documentation/AppKit/NSTextField/BezelStyle-swift.enum
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/bezelStyle
func (t NSTextFieldCell) BezelStyle() NSTextFieldBezelStyle {
	rv := objc.Send[NSTextFieldBezelStyle](t.ID, objc.Sel("bezelStyle"))
	return NSTextFieldBezelStyle(rv)
}
func (t NSTextFieldCell) SetBezelStyle(value NSTextFieldBezelStyle) {
	objc.Send[struct{}](t.ID, objc.Sel("setBezelStyle:"), value)
}

// The color of the cell’s background.
//
// # Discussion
// 
// The background color is drawn behind the cell’s text.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/backgroundColor
func (t NSTextFieldCell) BackgroundColor() INSColor {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("backgroundColor"))
	return NSColorFromID(objc.ID(rv))
}
func (t NSTextFieldCell) SetBackgroundColor(value INSColor) {
	objc.Send[struct{}](t.ID, objc.Sel("setBackgroundColor:"), value)
}

// A Boolean value that indicates whether the cell draws its background color.
//
// # Discussion
// 
// When the value of this property is [true], the cell draws its background
// color. In order to prevent inconsistent rendering, background color
// rendering is automatically disabled for rounded-bezel text fields.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/drawsBackground
func (t NSTextFieldCell) DrawsBackground() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("drawsBackground"))
	return rv
}
func (t NSTextFieldCell) SetDrawsBackground(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setDrawsBackground:"), value)
}

// The placeholder text for the cell, specified as a plain text string.
//
// # Discussion
// 
// Assigning a new value to this property also clears out any value set for
// the [PlaceholderAttributedString] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/placeholderString
func (t NSTextFieldCell) PlaceholderString() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("placeholderString"))
	return foundation.NSStringFromID(rv).String()
}
func (t NSTextFieldCell) SetPlaceholderString(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setPlaceholderString:"), objc.String(value))
}

// The placeholder text for the cell, specified as an attributed string.
//
// # Discussion
// 
// Assigning a new value to this property also clears out any value set for
// the [PlaceholderString] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/placeholderAttributedString
func (t NSTextFieldCell) PlaceholderAttributedString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("placeholderAttributedString"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (t NSTextFieldCell) SetPlaceholderAttributedString(value foundation.NSAttributedString) {
	objc.Send[struct{}](t.ID, objc.Sel("setPlaceholderAttributedString:"), value)
}

// An array of locale identifiers that represent the allowed input sources
// when the text field has the keyboard focus.
//
// # Discussion
// 
// The value of this property is an array of [NSString] objects, each of which
// contains a locale identifier. You can assign the meta-locale identifier,
// [NSAllRomanInputSourcesLocaleIdentifier], to specify input sources that are
// limited for Roman script editing.
//
// [NSAllRomanInputSourcesLocaleIdentifier]: https://developer.apple.com/documentation/AppKit/NSAllRomanInputSourcesLocaleIdentifier
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
//
// See: https://developer.apple.com/documentation/AppKit/NSTextFieldCell/allowedInputSourceLocales
func (t NSTextFieldCell) AllowedInputSourceLocales() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("allowedInputSourceLocales"))
	return objc.ConvertSliceToStrings(rv)
}
func (t NSTextFieldCell) SetAllowedInputSourceLocales(value []string) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowedInputSourceLocales:"), objectivec.StringSliceToNSArray(value))
}

