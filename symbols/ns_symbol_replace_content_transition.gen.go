// Code generated from Apple documentation for Symbols. DO NOT EDIT.

package symbols

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSymbolReplaceContentTransition] class.
var (
	_NSSymbolReplaceContentTransitionClass     NSSymbolReplaceContentTransitionClass
	_NSSymbolReplaceContentTransitionClassOnce sync.Once
)

func getNSSymbolReplaceContentTransitionClass() NSSymbolReplaceContentTransitionClass {
	_NSSymbolReplaceContentTransitionClassOnce.Do(func() {
		_NSSymbolReplaceContentTransitionClass = NSSymbolReplaceContentTransitionClass{class: objc.GetClass("NSSymbolReplaceContentTransition")}
	})
	return _NSSymbolReplaceContentTransitionClass
}

// GetNSSymbolReplaceContentTransitionClass returns the class object for NSSymbolReplaceContentTransition.
func GetNSSymbolReplaceContentTransitionClass() NSSymbolReplaceContentTransitionClass {
	return getNSSymbolReplaceContentTransitionClass()
}

type NSSymbolReplaceContentTransitionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSymbolReplaceContentTransitionClass) Alloc() NSSymbolReplaceContentTransition {
	rv := objc.Send[NSSymbolReplaceContentTransition](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A type that replaces the layers of one symbol-based image with those of
// another.
//
// # Overview
// 
// A replace transition animates the change from one symbol image to another.
// You choose from one of the predefined scaling animations: Down-Up, Off-Up,
// and Up-Up.
// 
// Down-Up: The initial symbol scales down as it’s removed, and the new
// symbol scales up as it’s added. Off-Up: The initial symbol is removed
// with no animation, and the new symbol scales up as it’s added. Up-Up: The
// initial symbol scales up as it’s removed, and the new symbol scales up as
// it’s added.
//
// # Determining effect scope
//
//   - [NSSymbolReplaceContentTransition.TransitionWithByLayer]: An effect that replaces each layer separately.
//   - [NSSymbolReplaceContentTransition.TransitionWithWholeSymbol]: An effect that replaces all layers simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolReplaceContentTransition
type NSSymbolReplaceContentTransition struct {
	NSSymbolContentTransition
}

// NSSymbolReplaceContentTransitionFromID constructs a [NSSymbolReplaceContentTransition] from an objc.ID.
//
// A type that replaces the layers of one symbol-based image with those of
// another.
func NSSymbolReplaceContentTransitionFromID(id objc.ID) NSSymbolReplaceContentTransition {
	return NSSymbolReplaceContentTransition{NSSymbolContentTransition: NSSymbolContentTransitionFromID(id)}
}
// Ensure NSSymbolReplaceContentTransition implements INSSymbolReplaceContentTransition.
var _ INSSymbolReplaceContentTransition = NSSymbolReplaceContentTransition{}

// An interface definition for the [NSSymbolReplaceContentTransition] class.
//
// # Determining effect scope
//
//   - [INSSymbolReplaceContentTransition.TransitionWithByLayer]: An effect that replaces each layer separately.
//   - [INSSymbolReplaceContentTransition.TransitionWithWholeSymbol]: An effect that replaces all layers simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolReplaceContentTransition
type INSSymbolReplaceContentTransition interface {
	INSSymbolContentTransition

	// Topic: Determining effect scope

	// An effect that replaces each layer separately.
	TransitionWithByLayer() INSSymbolReplaceContentTransition
	// An effect that replaces all layers simultaneously.
	TransitionWithWholeSymbol() INSSymbolReplaceContentTransition
}

// Init initializes the instance.
func (s NSSymbolReplaceContentTransition) Init() NSSymbolReplaceContentTransition {
	rv := objc.Send[NSSymbolReplaceContentTransition](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSymbolReplaceContentTransition) Autorelease() NSSymbolReplaceContentTransition {
	rv := objc.Send[NSSymbolReplaceContentTransition](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSymbolReplaceContentTransition creates a new NSSymbolReplaceContentTransition instance.
func NewNSSymbolReplaceContentTransition() NSSymbolReplaceContentTransition {
	class := getNSSymbolReplaceContentTransitionClass()
	rv := objc.Send[NSSymbolReplaceContentTransition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// An effect that replaces each layer separately.
//
// # Return Value
// 
// A copy of the transition that replaces each layer separately.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolReplaceContentTransition/transitionWithByLayer
func (s NSSymbolReplaceContentTransition) TransitionWithByLayer() INSSymbolReplaceContentTransition {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("transitionWithByLayer"))
	return NSSymbolReplaceContentTransitionFromID(rv)
}

// An effect that replaces all layers simultaneously.
//
// # Return Value
// 
// A copy of the transition that replaces all layers simultaneously.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolReplaceContentTransition/transitionWithWholeSymbol
func (s NSSymbolReplaceContentTransition) TransitionWithWholeSymbol() INSSymbolReplaceContentTransition {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("transitionWithWholeSymbol"))
	return NSSymbolReplaceContentTransitionFromID(rv)
}

// An effect that replaces the layers of one symbol-based image with those of
// another.
//
// # Return Value
// 
// A new instance of the replace transition.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolReplaceContentTransition/transition
func (_NSSymbolReplaceContentTransitionClass NSSymbolReplaceContentTransitionClass) Transition() NSSymbolReplaceContentTransition {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolReplaceContentTransitionClass.class), objc.Sel("transition"))
	return NSSymbolReplaceContentTransitionFromID(rv)
}

// An effect that replaces a symbol by scaling it down, and scaling a
// different symbol up.
//
// # Return Value
// 
// A new instance of the transition that uses the Down-Up animation.
//
// # Discussion
// 
// The initial symbol scales down as it’s removed, and the new symbol scales
// up as it’s added.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolReplaceContentTransition/replaceDownUpTransition
func (_NSSymbolReplaceContentTransitionClass NSSymbolReplaceContentTransitionClass) ReplaceDownUpTransition() NSSymbolReplaceContentTransition {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolReplaceContentTransitionClass.class), objc.Sel("replaceDownUpTransition"))
	return NSSymbolReplaceContentTransitionFromID(rv)
}

// An effect that replaces a symbol by removing it, and scaling a different
// symbol up.
//
// # Return Value
// 
// A new instance of the transition that uses the Off-Up animation.
//
// # Discussion
// 
// The initial symbol is removed with no animation, and the new symbol scales
// up as it’s added.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolReplaceContentTransition/replaceOffUpTransition
func (_NSSymbolReplaceContentTransitionClass NSSymbolReplaceContentTransitionClass) ReplaceOffUpTransition() NSSymbolReplaceContentTransition {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolReplaceContentTransitionClass.class), objc.Sel("replaceOffUpTransition"))
	return NSSymbolReplaceContentTransitionFromID(rv)
}

// An effect that replaces a symbol by scaling it up, and scaling a different
// symbol up.
//
// # Return Value
// 
// A new instance of the transition that uses the Up-Up animation.
//
// # Discussion
// 
// The initial symbol scales up as it’s removed, and the new symbol scales
// up as it’s added.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolReplaceContentTransition/replaceUpUpTransition
func (_NSSymbolReplaceContentTransitionClass NSSymbolReplaceContentTransitionClass) ReplaceUpUpTransition() NSSymbolReplaceContentTransition {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolReplaceContentTransitionClass.class), objc.Sel("replaceUpUpTransition"))
	return NSSymbolReplaceContentTransitionFromID(rv)
}

// Convenience initializer for a MagicReplace content transition with a
// configured Replace fallback.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolReplaceContentTransition/magicTransitionWithFallback:
func (_NSSymbolReplaceContentTransitionClass NSSymbolReplaceContentTransitionClass) MagicTransitionWithFallback(fallback INSSymbolReplaceContentTransition) NSSymbolMagicReplaceContentTransition {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolReplaceContentTransitionClass.class), objc.Sel("magicTransitionWithFallback:"), fallback)
	return NSSymbolMagicReplaceContentTransitionFromID(rv)
}

