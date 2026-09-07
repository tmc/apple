// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [PKGSystemAudioIndicator] class.
var (
	_PKGSystemAudioIndicatorClass     PKGSystemAudioIndicatorClass
	_PKGSystemAudioIndicatorClassOnce sync.Once
)

func getPKGSystemAudioIndicatorClass() PKGSystemAudioIndicatorClass {
	_PKGSystemAudioIndicatorClassOnce.Do(func() {
		_PKGSystemAudioIndicatorClass = PKGSystemAudioIndicatorClass{class: objc.GetClass("PKGSystemAudioIndicator")}
	})
	return _PKGSystemAudioIndicatorClass
}

// GetPKGSystemAudioIndicatorClass returns the class object for PKGSystemAudioIndicator.
func GetPKGSystemAudioIndicatorClass() PKGSystemAudioIndicatorClass {
	return getPKGSystemAudioIndicatorClass()
}

type PKGSystemAudioIndicatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (pc PKGSystemAudioIndicatorClass) Class() objc.Class {
	return pc.class
}

// Alloc allocates memory for a new instance of the class.
func (pc PKGSystemAudioIndicatorClass) Alloc() PKGSystemAudioIndicator {
	rv := objc.SendIfResponds[PKGSystemAudioIndicator](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

type PKGSystemAudioIndicator struct {
	PKGSystemStatusIndicator
}

// PKGSystemAudioIndicatorFromID constructs a [PKGSystemAudioIndicator] from an objc.ID.
func PKGSystemAudioIndicatorFromID(id objc.ID) PKGSystemAudioIndicator {
	return PKGSystemAudioIndicator{PKGSystemStatusIndicator: PKGSystemStatusIndicatorFromID(id)}
}

// Ensure PKGSystemAudioIndicator implements IPKGSystemAudioIndicator.
var _ IPKGSystemAudioIndicator = PKGSystemAudioIndicator{}

// An interface definition for the [PKGSystemAudioIndicator] class.
type IPKGSystemAudioIndicator interface {
	IPKGSystemStatusIndicator
}

// Init initializes the instance.
func (p PKGSystemAudioIndicator) Init() PKGSystemAudioIndicator {
	rv := objc.SendIfResponds[PKGSystemAudioIndicator](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p PKGSystemAudioIndicator) Autorelease() PKGSystemAudioIndicator {
	rv := objc.SendIfResponds[PKGSystemAudioIndicator](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewPKGSystemAudioIndicator creates a new PKGSystemAudioIndicator instance.
func NewPKGSystemAudioIndicator() PKGSystemAudioIndicator {
	class := getPKGSystemAudioIndicatorClass()
	rv := objc.SendIfResponds[PKGSystemAudioIndicator](objc.ID(class.class), objc.Sel("new"))
	return rv
}
