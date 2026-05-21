// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSliderAccessory] class.
var (
	_NSSliderAccessoryClass     NSSliderAccessoryClass
	_NSSliderAccessoryClassOnce sync.Once
)

func getNSSliderAccessoryClass() NSSliderAccessoryClass {
	_NSSliderAccessoryClassOnce.Do(func() {
		_NSSliderAccessoryClass = NSSliderAccessoryClass{class: objc.GetClass("NSSliderAccessory")}
	})
	return _NSSliderAccessoryClass
}

// GetNSSliderAccessoryClass returns the class object for NSSliderAccessory.
func GetNSSliderAccessoryClass() NSSliderAccessoryClass {
	return getNSSliderAccessoryClass()
}

type NSSliderAccessoryClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSliderAccessoryClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSliderAccessoryClass) Alloc() NSSliderAccessory {
	rv := objc.Send[NSSliderAccessory](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// # Instance Properties
//
//   - [NSSliderAccessory.Behavior]: The effect on interaction with the accessory.
//   - [NSSliderAccessory.SetBehavior]
//   - [NSSliderAccessory.IsEnabled]
//   - [NSSliderAccessory.SetEnabled]
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessory
type NSSliderAccessory struct {
	objectivec.Object
}

// NSSliderAccessoryFromID constructs a [NSSliderAccessory] from an objc.ID.
func NSSliderAccessoryFromID(id objc.ID) NSSliderAccessory {
	return NSSliderAccessory{objectivec.Object{ID: id}}
}

// NOTE: NSSliderAccessory adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSliderAccessory] class.
//
// # Instance Properties
//
//   - [INSSliderAccessory.Behavior]: The effect on interaction with the accessory.
//   - [INSSliderAccessory.SetBehavior]
//   - [INSSliderAccessory.IsEnabled]
//   - [INSSliderAccessory.SetEnabled]
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessory
type INSSliderAccessory interface {
	objectivec.IObject

	// Topic: Instance Properties

	// The effect on interaction with the accessory.
	Behavior() INSSliderAccessoryBehavior
	SetBehavior(value INSSliderAccessoryBehavior)
	IsEnabled() bool
	SetEnabled(value bool)

	InitWithCoder(coder foundation.INSCoder) NSSliderAccessory
	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s NSSliderAccessory) Init() NSSliderAccessory {
	rv := objc.Send[NSSliderAccessory](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSliderAccessory) Autorelease() NSSliderAccessory {
	rv := objc.Send[NSSliderAccessory](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSliderAccessory creates a new NSSliderAccessory instance.
func NewNSSliderAccessory() NSSliderAccessory {
	class := getNSSliderAccessoryClass()
	rv := objc.Send[NSSliderAccessory](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/init(coder:)
func NewSliderAccessoryWithCoder(coder foundation.INSCoder) NSSliderAccessory {
	instance := getNSSliderAccessoryClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSSliderAccessoryFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/init(image:)
func NewSliderAccessoryWithImage(image INSImage) NSSliderAccessory {
	rv := objc.Send[objc.ID](objc.ID(getNSSliderAccessoryClass().class), objc.Sel("accessoryWithImage:"), image)
	return NSSliderAccessoryFromID(rv)
}

// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/init(coder:)
func (s NSSliderAccessory) InitWithCoder(coder foundation.INSCoder) NSSliderAccessory {
	rv := objc.Send[NSSliderAccessory](s.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (s NSSliderAccessory) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The effect on interaction with the accessory.
//
// # Discussion
//
// The default value is `automaticBehavior`.
//
// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/behavior
func (s NSSliderAccessory) Behavior() INSSliderAccessoryBehavior {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("behavior"))
	return NSSliderAccessoryBehaviorFromID(objc.ID(rv))
}
func (s NSSliderAccessory) SetBehavior(value INSSliderAccessoryBehavior) {
	objc.Send[struct{}](s.ID, objc.Sel("setBehavior:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSSliderAccessory/isEnabled
func (s NSSliderAccessory) IsEnabled() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isEnabled"))
	return rv
}
func (s NSSliderAccessory) SetEnabled(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setEnabled:"), value)
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
func (o NSSliderAccessory) AccessibilityFrame() corefoundation.CGRect {
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
func (o NSSliderAccessory) AccessibilityParent() objectivec.IObject {
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
func (o NSSliderAccessory) AccessibilityIdentifier() string {
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
func (o NSSliderAccessory) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
}

// Protocol methods for NSAccessibilityProtocol

// Returns a Boolean value that determines whether the accessibility element
// participates in the accessibility hierarchy.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityElement()
func (o NSSliderAccessory) IsAccessibilityElement() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityElement"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// responds to user events.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEnabled()
func (o NSSliderAccessory) IsAccessibilityEnabled() bool {
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
func (o NSSliderAccessory) IsAccessibilitySelectorAllowed(selector objc.SEL) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelectorAllowed:"), selector)
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// contains protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityProtectedContent()
func (o NSSliderAccessory) IsAccessibilityProtectedContent() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityProtectedContent"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// is currently in a selected state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilitySelected()
func (o NSSliderAccessory) IsAccessibilitySelected() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelected"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// must have content for successful submission of a form.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityRequired()
func (o NSSliderAccessory) IsAccessibilityRequired() bool {
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
func (o NSSliderAccessory) AccessibilityStringForRange(range_ foundation.NSRange) string {
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
func (o NSSliderAccessory) AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString {
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
func (o NSSliderAccessory) AccessibilityRTFForRange(range_ foundation.NSRange) foundation.NSData {
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
func (o NSSliderAccessory) AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect {
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
func (o NSSliderAccessory) AccessibilityLineForIndex(index int) int {
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
func (o NSSliderAccessory) AccessibilityRangeForIndex(index int) foundation.NSRange {
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
func (o NSSliderAccessory) AccessibilityStyleRangeForIndex(index int) foundation.NSRange {
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
func (o NSSliderAccessory) AccessibilityRangeForLine(line int) foundation.NSRange {
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
func (o NSSliderAccessory) AccessibilityRangeForPosition(point corefoundation.CGPoint) foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForPosition:"), point)
	return rv
}

// Returns the Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityAlternateUIVisible()
func (o NSSliderAccessory) IsAccessibilityAlternateUIVisible() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
	return rv
}

// Returns a Boolean value that determines whether the window is the app’s
// main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMain()
func (o NSSliderAccessory) IsAccessibilityMain() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMain"))
	return rv
}

// Returns the Boolean value that determines whether the window is in a
// minimized state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityMinimized()
func (o NSSliderAccessory) IsAccessibilityMinimized() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMinimized"))
	return rv
}

// Returns a Boolean value that determines whether the window is modal.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityModal()
func (o NSSliderAccessory) IsAccessibilityModal() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityModal"))
	return rv
}

// Returns a Boolean value that determines whether the app is the frontmost
// app.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityFrontmost()
func (o NSSliderAccessory) IsAccessibilityFrontmost() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFrontmost"))
	return rv
}

// Returns a Boolean value that determines whether the app is in a hidden
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityHidden()
func (o NSSliderAccessory) IsAccessibilityHidden() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityHidden"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility
// element’s grid is in row major order or in column major order.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityOrderedByRow()
func (o NSSliderAccessory) IsAccessibilityOrderedByRow() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityOrderedByRow"))
	return rv
}

// Returns a Boolean value that determines whether the accessibility element
// is in an expanded state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityExpanded()
func (o NSSliderAccessory) IsAccessibilityExpanded() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityExpanded"))
	return rv
}

// Returns a Boolean value that determines whether the row is disclosing other
// rows.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityDisclosed()
func (o NSSliderAccessory) IsAccessibilityDisclosed() bool {
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
func (o NSSliderAccessory) AccessibilityCellForColumnRow(column int, row int) objectivec.IObject {
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
func (o NSSliderAccessory) AccessibilityLayoutPointForScreenPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
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
func (o NSSliderAccessory) AccessibilityLayoutSizeForScreenSize(size corefoundation.CGSize) corefoundation.CGSize {
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
func (o NSSliderAccessory) AccessibilityScreenPointForLayoutPoint(point corefoundation.CGPoint) corefoundation.CGPoint {
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
func (o NSSliderAccessory) AccessibilityScreenSizeForLayoutSize(size corefoundation.CGSize) corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](o.ID, objc.Sel("accessibilityScreenSizeForLayoutSize:"), size)
	return rv
}

// Returns a Boolean value that indicates whether the accessibility element is
// in an edited state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityProtocol/isAccessibilityEdited()
func (o NSSliderAccessory) IsAccessibilityEdited() bool {
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
func (o NSSliderAccessory) AccessibilityPerformCancel() bool {
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
func (o NSSliderAccessory) AccessibilityPerformConfirm() bool {
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
func (o NSSliderAccessory) AccessibilityPerformPick() bool {
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
func (o NSSliderAccessory) AccessibilityPerformPress() bool {
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
func (o NSSliderAccessory) AccessibilityPerformShowAlternateUI() bool {
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
func (o NSSliderAccessory) AccessibilityPerformShowDefaultUI() bool {
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
func (o NSSliderAccessory) AccessibilityPerformShowMenu() bool {
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
func (o NSSliderAccessory) AccessibilityPerformRaise() bool {
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
func (o NSSliderAccessory) AccessibilityPerformIncrement() bool {
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
func (o NSSliderAccessory) AccessibilityPerformDecrement() bool {
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
func (o NSSliderAccessory) AccessibilityPerformDelete() bool {
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
func (o NSSliderAccessory) AccessibilityActivationPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](o.ID, objc.Sel("accessibilityActivationPoint"))
	return corefoundation.CGPoint(rv)
}

func (o NSSliderAccessory) SetAccessibilityActivationPoint(value corefoundation.CGPoint) {
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
func (o NSSliderAccessory) AccessibilityAllowedValues() []foundation.NSNumber {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityAllowedValues"))
	result := make([]foundation.NSNumber, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = foundation.NSNumberFromID(id)
	}
	return result
}

func (o NSSliderAccessory) SetAccessibilityAllowedValues(value []foundation.NSNumber) {
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
func (o NSSliderAccessory) AccessibilityAlternateUIVisible() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
	return bool(rv)
}

func (o NSSliderAccessory) SetAccessibilityAlternateUIVisible(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAlternateUIVisible:"), value)
}

// The child accessibility element with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityApplicationFocusedUIElement
func (o NSSliderAccessory) AccessibilityApplicationFocusedUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityApplicationFocusedUIElement"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityApplicationFocusedUIElement(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityApplicationFocusedUIElement:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityAttributedUserInputLabels
func (o NSSliderAccessory) AccessibilityAttributedUserInputLabels() []foundation.NSAttributedString {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityAttributedUserInputLabels"))
	result := make([]foundation.NSAttributedString, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = foundation.NSAttributedStringFromID(id)
	}
	return result
}

func (o NSSliderAccessory) SetAccessibilityAttributedUserInputLabels(value []foundation.NSAttributedString) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityAttributedUserInputLabels:"), objectivec.IObjectSliceToNSArray(value))
}

// The child accessibility element that represents the window’s cancel
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCancelButton
func (o NSSliderAccessory) AccessibilityCancelButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCancelButton"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityCancelButton(value objectivec.IObject) {
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
func (o NSSliderAccessory) AccessibilityChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityChildren"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityChildren(value foundation.INSArray) {
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
func (o NSSliderAccessory) AccessibilityChildrenInNavigationOrder() []objectivec.IObject {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityChildrenInNavigationOrder"))
	result := make([]objectivec.IObject, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = objectivec.Object{ID: id}
	}
	return result
}

func (o NSSliderAccessory) SetAccessibilityChildrenInNavigationOrder(value []objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityChildrenInNavigationOrder:"), objectivec.IObjectSliceToNSArray(value))
}

// The clear button for the search field.
//
// # Discussion
//
// Use this property on a search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityClearButton
func (o NSSliderAccessory) AccessibilityClearButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityClearButton"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityClearButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityClearButton:"), value)
}

// The child accessibility element that represents the window’s close
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCloseButton
func (o NSSliderAccessory) AccessibilityCloseButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCloseButton"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityCloseButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCloseButton:"), value)
}

// The number of columns in the accessibility element’s grid.
//
// # Discussion
//
// Use this property for UI elements that present a grid of child elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumnCount
func (o NSSliderAccessory) AccessibilityColumnCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityColumnCount"))
	return int(rv)
}

func (o NSSliderAccessory) SetAccessibilityColumnCount(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnCount:"), value)
}

// The column header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumnHeaderUIElements
func (o NSSliderAccessory) AccessibilityColumnHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityColumnHeaderUIElements(value foundation.INSArray) {
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
func (o NSSliderAccessory) AccessibilityColumnIndexRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityColumnIndexRange"))
	return foundation.NSRange(rv)
}

func (o NSSliderAccessory) SetAccessibilityColumnIndexRange(value foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnIndexRange:"), value)
}

// The column titles for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumnTitles
func (o NSSliderAccessory) AccessibilityColumnTitles() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumnTitles"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityColumnTitles(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityColumnTitles:"), value)
}

// The column accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityColumns
func (o NSSliderAccessory) AccessibilityColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityColumns"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityColumns(value foundation.INSArray) {
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
func (o NSSliderAccessory) AccessibilityContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityContents"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityContents(value foundation.INSArray) {
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
func (o NSSliderAccessory) AccessibilityCriticalValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityCriticalValue"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityCriticalValue(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCriticalValue:"), value)
}

// The custom actions of the current accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityCustomActions
func (o NSSliderAccessory) AccessibilityCustomActions() []NSAccessibilityCustomAction {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityCustomActions"))
	result := make([]NSAccessibilityCustomAction, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = NSAccessibilityCustomActionFromID(id)
	}
	return result
}

func (o NSSliderAccessory) SetAccessibilityCustomActions(value []NSAccessibilityCustomAction) {
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
func (o NSSliderAccessory) AccessibilityCustomRotors() []NSAccessibilityCustomRotor {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityCustomRotors"))
	result := make([]NSAccessibilityCustomRotor, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = NSAccessibilityCustomRotorFromID(id)
	}
	return result
}

func (o NSSliderAccessory) SetAccessibilityCustomRotors(value []NSAccessibilityCustomRotor) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityCustomRotors:"), objectivec.IObjectSliceToNSArray(value))
}

// The decrement button for the stepper accessibility element.
//
// # Discussion
//
// Use this property on a stepper.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDecrementButton
func (o NSSliderAccessory) AccessibilityDecrementButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDecrementButton"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityDecrementButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDecrementButton:"), value)
}

// The child accessibility element that represents the window’s default
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDefaultButton
func (o NSSliderAccessory) AccessibilityDefaultButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDefaultButton"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityDefaultButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDefaultButton:"), value)
}

// A Boolean value that determines whether the row is disclosing other rows.
//
// # Discussion
//
// Use this property in the elements representing an outline’s row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosed
func (o NSSliderAccessory) AccessibilityDisclosed() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityDisclosed"))
	return bool(rv)
}

func (o NSSliderAccessory) SetAccessibilityDisclosed(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosed:"), value)
}

// The row disclosing the current row.
//
// # Discussion
//
// Use this property in the elements representing an outline’s row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosedByRow
func (o NSSliderAccessory) AccessibilityDisclosedByRow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedByRow"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityDisclosedByRow(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedByRow:"), value)
}

// The rows that the current row discloses.
//
// # Discussion
//
// Use this property in the elements representing an outline’s row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosedRows
func (o NSSliderAccessory) AccessibilityDisclosedRows() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDisclosedRows"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityDisclosedRows(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosedRows:"), value)
}

// The indention level for the row.
//
// # Discussion
//
// Use this property in the elements representing an outline’s row.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosureLevel
func (o NSSliderAccessory) AccessibilityDisclosureLevel() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityDisclosureLevel"))
	return int(rv)
}

func (o NSSliderAccessory) SetAccessibilityDisclosureLevel(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDisclosureLevel:"), value)
}

// The URL for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDocument
func (o NSSliderAccessory) AccessibilityDocument() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityDocument"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSSliderAccessory) SetAccessibilityDocument(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityDocument:"), objc.String(value))
}

