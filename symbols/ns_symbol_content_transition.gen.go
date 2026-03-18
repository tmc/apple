// Code generated from Apple documentation for Symbols. DO NOT EDIT.

package symbols

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSSymbolContentTransition] class.
var (
	_NSSymbolContentTransitionClass     NSSymbolContentTransitionClass
	_NSSymbolContentTransitionClassOnce sync.Once
)

func getNSSymbolContentTransitionClass() NSSymbolContentTransitionClass {
	_NSSymbolContentTransitionClassOnce.Do(func() {
		_NSSymbolContentTransitionClass = NSSymbolContentTransitionClass{class: objc.GetClass("NSSymbolContentTransition")}
	})
	return _NSSymbolContentTransitionClass
}

// GetNSSymbolContentTransitionClass returns the class object for NSSymbolContentTransition.
func GetNSSymbolContentTransitionClass() NSSymbolContentTransitionClass {
	return getNSSymbolContentTransitionClass()
}

type NSSymbolContentTransitionClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSymbolContentTransitionClass) Alloc() NSSymbolContentTransition {
	rv := objc.Send[NSSymbolContentTransition](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract base class for transitions you can apply to symbol-based
// images.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolContentTransition
type NSSymbolContentTransition struct {
	objectivec.Object
}

// NSSymbolContentTransitionFromID constructs a [NSSymbolContentTransition] from an objc.ID.
//
// An abstract base class for transitions you can apply to symbol-based
// images.
func NSSymbolContentTransitionFromID(id objc.ID) NSSymbolContentTransition {
	return NSSymbolContentTransition{objectivec.Object{ID: id}}
}
// NOTE: NSSymbolContentTransition adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSymbolContentTransition] class.
//
// See: https://developer.apple.com/documentation/Symbols/NSSymbolContentTransition
type INSSymbolContentTransition interface {
	objectivec.IObject

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (s NSSymbolContentTransition) Init() NSSymbolContentTransition {
	rv := objc.Send[NSSymbolContentTransition](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSymbolContentTransition) Autorelease() NSSymbolContentTransition {
	rv := objc.Send[NSSymbolContentTransition](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSymbolContentTransition creates a new NSSymbolContentTransition instance.
func NewNSSymbolContentTransition() NSSymbolContentTransition {
	class := getNSSymbolContentTransitionClass()
	rv := objc.Send[NSSymbolContentTransition](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (s NSSymbolContentTransition) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

