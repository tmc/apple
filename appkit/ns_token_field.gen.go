// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSTokenField] class.
var (
	_NSTokenFieldClass     NSTokenFieldClass
	_NSTokenFieldClassOnce sync.Once
)

func getNSTokenFieldClass() NSTokenFieldClass {
	_NSTokenFieldClassOnce.Do(func() {
		_NSTokenFieldClass = NSTokenFieldClass{class: objc.GetClass("NSTokenField")}
	})
	return _NSTokenFieldClass
}

// GetNSTokenFieldClass returns the class object for NSTokenField.
func GetNSTokenFieldClass() NSTokenFieldClass {
	return getNSTokenFieldClass()
}

type NSTokenFieldClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSTokenFieldClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSTokenFieldClass) Alloc() NSTokenField {
	rv := objc.Send[NSTokenField](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A text field that converts text into visually distinct tokens.
//
// # Overview
// 
// Use a token field when you want typed text to be transformed into
// “tokens”, which are visually distinct elements in the text field
// interface. For example, you might use a token field in a mail app to
// display email addresses for individual users. The distinct appearance of
// tokens makes them easy for users to distinguish from surrounding text.
// 
// [NSTokenField] uses an [NSTokenFieldCell] to implement much of the
// control’s functionality. [NSTokenField] provides cover methods for most
// methods of [NSTokenFieldCell], which invoke the corresponding cell method.
//
// # Configuring the Token Style
//
//   - [NSTokenField.TokenStyle]: The token style of the receiver.
//   - [NSTokenField.SetTokenStyle]
//
// # Configuring the Tokenizing Character Set
//
//   - [NSTokenField.TokenizingCharacterSet]: The recevier’s tokenizing character set to `characterSet`.
//   - [NSTokenField.SetTokenizingCharacterSet]
//
// # Configuring the Completion Delay
//
//   - [NSTokenField.CompletionDelay]: The receiver’s completion delay.
//   - [NSTokenField.SetCompletionDelay]
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenField
type NSTokenField struct {
	NSTextField
}

// NSTokenFieldFromID constructs a [NSTokenField] from an objc.ID.
//
// A text field that converts text into visually distinct tokens.
func NSTokenFieldFromID(id objc.ID) NSTokenField {
	return NSTokenField{NSTextField: NSTextFieldFromID(id)}
}
// NOTE: NSTokenField adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSTokenField] class.
//
// # Configuring the Token Style
//
//   - [INSTokenField.TokenStyle]: The token style of the receiver.
//   - [INSTokenField.SetTokenStyle]
//
// # Configuring the Tokenizing Character Set
//
//   - [INSTokenField.TokenizingCharacterSet]: The recevier’s tokenizing character set to `characterSet`.
//   - [INSTokenField.SetTokenizingCharacterSet]
//
// # Configuring the Completion Delay
//
//   - [INSTokenField.CompletionDelay]: The receiver’s completion delay.
//   - [INSTokenField.SetCompletionDelay]
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenField
type INSTokenField interface {
	INSTextField

	// Topic: Configuring the Token Style

	// The token style of the receiver.
	TokenStyle() NSTokenStyle
	SetTokenStyle(value NSTokenStyle)

	// Topic: Configuring the Tokenizing Character Set

	// The recevier’s tokenizing character set to `characterSet`.
	TokenizingCharacterSet() foundation.NSCharacterSet
	SetTokenizingCharacterSet(value foundation.NSCharacterSet)

	// Topic: Configuring the Completion Delay

	// The receiver’s completion delay.
	CompletionDelay() float64
	SetCompletionDelay(value float64)
}

// Init initializes the instance.
func (t NSTokenField) Init() NSTokenField {
	rv := objc.Send[NSTokenField](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSTokenField) Autorelease() NSTokenField {
	rv := objc.Send[NSTokenField](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSTokenField creates a new NSTokenField instance.
func NewNSTokenField() NSTokenField {
	class := getNSTokenFieldClass()
	rv := objc.Send[NSTokenField](objc.ID(class.class), objc.Sel("new"))
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
func NewTokenFieldLabelWithAttributedString(attributedStringValue foundation.NSAttributedString) NSTokenField {
	rv := objc.Send[objc.ID](objc.ID(getNSTokenFieldClass().class), objc.Sel("labelWithAttributedString:"), attributedStringValue)
	return NSTokenFieldFromID(rv)
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
func NewTokenFieldLabelWithString(stringValue string) NSTokenField {
	rv := objc.Send[objc.ID](objc.ID(getNSTokenFieldClass().class), objc.Sel("labelWithString:"), objc.String(stringValue))
	return NSTokenFieldFromID(rv)
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
func NewTokenFieldTextFieldWithString(stringValue string) NSTokenField {
	rv := objc.Send[objc.ID](objc.ID(getNSTokenFieldClass().class), objc.Sel("textFieldWithString:"), objc.String(stringValue))
	return NSTokenFieldFromID(rv)
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewTokenFieldWithCoder(coder foundation.INSCoder) NSTokenField {
	instance := getNSTokenFieldClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSTokenFieldFromID(rv)
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
func NewTokenFieldWithFrame(frameRect corefoundation.CGRect) NSTokenField {
	instance := getNSTokenFieldClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSTokenFieldFromID(rv)
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
func NewTokenFieldWrappingLabelWithString(stringValue string) NSTokenField {
	rv := objc.Send[objc.ID](objc.ID(getNSTokenFieldClass().class), objc.Sel("wrappingLabelWithString:"), objc.String(stringValue))
	return NSTokenFieldFromID(rv)
}

// The token style of the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenField/tokenStyle-swift.property
func (t NSTokenField) TokenStyle() NSTokenStyle {
	rv := objc.Send[NSTokenStyle](t.ID, objc.Sel("tokenStyle"))
	return NSTokenStyle(rv)
}
func (t NSTokenField) SetTokenStyle(value NSTokenStyle) {
	objc.Send[struct{}](t.ID, objc.Sel("setTokenStyle:"), value)
}
// The recevier’s tokenizing character set to `characterSet`.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenField/tokenizingCharacterSet
func (t NSTokenField) TokenizingCharacterSet() foundation.NSCharacterSet {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("tokenizingCharacterSet"))
	return foundation.NSCharacterSetFromID(objc.ID(rv))
}
func (t NSTokenField) SetTokenizingCharacterSet(value foundation.NSCharacterSet) {
	objc.Send[struct{}](t.ID, objc.Sel("setTokenizingCharacterSet:"), value)
}
// The receiver’s completion delay.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenField/completionDelay
func (t NSTokenField) CompletionDelay() float64 {
	rv := objc.Send[float64](t.ID, objc.Sel("completionDelay"))
	return rv
}
func (t NSTokenField) SetCompletionDelay(value float64) {
	objc.Send[struct{}](t.ID, objc.Sel("setCompletionDelay:"), value)
}

// Returns the default tokenizing character set.
//
// # Discussion
// 
// The default tokenizing character set is “,”.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenField/defaultTokenizingCharacterSet
func (_NSTokenFieldClass NSTokenFieldClass) DefaultTokenizingCharacterSet() foundation.NSCharacterSet {
	rv := objc.Send[objc.ID](objc.ID(_NSTokenFieldClass.class), objc.Sel("defaultTokenizingCharacterSet"))
	return foundation.NSCharacterSetFromID(objc.ID(rv))
}
// Returns the default completion delay.
//
// # Discussion
// 
// The default completion delay is 0.
//
// See: https://developer.apple.com/documentation/AppKit/NSTokenField/defaultCompletionDelay
func (_NSTokenFieldClass NSTokenFieldClass) DefaultCompletionDelay() float64 {
	rv := objc.Send[float64](objc.ID(_NSTokenFieldClass.class), objc.Sel("defaultCompletionDelay"))
	return rv
}

