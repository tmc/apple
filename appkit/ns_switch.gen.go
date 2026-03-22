// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSwitch] class.
var (
	_NSSwitchClass     NSSwitchClass
	_NSSwitchClassOnce sync.Once
)

func getNSSwitchClass() NSSwitchClass {
	_NSSwitchClassOnce.Do(func() {
		_NSSwitchClass = NSSwitchClass{class: objc.GetClass("NSSwitch")}
	})
	return _NSSwitchClass
}

// GetNSSwitchClass returns the class object for NSSwitch.
func GetNSSwitchClass() NSSwitchClass {
	return getNSSwitchClass()
}

type NSSwitchClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSwitchClass) Alloc() NSSwitch {
	rv := objc.Send[NSSwitch](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A control that offers a binary choice.
//
// # Overview
// 
// The [NSSwitch] class provides a simple interface for displaying and
// toggling a Boolean state, such as on/off. A switch toggles its [NSSwitch.State] and
// sends its [NSSwitch.Action] when clicked, activated through the keyboard, or tapped
// in the Touch Bar. [NSSwitch] also allows dragging between states, and if
// [Continuous] is [true], the switch sends its [NSSwitch.Action] for each change in
// position during the drag.
// 
// [NSSwitch] doesn’t use an instance of [NSCell] to provide its
// functionality. The [cellClass] class property and [cell] instance property
// both return [nil], and they ignore attempts to set a non-[nil] value.
// 
// For design guidance, see Human Interface Guidelines > [Toggles].
//
// [Toggles]: https://developer.apple.com/design/Human-Interface-Guidelines/toggles
// [cellClass]: https://developer.apple.com/documentation/AppKit/NSControl/cellClass
// [cell]: https://developer.apple.com/documentation/AppKit/NSControl/cell
// [nil]: https://developer.apple.com/documentation/ObjectiveC/nil-227m0
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Managing Switch State
//
//   - [NSSwitch.State]: The current position of the switch.
//   - [NSSwitch.SetState]
//
// See: https://developer.apple.com/documentation/AppKit/NSSwitch
type NSSwitch struct {
	NSControl
}

// NSSwitchFromID constructs a [NSSwitch] from an objc.ID.
//
// A control that offers a binary choice.
func NSSwitchFromID(id objc.ID) NSSwitch {
	return NSSwitch{NSControl: NSControlFromID(id)}
}
// NOTE: NSSwitch adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSwitch] class.
//
// # Managing Switch State
//
//   - [INSSwitch.State]: The current position of the switch.
//   - [INSSwitch.SetState]
//
// See: https://developer.apple.com/documentation/AppKit/NSSwitch
type INSSwitch interface {
	INSControl
	NSAccessibilityButton
	NSAccessibilitySwitch

	// Topic: Managing Switch State

	// The current position of the switch.
	State() NSControlStateValue
	SetState(value NSControlStateValue)

	// A Boolean value indicating whether the receiver’s cell sends its action message continuously to its target during mouse tracking.
	IsContinuous() bool
	SetIsContinuous(value bool)
}

// Init initializes the instance.
func (s NSSwitch) Init() NSSwitch {
	rv := objc.Send[NSSwitch](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSwitch) Autorelease() NSSwitch {
	rv := objc.Send[NSSwitch](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSwitch creates a new NSSwitch instance.
func NewNSSwitch() NSSwitch {
	class := getNSSwitchClass()
	rv := objc.Send[NSSwitch](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a control with data in an unarchiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSControl/init(coder:)
func NewSwitchWithCoder(coder foundation.INSCoder) NSSwitch {
	instance := getNSSwitchClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSSwitchFromID(rv)
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
func NewSwitchWithFrame(frameRect corefoundation.CGRect) NSSwitch {
	instance := getNSSwitchClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFrame:"), frameRect)
	return NSSwitchFromID(rv)
}

// Decrements the switch’s value.
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
// switch’s value.
//
// [valueChanged]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/valueChanged
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySwitch/accessibilityPerformDecrement()
func (s NSSwitch) AccessibilityPerformDecrement() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("accessibilityPerformDecrement"))
	return rv
}
// Increments the switch’s value.
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
// switch’s value.
//
// [valueChanged]: https://developer.apple.com/documentation/AppKit/NSAccessibility-swift.struct/Notification/valueChanged
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySwitch/accessibilityPerformIncrement()
func (s NSSwitch) AccessibilityPerformIncrement() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("accessibilityPerformIncrement"))
	return rv
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
func (s NSSwitch) AccessibilityPerformPress() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("accessibilityPerformPress"))
	return rv
}
// Returns the switch’s value.
//
// # Return Value
// 
// The value for the switch.
//
// # Discussion
// 
// This method is the getter for the [NSAccessibilityProtocol] protocol’s
// [accessibilityValue] property.
//
// [accessibilityValue]: https://developer.apple.com/documentation/AppKit/NSAccessibility-c.protocol/accessibilityValue
//
// See: https://developer.apple.com/documentation/AppKit/NSAccessibilitySwitch/accessibilityValue()
func (s NSSwitch) AccessibilityValue() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("accessibilityValue"))
	return foundation.NSStringFromID(rv).String()
}

// The current position of the switch.
//
// # Discussion
// 
// The values [off] and [on] indicate that the switch is in the off or on
// position. The switch treats any value other than [off] as on.
// 
// Setting this property through the [Animator] proxy animates the switch to
// the new value.
//
// [off]: https://developer.apple.com/documentation/AppKit/NSControl/StateValue/off
// [on]: https://developer.apple.com/documentation/AppKit/NSControl/StateValue/on
//
// See: https://developer.apple.com/documentation/AppKit/NSSwitch/state
func (s NSSwitch) State() NSControlStateValue {
	rv := objc.Send[NSControlStateValue](s.ID, objc.Sel("state"))
	return NSControlStateValue(rv)
}
func (s NSSwitch) SetState(value NSControlStateValue) {
	objc.Send[struct{}](s.ID, objc.Sel("setState:"), value)
}
// A Boolean value indicating whether the receiver’s cell sends its action
// message continuously to its target during mouse tracking.
//
// See: https://developer.apple.com/documentation/appkit/nscontrol/iscontinuous
func (s NSSwitch) IsContinuous() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("continuous"))
	return rv
}
func (s NSSwitch) SetIsContinuous(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setContinuous:"), value)
}

			// Protocol methods for NSAccessibilitySwitch
			
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
func (o NSSwitch) AccessibilityFrame() corefoundation.CGRect {
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
func (o NSSwitch) AccessibilityParent() objectivec.IObject {
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
func (o NSSwitch) AccessibilityIdentifier() string {
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
func (o NSSwitch) IsAccessibilityFocused() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessibilityFocused"))
	return rv
	}

