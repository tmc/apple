// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPopover] class.
var (
	_NSPopoverClass     NSPopoverClass
	_NSPopoverClassOnce sync.Once
)

func getNSPopoverClass() NSPopoverClass {
	_NSPopoverClassOnce.Do(func() {
		_NSPopoverClass = NSPopoverClass{class: objc.GetClass("NSPopover")}
	})
	return _NSPopoverClass
}

// GetNSPopoverClass returns the class object for NSPopover.
func GetNSPopoverClass() NSPopoverClass {
	return getNSPopoverClass()
}

type NSPopoverClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPopoverClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPopoverClass) Alloc() NSPopover {
	rv := objc.Send[NSPopover](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A means to display additional content related to existing content on the
// screen.
//
// # Overview
//
// The popover is positioned relative to the existing content and an anchor is
// used to express the relation between these two units of content. A popover
// has an appearance that specifies its visual characteristics, as well as a
// behavior that determines which user interactions will cause the popover to
// close. A transient popover is closed in response to most user interactions,
// whereas a semi-transient popover is closed when the user interacts with the
// window containing the popover’s positioning view. Popovers with
// application-defined behavior are not usually closed on the developer’s
// behalf.
//
// The system automatically positions each popover relative to its positioning
// view and moves the popover whenever its positioning view moves. A
// positioning rectangle within the positioning view can be specified for
// additional granularity.
//
// Popovers can be detached to become a separate window when they are dragged
// by implementing the appropriate delegate method.
//
// # Accessing a Popover’s Content View Controller
//
//   - [NSPopover.ContentViewController]: The view controller that manages the content of the popover.
//   - [NSPopover.SetContentViewController]
//
// # Managing a Popover’s Position and Size
//
//   - [NSPopover.Behavior]: Specifies the behavior of the popover.
//   - [NSPopover.SetBehavior]
//   - [NSPopover.ShowRelativeToRectOfViewPreferredEdge]: Shows the popover anchored to the specified view.
//   - [NSPopover.PositioningRect]: The rectangle within the positioning view relative to which the popover should be positioned.
//   - [NSPopover.SetPositioningRect]
//
// # Managing a Popover’s Appearance
//
//   - [NSPopover.Animates]: Specifies if the popover is to be animated.
//   - [NSPopover.SetAnimates]
//   - [NSPopover.ContentSize]: The content size of the popover.
//   - [NSPopover.SetContentSize]
//   - [NSPopover.IsShown]: The display state of the popover.
//   - [NSPopover.IsDetached]: A Boolean value that indicates whether the window created by a popover’s detachment is automatically created.
//
// # Closing a Popover
//
//   - [NSPopover.PerformClose]: Attempts to close the popover.
//   - [NSPopover.Close]: Forces the popover to close without consulting its delegate.
//
// # Getting and Setting the Delegate
//
//   - [NSPopover.Delegate]: The delegate of the popover.
//   - [NSPopover.SetDelegate]
//
// # Instance Properties
//
//   - [NSPopover.HasFullSizeContent]: A Boolean value that indicates whether the content view of the popover extends into the arrow region.
//   - [NSPopover.SetHasFullSizeContent]
//
// # Instance Methods
//
//   - [NSPopover.ShowRelativeToToolbarItem]: Shows the popover anchored to the specified toolbar item.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover
type NSPopover struct {
	NSResponder
}

// NSPopoverFromID constructs a [NSPopover] from an objc.ID.
//
// A means to display additional content related to existing content on the
// screen.
func NSPopoverFromID(id objc.ID) NSPopover {
	return NSPopover{NSResponder: NSResponderFromID(id)}
}

// NOTE: NSPopover adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPopover] class.
//
// # Accessing a Popover’s Content View Controller
//
//   - [INSPopover.ContentViewController]: The view controller that manages the content of the popover.
//   - [INSPopover.SetContentViewController]
//
// # Managing a Popover’s Position and Size
//
//   - [INSPopover.Behavior]: Specifies the behavior of the popover.
//   - [INSPopover.SetBehavior]
//   - [INSPopover.ShowRelativeToRectOfViewPreferredEdge]: Shows the popover anchored to the specified view.
//   - [INSPopover.PositioningRect]: The rectangle within the positioning view relative to which the popover should be positioned.
//   - [INSPopover.SetPositioningRect]
//
// # Managing a Popover’s Appearance
//
//   - [INSPopover.Animates]: Specifies if the popover is to be animated.
//   - [INSPopover.SetAnimates]
//   - [INSPopover.ContentSize]: The content size of the popover.
//   - [INSPopover.SetContentSize]
//   - [INSPopover.IsShown]: The display state of the popover.
//   - [INSPopover.IsDetached]: A Boolean value that indicates whether the window created by a popover’s detachment is automatically created.
//
// # Closing a Popover
//
//   - [INSPopover.PerformClose]: Attempts to close the popover.
//   - [INSPopover.Close]: Forces the popover to close without consulting its delegate.
//
// # Getting and Setting the Delegate
//
//   - [INSPopover.Delegate]: The delegate of the popover.
//   - [INSPopover.SetDelegate]
//
// # Instance Properties
//
//   - [INSPopover.HasFullSizeContent]: A Boolean value that indicates whether the content view of the popover extends into the arrow region.
//   - [INSPopover.SetHasFullSizeContent]
//
// # Instance Methods
//
//   - [INSPopover.ShowRelativeToToolbarItem]: Shows the popover anchored to the specified toolbar item.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover
type INSPopover interface {
	INSResponder
	NSAppearanceCustomization

	// Topic: Accessing a Popover’s Content View Controller

	// The view controller that manages the content of the popover.
	ContentViewController() INSViewController
	SetContentViewController(value INSViewController)

	// Topic: Managing a Popover’s Position and Size

	// Specifies the behavior of the popover.
	Behavior() NSPopoverBehavior
	SetBehavior(value NSPopoverBehavior)
	// Shows the popover anchored to the specified view.
	ShowRelativeToRectOfViewPreferredEdge(positioningRect corefoundation.CGRect, positioningView INSView, preferredEdge foundation.NSRectEdge)
	// The rectangle within the positioning view relative to which the popover should be positioned.
	PositioningRect() corefoundation.CGRect
	SetPositioningRect(value corefoundation.CGRect)

	// Topic: Managing a Popover’s Appearance

	// Specifies if the popover is to be animated.
	Animates() bool
	SetAnimates(value bool)
	// The content size of the popover.
	ContentSize() corefoundation.CGSize
	SetContentSize(value corefoundation.CGSize)
	// The display state of the popover.
	IsShown() bool
	// A Boolean value that indicates whether the window created by a popover’s detachment is automatically created.
	IsDetached() bool

	// Topic: Closing a Popover

	// Attempts to close the popover.
	PerformClose(sender objectivec.IObject)
	// Forces the popover to close without consulting its delegate.
	Close()

	// Topic: Getting and Setting the Delegate

	// The delegate of the popover.
	Delegate() NSPopoverDelegate
	SetDelegate(value NSPopoverDelegate)

	// Topic: Instance Properties

	// A Boolean value that indicates whether the content view of the popover extends into the arrow region.
	HasFullSizeContent() bool
	SetHasFullSizeContent(value bool)

	// Topic: Instance Methods

	// Shows the popover anchored to the specified toolbar item.
	ShowRelativeToToolbarItem(toolbarItem INSToolbarItem)

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
}

// Init initializes the instance.
func (p NSPopover) Init() NSPopover {
	rv := objc.Send[NSPopover](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPopover) Autorelease() NSPopover {
	rv := objc.Send[NSPopover](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPopover creates a new NSPopover instance.
func NewNSPopover() NSPopover {
	class := getNSPopoverClass()
	rv := objc.Send[NSPopover](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSPopover/init(coder:)
func NewPopoverWithCoder(coder foundation.INSCoder) NSPopover {
	instance := getNSPopoverClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSPopoverFromID(rv)
}

// Shows the popover anchored to the specified view.
//
// positioningRect: The rectangle within `positioningView` relative to which the popover should
// be positioned. Normally set to the bounds of `positioningView`. May be an
// empty rectangle, which will default to the bounds of `positioningView`.
//
// positioningView: The view relative to which the popover should be positioned. Causes the
// method to raise [invalidArgumentException] if `nil`.
//
// preferredEdge: The edge of `positioningView` the popover should prefer to be anchored to.
//
// # Discussion
//
// This method raises [internalInconsistencyException] if
// [NSPopover.ContentViewController] or the view controller’s view is `nil`.
// If the popover is already being shown, this method updates the anchored
// view, rectangle, and preferred edge. If the positioning view is not
// visible, this method does nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/show(relativeTo:of:preferredEdge:)
//
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
// [internalInconsistencyException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/internalInconsistencyException
func (p NSPopover) ShowRelativeToRectOfViewPreferredEdge(positioningRect corefoundation.CGRect, positioningView INSView, preferredEdge foundation.NSRectEdge) {
	objc.Send[objc.ID](p.ID, objc.Sel("showRelativeToRect:ofView:preferredEdge:"), positioningRect, positioningView, preferredEdge)
}

// Attempts to close the popover.
//
// sender: The sender of the action message.
//
// # Discussion
//
// The popover will not be closed if it has a delegate and the delegate
// implements the returns [PopoverShouldClose] method returning false, or if a
// subclass of the NSPopover class implements “ and returns false).
//
// The operation will fail if the popover is displaying a nested popover or if
// it has a child window. A window will attempt to close its popovers when it
// receives a [NSWindow.PerformClose] message.
//
// The popover animates out when closed unless the [NSPopover.Animates]
// property is set to false.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/performClose(_:)
func (p NSPopover) PerformClose(sender objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("performClose:"), sender)
}

// Forces the popover to close without consulting its delegate.
//
// # Discussion
//
// Any popovers nested within the popovers will also receive a
// [NSPopover.Close] message. When a window is closed in response to the
// [NSWindow.Close] message being sent, all of its popovers are closed. The
// popover animates out when closed unless the [NSPopover.Animates] property
// is set to false.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/close()
func (p NSPopover) Close() {
	objc.Send[objc.ID](p.ID, objc.Sel("close"))
}

// Shows the popover anchored to the specified toolbar item.
//
// toolbarItem: The toolbar item anchoring the popover.
//
// # Discussion
//
// Use this method to display a popover relative to a toolbar item. When the
// item is in the overflow menu, the popover presents itself from another
// appropriate affordance in the window. See
// [NSPopover.ShowRelativeToRectOfViewPreferredEdge] for popover behavior.
//
// This method raises an [invalidArgumentException] if it can’t locate the
// toolbar item. This could occur if the item isn’t in a toolbar, or because
// the toolbar isn’t in the window.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/show(relativeTo:)
//
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
func (p NSPopover) ShowRelativeToToolbarItem(toolbarItem INSToolbarItem) {
	objc.Send[objc.ID](p.ID, objc.Sel("showRelativeToToolbarItem:"), toolbarItem)
}

// Returns the activation point for the user interface element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityActivationPoint()
func (p NSPopover) AccessibilityActivationPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](p.ID, objc.Sel("accessibilityActivationPoint"))
	return corefoundation.CGPoint(rv)
}

// Returns the allowed values for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityAllowedValues()
func (p NSPopover) AccessibilityAllowedValues() []foundation.NSNumber {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("accessibilityAllowedValues"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSNumber {
		return foundation.NSNumberFromID(id)
	})
}

// Returns the child accessibility element with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityApplicationFocusedUIElement()
func (p NSPopover) AccessibilityApplicationFocusedUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityApplicationFocusedUIElement"))
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
func (p NSPopover) AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityAttributedStringForRange:"), range_)
	return foundation.NSAttributedStringFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityAttributedUserInputLabels()
func (p NSPopover) AccessibilityAttributedUserInputLabels() []foundation.NSAttributedString {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("accessibilityAttributedUserInputLabels"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSAttributedString {
		return foundation.NSAttributedStringFromID(id)
	})
}

// Returns the child accessibility element that represents the window’s
// cancel button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCancelButton()
func (p NSPopover) AccessibilityCancelButton() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityCancelButton"))
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
func (p NSPopover) AccessibilityCellForColumnRow(column int, row int) objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityCellForColumn:row:"), column, row)
	return objectivec.Object{ID: rv}
}

// Returns the child accessibility elements in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityChildren()
func (p NSPopover) AccessibilityChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityChildren"))
	return foundation.NSArrayFromID(rv)
}

// Returns the array of child accessibility elements in order for linear
// navigation.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityChildrenInNavigationOrder()
func (p NSPopover) AccessibilityChildrenInNavigationOrder() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("accessibilityChildrenInNavigationOrder"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns the clear button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityClearButton()
func (p NSPopover) AccessibilityClearButton() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityClearButton"))
	return objectivec.Object{ID: rv}
}

// Returns the child accessibility element that represents the window’s
// close button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCloseButton()
func (p NSPopover) AccessibilityCloseButton() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityCloseButton"))
	return objectivec.Object{ID: rv}
}

// Returns the number of columns in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnCount()
func (p NSPopover) AccessibilityColumnCount() int {
	rv := objc.Send[int](p.ID, objc.Sel("accessibilityColumnCount"))
	return rv
}

// Returns the column header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnHeaderUIElements()
func (p NSPopover) AccessibilityColumnHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityColumnHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the column index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnIndexRange()
func (p NSPopover) AccessibilityColumnIndexRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](p.ID, objc.Sel("accessibilityColumnIndexRange"))
	return foundation.NSRange(rv)
}

