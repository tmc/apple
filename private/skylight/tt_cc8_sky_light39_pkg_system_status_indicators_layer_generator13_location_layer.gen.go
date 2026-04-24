// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer] class.
var (
	_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerClass     TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerClass
	_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerClassOnce sync.Once
)

func getTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerClass() TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerClass {
	_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerClassOnce.Do(func() {
		_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerClass = TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerClass{class: objc.GetClass("_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer")}
	})
	return _TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerClass
}

// GetTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerClass returns the class object for _TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer.
func GetTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerClass() TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerClass {
	return getTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerClass()
}

type TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (tc TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerClass) Class() objc.Class {
	return tc.class
}

// Alloc allocates memory for a new instance of the class.
func (tc TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerClass) Alloc() TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer {
	rv := objc.Send[TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer
type TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer struct {
	TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer
}

// TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerFromID constructs a [TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer] from an objc.ID.
func TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerFromID(id objc.ID) TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer {
	return TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer{TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer: TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayerFromID(id)}
}

// Ensure TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer implements ITtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer.
var _ ITtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer = TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer{}

// An interface definition for the [TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer] class.
//
// See: https://developer.apple.com/documentation/SkyLight/_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer
type ITtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer interface {
	ITtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer
}

// Init initializes the instance.
func (t TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer) Init() TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer {
	rv := objc.Send[TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer) Autorelease() TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer {
	rv := objc.Send[TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer creates a new TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer instance.
func NewTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer() TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer {
	class := getTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerClass()
	rv := objc.Send[TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/SkyLight/_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer/initWithCoder:
func NewTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerWithCoder(coder objectivec.IObject) TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer {
	instance := getTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerFromID(rv)
}

// See: https://developer.apple.com/documentation/SkyLight/_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator14IndicatorLayer/initWithLayer:
func NewTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerWithLayer(layer objectivec.IObject) TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayer {
	instance := getTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLayer:"), layer)
	return TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator13LocationLayerFromID(rv)
}
