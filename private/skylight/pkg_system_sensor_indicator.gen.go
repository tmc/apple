// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [PKGSystemSensorIndicator] class.
var (
	_PKGSystemSensorIndicatorClass     PKGSystemSensorIndicatorClass
	_PKGSystemSensorIndicatorClassOnce sync.Once
)

func getPKGSystemSensorIndicatorClass() PKGSystemSensorIndicatorClass {
	_PKGSystemSensorIndicatorClassOnce.Do(func() {
		_PKGSystemSensorIndicatorClass = PKGSystemSensorIndicatorClass{class: objc.GetClass("PKGSystemSensorIndicator")}
	})
	return _PKGSystemSensorIndicatorClass
}

// GetPKGSystemSensorIndicatorClass returns the class object for PKGSystemSensorIndicator.
func GetPKGSystemSensorIndicatorClass() PKGSystemSensorIndicatorClass {
	return getPKGSystemSensorIndicatorClass()
}

type PKGSystemSensorIndicatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PKGSystemSensorIndicatorClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PKGSystemSensorIndicatorClass) Alloc() PKGSystemSensorIndicator {
	rv := objc.SendIfResponds[PKGSystemSensorIndicator](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

type PKGSystemSensorIndicator struct {
	PKGSystemStatusIndicator
}

// PKGSystemSensorIndicatorFromID constructs a [PKGSystemSensorIndicator] from an objc.ID.
func PKGSystemSensorIndicatorFromID(id objc.ID) PKGSystemSensorIndicator {
	return PKGSystemSensorIndicator{PKGSystemStatusIndicator: PKGSystemStatusIndicatorFromID(id)}
}

// Ensure PKGSystemSensorIndicator implements IPKGSystemSensorIndicator.
var _ IPKGSystemSensorIndicator = PKGSystemSensorIndicator{}

// An interface definition for the [PKGSystemSensorIndicator] class.
type IPKGSystemSensorIndicator interface {
	IPKGSystemStatusIndicator
}

// Init initializes the instance.
func (p PKGSystemSensorIndicator) Init() PKGSystemSensorIndicator {
	rv := objc.SendIfResponds[PKGSystemSensorIndicator](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PKGSystemSensorIndicator) Autorelease() PKGSystemSensorIndicator {
	rv := objc.SendIfResponds[PKGSystemSensorIndicator](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPKGSystemSensorIndicator creates a new PKGSystemSensorIndicator instance.
func NewPKGSystemSensorIndicator() PKGSystemSensorIndicator {
	class := getPKGSystemSensorIndicatorClass()
	rv := objc.SendIfResponds[PKGSystemSensorIndicator](objc.ID(class.class), objc.Sel("new"))
	return rv
}
