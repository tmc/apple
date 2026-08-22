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
	rv := objc.SendIfResponds[MLKNearestNeighborsClassifier](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLKNearestNeighborsClassifier.CancelUpdate]
//   - [MLKNearestNeighborsClassifier.ComputeClassProbabilitiesFrom]
//   - [MLKNearestNeighborsClassifier.ContinueWithUpdate]
//   - [MLKNearestNeighborsClassifier.SetContinueWithUpdate]
//   - [MLKNearestNeighborsClassifier.DefaultLabel]
//   - [MLKNearestNeighborsClassifier.SetDefaultLabel]
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
//   - [IMLKNearestNeighborsClassifier.ContinueWithUpdate]
//   - [IMLKNearestNeighborsClassifier.SetContinueWithUpdate]
//   - [IMLKNearestNeighborsClassifier.DefaultLabel]
//   - [IMLKNearestNeighborsClassifier.SetDefaultLabel]
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
type IMLKNearestNeighborsClassifier interface {
	IMLModel

	// Topic: Methods

	CancelUpdate()
	ComputeClassProbabilitiesFrom(probabilities []objectivec.IObject, from unsafe.Pointer)
	ContinueWithUpdate() bool
	SetContinueWithUpdate(value bool)
	DefaultLabel() objectivec.Object
	SetDefaultLabel(value objectivec.Object)
	Index() unsafe.Pointer
	SetIndex(value unsafe.Pointer)
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
	SetUpdateProgressHandlersDispatchQueue(handlers objectivec.IObject, queue objectivec.IObject)
	UpdateModelWithData(data objectivec.IObject)
	UpdateParameters() foundation.INSDictionary
	SetUpdateParameters(value foundation.INSDictionary)
	WeightingScheme() int64
	SetWeightingScheme(value int64)
	WriteToURLError(url foundation.NSURL) (bool, error)
	InitWithDescriptionConfigurationParametersDataPointsLabelsError(description objectivec.IObject, configuration objectivec.IObject, parameters objectivec.IObject, points unsafe.Pointer, labels objectivec.IObject) (MLKNearestNeighborsClassifier, error)
}

