// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphType] class.
var (
	_MPSGraphTypeClass     MPSGraphTypeClass
	_MPSGraphTypeClassOnce sync.Once
)

func getMPSGraphTypeClass() MPSGraphTypeClass {
	_MPSGraphTypeClassOnce.Do(func() {
		_MPSGraphTypeClass = MPSGraphTypeClass{class: objc.GetClass("MPSGraphType")}
	})
	return _MPSGraphTypeClass
}

// GetMPSGraphTypeClass returns the class object for MPSGraphType.
func GetMPSGraphTypeClass() MPSGraphTypeClass {
	return getMPSGraphTypeClass()
}

type MPSGraphTypeClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphTypeClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphTypeClass) Alloc() MPSGraphType {
	rv := objc.Send[MPSGraphType](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The base type class for types on tensors.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphType
type MPSGraphType struct {
	MPSGraphObject
}

// MPSGraphTypeFromID constructs a [MPSGraphType] from an objc.ID.
//
// The base type class for types on tensors.
func MPSGraphTypeFromID(id objc.ID) MPSGraphType {
	return MPSGraphType{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphType adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphType] class.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphType
type IMPSGraphType interface {
	IMPSGraphObject
}

// Init initializes the instance.
func (g MPSGraphType) Init() MPSGraphType {
	rv := objc.Send[MPSGraphType](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphType) Autorelease() MPSGraphType {
	rv := objc.Send[MPSGraphType](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphType creates a new MPSGraphType instance.
func NewMPSGraphType() MPSGraphType {
	class := getMPSGraphTypeClass()
	rv := objc.Send[MPSGraphType](objc.ID(class.class), objc.Sel("new"))
	return rv
}
