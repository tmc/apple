// Code generated from Apple documentation for Symbols. DO NOT EDIT.

package symbols

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSymbolWiggleEffect] class.
var (
	_NSSymbolWiggleEffectClass     NSSymbolWiggleEffectClass
	_NSSymbolWiggleEffectClassOnce sync.Once
)

func getNSSymbolWiggleEffectClass() NSSymbolWiggleEffectClass {
	_NSSymbolWiggleEffectClassOnce.Do(func() {
		_NSSymbolWiggleEffectClass = NSSymbolWiggleEffectClass{class: objc.GetClass("NSSymbolWiggleEffect")}
	})
	return _NSSymbolWiggleEffectClass
}

// GetNSSymbolWiggleEffectClass returns the class object for NSSymbolWiggleEffect.
func GetNSSymbolWiggleEffectClass() NSSymbolWiggleEffectClass {
	return getNSSymbolWiggleEffectClass()
}

type NSSymbolWiggleEffectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSymbolWiggleEffectClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSymbolWiggleEffectClass) Alloc() NSSymbolWiggleEffect {
	rv := objc.Send[NSSymbolWiggleEffect](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A symbol effect that applies the Wiggle animation to symbol images.
//
// # Overview
// 
// The Wiggle animation applies a transitory translation or rotation effect to
// the symbol.
//
// # Instance Methods
//
//   - [NSSymbolWiggleEffect.EffectWithByLayer]: Returns a copy of the effect that animates incrementally, by layer.
//   - [NSSymbolWiggleEffect.EffectWithWholeSymbol]: Returns a copy of the effect that animates all layers of the symbol simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolWiggleEffect
type NSSymbolWiggleEffect struct {
	NSSymbolEffect
}

// NSSymbolWiggleEffectFromID constructs a [NSSymbolWiggleEffect] from an objc.ID.
//
// A symbol effect that applies the Wiggle animation to symbol images.
func NSSymbolWiggleEffectFromID(id objc.ID) NSSymbolWiggleEffect {
	return NSSymbolWiggleEffect{NSSymbolEffect: NSSymbolEffectFromID(id)}
}
// Ensure NSSymbolWiggleEffect implements INSSymbolWiggleEffect.
var _ INSSymbolWiggleEffect = NSSymbolWiggleEffect{}

// An interface definition for the [NSSymbolWiggleEffect] class.
//
// # Instance Methods
//
//   - [INSSymbolWiggleEffect.EffectWithByLayer]: Returns a copy of the effect that animates incrementally, by layer.
//   - [INSSymbolWiggleEffect.EffectWithWholeSymbol]: Returns a copy of the effect that animates all layers of the symbol simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolWiggleEffect
type INSSymbolWiggleEffect interface {
	INSSymbolEffect

	// Topic: Instance Methods

	// Returns a copy of the effect that animates incrementally, by layer.
	EffectWithByLayer() INSSymbolWiggleEffect
	// Returns a copy of the effect that animates all layers of the symbol simultaneously.
	EffectWithWholeSymbol() INSSymbolWiggleEffect
}

// Init initializes the instance.
func (s NSSymbolWiggleEffect) Init() NSSymbolWiggleEffect {
	rv := objc.Send[NSSymbolWiggleEffect](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSymbolWiggleEffect) Autorelease() NSSymbolWiggleEffect {
	rv := objc.Send[NSSymbolWiggleEffect](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSymbolWiggleEffect creates a new NSSymbolWiggleEffect instance.
func NewNSSymbolWiggleEffect() NSSymbolWiggleEffect {
	class := getNSSymbolWiggleEffectClass()
	rv := objc.Send[NSSymbolWiggleEffect](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a copy of the effect that animates incrementally, by layer.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolWiggleEffect/effectWithByLayer
func (s NSSymbolWiggleEffect) EffectWithByLayer() INSSymbolWiggleEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithByLayer"))
	return NSSymbolWiggleEffectFromID(rv)
}
// Returns a copy of the effect that animates all layers of the symbol
// simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolWiggleEffect/effectWithWholeSymbol
func (s NSSymbolWiggleEffect) EffectWithWholeSymbol() INSSymbolWiggleEffect {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("effectWithWholeSymbol"))
	return NSSymbolWiggleEffectFromID(rv)
}

// The default wiggle effect, determined by the system.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolWiggleEffect/effect
func (_NSSymbolWiggleEffectClass NSSymbolWiggleEffectClass) Effect() NSSymbolWiggleEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolWiggleEffectClass.class), objc.Sel("effect"))
	return NSSymbolWiggleEffectFromID(rv)
}
// Convenience initializer for a wiggle effect that moves back and forth
// horizontally based on the current locale, starting by moving backward.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolWiggleEffect/wiggleBackwardEffect
func (_NSSymbolWiggleEffectClass NSSymbolWiggleEffectClass) WiggleBackwardEffect() NSSymbolWiggleEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolWiggleEffectClass.class), objc.Sel("wiggleBackwardEffect"))
	return NSSymbolWiggleEffectFromID(rv)
}
// Convenience initializer for a wiggle effect that rotates back and forth,
// starting by rotating clockwise.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolWiggleEffect/wiggleClockwiseEffect
func (_NSSymbolWiggleEffectClass NSSymbolWiggleEffectClass) WiggleClockwiseEffect() NSSymbolWiggleEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolWiggleEffectClass.class), objc.Sel("wiggleClockwiseEffect"))
	return NSSymbolWiggleEffectFromID(rv)
}
// Convenience initializer for a wiggle effect that rotates back and forth,
// starting by rotating counter-clockwise.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolWiggleEffect/wiggleCounterClockwiseEffect
func (_NSSymbolWiggleEffectClass NSSymbolWiggleEffectClass) WiggleCounterClockwiseEffect() NSSymbolWiggleEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolWiggleEffectClass.class), objc.Sel("wiggleCounterClockwiseEffect"))
	return NSSymbolWiggleEffectFromID(rv)
}
// Convenience initializer for a wiggle effect that moves back and forth along
// an axis, starting by moving toward a custom angle.
//
// # Discussion
// 
// The angle is in degrees moving clockwise from the positive x-axis.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolWiggleEffect/wiggleCustomAngleEffect:
func (_NSSymbolWiggleEffectClass NSSymbolWiggleEffectClass) WiggleCustomAngleEffect(angle float64) NSSymbolWiggleEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolWiggleEffectClass.class), objc.Sel("wiggleCustomAngleEffect:"), angle)
	return NSSymbolWiggleEffectFromID(rv)
}
// Convenience initializer for a wiggle effect that moves back and forth
// vertically, starting by moving down.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolWiggleEffect/wiggleDownEffect
func (_NSSymbolWiggleEffectClass NSSymbolWiggleEffectClass) WiggleDownEffect() NSSymbolWiggleEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolWiggleEffectClass.class), objc.Sel("wiggleDownEffect"))
	return NSSymbolWiggleEffectFromID(rv)
}
// Convenience initializer for a wiggle effect that moves back and forth
// horizontally based on the current locale, starting by moving forward.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolWiggleEffect/wiggleForwardEffect
func (_NSSymbolWiggleEffectClass NSSymbolWiggleEffectClass) WiggleForwardEffect() NSSymbolWiggleEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolWiggleEffectClass.class), objc.Sel("wiggleForwardEffect"))
	return NSSymbolWiggleEffectFromID(rv)
}
// Convenience initializer for a wiggle effect that moves back and forth
// horizontally, starting by moving left.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolWiggleEffect/wiggleLeftEffect
func (_NSSymbolWiggleEffectClass NSSymbolWiggleEffectClass) WiggleLeftEffect() NSSymbolWiggleEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolWiggleEffectClass.class), objc.Sel("wiggleLeftEffect"))
	return NSSymbolWiggleEffectFromID(rv)
}
// Convenience initializer for a wiggle effect that moves back and forth
// horizontally, starting by moving right.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolWiggleEffect/wiggleRightEffect
func (_NSSymbolWiggleEffectClass NSSymbolWiggleEffectClass) WiggleRightEffect() NSSymbolWiggleEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolWiggleEffectClass.class), objc.Sel("wiggleRightEffect"))
	return NSSymbolWiggleEffectFromID(rv)
}
// Convenience initializer for a wiggle effect that moves back and forth
// vertically, starting by moving up.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolWiggleEffect/wiggleUpEffect
func (_NSSymbolWiggleEffectClass NSSymbolWiggleEffectClass) WiggleUpEffect() NSSymbolWiggleEffect {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolWiggleEffectClass.class), objc.Sel("wiggleUpEffect"))
	return NSSymbolWiggleEffectFromID(rv)
}

