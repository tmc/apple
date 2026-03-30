// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSearchFieldCell] class.
var (
	_NSSearchFieldCellClass     NSSearchFieldCellClass
	_NSSearchFieldCellClassOnce sync.Once
)

func getNSSearchFieldCellClass() NSSearchFieldCellClass {
	_NSSearchFieldCellClassOnce.Do(func() {
		_NSSearchFieldCellClass = NSSearchFieldCellClass{class: objc.GetClass("NSSearchFieldCell")}
	})
	return _NSSearchFieldCellClass
}

// GetNSSearchFieldCellClass returns the class object for NSSearchFieldCell.
func GetNSSearchFieldCellClass() NSSearchFieldCellClass {
	return getNSSearchFieldCellClass()
}

type NSSearchFieldCellClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSearchFieldCellClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSearchFieldCellClass) Alloc() NSSearchFieldCell {
	rv := objc.Send[NSSearchFieldCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The programmatic interface for text fields that are used for text-based
// searches.
//
// # Overview
//
// The [NSSearchFieldCell] class defines the programmatic interface for text
// fields that are optimized for text-based searches. An [NSSearchFieldCell]
// object is “wrapped” by an [NSSearchField] control object, which
// directly inherits from the [NSTextField] class. The search field
// implemented by these classes presents a standard user interface for
// searches, including a search button, a cancel button, and a pop-up icon
// menu for listing recent search strings and custom search categories.
//
// When the user types and then pauses, the cell’s action message is sent to
// its target. You can query the cell’s string value for the current text to
// search for. Do not rely on the sender of the action to be an [NSMenu]
// object because the menu may change. If you need to change the menu, modify
// the search menu template and update the value in the [NSSearchFieldCell.SearchMenuTemplate]
// property.
//
// # Managing buttons
//
//   - [NSSearchFieldCell.SearchButtonCell]: The button cell used to display the search-button image.
//   - [NSSearchFieldCell.SetSearchButtonCell]
//   - [NSSearchFieldCell.ResetSearchButtonCell]: Resets the search button cell to its default attributes.
//   - [NSSearchFieldCell.CancelButtonCell]: The button cell used to display the cancel-button image.
//   - [NSSearchFieldCell.SetCancelButtonCell]
//   - [NSSearchFieldCell.ResetCancelButtonCell]: Resets the cancel button cell to its default attributes.
//
// # Custom layout
//
//   - [NSSearchFieldCell.SearchTextRectForBounds]: Modifies the bounding rectangle for the search-text field cell.
//   - [NSSearchFieldCell.SearchButtonRectForBounds]: Modifies the bounding rectangle for the search button cell.
//   - [NSSearchFieldCell.CancelButtonRectForBounds]: Modifies the bounding rectangle for the cancel button cell.
//
// # Managing menu templates
//
//   - [NSSearchFieldCell.SearchMenuTemplate]: The menu object used to dynamically construct the search field’s pop-up icon menu.
//   - [NSSearchFieldCell.SetSearchMenuTemplate]
//
// # Managing search modes
//
//   - [NSSearchFieldCell.SendsWholeSearchString]: A Boolean value indicating whether the cell calls its search action method when the user clicks the search button (or presses Return) or after each keystroke.
//   - [NSSearchFieldCell.SetSendsWholeSearchString]
//   - [NSSearchFieldCell.SendsSearchStringImmediately]: A Boolean value indicating whether the cell calls its action method immediately when an appropriate action occurs.
//   - [NSSearchFieldCell.SetSendsSearchStringImmediately]
//
// # Managing recent search strings
//
//   - [NSSearchFieldCell.MaximumRecents]: The maximum number of search strings that can appear in the search menu.
//   - [NSSearchFieldCell.SetMaximumRecents]
//   - [NSSearchFieldCell.RecentSearches]: An array of the recent search strings to display in the pop-up icon menu of the search field.
//   - [NSSearchFieldCell.SetRecentSearches]
//   - [NSSearchFieldCell.RecentsAutosaveName]: The autosave name under which the search field automatically saves the list of recent search strings.
//   - [NSSearchFieldCell.SetRecentsAutosaveName]
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell
type NSSearchFieldCell struct {
	NSTextFieldCell
}

// NSSearchFieldCellFromID constructs a [NSSearchFieldCell] from an objc.ID.
//
// The programmatic interface for text fields that are used for text-based
// searches.
func NSSearchFieldCellFromID(id objc.ID) NSSearchFieldCell {
	return NSSearchFieldCell{NSTextFieldCell: NSTextFieldCellFromID(id)}
}

// NOTE: NSSearchFieldCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSearchFieldCell] class.
//
// # Managing buttons
//
//   - [INSSearchFieldCell.SearchButtonCell]: The button cell used to display the search-button image.
//   - [INSSearchFieldCell.SetSearchButtonCell]
//   - [INSSearchFieldCell.ResetSearchButtonCell]: Resets the search button cell to its default attributes.
//   - [INSSearchFieldCell.CancelButtonCell]: The button cell used to display the cancel-button image.
//   - [INSSearchFieldCell.SetCancelButtonCell]
//   - [INSSearchFieldCell.ResetCancelButtonCell]: Resets the cancel button cell to its default attributes.
//
// # Custom layout
//
//   - [INSSearchFieldCell.SearchTextRectForBounds]: Modifies the bounding rectangle for the search-text field cell.
//   - [INSSearchFieldCell.SearchButtonRectForBounds]: Modifies the bounding rectangle for the search button cell.
//   - [INSSearchFieldCell.CancelButtonRectForBounds]: Modifies the bounding rectangle for the cancel button cell.
//
// # Managing menu templates
//
//   - [INSSearchFieldCell.SearchMenuTemplate]: The menu object used to dynamically construct the search field’s pop-up icon menu.
//   - [INSSearchFieldCell.SetSearchMenuTemplate]
//
// # Managing search modes
//
//   - [INSSearchFieldCell.SendsWholeSearchString]: A Boolean value indicating whether the cell calls its search action method when the user clicks the search button (or presses Return) or after each keystroke.
//   - [INSSearchFieldCell.SetSendsWholeSearchString]
//   - [INSSearchFieldCell.SendsSearchStringImmediately]: A Boolean value indicating whether the cell calls its action method immediately when an appropriate action occurs.
//   - [INSSearchFieldCell.SetSendsSearchStringImmediately]
//
// # Managing recent search strings
//
//   - [INSSearchFieldCell.MaximumRecents]: The maximum number of search strings that can appear in the search menu.
//   - [INSSearchFieldCell.SetMaximumRecents]
//   - [INSSearchFieldCell.RecentSearches]: An array of the recent search strings to display in the pop-up icon menu of the search field.
//   - [INSSearchFieldCell.SetRecentSearches]
//   - [INSSearchFieldCell.RecentsAutosaveName]: The autosave name under which the search field automatically saves the list of recent search strings.
//   - [INSSearchFieldCell.SetRecentsAutosaveName]
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell
type INSSearchFieldCell interface {
	INSTextFieldCell

	// Topic: Managing buttons

	// The button cell used to display the search-button image.
	SearchButtonCell() INSButtonCell
	SetSearchButtonCell(value INSButtonCell)
	// Resets the search button cell to its default attributes.
	ResetSearchButtonCell()
	// The button cell used to display the cancel-button image.
	CancelButtonCell() INSButtonCell
	SetCancelButtonCell(value INSButtonCell)
	// Resets the cancel button cell to its default attributes.
	ResetCancelButtonCell()

	// Topic: Custom layout

	// Modifies the bounding rectangle for the search-text field cell.
	SearchTextRectForBounds(rect corefoundation.CGRect) corefoundation.CGRect
	// Modifies the bounding rectangle for the search button cell.
	SearchButtonRectForBounds(rect corefoundation.CGRect) corefoundation.CGRect
	// Modifies the bounding rectangle for the cancel button cell.
	CancelButtonRectForBounds(rect corefoundation.CGRect) corefoundation.CGRect

	// Topic: Managing menu templates

	// The menu object used to dynamically construct the search field’s pop-up icon menu.
	SearchMenuTemplate() INSMenu
	SetSearchMenuTemplate(value INSMenu)

	// Topic: Managing search modes

	// A Boolean value indicating whether the cell calls its search action method when the user clicks the search button (or presses Return) or after each keystroke.
	SendsWholeSearchString() bool
	SetSendsWholeSearchString(value bool)
	// A Boolean value indicating whether the cell calls its action method immediately when an appropriate action occurs.
	SendsSearchStringImmediately() bool
	SetSendsSearchStringImmediately(value bool)

	// Topic: Managing recent search strings

	// The maximum number of search strings that can appear in the search menu.
	MaximumRecents() int
	SetMaximumRecents(value int)
	// An array of the recent search strings to display in the pop-up icon menu of the search field.
	RecentSearches() []string
	SetRecentSearches(value []string)
	// The autosave name under which the search field automatically saves the list of recent search strings.
	RecentsAutosaveName() NSSearchFieldRecentsAutosaveName
	SetRecentsAutosaveName(value NSSearchFieldRecentsAutosaveName)
}

// Init initializes the instance.
func (s NSSearchFieldCell) Init() NSSearchFieldCell {
	rv := objc.Send[NSSearchFieldCell](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSearchFieldCell) Autorelease() NSSearchFieldCell {
	rv := objc.Send[NSSearchFieldCell](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSearchFieldCell creates a new NSSearchFieldCell instance.
func NewNSSearchFieldCell() NSSearchFieldCell {
	class := getNSSearchFieldCellClass()
	rv := objc.Send[NSSearchFieldCell](objc.ID(class.class), objc.Sel("new"))
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
func NewSearchFieldCellImageCell(image INSImage) NSSearchFieldCell {
	instance := getNSSearchFieldCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSSearchFieldCellFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/init(textCell:)
func NewSearchFieldCellTextCell(string_ string) NSSearchFieldCell {
	instance := getNSSearchFieldCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSSearchFieldCellFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/init(coder:)
func NewSearchFieldCellWithCoder(coder foundation.INSCoder) NSSearchFieldCell {
	instance := getNSSearchFieldCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSSearchFieldCellFromID(rv)
}

// Resets the search button cell to its default attributes.
//
// # Discussion
//
// This method resets the target, action, regular image, and pressed image for
// the search button cell. By default, when users click the search button or
// press the Return key, the action defined for the receiver is sent to its
// designated target. This method gives you a way to customize the search
// button for specific situations and then reset the button defaults without
// having to undo changes individually.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/resetSearchButtonCell()
func (s NSSearchFieldCell) ResetSearchButtonCell() {
	objc.Send[objc.ID](s.ID, objc.Sel("resetSearchButtonCell"))
}

// Resets the cancel button cell to its default attributes.
//
// # Discussion
//
// This method resets the target, action, regular image, and pressed image for
// the cancel button cell. By default, when users click the cancel button, the
// “ action message is sent up the responder chain to the first [NSText]
// object that can handle it. This method gives you a way to customize the
// cancel button for specific situations and then reset the button defaults
// without having to undo changes individually.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/resetCancelButtonCell()
func (s NSSearchFieldCell) ResetCancelButtonCell() {
	objc.Send[objc.ID](s.ID, objc.Sel("resetCancelButtonCell"))
}

// Modifies the bounding rectangle for the search-text field cell.
//
// rect: The current bounding rectangle for the search text field.
//
// # Return Value
//
// The updated bounding rectangle to use for the search text field. The
// default value is the value passed into the `rect` parameter.
//
// # Discussion
//
// Subclasses can override this method to return a new bounding rectangle for
// the text-field cell object. You might use this method to provide a custom
// layout for the search field control.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/searchTextRect(forBounds:)
func (s NSSearchFieldCell) SearchTextRectForBounds(rect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("searchTextRectForBounds:"), rect)
	return corefoundation.CGRect(rv)
}

// Modifies the bounding rectangle for the search button cell.
//
// rect: The current bounding rectangle for the search button.
//
// # Return Value
//
// The updated bounding rectangle to use for the search button. The default
// value is the value passed into the `rect` parameter.
//
// # Discussion
//
// Subclasses can override this method to return a new bounding rectangle for
// the search button cell. You might use this method to provide a custom
// layout for the search field control.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/searchButtonRect(forBounds:)
func (s NSSearchFieldCell) SearchButtonRectForBounds(rect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("searchButtonRectForBounds:"), rect)
	return corefoundation.CGRect(rv)
}

// Modifies the bounding rectangle for the cancel button cell.
//
// rect: The current bounding rectangle for the cancel button.
//
// # Return Value
//
// The updated bounding rectangle to use for the cancel button. The default
// value is the value passed into the `rect` parameter.
//
// # Discussion
//
// Subclasses can override this method to return a new bounding rectangle for
// the cancel button cell. You might use this method to provide a custom
// layout for the search field control.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/cancelButtonRect(forBounds:)
func (s NSSearchFieldCell) CancelButtonRectForBounds(rect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("cancelButtonRectForBounds:"), rect)
	return corefoundation.CGRect(rv)
}

// The button cell used to display the search-button image.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/searchButtonCell
func (s NSSearchFieldCell) SearchButtonCell() INSButtonCell {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("searchButtonCell"))
	return NSButtonCellFromID(objc.ID(rv))
}
func (s NSSearchFieldCell) SetSearchButtonCell(value INSButtonCell) {
	objc.Send[struct{}](s.ID, objc.Sel("setSearchButtonCell:"), value)
}

// The button cell used to display the cancel-button image.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/cancelButtonCell
func (s NSSearchFieldCell) CancelButtonCell() INSButtonCell {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("cancelButtonCell"))
	return NSButtonCellFromID(objc.ID(rv))
}
func (s NSSearchFieldCell) SetCancelButtonCell(value INSButtonCell) {
	objc.Send[struct{}](s.ID, objc.Sel("setCancelButtonCell:"), value)
}

// The menu object used to dynamically construct the search field’s pop-up
// icon menu.
//
// # Discussion
//
// The cell looks for the tag constants described in [Menu tags] to determine
// how to populate the menu with items related to recent searches. For an
// example of how you might set up the search menu template, see [Configuring
// a Search Menu].
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/searchMenuTemplate
//
// [Configuring a Search Menu]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/SearchFields/Articles/MenuTemplate.html#//apple_ref/doc/uid/20002245
// [Menu tags]: https://developer.apple.com/documentation/AppKit/menu-tags
func (s NSSearchFieldCell) SearchMenuTemplate() INSMenu {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("searchMenuTemplate"))
	return NSMenuFromID(objc.ID(rv))
}
func (s NSSearchFieldCell) SetSearchMenuTemplate(value INSMenu) {
	objc.Send[struct{}](s.ID, objc.Sel("setSearchMenuTemplate:"), value)
}

// A Boolean value indicating whether the cell calls its search action method
// when the user clicks the search button (or presses Return) or after each
// keystroke.
//
// # Discussion
//
// When the value of this property is true, the cell calls its action method
// when the user clicks the search button or presses Return. When the value is
// false, the cell calls the action method after each keystroke. The default
// value of this property is false.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/sendsWholeSearchString
func (s NSSearchFieldCell) SendsWholeSearchString() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("sendsWholeSearchString"))
	return rv
}
func (s NSSearchFieldCell) SetSendsWholeSearchString(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setSendsWholeSearchString:"), value)
}

