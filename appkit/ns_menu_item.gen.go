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
//   - [NSMenuItem.Enabled]: A Boolean value that indicates whether the menu item is enabled.
//   - [NSMenuItem.SetEnabled]
//
// # Managing hidden status
//
//   - [NSMenuItem.Hidden]: A Boolean value that indicates whether the menu item is hidden.
//   - [NSMenuItem.SetHidden]
//   - [NSMenuItem.HiddenOrHasHiddenAncestor]: A Boolean value that indicates whether the menu item or any of its superitems is hidden.
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
//   - [NSMenuItem.SectionHeader]: A Boolean value indicating whether the menu item is a section header.
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
//   - [NSMenuItem.SeparatorItem]: A Boolean value indicating whether the menu item is a separator item.
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
//   - [NSMenuItem.Alternate]: A Boolean value that marks the menu item as an alternate to the previous menu item.
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
//   - [NSMenuItem.Highlighted]: A Boolean value that indicates whether the menu item should be drawn highlighted.
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
//   - [INSMenuItem.Enabled]: A Boolean value that indicates whether the menu item is enabled.
//   - [INSMenuItem.SetEnabled]
//
// # Managing hidden status
//
//   - [INSMenuItem.Hidden]: A Boolean value that indicates whether the menu item is hidden.
//   - [INSMenuItem.SetHidden]
//   - [INSMenuItem.HiddenOrHasHiddenAncestor]: A Boolean value that indicates whether the menu item or any of its superitems is hidden.
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
//   - [INSMenuItem.SectionHeader]: A Boolean value indicating whether the menu item is a section header.
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
//   - [INSMenuItem.SeparatorItem]: A Boolean value indicating whether the menu item is a separator item.
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
//   - [INSMenuItem.Alternate]: A Boolean value that marks the menu item as an alternate to the previous menu item.
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
//   - [INSMenuItem.Highlighted]: A Boolean value that indicates whether the menu item should be drawn highlighted.
//
// # Instance Properties
//
//   - [INSMenuItem.Subtitle]
//   - [INSMenuItem.SetSubtitle]
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItem
type INSMenuItem interface {
	objectivec.IObject
	NSUserInterfaceItemIdentification
	NSValidatedUserInterfaceItem

	// Topic: Creating a menu item

	// Returns an initialized instance of [NSMenuItem].
	InitWithTitleActionKeyEquivalent(string_ string, selector objc.SEL, charCode string) NSMenuItem
	InitWithCoder(coder foundation.INSCoder) NSMenuItem

	// Topic: Enabling a menu item

	// A Boolean value that indicates whether the menu item is enabled.
	Enabled() bool
	SetEnabled(value bool)

	// Topic: Managing hidden status

	// A Boolean value that indicates whether the menu item is hidden.
	Hidden() bool
	SetHidden(value bool)
	// A Boolean value that indicates whether the menu item or any of its superitems is hidden.
	HiddenOrHasHiddenAncestor() bool

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
	SectionHeader() bool

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
	SeparatorItem() bool

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
	Alternate() bool
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
	Highlighted() bool

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
func (m NSMenuItem) Enabled() bool {
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
func (m NSMenuItem) Hidden() bool {
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
func (m NSMenuItem) HiddenOrHasHiddenAncestor() bool {
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
func (m NSMenuItem) SectionHeader() bool {
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
func (m NSMenuItem) SeparatorItem() bool {
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
func (m NSMenuItem) Alternate() bool {
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
func (m NSMenuItem) Highlighted() bool {
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
func (m NSMenuItem) Identifier() NSUserInterfaceItemIdentifier {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("identifier"))
	return NSUserInterfaceItemIdentifier(foundation.NSStringFromID(rv).String())
}
func (m NSMenuItem) SetIdentifier(value NSUserInterfaceItemIdentifier) {
	objc.Send[struct{}](m.ID, objc.Sel("setIdentifier:"), objc.String(string(value)))
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

// Sets a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityElement(_:)
func (o NSMenuItem) SetAccessibilityElement(accessibilityElement bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityElement:"), accessibilityElement)
}

// Returns a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEnabled()
func (o NSMenuItem) IsAccessibilityEnabled() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEnabled"))
	return rv
}

// Sets a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityEnabled(_:)
func (o NSMenuItem) SetAccessibilityEnabled(accessibilityEnabled bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEnabled:"), accessibilityEnabled)
}

// Sets the accessibility element’s frame in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFrame(_:)
func (o NSMenuItem) SetAccessibilityFrame(accessibilityFrame corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrame:"), accessibilityFrame)
}

// Returns the help text for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHelp()
func (o NSMenuItem) AccessibilityHelp() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHelp"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the help text for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHelp(_:)
func (o NSMenuItem) SetAccessibilityHelp(accessibilityHelp string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHelp:"), objc.String(accessibilityHelp))
}

// Returns a short description of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLabel()
func (o NSMenuItem) AccessibilityLabel() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
}

// Sets a short description of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLabel(_:)
func (o NSMenuItem) SetAccessibilityLabel(accessibilityLabel string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabel:"), objc.String(accessibilityLabel))
}

// Returns the title of the accessibility element—for example, a button’s
// visible text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTitle()
func (o NSMenuItem) AccessibilityTitle() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitle"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the title of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTitle(_:)
func (o NSMenuItem) SetAccessibilityTitle(accessibilityTitle string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitle:"), objc.String(accessibilityTitle))
}

// Returns the accessibility element’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityValue()
func (o NSMenuItem) AccessibilityValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValue"))
	return objectivec.Object{ID: rv}
}

// Sets the accessibility element’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityValue(_:)
func (o NSMenuItem) SetAccessibilityValue(accessibilityValue objectivec.IObject) {
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
func (o NSMenuItem) IsAccessibilitySelectorAllowed(selector objc.SEL) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelectorAllowed:"), selector)
	return rv
}

// Returns the contents of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityContents()
func (o NSMenuItem) AccessibilityContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityContents"))
	return foundation.NSArrayFromID(rv)
}

// Sets the contents of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityContents(_:)
func (o NSMenuItem) SetAccessibilityContents(accessibilityContents objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityContents:"), accessibilityContents)
}

// Returns the critical value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCriticalValue()
func (o NSMenuItem) AccessibilityCriticalValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCriticalValue"))
	return objectivec.Object{ID: rv}
}

// Sets the critical value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCriticalValue(_:)
func (o NSMenuItem) SetAccessibilityCriticalValue(accessibilityCriticalValue objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCriticalValue:"), accessibilityCriticalValue)
}

// Sets the accessibility element’s identity.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityIdentifier(_:)
func (o NSMenuItem) SetAccessibilityIdentifier(accessibilityIdentifier string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIdentifier:"), objc.String(accessibilityIdentifier))
}

// Returns the maximum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMaxValue()
func (o NSMenuItem) AccessibilityMaxValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMaxValue"))
	return objectivec.Object{ID: rv}
}

// Sets the maximum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMaxValue(_:)
func (o NSMenuItem) SetAccessibilityMaxValue(accessibilityMaxValue objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMaxValue:"), accessibilityMaxValue)
}

// Returns the minimum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMinValue()
func (o NSMenuItem) AccessibilityMinValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinValue"))
	return objectivec.Object{ID: rv}
}

// Sets the minimum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMinValue(_:)
func (o NSMenuItem) SetAccessibilityMinValue(accessibilityMinValue objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinValue:"), accessibilityMinValue)
}

// Returns the orientation of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityOrientation()
func (o NSMenuItem) AccessibilityOrientation() NSAccessibilityOrientation {
	rv := objc.Send[NSAccessibilityOrientation](o.ID, objc.Sel("accessibilityOrientation"))
	return rv
}

// Sets the orientation of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityOrientation(_:)
func (o NSMenuItem) SetAccessibilityOrientation(accessibilityOrientation NSAccessibilityOrientation) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrientation:"), accessibilityOrientation)
}

