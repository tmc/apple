// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a stepper.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStepper
type NSAccessibilityStepper interface {
	objectivec.IObject

	// Returns a short description of the stepper.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStepper/accessibilityLabel()
	AccessibilityLabel() string

	// Decrements the stepper’s value.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStepper/accessibilityPerformDecrement()
	AccessibilityPerformDecrement() bool

	// Increments the stepper’s value.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStepper/accessibilityPerformIncrement()
	AccessibilityPerformIncrement() bool
}



// NSAccessibilityStepperObject wraps an existing Objective-C object that conforms to the NSAccessibilityStepper protocol.
type NSAccessibilityStepperObject struct {
	objectivec.Object
}
func (o NSAccessibilityStepperObject) BaseObject() objectivec.Object {
	return o.Object
}



// NSAccessibilityStepperObjectFromID constructs a [NSAccessibilityStepperObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSAccessibilityStepperObjectFromID(id objc.ID) NSAccessibilityStepperObject {
	return NSAccessibilityStepperObject{
		Object: objectivec.ObjectFromID(id),
	}
}




// Returns a short description of the stepper.
//
// # Return Value
// 
// The description of the stepper.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityLabel] property.
// 
// Do not include the control’s type in the label (for example, use
// [Volume], not `Volume stepper`). If possible use a single word. To help
// ensure that accessibility clients such as VoiceOver read the label with the
// correct intonation, start this label with a capital letter. Do not put a
// period at the end. Always localize the label.
//
// [accessibilityLabel]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLabel
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStepper/accessibilityLabel()

func (o NSAccessibilityStepperObject) AccessibilityLabel() string {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
	}

// Decrements the stepper’s value.
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
// This method must post an [valueChanged] notification after changing the
// stepper’s value.
//
// [valueChanged]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/valueChanged
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStepper/accessibilityPerformDecrement()

func (o NSAccessibilityStepperObject) AccessibilityPerformDecrement() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformDecrement"))
	return rv
	}

// Increments the stepper’s value.
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
// This method must post an [valueChanged] notification after changing the
// stepper’s value.
//
// [valueChanged]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/valueChanged
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStepper/accessibilityPerformIncrement()

func (o NSAccessibilityStepperObject) AccessibilityPerformIncrement() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformIncrement"))
	return rv
	}

// Returns the stepper’s value.
//
// # Return Value
// 
// The value for the stepper.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityValue] property.
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityStepper/accessibilityValue()

func (o NSAccessibilityStepperObject) AccessibilityValue() objectivec.IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityValue"))
	return objectivec.Object{ID: rv}
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

func (o NSAccessibilityStepperObject) AccessibilityFrame() corefoundation.CGRect {
	
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

func (o NSAccessibilityStepperObject) AccessibilityParent() objectivec.IObject {
	
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

func (o NSAccessibilityStepperObject) AccessibilityIdentifier() string {
	
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

func (o NSAccessibilityStepperObject) IsAccessibilityFocused() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}







