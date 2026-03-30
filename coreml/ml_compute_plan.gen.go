// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLComputePlan] class.
var (
	_MLComputePlanClass     MLComputePlanClass
	_MLComputePlanClassOnce sync.Once
)

func getMLComputePlanClass() MLComputePlanClass {
	_MLComputePlanClassOnce.Do(func() {
		_MLComputePlanClass = MLComputePlanClass{class: objc.GetClass("MLComputePlan")}
	})
	return _MLComputePlanClass
}

// GetMLComputePlanClass returns the class object for MLComputePlan.
func GetMLComputePlanClass() MLComputePlanClass {
	return getMLComputePlanClass()
}

type MLComputePlanClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLComputePlanClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLComputePlanClass) Alloc() MLComputePlan {
	rv := objc.Send[MLComputePlan](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class describing the plan for executing a model.
//
// # Overview
//
// The application can use the plan to estimate the necessary cost and
// resources of the model before running the predictions.
//
// # Getting the model structure
//
//   - [MLComputePlan.ModelStructure]: The model structure.
//
// # Getting the device usage
//
//   - [MLComputePlan.ComputeDeviceUsageForMLProgramOperation]: Returns The anticipated compute devices that would be used for executing an ML Program operation.
//   - [MLComputePlan.ComputeDeviceUsageForNeuralNetworkLayer]: Returns the anticipated compute devices that would be used for executing a NeuralNetwork layer.
//
// # Getting the estimated cost
//
//   - [MLComputePlan.EstimatedCostOfMLProgramOperation]: Returns the estimated cost of executing an ML Program operation.
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw
type MLComputePlan struct {
	objectivec.Object
}

// MLComputePlanFromID constructs a [MLComputePlan] from an objc.ID.
//
// A class describing the plan for executing a model.
func MLComputePlanFromID(id objc.ID) MLComputePlan {
	return MLComputePlan{objectivec.Object{ID: id}}
}

// Ensure MLComputePlan implements IMLComputePlan.
var _ IMLComputePlan = MLComputePlan{}

// An interface definition for the [MLComputePlan] class.
//
// # Getting the model structure
//
//   - [IMLComputePlan.ModelStructure]: The model structure.
//
// # Getting the device usage
//
//   - [IMLComputePlan.ComputeDeviceUsageForMLProgramOperation]: Returns The anticipated compute devices that would be used for executing an ML Program operation.
//   - [IMLComputePlan.ComputeDeviceUsageForNeuralNetworkLayer]: Returns the anticipated compute devices that would be used for executing a NeuralNetwork layer.
//
// # Getting the estimated cost
//
//   - [IMLComputePlan.EstimatedCostOfMLProgramOperation]: Returns the estimated cost of executing an ML Program operation.
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw
type IMLComputePlan interface {
	objectivec.IObject

	// Topic: Getting the model structure

	// The model structure.
	ModelStructure() IMLModelStructure

	// Topic: Getting the device usage

	// Returns The anticipated compute devices that would be used for executing an ML Program operation.
	ComputeDeviceUsageForMLProgramOperation(operation IMLModelStructureProgramOperation) IMLComputePlanDeviceUsage
	// Returns the anticipated compute devices that would be used for executing a NeuralNetwork layer.
	ComputeDeviceUsageForNeuralNetworkLayer(layer IMLModelStructureNeuralNetworkLayer) IMLComputePlanDeviceUsage

	// Topic: Getting the estimated cost

	// Returns the estimated cost of executing an ML Program operation.
	EstimatedCostOfMLProgramOperation(operation IMLModelStructureProgramOperation) IMLComputePlanCost
}

// Init initializes the instance.
func (c MLComputePlan) Init() MLComputePlan {
	rv := objc.Send[MLComputePlan](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c MLComputePlan) Autorelease() MLComputePlan {
	rv := objc.Send[MLComputePlan](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLComputePlan creates a new MLComputePlan instance.
func NewMLComputePlan() MLComputePlan {
	class := getMLComputePlanClass()
	rv := objc.Send[MLComputePlan](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns The anticipated compute devices that would be used for executing an
// ML Program operation.
//
// operation: An ML Program operation.
//
// # Return Value
//
// The anticipated compute devices that would be used for executing the
// operation or `nil`if the usage couldn’t be determined.
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw/computeDeviceUsageForMLProgramOperation:
func (c MLComputePlan) ComputeDeviceUsageForMLProgramOperation(operation IMLModelStructureProgramOperation) IMLComputePlanDeviceUsage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("computeDeviceUsageForMLProgramOperation:"), operation)
	return MLComputePlanDeviceUsageFromID(rv)
}

// Returns the anticipated compute devices that would be used for executing a
// NeuralNetwork layer.
//
// layer: A NeuralNetwork layer.
//
// # Return Value
//
// The anticipated compute devices that would be used for executing the layer
// or `nil` if the usage couldn’t be determined.
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw/computeDeviceUsageForNeuralNetworkLayer:
func (c MLComputePlan) ComputeDeviceUsageForNeuralNetworkLayer(layer IMLModelStructureNeuralNetworkLayer) IMLComputePlanDeviceUsage {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("computeDeviceUsageForNeuralNetworkLayer:"), layer)
	return MLComputePlanDeviceUsageFromID(rv)
}

// Returns the estimated cost of executing an ML Program operation.
//
// operation: An ML Program operation.
//
// # Return Value
//
// The estimated cost of executing the operation or nil if the cost couldn’t
// be estimated.
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw/estimatedCostOfMLProgramOperation:
func (c MLComputePlan) EstimatedCostOfMLProgramOperation(operation IMLModelStructureProgramOperation) IMLComputePlanCost {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("estimatedCostOfMLProgramOperation:"), operation)
	return MLComputePlanCostFromID(rv)
}

// Construct the compute plan of a model asynchronously given the location of
// its on-disk representation.
//
// url: The location of its on-disk representation (.mlmodelc directory).
//
// configuration: The model configuration.
//
// handler: When the compute plan is constructed successfully or unsuccessfully, the
// completion handler is invoked with a valid MLComputePlan instance or
// NSError object.
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw/loadContentsOfURL:configuration:completionHandler:
func (_MLComputePlanClass MLComputePlanClass) LoadContentsOfURLConfigurationCompletionHandler(url foundation.INSURL, configuration IMLModelConfiguration, handler MLComputePlanErrorHandler) {
	_block2, _ := NewMLComputePlanErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_MLComputePlanClass.class), objc.Sel("loadContentsOfURL:configuration:completionHandler:"), url, configuration, _block2)
}

// Construct the compute plan of a model asynchronously given the model asset.
//
// asset: The model asset.
//
// configuration: The model configuration.
//
// handler: When the compute plan is constructed successfully or unsuccessfully, the
// completion handler is invoked with a valid MLComputePlan instance or
// NSError object.
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw/loadModelAsset:configuration:completionHandler:
func (_MLComputePlanClass MLComputePlanClass) LoadModelAssetConfigurationCompletionHandler(asset IMLModelAsset, configuration IMLModelConfiguration, handler MLComputePlanErrorHandler) {
	_block2, _ := NewMLComputePlanErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_MLComputePlanClass.class), objc.Sel("loadModelAsset:configuration:completionHandler:"), asset, configuration, _block2)
}

