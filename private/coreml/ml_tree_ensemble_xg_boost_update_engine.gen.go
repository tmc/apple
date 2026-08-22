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
	rv := objc.SendIfResponds[MLTreeEnsembleXGBoostUpdateEngine](objc.ID(mc.class), objc.Sel("alloc"))
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
type IMLTreeEnsembleXGBoostUpdateEngine interface {
	IMLTreeEnsembleXGBoostClassifier

	// Topic: Methods

	CachedModel() unsafe.Pointer
	SetCachedModel(value unsafe.Pointer)
	CancelUpdate()
	ClassesByInt() unsafe.Pointer
	SetClassesByInt(value unsafe.Pointer)
	ClassesByString() unsafe.Pointer
	SetClassesByString(value unsafe.Pointer)
	ContinueWithUpdate() bool
	SetContinueWithUpdate(value bool)
	LoadParameterDescriptionsAndContainerFromConfigurationModelDescriptionError(configuration objectivec.IObject, description objectivec.IObject) (objectivec.IObject, error)
	MmappedModel() unsafe.Pointer
	SetMmappedModel(value unsafe.Pointer)
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
	SetUpdateProgressHandlersDispatchQueue(handlers objectivec.IObject, queue objectivec.IObject)
	UpdateModelWithData(data objectivec.IObject)
	UpdateParameters() objectivec.IObject
	WriteToURLError(url foundation.NSURL) (bool, error)
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
	Superclass() objectivec.Class
	SupportsConcurrentSubmissions() bool
}

