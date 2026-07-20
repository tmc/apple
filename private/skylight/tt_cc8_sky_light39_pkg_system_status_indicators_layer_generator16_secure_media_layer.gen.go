// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SecureMediaLayer] class.
var (
	_SecureMediaLayerClass     SecureMediaLayerClass
	_SecureMediaLayerClassOnce sync.Once
)

func getSecureMediaLayerClass() SecureMediaLayerClass {
	_SecureMediaLayerClassOnce.Do(func() {
		_SecureMediaLayerClass = SecureMediaLayerClass{class: objc.GetClass("_TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator16SecureMediaLayer")}
	})
	return _SecureMediaLayerClass
}

// GetSecureMediaLayerClass returns the class object for _TtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator16SecureMediaLayer.
func GetSecureMediaLayerClass() SecureMediaLayerClass {
	return getSecureMediaLayerClass()
}

type SecureMediaLayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SecureMediaLayerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SecureMediaLayerClass) Alloc() SecureMediaLayer {
	rv := objc.Send[SecureMediaLayer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

type SecureMediaLayer struct {
	MediaLayer
}

// SecureMediaLayerFromID constructs a [SecureMediaLayer] from an objc.ID.
func SecureMediaLayerFromID(id objc.ID) SecureMediaLayer {
	return SecureMediaLayer{MediaLayer: MediaLayerFromID(id)}
}

// Ensure SecureMediaLayer implements ISecureMediaLayer.
var _ ISecureMediaLayer = SecureMediaLayer{}

// An interface definition for the [SecureMediaLayer] class.
type ISecureMediaLayer interface {
	IMediaLayer
}

// Init initializes the instance.
func (s SecureMediaLayer) Init() SecureMediaLayer {
	rv := objc.Send[SecureMediaLayer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SecureMediaLayer) Autorelease() SecureMediaLayer {
	rv := objc.Send[SecureMediaLayer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSecureMediaLayer creates a new SecureMediaLayer instance.
func NewSecureMediaLayer() SecureMediaLayer {
	class := getSecureMediaLayerClass()
	rv := objc.Send[SecureMediaLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator16SecureMediaLayerWithCoder(coder objectivec.IObject) SecureMediaLayer {
	instance := getSecureMediaLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return SecureMediaLayerFromID(rv)
}

func NewTtCC8SkyLight39PKGSystemStatusIndicatorsLayerGenerator16SecureMediaLayerWithLayer(layer objectivec.IObject) SecureMediaLayer {
	instance := getSecureMediaLayerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLayer:"), layer)
	return SecureMediaLayerFromID(rv)
}