// Returns the column titles for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumnTitles()
func (p NSPopover) AccessibilityColumnTitles() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityColumnTitles"))
	return foundation.NSArrayFromID(rv)
}

// Returns the column accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityColumns()
func (p NSPopover) AccessibilityColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityColumns"))
	return foundation.NSArrayFromID(rv)
}

// Returns the contents of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityContents()
func (p NSPopover) AccessibilityContents() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityContents"))
	return foundation.NSArrayFromID(rv)
}

// Returns the critical value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCriticalValue()
func (p NSPopover) AccessibilityCriticalValue() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityCriticalValue"))
	return objectivec.Object{ID: rv}
}

// Returns the custom actions of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCustomActions()
func (p NSPopover) AccessibilityCustomActions() []NSAccessibilityCustomAction {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("accessibilityCustomActions"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSAccessibilityCustomAction {
		return NSAccessibilityCustomActionFromID(id)
	})
}

// Returns the custom rotors of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityCustomRotors()
func (p NSPopover) AccessibilityCustomRotors() []NSAccessibilityCustomRotor {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("accessibilityCustomRotors"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSAccessibilityCustomRotor {
		return NSAccessibilityCustomRotorFromID(id)
	})
}

// Returns the decrement button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDecrementButton()
func (p NSPopover) AccessibilityDecrementButton() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityDecrementButton"))
	return objectivec.Object{ID: rv}
}

// Returns the child accessibility element that represents the window’s
// default button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDefaultButton()
func (p NSPopover) AccessibilityDefaultButton() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityDefaultButton"))
	return objectivec.Object{ID: rv}
}

// Returns the row disclosing the current row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDisclosedByRow()
func (p NSPopover) AccessibilityDisclosedByRow() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityDisclosedByRow"))
	return objectivec.Object{ID: rv}
}

// Returns the rows that the current row discloses.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDisclosedRows()
func (p NSPopover) AccessibilityDisclosedRows() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityDisclosedRows"))
	return objectivec.Object{ID: rv}
}

// Returns the indention level for the row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDisclosureLevel()
func (p NSPopover) AccessibilityDisclosureLevel() int {
	rv := objc.Send[int](p.ID, objc.Sel("accessibilityDisclosureLevel"))
	return rv
}

// Returns the URL for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityDocument()
func (p NSPopover) AccessibilityDocument() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityDocument"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the icon for the app’s menu bar extra.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityExtrasMenuBar()
func (p NSPopover) AccessibilityExtrasMenuBar() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityExtrasMenuBar"))
	return objectivec.Object{ID: rv}
}

// Returns the filename for the file that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFilename()
func (p NSPopover) AccessibilityFilename() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityFilename"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the child window with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFocusedWindow()
func (p NSPopover) AccessibilityFocusedWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityFocusedWindow"))
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
func (p NSPopover) AccessibilityFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p.ID, objc.Sel("accessibilityFrame"))
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
func (p NSPopover) AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p.ID, objc.Sel("accessibilityFrameForRange:"), range_)
	return corefoundation.CGRect(rv)
}

// Returns the child accessibility element that represents the window’s
// full-screen button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFullScreenButton()
func (p NSPopover) AccessibilityFullScreenButton() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityFullScreenButton"))
	return objectivec.Object{ID: rv}
}

// Returns the child accessibility element that represents the window’s grow
// area.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityGrowArea()
func (p NSPopover) AccessibilityGrowArea() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityGrowArea"))
	return objectivec.Object{ID: rv}
}

// Returns the drag handle elements for the layout item element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHandles()
func (p NSPopover) AccessibilityHandles() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityHandles"))
	return foundation.NSArrayFromID(rv)
}

// Returns the header for the table view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHeader()
func (p NSPopover) AccessibilityHeader() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityHeader"))
	return objectivec.Object{ID: rv}
}

// Returns the help text for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHelp()
func (p NSPopover) AccessibilityHelp() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityHelp"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the horizontal scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHorizontalScrollBar()
func (p NSPopover) AccessibilityHorizontalScrollBar() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityHorizontalScrollBar"))
	return objectivec.Object{ID: rv}
}

// Returns the description of the layout area’s horizontal units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHorizontalUnitDescription()
func (p NSPopover) AccessibilityHorizontalUnitDescription() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityHorizontalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the units that the layout area uses for horizontal values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityHorizontalUnits()
func (p NSPopover) AccessibilityHorizontalUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](p.ID, objc.Sel("accessibilityHorizontalUnits"))
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
func (p NSPopover) AccessibilityIdentifier() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the increment button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityIncrementButton()
func (p NSPopover) AccessibilityIncrementButton() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityIncrementButton"))
	return objectivec.Object{ID: rv}
}

