// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMenuItem] class.
var (
	_NSMenuItemClass     NSMenuItemClass
	_NSMenuItemClassOnce sync.Once
)

func getNSMenuItemClass() NSMenuItemClass {
	_NSMenuItemClassOnce.Do(func() {
		_NSMenuItemClass = NSMenuItemClass{class: objc.GetClass("NSMenuItem")}
	})
	return _NSMenuItemClass
}

// GetNSMenuItemClass returns the class object for NSMenuItem.
func GetNSMenuItemClass() NSMenuItemClass {
	return getNSMenuItemClass()
}

type NSMenuItemClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSMenuItemClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMenuItemClass) Alloc() NSMenuItem {
	rv := objc.Send[NSMenuItem](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A command item in an app menu.
//
// # Overview
//
// The [NSMenuItem] class includes some private functionality needed to
// maintain binary compatibility with other components of Cocoa. Because of
// this fact, you can’t replace the [NSMenuItem] class with a different
// class, but you can subclass it if necessary.
//
// # Creating a menu item
//
//   - [NSMenuItem.InitWithTitleActionKeyEquivalent]: Returns an initialized instance of [NSMenuItem].
//   - [NSMenuItem.InitWithCoder]
//
// # Enabling a menu item
//
//   - [NSMenuItem.IsEnabled]: A Boolean value that indicates whether the menu item is enabled.
//   - [NSMenuItem.SetEnabled]
//
// # Managing hidden status
//
//   - [NSMenuItem.IsHidden]: A Boolean value that indicates whether the menu item is hidden.
//   - [NSMenuItem.SetHidden]
//   - [NSMenuItem.IsHiddenOrHasHiddenAncestor]: A Boolean value that indicates whether the menu item or any of its superitems is hidden.
//
// # Managing the target and action
//
//   - [NSMenuItem.Target]: The menu item’s target.
//   - [NSMenuItem.SetTarget]
//
// # Managing the title
//
//   - [NSMenuItem.Title]: The menu item’s title.
//   - [NSMenuItem.SetTitle]
//   - [NSMenuItem.AttributedTitle]: A custom string for a menu item.
//   - [NSMenuItem.SetAttributedTitle]
//
// # Managing the state
//
//   - [NSMenuItem.State]: The state of the menu item.
//   - [NSMenuItem.SetState]
//
// # Managing the image
//
//   - [NSMenuItem.Image]: The menu item’s image.
//   - [NSMenuItem.SetImage]
//   - [NSMenuItem.OnStateImage]: The image of the menu item that indicates an “on” state.
//   - [NSMenuItem.SetOnStateImage]
//   - [NSMenuItem.OffStateImage]: The image of the menu item that indicates an “off” state.
//   - [NSMenuItem.SetOffStateImage]
//   - [NSMenuItem.MixedStateImage]: The image of the menu item that indicates a “mixed” state, that is, a state neither “on” nor “off.”
//   - [NSMenuItem.SetMixedStateImage]
//
// # Managing the badge
//
//   - [NSMenuItem.Badge]: A badge used to provide additional quantitative information specific to the menu item, such as the number of available updates.
//   - [NSMenuItem.SetBadge]
//
// # Managing the section header
//
//   - [NSMenuItem.IsSectionHeader]: A Boolean value indicating whether the menu item is a section header.
//
// # Managing submenus
//
//   - [NSMenuItem.Submenu]: The submenu of the menu item.
//   - [NSMenuItem.SetSubmenu]
//   - [NSMenuItem.HasSubmenu]: A Boolean value that indicates whether the menu item has a submenu.
//   - [NSMenuItem.ParentItem]: The menu item whose submenu contains the receiver.
//
// # Managing the separator item
//
//   - [NSMenuItem.IsSeparatorItem]: A Boolean value indicating whether the menu item is a separator item.
//
// # Managing the owning menu
//
//   - [NSMenuItem.Menu]: The menu item’s menu.
//   - [NSMenuItem.SetMenu]
//
// # Managing key equivalents
//
//   - [NSMenuItem.KeyEquivalent]: The menu item’s unmodified key equivalent.
//   - [NSMenuItem.SetKeyEquivalent]
//   - [NSMenuItem.KeyEquivalentModifierMask]: The menu item’s keyboard equivalent modifiers.
//   - [NSMenuItem.SetKeyEquivalentModifierMask]
//
// # Managing user key equivalents
//
//   - [NSMenuItem.UserKeyEquivalent]: The user-assigned key equivalent for the menu item.
//   - [NSMenuItem.AllowsAutomaticKeyEquivalentLocalization]: A Boolean value that determines whether the system automatically remaps the keyboard shortcut to support localized keyboards.
//   - [NSMenuItem.SetAllowsAutomaticKeyEquivalentLocalization]
//   - [NSMenuItem.AllowsAutomaticKeyEquivalentMirroring]: A Boolean value that determines whether the system automatically swaps input strings for some keyboard shortcuts when the interface direction changes.
//   - [NSMenuItem.SetAllowsAutomaticKeyEquivalentMirroring]
//   - [NSMenuItem.AllowsKeyEquivalentWhenHidden]
//   - [NSMenuItem.SetAllowsKeyEquivalentWhenHidden]
//
// # Managing alternates
//
//   - [NSMenuItem.IsAlternate]: A Boolean value that marks the menu item as an alternate to the previous menu item.
//   - [NSMenuItem.SetAlternate]
//
// # Managing indentation levels
//
//   - [NSMenuItem.IndentationLevel]: The menu item indentation level for the menu item.
//   - [NSMenuItem.SetIndentationLevel]
//
// # Managing tool tips
//
//   - [NSMenuItem.ToolTip]: A help tag for the menu item.
//   - [NSMenuItem.SetToolTip]
//
// # Representing an object
//
//   - [NSMenuItem.RepresentedObject]: The object represented by the menu item.
//   - [NSMenuItem.SetRepresentedObject]
//
// # Managing the view
//
//   - [NSMenuItem.View]: The content view for the menu item.
//   - [NSMenuItem.SetView]
//
// # Getting highlighted status
//
//   - [NSMenuItem.IsHighlighted]: A Boolean value that indicates whether the menu item should be drawn highlighted.
//
// # Instance Properties
//
//   - [NSMenuItem.Subtitle]
//   - [NSMenuItem.SetSubtitle]
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem
type NSMenuItem struct {
	objectivec.Object
}

// NSMenuItemFromID constructs a [NSMenuItem] from an objc.ID.
//
// A command item in an app menu.
func NSMenuItemFromID(id objc.ID) NSMenuItem {
	return NSMenuItem{objectivec.Object{ID: id}}
}

// NOTE: NSMenuItem adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSMenuItem] class.
//
// # Creating a menu item
//
//   - [INSMenuItem.InitWithTitleActionKeyEquivalent]: Returns an initialized instance of [NSMenuItem].
//   - [INSMenuItem.InitWithCoder]
//
// # Enabling a menu item
//
//   - [INSMenuItem.IsEnabled]: A Boolean value that indicates whether the menu item is enabled.
//   - [INSMenuItem.SetEnabled]
//
// # Managing hidden status
//
//   - [INSMenuItem.IsHidden]: A Boolean value that indicates whether the menu item is hidden.
//   - [INSMenuItem.SetHidden]
//   - [INSMenuItem.IsHiddenOrHasHiddenAncestor]: A Boolean value that indicates whether the menu item or any of its superitems is hidden.
//
// # Managing the target and action
//
//   - [INSMenuItem.Target]: The menu item’s target.
//   - [INSMenuItem.SetTarget]
//
// # Managing the title
//
//   - [INSMenuItem.Title]: The menu item’s title.
//   - [INSMenuItem.SetTitle]
//   - [INSMenuItem.AttributedTitle]: A custom string for a menu item.
//   - [INSMenuItem.SetAttributedTitle]
//
// # Managing the state
//
//   - [INSMenuItem.State]: The state of the menu item.
//   - [INSMenuItem.SetState]
//
// # Managing the image
//
//   - [INSMenuItem.Image]: The menu item’s image.
//   - [INSMenuItem.SetImage]
//   - [INSMenuItem.OnStateImage]: The image of the menu item that indicates an “on” state.
//   - [INSMenuItem.SetOnStateImage]
//   - [INSMenuItem.OffStateImage]: The image of the menu item that indicates an “off” state.
//   - [INSMenuItem.SetOffStateImage]
//   - [INSMenuItem.MixedStateImage]: The image of the menu item that indicates a “mixed” state, that is, a state neither “on” nor “off.”
//   - [INSMenuItem.SetMixedStateImage]
//
// # Managing the badge
//
//   - [INSMenuItem.Badge]: A badge used to provide additional quantitative information specific to the menu item, such as the number of available updates.
//   - [INSMenuItem.SetBadge]
//
// # Managing the section header
//
//   - [INSMenuItem.IsSectionHeader]: A Boolean value indicating whether the menu item is a section header.
//
// # Managing submenus
//
//   - [INSMenuItem.Submenu]: The submenu of the menu item.
//   - [INSMenuItem.SetSubmenu]
//   - [INSMenuItem.HasSubmenu]: A Boolean value that indicates whether the menu item has a submenu.
//   - [INSMenuItem.ParentItem]: The menu item whose submenu contains the receiver.
//
// # Managing the separator item
//
//   - [INSMenuItem.IsSeparatorItem]: A Boolean value indicating whether the menu item is a separator item.
//
// # Managing the owning menu
//
//   - [INSMenuItem.Menu]: The menu item’s menu.
//   - [INSMenuItem.SetMenu]
//
// # Managing key equivalents
//
//   - [INSMenuItem.KeyEquivalent]: The menu item’s unmodified key equivalent.
//   - [INSMenuItem.SetKeyEquivalent]
//   - [INSMenuItem.KeyEquivalentModifierMask]: The menu item’s keyboard equivalent modifiers.
//   - [INSMenuItem.SetKeyEquivalentModifierMask]
//
// # Managing user key equivalents
//
//   - [INSMenuItem.UserKeyEquivalent]: The user-assigned key equivalent for the menu item.
//   - [INSMenuItem.AllowsAutomaticKeyEquivalentLocalization]: A Boolean value that determines whether the system automatically remaps the keyboard shortcut to support localized keyboards.
//   - [INSMenuItem.SetAllowsAutomaticKeyEquivalentLocalization]
//   - [INSMenuItem.AllowsAutomaticKeyEquivalentMirroring]: A Boolean value that determines whether the system automatically swaps input strings for some keyboard shortcuts when the interface direction changes.
//   - [INSMenuItem.SetAllowsAutomaticKeyEquivalentMirroring]
//   - [INSMenuItem.AllowsKeyEquivalentWhenHidden]
//   - [INSMenuItem.SetAllowsKeyEquivalentWhenHidden]
//
// # Managing alternates
//
//   - [INSMenuItem.IsAlternate]: A Boolean value that marks the menu item as an alternate to the previous menu item.
//   - [INSMenuItem.SetAlternate]
//
// # Managing indentation levels
//
//   - [INSMenuItem.IndentationLevel]: The menu item indentation level for the menu item.
//   - [INSMenuItem.SetIndentationLevel]
//
// # Managing tool tips
//
//   - [INSMenuItem.ToolTip]: A help tag for the menu item.
//   - [INSMenuItem.SetToolTip]
//
// # Representing an object
//
//   - [INSMenuItem.RepresentedObject]: The object represented by the menu item.
//   - [INSMenuItem.SetRepresentedObject]
//
// # Managing the view
//
//   - [INSMenuItem.View]: The content view for the menu item.
//   - [INSMenuItem.SetView]
//
// # Getting highlighted status
//
//   - [INSMenuItem.IsHighlighted]: A Boolean value that indicates whether the menu item should be drawn highlighted.
//
// # Instance Properties
//
//   - [INSMenuItem.Subtitle]
//   - [INSMenuItem.SetSubtitle]
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem
type INSMenuItem interface {
	objectivec.IObject
	NSValidatedUserInterfaceItem

	// Topic: Creating a menu item

	// Returns an initialized instance of [NSMenuItem].
	InitWithTitleActionKeyEquivalent(string_ string, selector objc.SEL, charCode string) NSMenuItem
	InitWithCoder(coder foundation.INSCoder) NSMenuItem

	// Topic: Enabling a menu item

	// A Boolean value that indicates whether the menu item is enabled.
	IsEnabled() bool
	SetEnabled(value bool)

	// Topic: Managing hidden status

	// A Boolean value that indicates whether the menu item is hidden.
	IsHidden() bool
	SetHidden(value bool)
	// A Boolean value that indicates whether the menu item or any of its superitems is hidden.
	IsHiddenOrHasHiddenAncestor() bool

	// Topic: Managing the target and action

	// The menu item’s target.
	Target() objectivec.IObject
	SetTarget(value objectivec.IObject)

	// Topic: Managing the title

	// The menu item’s title.
	Title() string
	SetTitle(value string)
	// A custom string for a menu item.
	AttributedTitle() foundation.NSAttributedString
	SetAttributedTitle(value foundation.NSAttributedString)

	// Topic: Managing the state

	// The state of the menu item.
	State() NSControlStateValue
	SetState(value NSControlStateValue)

	// Topic: Managing the image

	// The menu item’s image.
	Image() INSImage
	SetImage(value INSImage)
	// The image of the menu item that indicates an “on” state.
	OnStateImage() INSImage
	SetOnStateImage(value INSImage)
	// The image of the menu item that indicates an “off” state.
	OffStateImage() INSImage
	SetOffStateImage(value INSImage)
	// The image of the menu item that indicates a “mixed” state, that is, a state neither “on” nor “off.”
	MixedStateImage() INSImage
	SetMixedStateImage(value INSImage)

	// Topic: Managing the badge

	// A badge used to provide additional quantitative information specific to the menu item, such as the number of available updates.
	Badge() INSMenuItemBadge
	SetBadge(value INSMenuItemBadge)

	// Topic: Managing the section header

	// A Boolean value indicating whether the menu item is a section header.
	IsSectionHeader() bool

	// Topic: Managing submenus

	// The submenu of the menu item.
	Submenu() INSMenu
	SetSubmenu(value INSMenu)
	// A Boolean value that indicates whether the menu item has a submenu.
	HasSubmenu() bool
	// The menu item whose submenu contains the receiver.
	ParentItem() INSMenuItem

	// Topic: Managing the separator item

	// A Boolean value indicating whether the menu item is a separator item.
	IsSeparatorItem() bool

	// Topic: Managing the owning menu

	// The menu item’s menu.
	Menu() INSMenu
	SetMenu(value INSMenu)

	// Topic: Managing key equivalents

	// The menu item’s unmodified key equivalent.
	KeyEquivalent() string
	SetKeyEquivalent(value string)
	// The menu item’s keyboard equivalent modifiers.
	KeyEquivalentModifierMask() NSEventModifierFlags
	SetKeyEquivalentModifierMask(value NSEventModifierFlags)

	// Topic: Managing user key equivalents

	// The user-assigned key equivalent for the menu item.
	UserKeyEquivalent() string
	// A Boolean value that determines whether the system automatically remaps the keyboard shortcut to support localized keyboards.
	AllowsAutomaticKeyEquivalentLocalization() bool
	SetAllowsAutomaticKeyEquivalentLocalization(value bool)
	// A Boolean value that determines whether the system automatically swaps input strings for some keyboard shortcuts when the interface direction changes.
	AllowsAutomaticKeyEquivalentMirroring() bool
	SetAllowsAutomaticKeyEquivalentMirroring(value bool)
	AllowsKeyEquivalentWhenHidden() bool
	SetAllowsKeyEquivalentWhenHidden(value bool)

	// Topic: Managing alternates

	// A Boolean value that marks the menu item as an alternate to the previous menu item.
	IsAlternate() bool
	SetAlternate(value bool)

	// Topic: Managing indentation levels

	// The menu item indentation level for the menu item.
	IndentationLevel() int
	SetIndentationLevel(value int)

	// Topic: Managing tool tips

	// A help tag for the menu item.
	ToolTip() string
	SetToolTip(value string)

	// Topic: Representing an object

	// The object represented by the menu item.
	RepresentedObject() objectivec.IObject
	SetRepresentedObject(value objectivec.IObject)

	// Topic: Managing the view

	// The content view for the menu item.
	View() INSView
	SetView(value INSView)

	// Topic: Getting highlighted status

	// A Boolean value that indicates whether the menu item should be drawn highlighted.
	IsHighlighted() bool

	// Topic: Instance Properties

	Subtitle() string
	SetSubtitle(value string)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (m NSMenuItem) Init() NSMenuItem {
	rv := objc.Send[NSMenuItem](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMenuItem) Autorelease() NSMenuItem {
	rv := objc.Send[NSMenuItem](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMenuItem creates a new NSMenuItem instance.
func NewNSMenuItem() NSMenuItem {
	class := getNSMenuItemClass()
	rv := objc.Send[NSMenuItem](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/init(coder:)
func NewMenuItemWithCoder(coder foundation.INSCoder) NSMenuItem {
	instance := getNSMenuItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSMenuItemFromID(rv)
}

// Returns an initialized instance of [NSMenuItem].
//
// string: The title of the menu item. This value must not be `nil` (if there is no
// title, specify an empty [NSString]).
//
// selector: The action selector to be associated with the menu item. This value must be
// a valid selector or [NULL].
//
// charCode: A string representing a keyboard key to be used as the key equivalent. This
// value must not be `nil` (if there is no key equivalent, specify an empty
// [NSString]).
//
// # Return Value
//
// An instance of [NSMenuItem].
//
// # Discussion
//
// For instances of the [NSMenuItem] class, the default initial state is
// [NSOffState], the default on-state image is a check mark, and the default
// mixed-state image is a dash.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/init(title:action:keyEquivalent:)
func NewMenuItemWithTitleActionKeyEquivalent(string_ string, selector objc.SEL, charCode string) NSMenuItem {
	instance := getNSMenuItemClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTitle:action:keyEquivalent:"), objc.String(string_), selector, objc.String(charCode))
	return NSMenuItemFromID(rv)
}

// Returns an initialized instance of [NSMenuItem].
//
// string: The title of the menu item. This value must not be `nil` (if there is no
// title, specify an empty [NSString]).
//
// selector: The action selector to be associated with the menu item. This value must be
// a valid selector or [NULL].
//
// charCode: A string representing a keyboard key to be used as the key equivalent. This
// value must not be `nil` (if there is no key equivalent, specify an empty
// [NSString]).
//
// # Return Value
//
// An instance of [NSMenuItem].
//
// # Discussion
//
// For instances of the [NSMenuItem] class, the default initial state is
// [NSOffState], the default on-state image is a check mark, and the default
// mixed-state image is a dash.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/init(title:action:keyEquivalent:)
func (m NSMenuItem) InitWithTitleActionKeyEquivalent(string_ string, selector objc.SEL, charCode string) NSMenuItem {
	rv := objc.Send[NSMenuItem](m.ID, objc.Sel("initWithTitle:action:keyEquivalent:"), objc.String(string_), selector, objc.String(charCode))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/init(coder:)
func (m NSMenuItem) InitWithCoder(coder foundation.INSCoder) NSMenuItem {
	rv := objc.Send[NSMenuItem](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// A string that identifies the user interface item.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemIdentification/identifier
func (m NSMenuItem) Identifier() NSUserInterfaceItemIdentifier {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("identifier"))
	return NSUserInterfaceItemIdentifier(foundation.NSStringFromID(rv).String())
}
func (m NSMenuItem) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Returns a menu item that is used to separate logical groups of menu
// commands.
//
// # Return Value
//
// A menu item that is used to separate logical groups of menu commands.
//
// # Discussion
//
// This menu item is disabled. The default separator item is blank space.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/separator()
func (_NSMenuItemClass NSMenuItemClass) SeparatorItem() NSMenuItem {
	rv := objc.Send[objc.ID](objc.ID(_NSMenuItemClass.class), objc.Sel("separatorItem"))
	return NSMenuItemFromID(rv)
}

// Returns a menu item representing a section header for a logical grouping of
// menu commands.
//
// title: The title string to display on the section header.
//
// # Return Value
//
// A menu item representing a section header for a logical grouping of menu
// commands.
//
// # Discussion
//
// Use section headers to provide context to a group of menu items. Items
// created using this method are non-interactive and don’t perform actions.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/sectionHeaderWithTitle:
func (_NSMenuItemClass NSMenuItemClass) SectionHeaderWithTitle(title string) NSMenuItem {
	rv := objc.Send[objc.ID](objc.ID(_NSMenuItemClass.class), objc.Sel("sectionHeaderWithTitle:"), objc.String(title))
	return NSMenuItemFromID(rv)
}

// A Boolean value that indicates whether the menu item is enabled.
//
// # Discussion
//
// This property has no effect unless the menu in which the item will be added
// or is already a part of has been sent `NO`. If a menu item is disabled, its
// keyboard equivalent is also disabled. See the [NSMenuValidation] informal
// protocol specification for cautions regarding this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/isEnabled
func (m NSMenuItem) IsEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isEnabled"))
	return rv
}
func (m NSMenuItem) SetEnabled(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setEnabled:"), value)
}

// A Boolean value that indicates whether the menu item is hidden.
//
// # Discussion
//
// Hidden menu items (or items with a hidden superitem) do not appear in a
// menu and do not participate in command key matching.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/isHidden
func (m NSMenuItem) IsHidden() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isHidden"))
	return rv
}
func (m NSMenuItem) SetHidden(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setHidden:"), value)
}

// A Boolean value that indicates whether the menu item or any of its
// superitems is hidden.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/isHiddenOrHasHiddenAncestor
func (m NSMenuItem) IsHiddenOrHasHiddenAncestor() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isHiddenOrHasHiddenAncestor"))
	return rv
}

// The menu item’s target.
//
// # Discussion
//
// To ensure that a menu item’s target can receive commands while a modal
// dialog is open, the target object should return true in
// [NSPanel.WorksWhenModal].
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/target
func (m NSMenuItem) Target() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("target"))
	return objectivec.Object{ID: rv}
}
func (m NSMenuItem) SetTarget(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setTarget:"), value)
}

// The menu item’s action-method selector.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/action
func (m NSMenuItem) Action() objectivec.SEL {
	rv := objc.Send[objc.SEL](m.ID, objc.Sel("action"))
	return objectivec.SEL(rv)
}
func (m NSMenuItem) SetAction(value objectivec.SEL) {
	objc.Send[struct{}](m.ID, objc.Sel("setAction:"), value)
}

// The menu item’s title.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/title
func (m NSMenuItem) Title() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("title"))
	return foundation.NSStringFromID(rv).String()
}
func (m NSMenuItem) SetTitle(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setTitle:"), objc.String(value))
}

// A custom string for a menu item.
//
// # Discussion
//
// The attributed string is not archived in the old nib format.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/attributedTitle
func (m NSMenuItem) AttributedTitle() foundation.NSAttributedString {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("attributedTitle"))
	return foundation.NSAttributedStringFromID(objc.ID(rv))
}
func (m NSMenuItem) SetAttributedTitle(value foundation.NSAttributedString) {
	objc.Send[struct{}](m.ID, objc.Sel("setAttributedTitle:"), value)
}

// The menu item’s tag.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/tag
func (m NSMenuItem) Tag() int {
	rv := objc.Send[int](m.ID, objc.Sel("tag"))
	return rv
}
func (m NSMenuItem) SetTag(value int) {
	objc.Send[struct{}](m.ID, objc.Sel("setTag:"), value)
}

// The state of the menu item.
//
// # Discussion
//
// The image associated with the new state is displayed to the left of the
// menu item.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/state
func (m NSMenuItem) State() NSControlStateValue {
	rv := objc.Send[NSControlStateValue](m.ID, objc.Sel("state"))
	return NSControlStateValue(rv)
}
func (m NSMenuItem) SetState(value NSControlStateValue) {
	objc.Send[struct{}](m.ID, objc.Sel("setState:"), value)
}

// The menu item’s image.
//
// # Discussion
//
// The menu item’s image is not affected by changes in its state.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/image
func (m NSMenuItem) Image() INSImage {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("image"))
	return NSImageFromID(objc.ID(rv))
}
func (m NSMenuItem) SetImage(value INSImage) {
	objc.Send[struct{}](m.ID, objc.Sel("setImage:"), value)
}

// The image of the menu item that indicates an “on” state.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/onStateImage
func (m NSMenuItem) OnStateImage() INSImage {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("onStateImage"))
	return NSImageFromID(objc.ID(rv))
}
func (m NSMenuItem) SetOnStateImage(value INSImage) {
	objc.Send[struct{}](m.ID, objc.Sel("setOnStateImage:"), value)
}

