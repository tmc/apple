// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightSLSMenuBarAgentManager] class.
var (
	_SkyLightSLSMenuBarAgentManagerClass     SkyLightSLSMenuBarAgentManagerClass
	_SkyLightSLSMenuBarAgentManagerClassOnce sync.Once
)

func getSkyLightSLSMenuBarAgentManagerClass() SkyLightSLSMenuBarAgentManagerClass {
	_SkyLightSLSMenuBarAgentManagerClassOnce.Do(func() {
		_SkyLightSLSMenuBarAgentManagerClass = SkyLightSLSMenuBarAgentManagerClass{class: objc.GetClass("SkyLight.SLSMenuBarAgentManager")}
	})
	return _SkyLightSLSMenuBarAgentManagerClass
}

// GetSkyLightSLSMenuBarAgentManagerClass returns the class object for SkyLight.SLSMenuBarAgentManager.
func GetSkyLightSLSMenuBarAgentManagerClass() SkyLightSLSMenuBarAgentManagerClass {
	return getSkyLightSLSMenuBarAgentManagerClass()
}

type SkyLightSLSMenuBarAgentManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightSLSMenuBarAgentManagerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightSLSMenuBarAgentManagerClass) Alloc() SkyLightSLSMenuBarAgentManager {
	rv := objc.SendIfResponds[SkyLightSLSMenuBarAgentManager](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type SkyLightSLSMenuBarAgentManager struct {
	objectivec.Object
}

// SkyLightSLSMenuBarAgentManagerFromID constructs a [SkyLightSLSMenuBarAgentManager] from an objc.ID.
func SkyLightSLSMenuBarAgentManagerFromID(id objc.ID) SkyLightSLSMenuBarAgentManager {
	return SkyLightSLSMenuBarAgentManager{objectivec.Object{ID: id}}
}

// Ensure SkyLightSLSMenuBarAgentManager implements ISkyLightSLSMenuBarAgentManager.
var _ ISkyLightSLSMenuBarAgentManager = SkyLightSLSMenuBarAgentManager{}

// An interface definition for the [SkyLightSLSMenuBarAgentManager] class.
type ISkyLightSLSMenuBarAgentManager interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightSLSMenuBarAgentManager) Init() SkyLightSLSMenuBarAgentManager {
	rv := objc.SendIfResponds[SkyLightSLSMenuBarAgentManager](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightSLSMenuBarAgentManager) Autorelease() SkyLightSLSMenuBarAgentManager {
	rv := objc.SendIfResponds[SkyLightSLSMenuBarAgentManager](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightSLSMenuBarAgentManager creates a new SkyLightSLSMenuBarAgentManager instance.
func NewSkyLightSLSMenuBarAgentManager() SkyLightSLSMenuBarAgentManager {
	class := getSkyLightSLSMenuBarAgentManagerClass()
	rv := objc.SendIfResponds[SkyLightSLSMenuBarAgentManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}
