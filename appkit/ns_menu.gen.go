// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMenu] class.
var (
	_NSMenuClass     NSMenuClass
	_NSMenuClassOnce sync.Once
)

func getNSMenuClass() NSMenuClass {
	_NSMenuClassOnce.Do(func() {
		_NSMenuClass = NSMenuClass{class: objc.GetClass("NSMenu")}
	})
	return _NSMenuClass
}

// GetNSMenuClass returns the class object for NSMenu.
func GetNSMenuClass() NSMenuClass {
	return getNSMenuClass()
}

type NSMenuClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSMenuClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMenuClass) Alloc() NSMenu {
	rv := objc.Send[NSMenu](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that manages an app’s menus.
//
// # Managing the Menu Bar
//
//   - [NSMenu.MenuBarHeight]: The menu bar height for the main menu in pixels.
//
// # Creating an NSMenu Object
//
//   - [NSMenu.InitWithTitle]: Initializes and returns a menu having the specified title and with autoenabling of menu items turned on.
//   - [NSMenu.InitWithCoder]
//
// # Adding and Removing Menu Items
//
//   - [NSMenu.InsertItemAtIndex]: Inserts a menu item into the menu at a specific location.
//   - [NSMenu.InsertItemWithTitleActionKeyEquivalentAtIndex]: Creates and adds a menu item at a specified location in the menu.
//   - [NSMenu.AddItem]: Adds a menu item to the end of the menu.
//   - [NSMenu.AddItemWithTitleActionKeyEquivalent]: Creates a new menu item and adds it to the end of the menu.
//   - [NSMenu.RemoveItem]: Removes a menu item from the menu.
//   - [NSMenu.RemoveItemAtIndex]: Removes the menu item at a specified location in the menu.
//   - [NSMenu.ItemChanged]: Invoked when a menu item is modified visually (for example, its title changes).
//   - [NSMenu.RemoveAllItems]: Removes all the menu items in the menu.
//
// # Finding Menu Items
//
//   - [NSMenu.ItemWithTag]: Returns the first menu item in the menu with the specified tag.
//   - [NSMenu.ItemWithTitle]: Returns the first menu item in the menu with a specified title.
//   - [NSMenu.ItemAtIndex]: Returns the menu item at a specific location of the menu.
//   - [NSMenu.NumberOfItems]: The number of menu items in the menu, including separator items.
//   - [NSMenu.ItemArray]: An array containing the menu items in the menu.
//   - [NSMenu.SetItemArray]
//
// # Finding Indices of Menu Items
//
//   - [NSMenu.IndexOfItem]: Returns the index identifying the location of a specified menu item in the menu.
//   - [NSMenu.IndexOfItemWithTitle]: Returns the index of the first menu item in the menu that has a specified title.
//   - [NSMenu.IndexOfItemWithTag]: Returns the index of the first menu item in the menu identified by a tag.
//   - [NSMenu.IndexOfItemWithTargetAndAction]: Returns the index of the first menu item in the menu that has a specified action and target.
//   - [NSMenu.IndexOfItemWithRepresentedObject]: Returns the index of the first menu item in the menu that has a given represented object.
//   - [NSMenu.IndexOfItemWithSubmenu]: Returns the index of the menu item in the menu with the given submenu.
//
// # Managing Submenus
//
//   - [NSMenu.SetSubmenuForItem]: Assigns a menu to be a submenu of the menu controlled by a given menu item.
//   - [NSMenu.SubmenuAction]: The action method assigned to menu items that open submenus.
//   - [NSMenu.Supermenu]: The parent menu that contains the menu as a submenu.
//   - [NSMenu.SetSupermenu]
//
// # Enabling and Disabling Menu Items
//
//   - [NSMenu.AutoenablesItems]: Indicates whether the menu automatically enables and disables its menu items.
//   - [NSMenu.SetAutoenablesItems]
//   - [NSMenu.Update]: Enables or disables the menu items of the menu based on the NSMenuValidation informal protocol and sizes the menu to fit its current menu items if necessary.
//
// # Getting and Setting the Menu Font
//
//   - [NSMenu.Font]: The font of the menu and its submenus.
//   - [NSMenu.SetFont]
//
// # Handling Keyboard Equivalents
//
//   - [NSMenu.PerformKeyEquivalent]: Performs the action for the menu item that corresponds to the given key equivalent.
//
// # Simulating Mouse Clicks
//
//   - [NSMenu.PerformActionForItemAtIndex]: Causes the application to send the action message of a specified menu item to its target.
//
// # Managing the Title
//
//   - [NSMenu.Title]: The title of the menu.
//   - [NSMenu.SetTitle]
//
// # Selecting Items
//
//   - [NSMenu.SelectedItems]: The menu items that are currently selected.
//   - [NSMenu.SetSelectedItems]
//   - [NSMenu.SelectionMode]: The selection mode of the menu.
//   - [NSMenu.SetSelectionMode]
//
// # Configuring Menu Size
//
//   - [NSMenu.MinimumWidth]: The minimum width of the menu in screen coordinates.
//   - [NSMenu.SetMinimumWidth]
//   - [NSMenu.Size]: The size of the menu in screen coordinates
//
// # Getting Menu Properties
//
//   - [NSMenu.PropertiesToUpdate]: The available properties for the menu.
//
// # Managing Presentation Styles
//
//   - [NSMenu.PresentationStyle]: The presentation style of the menu.
//   - [NSMenu.SetPresentationStyle]
//
// # Displaying Contextual Menus
//
//   - [NSMenu.AllowsContextMenuPlugIns]: Indicates whether the pop-up menu allows appending of contextual menu plug-in items.
//   - [NSMenu.SetAllowsContextMenuPlugIns]
//
// # Displaying Context-Sensitive Help
//
//   - [NSMenu.PopUpMenuPositioningItemAtLocationInView]: Pops up the menu at the specified location.
//
// # Managing Display of the State Column
//
//   - [NSMenu.ShowsStateColumn]: Indicates whether the menu displays the state column.
//   - [NSMenu.SetShowsStateColumn]
//
// # Handling Highlighting
//
//   - [NSMenu.HighlightedItem]: Indicates the currently highlighted item in the menu.
//
// # Managing the User Interface
//
//   - [NSMenu.UserInterfaceLayoutDirection]: Configures the layout direction of menu items in the menu.
//   - [NSMenu.SetUserInterfaceLayoutDirection]
//
// # Managing the Delegate
//
//   - [NSMenu.Delegate]: The delegate of the menu.
//   - [NSMenu.SetDelegate]
//
// # Handling Tracking
//
//   - [NSMenu.CancelTracking]: Dismisses the menu and ends all menu tracking.
//   - [NSMenu.CancelTrackingWithoutAnimation]: Dismisses the menu and ends all menu tracking without displaying the associated animation.
//
// # Instance Properties
//
//   - [NSMenu.AutomaticallyInsertsWritingToolsItems]
//   - [NSMenu.SetAutomaticallyInsertsWritingToolsItems]
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu
type NSMenu struct {
	objectivec.Object
}

// NSMenuFromID constructs a [NSMenu] from an objc.ID.
//
// An object that manages an app’s menus.
func NSMenuFromID(id objc.ID) NSMenu {
	return NSMenu{objectivec.Object{ID: id}}
}

// NOTE: NSMenu adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMenu] class.
//
// # Managing the Menu Bar
//
//   - [INSMenu.MenuBarHeight]: The menu bar height for the main menu in pixels.
//
// # Creating an NSMenu Object
//
//   - [INSMenu.InitWithTitle]: Initializes and returns a menu having the specified title and with autoenabling of menu items turned on.
//   - [INSMenu.InitWithCoder]
//
// # Adding and Removing Menu Items
//
//   - [INSMenu.InsertItemAtIndex]: Inserts a menu item into the menu at a specific location.
//   - [INSMenu.InsertItemWithTitleActionKeyEquivalentAtIndex]: Creates and adds a menu item at a specified location in the menu.
//   - [INSMenu.AddItem]: Adds a menu item to the end of the menu.
//   - [INSMenu.AddItemWithTitleActionKeyEquivalent]: Creates a new menu item and adds it to the end of the menu.
//   - [INSMenu.RemoveItem]: Removes a menu item from the menu.
//   - [INSMenu.RemoveItemAtIndex]: Removes the menu item at a specified location in the menu.
//   - [INSMenu.ItemChanged]: Invoked when a menu item is modified visually (for example, its title changes).
//   - [INSMenu.RemoveAllItems]: Removes all the menu items in the menu.
//
// # Finding Menu Items
//
//   - [INSMenu.ItemWithTag]: Returns the first menu item in the menu with the specified tag.
//   - [INSMenu.ItemWithTitle]: Returns the first menu item in the menu with a specified title.
//   - [INSMenu.ItemAtIndex]: Returns the menu item at a specific location of the menu.
//   - [INSMenu.NumberOfItems]: The number of menu items in the menu, including separator items.
//   - [INSMenu.ItemArray]: An array containing the menu items in the menu.
//   - [INSMenu.SetItemArray]
//
// # Finding Indices of Menu Items
//
//   - [INSMenu.IndexOfItem]: Returns the index identifying the location of a specified menu item in the menu.
//   - [INSMenu.IndexOfItemWithTitle]: Returns the index of the first menu item in the menu that has a specified title.
//   - [INSMenu.IndexOfItemWithTag]: Returns the index of the first menu item in the menu identified by a tag.
//   - [INSMenu.IndexOfItemWithTargetAndAction]: Returns the index of the first menu item in the menu that has a specified action and target.
//   - [INSMenu.IndexOfItemWithRepresentedObject]: Returns the index of the first menu item in the menu that has a given represented object.
//   - [INSMenu.IndexOfItemWithSubmenu]: Returns the index of the menu item in the menu with the given submenu.
//
// # Managing Submenus
//
//   - [INSMenu.SetSubmenuForItem]: Assigns a menu to be a submenu of the menu controlled by a given menu item.
//   - [INSMenu.SubmenuAction]: The action method assigned to menu items that open submenus.
//   - [INSMenu.Supermenu]: The parent menu that contains the menu as a submenu.
//   - [INSMenu.SetSupermenu]
//
// # Enabling and Disabling Menu Items
//
//   - [INSMenu.AutoenablesItems]: Indicates whether the menu automatically enables and disables its menu items.
//   - [INSMenu.SetAutoenablesItems]
//   - [INSMenu.Update]: Enables or disables the menu items of the menu based on the NSMenuValidation informal protocol and sizes the menu to fit its current menu items if necessary.
//
// # Getting and Setting the Menu Font
//
//   - [INSMenu.Font]: The font of the menu and its submenus.
//   - [INSMenu.SetFont]
//
// # Handling Keyboard Equivalents
//
//   - [INSMenu.PerformKeyEquivalent]: Performs the action for the menu item that corresponds to the given key equivalent.
//
// # Simulating Mouse Clicks
//
//   - [INSMenu.PerformActionForItemAtIndex]: Causes the application to send the action message of a specified menu item to its target.
//
// # Managing the Title
//
//   - [INSMenu.Title]: The title of the menu.
//   - [INSMenu.SetTitle]
//
// # Selecting Items
//
//   - [INSMenu.SelectedItems]: The menu items that are currently selected.
//   - [INSMenu.SetSelectedItems]
//   - [INSMenu.SelectionMode]: The selection mode of the menu.
//   - [INSMenu.SetSelectionMode]
//
// # Configuring Menu Size
//
//   - [INSMenu.MinimumWidth]: The minimum width of the menu in screen coordinates.
//   - [INSMenu.SetMinimumWidth]
//   - [INSMenu.Size]: The size of the menu in screen coordinates
//
// # Getting Menu Properties
//
//   - [INSMenu.PropertiesToUpdate]: The available properties for the menu.
//
// # Managing Presentation Styles
//
//   - [INSMenu.PresentationStyle]: The presentation style of the menu.
//   - [INSMenu.SetPresentationStyle]
//
// # Displaying Contextual Menus
//
//   - [INSMenu.AllowsContextMenuPlugIns]: Indicates whether the pop-up menu allows appending of contextual menu plug-in items.
//   - [INSMenu.SetAllowsContextMenuPlugIns]
//
// # Displaying Context-Sensitive Help
//
//   - [INSMenu.PopUpMenuPositioningItemAtLocationInView]: Pops up the menu at the specified location.
//
// # Managing Display of the State Column
//
//   - [INSMenu.ShowsStateColumn]: Indicates whether the menu displays the state column.
//   - [INSMenu.SetShowsStateColumn]
//
// # Handling Highlighting
//
//   - [INSMenu.HighlightedItem]: Indicates the currently highlighted item in the menu.
//
// # Managing the User Interface
//
//   - [INSMenu.UserInterfaceLayoutDirection]: Configures the layout direction of menu items in the menu.
//   - [INSMenu.SetUserInterfaceLayoutDirection]
//
// # Managing the Delegate
//
//   - [INSMenu.Delegate]: The delegate of the menu.
//   - [INSMenu.SetDelegate]
//
// # Handling Tracking
//
//   - [INSMenu.CancelTracking]: Dismisses the menu and ends all menu tracking.
//   - [INSMenu.CancelTrackingWithoutAnimation]: Dismisses the menu and ends all menu tracking without displaying the associated animation.
//
// # Instance Properties
//
//   - [INSMenu.AutomaticallyInsertsWritingToolsItems]
//   - [INSMenu.SetAutomaticallyInsertsWritingToolsItems]
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu
type INSMenu interface {
	objectivec.IObject

	// Topic: Managing the Menu Bar

	// The menu bar height for the main menu in pixels.
	MenuBarHeight() float64

	// Topic: Creating an NSMenu Object

	// Initializes and returns a menu having the specified title and with autoenabling of menu items turned on.
	InitWithTitle(title string) NSMenu
	InitWithCoder(coder foundation.INSCoder) NSMenu

	// Topic: Adding and Removing Menu Items

	// Inserts a menu item into the menu at a specific location.
	InsertItemAtIndex(newItem INSMenuItem, index int)
	// Creates and adds a menu item at a specified location in the menu.
	InsertItemWithTitleActionKeyEquivalentAtIndex(string_ string, selector objc.SEL, charCode string, index int) INSMenuItem
	// Adds a menu item to the end of the menu.
	AddItem(newItem INSMenuItem)
	// Creates a new menu item and adds it to the end of the menu.
	AddItemWithTitleActionKeyEquivalent(string_ string, selector objc.SEL, charCode string) INSMenuItem
	// Removes a menu item from the menu.
	RemoveItem(item INSMenuItem)
	// Removes the menu item at a specified location in the menu.
	RemoveItemAtIndex(index int)
	// Invoked when a menu item is modified visually (for example, its title changes).
	ItemChanged(item INSMenuItem)
	// Removes all the menu items in the menu.
	RemoveAllItems()

	// Topic: Finding Menu Items

	// Returns the first menu item in the menu with the specified tag.
	ItemWithTag(tag int) INSMenuItem
	// Returns the first menu item in the menu with a specified title.
	ItemWithTitle(title string) INSMenuItem
	// Returns the menu item at a specific location of the menu.
	ItemAtIndex(index int) INSMenuItem
	// The number of menu items in the menu, including separator items.
	NumberOfItems() int
	// An array containing the menu items in the menu.
	ItemArray() []NSMenuItem
	SetItemArray(value []NSMenuItem)

	// Topic: Finding Indices of Menu Items

	// Returns the index identifying the location of a specified menu item in the menu.
	IndexOfItem(item INSMenuItem) int
	// Returns the index of the first menu item in the menu that has a specified title.
	IndexOfItemWithTitle(title string) int
	// Returns the index of the first menu item in the menu identified by a tag.
	IndexOfItemWithTag(tag int) int
	// Returns the index of the first menu item in the menu that has a specified action and target.
	IndexOfItemWithTargetAndAction(target objectivec.IObject, actionSelector objc.SEL) int
	// Returns the index of the first menu item in the menu that has a given represented object.
	IndexOfItemWithRepresentedObject(object objectivec.IObject) int
	// Returns the index of the menu item in the menu with the given submenu.
	IndexOfItemWithSubmenu(submenu INSMenu) int

	// Topic: Managing Submenus

	// Assigns a menu to be a submenu of the menu controlled by a given menu item.
	SetSubmenuForItem(menu INSMenu, item INSMenuItem)
	// The action method assigned to menu items that open submenus.
	SubmenuAction(sender objectivec.IObject)
	// The parent menu that contains the menu as a submenu.
	Supermenu() INSMenu
	SetSupermenu(value INSMenu)

	// Topic: Enabling and Disabling Menu Items

	// Indicates whether the menu automatically enables and disables its menu items.
	AutoenablesItems() bool
	SetAutoenablesItems(value bool)
	// Enables or disables the menu items of the menu based on the NSMenuValidation informal protocol and sizes the menu to fit its current menu items if necessary.
	Update()

	// Topic: Getting and Setting the Menu Font

	// The font of the menu and its submenus.
	Font() NSFont
	SetFont(value NSFont)

	// Topic: Handling Keyboard Equivalents

	// Performs the action for the menu item that corresponds to the given key equivalent.
	PerformKeyEquivalent(event INSEvent) bool

	// Topic: Simulating Mouse Clicks

	// Causes the application to send the action message of a specified menu item to its target.
	PerformActionForItemAtIndex(index int)

	// Topic: Managing the Title

	// The title of the menu.
	Title() string
	SetTitle(value string)

	// Topic: Selecting Items

	// The menu items that are currently selected.
	SelectedItems() []NSMenuItem
	SetSelectedItems(value []NSMenuItem)
	// The selection mode of the menu.
	SelectionMode() NSMenuSelectionMode
	SetSelectionMode(value NSMenuSelectionMode)

	// Topic: Configuring Menu Size

	// The minimum width of the menu in screen coordinates.
	MinimumWidth() float64
	SetMinimumWidth(value float64)
	// The size of the menu in screen coordinates
	Size() corefoundation.CGSize

	// Topic: Getting Menu Properties

	// The available properties for the menu.
	PropertiesToUpdate() NSMenuProperties

	// Topic: Managing Presentation Styles

	// The presentation style of the menu.
	PresentationStyle() NSMenuPresentationStyle
	SetPresentationStyle(value NSMenuPresentationStyle)

	// Topic: Displaying Contextual Menus

	// Indicates whether the pop-up menu allows appending of contextual menu plug-in items.
	AllowsContextMenuPlugIns() bool
	SetAllowsContextMenuPlugIns(value bool)

	// Topic: Displaying Context-Sensitive Help

	// Pops up the menu at the specified location.
	PopUpMenuPositioningItemAtLocationInView(item INSMenuItem, location corefoundation.CGPoint, view INSView) bool

	// Topic: Managing Display of the State Column

	// Indicates whether the menu displays the state column.
	ShowsStateColumn() bool
	SetShowsStateColumn(value bool)

	// Topic: Handling Highlighting

	// Indicates the currently highlighted item in the menu.
	HighlightedItem() INSMenuItem

	// Topic: Managing the User Interface

	// Configures the layout direction of menu items in the menu.
	UserInterfaceLayoutDirection() NSUserInterfaceLayoutDirection
	SetUserInterfaceLayoutDirection(value NSUserInterfaceLayoutDirection)

	// Topic: Managing the Delegate

	// The delegate of the menu.
	Delegate() NSMenuDelegate
	SetDelegate(value NSMenuDelegate)

	// Topic: Handling Tracking

	// Dismisses the menu and ends all menu tracking.
	CancelTracking()
	// Dismisses the menu and ends all menu tracking without displaying the associated animation.
	CancelTrackingWithoutAnimation()

	// Topic: Instance Properties

	AutomaticallyInsertsWritingToolsItems() bool
	SetAutomaticallyInsertsWritingToolsItems(value bool)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (m NSMenu) Init() NSMenu {
	rv := objc.Send[NSMenu](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMenu) Autorelease() NSMenu {
	rv := objc.Send[NSMenu](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMenu creates a new NSMenu instance.
func NewNSMenu() NSMenu {
	class := getNSMenuClass()
	rv := objc.Send[NSMenu](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSMenu/init(coder:)
func NewMenuWithCoder(coder foundation.INSCoder) NSMenu {
	instance := getNSMenuClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSMenuFromID(rv)
}

// Initializes and returns a menu having the specified title and with
// autoenabling of menu items turned on.
//
// title: The title to assign to the menu.
//
// # Return Value
//
// The initialized [NSMenu] object or `nil` if the object could not be
// initialized.
//
// # Discussion
//
// This method is the designated initializer for the class.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/init(title:)
func NewMenuWithTitle(title string) NSMenu {
	instance := getNSMenuClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTitle:"), objc.String(title))
	return NSMenuFromID(rv)
}

// Initializes and returns a menu having the specified title and with
// autoenabling of menu items turned on.
//
// title: The title to assign to the menu.
//
// # Return Value
//
// The initialized [NSMenu] object or `nil` if the object could not be
// initialized.
//
// # Discussion
//
// This method is the designated initializer for the class.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/init(title:)
func (m NSMenu) InitWithTitle(title string) NSMenu {
	rv := objc.Send[NSMenu](m.ID, objc.Sel("initWithTitle:"), objc.String(title))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSMenu/init(coder:)
func (m NSMenu) InitWithCoder(coder foundation.INSCoder) NSMenu {
	rv := objc.Send[NSMenu](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Inserts a menu item into the menu at a specific location.
//
// newItem: An object conforming to the [NSMenuItem] protocol that represents a menu
// item.
//
// index: An integer index identifying the location of the menu item in the menu.
//
// # Discussion
//
// This method posts an [didAddItemNotification], allowing interested
// observers to update as appropriate. This method is a primitive method. All
// item-addition methods end up calling this method, so this is where you
// should implement custom behavior on adding new items to a menu in a custom
// subclass. If the menu item already exists in another menu, it is not
// inserted and the method raises an exception of type
// [internalInconsistencyException].
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/insertItem(_:at:)
//
// [didAddItemNotification]: https://developer.apple.com/documentation/AppKit/NSMenu/didAddItemNotification
// [internalInconsistencyException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/internalInconsistencyException
func (m NSMenu) InsertItemAtIndex(newItem INSMenuItem, index int) {
	objc.Send[objc.ID](m.ID, objc.Sel("insertItem:atIndex:"), newItem, index)
}

// Creates and adds a menu item at a specified location in the menu.
//
// string: A string to be made the title of the menu item.
//
// selector: The action-message selector to assign to the menu item.
//
// charCode: A string identifying the key to use as a key equivalent for the menu item.
// If you do not want the menu item to have a key equivalent, `keyEquiv`
// should be an empty string (`@""`) and not `nil`.
//
// index: An integer index identifying the location of the menu item in the menu.
//
// # Return Value
//
// The new menu item (an object conforming to the NSMenuItem protocol) or
// `nil` if the item could not be created
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/insertItem(withTitle:action:keyEquivalent:at:)
func (m NSMenu) InsertItemWithTitleActionKeyEquivalentAtIndex(string_ string, selector objc.SEL, charCode string, index int) INSMenuItem {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("insertItemWithTitle:action:keyEquivalent:atIndex:"), objc.String(string_), selector, objc.String(charCode), index)
	return NSMenuItemFromID(rv)
}

// Adds a menu item to the end of the menu.
//
// newItem: The menu item (an object conforming to the NSMenuItem protocol) to add to
// the menu.
//
// # Discussion
//
// This method invokes [NSMenu.InsertItemAtIndex]. Thus, the menu does not
// accept the menu item if it already belongs to another menu. After adding
// the menu item, the menu updates itself.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/addItem(_:)
func (m NSMenu) AddItem(newItem INSMenuItem) {
	objc.Send[objc.ID](m.ID, objc.Sel("addItem:"), newItem)
}

// Creates a new menu item and adds it to the end of the menu.
//
// string: A string to be made the title of the menu item.
//
// selector: The action-message selector to assign to the menu item.
//
// charCode: A string identifying the key to use as a key equivalent for the menu item.
// If you do not want the menu item to have a key equivalent, `keyEquiv`
// should be an empty string (`@""`) and not `nil`.
//
// # Return Value
//
// The created menu item (an object conforming to the NSMenuItem protocol) or
// `nil` if the object couldn’t be created.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/addItem(withTitle:action:keyEquivalent:)
func (m NSMenu) AddItemWithTitleActionKeyEquivalent(string_ string, selector objc.SEL, charCode string) INSMenuItem {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("addItemWithTitle:action:keyEquivalent:"), objc.String(string_), selector, objc.String(charCode))
	return NSMenuItemFromID(rv)
}

// Removes a menu item from the menu.
//
// item: The menu item to remove.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/removeItem(_:)
func (m NSMenu) RemoveItem(item INSMenuItem) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeItem:"), item)
}

// Removes the menu item at a specified location in the menu.
//
// index: An integer index identifying the menu item.
//
// # Discussion
//
// After it removes the menu item, this method posts an
// [didRemoveItemNotification].
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/removeItem(at:)
//
// [didRemoveItemNotification]: https://developer.apple.com/documentation/AppKit/NSMenu/didRemoveItemNotification
func (m NSMenu) RemoveItemAtIndex(index int) {
	objc.Send[objc.ID](m.ID, objc.Sel("removeItemAtIndex:"), index)
}

// Invoked when a menu item is modified visually (for example, its title
// changes).
//
// item: The menu item that has visually changed.
//
// # Discussion
//
// This method is not called for changes involving the menu item’s action,
// target, represented object, or tag. Posts an [didChangeItemNotification].
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/itemChanged(_:)
//
// [didChangeItemNotification]: https://developer.apple.com/documentation/AppKit/NSMenu/didChangeItemNotification
func (m NSMenu) ItemChanged(item INSMenuItem) {
	objc.Send[objc.ID](m.ID, objc.Sel("itemChanged:"), item)
}

// Removes all the menu items in the menu.
//
// # Discussion
//
// This method is more efficient than removing menu items individually.
//
// Unlike the other remove methods, this method does not post
// [didChangeItemNotification] notifications.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/removeAllItems()
//
// [didChangeItemNotification]: https://developer.apple.com/documentation/AppKit/NSMenu/didChangeItemNotification
func (m NSMenu) RemoveAllItems() {
	objc.Send[objc.ID](m.ID, objc.Sel("removeAllItems"))
}

// Returns the first menu item in the menu with the specified tag.
//
// tag: A numeric tag associated with a menu item.
//
// # Return Value
//
// The found menu item (an object conforming to the NSMenuItem protocol) or
// `nil` if the object couldn’t be found.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/item(withTag:)
func (m NSMenu) ItemWithTag(tag int) INSMenuItem {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("itemWithTag:"), tag)
	return NSMenuItemFromID(rv)
}

// Returns the first menu item in the menu with a specified title.
//
// title: The title of a menu item.
//
// # Return Value
//
// The found menu item (an object conforming to the NSMenuItem protocol) or
// `nil` if the object couldn’t be found.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/item(withTitle:)
func (m NSMenu) ItemWithTitle(title string) INSMenuItem {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("itemWithTitle:"), objc.String(title))
	return NSMenuItemFromID(rv)
}

// Returns the menu item at a specific location of the menu.
//
// index: An integer index locating a menu item in a menu.
//
// # Return Value
//
// The found menu item (an object conforming to the NSMenuItem protocol) or
// `nil` if the object couldn’t be found.
//
// # Discussion
//
// This method raises an exception if `index` is out of bounds.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/item(at:)
func (m NSMenu) ItemAtIndex(index int) INSMenuItem {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("itemAtIndex:"), index)
	return NSMenuItemFromID(rv)
}

// Returns the index identifying the location of a specified menu item in the
// menu.
//
// item: A menu item—that is an object conforming to the NSMenuItem protocol.
//
// # Return Value
//
// The integer index of the menu item or, if no such menu item is in the menu,
// –1.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/index(of:)
func (m NSMenu) IndexOfItem(item INSMenuItem) int {
	rv := objc.Send[int](m.ID, objc.Sel("indexOfItem:"), item)
	return rv
}

// Returns the index of the first menu item in the menu that has a specified
// title.
//
// title: The title of a menu item in the menu.
//
// # Return Value
//
// The integer index of the menu item or, if no such menu item is in the menu,
// –1.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/indexOfItem(withTitle:)
func (m NSMenu) IndexOfItemWithTitle(title string) int {
	rv := objc.Send[int](m.ID, objc.Sel("indexOfItemWithTitle:"), objc.String(title))
	return rv
}

// Returns the index of the first menu item in the menu identified by a tag.
//
// tag: An integer tag associated with the menu item of the menu.
//
// # Return Value
//
// The integer index of the menu item or, if no such menu item is in the menu,
// –1.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/indexOfItem(withTag:)
func (m NSMenu) IndexOfItemWithTag(tag int) int {
	rv := objc.Send[int](m.ID, objc.Sel("indexOfItemWithTag:"), tag)
	return rv
}

// Returns the index of the first menu item in the menu that has a specified
// action and target.
//
// target: An object that is set as the target of a menu item of the menu.
//
// actionSelector: A selector identifying an action method. If `actionSelector` is [NULL], the
// first menu item in the menu that has target `anObject` is returned.
//
// # Return Value
//
// The integer index of the menu item or, if no such menu item is in the menu,
// –1.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/indexOfItem(withTarget:andAction:)
func (m NSMenu) IndexOfItemWithTargetAndAction(target objectivec.IObject, actionSelector objc.SEL) int {
	rv := objc.Send[int](m.ID, objc.Sel("indexOfItemWithTarget:andAction:"), target, actionSelector)
	return rv
}

// Returns the index of the first menu item in the menu that has a given
// represented object.
//
// object: A represented object of the menu.
//
// # Return Value
//
// The integer index of the menu item or, if no such menu item is in the menu,
// –1.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/indexOfItem(withRepresentedObject:)
func (m NSMenu) IndexOfItemWithRepresentedObject(object objectivec.IObject) int {
	rv := objc.Send[int](m.ID, objc.Sel("indexOfItemWithRepresentedObject:"), object)
	return rv
}

// Returns the index of the menu item in the menu with the given submenu.
//
// submenu: A menu object that is a menu item of the menu (that is, a submenu).
//
// # Return Value
//
// The integer index of the menu item or, if no such menu item is in the menu,
// –1.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/indexOfItem(withSubmenu:)
func (m NSMenu) IndexOfItemWithSubmenu(submenu INSMenu) int {
	rv := objc.Send[int](m.ID, objc.Sel("indexOfItemWithSubmenu:"), submenu)
	return rv
}

// Assigns a menu to be a submenu of the menu controlled by a given menu item.
//
// menu: A menu object that is to be a submenu of the menu.
//
// item: A menu item (that is, an object conforming to the NSMenuItem protocol) that
// controls `aMenu`. The method sets the action of `anItem` to
// [NSMenu.SubmenuAction].
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/setSubmenu(_:for:)
func (m NSMenu) SetSubmenuForItem(menu INSMenu, item INSMenuItem) {
	objc.Send[objc.ID](m.ID, objc.Sel("setSubmenu:forItem:"), menu, item)
}

// The action method assigned to menu items that open submenus.
//
// # Discussion
//
// You may override this method to implement different behavior. Never invoke
// this method directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/submenuAction(_:)
func (m NSMenu) SubmenuAction(sender objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("submenuAction:"), sender)
}

// Enables or disables the menu items of the menu based on the
// NSMenuValidation informal protocol and sizes the menu to fit its current
// menu items if necessary.
//
// # Discussion
//
// For more information, see [NSMenuValidation].
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/update()
func (m NSMenu) Update() {
	objc.Send[objc.ID](m.ID, objc.Sel("update"))
}

// Performs the action for the menu item that corresponds to the given key
// equivalent.
//
// event: An [NSEvent] object that represents a key-equivalent event.
//
// # Return Value
//
// Returns true if `event` is a key equivalent that the menu should handle,
// otherwise returns false.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/performKeyEquivalent(with:)
func (m NSMenu) PerformKeyEquivalent(event INSEvent) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("performKeyEquivalent:"), event)
	return rv
}

// Causes the application to send the action message of a specified menu item
// to its target.
//
// index: The integer index of a menu item.
//
// # Discussion
//
// If a target is not specified, the message is sent to the first responder.
// As a side effect, this method posts [willSendActionNotification] and
// [didSendActionNotification].
//
// In macOS 10.6 and later the “ no longer triggers menu validation. This is
// because validation is typically done during menu tracking or key equivalent
// matching, so the subsequent “ validation was redundant. To trigger
// validation explicitly, use invoke the [NSMenu.Update] method.
//
// In OS X v10.6 “, when called, now triggers highlighting in the menu bar.
// It also sends out appropriate accessibility notifications indicating the
// item was selected.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/performActionForItem(at:)
//
// [didSendActionNotification]: https://developer.apple.com/documentation/AppKit/NSMenu/didSendActionNotification
// [willSendActionNotification]: https://developer.apple.com/documentation/AppKit/NSMenu/willSendActionNotification
func (m NSMenu) PerformActionForItemAtIndex(index int) {
	objc.Send[objc.ID](m.ID, objc.Sel("performActionForItemAtIndex:"), index)
}

// Pops up the menu at the specified location.
//
// item: The menu item to be positioned at the specified location in the view.
//
// location: The location in the `view` coordinate system to display the menu item.
//
// view: The view to display the menu item over.
//
// # Return Value
//
// true if menu tracking ended because an item was selected, and false if menu
// tracking was cancelled for any reason.
//
// # Discussion
//
// Displays the menu as a pop-up menu. The top left corner of the specified
// item (if specified, `item` must be present in the menu) is positioned at
// the specified location in the specified view, interpreted in the view’s
// own coordinate system.
//
// If `item` is `nil`, the menu is positioned such that the top left of the
// menu content frame is at the given location.
//
// If `view` is `nil`, the location is interpreted in the screen coordinate
// system. This allows you to pop up a menu disconnected from any window.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/popUp(positioning:at:in:)
func (m NSMenu) PopUpMenuPositioningItemAtLocationInView(item INSMenuItem, location corefoundation.CGPoint, view INSView) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("popUpMenuPositioningItem:atLocation:inView:"), item, location, view)
	return rv
}