// The image of the menu item that indicates an “off” state.
//
// # Discussion
//
// By default there is no off-state image.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/offStateImage
func (m NSMenuItem) OffStateImage() INSImage {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("offStateImage"))
	return NSImageFromID(objc.ID(rv))
}
func (m NSMenuItem) SetOffStateImage(value INSImage) {
	objc.Send[struct{}](m.ID, objc.Sel("setOffStateImage:"), value)
}

// The image of the menu item that indicates a “mixed” state, that is, a
// state neither “on” nor “off.”
//
// # Discussion
//
// A mixed state is useful for indicating a mix of “off” and “on”
// attribute values in a group of selected objects, such as a selection of
// text containing boldface and plain (non-boldface) words. By default this is
// a horizontal line.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/mixedStateImage
func (m NSMenuItem) MixedStateImage() INSImage {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mixedStateImage"))
	return NSImageFromID(objc.ID(rv))
}
func (m NSMenuItem) SetMixedStateImage(value INSImage) {
	objc.Send[struct{}](m.ID, objc.Sel("setMixedStateImage:"), value)
}

// A badge used to provide additional quantitative information specific to the
// menu item, such as the number of available updates.
//
// # Discussion
//
// The default value of this property is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/badge
func (m NSMenuItem) Badge() INSMenuItemBadge {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("badge"))
	return NSMenuItemBadgeFromID(objc.ID(rv))
}
func (m NSMenuItem) SetBadge(value INSMenuItemBadge) {
	objc.Send[struct{}](m.ID, objc.Sel("setBadge:"), value)
}

