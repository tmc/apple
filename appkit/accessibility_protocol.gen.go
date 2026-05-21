// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The complete list of properties and methods for accessible elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol
type NSAccessibilityProtocol interface {
	objectivec.IObject

	// Returns a Boolean value that determines whether the accessibility element participates in the accessibility hierarchy.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityElement()
	IsAccessibilityElement() bool

	// Returns a Boolean value that determines whether the accessibility element responds to user events.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEnabled()
	IsAccessibilityEnabled() bool

	// Returns a Boolean value that indicates whether assistive apps can invoke the specified selector on the accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelectorAllowed(_:)
	IsAccessibilitySelectorAllowed(selector objc.SEL) bool

	// Returns a Boolean value that determines whether the accessibility element contains protected content.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityProtectedContent()
	IsAccessibilityProtectedContent() bool

	// Returns a Boolean value that determines whether the accessibility element is currently in a selected state.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelected()
	IsAccessibilitySelected() bool

	// Returns a Boolean value that determines whether the accessibility element has the keyboard focus.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityFocused()
	IsAccessibilityFocused() bool

	// Returns a Boolean value that determines whether the accessibility element must have content for successful submission of a form.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityRequired()
	IsAccessibilityRequired() bool

	// Returns the substring for the specified range.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityString(for:)
	AccessibilityStringForRange(range_ foundation.NSRange) string

	// Returns the attributed substring for the specified range of characters.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityAttributedString(for:)
	AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString

	// Returns the rich text format (RTF) data that describes the specified range of characters.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRTF(for:)
	AccessibilityRTFForRange(range_ foundation.NSRange) foundation.NSData

	// Returns the rectangle that encloses the specified range of characters.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityFrame(for:)
	AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect

	// Returns the line number for the line that contains the specified character index.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLine(for:)
	AccessibilityLineForIndex(index int) int

	// Returns the range of characters for the glyph that includes the specified character.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(for:)-6kv3
	AccessibilityRangeForIndex(index int) foundation.NSRange

	// Returns a range of characters that all have the same style as the specified character.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityStyleRange(for:)
	AccessibilityStyleRangeForIndex(index int) foundation.NSRange

	// Returns the range of characters in the specified line.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(forLine:)
	AccessibilityRangeForLine(line int) foundation.NSRange

	// Returns the range of characters for the glyph at the specified point.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityRange(for:)-1iudm
	AccessibilityRangeForPosition(point corefoundation.CGPoint) foundation.NSRange

	// Returns the Boolean value that determines whether the accessibility element’s alternative UI is currently visible.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityAlternateUIVisible()
	IsAccessibilityAlternateUIVisible() bool

	// Returns a Boolean value that determines whether the window is the app’s main window.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMain()
	IsAccessibilityMain() bool

	// Returns the Boolean value that determines whether the window is in a minimized state.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMinimized()
	IsAccessibilityMinimized() bool

	// Returns a Boolean value that determines whether the window is modal.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityModal()
	IsAccessibilityModal() bool

	// Returns a Boolean value that determines whether the app is the frontmost app.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityFrontmost()
	IsAccessibilityFrontmost() bool

	// Returns a Boolean value that determines whether the app is in a hidden state.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityHidden()
	IsAccessibilityHidden() bool

	// Returns a Boolean value that determines whether the accessibility element’s grid is in row major order or in column major order.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityOrderedByRow()
	IsAccessibilityOrderedByRow() bool

	// Returns a Boolean value that determines whether the accessibility element is in an expanded state.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityExpanded()
	IsAccessibilityExpanded() bool

	// Returns a Boolean value that determines whether the row is disclosing other rows.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityDisclosed()
	IsAccessibilityDisclosed() bool

	// Converts the provided point in screen coordinates to a point in the layout area’s coordinate system.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLayoutPoint(forScreenPoint:)
	AccessibilityLayoutPointForScreenPoint(point corefoundation.CGPoint) corefoundation.CGPoint

	// Converts the provided size in screen coordinates to a size in the layout area’s coordinate system.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityLayoutSize(forScreenSize:)
	AccessibilityLayoutSizeForScreenSize(size corefoundation.CGSize) corefoundation.CGSize

	// Converts the provided point in the layout area’s coordinates to a point in the screen’s coordinate system.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityScreenPoint(forLayoutPoint:)
	AccessibilityScreenPointForLayoutPoint(point corefoundation.CGPoint) corefoundation.CGPoint

	// Converts the provided size in the layout area’s coordinates to a size in the screen’s coordinate system.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityScreenSize(forLayoutSize:)
	AccessibilityScreenSizeForLayoutSize(size corefoundation.CGSize) corefoundation.CGSize

	// Returns a Boolean value that indicates whether the accessibility element is in an edited state.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEdited()
	IsAccessibilityEdited() bool

	// Cancels the current operation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformCancel()
	AccessibilityPerformCancel() bool

	// Simulates pressing Return in the accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformConfirm()
	AccessibilityPerformConfirm() bool

	// Selects the accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformPick()
	AccessibilityPerformPick() bool

	// Simulates clicking the accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformPress()
	AccessibilityPerformPress() bool

	// Displays the accessibility element’s alternative UI.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowAlternateUI()
	AccessibilityPerformShowAlternateUI() bool

	// Returns to the accessibility element’s original UI.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowDefaultUI()
	AccessibilityPerformShowDefaultUI() bool

	// Displays the menu accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformShowMenu()
	AccessibilityPerformShowMenu() bool

	// Brings the window to the front.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformRaise()
	AccessibilityPerformRaise() bool

	// Increments the accessibility element’s value.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformIncrement()
	AccessibilityPerformIncrement() bool

	// Decrements the accessibility element’s value.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDecrement()
	AccessibilityPerformDecrement() bool

	// Deletes the accessibility element’s value.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDelete()
	AccessibilityPerformDelete() bool

	// The activation point for the user interface element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityActivationPoint
	AccessibilityActivationPoint() corefoundation.CGPoint
	SetAccessibilityActivationPoint(value corefoundation.CGPoint)

	// The allowed values for the slider accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityAllowedValues
	AccessibilityAllowedValues() []foundation.NSNumber
	SetAccessibilityAllowedValues(value []foundation.NSNumber)

	// A Boolean value that determines whether the accessibility element’s alternative UI is currently visible.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityAlternateUIVisible
	AccessibilityAlternateUIVisible() bool
	SetAccessibilityAlternateUIVisible(value bool)

	// The child accessibility element with the current focus.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityApplicationFocusedUIElement
	AccessibilityApplicationFocusedUIElement() objectivec.IObject
	SetAccessibilityApplicationFocusedUIElement(value objectivec.IObject)

	// accessibilityAttributedUserInputLabels protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityAttributedUserInputLabels
	AccessibilityAttributedUserInputLabels() []foundation.NSAttributedString
	SetAccessibilityAttributedUserInputLabels(value []foundation.NSAttributedString)

	// The child accessibility element that represents the window’s cancel button.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCancelButton
	AccessibilityCancelButton() objectivec.IObject
	SetAccessibilityCancelButton(value objectivec.IObject)

	// The child accessibility elements in the accessibility hierarchy.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityChildren
	AccessibilityChildren() foundation.INSArray
	SetAccessibilityChildren(value foundation.INSArray)

	// An array of child accessibility elements in order for linear navigation.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityChildrenInNavigationOrder
	AccessibilityChildrenInNavigationOrder() []objectivec.IObject
	SetAccessibilityChildrenInNavigationOrder(value []objectivec.IObject)

	// The clear button for the search field.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityClearButton
	AccessibilityClearButton() objectivec.IObject
	SetAccessibilityClearButton(value objectivec.IObject)

	// The child accessibility element that represents the window’s close button.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCloseButton
	AccessibilityCloseButton() objectivec.IObject
	SetAccessibilityCloseButton(value objectivec.IObject)

	// The number of columns in the accessibility element’s grid.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumnCount
	AccessibilityColumnCount() int
	SetAccessibilityColumnCount(value int)

	// The column header accessibility elements for the table or outline.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumnHeaderUIElements
	AccessibilityColumnHeaderUIElements() foundation.INSArray
	SetAccessibilityColumnHeaderUIElements(value foundation.INSArray)

	// The column index range of the cell.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumnIndexRange
	AccessibilityColumnIndexRange() foundation.NSRange
	SetAccessibilityColumnIndexRange(value foundation.NSRange)

	// The column titles for the accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumnTitles
	AccessibilityColumnTitles() foundation.INSArray
	SetAccessibilityColumnTitles(value foundation.INSArray)

	// The column accessibility elements for the table or outline.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumns
	AccessibilityColumns() foundation.INSArray
	SetAccessibilityColumns(value foundation.INSArray)

	// The contents of the current accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityContents
	AccessibilityContents() foundation.INSArray
	SetAccessibilityContents(value foundation.INSArray)

	// The critical value for the level indicator.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCriticalValue
	AccessibilityCriticalValue() objectivec.IObject
	SetAccessibilityCriticalValue(value objectivec.IObject)

	// The custom actions of the current accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCustomActions
	AccessibilityCustomActions() []NSAccessibilityCustomAction
	SetAccessibilityCustomActions(value []NSAccessibilityCustomAction)

	// The custom rotors of the current accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCustomRotors
	AccessibilityCustomRotors() []NSAccessibilityCustomRotor
	SetAccessibilityCustomRotors(value []NSAccessibilityCustomRotor)

	// The decrement button for the stepper accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDecrementButton
	AccessibilityDecrementButton() objectivec.IObject
	SetAccessibilityDecrementButton(value objectivec.IObject)

	// The child accessibility element that represents the window’s default button.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDefaultButton
	AccessibilityDefaultButton() objectivec.IObject
	SetAccessibilityDefaultButton(value objectivec.IObject)

	// A Boolean value that determines whether the row is disclosing other rows.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosed
	AccessibilityDisclosed() bool
	SetAccessibilityDisclosed(value bool)

	// The row disclosing the current row.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosedByRow
	AccessibilityDisclosedByRow() objectivec.IObject
	SetAccessibilityDisclosedByRow(value objectivec.IObject)

	// The rows that the current row discloses.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosedRows
	AccessibilityDisclosedRows() objectivec.IObject
	SetAccessibilityDisclosedRows(value objectivec.IObject)

	// The indention level for the row.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosureLevel
	AccessibilityDisclosureLevel() int
	SetAccessibilityDisclosureLevel(value int)

	// The URL for the file that the accessibility element represents.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDocument
	AccessibilityDocument() string
	SetAccessibilityDocument(value string)

	// A Boolean value that indicates whether the accessibility element is in an edited state.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityEdited
	AccessibilityEdited() bool
	SetAccessibilityEdited(value bool)

	// A Boolean value that determines whether the accessibility element participates in the accessibility hierarchy.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityElement
	AccessibilityElement() bool
	SetAccessibilityElement(value bool)

	// A Boolean value that determines whether the accessibility element responds to user events.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityEnabled
	AccessibilityEnabled() bool
	SetAccessibilityEnabled(value bool)

	// A Boolean value that determines whether the accessibility element is in an expanded state.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityExpanded
	AccessibilityExpanded() bool
	SetAccessibilityExpanded(value bool)

	// The icon for the app’s menu bar extra.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityExtrasMenuBar
	AccessibilityExtrasMenuBar() objectivec.IObject
	SetAccessibilityExtrasMenuBar(value objectivec.IObject)

	// The filename for the file that the accessibility element represents.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFilename
	AccessibilityFilename() string
	SetAccessibilityFilename(value string)

	// A Boolean value that determines whether the accessibility element has the keyboard focus.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocused
	AccessibilityFocused() bool
	SetAccessibilityFocused(value bool)

	// The child window with the current focus.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocusedWindow
	AccessibilityFocusedWindow() objectivec.IObject
	SetAccessibilityFocusedWindow(value objectivec.IObject)

	// The accessibility element’s frame in screen coordinates.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
	AccessibilityFrame() corefoundation.CGRect
	SetAccessibilityFrame(value corefoundation.CGRect)

	// A Boolean value that determines whether the app is the frontmost app.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrontmost
	AccessibilityFrontmost() bool
	SetAccessibilityFrontmost(value bool)

	// The child accessibility element that represents the window’s full-screen button.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFullScreenButton
	AccessibilityFullScreenButton() objectivec.IObject
	SetAccessibilityFullScreenButton(value objectivec.IObject)

	// The child accessibility element that represents the window’s grow area.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityGrowArea
	AccessibilityGrowArea() objectivec.IObject
	SetAccessibilityGrowArea(value objectivec.IObject)

	// The drag handle accessibility elements for the layout item element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHandles
	AccessibilityHandles() foundation.INSArray
	SetAccessibilityHandles(value foundation.INSArray)

	// The header for the table view.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHeader
	AccessibilityHeader() objectivec.IObject
	SetAccessibilityHeader(value objectivec.IObject)

	// The help text for the accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHelp
	AccessibilityHelp() string
	SetAccessibilityHelp(value string)

	// A Boolean value that determines whether the app is in a hidden state.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHidden
	AccessibilityHidden() bool
	SetAccessibilityHidden(value bool)

	// The horizontal scroll bar for the scroll view.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHorizontalScrollBar
	AccessibilityHorizontalScrollBar() objectivec.IObject
	SetAccessibilityHorizontalScrollBar(value objectivec.IObject)

	// A description of the layout area’s horizontal units.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHorizontalUnitDescription
	AccessibilityHorizontalUnitDescription() string
	SetAccessibilityHorizontalUnitDescription(value string)

	// The units that the layout area uses for horizontal values.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHorizontalUnits
	AccessibilityHorizontalUnits() NSAccessibilityUnits
	SetAccessibilityHorizontalUnits(value NSAccessibilityUnits)

	// The accessibility element’s identity.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIdentifier
	AccessibilityIdentifier() string
	SetAccessibilityIdentifier(value string)

	// The increment button for the stepper accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIncrementButton
	AccessibilityIncrementButton() objectivec.IObject
	SetAccessibilityIncrementButton(value objectivec.IObject)

	// The index of the row or column that the accessibility element represents.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIndex
	AccessibilityIndex() int
	SetAccessibilityIndex(value int)

	// The line number that contains the insertion point.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityInsertionPointLineNumber
	AccessibilityInsertionPointLineNumber() int
	SetAccessibilityInsertionPointLineNumber(value int)

	// A short description of the accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLabel
	AccessibilityLabel() string
	SetAccessibilityLabel(value string)

	// The child label elements for the slider accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLabelUIElements
	AccessibilityLabelUIElements() foundation.INSArray
	SetAccessibilityLabelUIElements(value foundation.INSArray)

	// The value of the label accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLabelValue
	AccessibilityLabelValue() float32
	SetAccessibilityLabelValue(value float32)

	// The elements that have links with the accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLinkedUIElements
	AccessibilityLinkedUIElements() foundation.INSArray
	SetAccessibilityLinkedUIElements(value foundation.INSArray)

	// A Boolean value that determines whether the window is the app’s main window.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMain
	AccessibilityMain() bool
	SetAccessibilityMain(value bool)

	// The app’s main window.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMainWindow
	AccessibilityMainWindow() objectivec.IObject
	SetAccessibilityMainWindow(value objectivec.IObject)

	// The user interface element that functions as a marker group for the ruler accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMarkerGroupUIElement
	AccessibilityMarkerGroupUIElement() objectivec.IObject
	SetAccessibilityMarkerGroupUIElement(value objectivec.IObject)

	// A human-readable description of the marker type.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMarkerTypeDescription
	AccessibilityMarkerTypeDescription() string
	SetAccessibilityMarkerTypeDescription(value string)

	// An array of marker accessibility elements for the ruler.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMarkerUIElements
	AccessibilityMarkerUIElements() foundation.INSArray
	SetAccessibilityMarkerUIElements(value foundation.INSArray)

	// The marker values for the ruler.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMarkerValues
	AccessibilityMarkerValues() objectivec.IObject
	SetAccessibilityMarkerValues(value objectivec.IObject)

	// The maximum value for the accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMaxValue
	AccessibilityMaxValue() objectivec.IObject
	SetAccessibilityMaxValue(value objectivec.IObject)

	// The app’s menu bar.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMenuBar
	AccessibilityMenuBar() objectivec.IObject
	SetAccessibilityMenuBar(value objectivec.IObject)

	// The minimum value for the accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMinValue
	AccessibilityMinValue() objectivec.IObject
	SetAccessibilityMinValue(value objectivec.IObject)

	// The child accessibility element that represents the window’s minimize button.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMinimizeButton
	AccessibilityMinimizeButton() objectivec.IObject
	SetAccessibilityMinimizeButton(value objectivec.IObject)

	// A Boolean value that determines whether this window is in a minimized state.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMinimized
	AccessibilityMinimized() bool
	SetAccessibilityMinimized(value bool)

	// A Boolean value that determines whether the window is modal.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityModal
	AccessibilityModal() bool
	SetAccessibilityModal(value bool)

	// The contents that follow the divider accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityNextContents
	AccessibilityNextContents() foundation.INSArray
	SetAccessibilityNextContents(value foundation.INSArray)

	// The number of characters in the text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityNumberOfCharacters
	AccessibilityNumberOfCharacters() int
	SetAccessibilityNumberOfCharacters(value int)

	// A Boolean value that determines whether the accessibility element’s grid is in row major order or in column major order.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityOrderedByRow
	AccessibilityOrderedByRow() bool
	SetAccessibilityOrderedByRow(value bool)

	// The orientation of the accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityOrientation
	AccessibilityOrientation() NSAccessibilityOrientation
	SetAccessibilityOrientation(value NSAccessibilityOrientation)

	// The overflow button for the toolbar.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityOverflowButton
	AccessibilityOverflowButton() objectivec.IObject
	SetAccessibilityOverflowButton(value objectivec.IObject)

	// The accessibility element’s parent in the accessibility hierarchy.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityParent
	AccessibilityParent() objectivec.IObject
	SetAccessibilityParent(value objectivec.IObject)

	// The placeholder value for the accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityPlaceholderValue
	AccessibilityPlaceholderValue() string
	SetAccessibilityPlaceholderValue(value string)

	// The contents that precede the divider accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityPreviousContents
	AccessibilityPreviousContents() foundation.INSArray
	SetAccessibilityPreviousContents(value foundation.INSArray)

	// A Boolean value that determines whether the accessibility element contains protected content.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityProtectedContent
	AccessibilityProtectedContent() bool
	SetAccessibilityProtectedContent(value bool)

	// The child accessibility element that represents the window’s proxy icon.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityProxy
	AccessibilityProxy() objectivec.IObject
	SetAccessibilityProxy(value objectivec.IObject)

	// A Boolean value that determines whether the accessibility element must have content for successful submission of a form.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRequired
	AccessibilityRequired() bool
	SetAccessibilityRequired(value bool)

	// The type of interface element that the accessibility element represents.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRole
	AccessibilityRole() NSAccessibilityRole
	SetAccessibilityRole(value NSAccessibilityRole)

	// A localized, human-intelligible description of the accessibility element’s role, such as .
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRoleDescription
	AccessibilityRoleDescription() string
	SetAccessibilityRoleDescription(value string)

	// The number of rows in the accessibility element’s grid.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRowCount
	AccessibilityRowCount() int
	SetAccessibilityRowCount(value int)

	// The row header accessibility elements for the table or outline.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRowHeaderUIElements
	AccessibilityRowHeaderUIElements() foundation.INSArray
	SetAccessibilityRowHeaderUIElements(value foundation.INSArray)

	// The row index range of the cell.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRowIndexRange
	AccessibilityRowIndexRange() foundation.NSRange
	SetAccessibilityRowIndexRange(value foundation.NSRange)

	// The row accessibility elements for the table or outline.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRows
	AccessibilityRows() foundation.INSArray
	SetAccessibilityRows(value foundation.INSArray)

	// The type of markers for the ruler.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRulerMarkerType
	AccessibilityRulerMarkerType() NSAccessibilityRulerMarkerType
	SetAccessibilityRulerMarkerType(value NSAccessibilityRulerMarkerType)

	// The search button for the search field.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySearchButton
	AccessibilitySearchButton() objectivec.IObject
	SetAccessibilitySearchButton(value objectivec.IObject)

	// The search menu for the search field.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySearchMenu
	AccessibilitySearchMenu() objectivec.IObject
	SetAccessibilitySearchMenu(value objectivec.IObject)

	// A Boolean value that determines whether the accessibility element is currently in a selected state.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelected
	AccessibilitySelected() bool
	SetAccessibilitySelected(value bool)

	// The currently selected cells for the table.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedCells
	AccessibilitySelectedCells() foundation.INSArray
	SetAccessibilitySelectedCells(value foundation.INSArray)

	// The accessibility element’s currently selected children.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedChildren
	AccessibilitySelectedChildren() foundation.INSArray
	SetAccessibilitySelectedChildren(value foundation.INSArray)

	// The currently selected columns for the table or outline.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedColumns
	AccessibilitySelectedColumns() foundation.INSArray
	SetAccessibilitySelectedColumns(value foundation.INSArray)

	// The currently selected rows for the table or outline.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedRows
	AccessibilitySelectedRows() foundation.INSArray
	SetAccessibilitySelectedRows(value foundation.INSArray)

	// The currently selected text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedText
	AccessibilitySelectedText() string
	SetAccessibilitySelectedText(value string)

	// The range of the currently selected text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedTextRange
	AccessibilitySelectedTextRange() foundation.NSRange
	SetAccessibilitySelectedTextRange(value foundation.NSRange)

	// An array of ranges for the currently selected text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedTextRanges
	AccessibilitySelectedTextRanges() []foundation.NSValue
	SetAccessibilitySelectedTextRanges(value []foundation.NSValue)

	// The list of elements that the accessibility element is a title for.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityServesAsTitleForUIElements
	AccessibilityServesAsTitleForUIElements() foundation.INSArray
	SetAccessibilityServesAsTitleForUIElements(value foundation.INSArray)

	// The range of characters that the accessibility element displays.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySharedCharacterRange
	AccessibilitySharedCharacterRange() foundation.NSRange
	SetAccessibilitySharedCharacterRange(value foundation.NSRange)

	// An array of elements that shares the keyboard focus with the accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySharedFocusElements
	AccessibilitySharedFocusElements() foundation.INSArray
	SetAccessibilitySharedFocusElements(value foundation.INSArray)

	// Other elements that share text with the accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySharedTextUIElements
	AccessibilitySharedTextUIElements() foundation.INSArray
	SetAccessibilitySharedTextUIElements(value foundation.INSArray)

	// The menu currently displaying for the accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityShownMenu
	AccessibilityShownMenu() objectivec.IObject
	SetAccessibilityShownMenu(value objectivec.IObject)

	// The accessibility element’s sort direction.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySortDirection
	AccessibilitySortDirection() NSAccessibilitySortDirection
	SetAccessibilitySortDirection(value NSAccessibilitySortDirection)

	// An array that contains the views and splitter bar from the split view.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySplitters
	AccessibilitySplitters() foundation.INSArray
	SetAccessibilitySplitters(value foundation.INSArray)

	// The specialized interface element type that the accessibility element represents.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySubrole
	AccessibilitySubrole() NSAccessibilitySubrole
	SetAccessibilitySubrole(value NSAccessibilitySubrole)

	// The tab accessibility elements for the tab view.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityTabs
	AccessibilityTabs() foundation.INSArray
	SetAccessibilityTabs(value foundation.INSArray)

	// The title of the accessibility element—for example, a button’s visible text.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityTitle
	AccessibilityTitle() string
	SetAccessibilityTitle(value string)

	// A static text element that represents the accessibility element’s title.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityTitleUIElement
	AccessibilityTitleUIElement() objectivec.IObject
	SetAccessibilityTitleUIElement(value objectivec.IObject)

	// The child accessibility element that represents the window’s toolbar button.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityToolbarButton
	AccessibilityToolbarButton() objectivec.IObject
	SetAccessibilityToolbarButton(value objectivec.IObject)

	// The top-level element that contains the accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityTopLevelUIElement
	AccessibilityTopLevelUIElement() objectivec.IObject
	SetAccessibilityTopLevelUIElement(value objectivec.IObject)

	// The URL for the accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityURL
	AccessibilityURL() foundation.NSURL
	SetAccessibilityURL(value foundation.NSURL)

	// A human-readable description of the ruler’s units.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityUnitDescription
	AccessibilityUnitDescription() string
	SetAccessibilityUnitDescription(value string)

	// The units for the ruler.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityUnits
	AccessibilityUnits() NSAccessibilityUnits
	SetAccessibilityUnits(value NSAccessibilityUnits)

	// accessibilityUserInputLabels protocol.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityUserInputLabels
	AccessibilityUserInputLabels() []string
	SetAccessibilityUserInputLabels(value []string)

	// The accessibility element’s value.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
	AccessibilityValue() objectivec.IObject
	SetAccessibilityValue(value objectivec.IObject)

	// A human-readable description of the accessibility element’s value.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValueDescription
	AccessibilityValueDescription() string
	SetAccessibilityValueDescription(value string)

	// The vertical scroll bar for the scroll view.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVerticalScrollBar
	AccessibilityVerticalScrollBar() objectivec.IObject
	SetAccessibilityVerticalScrollBar(value objectivec.IObject)

	// A description of the layout area’s vertical units.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVerticalUnitDescription
	AccessibilityVerticalUnitDescription() string
	SetAccessibilityVerticalUnitDescription(value string)

	// The units that the layout area uses for vertical values.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVerticalUnits
	AccessibilityVerticalUnits() NSAccessibilityUnits
	SetAccessibilityVerticalUnits(value NSAccessibilityUnits)

	// The visible cells for the table.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleCells
	AccessibilityVisibleCells() foundation.INSArray
	SetAccessibilityVisibleCells(value foundation.INSArray)

	// The range of visible characters in the document.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleCharacterRange
	AccessibilityVisibleCharacterRange() foundation.NSRange
	SetAccessibilityVisibleCharacterRange(value foundation.NSRange)

	// The accessibility element’s visible child accessibility elements.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleChildren
	AccessibilityVisibleChildren() foundation.INSArray
	SetAccessibilityVisibleChildren(value foundation.INSArray)

	// The visible columns for the table or outline.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleColumns
	AccessibilityVisibleColumns() foundation.INSArray
	SetAccessibilityVisibleColumns(value foundation.INSArray)

	// The visible rows for the table or outline.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleRows
	AccessibilityVisibleRows() foundation.INSArray
	SetAccessibilityVisibleRows(value foundation.INSArray)

	// The warning value for the level indicator.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityWarningValue
	AccessibilityWarningValue() objectivec.IObject
	SetAccessibilityWarningValue(value objectivec.IObject)

	// The window that contains the accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityWindow
	AccessibilityWindow() objectivec.IObject
	SetAccessibilityWindow(value objectivec.IObject)

	// An array that contains all the app’s windows.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityWindows
	AccessibilityWindows() foundation.INSArray
	SetAccessibilityWindows(value foundation.INSArray)

	// The child accessibility element that represents the window’s zoom button.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityZoomButton
	AccessibilityZoomButton() objectivec.IObject
	SetAccessibilityZoomButton(value objectivec.IObject)
}

