// Code generated from Apple documentation for Symbols. DO NOT EDIT.

package symbols

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSymbolDrawOnEffect] class.
var (
	_NSSymbolDrawOnEffectClass     NSSymbolDrawOnEffectClass
	_NSSymbolDrawOnEffectClassOnce sync.Once
)

func getNSSymbolDrawOnEffectClass() NSSymbolDrawOnEffectClass {
	_NSSymbolDrawOnEffectClassOnce.Do(func() {
		_NSSymbolDrawOnEffectClass = NSSymbolDrawOnEffectClass{class: objc.GetClass("NSSymbolDrawOnEffect")}
	})
	return _NSSymbolDrawOnEffectClass
}

// GetNSSymbolDrawOnEffectClass returns the class object for NSSymbolDrawOnEffect.
func GetNSSymbolDrawOnEffectClass() NSSymbolDrawOnEffectClass {
	return getNSSymbolDrawOnEffectClass()
}

type NSSymbolDrawOnEffectClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSymbolDrawOnEffectClass) Alloc() NSSymbolDrawOnEffect {
	rv := objc.Send[NSSymbolDrawOnEffect](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A symbol effect that applies the DrawOn animation to symbol images.
//
// # Overview
// 
// The DrawOn animation makes the symbol visible either as a whole, or one
// motion group at a time, animating parts of the symbol with draw data.
//
// # Instance Methods
//
//   - [NSSymbolDrawOnEffect.EffectWithByLayer]: Returns a copy of the effect requesting an animation that applies separately to each motion group.
//   - [NSSymbolDrawOnEffect.EffectWithIndividually]: Returns a copy of the effect requesting an animation that applies separately to each motion group, where only one motion group is active at a time.
//   - [NSSymbolDrawOnEffect.EffectWithWholeSymbol]: Returns a copy of the effect requesting an animation that applies to all motion groups simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDrawOnEffect
type NSSymbolDrawOnEffect struct {
	NSSymbolEffect
}

// NSSymbolDrawOnEffectFromID constructs a [NSSymbolDrawOnEffect] from an objc.ID.
//
// A symbol effect that applies the DrawOn animation to symbol images.
func NSSymbolDrawOnEffectFromID(id objc.ID) NSSymbolDrawOnEffect {
	return NSSymbolDrawOnEffect{NSSymbolEffect: NSSymbolEffectFromID(id)}
}
// Ensure NSSymbolDrawOnEffect implements INSSymbolDrawOnEffect.
var _ INSSymbolDrawOnEffect = NSSymbolDrawOnEffect{}

// An interface definition for the [NSSymbolDrawOnEffect] class.
//
// # Instance Methods
//
//   - [INSSymbolDrawOnEffect.EffectWithByLayer]: Returns a copy of the effect requesting an animation that applies separately to each motion group.
//   - [INSSymbolDrawOnEffect.EffectWithIndividually]: Returns a copy of the effect requesting an animation that applies separately to each motion group, where only one motion group is active at a time.
//   - [INSSymbolDrawOnEffect.EffectWithWholeSymbol]: Returns a copy of the effect requesting an animation that applies to all motion groups simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDrawOnEffect
type INSSymbolDrawOnEffect interface {
	INSSymbolEffect

	// Topic: Instance Methods

	// Returns a copy of the effect requesting an animation that applies separately to each motion group.
	EffectWithByLayer() INSSymbolDrawOnEffect
	// Returns a copy of the effect requesting an animation that applies separately to each motion group, where only one motion group is active at a time.
	EffectWithIndividually() INSSymbolDrawOnEffect
	// Returns a copy of the effect requesting an animation that applies to all motion groups simultaneously.
	EffectWithWholeSymbol() INSSymbolDrawOnEffect
}

// Init initializes the instance.
func (s NSSymbolDrawOnEffect) Init() NSSymbolDrawOnEffect {
	rv := objc.Send[NSSymbolDrawOnEffect](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSymbolDrawOnEffect) Autorelease() NSSymbolDrawOnEffect {
	rv := objc.Send[NSSymbolDrawOnEffect](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSymbolDrawOnEffect creates a new NSSymbolDrawOnEffect instance.
func NewNSSymbolDrawOnEffect() NSSymbolDrawOnEffect {
	class := getNSSymbolDrawOnEffectClass()
	rv := objc.Send[NSSymbolDrawOnEffect](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a copy of the effect requesting an animation that applies
// separately to each motion group.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDrawOnEffect/effectWithByLayer
func (s NSSymbolDrawOnEffect) EffectWithByLayer() INSSymbolDrawOnEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithByLayer"))
	return NSSymbolDrawOnEffectFromID(rv)
}

// Returns a copy of the effect requesting an animation that applies
// separately to each motion group, where only one motion group is active at a
// time.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDrawOnEffect/effectWithIndividually
func (s NSSymbolDrawOnEffect) EffectWithIndividually() INSSymbolDrawOnEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithIndividually"))
	return NSSymbolDrawOnEffectFromID(rv)
}

// Returns a copy of the effect requesting an animation that applies to all
// motion groups simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDrawOnEffect/effectWithWholeSymbol
func (s NSSymbolDrawOnEffect) EffectWithWholeSymbol() INSSymbolDrawOnEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithWholeSymbol"))
	return NSSymbolDrawOnEffectFromID(rv)
}

// The default draw on effect, determined by the system.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDrawOnEffect/effect
func (_NSSymbolDrawOnEffectClass NSSymbolDrawOnEffectClass) Effect() NSSymbolDrawOnEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolDrawOnEffectClass.class), objc.Sel("effect"))
	return NSSymbolDrawOnEffectFromID(rv)
}

