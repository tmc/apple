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

//
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
func (p MLPipelineUpdateEngine) Init() MLPipelineUpdateEngine {
	rv := objc.Send[MLPipelineUpdateEngine](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p MLPipelineUpdateEngine) Autorelease() MLPipelineUpdateEngine {
	rv := objc.Send[MLPipelineUpdateEngine](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewMLPipelineUpdateEngine creates a new MLPipelineUpdateEngine instance.
func NewMLPipelineUpdateEngine() MLPipelineUpdateEngine {
	class := getMLPipelineUpdateEngineClass()
	rv := objc.Send[MLPipelineUpdateEngine](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
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
func (p MLPipelineUpdateEngine) CancelUpdate() {
	objc.Send[objc.ID](p.ID, objc.Sel("cancelUpdate"))
}
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/parameterValueForKey:error:
func (p MLPipelineUpdateEngine) ParameterValueForKeyError(key objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](p.ID, objc.Sel("parameterValueForKey:error:"), key, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/resumeUpdate
func (p MLPipelineUpdateEngine) ResumeUpdate() {
	objc.Send[objc.ID](p.ID, objc.Sel("resumeUpdate"))
}
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/resumeUpdateWithParameters:
func (p MLPipelineUpdateEngine) ResumeUpdateWithParameters(parameters objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("resumeUpdateWithParameters:"), parameters)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/setUpdateProgressHandlers:dispatchQueue:
func (p MLPipelineUpdateEngine) SetUpdateProgressHandlersDispatchQueue(handlers ErrorHandler, queue objectivec.IObject) {
_block0, _ := NewErrorBlock(handlers)
	objc.Send[objc.ID](p.ID, objc.Sel("setUpdateProgressHandlers:dispatchQueue:"), _block0, queue)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/updateModelWithData:
func (p MLPipelineUpdateEngine) UpdateModelWithData(data objectivec.IObject) {
	objc.Send[objc.ID](p.ID, objc.Sel("updateModelWithData:"), data)
}
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/writeToURL:error:
func (p MLPipelineUpdateEngine) WriteToURLError(url foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](p.ID, objc.Sel("writeToURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeToURL:error: returned NO with nil NSError")
	}
	return rv, nil

}
//
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/initFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:
func (p MLPipelineUpdateEngine) InitFromCompiledArchiveModelVersionInfoCompilerVersionInfoConfigurationError(archive unsafe.Pointer, info objectivec.IObject, info2 objectivec.IObject, configuration objectivec.IObject) (MLPipelineUpdateEngine, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](p.ID, objc.Sel("initFromCompiledArchive:modelVersionInfo:compilerVersionInfo:configuration:error:"), archive, info, info2, configuration, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return MLPipelineUpdateEngine{}, foundation.NSErrorFrom(errorPtr)
	}
	return MLPipelineUpdateEngineFromID(rv), nil

}

//
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
func (p MLPipelineUpdateEngine) Configuration() IMLModelConfiguration {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("configuration"))
	return MLModelConfigurationFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/debugDescription
func (p MLPipelineUpdateEngine) DebugDescription() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/description
func (p MLPipelineUpdateEngine) Description() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/dispatchQueue
func (p MLPipelineUpdateEngine) DispatchQueue() objectivec.Object {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("dispatchQueue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
func (p MLPipelineUpdateEngine) SetDispatchQueue(value objectivec.Object) {
	objc.Send[struct{}](p.ID, objc.Sel("setDispatchQueue:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/hash
func (p MLPipelineUpdateEngine) Hash() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("hash"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/metadata
func (p MLPipelineUpdateEngine) Metadata() IMLModelMetadata {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("metadata"))
	return MLModelMetadataFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/modelDescription
func (p MLPipelineUpdateEngine) ModelDescription() IMLModelDescription {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("modelDescription"))
	return MLModelDescriptionFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/predictionTypeForKTrace
func (p MLPipelineUpdateEngine) PredictionTypeForKTrace() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("predictionTypeForKTrace"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/progressHandlers
func (p MLPipelineUpdateEngine) ProgressHandlers() IMLUpdateProgressHandlers {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("progressHandlers"))
	return MLUpdateProgressHandlersFromID(objc.ID(rv))
}
func (p MLPipelineUpdateEngine) SetProgressHandlers(value IMLUpdateProgressHandlers) {
	objc.Send[struct{}](p.ID, objc.Sel("setProgressHandlers:"), value)
}
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/recordsPredictionEvent
func (p MLPipelineUpdateEngine) RecordsPredictionEvent() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("recordsPredictionEvent"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/signpostID
func (p MLPipelineUpdateEngine) SignpostID() uint64 {
	rv := objc.Send[uint64](p.ID, objc.Sel("signpostID"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/superclass
func (p MLPipelineUpdateEngine) Superclass() objc.Class {
	rv := objc.Send[objc.Class](p.ID, objc.Sel("superclass"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/supportsConcurrentSubmissions
func (p MLPipelineUpdateEngine) SupportsConcurrentSubmissions() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("supportsConcurrentSubmissions"))
	return rv
}
// See: https://developer.apple.com/documentation/CoreML/MLPipelineUpdateEngine/updatableModelIndicies
func (p MLPipelineUpdateEngine) UpdatableModelIndicies() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("updatableModelIndicies"))
	return objectivec.Object{ID: rv}
}

