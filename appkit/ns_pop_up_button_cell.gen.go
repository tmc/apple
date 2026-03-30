// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPopUpButtonCell] class.
var (
	_NSPopUpButtonCellClass     NSPopUpButtonCellClass
	_NSPopUpButtonCellClassOnce sync.Once
)

func getNSPopUpButtonCellClass() NSPopUpButtonCellClass {
	_NSPopUpButtonCellClassOnce.Do(func() {
		_NSPopUpButtonCellClass = NSPopUpButtonCellClass{class: objc.GetClass("NSPopUpButtonCell")}
	})
	return _NSPopUpButtonCellClass
}

// GetNSPopUpButtonCellClass returns the class object for NSPopUpButtonCell.
func GetNSPopUpButtonCellClass() NSPopUpButtonCellClass {
	return getNSPopUpButtonCellClass()
}

type NSPopUpButtonCellClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPopUpButtonCellClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPopUpButtonCellClass) Alloc() NSPopUpButtonCell {
	rv := objc.Send[NSPopUpButtonCell](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The [NSPopUpButtonCell] class defines the visual appearance of pop-up
// buttons that display pop-up or pull-down menus. Pop-up menus present the
// user with a set of choices, much the way radio buttons do, but using much
// less space. Pull-down menus also provide a set of choices but present the
// information in a slightly different way, usually to provide a set of
// commands from which the user can choose.
//
// # Overview
//
// The [NSPopUpButtonCell] class implements the user interface for the
// [NSPopUpButton] class.
//
// Changes made to a menu (such as adding, removing, or changing the items)
// are not apparent while the menu is being displayed or interacted with.
//
// # Initialization
//
//   - [NSPopUpButtonCell.InitTextCellPullsDown]: Returns an [NSPopUpButtonCell] object initialized with the specified title.
//
// # Accessing menu attributes
//
//   - [NSPopUpButtonCell.PullsDown]: A Boolean value that indicates the behavior of the button’s menu.
//   - [NSPopUpButtonCell.SetPullsDown]
//   - [NSPopUpButtonCell.AutoenablesItems]: A Boolean value that indicates if the button automatically enables and disables its items every time a user event occurs.
//   - [NSPopUpButtonCell.SetAutoenablesItems]
//   - [NSPopUpButtonCell.PreferredEdge]: The edge of the cell from which the menu should pop out when screen conditions are restrictive.
//   - [NSPopUpButtonCell.SetPreferredEdge]
//   - [NSPopUpButtonCell.UsesItemFromMenu]: A Boolean value that indicates if the control uses an item from the menu for its own title.
//   - [NSPopUpButtonCell.SetUsesItemFromMenu]
//   - [NSPopUpButtonCell.AltersStateOfSelectedItem]: A Boolean value that indicates if the pop-up button links the state of the selected menu item to the current selection.
//   - [NSPopUpButtonCell.SetAltersStateOfSelectedItem]
//   - [NSPopUpButtonCell.ArrowPosition]: The position of the arrow displayed on the button.
//   - [NSPopUpButtonCell.SetArrowPosition]
//
// # Adding and removing items
//
//   - [NSPopUpButtonCell.AddItemWithTitle]: Adds an item with the specified title to the end of the menu.
//   - [NSPopUpButtonCell.AddItemsWithTitles]: Adds multiple items to the end of the menu.
//   - [NSPopUpButtonCell.InsertItemWithTitleAtIndex]: Inserts an item at the specified position in the menu.
//   - [NSPopUpButtonCell.RemoveItemWithTitle]: Removes the item with the specified title from the menu.
//   - [NSPopUpButtonCell.RemoveItemAtIndex]: Removes the item at the specified index.
//   - [NSPopUpButtonCell.RemoveAllItems]: Removes all items in the receiver’s item menu.
//
// # Accessing the items
//
//   - [NSPopUpButtonCell.ItemArray]: An array of [NSMenuItem](<doc://com.apple.appkit/documentation/AppKit/NSMenuItem>) objects that represent the items in the menu.
//   - [NSPopUpButtonCell.NumberOfItems]: The number of items in the menu.
//   - [NSPopUpButtonCell.IndexOfItem]: Returns the index of the specified menu item.
//   - [NSPopUpButtonCell.IndexOfItemWithTitle]: Returns the index of the item with the specified title.
//   - [NSPopUpButtonCell.IndexOfItemWithTag]: Returns the index of the menu item with the specified tag.
//   - [NSPopUpButtonCell.IndexOfItemWithRepresentedObject]: Returns the index of the menu item that holds the specified represented object.
//   - [NSPopUpButtonCell.IndexOfItemWithTargetAndAction]: Returns the index of the menu item with the specified target and action.
//   - [NSPopUpButtonCell.ItemAtIndex]: Returns the menu item at the specified index.
//   - [NSPopUpButtonCell.ItemWithTitle]: Returns the menu item with the specified title.
//   - [NSPopUpButtonCell.LastItem]: The last item in the menu.
//
// # Dealing with selection
//
//   - [NSPopUpButtonCell.SelectItem]: Selects the specified menu item.
//   - [NSPopUpButtonCell.SelectItemAtIndex]: Selects the item in the menu at the specified index.
//   - [NSPopUpButtonCell.SelectItemWithTag]: Selects the menu item with the specified tag.
//   - [NSPopUpButtonCell.SelectItemWithTitle]: Selects the item with the specified title.
//   - [NSPopUpButtonCell.SelectedItem]: The menu item last selected by the user.
//   - [NSPopUpButtonCell.IndexOfSelectedItem]: The index of the item last selected by the user.
//   - [NSPopUpButtonCell.SynchronizeTitleAndSelectedItem]: Synchronizes the pop-up button’s displayed item with the currently selected menu item.
//
// # Title conveniences
//
//   - [NSPopUpButtonCell.ItemTitleAtIndex]: Returns the title of the item at the specified index.
//   - [NSPopUpButtonCell.ItemTitles]: An array of [NSString] objects containing the titles of every item in the menu.
//   - [NSPopUpButtonCell.TitleOfSelectedItem]: The title of the item last selected by the user.
//
// # Handling events and action messages
//
//   - [NSPopUpButtonCell.AttachPopUpWithFrameInView]: Sets up the receiver to display a menu.
//   - [NSPopUpButtonCell.DismissPopUp]: Dismisses the pop-up button’s menu by ordering its window out.
//   - [NSPopUpButtonCell.PerformClickWithFrameInView]: Displays the receiver’s menu and track mouse events in it.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell
type NSPopUpButtonCell struct {
	NSMenuItemCell
}

// NSPopUpButtonCellFromID constructs a [NSPopUpButtonCell] from an objc.ID.
//
// The [NSPopUpButtonCell] class defines the visual appearance of pop-up
// buttons that display pop-up or pull-down menus. Pop-up menus present the
// user with a set of choices, much the way radio buttons do, but using much
// less space. Pull-down menus also provide a set of choices but present the
// information in a slightly different way, usually to provide a set of
// commands from which the user can choose.
func NSPopUpButtonCellFromID(id objc.ID) NSPopUpButtonCell {
	return NSPopUpButtonCell{NSMenuItemCell: NSMenuItemCellFromID(id)}
}

// NOTE: NSPopUpButtonCell adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPopUpButtonCell] class.
//
// # Initialization
//
//   - [INSPopUpButtonCell.InitTextCellPullsDown]: Returns an [NSPopUpButtonCell] object initialized with the specified title.
//
// # Accessing menu attributes
//
//   - [INSPopUpButtonCell.PullsDown]: A Boolean value that indicates the behavior of the button’s menu.
//   - [INSPopUpButtonCell.SetPullsDown]
//   - [INSPopUpButtonCell.AutoenablesItems]: A Boolean value that indicates if the button automatically enables and disables its items every time a user event occurs.
//   - [INSPopUpButtonCell.SetAutoenablesItems]
//   - [INSPopUpButtonCell.PreferredEdge]: The edge of the cell from which the menu should pop out when screen conditions are restrictive.
//   - [INSPopUpButtonCell.SetPreferredEdge]
//   - [INSPopUpButtonCell.UsesItemFromMenu]: A Boolean value that indicates if the control uses an item from the menu for its own title.
//   - [INSPopUpButtonCell.SetUsesItemFromMenu]
//   - [INSPopUpButtonCell.AltersStateOfSelectedItem]: A Boolean value that indicates if the pop-up button links the state of the selected menu item to the current selection.
//   - [INSPopUpButtonCell.SetAltersStateOfSelectedItem]
//   - [INSPopUpButtonCell.ArrowPosition]: The position of the arrow displayed on the button.
//   - [INSPopUpButtonCell.SetArrowPosition]
//
// # Adding and removing items
//
//   - [INSPopUpButtonCell.AddItemWithTitle]: Adds an item with the specified title to the end of the menu.
//   - [INSPopUpButtonCell.AddItemsWithTitles]: Adds multiple items to the end of the menu.
//   - [INSPopUpButtonCell.InsertItemWithTitleAtIndex]: Inserts an item at the specified position in the menu.
//   - [INSPopUpButtonCell.RemoveItemWithTitle]: Removes the item with the specified title from the menu.
//   - [INSPopUpButtonCell.RemoveItemAtIndex]: Removes the item at the specified index.
//   - [INSPopUpButtonCell.RemoveAllItems]: Removes all items in the receiver’s item menu.
//
// # Accessing the items
//
//   - [INSPopUpButtonCell.ItemArray]: An array of [NSMenuItem](<doc://com.apple.appkit/documentation/AppKit/NSMenuItem>) objects that represent the items in the menu.
//   - [INSPopUpButtonCell.NumberOfItems]: The number of items in the menu.
//   - [INSPopUpButtonCell.IndexOfItem]: Returns the index of the specified menu item.
//   - [INSPopUpButtonCell.IndexOfItemWithTitle]: Returns the index of the item with the specified title.
//   - [INSPopUpButtonCell.IndexOfItemWithTag]: Returns the index of the menu item with the specified tag.
//   - [INSPopUpButtonCell.IndexOfItemWithRepresentedObject]: Returns the index of the menu item that holds the specified represented object.
//   - [INSPopUpButtonCell.IndexOfItemWithTargetAndAction]: Returns the index of the menu item with the specified target and action.
//   - [INSPopUpButtonCell.ItemAtIndex]: Returns the menu item at the specified index.
//   - [INSPopUpButtonCell.ItemWithTitle]: Returns the menu item with the specified title.
//   - [INSPopUpButtonCell.LastItem]: The last item in the menu.
//
// # Dealing with selection
//
//   - [INSPopUpButtonCell.SelectItem]: Selects the specified menu item.
//   - [INSPopUpButtonCell.SelectItemAtIndex]: Selects the item in the menu at the specified index.
//   - [INSPopUpButtonCell.SelectItemWithTag]: Selects the menu item with the specified tag.
//   - [INSPopUpButtonCell.SelectItemWithTitle]: Selects the item with the specified title.
//   - [INSPopUpButtonCell.SelectedItem]: The menu item last selected by the user.
//   - [INSPopUpButtonCell.IndexOfSelectedItem]: The index of the item last selected by the user.
//   - [INSPopUpButtonCell.SynchronizeTitleAndSelectedItem]: Synchronizes the pop-up button’s displayed item with the currently selected menu item.
//
// # Title conveniences
//
//   - [INSPopUpButtonCell.ItemTitleAtIndex]: Returns the title of the item at the specified index.
//   - [INSPopUpButtonCell.ItemTitles]: An array of [NSString] objects containing the titles of every item in the menu.
//   - [INSPopUpButtonCell.TitleOfSelectedItem]: The title of the item last selected by the user.
//
// # Handling events and action messages
//
//   - [INSPopUpButtonCell.AttachPopUpWithFrameInView]: Sets up the receiver to display a menu.
//   - [INSPopUpButtonCell.DismissPopUp]: Dismisses the pop-up button’s menu by ordering its window out.
//   - [INSPopUpButtonCell.PerformClickWithFrameInView]: Displays the receiver’s menu and track mouse events in it.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell
type INSPopUpButtonCell interface {
	INSMenuItemCell
	NSMenuItemValidation

	// Topic: Initialization

	// Returns an [NSPopUpButtonCell] object initialized with the specified title.
	InitTextCellPullsDown(stringValue string, pullDown bool) NSPopUpButtonCell

	// Topic: Accessing menu attributes

	// A Boolean value that indicates the behavior of the button’s menu.
	PullsDown() bool
	SetPullsDown(value bool)
	// A Boolean value that indicates if the button automatically enables and disables its items every time a user event occurs.
	AutoenablesItems() bool
	SetAutoenablesItems(value bool)
	// The edge of the cell from which the menu should pop out when screen conditions are restrictive.
	PreferredEdge() foundation.NSRectEdge
	SetPreferredEdge(value foundation.NSRectEdge)
	// A Boolean value that indicates if the control uses an item from the menu for its own title.
	UsesItemFromMenu() bool
	SetUsesItemFromMenu(value bool)
	// A Boolean value that indicates if the pop-up button links the state of the selected menu item to the current selection.
	AltersStateOfSelectedItem() bool
	SetAltersStateOfSelectedItem(value bool)
	// The position of the arrow displayed on the button.
	ArrowPosition() NSPopUpArrowPosition
	SetArrowPosition(value NSPopUpArrowPosition)

	// Topic: Adding and removing items

	// Adds an item with the specified title to the end of the menu.
	AddItemWithTitle(title string)
	// Adds multiple items to the end of the menu.
	AddItemsWithTitles(itemTitles []string)
	// Inserts an item at the specified position in the menu.
	InsertItemWithTitleAtIndex(title string, index int)
	// Removes the item with the specified title from the menu.
	RemoveItemWithTitle(title string)
	// Removes the item at the specified index.
	RemoveItemAtIndex(index int)
	// Removes all items in the receiver’s item menu.
	RemoveAllItems()

	// Topic: Accessing the items

	// An array of [NSMenuItem](<doc://com.apple.appkit/documentation/AppKit/NSMenuItem>) objects that represent the items in the menu.
	ItemArray() []NSMenuItem
	// The number of items in the menu.
	NumberOfItems() int
	// Returns the index of the specified menu item.
	IndexOfItem(item INSMenuItem) int
	// Returns the index of the item with the specified title.
	IndexOfItemWithTitle(title string) int
	// Returns the index of the menu item with the specified tag.
	IndexOfItemWithTag(tag int) int
	// Returns the index of the menu item that holds the specified represented object.
	IndexOfItemWithRepresentedObject(obj objectivec.IObject) int
	// Returns the index of the menu item with the specified target and action.
	IndexOfItemWithTargetAndAction(target objectivec.IObject, actionSelector objc.SEL) int
	// Returns the menu item at the specified index.
	ItemAtIndex(index int) INSMenuItem
	// Returns the menu item with the specified title.
	ItemWithTitle(title string) INSMenuItem
	// The last item in the menu.
	LastItem() INSMenuItem

	// Topic: Dealing with selection

	// Selects the specified menu item.
	SelectItem(item INSMenuItem)
	// Selects the item in the menu at the specified index.
	SelectItemAtIndex(index int)
	// Selects the menu item with the specified tag.
	SelectItemWithTag(tag int) bool
	// Selects the item with the specified title.
	SelectItemWithTitle(title string)
	// The menu item last selected by the user.
	SelectedItem() INSMenuItem
	// The index of the item last selected by the user.
	IndexOfSelectedItem() int
	// Synchronizes the pop-up button’s displayed item with the currently selected menu item.
	SynchronizeTitleAndSelectedItem()

	// Topic: Title conveniences

	// Returns the title of the item at the specified index.
	ItemTitleAtIndex(index int) string
	// An array of [NSString] objects containing the titles of every item in the menu.
	ItemTitles() []string
	// The title of the item last selected by the user.
	TitleOfSelectedItem() string

	// Topic: Handling events and action messages

	// Sets up the receiver to display a menu.
	AttachPopUpWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView)
	// Dismisses the pop-up button’s menu by ordering its window out.
	DismissPopUp()
	// Displays the receiver’s menu and track mouse events in it.
	PerformClickWithFrameInView(frame corefoundation.CGRect, controlView INSView)
}

// Init initializes the instance.
func (p NSPopUpButtonCell) Init() NSPopUpButtonCell {
	rv := objc.Send[NSPopUpButtonCell](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPopUpButtonCell) Autorelease() NSPopUpButtonCell {
	rv := objc.Send[NSPopUpButtonCell](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPopUpButtonCell creates a new NSPopUpButtonCell instance.
func NewNSPopUpButtonCell() NSPopUpButtonCell {
	class := getNSPopUpButtonCellClass()
	rv := objc.Send[NSPopUpButtonCell](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSButtonCell/init(imageCell:)
func NewPopUpButtonCellImageCell(image INSImage) NSPopUpButtonCell {
	instance := getNSPopUpButtonCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initImageCell:"), image)
	return NSPopUpButtonCellFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSMenuItemCell/init(textCell:)
func NewPopUpButtonCellTextCell(string_ string) NSPopUpButtonCell {
	instance := getNSPopUpButtonCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:"), objc.String(string_))
	return NSPopUpButtonCellFromID(rv)
}

// Returns an [NSPopUpButtonCell] object initialized with the specified title.
//
// stringValue: The title of the first menu. You may specify an empty string if you do not
// want to add an initial menu item.
//
// pullDown: true if you want the receiver to display a pull-down menu; otherwise, false
// if you want it to display a pop-up menu.
//
// # Return Value
//
// An initialized [NSPopUpButtonCell] object, or `nil` if the object could not
// be initialized.
//
// # Discussion
//
// This menu item is assigned the default pop-up button action that displays
// the menu. To set the action and target, use the setAction: and setTarget:
// methods of the item’s corresponding [NSMenuItem] object.
//
// This method is the designated initializer of the class.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/init(textCell:pullsDown:)
func NewPopUpButtonCellTextCellPullsDown(stringValue string, pullDown bool) NSPopUpButtonCell {
	instance := getNSPopUpButtonCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initTextCell:pullsDown:"), objc.String(stringValue), pullDown)
	return NSPopUpButtonCellFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/init(coder:)
func NewPopUpButtonCellWithCoder(coder foundation.INSCoder) NSPopUpButtonCell {
	instance := getNSPopUpButtonCellClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSPopUpButtonCellFromID(rv)
}

// Returns an [NSPopUpButtonCell] object initialized with the specified title.
//
// stringValue: The title of the first menu. You may specify an empty string if you do not
// want to add an initial menu item.
//
// pullDown: true if you want the receiver to display a pull-down menu; otherwise, false
// if you want it to display a pop-up menu.
//
// # Return Value
//
// An initialized [NSPopUpButtonCell] object, or `nil` if the object could not
// be initialized.
//
// # Discussion
//
// This menu item is assigned the default pop-up button action that displays
// the menu. To set the action and target, use the setAction: and setTarget:
// methods of the item’s corresponding [NSMenuItem] object.
//
// This method is the designated initializer of the class.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/init(textCell:pullsDown:)
func (p NSPopUpButtonCell) InitTextCellPullsDown(stringValue string, pullDown bool) NSPopUpButtonCell {
	rv := objc.Send[NSPopUpButtonCell](p.ID, objc.Sel("initTextCell:pullsDown:"), objc.String(stringValue), pullDown)
	return rv
}

// Adds an item with the specified title to the end of the menu.
//
// title: The title of the new menu item. If an item with the same title already
// exists in the menu, the existing item is removed and the new one is added.
//
// # Discussion
//
// The menu item uses the pop-up button’s default action and target, but you
// can change these using the setAction: and setTarget: methods of the
// corresponding [NSMenuItem] object.
//
// Because this method searches for duplicate items, it should not be used if
// you are adding an item to an already populated menu with more than a few
// hundred items. In a situation like this, add items directly to the
// button’s menu instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/addItem(withTitle:)
func (p NSPopUpButtonCell) AddItemWithTitle(title string) {
	objc.Send[objc.ID](p.ID, objc.Sel("addItemWithTitle:"), objc.String(title))
}

// Adds multiple items to the end of the menu.
//
// itemTitles: An array of [NSString] objects containing the titles of the items you want
// to add. Each string in the array should be unique. If an item with the same
// title already exists in the menu, the existing item is removed and the new
// one is added.
//
// # Discussion
//
// The new menu items use the pop-up button’s default action and target, but
// you can change these using the setAction: and setTarget: methods of the
// corresponding [NSMenuItem] object.
//
// If you want to move an item, it’s better to invoke [RemoveItemWithTitle]
// explicitly and then call this method. After adding the items, this method
// uses the [SynchronizeTitleAndSelectedItem] method to make sure the item
// being displayed matches the currently selected item.
//
// Because this method searches for duplicate items, it should not be used if
// you are adding items to an already populated menu with more than a few
// hundred items. In a situation like this, add items directly to the
// receiver’s menu instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/addItems(withTitles:)
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
func (p NSPopUpButtonCell) AddItemsWithTitles(itemTitles []string) {
	objc.Send[objc.ID](p.ID, objc.Sel("addItemsWithTitles:"), objectivec.StringSliceToNSArray(itemTitles))
}

// Inserts an item at the specified position in the menu.
//
// title: The title of the new item. If an item with the same title already exists in
// the menu, the existing item is removed and the new one is added
//
// index: The zero-based index at which to insert the item. Specifying `0` inserts
// the item at the top of the menu.
//
// # Discussion
//
// The value in `index` must represent a valid position in the array. The menu
// item at `index` and all those that follow it are shifted down one slot to
// make room for the new menu item.
//
// This method assigns the pop-up button’s default action and target to the
// new menu item. Use the menu item’s setAction: and setTarget: methods to
// assign a new action and target.
//
// Because this method searches for duplicate items, it should not be used if
// you are adding an item to an already populated menu with more than a few
// hundred items. In a situation like this, add items directly to the
// button’s menu instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/insertItem(withTitle:at:)
func (p NSPopUpButtonCell) InsertItemWithTitleAtIndex(title string, index int) {
	objc.Send[objc.ID](p.ID, objc.Sel("insertItemWithTitle:atIndex:"), objc.String(title), index)
}

// Removes the item with the specified title from the menu.
//
// title: The title of the item you want to remove. If no menu item exists with the
// specified title, this method triggers an assertion.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/removeItem(withTitle:)
func (p NSPopUpButtonCell) RemoveItemWithTitle(title string) {
	objc.Send[objc.ID](p.ID, objc.Sel("removeItemWithTitle:"), objc.String(title))
}

// Removes the item at the specified index.
//
// index: The zero-based index indicating which item to remove. Specifying `0`
// removes the item at the top of the menu. The index must be valid and
// non-negative.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/removeItem(at:)
func (p NSPopUpButtonCell) RemoveItemAtIndex(index int) {
	objc.Send[objc.ID](p.ID, objc.Sel("removeItemAtIndex:"), index)
}

// Removes all items in the receiver’s item menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/removeAllItems()
func (p NSPopUpButtonCell) RemoveAllItems() {
	objc.Send[objc.ID](p.ID, objc.Sel("removeAllItems"))
}

// Returns the index of the specified menu item.
//
// item: The menu item whose index you want.
//
// # Return Value
//
// The index of the item or `-1` if no such item was found.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/index(of:)
func (p NSPopUpButtonCell) IndexOfItem(item INSMenuItem) int {
	rv := objc.Send[int](p.ID, objc.Sel("indexOfItem:"), item)
	return rv
}

// Returns the index of the item with the specified title.
//
// title: The title of the item you want. You must not pass `nil` for this parameter.
//
// # Return Value
//
// The index of the item or `-1` if no item with the specified title was
// found.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/indexOfItem(withTitle:)
func (p NSPopUpButtonCell) IndexOfItemWithTitle(title string) int {
	rv := objc.Send[int](p.ID, objc.Sel("indexOfItemWithTitle:"), objc.String(title))
	return rv
}

// Returns the index of the menu item with the specified tag.
//
// tag: The tag of the menu item you want.
//
// # Return Value
//
// The index of the item or `-1` if no item with the specified tag was found.
//
// # Discussion
//
// Tags are values your application assigns to an object to identify it. You
// can assign tags to menu items using the setTag: method of [NSMenuItem].
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/indexOfItem(withTag:)
func (p NSPopUpButtonCell) IndexOfItemWithTag(tag int) int {
	rv := objc.Send[int](p.ID, objc.Sel("indexOfItemWithTag:"), tag)
	return rv
}

// Returns the index of the menu item that holds the specified represented
// object.
//
// obj: The represented object associated with a menu item.
//
// # Return Value
//
// The index of the menu item that owns the specified object, or `-1` if no
// such menu item was found.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/indexOfItem(withRepresentedObject:)
func (p NSPopUpButtonCell) IndexOfItemWithRepresentedObject(obj objectivec.IObject) int {
	rv := objc.Send[int](p.ID, objc.Sel("indexOfItemWithRepresentedObject:"), obj)
	return rv
}

// Returns the index of the menu item with the specified target and action.
//
// target: The target object associated with the menu item.
//
// actionSelector: The action method associated with the menu item.
//
// # Return Value
//
// The index of the menu item, or `-1` if no menu item contains the specified
// target and action.
//
// # Discussion
//
// If you specify [NULL] for the `actionSelector` parameter, the index of the
// first menu item with the specified target is returned.
//
// The [NSPopUpButtonCell] class assigns a default action and target to each
// menu item, but you can change these values using the setAction: and
// setTarget: methods of [NSMenuItem].
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/indexOfItem(withTarget:andAction:)
func (p NSPopUpButtonCell) IndexOfItemWithTargetAndAction(target objectivec.IObject, actionSelector objc.SEL) int {
	rv := objc.Send[int](p.ID, objc.Sel("indexOfItemWithTarget:andAction:"), target, actionSelector)
	return rv
}

// Returns the menu item at the specified index.
//
// index: The index of the item you want. The specified index must refer to an
// existing menu item.
//
// # Return Value
//
// The menu item, or `nil` if no item exists at the specified index.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/item(at:)
func (p NSPopUpButtonCell) ItemAtIndex(index int) INSMenuItem {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("itemAtIndex:"), index)
	return NSMenuItemFromID(rv)
}

// Returns the menu item with the specified title.
//
// title: The title of the menu item you want.
//
// # Return Value
//
// The menu item, or `nil` if no item with the specified title exists in the
// menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/item(withTitle:)
func (p NSPopUpButtonCell) ItemWithTitle(title string) INSMenuItem {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("itemWithTitle:"), objc.String(title))
	return NSMenuItemFromID(rv)
}

// Selects the specified menu item.
//
// item: The menu item to select, or `nil` if you want to deselect all menu items.
//
// # Discussion
//
// By default, selecting or deselecting a menu item from a pop-up menu changes
// its state. Selecting a menu item from a pull-down menu does not
// automatically alter the state of the item. To disassociate the current
// selection from the state of menu items, set the [AltersStateOfSelectedItem]
// property to false.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/select(_:)
func (p NSPopUpButtonCell) SelectItem(item INSMenuItem) {
	objc.Send[objc.ID](p.ID, objc.Sel("selectItem:"), item)
}

// Selects the item in the menu at the specified index.
//
// index: The index of the item you want to select, or `-1` you want to deselect all
// menu items.
//
// # Discussion
//
// By default, selecting or deselecting a menu item from a pop-up menu changes
// its state. Selecting a menu item from a pull-down menu does not
// automatically alter the state of the item. To disassociate the current
// selection from the state of menu items, set the [AltersStateOfSelectedItem]
// property to false.
//
// Subclassers can override this method to catch all select calls.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/selectItem(at:)
func (p NSPopUpButtonCell) SelectItemAtIndex(index int) {
	objc.Send[objc.ID](p.ID, objc.Sel("selectItemAtIndex:"), index)
}

// Selects the menu item with the specified tag.
//
// tag: The tag of the item you want to select.
//
// # Return Value
//
// true if the item was successfully selected; otherwise, false.
//
// # Discussion
//
// If no item with the specified tag is found, this method returns false and
// leaves the menu state unchanged.
//
// You typically assign tags to menu items from Interface Builder, but you can
// also assign them programmatically using the setTag: method of [NSMenuItem].
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/selectItem(withTag:)
func (p NSPopUpButtonCell) SelectItemWithTag(tag int) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("selectItemWithTag:"), tag)
	return rv
}

// Selects the item with the specified title.
//
// title: The title of the item to select. If you specify an empty string, or a
// string that does not match the title of a menu item, this method deselects
// the currently selected item.
//
// # Discussion
//
// By default, selecting or deselecting a menu item changes its state. To
// disassociate the current selection from the state of menu items, set the
// [AltersStateOfSelectedItem] property to false.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/selectItem(withTitle:)
func (p NSPopUpButtonCell) SelectItemWithTitle(title string) {
	objc.Send[objc.ID](p.ID, objc.Sel("selectItemWithTitle:"), objc.String(title))
}

// Synchronizes the pop-up button’s displayed item with the currently
// selected menu item.
//
// # Discussion
//
// If no item is currently selected, this method synchronizes the pop-up
// buttons displayed item with the first menu item. If the pop-up button cell
// does not get its displayed item from a menu item, this method does nothing.
//
// For pull-down menus, this method sets the displayed item to the title first
// menu item.
//
// If the pop-up button’s menu does not contain any menu items, this method
// sets the pop-up button’s displayed item to `nil`, resulting in nothing
// being displayed in the control.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/synchronizeTitleAndSelectedItem()
func (p NSPopUpButtonCell) SynchronizeTitleAndSelectedItem() {
	objc.Send[objc.ID](p.ID, objc.Sel("synchronizeTitleAndSelectedItem"))
}

// Returns the title of the item at the specified index.
//
// index: The index of the item you want.
//
// # Return Value
//
// The title of the item, or an empty string if no item exists at the
// specified index.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/itemTitle(at:)
func (p NSPopUpButtonCell) ItemTitleAtIndex(index int) string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("itemTitleAtIndex:"), index)
	return foundation.NSStringFromID(rv).String()
}

// Sets up the receiver to display a menu.
//
// cellFrame: The cell’s rectangle, specified in points in the coordinate system of the
// view in the `controlView` parameter. The menu is attached to this
// rectangle.
//
// controlView: The view in which to display the pop-up button’s menu.
//
// # Discussion
//
// This call sets up the popup button cell to display a menu, which occurs in
// [PerformClickWithFrameInView]. This method sets the cell’s control view
// and then highlights and redraws the cell. It does not show the menu.
//
// This method also posts an [willPopUpNotification]. (The [NSPopUpButton]
// object sends a corresponding [willPopUpNotification].)
//
// You normally do not call this method explicitly. It is called by the
// Application Kit automatically when the menu for the pop-up button is to be
// displayed.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/attachPopUp(withFrame:in:)
//
// [willPopUpNotification]: https://developer.apple.com/documentation/AppKit/NSPopUpButton/willPopUpNotification
func (p NSPopUpButtonCell) AttachPopUpWithFrameInView(cellFrame corefoundation.CGRect, controlView INSView) {
	objc.Send[objc.ID](p.ID, objc.Sel("attachPopUpWithFrame:inView:"), cellFrame, controlView)
}

// Dismisses the pop-up button’s menu by ordering its window out.
//
// # Discussion
//
// If the pop-up button was not displaying its menu, this method does nothing.
//
// You normally do not call this method explicitly. It is called by the
// Application Kit automatically to dismiss the menu for the pop-up button.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/dismissPopUp()
func (p NSPopUpButtonCell) DismissPopUp() {
	objc.Send[objc.ID](p.ID, objc.Sel("dismissPopUp"))
}

// Displays the receiver’s menu and track mouse events in it.
//
// frame: The cell’s rectangle, specified in points in the coordinate system of the
// view in the `controlView` parameter.
//
// controlView: The view in which to display the pop-up button’s menu.
//
// # Discussion
//
// You normally do not call this method explicitly. It is called by the
// Application Kit automatically to handle events in the pop-up button.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/performClick(withFrame:in:)
func (p NSPopUpButtonCell) PerformClickWithFrameInView(frame corefoundation.CGRect, controlView INSView) {
	objc.Send[objc.ID](p.ID, objc.Sel("performClickWithFrame:inView:"), frame, controlView)
}

// Implemented to override the default action of enabling or disabling a
// specific menu item.
//
// menuItem: An [NSMenuItem] object that represents the menu item.
//
// # Return Value
//
// true to enable `menuItem`, false to disable it.
//
// # Discussion
//
// The object implementing this method must be the target of `menuItem`. You
// can determine which menu item `menuItem` is by querying it for its tag or
// action.
//
// The following example disables the menu item associated with the
// `nextRecord` action method when the selected line in a table view is the
// last one; conversely, it disables the menu item with `priorRecord` as its
// action method when the selected row is the first one in the table view.
// (The `countryOrRegionKeys` array contains names that appear in the table
// view.)
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemValidation/validateMenuItem(_:)
func (p NSPopUpButtonCell) ValidateMenuItem(menuItem INSMenuItem) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("validateMenuItem:"), menuItem)
	return rv
}

