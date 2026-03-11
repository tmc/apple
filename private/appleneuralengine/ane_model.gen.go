// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
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

// Alloc allocates memory for a new instance of the class.
func (ac ANEModelClass) Alloc() ANEModel {
	rv := objc.Send[ANEModel](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}







//
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
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel
type ANEModel struct {
	objectivec.Object
}

// ANEModelFromID constructs a [ANEModel] from an objc.ID.
func ANEModelFromID(id objc.ID) ANEModel {
	return ANEModel{objectivec.Object{id}}
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
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel
type IANEModel interface {
	objectivec.IObject

	// Topic: Methods

	UUID() foundation.NSUUID
	CacheURLIdentifier() string
	SetCacheURLIdentifier(value string)
	EncodeWithCoder(coder objectivec.IObject)
	GetCacheURLIdentifier() objectivec.IObject
	GetUUID() objectivec.IObject
	IdentifierSource() int64
	InputSymbolIndicesForProcedureIndex(index uint32) objectivec.IObject
	IntermediateBufferHandle() uint64
	SetIntermediateBufferHandle(value uint64)
	IsEqualToModel(model objectivec.IObject) bool
	Key() string
	L() objectivec.IObject
	SetL(value objectivec.IObject)
	Mapper() *ANEProgramIOSurfacesMapper
	SetMapper(value *ANEProgramIOSurfacesMapper)
	ModelAttributes() foundation.INSDictionary
	SetModelAttributes(value foundation.INSDictionary)
	ModelURL() foundation.INSURL
	MpsConstants() foundation.INSDictionary
	OutputSymbolIndicesForProcedureIndex(index uint32) objectivec.IObject
	PerfStatsMask() uint32
	SetPerfStatsMask(value uint32)
	ProcedureInfoForProcedureIndex(index uint32) objectivec.IObject
	Program() *ANEProgramForEvaluation
	SetProgram(value *ANEProgramForEvaluation)
	ProgramHandle() uint64
	SetProgramHandle(value uint64)
	QueueDepth() int8
	SetQueueDepth(value int8)
	ResetOnUnload()
	ShallowCopy() objectivec.IObject
	SourceURL() foundation.INSURL
	State() uint64
	SetState(value uint64)
	String_id() uint64
	SetString_id(value uint64)
	SymbolIndicesForProcedureIndexIndexArrayKey(index uint32, key objectivec.IObject) objectivec.IObject
	UpdateModelAttributesState(attributes objectivec.IObject, state uint64)
	UpdateModelAttributesStateProgramHandleIntermediateBufferHandleQueueDepth(attributes objectivec.IObject, state uint64, handle uint64, handle2 uint64, depth int8)
	InitWithCoder(coder objectivec.IObject) ANEModel
	InitWithModelAtURLKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURL(url foundation.INSURL, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject, attributes objectivec.IObject, url2 bool) ANEModel
	InitWithModelAtURLSourceURLUUIDKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURLString_idGenerateNewStringId(url foundation.INSURL, url2 foundation.INSURL, uid objectivec.IObject, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject, attributes objectivec.IObject, url3 bool, string_id uint64, id bool) ANEModel
	InitWithModelAtURLSourceURLUUIDKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURLString_idGenerateNewStringIdMpsConstants(url foundation.INSURL, url2 foundation.INSURL, uid objectivec.IObject, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject, attributes objectivec.IObject, url3 bool, string_id uint64, id bool, constants objectivec.IObject) ANEModel
	InitWithModelIdentifier(identifier objectivec.IObject) ANEModel
}





// Init initializes the instance.
func (a ANEModel) Init() ANEModel {
	rv := objc.Send[ANEModel](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEModel) Autorelease() ANEModel {
	rv := objc.Send[ANEModel](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEModel creates a new ANEModel instance.
func NewANEModel() ANEModel {
	class := getANEModelClass()
	rv := objc.Send[ANEModel](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/initWithCoder:
func NewANEModelWithCoder(coder objectivec.IObject) ANEModel {
	instance := getANEModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return ANEModelFromID(rv)
}


//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/initWithModelAtURL:key:identifierSource:cacheURLIdentifier:modelAttributes:standardizeURL:
func NewANEModelWithModelAtURLKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURL(url foundation.INSURL, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject, attributes objectivec.IObject, url2 bool) ANEModel {
	instance := getANEModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelAtURL:key:identifierSource:cacheURLIdentifier:modelAttributes:standardizeURL:"), url, key, source, uRLIdentifier, attributes, url2)
	return ANEModelFromID(rv)
}


//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/initWithModelAtURL:sourceURL:UUID:key:identifierSource:cacheURLIdentifier:modelAttributes:standardizeURL:string_id:generateNewStringId:
func NewANEModelWithModelAtURLSourceURLUUIDKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURLString_idGenerateNewStringId(url foundation.INSURL, url2 foundation.INSURL, uid objectivec.IObject, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject, attributes objectivec.IObject, url3 bool, string_id uint64, id bool) ANEModel {
	instance := getANEModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelAtURL:sourceURL:UUID:key:identifierSource:cacheURLIdentifier:modelAttributes:standardizeURL:string_id:generateNewStringId:"), url, url2, uid, key, source, uRLIdentifier, attributes, url3, string_id, id)
	return ANEModelFromID(rv)
}


//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/initWithModelAtURL:sourceURL:UUID:key:identifierSource:cacheURLIdentifier:modelAttributes:standardizeURL:string_id:generateNewStringId:mpsConstants:
func NewANEModelWithModelAtURLSourceURLUUIDKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURLString_idGenerateNewStringIdMpsConstants(url foundation.INSURL, url2 foundation.INSURL, uid objectivec.IObject, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject, attributes objectivec.IObject, url3 bool, string_id uint64, id bool, constants objectivec.IObject) ANEModel {
	instance := getANEModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelAtURL:sourceURL:UUID:key:identifierSource:cacheURLIdentifier:modelAttributes:standardizeURL:string_id:generateNewStringId:mpsConstants:"), url, url2, uid, key, source, uRLIdentifier, attributes, url3, string_id, id, constants)
	return ANEModelFromID(rv)
}


//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/initWithModelIdentifier:
func NewANEModelWithModelIdentifier(identifier objectivec.IObject) ANEModel {
	instance := getANEModelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithModelIdentifier:"), identifier)
	return ANEModelFromID(rv)
}







//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/encodeWithCoder:
func (a ANEModel) EncodeWithCoder(coder objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/getCacheURLIdentifier
func (a ANEModel) GetCacheURLIdentifier() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("getCacheURLIdentifier"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/getUUID
func (a ANEModel) GetUUID() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("getUUID"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/inputSymbolIndicesForProcedureIndex:
func (a ANEModel) InputSymbolIndicesForProcedureIndex(index uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputSymbolIndicesForProcedureIndex:"), index)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/isEqualToModel:
func (a ANEModel) IsEqualToModel(model objectivec.IObject) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isEqualToModel:"), model)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/outputSymbolIndicesForProcedureIndex:
func (a ANEModel) OutputSymbolIndicesForProcedureIndex(index uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputSymbolIndicesForProcedureIndex:"), index)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/procedureInfoForProcedureIndex:
func (a ANEModel) ProcedureInfoForProcedureIndex(index uint32) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("procedureInfoForProcedureIndex:"), index)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/resetOnUnload
func (a ANEModel) ResetOnUnload() {
	objc.Send[objc.ID](a.ID, objc.Sel("resetOnUnload"))
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/shallowCopy
func (a ANEModel) ShallowCopy() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("shallowCopy"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/symbolIndicesForProcedureIndex:indexArrayKey:
func (a ANEModel) SymbolIndicesForProcedureIndexIndexArrayKey(index uint32, key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("symbolIndicesForProcedureIndex:indexArrayKey:"), index, key)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/updateModelAttributes:state:
func (a ANEModel) UpdateModelAttributesState(attributes objectivec.IObject, state uint64) {
	objc.Send[objc.ID](a.ID, objc.Sel("updateModelAttributes:state:"), attributes, state)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/updateModelAttributes:state:programHandle:intermediateBufferHandle:queueDepth:
func (a ANEModel) UpdateModelAttributesStateProgramHandleIntermediateBufferHandleQueueDepth(attributes objectivec.IObject, state uint64, handle uint64, handle2 uint64, depth int8) {
	objc.Send[objc.ID](a.ID, objc.Sel("updateModelAttributes:state:programHandle:intermediateBufferHandle:queueDepth:"), attributes, state, handle, handle2, depth)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/initWithCoder:
func (a ANEModel) InitWithCoder(coder objectivec.IObject) ANEModel {
	rv := objc.Send[ANEModel](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/initWithModelAtURL:key:identifierSource:cacheURLIdentifier:modelAttributes:standardizeURL:
func (a ANEModel) InitWithModelAtURLKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURL(url foundation.INSURL, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject, attributes objectivec.IObject, url2 bool) ANEModel {
	rv := objc.Send[ANEModel](a.ID, objc.Sel("initWithModelAtURL:key:identifierSource:cacheURLIdentifier:modelAttributes:standardizeURL:"), url, key, source, uRLIdentifier, attributes, url2)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/initWithModelAtURL:sourceURL:UUID:key:identifierSource:cacheURLIdentifier:modelAttributes:standardizeURL:string_id:generateNewStringId:
func (a ANEModel) InitWithModelAtURLSourceURLUUIDKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURLString_idGenerateNewStringId(url foundation.INSURL, url2 foundation.INSURL, uid objectivec.IObject, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject, attributes objectivec.IObject, url3 bool, string_id uint64, id bool) ANEModel {
	rv := objc.Send[ANEModel](a.ID, objc.Sel("initWithModelAtURL:sourceURL:UUID:key:identifierSource:cacheURLIdentifier:modelAttributes:standardizeURL:string_id:generateNewStringId:"), url, url2, uid, key, source, uRLIdentifier, attributes, url3, string_id, id)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/initWithModelAtURL:sourceURL:UUID:key:identifierSource:cacheURLIdentifier:modelAttributes:standardizeURL:string_id:generateNewStringId:mpsConstants:
func (a ANEModel) InitWithModelAtURLSourceURLUUIDKeyIdentifierSourceCacheURLIdentifierModelAttributesStandardizeURLString_idGenerateNewStringIdMpsConstants(url foundation.INSURL, url2 foundation.INSURL, uid objectivec.IObject, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject, attributes objectivec.IObject, url3 bool, string_id uint64, id bool, constants objectivec.IObject) ANEModel {
	rv := objc.Send[ANEModel](a.ID, objc.Sel("initWithModelAtURL:sourceURL:UUID:key:identifierSource:cacheURLIdentifier:modelAttributes:standardizeURL:string_id:generateNewStringId:mpsConstants:"), url, url2, uid, key, source, uRLIdentifier, attributes, url3, string_id, id, constants)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/initWithModelIdentifier:
func (a ANEModel) InitWithModelIdentifier(identifier objectivec.IObject) ANEModel {
	rv := objc.Send[ANEModel](a.ID, objc.Sel("initWithModelIdentifier:"), identifier)
	return rv
}





//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/correctFileURLFormat:
func (_ANEModelClass ANEModelClass) CorrectFileURLFormat(uRLFormat objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEModelClass.class), objc.Sel("correctFileURLFormat:"), uRLFormat)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/modelAtURL:key:
func (_ANEModelClass ANEModelClass) ModelAtURLKey(url foundation.INSURL, key objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEModelClass.class), objc.Sel("modelAtURL:key:"), url, key)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/modelAtURL:key:modelAttributes:
func (_ANEModelClass ANEModelClass) ModelAtURLKeyModelAttributes(url foundation.INSURL, key objectivec.IObject, attributes objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEModelClass.class), objc.Sel("modelAtURL:key:modelAttributes:"), url, key, attributes)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/modelAtURL:key:mpsConstants:
func (_ANEModelClass ANEModelClass) ModelAtURLKeyMpsConstants(url foundation.INSURL, key objectivec.IObject, constants objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEModelClass.class), objc.Sel("modelAtURL:key:mpsConstants:"), url, key, constants)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/modelAtURLWithCacheURLIdentifier:key:cacheURLIdentifier:
func (_ANEModelClass ANEModelClass) ModelAtURLWithCacheURLIdentifierKeyCacheURLIdentifier(uRLIdentifier objectivec.IObject, key objectivec.IObject, uRLIdentifier2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEModelClass.class), objc.Sel("modelAtURLWithCacheURLIdentifier:key:cacheURLIdentifier:"), uRLIdentifier, key, uRLIdentifier2)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/modelAtURLWithSourceURL:sourceURL:key:cacheURLIdentifier:
func (_ANEModelClass ANEModelClass) ModelAtURLWithSourceURLSourceURLKeyCacheURLIdentifier(url foundation.INSURL, url2 foundation.INSURL, key objectivec.IObject, uRLIdentifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEModelClass.class), objc.Sel("modelAtURLWithSourceURL:sourceURL:key:cacheURLIdentifier:"), url, url2, key, uRLIdentifier)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/modelAtURLWithSourceURL:sourceURL:key:identifierSource:cacheURLIdentifier:
func (_ANEModelClass ANEModelClass) ModelAtURLWithSourceURLSourceURLKeyIdentifierSourceCacheURLIdentifier(url foundation.INSURL, url2 foundation.INSURL, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEModelClass.class), objc.Sel("modelAtURLWithSourceURL:sourceURL:key:identifierSource:cacheURLIdentifier:"), url, url2, key, source, uRLIdentifier)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/modelAtURLWithSourceURL:sourceURL:key:identifierSource:cacheURLIdentifier:UUID:
func (_ANEModelClass ANEModelClass) ModelAtURLWithSourceURLSourceURLKeyIdentifierSourceCacheURLIdentifierUUID(url foundation.INSURL, url2 foundation.INSURL, key objectivec.IObject, source int64, uRLIdentifier objectivec.IObject, uid objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEModelClass.class), objc.Sel("modelAtURLWithSourceURL:sourceURL:key:identifierSource:cacheURLIdentifier:UUID:"), url, url2, key, source, uRLIdentifier, uid)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/modelWithCacheURLIdentifier:
func (_ANEModelClass ANEModelClass) ModelWithCacheURLIdentifier(uRLIdentifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEModelClass.class), objc.Sel("modelWithCacheURLIdentifier:"), uRLIdentifier)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/modelWithCacheURLIdentifier:UUID:
func (_ANEModelClass ANEModelClass) ModelWithCacheURLIdentifierUUID(uRLIdentifier objectivec.IObject, uid objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEModelClass.class), objc.Sel("modelWithCacheURLIdentifier:UUID:"), uRLIdentifier, uid)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/supportsSecureCoding
func (_ANEModelClass ANEModelClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_ANEModelClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}








// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/UUID
func (a ANEModel) UUID() foundation.NSUUID {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("UUID"))
	return foundation.NSUUIDFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/cacheURLIdentifier
func (a ANEModel) CacheURLIdentifier() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("cacheURLIdentifier"))
	return foundation.NSStringFromID(rv).String()
}
func (a ANEModel) SetCacheURLIdentifier(value string) {
	objc.Send[struct{}](a.ID, objc.Sel("setCacheURLIdentifier:"), objc.String(value))
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/identifierSource
func (a ANEModel) IdentifierSource() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("identifierSource"))
	return rv
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/intermediateBufferHandle
func (a ANEModel) IntermediateBufferHandle() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("intermediateBufferHandle"))
	return rv
}
func (a ANEModel) SetIntermediateBufferHandle(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setIntermediateBufferHandle:"), value)
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/key
func (a ANEModel) Key() string {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("key"))
	return foundation.NSStringFromID(rv).String()
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/l
func (a ANEModel) L() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("l"))
	return objectivec.Object{ID: rv}
}
func (a ANEModel) SetL(value objectivec.IObject) {
	objc.Send[struct{}](a.ID, objc.Sel("setL:"), value)
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/mapper
func (a ANEModel) Mapper() *ANEProgramIOSurfacesMapper {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("mapper"))
	if rv == 0 {
		return nil
	}
	val := ANEProgramIOSurfacesMapperFromID(objc.ID(rv))
	return &val
}
func (a ANEModel) SetMapper(value *ANEProgramIOSurfacesMapper) {
	objc.Send[struct{}](a.ID, objc.Sel("setMapper:"), value)
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/modelAttributes
func (a ANEModel) ModelAttributes() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("modelAttributes"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (a ANEModel) SetModelAttributes(value foundation.INSDictionary) {
	objc.Send[struct{}](a.ID, objc.Sel("setModelAttributes:"), value)
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/modelURL
func (a ANEModel) ModelURL() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("modelURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/mpsConstants
func (a ANEModel) MpsConstants() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("mpsConstants"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/perfStatsMask
func (a ANEModel) PerfStatsMask() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("perfStatsMask"))
	return rv
}
func (a ANEModel) SetPerfStatsMask(value uint32) {
	objc.Send[struct{}](a.ID, objc.Sel("setPerfStatsMask:"), value)
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/program
func (a ANEModel) Program() *ANEProgramForEvaluation {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("program"))
	if rv == 0 {
		return nil
	}
	val := ANEProgramForEvaluationFromID(objc.ID(rv))
	return &val
}
func (a ANEModel) SetProgram(value *ANEProgramForEvaluation) {
	objc.Send[struct{}](a.ID, objc.Sel("setProgram:"), value)
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/programHandle
func (a ANEModel) ProgramHandle() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("programHandle"))
	return rv
}
func (a ANEModel) SetProgramHandle(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setProgramHandle:"), value)
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/queueDepth
func (a ANEModel) QueueDepth() int8 {
	rv := objc.Send[int8](a.ID, objc.Sel("queueDepth"))
	return rv
}
func (a ANEModel) SetQueueDepth(value int8) {
	objc.Send[struct{}](a.ID, objc.Sel("setQueueDepth:"), value)
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/sourceURL
func (a ANEModel) SourceURL() foundation.INSURL {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sourceURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/state
func (a ANEModel) State() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("state"))
	return rv
}
func (a ANEModel) SetState(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setState:"), value)
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEModel/string_id
func (a ANEModel) String_id() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("string_id"))
	return rv
}
func (a ANEModel) SetString_id(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setString_id:"), value)
}

















