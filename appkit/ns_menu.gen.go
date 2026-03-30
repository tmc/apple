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
	NSAppearanceCustomization
	NSUserInterfaceItemIdentification

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
// This method invokes [InsertItemAtIndex]. Thus, the menu does not accept the
// menu item if it already belongs to another menu. After adding the menu
// item, the menu updates itself.
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
// [SubmenuAction].
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
// validation explicitly, use invoke the [Update] method.
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
// menu’s font using the [Font] property, then pass `nil` for the `font`
// parameter.
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
// [UserInterfaceLayoutDirection] in [NSApplication].
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
func (m NSMenu) Appearance() INSAppearance {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("appearance"))
	return NSAppearanceFromID(objc.ID(rv))
}
func (m NSMenu) SetAppearance(value INSAppearance) {
	objc.Send[struct{}](m.ID, objc.Sel("setAppearance:"), value)
}

// The appearance that will be used when the receiver is drawn onscreen, in an
// [NSAppearance] object. (read-only)
//
// # Discussion
//
// The default value for this property is provided by the nearest ancestor of
// the receiver that has set an appearance.
//
// You can use this property to ensure that an offscreen view sets the
// appropriate current appearance when it draws onscreen.
//
// See: https://developer.apple.com/documentation/AppKit/NSAppearanceCustomization/effectiveAppearance
func (m NSMenu) EffectiveAppearance() INSAppearance {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("effectiveAppearance"))
	return NSAppearanceFromID(objc.ID(rv))
}

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
func (m NSMenu) Identifier() NSUserInterfaceItemIdentifier {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("identifier"))
	return NSUserInterfaceItemIdentifier(foundation.NSStringFromID(rv).String())
}
func (m NSMenu) SetIdentifier(value NSUserInterfaceItemIdentifier) {
	objc.Send[struct{}](m.ID, objc.Sel("setIdentifier:"), objc.String(string(value)))
}

// Protocol methods for NSAccessibilityElementProtocol

// Returns the accessibility element’s frame in screen coordinates.
//
// # Return Value
//
// The element’s frame in screen coordinates.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFrame] property. This method is called whenever accessibility
// clients request the [size] or [position] attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityFrame()
//
// [accessibilityFrame]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
// [position]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
// [size]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
func (o NSMenu) AccessibilityFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrame"))
	return rv
}

// Returns the accessibility element’s parent in the accessibility
// hierarchy.
//
// # Return Value
//
// The element’s parent in the accessibility hierarchy.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityParent] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityParent()
//
// [accessibilityParent]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityParent
func (o NSMenu) AccessibilityParent() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityParent"))
	return objectivec.Object{ID: rv}
}

// Returns the accessibility element’s identity.
//
// # Return Value
//
// Returns the unique ID for the accessibility element. It is often used in
// automated testing.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityIdentifier] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityIdentifier()
//
// [accessibilityIdentifier]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIdentifier
func (o NSMenu) AccessibilityIdentifier() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// Returns a Boolean value that indicates whether the accessibility element
// has the keyboard focus.
//
// # Return Value
//
// true if this element has the keyboard focus; otherwise, false.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFocused] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/isAccessibilityFocused()
//
// [accessibilityFocused]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocused
func (o NSMenu) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
}

// Protocol methods for NSAccessibilityProtocol

// Returns a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityElement()
func (o NSMenu) IsAccessibilityElement() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityElement"))
	return rv
}

// Sets a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityElement(_:)
func (o NSMenu) SetAccessibilityElement(accessibilityElement bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityElement:"), accessibilityElement)
}

// Returns a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEnabled()
func (o NSMenu) IsAccessibilityEnabled() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEnabled"))
	return rv
}

// Sets a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityEnabled(_:)
func (o NSMenu) SetAccessibilityEnabled(accessibilityEnabled bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEnabled:"), accessibilityEnabled)
}

// Sets the accessibility element’s frame in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFrame(_:)
func (o NSMenu) SetAccessibilityFrame(accessibilityFrame corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrame:"), accessibilityFrame)
}

// Returns the help text for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHelp()
func (o NSMenu) AccessibilityHelp() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHelp"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the help text for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHelp(_:)
func (o NSMenu) SetAccessibilityHelp(accessibilityHelp string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHelp:"), objc.String(accessibilityHelp))
}

// Returns a short description of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLabel()
func (o NSMenu) AccessibilityLabel() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
}

// Sets a short description of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLabel(_:)
func (o NSMenu) SetAccessibilityLabel(accessibilityLabel string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabel:"), objc.String(accessibilityLabel))
}

// Returns the title of the accessibility element—for example, a button’s
// visible text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTitle()
func (o NSMenu) AccessibilityTitle() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitle"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the title of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTitle(_:)
func (o NSMenu) SetAccessibilityTitle(accessibilityTitle string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitle:"), objc.String(accessibilityTitle))
}

// Returns the accessibility element’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityValue()
func (o NSMenu) AccessibilityValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValue"))
	return objectivec.Object{ID: rv}
}

// Sets the accessibility element’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityValue(_:)
func (o NSMenu) SetAccessibilityValue(accessibilityValue objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValue:"), accessibilityValue)
}

// Returns a Boolean value that indicates whether assistive apps can invoke
// the specified selector on the accessibility element.
//
// selector: The selector to check.
//
// # Return Value
//
// true, if accessibility clients can call the selector; otherwise, false.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelectorAllowed(_:)
func (o NSMenu) IsAccessibilitySelectorAllowed(selector objc.SEL) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelectorAllowed:"), selector)
	return rv
}

// Returns the contents of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityContents()
func (o NSMenu) AccessibilityContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityContents"))
	return foundation.NSArrayFromID(rv)
}

// Sets the contents of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityContents(_:)
func (o NSMenu) SetAccessibilityContents(accessibilityContents objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityContents:"), accessibilityContents)
}

// Returns the critical value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCriticalValue()
func (o NSMenu) AccessibilityCriticalValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCriticalValue"))
	return objectivec.Object{ID: rv}
}

// Sets the critical value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCriticalValue(_:)
func (o NSMenu) SetAccessibilityCriticalValue(accessibilityCriticalValue objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCriticalValue:"), accessibilityCriticalValue)
}

// Sets the accessibility element’s identity.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityIdentifier(_:)
func (o NSMenu) SetAccessibilityIdentifier(accessibilityIdentifier string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIdentifier:"), objc.String(accessibilityIdentifier))
}

// Returns the maximum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMaxValue()
func (o NSMenu) AccessibilityMaxValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMaxValue"))
	return objectivec.Object{ID: rv}
}

// Sets the maximum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMaxValue(_:)
func (o NSMenu) SetAccessibilityMaxValue(accessibilityMaxValue objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMaxValue:"), accessibilityMaxValue)
}

// Returns the minimum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMinValue()
func (o NSMenu) AccessibilityMinValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinValue"))
	return objectivec.Object{ID: rv}
}

// Sets the minimum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMinValue(_:)
func (o NSMenu) SetAccessibilityMinValue(accessibilityMinValue objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinValue:"), accessibilityMinValue)
}