// Dismisses the menu and ends all menu tracking.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/cancelTracking()
func (m NSMenu) CancelTracking() {
	objc.Send[objc.ID](m.ID, objc.Sel("cancelTracking"))
}

// Dismisses the menu and ends all menu tracking without displaying the
// associated animation.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/cancelTrackingWithoutAnimation()
func (m NSMenu) CancelTrackingWithoutAnimation() {
	objc.Send[objc.ID](m.ID, objc.Sel("cancelTrackingWithoutAnimation"))
}

// The appearance of the receiver, in an [NSAppearance] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/appearance
func (m NSMenu) Appearance() INSAppearance {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("appearance"))
	return NSAppearanceFromID(rv)
}

// The appearance that will be used when the receiver is drawn onscreen, in an
// [NSAppearance] object. (read-only)
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/effectiveAppearance
func (m NSMenu) EffectiveAppearance() INSAppearance {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("effectiveAppearance"))
	return NSAppearanceFromID(rv)
}

// A string that identifies the user interface item.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemIdentification/identifier
func (m NSMenu) Identifier() NSUserInterfaceItemIdentifier {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("identifier"))
	return NSUserInterfaceItemIdentifier(foundation.NSStringFromID(rv).String())
}
func (m NSMenu) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns a Boolean value that indicates whether the menu bar is visible.
//
// # Return Value
//
// true if the menu bar is visible and selectable, otherwise false.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/menuBarVisible()
func (_NSMenuClass NSMenuClass) MenuBarVisible() bool {
	rv := objc.Send[bool](objc.ID(_NSMenuClass.class), objc.Sel("menuBarVisible"))
	return rv
}

