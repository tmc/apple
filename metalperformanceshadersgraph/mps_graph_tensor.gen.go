// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphTensor] class.
var (
	_MPSGraphTensorClass     MPSGraphTensorClass
	_MPSGraphTensorClassOnce sync.Once
)

func getMPSGraphTensorClass() MPSGraphTensorClass {
	_MPSGraphTensorClassOnce.Do(func() {
		_MPSGraphTensorClass = MPSGraphTensorClass{class: objc.GetClass("MPSGraphTensor")}
	})
	return _MPSGraphTensorClass
}

// GetMPSGraphTensorClass returns the class object for MPSGraphTensor.
func GetMPSGraphTensorClass() MPSGraphTensorClass {
	return getMPSGraphTensorClass()
}

type MPSGraphTensorClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MPSGraphTensorClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MPSGraphTensorClass) Alloc() MPSGraphTensor {
	rv := objc.Send[MPSGraphTensor](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// The symbolic representation of a compute data type.
//
// # Overview
//
// [NSCopy] will take a refrence, this is so [NSDictionary] can work with the
// tensor. All tensors are created, owned and destroyed by the MPSGraph
//
// # Instance Properties
//
//   - [MPSGraphTensor.DataType]: The data type of the tensor.
//   - [MPSGraphTensor.Operation]: The operation responsible for creating this tensor.
//   - [MPSGraphTensor.Shape]: The shape of the tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensor
type MPSGraphTensor struct {
	MPSGraphObject
}

// MPSGraphTensorFromID constructs a [MPSGraphTensor] from an objc.ID.
//
// The symbolic representation of a compute data type.
func MPSGraphTensorFromID(id objc.ID) MPSGraphTensor {
	return MPSGraphTensor{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphTensor adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphTensor] class.
//
// # Instance Properties
//
//   - [IMPSGraphTensor.DataType]: The data type of the tensor.
//   - [IMPSGraphTensor.Operation]: The operation responsible for creating this tensor.
//   - [IMPSGraphTensor.Shape]: The shape of the tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensor
type IMPSGraphTensor interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// The data type of the tensor.
	DataType() uint32
	// The operation responsible for creating this tensor.
	Operation() IMPSGraphOperation
	// The shape of the tensor.
	Shape() foundation.NSArray
}

// Init initializes the instance.
func (g MPSGraphTensor) Init() MPSGraphTensor {
	rv := objc.Send[MPSGraphTensor](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphTensor) Autorelease() MPSGraphTensor {
	rv := objc.Send[MPSGraphTensor](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphTensor creates a new MPSGraphTensor instance.
func NewMPSGraphTensor() MPSGraphTensor {
	class := getMPSGraphTensorClass()
	rv := objc.Send[MPSGraphTensor](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The data type of the tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensor/dataType
func (g MPSGraphTensor) DataType() uint32 {
	rv := objc.Send[uint32](g.ID, objc.Sel("dataType"))
	return rv
}

// The operation responsible for creating this tensor.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensor/operation
func (g MPSGraphTensor) Operation() IMPSGraphOperation {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("operation"))
	return MPSGraphOperationFromID(objc.ID(rv))
}

// The shape of the tensor.
//
// # Discussion
//
// Nil shape represents an unranked tensor. -1 value for a dimension
// represents that it will be resolved via shape inference at runtime and it
// can be anything.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphTensor/shape
func (g MPSGraphTensor) Shape() foundation.NSArray {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("shape"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