// Returns a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityProtectedContent()
func (o NSMenuItem) IsAccessibilityProtectedContent() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityProtectedContent"))
	return rv
}

// Sets a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityProtectedContent(_:)
func (o NSMenuItem) SetAccessibilityProtectedContent(accessibilityProtectedContent bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProtectedContent:"), accessibilityProtectedContent)
}

// Returns a Boolean value that determines whether the accessibility element
// is currently in a selected state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelected()
func (o NSMenuItem) IsAccessibilitySelected() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelected"))
	return rv
}

// Sets a Boolean value that determines whether the accessibility element is
// currently in a selected state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelected(_:)
func (o NSMenuItem) SetAccessibilitySelected(accessibilitySelected bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelected:"), accessibilitySelected)
}

// Returns the URL for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityURL()
func (o NSMenuItem) AccessibilityURL() foundation.INSURL {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityURL"))
	return foundation.NSURLFromID(rv)
}

// Sets the URL for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityURL(_:)
func (o NSMenuItem) SetAccessibilityURL(accessibilityURL foundation.INSURL) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityURL:"), accessibilityURL)
}

// Returns the human-readable description of the accessibility element’s
// value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityValueDescription()
func (o NSMenuItem) AccessibilityValueDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValueDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the human-readable description of the accessibility element’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityValueDescription(_:)
func (o NSMenuItem) SetAccessibilityValueDescription(accessibilityValueDescription string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValueDescription:"), objc.String(accessibilityValueDescription))
}

// Returns the warning value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityWarningValue()
func (o NSMenuItem) AccessibilityWarningValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWarningValue"))
	return objectivec.Object{ID: rv}
}

// Sets the warning value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityWarningValue(_:)
func (o NSMenuItem) SetAccessibilityWarningValue(accessibilityWarningValue objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWarningValue:"), accessibilityWarningValue)
}

// Returns the child accessibility elements in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityChildren()
func (o NSMenuItem) AccessibilityChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityChildren"))
	return foundation.NSArrayFromID(rv)
}

// Sets the child accessibility elements in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityChildren(_:)
func (o NSMenuItem) SetAccessibilityChildren(accessibilityChildren objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildren:"), accessibilityChildren)
}

// Returns the array of child accessibility elements in order for linear
// navigation.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityChildrenInNavigationOrder()
func (o NSMenuItem) AccessibilityChildrenInNavigationOrder() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityChildrenInNavigationOrder"))
	return foundation.NSArrayFromID(rv)
}

// Sets the array of child accessibility elements in order for linear
// navigation.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityChildrenInNavigationOrder(_:)
func (o NSMenuItem) SetAccessibilityChildrenInNavigationOrder(accessibilityChildrenInNavigationOrder objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildrenInNavigationOrder:"), accessibilityChildrenInNavigationOrder)
}

// Sets the accessibility element’s parent in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityParent(_:)
func (o NSMenuItem) SetAccessibilityParent(accessibilityParent objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityParent:"), accessibilityParent)
}

// Returns the accessibility element’s currently selected children.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedChildren()
func (o NSMenuItem) AccessibilitySelectedChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedChildren"))
	return foundation.NSArrayFromID(rv)
}

// Sets the accessibility element’s currently selected children.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedChildren(_:)
func (o NSMenuItem) SetAccessibilitySelectedChildren(accessibilitySelectedChildren objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedChildren:"), accessibilitySelectedChildren)
}

// Returns the top-level element that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTopLevelUIElement()
func (o NSMenuItem) AccessibilityTopLevelUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTopLevelUIElement"))
	return objectivec.Object{ID: rv}
}

// Sets the top-level element that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTopLevelUIElement(_:)
func (o NSMenuItem) SetAccessibilityTopLevelUIElement(accessibilityTopLevelUIElement objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTopLevelUIElement:"), accessibilityTopLevelUIElement)
}

// Returns the accessibility element’s visible child accessibility elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleChildren()
func (o NSMenuItem) AccessibilityVisibleChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleChildren"))
	return foundation.NSArrayFromID(rv)
}

// Sets the accessibility element’s visible child accessibility elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleChildren(_:)
func (o NSMenuItem) SetAccessibilityVisibleChildren(accessibilityVisibleChildren objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleChildren:"), accessibilityVisibleChildren)
}

// Returns the child accessibility element with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityApplicationFocusedUIElement()
func (o NSMenuItem) AccessibilityApplicationFocusedUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityApplicationFocusedUIElement"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityApplicationFocusedUIElement(_:)
func (o NSMenuItem) SetAccessibilityApplicationFocusedUIElement(accessibilityApplicationFocusedUIElement objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityApplicationFocusedUIElement:"), accessibilityApplicationFocusedUIElement)
}

// Sets a Boolean value that determines whether the accessibility element has
// the keyboard focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFocused(_:)
func (o NSMenuItem) SetAccessibilityFocused(accessibilityFocused bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocused:"), accessibilityFocused)
}

