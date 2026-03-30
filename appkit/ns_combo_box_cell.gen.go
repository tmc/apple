// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSComboBoxCell] class.
var (
	_NSComboBoxCellClass     NSComboBoxCellClass
	_NSComboBoxCellClassOnce sync.Once
)

func getNSComboBoxCellClass() NSComboBoxCellClass {
	_NSComboBoxCellClassOnce.Do(func() {
		_NSComboBoxCellClass = NSComboBoxCellClass{class: objc.GetClass("NSComboBoxCell")}
	})
	return _NSComboBoxCellClass
}

// GetNSComboBoxCellClass returns the class object for NSComboBoxCell.
func GetNSComboBoxCellClass() NSComboBoxCellClass {
	return getNSComboBoxCellClass()
}

type NSComboBoxCellClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSComboBoxCellClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSComboBoxCellClass) Alloc() NSComboBoxCell {
	rv := objc.Send[NSComboBoxCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The user interface of a combo box.
//
// # Overview
//
// [NSComboBoxCell] is a subclass of [NSTextFieldCell] used to implement the
// user interface of “combo boxes” (see [NSComboBox] for information on
// how combo boxes look and work). The [NSComboBox] subclass of [NSTextField]
// uses a single [NSComboBoxCell], and essentially all of the [NSComboBox]
// class’s methods simply invoke the corresponding [NSComboBoxCell] method.
//
// Also see the [NSComboBoxCellDataSource] protocol, which declares the
// methods that an [NSComboBoxCell] object uses to access the contents of its
// data source object.
//
// # Setting Display Attributes
//
//   - [NSComboBoxCell.HasVerticalScroller]: A Boolean value that indicates if the combo box displays a vertical scroller.
//   - [NSComboBoxCell.SetHasVerticalScroller]
//   - [NSComboBoxCell.ButtonBordered]: A Boolean value that indicates whether the combo box button displays a border.
//   - [NSComboBoxCell.SetButtonBordered]
//   - [NSComboBoxCell.IntercellSpacing]: The spacing between cells in the combo box’s pop-up list.
//   - [NSComboBoxCell.SetIntercellSpacing]
//   - [NSComboBoxCell.ItemHeight]: The height of each item in the combo box’s pop-up list.
//   - [NSComboBoxCell.SetItemHeight]
//   - [NSComboBoxCell.NumberOfVisibleItems]: The maximum number of items visible in the pop-up list at any one time.
//   - [NSComboBoxCell.SetNumberOfVisibleItems]
//
// # Accessing a Data Source
//
//   - [NSComboBoxCell.DataSource]: The object that provides the data displayed in the combo box’s pop-up list.
//   - [NSComboBoxCell.SetDataSource]
//   - [NSComboBoxCell.UsesDataSource]: A Boolean value that indicates if the combo box uses an external data source to populate its pop-up list.
//   - [NSComboBoxCell.SetUsesDataSource]
//
// # Working with an Internal List
//
//   - [NSComboBoxCell.AddItemsWithObjectValues]: Adds multiple objects to the internal item list.
//   - [NSComboBoxCell.AddItemWithObjectValue]: Adds the specified object to the internal item list.
//   - [NSComboBoxCell.InsertItemWithObjectValueAtIndex]: Inserts an object at the specified location in the internal item list.
//   - [NSComboBoxCell.ObjectValues]: The combo box’s internal item list in an array.
//   - [NSComboBoxCell.RemoveAllItems]: Removes all items from the combo box’s internal item list.
//   - [NSComboBoxCell.RemoveItemAtIndex]: Removes the object at the specified location from the combo box’s internal item list.
//   - [NSComboBoxCell.RemoveItemWithObjectValue]: Removes all occurrences of the specified object from the combo box’s internal item list.
//   - [NSComboBoxCell.NumberOfItems]: The total number of items in the pop-up list.
//
// # Manipulating the Displayed List
//
//   - [NSComboBoxCell.IndexOfItemWithObjectValue]: Searches the combo box’s internal item list for the given object and returns the matching index number.
//   - [NSComboBoxCell.ItemObjectValueAtIndex]: Returns the object located at the specified location in the internal item list.
//   - [NSComboBoxCell.NoteNumberOfItemsChanged]: Informs the combo box that the number of items in its data source has changed.
//   - [NSComboBoxCell.ReloadData]: Marks the combo box as needing redisplay, so that it will reload the data for visible pop-up items and draw the new values.
//   - [NSComboBoxCell.ScrollItemAtIndexToTop]: Scrolls the combo box’s pop-up list vertically so that the item at the given index is as close to the top as possible.
//   - [NSComboBoxCell.ScrollItemAtIndexToVisible]: Scrolls the combo box’s pop-up list vertically so that the item at the given index is visible.
//
// # Manipulating the Selection
//
//   - [NSComboBoxCell.DeselectItemAtIndex]: Deselects the pop-up list item at the given index if it’s selected.
//   - [NSComboBoxCell.IndexOfSelectedItem]: The index of the last item selected from the pop-up list.
//   - [NSComboBoxCell.ObjectValueOfSelectedItem]: The object corresponding to the last item selected from the pop-up list.
//   - [NSComboBoxCell.SelectItemAtIndex]: Selects the pop-up list row at the given index.
//   - [NSComboBoxCell.SelectItemWithObjectValue]: Selects the first pop-up list item that corresponds to the specified object.
//
// # Completing the Text Field
//
//   - [NSComboBoxCell.CompletedString]: Returns a string from the combo box’s pop-up list that starts with the given substring.
//   - [NSComboBoxCell.Completes]: A Boolean value that indicates if the combo box tries to complete text entered by the user.
//   - [NSComboBoxCell.SetCompletes]
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell
type NSComboBoxCell struct {
	NSTextFieldCell
}

// NSComboBoxCellFromID constructs a [NSComboBoxCell] from an objc.ID.
//
// The user interface of a combo box.
func NSComboBoxCellFromID(id objc.ID) NSComboBoxCell {
	return NSComboBoxCell{NSTextFieldCell: NSTextFieldCellFromID(id)}
}

// NOTE: NSComboBoxCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSComboBoxCell] class.
//
// # Setting Display Attributes
//
//   - [INSComboBoxCell.HasVerticalScroller]: A Boolean value that indicates if the combo box displays a vertical scroller.
//   - [INSComboBoxCell.SetHasVerticalScroller]
//   - [INSComboBoxCell.ButtonBordered]: A Boolean value that indicates whether the combo box button displays a border.
//   - [INSComboBoxCell.SetButtonBordered]
//   - [INSComboBoxCell.IntercellSpacing]: The spacing between cells in the combo box’s pop-up list.
//   - [INSComboBoxCell.SetIntercellSpacing]
//   - [INSComboBoxCell.ItemHeight]: The height of each item in the combo box’s pop-up list.
//   - [INSComboBoxCell.SetItemHeight]
//   - [INSComboBoxCell.NumberOfVisibleItems]: The maximum number of items visible in the pop-up list at any one time.
//   - [INSComboBoxCell.SetNumberOfVisibleItems]
//
// # Accessing a Data Source
//
//   - [INSComboBoxCell.DataSource]: The object that provides the data displayed in the combo box’s pop-up list.
//   - [INSComboBoxCell.SetDataSource]
//   - [INSComboBoxCell.UsesDataSource]: A Boolean value that indicates if the combo box uses an external data source to populate its pop-up list.
//   - [INSComboBoxCell.SetUsesDataSource]
//
// # Working with an Internal List
//
//   - [INSComboBoxCell.AddItemsWithObjectValues]: Adds multiple objects to the internal item list.
//   - [INSComboBoxCell.AddItemWithObjectValue]: Adds the specified object to the internal item list.
//   - [INSComboBoxCell.InsertItemWithObjectValueAtIndex]: Inserts an object at the specified location in the internal item list.
//   - [INSComboBoxCell.ObjectValues]: The combo box’s internal item list in an array.
//   - [INSComboBoxCell.RemoveAllItems]: Removes all items from the combo box’s internal item list.
//   - [INSComboBoxCell.RemoveItemAtIndex]: Removes the object at the specified location from the combo box’s internal item list.
//   - [INSComboBoxCell.RemoveItemWithObjectValue]: Removes all occurrences of the specified object from the combo box’s internal item list.
//   - [INSComboBoxCell.NumberOfItems]: The total number of items in the pop-up list.
//
// # Manipulating the Displayed List
//
//   - [INSComboBoxCell.IndexOfItemWithObjectValue]: Searches the combo box’s internal item list for the given object and returns the matching index number.
//   - [INSComboBoxCell.ItemObjectValueAtIndex]: Returns the object located at the specified location in the internal item list.
//   - [INSComboBoxCell.NoteNumberOfItemsChanged]: Informs the combo box that the number of items in its data source has changed.
//   - [INSComboBoxCell.ReloadData]: Marks the combo box as needing redisplay, so that it will reload the data for visible pop-up items and draw the new values.
//   - [INSComboBoxCell.ScrollItemAtIndexToTop]: Scrolls the combo box’s pop-up list vertically so that the item at the given index is as close to the top as possible.
//   - [INSComboBoxCell.ScrollItemAtIndexToVisible]: Scrolls the combo box’s pop-up list vertically so that the item at the given index is visible.
//
// # Manipulating the Selection
//
//   - [INSComboBoxCell.DeselectItemAtIndex]: Deselects the pop-up list item at the given index if it’s selected.
//   - [INSComboBoxCell.IndexOfSelectedItem]: The index of the last item selected from the pop-up list.
//   - [INSComboBoxCell.ObjectValueOfSelectedItem]: The object corresponding to the last item selected from the pop-up list.
//   - [INSComboBoxCell.SelectItemAtIndex]: Selects the pop-up list row at the given index.
//   - [INSComboBoxCell.SelectItemWithObjectValue]: Selects the first pop-up list item that corresponds to the specified object.
//
// # Completing the Text Field
//
//   - [INSComboBoxCell.CompletedString]: Returns a string from the combo box’s pop-up list that starts with the given substring.
//   - [INSComboBoxCell.Completes]: A Boolean value that indicates if the combo box tries to complete text entered by the user.
//   - [INSComboBoxCell.SetCompletes]
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell
type INSComboBoxCell interface {
	INSTextFieldCell

	// Topic: Setting Display Attributes

	// A Boolean value that indicates if the combo box displays a vertical scroller.
	HasVerticalScroller() bool
	SetHasVerticalScroller(value bool)
	// A Boolean value that indicates whether the combo box button displays a border.
	ButtonBordered() bool
	SetButtonBordered(value bool)
	// The spacing between cells in the combo box’s pop-up list.
	IntercellSpacing() corefoundation.CGSize
	SetIntercellSpacing(value corefoundation.CGSize)
	// The height of each item in the combo box’s pop-up list.
	ItemHeight() float64
	SetItemHeight(value float64)
	// The maximum number of items visible in the pop-up list at any one time.
	NumberOfVisibleItems() int
	SetNumberOfVisibleItems(value int)

	// Topic: Accessing a Data Source

	// The object that provides the data displayed in the combo box’s pop-up list.
	DataSource() NSComboBoxCellDataSource
	SetDataSource(value NSComboBoxCellDataSource)
	// A Boolean value that indicates if the combo box uses an external data source to populate its pop-up list.
	UsesDataSource() bool
	SetUsesDataSource(value bool)

	// Topic: Working with an Internal List

	// Adds multiple objects to the internal item list.
	AddItemsWithObjectValues(objects foundation.INSArray)
	// Adds the specified object to the internal item list.
	AddItemWithObjectValue(object objectivec.IObject)
	// Inserts an object at the specified location in the internal item list.
	InsertItemWithObjectValueAtIndex(object objectivec.IObject, index int)
	// The combo box’s internal item list in an array.
	ObjectValues() foundation.INSArray
	// Removes all items from the combo box’s internal item list.
	RemoveAllItems()
	// Removes the object at the specified location from the combo box’s internal item list.
	RemoveItemAtIndex(index int)
	// Removes all occurrences of the specified object from the combo box’s internal item list.
	RemoveItemWithObjectValue(object objectivec.IObject)
	// The total number of items in the pop-up list.
	NumberOfItems() int

	// Topic: Manipulating the Displayed List

	// Searches the combo box’s internal item list for the given object and returns the matching index number.
	IndexOfItemWithObjectValue(object objectivec.IObject) int
	// Returns the object located at the specified location in the internal item list.
	ItemObjectValueAtIndex(index int) objectivec.IObject
	// Informs the combo box that the number of items in its data source has changed.
	NoteNumberOfItemsChanged()
	// Marks the combo box as needing redisplay, so that it will reload the data for visible pop-up items and draw the new values.
	ReloadData()
	// Scrolls the combo box’s pop-up list vertically so that the item at the given index is as close to the top as possible.
	ScrollItemAtIndexToTop(index int)
	// Scrolls the combo box’s pop-up list vertically so that the item at the given index is visible.
	ScrollItemAtIndexToVisible(index int)

	// Topic: Manipulating the Selection

	// Deselects the pop-up list item at the given index if it’s selected.
	DeselectItemAtIndex(index int)
	// The index of the last item selected from the pop-up list.
	IndexOfSelectedItem() int
	// The object corresponding to the last item selected from the pop-up list.
	ObjectValueOfSelectedItem() objectivec.IObject
	// Selects the pop-up list row at the given index.
	SelectItemAtIndex(index int)
	// Selects the first pop-up list item that corresponds to the specified object.
	SelectItemWithObjectValue(object objectivec.IObject)

	// Topic: Completing the Text Field

	// Returns a string from the combo box’s pop-up list that starts with the given substring.
	CompletedString(string_ string) string
	// A Boolean value that indicates if the combo box tries to complete text entered by the user.
	Completes() bool
	SetCompletes(value bool)
}

// Init initializes the instance.
func (c NSComboBoxCell) Init() NSComboBoxCell {
	rv := objc.Send[NSComboBoxCell](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSComboBoxCell) Autorelease() NSComboBoxCell {
	rv := objc.Send[NSComboBoxCell](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSComboBoxCell creates a new NSComboBoxCell instance.
func NewNSComboBoxCell() NSComboBoxCell {
	class := getNSComboBoxCellClass()
	rv := objc.Send[NSComboBoxCell](objc.ID(class.class), objc.Sel("new"))
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
func NewComboBoxCellImageCell(image INSImage) NSComboBoxCell {
	instance := getNSComboBoxCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSComboBoxCellFromID(rv)
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
func NewComboBoxCellTextCell(string_ string) NSComboBoxCell {
	instance := getNSComboBoxCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSComboBoxCellFromID(rv)
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
func NewComboBoxCellWithCoder(coder foundation.INSCoder) NSComboBoxCell {
	instance := getNSComboBoxCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSComboBoxCellFromID(rv)
}

// Adds multiple objects to the internal item list.
//
// objects: The object to add to the end of the combo box’s internal item list.
//
// # Discussion
//
// This method logs a warning if [UsesDataSource] is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/addItems(withObjectValues:)
func (c NSComboBoxCell) AddItemsWithObjectValues(objects foundation.INSArray) {
	objc.Send[objc.ID](c.ID, objc.Sel("addItemsWithObjectValues:"), objects)
}

// Adds the specified object to the internal item list.
//
// object: The object to add to the end of the combo box’s internal item list.
//
// # Discussion
//
// This method logs a warning if [UsesDataSource] is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/addItem(withObjectValue:)
func (c NSComboBoxCell) AddItemWithObjectValue(object objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("addItemWithObjectValue:"), object)
}

// Inserts an object at the specified location in the internal item list.
//
// object: The object to add to the combo box’s internal item list.
//
// index: The index at which to add the specified object. The previous item at
// `index`—along with all following items—is shifted down one slot to make
// room.
//
// # Discussion
//
// This method logs a warning if [UsesDataSource] is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/insertItem(withObjectValue:at:)
func (c NSComboBoxCell) InsertItemWithObjectValueAtIndex(object objectivec.IObject, index int) {
	objc.Send[objc.ID](c.ID, objc.Sel("insertItemWithObjectValue:atIndex:"), object, index)
}

// Removes all items from the combo box’s internal item list.
//
// # Discussion
//
// This method logs a warning if [UsesDataSource] is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/removeAllItems()
func (c NSComboBoxCell) RemoveAllItems() {
	objc.Send[objc.ID](c.ID, objc.Sel("removeAllItems"))
}

// Removes the object at the specified location from the combo box’s
// internal item list.
//
// index: The index of the object to remove from the combo box’s internal item
// list. All items beyond `index` are moved up one slot to fill the gap.
//
// # Discussion
//
// The removed object receives a `release` message. This method raises an
// [NSRangeException] if `index` is beyond the end of the list and logs a
// warning if [UsesDataSource] is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/removeItem(at:)
func (c NSComboBoxCell) RemoveItemAtIndex(index int) {
	objc.Send[objc.ID](c.ID, objc.Sel("removeItemAtIndex:"), index)
}

// Removes all occurrences of the specified object from the combo box’s
// internal item list.
//
// object: The object to remove from the combo box’s internal item list. Objects are
// considered equal if they have the same id or if “ returns true.
//
// # Discussion
//
// This method logs a warning if [UsesDataSource] is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/removeItem(withObjectValue:)
func (c NSComboBoxCell) RemoveItemWithObjectValue(object objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("removeItemWithObjectValue:"), object)
}

// Searches the combo box’s internal item list for the given object and
// returns the matching index number.
//
// object: The object for which to return the index.
//
// # Return Value
//
// The lowest index whose corresponding value is equal to `anObject`. Objects
// are considered equal if they have the same id or if “ returns true. If
// none of the objects in the combo box’s internal item list is equal to
// `anObject`, [IndexOfItemWithObjectValue] returns [NSNotFound].
//
// # Discussion
//
// This method logs a warning if [UsesDataSource] is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/indexOfItem(withObjectValue:)
func (c NSComboBoxCell) IndexOfItemWithObjectValue(object objectivec.IObject) int {
	rv := objc.Send[int](c.ID, objc.Sel("indexOfItemWithObjectValue:"), object)
	return rv
}

// Returns the object located at the specified location in the internal item
// list.
//
// index: The index of the object to return. If `index` is beyond the end of the
// list, an [NSRangeException] is raised.
//
// # Return Value
//
// The object at the given location in the combo box’s internal item list.
//
// # Discussion
//
// This method logs a warning if [UsesDataSource] is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/itemObjectValue(at:)
func (c NSComboBoxCell) ItemObjectValueAtIndex(index int) objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("itemObjectValueAtIndex:"), index)
	return objectivec.Object{ID: rv}
}

// Informs the combo box that the number of items in its data source has
// changed.
//
// # Discussion
//
// This method allows the combo box to update the scrollers in its displayed
// pop-up list without actually reloading data into the combo box. It is
// particularly useful for a data source that continually receives data in the
// background over a period of time, in which case the [NSComboBoxCell] can
// remain responsive to the user while the data is received.
//
// See the NSComboBoxCellDataSource informal protocol specification for
// information on the messages an [NSComboBoxCell] sends to its data source.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/noteNumberOfItemsChanged()
func (c NSComboBoxCell) NoteNumberOfItemsChanged() {
	objc.Send[objc.ID](c.ID, objc.Sel("noteNumberOfItemsChanged"))
}

// Marks the combo box as needing redisplay, so that it will reload the data
// for visible pop-up items and draw the new values.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/reloadData()
func (c NSComboBoxCell) ReloadData() {
	objc.Send[objc.ID](c.ID, objc.Sel("reloadData"))
}

// Scrolls the combo box’s pop-up list vertically so that the item at the
// given index is as close to the top as possible.
//
// index: The index of the item to scroll to the top.
//
// # Discussion
//
// The pop-up list need not be displayed at the time this method is invoked.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/scrollItemAtIndexToTop(_:)
func (c NSComboBoxCell) ScrollItemAtIndexToTop(index int) {
	objc.Send[objc.ID](c.ID, objc.Sel("scrollItemAtIndexToTop:"), index)
}

// Scrolls the combo box’s pop-up list vertically so that the item at the
// given index is visible.
//
// index: The index of the item to make visible.
//
// # Discussion
//
// The pop-up list need not be displayed at the time this method is invoked.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/scrollItemAtIndexToVisible(_:)
func (c NSComboBoxCell) ScrollItemAtIndexToVisible(index int) {
	objc.Send[objc.ID](c.ID, objc.Sel("scrollItemAtIndexToVisible:"), index)
}

// Deselects the pop-up list item at the given index if it’s selected.
//
// index: The index of the item to deselect.
//
// # Discussion
//
// If the selection does in fact change, this method posts an
// [selectionDidChangeNotification] to the default notification center.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/deselectItem(at:)
//
// [selectionDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSComboBox/selectionDidChangeNotification
func (c NSComboBoxCell) DeselectItemAtIndex(index int) {
	objc.Send[objc.ID](c.ID, objc.Sel("deselectItemAtIndex:"), index)
}

// Selects the pop-up list row at the given index.
//
// index: The index of the row to select.
//
// # Discussion
//
// Posts an [selectionDidChangeNotification] to the default notification
// center if the selection does in fact change. Note that this method does not
// alter the contents of the combo box cell’s text field—see [Combo Box
// Programming Topics] for more information.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/selectItem(at:)
//
// [Combo Box Programming Topics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ComboBox/ComboBox.html#//apple_ref/doc/uid/10000020i
// [selectionDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSComboBox/selectionDidChangeNotification
func (c NSComboBoxCell) SelectItemAtIndex(index int) {
	objc.Send[objc.ID](c.ID, objc.Sel("selectItemAtIndex:"), index)
}

// Selects the first pop-up list item that corresponds to the specified
// object.
//
// object: The object for which to select the corresponding pop-up list item. Objects
// are considered equal if they have the same id or if “ returns true.
//
// # Discussion
//
// This method logs a warning if [UsesDataSource] is true. Posts an
// [selectionDidChangeNotification] to the default notification center if the
// selection does in fact change. Note that this method doesn’t alter the
// contents of the combo box cell’s text field—see [Setting the Combo
// Box’s Value] for more information.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/selectItem(withObjectValue:)
//
// [Setting the Combo Box’s Value]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ComboBox/Tasks/SettingComboBoxValue.html#//apple_ref/doc/uid/20000256
// [selectionDidChangeNotification]: https://developer.apple.com/documentation/AppKit/NSComboBox/selectionDidChangeNotification
func (c NSComboBoxCell) SelectItemWithObjectValue(object objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("selectItemWithObjectValue:"), object)
}

// Returns a string from the combo box’s pop-up list that starts with the
// given substring.
//
// string: The substring to search for. This is what the user entered in the combo
// box’s text field.
//
// # Return Value
//
// The string from the combo box’s pop-up list that starts with the
// specified substring or `nil` if there is no such string.
//
// # Discussion
//
// The default implementation of this method first checks whether the combo
// box uses a data source and whether the data source responds to
// comboBox:completedString: or comboBoxCell:completedString:. If so, the
// combo box cell returns that method’s return value. Otherwise, this method
// goes through the combo box’s items one by one and returns an item that
// starts with `substring`.
//
// Override this method only if your subclass completes strings differently.
// The overriding method does not need to call the superclass’s method.
// Generally, you do not need to call this method directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/completedString(_:)
func (c NSComboBoxCell) CompletedString(string_ string) string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("completedString:"), objc.String(string_))
	return foundation.NSStringFromID(rv).String()
}

// A Boolean value that indicates if the combo box displays a vertical
// scroller.
//
// # Discussion
//
// When the value of this property is true, the combo box displays a vertical
// scroller; when the value is false, it does not. The default value of this
// property is true. Note that the scroller is displayed even if the pop-up
// list contains fewer items than will fit in the area specified for display.
//
// If you set this property to false and the combo box cell has more list
// items (either in its internal item list or from its data source) than are
// allowed by [NumberOfVisibleItems], only a subset are displayed. The
// [NSComboBoxCell] `scroll...` methods can be used to position this subset
// within the pop-up list.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/hasVerticalScroller
func (c NSComboBoxCell) HasVerticalScroller() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasVerticalScroller"))
	return rv
}
func (c NSComboBoxCell) SetHasVerticalScroller(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setHasVerticalScroller:"), value)
}