// Returns the orientation of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityOrientation()
func (o NSMenu) AccessibilityOrientation() NSAccessibilityOrientation {
	rv := objc.Send[NSAccessibilityOrientation](o.ID, objc.Sel("accessibilityOrientation"))
	return rv
}

// Sets the orientation of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityOrientation(_:)
func (o NSMenu) SetAccessibilityOrientation(accessibilityOrientation NSAccessibilityOrientation) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrientation:"), accessibilityOrientation)
}

// Returns a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityProtectedContent()
func (o NSMenu) IsAccessibilityProtectedContent() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityProtectedContent"))
	return rv
}

// Sets a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityProtectedContent(_:)
func (o NSMenu) SetAccessibilityProtectedContent(accessibilityProtectedContent bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProtectedContent:"), accessibilityProtectedContent)
}

// Returns a Boolean value that determines whether the accessibility element
// is currently in a selected state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelected()
func (o NSMenu) IsAccessibilitySelected() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelected"))
	return rv
}

// Sets a Boolean value that determines whether the accessibility element is
// currently in a selected state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelected(_:)
func (o NSMenu) SetAccessibilitySelected(accessibilitySelected bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelected:"), accessibilitySelected)
}

// Returns the URL for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityURL()
func (o NSMenu) AccessibilityURL() foundation.INSURL {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityURL"))
	return foundation.NSURLFromID(rv)
}

// Sets the URL for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityURL(_:)
func (o NSMenu) SetAccessibilityURL(accessibilityURL foundation.INSURL) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityURL:"), accessibilityURL)
}

// Returns the human-readable description of the accessibility element’s
// value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityValueDescription()
func (o NSMenu) AccessibilityValueDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValueDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the human-readable description of the accessibility element’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityValueDescription(_:)
func (o NSMenu) SetAccessibilityValueDescription(accessibilityValueDescription string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValueDescription:"), objc.String(accessibilityValueDescription))
}

// Returns the warning value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityWarningValue()
func (o NSMenu) AccessibilityWarningValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWarningValue"))
	return objectivec.Object{ID: rv}
}

// Sets the warning value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityWarningValue(_:)
func (o NSMenu) SetAccessibilityWarningValue(accessibilityWarningValue objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWarningValue:"), accessibilityWarningValue)
}

// Returns the child accessibility elements in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityChildren()
func (o NSMenu) AccessibilityChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityChildren"))
	return foundation.NSArrayFromID(rv)
}

// Sets the child accessibility elements in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityChildren(_:)
func (o NSMenu) SetAccessibilityChildren(accessibilityChildren objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildren:"), accessibilityChildren)
}

// Returns the array of child accessibility elements in order for linear
// navigation.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityChildrenInNavigationOrder()
func (o NSMenu) AccessibilityChildrenInNavigationOrder() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityChildrenInNavigationOrder"))
	return foundation.NSArrayFromID(rv)
}

// Sets the array of child accessibility elements in order for linear
// navigation.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityChildrenInNavigationOrder(_:)
func (o NSMenu) SetAccessibilityChildrenInNavigationOrder(accessibilityChildrenInNavigationOrder objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildrenInNavigationOrder:"), accessibilityChildrenInNavigationOrder)
}

// Sets the accessibility element’s parent in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityParent(_:)
func (o NSMenu) SetAccessibilityParent(accessibilityParent objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityParent:"), accessibilityParent)
}

// Returns the accessibility element’s currently selected children.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedChildren()
func (o NSMenu) AccessibilitySelectedChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedChildren"))
	return foundation.NSArrayFromID(rv)
}

// Sets the accessibility element’s currently selected children.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedChildren(_:)
func (o NSMenu) SetAccessibilitySelectedChildren(accessibilitySelectedChildren objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedChildren:"), accessibilitySelectedChildren)
}

// Returns the top-level element that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTopLevelUIElement()
func (o NSMenu) AccessibilityTopLevelUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTopLevelUIElement"))
	return objectivec.Object{ID: rv}
}

// Sets the top-level element that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTopLevelUIElement(_:)
func (o NSMenu) SetAccessibilityTopLevelUIElement(accessibilityTopLevelUIElement objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTopLevelUIElement:"), accessibilityTopLevelUIElement)
}

// Returns the accessibility element’s visible child accessibility elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleChildren()
func (o NSMenu) AccessibilityVisibleChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleChildren"))
	return foundation.NSArrayFromID(rv)
}

// Sets the accessibility element’s visible child accessibility elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleChildren(_:)
func (o NSMenu) SetAccessibilityVisibleChildren(accessibilityVisibleChildren objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleChildren:"), accessibilityVisibleChildren)
}

// Returns the child accessibility element with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityApplicationFocusedUIElement()
func (o NSMenu) AccessibilityApplicationFocusedUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityApplicationFocusedUIElement"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityApplicationFocusedUIElement(_:)
func (o NSMenu) SetAccessibilityApplicationFocusedUIElement(accessibilityApplicationFocusedUIElement objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityApplicationFocusedUIElement:"), accessibilityApplicationFocusedUIElement)
}

// Sets a Boolean value that determines whether the accessibility element has
// the keyboard focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFocused(_:)
func (o NSMenu) SetAccessibilityFocused(accessibilityFocused bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocused:"), accessibilityFocused)
}

// Returns the child window with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFocusedWindow()
func (o NSMenu) AccessibilityFocusedWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFocusedWindow"))
	return objectivec.Object{ID: rv}
}

// Sets the child window with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFocusedWindow(_:)
func (o NSMenu) SetAccessibilityFocusedWindow(accessibilityFocusedWindow objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocusedWindow:"), accessibilityFocusedWindow)
}

// Returns the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySharedFocusElements()
func (o NSMenu) AccessibilitySharedFocusElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedFocusElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySharedFocusElements(_:)
func (o NSMenu) SetAccessibilitySharedFocusElements(accessibilitySharedFocusElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedFocusElements:"), accessibilitySharedFocusElements)
}

// Returns a Boolean value that determines whether the accessibility element
// must have content for successful submission of a form.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityRequired()
func (o NSMenu) IsAccessibilityRequired() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityRequired"))
	return rv
}

// Sets a Boolean value that determines whether the accessibility element must
// have content for successful submission of a form.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRequired(_:)
func (o NSMenu) SetAccessibilityRequired(accessibilityRequired bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRequired:"), accessibilityRequired)
}

// Returns the type of interface element that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRole()
func (o NSMenu) AccessibilityRole() NSAccessibilityRole {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRole"))
	return NSAccessibilityRole(foundation.NSStringFromID(rv).String())
}

// Sets the type of interface element that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRole(_:)
func (o NSMenu) SetAccessibilityRole(accessibilityRole NSAccessibilityRole) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRole:"), objc.String(string(accessibilityRole)))
}

// Returns a localized, human-intelligible description of the accessibility
// element’s role, such as .
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRoleDescription()
func (o NSMenu) AccessibilityRoleDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRoleDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the localized, human-intelligible description of the accessibility
// element’s role, such as .
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRoleDescription(_:)
func (o NSMenu) SetAccessibilityRoleDescription(accessibilityRoleDescription string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRoleDescription:"), objc.String(accessibilityRoleDescription))
}

