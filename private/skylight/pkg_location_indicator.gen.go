// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [PKGLocationIndicator] class.
var (
	_PKGLocationIndicatorClass     PKGLocationIndicatorClass
	_PKGLocationIndicatorClassOnce sync.Once
)

func getPKGLocationIndicatorClass() PKGLocationIndicatorClass {
	_PKGLocationIndicatorClassOnce.Do(func() {
		_PKGLocationIndicatorClass = PKGLocationIndicatorClass{class: objc.GetClass("PKGLocationIndicator")}
	})
	return _PKGLocationIndicatorClass
}

// GetPKGLocationIndicatorClass returns the class object for PKGLocationIndicator.
func GetPKGLocationIndicatorClass() PKGLocationIndicatorClass {
	return getPKGLocationIndicatorClass()
}

type PKGLocationIndicatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PKGLocationIndicatorClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PKGLocationIndicatorClass) Alloc() PKGLocationIndicator {
	rv := objc.SendIfResponds[PKGLocationIndicator](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

type PKGLocationIndicator struct {
	PKGSystemStatusIndicator
}

// PKGLocationIndicatorFromID constructs a [PKGLocationIndicator] from an objc.ID.
func PKGLocationIndicatorFromID(id objc.ID) PKGLocationIndicator {
	return PKGLocationIndicator{PKGSystemStatusIndicator: PKGSystemStatusIndicatorFromID(id)}
}

// Ensure PKGLocationIndicator implements IPKGLocationIndicator.
var _ IPKGLocationIndicator = PKGLocationIndicator{}

// An interface definition for the [PKGLocationIndicator] class.
type IPKGLocationIndicator interface {
	IPKGSystemStatusIndicator
}

// Init initializes the instance.
func (p PKGLocationIndicator) Init() PKGLocationIndicator {
	rv := objc.SendIfResponds[PKGLocationIndicator](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PKGLocationIndicator) Autorelease() PKGLocationIndicator {
	rv := objc.SendIfResponds[PKGLocationIndicator](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPKGLocationIndicator creates a new PKGLocationIndicator instance.
func NewPKGLocationIndicator() PKGLocationIndicator {
	class := getPKGLocationIndicatorClass()
	rv := objc.SendIfResponds[PKGLocationIndicator](objc.ID(class.class), objc.Sel("new"))
	return rv
}