// NSAccessibilityProtocolObject wraps an existing Objective-C object that conforms to the NSAccessibilityProtocol protocol.
type NSAccessibilityProtocolObject struct {
	objectivec.Object
}

func (o NSAccessibilityProtocolObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSAccessibilityProtocolObjectFromID constructs a [NSAccessibilityProtocolObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSAccessibilityProtocolObjectFromID(id objc.ID) NSAccessibilityProtocolObject {
	return NSAccessibilityProtocolObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityElement()
func (o NSAccessibilityProtocolObject) IsAccessibilityElement() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityElement"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEnabled()
func (o NSAccessibilityProtocolObject) IsAccessibilityEnabled() bool {
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
func (o NSAccessibilityProtocolObject) IsAccessibilitySelectorAllowed(selector objc.SEL) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelectorAllowed:"), selector)
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityProtectedContent()
func (o NSAccessibilityProtocolObject) IsAccessibilityProtectedContent() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityProtectedContent"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// is currently in a selected state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelected()
func (o NSAccessibilityProtocolObject) IsAccessibilitySelected() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelected"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// has the keyboard focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityFocused()
func (o NSAccessibilityProtocolObject) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// must have content for successful submission of a form.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityRequired()
func (o NSAccessibilityProtocolObject) IsAccessibilityRequired() bool {
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
func (o NSAccessibilityProtocolObject) AccessibilityStringForRange(range_ foundation.NSRange) string {
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
func (o NSAccessibilityProtocolObject) AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString {
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
func (o NSAccessibilityProtocolObject) AccessibilityRTFForRange(range_ foundation.NSRange) foundation.NSData {
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
func (o NSAccessibilityProtocolObject) AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect {
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
func (o NSAccessibilityProtocolObject) AccessibilityLineForIndex(index int) int {
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
func (o NSAccessibilityProtocolObject) AccessibilityRangeForIndex(index int) foundation.NSRange {
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
func (o NSAccessibilityProtocolObject) AccessibilityStyleRangeForIndex(index int) foundation.NSRange {
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
func (o NSAccessibilityProtocolObject) AccessibilityRangeForLine(line int) foundation.NSRange {
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
func (o NSAccessibilityProtocolObject) AccessibilityRangeForPosition(point corefoundation.CGPoint) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForPosition:"), point)
	return rv
}

// Returns the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityAlternateUIVisible()
func (o NSAccessibilityProtocolObject) IsAccessibilityAlternateUIVisible() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
	return rv
}

// Returns a Boolean value that determines whether the window is the app’s
// main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMain()
func (o NSAccessibilityProtocolObject) IsAccessibilityMain() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMain"))
	return rv
}

// Returns the Boolean value that determines whether the window is in a
// minimized state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMinimized()
func (o NSAccessibilityProtocolObject) IsAccessibilityMinimized() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMinimized"))
	return rv
}

// Returns a Boolean value that determines whether the window is modal.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityModal()
func (o NSAccessibilityProtocolObject) IsAccessibilityModal() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityModal"))
	return rv
}

// Returns a Boolean value that determines whether the app is the frontmost
// app.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityFrontmost()
func (o NSAccessibilityProtocolObject) IsAccessibilityFrontmost() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFrontmost"))
	return rv
}

// Returns a Boolean value that determines whether the app is in a hidden
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityHidden()
func (o NSAccessibilityProtocolObject) IsAccessibilityHidden() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityHidden"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility
// element’s grid is in row major order or in column major order.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityOrderedByRow()
func (o NSAccessibilityProtocolObject) IsAccessibilityOrderedByRow() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityOrderedByRow"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// is in an expanded state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityExpanded()
func (o NSAccessibilityProtocolObject) IsAccessibilityExpanded() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityExpanded"))
	return rv
}

// Returns a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityDisclosed()
func (o NSAccessibilityProtocolObject) IsAccessibilityDisclosed() bool {
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
func (o NSAccessibilityProtocolObject) AccessibilityCellForColumnRow(column int, row int) objectivec.IObject {
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
func (o NSAccessibilityProtocolObject) AccessibilityLayoutPointForScreenPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
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
func (o NSAccessibilityProtocolObject) AccessibilityLayoutSizeForScreenSize(size corefoundation.CGSize) corefoundation.CGSize {
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
func (o NSAccessibilityProtocolObject) AccessibilityScreenPointForLayoutPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
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
func (o NSAccessibilityProtocolObject) AccessibilityScreenSizeForLayoutSize(size corefoundation.CGSize) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("accessibilityScreenSizeForLayoutSize:"), size)
	return rv
}

// Returns a Boolean value that indicates whether the accessibility element is
// in an edited state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEdited()
func (o NSAccessibilityProtocolObject) IsAccessibilityEdited() bool {
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
func (o NSAccessibilityProtocolObject) AccessibilityPerformCancel() bool {
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
func (o NSAccessibilityProtocolObject) AccessibilityPerformConfirm() bool {
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
func (o NSAccessibilityProtocolObject) AccessibilityPerformPick() bool {
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
func (o NSAccessibilityProtocolObject) AccessibilityPerformPress() bool {
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
func (o NSAccessibilityProtocolObject) AccessibilityPerformShowAlternateUI() bool {
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
func (o NSAccessibilityProtocolObject) AccessibilityPerformShowDefaultUI() bool {
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
func (o NSAccessibilityProtocolObject) AccessibilityPerformShowMenu() bool {
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
func (o NSAccessibilityProtocolObject) AccessibilityPerformRaise() bool {
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
func (o NSAccessibilityProtocolObject) AccessibilityPerformIncrement() bool {
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
func (o NSAccessibilityProtocolObject) AccessibilityPerformDecrement() bool {
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
func (o NSAccessibilityProtocolObject) AccessibilityPerformDelete() bool {
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
func (o NSAccessibilityProtocolObject) AccessibilityActivationPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityActivationPoint"))
	return corefoundation.CGPoint(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityActivationPoint(value corefoundation.CGPoint) {
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
func (o NSAccessibilityProtocolObject) AccessibilityAllowedValues() []foundation.NSNumber {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityAllowedValues"))
	result := make([]foundation.NSNumber, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = foundation.NSNumberFromID(id)
	}
	return result
}

func (o NSAccessibilityProtocolObject) SetAccessibilityAllowedValues(value []foundation.NSNumber) {
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
func (o NSAccessibilityProtocolObject) AccessibilityAlternateUIVisible() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
	return bool(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityAlternateUIVisible(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAlternateUIVisible:"), value)
}

// The child accessibility element with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityApplicationFocusedUIElement
func (o NSAccessibilityProtocolObject) AccessibilityApplicationFocusedUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityApplicationFocusedUIElement"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityApplicationFocusedUIElement(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityApplicationFocusedUIElement:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityAttributedUserInputLabels
func (o NSAccessibilityProtocolObject) AccessibilityAttributedUserInputLabels() []foundation.NSAttributedString {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityAttributedUserInputLabels"))
	result := make([]foundation.NSAttributedString, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = foundation.NSAttributedStringFromID(id)
	}
	return result
}

func (o NSAccessibilityProtocolObject) SetAccessibilityAttributedUserInputLabels(value []foundation.NSAttributedString) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAttributedUserInputLabels:"), objectivec.IObjectSliceToNSArray(value))
}

// The child accessibility element that represents the window’s cancel
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCancelButton
func (o NSAccessibilityProtocolObject) AccessibilityCancelButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCancelButton"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityCancelButton(value objectivec.IObject) {
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
func (o NSAccessibilityProtocolObject) AccessibilityChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityChildren"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityChildren(value foundation.INSArray) {
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
func (o NSAccessibilityProtocolObject) AccessibilityChildrenInNavigationOrder() []objectivec.IObject {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityChildrenInNavigationOrder"))
	result := make([]objectivec.IObject, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = objectivec.Object{ID: id}
	}
	return result
}

func (o NSAccessibilityProtocolObject) SetAccessibilityChildrenInNavigationOrder(value []objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildrenInNavigationOrder:"), objectivec.IObjectSliceToNSArray(value))
}

// The clear button for the search field.
//
// # Discussion
//
// Use this property on a search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityClearButton
func (o NSAccessibilityProtocolObject) AccessibilityClearButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityClearButton"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityClearButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityClearButton:"), value)
}

// The child accessibility element that represents the window’s close
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCloseButton
func (o NSAccessibilityProtocolObject) AccessibilityCloseButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCloseButton"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityCloseButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCloseButton:"), value)
}

// The number of columns in the accessibility element’s grid.
//
// # Discussion
//
// Use this property for UI elements that present a grid of child elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumnCount
func (o NSAccessibilityProtocolObject) AccessibilityColumnCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityColumnCount"))
	return int(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityColumnCount(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnCount:"), value)
}

// The column header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumnHeaderUIElements
func (o NSAccessibilityProtocolObject) AccessibilityColumnHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityColumnHeaderUIElements(value foundation.INSArray) {
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
func (o NSAccessibilityProtocolObject) AccessibilityColumnIndexRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityColumnIndexRange"))
	return foundation.NSRange(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityColumnIndexRange(value foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnIndexRange:"), value)
}

// The column titles for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumnTitles
func (o NSAccessibilityProtocolObject) AccessibilityColumnTitles() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnTitles"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityColumnTitles(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnTitles:"), value)
}

// The column accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumns
func (o NSAccessibilityProtocolObject) AccessibilityColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumns"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityColumns(value foundation.INSArray) {
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
func (o NSAccessibilityProtocolObject) AccessibilityContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityContents"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityContents(value foundation.INSArray) {
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
func (o NSAccessibilityProtocolObject) AccessibilityCriticalValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCriticalValue"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityCriticalValue(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCriticalValue:"), value)
}

// The custom actions of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCustomActions
func (o NSAccessibilityProtocolObject) AccessibilityCustomActions() []NSAccessibilityCustomAction {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityCustomActions"))
	result := make([]NSAccessibilityCustomAction, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = NSAccessibilityCustomActionFromID(id)
	}
	return result
}

func (o NSAccessibilityProtocolObject) SetAccessibilityCustomActions(value []NSAccessibilityCustomAction) {
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
func (o NSAccessibilityProtocolObject) AccessibilityCustomRotors() []NSAccessibilityCustomRotor {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityCustomRotors"))
	result := make([]NSAccessibilityCustomRotor, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = NSAccessibilityCustomRotorFromID(id)
	}
	return result
}

func (o NSAccessibilityProtocolObject) SetAccessibilityCustomRotors(value []NSAccessibilityCustomRotor) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomRotors:"), objectivec.IObjectSliceToNSArray(value))
}

// The decrement button for the stepper accessibility element.
//
// # Discussion
//
// Use this property on a stepper.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDecrementButton
func (o NSAccessibilityProtocolObject) AccessibilityDecrementButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDecrementButton"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityDecrementButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDecrementButton:"), value)
}

// The child accessibility element that represents the window’s default
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDefaultButton
func (o NSAccessibilityProtocolObject) AccessibilityDefaultButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDefaultButton"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityDefaultButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDefaultButton:"), value)
}

// A Boolean value that determines whether the row is disclosing other rows.
//
// # Discussion
//
// Use this property in the elements representing an outline’s row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosed
func (o NSAccessibilityProtocolObject) AccessibilityDisclosed() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityDisclosed"))
	return bool(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityDisclosed(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosed:"), value)
}

// The row disclosing the current row.
//
// # Discussion
//
// Use this property in the elements representing an outline’s row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosedByRow
func (o NSAccessibilityProtocolObject) AccessibilityDisclosedByRow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedByRow"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityDisclosedByRow(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedByRow:"), value)
}

// The rows that the current row discloses.
//
// # Discussion
//
// Use this property in the elements representing an outline’s row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosedRows
func (o NSAccessibilityProtocolObject) AccessibilityDisclosedRows() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedRows"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityDisclosedRows(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedRows:"), value)
}

// The indention level for the row.
//
// # Discussion
//
// Use this property in the elements representing an outline’s row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosureLevel
func (o NSAccessibilityProtocolObject) AccessibilityDisclosureLevel() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityDisclosureLevel"))
	return int(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityDisclosureLevel(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosureLevel:"), value)
}

// The URL for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDocument
func (o NSAccessibilityProtocolObject) AccessibilityDocument() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDocument"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSAccessibilityProtocolObject) SetAccessibilityDocument(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDocument:"), objc.String(value))
}

