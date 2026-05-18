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

// The class instance for the [MLTreeEnsembleXGBoostUpdateEngine] class.
var (
	_MLTreeEnsembleXGBoostUpdateEngineClass     MLTreeEnsembleXGBoostUpdateEngineClass
	_MLTreeEnsembleXGBoostUpdateEngineClassOnce sync.Once
)

func getMLTreeEnsembleXGBoostUpdateEngineClass() MLTreeEnsembleXGBoostUpdateEngineClass {
	_MLTreeEnsembleXGBoostUpdateEngineClassOnce.Do(func() {
		_MLTreeEnsembleXGBoostUpdateEngineClass = MLTreeEnsembleXGBoostUpdateEngineClass{class: objc.GetClass("MLTreeEnsembleXGBoostUpdateEngine")}
	})
	return _MLTreeEnsembleXGBoostUpdateEngineClass
}

// GetMLTreeEnsembleXGBoostUpdateEngineClass returns the class object for MLTreeEnsembleXGBoostUpdateEngine.
func GetMLTreeEnsembleXGBoostUpdateEngineClass() MLTreeEnsembleXGBoostUpdateEngineClass {
	return getMLTreeEnsembleXGBoostUpdateEngineClass()
}

type MLTreeEnsembleXGBoostUpdateEngineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLTreeEnsembleXGBoostUpdateEngineClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLTreeEnsembleXGBoostUpdateEngineClass) Alloc() MLTreeEnsembleXGBoostUpdateEngine {
	rv := objc.Send[MLTreeEnsembleXGBoostUpdateEngine](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLTreeEnsembleXGBoostUpdateEngine.CachedModel]
//   - [MLTreeEnsembleXGBoostUpdateEngine.SetCachedModel]
//   - [MLTreeEnsembleXGBoostUpdateEngine.CancelUpdate]
//   - [MLTreeEnsembleXGBoostUpdateEngine.ClassesByInt]
//   - [MLTreeEnsembleXGBoostUpdateEngine.SetClassesByInt]
//   - [MLTreeEnsembleXGBoostUpdateEngine.ClassesByString]
//   - [MLTreeEnsembleXGBoostUpdateEngine.SetClassesByString]
//   - [MLTreeEnsembleXGBoostUpdateEngine.ContinueWithUpdate]
//   - [MLTreeEnsembleXGBoostUpdateEngine.SetContinueWithUpdate]
//   - [MLTreeEnsembleXGBoostUpdateEngine.LoadParameterDescriptionsAndContainerFromConfigurationModelDescriptionError]
//   - [MLTreeEnsembleXGBoostUpdateEngine.MmappedModel]
//   - [MLTreeEnsembleXGBoostUpdateEngine.SetMmappedModel]
//   - [MLTreeEnsembleXGBoostUpdateEngine.NumDimensions]
//   - [MLTreeEnsembleXGBoostUpdateEngine.SetNumDimensions]
//   - [MLTreeEnsembleXGBoostUpdateEngine.ParameterContainer]
//   - [MLTreeEnsembleXGBoostUpdateEngine.SetParameterContainer]
//   - [MLTreeEnsembleXGBoostUpdateEngine.ParameterValueForKey]
//   - [MLTreeEnsembleXGBoostUpdateEngine.ParameterValueForKeyError]
//   - [MLTreeEnsembleXGBoostUpdateEngine.Personalization]
//   - [MLTreeEnsembleXGBoostUpdateEngine.SetPersonalization]
//   - [MLTreeEnsembleXGBoostUpdateEngine.ProgressHandlers]
//   - [MLTreeEnsembleXGBoostUpdateEngine.SetProgressHandlers]
//   - [MLTreeEnsembleXGBoostUpdateEngine.ProgressHandlersDispatchQueue]
//   - [MLTreeEnsembleXGBoostUpdateEngine.SetProgressHandlersDispatchQueue]
//   - [MLTreeEnsembleXGBoostUpdateEngine.ResumeUpdate]
//   - [MLTreeEnsembleXGBoostUpdateEngine.ResumeUpdateWithParameters]
//   - [MLTreeEnsembleXGBoostUpdateEngine.SetBoosterParametersError]
//   - [MLTreeEnsembleXGBoostUpdateEngine.SetUpdateProgressHandlersDispatchQueue]
//   - [MLTreeEnsembleXGBoostUpdateEngine.UpdateModelWithData]
//   - [MLTreeEnsembleXGBoostUpdateEngine.UpdateParameters]
//   - [MLTreeEnsembleXGBoostUpdateEngine.WriteToURLError]
//   - [MLTreeEnsembleXGBoostUpdateEngine.InitWithCompiledArchiveConfigurationError]
//   - [MLTreeEnsembleXGBoostUpdateEngine.Configuration]
//   - [MLTreeEnsembleXGBoostUpdateEngine.DebugDescription]
//   - [MLTreeEnsembleXGBoostUpdateEngine.Description]
//   - [MLTreeEnsembleXGBoostUpdateEngine.Hash]
//   - [MLTreeEnsembleXGBoostUpdateEngine.Metadata]
//   - [MLTreeEnsembleXGBoostUpdateEngine.ModelDescription]
//   - [MLTreeEnsembleXGBoostUpdateEngine.PredictionTypeForKTrace]
//   - [MLTreeEnsembleXGBoostUpdateEngine.RecordsPredictionEvent]
//   - [MLTreeEnsembleXGBoostUpdateEngine.SignpostID]
//   - [MLTreeEnsembleXGBoostUpdateEngine.Superclass]
//   - [MLTreeEnsembleXGBoostUpdateEngine.SupportsConcurrentSubmissions]
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine
type MLTreeEnsembleXGBoostUpdateEngine struct {
	MLTreeEnsembleXGBoostClassifier
}

// MLTreeEnsembleXGBoostUpdateEngineFromID constructs a [MLTreeEnsembleXGBoostUpdateEngine] from an objc.ID.
func MLTreeEnsembleXGBoostUpdateEngineFromID(id objc.ID) MLTreeEnsembleXGBoostUpdateEngine {
	return MLTreeEnsembleXGBoostUpdateEngine{MLTreeEnsembleXGBoostClassifier: MLTreeEnsembleXGBoostClassifierFromID(id)}
}

// Ensure MLTreeEnsembleXGBoostUpdateEngine implements IMLTreeEnsembleXGBoostUpdateEngine.
var _ IMLTreeEnsembleXGBoostUpdateEngine = MLTreeEnsembleXGBoostUpdateEngine{}

// An interface definition for the [MLTreeEnsembleXGBoostUpdateEngine] class.
//
// # Methods
//
//   - [IMLTreeEnsembleXGBoostUpdateEngine.CachedModel]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.SetCachedModel]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.CancelUpdate]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.ClassesByInt]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.SetClassesByInt]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.ClassesByString]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.SetClassesByString]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.ContinueWithUpdate]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.SetContinueWithUpdate]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.LoadParameterDescriptionsAndContainerFromConfigurationModelDescriptionError]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.MmappedModel]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.SetMmappedModel]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.NumDimensions]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.SetNumDimensions]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.ParameterContainer]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.SetParameterContainer]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.ParameterValueForKey]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.ParameterValueForKeyError]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.Personalization]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.SetPersonalization]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.ProgressHandlers]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.SetProgressHandlers]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.ProgressHandlersDispatchQueue]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.SetProgressHandlersDispatchQueue]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.ResumeUpdate]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.ResumeUpdateWithParameters]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.SetBoosterParametersError]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.SetUpdateProgressHandlersDispatchQueue]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.UpdateModelWithData]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.UpdateParameters]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.WriteToURLError]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.InitWithCompiledArchiveConfigurationError]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.Configuration]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.DebugDescription]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.Description]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.Hash]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.Metadata]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.ModelDescription]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.PredictionTypeForKTrace]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.RecordsPredictionEvent]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.SignpostID]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.Superclass]
//   - [IMLTreeEnsembleXGBoostUpdateEngine.SupportsConcurrentSubmissions]
//
// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine
type IMLTreeEnsembleXGBoostUpdateEngine interface {
	IMLTreeEnsembleXGBoostClassifier

	// Topic: Methods

	CachedModel() objectivec.IObject
	SetCachedModel(value objectivec.IObject)
	CancelUpdate()
	ClassesByInt() objectivec.IObject
	SetClassesByInt(value objectivec.IObject)
	ClassesByString() objectivec.IObject
	SetClassesByString(value objectivec.IObject)
	ContinueWithUpdate() bool
	SetContinueWithUpdate(value bool)
	LoadParameterDescriptionsAndContainerFromConfigurationModelDescriptionError(configuration objectivec.IObject, description objectivec.IObject) (objectivec.IObject, error)
	MmappedModel() objectivec.IObject
	SetMmappedModel(value objectivec.IObject)
	NumDimensions() uint64
	SetNumDimensions(value uint64)
	ParameterContainer() IMLParameterContainer
	SetParameterContainer(value IMLParameterContainer)
	ParameterValueForKey(key objectivec.IObject) objectivec.IObject
	ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error)
	Personalization() bool
	SetPersonalization(value bool)
	ProgressHandlers() IMLUpdateProgressHandlers
	SetProgressHandlers(value IMLUpdateProgressHandlers)
	ProgressHandlersDispatchQueue() objectivec.Object
	SetProgressHandlersDispatchQueue(value objectivec.Object)
	ResumeUpdate()
	ResumeUpdateWithParameters(parameters objectivec.IObject)
	SetBoosterParametersError(parameters unsafe.Pointer) (bool, error)
	SetUpdateProgressHandlersDispatchQueue(handlers ErrorHandler, queue objectivec.IObject)
	UpdateModelWithData(data objectivec.IObject)
	UpdateParameters() objectivec.IObject
	WriteToURLError(url foundation.INSURL) (bool, error)
	InitWithCompiledArchiveConfigurationError(archive unsafe.Pointer, configuration objectivec.IObject) (MLTreeEnsembleXGBoostUpdateEngine, error)
	Configuration() IMLModelConfiguration
	DebugDescription() string
	Description() string
	Hash() uint64
	Metadata() IMLModelMetadata
	ModelDescription() IMLModelDescription
	PredictionTypeForKTrace() uint64
	RecordsPredictionEvent() bool
	SignpostID() uint64
	Superclass() objc.Class
	SupportsConcurrentSubmissions() bool
}

