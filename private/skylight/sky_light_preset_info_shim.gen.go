// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightPresetInfoShim] class.
var (
	_SkyLightPresetInfoShimClass     SkyLightPresetInfoShimClass
	_SkyLightPresetInfoShimClassOnce sync.Once
)

func getSkyLightPresetInfoShimClass() SkyLightPresetInfoShimClass {
	_SkyLightPresetInfoShimClassOnce.Do(func() {
		_SkyLightPresetInfoShimClass = SkyLightPresetInfoShimClass{class: objc.GetClass("SkyLight.PresetInfoShim")}
	})
	return _SkyLightPresetInfoShimClass
}

// GetSkyLightPresetInfoShimClass returns the class object for SkyLight.PresetInfoShim.
func GetSkyLightPresetInfoShimClass() SkyLightPresetInfoShimClass {
	return getSkyLightPresetInfoShimClass()
}

type SkyLightPresetInfoShimClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightPresetInfoShimClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightPresetInfoShimClass) Alloc() SkyLightPresetInfoShim {
	rv := objc.Send[SkyLightPresetInfoShim](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SkyLight.PresetInfoShim
type SkyLightPresetInfoShim struct {
	objectivec.Object
}

// SkyLightPresetInfoShimFromID constructs a [SkyLightPresetInfoShim] from an objc.ID.
func SkyLightPresetInfoShimFromID(id objc.ID) SkyLightPresetInfoShim {
	return SkyLightPresetInfoShim{objectivec.Object{ID: id}}
}

// NOTE: SkyLightPresetInfoShim struct embeds objectivec.Object (parent type unavailable) but
// ISkyLightPresetInfoShim embeds the parent interface; skip compile-time assertion.

// An interface definition for the [SkyLightPresetInfoShim] class.
//
// See: https://developer.apple.com/documentation/SkyLight/SkyLight.PresetInfoShim
type ISkyLightPresetInfoShim interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightPresetInfoShim) Init() SkyLightPresetInfoShim {
	rv := objc.Send[SkyLightPresetInfoShim](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightPresetInfoShim) Autorelease() SkyLightPresetInfoShim {
	rv := objc.Send[SkyLightPresetInfoShim](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightPresetInfoShim creates a new SkyLightPresetInfoShim instance.
func NewSkyLightPresetInfoShim() SkyLightPresetInfoShim {
	class := getSkyLightPresetInfoShimClass()
	rv := objc.Send[SkyLightPresetInfoShim](objc.ID(class.class), objc.Sel("new"))
	return rv
}