// A Boolean value that indicates whether the accessibility element is in an
// edited state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityEdited
func (o NSAccessibilityProtocolObject) AccessibilityEdited() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEdited"))
	return bool(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityEdited(value bool) {
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
func (o NSAccessibilityProtocolObject) AccessibilityElement() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityElement"))
	return bool(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityElement(value bool) {
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
func (o NSAccessibilityProtocolObject) AccessibilityEnabled() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEnabled"))
	return bool(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityEnabled(value bool) {
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
func (o NSAccessibilityProtocolObject) AccessibilityExpanded() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityExpanded"))
	return bool(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityExpanded(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExpanded:"), value)
}

// The icon for the app’s menu bar extra.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityExtrasMenuBar
func (o NSAccessibilityProtocolObject) AccessibilityExtrasMenuBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityExtrasMenuBar"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityExtrasMenuBar(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExtrasMenuBar:"), value)
}

// The filename for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFilename
func (o NSAccessibilityProtocolObject) AccessibilityFilename() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFilename"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSAccessibilityProtocolObject) SetAccessibilityFilename(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFilename:"), objc.String(value))
}

// A Boolean value that determines whether the accessibility element has the
// keyboard focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocused
func (o NSAccessibilityProtocolObject) AccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return bool(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityFocused(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocused:"), value)
}

// The child window with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocusedWindow
func (o NSAccessibilityProtocolObject) AccessibilityFocusedWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFocusedWindow"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityFocusedWindow(value objectivec.IObject) {
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
func (o NSAccessibilityProtocolObject) AccessibilityFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrame"))
	return corefoundation.CGRect(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityFrame(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrame:"), value)
}

// A Boolean value that determines whether the app is the frontmost app.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrontmost
func (o NSAccessibilityProtocolObject) AccessibilityFrontmost() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFrontmost"))
	return bool(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityFrontmost(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrontmost:"), value)
}

// The child accessibility element that represents the window’s full-screen
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFullScreenButton
func (o NSAccessibilityProtocolObject) AccessibilityFullScreenButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFullScreenButton"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityFullScreenButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFullScreenButton:"), value)
}

// The child accessibility element that represents the window’s grow area.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityGrowArea
func (o NSAccessibilityProtocolObject) AccessibilityGrowArea() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityGrowArea"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityGrowArea(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityGrowArea:"), value)
}

