// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalGLDrawable] class.
var (
	_IOGPUMetalGLDrawableClass     IOGPUMetalGLDrawableClass
	_IOGPUMetalGLDrawableClassOnce sync.Once
)

func getIOGPUMetalGLDrawableClass() IOGPUMetalGLDrawableClass {
	_IOGPUMetalGLDrawableClassOnce.Do(func() {
		_IOGPUMetalGLDrawableClass = IOGPUMetalGLDrawableClass{class: objc.GetClass("IOGPUMetalGLDrawable")}
	})
	return _IOGPUMetalGLDrawableClass
}

// GetIOGPUMetalGLDrawableClass returns the class object for IOGPUMetalGLDrawable.
func GetIOGPUMetalGLDrawableClass() IOGPUMetalGLDrawableClass {
	return getIOGPUMetalGLDrawableClass()
}

type IOGPUMetalGLDrawableClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalGLDrawableClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalGLDrawableClass) Alloc() IOGPUMetalGLDrawable {
	rv := objc.SendIfResponds[IOGPUMetalGLDrawable](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalGLDrawable.ClearDrawable]
//   - [IOGPUMetalGLDrawable.Height]
//   - [IOGPUMetalGLDrawable.LookupIOSurfaceAtIndex]
//   - [IOGPUMetalGLDrawable.SetDrawableSurfaceModeColorDepthModeFaceLevelVolatileFixedSourceScaleOptionsScaledWidthScaledHeight]
//   - [IOGPUMetalGLDrawable.SetSwapIntervalLimit]
//   - [IOGPUMetalGLDrawable.SetSwapRectXYWH]
//   - [IOGPUMetalGLDrawable.SignalSharedEventValueOperation]
//   - [IOGPUMetalGLDrawable.SurfaceHeight]
//   - [IOGPUMetalGLDrawable.SurfaceWidth]
//   - [IOGPUMetalGLDrawable.Width]
//   - [IOGPUMetalGLDrawable.WindowMode]
//   - [IOGPUMetalGLDrawable.InitWithDevice]
type IOGPUMetalGLDrawable struct {
	objectivec.Object
}

// IOGPUMetalGLDrawableFromID constructs a [IOGPUMetalGLDrawable] from an objc.ID.
func IOGPUMetalGLDrawableFromID(id objc.ID) IOGPUMetalGLDrawable {
	return IOGPUMetalGLDrawable{objectivec.Object{ID: id}}
}

// Ensure IOGPUMetalGLDrawable implements IIOGPUMetalGLDrawable.
var _ IIOGPUMetalGLDrawable = IOGPUMetalGLDrawable{}

// An interface definition for the [IOGPUMetalGLDrawable] class.
//
// # Methods
//
//   - [IIOGPUMetalGLDrawable.ClearDrawable]
//   - [IIOGPUMetalGLDrawable.Height]
//   - [IIOGPUMetalGLDrawable.LookupIOSurfaceAtIndex]
//   - [IIOGPUMetalGLDrawable.SetDrawableSurfaceModeColorDepthModeFaceLevelVolatileFixedSourceScaleOptionsScaledWidthScaledHeight]
//   - [IIOGPUMetalGLDrawable.SetSwapIntervalLimit]
//   - [IIOGPUMetalGLDrawable.SetSwapRectXYWH]
//   - [IIOGPUMetalGLDrawable.SignalSharedEventValueOperation]
//   - [IIOGPUMetalGLDrawable.SurfaceHeight]
//   - [IIOGPUMetalGLDrawable.SurfaceWidth]
//   - [IIOGPUMetalGLDrawable.Width]
//   - [IIOGPUMetalGLDrawable.WindowMode]
//   - [IIOGPUMetalGLDrawable.InitWithDevice]
type IIOGPUMetalGLDrawable interface {
	objectivec.IObject

	// Topic: Methods

	ClearDrawable() int
	Height() uint32
	LookupIOSurfaceAtIndex(index uint64) objectivec.IObject
	SetDrawableSurfaceModeColorDepthModeFaceLevelVolatileFixedSourceScaleOptionsScaledWidthScaledHeight(surface uint32, mode uint64, mode2 uint32, face uint32, level uint32, volatile uint32, source uint32, options uint32, width uint32, height uint32) int
	SetSwapIntervalLimit(interval int, limit int) int
	SetSwapRectXYWH(x uint32, y uint32, w uint32, h uint32) int
	SignalSharedEventValueOperation(event objectivec.IObject, value uint64, operation uint64) int
	SurfaceHeight() uint32
	SurfaceWidth() uint32
	Width() uint32
	WindowMode() uint64
	InitWithDevice(device objectivec.IObject) IOGPUMetalGLDrawable
}

// Init initializes the instance.
func (i IOGPUMetalGLDrawable) Init() IOGPUMetalGLDrawable {
	rv := objc.SendIfResponds[IOGPUMetalGLDrawable](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalGLDrawable) Autorelease() IOGPUMetalGLDrawable {
	rv := objc.SendIfResponds[IOGPUMetalGLDrawable](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalGLDrawable creates a new IOGPUMetalGLDrawable instance.
func NewIOGPUMetalGLDrawable() IOGPUMetalGLDrawable {
	class := getIOGPUMetalGLDrawableClass()
	rv := objc.SendIfResponds[IOGPUMetalGLDrawable](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMetalGLDrawableWithDevice(device objectivec.IObject) IOGPUMetalGLDrawable {
	instance := getIOGPUMetalGLDrawableClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return IOGPUMetalGLDrawableFromID(rv)
}

func (i IOGPUMetalGLDrawable) ClearDrawable() int {
	rv := objc.SendIfResponds[int](i.ID, objc.Sel("clearDrawable"))
	return rv
}
func (i IOGPUMetalGLDrawable) Height() uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("height"))
	return rv
}
func (i IOGPUMetalGLDrawable) LookupIOSurfaceAtIndex(index uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("lookupIOSurfaceAtIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalGLDrawable) SetDrawableSurfaceModeColorDepthModeFaceLevelVolatileFixedSourceScaleOptionsScaledWidthScaledHeight(surface uint32, mode uint64, mode2 uint32, face uint32, level uint32, volatile uint32, source uint32, options uint32, width uint32, height uint32) int {
	rv := objc.SendIfResponds[int](i.ID, objc.Sel("setDrawableSurface:mode:colorDepthMode:face:level:volatile:fixedSource:scaleOptions:scaledWidth:scaledHeight:"), surface, mode, mode2, face, level, volatile, source, options, width, height)
	return rv
}
func (i IOGPUMetalGLDrawable) SetSwapIntervalLimit(interval int, limit int) int {
	rv := objc.SendIfResponds[int](i.ID, objc.Sel("setSwapInterval:limit:"), interval, limit)
	return rv
}
func (i IOGPUMetalGLDrawable) SetSwapRectXYWH(x uint32, y uint32, w uint32, h uint32) int {
	rv := objc.SendIfResponds[int](i.ID, objc.Sel("setSwapRectX:y:w:h:"), x, y, w, h)
	return rv
}
func (i IOGPUMetalGLDrawable) SignalSharedEventValueOperation(event objectivec.IObject, value uint64, operation uint64) int {
	rv := objc.SendIfResponds[int](i.ID, objc.Sel("signalSharedEvent:value:operation:"), event, value, operation)
	return rv
}
func (i IOGPUMetalGLDrawable) SurfaceHeight() uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("surfaceHeight"))
	return rv
}
func (i IOGPUMetalGLDrawable) SurfaceWidth() uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("surfaceWidth"))
	return rv
}
func (i IOGPUMetalGLDrawable) Width() uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("width"))
	return rv
}
func (i IOGPUMetalGLDrawable) WindowMode() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("windowMode"))
	return rv
}
func (i IOGPUMetalGLDrawable) InitWithDevice(device objectivec.IObject) IOGPUMetalGLDrawable {
	rv := objc.SendIfResponds[IOGPUMetalGLDrawable](i.ID, objc.Sel("initWithDevice:"), device)
	return rv
}
