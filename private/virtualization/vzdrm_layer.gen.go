// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/quartzcore"
)

// The class instance for the [VZDRMLayer] class.
var (
	_VZDRMLayerClass     VZDRMLayerClass
	_VZDRMLayerClassOnce sync.Once
)

func getVZDRMLayerClass() VZDRMLayerClass {
	_VZDRMLayerClassOnce.Do(func() {
		_VZDRMLayerClass = VZDRMLayerClass{class: objc.GetClass("_VZDRMLayer")}
	})
	return _VZDRMLayerClass
}

// GetVZDRMLayerClass returns the class object for _VZDRMLayer.
func GetVZDRMLayerClass() VZDRMLayerClass {
	return getVZDRMLayerClass()
}

type VZDRMLayerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZDRMLayerClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZDRMLayerClass) Alloc() VZDRMLayer {
	rv := objc.SendIfResponds[VZDRMLayer](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZDRMLayer.LayerDidResize]
//   - [VZDRMLayer.InitForTestingContent]
//   - [VZDRMLayer.InitWithParentLayer]
type VZDRMLayer struct {
	quartzcore.CALayer
}

// VZDRMLayerFromID constructs a [VZDRMLayer] from an objc.ID.
func VZDRMLayerFromID(id objc.ID) VZDRMLayer {
	return VZDRMLayer{CALayer: quartzcore.CALayerFromID(id)}
}

// Ensure VZDRMLayer implements IVZDRMLayer.
var _ IVZDRMLayer = VZDRMLayer{}

// An interface definition for the [VZDRMLayer] class.
//
// # Methods
//
//   - [IVZDRMLayer.LayerDidResize]
//   - [IVZDRMLayer.InitForTestingContent]
//   - [IVZDRMLayer.InitWithParentLayer]
type IVZDRMLayer interface {
	quartzcore.ICALayer

	// Topic: Methods

	LayerDidResize(resize corefoundation.CGSize)
	InitForTestingContent(testing objectivec.IObject, content objectivec.IObject) VZDRMLayer
	InitWithParentLayer(layer objectivec.IObject) VZDRMLayer
}

// Init initializes the instance.
func (v VZDRMLayer) Init() VZDRMLayer {
	rv := objc.SendIfResponds[VZDRMLayer](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZDRMLayer) Autorelease() VZDRMLayer {
	rv := objc.SendIfResponds[VZDRMLayer](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZDRMLayer creates a new VZDRMLayer instance.
func NewVZDRMLayer() VZDRMLayer {
	class := getVZDRMLayerClass()
	rv := objc.SendIfResponds[VZDRMLayer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewVZDRMLayerForTestingContent(testing objectivec.IObject, content objectivec.IObject) VZDRMLayer {
	instance := getVZDRMLayerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initForTesting:content:"), testing, content)
	return VZDRMLayerFromID(rv)
}

func NewVZDRMLayerWithParentLayer(layer objectivec.IObject) VZDRMLayer {
	instance := getVZDRMLayerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithParentLayer:"), layer)
	return VZDRMLayerFromID(rv)
}

func (v VZDRMLayer) LayerDidResize(resize corefoundation.CGSize) {
	objc.SendIfResponds[objc.ID](v.ID, objc.Sel("layerDidResize:"), resize)
}
func (v VZDRMLayer) InitForTestingContent(testing objectivec.IObject, content objectivec.IObject) VZDRMLayer {
	rv := objc.SendIfResponds[VZDRMLayer](v.ID, objc.Sel("initForTesting:content:"), testing, content)
	return rv
}
func (v VZDRMLayer) InitWithParentLayer(layer objectivec.IObject) VZDRMLayer {
	rv := objc.SendIfResponds[VZDRMLayer](v.ID, objc.Sel("initWithParentLayer:"), layer)
	return rv
}

func (_VZDRMLayerClass VZDRMLayerClass) IsSupported() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_VZDRMLayerClass.class), objc.Sel("isSupported"))
	return rv
}
