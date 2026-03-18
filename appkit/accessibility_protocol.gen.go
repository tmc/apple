// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The complete list of properties and methods for accessible elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol
type NSAccessibilityProtocol interface {
	objectivec.IObject

	// Returns a Boolean value that determines whether the accessibility element participates in the accessibility hierarchy.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityelement()
	IsAccessibilityElement() bool

	// Sets a Boolean value that determines whether the accessibility element participates in the accessibility hierarchy.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityelement(_:)
	SetAccessibilityElement(accessibilityElement bool)

	// Returns a Boolean value that determines whether the accessibility element responds to user events.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityenabled()
	IsAccessibilityEnabled() bool

	// Sets a Boolean value that determines whether the accessibility element responds to user events.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityenabled(_:)
	SetAccessibilityEnabled(accessibilityEnabled bool)

	// Returns the accessibility element’s frame in screen coordinates.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityframe()
	AccessibilityFrame() corefoundation.CGRect

	// Sets the accessibility element’s frame in screen coordinates.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityframe(_:)
	SetAccessibilityFrame(accessibilityFrame corefoundation.CGRect)

	// Returns the help text for the accessibility element.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhelp()
	AccessibilityHelp() string

	// Sets the help text for the accessibility element.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhelp(_:)
	SetAccessibilityHelp(accessibilityHelp string)

	// Returns a short description of the accessibility element.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylabel()
	AccessibilityLabel() string

	// Sets a short description of the accessibility element.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylabel(_:)
	SetAccessibilityLabel(accessibilityLabel string)

	// Returns the title of the accessibility element—for example, a button’s visible text.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytitle()
	AccessibilityTitle() string

	// Sets the title of the accessibility element.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytitle(_:)
	SetAccessibilityTitle(accessibilityTitle string)

	// Returns a Boolean value that indicates whether assistive apps can invoke the specified selector on the accessibility element.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelectorAllowed(_:)
	IsAccessibilitySelectorAllowed(selector objc.SEL) bool

	// Returns the accessibility element’s identity.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityidentifier()
	AccessibilityIdentifier() string

	// Sets the accessibility element’s identity.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityidentifier(_:)
	SetAccessibilityIdentifier(accessibilityIdentifier string)

	// Returns the orientation of the accessibility element.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityorientation()
	AccessibilityOrientation() NSAccessibilityOrientation

	// Sets the orientation of the accessibility element.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityorientation(_:)
	SetAccessibilityOrientation(accessibilityOrientation NSAccessibilityOrientation)

	// Returns a Boolean value that determines whether the accessibility element contains protected content.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityprotectedcontent()
	IsAccessibilityProtectedContent() bool

	// Sets a Boolean value that determines whether the accessibility element contains protected content.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityprotectedcontent(_:)
	SetAccessibilityProtectedContent(accessibilityProtectedContent bool)

	// Returns a Boolean value that determines whether the accessibility element is currently in a selected state.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityselected()
	IsAccessibilitySelected() bool

	// Sets a Boolean value that determines whether the accessibility element is currently in a selected state.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselected(_:)
	SetAccessibilitySelected(accessibilitySelected bool)

	// Returns the URL for the accessibility element.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityurl()
	AccessibilityURL() foundation.INSURL

	// Sets the URL for the accessibility element.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityurl(_:)
	SetAccessibilityURL(accessibilityURL foundation.INSURL)

	// Returns the human-readable description of the accessibility element’s value.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvaluedescription()
	AccessibilityValueDescription() string

	// Sets the human-readable description of the accessibility element’s value.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvaluedescription(_:)
	SetAccessibilityValueDescription(accessibilityValueDescription string)

	// Returns the array of child accessibility elements in order for linear navigation.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitychildreninnavigationorder()
	AccessibilityChildrenInNavigationOrder() unsafe.Pointer

	// Returns a Boolean value that determines whether the accessibility element has the keyboard focus.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityfocused()
	IsAccessibilityFocused() bool

	// Sets a Boolean value that determines whether the accessibility element has the keyboard focus.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfocused(_:)
	SetAccessibilityFocused(accessibilityFocused bool)

	// Returns a Boolean value that determines whether the accessibility element must have content for successful submission of a form.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityrequired()
	IsAccessibilityRequired() bool

	// Sets a Boolean value that determines whether the accessibility element must have content for successful submission of a form.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrequired(_:)
	SetAccessibilityRequired(accessibilityRequired bool)

	// Returns the type of interface element that the accessibility element represents.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrole()
	AccessibilityRole() NSAccessibilityRole

	// Sets the type of interface element that the accessibility element represents.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrole(_:)
	SetAccessibilityRole(accessibilityRole NSAccessibilityRole)

	// Returns a localized, human-intelligible description of the accessibility element’s role, such as
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityroledescription()
	AccessibilityRoleDescription() string

	// Sets the localized, human-intelligible description of the accessibility element’s role, such as
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityroledescription(_:)
	SetAccessibilityRoleDescription(accessibilityRoleDescription string)

	// Returns the specialized interface element type that the accessibility element represents.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysubrole()
	AccessibilitySubrole() NSAccessibilitySubrole

	// Sets the specialized interface element type that the accessibility element represents.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysubrole(_:)
	SetAccessibilitySubrole(accessibilitySubrole NSAccessibilitySubrole)

	// Returns the custom actions of the current accessibility element.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycustomactions()
	AccessibilityCustomActions() INSAccessibilityCustomAction

	// Returns the custom rotors of the current accessibility element.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycustomrotors()
	AccessibilityCustomRotors() INSAccessibilityCustomRotor

	// Returns the line number that contains the insertion point.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityinsertionpointlinenumber()
	AccessibilityInsertionPointLineNumber() int

	// Sets the line number that contains the insertion point.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityinsertionpointlinenumber(_:)
	SetAccessibilityInsertionPointLineNumber(accessibilityInsertionPointLineNumber int)

	// Returns the number of characters in the text.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitynumberofcharacters()
	AccessibilityNumberOfCharacters() int

	// Sets the number of characters in the text.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitynumberofcharacters(_:)
	SetAccessibilityNumberOfCharacters(accessibilityNumberOfCharacters int)

	// Returns the placeholder value for the accessibility element.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityplaceholdervalue()
	AccessibilityPlaceholderValue() string

	// Sets the placeholder value for the accessibility element.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityplaceholdervalue(_:)
	SetAccessibilityPlaceholderValue(accessibilityPlaceholderValue string)

	// Returns the currently selected text.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtext()
	AccessibilitySelectedText() string

	// Sets the currently selected text.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedtext(_:)
	SetAccessibilitySelectedText(accessibilitySelectedText string)

	// Returns the range of the currently selected text.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtextrange()
	AccessibilitySelectedTextRange() foundation.NSRange

	// Sets the range of the currently selected text.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedtextrange(_:)
	SetAccessibilitySelectedTextRange(accessibilitySelectedTextRange foundation.NSRange)

	// Returns an array of ranges for the currently selected text.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtextranges()
	AccessibilitySelectedTextRanges() foundation.NSValue

	// Returns the range of characters that the accessibility element displays.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysharedcharacterrange()
	AccessibilitySharedCharacterRange() foundation.NSRange

	// Sets the range of characters that the accessibility element displays.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysharedcharacterrange(_:)
	SetAccessibilitySharedCharacterRange(accessibilitySharedCharacterRange foundation.NSRange)

	// Returns the range of visible characters in the document.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblecharacterrange()
	AccessibilityVisibleCharacterRange() foundation.NSRange

	// Sets the range of visible characters in the document.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblecharacterrange(_:)
	SetAccessibilityVisibleCharacterRange(accessibilityVisibleCharacterRange foundation.NSRange)

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
	AccessibilityRTFForRange(range_ foundation.NSRange) foundation.INSData

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

	// Returns the activation point for the user interface element.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityactivationpoint()
	AccessibilityActivationPoint() corefoundation.CGPoint

	// Sets the activation point for the user interface element.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityactivationpoint(_:)
	SetAccessibilityActivationPoint(accessibilityActivationPoint corefoundation.CGPoint)

	// Returns the Boolean value that determines whether the accessibility element’s alternative UI is currently visible.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityalternateuivisible()
	IsAccessibilityAlternateUIVisible() bool

	// Sets the Boolean value that determines whether the accessibility element’s alternative UI is currently visible.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityalternateuivisible(_:)
	SetAccessibilityAlternateUIVisible(accessibilityAlternateUIVisible bool)

	// Returns a Boolean value that determines whether the window is the app’s main window.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitymain()
	IsAccessibilityMain() bool

	// Sets a Boolean value that determines whether the window is the app’s main window.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymain(_:)
	SetAccessibilityMain(accessibilityMain bool)

	// Returns the Boolean value that determines whether the window is in a minimized state.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityminimized()
	IsAccessibilityMinimized() bool

	// Sets the Boolean value that determines whether the window is in a minimized state.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityminimized(_:)
	SetAccessibilityMinimized(accessibilityMinimized bool)

	// Returns a Boolean value that determines whether the window is modal.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitymodal()
	IsAccessibilityModal() bool

	// Sets a Boolean value that determines whether the window is modal.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymodal(_:)
	SetAccessibilityModal(accessibilityModal bool)

	// Returns a Boolean value that determines whether the app is the frontmost app.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityfrontmost()
	IsAccessibilityFrontmost() bool

	// Sets a Boolean value that determines whether the app is the frontmost app.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfrontmost(_:)
	SetAccessibilityFrontmost(accessibilityFrontmost bool)

	// Returns a Boolean value that determines whether the app is in a hidden state.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityhidden()
	IsAccessibilityHidden() bool

	// Sets a Boolean value that determines whether the app is in a hidden state.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhidden(_:)
	SetAccessibilityHidden(accessibilityHidden bool)

	// Returns the number of columns in the accessibility element’s grid.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumncount()
	AccessibilityColumnCount() int

	// Sets the number of columns in the accessibility element’s grid.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumncount(_:)
	SetAccessibilityColumnCount(accessibilityColumnCount int)

	// Returns a Boolean value that determines whether the accessibility element’s grid is in row major order or in column major order.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityorderedbyrow()
	IsAccessibilityOrderedByRow() bool

	// Sets a Boolean value that determines whether the element’s grid is in row major order or in column major order.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityorderedbyrow(_:)
	SetAccessibilityOrderedByRow(accessibilityOrderedByRow bool)

	// Returns the number of rows in the accessibility element’s grid.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrowcount()
	AccessibilityRowCount() int

	// Sets the number of rows in the accessibility element’s grid.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrowcount(_:)
	SetAccessibilityRowCount(accessibilityRowCount int)

	// Returns a Boolean value that determines whether the accessibility element is in an expanded state.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityexpanded()
	IsAccessibilityExpanded() bool

	// Sets a Boolean value that determines whether accessibility element is in an expanded state.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityexpanded(_:)
	SetAccessibilityExpanded(accessibilityExpanded bool)

	// Returns the index of the row or column that the accessibility element represents.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityindex()
	AccessibilityIndex() int

	// Sets the index of the row or column that the accessibility element represents.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityindex(_:)
	SetAccessibilityIndex(accessibilityIndex int)

	// Returns the accessibility element’s sort direction.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysortdirection()
	AccessibilitySortDirection() NSAccessibilitySortDirection

	// Sets the accessibility element’s sort direction.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysortdirection(_:)
	SetAccessibilitySortDirection(accessibilitySortDirection NSAccessibilitySortDirection)

	// Returns a Boolean value that determines whether the row is disclosing other rows.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitydisclosed()
	IsAccessibilityDisclosed() bool

	// Sets a Boolean value that determines whether the row is disclosing other rows.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosed(_:)
	SetAccessibilityDisclosed(accessibilityDisclosed bool)

	// Returns the indention level for the row.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydisclosurelevel()
	AccessibilityDisclosureLevel() int

	// Sets the indention level for the row.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosurelevel(_:)
	SetAccessibilityDisclosureLevel(accessibilityDisclosureLevel int)

	// Returns the column index range of the cell.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumnindexrange()
	AccessibilityColumnIndexRange() foundation.NSRange

	// Sets the column index range of the cell.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumnindexrange(_:)
	SetAccessibilityColumnIndexRange(accessibilityColumnIndexRange foundation.NSRange)

	// Returns the row index range of the cell.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrowindexrange()
	AccessibilityRowIndexRange() foundation.NSRange

	// Sets the row index range of the cell.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrowindexrange(_:)
	SetAccessibilityRowIndexRange(accessibilityRowIndexRange foundation.NSRange)

	// Returns the units that the layout area uses for horizontal values.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhorizontalunits()
	AccessibilityHorizontalUnits() NSAccessibilityUnits

	// Sets the units that the layout area uses for horizontal values.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhorizontalunits(_:)
	SetAccessibilityHorizontalUnits(accessibilityHorizontalUnits NSAccessibilityUnits)

	// Returns the description of the layout area’s horizontal units.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhorizontalunitdescription()
	AccessibilityHorizontalUnitDescription() string

	// Sets the description of the layout area’s horizontal units.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhorizontalunitdescription(_:)
	SetAccessibilityHorizontalUnitDescription(accessibilityHorizontalUnitDescription string)

	// Returns the units that the layout area uses for vertical values.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityverticalunits()
	AccessibilityVerticalUnits() NSAccessibilityUnits

	// Sets the units that the layout area uses for vertical values.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityverticalunits(_:)
	SetAccessibilityVerticalUnits(accessibilityVerticalUnits NSAccessibilityUnits)

	// Returns the description of the layout area’s vertical units.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityverticalunitdescription()
	AccessibilityVerticalUnitDescription() string

	// Sets the description of the layout area’s vertical units.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityverticalunitdescription(_:)
	SetAccessibilityVerticalUnitDescription(accessibilityVerticalUnitDescription string)

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

	// Returns the allowed values for the slider accessibility element.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityallowedvalues()
	AccessibilityAllowedValues() foundation.NSNumber

	// Returns the value of the label accessibility element.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylabelvalue()
	AccessibilityLabelValue() float64

	// Sets the value of the label accessibility element.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylabelvalue(_:)
	SetAccessibilityLabelValue(accessibilityLabelValue float64)

	// Returns the human-readable description of the marker type.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkertypedescription()
	AccessibilityMarkerTypeDescription() string

	// Sets the human-readable description of the marker type.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkertypedescription(_:)
	SetAccessibilityMarkerTypeDescription(accessibilityMarkerTypeDescription string)

	// Returns the type of markers for the ruler.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrulermarkertype()
	AccessibilityRulerMarkerType() NSAccessibilityRulerMarkerType

	// Sets the type of markers for the ruler.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrulermarkertype(_:)
	SetAccessibilityRulerMarkerType(accessibilityRulerMarkerType NSAccessibilityRulerMarkerType)

	// Returns the units for the ruler.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityunits()
	AccessibilityUnits() NSAccessibilityUnits

	// Sets the units used for the ruler.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityunits(_:)
	SetAccessibilityUnits(accessibilityUnits NSAccessibilityUnits)

	// Returns the human-readable description of the ruler’s units.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityunitdescription()
	AccessibilityUnitDescription() string

	// Sets the human-readable description of the ruler’s units.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityunitdescription(_:)
	SetAccessibilityUnitDescription(accessibilityUnitDescription string)

	// Returns the URL for the file that the accessibility element represents.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydocument()
	AccessibilityDocument() string

	// Sets the URL for the file that the accessibility element represents.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydocument(_:)
	SetAccessibilityDocument(accessibilityDocument string)

	// Returns a Boolean value that indicates whether the accessibility element is in an edited state.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityedited()
	IsAccessibilityEdited() bool

	// Sets a Boolean value that indicates whether the accessibility element is in an edited state.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityedited(_:)
	SetAccessibilityEdited(accessibilityEdited bool)

	// Returns the filename for the file that the accessibility element represents.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityfilename()
	AccessibilityFilename() string

	// Sets the filename for the file that the accessibility element represents.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfilename(_:)
	SetAccessibilityFilename(accessibilityFilename string)

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

	// AccessibilityAttributedUserInputLabels protocol.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityattributeduserinputlabels()
	AccessibilityAttributedUserInputLabels() foundation.NSAttributedString

	// AccessibilityUserInputLabels protocol.
	//
	// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityuserinputlabels()
	AccessibilityUserInputLabels() string
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
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityelement()

func (o NSAccessibilityProtocolObject) IsAccessibilityElement() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityElement"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityelement(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityElement(accessibilityElement bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityElement:"), accessibilityElement)
	}

