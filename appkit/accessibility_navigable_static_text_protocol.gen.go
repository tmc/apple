// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as navigable static text.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityNavigableStaticText
type NSAccessibilityNavigableStaticText interface {
	objectivec.IObject
	NSAccessibilityStaticText

	// Returns the rectangle that encloses the specified range of characters.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityNavigableStaticText/accessibilityFrame(for:)
	AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect

	// Returns the line number for the line that contains the specified character index.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityNavigableStaticText/accessibilityLine(for:)
	AccessibilityLineForIndex(index int) int

	// Returns the range of characters in the specified line.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityNavigableStaticText/accessibilityRange(forLine:)
	AccessibilityRangeForLine(lineNumber int) foundation.NSRange

	// Returns the substring for the specified range.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityNavigableStaticText/accessibilityString(for:)
	AccessibilityStringForRange(range_ foundation.NSRange) string
}



// NSAccessibilityNavigableStaticTextObject wraps an existing Objective-C object that conforms to the NSAccessibilityNavigableStaticText protocol.
type NSAccessibilityNavigableStaticTextObject struct {
	objectivec.Object
}
func (o NSAccessibilityNavigableStaticTextObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSAccessibilityNavigableStaticTextObjectFromID constructs a [NSAccessibilityNavigableStaticTextObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSAccessibilityNavigableStaticTextObjectFromID(id objc.ID) NSAccessibilityNavigableStaticTextObject {
	return NSAccessibilityNavigableStaticTextObject{
		Object: objectivec.ObjectFromID(id),
	}
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
// If the range crosses a line boundary, the returned rectangle will fully
// enclose all the lines of characters.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityNavigableStaticText/accessibilityFrame(for:)

func (o NSAccessibilityNavigableStaticTextObject) AccessibilityFrameForRange(range_ foundation.NSRange) corefoundation.CGRect {
	
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
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityNavigableStaticText/accessibilityLine(for:)

func (o NSAccessibilityNavigableStaticTextObject) AccessibilityLineForIndex(index int) int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityLineForIndex:"), index)
	return rv
	}

// Returns the range of characters in the specified line.
//
// lineNumber: The line number to be examined.
//
// # Return Value
// 
// The range of characters for the specified line number. If the line ends
// with a newline character, including the newline is preferred.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityNavigableStaticText/accessibilityRange(forLine:)

func (o NSAccessibilityNavigableStaticTextObject) AccessibilityRangeForLine(lineNumber int) foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityRangeForLine:"), lineNumber)
	return rv
	}

// Returns the substring for the specified range.
//
// range: A range of characters contained by this element.
//
// # Return Value
// 
// The substring specified by the given range.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityNavigableStaticText/accessibilityString(for:)

func (o NSAccessibilityNavigableStaticTextObject) AccessibilityStringForRange(range_ foundation.NSRange) string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityStringForRange:"), range_)
	return foundation.NSStringFromID(rv).String()
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
// [accessibilityFrame] property. This method is called whenever accessibility
// clients request the [size] or [position] attributes.
//
// [accessibilityFrame]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
// [position]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
// [size]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityFrame()

func (o NSAccessibilityNavigableStaticTextObject) AccessibilityFrame() corefoundation.CGRect {
	
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
// [accessibilityParent]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityParent
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityParent()

func (o NSAccessibilityNavigableStaticTextObject) AccessibilityParent() objectivec.IObject {
	
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
// [accessibilityIdentifier]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIdentifier
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityIdentifier()

func (o NSAccessibilityNavigableStaticTextObject) AccessibilityIdentifier() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityIdentifier"))
	return foundation.NSStringFromID(rv).String()
	}

// Returns a Boolean value that indicates whether the accessibility element
// has the keyboard focus.
//
// # Return Value
// 
// [true] if this element has the keyboard focus; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityFocused] property.
//
// [accessibilityFocused]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFocused
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/isAccessibilityFocused()

func (o NSAccessibilityNavigableStaticTextObject) IsAccessibilityFocused() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
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
// [accessibilityValue] property.
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStaticText/accessibilityValue()

func (o NSAccessibilityNavigableStaticTextObject) AccessibilityValue() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValue"))
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
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStaticText/accessibilityAttributedString(for:)

func (o NSAccessibilityNavigableStaticTextObject) AccessibilityAttributedStringForRange(range_ foundation.NSRange) foundation.NSAttributedString {
	
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
// [accessibilityVisibleCharacterRange] property.
//
// [accessibilityVisibleCharacterRange]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityVisibleCharacterRange
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStaticText/accessibilityVisibleCharacterRange()

func (o NSAccessibilityNavigableStaticTextObject) AccessibilityVisibleCharacterRange() foundation.NSRange {
	
	rv := objc.Send[foundation.NSRange](o.ID, objc.Sel("accessibilityVisibleCharacterRange"))
	return rv
	}







