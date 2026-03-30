// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"context"
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLModelStructure] class.
var (
	_MLModelStructureClass     MLModelStructureClass
	_MLModelStructureClassOnce sync.Once
)

func getMLModelStructureClass() MLModelStructureClass {
	_MLModelStructureClassOnce.Do(func() {
		_MLModelStructureClass = MLModelStructureClass{class: objc.GetClass("MLModelStructure")}
	})
	return _MLModelStructureClass
}

// GetMLModelStructureClass returns the class object for MLModelStructure.
func GetMLModelStructureClass() MLModelStructureClass {
	return getMLModelStructureClass()
}

type MLModelStructureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelStructureClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelStructureClass) Alloc() MLModelStructure {
	rv := objc.Send[MLModelStructure](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// A class representing the structure of a model.
//
// # Overview
//
// # Model structures
//
//   - [MLModelStructure.NeuralNetwork]: If the model is of NeuralNetwork type then it is the structure of the NeuralNetwork otherwise `nil`.
//   - [MLModelStructure.Pipeline]: If the model is of Pipeline type then it is the structure of the Pipeline otherwise `nil`.
//   - [MLModelStructure.Program]: If the model is of ML Program type then it is the structure of the ML Program otherwise `nil`.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructure-c.class
type MLModelStructure struct {
	objectivec.Object
}

// MLModelStructureFromID constructs a [MLModelStructure] from an objc.ID.
//
// A class representing the structure of a model.
func MLModelStructureFromID(id objc.ID) MLModelStructure {
	return MLModelStructure{objectivec.Object{ID: id}}
}

// Ensure MLModelStructure implements IMLModelStructure.
var _ IMLModelStructure = MLModelStructure{}

// An interface definition for the [MLModelStructure] class.
//
// # Model structures
//
//   - [IMLModelStructure.NeuralNetwork]: If the model is of NeuralNetwork type then it is the structure of the NeuralNetwork otherwise `nil`.
//   - [IMLModelStructure.Pipeline]: If the model is of Pipeline type then it is the structure of the Pipeline otherwise `nil`.
//   - [IMLModelStructure.Program]: If the model is of ML Program type then it is the structure of the ML Program otherwise `nil`.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructure-c.class
type IMLModelStructure interface {
	objectivec.IObject

	// Topic: Model structures

	// If the model is of NeuralNetwork type then it is the structure of the NeuralNetwork otherwise `nil`.
	NeuralNetwork() IMLModelStructureNeuralNetwork
	// If the model is of Pipeline type then it is the structure of the Pipeline otherwise `nil`.
	Pipeline() IMLModelStructurePipeline
	// If the model is of ML Program type then it is the structure of the ML Program otherwise `nil`.
	Program() IMLModelStructureProgram
}

// Init initializes the instance.
func (m MLModelStructure) Init() MLModelStructure {
	rv := objc.Send[MLModelStructure](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelStructure) Autorelease() MLModelStructure {
	rv := objc.Send[MLModelStructure](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelStructure creates a new MLModelStructure instance.
func NewMLModelStructure() MLModelStructure {
	class := getMLModelStructureClass()
	rv := objc.Send[MLModelStructure](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Construct the model structure asynchronously given the location of its
// on-disk representation.
//
// url: The location of its on-disk representation (.mlmodelc directory).
//
// handler: When the model structure is constructed successfully or unsuccessfully, the
// completion handler is invoked with a valid MLModelStructure instance or
// NSError object.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructure-c.class/loadContentsOfURL:completionHandler:
func (_MLModelStructureClass MLModelStructureClass) LoadContentsOfURLCompletionHandler(url foundation.INSURL, handler MLModelStructureErrorHandler) {
	_block1, _ := NewMLModelStructureErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_MLModelStructureClass.class), objc.Sel("loadContentsOfURL:completionHandler:"), url, _block1)
}

// Construct the model structure asynchronously given the model asset.
//
// asset: The model asset.
//
// handler: When the model structure is constructed successfully or unsuccessfully, the
// completion handler is invoked with a valid MLModelStructure instance or
// NSError object.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructure-c.class/loadModelAsset:completionHandler:
func (_MLModelStructureClass MLModelStructureClass) LoadModelAssetCompletionHandler(asset IMLModelAsset, handler MLModelStructureErrorHandler) {
	_block1, _ := NewMLModelStructureErrorBlock(handler)
	objc.Send[objc.ID](objc.ID(_MLModelStructureClass.class), objc.Sel("loadModelAsset:completionHandler:"), asset, _block1)
}

// If the model is of NeuralNetwork type then it is the structure of the
// NeuralNetwork otherwise `nil`.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructure-c.class/neuralNetwork
func (m MLModelStructure) NeuralNetwork() IMLModelStructureNeuralNetwork {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("neuralNetwork"))
	return MLModelStructureNeuralNetworkFromID(objc.ID(rv))
}

// If the model is of Pipeline type then it is the structure of the Pipeline
// otherwise `nil`.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructure-c.class/pipeline
func (m MLModelStructure) Pipeline() IMLModelStructurePipeline {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("pipeline"))
	return MLModelStructurePipelineFromID(objc.ID(rv))
}

// If the model is of ML Program type then it is the structure of the ML
// Program otherwise `nil`.
//
// See: https://developer.apple.com/documentation/CoreML/MLModelStructure-c.class/program
func (m MLModelStructure) Program() IMLModelStructureProgram {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("program"))
	return MLModelStructureProgramFromID(objc.ID(rv))
}

// LoadContentsOfURL is a synchronous wrapper around [MLModelStructure.LoadContentsOfURLCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (mc MLModelStructureClass) LoadContentsOfURL(ctx context.Context, url foundation.INSURL) (*MLModelStructure, error) {
	type result struct {
		val *MLModelStructure
		err error
	}
	done := make(chan result, 1)
	mc.LoadContentsOfURLCompletionHandler(url, func(val *MLModelStructure, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// LoadModelAsset is a synchronous wrapper around [MLModelStructure.LoadModelAssetCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (mc MLModelStructureClass) LoadModelAsset(ctx context.Context, asset IMLModelAsset) (*MLModelStructure, error) {
	type result struct {
		val *MLModelStructure
		err error
	}
	done := make(chan result, 1)
	mc.LoadModelAssetCompletionHandler(asset, func(val *MLModelStructure, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}
