// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/quartzcore"
)

// The class instance for the [IndicatorLayer] class.
var (
	_IndicatorLayerClass     IndicatorLayerClass
	_IndicatorLayerClassOnce sync.Once
)

func getIndicatorLayerClass() IndicatorLayerClass {
	_IndicatorLayerClassOnce.Do(func() {
		_IndicatorLayerClass = IndicatorLayerClass{class: objc.GetClass("_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer")}
	})
	return _IndicatorLayerClass
}

// GetIndicatorLayerClass returns the class object for _TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer.
func GetIndicatorLayerClass() IndicatorLayerClass {
	return getIndicatorLayerClass()
}

type IndicatorLayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IndicatorLayerClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IndicatorLayerClass) Alloc() IndicatorLayer {
	rv := objc.Send[IndicatorLayer](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

type IndicatorLayer struct {
	quartzcore.CALayer
}

// IndicatorLayerFromID constructs a [IndicatorLayer] from an objc.ID.
func IndicatorLayerFromID(id objc.ID) IndicatorLayer {
	return IndicatorLayer{CALayer: quartzcore.CALayerFromID(id)}
}

// Ensure IndicatorLayer implements IIndicatorLayer.
var _ IIndicatorLayer = IndicatorLayer{}

// An interface definition for the [IndicatorLayer] class.
type IIndicatorLayer interface {
	quartzcore.ICALayer
}

// Init initializes the instance.
func (i IndicatorLayer) Init() IndicatorLayer {
	rv := objc.Send[IndicatorLayer](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IndicatorLayer) Autorelease() IndicatorLayer {
	rv := objc.Send[IndicatorLayer](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIndicatorLayer creates a new IndicatorLayer instance.
func NewIndicatorLayer() IndicatorLayer {
	class := getIndicatorLayerClass()
	rv := objc.Send[IndicatorLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerWithCoder(coder objectivec.IObject) IndicatorLayer {
	instance := getIndicatorLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return IndicatorLayerFromID(rv)
}

func NewTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerWithLayer(layer objectivec.IObject) IndicatorLayer {
	instance := getIndicatorLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLayer:"), layer)
	return IndicatorLayerFromID(rv)
}