// Returns the index of the row or column that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityIndex()
func (p NSPopover) AccessibilityIndex() int {
	rv := objc.Send[int](p.ID, objc.Sel("accessibilityIndex"))
	return rv
}

// Returns the line number that contains the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityInsertionPointLineNumber()
func (p NSPopover) AccessibilityInsertionPointLineNumber() int {
	rv := objc.Send[int](p.ID, objc.Sel("accessibilityInsertionPointLineNumber"))
	return rv
}

// Returns a short description of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLabel()
func (p NSPopover) AccessibilityLabel() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the child label elements for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLabelUIElements()
func (p NSPopover) AccessibilityLabelUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityLabelUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the value of the label accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLabelValue()
func (p NSPopover) AccessibilityLabelValue() float32 {
	rv := objc.Send[float32](p.ID, objc.Sel("accessibilityLabelValue"))
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
func (p NSPopover) AccessibilityLayoutPointForScreenPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](p.ID, objc.Sel("accessibilityLayoutPointForScreenPoint:"), point)
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
func (p NSPopover) AccessibilityLayoutSizeForScreenSize(size corefoundation.CGSize) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](p.ID, objc.Sel("accessibilityLayoutSizeForScreenSize:"), size)
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
func (p NSPopover) AccessibilityLineForIndex(index int) int {
	rv := objc.Send[int](p.ID, objc.Sel("accessibilityLineForIndex:"), index)
	return rv
}

// Returns the elements that have links with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLinkedUIElements()
func (p NSPopover) AccessibilityLinkedUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityLinkedUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the app’s main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMainWindow()
func (p NSPopover) AccessibilityMainWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityMainWindow"))
	return objectivec.Object{ID: rv}
}

// Returns the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerGroupUIElement()
func (p NSPopover) AccessibilityMarkerGroupUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityMarkerGroupUIElement"))
	return objectivec.Object{ID: rv}
}

// Returns the human-readable description of the marker type.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerTypeDescription()
func (p NSPopover) AccessibilityMarkerTypeDescription() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityMarkerTypeDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the array of marker accessibility elements for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerUIElements()
func (p NSPopover) AccessibilityMarkerUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityMarkerUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the marker values for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMarkerValues()
func (p NSPopover) AccessibilityMarkerValues() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityMarkerValues"))
	return objectivec.Object{ID: rv}
}

// Returns the maximum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMaxValue()
func (p NSPopover) AccessibilityMaxValue() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityMaxValue"))
	return objectivec.Object{ID: rv}
}

// Returns the app’s menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMenuBar()
func (p NSPopover) AccessibilityMenuBar() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityMenuBar"))
	return objectivec.Object{ID: rv}
}

// Returns the minimum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMinValue()
func (p NSPopover) AccessibilityMinValue() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityMinValue"))
	return objectivec.Object{ID: rv}
}

// Returns the child accessibility element that represents the window’s
// minimize button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityMinimizeButton()
func (p NSPopover) AccessibilityMinimizeButton() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityMinimizeButton"))
	return objectivec.Object{ID: rv}
}

// Returns the contents that follow the divider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityNextContents()
func (p NSPopover) AccessibilityNextContents() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityNextContents"))
	return foundation.NSArrayFromID(rv)
}

// Returns the number of characters in the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityNumberOfCharacters()
func (p NSPopover) AccessibilityNumberOfCharacters() int {
	rv := objc.Send[int](p.ID, objc.Sel("accessibilityNumberOfCharacters"))
	return rv
}

// Returns the orientation of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityOrientation()
func (p NSPopover) AccessibilityOrientation() NSAccessibilityOrientation {
	rv := objc.Send[NSAccessibilityOrientation](p.ID, objc.Sel("accessibilityOrientation"))
	return NSAccessibilityOrientation(rv)
}

// Returns the overflow button for the toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityOverflowButton()
func (p NSPopover) AccessibilityOverflowButton() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityOverflowButton"))
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
func (p NSPopover) AccessibilityParent() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityParent"))
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
func (p NSPopover) AccessibilityPerformCancel() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("accessibilityPerformCancel"))
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
func (p NSPopover) AccessibilityPerformConfirm() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("accessibilityPerformConfirm"))
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
func (p NSPopover) AccessibilityPerformDecrement() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("accessibilityPerformDecrement"))
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
func (p NSPopover) AccessibilityPerformDelete() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("accessibilityPerformDelete"))
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
func (p NSPopover) AccessibilityPerformIncrement() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("accessibilityPerformIncrement"))
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
func (p NSPopover) AccessibilityPerformPick() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("accessibilityPerformPick"))
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
func (p NSPopover) AccessibilityPerformPress() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("accessibilityPerformPress"))
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
func (p NSPopover) AccessibilityPerformRaise() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("accessibilityPerformRaise"))
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
func (p NSPopover) AccessibilityPerformShowAlternateUI() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("accessibilityPerformShowAlternateUI"))
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
func (p NSPopover) AccessibilityPerformShowDefaultUI() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("accessibilityPerformShowDefaultUI"))
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
func (p NSPopover) AccessibilityPerformShowMenu() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("accessibilityPerformShowMenu"))
	return rv
}

// Returns the placeholder value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPlaceholderValue()
func (p NSPopover) AccessibilityPlaceholderValue() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityPlaceholderValue"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the contents that precede the divider accessibility element.
//
// # Return Value
//
// Sets the contents preceding this divider element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPreviousContents()
func (p NSPopover) AccessibilityPreviousContents() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityPreviousContents"))
	return foundation.NSArrayFromID(rv)
}

// Returns the child accessibility element that represents the window’s
// proxy icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityProxy()
func (p NSPopover) AccessibilityProxy() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityProxy"))
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
func (p NSPopover) AccessibilityRTFForRange(range_ foundation.NSRange) foundation.NSData {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityRTFForRange:"), range_)
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
func (p NSPopover) AccessibilityRangeForIndex(index int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](p.ID, objc.Sel("accessibilityRangeForIndex:"), index)
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
func (p NSPopover) AccessibilityRangeForLine(line int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](p.ID, objc.Sel("accessibilityRangeForLine:"), line)
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
func (p NSPopover) AccessibilityRangeForPosition(point corefoundation.CGPoint) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](p.ID, objc.Sel("accessibilityRangeForPosition:"), point)
	return foundation.NSRange(rv)
}

// Returns the type of interface element that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRole()
func (p NSPopover) AccessibilityRole() NSAccessibilityRole {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityRole"))
	return NSAccessibilityRole(foundation.NSStringFromID(rv).String())
}

// Returns a localized, human-intelligible description of the accessibility
// element’s role, such as radio button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRoleDescription()
func (p NSPopover) AccessibilityRoleDescription() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityRoleDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the number of rows in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRowCount()
func (p NSPopover) AccessibilityRowCount() int {
	rv := objc.Send[int](p.ID, objc.Sel("accessibilityRowCount"))
	return rv
}

// Returns the row header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRowHeaderUIElements()
func (p NSPopover) AccessibilityRowHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityRowHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the row index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRowIndexRange()
func (p NSPopover) AccessibilityRowIndexRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](p.ID, objc.Sel("accessibilityRowIndexRange"))
	return foundation.NSRange(rv)
}

// Returns the row accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRows()
func (p NSPopover) AccessibilityRows() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityRows"))
	return foundation.NSArrayFromID(rv)
}

// Returns the type of markers for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRulerMarkerType()
func (p NSPopover) AccessibilityRulerMarkerType() NSAccessibilityRulerMarkerType {
	rv := objc.Send[NSAccessibilityRulerMarkerType](p.ID, objc.Sel("accessibilityRulerMarkerType"))
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
func (p NSPopover) AccessibilityScreenPointForLayoutPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](p.ID, objc.Sel("accessibilityScreenPointForLayoutPoint:"), point)
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
func (p NSPopover) AccessibilityScreenSizeForLayoutSize(size corefoundation.CGSize) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](p.ID, objc.Sel("accessibilityScreenSizeForLayoutSize:"), size)
	return corefoundation.CGSize(rv)
}

// Returns the search button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySearchButton()
func (p NSPopover) AccessibilitySearchButton() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilitySearchButton"))
	return objectivec.Object{ID: rv}
}

// Returns the search menu for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySearchMenu()
func (p NSPopover) AccessibilitySearchMenu() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilitySearchMenu"))
	return objectivec.Object{ID: rv}
}

// Returns the currently selected cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedCells()
func (p NSPopover) AccessibilitySelectedCells() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilitySelectedCells"))
	return foundation.NSArrayFromID(rv)
}

// Returns the accessibility element’s currently selected children.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedChildren()
func (p NSPopover) AccessibilitySelectedChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilitySelectedChildren"))
	return foundation.NSArrayFromID(rv)
}