// Returns the child window with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFocusedWindow()
func (o NSMenuItem) AccessibilityFocusedWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFocusedWindow"))
	return objectivec.Object{ID: rv}
}

// Sets the child window with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFocusedWindow(_:)
func (o NSMenuItem) SetAccessibilityFocusedWindow(accessibilityFocusedWindow objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocusedWindow:"), accessibilityFocusedWindow)
}

// Returns the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySharedFocusElements()
func (o NSMenuItem) AccessibilitySharedFocusElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedFocusElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySharedFocusElements(_:)
func (o NSMenuItem) SetAccessibilitySharedFocusElements(accessibilitySharedFocusElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedFocusElements:"), accessibilitySharedFocusElements)
}

// Returns a Boolean value that determines whether the accessibility element
// must have content for successful submission of a form.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityRequired()
func (o NSMenuItem) IsAccessibilityRequired() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityRequired"))
	return rv
}

// Sets a Boolean value that determines whether the accessibility element must
// have content for successful submission of a form.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRequired(_:)
func (o NSMenuItem) SetAccessibilityRequired(accessibilityRequired bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRequired:"), accessibilityRequired)
}

// Returns the type of interface element that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRole()
func (o NSMenuItem) AccessibilityRole() NSAccessibilityRole {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRole"))
	return NSAccessibilityRole(foundation.NSStringFromID(rv).String())
}

// Sets the type of interface element that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRole(_:)
func (o NSMenuItem) SetAccessibilityRole(accessibilityRole NSAccessibilityRole) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRole:"), objc.String(string(accessibilityRole)))
}

// Returns a localized, human-intelligible description of the accessibility
// element’s role, such as .
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRoleDescription()
func (o NSMenuItem) AccessibilityRoleDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRoleDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the localized, human-intelligible description of the accessibility
// element’s role, such as .
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRoleDescription(_:)
func (o NSMenuItem) SetAccessibilityRoleDescription(accessibilityRoleDescription string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRoleDescription:"), objc.String(accessibilityRoleDescription))
}

// Returns the specialized interface element type that the accessibility
// element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySubrole()
func (o NSMenuItem) AccessibilitySubrole() NSAccessibilitySubrole {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySubrole"))
	return NSAccessibilitySubrole(foundation.NSStringFromID(rv).String())
}

// Sets the specialized interface element type that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySubrole(_:)
func (o NSMenuItem) SetAccessibilitySubrole(accessibilitySubrole NSAccessibilitySubrole) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySubrole:"), objc.String(string(accessibilitySubrole)))
}

// Returns the custom actions of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCustomActions()
func (o NSMenuItem) AccessibilityCustomActions() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCustomActions"))
	return foundation.NSArrayFromID(rv)
}

// Sets the custom actions of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCustomActions(_:)
func (o NSMenuItem) SetAccessibilityCustomActions(accessibilityCustomActions objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomActions:"), accessibilityCustomActions)
}

// Returns the custom rotors of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCustomRotors()
func (o NSMenuItem) AccessibilityCustomRotors() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCustomRotors"))
	return foundation.NSArrayFromID(rv)
}

// Sets the custom rotors of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCustomRotors(_:)
func (o NSMenuItem) SetAccessibilityCustomRotors(accessibilityCustomRotors objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomRotors:"), accessibilityCustomRotors)
}

// Returns the line number that contains the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityInsertionPointLineNumber()
func (o NSMenuItem) AccessibilityInsertionPointLineNumber() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityInsertionPointLineNumber"))
	return rv
}

// Sets the line number that contains the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityInsertionPointLineNumber(_:)
func (o NSMenuItem) SetAccessibilityInsertionPointLineNumber(accessibilityInsertionPointLineNumber int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityInsertionPointLineNumber:"), accessibilityInsertionPointLineNumber)
}

// Returns the number of characters in the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityNumberOfCharacters()
func (o NSMenuItem) AccessibilityNumberOfCharacters() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityNumberOfCharacters"))
	return rv
}

// Sets the number of characters in the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityNumberOfCharacters(_:)
func (o NSMenuItem) SetAccessibilityNumberOfCharacters(accessibilityNumberOfCharacters int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNumberOfCharacters:"), accessibilityNumberOfCharacters)
}

// Returns the placeholder value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPlaceholderValue()
func (o NSMenuItem) AccessibilityPlaceholderValue() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPlaceholderValue"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the placeholder value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityPlaceholderValue(_:)
func (o NSMenuItem) SetAccessibilityPlaceholderValue(accessibilityPlaceholderValue string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPlaceholderValue:"), objc.String(accessibilityPlaceholderValue))
}

// Returns the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedText()
func (o NSMenuItem) AccessibilitySelectedText() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedText"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedText(_:)
func (o NSMenuItem) SetAccessibilitySelectedText(accessibilitySelectedText string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedText:"), objc.String(accessibilitySelectedText))
}

// Returns the range of the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedTextRange()
func (o NSMenuItem) AccessibilitySelectedTextRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySelectedTextRange"))
	return rv
}

// Sets the range of the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedTextRange(_:)
func (o NSMenuItem) SetAccessibilitySelectedTextRange(accessibilitySelectedTextRange foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRange:"), accessibilitySelectedTextRange)
}

// Returns an array of ranges for the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedTextRanges()
func (o NSMenuItem) AccessibilitySelectedTextRanges() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedTextRanges"))
	return foundation.NSArrayFromID(rv)
}

// Sets an array of ranges for the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedTextRanges(_:)
func (o NSMenuItem) SetAccessibilitySelectedTextRanges(accessibilitySelectedTextRanges objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRanges:"), accessibilitySelectedTextRanges)
}

// Returns the range of characters that the accessibility element displays.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySharedCharacterRange()
func (o NSMenuItem) AccessibilitySharedCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySharedCharacterRange"))
	return rv
}

// Sets the range of characters that the accessibility element displays.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySharedCharacterRange(_:)
func (o NSMenuItem) SetAccessibilitySharedCharacterRange(accessibilitySharedCharacterRange foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedCharacterRange:"), accessibilitySharedCharacterRange)
}

