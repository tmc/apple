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
	rv := objc.Send[MLModelDescription](objc.ID(mc.class), objc.Sel("alloc"))
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
//   - [MLModelDescription.InitFromRawModelDescriptionSpecification]
//   - [MLModelDescription.InitFromRawModelSpecification]
//   - [MLModelDescription.InitFromSingleFunctionModelDescriptionSpecification]
//   - [MLModelDescription.InitFromSingleFunctionModelSpecification]
//   - [MLModelDescription.InitWithCoder]
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
//
// See: https://developer.apple.com/documentation/CoreML/MLModelDescription
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
//   - [IMLModelDescription.InitFromRawModelDescriptionSpecification]
//   - [IMLModelDescription.InitFromRawModelSpecification]
//   - [IMLModelDescription.InitFromSingleFunctionModelDescriptionSpecification]
//   - [IMLModelDescription.InitFromSingleFunctionModelSpecification]
//   - [IMLModelDescription.InitWithCoder]
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
//
// See: https://developer.apple.com/documentation/CoreML/MLModelDescription
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
	ModelURL() foundation.INSURL
	SetModelURL(value foundation.INSURL)
	OutputFeatureNames() foundation.INSOrderedSet
	PredictedClassFeatureDescription() IMLFeatureDescription
	PredictedValueFeatureDescription() IMLFeatureDescription
	StateFeatureNames() foundation.INSOrderedSet
	ValidateAsClassifierDescriptionAndReturnError() (bool, error)
	ValidateAsRegressorDescriptionAndReturnError() (bool, error)
	VerifyInputError(input objectivec.IObject) (bool, error)
	InitFromModelDescriptionSpecification(specification unsafe.Pointer) MLModelDescription
	InitFromRawModelDescriptionSpecification(specification unsafe.Pointer) MLModelDescription
	InitFromRawModelSpecification(specification unsafe.Pointer) MLModelDescription
	InitFromSingleFunctionModelDescriptionSpecification(specification unsafe.Pointer) MLModelDescription
	InitFromSingleFunctionModelSpecification(specification unsafe.Pointer) MLModelDescription
	InitWithCoder(coder foundation.INSCoder) MLModelDescription
	InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionDescriptionsIsUpdatableTrainingInputDescriptionsParameterDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesMetadataDefaultFunctionNameFunctionNameClassLabelsModelURLModelPath(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions3 objectivec.IObject, updatable bool, descriptions4 objectivec.IObject, descriptions5 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, metadata objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, labels objectivec.IObject, url foundation.INSURL, path objectivec.IObject) MLModelDescription
	InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameMetadata(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, metadata objectivec.IObject) MLModelDescription
	InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameTrainingInputDescriptionsMetadata(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions3 objectivec.IObject, metadata objectivec.IObject) MLModelDescription
	InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameTrainingInputDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesMetadata(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions3 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, metadata objectivec.IObject) MLModelDescription
	InitWithInputDescriptionsOutputDescriptionsStateDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionDescriptionsIsUpdatableTrainingInputDescriptionsParameterDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesOrderedStateFeatureNamesMetadataDefaultFunctionNameFunctionNameClassLabelsModelURLModelPath(descriptions objectivec.IObject, descriptions2 objectivec.IObject, descriptions3 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions4 objectivec.IObject, updatable bool, descriptions5 objectivec.IObject, descriptions6 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, names3 objectivec.IObject, metadata objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, labels objectivec.IObject, url foundation.INSURL, path objectivec.IObject) MLModelDescription
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
	rv := objc.Send[MLModelDescription](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLModelDescription) Autorelease() MLModelDescription {
	rv := objc.Send[MLModelDescription](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLModelDescription creates a new MLModelDescription instance.
func NewMLModelDescription() MLModelDescription {
	class := getMLModelDescriptionClass()
	rv := objc.Send[MLModelDescription](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initFromModelDescriptionSpecification:
func NewModelDescriptionFromModelDescriptionSpecification(specification unsafe.Pointer) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initFromModelDescriptionSpecification:"), specification)
	return MLModelDescriptionFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initFromRawModelDescriptionSpecification:
func NewModelDescriptionFromRawModelDescriptionSpecification(specification unsafe.Pointer) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initFromRawModelDescriptionSpecification:"), specification)
	return MLModelDescriptionFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initFromRawModelSpecification:
func NewModelDescriptionFromRawModelSpecification(specification unsafe.Pointer) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initFromRawModelSpecification:"), specification)
	return MLModelDescriptionFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initFromSingleFunctionModelDescriptionSpecification:
func NewModelDescriptionFromSingleFunctionModelDescriptionSpecification(specification unsafe.Pointer) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initFromSingleFunctionModelDescriptionSpecification:"), specification)
	return MLModelDescriptionFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initFromSingleFunctionModelSpecification:
func NewModelDescriptionFromSingleFunctionModelSpecification(specification unsafe.Pointer) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initFromSingleFunctionModelSpecification:"), specification)
	return MLModelDescriptionFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initWithCoder:
func NewModelDescriptionWithCoder(coder objectivec.IObject) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return MLModelDescriptionFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:functionDescriptions:isUpdatable:trainingInputDescriptions:parameterDescriptions:orderedInputFeatureNames:orderedOutputFeatureNames:metadata:defaultFunctionName:functionName:classLabels:modelURL:modelPath:
func NewModelDescriptionWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionDescriptionsIsUpdatableTrainingInputDescriptionsParameterDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesMetadataDefaultFunctionNameFunctionNameClassLabelsModelURLModelPath(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions3 objectivec.IObject, updatable bool, descriptions4 objectivec.IObject, descriptions5 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, metadata objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, labels objectivec.IObject, url foundation.INSURL, path objectivec.IObject) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:functionDescriptions:isUpdatable:trainingInputDescriptions:parameterDescriptions:orderedInputFeatureNames:orderedOutputFeatureNames:metadata:defaultFunctionName:functionName:classLabels:modelURL:modelPath:"), descriptions, descriptions2, name, name2, descriptions3, updatable, descriptions4, descriptions5, names, names2, metadata, name3, name4, labels, url, path)
	return MLModelDescriptionFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:metadata:
func NewModelDescriptionWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameMetadata(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, metadata objectivec.IObject) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:metadata:"), descriptions, descriptions2, name, name2, metadata)
	return MLModelDescriptionFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:trainingInputDescriptions:metadata:
func NewModelDescriptionWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameTrainingInputDescriptionsMetadata(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions3 objectivec.IObject, metadata objectivec.IObject) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:trainingInputDescriptions:metadata:"), descriptions, descriptions2, name, name2, descriptions3, metadata)
	return MLModelDescriptionFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:trainingInputDescriptions:orderedInputFeatureNames:orderedOutputFeatureNames:metadata:
func NewModelDescriptionWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameTrainingInputDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesMetadata(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions3 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, metadata objectivec.IObject) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:trainingInputDescriptions:orderedInputFeatureNames:orderedOutputFeatureNames:metadata:"), descriptions, descriptions2, name, name2, descriptions3, names, names2, metadata)
	return MLModelDescriptionFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initWithInputDescriptions:outputDescriptions:stateDescriptions:predictedFeatureName:predictedProbabilitiesName:functionDescriptions:isUpdatable:trainingInputDescriptions:parameterDescriptions:orderedInputFeatureNames:orderedOutputFeatureNames:orderedStateFeatureNames:metadata:defaultFunctionName:functionName:classLabels:modelURL:modelPath:
func NewModelDescriptionWithInputDescriptionsOutputDescriptionsStateDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionDescriptionsIsUpdatableTrainingInputDescriptionsParameterDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesOrderedStateFeatureNamesMetadataDefaultFunctionNameFunctionNameClassLabelsModelURLModelPath(descriptions objectivec.IObject, descriptions2 objectivec.IObject, descriptions3 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions4 objectivec.IObject, updatable bool, descriptions5 objectivec.IObject, descriptions6 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, names3 objectivec.IObject, metadata objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, labels objectivec.IObject, url foundation.INSURL, path objectivec.IObject) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:stateDescriptions:predictedFeatureName:predictedProbabilitiesName:functionDescriptions:isUpdatable:trainingInputDescriptions:parameterDescriptions:orderedInputFeatureNames:orderedOutputFeatureNames:orderedStateFeatureNames:metadata:defaultFunctionName:functionName:classLabels:modelURL:modelPath:"), descriptions, descriptions2, descriptions3, name, name2, descriptions4, updatable, descriptions5, descriptions6, names, names2, names3, metadata, name3, name4, labels, url, path)
	return MLModelDescriptionFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initWithInputDescriptions:outputDescriptions:stateDescriptions:predictedFeatureName:predictedProbabilitiesName:functionName:
func NewModelDescriptionWithInputDescriptionsOutputDescriptionsStateDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionName(descriptions objectivec.IObject, descriptions2 objectivec.IObject, descriptions3 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject) MLModelDescription {
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:stateDescriptions:predictedFeatureName:predictedProbabilitiesName:functionName:"), descriptions, descriptions2, descriptions3, name, name2, name3)
	return MLModelDescriptionFromID(rv)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initWithModelDescriptionSpecification:error:
func NewModelDescriptionWithModelDescriptionSpecificationError(specification unsafe.Pointer) (MLModelDescription, error) {
	var errorPtr objc.ID
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelDescriptionSpecification:error:"), specification, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelDescription{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelDescriptionFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initWithModelSpecification:error:
func NewModelDescriptionWithModelSpecificationError(specification unsafe.Pointer) (MLModelDescription, error) {
	var errorPtr objc.ID
	instance := getMLModelDescriptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelSpecification:error:"), specification, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelDescription{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelDescriptionFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/debugQuickLookObject
func (m MLModelDescription) DebugQuickLookObject() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugQuickLookObject"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/defaultFunctionName
func (m MLModelDescription) DefaultFunctionName() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("defaultFunctionName"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/functionDescriptions
func (m MLModelDescription) FunctionDescriptions() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("functionDescriptions"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/hasEnumeratedShapeInputs
func (m MLModelDescription) HasEnumeratedShapeInputs() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasEnumeratedShapeInputs"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/hasRangeShapeInputs
func (m MLModelDescription) HasRangeShapeInputs() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("hasRangeShapeInputs"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/isEqualToDescription:
func (m MLModelDescription) IsEqualToDescription(description objectivec.IObject) bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isEqualToDescription:"), description)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/modelDescriptionBySettingMetadata:
func (m MLModelDescription) ModelDescriptionBySettingMetadata(metadata objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelDescriptionBySettingMetadata:"), metadata)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/validateAsClassifierDescriptionAndReturnError:
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

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/validateAsRegressorDescriptionAndReturnError:
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

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/verifyInput:error:
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

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initFromModelDescriptionSpecification:
func (m MLModelDescription) InitFromModelDescriptionSpecification(specification unsafe.Pointer) MLModelDescription {
	rv := objc.Send[MLModelDescription](m.ID, objc.Sel("initFromModelDescriptionSpecification:"), specification)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initFromRawModelDescriptionSpecification:
func (m MLModelDescription) InitFromRawModelDescriptionSpecification(specification unsafe.Pointer) MLModelDescription {
	rv := objc.Send[MLModelDescription](m.ID, objc.Sel("initFromRawModelDescriptionSpecification:"), specification)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initFromRawModelSpecification:
func (m MLModelDescription) InitFromRawModelSpecification(specification unsafe.Pointer) MLModelDescription {
	rv := objc.Send[MLModelDescription](m.ID, objc.Sel("initFromRawModelSpecification:"), specification)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initFromSingleFunctionModelDescriptionSpecification:
func (m MLModelDescription) InitFromSingleFunctionModelDescriptionSpecification(specification unsafe.Pointer) MLModelDescription {
	rv := objc.Send[MLModelDescription](m.ID, objc.Sel("initFromSingleFunctionModelDescriptionSpecification:"), specification)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initFromSingleFunctionModelSpecification:
func (m MLModelDescription) InitFromSingleFunctionModelSpecification(specification unsafe.Pointer) MLModelDescription {
	rv := objc.Send[MLModelDescription](m.ID, objc.Sel("initFromSingleFunctionModelSpecification:"), specification)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initWithCoder:
func (m MLModelDescription) InitWithCoder(coder foundation.INSCoder) MLModelDescription {
	rv := objc.Send[MLModelDescription](m.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:functionDescriptions:isUpdatable:trainingInputDescriptions:parameterDescriptions:orderedInputFeatureNames:orderedOutputFeatureNames:metadata:defaultFunctionName:functionName:classLabels:modelURL:modelPath:
func (m MLModelDescription) InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionDescriptionsIsUpdatableTrainingInputDescriptionsParameterDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesMetadataDefaultFunctionNameFunctionNameClassLabelsModelURLModelPath(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions3 objectivec.IObject, updatable bool, descriptions4 objectivec.IObject, descriptions5 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, metadata objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, labels objectivec.IObject, url foundation.INSURL, path objectivec.IObject) MLModelDescription {
	rv := objc.Send[MLModelDescription](m.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:functionDescriptions:isUpdatable:trainingInputDescriptions:parameterDescriptions:orderedInputFeatureNames:orderedOutputFeatureNames:metadata:defaultFunctionName:functionName:classLabels:modelURL:modelPath:"), descriptions, descriptions2, name, name2, descriptions3, updatable, descriptions4, descriptions5, names, names2, metadata, name3, name4, labels, url, path)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:metadata:
func (m MLModelDescription) InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameMetadata(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, metadata objectivec.IObject) MLModelDescription {
	rv := objc.Send[MLModelDescription](m.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:metadata:"), descriptions, descriptions2, name, name2, metadata)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:trainingInputDescriptions:metadata:
func (m MLModelDescription) InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameTrainingInputDescriptionsMetadata(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions3 objectivec.IObject, metadata objectivec.IObject) MLModelDescription {
	rv := objc.Send[MLModelDescription](m.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:trainingInputDescriptions:metadata:"), descriptions, descriptions2, name, name2, descriptions3, metadata)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:trainingInputDescriptions:orderedInputFeatureNames:orderedOutputFeatureNames:metadata:
func (m MLModelDescription) InitWithInputDescriptionsOutputDescriptionsPredictedFeatureNamePredictedProbabilitiesNameTrainingInputDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesMetadata(descriptions objectivec.IObject, descriptions2 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions3 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, metadata objectivec.IObject) MLModelDescription {
	rv := objc.Send[MLModelDescription](m.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:predictedFeatureName:predictedProbabilitiesName:trainingInputDescriptions:orderedInputFeatureNames:orderedOutputFeatureNames:metadata:"), descriptions, descriptions2, name, name2, descriptions3, names, names2, metadata)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initWithInputDescriptions:outputDescriptions:stateDescriptions:predictedFeatureName:predictedProbabilitiesName:functionDescriptions:isUpdatable:trainingInputDescriptions:parameterDescriptions:orderedInputFeatureNames:orderedOutputFeatureNames:orderedStateFeatureNames:metadata:defaultFunctionName:functionName:classLabels:modelURL:modelPath:
func (m MLModelDescription) InitWithInputDescriptionsOutputDescriptionsStateDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionDescriptionsIsUpdatableTrainingInputDescriptionsParameterDescriptionsOrderedInputFeatureNamesOrderedOutputFeatureNamesOrderedStateFeatureNamesMetadataDefaultFunctionNameFunctionNameClassLabelsModelURLModelPath(descriptions objectivec.IObject, descriptions2 objectivec.IObject, descriptions3 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, descriptions4 objectivec.IObject, updatable bool, descriptions5 objectivec.IObject, descriptions6 objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, names3 objectivec.IObject, metadata objectivec.IObject, name3 objectivec.IObject, name4 objectivec.IObject, labels objectivec.IObject, url foundation.INSURL, path objectivec.IObject) MLModelDescription {
	rv := objc.Send[MLModelDescription](m.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:stateDescriptions:predictedFeatureName:predictedProbabilitiesName:functionDescriptions:isUpdatable:trainingInputDescriptions:parameterDescriptions:orderedInputFeatureNames:orderedOutputFeatureNames:orderedStateFeatureNames:metadata:defaultFunctionName:functionName:classLabels:modelURL:modelPath:"), descriptions, descriptions2, descriptions3, name, name2, descriptions4, updatable, descriptions5, descriptions6, names, names2, names3, metadata, name3, name4, labels, url, path)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initWithInputDescriptions:outputDescriptions:stateDescriptions:predictedFeatureName:predictedProbabilitiesName:functionName:
func (m MLModelDescription) InitWithInputDescriptionsOutputDescriptionsStateDescriptionsPredictedFeatureNamePredictedProbabilitiesNameFunctionName(descriptions objectivec.IObject, descriptions2 objectivec.IObject, descriptions3 objectivec.IObject, name objectivec.IObject, name2 objectivec.IObject, name3 objectivec.IObject) MLModelDescription {
	rv := objc.Send[MLModelDescription](m.ID, objc.Sel("initWithInputDescriptions:outputDescriptions:stateDescriptions:predictedFeatureName:predictedProbabilitiesName:functionName:"), descriptions, descriptions2, descriptions3, name, name2, name3)
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initWithModelDescriptionSpecification:error:
func (m MLModelDescription) InitWithModelDescriptionSpecificationError(specification unsafe.Pointer) (MLModelDescription, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithModelDescriptionSpecification:error:"), specification, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelDescription{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelDescriptionFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/initWithModelSpecification:error:
func (m MLModelDescription) InitWithModelSpecificationError(specification unsafe.Pointer) (MLModelDescription, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithModelSpecification:error:"), specification, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLModelDescription{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLModelDescriptionFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/metadataWithFormat:
func (_MLModelDescriptionClass MLModelDescriptionClass) MetadataWithFormat(format unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelDescriptionClass.class), objc.Sel("metadataWithFormat:"), format)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/metadataWithSpecification:
func (_MLModelDescriptionClass MLModelDescriptionClass) MetadataWithSpecification(specification unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLModelDescriptionClass.class), objc.Sel("metadataWithSpecification:"), specification)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/supportsSecureCoding
func (_MLModelDescriptionClass MLModelDescriptionClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_MLModelDescriptionClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/classLabels
func (m MLModelDescription) ClassLabels() foundation.INSArray {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classLabels"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (m MLModelDescription) SetClassLabels(value foundation.INSArray) {
	objc.Send[struct{}](m.ID, objc.Sel("setClassLabels:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/classProbabilityFeatureDescription
func (m MLModelDescription) ClassProbabilityFeatureDescription() IMLFeatureDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classProbabilityFeatureDescription"))
	return MLFeatureDescriptionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/functionName
func (m MLModelDescription) FunctionName() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("functionName"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/inputFeatureNames
func (m MLModelDescription) InputFeatureNames() foundation.INSOrderedSet {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("inputFeatureNames"))
	return foundation.NSOrderedSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/isUpdatable
func (m MLModelDescription) IsUpdatable() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isUpdatable"))
	return rv
}
func (m MLModelDescription) SetIsUpdatable(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setIsUpdatable:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/modelPath
func (m MLModelDescription) ModelPath() IMLLayerPath {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelPath"))
	return MLLayerPathFromID(objc.ID(rv))
}
func (m MLModelDescription) SetModelPath(value IMLLayerPath) {
	objc.Send[struct{}](m.ID, objc.Sel("setModelPath:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/modelURL
func (m MLModelDescription) ModelURL() foundation.INSURL {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (m MLModelDescription) SetModelURL(value foundation.INSURL) {
	objc.Send[struct{}](m.ID, objc.Sel("setModelURL:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/outputFeatureNames
func (m MLModelDescription) OutputFeatureNames() foundation.INSOrderedSet {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("outputFeatureNames"))
	return foundation.NSOrderedSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/parameterDescriptionsByKey
func (m MLModelDescription) ParameterDescriptionsByKey() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parameterDescriptionsByKey"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLModelDescription) SetParameterDescriptionsByKey(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setParameterDescriptionsByKey:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedClassFeatureDescription
func (m MLModelDescription) PredictedClassFeatureDescription() IMLFeatureDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictedClassFeatureDescription"))
	return MLFeatureDescriptionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/predictedValueFeatureDescription
func (m MLModelDescription) PredictedValueFeatureDescription() IMLFeatureDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("predictedValueFeatureDescription"))
	return MLFeatureDescriptionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/stateFeatureNames
func (m MLModelDescription) StateFeatureNames() foundation.INSOrderedSet {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("stateFeatureNames"))
	return foundation.NSOrderedSetFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLModelDescription/trainingInputDescriptionsByName
func (m MLModelDescription) TrainingInputDescriptionsByName() foundation.INSDictionary {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("trainingInputDescriptionsByName"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (m MLModelDescription) SetTrainingInputDescriptionsByName(value foundation.INSDictionary) {
	objc.Send[struct{}](m.ID, objc.Sel("setTrainingInputDescriptionsByName:"), value)
}
