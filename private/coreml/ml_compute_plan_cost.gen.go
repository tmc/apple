// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLComputePlanCost] class.
var (
	_MLComputePlanCostClass     MLComputePlanCostClass
	_MLComputePlanCostClassOnce sync.Once
)

func getMLComputePlanCostClass() MLComputePlanCostClass {
	_MLComputePlanCostClassOnce.Do(func() {
		_MLComputePlanCostClass = MLComputePlanCostClass{class: objc.GetClass("MLComputePlanCost")}
	})
	return _MLComputePlanCostClass
}

// GetMLComputePlanCostClass returns the class object for MLComputePlanCost.
func GetMLComputePlanCostClass() MLComputePlanCostClass {
	return getMLComputePlanCostClass()
}

type MLComputePlanCostClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLComputePlanCostClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLComputePlanCostClass) Alloc() MLComputePlanCost {
	rv := objc.Send[MLComputePlanCost](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLComputePlanCost.InitWithWeight]
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlanCost
type MLComputePlanCost struct {
	objectivec.Object
}

// MLComputePlanCostFromID constructs a [MLComputePlanCost] from an objc.ID.
func MLComputePlanCostFromID(id objc.ID) MLComputePlanCost {
	return MLComputePlanCost{objectivec.Object{ID: id}}
}

// Ensure MLComputePlanCost implements IMLComputePlanCost.
var _ IMLComputePlanCost = MLComputePlanCost{}

// An interface definition for the [MLComputePlanCost] class.
//
// # Methods
//
//   - [IMLComputePlanCost.InitWithWeight]
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlanCost
type IMLComputePlanCost interface {
	objectivec.IObject

	// Topic: Methods

	InitWithWeight(weight float64) MLComputePlanCost
}

// Init initializes the instance.
func (c MLComputePlanCost) Init() MLComputePlanCost {
	rv := objc.Send[MLComputePlanCost](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLComputePlanCost) Autorelease() MLComputePlanCost {
	rv := objc.Send[MLComputePlanCost](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLComputePlanCost creates a new MLComputePlanCost instance.
func NewMLComputePlanCost() MLComputePlanCost {
	class := getMLComputePlanCostClass()
	rv := objc.Send[MLComputePlanCost](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLComputePlanCost/initWithWeight:
func NewComputePlanCostWithWeight(weight float64) MLComputePlanCost {
	instance := getMLComputePlanCostClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWeight:"), weight)
	return MLComputePlanCostFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLComputePlanCost/initWithWeight:
func (c MLComputePlanCost) InitWithWeight(weight float64) MLComputePlanCost {
	rv := objc.Send[MLComputePlanCost](c.ID, objc.Sel("initWithWeight:"), weight)
	return rv
}
