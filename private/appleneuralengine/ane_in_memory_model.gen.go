// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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

// Class returns the underlying Objective-C class pointer.
func (ac ANEInMemoryModelClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEInMemoryModelClass) Alloc() ANEInMemoryModel {
	rv := objc.SendIfResponds[ANEInMemoryModel](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

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
type ANEInMemoryModel struct {
	objectivec.Object
}

// ANEInMemoryModelFromID constructs a [ANEInMemoryModel] from an objc.ID.
func ANEInMemoryModelFromID(id objc.ID) ANEInMemoryModel {
	return ANEInMemoryModel{objectivec.Object{ID: id}}
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
type IANEInMemoryModel interface {
	objectivec.IObject

	// Topic: Methods

	CompileWithQoSOptionsError(s uint32, options objectivec.IObject) (bool, error)
	CompiledModelExists() bool
	CompilerOptionsFileName() string
	SetCompilerOptionsFileName(value string)
	CompilerOptionsWithOptionsIsCompiledModelCached(options objectivec.IObject, cached bool) objectivec.IObject
	Descriptor() IANEInMemoryModelDescriptor
	SetDescriptor(value IANEInMemoryModelDescriptor)
	EvaluateWithQoSOptionsRequestError(s uint32, options objectivec.IObject, request objectivec.IObject) (bool, error)
	HexStringIdentifier() string
	IntermediateBufferHandle() uint64
	SetIntermediateBufferHandle(value uint64)
	IsMILModel() bool
	LoadWithQoSOptionsError(s uint32, options objectivec.IObject) (bool, error)
	LocalModelPath() objectivec.IObject
	MapIOSurfacesWithRequestCacheInferenceError(request objectivec.IObject, inference bool) (bool, error)
	Model() IANEModel
	SetModel(value IANEModel)
	ModelAttributes() foundation.INSDictionary
	SetModelAttributes(value foundation.INSDictionary)
	ModelURL() foundation.NSURL
	SetModelURL(value foundation.NSURL)
	PerfStatsMask() uint32
	SetPerfStatsMask(value uint32)
	Program() IANEProgramForEvaluation
	SetProgram(value IANEProgramForEvaluation)
	ProgramHandle() uint64
	SetProgramHandle(value uint64)
	PurgeCompiledModel()
	QueueDepth() int8
	SetQueueDepth(value int8)
	SaveModelFiles() objectivec.IObject
	SharedConnection() IANEClient
	State() uint64
	SetState(value uint64)
	String_id() uint64
	UnloadWithQoSError(s uint32) (bool, error)
	UnmapIOSurfacesWithRequest(request objectivec.IObject)
	InitWithDesctiptor(desctiptor objectivec.IObject) ANEInMemoryModel
}

// Init initializes the instance.
func (a ANEInMemoryModel) Init() ANEInMemoryModel {
	rv := objc.SendIfResponds[ANEInMemoryModel](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEInMemoryModel) Autorelease() ANEInMemoryModel {
	rv := objc.SendIfResponds[ANEInMemoryModel](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEInMemoryModel creates a new ANEInMemoryModel instance.
func NewANEInMemoryModel() ANEInMemoryModel {
	class := getANEInMemoryModelClass()
	rv := objc.SendIfResponds[ANEInMemoryModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewANEInMemoryModelWithDesctiptor(desctiptor objectivec.IObject) ANEInMemoryModel {
	instance := getANEInMemoryModelClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDesctiptor:"), desctiptor)
	return ANEInMemoryModelFromID(rv)
}

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
func (a ANEInMemoryModel) CompiledModelExists() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("compiledModelExists"))
	return rv
}
func (a ANEInMemoryModel) CompilerOptionsWithOptionsIsCompiledModelCached(options objectivec.IObject, cached bool) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("compilerOptionsWithOptions:isCompiledModelCached:"), options, cached)
	return objectivec.Object{ID: rv}
}
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
func (a ANEInMemoryModel) LocalModelPath() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("localModelPath"))
	return objectivec.Object{ID: rv}
}
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
func (a ANEInMemoryModel) PurgeCompiledModel() {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("purgeCompiledModel"))
}
func (a ANEInMemoryModel) SaveModelFiles() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("saveModelFiles"))
	return objectivec.Object{ID: rv}
}
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
func (a ANEInMemoryModel) UnmapIOSurfacesWithRequest(request objectivec.IObject) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("unmapIOSurfacesWithRequest:"), request)
}
func (a ANEInMemoryModel) InitWithDesctiptor(desctiptor objectivec.IObject) ANEInMemoryModel {
	rv := objc.SendIfResponds[ANEInMemoryModel](a.ID, objc.Sel("initWithDesctiptor:"), desctiptor)
	return rv
}

