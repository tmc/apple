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

	// Returns the activation point for the user interface element.
	AccessibilityActivationPoint() corefoundation.CGPoint
	// Returns the allowed values for the slider accessibility element.
	AccessibilityAllowedValues() []foundation.NSNumber
	// Returns the child accessibility element with the current focus.
	AccessibilityApplicationFocusedUIElement() objectivec.IObject
	// Returns the attributed substring for the specified range of characters.
	AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString
	AccessibilityAttributedUserInputLabels() []foundation.NSAttributedString
	// Returns the child accessibility element that represents the window’s cancel button.
	AccessibilityCancelButton() objectivec.IObject
	// Returns the cell at the specified column and row.
	AccessibilityCellForColumnRow(column int, row int) objectivec.IObject
	// Returns the child accessibility elements in the accessibility hierarchy.
	AccessibilityChildren() foundation.INSArray
	// Returns the array of child accessibility elements in order for linear navigation.
	AccessibilityChildrenInNavigationOrder() []objectivec.IObject
	// Returns the clear button for the search field.
	AccessibilityClearButton() objectivec.IObject
	// Returns the child accessibility element that represents the window’s close button.
	AccessibilityCloseButton() objectivec.IObject
	// Returns the number of columns in the accessibility element’s grid.
	AccessibilityColumnCount() int
	// Returns the column header accessibility elements for the table or outline.
	AccessibilityColumnHeaderUIElements() foundation.INSArray
	// Returns the column index range of the cell.
	AccessibilityColumnIndexRange() foundation.NSRange
	// Returns the column titles for the accessibility element.
	AccessibilityColumnTitles() foundation.INSArray
	// Returns the column accessibility elements for the table or outline.
	AccessibilityColumns() foundation.INSArray
	// Returns the contents of the current accessibility element.
	AccessibilityContents() foundation.INSArray
	// Returns the critical value for the level indicator.
	AccessibilityCriticalValue() objectivec.IObject
	// Returns the custom actions of the current accessibility element.
	AccessibilityCustomActions() []NSAccessibilityCustomAction
	// Returns the custom rotors of the current accessibility element.
	AccessibilityCustomRotors() []NSAccessibilityCustomRotor
	// Returns the decrement button for the stepper accessibility element.
	AccessibilityDecrementButton() objectivec.IObject
	// Returns the child accessibility element that represents the window’s default button.
	AccessibilityDefaultButton() objectivec.IObject
	// Returns the row disclosing the current row.
	AccessibilityDisclosedByRow() objectivec.IObject
	// Returns the rows that the current row discloses.
	AccessibilityDisclosedRows() objectivec.IObject
	// Returns the indention level for the row.
	AccessibilityDisclosureLevel() int
	// Returns the URL for the file that the accessibility element represents.
	AccessibilityDocument() string
	// Returns the icon for the app’s menu bar extra.
	AccessibilityExtrasMenuBar() objectivec.IObject
	// Returns the filename for the file that the accessibility element represents.
	AccessibilityFilename() string
	// Returns the child window with the current focus.
	AccessibilityFocusedWindow() objectivec.IObject
	// Returns the accessibility element’s frame in screen coordinates.
	AccessibilityFrame() corefoundation.CGRect
	// Returns the rectangle that encloses the specified range of characters.
	AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect
	// Returns the child accessibility element that represents the window’s full-screen button.
	AccessibilityFullScreenButton() objectivec.IObject
	// Returns the child accessibility element that represents the window’s grow area.
	AccessibilityGrowArea() objectivec.IObject
	// Returns the drag handle elements for the layout item element.
	AccessibilityHandles() foundation.INSArray
	// Returns the header for the table view.
	AccessibilityHeader() objectivec.IObject
	// Returns the help text for the accessibility element.
	AccessibilityHelp() string
	// Returns the horizontal scroll bar for the scroll view.
	AccessibilityHorizontalScrollBar() objectivec.IObject
	// Returns the description of the layout area’s horizontal units.
	AccessibilityHorizontalUnitDescription() string
	// Returns the units that the layout area uses for horizontal values.
	AccessibilityHorizontalUnits() NSAccessibilityUnits
	// Returns the accessibility element’s identity.
	AccessibilityIdentifier() string
	// Returns the increment button for the stepper accessibility element.
	AccessibilityIncrementButton() objectivec.IObject
	// Returns the index of the row or column that the accessibility element represents.
	AccessibilityIndex() int
	// Returns the line number that contains the insertion point.
	AccessibilityInsertionPointLineNumber() int
	// Returns a short description of the accessibility element.
	AccessibilityLabel() string
	// Returns the child label elements for the slider accessibility element.
	AccessibilityLabelUIElements() foundation.INSArray
	// Returns the value of the label accessibility element.
	AccessibilityLabelValue() float32
	// Converts the provided point in screen coordinates to a point in the layout area’s coordinate system.
	AccessibilityLayoutPointForScreenPoint(point corefoundation.CGPoint) corefoundation.CGPoint
	// Converts the provided size in screen coordinates to a size in the layout area’s coordinate system.
	AccessibilityLayoutSizeForScreenSize(size corefoundation.CGSize) corefoundation.CGSize
	// Returns the line number for the line that contains the specified character index.
	AccessibilityLineForIndex(index int) int
	// Returns the elements that have links with the accessibility element.
	AccessibilityLinkedUIElements() foundation.INSArray
	// Returns the app’s main window.
	AccessibilityMainWindow() objectivec.IObject
	// Returns the user interface element that functions as a marker group for the ruler accessibility element.
	AccessibilityMarkerGroupUIElement() objectivec.IObject
	// Returns the human-readable description of the marker type.
	AccessibilityMarkerTypeDescription() string
	// Returns the array of marker accessibility elements for the ruler.
	AccessibilityMarkerUIElements() foundation.INSArray
	// Returns the marker values for the ruler.
	AccessibilityMarkerValues() objectivec.IObject
	// Returns the maximum value for the accessibility element.
	AccessibilityMaxValue() objectivec.IObject
	// Returns the app’s menu bar.
	AccessibilityMenuBar() objectivec.IObject
	// Returns the minimum value for the accessibility element.
	AccessibilityMinValue() objectivec.IObject
	// Returns the child accessibility element that represents the window’s minimize button.
	AccessibilityMinimizeButton() objectivec.IObject
	// Returns the contents that follow the divider accessibility element.
	AccessibilityNextContents() foundation.INSArray
	// Returns the number of characters in the text.
	AccessibilityNumberOfCharacters() int
	// Returns the orientation of the accessibility element.
	AccessibilityOrientation() NSAccessibilityOrientation
	// Returns the overflow button for the toolbar.
	AccessibilityOverflowButton() objectivec.IObject
	// Returns the accessibility element’s parent in the accessibility hierarchy.
	AccessibilityParent() objectivec.IObject
	// Cancels the current operation.
	AccessibilityPerformCancel() bool
	// Simulates pressing Return in the accessibility element.
	AccessibilityPerformConfirm() bool
	// Decrements the accessibility element’s value.
	AccessibilityPerformDecrement() bool
	// Deletes the accessibility element’s value.
	AccessibilityPerformDelete() bool
	// Increments the accessibility element’s value.
	AccessibilityPerformIncrement() bool
	// Selects the accessibility element.
	AccessibilityPerformPick() bool
	// Simulates clicking the accessibility element.
	AccessibilityPerformPress() bool
	// Brings the window to the front.
	AccessibilityPerformRaise() bool
	// Displays the accessibility element’s alternative UI.
	AccessibilityPerformShowAlternateUI() bool
	// Returns to the accessibility element’s original UI.
	AccessibilityPerformShowDefaultUI() bool
	// Displays the menu accessibility element.
	AccessibilityPerformShowMenu() bool
	// Returns the placeholder value for the accessibility element.
	AccessibilityPlaceholderValue() string
	// Returns the contents that precede the divider accessibility element.
	AccessibilityPreviousContents() foundation.INSArray
	// Returns the child accessibility element that represents the window’s proxy icon.
	AccessibilityProxy() objectivec.IObject
	// Returns the rich text format (RTF) data that describes the specified range of characters.
	AccessibilityRTFForRange(range_ foundation.NSRange) foundation.NSData
	// Returns the range of characters for the glyph that includes the specified character.
	AccessibilityRangeForIndex(index int) foundation.NSRange
	// Returns the range of characters in the specified line.
	AccessibilityRangeForLine(line int) foundation.NSRange
	// Returns the range of characters for the glyph at the specified point.
	AccessibilityRangeForPosition(point corefoundation.CGPoint) foundation.NSRange
	// Returns the type of interface element that the accessibility element represents.
	AccessibilityRole() NSAccessibilityRole
	// Returns a localized, human-intelligible description of the accessibility element’s role, such as radio button.
	AccessibilityRoleDescription() string
	// Returns the number of rows in the accessibility element’s grid.
	AccessibilityRowCount() int
	// Returns the row header accessibility elements for the table or outline.
	AccessibilityRowHeaderUIElements() foundation.INSArray
	// Returns the row index range of the cell.
	AccessibilityRowIndexRange() foundation.NSRange
	// Returns the row accessibility elements for the table or outline.
	AccessibilityRows() foundation.INSArray
	// Returns the type of markers for the ruler.
	AccessibilityRulerMarkerType() NSAccessibilityRulerMarkerType
	// Converts the provided point in the layout area’s coordinates to a point in the screen’s coordinate system.
	AccessibilityScreenPointForLayoutPoint(point corefoundation.CGPoint) corefoundation.CGPoint
	// Converts the provided size in the layout area’s coordinates to a size in the screen’s coordinate system.
	AccessibilityScreenSizeForLayoutSize(size corefoundation.CGSize) corefoundation.CGSize
	// Returns the search button for the search field.
	AccessibilitySearchButton() objectivec.IObject
	// Returns the search menu for the search field.
	AccessibilitySearchMenu() objectivec.IObject
	// Returns the currently selected cells for the table.
	AccessibilitySelectedCells() foundation.INSArray
	// Returns the accessibility element’s currently selected children.
	AccessibilitySelectedChildren() foundation.INSArray
	// Returns the currently selected columns for the table or outline.
	AccessibilitySelectedColumns() foundation.INSArray
	// Returns the currently selected rows for the table or outline.
	AccessibilitySelectedRows() foundation.INSArray
	// Returns the currently selected text.
	AccessibilitySelectedText() string
	// Returns the range of the currently selected text.
	AccessibilitySelectedTextRange() foundation.NSRange
	// Returns an array of ranges for the currently selected text.
	AccessibilitySelectedTextRanges() []foundation.NSValue
	// Returns the list of elements that the accessibility element is a title for.
	AccessibilityServesAsTitleForUIElements() foundation.INSArray
	// Returns the range of characters that the accessibility element displays.
	AccessibilitySharedCharacterRange() foundation.NSRange
	// Returns the array of elements that shares the keyboard focus with the accessibility element.
	AccessibilitySharedFocusElements() foundation.INSArray
	// Returns the other elements that share text with the accessibility element.
	AccessibilitySharedTextUIElements() foundation.INSArray
	// Returns the menu currently displaying for the accessibility element.
	AccessibilityShownMenu() objectivec.IObject
	// Returns the accessibility element’s sort direction.
	AccessibilitySortDirection() NSAccessibilitySortDirection
	// Returns an array that contains the views and splitter bar from the split view.
	AccessibilitySplitters() foundation.INSArray
	// Returns the substring for the specified range.
	AccessibilityStringForRange(range_ foundation.NSRange) string
	// Returns a range of characters that all have the same style as the specified character.
	AccessibilityStyleRangeForIndex(index int) foundation.NSRange
	// Returns the specialized interface element type that the accessibility element represents.
	AccessibilitySubrole() NSAccessibilitySubrole
	// Returns the tab accessibility elements for the tab view.
	AccessibilityTabs() foundation.INSArray
	// Returns the title of the accessibility element—for example, a button’s visible text.
	AccessibilityTitle() string
	// Returns the static text element that represents the accessibility element’s title.
	AccessibilityTitleUIElement() objectivec.IObject
	// Returns the child accessibility element that represents the window’s toolbar button.
	AccessibilityToolbarButton() objectivec.IObject
	// Returns the top-level element that contains the accessibility element.
	AccessibilityTopLevelUIElement() objectivec.IObject
	// Returns the URL for the accessibility element.
	AccessibilityURL() foundation.NSURL
	// Returns the human-readable description of the ruler’s units.
	AccessibilityUnitDescription() string
	// Returns the units for the ruler.
	AccessibilityUnits() NSAccessibilityUnits
	AccessibilityUserInputLabels() []string
	// Returns the accessibility element’s value.
	AccessibilityValue() objectivec.IObject
	// Returns the human-readable description of the accessibility element’s value.
	AccessibilityValueDescription() string
	// Returns the vertical scroll bar for the scroll view.
	AccessibilityVerticalScrollBar() objectivec.IObject
	// Returns the description of the layout area’s vertical units.
	AccessibilityVerticalUnitDescription() string
	// Returns the units that the layout area uses for vertical values.
	AccessibilityVerticalUnits() NSAccessibilityUnits
	// Returns the visible cells for the table.
	AccessibilityVisibleCells() foundation.INSArray
	// Returns the range of visible characters in the document.
	AccessibilityVisibleCharacterRange() foundation.NSRange
	// Returns the accessibility element’s visible child accessibility elements.
	AccessibilityVisibleChildren() foundation.INSArray
	// Returns the visible columns for the table or outline.
	AccessibilityVisibleColumns() foundation.INSArray
	// Returns the visible rows for the table or outline.
	AccessibilityVisibleRows() foundation.INSArray
	// Returns the warning value for the level indicator.
	AccessibilityWarningValue() objectivec.IObject
	// Returns the window that contains the accessibility element.
	AccessibilityWindow() objectivec.IObject
	// Returns an array that contains all the app’s windows.
	AccessibilityWindows() foundation.INSArray
	// Returns the child accessibility element that represents the window’s zoom button.
	AccessibilityZoomButton() objectivec.IObject
	// A string that identifies the user interface item.
	Identifier() NSUserInterfaceItemIdentifier
	// Returns the Boolean value that determines whether the accessibility element’s alternative UI is currently visible.
	IsAccessibilityAlternateUIVisible() bool
	// Returns a Boolean value that determines whether the row is disclosing other rows.
	IsAccessibilityDisclosed() bool
	// Returns a Boolean value that indicates whether the accessibility element is in an edited state.
	IsAccessibilityEdited() bool
	// Returns a Boolean value that determines whether the accessibility element participates in the accessibility hierarchy.
	IsAccessibilityElement() bool
	// Returns a Boolean value that determines whether the accessibility element responds to user events.
	IsAccessibilityEnabled() bool
	// Returns a Boolean value that determines whether the accessibility element is in an expanded state.
	IsAccessibilityExpanded() bool
	// Returns a Boolean value that indicates whether the accessibility element has the keyboard focus.
	IsAccessibilityFocused() bool
	// Returns a Boolean value that determines whether the app is the frontmost app.
	IsAccessibilityFrontmost() bool
	// Returns a Boolean value that determines whether the app is in a hidden state.
	IsAccessibilityHidden() bool
	// Returns a Boolean value that determines whether the window is the app’s main window.
	IsAccessibilityMain() bool
	// Returns the Boolean value that determines whether the window is in a minimized state.
	IsAccessibilityMinimized() bool
	// Returns a Boolean value that determines whether the window is modal.
	IsAccessibilityModal() bool
	// Returns a Boolean value that determines whether the accessibility element’s grid is in row major order or in column major order.
	IsAccessibilityOrderedByRow() bool
	// Returns a Boolean value that determines whether the accessibility element contains protected content.
	IsAccessibilityProtectedContent() bool
	// Returns a Boolean value that determines whether the accessibility element must have content for successful submission of a form.
	IsAccessibilityRequired() bool
	// Returns a Boolean value that determines whether the accessibility element is currently in a selected state.
	IsAccessibilitySelected() bool
	// Returns a Boolean value that indicates whether assistive apps can invoke the specified selector on the accessibility element.
	IsAccessibilitySelectorAllowed(selector objc.SEL) bool
	// Sets the activation point for the user interface element.
	SetAccessibilityActivationPoint(accessibilityActivationPoint corefoundation.CGPoint)
	// Sets the allowed values for the slider accessibility element.
	SetAccessibilityAllowedValues(accessibilityAllowedValues []foundation.NSNumber)
	// Sets the Boolean value that determines whether the accessibility element’s alternative UI is currently visible.
	SetAccessibilityAlternateUIVisible(accessibilityAlternateUIVisible bool)
	// Sets the child accessibility element with the current focus.
	SetAccessibilityApplicationFocusedUIElement(accessibilityApplicationFocusedUIElement objectivec.IObject)
	SetAccessibilityAttributedUserInputLabels(accessibilityAttributedUserInputLabels []foundation.NSAttributedString)
	// Sets the child accessibility element that represents the window’s cancel button.
	SetAccessibilityCancelButton(accessibilityCancelButton objectivec.IObject)
	// Sets the child accessibility elements in the accessibility hierarchy.
	SetAccessibilityChildren(accessibilityChildren foundation.INSArray)
	// Sets the array of child accessibility elements in order for linear navigation.
	SetAccessibilityChildrenInNavigationOrder(accessibilityChildrenInNavigationOrder []objectivec.IObject)
	// Sets the clear button for the search field.
	SetAccessibilityClearButton(accessibilityClearButton objectivec.IObject)
	// Sets the child accessibility element that represents the window’s close button.
	SetAccessibilityCloseButton(accessibilityCloseButton objectivec.IObject)
	// Sets the number of columns in the accessibility element’s grid.
	SetAccessibilityColumnCount(accessibilityColumnCount int)
	// Sets the column header accessibility elements for the table or outline.
	SetAccessibilityColumnHeaderUIElements(accessibilityColumnHeaderUIElements foundation.INSArray)
	// Sets the column index range of the cell.
	SetAccessibilityColumnIndexRange(accessibilityColumnIndexRange foundation.NSRange)
	// Sets the column titles for the accessibility element.
	SetAccessibilityColumnTitles(accessibilityColumnTitles foundation.INSArray)
	// Sets the column accessibility elements for the table or outline.
	SetAccessibilityColumns(accessibilityColumns foundation.INSArray)
	// Sets the contents of the current accessibility element.
	SetAccessibilityContents(accessibilityContents foundation.INSArray)
	// Sets the critical value for the level indicator.
	SetAccessibilityCriticalValue(accessibilityCriticalValue objectivec.IObject)
	// Sets the custom actions of the current accessibility element.
	SetAccessibilityCustomActions(accessibilityCustomActions []NSAccessibilityCustomAction)
	// Sets the custom rotors of the current accessibility element.
	SetAccessibilityCustomRotors(accessibilityCustomRotors []NSAccessibilityCustomRotor)
	// Sets the decrement button for the stepper accessibility element.
	SetAccessibilityDecrementButton(accessibilityDecrementButton objectivec.IObject)
	// Sets the child accessibility element that represents the window’s default button.
	SetAccessibilityDefaultButton(accessibilityDefaultButton objectivec.IObject)
	// Sets a Boolean value that determines whether the row is disclosing other rows.
	SetAccessibilityDisclosed(accessibilityDisclosed bool)
	// Sets the row disclosing the current row.
	SetAccessibilityDisclosedByRow(accessibilityDisclosedByRow objectivec.IObject)
	// Sets the rows that the current row discloses.
	SetAccessibilityDisclosedRows(accessibilityDisclosedRows objectivec.IObject)
	// Sets the indention level for the row.
	SetAccessibilityDisclosureLevel(accessibilityDisclosureLevel int)
	// Sets the URL for the file that the accessibility element represents.
	SetAccessibilityDocument(accessibilityDocument string)
	// Sets a Boolean value that indicates whether the accessibility element is in an edited state.
	SetAccessibilityEdited(accessibilityEdited bool)
	// Sets a Boolean value that determines whether the accessibility element participates in the accessibility hierarchy.
	SetAccessibilityElement(accessibilityElement bool)
	// Sets a Boolean value that determines whether the accessibility element responds to user events.
	SetAccessibilityEnabled(accessibilityEnabled bool)
	// Sets a Boolean value that determines whether accessibility element is in an expanded state.
	SetAccessibilityExpanded(accessibilityExpanded bool)
	// Sets the icon for the app’s menu bar extra.
	SetAccessibilityExtrasMenuBar(accessibilityExtrasMenuBar objectivec.IObject)
	// Sets the filename for the file that the accessibility element represents.
	SetAccessibilityFilename(accessibilityFilename string)
	// Sets a Boolean value that determines whether the accessibility element has the keyboard focus.
	SetAccessibilityFocused(accessibilityFocused bool)
	// Sets the child window with the current focus.
	SetAccessibilityFocusedWindow(accessibilityFocusedWindow objectivec.IObject)
	// Sets the accessibility element’s frame in screen coordinates.
	SetAccessibilityFrame(accessibilityFrame corefoundation.CGRect)
	// Sets a Boolean value that determines whether the app is the frontmost app.
	SetAccessibilityFrontmost(accessibilityFrontmost bool)
	// Sets the child accessibility element that represents the window’s full-screen button.
	SetAccessibilityFullScreenButton(accessibilityFullScreenButton objectivec.IObject)
	// Sets the child accessibility element that represents the window’s grow area.
	SetAccessibilityGrowArea(accessibilityGrowArea objectivec.IObject)
	// Sets the drag handle accessibility elements for the layout item element.
	SetAccessibilityHandles(accessibilityHandles foundation.INSArray)
	// Sets the header for the table view.
	SetAccessibilityHeader(accessibilityHeader objectivec.IObject)
	// Sets the help text for the accessibility element.
	SetAccessibilityHelp(accessibilityHelp string)
	// Sets a Boolean value that determines whether the app is in a hidden state.
	SetAccessibilityHidden(accessibilityHidden bool)
	// Sets the horizontal scroll bar for the scroll view.
	SetAccessibilityHorizontalScrollBar(accessibilityHorizontalScrollBar objectivec.IObject)
	// Sets the description of the layout area’s horizontal units.
	SetAccessibilityHorizontalUnitDescription(accessibilityHorizontalUnitDescription string)
	// Sets the units that the layout area uses for horizontal values.
	SetAccessibilityHorizontalUnits(accessibilityHorizontalUnits NSAccessibilityUnits)
	// Sets the accessibility element’s identity.
	SetAccessibilityIdentifier(accessibilityIdentifier string)
	// Sets the increment button for the stepper accessibility element.
	SetAccessibilityIncrementButton(accessibilityIncrementButton objectivec.IObject)
	// Sets the index of the row or column that the accessibility element represents.
	SetAccessibilityIndex(accessibilityIndex int)
	// Sets the line number that contains the insertion point.
	SetAccessibilityInsertionPointLineNumber(accessibilityInsertionPointLineNumber int)
	// Sets a short description of the accessibility element.
	SetAccessibilityLabel(accessibilityLabel string)
	// Sets the child label elements for the slider accessibility element.
	SetAccessibilityLabelUIElements(accessibilityLabelUIElements foundation.INSArray)
	// Sets the value of the label accessibility element.
	SetAccessibilityLabelValue(accessibilityLabelValue float32)
	// Sets the elements that have links with the accessibility element.
	SetAccessibilityLinkedUIElements(accessibilityLinkedUIElements foundation.INSArray)
	// Sets a Boolean value that determines whether the window is the app’s main window.
	SetAccessibilityMain(accessibilityMain bool)
	// Sets the app’s main window.
	SetAccessibilityMainWindow(accessibilityMainWindow objectivec.IObject)
	// Sets the user interface element that functions as a marker group for the ruler accessibility element.
	SetAccessibilityMarkerGroupUIElement(accessibilityMarkerGroupUIElement objectivec.IObject)
	// Sets the human-readable description of the marker type.
	SetAccessibilityMarkerTypeDescription(accessibilityMarkerTypeDescription string)
	// Sets the array of marker accessibility elements for the ruler.
	SetAccessibilityMarkerUIElements(accessibilityMarkerUIElements foundation.INSArray)
	// Sets the marker values for the ruler.
	SetAccessibilityMarkerValues(accessibilityMarkerValues objectivec.IObject)
	// Sets the maximum value for the accessibility element.
	SetAccessibilityMaxValue(accessibilityMaxValue objectivec.IObject)
	// Sets the app’s menu bar.
	SetAccessibilityMenuBar(accessibilityMenuBar objectivec.IObject)
	// Sets the minimum value for the accessibility element.
	SetAccessibilityMinValue(accessibilityMinValue objectivec.IObject)
	// Sets the child accessibility element that represents the window’s minimize button.
	SetAccessibilityMinimizeButton(accessibilityMinimizeButton objectivec.IObject)
	// Sets the Boolean value that determines whether the window is in a minimized state.
	SetAccessibilityMinimized(accessibilityMinimized bool)
	// Sets a Boolean value that determines whether the window is modal.
	SetAccessibilityModal(accessibilityModal bool)
	// Sets the contents that follow the divider accessibility element.
	SetAccessibilityNextContents(accessibilityNextContents foundation.INSArray)
	// Sets the number of characters in the text.
	SetAccessibilityNumberOfCharacters(accessibilityNumberOfCharacters int)
	// Sets a Boolean value that determines whether the element’s grid is in row major order or in column major order.
	SetAccessibilityOrderedByRow(accessibilityOrderedByRow bool)
	// Sets the orientation of the accessibility element.
	SetAccessibilityOrientation(accessibilityOrientation NSAccessibilityOrientation)
	// Sets the overflow button for the toolbar.
	SetAccessibilityOverflowButton(accessibilityOverflowButton objectivec.IObject)
	// Sets the accessibility element’s parent in the accessibility hierarchy.
	SetAccessibilityParent(accessibilityParent objectivec.IObject)
	// Sets the placeholder value for the accessibility element.
	SetAccessibilityPlaceholderValue(accessibilityPlaceholderValue string)
	// Sets the contents that precede the divider accessibility element.
	SetAccessibilityPreviousContents(accessibilityPreviousContents foundation.INSArray)
	// Sets a Boolean value that determines whether the accessibility element contains protected content.
	SetAccessibilityProtectedContent(accessibilityProtectedContent bool)
	// Sets the child accessibility element that represents the window’s proxy icon.
	SetAccessibilityProxy(accessibilityProxy objectivec.IObject)
	// Sets a Boolean value that determines whether the accessibility element must have content for successful submission of a form.
	SetAccessibilityRequired(accessibilityRequired bool)
	// Sets the type of interface element that the accessibility element represents.
	SetAccessibilityRole(accessibilityRole NSAccessibilityRole)
	// Sets the localized, human-intelligible description of the accessibility element’s role, such as radio button.
	SetAccessibilityRoleDescription(accessibilityRoleDescription string)
	// Sets the number of rows in the accessibility element’s grid.
	SetAccessibilityRowCount(accessibilityRowCount int)
	// Sets the row header accessibility elements for the table or outline.
	SetAccessibilityRowHeaderUIElements(accessibilityRowHeaderUIElements foundation.INSArray)
	// Sets the row index range of the cell.
	SetAccessibilityRowIndexRange(accessibilityRowIndexRange foundation.NSRange)
	// Sets the row accessibility elements for the table or outline.
	SetAccessibilityRows(accessibilityRows foundation.INSArray)
	// Sets the type of markers for the ruler.
	SetAccessibilityRulerMarkerType(accessibilityRulerMarkerType NSAccessibilityRulerMarkerType)
	// Sets the search button for the search field.
	SetAccessibilitySearchButton(accessibilitySearchButton objectivec.IObject)
	// Sets the search menu for the search field.
	SetAccessibilitySearchMenu(accessibilitySearchMenu objectivec.IObject)
	// Sets a Boolean value that determines whether the accessibility element is currently in a selected state.
	SetAccessibilitySelected(accessibilitySelected bool)
	// Sets the currently selected cells for the table.
	SetAccessibilitySelectedCells(accessibilitySelectedCells foundation.INSArray)
	// Sets the accessibility element’s currently selected children.
	SetAccessibilitySelectedChildren(accessibilitySelectedChildren foundation.INSArray)
	// Sets the currently selected columns for the table or outline.
	SetAccessibilitySelectedColumns(accessibilitySelectedColumns foundation.INSArray)
	// Sets the currently selected rows for the table or outline.
	SetAccessibilitySelectedRows(accessibilitySelectedRows foundation.INSArray)
	// Sets the currently selected text.
	SetAccessibilitySelectedText(accessibilitySelectedText string)
	// Sets the range of the currently selected text.
	SetAccessibilitySelectedTextRange(accessibilitySelectedTextRange foundation.NSRange)
	// Sets an array of ranges for the currently selected text.
	SetAccessibilitySelectedTextRanges(accessibilitySelectedTextRanges []foundation.NSValue)
	// Sets the list of elements that the accessibility element is a title for.
	SetAccessibilityServesAsTitleForUIElements(accessibilityServesAsTitleForUIElements foundation.INSArray)
	// Sets the range of characters that the accessibility element displays.
	SetAccessibilitySharedCharacterRange(accessibilitySharedCharacterRange foundation.NSRange)
	// Sets the array of elements that shares the keyboard focus with the accessibility element.
	SetAccessibilitySharedFocusElements(accessibilitySharedFocusElements foundation.INSArray)
	// Sets the other elements that share text with the accessibility element.
	SetAccessibilitySharedTextUIElements(accessibilitySharedTextUIElements foundation.INSArray)
	// Sets the menu currently displaying for the accessibility element.
	SetAccessibilityShownMenu(accessibilityShownMenu objectivec.IObject)
	// Sets the accessibility element’s sort direction.
	SetAccessibilitySortDirection(accessibilitySortDirection NSAccessibilitySortDirection)
	// Sets the array that contains the views and splitter bar from the split view.
	SetAccessibilitySplitters(accessibilitySplitters foundation.INSArray)
	// Sets the specialized interface element type that the accessibility element represents.
	SetAccessibilitySubrole(accessibilitySubrole NSAccessibilitySubrole)
	// Sets the tab accessibility elements for the tab view.
	SetAccessibilityTabs(accessibilityTabs foundation.INSArray)
	// Sets the title of the accessibility element.
	SetAccessibilityTitle(accessibilityTitle string)
	// Sets the static text element that represents the accessibility element’s title.
	SetAccessibilityTitleUIElement(accessibilityTitleUIElement objectivec.IObject)
	// Sets the child accessibility element that represents the window’s toolbar button.
	SetAccessibilityToolbarButton(accessibilityToolbarButton objectivec.IObject)
	// Sets the top-level element that contains the accessibility element.
	SetAccessibilityTopLevelUIElement(accessibilityTopLevelUIElement objectivec.IObject)
	// Sets the URL for the accessibility element.
	SetAccessibilityURL(accessibilityURL foundation.NSURL)
	// Sets the human-readable description of the ruler’s units.
	SetAccessibilityUnitDescription(accessibilityUnitDescription string)
	// Sets the units used for the ruler.
	SetAccessibilityUnits(accessibilityUnits NSAccessibilityUnits)
	SetAccessibilityUserInputLabels(accessibilityUserInputLabels []string)
	// Sets the accessibility element’s value.
	SetAccessibilityValue(accessibilityValue objectivec.IObject)
	// Sets the human-readable description of the accessibility element’s value.
	SetAccessibilityValueDescription(accessibilityValueDescription string)
	// Sets the vertical scroll bar for the scroll view.
	SetAccessibilityVerticalScrollBar(accessibilityVerticalScrollBar objectivec.IObject)
	// Sets the description of the layout area’s vertical units.
	SetAccessibilityVerticalUnitDescription(accessibilityVerticalUnitDescription string)
	// Sets the units that the layout area uses for vertical values.
	SetAccessibilityVerticalUnits(accessibilityVerticalUnits NSAccessibilityUnits)
	// Sets the visible cells for the table.
	SetAccessibilityVisibleCells(accessibilityVisibleCells foundation.INSArray)
	// Sets the range of visible characters in the document.
	SetAccessibilityVisibleCharacterRange(accessibilityVisibleCharacterRange foundation.NSRange)
	// Sets the accessibility element’s visible child accessibility elements.
	SetAccessibilityVisibleChildren(accessibilityVisibleChildren foundation.INSArray)
	// Sets the visible columns for the table or outline.
	SetAccessibilityVisibleColumns(accessibilityVisibleColumns foundation.INSArray)
	// Sets the visible rows for the table or outline.
	SetAccessibilityVisibleRows(accessibilityVisibleRows foundation.INSArray)
	// Sets the warning value for the level indicator.
	SetAccessibilityWarningValue(accessibilityWarningValue objectivec.IObject)
	// Sets the window that contains the accessibility element.
	SetAccessibilityWindow(accessibilityWindow objectivec.IObject)
	// Sets the array that contains all the app’s windows.
	SetAccessibilityWindows(accessibilityWindows foundation.INSArray)
	// Sets the child accessibility element that represents the window’s zoom button.
	SetAccessibilityZoomButton(accessibilityZoomButton objectivec.IObject)
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

