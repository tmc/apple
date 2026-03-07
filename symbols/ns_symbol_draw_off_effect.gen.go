// Code generated from Apple documentation for Symbols. DO NOT EDIT.

package symbols

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSymbolDrawOffEffect] class.
var (
	_NSSymbolDrawOffEffectClass     NSSymbolDrawOffEffectClass
	_NSSymbolDrawOffEffectClassOnce sync.Once
)

func getNSSymbolDrawOffEffectClass() NSSymbolDrawOffEffectClass {
	_NSSymbolDrawOffEffectClassOnce.Do(func() {
		_NSSymbolDrawOffEffectClass = NSSymbolDrawOffEffectClass{class: objc.GetClass("NSSymbolDrawOffEffect")}
	})
	return _NSSymbolDrawOffEffectClass
}

// GetNSSymbolDrawOffEffectClass returns the class object for NSSymbolDrawOffEffect.
func GetNSSymbolDrawOffEffectClass() NSSymbolDrawOffEffectClass {
	return getNSSymbolDrawOffEffectClass()
}

type NSSymbolDrawOffEffectClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSymbolDrawOffEffectClass) Alloc() NSSymbolDrawOffEffect {
	rv := objc.Send[NSSymbolDrawOffEffect](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A symbol effect that applies the DrawOff animation to symbol images.
//
// # Overview
// 
// The DrawOff animation makes the symbol hidden either as a whole, or one
// motion group at a time, animating parts of the symbol with draw data.
//
// # Instance Methods
//
//   - [NSSymbolDrawOffEffect.EffectWithByLayer]: Returns a copy of the effect requesting an animation that applies separately to each motion group.
//   - [NSSymbolDrawOffEffect.EffectWithIndividually]: Returns a copy of the effect requesting an animation that applies separately to each motion group, where only one motion group is active at a time.
//   - [NSSymbolDrawOffEffect.EffectWithNonReversed]: Returns a copy of the effect that only animates forwards. This cancels the reversed variant.
//   - [NSSymbolDrawOffEffect.EffectWithReversed]: Returns a copy of the effect that animates in reverse. This cancels the nonReversed variant.
//   - [NSSymbolDrawOffEffect.EffectWithWholeSymbol]: Returns a copy of the effect requesting an animation that applies to all motion groups simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDrawOffEffect
type NSSymbolDrawOffEffect struct {
	NSSymbolEffect
}

// NSSymbolDrawOffEffectFromID constructs a [NSSymbolDrawOffEffect] from an objc.ID.
//
// A symbol effect that applies the DrawOff animation to symbol images.
func NSSymbolDrawOffEffectFromID(id objc.ID) NSSymbolDrawOffEffect {
	return NSSymbolDrawOffEffect{NSSymbolEffect: NSSymbolEffectFromID(id)}
}
// Ensure NSSymbolDrawOffEffect implements INSSymbolDrawOffEffect.
var _ INSSymbolDrawOffEffect = NSSymbolDrawOffEffect{}





// An interface definition for the [NSSymbolDrawOffEffect] class.
//
// # Instance Methods
//
//   - [INSSymbolDrawOffEffect.EffectWithByLayer]: Returns a copy of the effect requesting an animation that applies separately to each motion group.
//   - [INSSymbolDrawOffEffect.EffectWithIndividually]: Returns a copy of the effect requesting an animation that applies separately to each motion group, where only one motion group is active at a time.
//   - [INSSymbolDrawOffEffect.EffectWithNonReversed]: Returns a copy of the effect that only animates forwards. This cancels the reversed variant.
//   - [INSSymbolDrawOffEffect.EffectWithReversed]: Returns a copy of the effect that animates in reverse. This cancels the nonReversed variant.
//   - [INSSymbolDrawOffEffect.EffectWithWholeSymbol]: Returns a copy of the effect requesting an animation that applies to all motion groups simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDrawOffEffect
type INSSymbolDrawOffEffect interface {
	INSSymbolEffect

	// Topic: Instance Methods

	// Returns a copy of the effect requesting an animation that applies separately to each motion group.
	EffectWithByLayer() INSSymbolDrawOffEffect
	// Returns a copy of the effect requesting an animation that applies separately to each motion group, where only one motion group is active at a time.
	EffectWithIndividually() INSSymbolDrawOffEffect
	// Returns a copy of the effect that only animates forwards. This cancels the reversed variant.
	EffectWithNonReversed() INSSymbolDrawOffEffect
	// Returns a copy of the effect that animates in reverse. This cancels the nonReversed variant.
	EffectWithReversed() INSSymbolDrawOffEffect
	// Returns a copy of the effect requesting an animation that applies to all motion groups simultaneously.
	EffectWithWholeSymbol() INSSymbolDrawOffEffect
}





// Init initializes the instance.
func (s NSSymbolDrawOffEffect) Init() NSSymbolDrawOffEffect {
	rv := objc.Send[NSSymbolDrawOffEffect](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSymbolDrawOffEffect) Autorelease() NSSymbolDrawOffEffect {
	rv := objc.Send[NSSymbolDrawOffEffect](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSymbolDrawOffEffect creates a new NSSymbolDrawOffEffect instance.
func NewNSSymbolDrawOffEffect() NSSymbolDrawOffEffect {
	class := getNSSymbolDrawOffEffectClass()
	rv := objc.Send[NSSymbolDrawOffEffect](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Returns a copy of the effect requesting an animation that applies
// separately to each motion group.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDrawOffEffect/effectWithByLayer
func (s NSSymbolDrawOffEffect) EffectWithByLayer() INSSymbolDrawOffEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithByLayer"))
	return NSSymbolDrawOffEffectFromID(rv)
}

// Returns a copy of the effect requesting an animation that applies
// separately to each motion group, where only one motion group is active at a
// time.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDrawOffEffect/effectWithIndividually
func (s NSSymbolDrawOffEffect) EffectWithIndividually() INSSymbolDrawOffEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithIndividually"))
	return NSSymbolDrawOffEffectFromID(rv)
}

// Returns a copy of the effect that only animates forwards. This cancels the
// reversed variant.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDrawOffEffect/effectWithNonReversed
func (s NSSymbolDrawOffEffect) EffectWithNonReversed() INSSymbolDrawOffEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithNonReversed"))
	return NSSymbolDrawOffEffectFromID(rv)
}

// Returns a copy of the effect that animates in reverse. This cancels the
// nonReversed variant.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDrawOffEffect/effectWithReversed
func (s NSSymbolDrawOffEffect) EffectWithReversed() INSSymbolDrawOffEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithReversed"))
	return NSSymbolDrawOffEffectFromID(rv)
}

// Returns a copy of the effect requesting an animation that applies to all
// motion groups simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDrawOffEffect/effectWithWholeSymbol
func (s NSSymbolDrawOffEffect) EffectWithWholeSymbol() INSSymbolDrawOffEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithWholeSymbol"))
	return NSSymbolDrawOffEffectFromID(rv)
}





// The default draw off effect, determined by the system.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDrawOffEffect/effect
func (_NSSymbolDrawOffEffectClass NSSymbolDrawOffEffectClass) Effect() NSSymbolDrawOffEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolDrawOffEffectClass.class), objc.Sel("effect"))
	return NSSymbolDrawOffEffectFromID(rv)
}






