// A Boolean value indicating whether the cell calls its action method
// immediately when an appropriate action occurs.
//
// # Discussion
//
// When the value of this property is true, the cell calls its action method
// immediately upon notification of any changes to the search field. When the
// value is false, the cell pauses briefly after receiving a notification and
// then calls its action method. Pausing gives the user an opportunity to type
// more text into the search field and minimize the number of searches that
// are performed.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/sendsSearchStringImmediately
func (s NSSearchFieldCell) SendsSearchStringImmediately() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("sendsSearchStringImmediately"))
	return rv
}
func (s NSSearchFieldCell) SetSendsSearchStringImmediately(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setSendsSearchStringImmediately:"), value)
}

// The maximum number of search strings that can appear in the search menu.
//
// # Discussion
//
// The value of this property must be between `0` and `254`. Specifying a
// negative value for the property sets it to the default value, which is
// `10`. Specifying a value greater than `254` sets the property to `254`.
//
// When the maximum number of search strings is exceeded, the oldest search
// string on the menu is dropped.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/maximumRecents
func (s NSSearchFieldCell) MaximumRecents() int {
	rv := objc.Send[int](s.ID, objc.Sel("maximumRecents"))
	return rv
}
func (s NSSearchFieldCell) SetMaximumRecents(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setMaximumRecents:"), value)
}

