// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSecureTextFieldCell] class.
var (
	_NSSecureTextFieldCellClass     NSSecureTextFieldCellClass
	_NSSecureTextFieldCellClassOnce sync.Once
)

func getNSSecureTextFieldCellClass() NSSecureTextFieldCellClass {
	_NSSecureTextFieldCellClassOnce.Do(func() {
		_NSSecureTextFieldCellClass = NSSecureTextFieldCellClass{class: objc.GetClass("NSSecureTextFieldCell")}
	})
	return _NSSecureTextFieldCellClass
}

// GetNSSecureTextFieldCellClass returns the class object for NSSecureTextFieldCell.
func GetNSSecureTextFieldCellClass() NSSecureTextFieldCellClass {
	return getNSSecureTextFieldCellClass()
}

type NSSecureTextFieldCellClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSecureTextFieldCellClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSecureTextFieldCellClass) Alloc() NSSecureTextFieldCell {
	rv := objc.Send[NSSecureTextFieldCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A text field whose value is hidden from the user.
//
// # Overview
//
// [NSSecureTextFieldCell] works with [NSSecureTextField] and overrides the
// general cell use of the field editor to provide its own field editor, which
// doesn’t display text or allow the user to cut or copy its value.
//
// # Working with character echo
//
//   - [NSSecureTextFieldCell.EchosBullets]: A Boolean that indicates whether the receiver echoes a bullet character rather than each character typed.
//   - [NSSecureTextFieldCell.SetEchosBullets]
//
// See: https://developer.apple.com/documentation/AppKit/NSSecureTextFieldCell
type NSSecureTextFieldCell struct {
	NSTextFieldCell
}

// NSSecureTextFieldCellFromID constructs a [NSSecureTextFieldCell] from an objc.ID.
//
// A text field whose value is hidden from the user.
func NSSecureTextFieldCellFromID(id objc.ID) NSSecureTextFieldCell {
	return NSSecureTextFieldCell{NSTextFieldCell: NSTextFieldCellFromID(id)}
}

// NOTE: NSSecureTextFieldCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSecureTextFieldCell] class.
//
// # Working with character echo
//
//   - [INSSecureTextFieldCell.EchosBullets]: A Boolean that indicates whether the receiver echoes a bullet character rather than each character typed.
//   - [INSSecureTextFieldCell.SetEchosBullets]
//
// See: https://developer.apple.com/documentation/AppKit/NSSecureTextFieldCell
type INSSecureTextFieldCell interface {
	INSTextFieldCell

	// Topic: Working with character echo

	// A Boolean that indicates whether the receiver echoes a bullet character rather than each character typed.
	EchosBullets() bool
	SetEchosBullets(value bool)
}

// Init initializes the instance.
func (s NSSecureTextFieldCell) Init() NSSecureTextFieldCell {
	rv := objc.Send[NSSecureTextFieldCell](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSecureTextFieldCell) Autorelease() NSSecureTextFieldCell {
	rv := objc.Send[NSSecureTextFieldCell](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSecureTextFieldCell creates a new NSSecureTextFieldCell instance.
func NewNSSecureTextFieldCell() NSSecureTextFieldCell {
	class := getNSSecureTextFieldCellClass()
	rv := objc.Send[NSSecureTextFieldCell](objc.ID(class.class), objc.Sel("new"))
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
func NewSecureTextFieldCellImageCell(image INSImage) NSSecureTextFieldCell {
	instance := getNSSecureTextFieldCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSSecureTextFieldCellFromID(rv)
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
func NewSecureTextFieldCellTextCell(string_ string) NSSecureTextFieldCell {
	instance := getNSSecureTextFieldCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSSecureTextFieldCellFromID(rv)
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
func NewSecureTextFieldCellWithCoder(coder foundation.INSCoder) NSSecureTextFieldCell {
	instance := getNSSecureTextFieldCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSSecureTextFieldCellFromID(rv)
}

// A Boolean that indicates whether the receiver echoes a bullet character
// rather than each character typed.
//
// # Discussion
//
// Default is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSSecureTextFieldCell/echosBullets
func (s NSSecureTextFieldCell) EchosBullets() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("echosBullets"))
	return rv
}
func (s NSSecureTextFieldCell) SetEchosBullets(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setEchosBullets:"), value)
}