// Returns a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityenabled()

func (o NSAccessibilityProtocolObject) IsAccessibilityEnabled() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEnabled"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityenabled(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityEnabled(accessibilityEnabled bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEnabled:"), accessibilityEnabled)
	}

// Returns the accessibility element’s frame in screen coordinates.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityframe()

func (o NSAccessibilityProtocolObject) AccessibilityFrame() corefoundation.CGRect {
	
	rv := objc.Send[corefoundation.CGRect](o.ID, objc.Sel("accessibilityFrame"))
	return rv
	}

// Sets the accessibility element’s frame in screen coordinates.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityframe(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityFrame(accessibilityFrame corefoundation.CGRect) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrame:"), accessibilityFrame)
	}

// Returns the help text for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhelp()

func (o NSAccessibilityProtocolObject) AccessibilityHelp() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHelp"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the help text for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhelp(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityHelp(accessibilityHelp string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHelp:"), objc.String(accessibilityHelp))
	}

// Returns a short description of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylabel()

func (o NSAccessibilityProtocolObject) AccessibilityLabel() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets a short description of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylabel(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityLabel(accessibilityLabel string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabel:"), objc.String(accessibilityLabel))
	}

// Returns the title of the accessibility element—for example, a button’s
// visible text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytitle()

func (o NSAccessibilityProtocolObject) AccessibilityTitle() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitle"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the title of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytitle(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityTitle(accessibilityTitle string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitle:"), objc.String(accessibilityTitle))
	}

