// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEInMemoryModel] class.
var (
	_ANEInMemoryModelClass     ANEInMemoryModelClass
	_ANEInMemoryModelClassOnce sync.Once
)

func getANEInMemoryModelClass() ANEInMemoryModelClass {
	_ANEInMemoryModelClassOnce.Do(func() {
		_ANEInMemoryModelClass = ANEInMemoryModelClass{class: objc.GetClass("_ANEInMemoryModel")}
	})
	return _ANEInMemoryModelClass
}

// GetANEInMemoryModelClass returns the class object for _ANEInMemoryModel.
func GetANEInMemoryModelClass() ANEInMemoryModelClass {
	return getANEInMemoryModelClass()
}

type ANEInMemoryModelClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEInMemoryModelClass) Alloc() ANEInMemoryModel {
	rv := objc.Send[ANEInMemoryModel](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ANEInMemoryModel.CompileWithQoSOptionsError]
//   - [ANEInMemoryModel.CompiledModelExists]
//   - [ANEInMemoryModel.CompilerOptionsFileName]
//   - [ANEInMemoryModel.SetCompilerOptionsFileName]
//   - [ANEInMemoryModel.CompilerOptionsWithOptionsIsCompiledModelCached]
//   - [ANEInMemoryModel.Descriptor]
//   - [ANEInMemoryModel.SetDescriptor]
//   - [ANEInMemoryModel.EvaluateWithQoSOptionsRequestError]
//   - [ANEInMemoryModel.HexStringIdentifier]
//   - [ANEInMemoryModel.IntermediateBufferHandle]
//   - [ANEInMemoryModel.SetIntermediateBufferHandle]
//   - [ANEInMemoryModel.IsMILModel]
//   - [ANEInMemoryModel.LoadWithQoSOptionsError]
//   - [ANEInMemoryModel.LocalModelPath]
//   - [ANEInMemoryModel.MapIOSurfacesWithRequestCacheInferenceError]
//   - [ANEInMemoryModel.Model]
//   - [ANEInMemoryModel.SetModel]
//   - [ANEInMemoryModel.ModelAttributes]
//   - [ANEInMemoryModel.SetModelAttributes]
//   - [ANEInMemoryModel.ModelURL]
//   - [ANEInMemoryModel.SetModelURL]
//   - [ANEInMemoryModel.PerfStatsMask]
//   - [ANEInMemoryModel.SetPerfStatsMask]
//   - [ANEInMemoryModel.Program]
//   - [ANEInMemoryModel.SetProgram]
//   - [ANEInMemoryModel.ProgramHandle]
//   - [ANEInMemoryModel.SetProgramHandle]
//   - [ANEInMemoryModel.PurgeCompiledModel]
//   - [ANEInMemoryModel.QueueDepth]
//   - [ANEInMemoryModel.SetQueueDepth]
//   - [ANEInMemoryModel.SaveModelFiles]
//   - [ANEInMemoryModel.SharedConnection]
//   - [ANEInMemoryModel.State]
//   - [ANEInMemoryModel.SetState]
//   - [ANEInMemoryModel.String_id]
//   - [ANEInMemoryModel.UnloadWithQoSError]
//   - [ANEInMemoryModel.UnmapIOSurfacesWithRequest]
//   - [ANEInMemoryModel.InitWithDesctiptor]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel
type ANEInMemoryModel struct {
	objectivec.Object
}

// ANEInMemoryModelFromID constructs a [ANEInMemoryModel] from an objc.ID.
func ANEInMemoryModelFromID(id objc.ID) ANEInMemoryModel {
	return ANEInMemoryModel{objectivec.Object{id}}
}
// Ensure ANEInMemoryModel implements IANEInMemoryModel.
var _ IANEInMemoryModel = ANEInMemoryModel{}

// An interface definition for the [ANEInMemoryModel] class.
//
// # Methods
//
//   - [IANEInMemoryModel.CompileWithQoSOptionsError]
//   - [IANEInMemoryModel.CompiledModelExists]
//   - [IANEInMemoryModel.CompilerOptionsFileName]
//   - [IANEInMemoryModel.SetCompilerOptionsFileName]
//   - [IANEInMemoryModel.CompilerOptionsWithOptionsIsCompiledModelCached]
//   - [IANEInMemoryModel.Descriptor]
//   - [IANEInMemoryModel.SetDescriptor]
//   - [IANEInMemoryModel.EvaluateWithQoSOptionsRequestError]
//   - [IANEInMemoryModel.HexStringIdentifier]
//   - [IANEInMemoryModel.IntermediateBufferHandle]
//   - [IANEInMemoryModel.SetIntermediateBufferHandle]
//   - [IANEInMemoryModel.IsMILModel]
//   - [IANEInMemoryModel.LoadWithQoSOptionsError]
//   - [IANEInMemoryModel.LocalModelPath]
//   - [IANEInMemoryModel.MapIOSurfacesWithRequestCacheInferenceError]
//   - [IANEInMemoryModel.Model]
//   - [IANEInMemoryModel.SetModel]
//   - [IANEInMemoryModel.ModelAttributes]
//   - [IANEInMemoryModel.SetModelAttributes]
//   - [IANEInMemoryModel.ModelURL]
//   - [IANEInMemoryModel.SetModelURL]
//   - [IANEInMemoryModel.PerfStatsMask]
//   - [IANEInMemoryModel.SetPerfStatsMask]
//   - [IANEInMemoryModel.Program]
//   - [IANEInMemoryModel.SetProgram]
//   - [IANEInMemoryModel.ProgramHandle]
//   - [IANEInMemoryModel.SetProgramHandle]
//   - [IANEInMemoryModel.PurgeCompiledModel]
//   - [IANEInMemoryModel.QueueDepth]
//   - [IANEInMemoryModel.SetQueueDepth]
//   - [IANEInMemoryModel.SaveModelFiles]
//   - [IANEInMemoryModel.SharedConnection]
//   - [IANEInMemoryModel.State]
//   - [IANEInMemoryModel.SetState]
//   - [IANEInMemoryModel.String_id]
//   - [IANEInMemoryModel.UnloadWithQoSError]
//   - [IANEInMemoryModel.UnmapIOSurfacesWithRequest]
//   - [IANEInMemoryModel.InitWithDesctiptor]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel
type IANEInMemoryModel interface {
	objectivec.IObject

	// Topic: Methods

	CompileWithQoSOptionsError(s uint32, options objectivec.IObject) (bool, error)
	CompiledModelExists() bool
	CompilerOptionsFileName() string
	SetCompilerOptionsFileName(value string)
	CompilerOptionsWithOptionsIsCompiledModelCached(options objectivec.IObject, cached bool) objectivec.IObject
	Descriptor() *ANEInMemoryModelDescriptor
	SetDescriptor(value *ANEInMemoryModelDescriptor)
	EvaluateWithQoSOptionsRequestError(s uint32, options objectivec.IObject, request objectivec.IObject) (bool, error)
	HexStringIdentifier() string
	IntermediateBufferHandle() uint64
	SetIntermediateBufferHandle(value uint64)
	IsMILModel() bool
	LoadWithQoSOptionsError(s uint32, options objectivec.IObject) (bool, error)
	LocalModelPath() objectivec.IObject
	MapIOSurfacesWithRequestCacheInferenceError(request objectivec.IObject, inference bool) (bool, error)
	Model() *ANEModel
	SetModel(value *ANEModel)
	ModelAttributes() foundation.INSDictionary
	SetModelAttributes(value foundation.INSDictionary)
	ModelURL() foundation.INSURL
	SetModelURL(value foundation.INSURL)
	PerfStatsMask() uint32
	SetPerfStatsMask(value uint32)
	Program() *ANEProgramForEvaluation
	SetProgram(value *ANEProgramForEvaluation)
	ProgramHandle() uint64
	SetProgramHandle(value uint64)
	PurgeCompiledModel()
	QueueDepth() int8
	SetQueueDepth(value int8)
	SaveModelFiles() objectivec.IObject
	SharedConnection() *ANEClient
	State() uint64
	SetState(value uint64)
	String_id() uint64
	UnloadWithQoSError(s uint32) (bool, error)
	UnmapIOSurfacesWithRequest(request objectivec.IObject)
	InitWithDesctiptor(desctiptor objectivec.IObject) ANEInMemoryModel
}

// Init initializes the instance.
func (a ANEInMemoryModel) Init() ANEInMemoryModel {
	rv := objc.Send[ANEInMemoryModel](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEInMemoryModel) Autorelease() ANEInMemoryModel {
	rv := objc.Send[ANEInMemoryModel](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEInMemoryModel creates a new ANEInMemoryModel instance.
func NewANEInMemoryModel() ANEInMemoryModel {
	class := getANEInMemoryModelClass()
	rv := objc.Send[ANEInMemoryModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/initWithDesctiptor:
func NewANEInMemoryModelWithDesctiptor(desctiptor objectivec.IObject) ANEInMemoryModel {
	instance := getANEInMemoryModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDesctiptor:"), desctiptor)
	return ANEInMemoryModelFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/compileWithQoS:options:error:
func (a ANEInMemoryModel) CompileWithQoSOptionsError(s uint32, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("compileWithQoS:options:error:"), s, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("compileWithQoS:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/compiledModelExists
func (a ANEInMemoryModel) CompiledModelExists() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("compiledModelExists"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/compilerOptionsWithOptions:isCompiledModelCached:
func (a ANEInMemoryModel) CompilerOptionsWithOptionsIsCompiledModelCached(options objectivec.IObject, cached bool) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("compilerOptionsWithOptions:isCompiledModelCached:"), options, cached)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/evaluateWithQoS:options:request:error:
func (a ANEInMemoryModel) EvaluateWithQoSOptionsRequestError(s uint32, options objectivec.IObject, request objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("evaluateWithQoS:options:request:error:"), s, options, request, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("evaluateWithQoS:options:request:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/loadWithQoS:options:error:
func (a ANEInMemoryModel) LoadWithQoSOptionsError(s uint32, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("loadWithQoS:options:error:"), s, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("loadWithQoS:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/localModelPath
func (a ANEInMemoryModel) LocalModelPath() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("localModelPath"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/mapIOSurfacesWithRequest:cacheInference:error:
func (a ANEInMemoryModel) MapIOSurfacesWithRequestCacheInferenceError(request objectivec.IObject, inference bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("mapIOSurfacesWithRequest:cacheInference:error:"), request, inference, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("mapIOSurfacesWithRequest:cacheInference:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/purgeCompiledModel
func (a ANEInMemoryModel) PurgeCompiledModel() {
	objc.Send[objc.ID](a.ID, objc.Sel("purgeCompiledModel"))
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/saveModelFiles
func (a ANEInMemoryModel) SaveModelFiles() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("saveModelFiles"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/unloadWithQoS:error:
func (a ANEInMemoryModel) UnloadWithQoSError(s uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("unloadWithQoS:error:"), s, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("unloadWithQoS:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/unmapIOSurfacesWithRequest:
func (a ANEInMemoryModel) UnmapIOSurfacesWithRequest(request objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("unmapIOSurfacesWithRequest:"), request)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/initWithDesctiptor:
func (a ANEInMemoryModel) InitWithDesctiptor(desctiptor objectivec.IObject) ANEInMemoryModel {
	rv := objc.Send[ANEInMemoryModel](a.ID, objc.Sel("initWithDesctiptor:"), desctiptor)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/inMemoryModelWithDescriptor:
func (_ANEInMemoryModelClass ANEInMemoryModelClass) InMemoryModelWithDescriptor(descriptor objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEInMemoryModelClass.class), objc.Sel("inMemoryModelWithDescriptor:"), descriptor)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/compilerOptionsFileName
func (a ANEInMemoryModel) CompilerOptionsFileName() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("compilerOptionsFileName"))
	return foundation.NSStringFromID(rv).String()
}
func (a ANEInMemoryModel) SetCompilerOptionsFileName(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setCompilerOptionsFileName:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/descriptor
func (a ANEInMemoryModel) Descriptor() *ANEInMemoryModelDescriptor {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("descriptor"))
	if rv == 0 {
		return nil
	}
	val := ANEInMemoryModelDescriptorFromID(objc.ID(rv))
	return &val
}
func (a ANEInMemoryModel) SetDescriptor(value *ANEInMemoryModelDescriptor) {
	if value == nil {
		objc.Send[struct{}](a.ID, objc.Sel("setDescriptor:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](a.ID, objc.Sel("setDescriptor:"), value)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/hexStringIdentifier
func (a ANEInMemoryModel) HexStringIdentifier() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("hexStringIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/intermediateBufferHandle
func (a ANEInMemoryModel) IntermediateBufferHandle() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("intermediateBufferHandle"))
	return rv
}
func (a ANEInMemoryModel) SetIntermediateBufferHandle(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setIntermediateBufferHandle:"), value)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/isMILModel
func (a ANEInMemoryModel) IsMILModel() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isMILModel"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/model
func (a ANEInMemoryModel) Model() *ANEModel {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("model"))
	if rv == 0 {
		return nil
	}
	val := ANEModelFromID(objc.ID(rv))
	return &val
}
func (a ANEInMemoryModel) SetModel(value *ANEModel) {
	if value == nil {
		objc.Send[struct{}](a.ID, objc.Sel("setModel:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](a.ID, objc.Sel("setModel:"), value)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/modelAttributes
func (a ANEInMemoryModel) ModelAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("modelAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (a ANEInMemoryModel) SetModelAttributes(value foundation.INSDictionary) {
	objc.Send[struct{}](a.ID, objc.Sel("setModelAttributes:"), value)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/modelURL
func (a ANEInMemoryModel) ModelURL() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("modelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (a ANEInMemoryModel) SetModelURL(value foundation.INSURL) {
	objc.Send[struct{}](a.ID, objc.Sel("setModelURL:"), value)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/perfStatsMask
func (a ANEInMemoryModel) PerfStatsMask() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("perfStatsMask"))
	return rv
}
func (a ANEInMemoryModel) SetPerfStatsMask(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setPerfStatsMask:"), value)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/program
func (a ANEInMemoryModel) Program() *ANEProgramForEvaluation {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("program"))
	if rv == 0 {
		return nil
	}
	val := ANEProgramForEvaluationFromID(objc.ID(rv))
	return &val
}
func (a ANEInMemoryModel) SetProgram(value *ANEProgramForEvaluation) {
	if value == nil {
		objc.Send[struct{}](a.ID, objc.Sel("setProgram:"), objc.ID(0))
		return
	}
	objc.Send[struct{}](a.ID, objc.Sel("setProgram:"), value)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/programHandle
func (a ANEInMemoryModel) ProgramHandle() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("programHandle"))
	return rv
}
func (a ANEInMemoryModel) SetProgramHandle(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setProgramHandle:"), value)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/queueDepth
func (a ANEInMemoryModel) QueueDepth() int8 {
	rv := objc.Send[int8](a.ID, objc.Sel("queueDepth"))
	return rv
}
func (a ANEInMemoryModel) SetQueueDepth(value int8) {
	objc.Send[struct{}](a.ID, objc.Sel("setQueueDepth:"), value)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/sharedConnection
func (a ANEInMemoryModel) SharedConnection() *ANEClient {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sharedConnection"))
	if rv == 0 {
		return nil
	}
	val := ANEClientFromID(objc.ID(rv))
	return &val
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/state
func (a ANEInMemoryModel) State() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("state"))
	return rv
}
func (a ANEInMemoryModel) SetState(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setState:"), value)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInMemoryModel/string_id
func (a ANEInMemoryModel) String_id() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("string_id"))
	return rv
}

