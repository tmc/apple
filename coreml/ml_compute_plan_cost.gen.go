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

// Alloc allocates memory for a new instance of the class.
func (mc MLComputePlanCostClass) Alloc() MLComputePlanCost {
	rv := objc.Send[MLComputePlanCost](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class that represents the estimated cost of executing a layer or
// operation.
//
// # Accessing the weight
//
//   - [MLComputePlanCost.Weight]: The estimated workload of executing the operation over the total model execution. The value is between [0.0, 1.0].
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlanCost
type MLComputePlanCost struct {
	objectivec.Object
}

// MLComputePlanCostFromID constructs a [MLComputePlanCost] from an objc.ID.
//
// A class that represents the estimated cost of executing a layer or
// operation.
func MLComputePlanCostFromID(id objc.ID) MLComputePlanCost {
	return MLComputePlanCost{objectivec.Object{ID: id}}
}
// Ensure MLComputePlanCost implements IMLComputePlanCost.
var _ IMLComputePlanCost = MLComputePlanCost{}

// An interface definition for the [MLComputePlanCost] class.
//
// # Accessing the weight
//
//   - [IMLComputePlanCost.Weight]: The estimated workload of executing the operation over the total model execution. The value is between [0.0, 1.0].
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlanCost
type IMLComputePlanCost interface {
	objectivec.IObject

	// Topic: Accessing the weight

	// The estimated workload of executing the operation over the total model execution. The value is between [0.0, 1.0].
	Weight() float64
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

// The estimated workload of executing the operation over the total model
// execution. The value is between [0.0, 1.0].
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlanCost/weight
func (c MLComputePlanCost) Weight() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("weight"))
	return rv
}