// Returns the accessibility element’s value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvalue()

func (o NSAccessibilityProtocolObject) AccessibilityValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValue"))
	return objectivec.Object{ID: rv}
	}

// Sets the accessibility element’s value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvalue(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityValue(accessibilityValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValue:"), accessibilityValue)
	}

// Returns a Boolean value that indicates whether assistive apps can invoke
// the specified selector on the accessibility element.
//
// selector: The selector to check.
//
// # Return Value
// 
// [true], if accessibility clients can call the selector; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelectorAllowed(_:)

func (o NSAccessibilityProtocolObject) IsAccessibilitySelectorAllowed(selector objc.SEL) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelectorAllowed:"), selector)
	return rv
	}

// Returns the contents of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycontents()

func (o NSAccessibilityProtocolObject) AccessibilityContents() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityContents"))
	return objectivec.Object{ID: rv}
	}

// Sets the contents of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycontents(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityContents(accessibilityContents foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityContents:"), accessibilityContents)
	}

// Returns the critical value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycriticalvalue()

func (o NSAccessibilityProtocolObject) AccessibilityCriticalValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCriticalValue"))
	return objectivec.Object{ID: rv}
	}

// Sets the critical value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycriticalvalue(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityCriticalValue(accessibilityCriticalValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCriticalValue:"), accessibilityCriticalValue)
	}

// Returns the accessibility element’s identity.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityidentifier()

func (o NSAccessibilityProtocolObject) AccessibilityIdentifier() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIdentifier"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the accessibility element’s identity.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityidentifier(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityIdentifier(accessibilityIdentifier string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIdentifier:"), objc.String(accessibilityIdentifier))
	}

// Returns the maximum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymaxvalue()

func (o NSAccessibilityProtocolObject) AccessibilityMaxValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMaxValue"))
	return objectivec.Object{ID: rv}
	}

// Sets the maximum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymaxvalue(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityMaxValue(accessibilityMaxValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMaxValue:"), accessibilityMaxValue)
	}

// Returns the minimum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityminvalue()

func (o NSAccessibilityProtocolObject) AccessibilityMinValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinValue"))
	return objectivec.Object{ID: rv}
	}

// Sets the minimum value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityminvalue(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityMinValue(accessibilityMinValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinValue:"), accessibilityMinValue)
	}

// Returns the orientation of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityorientation()

func (o NSAccessibilityProtocolObject) AccessibilityOrientation() NSAccessibilityOrientation {
	
	rv := objc.Send[NSAccessibilityOrientation](o.ID, objc.Sel("accessibilityOrientation"))
	return rv
	}

// Sets the orientation of the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityorientation(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityOrientation(accessibilityOrientation NSAccessibilityOrientation) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrientation:"), accessibilityOrientation)
	}

// Returns a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityprotectedcontent()

func (o NSAccessibilityProtocolObject) IsAccessibilityProtectedContent() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityProtectedContent"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityprotectedcontent(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityProtectedContent(accessibilityProtectedContent bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProtectedContent:"), accessibilityProtectedContent)
	}

// Returns a Boolean value that determines whether the accessibility element
// is currently in a selected state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityselected()

func (o NSAccessibilityProtocolObject) IsAccessibilitySelected() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelected"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element is
// currently in a selected state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselected(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilitySelected(accessibilitySelected bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelected:"), accessibilitySelected)
	}

// Returns the URL for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityurl()

func (o NSAccessibilityProtocolObject) AccessibilityURL() foundation.INSURL {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityURL"))
	return foundation.NSURLFromID(rv)
	}

// Sets the URL for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityurl(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityURL(accessibilityURL foundation.INSURL) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityURL:"), accessibilityURL)
	}

// Returns the human-readable description of the accessibility element’s
// value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvaluedescription()

func (o NSAccessibilityProtocolObject) AccessibilityValueDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValueDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the human-readable description of the accessibility element’s value.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvaluedescription(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityValueDescription(accessibilityValueDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValueDescription:"), objc.String(accessibilityValueDescription))
	}

// Returns the warning value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitywarningvalue()

func (o NSAccessibilityProtocolObject) AccessibilityWarningValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWarningValue"))
	return objectivec.Object{ID: rv}
	}

// Sets the warning value for the level indicator.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitywarningvalue(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityWarningValue(accessibilityWarningValue objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWarningValue:"), accessibilityWarningValue)
	}

// Returns the child accessibility elements in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitychildren()

func (o NSAccessibilityProtocolObject) AccessibilityChildren() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityChildren"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility elements in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitychildren(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityChildren(accessibilityChildren foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildren:"), accessibilityChildren)
	}

// Returns the array of child accessibility elements in order for linear
// navigation.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitychildreninnavigationorder()

func (o NSAccessibilityProtocolObject) AccessibilityChildrenInNavigationOrder() unsafe.Pointer {
	
	rv := objc.Send[unsafe.Pointer](o.ID, objc.Sel("accessibilityChildrenInNavigationOrder"))
	return rv
	}

// Sets the array of child accessibility elements in order for linear
// navigation.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitychildreninnavigationorder(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityChildrenInNavigationOrder(accessibilityChildrenInNavigationOrder foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildrenInNavigationOrder:"), accessibilityChildrenInNavigationOrder)
	}

// Returns the accessibility element’s parent in the accessibility
// hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityparent()

