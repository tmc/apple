// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightSystemGestureClientStructuralRegion] class.
var (
	_SkyLightSystemGestureClientStructuralRegionClass     SkyLightSystemGestureClientStructuralRegionClass
	_SkyLightSystemGestureClientStructuralRegionClassOnce sync.Once
)

func getSkyLightSystemGestureClientStructuralRegionClass() SkyLightSystemGestureClientStructuralRegionClass {
	_SkyLightSystemGestureClientStructuralRegionClassOnce.Do(func() {
		_SkyLightSystemGestureClientStructuralRegionClass = SkyLightSystemGestureClientStructuralRegionClass{class: objc.GetClass("SkyLight.SystemGestureClientStructuralRegion")}
	})
	return _SkyLightSystemGestureClientStructuralRegionClass
}

// GetSkyLightSystemGestureClientStructuralRegionClass returns the class object for SkyLight.SystemGestureClientStructuralRegion.
func GetSkyLightSystemGestureClientStructuralRegionClass() SkyLightSystemGestureClientStructuralRegionClass {
	return getSkyLightSystemGestureClientStructuralRegionClass()
}

type SkyLightSystemGestureClientStructuralRegionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightSystemGestureClientStructuralRegionClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightSystemGestureClientStructuralRegionClass) Alloc() SkyLightSystemGestureClientStructuralRegion {
	rv := objc.SendIfResponds[SkyLightSystemGestureClientStructuralRegion](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type SkyLightSystemGestureClientStructuralRegion struct {
	objectivec.Object
}

// SkyLightSystemGestureClientStructuralRegionFromID constructs a [SkyLightSystemGestureClientStructuralRegion] from an objc.ID.
func SkyLightSystemGestureClientStructuralRegionFromID(id objc.ID) SkyLightSystemGestureClientStructuralRegion {
	return SkyLightSystemGestureClientStructuralRegion{objectivec.Object{ID: id}}
}

// Ensure SkyLightSystemGestureClientStructuralRegion implements ISkyLightSystemGestureClientStructuralRegion.
var _ ISkyLightSystemGestureClientStructuralRegion = SkyLightSystemGestureClientStructuralRegion{}

// An interface definition for the [SkyLightSystemGestureClientStructuralRegion] class.
type ISkyLightSystemGestureClientStructuralRegion interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightSystemGestureClientStructuralRegion) Init() SkyLightSystemGestureClientStructuralRegion {
	rv := objc.SendIfResponds[SkyLightSystemGestureClientStructuralRegion](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightSystemGestureClientStructuralRegion) Autorelease() SkyLightSystemGestureClientStructuralRegion {
	rv := objc.SendIfResponds[SkyLightSystemGestureClientStructuralRegion](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightSystemGestureClientStructuralRegion creates a new SkyLightSystemGestureClientStructuralRegion instance.
func NewSkyLightSystemGestureClientStructuralRegion() SkyLightSystemGestureClientStructuralRegion {
	class := getSkyLightSystemGestureClientStructuralRegionClass()
	rv := objc.SendIfResponds[SkyLightSystemGestureClientStructuralRegion](objc.ID(class.class), objc.Sel("new"))
	return rv
}