// Returns the activation point for the user interface element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityActivationPoint()
func (m NSMenuItem) AccessibilityActivationPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](m.ID, objc.Sel("accessibilityActivationPoint"))
	return corefoundation.CGPoint(rv)
}

// Returns the allowed values for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityAllowedValues()
func (m NSMenuItem) AccessibilityAllowedValues() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("accessibilityAllowedValues"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// Returns the child accessibility element with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityApplicationFocusedUIElement()
func (m NSMenuItem) AccessibilityApplicationFocusedUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityApplicationFocusedUIElement"))
	return objectivec.Object{ID: rv}
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
func (m NSMenuItem) AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityAttributedStringForRange:"), range_)
	return foundation.NSAttributedStringFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityAttributedUserInputLabels()
func (m NSMenuItem) AccessibilityAttributedUserInputLabels() []foundation.NSAttributedString {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("accessibilityAttributedUserInputLabels"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSAttributedString {
		return foundation.NSAttributedStringFromID(id)
	})
}

// Returns the child accessibility element that represents the window’s
// cancel button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCancelButton()
func (m NSMenuItem) AccessibilityCancelButton() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityCancelButton"))
	return objectivec.Object{ID: rv}
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
func (m NSMenuItem) AccessibilityCellForColumnRow(column int, row int) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityCellForColumn:row:"), column, row)
	return objectivec.Object{ID: rv}
}