// Init initializes the instance.
func (m MLTreeEnsembleXGBoostUpdateEngine) Init() MLTreeEnsembleXGBoostUpdateEngine {
	rv := objc.SendIfResponds[MLTreeEnsembleXGBoostUpdateEngine](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLTreeEnsembleXGBoostUpdateEngine) Autorelease() MLTreeEnsembleXGBoostUpdateEngine {
	rv := objc.SendIfResponds[MLTreeEnsembleXGBoostUpdateEngine](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLTreeEnsembleXGBoostUpdateEngine creates a new MLTreeEnsembleXGBoostUpdateEngine instance.
func NewMLTreeEnsembleXGBoostUpdateEngine() MLTreeEnsembleXGBoostUpdateEngine {
	class := getMLTreeEnsembleXGBoostUpdateEngineClass()
	rv := objc.SendIfResponds[MLTreeEnsembleXGBoostUpdateEngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewTreeEnsembleXGBoostUpdateEngineWithCompiledArchiveConfigurationError(archive unsafe.Pointer, configuration objectivec.IObject) (MLTreeEnsembleXGBoostUpdateEngine, error) {
	var errorPtr objc.ID
	instance := getMLTreeEnsembleXGBoostUpdateEngineClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCompiledArchive:configuration:error:"), archive, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLTreeEnsembleXGBoostUpdateEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLTreeEnsembleXGBoostUpdateEngine{}, objc.ErrInitFailed
	}
	return MLTreeEnsembleXGBoostUpdateEngineFromID(rv), nil
}

func NewTreeEnsembleXGBoostUpdateEngineWithDescriptionConfigurationIndexToStringLabelArrayIndexToIntLabelArrayModelURLError(description objectivec.IObject, configuration objectivec.IObject, array unsafe.Pointer, array2 unsafe.Pointer, url foundation.NSURL) (MLTreeEnsembleXGBoostUpdateEngine, error) {
	var errorPtr objc.ID
	instance := getMLTreeEnsembleXGBoostUpdateEngineClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDescription:configuration:indexToStringLabelArray:indexToIntLabelArray:modelURL:error:"), description, configuration, array, array2, url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLTreeEnsembleXGBoostUpdateEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	if rv == 0 {
		return MLTreeEnsembleXGBoostUpdateEngine{}, objc.ErrInitFailed
	}
	return MLTreeEnsembleXGBoostUpdateEngineFromID(rv), nil
}

func (m MLTreeEnsembleXGBoostUpdateEngine) CancelUpdate() {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("cancelUpdate"))
}
func (m MLTreeEnsembleXGBoostUpdateEngine) LoadParameterDescriptionsAndContainerFromConfigurationModelDescriptionError(configuration objectivec.IObject, description objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("loadParameterDescriptionsAndContainerFromConfiguration:modelDescription:error:"), configuration, description, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLTreeEnsembleXGBoostUpdateEngine) ParameterValueForKey(key objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("parameterValueForKey:"), key)
	return objectivec.Object{ID: rv}
}
func (m MLTreeEnsembleXGBoostUpdateEngine) ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parameterValueForKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
func (m MLTreeEnsembleXGBoostUpdateEngine) ResumeUpdate() {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("resumeUpdate"))
}
func (m MLTreeEnsembleXGBoostUpdateEngine) ResumeUpdateWithParameters(parameters objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("resumeUpdateWithParameters:"), parameters)
}
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
func (m MLTreeEnsembleXGBoostUpdateEngine) SetUpdateProgressHandlersDispatchQueue(handlers objectivec.IObject, queue objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("setUpdateProgressHandlers:dispatchQueue:"), handlers, queue)
}
func (m MLTreeEnsembleXGBoostUpdateEngine) UpdateModelWithData(data objectivec.IObject) {
	objc.SendIfResponds[objc.ID](m.ID, objc.Sel("updateModelWithData:"), data)
}
func (m MLTreeEnsembleXGBoostUpdateEngine) UpdateParameters() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("updateParameters"))
	return objectivec.Object{ID: rv}
}
func (m MLTreeEnsembleXGBoostUpdateEngine) WriteToURLError(url foundation.NSURL) (bool, error) {
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
func (m MLTreeEnsembleXGBoostUpdateEngine) InitWithCompiledArchiveConfigurationError(archive unsafe.Pointer, configuration objectivec.IObject) (MLTreeEnsembleXGBoostUpdateEngine, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initWithCompiledArchive:configuration:error:"), archive, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLTreeEnsembleXGBoostUpdateEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLTreeEnsembleXGBoostUpdateEngineFromID(rv), nil

}

func (_MLTreeEnsembleXGBoostUpdateEngineClass MLTreeEnsembleXGBoostUpdateEngineClass) LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLTreeEnsembleXGBoostUpdateEngineClass.class), objc.Sel("loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

func (m MLTreeEnsembleXGBoostUpdateEngine) CachedModel() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("cachedModel"))
	return rv
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SetCachedModel(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setCachedModel:"), value)
}
func (m MLTreeEnsembleXGBoostUpdateEngine) ClassesByInt() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("classesByInt"))
	return rv
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SetClassesByInt(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setClassesByInt:"), value)
}
func (m MLTreeEnsembleXGBoostUpdateEngine) ClassesByString() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("classesByString"))
	return rv
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SetClassesByString(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setClassesByString:"), value)
}
func (m MLTreeEnsembleXGBoostUpdateEngine) Configuration() IMLModelConfiguration {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
func (m MLTreeEnsembleXGBoostUpdateEngine) ContinueWithUpdate() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("continueWithUpdate"))
	return rv
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SetContinueWithUpdate(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setContinueWithUpdate:"), value)
}
func (m MLTreeEnsembleXGBoostUpdateEngine) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLTreeEnsembleXGBoostUpdateEngine) Description() string {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (m MLTreeEnsembleXGBoostUpdateEngine) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("hash"))
	return rv
}
func (m MLTreeEnsembleXGBoostUpdateEngine) Metadata() IMLModelMetadata {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("metadata"))
	return MLModelMetadataFromID(objc.ID(rv))
}
func (m MLTreeEnsembleXGBoostUpdateEngine) MmappedModel() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](m.ID, objc.Sel("mmappedModel"))
	return rv
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SetMmappedModel(value unsafe.Pointer) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setMmappedModel:"), value)
}
func (m MLTreeEnsembleXGBoostUpdateEngine) ModelDescription() IMLModelDescription {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}
func (m MLTreeEnsembleXGBoostUpdateEngine) NumDimensions() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("numDimensions"))
	return rv
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SetNumDimensions(value uint64) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setNumDimensions:"), value)
}
func (m MLTreeEnsembleXGBoostUpdateEngine) ParameterContainer() IMLParameterContainer {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("parameterContainer"))
	return MLParameterContainerFromID(objc.ID(rv))
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SetParameterContainer(value IMLParameterContainer) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setParameterContainer:"), value)
}
func (m MLTreeEnsembleXGBoostUpdateEngine) Personalization() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("personalization"))
	return rv
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SetPersonalization(value bool) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setPersonalization:"), value)
}
func (m MLTreeEnsembleXGBoostUpdateEngine) PredictionTypeForKTrace() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("predictionTypeForKTrace"))
	return rv
}
func (m MLTreeEnsembleXGBoostUpdateEngine) ProgressHandlers() IMLUpdateProgressHandlers {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("progressHandlers"))
	return MLUpdateProgressHandlersFromID(objc.ID(rv))
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SetProgressHandlers(value IMLUpdateProgressHandlers) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setProgressHandlers:"), value)
}
func (m MLTreeEnsembleXGBoostUpdateEngine) ProgressHandlersDispatchQueue() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](m.ID, objc.Sel("progressHandlersDispatchQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SetProgressHandlersDispatchQueue(value objectivec.Object) {
	objc.SendIfResponds[struct{}](m.ID, objc.Sel("setProgressHandlersDispatchQueue:"), value)
}
func (m MLTreeEnsembleXGBoostUpdateEngine) RecordsPredictionEvent() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("recordsPredictionEvent"))
	return rv
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SignpostID() uint64 {
	rv := objc.SendIfResponds[uint64](m.ID, objc.Sel("signpostID"))
	return rv
}
func (m MLTreeEnsembleXGBoostUpdateEngine) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](m.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
func (m MLTreeEnsembleXGBoostUpdateEngine) SupportsConcurrentSubmissions() bool {
	rv := objc.SendIfResponds[bool](m.ID, objc.Sel("supportsConcurrentSubmissions"))
	return rv
}