// Sets whether the menu bar is visible and selectable by the user.
//
// visible: true if the menu bar should be visible and selectable, otherwise false.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/setMenuBarVisible(_:)
func (_NSMenuClass NSMenuClass) SetMenuBarVisible(visible bool) {
	objc.Send[objc.ID](objc.ID(_NSMenuClass.class), objc.Sel("setMenuBarVisible:"), visible)
}

// Displays a contextual menu over a view for an event.
//
// menu: The menu object to use for the contextual menu.
//
// event: An [NSEvent] object representing the event.
//
// view: The view object over which to display the contextual menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/popUpContextMenu(_:with:for:)
func (_NSMenuClass NSMenuClass) PopUpContextMenuWithEventForView(menu INSMenu, event INSEvent, view INSView) {
	objc.Send[objc.ID](objc.ID(_NSMenuClass.class), objc.Sel("popUpContextMenu:withEvent:forView:"), menu, event, view)
}

// Displays a contextual menu over a view for an event using a specified font.
//
// menu: The menu object to use for the contextual menu.
//
// event: An [NSEvent] object representing the event.
//
// view: The view object over which to display the contextual menu.
//
// font: An [NSFont] object representing the font for the contextual menu. If you
// pass in `nil` for the font, the method uses the default font for `menu`.
//
// # Discussion
//
// Specifying a font using the font parameter is discouraged. Instead, set the
// menu’s font using the [NSMenu.Font] property, then pass `nil` for the
// `font` parameter.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/popUpContextMenu(_:with:for:with:)
func (_NSMenuClass NSMenuClass) PopUpContextMenuWithEventForViewWithFont(menu INSMenu, event INSEvent, view INSView, font NSFont) {
	objc.Send[objc.ID](objc.ID(_NSMenuClass.class), objc.Sel("popUpContextMenu:withEvent:forView:withFont:"), menu, event, view, font)
}