// Returns the child accessibility elements in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityChildren()
func (m NSMenuItem) AccessibilityChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityChildren"))
	return foundation.NSArrayFromID(rv)
}

// Returns the array of child accessibility elements in order for linear
// navigation.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityChildrenInNavigationOrder()
func (m NSMenuItem) AccessibilityChildrenInNavigationOrder() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("accessibilityChildrenInNavigationOrder"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns the clear button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityClearButton()
func (m NSMenuItem) AccessibilityClearButton() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityClearButton"))
	return objectivec.Object{ID: rv}
}

// Returns the child accessibility element that represents the window’s
// close button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCloseButton()
func (m NSMenuItem) AccessibilityCloseButton() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityCloseButton"))
	return objectivec.Object{ID: rv}
}

// Returns the number of columns in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnCount()
func (m NSMenuItem) AccessibilityColumnCount() int {
	rv := objc.Send[int](m.ID, objc.Sel("accessibilityColumnCount"))
	return rv
}

// Returns the column header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnHeaderUIElements()
func (m NSMenuItem) AccessibilityColumnHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityColumnHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the column index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnIndexRange()
func (m NSMenuItem) AccessibilityColumnIndexRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](m.ID, objc.Sel("accessibilityColumnIndexRange"))
	return foundation.NSRange(rv)
}

