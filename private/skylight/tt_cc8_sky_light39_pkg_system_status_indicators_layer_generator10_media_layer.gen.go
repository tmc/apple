// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer] class.
var (
	_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerClass     TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerClass
	_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerClassOnce sync.Once
)

func getTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerClass() TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerClass {
	_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerClassOnce.Do(func() {
		_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerClass = TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerClass{class: objc.GetClass("_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer")}
	})
	return _TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerClass
}

// GetTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerClass returns the class object for _TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer.
func GetTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerClass() TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerClass {
	return getTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerClass()
}

type TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerClass) Alloc() TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer {
	rv := objc.Send[TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer
type TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer struct {
	TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer
}

// TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerFromID constructs a [TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer] from an objc.ID.
func TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerFromID(id objc.ID) TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer {
	return TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer{TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer: TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerFromID(id)}
}

// Ensure TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer implements ITtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer.
var _ ITtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer = TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer{}

// An interface definition for the [TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer] class.
//
// See: https://developer.apple.com/documentation/SkyLight/_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer
type ITtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer interface {
	ITtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer
}

// Init initializes the instance.
func (t TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer) Init() TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer {
	rv := objc.Send[TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer) Autorelease() TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer {
	rv := objc.Send[TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer creates a new TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer instance.
func NewTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer() TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer {
	class := getTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerClass()
	rv := objc.Send[TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer/initWithCoder:
func NewTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerWithCoder(coder objectivec.IObject) TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer {
	instance := getTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer/initWithLayer:
func NewTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerWithLayer(layer objectivec.IObject) TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer {
	instance := getTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLayer:"), layer)
	return TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerFromID(rv)
}
