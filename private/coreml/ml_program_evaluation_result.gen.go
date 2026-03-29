// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLProgramEvaluationResult] class.
var (
	_MLProgramEvaluationResultClass     MLProgramEvaluationResultClass
	_MLProgramEvaluationResultClassOnce sync.Once
)

func getMLProgramEvaluationResultClass() MLProgramEvaluationResultClass {
	_MLProgramEvaluationResultClassOnce.Do(func() {
		_MLProgramEvaluationResultClass = MLProgramEvaluationResultClass{class: objc.GetClass("MLProgramEvaluationResult")}
	})
	return _MLProgramEvaluationResultClass
}

// GetMLProgramEvaluationResultClass returns the class object for MLProgramEvaluationResult.
func GetMLProgramEvaluationResultClass() MLProgramEvaluationResultClass {
	return getMLProgramEvaluationResultClass()
}

type MLProgramEvaluationResultClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLProgramEvaluationResultClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLProgramEvaluationResultClass) Alloc() MLProgramEvaluationResult {
	rv := objc.Send[MLProgramEvaluationResult](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLProgramEvaluationResult.EvaluationMetrics]
//   - [MLProgramEvaluationResult.SetEvaluationMetrics]
//   - [MLProgramEvaluationResult.Loss]
//   - [MLProgramEvaluationResult.SetLoss]
// See: https://developer.apple.com/documentation/CoreML/MLProgramEvaluationResult
type MLProgramEvaluationResult struct {
	objectivec.Object
}

// MLProgramEvaluationResultFromID constructs a [MLProgramEvaluationResult] from an objc.ID.
func MLProgramEvaluationResultFromID(id objc.ID) MLProgramEvaluationResult {
	return MLProgramEvaluationResult{objectivec.Object{ID: id}}
}
// Ensure MLProgramEvaluationResult implements IMLProgramEvaluationResult.
var _ IMLProgramEvaluationResult = MLProgramEvaluationResult{}

// An interface definition for the [MLProgramEvaluationResult] class.
//
// # Methods
//
//   - [IMLProgramEvaluationResult.EvaluationMetrics]
//   - [IMLProgramEvaluationResult.SetEvaluationMetrics]
//   - [IMLProgramEvaluationResult.Loss]
//   - [IMLProgramEvaluationResult.SetLoss]
//
// See: https://developer.apple.com/documentation/CoreML/MLProgramEvaluationResult
type IMLProgramEvaluationResult interface {
	objectivec.IObject

	// Topic: Methods

	EvaluationMetrics() objectivec.IObject
	SetEvaluationMetrics(value objectivec.IObject)
	Loss() float64
	SetLoss(value float64)
}

// Init initializes the instance.
func (p MLProgramEvaluationResult) Init() MLProgramEvaluationResult {
	rv := objc.Send[MLProgramEvaluationResult](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLProgramEvaluationResult) Autorelease() MLProgramEvaluationResult {
	rv := objc.Send[MLProgramEvaluationResult](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLProgramEvaluationResult creates a new MLProgramEvaluationResult instance.
func NewMLProgramEvaluationResult() MLProgramEvaluationResult {
	class := getMLProgramEvaluationResultClass()
	rv := objc.Send[MLProgramEvaluationResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramEvaluationResult/evaluationMetrics
func (p MLProgramEvaluationResult) EvaluationMetrics() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("evaluationMetrics"))
	return objectivec.Object{ID: rv}
}
func (p MLProgramEvaluationResult) SetEvaluationMetrics(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setEvaluationMetrics:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLProgramEvaluationResult/loss
func (p MLProgramEvaluationResult) Loss() float64 {
	rv := objc.Send[float64](p.ID, objc.Sel("loss"))
	return rv
}
func (p MLProgramEvaluationResult) SetLoss(value float64) {
	objc.Send[struct{}](p.ID, objc.Sel("setLoss:"), value)
}