// Creates a palette style menu displaying user-selectable color tags.
//
// colors: The display colors for the menu items.
//
// itemTitles: The menu item titles.
//
// onSelectionChange: The closure to invoke when someone selects the menu item.
//
// # Return Value
//
// A menu in the palette presentation style.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/paletteMenuWithColors:titles:selectionHandler:
func (_NSMenuClass NSMenuClass) PaletteMenuWithColorsTitlesSelectionHandler(colors []NSColor, itemTitles []string, onSelectionChange MenuHandler) NSMenu {
	_block2, _ := NewMenuBlock(onSelectionChange)
	rv := objc.Send[objc.ID](objc.ID(_NSMenuClass.class), objc.Sel("paletteMenuWithColors:titles:selectionHandler:"), colors, itemTitles, _block2)
	return NSMenuFromID(rv)
}

// Creates a palette style menu displaying user-selectable color tags that
// tint using the specified array of colors.
//
// colors: The display colors for the menu items.
//
// itemTitles: The menu item titles.
//
// image: The image the system displays for the menu items.
//
// onSelectionChange: The closure to invoke when someone selects the menu item.
//
// # Return Value
//
// A menu in the palette presentation style.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/paletteMenuWithColors:titles:templateImage:selectionHandler:
func (_NSMenuClass NSMenuClass) PaletteMenuWithColorsTitlesTemplateImageSelectionHandler(colors []NSColor, itemTitles []string, image INSImage, onSelectionChange MenuHandler) NSMenu {
	_block3, _ := NewMenuBlock(onSelectionChange)
	rv := objc.Send[objc.ID](objc.ID(_NSMenuClass.class), objc.Sel("paletteMenuWithColors:titles:templateImage:selectionHandler:"), colors, itemTitles, image, _block3)
	return NSMenuFromID(rv)
}