// Init initializes the instance.
func (m MLTreeEnsembleXGBoostUpdateEngine) Init() MLTreeEnsembleXGBoostUpdateEngine {
	rv := objc.Send[MLTreeEnsembleXGBoostUpdateEngine](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLTreeEnsembleXGBoostUpdateEngine) Autorelease() MLTreeEnsembleXGBoostUpdateEngine {
	rv := objc.Send[MLTreeEnsembleXGBoostUpdateEngine](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLTreeEnsembleXGBoostUpdateEngine creates a new MLTreeEnsembleXGBoostUpdateEngine instance.
func NewMLTreeEnsembleXGBoostUpdateEngine() MLTreeEnsembleXGBoostUpdateEngine {
	class := getMLTreeEnsembleXGBoostUpdateEngineClass()
	rv := objc.Send[MLTreeEnsembleXGBoostUpdateEngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/initWithCompiledArchive:configuration:error:
func NewTreeEnsembleXGBoostUpdateEngineWithCompiledArchiveConfigurationError(archive unsafe.Pointer, configuration objectivec.IObject) (MLTreeEnsembleXGBoostUpdateEngine, error) {
	var errorPtr objc.ID
	instance := getMLTreeEnsembleXGBoostUpdateEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCompiledArchive:configuration:error:"), archive, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLTreeEnsembleXGBoostUpdateEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLTreeEnsembleXGBoostUpdateEngineFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostClassifier/initWithDescription:configuration:indexToStringLabelArray:indexToIntLabelArray:modelURL:error:
func NewTreeEnsembleXGBoostUpdateEngineWithDescriptionConfigurationIndexToStringLabelArrayIndexToIntLabelArrayModelURLError(description objectivec.IObject, configuration objectivec.IObject, array objectivec.IObject, array2 objectivec.IObject, url foundation.INSURL) (MLTreeEnsembleXGBoostUpdateEngine, error) {
	var errorPtr objc.ID
	instance := getMLTreeEnsembleXGBoostUpdateEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:indexToStringLabelArray:indexToIntLabelArray:modelURL:error:"), description, configuration, array, array2, url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLTreeEnsembleXGBoostUpdateEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLTreeEnsembleXGBoostUpdateEngineFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/cancelUpdate
func (m MLTreeEnsembleXGBoostUpdateEngine) CancelUpdate() {
	objc.Send[objc.ID](m.ID, objc.Sel("cancelUpdate"))
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/loadParameterDescriptionsAndContainerFromConfiguration:modelDescription:error:
func (m MLTreeEnsembleXGBoostUpdateEngine) LoadParameterDescriptionsAndContainerFromConfigurationModelDescriptionError(configuration objectivec.IObject, description objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("loadParameterDescriptionsAndContainerFromConfiguration:modelDescription:error:"), configuration, description, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/parameterValueForKey:
func (m MLTreeEnsembleXGBoostUpdateEngine) ParameterValueForKey(key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parameterValueForKey:"), key)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/parameterValueForKey:error:
func (m MLTreeEnsembleXGBoostUpdateEngine) ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parameterValueForKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/resumeUpdate
func (m MLTreeEnsembleXGBoostUpdateEngine) ResumeUpdate() {
	objc.Send[objc.ID](m.ID, objc.Sel("resumeUpdate"))
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/resumeUpdateWithParameters:
func (m MLTreeEnsembleXGBoostUpdateEngine) ResumeUpdateWithParameters(parameters objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("resumeUpdateWithParameters:"), parameters)
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/setBoosterParameters:error:
func (m MLTreeEnsembleXGBoostUpdateEngine) SetBoosterParametersError(parameters unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](m.ID, objc.Sel("setBoosterParameters:error:"), parameters, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setBoosterParameters:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/setUpdateProgressHandlers:dispatchQueue:
func (m MLTreeEnsembleXGBoostUpdateEngine) SetUpdateProgressHandlersDispatchQueue(handlers ErrorHandler, queue objectivec.IObject) {
	_block0, _ := NewErrorBlock(handlers)
	objc.Send[objc.ID](m.ID, objc.Sel("setUpdateProgressHandlers:dispatchQueue:"), _block0, queue)
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/updateModelWithData:
func (m MLTreeEnsembleXGBoostUpdateEngine) UpdateModelWithData(data objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("updateModelWithData:"), data)
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/updateParameters
func (m MLTreeEnsembleXGBoostUpdateEngine) UpdateParameters() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("updateParameters"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/writeToURL:error:
func (m MLTreeEnsembleXGBoostUpdateEngine) WriteToURLError(url foundation.INSURL) (bool, error) {
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

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/initWithCompiledArchive:configuration:error:
func (m MLTreeEnsembleXGBoostUpdateEngine) InitWithCompiledArchiveConfigurationError(archive unsafe.Pointer, configuration objectivec.IObject) (MLTreeEnsembleXGBoostUpdateEngine, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithCompiledArchive:configuration:error:"), archive, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLTreeEnsembleXGBoostUpdateEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLTreeEnsembleXGBoostUpdateEngineFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:
func (_MLTreeEnsembleXGBoostUpdateEngineClass MLTreeEnsembleXGBoostUpdateEngineClass) LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLTreeEnsembleXGBoostUpdateEngineClass.class), objc.Sel("loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/cachedModel
func (m MLTreeEnsembleXGBoostUpdateEngine) CachedModel() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("cachedModel"))
	return objectivec.Object{ID: rv}
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SetCachedModel(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setCachedModel:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/classesByInt
func (m MLTreeEnsembleXGBoostUpdateEngine) ClassesByInt() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classesByInt"))
	return objectivec.Object{ID: rv}
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SetClassesByInt(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setClassesByInt:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/classesByString
func (m MLTreeEnsembleXGBoostUpdateEngine) ClassesByString() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("classesByString"))
	return objectivec.Object{ID: rv}
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SetClassesByString(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setClassesByString:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/configuration
func (m MLTreeEnsembleXGBoostUpdateEngine) Configuration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/continueWithUpdate
func (m MLTreeEnsembleXGBoostUpdateEngine) ContinueWithUpdate() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("continueWithUpdate"))
	return rv
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SetContinueWithUpdate(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setContinueWithUpdate:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/debugDescription
func (m MLTreeEnsembleXGBoostUpdateEngine) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/description
func (m MLTreeEnsembleXGBoostUpdateEngine) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/hash
func (m MLTreeEnsembleXGBoostUpdateEngine) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/metadata
func (m MLTreeEnsembleXGBoostUpdateEngine) Metadata() IMLModelMetadata {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("metadata"))
	return MLModelMetadataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/mmappedModel
func (m MLTreeEnsembleXGBoostUpdateEngine) MmappedModel() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("mmappedModel"))
	return objectivec.Object{ID: rv}
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SetMmappedModel(value objectivec.IObject) {
	objc.Send[struct{}](m.ID, objc.Sel("setMmappedModel:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/modelDescription
func (m MLTreeEnsembleXGBoostUpdateEngine) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/numDimensions
func (m MLTreeEnsembleXGBoostUpdateEngine) NumDimensions() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("numDimensions"))
	return rv
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SetNumDimensions(value uint64) {
	objc.Send[struct{}](m.ID, objc.Sel("setNumDimensions:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/parameterContainer
func (m MLTreeEnsembleXGBoostUpdateEngine) ParameterContainer() IMLParameterContainer {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parameterContainer"))
	return MLParameterContainerFromID(objc.ID(rv))
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SetParameterContainer(value IMLParameterContainer) {
	objc.Send[struct{}](m.ID, objc.Sel("setParameterContainer:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/personalization
func (m MLTreeEnsembleXGBoostUpdateEngine) Personalization() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("personalization"))
	return rv
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SetPersonalization(value bool) {
	objc.Send[struct{}](m.ID, objc.Sel("setPersonalization:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/predictionTypeForKTrace
func (m MLTreeEnsembleXGBoostUpdateEngine) PredictionTypeForKTrace() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("predictionTypeForKTrace"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/progressHandlers
func (m MLTreeEnsembleXGBoostUpdateEngine) ProgressHandlers() IMLUpdateProgressHandlers {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("progressHandlers"))
	return MLUpdateProgressHandlersFromID(objc.ID(rv))
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SetProgressHandlers(value IMLUpdateProgressHandlers) {
	objc.Send[struct{}](m.ID, objc.Sel("setProgressHandlers:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/progressHandlersDispatchQueue
func (m MLTreeEnsembleXGBoostUpdateEngine) ProgressHandlersDispatchQueue() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("progressHandlersDispatchQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SetProgressHandlersDispatchQueue(value objectivec.Object) {
	objc.Send[struct{}](m.ID, objc.Sel("setProgressHandlersDispatchQueue:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/recordsPredictionEvent
func (m MLTreeEnsembleXGBoostUpdateEngine) RecordsPredictionEvent() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("recordsPredictionEvent"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/signpostID
func (m MLTreeEnsembleXGBoostUpdateEngine) SignpostID() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("signpostID"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/superclass
func (m MLTreeEnsembleXGBoostUpdateEngine) Superclass() objc.Class {
	rv := objc.Send[objc.Class](m.ID, objc.Sel("superclass"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLTreeEnsembleXGBoostUpdateEngine/supportsConcurrentSubmissions
func (m MLTreeEnsembleXGBoostUpdateEngine) SupportsConcurrentSubmissions() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportsConcurrentSubmissions"))
	return rv
}
