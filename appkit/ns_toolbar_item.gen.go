// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSToolbarItem] class.
var (
	_NSToolbarItemClass     NSToolbarItemClass
	_NSToolbarItemClassOnce sync.Once
)

func getNSToolbarItemClass() NSToolbarItemClass {
	_NSToolbarItemClassOnce.Do(func() {
		_NSToolbarItemClass = NSToolbarItemClass{class: objc.GetClass("NSToolbarItem")}
	})
	return _NSToolbarItemClass
}

// GetNSToolbarItemClass returns the class object for NSToolbarItem.
func GetNSToolbarItemClass() NSToolbarItemClass {
	return getNSToolbarItemClass()
}

type NSToolbarItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSToolbarItemClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSToolbarItemClass) Alloc() NSToolbarItem {
	rv := objc.Send[NSToolbarItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A single item that appears in a window’s toolbar.
//
// # Overview
// 
// An [NSToolbarItem] object displays an image and text string in the toolbar
// area of a window. You can also create toolbar items that display custom
// views you provide. Toolbar items provide fast access to common commands or
// features in the window. For example, the Finder window uses toolbar items
// to help someone navigate the file system.
// 
// You typically create toolbar items at the same time you create your
// window’s toolbar. The system provides some standard items like spacers
// you can include in your toolbar. It also provides items that display
// standard interfaces like the color panel or font panel. For any custom
// toolbar items you create, provide an action method to call when someone
// clicks the item.
// 
// You can display your toolbar item’s content using a custom view if you
// prefer, rather than an image and text label. If you specify an
// [NSSearchField] object for the view, the system automatically adjusts the
// minimum and maximum size of the search field to the system-standard values.
//
// # Creating a toolbar item
//
//   - [NSToolbarItem.InitWithItemIdentifier]: Creates a toolbar item with the specified identifier.
//
// # Getting the toolbar item’s identity
//
//   - [NSToolbarItem.ItemIdentifier]: The value you use to identify the toolbar item.
//
// # Describing the item
//
//   - [NSToolbarItem.PossibleLabels]: The set of labels that the item might display.
//   - [NSToolbarItem.SetPossibleLabels]
//   - [NSToolbarItem.Label]: The label that appears for this item in the toolbar.
//   - [NSToolbarItem.SetLabel]
//   - [NSToolbarItem.PaletteLabel]: The label that appears when the toolbar item is in the customization palette.
//   - [NSToolbarItem.SetPaletteLabel]
//   - [NSToolbarItem.Title]: The title of the toolbar item.
//   - [NSToolbarItem.SetTitle]
//   - [NSToolbarItem.ToolTip]: The tooltip to display when someone hovers over the item in the toolbar.
//   - [NSToolbarItem.SetToolTip]
//
// # Getting the item’s visual appearance
//
//   - [NSToolbarItem.Image]: The image to display for the toolbar item.
//   - [NSToolbarItem.SetImage]
//   - [NSToolbarItem.BackgroundTintColor]
//   - [NSToolbarItem.SetBackgroundTintColor]
//   - [NSToolbarItem.View]: The custom view you use to draw the toolbar item.
//   - [NSToolbarItem.SetView]
//
// # Performing the item’s action
//
//   - [NSToolbarItem.Target]: The object that defines the action method the toolbar item calls when clicked.
//   - [NSToolbarItem.SetTarget]
//
// # Configuring the item’s menu
//
//   - [NSToolbarItem.MenuFormRepresentation]: The menu item to use when the toolbar item is in the overflow menu.
//   - [NSToolbarItem.SetMenuFormRepresentation]
//
// # Getting the item’s configuration
//
//   - [NSToolbarItem.Visible]: A Boolean value that indicates whether the item is currently visible in the toolbar, and not in the overflow menu.
//   - [NSToolbarItem.Hidden]
//   - [NSToolbarItem.SetHidden]
//   - [NSToolbarItem.Bordered]: A Boolean value that indicates whether the toolbar item has a bordered style.
//   - [NSToolbarItem.SetBordered]
//   - [NSToolbarItem.Navigational]: A Boolean value that indicates whether the item behaves as a navigation item in the toolbar.
//   - [NSToolbarItem.SetNavigational]
//   - [NSToolbarItem.Enabled]: A Boolean value that indicates whether the item is enabled.
//   - [NSToolbarItem.SetEnabled]
//   - [NSToolbarItem.Badge]: A badge that can be attached to an NSToolbarItem. This provides a way to display small visual indicators that can be used to highlight important information, such as unread notifications or status indicators.
//   - [NSToolbarItem.SetBadge]
//   - [NSToolbarItem.Style]: Defines the toolbar item’s appearance. The default style is plain. Prominent style tints the background. If a background tint color is set, it uses it; otherwise, it uses the app’s or system’s accent color. If grouped with other items, it moves to its own to avoid tinting other items’ background.
//   - [NSToolbarItem.SetStyle]
//   - [NSToolbarItem.VisibilityPriority]: The display priority associated with the toolbar item.
//   - [NSToolbarItem.SetVisibilityPriority]
//
// # Getting the parent toolbar
//
//   - [NSToolbarItem.Toolbar]: The toolbar that currently includes the item.
//
// # Validating the item
//
//   - [NSToolbarItem.Autovalidates]: A Boolean value that indicates whether the toolbar automatically validates the item.
//   - [NSToolbarItem.SetAutovalidates]
//   - [NSToolbarItem.Validate]: Validates the toolbar item’s menu and its ability to perfrom its action.
//
// # Deprecated
//
//   - [NSToolbarItem.AllowsDuplicatesInToolbar]: A Boolean value that indicates whether the toolbar item can appear more than once in a toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem
type NSToolbarItem struct {
	objectivec.Object
}

// NSToolbarItemFromID constructs a [NSToolbarItem] from an objc.ID.
//
// A single item that appears in a window’s toolbar.
func NSToolbarItemFromID(id objc.ID) NSToolbarItem {
	return NSToolbarItem{objectivec.Object{ID: id}}
}
// NOTE: NSToolbarItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSToolbarItem] class.
//
// # Creating a toolbar item
//
//   - [INSToolbarItem.InitWithItemIdentifier]: Creates a toolbar item with the specified identifier.
//
// # Getting the toolbar item’s identity
//
//   - [INSToolbarItem.ItemIdentifier]: The value you use to identify the toolbar item.
//
// # Describing the item
//
//   - [INSToolbarItem.PossibleLabels]: The set of labels that the item might display.
//   - [INSToolbarItem.SetPossibleLabels]
//   - [INSToolbarItem.Label]: The label that appears for this item in the toolbar.
//   - [INSToolbarItem.SetLabel]
//   - [INSToolbarItem.PaletteLabel]: The label that appears when the toolbar item is in the customization palette.
//   - [INSToolbarItem.SetPaletteLabel]
//   - [INSToolbarItem.Title]: The title of the toolbar item.
//   - [INSToolbarItem.SetTitle]
//   - [INSToolbarItem.ToolTip]: The tooltip to display when someone hovers over the item in the toolbar.
//   - [INSToolbarItem.SetToolTip]
//
// # Getting the item’s visual appearance
//
//   - [INSToolbarItem.Image]: The image to display for the toolbar item.
//   - [INSToolbarItem.SetImage]
//   - [INSToolbarItem.BackgroundTintColor]
//   - [INSToolbarItem.SetBackgroundTintColor]
//   - [INSToolbarItem.View]: The custom view you use to draw the toolbar item.
//   - [INSToolbarItem.SetView]
//
// # Performing the item’s action
//
//   - [INSToolbarItem.Target]: The object that defines the action method the toolbar item calls when clicked.
//   - [INSToolbarItem.SetTarget]
//
// # Configuring the item’s menu
//
//   - [INSToolbarItem.MenuFormRepresentation]: The menu item to use when the toolbar item is in the overflow menu.
//   - [INSToolbarItem.SetMenuFormRepresentation]
//
// # Getting the item’s configuration
//
//   - [INSToolbarItem.Visible]: A Boolean value that indicates whether the item is currently visible in the toolbar, and not in the overflow menu.
//   - [INSToolbarItem.Hidden]
//   - [INSToolbarItem.SetHidden]
//   - [INSToolbarItem.Bordered]: A Boolean value that indicates whether the toolbar item has a bordered style.
//   - [INSToolbarItem.SetBordered]
//   - [INSToolbarItem.Navigational]: A Boolean value that indicates whether the item behaves as a navigation item in the toolbar.
//   - [INSToolbarItem.SetNavigational]
//   - [INSToolbarItem.Enabled]: A Boolean value that indicates whether the item is enabled.
//   - [INSToolbarItem.SetEnabled]
//   - [INSToolbarItem.Badge]: A badge that can be attached to an NSToolbarItem. This provides a way to display small visual indicators that can be used to highlight important information, such as unread notifications or status indicators.
//   - [INSToolbarItem.SetBadge]
//   - [INSToolbarItem.Style]: Defines the toolbar item’s appearance. The default style is plain. Prominent style tints the background. If a background tint color is set, it uses it; otherwise, it uses the app’s or system’s accent color. If grouped with other items, it moves to its own to avoid tinting other items’ background.
//   - [INSToolbarItem.SetStyle]
//   - [INSToolbarItem.VisibilityPriority]: The display priority associated with the toolbar item.
//   - [INSToolbarItem.SetVisibilityPriority]
//
// # Getting the parent toolbar
//
//   - [INSToolbarItem.Toolbar]: The toolbar that currently includes the item.
//
// # Validating the item
//
//   - [INSToolbarItem.Autovalidates]: A Boolean value that indicates whether the toolbar automatically validates the item.
//   - [INSToolbarItem.SetAutovalidates]
//   - [INSToolbarItem.Validate]: Validates the toolbar item’s menu and its ability to perfrom its action.
//
// # Deprecated
//
//   - [INSToolbarItem.AllowsDuplicatesInToolbar]: A Boolean value that indicates whether the toolbar item can appear more than once in a toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem
type INSToolbarItem interface {
	objectivec.IObject
	NSMenuItemValidation
	NSValidatedUserInterfaceItem

	// Topic: Creating a toolbar item

	// Creates a toolbar item with the specified identifier.
	InitWithItemIdentifier(itemIdentifier NSToolbarItemIdentifier) NSToolbarItem

	// Topic: Getting the toolbar item’s identity

	// The value you use to identify the toolbar item.
	ItemIdentifier() NSToolbarItemIdentifier

	// Topic: Describing the item

	// The set of labels that the item might display.
	PossibleLabels() foundation.INSSet
	SetPossibleLabels(value foundation.INSSet)
	// The label that appears for this item in the toolbar.
	Label() string
	SetLabel(value string)
	// The label that appears when the toolbar item is in the customization palette.
	PaletteLabel() string
	SetPaletteLabel(value string)
	// The title of the toolbar item.
	Title() string
	SetTitle(value string)
	// The tooltip to display when someone hovers over the item in the toolbar.
	ToolTip() string
	SetToolTip(value string)

	// Topic: Getting the item’s visual appearance

	// The image to display for the toolbar item.
	Image() objectivec.Object
	SetImage(value objectivec.Object)
	BackgroundTintColor() objectivec.Object
	SetBackgroundTintColor(value objectivec.Object)
	// The custom view you use to draw the toolbar item.
	View() INSView
	SetView(value INSView)

	// Topic: Performing the item’s action

	// The object that defines the action method the toolbar item calls when clicked.
	Target() objectivec.IObject
	SetTarget(value objectivec.IObject)

	// Topic: Configuring the item’s menu

	// The menu item to use when the toolbar item is in the overflow menu.
	MenuFormRepresentation() INSMenuItem
	SetMenuFormRepresentation(value INSMenuItem)

	// Topic: Getting the item’s configuration

	// A Boolean value that indicates whether the item is currently visible in the toolbar, and not in the overflow menu.
	Visible() bool
	Hidden() bool
	SetHidden(value bool)
	// A Boolean value that indicates whether the toolbar item has a bordered style.
	Bordered() bool
	SetBordered(value bool)
	// A Boolean value that indicates whether the item behaves as a navigation item in the toolbar.
	Navigational() bool
	SetNavigational(value bool)
	// A Boolean value that indicates whether the item is enabled.
	Enabled() bool
	SetEnabled(value bool)
	// A badge that can be attached to an NSToolbarItem. This provides a way to display small visual indicators that can be used to highlight important information, such as unread notifications or status indicators.
	Badge() INSItemBadge
	SetBadge(value INSItemBadge)
	// Defines the toolbar item’s appearance. The default style is plain. Prominent style tints the background. If a background tint color is set, it uses it; otherwise, it uses the app’s or system’s accent color. If grouped with other items, it moves to its own to avoid tinting other items’ background.
	Style() NSToolbarItemStyle
	SetStyle(value NSToolbarItemStyle)
	// The display priority associated with the toolbar item.
	VisibilityPriority() NSToolbarItemVisibilityPriority
	SetVisibilityPriority(value NSToolbarItemVisibilityPriority)

	// Topic: Getting the parent toolbar

	// The toolbar that currently includes the item.
	Toolbar() INSToolbar

	// Topic: Validating the item

	// A Boolean value that indicates whether the toolbar automatically validates the item.
	Autovalidates() bool
	SetAutovalidates(value bool)
	// Validates the toolbar item’s menu and its ability to perfrom its action.
	Validate()

	// Topic: Deprecated

	// A Boolean value that indicates whether the toolbar item can appear more than once in a toolbar.
	AllowsDuplicatesInToolbar() bool
}

// Init initializes the instance.
func (t NSToolbarItem) Init() NSToolbarItem {
	rv := objc.Send[NSToolbarItem](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSToolbarItem) Autorelease() NSToolbarItem {
	rv := objc.Send[NSToolbarItem](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSToolbarItem creates a new NSToolbarItem instance.
func NewNSToolbarItem() NSToolbarItem {
	class := getNSToolbarItemClass()
	rv := objc.Send[NSToolbarItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a toolbar item with the specified identifier.
//
// itemIdentifier: The identifier for the toolbar item. You use this value to identify the
// item within your app, so you don’t need to localize it. For example, your
// toolbar delegate uses this value to identify the specific toolbar item.
//
// # Return Value
// 
// A new toolbar item.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/init(itemIdentifier:)
func NewToolbarItemWithItemIdentifier(itemIdentifier NSToolbarItemIdentifier) NSToolbarItem {
	instance := getNSToolbarItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithItemIdentifier:"), objc.String(string(itemIdentifier)))
	return NSToolbarItemFromID(rv)
}

// Creates a toolbar item with the specified identifier.
//
// itemIdentifier: The identifier for the toolbar item. You use this value to identify the
// item within your app, so you don’t need to localize it. For example, your
// toolbar delegate uses this value to identify the specific toolbar item.
//
// # Return Value
// 
// A new toolbar item.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/init(itemIdentifier:)
func (t NSToolbarItem) InitWithItemIdentifier(itemIdentifier NSToolbarItemIdentifier) NSToolbarItem {
	rv := objc.Send[NSToolbarItem](t.ID, objc.Sel("initWithItemIdentifier:"), objc.String(string(itemIdentifier)))
	return rv
}
// Validates the toolbar item’s menu and its ability to perfrom its action.
//
// # Discussion
// 
// Typically, you don’t call this method directly. When automatic validation
// is enabled, the toolbar calls this method to validate the item. For
// standard toolbar items — that is, items without a custom view — the
// validation process checks whether the item can perform its associated
// action successfully and enables or disables the item accordingly. The
// process also validates the associated menu item. When someone switches to
// another app or window, the automatic validation process disables the item
// automatically.
// 
// If the toolbar item has a custom view, subclass [NSToolbarItem] and
// override this method to perform the validation yourself. After you validate
// your custom toolbar item, update the [Enabled] property. You don’t need
// to call `super` in your implementation.
// 
// If you disable automatic validation, toolbar items remain enabled and
// clickable, including when someone switches to another app or window.
// However, you can still call this method manually to validate the toolbar
// item.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/validate()
func (t NSToolbarItem) Validate() {
	objc.Send[objc.ID](t.ID, objc.Sel("validate"))
}
// Implemented to override the default action of enabling or disabling a
// specific menu item.
//
// menuItem: An [NSMenuItem] object that represents the menu item.
//
// # Return Value
// 
// [true] to enable `menuItem`, [false] to disable it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
func (t NSToolbarItem) ValidateMenuItem(menuItem INSMenuItem) bool {
	rv := objc.Send[bool](t.ID, objc.Sel("validateMenuItem:"), menuItem)
	return rv
}

// The value you use to identify the toolbar item.
//
// # Discussion
// 
// Use this property to distinguish one toolbar item from another on the same
// toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/itemIdentifier
func (t NSToolbarItem) ItemIdentifier() NSToolbarItemIdentifier {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("itemIdentifier"))
	return NSToolbarItemIdentifier(foundation.NSStringFromID(rv).String())
}
// The set of labels that the item might display.
//
// # Discussion
// 
// Use this property to specify all of the labels you might possibly use for
// the toolbar item. Specify all strings in the current locale. To ensure
// there’s space for the longest label, the item sizes itself using the
// strings you provide.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/possibleLabels
func (t NSToolbarItem) PossibleLabels() foundation.INSSet {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("possibleLabels"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (t NSToolbarItem) SetPossibleLabels(value foundation.INSSet) {
	objc.Send[struct{}](t.ID, objc.Sel("setPossibleLabels:"), value)
}
// The label that appears for this item in the toolbar.
//
// # Discussion
// 
// For a discussion on labels, see [Setting a Toolbar Item’s
// Representation].
//
// [Setting a Toolbar Item’s Representation]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Toolbars/Tasks/SettingTBItemRep.html#//apple_ref/doc/uid/20000722
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/label
func (t NSToolbarItem) Label() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (t NSToolbarItem) SetLabel(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setLabel:"), objc.String(value))
}
// The label that appears when the toolbar item is in the customization
// palette.
//
// # Discussion
// 
// If you support toolbar customizations, you must provide palette labels for
// your items. In most cases, you can apply the same value to this property
// and the [Label] property. However, you might use this property to offer a
// more descriptive string, or to provide a label string when the [Label]
// property contains an empty string.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/paletteLabel
func (t NSToolbarItem) PaletteLabel() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("paletteLabel"))
	return foundation.NSStringFromID(rv).String()
}
func (t NSToolbarItem) SetPaletteLabel(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setPaletteLabel:"), objc.String(value))
}
// The title of the toolbar item.
//
// # Discussion
// 
// If you assign a custom view to the toolbar item, modifying this property
// updates the [Title] property of the view if one exists. If the toolbar item
// contains a button, modifying this property updates the button title. If the
// item doesn’t contain a custom view, the toolbar item manages the content
// directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/title
func (t NSToolbarItem) Title() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
func (t NSToolbarItem) SetTitle(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setTitle:"), objc.String(value))
}
// The tooltip to display when someone hovers over the item in the toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/toolTip
func (t NSToolbarItem) ToolTip() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("toolTip"))
	return foundation.NSStringFromID(rv).String()
}
func (t NSToolbarItem) SetToolTip(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setToolTip:"), objc.String(value))
}
// The image to display for the toolbar item.
//
// # Discussion
// 
// If you assign a custom view to the toolbar item, modifying this property
// updates the `image` property of the view, if one exists. If the item
// doesn’t contain a custom view, the toolbar item manages the image content
// directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/image
func (t NSToolbarItem) Image() objectivec.Object {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("image"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (t NSToolbarItem) SetImage(value objectivec.Object) {
	objc.Send[struct{}](t.ID, objc.Sel("setImage:"), value)
}
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/backgroundTintColor
func (t NSToolbarItem) BackgroundTintColor() objectivec.Object {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("backgroundTintColor"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (t NSToolbarItem) SetBackgroundTintColor(value objectivec.Object) {
	objc.Send[struct{}](t.ID, objc.Sel("setBackgroundTintColor:"), value)
}
// The custom view you use to draw the toolbar item.
//
// # Discussion
// 
// Many properties of [NSToolbarItem] automatically forward changes to the
// associated custom [NSView] object, if the view has a property or accessor
// method with a matching name.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/view
func (t NSToolbarItem) View() INSView {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("view"))
	return NSViewFromID(objc.ID(rv))
}
func (t NSToolbarItem) SetView(value INSView) {
	objc.Send[struct{}](t.ID, objc.Sel("setView:"), value)
}
// The object that defines the action method the toolbar item calls when
// clicked.
//
// # Discussion
// 
// If you set this property to `nil`, the toolbar item attempts to execute its
// action method on the first responder. If the first responder doesn’t
// implement the action method, it forwards the request up the responder
// chain.
// 
// If you assign a custom view to the toolbar item, modifying this property
// updates the `target` property of the view, if one exists. If the item
// doesn’t contain a custom view, the toolbar item manages the target object
// directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/target
func (t NSToolbarItem) Target() objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("target"))
	return objectivec.Object{ID: rv}
}
func (t NSToolbarItem) SetTarget(value objectivec.IObject) {
	objc.Send[struct{}](t.ID, objc.Sel("setTarget:"), value)
}
// The action method to call when someone clicks on the toolbar item.
//
// # Discussion
// 
// If you assign a custom view to the toolbar item, modifying this property
// updates the `action` property of the view or calls the `` method, if one of
// them exists. If the item doesn’t contain a custom view, the toolbar item
// manages the action directly.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/action
func (t NSToolbarItem) Action() objc.SEL {
	rv := objc.Send[objc.SEL](t.ID, objc.Sel("action"))
	return rv
}
func (t NSToolbarItem) SetAction(value objc.SEL) {
	objc.Send[struct{}](t.ID, objc.Sel("setAction:"), value)
}
// The menu item to use when the toolbar item is in the overflow menu.
//
// # Discussion
// 
// The toolbar provides an initial default menu form representation that uses
// the toolbar item’s label as the menu item’s title. You can customize
// this menu item by changing the title or adding a submenu. When the toolbar
// is in text only mode, this menu item provides the text for the toolbar
// item. If the menu item in this property has a submenu and is visible,
// clicking the toolbar item displays that submenu. If the toolbar item
// isn’t visible because it’s in the overflow menu, the menu item and
// submenu appear there.
// 
// For a discussion on menu forms, see [Setting a Toolbar Item’s
// Representation].
//
// [Setting a Toolbar Item’s Representation]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Toolbars/Tasks/SettingTBItemRep.html#//apple_ref/doc/uid/20000722
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/menuFormRepresentation
func (t NSToolbarItem) MenuFormRepresentation() INSMenuItem {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("menuFormRepresentation"))
	return NSMenuItemFromID(objc.ID(rv))
}
func (t NSToolbarItem) SetMenuFormRepresentation(value INSMenuItem) {
	objc.Send[struct{}](t.ID, objc.Sel("setMenuFormRepresentation:"), value)
}
// A Boolean value that indicates whether the item is currently visible in the
// toolbar, and not in the overflow menu.
//
// # Discussion
// 
// The value of this property is [true] when the item is visible in the
// toolbar, and [false] when it isn’t in the toolbar or is present in the
// toolbar’s overflow menu. This property is key-value observing (KVO)
// compliant.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isVisible
func (t NSToolbarItem) Visible() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isVisible"))
	return rv
}
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isHidden
func (t NSToolbarItem) Hidden() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isHidden"))
	return rv
}
func (t NSToolbarItem) SetHidden(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setHidden:"), value)
}
// A Boolean value that indicates whether the toolbar item has a bordered
// style.
//
// # Discussion
// 
// If your toolbar item displays a custom view, modifying this property
// applies the image to the view’s [Bordered] property, if one exists. The
// default value of this property is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isBordered
func (t NSToolbarItem) Bordered() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isBordered"))
	return rv
}
func (t NSToolbarItem) SetBordered(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setBordered:"), value)
}
// A Boolean value that indicates whether the item behaves as a navigation
// item in the toolbar.
//
// # Discussion
// 
// Mark a toolbar item as navigation if you use it to navigate around your
// content. When you set this property to [true], the system can position
// navigation items outside of the normal list of items in the toolbar. For
// example, the back and forward buttons in Finder windows are navigational,
// and the system positions them at the leading edge of the window’s title
// area. Specify the initial order of the items using the
// [ToolbarDefaultItemIdentifiers] method of the toolbar delegate object.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isNavigational
func (t NSToolbarItem) Navigational() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isNavigational"))
	return rv
}
func (t NSToolbarItem) SetNavigational(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setNavigational:"), value)
}
// A Boolean value that indicates whether the item is enabled.
//
// # Discussion
// 
// If the value of this property is [true], the item is enabled. If the
// [Autovalidates] property is true, changing the value of this property has
// no effect. Instead, the validation process enables and disables the toolbar
// item as appropriate.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/isEnabled
func (t NSToolbarItem) Enabled() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isEnabled"))
	return rv
}
func (t NSToolbarItem) SetEnabled(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setEnabled:"), value)
}
// A badge that can be attached to an NSToolbarItem. This provides a way to
// display small visual indicators that can be used to highlight important
// information, such as unread notifications or status indicators.
//
// See: https://developer.apple.com/documentation/appkit/nstoolbaritem/badge-17r3r
func (t NSToolbarItem) Badge() INSItemBadge {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("badge"))
	return NSItemBadgeFromID(objc.ID(rv))
}
func (t NSToolbarItem) SetBadge(value INSItemBadge) {
	objc.Send[struct{}](t.ID, objc.Sel("setBadge:"), value)
}
// Defines the toolbar item’s appearance. The default style is plain.
// Prominent style tints the background. If a background tint color is set, it
// uses it; otherwise, it uses the app’s or system’s accent color. If
// grouped with other items, it moves to its own to avoid tinting other
// items’ background.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/style-swift.property
func (t NSToolbarItem) Style() NSToolbarItemStyle {
	rv := objc.Send[NSToolbarItemStyle](t.ID, objc.Sel("style"))
	return NSToolbarItemStyle(rv)
}
func (t NSToolbarItem) SetStyle(value NSToolbarItemStyle) {
	objc.Send[struct{}](t.ID, objc.Sel("setStyle:"), value)
}
// The display priority associated with the toolbar item.
//
// # Discussion
// 
// The default value of this property is [standard]. Assign a higher priority
// to give preference to the toolbar item when space is limited.
// 
// When a toolbar doesn’t have enough space to fit all of its items, it
// pushes lower-priority items to the overflow menu first. When two or more
// items have the same priority, the toolbar removes them one at a time
// starting from the trailing edge.
//
// [standard]: https://developer.apple.com/documentation/AppKit/NSToolbarItem/VisibilityPriority-swift.struct/standard
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/visibilityPriority-swift.property
func (t NSToolbarItem) VisibilityPriority() NSToolbarItemVisibilityPriority {
	rv := objc.Send[NSToolbarItemVisibilityPriority](t.ID, objc.Sel("visibilityPriority"))
	return NSToolbarItemVisibilityPriority(rv)
}
func (t NSToolbarItem) SetVisibilityPriority(value NSToolbarItemVisibilityPriority) {
	objc.Send[struct{}](t.ID, objc.Sel("setVisibilityPriority:"), value)
}
// An integer tag you can use to identify the toolbar item.
//
// # Discussion
// 
// The toolbar doesn’t use this value. You can use it for your own custom
// purposes.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/tag
func (t NSToolbarItem) Tag() int {
	rv := objc.Send[int](t.ID, objc.Sel("tag"))
	return rv
}
func (t NSToolbarItem) SetTag(value int) {
	objc.Send[struct{}](t.ID, objc.Sel("setTag:"), value)
}
// The toolbar that currently includes the item.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/toolbar
func (t NSToolbarItem) Toolbar() INSToolbar {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("toolbar"))
	return NSToolbarFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the toolbar automatically validates
// the item.
//
// # Discussion
// 
// If the value of this property is [true], the toolbar automatically
// validates the item; otherwise, it doesn’t validate the item
// automatically. The default value of this property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/autovalidates
func (t NSToolbarItem) Autovalidates() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("autovalidates"))
	return rv
}
func (t NSToolbarItem) SetAutovalidates(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAutovalidates:"), value)
}
// A Boolean value that indicates whether the toolbar item can appear more
// than once in a toolbar.
//
// # Discussion
// 
// If the value in this property is [true], the toolbar allows someone to drag
// more than one copy of the toolbar item from the customization palette. If
// the value of this property is [false], the toolbar prevents someone from
// dragging more than one copy of the item from the customization palette.
// 
// By default, if an item with the same identifier is already in the toolbar,
// dragging it in again will effectively move it to the new position.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbarItem/allowsDuplicatesInToolbar
func (t NSToolbarItem) AllowsDuplicatesInToolbar() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsDuplicatesInToolbar"))
	return rv
}

			// Protocol methods for NSMenuItemValidation
			

			// Protocol methods for NSValidatedUserInterfaceItem
			

