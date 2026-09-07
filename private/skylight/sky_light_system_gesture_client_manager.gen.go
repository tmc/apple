// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightSystemGestureClientManager] class.
var (
	_SkyLightSystemGestureClientManagerClass     SkyLightSystemGestureClientManagerClass
	_SkyLightSystemGestureClientManagerClassOnce sync.Once
)

func getSkyLightSystemGestureClientManagerClass() SkyLightSystemGestureClientManagerClass {
	_SkyLightSystemGestureClientManagerClassOnce.Do(func() {
		_SkyLightSystemGestureClientManagerClass = SkyLightSystemGestureClientManagerClass{class: objc.GetClass("SkyLight.SystemGestureClientManager")}
	})
	return _SkyLightSystemGestureClientManagerClass
}

// GetSkyLightSystemGestureClientManagerClass returns the class object for SkyLight.SystemGestureClientManager.
func GetSkyLightSystemGestureClientManagerClass() SkyLightSystemGestureClientManagerClass {
	return getSkyLightSystemGestureClientManagerClass()
}

type SkyLightSystemGestureClientManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightSystemGestureClientManagerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightSystemGestureClientManagerClass) Alloc() SkyLightSystemGestureClientManager {
	rv := objc.SendIfResponds[SkyLightSystemGestureClientManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type SkyLightSystemGestureClientManager struct {
	objectivec.Object
}

// SkyLightSystemGestureClientManagerFromID constructs a [SkyLightSystemGestureClientManager] from an objc.ID.
func SkyLightSystemGestureClientManagerFromID(id objc.ID) SkyLightSystemGestureClientManager {
	return SkyLightSystemGestureClientManager{objectivec.Object{ID: id}}
}

// Ensure SkyLightSystemGestureClientManager implements ISkyLightSystemGestureClientManager.
var _ ISkyLightSystemGestureClientManager = SkyLightSystemGestureClientManager{}

// An interface definition for the [SkyLightSystemGestureClientManager] class.
type ISkyLightSystemGestureClientManager interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightSystemGestureClientManager) Init() SkyLightSystemGestureClientManager {
	rv := objc.SendIfResponds[SkyLightSystemGestureClientManager](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightSystemGestureClientManager) Autorelease() SkyLightSystemGestureClientManager {
	rv := objc.SendIfResponds[SkyLightSystemGestureClientManager](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightSystemGestureClientManager creates a new SkyLightSystemGestureClientManager instance.
func NewSkyLightSystemGestureClientManager() SkyLightSystemGestureClientManager {
	class := getSkyLightSystemGestureClientManagerClass()
	rv := objc.SendIfResponds[SkyLightSystemGestureClientManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}