// The drag handle accessibility elements for the layout item element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHandles
func (o NSAccessibilityProtocolObject) AccessibilityHandles() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHandles"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityHandles(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHandles:"), value)
}

// The header for the table view.
//
// # Discussion
//
// Use this property on a table view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHeader
func (o NSAccessibilityProtocolObject) AccessibilityHeader() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHeader"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityHeader(value objectivec.IObject) {
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
func (o NSAccessibilityProtocolObject) AccessibilityHelp() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHelp"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSAccessibilityProtocolObject) SetAccessibilityHelp(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHelp:"), objc.String(value))
}

// A Boolean value that determines whether the app is in a hidden state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHidden
func (o NSAccessibilityProtocolObject) AccessibilityHidden() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityHidden"))
	return bool(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityHidden(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHidden:"), value)
}

// The horizontal scroll bar for the scroll view.
//
// # Discussion
//
// Use this property on a scrollable view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHorizontalScrollBar
func (o NSAccessibilityProtocolObject) AccessibilityHorizontalScrollBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalScrollBar"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityHorizontalScrollBar(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalScrollBar:"), value)
}

// A description of the layout area’s horizontal units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHorizontalUnitDescription
func (o NSAccessibilityProtocolObject) AccessibilityHorizontalUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSAccessibilityProtocolObject) SetAccessibilityHorizontalUnitDescription(value string) {
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
func (o NSAccessibilityProtocolObject) AccessibilityHorizontalUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityHorizontalUnits"))
	return NSAccessibilityUnits(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityHorizontalUnits(value NSAccessibilityUnits) {
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
func (o NSAccessibilityProtocolObject) AccessibilityIdentifier() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSAccessibilityProtocolObject) SetAccessibilityIdentifier(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIdentifier:"), objc.String(value))
}

// The increment button for the stepper accessibility element.
//
// # Discussion
//
// Use this property on a stepper.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIncrementButton
func (o NSAccessibilityProtocolObject) AccessibilityIncrementButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIncrementButton"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityIncrementButton(value objectivec.IObject) {
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
func (o NSAccessibilityProtocolObject) AccessibilityIndex() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityIndex"))
	return int(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityIndex(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIndex:"), value)
}

// The line number that contains the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityInsertionPointLineNumber
func (o NSAccessibilityProtocolObject) AccessibilityInsertionPointLineNumber() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityInsertionPointLineNumber"))
	return int(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityInsertionPointLineNumber(value int) {
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
func (o NSAccessibilityProtocolObject) AccessibilityLabel() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSAccessibilityProtocolObject) SetAccessibilityLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabel:"), objc.String(value))
}

// The child label elements for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLabelUIElements
func (o NSAccessibilityProtocolObject) AccessibilityLabelUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabelUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityLabelUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelUIElements:"), value)
}

// The value of the label accessibility element.
//
// # Discussion
//
// Use this property on a slider element’s labels.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLabelValue
func (o NSAccessibilityProtocolObject) AccessibilityLabelValue() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("accessibilityLabelValue"))
	return float32(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityLabelValue(value float32) {
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
func (o NSAccessibilityProtocolObject) AccessibilityLinkedUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLinkedUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityLinkedUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLinkedUIElements:"), value)
}

// A Boolean value that determines whether the window is the app’s main
// window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMain
func (o NSAccessibilityProtocolObject) AccessibilityMain() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMain"))
	return bool(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityMain(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMain:"), value)
}

// The app’s main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMainWindow
func (o NSAccessibilityProtocolObject) AccessibilityMainWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMainWindow"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityMainWindow(value objectivec.IObject) {
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
func (o NSAccessibilityProtocolObject) AccessibilityMarkerGroupUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerGroupUIElement"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityMarkerGroupUIElement(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerGroupUIElement:"), value)
}

