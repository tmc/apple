// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSToolbar] class.
var (
	_NSToolbarClass     NSToolbarClass
	_NSToolbarClassOnce sync.Once
)

func getNSToolbarClass() NSToolbarClass {
	_NSToolbarClassOnce.Do(func() {
		_NSToolbarClass = NSToolbarClass{class: objc.GetClass("NSToolbar")}
	})
	return _NSToolbarClass
}

// GetNSToolbarClass returns the class object for NSToolbar.
func GetNSToolbarClass() NSToolbarClass {
	return getNSToolbarClass()
}

type NSToolbarClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSToolbarClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSToolbarClass) Alloc() NSToolbar {
	rv := objc.Send[NSToolbar](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that manages the space above your app’s custom content and
// either below or integrated with the window’s title bar.
//
// # Overview
//
// An [NSToolbar] object manages the controls and views that apply to the main
// window’s content area. Toolbars provide convenient access to the commands
// and features people use most often. Toolbars are also user-configurable and
// support the display of an interactive customization palette.
//
// Create and configure your toolbar programmatically or using Interface
// Builder. Add items to the toolbar that correspond to the commands you want
// to feature in your window. Each item has a corresponding [NSToolbarItem]
// object, which you use to make changes. Each toolbar manages a unique set of
// items, but you can synchronize the items and state of multiple toolbars by
// assigning the same value to their [NSToolbar.Identifier] properties.
//
// For more information about how to use toolbars, see [Integrating a Toolbar
// and Touch Bar into Your App].
//
// # Creating an toolbar object
//
//   - [NSToolbar.InitWithIdentifier]: Creates a newly allocated toolbar with the specified identifier.
//
// # Configuring the toolbar contents
//
//   - [NSToolbar.Delegate]: The object you use to customize the toolbar contents and configuration.
//   - [NSToolbar.SetDelegate]
//
// # Getting the toolbar’s identity
//
//   - [NSToolbar.Identifier]: The value you use to identify the toolbar in your app.
//
// # Configuring the toolbar’s behavior
//
//   - [NSToolbar.Visible]: A Boolean value that indicates whether the toolbar is visible.
//   - [NSToolbar.SetVisible]
//   - [NSToolbar.DisplayMode]: A value that indicates whether the toolbar displays items using a name, icon, or combination of elements.
//   - [NSToolbar.SetDisplayMode]
//   - [NSToolbar.ShowsBaselineSeparator]: A Boolean value that indicates whether the toolbar shows the separator between the toolbar and the main window contents.
//   - [NSToolbar.SetShowsBaselineSeparator]
//   - [NSToolbar.AllowsUserCustomization]: A Boolean value that indicates whether users can modify the contents of the toolbar.
//   - [NSToolbar.SetAllowsUserCustomization]
//   - [NSToolbar.AllowsExtensionItems]: A Boolean value that indicates whether the toolbar can add items for Action extensions.
//   - [NSToolbar.SetAllowsExtensionItems]
//
// # Managing items on the toolbar
//
//   - [NSToolbar.Items]: An array containing the toolbar’s current items, in order.
//   - [NSToolbar.VisibleItems]: An array containing the toolbar’s currently visible items.
//   - [NSToolbar.CenteredItemIdentifiers]: The set of custom items to display in the center of the toolbar.
//   - [NSToolbar.SetCenteredItemIdentifiers]
//   - [NSToolbar.SelectedItemIdentifier]: The identifier of the toolbar’s currently selected item.
//   - [NSToolbar.SetSelectedItemIdentifier]
//   - [NSToolbar.InsertItemWithItemIdentifierAtIndex]: Inserts an item into the toolbar at the specified index.
//   - [NSToolbar.RemoveItemAtIndex]: Removes the item at the specified index in the toolbar.
//
// # Autosaving the configuration
//
//   - [NSToolbar.AutosavesConfiguration]: A Boolean value that indicates whether the toolbar autosaves its configuration.
//   - [NSToolbar.SetAutosavesConfiguration]
//   - [NSToolbar.ConfigurationDictionary]: A dictionary containing the current configuration details for the toolbar.
//
// # Displaying the customization palette
//
//   - [NSToolbar.RunCustomizationPalette]: Displays the toolbar’s customization palette and handles any user-initiated customizations.
//   - [NSToolbar.CustomizationPaletteIsRunning]: A Boolean value that indicates whether the toolbar’s customization palette is in use.
//
// # Validating visible items
//
//   - [NSToolbar.ValidateVisibleItems]: Validates the toolbar’s visible items during a window update.
//
// # Deprecated
//
//   - [NSToolbar.CenteredItemIdentifier]: The item to display in the center of the toolbar.
//   - [NSToolbar.SetCenteredItemIdentifier]
//   - [NSToolbar.SizeMode]: The toolbar’s size mode.
//   - [NSToolbar.SetSizeMode]
//
// # Instance Properties
//
//   - [NSToolbar.AllowsDisplayModeCustomization]: Whether or not the user is allowed to change display modes at run time. This functionality is independent of customizing the order of the items themselves. Only disable when the functionality or legibility of your toolbar could not be improved by another display mode. The user’s selection will be persisted using the toolbar’s `identifier` when `autosavesConfiguration` is enabled. The default is YES for apps linked on macOS 15.0 and above.
//   - [NSToolbar.SetAllowsDisplayModeCustomization]
//   - [NSToolbar.ItemIdentifiers]: An array of itemIdentifiers that represent the current items in the toolbar. Setting this property will set the current items in the toolbar by diffing against items that already exist. Use this with great caution if `allowsUserCustomization` is enabled as it will override any customizations the user has made. This property is key value observable.
//   - [NSToolbar.SetItemIdentifiers]
//
// # Instance Methods
//
//   - [NSToolbar.RemoveItemWithItemIdentifier]: Removes the item with matching `itemIdentifier` in the receiving toolbar. If multiple items share the same identifier (as is the case with space items) all matching items will be removed. To remove only a single space item, use `-` instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar
//
// [Integrating a Toolbar and Touch Bar into Your App]: https://developer.apple.com/documentation/AppKit/integrating-a-toolbar-and-touch-bar-into-your-app
type NSToolbar struct {
	objectivec.Object
}

// NSToolbarFromID constructs a [NSToolbar] from an objc.ID.
//
// An object that manages the space above your app’s custom content and
// either below or integrated with the window’s title bar.
func NSToolbarFromID(id objc.ID) NSToolbar {
	return NSToolbar{objectivec.Object{ID: id}}
}

// NOTE: NSToolbar adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSToolbar] class.
//
// # Creating an toolbar object
//
//   - [INSToolbar.InitWithIdentifier]: Creates a newly allocated toolbar with the specified identifier.
//
// # Configuring the toolbar contents
//
//   - [INSToolbar.Delegate]: The object you use to customize the toolbar contents and configuration.
//   - [INSToolbar.SetDelegate]
//
// # Getting the toolbar’s identity
//
//   - [INSToolbar.Identifier]: The value you use to identify the toolbar in your app.
//
// # Configuring the toolbar’s behavior
//
//   - [INSToolbar.Visible]: A Boolean value that indicates whether the toolbar is visible.
//   - [INSToolbar.SetVisible]
//   - [INSToolbar.DisplayMode]: A value that indicates whether the toolbar displays items using a name, icon, or combination of elements.
//   - [INSToolbar.SetDisplayMode]
//   - [INSToolbar.ShowsBaselineSeparator]: A Boolean value that indicates whether the toolbar shows the separator between the toolbar and the main window contents.
//   - [INSToolbar.SetShowsBaselineSeparator]
//   - [INSToolbar.AllowsUserCustomization]: A Boolean value that indicates whether users can modify the contents of the toolbar.
//   - [INSToolbar.SetAllowsUserCustomization]
//   - [INSToolbar.AllowsExtensionItems]: A Boolean value that indicates whether the toolbar can add items for Action extensions.
//   - [INSToolbar.SetAllowsExtensionItems]
//
// # Managing items on the toolbar
//
//   - [INSToolbar.Items]: An array containing the toolbar’s current items, in order.
//   - [INSToolbar.VisibleItems]: An array containing the toolbar’s currently visible items.
//   - [INSToolbar.CenteredItemIdentifiers]: The set of custom items to display in the center of the toolbar.
//   - [INSToolbar.SetCenteredItemIdentifiers]
//   - [INSToolbar.SelectedItemIdentifier]: The identifier of the toolbar’s currently selected item.
//   - [INSToolbar.SetSelectedItemIdentifier]
//   - [INSToolbar.InsertItemWithItemIdentifierAtIndex]: Inserts an item into the toolbar at the specified index.
//   - [INSToolbar.RemoveItemAtIndex]: Removes the item at the specified index in the toolbar.
//
// # Autosaving the configuration
//
//   - [INSToolbar.AutosavesConfiguration]: A Boolean value that indicates whether the toolbar autosaves its configuration.
//   - [INSToolbar.SetAutosavesConfiguration]
//   - [INSToolbar.ConfigurationDictionary]: A dictionary containing the current configuration details for the toolbar.
//
// # Displaying the customization palette
//
//   - [INSToolbar.RunCustomizationPalette]: Displays the toolbar’s customization palette and handles any user-initiated customizations.
//   - [INSToolbar.CustomizationPaletteIsRunning]: A Boolean value that indicates whether the toolbar’s customization palette is in use.
//
// # Validating visible items
//
//   - [INSToolbar.ValidateVisibleItems]: Validates the toolbar’s visible items during a window update.
//
// # Deprecated
//
//   - [INSToolbar.CenteredItemIdentifier]: The item to display in the center of the toolbar.
//   - [INSToolbar.SetCenteredItemIdentifier]
//   - [INSToolbar.SizeMode]: The toolbar’s size mode.
//   - [INSToolbar.SetSizeMode]
//
// # Instance Properties
//
//   - [INSToolbar.AllowsDisplayModeCustomization]: Whether or not the user is allowed to change display modes at run time. This functionality is independent of customizing the order of the items themselves. Only disable when the functionality or legibility of your toolbar could not be improved by another display mode. The user’s selection will be persisted using the toolbar’s `identifier` when `autosavesConfiguration` is enabled. The default is YES for apps linked on macOS 15.0 and above.
//   - [INSToolbar.SetAllowsDisplayModeCustomization]
//   - [INSToolbar.ItemIdentifiers]: An array of itemIdentifiers that represent the current items in the toolbar. Setting this property will set the current items in the toolbar by diffing against items that already exist. Use this with great caution if `allowsUserCustomization` is enabled as it will override any customizations the user has made. This property is key value observable.
//   - [INSToolbar.SetItemIdentifiers]
//
// # Instance Methods
//
//   - [INSToolbar.RemoveItemWithItemIdentifier]: Removes the item with matching `itemIdentifier` in the receiving toolbar. If multiple items share the same identifier (as is the case with space items) all matching items will be removed. To remove only a single space item, use `-` instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar
type INSToolbar interface {
	objectivec.IObject

	// Topic: Creating an toolbar object

	// Creates a newly allocated toolbar with the specified identifier.
	InitWithIdentifier(identifier NSToolbarIdentifier) NSToolbar

	// Topic: Configuring the toolbar contents

	// The object you use to customize the toolbar contents and configuration.
	Delegate() NSToolbarDelegate
	SetDelegate(value NSToolbarDelegate)

	// Topic: Getting the toolbar’s identity

	// The value you use to identify the toolbar in your app.
	Identifier() NSToolbarIdentifier

	// Topic: Configuring the toolbar’s behavior

	// A Boolean value that indicates whether the toolbar is visible.
	Visible() bool
	SetVisible(value bool)
	// A value that indicates whether the toolbar displays items using a name, icon, or combination of elements.
	DisplayMode() NSToolbarDisplayMode
	SetDisplayMode(value NSToolbarDisplayMode)
	// A Boolean value that indicates whether the toolbar shows the separator between the toolbar and the main window contents.
	ShowsBaselineSeparator() bool
	SetShowsBaselineSeparator(value bool)
	// A Boolean value that indicates whether users can modify the contents of the toolbar.
	AllowsUserCustomization() bool
	SetAllowsUserCustomization(value bool)
	// A Boolean value that indicates whether the toolbar can add items for Action extensions.
	AllowsExtensionItems() bool
	SetAllowsExtensionItems(value bool)

	// Topic: Managing items on the toolbar

	// An array containing the toolbar’s current items, in order.
	Items() []NSToolbarItem
	// An array containing the toolbar’s currently visible items.
	VisibleItems() []NSToolbarItem
	// The set of custom items to display in the center of the toolbar.
	CenteredItemIdentifiers() foundation.INSSet
	SetCenteredItemIdentifiers(value foundation.INSSet)
	// The identifier of the toolbar’s currently selected item.
	SelectedItemIdentifier() NSToolbarItemIdentifier
	SetSelectedItemIdentifier(value NSToolbarItemIdentifier)
	// Inserts an item into the toolbar at the specified index.
	InsertItemWithItemIdentifierAtIndex(itemIdentifier NSToolbarItemIdentifier, index int)
	// Removes the item at the specified index in the toolbar.
	RemoveItemAtIndex(index int)

	// Topic: Autosaving the configuration

	// A Boolean value that indicates whether the toolbar autosaves its configuration.
	AutosavesConfiguration() bool
	SetAutosavesConfiguration(value bool)
	// A dictionary containing the current configuration details for the toolbar.
	ConfigurationDictionary() foundation.INSDictionary

	// Topic: Displaying the customization palette

	// Displays the toolbar’s customization palette and handles any user-initiated customizations.
	RunCustomizationPalette(sender objectivec.IObject)
	// A Boolean value that indicates whether the toolbar’s customization palette is in use.
	CustomizationPaletteIsRunning() bool

	// Topic: Validating visible items

	// Validates the toolbar’s visible items during a window update.
	ValidateVisibleItems()

	// Topic: Deprecated

	// The item to display in the center of the toolbar.
	CenteredItemIdentifier() NSToolbarItemIdentifier
	SetCenteredItemIdentifier(value NSToolbarItemIdentifier)
	// The toolbar’s size mode.
	SizeMode() NSToolbarSizeMode
	SetSizeMode(value NSToolbarSizeMode)

	// Topic: Instance Properties

	// Whether or not the user is allowed to change display modes at run time. This functionality is independent of customizing the order of the items themselves. Only disable when the functionality or legibility of your toolbar could not be improved by another display mode. The user’s selection will be persisted using the toolbar’s `identifier` when `autosavesConfiguration` is enabled. The default is YES for apps linked on macOS 15.0 and above.
	AllowsDisplayModeCustomization() bool
	SetAllowsDisplayModeCustomization(value bool)
	// An array of itemIdentifiers that represent the current items in the toolbar. Setting this property will set the current items in the toolbar by diffing against items that already exist. Use this with great caution if `allowsUserCustomization` is enabled as it will override any customizations the user has made. This property is key value observable.
	ItemIdentifiers() []string
	SetItemIdentifiers(value []string)

	// Topic: Instance Methods

	// Removes the item with matching `itemIdentifier` in the receiving toolbar. If multiple items share the same identifier (as is the case with space items) all matching items will be removed. To remove only a single space item, use `-` instead.
	RemoveItemWithItemIdentifier(itemIdentifier NSToolbarItemIdentifier)
}

// Init initializes the instance.
func (t NSToolbar) Init() NSToolbar {
	rv := objc.Send[NSToolbar](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t NSToolbar) Autorelease() NSToolbar {
	rv := objc.Send[NSToolbar](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSToolbar creates a new NSToolbar instance.
func NewNSToolbar() NSToolbar {
	class := getNSToolbarClass()
	rv := objc.Send[NSToolbar](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a newly allocated toolbar with the specified identifier.
//
// identifier: A string used by the class to identify the kind of the toolbar.
//
// # Return Value
//
// The initialized toolbar object.
//
// # Discussion
//
// `identifier` is never seen by users and should not be localized. See the
// [Identifier] property for important information.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/init(identifier:)
func NewToolbarWithIdentifier(identifier NSToolbarIdentifier) NSToolbar {
	instance := getNSToolbarClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIdentifier:"), objc.String(string(identifier)))
	return NSToolbarFromID(rv)
}

// Creates a newly allocated toolbar with the specified identifier.
//
// identifier: A string used by the class to identify the kind of the toolbar.
//
// # Return Value
//
// The initialized toolbar object.
//
// # Discussion
//
// `identifier` is never seen by users and should not be localized. See the
// [Identifier] property for important information.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/init(identifier:)
func (t NSToolbar) InitWithIdentifier(identifier NSToolbarIdentifier) NSToolbar {
	rv := objc.Send[NSToolbar](t.ID, objc.Sel("initWithIdentifier:"), objc.String(string(identifier)))
	return rv
}

// Inserts an item into the toolbar at the specified index.
//
// itemIdentifier: The identifier of the toolbar item to insert.
//
// index: The index at which to insert the item.
//
// # Discussion
//
// Typically, you don’t call this method directly from your code. Instead,
// you specify your toolbar’s allowed items, and the set of default items
// you want to appear. After that, you let the user customize the toolbar.
//
// Any changes you make to the toolbar appear in all [NSToolbar] objects with
// the same [Identifier] value. If a toolbar item with the specified
// identifier isn’t available, the toolbar calls the
// [ToolbarItemForItemIdentifierWillBeInsertedIntoToolbar] method of its
// delegate to get the item. This method does not trigger a call to your
// delegate’s [ToolbarItemIdentifierCanBeInsertedAtIndex] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/insertItem(withItemIdentifier:at:)
func (t NSToolbar) InsertItemWithItemIdentifierAtIndex(itemIdentifier NSToolbarItemIdentifier, index int) {
	objc.Send[objc.ID](t.ID, objc.Sel("insertItemWithItemIdentifier:atIndex:"), objc.String(string(itemIdentifier)), index)
}

// Removes the item at the specified index in the toolbar.
//
// index: The index of the item to remove.
//
// # Discussion
//
// Typically, you don’t call this method directly from your code. Instead,
// you specify your toolbar’s allowed items, and the set of default items
// you want to appear. After that, you let the user customize the toolbar.
//
// Any changes you make to the toolbar appear in all [NSToolbar] objects with
// the same [Identifier] value. This method does not trigger a call to your
// delegate’s [ToolbarItemIdentifierCanBeInsertedAtIndex] method for the
// removal of the item.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/removeItem(at:)
func (t NSToolbar) RemoveItemAtIndex(index int) {
	objc.Send[objc.ID](t.ID, objc.Sel("removeItemAtIndex:"), index)
}

// Displays the toolbar’s customization palette and handles any
// user-initiated customizations.
//
// sender: The control sending the message.
//
// # Discussion
//
// While the customization palette is visible, the toolbar calls methods of
// its delegate to manage configuration changes.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/runCustomizationPalette(_:)
func (t NSToolbar) RunCustomizationPalette(sender objectivec.IObject) {
	objc.Send[objc.ID](t.ID, objc.Sel("runCustomizationPalette:"), sender)
}

// Validates the toolbar’s visible items during a window update.
//
// # Discussion
//
// Typically, you override this method and use it to customize the validation
// process. The default implementation calls the [Validate] method of each
// visible item in the toolbar, including items that have their
// [Autovalidates] property set to false. If you override this method, call
// `super` at some point during your implementation.
//
// The toolbar doesn’t validate its content for some events, including
// [NSLeftMouseDragged], [NSRightMouseDragged], [NSOtherMouseDragged],
// [NSMouseEntered], [NSMouseExited], [NSScrollWheel], [NSCursorUpdate],
// [NSKeyDown]. In addition, the toolbar defers validation for [NSKeyUp] and
// [NSFlagsChanged] events until a pause of 0.85 seconds occurs or the window
// processes an event other than [NSKeyUp] or [NSFlagsChanged]. So a rapid
// sequence of key events doesn’t trigger any validation.
//
// To trigger validation for all toolbars, call the app’s
// [SetWindowsNeedUpdate] method with a value of true.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/validateVisibleItems()
//
// [NSCursorUpdate]: https://developer.apple.com/documentation/AppKit/NSCursorUpdate
// [NSFlagsChanged]: https://developer.apple.com/documentation/AppKit/NSFlagsChanged
// [NSKeyDown]: https://developer.apple.com/documentation/AppKit/NSKeyDown
// [NSKeyUp]: https://developer.apple.com/documentation/AppKit/NSKeyUp
// [NSLeftMouseDragged]: https://developer.apple.com/documentation/AppKit/NSLeftMouseDragged
// [NSMouseEntered]: https://developer.apple.com/documentation/AppKit/NSMouseEntered
// [NSMouseExited]: https://developer.apple.com/documentation/AppKit/NSMouseExited
// [NSOtherMouseDragged]: https://developer.apple.com/documentation/AppKit/NSOtherMouseDragged
// [NSRightMouseDragged]: https://developer.apple.com/documentation/AppKit/NSRightMouseDragged
// [NSScrollWheel]: https://developer.apple.com/documentation/AppKit/NSScrollWheel
func (t NSToolbar) ValidateVisibleItems() {
	objc.Send[objc.ID](t.ID, objc.Sel("validateVisibleItems"))
}

// Removes the item with matching `itemIdentifier` in the receiving toolbar.
// If multiple items share the same identifier (as is the case with space
// items) all matching items will be removed. To remove only a single space
// item, use `-` instead.
//
// # Discussion
//
// Any change made will be propagated immediately to all other toolbars with
// the same identifier.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/removeItem(identifier:)
func (t NSToolbar) RemoveItemWithItemIdentifier(itemIdentifier NSToolbarItemIdentifier) {
	objc.Send[objc.ID](t.ID, objc.Sel("removeItemWithItemIdentifier:"), objc.String(string(itemIdentifier)))
}

// The object you use to customize the toolbar contents and configuration.
//
// # Discussion
//
// Assign an object to this property if you customize the toolbar’s
// behavior. The object you assign to this property must adopt the
// [NSToolbarDelegate] protocol.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/delegate
func (t NSToolbar) Delegate() NSToolbarDelegate {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("delegate"))
	return NSToolbarDelegateObjectFromID(rv)
}
func (t NSToolbar) SetDelegate(value NSToolbarDelegate) {
	objc.Send[struct{}](t.ID, objc.Sel("setDelegate:"), value)
}

// The value you use to identify the toolbar in your app.
//
// # Discussion
//
// Use this property to distinguish toolbars in your app. Multiple toolbars
// can share the same identifier, and you might do so for windows that display
// similar content. When two or more toolbars share an identifier, they
// synchronize their state and display the same set of items.
//
// If the toolbar autosaves its contents, the system associates the
// configuration data with this identifier.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/identifier-swift.property
func (t NSToolbar) Identifier() NSToolbarIdentifier {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("identifier"))
	return NSToolbarIdentifier(foundation.NSStringFromID(rv).String())
}

// A Boolean value that indicates whether the toolbar is visible.
//
// # Discussion
//
// If the value of this property is true, the toolbar is visible; otherwise,
// it’s false. Change the value to hide or show the toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/isVisible
func (t NSToolbar) Visible() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("isVisible"))
	return rv
}
func (t NSToolbar) SetVisible(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setVisible:"), value)
}

// A value that indicates whether the toolbar displays items using a name,
// icon, or combination of elements.
//
// # Discussion
//
// The default value of this property is [NSToolbarDisplayModeDefault]. For a
// list of possible values, see [NSToolbar.DisplayMode].
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/displayMode-swift.property
//
// [NSToolbar.DisplayMode]: https://developer.apple.com/documentation/AppKit/NSToolbar/DisplayMode-swift.enum
func (t NSToolbar) DisplayMode() NSToolbarDisplayMode {
	rv := objc.Send[NSToolbarDisplayMode](t.ID, objc.Sel("displayMode"))
	return NSToolbarDisplayMode(rv)
}
func (t NSToolbar) SetDisplayMode(value NSToolbarDisplayMode) {
	objc.Send[struct{}](t.ID, objc.Sel("setDisplayMode:"), value)
}

// A Boolean value that indicates whether the toolbar shows the separator
// between the toolbar and the main window contents.
//
// # Discussion
//
// If the value of this property is true, the toolbar shows the separator
// between the toolbar and the main window contents. If the value is false,
// the toolbar doesn’t show the separator. The default value is false.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/showsBaselineSeparator
func (t NSToolbar) ShowsBaselineSeparator() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("showsBaselineSeparator"))
	return rv
}
func (t NSToolbar) SetShowsBaselineSeparator(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setShowsBaselineSeparator:"), value)
}

// A Boolean value that indicates whether users can modify the contents of the
// toolbar.
//
// # Discussion
//
// If the value of this property is true, the toolbar enables the Customize
// Toolbar… menu item. If the value is false, the toolbar disables this menu
// item. The Customize Toolbar… menu item lets people change the items on
// the toolbar, rearrange their positions, and change the toolbar’s display
// mode. This attribute does not affect someone’s ability to show or hide
// the toolbar. The default value of this property is false.
//
// You can change the value of this property at any time to change your
// toolbar’s customization behavior. For example, you might prevent toolbar
// customizations while your app processes some other event. If you set this
// property to true, set the [AutosavesConfiguration] property to true to
// persist any customizations.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/allowsUserCustomization
func (t NSToolbar) AllowsUserCustomization() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsUserCustomization"))
	return rv
}
func (t NSToolbar) SetAllowsUserCustomization(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsUserCustomization:"), value)
}

