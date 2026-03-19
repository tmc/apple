// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSComboBox] class.
var (
	_NSComboBoxClass     NSComboBoxClass
	_NSComboBoxClassOnce sync.Once
)

func getNSComboBoxClass() NSComboBoxClass {
	_NSComboBoxClassOnce.Do(func() {
		_NSComboBoxClass = NSComboBoxClass{class: objc.GetClass("NSComboBox")}
	})
	return _NSComboBoxClass
}

// GetNSComboBoxClass returns the class object for NSComboBox.
func GetNSComboBoxClass() NSComboBoxClass {
	return getNSComboBoxClass()
}

type NSComboBoxClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSComboBoxClass) Alloc() NSComboBox {
	rv := objc.Send[NSComboBox](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A view that displays a list of values in a pop-up menu where the user
// selects a value or types in a custom value.
//
// # Overview
// 
// A combo box combines the behavior of an [NSTextField] object with an
// [NSPopUpButton] object. A combo box displays a list of values from a pop-up
// list, but also provides a means for users to type in custom values. For
// example, here’s a combo box in its initial state.
// 
// [media-4305420]
// 
// Clicking in the text portion of the control allows the user to edit the
// current value. When the user clicks the down arrow at the right side of the
// text field, the pop-up list appears.
// 
// [media-4305419]
// 
// The [NSComboBox] class uses [NSComboBoxCell] to implement its user
// interface.
// 
// Also see the [NSComboBoxDataSource] protocol, which declares the methods
// that [NSComboBox] uses to access the contents of its data source object.
//
// # Setting Display Attributes
//
//   - [NSComboBox.HasVerticalScroller]: A Boolean value indicating whether the combo box has a vertical scroller.
//   - [NSComboBox.SetHasVerticalScroller]
//   - [NSComboBox.IntercellSpacing]: The horizontal and vertical spacing between cells in the pop-up list.
//   - [NSComboBox.SetIntercellSpacing]
//   - [NSComboBox.ButtonBordered]: A Boolean value indicating whether the combo box displays a border.
//   - [NSComboBox.SetButtonBordered]
//   - [NSComboBox.ItemHeight]: The height of each item in the pop-up list.
//   - [NSComboBox.SetItemHeight]
//   - [NSComboBox.NumberOfVisibleItems]: The maximum number of visible items to display in the pop-up list at one time.
//   - [NSComboBox.SetNumberOfVisibleItems]
//
// # Setting a Data Source
//
//   - [NSComboBox.DataSource]: The object that provides the item data for the combo box.
//   - [NSComboBox.SetDataSource]
//   - [NSComboBox.UsesDataSource]: A Boolean value indicating whether the combo box retrieves its items from a data source object.
//   - [NSComboBox.SetUsesDataSource]
//
// # Configuring the Combo Box Items
//
//   - [NSComboBox.AddItemsWithObjectValues]: Adds multiple objects to the end of the receiver’s internal item list.
//   - [NSComboBox.AddItemWithObjectValue]: Adds an object to the end of the receiver’s internal item list.
//   - [NSComboBox.InsertItemWithObjectValueAtIndex]: Inserts an object at the specified location in the receiver’s internal item list.
//   - [NSComboBox.ObjectValues]: An array of the items from the combo box’s internal list.
//   - [NSComboBox.RemoveAllItems]: Removes all items from the receiver’s internal item list.
//   - [NSComboBox.RemoveItemAtIndex]: Removes the object at the specified location from the receiver’s internal item list.
//   - [NSComboBox.RemoveItemWithObjectValue]: Removes all occurrences of the given object from the receiver’s internal item list.
//   - [NSComboBox.NumberOfItems]: The total number of items in the pop-up list.
//
// # Manipulating the Displayed List
//
//   - [NSComboBox.IndexOfItemWithObjectValue]: Searches the receiver’s internal item list for the specified object and returns the lowest matching index.
//   - [NSComboBox.ItemObjectValueAtIndex]: Returns the object located at the given index within the receiver’s internal item list.
//   - [NSComboBox.NoteNumberOfItemsChanged]: Informs the receiver that the number of items in its data source has changed.
//   - [NSComboBox.ReloadData]: Marks the receiver as needing redisplay, so that it will reload the data for visible pop-up items and draw the new values.
//   - [NSComboBox.ScrollItemAtIndexToTop]: Scrolls the receiver’s pop-up list vertically so that the item at the specified index is as close to the top as possible.
//   - [NSComboBox.ScrollItemAtIndexToVisible]: Scrolls the receiver’s pop-up list vertically so that the item at the specified index is visible.
//
// # Manipulating the Selection
//
//   - [NSComboBox.DeselectItemAtIndex]: Deselects the pop-up list item at the specified index if it’s selected.
//   - [NSComboBox.IndexOfSelectedItem]: The index of the last item selected from the pop-up list.
//   - [NSComboBox.ObjectValueOfSelectedItem]: The object corresponding to the last item selected from the pop-up list.
//   - [NSComboBox.SelectItemAtIndex]: Selects the pop-up list row at the given index.
//   - [NSComboBox.SelectItemWithObjectValue]: Selects the first pop-up list item that corresponds to the given object.
//
// # Completing the Text Field
//
//   - [NSComboBox.Completes]: A Boolean value indicating whether the combo box tries to complete what the user types.
//   - [NSComboBox.SetCompletes]
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox
type NSComboBox struct {
	NSTextField
}

// NSComboBoxFromID constructs a [NSComboBox] from an objc.ID.
//
// A view that displays a list of values in a pop-up menu where the user
// selects a value or types in a custom value.
func NSComboBoxFromID(id objc.ID) NSComboBox {
	return NSComboBox{NSTextField: NSTextFieldFromID(id)}
}
// NOTE: NSComboBox adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSComboBox] class.
//
// # Setting Display Attributes
//
//   - [INSComboBox.HasVerticalScroller]: A Boolean value indicating whether the combo box has a vertical scroller.
//   - [INSComboBox.SetHasVerticalScroller]
//   - [INSComboBox.IntercellSpacing]: The horizontal and vertical spacing between cells in the pop-up list.
//   - [INSComboBox.SetIntercellSpacing]
//   - [INSComboBox.ButtonBordered]: A Boolean value indicating whether the combo box displays a border.
//   - [INSComboBox.SetButtonBordered]
//   - [INSComboBox.ItemHeight]: The height of each item in the pop-up list.
//   - [INSComboBox.SetItemHeight]
//   - [INSComboBox.NumberOfVisibleItems]: The maximum number of visible items to display in the pop-up list at one time.
//   - [INSComboBox.SetNumberOfVisibleItems]
//
// # Setting a Data Source
//
//   - [INSComboBox.DataSource]: The object that provides the item data for the combo box.
//   - [INSComboBox.SetDataSource]
//   - [INSComboBox.UsesDataSource]: A Boolean value indicating whether the combo box retrieves its items from a data source object.
//   - [INSComboBox.SetUsesDataSource]
//
// # Configuring the Combo Box Items
//
//   - [INSComboBox.AddItemsWithObjectValues]: Adds multiple objects to the end of the receiver’s internal item list.
//   - [INSComboBox.AddItemWithObjectValue]: Adds an object to the end of the receiver’s internal item list.
//   - [INSComboBox.InsertItemWithObjectValueAtIndex]: Inserts an object at the specified location in the receiver’s internal item list.
//   - [INSComboBox.ObjectValues]: An array of the items from the combo box’s internal list.
//   - [INSComboBox.RemoveAllItems]: Removes all items from the receiver’s internal item list.
//   - [INSComboBox.RemoveItemAtIndex]: Removes the object at the specified location from the receiver’s internal item list.
//   - [INSComboBox.RemoveItemWithObjectValue]: Removes all occurrences of the given object from the receiver’s internal item list.
//   - [INSComboBox.NumberOfItems]: The total number of items in the pop-up list.
//
// # Manipulating the Displayed List
//
//   - [INSComboBox.IndexOfItemWithObjectValue]: Searches the receiver’s internal item list for the specified object and returns the lowest matching index.
//   - [INSComboBox.ItemObjectValueAtIndex]: Returns the object located at the given index within the receiver’s internal item list.
//   - [INSComboBox.NoteNumberOfItemsChanged]: Informs the receiver that the number of items in its data source has changed.
//   - [INSComboBox.ReloadData]: Marks the receiver as needing redisplay, so that it will reload the data for visible pop-up items and draw the new values.
//   - [INSComboBox.ScrollItemAtIndexToTop]: Scrolls the receiver’s pop-up list vertically so that the item at the specified index is as close to the top as possible.
//   - [INSComboBox.ScrollItemAtIndexToVisible]: Scrolls the receiver’s pop-up list vertically so that the item at the specified index is visible.
//
// # Manipulating the Selection
//
//   - [INSComboBox.DeselectItemAtIndex]: Deselects the pop-up list item at the specified index if it’s selected.
//   - [INSComboBox.IndexOfSelectedItem]: The index of the last item selected from the pop-up list.
//   - [INSComboBox.ObjectValueOfSelectedItem]: The object corresponding to the last item selected from the pop-up list.
//   - [INSComboBox.SelectItemAtIndex]: Selects the pop-up list row at the given index.
//   - [INSComboBox.SelectItemWithObjectValue]: Selects the first pop-up list item that corresponds to the given object.
//
// # Completing the Text Field
//
//   - [INSComboBox.Completes]: A Boolean value indicating whether the combo box tries to complete what the user types.
//   - [INSComboBox.SetCompletes]
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox
type INSComboBox interface {
	INSTextField

	// Topic: Setting Display Attributes

	// A Boolean value indicating whether the combo box has a vertical scroller.
	HasVerticalScroller() bool
	SetHasVerticalScroller(value bool)
	// The horizontal and vertical spacing between cells in the pop-up list.
	IntercellSpacing() corefoundation.CGSize
	SetIntercellSpacing(value corefoundation.CGSize)
	// A Boolean value indicating whether the combo box displays a border.
	ButtonBordered() bool
	SetButtonBordered(value bool)
	// The height of each item in the pop-up list.
	ItemHeight() float64
	SetItemHeight(value float64)
	// The maximum number of visible items to display in the pop-up list at one time.
	NumberOfVisibleItems() int
	SetNumberOfVisibleItems(value int)

	// Topic: Setting a Data Source

	// The object that provides the item data for the combo box.
	DataSource() NSComboBoxDataSource
	SetDataSource(value NSComboBoxDataSource)
	// A Boolean value indicating whether the combo box retrieves its items from a data source object.
	UsesDataSource() bool
	SetUsesDataSource(value bool)

	// Topic: Configuring the Combo Box Items

	// Adds multiple objects to the end of the receiver’s internal item list.
	AddItemsWithObjectValues(objects foundation.INSArray)
	// Adds an object to the end of the receiver’s internal item list.
	AddItemWithObjectValue(object objectivec.IObject)
	// Inserts an object at the specified location in the receiver’s internal item list.
	InsertItemWithObjectValueAtIndex(object objectivec.IObject, index int)
	// An array of the items from the combo box’s internal list.
	ObjectValues() foundation.INSArray
	// Removes all items from the receiver’s internal item list.
	RemoveAllItems()
	// Removes the object at the specified location from the receiver’s internal item list.
	RemoveItemAtIndex(index int)
	// Removes all occurrences of the given object from the receiver’s internal item list.
	RemoveItemWithObjectValue(object objectivec.IObject)
	// The total number of items in the pop-up list.
	NumberOfItems() int

	// Topic: Manipulating the Displayed List

	// Searches the receiver’s internal item list for the specified object and returns the lowest matching index.
	IndexOfItemWithObjectValue(object objectivec.IObject) int
	// Returns the object located at the given index within the receiver’s internal item list.
	ItemObjectValueAtIndex(index int) objectivec.IObject
	// Informs the receiver that the number of items in its data source has changed.
	NoteNumberOfItemsChanged()
	// Marks the receiver as needing redisplay, so that it will reload the data for visible pop-up items and draw the new values.
	ReloadData()
	// Scrolls the receiver’s pop-up list vertically so that the item at the specified index is as close to the top as possible.
	ScrollItemAtIndexToTop(index int)
	// Scrolls the receiver’s pop-up list vertically so that the item at the specified index is visible.
	ScrollItemAtIndexToVisible(index int)

	// Topic: Manipulating the Selection

	// Deselects the pop-up list item at the specified index if it’s selected.
	DeselectItemAtIndex(index int)
	// The index of the last item selected from the pop-up list.
	IndexOfSelectedItem() int
	// The object corresponding to the last item selected from the pop-up list.
	ObjectValueOfSelectedItem() objectivec.IObject
	// Selects the pop-up list row at the given index.
	SelectItemAtIndex(index int)
	// Selects the first pop-up list item that corresponds to the given object.
	SelectItemWithObjectValue(object objectivec.IObject)

	// Topic: Completing the Text Field

	// A Boolean value indicating whether the combo box tries to complete what the user types.
	Completes() bool
	SetCompletes(value bool)
}

// Init initializes the instance.
func (c NSComboBox) Init() NSComboBox {
	rv := objc.Send[NSComboBox](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSComboBox) Autorelease() NSComboBox {
	rv := objc.Send[NSComboBox](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSComboBox creates a new NSComboBox instance.
func NewNSComboBox() NSComboBox {
	class := getNSComboBoxClass()
	rv := objc.Send[NSComboBox](objc.ID(class.class), objc.Sel("new"))
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
func NewComboBoxLabelWithAttributedString(attributedStringValue foundation.NSAttributedString) NSComboBox {
	rv := objc.Send[objc.ID](objc.ID(getNSComboBoxClass().class), objc.Sel("labelWithAttributedString:"), attributedStringValue)
	return NSComboBoxFromID(rv)
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
func NewComboBoxLabelWithString(stringValue string) NSComboBox {
	rv := objc.Send[objc.ID](objc.ID(getNSComboBoxClass().class), objc.Sel("labelWithString:"), objc.String(stringValue))
	return NSComboBoxFromID(rv)
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
func NewComboBoxTextFieldWithString(stringValue string) NSComboBox {
	rv := objc.Send[objc.ID](objc.ID(getNSComboBoxClass().class), objc.Sel("textFieldWithString:"), objc.String(stringValue))
	return NSComboBoxFromID(rv)
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewComboBoxWithCoder(coder foundation.INSCoder) NSComboBox {
	instance := getNSComboBoxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSComboBoxFromID(rv)
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
func NewComboBoxWithFrame(frameRect corefoundation.CGRect) NSComboBox {
	instance := getNSComboBoxClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSComboBoxFromID(rv)
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
func NewComboBoxWrappingLabelWithString(stringValue string) NSComboBox {
	rv := objc.Send[objc.ID](objc.ID(getNSComboBoxClass().class), objc.Sel("wrappingLabelWithString:"), objc.String(stringValue))
	return NSComboBoxFromID(rv)
}

// Adds multiple objects to the end of the receiver’s internal item list.
//
// objects: An array of the objects to add to the internal item list.
//
// # Discussion
// 
// This method logs a warning if the [UsesDataSource] property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/addItems(withObjectValues:)
func (c NSComboBox) AddItemsWithObjectValues(objects foundation.INSArray) {
	objc.Send[objc.ID](c.ID, objc.Sel("addItemsWithObjectValues:"), objects)
}
// Adds an object to the end of the receiver’s internal item list.
//
// object: The object to add to the internal item list.
//
// # Discussion
// 
// This method logs a warning if the [UsesDataSource] property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/addItem(withObjectValue:)
func (c NSComboBox) AddItemWithObjectValue(object objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("addItemWithObjectValue:"), object)
}
// Inserts an object at the specified location in the receiver’s internal
// item list.
//
// object: The object to add to the internal item list.
//
// index: The index in the list at which to add the new object. The previous item at
// `index`—along with all following items—is shifted down one slot to make
// room
//
// # Discussion
// 
// This method logs a warning if the [UsesDataSource] property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/insertItem(withObjectValue:at:)
func (c NSComboBox) InsertItemWithObjectValueAtIndex(object objectivec.IObject, index int) {
	objc.Send[objc.ID](c.ID, objc.Sel("insertItemWithObjectValue:atIndex:"), object, index)
}
// Removes all items from the receiver’s internal item list.
//
// # Discussion
// 
// This method logs a warning if the [UsesDataSource] property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/removeAllItems()
func (c NSComboBox) RemoveAllItems() {
	objc.Send[objc.ID](c.ID, objc.Sel("removeAllItems"))
}
// Removes the object at the specified location from the receiver’s internal
// item list.
//
// index: The index of the object to remove. All items beyond `index` are moved up
// one slot to fill the gap.
//
// # Discussion
// 
// The removed object receives a `release` message. This method raises an
// [NSRangeException] if `index` is beyond the end of the list and logs a
// warning if the [UsesDataSource] property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/removeItem(at:)
func (c NSComboBox) RemoveItemAtIndex(index int) {
	objc.Send[objc.ID](c.ID, objc.Sel("removeItemAtIndex:"), index)
}
// Removes all occurrences of the given object from the receiver’s internal
// item list.
//
// object: The object to remove from the internal item list. Objects are considered
// equal if they have the same id or if `` returns [true].
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method logs a warning if the [UsesDataSource] property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/removeItem(withObjectValue:)
func (c NSComboBox) RemoveItemWithObjectValue(object objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("removeItemWithObjectValue:"), object)
}
// Searches the receiver’s internal item list for the specified object and
// returns the lowest matching index.
//
// object: The object for which to return the index.
//
// # Return Value
// 
// The lowest index in the internal item list whose corresponding value is
// equal to that of the specified object. Objects are considered equal if they
// have the same id or if `` returns [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If none of the objects in the receiver’s internal item list are equal to
// `anObject`, [IndexOfItemWithObjectValue] returns [NSNotFound].
// 
// # Discussion
// 
// This method logs a warning if the [UsesDataSource] property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/indexOfItem(withObjectValue:)
func (c NSComboBox) IndexOfItemWithObjectValue(object objectivec.IObject) int {
	rv := objc.Send[int](c.ID, objc.Sel("indexOfItemWithObjectValue:"), object)
	return rv
}
// Returns the object located at the given index within the receiver’s
// internal item list.
//
// index: The index of the object to retrieve. If `index` is beyond the end of the
// list, an [NSRangeException] is raised.
//
// # Return Value
// 
// The object located at the specified index in the internal item list.
//
// # Discussion
// 
// This method logs a warning if the [UsesDataSource] property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/itemObjectValue(at:)
func (c NSComboBox) ItemObjectValueAtIndex(index int) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("itemObjectValueAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
// Informs the receiver that the number of items in its data source has
// changed.
//
// # Discussion
// 
// This method allows the receiver to update the scrollers in its displayed
// pop-up list without actually reloading data into the receiver. It is
// particularly useful for a data source that continually receives data in the
// background over a period of time, in which case the [NSComboBox] can remain
// responsive to the user while the data is received.
// 
// See the NSComboBoxDataSource informal protocol specification for
// information on the messages an [NSComboBox] sends to its data source.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/noteNumberOfItemsChanged()
func (c NSComboBox) NoteNumberOfItemsChanged() {
	objc.Send[objc.ID](c.ID, objc.Sel("noteNumberOfItemsChanged"))
}
// Marks the receiver as needing redisplay, so that it will reload the data
// for visible pop-up items and draw the new values.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/reloadData()
func (c NSComboBox) ReloadData() {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadData"))
}
// Scrolls the receiver’s pop-up list vertically so that the item at the
// specified index is as close to the top as possible.
//
// index: The index of the item to scroll to the top.
//
// # Discussion
// 
// The pop-up list need not be displayed at the time this method is invoked.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/scrollItemAtIndexToTop(_:)
func (c NSComboBox) ScrollItemAtIndexToTop(index int) {
	objc.Send[objc.ID](c.ID, objc.Sel("scrollItemAtIndexToTop:"), index)
}
// Scrolls the receiver’s pop-up list vertically so that the item at the
// specified index is visible.
//
// index: The index of the item to make visible.
//
// # Discussion
// 
// The pop-up list need not be displayed at the time this method is invoked.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/scrollItemAtIndexToVisible(_:)
func (c NSComboBox) ScrollItemAtIndexToVisible(index int) {
	objc.Send[objc.ID](c.ID, objc.Sel("scrollItemAtIndexToVisible:"), index)
}
// Deselects the pop-up list item at the specified index if it’s selected.
//
// index: The index of the item to deselect.
//
// # Discussion
// 
// If the selection does in fact change, this method posts an
// [selectionDidChangeNotification] to the default notification center.
//
// [selectionDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSComboBox/selectionDidChangeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/deselectItem(at:)
func (c NSComboBox) DeselectItemAtIndex(index int) {
	objc.Send[objc.ID](c.ID, objc.Sel("deselectItemAtIndex:"), index)
}
// Selects the pop-up list row at the given index.
//
// index: The index of the item to select in the pop-up list.
//
// # Discussion
// 
// Posts an [selectionDidChangeNotification] to the default notification
// center if the selection does in fact change. Note that this method does not
// alter the contents of the combo box’s text field—see [Setting the Combo
// Box’s Value] for more information.
//
// [Setting the Combo Box’s Value]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ComboBox/Tasks/SettingComboBoxValue.html#//apple_ref/doc/uid/20000256
// [selectionDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSComboBox/selectionDidChangeNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/selectItem(at:)
func (c NSComboBox) SelectItemAtIndex(index int) {
	objc.Send[objc.ID](c.ID, objc.Sel("selectItemAtIndex:"), index)
}
// Selects the first pop-up list item that corresponds to the given object.
//
// object: The object to select in the pop-up list. Objects are considered equal if
// they have the same id or if `` returns [true].
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method logs a warning if [UsesDataSource] returns [true]. Posts an
// [selectionDidChangeNotification] to the default notification center if the
// selection does in fact change. Note that this method doesn’t alter the
// contents of the combo box’s text field—see [Setting the Combo Box’s
// Value] for more information.
//
// [Setting the Combo Box’s Value]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ComboBox/Tasks/SettingComboBoxValue.html#//apple_ref/doc/uid/20000256
// [selectionDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSComboBox/selectionDidChangeNotification
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/selectItem(withObjectValue:)
func (c NSComboBox) SelectItemWithObjectValue(object objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("selectItemWithObjectValue:"), object)
}

// A Boolean value indicating whether the combo box has a vertical scroller.
//
// # Discussion
// 
// When the value of this property is [true], the combo box displays a
// vertical scroller even when the pop-up list contains few enough items that
// a scroller is not needed. The default value of this property is [true].
// 
// If the value of this property is [false] and the combo box has more list
// items (either in its internal item list or from its data source) than are
// allowed by [NumberOfVisibleItems], only a subset of items are displayed.
// The [NSComboBox] class’ `scroll...` methods can be used to position this
// subset within the pop-up list.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/hasVerticalScroller
func (c NSComboBox) HasVerticalScroller() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasVerticalScroller"))
	return rv
}
func (c NSComboBox) SetHasVerticalScroller(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setHasVerticalScroller:"), value)
}
// The horizontal and vertical spacing between cells in the pop-up list.
//
// # Discussion
// 
// Spacing values are measured in points. The default spacing is (`3.0`,
// `2.0`).
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/intercellSpacing
func (c NSComboBox) IntercellSpacing() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("intercellSpacing"))
	return corefoundation.CGSize(rv)
}
func (c NSComboBox) SetIntercellSpacing(value corefoundation.CGSize) {
	objc.Send[struct{}](c.ID, objc.Sel("setIntercellSpacing:"), value)
}
// A Boolean value indicating whether the combo box displays a border.
//
// # Discussion
// 
// When the value of this property is [true], the combo box displays a border.
// For example, when displaying a combo box in a table, it is often useful to
// display the combo box without a border. The default value of this property
// is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/isButtonBordered
func (c NSComboBox) ButtonBordered() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isButtonBordered"))
	return rv
}
func (c NSComboBox) SetButtonBordered(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setButtonBordered:"), value)
}
// The height of each item in the pop-up list.
//
// # Discussion
// 
// The height of items is measured in points. The default item height is
// `16.0` points.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/itemHeight
func (c NSComboBox) ItemHeight() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("itemHeight"))
	return rv
}
func (c NSComboBox) SetItemHeight(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setItemHeight:"), value)
}
// The maximum number of visible items to display in the pop-up list at one
// time.
//
// # Discussion
// 
// Use this property to configure how many items can be displayed at the same
// time. If the combo box has a scroller, the user can scroll to view
// additional items beyond the visible range.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/numberOfVisibleItems
func (c NSComboBox) NumberOfVisibleItems() int {
	rv := objc.Send[int](c.ID, objc.Sel("numberOfVisibleItems"))
	return rv
}
func (c NSComboBox) SetNumberOfVisibleItems(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setNumberOfVisibleItems:"), value)
}
// The object that provides the item data for the combo box.
//
// # Discussion
// 
// Assigning an object to this property does not automatically set the
// [UsesDataSource] property to [true]. If the [UsesDataSource] property is
// [false], accessing this property logs a warning. The default value of this
// property is `nil`.
// 
// For information about how to implement a combo box data source, see
// [NSComboBoxDataSource].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/dataSource
func (c NSComboBox) DataSource() NSComboBoxDataSource {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dataSource"))
	return NSComboBoxDataSourceObjectFromID(rv)
}
func (c NSComboBox) SetDataSource(value NSComboBoxDataSource) {
	objc.Send[struct{}](c.ID, objc.Sel("setDataSource:"), value)
}
// A Boolean value indicating whether the combo box retrieves its items from a
// data source object.
//
// # Discussion
// 
// When the value of this property is [true], the combo box retrieves its
// items from the object in the [DataSource] property. When the value is
// [false], the combo box manages an internal list of items, which it gets
// from the ones specified at design time and the ones you add
// programmatically.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/usesDataSource
func (c NSComboBox) UsesDataSource() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("usesDataSource"))
	return rv
}
func (c NSComboBox) SetUsesDataSource(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setUsesDataSource:"), value)
}
// An array of the items from the combo box’s internal list.
//
// # Discussion
// 
// The array contains the objects you added or inserted into the combo box, so
// the type of each object can vary. Accessing this property logs a warning if
// the [UsesDataSource] property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/objectValues
func (c NSComboBox) ObjectValues() foundation.INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("objectValues"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// The total number of items in the pop-up list.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/numberOfItems
func (c NSComboBox) NumberOfItems() int {
	rv := objc.Send[int](c.ID, objc.Sel("numberOfItems"))
	return rv
}
// The index of the last item selected from the pop-up list.
//
// # Discussion
// 
// The value of this property is `-1` if no item is selected; otherwise, it is
// the index of the selected item. Nothing is selected in a newly initialized
// combo box.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/indexOfSelectedItem
func (c NSComboBox) IndexOfSelectedItem() int {
	rv := objc.Send[int](c.ID, objc.Sel("indexOfSelectedItem"))
	return rv
}
// The object corresponding to the last item selected from the pop-up list.
//
// # Discussion
// 
// For combo boxes that use their own internally maintained list of items,
// this property contains the object in that list that is selected. If no item
// is selected, the value in this property is `nil`. Nothing is selected in a
// newly initialized combo box. This method logs a warning if the
// [UsesDataSource] property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/objectValueOfSelectedItem
func (c NSComboBox) ObjectValueOfSelectedItem() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("objectValueOfSelectedItem"))
	return objectivec.Object{ID: rv}
}
// A Boolean value indicating whether the combo box tries to complete what the
// user types.
//
// # Discussion
// 
// When the value of this property is [true], the combo box tries to complete
// what the user is typing. Every time the user types a new character, the
// combo box uses the [CompletedString] method of its cell to get the new
// value. If the string returned by that method is longer than the string
// typed by the user, the combo box replaces the existing string with the
// returned string and selects the additional characters. If the user is
// deleting characters or adds characters somewhere besides the end of the
// string, the combo box does not try to complete it.
// 
// When the value of this property is [false], the combo box does not try to
// complete the string typed by the user.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBox/completes
func (c NSComboBox) Completes() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("completes"))
	return rv
}
func (c NSComboBox) SetCompletes(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setCompletes:"), value)
}