// A Boolean value that indicates the behavior of the button’s menu.
//
// # Discussion
//
// When the value of this property is true, the menu behaves like a pull-down
// menu; when the value is false, it behaves like a pop-up menu. If you use
// this property to change the menu type from a pop-up menu to a pull-down
// menu, and the cell alters the state of its selected items, the state of the
// currently selected item is set to [NSOffState] before the menu type is
// changed.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/pullsDown
//
// [NSOffState]: https://developer.apple.com/documentation/AppKit/NSOffState
func (p NSPopUpButtonCell) PullsDown() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("pullsDown"))
	return rv
}
func (p NSPopUpButtonCell) SetPullsDown(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setPullsDown:"), value)
}

// A Boolean value that indicates if the button automatically enables and
// disables its items every time a user event occurs.
//
// # Discussion
//
// When the value of this property is true, the button automatically enables
// and disables items. The default value is true. For more information about
// enabling and disabling menu items, see NSMenuValidation.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/autoenablesItems
func (p NSPopUpButtonCell) AutoenablesItems() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("autoenablesItems"))
	return rv
}
func (p NSPopUpButtonCell) SetAutoenablesItems(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAutoenablesItems:"), value)
}

// The edge of the cell from which the menu should pop out when screen
// conditions are restrictive.
//
// # Discussion
//
// At display time, if attaching the menu to the preferred edge would cause
// part of the menu to be obscured, the pop-up button may use a different
// edge. If no preferred edge is set, the pop-up button uses the bottom edge
// by default, which is [NSMaxYEdge] for flipped views or [NSMinYEdge] for
// unflipped views. Additional values for this property include [NSMinXEdge]
// and [NSMaxXEdge].
//
// The exact location of the arrow is determined by examining the value of
// this property and [ArrowPosition].
//
// - If the arrow position is [NSPopUpArrowAtCenter], the arrow stays in the
// center of the button and the value of this property determines which edge
// the arrow points to: [NSMinXEdge] points to the left, [NSMaxYEdge] points
// to the top, [NSMaxXEdge] points to the right, and [NSMinYEdge] points to
// the bottom. - If the arrow position is [NSPopUpArrowAtBottom], the value of
// this property determines which edge at which the arrow is placed:
// [NSMinXEdge] places the arrow at the center of the left side, pointing to
// the left, [NSMinYEdge] places the arrow at bottom right corner, pointing
// up, [NSMaxXEdge] places the arrow at the center of the right side, pointing
// to the right, and [NSMaxYEdge] places the arrow at the bottom right corner,
// pointing down.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/preferredEdge
//
// [NSMaxXEdge]: https://developer.apple.com/documentation/Foundation/NSRectEdge/NSMaxXEdge
// [NSMaxYEdge]: https://developer.apple.com/documentation/Foundation/NSRectEdge/NSMaxYEdge
// [NSMinXEdge]: https://developer.apple.com/documentation/Foundation/NSRectEdge/NSMinXEdge
// [NSMinYEdge]: https://developer.apple.com/documentation/Foundation/NSRectEdge/NSMinYEdge
func (p NSPopUpButtonCell) PreferredEdge() foundation.NSRectEdge {
	rv := objc.Send[foundation.NSRectEdge](p.ID, objc.Sel("preferredEdge"))
	return foundation.NSRectEdge(rv)
}
func (p NSPopUpButtonCell) SetPreferredEdge(value foundation.NSRectEdge) {
	objc.Send[struct{}](p.ID, objc.Sel("setPreferredEdge:"), value)
}

