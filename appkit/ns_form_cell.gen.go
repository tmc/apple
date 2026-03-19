// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSFormCell] class.
var (
	_NSFormCellClass     NSFormCellClass
	_NSFormCellClassOnce sync.Once
)

func getNSFormCellClass() NSFormCellClass {
	_NSFormCellClassOnce.Do(func() {
		_NSFormCellClass = NSFormCellClass{class: objc.GetClass("NSFormCell")}
	})
	return _NSFormCellClass
}

// GetNSFormCellClass returns the class object for NSFormCell.
func GetNSFormCellClass() NSFormCellClass {
	return getNSFormCellClass()
}

type NSFormCellClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFormCellClass) Alloc() NSFormCell {
	rv := objc.Send[NSFormCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The [NSFormCell] class is used to implement text entry fields in a form.
// The left part of an [NSFormCell] object contains a title. The right part
// contains an editable text entry field.
//
// # Overview
// 
// An [NSFormCell] object implements the user interface of an [NSForm] object.
//
// [NSForm]: https://developer.apple.com/documentation/AppKit/NSForm
//
// # Accessing a Cell’s Title
//
//   - [NSFormCell.AttributedTitle]: The title of the cell as an attributed string.
//   - [NSFormCell.SetAttributedTitle]
//   - [NSFormCell.TitleAlignment]: The alignment of the title.
//   - [NSFormCell.SetTitleAlignment]
//   - [NSFormCell.TitleBaseWritingDirection]: The default writing direction used to render the form cell’s title.
//   - [NSFormCell.SetTitleBaseWritingDirection]
//   - [NSFormCell.TitleFont]: The font used to draw cell’s title.
//   - [NSFormCell.SetTitleFont]
//   - [NSFormCell.TitleWidth]: The width of the title field.
//   - [NSFormCell.SetTitleWidth]
//
// # Asking About Placeholder Values
//
//   - [NSFormCell.PlaceholderAttributedString]: The cell’s attributed placeholder string.
//   - [NSFormCell.SetPlaceholderAttributedString]
//   - [NSFormCell.PlaceholderString]: The cell’s plain text placeholder string.
//   - [NSFormCell.SetPlaceholderString]
//
// # Sizing for Auto Layout
//
//   - [NSFormCell.PreferredTextFieldWidth]: The preferred text field width.
//   - [NSFormCell.SetPreferredTextFieldWidth]
//
// See: https://developer.apple.com/documentation/AppKit/NSFormCell
type NSFormCell struct {
	NSActionCell
}

// NSFormCellFromID constructs a [NSFormCell] from an objc.ID.
//
// The [NSFormCell] class is used to implement text entry fields in a form.
// The left part of an [NSFormCell] object contains a title. The right part
// contains an editable text entry field.
func NSFormCellFromID(id objc.ID) NSFormCell {
	return NSFormCell{NSActionCell: NSActionCellFromID(id)}
}
// NOTE: NSFormCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFormCell] class.
//
// # Accessing a Cell’s Title
//
//   - [INSFormCell.AttributedTitle]: The title of the cell as an attributed string.
//   - [INSFormCell.SetAttributedTitle]
//   - [INSFormCell.TitleAlignment]: The alignment of the title.
//   - [INSFormCell.SetTitleAlignment]
//   - [INSFormCell.TitleBaseWritingDirection]: The default writing direction used to render the form cell’s title.
//   - [INSFormCell.SetTitleBaseWritingDirection]
//   - [INSFormCell.TitleFont]: The font used to draw cell’s title.
//   - [INSFormCell.SetTitleFont]
//   - [INSFormCell.TitleWidth]: The width of the title field.
//   - [INSFormCell.SetTitleWidth]
//
// # Asking About Placeholder Values
//
//   - [INSFormCell.PlaceholderAttributedString]: The cell’s attributed placeholder string.
//   - [INSFormCell.SetPlaceholderAttributedString]
//   - [INSFormCell.PlaceholderString]: The cell’s plain text placeholder string.
//   - [INSFormCell.SetPlaceholderString]
//
// # Sizing for Auto Layout
//
//   - [INSFormCell.PreferredTextFieldWidth]: The preferred text field width.
//   - [INSFormCell.SetPreferredTextFieldWidth]
//
// See: https://developer.apple.com/documentation/AppKit/NSFormCell
type INSFormCell interface {
	INSActionCell

	// Topic: Accessing a Cell’s Title

	// The title of the cell as an attributed string.
	AttributedTitle() foundation.NSAttributedString
	SetAttributedTitle(value foundation.NSAttributedString)
	// The alignment of the title.
	TitleAlignment() NSTextAlignment
	SetTitleAlignment(value NSTextAlignment)
	// The default writing direction used to render the form cell’s title.
	TitleBaseWritingDirection() NSWritingDirection
	SetTitleBaseWritingDirection(value NSWritingDirection)
	// The font used to draw cell’s title.
	TitleFont() NSFont
	SetTitleFont(value NSFont)
	// The width of the title field.
	TitleWidth() float64
	SetTitleWidth(value float64)

	// Topic: Asking About Placeholder Values

	// The cell’s attributed placeholder string.
	PlaceholderAttributedString() foundation.NSAttributedString
	SetPlaceholderAttributedString(value foundation.NSAttributedString)
	// The cell’s plain text placeholder string.
	PlaceholderString() string
	SetPlaceholderString(value string)

	// Topic: Sizing for Auto Layout

	// The preferred text field width.
	PreferredTextFieldWidth() float64
	SetPreferredTextFieldWidth(value float64)
}

// Init initializes the instance.
func (f NSFormCell) Init() NSFormCell {
	rv := objc.Send[NSFormCell](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFormCell) Autorelease() NSFormCell {
	rv := objc.Send[NSFormCell](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFormCell creates a new NSFormCell instance.
func NewNSFormCell() NSFormCell {
	class := getNSFormCellClass()
	rv := objc.Send[NSFormCell](objc.ID(class.class), objc.Sel("new"))
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
func NewFormCellImageCell(image INSImage) NSFormCell {
	instance := getNSFormCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSFormCellFromID(rv)
}

// Returns an [NSFormCell] object initialized with the specified title string.
//
// string: The title for the new form cell object.
//
// # Return Value
// 
// An initialized [NSFormCell] object.
//
// # Discussion
// 
// The contents of the cell’s editable text entry field are set to the empty
// string (`@""`). The font for both title and text is the user’s chosen
// system font in 12.0 point, and the text area is drawn with a bezel. This
// method is the designated initializer for the [NSFormCell] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSFormCell/init(textCell:)
func NewFormCellTextCell(string_ string) NSFormCell {
	instance := getNSFormCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSFormCellFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppKit/NSFormCell/init(coder:)
func NewFormCellWithCoder(coder foundation.INSCoder) NSFormCell {
	instance := getNSFormCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSFormCellFromID(rv)
}

// The title of the cell as an attributed string.
//
// See: https://developer.apple.com/documentation/AppKit/NSFormCell/attributedTitle
func (f NSFormCell) AttributedTitle() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("attributedTitle"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (f NSFormCell) SetAttributedTitle(value foundation.NSAttributedString) {
	objc.Send[struct{}](f.ID, objc.Sel("setAttributedTitle:"), value)
}
// The alignment of the title.
//
// # Discussion
// 
// The alignment can be one of the following values: [NSLeftTextAlignment],
// [NSCenterTextAlignment], or [NSRightTextAlignment]. The default alignment
// is [NSRightTextAlignment].
//
// See: https://developer.apple.com/documentation/AppKit/NSFormCell/titleAlignment
func (f NSFormCell) TitleAlignment() NSTextAlignment {
	rv := objc.Send[NSTextAlignment](f.ID, objc.Sel("titleAlignment"))
	return NSTextAlignment(rv)
}
func (f NSFormCell) SetTitleAlignment(value NSTextAlignment) {
	objc.Send[struct{}](f.ID, objc.Sel("setTitleAlignment:"), value)
}
// The default writing direction used to render the form cell’s title.
//
// # Discussion
// 
// Can be one of the following constants: [NSWritingDirectionNatural],
// [NSWritingDirectionLeftToRight], or [NSWritingDirectionRightToLeft].
//
// See: https://developer.apple.com/documentation/AppKit/NSFormCell/titleBaseWritingDirection
func (f NSFormCell) TitleBaseWritingDirection() NSWritingDirection {
	rv := objc.Send[NSWritingDirection](f.ID, objc.Sel("titleBaseWritingDirection"))
	return NSWritingDirection(rv)
}
func (f NSFormCell) SetTitleBaseWritingDirection(value NSWritingDirection) {
	objc.Send[struct{}](f.ID, objc.Sel("setTitleBaseWritingDirection:"), value)
}
// The font used to draw cell’s title.
//
// See: https://developer.apple.com/documentation/AppKit/NSFormCell/titleFont
func (f NSFormCell) TitleFont() NSFont {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("titleFont"))
	return NSFontFromID(objc.ID(rv))
}
func (f NSFormCell) SetTitleFont(value NSFont) {
	objc.Send[struct{}](f.ID, objc.Sel("setTitleFont:"), value)
}
// The width of the title field.
//
// # Discussion
// 
// The width of the title field, measured in points in the user coordinate
// space. You usually do not need to set this property. AppKit automatically
// sets the title width whenever the title changes. If the automatic width
// doesn’t suit your needs, though, you can use this property to set the
// width explicitly.
// 
// After you have set the width this way, AppKit stops setting the width
// automatically, so you must set this property every time the title changes.
// If you want AppKit to resume automatic width assignments, set this property
// to a negative value.
//
// See: https://developer.apple.com/documentation/AppKit/NSFormCell/titleWidth
func (f NSFormCell) TitleWidth() float64 {
	rv := objc.Send[float64](f.ID, objc.Sel("titleWidth"))
	return rv
}
func (f NSFormCell) SetTitleWidth(value float64) {
	objc.Send[struct{}](f.ID, objc.Sel("setTitleWidth:"), value)
}
// The cell’s attributed placeholder string.
//
// # Discussion
// 
// If this property returns `nil`, you can also call `placeholderString` to
// see if the cell has a plain text placeholder string.
//
// See: https://developer.apple.com/documentation/AppKit/NSFormCell/placeholderAttributedString
func (f NSFormCell) PlaceholderAttributedString() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("placeholderAttributedString"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (f NSFormCell) SetPlaceholderAttributedString(value foundation.NSAttributedString) {
	objc.Send[struct{}](f.ID, objc.Sel("setPlaceholderAttributedString:"), value)
}
// The cell’s plain text placeholder string.
//
// # Discussion
// 
// If this property returns `nil`, you can also call
// `placeholderAttributedString` to see if the cell has an attributed
// placeholder string. Note that invoking this method clears out any
// attributed string set by the [PlaceholderAttributedString] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSFormCell/placeholderString
func (f NSFormCell) PlaceholderString() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("placeholderString"))
	return foundation.NSStringFromID(rv).String()
}
func (f NSFormCell) SetPlaceholderString(value string) {
	objc.Send[struct{}](f.ID, objc.Sel("setPlaceholderString:"), objc.String(value))
}
// The preferred text field width.
//
// # Discussion
// 
// The preferred width is reflected in the cell’s `cellSize`, which will be
// large enough to accommodate the title, bezel, and a text field of width
// `preferredTextWidth`. It is also reflected in the [intrinsicContentSize] of
// the [NSForm] object. That is, under Auto Layout, the form will try to size
// itself so that the text field cell is the given width, according to the
// usual content size constraint priorities.
// 
// If the width is negative, the cel size matches the historic behavior, which
// is that it is large enough to accommodate the title, bezel, and the current
// text.
// 
// This property can aid migration to Auto Layout, and is sufficient for
// simple cases. However, for new apps, it’s recommended that you use an
// [NSTextField] instance directly instead of a form.
// 
// The default value of this property is `-1`.
//
// [NSForm]: https://developer.apple.com/documentation/AppKit/NSForm
// [intrinsicContentSize]: https://developer.apple.com/documentation/AppKit/NSView/intrinsicContentSize
//
// See: https://developer.apple.com/documentation/AppKit/NSFormCell/preferredTextFieldWidth
func (f NSFormCell) PreferredTextFieldWidth() float64 {
	rv := objc.Send[float64](f.ID, objc.Sel("preferredTextFieldWidth"))
	return rv
}
func (f NSFormCell) SetPreferredTextFieldWidth(value float64) {
	objc.Send[struct{}](f.ID, objc.Sel("setPreferredTextFieldWidth:"), value)
}