// A Boolean value that indicates whether the toolbar can add items for Action
// extensions.
//
// # Discussion
//
// If the value of this property is true, the toolbar can dynamically create
// toolbar items for Action extensions in the toolbar configuration panel. The
// toolbar can only add an Action extension if its `Info.Plist()` file
// contains the [NSExtensionServiceAllowsToolbarItem] key with the value true.
// The default value of this property is false.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/allowsExtensionItems
//
// [NSExtensionServiceAllowsToolbarItem]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/NSExtension/NSExtensionAttributes/NSExtensionServiceAllowsToolbarItem
func (t NSToolbar) AllowsExtensionItems() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsExtensionItems"))
	return rv
}
func (t NSToolbar) SetAllowsExtensionItems(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsExtensionItems:"), value)
}

// An array containing the toolbar’s current items, in order.
//
// # Discussion
//
// To specify the default order of your toolbar’s items, implement the
// [ToolbarDefaultItemIdentifiers] method in your toolbar delegate object. Use
// other methods of your delegate object to manage the placement of items in
// the toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/items
func (t NSToolbar) Items() []NSToolbarItem {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("items"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSToolbarItem {
		return NSToolbarItemFromID(id)
	})
}

// An array containing the toolbar’s currently visible items.
//
// # Discussion
//
// This property doesn’t contain items in the overflow menu because those
// items aren’t visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/visibleItems
func (t NSToolbar) VisibleItems() []NSToolbarItem {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("visibleItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSToolbarItem {
		return NSToolbarItemFromID(id)
	})
}

