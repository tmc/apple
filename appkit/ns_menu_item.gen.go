// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

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
// dialog is open, the target object should return true in [WorksWhenModal].
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
func (m NSMenuItem) Action() objc.SEL {
	rv := objc.Send[objc.SEL](m.ID, objc.Sel("action"))
	return rv
}
func (m NSMenuItem) SetAction(value objc.SEL) {
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
// [AllowsAutomaticKeyEquivalentMirroring] property.
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
