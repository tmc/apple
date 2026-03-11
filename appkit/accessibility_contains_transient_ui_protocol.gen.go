// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to support dynamic UI changes.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityContainsTransientUI
type NSAccessibilityContainsTransientUI interface {
	objectivec.IObject
	NSAccessibilityElementProtocol

	// Displays the accessibility element’s alternative UI.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityContainsTransientUI/accessibilityPerformShowAlternateUI()
	AccessibilityPerformShowAlternateUI() bool

	// Returns to the accessibility element’s original UI.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityContainsTransientUI/accessibilityPerformShowDefaultUI()
	AccessibilityPerformShowDefaultUI() bool

	// Returns a Boolean value that determines whether the accessibility element’s alternative UI is currently visible.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityContainsTransientUI/isAccessibilityAlternateUIVisible()
	IsAccessibilityAlternateUIVisible() bool
}



// NSAccessibilityContainsTransientUIObject wraps an existing Objective-C object that conforms to the NSAccessibilityContainsTransientUI protocol.
type NSAccessibilityContainsTransientUIObject struct {
	objectivec.Object
}
func (o NSAccessibilityContainsTransientUIObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSAccessibilityContainsTransientUIObjectFromID constructs a [NSAccessibilityContainsTransientUIObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSAccessibilityContainsTransientUIObjectFromID(id objc.ID) NSAccessibilityContainsTransientUIObject {
	return NSAccessibilityContainsTransientUIObject{
		Object: objectivec.ObjectFromID(id),
	}
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
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityContainsTransientUI/accessibilityPerformShowAlternateUI()

func (o NSAccessibilityContainsTransientUIObject) AccessibilityPerformShowAlternateUI() bool {
	
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
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityContainsTransientUI/accessibilityPerformShowDefaultUI()

func (o NSAccessibilityContainsTransientUIObject) AccessibilityPerformShowDefaultUI() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformShowDefaultUI"))
	return rv
	}

// Returns a Boolean value that determines whether the accessibility
// element’s alternative UI is currently visible.
//
// # Return Value
// 
// Use this property for elements that present an alternate UI—for example,
// when the pointer hovers over an interface element for a few seconds.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityContainsTransientUI/isAccessibilityAlternateUIVisible()

func (o NSAccessibilityContainsTransientUIObject) IsAccessibilityAlternateUIVisible() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityAlternateUIVisible"))
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

func (o NSAccessibilityContainsTransientUIObject) AccessibilityFrame() corefoundation.CGRect {
	
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

func (o NSAccessibilityContainsTransientUIObject) AccessibilityParent() objectivec.IObject {
	
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

func (o NSAccessibilityContainsTransientUIObject) AccessibilityIdentifier() string {
	
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

func (o NSAccessibilityContainsTransientUIObject) IsAccessibilityFocused() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}