// A Boolean value that indicates whether the accessibility element is in an
// edited state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityEdited
func (o NSSliderAccessory) AccessibilityEdited() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEdited"))
	return bool(rv)
}

func (o NSSliderAccessory) SetAccessibilityEdited(value bool) {
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
func (o NSSliderAccessory) AccessibilityElement() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityElement"))
	return bool(rv)
}

func (o NSSliderAccessory) SetAccessibilityElement(value bool) {
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
func (o NSSliderAccessory) AccessibilityEnabled() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityEnabled"))
	return bool(rv)
}

func (o NSSliderAccessory) SetAccessibilityEnabled(value bool) {
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
func (o NSSliderAccessory) AccessibilityExpanded() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityExpanded"))
	return bool(rv)
}

func (o NSSliderAccessory) SetAccessibilityExpanded(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExpanded:"), value)
}

// The icon for the app’s menu bar extra.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityExtrasMenuBar
func (o NSSliderAccessory) AccessibilityExtrasMenuBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityExtrasMenuBar"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityExtrasMenuBar(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityExtrasMenuBar:"), value)
}

// The filename for the file that the accessibility element represents.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFilename
func (o NSSliderAccessory) AccessibilityFilename() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFilename"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSSliderAccessory) SetAccessibilityFilename(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFilename:"), objc.String(value))
}

