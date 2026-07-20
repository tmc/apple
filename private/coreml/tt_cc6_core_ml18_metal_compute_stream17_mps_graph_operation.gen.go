// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MPSGraphOperation] class.
var (
	_MPSGraphOperationClass     MPSGraphOperationClass
	_MPSGraphOperationClassOnce sync.Once
)

func getMPSGraphOperationClass() MPSGraphOperationClass {
	_MPSGraphOperationClassOnce.Do(func() {
		_MPSGraphOperationClass = MPSGraphOperationClass{class: objc.GetClass("_TtCC6CoreML18MetalComputeStream17MPSGraphOperation")}
	})
	return _MPSGraphOperationClass
}

// GetMPSGraphOperationClass returns the class object for _TtCC6CoreML18MetalComputeStream17MPSGraphOperation.
func GetMPSGraphOperationClass() MPSGraphOperationClass {
	return getMPSGraphOperationClass()
}

type MPSGraphOperationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphOperationClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphOperationClass) Alloc() MPSGraphOperation {
	rv := objc.Send[MPSGraphOperation](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

type MPSGraphOperation struct {
	objectivec.Object
}

// MPSGraphOperationFromID constructs a [MPSGraphOperation] from an objc.ID.
func MPSGraphOperationFromID(id objc.ID) MPSGraphOperation {
	return MPSGraphOperation{objectivec.Object{ID: id}}
}

// NOTE: MPSGraphOperation struct embeds objectivec.Object (parent type unavailable) but
// IMPSGraphOperation embeds the parent interface; skip compile-time assertion.

// An interface definition for the [MPSGraphOperation] class.
type IMPSGraphOperation interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MPSGraphOperation) Init() MPSGraphOperation {
	rv := objc.Send[MPSGraphOperation](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MPSGraphOperation) Autorelease() MPSGraphOperation {
	rv := objc.Send[MPSGraphOperation](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphOperation creates a new MPSGraphOperation instance.
func NewMPSGraphOperation() MPSGraphOperation {
	class := getMPSGraphOperationClass()
	rv := objc.Send[MPSGraphOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}