// An array of the recent search strings to display in the pop-up icon menu of
// the search field.
//
// # Discussion
//
// This property contains an array of [NSString] objects, each of which
// contains a search string either displayed in the search menu or from a
// recent autosave archive. If there have been no recent searches and no prior
// searches saved under an autosave name, this array may be empty. When
// loading your interface, you might set the value of this property to a set
// of saved search strings.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/recentSearches
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
func (s NSSearchFieldCell) RecentSearches() []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("recentSearches"))
	return objc.ConvertSliceToStrings(rv)
}
func (s NSSearchFieldCell) SetRecentSearches(value []string) {
	objc.Send[struct{}](s.ID, objc.Sel("setRecentSearches:"), objectivec.StringSliceToNSArray(value))
}

// The autosave name under which the search field automatically saves the list
// of recent search strings.
//
// # Discussion
//
// The autosave name is used as a key in the standard user defaults to save
// the recent searches. If you specify `nil` or an empty string for this
// parameter, no autosave name is set and searches are not automatically
// saved.
//
// See: https://developer.apple.com/documentation/AppKit/NSSearchFieldCell/recentsAutosaveName
func (s NSSearchFieldCell) RecentsAutosaveName() NSSearchFieldRecentsAutosaveName {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("recentsAutosaveName"))
	return NSSearchFieldRecentsAutosaveName(foundation.NSStringFromID(rv).String())
}
func (s NSSearchFieldCell) SetRecentsAutosaveName(value NSSearchFieldRecentsAutosaveName) {
	objc.Send[struct{}](s.ID, objc.Sel("setRecentsAutosaveName:"), objc.String(string(value)))
}
