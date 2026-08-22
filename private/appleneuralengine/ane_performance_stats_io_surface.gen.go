// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEPerformanceStatsIOSurface] class.
var (
	_ANEPerformanceStatsIOSurfaceClass     ANEPerformanceStatsIOSurfaceClass
	_ANEPerformanceStatsIOSurfaceClassOnce sync.Once
)

func getANEPerformanceStatsIOSurfaceClass() ANEPerformanceStatsIOSurfaceClass {
	_ANEPerformanceStatsIOSurfaceClassOnce.Do(func() {
		_ANEPerformanceStatsIOSurfaceClass = ANEPerformanceStatsIOSurfaceClass{class: objc.GetClass("_ANEPerformanceStatsIOSurface")}
	})
	return _ANEPerformanceStatsIOSurfaceClass
}

// GetANEPerformanceStatsIOSurfaceClass returns the class object for _ANEPerformanceStatsIOSurface.
func GetANEPerformanceStatsIOSurfaceClass() ANEPerformanceStatsIOSurfaceClass {
	return getANEPerformanceStatsIOSurfaceClass()
}

type ANEPerformanceStatsIOSurfaceClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEPerformanceStatsIOSurfaceClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEPerformanceStatsIOSurfaceClass) Alloc() ANEPerformanceStatsIOSurface {
	rv := objc.SendIfResponds[ANEPerformanceStatsIOSurface](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ANEPerformanceStatsIOSurface.StatType]
//   - [ANEPerformanceStatsIOSurface.Stats]
//   - [ANEPerformanceStatsIOSurface.InitWithIOSurfaceStatType]
type ANEPerformanceStatsIOSurface struct {
	objectivec.Object
}

// ANEPerformanceStatsIOSurfaceFromID constructs a [ANEPerformanceStatsIOSurface] from an objc.ID.
func ANEPerformanceStatsIOSurfaceFromID(id objc.ID) ANEPerformanceStatsIOSurface {
	return ANEPerformanceStatsIOSurface{objectivec.Object{ID: id}}
}

// Ensure ANEPerformanceStatsIOSurface implements IANEPerformanceStatsIOSurface.
var _ IANEPerformanceStatsIOSurface = ANEPerformanceStatsIOSurface{}

// An interface definition for the [ANEPerformanceStatsIOSurface] class.
//
// # Methods
//
//   - [IANEPerformanceStatsIOSurface.StatType]
//   - [IANEPerformanceStatsIOSurface.Stats]
//   - [IANEPerformanceStatsIOSurface.InitWithIOSurfaceStatType]
type IANEPerformanceStatsIOSurface interface {
	objectivec.IObject

	// Topic: Methods

	StatType() int64
	Stats() IANEIOSurfaceObject
	InitWithIOSurfaceStatType(iOSurface objectivec.IObject, type_ int64) ANEPerformanceStatsIOSurface
}

// Init initializes the instance.
func (a ANEPerformanceStatsIOSurface) Init() ANEPerformanceStatsIOSurface {
	rv := objc.SendIfResponds[ANEPerformanceStatsIOSurface](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEPerformanceStatsIOSurface) Autorelease() ANEPerformanceStatsIOSurface {
	rv := objc.SendIfResponds[ANEPerformanceStatsIOSurface](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEPerformanceStatsIOSurface creates a new ANEPerformanceStatsIOSurface instance.
func NewANEPerformanceStatsIOSurface() ANEPerformanceStatsIOSurface {
	class := getANEPerformanceStatsIOSurfaceClass()
	rv := objc.SendIfResponds[ANEPerformanceStatsIOSurface](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewANEPerformanceStatsIOSurfaceWithIOSurfaceStatType(iOSurface objectivec.IObject, type_ int64) ANEPerformanceStatsIOSurface {
	instance := getANEPerformanceStatsIOSurfaceClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithIOSurface:statType:"), iOSurface, type_)
	return ANEPerformanceStatsIOSurfaceFromID(rv)
}

func (a ANEPerformanceStatsIOSurface) InitWithIOSurfaceStatType(iOSurface objectivec.IObject, type_ int64) ANEPerformanceStatsIOSurface {
	rv := objc.SendIfResponds[ANEPerformanceStatsIOSurface](a.ID, objc.Sel("initWithIOSurface:statType:"), iOSurface, type_)
	return rv
}

func (_ANEPerformanceStatsIOSurfaceClass ANEPerformanceStatsIOSurfaceClass) ObjectWithIOSurfaceStatType(iOSurface objectivec.IObject, type_ int64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEPerformanceStatsIOSurfaceClass.class), objc.Sel("objectWithIOSurface:statType:"), iOSurface, type_)
	return objectivec.Object{ID: rv}
}

func (a ANEPerformanceStatsIOSurface) StatType() int64 {
	rv := objc.SendIfResponds[int64](a.ID, objc.Sel("statType"))
	return rv
}
func (a ANEPerformanceStatsIOSurface) Stats() IANEIOSurfaceObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("stats"))
	return ANEIOSurfaceObjectFromID(objc.ID(rv))
}
