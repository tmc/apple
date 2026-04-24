// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightPKGSystemStatusIndicatorsAnimator] class.
var (
	_SkyLightPKGSystemStatusIndicatorsAnimatorClass     SkyLightPKGSystemStatusIndicatorsAnimatorClass
	_SkyLightPKGSystemStatusIndicatorsAnimatorClassOnce sync.Once
)

func getSkyLightPKGSystemStatusIndicatorsAnimatorClass() SkyLightPKGSystemStatusIndicatorsAnimatorClass {
	_SkyLightPKGSystemStatusIndicatorsAnimatorClassOnce.Do(func() {
		_SkyLightPKGSystemStatusIndicatorsAnimatorClass = SkyLightPKGSystemStatusIndicatorsAnimatorClass{class: objc.GetClass("SkyLight.PKGSystemStatusIndicatorsAnimator")}
	})
	return _SkyLightPKGSystemStatusIndicatorsAnimatorClass
}

// GetSkyLightPKGSystemStatusIndicatorsAnimatorClass returns the class object for SkyLight.PKGSystemStatusIndicatorsAnimator.
func GetSkyLightPKGSystemStatusIndicatorsAnimatorClass() SkyLightPKGSystemStatusIndicatorsAnimatorClass {
	return getSkyLightPKGSystemStatusIndicatorsAnimatorClass()
}

type SkyLightPKGSystemStatusIndicatorsAnimatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightPKGSystemStatusIndicatorsAnimatorClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightPKGSystemStatusIndicatorsAnimatorClass) Alloc() SkyLightPKGSystemStatusIndicatorsAnimator {
	rv := objc.Send[SkyLightPKGSystemStatusIndicatorsAnimator](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SkyLight.PKGSystemStatusIndicatorsAnimator
type SkyLightPKGSystemStatusIndicatorsAnimator struct {
	objectivec.Object
}

// SkyLightPKGSystemStatusIndicatorsAnimatorFromID constructs a [SkyLightPKGSystemStatusIndicatorsAnimator] from an objc.ID.
func SkyLightPKGSystemStatusIndicatorsAnimatorFromID(id objc.ID) SkyLightPKGSystemStatusIndicatorsAnimator {
	return SkyLightPKGSystemStatusIndicatorsAnimator{objectivec.Object{ID: id}}
}

// NOTE: SkyLightPKGSystemStatusIndicatorsAnimator struct embeds objectivec.Object (parent type unavailable) but
// ISkyLightPKGSystemStatusIndicatorsAnimator embeds the parent interface; skip compile-time assertion.

// An interface definition for the [SkyLightPKGSystemStatusIndicatorsAnimator] class.
//
// See: https://developer.apple.com/documentation/SkyLight/SkyLight.PKGSystemStatusIndicatorsAnimator
type ISkyLightPKGSystemStatusIndicatorsAnimator interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightPKGSystemStatusIndicatorsAnimator) Init() SkyLightPKGSystemStatusIndicatorsAnimator {
	rv := objc.Send[SkyLightPKGSystemStatusIndicatorsAnimator](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightPKGSystemStatusIndicatorsAnimator) Autorelease() SkyLightPKGSystemStatusIndicatorsAnimator {
	rv := objc.Send[SkyLightPKGSystemStatusIndicatorsAnimator](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightPKGSystemStatusIndicatorsAnimator creates a new SkyLightPKGSystemStatusIndicatorsAnimator instance.
func NewSkyLightPKGSystemStatusIndicatorsAnimator() SkyLightPKGSystemStatusIndicatorsAnimator {
	class := getSkyLightPKGSystemStatusIndicatorsAnimatorClass()
	rv := objc.Send[SkyLightPKGSystemStatusIndicatorsAnimator](objc.ID(class.class), objc.Sel("new"))
	return rv
}