// A Boolean value that indicates whether the combo box button displays a
// border.
//
// # Discussion
//
// When the value of this property is true, the button has a border; when it
// is false, the button is borderless. For example, it is often useful when
// using a combo box in an [NSTableView] to display the button without a
// border.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/isButtonBordered
func (c NSComboBoxCell) ButtonBordered() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isButtonBordered"))
	return rv
}
func (c NSComboBoxCell) SetButtonBordered(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setButtonBordered:"), value)
}

// The spacing between cells in the combo box’s pop-up list.
//
// # Discussion
//
// The value of this property is the horizontal and vertical spacing between
// cells in the combo box’s pop-up list. The default spacing is (3.0, 2.0).
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/intercellSpacing
func (c NSComboBoxCell) IntercellSpacing() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("intercellSpacing"))
	return corefoundation.CGSize(rv)
}
func (c NSComboBoxCell) SetIntercellSpacing(value corefoundation.CGSize) {
	objc.Send[struct{}](c.ID, objc.Sel("setIntercellSpacing:"), value)
}

// The height of each item in the combo box’s pop-up list.
//
// # Discussion
//
// The default item height is 16.0.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/itemHeight
func (c NSComboBoxCell) ItemHeight() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("itemHeight"))
	return rv
}
func (c NSComboBoxCell) SetItemHeight(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setItemHeight:"), value)
}