// Returns the other elements that share text with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySharedTextUIElements()
func (o NSMenuItem) AccessibilitySharedTextUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedTextUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the other elements that share text with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySharedTextUIElements(_:)
func (o NSMenuItem) SetAccessibilitySharedTextUIElements(accessibilitySharedTextUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedTextUIElements:"), accessibilitySharedTextUIElements)
}

// Returns the range of visible characters in the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleCharacterRange()
func (o NSMenuItem) AccessibilityVisibleCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityVisibleCharacterRange"))
	return rv
}

// Sets the range of visible characters in the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleCharacterRange(_:)
func (o NSMenuItem) SetAccessibilityVisibleCharacterRange(accessibilityVisibleCharacterRange foundation.NSRange) {
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
func (o NSMenuItem) AccessibilityRTFForRange(range_ foundation.NSRange) foundation.INSData {
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

// Returns the activation point for the user interface element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityActivationPoint()
func (o NSMenuItem) AccessibilityActivationPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityActivationPoint"))
	return rv
}

// Sets the activation point for the user interface element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityActivationPoint(_:)
func (o NSMenuItem) SetAccessibilityActivationPoint(accessibilityActivationPoint corefoundation.CGPoint) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityActivationPoint:"), accessibilityActivationPoint)
}

// Returns the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityAlternateUIVisible()
func (o NSMenuItem) IsAccessibilityAlternateUIVisible() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
	return rv
}

// Sets the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityAlternateUIVisible(_:)
func (o NSMenuItem) SetAccessibilityAlternateUIVisible(accessibilityAlternateUIVisible bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAlternateUIVisible:"), accessibilityAlternateUIVisible)
}

// Returns the child accessibility element that represents the window’s
// cancel button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCancelButton()
func (o NSMenuItem) AccessibilityCancelButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCancelButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s cancel
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCancelButton(_:)
func (o NSMenuItem) SetAccessibilityCancelButton(accessibilityCancelButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCancelButton:"), accessibilityCancelButton)
}

// Returns the child accessibility element that represents the window’s
// close button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCloseButton()
func (o NSMenuItem) AccessibilityCloseButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCloseButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s close
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCloseButton(_:)
func (o NSMenuItem) SetAccessibilityCloseButton(accessibilityCloseButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCloseButton:"), accessibilityCloseButton)
}

// Returns the child accessibility element that represents the window’s
// default button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDefaultButton()
func (o NSMenuItem) AccessibilityDefaultButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDefaultButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s default
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDefaultButton(_:)
func (o NSMenuItem) SetAccessibilityDefaultButton(accessibilityDefaultButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDefaultButton:"), accessibilityDefaultButton)
}

// Returns the child accessibility element that represents the window’s
// full-screen button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFullScreenButton()
func (o NSMenuItem) AccessibilityFullScreenButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFullScreenButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s
// full-screen button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFullScreenButton(_:)
func (o NSMenuItem) SetAccessibilityFullScreenButton(accessibilityFullScreenButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFullScreenButton:"), accessibilityFullScreenButton)
}

// Returns the child accessibility element that represents the window’s grow
// area.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityGrowArea()
func (o NSMenuItem) AccessibilityGrowArea() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityGrowArea"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s grow
// area.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityGrowArea(_:)
func (o NSMenuItem) SetAccessibilityGrowArea(accessibilityGrowArea objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityGrowArea:"), accessibilityGrowArea)
}

// Returns a Boolean value that determines whether the window is the app’s
// main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMain()
func (o NSMenuItem) IsAccessibilityMain() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMain"))
	return rv
}

// Sets a Boolean value that determines whether the window is the app’s main
// window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMain(_:)
func (o NSMenuItem) SetAccessibilityMain(accessibilityMain bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMain:"), accessibilityMain)
}

// Returns the child accessibility element that represents the window’s
// minimize button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMinimizeButton()
func (o NSMenuItem) AccessibilityMinimizeButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinimizeButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s
// minimize button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMinimizeButton(_:)
func (o NSMenuItem) SetAccessibilityMinimizeButton(accessibilityMinimizeButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimizeButton:"), accessibilityMinimizeButton)
}

// Returns the Boolean value that determines whether the window is in a
// minimized state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMinimized()
func (o NSMenuItem) IsAccessibilityMinimized() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMinimized"))
	return rv
}

// Sets the Boolean value that determines whether the window is in a minimized
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMinimized(_:)
func (o NSMenuItem) SetAccessibilityMinimized(accessibilityMinimized bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimized:"), accessibilityMinimized)
}

// Returns a Boolean value that determines whether the window is modal.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityModal()
func (o NSMenuItem) IsAccessibilityModal() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityModal"))
	return rv
}

// Sets a Boolean value that determines whether the window is modal.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityModal(_:)
func (o NSMenuItem) SetAccessibilityModal(accessibilityModal bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityModal:"), accessibilityModal)
}

// Returns the child accessibility element that represents the window’s
// proxy icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityProxy()
func (o NSMenuItem) AccessibilityProxy() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityProxy"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s proxy
// icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityProxy(_:)
func (o NSMenuItem) SetAccessibilityProxy(accessibilityProxy objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProxy:"), accessibilityProxy)
}

// Returns the menu currently displaying for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityShownMenu()
func (o NSMenuItem) AccessibilityShownMenu() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityShownMenu"))
	return objectivec.Object{ID: rv}
}

// Sets the menu currently displaying for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityShownMenu(_:)
func (o NSMenuItem) SetAccessibilityShownMenu(accessibilityShownMenu objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityShownMenu:"), accessibilityShownMenu)
}

// Returns the child accessibility element that represents the window’s
// toolbar button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityToolbarButton()
func (o NSMenuItem) AccessibilityToolbarButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityToolbarButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s toolbar
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityToolbarButton(_:)
func (o NSMenuItem) SetAccessibilityToolbarButton(accessibilityToolbarButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityToolbarButton:"), accessibilityToolbarButton)
}

// Returns the window that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityWindow()
func (o NSMenuItem) AccessibilityWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindow"))
	return objectivec.Object{ID: rv}
}

// Sets the window that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityWindow(_:)
func (o NSMenuItem) SetAccessibilityWindow(accessibilityWindow objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindow:"), accessibilityWindow)
}

// Returns the child accessibility element that represents the window’s zoom
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityZoomButton()
func (o NSMenuItem) AccessibilityZoomButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityZoomButton"))
	return objectivec.Object{ID: rv}
}

