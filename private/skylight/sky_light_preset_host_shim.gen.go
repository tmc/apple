// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightPresetHostShim] class.
var (
	_SkyLightPresetHostShimClass     SkyLightPresetHostShimClass
	_SkyLightPresetHostShimClassOnce sync.Once
)

func getSkyLightPresetHostShimClass() SkyLightPresetHostShimClass {
	_SkyLightPresetHostShimClassOnce.Do(func() {
		_SkyLightPresetHostShimClass = SkyLightPresetHostShimClass{class: objc.GetClass("SkyLight.PresetHostShim")}
	})
	return _SkyLightPresetHostShimClass
}

// GetSkyLightPresetHostShimClass returns the class object for SkyLight.PresetHostShim.
func GetSkyLightPresetHostShimClass() SkyLightPresetHostShimClass {
	return getSkyLightPresetHostShimClass()
}

type SkyLightPresetHostShimClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightPresetHostShimClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightPresetHostShimClass) Alloc() SkyLightPresetHostShim {
	rv := objc.Send[SkyLightPresetHostShim](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SkyLight.PresetHostShim
type SkyLightPresetHostShim struct {
	objectivec.Object
}

// SkyLightPresetHostShimFromID constructs a [SkyLightPresetHostShim] from an objc.ID.
func SkyLightPresetHostShimFromID(id objc.ID) SkyLightPresetHostShim {
	return SkyLightPresetHostShim{objectivec.Object{ID: id}}
}

// NOTE: SkyLightPresetHostShim struct embeds objectivec.Object (parent type unavailable) but
// ISkyLightPresetHostShim embeds the parent interface; skip compile-time assertion.

// An interface definition for the [SkyLightPresetHostShim] class.
//
// See: https://developer.apple.com/documentation/SkyLight/SkyLight.PresetHostShim
type ISkyLightPresetHostShim interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightPresetHostShim) Init() SkyLightPresetHostShim {
	rv := objc.Send[SkyLightPresetHostShim](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightPresetHostShim) Autorelease() SkyLightPresetHostShim {
	rv := objc.Send[SkyLightPresetHostShim](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightPresetHostShim creates a new SkyLightPresetHostShim instance.
func NewSkyLightPresetHostShim() SkyLightPresetHostShim {
	class := getSkyLightPresetHostShimClass()
	rv := objc.Send[SkyLightPresetHostShim](objc.ID(class.class), objc.Sel("new"))
	return rv
}
