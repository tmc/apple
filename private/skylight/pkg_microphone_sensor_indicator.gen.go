// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [PKGMicrophoneSensorIndicator] class.
var (
	_PKGMicrophoneSensorIndicatorClass     PKGMicrophoneSensorIndicatorClass
	_PKGMicrophoneSensorIndicatorClassOnce sync.Once
)

func getPKGMicrophoneSensorIndicatorClass() PKGMicrophoneSensorIndicatorClass {
	_PKGMicrophoneSensorIndicatorClassOnce.Do(func() {
		_PKGMicrophoneSensorIndicatorClass = PKGMicrophoneSensorIndicatorClass{class: objc.GetClass("PKGMicrophoneSensorIndicator")}
	})
	return _PKGMicrophoneSensorIndicatorClass
}

// GetPKGMicrophoneSensorIndicatorClass returns the class object for PKGMicrophoneSensorIndicator.
func GetPKGMicrophoneSensorIndicatorClass() PKGMicrophoneSensorIndicatorClass {
	return getPKGMicrophoneSensorIndicatorClass()
}

type PKGMicrophoneSensorIndicatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PKGMicrophoneSensorIndicatorClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PKGMicrophoneSensorIndicatorClass) Alloc() PKGMicrophoneSensorIndicator {
	rv := objc.SendIfResponds[PKGMicrophoneSensorIndicator](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [PKGMicrophoneSensorIndicator.BuiltIn]
//   - [PKGMicrophoneSensorIndicator.SetBuiltIn]
type PKGMicrophoneSensorIndicator struct {
	PKGSystemStatusIndicator
}

// PKGMicrophoneSensorIndicatorFromID constructs a [PKGMicrophoneSensorIndicator] from an objc.ID.
func PKGMicrophoneSensorIndicatorFromID(id objc.ID) PKGMicrophoneSensorIndicator {
	return PKGMicrophoneSensorIndicator{PKGSystemStatusIndicator: PKGSystemStatusIndicatorFromID(id)}
}

// Ensure PKGMicrophoneSensorIndicator implements IPKGMicrophoneSensorIndicator.
var _ IPKGMicrophoneSensorIndicator = PKGMicrophoneSensorIndicator{}

// An interface definition for the [PKGMicrophoneSensorIndicator] class.
//
// # Methods
//
//   - [IPKGMicrophoneSensorIndicator.BuiltIn]
//   - [IPKGMicrophoneSensorIndicator.SetBuiltIn]
type IPKGMicrophoneSensorIndicator interface {
	IPKGSystemStatusIndicator

	// Topic: Methods

	BuiltIn() bool
	SetBuiltIn(value bool)
}

// Init initializes the instance.
func (p PKGMicrophoneSensorIndicator) Init() PKGMicrophoneSensorIndicator {
	rv := objc.SendIfResponds[PKGMicrophoneSensorIndicator](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PKGMicrophoneSensorIndicator) Autorelease() PKGMicrophoneSensorIndicator {
	rv := objc.SendIfResponds[PKGMicrophoneSensorIndicator](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPKGMicrophoneSensorIndicator creates a new PKGMicrophoneSensorIndicator instance.
func NewPKGMicrophoneSensorIndicator() PKGMicrophoneSensorIndicator {
	class := getPKGMicrophoneSensorIndicatorClass()
	rv := objc.SendIfResponds[PKGMicrophoneSensorIndicator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (p PKGMicrophoneSensorIndicator) BuiltIn() bool {
	rv := objc.SendIfResponds[bool](p.ID, objc.Sel("builtIn"))
	return rv
}
func (p PKGMicrophoneSensorIndicator) SetBuiltIn(value bool) {
	objc.SendIfResponds[struct{}](p.ID, objc.Sel("setBuiltIn:"), value)
}
