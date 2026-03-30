// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSTokenFieldCell] class.
var (
	_NSTokenFieldCellClass     NSTokenFieldCellClass
	_NSTokenFieldCellClassOnce sync.Once
)

func getNSTokenFieldCellClass() NSTokenFieldCellClass {
	_NSTokenFieldCellClassOnce.Do(func() {
		_NSTokenFieldCellClass = NSTokenFieldCellClass{class: objc.GetClass("NSTokenFieldCell")}
	})
	return _NSTokenFieldCellClass
}

// GetNSTokenFieldCellClass returns the class object for NSTokenFieldCell.
func GetNSTokenFieldCellClass() NSTokenFieldCellClass {
	return getNSTokenFieldCellClass()
}

type NSTokenFieldCellClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTokenFieldCellClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTokenFieldCellClass) Alloc() NSTokenFieldCell {
	rv := objc.Send[NSTokenFieldCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A text field cell subclass that enables tokenized editing of an array of
// objects.
//
// # Overview
//
// [NSTokenFieldCell] is a subclass of [NSTextFieldCell] that provides
// tokenized editing of an array of objects similar to the address field in
// the Mail app. The objects may be strings or objects that can be represented
// as strings. A single token field cell can be presented in an [NSTokenField]
// control.
//
// # Managing the Token Style
//
//   - [NSTokenFieldCell.TokenStyle]: The token style of the receiver.
//   - [NSTokenFieldCell.SetTokenStyle]
//
// # Managing the Tokenizing Character Set
//
//   - [NSTokenFieldCell.TokenizingCharacterSet]: The receiver’s tokenizing character set to a given character set.
//   - [NSTokenFieldCell.SetTokenizingCharacterSet]
//
// # Configuring the Completion Delay
//
//   - [NSTokenFieldCell.CompletionDelay]: The receiver’s completion delay to a given delay.
//   - [NSTokenFieldCell.SetCompletionDelay]
//
// # Managing the Delegate
//
//   - [NSTokenFieldCell.Delegate]: The receiver’s delegate.
//   - [NSTokenFieldCell.SetDelegate]
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell
type NSTokenFieldCell struct {
	NSTextFieldCell
}

// NSTokenFieldCellFromID constructs a [NSTokenFieldCell] from an objc.ID.
//
// A text field cell subclass that enables tokenized editing of an array of
// objects.
func NSTokenFieldCellFromID(id objc.ID) NSTokenFieldCell {
	return NSTokenFieldCell{NSTextFieldCell: NSTextFieldCellFromID(id)}
}

// NOTE: NSTokenFieldCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTokenFieldCell] class.
//
// # Managing the Token Style
//
//   - [INSTokenFieldCell.TokenStyle]: The token style of the receiver.
//   - [INSTokenFieldCell.SetTokenStyle]
//
// # Managing the Tokenizing Character Set
//
//   - [INSTokenFieldCell.TokenizingCharacterSet]: The receiver’s tokenizing character set to a given character set.
//   - [INSTokenFieldCell.SetTokenizingCharacterSet]
//
// # Configuring the Completion Delay
//
//   - [INSTokenFieldCell.CompletionDelay]: The receiver’s completion delay to a given delay.
//   - [INSTokenFieldCell.SetCompletionDelay]
//
// # Managing the Delegate
//
//   - [INSTokenFieldCell.Delegate]: The receiver’s delegate.
//   - [INSTokenFieldCell.SetDelegate]
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell
type INSTokenFieldCell interface {
	INSTextFieldCell

	// Topic: Managing the Token Style

	// The token style of the receiver.
	TokenStyle() NSTokenStyle
	SetTokenStyle(value NSTokenStyle)

	// Topic: Managing the Tokenizing Character Set

	// The receiver’s tokenizing character set to a given character set.
	TokenizingCharacterSet() foundation.NSCharacterSet
	SetTokenizingCharacterSet(value foundation.NSCharacterSet)

	// Topic: Configuring the Completion Delay

	// The receiver’s completion delay to a given delay.
	CompletionDelay() float64
	SetCompletionDelay(value float64)

	// Topic: Managing the Delegate

	// The receiver’s delegate.
	Delegate() NSTokenFieldCellDelegate
	SetDelegate(value NSTokenFieldCellDelegate)
}

// Init initializes the instance.
func (t NSTokenFieldCell) Init() NSTokenFieldCell {
	rv := objc.Send[NSTokenFieldCell](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTokenFieldCell) Autorelease() NSTokenFieldCell {
	rv := objc.Send[NSTokenFieldCell](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTokenFieldCell creates a new NSTokenFieldCell instance.
func NewNSTokenFieldCell() NSTokenFieldCell {
	class := getNSTokenFieldCellClass()
	rv := objc.Send[NSTokenFieldCell](objc.ID(class.class), objc.Sel("new"))
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
func NewTokenFieldCellImageCell(image INSImage) NSTokenFieldCell {
	instance := getNSTokenFieldCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSTokenFieldCellFromID(rv)
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
func NewTokenFieldCellTextCell(string_ string) NSTokenFieldCell {
	instance := getNSTokenFieldCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSTokenFieldCellFromID(rv)
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
func NewTokenFieldCellWithCoder(coder foundation.INSCoder) NSTokenFieldCell {
	instance := getNSTokenFieldCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTokenFieldCellFromID(rv)
}

// The token style of the receiver.
//
// # Discussion
//
// The valid values are described in [NSTokenField.TokenStyle].
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/tokenStyle
//
// [NSTokenField.TokenStyle]: https://developer.apple.com/documentation/AppKit/NSTokenField/TokenStyle-swift.enum
func (t NSTokenFieldCell) TokenStyle() NSTokenStyle {
	rv := objc.Send[NSTokenStyle](t.ID, objc.Sel("tokenStyle"))
	return NSTokenStyle(rv)
}
func (t NSTokenFieldCell) SetTokenStyle(value NSTokenStyle) {
	objc.Send[struct{}](t.ID, objc.Sel("setTokenStyle:"), value)
}

// The receiver’s tokenizing character set to a given character set.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/tokenizingCharacterSet
func (t NSTokenFieldCell) TokenizingCharacterSet() foundation.NSCharacterSet {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tokenizingCharacterSet"))
	return foundation.NSCharacterSetFromID(objc.ID(rv))
}
func (t NSTokenFieldCell) SetTokenizingCharacterSet(value foundation.NSCharacterSet) {
	objc.Send[struct{}](t.ID, objc.Sel("setTokenizingCharacterSet:"), value)
}

// The receiver’s completion delay to a given delay.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/completionDelay
func (t NSTokenFieldCell) CompletionDelay() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("completionDelay"))
	return rv
}
func (t NSTokenFieldCell) SetCompletionDelay(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setCompletionDelay:"), value)
}

// The receiver’s delegate.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/delegate
func (t NSTokenFieldCell) Delegate() NSTokenFieldCellDelegate {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("delegate"))
	return NSTokenFieldCellDelegateObjectFromID(rv)
}
func (t NSTokenFieldCell) SetDelegate(value NSTokenFieldCellDelegate) {
	objc.Send[struct{}](t.ID, objc.Sel("setDelegate:"), value)
}

// Returns the default tokenizing character set.
//
// # Return Value
//
// The default tokenizing character set.
//
// # Discussion
//
// The default tokenizing character set contains the single character
// “`,`”.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/defaultTokenizingCharacterSet
func (_NSTokenFieldCellClass NSTokenFieldCellClass) DefaultTokenizingCharacterSet() foundation.NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSTokenFieldCellClass.class), objc.Sel("defaultTokenizingCharacterSet"))
	return foundation.NSCharacterSetFromID(objc.ID(rv))
}

// Returns the default completion delay.
//
// # Return Value
//
// The default completion delay.
//
// # Discussion
//
// The default completion delay is `0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenFieldCell/defaultCompletionDelay
func (_NSTokenFieldCellClass NSTokenFieldCellClass) DefaultCompletionDelay() float64 {
	rv := objc.Send[float64](objc.ID(_NSTokenFieldCellClass.class), objc.Sel("defaultCompletionDelay"))
	return rv
}
