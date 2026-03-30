// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A role-based protocol that declares the minimum interface necessary for an accessibility element to act as a slider.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySlider
type NSAccessibilitySlider interface {
	objectivec.IObject
	NSAccessibilityElementProtocol

	// Returns a short description of the slider.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySlider/accessibilityLabel()
	AccessibilityLabel() string

	// Decrements the slider’s value.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySlider/accessibilityPerformDecrement()
	AccessibilityPerformDecrement() bool

	// Increments the slider’s value.
	//
	// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySlider/accessibilityPerformIncrement()
	AccessibilityPerformIncrement() bool
}

// NSAccessibilitySliderObject wraps an existing Objective-C object that conforms to the NSAccessibilitySlider protocol.
type NSAccessibilitySliderObject struct {
	objectivec.Object
}

func (o NSAccessibilitySliderObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSAccessibilitySliderObjectFromID constructs a [NSAccessibilitySliderObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSAccessibilitySliderObjectFromID(id objc.ID) NSAccessibilitySliderObject {
	return NSAccessibilitySliderObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Returns a short description of the slider.
//
// # Return Value
//
// The description of the slider.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityLabel] property.
//
// Do not include the control’s type in the label (for example, use
// [Volume], not `Volume slider`). If possible use a single word. To help
// ensure that accessibility clients such as VoiceOver read the label with the
// correct intonation, start this label with a capital letter. Do not put a
// period at the end. Always localize the label.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySlider/accessibilityLabel()
//
// [accessibilityLabel]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityLabel
func (o NSAccessibilitySliderObject) AccessibilityLabel() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("accessibilityLabel"))
	return foundation.NSStringFromID(rv).String()
}

// Decrements the slider’s value.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// This method must post an [valueChanged] notification after changing the
// slider’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySlider/accessibilityPerformDecrement()
//
// [valueChanged]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/valueChanged
func (o NSAccessibilitySliderObject) AccessibilityPerformDecrement() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformDecrement"))
	return rv
}

// Increments the slider’s value.
//
// # Return Value
//
// true if the action was successfully triggered; otherwise, false. This
// method does not indicate the success or failure of the action, just the
// fact that the action was successfully triggered.
//
// # Discussion
//
// This method must post an [valueChanged] notification after changing the
// slider’s value.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySlider/accessibilityPerformIncrement()
//
// [valueChanged]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/valueChanged
func (o NSAccessibilitySliderObject) AccessibilityPerformIncrement() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("accessibilityPerformIncrement"))
	return rv
}

// Returns the slider’s value.
//
// # Return Value
//
// The value for the slider.
//
// # Discussion
//
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityValue] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySlider/accessibilityValue()
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
func (o NSAccessibilitySliderObject) AccessibilityValue() objectivec.IObject {
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
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilityElementProtocol/accessibilityFrame()
//
// [accessibilityFrame]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityFrame
// [position]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/position
// [size]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Attribute/size
func (o NSAccessibilitySliderObject) AccessibilityFrame() corefoundation.CGRect {
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
func (o NSAccessibilitySliderObject) AccessibilityParent() objectivec.IObject {
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
func (o NSAccessibilitySliderObject) AccessibilityIdentifier() string {
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
func (o NSAccessibilitySliderObject) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
}