// A Boolean value that determines whether the accessibility element has the
// keyboard focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocused
func (o NSSliderAccessory) AccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return bool(rv)
}

func (o NSSliderAccessory) SetAccessibilityFocused(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFocused:"), value)
}

// The child window with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocusedWindow
func (o NSSliderAccessory) AccessibilityFocusedWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFocusedWindow"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityFocusedWindow(value objectivec.IObject) {
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
func (o NSSliderAccessory) SetAccessibilityFrame(value corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrame:"), value)
}

// A Boolean value that determines whether the app is the frontmost app.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrontmost
func (o NSSliderAccessory) AccessibilityFrontmost() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFrontmost"))
	return bool(rv)
}

func (o NSSliderAccessory) SetAccessibilityFrontmost(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrontmost:"), value)
}

// The child accessibility element that represents the window’s full-screen
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFullScreenButton
func (o NSSliderAccessory) AccessibilityFullScreenButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFullScreenButton"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityFullScreenButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFullScreenButton:"), value)
}

// The child accessibility element that represents the window’s grow area.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityGrowArea
func (o NSSliderAccessory) AccessibilityGrowArea() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityGrowArea"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityGrowArea(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityGrowArea:"), value)
}

// The drag handle accessibility elements for the layout item element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHandles
func (o NSSliderAccessory) AccessibilityHandles() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHandles"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityHandles(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHandles:"), value)
}

