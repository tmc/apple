// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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
	rv := objc.SendIfResponds[MLBayesianProbitRegression](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLBayesianProbitRegression.FeatureCount]
//   - [MLBayesianProbitRegression.ConvertOutputFeatureToPredictionValuesEventImportanceError]
//   - [MLBayesianProbitRegression.CreateCheckpoint]
//   - [MLBayesianProbitRegression.CreateRegressorResult]
//   - [MLBayesianProbitRegression.GetArrayFeatureValue]
//   - [MLBayesianProbitRegression.GetFeatureValueForNameWithType]
//   - [MLBayesianProbitRegression.GetOptimism]
//   - [MLBayesianProbitRegression.GetSamplingScale]
//   - [MLBayesianProbitRegression.GetSamplingTruncation]
//   - [MLBayesianProbitRegression.IsEqualToBopr]
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
type MLBayesianProbitRegression struct {
	objectivec.Object
}

// MLBayesianProbitRegressionFromID constructs a [MLBayesianProbitRegression] from an objc.ID.
func MLBayesianProbitRegressionFromID(id objc.ID) MLBayesianProbitRegression {
	return MLBayesianProbitRegression{objectivec.Object{ID: id}}
}

// Ensure MLBayesianProbitRegression implements IMLBayesianProbitRegression.
var _ IMLBayesianProbitRegression = MLBayesianProbitRegression{}

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
//   - [IMLBayesianProbitRegression.GetOptimism]
//   - [IMLBayesianProbitRegression.GetSamplingScale]
//   - [IMLBayesianProbitRegression.GetSamplingTruncation]
//   - [IMLBayesianProbitRegression.IsEqualToBopr]
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
type IMLBayesianProbitRegression interface {
	objectivec.IObject

	// Topic: Methods

	FeatureCount() uint64
	ConvertOutputFeatureToPredictionValuesEventImportanceError(values objectivec.IObject) (bool, float64, error)
	CreateCheckpoint()
	CreateRegressorResult(result *Prediction) objectivec.IObject
	GetArrayFeatureValue(value objectivec.IObject) objectivec.IObject
	GetFeatureValueForNameWithType(value objectivec.IObject, name objectivec.IObject, type_ int64) float64
	GetOptimism(optimism objectivec.IObject) float64
	GetSamplingScale(scale objectivec.IObject) float64
	GetSamplingTruncation(truncation objectivec.IObject) float64
	IsEqualToBopr(bopr objectivec.IObject) bool
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
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (m MLBayesianProbitRegression) Init() MLBayesianProbitRegression {
	rv := objc.SendIfResponds[MLBayesianProbitRegression](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLBayesianProbitRegression) Autorelease() MLBayesianProbitRegression {
	rv := objc.SendIfResponds[MLBayesianProbitRegression](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLBayesianProbitRegression creates a new MLBayesianProbitRegression instance.
func NewMLBayesianProbitRegression() MLBayesianProbitRegression {
	class := getMLBayesianProbitRegressionClass()
	rv := objc.SendIfResponds[MLBayesianProbitRegression](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewBayesianProbitRegressionWithDescriptionNumberOfFeaturesPriorMean(description objectivec.IObject, features int64, mean objectivec.IObject) MLBayesianProbitRegression {
	instance := getMLBayesianProbitRegressionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescription:numberOfFeatures:priorMean:"), description, features, mean)
	return MLBayesianProbitRegressionFromID(rv)
}

func NewBayesianProbitRegressionWithDescriptionNumberOfFeaturesPriorMeanRegressionInputNameOptimismInputNameSamplingScaleInputNameSamplingTruncationInputNameMeanOutputNameVarianceOutputNamePessimisticProbabilityOutputNameSampledProbabilityOutputName(description objectivec.IObject, features int64, mean objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, name5 objectivec.IObject, name6 objectivec.IObject, name7 objectivec.IObject, name8 objectivec.IObject) MLBayesianProbitRegression {
	instance := getMLBayesianProbitRegressionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescription:numberOfFeatures:priorMean:regressionInputName:optimismInputName:samplingScaleInputName:samplingTruncationInputName:meanOutputName:varianceOutputName:pessimisticProbabilityOutputName:sampledProbabilityOutputName:"), description, features, mean, name, name2, name3, name4, name5, name6, name7, name8)
	return MLBayesianProbitRegressionFromID(rv)
}

func NewBayesianProbitRegressionWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLBayesianProbitRegression, error) {
	var errorPtr objc.ID
	instance := getMLBayesianProbitRegressionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLBayesianProbitRegression{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLBayesianProbitRegression{}, objc.ErrInitFailed
	}
	return MLBayesianProbitRegressionFromID(rv), nil
}

func (m MLBayesianProbitRegression) FeatureCount() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("FeatureCount"))
	return rv
}
func (m MLBayesianProbitRegression) ConvertOutputFeatureToPredictionValuesEventImportanceError(values objectivec.IObject) (bool, float64, error) {
	var event bool
	var importance float64
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("convertOutputFeatureToPredictionValues:event:importance:error:"), values, unsafe.Pointer(&event), unsafe.Pointer(&importance), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, 0.0, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, 0.0, errors.New("convertOutputFeatureToPredictionValues:event:importance:error: returned NO with nil NSError")
	}
	return event, importance, nil
}
func (m MLBayesianProbitRegression) CreateCheckpoint() {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("createCheckpoint"))
}
func (m MLBayesianProbitRegression) CreateRegressorResult(result *Prediction) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("createRegressorResult:"), unsafe.Pointer(result))
	return objectivec.Object{ID: rv}
}
func (m MLBayesianProbitRegression) GetArrayFeatureValue(value objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("getArrayFeatureValue:"), value)
	return objectivec.Object{ID: rv}
}
func (m MLBayesianProbitRegression) GetFeatureValueForNameWithType(value objectivec.IObject, name objectivec.IObject, type_ int64) float64 {
	rv := objc.SendIfResponds[float64](m.ID, objc.Sel("getFeatureValue:forName:withType:"), value, name, type_)
	return rv
}
func (m MLBayesianProbitRegression) GetOptimism(optimism objectivec.IObject) float64 {
	rv := objc.SendIfResponds[float64](m.ID, objc.Sel("getOptimism:"), optimism)
	return rv
}
func (m MLBayesianProbitRegression) GetSamplingScale(scale objectivec.IObject) float64 {
	rv := objc.SendIfResponds[float64](m.ID, objc.Sel("getSamplingScale:"), scale)
	return rv
}
func (m MLBayesianProbitRegression) GetSamplingTruncation(truncation objectivec.IObject) float64 {
	rv := objc.SendIfResponds[float64](m.ID, objc.Sel("getSamplingTruncation:"), truncation)
	return rv
}
func (m MLBayesianProbitRegression) IsEqualToBopr(bopr objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("isEqualToBopr:"), bopr)
	return rv
}
func (m MLBayesianProbitRegression) RegressOptionsError(regress objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("regress:options:error:"), regress, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLBayesianProbitRegression) Reset() {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("reset"))
}
func (m MLBayesianProbitRegression) ResetToLastCheckpointBeforeDate(date objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("resetToLastCheckpointBeforeDate:"), date)
}
func (m MLBayesianProbitRegression) SaveModelToSpecification(specification []objectivec.IObject) unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("saveModelToSpecification:"), objectivec.IObjectSliceToNSArray(specification))
	return rv
}
func (m MLBayesianProbitRegression) SetFeatureCount(count uint64) bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("setFeatureCount:"), count)
	return rv
}
func (m MLBayesianProbitRegression) SetInputFeatureNameTo(name []objectivec.IObject, to objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("setInputFeatureName:to:"), objectivec.IObjectSliceToNSArray(name), to)
	return rv
}
func (m MLBayesianProbitRegression) SetOutputFeatureNameTo(name []objectivec.IObject, to objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("setOutputFeatureName:to:"), objectivec.IObjectSliceToNSArray(name), to)
	return rv
}
func (m MLBayesianProbitRegression) UpdateModelFromFeaturesToTargetError(features objectivec.IObject, target objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("updateModelFromFeatures:toTarget:error:"), features, target, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateModelFromFeatures:toTarget:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLBayesianProbitRegression) UpdateModelFromFeaturesToTargetOptionsError(features objectivec.IObject, target objectivec.IObject, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("updateModelFromFeatures:toTarget:options:error:"), features, target, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateModelFromFeatures:toTarget:options:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLBayesianProbitRegression) InitWithDescriptionNumberOfFeaturesPriorMean(description objectivec.IObject, features int64, mean objectivec.IObject) MLBayesianProbitRegression {
	rv := objc.SendIfResponds[MLBayesianProbitRegression](m.ID, objc.Sel("initWithDescription:numberOfFeatures:priorMean:"), description, features, mean)
	return rv
}
func (m MLBayesianProbitRegression) InitWithDescriptionNumberOfFeaturesPriorMeanRegressionInputNameOptimismInputNameSamplingScaleInputNameSamplingTruncationInputNameMeanOutputNameVarianceOutputNamePessimisticProbabilityOutputNameSampledProbabilityOutputName(description objectivec.IObject, features int64, mean objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, name5 objectivec.IObject, name6 objectivec.IObject, name7 objectivec.IObject, name8 objectivec.IObject) MLBayesianProbitRegression {
	rv := objc.SendIfResponds[MLBayesianProbitRegression](m.ID, objc.Sel("initWithDescription:numberOfFeatures:priorMean:regressionInputName:optimismInputName:samplingScaleInputName:samplingTruncationInputName:meanOutputName:varianceOutputName:pessimisticProbabilityOutputName:sampledProbabilityOutputName:"), description, features, mean, name, name2, name3, name4, name5, name6, name7, name8)
	return rv
}
func (m MLBayesianProbitRegression) InitWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLBayesianProbitRegression, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLBayesianProbitRegression{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLBayesianProbitRegressionFromID(rv), nil

}

func (_MLBayesianProbitRegressionClass MLBayesianProbitRegressionClass) LoadModelFromSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLBayesianProbitRegressionClass.class), objc.Sel("loadModelFromSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLBayesianProbitRegressionClass MLBayesianProbitRegressionClass) ModelWithContentsOfURLError(url foundation.NSURL) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLBayesianProbitRegressionClass.class), objc.Sel("modelWithContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLBayesianProbitRegressionClass MLBayesianProbitRegressionClass) SetFeatureNameToDescriptions(name []objectivec.IObject, to objectivec.IObject, descriptions objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLBayesianProbitRegressionClass.class), objc.Sel("setFeatureName:to:descriptions:"), objectivec.IObjectSliceToNSArray(name), to, descriptions)
	return rv
}
func (_MLBayesianProbitRegressionClass MLBayesianProbitRegressionClass) ValidateModelDescription(description objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLBayesianProbitRegressionClass.class), objc.Sel("validateModelDescription:"), description)
	return rv
}

func (m MLBayesianProbitRegression) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLBayesianProbitRegression) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLBayesianProbitRegression) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLBayesianProbitRegression) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