// A Boolean value that indicates if the control uses an item from the menu
// for its own title.
//
// # Discussion
//
// When the value of this property is true, a pull-down menu uses the title of
// the first menu item, and a pop-up menu uses the title of the currently
// selected menu (if no menu item is selected, the pop-up button displays no
// item and is drawn empty). When the value is false, the menu item set with
// [MenuItem] ([NSMenuItem]) is always displayed. The default value is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/usesItemFromMenu
func (p NSPopUpButtonCell) UsesItemFromMenu() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("usesItemFromMenu"))
	return rv
}
func (p NSPopUpButtonCell) SetUsesItemFromMenu(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setUsesItemFromMenu:"), value)
}

// A Boolean value that indicates if the pop-up button links the state of the
// selected menu item to the current selection.
//
// # Discussion
//
// When the value of this property is true (which is the default value), the
// state of the selected item is set to [NSOnState]. When the value of this
// property is false, the items in the menu are left alone. When you change
// the value of this property, the state of the currently selected item is
// updated appropriately.
//
// Note that this property affects only pop-up buttons (it is ignored for
// pull-down menus).
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/altersStateOfSelectedItem
//
// [NSOnState]: https://developer.apple.com/documentation/AppKit/NSOnState
func (p NSPopUpButtonCell) AltersStateOfSelectedItem() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("altersStateOfSelectedItem"))
	return rv
}
func (p NSPopUpButtonCell) SetAltersStateOfSelectedItem(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAltersStateOfSelectedItem:"), value)
}