// The header for the table view.
//
// # Discussion
//
// Use this property on a table view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHeader
func (o NSSliderAccessory) AccessibilityHeader() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHeader"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityHeader(value objectivec.IObject) {
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
func (o NSSliderAccessory) AccessibilityHelp() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHelp"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSSliderAccessory) SetAccessibilityHelp(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHelp:"), objc.String(value))
}

// A Boolean value that determines whether the app is in a hidden state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHidden
func (o NSSliderAccessory) AccessibilityHidden() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityHidden"))
	return bool(rv)
}

func (o NSSliderAccessory) SetAccessibilityHidden(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHidden:"), value)
}

// The horizontal scroll bar for the scroll view.
//
// # Discussion
//
// Use this property on a scrollable view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHorizontalScrollBar
func (o NSSliderAccessory) AccessibilityHorizontalScrollBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalScrollBar"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityHorizontalScrollBar(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityHorizontalScrollBar:"), value)
}

// A description of the layout area’s horizontal units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityHorizontalUnitDescription
func (o NSSliderAccessory) AccessibilityHorizontalUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityHorizontalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSSliderAccessory) SetAccessibilityHorizontalUnitDescription(value string) {
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
func (o NSSliderAccessory) AccessibilityHorizontalUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityHorizontalUnits"))
	return NSAccessibilityUnits(rv)
}

