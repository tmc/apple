// Code generated from Apple documentation for Symbols. DO NOT EDIT.

package symbols

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSymbolBounceEffect] class.
var (
	_NSSymbolBounceEffectClass     NSSymbolBounceEffectClass
	_NSSymbolBounceEffectClassOnce sync.Once
)

func getNSSymbolBounceEffectClass() NSSymbolBounceEffectClass {
	_NSSymbolBounceEffectClassOnce.Do(func() {
		_NSSymbolBounceEffectClass = NSSymbolBounceEffectClass{class: objc.GetClass("NSSymbolBounceEffect")}
	})
	return _NSSymbolBounceEffectClass
}

// GetNSSymbolBounceEffectClass returns the class object for NSSymbolBounceEffect.
func GetNSSymbolBounceEffectClass() NSSymbolBounceEffectClass {
	return getNSSymbolBounceEffectClass()
}

type NSSymbolBounceEffectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSymbolBounceEffectClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSymbolBounceEffectClass) Alloc() NSSymbolBounceEffect {
	rv := objc.Send[NSSymbolBounceEffect](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A type that applies a transitory scaling effect, or bounce, to the layers
// in a symbol-based image separately or as a whole.
//
// # Overview
// 
// A bounce animation draws attention to a symbol by applying a brief scaling
// operation to the symbol’s layers. You can choose to scale the symbol up
// or down as it bounces.
//
// # Determining effect scope
//
//   - [NSSymbolBounceEffect.EffectWithByLayer]: An effect that bounces each layer separately.
//   - [NSSymbolBounceEffect.EffectWithWholeSymbol]: An effect that bounces all layers simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolBounceEffect
type NSSymbolBounceEffect struct {
	NSSymbolEffect
}

// NSSymbolBounceEffectFromID constructs a [NSSymbolBounceEffect] from an objc.ID.
//
// A type that applies a transitory scaling effect, or bounce, to the layers
// in a symbol-based image separately or as a whole.
func NSSymbolBounceEffectFromID(id objc.ID) NSSymbolBounceEffect {
	return NSSymbolBounceEffect{NSSymbolEffect: NSSymbolEffectFromID(id)}
}
// Ensure NSSymbolBounceEffect implements INSSymbolBounceEffect.
var _ INSSymbolBounceEffect = NSSymbolBounceEffect{}

// An interface definition for the [NSSymbolBounceEffect] class.
//
// # Determining effect scope
//
//   - [INSSymbolBounceEffect.EffectWithByLayer]: An effect that bounces each layer separately.
//   - [INSSymbolBounceEffect.EffectWithWholeSymbol]: An effect that bounces all layers simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolBounceEffect
type INSSymbolBounceEffect interface {
	INSSymbolEffect

	// Topic: Determining effect scope

	// An effect that bounces each layer separately.
	EffectWithByLayer() INSSymbolBounceEffect
	// An effect that bounces all layers simultaneously.
	EffectWithWholeSymbol() INSSymbolBounceEffect
}

// Init initializes the instance.
func (s NSSymbolBounceEffect) Init() NSSymbolBounceEffect {
	rv := objc.Send[NSSymbolBounceEffect](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSymbolBounceEffect) Autorelease() NSSymbolBounceEffect {
	rv := objc.Send[NSSymbolBounceEffect](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSymbolBounceEffect creates a new NSSymbolBounceEffect instance.
func NewNSSymbolBounceEffect() NSSymbolBounceEffect {
	class := getNSSymbolBounceEffectClass()
	rv := objc.Send[NSSymbolBounceEffect](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An effect that bounces each layer separately.
//
// # Return Value
// 
// A copy of the effect options that bounces each layer separately.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolBounceEffect/effectWithByLayer
func (s NSSymbolBounceEffect) EffectWithByLayer() INSSymbolBounceEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithByLayer"))
	return NSSymbolBounceEffectFromID(rv)
}
// An effect that bounces all layers simultaneously.
//
// # Return Value
// 
// A copy of the effect options that bounces all layers simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolBounceEffect/effectWithWholeSymbol
func (s NSSymbolBounceEffect) EffectWithWholeSymbol() INSSymbolBounceEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithWholeSymbol"))
	return NSSymbolBounceEffectFromID(rv)
}

// An effect that bounces the symbol downward.
//
// # Return Value
// 
// A new instance of the bounce effect options that bounces the symbol
// downward.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolBounceEffect/bounceDownEffect
func (_NSSymbolBounceEffectClass NSSymbolBounceEffectClass) BounceDownEffect() NSSymbolBounceEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolBounceEffectClass.class), objc.Sel("bounceDownEffect"))
	return NSSymbolBounceEffectFromID(rv)
}
// An effect that bounces the symbol upward.
//
// # Return Value
// 
// A new instance of the bounce effect options that bounces the symbol upward.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolBounceEffect/bounceUpEffect
func (_NSSymbolBounceEffectClass NSSymbolBounceEffectClass) BounceUpEffect() NSSymbolBounceEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolBounceEffectClass.class), objc.Sel("bounceUpEffect"))
	return NSSymbolBounceEffectFromID(rv)
}
// An animation that applies a transitory scaling effect, or bounce, to the
// layers in a symbol-based image separately or as a whole.
//
// # Return Value
// 
// A new instance of the bounce effect options.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolBounceEffect/effect
func (_NSSymbolBounceEffectClass NSSymbolBounceEffectClass) Effect() NSSymbolBounceEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolBounceEffectClass.class), objc.Sel("effect"))
	return NSSymbolBounceEffectFromID(rv)
}