// A Boolean value indicating whether the menu item is a section header.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/isSectionHeader
func (m NSMenuItem) IsSectionHeader() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isSectionHeader"))
	return rv
}

// The submenu of the menu item.
//
// # Discussion
//
// The default implementation of the [NSMenuItem] class raises an exception if
// `aSubmenu` already has a supermenu.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/submenu
func (m NSMenuItem) Submenu() INSMenu {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("submenu"))
	return NSMenuFromID(objc.ID(rv))
}
func (m NSMenuItem) SetSubmenu(value INSMenu) {
	objc.Send[struct{}](m.ID, objc.Sel("setSubmenu:"), value)
}

// A Boolean value that indicates whether the menu item has a submenu.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/hasSubmenu
func (m NSMenuItem) HasSubmenu() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasSubmenu"))
	return rv
}

// The menu item whose submenu contains the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/parent
func (m NSMenuItem) ParentItem() INSMenuItem {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parentItem"))
	return NSMenuItemFromID(objc.ID(rv))
}

// A Boolean value indicating whether the menu item is a separator item.
//
// # Discussion
//
// This menu item is disabled. The default separator item is blank space.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/isSeparatorItem
func (m NSMenuItem) IsSeparatorItem() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isSeparatorItem"))
	return rv
}

// The menu item’s menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/menu
func (m NSMenuItem) Menu() INSMenu {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("menu"))
	return NSMenuFromID(objc.ID(rv))
}
func (m NSMenuItem) SetMenu(value INSMenu) {
	objc.Send[struct{}](m.ID, objc.Sel("setMenu:"), value)
}

// The menu item’s unmodified key equivalent.
//
// # Discussion
//
// If you want to specify the Backspace key as the key equivalent for a menu
// item, use a single character string with [NSBackspaceCharacter] (defined in
// `NSText.H()` as `0x08`) and for the Forward Delete key, use
// [NSDeleteCharacter] (defined in `NSText.H()` as `0x7F`). Note that these
// are not the same characters you get from an [NSEvent] key-down event when
// pressing those keys.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/keyEquivalent
func (m NSMenuItem) KeyEquivalent() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("keyEquivalent"))
	return foundation.NSStringFromID(rv).String()
}
func (m NSMenuItem) SetKeyEquivalent(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setKeyEquivalent:"), objc.String(value))
}

// The menu item’s keyboard equivalent modifiers.
//
// # Discussion
//
// [NSShiftKeyMask] is a valid modifier for any key equivalent in `mask`. This
// allows you to specify key-equivalents such as Command-Shift-1 that are
// consistent across all keyboards. However, with a few exceptions (such as
// the German “ß” character), a lowercase character with [NSShiftKeyMask]
// is interpreted the same as the uppercase character without that mask. For
// example, Command-Shift-c and Command-C are considered to be identical key
// equivalents.
//
// See the [NSEvent] class specification for more information about modifier
// mask values.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/keyEquivalentModifierMask
func (m NSMenuItem) KeyEquivalentModifierMask() NSEventModifierFlags {
	rv := objc.Send[NSEventModifierFlags](m.ID, objc.Sel("keyEquivalentModifierMask"))
	return NSEventModifierFlags(rv)
}
func (m NSMenuItem) SetKeyEquivalentModifierMask(value NSEventModifierFlags) {
	objc.Send[struct{}](m.ID, objc.Sel("setKeyEquivalentModifierMask:"), value)
}

// The user-assigned key equivalent for the menu item.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/userKeyEquivalent
func (m NSMenuItem) UserKeyEquivalent() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("userKeyEquivalent"))
	return foundation.NSStringFromID(rv).String()
}

// A Boolean value that determines whether the system automatically remaps the
// keyboard shortcut to support localized keyboards.
//
// # Discussion
//
// A keyboard shortcut you specify in one language might be difficult or
// impossible to reproduce on a keyboard with a different character set or
// layout. Localized keyboards sometimes rearrange punctuation marks or
// replace them altogether to make room for a language’s required
// characters. The new locations of those keys might make it difficult to use
// your menu item’s current shortcut. To ensure your shortcuts are always
// usable, the system can automatically remap shortcuts, as needed, to
// accommodate the connected keyboard.
//
// When the value of this property is true, the system automatically remaps
// this menu item’s shortcut when that shortcut is unreachable on the
// current keyboard. The system doesn’t remap shortcuts when the input keys
// have identical positions on both keyboards, or when the shortcut is still
// easily reachable on the current keyboard. The remapping is transparent to
// your app.
//
// If you already localize your app’s shortcuts for different languages, or
// if you permit someone to customize your app’s shortcuts, you can set this
// property to false to disable the automatic remapping behavior. When you set
// this property to false, the system doesn’t change the shortcut for your
// menu items. Instead, you are responsible for making any required changes to
// support localized keyboards. Setting this property to false also disables
// the automatic mirroring of shortcuts, as described by the
// [NSMenuItem.AllowsAutomaticKeyEquivalentMirroring] property.
//
// The default value of this property is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/allowsAutomaticKeyEquivalentLocalization
func (m NSMenuItem) AllowsAutomaticKeyEquivalentLocalization() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("allowsAutomaticKeyEquivalentLocalization"))
	return rv
}
func (m NSMenuItem) SetAllowsAutomaticKeyEquivalentLocalization(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setAllowsAutomaticKeyEquivalentLocalization:"), value)
}

