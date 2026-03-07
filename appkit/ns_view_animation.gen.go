// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSViewAnimation] class.
var (
	_NSViewAnimationClass     NSViewAnimationClass
	_NSViewAnimationClassOnce sync.Once
)

func getNSViewAnimationClass() NSViewAnimationClass {
	_NSViewAnimationClassOnce.Do(func() {
		_NSViewAnimationClass = NSViewAnimationClass{class: objc.GetClass("NSViewAnimation")}
	})
	return _NSViewAnimationClass
}

// GetNSViewAnimationClass returns the class object for NSViewAnimation.
func GetNSViewAnimationClass() NSViewAnimationClass {
	return getNSViewAnimationClass()
}

type NSViewAnimationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSViewAnimationClass) Alloc() NSViewAnimation {
	rv := objc.Send[NSViewAnimation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// An animation of an app’s views, limited to changes in frame location and
// size, and to fade-in and fade-out effects.
//
// # Overview
// 
// An [NSViewAnimation] object takes an array of dictionaries from which it
// determines the objects to animate and the effects to apply to them. Each
// dictionary must have a target object and, optionally, properties that
// specify beginning and ending frame and whether to fade in or fade out. (See
// [NSViewAnimationKey] for further information.) Animations with
// [NSViewAnimation] are, by default, in non-blocking mode over a duration of
// 0.5 seconds using the ease in-out animation curve. But you can configure
// the animation to have any duration, curve, frame rate, and blocking mode.
// You may also set progress marks, assign a delegate, and implement
// delegation methods in order to animate view and windows concurrent with the
// ones specified as targets in the view-animation dictionary.
// 
// Invoking the [NSAnimation] [StopAnimation] method on a running
// [NSViewAnimation] object moves the animation to the end frame.
//
// # Initializing an NSViewAnimation object
//
//   - [NSViewAnimation.InitWithViewAnimations]: Returns an [NSViewAnimation] object initialized with the supplied information.
//
// # Getting and setting view-animation dictionaries
//
//   - [NSViewAnimation.ViewAnimations]: The dictionaries defining the objects to animate.
//   - [NSViewAnimation.SetViewAnimations]
//
// See: https://developer.apple.com/documentation/AppKit/NSViewAnimation
type NSViewAnimation struct {
	NSAnimation
}

// NSViewAnimationFromID constructs a [NSViewAnimation] from an objc.ID.
//
// An animation of an app’s views, limited to changes in frame location and
// size, and to fade-in and fade-out effects.
func NSViewAnimationFromID(id objc.ID) NSViewAnimation {
	return NSViewAnimation{NSAnimation: NSAnimationFromID(id)}
}
// NOTE: NSViewAnimation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSViewAnimation] class.
//
// # Initializing an NSViewAnimation object
//
//   - [INSViewAnimation.InitWithViewAnimations]: Returns an [NSViewAnimation] object initialized with the supplied information.
//
// # Getting and setting view-animation dictionaries
//
//   - [INSViewAnimation.ViewAnimations]: The dictionaries defining the objects to animate.
//   - [INSViewAnimation.SetViewAnimations]
//
// See: https://developer.apple.com/documentation/AppKit/NSViewAnimation
type INSViewAnimation interface {
	INSAnimation

	// Topic: Initializing an NSViewAnimation object

	// Returns an [NSViewAnimation] object initialized with the supplied information.
	InitWithViewAnimations(viewAnimations foundation.INSDictionary) NSViewAnimation

	// Topic: Getting and setting view-animation dictionaries

	// The dictionaries defining the objects to animate.
	ViewAnimations() foundation.INSDictionary
	SetViewAnimations(value foundation.INSDictionary)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (v NSViewAnimation) Init() NSViewAnimation {
	rv := objc.Send[NSViewAnimation](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v NSViewAnimation) Autorelease() NSViewAnimation {
	rv := objc.Send[NSViewAnimation](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSViewAnimation creates a new NSViewAnimation instance.
func NewNSViewAnimation() NSViewAnimation {
	class := getNSViewAnimationClass()
	rv := objc.Send[NSViewAnimation](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/init(coder:)
func NewViewAnimationWithCoder(coder foundation.INSCoder) NSViewAnimation {
	instance := getNSViewAnimationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSViewAnimationFromID(rv)
}


// Returns an [NSAnimation] object initialized with the specified duration and
// animation-curve values.
//
// duration: The number of seconds over which the animation occurs. Specifying a
// negative number raises an exception.
//
// animationCurve: An [NSAnimationCurve] constant that describes the relative speed of the
// animation over its course; if it is zero, the default curve
// ([NSAnimationEaseInOut]) is used.
//
// # Return Value
// 
// An initialized [NSAnimation] instance. Returns `nil` if the object could
// not be initialized.
//
// # Discussion
// 
// You can always later change the duration of an [NSAnimation] object by
// changing the [Duration] property, even while the animation is running. See
// “Constants” for descriptions of the NSAnimationCurve constants.
//
// See: https://developer.apple.com/documentation/AppKit/NSAnimation/init(duration:animationCurve:)
func NewViewAnimationWithDurationAnimationCurve(duration float64, animationCurve NSAnimationCurve) NSViewAnimation {
	instance := getNSViewAnimationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDuration:animationCurve:"), duration, animationCurve)
	return NSViewAnimationFromID(rv)
}


// Returns an [NSViewAnimation] object initialized with the supplied
// information.
//
// viewAnimations: An array of [NSDictionary] objects. Each dictionary specifies a view or
// window to animate and the effect to apply. `viewAnimations` can be `nil`,
// but you must later set the required array of dictionaries with
// [ViewAnimations] if you want to use the capabilities of the
// [NSViewAnimation] class. See`View Animation Dictionary Keys` for a
// description of valid keys and values for dictionaries in `viewAnimations`.
// //
// [NSDictionary]: https://developer.apple.com/documentation/Foundation/NSDictionary
//
// # Return Value
// 
// The created [NSViewAnimation] object or `nil` if there was a problem
// initializing the object.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewAnimation/init(viewAnimations:)
func NewViewAnimationWithViewAnimations(viewAnimations foundation.INSDictionary) NSViewAnimation {
	instance := getNSViewAnimationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithViewAnimations:"), viewAnimations)
	return NSViewAnimationFromID(rv)
}







// Returns an [NSViewAnimation] object initialized with the supplied
// information.
//
// viewAnimations: An array of [NSDictionary] objects. Each dictionary specifies a view or
// window to animate and the effect to apply. `viewAnimations` can be `nil`,
// but you must later set the required array of dictionaries with
// [ViewAnimations] if you want to use the capabilities of the
// [NSViewAnimation] class. See`View Animation Dictionary Keys` for a
// description of valid keys and values for dictionaries in `viewAnimations`.
// //
// [NSDictionary]: https://developer.apple.com/documentation/Foundation/NSDictionary
//
// # Return Value
// 
// The created [NSViewAnimation] object or `nil` if there was a problem
// initializing the object.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewAnimation/init(viewAnimations:)
func (v NSViewAnimation) InitWithViewAnimations(viewAnimations foundation.INSDictionary) NSViewAnimation {
	rv := objc.Send[NSViewAnimation](v.ID, objc.Sel("initWithViewAnimations:"), viewAnimations)
	return rv
}
func (v NSViewAnimation) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](v.ID, objc.Sel("encodeWithCoder:"), coder)
}












// The dictionaries defining the objects to animate.
//
// See: https://developer.apple.com/documentation/AppKit/NSViewAnimation/viewAnimations
func (v NSViewAnimation) ViewAnimations() foundation.INSDictionary {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("viewAnimations"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (v NSViewAnimation) SetViewAnimations(value foundation.INSDictionary) {
	objc.Send[struct{}](v.ID, objc.Sel("setViewAnimations:"), value)
}



