// Returns the currently selected columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedColumns()
func (p NSPopover) AccessibilitySelectedColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilitySelectedColumns"))
	return foundation.NSArrayFromID(rv)
}

// Returns the currently selected rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedRows()
func (p NSPopover) AccessibilitySelectedRows() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilitySelectedRows"))
	return foundation.NSArrayFromID(rv)
}

// Returns the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedText()
func (p NSPopover) AccessibilitySelectedText() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilitySelectedText"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the range of the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedTextRange()
func (p NSPopover) AccessibilitySelectedTextRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](p.ID, objc.Sel("accessibilitySelectedTextRange"))
	return foundation.NSRange(rv)
}

// Returns an array of ranges for the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySelectedTextRanges()
func (p NSPopover) AccessibilitySelectedTextRanges() []foundation.NSValue {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("accessibilitySelectedTextRanges"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSValue {
		return foundation.NSValueFromID(id)
	})
}

// Returns the list of elements that the accessibility element is a title for.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityServesAsTitleForUIElements()
func (p NSPopover) AccessibilityServesAsTitleForUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityServesAsTitleForUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the range of characters that the accessibility element displays.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySharedCharacterRange()
func (p NSPopover) AccessibilitySharedCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](p.ID, objc.Sel("accessibilitySharedCharacterRange"))
	return foundation.NSRange(rv)
}

// Returns the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySharedFocusElements()
func (p NSPopover) AccessibilitySharedFocusElements() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilitySharedFocusElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the other elements that share text with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySharedTextUIElements()
func (p NSPopover) AccessibilitySharedTextUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilitySharedTextUIElements"))
	return foundation.NSArrayFromID(rv)
}

// Returns the menu currently displaying for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityShownMenu()
func (p NSPopover) AccessibilityShownMenu() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityShownMenu"))
	return objectivec.Object{ID: rv}
}

// Returns the accessibility element’s sort direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySortDirection()
func (p NSPopover) AccessibilitySortDirection() NSAccessibilitySortDirection {
	rv := objc.Send[NSAccessibilitySortDirection](p.ID, objc.Sel("accessibilitySortDirection"))
	return NSAccessibilitySortDirection(rv)
}

// Returns an array that contains the views and splitter bar from the split
// view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySplitters()
func (p NSPopover) AccessibilitySplitters() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilitySplitters"))
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
func (p NSPopover) AccessibilityStringForRange(range_ foundation.NSRange) string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityStringForRange:"), range_)
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
func (p NSPopover) AccessibilityStyleRangeForIndex(index int) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](p.ID, objc.Sel("accessibilityStyleRangeForIndex:"), index)
	return foundation.NSRange(rv)
}

// Returns the specialized interface element type that the accessibility
// element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilitySubrole()
func (p NSPopover) AccessibilitySubrole() NSAccessibilitySubrole {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilitySubrole"))
	return NSAccessibilitySubrole(foundation.NSStringFromID(rv).String())
}

// Returns the tab accessibility elements for the tab view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTabs()
func (p NSPopover) AccessibilityTabs() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityTabs"))
	return foundation.NSArrayFromID(rv)
}

// Returns the title of the accessibility element—for example, a button’s
// visible text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTitle()
func (p NSPopover) AccessibilityTitle() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityTitle"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the static text element that represents the accessibility
// element’s title.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTitleUIElement()
func (p NSPopover) AccessibilityTitleUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityTitleUIElement"))
	return objectivec.Object{ID: rv}
}

// Returns the child accessibility element that represents the window’s
// toolbar button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityToolbarButton()
func (p NSPopover) AccessibilityToolbarButton() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityToolbarButton"))
	return objectivec.Object{ID: rv}
}

// Returns the top-level element that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityTopLevelUIElement()
func (p NSPopover) AccessibilityTopLevelUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityTopLevelUIElement"))
	return objectivec.Object{ID: rv}
}

// Returns the URL for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityURL()
func (p NSPopover) AccessibilityURL() foundation.NSURL {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityURL"))
	return foundation.NSURLFromID(rv)
}

// Returns the human-readable description of the ruler’s units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityUnitDescription()
func (p NSPopover) AccessibilityUnitDescription() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the units for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityUnits()
func (p NSPopover) AccessibilityUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](p.ID, objc.Sel("accessibilityUnits"))
	return NSAccessibilityUnits(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityUserInputLabels()
func (p NSPopover) AccessibilityUserInputLabels() []string {
	rv := objc.Send[[]objc.ID](p.ID, objc.Sel("accessibilityUserInputLabels"))
	return objc.ConvertSliceToStrings(rv)
}

// Returns the accessibility element’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityValue()
func (p NSPopover) AccessibilityValue() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityValue"))
	return objectivec.Object{ID: rv}
}

// Returns the human-readable description of the accessibility element’s
// value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityValueDescription()
func (p NSPopover) AccessibilityValueDescription() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityValueDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the vertical scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVerticalScrollBar()
func (p NSPopover) AccessibilityVerticalScrollBar() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityVerticalScrollBar"))
	return objectivec.Object{ID: rv}
}

// Returns the description of the layout area’s vertical units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVerticalUnitDescription()
func (p NSPopover) AccessibilityVerticalUnitDescription() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityVerticalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the units that the layout area uses for vertical values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVerticalUnits()
func (p NSPopover) AccessibilityVerticalUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](p.ID, objc.Sel("accessibilityVerticalUnits"))
	return NSAccessibilityUnits(rv)
}

// Returns the visible cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleCells()
func (p NSPopover) AccessibilityVisibleCells() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityVisibleCells"))
	return foundation.NSArrayFromID(rv)
}

// Returns the range of visible characters in the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleCharacterRange()
func (p NSPopover) AccessibilityVisibleCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](p.ID, objc.Sel("accessibilityVisibleCharacterRange"))
	return foundation.NSRange(rv)
}

// Returns the accessibility element’s visible child accessibility elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleChildren()
func (p NSPopover) AccessibilityVisibleChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityVisibleChildren"))
	return foundation.NSArrayFromID(rv)
}

// Returns the visible columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleColumns()
func (p NSPopover) AccessibilityVisibleColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityVisibleColumns"))
	return foundation.NSArrayFromID(rv)
}

// Returns the visible rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityVisibleRows()
func (p NSPopover) AccessibilityVisibleRows() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityVisibleRows"))
	return foundation.NSArrayFromID(rv)
}

// Returns the warning value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityWarningValue()
func (p NSPopover) AccessibilityWarningValue() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityWarningValue"))
	return objectivec.Object{ID: rv}
}

// Returns the window that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityWindow()
func (p NSPopover) AccessibilityWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityWindow"))
	return objectivec.Object{ID: rv}
}

// Returns an array that contains all the app’s windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityWindows()
func (p NSPopover) AccessibilityWindows() foundation.INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityWindows"))
	return foundation.NSArrayFromID(rv)
}

// Returns the child accessibility element that represents the window’s zoom
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityZoomButton()
func (p NSPopover) AccessibilityZoomButton() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("accessibilityZoomButton"))
	return objectivec.Object{ID: rv}
}

// Returns the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityAlternateUIVisible()
func (p NSPopover) IsAccessibilityAlternateUIVisible() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
	return rv
}

// Returns a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityDisclosed()
func (p NSPopover) IsAccessibilityDisclosed() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAccessibilityDisclosed"))
	return rv
}

// Returns a Boolean value that indicates whether the accessibility element is
// in an edited state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEdited()
func (p NSPopover) IsAccessibilityEdited() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAccessibilityEdited"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityElement()
func (p NSPopover) IsAccessibilityElement() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAccessibilityElement"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEnabled()
func (p NSPopover) IsAccessibilityEnabled() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAccessibilityEnabled"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// is in an expanded state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityExpanded()
func (p NSPopover) IsAccessibilityExpanded() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAccessibilityExpanded"))
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
func (p NSPopover) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAccessibilityFocused"))
	return rv
}

// Returns a Boolean value that determines whether the app is the frontmost
// app.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityFrontmost()
func (p NSPopover) IsAccessibilityFrontmost() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAccessibilityFrontmost"))
	return rv
}

// Returns a Boolean value that determines whether the app is in a hidden
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityHidden()
func (p NSPopover) IsAccessibilityHidden() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAccessibilityHidden"))
	return rv
}

// Returns a Boolean value that determines whether the window is the app’s
// main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMain()
func (p NSPopover) IsAccessibilityMain() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAccessibilityMain"))
	return rv
}

// Returns the Boolean value that determines whether the window is in a
// minimized state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMinimized()
func (p NSPopover) IsAccessibilityMinimized() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAccessibilityMinimized"))
	return rv
}

