// Code generated from Apple documentation for Symbols. DO NOT EDIT.

package symbols

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSymbolDisappearEffect] class.
var (
	_NSSymbolDisappearEffectClass     NSSymbolDisappearEffectClass
	_NSSymbolDisappearEffectClassOnce sync.Once
)

func getNSSymbolDisappearEffectClass() NSSymbolDisappearEffectClass {
	_NSSymbolDisappearEffectClassOnce.Do(func() {
		_NSSymbolDisappearEffectClass = NSSymbolDisappearEffectClass{class: objc.GetClass("NSSymbolDisappearEffect")}
	})
	return _NSSymbolDisappearEffectClass
}

// GetNSSymbolDisappearEffectClass returns the class object for NSSymbolDisappearEffect.
func GetNSSymbolDisappearEffectClass() NSSymbolDisappearEffectClass {
	return getNSSymbolDisappearEffectClass()
}

type NSSymbolDisappearEffectClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSymbolDisappearEffectClass) Alloc() NSSymbolDisappearEffect {
	rv := objc.Send[NSSymbolDisappearEffect](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A type that makes the layers of a symbol-based image disappear separately
// or as a whole.
//
// # Overview
// 
// A disappear transition causes a symbol to become invisible using a scaling
// animation. You can choose to scale the image up or down and to animate the
// symbol by individual layers or as a whole.
//
// # Determining effect scope
//
//   - [NSSymbolDisappearEffect.EffectWithByLayer]: An effect that makes each layer disappear separately.
//   - [NSSymbolDisappearEffect.EffectWithWholeSymbol]: An effect that makes all layers disappear simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDisappearEffect
type NSSymbolDisappearEffect struct {
	NSSymbolEffect
}

// NSSymbolDisappearEffectFromID constructs a [NSSymbolDisappearEffect] from an objc.ID.
//
// A type that makes the layers of a symbol-based image disappear separately
// or as a whole.
func NSSymbolDisappearEffectFromID(id objc.ID) NSSymbolDisappearEffect {
	return NSSymbolDisappearEffect{NSSymbolEffect: NSSymbolEffectFromID(id)}
}
// Ensure NSSymbolDisappearEffect implements INSSymbolDisappearEffect.
var _ INSSymbolDisappearEffect = NSSymbolDisappearEffect{}

// An interface definition for the [NSSymbolDisappearEffect] class.
//
// # Determining effect scope
//
//   - [INSSymbolDisappearEffect.EffectWithByLayer]: An effect that makes each layer disappear separately.
//   - [INSSymbolDisappearEffect.EffectWithWholeSymbol]: An effect that makes all layers disappear simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDisappearEffect
type INSSymbolDisappearEffect interface {
	INSSymbolEffect

	// Topic: Determining effect scope

	// An effect that makes each layer disappear separately.
	EffectWithByLayer() INSSymbolDisappearEffect
	// An effect that makes all layers disappear simultaneously.
	EffectWithWholeSymbol() INSSymbolDisappearEffect
}

// Init initializes the instance.
func (s NSSymbolDisappearEffect) Init() NSSymbolDisappearEffect {
	rv := objc.Send[NSSymbolDisappearEffect](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSymbolDisappearEffect) Autorelease() NSSymbolDisappearEffect {
	rv := objc.Send[NSSymbolDisappearEffect](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSymbolDisappearEffect creates a new NSSymbolDisappearEffect instance.
func NewNSSymbolDisappearEffect() NSSymbolDisappearEffect {
	class := getNSSymbolDisappearEffectClass()
	rv := objc.Send[NSSymbolDisappearEffect](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An effect that makes each layer disappear separately.
//
// # Return Value
// 
// A copy of the effect options that makes each layer disappear separately.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDisappearEffect/effectWithByLayer
func (s NSSymbolDisappearEffect) EffectWithByLayer() INSSymbolDisappearEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithByLayer"))
	return NSSymbolDisappearEffectFromID(rv)
}
// An effect that makes all layers disappear simultaneously.
//
// # Return Value
// 
// A copy of the effect options that makes all layers disappear
// simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDisappearEffect/effectWithWholeSymbol
func (s NSSymbolDisappearEffect) EffectWithWholeSymbol() INSSymbolDisappearEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithWholeSymbol"))
	return NSSymbolDisappearEffectFromID(rv)
}

// An effect that scales the symbol down as it disappears.
//
// # Return Value
// 
// A new instance of the disappear effect options that makes the symbol scale
// down as it disappears.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDisappearEffect/disappearDownEffect
func (_NSSymbolDisappearEffectClass NSSymbolDisappearEffectClass) DisappearDownEffect() NSSymbolDisappearEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolDisappearEffectClass.class), objc.Sel("disappearDownEffect"))
	return NSSymbolDisappearEffectFromID(rv)
}
// An effect that scales the symbol up as it disappears.
//
// # Return Value
// 
// A new instance of the disappear effect options that makes the symbol scale
// up as it disappears.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDisappearEffect/disappearUpEffect
func (_NSSymbolDisappearEffectClass NSSymbolDisappearEffectClass) DisappearUpEffect() NSSymbolDisappearEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolDisappearEffectClass.class), objc.Sel("disappearUpEffect"))
	return NSSymbolDisappearEffectFromID(rv)
}
// An animation that makes the layers of a symbol-based image disappear
// separately or as a whole.
//
// # Return Value
// 
// A new instance of the disappear effect options.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolDisappearEffect/effect
func (_NSSymbolDisappearEffectClass NSSymbolDisappearEffectClass) Effect() NSSymbolDisappearEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolDisappearEffectClass.class), objc.Sel("effect"))
	return NSSymbolDisappearEffectFromID(rv)
}

