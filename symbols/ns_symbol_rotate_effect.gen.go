// Code generated from Apple documentation for Symbols. DO NOT EDIT.

package symbols

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSymbolRotateEffect] class.
var (
	_NSSymbolRotateEffectClass     NSSymbolRotateEffectClass
	_NSSymbolRotateEffectClassOnce sync.Once
)

func getNSSymbolRotateEffectClass() NSSymbolRotateEffectClass {
	_NSSymbolRotateEffectClassOnce.Do(func() {
		_NSSymbolRotateEffectClass = NSSymbolRotateEffectClass{class: objc.GetClass("NSSymbolRotateEffect")}
	})
	return _NSSymbolRotateEffectClass
}

// GetNSSymbolRotateEffectClass returns the class object for NSSymbolRotateEffect.
func GetNSSymbolRotateEffectClass() NSSymbolRotateEffectClass {
	return getNSSymbolRotateEffectClass()
}

type NSSymbolRotateEffectClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSymbolRotateEffectClass) Alloc() NSSymbolRotateEffect {
	rv := objc.Send[NSSymbolRotateEffect](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A symbol effect that applies the Rotate animation to symbol images.
//
// # Overview
// 
// The Rotate animation rotates parts of a symbol around a symbol-provided
// anchor point.
//
// # Instance Methods
//
//   - [NSSymbolRotateEffect.EffectWithByLayer]: Returns a copy of the effect that animates incrementally, by layer.
//   - [NSSymbolRotateEffect.EffectWithWholeSymbol]: Returns a copy of the effect that animates all layers of the symbol simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolRotateEffect
type NSSymbolRotateEffect struct {
	NSSymbolEffect
}

// NSSymbolRotateEffectFromID constructs a [NSSymbolRotateEffect] from an objc.ID.
//
// A symbol effect that applies the Rotate animation to symbol images.
func NSSymbolRotateEffectFromID(id objc.ID) NSSymbolRotateEffect {
	return NSSymbolRotateEffect{NSSymbolEffect: NSSymbolEffectFromID(id)}
}
// Ensure NSSymbolRotateEffect implements INSSymbolRotateEffect.
var _ INSSymbolRotateEffect = NSSymbolRotateEffect{}





// An interface definition for the [NSSymbolRotateEffect] class.
//
// # Instance Methods
//
//   - [INSSymbolRotateEffect.EffectWithByLayer]: Returns a copy of the effect that animates incrementally, by layer.
//   - [INSSymbolRotateEffect.EffectWithWholeSymbol]: Returns a copy of the effect that animates all layers of the symbol simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolRotateEffect
type INSSymbolRotateEffect interface {
	INSSymbolEffect

	// Topic: Instance Methods

	// Returns a copy of the effect that animates incrementally, by layer.
	EffectWithByLayer() INSSymbolRotateEffect
	// Returns a copy of the effect that animates all layers of the symbol simultaneously.
	EffectWithWholeSymbol() INSSymbolRotateEffect
}





// Init initializes the instance.
func (s NSSymbolRotateEffect) Init() NSSymbolRotateEffect {
	rv := objc.Send[NSSymbolRotateEffect](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSymbolRotateEffect) Autorelease() NSSymbolRotateEffect {
	rv := objc.Send[NSSymbolRotateEffect](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSymbolRotateEffect creates a new NSSymbolRotateEffect instance.
func NewNSSymbolRotateEffect() NSSymbolRotateEffect {
	class := getNSSymbolRotateEffectClass()
	rv := objc.Send[NSSymbolRotateEffect](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Returns a copy of the effect that animates incrementally, by layer.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolRotateEffect/effectWithByLayer
func (s NSSymbolRotateEffect) EffectWithByLayer() INSSymbolRotateEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithByLayer"))
	return NSSymbolRotateEffectFromID(rv)
}

// Returns a copy of the effect that animates all layers of the symbol
// simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolRotateEffect/effectWithWholeSymbol
func (s NSSymbolRotateEffect) EffectWithWholeSymbol() INSSymbolRotateEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithWholeSymbol"))
	return NSSymbolRotateEffectFromID(rv)
}





// The default rotate effect, determined by the system.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolRotateEffect/effect
func (_NSSymbolRotateEffectClass NSSymbolRotateEffectClass) Effect() NSSymbolRotateEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolRotateEffectClass.class), objc.Sel("effect"))
	return NSSymbolRotateEffectFromID(rv)
}

// Convenience initializer for a rotate effect that rotates clockwise.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolRotateEffect/rotateClockwiseEffect
func (_NSSymbolRotateEffectClass NSSymbolRotateEffectClass) RotateClockwiseEffect() NSSymbolRotateEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolRotateEffectClass.class), objc.Sel("rotateClockwiseEffect"))
	return NSSymbolRotateEffectFromID(rv)
}

// Convenience initializer for a rotate effect that rotates counter-clockwise.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolRotateEffect/rotateCounterClockwiseEffect
func (_NSSymbolRotateEffectClass NSSymbolRotateEffectClass) RotateCounterClockwiseEffect() NSSymbolRotateEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolRotateEffectClass.class), objc.Sel("rotateCounterClockwiseEffect"))
	return NSSymbolRotateEffectFromID(rv)
}






















