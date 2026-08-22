// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/os"
)

// The class instance for the [ANEModel] class.
var (
	_ANEModelClass     ANEModelClass
	_ANEModelClassOnce sync.Once
)

func getANEModelClass() ANEModelClass {
	_ANEModelClassOnce.Do(func() {
		_ANEModelClass = ANEModelClass{class: objc.GetClass("_ANEModel")}
	})
	return _ANEModelClass
}

// GetANEModelClass returns the class object for _ANEModel.
func GetANEModelClass() ANEModelClass {
	return getANEModelClass()
}

type ANEModelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEModelClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEModelClass) Alloc() ANEModel {
	rv := objc.SendIfResponds[ANEModel](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ANEModel.UUID]
//   - [ANEModel.CacheURLIdentifier]
//   - [ANEModel.SetCacheURLIdentifier]
//   - [ANEModel.EncodeWithCoder]
//   - [ANEModel.GetCacheURLIdentifier]
//   - [ANEModel.GetUUID]
//   - [ANEModel.IdentifierSource]
//   - [ANEModel.InputSymbolIndicesForProcedureIndex]
//   - [ANEModel.IntermediateBufferHandle]
//   - [ANEModel.SetIntermediateBufferHandle]
//   - [ANEModel.IsEqualToModel]
//   - [ANEModel.Key]
//   - [ANEModel.L]
//   - [ANEModel.SetL]
//   - [ANEModel.Mapper]
//   - [ANEModel.SetMapper]
//   - [ANEModel.ModelAttributes]
//   - [ANEModel.SetModelAttributes]
//   - [ANEModel.ModelURL]
//   - [ANEModel.MpsConstants]
//   - [ANEModel.OutputSymbolIndicesForProcedureIndex]
//   - [ANEModel.PerfStatsMask]
//   - [ANEModel.SetPerfStatsMask]
//   - [ANEModel.ProcedureInfoForProcedureIndex]
//   - [ANEModel.Program]
//   - [ANEModel.SetProgram]
//   - [ANEModel.ProgramHandle]
//   - [ANEModel.SetProgramHandle]
//   - [ANEModel.QueueDepth]
//   - [ANEModel.SetQueueDepth]
//   - [ANEModel.ResetOnUnload]
//   - [ANEModel.ShallowCopy]
//   - [ANEModel.SourceURL]
//   - [ANEModel.State]
//   - [ANEModel.SetState]
//   - [ANEModel.String_id]
//   - [ANEModel.SetString_id]
//   - [ANEModel.SymbolIndicesForProcedureIndexIndexArrayKey]
//   - [ANEModel.UpdateModelAttributesState]
//   - [ANEModel.UpdateModelAttributesStateProgramHandleIntermediateBufferHandleQueueDepth]
//   - [ANEModel.InitWithCoder]
//   - [ANEModel.InitWithModelAtURLKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURL]
//   - [ANEModel.InitWithModelAtURLSourceURLUUIDKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURLString_idGenerateNewStringId]
//   - [ANEModel.InitWithModelAtURLSourceURLUUIDKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURLString_idGenerateNewStringIdMpsConstants]
//   - [ANEModel.InitWithModelIdentifier]
type ANEModel struct {
	objectivec.Object
}

// ANEModelFromID constructs a [ANEModel] from an objc.ID.
func ANEModelFromID(id objc.ID) ANEModel {
	return ANEModel{objectivec.Object{ID: id}}
}

// Ensure ANEModel implements IANEModel.
var _ IANEModel = ANEModel{}

// An interface definition for the [ANEModel] class.
//
// # Methods
//
//   - [IANEModel.UUID]
//   - [IANEModel.CacheURLIdentifier]
//   - [IANEModel.SetCacheURLIdentifier]
//   - [IANEModel.EncodeWithCoder]
//   - [IANEModel.GetCacheURLIdentifier]
//   - [IANEModel.GetUUID]
//   - [IANEModel.IdentifierSource]
//   - [IANEModel.InputSymbolIndicesForProcedureIndex]
//   - [IANEModel.IntermediateBufferHandle]
//   - [IANEModel.SetIntermediateBufferHandle]
//   - [IANEModel.IsEqualToModel]
//   - [IANEModel.Key]
//   - [IANEModel.L]
//   - [IANEModel.SetL]
//   - [IANEModel.Mapper]
//   - [IANEModel.SetMapper]
//   - [IANEModel.ModelAttributes]
//   - [IANEModel.SetModelAttributes]
//   - [IANEModel.ModelURL]
//   - [IANEModel.MpsConstants]
//   - [IANEModel.OutputSymbolIndicesForProcedureIndex]
//   - [IANEModel.PerfStatsMask]
//   - [IANEModel.SetPerfStatsMask]
//   - [IANEModel.ProcedureInfoForProcedureIndex]
//   - [IANEModel.Program]
//   - [IANEModel.SetProgram]
//   - [IANEModel.ProgramHandle]
//   - [IANEModel.SetProgramHandle]
//   - [IANEModel.QueueDepth]
//   - [IANEModel.SetQueueDepth]
//   - [IANEModel.ResetOnUnload]
//   - [IANEModel.ShallowCopy]
//   - [IANEModel.SourceURL]
//   - [IANEModel.State]
//   - [IANEModel.SetState]
//   - [IANEModel.String_id]
//   - [IANEModel.SetString_id]
//   - [IANEModel.SymbolIndicesForProcedureIndexIndexArrayKey]
//   - [IANEModel.UpdateModelAttributesState]
//   - [IANEModel.UpdateModelAttributesStateProgramHandleIntermediateBufferHandleQueueDepth]
//   - [IANEModel.InitWithCoder]
//   - [IANEModel.InitWithModelAtURLKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURL]
//   - [IANEModel.InitWithModelAtURLSourceURLUUIDKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURLString_idGenerateNewStringId]
//   - [IANEModel.InitWithModelAtURLSourceURLUUIDKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURLString_idGenerateNewStringIdMpsConstants]
//   - [IANEModel.InitWithModelIdentifier]
type IANEModel interface {
	objectivec.IObject

	// Topic: Methods

	UUID() foundation.NSUUID
	CacheURLIdentifier() string
	SetCacheURLIdentifier(value string)
	EncodeWithCoder(coder foundation.INSCoder)
	GetCacheURLIdentifier() objectivec.IObject
	GetUUID() objectivec.IObject
	IdentifierSource() int64
	InputSymbolIndicesForProcedureIndex(index uint32) objectivec.IObject
	IntermediateBufferHandle() uint64
	SetIntermediateBufferHandle(value uint64)
	IsEqualToModel(model objectivec.IObject) bool
	Key() string
	L() os.OSUnfairLockS
	SetL(value os.OSUnfairLockS)
	Mapper() IANEProgramIOSurfacesMapper
	SetMapper(value IANEProgramIOSurfacesMapper)
	ModelAttributes() foundation.INSDictionary
	SetModelAttributes(value foundation.INSDictionary)
	ModelURL() foundation.NSURL
	MpsConstants() foundation.INSDictionary
	OutputSymbolIndicesForProcedureIndex(index uint32) objectivec.IObject
	PerfStatsMask() uint32
	SetPerfStatsMask(value uint32)
	ProcedureInfoForProcedureIndex(index uint32) objectivec.IObject
	Program() IANEProgramForEvaluation
	SetProgram(value IANEProgramForEvaluation)
	ProgramHandle() uint64
	SetProgramHandle(value uint64)
	QueueDepth() int8
	SetQueueDepth(value int8)
	ResetOnUnload()
	ShallowCopy() objectivec.IObject
	SourceURL() foundation.NSURL
	State() uint64
	SetState(value uint64)
	String_id() uint64
	SetString_id(value uint64)
	SymbolIndicesForProcedureIndexIndexArrayKey(index uint32, key objectivec.IObject) objectivec.IObject
	UpdateModelAttributesState(attributes objectivec.IObject, state uint64)
	UpdateModelAttributesStateProgramHandleIntermediateBufferHandleQueueDepth(attributes objectivec.IObject, state uint64, handle uint64, handle2 uint64, depth int8)
	InitWithCoder(coder foundation.INSCoder) ANEModel
	InitWithModelAtURLKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURL(url foundation.NSURL, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject, attributes objectivec.IObject, url2 bool) ANEModel
	InitWithModelAtURLSourceURLUUIDKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURLString_idGenerateNewStringId(url foundation.NSURL, url2 foundation.NSURL, uid objectivec.IObject, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject, attributes objectivec.IObject, url3 bool, string_id uint64, id bool) ANEModel
	InitWithModelAtURLSourceURLUUIDKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURLString_idGenerateNewStringIdMpsConstants(url foundation.NSURL, url2 foundation.NSURL, uid objectivec.IObject, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject, attributes objectivec.IObject, url3 bool, string_id uint64, id bool, constants objectivec.IObject) ANEModel
	InitWithModelIdentifier(identifier objectivec.IObject) ANEModel
}

// Init initializes the instance.
func (a ANEModel) Init() ANEModel {
	rv := objc.SendIfResponds[ANEModel](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEModel) Autorelease() ANEModel {
	rv := objc.SendIfResponds[ANEModel](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEModel creates a new ANEModel instance.
func NewANEModel() ANEModel {
	class := getANEModelClass()
	rv := objc.SendIfResponds[ANEModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewANEModelWithCoder(coder objectivec.IObject) ANEModel {
	instance := getANEModelClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return ANEModelFromID(rv)
}

func NewANEModelWithModelAtURLKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURL(url foundation.NSURL, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject, attributes objectivec.IObject, url2 bool) ANEModel {
	instance := getANEModelClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithModelAtURL:key:identifierSource:cacheURLIdentifier:modelAttributes:standardizeURL:"), url, key, source, uRLIdentifier, attributes, url2)
	return ANEModelFromID(rv)
}

func NewANEModelWithModelAtURLSourceURLUUIDKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURLString_idGenerateNewStringId(url foundation.NSURL, url2 foundation.NSURL, uid objectivec.IObject, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject, attributes objectivec.IObject, url3 bool, string_id uint64, id bool) ANEModel {
	instance := getANEModelClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithModelAtURL:sourceURL:UUID:key:identifierSource:cacheURLIdentifier:modelAttributes:standardizeURL:string_id:generateNewStringId:"), url, url2, uid, key, source, uRLIdentifier, attributes, url3, string_id, id)
	return ANEModelFromID(rv)
}

func NewANEModelWithModelAtURLSourceURLUUIDKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURLString_idGenerateNewStringIdMpsConstants(url foundation.NSURL, url2 foundation.NSURL, uid objectivec.IObject, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject, attributes objectivec.IObject, url3 bool, string_id uint64, id bool, constants objectivec.IObject) ANEModel {
	instance := getANEModelClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithModelAtURL:sourceURL:UUID:key:identifierSource:cacheURLIdentifier:modelAttributes:standardizeURL:string_id:generateNewStringId:mpsConstants:"), url, url2, uid, key, source, uRLIdentifier, attributes, url3, string_id, id, constants)
	return ANEModelFromID(rv)
}

func NewANEModelWithModelIdentifier(identifier objectivec.IObject) ANEModel {
	instance := getANEModelClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithModelIdentifier:"), identifier)
	return ANEModelFromID(rv)
}

func (a ANEModel) EncodeWithCoder(coder foundation.INSCoder) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}
func (a ANEModel) GetCacheURLIdentifier() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("getCacheURLIdentifier"))
	return objectivec.Object{ID: rv}
}
func (a ANEModel) GetUUID() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("getUUID"))
	return objectivec.Object{ID: rv}
}
func (a ANEModel) InputSymbolIndicesForProcedureIndex(index uint32) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("inputSymbolIndicesForProcedureIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (a ANEModel) IsEqualToModel(model objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("isEqualToModel:"), model)
	return rv
}
func (a ANEModel) OutputSymbolIndicesForProcedureIndex(index uint32) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("outputSymbolIndicesForProcedureIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (a ANEModel) ProcedureInfoForProcedureIndex(index uint32) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("procedureInfoForProcedureIndex:"), index)
	return objectivec.Object{ID: rv}
}
func (a ANEModel) ResetOnUnload() {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("resetOnUnload"))
}
func (a ANEModel) ShallowCopy() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("shallowCopy"))
	return objectivec.Object{ID: rv}
}
func (a ANEModel) SymbolIndicesForProcedureIndexIndexArrayKey(index uint32, key objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("symbolIndicesForProcedureIndex:indexArrayKey:"), index, key)
	return objectivec.Object{ID: rv}
}
func (a ANEModel) UpdateModelAttributesState(attributes objectivec.IObject, state uint64) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("updateModelAttributes:state:"), attributes, state)
}
func (a ANEModel) UpdateModelAttributesStateProgramHandleIntermediateBufferHandleQueueDepth(attributes objectivec.IObject, state uint64, handle uint64, handle2 uint64, depth int8) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("updateModelAttributes:state:programHandle:intermediateBufferHandle:queueDepth:"), attributes, state, handle, handle2, depth)
}
func (a ANEModel) InitWithCoder(coder foundation.INSCoder) ANEModel {
	rv := objc.SendIfResponds[ANEModel](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
func (a ANEModel) InitWithModelAtURLKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURL(url foundation.NSURL, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject, attributes objectivec.IObject, url2 bool) ANEModel {
	rv := objc.SendIfResponds[ANEModel](a.ID, objc.Sel("initWithModelAtURL:key:identifierSource:cacheURLIdentifier:modelAttributes:standardizeURL:"), url, key, source, uRLIdentifier, attributes, url2)
	return rv
}
func (a ANEModel) InitWithModelAtURLSourceURLUUIDKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURLString_idGenerateNewStringId(url foundation.NSURL, url2 foundation.NSURL, uid objectivec.IObject, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject, attributes objectivec.IObject, url3 bool, string_id uint64, id bool) ANEModel {
	rv := objc.SendIfResponds[ANEModel](a.ID, objc.Sel("initWithModelAtURL:sourceURL:UUID:key:identifierSource:cacheURLIdentifier:modelAttributes:standardizeURL:string_id:generateNewStringId:"), url, url2, uid, key, source, uRLIdentifier, attributes, url3, string_id, id)
	return rv
}
func (a ANEModel) InitWithModelAtURLSourceURLUUIDKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURLString_idGenerateNewStringIdMpsConstants(url foundation.NSURL, url2 foundation.NSURL, uid objectivec.IObject, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject, attributes objectivec.IObject, url3 bool, string_id uint64, id bool, constants objectivec.IObject) ANEModel {
	rv := objc.SendIfResponds[ANEModel](a.ID, objc.Sel("initWithModelAtURL:sourceURL:UUID:key:identifierSource:cacheURLIdentifier:modelAttributes:standardizeURL:string_id:generateNewStringId:mpsConstants:"), url, url2, uid, key, source, uRLIdentifier, attributes, url3, string_id, id, constants)
	return rv
}
func (a ANEModel) InitWithModelIdentifier(identifier objectivec.IObject) ANEModel {
	rv := objc.SendIfResponds[ANEModel](a.ID, objc.Sel("initWithModelIdentifier:"), identifier)
	return rv
}

func (_ANEModelClass ANEModelClass) CorrectFileURLFormat(uRLFormat objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEModelClass.class), objc.Sel("correctFileURLFormat:"), uRLFormat)
	return objectivec.Object{ID: rv}
}
func (_ANEModelClass ANEModelClass) ModelAtURLKey(url foundation.NSURL, key objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEModelClass.class), objc.Sel("modelAtURL:key:"), url, key)
	return objectivec.Object{ID: rv}
}
func (_ANEModelClass ANEModelClass) ModelAtURLKeyModelAttributes(url foundation.NSURL, key objectivec.IObject, attributes objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEModelClass.class), objc.Sel("modelAtURL:key:modelAttributes:"), url, key, attributes)
	return objectivec.Object{ID: rv}
}
func (_ANEModelClass ANEModelClass) ModelAtURLKeyMpsConstants(url foundation.NSURL, key objectivec.IObject, constants objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEModelClass.class), objc.Sel("modelAtURL:key:mpsConstants:"), url, key, constants)
	return objectivec.Object{ID: rv}
}
func (_ANEModelClass ANEModelClass) ModelAtURLWithCacheURLIdentifierKeyCacheURLIdentifier(uRLIdentifier objectivec.IObject, key objectivec.IObject, uRLIdentifier2 objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEModelClass.class), objc.Sel("modelAtURLWithCacheURLIdentifier:key:cacheURLIdentifier:"), uRLIdentifier, key, uRLIdentifier2)
	return objectivec.Object{ID: rv}
}
func (_ANEModelClass ANEModelClass) ModelAtURLWithSourceURLSourceURLKeyCacheURLIdentifier(url foundation.NSURL, url2 foundation.NSURL, key objectivec.IObject, uRLIdentifier objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEModelClass.class), objc.Sel("modelAtURLWithSourceURL:sourceURL:key:cacheURLIdentifier:"), url, url2, key, uRLIdentifier)
	return objectivec.Object{ID: rv}
}
func (_ANEModelClass ANEModelClass) ModelAtURLWithSourceURLSourceURLKeyIdentifierSourceCacheURLIdentifier(url foundation.NSURL, url2 foundation.NSURL, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEModelClass.class), objc.Sel("modelAtURLWithSourceURL:sourceURL:key:identifierSource:cacheURLIdentifier:"), url, url2, key, source, uRLIdentifier)
	return objectivec.Object{ID: rv}
}
func (_ANEModelClass ANEModelClass) ModelAtURLWithSourceURLSourceURLKeyIdentifierSourceCacheURLIdentifierUUID(url foundation.NSURL, url2 foundation.NSURL, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject, uid objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEModelClass.class), objc.Sel("modelAtURLWithSourceURL:sourceURL:key:identifierSource:cacheURLIdentifier:UUID:"), url, url2, key, source, uRLIdentifier, uid)
	return objectivec.Object{ID: rv}
}
func (_ANEModelClass ANEModelClass) ModelWithCacheURLIdentifier(uRLIdentifier objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEModelClass.class), objc.Sel("modelWithCacheURLIdentifier:"), uRLIdentifier)
	return objectivec.Object{ID: rv}
}
func (_ANEModelClass ANEModelClass) ModelWithCacheURLIdentifierUUID(uRLIdentifier objectivec.IObject, uid objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEModelClass.class), objc.Sel("modelWithCacheURLIdentifier:UUID:"), uRLIdentifier, uid)
	return objectivec.Object{ID: rv}
}
func (_ANEModelClass ANEModelClass) SupportsSecureCoding() bool {
	rv := objc.SendIfResponds[bool](objc.ID(_ANEModelClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

func (a ANEModel) UUID() foundation.NSUUID {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("UUID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}
func (a ANEModel) CacheURLIdentifier() string {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("cacheURLIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (a ANEModel) SetCacheURLIdentifier(value string) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setCacheURLIdentifier:"), objc.String(value))
}
func (a ANEModel) IdentifierSource() int64 {
	rv := objc.SendIfResponds[int64](a.ID, objc.Sel("identifierSource"))
	return rv
}
func (a ANEModel) IntermediateBufferHandle() uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("intermediateBufferHandle"))
	return rv
}
func (a ANEModel) SetIntermediateBufferHandle(value uint64) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setIntermediateBufferHandle:"), value)
}
func (a ANEModel) Key() string {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("key"))
	return foundation.NSStringFromID(rv).String()
}
func (a ANEModel) L() os.OSUnfairLockS {
	rv := objc.SendIfResponds[os.OSUnfairLockS](a.ID, objc.Sel("l"))
	return os.OSUnfairLockS(rv)
}
func (a ANEModel) SetL(value os.OSUnfairLockS) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setL:"), value)
}
func (a ANEModel) Mapper() IANEProgramIOSurfacesMapper {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("mapper"))
	return ANEProgramIOSurfacesMapperFromID(objc.ID(rv))
}
func (a ANEModel) SetMapper(value IANEProgramIOSurfacesMapper) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setMapper:"), value)
}
func (a ANEModel) ModelAttributes() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("modelAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (a ANEModel) SetModelAttributes(value foundation.INSDictionary) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setModelAttributes:"), value)
}
func (a ANEModel) ModelURL() foundation.NSURL {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("modelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (a ANEModel) MpsConstants() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("mpsConstants"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (a ANEModel) PerfStatsMask() uint32 {
	rv := objc.SendIfResponds[uint32](a.ID, objc.Sel("perfStatsMask"))
	return rv
}
func (a ANEModel) SetPerfStatsMask(value uint32) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setPerfStatsMask:"), value)
}
func (a ANEModel) Program() IANEProgramForEvaluation {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("program"))
	return ANEProgramForEvaluationFromID(objc.ID(rv))
}
func (a ANEModel) SetProgram(value IANEProgramForEvaluation) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setProgram:"), value)
}
func (a ANEModel) ProgramHandle() uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("programHandle"))
	return rv
}
func (a ANEModel) SetProgramHandle(value uint64) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setProgramHandle:"), value)
}
func (a ANEModel) QueueDepth() int8 {
	rv := objc.SendIfResponds[int8](a.ID, objc.Sel("queueDepth"))
	return rv
}
func (a ANEModel) SetQueueDepth(value int8) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setQueueDepth:"), value)
}
func (a ANEModel) SourceURL() foundation.NSURL {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("sourceURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (a ANEModel) State() uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("state"))
	return rv
}
func (a ANEModel) SetState(value uint64) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setState:"), value)
}
func (a ANEModel) String_id() uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("string_id"))
	return rv
}
func (a ANEModel) SetString_id(value uint64) {
	objc.SendIfResponds[struct{}](a.ID, objc.Sel("setString_id:"), value)
}