// A Boolean value that determines whether the system automatically swaps
// input strings for some keyboard shortcuts when the interface direction
// changes.
//
// # Discussion
//
// When a menu item represents a direction-related action, it’s common to
// specify an input string that conveys that direction. For example, Finder
// uses Command-[ to go back to the previous page, and Command-] to go forward
// to the next page. Because directions are different in left-to-right and
// right-to-left interfaces, this property lets the system swap some input
// strings to match the current language direction.
//
// When the value of this property is true, macOS 12 and later automatically
// swaps input strings that contain brackets `[]`, braces `{}`, parenthesis
// `()`, angle brackets “, or arrow keys when the interface directionality
// changes. This behavior eliminates the need for you to create different menu
// items for left-to-right and right-to-left interfaces. Set this property to
// false if you already change this item’s shortcut to support both
// left-to-right and right-to-left interfaces. You might also set it to false
// to keep the same shortcut regardless of the interface’s directionality.
//
// The default value of this property is true. However, if you set the
// [allowsAutomaticLocalization] property to false, the system disables this
// feature regardless of the property’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/allowsAutomaticKeyEquivalentMirroring
//
// [allowsAutomaticLocalization]: https://developer.apple.com/documentation/UIKit/UIKeyCommand/allowsAutomaticLocalization
func (m NSMenuItem) AllowsAutomaticKeyEquivalentMirroring() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("allowsAutomaticKeyEquivalentMirroring"))
	return rv
}
func (m NSMenuItem) SetAllowsAutomaticKeyEquivalentMirroring(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setAllowsAutomaticKeyEquivalentMirroring:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/allowsKeyEquivalentWhenHidden
func (m NSMenuItem) AllowsKeyEquivalentWhenHidden() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("allowsKeyEquivalentWhenHidden"))
	return rv
}
func (m NSMenuItem) SetAllowsKeyEquivalentWhenHidden(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setAllowsKeyEquivalentWhenHidden:"), value)
}

// A Boolean value that marks the menu item as an alternate to the previous
// menu item.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/isAlternate
func (m NSMenuItem) IsAlternate() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isAlternate"))
	return rv
}
func (m NSMenuItem) SetAlternate(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setAlternate:"), value)
}

// The menu item indentation level for the menu item.
//
// # Discussion
//
// The `indentationLevel` value is archived.
//
// Value is from 0 to 15. The default indentation level is 0.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/indentationLevel
func (m NSMenuItem) IndentationLevel() int {
	rv := objc.Send[int](m.ID, objc.Sel("indentationLevel"))
	return rv
}
func (m NSMenuItem) SetIndentationLevel(value int) {
	objc.Send[struct{}](m.ID, objc.Sel("setIndentationLevel:"), value)
}

// A help tag for the menu item.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/toolTip
func (m NSMenuItem) ToolTip() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("toolTip"))
	return foundation.NSStringFromID(rv).String()
}
func (m NSMenuItem) SetToolTip(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setToolTip:"), objc.String(value))
}

// The object represented by the menu item.
//
// # Discussion
//
// By setting a represented object for a menu item, you make an association
// between the menu item and that object. The represented object functions as
// a more specific form of tag that allows you to associate any object, not
// just an arbitrary integer, with the items in a menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/representedObject
func (m NSMenuItem) RepresentedObject() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("representedObject"))
	return objectivec.Object{ID: rv}
}
func (m NSMenuItem) SetRepresentedObject(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setRepresentedObject:"), value)
}

// The content view for the menu item.
//
// # Discussion
//
// A menu item with a view does not draw its title, state, font, or other
// standard drawing attributes, and assigns drawing responsibility entirely to
// the view. Keyboard equivalents and type-select continue to use the key
// equivalent and title as normal. For more details, see [Application Menu and
// Pop-up List Programming Topics]
//
// By default, a menu item has a `nil` view.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/view
//
// [Application Menu and Pop-up List Programming Topics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/MenuList/MenuList.html#//apple_ref/doc/uid/10000032i
func (m NSMenuItem) View() INSView {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("view"))
	return NSViewFromID(objc.ID(rv))
}
func (m NSMenuItem) SetView(value INSView) {
	objc.Send[struct{}](m.ID, objc.Sel("setView:"), value)
}

// A Boolean value that indicates whether the menu item should be drawn
// highlighted.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/isHighlighted
func (m NSMenuItem) IsHighlighted() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isHighlighted"))
	return rv
}

// # Discussion
//
// Used to specify a standard subtitle for the menu item.
//
// The subtitle is displayed below the standard title.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/subtitle
func (m NSMenuItem) Subtitle() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("subtitle"))
	return foundation.NSStringFromID(rv).String()
}
func (m NSMenuItem) SetSubtitle(value string) {
	objc.Send[struct{}](m.ID, objc.Sel("setSubtitle:"), objc.String(value))
}

// Returns a Boolean value that indicates whether menu items conform to user
// preferences for key equivalents.
//
// # Return Value
//
// true if menu items conform to user preferences for key equivalents;
// otherwise, false.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/usesUserKeyEquivalents
func (_NSMenuItemClass NSMenuItemClass) UsesUserKeyEquivalents() bool {
	rv := objc.Send[bool](objc.ID(_NSMenuItemClass.class), objc.Sel("usesUserKeyEquivalents"))
	return rv
}
func (_NSMenuItemClass NSMenuItemClass) SetUsesUserKeyEquivalents(value bool) {
	objc.Send[struct{}](objc.ID(_NSMenuItemClass.class), objc.Sel("setUsesUserKeyEquivalents:"), value)
}

