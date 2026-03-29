// Code generated from Apple documentation for Symbols. DO NOT EDIT.

package symbols

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSymbolEffect] class.
var (
	_NSSymbolEffectClass     NSSymbolEffectClass
	_NSSymbolEffectClassOnce sync.Once
)

func getNSSymbolEffectClass() NSSymbolEffectClass {
	_NSSymbolEffectClassOnce.Do(func() {
		_NSSymbolEffectClass = NSSymbolEffectClass{class: objc.GetClass("NSSymbolEffect")}
	})
	return _NSSymbolEffectClass
}

// GetNSSymbolEffectClass returns the class object for NSSymbolEffect.
func GetNSSymbolEffectClass() NSSymbolEffectClass {
	return getNSSymbolEffectClass()
}

type NSSymbolEffectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSymbolEffectClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSymbolEffectClass) Alloc() NSSymbolEffect {
	rv := objc.Send[NSSymbolEffect](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract base class for effects that you can apply to a symbol-based
// image.
//
// # Overview
// 
// You don’t use this class directly. Instead, use a class that inherits
// from this one, such as [NSSymbolBounceEffect].
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolEffect
type NSSymbolEffect struct {
	objectivec.Object
}

// NSSymbolEffectFromID constructs a [NSSymbolEffect] from an objc.ID.
//
// An abstract base class for effects that you can apply to a symbol-based
// image.
func NSSymbolEffectFromID(id objc.ID) NSSymbolEffect {
	return NSSymbolEffect{objectivec.Object{ID: id}}
}
// NOTE: NSSymbolEffect adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSymbolEffect] class.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolEffect
type INSSymbolEffect interface {
	objectivec.IObject

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s NSSymbolEffect) Init() NSSymbolEffect {
	rv := objc.Send[NSSymbolEffect](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSymbolEffect) Autorelease() NSSymbolEffect {
	rv := objc.Send[NSSymbolEffect](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSymbolEffect creates a new NSSymbolEffect instance.
func NewNSSymbolEffect() NSSymbolEffect {
	class := getNSSymbolEffectClass()
	rv := objc.Send[NSSymbolEffect](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (s NSSymbolEffect) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

