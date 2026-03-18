// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLOptimizationHints] class.
var (
	_MLOptimizationHintsClass     MLOptimizationHintsClass
	_MLOptimizationHintsClassOnce sync.Once
)

func getMLOptimizationHintsClass() MLOptimizationHintsClass {
	_MLOptimizationHintsClassOnce.Do(func() {
		_MLOptimizationHintsClass = MLOptimizationHintsClass{class: objc.GetClass("MLOptimizationHints")}
	})
	return _MLOptimizationHintsClass
}

// GetMLOptimizationHintsClass returns the class object for MLOptimizationHints.
func GetMLOptimizationHintsClass() MLOptimizationHintsClass {
	return getMLOptimizationHintsClass()
}

type MLOptimizationHintsClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLOptimizationHintsClass) Alloc() MLOptimizationHints {
	rv := objc.Send[MLOptimizationHints](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// MLOptimizationHints
//
// # Overview
// 
// An object to hold hints that CoreML could use for further optimization
//
// # Getting the reshape frequency
//
//   - [MLOptimizationHints.ReshapeFrequency]: The anticipated reshape frequency
//   - [MLOptimizationHints.SetReshapeFrequency]
//
// # Getting the specialization strategy
//
//   - [MLOptimizationHints.SpecializationStrategy]: Optimization strategy for the model specialization.
//   - [MLOptimizationHints.SetSpecializationStrategy]
//
// See: https://developer.apple.com/documentation/CoreML/MLOptimizationHints-c.class
type MLOptimizationHints struct {
	objectivec.Object
}

// MLOptimizationHintsFromID constructs a [MLOptimizationHints] from an objc.ID.
//
// MLOptimizationHints
func MLOptimizationHintsFromID(id objc.ID) MLOptimizationHints {
	return MLOptimizationHints{objectivec.Object{ID: id}}
}
// NOTE: MLOptimizationHints adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [MLOptimizationHints] class.
//
// # Getting the reshape frequency
//
//   - [IMLOptimizationHints.ReshapeFrequency]: The anticipated reshape frequency
//   - [IMLOptimizationHints.SetReshapeFrequency]
//
// # Getting the specialization strategy
//
//   - [IMLOptimizationHints.SpecializationStrategy]: Optimization strategy for the model specialization.
//   - [IMLOptimizationHints.SetSpecializationStrategy]
//
// See: https://developer.apple.com/documentation/CoreML/MLOptimizationHints-c.class
type IMLOptimizationHints interface {
	objectivec.IObject

	// Topic: Getting the reshape frequency

	// The anticipated reshape frequency
	ReshapeFrequency() MLReshapeFrequencyHint
	SetReshapeFrequency(value MLReshapeFrequencyHint)

	// Topic: Getting the specialization strategy

	// Optimization strategy for the model specialization.
	SpecializationStrategy() MLSpecializationStrategy
	SetSpecializationStrategy(value MLSpecializationStrategy)

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (o MLOptimizationHints) Init() MLOptimizationHints {
	rv := objc.Send[MLOptimizationHints](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o MLOptimizationHints) Autorelease() MLOptimizationHints {
	rv := objc.Send[MLOptimizationHints](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLOptimizationHints creates a new MLOptimizationHints instance.
func NewMLOptimizationHints() MLOptimizationHints {
	class := getMLOptimizationHintsClass()
	rv := objc.Send[MLOptimizationHints](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (o MLOptimizationHints) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](o.ID, objc.Sel("encodeWithCoder:"), coder)
}

// The anticipated reshape frequency
//
// # Discussion
// 
// CoreML framework needs to reshape the model with new shapes for models with
// flexible input. Specify the anticipated reshape frequency (frequent or
// infrequent), so that the framework can optimize for fast shape switching or
// fast prediction on seen shapes.
// 
// The default value is frequent, which means CoreML tries to switch to new
// shapes as fast as possible
//
// See: https://developer.apple.com/documentation/CoreML/MLOptimizationHints-c.class/reshapeFrequency
func (o MLOptimizationHints) ReshapeFrequency() MLReshapeFrequencyHint {
	rv := objc.Send[MLReshapeFrequencyHint](o.ID, objc.Sel("reshapeFrequency"))
	return MLReshapeFrequencyHint(rv)
}
func (o MLOptimizationHints) SetReshapeFrequency(value MLReshapeFrequencyHint) {
	objc.Send[struct{}](o.ID, objc.Sel("setReshapeFrequency:"), value)
}

// Optimization strategy for the model specialization.
//
// # Discussion
// 
// Core ML segments the model’s compute graph and optimizes each segment for
// the target compute device. This process can affect the model loading time
// and the prediction latency.
// 
// Use this option to tailor the specialization strategy for your application.
//
// See: https://developer.apple.com/documentation/CoreML/MLOptimizationHints-c.class/specializationStrategy
func (o MLOptimizationHints) SpecializationStrategy() MLSpecializationStrategy {
	rv := objc.Send[MLSpecializationStrategy](o.ID, objc.Sel("specializationStrategy"))
	return MLSpecializationStrategy(rv)
}
func (o MLOptimizationHints) SetSpecializationStrategy(value MLSpecializationStrategy) {
	objc.Send[struct{}](o.ID, objc.Sel("setSpecializationStrategy:"), value)
}

