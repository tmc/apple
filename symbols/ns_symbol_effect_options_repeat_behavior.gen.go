// Code generated from Apple documentation for Symbols. DO NOT EDIT.

package symbols

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSymbolEffectOptionsRepeatBehavior] class.
var (
	_NSSymbolEffectOptionsRepeatBehaviorClass     NSSymbolEffectOptionsRepeatBehaviorClass
	_NSSymbolEffectOptionsRepeatBehaviorClassOnce sync.Once
)

func getNSSymbolEffectOptionsRepeatBehaviorClass() NSSymbolEffectOptionsRepeatBehaviorClass {
	_NSSymbolEffectOptionsRepeatBehaviorClassOnce.Do(func() {
		_NSSymbolEffectOptionsRepeatBehaviorClass = NSSymbolEffectOptionsRepeatBehaviorClass{class: objc.GetClass("NSSymbolEffectOptionsRepeatBehavior")}
	})
	return _NSSymbolEffectOptionsRepeatBehaviorClass
}

// GetNSSymbolEffectOptionsRepeatBehaviorClass returns the class object for NSSymbolEffectOptionsRepeatBehavior.
func GetNSSymbolEffectOptionsRepeatBehaviorClass() NSSymbolEffectOptionsRepeatBehaviorClass {
	return getNSSymbolEffectOptionsRepeatBehaviorClass()
}

type NSSymbolEffectOptionsRepeatBehaviorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSymbolEffectOptionsRepeatBehaviorClass) Alloc() NSSymbolEffectOptionsRepeatBehavior {
	rv := objc.Send[NSSymbolEffectOptionsRepeatBehavior](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The behavior of repetition to use when a symbol effect is animating.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolEffectOptionsRepeatBehavior
type NSSymbolEffectOptionsRepeatBehavior struct {
	objectivec.Object
}

// NSSymbolEffectOptionsRepeatBehaviorFromID constructs a [NSSymbolEffectOptionsRepeatBehavior] from an objc.ID.
//
// The behavior of repetition to use when a symbol effect is animating.
func NSSymbolEffectOptionsRepeatBehaviorFromID(id objc.ID) NSSymbolEffectOptionsRepeatBehavior {
	return NSSymbolEffectOptionsRepeatBehavior{objectivec.Object{ID: id}}
}
// NOTE: NSSymbolEffectOptionsRepeatBehavior adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSymbolEffectOptionsRepeatBehavior] class.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolEffectOptionsRepeatBehavior
type INSSymbolEffectOptionsRepeatBehavior interface {
	objectivec.IObject

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s NSSymbolEffectOptionsRepeatBehavior) Init() NSSymbolEffectOptionsRepeatBehavior {
	rv := objc.Send[NSSymbolEffectOptionsRepeatBehavior](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSymbolEffectOptionsRepeatBehavior) Autorelease() NSSymbolEffectOptionsRepeatBehavior {
	rv := objc.Send[NSSymbolEffectOptionsRepeatBehavior](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSymbolEffectOptionsRepeatBehavior creates a new NSSymbolEffectOptionsRepeatBehavior instance.
func NewNSSymbolEffectOptionsRepeatBehavior() NSSymbolEffectOptionsRepeatBehavior {
	class := getNSSymbolEffectOptionsRepeatBehaviorClass()
	rv := objc.Send[NSSymbolEffectOptionsRepeatBehavior](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (s NSSymbolEffectOptionsRepeatBehavior) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Creates and returns a repeat behavior that prefers to repeat indefinitely,
// using continuous animations if available. Continuous animations have an
// intro, a body that runs as long as the effect is enabled, and an outro. If
// available these animations provide a smoother animation when an effect
// repeats indefinitely.
//
// # Return Value
// 
// A new behavior that prefers to repeat indefinitely with continuous
// animations.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolEffectOptionsRepeatBehavior/behaviorContinuous
func (_NSSymbolEffectOptionsRepeatBehaviorClass NSSymbolEffectOptionsRepeatBehaviorClass) BehaviorContinuous() NSSymbolEffectOptionsRepeatBehavior {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolEffectOptionsRepeatBehaviorClass.class), objc.Sel("behaviorContinuous"))
	return NSSymbolEffectOptionsRepeatBehaviorFromID(rv)
}

// Creates and returns a repeat behavior that prefers to repeat indefinitely
// using periodic animations. Periodic animations play the effect at regular
// intervals starting and stopping each time.
//
// # Return Value
// 
// A new behavior that prefers to repeat indefinitely using periodic
// animations.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolEffectOptionsRepeatBehavior/behaviorPeriodic
func (_NSSymbolEffectOptionsRepeatBehaviorClass NSSymbolEffectOptionsRepeatBehaviorClass) BehaviorPeriodic() NSSymbolEffectOptionsRepeatBehavior {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolEffectOptionsRepeatBehaviorClass.class), objc.Sel("behaviorPeriodic"))
	return NSSymbolEffectOptionsRepeatBehaviorFromID(rv)
}

// Creates and returns a repeat behavior with a preferred play count using
// periodic animations. Periodic animations play the effect at regular
// intervals starting and stopping each time.
//
// count: The preferred number of times to play the effect. Very large or small
// values may be clamped.
//
// # Return Value
// 
// A new behavior with the preferred play count using periodic animations.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolEffectOptionsRepeatBehavior/behaviorPeriodicWithCount:
func (_NSSymbolEffectOptionsRepeatBehaviorClass NSSymbolEffectOptionsRepeatBehaviorClass) BehaviorPeriodicWithCount(count int) NSSymbolEffectOptionsRepeatBehavior {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolEffectOptionsRepeatBehaviorClass.class), objc.Sel("behaviorPeriodicWithCount:"), count)
	return NSSymbolEffectOptionsRepeatBehaviorFromID(rv)
}

// Creates and returns a repeat behavior with a preferred play count and delay
// using periodic animations. Periodic animations play the effect at regular
// intervals starting and stopping each time.
//
// count: The preferred number of times to play the effect. Very large or small
// values may be clamped.
//
// delay: The preferred delay between repetitions, in seconds.
//
// # Return Value
// 
// A new behavior with the preferred play count and delay using periodic
// animations.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolEffectOptionsRepeatBehavior/behaviorPeriodicWithCount:delay:
func (_NSSymbolEffectOptionsRepeatBehaviorClass NSSymbolEffectOptionsRepeatBehaviorClass) BehaviorPeriodicWithCountDelay(count int, delay float64) NSSymbolEffectOptionsRepeatBehavior {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolEffectOptionsRepeatBehaviorClass.class), objc.Sel("behaviorPeriodicWithCount:delay:"), count, delay)
	return NSSymbolEffectOptionsRepeatBehaviorFromID(rv)
}

// Creates and returns a repeat behavior with a preferred repeat delay using
// periodic animations. Periodic animations play the effect at regular
// intervals starting and stopping each time.
//
// delay: The preferred delay between repetitions, in seconds.
//
// # Return Value
// 
// A new behavior that prefers to repeat indefinitely with a specified delay
// using periodic animations.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolEffectOptionsRepeatBehavior/behaviorPeriodicWithDelay:
func (_NSSymbolEffectOptionsRepeatBehaviorClass NSSymbolEffectOptionsRepeatBehaviorClass) BehaviorPeriodicWithDelay(delay float64) NSSymbolEffectOptionsRepeatBehavior {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolEffectOptionsRepeatBehaviorClass.class), objc.Sel("behaviorPeriodicWithDelay:"), delay)
	return NSSymbolEffectOptionsRepeatBehaviorFromID(rv)
}

