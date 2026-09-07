// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTShaderProfilerRegisterPressureView] class.
var (
	_GTShaderProfilerRegisterPressureViewClass     GTShaderProfilerRegisterPressureViewClass
	_GTShaderProfilerRegisterPressureViewClassOnce sync.Once
)

func getGTShaderProfilerRegisterPressureViewClass() GTShaderProfilerRegisterPressureViewClass {
	_GTShaderProfilerRegisterPressureViewClassOnce.Do(func() {
		_GTShaderProfilerRegisterPressureViewClass = GTShaderProfilerRegisterPressureViewClass{class: objc.GetClass("GTShaderProfilerRegisterPressureView")}
	})
	return _GTShaderProfilerRegisterPressureViewClass
}

// GetGTShaderProfilerRegisterPressureViewClass returns the class object for GTShaderProfilerRegisterPressureView.
func GetGTShaderProfilerRegisterPressureViewClass() GTShaderProfilerRegisterPressureViewClass {
	return getGTShaderProfilerRegisterPressureViewClass()
}

type GTShaderProfilerRegisterPressureViewClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTShaderProfilerRegisterPressureViewClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTShaderProfilerRegisterPressureViewClass) Alloc() GTShaderProfilerRegisterPressureView {
	rv := objc.SendIfResponds[GTShaderProfilerRegisterPressureView](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTShaderProfilerRegisterPressureView.Binary]
//   - [GTShaderProfilerRegisterPressureView.LoadFromDict]
//   - [GTShaderProfilerRegisterPressureView.MaxTheoriticalOccupancy]
//   - [GTShaderProfilerRegisterPressureView.RegisterPressureForAddress]
//   - [GTShaderProfilerRegisterPressureView.InitWithDictBinaryGpu]
type GTShaderProfilerRegisterPressureView struct {
	objectivec.Object
}

// GTShaderProfilerRegisterPressureViewFromID constructs a [GTShaderProfilerRegisterPressureView] from an objc.ID.
func GTShaderProfilerRegisterPressureViewFromID(id objc.ID) GTShaderProfilerRegisterPressureView {
	return GTShaderProfilerRegisterPressureView{objectivec.Object{ID: id}}
}

// Ensure GTShaderProfilerRegisterPressureView implements IGTShaderProfilerRegisterPressureView.
var _ IGTShaderProfilerRegisterPressureView = GTShaderProfilerRegisterPressureView{}

// An interface definition for the [GTShaderProfilerRegisterPressureView] class.
//
// # Methods
//
//   - [IGTShaderProfilerRegisterPressureView.Binary]
//   - [IGTShaderProfilerRegisterPressureView.LoadFromDict]
//   - [IGTShaderProfilerRegisterPressureView.MaxTheoriticalOccupancy]
//   - [IGTShaderProfilerRegisterPressureView.RegisterPressureForAddress]
//   - [IGTShaderProfilerRegisterPressureView.InitWithDictBinaryGpu]
type IGTShaderProfilerRegisterPressureView interface {
	objectivec.IObject

	// Topic: Methods

	Binary() unsafe.Pointer
	LoadFromDict(dict objectivec.IObject)
	MaxTheoriticalOccupancy() float32
	RegisterPressureForAddress(address uint32) objectivec.IObject
	InitWithDictBinaryGpu(dict objectivec.IObject, binary objectivec.IObject, gpu uint32) GTShaderProfilerRegisterPressureView
}

// Init initializes the instance.
func (g GTShaderProfilerRegisterPressureView) Init() GTShaderProfilerRegisterPressureView {
	rv := objc.SendIfResponds[GTShaderProfilerRegisterPressureView](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTShaderProfilerRegisterPressureView) Autorelease() GTShaderProfilerRegisterPressureView {
	rv := objc.SendIfResponds[GTShaderProfilerRegisterPressureView](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTShaderProfilerRegisterPressureView creates a new GTShaderProfilerRegisterPressureView instance.
func NewGTShaderProfilerRegisterPressureView() GTShaderProfilerRegisterPressureView {
	class := getGTShaderProfilerRegisterPressureViewClass()
	rv := objc.SendIfResponds[GTShaderProfilerRegisterPressureView](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGTShaderProfilerRegisterPressureViewWithDictBinaryGpu(dict objectivec.IObject, binary objectivec.IObject, gpu uint32) GTShaderProfilerRegisterPressureView {
	instance := getGTShaderProfilerRegisterPressureViewClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDict:binary:gpu:"), dict, binary, gpu)
	return GTShaderProfilerRegisterPressureViewFromID(rv)
}

func (g GTShaderProfilerRegisterPressureView) LoadFromDict(dict objectivec.IObject) {
	objc.SendIfResponds[objc.ID](g.ID, objc.Sel("loadFromDict:"), dict)
}
func (g GTShaderProfilerRegisterPressureView) MaxTheoriticalOccupancy() float32 {
	rv := objc.SendIfResponds[float32](g.ID, objc.Sel("maxTheoriticalOccupancy"))
	return rv
}
func (g GTShaderProfilerRegisterPressureView) RegisterPressureForAddress(address uint32) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](g.ID, objc.Sel("registerPressureForAddress:"), address)
	return objectivec.Object{ID: rv}
}
func (g GTShaderProfilerRegisterPressureView) InitWithDictBinaryGpu(dict objectivec.IObject, binary objectivec.IObject, gpu uint32) GTShaderProfilerRegisterPressureView {
	rv := objc.SendIfResponds[GTShaderProfilerRegisterPressureView](g.ID, objc.Sel("initWithDict:binary:gpu:"), dict, binary, gpu)
	return rv
}

func (_GTShaderProfilerRegisterPressureViewClass GTShaderProfilerRegisterPressureViewClass) MaxTheoriticalOccupancyWithRegisterCountGpu(count uint32, gpu uint32) float32 {
	rv := objc.SendIfResponds[float32](objc.ID(_GTShaderProfilerRegisterPressureViewClass.class), objc.Sel("maxTheoriticalOccupancyWithRegisterCount:gpu:"), count, gpu)
	return rv
}

func (g GTShaderProfilerRegisterPressureView) Binary() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](g.ID, objc.Sel("binary"))
	return rv
}
