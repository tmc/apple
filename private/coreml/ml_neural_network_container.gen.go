// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MLNeuralNetworkContainer] class.
var (
	_MLNeuralNetworkContainerClass     MLNeuralNetworkContainerClass
	_MLNeuralNetworkContainerClassOnce sync.Once
)

func getMLNeuralNetworkContainerClass() MLNeuralNetworkContainerClass {
	_MLNeuralNetworkContainerClassOnce.Do(func() {
		_MLNeuralNetworkContainerClass = MLNeuralNetworkContainerClass{class: objc.GetClass("MLNeuralNetworkContainer")}
	})
	return _MLNeuralNetworkContainerClass
}

// GetMLNeuralNetworkContainerClass returns the class object for MLNeuralNetworkContainer.
func GetMLNeuralNetworkContainerClass() MLNeuralNetworkContainerClass {
	return getMLNeuralNetworkContainerClass()
}

type MLNeuralNetworkContainerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLNeuralNetworkContainerClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLNeuralNetworkContainerClass) Alloc() MLNeuralNetworkContainer {
	rv := objc.Send[MLNeuralNetworkContainer](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [MLNeuralNetworkContainer.ActiveFunction]
//   - [MLNeuralNetworkContainer.ClassLabels]
//   - [MLNeuralNetworkContainer.SetClassLabels]
//   - [MLNeuralNetworkContainer.ClassScoreVectorName]
//   - [MLNeuralNetworkContainer.SetClassScoreVectorName]
//   - [MLNeuralNetworkContainer.CompilerOutput]
//   - [MLNeuralNetworkContainer.SetCompilerOutput]
//   - [MLNeuralNetworkContainer.CompilerVersionInfo]
//   - [MLNeuralNetworkContainer.SetCompilerVersionInfo]
//   - [MLNeuralNetworkContainer.ConfigurationList]
//   - [MLNeuralNetworkContainer.SetConfigurationList]
//   - [MLNeuralNetworkContainer.Engine]
//   - [MLNeuralNetworkContainer.SetEngine]
//   - [MLNeuralNetworkContainer.HasBidirectionalLayer]
//   - [MLNeuralNetworkContainer.SetHasBidirectionalLayer]
//   - [MLNeuralNetworkContainer.HasDynamicLayer]
//   - [MLNeuralNetworkContainer.SetHasDynamicLayer]
//   - [MLNeuralNetworkContainer.HasOptionalInputSequenceConcat]
//   - [MLNeuralNetworkContainer.SetHasOptionalInputSequenceConcat]
//   - [MLNeuralNetworkContainer.ImageParameters]
//   - [MLNeuralNetworkContainer.SetImageParameters]
//   - [MLNeuralNetworkContainer.ImagePreprocessingParams]
//   - [MLNeuralNetworkContainer.SetImagePreprocessingParams]
//   - [MLNeuralNetworkContainer.InputDescription]
//   - [MLNeuralNetworkContainer.SetInputDescription]
//   - [MLNeuralNetworkContainer.InputLayerNames]
//   - [MLNeuralNetworkContainer.SetInputLayerNames]
//   - [MLNeuralNetworkContainer.ModelDescription]
//   - [MLNeuralNetworkContainer.ModelFilePath]
//   - [MLNeuralNetworkContainer.SetModelFilePath]
//   - [MLNeuralNetworkContainer.ModelIsEncrypted]
//   - [MLNeuralNetworkContainer.SetModelIsEncrypted]
//   - [MLNeuralNetworkContainer.ModelIsMIL]
//   - [MLNeuralNetworkContainer.SetModelIsMIL]
//   - [MLNeuralNetworkContainer.ModelIsTrainingProgram]
//   - [MLNeuralNetworkContainer.SetModelIsTrainingProgram]
//   - [MLNeuralNetworkContainer.ModelVersionInfo]
//   - [MLNeuralNetworkContainer.SetModelVersionInfo]
//   - [MLNeuralNetworkContainer.Name]
//   - [MLNeuralNetworkContainer.SetName]
//   - [MLNeuralNetworkContainer.NdArrayInterpretation]
//   - [MLNeuralNetworkContainer.SetNdArrayInterpretation]
//   - [MLNeuralNetworkContainer.OptionalInputDefaultValues]
//   - [MLNeuralNetworkContainer.SetOptionalInputDefaultValues]
//   - [MLNeuralNetworkContainer.OptionalInputTypes]
//   - [MLNeuralNetworkContainer.SetOptionalInputTypes]
//   - [MLNeuralNetworkContainer.OutputDescription]
//   - [MLNeuralNetworkContainer.SetOutputDescription]
//   - [MLNeuralNetworkContainer.OutputLayerNames]
//   - [MLNeuralNetworkContainer.SetOutputLayerNames]
//   - [MLNeuralNetworkContainer.Precision]
//   - [MLNeuralNetworkContainer.SetPrecision]
//   - [MLNeuralNetworkContainer.UpdatableModelCompiledParams]
//   - [MLNeuralNetworkContainer.SetUpdatableModelCompiledParams]
//   - [MLNeuralNetworkContainer.InitWithFeatureDescriptionsModelDescriptionOutputLayerNamesClassScoreVectorNameClassLabelsIsEncryptedModelVersionInfo]
//   - [MLNeuralNetworkContainer.InitWithFeatureDescriptionsModelDescriptionOutputLayerNamesClassScoreVectorNameClassLabelsIsEncryptedModelVersionInfoCompilerVersionInfo]
//   - [MLNeuralNetworkContainer.InitWithFilePathInputLayerNamesOutputLayerNamesParameters]
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer
type MLNeuralNetworkContainer struct {
	objectivec.Object
}

// MLNeuralNetworkContainerFromID constructs a [MLNeuralNetworkContainer] from an objc.ID.
func MLNeuralNetworkContainerFromID(id objc.ID) MLNeuralNetworkContainer {
	return MLNeuralNetworkContainer{objectivec.Object{ID: id}}
}
// Ensure MLNeuralNetworkContainer implements IMLNeuralNetworkContainer.
var _ IMLNeuralNetworkContainer = MLNeuralNetworkContainer{}

// An interface definition for the [MLNeuralNetworkContainer] class.
//
// # Methods
//
//   - [IMLNeuralNetworkContainer.ActiveFunction]
//   - [IMLNeuralNetworkContainer.ClassLabels]
//   - [IMLNeuralNetworkContainer.SetClassLabels]
//   - [IMLNeuralNetworkContainer.ClassScoreVectorName]
//   - [IMLNeuralNetworkContainer.SetClassScoreVectorName]
//   - [IMLNeuralNetworkContainer.CompilerOutput]
//   - [IMLNeuralNetworkContainer.SetCompilerOutput]
//   - [IMLNeuralNetworkContainer.CompilerVersionInfo]
//   - [IMLNeuralNetworkContainer.SetCompilerVersionInfo]
//   - [IMLNeuralNetworkContainer.ConfigurationList]
//   - [IMLNeuralNetworkContainer.SetConfigurationList]
//   - [IMLNeuralNetworkContainer.Engine]
//   - [IMLNeuralNetworkContainer.SetEngine]
//   - [IMLNeuralNetworkContainer.HasBidirectionalLayer]
//   - [IMLNeuralNetworkContainer.SetHasBidirectionalLayer]
//   - [IMLNeuralNetworkContainer.HasDynamicLayer]
//   - [IMLNeuralNetworkContainer.SetHasDynamicLayer]
//   - [IMLNeuralNetworkContainer.HasOptionalInputSequenceConcat]
//   - [IMLNeuralNetworkContainer.SetHasOptionalInputSequenceConcat]
//   - [IMLNeuralNetworkContainer.ImageParameters]
//   - [IMLNeuralNetworkContainer.SetImageParameters]
//   - [IMLNeuralNetworkContainer.ImagePreprocessingParams]
//   - [IMLNeuralNetworkContainer.SetImagePreprocessingParams]
//   - [IMLNeuralNetworkContainer.InputDescription]
//   - [IMLNeuralNetworkContainer.SetInputDescription]
//   - [IMLNeuralNetworkContainer.InputLayerNames]
//   - [IMLNeuralNetworkContainer.SetInputLayerNames]
//   - [IMLNeuralNetworkContainer.ModelDescription]
//   - [IMLNeuralNetworkContainer.ModelFilePath]
//   - [IMLNeuralNetworkContainer.SetModelFilePath]
//   - [IMLNeuralNetworkContainer.ModelIsEncrypted]
//   - [IMLNeuralNetworkContainer.SetModelIsEncrypted]
//   - [IMLNeuralNetworkContainer.ModelIsMIL]
//   - [IMLNeuralNetworkContainer.SetModelIsMIL]
//   - [IMLNeuralNetworkContainer.ModelIsTrainingProgram]
//   - [IMLNeuralNetworkContainer.SetModelIsTrainingProgram]
//   - [IMLNeuralNetworkContainer.ModelVersionInfo]
//   - [IMLNeuralNetworkContainer.SetModelVersionInfo]
//   - [IMLNeuralNetworkContainer.Name]
//   - [IMLNeuralNetworkContainer.SetName]
//   - [IMLNeuralNetworkContainer.NdArrayInterpretation]
//   - [IMLNeuralNetworkContainer.SetNdArrayInterpretation]
//   - [IMLNeuralNetworkContainer.OptionalInputDefaultValues]
//   - [IMLNeuralNetworkContainer.SetOptionalInputDefaultValues]
//   - [IMLNeuralNetworkContainer.OptionalInputTypes]
//   - [IMLNeuralNetworkContainer.SetOptionalInputTypes]
//   - [IMLNeuralNetworkContainer.OutputDescription]
//   - [IMLNeuralNetworkContainer.SetOutputDescription]
//   - [IMLNeuralNetworkContainer.OutputLayerNames]
//   - [IMLNeuralNetworkContainer.SetOutputLayerNames]
//   - [IMLNeuralNetworkContainer.Precision]
//   - [IMLNeuralNetworkContainer.SetPrecision]
//   - [IMLNeuralNetworkContainer.UpdatableModelCompiledParams]
//   - [IMLNeuralNetworkContainer.SetUpdatableModelCompiledParams]
//   - [IMLNeuralNetworkContainer.InitWithFeatureDescriptionsModelDescriptionOutputLayerNamesClassScoreVectorNameClassLabelsIsEncryptedModelVersionInfo]
//   - [IMLNeuralNetworkContainer.InitWithFeatureDescriptionsModelDescriptionOutputLayerNamesClassScoreVectorNameClassLabelsIsEncryptedModelVersionInfoCompilerVersionInfo]
//   - [IMLNeuralNetworkContainer.InitWithFilePathInputLayerNamesOutputLayerNamesParameters]
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer
type IMLNeuralNetworkContainer interface {
	objectivec.IObject

	// Topic: Methods

	ActiveFunction() string
	ClassLabels() foundation.INSArray
	SetClassLabels(value foundation.INSArray)
	ClassScoreVectorName() string
	SetClassScoreVectorName(value string)
	CompilerOutput() IMLCompilerNeuralNetworkOutput
	SetCompilerOutput(value IMLCompilerNeuralNetworkOutput)
	CompilerVersionInfo() IMLVersionInfo
	SetCompilerVersionInfo(value IMLVersionInfo)
	ConfigurationList() foundation.INSArray
	SetConfigurationList(value foundation.INSArray)
	Engine() int
	SetEngine(value int)
	HasBidirectionalLayer() bool
	SetHasBidirectionalLayer(value bool)
	HasDynamicLayer() bool
	SetHasDynamicLayer(value bool)
	HasOptionalInputSequenceConcat() bool
	SetHasOptionalInputSequenceConcat(value bool)
	ImageParameters() foundation.INSDictionary
	SetImageParameters(value foundation.INSDictionary)
	ImagePreprocessingParams() foundation.INSDictionary
	SetImagePreprocessingParams(value foundation.INSDictionary)
	InputDescription() foundation.INSDictionary
	SetInputDescription(value foundation.INSDictionary)
	InputLayerNames() foundation.INSArray
	SetInputLayerNames(value foundation.INSArray)
	ModelDescription() IMLModelDescription
	ModelFilePath() string
	SetModelFilePath(value string)
	ModelIsEncrypted() bool
	SetModelIsEncrypted(value bool)
	ModelIsMIL() bool
	SetModelIsMIL(value bool)
	ModelIsTrainingProgram() bool
	SetModelIsTrainingProgram(value bool)
	ModelVersionInfo() IMLVersionInfo
	SetModelVersionInfo(value IMLVersionInfo)
	Name() string
	SetName(value string)
	NdArrayInterpretation() bool
	SetNdArrayInterpretation(value bool)
	OptionalInputDefaultValues() foundation.INSDictionary
	SetOptionalInputDefaultValues(value foundation.INSDictionary)
	OptionalInputTypes() foundation.INSDictionary
	SetOptionalInputTypes(value foundation.INSDictionary)
	OutputDescription() foundation.INSDictionary
	SetOutputDescription(value foundation.INSDictionary)
	OutputLayerNames() foundation.INSArray
	SetOutputLayerNames(value foundation.INSArray)
	Precision() int
	SetPrecision(value int)
	UpdatableModelCompiledParams() IMLNeuralNetworksCompileTimeParams
	SetUpdatableModelCompiledParams(value IMLNeuralNetworksCompileTimeParams)
	InitWithFeatureDescriptionsModelDescriptionOutputLayerNamesClassScoreVectorNameClassLabelsIsEncryptedModelVersionInfo(descriptions objectivec.IObject, description objectivec.IObject, names objectivec.IObject, name objectivec.IObject, labels objectivec.IObject, encrypted bool, info objectivec.IObject) MLNeuralNetworkContainer
	InitWithFeatureDescriptionsModelDescriptionOutputLayerNamesClassScoreVectorNameClassLabelsIsEncryptedModelVersionInfoCompilerVersionInfo(descriptions objectivec.IObject, description objectivec.IObject, names objectivec.IObject, name objectivec.IObject, labels objectivec.IObject, encrypted bool, info objectivec.IObject, info2 objectivec.IObject) MLNeuralNetworkContainer
	InitWithFilePathInputLayerNamesOutputLayerNamesParameters(path objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, parameters objectivec.IObject) MLNeuralNetworkContainer
}

// Init initializes the instance.
func (n MLNeuralNetworkContainer) Init() MLNeuralNetworkContainer {
	rv := objc.Send[MLNeuralNetworkContainer](n.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (n MLNeuralNetworkContainer) Autorelease() MLNeuralNetworkContainer {
	rv := objc.Send[MLNeuralNetworkContainer](n.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLNeuralNetworkContainer creates a new MLNeuralNetworkContainer instance.
func NewMLNeuralNetworkContainer() MLNeuralNetworkContainer {
	class := getMLNeuralNetworkContainerClass()
	rv := objc.Send[MLNeuralNetworkContainer](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/initWithFeatureDescriptions:modelDescription:outputLayerNames:classScoreVectorName:classLabels:isEncrypted:modelVersionInfo:
func NewNeuralNetworkContainerWithFeatureDescriptionsModelDescriptionOutputLayerNamesClassScoreVectorNameClassLabelsIsEncryptedModelVersionInfo(descriptions objectivec.IObject, description objectivec.IObject, names objectivec.IObject, name objectivec.IObject, labels objectivec.IObject, encrypted bool, info objectivec.IObject) MLNeuralNetworkContainer {
	instance := getMLNeuralNetworkContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFeatureDescriptions:modelDescription:outputLayerNames:classScoreVectorName:classLabels:isEncrypted:modelVersionInfo:"), descriptions, description, names, name, labels, encrypted, info)
	return MLNeuralNetworkContainerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/initWithFeatureDescriptions:modelDescription:outputLayerNames:classScoreVectorName:classLabels:isEncrypted:modelVersionInfo:compilerVersionInfo:
func NewNeuralNetworkContainerWithFeatureDescriptionsModelDescriptionOutputLayerNamesClassScoreVectorNameClassLabelsIsEncryptedModelVersionInfoCompilerVersionInfo(descriptions objectivec.IObject, description objectivec.IObject, names objectivec.IObject, name objectivec.IObject, labels objectivec.IObject, encrypted bool, info objectivec.IObject, info2 objectivec.IObject) MLNeuralNetworkContainer {
	instance := getMLNeuralNetworkContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFeatureDescriptions:modelDescription:outputLayerNames:classScoreVectorName:classLabels:isEncrypted:modelVersionInfo:compilerVersionInfo:"), descriptions, description, names, name, labels, encrypted, info, info2)
	return MLNeuralNetworkContainerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/initWithFilePath:inputLayerNames:outputLayerNames:parameters:
func NewNeuralNetworkContainerWithFilePathInputLayerNamesOutputLayerNamesParameters(path objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, parameters objectivec.IObject) MLNeuralNetworkContainer {
	instance := getMLNeuralNetworkContainerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithFilePath:inputLayerNames:outputLayerNames:parameters:"), path, names, names2, parameters)
	return MLNeuralNetworkContainerFromID(rv)
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/initWithFeatureDescriptions:modelDescription:outputLayerNames:classScoreVectorName:classLabels:isEncrypted:modelVersionInfo:
func (n MLNeuralNetworkContainer) InitWithFeatureDescriptionsModelDescriptionOutputLayerNamesClassScoreVectorNameClassLabelsIsEncryptedModelVersionInfo(descriptions objectivec.IObject, description objectivec.IObject, names objectivec.IObject, name objectivec.IObject, labels objectivec.IObject, encrypted bool, info objectivec.IObject) MLNeuralNetworkContainer {
	rv := objc.Send[MLNeuralNetworkContainer](n.ID, objc.Sel("initWithFeatureDescriptions:modelDescription:outputLayerNames:classScoreVectorName:classLabels:isEncrypted:modelVersionInfo:"), descriptions, description, names, name, labels, encrypted, info)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/initWithFeatureDescriptions:modelDescription:outputLayerNames:classScoreVectorName:classLabels:isEncrypted:modelVersionInfo:compilerVersionInfo:
func (n MLNeuralNetworkContainer) InitWithFeatureDescriptionsModelDescriptionOutputLayerNamesClassScoreVectorNameClassLabelsIsEncryptedModelVersionInfoCompilerVersionInfo(descriptions objectivec.IObject, description objectivec.IObject, names objectivec.IObject, name objectivec.IObject, labels objectivec.IObject, encrypted bool, info objectivec.IObject, info2 objectivec.IObject) MLNeuralNetworkContainer {
	rv := objc.Send[MLNeuralNetworkContainer](n.ID, objc.Sel("initWithFeatureDescriptions:modelDescription:outputLayerNames:classScoreVectorName:classLabels:isEncrypted:modelVersionInfo:compilerVersionInfo:"), descriptions, description, names, name, labels, encrypted, info, info2)
	return rv
}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/initWithFilePath:inputLayerNames:outputLayerNames:parameters:
func (n MLNeuralNetworkContainer) InitWithFilePathInputLayerNamesOutputLayerNamesParameters(path objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, parameters objectivec.IObject) MLNeuralNetworkContainer {
	rv := objc.Send[MLNeuralNetworkContainer](n.ID, objc.Sel("initWithFilePath:inputLayerNames:outputLayerNames:parameters:"), path, names, names2, parameters)
	return rv
}

//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/containerFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:
func (_MLNeuralNetworkContainerClass MLNeuralNetworkContainerClass) ContainerFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkContainerClass.class), objc.Sel("containerFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/containerFromCompiledArchiveCommon:filename:modelVersionInfo:compilerVersionInfo:configuration:error:
func (_MLNeuralNetworkContainerClass MLNeuralNetworkContainerClass) ContainerFromCompiledArchiveCommonFilenameModelVersionInfoCompilerVersionInfoConfigurationError(common unsafe.Pointer, filename objectivec.IObject, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkContainerClass.class), objc.Sel("containerFromCompiledArchiveCommon:filename:modelVersionInfo:compilerVersionInfo:configuration:error:"), common, filename, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/containerFromFilePath:inputLayerNames:outputLayerNames:parameters:
func (_MLNeuralNetworkContainerClass MLNeuralNetworkContainerClass) ContainerFromFilePathInputLayerNamesOutputLayerNamesParameters(path objectivec.IObject, names objectivec.IObject, names2 objectivec.IObject, parameters objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_MLNeuralNetworkContainerClass.class), objc.Sel("containerFromFilePath:inputLayerNames:outputLayerNames:parameters:"), path, names, names2, parameters)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/activeFunction
func (n MLNeuralNetworkContainer) ActiveFunction() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("activeFunction"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/classLabels
func (n MLNeuralNetworkContainer) ClassLabels() foundation.INSArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("classLabels"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (n MLNeuralNetworkContainer) SetClassLabels(value foundation.INSArray) {
	objc.Send[struct{}](n.ID, objc.Sel("setClassLabels:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/classScoreVectorName
func (n MLNeuralNetworkContainer) ClassScoreVectorName() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("classScoreVectorName"))
	return foundation.NSStringFromID(rv).String()
}
func (n MLNeuralNetworkContainer) SetClassScoreVectorName(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setClassScoreVectorName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/compilerOutput
func (n MLNeuralNetworkContainer) CompilerOutput() IMLCompilerNeuralNetworkOutput {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("compilerOutput"))
	return MLCompilerNeuralNetworkOutputFromID(objc.ID(rv))
}
func (n MLNeuralNetworkContainer) SetCompilerOutput(value IMLCompilerNeuralNetworkOutput) {
	objc.Send[struct{}](n.ID, objc.Sel("setCompilerOutput:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/compilerVersionInfo
func (n MLNeuralNetworkContainer) CompilerVersionInfo() IMLVersionInfo {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("compilerVersionInfo"))
	return MLVersionInfoFromID(objc.ID(rv))
}
func (n MLNeuralNetworkContainer) SetCompilerVersionInfo(value IMLVersionInfo) {
	objc.Send[struct{}](n.ID, objc.Sel("setCompilerVersionInfo:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/configurationList
func (n MLNeuralNetworkContainer) ConfigurationList() foundation.INSArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("configurationList"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (n MLNeuralNetworkContainer) SetConfigurationList(value foundation.INSArray) {
	objc.Send[struct{}](n.ID, objc.Sel("setConfigurationList:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/engine
func (n MLNeuralNetworkContainer) Engine() int {
	rv := objc.Send[int](n.ID, objc.Sel("engine"))
	return rv
}
func (n MLNeuralNetworkContainer) SetEngine(value int) {
	objc.Send[struct{}](n.ID, objc.Sel("setEngine:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/hasBidirectionalLayer
func (n MLNeuralNetworkContainer) HasBidirectionalLayer() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasBidirectionalLayer"))
	return rv
}
func (n MLNeuralNetworkContainer) SetHasBidirectionalLayer(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setHasBidirectionalLayer:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/hasDynamicLayer
func (n MLNeuralNetworkContainer) HasDynamicLayer() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasDynamicLayer"))
	return rv
}
func (n MLNeuralNetworkContainer) SetHasDynamicLayer(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setHasDynamicLayer:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/hasOptionalInputSequenceConcat
func (n MLNeuralNetworkContainer) HasOptionalInputSequenceConcat() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("hasOptionalInputSequenceConcat"))
	return rv
}
func (n MLNeuralNetworkContainer) SetHasOptionalInputSequenceConcat(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setHasOptionalInputSequenceConcat:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/imageParameters
func (n MLNeuralNetworkContainer) ImageParameters() foundation.INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("imageParameters"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (n MLNeuralNetworkContainer) SetImageParameters(value foundation.INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setImageParameters:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/imagePreprocessingParams
func (n MLNeuralNetworkContainer) ImagePreprocessingParams() foundation.INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("imagePreprocessingParams"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (n MLNeuralNetworkContainer) SetImagePreprocessingParams(value foundation.INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setImagePreprocessingParams:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/inputDescription
func (n MLNeuralNetworkContainer) InputDescription() foundation.INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("inputDescription"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (n MLNeuralNetworkContainer) SetInputDescription(value foundation.INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setInputDescription:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/inputLayerNames
func (n MLNeuralNetworkContainer) InputLayerNames() foundation.INSArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("inputLayerNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (n MLNeuralNetworkContainer) SetInputLayerNames(value foundation.INSArray) {
	objc.Send[struct{}](n.ID, objc.Sel("setInputLayerNames:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/modelDescription
func (n MLNeuralNetworkContainer) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/modelFilePath
func (n MLNeuralNetworkContainer) ModelFilePath() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("modelFilePath"))
	return foundation.NSStringFromID(rv).String()
}
func (n MLNeuralNetworkContainer) SetModelFilePath(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setModelFilePath:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/modelIsEncrypted
func (n MLNeuralNetworkContainer) ModelIsEncrypted() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("modelIsEncrypted"))
	return rv
}
func (n MLNeuralNetworkContainer) SetModelIsEncrypted(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setModelIsEncrypted:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/modelIsMIL
func (n MLNeuralNetworkContainer) ModelIsMIL() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("modelIsMIL"))
	return rv
}
func (n MLNeuralNetworkContainer) SetModelIsMIL(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setModelIsMIL:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/modelIsTrainingProgram
func (n MLNeuralNetworkContainer) ModelIsTrainingProgram() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("modelIsTrainingProgram"))
	return rv
}
func (n MLNeuralNetworkContainer) SetModelIsTrainingProgram(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setModelIsTrainingProgram:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/modelVersionInfo
func (n MLNeuralNetworkContainer) ModelVersionInfo() IMLVersionInfo {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("modelVersionInfo"))
	return MLVersionInfoFromID(objc.ID(rv))
}
func (n MLNeuralNetworkContainer) SetModelVersionInfo(value IMLVersionInfo) {
	objc.Send[struct{}](n.ID, objc.Sel("setModelVersionInfo:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/name
func (n MLNeuralNetworkContainer) Name() string {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (n MLNeuralNetworkContainer) SetName(value string) {
	objc.Send[struct{}](n.ID, objc.Sel("setName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/ndArrayInterpretation
func (n MLNeuralNetworkContainer) NdArrayInterpretation() bool {
	rv := objc.Send[bool](n.ID, objc.Sel("ndArrayInterpretation"))
	return rv
}
func (n MLNeuralNetworkContainer) SetNdArrayInterpretation(value bool) {
	objc.Send[struct{}](n.ID, objc.Sel("setNdArrayInterpretation:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/optionalInputDefaultValues
func (n MLNeuralNetworkContainer) OptionalInputDefaultValues() foundation.INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("optionalInputDefaultValues"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (n MLNeuralNetworkContainer) SetOptionalInputDefaultValues(value foundation.INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setOptionalInputDefaultValues:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/optionalInputTypes
func (n MLNeuralNetworkContainer) OptionalInputTypes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("optionalInputTypes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (n MLNeuralNetworkContainer) SetOptionalInputTypes(value foundation.INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setOptionalInputTypes:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/outputDescription
func (n MLNeuralNetworkContainer) OutputDescription() foundation.INSDictionary {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("outputDescription"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (n MLNeuralNetworkContainer) SetOutputDescription(value foundation.INSDictionary) {
	objc.Send[struct{}](n.ID, objc.Sel("setOutputDescription:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/outputLayerNames
func (n MLNeuralNetworkContainer) OutputLayerNames() foundation.INSArray {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("outputLayerNames"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (n MLNeuralNetworkContainer) SetOutputLayerNames(value foundation.INSArray) {
	objc.Send[struct{}](n.ID, objc.Sel("setOutputLayerNames:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/precision
func (n MLNeuralNetworkContainer) Precision() int {
	rv := objc.Send[int](n.ID, objc.Sel("precision"))
	return rv
}
func (n MLNeuralNetworkContainer) SetPrecision(value int) {
	objc.Send[struct{}](n.ID, objc.Sel("setPrecision:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLNeuralNetworkContainer/updatableModelCompiledParams
func (n MLNeuralNetworkContainer) UpdatableModelCompiledParams() IMLNeuralNetworksCompileTimeParams {
	rv := objc.Send[objc.ID](n.ID, objc.Sel("updatableModelCompiledParams"))
	return MLNeuralNetworksCompileTimeParamsFromID(objc.ID(rv))
}
func (n MLNeuralNetworkContainer) SetUpdatableModelCompiledParams(value IMLNeuralNetworksCompileTimeParams) {
	objc.Send[struct{}](n.ID, objc.Sel("setUpdatableModelCompiledParams:"), value)
}