// An array of standard menu items related to Writing Tools. Each call to this
// method returns an array of newly allocated instances of NSMenuItem.
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem/writingToolsItems
func (_NSMenuItemClass NSMenuItemClass) WritingToolsItems() []NSMenuItem {
	rv := objc.Send[[]objc.ID](objc.ID(_NSMenuItemClass.class), objc.Sel("writingToolsItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSMenuItem {
		return NSMenuItemFromID(id)
	})
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
func (o NSMenuItem) AccessibilityFrame() corefoundation.CGRect {
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
func (o NSMenuItem) AccessibilityParent() objectivec.IObject {
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
func (o NSMenuItem) AccessibilityIdentifier() string {
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
func (o NSMenuItem) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
}

// Protocol methods for NSAccessibilityProtocol

// Returns a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityElement()
func (o NSMenuItem) IsAccessibilityElement() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityElement"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEnabled()
func (o NSMenuItem) IsAccessibilityEnabled() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEnabled"))
	return rv
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
func (o NSMenuItem) IsAccessibilitySelectorAllowed(selector objc.SEL) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelectorAllowed:"), selector)
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityProtectedContent()
func (o NSMenuItem) IsAccessibilityProtectedContent() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityProtectedContent"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// is currently in a selected state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelected()
func (o NSMenuItem) IsAccessibilitySelected() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelected"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// must have content for successful submission of a form.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityRequired()
func (o NSMenuItem) IsAccessibilityRequired() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityRequired"))
	return rv
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
func (o NSMenuItem) AccessibilityStringForRange(range_ foundation.NSRange) string {
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
func (o NSMenuItem) AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString {
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
func (o NSMenuItem) AccessibilityRTFForRange(range_ foundation.NSRange) foundation.NSData {
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
func (o NSMenuItem) AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect {
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
func (o NSMenuItem) AccessibilityLineForIndex(index int) int {
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
func (o NSMenuItem) AccessibilityRangeForIndex(index int) foundation.NSRange {
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
func (o NSMenuItem) AccessibilityStyleRangeForIndex(index int) foundation.NSRange {
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
func (o NSMenuItem) AccessibilityRangeForLine(line int) foundation.NSRange {
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
func (o NSMenuItem) AccessibilityRangeForPosition(point corefoundation.CGPoint) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForPosition:"), point)
	return rv
}

// Returns the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityAlternateUIVisible()
func (o NSMenuItem) IsAccessibilityAlternateUIVisible() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
	return rv
}

// Returns a Boolean value that determines whether the window is the app’s
// main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMain()
func (o NSMenuItem) IsAccessibilityMain() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMain"))
	return rv
}

// Returns the Boolean value that determines whether the window is in a
// minimized state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMinimized()
func (o NSMenuItem) IsAccessibilityMinimized() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMinimized"))
	return rv
}

// Returns a Boolean value that determines whether the window is modal.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityModal()
func (o NSMenuItem) IsAccessibilityModal() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityModal"))
	return rv
}

// Returns a Boolean value that determines whether the app is the frontmost
// app.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityFrontmost()
func (o NSMenuItem) IsAccessibilityFrontmost() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFrontmost"))
	return rv
}

// Returns a Boolean value that determines whether the app is in a hidden
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityHidden()
func (o NSMenuItem) IsAccessibilityHidden() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityHidden"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility
// element’s grid is in row major order or in column major order.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityOrderedByRow()
func (o NSMenuItem) IsAccessibilityOrderedByRow() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityOrderedByRow"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// is in an expanded state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityExpanded()
func (o NSMenuItem) IsAccessibilityExpanded() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityExpanded"))
	return rv
}

// Returns a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityDisclosed()
func (o NSMenuItem) IsAccessibilityDisclosed() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityDisclosed"))
	return rv
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
func (o NSMenuItem) AccessibilityCellForColumnRow(column int, row int) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCellForColumn:row:"), column, row)
	return objectivec.Object{ID: rv}
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
func (o NSMenuItem) AccessibilityLayoutPointForScreenPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
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
func (o NSMenuItem) AccessibilityLayoutSizeForScreenSize(size corefoundation.CGSize) corefoundation.CGSize {
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
func (o NSMenuItem) AccessibilityScreenPointForLayoutPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
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
func (o NSMenuItem) AccessibilityScreenSizeForLayoutSize(size corefoundation.CGSize) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("accessibilityScreenSizeForLayoutSize:"), size)
	return rv
}

// Returns a Boolean value that indicates whether the accessibility element is
// in an edited state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEdited()
func (o NSMenuItem) IsAccessibilityEdited() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEdited"))
	return rv
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
func (o NSMenuItem) AccessibilityPerformCancel() bool {
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
func (o NSMenuItem) AccessibilityPerformConfirm() bool {
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
func (o NSMenuItem) AccessibilityPerformPick() bool {
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
func (o NSMenuItem) AccessibilityPerformPress() bool {
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
func (o NSMenuItem) AccessibilityPerformShowAlternateUI() bool {
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
func (o NSMenuItem) AccessibilityPerformShowDefaultUI() bool {
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
func (o NSMenuItem) AccessibilityPerformShowMenu() bool {
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
func (o NSMenuItem) AccessibilityPerformRaise() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformRaise"))
	return rv
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
func (o NSMenuItem) AccessibilityPerformIncrement() bool {
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
func (o NSMenuItem) AccessibilityPerformDecrement() bool {
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
func (o NSMenuItem) AccessibilityPerformDelete() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformDelete"))
	return rv
}

// The activation point for the user interface element.
//
// # Discussion
//
// The activation point in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityActivationPoint
func (o NSMenuItem) AccessibilityActivationPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityActivationPoint"))
	return corefoundation.CGPoint(rv)
}

func (o NSMenuItem) SetAccessibilityActivationPoint(value corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityActivationPoint:"), value)
}

// The allowed values for the slider accessibility element.
//
// # Discussion
//
// Use this property if the slider can be set only to predefined values (for
// example, if the slider’s level indicator automatically snaps to the
// closest integer values between 0 and 100).
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityAllowedValues
func (o NSMenuItem) AccessibilityAllowedValues() []foundation.NSNumber {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityAllowedValues"))
	result := make([]foundation.NSNumber, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = foundation.NSNumberFromID(id)
	}
	return result
}

func (o NSMenuItem) SetAccessibilityAllowedValues(value []foundation.NSNumber) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAllowedValues:"), objectivec.IObjectSliceToNSArray(value))
}

// A Boolean value that determines whether the accessibility element’s
// alternative UI is currently visible.
//
// # Discussion
//
// Use this property for elements that present an alternative UI—for
// example, when the pointer hovers over an interface element for a few
// seconds.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityAlternateUIVisible
func (o NSMenuItem) AccessibilityAlternateUIVisible() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
	return bool(rv)
}

func (o NSMenuItem) SetAccessibilityAlternateUIVisible(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAlternateUIVisible:"), value)
}

// The child accessibility element with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityApplicationFocusedUIElement
func (o NSMenuItem) AccessibilityApplicationFocusedUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityApplicationFocusedUIElement"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityApplicationFocusedUIElement(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityApplicationFocusedUIElement:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityAttributedUserInputLabels
func (o NSMenuItem) AccessibilityAttributedUserInputLabels() []foundation.NSAttributedString {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityAttributedUserInputLabels"))
	result := make([]foundation.NSAttributedString, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = foundation.NSAttributedStringFromID(id)
	}
	return result
}

func (o NSMenuItem) SetAccessibilityAttributedUserInputLabels(value []foundation.NSAttributedString) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAttributedUserInputLabels:"), objectivec.IObjectSliceToNSArray(value))
}

// The child accessibility element that represents the window’s cancel
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCancelButton
func (o NSMenuItem) AccessibilityCancelButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCancelButton"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityCancelButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCancelButton:"), value)
}

// The child accessibility elements in the accessibility hierarchy.
//
// # Discussion
//
// This property contains references to child elements in the accessibility
// hierarchy. If you create an [NSView] subclass, you don’t typically need
// to set this value. The system automatically populates the
// `accessibilityChildren` property with descendants in the view hierarchy
// that are also in the accessibility hierarchy. If you use an
// [NSAccessibilityElement] subclass to represent an interface element that is
// not backed by a view, you can either set the `accessibilityChildren`
// property or you can call the
// [NSAccessibilityElement.AccessibilityAddChildElement] convenience method.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityChildren
func (o NSMenuItem) AccessibilityChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityChildren"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilityChildren(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildren:"), value)
}

// An array of child accessibility elements in order for linear navigation.
//
// # Discussion
//
// The array should match all elements found in [accessibilityChildren],
// rearranged in an easily navigable order.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityChildrenInNavigationOrder
//
// [accessibilityChildren]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityChildren
func (o NSMenuItem) AccessibilityChildrenInNavigationOrder() []objectivec.IObject {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityChildrenInNavigationOrder"))
	result := make([]objectivec.IObject, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = objectivec.Object{ID: id}
	}
	return result
}

func (o NSMenuItem) SetAccessibilityChildrenInNavigationOrder(value []objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildrenInNavigationOrder:"), objectivec.IObjectSliceToNSArray(value))
}

// The clear button for the search field.
//
// # Discussion
//
// Use this property on a search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityClearButton
func (o NSMenuItem) AccessibilityClearButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityClearButton"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityClearButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityClearButton:"), value)
}

// The child accessibility element that represents the window’s close
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCloseButton
func (o NSMenuItem) AccessibilityCloseButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCloseButton"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityCloseButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCloseButton:"), value)
}

// The number of columns in the accessibility element’s grid.
//
// # Discussion
//
// Use this property for UI elements that present a grid of child elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumnCount
func (o NSMenuItem) AccessibilityColumnCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityColumnCount"))
	return int(rv)
}

func (o NSMenuItem) SetAccessibilityColumnCount(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnCount:"), value)
}

// The column header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumnHeaderUIElements
func (o NSMenuItem) AccessibilityColumnHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilityColumnHeaderUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnHeaderUIElements:"), value)
}

// The column index range of the cell.
//
// # Discussion
//
// This property contains the column’s starting index and index span in the
// table. Use this property in the elements representing a table’s cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumnIndexRange
func (o NSMenuItem) AccessibilityColumnIndexRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityColumnIndexRange"))
	return foundation.NSRange(rv)
}

func (o NSMenuItem) SetAccessibilityColumnIndexRange(value foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnIndexRange:"), value)
}

// The column titles for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumnTitles
func (o NSMenuItem) AccessibilityColumnTitles() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnTitles"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilityColumnTitles(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnTitles:"), value)
}

// The column accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumns
func (o NSMenuItem) AccessibilityColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumns"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilityColumns(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumns:"), value)
}

// The contents of the current accessibility element.
//
// # Discussion
//
// This property is used by container elements. It holds an array of the
// container’s contents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityContents
func (o NSMenuItem) AccessibilityContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityContents"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilityContents(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityContents:"), value)
}

// The critical value for the level indicator.
//
// # Discussion
//
// Use this property for elements such as the battery level indicator. This
// property sets a boundary value. If the element’s value exceeds the
// boundary value, the element has reached a critical stage.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCriticalValue
func (o NSMenuItem) AccessibilityCriticalValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCriticalValue"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityCriticalValue(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCriticalValue:"), value)
}

// The custom actions of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCustomActions
func (o NSMenuItem) AccessibilityCustomActions() []NSAccessibilityCustomAction {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityCustomActions"))
	result := make([]NSAccessibilityCustomAction, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = NSAccessibilityCustomActionFromID(id)
	}
	return result
}

func (o NSMenuItem) SetAccessibilityCustomActions(value []NSAccessibilityCustomAction) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomActions:"), objectivec.IObjectSliceToNSArray(value))
}

// The custom rotors of the current accessibility element.
//
// # Discussion
//
// Custom rotors are lists of items of a specific category. For example, a
// “headings” rotor returns a list of headings a given document.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCustomRotors
func (o NSMenuItem) AccessibilityCustomRotors() []NSAccessibilityCustomRotor {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityCustomRotors"))
	result := make([]NSAccessibilityCustomRotor, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = NSAccessibilityCustomRotorFromID(id)
	}
	return result
}

func (o NSMenuItem) SetAccessibilityCustomRotors(value []NSAccessibilityCustomRotor) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomRotors:"), objectivec.IObjectSliceToNSArray(value))
}

// The decrement button for the stepper accessibility element.
//
// # Discussion
//
// Use this property on a stepper.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDecrementButton
func (o NSMenuItem) AccessibilityDecrementButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDecrementButton"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityDecrementButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDecrementButton:"), value)
}

// The child accessibility element that represents the window’s default
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDefaultButton
func (o NSMenuItem) AccessibilityDefaultButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDefaultButton"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityDefaultButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDefaultButton:"), value)
}

// A Boolean value that determines whether the row is disclosing other rows.
//
// # Discussion
//
// Use this property in the elements representing an outline’s row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosed
func (o NSMenuItem) AccessibilityDisclosed() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityDisclosed"))
	return bool(rv)
}

func (o NSMenuItem) SetAccessibilityDisclosed(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosed:"), value)
}

// The row disclosing the current row.
//
// # Discussion
//
// Use this property in the elements representing an outline’s row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosedByRow
func (o NSMenuItem) AccessibilityDisclosedByRow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedByRow"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityDisclosedByRow(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedByRow:"), value)
}

// The rows that the current row discloses.
//
// # Discussion
//
// Use this property in the elements representing an outline’s row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosedRows
func (o NSMenuItem) AccessibilityDisclosedRows() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedRows"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityDisclosedRows(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedRows:"), value)
}

// The indention level for the row.
//
// # Discussion
//
// Use this property in the elements representing an outline’s row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosureLevel
func (o NSMenuItem) AccessibilityDisclosureLevel() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityDisclosureLevel"))
	return int(rv)
}

func (o NSMenuItem) SetAccessibilityDisclosureLevel(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosureLevel:"), value)
}

// The URL for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDocument
func (o NSMenuItem) AccessibilityDocument() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDocument"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSMenuItem) SetAccessibilityDocument(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDocument:"), objc.String(value))
}

// A Boolean value that indicates whether the accessibility element is in an
// edited state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityEdited
func (o NSMenuItem) AccessibilityEdited() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEdited"))
	return bool(rv)
}