// Returns the specialized interface element type that the accessibility
// element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySubrole()
func (o NSMenu) AccessibilitySubrole() NSAccessibilitySubrole {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySubrole"))
	return NSAccessibilitySubrole(foundation.NSStringFromID(rv).String())
}

// Sets the specialized interface element type that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySubrole(_:)
func (o NSMenu) SetAccessibilitySubrole(accessibilitySubrole NSAccessibilitySubrole) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySubrole:"), objc.String(string(accessibilitySubrole)))
}

// Returns the custom actions of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCustomActions()
func (o NSMenu) AccessibilityCustomActions() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCustomActions"))
	return foundation.NSArrayFromID(rv)
}

// Sets the custom actions of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCustomActions(_:)
func (o NSMenu) SetAccessibilityCustomActions(accessibilityCustomActions objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomActions:"), accessibilityCustomActions)
}

// Returns the custom rotors of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCustomRotors()
func (o NSMenu) AccessibilityCustomRotors() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCustomRotors"))
	return foundation.NSArrayFromID(rv)
}

// Sets the custom rotors of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCustomRotors(_:)
func (o NSMenu) SetAccessibilityCustomRotors(accessibilityCustomRotors objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomRotors:"), accessibilityCustomRotors)
}

// Returns the line number that contains the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityInsertionPointLineNumber()
func (o NSMenu) AccessibilityInsertionPointLineNumber() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityInsertionPointLineNumber"))
	return rv
}

// Sets the line number that contains the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityInsertionPointLineNumber(_:)
func (o NSMenu) SetAccessibilityInsertionPointLineNumber(accessibilityInsertionPointLineNumber int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityInsertionPointLineNumber:"), accessibilityInsertionPointLineNumber)
}

// Returns the number of characters in the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityNumberOfCharacters()
func (o NSMenu) AccessibilityNumberOfCharacters() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityNumberOfCharacters"))
	return rv
}

// Sets the number of characters in the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityNumberOfCharacters(_:)
func (o NSMenu) SetAccessibilityNumberOfCharacters(accessibilityNumberOfCharacters int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNumberOfCharacters:"), accessibilityNumberOfCharacters)
}

// Returns the placeholder value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPlaceholderValue()
func (o NSMenu) AccessibilityPlaceholderValue() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPlaceholderValue"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the placeholder value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityPlaceholderValue(_:)
func (o NSMenu) SetAccessibilityPlaceholderValue(accessibilityPlaceholderValue string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPlaceholderValue:"), objc.String(accessibilityPlaceholderValue))
}

// Returns the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedText()
func (o NSMenu) AccessibilitySelectedText() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedText"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedText(_:)
func (o NSMenu) SetAccessibilitySelectedText(accessibilitySelectedText string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedText:"), objc.String(accessibilitySelectedText))
}

// Returns the range of the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedTextRange()
func (o NSMenu) AccessibilitySelectedTextRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySelectedTextRange"))
	return rv
}

// Sets the range of the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedTextRange(_:)
func (o NSMenu) SetAccessibilitySelectedTextRange(accessibilitySelectedTextRange foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRange:"), accessibilitySelectedTextRange)
}

// Returns an array of ranges for the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedTextRanges()
func (o NSMenu) AccessibilitySelectedTextRanges() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedTextRanges"))
	return foundation.NSArrayFromID(rv)
}

// Sets an array of ranges for the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedTextRanges(_:)
func (o NSMenu) SetAccessibilitySelectedTextRanges(accessibilitySelectedTextRanges objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRanges:"), accessibilitySelectedTextRanges)
}

// Returns the range of characters that the accessibility element displays.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySharedCharacterRange()
func (o NSMenu) AccessibilitySharedCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySharedCharacterRange"))
	return rv
}

// Sets the range of characters that the accessibility element displays.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySharedCharacterRange(_:)
func (o NSMenu) SetAccessibilitySharedCharacterRange(accessibilitySharedCharacterRange foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedCharacterRange:"), accessibilitySharedCharacterRange)
}

// Returns the other elements that share text with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySharedTextUIElements()
func (o NSMenu) AccessibilitySharedTextUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedTextUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the other elements that share text with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySharedTextUIElements(_:)
func (o NSMenu) SetAccessibilitySharedTextUIElements(accessibilitySharedTextUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedTextUIElements:"), accessibilitySharedTextUIElements)
}

// Returns the range of visible characters in the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleCharacterRange()
func (o NSMenu) AccessibilityVisibleCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityVisibleCharacterRange"))
	return rv
}

// Sets the range of visible characters in the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleCharacterRange(_:)
func (o NSMenu) SetAccessibilityVisibleCharacterRange(accessibilityVisibleCharacterRange foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleCharacterRange:"), accessibilityVisibleCharacterRange)
}

// Returns the substring for the specified range.
//
// range: A range of characters contained by the element.
//
// # Return Value
//
// The substring specified by the given range.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityString(for:)
func (o NSMenu) AccessibilityStringForRange(range_ foundation.NSRange) string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityStringForRange:"), range_)
	return foundation.NSStringFromID(rv).String()
}

// Returns the attributed substring for the specified range of characters.
//
// range: The range of characters.
//
// # Return Value
//
// An attributed string representing the specified characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityAttributedString(for:)
func (o NSMenu) AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAttributedStringForRange:"), range_)
	return foundation.NSAttributedStringFromID(rv)
}

// Returns the rich text format (RTF) data that describes the specified range
// of characters.
//
// range: The range of characters.
//
// # Return Value
//
// A data object containing an RTF representation of the specified characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRTF(for:)
func (o NSMenu) AccessibilityRTFForRange(range_ foundation.NSRange) foundation.INSData {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRTFForRange:"), range_)
	return foundation.NSDataFromID(rv)
}

// Returns the rectangle that encloses the specified range of characters.
//
// range: The range of characters.
//
// # Return Value
//
// The rectangle that encloses the specified characters.
//
// # Discussion
//
// If the range crosses a line boundary, the returned rectangle fully encloses
// all the lines of characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFrame(for:)
func (o NSMenu) AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrameForRange:"), range_)
	return rv
}

// Returns the line number for the line that contains the specified character
// index.
//
// index: The index for a character.
//
// # Return Value
//
// The line number for the line holding the specified character index.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLine(for:)
func (o NSMenu) AccessibilityLineForIndex(index int) int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityLineForIndex:"), index)
	return rv
}

// Returns the range of characters for the glyph that includes the specified
// character.
//
// index: The specified character.
//
// # Return Value
//
// The range of characters for the glyph.
//
// # Discussion
//
// This value always includes the specified character but may include
// additional characters if that character is part of a multicharacter glyph.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(for:)-6kv3
func (o NSMenu) AccessibilityRangeForIndex(index int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForIndex:"), index)
	return rv
}