// A human-readable description of the marker type.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMarkerTypeDescription
func (o NSAccessibilityProtocolObject) AccessibilityMarkerTypeDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerTypeDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSAccessibilityProtocolObject) SetAccessibilityMarkerTypeDescription(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerTypeDescription:"), objc.String(value))
}

// An array of marker accessibility elements for the ruler.
//
// # Discussion
//
// Use this property on a ruler element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMarkerUIElements
func (o NSAccessibilityProtocolObject) AccessibilityMarkerUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityMarkerUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerUIElements:"), value)
}

// The marker values for the ruler.
//
// # Discussion
//
// Use this property on a ruler element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMarkerValues
func (o NSAccessibilityProtocolObject) AccessibilityMarkerValues() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerValues"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityMarkerValues(value objectivec.IObject) {
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
func (o NSAccessibilityProtocolObject) AccessibilityMaxValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMaxValue"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityMaxValue(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMaxValue:"), value)
}

// The app’s menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMenuBar
func (o NSAccessibilityProtocolObject) AccessibilityMenuBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMenuBar"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityMenuBar(value objectivec.IObject) {
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
func (o NSAccessibilityProtocolObject) AccessibilityMinValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinValue"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityMinValue(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinValue:"), value)
}

// The child accessibility element that represents the window’s minimize
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMinimizeButton
func (o NSAccessibilityProtocolObject) AccessibilityMinimizeButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinimizeButton"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityMinimizeButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimizeButton:"), value)
}