func (o NSMenuItem) SetAccessibilityEdited(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEdited:"), value)
}

// A Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// # Discussion
//
// Use this property to expose this object to accessibility clients as a
// functional interface element. For example, when you place a button in a
// window, the system typically creates a button cell inside a button control
// inside a container view inside a window. Users, however, don’t care about
// the view hierarchy details. They should only be told that there’s a
// button in a window.
//
// If this property is set to false, accessibility clients ignore this
// element. By default, [NSView] and its subclasses set this value to false;
// however, if your [NSView] subclass adopts one of the accessibility
// protocols, the system changes the default value to true.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityElement
func (o NSMenuItem) AccessibilityElement() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityElement"))
	return bool(rv)
}

func (o NSMenuItem) SetAccessibilityElement(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityElement:"), value)
}

// A Boolean value that determines whether the accessibility element responds
// to user events.
//
// # Discussion
//
// Returns YES if the element is enabled; otherwise, NO. Enabled elements
// respond to user events.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityEnabled
func (o NSMenuItem) AccessibilityEnabled() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEnabled"))
	return bool(rv)
}

func (o NSMenuItem) SetAccessibilityEnabled(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEnabled:"), value)
}

// A Boolean value that determines whether the accessibility element is in an
// expanded state.
//
// # Discussion
//
// Use this property on elements that can expand to reveal additional
// information, such as outline rows and combo boxes.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityExpanded
func (o NSMenuItem) AccessibilityExpanded() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityExpanded"))
	return bool(rv)
}

func (o NSMenuItem) SetAccessibilityExpanded(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExpanded:"), value)
}

// The icon for the app’s menu bar extra.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityExtrasMenuBar
func (o NSMenuItem) AccessibilityExtrasMenuBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityExtrasMenuBar"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityExtrasMenuBar(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExtrasMenuBar:"), value)
}

// The filename for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFilename
func (o NSMenuItem) AccessibilityFilename() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFilename"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSMenuItem) SetAccessibilityFilename(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFilename:"), objc.String(value))
}

// A Boolean value that determines whether the accessibility element has the
// keyboard focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocused
func (o NSMenuItem) AccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return bool(rv)
}

func (o NSMenuItem) SetAccessibilityFocused(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocused:"), value)
}

// The child window with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocusedWindow
func (o NSMenuItem) AccessibilityFocusedWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFocusedWindow"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityFocusedWindow(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocusedWindow:"), value)
}

// The accessibility element’s frame in screen coordinates.
//
// # Discussion
//
// This property is accessed by the system whenever an accessibility client
// requests the element’s size or position.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
func (o NSMenuItem) SetAccessibilityFrame(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrame:"), value)
}

// A Boolean value that determines whether the app is the frontmost app.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrontmost
func (o NSMenuItem) AccessibilityFrontmost() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFrontmost"))
	return bool(rv)
}

func (o NSMenuItem) SetAccessibilityFrontmost(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrontmost:"), value)
}

// The child accessibility element that represents the window’s full-screen
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFullScreenButton
func (o NSMenuItem) AccessibilityFullScreenButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFullScreenButton"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityFullScreenButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFullScreenButton:"), value)
}

// The child accessibility element that represents the window’s grow area.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityGrowArea
func (o NSMenuItem) AccessibilityGrowArea() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityGrowArea"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityGrowArea(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityGrowArea:"), value)
}

// The drag handle accessibility elements for the layout item element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHandles
func (o NSMenuItem) AccessibilityHandles() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHandles"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilityHandles(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHandles:"), value)
}

// The header for the table view.
//
// # Discussion
//
// Use this property on a table view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHeader
func (o NSMenuItem) AccessibilityHeader() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHeader"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityHeader(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHeader:"), value)
}

// The help text for the accessibility element.
//
// # Discussion
//
// Use this property only when the results of activating this element are not
// obvious from the element’s label. This string functions as a tooltip. For
// example, VoiceOver reads this string when you pause over a control. To help
// ensure that accessibility clients like VoiceOver read the help text with
// the proper inflection, begin this string with a verb, capitalize the first
// letter, and end the string with a period. Always localize this string. The
// default value is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHelp
func (o NSMenuItem) AccessibilityHelp() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHelp"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSMenuItem) SetAccessibilityHelp(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHelp:"), objc.String(value))
}

// A Boolean value that determines whether the app is in a hidden state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHidden
func (o NSMenuItem) AccessibilityHidden() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityHidden"))
	return bool(rv)
}

func (o NSMenuItem) SetAccessibilityHidden(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHidden:"), value)
}

// The horizontal scroll bar for the scroll view.
//
// # Discussion
//
// Use this property on a scrollable view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHorizontalScrollBar
func (o NSMenuItem) AccessibilityHorizontalScrollBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalScrollBar"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityHorizontalScrollBar(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalScrollBar:"), value)
}

// A description of the layout area’s horizontal units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHorizontalUnitDescription
func (o NSMenuItem) AccessibilityHorizontalUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSMenuItem) SetAccessibilityHorizontalUnitDescription(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnitDescription:"), objc.String(value))
}

// The units that the layout area uses for horizontal values.
//
// # Discussion
//
// For a list of possible values, see [NSAccessibilityUnits].
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHorizontalUnits
//
// [NSAccessibilityUnits]: https://developer.apple.com/documentation/AppKit/NSAccessibilityUnits
func (o NSMenuItem) AccessibilityHorizontalUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityHorizontalUnits"))
	return NSAccessibilityUnits(rv)
}

func (o NSMenuItem) SetAccessibilityHorizontalUnits(value NSAccessibilityUnits) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnits:"), value)
}

// The accessibility element’s identity.
//
// # Discussion
//
// This property holds the unique ID for the accessibility element. It is
// often used in automated testing.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIdentifier
func (o NSMenuItem) SetAccessibilityIdentifier(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIdentifier:"), objc.String(value))
}

// The increment button for the stepper accessibility element.
//
// # Discussion
//
// Use this property on a stepper.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIncrementButton
func (o NSMenuItem) AccessibilityIncrementButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIncrementButton"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityIncrementButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIncrementButton:"), value)
}

// The index of the row or column that the accessibility element represents.
//
// # Discussion
//
// Use this property for any element that can be accessed through an index:
// cells, rows, columns, and so forth.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIndex
func (o NSMenuItem) AccessibilityIndex() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityIndex"))
	return int(rv)
}

func (o NSMenuItem) SetAccessibilityIndex(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIndex:"), value)
}

// The line number that contains the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityInsertionPointLineNumber
func (o NSMenuItem) AccessibilityInsertionPointLineNumber() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityInsertionPointLineNumber"))
	return int(rv)
}

func (o NSMenuItem) SetAccessibilityInsertionPointLineNumber(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityInsertionPointLineNumber:"), value)
}

// A short description of the accessibility element.
//
// # Discussion
//
// Do not include the accessibility element’s type in the label (for
// example, write [Play], not `Play button`.). If possible, use a single word.
// To help ensure that accessibility clients such as VoiceOver read the label
// with the correct intonation, start this label with a capital letter. Do not
// put a period at the end. Always localize the label.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLabel
func (o NSMenuItem) AccessibilityLabel() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSMenuItem) SetAccessibilityLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabel:"), objc.String(value))
}

// The child label elements for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLabelUIElements
func (o NSMenuItem) AccessibilityLabelUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabelUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilityLabelUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelUIElements:"), value)
}

// The value of the label accessibility element.
//
// # Discussion
//
// Use this property on a slider element’s labels.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLabelValue
func (o NSMenuItem) AccessibilityLabelValue() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("accessibilityLabelValue"))
	return float32(rv)
}

func (o NSMenuItem) SetAccessibilityLabelValue(value float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelValue:"), value)
}

// The elements that have links with the accessibility element.
//
// # Discussion
//
// Use this property to define a relationship between different user interface
// elements. For example, use this property to link a list item with contents
// displayed in another pane or window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLinkedUIElements
func (o NSMenuItem) AccessibilityLinkedUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLinkedUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilityLinkedUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLinkedUIElements:"), value)
}

// A Boolean value that determines whether the window is the app’s main
// window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMain
func (o NSMenuItem) AccessibilityMain() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMain"))
	return bool(rv)
}

func (o NSMenuItem) SetAccessibilityMain(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMain:"), value)
}

// The app’s main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMainWindow
func (o NSMenuItem) AccessibilityMainWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMainWindow"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityMainWindow(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMainWindow:"), value)
}

// The user interface element that functions as a marker group for the ruler
// accessibility element.
//
// # Discussion
//
// Use this property on a ruler element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMarkerGroupUIElement
func (o NSMenuItem) AccessibilityMarkerGroupUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerGroupUIElement"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityMarkerGroupUIElement(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerGroupUIElement:"), value)
}

// A human-readable description of the marker type.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMarkerTypeDescription
func (o NSMenuItem) AccessibilityMarkerTypeDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerTypeDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSMenuItem) SetAccessibilityMarkerTypeDescription(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerTypeDescription:"), objc.String(value))
}

// An array of marker accessibility elements for the ruler.
//
// # Discussion
//
// Use this property on a ruler element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMarkerUIElements
func (o NSMenuItem) AccessibilityMarkerUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilityMarkerUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerUIElements:"), value)
}

// The marker values for the ruler.
//
// # Discussion
//
// Use this property on a ruler element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMarkerValues
func (o NSMenuItem) AccessibilityMarkerValues() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerValues"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityMarkerValues(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerValues:"), value)
}

// The maximum value for the accessibility element.
//
// # Discussion
//
// This property is set to `nil` by default. Only a few AppKit controls (for
// example, [NSSliderCell]) support this value. Set this property only when
// the element has an [accessibilityValue] property and you want to define the
// maximum possible value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMaxValue
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
func (o NSMenuItem) AccessibilityMaxValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMaxValue"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityMaxValue(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMaxValue:"), value)
}

// The app’s menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMenuBar
func (o NSMenuItem) AccessibilityMenuBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMenuBar"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityMenuBar(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMenuBar:"), value)
}