// Returns a range of characters that all have the same style as the specified
// character.
//
// index: The index of the specified character.
//
// # Return Value
//
// A range of characters with the same style as the specified character.
//
// # Discussion
//
// This method returns a range of characters that meet two conditions: The
// range must include the specified character, and all the other characters in
// the range must match the specified character’s style. If none of the
// adjacent characters match the specified character’s style, the method
// returns only the specified character.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityStyleRange(for:)
func (o NSMenu) AccessibilityStyleRangeForIndex(index int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityStyleRangeForIndex:"), index)
	return rv
}

// Returns the range of characters in the specified line.
//
// line: The line number to be examined.
//
// # Return Value
//
// The range of characters for the specified line number. If the line ends
// with a newline character, including the newline is preferred.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(forLine:)
func (o NSMenu) AccessibilityRangeForLine(line int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForLine:"), line)
	return rv
}

// Returns the range of characters for the glyph at the specified point.
//
// point: A point in screen coordinates.
//
// # Return Value
//
// The range of characters that make up the glyph at the given point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(for:)-1iudm
func (o NSMenu) AccessibilityRangeForPosition(point corefoundation.CGPoint) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForPosition:"), point)
	return rv
}

// Returns the activation point for the user interface element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityActivationPoint()
func (o NSMenu) AccessibilityActivationPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityActivationPoint"))
	return rv
}

// Sets the activation point for the user interface element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityActivationPoint(_:)
func (o NSMenu) SetAccessibilityActivationPoint(accessibilityActivationPoint corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityActivationPoint:"), accessibilityActivationPoint)
}

// Returns the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityAlternateUIVisible()
func (o NSMenu) IsAccessibilityAlternateUIVisible() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
	return rv
}

// Sets the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityAlternateUIVisible(_:)
func (o NSMenu) SetAccessibilityAlternateUIVisible(accessibilityAlternateUIVisible bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAlternateUIVisible:"), accessibilityAlternateUIVisible)
}

// Returns the child accessibility element that represents the window’s
// cancel button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCancelButton()
func (o NSMenu) AccessibilityCancelButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCancelButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s cancel
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCancelButton(_:)
func (o NSMenu) SetAccessibilityCancelButton(accessibilityCancelButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCancelButton:"), accessibilityCancelButton)
}

// Returns the child accessibility element that represents the window’s
// close button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCloseButton()
func (o NSMenu) AccessibilityCloseButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCloseButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s close
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCloseButton(_:)
func (o NSMenu) SetAccessibilityCloseButton(accessibilityCloseButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCloseButton:"), accessibilityCloseButton)
}

// Returns the child accessibility element that represents the window’s
// default button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDefaultButton()
func (o NSMenu) AccessibilityDefaultButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDefaultButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s default
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDefaultButton(_:)
func (o NSMenu) SetAccessibilityDefaultButton(accessibilityDefaultButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDefaultButton:"), accessibilityDefaultButton)
}

// Returns the child accessibility element that represents the window’s
// full-screen button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFullScreenButton()
func (o NSMenu) AccessibilityFullScreenButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFullScreenButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s
// full-screen button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFullScreenButton(_:)
func (o NSMenu) SetAccessibilityFullScreenButton(accessibilityFullScreenButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFullScreenButton:"), accessibilityFullScreenButton)
}

// Returns the child accessibility element that represents the window’s grow
// area.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityGrowArea()
func (o NSMenu) AccessibilityGrowArea() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityGrowArea"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s grow
// area.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityGrowArea(_:)
func (o NSMenu) SetAccessibilityGrowArea(accessibilityGrowArea objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityGrowArea:"), accessibilityGrowArea)
}

// Returns a Boolean value that determines whether the window is the app’s
// main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMain()
func (o NSMenu) IsAccessibilityMain() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMain"))
	return rv
}

// Sets a Boolean value that determines whether the window is the app’s main
// window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMain(_:)
func (o NSMenu) SetAccessibilityMain(accessibilityMain bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMain:"), accessibilityMain)
}

// Returns the child accessibility element that represents the window’s
// minimize button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMinimizeButton()
func (o NSMenu) AccessibilityMinimizeButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinimizeButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s
// minimize button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMinimizeButton(_:)
func (o NSMenu) SetAccessibilityMinimizeButton(accessibilityMinimizeButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimizeButton:"), accessibilityMinimizeButton)
}

// Returns the Boolean value that determines whether the window is in a
// minimized state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMinimized()
func (o NSMenu) IsAccessibilityMinimized() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMinimized"))
	return rv
}

// Sets the Boolean value that determines whether the window is in a minimized
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMinimized(_:)
func (o NSMenu) SetAccessibilityMinimized(accessibilityMinimized bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimized:"), accessibilityMinimized)
}

// Returns a Boolean value that determines whether the window is modal.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityModal()
func (o NSMenu) IsAccessibilityModal() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityModal"))
	return rv
}

// Sets a Boolean value that determines whether the window is modal.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityModal(_:)
func (o NSMenu) SetAccessibilityModal(accessibilityModal bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityModal:"), accessibilityModal)
}

// Returns the child accessibility element that represents the window’s
// proxy icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityProxy()
func (o NSMenu) AccessibilityProxy() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityProxy"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s proxy
// icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityProxy(_:)
func (o NSMenu) SetAccessibilityProxy(accessibilityProxy objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProxy:"), accessibilityProxy)
}

// Returns the menu currently displaying for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityShownMenu()
func (o NSMenu) AccessibilityShownMenu() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityShownMenu"))
	return objectivec.Object{ID: rv}
}

// Sets the menu currently displaying for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityShownMenu(_:)
func (o NSMenu) SetAccessibilityShownMenu(accessibilityShownMenu objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityShownMenu:"), accessibilityShownMenu)
}

// Returns the child accessibility element that represents the window’s
// toolbar button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityToolbarButton()
func (o NSMenu) AccessibilityToolbarButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityToolbarButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s toolbar
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityToolbarButton(_:)
func (o NSMenu) SetAccessibilityToolbarButton(accessibilityToolbarButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityToolbarButton:"), accessibilityToolbarButton)
}

// Returns the window that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityWindow()
func (o NSMenu) AccessibilityWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindow"))
	return objectivec.Object{ID: rv}
}

// Sets the window that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityWindow(_:)
func (o NSMenu) SetAccessibilityWindow(accessibilityWindow objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindow:"), accessibilityWindow)
}

// Returns the child accessibility element that represents the window’s zoom
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityZoomButton()
func (o NSMenu) AccessibilityZoomButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityZoomButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s zoom
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityZoomButton(_:)
func (o NSMenu) SetAccessibilityZoomButton(accessibilityZoomButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityZoomButton:"), accessibilityZoomButton)
}

// Returns the icon for the app’s menu bar extra.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityExtrasMenuBar()
func (o NSMenu) AccessibilityExtrasMenuBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityExtrasMenuBar"))
	return objectivec.Object{ID: rv}
}

