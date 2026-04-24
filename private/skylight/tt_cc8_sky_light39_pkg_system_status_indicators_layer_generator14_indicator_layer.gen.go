// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/quartzcore"
)

// The class instance for the [TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer] class.
var (
	_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerClass     TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerClass
	_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerClassOnce sync.Once
)

func getTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerClass() TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerClass {
	_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerClassOnce.Do(func() {
		_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerClass = TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerClass{class: objc.GetClass("_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer")}
	})
	return _TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerClass
}

// GetTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerClass returns the class object for _TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer.
func GetTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerClass() TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerClass {
	return getTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerClass()
}

type TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerClass) Alloc() TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer {
	rv := objc.Send[TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer
type TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer struct {
	quartzcore.CALayer
}

// TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerFromID constructs a [TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer] from an objc.ID.
func TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerFromID(id objc.ID) TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer {
	return TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer{CALayer: quartzcore.CALayerFromID(id)}
}

// Ensure TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer implements ITtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer.
var _ ITtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer = TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer{}

// An interface definition for the [TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer] class.
//
// See: https://developer.apple.com/documentation/SkyLight/_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer
type ITtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer interface {
	quartzcore.ICALayer
}

// Init initializes the instance.
func (t TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer) Init() TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer {
	rv := objc.Send[TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer) Autorelease() TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer {
	rv := objc.Send[TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer creates a new TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer instance.
func NewTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer() TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer {
	class := getTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerClass()
	rv := objc.Send[TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer/initWithCoder:
func NewTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerWithCoder(coder objectivec.IObject) TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer {
	instance := getTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer/initWithLayer:
func NewTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerWithLayer(layer objectivec.IObject) TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer {
	instance := getTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLayer:"), layer)
	return TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerFromID(rv)
}