// The set of custom items to display in the center of the toolbar.
//
// # Discussion
//
// Set this property to the items you want to appear together in the center of
// the toolbar. Specify the initial order of the items using the
// [ToolbarDefaultItemIdentifiers] method of your toolbar delegate object.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/centeredItemIdentifiers
func (t NSToolbar) CenteredItemIdentifiers() foundation.INSSet {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("centeredItemIdentifiers"))
	return foundation.NSSetFromID(objc.ID(rv))
}
func (t NSToolbar) SetCenteredItemIdentifiers(value foundation.INSSet) {
	objc.Send[struct{}](t.ID, objc.Sel("setCenteredItemIdentifiers:"), value)
}

// The identifier of the toolbar’s currently selected item.
//
// # Discussion
//
// The value of this property is `nil` if the toolbar doesn’t have a
// selected item.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/selectedItemIdentifier
func (t NSToolbar) SelectedItemIdentifier() NSToolbarItemIdentifier {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("selectedItemIdentifier"))
	return NSToolbarItemIdentifier(foundation.NSStringFromID(rv).String())
}
func (t NSToolbar) SetSelectedItemIdentifier(value NSToolbarItemIdentifier) {
	objc.Send[struct{}](t.ID, objc.Sel("setSelectedItemIdentifier:"), objc.String(string(value)))
}