// Sets the icon for the app’s menu bar extra.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityExtrasMenuBar(_:)
func (o NSMenu) SetAccessibilityExtrasMenuBar(accessibilityExtrasMenuBar objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExtrasMenuBar:"), accessibilityExtrasMenuBar)
}

// Returns a Boolean value that determines whether the app is the frontmost
// app.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityFrontmost()
func (o NSMenu) IsAccessibilityFrontmost() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFrontmost"))
	return rv
}

// Sets a Boolean value that determines whether the app is the frontmost app.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFrontmost(_:)
func (o NSMenu) SetAccessibilityFrontmost(accessibilityFrontmost bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrontmost:"), accessibilityFrontmost)
}

// Returns a Boolean value that determines whether the app is in a hidden
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityHidden()
func (o NSMenu) IsAccessibilityHidden() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityHidden"))
	return rv
}

// Sets a Boolean value that determines whether the app is in a hidden state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHidden(_:)
func (o NSMenu) SetAccessibilityHidden(accessibilityHidden bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHidden:"), accessibilityHidden)
}

// Returns the app’s main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMainWindow()
func (o NSMenu) AccessibilityMainWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMainWindow"))
	return objectivec.Object{ID: rv}
}

// Sets the app’s main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMainWindow(_:)
func (o NSMenu) SetAccessibilityMainWindow(accessibilityMainWindow objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMainWindow:"), accessibilityMainWindow)
}

// Returns the app’s menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMenuBar()
func (o NSMenu) AccessibilityMenuBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMenuBar"))
	return objectivec.Object{ID: rv}
}

// Sets the app’s menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMenuBar(_:)
func (o NSMenu) SetAccessibilityMenuBar(accessibilityMenuBar objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMenuBar:"), accessibilityMenuBar)
}

// Returns an array that contains all the app’s windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityWindows()
func (o NSMenu) AccessibilityWindows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindows"))
	return foundation.NSArrayFromID(rv)
}

// Sets the array that contains all the app’s windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityWindows(_:)
func (o NSMenu) SetAccessibilityWindows(accessibilityWindows objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindows:"), accessibilityWindows)
}

// Returns the number of columns in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnCount()
func (o NSMenu) AccessibilityColumnCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityColumnCount"))
	return rv
}

// Sets the number of columns in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnCount(_:)
func (o NSMenu) SetAccessibilityColumnCount(accessibilityColumnCount int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnCount:"), accessibilityColumnCount)
}

// Returns a Boolean value that determines whether the accessibility
// element’s grid is in row major order or in column major order.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityOrderedByRow()
func (o NSMenu) IsAccessibilityOrderedByRow() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityOrderedByRow"))
	return rv
}

// Sets a Boolean value that determines whether the element’s grid is in row
// major order or in column major order.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityOrderedByRow(_:)
func (o NSMenu) SetAccessibilityOrderedByRow(accessibilityOrderedByRow bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrderedByRow:"), accessibilityOrderedByRow)
}

// Returns the number of rows in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRowCount()
func (o NSMenu) AccessibilityRowCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityRowCount"))
	return rv
}

// Sets the number of rows in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRowCount(_:)
func (o NSMenu) SetAccessibilityRowCount(accessibilityRowCount int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowCount:"), accessibilityRowCount)
}

// Returns the horizontal scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHorizontalScrollBar()
func (o NSMenu) AccessibilityHorizontalScrollBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalScrollBar"))
	return objectivec.Object{ID: rv}
}

// Sets the horizontal scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHorizontalScrollBar(_:)
func (o NSMenu) SetAccessibilityHorizontalScrollBar(accessibilityHorizontalScrollBar objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalScrollBar:"), accessibilityHorizontalScrollBar)
}

// Returns the vertical scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVerticalScrollBar()
func (o NSMenu) AccessibilityVerticalScrollBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalScrollBar"))
	return objectivec.Object{ID: rv}
}

// Sets the vertical scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVerticalScrollBar(_:)
func (o NSMenu) SetAccessibilityVerticalScrollBar(accessibilityVerticalScrollBar objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalScrollBar:"), accessibilityVerticalScrollBar)
}

// Returns the column header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnHeaderUIElements()
func (o NSMenu) AccessibilityColumnHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the column header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnHeaderUIElements(_:)
func (o NSMenu) SetAccessibilityColumnHeaderUIElements(accessibilityColumnHeaderUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnHeaderUIElements:"), accessibilityColumnHeaderUIElements)
}

// Returns the column accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumns()
func (o NSMenu) AccessibilityColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumns"))
	return foundation.NSArrayFromID(rv)
}

// Sets the column accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumns(_:)
func (o NSMenu) SetAccessibilityColumns(accessibilityColumns objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumns:"), accessibilityColumns)
}

// Returns the column titles for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnTitles()
func (o NSMenu) AccessibilityColumnTitles() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnTitles"))
	return foundation.NSArrayFromID(rv)
}

// Sets the column titles for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnTitles(_:)
func (o NSMenu) SetAccessibilityColumnTitles(accessibilityColumnTitles objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnTitles:"), accessibilityColumnTitles)
}

// Returns a Boolean value that determines whether the accessibility element
// is in an expanded state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityExpanded()
func (o NSMenu) IsAccessibilityExpanded() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityExpanded"))
	return rv
}

// Sets a Boolean value that determines whether accessibility element is in an
// expanded state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityExpanded(_:)
func (o NSMenu) SetAccessibilityExpanded(accessibilityExpanded bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExpanded:"), accessibilityExpanded)
}

// Returns the header for the table view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHeader()
func (o NSMenu) AccessibilityHeader() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHeader"))
	return objectivec.Object{ID: rv}
}

// Sets the header for the table view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHeader(_:)
func (o NSMenu) SetAccessibilityHeader(accessibilityHeader objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHeader:"), accessibilityHeader)
}

// Returns the index of the row or column that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityIndex()
func (o NSMenu) AccessibilityIndex() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityIndex"))
	return rv
}

// Sets the index of the row or column that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityIndex(_:)
func (o NSMenu) SetAccessibilityIndex(accessibilityIndex int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIndex:"), accessibilityIndex)
}

// Returns the row header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRowHeaderUIElements()
func (o NSMenu) AccessibilityRowHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRowHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the row header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRowHeaderUIElements(_:)
func (o NSMenu) SetAccessibilityRowHeaderUIElements(accessibilityRowHeaderUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowHeaderUIElements:"), accessibilityRowHeaderUIElements)
}

// Returns the row accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRows()
func (o NSMenu) AccessibilityRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRows"))
	return foundation.NSArrayFromID(rv)
}

// Sets the row accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRows(_:)
func (o NSMenu) SetAccessibilityRows(accessibilityRows objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRows:"), accessibilityRows)
}

// Returns the currently selected columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedColumns()
func (o NSMenu) AccessibilitySelectedColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedColumns"))
	return foundation.NSArrayFromID(rv)
}

// Sets the currently selected columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedColumns(_:)
func (o NSMenu) SetAccessibilitySelectedColumns(accessibilitySelectedColumns objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedColumns:"), accessibilitySelectedColumns)
}

