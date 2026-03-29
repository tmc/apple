// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSSecureTextField] class.
var (
	_NSSecureTextFieldClass     NSSecureTextFieldClass
	_NSSecureTextFieldClassOnce sync.Once
)

func getNSSecureTextFieldClass() NSSecureTextFieldClass {
	_NSSecureTextFieldClassOnce.Do(func() {
		_NSSecureTextFieldClass = NSSecureTextFieldClass{class: objc.GetClass("NSSecureTextField")}
	})
	return _NSSecureTextFieldClass
}

// GetNSSecureTextFieldClass returns the class object for NSSecureTextField.
func GetNSSecureTextFieldClass() NSSecureTextFieldClass {
	return getNSSecureTextFieldClass()
}

type NSSecureTextFieldClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSecureTextFieldClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSecureTextFieldClass) Alloc() NSSecureTextField {
	rv := objc.Send[NSSecureTextField](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A text field that hides the typed text.
//
// # Overview
// 
// A secure text field is suitable for use as a password-entry object or for
// any item in which the text value must be kept secret. [NSSecureTextField]
// uses [NSSecureTextFieldCell] to implement its user interface.
//
// See: https://developer.apple.com/documentation/AppKit/NSSecureTextField
type NSSecureTextField struct {
	NSTextField
}

// NSSecureTextFieldFromID constructs a [NSSecureTextField] from an objc.ID.
//
// A text field that hides the typed text.
func NSSecureTextFieldFromID(id objc.ID) NSSecureTextField {
	return NSSecureTextField{NSTextField: NSTextFieldFromID(id)}
}
// NOTE: NSSecureTextField adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSecureTextField] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSSecureTextField
type INSSecureTextField interface {
	INSTextField
}

// Init initializes the instance.
func (s NSSecureTextField) Init() NSSecureTextField {
	rv := objc.Send[NSSecureTextField](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSecureTextField) Autorelease() NSSecureTextField {
	rv := objc.Send[NSSecureTextField](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSecureTextField creates a new NSSecureTextField instance.
func NewNSSecureTextField() NSSecureTextField {
	class := getNSSecureTextFieldClass()
	rv := objc.Send[NSSecureTextField](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a text field for use as a static label that displays styled text,
// doesn’t wrap, and doesn’t have selectable text.
//
// attributedStringValue: An attributed string to use as the content of the label.
//
// # Return Value
// 
// A text field that displays the specified attributed string as a static
// label.
//
// # Discussion
// 
// The text field determines its line-break mode by inspecting the paragraph
// style attributes in the attributed string.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/init(labelWithAttributedString:)
func NewSecureTextFieldLabelWithAttributedString(attributedStringValue foundation.NSAttributedString) NSSecureTextField {
	rv := objc.Send[objc.ID](objc.ID(getNSSecureTextFieldClass().class), objc.Sel("labelWithAttributedString:"), attributedStringValue)
	return NSSecureTextFieldFromID(rv)
}

// Initializes a text field for use as a static label that uses the system
// default font, doesn’t wrap, and doesn’t have selectable text.
//
// stringValue: A string to use as the content of the label.
//
// # Return Value
// 
// A text field that displays the specified string as a static label.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/init(labelWithString:)
func NewSecureTextFieldLabelWithString(stringValue string) NSSecureTextField {
	rv := objc.Send[objc.ID](objc.ID(getNSSecureTextFieldClass().class), objc.Sel("labelWithString:"), objc.String(stringValue))
	return NSSecureTextFieldFromID(rv)
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewSecureTextFieldWithCoder(coder foundation.INSCoder) NSSecureTextField {
	instance := getNSSecureTextFieldClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSSecureTextFieldFromID(rv)
}

// Initializes a control with the specified frame rectangle.
//
// frameRect: The rectangle of the control, specified in points in the coordinate space
// of the enclosing view.
//
// # Return Value
// 
// An initialized control object, or `nil` if the object couldn’t be
// initialized.
//
// # Discussion
// 
// If a cell has been specified for controls of this type, this method also
// creates an instance of the cell. Because [NSControl] is an abstract class,
// invocations of this method should appear only in the designated
// initializers of subclasses; that is, there should always be a more specific
// designated initializer for the subclass, as this method is the designated
// initializer for [NSControl].
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(frame:)
func NewSecureTextFieldWithFrame(frameRect corefoundation.CGRect) NSSecureTextField {
	instance := getNSSecureTextFieldClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSSecureTextFieldFromID(rv)
}

// Initializes a single-line editable text field for user input using the
// system default font and standard visual appearance.
//
// stringValue: A string to use as the initial content of the editable text field.
//
// # Return Value
// 
// A single-line editable text field that displays the specified string.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/init(string:)
func NewSecureTextFieldWithString(stringValue string) NSSecureTextField {
	rv := objc.Send[objc.ID](objc.ID(getNSSecureTextFieldClass().class), objc.Sel("textFieldWithString:"), objc.String(stringValue))
	return NSSecureTextFieldFromID(rv)
}

// Initializes a text field for use as a multiline static label with
// selectable text that uses the system default font.
//
// stringValue: A string to use as the initial content of the editable text field.
//
// # Return Value
// 
// A multiline text field that displays the specified string.
//
// See: https://developer.apple.com/documentation/AppKit/NSTextField/init(wrappingLabelWithString:)
func NewSecureTextFieldWrappingLabelWithString(stringValue string) NSSecureTextField {
	rv := objc.Send[objc.ID](objc.ID(getNSSecureTextFieldClass().class), objc.Sel("wrappingLabelWithString:"), objc.String(stringValue))
	return NSSecureTextFieldFromID(rv)
}

