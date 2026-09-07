// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

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
	rv := objc.SendIfResponds[MLProgramEvaluationResult](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLProgramEvaluationResult.EvaluationMetrics]
//   - [MLProgramEvaluationResult.SetEvaluationMetrics]
//   - [MLProgramEvaluationResult.Loss]
//   - [MLProgramEvaluationResult.SetLoss]
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
type IMLProgramEvaluationResult interface {
	objectivec.IObject

	// Topic: Methods

	EvaluationMetrics() unsafe.Pointer
	SetEvaluationMetrics(value unsafe.Pointer)
	Loss() float64
	SetLoss(value float64)
}

// Init initializes the instance.
func (m MLProgramEvaluationResult) Init() MLProgramEvaluationResult {
	rv := objc.SendIfResponds[MLProgramEvaluationResult](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLProgramEvaluationResult) Autorelease() MLProgramEvaluationResult {
	rv := objc.SendIfResponds[MLProgramEvaluationResult](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLProgramEvaluationResult creates a new MLProgramEvaluationResult instance.
func NewMLProgramEvaluationResult() MLProgramEvaluationResult {
	class := getMLProgramEvaluationResultClass()
	rv := objc.SendIfResponds[MLProgramEvaluationResult](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (m MLProgramEvaluationResult) EvaluationMetrics() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("evaluationMetrics"))
	return rv
}
func (m MLProgramEvaluationResult) SetEvaluationMetrics(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setEvaluationMetrics:"), value)
}
func (m MLProgramEvaluationResult) Loss() float64 {
	rv := objc.SendIfResponds[float64](m.ID, objc.Sel("loss"))
	return rv
}
func (m MLProgramEvaluationResult) SetLoss(value float64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setLoss:"), value)
}
