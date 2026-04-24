// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightPKGSystemStatusIndicatorsDisplayItemGenerator] class.
var (
	_SkyLightPKGSystemStatusIndicatorsDisplayItemGeneratorClass     SkyLightPKGSystemStatusIndicatorsDisplayItemGeneratorClass
	_SkyLightPKGSystemStatusIndicatorsDisplayItemGeneratorClassOnce sync.Once
)

func getSkyLightPKGSystemStatusIndicatorsDisplayItemGeneratorClass() SkyLightPKGSystemStatusIndicatorsDisplayItemGeneratorClass {
	_SkyLightPKGSystemStatusIndicatorsDisplayItemGeneratorClassOnce.Do(func() {
		_SkyLightPKGSystemStatusIndicatorsDisplayItemGeneratorClass = SkyLightPKGSystemStatusIndicatorsDisplayItemGeneratorClass{class: objc.GetClass("SkyLight.PKGSystemStatusIndicatorsDisplayItemGenerator")}
	})
	return _SkyLightPKGSystemStatusIndicatorsDisplayItemGeneratorClass
}

// GetSkyLightPKGSystemStatusIndicatorsDisplayItemGeneratorClass returns the class object for SkyLight.PKGSystemStatusIndicatorsDisplayItemGenerator.
func GetSkyLightPKGSystemStatusIndicatorsDisplayItemGeneratorClass() SkyLightPKGSystemStatusIndicatorsDisplayItemGeneratorClass {
	return getSkyLightPKGSystemStatusIndicatorsDisplayItemGeneratorClass()
}

type SkyLightPKGSystemStatusIndicatorsDisplayItemGeneratorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightPKGSystemStatusIndicatorsDisplayItemGeneratorClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightPKGSystemStatusIndicatorsDisplayItemGeneratorClass) Alloc() SkyLightPKGSystemStatusIndicatorsDisplayItemGenerator {
	rv := objc.Send[SkyLightPKGSystemStatusIndicatorsDisplayItemGenerator](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/SkyLight.PKGSystemStatusIndicatorsDisplayItemGenerator
type SkyLightPKGSystemStatusIndicatorsDisplayItemGenerator struct {
	objectivec.Object
}

// SkyLightPKGSystemStatusIndicatorsDisplayItemGeneratorFromID constructs a [SkyLightPKGSystemStatusIndicatorsDisplayItemGenerator] from an objc.ID.
func SkyLightPKGSystemStatusIndicatorsDisplayItemGeneratorFromID(id objc.ID) SkyLightPKGSystemStatusIndicatorsDisplayItemGenerator {
	return SkyLightPKGSystemStatusIndicatorsDisplayItemGenerator{objectivec.Object{ID: id}}
}

// NOTE: SkyLightPKGSystemStatusIndicatorsDisplayItemGenerator struct embeds objectivec.Object (parent type unavailable) but
// ISkyLightPKGSystemStatusIndicatorsDisplayItemGenerator embeds the parent interface; skip compile-time assertion.

// An interface definition for the [SkyLightPKGSystemStatusIndicatorsDisplayItemGenerator] class.
//
// See: https://developer.apple.com/documentation/SkyLight/SkyLight.PKGSystemStatusIndicatorsDisplayItemGenerator
type ISkyLightPKGSystemStatusIndicatorsDisplayItemGenerator interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightPKGSystemStatusIndicatorsDisplayItemGenerator) Init() SkyLightPKGSystemStatusIndicatorsDisplayItemGenerator {
	rv := objc.Send[SkyLightPKGSystemStatusIndicatorsDisplayItemGenerator](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightPKGSystemStatusIndicatorsDisplayItemGenerator) Autorelease() SkyLightPKGSystemStatusIndicatorsDisplayItemGenerator {
	rv := objc.Send[SkyLightPKGSystemStatusIndicatorsDisplayItemGenerator](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightPKGSystemStatusIndicatorsDisplayItemGenerator creates a new SkyLightPKGSystemStatusIndicatorsDisplayItemGenerator instance.
func NewSkyLightPKGSystemStatusIndicatorsDisplayItemGenerator() SkyLightPKGSystemStatusIndicatorsDisplayItemGenerator {
	class := getSkyLightPKGSystemStatusIndicatorsDisplayItemGeneratorClass()
	rv := objc.Send[SkyLightPKGSystemStatusIndicatorsDisplayItemGenerator](objc.ID(class.class), objc.Sel("new"))
	return rv
}
