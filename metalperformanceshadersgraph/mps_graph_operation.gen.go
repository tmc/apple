// Code generated from Apple documentation for MetalPerformanceShadersGraph. DO NOT EDIT.

package metalperformanceshadersgraph

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// The class instance for the [MPSGraphOperation] class.
var (
	_MPSGraphOperationClass     MPSGraphOperationClass
	_MPSGraphOperationClassOnce sync.Once
)

func getMPSGraphOperationClass() MPSGraphOperationClass {
	_MPSGraphOperationClassOnce.Do(func() {
		_MPSGraphOperationClass = MPSGraphOperationClass{class: objc.GetClass("MPSGraphOperation")}
	})
	return _MPSGraphOperationClass
}

// GetMPSGraphOperationClass returns the class object for MPSGraphOperation.
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

// A symbolic representation of a compute operation.
//
// # Overview
//
// [NSCopy] will take a refrence, this is so [NSDictionary] can work with the
// tensor. All operations are created, owned and destroyed by the graph.
//
// # Instance Properties
//
//   - [MPSGraphOperation.ControlDependencies]: The set of operations guaranteed to execute before this operation.
//   - [MPSGraphOperation.Graph]: The graph on which the operation is defined.
//   - [MPSGraphOperation.InputTensors]: The input tensors of the operation.
//   - [MPSGraphOperation.Name]: Name of the operation.
//   - [MPSGraphOperation.OutputTensors]: The output tensors of the operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOperation
type MPSGraphOperation struct {
	MPSGraphObject
}

// MPSGraphOperationFromID constructs a [MPSGraphOperation] from an objc.ID.
//
// A symbolic representation of a compute operation.
func MPSGraphOperationFromID(id objc.ID) MPSGraphOperation {
	return MPSGraphOperation{MPSGraphObject: MPSGraphObjectFromID(id)}
}

// NOTE: MPSGraphOperation adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MPSGraphOperation] class.
//
// # Instance Properties
//
//   - [IMPSGraphOperation.ControlDependencies]: The set of operations guaranteed to execute before this operation.
//   - [IMPSGraphOperation.Graph]: The graph on which the operation is defined.
//   - [IMPSGraphOperation.InputTensors]: The input tensors of the operation.
//   - [IMPSGraphOperation.Name]: Name of the operation.
//   - [IMPSGraphOperation.OutputTensors]: The output tensors of the operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOperation
type IMPSGraphOperation interface {
	IMPSGraphObject

	// Topic: Instance Properties

	// The set of operations guaranteed to execute before this operation.
	ControlDependencies() []MPSGraphOperation
	// The graph on which the operation is defined.
	Graph() IMPSGraph
	// The input tensors of the operation.
	InputTensors() []MPSGraphTensor
	// Name of the operation.
	Name() string
	// The output tensors of the operation.
	OutputTensors() []MPSGraphTensor
}

// Init initializes the instance.
func (g MPSGraphOperation) Init() MPSGraphOperation {
	rv := objc.Send[MPSGraphOperation](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g MPSGraphOperation) Autorelease() MPSGraphOperation {
	rv := objc.Send[MPSGraphOperation](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewMPSGraphOperation creates a new MPSGraphOperation instance.
func NewMPSGraphOperation() MPSGraphOperation {
	class := getMPSGraphOperationClass()
	rv := objc.Send[MPSGraphOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The set of operations guaranteed to execute before this operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOperation/controlDependencies
func (g MPSGraphOperation) ControlDependencies() []MPSGraphOperation {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("controlDependencies"))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphOperation {
		return MPSGraphOperationFromID(id)
	})
}

// The graph on which the operation is defined.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOperation/graph
func (g MPSGraphOperation) Graph() IMPSGraph {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("graph"))
	return MPSGraphFromID(objc.ID(rv))
}

// The input tensors of the operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOperation/inputTensors
func (g MPSGraphOperation) InputTensors() []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("inputTensors"))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}

// Name of the operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOperation/name
func (g MPSGraphOperation) Name() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}

// The output tensors of the operation.
//
// See: https://developer.apple.com/documentation/MetalPerformanceShadersGraph/MPSGraphOperation/outputTensors
func (g MPSGraphOperation) OutputTensors() []MPSGraphTensor {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("outputTensors"))
	return objc.ConvertSlice(rv, func(id objc.ID) MPSGraphTensor {
		return MPSGraphTensorFromID(id)
	})
}