// The menu bar height for the main menu in pixels.
//
// # Discussion
//
// For the main menu, the value of this property is a value of type [CGFloat],
// indicating the height of the menu bar in pixels. For any other menu, the
// value of this property is `0`.
//
// This property supersedes the `menuBarHeight` class method of the
// [NSMenuView] class.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/menuBarHeight
func (m NSMenu) MenuBarHeight() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("menuBarHeight"))
	return rv
}

// The number of menu items in the menu, including separator items.
//
// # Discussion
//
// This property contains a value of type [NSInteger] that indicates the
// number of menu items in the menu, including separator items.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/numberOfItems
func (m NSMenu) NumberOfItems() int {
	rv := objc.Send[int](m.ID, objc.Sel("numberOfItems"))
	return rv
}

// An array containing the menu items in the menu.
//
// # Discussion
//
// This property contains an array of menu items in the menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/items
func (m NSMenu) ItemArray() []NSMenuItem {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("itemArray"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSMenuItem {
		return NSMenuItemFromID(id)
	})
}
func (m NSMenu) SetItemArray(value []NSMenuItem) {
	objc.Send[struct{}](m.ID, objc.Sel("setItemArray:"), objectivec.IObjectSliceToNSArray(value))
}

// The parent menu that contains the menu as a submenu.
//
// # Discussion
//
// This property contains a value of type [NSMenu] representing the the parent
// menu that contains the menu as a submenu. If the menu has no parent menu,
// then the value of this property is `nil`.
//
// You should never invoke the setter method for this property directly. The
// setter method is called automatically when changes to the parent menu
// occur. You can, however, override the setter method for this property in
// order to take action when changes to the parent menu occur.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/supermenu
func (m NSMenu) Supermenu() INSMenu {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("supermenu"))
	return NSMenuFromID(objc.ID(rv))
}
func (m NSMenu) SetSupermenu(value INSMenu) {
	objc.Send[struct{}](m.ID, objc.Sel("setSupermenu:"), value)
}

