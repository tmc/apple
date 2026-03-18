// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSClickGestureRecognizer] class.
var (
	_NSClickGestureRecognizerClass     NSClickGestureRecognizerClass
	_NSClickGestureRecognizerClassOnce sync.Once
)

func getNSClickGestureRecognizerClass() NSClickGestureRecognizerClass {
	_NSClickGestureRecognizerClassOnce.Do(func() {
		_NSClickGestureRecognizerClass = NSClickGestureRecognizerClass{class: objc.GetClass("NSClickGestureRecognizer")}
	})
	return _NSClickGestureRecognizerClass
}

// GetNSClickGestureRecognizerClass returns the class object for NSClickGestureRecognizer.
func GetNSClickGestureRecognizerClass() NSClickGestureRecognizerClass {
	return getNSClickGestureRecognizerClass()
}

type NSClickGestureRecognizerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSClickGestureRecognizerClass) Alloc() NSClickGestureRecognizer {
	rv := objc.Send[NSClickGestureRecognizer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A discrete gesture recognizer that tracks a specified number of mouse
// clicks.
//
// # Overview
// 
// When configuring this gesture recognizer, you can specify which mouse
// buttons must be clicked and how many clicks must occur before the action
// method is called. The user must click the specified mouse button the
// required number of times without dragging the mouse for the gesture to be
// recognized.
// 
// The gesture recognizer automatically sets the values of the
// [NSClickGestureRecognizer.DelaysPrimaryMouseButtonEvents], [NSClickGestureRecognizer.DelaysSecondaryMouseButtonEvents], and
// [NSClickGestureRecognizer.DelaysOtherMouseButtonEvents] properties to [true] for each button in the
// [NSClickGestureRecognizer.ButtonMask] property.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Configuring the Gesture
//
//   - [NSClickGestureRecognizer.ButtonMask]: A bit mask of the button (or buttons) required to recognize this click.
//   - [NSClickGestureRecognizer.SetButtonMask]
//   - [NSClickGestureRecognizer.NumberOfClicksRequired]: The number of clicks required to match.
//   - [NSClickGestureRecognizer.SetNumberOfClicksRequired]
//   - [NSClickGestureRecognizer.NumberOfTouchesRequired]: The number of touches required in an [NSTouchBar](<doc://com.apple.appkit/documentation/AppKit/NSTouchBar>) object for the gesture recognizer to match.
//   - [NSClickGestureRecognizer.SetNumberOfTouchesRequired]
//
// See: https://developer.apple.com/documentation/AppKit/NSClickGestureRecognizer
type NSClickGestureRecognizer struct {
	NSGestureRecognizer
}

// NSClickGestureRecognizerFromID constructs a [NSClickGestureRecognizer] from an objc.ID.
//
// A discrete gesture recognizer that tracks a specified number of mouse
// clicks.
func NSClickGestureRecognizerFromID(id objc.ID) NSClickGestureRecognizer {
	return NSClickGestureRecognizer{NSGestureRecognizer: NSGestureRecognizerFromID(id)}
}
// NOTE: NSClickGestureRecognizer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSClickGestureRecognizer] class.
//
// # Configuring the Gesture
//
//   - [INSClickGestureRecognizer.ButtonMask]: A bit mask of the button (or buttons) required to recognize this click.
//   - [INSClickGestureRecognizer.SetButtonMask]
//   - [INSClickGestureRecognizer.NumberOfClicksRequired]: The number of clicks required to match.
//   - [INSClickGestureRecognizer.SetNumberOfClicksRequired]
//   - [INSClickGestureRecognizer.NumberOfTouchesRequired]: The number of touches required in an [NSTouchBar](<doc://com.apple.appkit/documentation/AppKit/NSTouchBar>) object for the gesture recognizer to match.
//   - [INSClickGestureRecognizer.SetNumberOfTouchesRequired]
//
// See: https://developer.apple.com/documentation/AppKit/NSClickGestureRecognizer
type INSClickGestureRecognizer interface {
	INSGestureRecognizer

	// Topic: Configuring the Gesture

	// A bit mask of the button (or buttons) required to recognize this click.
	ButtonMask() uint
	SetButtonMask(value uint)
	// The number of clicks required to match.
	NumberOfClicksRequired() int
	SetNumberOfClicksRequired(value int)
	// The number of touches required in an [NSTouchBar](<doc://com.apple.appkit/documentation/AppKit/NSTouchBar>) object for the gesture recognizer to match.
	NumberOfTouchesRequired() int
	SetNumberOfTouchesRequired(value int)
}

// Init initializes the instance.
func (c NSClickGestureRecognizer) Init() NSClickGestureRecognizer {
	rv := objc.Send[NSClickGestureRecognizer](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSClickGestureRecognizer) Autorelease() NSClickGestureRecognizer {
	rv := objc.Send[NSClickGestureRecognizer](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSClickGestureRecognizer creates a new NSClickGestureRecognizer instance.
func NewNSClickGestureRecognizer() NSClickGestureRecognizer {
	class := getNSClickGestureRecognizerClass()
	rv := objc.Send[NSClickGestureRecognizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/init(coder:)
func NewClickGestureRecognizerWithCoder(coder foundation.INSCoder) NSClickGestureRecognizer {
	instance := getNSClickGestureRecognizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSClickGestureRecognizerFromID(rv)
}

// Initializes the gesture recognizer with the specified target and action
// information.
//
// target: The object whose action method is called when the gesture is recognized.
// You must not specify `nil` for this parameter.
//
// action: A selector that identifies the method to call when the gesture is
// recognized. This method must be implemented by the object in `target`. You
// must not specify `nil` for this parameter.
//
// # Return Value
// 
// The initialized gesture recognizer object or `nil` if an error occurred.
//
// # Discussion
// 
// This method is the designated initializer. Subclasses must call this method
// from their own custom initialization methods. Call the method before
// performing other tasks.
// 
// This method records the specified `target` and `action` values and prepares
// the gesture recognizer for use.
// 
// The `action` method must have one of the following signatures:
//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/init(target:action:)
func NewClickGestureRecognizerWithTargetAction(target objectivec.IObject, action objc.SEL) NSClickGestureRecognizer {
	instance := getNSClickGestureRecognizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTarget:action:"), target, action)
	return NSClickGestureRecognizerFromID(rv)
}

// A bit mask of the button (or buttons) required to recognize this click.
//
// # Discussion
// 
// Bit `0` represents the primary button, bit `1` is the secondary button, and
// so on. So to track clicks of the secondary button, assign the value `0x2`
// (which corresponds to a `1` in bit `1`) to this property. The default value
// of this property is `0x1`, which detects clicks in the primary mouse
// button.
// 
// Changing the value of this property also sets the values of the
// [DelaysPrimaryMouseButtonEvents], [DelaysSecondaryMouseButtonEvents], and
// [DelaysOtherMouseButtonEvents] properties to [true] for each of the buttons
// you specified.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSClickGestureRecognizer/buttonMask
func (c NSClickGestureRecognizer) ButtonMask() uint {
	rv := objc.Send[uint](c.ID, objc.Sel("buttonMask"))
	return rv
}
func (c NSClickGestureRecognizer) SetButtonMask(value uint) {
	objc.Send[struct{}](c.ID, objc.Sel("setButtonMask:"), value)
}

// The number of clicks required to match.
//
// # Discussion
// 
// The value in this property should always be a positive integer. Negative
// integers or `0` cause this object to never recognize its gesture. The
// default value of this property is `1`.
//
// See: https://developer.apple.com/documentation/AppKit/NSClickGestureRecognizer/numberOfClicksRequired
func (c NSClickGestureRecognizer) NumberOfClicksRequired() int {
	rv := objc.Send[int](c.ID, objc.Sel("numberOfClicksRequired"))
	return rv
}
func (c NSClickGestureRecognizer) SetNumberOfClicksRequired(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setNumberOfClicksRequired:"), value)
}

// The number of touches required in an [NSTouchBar] object for the gesture
// recognizer to match.
//
// See: https://developer.apple.com/documentation/AppKit/NSClickGestureRecognizer/numberOfTouchesRequired
func (c NSClickGestureRecognizer) NumberOfTouchesRequired() int {
	rv := objc.Send[int](c.ID, objc.Sel("numberOfTouchesRequired"))
	return rv
}
func (c NSClickGestureRecognizer) SetNumberOfTouchesRequired(value int) {
	objc.Send[struct{}](c.ID, objc.Sel("setNumberOfTouchesRequired:"), value)
}

