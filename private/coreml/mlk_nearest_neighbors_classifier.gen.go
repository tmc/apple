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

// The class instance for the [MLKNearestNeighborsClassifier] class.
var (
	_MLKNearestNeighborsClassifierClass     MLKNearestNeighborsClassifierClass
	_MLKNearestNeighborsClassifierClassOnce sync.Once
)

func getMLKNearestNeighborsClassifierClass() MLKNearestNeighborsClassifierClass {
	_MLKNearestNeighborsClassifierClassOnce.Do(func() {
		_MLKNearestNeighborsClassifierClass = MLKNearestNeighborsClassifierClass{class: objc.GetClass("MLKNearestNeighborsClassifier")}
	})
	return _MLKNearestNeighborsClassifierClass
}

// GetMLKNearestNeighborsClassifierClass returns the class object for MLKNearestNeighborsClassifier.
func GetMLKNearestNeighborsClassifierClass() MLKNearestNeighborsClassifierClass {
	return getMLKNearestNeighborsClassifierClass()
}

type MLKNearestNeighborsClassifierClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLKNearestNeighborsClassifierClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLKNearestNeighborsClassifierClass) Alloc() MLKNearestNeighborsClassifier {
	rv := objc.Send[MLKNearestNeighborsClassifier](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLKNearestNeighborsClassifier.CancelUpdate]
//   - [MLKNearestNeighborsClassifier.ComputeClassProbabilitiesFrom]
//   - [MLKNearestNeighborsClassifier.ComputeKClosestLabels]
//   - [MLKNearestNeighborsClassifier.ContinueWithUpdate]
//   - [MLKNearestNeighborsClassifier.SetContinueWithUpdate]
//   - [MLKNearestNeighborsClassifier.DefaultLabel]
//   - [MLKNearestNeighborsClassifier.SetDefaultLabel]
//   - [MLKNearestNeighborsClassifier.ExtractNearestNeighborLabelsDistancesIndicesFrom]
//   - [MLKNearestNeighborsClassifier.Index]
//   - [MLKNearestNeighborsClassifier.SetIndex]
//   - [MLKNearestNeighborsClassifier.IndexType]
//   - [MLKNearestNeighborsClassifier.SetIndexType]
//   - [MLKNearestNeighborsClassifier.InputMultiArrayError]
//   - [MLKNearestNeighborsClassifier.LabelType]
//   - [MLKNearestNeighborsClassifier.SetLabelType]
//   - [MLKNearestNeighborsClassifier.LabelsForDataPoints]
//   - [MLKNearestNeighborsClassifier.SetLabelsForDataPoints]
//   - [MLKNearestNeighborsClassifier.LabelsSet]
//   - [MLKNearestNeighborsClassifier.SetLabelsSet]
//   - [MLKNearestNeighborsClassifier.NearestDistancesFeatureName]
//   - [MLKNearestNeighborsClassifier.SetNearestDistancesFeatureName]
//   - [MLKNearestNeighborsClassifier.NearestLabelsFeatureName]
//   - [MLKNearestNeighborsClassifier.SetNearestLabelsFeatureName]
//   - [MLKNearestNeighborsClassifier.NumberOfDimensions]
//   - [MLKNearestNeighborsClassifier.SetNumberOfDimensions]
//   - [MLKNearestNeighborsClassifier.PackageOutputWithPredictedLabelClassProbabilitiesNearestLabelsNearestDistancesNearestFeatureIndices]
//   - [MLKNearestNeighborsClassifier.ParameterContainer]
//   - [MLKNearestNeighborsClassifier.SetParameterContainer]
//   - [MLKNearestNeighborsClassifier.ParameterValueForKeyError]
//   - [MLKNearestNeighborsClassifier.PredictionFromFeaturesOptionsError]
//   - [MLKNearestNeighborsClassifier.ProgressHandlers]
//   - [MLKNearestNeighborsClassifier.SetProgressHandlers]
//   - [MLKNearestNeighborsClassifier.ProgressHandlersDispatchQueue]
//   - [MLKNearestNeighborsClassifier.SetProgressHandlersDispatchQueue]
//   - [MLKNearestNeighborsClassifier.ResumeUpdate]
//   - [MLKNearestNeighborsClassifier.ResumeUpdateWithParameters]
//   - [MLKNearestNeighborsClassifier.SetUpdateProgressHandlersDispatchQueue]
//   - [MLKNearestNeighborsClassifier.UpdateModelWithData]
//   - [MLKNearestNeighborsClassifier.UpdateParameters]
//   - [MLKNearestNeighborsClassifier.SetUpdateParameters]
//   - [MLKNearestNeighborsClassifier.WeightingScheme]
//   - [MLKNearestNeighborsClassifier.SetWeightingScheme]
//   - [MLKNearestNeighborsClassifier.WriteToURLError]
//   - [MLKNearestNeighborsClassifier.InitWithDescriptionConfigurationParametersDataPointsLabelsError]
//   - [MLKNearestNeighborsClassifier.Metadata]
//
// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier
type MLKNearestNeighborsClassifier struct {
	MLModel
}

// MLKNearestNeighborsClassifierFromID constructs a [MLKNearestNeighborsClassifier] from an objc.ID.
func MLKNearestNeighborsClassifierFromID(id objc.ID) MLKNearestNeighborsClassifier {
	return MLKNearestNeighborsClassifier{MLModel: MLModelFromID(id)}
}

// Ensure MLKNearestNeighborsClassifier implements IMLKNearestNeighborsClassifier.
var _ IMLKNearestNeighborsClassifier = MLKNearestNeighborsClassifier{}

// An interface definition for the [MLKNearestNeighborsClassifier] class.
//
// # Methods
//
//   - [IMLKNearestNeighborsClassifier.CancelUpdate]
//   - [IMLKNearestNeighborsClassifier.ComputeClassProbabilitiesFrom]
//   - [IMLKNearestNeighborsClassifier.ComputeKClosestLabels]
//   - [IMLKNearestNeighborsClassifier.ContinueWithUpdate]
//   - [IMLKNearestNeighborsClassifier.SetContinueWithUpdate]
//   - [IMLKNearestNeighborsClassifier.DefaultLabel]
//   - [IMLKNearestNeighborsClassifier.SetDefaultLabel]
//   - [IMLKNearestNeighborsClassifier.ExtractNearestNeighborLabelsDistancesIndicesFrom]
//   - [IMLKNearestNeighborsClassifier.Index]
//   - [IMLKNearestNeighborsClassifier.SetIndex]
//   - [IMLKNearestNeighborsClassifier.IndexType]
//   - [IMLKNearestNeighborsClassifier.SetIndexType]
//   - [IMLKNearestNeighborsClassifier.InputMultiArrayError]
//   - [IMLKNearestNeighborsClassifier.LabelType]
//   - [IMLKNearestNeighborsClassifier.SetLabelType]
//   - [IMLKNearestNeighborsClassifier.LabelsForDataPoints]
//   - [IMLKNearestNeighborsClassifier.SetLabelsForDataPoints]
//   - [IMLKNearestNeighborsClassifier.LabelsSet]
//   - [IMLKNearestNeighborsClassifier.SetLabelsSet]
//   - [IMLKNearestNeighborsClassifier.NearestDistancesFeatureName]
//   - [IMLKNearestNeighborsClassifier.SetNearestDistancesFeatureName]
//   - [IMLKNearestNeighborsClassifier.NearestLabelsFeatureName]
//   - [IMLKNearestNeighborsClassifier.SetNearestLabelsFeatureName]
//   - [IMLKNearestNeighborsClassifier.NumberOfDimensions]
//   - [IMLKNearestNeighborsClassifier.SetNumberOfDimensions]
//   - [IMLKNearestNeighborsClassifier.PackageOutputWithPredictedLabelClassProbabilitiesNearestLabelsNearestDistancesNearestFeatureIndices]
//   - [IMLKNearestNeighborsClassifier.ParameterContainer]
//   - [IMLKNearestNeighborsClassifier.SetParameterContainer]
//   - [IMLKNearestNeighborsClassifier.ParameterValueForKeyError]
//   - [IMLKNearestNeighborsClassifier.PredictionFromFeaturesOptionsError]
//   - [IMLKNearestNeighborsClassifier.ProgressHandlers]
//   - [IMLKNearestNeighborsClassifier.SetProgressHandlers]
//   - [IMLKNearestNeighborsClassifier.ProgressHandlersDispatchQueue]
//   - [IMLKNearestNeighborsClassifier.SetProgressHandlersDispatchQueue]
//   - [IMLKNearestNeighborsClassifier.ResumeUpdate]
//   - [IMLKNearestNeighborsClassifier.ResumeUpdateWithParameters]
//   - [IMLKNearestNeighborsClassifier.SetUpdateProgressHandlersDispatchQueue]
//   - [IMLKNearestNeighborsClassifier.UpdateModelWithData]
//   - [IMLKNearestNeighborsClassifier.UpdateParameters]
//   - [IMLKNearestNeighborsClassifier.SetUpdateParameters]
//   - [IMLKNearestNeighborsClassifier.WeightingScheme]
//   - [IMLKNearestNeighborsClassifier.SetWeightingScheme]
//   - [IMLKNearestNeighborsClassifier.WriteToURLError]
//   - [IMLKNearestNeighborsClassifier.InitWithDescriptionConfigurationParametersDataPointsLabelsError]
//   - [IMLKNearestNeighborsClassifier.Metadata]
//
// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier
type IMLKNearestNeighborsClassifier interface {
	IMLModel

	// Topic: Methods

	CancelUpdate()
	ComputeClassProbabilitiesFrom(probabilities []objectivec.IObject, from unsafe.Pointer)
	ComputeKClosestLabels(labels objectivec.IObject) objectivec.IObject
	ContinueWithUpdate() bool
	SetContinueWithUpdate(value bool)
	DefaultLabel() objectivec.Object
	SetDefaultLabel(value objectivec.Object)
	ExtractNearestNeighborLabelsDistancesIndicesFrom(labels []objectivec.IObject, distances []objectivec.IObject, indices []objectivec.IObject, from objectivec.IObject)
	Index() objectivec.IObject
	SetIndex(value objectivec.IObject)
	IndexType() int64
	SetIndexType(value int64)
	InputMultiArrayError(array objectivec.IObject) (objectivec.IObject, error)
	LabelType() int64
	SetLabelType(value int64)
	LabelsForDataPoints() foundation.INSArray
	SetLabelsForDataPoints(value foundation.INSArray)
	LabelsSet() foundation.INSOrderedSet
	SetLabelsSet(value foundation.INSOrderedSet)
	NearestDistancesFeatureName() string
	SetNearestDistancesFeatureName(value string)
	NearestLabelsFeatureName() string
	SetNearestLabelsFeatureName(value string)
	NumberOfDimensions() uint64
	SetNumberOfDimensions(value uint64)
	PackageOutputWithPredictedLabelClassProbabilitiesNearestLabelsNearestDistancesNearestFeatureIndices(label objectivec.IObject, probabilities objectivec.IObject, labels objectivec.IObject, distances objectivec.IObject, indices objectivec.IObject) objectivec.IObject
	ParameterContainer() IMLParameterContainer
	SetParameterContainer(value IMLParameterContainer)
	ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error)
	PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error)
	ProgressHandlers() IMLUpdateProgressHandlers
	SetProgressHandlers(value IMLUpdateProgressHandlers)
	ProgressHandlersDispatchQueue() objectivec.Object
	SetProgressHandlersDispatchQueue(value objectivec.Object)
	ResumeUpdate()
	ResumeUpdateWithParameters(parameters objectivec.IObject)
	SetUpdateProgressHandlersDispatchQueue(handlers ErrorHandler, queue objectivec.IObject)
	UpdateModelWithData(data objectivec.IObject)
	UpdateParameters() foundation.INSDictionary
	SetUpdateParameters(value foundation.INSDictionary)
	WeightingScheme() int64
	SetWeightingScheme(value int64)
	WriteToURLError(url foundation.INSURL) (bool, error)
	InitWithDescriptionConfigurationParametersDataPointsLabelsError(description objectivec.IObject, configuration objectivec.IObject, parameters objectivec.IObject, points unsafe.Pointer, labels objectivec.IObject) (MLKNearestNeighborsClassifier, error)
	Metadata() IMLModelMetadata
}

