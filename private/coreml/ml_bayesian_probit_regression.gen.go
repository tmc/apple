// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLBayesianProbitRegression] class.
var (
	_MLBayesianProbitRegressionClass     MLBayesianProbitRegressionClass
	_MLBayesianProbitRegressionClassOnce sync.Once
)

func getMLBayesianProbitRegressionClass() MLBayesianProbitRegressionClass {
	_MLBayesianProbitRegressionClassOnce.Do(func() {
		_MLBayesianProbitRegressionClass = MLBayesianProbitRegressionClass{class: objc.GetClass("MLBayesianProbitRegression")}
	})
	return _MLBayesianProbitRegressionClass
}

// GetMLBayesianProbitRegressionClass returns the class object for MLBayesianProbitRegression.
func GetMLBayesianProbitRegressionClass() MLBayesianProbitRegressionClass {
	return getMLBayesianProbitRegressionClass()
}

type MLBayesianProbitRegressionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLBayesianProbitRegressionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLBayesianProbitRegressionClass) Alloc() MLBayesianProbitRegression {
	rv := objc.Send[MLBayesianProbitRegression](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLBayesianProbitRegression.FeatureCount]
//   - [MLBayesianProbitRegression.ConvertOutputFeatureToPredictionValuesEventImportanceError]
//   - [MLBayesianProbitRegression.CreateCheckpoint]
//   - [MLBayesianProbitRegression.CreateRegressorResult]
//   - [MLBayesianProbitRegression.GetArrayFeatureValue]
//   - [MLBayesianProbitRegression.GetFeatureValueForNameWithType]
//   - [MLBayesianProbitRegression.GetOneHotFeatureValuesError]
//   - [MLBayesianProbitRegression.GetOptimism]
//   - [MLBayesianProbitRegression.GetSamplingScale]
//   - [MLBayesianProbitRegression.GetSamplingTruncation]
//   - [MLBayesianProbitRegression.IsEqualToBopr]
//   - [MLBayesianProbitRegression.Model]
//   - [MLBayesianProbitRegression.RegressOptionsError]
//   - [MLBayesianProbitRegression.Reset]
//   - [MLBayesianProbitRegression.ResetToLastCheckpointBeforeDate]
//   - [MLBayesianProbitRegression.SaveModelToSpecification]
//   - [MLBayesianProbitRegression.SetFeatureCount]
//   - [MLBayesianProbitRegression.SetInputFeatureNameTo]
//   - [MLBayesianProbitRegression.SetOutputFeatureNameTo]
//   - [MLBayesianProbitRegression.UpdateModelFromFeaturesToTargetError]
//   - [MLBayesianProbitRegression.UpdateModelFromFeaturesToTargetOptionsError]
//   - [MLBayesianProbitRegression.InitWithDescriptionNumberOfFeaturesPriorMean]
//   - [MLBayesianProbitRegression.InitWithDescriptionNumberOfFeaturesPriorMeanRegressionInputNameOptimismInputNameSamplingScaleInputNameSamplingTruncationInputNameMeanOutputNameVarianceOutputNamePessimisticProbabilityOutputNameSampledProbabilityOutputName]
//   - [MLBayesianProbitRegression.InitWithSpecificationConfigurationError]
//   - [MLBayesianProbitRegression.DebugDescription]
//   - [MLBayesianProbitRegression.Description]
//   - [MLBayesianProbitRegression.Hash]
//   - [MLBayesianProbitRegression.Superclass]
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression
type MLBayesianProbitRegression struct {
	objectivec.Object
}

// MLBayesianProbitRegressionFromID constructs a [MLBayesianProbitRegression] from an objc.ID.
func MLBayesianProbitRegressionFromID(id objc.ID) MLBayesianProbitRegression {
	return MLBayesianProbitRegression{objectivec.Object{ID: id}}
}
// NOTE: MLBayesianProbitRegression struct embeds objectivec.Object (parent type unavailable) but
// IMLBayesianProbitRegression embeds the parent interface; skip compile-time assertion.

// An interface definition for the [MLBayesianProbitRegression] class.
//
// # Methods
//
//   - [IMLBayesianProbitRegression.FeatureCount]
//   - [IMLBayesianProbitRegression.ConvertOutputFeatureToPredictionValuesEventImportanceError]
//   - [IMLBayesianProbitRegression.CreateCheckpoint]
//   - [IMLBayesianProbitRegression.CreateRegressorResult]
//   - [IMLBayesianProbitRegression.GetArrayFeatureValue]
//   - [IMLBayesianProbitRegression.GetFeatureValueForNameWithType]
//   - [IMLBayesianProbitRegression.GetOneHotFeatureValuesError]
//   - [IMLBayesianProbitRegression.GetOptimism]
//   - [IMLBayesianProbitRegression.GetSamplingScale]
//   - [IMLBayesianProbitRegression.GetSamplingTruncation]
//   - [IMLBayesianProbitRegression.IsEqualToBopr]
//   - [IMLBayesianProbitRegression.Model]
//   - [IMLBayesianProbitRegression.RegressOptionsError]
//   - [IMLBayesianProbitRegression.Reset]
//   - [IMLBayesianProbitRegression.ResetToLastCheckpointBeforeDate]
//   - [IMLBayesianProbitRegression.SaveModelToSpecification]
//   - [IMLBayesianProbitRegression.SetFeatureCount]
//   - [IMLBayesianProbitRegression.SetInputFeatureNameTo]
//   - [IMLBayesianProbitRegression.SetOutputFeatureNameTo]
//   - [IMLBayesianProbitRegression.UpdateModelFromFeaturesToTargetError]
//   - [IMLBayesianProbitRegression.UpdateModelFromFeaturesToTargetOptionsError]
//   - [IMLBayesianProbitRegression.InitWithDescriptionNumberOfFeaturesPriorMean]
//   - [IMLBayesianProbitRegression.InitWithDescriptionNumberOfFeaturesPriorMeanRegressionInputNameOptimismInputNameSamplingScaleInputNameSamplingTruncationInputNameMeanOutputNameVarianceOutputNamePessimisticProbabilityOutputNameSampledProbabilityOutputName]
//   - [IMLBayesianProbitRegression.InitWithSpecificationConfigurationError]
//   - [IMLBayesianProbitRegression.DebugDescription]
//   - [IMLBayesianProbitRegression.Description]
//   - [IMLBayesianProbitRegression.Hash]
//   - [IMLBayesianProbitRegression.Superclass]
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression
type IMLBayesianProbitRegression interface {
	IMLRegressor

	// Topic: Methods

	FeatureCount() uint64
	ConvertOutputFeatureToPredictionValuesEventImportanceError(values objectivec.IObject) (bool, float64, error)
	CreateCheckpoint()
	CreateRegressorResult(result unsafe.Pointer) objectivec.IObject
	GetArrayFeatureValue(value objectivec.IObject) objectivec.IObject
	GetFeatureValueForNameWithType(value objectivec.IObject, name objectivec.IObject, type_ int64) float64
	GetOneHotFeatureValuesError(values objectivec.IObject) (objectivec.IObject, error)
	GetOptimism(optimism objectivec.IObject) float64
	GetSamplingScale(scale objectivec.IObject) float64
	GetSamplingTruncation(truncation objectivec.IObject) float64
	IsEqualToBopr(bopr objectivec.IObject) bool
	Model() objectivec.IObject
	RegressOptionsError(regress objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	Reset()
	ResetToLastCheckpointBeforeDate(date objectivec.IObject)
	SaveModelToSpecification(specification []objectivec.IObject) unsafe.Pointer
	SetFeatureCount(count uint64) bool
	SetInputFeatureNameTo(name []objectivec.IObject, to objectivec.IObject) bool
	SetOutputFeatureNameTo(name []objectivec.IObject, to objectivec.IObject) bool
	UpdateModelFromFeaturesToTargetError(features objectivec.IObject, target objectivec.IObject) (bool, error)
	UpdateModelFromFeaturesToTargetOptionsError(features objectivec.IObject, target objectivec.IObject, options objectivec.IObject) (bool, error)
	InitWithDescriptionNumberOfFeaturesPriorMean(description objectivec.IObject, features int64, mean objectivec.IObject) MLBayesianProbitRegression
	InitWithDescriptionNumberOfFeaturesPriorMeanRegressionInputNameOptimismInputNameSamplingScaleInputNameSamplingTruncationInputNameMeanOutputNameVarianceOutputNamePessimisticProbabilityOutputNameSampledProbabilityOutputName(description objectivec.IObject, features int64, mean objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, name5 objectivec.IObject, name6 objectivec.IObject, name7 objectivec.IObject, name8 objectivec.IObject) MLBayesianProbitRegression
	InitWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLBayesianProbitRegression, error)
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objc.Class
}

// Init initializes the instance.
func (b MLBayesianProbitRegression) Init() MLBayesianProbitRegression {
	rv := objc.Send[MLBayesianProbitRegression](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b MLBayesianProbitRegression) Autorelease() MLBayesianProbitRegression {
	rv := objc.Send[MLBayesianProbitRegression](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLBayesianProbitRegression creates a new MLBayesianProbitRegression instance.
func NewMLBayesianProbitRegression() MLBayesianProbitRegression {
	class := getMLBayesianProbitRegressionClass()
	rv := objc.Send[MLBayesianProbitRegression](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/initWithDescription:numberOfFeatures:priorMean:
func NewBayesianProbitRegressionWithDescriptionNumberOfFeaturesPriorMean(description objectivec.IObject, features int64, mean objectivec.IObject) MLBayesianProbitRegression {
	instance := getMLBayesianProbitRegressionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:numberOfFeatures:priorMean:"), description, features, mean)
	return MLBayesianProbitRegressionFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/initWithDescription:numberOfFeatures:priorMean:regressionInputName:optimismInputName:samplingScaleInputName:samplingTruncationInputName:meanOutputName:varianceOutputName:pessimisticProbabilityOutputName:sampledProbabilityOutputName:
func NewBayesianProbitRegressionWithDescriptionNumberOfFeaturesPriorMeanRegressionInputNameOptimismInputNameSamplingScaleInputNameSamplingTruncationInputNameMeanOutputNameVarianceOutputNamePessimisticProbabilityOutputNameSampledProbabilityOutputName(description objectivec.IObject, features int64, mean objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, name5 objectivec.IObject, name6 objectivec.IObject, name7 objectivec.IObject, name8 objectivec.IObject) MLBayesianProbitRegression {
	instance := getMLBayesianProbitRegressionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:numberOfFeatures:priorMean:regressionInputName:optimismInputName:samplingScaleInputName:samplingTruncationInputName:meanOutputName:varianceOutputName:pessimisticProbabilityOutputName:sampledProbabilityOutputName:"), description, features, mean, name, name2, name3, name4, name5, name6, name7, name8)
	return MLBayesianProbitRegressionFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/initWithSpecification:configuration:error:
func NewBayesianProbitRegressionWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLBayesianProbitRegression, error) {
	var errorPtr objc.ID
	instance := getMLBayesianProbitRegressionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLBayesianProbitRegression{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLBayesianProbitRegressionFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/FeatureCount
func (b MLBayesianProbitRegression) FeatureCount() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("FeatureCount"))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/convertOutputFeatureToPredictionValues:event:importance:error:
func (b MLBayesianProbitRegression) ConvertOutputFeatureToPredictionValuesEventImportanceError(values objectivec.IObject) (bool, float64, error) {
	var event bool
	var importance float64
	var errorPtr objc.ID
	rv := objc.Send[bool](b.ID, objc.Sel("convertOutputFeatureToPredictionValues:event:importance:error:"), values, unsafe.Pointer(&event), unsafe.Pointer(&importance), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, 0.0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, 0.0, errors.New("convertOutputFeatureToPredictionValues:event:importance:error: returned NO with nil NSError")
	}
	return event, importance, nil
}
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/createCheckpoint
func (b MLBayesianProbitRegression) CreateCheckpoint() {
	objc.Send[objc.ID](b.ID, objc.Sel("createCheckpoint"))
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/createRegressorResult:
func (b MLBayesianProbitRegression) CreateRegressorResult(result unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("createRegressorResult:"), result)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/getArrayFeatureValue:
func (b MLBayesianProbitRegression) GetArrayFeatureValue(value objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("getArrayFeatureValue:"), value)
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/getFeatureValue:forName:withType:
func (b MLBayesianProbitRegression) GetFeatureValueForNameWithType(value objectivec.IObject, name objectivec.IObject, type_ int64) float64 {
	rv := objc.Send[float64](b.ID, objc.Sel("getFeatureValue:forName:withType:"), value, name, type_)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/getOneHotFeatureValues:error:
func (b MLBayesianProbitRegression) GetOneHotFeatureValuesError(values objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](b.ID, objc.Sel("getOneHotFeatureValues:error:"), values, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/getOptimism:
func (b MLBayesianProbitRegression) GetOptimism(optimism objectivec.IObject) float64 {
	rv := objc.Send[float64](b.ID, objc.Sel("getOptimism:"), optimism)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/getSamplingScale:
func (b MLBayesianProbitRegression) GetSamplingScale(scale objectivec.IObject) float64 {
	rv := objc.Send[float64](b.ID, objc.Sel("getSamplingScale:"), scale)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/getSamplingTruncation:
func (b MLBayesianProbitRegression) GetSamplingTruncation(truncation objectivec.IObject) float64 {
	rv := objc.Send[float64](b.ID, objc.Sel("getSamplingTruncation:"), truncation)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/isEqualToBopr:
func (b MLBayesianProbitRegression) IsEqualToBopr(bopr objectivec.IObject) bool {
	rv := objc.Send[bool](b.ID, objc.Sel("isEqualToBopr:"), bopr)
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/model
func (b MLBayesianProbitRegression) Model() objectivec.IObject {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("model"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/regress:options:error:
func (b MLBayesianProbitRegression) RegressOptionsError(regress objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](b.ID, objc.Sel("regress:options:error:"), regress, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/reset
func (b MLBayesianProbitRegression) Reset() {
	objc.Send[objc.ID](b.ID, objc.Sel("reset"))
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/resetToLastCheckpointBeforeDate:
func (b MLBayesianProbitRegression) ResetToLastCheckpointBeforeDate(date objectivec.IObject) {
	objc.Send[objc.ID](b.ID, objc.Sel("resetToLastCheckpointBeforeDate:"), date)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/saveModelToSpecification:
func (b MLBayesianProbitRegression) SaveModelToSpecification(specification []objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](b.ID, objc.Sel("saveModelToSpecification:"), objectivec.IObjectSliceToNSArray(specification))
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/setFeatureCount:
func (b MLBayesianProbitRegression) SetFeatureCount(count uint64) bool {
	rv := objc.Send[bool](b.ID, objc.Sel("setFeatureCount:"), count)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/setInputFeatureName:to:
func (b MLBayesianProbitRegression) SetInputFeatureNameTo(name []objectivec.IObject, to objectivec.IObject) bool {
	rv := objc.Send[bool](b.ID, objc.Sel("setInputFeatureName:to:"), objectivec.IObjectSliceToNSArray(name), to)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/setOutputFeatureName:to:
func (b MLBayesianProbitRegression) SetOutputFeatureNameTo(name []objectivec.IObject, to objectivec.IObject) bool {
	rv := objc.Send[bool](b.ID, objc.Sel("setOutputFeatureName:to:"), objectivec.IObjectSliceToNSArray(name), to)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/updateModelFromFeatures:toTarget:error:
func (b MLBayesianProbitRegression) UpdateModelFromFeaturesToTargetError(features objectivec.IObject, target objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](b.ID, objc.Sel("updateModelFromFeatures:toTarget:error:"), features, target, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateModelFromFeatures:toTarget:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/updateModelFromFeatures:toTarget:options:error:
func (b MLBayesianProbitRegression) UpdateModelFromFeaturesToTargetOptionsError(features objectivec.IObject, target objectivec.IObject, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](b.ID, objc.Sel("updateModelFromFeatures:toTarget:options:error:"), features, target, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateModelFromFeatures:toTarget:options:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/initWithDescription:numberOfFeatures:priorMean:
func (b MLBayesianProbitRegression) InitWithDescriptionNumberOfFeaturesPriorMean(description objectivec.IObject, features int64, mean objectivec.IObject) MLBayesianProbitRegression {
	rv := objc.Send[MLBayesianProbitRegression](b.ID, objc.Sel("initWithDescription:numberOfFeatures:priorMean:"), description, features, mean)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/initWithDescription:numberOfFeatures:priorMean:regressionInputName:optimismInputName:samplingScaleInputName:samplingTruncationInputName:meanOutputName:varianceOutputName:pessimisticProbabilityOutputName:sampledProbabilityOutputName:
func (b MLBayesianProbitRegression) InitWithDescriptionNumberOfFeaturesPriorMeanRegressionInputNameOptimismInputNameSamplingScaleInputNameSamplingTruncationInputNameMeanOutputNameVarianceOutputNamePessimisticProbabilityOutputNameSampledProbabilityOutputName(description objectivec.IObject, features int64, mean objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, name5 objectivec.IObject, name6 objectivec.IObject, name7 objectivec.IObject, name8 objectivec.IObject) MLBayesianProbitRegression {
	rv := objc.Send[MLBayesianProbitRegression](b.ID, objc.Sel("initWithDescription:numberOfFeatures:priorMean:regressionInputName:optimismInputName:samplingScaleInputName:samplingTruncationInputName:meanOutputName:varianceOutputName:pessimisticProbabilityOutputName:sampledProbabilityOutputName:"), description, features, mean, name, name2, name3, name4, name5, name6, name7, name8)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/initWithSpecification:configuration:error:
func (b MLBayesianProbitRegression) InitWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLBayesianProbitRegression, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](b.ID, objc.Sel("initWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLBayesianProbitRegression{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLBayesianProbitRegressionFromID(rv), nil

}

//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/loadModelFromSpecification:configuration:error:
func (_MLBayesianProbitRegressionClass MLBayesianProbitRegressionClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLBayesianProbitRegressionClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/modelWithContentsOfURL:error:
func (_MLBayesianProbitRegressionClass MLBayesianProbitRegressionClass) ModelWithContentsOfURLError(url foundation.INSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLBayesianProbitRegressionClass.class), objc.Sel("modelWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/setFeatureName:to:descriptions:
func (_MLBayesianProbitRegressionClass MLBayesianProbitRegressionClass) SetFeatureNameToDescriptions(name []objectivec.IObject, to objectivec.IObject, descriptions objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_MLBayesianProbitRegressionClass.class), objc.Sel("setFeatureName:to:descriptions:"), objectivec.IObjectSliceToNSArray(name), to, descriptions)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/validateModelDescription:
func (_MLBayesianProbitRegressionClass MLBayesianProbitRegressionClass) ValidateModelDescription(description objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_MLBayesianProbitRegressionClass.class), objc.Sel("validateModelDescription:"), description)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/debugDescription
func (b MLBayesianProbitRegression) DebugDescription() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/description
func (b MLBayesianProbitRegression) Description() string {
	rv := objc.Send[objc.ID](b.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/hash
func (b MLBayesianProbitRegression) Hash() uint64 {
	rv := objc.Send[uint64](b.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLBayesianProbitRegression/superclass
func (b MLBayesianProbitRegression) Superclass() objc.Class {
	rv := objc.Send[objc.Class](b.ID, objc.Sel("superclass"))
	return rv
}