func (o NSSliderAccessory) SetAccessibilityHorizontalUnits(value NSAccessibilityUnits) {
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
func (o NSSliderAccessory) SetAccessibilityIdentifier(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIdentifier:"), objc.String(value))
}

// The increment button for the stepper accessibility element.
//
// # Discussion
//
// Use this property on a stepper.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIncrementButton
func (o NSSliderAccessory) AccessibilityIncrementButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIncrementButton"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityIncrementButton(value objectivec.IObject) {
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
func (o NSSliderAccessory) AccessibilityIndex() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityIndex"))
	return int(rv)
}

func (o NSSliderAccessory) SetAccessibilityIndex(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityIndex:"), value)
}

// The line number that contains the insertion point.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityInsertionPointLineNumber
func (o NSSliderAccessory) AccessibilityInsertionPointLineNumber() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityInsertionPointLineNumber"))
	return int(rv)
}

func (o NSSliderAccessory) SetAccessibilityInsertionPointLineNumber(value int) {
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
func (o NSSliderAccessory) AccessibilityLabel() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSSliderAccessory) SetAccessibilityLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabel:"), objc.String(value))
}

// The child label elements for the slider accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLabelUIElements
func (o NSSliderAccessory) AccessibilityLabelUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabelUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityLabelUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLabelUIElements:"), value)
}

// The value of the label accessibility element.
//
// # Discussion
//
// Use this property on a slider element’s labels.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLabelValue
func (o NSSliderAccessory) AccessibilityLabelValue() float32 {
	rv := objc.Send[float32](o.ID, objc.Sel("accessibilityLabelValue"))
	return float32(rv)
}

func (o NSSliderAccessory) SetAccessibilityLabelValue(value float32) {
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
func (o NSSliderAccessory) AccessibilityLinkedUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLinkedUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityLinkedUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityLinkedUIElements:"), value)
}

// A Boolean value that determines whether the window is the app’s main
// window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMain
func (o NSSliderAccessory) AccessibilityMain() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMain"))
	return bool(rv)
}

func (o NSSliderAccessory) SetAccessibilityMain(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMain:"), value)
}

// The app’s main window.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMainWindow
func (o NSSliderAccessory) AccessibilityMainWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMainWindow"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityMainWindow(value objectivec.IObject) {
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
func (o NSSliderAccessory) AccessibilityMarkerGroupUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerGroupUIElement"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityMarkerGroupUIElement(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerGroupUIElement:"), value)
}

// A human-readable description of the marker type.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMarkerTypeDescription
func (o NSSliderAccessory) AccessibilityMarkerTypeDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerTypeDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSSliderAccessory) SetAccessibilityMarkerTypeDescription(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerTypeDescription:"), objc.String(value))
}

// An array of marker accessibility elements for the ruler.
//
// # Discussion
//
// Use this property on a ruler element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMarkerUIElements
func (o NSSliderAccessory) AccessibilityMarkerUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityMarkerUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMarkerUIElements:"), value)
}

// The marker values for the ruler.
//
// # Discussion
//
// Use this property on a ruler element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMarkerValues
func (o NSSliderAccessory) AccessibilityMarkerValues() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMarkerValues"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityMarkerValues(value objectivec.IObject) {
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
func (o NSSliderAccessory) AccessibilityMaxValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMaxValue"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityMaxValue(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMaxValue:"), value)
}

// The app’s menu bar.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMenuBar
func (o NSSliderAccessory) AccessibilityMenuBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMenuBar"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityMenuBar(value objectivec.IObject) {
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
func (o NSSliderAccessory) AccessibilityMinValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinValue"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityMinValue(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinValue:"), value)
}

// The child accessibility element that represents the window’s minimize
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMinimizeButton
func (o NSSliderAccessory) AccessibilityMinimizeButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityMinimizeButton"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityMinimizeButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimizeButton:"), value)
}

// A Boolean value that determines whether this window is in a minimized
// state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityMinimized
func (o NSSliderAccessory) AccessibilityMinimized() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityMinimized"))
	return bool(rv)
}