// Returns a Boolean value that determines whether the window is modal.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityModal()
func (p NSPopover) IsAccessibilityModal() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAccessibilityModal"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility
// element’s grid is in row major order or in column major order.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityOrderedByRow()
func (p NSPopover) IsAccessibilityOrderedByRow() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAccessibilityOrderedByRow"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityProtectedContent()
func (p NSPopover) IsAccessibilityProtectedContent() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAccessibilityProtectedContent"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// must have content for successful submission of a form.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityRequired()
func (p NSPopover) IsAccessibilityRequired() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAccessibilityRequired"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// is currently in a selected state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelected()
func (p NSPopover) IsAccessibilitySelected() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAccessibilitySelected"))
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
func (p NSPopover) IsAccessibilitySelectorAllowed(selector objc.SEL) bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isAccessibilitySelectorAllowed:"), selector)
	return rv
}

// Sets the activation point for the user interface element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityActivationPoint(_:)
func (p NSPopover) SetAccessibilityActivationPoint(accessibilityActivationPoint corefoundation.CGPoint) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityActivationPoint:"), accessibilityActivationPoint)
}

// Sets the allowed values for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityAllowedValues(_:)
func (p NSPopover) SetAccessibilityAllowedValues(accessibilityAllowedValues []foundation.NSNumber) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityAllowedValues:"), objectivec.IObjectSliceToNSArray(accessibilityAllowedValues))
}

// Sets the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityAlternateUIVisible(_:)
func (p NSPopover) SetAccessibilityAlternateUIVisible(accessibilityAlternateUIVisible bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityAlternateUIVisible:"), accessibilityAlternateUIVisible)
}

// Sets the child accessibility element with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityApplicationFocusedUIElement(_:)
func (p NSPopover) SetAccessibilityApplicationFocusedUIElement(accessibilityApplicationFocusedUIElement objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityApplicationFocusedUIElement:"), accessibilityApplicationFocusedUIElement)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityAttributedUserInputLabels(_:)
func (p NSPopover) SetAccessibilityAttributedUserInputLabels(accessibilityAttributedUserInputLabels []foundation.NSAttributedString) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityAttributedUserInputLabels:"), objectivec.IObjectSliceToNSArray(accessibilityAttributedUserInputLabels))
}

// Sets the child accessibility element that represents the window’s cancel
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCancelButton(_:)
func (p NSPopover) SetAccessibilityCancelButton(accessibilityCancelButton objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityCancelButton:"), accessibilityCancelButton)
}

// Sets the child accessibility elements in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityChildren(_:)
func (p NSPopover) SetAccessibilityChildren(accessibilityChildren foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityChildren:"), accessibilityChildren)
}

// Sets the array of child accessibility elements in order for linear
// navigation.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityChildrenInNavigationOrder(_:)
func (p NSPopover) SetAccessibilityChildrenInNavigationOrder(accessibilityChildrenInNavigationOrder []objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityChildrenInNavigationOrder:"), objectivec.IObjectSliceToNSArray(accessibilityChildrenInNavigationOrder))
}

// Sets the clear button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityClearButton(_:)
func (p NSPopover) SetAccessibilityClearButton(accessibilityClearButton objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityClearButton:"), accessibilityClearButton)
}

// Sets the child accessibility element that represents the window’s close
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCloseButton(_:)
func (p NSPopover) SetAccessibilityCloseButton(accessibilityCloseButton objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityCloseButton:"), accessibilityCloseButton)
}

// Sets the number of columns in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnCount(_:)
func (p NSPopover) SetAccessibilityColumnCount(accessibilityColumnCount int) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityColumnCount:"), accessibilityColumnCount)
}

// Sets the column header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnHeaderUIElements(_:)
func (p NSPopover) SetAccessibilityColumnHeaderUIElements(accessibilityColumnHeaderUIElements foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityColumnHeaderUIElements:"), accessibilityColumnHeaderUIElements)
}

// Sets the column index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnIndexRange(_:)
func (p NSPopover) SetAccessibilityColumnIndexRange(accessibilityColumnIndexRange foundation.NSRange) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityColumnIndexRange:"), accessibilityColumnIndexRange)
}

// Sets the column titles for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumnTitles(_:)
func (p NSPopover) SetAccessibilityColumnTitles(accessibilityColumnTitles foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityColumnTitles:"), accessibilityColumnTitles)
}

// Sets the column accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityColumns(_:)
func (p NSPopover) SetAccessibilityColumns(accessibilityColumns foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityColumns:"), accessibilityColumns)
}

// Sets the contents of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityContents(_:)
func (p NSPopover) SetAccessibilityContents(accessibilityContents foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityContents:"), accessibilityContents)
}

// Sets the critical value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCriticalValue(_:)
func (p NSPopover) SetAccessibilityCriticalValue(accessibilityCriticalValue objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityCriticalValue:"), accessibilityCriticalValue)
}

// Sets the custom actions of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCustomActions(_:)
func (p NSPopover) SetAccessibilityCustomActions(accessibilityCustomActions []NSAccessibilityCustomAction) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityCustomActions:"), objectivec.IObjectSliceToNSArray(accessibilityCustomActions))
}

// Sets the custom rotors of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityCustomRotors(_:)
func (p NSPopover) SetAccessibilityCustomRotors(accessibilityCustomRotors []NSAccessibilityCustomRotor) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityCustomRotors:"), objectivec.IObjectSliceToNSArray(accessibilityCustomRotors))
}

// Sets the decrement button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDecrementButton(_:)
func (p NSPopover) SetAccessibilityDecrementButton(accessibilityDecrementButton objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityDecrementButton:"), accessibilityDecrementButton)
}

// Sets the child accessibility element that represents the window’s default
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDefaultButton(_:)
func (p NSPopover) SetAccessibilityDefaultButton(accessibilityDefaultButton objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityDefaultButton:"), accessibilityDefaultButton)
}

// Sets a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosed(_:)
func (p NSPopover) SetAccessibilityDisclosed(accessibilityDisclosed bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityDisclosed:"), accessibilityDisclosed)
}

// Sets the row disclosing the current row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosedByRow(_:)
func (p NSPopover) SetAccessibilityDisclosedByRow(accessibilityDisclosedByRow objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityDisclosedByRow:"), accessibilityDisclosedByRow)
}

// Sets the rows that the current row discloses.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosedRows(_:)
func (p NSPopover) SetAccessibilityDisclosedRows(accessibilityDisclosedRows objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityDisclosedRows:"), accessibilityDisclosedRows)
}

// Sets the indention level for the row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDisclosureLevel(_:)
func (p NSPopover) SetAccessibilityDisclosureLevel(accessibilityDisclosureLevel int) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityDisclosureLevel:"), accessibilityDisclosureLevel)
}

// Sets the URL for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityDocument(_:)
func (p NSPopover) SetAccessibilityDocument(accessibilityDocument string) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityDocument:"), objc.String(accessibilityDocument))
}

// Sets a Boolean value that indicates whether the accessibility element is in
// an edited state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityEdited(_:)
func (p NSPopover) SetAccessibilityEdited(accessibilityEdited bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityEdited:"), accessibilityEdited)
}

// Sets a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityElement(_:)
func (p NSPopover) SetAccessibilityElement(accessibilityElement bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityElement:"), accessibilityElement)
}

// Sets a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityEnabled(_:)
func (p NSPopover) SetAccessibilityEnabled(accessibilityEnabled bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityEnabled:"), accessibilityEnabled)
}

// Sets a Boolean value that determines whether accessibility element is in an
// expanded state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityExpanded(_:)
func (p NSPopover) SetAccessibilityExpanded(accessibilityExpanded bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityExpanded:"), accessibilityExpanded)
}

// Sets the icon for the app’s menu bar extra.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityExtrasMenuBar(_:)
func (p NSPopover) SetAccessibilityExtrasMenuBar(accessibilityExtrasMenuBar objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityExtrasMenuBar:"), accessibilityExtrasMenuBar)
}

// Sets the filename for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFilename(_:)
func (p NSPopover) SetAccessibilityFilename(accessibilityFilename string) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityFilename:"), objc.String(accessibilityFilename))
}

// Sets a Boolean value that determines whether the accessibility element has
// the keyboard focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFocused(_:)
func (p NSPopover) SetAccessibilityFocused(accessibilityFocused bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityFocused:"), accessibilityFocused)
}

// Sets the child window with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFocusedWindow(_:)
func (p NSPopover) SetAccessibilityFocusedWindow(accessibilityFocusedWindow objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityFocusedWindow:"), accessibilityFocusedWindow)
}

// Sets the accessibility element’s frame in screen coordinates.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFrame(_:)
func (p NSPopover) SetAccessibilityFrame(accessibilityFrame corefoundation.CGRect) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityFrame:"), accessibilityFrame)
}

// Sets a Boolean value that determines whether the app is the frontmost app.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFrontmost(_:)
func (p NSPopover) SetAccessibilityFrontmost(accessibilityFrontmost bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityFrontmost:"), accessibilityFrontmost)
}