// Init initializes the instance.
func (m MLKNearestNeighborsClassifier) Init() MLKNearestNeighborsClassifier {
	rv := objc.SendIfResponds[MLKNearestNeighborsClassifier](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLKNearestNeighborsClassifier) Autorelease() MLKNearestNeighborsClassifier {
	rv := objc.SendIfResponds[MLKNearestNeighborsClassifier](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLKNearestNeighborsClassifier creates a new MLKNearestNeighborsClassifier instance.
func NewMLKNearestNeighborsClassifier() MLKNearestNeighborsClassifier {
	class := getMLKNearestNeighborsClassifierClass()
	rv := objc.SendIfResponds[MLKNearestNeighborsClassifier](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewKNearestNeighborsClassifierDescriptionOnlyWithSpecificationConfigurationError(specification unsafe.Pointer, configuration objectivec.IObject) (MLKNearestNeighborsClassifier, error) {
	var errorPtr objc.ID
	instance := getMLKNearestNeighborsClassifierClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initDescriptionOnlyWithSpecification:configuration:error:"), specification, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLKNearestNeighborsClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLKNearestNeighborsClassifier{}, objc.ErrInitFailed
	}
	return MLKNearestNeighborsClassifierFromID(rv), nil
}

func NewKNearestNeighborsClassifierInterfaceAndMetadataWithCompiledArchiveError(archive unsafe.Pointer) (MLKNearestNeighborsClassifier, error) {
	var errorPtr objc.ID
	instance := getMLKNearestNeighborsClassifierClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initInterfaceAndMetadataWithCompiledArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLKNearestNeighborsClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLKNearestNeighborsClassifier{}, objc.ErrInitFailed
	}
	return MLKNearestNeighborsClassifierFromID(rv), nil
}

func NewKNearestNeighborsClassifierWithCoder(coder objectivec.IObject) MLKNearestNeighborsClassifier {
	instance := getMLKNearestNeighborsClassifierClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLKNearestNeighborsClassifierFromID(rv)
}

func NewKNearestNeighborsClassifierWithConfiguration(configuration objectivec.IObject) MLKNearestNeighborsClassifier {
	instance := getMLKNearestNeighborsClassifierClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithConfiguration:"), configuration)
	return MLKNearestNeighborsClassifierFromID(rv)
}

func NewKNearestNeighborsClassifierWithDescription(description objectivec.IObject) MLKNearestNeighborsClassifier {
	instance := getMLKNearestNeighborsClassifierClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescription:"), description)
	return MLKNearestNeighborsClassifierFromID(rv)
}

func NewKNearestNeighborsClassifierWithDescriptionConfiguration(description objectivec.IObject, configuration objectivec.IObject) MLKNearestNeighborsClassifier {
	instance := getMLKNearestNeighborsClassifierClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:"), description, configuration)
	return MLKNearestNeighborsClassifierFromID(rv)
}

func NewKNearestNeighborsClassifierWithDescriptionConfigurationParametersDataPointsLabelsError(description objectivec.IObject, configuration objectivec.IObject, parameters objectivec.IObject, points unsafe.Pointer, labels objectivec.IObject) (MLKNearestNeighborsClassifier, error) {
	var errorPtr objc.ID
	instance := getMLKNearestNeighborsClassifierClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:parameters:dataPoints:labels:error:"), description, configuration, parameters, points, labels, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLKNearestNeighborsClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLKNearestNeighborsClassifier{}, objc.ErrInitFailed
	}
	return MLKNearestNeighborsClassifierFromID(rv), nil
}

func NewKNearestNeighborsClassifierWithNameInputDescriptionOutputDescriptionOrderedInputFeatureNamesOrderedOutputFeatureNamesConfiguration(name objectivec.IObject, description objectivec.IObject, description2 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, configuration objectivec.IObject) MLKNearestNeighborsClassifier {
	instance := getMLKNearestNeighborsClassifierClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithName:inputDescription:outputDescription:orderedInputFeatureNames:orderedOutputFeatureNames:configuration:"), name, description, description2, names, names2, configuration)
	return MLKNearestNeighborsClassifierFromID(rv)
}

func (m MLKNearestNeighborsClassifier) CancelUpdate() {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("cancelUpdate"))
}
func (m MLKNearestNeighborsClassifier) ComputeClassProbabilitiesFrom(probabilities []objectivec.IObject, from unsafe.Pointer) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("computeClassProbabilities:from:"), objectivec.IObjectSliceToNSArray(probabilities), from)
}
func (m MLKNearestNeighborsClassifier) InputMultiArrayError(array objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputMultiArray:error:"), array, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLKNearestNeighborsClassifier) PackageOutputWithPredictedLabelClassProbabilitiesNearestLabelsNearestDistancesNearestFeatureIndices(label objectivec.IObject, probabilities objectivec.IObject, labels objectivec.IObject, distances objectivec.IObject, indices objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("packageOutputWithPredictedLabel:classProbabilities:nearestLabels:nearestDistances:nearestFeatureIndices:"), label, probabilities, labels, distances, indices)
	return objectivec.Object{ID: rv}
}
func (m MLKNearestNeighborsClassifier) ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parameterValueForKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLKNearestNeighborsClassifier) PredictionFromFeaturesOptionsError(features objectivec.IObject, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictionFromFeatures:options:error:"), features, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLKNearestNeighborsClassifier) ResumeUpdate() {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("resumeUpdate"))
}
func (m MLKNearestNeighborsClassifier) ResumeUpdateWithParameters(parameters objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("resumeUpdateWithParameters:"), parameters)
}
func (m MLKNearestNeighborsClassifier) SetUpdateProgressHandlersDispatchQueue(handlers objectivec.IObject, queue objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("setUpdateProgressHandlers:dispatchQueue:"), handlers, queue)
}
func (m MLKNearestNeighborsClassifier) UpdateModelWithData(data objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("updateModelWithData:"), data)
}
func (m MLKNearestNeighborsClassifier) WriteToURLError(url foundation.NSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("writeToURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeToURL:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLKNearestNeighborsClassifier) InitWithDescriptionConfigurationParametersDataPointsLabelsError(description objectivec.IObject, configuration objectivec.IObject, parameters objectivec.IObject, points unsafe.Pointer, labels objectivec.IObject) (MLKNearestNeighborsClassifier, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithDescription:configuration:parameters:dataPoints:labels:error:"), description, configuration, parameters, points, labels, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLKNearestNeighborsClassifier{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLKNearestNeighborsClassifierFromID(rv), nil

}

func (_MLKNearestNeighborsClassifierClass MLKNearestNeighborsClassifierClass) CompileSpecificationToArchiveOptionsError(specification unsafe.Pointer, archive unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLKNearestNeighborsClassifierClass.class), objc.Sel("compileSpecification:toArchive:options:error:"), specification, archive, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLKNearestNeighborsClassifierClass MLKNearestNeighborsClassifierClass) CompiledVersionForSpecificationOptionsError(specification unsafe.Pointer, options objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLKNearestNeighborsClassifierClass.class), objc.Sel("compiledVersionForSpecification:options:error:"), specification, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLKNearestNeighborsClassifierClass MLKNearestNeighborsClassifierClass) LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLKNearestNeighborsClassifierClass.class), objc.Sel("loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (_MLKNearestNeighborsClassifierClass MLKNearestNeighborsClassifierClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLKNearestNeighborsClassifierClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (m MLKNearestNeighborsClassifier) ContinueWithUpdate() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("continueWithUpdate"))
	return rv
}
func (m MLKNearestNeighborsClassifier) SetContinueWithUpdate(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setContinueWithUpdate:"), value)
}
func (m MLKNearestNeighborsClassifier) DefaultLabel() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("defaultLabel"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (m MLKNearestNeighborsClassifier) SetDefaultLabel(value objectivec.Object) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setDefaultLabel:"), value)
}
func (m MLKNearestNeighborsClassifier) Index() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("index"))
	return rv
}
func (m MLKNearestNeighborsClassifier) SetIndex(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setIndex:"), value)
}
func (m MLKNearestNeighborsClassifier) IndexType() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("indexType"))
	return rv
}
func (m MLKNearestNeighborsClassifier) SetIndexType(value int64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setIndexType:"), value)
}
func (m MLKNearestNeighborsClassifier) LabelType() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("labelType"))
	return rv
}
func (m MLKNearestNeighborsClassifier) SetLabelType(value int64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setLabelType:"), value)
}
func (m MLKNearestNeighborsClassifier) LabelsForDataPoints() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("labelsForDataPoints"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLKNearestNeighborsClassifier) SetLabelsForDataPoints(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setLabelsForDataPoints:"), value)
}
func (m MLKNearestNeighborsClassifier) LabelsSet() foundation.INSOrderedSet {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("labelsSet"))
	return foundation.NSOrderedSetFromID(objc.ID(rv))
}
func (m MLKNearestNeighborsClassifier) SetLabelsSet(value foundation.INSOrderedSet) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setLabelsSet:"), value)
}
func (m MLKNearestNeighborsClassifier) NearestDistancesFeatureName() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("nearestDistancesFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLKNearestNeighborsClassifier) SetNearestDistancesFeatureName(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setNearestDistancesFeatureName:"), objc.String(value))
}
func (m MLKNearestNeighborsClassifier) NearestLabelsFeatureName() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("nearestLabelsFeatureName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLKNearestNeighborsClassifier) SetNearestLabelsFeatureName(value string) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setNearestLabelsFeatureName:"), objc.String(value))
}
func (m MLKNearestNeighborsClassifier) NumberOfDimensions() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("numberOfDimensions"))
	return rv
}
func (m MLKNearestNeighborsClassifier) SetNumberOfDimensions(value uint64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setNumberOfDimensions:"), value)
}
func (m MLKNearestNeighborsClassifier) ParameterContainer() IMLParameterContainer {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("parameterContainer"))
	return MLParameterContainerFromID(objc.ID(rv))
}
func (m MLKNearestNeighborsClassifier) SetParameterContainer(value IMLParameterContainer) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setParameterContainer:"), value)
}
func (m MLKNearestNeighborsClassifier) ProgressHandlers() IMLUpdateProgressHandlers {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("progressHandlers"))
	return MLUpdateProgressHandlersFromID(objc.ID(rv))
}
func (m MLKNearestNeighborsClassifier) SetProgressHandlers(value IMLUpdateProgressHandlers) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setProgressHandlers:"), value)
}
func (m MLKNearestNeighborsClassifier) ProgressHandlersDispatchQueue() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("progressHandlersDispatchQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (m MLKNearestNeighborsClassifier) SetProgressHandlersDispatchQueue(value objectivec.Object) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setProgressHandlersDispatchQueue:"), value)
}
func (m MLKNearestNeighborsClassifier) UpdateParameters() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("updateParameters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLKNearestNeighborsClassifier) SetUpdateParameters(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setUpdateParameters:"), value)
}
func (m MLKNearestNeighborsClassifier) WeightingScheme() int64 {
	rv := objc.SendIfResponds[int64](m.ID, objc.Sel("weightingScheme"))
	return rv
}
func (m MLKNearestNeighborsClassifier) SetWeightingScheme(value int64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setWeightingScheme:"), value)
}