// Sets the child accessibility element that represents the window’s zoom
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityZoomButton(_:)
func (o NSMenuItem) SetAccessibilityZoomButton(accessibilityZoomButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityZoomButton:"), accessibilityZoomButton)
}

// Returns the icon for the app’s menu bar extra.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityExtrasMenuBar()
func (o NSMenuItem) AccessibilityExtrasMenuBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityExtrasMenuBar"))
	return objectivec.Object{ID: rv}
}

// Sets the icon for the app’s menu bar extra.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityExtrasMenuBar(_:)
func (o NSMenuItem) SetAccessibilityExtrasMenuBar(accessibilityExtrasMenuBar objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExtrasMenuBar:"), accessibilityExtrasMenuBar)
}

// Returns a Boolean value that determines whether the app is the frontmost
// app.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityFrontmost()
func (o NSMenuItem) IsAccessibilityFrontmost() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFrontmost"))
	return rv
}

// Sets a Boolean value that determines whether the app is the frontmost app.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFrontmost(_:)
func (o NSMenuItem) SetAccessibilityFrontmost(accessibilityFrontmost bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrontmost:"), accessibilityFrontmost)
}

// Returns a Boolean value that determines whether the app is in a hidden
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityHidden()
func (o NSMenuItem) IsAccessibilityHidden() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityHidden"))
	return rv
}

// Sets a Boolean value that determines whether the app is in a hidden state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHidden(_:)
func (o NSMenuItem) SetAccessibilityHidden(accessibilityHidden bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHidden:"), accessibilityHidden)
}

// Returns the app’s main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMainWindow()
func (o NSMenuItem) AccessibilityMainWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMainWindow"))
	return objectivec.Object{ID: rv}
}

// Sets the app’s main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMainWindow(_:)
func (o NSMenuItem) SetAccessibilityMainWindow(accessibilityMainWindow objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMainWindow:"), accessibilityMainWindow)
}

// Returns the app’s menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMenuBar()
func (o NSMenuItem) AccessibilityMenuBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMenuBar"))
	return objectivec.Object{ID: rv}
}

// Sets the app’s menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMenuBar(_:)
func (o NSMenuItem) SetAccessibilityMenuBar(accessibilityMenuBar objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMenuBar:"), accessibilityMenuBar)
}

// Returns an array that contains all the app’s windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityWindows()
func (o NSMenuItem) AccessibilityWindows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindows"))
	return foundation.NSArrayFromID(rv)
}

// Sets the array that contains all the app’s windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityWindows(_:)
func (o NSMenuItem) SetAccessibilityWindows(accessibilityWindows objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindows:"), accessibilityWindows)
}

// Returns the number of columns in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnCount()
func (o NSMenuItem) AccessibilityColumnCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityColumnCount"))
	return rv
}

// Sets the number of columns in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnCount(_:)
func (o NSMenuItem) SetAccessibilityColumnCount(accessibilityColumnCount int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnCount:"), accessibilityColumnCount)
}

// Returns a Boolean value that determines whether the accessibility
// element’s grid is in row major order or in column major order.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityOrderedByRow()
func (o NSMenuItem) IsAccessibilityOrderedByRow() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityOrderedByRow"))
	return rv
}

// Sets a Boolean value that determines whether the element’s grid is in row
// major order or in column major order.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityOrderedByRow(_:)
func (o NSMenuItem) SetAccessibilityOrderedByRow(accessibilityOrderedByRow bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrderedByRow:"), accessibilityOrderedByRow)
}

// Returns the number of rows in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRowCount()
func (o NSMenuItem) AccessibilityRowCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityRowCount"))
	return rv
}

// Sets the number of rows in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRowCount(_:)
func (o NSMenuItem) SetAccessibilityRowCount(accessibilityRowCount int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowCount:"), accessibilityRowCount)
}

// Returns the horizontal scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHorizontalScrollBar()
func (o NSMenuItem) AccessibilityHorizontalScrollBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalScrollBar"))
	return objectivec.Object{ID: rv}
}

// Sets the horizontal scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHorizontalScrollBar(_:)
func (o NSMenuItem) SetAccessibilityHorizontalScrollBar(accessibilityHorizontalScrollBar objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalScrollBar:"), accessibilityHorizontalScrollBar)
}

// Returns the vertical scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVerticalScrollBar()
func (o NSMenuItem) AccessibilityVerticalScrollBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalScrollBar"))
	return objectivec.Object{ID: rv}
}

// Sets the vertical scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVerticalScrollBar(_:)
func (o NSMenuItem) SetAccessibilityVerticalScrollBar(accessibilityVerticalScrollBar objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalScrollBar:"), accessibilityVerticalScrollBar)
}

// Returns the column header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnHeaderUIElements()
func (o NSMenuItem) AccessibilityColumnHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the column header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnHeaderUIElements(_:)
func (o NSMenuItem) SetAccessibilityColumnHeaderUIElements(accessibilityColumnHeaderUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnHeaderUIElements:"), accessibilityColumnHeaderUIElements)
}

// Returns the column accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumns()
func (o NSMenuItem) AccessibilityColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumns"))
	return foundation.NSArrayFromID(rv)
}

// Sets the column accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumns(_:)
func (o NSMenuItem) SetAccessibilityColumns(accessibilityColumns objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumns:"), accessibilityColumns)
}

// Returns the column titles for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnTitles()
func (o NSMenuItem) AccessibilityColumnTitles() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnTitles"))
	return foundation.NSArrayFromID(rv)
}

// Sets the column titles for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnTitles(_:)
func (o NSMenuItem) SetAccessibilityColumnTitles(accessibilityColumnTitles objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnTitles:"), accessibilityColumnTitles)
}

// Returns a Boolean value that determines whether the accessibility element
// is in an expanded state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityExpanded()
func (o NSMenuItem) IsAccessibilityExpanded() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityExpanded"))
	return rv
}

// Sets a Boolean value that determines whether accessibility element is in an
// expanded state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityExpanded(_:)
func (o NSMenuItem) SetAccessibilityExpanded(accessibilityExpanded bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExpanded:"), accessibilityExpanded)
}

// Returns the header for the table view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHeader()
func (o NSMenuItem) AccessibilityHeader() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHeader"))
	return objectivec.Object{ID: rv}
}

// Sets the header for the table view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHeader(_:)
func (o NSMenuItem) SetAccessibilityHeader(accessibilityHeader objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHeader:"), accessibilityHeader)
}

