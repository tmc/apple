// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as static text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStaticText
type NSAccessibilityStaticText interface {
	objectivec.IObject
	NSAccessibilityElementProtocol

	// Returns the text that the accessibility element displays.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStaticText/accessibilityValue()
	//
	// SDK narrowing: NSAccessibilityStaticText declares this as - (nullable NSString *)accessibilityValue.
	// Go has no covariance, so the base declaration on NSAccessibility (rendered on INSView) is used.
	// Convert with foundation.NSStringFromID(v.GetID()) to recover the narrowed type.
	AccessibilityValue() objectivec.IObject
}

// NSAccessibilityStaticTextObject wraps an existing Objective-C object that conforms to the NSAccessibilityStaticText protocol.
type NSAccessibilityStaticTextObject struct {
	objectivec.Object
}

func (o NSAccessibilityStaticTextObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSAccessibilityStaticTextObjectFromID constructs a [NSAccessibilityStaticTextObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSAccessibilityStaticTextObjectFromID(id objc.ID) NSAccessibilityStaticTextObject {
	return NSAccessibilityStaticTextObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the text that the accessibility element displays.
//
// # Return Value
//
// The text displayed by the element.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilityValue] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStaticText/accessibilityValue()
//
// SDK narrowing: NSAccessibilityStaticText declares this as - (nullable NSString *)accessibilityValue.
// Go has no covariance, so the base declaration on NSAccessibility (rendered on INSView) is used.
// Convert with foundation.NSStringFromID(v.GetID()) to recover the narrowed type.
func (o NSAccessibilityStaticTextObject) AccessibilityValue() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValue"))
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
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStaticText/accessibilityAttributedString(for:)
func (o NSAccessibilityStaticTextObject) AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityAttributedStringForRange:"), range_)
	return foundation.NSAttributedStringFromID(rv)
}

// Returns the range of visible characters in the document.
//
// # Return Value
//
// The range of the visible characters in the document. This method should
// return the range for entire lines. Characters that are horizontally clipped
// are included in this range.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [NSWindow.AccessibilityVisibleCharacterRange] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStaticText/accessibilityVisibleCharacterRange()
func (o NSAccessibilityStaticTextObject) AccessibilityVisibleCharacterRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityVisibleCharacterRange"))
	return rv
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
func (o NSAccessibilityStaticTextObject) AccessibilityFrame() corefoundation.CGRect {
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
// [NSWindow.AccessibilityParent] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityParent()
func (o NSAccessibilityStaticTextObject) AccessibilityParent() objectivec.IObject {
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
// [NSWindow.AccessibilityIdentifier] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityIdentifier()
func (o NSAccessibilityStaticTextObject) AccessibilityIdentifier() string {
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
// [NSWindow.AccessibilityFocused] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/isAccessibilityFocused()
func (o NSAccessibilityStaticTextObject) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
}
