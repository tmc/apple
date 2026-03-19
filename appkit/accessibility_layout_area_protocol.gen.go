// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a layout area.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityLayoutArea
type NSAccessibilityLayoutArea interface {
	objectivec.IObject
	NSAccessibilityElementProtocol
	NSAccessibilityGroup

	// Returns the accessibility element’s children in the accessibility hierarchy.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityLayoutArea/accessibilityChildren()
	AccessibilityChildren() foundation.INSArray

	// Returns a short description of the layout area.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityLayoutArea/accessibilityLabel()
	AccessibilityLabel() string

	// Returns the layout area’s currently selected children.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityLayoutArea/accessibilitySelectedChildren()
	AccessibilitySelectedChildren() foundation.INSArray
}

// NSAccessibilityLayoutAreaObject wraps an existing Objective-C object that conforms to the NSAccessibilityLayoutArea protocol.
type NSAccessibilityLayoutAreaObject struct {
	objectivec.Object
}
func (o NSAccessibilityLayoutAreaObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSAccessibilityLayoutAreaObjectFromID constructs a [NSAccessibilityLayoutAreaObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSAccessibilityLayoutAreaObjectFromID(id objc.ID) NSAccessibilityLayoutAreaObject {
	return NSAccessibilityLayoutAreaObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns the accessibility element’s children in the accessibility
// hierarchy.
//
// # Return Value
// 
// An array that contains references to this element’s children in the
// accessibility hierarchy.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityChildren] property.
//
// [accessibilityChildren]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityChildren
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityLayoutArea/accessibilityChildren()
func (o NSAccessibilityLayoutAreaObject) AccessibilityChildren() foundation.INSArray {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityChildren"))
	return foundation.NSArrayFromID(rv)
	}
// The child accessibility element with the current focus.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityLayoutArea/accessibilityFocusedUIElement
func (o NSAccessibilityLayoutAreaObject) AccessibilityFocusedUIElement() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityFocusedUIElement"))
	return objectivec.Object{ID: rv}
	}
// Returns a short description of the layout area.
//
// # Return Value
// 
// The description of the layout area.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityLabel] property.
// 
// Do not include the control’s type in the label (for example, use
// [Canvas], not `Canvas Layout Area`). If possible use a single word. To help
// ensure that accessibility clients such as VoiceOver read the label with the
// correct intonation, start this label with a capital letter. Do not put a
// period at the end. Always localize the label.
//
// [accessibilityLabel]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLabel
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityLayoutArea/accessibilityLabel()
func (o NSAccessibilityLayoutAreaObject) AccessibilityLabel() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
	}
// Returns the layout area’s currently selected children.
//
// # Return Value
// 
// An array containing the currently selected children. If no children are
// selected, this method returns an empty array.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilitySelectedChildren] property.
//
// [accessibilitySelectedChildren]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilitySelectedChildren
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityLayoutArea/accessibilitySelectedChildren()
func (o NSAccessibilityLayoutAreaObject) AccessibilitySelectedChildren() foundation.INSArray {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilitySelectedChildren"))
	return foundation.NSArrayFromID(rv)
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
func (o NSAccessibilityLayoutAreaObject) AccessibilityFrame() corefoundation.CGRect {
	
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
func (o NSAccessibilityLayoutAreaObject) AccessibilityParent() objectivec.IObject {
	
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
func (o NSAccessibilityLayoutAreaObject) AccessibilityIdentifier() string {
	
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
func (o NSAccessibilityLayoutAreaObject) IsAccessibilityFocused() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}

