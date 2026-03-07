// Code generated from Apple documentation for Symbols. DO NOT EDIT.

package symbols

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSymbolScaleEffect] class.
var (
	_NSSymbolScaleEffectClass     NSSymbolScaleEffectClass
	_NSSymbolScaleEffectClassOnce sync.Once
)

func getNSSymbolScaleEffectClass() NSSymbolScaleEffectClass {
	_NSSymbolScaleEffectClassOnce.Do(func() {
		_NSSymbolScaleEffectClass = NSSymbolScaleEffectClass{class: objc.GetClass("NSSymbolScaleEffect")}
	})
	return _NSSymbolScaleEffectClass
}

// GetNSSymbolScaleEffectClass returns the class object for NSSymbolScaleEffect.
func GetNSSymbolScaleEffectClass() NSSymbolScaleEffectClass {
	return getNSSymbolScaleEffectClass()
}

type NSSymbolScaleEffectClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSymbolScaleEffectClass) Alloc() NSSymbolScaleEffect {
	rv := objc.Send[NSSymbolScaleEffect](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A type that scales the layers in a symbol-based image separately or as a
// whole.
//
// # Overview
// 
// A scale animation draws attention to a symbol by changing the symbol’s
// scale indefinitely. You can choose to scale the symbol up or down.
//
// # Determining effect scope
//
//   - [NSSymbolScaleEffect.EffectWithByLayer]: An effect that scales each layer separately.
//   - [NSSymbolScaleEffect.EffectWithWholeSymbol]: An effect that scales all layers simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolScaleEffect
type NSSymbolScaleEffect struct {
	NSSymbolEffect
}

// NSSymbolScaleEffectFromID constructs a [NSSymbolScaleEffect] from an objc.ID.
//
// A type that scales the layers in a symbol-based image separately or as a
// whole.
func NSSymbolScaleEffectFromID(id objc.ID) NSSymbolScaleEffect {
	return NSSymbolScaleEffect{NSSymbolEffect: NSSymbolEffectFromID(id)}
}
// Ensure NSSymbolScaleEffect implements INSSymbolScaleEffect.
var _ INSSymbolScaleEffect = NSSymbolScaleEffect{}





// An interface definition for the [NSSymbolScaleEffect] class.
//
// # Determining effect scope
//
//   - [INSSymbolScaleEffect.EffectWithByLayer]: An effect that scales each layer separately.
//   - [INSSymbolScaleEffect.EffectWithWholeSymbol]: An effect that scales all layers simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolScaleEffect
type INSSymbolScaleEffect interface {
	INSSymbolEffect

	// Topic: Determining effect scope

	// An effect that scales each layer separately.
	EffectWithByLayer() INSSymbolScaleEffect
	// An effect that scales all layers simultaneously.
	EffectWithWholeSymbol() INSSymbolScaleEffect
}





// Init initializes the instance.
func (s NSSymbolScaleEffect) Init() NSSymbolScaleEffect {
	rv := objc.Send[NSSymbolScaleEffect](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSymbolScaleEffect) Autorelease() NSSymbolScaleEffect {
	rv := objc.Send[NSSymbolScaleEffect](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSymbolScaleEffect creates a new NSSymbolScaleEffect instance.
func NewNSSymbolScaleEffect() NSSymbolScaleEffect {
	class := getNSSymbolScaleEffectClass()
	rv := objc.Send[NSSymbolScaleEffect](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// An effect that scales each layer separately.
//
// # Return Value
// 
// A copy of the effect options that scales each layer separately.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolScaleEffect/effectWithByLayer
func (s NSSymbolScaleEffect) EffectWithByLayer() INSSymbolScaleEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithByLayer"))
	return NSSymbolScaleEffectFromID(rv)
}

// An effect that scales all layers simultaneously.
//
// # Return Value
// 
// A copy of the effect options that scales all layers simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolScaleEffect/effectWithWholeSymbol
func (s NSSymbolScaleEffect) EffectWithWholeSymbol() INSSymbolScaleEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithWholeSymbol"))
	return NSSymbolScaleEffectFromID(rv)
}





// An animation that scales the layers in a symbol-based image separately or
// as a whole.
//
// # Return Value
// 
// A new instance of the scale effect options.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolScaleEffect/effect
func (_NSSymbolScaleEffectClass NSSymbolScaleEffectClass) Effect() NSSymbolScaleEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolScaleEffectClass.class), objc.Sel("effect"))
	return NSSymbolScaleEffectFromID(rv)
}

// An effect that scales the symbol down.
//
// # Return Value
// 
// A new instance of the scale effect options that scales down the symbol.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolScaleEffect/scaleDownEffect
func (_NSSymbolScaleEffectClass NSSymbolScaleEffectClass) ScaleDownEffect() NSSymbolScaleEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolScaleEffectClass.class), objc.Sel("scaleDownEffect"))
	return NSSymbolScaleEffectFromID(rv)
}

// An effect that scales the symbol up.
//
// # Return Value
// 
// A new instance of the scale effect options that scales up the symbol.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolScaleEffect/scaleUpEffect
func (_NSSymbolScaleEffectClass NSSymbolScaleEffectClass) ScaleUpEffect() NSSymbolScaleEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolScaleEffectClass.class), objc.Sel("scaleUpEffect"))
	return NSSymbolScaleEffectFromID(rv)
}






















