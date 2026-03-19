// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSRotationGestureRecognizer] class.
var (
	_NSRotationGestureRecognizerClass     NSRotationGestureRecognizerClass
	_NSRotationGestureRecognizerClassOnce sync.Once
)

func getNSRotationGestureRecognizerClass() NSRotationGestureRecognizerClass {
	_NSRotationGestureRecognizerClassOnce.Do(func() {
		_NSRotationGestureRecognizerClass = NSRotationGestureRecognizerClass{class: objc.GetClass("NSRotationGestureRecognizer")}
	})
	return _NSRotationGestureRecognizerClass
}

// GetNSRotationGestureRecognizerClass returns the class object for NSRotationGestureRecognizer.
func GetNSRotationGestureRecognizerClass() NSRotationGestureRecognizerClass {
	return getNSRotationGestureRecognizerClass()
}

type NSRotationGestureRecognizerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSRotationGestureRecognizerClass) Alloc() NSRotationGestureRecognizer {
	rv := objc.Send[NSRotationGestureRecognizer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A continuous gesture recognizer that tracks two trackpad touches moving
// opposite each other in a circular motion.
//
// # Overview
// 
// This rotation gesture implies that the underlying view should rotate in a
// matching direction. The gesture is recognized when the trackpad touches
// end.
// 
// Upon creation, the gesture recognizer sets the value of the
// [NSRotationGestureRecognizer.DelaysRotationEvents] property to [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Interpreting the Gesture
//
//   - [NSRotationGestureRecognizer.Rotation]: The rotation of the gesture in radians.
//   - [NSRotationGestureRecognizer.SetRotation]
//   - [NSRotationGestureRecognizer.RotationInDegrees]: The rotation of the gesture in degrees.
//   - [NSRotationGestureRecognizer.SetRotationInDegrees]
//
// See: https://developer.apple.com/documentation/AppKit/NSRotationGestureRecognizer
type NSRotationGestureRecognizer struct {
	NSGestureRecognizer
}

// NSRotationGestureRecognizerFromID constructs a [NSRotationGestureRecognizer] from an objc.ID.
//
// A continuous gesture recognizer that tracks two trackpad touches moving
// opposite each other in a circular motion.
func NSRotationGestureRecognizerFromID(id objc.ID) NSRotationGestureRecognizer {
	return NSRotationGestureRecognizer{NSGestureRecognizer: NSGestureRecognizerFromID(id)}
}
// NOTE: NSRotationGestureRecognizer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSRotationGestureRecognizer] class.
//
// # Interpreting the Gesture
//
//   - [INSRotationGestureRecognizer.Rotation]: The rotation of the gesture in radians.
//   - [INSRotationGestureRecognizer.SetRotation]
//   - [INSRotationGestureRecognizer.RotationInDegrees]: The rotation of the gesture in degrees.
//   - [INSRotationGestureRecognizer.SetRotationInDegrees]
//
// See: https://developer.apple.com/documentation/AppKit/NSRotationGestureRecognizer
type INSRotationGestureRecognizer interface {
	INSGestureRecognizer

	// Topic: Interpreting the Gesture

	// The rotation of the gesture in radians.
	Rotation() float64
	SetRotation(value float64)
	// The rotation of the gesture in degrees.
	RotationInDegrees() float64
	SetRotationInDegrees(value float64)
}

// Init initializes the instance.
func (r NSRotationGestureRecognizer) Init() NSRotationGestureRecognizer {
	rv := objc.Send[NSRotationGestureRecognizer](r.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r NSRotationGestureRecognizer) Autorelease() NSRotationGestureRecognizer {
	rv := objc.Send[NSRotationGestureRecognizer](r.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSRotationGestureRecognizer creates a new NSRotationGestureRecognizer instance.
func NewNSRotationGestureRecognizer() NSRotationGestureRecognizer {
	class := getNSRotationGestureRecognizerClass()
	rv := objc.Send[NSRotationGestureRecognizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/init(coder:)
func NewRotationGestureRecognizerWithCoder(coder foundation.INSCoder) NSRotationGestureRecognizer {
	instance := getNSRotationGestureRecognizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSRotationGestureRecognizerFromID(rv)
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
func NewRotationGestureRecognizerWithTargetAction(target objectivec.IObject, action objc.SEL) NSRotationGestureRecognizer {
	instance := getNSRotationGestureRecognizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTarget:action:"), target, action)
	return NSRotationGestureRecognizerFromID(rv)
}

// The rotation of the gesture in radians.
//
// # Discussion
// 
// This property contains the current rotation in effect for the gesture.
// Changing the value in this property also updates the value in the
// [RotationInDegrees] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSRotationGestureRecognizer/rotation
func (r NSRotationGestureRecognizer) Rotation() float64 {
	rv := objc.Send[float64](r.ID, objc.Sel("rotation"))
	return rv
}
func (r NSRotationGestureRecognizer) SetRotation(value float64) {
	objc.Send[struct{}](r.ID, objc.Sel("setRotation:"), value)
}
// The rotation of the gesture in degrees.
//
// # Discussion
// 
// This property contains the current rotation in effect for the gesture.
// Changing the value in this property also updates the value in the
// [Rotation] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSRotationGestureRecognizer/rotationInDegrees
func (r NSRotationGestureRecognizer) RotationInDegrees() float64 {
	rv := objc.Send[float64](r.ID, objc.Sel("rotationInDegrees"))
	return rv
}
func (r NSRotationGestureRecognizer) SetRotationInDegrees(value float64) {
	objc.Send[struct{}](r.ID, objc.Sel("setRotationInDegrees:"), value)
}