// The position of the arrow displayed on the button.
//
// # Discussion
//
// When the value of this property is [NSPopUpNoArrow], the control displays
// no arrow. [NSPopUpArrowAtCenter] displays the arrow centered horizontally
// within the cell and [NSPopUpArrowAtBottom] displays the arrow at the edge
// of the cell. This property is used with [PreferredEdge] to determine the
// exact location and orientation of the arrow.
//
// This property applies to only bezel style and borderless pop-up buttons.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/arrowPosition
func (p NSPopUpButtonCell) ArrowPosition() NSPopUpArrowPosition {
	rv := objc.Send[NSPopUpArrowPosition](p.ID, objc.Sel("arrowPosition"))
	return NSPopUpArrowPosition(rv)
}
func (p NSPopUpButtonCell) SetArrowPosition(value NSPopUpArrowPosition) {
	objc.Send[struct{}](p.ID, objc.Sel("setArrowPosition:"), value)
}

// An array of [NSMenuItem] objects that represent the items in the menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/itemArray
func (p NSPopUpButtonCell) ItemArray() []NSMenuItem {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("itemArray"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSMenuItem {
		return NSMenuItemFromID(id)
	})
}

// The number of items in the menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/numberOfItems
func (p NSPopUpButtonCell) NumberOfItems() int {
	rv := objc.Send[int](p.ID, objc.Sel("numberOfItems"))
	return rv
}

