// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightPKGGestureCoordinator] class.
var (
	_SkyLightPKGGestureCoordinatorClass     SkyLightPKGGestureCoordinatorClass
	_SkyLightPKGGestureCoordinatorClassOnce sync.Once
)

func getSkyLightPKGGestureCoordinatorClass() SkyLightPKGGestureCoordinatorClass {
	_SkyLightPKGGestureCoordinatorClassOnce.Do(func() {
		_SkyLightPKGGestureCoordinatorClass = SkyLightPKGGestureCoordinatorClass{class: objc.GetClass("SkyLight.PKGGestureCoordinator")}
	})
	return _SkyLightPKGGestureCoordinatorClass
}

// GetSkyLightPKGGestureCoordinatorClass returns the class object for SkyLight.PKGGestureCoordinator.
func GetSkyLightPKGGestureCoordinatorClass() SkyLightPKGGestureCoordinatorClass {
	return getSkyLightPKGGestureCoordinatorClass()
}

type SkyLightPKGGestureCoordinatorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightPKGGestureCoordinatorClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightPKGGestureCoordinatorClass) Alloc() SkyLightPKGGestureCoordinator {
	rv := objc.SendIfResponds[SkyLightPKGGestureCoordinator](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type SkyLightPKGGestureCoordinator struct {
	objectivec.Object
}

// SkyLightPKGGestureCoordinatorFromID constructs a [SkyLightPKGGestureCoordinator] from an objc.ID.
func SkyLightPKGGestureCoordinatorFromID(id objc.ID) SkyLightPKGGestureCoordinator {
	return SkyLightPKGGestureCoordinator{objectivec.Object{ID: id}}
}

// Ensure SkyLightPKGGestureCoordinator implements ISkyLightPKGGestureCoordinator.
var _ ISkyLightPKGGestureCoordinator = SkyLightPKGGestureCoordinator{}

// An interface definition for the [SkyLightPKGGestureCoordinator] class.
type ISkyLightPKGGestureCoordinator interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightPKGGestureCoordinator) Init() SkyLightPKGGestureCoordinator {
	rv := objc.SendIfResponds[SkyLightPKGGestureCoordinator](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightPKGGestureCoordinator) Autorelease() SkyLightPKGGestureCoordinator {
	rv := objc.SendIfResponds[SkyLightPKGGestureCoordinator](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightPKGGestureCoordinator creates a new SkyLightPKGGestureCoordinator instance.
func NewSkyLightPKGGestureCoordinator() SkyLightPKGGestureCoordinator {
	class := getSkyLightPKGGestureCoordinatorClass()
	rv := objc.SendIfResponds[SkyLightPKGGestureCoordinator](objc.ID(class.class), objc.Sel("new"))
	return rv
}