// A Boolean value that determines whether this window is in a minimized
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMinimized
func (o NSAccessibilityProtocolObject) AccessibilityMinimized() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMinimized"))
	return bool(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityMinimized(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimized:"), value)
}

// A Boolean value that determines whether the window is modal.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityModal
func (o NSAccessibilityProtocolObject) AccessibilityModal() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityModal"))
	return bool(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityModal(value bool) {
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
func (o NSAccessibilityProtocolObject) AccessibilityNextContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityNextContents"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityNextContents(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNextContents:"), value)
}

// The number of characters in the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityNumberOfCharacters
func (o NSAccessibilityProtocolObject) AccessibilityNumberOfCharacters() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityNumberOfCharacters"))
	return int(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityNumberOfCharacters(value int) {
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
func (o NSAccessibilityProtocolObject) AccessibilityOrderedByRow() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityOrderedByRow"))
	return bool(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityOrderedByRow(value bool) {
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
func (o NSAccessibilityProtocolObject) AccessibilityOrientation() NSAccessibilityOrientation {
	rv := objc.Send[NSAccessibilityOrientation](o.ID, objc.Sel("accessibilityOrientation"))
	return NSAccessibilityOrientation(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityOrientation(value NSAccessibilityOrientation) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrientation:"), value)
}

// The overflow button for the toolbar.
//
// # Discussion
//
// Use this property on a toolbar element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityOverflowButton
func (o NSAccessibilityProtocolObject) AccessibilityOverflowButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityOverflowButton"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityOverflowButton(value objectivec.IObject) {
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
func (o NSAccessibilityProtocolObject) AccessibilityParent() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityParent"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityParent(value objectivec.IObject) {
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
func (o NSAccessibilityProtocolObject) AccessibilityPlaceholderValue() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPlaceholderValue"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSAccessibilityProtocolObject) SetAccessibilityPlaceholderValue(value string) {
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
func (o NSAccessibilityProtocolObject) AccessibilityPreviousContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPreviousContents"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityPreviousContents(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPreviousContents:"), value)
}

// A Boolean value that determines whether the accessibility element contains
// protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityProtectedContent
func (o NSAccessibilityProtocolObject) AccessibilityProtectedContent() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityProtectedContent"))
	return bool(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityProtectedContent(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProtectedContent:"), value)
}

// The child accessibility element that represents the window’s proxy icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityProxy
func (o NSAccessibilityProtocolObject) AccessibilityProxy() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityProxy"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityProxy(value objectivec.IObject) {
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
func (o NSAccessibilityProtocolObject) AccessibilityRequired() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityRequired"))
	return bool(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityRequired(value bool) {
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
func (o NSAccessibilityProtocolObject) AccessibilityRole() NSAccessibilityRole {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRole"))
	return NSAccessibilityRole(foundation.NSStringFromID(rv).String())
}

func (o NSAccessibilityProtocolObject) SetAccessibilityRole(value NSAccessibilityRole) {
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
func (o NSAccessibilityProtocolObject) AccessibilityRoleDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRoleDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSAccessibilityProtocolObject) SetAccessibilityRoleDescription(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRoleDescription:"), objc.String(value))
}

// The number of rows in the accessibility element’s grid.
//
// # Discussion
//
// Use this property for elements that present a grid of child elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRowCount
func (o NSAccessibilityProtocolObject) AccessibilityRowCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityRowCount"))
	return int(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityRowCount(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowCount:"), value)
}

// The row header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRowHeaderUIElements
func (o NSAccessibilityProtocolObject) AccessibilityRowHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRowHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityRowHeaderUIElements(value foundation.INSArray) {
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
func (o NSAccessibilityProtocolObject) AccessibilityRowIndexRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRowIndexRange"))
	return foundation.NSRange(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityRowIndexRange(value foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowIndexRange:"), value)
}

// The row accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRows
func (o NSAccessibilityProtocolObject) AccessibilityRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRows"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityRows(value foundation.INSArray) {
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
func (o NSAccessibilityProtocolObject) AccessibilityRulerMarkerType() NSAccessibilityRulerMarkerType {
	rv := objc.Send[NSAccessibilityRulerMarkerType](o.ID, objc.Sel("accessibilityRulerMarkerType"))
	return NSAccessibilityRulerMarkerType(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityRulerMarkerType(value NSAccessibilityRulerMarkerType) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRulerMarkerType:"), value)
}

// The search button for the search field.
//
// # Discussion
//
// Use this property on a search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySearchButton
func (o NSAccessibilityProtocolObject) AccessibilitySearchButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchButton"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilitySearchButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchButton:"), value)
}

// The search menu for the search field.
//
// # Discussion
//
// Use this property on a search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySearchMenu
func (o NSAccessibilityProtocolObject) AccessibilitySearchMenu() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchMenu"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilitySearchMenu(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchMenu:"), value)
}

// A Boolean value that determines whether the accessibility element is
// currently in a selected state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelected
func (o NSAccessibilityProtocolObject) AccessibilitySelected() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelected"))
	return bool(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilitySelected(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelected:"), value)
}

// The currently selected cells for the table.
//
// # Discussion
//
// This property is required for all elements that act like cell-based tables.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedCells
func (o NSAccessibilityProtocolObject) AccessibilitySelectedCells() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedCells"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilitySelectedCells(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedCells:"), value)
}

// The accessibility element’s currently selected children.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedChildren
func (o NSAccessibilityProtocolObject) AccessibilitySelectedChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedChildren"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilitySelectedChildren(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedChildren:"), value)
}