// The last item in the menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/lastItem
func (p NSPopUpButtonCell) LastItem() INSMenuItem {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("lastItem"))
	return NSMenuItemFromID(objc.ID(rv))
}

// The menu item last selected by the user.
//
// # Discussion
//
// The value of this property is the menu item that is currently selected, or
// `nil` if no item is selected. The last selected menu item is the one that
// was highlighted when the user released the mouse button. It is possible for
// a pull-down menu’s selected item to be its first item.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/selectedItem
func (p NSPopUpButtonCell) SelectedItem() INSMenuItem {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("selectedItem"))
	return NSMenuItemFromID(objc.ID(rv))
}

// The index of the item last selected by the user.
//
// # Discussion
//
// The value of this property is the index of the selected item, or `-1` if no
// item is selected.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/indexOfSelectedItem
func (p NSPopUpButtonCell) IndexOfSelectedItem() int {
	rv := objc.Send[int](p.ID, objc.Sel("indexOfSelectedItem"))
	return rv
}

// An array of [NSString] objects containing the titles of every item in the
// menu.
//
// # Discussion
//
// The titles appear in the order in which the items appear in the menu. If
// the menu contains separator items, the array contains an empty string
// (`@""`) for each separator item.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/itemTitles
func (p NSPopUpButtonCell) ItemTitles() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("itemTitles"))
	return objc.ConvertSliceToStrings(rv)
}

// The title of the item last selected by the user.
//
// # Discussion
//
// The value of this property is the title of the selected menu item, or an
// empty string if no item is selected.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButtonCell/titleOfSelectedItem
func (p NSPopUpButtonCell) TitleOfSelectedItem() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("titleOfSelectedItem"))
	return foundation.NSStringFromID(rv).String()
}

// Protocol methods for NSMenuItemValidation