// The minimum value for the accessibility element.
//
// # Discussion
//
// This property is set to `nil` by default. Only a few AppKit controls (for
// example, [NSSliderCell]) support this value. Set this property only when
// the element has an [accessibilityValue] property and you want to define the
// minimum possible value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMinValue
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
func (o NSMenuItem) AccessibilityMinValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinValue"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityMinValue(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinValue:"), value)
}

// The child accessibility element that represents the window’s minimize
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMinimizeButton
func (o NSMenuItem) AccessibilityMinimizeButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinimizeButton"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityMinimizeButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimizeButton:"), value)
}

// A Boolean value that determines whether this window is in a minimized
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMinimized
func (o NSMenuItem) AccessibilityMinimized() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMinimized"))
	return bool(rv)
}

func (o NSMenuItem) SetAccessibilityMinimized(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimized:"), value)
}

// A Boolean value that determines whether the window is modal.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityModal
func (o NSMenuItem) AccessibilityModal() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityModal"))
	return bool(rv)
}

func (o NSMenuItem) SetAccessibilityModal(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityModal:"), value)
}

// The contents that follow the divider accessibility element.
//
// # Discussion
//
// For example, use this property to set the subview adjacent to a split
// view’s splitter element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityNextContents
func (o NSMenuItem) AccessibilityNextContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityNextContents"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilityNextContents(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNextContents:"), value)
}

// The number of characters in the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityNumberOfCharacters
func (o NSMenuItem) AccessibilityNumberOfCharacters() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityNumberOfCharacters"))
	return int(rv)
}

func (o NSMenuItem) SetAccessibilityNumberOfCharacters(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNumberOfCharacters:"), value)
}

// A Boolean value that determines whether the accessibility element’s grid
// is in row major order or in column major order.
//
// # Discussion
//
// Use this property for UI elements that present a grid of child elements.
// Set the property to true if the grid is ordered row major; otherwise, set
// to false.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityOrderedByRow
func (o NSMenuItem) AccessibilityOrderedByRow() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityOrderedByRow"))
	return bool(rv)
}

func (o NSMenuItem) SetAccessibilityOrderedByRow(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrderedByRow:"), value)
}

// The orientation of the accessibility element.
//
// # Discussion
//
// This property can hold either the [NSAccessibilityOrientationHorizontal]
// value or the [NSAccessibilityOrientationVertical] value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityOrientation
func (o NSMenuItem) AccessibilityOrientation() NSAccessibilityOrientation {
	rv := objc.Send[NSAccessibilityOrientation](o.ID, objc.Sel("accessibilityOrientation"))
	return NSAccessibilityOrientation(rv)
}

func (o NSMenuItem) SetAccessibilityOrientation(value NSAccessibilityOrientation) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrientation:"), value)
}

// The overflow button for the toolbar.
//
// # Discussion
//
// Use this property on a toolbar element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityOverflowButton
func (o NSMenuItem) AccessibilityOverflowButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityOverflowButton"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityOverflowButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOverflowButton:"), value)
}

// The accessibility element’s parent in the accessibility hierarchy.
//
// # Discussion
//
// This property must contain a reference to another element in the
// accessibility hierarchy. If you create an [NSView] subclass, you don’t
// typically need to set this value. The system automatically sets the parent
// to the nearest ancestor in the view hierarchy that is also in the
// accessibility hierarchy. If you use an [NSAccessibilityElement] subclass to
// represent an interface element that is not backed by a view, you can either
// set the parent property or you can call the
// [NSAccessibilityElementClass.AccessibilityElementWithRoleFrameLabelParent]
// convenience method, which sets it automatically.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityParent
func (o NSMenuItem) SetAccessibilityParent(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityParent:"), value)
}

// The placeholder value for the accessibility element.
//
// # Discussion
//
// Use this property for accessibility elements that support placeholder
// values, such as text fields.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityPlaceholderValue
func (o NSMenuItem) AccessibilityPlaceholderValue() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPlaceholderValue"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSMenuItem) SetAccessibilityPlaceholderValue(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPlaceholderValue:"), objc.String(value))
}

// The contents that precede the divider accessibility element.
//
// # Discussion
//
// For example, use this property to set the subview adjacent to a split
// view’s splitter element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityPreviousContents
func (o NSMenuItem) AccessibilityPreviousContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPreviousContents"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilityPreviousContents(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPreviousContents:"), value)
}

// A Boolean value that determines whether the accessibility element contains
// protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityProtectedContent
func (o NSMenuItem) AccessibilityProtectedContent() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityProtectedContent"))
	return bool(rv)
}

func (o NSMenuItem) SetAccessibilityProtectedContent(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProtectedContent:"), value)
}

// The child accessibility element that represents the window’s proxy icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityProxy
func (o NSMenuItem) AccessibilityProxy() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityProxy"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityProxy(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProxy:"), value)
}

// A Boolean value that determines whether the accessibility element must have
// content for successful submission of a form.
//
// # Discussion
//
// Returns YES if the element is required to have content; otherwise, NO.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRequired
func (o NSMenuItem) AccessibilityRequired() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityRequired"))
	return bool(rv)
}

func (o NSMenuItem) SetAccessibilityRequired(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRequired:"), value)
}

// The type of interface element that the accessibility element represents.
//
// # Discussion
//
// This property contains a nonlocalized string that defines the element’s
// role in the app. For a list of possible roles, see [Roles]. This property
// is set automatically when you adopt one of the accessibility protocols.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRole
func (o NSMenuItem) AccessibilityRole() NSAccessibilityRole {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRole"))
	return NSAccessibilityRole(foundation.NSStringFromID(rv).String())
}

func (o NSMenuItem) SetAccessibilityRole(value NSAccessibilityRole) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRole:"), objc.String(string(value)))
}

// A localized, human-intelligible description of the accessibility
// element’s role, such as .
//
// # Discussion
//
// This property is set automatically based on the value of the
// [accessibilityRole] property; however, you can customize the value of this
// property to better describe your element’s role. Keep role descriptions
// short. If possible, use a single word. These descriptions should be noun
// phrases, all lowercase, with no period at the end.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRoleDescription
//
// [accessibilityRole]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRole
func (o NSMenuItem) AccessibilityRoleDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRoleDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSMenuItem) SetAccessibilityRoleDescription(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRoleDescription:"), objc.String(value))
}

// The number of rows in the accessibility element’s grid.
//
// # Discussion
//
// Use this property for elements that present a grid of child elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRowCount
func (o NSMenuItem) AccessibilityRowCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityRowCount"))
	return int(rv)
}

func (o NSMenuItem) SetAccessibilityRowCount(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowCount:"), value)
}

// The row header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRowHeaderUIElements
func (o NSMenuItem) AccessibilityRowHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRowHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilityRowHeaderUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowHeaderUIElements:"), value)
}

// The row index range of the cell.
//
// # Discussion
//
// This property contains the row’s starting index and index span in the
// table. Use this property in the elements representing a table’s cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRowIndexRange
func (o NSMenuItem) AccessibilityRowIndexRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRowIndexRange"))
	return foundation.NSRange(rv)
}

func (o NSMenuItem) SetAccessibilityRowIndexRange(value foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowIndexRange:"), value)
}

// The row accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRows
func (o NSMenuItem) AccessibilityRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRows"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilityRows(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRows:"), value)
}

// The type of markers for the ruler.
//
// # Discussion
//
// Use this property on a ruler element. For a complete list of marker types,
// see [NSAccessibilityRulerMarkerType].
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRulerMarkerType
//
// [NSAccessibilityRulerMarkerType]: https://developer.apple.com/documentation/AppKit/NSAccessibilityRulerMarkerType
func (o NSMenuItem) AccessibilityRulerMarkerType() NSAccessibilityRulerMarkerType {
	rv := objc.Send[NSAccessibilityRulerMarkerType](o.ID, objc.Sel("accessibilityRulerMarkerType"))
	return NSAccessibilityRulerMarkerType(rv)
}

func (o NSMenuItem) SetAccessibilityRulerMarkerType(value NSAccessibilityRulerMarkerType) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRulerMarkerType:"), value)
}

// The search button for the search field.
//
// # Discussion
//
// Use this property on a search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySearchButton
func (o NSMenuItem) AccessibilitySearchButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchButton"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilitySearchButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchButton:"), value)
}

// The search menu for the search field.
//
// # Discussion
//
// Use this property on a search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySearchMenu
func (o NSMenuItem) AccessibilitySearchMenu() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchMenu"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilitySearchMenu(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchMenu:"), value)
}

// A Boolean value that determines whether the accessibility element is
// currently in a selected state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelected
func (o NSMenuItem) AccessibilitySelected() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelected"))
	return bool(rv)
}

func (o NSMenuItem) SetAccessibilitySelected(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelected:"), value)
}

// The currently selected cells for the table.
//
// # Discussion
//
// This property is required for all elements that act like cell-based tables.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedCells
func (o NSMenuItem) AccessibilitySelectedCells() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedCells"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilitySelectedCells(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedCells:"), value)
}

// The accessibility element’s currently selected children.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedChildren
func (o NSMenuItem) AccessibilitySelectedChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedChildren"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilitySelectedChildren(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedChildren:"), value)
}

// The currently selected columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedColumns
func (o NSMenuItem) AccessibilitySelectedColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedColumns"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilitySelectedColumns(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedColumns:"), value)
}

// The currently selected rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedRows
func (o NSMenuItem) AccessibilitySelectedRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedRows"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilitySelectedRows(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedRows:"), value)
}

// The currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedText
func (o NSMenuItem) AccessibilitySelectedText() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedText"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSMenuItem) SetAccessibilitySelectedText(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedText:"), objc.String(value))
}

// The range of the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedTextRange
func (o NSMenuItem) AccessibilitySelectedTextRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySelectedTextRange"))
	return foundation.NSRange(rv)
}

func (o NSMenuItem) SetAccessibilitySelectedTextRange(value foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRange:"), value)
}

// An array of ranges for the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedTextRanges
func (o NSMenuItem) AccessibilitySelectedTextRanges() []foundation.NSValue {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilitySelectedTextRanges"))
	result := make([]foundation.NSValue, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = foundation.NSValueFromID(id)
	}
	return result
}

func (o NSMenuItem) SetAccessibilitySelectedTextRanges(value []foundation.NSValue) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRanges:"), objectivec.IObjectSliceToNSArray(value))
}