// The maximum number of items visible in the pop-up list at any one time.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/numberOfVisibleItems
func (c NSComboBoxCell) NumberOfVisibleItems() int {
	rv := objc.Send[int](c.ID, objc.Sel("numberOfVisibleItems"))
	return rv
}
func (c NSComboBoxCell) SetNumberOfVisibleItems(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setNumberOfVisibleItems:"), value)
}

// The object that provides the data displayed in the combo box’s pop-up
// list.
//
// # Discussion
//
// The value of this property should be an object that implements the
// appropriate methods of the NSComboBoxCellDataSource informal protocol. Note
// that setting this property doesn’t automatically set [UsesDataSource] to
// false and in fact logs a warning if [UsesDataSource] is false. If you set
// this property to an object that doesn’t respond to either
// numberOfItemsInComboBoxCell: or comboBoxCell:objectValueForItemAtIndex:, a
// warning is logged if [UsesDataSource] is false. See the class description
// and the NSComboBoxCellDataSource informal protocol specification for more
// information on combo box cell data source objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/dataSource
func (c NSComboBoxCell) DataSource() NSComboBoxCellDataSource {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("dataSource"))
	return NSComboBoxCellDataSourceObjectFromID(rv)
}
func (c NSComboBoxCell) SetDataSource(value NSComboBoxCellDataSource) {
	objc.Send[struct{}](c.ID, objc.Sel("setDataSource:"), value)
}