// A Boolean value that indicates whether the toolbar autosaves its
// configuration.
//
// # Discussion
//
// When the value of this property is true, the toolbar automatically writes
// any configuration changes to user defaults. It associates the configuration
// details with the value in its identifier property. If mutliple toolbars
// share the same identifier, they all share the same configuration settings.
// When the value of this property is false, the toolbar doesn’t save
// changes and reverts to the default configuration when the app relaunches.
//
// The default of this property is false.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/autosavesConfiguration
func (t NSToolbar) AutosavesConfiguration() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("autosavesConfiguration"))
	return rv
}
func (t NSToolbar) SetAutosavesConfiguration(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAutosavesConfiguration:"), value)
}

// A dictionary containing the current configuration details for the toolbar.
//
// # Discussion
//
// Use this property to retrieve the toolbar’s configuration details so you
// can save them to disk yourself. The dictionary in this property contains
// the identifiers of the current toolbar items and the values of important
// properties such as [DisplayMode] and [Visible].
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/configuration
func (t NSToolbar) ConfigurationDictionary() foundation.INSDictionary {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("configurationDictionary"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the toolbar’s customization
// palette is in use.
//
// # Discussion
//
// The value of this property is true when the toolbar’s customization
// palette is running; otherwise, it’s false.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/customizationPaletteIsRunning
func (t NSToolbar) CustomizationPaletteIsRunning() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("customizationPaletteIsRunning"))
	return rv
}

