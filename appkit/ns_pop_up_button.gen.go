// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPopUpButton] class.
var (
	_NSPopUpButtonClass     NSPopUpButtonClass
	_NSPopUpButtonClassOnce sync.Once
)

func getNSPopUpButtonClass() NSPopUpButtonClass {
	_NSPopUpButtonClassOnce.Do(func() {
		_NSPopUpButtonClass = NSPopUpButtonClass{class: objc.GetClass("NSPopUpButton")}
	})
	return _NSPopUpButtonClass
}

// GetNSPopUpButtonClass returns the class object for NSPopUpButton.
func GetNSPopUpButtonClass() NSPopUpButtonClass {
	return getNSPopUpButtonClass()
}

type NSPopUpButtonClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPopUpButtonClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPopUpButtonClass) Alloc() NSPopUpButton {
	rv := objc.Send[NSPopUpButton](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A control for selecting an item from a list.
//
// # Overview
// 
// An [NSPopUpButton] object uses an [NSPopUpButtonCell] object to implement
// its user interface.
// 
// Note that while a menu is tracking user input, programmatic changes to the
// menu, such as adding, removing, or changing items on the menu, is not
// reflected.
//
// # Initializing an NSPopUpButton
//
//   - [NSPopUpButton.InitWithFramePullsDown]: Returns an [NSPopUpButton] object initialized to the specified dimensions.
//
// # Setting the type of menu
//
//   - [NSPopUpButton.PullsDown]: A Boolean value indicating whether the button displays a pull-down or pop-up menu.
//   - [NSPopUpButton.SetPullsDown]
//   - [NSPopUpButton.AutoenablesItems]: A Boolean value indicating whether the button enables and disables its items every time a user event occurs.
//   - [NSPopUpButton.SetAutoenablesItems]
//
// # Inserting and deleting items
//
//   - [NSPopUpButton.AddItemWithTitle]: Adds an item with the specified title to the end of the menu.
//   - [NSPopUpButton.AddItemsWithTitles]: Adds multiple items to the end of the menu.
//   - [NSPopUpButton.InsertItemWithTitleAtIndex]: Inserts an item at the specified position in the menu.
//   - [NSPopUpButton.RemoveAllItems]: Removes all items in the receiver’s item menu.
//   - [NSPopUpButton.RemoveItemWithTitle]: Removes the item with the specified title from the menu.
//   - [NSPopUpButton.RemoveItemAtIndex]: Removes the item at the specified index.
//
// # Getting the user’s selection
//
//   - [NSPopUpButton.SelectedItem]: The menu item that was last selected by the user.
//   - [NSPopUpButton.TitleOfSelectedItem]: The title of the item that was last selected by the user.
//   - [NSPopUpButton.IndexOfSelectedItem]: The index of the item that was last selected by the user.
//
// # Setting the current selection
//
//   - [NSPopUpButton.SelectItem]: Selects the specified menu item.
//   - [NSPopUpButton.SelectItemAtIndex]: Selects the item in the menu at the specified index.
//   - [NSPopUpButton.SelectItemWithTag]: Selects the menu item with the specified tag.
//   - [NSPopUpButton.SelectItemWithTitle]: Selects the item with the specified title.
//
// # Getting menu items
//
//   - [NSPopUpButton.NumberOfItems]: The number of items in the menu.
//   - [NSPopUpButton.ItemArray]: The array of menu item objects associated with the button.
//   - [NSPopUpButton.ItemAtIndex]: Returns the menu item at the specified index.
//   - [NSPopUpButton.ItemTitleAtIndex]: Returns the title of the item at the specified index.
//   - [NSPopUpButton.ItemTitles]: An array of strings corresponding to the titles of the items in the menu.
//   - [NSPopUpButton.ItemWithTitle]: Returns the menu item with the specified title.
//   - [NSPopUpButton.LastItem]: The last item in the menu.
//
// # Getting the indices of menu items
//
//   - [NSPopUpButton.IndexOfItem]: Returns the index of the specified menu item.
//   - [NSPopUpButton.IndexOfItemWithTag]: Returns the index of the menu item with the specified tag.
//   - [NSPopUpButton.IndexOfItemWithTitle]: Returns the index of the item with the specified title.
//   - [NSPopUpButton.IndexOfItemWithRepresentedObject]: Returns the index of the menu item that holds the specified represented object.
//   - [NSPopUpButton.IndexOfItemWithTargetAndAction]: Returns the index of the menu item with the specified target and action.
//
// # Setting the cell edge to pop out in restricted situations
//
//   - [NSPopUpButton.PreferredEdge]: The edge of the button on which to display the menu when screen space is constrained.
//   - [NSPopUpButton.SetPreferredEdge]
//
// # Setting the state
//
//   - [NSPopUpButton.SynchronizeTitleAndSelectedItem]: Ensures that the item being displayed by the receiver agrees with the selected item.
//
// # Instance Properties
//
//   - [NSPopUpButton.AltersStateOfSelectedItem]: When the value of this property is [YES], the selected menu item’s `state` is set to [NSControlStateValueOn]. When the value of this property is [NO], the menu item’s `state` is not changed. When this property changes, the `state` of the currently selected item is updated appropriately. This property is ignored for pull-down buttons.
//   - [NSPopUpButton.SetAltersStateOfSelectedItem]
//   - [NSPopUpButton.UsesItemFromMenu]: When `usesItemFromMenu` is [YES], a pull-down button uses the title of the first menu item and hides the first menu item. A pop-up button uses the title of the currently selected menu. The default value is [YES].
//   - [NSPopUpButton.SetUsesItemFromMenu]
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton
type NSPopUpButton struct {
	NSButton
}

// NSPopUpButtonFromID constructs a [NSPopUpButton] from an objc.ID.
//
// A control for selecting an item from a list.
func NSPopUpButtonFromID(id objc.ID) NSPopUpButton {
	return NSPopUpButton{NSButton: NSButtonFromID(id)}
}
// NOTE: NSPopUpButton adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPopUpButton] class.
//
// # Initializing an NSPopUpButton
//
//   - [INSPopUpButton.InitWithFramePullsDown]: Returns an [NSPopUpButton] object initialized to the specified dimensions.
//
// # Setting the type of menu
//
//   - [INSPopUpButton.PullsDown]: A Boolean value indicating whether the button displays a pull-down or pop-up menu.
//   - [INSPopUpButton.SetPullsDown]
//   - [INSPopUpButton.AutoenablesItems]: A Boolean value indicating whether the button enables and disables its items every time a user event occurs.
//   - [INSPopUpButton.SetAutoenablesItems]
//
// # Inserting and deleting items
//
//   - [INSPopUpButton.AddItemWithTitle]: Adds an item with the specified title to the end of the menu.
//   - [INSPopUpButton.AddItemsWithTitles]: Adds multiple items to the end of the menu.
//   - [INSPopUpButton.InsertItemWithTitleAtIndex]: Inserts an item at the specified position in the menu.
//   - [INSPopUpButton.RemoveAllItems]: Removes all items in the receiver’s item menu.
//   - [INSPopUpButton.RemoveItemWithTitle]: Removes the item with the specified title from the menu.
//   - [INSPopUpButton.RemoveItemAtIndex]: Removes the item at the specified index.
//
// # Getting the user’s selection
//
//   - [INSPopUpButton.SelectedItem]: The menu item that was last selected by the user.
//   - [INSPopUpButton.TitleOfSelectedItem]: The title of the item that was last selected by the user.
//   - [INSPopUpButton.IndexOfSelectedItem]: The index of the item that was last selected by the user.
//
// # Setting the current selection
//
//   - [INSPopUpButton.SelectItem]: Selects the specified menu item.
//   - [INSPopUpButton.SelectItemAtIndex]: Selects the item in the menu at the specified index.
//   - [INSPopUpButton.SelectItemWithTag]: Selects the menu item with the specified tag.
//   - [INSPopUpButton.SelectItemWithTitle]: Selects the item with the specified title.
//
// # Getting menu items
//
//   - [INSPopUpButton.NumberOfItems]: The number of items in the menu.
//   - [INSPopUpButton.ItemArray]: The array of menu item objects associated with the button.
//   - [INSPopUpButton.ItemAtIndex]: Returns the menu item at the specified index.
//   - [INSPopUpButton.ItemTitleAtIndex]: Returns the title of the item at the specified index.
//   - [INSPopUpButton.ItemTitles]: An array of strings corresponding to the titles of the items in the menu.
//   - [INSPopUpButton.ItemWithTitle]: Returns the menu item with the specified title.
//   - [INSPopUpButton.LastItem]: The last item in the menu.
//
// # Getting the indices of menu items
//
//   - [INSPopUpButton.IndexOfItem]: Returns the index of the specified menu item.
//   - [INSPopUpButton.IndexOfItemWithTag]: Returns the index of the menu item with the specified tag.
//   - [INSPopUpButton.IndexOfItemWithTitle]: Returns the index of the item with the specified title.
//   - [INSPopUpButton.IndexOfItemWithRepresentedObject]: Returns the index of the menu item that holds the specified represented object.
//   - [INSPopUpButton.IndexOfItemWithTargetAndAction]: Returns the index of the menu item with the specified target and action.
//
// # Setting the cell edge to pop out in restricted situations
//
//   - [INSPopUpButton.PreferredEdge]: The edge of the button on which to display the menu when screen space is constrained.
//   - [INSPopUpButton.SetPreferredEdge]
//
// # Setting the state
//
//   - [INSPopUpButton.SynchronizeTitleAndSelectedItem]: Ensures that the item being displayed by the receiver agrees with the selected item.
//
// # Instance Properties
//
//   - [INSPopUpButton.AltersStateOfSelectedItem]: When the value of this property is [YES], the selected menu item’s `state` is set to [NSControlStateValueOn]. When the value of this property is [NO], the menu item’s `state` is not changed. When this property changes, the `state` of the currently selected item is updated appropriately. This property is ignored for pull-down buttons.
//   - [INSPopUpButton.SetAltersStateOfSelectedItem]
//   - [INSPopUpButton.UsesItemFromMenu]: When `usesItemFromMenu` is [YES], a pull-down button uses the title of the first menu item and hides the first menu item. A pop-up button uses the title of the currently selected menu. The default value is [YES].
//   - [INSPopUpButton.SetUsesItemFromMenu]
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton
type INSPopUpButton interface {
	INSButton

	// Topic: Initializing an NSPopUpButton

	// Returns an [NSPopUpButton] object initialized to the specified dimensions.
	InitWithFramePullsDown(buttonFrame corefoundation.CGRect, flag bool) NSPopUpButton

	// Topic: Setting the type of menu

	// A Boolean value indicating whether the button displays a pull-down or pop-up menu.
	PullsDown() bool
	SetPullsDown(value bool)
	// A Boolean value indicating whether the button enables and disables its items every time a user event occurs.
	AutoenablesItems() bool
	SetAutoenablesItems(value bool)

	// Topic: Inserting and deleting items

	// Adds an item with the specified title to the end of the menu.
	AddItemWithTitle(title string)
	// Adds multiple items to the end of the menu.
	AddItemsWithTitles(itemTitles []string)
	// Inserts an item at the specified position in the menu.
	InsertItemWithTitleAtIndex(title string, index int)
	// Removes all items in the receiver’s item menu.
	RemoveAllItems()
	// Removes the item with the specified title from the menu.
	RemoveItemWithTitle(title string)
	// Removes the item at the specified index.
	RemoveItemAtIndex(index int)

	// Topic: Getting the user’s selection

	// The menu item that was last selected by the user.
	SelectedItem() INSMenuItem
	// The title of the item that was last selected by the user.
	TitleOfSelectedItem() string
	// The index of the item that was last selected by the user.
	IndexOfSelectedItem() int

	// Topic: Setting the current selection

	// Selects the specified menu item.
	SelectItem(item INSMenuItem)
	// Selects the item in the menu at the specified index.
	SelectItemAtIndex(index int)
	// Selects the menu item with the specified tag.
	SelectItemWithTag(tag int) bool
	// Selects the item with the specified title.
	SelectItemWithTitle(title string)

	// Topic: Getting menu items

	// The number of items in the menu.
	NumberOfItems() int
	// The array of menu item objects associated with the button.
	ItemArray() []NSMenuItem
	// Returns the menu item at the specified index.
	ItemAtIndex(index int) INSMenuItem
	// Returns the title of the item at the specified index.
	ItemTitleAtIndex(index int) string
	// An array of strings corresponding to the titles of the items in the menu.
	ItemTitles() []string
	// Returns the menu item with the specified title.
	ItemWithTitle(title string) INSMenuItem
	// The last item in the menu.
	LastItem() INSMenuItem

	// Topic: Getting the indices of menu items

	// Returns the index of the specified menu item.
	IndexOfItem(item INSMenuItem) int
	// Returns the index of the menu item with the specified tag.
	IndexOfItemWithTag(tag int) int
	// Returns the index of the item with the specified title.
	IndexOfItemWithTitle(title string) int
	// Returns the index of the menu item that holds the specified represented object.
	IndexOfItemWithRepresentedObject(obj objectivec.IObject) int
	// Returns the index of the menu item with the specified target and action.
	IndexOfItemWithTargetAndAction(target objectivec.IObject, actionSelector objc.SEL) int

	// Topic: Setting the cell edge to pop out in restricted situations

	// The edge of the button on which to display the menu when screen space is constrained.
	PreferredEdge() foundation.NSRectEdge
	SetPreferredEdge(value foundation.NSRectEdge)

	// Topic: Setting the state

	// Ensures that the item being displayed by the receiver agrees with the selected item.
	SynchronizeTitleAndSelectedItem()

	// Topic: Instance Properties

	// When the value of this property is [YES], the selected menu item’s `state` is set to [NSControlStateValueOn]. When the value of this property is [NO], the menu item’s `state` is not changed. When this property changes, the `state` of the currently selected item is updated appropriately. This property is ignored for pull-down buttons.
	AltersStateOfSelectedItem() bool
	SetAltersStateOfSelectedItem(value bool)
	// When `usesItemFromMenu` is [YES], a pull-down button uses the title of the first menu item and hides the first menu item. A pop-up button uses the title of the currently selected menu. The default value is [YES].
	UsesItemFromMenu() bool
	SetUsesItemFromMenu(value bool)
}

// Init initializes the instance.
func (p NSPopUpButton) Init() NSPopUpButton {
	rv := objc.Send[NSPopUpButton](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPopUpButton) Autorelease() NSPopUpButton {
	rv := objc.Send[NSPopUpButton](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPopUpButton creates a new NSPopUpButton instance.
func NewNSPopUpButton() NSPopUpButton {
	class := getNSPopUpButtonClass()
	rv := objc.Send[NSPopUpButton](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a standard checkbox with the title you specify.
//
// title: The localized title string to display on the button.
//
// target: The target object that receives action messages from the button.
//
// action: The action message the button sends to the target.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/init(checkboxWithTitle:target:action:)
func NewPopUpButtonCheckboxWithTitleTargetAction(title string, target objectivec.IObject, action objc.SEL) NSPopUpButton {
	rv := objc.Send[objc.ID](objc.ID(getNSPopUpButtonClass().class), objc.Sel("checkboxWithTitle:target:action:"), objc.String(title), target, action)
	return NSPopUpButtonFromID(rv)
}

// Creates a standard radio button with the title you specify.
//
// title: The localized title string to display on the button.
//
// target: The target object that receives action messages from the button.
//
// action: The action message the button sends to the target.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/init(radioButtonWithTitle:target:action:)
func NewPopUpButtonRadioButtonWithTitleTargetAction(title string, target objectivec.IObject, action objc.SEL) NSPopUpButton {
	rv := objc.Send[objc.ID](objc.ID(getNSPopUpButtonClass().class), objc.Sel("radioButtonWithTitle:target:action:"), objc.String(title), target, action)
	return NSPopUpButtonFromID(rv)
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewPopUpButtonWithCoder(coder foundation.INSCoder) NSPopUpButton {
	instance := getNSPopUpButtonClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSPopUpButtonFromID(rv)
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
func NewPopUpButtonWithFrame(frameRect corefoundation.CGRect) NSPopUpButton {
	instance := getNSPopUpButtonClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSPopUpButtonFromID(rv)
}

// Returns an [NSPopUpButton] object initialized to the specified dimensions.
//
// buttonFrame: The frame rectangle for the button, specified in the parent view’s
// coordinate system.
//
// flag: [true] if you want the receiver to display a pull-down menu; otherwise,
// [false] if you want it to display a pop-up menu.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An initialized [NSPopUpButton] object, or `nil` if the object could not be
// initialized.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/init(frame:pullsDown:)
func NewPopUpButtonWithFramePullsDown(buttonFrame corefoundation.CGRect, flag bool) NSPopUpButton {
	instance := getNSPopUpButtonClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:pullsDown:"), buttonFrame, flag)
	return NSPopUpButtonFromID(rv)
}

// Creates a standard push button with the image you specify.
//
// image: The image to display in the body of the button.
//
// target: The target object that receives action messages from the button.
//
// action: The action message the button sends to the target.
//
// # Discussion
// 
// Set the image’s [AccessibilityDescription] property to ensure
// accessibility for this control.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/init(image:target:action:)
func NewPopUpButtonWithImageTargetAction(image INSImage, target objectivec.IObject, action objc.SEL) NSPopUpButton {
	rv := objc.Send[objc.ID](objc.ID(getNSPopUpButtonClass().class), objc.Sel("buttonWithImage:target:action:"), image, target, action)
	return NSPopUpButtonFromID(rv)
}

// Creates a standard push button with a title and image.
//
// title: The localized title string to display on the button.
//
// image: The image to display in the body of the button.
//
// target: The target object that receives action messages from the button.
//
// action: The action message the button sends to the target.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/init(title:image:target:action:)
func NewPopUpButtonWithTitleImageTargetAction(title string, image INSImage, target objectivec.IObject, action objc.SEL) NSPopUpButton {
	rv := objc.Send[objc.ID](objc.ID(getNSPopUpButtonClass().class), objc.Sel("buttonWithTitle:image:target:action:"), objc.String(title), image, target, action)
	return NSPopUpButtonFromID(rv)
}

// Creates a standard push button with the title you specify.
//
// title: The localized title string to display on the button.
//
// target: The target object that receives action messages from the control.
//
// action: The action message the button sends to the target.
//
// See: https://developer.apple.com/documentation/AppKit/NSButton/init(title:target:action:)
func NewPopUpButtonWithTitleTargetAction(title string, target objectivec.IObject, action objc.SEL) NSPopUpButton {
	rv := objc.Send[objc.ID](objc.ID(getNSPopUpButtonClass().class), objc.Sel("buttonWithTitle:target:action:"), objc.String(title), target, action)
	return NSPopUpButtonFromID(rv)
}

// Returns an [NSPopUpButton] object initialized to the specified dimensions.
//
// buttonFrame: The frame rectangle for the button, specified in the parent view’s
// coordinate system.
//
// flag: [true] if you want the receiver to display a pull-down menu; otherwise,
// [false] if you want it to display a pop-up menu.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// An initialized [NSPopUpButton] object, or `nil` if the object could not be
// initialized.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/init(frame:pullsDown:)
func (p NSPopUpButton) InitWithFramePullsDown(buttonFrame corefoundation.CGRect, flag bool) NSPopUpButton {
	rv := objc.Send[NSPopUpButton](p.ID, objc.Sel("initWithFrame:pullsDown:"), buttonFrame, flag)
	return rv
}
// Adds an item with the specified title to the end of the menu.
//
// title: The title of the menu-item entry. If an item with the same title already
// exists in the menu, the existing item is removed and the new one is added.
//
// # Discussion
// 
// If you want to move an item, it’s better to invoke [RemoveItemWithTitle]
// explicitly and then send this method. After adding the item, this method
// calls the [SynchronizeTitleAndSelectedItem] method to make sure the item
// being displayed matches the currently selected item.
// 
// Since this method searches for duplicate items, it should not be used if
// you are adding an item to an already populated menu with more than a few
// hundred items. Add items directly to the receiver’s menu instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/addItem(withTitle:)
func (p NSPopUpButton) AddItemWithTitle(title string) {
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
// If you want to move an item, it’s better to invoke [RemoveItemWithTitle]
// explicitly and then send this method. After adding the items, this method
// uses the [SynchronizeTitleAndSelectedItem] method to make sure the item
// being displayed matches the currently selected item.
// 
// Since this method searches for duplicate items, it should not be used if
// you are adding items to an already populated menu with more than a few
// hundred items. Add items directly to the receiver’s menu instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/addItems(withTitles:)
func (p NSPopUpButton) AddItemsWithTitles(itemTitles []string) {
	objc.Send[objc.ID](p.ID, objc.Sel("addItemsWithTitles:"), objectivec.StringSliceToNSArray(itemTitles))
}
// Inserts an item at the specified position in the menu.
//
// title: The title of the new item. If an item with the same title already exists in
// the menu, the existing item is removed and the new one is added
//
// index: The zero-based index at which to insert the item. Specifying 0 inserts the
// item at the top of the menu.
//
// # Discussion
// 
// If you want to move an item, it’s better to invoke [RemoveItemWithTitle]
// explicitly and then send this method. After adding the item, this method
// uses the [SynchronizeTitleAndSelectedItem] method to make sure the item
// displayed matches the currently selected item.
// 
// Since this method searches for duplicate items, it should not be used if
// you are adding an item to an already populated menu with more than a few
// hundred items. Add items directly to the receiver’s menu instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/insertItem(withTitle:at:)
func (p NSPopUpButton) InsertItemWithTitleAtIndex(title string, index int) {
	objc.Send[objc.ID](p.ID, objc.Sel("insertItemWithTitle:atIndex:"), objc.String(title), index)
}
// Removes all items in the receiver’s item menu.
//
// # Discussion
// 
// After removing the items, this method uses the
// [SynchronizeTitleAndSelectedItem] method to refresh the menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/removeAllItems()
func (p NSPopUpButton) RemoveAllItems() {
	objc.Send[objc.ID](p.ID, objc.Sel("removeAllItems"))
}
// Removes the item with the specified title from the menu.
//
// title: The title of the item you want to remove. If no menu item exists with the
// specified title, this method triggers an assertion.
//
// # Discussion
// 
// This method removes the first item it finds with the specified name. This
// method then uses [SynchronizeTitleAndSelectedItem] to refresh the menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/removeItem(withTitle:)
func (p NSPopUpButton) RemoveItemWithTitle(title string) {
	objc.Send[objc.ID](p.ID, objc.Sel("removeItemWithTitle:"), objc.String(title))
}
// Removes the item at the specified index.
//
// index: The zero-based index indicating which item to remove. Specifying 0 removes
// the item at the top of the menu.
//
// # Discussion
// 
// After removing the item, this method uses the
// [SynchronizeTitleAndSelectedItem] method to make sure the title displayed
// matches the currently selected item.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/removeItem(at:)
func (p NSPopUpButton) RemoveItemAtIndex(index int) {
	objc.Send[objc.ID](p.ID, objc.Sel("removeItemAtIndex:"), index)
}
// Selects the specified menu item.
//
// item: The menu item to select, or `nil` if you want to deselect all menu items.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/select(_:)
func (p NSPopUpButton) SelectItem(item INSMenuItem) {
	objc.Send[objc.ID](p.ID, objc.Sel("selectItem:"), item)
}
// Selects the item in the menu at the specified index.
//
// index: The index of the item you want to select, or `-1` you want to deselect all
// menu items.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/selectItem(at:)
func (p NSPopUpButton) SelectItemAtIndex(index int) {
	objc.Send[objc.ID](p.ID, objc.Sel("selectItemAtIndex:"), index)
}
// Selects the menu item with the specified tag.
//
// tag: The tag of the item you want to select.
//
// # Return Value
// 
// [true] if the item was successfully selected; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If no item with the specified tag is found, this method returns [false] and
// leaves the menu state unchanged.
// 
// You typically assign tags to menu items from Interface Builder, but you can
// also assign them programmatically using the setTag: method of [NSMenuItem].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/selectItem(withTag:)
func (p NSPopUpButton) SelectItemWithTag(tag int) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("selectItemWithTag:"), tag)
	return rv
}
// Selects the item with the specified title.
//
// title: The title of the item to select. If you specify an empty string, or a
// string that does not match the title of a menu item, this method deselects
// the currently selected item.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/selectItem(withTitle:)
func (p NSPopUpButton) SelectItemWithTitle(title string) {
	objc.Send[objc.ID](p.ID, objc.Sel("selectItemWithTitle:"), objc.String(title))
}
// Returns the menu item at the specified index.
//
// index: The index of the item you want.
//
// # Return Value
// 
// The menu item, or `nil` if no item exists at the specified index.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/item(at:)
func (p NSPopUpButton) ItemAtIndex(index int) INSMenuItem {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("itemAtIndex:"), index)
	return NSMenuItemFromID(rv)
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
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/itemTitle(at:)
func (p NSPopUpButton) ItemTitleAtIndex(index int) string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("itemTitleAtIndex:"), index)
	return foundation.NSStringFromID(rv).String()
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
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/item(withTitle:)
func (p NSPopUpButton) ItemWithTitle(title string) INSMenuItem {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("itemWithTitle:"), objc.String(title))
	return NSMenuItemFromID(rv)
}
// Returns the index of the specified menu item.
//
// item: The menu item whose index you want.
//
// # Return Value
// 
// The index of the item or `-1` if no such item was found.
//
// # Discussion
// 
// This method invokes the method of the same name of its [NSPopUpButtonCell]
// object.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/index(of:)
func (p NSPopUpButton) IndexOfItem(item INSMenuItem) int {
	rv := objc.Send[int](p.ID, objc.Sel("indexOfItem:"), item)
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
// This method invokes the method of the same name of its [NSPopUpButtonCell]
// object.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/indexOfItem(withTag:)
func (p NSPopUpButton) IndexOfItemWithTag(tag int) int {
	rv := objc.Send[int](p.ID, objc.Sel("indexOfItemWithTag:"), tag)
	return rv
}
// Returns the index of the item with the specified title.
//
// title: The title of the item you want.
//
// # Return Value
// 
// The index of the item or `-1` if no item with the specified title was
// found.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/indexOfItem(withTitle:)
func (p NSPopUpButton) IndexOfItemWithTitle(title string) int {
	rv := objc.Send[int](p.ID, objc.Sel("indexOfItemWithTitle:"), objc.String(title))
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
// # Discussion
// 
// Represented objects bear some direct relation to the title or image of a
// menu item; for example, an item entitled “100” might have an [NSNumber]
// object encapsulating that value as its represented object. This method
// invokes the method of the same name of its [NSPopUpButtonCell] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/indexOfItem(withRepresentedObject:)
func (p NSPopUpButton) IndexOfItemWithRepresentedObject(obj objectivec.IObject) int {
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
// first menu item with the specified target is returned. This method invokes
// the method of the same name of its [NSPopUpButtonCell] object.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/indexOfItem(withTarget:andAction:)
func (p NSPopUpButton) IndexOfItemWithTargetAndAction(target objectivec.IObject, actionSelector objc.SEL) int {
	rv := objc.Send[int](p.ID, objc.Sel("indexOfItemWithTarget:andAction:"), target, actionSelector)
	return rv
}
// Ensures that the item being displayed by the receiver agrees with the
// selected item.
//
// # Discussion
// 
// If there’s no selected item, this method selects the first item in the
// item menu and sets the receiver’s item to match. For pull-down menus,
// this method makes sure that the first item is being displayed (the
// [NSPopUpButtonCell] object must be set to use the selected menu item, which
// happens by default).
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/synchronizeTitleAndSelectedItem()
func (p NSPopUpButton) SynchronizeTitleAndSelectedItem() {
	objc.Send[objc.ID](p.ID, objc.Sel("synchronizeTitleAndSelectedItem"))
}

// A Boolean value indicating whether the button displays a pull-down or
// pop-up menu.
//
// # Discussion
// 
// When the value of this property is [true], the button displays a pull-down
// menu; otherwise, it displays a pop-up menu. This property does not affect
// the contents of the menu; it affects only the style of the menu.
// 
// When changing the menu type to a pull-down menu, if the menu was a pop-up
// menu and the cell alters the state of its selected items, this method sets
// the state of the currently selected item to [NSStateOff] before changing
// the menu type.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/pullsDown
func (p NSPopUpButton) PullsDown() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("pullsDown"))
	return rv
}
func (p NSPopUpButton) SetPullsDown(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setPullsDown:"), value)
}
// A Boolean value indicating whether the button enables and disables its
// items every time a user event occurs.
//
// # Discussion
// 
// When the value of this property is [true], user events cause the button to
// enable and disable its items automatically according to the
// NSMenuValidation protocol specification.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/autoenablesItems
func (p NSPopUpButton) AutoenablesItems() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("autoenablesItems"))
	return rv
}
func (p NSPopUpButton) SetAutoenablesItems(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAutoenablesItems:"), value)
}
// The menu item that was last selected by the user.
//
// # Discussion
// 
// The last selected menu item is the one that was highlighted when the user
// released the mouse button. It is possible for a pull-down menu’s selected
// item to be its first item. If no item is selected, the value in this
// property is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/selectedItem
func (p NSPopUpButton) SelectedItem() INSMenuItem {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("selectedItem"))
	return NSMenuItemFromID(objc.ID(rv))
}
// The title of the item that was last selected by the user.
//
// # Discussion
// 
// If no item is selected, the value in this property is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/titleOfSelectedItem
func (p NSPopUpButton) TitleOfSelectedItem() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("titleOfSelectedItem"))
	return foundation.NSStringFromID(rv).String()
}
// The index of the item that was last selected by the user.
//
// # Discussion
// 
// If no item is selected, the value in this property is `-1`.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/indexOfSelectedItem
func (p NSPopUpButton) IndexOfSelectedItem() int {
	rv := objc.Send[int](p.ID, objc.Sel("indexOfSelectedItem"))
	return rv
}
// The number of items in the menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/numberOfItems
func (p NSPopUpButton) NumberOfItems() int {
	rv := objc.Send[int](p.ID, objc.Sel("numberOfItems"))
	return rv
}
// The array of menu item objects associated with the button.
//
// # Discussion
// 
// This property contains an array of [NSMenuItem] objects representing the
// items in the menu. Usually, you access menu items using the methods and
// properties of this class rather than accessing the items directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/itemArray
func (p NSPopUpButton) ItemArray() []NSMenuItem {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("itemArray"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSMenuItem {
		return NSMenuItemFromID(id)
	})
}
// An array of strings corresponding to the titles of the items in the menu.
//
// # Discussion
// 
// This property contains an array of [NSString] objects, each of which
// contains the title of an item in the menu. The order of the titles in this
// array matches the order of the items in the menu. If the menu contains
// separator items, the array contains an empty string for each separator
// item.
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/itemTitles
func (p NSPopUpButton) ItemTitles() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("itemTitles"))
	return objc.ConvertSliceToStrings(rv)
}
// The last item in the menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/lastItem
func (p NSPopUpButton) LastItem() INSMenuItem {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("lastItem"))
	return NSMenuItemFromID(objc.ID(rv))
}
// The edge of the button on which to display the menu when screen space is
// constrained.
//
// # Discussion
// 
// Possible values include [NSMinXEdge], [NSMinYEdge], [NSMaxXEdge], or
// [NSMaxYEdge]. For pull-down menus, the default behavior is to position the
// menu under the button. The bottom edge corresponds to the value
// [NSMaxYEdge] for flipped views or [NSMinYEdge] for unflipped views. For
// most pop-up menus, the [NSPopUpButton] object attempts to show the selected
// item directly over the button.
//
// [NSMaxXEdge]: https://developer.apple.com/documentation/Foundation/NSRectEdge/NSMaxXEdge
// [NSMaxYEdge]: https://developer.apple.com/documentation/Foundation/NSRectEdge/NSMaxYEdge
// [NSMinXEdge]: https://developer.apple.com/documentation/Foundation/NSRectEdge/NSMinXEdge
// [NSMinYEdge]: https://developer.apple.com/documentation/Foundation/NSRectEdge/NSMinYEdge
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/preferredEdge
func (p NSPopUpButton) PreferredEdge() foundation.NSRectEdge {
	rv := objc.Send[foundation.NSRectEdge](p.ID, objc.Sel("preferredEdge"))
	return foundation.NSRectEdge(rv)
}
func (p NSPopUpButton) SetPreferredEdge(value foundation.NSRectEdge) {
	objc.Send[struct{}](p.ID, objc.Sel("setPreferredEdge:"), value)
}
// When the value of this property is [YES], the selected menu item’s
// `state` is set to [NSControlStateValueOn]. When the value of this property
// is [NO], the menu item’s `state` is not changed. When this property
// changes, the `state` of the currently selected item is updated
// appropriately. This property is ignored for pull-down buttons.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/altersStateOfSelectedItem
func (p NSPopUpButton) AltersStateOfSelectedItem() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("altersStateOfSelectedItem"))
	return rv
}
func (p NSPopUpButton) SetAltersStateOfSelectedItem(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAltersStateOfSelectedItem:"), value)
}
// When `usesItemFromMenu` is [YES], a pull-down button uses the title of the
// first menu item and hides the first menu item. A pop-up button uses the
// title of the currently selected menu. The default value is [YES].
//
// See: https://developer.apple.com/documentation/AppKit/NSPopUpButton/usesItemFromMenu
func (p NSPopUpButton) UsesItemFromMenu() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("usesItemFromMenu"))
	return rv
}
func (p NSPopUpButton) SetUsesItemFromMenu(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setUsesItemFromMenu:"), value)
}

