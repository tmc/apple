// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightStructuralRegionGestureContainer] class.
var (
	_SkyLightStructuralRegionGestureContainerClass     SkyLightStructuralRegionGestureContainerClass
	_SkyLightStructuralRegionGestureContainerClassOnce sync.Once
)

func getSkyLightStructuralRegionGestureContainerClass() SkyLightStructuralRegionGestureContainerClass {
	_SkyLightStructuralRegionGestureContainerClassOnce.Do(func() {
		_SkyLightStructuralRegionGestureContainerClass = SkyLightStructuralRegionGestureContainerClass{class: objc.GetClass("SkyLight.StructuralRegionGestureContainer")}
	})
	return _SkyLightStructuralRegionGestureContainerClass
}

// GetSkyLightStructuralRegionGestureContainerClass returns the class object for SkyLight.StructuralRegionGestureContainer.
func GetSkyLightStructuralRegionGestureContainerClass() SkyLightStructuralRegionGestureContainerClass {
	return getSkyLightStructuralRegionGestureContainerClass()
}

type SkyLightStructuralRegionGestureContainerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightStructuralRegionGestureContainerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightStructuralRegionGestureContainerClass) Alloc() SkyLightStructuralRegionGestureContainer {
	rv := objc.SendIfResponds[SkyLightStructuralRegionGestureContainer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type SkyLightStructuralRegionGestureContainer struct {
	objectivec.Object
}

// SkyLightStructuralRegionGestureContainerFromID constructs a [SkyLightStructuralRegionGestureContainer] from an objc.ID.
func SkyLightStructuralRegionGestureContainerFromID(id objc.ID) SkyLightStructuralRegionGestureContainer {
	return SkyLightStructuralRegionGestureContainer{objectivec.Object{ID: id}}
}

// Ensure SkyLightStructuralRegionGestureContainer implements ISkyLightStructuralRegionGestureContainer.
var _ ISkyLightStructuralRegionGestureContainer = SkyLightStructuralRegionGestureContainer{}

// An interface definition for the [SkyLightStructuralRegionGestureContainer] class.
type ISkyLightStructuralRegionGestureContainer interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightStructuralRegionGestureContainer) Init() SkyLightStructuralRegionGestureContainer {
	rv := objc.SendIfResponds[SkyLightStructuralRegionGestureContainer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightStructuralRegionGestureContainer) Autorelease() SkyLightStructuralRegionGestureContainer {
	rv := objc.SendIfResponds[SkyLightStructuralRegionGestureContainer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightStructuralRegionGestureContainer creates a new SkyLightStructuralRegionGestureContainer instance.
func NewSkyLightStructuralRegionGestureContainer() SkyLightStructuralRegionGestureContainer {
	class := getSkyLightStructuralRegionGestureContainerClass()
	rv := objc.SendIfResponds[SkyLightStructuralRegionGestureContainer](objc.ID(class.class), objc.Sel("new"))
	return rv
}
