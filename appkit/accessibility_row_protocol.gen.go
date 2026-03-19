// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a row for a table, list, or outline view.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityRow
type NSAccessibilityRow interface {
	objectivec.IObject
	NSAccessibilityElementProtocol
	NSAccessibilityGroup

	// Returns the index for the row.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityRow/accessibilityIndex()
	AccessibilityIndex() int
}

// NSAccessibilityRowObject wraps an existing Objective-C object that conforms to the NSAccessibilityRow protocol.
type NSAccessibilityRowObject struct {
	objectivec.Object
}
func (o NSAccessibilityRowObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSAccessibilityRowObjectFromID constructs a [NSAccessibilityRowObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSAccessibilityRowObjectFromID(id objc.ID) NSAccessibilityRowObject {
	return NSAccessibilityRowObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the index for the row.
//
// # Return Value
// 
// The index for the row.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityIndex] property.
//
// [accessibilityIndex]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityIndex
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityRow/accessibilityIndex()
func (o NSAccessibilityRowObject) AccessibilityIndex() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityIndex"))
	return rv
	}
// Returns the indention level for the row.
//
// # Return Value
// 
// The disclosure level for the row.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityDisclosureLevel] property.
//
// [accessibilityDisclosureLevel]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityDisclosureLevel
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityRow/accessibilityDisclosureLevel()
func (o NSAccessibilityRowObject) AccessibilityDisclosureLevel() int {
	
	rv := objc.Send[int](o.ID, objc.Sel("accessibilityDisclosureLevel"))
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
// [accessibilityFrame] property. This method is called whenever accessibility
// clients request the [size] or [position] attributes.
//
// [accessibilityFrame]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
// [position]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
// [size]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityFrame()
func (o NSAccessibilityRowObject) AccessibilityFrame() corefoundation.CGRect {
	
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
func (o NSAccessibilityRowObject) AccessibilityParent() objectivec.IObject {
	
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
func (o NSAccessibilityRowObject) AccessibilityIdentifier() string {
	
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
func (o NSAccessibilityRowObject) IsAccessibilityFocused() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}