func (o NSSliderAccessory) SetAccessibilityMinimized(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityMinimized:"), value)
}

// A Boolean value that determines whether the window is modal.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityModal
func (o NSSliderAccessory) AccessibilityModal() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityModal"))
	return bool(rv)
}

func (o NSSliderAccessory) SetAccessibilityModal(value bool) {
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
func (o NSSliderAccessory) AccessibilityNextContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityNextContents"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityNextContents(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityNextContents:"), value)
}

// The number of characters in the text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityNumberOfCharacters
func (o NSSliderAccessory) AccessibilityNumberOfCharacters() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityNumberOfCharacters"))
	return int(rv)
}

func (o NSSliderAccessory) SetAccessibilityNumberOfCharacters(value int) {
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
func (o NSSliderAccessory) AccessibilityOrderedByRow() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityOrderedByRow"))
	return bool(rv)
}

func (o NSSliderAccessory) SetAccessibilityOrderedByRow(value bool) {
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
func (o NSSliderAccessory) AccessibilityOrientation() NSAccessibilityOrientation {
	rv := objc.Send[NSAccessibilityOrientation](o.ID, objc.Sel("accessibilityOrientation"))
	return NSAccessibilityOrientation(rv)
}

func (o NSSliderAccessory) SetAccessibilityOrientation(value NSAccessibilityOrientation) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityOrientation:"), value)
}

// The overflow button for the toolbar.
//
// # Discussion
//
// Use this property on a toolbar element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityOverflowButton
func (o NSSliderAccessory) AccessibilityOverflowButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityOverflowButton"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityOverflowButton(value objectivec.IObject) {
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
func (o NSSliderAccessory) SetAccessibilityParent(value objectivec.IObject) {
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
func (o NSSliderAccessory) AccessibilityPlaceholderValue() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPlaceholderValue"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSSliderAccessory) SetAccessibilityPlaceholderValue(value string) {
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
func (o NSSliderAccessory) AccessibilityPreviousContents() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityPreviousContents"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityPreviousContents(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityPreviousContents:"), value)
}

// A Boolean value that determines whether the accessibility element contains
// protected content.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityProtectedContent
func (o NSSliderAccessory) AccessibilityProtectedContent() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityProtectedContent"))
	return bool(rv)
}

func (o NSSliderAccessory) SetAccessibilityProtectedContent(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityProtectedContent:"), value)
}

// The child accessibility element that represents the window’s proxy icon.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityProxy
func (o NSSliderAccessory) AccessibilityProxy() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityProxy"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityProxy(value objectivec.IObject) {
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
func (o NSSliderAccessory) AccessibilityRequired() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityRequired"))
	return bool(rv)
}

func (o NSSliderAccessory) SetAccessibilityRequired(value bool) {
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
func (o NSSliderAccessory) AccessibilityRole() NSAccessibilityRole {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRole"))
	return NSAccessibilityRole(foundation.NSStringFromID(rv).String())
}

func (o NSSliderAccessory) SetAccessibilityRole(value NSAccessibilityRole) {
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
func (o NSSliderAccessory) AccessibilityRoleDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRoleDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSSliderAccessory) SetAccessibilityRoleDescription(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRoleDescription:"), objc.String(value))
}

// The number of rows in the accessibility element’s grid.
//
// # Discussion
//
// Use this property for elements that present a grid of child elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRowCount
func (o NSSliderAccessory) AccessibilityRowCount() int {
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityRowCount"))
	return int(rv)
}

func (o NSSliderAccessory) SetAccessibilityRowCount(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowCount:"), value)
}

// The row header accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRowHeaderUIElements
func (o NSSliderAccessory) AccessibilityRowHeaderUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRowHeaderUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityRowHeaderUIElements(value foundation.INSArray) {
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
func (o NSSliderAccessory) AccessibilityRowIndexRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRowIndexRange"))
	return foundation.NSRange(rv)
}

func (o NSSliderAccessory) SetAccessibilityRowIndexRange(value foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRowIndexRange:"), value)
}

// The row accessibility elements for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityRows
func (o NSSliderAccessory) AccessibilityRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityRows"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityRows(value foundation.INSArray) {
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
func (o NSSliderAccessory) AccessibilityRulerMarkerType() NSAccessibilityRulerMarkerType {
	rv := objc.Send[NSAccessibilityRulerMarkerType](o.ID, objc.Sel("accessibilityRulerMarkerType"))
	return NSAccessibilityRulerMarkerType(rv)
}

func (o NSSliderAccessory) SetAccessibilityRulerMarkerType(value NSAccessibilityRulerMarkerType) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityRulerMarkerType:"), value)
}

// The search button for the search field.
//
// # Discussion
//
// Use this property on a search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySearchButton
func (o NSSliderAccessory) AccessibilitySearchButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchButton"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilitySearchButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchButton:"), value)
}