// Returns the index of the row or column that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityIndex()
func (o NSMenuItem) AccessibilityIndex() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityIndex"))
	return rv
}

// Sets the index of the row or column that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityIndex(_:)
func (o NSMenuItem) SetAccessibilityIndex(accessibilityIndex int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIndex:"), accessibilityIndex)
}

// Returns the row header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRowHeaderUIElements()
func (o NSMenuItem) AccessibilityRowHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRowHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the row header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRowHeaderUIElements(_:)
func (o NSMenuItem) SetAccessibilityRowHeaderUIElements(accessibilityRowHeaderUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowHeaderUIElements:"), accessibilityRowHeaderUIElements)
}

// Returns the row accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRows()
func (o NSMenuItem) AccessibilityRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRows"))
	return foundation.NSArrayFromID(rv)
}

// Sets the row accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRows(_:)
func (o NSMenuItem) SetAccessibilityRows(accessibilityRows objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRows:"), accessibilityRows)
}

// Returns the currently selected columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedColumns()
func (o NSMenuItem) AccessibilitySelectedColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedColumns"))
	return foundation.NSArrayFromID(rv)
}

// Sets the currently selected columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedColumns(_:)
func (o NSMenuItem) SetAccessibilitySelectedColumns(accessibilitySelectedColumns objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedColumns:"), accessibilitySelectedColumns)
}

// Returns the currently selected rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedRows()
func (o NSMenuItem) AccessibilitySelectedRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedRows"))
	return foundation.NSArrayFromID(rv)
}

// Sets the currently selected rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedRows(_:)
func (o NSMenuItem) SetAccessibilitySelectedRows(accessibilitySelectedRows objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedRows:"), accessibilitySelectedRows)
}

// Returns the accessibility element’s sort direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySortDirection()
func (o NSMenuItem) AccessibilitySortDirection() NSAccessibilitySortDirection {
	rv := objc.Send[NSAccessibilitySortDirection](o.ID, objc.Sel("accessibilitySortDirection"))
	return rv
}

// Sets the accessibility element’s sort direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySortDirection(_:)
func (o NSMenuItem) SetAccessibilitySortDirection(accessibilitySortDirection NSAccessibilitySortDirection) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySortDirection:"), accessibilitySortDirection)
}

// Returns the visible columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleColumns()
func (o NSMenuItem) AccessibilityVisibleColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleColumns"))
	return foundation.NSArrayFromID(rv)
}

// Sets the visible columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleColumns(_:)
func (o NSMenuItem) SetAccessibilityVisibleColumns(accessibilityVisibleColumns objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleColumns:"), accessibilityVisibleColumns)
}

// Returns the visible rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleRows()
func (o NSMenuItem) AccessibilityVisibleRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleRows"))
	return foundation.NSArrayFromID(rv)
}

// Sets the visible rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleRows(_:)
func (o NSMenuItem) SetAccessibilityVisibleRows(accessibilityVisibleRows objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleRows:"), accessibilityVisibleRows)
}

// Returns a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityDisclosed()
func (o NSMenuItem) IsAccessibilityDisclosed() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityDisclosed"))
	return rv
}

// Sets a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosed(_:)
func (o NSMenuItem) SetAccessibilityDisclosed(accessibilityDisclosed bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosed:"), accessibilityDisclosed)
}

// Returns the row disclosing the current row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDisclosedByRow()
func (o NSMenuItem) AccessibilityDisclosedByRow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedByRow"))
	return objectivec.Object{ID: rv}
}

// Sets the row disclosing the current row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosedByRow(_:)
func (o NSMenuItem) SetAccessibilityDisclosedByRow(accessibilityDisclosedByRow objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedByRow:"), accessibilityDisclosedByRow)
}

// Returns the rows that the current row discloses.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDisclosedRows()
func (o NSMenuItem) AccessibilityDisclosedRows() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedRows"))
	return objectivec.Object{ID: rv}
}

// Sets the rows that the current row discloses.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosedRows(_:)
func (o NSMenuItem) SetAccessibilityDisclosedRows(accessibilityDisclosedRows objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedRows:"), accessibilityDisclosedRows)
}

// Returns the indention level for the row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDisclosureLevel()
func (o NSMenuItem) AccessibilityDisclosureLevel() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityDisclosureLevel"))
	return rv
}

// Sets the indention level for the row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosureLevel(_:)
func (o NSMenuItem) SetAccessibilityDisclosureLevel(accessibilityDisclosureLevel int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosureLevel:"), accessibilityDisclosureLevel)
}

// Returns the column index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnIndexRange()
func (o NSMenuItem) AccessibilityColumnIndexRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityColumnIndexRange"))
	return rv
}

// Sets the column index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnIndexRange(_:)
func (o NSMenuItem) SetAccessibilityColumnIndexRange(accessibilityColumnIndexRange foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnIndexRange:"), accessibilityColumnIndexRange)
}

// Returns the row index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRowIndexRange()
func (o NSMenuItem) AccessibilityRowIndexRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRowIndexRange"))
	return rv
}

// Sets the row index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRowIndexRange(_:)
func (o NSMenuItem) SetAccessibilityRowIndexRange(accessibilityRowIndexRange foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowIndexRange:"), accessibilityRowIndexRange)
}

// Returns the currently selected cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedCells()
func (o NSMenuItem) AccessibilitySelectedCells() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedCells"))
	return foundation.NSArrayFromID(rv)
}

// Sets the currently selected cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedCells(_:)
func (o NSMenuItem) SetAccessibilitySelectedCells(accessibilitySelectedCells objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedCells:"), accessibilitySelectedCells)
}

// Returns the visible cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleCells()
func (o NSMenuItem) AccessibilityVisibleCells() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleCells"))
	return foundation.NSArrayFromID(rv)
}

// Sets the visible cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleCells(_:)
func (o NSMenuItem) SetAccessibilityVisibleCells(accessibilityVisibleCells objectivec.IObject) {
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
func (o NSMenuItem) AccessibilityCellForColumnRow(column int, row int) objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCellForColumn:row:"), column, row)
	return objectivec.Object{ID: rv}
}

// Returns the drag handle elements for the layout item element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHandles()
func (o NSMenuItem) AccessibilityHandles() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHandles"))
	return foundation.NSArrayFromID(rv)
}

