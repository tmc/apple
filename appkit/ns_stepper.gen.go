// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSStepper] class.
var (
	_NSStepperClass     NSStepperClass
	_NSStepperClassOnce sync.Once
)

func getNSStepperClass() NSStepperClass {
	_NSStepperClassOnce.Do(func() {
		_NSStepperClass = NSStepperClass{class: objc.GetClass("NSStepper")}
	})
	return _NSStepperClass
}

// GetNSStepperClass returns the class object for NSStepper.
func GetNSStepperClass() NSStepperClass {
	return getNSStepperClass()
}

type NSStepperClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSStepperClass) Alloc() NSStepper {
	rv := objc.Send[NSStepper](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An interface with up and down arrow buttons for incrementing or
// decrementing a value.
//
// # Overview
// 
// A stepper consists of two small arrows that can increment and decrement a
// value that appears beside it, such as a date or time. The illustration
// below shows a stepper to the right of a text field, which would show the
// stepper’s value.
// 
// [media-2555815]
// 
// The [NSStepper] class uses the [NSStepperCell] class to implement its user
// interface.
//
// # Specifying value range
//
//   - [NSStepper.MaxValue]: The stepper’s maximum value.
//   - [NSStepper.SetMaxValue]
//   - [NSStepper.MinValue]: The stepper’s minimum value.
//   - [NSStepper.SetMinValue]
//   - [NSStepper.Increment]: The amount by which the receiver changes with each increment or decrement.
//   - [NSStepper.SetIncrement]
//
// # Specifying how the stepper responds
//
//   - [NSStepper.Autorepeat]: A Boolean value that indicates how the stepper responds to mouse events.
//   - [NSStepper.SetAutorepeat]
//   - [NSStepper.ValueWraps]: A Boolean value that indicates whether the stepper wraps around the minimum and maximum values.
//   - [NSStepper.SetValueWraps]
//
// See: https://developer.apple.com/documentation/AppKit/NSStepper
type NSStepper struct {
	NSControl
}

// NSStepperFromID constructs a [NSStepper] from an objc.ID.
//
// An interface with up and down arrow buttons for incrementing or
// decrementing a value.
func NSStepperFromID(id objc.ID) NSStepper {
	return NSStepper{NSControl: NSControlFromID(id)}
}
// NOTE: NSStepper adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSStepper] class.
//
// # Specifying value range
//
//   - [INSStepper.MaxValue]: The stepper’s maximum value.
//   - [INSStepper.SetMaxValue]
//   - [INSStepper.MinValue]: The stepper’s minimum value.
//   - [INSStepper.SetMinValue]
//   - [INSStepper.Increment]: The amount by which the receiver changes with each increment or decrement.
//   - [INSStepper.SetIncrement]
//
// # Specifying how the stepper responds
//
//   - [INSStepper.Autorepeat]: A Boolean value that indicates how the stepper responds to mouse events.
//   - [INSStepper.SetAutorepeat]
//   - [INSStepper.ValueWraps]: A Boolean value that indicates whether the stepper wraps around the minimum and maximum values.
//   - [INSStepper.SetValueWraps]
//
// See: https://developer.apple.com/documentation/AppKit/NSStepper
type INSStepper interface {
	INSControl
	NSAccessibilityStepper

	// Topic: Specifying value range

	// The stepper’s maximum value.
	MaxValue() float64
	SetMaxValue(value float64)
	// The stepper’s minimum value.
	MinValue() float64
	SetMinValue(value float64)
	// The amount by which the receiver changes with each increment or decrement.
	Increment() float64
	SetIncrement(value float64)

	// Topic: Specifying how the stepper responds

	// A Boolean value that indicates how the stepper responds to mouse events.
	Autorepeat() bool
	SetAutorepeat(value bool)
	// A Boolean value that indicates whether the stepper wraps around the minimum and maximum values.
	ValueWraps() bool
	SetValueWraps(value bool)
}

// Init initializes the instance.
func (s NSStepper) Init() NSStepper {
	rv := objc.Send[NSStepper](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSStepper) Autorelease() NSStepper {
	rv := objc.Send[NSStepper](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSStepper creates a new NSStepper instance.
func NewNSStepper() NSStepper {
	class := getNSStepperClass()
	rv := objc.Send[NSStepper](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewStepperWithCoder(coder foundation.INSCoder) NSStepper {
	instance := getNSStepperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSStepperFromID(rv)
}

// Initializes a control with the specified frame rectangle.
//
// frameRect: The rectangle of the control, specified in points in the coordinate space
// of the enclosing view.
//
// # Return Value
// 
// An initialized control object, or `nil` if the object couldn’t be
// initialized.
//
// # Discussion
// 
// If a cell has been specified for controls of this type, this method also
// creates an instance of the cell. Because [NSControl] is an abstract class,
// invocations of this method should appear only in the designated
// initializers of subclasses; that is, there should always be a more specific
// designated initializer for the subclass, as this method is the designated
// initializer for [NSControl].
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(frame:)
func NewStepperWithFrame(frameRect corefoundation.CGRect) NSStepper {
	instance := getNSStepperClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSStepperFromID(rv)
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
func (s NSStepper) AccessibilityPerformDecrement() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("accessibilityPerformDecrement"))
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
func (s NSStepper) AccessibilityPerformIncrement() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("accessibilityPerformIncrement"))
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
func (s NSStepper) AccessibilityValue() objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("accessibilityValue"))
	return objectivec.Object{ID: rv}
}

// The stepper’s maximum value.
//
// # Discussion
// 
// The default is 59.
//
// See: https://developer.apple.com/documentation/AppKit/NSStepper/maxValue
func (s NSStepper) MaxValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("maxValue"))
	return rv
}
func (s NSStepper) SetMaxValue(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMaxValue:"), value)
}
// The stepper’s minimum value.
//
// # Discussion
// 
// The default is 0.
//
// See: https://developer.apple.com/documentation/AppKit/NSStepper/minValue
func (s NSStepper) MinValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("minValue"))
	return rv
}
func (s NSStepper) SetMinValue(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setMinValue:"), value)
}
// The amount by which the receiver changes with each increment or decrement.
//
// # Discussion
// 
// The default is 1.
//
// See: https://developer.apple.com/documentation/AppKit/NSStepper/increment
func (s NSStepper) Increment() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("increment"))
	return rv
}
func (s NSStepper) SetIncrement(value float64) {
	objc.Send[struct{}](s.ID, objc.Sel("setIncrement:"), value)
}
// A Boolean value that indicates how the stepper responds to mouse events.
//
// # Discussion
// 
// [true] if the first mouse down does one increment (or decrement) and, after
// a delay of 0.5 seconds, increments (or decrements) at a rate of ten times
// per second. [false] if the receiver does one increment (decrement) on a
// mouse up. The default is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSStepper/autorepeat
func (s NSStepper) Autorepeat() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("autorepeat"))
	return rv
}
func (s NSStepper) SetAutorepeat(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setAutorepeat:"), value)
}
// A Boolean value that indicates whether the stepper wraps around the minimum
// and maximum values.
//
// # Discussion
// 
// [true] if, when incrementing or decrementing, the value wraps around to the
// minimum or maximum. [false] if the value stays pinned at the minimum or
// maximum. The default is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSStepper/valueWraps
func (s NSStepper) ValueWraps() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("valueWraps"))
	return rv
}
func (s NSStepper) SetValueWraps(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setValueWraps:"), value)
}

			// Protocol methods for NSAccessibilityStepper
			
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
func (o NSStepper) AccessibilityFrame() corefoundation.CGRect {
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
func (o NSStepper) AccessibilityParent() objectivec.IObject {
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
func (o NSStepper) AccessibilityIdentifier() string {
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
func (o NSStepper) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}