// Returns the currently selected rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedRows()
func (o NSMenu) AccessibilitySelectedRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedRows"))
	return foundation.NSArrayFromID(rv)
}

// Sets the currently selected rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedRows(_:)
func (o NSMenu) SetAccessibilitySelectedRows(accessibilitySelectedRows objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedRows:"), accessibilitySelectedRows)
}

// Returns the accessibility element’s sort direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySortDirection()
func (o NSMenu) AccessibilitySortDirection() NSAccessibilitySortDirection {
	rv := objc.Send[NSAccessibilitySortDirection](o.ID, objc.Sel("accessibilitySortDirection"))
	return rv
}

// Sets the accessibility element’s sort direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySortDirection(_:)
func (o NSMenu) SetAccessibilitySortDirection(accessibilitySortDirection NSAccessibilitySortDirection) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySortDirection:"), accessibilitySortDirection)
}

// Returns the visible columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleColumns()
func (o NSMenu) AccessibilityVisibleColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleColumns"))
	return foundation.NSArrayFromID(rv)
}

// Sets the visible columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleColumns(_:)
func (o NSMenu) SetAccessibilityVisibleColumns(accessibilityVisibleColumns objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleColumns:"), accessibilityVisibleColumns)
}

// Returns the visible rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleRows()
func (o NSMenu) AccessibilityVisibleRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleRows"))
	return foundation.NSArrayFromID(rv)
}

// Sets the visible rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleRows(_:)
func (o NSMenu) SetAccessibilityVisibleRows(accessibilityVisibleRows objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleRows:"), accessibilityVisibleRows)
}

// Returns a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityDisclosed()
func (o NSMenu) IsAccessibilityDisclosed() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityDisclosed"))
	return rv
}

// Sets a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosed(_:)
func (o NSMenu) SetAccessibilityDisclosed(accessibilityDisclosed bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosed:"), accessibilityDisclosed)
}

// Returns the row disclosing the current row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDisclosedByRow()
func (o NSMenu) AccessibilityDisclosedByRow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedByRow"))
	return objectivec.Object{ID: rv}
}

// Sets the row disclosing the current row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosedByRow(_:)
func (o NSMenu) SetAccessibilityDisclosedByRow(accessibilityDisclosedByRow objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedByRow:"), accessibilityDisclosedByRow)
}

// Returns the rows that the current row discloses.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDisclosedRows()
func (o NSMenu) AccessibilityDisclosedRows() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedRows"))
	return objectivec.Object{ID: rv}
}

// Sets the rows that the current row discloses.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosedRows(_:)
func (o NSMenu) SetAccessibilityDisclosedRows(accessibilityDisclosedRows objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedRows:"), accessibilityDisclosedRows)
}

// Returns the indention level for the row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDisclosureLevel()
func (o NSMenu) AccessibilityDisclosureLevel() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityDisclosureLevel"))
	return rv
}

// Sets the indention level for the row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosureLevel(_:)
func (o NSMenu) SetAccessibilityDisclosureLevel(accessibilityDisclosureLevel int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosureLevel:"), accessibilityDisclosureLevel)
}

// Returns the column index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnIndexRange()
func (o NSMenu) AccessibilityColumnIndexRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityColumnIndexRange"))
	return rv
}

// Sets the column index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnIndexRange(_:)
func (o NSMenu) SetAccessibilityColumnIndexRange(accessibilityColumnIndexRange foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnIndexRange:"), accessibilityColumnIndexRange)
}

// Returns the row index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRowIndexRange()
func (o NSMenu) AccessibilityRowIndexRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRowIndexRange"))
	return rv
}

// Sets the row index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRowIndexRange(_:)
func (o NSMenu) SetAccessibilityRowIndexRange(accessibilityRowIndexRange foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowIndexRange:"), accessibilityRowIndexRange)
}

// Returns the currently selected cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedCells()
func (o NSMenu) AccessibilitySelectedCells() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedCells"))
	return foundation.NSArrayFromID(rv)
}

// Sets the currently selected cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedCells(_:)
func (o NSMenu) SetAccessibilitySelectedCells(accessibilitySelectedCells objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedCells:"), accessibilitySelectedCells)
}

// Returns the visible cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleCells()
func (o NSMenu) AccessibilityVisibleCells() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleCells"))
	return foundation.NSArrayFromID(rv)
}

// Sets the visible cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleCells(_:)
func (o NSMenu) SetAccessibilityVisibleCells(accessibilityVisibleCells objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleCells:"), accessibilityVisibleCells)
}

// Returns the cell at the specified column and row.
//
// column: The column index.
//
// row: The row index.
//
// # Return Value
//
// The cell specified by the column and row indexes.
//
// # Discussion
//
// This property is required for all elements that function as cell-based
// tables.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCell(forColumn:row:)
func (o NSMenu) AccessibilityCellForColumnRow(column int, row int) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCellForColumn:row:"), column, row)
	return objectivec.Object{ID: rv}
}

// Returns the drag handle elements for the layout item element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHandles()
func (o NSMenu) AccessibilityHandles() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHandles"))
	return foundation.NSArrayFromID(rv)
}

// Sets the drag handle accessibility elements for the layout item element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHandles(_:)
func (o NSMenu) SetAccessibilityHandles(accessibilityHandles objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHandles:"), accessibilityHandles)
}

// Returns the units that the layout area uses for horizontal values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHorizontalUnits()
func (o NSMenu) AccessibilityHorizontalUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityHorizontalUnits"))
	return rv
}

// Sets the units that the layout area uses for horizontal values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHorizontalUnits(_:)
func (o NSMenu) SetAccessibilityHorizontalUnits(accessibilityHorizontalUnits NSAccessibilityUnits) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnits:"), accessibilityHorizontalUnits)
}

// Returns the description of the layout area’s horizontal units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHorizontalUnitDescription()
func (o NSMenu) AccessibilityHorizontalUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the description of the layout area’s horizontal units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHorizontalUnitDescription(_:)
func (o NSMenu) SetAccessibilityHorizontalUnitDescription(accessibilityHorizontalUnitDescription string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnitDescription:"), objc.String(accessibilityHorizontalUnitDescription))
}

// Returns the units that the layout area uses for vertical values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVerticalUnits()
func (o NSMenu) AccessibilityVerticalUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityVerticalUnits"))
	return rv
}

// Sets the units that the layout area uses for vertical values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVerticalUnits(_:)
func (o NSMenu) SetAccessibilityVerticalUnits(accessibilityVerticalUnits NSAccessibilityUnits) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnits:"), accessibilityVerticalUnits)
}

// Returns the description of the layout area’s vertical units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVerticalUnitDescription()
func (o NSMenu) AccessibilityVerticalUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the description of the layout area’s vertical units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVerticalUnitDescription(_:)
func (o NSMenu) SetAccessibilityVerticalUnitDescription(accessibilityVerticalUnitDescription string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnitDescription:"), objc.String(accessibilityVerticalUnitDescription))
}