// Indicates whether the menu automatically enables and disables its menu
// items.
//
// # Discussion
//
// This property contains a Boolean value, indicating whether the menu
// automatically enables and disables its menu items. If set to true, menu
// items of the menu are automatically enabled and disabled according to rules
// computed by the NSMenuValidation informal protocol. By default, [NSMenu]
// objects autoenable their menu items.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/autoenablesItems
func (m NSMenu) AutoenablesItems() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("autoenablesItems"))
	return rv
}
func (m NSMenu) SetAutoenablesItems(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setAutoenablesItems:"), value)
}

// The font of the menu and its submenus.
//
// # Discussion
//
// This property contains a font object of the menu and its submenus that
// don’t specify fonts of their own.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/font
func (m NSMenu) Font() NSFont {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("font"))
	return NSFontFromID(objc.ID(rv))
}
func (m NSMenu) SetFont(value NSFont) {
	objc.Send[struct{}](m.ID, objc.Sel("setFont:"), value)
}

// The title of the menu.
//
// # Discussion
//
// This property contains a string value indicating the title of the menu. If
// the menu is a submenu of the application’s main menu, then the title of
// the menu appears in the menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/title
func (m NSMenu) Title() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
func (m NSMenu) SetTitle(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setTitle:"), objc.String(value))
}

