// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a layout item.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityLayoutItem
type NSAccessibilityLayoutItem interface {
	objectivec.IObject
	NSAccessibilityElementProtocol
	NSAccessibilityGroup
}

// NSAccessibilityLayoutItemObject wraps an existing Objective-C object that conforms to the NSAccessibilityLayoutItem protocol.
type NSAccessibilityLayoutItemObject struct {
	objectivec.Object
}

func (o NSAccessibilityLayoutItemObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSAccessibilityLayoutItemObjectFromID constructs a [NSAccessibilityLayoutItemObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSAccessibilityLayoutItemObjectFromID(id objc.ID) NSAccessibilityLayoutItemObject {
	return NSAccessibilityLayoutItemObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Sets the accessibility element’s frame.
//
// frame: The new frame in screen coordinates.
//
// # Discussion
//
// This method is the setter method for the [NSAccessibilityProtocol]
// protocol’s [accessibilityFrame] property. By implementing this method,
// you allow accessibility clients to modify this element’s frame.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityLayoutItem/setAccessibilityFrame(_:)
//
// [accessibilityFrame]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
func (o NSAccessibilityLayoutItemObject) SetAccessibilityFrame(frame corefoundation.CGRect) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessibilityFrame:"), frame)
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
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityFrame()
//
// [accessibilityFrame]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
// [position]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
// [size]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
func (o NSAccessibilityLayoutItemObject) AccessibilityFrame() corefoundation.CGRect {
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
func (o NSAccessibilityLayoutItemObject) AccessibilityParent() objectivec.IObject {
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
func (o NSAccessibilityLayoutItemObject) AccessibilityIdentifier() string {
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
func (o NSAccessibilityLayoutItemObject) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
}