// The list of elements that the accessibility element is a title for.
//
// # Discussion
//
// Use on a static text label to associate that label with one or more user
// interface elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityServesAsTitleForUIElements
func (o NSMenuItem) AccessibilityServesAsTitleForUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityServesAsTitleForUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilityServesAsTitleForUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityServesAsTitleForUIElements:"), value)
}

// The range of characters that the accessibility element displays.
//
// # Discussion
//
// Use this property to manage text that is split across multiple
// elements—for example, an ebook reader that splits the text into multiple
// pages.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySharedCharacterRange
func (o NSMenuItem) AccessibilitySharedCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySharedCharacterRange"))
	return foundation.NSRange(rv)
}

func (o NSMenuItem) SetAccessibilitySharedCharacterRange(value foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedCharacterRange:"), value)
}

// An array of elements that shares the keyboard focus with the accessibility
// element.
//
// # Discussion
//
// Use this property to manage elements that share the keyboard focus—for
// example, a search field with completion menu below it.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySharedFocusElements
func (o NSMenuItem) AccessibilitySharedFocusElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedFocusElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilitySharedFocusElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedFocusElements:"), value)
}

// Other elements that share text with the accessibility element.
//
// # Discussion
//
// Use this property to manage text that is split across multiple
// elements—for example, an ebook reader that splits the text into multiple
// pages.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySharedTextUIElements
func (o NSMenuItem) AccessibilitySharedTextUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedTextUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilitySharedTextUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedTextUIElements:"), value)
}

// The menu currently displaying for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityShownMenu
func (o NSMenuItem) AccessibilityShownMenu() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityShownMenu"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityShownMenu(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityShownMenu:"), value)
}

// The accessibility element’s sort direction.
//
// # Discussion
//
// Used by an element with an [button] role and an
// [NSAccessibilitySortButtonRole] subrole. For a list of possible sort
// directions, see [NSAccessibilitySortDirection].
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySortDirection
//
// [NSAccessibilitySortButtonRole]: https://developer.apple.com/documentation/AppKit/NSAccessibilitySortButtonRole
// [NSAccessibilitySortDirection]: https://developer.apple.com/documentation/AppKit/NSAccessibilitySortDirection
// [button]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/button
func (o NSMenuItem) AccessibilitySortDirection() NSAccessibilitySortDirection {
	rv := objc.Send[NSAccessibilitySortDirection](o.ID, objc.Sel("accessibilitySortDirection"))
	return NSAccessibilitySortDirection(rv)
}

func (o NSMenuItem) SetAccessibilitySortDirection(value NSAccessibilitySortDirection) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySortDirection:"), value)
}

// An array that contains the views and splitter bar from the split view.
//
// # Discussion
//
// Use this property on a split view element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySplitters
func (o NSMenuItem) AccessibilitySplitters() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySplitters"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilitySplitters(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySplitters:"), value)
}

// The specialized interface element type that the accessibility element
// represents.
//
// # Discussion
//
// For a list of possible subroles, see [Subroles].
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySubrole
func (o NSMenuItem) AccessibilitySubrole() NSAccessibilitySubrole {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySubrole"))
	return NSAccessibilitySubrole(foundation.NSStringFromID(rv).String())
}

func (o NSMenuItem) SetAccessibilitySubrole(value NSAccessibilitySubrole) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySubrole:"), objc.String(string(value)))
}

// The tab accessibility elements for the tab view.
//
// # Discussion
//
// Use this property on a tab view element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityTabs
func (o NSMenuItem) AccessibilityTabs() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTabs"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilityTabs(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTabs:"), value)
}

// The title of the accessibility element—for example, a button’s visible
// text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityTitle
func (o NSMenuItem) AccessibilityTitle() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitle"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSMenuItem) SetAccessibilityTitle(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitle:"), objc.String(value))
}

// A static text element that represents the accessibility element’s title.
//
// # Discussion
//
// Use this property to associate a static text label with another
// element—for example, to associate a label with its corresponding text
// field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityTitleUIElement
func (o NSMenuItem) AccessibilityTitleUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitleUIElement"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityTitleUIElement(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitleUIElement:"), value)
}

// The child accessibility element that represents the window’s toolbar
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityToolbarButton
func (o NSMenuItem) AccessibilityToolbarButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityToolbarButton"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityToolbarButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityToolbarButton:"), value)
}

// The top-level element that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityTopLevelUIElement
func (o NSMenuItem) AccessibilityTopLevelUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTopLevelUIElement"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityTopLevelUIElement(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTopLevelUIElement:"), value)
}

// The URL for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityURL
func (o NSMenuItem) AccessibilityURL() foundation.NSURL {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityURL"))
	return foundation.NSURLFromID(rv)
}

func (o NSMenuItem) SetAccessibilityURL(value foundation.NSURL) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityURL:"), value)
}

// A human-readable description of the ruler’s units.
//
// # Discussion
//
// Use this property on a ruler element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityUnitDescription
func (o NSMenuItem) AccessibilityUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSMenuItem) SetAccessibilityUnitDescription(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnitDescription:"), objc.String(value))
}

// The units for the ruler.
//
// # Discussion
//
// Use this property on a ruler element. For a complete list of units, see
// [NSAccessibilityUnits].
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityUnits
//
// [NSAccessibilityUnits]: https://developer.apple.com/documentation/AppKit/NSAccessibilityUnits
func (o NSMenuItem) AccessibilityUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityUnits"))
	return NSAccessibilityUnits(rv)
}

func (o NSMenuItem) SetAccessibilityUnits(value NSAccessibilityUnits) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnits:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityUserInputLabels
func (o NSMenuItem) AccessibilityUserInputLabels() []string {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityUserInputLabels"))
	return objc.ConvertSliceToStrings(rvIDs)
}

func (o NSMenuItem) SetAccessibilityUserInputLabels(value []string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUserInputLabels:"), objectivec.StringSliceToNSArray(value))
}

// The accessibility element’s value.
//
// # Discussion
//
// The accessibility protocols for roles that support values typically
// redefine this property to take a more specific value type. For example, the
// [staticText] protocol uses [NSString] values, and the [progressIndicator]
// protocol uses [NSNumber] values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [progressIndicator]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/progressIndicator
// [staticText]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Role/staticText
func (o NSMenuItem) AccessibilityValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValue"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityValue(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValue:"), value)
}

// A human-readable description of the accessibility element’s value.
//
// # Discussion
//
// Use this property to provide a more useful description of the accessibility
// element’s raw value. For example, you might set the value to `600`, but
// set the description to `10 minutes`. Always localize this description.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValueDescription
func (o NSMenuItem) AccessibilityValueDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValueDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSMenuItem) SetAccessibilityValueDescription(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValueDescription:"), objc.String(value))
}

// The vertical scroll bar for the scroll view.
//
// # Discussion
//
// Use this property on a scrollable view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVerticalScrollBar
func (o NSMenuItem) AccessibilityVerticalScrollBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalScrollBar"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityVerticalScrollBar(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalScrollBar:"), value)
}

// A description of the layout area’s vertical units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVerticalUnitDescription
func (o NSMenuItem) AccessibilityVerticalUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSMenuItem) SetAccessibilityVerticalUnitDescription(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnitDescription:"), objc.String(value))
}

// The units that the layout area uses for vertical values.
//
// # Discussion
//
// For a list of possible values, see [NSAccessibilityUnits].
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVerticalUnits
//
// [NSAccessibilityUnits]: https://developer.apple.com/documentation/AppKit/NSAccessibilityUnits
func (o NSMenuItem) AccessibilityVerticalUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityVerticalUnits"))
	return NSAccessibilityUnits(rv)
}

func (o NSMenuItem) SetAccessibilityVerticalUnits(value NSAccessibilityUnits) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnits:"), value)
}

// The visible cells for the table.
//
// # Discussion
//
// This property is required for all elements that act like cell-based tables.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleCells
func (o NSMenuItem) AccessibilityVisibleCells() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleCells"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilityVisibleCells(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleCells:"), value)
}

// The range of visible characters in the document.
//
// # Discussion
//
// Use this property to store the range for entire lines. Characters that are
// horizontally clipped are included in this range.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleCharacterRange
func (o NSMenuItem) AccessibilityVisibleCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityVisibleCharacterRange"))
	return foundation.NSRange(rv)
}

func (o NSMenuItem) SetAccessibilityVisibleCharacterRange(value foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleCharacterRange:"), value)
}

// The accessibility element’s visible child accessibility elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleChildren
func (o NSMenuItem) AccessibilityVisibleChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleChildren"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilityVisibleChildren(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleChildren:"), value)
}

// The visible columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleColumns
func (o NSMenuItem) AccessibilityVisibleColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleColumns"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilityVisibleColumns(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleColumns:"), value)
}

// The visible rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleRows
func (o NSMenuItem) AccessibilityVisibleRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleRows"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilityVisibleRows(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleRows:"), value)
}

// The warning value for the level indicator.
//
// # Discussion
//
// Use this property for elements such as the battery level indicator. This
// property sets a boundary value. If the element’s value exceeds the
// boundary value, the element has reached a warning stage.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityWarningValue
func (o NSMenuItem) AccessibilityWarningValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWarningValue"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityWarningValue(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWarningValue:"), value)
}

// The window that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityWindow
func (o NSMenuItem) AccessibilityWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindow"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityWindow(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindow:"), value)
}

// An array that contains all the app’s windows.
//
// # Discussion
//
// Use on the app element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityWindows
func (o NSMenuItem) AccessibilityWindows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindows"))
	return foundation.NSArrayFromID(rv)
}

func (o NSMenuItem) SetAccessibilityWindows(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindows:"), value)
}

// The child accessibility element that represents the window’s zoom button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityZoomButton
func (o NSMenuItem) AccessibilityZoomButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityZoomButton"))
	return objectivec.Object{ID: rv}
}

func (o NSMenuItem) SetAccessibilityZoomButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityZoomButton:"), value)
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
func (o NSMenuItem) SetIdentifier(value NSUserInterfaceItemIdentifier) {
	objc.Send[struct{}](o.ID, objc.Sel("setIdentifier:"), objc.String(string(value)))
}

// Protocol methods for NSValidatedUserInterfaceItem
