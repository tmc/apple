// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSGraphObject] class.
var (
	_MPSGraphObjectClass     MPSGraphObjectClass
	_MPSGraphObjectClassOnce sync.Once
)

func getMPSGraphObjectClass() MPSGraphObjectClass {
	_MPSGraphObjectClassOnce.Do(func() {
		_MPSGraphObjectClass = MPSGraphObjectClass{class: objc.GetClass("MPSGraphObject")}
	})
	return _MPSGraphObjectClass
}

// GetMPSGraphObjectClass returns the class object for MPSGraphObject.
func GetMPSGraphObjectClass() MPSGraphObjectClass {
	return getMPSGraphObjectClass()
}

type MPSGraphObjectClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphObjectClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphObjectClass) Alloc() MPSGraphObject {
	rv := objc.Send[MPSGraphObject](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The common base class for all Metal Performance Shaders Graph objects.
//
// # Overview
//
// Only the child classes should be used.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphObject
type MPSGraphObject struct {
	objectivec.Object
}

// MPSGraphObjectFromID constructs a [MPSGraphObject] from an objc.ID.
//
// The common base class for all Metal Performance Shaders Graph objects.
func MPSGraphObjectFromID(id objc.ID) MPSGraphObject {
	return MPSGraphObject{objectivec.Object{ID: id}}
}

// NOTE: MPSGraphObject adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphObject] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphObject
type IMPSGraphObject interface {
	objectivec.IObject
}

// Init initializes the instance.
func (g MPSGraphObject) Init() MPSGraphObject {
	rv := objc.Send[MPSGraphObject](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphObject) Autorelease() MPSGraphObject {
	rv := objc.Send[MPSGraphObject](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphObject creates a new MPSGraphObject instance.
func NewMPSGraphObject() MPSGraphObject {
	class := getMPSGraphObjectClass()
	rv := objc.Send[MPSGraphObject](objc.ID(class.class), objc.Sel("new"))
	return rv
}
