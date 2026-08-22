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

// The class instance for the [MLModelDescription] class.
var (
	_MLModelDescriptionClass     MLModelDescriptionClass
	_MLModelDescriptionClassOnce sync.Once
)

func getMLModelDescriptionClass() MLModelDescriptionClass {
	_MLModelDescriptionClassOnce.Do(func() {
		_MLModelDescriptionClass = MLModelDescriptionClass{class: objc.GetClass("MLModelDescription")}
	})
	return _MLModelDescriptionClass
}

// GetMLModelDescriptionClass returns the class object for MLModelDescription.
func GetMLModelDescriptionClass() MLModelDescriptionClass {
	return getMLModelDescriptionClass()
}

type MLModelDescriptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLModelDescriptionClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLModelDescriptionClass) Alloc() MLModelDescription {
	rv := objc.SendIfResponds[MLModelDescription](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLModelDescription.ClassProbabilityFeatureDescription]
//   - [MLModelDescription.DebugQuickLookObject]
//   - [MLModelDescription.DefaultFunctionName]
//   - [MLModelDescription.FunctionDescriptions]
//   - [MLModelDescription.FunctionName]
//   - [MLModelDescription.HasEnumeratedShapeInputs]
//   - [MLModelDescription.HasRangeShapeInputs]
//   - [MLModelDescription.InputFeatureNames]
//   - [MLModelDescription.IsEqualToDescription]
//   - [MLModelDescription.ModelDescriptionBySettingMetadata]
//   - [MLModelDescription.ModelPath]
//   - [MLModelDescription.SetModelPath]
//   - [MLModelDescription.ModelURL]
//   - [MLModelDescription.SetModelURL]
//   - [MLModelDescription.OutputFeatureNames]
//   - [MLModelDescription.PredictedClassFeatureDescription]
//   - [MLModelDescription.PredictedValueFeatureDescription]
//   - [MLModelDescription.StateFeatureNames]
//   - [MLModelDescription.ValidateAsClassifierDescriptionAndReturnError]
//   - [MLModelDescription.ValidateAsRegressorDescriptionAndReturnError]
//   - [MLModelDescription.VerifyInputError]
//   - [MLModelDescription.InitFromModelDescriptionSpecification]
//   - [MLModelDescription.InitFromRawCompiledModelArchiveError]
//   - [MLModelDescription.InitFromRawModelDescriptionSpecification]
//   - [MLModelDescription.InitFromRawModelSpecification]
//   - [MLModelDescription.InitFromSingleFunctionCompiledModelArchiveError]
//   - [MLModelDescription.InitFromSingleFunctionModelDescriptionSpecification]
//   - [MLModelDescription.InitFromSingleFunctionModelSpecification]
//   - [MLModelDescription.InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionDescriptionsIsUpdatableTrainingInputDescriptionsParameterDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesMetadataDefaultFunctionNameFunctionNameClassLabelsModelURLModelPath]
//   - [MLModelDescription.InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameMetadata]
//   - [MLModelDescription.InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameTrainingInputDescriptionsMetadata]
//   - [MLModelDescription.InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameTrainingInputDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesMetadata]
//   - [MLModelDescription.InitWithInputDescriptionsOutputDescriptionsStateDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionDescriptionsIsUpdatableTrainingInputDescriptionsParameterDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesOrderedStateFeatureNamesMetadataDefaultFunctionNameFunctionNameClassLabelsModelURLModelPath]
//   - [MLModelDescription.InitWithInputDescriptionsOutputDescriptionsStateDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionName]
//   - [MLModelDescription.InitWithModelDescriptionSpecificationError]
//   - [MLModelDescription.InitWithModelSpecificationError]
//   - [MLModelDescription.ClassLabels]
//   - [MLModelDescription.SetClassLabels]
//   - [MLModelDescription.IsUpdatable]
//   - [MLModelDescription.SetIsUpdatable]
//   - [MLModelDescription.ParameterDescriptionsByKey]
//   - [MLModelDescription.SetParameterDescriptionsByKey]
//   - [MLModelDescription.TrainingInputDescriptionsByName]
//   - [MLModelDescription.SetTrainingInputDescriptionsByName]
type MLModelDescription struct {
	objectivec.Object
}

// MLModelDescriptionFromID constructs a [MLModelDescription] from an objc.ID.
func MLModelDescriptionFromID(id objc.ID) MLModelDescription {
	return MLModelDescription{objectivec.Object{ID: id}}
}

// Ensure MLModelDescription implements IMLModelDescription.
var _ IMLModelDescription = MLModelDescription{}

// An interface definition for the [MLModelDescription] class.
//
// # Methods
//
//   - [IMLModelDescription.ClassProbabilityFeatureDescription]
//   - [IMLModelDescription.DebugQuickLookObject]
//   - [IMLModelDescription.DefaultFunctionName]
//   - [IMLModelDescription.FunctionDescriptions]
//   - [IMLModelDescription.FunctionName]
//   - [IMLModelDescription.HasEnumeratedShapeInputs]
//   - [IMLModelDescription.HasRangeShapeInputs]
//   - [IMLModelDescription.InputFeatureNames]
//   - [IMLModelDescription.IsEqualToDescription]
//   - [IMLModelDescription.ModelDescriptionBySettingMetadata]
//   - [IMLModelDescription.ModelPath]
//   - [IMLModelDescription.SetModelPath]
//   - [IMLModelDescription.ModelURL]
//   - [IMLModelDescription.SetModelURL]
//   - [IMLModelDescription.OutputFeatureNames]
//   - [IMLModelDescription.PredictedClassFeatureDescription]
//   - [IMLModelDescription.PredictedValueFeatureDescription]
//   - [IMLModelDescription.StateFeatureNames]
//   - [IMLModelDescription.ValidateAsClassifierDescriptionAndReturnError]
//   - [IMLModelDescription.ValidateAsRegressorDescriptionAndReturnError]
//   - [IMLModelDescription.VerifyInputError]
//   - [IMLModelDescription.InitFromModelDescriptionSpecification]
//   - [IMLModelDescription.InitFromRawCompiledModelArchiveError]
//   - [IMLModelDescription.InitFromRawModelDescriptionSpecification]
//   - [IMLModelDescription.InitFromRawModelSpecification]
//   - [IMLModelDescription.InitFromSingleFunctionCompiledModelArchiveError]
//   - [IMLModelDescription.InitFromSingleFunctionModelDescriptionSpecification]
//   - [IMLModelDescription.InitFromSingleFunctionModelSpecification]
//   - [IMLModelDescription.InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionDescriptionsIsUpdatableTrainingInputDescriptionsParameterDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesMetadataDefaultFunctionNameFunctionNameClassLabelsModelURLModelPath]
//   - [IMLModelDescription.InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameMetadata]
//   - [IMLModelDescription.InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameTrainingInputDescriptionsMetadata]
//   - [IMLModelDescription.InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameTrainingInputDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesMetadata]
//   - [IMLModelDescription.InitWithInputDescriptionsOutputDescriptionsStateDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionDescriptionsIsUpdatableTrainingInputDescriptionsParameterDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesOrderedStateFeatureNamesMetadataDefaultFunctionNameFunctionNameClassLabelsModelURLModelPath]
//   - [IMLModelDescription.InitWithInputDescriptionsOutputDescriptionsStateDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionName]
//   - [IMLModelDescription.InitWithModelDescriptionSpecificationError]
//   - [IMLModelDescription.InitWithModelSpecificationError]
//   - [IMLModelDescription.ClassLabels]
//   - [IMLModelDescription.SetClassLabels]
//   - [IMLModelDescription.IsUpdatable]
//   - [IMLModelDescription.SetIsUpdatable]
//   - [IMLModelDescription.ParameterDescriptionsByKey]
//   - [IMLModelDescription.SetParameterDescriptionsByKey]
//   - [IMLModelDescription.TrainingInputDescriptionsByName]
//   - [IMLModelDescription.SetTrainingInputDescriptionsByName]
type IMLModelDescription interface {
	objectivec.IObject

	// Topic: Methods

	ClassProbabilityFeatureDescription() IMLFeatureDescription
	DebugQuickLookObject() objectivec.IObject
	DefaultFunctionName() objectivec.IObject
	FunctionDescriptions() objectivec.IObject
	FunctionName() string
	HasEnumeratedShapeInputs() bool
	HasRangeShapeInputs() bool
	InputFeatureNames() foundation.INSOrderedSet
	IsEqualToDescription(description objectivec.IObject) bool
	ModelDescriptionBySettingMetadata(metadata objectivec.IObject) objectivec.IObject
	ModelPath() IMLLayerPath
	SetModelPath(value IMLLayerPath)
	ModelURL() foundation.NSURL
	SetModelURL(value foundation.NSURL)
	OutputFeatureNames() foundation.INSOrderedSet
	PredictedClassFeatureDescription() IMLFeatureDescription
	PredictedValueFeatureDescription() IMLFeatureDescription
	StateFeatureNames() foundation.INSOrderedSet
	ValidateAsClassifierDescriptionAndReturnError() (bool, error)
	ValidateAsRegressorDescriptionAndReturnError() (bool, error)
	VerifyInputError(input objectivec.IObject) (bool, error)
	InitFromModelDescriptionSpecification(specification unsafe.Pointer) MLModelDescription
	InitFromRawCompiledModelArchiveError(archive MLModelInputArchiverRef) (MLModelDescription, error)
	InitFromRawModelDescriptionSpecification(specification unsafe.Pointer) MLModelDescription
	InitFromRawModelSpecification(specification unsafe.Pointer) MLModelDescription
	InitFromSingleFunctionCompiledModelArchiveError(archive MLModelInputArchiverRef) (MLModelDescription, error)
	InitFromSingleFunctionModelDescriptionSpecification(specification unsafe.Pointer) MLModelDescription
	InitFromSingleFunctionModelSpecification(specification unsafe.Pointer) MLModelDescription
	InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionDescriptionsIsUpdatableTrainingInputDescriptionsParameterDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesMetadataDefaultFunctionNameFunctionNameClassLabelsModelURLModelPath(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions3 objectivec.IObject, updatable bool, descriptions4 objectivec.IObject, descriptions5 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, metadata objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, labels objectivec.IObject, url foundation.NSURL, path objectivec.IObject) MLModelDescription
	InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameMetadata(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, metadata objectivec.IObject) MLModelDescription
	InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameTrainingInputDescriptionsMetadata(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions3 objectivec.IObject, metadata objectivec.IObject) MLModelDescription
	InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameTrainingInputDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesMetadata(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions3 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, metadata objectivec.IObject) MLModelDescription
	InitWithInputDescriptionsOutputDescriptionsStateDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionDescriptionsIsUpdatableTrainingInputDescriptionsParameterDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesOrderedStateFeatureNamesMetadataDefaultFunctionNameFunctionNameClassLabelsModelURLModelPath(descriptions objectivec.IObject, descriptions2 objectivec.IObject, descriptions3 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions4 objectivec.IObject, updatable bool, descriptions5 objectivec.IObject, descriptions6 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, names3 objectivec.IObject, metadata objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, labels objectivec.IObject, url foundation.NSURL, path objectivec.IObject) MLModelDescription
	InitWithInputDescriptionsOutputDescriptionsStateDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionName(descriptions objectivec.IObject, descriptions2 objectivec.IObject, descriptions3 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject) MLModelDescription
	InitWithModelDescriptionSpecificationError(specification unsafe.Pointer) (MLModelDescription, error)
	InitWithModelSpecificationError(specification unsafe.Pointer) (MLModelDescription, error)
	ClassLabels() foundation.INSArray
	SetClassLabels(value foundation.INSArray)
	IsUpdatable() bool
	SetIsUpdatable(value bool)
	ParameterDescriptionsByKey() foundation.INSDictionary
	SetParameterDescriptionsByKey(value foundation.INSDictionary)
	TrainingInputDescriptionsByName() foundation.INSDictionary
	SetTrainingInputDescriptionsByName(value foundation.INSDictionary)
}

// Init initializes the instance.
func (m MLModelDescription) Init() MLModelDescription {
	rv := objc.SendIfResponds[MLModelDescription](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelDescription) Autorelease() MLModelDescription {
	rv := objc.SendIfResponds[MLModelDescription](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelDescription creates a new MLModelDescription instance.
func NewMLModelDescription() MLModelDescription {
	class := getMLModelDescriptionClass()
	rv := objc.SendIfResponds[MLModelDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewModelDescriptionFromModelDescriptionSpecification(specification unsafe.Pointer) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initFromModelDescriptionSpecification:"), specification)
	return MLModelDescriptionFromID(rv)
}

func NewModelDescriptionFromRawCompiledModelArchiveError(archive MLModelInputArchiverRef) (MLModelDescription, error) {
	var errorPtr objc.ID
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initFromRawCompiledModelArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelDescription{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLModelDescription{}, objc.ErrInitFailed
	}
	return MLModelDescriptionFromID(rv), nil
}

func NewModelDescriptionFromRawModelDescriptionSpecification(specification unsafe.Pointer) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initFromRawModelDescriptionSpecification:"), specification)
	return MLModelDescriptionFromID(rv)
}

func NewModelDescriptionFromRawModelSpecification(specification unsafe.Pointer) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initFromRawModelSpecification:"), specification)
	return MLModelDescriptionFromID(rv)
}

func NewModelDescriptionFromSingleFunctionCompiledModelArchiveError(archive MLModelInputArchiverRef) (MLModelDescription, error) {
	var errorPtr objc.ID
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initFromSingleFunctionCompiledModelArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelDescription{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLModelDescription{}, objc.ErrInitFailed
	}
	return MLModelDescriptionFromID(rv), nil
}

func NewModelDescriptionFromSingleFunctionModelDescriptionSpecification(specification unsafe.Pointer) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initFromSingleFunctionModelDescriptionSpecification:"), specification)
	return MLModelDescriptionFromID(rv)
}

func NewModelDescriptionFromSingleFunctionModelSpecification(specification unsafe.Pointer) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initFromSingleFunctionModelSpecification:"), specification)
	return MLModelDescriptionFromID(rv)
}

func NewModelDescriptionWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionDescriptionsIsUpdatableTrainingInputDescriptionsParameterDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesMetadataDefaultFunctionNameFunctionNameClassLabelsModelURLModelPath(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions3 objectivec.IObject, updatable bool, descriptions4 objectivec.IObject, descriptions5 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, metadata objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, labels objectivec.IObject, url foundation.NSURL, path objectivec.IObject) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:functionDescriptions:isUpdatable:trainingInputDescriptions:parameterDescriptions:orderedInputFeatureNames:orderedOutputFeatureNames:metadata:defaultFunctionName:functionName:classLabels:modelURL:modelPath:"), descriptions, descriptions2, name, name2, descriptions3, updatable, descriptions4, descriptions5, names, names2, metadata, name3, name4, labels, url, path)
	return MLModelDescriptionFromID(rv)
}

func NewModelDescriptionWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameMetadata(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, metadata objectivec.IObject) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:metadata:"), descriptions, descriptions2, name, name2, metadata)
	return MLModelDescriptionFromID(rv)
}

func NewModelDescriptionWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameTrainingInputDescriptionsMetadata(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions3 objectivec.IObject, metadata objectivec.IObject) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:trainingInputDescriptions:metadata:"), descriptions, descriptions2, name, name2, descriptions3, metadata)
	return MLModelDescriptionFromID(rv)
}

func NewModelDescriptionWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameTrainingInputDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesMetadata(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions3 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, metadata objectivec.IObject) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:trainingInputDescriptions:orderedInputFeatureNames:orderedOutputFeatureNames:metadata:"), descriptions, descriptions2, name, name2, descriptions3, names, names2, metadata)
	return MLModelDescriptionFromID(rv)
}

func NewModelDescriptionWithInputDescriptionsOutputDescriptionsStateDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionDescriptionsIsUpdatableTrainingInputDescriptionsParameterDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesOrderedStateFeatureNamesMetadataDefaultFunctionNameFunctionNameClassLabelsModelURLModelPath(descriptions objectivec.IObject, descriptions2 objectivec.IObject, descriptions3 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions4 objectivec.IObject, updatable bool, descriptions5 objectivec.IObject, descriptions6 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, names3 objectivec.IObject, metadata objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, labels objectivec.IObject, url foundation.NSURL, path objectivec.IObject) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:stateDescriptions:predictedFeatureName:predictedProbabilitiesName:functionDescriptions:isUpdatable:trainingInputDescriptions:parameterDescriptions:orderedInputFeatureNames:orderedOutputFeatureNames:orderedStateFeatureNames:metadata:defaultFunctionName:functionName:classLabels:modelURL:modelPath:"), descriptions, descriptions2, descriptions3, name, name2, descriptions4, updatable, descriptions5, descriptions6, names, names2, names3, metadata, name3, name4, labels, url, path)
	return MLModelDescriptionFromID(rv)
}

