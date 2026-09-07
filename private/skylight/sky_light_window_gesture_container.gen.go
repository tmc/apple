// Code generated from Apple documentation for skylight. DO NOT EDIT.

package skylight

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SkyLightWindowGestureContainer] class.
var (
	_SkyLightWindowGestureContainerClass     SkyLightWindowGestureContainerClass
	_SkyLightWindowGestureContainerClassOnce sync.Once
)

func getSkyLightWindowGestureContainerClass() SkyLightWindowGestureContainerClass {
	_SkyLightWindowGestureContainerClassOnce.Do(func() {
		_SkyLightWindowGestureContainerClass = SkyLightWindowGestureContainerClass{class: objc.GetClass("SkyLight.WindowGestureContainer")}
	})
	return _SkyLightWindowGestureContainerClass
}

// GetSkyLightWindowGestureContainerClass returns the class object for SkyLight.WindowGestureContainer.
func GetSkyLightWindowGestureContainerClass() SkyLightWindowGestureContainerClass {
	return getSkyLightWindowGestureContainerClass()
}

type SkyLightWindowGestureContainerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SkyLightWindowGestureContainerClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SkyLightWindowGestureContainerClass) Alloc() SkyLightWindowGestureContainer {
	rv := objc.SendIfResponds[SkyLightWindowGestureContainer](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [SkyLightWindowGestureContainer.CreateDebugDescription]
//   - [SkyLightWindowGestureContainer.Invalidate]
//   - [SkyLightWindowGestureContainer.InitWithInternalCGXWindow]
type SkyLightWindowGestureContainer struct {
	objectivec.Object
}

// SkyLightWindowGestureContainerFromID constructs a [SkyLightWindowGestureContainer] from an objc.ID.
func SkyLightWindowGestureContainerFromID(id objc.ID) SkyLightWindowGestureContainer {
	return SkyLightWindowGestureContainer{objectivec.Object{ID: id}}
}

// Ensure SkyLightWindowGestureContainer implements ISkyLightWindowGestureContainer.
var _ ISkyLightWindowGestureContainer = SkyLightWindowGestureContainer{}

// An interface definition for the [SkyLightWindowGestureContainer] class.
//
// # Methods
//
//   - [ISkyLightWindowGestureContainer.CreateDebugDescription]
//   - [ISkyLightWindowGestureContainer.Invalidate]
//   - [ISkyLightWindowGestureContainer.InitWithInternalCGXWindow]
type ISkyLightWindowGestureContainer interface {
	objectivec.IObject

	// Topic: Methods

	CreateDebugDescription() objectivec.IObject
	Invalidate()
	InitWithInternalCGXWindow(cGXWindow unsafe.Pointer) SkyLightWindowGestureContainer
}

// Init initializes the instance.
func (s SkyLightWindowGestureContainer) Init() SkyLightWindowGestureContainer {
	rv := objc.SendIfResponds[SkyLightWindowGestureContainer](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SkyLightWindowGestureContainer) Autorelease() SkyLightWindowGestureContainer {
	rv := objc.SendIfResponds[SkyLightWindowGestureContainer](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSkyLightWindowGestureContainer creates a new SkyLightWindowGestureContainer instance.
func NewSkyLightWindowGestureContainer() SkyLightWindowGestureContainer {
	class := getSkyLightWindowGestureContainerClass()
	rv := objc.SendIfResponds[SkyLightWindowGestureContainer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewSkyLightWindowGestureContainerWithInternalCGXWindow(cGXWindow unsafe.Pointer) SkyLightWindowGestureContainer {
	instance := getSkyLightWindowGestureContainerClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithInternalCGXWindow:"), cGXWindow)
	return SkyLightWindowGestureContainerFromID(rv)
}

func (s SkyLightWindowGestureContainer) CreateDebugDescription() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](s.ID, objc.Sel("createDebugDescription"))
	return objectivec.Object{ID: rv}
}
func (s SkyLightWindowGestureContainer) Invalidate() {
	objc.SendIfResponds[objc.ID](s.ID, objc.Sel("invalidate"))
}
func (s SkyLightWindowGestureContainer) InitWithInternalCGXWindow(cGXWindow unsafe.Pointer) SkyLightWindowGestureContainer {
	rv := objc.SendIfResponds[SkyLightWindowGestureContainer](s.ID, objc.Sel("initWithInternalCGXWindow:"), cGXWindow)
	return rv
}
