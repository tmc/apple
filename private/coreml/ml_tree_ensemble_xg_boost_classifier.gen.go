// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLTreeEnsembleXGBoostClassifier] class.
var (
	_MLTreeEnsembleXGBoostClassifierClass     MLTreeEnsembleXGBoostClassifierClass
	_MLTreeEnsembleXGBoostClassifierClassOnce sync.Once
)

func getMLTreeEnsembleXGBoostClassifierClass() MLTreeEnsembleXGBoostClassifierClass {
	_MLTreeEnsembleXGBoostClassifierClassOnce.Do(func() {
		_MLTreeEnsembleXGBoostClassifierClass = MLTreeEnsembleXGBoostClassifierClass{class: objc.GetClass("MLTreeEnsembleXGBoostClassifier")}
	})
	return _MLTreeEnsembleXGBoostClassifierClass
}

// GetMLTreeEnsembleXGBoostClassifierClass returns the class object for MLTreeEnsembleXGBoostClassifier.
func GetMLTreeEnsembleXGBoostClassifierClass() MLTreeEnsembleXGBoostClassifierClass {
	return getMLTreeEnsembleXGBoostClassifierClass()
}

type MLTreeEnsembleXGBoostClassifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLTreeEnsembleXGBoostClassifierClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLTreeEnsembleXGBoostClassifierClass) Alloc() MLTreeEnsembleXGBoostClassifier {
	rv := objc.SendIfResponds[MLTreeEnsembleXGBoostClassifier](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLTreeEnsembleXGBoostClassifier.BatchProviderFromXGboostResultsLengthError]
//   - [MLTreeEnsembleXGBoostClassifier.Booster]
//   - [MLTreeEnsembleXGBoostClassifier.SetBooster]
//   - [MLTreeEnsembleXGBoostClassifier.FeatureProviderArrayFromXGBoostResultLengthError]
//   - [MLTreeEnsembleXGBoostClassifier.FeatureProviderFromXGboostResultsLengthError]
//   - [MLTreeEnsembleXGBoostClassifier.InitializeAndvalidateObjectiveAndNumClassesWithConfigurationError]
//   - [MLTreeEnsembleXGBoostClassifier.InitializeBoosterIfOneExists]
//   - [MLTreeEnsembleXGBoostClassifier.LabelsForDataPoints]
//   - [MLTreeEnsembleXGBoostClassifier.SetLabelsForDataPoints]
//   - [MLTreeEnsembleXGBoostClassifier.ModelURL]
//   - [MLTreeEnsembleXGBoostClassifier.SetModelURL]
//   - [MLTreeEnsembleXGBoostClassifier.NumClasses]
//   - [MLTreeEnsembleXGBoostClassifier.SetNumClasses]
//   - [MLTreeEnsembleXGBoostClassifier.Objective]
//   - [MLTreeEnsembleXGBoostClassifier.SetObjective]
//   - [MLTreeEnsembleXGBoostClassifier.PackageOutputWithPredictedLabelClassProbabilities]
//   - [MLTreeEnsembleXGBoostClassifier.PopulateXGBoostDataFormatTrainingDataDataIndexInputNameNeedLabelsError]
//   - [MLTreeEnsembleXGBoostClassifier.PredictionFromFeaturesError]
//   - [MLTreeEnsembleXGBoostClassifier.PredictionFromFeaturesOptionsError]
//   - [MLTreeEnsembleXGBoostClassifier.PredictionsFromBatchError]
//   - [MLTreeEnsembleXGBoostClassifier.PredictionsFromBatchOptionsError]
type MLTreeEnsembleXGBoostClassifier struct {
	objectivec.Object
}

// MLTreeEnsembleXGBoostClassifierFromID constructs a [MLTreeEnsembleXGBoostClassifier] from an objc.ID.
func MLTreeEnsembleXGBoostClassifierFromID(id objc.ID) MLTreeEnsembleXGBoostClassifier {
	return MLTreeEnsembleXGBoostClassifier{objectivec.Object{ID: id}}
}

// Ensure MLTreeEnsembleXGBoostClassifier implements IMLTreeEnsembleXGBoostClassifier.
var _ IMLTreeEnsembleXGBoostClassifier = MLTreeEnsembleXGBoostClassifier{}

// An interface definition for the [MLTreeEnsembleXGBoostClassifier] class.
//
// # Methods
//
//   - [IMLTreeEnsembleXGBoostClassifier.BatchProviderFromXGboostResultsLengthError]
//   - [IMLTreeEnsembleXGBoostClassifier.Booster]
//   - [IMLTreeEnsembleXGBoostClassifier.SetBooster]
//   - [IMLTreeEnsembleXGBoostClassifier.FeatureProviderArrayFromXGBoostResultLengthError]
//   - [IMLTreeEnsembleXGBoostClassifier.FeatureProviderFromXGboostResultsLengthError]
//   - [IMLTreeEnsembleXGBoostClassifier.InitializeAndvalidateObjectiveAndNumClassesWithConfigurationError]
//   - [IMLTreeEnsembleXGBoostClassifier.InitializeBoosterIfOneExists]
//   - [IMLTreeEnsembleXGBoostClassifier.LabelsForDataPoints]
//   - [IMLTreeEnsembleXGBoostClassifier.SetLabelsForDataPoints]
//   - [IMLTreeEnsembleXGBoostClassifier.ModelURL]
//   - [IMLTreeEnsembleXGBoostClassifier.SetModelURL]
//   - [IMLTreeEnsembleXGBoostClassifier.NumClasses]
//   - [IMLTreeEnsembleXGBoostClassifier.SetNumClasses]
//   - [IMLTreeEnsembleXGBoostClassifier.Objective]
//   - [IMLTreeEnsembleXGBoostClassifier.SetObjective]
//   - [IMLTreeEnsembleXGBoostClassifier.PackageOutputWithPredictedLabelClassProbabilities]
//   - [IMLTreeEnsembleXGBoostClassifier.PopulateXGBoostDataFormatTrainingDataDataIndexInputNameNeedLabelsError]
//   - [IMLTreeEnsembleXGBoostClassifier.PredictionFromFeaturesError]
//   - [IMLTreeEnsembleXGBoostClassifier.PredictionFromFeaturesOptionsError]
//   - [IMLTreeEnsembleXGBoostClassifier.PredictionsFromBatchError]
//   - [IMLTreeEnsembleXGBoostClassifier.PredictionsFromBatchOptionsError]
type IMLTreeEnsembleXGBoostClassifier interface {
	objectivec.IObject

	// Topic: Methods

	BatchProviderFromXGboostResultsLengthError(results *float32, length uint64) (objectivec.IObject, error)
	Booster() unsafe.Pointer
	SetBooster(value unsafe.Pointer)
	FeatureProviderArrayFromXGBoostResultLengthError(result *float32, length uint64) (objectivec.IObject, error)
	FeatureProviderFromXGboostResultsLengthError(results *float32, length uint64) (objectivec.IObject, error)
	InitializeAndvalidateObjectiveAndNumClassesWithConfigurationError(configuration objectivec.IObject) (MLTreeEnsembleXGBoostClassifier, error)
	InitializeBoosterIfOneExists() MLTreeEnsembleXGBoostClassifier
	LabelsForDataPoints() foundation.INSArray
	SetLabelsForDataPoints(value foundation.INSArray)
	ModelURL() foundation.NSURL
	SetModelURL(value foundation.NSURL)
	NumClasses() uint64
	SetNumClasses(value uint64)
	Objective() string
	SetObjective(value string)
	PackageOutputWithPredictedLabelClassProbabilities(label objectivec.IObject, probabilities objectivec.IObject) objectivec.IObject
	PopulateXGBoostDataFormatTrainingDataDataIndexInputNameNeedLabelsError(format unsafe.Pointer, data objectivec.IObject, index int64, name objectivec.IObject, labels bool) error
	PredictionFromFeaturesError(features objectivec.IObject) (objectivec.IObject, error)
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	PredictionsFromBatchError(batch objectivec.IObject) (objectivec.IObject, error)
	PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
}

// Init initializes the instance.
func (m MLTreeEnsembleXGBoostClassifier) Init() MLTreeEnsembleXGBoostClassifier {
	rv := objc.SendIfResponds[MLTreeEnsembleXGBoostClassifier](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLTreeEnsembleXGBoostClassifier) Autorelease() MLTreeEnsembleXGBoostClassifier {
	rv := objc.SendIfResponds[MLTreeEnsembleXGBoostClassifier](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLTreeEnsembleXGBoostClassifier creates a new MLTreeEnsembleXGBoostClassifier instance.
func NewMLTreeEnsembleXGBoostClassifier() MLTreeEnsembleXGBoostClassifier {
	class := getMLTreeEnsembleXGBoostClassifierClass()
	rv := objc.SendIfResponds[MLTreeEnsembleXGBoostClassifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewTreeEnsembleXGBoostClassifierWithDescriptionConfigurationIndexToStringLabelArrayIndexToIntLabelArrayModelURLError(description objectivec.IObject, configuration objectivec.IObject, array unsafe.Pointer, array2 unsafe.Pointer, url foundation.NSURL) (MLTreeEnsembleXGBoostClassifier, error) {
	var errorPtr objc.ID
	instance := getMLTreeEnsembleXGBoostClassifierClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:indexToStringLabelArray:indexToIntLabelArray:modelURL:error:"), description, configuration, array, array2, url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLTreeEnsembleXGBoostClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLTreeEnsembleXGBoostClassifier{}, objc.ErrInitFailed
	}
	return MLTreeEnsembleXGBoostClassifierFromID(rv), nil
}

func (m MLTreeEnsembleXGBoostClassifier) BatchProviderFromXGboostResultsLengthError(results *float32, length uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("batchProviderFromXGboostResults:length:error:"), results, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLTreeEnsembleXGBoostClassifier) FeatureProviderArrayFromXGBoostResultLengthError(result *float32, length uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("featureProviderArrayFromXGBoostResult:length:error:"), result, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLTreeEnsembleXGBoostClassifier) FeatureProviderFromXGboostResultsLengthError(results *float32, length uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("featureProviderFromXGboostResults:length:error:"), results, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLTreeEnsembleXGBoostClassifier) InitializeAndvalidateObjectiveAndNumClassesWithConfigurationError(configuration objectivec.IObject) (MLTreeEnsembleXGBoostClassifier, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initializeAndvalidateObjectiveAndNumClassesWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLTreeEnsembleXGBoostClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLTreeEnsembleXGBoostClassifierFromID(rv), nil

}
func (m MLTreeEnsembleXGBoostClassifier) InitializeBoosterIfOneExists() MLTreeEnsembleXGBoostClassifier {
	rv := objc.SendIfResponds[MLTreeEnsembleXGBoostClassifier](m.ID, objc.Sel("initializeBoosterIfOneExists"))
	return rv
}
func (m MLTreeEnsembleXGBoostClassifier) PackageOutputWithPredictedLabelClassProbabilities(label objectivec.IObject, probabilities objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("packageOutputWithPredictedLabel:classProbabilities:"), label, probabilities)
	return objectivec.Object{ID: rv}
}
func (m MLTreeEnsembleXGBoostClassifier) PopulateXGBoostDataFormatTrainingDataDataIndexInputNameNeedLabelsError(format unsafe.Pointer, data objectivec.IObject, index int64, name objectivec.IObject, labels bool) error {
	var errorPtr objc.ID
	objc.Send[struct{}](m.ID, objc.Sel("populateXGBoostDataFormat:trainingData:dataIndex:inputName:needLabels:error:"), format, data, index, name, labels, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSErrorFrom(errorPtr)
	}
	return nil

}
func (m MLTreeEnsembleXGBoostClassifier) PredictionFromFeaturesError(features objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:error:"), features, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLTreeEnsembleXGBoostClassifier) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLTreeEnsembleXGBoostClassifier) PredictionsFromBatchError(batch objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionsFromBatch:error:"), batch, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLTreeEnsembleXGBoostClassifier) PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionsFromBatch:options:error:"), batch, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

func (m MLTreeEnsembleXGBoostClassifier) Booster() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("booster"))
	return rv
}
func (m MLTreeEnsembleXGBoostClassifier) SetBooster(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setBooster:"), value)
}
func (m MLTreeEnsembleXGBoostClassifier) LabelsForDataPoints() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("labelsForDataPoints"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLTreeEnsembleXGBoostClassifier) SetLabelsForDataPoints(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setLabelsForDataPoints:"), value)
}
func (m MLTreeEnsembleXGBoostClassifier) ModelURL() foundation.NSURL {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (m MLTreeEnsembleXGBoostClassifier) SetModelURL(value foundation.NSURL) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setModelURL:"), value)
}
func (m MLTreeEnsembleXGBoostClassifier) NumClasses() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("numClasses"))
	return rv
}
func (m MLTreeEnsembleXGBoostClassifier) SetNumClasses(value uint64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setNumClasses:"), value)
}
func (m MLTreeEnsembleXGBoostClassifier) Objective() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("objective"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLTreeEnsembleXGBoostClassifier) SetObjective(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setObjective:"), objc.String(value))
}