func (o NSAccessibilityProtocolObject) AccessibilityParent() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityParent"))
	return objectivec.Object{ID: rv}
	}

// Sets the accessibility element’s parent in the accessibility hierarchy.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityparent(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityParent(accessibilityParent objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityParent:"), accessibilityParent)
	}

// Returns the accessibility element’s currently selected children.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedchildren()

func (o NSAccessibilityProtocolObject) AccessibilitySelectedChildren() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedChildren"))
	return objectivec.Object{ID: rv}
	}

// Sets the accessibility element’s currently selected children.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedchildren(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilitySelectedChildren(accessibilitySelectedChildren foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedChildren:"), accessibilitySelectedChildren)
	}

// Returns the top-level element that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytopleveluielement()

func (o NSAccessibilityProtocolObject) AccessibilityTopLevelUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTopLevelUIElement"))
	return objectivec.Object{ID: rv}
	}

// Sets the top-level element that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytopleveluielement(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityTopLevelUIElement(accessibilityTopLevelUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTopLevelUIElement:"), accessibilityTopLevelUIElement)
	}

// Returns the accessibility element’s visible child accessibility elements.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblechildren()

func (o NSAccessibilityProtocolObject) AccessibilityVisibleChildren() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleChildren"))
	return objectivec.Object{ID: rv}
	}

// Sets the accessibility element’s visible child accessibility elements.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblechildren(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityVisibleChildren(accessibilityVisibleChildren foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleChildren:"), accessibilityVisibleChildren)
	}

// Returns the child accessibility element with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityapplicationfocuseduielement()

func (o NSAccessibilityProtocolObject) AccessibilityApplicationFocusedUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityApplicationFocusedUIElement"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityapplicationfocuseduielement(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityApplicationFocusedUIElement(accessibilityApplicationFocusedUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityApplicationFocusedUIElement:"), accessibilityApplicationFocusedUIElement)
	}

// Returns a Boolean value that determines whether the accessibility element
// has the keyboard focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityfocused()

func (o NSAccessibilityProtocolObject) IsAccessibilityFocused() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element has
// the keyboard focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfocused(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityFocused(accessibilityFocused bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocused:"), accessibilityFocused)
	}

// Returns the child window with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityfocusedwindow()

func (o NSAccessibilityProtocolObject) AccessibilityFocusedWindow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFocusedWindow"))
	return objectivec.Object{ID: rv}
	}

// Sets the child window with the current focus.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfocusedwindow(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityFocusedWindow(accessibilityFocusedWindow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocusedWindow:"), accessibilityFocusedWindow)
	}

// Returns the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysharedfocuselements()

func (o NSAccessibilityProtocolObject) AccessibilitySharedFocusElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedFocusElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the array of elements that shares the keyboard focus with the
// accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysharedfocuselements(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilitySharedFocusElements(accessibilitySharedFocusElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedFocusElements:"), accessibilitySharedFocusElements)
	}

// Returns a Boolean value that determines whether the accessibility element
// must have content for successful submission of a form.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityrequired()

func (o NSAccessibilityProtocolObject) IsAccessibilityRequired() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityRequired"))
	return rv
	}

// Sets a Boolean value that determines whether the accessibility element must
// have content for successful submission of a form.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrequired(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityRequired(accessibilityRequired bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRequired:"), accessibilityRequired)
	}

// Returns the type of interface element that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrole()

func (o NSAccessibilityProtocolObject) AccessibilityRole() NSAccessibilityRole {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRole"))
	return NSAccessibilityRole(foundation.NSStringFromID(rv).String())
	}

// Sets the type of interface element that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrole(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityRole(accessibilityRole NSAccessibilityRole) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRole:"), objc.String(string(accessibilityRole)))
	}

// Returns a localized, human-intelligible description of the accessibility
// element’s role, such as
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityroledescription()

func (o NSAccessibilityProtocolObject) AccessibilityRoleDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRoleDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the localized, human-intelligible description of the accessibility
// element’s role, such as
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityroledescription(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityRoleDescription(accessibilityRoleDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRoleDescription:"), objc.String(accessibilityRoleDescription))
	}

// Returns the specialized interface element type that the accessibility
// element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysubrole()

func (o NSAccessibilityProtocolObject) AccessibilitySubrole() NSAccessibilitySubrole {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySubrole"))
	return NSAccessibilitySubrole(foundation.NSStringFromID(rv).String())
	}

// Sets the specialized interface element type that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysubrole(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilitySubrole(accessibilitySubrole NSAccessibilitySubrole) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySubrole:"), objc.String(string(accessibilitySubrole)))
	}

// Returns the custom actions of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycustomactions()

func (o NSAccessibilityProtocolObject) AccessibilityCustomActions() INSAccessibilityCustomAction {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCustomActions"))
	return NSAccessibilityCustomActionFromID(rv)
	}

// Sets the custom actions of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycustomactions(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityCustomActions(accessibilityCustomActions foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomActions:"), accessibilityCustomActions)
	}

// Returns the custom rotors of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycustomrotors()

func (o NSAccessibilityProtocolObject) AccessibilityCustomRotors() INSAccessibilityCustomRotor {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCustomRotors"))
	return NSAccessibilityCustomRotorFromID(rv)
	}

// Sets the custom rotors of the current accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycustomrotors(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityCustomRotors(accessibilityCustomRotors foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomRotors:"), accessibilityCustomRotors)
	}

// Returns the line number that contains the insertion point.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityinsertionpointlinenumber()

func (o NSAccessibilityProtocolObject) AccessibilityInsertionPointLineNumber() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityInsertionPointLineNumber"))
	return rv
	}

// Sets the line number that contains the insertion point.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityinsertionpointlinenumber(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityInsertionPointLineNumber(accessibilityInsertionPointLineNumber int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityInsertionPointLineNumber:"), accessibilityInsertionPointLineNumber)
	}

// Returns the number of characters in the text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitynumberofcharacters()

func (o NSAccessibilityProtocolObject) AccessibilityNumberOfCharacters() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityNumberOfCharacters"))
	return rv
	}

// Sets the number of characters in the text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitynumberofcharacters(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityNumberOfCharacters(accessibilityNumberOfCharacters int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNumberOfCharacters:"), accessibilityNumberOfCharacters)
	}

// Returns the placeholder value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityplaceholdervalue()

func (o NSAccessibilityProtocolObject) AccessibilityPlaceholderValue() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPlaceholderValue"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the placeholder value for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityplaceholdervalue(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityPlaceholderValue(accessibilityPlaceholderValue string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPlaceholderValue:"), objc.String(accessibilityPlaceholderValue))
	}

// Returns the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtext()

func (o NSAccessibilityProtocolObject) AccessibilitySelectedText() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedText"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedtext(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilitySelectedText(accessibilitySelectedText string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedText:"), objc.String(accessibilitySelectedText))
	}

// Returns the range of the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtextrange()

func (o NSAccessibilityProtocolObject) AccessibilitySelectedTextRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySelectedTextRange"))
	return rv
	}

// Sets the range of the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedtextrange(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilitySelectedTextRange(accessibilitySelectedTextRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRange:"), accessibilitySelectedTextRange)
	}

// Returns an array of ranges for the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedtextranges()

func (o NSAccessibilityProtocolObject) AccessibilitySelectedTextRanges() foundation.NSValue {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedTextRanges"))
	return foundation.NSValueFromID(rv)
	}

// Sets an array of ranges for the currently selected text.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedtextranges(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilitySelectedTextRanges(accessibilitySelectedTextRanges foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRanges:"), accessibilitySelectedTextRanges)
	}

// Returns the range of characters that the accessibility element displays.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysharedcharacterrange()

func (o NSAccessibilityProtocolObject) AccessibilitySharedCharacterRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySharedCharacterRange"))
	return rv
	}