// Sets the child accessibility element that represents the window’s
// full-screen button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityFullScreenButton(_:)
func (p NSPopover) SetAccessibilityFullScreenButton(accessibilityFullScreenButton objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityFullScreenButton:"), accessibilityFullScreenButton)
}

// Sets the child accessibility element that represents the window’s grow
// area.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityGrowArea(_:)
func (p NSPopover) SetAccessibilityGrowArea(accessibilityGrowArea objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityGrowArea:"), accessibilityGrowArea)
}

// Sets the drag handle accessibility elements for the layout item element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHandles(_:)
func (p NSPopover) SetAccessibilityHandles(accessibilityHandles foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityHandles:"), accessibilityHandles)
}

// Sets the header for the table view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHeader(_:)
func (p NSPopover) SetAccessibilityHeader(accessibilityHeader objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityHeader:"), accessibilityHeader)
}

// Sets the help text for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHelp(_:)
func (p NSPopover) SetAccessibilityHelp(accessibilityHelp string) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityHelp:"), objc.String(accessibilityHelp))
}

// Sets a Boolean value that determines whether the app is in a hidden state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHidden(_:)
func (p NSPopover) SetAccessibilityHidden(accessibilityHidden bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityHidden:"), accessibilityHidden)
}

// Sets the horizontal scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHorizontalScrollBar(_:)
func (p NSPopover) SetAccessibilityHorizontalScrollBar(accessibilityHorizontalScrollBar objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityHorizontalScrollBar:"), accessibilityHorizontalScrollBar)
}

// Sets the description of the layout area’s horizontal units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHorizontalUnitDescription(_:)
func (p NSPopover) SetAccessibilityHorizontalUnitDescription(accessibilityHorizontalUnitDescription string) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityHorizontalUnitDescription:"), objc.String(accessibilityHorizontalUnitDescription))
}

// Sets the units that the layout area uses for horizontal values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityHorizontalUnits(_:)
func (p NSPopover) SetAccessibilityHorizontalUnits(accessibilityHorizontalUnits NSAccessibilityUnits) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityHorizontalUnits:"), accessibilityHorizontalUnits)
}

// Sets the accessibility element’s identity.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityIdentifier(_:)
func (p NSPopover) SetAccessibilityIdentifier(accessibilityIdentifier string) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityIdentifier:"), objc.String(accessibilityIdentifier))
}

// Sets the increment button for the stepper accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityIncrementButton(_:)
func (p NSPopover) SetAccessibilityIncrementButton(accessibilityIncrementButton objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityIncrementButton:"), accessibilityIncrementButton)
}

// Sets the index of the row or column that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityIndex(_:)
func (p NSPopover) SetAccessibilityIndex(accessibilityIndex int) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityIndex:"), accessibilityIndex)
}

// Sets the line number that contains the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityInsertionPointLineNumber(_:)
func (p NSPopover) SetAccessibilityInsertionPointLineNumber(accessibilityInsertionPointLineNumber int) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityInsertionPointLineNumber:"), accessibilityInsertionPointLineNumber)
}

// Sets a short description of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLabel(_:)
func (p NSPopover) SetAccessibilityLabel(accessibilityLabel string) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityLabel:"), objc.String(accessibilityLabel))
}

// Sets the child label elements for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLabelUIElements(_:)
func (p NSPopover) SetAccessibilityLabelUIElements(accessibilityLabelUIElements foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityLabelUIElements:"), accessibilityLabelUIElements)
}

// Sets the value of the label accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLabelValue(_:)
func (p NSPopover) SetAccessibilityLabelValue(accessibilityLabelValue float32) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityLabelValue:"), accessibilityLabelValue)
}

// Sets the elements that have links with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityLinkedUIElements(_:)
func (p NSPopover) SetAccessibilityLinkedUIElements(accessibilityLinkedUIElements foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityLinkedUIElements:"), accessibilityLinkedUIElements)
}

// Sets a Boolean value that determines whether the window is the app’s main
// window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMain(_:)
func (p NSPopover) SetAccessibilityMain(accessibilityMain bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityMain:"), accessibilityMain)
}

// Sets the app’s main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMainWindow(_:)
func (p NSPopover) SetAccessibilityMainWindow(accessibilityMainWindow objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityMainWindow:"), accessibilityMainWindow)
}

// Sets the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerGroupUIElement(_:)
func (p NSPopover) SetAccessibilityMarkerGroupUIElement(accessibilityMarkerGroupUIElement objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityMarkerGroupUIElement:"), accessibilityMarkerGroupUIElement)
}

// Sets the human-readable description of the marker type.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerTypeDescription(_:)
func (p NSPopover) SetAccessibilityMarkerTypeDescription(accessibilityMarkerTypeDescription string) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityMarkerTypeDescription:"), objc.String(accessibilityMarkerTypeDescription))
}

// Sets the array of marker accessibility elements for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerUIElements(_:)
func (p NSPopover) SetAccessibilityMarkerUIElements(accessibilityMarkerUIElements foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityMarkerUIElements:"), accessibilityMarkerUIElements)
}

// Sets the marker values for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMarkerValues(_:)
func (p NSPopover) SetAccessibilityMarkerValues(accessibilityMarkerValues objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityMarkerValues:"), accessibilityMarkerValues)
}

// Sets the maximum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMaxValue(_:)
func (p NSPopover) SetAccessibilityMaxValue(accessibilityMaxValue objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityMaxValue:"), accessibilityMaxValue)
}

// Sets the app’s menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMenuBar(_:)
func (p NSPopover) SetAccessibilityMenuBar(accessibilityMenuBar objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityMenuBar:"), accessibilityMenuBar)
}

// Sets the minimum value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMinValue(_:)
func (p NSPopover) SetAccessibilityMinValue(accessibilityMinValue objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityMinValue:"), accessibilityMinValue)
}

// Sets the child accessibility element that represents the window’s
// minimize button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMinimizeButton(_:)
func (p NSPopover) SetAccessibilityMinimizeButton(accessibilityMinimizeButton objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityMinimizeButton:"), accessibilityMinimizeButton)
}

// Sets the Boolean value that determines whether the window is in a minimized
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityMinimized(_:)
func (p NSPopover) SetAccessibilityMinimized(accessibilityMinimized bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityMinimized:"), accessibilityMinimized)
}

// Sets a Boolean value that determines whether the window is modal.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityModal(_:)
func (p NSPopover) SetAccessibilityModal(accessibilityModal bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityModal:"), accessibilityModal)
}

// Sets the contents that follow the divider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityNextContents(_:)
func (p NSPopover) SetAccessibilityNextContents(accessibilityNextContents foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityNextContents:"), accessibilityNextContents)
}

// Sets the number of characters in the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityNumberOfCharacters(_:)
func (p NSPopover) SetAccessibilityNumberOfCharacters(accessibilityNumberOfCharacters int) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityNumberOfCharacters:"), accessibilityNumberOfCharacters)
}

// Sets a Boolean value that determines whether the element’s grid is in row
// major order or in column major order.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityOrderedByRow(_:)
func (p NSPopover) SetAccessibilityOrderedByRow(accessibilityOrderedByRow bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityOrderedByRow:"), accessibilityOrderedByRow)
}

// Sets the orientation of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityOrientation(_:)
func (p NSPopover) SetAccessibilityOrientation(accessibilityOrientation NSAccessibilityOrientation) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityOrientation:"), accessibilityOrientation)
}

// Sets the overflow button for the toolbar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityOverflowButton(_:)
func (p NSPopover) SetAccessibilityOverflowButton(accessibilityOverflowButton objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityOverflowButton:"), accessibilityOverflowButton)
}

// Sets the accessibility element’s parent in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityParent(_:)
func (p NSPopover) SetAccessibilityParent(accessibilityParent objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityParent:"), accessibilityParent)
}

// Sets the placeholder value for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityPlaceholderValue(_:)
func (p NSPopover) SetAccessibilityPlaceholderValue(accessibilityPlaceholderValue string) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityPlaceholderValue:"), objc.String(accessibilityPlaceholderValue))
}

// Sets the contents that precede the divider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityPreviousContents(_:)
func (p NSPopover) SetAccessibilityPreviousContents(accessibilityPreviousContents foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityPreviousContents:"), accessibilityPreviousContents)
}

// Sets a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityProtectedContent(_:)
func (p NSPopover) SetAccessibilityProtectedContent(accessibilityProtectedContent bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityProtectedContent:"), accessibilityProtectedContent)
}

// Sets the child accessibility element that represents the window’s proxy
// icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityProxy(_:)
func (p NSPopover) SetAccessibilityProxy(accessibilityProxy objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityProxy:"), accessibilityProxy)
}

// Sets a Boolean value that determines whether the accessibility element must
// have content for successful submission of a form.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRequired(_:)
func (p NSPopover) SetAccessibilityRequired(accessibilityRequired bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityRequired:"), accessibilityRequired)
}

