// Code generated from Apple documentation for Symbols. DO NOT EDIT.

package symbols

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSymbolAppearEffect] class.
var (
	_NSSymbolAppearEffectClass     NSSymbolAppearEffectClass
	_NSSymbolAppearEffectClassOnce sync.Once
)

func getNSSymbolAppearEffectClass() NSSymbolAppearEffectClass {
	_NSSymbolAppearEffectClassOnce.Do(func() {
		_NSSymbolAppearEffectClass = NSSymbolAppearEffectClass{class: objc.GetClass("NSSymbolAppearEffect")}
	})
	return _NSSymbolAppearEffectClass
}

// GetNSSymbolAppearEffectClass returns the class object for NSSymbolAppearEffect.
func GetNSSymbolAppearEffectClass() NSSymbolAppearEffectClass {
	return getNSSymbolAppearEffectClass()
}

type NSSymbolAppearEffectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSymbolAppearEffectClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSymbolAppearEffectClass) Alloc() NSSymbolAppearEffect {
	rv := objc.Send[NSSymbolAppearEffect](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A type that makes the layers of a symbol-based image appear separately or
// as a whole.
//
// # Overview
// 
// An appear transition causes a symbol to become visible using a scaling
// animation. You can choose to scale the image up or down and to animate the
// symbol by individual layers or as a whole.
//
// # Determining effect scope
//
//   - [NSSymbolAppearEffect.EffectWithByLayer]: An effect that makes each layer appear separately.
//   - [NSSymbolAppearEffect.EffectWithWholeSymbol]: An effect that makes all layers appear simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolAppearEffect
type NSSymbolAppearEffect struct {
	NSSymbolEffect
}

// NSSymbolAppearEffectFromID constructs a [NSSymbolAppearEffect] from an objc.ID.
//
// A type that makes the layers of a symbol-based image appear separately or
// as a whole.
func NSSymbolAppearEffectFromID(id objc.ID) NSSymbolAppearEffect {
	return NSSymbolAppearEffect{NSSymbolEffect: NSSymbolEffectFromID(id)}
}
// Ensure NSSymbolAppearEffect implements INSSymbolAppearEffect.
var _ INSSymbolAppearEffect = NSSymbolAppearEffect{}

// An interface definition for the [NSSymbolAppearEffect] class.
//
// # Determining effect scope
//
//   - [INSSymbolAppearEffect.EffectWithByLayer]: An effect that makes each layer appear separately.
//   - [INSSymbolAppearEffect.EffectWithWholeSymbol]: An effect that makes all layers appear simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolAppearEffect
type INSSymbolAppearEffect interface {
	INSSymbolEffect

	// Topic: Determining effect scope

	// An effect that makes each layer appear separately.
	EffectWithByLayer() INSSymbolAppearEffect
	// An effect that makes all layers appear simultaneously.
	EffectWithWholeSymbol() INSSymbolAppearEffect
}

// Init initializes the instance.
func (s NSSymbolAppearEffect) Init() NSSymbolAppearEffect {
	rv := objc.Send[NSSymbolAppearEffect](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSymbolAppearEffect) Autorelease() NSSymbolAppearEffect {
	rv := objc.Send[NSSymbolAppearEffect](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSymbolAppearEffect creates a new NSSymbolAppearEffect instance.
func NewNSSymbolAppearEffect() NSSymbolAppearEffect {
	class := getNSSymbolAppearEffectClass()
	rv := objc.Send[NSSymbolAppearEffect](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An effect that makes each layer appear separately.
//
// # Return Value
// 
// An effect that makes each layer appear separately.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolAppearEffect/effectWithByLayer
func (s NSSymbolAppearEffect) EffectWithByLayer() INSSymbolAppearEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithByLayer"))
	return NSSymbolAppearEffectFromID(rv)
}
// An effect that makes all layers appear simultaneously.
//
// # Return Value
// 
// A copy of the effect options that makes all layers appear simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolAppearEffect/effectWithWholeSymbol
func (s NSSymbolAppearEffect) EffectWithWholeSymbol() INSSymbolAppearEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithWholeSymbol"))
	return NSSymbolAppearEffectFromID(rv)
}

// An effect that makes the symbol scale down as it appears.
//
// # Return Value
// 
// A new instance of the appear effect options that makes the symbol scale
// down as it appears.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolAppearEffect/appearDownEffect
func (_NSSymbolAppearEffectClass NSSymbolAppearEffectClass) AppearDownEffect() NSSymbolAppearEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolAppearEffectClass.class), objc.Sel("appearDownEffect"))
	return NSSymbolAppearEffectFromID(rv)
}
// An effect that makes the symbol scale up as it appears.
//
// # Return Value
// 
// A new instance of the appear effect options that makes the symbol scale up
// as it appears.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolAppearEffect/appearUpEffect
func (_NSSymbolAppearEffectClass NSSymbolAppearEffectClass) AppearUpEffect() NSSymbolAppearEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolAppearEffectClass.class), objc.Sel("appearUpEffect"))
	return NSSymbolAppearEffectFromID(rv)
}
// An animation that makes the layers of a symbol-based image appear
// separately or as a whole.
//
// # Return Value
// 
// A new instance of the appear effect options.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolAppearEffect/effect
func (_NSSymbolAppearEffectClass NSSymbolAppearEffectClass) Effect() NSSymbolAppearEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolAppearEffectClass.class), objc.Sel("effect"))
	return NSSymbolAppearEffectFromID(rv)
}