// Sets the range of characters that the accessibility element displays.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysharedcharacterrange(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilitySharedCharacterRange(accessibilitySharedCharacterRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedCharacterRange:"), accessibilitySharedCharacterRange)
	}

// Returns the other elements that share text with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysharedtextuielements()

func (o NSAccessibilityProtocolObject) AccessibilitySharedTextUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedTextUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the other elements that share text with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysharedtextuielements(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilitySharedTextUIElements(accessibilitySharedTextUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedTextUIElements:"), accessibilitySharedTextUIElements)
	}

// Returns the range of visible characters in the document.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblecharacterrange()

func (o NSAccessibilityProtocolObject) AccessibilityVisibleCharacterRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityVisibleCharacterRange"))
	return rv
	}

// Sets the range of visible characters in the document.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblecharacterrange(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityVisibleCharacterRange(accessibilityVisibleCharacterRange foundation.NSRange) {
	
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

func (o NSAccessibilityProtocolObject) AccessibilityRTFForRange(range_ foundation.NSRange) foundation.INSData {
	
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

// Returns the activation point for the user interface element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityactivationpoint()

func (o NSAccessibilityProtocolObject) AccessibilityActivationPoint() corefoundation.CGPoint {
	
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityActivationPoint"))
	return rv
	}

// Sets the activation point for the user interface element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityactivationpoint(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityActivationPoint(accessibilityActivationPoint corefoundation.CGPoint) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityActivationPoint:"), accessibilityActivationPoint)
	}

// Returns the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityalternateuivisible()

func (o NSAccessibilityProtocolObject) IsAccessibilityAlternateUIVisible() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
	return rv
	}

// Sets the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityalternateuivisible(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityAlternateUIVisible(accessibilityAlternateUIVisible bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAlternateUIVisible:"), accessibilityAlternateUIVisible)
	}

// Returns the child accessibility element that represents the window’s
// cancel button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycancelbutton()

func (o NSAccessibilityProtocolObject) AccessibilityCancelButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCancelButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s cancel
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycancelbutton(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityCancelButton(accessibilityCancelButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCancelButton:"), accessibilityCancelButton)
	}

// Returns the child accessibility element that represents the window’s
// close button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityclosebutton()

func (o NSAccessibilityProtocolObject) AccessibilityCloseButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCloseButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s close
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityclosebutton(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityCloseButton(accessibilityCloseButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCloseButton:"), accessibilityCloseButton)
	}

// Returns the child accessibility element that represents the window’s
// default button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydefaultbutton()

func (o NSAccessibilityProtocolObject) AccessibilityDefaultButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDefaultButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s default
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydefaultbutton(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityDefaultButton(accessibilityDefaultButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDefaultButton:"), accessibilityDefaultButton)
	}

// Returns the child accessibility element that represents the window’s
// full-screen button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityfullscreenbutton()

func (o NSAccessibilityProtocolObject) AccessibilityFullScreenButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFullScreenButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s
// full-screen button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfullscreenbutton(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityFullScreenButton(accessibilityFullScreenButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFullScreenButton:"), accessibilityFullScreenButton)
	}

// Returns the child accessibility element that represents the window’s grow
// area.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitygrowarea()

func (o NSAccessibilityProtocolObject) AccessibilityGrowArea() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityGrowArea"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s grow
// area.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitygrowarea(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityGrowArea(accessibilityGrowArea objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityGrowArea:"), accessibilityGrowArea)
	}

// Returns a Boolean value that determines whether the window is the app’s
// main window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitymain()

func (o NSAccessibilityProtocolObject) IsAccessibilityMain() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMain"))
	return rv
	}

// Sets a Boolean value that determines whether the window is the app’s main
// window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymain(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityMain(accessibilityMain bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMain:"), accessibilityMain)
	}

// Returns the child accessibility element that represents the window’s
// minimize button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityminimizebutton()

func (o NSAccessibilityProtocolObject) AccessibilityMinimizeButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinimizeButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s
// minimize button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityminimizebutton(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityMinimizeButton(accessibilityMinimizeButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimizeButton:"), accessibilityMinimizeButton)
	}

// Returns the Boolean value that determines whether the window is in a
// minimized state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityminimized()

func (o NSAccessibilityProtocolObject) IsAccessibilityMinimized() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMinimized"))
	return rv
	}

// Sets the Boolean value that determines whether the window is in a minimized
// state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityminimized(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityMinimized(accessibilityMinimized bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimized:"), accessibilityMinimized)
	}

// Returns a Boolean value that determines whether the window is modal.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitymodal()

func (o NSAccessibilityProtocolObject) IsAccessibilityModal() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityModal"))
	return rv
	}

// Sets a Boolean value that determines whether the window is modal.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymodal(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityModal(accessibilityModal bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityModal:"), accessibilityModal)
	}

// Returns the child accessibility element that represents the window’s
// proxy icon.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityproxy()

func (o NSAccessibilityProtocolObject) AccessibilityProxy() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityProxy"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s proxy
// icon.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityproxy(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityProxy(accessibilityProxy objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProxy:"), accessibilityProxy)
	}

// Returns the menu currently displaying for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityshownmenu()

func (o NSAccessibilityProtocolObject) AccessibilityShownMenu() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityShownMenu"))
	return objectivec.Object{ID: rv}
	}

// Sets the menu currently displaying for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityshownmenu(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityShownMenu(accessibilityShownMenu objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityShownMenu:"), accessibilityShownMenu)
	}

// Returns the child accessibility element that represents the window’s
// toolbar button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytoolbarbutton()

func (o NSAccessibilityProtocolObject) AccessibilityToolbarButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityToolbarButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s toolbar
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytoolbarbutton(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityToolbarButton(accessibilityToolbarButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityToolbarButton:"), accessibilityToolbarButton)
	}

// Returns the window that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitywindow()

func (o NSAccessibilityProtocolObject) AccessibilityWindow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindow"))
	return objectivec.Object{ID: rv}
	}

// Sets the window that contains the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitywindow(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityWindow(accessibilityWindow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindow:"), accessibilityWindow)
	}

// Returns the child accessibility element that represents the window’s zoom
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityzoombutton()

func (o NSAccessibilityProtocolObject) AccessibilityZoomButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityZoomButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the child accessibility element that represents the window’s zoom
// button.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityzoombutton(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityZoomButton(accessibilityZoomButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityZoomButton:"), accessibilityZoomButton)
	}

// Returns the icon for the app’s menu bar extra.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityextrasmenubar()

func (o NSAccessibilityProtocolObject) AccessibilityExtrasMenuBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityExtrasMenuBar"))
	return objectivec.Object{ID: rv}
	}

// Sets the icon for the app’s menu bar extra.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityextrasmenubar(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityExtrasMenuBar(accessibilityExtrasMenuBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExtrasMenuBar:"), accessibilityExtrasMenuBar)
	}

// Returns a Boolean value that determines whether the app is the frontmost
// app.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityfrontmost()

func (o NSAccessibilityProtocolObject) IsAccessibilityFrontmost() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFrontmost"))
	return rv
	}

// Sets a Boolean value that determines whether the app is the frontmost app.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfrontmost(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityFrontmost(accessibilityFrontmost bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrontmost:"), accessibilityFrontmost)
	}

// Returns a Boolean value that determines whether the app is in a hidden
// state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityhidden()

func (o NSAccessibilityProtocolObject) IsAccessibilityHidden() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityHidden"))
	return rv
	}

// Sets a Boolean value that determines whether the app is in a hidden state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhidden(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityHidden(accessibilityHidden bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHidden:"), accessibilityHidden)
	}

// Returns the app’s main window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymainwindow()

func (o NSAccessibilityProtocolObject) AccessibilityMainWindow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMainWindow"))
	return objectivec.Object{ID: rv}
	}

