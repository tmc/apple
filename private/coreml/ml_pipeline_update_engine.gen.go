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

// The class instance for the [MLPipelineUpdateEngine] class.
var (
	_MLPipelineUpdateEngineClass     MLPipelineUpdateEngineClass
	_MLPipelineUpdateEngineClassOnce sync.Once
)

func getMLPipelineUpdateEngineClass() MLPipelineUpdateEngineClass {
	_MLPipelineUpdateEngineClassOnce.Do(func() {
		_MLPipelineUpdateEngineClass = MLPipelineUpdateEngineClass{class: objc.GetClass("MLPipelineUpdateEngine")}
	})
	return _MLPipelineUpdateEngineClass
}

// GetMLPipelineUpdateEngineClass returns the class object for MLPipelineUpdateEngine.
func GetMLPipelineUpdateEngineClass() MLPipelineUpdateEngineClass {
	return getMLPipelineUpdateEngineClass()
}

type MLPipelineUpdateEngineClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MLPipelineUpdateEngineClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MLPipelineUpdateEngineClass) Alloc() MLPipelineUpdateEngine {
	rv := objc.Send[MLPipelineUpdateEngine](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [MLPipelineUpdateEngine.CancelUpdate]
//   - [MLPipelineUpdateEngine.DispatchQueue]
//   - [MLPipelineUpdateEngine.SetDispatchQueue]
//   - [MLPipelineUpdateEngine.ParameterValueForKeyError]
//   - [MLPipelineUpdateEngine.ProgressHandlers]
//   - [MLPipelineUpdateEngine.SetProgressHandlers]
//   - [MLPipelineUpdateEngine.ResumeUpdate]
//   - [MLPipelineUpdateEngine.ResumeUpdateWithParameters]
//   - [MLPipelineUpdateEngine.SetUpdateProgressHandlersDispatchQueue]
//   - [MLPipelineUpdateEngine.UpdatableModelIndicies]
//   - [MLPipelineUpdateEngine.UpdateModelWithData]
//   - [MLPipelineUpdateEngine.WriteToURLError]
//   - [MLPipelineUpdateEngine.InitFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError]
//   - [MLPipelineUpdateEngine.Configuration]
//   - [MLPipelineUpdateEngine.DebugDescription]
//   - [MLPipelineUpdateEngine.Description]
//   - [MLPipelineUpdateEngine.Hash]
//   - [MLPipelineUpdateEngine.Metadata]
//   - [MLPipelineUpdateEngine.ModelDescription]
//   - [MLPipelineUpdateEngine.PredictionTypeForKTrace]
//   - [MLPipelineUpdateEngine.RecordsPredictionEvent]
//   - [MLPipelineUpdateEngine.SignpostID]
//   - [MLPipelineUpdateEngine.Superclass]
//   - [MLPipelineUpdateEngine.SupportsConcurrentSubmissions]
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine
type MLPipelineUpdateEngine struct {
	objectivec.Object
}

// MLPipelineUpdateEngineFromID constructs a [MLPipelineUpdateEngine] from an objc.ID.
func MLPipelineUpdateEngineFromID(id objc.ID) MLPipelineUpdateEngine {
	return MLPipelineUpdateEngine{objectivec.Object{ID: id}}
}

// NOTE: MLPipelineUpdateEngine struct embeds objectivec.Object (parent type unavailable) but
// IMLPipelineUpdateEngine embeds the parent interface; skip compile-time assertion.

// An interface definition for the [MLPipelineUpdateEngine] class.
//
// # Methods
//
//   - [IMLPipelineUpdateEngine.CancelUpdate]
//   - [IMLPipelineUpdateEngine.DispatchQueue]
//   - [IMLPipelineUpdateEngine.SetDispatchQueue]
//   - [IMLPipelineUpdateEngine.ParameterValueForKeyError]
//   - [IMLPipelineUpdateEngine.ProgressHandlers]
//   - [IMLPipelineUpdateEngine.SetProgressHandlers]
//   - [IMLPipelineUpdateEngine.ResumeUpdate]
//   - [IMLPipelineUpdateEngine.ResumeUpdateWithParameters]
//   - [IMLPipelineUpdateEngine.SetUpdateProgressHandlersDispatchQueue]
//   - [IMLPipelineUpdateEngine.UpdatableModelIndicies]
//   - [IMLPipelineUpdateEngine.UpdateModelWithData]
//   - [IMLPipelineUpdateEngine.WriteToURLError]
//   - [IMLPipelineUpdateEngine.InitFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError]
//   - [IMLPipelineUpdateEngine.Configuration]
//   - [IMLPipelineUpdateEngine.DebugDescription]
//   - [IMLPipelineUpdateEngine.Description]
//   - [IMLPipelineUpdateEngine.Hash]
//   - [IMLPipelineUpdateEngine.Metadata]
//   - [IMLPipelineUpdateEngine.ModelDescription]
//   - [IMLPipelineUpdateEngine.PredictionTypeForKTrace]
//   - [IMLPipelineUpdateEngine.RecordsPredictionEvent]
//   - [IMLPipelineUpdateEngine.SignpostID]
//   - [IMLPipelineUpdateEngine.Superclass]
//   - [IMLPipelineUpdateEngine.SupportsConcurrentSubmissions]
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine
type IMLPipelineUpdateEngine interface {
	IMLPipeline

	// Topic: Methods

	CancelUpdate()
	DispatchQueue() objectivec.Object
	SetDispatchQueue(value objectivec.Object)
	ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error)
	ProgressHandlers() IMLUpdateProgressHandlers
	SetProgressHandlers(value IMLUpdateProgressHandlers)
	ResumeUpdate()
	ResumeUpdateWithParameters(parameters objectivec.IObject)
	SetUpdateProgressHandlersDispatchQueue(handlers ErrorHandler, queue objectivec.IObject)
	UpdatableModelIndicies() objectivec.IObject
	UpdateModelWithData(data objectivec.IObject)
	WriteToURLError(url foundation.INSURL) (bool, error)
	InitFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (MLPipelineUpdateEngine, error)
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
func (m MLPipelineUpdateEngine) Init() MLPipelineUpdateEngine {
	rv := objc.Send[MLPipelineUpdateEngine](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MLPipelineUpdateEngine) Autorelease() MLPipelineUpdateEngine {
	rv := objc.Send[MLPipelineUpdateEngine](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPipelineUpdateEngine creates a new MLPipelineUpdateEngine instance.
func NewMLPipelineUpdateEngine() MLPipelineUpdateEngine {
	class := getMLPipelineUpdateEngineClass()
	rv := objc.Send[MLPipelineUpdateEngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/initFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:
func NewPipelineUpdateEngineFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (MLPipelineUpdateEngine, error) {
	var errorPtr objc.ID
	instance := getMLPipelineUpdateEngineClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLPipelineUpdateEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLPipelineUpdateEngineFromID(rv), nil
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/cancelUpdate
func (m MLPipelineUpdateEngine) CancelUpdate() {
	objc.Send[objc.ID](m.ID, objc.Sel("cancelUpdate"))
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/parameterValueForKey:error:
func (m MLPipelineUpdateEngine) ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("parameterValueForKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/resumeUpdate
func (m MLPipelineUpdateEngine) ResumeUpdate() {
	objc.Send[objc.ID](m.ID, objc.Sel("resumeUpdate"))
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/resumeUpdateWithParameters:
func (m MLPipelineUpdateEngine) ResumeUpdateWithParameters(parameters objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("resumeUpdateWithParameters:"), parameters)
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/setUpdateProgressHandlers:dispatchQueue:
func (m MLPipelineUpdateEngine) SetUpdateProgressHandlersDispatchQueue(handlers ErrorHandler, queue objectivec.IObject) {
	_block0, _ := NewErrorBlock(handlers)
	objc.Send[objc.ID](m.ID, objc.Sel("setUpdateProgressHandlers:dispatchQueue:"), _block0, queue)
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/updateModelWithData:
func (m MLPipelineUpdateEngine) UpdateModelWithData(data objectivec.IObject) {
	objc.Send[objc.ID](m.ID, objc.Sel("updateModelWithData:"), data)
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/writeToURL:error:
func (m MLPipelineUpdateEngine) WriteToURLError(url foundation.INSURL) (bool, error) {
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

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/initFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:
func (m MLPipelineUpdateEngine) InitFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (MLPipelineUpdateEngine, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](m.ID, objc.Sel("initFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLPipelineUpdateEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLPipelineUpdateEngineFromID(rv), nil

}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:
func (_MLPipelineUpdateEngineClass MLPipelineUpdateEngineClass) LoadModelFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_MLPipelineUpdateEngineClass.class), objc.Sel("loadModelFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/configuration
func (m MLPipelineUpdateEngine) Configuration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/debugDescription
func (m MLPipelineUpdateEngine) DebugDescription() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/description
func (m MLPipelineUpdateEngine) Description() string {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/dispatchQueue
func (m MLPipelineUpdateEngine) DispatchQueue() objectivec.Object {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("dispatchQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (m MLPipelineUpdateEngine) SetDispatchQueue(value objectivec.Object) {
	objc.Send[struct{}](m.ID, objc.Sel("setDispatchQueue:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/hash
func (m MLPipelineUpdateEngine) Hash() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("hash"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/metadata
func (m MLPipelineUpdateEngine) Metadata() IMLModelMetadata {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("metadata"))
	return MLModelMetadataFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/modelDescription
func (m MLPipelineUpdateEngine) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/predictionTypeForKTrace
func (m MLPipelineUpdateEngine) PredictionTypeForKTrace() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("predictionTypeForKTrace"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/progressHandlers
func (m MLPipelineUpdateEngine) ProgressHandlers() IMLUpdateProgressHandlers {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("progressHandlers"))
	return MLUpdateProgressHandlersFromID(objc.ID(rv))
}
func (m MLPipelineUpdateEngine) SetProgressHandlers(value IMLUpdateProgressHandlers) {
	objc.Send[struct{}](m.ID, objc.Sel("setProgressHandlers:"), value)
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/recordsPredictionEvent
func (m MLPipelineUpdateEngine) RecordsPredictionEvent() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("recordsPredictionEvent"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/signpostID
func (m MLPipelineUpdateEngine) SignpostID() uint64 {
	rv := objc.Send[uint64](m.ID, objc.Sel("signpostID"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/superclass
func (m MLPipelineUpdateEngine) Superclass() objc.Class {
	rv := objc.Send[objc.Class](m.ID, objc.Sel("superclass"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/supportsConcurrentSubmissions
func (m MLPipelineUpdateEngine) SupportsConcurrentSubmissions() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("supportsConcurrentSubmissions"))
	return rv
}

// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/updatableModelIndicies
func (m MLPipelineUpdateEngine) UpdatableModelIndicies() objectivec.IObject {
	rv := objc.Send[objc.ID](m.ID, objc.Sel("updatableModelIndicies"))
	return objectivec.Object{ID: rv}
}