// The search menu for the search field.
//
// # Discussion
//
// Use this property on a search field.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySearchMenu
func (o NSSliderAccessory) AccessibilitySearchMenu() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySearchMenu"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilitySearchMenu(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySearchMenu:"), value)
}

// A Boolean value that determines whether the accessibility element is
// currently in a selected state.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelected
func (o NSSliderAccessory) AccessibilitySelected() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilitySelected"))
	return bool(rv)
}

func (o NSSliderAccessory) SetAccessibilitySelected(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelected:"), value)
}

// The currently selected cells for the table.
//
// # Discussion
//
// This property is required for all elements that act like cell-based tables.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedCells
func (o NSSliderAccessory) AccessibilitySelectedCells() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedCells"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilitySelectedCells(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedCells:"), value)
}

// The accessibility element’s currently selected children.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedChildren
func (o NSSliderAccessory) AccessibilitySelectedChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedChildren"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilitySelectedChildren(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedChildren:"), value)
}

// The currently selected columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedColumns
func (o NSSliderAccessory) AccessibilitySelectedColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedColumns"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilitySelectedColumns(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedColumns:"), value)
}

// The currently selected rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedRows
func (o NSSliderAccessory) AccessibilitySelectedRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedRows"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilitySelectedRows(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedRows:"), value)
}

// The currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedText
func (o NSSliderAccessory) AccessibilitySelectedText() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedText"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSSliderAccessory) SetAccessibilitySelectedText(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedText:"), objc.String(value))
}

// The range of the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedTextRange
func (o NSSliderAccessory) AccessibilitySelectedTextRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySelectedTextRange"))
	return foundation.NSRange(rv)
}

func (o NSSliderAccessory) SetAccessibilitySelectedTextRange(value foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySelectedTextRange:"), value)
}

// An array of ranges for the currently selected text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedTextRanges
func (o NSSliderAccessory) AccessibilitySelectedTextRanges() []foundation.NSValue {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilitySelectedTextRanges"))
	result := make([]foundation.NSValue, len(rvIDs))
	for i, id := range rvIDs {
		result[i] = foundation.NSValueFromID(id)
	}
	return result
}

func (o NSSliderAccessory) SetAccessibilitySelectedTextRanges(value []foundation.NSValue) {
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
func (o NSSliderAccessory) AccessibilityServesAsTitleForUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityServesAsTitleForUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityServesAsTitleForUIElements(value foundation.INSArray) {
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
func (o NSSliderAccessory) AccessibilitySharedCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilitySharedCharacterRange"))
	return foundation.NSRange(rv)
}

func (o NSSliderAccessory) SetAccessibilitySharedCharacterRange(value foundation.NSRange) {
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
func (o NSSliderAccessory) AccessibilitySharedFocusElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedFocusElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilitySharedFocusElements(value foundation.INSArray) {
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
func (o NSSliderAccessory) AccessibilitySharedTextUIElements() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySharedTextUIElements"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilitySharedTextUIElements(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySharedTextUIElements:"), value)
}

// The menu currently displaying for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityShownMenu
func (o NSSliderAccessory) AccessibilityShownMenu() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityShownMenu"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityShownMenu(value objectivec.IObject) {
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
func (o NSSliderAccessory) AccessibilitySortDirection() NSAccessibilitySortDirection {
	rv := objc.Send[NSAccessibilitySortDirection](o.ID, objc.Sel("accessibilitySortDirection"))
	return NSAccessibilitySortDirection(rv)
}

func (o NSSliderAccessory) SetAccessibilitySortDirection(value NSAccessibilitySortDirection) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySortDirection:"), value)
}

// An array that contains the views and splitter bar from the split view.
//
// # Discussion
//
// Use this property on a split view element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySplitters
func (o NSSliderAccessory) AccessibilitySplitters() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySplitters"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilitySplitters(value foundation.INSArray) {
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
func (o NSSliderAccessory) AccessibilitySubrole() NSAccessibilitySubrole {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySubrole"))
	return NSAccessibilitySubrole(foundation.NSStringFromID(rv).String())
}

func (o NSSliderAccessory) SetAccessibilitySubrole(value NSAccessibilitySubrole) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilitySubrole:"), objc.String(string(value)))
}

// The tab accessibility elements for the tab view.
//
// # Discussion
//
// Use this property on a tab view element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityTabs
func (o NSSliderAccessory) AccessibilityTabs() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTabs"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityTabs(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTabs:"), value)
}

// The title of the accessibility element—for example, a button’s visible
// text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityTitle
func (o NSSliderAccessory) AccessibilityTitle() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitle"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSSliderAccessory) SetAccessibilityTitle(value string) {
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
func (o NSSliderAccessory) AccessibilityTitleUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTitleUIElement"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityTitleUIElement(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTitleUIElement:"), value)
}