// Init initializes the instance.
func (k MLKNearestNeighborsClassifier) Init() MLKNearestNeighborsClassifier {
	rv := objc.Send[MLKNearestNeighborsClassifier](k.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k MLKNearestNeighborsClassifier) Autorelease() MLKNearestNeighborsClassifier {
	rv := objc.Send[MLKNearestNeighborsClassifier](k.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLKNearestNeighborsClassifier creates a new MLKNearestNeighborsClassifier instance.
func NewMLKNearestNeighborsClassifier() MLKNearestNeighborsClassifier {
	class := getMLKNearestNeighborsClassifierClass()
	rv := objc.Send[MLKNearestNeighborsClassifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initDescriptionOnlyWithSpecification:configuration:error:
func NewKNearestNeighborsClassifierDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLKNearestNeighborsClassifier, error) {
	var errorPtr objc.ID
	instance := getMLKNearestNeighborsClassifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLKNearestNeighborsClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLKNearestNeighborsClassifierFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initInterfaceAndMetadataWithCompiledArchive:error:
func NewKNearestNeighborsClassifierInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLKNearestNeighborsClassifier, error) {
	var errorPtr objc.ID
	instance := getMLKNearestNeighborsClassifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLKNearestNeighborsClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLKNearestNeighborsClassifierFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/initWithCoder:
func NewKNearestNeighborsClassifierWithCoder(coder objectivec.IObject) MLKNearestNeighborsClassifier {
	instance := getMLKNearestNeighborsClassifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLKNearestNeighborsClassifierFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithConfiguration:
func NewKNearestNeighborsClassifierWithConfiguration(configuration objectivec.IObject) MLKNearestNeighborsClassifier {
	instance := getMLKNearestNeighborsClassifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLKNearestNeighborsClassifierFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:
func NewKNearestNeighborsClassifierWithDescription(description objectivec.IObject) MLKNearestNeighborsClassifier {
	instance := getMLKNearestNeighborsClassifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLKNearestNeighborsClassifierFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithDescription:configuration:
func NewKNearestNeighborsClassifierWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLKNearestNeighborsClassifier {
	instance := getMLKNearestNeighborsClassifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLKNearestNeighborsClassifierFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/initWithDescription:configuration:parameters:dataPoints:labels:error:
func NewKNearestNeighborsClassifierWithDescriptionConfigurationParametersDataPointsLabelsError(description objectivec.IObject, configuration objectivec.IObject, parameters objectivec.IObject, points unsafe.Pointer, labels objectivec.IObject) (MLKNearestNeighborsClassifier, error) {
	var errorPtr objc.ID
	instance := getMLKNearestNeighborsClassifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:parameters:dataPoints:labels:error:"), description, configuration, parameters, points, labels, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLKNearestNeighborsClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLKNearestNeighborsClassifierFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModel/initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:
func NewKNearestNeighborsClassifierWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLKNearestNeighborsClassifier {
	instance := getMLKNearestNeighborsClassifierClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLKNearestNeighborsClassifierFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/cancelUpdate
func (k MLKNearestNeighborsClassifier) CancelUpdate() {
	objc.Send[objc.ID](k.ID, objc.Sel("cancelUpdate"))
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/computeClassProbabilities:from:
func (k MLKNearestNeighborsClassifier) ComputeClassProbabilitiesFrom(probabilities []objectivec.IObject, from unsafe.Pointer) {
	objc.Send[objc.ID](k.ID, objc.Sel("computeClassProbabilities:from:"), objectivec.IObjectSliceToNSArray(probabilities), from)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/computeKClosestLabels:
func (k MLKNearestNeighborsClassifier) ComputeKClosestLabels(labels objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("computeKClosestLabels:"), labels)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/extractNearestNeighborLabels:distances:indices:from:
func (k MLKNearestNeighborsClassifier) ExtractNearestNeighborLabelsDistancesIndicesFrom(labels []objectivec.IObject, distances []objectivec.IObject, indices []objectivec.IObject, from objectivec.IObject) {
	objc.Send[objc.ID](k.ID, objc.Sel("extractNearestNeighborLabels:distances:indices:from:"), objectivec.IObjectSliceToNSArray(labels), objectivec.IObjectSliceToNSArray(distances), objectivec.IObjectSliceToNSArray(indices), from)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/inputMultiArray:error:
func (k MLKNearestNeighborsClassifier) InputMultiArrayError(array objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](k.ID, objc.Sel("inputMultiArray:error:"), array, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/packageOutputWithPredictedLabel:classProbabilities:nearestLabels:nearestDistances:nearestFeatureIndices:
func (k MLKNearestNeighborsClassifier) PackageOutputWithPredictedLabelClassProbabilitiesNearestLabelsNearestDistancesNearestFeatureIndices(label objectivec.IObject, probabilities objectivec.IObject, labels objectivec.IObject, distances objectivec.IObject, indices objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("packageOutputWithPredictedLabel:classProbabilities:nearestLabels:nearestDistances:nearestFeatureIndices:"), label, probabilities, labels, distances, indices)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/parameterValueForKey:error:
func (k MLKNearestNeighborsClassifier) ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](k.ID, objc.Sel("parameterValueForKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/predictionFromFeatures:options:error:
func (k MLKNearestNeighborsClassifier) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](k.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/resumeUpdate
func (k MLKNearestNeighborsClassifier) ResumeUpdate() {
	objc.Send[objc.ID](k.ID, objc.Sel("resumeUpdate"))
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/resumeUpdateWithParameters:
func (k MLKNearestNeighborsClassifier) ResumeUpdateWithParameters(parameters objectivec.IObject) {
	objc.Send[objc.ID](k.ID, objc.Sel("resumeUpdateWithParameters:"), parameters)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/setUpdateProgressHandlers:dispatchQueue:
func (k MLKNearestNeighborsClassifier) SetUpdateProgressHandlersDispatchQueue(handlers ErrorHandler, queue objectivec.IObject) {
	_block0, _ := NewErrorBlock(handlers)
	objc.Send[objc.ID](k.ID, objc.Sel("setUpdateProgressHandlers:dispatchQueue:"), _block0, queue)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/updateModelWithData:
func (k MLKNearestNeighborsClassifier) UpdateModelWithData(data objectivec.IObject) {
	objc.Send[objc.ID](k.ID, objc.Sel("updateModelWithData:"), data)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/writeToURL:error:
func (k MLKNearestNeighborsClassifier) WriteToURLError(url foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](k.ID, objc.Sel("writeToURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeToURL:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/initWithDescription:configuration:parameters:dataPoints:labels:error:
func (k MLKNearestNeighborsClassifier) InitWithDescriptionConfigurationParametersDataPointsLabelsError(description objectivec.IObject, configuration objectivec.IObject, parameters objectivec.IObject, points unsafe.Pointer, labels objectivec.IObject) (MLKNearestNeighborsClassifier, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](k.ID, objc.Sel("initWithDescription:configuration:parameters:dataPoints:labels:error:"), description, configuration, parameters, points, labels, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLKNearestNeighborsClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLKNearestNeighborsClassifierFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/compileSpecification:toArchive:options:error:
func (_MLKNearestNeighborsClassifierClass MLKNearestNeighborsClassifierClass) CompileSpecificationToArchiveOptionsError(specification unsafe.Pointer, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLKNearestNeighborsClassifierClass.class), objc.Sel("compileSpecification:toArchive:options:error:"), specification, archive, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/compiledVersionForSpecification:options:error:
func (_MLKNearestNeighborsClassifierClass MLKNearestNeighborsClassifierClass) CompiledVersionForSpecificationOptionsError(specification unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLKNearestNeighborsClassifierClass.class), objc.Sel("compiledVersionForSpecification:options:error:"), specification, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:
func (_MLKNearestNeighborsClassifierClass MLKNearestNeighborsClassifierClass) LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLKNearestNeighborsClassifierClass.class), objc.Sel("loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/supportsSecureCoding
func (_MLKNearestNeighborsClassifierClass MLKNearestNeighborsClassifierClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLKNearestNeighborsClassifierClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/continueWithUpdate
func (k MLKNearestNeighborsClassifier) ContinueWithUpdate() bool {
	rv := objc.Send[bool](k.ID, objc.Sel("continueWithUpdate"))
	return rv
}
func (k MLKNearestNeighborsClassifier) SetContinueWithUpdate(value bool) {
	objc.Send[struct{}](k.ID, objc.Sel("setContinueWithUpdate:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/defaultLabel
func (k MLKNearestNeighborsClassifier) DefaultLabel() objectivec.Object {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("defaultLabel"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (k MLKNearestNeighborsClassifier) SetDefaultLabel(value objectivec.Object) {
	objc.Send[struct{}](k.ID, objc.Sel("setDefaultLabel:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/index
func (k MLKNearestNeighborsClassifier) Index() objectivec.IObject {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("index"))
	return objectivec.Object{ID: rv}
}
func (k MLKNearestNeighborsClassifier) SetIndex(value objectivec.IObject) {
	objc.Send[struct{}](k.ID, objc.Sel("setIndex:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/indexType
func (k MLKNearestNeighborsClassifier) IndexType() int64 {
	rv := objc.Send[int64](k.ID, objc.Sel("indexType"))
	return rv
}
func (k MLKNearestNeighborsClassifier) SetIndexType(value int64) {
	objc.Send[struct{}](k.ID, objc.Sel("setIndexType:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/labelType
func (k MLKNearestNeighborsClassifier) LabelType() int64 {
	rv := objc.Send[int64](k.ID, objc.Sel("labelType"))
	return rv
}
func (k MLKNearestNeighborsClassifier) SetLabelType(value int64) {
	objc.Send[struct{}](k.ID, objc.Sel("setLabelType:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/labelsForDataPoints
func (k MLKNearestNeighborsClassifier) LabelsForDataPoints() foundation.INSArray {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("labelsForDataPoints"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (k MLKNearestNeighborsClassifier) SetLabelsForDataPoints(value foundation.INSArray) {
	objc.Send[struct{}](k.ID, objc.Sel("setLabelsForDataPoints:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/labelsSet
func (k MLKNearestNeighborsClassifier) LabelsSet() foundation.INSOrderedSet {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("labelsSet"))
	return foundation.NSOrderedSetFromID(objc.ID(rv))
}
func (k MLKNearestNeighborsClassifier) SetLabelsSet(value foundation.INSOrderedSet) {
	objc.Send[struct{}](k.ID, objc.Sel("setLabelsSet:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/metadata
func (k MLKNearestNeighborsClassifier) Metadata() IMLModelMetadata {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("metadata"))
	return MLModelMetadataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/nearestDistancesFeatureName
func (k MLKNearestNeighborsClassifier) NearestDistancesFeatureName() string {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("nearestDistancesFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (k MLKNearestNeighborsClassifier) SetNearestDistancesFeatureName(value string) {
	objc.Send[struct{}](k.ID, objc.Sel("setNearestDistancesFeatureName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/nearestLabelsFeatureName
func (k MLKNearestNeighborsClassifier) NearestLabelsFeatureName() string {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("nearestLabelsFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (k MLKNearestNeighborsClassifier) SetNearestLabelsFeatureName(value string) {
	objc.Send[struct{}](k.ID, objc.Sel("setNearestLabelsFeatureName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/numberOfDimensions
func (k MLKNearestNeighborsClassifier) NumberOfDimensions() uint64 {
	rv := objc.Send[uint64](k.ID, objc.Sel("numberOfDimensions"))
	return rv
}
func (k MLKNearestNeighborsClassifier) SetNumberOfDimensions(value uint64) {
	objc.Send[struct{}](k.ID, objc.Sel("setNumberOfDimensions:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/parameterContainer
func (k MLKNearestNeighborsClassifier) ParameterContainer() IMLParameterContainer {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("parameterContainer"))
	return MLParameterContainerFromID(objc.ID(rv))
}
func (k MLKNearestNeighborsClassifier) SetParameterContainer(value IMLParameterContainer) {
	objc.Send[struct{}](k.ID, objc.Sel("setParameterContainer:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/progressHandlers
func (k MLKNearestNeighborsClassifier) ProgressHandlers() IMLUpdateProgressHandlers {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("progressHandlers"))
	return MLUpdateProgressHandlersFromID(objc.ID(rv))
}
func (k MLKNearestNeighborsClassifier) SetProgressHandlers(value IMLUpdateProgressHandlers) {
	objc.Send[struct{}](k.ID, objc.Sel("setProgressHandlers:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/progressHandlersDispatchQueue
func (k MLKNearestNeighborsClassifier) ProgressHandlersDispatchQueue() objectivec.Object {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("progressHandlersDispatchQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (k MLKNearestNeighborsClassifier) SetProgressHandlersDispatchQueue(value objectivec.Object) {
	objc.Send[struct{}](k.ID, objc.Sel("setProgressHandlersDispatchQueue:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/updateParameters
func (k MLKNearestNeighborsClassifier) UpdateParameters() foundation.INSDictionary {
	rv := objc.Send[objc.ID](k.ID, objc.Sel("updateParameters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (k MLKNearestNeighborsClassifier) SetUpdateParameters(value foundation.INSDictionary) {
	objc.Send[struct{}](k.ID, objc.Sel("setUpdateParameters:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLKNearestNeighborsClassifier/weightingScheme
func (k MLKNearestNeighborsClassifier) WeightingScheme() int64 {
	rv := objc.Send[int64](k.ID, objc.Sel("weightingScheme"))
	return rv
}
func (k MLKNearestNeighborsClassifier) SetWeightingScheme(value int64) {
	objc.Send[struct{}](k.ID, objc.Sel("setWeightingScheme:"), value)
}