// Returns the column titles for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnTitles()
func (m NSMenuItem) AccessibilityColumnTitles() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityColumnTitles"))
	return foundation.NSArrayFromID(rv)
}

// Returns the column accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumns()
func (m NSMenuItem) AccessibilityColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityColumns"))
	return foundation.NSArrayFromID(rv)
}

// Returns the contents of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityContents()
func (m NSMenuItem) AccessibilityContents() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityContents"))
	return foundation.NSArrayFromID(rv)
}

// Returns the critical value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCriticalValue()
func (m NSMenuItem) AccessibilityCriticalValue() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityCriticalValue"))
	return objectivec.Object{ID: rv}
}

// Returns the custom actions of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCustomActions()
func (m NSMenuItem) AccessibilityCustomActions() []NSAccessibilityCustomAction {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("accessibilityCustomActions"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSAccessibilityCustomAction {
		return NSAccessibilityCustomActionFromID(id)
	})
}

// Returns the custom rotors of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCustomRotors()
func (m NSMenuItem) AccessibilityCustomRotors() []NSAccessibilityCustomRotor {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("accessibilityCustomRotors"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSAccessibilityCustomRotor {
		return NSAccessibilityCustomRotorFromID(id)
	})
}

// Returns the decrement button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDecrementButton()
func (m NSMenuItem) AccessibilityDecrementButton() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityDecrementButton"))
	return objectivec.Object{ID: rv}
}

// Returns the child accessibility element that represents the window’s
// default button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDefaultButton()
func (m NSMenuItem) AccessibilityDefaultButton() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityDefaultButton"))
	return objectivec.Object{ID: rv}
}

// Returns the row disclosing the current row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDisclosedByRow()
func (m NSMenuItem) AccessibilityDisclosedByRow() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityDisclosedByRow"))
	return objectivec.Object{ID: rv}
}

// Returns the rows that the current row discloses.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDisclosedRows()
func (m NSMenuItem) AccessibilityDisclosedRows() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityDisclosedRows"))
	return objectivec.Object{ID: rv}
}

// Returns the indention level for the row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDisclosureLevel()
func (m NSMenuItem) AccessibilityDisclosureLevel() int {
	rv := objc.Send[int](m.ID, objc.Sel("accessibilityDisclosureLevel"))
	return rv
}

// Returns the URL for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDocument()
func (m NSMenuItem) AccessibilityDocument() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityDocument"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the icon for the app’s menu bar extra.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityExtrasMenuBar()
func (m NSMenuItem) AccessibilityExtrasMenuBar() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityExtrasMenuBar"))
	return objectivec.Object{ID: rv}
}

// Returns the filename for the file that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFilename()
func (m NSMenuItem) AccessibilityFilename() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityFilename"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the child window with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFocusedWindow()
func (m NSMenuItem) AccessibilityFocusedWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityFocusedWindow"))
	return objectivec.Object{ID: rv}
}

// Returns the accessibility element’s frame in screen coordinates.
//
// # Return Value
//
// The element’s frame in screen coordinates.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilityFrame] property. This method is called whenever
// accessibility clients request the [size] or [position] attributes.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityFrame()
//
// [position]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
// [size]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
func (m NSMenuItem) AccessibilityFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m.ID, objc.Sel("accessibilityFrame"))
	return corefoundation.CGRect(rv)
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
func (m NSMenuItem) AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](m.ID, objc.Sel("accessibilityFrameForRange:"), range_)
	return corefoundation.CGRect(rv)
}

// Returns the child accessibility element that represents the window’s
// full-screen button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFullScreenButton()
func (m NSMenuItem) AccessibilityFullScreenButton() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityFullScreenButton"))
	return objectivec.Object{ID: rv}
}

// Returns the child accessibility element that represents the window’s grow
// area.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityGrowArea()
func (m NSMenuItem) AccessibilityGrowArea() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityGrowArea"))
	return objectivec.Object{ID: rv}
}

// Returns the drag handle elements for the layout item element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHandles()
func (m NSMenuItem) AccessibilityHandles() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityHandles"))
	return foundation.NSArrayFromID(rv)
}

// Returns the header for the table view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHeader()
func (m NSMenuItem) AccessibilityHeader() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityHeader"))
	return objectivec.Object{ID: rv}
}

// Returns the help text for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHelp()
func (m NSMenuItem) AccessibilityHelp() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityHelp"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the horizontal scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHorizontalScrollBar()
func (m NSMenuItem) AccessibilityHorizontalScrollBar() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityHorizontalScrollBar"))
	return objectivec.Object{ID: rv}
}

// Returns the description of the layout area’s horizontal units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHorizontalUnitDescription()
func (m NSMenuItem) AccessibilityHorizontalUnitDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityHorizontalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the units that the layout area uses for horizontal values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHorizontalUnits()
func (m NSMenuItem) AccessibilityHorizontalUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](m.ID, objc.Sel("accessibilityHorizontalUnits"))
	return NSAccessibilityUnits(rv)
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
// [NSWindow.AccessibilityIdentifier] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityIdentifier()
func (m NSMenuItem) AccessibilityIdentifier() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the increment button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityIncrementButton()
func (m NSMenuItem) AccessibilityIncrementButton() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityIncrementButton"))
	return objectivec.Object{ID: rv}
}

// Returns the index of the row or column that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityIndex()
func (m NSMenuItem) AccessibilityIndex() int {
	rv := objc.Send[int](m.ID, objc.Sel("accessibilityIndex"))
	return rv
}

// Returns the line number that contains the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityInsertionPointLineNumber()
func (m NSMenuItem) AccessibilityInsertionPointLineNumber() int {
	rv := objc.Send[int](m.ID, objc.Sel("accessibilityInsertionPointLineNumber"))
	return rv
}

// Returns a short description of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLabel()
func (m NSMenuItem) AccessibilityLabel() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the child label elements for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLabelUIElements()
func (m NSMenuItem) AccessibilityLabelUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityLabelUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the value of the label accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLabelValue()
func (m NSMenuItem) AccessibilityLabelValue() float32 {
	rv := objc.Send[float32](m.ID, objc.Sel("accessibilityLabelValue"))
	return rv
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
func (m NSMenuItem) AccessibilityLayoutPointForScreenPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](m.ID, objc.Sel("accessibilityLayoutPointForScreenPoint:"), point)
	return corefoundation.CGPoint(rv)
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
func (m NSMenuItem) AccessibilityLayoutSizeForScreenSize(size corefoundation.CGSize) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m.ID, objc.Sel("accessibilityLayoutSizeForScreenSize:"), size)
	return corefoundation.CGSize(rv)
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
func (m NSMenuItem) AccessibilityLineForIndex(index int) int {
	rv := objc.Send[int](m.ID, objc.Sel("accessibilityLineForIndex:"), index)
	return rv
}

// Returns the elements that have links with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLinkedUIElements()
func (m NSMenuItem) AccessibilityLinkedUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityLinkedUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the app’s main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMainWindow()
func (m NSMenuItem) AccessibilityMainWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityMainWindow"))
	return objectivec.Object{ID: rv}
}

// Returns the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerGroupUIElement()
func (m NSMenuItem) AccessibilityMarkerGroupUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityMarkerGroupUIElement"))
	return objectivec.Object{ID: rv}
}

// Returns the human-readable description of the marker type.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerTypeDescription()
func (m NSMenuItem) AccessibilityMarkerTypeDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityMarkerTypeDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the array of marker accessibility elements for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerUIElements()
func (m NSMenuItem) AccessibilityMarkerUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityMarkerUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the marker values for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerValues()
func (m NSMenuItem) AccessibilityMarkerValues() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityMarkerValues"))
	return objectivec.Object{ID: rv}
}

// Returns the maximum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMaxValue()
func (m NSMenuItem) AccessibilityMaxValue() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityMaxValue"))
	return objectivec.Object{ID: rv}
}

// Returns the app’s menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMenuBar()
func (m NSMenuItem) AccessibilityMenuBar() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityMenuBar"))
	return objectivec.Object{ID: rv}
}

// Returns the minimum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMinValue()
func (m NSMenuItem) AccessibilityMinValue() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityMinValue"))
	return objectivec.Object{ID: rv}
}

// Returns the child accessibility element that represents the window’s
// minimize button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMinimizeButton()
func (m NSMenuItem) AccessibilityMinimizeButton() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityMinimizeButton"))
	return objectivec.Object{ID: rv}
}

// Returns the contents that follow the divider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityNextContents()
func (m NSMenuItem) AccessibilityNextContents() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityNextContents"))
	return foundation.NSArrayFromID(rv)
}

// Returns the number of characters in the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityNumberOfCharacters()
func (m NSMenuItem) AccessibilityNumberOfCharacters() int {
	rv := objc.Send[int](m.ID, objc.Sel("accessibilityNumberOfCharacters"))
	return rv
}

