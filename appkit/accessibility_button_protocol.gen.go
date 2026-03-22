// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a button.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityButton
type NSAccessibilityButton interface {
	objectivec.IObject
	NSAccessibilityElementProtocol

	// Returns a short description of the button.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityButton/accessibilityLabel()
	AccessibilityLabel() string

	// Simulates clicking the button.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityButton/accessibilityPerformPress()
	AccessibilityPerformPress() bool
}

// NSAccessibilityButtonObject wraps an existing Objective-C object that conforms to the NSAccessibilityButton protocol.
type NSAccessibilityButtonObject struct {
	objectivec.Object
}
func (o NSAccessibilityButtonObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSAccessibilityButtonObjectFromID constructs a [NSAccessibilityButtonObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSAccessibilityButtonObjectFromID(id objc.ID) NSAccessibilityButtonObject {
	return NSAccessibilityButtonObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns a short description of the button.
//
// # Return Value
// 
// The description of this button.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityLabel] property.
// 
// Do not include the accessibility element’s type in the label (for
// example, write [Play] not `Play button`.). If possible, use a single word.
// To help ensure that accessibility clients like VoiceOver read the label
// with the correct intonation, this label should start with a capital letter.
// Do not put a period at the end. Always localize the label.
//
// [accessibilityLabel]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLabel
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityButton/accessibilityLabel()
func (o NSAccessibilityButtonObject) AccessibilityLabel() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
	}
// Simulates clicking the button.
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
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityButton/accessibilityPerformPress()
func (o NSAccessibilityButtonObject) AccessibilityPerformPress() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformPress"))
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
func (o NSAccessibilityButtonObject) AccessibilityFrame() corefoundation.CGRect {
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
func (o NSAccessibilityButtonObject) AccessibilityParent() objectivec.IObject {
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
func (o NSAccessibilityButtonObject) AccessibilityIdentifier() string {
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
func (o NSAccessibilityButtonObject) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}

