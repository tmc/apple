// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [LocationLayer] class.
var (
	_LocationLayerClass     LocationLayerClass
	_LocationLayerClassOnce sync.Once
)

func getLocationLayerClass() LocationLayerClass {
	_LocationLayerClassOnce.Do(func() {
		_LocationLayerClass = LocationLayerClass{class: objc.GetClass("_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer")}
	})
	return _LocationLayerClass
}

// GetLocationLayerClass returns the class object for _TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer.
func GetLocationLayerClass() LocationLayerClass {
	return getLocationLayerClass()
}

type LocationLayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (lc LocationLayerClass) Class() objc.Class {
	return lc.class
}

// Alloc allocates memory for a new instance of the class.
func (lc LocationLayerClass) Alloc() LocationLayer {
	rv := objc.SendIfResponds[LocationLayer](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

type LocationLayer struct {
	IndicatorLayer
}

// LocationLayerFromID constructs a [LocationLayer] from an objc.ID.
func LocationLayerFromID(id objc.ID) LocationLayer {
	return LocationLayer{IndicatorLayer: IndicatorLayerFromID(id)}
}

// Ensure LocationLayer implements ILocationLayer.
var _ ILocationLayer = LocationLayer{}

// An interface definition for the [LocationLayer] class.
type ILocationLayer interface {
	IIndicatorLayer
}

// Init initializes the instance.
func (l LocationLayer) Init() LocationLayer {
	rv := objc.SendIfResponds[LocationLayer](l.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l LocationLayer) Autorelease() LocationLayer {
	rv := objc.SendIfResponds[LocationLayer](l.ID, objc.Sel("autorelease"))
	return rv
}

// NewLocationLayer creates a new LocationLayer instance.
func NewLocationLayer() LocationLayer {
	class := getLocationLayerClass()
	rv := objc.SendIfResponds[LocationLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerWithCoder(coder objectivec.IObject) LocationLayer {
	instance := getLocationLayerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return LocationLayerFromID(rv)
}

func NewTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerWithLayer(layer objectivec.IObject) LocationLayer {
	instance := getLocationLayerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithLayer:"), layer)
	return LocationLayerFromID(rv)
}