// The model structure.
//
// See: https://developer.apple.com/documentation/CoreML/MLComputePlan-85vdw/modelStructure
func (c MLComputePlan) ModelStructure() IMLModelStructure {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("modelStructure"))
	return MLModelStructureFromID(objc.ID(rv))
}

// LoadContentsOfURLConfiguration is a synchronous wrapper around [MLComputePlan.LoadContentsOfURLConfigurationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (cc MLComputePlanClass) LoadContentsOfURLConfiguration(ctx context.Context, url foundation.INSURL, configuration IMLModelConfiguration) (*MLComputePlan, error) {
	type result struct {
		val *MLComputePlan
		err error
	}
	done := make(chan result, 1)
	cc.LoadContentsOfURLConfigurationCompletionHandler(url, configuration, func(val *MLComputePlan, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// LoadModelAssetConfiguration is a synchronous wrapper around [MLComputePlan.LoadModelAssetConfigurationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (cc MLComputePlanClass) LoadModelAssetConfiguration(ctx context.Context, asset IMLModelAsset, configuration IMLModelConfiguration) (*MLComputePlan, error) {
	type result struct {
		val *MLComputePlan
		err error
	}
	done := make(chan result, 1)
	cc.LoadModelAssetConfigurationCompletionHandler(asset, configuration, func(val *MLComputePlan, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