// Returns the orientation of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityOrientation()
func (m NSMenuItem) AccessibilityOrientation() NSAccessibilityOrientation {
	rv := objc.Send[NSAccessibilityOrientation](m.ID, objc.Sel("accessibilityOrientation"))
	return NSAccessibilityOrientation(rv)
}

// Returns the overflow button for the toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityOverflowButton()
func (m NSMenuItem) AccessibilityOverflowButton() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityOverflowButton"))
	return objectivec.Object{ID: rv}
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
// [NSWindow.AccessibilityParent] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityParent()
func (m NSMenuItem) AccessibilityParent() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityParent"))
	return objectivec.Object{ID: rv}
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
func (m NSMenuItem) AccessibilityPerformCancel() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("accessibilityPerformCancel"))
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
func (m NSMenuItem) AccessibilityPerformConfirm() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("accessibilityPerformConfirm"))
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
// Use this method on elements that have an adjustable
// [NSWindow.AccessibilityValue] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDecrement()
func (m NSMenuItem) AccessibilityPerformDecrement() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("accessibilityPerformDecrement"))
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
func (m NSMenuItem) AccessibilityPerformDelete() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("accessibilityPerformDelete"))
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
// Use this method on elements that have an adjustable
// [NSWindow.AccessibilityValue] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformIncrement()
func (m NSMenuItem) AccessibilityPerformIncrement() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("accessibilityPerformIncrement"))
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
func (m NSMenuItem) AccessibilityPerformPick() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("accessibilityPerformPick"))
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
func (m NSMenuItem) AccessibilityPerformPress() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("accessibilityPerformPress"))
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
func (m NSMenuItem) AccessibilityPerformRaise() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("accessibilityPerformRaise"))
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
func (m NSMenuItem) AccessibilityPerformShowAlternateUI() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("accessibilityPerformShowAlternateUI"))
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
func (m NSMenuItem) AccessibilityPerformShowDefaultUI() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("accessibilityPerformShowDefaultUI"))
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
func (m NSMenuItem) AccessibilityPerformShowMenu() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("accessibilityPerformShowMenu"))
	return rv
}

// Returns the placeholder value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPlaceholderValue()
func (m NSMenuItem) AccessibilityPlaceholderValue() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityPlaceholderValue"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the contents that precede the divider accessibility element.
//
// # Return Value
//
// Sets the contents preceding this divider element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPreviousContents()
func (m NSMenuItem) AccessibilityPreviousContents() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityPreviousContents"))
	return foundation.NSArrayFromID(rv)
}

// Returns the child accessibility element that represents the window’s
// proxy icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityProxy()
func (m NSMenuItem) AccessibilityProxy() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityProxy"))
	return objectivec.Object{ID: rv}
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
func (m NSMenuItem) AccessibilityRTFForRange(range_ foundation.NSRange) foundation.NSData {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityRTFForRange:"), range_)
	return foundation.NSDataFromID(rv)
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
func (m NSMenuItem) AccessibilityRangeForIndex(index int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](m.ID, objc.Sel("accessibilityRangeForIndex:"), index)
	return foundation.NSRange(rv)
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
func (m NSMenuItem) AccessibilityRangeForLine(line int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](m.ID, objc.Sel("accessibilityRangeForLine:"), line)
	return foundation.NSRange(rv)
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
func (m NSMenuItem) AccessibilityRangeForPosition(point corefoundation.CGPoint) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](m.ID, objc.Sel("accessibilityRangeForPosition:"), point)
	return foundation.NSRange(rv)
}

// Returns the type of interface element that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRole()
func (m NSMenuItem) AccessibilityRole() NSAccessibilityRole {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityRole"))
	return NSAccessibilityRole(foundation.NSStringFromID(rv).String())
}

// Returns a localized, human-intelligible description of the accessibility
// element’s role, such as radio button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRoleDescription()
func (m NSMenuItem) AccessibilityRoleDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityRoleDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the number of rows in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRowCount()
func (m NSMenuItem) AccessibilityRowCount() int {
	rv := objc.Send[int](m.ID, objc.Sel("accessibilityRowCount"))
	return rv
}

// Returns the row header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRowHeaderUIElements()
func (m NSMenuItem) AccessibilityRowHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityRowHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the row index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRowIndexRange()
func (m NSMenuItem) AccessibilityRowIndexRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](m.ID, objc.Sel("accessibilityRowIndexRange"))
	return foundation.NSRange(rv)
}

// Returns the row accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRows()
func (m NSMenuItem) AccessibilityRows() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityRows"))
	return foundation.NSArrayFromID(rv)
}

// Returns the type of markers for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRulerMarkerType()
func (m NSMenuItem) AccessibilityRulerMarkerType() NSAccessibilityRulerMarkerType {
	rv := objc.Send[NSAccessibilityRulerMarkerType](m.ID, objc.Sel("accessibilityRulerMarkerType"))
	return NSAccessibilityRulerMarkerType(rv)
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
func (m NSMenuItem) AccessibilityScreenPointForLayoutPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](m.ID, objc.Sel("accessibilityScreenPointForLayoutPoint:"), point)
	return corefoundation.CGPoint(rv)
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
func (m NSMenuItem) AccessibilityScreenSizeForLayoutSize(size corefoundation.CGSize) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](m.ID, objc.Sel("accessibilityScreenSizeForLayoutSize:"), size)
	return corefoundation.CGSize(rv)
}

// Returns the search button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySearchButton()
func (m NSMenuItem) AccessibilitySearchButton() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilitySearchButton"))
	return objectivec.Object{ID: rv}
}

// Returns the search menu for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySearchMenu()
func (m NSMenuItem) AccessibilitySearchMenu() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilitySearchMenu"))
	return objectivec.Object{ID: rv}
}

// Returns the currently selected cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedCells()
func (m NSMenuItem) AccessibilitySelectedCells() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilitySelectedCells"))
	return foundation.NSArrayFromID(rv)
}

// Returns the accessibility element’s currently selected children.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedChildren()
func (m NSMenuItem) AccessibilitySelectedChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilitySelectedChildren"))
	return foundation.NSArrayFromID(rv)
}

// Returns the currently selected columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedColumns()
func (m NSMenuItem) AccessibilitySelectedColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilitySelectedColumns"))
	return foundation.NSArrayFromID(rv)
}

// Returns the currently selected rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedRows()
func (m NSMenuItem) AccessibilitySelectedRows() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilitySelectedRows"))
	return foundation.NSArrayFromID(rv)
}

// Returns the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedText()
func (m NSMenuItem) AccessibilitySelectedText() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilitySelectedText"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the range of the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedTextRange()
func (m NSMenuItem) AccessibilitySelectedTextRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](m.ID, objc.Sel("accessibilitySelectedTextRange"))
	return foundation.NSRange(rv)
}

// Returns an array of ranges for the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedTextRanges()
func (m NSMenuItem) AccessibilitySelectedTextRanges() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("accessibilitySelectedTextRanges"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}

// Returns the list of elements that the accessibility element is a title for.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityServesAsTitleForUIElements()
func (m NSMenuItem) AccessibilityServesAsTitleForUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityServesAsTitleForUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the range of characters that the accessibility element displays.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySharedCharacterRange()
func (m NSMenuItem) AccessibilitySharedCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](m.ID, objc.Sel("accessibilitySharedCharacterRange"))
	return foundation.NSRange(rv)
}

// Returns the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySharedFocusElements()
func (m NSMenuItem) AccessibilitySharedFocusElements() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilitySharedFocusElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the other elements that share text with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySharedTextUIElements()
func (m NSMenuItem) AccessibilitySharedTextUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilitySharedTextUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the menu currently displaying for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityShownMenu()
func (m NSMenuItem) AccessibilityShownMenu() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityShownMenu"))
	return objectivec.Object{ID: rv}
}

// Returns the accessibility element’s sort direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySortDirection()
func (m NSMenuItem) AccessibilitySortDirection() NSAccessibilitySortDirection {
	rv := objc.Send[NSAccessibilitySortDirection](m.ID, objc.Sel("accessibilitySortDirection"))
	return NSAccessibilitySortDirection(rv)
}

// Returns an array that contains the views and splitter bar from the split
// view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySplitters()
func (m NSMenuItem) AccessibilitySplitters() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilitySplitters"))
	return foundation.NSArrayFromID(rv)
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
func (m NSMenuItem) AccessibilityStringForRange(range_ foundation.NSRange) string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityStringForRange:"), range_)
	return foundation.NSStringFromID(rv).String()
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
func (m NSMenuItem) AccessibilityStyleRangeForIndex(index int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](m.ID, objc.Sel("accessibilityStyleRangeForIndex:"), index)
	return foundation.NSRange(rv)
}

// Returns the specialized interface element type that the accessibility
// element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySubrole()
func (m NSMenuItem) AccessibilitySubrole() NSAccessibilitySubrole {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilitySubrole"))
	return NSAccessibilitySubrole(foundation.NSStringFromID(rv).String())
}

// Returns the tab accessibility elements for the tab view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTabs()
func (m NSMenuItem) AccessibilityTabs() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityTabs"))
	return foundation.NSArrayFromID(rv)
}

// Returns the title of the accessibility element—for example, a button’s
// visible text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTitle()
func (m NSMenuItem) AccessibilityTitle() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityTitle"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the static text element that represents the accessibility
// element’s title.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTitleUIElement()
func (m NSMenuItem) AccessibilityTitleUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityTitleUIElement"))
	return objectivec.Object{ID: rv}
}

// Returns the child accessibility element that represents the window’s
// toolbar button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityToolbarButton()
func (m NSMenuItem) AccessibilityToolbarButton() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityToolbarButton"))
	return objectivec.Object{ID: rv}
}

// Returns the top-level element that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTopLevelUIElement()
func (m NSMenuItem) AccessibilityTopLevelUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityTopLevelUIElement"))
	return objectivec.Object{ID: rv}
}

// Returns the URL for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityURL()
func (m NSMenuItem) AccessibilityURL() foundation.NSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityURL"))
	return foundation.NSURLFromID(rv)
}

// Returns the human-readable description of the ruler’s units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityUnitDescription()
func (m NSMenuItem) AccessibilityUnitDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the units for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityUnits()
func (m NSMenuItem) AccessibilityUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](m.ID, objc.Sel("accessibilityUnits"))
	return NSAccessibilityUnits(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityUserInputLabels()
func (m NSMenuItem) AccessibilityUserInputLabels() []string {
	rv := objc.Send[[]objc.ID](m.ID, objc.Sel("accessibilityUserInputLabels"))
	return objc.ConvertSliceToStrings(rv)
}

// Returns the accessibility element’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityValue()
func (m NSMenuItem) AccessibilityValue() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityValue"))
	return objectivec.Object{ID: rv}
}

// Returns the human-readable description of the accessibility element’s
// value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityValueDescription()
func (m NSMenuItem) AccessibilityValueDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityValueDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the vertical scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVerticalScrollBar()
func (m NSMenuItem) AccessibilityVerticalScrollBar() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityVerticalScrollBar"))
	return objectivec.Object{ID: rv}
}

