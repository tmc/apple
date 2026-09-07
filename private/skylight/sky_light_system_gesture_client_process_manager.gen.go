// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightSystemGestureClientProcessManager] class.
var (
	_SkyLightSystemGestureClientProcessManagerClass     SkyLightSystemGestureClientProcessManagerClass
	_SkyLightSystemGestureClientProcessManagerClassOnce sync.Once
)

func getSkyLightSystemGestureClientProcessManagerClass() SkyLightSystemGestureClientProcessManagerClass {
	_SkyLightSystemGestureClientProcessManagerClassOnce.Do(func() {
		_SkyLightSystemGestureClientProcessManagerClass = SkyLightSystemGestureClientProcessManagerClass{class: objc.GetClass("SkyLight.SystemGestureClientProcessManager")}
	})
	return _SkyLightSystemGestureClientProcessManagerClass
}

// GetSkyLightSystemGestureClientProcessManagerClass returns the class object for SkyLight.SystemGestureClientProcessManager.
func GetSkyLightSystemGestureClientProcessManagerClass() SkyLightSystemGestureClientProcessManagerClass {
	return getSkyLightSystemGestureClientProcessManagerClass()
}

type SkyLightSystemGestureClientProcessManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightSystemGestureClientProcessManagerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightSystemGestureClientProcessManagerClass) Alloc() SkyLightSystemGestureClientProcessManager {
	rv := objc.SendIfResponds[SkyLightSystemGestureClientProcessManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type SkyLightSystemGestureClientProcessManager struct {
	objectivec.Object
}

// SkyLightSystemGestureClientProcessManagerFromID constructs a [SkyLightSystemGestureClientProcessManager] from an objc.ID.
func SkyLightSystemGestureClientProcessManagerFromID(id objc.ID) SkyLightSystemGestureClientProcessManager {
	return SkyLightSystemGestureClientProcessManager{objectivec.Object{ID: id}}
}

// Ensure SkyLightSystemGestureClientProcessManager implements ISkyLightSystemGestureClientProcessManager.
var _ ISkyLightSystemGestureClientProcessManager = SkyLightSystemGestureClientProcessManager{}

// An interface definition for the [SkyLightSystemGestureClientProcessManager] class.
type ISkyLightSystemGestureClientProcessManager interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightSystemGestureClientProcessManager) Init() SkyLightSystemGestureClientProcessManager {
	rv := objc.SendIfResponds[SkyLightSystemGestureClientProcessManager](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightSystemGestureClientProcessManager) Autorelease() SkyLightSystemGestureClientProcessManager {
	rv := objc.SendIfResponds[SkyLightSystemGestureClientProcessManager](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightSystemGestureClientProcessManager creates a new SkyLightSystemGestureClientProcessManager instance.
func NewSkyLightSystemGestureClientProcessManager() SkyLightSystemGestureClientProcessManager {
	class := getSkyLightSystemGestureClientProcessManagerClass()
	rv := objc.SendIfResponds[SkyLightSystemGestureClientProcessManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}
