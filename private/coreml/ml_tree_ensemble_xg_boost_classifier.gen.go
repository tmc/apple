// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
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
	rv := objc.Send[MLTreeEnsembleXGBoostClassifier](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
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
//   - [MLTreeEnsembleXGBoostClassifier.LoadLabelsWithStringLabelsIntLabels]
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
//   - [MLTreeEnsembleXGBoostClassifier.XgBoostDataFormatFromBatchProviderNeedLabelsError]
//   - [MLTreeEnsembleXGBoostClassifier.XgBoostDataFormatFromFeatureProviderError]
//   - [MLTreeEnsembleXGBoostClassifier.InitWithDescriptionConfigurationIndexToStringLabelArrayIndexToIntLabelArrayModelURLError]
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier
type MLTreeEnsembleXGBoostClassifier struct {
	objectivec.Object
}

// MLTreeEnsembleXGBoostClassifierFromID constructs a [MLTreeEnsembleXGBoostClassifier] from an objc.ID.
func MLTreeEnsembleXGBoostClassifierFromID(id objc.ID) MLTreeEnsembleXGBoostClassifier {
	return MLTreeEnsembleXGBoostClassifier{objectivec.Object{ID: id}}
}
// NOTE: MLTreeEnsembleXGBoostClassifier struct embeds objectivec.Object (parent type unavailable) but
// IMLTreeEnsembleXGBoostClassifier embeds the parent interface; skip compile-time assertion.

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
//   - [IMLTreeEnsembleXGBoostClassifier.LoadLabelsWithStringLabelsIntLabels]
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
//   - [IMLTreeEnsembleXGBoostClassifier.XgBoostDataFormatFromBatchProviderNeedLabelsError]
//   - [IMLTreeEnsembleXGBoostClassifier.XgBoostDataFormatFromFeatureProviderError]
//   - [IMLTreeEnsembleXGBoostClassifier.InitWithDescriptionConfigurationIndexToStringLabelArrayIndexToIntLabelArrayModelURLError]
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier
type IMLTreeEnsembleXGBoostClassifier interface {
	IMLClassifier

	// Topic: Methods

	BatchProviderFromXGboostResultsLengthError(results unsafe.Pointer, length uint64) (objectivec.IObject, error)
	Booster() unsafe.Pointer
	SetBooster(value unsafe.Pointer)
	FeatureProviderArrayFromXGBoostResultLengthError(result unsafe.Pointer, length uint64) (objectivec.IObject, error)
	FeatureProviderFromXGboostResultsLengthError(results unsafe.Pointer, length uint64) (objectivec.IObject, error)
	InitializeAndvalidateObjectiveAndNumClassesWithConfigurationError(configuration objectivec.IObject) (MLTreeEnsembleXGBoostClassifier, error)
	InitializeBoosterIfOneExists() MLTreeEnsembleXGBoostClassifier
	LabelsForDataPoints() foundation.INSArray
	SetLabelsForDataPoints(value foundation.INSArray)
	LoadLabelsWithStringLabelsIntLabels(labels objectivec.IObject, labels2 objectivec.IObject) objectivec.IObject
	ModelURL() foundation.INSURL
	SetModelURL(value foundation.INSURL)
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
	XgBoostDataFormatFromBatchProviderNeedLabelsError(provider objectivec.IObject, labels bool) (objectivec.IObject, error)
	XgBoostDataFormatFromFeatureProviderError(provider objectivec.IObject) (objectivec.IObject, error)
	InitWithDescriptionConfigurationIndexToStringLabelArrayIndexToIntLabelArrayModelURLError(description objectivec.IObject, configuration objectivec.IObject, array objectivec.IObject, array2 objectivec.IObject, url foundation.INSURL) (MLTreeEnsembleXGBoostClassifier, error)
}

// Init initializes the instance.
func (t MLTreeEnsembleXGBoostClassifier) Init() MLTreeEnsembleXGBoostClassifier {
	rv := objc.Send[MLTreeEnsembleXGBoostClassifier](t.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t MLTreeEnsembleXGBoostClassifier) Autorelease() MLTreeEnsembleXGBoostClassifier {
	rv := objc.Send[MLTreeEnsembleXGBoostClassifier](t.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLTreeEnsembleXGBoostClassifier creates a new MLTreeEnsembleXGBoostClassifier instance.
func NewMLTreeEnsembleXGBoostClassifier() MLTreeEnsembleXGBoostClassifier {
	class := getMLTreeEnsembleXGBoostClassifierClass()
	rv := objc.Send[MLTreeEnsembleXGBoostClassifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/initWithDescription:configuration:indexToStringLabelArray:indexToIntLabelArray:modelURL:error:
func NewTreeEnsembleXGBoostClassifierWithDescriptionConfigurationIndexToStringLabelArrayIndexToIntLabelArrayModelURLError(description objectivec.IObject, configuration objectivec.IObject, array objectivec.IObject, array2 objectivec.IObject, url foundation.INSURL) (MLTreeEnsembleXGBoostClassifier, error) {
	var errorPtr objc.ID
	instance := getMLTreeEnsembleXGBoostClassifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:indexToStringLabelArray:indexToIntLabelArray:modelURL:error:"), description, configuration, array, array2, url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLTreeEnsembleXGBoostClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLTreeEnsembleXGBoostClassifierFromID(rv), nil
}

//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/batchProviderFromXGboostResults:length:error:
func (t MLTreeEnsembleXGBoostClassifier) BatchProviderFromXGboostResultsLengthError(results unsafe.Pointer, length uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("batchProviderFromXGboostResults:length:error:"), results, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/featureProviderArrayFromXGBoostResult:length:error:
func (t MLTreeEnsembleXGBoostClassifier) FeatureProviderArrayFromXGBoostResultLengthError(result unsafe.Pointer, length uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("featureProviderArrayFromXGBoostResult:length:error:"), result, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/featureProviderFromXGboostResults:length:error:
func (t MLTreeEnsembleXGBoostClassifier) FeatureProviderFromXGboostResultsLengthError(results unsafe.Pointer, length uint64) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("featureProviderFromXGboostResults:length:error:"), results, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/initializeAndvalidateObjectiveAndNumClassesWithConfiguration:error:
func (t MLTreeEnsembleXGBoostClassifier) InitializeAndvalidateObjectiveAndNumClassesWithConfigurationError(configuration objectivec.IObject) (MLTreeEnsembleXGBoostClassifier, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("initializeAndvalidateObjectiveAndNumClassesWithConfiguration:error:"), configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLTreeEnsembleXGBoostClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLTreeEnsembleXGBoostClassifierFromID(rv), nil

}
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/initializeBoosterIfOneExists
func (t MLTreeEnsembleXGBoostClassifier) InitializeBoosterIfOneExists() MLTreeEnsembleXGBoostClassifier {
	rv := objc.Send[MLTreeEnsembleXGBoostClassifier](t.ID, objc.Sel("initializeBoosterIfOneExists"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/loadLabelsWithStringLabels:intLabels:
func (t MLTreeEnsembleXGBoostClassifier) LoadLabelsWithStringLabelsIntLabels(labels objectivec.IObject, labels2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("loadLabelsWithStringLabels:intLabels:"), labels, labels2)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/packageOutputWithPredictedLabel:classProbabilities:
func (t MLTreeEnsembleXGBoostClassifier) PackageOutputWithPredictedLabelClassProbabilities(label objectivec.IObject, probabilities objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("packageOutputWithPredictedLabel:classProbabilities:"), label, probabilities)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/populateXGBoostDataFormat:trainingData:dataIndex:inputName:needLabels:error:
func (t MLTreeEnsembleXGBoostClassifier) PopulateXGBoostDataFormatTrainingDataDataIndexInputNameNeedLabelsError(format unsafe.Pointer, data objectivec.IObject, index int64, name objectivec.IObject, labels bool) error {
	var errorPtr objc.ID
	objc.Send[struct{}](t.ID, objc.Sel("populateXGBoostDataFormat:trainingData:dataIndex:inputName:needLabels:error:"), format, data, index, name, labels, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSErrorFrom(errorPtr)
	}
	return nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/predictionFromFeatures:error:
func (t MLTreeEnsembleXGBoostClassifier) PredictionFromFeaturesError(features objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("predictionFromFeatures:error:"), features, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/predictionFromFeatures:options:error:
func (t MLTreeEnsembleXGBoostClassifier) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/predictionsFromBatch:error:
func (t MLTreeEnsembleXGBoostClassifier) PredictionsFromBatchError(batch objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("predictionsFromBatch:error:"), batch, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/predictionsFromBatch:options:error:
func (t MLTreeEnsembleXGBoostClassifier) PredictionsFromBatchOptionsError(batch objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("predictionsFromBatch:options:error:"), batch, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/xgBoostDataFormatFromBatchProvider:needLabels:error:
func (t MLTreeEnsembleXGBoostClassifier) XgBoostDataFormatFromBatchProviderNeedLabelsError(provider objectivec.IObject, labels bool) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("xgBoostDataFormatFromBatchProvider:needLabels:error:"), provider, labels, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/xgBoostDataFormatFromFeatureProvider:error:
func (t MLTreeEnsembleXGBoostClassifier) XgBoostDataFormatFromFeatureProviderError(provider objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("xgBoostDataFormatFromFeatureProvider:error:"), provider, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/initWithDescription:configuration:indexToStringLabelArray:indexToIntLabelArray:modelURL:error:
func (t MLTreeEnsembleXGBoostClassifier) InitWithDescriptionConfigurationIndexToStringLabelArrayIndexToIntLabelArrayModelURLError(description objectivec.IObject, configuration objectivec.IObject, array objectivec.IObject, array2 objectivec.IObject, url foundation.INSURL) (MLTreeEnsembleXGBoostClassifier, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](t.ID, objc.Sel("initWithDescription:configuration:indexToStringLabelArray:indexToIntLabelArray:modelURL:error:"), description, configuration, array, array2, url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLTreeEnsembleXGBoostClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLTreeEnsembleXGBoostClassifierFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/booster
func (t MLTreeEnsembleXGBoostClassifier) Booster() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t.ID, objc.Sel("booster"))
	return rv
}
func (t MLTreeEnsembleXGBoostClassifier) SetBooster(value unsafe.Pointer) {
	objc.Send[struct{}](t.ID, objc.Sel("setBooster:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/labelsForDataPoints
func (t MLTreeEnsembleXGBoostClassifier) LabelsForDataPoints() foundation.INSArray {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("labelsForDataPoints"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (t MLTreeEnsembleXGBoostClassifier) SetLabelsForDataPoints(value foundation.INSArray) {
	objc.Send[struct{}](t.ID, objc.Sel("setLabelsForDataPoints:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/modelURL
func (t MLTreeEnsembleXGBoostClassifier) ModelURL() foundation.INSURL {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("modelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (t MLTreeEnsembleXGBoostClassifier) SetModelURL(value foundation.INSURL) {
	objc.Send[struct{}](t.ID, objc.Sel("setModelURL:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/numClasses
func (t MLTreeEnsembleXGBoostClassifier) NumClasses() uint64 {
	rv := objc.Send[uint64](t.ID, objc.Sel("numClasses"))
	return rv
}
func (t MLTreeEnsembleXGBoostClassifier) SetNumClasses(value uint64) {
	objc.Send[struct{}](t.ID, objc.Sel("setNumClasses:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/objective
func (t MLTreeEnsembleXGBoostClassifier) Objective() string {
	rv := objc.Send[objc.ID](t.ID, objc.Sel("objective"))
	return foundation.NSStringFromID(rv).String()
}
func (t MLTreeEnsembleXGBoostClassifier) SetObjective(value string) {
	objc.Send[struct{}](t.ID, objc.Sel("setObjective:"), objc.String(value))
}