// A Boolean value that indicates if the combo box uses an external data
// source to populate its pop-up list.
//
// # Discussion
//
// When the value of this property is true, the combo box uses an external
// data source to populate its pop-up list; when it is false, the combo box
// uses an internal item list.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/usesDataSource
func (c NSComboBoxCell) UsesDataSource() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("usesDataSource"))
	return rv
}
func (c NSComboBoxCell) SetUsesDataSource(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setUsesDataSource:"), value)
}

// The combo box’s internal item list in an array.
//
// # Discussion
//
// Accessing this property logs a warning if [UsesDataSource] is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/objectValues
func (c NSComboBoxCell) ObjectValues() foundation.INSArray {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("objectValues"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// The total number of items in the pop-up list.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/numberOfItems
func (c NSComboBoxCell) NumberOfItems() int {
	rv := objc.Send[int](c.ID, objc.Sel("numberOfItems"))
	return rv
}

// The index of the last item selected from the pop-up list.
//
// # Discussion
//
// The index of the last item selected from the combo box’s pop-up list or
// –1 if no item is selected. Note that nothing is initially selected in a
// newly initialized combo box cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/indexOfSelectedItem
func (c NSComboBoxCell) IndexOfSelectedItem() int {
	rv := objc.Send[int](c.ID, objc.Sel("indexOfSelectedItem"))
	return rv
}

// The object corresponding to the last item selected from the pop-up list.
//
// # Discussion
//
// The value of this property is the object from the combo box’s internal
// item list corresponding to the last item selected from the pop-up list, or
// `nil` if no item is selected.
//
// Note that nothing is initially selected in a newly initialized combo box
// cell. Accessing this property logs a warning if [UsesDataSource] is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/objectValueOfSelectedItem
func (c NSComboBoxCell) ObjectValueOfSelectedItem() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("objectValueOfSelectedItem"))
	return objectivec.Object{ID: rv}
}

// A Boolean value that indicates if the combo box tries to complete text
// entered by the user.
//
// # Discussion
//
// When the value of this property is true, the combo box tries to complete
// what the user types in the text field and every time the user adds
// characters to the end of the text field, the combo box calls
// [CompletedString]; when it is false, it does not.
//
// If [CompletedString] returns a string that’s longer than the existing
// string, the combo box replaces the existing string with the returned string
// and selects the additional characters. If the user is deleting characters
// or adds characters somewhere besides the end of the string, the combo box
// does not try to complete it.
//
// See: https://developer.apple.com/documentation/AppKit/NSComboBoxCell/completes
func (c NSComboBoxCell) Completes() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("completes"))
	return rv
}
func (c NSComboBoxCell) SetCompletes(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setCompletes:"), value)
}
