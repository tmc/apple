// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPressGestureRecognizer] class.
var (
	_NSPressGestureRecognizerClass     NSPressGestureRecognizerClass
	_NSPressGestureRecognizerClassOnce sync.Once
)

func getNSPressGestureRecognizerClass() NSPressGestureRecognizerClass {
	_NSPressGestureRecognizerClassOnce.Do(func() {
		_NSPressGestureRecognizerClass = NSPressGestureRecognizerClass{class: objc.GetClass("NSPressGestureRecognizer")}
	})
	return _NSPressGestureRecognizerClass
}

// GetNSPressGestureRecognizerClass returns the class object for NSPressGestureRecognizer.
func GetNSPressGestureRecognizerClass() NSPressGestureRecognizerClass {
	return getNSPressGestureRecognizerClass()
}

type NSPressGestureRecognizerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPressGestureRecognizerClass) Alloc() NSPressGestureRecognizer {
	rv := objc.Send[NSPressGestureRecognizer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A discrete gesture recognizer that tracks whether the user holds down a
// mouse button for a minimum amount of time before releasing it.
//
// # Overview
// 
// Use a press gesture recognizer to configure which button the user must hold
// and the length of time they must hold it. You can also specify how far the
// mouse can move for a valid gesture.
// 
// Upon creation, the gesture recognizer recognizes press gestures involving
// only the primary button. It also delays sending primary button events to
// the view by setting the [NSPressGestureRecognizer.DelaysPrimaryMouseButtonEvents] property to
// [true]. To change the set of buttons to track, modify the [NSPressGestureRecognizer.ButtonMask]
// property.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Configuring the Gesture Recognizer
//
//   - [NSPressGestureRecognizer.AllowableMovement]: The maximum movement of the mouse in the view before the gesture fails.
//   - [NSPressGestureRecognizer.SetAllowableMovement]
//   - [NSPressGestureRecognizer.ButtonMask]: A bit mask of the buttons required to recognize this press.
//   - [NSPressGestureRecognizer.SetButtonMask]
//   - [NSPressGestureRecognizer.MinimumPressDuration]: The minimum time (in seconds) that the user must hold the mouse button in the view for a valid gesture.
//   - [NSPressGestureRecognizer.SetMinimumPressDuration]
//   - [NSPressGestureRecognizer.NumberOfTouchesRequired]: The number of necessary touches on a Touch Bar for the gesture recognizer to match.
//   - [NSPressGestureRecognizer.SetNumberOfTouchesRequired]
//
// See: https://developer.apple.com/documentation/AppKit/NSPressGestureRecognizer
type NSPressGestureRecognizer struct {
	NSGestureRecognizer
}

// NSPressGestureRecognizerFromID constructs a [NSPressGestureRecognizer] from an objc.ID.
//
// A discrete gesture recognizer that tracks whether the user holds down a
// mouse button for a minimum amount of time before releasing it.
func NSPressGestureRecognizerFromID(id objc.ID) NSPressGestureRecognizer {
	return NSPressGestureRecognizer{NSGestureRecognizer: NSGestureRecognizerFromID(id)}
}
// NOTE: NSPressGestureRecognizer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPressGestureRecognizer] class.
//
// # Configuring the Gesture Recognizer
//
//   - [INSPressGestureRecognizer.AllowableMovement]: The maximum movement of the mouse in the view before the gesture fails.
//   - [INSPressGestureRecognizer.SetAllowableMovement]
//   - [INSPressGestureRecognizer.ButtonMask]: A bit mask of the buttons required to recognize this press.
//   - [INSPressGestureRecognizer.SetButtonMask]
//   - [INSPressGestureRecognizer.MinimumPressDuration]: The minimum time (in seconds) that the user must hold the mouse button in the view for a valid gesture.
//   - [INSPressGestureRecognizer.SetMinimumPressDuration]
//   - [INSPressGestureRecognizer.NumberOfTouchesRequired]: The number of necessary touches on a Touch Bar for the gesture recognizer to match.
//   - [INSPressGestureRecognizer.SetNumberOfTouchesRequired]
//
// See: https://developer.apple.com/documentation/AppKit/NSPressGestureRecognizer
type INSPressGestureRecognizer interface {
	INSGestureRecognizer

	// Topic: Configuring the Gesture Recognizer

	// The maximum movement of the mouse in the view before the gesture fails.
	AllowableMovement() float64
	SetAllowableMovement(value float64)
	// A bit mask of the buttons required to recognize this press.
	ButtonMask() uint
	SetButtonMask(value uint)
	// The minimum time (in seconds) that the user must hold the mouse button in the view for a valid gesture.
	MinimumPressDuration() float64
	SetMinimumPressDuration(value float64)
	// The number of necessary touches on a Touch Bar for the gesture recognizer to match.
	NumberOfTouchesRequired() int
	SetNumberOfTouchesRequired(value int)
}

// Init initializes the instance.
func (p NSPressGestureRecognizer) Init() NSPressGestureRecognizer {
	rv := objc.Send[NSPressGestureRecognizer](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPressGestureRecognizer) Autorelease() NSPressGestureRecognizer {
	rv := objc.Send[NSPressGestureRecognizer](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPressGestureRecognizer creates a new NSPressGestureRecognizer instance.
func NewNSPressGestureRecognizer() NSPressGestureRecognizer {
	class := getNSPressGestureRecognizerClass()
	rv := objc.Send[NSPressGestureRecognizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/init(coder:)
func NewPressGestureRecognizerWithCoder(coder foundation.INSCoder) NSPressGestureRecognizer {
	instance := getNSPressGestureRecognizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSPressGestureRecognizerFromID(rv)
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
func NewPressGestureRecognizerWithTargetAction(target objectivec.IObject, action objc.SEL) NSPressGestureRecognizer {
	instance := getNSPressGestureRecognizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTarget:action:"), target, action)
	return NSPressGestureRecognizerFromID(rv)
}

// The maximum movement of the mouse in the view before the gesture fails.
//
// # Discussion
// 
// The mouse must move by the specified amount along either axis for the
// gesture to fail. The distance is measured in points. The default value of
// this property is the same as the double-click distance.
//
// See: https://developer.apple.com/documentation/AppKit/NSPressGestureRecognizer/allowableMovement
func (p NSPressGestureRecognizer) AllowableMovement() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("allowableMovement"))
	return rv
}
func (p NSPressGestureRecognizer) SetAllowableMovement(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setAllowableMovement:"), value)
}

// A bit mask of the buttons required to recognize this press.
//
// # Discussion
// 
// Bit 0 represents the primary button, bit 1 is the secondary button, and so
// on. So to track clicks of the secondary button, assign the value `0x2`
// (which corresponds to a `1` in bit 1) to this property. The default value
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
// See: https://developer.apple.com/documentation/AppKit/NSPressGestureRecognizer/buttonMask
func (p NSPressGestureRecognizer) ButtonMask() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("buttonMask"))
	return rv
}
func (p NSPressGestureRecognizer) SetButtonMask(value uint) {
	objc.Send[struct{}](p.ID, objc.Sel("setButtonMask:"), value)
}

// The minimum time (in seconds) that the user must hold the mouse button in
// the view for a valid gesture.
//
// # Discussion
// 
// The default value of this property is the same as the current double-click
// interval.
//
// See: https://developer.apple.com/documentation/AppKit/NSPressGestureRecognizer/minimumPressDuration
func (p NSPressGestureRecognizer) MinimumPressDuration() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("minimumPressDuration"))
	return rv
}
func (p NSPressGestureRecognizer) SetMinimumPressDuration(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setMinimumPressDuration:"), value)
}

// The number of necessary touches on a Touch Bar for the gesture recognizer
// to match.
//
// See: https://developer.apple.com/documentation/AppKit/NSPressGestureRecognizer/numberOfTouchesRequired
func (p NSPressGestureRecognizer) NumberOfTouchesRequired() int {
	rv := objc.Send[int](p.ID, objc.Sel("numberOfTouchesRequired"))
	return rv
}
func (p NSPressGestureRecognizer) SetNumberOfTouchesRequired(value int) {
	objc.Send[struct{}](p.ID, objc.Sel("setNumberOfTouchesRequired:"), value)
}