// Returns the description of the layout area’s vertical units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVerticalUnitDescription()
func (m NSMenuItem) AccessibilityVerticalUnitDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityVerticalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the units that the layout area uses for vertical values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVerticalUnits()
func (m NSMenuItem) AccessibilityVerticalUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](m.ID, objc.Sel("accessibilityVerticalUnits"))
	return NSAccessibilityUnits(rv)
}

// Returns the visible cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleCells()
func (m NSMenuItem) AccessibilityVisibleCells() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityVisibleCells"))
	return foundation.NSArrayFromID(rv)
}

// Returns the range of visible characters in the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleCharacterRange()
func (m NSMenuItem) AccessibilityVisibleCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](m.ID, objc.Sel("accessibilityVisibleCharacterRange"))
	return foundation.NSRange(rv)
}

// Returns the accessibility element’s visible child accessibility elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleChildren()
func (m NSMenuItem) AccessibilityVisibleChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityVisibleChildren"))
	return foundation.NSArrayFromID(rv)
}

// Returns the visible columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleColumns()
func (m NSMenuItem) AccessibilityVisibleColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityVisibleColumns"))
	return foundation.NSArrayFromID(rv)
}

// Returns the visible rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleRows()
func (m NSMenuItem) AccessibilityVisibleRows() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityVisibleRows"))
	return foundation.NSArrayFromID(rv)
}

// Returns the warning value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityWarningValue()
func (m NSMenuItem) AccessibilityWarningValue() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityWarningValue"))
	return objectivec.Object{ID: rv}
}

// Returns the window that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityWindow()
func (m NSMenuItem) AccessibilityWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityWindow"))
	return objectivec.Object{ID: rv}
}

// Returns an array that contains all the app’s windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityWindows()
func (m NSMenuItem) AccessibilityWindows() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityWindows"))
	return foundation.NSArrayFromID(rv)
}

// Returns the child accessibility element that represents the window’s zoom
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityZoomButton()
func (m NSMenuItem) AccessibilityZoomButton() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("accessibilityZoomButton"))
	return objectivec.Object{ID: rv}
}

// A string that identifies the user interface item.
//
// See: https://developer.apple.com/documentation/AppKit/NSUserInterfaceItemIdentification/identifier
func (m NSMenuItem) Identifier() NSUserInterfaceItemIdentifier {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("identifier"))
	return NSUserInterfaceItemIdentifier(foundation.NSStringFromID(rv).String())
}

// Returns the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityAlternateUIVisible()
func (m NSMenuItem) IsAccessibilityAlternateUIVisible() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
	return rv
}

// Returns a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityDisclosed()
func (m NSMenuItem) IsAccessibilityDisclosed() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isAccessibilityDisclosed"))
	return rv
}

// Returns a Boolean value that indicates whether the accessibility element is
// in an edited state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEdited()
func (m NSMenuItem) IsAccessibilityEdited() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isAccessibilityEdited"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityElement()
func (m NSMenuItem) IsAccessibilityElement() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isAccessibilityElement"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEnabled()
func (m NSMenuItem) IsAccessibilityEnabled() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isAccessibilityEnabled"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// is in an expanded state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityExpanded()
func (m NSMenuItem) IsAccessibilityExpanded() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isAccessibilityExpanded"))
	return rv
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
// [NSWindow.AccessibilityFocused] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/isAccessibilityFocused()
func (m NSMenuItem) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isAccessibilityFocused"))
	return rv
}

// Returns a Boolean value that determines whether the app is the frontmost
// app.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityFrontmost()
func (m NSMenuItem) IsAccessibilityFrontmost() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isAccessibilityFrontmost"))
	return rv
}

// Returns a Boolean value that determines whether the app is in a hidden
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityHidden()
func (m NSMenuItem) IsAccessibilityHidden() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isAccessibilityHidden"))
	return rv
}

// Returns a Boolean value that determines whether the window is the app’s
// main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMain()
func (m NSMenuItem) IsAccessibilityMain() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isAccessibilityMain"))
	return rv
}

// Returns the Boolean value that determines whether the window is in a
// minimized state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMinimized()
func (m NSMenuItem) IsAccessibilityMinimized() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isAccessibilityMinimized"))
	return rv
}

// Returns a Boolean value that determines whether the window is modal.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityModal()
func (m NSMenuItem) IsAccessibilityModal() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isAccessibilityModal"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility
// element’s grid is in row major order or in column major order.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityOrderedByRow()
func (m NSMenuItem) IsAccessibilityOrderedByRow() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isAccessibilityOrderedByRow"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityProtectedContent()
func (m NSMenuItem) IsAccessibilityProtectedContent() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isAccessibilityProtectedContent"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// must have content for successful submission of a form.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityRequired()
func (m NSMenuItem) IsAccessibilityRequired() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isAccessibilityRequired"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// is currently in a selected state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelected()
func (m NSMenuItem) IsAccessibilitySelected() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isAccessibilitySelected"))
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
func (m NSMenuItem) IsAccessibilitySelectorAllowed(selector objc.SEL) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isAccessibilitySelectorAllowed:"), selector)
	return rv
}

// Sets the activation point for the user interface element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityActivationPoint(_:)
func (m NSMenuItem) SetAccessibilityActivationPoint(accessibilityActivationPoint corefoundation.CGPoint) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityActivationPoint:"), accessibilityActivationPoint)
}

// Sets the allowed values for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityAllowedValues(_:)
func (m NSMenuItem) SetAccessibilityAllowedValues(accessibilityAllowedValues []foundation.NSNumber) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityAllowedValues:"), objectivec.IObjectSliceToNSArray(accessibilityAllowedValues))
}

// Sets the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityAlternateUIVisible(_:)
func (m NSMenuItem) SetAccessibilityAlternateUIVisible(accessibilityAlternateUIVisible bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityAlternateUIVisible:"), accessibilityAlternateUIVisible)
}

// Sets the child accessibility element with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityApplicationFocusedUIElement(_:)
func (m NSMenuItem) SetAccessibilityApplicationFocusedUIElement(accessibilityApplicationFocusedUIElement objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityApplicationFocusedUIElement:"), accessibilityApplicationFocusedUIElement)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityAttributedUserInputLabels(_:)
func (m NSMenuItem) SetAccessibilityAttributedUserInputLabels(accessibilityAttributedUserInputLabels []foundation.NSAttributedString) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityAttributedUserInputLabels:"), objectivec.IObjectSliceToNSArray(accessibilityAttributedUserInputLabels))
}

// Sets the child accessibility element that represents the window’s cancel
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCancelButton(_:)
func (m NSMenuItem) SetAccessibilityCancelButton(accessibilityCancelButton objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityCancelButton:"), accessibilityCancelButton)
}

// Sets the child accessibility elements in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityChildren(_:)
func (m NSMenuItem) SetAccessibilityChildren(accessibilityChildren foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityChildren:"), accessibilityChildren)
}

// Sets the array of child accessibility elements in order for linear
// navigation.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityChildrenInNavigationOrder(_:)
func (m NSMenuItem) SetAccessibilityChildrenInNavigationOrder(accessibilityChildrenInNavigationOrder []objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityChildrenInNavigationOrder:"), objectivec.IObjectSliceToNSArray(accessibilityChildrenInNavigationOrder))
}

// Sets the clear button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityClearButton(_:)
func (m NSMenuItem) SetAccessibilityClearButton(accessibilityClearButton objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityClearButton:"), accessibilityClearButton)
}

// Sets the child accessibility element that represents the window’s close
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCloseButton(_:)
func (m NSMenuItem) SetAccessibilityCloseButton(accessibilityCloseButton objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityCloseButton:"), accessibilityCloseButton)
}

// Sets the number of columns in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnCount(_:)
func (m NSMenuItem) SetAccessibilityColumnCount(accessibilityColumnCount int) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityColumnCount:"), accessibilityColumnCount)
}

// Sets the column header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnHeaderUIElements(_:)
func (m NSMenuItem) SetAccessibilityColumnHeaderUIElements(accessibilityColumnHeaderUIElements foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityColumnHeaderUIElements:"), accessibilityColumnHeaderUIElements)
}

// Sets the column index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnIndexRange(_:)
func (m NSMenuItem) SetAccessibilityColumnIndexRange(accessibilityColumnIndexRange foundation.NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityColumnIndexRange:"), accessibilityColumnIndexRange)
}

// Sets the column titles for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnTitles(_:)
func (m NSMenuItem) SetAccessibilityColumnTitles(accessibilityColumnTitles foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityColumnTitles:"), accessibilityColumnTitles)
}

// Sets the column accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumns(_:)
func (m NSMenuItem) SetAccessibilityColumns(accessibilityColumns foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityColumns:"), accessibilityColumns)
}

// Sets the contents of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityContents(_:)
func (m NSMenuItem) SetAccessibilityContents(accessibilityContents foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityContents:"), accessibilityContents)
}

// Sets the critical value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCriticalValue(_:)
func (m NSMenuItem) SetAccessibilityCriticalValue(accessibilityCriticalValue objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityCriticalValue:"), accessibilityCriticalValue)
}

// Sets the custom actions of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCustomActions(_:)
func (m NSMenuItem) SetAccessibilityCustomActions(accessibilityCustomActions []NSAccessibilityCustomAction) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityCustomActions:"), objectivec.IObjectSliceToNSArray(accessibilityCustomActions))
}

// Sets the custom rotors of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCustomRotors(_:)
func (m NSMenuItem) SetAccessibilityCustomRotors(accessibilityCustomRotors []NSAccessibilityCustomRotor) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityCustomRotors:"), objectivec.IObjectSliceToNSArray(accessibilityCustomRotors))
}

// Sets the decrement button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDecrementButton(_:)
func (m NSMenuItem) SetAccessibilityDecrementButton(accessibilityDecrementButton objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityDecrementButton:"), accessibilityDecrementButton)
}

// Sets the child accessibility element that represents the window’s default
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDefaultButton(_:)
func (m NSMenuItem) SetAccessibilityDefaultButton(accessibilityDefaultButton objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityDefaultButton:"), accessibilityDefaultButton)
}

// Sets a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosed(_:)
func (m NSMenuItem) SetAccessibilityDisclosed(accessibilityDisclosed bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityDisclosed:"), accessibilityDisclosed)
}

// Sets the row disclosing the current row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosedByRow(_:)
func (m NSMenuItem) SetAccessibilityDisclosedByRow(accessibilityDisclosedByRow objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityDisclosedByRow:"), accessibilityDisclosedByRow)
}

// Sets the rows that the current row discloses.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosedRows(_:)
func (m NSMenuItem) SetAccessibilityDisclosedRows(accessibilityDisclosedRows objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityDisclosedRows:"), accessibilityDisclosedRows)
}

// Sets the indention level for the row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosureLevel(_:)
func (m NSMenuItem) SetAccessibilityDisclosureLevel(accessibilityDisclosureLevel int) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityDisclosureLevel:"), accessibilityDisclosureLevel)
}

// Sets the URL for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDocument(_:)
func (m NSMenuItem) SetAccessibilityDocument(accessibilityDocument string) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityDocument:"), objc.String(accessibilityDocument))
}

// Sets a Boolean value that indicates whether the accessibility element is in
// an edited state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityEdited(_:)
func (m NSMenuItem) SetAccessibilityEdited(accessibilityEdited bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityEdited:"), accessibilityEdited)
}