// Sets the app’s main window.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymainwindow(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityMainWindow(accessibilityMainWindow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMainWindow:"), accessibilityMainWindow)
	}

// Returns the app’s menu bar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymenubar()

func (o NSAccessibilityProtocolObject) AccessibilityMenuBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMenuBar"))
	return objectivec.Object{ID: rv}
	}

// Sets the app’s menu bar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymenubar(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityMenuBar(accessibilityMenuBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMenuBar:"), accessibilityMenuBar)
	}

// Returns an array that contains all the app’s windows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitywindows()

func (o NSAccessibilityProtocolObject) AccessibilityWindows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindows"))
	return objectivec.Object{ID: rv}
	}

// Sets the array that contains all the app’s windows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitywindows(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityWindows(accessibilityWindows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindows:"), accessibilityWindows)
	}

// Returns the number of columns in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumncount()

func (o NSAccessibilityProtocolObject) AccessibilityColumnCount() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityColumnCount"))
	return rv
	}

// Sets the number of columns in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumncount(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityColumnCount(accessibilityColumnCount int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnCount:"), accessibilityColumnCount)
	}

// Returns a Boolean value that determines whether the accessibility
// element’s grid is in row major order or in column major order.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityorderedbyrow()

func (o NSAccessibilityProtocolObject) IsAccessibilityOrderedByRow() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityOrderedByRow"))
	return rv
	}

// Sets a Boolean value that determines whether the element’s grid is in row
// major order or in column major order.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityorderedbyrow(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityOrderedByRow(accessibilityOrderedByRow bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrderedByRow:"), accessibilityOrderedByRow)
	}

// Returns the number of rows in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrowcount()

func (o NSAccessibilityProtocolObject) AccessibilityRowCount() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityRowCount"))
	return rv
	}

// Sets the number of rows in the accessibility element’s grid.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrowcount(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityRowCount(accessibilityRowCount int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowCount:"), accessibilityRowCount)
	}

// Returns the horizontal scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhorizontalscrollbar()

func (o NSAccessibilityProtocolObject) AccessibilityHorizontalScrollBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalScrollBar"))
	return objectivec.Object{ID: rv}
	}

// Sets the horizontal scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhorizontalscrollbar(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityHorizontalScrollBar(accessibilityHorizontalScrollBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalScrollBar:"), accessibilityHorizontalScrollBar)
	}

// Returns the vertical scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityverticalscrollbar()

func (o NSAccessibilityProtocolObject) AccessibilityVerticalScrollBar() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalScrollBar"))
	return objectivec.Object{ID: rv}
	}

// Sets the vertical scroll bar for the scroll view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityverticalscrollbar(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityVerticalScrollBar(accessibilityVerticalScrollBar objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalScrollBar:"), accessibilityVerticalScrollBar)
	}

// Returns the column header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumnheaderuielements()

func (o NSAccessibilityProtocolObject) AccessibilityColumnHeaderUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnHeaderUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the column header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumnheaderuielements(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityColumnHeaderUIElements(accessibilityColumnHeaderUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnHeaderUIElements:"), accessibilityColumnHeaderUIElements)
	}

// Returns the column accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumns()

func (o NSAccessibilityProtocolObject) AccessibilityColumns() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumns"))
	return objectivec.Object{ID: rv}
	}

// Sets the column accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumns(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityColumns(accessibilityColumns foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumns:"), accessibilityColumns)
	}

// Returns the column titles for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumntitles()

func (o NSAccessibilityProtocolObject) AccessibilityColumnTitles() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnTitles"))
	return objectivec.Object{ID: rv}
	}

// Sets the column titles for the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumntitles(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityColumnTitles(accessibilityColumnTitles foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnTitles:"), accessibilityColumnTitles)
	}

// Returns a Boolean value that determines whether the accessibility element
// is in an expanded state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityexpanded()

func (o NSAccessibilityProtocolObject) IsAccessibilityExpanded() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityExpanded"))
	return rv
	}

// Sets a Boolean value that determines whether accessibility element is in an
// expanded state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityexpanded(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityExpanded(accessibilityExpanded bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExpanded:"), accessibilityExpanded)
	}

// Returns the header for the table view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityheader()

func (o NSAccessibilityProtocolObject) AccessibilityHeader() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHeader"))
	return objectivec.Object{ID: rv}
	}

// Sets the header for the table view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityheader(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityHeader(accessibilityHeader objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHeader:"), accessibilityHeader)
	}

// Returns the index of the row or column that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityindex()

func (o NSAccessibilityProtocolObject) AccessibilityIndex() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityIndex"))
	return rv
	}

// Sets the index of the row or column that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityindex(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityIndex(accessibilityIndex int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIndex:"), accessibilityIndex)
	}

// Returns the row header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrowheaderuielements()

func (o NSAccessibilityProtocolObject) AccessibilityRowHeaderUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRowHeaderUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the row header accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrowheaderuielements(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityRowHeaderUIElements(accessibilityRowHeaderUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowHeaderUIElements:"), accessibilityRowHeaderUIElements)
	}

// Returns the row accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrows()

func (o NSAccessibilityProtocolObject) AccessibilityRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRows"))
	return objectivec.Object{ID: rv}
	}

// Sets the row accessibility elements for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrows(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityRows(accessibilityRows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRows:"), accessibilityRows)
	}

// Returns the currently selected columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedcolumns()

func (o NSAccessibilityProtocolObject) AccessibilitySelectedColumns() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedColumns"))
	return objectivec.Object{ID: rv}
	}

// Sets the currently selected columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedcolumns(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilitySelectedColumns(accessibilitySelectedColumns foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedColumns:"), accessibilitySelectedColumns)
	}

// Returns the currently selected rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedrows()

func (o NSAccessibilityProtocolObject) AccessibilitySelectedRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedRows"))
	return objectivec.Object{ID: rv}
	}

// Sets the currently selected rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedrows(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilitySelectedRows(accessibilitySelectedRows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedRows:"), accessibilitySelectedRows)
	}

// Returns the accessibility element’s sort direction.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysortdirection()

func (o NSAccessibilityProtocolObject) AccessibilitySortDirection() NSAccessibilitySortDirection {
	
	rv := objc.Send[NSAccessibilitySortDirection](o.ID, objc.Sel("accessibilitySortDirection"))
	return rv
	}

// Sets the accessibility element’s sort direction.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysortdirection(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilitySortDirection(accessibilitySortDirection NSAccessibilitySortDirection) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySortDirection:"), accessibilitySortDirection)
	}

// Returns the visible columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblecolumns()

func (o NSAccessibilityProtocolObject) AccessibilityVisibleColumns() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleColumns"))
	return objectivec.Object{ID: rv}
	}

// Sets the visible columns for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblecolumns(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityVisibleColumns(accessibilityVisibleColumns foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleColumns:"), accessibilityVisibleColumns)
	}

// Returns the visible rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblerows()

func (o NSAccessibilityProtocolObject) AccessibilityVisibleRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleRows"))
	return objectivec.Object{ID: rv}
	}

// Sets the visible rows for the table or outline.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblerows(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityVisibleRows(accessibilityVisibleRows foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleRows:"), accessibilityVisibleRows)
	}

// Returns a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilitydisclosed()

func (o NSAccessibilityProtocolObject) IsAccessibilityDisclosed() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityDisclosed"))
	return rv
	}

// Sets a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosed(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityDisclosed(accessibilityDisclosed bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosed:"), accessibilityDisclosed)
	}

// Returns the row disclosing the current row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydisclosedbyrow()

func (o NSAccessibilityProtocolObject) AccessibilityDisclosedByRow() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedByRow"))
	return objectivec.Object{ID: rv}
	}

