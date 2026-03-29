// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSearchField] class.
var (
	_NSSearchFieldClass     NSSearchFieldClass
	_NSSearchFieldClassOnce sync.Once
)

func getNSSearchFieldClass() NSSearchFieldClass {
	_NSSearchFieldClassOnce.Do(func() {
		_NSSearchFieldClass = NSSearchFieldClass{class: objc.GetClass("NSSearchField")}
	})
	return _NSSearchFieldClass
}

// GetNSSearchFieldClass returns the class object for NSSearchField.
func GetNSSearchFieldClass() NSSearchFieldClass {
	return getNSSearchFieldClass()
}

type NSSearchFieldClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSearchFieldClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSearchFieldClass) Alloc() NSSearchField {
	rv := objc.Send[NSSearchField](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A text field optimized for performing text-based searches.
//
// # Overview
// 
// [NSSearchField] provides a customized text field for entering search data.
// The class also provides a search button, a cancel button, and a pop-up icon
// menu for listing recent search strings and custom search categories.
// 
// An [NSSearchField] object wraps an [NSSearchFieldCell] object. The cell
// provides access to most search field attributes and a comprehensive
// programmatic interface for manipulating the search field. You can use an
// [NSSearchField] object to manipulate some aspects of the search field.
// 
// For additional information about search fields and how to implement them,
// see the [NSSearchFieldCell] class.
//
// # Managing Menu Templates
//
//   - [NSSearchField.SearchMenuTemplate]: The menu object used to dynamically construct the search field’s pop-up icon menu.
//   - [NSSearchField.SetSearchMenuTemplate]
//
// # Managing Search Modes
//
//   - [NSSearchField.SendsSearchStringImmediately]: A Boolean value indicating whether the cell calls its action method immediately when an appropriate action occurs.
//   - [NSSearchField.SetSendsSearchStringImmediately]
//   - [NSSearchField.SendsWholeSearchString]: A Boolean value indicating whether the cell calls its search action method when the user clicks the search button or presses Return, or after each keystroke.
//   - [NSSearchField.SetSendsWholeSearchString]
//
// # Managing Recent Searches
//
//   - [NSSearchField.RecentSearches]: The list of recent search strings for the control.
//   - [NSSearchField.SetRecentSearches]
//   - [NSSearchField.MaximumRecents]: The maximum number of search strings that can appear in the search menu.
//   - [NSSearchField.SetMaximumRecents]
//   - [NSSearchField.RecentsAutosaveName]: The name under which the search field automatically archives the list of recent search strings.
//   - [NSSearchField.SetRecentsAutosaveName]
//
// # Getting Search Field Metrics
//
//   - [NSSearchField.CancelButtonBounds]: The rectangle for the cancel button within the bounds of the search field.
//   - [NSSearchField.SearchButtonBounds]: The rectangle for the search button within the bounds of the search field.
//   - [NSSearchField.SearchTextBounds]: The rectangle for the search text within the bounds of the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchField
type NSSearchField struct {
	NSTextField
}

// NSSearchFieldFromID constructs a [NSSearchField] from an objc.ID.
//
// A text field optimized for performing text-based searches.
func NSSearchFieldFromID(id objc.ID) NSSearchField {
	return NSSearchField{NSTextField: NSTextFieldFromID(id)}
}
// NOTE: NSSearchField adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSearchField] class.
//
// # Managing Menu Templates
//
//   - [INSSearchField.SearchMenuTemplate]: The menu object used to dynamically construct the search field’s pop-up icon menu.
//   - [INSSearchField.SetSearchMenuTemplate]
//
// # Managing Search Modes
//
//   - [INSSearchField.SendsSearchStringImmediately]: A Boolean value indicating whether the cell calls its action method immediately when an appropriate action occurs.
//   - [INSSearchField.SetSendsSearchStringImmediately]
//   - [INSSearchField.SendsWholeSearchString]: A Boolean value indicating whether the cell calls its search action method when the user clicks the search button or presses Return, or after each keystroke.
//   - [INSSearchField.SetSendsWholeSearchString]
//
// # Managing Recent Searches
//
//   - [INSSearchField.RecentSearches]: The list of recent search strings for the control.
//   - [INSSearchField.SetRecentSearches]
//   - [INSSearchField.MaximumRecents]: The maximum number of search strings that can appear in the search menu.
//   - [INSSearchField.SetMaximumRecents]
//   - [INSSearchField.RecentsAutosaveName]: The name under which the search field automatically archives the list of recent search strings.
//   - [INSSearchField.SetRecentsAutosaveName]
//
// # Getting Search Field Metrics
//
//   - [INSSearchField.CancelButtonBounds]: The rectangle for the cancel button within the bounds of the search field.
//   - [INSSearchField.SearchButtonBounds]: The rectangle for the search button within the bounds of the search field.
//   - [INSSearchField.SearchTextBounds]: The rectangle for the search text within the bounds of the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchField
type INSSearchField interface {
	INSTextField

	// Topic: Managing Menu Templates

	// The menu object used to dynamically construct the search field’s pop-up icon menu.
	SearchMenuTemplate() INSMenu
	SetSearchMenuTemplate(value INSMenu)

	// Topic: Managing Search Modes

	// A Boolean value indicating whether the cell calls its action method immediately when an appropriate action occurs.
	SendsSearchStringImmediately() bool
	SetSendsSearchStringImmediately(value bool)
	// A Boolean value indicating whether the cell calls its search action method when the user clicks the search button or presses Return, or after each keystroke.
	SendsWholeSearchString() bool
	SetSendsWholeSearchString(value bool)

	// Topic: Managing Recent Searches

	// The list of recent search strings for the control.
	RecentSearches() []string
	SetRecentSearches(value []string)
	// The maximum number of search strings that can appear in the search menu.
	MaximumRecents() int
	SetMaximumRecents(value int)
	// The name under which the search field automatically archives the list of recent search strings.
	RecentsAutosaveName() NSSearchFieldRecentsAutosaveName
	SetRecentsAutosaveName(value NSSearchFieldRecentsAutosaveName)

	// Topic: Getting Search Field Metrics

	// The rectangle for the cancel button within the bounds of the search field.
	CancelButtonBounds() corefoundation.CGRect
	// The rectangle for the search button within the bounds of the search field.
	SearchButtonBounds() corefoundation.CGRect
	// The rectangle for the search text within the bounds of the search field.
	SearchTextBounds() corefoundation.CGRect
}

// Init initializes the instance.
func (s NSSearchField) Init() NSSearchField {
	rv := objc.Send[NSSearchField](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSearchField) Autorelease() NSSearchField {
	rv := objc.Send[NSSearchField](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSearchField creates a new NSSearchField instance.
func NewNSSearchField() NSSearchField {
	class := getNSSearchFieldClass()
	rv := objc.Send[NSSearchField](objc.ID(class.class), objc.Sel("new"))
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
func NewSearchFieldLabelWithAttributedString(attributedStringValue foundation.NSAttributedString) NSSearchField {
	rv := objc.Send[objc.ID](objc.ID(getNSSearchFieldClass().class), objc.Sel("labelWithAttributedString:"), attributedStringValue)
	return NSSearchFieldFromID(rv)
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
func NewSearchFieldLabelWithString(stringValue string) NSSearchField {
	rv := objc.Send[objc.ID](objc.ID(getNSSearchFieldClass().class), objc.Sel("labelWithString:"), objc.String(stringValue))
	return NSSearchFieldFromID(rv)
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
func NewSearchFieldTextFieldWithString(stringValue string) NSSearchField {
	rv := objc.Send[objc.ID](objc.ID(getNSSearchFieldClass().class), objc.Sel("textFieldWithString:"), objc.String(stringValue))
	return NSSearchFieldFromID(rv)
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewSearchFieldWithCoder(coder foundation.INSCoder) NSSearchField {
	instance := getNSSearchFieldClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSSearchFieldFromID(rv)
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
func NewSearchFieldWithFrame(frameRect corefoundation.CGRect) NSSearchField {
	instance := getNSSearchFieldClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSSearchFieldFromID(rv)
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
func NewSearchFieldWrappingLabelWithString(stringValue string) NSSearchField {
	rv := objc.Send[objc.ID](objc.ID(getNSSearchFieldClass().class), objc.Sel("wrappingLabelWithString:"), objc.String(stringValue))
	return NSSearchFieldFromID(rv)
}

// The menu object used to dynamically construct the search field’s pop-up
// icon menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchField/searchMenuTemplate
func (s NSSearchField) SearchMenuTemplate() INSMenu {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("searchMenuTemplate"))
	return NSMenuFromID(objc.ID(rv))
}
func (s NSSearchField) SetSearchMenuTemplate(value INSMenu) {
	objc.Send[struct{}](s.ID, objc.Sel("setSearchMenuTemplate:"), value)
}
// A Boolean value indicating whether the cell calls its action method
// immediately when an appropriate action occurs.
//
// # Discussion
// 
// When the value of this property is YES, the field calls its action method
// immediately upon notification of any changes to the search field. When the
// value is NO, the field pauses briefly after receiving a notification and
// then calls its action method. Pausing gives the user an opportunity to type
// more text into the search field and minimize the number of searches that
// are performed.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchField/sendsSearchStringImmediately
func (s NSSearchField) SendsSearchStringImmediately() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("sendsSearchStringImmediately"))
	return rv
}
func (s NSSearchField) SetSendsSearchStringImmediately(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setSendsSearchStringImmediately:"), value)
}
// A Boolean value indicating whether the cell calls its search action method
// when the user clicks the search button or presses Return, or after each
// keystroke.
//
// # Discussion
// 
// When the value of this property is yes, the field calls its action method
// when the user clicks the search button or presses Return. When the value is
// NO, the field calls the action method after each keystroke. The default
// value of this property is no.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchField/sendsWholeSearchString
func (s NSSearchField) SendsWholeSearchString() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("sendsWholeSearchString"))
	return rv
}
func (s NSSearchField) SetSendsWholeSearchString(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setSendsWholeSearchString:"), value)
}
// The list of recent search strings for the control.
//
// # Discussion
// 
// An array of [NSString] objects, each of which contains a search string
// either displayed in the search menu or from a recent autosave archive. If
// there have been no recent searches and no prior searches saved under an
// autosave name, this array may be empty.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchField/recentSearches
func (s NSSearchField) RecentSearches() []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("recentSearches"))
	return objc.ConvertSliceToStrings(rv)
}
func (s NSSearchField) SetRecentSearches(value []string) {
	objc.Send[struct{}](s.ID, objc.Sel("setRecentSearches:"), objectivec.StringSliceToNSArray(value))
}
// The maximum number of search strings that can appear in the search menu.
//
// # Discussion
// 
// The value of this property must be between 0 and 254. Specifying a negative
// value for the property sets it to the default value, which is 10.
// Specifying a value greater than 254 sets the property to 254.
// 
// When the maximum number of search strings is exceeded, the oldest search
// string on the menu is dropped.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchField/maximumRecents
func (s NSSearchField) MaximumRecents() int {
	rv := objc.Send[int](s.ID, objc.Sel("maximumRecents"))
	return rv
}
func (s NSSearchField) SetMaximumRecents(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setMaximumRecents:"), value)
}
// The name under which the search field automatically archives the list of
// recent search strings.
//
// # Discussion
// 
// Used as a key in the standard user defaults to save the recent searches. If
// you specify `nil` or an empty string for this property, no autosave name is
// set and searches aren’t autosaved.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchField/recentsAutosaveName-swift.property
func (s NSSearchField) RecentsAutosaveName() NSSearchFieldRecentsAutosaveName {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("recentsAutosaveName"))
	return NSSearchFieldRecentsAutosaveName(foundation.NSStringFromID(rv).String())
}
func (s NSSearchField) SetRecentsAutosaveName(value NSSearchFieldRecentsAutosaveName) {
	objc.Send[struct{}](s.ID, objc.Sel("setRecentsAutosaveName:"), objc.String(string(value)))
}
// The rectangle for the cancel button within the bounds of the search field.
//
// # Discussion
// 
// Subclasses can override `cancelButtonBounds` for custom layout purposes.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchField/cancelButtonBounds
func (s NSSearchField) CancelButtonBounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("cancelButtonBounds"))
	return corefoundation.CGRect(rv)
}
// The rectangle for the search button within the bounds of the search field.
//
// # Discussion
// 
// Subclasses can override `searchButtonBounds` for custom layout purposes.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchField/searchButtonBounds
func (s NSSearchField) SearchButtonBounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("searchButtonBounds"))
	return corefoundation.CGRect(rv)
}
// The rectangle for the search text within the bounds of the search field.
//
// # Discussion
// 
// Subclasses can override `searchTextBounds` for custom layout purposes.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchField/searchTextBounds
func (s NSSearchField) SearchTextBounds() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("searchTextBounds"))
	return corefoundation.CGRect(rv)
}