// The menu items that are currently selected.
//
// # Discussion
//
// An item selects when its state is [on]. If the tracking mode is
// [NSMenuSelectionModeSelectOne] or [NSMenuSelectionModeSelectAny], the
// property only selects or returns menu items whose show-target action is
// `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/selectedItems
//
// [on]: https://developer.apple.com/documentation/AppKit/NSControl/StateValue/on
func (m NSMenu) SelectedItems() []NSMenuItem {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("selectedItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSMenuItem {
		return NSMenuItemFromID(id)
	})
}
func (m NSMenu) SetSelectedItems(value []NSMenuItem) {
	objc.Send[struct{}](m.ID, objc.Sel("setSelectedItems:"), objectivec.IObjectSliceToNSArray(value))
}

// The selection mode of the menu.
//
// # Discussion
//
// The selection mode only affects menu items that belong to the same
// selection group. A selection group consists of the items with the same
// target-action.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/selectionMode-swift.property
func (m NSMenu) SelectionMode() NSMenuSelectionMode {
	rv := objc.Send[NSMenuSelectionMode](m.ID, objc.Sel("selectionMode"))
	return NSMenuSelectionMode(rv)
}
func (m NSMenu) SetSelectionMode(value NSMenuSelectionMode) {
	objc.Send[struct{}](m.ID, objc.Sel("setSelectionMode:"), value)
}

// The minimum width of the menu in screen coordinates.
//
// # Discussion
//
// This property contains a value of type [CGFloat], indicating the minimum
// width of the menu in screen coordinates.
//
// The menu will not draw smaller than its minimum width, but may draw larger
// if it needs more space. The default value for this property is `0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/minimumWidth
func (m NSMenu) MinimumWidth() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("minimumWidth"))
	return rv
}
func (m NSMenu) SetMinimumWidth(value float64) {
	objc.Send[struct{}](m.ID, objc.Sel("setMinimumWidth:"), value)
}

// The size of the menu in screen coordinates
//
// # Discussion
//
// This property contains a value of type [NSSize], indicating the size of the
// menu in screen coordinates.
//
// The menu may draw at a smaller size when shown, depending on its
// positioning and display configuration.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/size
func (m NSMenu) Size() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m.ID, objc.Sel("size"))
	return corefoundation.CGSize(rv)
}

// The available properties for the menu.
//
// # Discussion
//
// This property contains a bitwise-C [OR] set of [NSMenu.Properties] values
// that are applicable to the menu.
//
// This property may be queried from specific callbacks to determine which
// menu properties are defined, and whether or not they are relevant to
// changes you need to make to the menu. This property is intended to allow
// for more efficient updating of the menu in certain circumstances.
//
// For example, if the [NSMenuPropertyItemImage] property isn’t set, your
// delegate doesn’t need to spend time updating the images of the menu
// items, because the images aren’t needed (for example, during
// key-equivalent matching).
//
// You have to update a menu property only if it has changed since you last
// set it, even if the corresponding bit is `1`. For example, if the title of
// a menu item never changes, you have to set it only once.
//
// Accessing this property is optional; it is always acceptable to fully
// update all properties of the menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/propertiesToUpdate
//
// [NSMenu.Properties]: https://developer.apple.com/documentation/AppKit/NSMenu/Properties
func (m NSMenu) PropertiesToUpdate() NSMenuProperties {
	rv := objc.Send[NSMenuProperties](m.ID, objc.Sel("propertiesToUpdate"))
	return NSMenuProperties(rv)
}

