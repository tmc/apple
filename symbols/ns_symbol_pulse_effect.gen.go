// Code generated from Apple documentation for Symbols. DO NOT EDIT.

package symbols

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSymbolPulseEffect] class.
var (
	_NSSymbolPulseEffectClass     NSSymbolPulseEffectClass
	_NSSymbolPulseEffectClassOnce sync.Once
)

func getNSSymbolPulseEffectClass() NSSymbolPulseEffectClass {
	_NSSymbolPulseEffectClassOnce.Do(func() {
		_NSSymbolPulseEffectClass = NSSymbolPulseEffectClass{class: objc.GetClass("NSSymbolPulseEffect")}
	})
	return _NSSymbolPulseEffectClass
}

// GetNSSymbolPulseEffectClass returns the class object for NSSymbolPulseEffect.
func GetNSSymbolPulseEffectClass() NSSymbolPulseEffectClass {
	return getNSSymbolPulseEffectClass()
}

type NSSymbolPulseEffectClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSymbolPulseEffectClass) Alloc() NSSymbolPulseEffect {
	rv := objc.Send[NSSymbolPulseEffect](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A type that fades the opacity of some or all layers in a symbol-based
// image.
//
// # Overview
// 
// A pulse animation applies an opacity ramp to the layers in a symbol. You
// can choose to animate only layers marked as “always-pulses” or all
// layers simultaneously. Participating layers reduce their opacity to a
// minimum value before returning to fully opaque.
//
// # Determining effect scope
//
//   - [NSSymbolPulseEffect.EffectWithByLayer]: A copy of the effect requesting an animation that pulses only the layers marked to always pulse.
//   - [NSSymbolPulseEffect.EffectWithWholeSymbol]: A copy of the effect requesting an animation that pulses all layers simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolPulseEffect
type NSSymbolPulseEffect struct {
	NSSymbolEffect
}

// NSSymbolPulseEffectFromID constructs a [NSSymbolPulseEffect] from an objc.ID.
//
// A type that fades the opacity of some or all layers in a symbol-based
// image.
func NSSymbolPulseEffectFromID(id objc.ID) NSSymbolPulseEffect {
	return NSSymbolPulseEffect{NSSymbolEffect: NSSymbolEffectFromID(id)}
}
// Ensure NSSymbolPulseEffect implements INSSymbolPulseEffect.
var _ INSSymbolPulseEffect = NSSymbolPulseEffect{}





// An interface definition for the [NSSymbolPulseEffect] class.
//
// # Determining effect scope
//
//   - [INSSymbolPulseEffect.EffectWithByLayer]: A copy of the effect requesting an animation that pulses only the layers marked to always pulse.
//   - [INSSymbolPulseEffect.EffectWithWholeSymbol]: A copy of the effect requesting an animation that pulses all layers simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolPulseEffect
type INSSymbolPulseEffect interface {
	INSSymbolEffect

	// Topic: Determining effect scope

	// A copy of the effect requesting an animation that pulses only the layers marked to always pulse.
	EffectWithByLayer() INSSymbolPulseEffect
	// A copy of the effect requesting an animation that pulses all layers simultaneously.
	EffectWithWholeSymbol() INSSymbolPulseEffect
}





// Init initializes the instance.
func (s NSSymbolPulseEffect) Init() NSSymbolPulseEffect {
	rv := objc.Send[NSSymbolPulseEffect](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSymbolPulseEffect) Autorelease() NSSymbolPulseEffect {
	rv := objc.Send[NSSymbolPulseEffect](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSymbolPulseEffect creates a new NSSymbolPulseEffect instance.
func NewNSSymbolPulseEffect() NSSymbolPulseEffect {
	class := getNSSymbolPulseEffectClass()
	rv := objc.Send[NSSymbolPulseEffect](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// A copy of the effect requesting an animation that pulses only the layers
// marked to always pulse.
//
// # Return Value
// 
// A copy of the effect options that pulses only the layers marked to always
// pulse.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolPulseEffect/effectWithByLayer
func (s NSSymbolPulseEffect) EffectWithByLayer() INSSymbolPulseEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithByLayer"))
	return NSSymbolPulseEffectFromID(rv)
}

// A copy of the effect requesting an animation that pulses all layers
// simultaneously.
//
// # Return Value
// 
// A copy of the effect options that pulses all layers simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolPulseEffect/effectWithWholeSymbol
func (s NSSymbolPulseEffect) EffectWithWholeSymbol() INSSymbolPulseEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithWholeSymbol"))
	return NSSymbolPulseEffectFromID(rv)
}





// The default pulse effect, determined by the system.
//
// # Return Value
// 
// A new instance of the pulse effect options.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolPulseEffect/effect
func (_NSSymbolPulseEffectClass NSSymbolPulseEffectClass) Effect() NSSymbolPulseEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolPulseEffectClass.class), objc.Sel("effect"))
	return NSSymbolPulseEffectFromID(rv)
}






