// Sets the type of interface element that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRole(_:)
func (p NSPopover) SetAccessibilityRole(accessibilityRole NSAccessibilityRole) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityRole:"), objc.String(string(accessibilityRole)))
}

// Sets the localized, human-intelligible description of the accessibility
// element’s role, such as radio button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRoleDescription(_:)
func (p NSPopover) SetAccessibilityRoleDescription(accessibilityRoleDescription string) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityRoleDescription:"), objc.String(accessibilityRoleDescription))
}

// Sets the number of rows in the accessibility element’s grid.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRowCount(_:)
func (p NSPopover) SetAccessibilityRowCount(accessibilityRowCount int) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityRowCount:"), accessibilityRowCount)
}

// Sets the row header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRowHeaderUIElements(_:)
func (p NSPopover) SetAccessibilityRowHeaderUIElements(accessibilityRowHeaderUIElements foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityRowHeaderUIElements:"), accessibilityRowHeaderUIElements)
}

// Sets the row index range of the cell.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRowIndexRange(_:)
func (p NSPopover) SetAccessibilityRowIndexRange(accessibilityRowIndexRange foundation.NSRange) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityRowIndexRange:"), accessibilityRowIndexRange)
}

// Sets the row accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRows(_:)
func (p NSPopover) SetAccessibilityRows(accessibilityRows foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityRows:"), accessibilityRows)
}

// Sets the type of markers for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityRulerMarkerType(_:)
func (p NSPopover) SetAccessibilityRulerMarkerType(accessibilityRulerMarkerType NSAccessibilityRulerMarkerType) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityRulerMarkerType:"), accessibilityRulerMarkerType)
}

// Sets the search button for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySearchButton(_:)
func (p NSPopover) SetAccessibilitySearchButton(accessibilitySearchButton objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilitySearchButton:"), accessibilitySearchButton)
}

// Sets the search menu for the search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySearchMenu(_:)
func (p NSPopover) SetAccessibilitySearchMenu(accessibilitySearchMenu objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilitySearchMenu:"), accessibilitySearchMenu)
}

// Sets a Boolean value that determines whether the accessibility element is
// currently in a selected state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelected(_:)
func (p NSPopover) SetAccessibilitySelected(accessibilitySelected bool) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilitySelected:"), accessibilitySelected)
}

// Sets the currently selected cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedCells(_:)
func (p NSPopover) SetAccessibilitySelectedCells(accessibilitySelectedCells foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilitySelectedCells:"), accessibilitySelectedCells)
}

// Sets the accessibility element’s currently selected children.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedChildren(_:)
func (p NSPopover) SetAccessibilitySelectedChildren(accessibilitySelectedChildren foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilitySelectedChildren:"), accessibilitySelectedChildren)
}

// Sets the currently selected columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedColumns(_:)
func (p NSPopover) SetAccessibilitySelectedColumns(accessibilitySelectedColumns foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilitySelectedColumns:"), accessibilitySelectedColumns)
}

// Sets the currently selected rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedRows(_:)
func (p NSPopover) SetAccessibilitySelectedRows(accessibilitySelectedRows foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilitySelectedRows:"), accessibilitySelectedRows)
}

// Sets the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedText(_:)
func (p NSPopover) SetAccessibilitySelectedText(accessibilitySelectedText string) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilitySelectedText:"), objc.String(accessibilitySelectedText))
}

// Sets the range of the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedTextRange(_:)
func (p NSPopover) SetAccessibilitySelectedTextRange(accessibilitySelectedTextRange foundation.NSRange) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilitySelectedTextRange:"), accessibilitySelectedTextRange)
}

// Sets an array of ranges for the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySelectedTextRanges(_:)
func (p NSPopover) SetAccessibilitySelectedTextRanges(accessibilitySelectedTextRanges []foundation.NSValue) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilitySelectedTextRanges:"), objectivec.IObjectSliceToNSArray(accessibilitySelectedTextRanges))
}

// Sets the list of elements that the accessibility element is a title for.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityServesAsTitleForUIElements(_:)
func (p NSPopover) SetAccessibilityServesAsTitleForUIElements(accessibilityServesAsTitleForUIElements foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityServesAsTitleForUIElements:"), accessibilityServesAsTitleForUIElements)
}

// Sets the range of characters that the accessibility element displays.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySharedCharacterRange(_:)
func (p NSPopover) SetAccessibilitySharedCharacterRange(accessibilitySharedCharacterRange foundation.NSRange) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilitySharedCharacterRange:"), accessibilitySharedCharacterRange)
}

// Sets the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySharedFocusElements(_:)
func (p NSPopover) SetAccessibilitySharedFocusElements(accessibilitySharedFocusElements foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilitySharedFocusElements:"), accessibilitySharedFocusElements)
}

// Sets the other elements that share text with the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySharedTextUIElements(_:)
func (p NSPopover) SetAccessibilitySharedTextUIElements(accessibilitySharedTextUIElements foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilitySharedTextUIElements:"), accessibilitySharedTextUIElements)
}

// Sets the menu currently displaying for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityShownMenu(_:)
func (p NSPopover) SetAccessibilityShownMenu(accessibilityShownMenu objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityShownMenu:"), accessibilityShownMenu)
}

// Sets the accessibility element’s sort direction.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySortDirection(_:)
func (p NSPopover) SetAccessibilitySortDirection(accessibilitySortDirection NSAccessibilitySortDirection) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilitySortDirection:"), accessibilitySortDirection)
}

// Sets the array that contains the views and splitter bar from the split
// view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySplitters(_:)
func (p NSPopover) SetAccessibilitySplitters(accessibilitySplitters foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilitySplitters:"), accessibilitySplitters)
}

// Sets the specialized interface element type that the accessibility element
// represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilitySubrole(_:)
func (p NSPopover) SetAccessibilitySubrole(accessibilitySubrole NSAccessibilitySubrole) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilitySubrole:"), objc.String(string(accessibilitySubrole)))
}

// Sets the tab accessibility elements for the tab view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTabs(_:)
func (p NSPopover) SetAccessibilityTabs(accessibilityTabs foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityTabs:"), accessibilityTabs)
}

// Sets the title of the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTitle(_:)
func (p NSPopover) SetAccessibilityTitle(accessibilityTitle string) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityTitle:"), objc.String(accessibilityTitle))
}

// Sets the static text element that represents the accessibility element’s
// title.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTitleUIElement(_:)
func (p NSPopover) SetAccessibilityTitleUIElement(accessibilityTitleUIElement objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityTitleUIElement:"), accessibilityTitleUIElement)
}

// Sets the child accessibility element that represents the window’s toolbar
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityToolbarButton(_:)
func (p NSPopover) SetAccessibilityToolbarButton(accessibilityToolbarButton objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityToolbarButton:"), accessibilityToolbarButton)
}

// Sets the top-level element that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityTopLevelUIElement(_:)
func (p NSPopover) SetAccessibilityTopLevelUIElement(accessibilityTopLevelUIElement objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityTopLevelUIElement:"), accessibilityTopLevelUIElement)
}

// Sets the URL for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityURL(_:)
func (p NSPopover) SetAccessibilityURL(accessibilityURL foundation.NSURL) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityURL:"), accessibilityURL)
}

// Sets the human-readable description of the ruler’s units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityUnitDescription(_:)
func (p NSPopover) SetAccessibilityUnitDescription(accessibilityUnitDescription string) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityUnitDescription:"), objc.String(accessibilityUnitDescription))
}

// Sets the units used for the ruler.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityUnits(_:)
func (p NSPopover) SetAccessibilityUnits(accessibilityUnits NSAccessibilityUnits) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityUnits:"), accessibilityUnits)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityUserInputLabels(_:)
func (p NSPopover) SetAccessibilityUserInputLabels(accessibilityUserInputLabels []string) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityUserInputLabels:"), objectivec.StringSliceToNSArray(accessibilityUserInputLabels))
}

// Sets the accessibility element’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityValue(_:)
func (p NSPopover) SetAccessibilityValue(accessibilityValue objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityValue:"), accessibilityValue)
}

// Sets the human-readable description of the accessibility element’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityValueDescription(_:)
func (p NSPopover) SetAccessibilityValueDescription(accessibilityValueDescription string) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityValueDescription:"), objc.String(accessibilityValueDescription))
}

// Sets the vertical scroll bar for the scroll view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVerticalScrollBar(_:)
func (p NSPopover) SetAccessibilityVerticalScrollBar(accessibilityVerticalScrollBar objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityVerticalScrollBar:"), accessibilityVerticalScrollBar)
}

// Sets the description of the layout area’s vertical units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVerticalUnitDescription(_:)
func (p NSPopover) SetAccessibilityVerticalUnitDescription(accessibilityVerticalUnitDescription string) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityVerticalUnitDescription:"), objc.String(accessibilityVerticalUnitDescription))
}

