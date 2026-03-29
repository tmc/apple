// Code generated from Apple documentation for Symbols. DO NOT EDIT.

package symbols

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSymbolBreatheEffect] class.
var (
	_NSSymbolBreatheEffectClass     NSSymbolBreatheEffectClass
	_NSSymbolBreatheEffectClassOnce sync.Once
)

func getNSSymbolBreatheEffectClass() NSSymbolBreatheEffectClass {
	_NSSymbolBreatheEffectClassOnce.Do(func() {
		_NSSymbolBreatheEffectClass = NSSymbolBreatheEffectClass{class: objc.GetClass("NSSymbolBreatheEffect")}
	})
	return _NSSymbolBreatheEffectClass
}

// GetNSSymbolBreatheEffectClass returns the class object for NSSymbolBreatheEffect.
func GetNSSymbolBreatheEffectClass() NSSymbolBreatheEffectClass {
	return getNSSymbolBreatheEffectClass()
}

type NSSymbolBreatheEffectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSymbolBreatheEffectClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSymbolBreatheEffectClass) Alloc() NSSymbolBreatheEffect {
	rv := objc.Send[NSSymbolBreatheEffect](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A symbol effect that applies the Breathe animation to symbol images.
//
// # Overview
// 
// The Breathe animation smoothly scales a symbol up and down.
//
// # Instance Methods
//
//   - [NSSymbolBreatheEffect.EffectWithByLayer]: Returns a copy of the effect that animates incrementally, by layer.
//   - [NSSymbolBreatheEffect.EffectWithWholeSymbol]: Returns a copy of the effect that animates all layers of the symbol simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolBreatheEffect
type NSSymbolBreatheEffect struct {
	NSSymbolEffect
}

// NSSymbolBreatheEffectFromID constructs a [NSSymbolBreatheEffect] from an objc.ID.
//
// A symbol effect that applies the Breathe animation to symbol images.
func NSSymbolBreatheEffectFromID(id objc.ID) NSSymbolBreatheEffect {
	return NSSymbolBreatheEffect{NSSymbolEffect: NSSymbolEffectFromID(id)}
}
// Ensure NSSymbolBreatheEffect implements INSSymbolBreatheEffect.
var _ INSSymbolBreatheEffect = NSSymbolBreatheEffect{}

// An interface definition for the [NSSymbolBreatheEffect] class.
//
// # Instance Methods
//
//   - [INSSymbolBreatheEffect.EffectWithByLayer]: Returns a copy of the effect that animates incrementally, by layer.
//   - [INSSymbolBreatheEffect.EffectWithWholeSymbol]: Returns a copy of the effect that animates all layers of the symbol simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolBreatheEffect
type INSSymbolBreatheEffect interface {
	INSSymbolEffect

	// Topic: Instance Methods

	// Returns a copy of the effect that animates incrementally, by layer.
	EffectWithByLayer() INSSymbolBreatheEffect
	// Returns a copy of the effect that animates all layers of the symbol simultaneously.
	EffectWithWholeSymbol() INSSymbolBreatheEffect
}

// Init initializes the instance.
func (s NSSymbolBreatheEffect) Init() NSSymbolBreatheEffect {
	rv := objc.Send[NSSymbolBreatheEffect](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSymbolBreatheEffect) Autorelease() NSSymbolBreatheEffect {
	rv := objc.Send[NSSymbolBreatheEffect](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSymbolBreatheEffect creates a new NSSymbolBreatheEffect instance.
func NewNSSymbolBreatheEffect() NSSymbolBreatheEffect {
	class := getNSSymbolBreatheEffectClass()
	rv := objc.Send[NSSymbolBreatheEffect](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a copy of the effect that animates incrementally, by layer.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolBreatheEffect/effectWithByLayer
func (s NSSymbolBreatheEffect) EffectWithByLayer() INSSymbolBreatheEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithByLayer"))
	return NSSymbolBreatheEffectFromID(rv)
}
// Returns a copy of the effect that animates all layers of the symbol
// simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolBreatheEffect/effectWithWholeSymbol
func (s NSSymbolBreatheEffect) EffectWithWholeSymbol() INSSymbolBreatheEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithWholeSymbol"))
	return NSSymbolBreatheEffectFromID(rv)
}

// Convenience initializer for a breathe effect that makes the symbol breathe
// with no other styling.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolBreatheEffect/breathePlainEffect
func (_NSSymbolBreatheEffectClass NSSymbolBreatheEffectClass) BreathePlainEffect() NSSymbolBreatheEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolBreatheEffectClass.class), objc.Sel("breathePlainEffect"))
	return NSSymbolBreatheEffectFromID(rv)
}
// Convenience initializer for a breathe effect that pulses layers as they
// breathe.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolBreatheEffect/breathePulseEffect
func (_NSSymbolBreatheEffectClass NSSymbolBreatheEffectClass) BreathePulseEffect() NSSymbolBreatheEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolBreatheEffectClass.class), objc.Sel("breathePulseEffect"))
	return NSSymbolBreatheEffectFromID(rv)
}
// The default breathe effect, determined by the system.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolBreatheEffect/effect
func (_NSSymbolBreatheEffectClass NSSymbolBreatheEffectClass) Effect() NSSymbolBreatheEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolBreatheEffectClass.class), objc.Sel("effect"))
	return NSSymbolBreatheEffectFromID(rv)
}

