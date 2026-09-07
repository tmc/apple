// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [PKGScreenCaptureIndicator] class.
var (
	_PKGScreenCaptureIndicatorClass     PKGScreenCaptureIndicatorClass
	_PKGScreenCaptureIndicatorClassOnce sync.Once
)

func getPKGScreenCaptureIndicatorClass() PKGScreenCaptureIndicatorClass {
	_PKGScreenCaptureIndicatorClassOnce.Do(func() {
		_PKGScreenCaptureIndicatorClass = PKGScreenCaptureIndicatorClass{class: objc.GetClass("PKGScreenCaptureIndicator")}
	})
	return _PKGScreenCaptureIndicatorClass
}

// GetPKGScreenCaptureIndicatorClass returns the class object for PKGScreenCaptureIndicator.
func GetPKGScreenCaptureIndicatorClass() PKGScreenCaptureIndicatorClass {
	return getPKGScreenCaptureIndicatorClass()
}

type PKGScreenCaptureIndicatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PKGScreenCaptureIndicatorClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PKGScreenCaptureIndicatorClass) Alloc() PKGScreenCaptureIndicator {
	rv := objc.SendIfResponds[PKGScreenCaptureIndicator](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

type PKGScreenCaptureIndicator struct {
	PKGSystemStatusIndicator
}

// PKGScreenCaptureIndicatorFromID constructs a [PKGScreenCaptureIndicator] from an objc.ID.
func PKGScreenCaptureIndicatorFromID(id objc.ID) PKGScreenCaptureIndicator {
	return PKGScreenCaptureIndicator{PKGSystemStatusIndicator: PKGSystemStatusIndicatorFromID(id)}
}

// Ensure PKGScreenCaptureIndicator implements IPKGScreenCaptureIndicator.
var _ IPKGScreenCaptureIndicator = PKGScreenCaptureIndicator{}

// An interface definition for the [PKGScreenCaptureIndicator] class.
type IPKGScreenCaptureIndicator interface {
	IPKGSystemStatusIndicator
}

// Init initializes the instance.
func (p PKGScreenCaptureIndicator) Init() PKGScreenCaptureIndicator {
	rv := objc.SendIfResponds[PKGScreenCaptureIndicator](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PKGScreenCaptureIndicator) Autorelease() PKGScreenCaptureIndicator {
	rv := objc.SendIfResponds[PKGScreenCaptureIndicator](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPKGScreenCaptureIndicator creates a new PKGScreenCaptureIndicator instance.
func NewPKGScreenCaptureIndicator() PKGScreenCaptureIndicator {
	class := getPKGScreenCaptureIndicatorClass()
	rv := objc.SendIfResponds[PKGScreenCaptureIndicator](objc.ID(class.class), objc.Sel("new"))
	return rv
}