// Sets the row disclosing the current row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosedbyrow(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityDisclosedByRow(accessibilityDisclosedByRow objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedByRow:"), accessibilityDisclosedByRow)
	}

// Returns the rows that the current row discloses.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydisclosedrows()

func (o NSAccessibilityProtocolObject) AccessibilityDisclosedRows() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedRows"))
	return objectivec.Object{ID: rv}
	}

// Sets the rows that the current row discloses.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosedrows(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityDisclosedRows(accessibilityDisclosedRows objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedRows:"), accessibilityDisclosedRows)
	}

// Returns the indention level for the row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydisclosurelevel()

func (o NSAccessibilityProtocolObject) AccessibilityDisclosureLevel() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityDisclosureLevel"))
	return rv
	}

// Sets the indention level for the row.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydisclosurelevel(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityDisclosureLevel(accessibilityDisclosureLevel int) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosureLevel:"), accessibilityDisclosureLevel)
	}

// Returns the column index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitycolumnindexrange()

func (o NSAccessibilityProtocolObject) AccessibilityColumnIndexRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityColumnIndexRange"))
	return rv
	}

// Sets the column index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitycolumnindexrange(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityColumnIndexRange(accessibilityColumnIndexRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnIndexRange:"), accessibilityColumnIndexRange)
	}

// Returns the row index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrowindexrange()

func (o NSAccessibilityProtocolObject) AccessibilityRowIndexRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRowIndexRange"))
	return rv
	}

// Sets the row index range of the cell.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrowindexrange(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityRowIndexRange(accessibilityRowIndexRange foundation.NSRange) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowIndexRange:"), accessibilityRowIndexRange)
	}

// Returns the currently selected cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityselectedcells()

func (o NSAccessibilityProtocolObject) AccessibilitySelectedCells() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedCells"))
	return objectivec.Object{ID: rv}
	}

// Sets the currently selected cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityselectedcells(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilitySelectedCells(accessibilitySelectedCells foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedCells:"), accessibilitySelectedCells)
	}

// Returns the visible cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityvisiblecells()

func (o NSAccessibilityProtocolObject) AccessibilityVisibleCells() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleCells"))
	return objectivec.Object{ID: rv}
	}

// Sets the visible cells for the table.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityvisiblecells(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityVisibleCells(accessibilityVisibleCells foundation.INSArray) {
	
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

func (o NSAccessibilityProtocolObject) AccessibilityCellForColumnRow(column int, row int) objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCellForColumn:row:"), column, row)
	return objectivec.Object{ID: rv}
	}

// Returns the drag handle elements for the layout item element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhandles()

func (o NSAccessibilityProtocolObject) AccessibilityHandles() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHandles"))
	return objectivec.Object{ID: rv}
	}

// Sets the drag handle accessibility elements for the layout item element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhandles(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityHandles(accessibilityHandles foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHandles:"), accessibilityHandles)
	}

// Returns the units that the layout area uses for horizontal values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhorizontalunits()

func (o NSAccessibilityProtocolObject) AccessibilityHorizontalUnits() NSAccessibilityUnits {
	
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityHorizontalUnits"))
	return rv
	}

// Sets the units that the layout area uses for horizontal values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhorizontalunits(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityHorizontalUnits(accessibilityHorizontalUnits NSAccessibilityUnits) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnits:"), accessibilityHorizontalUnits)
	}

// Returns the description of the layout area’s horizontal units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityhorizontalunitdescription()

func (o NSAccessibilityProtocolObject) AccessibilityHorizontalUnitDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the description of the layout area’s horizontal units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityhorizontalunitdescription(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityHorizontalUnitDescription(accessibilityHorizontalUnitDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalUnitDescription:"), objc.String(accessibilityHorizontalUnitDescription))
	}

// Returns the units that the layout area uses for vertical values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityverticalunits()

func (o NSAccessibilityProtocolObject) AccessibilityVerticalUnits() NSAccessibilityUnits {
	
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityVerticalUnits"))
	return rv
	}

// Sets the units that the layout area uses for vertical values.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityverticalunits(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityVerticalUnits(accessibilityVerticalUnits NSAccessibilityUnits) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnits:"), accessibilityVerticalUnits)
	}

// Returns the description of the layout area’s vertical units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityverticalunitdescription()

func (o NSAccessibilityProtocolObject) AccessibilityVerticalUnitDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the description of the layout area’s vertical units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityverticalunitdescription(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityVerticalUnitDescription(accessibilityVerticalUnitDescription string) {
	
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

// Returns the allowed values for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityallowedvalues()

func (o NSAccessibilityProtocolObject) AccessibilityAllowedValues() foundation.NSNumber {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAllowedValues"))
	return foundation.NSNumberFromID(rv)
	}

// Sets the allowed values for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityallowedvalues(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityAllowedValues(accessibilityAllowedValues foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAllowedValues:"), accessibilityAllowedValues)
	}

// Returns the child label elements for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylabeluielements()

func (o NSAccessibilityProtocolObject) AccessibilityLabelUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabelUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the child label elements for the slider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylabeluielements(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityLabelUIElements(accessibilityLabelUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelUIElements:"), accessibilityLabelUIElements)
	}

// Returns the value of the label accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylabelvalue()

func (o NSAccessibilityProtocolObject) AccessibilityLabelValue() float64 {
	
	rv := objc.Send[float64](o.ID, objc.Sel("accessibilityLabelValue"))
	return rv
	}

// Sets the value of the label accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylabelvalue(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityLabelValue(accessibilityLabelValue float64) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelValue:"), accessibilityLabelValue)
	}

// Returns the contents that follow the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitynextcontents()

func (o NSAccessibilityProtocolObject) AccessibilityNextContents() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityNextContents"))
	return objectivec.Object{ID: rv}
	}

// Sets the contents that follow the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitynextcontents(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityNextContents(accessibilityNextContents foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNextContents:"), accessibilityNextContents)
	}

// Returns the contents that precede the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitypreviouscontents()

func (o NSAccessibilityProtocolObject) AccessibilityPreviousContents() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPreviousContents"))
	return objectivec.Object{ID: rv}
	}

// Sets the contents that precede the divider accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitypreviouscontents(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityPreviousContents(accessibilityPreviousContents foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPreviousContents:"), accessibilityPreviousContents)
	}

// Returns an array that contains the views and splitter bar from the split
// view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysplitters()

func (o NSAccessibilityProtocolObject) AccessibilitySplitters() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySplitters"))
	return objectivec.Object{ID: rv}
	}

// Sets the array that contains the views and splitter bar from the split
// view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysplitters(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilitySplitters(accessibilitySplitters foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySplitters:"), accessibilitySplitters)
	}

// Returns the overflow button for the toolbar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityoverflowbutton()

func (o NSAccessibilityProtocolObject) AccessibilityOverflowButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityOverflowButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the overflow button for the toolbar.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityoverflowbutton(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityOverflowButton(accessibilityOverflowButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOverflowButton:"), accessibilityOverflowButton)
	}

// Returns the tab accessibility elements for the tab view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytabs()

func (o NSAccessibilityProtocolObject) AccessibilityTabs() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTabs"))
	return objectivec.Object{ID: rv}
	}

// Sets the tab accessibility elements for the tab view.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytabs(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityTabs(accessibilityTabs foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTabs:"), accessibilityTabs)
	}

// Returns the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkergroupuielement()

func (o NSAccessibilityProtocolObject) AccessibilityMarkerGroupUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerGroupUIElement"))
	return objectivec.Object{ID: rv}
	}

// Sets the user interface element that functions as a marker group for the
// ruler accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkergroupuielement(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityMarkerGroupUIElement(accessibilityMarkerGroupUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerGroupUIElement:"), accessibilityMarkerGroupUIElement)
	}

// Returns the human-readable description of the marker type.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkertypedescription()