func NewModelDescriptionWithInputDescriptionsOutputDescriptionsStateDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionName(descriptions objectivec.IObject, descriptions2 objectivec.IObject, descriptions3 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:stateDescriptions:predictedFeatureName:predictedProbabilitiesName:functionName:"), descriptions, descriptions2, descriptions3, name, name2, name3)
	return MLModelDescriptionFromID(rv)
}

func NewModelDescriptionWithModelDescriptionSpecificationError(specification unsafe.Pointer) (MLModelDescription, error) {
	var errorPtr objc.ID
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithModelDescriptionSpecification:error:"), specification, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelDescription{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLModelDescription{}, objc.ErrInitFailed
	}
	return MLModelDescriptionFromID(rv), nil
}

func NewModelDescriptionWithModelSpecificationError(specification unsafe.Pointer) (MLModelDescription, error) {
	var errorPtr objc.ID
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithModelSpecification:error:"), specification, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelDescription{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLModelDescription{}, objc.ErrInitFailed
	}
	return MLModelDescriptionFromID(rv), nil
}

func (m MLModelDescription) DebugQuickLookObject() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugQuickLookObject"))
	return objectivec.Object{ID: rv}
}
func (m MLModelDescription) DefaultFunctionName() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("defaultFunctionName"))
	return objectivec.Object{ID: rv}
}
func (m MLModelDescription) FunctionDescriptions() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("functionDescriptions"))
	return objectivec.Object{ID: rv}
}
func (m MLModelDescription) HasEnumeratedShapeInputs() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasEnumeratedShapeInputs"))
	return rv
}
func (m MLModelDescription) HasRangeShapeInputs() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("hasRangeShapeInputs"))
	return rv
}
func (m MLModelDescription) IsEqualToDescription(description objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("isEqualToDescription:"), description)
	return rv
}
func (m MLModelDescription) ModelDescriptionBySettingMetadata(metadata objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelDescriptionBySettingMetadata:"), metadata)
	return objectivec.Object{ID: rv}
}
func (m MLModelDescription) ValidateAsClassifierDescriptionAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("validateAsClassifierDescriptionAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateAsClassifierDescriptionAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLModelDescription) ValidateAsRegressorDescriptionAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("validateAsRegressorDescriptionAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("validateAsRegressorDescriptionAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLModelDescription) VerifyInputError(input objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("verifyInput:error:"), input, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("verifyInput:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (m MLModelDescription) InitFromModelDescriptionSpecification(specification unsafe.Pointer) MLModelDescription {
	rv := objc.SendIfResponds[MLModelDescription](m.ID, objc.Sel("initFromModelDescriptionSpecification:"), specification)
	return rv
}
func (m MLModelDescription) InitFromRawCompiledModelArchiveError(archive MLModelInputArchiverRef) (MLModelDescription, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initFromRawCompiledModelArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelDescription{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelDescriptionFromID(rv), nil

}
func (m MLModelDescription) InitFromRawModelDescriptionSpecification(specification unsafe.Pointer) MLModelDescription {
	rv := objc.SendIfResponds[MLModelDescription](m.ID, objc.Sel("initFromRawModelDescriptionSpecification:"), specification)
	return rv
}
func (m MLModelDescription) InitFromRawModelSpecification(specification unsafe.Pointer) MLModelDescription {
	rv := objc.SendIfResponds[MLModelDescription](m.ID, objc.Sel("initFromRawModelSpecification:"), specification)
	return rv
}
func (m MLModelDescription) InitFromSingleFunctionCompiledModelArchiveError(archive MLModelInputArchiverRef) (MLModelDescription, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initFromSingleFunctionCompiledModelArchive:error:"), archive, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelDescription{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelDescriptionFromID(rv), nil

}
func (m MLModelDescription) InitFromSingleFunctionModelDescriptionSpecification(specification unsafe.Pointer) MLModelDescription {
	rv := objc.SendIfResponds[MLModelDescription](m.ID, objc.Sel("initFromSingleFunctionModelDescriptionSpecification:"), specification)
	return rv
}
func (m MLModelDescription) InitFromSingleFunctionModelSpecification(specification unsafe.Pointer) MLModelDescription {
	rv := objc.SendIfResponds[MLModelDescription](m.ID, objc.Sel("initFromSingleFunctionModelSpecification:"), specification)
	return rv
}
func (m MLModelDescription) InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionDescriptionsIsUpdatableTrainingInputDescriptionsParameterDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesMetadataDefaultFunctionNameFunctionNameClassLabelsModelURLModelPath(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions3 objectivec.IObject, updatable bool, descriptions4 objectivec.IObject, descriptions5 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, metadata objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, labels objectivec.IObject, url foundation.NSURL, path objectivec.IObject) MLModelDescription {
	rv := objc.SendIfResponds[MLModelDescription](m.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:functionDescriptions:isUpdatable:trainingInputDescriptions:parameterDescriptions:orderedInputFeatureNames:orderedOutputFeatureNames:metadata:defaultFunctionName:functionName:classLabels:modelURL:modelPath:"), descriptions, descriptions2, name, name2, descriptions3, updatable, descriptions4, descriptions5, names, names2, metadata, name3, name4, labels, url, path)
	return rv
}
func (m MLModelDescription) InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameMetadata(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, metadata objectivec.IObject) MLModelDescription {
	rv := objc.SendIfResponds[MLModelDescription](m.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:metadata:"), descriptions, descriptions2, name, name2, metadata)
	return rv
}
func (m MLModelDescription) InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameTrainingInputDescriptionsMetadata(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions3 objectivec.IObject, metadata objectivec.IObject) MLModelDescription {
	rv := objc.SendIfResponds[MLModelDescription](m.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:trainingInputDescriptions:metadata:"), descriptions, descriptions2, name, name2, descriptions3, metadata)
	return rv
}
func (m MLModelDescription) InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameTrainingInputDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesMetadata(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions3 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, metadata objectivec.IObject) MLModelDescription {
	rv := objc.SendIfResponds[MLModelDescription](m.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:trainingInputDescriptions:orderedInputFeatureNames:orderedOutputFeatureNames:metadata:"), descriptions, descriptions2, name, name2, descriptions3, names, names2, metadata)
	return rv
}
func (m MLModelDescription) InitWithInputDescriptionsOutputDescriptionsStateDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionDescriptionsIsUpdatableTrainingInputDescriptionsParameterDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesOrderedStateFeatureNamesMetadataDefaultFunctionNameFunctionNameClassLabelsModelURLModelPath(descriptions objectivec.IObject, descriptions2 objectivec.IObject, descriptions3 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions4 objectivec.IObject, updatable bool, descriptions5 objectivec.IObject, descriptions6 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, names3 objectivec.IObject, metadata objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, labels objectivec.IObject, url foundation.NSURL, path objectivec.IObject) MLModelDescription {
	rv := objc.SendIfResponds[MLModelDescription](m.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:stateDescriptions:predictedFeatureName:predictedProbabilitiesName:functionDescriptions:isUpdatable:trainingInputDescriptions:parameterDescriptions:orderedInputFeatureNames:orderedOutputFeatureNames:orderedStateFeatureNames:metadata:defaultFunctionName:functionName:classLabels:modelURL:modelPath:"), descriptions, descriptions2, descriptions3, name, name2, descriptions4, updatable, descriptions5, descriptions6, names, names2, names3, metadata, name3, name4, labels, url, path)
	return rv
}
func (m MLModelDescription) InitWithInputDescriptionsOutputDescriptionsStateDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionName(descriptions objectivec.IObject, descriptions2 objectivec.IObject, descriptions3 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject) MLModelDescription {
	rv := objc.SendIfResponds[MLModelDescription](m.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:stateDescriptions:predictedFeatureName:predictedProbabilitiesName:functionName:"), descriptions, descriptions2, descriptions3, name, name2, name3)
	return rv
}
func (m MLModelDescription) InitWithModelDescriptionSpecificationError(specification unsafe.Pointer) (MLModelDescription, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithModelDescriptionSpecification:error:"), specification, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelDescription{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelDescriptionFromID(rv), nil

}
func (m MLModelDescription) InitWithModelSpecificationError(specification unsafe.Pointer) (MLModelDescription, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithModelSpecification:error:"), specification, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelDescription{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelDescriptionFromID(rv), nil

}

func (_MLModelDescriptionClass MLModelDescriptionClass) MetadataWithFormat(format unsafe.Pointer) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelDescriptionClass.class), objc.Sel("metadataWithFormat:"), format)
	return objectivec.Object{ID: rv}
}
func (_MLModelDescriptionClass MLModelDescriptionClass) MetadataWithSpecification(specification unsafe.Pointer) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_MLModelDescriptionClass.class), objc.Sel("metadataWithSpecification:"), specification)
	return objectivec.Object{ID: rv}
}
func (_MLModelDescriptionClass MLModelDescriptionClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_MLModelDescriptionClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (m MLModelDescription) ClassLabels() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("classLabels"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLModelDescription) SetClassLabels(value foundation.INSArray) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setClassLabels:"), value)
}
func (m MLModelDescription) ClassProbabilityFeatureDescription() IMLFeatureDescription {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("classProbabilityFeatureDescription"))
	return MLFeatureDescriptionFromID(objc.ID(rv))
}
func (m MLModelDescription) FunctionName() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("functionName"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLModelDescription) InputFeatureNames() foundation.INSOrderedSet {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("inputFeatureNames"))
	return foundation.NSOrderedSetFromID(objc.ID(rv))
}
func (m MLModelDescription) IsUpdatable() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("isUpdatable"))
	return rv
}
func (m MLModelDescription) SetIsUpdatable(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setIsUpdatable:"), value)
}
func (m MLModelDescription) ModelPath() IMLLayerPath {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelPath"))
	return MLLayerPathFromID(objc.ID(rv))
}
func (m MLModelDescription) SetModelPath(value IMLLayerPath) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setModelPath:"), value)
}
func (m MLModelDescription) ModelURL() foundation.NSURL {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (m MLModelDescription) SetModelURL(value foundation.NSURL) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setModelURL:"), value)
}
func (m MLModelDescription) OutputFeatureNames() foundation.INSOrderedSet {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("outputFeatureNames"))
	return foundation.NSOrderedSetFromID(objc.ID(rv))
}
func (m MLModelDescription) ParameterDescriptionsByKey() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("parameterDescriptionsByKey"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLModelDescription) SetParameterDescriptionsByKey(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setParameterDescriptionsByKey:"), value)
}
func (m MLModelDescription) PredictedClassFeatureDescription() IMLFeatureDescription {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("predictedClassFeatureDescription"))
	return MLFeatureDescriptionFromID(objc.ID(rv))
}
func (m MLModelDescription) PredictedValueFeatureDescription() IMLFeatureDescription {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("predictedValueFeatureDescription"))
	return MLFeatureDescriptionFromID(objc.ID(rv))
}
func (m MLModelDescription) StateFeatureNames() foundation.INSOrderedSet {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("stateFeatureNames"))
	return foundation.NSOrderedSetFromID(objc.ID(rv))
}
func (m MLModelDescription) TrainingInputDescriptionsByName() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("trainingInputDescriptionsByName"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLModelDescription) SetTrainingInputDescriptionsByName(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setTrainingInputDescriptionsByName:"), value)
}