// Sets a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityElement(_:)
func (m NSMenuItem) SetAccessibilityElement(accessibilityElement bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityElement:"), accessibilityElement)
}

// Sets a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityEnabled(_:)
func (m NSMenuItem) SetAccessibilityEnabled(accessibilityEnabled bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityEnabled:"), accessibilityEnabled)
}

// Sets a Boolean value that determines whether accessibility element is in an
// expanded state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityExpanded(_:)
func (m NSMenuItem) SetAccessibilityExpanded(accessibilityExpanded bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityExpanded:"), accessibilityExpanded)
}

// Sets the icon for the app’s menu bar extra.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityExtrasMenuBar(_:)
func (m NSMenuItem) SetAccessibilityExtrasMenuBar(accessibilityExtrasMenuBar objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityExtrasMenuBar:"), accessibilityExtrasMenuBar)
}

// Sets the filename for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFilename(_:)
func (m NSMenuItem) SetAccessibilityFilename(accessibilityFilename string) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityFilename:"), objc.String(accessibilityFilename))
}

// Sets a Boolean value that determines whether the accessibility element has
// the keyboard focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFocused(_:)
func (m NSMenuItem) SetAccessibilityFocused(accessibilityFocused bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityFocused:"), accessibilityFocused)
}

// Sets the child window with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFocusedWindow(_:)
func (m NSMenuItem) SetAccessibilityFocusedWindow(accessibilityFocusedWindow objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityFocusedWindow:"), accessibilityFocusedWindow)
}

// Sets the accessibility element’s frame in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFrame(_:)
func (m NSMenuItem) SetAccessibilityFrame(accessibilityFrame corefoundation.CGRect) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityFrame:"), accessibilityFrame)
}

// Sets a Boolean value that determines whether the app is the frontmost app.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFrontmost(_:)
func (m NSMenuItem) SetAccessibilityFrontmost(accessibilityFrontmost bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityFrontmost:"), accessibilityFrontmost)
}

// Sets the child accessibility element that represents the window’s
// full-screen button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFullScreenButton(_:)
func (m NSMenuItem) SetAccessibilityFullScreenButton(accessibilityFullScreenButton objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityFullScreenButton:"), accessibilityFullScreenButton)
}

// Sets the child accessibility element that represents the window’s grow
// area.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityGrowArea(_:)
func (m NSMenuItem) SetAccessibilityGrowArea(accessibilityGrowArea objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityGrowArea:"), accessibilityGrowArea)
}

// Sets the drag handle accessibility elements for the layout item element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHandles(_:)
func (m NSMenuItem) SetAccessibilityHandles(accessibilityHandles foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityHandles:"), accessibilityHandles)
}

// Sets the header for the table view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHeader(_:)
func (m NSMenuItem) SetAccessibilityHeader(accessibilityHeader objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityHeader:"), accessibilityHeader)
}

// Sets the help text for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHelp(_:)
func (m NSMenuItem) SetAccessibilityHelp(accessibilityHelp string) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityHelp:"), objc.String(accessibilityHelp))
}

// Sets a Boolean value that determines whether the app is in a hidden state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHidden(_:)
func (m NSMenuItem) SetAccessibilityHidden(accessibilityHidden bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityHidden:"), accessibilityHidden)
}

// Sets the horizontal scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHorizontalScrollBar(_:)
func (m NSMenuItem) SetAccessibilityHorizontalScrollBar(accessibilityHorizontalScrollBar objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityHorizontalScrollBar:"), accessibilityHorizontalScrollBar)
}

// Sets the description of the layout area’s horizontal units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHorizontalUnitDescription(_:)
func (m NSMenuItem) SetAccessibilityHorizontalUnitDescription(accessibilityHorizontalUnitDescription string) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityHorizontalUnitDescription:"), objc.String(accessibilityHorizontalUnitDescription))
}

// Sets the units that the layout area uses for horizontal values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHorizontalUnits(_:)
func (m NSMenuItem) SetAccessibilityHorizontalUnits(accessibilityHorizontalUnits NSAccessibilityUnits) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityHorizontalUnits:"), accessibilityHorizontalUnits)
}

// Sets the accessibility element’s identity.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityIdentifier(_:)
func (m NSMenuItem) SetAccessibilityIdentifier(accessibilityIdentifier string) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityIdentifier:"), objc.String(accessibilityIdentifier))
}

// Sets the increment button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityIncrementButton(_:)
func (m NSMenuItem) SetAccessibilityIncrementButton(accessibilityIncrementButton objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityIncrementButton:"), accessibilityIncrementButton)
}

// Sets the index of the row or column that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityIndex(_:)
func (m NSMenuItem) SetAccessibilityIndex(accessibilityIndex int) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityIndex:"), accessibilityIndex)
}

// Sets the line number that contains the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityInsertionPointLineNumber(_:)
func (m NSMenuItem) SetAccessibilityInsertionPointLineNumber(accessibilityInsertionPointLineNumber int) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityInsertionPointLineNumber:"), accessibilityInsertionPointLineNumber)
}

// Sets a short description of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLabel(_:)
func (m NSMenuItem) SetAccessibilityLabel(accessibilityLabel string) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityLabel:"), objc.String(accessibilityLabel))
}

// Sets the child label elements for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLabelUIElements(_:)
func (m NSMenuItem) SetAccessibilityLabelUIElements(accessibilityLabelUIElements foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityLabelUIElements:"), accessibilityLabelUIElements)
}

// Sets the value of the label accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLabelValue(_:)
func (m NSMenuItem) SetAccessibilityLabelValue(accessibilityLabelValue float32) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityLabelValue:"), accessibilityLabelValue)
}

// Sets the elements that have links with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLinkedUIElements(_:)
func (m NSMenuItem) SetAccessibilityLinkedUIElements(accessibilityLinkedUIElements foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityLinkedUIElements:"), accessibilityLinkedUIElements)
}

// Sets a Boolean value that determines whether the window is the app’s main
// window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMain(_:)
func (m NSMenuItem) SetAccessibilityMain(accessibilityMain bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityMain:"), accessibilityMain)
}

// Sets the app’s main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMainWindow(_:)
func (m NSMenuItem) SetAccessibilityMainWindow(accessibilityMainWindow objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityMainWindow:"), accessibilityMainWindow)
}

// Sets the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerGroupUIElement(_:)
func (m NSMenuItem) SetAccessibilityMarkerGroupUIElement(accessibilityMarkerGroupUIElement objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityMarkerGroupUIElement:"), accessibilityMarkerGroupUIElement)
}

// Sets the human-readable description of the marker type.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerTypeDescription(_:)
func (m NSMenuItem) SetAccessibilityMarkerTypeDescription(accessibilityMarkerTypeDescription string) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityMarkerTypeDescription:"), objc.String(accessibilityMarkerTypeDescription))
}

// Sets the array of marker accessibility elements for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerUIElements(_:)
func (m NSMenuItem) SetAccessibilityMarkerUIElements(accessibilityMarkerUIElements foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityMarkerUIElements:"), accessibilityMarkerUIElements)
}

// Sets the marker values for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerValues(_:)
func (m NSMenuItem) SetAccessibilityMarkerValues(accessibilityMarkerValues objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityMarkerValues:"), accessibilityMarkerValues)
}

// Sets the maximum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMaxValue(_:)
func (m NSMenuItem) SetAccessibilityMaxValue(accessibilityMaxValue objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityMaxValue:"), accessibilityMaxValue)
}

// Sets the app’s menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMenuBar(_:)
func (m NSMenuItem) SetAccessibilityMenuBar(accessibilityMenuBar objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityMenuBar:"), accessibilityMenuBar)
}

// Sets the minimum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMinValue(_:)
func (m NSMenuItem) SetAccessibilityMinValue(accessibilityMinValue objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityMinValue:"), accessibilityMinValue)
}

// Sets the child accessibility element that represents the window’s
// minimize button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMinimizeButton(_:)
func (m NSMenuItem) SetAccessibilityMinimizeButton(accessibilityMinimizeButton objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityMinimizeButton:"), accessibilityMinimizeButton)
}

// Sets the Boolean value that determines whether the window is in a minimized
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMinimized(_:)
func (m NSMenuItem) SetAccessibilityMinimized(accessibilityMinimized bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityMinimized:"), accessibilityMinimized)
}

// Sets a Boolean value that determines whether the window is modal.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityModal(_:)
func (m NSMenuItem) SetAccessibilityModal(accessibilityModal bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityModal:"), accessibilityModal)
}

// Sets the contents that follow the divider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityNextContents(_:)
func (m NSMenuItem) SetAccessibilityNextContents(accessibilityNextContents foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityNextContents:"), accessibilityNextContents)
}

// Sets the number of characters in the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityNumberOfCharacters(_:)
func (m NSMenuItem) SetAccessibilityNumberOfCharacters(accessibilityNumberOfCharacters int) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityNumberOfCharacters:"), accessibilityNumberOfCharacters)
}

// Sets a Boolean value that determines whether the element’s grid is in row
// major order or in column major order.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityOrderedByRow(_:)
func (m NSMenuItem) SetAccessibilityOrderedByRow(accessibilityOrderedByRow bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityOrderedByRow:"), accessibilityOrderedByRow)
}

// Sets the orientation of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityOrientation(_:)
func (m NSMenuItem) SetAccessibilityOrientation(accessibilityOrientation NSAccessibilityOrientation) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityOrientation:"), accessibilityOrientation)
}

// Sets the overflow button for the toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityOverflowButton(_:)
func (m NSMenuItem) SetAccessibilityOverflowButton(accessibilityOverflowButton objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityOverflowButton:"), accessibilityOverflowButton)
}

// Sets the accessibility element’s parent in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityParent(_:)
func (m NSMenuItem) SetAccessibilityParent(accessibilityParent objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityParent:"), accessibilityParent)
}

// Sets the placeholder value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityPlaceholderValue(_:)
func (m NSMenuItem) SetAccessibilityPlaceholderValue(accessibilityPlaceholderValue string) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityPlaceholderValue:"), objc.String(accessibilityPlaceholderValue))
}

// Sets the contents that precede the divider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityPreviousContents(_:)
func (m NSMenuItem) SetAccessibilityPreviousContents(accessibilityPreviousContents foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityPreviousContents:"), accessibilityPreviousContents)
}

// Sets a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityProtectedContent(_:)
func (m NSMenuItem) SetAccessibilityProtectedContent(accessibilityProtectedContent bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityProtectedContent:"), accessibilityProtectedContent)
}

// Sets the child accessibility element that represents the window’s proxy
// icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityProxy(_:)
func (m NSMenuItem) SetAccessibilityProxy(accessibilityProxy objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityProxy:"), accessibilityProxy)
}

// Sets a Boolean value that determines whether the accessibility element must
// have content for successful submission of a form.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRequired(_:)
func (m NSMenuItem) SetAccessibilityRequired(accessibilityRequired bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityRequired:"), accessibilityRequired)
}