// The item to display in the center of the toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/centeredItemIdentifier
func (t NSToolbar) CenteredItemIdentifier() NSToolbarItemIdentifier {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("centeredItemIdentifier"))
	return NSToolbarItemIdentifier(foundation.NSStringFromID(rv).String())
}
func (t NSToolbar) SetCenteredItemIdentifier(value NSToolbarItemIdentifier) {
	objc.Send[struct{}](t.ID, objc.Sel("setCenteredItemIdentifier:"), objc.String(string(value)))
}

// The toolbar’s size mode.
//
// # Discussion
//
// If there’s no icon of the given size for a toolbar item, the toolbar item
// creates one by scaling an icon of another size.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/sizeMode-swift.property
func (t NSToolbar) SizeMode() NSToolbarSizeMode {
	rv := objc.Send[NSToolbarSizeMode](t.ID, objc.Sel("sizeMode"))
	return NSToolbarSizeMode(rv)
}
func (t NSToolbar) SetSizeMode(value NSToolbarSizeMode) {
	objc.Send[struct{}](t.ID, objc.Sel("setSizeMode:"), value)
}

// Whether or not the user is allowed to change display modes at run time.
// This functionality is independent of customizing the order of the items
// themselves. Only disable when the functionality or legibility of your
// toolbar could not be improved by another display mode. The user’s
// selection will be persisted using the toolbar’s `identifier` when
// `autosavesConfiguration` is enabled. The default is YES for apps linked on
// macOS 15.0 and above.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/allowsDisplayModeCustomization
func (t NSToolbar) AllowsDisplayModeCustomization() bool {
	rv := objc.Send[bool](t.ID, objc.Sel("allowsDisplayModeCustomization"))
	return rv
}
func (t NSToolbar) SetAllowsDisplayModeCustomization(value bool) {
	objc.Send[struct{}](t.ID, objc.Sel("setAllowsDisplayModeCustomization:"), value)
}

// An array of itemIdentifiers that represent the current items in the
// toolbar. Setting this property will set the current items in the toolbar by
// diffing against items that already exist. Use this with great caution if
// `allowsUserCustomization` is enabled as it will override any customizations
// the user has made. This property is key value observable.
//
// See: https://developer.apple.com/documentation/AppKit/NSToolbar/itemIdentifiers
func (t NSToolbar) ItemIdentifiers() []string {
	rv := objc.Send[[]objc.ID](t.ID, objc.Sel("itemIdentifiers"))
	return objc.ConvertSliceToStrings(rv)
}
func (t NSToolbar) SetItemIdentifiers(value []string) {
	objc.Send[struct{}](t.ID, objc.Sel("setItemIdentifiers:"), objectivec.StringSliceToNSArray(value))
}
