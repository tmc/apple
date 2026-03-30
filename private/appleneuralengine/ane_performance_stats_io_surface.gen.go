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
	rv := objc.Send[ANEPerformanceStatsIOSurface](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ANEPerformanceStatsIOSurface.StatType]
//   - [ANEPerformanceStatsIOSurface.Stats]
//   - [ANEPerformanceStatsIOSurface.InitWithIOSurfaceStatType]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStatsIOSurface
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
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStatsIOSurface
type IANEPerformanceStatsIOSurface interface {
	objectivec.IObject

	// Topic: Methods

	StatType() int64
	Stats() *ANEIOSurfaceObject
	InitWithIOSurfaceStatType(iOSurface objectivec.IObject, type_ int64) ANEPerformanceStatsIOSurface
}

// Init initializes the instance.
func (a ANEPerformanceStatsIOSurface) Init() ANEPerformanceStatsIOSurface {
	rv := objc.Send[ANEPerformanceStatsIOSurface](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEPerformanceStatsIOSurface) Autorelease() ANEPerformanceStatsIOSurface {
	rv := objc.Send[ANEPerformanceStatsIOSurface](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEPerformanceStatsIOSurface creates a new ANEPerformanceStatsIOSurface instance.
func NewANEPerformanceStatsIOSurface() ANEPerformanceStatsIOSurface {
	class := getANEPerformanceStatsIOSurfaceClass()
	rv := objc.Send[ANEPerformanceStatsIOSurface](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStatsIOSurface/initWithIOSurface:statType:
func NewANEPerformanceStatsIOSurfaceWithIOSurfaceStatType(iOSurface objectivec.IObject, type_ int64) ANEPerformanceStatsIOSurface {
	instance := getANEPerformanceStatsIOSurfaceClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIOSurface:statType:"), iOSurface, type_)
	return ANEPerformanceStatsIOSurfaceFromID(rv)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStatsIOSurface/initWithIOSurface:statType:
func (a ANEPerformanceStatsIOSurface) InitWithIOSurfaceStatType(iOSurface objectivec.IObject, type_ int64) ANEPerformanceStatsIOSurface {
	rv := objc.Send[ANEPerformanceStatsIOSurface](a.ID, objc.Sel("initWithIOSurface:statType:"), iOSurface, type_)
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStatsIOSurface/objectWithIOSurface:statType:
func (_ANEPerformanceStatsIOSurfaceClass ANEPerformanceStatsIOSurfaceClass) ObjectWithIOSurfaceStatType(iOSurface objectivec.IObject, type_ int64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEPerformanceStatsIOSurfaceClass.class), objc.Sel("objectWithIOSurface:statType:"), iOSurface, type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStatsIOSurface/statType
func (a ANEPerformanceStatsIOSurface) StatType() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("statType"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEPerformanceStatsIOSurface/stats
func (a ANEPerformanceStatsIOSurface) Stats() *ANEIOSurfaceObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("stats"))
	if rv == 0 {
		return nil
	}
	val := ANEIOSurfaceObjectFromID(objc.ID(rv))
	return &val
}