// The currently selected columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedColumns
func (o NSAccessibilityProtocolObject) AccessibilitySelectedColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedColumns"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilitySelectedColumns(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedColumns:"), value)
}

// The currently selected rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedRows
func (o NSAccessibilityProtocolObject) AccessibilitySelectedRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedRows"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilitySelectedRows(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedRows:"), value)
}

// The currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedText
func (o NSAccessibilityProtocolObject) AccessibilitySelectedText() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedText"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSAccessibilityProtocolObject) SetAccessibilitySelectedText(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedText:"), objc.String(value))
}

// The range of the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedTextRange
func (o NSAccessibilityProtocolObject) AccessibilitySelectedTextRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySelectedTextRange"))
	return foundation.NSRange(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilitySelectedTextRange(value foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRange:"), value)
}

// An array of ranges for the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedTextRanges
func (o NSAccessibilityProtocolObject) AccessibilitySelectedTextRanges() []foundation.NSValue {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilitySelectedTextRanges"))
	result := make([]foundation.NSValue, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = foundation.NSValueFromID(id)
	}
	return result
}

func (o NSAccessibilityProtocolObject) SetAccessibilitySelectedTextRanges(value []foundation.NSValue) {
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
func (o NSAccessibilityProtocolObject) AccessibilityServesAsTitleForUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityServesAsTitleForUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityServesAsTitleForUIElements(value foundation.INSArray) {
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
func (o NSAccessibilityProtocolObject) AccessibilitySharedCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySharedCharacterRange"))
	return foundation.NSRange(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilitySharedCharacterRange(value foundation.NSRange) {
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
func (o NSAccessibilityProtocolObject) AccessibilitySharedFocusElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedFocusElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilitySharedFocusElements(value foundation.INSArray) {
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
func (o NSAccessibilityProtocolObject) AccessibilitySharedTextUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedTextUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilitySharedTextUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedTextUIElements:"), value)
}

// The menu currently displaying for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityShownMenu
func (o NSAccessibilityProtocolObject) AccessibilityShownMenu() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityShownMenu"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityShownMenu(value objectivec.IObject) {
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
func (o NSAccessibilityProtocolObject) AccessibilitySortDirection() NSAccessibilitySortDirection {
	rv := objc.Send[NSAccessibilitySortDirection](o.ID, objc.Sel("accessibilitySortDirection"))
	return NSAccessibilitySortDirection(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilitySortDirection(value NSAccessibilitySortDirection) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySortDirection:"), value)
}

// An array that contains the views and splitter bar from the split view.
//
// # Discussion
//
// Use this property on a split view element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySplitters
func (o NSAccessibilityProtocolObject) AccessibilitySplitters() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySplitters"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilitySplitters(value foundation.INSArray) {
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
func (o NSAccessibilityProtocolObject) AccessibilitySubrole() NSAccessibilitySubrole {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySubrole"))
	return NSAccessibilitySubrole(foundation.NSStringFromID(rv).String())
}

func (o NSAccessibilityProtocolObject) SetAccessibilitySubrole(value NSAccessibilitySubrole) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySubrole:"), objc.String(string(value)))
}

// The tab accessibility elements for the tab view.
//
// # Discussion
//
// Use this property on a tab view element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityTabs
func (o NSAccessibilityProtocolObject) AccessibilityTabs() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTabs"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityTabs(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTabs:"), value)
}

// The title of the accessibility element—for example, a button’s visible
// text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityTitle
func (o NSAccessibilityProtocolObject) AccessibilityTitle() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitle"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSAccessibilityProtocolObject) SetAccessibilityTitle(value string) {
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
func (o NSAccessibilityProtocolObject) AccessibilityTitleUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitleUIElement"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityTitleUIElement(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitleUIElement:"), value)
}

