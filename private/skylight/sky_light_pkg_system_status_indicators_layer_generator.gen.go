// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightPKGSystemStatusIndicatorsLayerGenerator] class.
var (
	_SkyLightPKGSystemStatusIndicatorsLayerGeneratorClass     SkyLightPKGSystemStatusIndicatorsLayerGeneratorClass
	_SkyLightPKGSystemStatusIndicatorsLayerGeneratorClassOnce sync.Once
)

func getSkyLightPKGSystemStatusIndicatorsLayerGeneratorClass() SkyLightPKGSystemStatusIndicatorsLayerGeneratorClass {
	_SkyLightPKGSystemStatusIndicatorsLayerGeneratorClassOnce.Do(func() {
		_SkyLightPKGSystemStatusIndicatorsLayerGeneratorClass = SkyLightPKGSystemStatusIndicatorsLayerGeneratorClass{class: objc.GetClass("SkyLight.PKGSystemStatusIndicatorsLayerGenerator")}
	})
	return _SkyLightPKGSystemStatusIndicatorsLayerGeneratorClass
}

// GetSkyLightPKGSystemStatusIndicatorsLayerGeneratorClass returns the class object for SkyLight.PKGSystemStatusIndicatorsLayerGenerator.
func GetSkyLightPKGSystemStatusIndicatorsLayerGeneratorClass() SkyLightPKGSystemStatusIndicatorsLayerGeneratorClass {
	return getSkyLightPKGSystemStatusIndicatorsLayerGeneratorClass()
}

type SkyLightPKGSystemStatusIndicatorsLayerGeneratorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightPKGSystemStatusIndicatorsLayerGeneratorClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightPKGSystemStatusIndicatorsLayerGeneratorClass) Alloc() SkyLightPKGSystemStatusIndicatorsLayerGenerator {
	rv := objc.SendIfResponds[SkyLightPKGSystemStatusIndicatorsLayerGenerator](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type SkyLightPKGSystemStatusIndicatorsLayerGenerator struct {
	objectivec.Object
}

// SkyLightPKGSystemStatusIndicatorsLayerGeneratorFromID constructs a [SkyLightPKGSystemStatusIndicatorsLayerGenerator] from an objc.ID.
func SkyLightPKGSystemStatusIndicatorsLayerGeneratorFromID(id objc.ID) SkyLightPKGSystemStatusIndicatorsLayerGenerator {
	return SkyLightPKGSystemStatusIndicatorsLayerGenerator{objectivec.Object{ID: id}}
}

// Ensure SkyLightPKGSystemStatusIndicatorsLayerGenerator implements ISkyLightPKGSystemStatusIndicatorsLayerGenerator.
var _ ISkyLightPKGSystemStatusIndicatorsLayerGenerator = SkyLightPKGSystemStatusIndicatorsLayerGenerator{}

// An interface definition for the [SkyLightPKGSystemStatusIndicatorsLayerGenerator] class.
type ISkyLightPKGSystemStatusIndicatorsLayerGenerator interface {
	objectivec.IObject
}

// Init initializes the instance.
func (s SkyLightPKGSystemStatusIndicatorsLayerGenerator) Init() SkyLightPKGSystemStatusIndicatorsLayerGenerator {
	rv := objc.SendIfResponds[SkyLightPKGSystemStatusIndicatorsLayerGenerator](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightPKGSystemStatusIndicatorsLayerGenerator) Autorelease() SkyLightPKGSystemStatusIndicatorsLayerGenerator {
	rv := objc.SendIfResponds[SkyLightPKGSystemStatusIndicatorsLayerGenerator](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightPKGSystemStatusIndicatorsLayerGenerator creates a new SkyLightPKGSystemStatusIndicatorsLayerGenerator instance.
func NewSkyLightPKGSystemStatusIndicatorsLayerGenerator() SkyLightPKGSystemStatusIndicatorsLayerGenerator {
	class := getSkyLightPKGSystemStatusIndicatorsLayerGeneratorClass()
	rv := objc.SendIfResponds[SkyLightPKGSystemStatusIndicatorsLayerGenerator](objc.ID(class.class), objc.Sel("new"))
	return rv
}