// The presentation style of the menu.
//
// # Discussion
//
// This property isn’t respected if the menu is the main menu of the app.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/presentationStyle-swift.property
func (m NSMenu) PresentationStyle() NSMenuPresentationStyle {
	rv := objc.Send[NSMenuPresentationStyle](m.ID, objc.Sel("presentationStyle"))
	return NSMenuPresentationStyle(rv)
}
func (m NSMenu) SetPresentationStyle(value NSMenuPresentationStyle) {
	objc.Send[struct{}](m.ID, objc.Sel("setPresentationStyle:"), value)
}

// Indicates whether the pop-up menu allows appending of contextual menu
// plug-in items.
//
// # Discussion
//
// This property contains a Boolean value indicating whether the pop-up menu
// allows appending of contextual menu plug-in items.
//
// Contextual menu plug-ins are system-wide services provided by other
// applications. For example, a contextual menu plug-in might provide an
// “Open URL…” service. If you enable context menu plug-ins, your
// application’s contextual menu will display the appropriate items for the
// currently selected data type.
//
// The default value for this property is true.
//
// See [Services Implementation Guide] for more information on contextual menu
// plug-ins.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/allowsContextMenuPlugIns
//
// [Services Implementation Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/SysServices/introduction.html#//apple_ref/doc/uid/10000101i
func (m NSMenu) AllowsContextMenuPlugIns() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("allowsContextMenuPlugIns"))
	return rv
}
func (m NSMenu) SetAllowsContextMenuPlugIns(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setAllowsContextMenuPlugIns:"), value)
}

// Indicates whether the menu displays the state column.
//
// # Discussion
//
// This property contains a Boolean value indicating whether the menu displays
// the state column. The default value for this property is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/showsStateColumn
func (m NSMenu) ShowsStateColumn() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("showsStateColumn"))
	return rv
}
func (m NSMenu) SetShowsStateColumn(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setShowsStateColumn:"), value)
}

// Indicates the currently highlighted item in the menu.
//
// # Discussion
//
// This property indicates the currently highlighted item in the menu. If no
// menu is highlighted, this property has a value of `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/highlightedItem
func (m NSMenu) HighlightedItem() INSMenuItem {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("highlightedItem"))
	return NSMenuItemFromID(objc.ID(rv))
}

// Configures the layout direction of menu items in the menu.
//
// # Discussion
//
// This property configures the layout direction (a value of type
// [NSUserInterfaceLayoutDirection]) of menu items in the menu. If no layout
// direction is explicitly set for a menu, then the menu defaults to the
// layout direction specified for the application object. See
// [NSApplication.UserInterfaceLayoutDirection] in [NSApplication].
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/userInterfaceLayoutDirection
//
// [NSUserInterfaceLayoutDirection]: https://developer.apple.com/documentation/AppKit/NSUserInterfaceLayoutDirection
func (m NSMenu) UserInterfaceLayoutDirection() NSUserInterfaceLayoutDirection {
	rv := objc.Send[NSUserInterfaceLayoutDirection](m.ID, objc.Sel("userInterfaceLayoutDirection"))
	return NSUserInterfaceLayoutDirection(rv)
}
func (m NSMenu) SetUserInterfaceLayoutDirection(value NSUserInterfaceLayoutDirection) {
	objc.Send[struct{}](m.ID, objc.Sel("setUserInterfaceLayoutDirection:"), value)
}

// The delegate of the menu.
//
// # Discussion
//
// This property indicates the delegate of the menu.
//
// You can use the delegate to populate a menu just before it is drawn and to
// check for key equivalents without creating a menu item.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenu/delegate
func (m NSMenu) Delegate() NSMenuDelegate {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("delegate"))
	return NSMenuDelegateObjectFromID(rv)
}
func (m NSMenu) SetDelegate(value NSMenuDelegate) {
	objc.Send[struct{}](m.ID, objc.Sel("setDelegate:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSMenu/automaticallyInsertsWritingToolsItems
func (m NSMenu) AutomaticallyInsertsWritingToolsItems() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("automaticallyInsertsWritingToolsItems"))
	return rv
}
func (m NSMenu) SetAutomaticallyInsertsWritingToolsItems(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setAutomaticallyInsertsWritingToolsItems:"), value)
}

// Protocol methods for NSAppearanceCustomization

// The appearance of the receiver, in an [NSAppearance] object.
//
// # Discussion
//
// The default value for this property is `nil`, which means that the receiver
// uses the appearance it inherits from the nearest ancestor that has set an
// appearance. When you set `appearance` to a non-`nil` value, the receiver
// and the views it contains use the specified appearance.
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/appearance
func (o NSMenu) SetAppearance(value INSAppearance) {
	objc.Send[struct{}](o.ID, objc.Sel("setAppearance:"), value)
}

// Protocol methods for NSUserInterfaceItemIdentification

// A string that identifies the user interface item.
//
// # Discussion
//
// Identifiers are used during window restoration operations to uniquely
// identify the windows of the application. You can set the value of this
// string programmatically or in Interface Builder. If you create an item in
// Interface Builder and do not set a value for this string, a unique value is
// created for the item when the nib file is loaded. For programmatically
// created views, you typically set this value after creating the item but
// before adding it to a window.
//
// You should not change the value of a window’s identifier after adding any
// views to the window. For views and controls in a window, the value you
// specify for this string must be unique on a per-window basis.
//
// The slash (`/`), backslash (`\`), or colon (`:`) characters are reserved
// and must not be used in your custom identifiers. Similarly, Apple reserves
// all identifiers beginning with an underscore (`_`) character. Applications
// and frameworks should use a consistent prefix for their identifiers to
// avoid collisions with other frameworks. For a list of prefixes used by the
// system frameworks, see [OS X Frameworks] in [Mac Technology Overview].
//
// If you are subclassing a class from one of the system frameworks, do not
// override the accessor methods of this protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemIdentification/identifier
//
// [Mac Technology Overview]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/OSX_Technology_Overview/About/About.html#//apple_ref/doc/uid/TP40001067
// [OS X Frameworks]: https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/OSX_Technology_Overview/SystemFrameworks/SystemFrameworks.html#//apple_ref/doc/uid/TP40001067-CH210
func (o NSMenu) SetIdentifier(value NSUserInterfaceItemIdentifier) {
	objc.Send[struct{}](o.ID, objc.Sel("setIdentifier:"), objc.String(string(value)))
}