// The child accessibility element that represents the window’s toolbar
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityToolbarButton
func (o NSAccessibilityProtocolObject) AccessibilityToolbarButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityToolbarButton"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityToolbarButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityToolbarButton:"), value)
}

// The top-level element that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityTopLevelUIElement
func (o NSAccessibilityProtocolObject) AccessibilityTopLevelUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTopLevelUIElement"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityTopLevelUIElement(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTopLevelUIElement:"), value)
}

// The URL for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityURL
func (o NSAccessibilityProtocolObject) AccessibilityURL() foundation.NSURL {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityURL"))
	return foundation.NSURLFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityURL(value foundation.NSURL) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityURL:"), value)
}

// A human-readable description of the ruler’s units.
//
// # Discussion
//
// Use this property on a ruler element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityUnitDescription
func (o NSAccessibilityProtocolObject) AccessibilityUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSAccessibilityProtocolObject) SetAccessibilityUnitDescription(value string) {
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
func (o NSAccessibilityProtocolObject) AccessibilityUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityUnits"))
	return NSAccessibilityUnits(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityUnits(value NSAccessibilityUnits) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnits:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityUserInputLabels
func (o NSAccessibilityProtocolObject) AccessibilityUserInputLabels() []string {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityUserInputLabels"))
	return objc.ConvertSliceToStrings(rvIDs)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityUserInputLabels(value []string) {
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
func (o NSAccessibilityProtocolObject) AccessibilityValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValue"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityValue(value objectivec.IObject) {
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
func (o NSAccessibilityProtocolObject) AccessibilityValueDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValueDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSAccessibilityProtocolObject) SetAccessibilityValueDescription(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValueDescription:"), objc.String(value))
}

// The vertical scroll bar for the scroll view.
//
// # Discussion
//
// Use this property on a scrollable view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVerticalScrollBar
func (o NSAccessibilityProtocolObject) AccessibilityVerticalScrollBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalScrollBar"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityVerticalScrollBar(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalScrollBar:"), value)
}

// A description of the layout area’s vertical units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVerticalUnitDescription
func (o NSAccessibilityProtocolObject) AccessibilityVerticalUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSAccessibilityProtocolObject) SetAccessibilityVerticalUnitDescription(value string) {
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
func (o NSAccessibilityProtocolObject) AccessibilityVerticalUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityVerticalUnits"))
	return NSAccessibilityUnits(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityVerticalUnits(value NSAccessibilityUnits) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnits:"), value)
}

// The visible cells for the table.
//
// # Discussion
//
// This property is required for all elements that act like cell-based tables.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleCells
func (o NSAccessibilityProtocolObject) AccessibilityVisibleCells() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleCells"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityVisibleCells(value foundation.INSArray) {
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
func (o NSAccessibilityProtocolObject) AccessibilityVisibleCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityVisibleCharacterRange"))
	return foundation.NSRange(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityVisibleCharacterRange(value foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleCharacterRange:"), value)
}

// The accessibility element’s visible child accessibility elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleChildren
func (o NSAccessibilityProtocolObject) AccessibilityVisibleChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleChildren"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityVisibleChildren(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleChildren:"), value)
}

// The visible columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleColumns
func (o NSAccessibilityProtocolObject) AccessibilityVisibleColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleColumns"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityVisibleColumns(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleColumns:"), value)
}

// The visible rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleRows
func (o NSAccessibilityProtocolObject) AccessibilityVisibleRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleRows"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityVisibleRows(value foundation.INSArray) {
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
func (o NSAccessibilityProtocolObject) AccessibilityWarningValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWarningValue"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityWarningValue(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWarningValue:"), value)
}

// The window that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityWindow
func (o NSAccessibilityProtocolObject) AccessibilityWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindow"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityWindow(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindow:"), value)
}

// An array that contains all the app’s windows.
//
// # Discussion
//
// Use on the app element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityWindows
func (o NSAccessibilityProtocolObject) AccessibilityWindows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindows"))
	return foundation.NSArrayFromID(rv)
}

func (o NSAccessibilityProtocolObject) SetAccessibilityWindows(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindows:"), value)
}

// The child accessibility element that represents the window’s zoom button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityZoomButton
func (o NSAccessibilityProtocolObject) AccessibilityZoomButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityZoomButton"))
	return objectivec.Object{ID: rv}
}

func (o NSAccessibilityProtocolObject) SetAccessibilityZoomButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityZoomButton:"), value)
}
