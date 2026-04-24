// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightPKGSystemStatusIndicatorsTimer] class.
var (
	_SkyLightPKGSystemStatusIndicatorsTimerClass     SkyLightPKGSystemStatusIndicatorsTimerClass
	_SkyLightPKGSystemStatusIndicatorsTimerClassOnce sync.Once
)

func getSkyLightPKGSystemStatusIndicatorsTimerClass() SkyLightPKGSystemStatusIndicatorsTimerClass {
	_SkyLightPKGSystemStatusIndicatorsTimerClassOnce.Do(func() {
		_SkyLightPKGSystemStatusIndicatorsTimerClass = SkyLightPKGSystemStatusIndicatorsTimerClass{class: objc.GetClass("SkyLight.PKGSystemStatusIndicatorsTimer")}
	})
	return _SkyLightPKGSystemStatusIndicatorsTimerClass
}

// GetSkyLightPKGSystemStatusIndicatorsTimerClass returns the class object for SkyLight.PKGSystemStatusIndicatorsTimer.
func GetSkyLightPKGSystemStatusIndicatorsTimerClass() SkyLightPKGSystemStatusIndicatorsTimerClass {
	return getSkyLightPKGSystemStatusIndicatorsTimerClass()
}

type SkyLightPKGSystemStatusIndicatorsTimerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightPKGSystemStatusIndicatorsTimerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightPKGSystemStatusIndicatorsTimerClass) Alloc() SkyLightPKGSystemStatusIndicatorsTimer {
	rv := objc.Send[SkyLightPKGSystemStatusIndicatorsTimer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SkyLight.PKGSystemStatusIndicatorsTimer
type SkyLightPKGSystemStatusIndicatorsTimer struct {
	objectivec.Object
}

// SkyLightPKGSystemStatusIndicatorsTimerFromID constructs a [SkyLightPKGSystemStatusIndicatorsTimer] from an objc.ID.
func SkyLightPKGSystemStatusIndicatorsTimerFromID(id objc.ID) SkyLightPKGSystemStatusIndicatorsTimer {
	return SkyLightPKGSystemStatusIndicatorsTimer{objectivec.Object{ID: id}}
}

// NOTE: SkyLightPKGSystemStatusIndicatorsTimer struct embeds objectivec.Object (parent type unavailable) but
// ISkyLightPKGSystemStatusIndicatorsTimer embeds the parent interface; skip compile-time assertion.

// An interface definition for the [SkyLightPKGSystemStatusIndicatorsTimer] class.
//
// See: https://developer.apple.com/documentation/SkyLight/SkyLight.PKGSystemStatusIndicatorsTimer
type ISkyLightPKGSystemStatusIndicatorsTimer interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightPKGSystemStatusIndicatorsTimer) Init() SkyLightPKGSystemStatusIndicatorsTimer {
	rv := objc.Send[SkyLightPKGSystemStatusIndicatorsTimer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightPKGSystemStatusIndicatorsTimer) Autorelease() SkyLightPKGSystemStatusIndicatorsTimer {
	rv := objc.Send[SkyLightPKGSystemStatusIndicatorsTimer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightPKGSystemStatusIndicatorsTimer creates a new SkyLightPKGSystemStatusIndicatorsTimer instance.
func NewSkyLightPKGSystemStatusIndicatorsTimer() SkyLightPKGSystemStatusIndicatorsTimer {
	class := getSkyLightPKGSystemStatusIndicatorsTimerClass()
	rv := objc.Send[SkyLightPKGSystemStatusIndicatorsTimer](objc.ID(class.class), objc.Sel("new"))
	return rv
}
