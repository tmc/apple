// Code generated from Apple documentation for Symbols. DO NOT EDIT.

package symbols

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [NSSymbolMagicReplaceContentTransition] class.
var (
	_NSSymbolMagicReplaceContentTransitionClass     NSSymbolMagicReplaceContentTransitionClass
	_NSSymbolMagicReplaceContentTransitionClassOnce sync.Once
)

func getNSSymbolMagicReplaceContentTransitionClass() NSSymbolMagicReplaceContentTransitionClass {
	_NSSymbolMagicReplaceContentTransitionClassOnce.Do(func() {
		_NSSymbolMagicReplaceContentTransitionClass = NSSymbolMagicReplaceContentTransitionClass{class: objc.GetClass("NSSymbolMagicReplaceContentTransition")}
	})
	return _NSSymbolMagicReplaceContentTransitionClass
}

// GetNSSymbolMagicReplaceContentTransitionClass returns the class object for NSSymbolMagicReplaceContentTransition.
func GetNSSymbolMagicReplaceContentTransitionClass() NSSymbolMagicReplaceContentTransitionClass {
	return getNSSymbolMagicReplaceContentTransitionClass()
}

type NSSymbolMagicReplaceContentTransitionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSymbolMagicReplaceContentTransitionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSymbolMagicReplaceContentTransitionClass) Alloc() NSSymbolMagicReplaceContentTransition {
	rv := objc.Send[NSSymbolMagicReplaceContentTransition](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A symbol effect applies the MagicReplace animation to symbol images.
//
// # Overview
//
// The MagicReplace effect animates common elements across symbol images.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolMagicReplaceContentTransition
type NSSymbolMagicReplaceContentTransition struct {
	NSSymbolContentTransition
}

// NSSymbolMagicReplaceContentTransitionFromID constructs a [NSSymbolMagicReplaceContentTransition] from an objc.ID.
//
// A symbol effect applies the MagicReplace animation to symbol images.
func NSSymbolMagicReplaceContentTransitionFromID(id objc.ID) NSSymbolMagicReplaceContentTransition {
	return NSSymbolMagicReplaceContentTransition{NSSymbolContentTransition: NSSymbolContentTransitionFromID(id)}
}

// Ensure NSSymbolMagicReplaceContentTransition implements INSSymbolMagicReplaceContentTransition.
var _ INSSymbolMagicReplaceContentTransition = NSSymbolMagicReplaceContentTransition{}

// An interface definition for the [NSSymbolMagicReplaceContentTransition] class.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolMagicReplaceContentTransition
type INSSymbolMagicReplaceContentTransition interface {
	INSSymbolContentTransition
}

// Init initializes the instance.
func (s NSSymbolMagicReplaceContentTransition) Init() NSSymbolMagicReplaceContentTransition {
	rv := objc.Send[NSSymbolMagicReplaceContentTransition](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSymbolMagicReplaceContentTransition) Autorelease() NSSymbolMagicReplaceContentTransition {
	rv := objc.Send[NSSymbolMagicReplaceContentTransition](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSymbolMagicReplaceContentTransition creates a new NSSymbolMagicReplaceContentTransition instance.
func NewNSSymbolMagicReplaceContentTransition() NSSymbolMagicReplaceContentTransition {
	class := getNSSymbolMagicReplaceContentTransitionClass()
	rv := objc.Send[NSSymbolMagicReplaceContentTransition](objc.ID(class.class), objc.Sel("new"))
	return rv
}
