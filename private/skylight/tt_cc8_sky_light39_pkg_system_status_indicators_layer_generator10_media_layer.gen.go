// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MediaLayer] class.
var (
	_MediaLayerClass     MediaLayerClass
	_MediaLayerClassOnce sync.Once
)

func getMediaLayerClass() MediaLayerClass {
	_MediaLayerClassOnce.Do(func() {
		_MediaLayerClass = MediaLayerClass{class: objc.GetClass("_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer")}
	})
	return _MediaLayerClass
}

// GetMediaLayerClass returns the class object for _TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayer.
func GetMediaLayerClass() MediaLayerClass {
	return getMediaLayerClass()
}

type MediaLayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MediaLayerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MediaLayerClass) Alloc() MediaLayer {
	rv := objc.SendIfResponds[MediaLayer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

type MediaLayer struct {
	IndicatorLayer
}

// MediaLayerFromID constructs a [MediaLayer] from an objc.ID.
func MediaLayerFromID(id objc.ID) MediaLayer {
	return MediaLayer{IndicatorLayer: IndicatorLayerFromID(id)}
}

// Ensure MediaLayer implements IMediaLayer.
var _ IMediaLayer = MediaLayer{}

// An interface definition for the [MediaLayer] class.
type IMediaLayer interface {
	IIndicatorLayer
}

// Init initializes the instance.
func (m MediaLayer) Init() MediaLayer {
	rv := objc.SendIfResponds[MediaLayer](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MediaLayer) Autorelease() MediaLayer {
	rv := objc.SendIfResponds[MediaLayer](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMediaLayer creates a new MediaLayer instance.
func NewMediaLayer() MediaLayer {
	class := getMediaLayerClass()
	rv := objc.SendIfResponds[MediaLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerWithCoder(coder objectivec.IObject) MediaLayer {
	instance := getMediaLayerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MediaLayerFromID(rv)
}

func NewTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator10MediaLayerWithLayer(layer objectivec.IObject) MediaLayer {
	instance := getMediaLayerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithLayer:"), layer)
	return MediaLayerFromID(rv)
}
