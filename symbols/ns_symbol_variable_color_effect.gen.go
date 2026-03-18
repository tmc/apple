// Code generated from Apple documentation for Symbols. DO NOT EDIT.

package symbols

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSymbolVariableColorEffect] class.
var (
	_NSSymbolVariableColorEffectClass     NSSymbolVariableColorEffectClass
	_NSSymbolVariableColorEffectClassOnce sync.Once
)

func getNSSymbolVariableColorEffectClass() NSSymbolVariableColorEffectClass {
	_NSSymbolVariableColorEffectClassOnce.Do(func() {
		_NSSymbolVariableColorEffectClass = NSSymbolVariableColorEffectClass{class: objc.GetClass("NSSymbolVariableColorEffect")}
	})
	return _NSSymbolVariableColorEffectClass
}

// GetNSSymbolVariableColorEffectClass returns the class object for NSSymbolVariableColorEffect.
func GetNSSymbolVariableColorEffectClass() NSSymbolVariableColorEffectClass {
	return getNSSymbolVariableColorEffectClass()
}

type NSSymbolVariableColorEffectClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSymbolVariableColorEffectClass) Alloc() NSSymbolVariableColorEffect {
	rv := objc.Send[NSSymbolVariableColorEffect](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A type that replaces the opacity of variable layers in a symbol-based image
// in a repeatable sequence.
//
// # Overview
// 
// A variable color animation draws attention to a symbol by changing the
// opacity of the symbol’s layers. You can choose to apply the effect to
// layers either cumulatively or iteratively. For cumulative animations, each
// layer’s opacity remains changed until the end of the animation cycle. For
// iterative animations, each layer’s opacity changes briefly before
// returning to its original state.
//
// # Controlling fill style
//
//   - [NSSymbolVariableColorEffect.EffectWithCumulative]: An effect that enables each layer of a symbol-based image in sequence.
//   - [NSSymbolVariableColorEffect.EffectWithIterative]: An effect that momentarily enables each layer of a symbol-based image in sequence.
//
// # Changing playback style
//
//   - [NSSymbolVariableColorEffect.EffectWithNonReversing]: An effect that doesn’t reverse each time it repeats.
//   - [NSSymbolVariableColorEffect.EffectWithReversing]: An effect that reverses each time it repeats.
//
// # Affecting inactive layers
//
//   - [NSSymbolVariableColorEffect.EffectWithDimInactiveLayers]: An effect that dims inactive layers in a symbol-based image.
//   - [NSSymbolVariableColorEffect.EffectWithHideInactiveLayers]: An effect that hides inactive layers in a symbol-based image.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolVariableColorEffect
type NSSymbolVariableColorEffect struct {
	NSSymbolEffect
}

// NSSymbolVariableColorEffectFromID constructs a [NSSymbolVariableColorEffect] from an objc.ID.
//
// A type that replaces the opacity of variable layers in a symbol-based image
// in a repeatable sequence.
func NSSymbolVariableColorEffectFromID(id objc.ID) NSSymbolVariableColorEffect {
	return NSSymbolVariableColorEffect{NSSymbolEffect: NSSymbolEffectFromID(id)}
}
// Ensure NSSymbolVariableColorEffect implements INSSymbolVariableColorEffect.
var _ INSSymbolVariableColorEffect = NSSymbolVariableColorEffect{}

// An interface definition for the [NSSymbolVariableColorEffect] class.
//
// # Controlling fill style
//
//   - [INSSymbolVariableColorEffect.EffectWithCumulative]: An effect that enables each layer of a symbol-based image in sequence.
//   - [INSSymbolVariableColorEffect.EffectWithIterative]: An effect that momentarily enables each layer of a symbol-based image in sequence.
//
// # Changing playback style
//
//   - [INSSymbolVariableColorEffect.EffectWithNonReversing]: An effect that doesn’t reverse each time it repeats.
//   - [INSSymbolVariableColorEffect.EffectWithReversing]: An effect that reverses each time it repeats.
//
// # Affecting inactive layers
//
//   - [INSSymbolVariableColorEffect.EffectWithDimInactiveLayers]: An effect that dims inactive layers in a symbol-based image.
//   - [INSSymbolVariableColorEffect.EffectWithHideInactiveLayers]: An effect that hides inactive layers in a symbol-based image.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolVariableColorEffect
type INSSymbolVariableColorEffect interface {
	INSSymbolEffect

	// Topic: Controlling fill style

	// An effect that enables each layer of a symbol-based image in sequence.
	EffectWithCumulative() INSSymbolVariableColorEffect
	// An effect that momentarily enables each layer of a symbol-based image in sequence.
	EffectWithIterative() INSSymbolVariableColorEffect

	// Topic: Changing playback style

	// An effect that doesn’t reverse each time it repeats.
	EffectWithNonReversing() INSSymbolVariableColorEffect
	// An effect that reverses each time it repeats.
	EffectWithReversing() INSSymbolVariableColorEffect

	// Topic: Affecting inactive layers

	// An effect that dims inactive layers in a symbol-based image.
	EffectWithDimInactiveLayers() INSSymbolVariableColorEffect
	// An effect that hides inactive layers in a symbol-based image.
	EffectWithHideInactiveLayers() INSSymbolVariableColorEffect
}

// Init initializes the instance.
func (s NSSymbolVariableColorEffect) Init() NSSymbolVariableColorEffect {
	rv := objc.Send[NSSymbolVariableColorEffect](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSymbolVariableColorEffect) Autorelease() NSSymbolVariableColorEffect {
	rv := objc.Send[NSSymbolVariableColorEffect](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSymbolVariableColorEffect creates a new NSSymbolVariableColorEffect instance.
func NewNSSymbolVariableColorEffect() NSSymbolVariableColorEffect {
	class := getNSSymbolVariableColorEffectClass()
	rv := objc.Send[NSSymbolVariableColorEffect](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An effect that enables each layer of a symbol-based image in sequence.
//
// # Return Value
// 
// A copy of the symbol effect options that uses the `cumulative` animation.
//
// # Discussion
// 
// This effect enables each successive variable layer, and the layer remains
// enabled until the end of the animation cycle. This effect cancels the
// [iterative] variant.
//
// [iterative]: https://developer.apple.com/documentation/Symbols/VariableColorSymbolEffect/iterative
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolVariableColorEffect/effectWithCumulative
func (s NSSymbolVariableColorEffect) EffectWithCumulative() INSSymbolVariableColorEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithCumulative"))
	return NSSymbolVariableColorEffectFromID(rv)
}

// An effect that momentarily enables each layer of a symbol-based image in
// sequence.
//
// # Return Value
// 
// A copy of the symbol effect options that uses the `iterative` animation.
//
// # Discussion
// 
// This effect enables each successive variable layer for a short period of
// time, and then disables the layer until the animation cycle ends. This
// effect cancels the [cumulative] variant.
//
// [cumulative]: https://developer.apple.com/documentation/Symbols/VariableColorSymbolEffect/cumulative
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolVariableColorEffect/effectWithIterative
func (s NSSymbolVariableColorEffect) EffectWithIterative() INSSymbolVariableColorEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithIterative"))
	return NSSymbolVariableColorEffectFromID(rv)
}

// An effect that doesn’t reverse each time it repeats.
//
// # Return Value
// 
// A copy of the symbol effect options that doesn’t reverse each time it
// repeats.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolVariableColorEffect/effectWithNonReversing
func (s NSSymbolVariableColorEffect) EffectWithNonReversing() INSSymbolVariableColorEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithNonReversing"))
	return NSSymbolVariableColorEffectFromID(rv)
}

// An effect that reverses each time it repeats.
//
// # Return Value
// 
// A copy of the symbol effect options that reverses each time it repeats.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolVariableColorEffect/effectWithReversing
func (s NSSymbolVariableColorEffect) EffectWithReversing() INSSymbolVariableColorEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithReversing"))
	return NSSymbolVariableColorEffectFromID(rv)
}

// An effect that dims inactive layers in a symbol-based image.
//
// # Return Value
// 
// A copy of the symbol effect options that dims inactive layers.
//
// # Discussion
// 
// This effect draws inactive layers with reduced, but nonzero, opacity.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolVariableColorEffect/effectWithDimInactiveLayers
func (s NSSymbolVariableColorEffect) EffectWithDimInactiveLayers() INSSymbolVariableColorEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithDimInactiveLayers"))
	return NSSymbolVariableColorEffectFromID(rv)
}

// An effect that hides inactive layers in a symbol-based image.
//
// # Return Value
// 
// A copy of the symbol effect options that hides inactive layers.
//
// # Discussion
// 
// This effect hides inactive layers completely, rather than drawing them with
// reduced, but nonzero, opacity.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolVariableColorEffect/effectWithHideInactiveLayers
func (s NSSymbolVariableColorEffect) EffectWithHideInactiveLayers() INSSymbolVariableColorEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithHideInactiveLayers"))
	return NSSymbolVariableColorEffectFromID(rv)
}

// An animation that replaces the opacity of variable layers in a symbol-based
// image in a repeatable sequence.
//
// # Return Value
// 
// A new instance of the variable color effect options.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolVariableColorEffect/effect
func (_NSSymbolVariableColorEffectClass NSSymbolVariableColorEffectClass) Effect() NSSymbolVariableColorEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolVariableColorEffectClass.class), objc.Sel("effect"))
	return NSSymbolVariableColorEffectFromID(rv)
}

