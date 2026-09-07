// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightSystemGestureClientProcess] class.
var (
	_SkyLightSystemGestureClientProcessClass     SkyLightSystemGestureClientProcessClass
	_SkyLightSystemGestureClientProcessClassOnce sync.Once
)

func getSkyLightSystemGestureClientProcessClass() SkyLightSystemGestureClientProcessClass {
	_SkyLightSystemGestureClientProcessClassOnce.Do(func() {
		_SkyLightSystemGestureClientProcessClass = SkyLightSystemGestureClientProcessClass{class: objc.GetClass("SkyLight.SystemGestureClientProcess")}
	})
	return _SkyLightSystemGestureClientProcessClass
}

// GetSkyLightSystemGestureClientProcessClass returns the class object for SkyLight.SystemGestureClientProcess.
func GetSkyLightSystemGestureClientProcessClass() SkyLightSystemGestureClientProcessClass {
	return getSkyLightSystemGestureClientProcessClass()
}

type SkyLightSystemGestureClientProcessClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightSystemGestureClientProcessClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightSystemGestureClientProcessClass) Alloc() SkyLightSystemGestureClientProcess {
	rv := objc.SendIfResponds[SkyLightSystemGestureClientProcess](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type SkyLightSystemGestureClientProcess struct {
	objectivec.Object
}

// SkyLightSystemGestureClientProcessFromID constructs a [SkyLightSystemGestureClientProcess] from an objc.ID.
func SkyLightSystemGestureClientProcessFromID(id objc.ID) SkyLightSystemGestureClientProcess {
	return SkyLightSystemGestureClientProcess{objectivec.Object{ID: id}}
}

// Ensure SkyLightSystemGestureClientProcess implements ISkyLightSystemGestureClientProcess.
var _ ISkyLightSystemGestureClientProcess = SkyLightSystemGestureClientProcess{}

// An interface definition for the [SkyLightSystemGestureClientProcess] class.
type ISkyLightSystemGestureClientProcess interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightSystemGestureClientProcess) Init() SkyLightSystemGestureClientProcess {
	rv := objc.SendIfResponds[SkyLightSystemGestureClientProcess](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightSystemGestureClientProcess) Autorelease() SkyLightSystemGestureClientProcess {
	rv := objc.SendIfResponds[SkyLightSystemGestureClientProcess](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightSystemGestureClientProcess creates a new SkyLightSystemGestureClientProcess instance.
func NewSkyLightSystemGestureClientProcess() SkyLightSystemGestureClientProcess {
	class := getSkyLightSystemGestureClientProcessClass()
	rv := objc.SendIfResponds[SkyLightSystemGestureClientProcess](objc.ID(class.class), objc.Sel("new"))
	return rv
}