// Converts the provided point in screen coordinates to a point in the layout
// area’s coordinate system.
//
// point: A point in the screen’s coordinate system.
//
// # Return Value
//
// A point in the layout area’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLayoutPoint(forScreenPoint:)
func (o NSMenu) AccessibilityLayoutPointForScreenPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityLayoutPointForScreenPoint:"), point)
	return rv
}

// Converts the provided size in screen coordinates to a size in the layout
// area’s coordinate system.
//
// size: A size in the screen’s coordinate system.
//
// # Return Value
//
// A size in the layout area’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLayoutSize(forScreenSize:)
func (o NSMenu) AccessibilityLayoutSizeForScreenSize(size corefoundation.CGSize) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("accessibilityLayoutSizeForScreenSize:"), size)
	return rv
}

// Converts the provided point in the layout area’s coordinates to a point
// in the screen’s coordinate system.
//
// point: A point in the layout area’s coordinate system.
//
// # Return Value
//
// A point in the screen’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityScreenPoint(forLayoutPoint:)
func (o NSMenu) AccessibilityScreenPointForLayoutPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityScreenPointForLayoutPoint:"), point)
	return rv
}

// Converts the provided size in the layout area’s coordinates to a size in
// the screen’s coordinate system.
//
// size: A size in the layout area’s coordinate system.
//
// # Return Value
//
// A size in the screen’s coordinate system.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityScreenSize(forLayoutSize:)
func (o NSMenu) AccessibilityScreenSizeForLayoutSize(size corefoundation.CGSize) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("accessibilityScreenSizeForLayoutSize:"), size)
	return rv
}

// Returns the allowed values for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityAllowedValues()
func (o NSMenu) AccessibilityAllowedValues() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAllowedValues"))
	return foundation.NSArrayFromID(rv)
}

// Sets the allowed values for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityAllowedValues(_:)
func (o NSMenu) SetAccessibilityAllowedValues(accessibilityAllowedValues objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAllowedValues:"), accessibilityAllowedValues)
}

// Returns the child label elements for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLabelUIElements()
func (o NSMenu) AccessibilityLabelUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabelUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the child label elements for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLabelUIElements(_:)
func (o NSMenu) SetAccessibilityLabelUIElements(accessibilityLabelUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelUIElements:"), accessibilityLabelUIElements)
}

// Returns the value of the label accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLabelValue()
func (o NSMenu) AccessibilityLabelValue() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("accessibilityLabelValue"))
	return rv
}

// Sets the value of the label accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLabelValue(_:)
func (o NSMenu) SetAccessibilityLabelValue(accessibilityLabelValue float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelValue:"), accessibilityLabelValue)
}

// Returns the contents that follow the divider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityNextContents()
func (o NSMenu) AccessibilityNextContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityNextContents"))
	return foundation.NSArrayFromID(rv)
}

// Sets the contents that follow the divider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityNextContents(_:)
func (o NSMenu) SetAccessibilityNextContents(accessibilityNextContents objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNextContents:"), accessibilityNextContents)
}

// Returns the contents that precede the divider accessibility element.
//
// # Return Value
//
// Sets the contents preceding this divider element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPreviousContents()
func (o NSMenu) AccessibilityPreviousContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPreviousContents"))
	return foundation.NSArrayFromID(rv)
}

// Sets the contents that precede the divider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityPreviousContents(_:)
func (o NSMenu) SetAccessibilityPreviousContents(accessibilityPreviousContents objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPreviousContents:"), accessibilityPreviousContents)
}

// Returns an array that contains the views and splitter bar from the split
// view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySplitters()
func (o NSMenu) AccessibilitySplitters() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySplitters"))
	return foundation.NSArrayFromID(rv)
}

// Sets the array that contains the views and splitter bar from the split
// view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySplitters(_:)
func (o NSMenu) SetAccessibilitySplitters(accessibilitySplitters objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySplitters:"), accessibilitySplitters)
}

// Returns the overflow button for the toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityOverflowButton()
func (o NSMenu) AccessibilityOverflowButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityOverflowButton"))
	return objectivec.Object{ID: rv}
}

// Sets the overflow button for the toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityOverflowButton(_:)
func (o NSMenu) SetAccessibilityOverflowButton(accessibilityOverflowButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOverflowButton:"), accessibilityOverflowButton)
}

// Returns the tab accessibility elements for the tab view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTabs()
func (o NSMenu) AccessibilityTabs() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTabs"))
	return foundation.NSArrayFromID(rv)
}

// Sets the tab accessibility elements for the tab view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTabs(_:)
func (o NSMenu) SetAccessibilityTabs(accessibilityTabs objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTabs:"), accessibilityTabs)
}

// Returns the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerGroupUIElement()
func (o NSMenu) AccessibilityMarkerGroupUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerGroupUIElement"))
	return objectivec.Object{ID: rv}
}

// Sets the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerGroupUIElement(_:)
func (o NSMenu) SetAccessibilityMarkerGroupUIElement(accessibilityMarkerGroupUIElement objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerGroupUIElement:"), accessibilityMarkerGroupUIElement)
}

// Returns the human-readable description of the marker type.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerTypeDescription()
func (o NSMenu) AccessibilityMarkerTypeDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerTypeDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the human-readable description of the marker type.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerTypeDescription(_:)
func (o NSMenu) SetAccessibilityMarkerTypeDescription(accessibilityMarkerTypeDescription string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerTypeDescription:"), objc.String(accessibilityMarkerTypeDescription))
}

// Returns the array of marker accessibility elements for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerUIElements()
func (o NSMenu) AccessibilityMarkerUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the array of marker accessibility elements for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerUIElements(_:)
func (o NSMenu) SetAccessibilityMarkerUIElements(accessibilityMarkerUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerUIElements:"), accessibilityMarkerUIElements)
}

// Returns the marker values for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerValues()
func (o NSMenu) AccessibilityMarkerValues() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerValues"))
	return objectivec.Object{ID: rv}
}

// Sets the marker values for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerValues(_:)
func (o NSMenu) SetAccessibilityMarkerValues(accessibilityMarkerValues objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerValues:"), accessibilityMarkerValues)
}

// Returns the type of markers for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRulerMarkerType()
func (o NSMenu) AccessibilityRulerMarkerType() NSAccessibilityRulerMarkerType {
	rv := objc.Send[NSAccessibilityRulerMarkerType](o.ID, objc.Sel("accessibilityRulerMarkerType"))
	return rv
}

// Sets the type of markers for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRulerMarkerType(_:)
func (o NSMenu) SetAccessibilityRulerMarkerType(accessibilityRulerMarkerType NSAccessibilityRulerMarkerType) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRulerMarkerType:"), accessibilityRulerMarkerType)
}

// Returns the units for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityUnits()
func (o NSMenu) AccessibilityUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityUnits"))
	return rv
}

// Sets the units used for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityUnits(_:)
func (o NSMenu) SetAccessibilityUnits(accessibilityUnits NSAccessibilityUnits) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnits:"), accessibilityUnits)
}

// Returns the human-readable description of the ruler’s units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityUnitDescription()
func (o NSMenu) AccessibilityUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the human-readable description of the ruler’s units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityUnitDescription(_:)
func (o NSMenu) SetAccessibilityUnitDescription(accessibilityUnitDescription string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnitDescription:"), objc.String(accessibilityUnitDescription))
}