// The child accessibility element that represents the window’s toolbar
// button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityToolbarButton
func (o NSSliderAccessory) AccessibilityToolbarButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityToolbarButton"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityToolbarButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityToolbarButton:"), value)
}

// The top-level element that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityTopLevelUIElement
func (o NSSliderAccessory) AccessibilityTopLevelUIElement() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityTopLevelUIElement"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityTopLevelUIElement(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityTopLevelUIElement:"), value)
}

// The URL for the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityURL
func (o NSSliderAccessory) AccessibilityURL() foundation.NSURL {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityURL"))
	return foundation.NSURLFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityURL(value foundation.NSURL) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityURL:"), value)
}

// A human-readable description of the ruler’s units.
//
// # Discussion
//
// Use this property on a ruler element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityUnitDescription
func (o NSSliderAccessory) AccessibilityUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSSliderAccessory) SetAccessibilityUnitDescription(value string) {
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
func (o NSSliderAccessory) AccessibilityUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityUnits"))
	return NSAccessibilityUnits(rv)
}

func (o NSSliderAccessory) SetAccessibilityUnits(value NSAccessibilityUnits) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityUnits:"), value)
}

// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityUserInputLabels
func (o NSSliderAccessory) AccessibilityUserInputLabels() []string {
	rvIDs := objc.Send[[]objc.ID](o.ID, objc.Sel("accessibilityUserInputLabels"))
	return objc.ConvertSliceToStrings(rvIDs)
}

func (o NSSliderAccessory) SetAccessibilityUserInputLabels(value []string) {
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
func (o NSSliderAccessory) AccessibilityValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValue"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityValue(value objectivec.IObject) {
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
func (o NSSliderAccessory) AccessibilityValueDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValueDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSSliderAccessory) SetAccessibilityValueDescription(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityValueDescription:"), objc.String(value))
}

// The vertical scroll bar for the scroll view.
//
// # Discussion
//
// Use this property on a scrollable view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVerticalScrollBar
func (o NSSliderAccessory) AccessibilityVerticalScrollBar() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalScrollBar"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityVerticalScrollBar(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalScrollBar:"), value)
}

// A description of the layout area’s vertical units.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVerticalUnitDescription
func (o NSSliderAccessory) AccessibilityVerticalUnitDescription() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVerticalUnitDescription"))
	return foundation.NSStringFromID(rv).String()
}

func (o NSSliderAccessory) SetAccessibilityVerticalUnitDescription(value string) {
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
func (o NSSliderAccessory) AccessibilityVerticalUnits() NSAccessibilityUnits {
	rv := objc.Send[NSAccessibilityUnits](o.ID, objc.Sel("accessibilityVerticalUnits"))
	return NSAccessibilityUnits(rv)
}

func (o NSSliderAccessory) SetAccessibilityVerticalUnits(value NSAccessibilityUnits) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVerticalUnits:"), value)
}

// The visible cells for the table.
//
// # Discussion
//
// This property is required for all elements that act like cell-based tables.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleCells
func (o NSSliderAccessory) AccessibilityVisibleCells() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleCells"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityVisibleCells(value foundation.INSArray) {
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
func (o NSSliderAccessory) AccessibilityVisibleCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityVisibleCharacterRange"))
	return foundation.NSRange(rv)
}

func (o NSSliderAccessory) SetAccessibilityVisibleCharacterRange(value foundation.NSRange) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleCharacterRange:"), value)
}

// The accessibility element’s visible child accessibility elements.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleChildren
func (o NSSliderAccessory) AccessibilityVisibleChildren() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleChildren"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityVisibleChildren(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleChildren:"), value)
}

// The visible columns for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleColumns
func (o NSSliderAccessory) AccessibilityVisibleColumns() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleColumns"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityVisibleColumns(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityVisibleColumns:"), value)
}

// The visible rows for the table or outline.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleRows
func (o NSSliderAccessory) AccessibilityVisibleRows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityVisibleRows"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityVisibleRows(value foundation.INSArray) {
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
func (o NSSliderAccessory) AccessibilityWarningValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWarningValue"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityWarningValue(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWarningValue:"), value)
}

// The window that contains the accessibility element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityWindow
func (o NSSliderAccessory) AccessibilityWindow() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindow"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityWindow(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindow:"), value)
}

// An array that contains all the app’s windows.
//
// # Discussion
//
// Use on the app element.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityWindows
func (o NSSliderAccessory) AccessibilityWindows() foundation.INSArray {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityWindows"))
	return foundation.NSArrayFromID(rv)
}

func (o NSSliderAccessory) SetAccessibilityWindows(value foundation.INSArray) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityWindows:"), value)
}

// The child accessibility element that represents the window’s zoom button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityZoomButton
func (o NSSliderAccessory) AccessibilityZoomButton() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityZoomButton"))
	return objectivec.Object{ID: rv}
}

func (o NSSliderAccessory) SetAccessibilityZoomButton(value objectivec.IObject) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityZoomButton:"), value)
}
