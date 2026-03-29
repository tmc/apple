// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPanGestureRecognizer] class.
var (
	_NSPanGestureRecognizerClass     NSPanGestureRecognizerClass
	_NSPanGestureRecognizerClassOnce sync.Once
)

func getNSPanGestureRecognizerClass() NSPanGestureRecognizerClass {
	_NSPanGestureRecognizerClassOnce.Do(func() {
		_NSPanGestureRecognizerClass = NSPanGestureRecognizerClass{class: objc.GetClass("NSPanGestureRecognizer")}
	})
	return _NSPanGestureRecognizerClass
}

// GetNSPanGestureRecognizerClass returns the class object for NSPanGestureRecognizer.
func GetNSPanGestureRecognizerClass() NSPanGestureRecognizerClass {
	return getNSPanGestureRecognizerClass()
}

type NSPanGestureRecognizerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPanGestureRecognizerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPanGestureRecognizerClass) Alloc() NSPanGestureRecognizer {
	rv := objc.Send[NSPanGestureRecognizer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A continuous gesture recognizer for panning gestures.
//
// # Overview
// 
// The gesture is recognized when the user clicks all of specified buttons,
// drags the mouse, and releases one or more of the buttons. Use the pan
// gesture recognizer object to retrieve the distance traveled during the pan
// and the location of the mouse as it pans.
// 
// Upon creation, the gesture recognizer is configured to recognize pan
// gestures involving only the primary button. It also delays sending primary
// button events to the view by setting the [NSPanGestureRecognizer.DelaysPrimaryMouseButtonEvents]
// property to [true]. To change the set of buttons to track, modify the
// [NSPanGestureRecognizer.ButtonMask] property.
// 
// In this gesture recognizer, the [LocationInView] method always reports the
// current mouse point, which changes as the user drags the mouse.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Configuring the Gesture Recognizer
//
//   - [NSPanGestureRecognizer.ButtonMask]: A bit mask of the button (or buttons) required to recognize this gesture.
//   - [NSPanGestureRecognizer.SetButtonMask]
//   - [NSPanGestureRecognizer.NumberOfTouchesRequired]: The number of necessary touches on a Touch Bar for the gesture recognizer to match.
//   - [NSPanGestureRecognizer.SetNumberOfTouchesRequired]
//
// # Tracking the Location and Velocity of the Gesture
//
//   - [NSPanGestureRecognizer.TranslationInView]: The distance traveled by the mouse during the gesture.
//   - [NSPanGestureRecognizer.SetTranslationInView]: Changes the current translation value of the gesture recognizer.
//   - [NSPanGestureRecognizer.VelocityInView]: The velocity of the pan, measured in points per second.
//
// See: https://developer.apple.com/documentation/AppKit/NSPanGestureRecognizer
type NSPanGestureRecognizer struct {
	NSGestureRecognizer
}

// NSPanGestureRecognizerFromID constructs a [NSPanGestureRecognizer] from an objc.ID.
//
// A continuous gesture recognizer for panning gestures.
func NSPanGestureRecognizerFromID(id objc.ID) NSPanGestureRecognizer {
	return NSPanGestureRecognizer{NSGestureRecognizer: NSGestureRecognizerFromID(id)}
}
// NOTE: NSPanGestureRecognizer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPanGestureRecognizer] class.
//
// # Configuring the Gesture Recognizer
//
//   - [INSPanGestureRecognizer.ButtonMask]: A bit mask of the button (or buttons) required to recognize this gesture.
//   - [INSPanGestureRecognizer.SetButtonMask]
//   - [INSPanGestureRecognizer.NumberOfTouchesRequired]: The number of necessary touches on a Touch Bar for the gesture recognizer to match.
//   - [INSPanGestureRecognizer.SetNumberOfTouchesRequired]
//
// # Tracking the Location and Velocity of the Gesture
//
//   - [INSPanGestureRecognizer.TranslationInView]: The distance traveled by the mouse during the gesture.
//   - [INSPanGestureRecognizer.SetTranslationInView]: Changes the current translation value of the gesture recognizer.
//   - [INSPanGestureRecognizer.VelocityInView]: The velocity of the pan, measured in points per second.
//
// See: https://developer.apple.com/documentation/AppKit/NSPanGestureRecognizer
type INSPanGestureRecognizer interface {
	INSGestureRecognizer

	// Topic: Configuring the Gesture Recognizer

	// A bit mask of the button (or buttons) required to recognize this gesture.
	ButtonMask() uint
	SetButtonMask(value uint)
	// The number of necessary touches on a Touch Bar for the gesture recognizer to match.
	NumberOfTouchesRequired() int
	SetNumberOfTouchesRequired(value int)

	// Topic: Tracking the Location and Velocity of the Gesture

	// The distance traveled by the mouse during the gesture.
	TranslationInView(view INSView) corefoundation.CGPoint
	// Changes the current translation value of the gesture recognizer.
	SetTranslationInView(translation corefoundation.CGPoint, view INSView)
	// The velocity of the pan, measured in points per second.
	VelocityInView(view INSView) corefoundation.CGPoint
}

// Init initializes the instance.
func (p NSPanGestureRecognizer) Init() NSPanGestureRecognizer {
	rv := objc.Send[NSPanGestureRecognizer](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPanGestureRecognizer) Autorelease() NSPanGestureRecognizer {
	rv := objc.Send[NSPanGestureRecognizer](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPanGestureRecognizer creates a new NSPanGestureRecognizer instance.
func NewNSPanGestureRecognizer() NSPanGestureRecognizer {
	class := getNSPanGestureRecognizerClass()
	rv := objc.Send[NSPanGestureRecognizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/init(coder:)
func NewPanGestureRecognizerWithCoder(coder foundation.INSCoder) NSPanGestureRecognizer {
	instance := getNSPanGestureRecognizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSPanGestureRecognizerFromID(rv)
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
func NewPanGestureRecognizerWithTargetAction(target objectivec.IObject, action objc.SEL) NSPanGestureRecognizer {
	instance := getNSPanGestureRecognizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTarget:action:"), target, action)
	return NSPanGestureRecognizerFromID(rv)
}

// The distance traveled by the mouse during the gesture.
//
// view: The view in whose coordinate system the translation of the pan gesture
// should be computed. The view’s transform is applied to the distance
// values.
//
// # Return Value
// 
// A point whose x and y values correspond to the total distance travelled
// since the beginning of the gesture.
//
// # Discussion
// 
// The x and y values of the returned point report the total translation over
// time. They are not delta values from the last time that the translation was
// reported. To determine the starting point of the gesture, subtract the
// current translation values from the current location of the mouse in the
// same view.
//
// See: https://developer.apple.com/documentation/AppKit/NSPanGestureRecognizer/translation(in:)
func (p NSPanGestureRecognizer) TranslationInView(view INSView) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](p.ID, objc.Sel("translationInView:"), view)
	return corefoundation.CGPoint(rv)
}
// Changes the current translation value of the gesture recognizer.
//
// translation: The new translation values to use in the gesture recognizer.
//
// view: The view in whose coordinate system you specified the new translation
// value. Specifying `nil` resets the previous translation value.
//
// # Discussion
// 
// This method changes the current translation value of the gesture
// recognizer. Changing the value resets the velocity of the pan. You might
// call this method at mouse-down time to adjust the translation value and
// make it relative to some specific point in your view.
//
// See: https://developer.apple.com/documentation/AppKit/NSPanGestureRecognizer/setTranslation(_:in:)
func (p NSPanGestureRecognizer) SetTranslationInView(translation corefoundation.CGPoint, view INSView) {
	objc.Send[objc.ID](p.ID, objc.Sel("setTranslation:inView:"), translation, view)
}
// The velocity of the pan, measured in points per second.
//
// view: The view that provides the coordinate system for computing the velocity
// value. This parameter must not be `nil`.
//
// # Return Value
// 
// The horizontal and vertical velocity of the pan gesture. These values are
// relative to the specified view.
//
// See: https://developer.apple.com/documentation/AppKit/NSPanGestureRecognizer/velocity(in:)
func (p NSPanGestureRecognizer) VelocityInView(view INSView) corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](p.ID, objc.Sel("velocityInView:"), view)
	return corefoundation.CGPoint(rv)
}

// A bit mask of the button (or buttons) required to recognize this gesture.
//
// # Discussion
// 
// Bit 0 represents the primary button, bit 1 is the secondary button, and so
// on. So to track clicks of the secondary button, assign the value `0x2`
// (which corresponds to a `1` in bit 1) to this property. The default value
// of this property is 0x1, which detects clicks in the primary mouse button.
// 
// Changing the value of this property also sets the values of the
// [DelaysPrimaryMouseButtonEvents], [DelaysSecondaryMouseButtonEvents], and
// [DelaysOtherMouseButtonEvents] properties to [true] for each of the buttons
// you specified.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSPanGestureRecognizer/buttonMask
func (p NSPanGestureRecognizer) ButtonMask() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("buttonMask"))
	return rv
}
func (p NSPanGestureRecognizer) SetButtonMask(value uint) {
	objc.Send[struct{}](p.ID, objc.Sel("setButtonMask:"), value)
}
// The number of necessary touches on a Touch Bar for the gesture recognizer
// to match.
//
// See: https://developer.apple.com/documentation/AppKit/NSPanGestureRecognizer/numberOfTouchesRequired
func (p NSPanGestureRecognizer) NumberOfTouchesRequired() int {
	rv := objc.Send[int](p.ID, objc.Sel("numberOfTouchesRequired"))
	return rv
}
func (p NSPanGestureRecognizer) SetNumberOfTouchesRequired(value int) {
	objc.Send[struct{}](p.ID, objc.Sel("setNumberOfTouchesRequired:"), value)
}