// Sets the type of interface element that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRole(_:)
func (m NSMenuItem) SetAccessibilityRole(accessibilityRole NSAccessibilityRole) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityRole:"), objc.String(string(accessibilityRole)))
}

// Sets the localized, human-intelligible description of the accessibility
// element’s role, such as radio button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRoleDescription(_:)
func (m NSMenuItem) SetAccessibilityRoleDescription(accessibilityRoleDescription string) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityRoleDescription:"), objc.String(accessibilityRoleDescription))
}

// Sets the number of rows in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRowCount(_:)
func (m NSMenuItem) SetAccessibilityRowCount(accessibilityRowCount int) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityRowCount:"), accessibilityRowCount)
}

// Sets the row header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRowHeaderUIElements(_:)
func (m NSMenuItem) SetAccessibilityRowHeaderUIElements(accessibilityRowHeaderUIElements foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityRowHeaderUIElements:"), accessibilityRowHeaderUIElements)
}

// Sets the row index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRowIndexRange(_:)
func (m NSMenuItem) SetAccessibilityRowIndexRange(accessibilityRowIndexRange foundation.NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityRowIndexRange:"), accessibilityRowIndexRange)
}

// Sets the row accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRows(_:)
func (m NSMenuItem) SetAccessibilityRows(accessibilityRows foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityRows:"), accessibilityRows)
}

// Sets the type of markers for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRulerMarkerType(_:)
func (m NSMenuItem) SetAccessibilityRulerMarkerType(accessibilityRulerMarkerType NSAccessibilityRulerMarkerType) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityRulerMarkerType:"), accessibilityRulerMarkerType)
}

// Sets the search button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySearchButton(_:)
func (m NSMenuItem) SetAccessibilitySearchButton(accessibilitySearchButton objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilitySearchButton:"), accessibilitySearchButton)
}

// Sets the search menu for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySearchMenu(_:)
func (m NSMenuItem) SetAccessibilitySearchMenu(accessibilitySearchMenu objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilitySearchMenu:"), accessibilitySearchMenu)
}

// Sets a Boolean value that determines whether the accessibility element is
// currently in a selected state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelected(_:)
func (m NSMenuItem) SetAccessibilitySelected(accessibilitySelected bool) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilitySelected:"), accessibilitySelected)
}

// Sets the currently selected cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedCells(_:)
func (m NSMenuItem) SetAccessibilitySelectedCells(accessibilitySelectedCells foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilitySelectedCells:"), accessibilitySelectedCells)
}

// Sets the accessibility element’s currently selected children.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedChildren(_:)
func (m NSMenuItem) SetAccessibilitySelectedChildren(accessibilitySelectedChildren foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilitySelectedChildren:"), accessibilitySelectedChildren)
}

// Sets the currently selected columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedColumns(_:)
func (m NSMenuItem) SetAccessibilitySelectedColumns(accessibilitySelectedColumns foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilitySelectedColumns:"), accessibilitySelectedColumns)
}

// Sets the currently selected rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedRows(_:)
func (m NSMenuItem) SetAccessibilitySelectedRows(accessibilitySelectedRows foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilitySelectedRows:"), accessibilitySelectedRows)
}

// Sets the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedText(_:)
func (m NSMenuItem) SetAccessibilitySelectedText(accessibilitySelectedText string) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilitySelectedText:"), objc.String(accessibilitySelectedText))
}

// Sets the range of the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedTextRange(_:)
func (m NSMenuItem) SetAccessibilitySelectedTextRange(accessibilitySelectedTextRange foundation.NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilitySelectedTextRange:"), accessibilitySelectedTextRange)
}

// Sets an array of ranges for the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedTextRanges(_:)
func (m NSMenuItem) SetAccessibilitySelectedTextRanges(accessibilitySelectedTextRanges []foundation.NSValue) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilitySelectedTextRanges:"), objectivec.IObjectSliceToNSArray(accessibilitySelectedTextRanges))
}

// Sets the list of elements that the accessibility element is a title for.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityServesAsTitleForUIElements(_:)
func (m NSMenuItem) SetAccessibilityServesAsTitleForUIElements(accessibilityServesAsTitleForUIElements foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityServesAsTitleForUIElements:"), accessibilityServesAsTitleForUIElements)
}

// Sets the range of characters that the accessibility element displays.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySharedCharacterRange(_:)
func (m NSMenuItem) SetAccessibilitySharedCharacterRange(accessibilitySharedCharacterRange foundation.NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilitySharedCharacterRange:"), accessibilitySharedCharacterRange)
}

// Sets the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySharedFocusElements(_:)
func (m NSMenuItem) SetAccessibilitySharedFocusElements(accessibilitySharedFocusElements foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilitySharedFocusElements:"), accessibilitySharedFocusElements)
}

// Sets the other elements that share text with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySharedTextUIElements(_:)
func (m NSMenuItem) SetAccessibilitySharedTextUIElements(accessibilitySharedTextUIElements foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilitySharedTextUIElements:"), accessibilitySharedTextUIElements)
}

// Sets the menu currently displaying for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityShownMenu(_:)
func (m NSMenuItem) SetAccessibilityShownMenu(accessibilityShownMenu objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityShownMenu:"), accessibilityShownMenu)
}

// Sets the accessibility element’s sort direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySortDirection(_:)
func (m NSMenuItem) SetAccessibilitySortDirection(accessibilitySortDirection NSAccessibilitySortDirection) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilitySortDirection:"), accessibilitySortDirection)
}

// Sets the array that contains the views and splitter bar from the split
// view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySplitters(_:)
func (m NSMenuItem) SetAccessibilitySplitters(accessibilitySplitters foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilitySplitters:"), accessibilitySplitters)
}

// Sets the specialized interface element type that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySubrole(_:)
func (m NSMenuItem) SetAccessibilitySubrole(accessibilitySubrole NSAccessibilitySubrole) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilitySubrole:"), objc.String(string(accessibilitySubrole)))
}

// Sets the tab accessibility elements for the tab view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTabs(_:)
func (m NSMenuItem) SetAccessibilityTabs(accessibilityTabs foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityTabs:"), accessibilityTabs)
}

// Sets the title of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTitle(_:)
func (m NSMenuItem) SetAccessibilityTitle(accessibilityTitle string) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityTitle:"), objc.String(accessibilityTitle))
}

// Sets the static text element that represents the accessibility element’s
// title.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTitleUIElement(_:)
func (m NSMenuItem) SetAccessibilityTitleUIElement(accessibilityTitleUIElement objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityTitleUIElement:"), accessibilityTitleUIElement)
}

// Sets the child accessibility element that represents the window’s toolbar
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityToolbarButton(_:)
func (m NSMenuItem) SetAccessibilityToolbarButton(accessibilityToolbarButton objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityToolbarButton:"), accessibilityToolbarButton)
}

// Sets the top-level element that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTopLevelUIElement(_:)
func (m NSMenuItem) SetAccessibilityTopLevelUIElement(accessibilityTopLevelUIElement objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityTopLevelUIElement:"), accessibilityTopLevelUIElement)
}

// Sets the URL for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityURL(_:)
func (m NSMenuItem) SetAccessibilityURL(accessibilityURL foundation.NSURL) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityURL:"), accessibilityURL)
}

// Sets the human-readable description of the ruler’s units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityUnitDescription(_:)
func (m NSMenuItem) SetAccessibilityUnitDescription(accessibilityUnitDescription string) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityUnitDescription:"), objc.String(accessibilityUnitDescription))
}

// Sets the units used for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityUnits(_:)
func (m NSMenuItem) SetAccessibilityUnits(accessibilityUnits NSAccessibilityUnits) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityUnits:"), accessibilityUnits)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityUserInputLabels(_:)
func (m NSMenuItem) SetAccessibilityUserInputLabels(accessibilityUserInputLabels []string) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityUserInputLabels:"), objectivec.StringSliceToNSArray(accessibilityUserInputLabels))
}

// Sets the accessibility element’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityValue(_:)
func (m NSMenuItem) SetAccessibilityValue(accessibilityValue objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityValue:"), accessibilityValue)
}

// Sets the human-readable description of the accessibility element’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityValueDescription(_:)
func (m NSMenuItem) SetAccessibilityValueDescription(accessibilityValueDescription string) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityValueDescription:"), objc.String(accessibilityValueDescription))
}

// Sets the vertical scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVerticalScrollBar(_:)
func (m NSMenuItem) SetAccessibilityVerticalScrollBar(accessibilityVerticalScrollBar objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityVerticalScrollBar:"), accessibilityVerticalScrollBar)
}

// Sets the description of the layout area’s vertical units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVerticalUnitDescription(_:)
func (m NSMenuItem) SetAccessibilityVerticalUnitDescription(accessibilityVerticalUnitDescription string) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityVerticalUnitDescription:"), objc.String(accessibilityVerticalUnitDescription))
}

// Sets the units that the layout area uses for vertical values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVerticalUnits(_:)
func (m NSMenuItem) SetAccessibilityVerticalUnits(accessibilityVerticalUnits NSAccessibilityUnits) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityVerticalUnits:"), accessibilityVerticalUnits)
}

// Sets the visible cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleCells(_:)
func (m NSMenuItem) SetAccessibilityVisibleCells(accessibilityVisibleCells foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityVisibleCells:"), accessibilityVisibleCells)
}

// Sets the range of visible characters in the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleCharacterRange(_:)
func (m NSMenuItem) SetAccessibilityVisibleCharacterRange(accessibilityVisibleCharacterRange foundation.NSRange) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityVisibleCharacterRange:"), accessibilityVisibleCharacterRange)
}

// Sets the accessibility element’s visible child accessibility elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleChildren(_:)
func (m NSMenuItem) SetAccessibilityVisibleChildren(accessibilityVisibleChildren foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityVisibleChildren:"), accessibilityVisibleChildren)
}

// Sets the visible columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleColumns(_:)
func (m NSMenuItem) SetAccessibilityVisibleColumns(accessibilityVisibleColumns foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityVisibleColumns:"), accessibilityVisibleColumns)
}

// Sets the visible rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleRows(_:)
func (m NSMenuItem) SetAccessibilityVisibleRows(accessibilityVisibleRows foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityVisibleRows:"), accessibilityVisibleRows)
}

// Sets the warning value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityWarningValue(_:)
func (m NSMenuItem) SetAccessibilityWarningValue(accessibilityWarningValue objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityWarningValue:"), accessibilityWarningValue)
}

// Sets the window that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityWindow(_:)
func (m NSMenuItem) SetAccessibilityWindow(accessibilityWindow objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityWindow:"), accessibilityWindow)
}

// Sets the array that contains all the app’s windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityWindows(_:)
func (m NSMenuItem) SetAccessibilityWindows(accessibilityWindows foundation.INSArray) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityWindows:"), accessibilityWindows)
}

// Sets the child accessibility element that represents the window’s zoom
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityZoomButton(_:)
func (m NSMenuItem) SetAccessibilityZoomButton(accessibilityZoomButton objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("setAccessibilityZoomButton:"), accessibilityZoomButton)
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

// Protocol methods for NSAccessibilityProtocol

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