// Sets the units that the layout area uses for vertical values.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVerticalUnits(_:)
func (p NSPopover) SetAccessibilityVerticalUnits(accessibilityVerticalUnits NSAccessibilityUnits) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityVerticalUnits:"), accessibilityVerticalUnits)
}

// Sets the visible cells for the table.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleCells(_:)
func (p NSPopover) SetAccessibilityVisibleCells(accessibilityVisibleCells foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityVisibleCells:"), accessibilityVisibleCells)
}

// Sets the range of visible characters in the document.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleCharacterRange(_:)
func (p NSPopover) SetAccessibilityVisibleCharacterRange(accessibilityVisibleCharacterRange foundation.NSRange) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityVisibleCharacterRange:"), accessibilityVisibleCharacterRange)
}

// Sets the accessibility element’s visible child accessibility elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleChildren(_:)
func (p NSPopover) SetAccessibilityVisibleChildren(accessibilityVisibleChildren foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityVisibleChildren:"), accessibilityVisibleChildren)
}

// Sets the visible columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleColumns(_:)
func (p NSPopover) SetAccessibilityVisibleColumns(accessibilityVisibleColumns foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityVisibleColumns:"), accessibilityVisibleColumns)
}

// Sets the visible rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityVisibleRows(_:)
func (p NSPopover) SetAccessibilityVisibleRows(accessibilityVisibleRows foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityVisibleRows:"), accessibilityVisibleRows)
}

// Sets the warning value for the level indicator.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityWarningValue(_:)
func (p NSPopover) SetAccessibilityWarningValue(accessibilityWarningValue objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityWarningValue:"), accessibilityWarningValue)
}

// Sets the window that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityWindow(_:)
func (p NSPopover) SetAccessibilityWindow(accessibilityWindow objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityWindow:"), accessibilityWindow)
}

// Sets the array that contains all the app’s windows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityWindows(_:)
func (p NSPopover) SetAccessibilityWindows(accessibilityWindows foundation.INSArray) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityWindows:"), accessibilityWindows)
}

// Sets the child accessibility element that represents the window’s zoom
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/setAccessibilityZoomButton(_:)
func (p NSPopover) SetAccessibilityZoomButton(accessibilityZoomButton objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("setAccessibilityZoomButton:"), accessibilityZoomButton)
}

// The view controller that manages the content of the popover.
//
// # Discussion
//
// You must set the content view controller of the popover before the popover
// is shown. Changes to the popover’s content view controller while the
// popover is shown will cause the popover to animate if the
// [NSPopover.Animates] property is true.
//
// The default value is `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/contentViewController
func (p NSPopover) ContentViewController() INSViewController {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("contentViewController"))
	return NSViewControllerFromID(objc.ID(rv))
}
func (p NSPopover) SetContentViewController(value INSViewController) {
	objc.Send[struct{}](p.ID, objc.Sel("setContentViewController:"), value)
}

// Specifies the behavior of the popover.
//
// # Discussion
//
// The default value is [NSPopoverBehaviorApplicationDefined]. See
// [NSPopover.Behavior] for possible value.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/behavior-swift.property
//
// [NSPopover.Behavior]: https://developer.apple.com/documentation/AppKit/NSPopover/Behavior-swift.enum
func (p NSPopover) Behavior() NSPopoverBehavior {
	rv := objc.Send[NSPopoverBehavior](p.ID, objc.Sel("behavior"))
	return NSPopoverBehavior(rv)
}
func (p NSPopover) SetBehavior(value NSPopoverBehavior) {
	objc.Send[struct{}](p.ID, objc.Sel("setBehavior:"), value)
}

// The rectangle within the positioning view relative to which the popover
// should be positioned.
//
// # Discussion
//
// Popovers are positioned relative to a positioning view and are
// automatically moved when the location or size of the positioning view
// changes.
//
// Sometimes it is desirable to position popovers relative to a rectangle
// within the positioning view. In this case, you must update the
// `positioningRect` property whenever this rectangle changes.
//
// This property is exposed as a read-only binding.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/positioningRect
func (p NSPopover) PositioningRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](p.ID, objc.Sel("positioningRect"))
	return corefoundation.CGRect(rv)
}
func (p NSPopover) SetPositioningRect(value corefoundation.CGRect) {
	objc.Send[struct{}](p.ID, objc.Sel("setPositioningRect:"), value)
}

// The appearance of the popover.
//
// # Discussion
//
// If no appearance is specified, the popover’s effective appearance
// defaults to [vibrantLight].
//
// In apps that run in macOS 10.10 and later, the previous property type of
// [NSPopover.Appearance] is deprecated. In apps that run in OS X v10.9 and
// earlier, the [aqua] appearance is automatically set on popover content.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/appearance-swift.property
//
// [NSPopover.Appearance]: https://developer.apple.com/documentation/AppKit/NSPopover/Appearance-swift.enum
// [aqua]: https://developer.apple.com/documentation/AppKit/NSAppearance/Name-swift.struct/aqua
// [vibrantLight]: https://developer.apple.com/documentation/AppKit/NSAppearance/Name-swift.struct/vibrantLight
func (p NSPopover) Appearance() INSAppearance {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("appearance"))
	return NSAppearanceFromID(objc.ID(rv))
}
func (p NSPopover) SetAppearance(value INSAppearance) {
	objc.Send[struct{}](p.ID, objc.Sel("setAppearance:"), value)
}

// The appearance that will be used when the popover is displayed onscreen.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/effectiveAppearance
func (p NSPopover) EffectiveAppearance() INSAppearance {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("effectiveAppearance"))
	return NSAppearanceFromID(objc.ID(rv))
}

// Specifies if the popover is to be animated.
//
// # Discussion
//
// A popover may be animated when it shows, closes, moves, or appears to
// transition to a detachable window. This property also controls whether the
// popover animates when the content view or content size changes.
//
// The system does not guarantee which behaviors will be animated or that this
// property will be respected; it is regarded as a hint.
//
// The default value is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/animates
func (p NSPopover) Animates() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("animates"))
	return rv
}
func (p NSPopover) SetAnimates(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setAnimates:"), value)
}

// The content size of the popover.
//
// # Discussion
//
// The popover’s content size is set to match the size of the content view
// when the content view controller is set.
//
// Changes to the content size of the popover will cause the popover to
// animate while it is shown if the [NSPopover.Animates] property is true.
//
// This property is exposed as a read-only binding.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/contentSize
func (p NSPopover) ContentSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](p.ID, objc.Sel("contentSize"))
	return corefoundation.CGSize(rv)
}
func (p NSPopover) SetContentSize(value corefoundation.CGSize) {
	objc.Send[struct{}](p.ID, objc.Sel("setContentSize:"), value)
}

// The display state of the popover.
//
// # Discussion
//
// The value is true if the popover is being shown, false otherwise.
//
// The popover is considered to be shown from the point when
// [NSPopover.ShowRelativeToRectOfViewPreferredEdge] is invoked. A popover is
// considered closed in response to an invocation of either [NSPopover.Close]
// or [NSPopover.PerformClose].
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/isShown
func (p NSPopover) IsShown() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isShown"))
	return rv
}

// A Boolean value that indicates whether the window created by a popover’s
// detachment is automatically created.
//
// # Discussion
//
// When [NSPopover.Detached] is true, the detached window is automatically
// created. This property does not apply when detaching a popover results in a
// window returned by [DetachableWindowForPopover].
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/isDetached
func (p NSPopover) IsDetached() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("isDetached"))
	return rv
}

// The delegate of the popover.
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/delegate
func (p NSPopover) Delegate() NSPopoverDelegate {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("delegate"))
	return NSPopoverDelegateObjectFromID(rv)
}
func (p NSPopover) SetDelegate(value NSPopoverDelegate) {
	objc.Send[struct{}](p.ID, objc.Sel("setDelegate:"), value)
}

// A Boolean value that indicates whether the content view of the popover
// extends into the arrow region.
//
// # Discussion
//
// Setting the value of this property to true extends the frame of the content
// view by the height of the arrow region on all four sides of the frame. This
// causes the [NSPopover.ContentViewController] view to extend to the
// window’s bounds.
//
// [media-4304810]
//
// Use the [NSView.SafeAreaLayoutGuide] of the
// [NSPopover.ContentViewController] view to ensure that your content is fully
// visible and doesn’t become clipped when displayed.
//
// [media-4304811]
//
// Setting this value to false doesn’t extend the
// [NSPopover.ContentViewController] view fully into the arrow region. The
// default value for this property is false.
//
// [media-4304812]
//
// See: https://developer.apple.com/documentation/AppKit/NSPopover/hasFullSizeContent
func (p NSPopover) HasFullSizeContent() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("hasFullSizeContent"))
	return rv
}
func (p NSPopover) SetHasFullSizeContent(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setHasFullSizeContent:"), value)
}

// Protocol methods for NSAccessibilityElementProtocol

// Protocol methods for NSAccessibilityProtocol

// Protocol methods for NSAppearanceCustomization
