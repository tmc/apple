// Code generated from Apple documentation for Symbols. DO NOT EDIT.

package symbols

import (
	"sync"
	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSymbolAutomaticContentTransition] class.
var (
	_NSSymbolAutomaticContentTransitionClass     NSSymbolAutomaticContentTransitionClass
	_NSSymbolAutomaticContentTransitionClassOnce sync.Once
)

func getNSSymbolAutomaticContentTransitionClass() NSSymbolAutomaticContentTransitionClass {
	_NSSymbolAutomaticContentTransitionClassOnce.Do(func() {
		_NSSymbolAutomaticContentTransitionClass = NSSymbolAutomaticContentTransitionClass{class: objc.GetClass("NSSymbolAutomaticContentTransition")}
	})
	return _NSSymbolAutomaticContentTransitionClass
}

// GetNSSymbolAutomaticContentTransitionClass returns the class object for NSSymbolAutomaticContentTransition.
func GetNSSymbolAutomaticContentTransitionClass() NSSymbolAutomaticContentTransitionClass {
	return getNSSymbolAutomaticContentTransitionClass()
}

type NSSymbolAutomaticContentTransitionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSymbolAutomaticContentTransitionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSymbolAutomaticContentTransitionClass) Alloc() NSSymbolAutomaticContentTransition {
	rv := objc.Send[NSSymbolAutomaticContentTransition](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A type that applies the default animation to a symbol-based image in a
// context-sensitive manner.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolAutomaticContentTransition
type NSSymbolAutomaticContentTransition struct {
	NSSymbolContentTransition
}

// NSSymbolAutomaticContentTransitionFromID constructs a [NSSymbolAutomaticContentTransition] from an objc.ID.
//
// A type that applies the default animation to a symbol-based image in a
// context-sensitive manner.
func NSSymbolAutomaticContentTransitionFromID(id objc.ID) NSSymbolAutomaticContentTransition {
	return NSSymbolAutomaticContentTransition{NSSymbolContentTransition: NSSymbolContentTransitionFromID(id)}
}
// Ensure NSSymbolAutomaticContentTransition implements INSSymbolAutomaticContentTransition.
var _ INSSymbolAutomaticContentTransition = NSSymbolAutomaticContentTransition{}

// An interface definition for the [NSSymbolAutomaticContentTransition] class.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolAutomaticContentTransition
type INSSymbolAutomaticContentTransition interface {
	INSSymbolContentTransition
}

// Init initializes the instance.
func (s NSSymbolAutomaticContentTransition) Init() NSSymbolAutomaticContentTransition {
	rv := objc.Send[NSSymbolAutomaticContentTransition](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSymbolAutomaticContentTransition) Autorelease() NSSymbolAutomaticContentTransition {
	rv := objc.Send[NSSymbolAutomaticContentTransition](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSymbolAutomaticContentTransition creates a new NSSymbolAutomaticContentTransition instance.
func NewNSSymbolAutomaticContentTransition() NSSymbolAutomaticContentTransition {
	class := getNSSymbolAutomaticContentTransitionClass()
	rv := objc.Send[NSSymbolAutomaticContentTransition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A transition that applies the default animation to a symbol-based image in
// a context-sensitive manner.
//
// # Return Value
// 
// A new instance of the automatic transition.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolAutomaticContentTransition/transition
func (_NSSymbolAutomaticContentTransitionClass NSSymbolAutomaticContentTransitionClass) Transition() NSSymbolAutomaticContentTransition {
	rv := objc.Send[objc.ID](objc.ID(_NSSymbolAutomaticContentTransitionClass.class), objc.Sel("transition"))
	return NSSymbolAutomaticContentTransitionFromID(rv)
}

