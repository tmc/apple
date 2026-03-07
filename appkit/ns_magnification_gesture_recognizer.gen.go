// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMagnificationGestureRecognizer] class.
var (
	_NSMagnificationGestureRecognizerClass     NSMagnificationGestureRecognizerClass
	_NSMagnificationGestureRecognizerClassOnce sync.Once
)

func getNSMagnificationGestureRecognizerClass() NSMagnificationGestureRecognizerClass {
	_NSMagnificationGestureRecognizerClassOnce.Do(func() {
		_NSMagnificationGestureRecognizerClass = NSMagnificationGestureRecognizerClass{class: objc.GetClass("NSMagnificationGestureRecognizer")}
	})
	return _NSMagnificationGestureRecognizerClass
}

// GetNSMagnificationGestureRecognizerClass returns the class object for NSMagnificationGestureRecognizer.
func GetNSMagnificationGestureRecognizerClass() NSMagnificationGestureRecognizerClass {
	return getNSMagnificationGestureRecognizerClass()
}

type NSMagnificationGestureRecognizerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMagnificationGestureRecognizerClass) Alloc() NSMagnificationGestureRecognizer {
	rv := objc.Send[NSMagnificationGestureRecognizer](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A continuous gesture recognizer that tracks a pinch gesture that magnifies
// content.
//
// # Overview
// 
// This object tracks pinch gestures on a track pad or other input device and
// stores the resulting magnification value for you to use in your code.
// 
// This gesture recognizer automatically sets the value of the
// [NSMagnificationGestureRecognizer.DelaysMagnificationEvents] property to [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Finding the Magnification Factor
//
//   - [NSMagnificationGestureRecognizer.Magnification]: The amount of magnification to apply.
//   - [NSMagnificationGestureRecognizer.SetMagnification]
//
// See: https://developer.apple.com/documentation/AppKit/NSMagnificationGestureRecognizer
type NSMagnificationGestureRecognizer struct {
	NSGestureRecognizer
}

// NSMagnificationGestureRecognizerFromID constructs a [NSMagnificationGestureRecognizer] from an objc.ID.
//
// A continuous gesture recognizer that tracks a pinch gesture that magnifies
// content.
func NSMagnificationGestureRecognizerFromID(id objc.ID) NSMagnificationGestureRecognizer {
	return NSMagnificationGestureRecognizer{NSGestureRecognizer: NSGestureRecognizerFromID(id)}
}
// NOTE: NSMagnificationGestureRecognizer adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSMagnificationGestureRecognizer] class.
//
// # Finding the Magnification Factor
//
//   - [INSMagnificationGestureRecognizer.Magnification]: The amount of magnification to apply.
//   - [INSMagnificationGestureRecognizer.SetMagnification]
//
// See: https://developer.apple.com/documentation/AppKit/NSMagnificationGestureRecognizer
type INSMagnificationGestureRecognizer interface {
	INSGestureRecognizer

	// Topic: Finding the Magnification Factor

	// The amount of magnification to apply.
	Magnification() float64
	SetMagnification(value float64)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (m NSMagnificationGestureRecognizer) Init() NSMagnificationGestureRecognizer {
	rv := objc.Send[NSMagnificationGestureRecognizer](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMagnificationGestureRecognizer) Autorelease() NSMagnificationGestureRecognizer {
	rv := objc.Send[NSMagnificationGestureRecognizer](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMagnificationGestureRecognizer creates a new NSMagnificationGestureRecognizer instance.
func NewNSMagnificationGestureRecognizer() NSMagnificationGestureRecognizer {
	class := getNSMagnificationGestureRecognizerClass()
	rv := objc.Send[NSMagnificationGestureRecognizer](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppKit/NSGestureRecognizer/init(coder:)
func NewMagnificationGestureRecognizerWithCoder(coder foundation.INSCoder) NSMagnificationGestureRecognizer {
	instance := getNSMagnificationGestureRecognizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSMagnificationGestureRecognizerFromID(rv)
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
func NewMagnificationGestureRecognizerWithTargetAction(target objectivec.IObject, action objc.SEL) NSMagnificationGestureRecognizer {
	instance := getNSMagnificationGestureRecognizerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTarget:action:"), target, action)
	return NSMagnificationGestureRecognizerFromID(rv)
}






func (m NSMagnificationGestureRecognizer) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](m.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The amount of magnification to apply.
//
// # Discussion
// 
// This property contains the current magnification in effect for the gesture.
// Add the value of this property to `1.0` to get the scale factor to apply to
// your content.
//
// See: https://developer.apple.com/documentation/AppKit/NSMagnificationGestureRecognizer/magnification
func (m NSMagnificationGestureRecognizer) Magnification() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("magnification"))
	return rv
}
func (m NSMagnificationGestureRecognizer) SetMagnification(value float64) {
	objc.Send[struct{}](m.ID, objc.Sel("setMagnification:"), value)
}


