// Sets the drag handle accessibility elements for the layout item element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHandles(_:)
func (o NSMenuItem) SetAccessibilityHandles(accessibilityHandles objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHandles:"), accessibilityHandles)
}

// Returns the units that the layout area uses for horizontal values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHorizontalUnits()
func (o NSMenuItem) AccessibilityHorizontalUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityHorizontalUnits"))
	return rv
}

// Sets the units that the layout area uses for horizontal values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHorizontalUnits(_:)
func (o NSMenuItem) SetAccessibilityHorizontalUnits(accessibilityHorizontalUnits NSAccessibilityUnits) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnits:"), accessibilityHorizontalUnits)
}

// Returns the description of the layout area’s horizontal units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHorizontalUnitDescription()
func (o NSMenuItem) AccessibilityHorizontalUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the description of the layout area’s horizontal units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHorizontalUnitDescription(_:)
func (o NSMenuItem) SetAccessibilityHorizontalUnitDescription(accessibilityHorizontalUnitDescription string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnitDescription:"), objc.String(accessibilityHorizontalUnitDescription))
}

// Returns the units that the layout area uses for vertical values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVerticalUnits()
func (o NSMenuItem) AccessibilityVerticalUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityVerticalUnits"))
	return rv
}

// Sets the units that the layout area uses for vertical values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVerticalUnits(_:)
func (o NSMenuItem) SetAccessibilityVerticalUnits(accessibilityVerticalUnits NSAccessibilityUnits) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnits:"), accessibilityVerticalUnits)
}

// Returns the description of the layout area’s vertical units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVerticalUnitDescription()
func (o NSMenuItem) AccessibilityVerticalUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the description of the layout area’s vertical units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVerticalUnitDescription(_:)
func (o NSMenuItem) SetAccessibilityVerticalUnitDescription(accessibilityVerticalUnitDescription string) {
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

// Returns the allowed values for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityAllowedValues()
func (o NSMenuItem) AccessibilityAllowedValues() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAllowedValues"))
	return foundation.NSArrayFromID(rv)
}

// Sets the allowed values for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityAllowedValues(_:)
func (o NSMenuItem) SetAccessibilityAllowedValues(accessibilityAllowedValues objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAllowedValues:"), accessibilityAllowedValues)
}

// Returns the child label elements for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLabelUIElements()
func (o NSMenuItem) AccessibilityLabelUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabelUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the child label elements for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLabelUIElements(_:)
func (o NSMenuItem) SetAccessibilityLabelUIElements(accessibilityLabelUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelUIElements:"), accessibilityLabelUIElements)
}

// Returns the value of the label accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLabelValue()
func (o NSMenuItem) AccessibilityLabelValue() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("accessibilityLabelValue"))
	return rv
}

// Sets the value of the label accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLabelValue(_:)
func (o NSMenuItem) SetAccessibilityLabelValue(accessibilityLabelValue float32) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelValue:"), accessibilityLabelValue)
}

// Returns the contents that follow the divider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityNextContents()
func (o NSMenuItem) AccessibilityNextContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityNextContents"))
	return foundation.NSArrayFromID(rv)
}

// Sets the contents that follow the divider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityNextContents(_:)
func (o NSMenuItem) SetAccessibilityNextContents(accessibilityNextContents objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNextContents:"), accessibilityNextContents)
}

// Returns the contents that precede the divider accessibility element.
//
// # Return Value
//
// Sets the contents preceding this divider element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPreviousContents()
func (o NSMenuItem) AccessibilityPreviousContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPreviousContents"))
	return foundation.NSArrayFromID(rv)
}

// Sets the contents that precede the divider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityPreviousContents(_:)
func (o NSMenuItem) SetAccessibilityPreviousContents(accessibilityPreviousContents objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPreviousContents:"), accessibilityPreviousContents)
}

// Returns an array that contains the views and splitter bar from the split
// view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySplitters()
func (o NSMenuItem) AccessibilitySplitters() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySplitters"))
	return foundation.NSArrayFromID(rv)
}

// Sets the array that contains the views and splitter bar from the split
// view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySplitters(_:)
func (o NSMenuItem) SetAccessibilitySplitters(accessibilitySplitters objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySplitters:"), accessibilitySplitters)
}

// Returns the overflow button for the toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityOverflowButton()
func (o NSMenuItem) AccessibilityOverflowButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityOverflowButton"))
	return objectivec.Object{ID: rv}
}

// Sets the overflow button for the toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityOverflowButton(_:)
func (o NSMenuItem) SetAccessibilityOverflowButton(accessibilityOverflowButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOverflowButton:"), accessibilityOverflowButton)
}

// Returns the tab accessibility elements for the tab view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTabs()
func (o NSMenuItem) AccessibilityTabs() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTabs"))
	return foundation.NSArrayFromID(rv)
}

// Sets the tab accessibility elements for the tab view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTabs(_:)
func (o NSMenuItem) SetAccessibilityTabs(accessibilityTabs objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTabs:"), accessibilityTabs)
}

// Returns the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerGroupUIElement()
func (o NSMenuItem) AccessibilityMarkerGroupUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerGroupUIElement"))
	return objectivec.Object{ID: rv}
}

// Sets the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerGroupUIElement(_:)
func (o NSMenuItem) SetAccessibilityMarkerGroupUIElement(accessibilityMarkerGroupUIElement objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerGroupUIElement:"), accessibilityMarkerGroupUIElement)
}

// Returns the human-readable description of the marker type.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerTypeDescription()
func (o NSMenuItem) AccessibilityMarkerTypeDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerTypeDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the human-readable description of the marker type.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerTypeDescription(_:)
func (o NSMenuItem) SetAccessibilityMarkerTypeDescription(accessibilityMarkerTypeDescription string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerTypeDescription:"), objc.String(accessibilityMarkerTypeDescription))
}

// Returns the array of marker accessibility elements for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerUIElements()
func (o NSMenuItem) AccessibilityMarkerUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the array of marker accessibility elements for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerUIElements(_:)
func (o NSMenuItem) SetAccessibilityMarkerUIElements(accessibilityMarkerUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerUIElements:"), accessibilityMarkerUIElements)
}

// Returns the marker values for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerValues()
func (o NSMenuItem) AccessibilityMarkerValues() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerValues"))
	return objectivec.Object{ID: rv}
}

