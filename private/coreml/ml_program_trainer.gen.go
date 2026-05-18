// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLProgramTrainer] class.
var (
	_MLProgramTrainerClass     MLProgramTrainerClass
	_MLProgramTrainerClassOnce sync.Once
)

func getMLProgramTrainerClass() MLProgramTrainerClass {
	_MLProgramTrainerClassOnce.Do(func() {
		_MLProgramTrainerClass = MLProgramTrainerClass{class: objc.GetClass("MLProgramTrainer")}
	})
	return _MLProgramTrainerClass
}

// GetMLProgramTrainerClass returns the class object for MLProgramTrainer.
func GetMLProgramTrainerClass() MLProgramTrainerClass {
	return getMLProgramTrainerClass()
}

type MLProgramTrainerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLProgramTrainerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLProgramTrainerClass) Alloc() MLProgramTrainer {
	rv := objc.Send[MLProgramTrainer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLProgramTrainer.AttachLearningRateToFeatures]
//   - [MLProgramTrainer.Context]
//   - [MLProgramTrainer.SetContext]
//   - [MLProgramTrainer.CopyCurrentTrainingDelta]
//   - [MLProgramTrainer.CurrentUpdatedWeights]
//   - [MLProgramTrainer.SetCurrentUpdatedWeights]
//   - [MLProgramTrainer.EvaluateUsingTestDataError]
//   - [MLProgramTrainer.EvaluateUsingTestDataEvaluationMetricNamesError]
//   - [MLProgramTrainer.EvaluateUsingTestDataEvaluationMetricNamesEvaluateOnTrainFunctionError]
//   - [MLProgramTrainer.Evaluator]
//   - [MLProgramTrainer.SetEvaluator]
//   - [MLProgramTrainer.FlattenFeaturesOrderedFeatures]
//   - [MLProgramTrainer.InferenceModel]
//   - [MLProgramTrainer.LearningRate]
//   - [MLProgramTrainer.SetLearningRate]
//   - [MLProgramTrainer.OrderedTrainableWeightsNames]
//   - [MLProgramTrainer.Program]
//   - [MLProgramTrainer.SetProgram]
//   - [MLProgramTrainer.TrainUsingTrainingDataError]
//   - [MLProgramTrainer.TrainUsingTrainingDataEvaluationMetricNamesError]
//   - [MLProgramTrainer.InitWithProgramLearningRateError]
//
// See: https://developer.apple.com/documentation/CoreML/MLProgramTrainer
type MLProgramTrainer struct {
	objectivec.Object
}

// MLProgramTrainerFromID constructs a [MLProgramTrainer] from an objc.ID.
func MLProgramTrainerFromID(id objc.ID) MLProgramTrainer {
	return MLProgramTrainer{objectivec.Object{ID: id}}
}

// Ensure MLProgramTrainer implements IMLProgramTrainer.
var _ IMLProgramTrainer = MLProgramTrainer{}

// An interface definition for the [MLProgramTrainer] class.
//
// # Methods
//
//   - [IMLProgramTrainer.AttachLearningRateToFeatures]
//   - [IMLProgramTrainer.Context]
//   - [IMLProgramTrainer.SetContext]
//   - [IMLProgramTrainer.CopyCurrentTrainingDelta]
//   - [IMLProgramTrainer.CurrentUpdatedWeights]
//   - [IMLProgramTrainer.SetCurrentUpdatedWeights]
//   - [IMLProgramTrainer.EvaluateUsingTestDataError]
//   - [IMLProgramTrainer.EvaluateUsingTestDataEvaluationMetricNamesError]
//   - [IMLProgramTrainer.EvaluateUsingTestDataEvaluationMetricNamesEvaluateOnTrainFunctionError]
//   - [IMLProgramTrainer.Evaluator]
//   - [IMLProgramTrainer.SetEvaluator]
//   - [IMLProgramTrainer.FlattenFeaturesOrderedFeatures]
//   - [IMLProgramTrainer.InferenceModel]
//   - [IMLProgramTrainer.LearningRate]
//   - [IMLProgramTrainer.SetLearningRate]
//   - [IMLProgramTrainer.OrderedTrainableWeightsNames]
//   - [IMLProgramTrainer.Program]
//   - [IMLProgramTrainer.SetProgram]
//   - [IMLProgramTrainer.TrainUsingTrainingDataError]
//   - [IMLProgramTrainer.TrainUsingTrainingDataEvaluationMetricNamesError]
//   - [IMLProgramTrainer.InitWithProgramLearningRateError]
//
// See: https://developer.apple.com/documentation/CoreML/MLProgramTrainer
type IMLProgramTrainer interface {
	objectivec.IObject

	// Topic: Methods

	AttachLearningRateToFeatures(features objectivec.IObject) objectivec.IObject
	Context() IMLProgramContext
	SetContext(value IMLProgramContext)
	CopyCurrentTrainingDelta() objectivec.IObject
	CurrentUpdatedWeights() objectivec.IObject
	SetCurrentUpdatedWeights(value objectivec.IObject)
	EvaluateUsingTestDataError(data objectivec.IObject) (objectivec.IObject, error)
	EvaluateUsingTestDataEvaluationMetricNamesError(data objectivec.IObject, names objectivec.IObject) (objectivec.IObject, error)
	EvaluateUsingTestDataEvaluationMetricNamesEvaluateOnTrainFunctionError(data objectivec.IObject, names objectivec.IObject, function bool) (objectivec.IObject, error)
	Evaluator() IMLProgramEvaluator
	SetEvaluator(value IMLProgramEvaluator)
	FlattenFeaturesOrderedFeatures(features objectivec.IObject, features2 objectivec.IObject) objectivec.IObject
	InferenceModel() objectivec.IObject
	LearningRate() float64
	SetLearningRate(value float64)
	OrderedTrainableWeightsNames() objectivec.IObject
	Program() objectivec.IObject
	SetProgram(value objectivec.IObject)
	TrainUsingTrainingDataError(data objectivec.IObject) (objectivec.IObject, error)
	TrainUsingTrainingDataEvaluationMetricNamesError(data objectivec.IObject, names objectivec.IObject) (objectivec.IObject, error)
	InitWithProgramLearningRateError(program objectivec.IObject, rate float64) (MLProgramTrainer, error)
}

// Init initializes the instance.
func (m MLProgramTrainer) Init() MLProgramTrainer {
	rv := objc.Send[MLProgramTrainer](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLProgramTrainer) Autorelease() MLProgramTrainer {
	rv := objc.Send[MLProgramTrainer](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLProgramTrainer creates a new MLProgramTrainer instance.
func NewMLProgramTrainer() MLProgramTrainer {
	class := getMLProgramTrainerClass()
	rv := objc.Send[MLProgramTrainer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramTrainer/initWithProgram:learningRate:error:
func NewProgramTrainerWithProgramLearningRateError(program objectivec.IObject, rate float64) (MLProgramTrainer, error) {
	var errorPtr objc.ID
	instance := getMLProgramTrainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithProgram:learningRate:error:"), program, rate, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLProgramTrainer{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLProgramTrainerFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramTrainer/attachLearningRateToFeatures:
func (m MLProgramTrainer) AttachLearningRateToFeatures(features objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("attachLearningRateToFeatures:"), features)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramTrainer/copyCurrentTrainingDelta
func (m MLProgramTrainer) CopyCurrentTrainingDelta() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("copyCurrentTrainingDelta"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramTrainer/evaluateUsingTestData:error:
func (m MLProgramTrainer) EvaluateUsingTestDataError(data objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("evaluateUsingTestData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLProgramTrainer/evaluateUsingTestData:evaluationMetricNames:error:
func (m MLProgramTrainer) EvaluateUsingTestDataEvaluationMetricNamesError(data objectivec.IObject, names objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("evaluateUsingTestData:evaluationMetricNames:error:"), data, names, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLProgramTrainer/evaluateUsingTestData:evaluationMetricNames:evaluateOnTrainFunction:error:
func (m MLProgramTrainer) EvaluateUsingTestDataEvaluationMetricNamesEvaluateOnTrainFunctionError(data objectivec.IObject, names objectivec.IObject, function bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("evaluateUsingTestData:evaluationMetricNames:evaluateOnTrainFunction:error:"), data, names, function, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLProgramTrainer/flattenFeatures:orderedFeatures:
func (m MLProgramTrainer) FlattenFeaturesOrderedFeatures(features objectivec.IObject, features2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("flattenFeatures:orderedFeatures:"), features, features2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramTrainer/orderedTrainableWeightsNames
func (m MLProgramTrainer) OrderedTrainableWeightsNames() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("orderedTrainableWeightsNames"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramTrainer/trainUsingTrainingData:error:
func (m MLProgramTrainer) TrainUsingTrainingDataError(data objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("trainUsingTrainingData:error:"), data, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLProgramTrainer/trainUsingTrainingData:evaluationMetricNames:error:
func (m MLProgramTrainer) TrainUsingTrainingDataEvaluationMetricNamesError(data objectivec.IObject, names objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("trainUsingTrainingData:evaluationMetricNames:error:"), data, names, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLProgramTrainer/initWithProgram:learningRate:error:
func (m MLProgramTrainer) InitWithProgramLearningRateError(program objectivec.IObject, rate float64) (MLProgramTrainer, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithProgram:learningRate:error:"), program, rate, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLProgramTrainer{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLProgramTrainerFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLProgramTrainer/context
func (m MLProgramTrainer) Context() IMLProgramContext {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("context"))
	return MLProgramContextFromID(objc.ID(rv))
}
func (m MLProgramTrainer) SetContext(value IMLProgramContext) {
	objc.Send[struct{}](m.ID, objc.Sel("setContext:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramTrainer/currentUpdatedWeights
func (m MLProgramTrainer) CurrentUpdatedWeights() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("currentUpdatedWeights"))
	return objectivec.Object{ID: rv}
}
func (m MLProgramTrainer) SetCurrentUpdatedWeights(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setCurrentUpdatedWeights:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramTrainer/evaluator
func (m MLProgramTrainer) Evaluator() IMLProgramEvaluator {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("evaluator"))
	return MLProgramEvaluatorFromID(objc.ID(rv))
}
func (m MLProgramTrainer) SetEvaluator(value IMLProgramEvaluator) {
	objc.Send[struct{}](m.ID, objc.Sel("setEvaluator:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramTrainer/inferenceModel
func (m MLProgramTrainer) InferenceModel() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inferenceModel"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramTrainer/learningRate
func (m MLProgramTrainer) LearningRate() float64 {
	rv := objc.Send[float64](m.ID, objc.Sel("learningRate"))
	return rv
}
func (m MLProgramTrainer) SetLearningRate(value float64) {
	objc.Send[struct{}](m.ID, objc.Sel("setLearningRate:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLProgramTrainer/program
func (m MLProgramTrainer) Program() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("program"))
	return objectivec.Object{ID: rv}
}
func (m MLProgramTrainer) SetProgram(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setProgram:"), value)
}
