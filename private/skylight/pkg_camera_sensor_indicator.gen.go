// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [PKGCameraSensorIndicator] class.
var (
	_PKGCameraSensorIndicatorClass     PKGCameraSensorIndicatorClass
	_PKGCameraSensorIndicatorClassOnce sync.Once
)

func getPKGCameraSensorIndicatorClass() PKGCameraSensorIndicatorClass {
	_PKGCameraSensorIndicatorClassOnce.Do(func() {
		_PKGCameraSensorIndicatorClass = PKGCameraSensorIndicatorClass{class: objc.GetClass("PKGCameraSensorIndicator")}
	})
	return _PKGCameraSensorIndicatorClass
}

// GetPKGCameraSensorIndicatorClass returns the class object for PKGCameraSensorIndicator.
func GetPKGCameraSensorIndicatorClass() PKGCameraSensorIndicatorClass {
	return getPKGCameraSensorIndicatorClass()
}

type PKGCameraSensorIndicatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PKGCameraSensorIndicatorClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PKGCameraSensorIndicatorClass) Alloc() PKGCameraSensorIndicator {
	rv := objc.SendIfResponds[PKGCameraSensorIndicator](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [PKGCameraSensorIndicator.BuiltIn]
//   - [PKGCameraSensorIndicator.SetBuiltIn]
type PKGCameraSensorIndicator struct {
	PKGSystemStatusIndicator
}

// PKGCameraSensorIndicatorFromID constructs a [PKGCameraSensorIndicator] from an objc.ID.
func PKGCameraSensorIndicatorFromID(id objc.ID) PKGCameraSensorIndicator {
	return PKGCameraSensorIndicator{PKGSystemStatusIndicator: PKGSystemStatusIndicatorFromID(id)}
}

// Ensure PKGCameraSensorIndicator implements IPKGCameraSensorIndicator.
var _ IPKGCameraSensorIndicator = PKGCameraSensorIndicator{}

// An interface definition for the [PKGCameraSensorIndicator] class.
//
// # Methods
//
//   - [IPKGCameraSensorIndicator.BuiltIn]
//   - [IPKGCameraSensorIndicator.SetBuiltIn]
type IPKGCameraSensorIndicator interface {
	IPKGSystemStatusIndicator

	// Topic: Methods

	BuiltIn() bool
	SetBuiltIn(value bool)
}

// Init initializes the instance.
func (p PKGCameraSensorIndicator) Init() PKGCameraSensorIndicator {
	rv := objc.SendIfResponds[PKGCameraSensorIndicator](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PKGCameraSensorIndicator) Autorelease() PKGCameraSensorIndicator {
	rv := objc.SendIfResponds[PKGCameraSensorIndicator](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPKGCameraSensorIndicator creates a new PKGCameraSensorIndicator instance.
func NewPKGCameraSensorIndicator() PKGCameraSensorIndicator {
	class := getPKGCameraSensorIndicatorClass()
	rv := objc.SendIfResponds[PKGCameraSensorIndicator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (p PKGCameraSensorIndicator) BuiltIn() bool {
	rv := objc.SendIfResponds[bool](p.ID, objc.Sel("builtIn"))
	return rv
}
func (p PKGCameraSensorIndicator) SetBuiltIn(value bool) {
	objc.SendIfResponds[struct{}](p.ID, objc.Sel("setBuiltIn:"), value)
}