// Sets the marker values for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerValues(_:)
func (o NSMenuItem) SetAccessibilityMarkerValues(accessibilityMarkerValues objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerValues:"), accessibilityMarkerValues)
}

// Returns the type of markers for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRulerMarkerType()
func (o NSMenuItem) AccessibilityRulerMarkerType() NSAccessibilityRulerMarkerType {
	rv := objc.Send[NSAccessibilityRulerMarkerType](o.ID, objc.Sel("accessibilityRulerMarkerType"))
	return rv
}

// Sets the type of markers for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRulerMarkerType(_:)
func (o NSMenuItem) SetAccessibilityRulerMarkerType(accessibilityRulerMarkerType NSAccessibilityRulerMarkerType) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRulerMarkerType:"), accessibilityRulerMarkerType)
}

// Returns the units for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityUnits()
func (o NSMenuItem) AccessibilityUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityUnits"))
	return rv
}

// Sets the units used for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityUnits(_:)
func (o NSMenuItem) SetAccessibilityUnits(accessibilityUnits NSAccessibilityUnits) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnits:"), accessibilityUnits)
}

// Returns the human-readable description of the ruler’s units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityUnitDescription()
func (o NSMenuItem) AccessibilityUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the human-readable description of the ruler’s units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityUnitDescription(_:)
func (o NSMenuItem) SetAccessibilityUnitDescription(accessibilityUnitDescription string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnitDescription:"), objc.String(accessibilityUnitDescription))
}

// Returns the URL for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDocument()
func (o NSMenuItem) AccessibilityDocument() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDocument"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the URL for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDocument(_:)
func (o NSMenuItem) SetAccessibilityDocument(accessibilityDocument string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDocument:"), objc.String(accessibilityDocument))
}

// Returns a Boolean value that indicates whether the accessibility element is
// in an edited state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEdited()
func (o NSMenuItem) IsAccessibilityEdited() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEdited"))
	return rv
}

// Sets a Boolean value that indicates whether the accessibility element is in
// an edited state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityEdited(_:)
func (o NSMenuItem) SetAccessibilityEdited(accessibilityEdited bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEdited:"), accessibilityEdited)
}

// Returns the filename for the file that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFilename()
func (o NSMenuItem) AccessibilityFilename() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFilename"))
	return foundation.NSStringFromID(rv).String()
}

// Sets the filename for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFilename(_:)
func (o NSMenuItem) SetAccessibilityFilename(accessibilityFilename string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFilename:"), objc.String(accessibilityFilename))
}

// Returns the elements that have links with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLinkedUIElements()
func (o NSMenuItem) AccessibilityLinkedUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLinkedUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the elements that have links with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLinkedUIElements(_:)
func (o NSMenuItem) SetAccessibilityLinkedUIElements(accessibilityLinkedUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLinkedUIElements:"), accessibilityLinkedUIElements)
}

// Returns the list of elements that the accessibility element is a title for.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityServesAsTitleForUIElements()
func (o NSMenuItem) AccessibilityServesAsTitleForUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityServesAsTitleForUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Sets the list of elements that the accessibility element is a title for.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityServesAsTitleForUIElements(_:)
func (o NSMenuItem) SetAccessibilityServesAsTitleForUIElements(accessibilityServesAsTitleForUIElements objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityServesAsTitleForUIElements:"), accessibilityServesAsTitleForUIElements)
}

// Returns the static text element that represents the accessibility
// element’s title.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTitleUIElement()
func (o NSMenuItem) AccessibilityTitleUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitleUIElement"))
	return objectivec.Object{ID: rv}
}

// Sets the static text element that represents the accessibility element’s
// title.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTitleUIElement(_:)
func (o NSMenuItem) SetAccessibilityTitleUIElement(accessibilityTitleUIElement objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitleUIElement:"), accessibilityTitleUIElement)
}

// Returns the clear button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityClearButton()
func (o NSMenuItem) AccessibilityClearButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityClearButton"))
	return objectivec.Object{ID: rv}
}

// Sets the clear button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityClearButton(_:)
func (o NSMenuItem) SetAccessibilityClearButton(accessibilityClearButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityClearButton:"), accessibilityClearButton)
}

// Returns the search button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySearchButton()
func (o NSMenuItem) AccessibilitySearchButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchButton"))
	return objectivec.Object{ID: rv}
}

// Sets the search button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySearchButton(_:)
func (o NSMenuItem) SetAccessibilitySearchButton(accessibilitySearchButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchButton:"), accessibilitySearchButton)
}

// Returns the search menu for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySearchMenu()
func (o NSMenuItem) AccessibilitySearchMenu() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchMenu"))
	return objectivec.Object{ID: rv}
}

// Sets the search menu for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySearchMenu(_:)
func (o NSMenuItem) SetAccessibilitySearchMenu(accessibilitySearchMenu objectivec.IObject) {
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

// Returns the increment button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityIncrementButton()
func (o NSMenuItem) AccessibilityIncrementButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIncrementButton"))
	return objectivec.Object{ID: rv}
}

// Sets the increment button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityIncrementButton(_:)
func (o NSMenuItem) SetAccessibilityIncrementButton(accessibilityIncrementButton objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIncrementButton:"), accessibilityIncrementButton)
}

// Returns the decrement button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDecrementButton()
func (o NSMenuItem) AccessibilityDecrementButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDecrementButton"))
	return objectivec.Object{ID: rv}
}

// Sets the decrement button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDecrementButton(_:)
func (o NSMenuItem) SetAccessibilityDecrementButton(accessibilityDecrementButton objectivec.IObject) {
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

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityAttributedUserInputLabels()
func (o NSMenuItem) AccessibilityAttributedUserInputLabels() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAttributedUserInputLabels"))
	return foundation.NSArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityUserInputLabels()
func (o NSMenuItem) AccessibilityUserInputLabels() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUserInputLabels"))
	return foundation.NSArrayFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityAttributedUserInputLabels(_:)
func (o NSMenuItem) SetAccessibilityAttributedUserInputLabels(accessibilityAttributedUserInputLabels objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAttributedUserInputLabels:"), accessibilityAttributedUserInputLabels)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityUserInputLabels(_:)
func (o NSMenuItem) SetAccessibilityUserInputLabels(accessibilityUserInputLabels objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUserInputLabels:"), accessibilityUserInputLabels)
}

// Protocol methods for NSUserInterfaceItemIdentification

// Protocol methods for NSValidatedUserInterfaceItem