func (o NSAccessibilityProtocolObject) AccessibilityMarkerTypeDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerTypeDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the human-readable description of the marker type.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkertypedescription(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityMarkerTypeDescription(accessibilityMarkerTypeDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerTypeDescription:"), objc.String(accessibilityMarkerTypeDescription))
	}

// Returns the array of marker accessibility elements for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkeruielements()

func (o NSAccessibilityProtocolObject) AccessibilityMarkerUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the array of marker accessibility elements for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkeruielements(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityMarkerUIElements(accessibilityMarkerUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerUIElements:"), accessibilityMarkerUIElements)
	}

// Returns the marker values for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitymarkervalues()

func (o NSAccessibilityProtocolObject) AccessibilityMarkerValues() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerValues"))
	return objectivec.Object{ID: rv}
	}

// Sets the marker values for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitymarkervalues(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityMarkerValues(accessibilityMarkerValues objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerValues:"), accessibilityMarkerValues)
	}

// Returns the type of markers for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityrulermarkertype()

func (o NSAccessibilityProtocolObject) AccessibilityRulerMarkerType() NSAccessibilityRulerMarkerType {
	
	rv := objc.Send[NSAccessibilityRulerMarkerType](o.ID, objc.Sel("accessibilityRulerMarkerType"))
	return rv
	}

// Sets the type of markers for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityrulermarkertype(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityRulerMarkerType(accessibilityRulerMarkerType NSAccessibilityRulerMarkerType) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRulerMarkerType:"), accessibilityRulerMarkerType)
	}

// Returns the units for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityunits()

func (o NSAccessibilityProtocolObject) AccessibilityUnits() NSAccessibilityUnits {
	
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityUnits"))
	return rv
	}

// Sets the units used for the ruler.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityunits(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityUnits(accessibilityUnits NSAccessibilityUnits) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnits:"), accessibilityUnits)
	}

// Returns the human-readable description of the ruler’s units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityunitdescription()

func (o NSAccessibilityProtocolObject) AccessibilityUnitDescription() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUnitDescription"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the human-readable description of the ruler’s units.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityunitdescription(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityUnitDescription(accessibilityUnitDescription string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnitDescription:"), objc.String(accessibilityUnitDescription))
	}

// Returns the URL for the file that the accessibility element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydocument()

func (o NSAccessibilityProtocolObject) AccessibilityDocument() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDocument"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the URL for the file that the accessibility element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydocument(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityDocument(accessibilityDocument string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDocument:"), objc.String(accessibilityDocument))
	}

// Returns a Boolean value that indicates whether the accessibility element is
// in an edited state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/isaccessibilityedited()

func (o NSAccessibilityProtocolObject) IsAccessibilityEdited() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEdited"))
	return rv
	}

// Sets a Boolean value that indicates whether the accessibility element is in
// an edited state.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityedited(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityEdited(accessibilityEdited bool) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityEdited:"), accessibilityEdited)
	}

// Returns the filename for the file that the accessibility element
// represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityfilename()

func (o NSAccessibilityProtocolObject) AccessibilityFilename() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFilename"))
	return foundation.NSStringFromID(rv).String()
	}

// Sets the filename for the file that the accessibility element represents.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityfilename(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityFilename(accessibilityFilename string) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFilename:"), objc.String(accessibilityFilename))
	}

// Returns the elements that have links with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitylinkeduielements()

func (o NSAccessibilityProtocolObject) AccessibilityLinkedUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLinkedUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the elements that have links with the accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitylinkeduielements(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityLinkedUIElements(accessibilityLinkedUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLinkedUIElements:"), accessibilityLinkedUIElements)
	}

// Returns the list of elements that the accessibility element is a title for.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityservesastitleforuielements()

func (o NSAccessibilityProtocolObject) AccessibilityServesAsTitleForUIElements() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityServesAsTitleForUIElements"))
	return objectivec.Object{ID: rv}
	}

// Sets the list of elements that the accessibility element is a title for.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityservesastitleforuielements(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityServesAsTitleForUIElements(accessibilityServesAsTitleForUIElements foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityServesAsTitleForUIElements:"), accessibilityServesAsTitleForUIElements)
	}

// Returns the static text element that represents the accessibility
// element’s title.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitytitleuielement()

func (o NSAccessibilityProtocolObject) AccessibilityTitleUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitleUIElement"))
	return objectivec.Object{ID: rv}
	}

// Sets the static text element that represents the accessibility element’s
// title.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitytitleuielement(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityTitleUIElement(accessibilityTitleUIElement objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitleUIElement:"), accessibilityTitleUIElement)
	}

// Returns the clear button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityclearbutton()

func (o NSAccessibilityProtocolObject) AccessibilityClearButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityClearButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the clear button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityclearbutton(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityClearButton(accessibilityClearButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityClearButton:"), accessibilityClearButton)
	}

// Returns the search button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysearchbutton()

func (o NSAccessibilityProtocolObject) AccessibilitySearchButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the search button for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysearchbutton(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilitySearchButton(accessibilitySearchButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchButton:"), accessibilitySearchButton)
	}

// Returns the search menu for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitysearchmenu()

func (o NSAccessibilityProtocolObject) AccessibilitySearchMenu() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchMenu"))
	return objectivec.Object{ID: rv}
	}

// Sets the search menu for the search field.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitysearchmenu(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilitySearchMenu(accessibilitySearchMenu objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchMenu:"), accessibilitySearchMenu)
	}

// Cancels the current operation.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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

// Returns the increment button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityincrementbutton()

func (o NSAccessibilityProtocolObject) AccessibilityIncrementButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIncrementButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the increment button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityincrementbutton(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityIncrementButton(accessibilityIncrementButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIncrementButton:"), accessibilityIncrementButton)
	}

// Returns the decrement button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/accessibilitydecrementbutton()

func (o NSAccessibilityProtocolObject) AccessibilityDecrementButton() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDecrementButton"))
	return objectivec.Object{ID: rv}
	}

// Sets the decrement button for the stepper accessibility element.
//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilitydecrementbutton(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityDecrementButton(accessibilityDecrementButton objectivec.IObject) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDecrementButton:"), accessibilityDecrementButton)
	}

// Increments the accessibility element’s value.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that have an adjustable [accessibilityValue]
// property.
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformIncrement()

func (o NSAccessibilityProtocolObject) AccessibilityPerformIncrement() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformIncrement"))
	return rv
	}

// Decrements the accessibility element’s value.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Use this method on elements that have an adjustable [accessibilityValue]
// property.
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/accessibilityPerformDecrement()

func (o NSAccessibilityProtocolObject) AccessibilityPerformDecrement() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformDecrement"))
	return rv
	}

// Deletes the accessibility element’s value.
//
// # Return Value
// 
// [true] if the action was successfully triggered; otherwise, [false]. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
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

// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityattributeduserinputlabels()

func (o NSAccessibilityProtocolObject) AccessibilityAttributedUserInputLabels() foundation.NSAttributedString {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAttributedUserInputLabels"))
	return foundation.NSAttributedStringFromID(rv)
	}

// See: /documentation/appkit/nsaccessibilityprotocol/accessibilityuserinputlabels()

func (o NSAccessibilityProtocolObject) AccessibilityUserInputLabels() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUserInputLabels"))
	return foundation.NSStringFromID(rv).String()
	}

//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityattributeduserinputlabels(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityAttributedUserInputLabels(accessibilityAttributedUserInputLabels foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAttributedUserInputLabels:"), accessibilityAttributedUserInputLabels)
	}

//
// See: /documentation/appkit/nsaccessibilityprotocol/setaccessibilityuserinputlabels(_:)

func (o NSAccessibilityProtocolObject) SetAccessibilityUserInputLabels(accessibilityUserInputLabels foundation.INSArray) {
	
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUserInputLabels:"), accessibilityUserInputLabels)
	}

