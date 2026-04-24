// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightPresetSpecificationsShim] class.
var (
	_SkyLightPresetSpecificationsShimClass     SkyLightPresetSpecificationsShimClass
	_SkyLightPresetSpecificationsShimClassOnce sync.Once
)

func getSkyLightPresetSpecificationsShimClass() SkyLightPresetSpecificationsShimClass {
	_SkyLightPresetSpecificationsShimClassOnce.Do(func() {
		_SkyLightPresetSpecificationsShimClass = SkyLightPresetSpecificationsShimClass{class: objc.GetClass("SkyLight.PresetSpecificationsShim")}
	})
	return _SkyLightPresetSpecificationsShimClass
}

// GetSkyLightPresetSpecificationsShimClass returns the class object for SkyLight.PresetSpecificationsShim.
func GetSkyLightPresetSpecificationsShimClass() SkyLightPresetSpecificationsShimClass {
	return getSkyLightPresetSpecificationsShimClass()
}

type SkyLightPresetSpecificationsShimClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightPresetSpecificationsShimClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightPresetSpecificationsShimClass) Alloc() SkyLightPresetSpecificationsShim {
	rv := objc.Send[SkyLightPresetSpecificationsShim](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SkyLight.PresetSpecificationsShim
type SkyLightPresetSpecificationsShim struct {
	objectivec.Object
}

// SkyLightPresetSpecificationsShimFromID constructs a [SkyLightPresetSpecificationsShim] from an objc.ID.
func SkyLightPresetSpecificationsShimFromID(id objc.ID) SkyLightPresetSpecificationsShim {
	return SkyLightPresetSpecificationsShim{objectivec.Object{ID: id}}
}

// NOTE: SkyLightPresetSpecificationsShim struct embeds objectivec.Object (parent type unavailable) but
// ISkyLightPresetSpecificationsShim embeds the parent interface; skip compile-time assertion.

// An interface definition for the [SkyLightPresetSpecificationsShim] class.
//
// See: https://developer.apple.com/documentation/SkyLight/SkyLight.PresetSpecificationsShim
type ISkyLightPresetSpecificationsShim interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightPresetSpecificationsShim) Init() SkyLightPresetSpecificationsShim {
	rv := objc.Send[SkyLightPresetSpecificationsShim](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightPresetSpecificationsShim) Autorelease() SkyLightPresetSpecificationsShim {
	rv := objc.Send[SkyLightPresetSpecificationsShim](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightPresetSpecificationsShim creates a new SkyLightPresetSpecificationsShim instance.
func NewSkyLightPresetSpecificationsShim() SkyLightPresetSpecificationsShim {
	class := getSkyLightPresetSpecificationsShimClass()
	rv := objc.Send[SkyLightPresetSpecificationsShim](objc.ID(class.class), objc.Sel("new"))
	return rv
}