func (_ANEInMemoryModelClass ANEInMemoryModelClass) InMemoryModelWithDescriptor(descriptor objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEInMemoryModelClass.class), objc.Sel("inMemoryModelWithDescriptor:"), descriptor)
	return objectivec.Object{ID: rv}
}

func (a ANEInMemoryModel) CompilerOptionsFileName() string {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("compilerOptionsFileName"))
	return foundation.NSStringFromID(rv).String()
}
func (a ANEInMemoryModel) SetCompilerOptionsFileName(value string) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setCompilerOptionsFileName:"), objc.String(value))
}
func (a ANEInMemoryModel) Descriptor() IANEInMemoryModelDescriptor {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("descriptor"))
	return ANEInMemoryModelDescriptorFromID(objc.ID(rv))
}
func (a ANEInMemoryModel) SetDescriptor(value IANEInMemoryModelDescriptor) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setDescriptor:"), value)
}
func (a ANEInMemoryModel) HexStringIdentifier() string {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("hexStringIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (a ANEInMemoryModel) IntermediateBufferHandle() uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("intermediateBufferHandle"))
	return rv
}
func (a ANEInMemoryModel) SetIntermediateBufferHandle(value uint64) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setIntermediateBufferHandle:"), value)
}
func (a ANEInMemoryModel) IsMILModel() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("isMILModel"))
	return rv
}
func (a ANEInMemoryModel) Model() IANEModel {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("model"))
	return ANEModelFromID(objc.ID(rv))
}
func (a ANEInMemoryModel) SetModel(value IANEModel) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setModel:"), value)
}
func (a ANEInMemoryModel) ModelAttributes() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("modelAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (a ANEInMemoryModel) SetModelAttributes(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setModelAttributes:"), value)
}
func (a ANEInMemoryModel) ModelURL() foundation.NSURL {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("modelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (a ANEInMemoryModel) SetModelURL(value foundation.NSURL) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setModelURL:"), value)
}
func (a ANEInMemoryModel) PerfStatsMask() uint32 {
	rv := objc.SendIfResponds[uint32](a.ID, objc.Sel("perfStatsMask"))
	return rv
}
func (a ANEInMemoryModel) SetPerfStatsMask(value uint32) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setPerfStatsMask:"), value)
}
func (a ANEInMemoryModel) Program() IANEProgramForEvaluation {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("program"))
	return ANEProgramForEvaluationFromID(objc.ID(rv))
}
func (a ANEInMemoryModel) SetProgram(value IANEProgramForEvaluation) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setProgram:"), value)
}
func (a ANEInMemoryModel) ProgramHandle() uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("programHandle"))
	return rv
}
func (a ANEInMemoryModel) SetProgramHandle(value uint64) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setProgramHandle:"), value)
}
func (a ANEInMemoryModel) QueueDepth() int8 {
	rv := objc.SendIfResponds[int8](a.ID, objc.Sel("queueDepth"))
	return rv
}
func (a ANEInMemoryModel) SetQueueDepth(value int8) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setQueueDepth:"), value)
}
func (a ANEInMemoryModel) SharedConnection() IANEClient {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("sharedConnection"))
	return ANEClientFromID(objc.ID(rv))
}
func (a ANEInMemoryModel) State() uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("state"))
	return rv
}
func (a ANEInMemoryModel) SetState(value uint64) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setState:"), value)
}
func (a ANEInMemoryModel) String_id() uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("string_id"))
	return rv
}
