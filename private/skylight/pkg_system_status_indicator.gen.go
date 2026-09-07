// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [PKGSystemStatusIndicator] class.
var (
	_PKGSystemStatusIndicatorClass     PKGSystemStatusIndicatorClass
	_PKGSystemStatusIndicatorClassOnce sync.Once
)

func getPKGSystemStatusIndicatorClass() PKGSystemStatusIndicatorClass {
	_PKGSystemStatusIndicatorClassOnce.Do(func() {
		_PKGSystemStatusIndicatorClass = PKGSystemStatusIndicatorClass{class: objc.GetClass("PKGSystemStatusIndicator")}
	})
	return _PKGSystemStatusIndicatorClass
}

// GetPKGSystemStatusIndicatorClass returns the class object for PKGSystemStatusIndicator.
func GetPKGSystemStatusIndicatorClass() PKGSystemStatusIndicatorClass {
	return getPKGSystemStatusIndicatorClass()
}

type PKGSystemStatusIndicatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PKGSystemStatusIndicatorClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PKGSystemStatusIndicatorClass) Alloc() PKGSystemStatusIndicator {
	rv := objc.SendIfResponds[PKGSystemStatusIndicator](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

type PKGSystemStatusIndicator struct {
	objectivec.Object
}

// PKGSystemStatusIndicatorFromID constructs a [PKGSystemStatusIndicator] from an objc.ID.
func PKGSystemStatusIndicatorFromID(id objc.ID) PKGSystemStatusIndicator {
	return PKGSystemStatusIndicator{objectivec.Object{ID: id}}
}

// Ensure PKGSystemStatusIndicator implements IPKGSystemStatusIndicator.
var _ IPKGSystemStatusIndicator = PKGSystemStatusIndicator{}

// An interface definition for the [PKGSystemStatusIndicator] class.
type IPKGSystemStatusIndicator interface {
	objectivec.IObject
}

// Init initializes the instance.
func (p PKGSystemStatusIndicator) Init() PKGSystemStatusIndicator {
	rv := objc.SendIfResponds[PKGSystemStatusIndicator](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PKGSystemStatusIndicator) Autorelease() PKGSystemStatusIndicator {
	rv := objc.SendIfResponds[PKGSystemStatusIndicator](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPKGSystemStatusIndicator creates a new PKGSystemStatusIndicator instance.
func NewPKGSystemStatusIndicator() PKGSystemStatusIndicator {
	class := getPKGSystemStatusIndicatorClass()
	rv := objc.SendIfResponds[PKGSystemStatusIndicator](objc.ID(class.class), objc.Sel("new"))
	return rv
}