// Returns the URL for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDocument()
func (o NSMenu) AccessibilityDocument() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDocument"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the URL for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDocument(_:)
func (o NSMenu) SetAccessibilityDocument(accessibilityDocument string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDocument:"), objc.String(accessibilityDocument))
}

// Returns a Boolean value that indicates whether the accessibility element is
// in an edited state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEdited()
func (o NSMenu) IsAccessibilityEdited() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEdited"))
	return rv
}

// Sets a Boolean value that indicates whether the accessibility element is in
// an edited state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityEdited(_:)
func (o NSMenu) SetAccessibilityEdited(accessibilityEdited bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEdited:"), accessibilityEdited)
}

// Returns the filename for the file that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFilename()
func (o NSMenu) AccessibilityFilename() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFilename"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the filename for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFilename(_:)
func (o NSMenu) SetAccessibilityFilename(accessibilityFilename string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFilename:"), objc.String(accessibilityFilename))
}

// Returns the elements that have links with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLinkedUIElements()
func (o NSMenu) AccessibilityLinkedUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLinkedUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the elements that have links with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLinkedUIElements(_:)
func (o NSMenu) SetAccessibilityLinkedUIElements(accessibilityLinkedUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLinkedUIElements:"), accessibilityLinkedUIElements)
}

// Returns the list of elements that the accessibility element is a title for.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityServesAsTitleForUIElements()
func (o NSMenu) AccessibilityServesAsTitleForUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityServesAsTitleForUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the list of elements that the accessibility element is a title for.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityServesAsTitleForUIElements(_:)
func (o NSMenu) SetAccessibilityServesAsTitleForUIElements(accessibilityServesAsTitleForUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityServesAsTitleForUIElements:"), accessibilityServesAsTitleForUIElements)
}

// Returns the static text element that represents the accessibility
// element’s title.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTitleUIElement()
func (o NSMenu) AccessibilityTitleUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitleUIElement"))
	return objectivec.Object{ID: rv}
}

// Sets the static text element that represents the accessibility element’s
// title.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTitleUIElement(_:)
func (o NSMenu) SetAccessibilityTitleUIElement(accessibilityTitleUIElement objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitleUIElement:"), accessibilityTitleUIElement)
}

// Returns the clear button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityClearButton()
func (o NSMenu) AccessibilityClearButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityClearButton"))
	return objectivec.Object{ID: rv}
}

// Sets the clear button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityClearButton(_:)
func (o NSMenu) SetAccessibilityClearButton(accessibilityClearButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityClearButton:"), accessibilityClearButton)
}

// Returns the search button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySearchButton()
func (o NSMenu) AccessibilitySearchButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchButton"))
	return objectivec.Object{ID: rv}
}

// Sets the search button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySearchButton(_:)
func (o NSMenu) SetAccessibilitySearchButton(accessibilitySearchButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchButton:"), accessibilitySearchButton)
}

// Returns the search menu for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySearchMenu()
func (o NSMenu) AccessibilitySearchMenu() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchMenu"))
	return objectivec.Object{ID: rv}
}

// Sets the search menu for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySearchMenu(_:)
func (o NSMenu) SetAccessibilitySearchMenu(accessibilitySearchMenu objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchMenu:"), accessibilitySearchMenu)
}

// Cancels the current operation.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformCancel()
func (o NSMenu) AccessibilityPerformCancel() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformCancel"))
	return rv
}

// Simulates pressing Return in the accessibility element.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on elements that take keyboard input, such as a text field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformConfirm()
func (o NSMenu) AccessibilityPerformConfirm() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformConfirm"))
	return rv
}

// Selects the accessibility element.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on selectable elements, such as a menu item.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformPick()
func (o NSMenu) AccessibilityPerformPick() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformPick"))
	return rv
}

// Simulates clicking the accessibility element.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on elements that behave like buttons.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformPress()
func (o NSMenu) AccessibilityPerformPress() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformPress"))
	return rv
}

// Displays the accessibility element’s alternative UI.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method to trigger changes to the UI due to a mouse-hover or
// similar event.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowAlternateUI()
func (o NSMenu) AccessibilityPerformShowAlternateUI() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowAlternateUI"))
	return rv
}

// Returns to the accessibility element’s original UI.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Call this method after successfully calling
// [AccessibilityPerformShowAlternateUI] to return to the original UI.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowDefaultUI()
func (o NSMenu) AccessibilityPerformShowDefaultUI() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowDefaultUI"))
	return rv
}

// Displays the menu accessibility element.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method to display the contextual menu for the element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowMenu()
func (o NSMenu) AccessibilityPerformShowMenu() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowMenu"))
	return rv
}

// Brings the window to the front.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// The window behaves as if you had clicked on the window’s title bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformRaise()
func (o NSMenu) AccessibilityPerformRaise() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformRaise"))
	return rv
}

// Returns the increment button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityIncrementButton()
func (o NSMenu) AccessibilityIncrementButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIncrementButton"))
	return objectivec.Object{ID: rv}
}

// Sets the increment button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityIncrementButton(_:)
func (o NSMenu) SetAccessibilityIncrementButton(accessibilityIncrementButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIncrementButton:"), accessibilityIncrementButton)
}

// Returns the decrement button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDecrementButton()
func (o NSMenu) AccessibilityDecrementButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDecrementButton"))
	return objectivec.Object{ID: rv}
}

// Sets the decrement button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDecrementButton(_:)
func (o NSMenu) SetAccessibilityDecrementButton(accessibilityDecrementButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDecrementButton:"), accessibilityDecrementButton)
}

// Increments the accessibility element’s value.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on elements that have an adjustable [accessibilityValue]
// property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformIncrement()
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
func (o NSMenu) AccessibilityPerformIncrement() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformIncrement"))
	return rv
}

// Decrements the accessibility element’s value.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on elements that have an adjustable [accessibilityValue]
// property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDecrement()
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
func (o NSMenu) AccessibilityPerformDecrement() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformDecrement"))
	return rv
}

// Deletes the accessibility element’s value.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// Use this method on elements with values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDelete()
func (o NSMenu) AccessibilityPerformDelete() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformDelete"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityAttributedUserInputLabels()
func (o NSMenu) AccessibilityAttributedUserInputLabels() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAttributedUserInputLabels"))
	return foundation.NSArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityUserInputLabels()
func (o NSMenu) AccessibilityUserInputLabels() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUserInputLabels"))
	return foundation.NSArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityAttributedUserInputLabels(_:)
func (o NSMenu) SetAccessibilityAttributedUserInputLabels(accessibilityAttributedUserInputLabels objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAttributedUserInputLabels:"), accessibilityAttributedUserInputLabels)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityUserInputLabels(_:)
func (o NSMenu) SetAccessibilityUserInputLabels(accessibilityUserInputLabels objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUserInputLabels:"), accessibilityUserInputLabels)
}

// Protocol methods for NSAppearanceCustomization

// Protocol methods for NSUserInterfaceItemIdentification
