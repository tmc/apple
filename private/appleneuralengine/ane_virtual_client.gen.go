// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEVirtualClient] class.
var (
	_ANEVirtualClientClass     ANEVirtualClientClass
	_ANEVirtualClientClassOnce sync.Once
)

func getANEVirtualClientClass() ANEVirtualClientClass {
	_ANEVirtualClientClassOnce.Do(func() {
		_ANEVirtualClientClass = ANEVirtualClientClass{class: objc.GetClass("_ANEVirtualClient")}
	})
	return _ANEVirtualClientClass
}

// GetANEVirtualClientClass returns the class object for _ANEVirtualClient.
func GetANEVirtualClientClass() ANEVirtualClientClass {
	return getANEVirtualClientClass()
}

type ANEVirtualClientClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEVirtualClientClass) Alloc() ANEVirtualClient {
	rv := objc.Send[ANEVirtualClient](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ANEVirtualClient.AneArchitectureTypeStr]
//   - [ANEVirtualClient.AneBoardtype]
//   - [ANEVirtualClient.BeginRealTimeTask]
//   - [ANEVirtualClient.CallIOUserClientInParamsOutParams]
//   - [ANEVirtualClient.CallIOUserClientWithDictionaryInDictionaryError]
//   - [ANEVirtualClient.CheckKernReturnValueSelectorOutParams]
//   - [ANEVirtualClient.CompileModelOptionsQosError]
//   - [ANEVirtualClient.CompiledModelExistsFor]
//   - [ANEVirtualClient.CompiledModelExistsMatchingHash]
//   - [ANEVirtualClient.Connect]
//   - [ANEVirtualClient.CopyAllModelFilesDictionaryIoSurfaceRefs]
//   - [ANEVirtualClient.CopyDictionaryToIOSurfaceCopiedDataSizeCreatedIOSID]
//   - [ANEVirtualClient.CopyErrorValue]
//   - [ANEVirtualClient.CopyErrorValueVmData]
//   - [ANEVirtualClient.CopyFilesInDirectoryToIOSurfacesIoSurfaceRefsIoSurfaceSizesFileNames]
//   - [ANEVirtualClient.CopyModelOptionsVmData]
//   - [ANEVirtualClient.CopyModelMetaDataOptionsDictionaryVmData]
//   - [ANEVirtualClient.CopyModelOptionFilesOptionsDictionaryVmData]
//   - [ANEVirtualClient.CopyModelOptionFilesOptionsVmData]
//   - [ANEVirtualClient.CopyOptionsDictionaryVmData]
//   - [ANEVirtualClient.CopyOptionsVmData]
//   - [ANEVirtualClient.CopyToIOSurfaceLengthIoSID]
//   - [ANEVirtualClient.CopyToIOSurfaceSizeIoSID]
//   - [ANEVirtualClient.DoEvaluateWithModelOptionsRequestQosCompletionEventError]
//   - [ANEVirtualClient.DoEvaluateWithModelLegacyOptionsRequestQosCompletionEventError]
//   - [ANEVirtualClient.DoJsonParsingMatchWeightName]
//   - [ANEVirtualClient.DoMapIOSurfacesWithModelRequestCacheInferenceError]
//   - [ANEVirtualClient.Echo]
//   - [ANEVirtualClient.EndRealTimeTask]
//   - [ANEVirtualClient.EvaluateWithModelOptionsRequestQosError]
//   - [ANEVirtualClient.ExchangeBuildVersionInfo]
//   - [ANEVirtualClient.GetDeviceInfo]
//   - [ANEVirtualClient.GetModelAttribute]
//   - [ANEVirtualClient.GetValidateNetworkVersion]
//   - [ANEVirtualClient.HasANE]
//   - [ANEVirtualClient.HostBuildVersionStr]
//   - [ANEVirtualClient.IsInternalBuild]
//   - [ANEVirtualClient.LoadModelOptionsQosError]
//   - [ANEVirtualClient.LoadModelNewInstanceOptionsModelInstParamsQosError]
//   - [ANEVirtualClient.LoadModelNewInstanceLegacyOptionsModelInstParamsQosError]
//   - [ANEVirtualClient.MapIOSurfacesWithModelRequestCacheInferenceError]
//   - [ANEVirtualClient.NegotiatedCapabilityMask]
//   - [ANEVirtualClient.NegotiatedDataInterfaceVersion]
//   - [ANEVirtualClient.NumANECores]
//   - [ANEVirtualClient.NumANEs]
//   - [ANEVirtualClient.OutputDictIOSurfaceSize]
//   - [ANEVirtualClient.ParallelDecompressedData]
//   - [ANEVirtualClient.PrintDictionary]
//   - [ANEVirtualClient.PurgeCompiledModel]
//   - [ANEVirtualClient.PurgeCompiledModelMatchingHash]
//   - [ANEVirtualClient.Queue]
//   - [ANEVirtualClient.ReadWeightFilename]
//   - [ANEVirtualClient.ReleaseIOSurfaces]
//   - [ANEVirtualClient.SendGuestBuildVersion]
//   - [ANEVirtualClient.SessionHintWithModelHintOptionsReportError]
//   - [ANEVirtualClient.TransferFileToHostWithPathWithChunkSizeWithUUIDWithModelInputPathOverWriteFileNameWith]
//   - [ANEVirtualClient.UnloadModelOptionsQosError]
//   - [ANEVirtualClient.UpdateErrorError]
//   - [ANEVirtualClient.UpdatePerformanceStats]
//   - [ANEVirtualClient.ValidateEnvironmentForPrecompiledBinarySupport]
//   - [ANEVirtualClient.ValidateNetworkCreateUuidFunctionDirectoryPathScratchPadPathMilTextData]
//   - [ANEVirtualClient.ValidateNetworkCreateMLIRValidation_params]
//   - [ANEVirtualClient.InitWithSingletonAccess]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient
type ANEVirtualClient struct {
	objectivec.Object
}

// ANEVirtualClientFromID constructs a [ANEVirtualClient] from an objc.ID.
func ANEVirtualClientFromID(id objc.ID) ANEVirtualClient {
	return ANEVirtualClient{objectivec.Object{id}}
}
// Ensure ANEVirtualClient implements IANEVirtualClient.
var _ IANEVirtualClient = ANEVirtualClient{}

// An interface definition for the [ANEVirtualClient] class.
//
// # Methods
//
//   - [IANEVirtualClient.AneArchitectureTypeStr]
//   - [IANEVirtualClient.AneBoardtype]
//   - [IANEVirtualClient.BeginRealTimeTask]
//   - [IANEVirtualClient.CallIOUserClientInParamsOutParams]
//   - [IANEVirtualClient.CallIOUserClientWithDictionaryInDictionaryError]
//   - [IANEVirtualClient.CheckKernReturnValueSelectorOutParams]
//   - [IANEVirtualClient.CompileModelOptionsQosError]
//   - [IANEVirtualClient.CompiledModelExistsFor]
//   - [IANEVirtualClient.CompiledModelExistsMatchingHash]
//   - [IANEVirtualClient.Connect]
//   - [IANEVirtualClient.CopyAllModelFilesDictionaryIoSurfaceRefs]
//   - [IANEVirtualClient.CopyDictionaryToIOSurfaceCopiedDataSizeCreatedIOSID]
//   - [IANEVirtualClient.CopyErrorValue]
//   - [IANEVirtualClient.CopyErrorValueVmData]
//   - [IANEVirtualClient.CopyFilesInDirectoryToIOSurfacesIoSurfaceRefsIoSurfaceSizesFileNames]
//   - [IANEVirtualClient.CopyModelOptionsVmData]
//   - [IANEVirtualClient.CopyModelMetaDataOptionsDictionaryVmData]
//   - [IANEVirtualClient.CopyModelOptionFilesOptionsDictionaryVmData]
//   - [IANEVirtualClient.CopyModelOptionFilesOptionsVmData]
//   - [IANEVirtualClient.CopyOptionsDictionaryVmData]
//   - [IANEVirtualClient.CopyOptionsVmData]
//   - [IANEVirtualClient.CopyToIOSurfaceLengthIoSID]
//   - [IANEVirtualClient.CopyToIOSurfaceSizeIoSID]
//   - [IANEVirtualClient.DoEvaluateWithModelOptionsRequestQosCompletionEventError]
//   - [IANEVirtualClient.DoEvaluateWithModelLegacyOptionsRequestQosCompletionEventError]
//   - [IANEVirtualClient.DoJsonParsingMatchWeightName]
//   - [IANEVirtualClient.DoMapIOSurfacesWithModelRequestCacheInferenceError]
//   - [IANEVirtualClient.Echo]
//   - [IANEVirtualClient.EndRealTimeTask]
//   - [IANEVirtualClient.EvaluateWithModelOptionsRequestQosError]
//   - [IANEVirtualClient.ExchangeBuildVersionInfo]
//   - [IANEVirtualClient.GetDeviceInfo]
//   - [IANEVirtualClient.GetModelAttribute]
//   - [IANEVirtualClient.GetValidateNetworkVersion]
//   - [IANEVirtualClient.HasANE]
//   - [IANEVirtualClient.HostBuildVersionStr]
//   - [IANEVirtualClient.IsInternalBuild]
//   - [IANEVirtualClient.LoadModelOptionsQosError]
//   - [IANEVirtualClient.LoadModelNewInstanceOptionsModelInstParamsQosError]
//   - [IANEVirtualClient.LoadModelNewInstanceLegacyOptionsModelInstParamsQosError]
//   - [IANEVirtualClient.MapIOSurfacesWithModelRequestCacheInferenceError]
//   - [IANEVirtualClient.NegotiatedCapabilityMask]
//   - [IANEVirtualClient.NegotiatedDataInterfaceVersion]
//   - [IANEVirtualClient.NumANECores]
//   - [IANEVirtualClient.NumANEs]
//   - [IANEVirtualClient.OutputDictIOSurfaceSize]
//   - [IANEVirtualClient.ParallelDecompressedData]
//   - [IANEVirtualClient.PrintDictionary]
//   - [IANEVirtualClient.PurgeCompiledModel]
//   - [IANEVirtualClient.PurgeCompiledModelMatchingHash]
//   - [IANEVirtualClient.Queue]
//   - [IANEVirtualClient.ReadWeightFilename]
//   - [IANEVirtualClient.ReleaseIOSurfaces]
//   - [IANEVirtualClient.SendGuestBuildVersion]
//   - [IANEVirtualClient.SessionHintWithModelHintOptionsReportError]
//   - [IANEVirtualClient.TransferFileToHostWithPathWithChunkSizeWithUUIDWithModelInputPathOverWriteFileNameWith]
//   - [IANEVirtualClient.UnloadModelOptionsQosError]
//   - [IANEVirtualClient.UpdateErrorError]
//   - [IANEVirtualClient.UpdatePerformanceStats]
//   - [IANEVirtualClient.ValidateEnvironmentForPrecompiledBinarySupport]
//   - [IANEVirtualClient.ValidateNetworkCreateUuidFunctionDirectoryPathScratchPadPathMilTextData]
//   - [IANEVirtualClient.ValidateNetworkCreateMLIRValidation_params]
//   - [IANEVirtualClient.InitWithSingletonAccess]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient
type IANEVirtualClient interface {
	objectivec.IObject

	// Topic: Methods

	AneArchitectureTypeStr() objectivec.IObject
	AneBoardtype() int64
	BeginRealTimeTask() bool
	CallIOUserClientInParamsOutParams(client uint32, params unsafe.Pointer, params2 unsafe.Pointer) bool
	CallIOUserClientWithDictionaryInDictionaryError(dictionary uint32, dictionary2 objectivec.IObject) (objectivec.IObject, error)
	CheckKernReturnValueSelectorOutParams(value int, selector uint32, params unsafe.Pointer)
	CompileModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error)
	CompiledModelExistsFor(for_ objectivec.IObject) bool
	CompiledModelExistsMatchingHash(hash objectivec.IObject) bool
	Connect() uint32
	CopyAllModelFilesDictionaryIoSurfaceRefs(files objectivec.IObject, dictionary objectivec.IObject, refs objectivec.IObject) bool
	CopyDictionaryToIOSurfaceCopiedDataSizeCreatedIOSID(iOSurface objectivec.IObject, size unsafe.Pointer, iosid unsafe.Pointer) coregraphics.IOSurfaceRef
	CopyErrorValue(value unsafe.Pointer)
	CopyErrorValueVmData(value objectivec.IObject, data unsafe.Pointer)
	CopyFilesInDirectoryToIOSurfacesIoSurfaceRefsIoSurfaceSizesFileNames(iOSurfaces objectivec.IObject, refs objectivec.IObject, sizes objectivec.IObject, names objectivec.IObject) bool
	CopyModelOptionsVmData(model objectivec.IObject, options objectivec.IObject, data unsafe.Pointer)
	CopyModelMetaDataOptionsDictionaryVmData(data objectivec.IObject, options objectivec.IObject, dictionary objectivec.IObject, data2 unsafe.Pointer)
	CopyModelOptionFilesOptionsDictionaryVmData(files objectivec.IObject, options objectivec.IObject, dictionary objectivec.IObject, data unsafe.Pointer)
	CopyModelOptionFilesOptionsVmData(files objectivec.IObject, options objectivec.IObject, data unsafe.Pointer)
	CopyOptionsDictionaryVmData(options objectivec.IObject, dictionary objectivec.IObject, data unsafe.Pointer)
	CopyOptionsVmData(options objectivec.IObject, data unsafe.Pointer)
	CopyToIOSurfaceLengthIoSID(iOSurface objectivec.IObject, length uint64, sid unsafe.Pointer) coregraphics.IOSurfaceRef
	CopyToIOSurfaceSizeIoSID(iOSurface []byte, size uint64, sid unsafe.Pointer) coregraphics.IOSurfaceRef
	DoEvaluateWithModelOptionsRequestQosCompletionEventError(model objectivec.IObject, options objectivec.IObject, request objectivec.IObject, qos uint32, event objectivec.IObject) (bool, error)
	DoEvaluateWithModelLegacyOptionsRequestQosCompletionEventError(legacy objectivec.IObject, options objectivec.IObject, request objectivec.IObject, qos uint32, event objectivec.IObject) (bool, error)
	DoJsonParsingMatchWeightName(name objectivec.IObject) objectivec.IObject
	DoMapIOSurfacesWithModelRequestCacheInferenceError(model objectivec.IObject, request objectivec.IObject, inference bool) (bool, error)
	Echo(echo objectivec.IObject) bool
	EndRealTimeTask() bool
	EvaluateWithModelOptionsRequestQosError(model objectivec.IObject, options objectivec.IObject, request objectivec.IObject, qos uint32) (bool, error)
	ExchangeBuildVersionInfo() objectivec.IObject
	GetDeviceInfo() objectivec.IObject
	GetModelAttribute(attribute unsafe.Pointer) objectivec.IObject
	GetValidateNetworkVersion() uint64
	HasANE() bool
	HostBuildVersionStr() objectivec.IObject
	IsInternalBuild() bool
	LoadModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error)
	LoadModelNewInstanceOptionsModelInstParamsQosError(instance objectivec.IObject, options objectivec.IObject, params objectivec.IObject, qos uint32) (bool, error)
	LoadModelNewInstanceLegacyOptionsModelInstParamsQosError(legacy objectivec.IObject, options objectivec.IObject, params objectivec.IObject, qos uint32) (bool, error)
	MapIOSurfacesWithModelRequestCacheInferenceError(model objectivec.IObject, request objectivec.IObject, inference bool) (bool, error)
	NegotiatedCapabilityMask() uint64
	NegotiatedDataInterfaceVersion() uint32
	NumANECores() uint32
	NumANEs() uint32
	OutputDictIOSurfaceSize() uint64
	ParallelDecompressedData(data objectivec.IObject) objectivec.IObject
	PrintDictionary(dictionary objectivec.IObject)
	PurgeCompiledModel(model objectivec.IObject)
	PurgeCompiledModelMatchingHash(hash objectivec.IObject)
	Queue() objectivec.Object
	ReadWeightFilename(filename objectivec.IObject) objectivec.IObject
	ReleaseIOSurfaces(iOSurfaces unsafe.Pointer)
	SendGuestBuildVersion()
	SessionHintWithModelHintOptionsReportError(model objectivec.IObject, hint objectivec.IObject, options objectivec.IObject, report objectivec.IObject) (bool, error)
	TransferFileToHostWithPathWithChunkSizeWithUUIDWithModelInputPathOverWriteFileNameWith(path objectivec.IObject, size uint32, uuid objectivec.IObject, path2 objectivec.IObject, with objectivec.IObject) bool
	UnloadModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error)
	UpdateErrorError(error_ unsafe.Pointer) (bool, error)
	UpdatePerformanceStats(stats unsafe.Pointer) objectivec.IObject
	ValidateEnvironmentForPrecompiledBinarySupport() bool
	ValidateNetworkCreateUuidFunctionDirectoryPathScratchPadPathMilTextData(create uint64, uuid objectivec.IObject, function objectivec.IObject, path objectivec.IObject, path2 objectivec.IObject, data objectivec.IObject) objectivec.IObject
	ValidateNetworkCreateMLIRValidation_params(mlir uint64, validation_params objectivec.IObject) objectivec.IObject
	InitWithSingletonAccess() ANEVirtualClient
}

// Init initializes the instance.
func (a ANEVirtualClient) Init() ANEVirtualClient {
	rv := objc.Send[ANEVirtualClient](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEVirtualClient) Autorelease() ANEVirtualClient {
	rv := objc.Send[ANEVirtualClient](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEVirtualClient creates a new ANEVirtualClient instance.
func NewANEVirtualClient() ANEVirtualClient {
	class := getANEVirtualClientClass()
	rv := objc.Send[ANEVirtualClient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/initWithSingletonAccess
func NewANEVirtualClientWithSingletonAccess() ANEVirtualClient {
	instance := getANEVirtualClientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSingletonAccess"))
	return ANEVirtualClientFromID(rv)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/aneArchitectureTypeStr
func (a ANEVirtualClient) AneArchitectureTypeStr() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("aneArchitectureTypeStr"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/aneBoardtype
func (a ANEVirtualClient) AneBoardtype() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("aneBoardtype"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/beginRealTimeTask
func (a ANEVirtualClient) BeginRealTimeTask() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("beginRealTimeTask"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/callIOUserClient:inParams:outParams:
func (a ANEVirtualClient) CallIOUserClientInParamsOutParams(client uint32, params unsafe.Pointer, params2 unsafe.Pointer) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("callIOUserClient:inParams:outParams:"), client, params, params2)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/callIOUserClientWithDictionary:inDictionary:error:
func (a ANEVirtualClient) CallIOUserClientWithDictionaryInDictionaryError(dictionary uint32, dictionary2 objectivec.IObject) (objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](a.ID, objc.Sel("callIOUserClientWithDictionary:inDictionary:error:"), dictionary, dictionary2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objectivec.Object{ID: rv}, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/checkKernReturnValue:selector:outParams:
func (a ANEVirtualClient) CheckKernReturnValueSelectorOutParams(value int, selector uint32, params unsafe.Pointer) {
	objc.Send[objc.ID](a.ID, objc.Sel("checkKernReturnValue:selector:outParams:"), value, selector, params)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/compileModel:options:qos:error:
func (a ANEVirtualClient) CompileModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("compileModel:options:qos:error:"), model, options, qos, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("compileModel:options:qos:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/compiledModelExistsFor:
func (a ANEVirtualClient) CompiledModelExistsFor(for_ objectivec.IObject) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("compiledModelExistsFor:"), for_)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/compiledModelExistsMatchingHash:
func (a ANEVirtualClient) CompiledModelExistsMatchingHash(hash objectivec.IObject) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("compiledModelExistsMatchingHash:"), hash)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/copyAllModelFiles:dictionary:ioSurfaceRefs:
func (a ANEVirtualClient) CopyAllModelFilesDictionaryIoSurfaceRefs(files objectivec.IObject, dictionary objectivec.IObject, refs objectivec.IObject) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("copyAllModelFiles:dictionary:ioSurfaceRefs:"), files, dictionary, refs)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/copyDictionaryToIOSurface:copiedDataSize:createdIOSID:
func (a ANEVirtualClient) CopyDictionaryToIOSurfaceCopiedDataSizeCreatedIOSID(iOSurface objectivec.IObject, size unsafe.Pointer, iosid unsafe.Pointer) coregraphics.IOSurfaceRef {
	rv := objc.Send[coregraphics.IOSurfaceRef](a.ID, objc.Sel("copyDictionaryToIOSurface:copiedDataSize:createdIOSID:"), iOSurface, size, iosid)
	return coregraphics.IOSurfaceRef(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/copyErrorValue:
func (a ANEVirtualClient) CopyErrorValue(value unsafe.Pointer) {
	objc.Send[objc.ID](a.ID, objc.Sel("copyErrorValue:"), value)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/copyErrorValue:vmData:
func (a ANEVirtualClient) CopyErrorValueVmData(value objectivec.IObject, data unsafe.Pointer) {
	objc.Send[objc.ID](a.ID, objc.Sel("copyErrorValue:vmData:"), value, data)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/copyFilesInDirectoryToIOSurfaces:ioSurfaceRefs:ioSurfaceSizes:fileNames:
func (a ANEVirtualClient) CopyFilesInDirectoryToIOSurfacesIoSurfaceRefsIoSurfaceSizesFileNames(iOSurfaces objectivec.IObject, refs objectivec.IObject, sizes objectivec.IObject, names objectivec.IObject) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("copyFilesInDirectoryToIOSurfaces:ioSurfaceRefs:ioSurfaceSizes:fileNames:"), iOSurfaces, refs, sizes, names)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/copyModel:options:vmData:
func (a ANEVirtualClient) CopyModelOptionsVmData(model objectivec.IObject, options objectivec.IObject, data unsafe.Pointer) {
	objc.Send[objc.ID](a.ID, objc.Sel("copyModel:options:vmData:"), model, options, data)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/copyModelMetaData:options:dictionary:vmData:
func (a ANEVirtualClient) CopyModelMetaDataOptionsDictionaryVmData(data objectivec.IObject, options objectivec.IObject, dictionary objectivec.IObject, data2 unsafe.Pointer) {
	objc.Send[objc.ID](a.ID, objc.Sel("copyModelMetaData:options:dictionary:vmData:"), data, options, dictionary, data2)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/copyModelOptionFiles:options:dictionary:vmData:
func (a ANEVirtualClient) CopyModelOptionFilesOptionsDictionaryVmData(files objectivec.IObject, options objectivec.IObject, dictionary objectivec.IObject, data unsafe.Pointer) {
	objc.Send[objc.ID](a.ID, objc.Sel("copyModelOptionFiles:options:dictionary:vmData:"), files, options, dictionary, data)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/copyModelOptionFiles:options:vmData:
func (a ANEVirtualClient) CopyModelOptionFilesOptionsVmData(files objectivec.IObject, options objectivec.IObject, data unsafe.Pointer) {
	objc.Send[objc.ID](a.ID, objc.Sel("copyModelOptionFiles:options:vmData:"), files, options, data)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/copyOptions:dictionary:vmData:
func (a ANEVirtualClient) CopyOptionsDictionaryVmData(options objectivec.IObject, dictionary objectivec.IObject, data unsafe.Pointer) {
	objc.Send[objc.ID](a.ID, objc.Sel("copyOptions:dictionary:vmData:"), options, dictionary, data)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/copyOptions:vmData:
func (a ANEVirtualClient) CopyOptionsVmData(options objectivec.IObject, data unsafe.Pointer) {
	objc.Send[objc.ID](a.ID, objc.Sel("copyOptions:vmData:"), options, data)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/copyToIOSurface:length:ioSID:
func (a ANEVirtualClient) CopyToIOSurfaceLengthIoSID(iOSurface objectivec.IObject, length uint64, sid unsafe.Pointer) coregraphics.IOSurfaceRef {
	rv := objc.Send[coregraphics.IOSurfaceRef](a.ID, objc.Sel("copyToIOSurface:length:ioSID:"), iOSurface, length, sid)
	return coregraphics.IOSurfaceRef(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/copyToIOSurface:size:ioSID:
func (a ANEVirtualClient) CopyToIOSurfaceSizeIoSID(iOSurface []byte, size uint64, sid unsafe.Pointer) coregraphics.IOSurfaceRef {
	rv := objc.Send[coregraphics.IOSurfaceRef](a.ID, objc.Sel("copyToIOSurface:size:ioSID:"), unsafe.Pointer(unsafe.SliceData(iOSurface)), size, sid)
	return coregraphics.IOSurfaceRef(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/doEvaluateWithModel:options:request:qos:completionEvent:error:
func (a ANEVirtualClient) DoEvaluateWithModelOptionsRequestQosCompletionEventError(model objectivec.IObject, options objectivec.IObject, request objectivec.IObject, qos uint32, event objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("doEvaluateWithModel:options:request:qos:completionEvent:error:"), model, options, request, qos, event, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("doEvaluateWithModel:options:request:qos:completionEvent:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/doEvaluateWithModelLegacy:options:request:qos:completionEvent:error:
func (a ANEVirtualClient) DoEvaluateWithModelLegacyOptionsRequestQosCompletionEventError(legacy objectivec.IObject, options objectivec.IObject, request objectivec.IObject, qos uint32, event objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("doEvaluateWithModelLegacy:options:request:qos:completionEvent:error:"), legacy, options, request, qos, event, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("doEvaluateWithModelLegacy:options:request:qos:completionEvent:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/doJsonParsingMatchWeightName:
func (a ANEVirtualClient) DoJsonParsingMatchWeightName(name objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("doJsonParsingMatchWeightName:"), name)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/doMapIOSurfacesWithModel:request:cacheInference:error:
func (a ANEVirtualClient) DoMapIOSurfacesWithModelRequestCacheInferenceError(model objectivec.IObject, request objectivec.IObject, inference bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("doMapIOSurfacesWithModel:request:cacheInference:error:"), model, request, inference, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("doMapIOSurfacesWithModel:request:cacheInference:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/echo:
func (a ANEVirtualClient) Echo(echo objectivec.IObject) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("echo:"), echo)
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/endRealTimeTask
func (a ANEVirtualClient) EndRealTimeTask() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("endRealTimeTask"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/evaluateWithModel:options:request:qos:error:
func (a ANEVirtualClient) EvaluateWithModelOptionsRequestQosError(model objectivec.IObject, options objectivec.IObject, request objectivec.IObject, qos uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("evaluateWithModel:options:request:qos:error:"), model, options, request, qos, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("evaluateWithModel:options:request:qos:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/exchangeBuildVersionInfo
func (a ANEVirtualClient) ExchangeBuildVersionInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("exchangeBuildVersionInfo"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/getDeviceInfo
func (a ANEVirtualClient) GetDeviceInfo() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("getDeviceInfo"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/getModelAttribute:
func (a ANEVirtualClient) GetModelAttribute(attribute unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("getModelAttribute:"), attribute)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/getValidateNetworkVersion
func (a ANEVirtualClient) GetValidateNetworkVersion() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("getValidateNetworkVersion"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/hasANE
func (a ANEVirtualClient) HasANE() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("hasANE"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/hostBuildVersionStr
func (a ANEVirtualClient) HostBuildVersionStr() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("hostBuildVersionStr"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/isInternalBuild
func (a ANEVirtualClient) IsInternalBuild() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isInternalBuild"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/loadModel:options:qos:error:
func (a ANEVirtualClient) LoadModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("loadModel:options:qos:error:"), model, options, qos, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("loadModel:options:qos:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/loadModelNewInstance:options:modelInstParams:qos:error:
func (a ANEVirtualClient) LoadModelNewInstanceOptionsModelInstParamsQosError(instance objectivec.IObject, options objectivec.IObject, params objectivec.IObject, qos uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("loadModelNewInstance:options:modelInstParams:qos:error:"), instance, options, params, qos, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("loadModelNewInstance:options:modelInstParams:qos:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/loadModelNewInstanceLegacy:options:modelInstParams:qos:error:
func (a ANEVirtualClient) LoadModelNewInstanceLegacyOptionsModelInstParamsQosError(legacy objectivec.IObject, options objectivec.IObject, params objectivec.IObject, qos uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("loadModelNewInstanceLegacy:options:modelInstParams:qos:error:"), legacy, options, params, qos, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("loadModelNewInstanceLegacy:options:modelInstParams:qos:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/mapIOSurfacesWithModel:request:cacheInference:error:
func (a ANEVirtualClient) MapIOSurfacesWithModelRequestCacheInferenceError(model objectivec.IObject, request objectivec.IObject, inference bool) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("mapIOSurfacesWithModel:request:cacheInference:error:"), model, request, inference, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("mapIOSurfacesWithModel:request:cacheInference:error: returned NO with nil NSError")
	}
	return rv, nil

}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/negotiatedCapabilityMask
func (a ANEVirtualClient) NegotiatedCapabilityMask() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("negotiatedCapabilityMask"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/negotiatedDataInterfaceVersion
func (a ANEVirtualClient) NegotiatedDataInterfaceVersion() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("negotiatedDataInterfaceVersion"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/numANECores
func (a ANEVirtualClient) NumANECores() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("numANECores"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/numANEs
func (a ANEVirtualClient) NumANEs() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("numANEs"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/outputDictIOSurfaceSize
func (a ANEVirtualClient) OutputDictIOSurfaceSize() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("outputDictIOSurfaceSize"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/parallelDecompressedData:
func (a ANEVirtualClient) ParallelDecompressedData(data objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("parallelDecompressedData:"), data)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/printDictionary:
func (a ANEVirtualClient) PrintDictionary(dictionary objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("printDictionary:"), dictionary)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/purgeCompiledModel:
func (a ANEVirtualClient) PurgeCompiledModel(model objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("purgeCompiledModel:"), model)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/purgeCompiledModelMatchingHash:
func (a ANEVirtualClient) PurgeCompiledModelMatchingHash(hash objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("purgeCompiledModelMatchingHash:"), hash)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/readWeightFilename:
func (a ANEVirtualClient) ReadWeightFilename(filename objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("readWeightFilename:"), filename)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/releaseIOSurfaces:
func (a ANEVirtualClient) ReleaseIOSurfaces(iOSurfaces unsafe.Pointer) {
	objc.Send[objc.ID](a.ID, objc.Sel("releaseIOSurfaces:"), iOSurfaces)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/sendGuestBuildVersion
func (a ANEVirtualClient) SendGuestBuildVersion() {
	objc.Send[objc.ID](a.ID, objc.Sel("sendGuestBuildVersion"))
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/sessionHintWithModel:hint:options:report:error:
func (a ANEVirtualClient) SessionHintWithModelHintOptionsReportError(model objectivec.IObject, hint objectivec.IObject, options objectivec.IObject, report objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("sessionHintWithModel:hint:options:report:error:"), model, hint, options, report, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("sessionHintWithModel:hint:options:report:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/transferFileToHostWithPath:withChunkSize:withUUID:withModelInputPath:overWriteFileNameWith:
func (a ANEVirtualClient) TransferFileToHostWithPathWithChunkSizeWithUUIDWithModelInputPathOverWriteFileNameWith(path objectivec.IObject, size uint32, uuid objectivec.IObject, path2 objectivec.IObject, with objectivec.IObject) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("transferFileToHostWithPath:withChunkSize:withUUID:withModelInputPath:overWriteFileNameWith:"), path, size, uuid, path2, with)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/unloadModel:options:qos:error:
func (a ANEVirtualClient) UnloadModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("unloadModel:options:qos:error:"), model, options, qos, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("unloadModel:options:qos:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/updateError:error:
func (a ANEVirtualClient) UpdateErrorError(error_ unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("updateError:error:"), error_, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateError:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/updatePerformanceStats:
func (a ANEVirtualClient) UpdatePerformanceStats(stats unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("updatePerformanceStats:"), stats)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/validateEnvironmentForPrecompiledBinarySupport
func (a ANEVirtualClient) ValidateEnvironmentForPrecompiledBinarySupport() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("validateEnvironmentForPrecompiledBinarySupport"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/validateNetworkCreate:uuid:function:directoryPath:scratchPadPath:milTextData:
func (a ANEVirtualClient) ValidateNetworkCreateUuidFunctionDirectoryPathScratchPadPathMilTextData(create uint64, uuid objectivec.IObject, function objectivec.IObject, path objectivec.IObject, path2 objectivec.IObject, data objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("validateNetworkCreate:uuid:function:directoryPath:scratchPadPath:milTextData:"), create, uuid, function, path, path2, data)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/validateNetworkCreateMLIR:validation_params:
func (a ANEVirtualClient) ValidateNetworkCreateMLIRValidation_params(mlir uint64, validation_params objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("validateNetworkCreateMLIR:validation_params:"), mlir, validation_params)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/initWithSingletonAccess
func (a ANEVirtualClient) InitWithSingletonAccess() ANEVirtualClient {
	rv := objc.Send[ANEVirtualClient](a.ID, objc.Sel("initWithSingletonAccess"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/copyDictionaryDataToStruct:dictionary:
func (_ANEVirtualClientClass ANEVirtualClientClass) CopyDictionaryDataToStructDictionary(struct_ unsafe.Pointer, dictionary objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("copyDictionaryDataToStruct:dictionary:"), struct_, dictionary)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/copyLLIRBundleToIOSurface:writtenDataSize:
func (_ANEVirtualClientClass ANEVirtualClientClass) CopyLLIRBundleToIOSurfaceWrittenDataSize(iOSurface objectivec.IObject, size unsafe.Pointer) coregraphics.IOSurfaceRef {
	rv := objc.Send[coregraphics.IOSurfaceRef](objc.ID(_ANEVirtualClientClass.class), objc.Sel("copyLLIRBundleToIOSurface:writtenDataSize:"), iOSurface, size)
	return coregraphics.IOSurfaceRef(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/createIOSurface:ioSID:
func (_ANEVirtualClientClass ANEVirtualClientClass) CreateIOSurfaceIoSID(iOSurface uint64, sid unsafe.Pointer) coregraphics.IOSurfaceRef {
	rv := objc.Send[coregraphics.IOSurfaceRef](objc.ID(_ANEVirtualClientClass.class), objc.Sel("createIOSurface:ioSID:"), iOSurface, sid)
	return coregraphics.IOSurfaceRef(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/dictionaryGetInt64ForKey:key:
func (_ANEVirtualClientClass ANEVirtualClientClass) DictionaryGetInt64ForKeyKey(key objectivec.IObject, key2 objectivec.IObject) int64 {
	rv := objc.Send[int64](objc.ID(_ANEVirtualClientClass.class), objc.Sel("dictionaryGetInt64ForKey:key:"), key, key2)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/dictionaryGetInt8ForKey:key:
func (_ANEVirtualClientClass ANEVirtualClientClass) DictionaryGetInt8ForKeyKey(key objectivec.IObject, key2 objectivec.IObject) int8 {
	rv := objc.Send[int8](objc.ID(_ANEVirtualClientClass.class), objc.Sel("dictionaryGetInt8ForKey:key:"), key, key2)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/dictionaryGetNSStringForKey:key:
func (_ANEVirtualClientClass ANEVirtualClientClass) DictionaryGetNSStringForKeyKey(key objectivec.IObject, key2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("dictionaryGetNSStringForKey:key:"), key, key2)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/dictionaryGetUInt32ForKey:key:
func (_ANEVirtualClientClass ANEVirtualClientClass) DictionaryGetUInt32ForKeyKey(key objectivec.IObject, key2 objectivec.IObject) uint32 {
	rv := objc.Send[uint32](objc.ID(_ANEVirtualClientClass.class), objc.Sel("dictionaryGetUInt32ForKey:key:"), key, key2)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/dictionaryGetUInt64ForKey:key:
func (_ANEVirtualClientClass ANEVirtualClientClass) DictionaryGetUInt64ForKeyKey(key objectivec.IObject, key2 objectivec.IObject) uint64 {
	rv := objc.Send[uint64](objc.ID(_ANEVirtualClientClass.class), objc.Sel("dictionaryGetUInt64ForKey:key:"), key, key2)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/freeModelFileIOSurfaces:
func (_ANEVirtualClientClass ANEVirtualClientClass) FreeModelFileIOSurfaces(iOSurfaces objectivec.IObject) {
	objc.Send[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("freeModelFileIOSurfaces:"), iOSurfaces)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/getCFDictionaryFromIOSurface:dataLength:
func (_ANEVirtualClientClass ANEVirtualClientClass) GetCFDictionaryFromIOSurfaceDataLength(iOSurface coregraphics.IOSurfaceRef, length uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("getCFDictionaryFromIOSurface:dataLength:"), iOSurface, length)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/getCodeSigningIdentity
func (_ANEVirtualClientClass ANEVirtualClientClass) GetCodeSigningIdentity() objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("getCodeSigningIdentity"))
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/getDictionaryWithJSONEncodingFromIOSurface:withArchivedDataSize:
func (_ANEVirtualClientClass ANEVirtualClientClass) GetDictionaryWithJSONEncodingFromIOSurfaceWithArchivedDataSize(iOSurface coregraphics.IOSurfaceRef, size uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("getDictionaryWithJSONEncodingFromIOSurface:withArchivedDataSize:"), iOSurface, size)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/getObjectFromIOSurface:classType:length:
func (_ANEVirtualClientClass ANEVirtualClientClass) GetObjectFromIOSurfaceClassTypeLength(iOSurface coregraphics.IOSurfaceRef, type_ objc.Class, length uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("getObjectFromIOSurface:classType:length:"), iOSurface, type_, length)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/printIOSurfaceDataInBytes:
func (_ANEVirtualClientClass ANEVirtualClientClass) PrintIOSurfaceDataInBytes(bytes coregraphics.IOSurfaceRef) {
	objc.Send[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("printIOSurfaceDataInBytes:"), bytes)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/printStruct:
func (_ANEVirtualClientClass ANEVirtualClientClass) PrintStruct(struct_ unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("printStruct:"), struct_)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/setCodeSigningIdentity:
func (_ANEVirtualClientClass ANEVirtualClientClass) SetCodeSigningIdentity(identity objectivec.IObject) bool {
	rv := objc.Send[bool](objc.ID(_ANEVirtualClientClass.class), objc.Sel("setCodeSigningIdentity:"), identity)
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/sharedConnection
func (_ANEVirtualClientClass ANEVirtualClientClass) SharedConnection() *ANEVirtualClient {
	rv := objc.Send[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("sharedConnection"))
	if rv == 0 {
		return nil
	}
	val := ANEVirtualClientFromID(rv)
	return &val
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/shouldUsePrecompiledPath:options:shouldUseChunking:chunkingThreshold:
func (_ANEVirtualClientClass ANEVirtualClientClass) ShouldUsePrecompiledPathOptionsShouldUseChunkingChunkingThreshold(path objectivec.IObject, options objectivec.IObject, threshold uint64) (bool, bool) {
	var chunking bool
	rv := objc.Send[bool](objc.ID(_ANEVirtualClientClass.class), objc.Sel("shouldUsePrecompiledPath:options:shouldUseChunking:chunkingThreshold:"), path, options, unsafe.Pointer(&chunking), threshold)
	return chunking, rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/updateError:errorLength:error:
func (_ANEVirtualClientClass ANEVirtualClientClass) UpdateErrorErrorLengthError(error_ coregraphics.IOSurfaceRef, length uint64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_ANEVirtualClientClass.class), objc.Sel("updateError:errorLength:error:"), error_, length, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateError:errorLength:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/updateError:errorLength:errorCode:error:
func (_ANEVirtualClientClass ANEVirtualClientClass) UpdateErrorErrorLengthErrorCodeError(error_ coregraphics.IOSurfaceRef, length uint64, code int64) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_ANEVirtualClientClass.class), objc.Sel("updateError:errorLength:errorCode:error:"), error_, length, code, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateError:errorLength:errorCode:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/updatePerformanceStats:performanceStatsLength:perfStatsRawIOSurfaceRef:performanceStatsRawLength:hwExecutionTime:
func (_ANEVirtualClientClass ANEVirtualClientClass) UpdatePerformanceStatsPerformanceStatsLengthPerfStatsRawIOSurfaceRefPerformanceStatsRawLengthHwExecutionTime(stats coregraphics.IOSurfaceRef, length uint64, ref coregraphics.IOSurfaceRef, length2 uint64, time uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("updatePerformanceStats:performanceStatsLength:perfStatsRawIOSurfaceRef:performanceStatsRawLength:hwExecutionTime:"), stats, length, ref, length2, time)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/connect
func (a ANEVirtualClient) Connect() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("connect"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEVirtualClient/queue
func (a ANEVirtualClient) Queue() objectivec.Object {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("queue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

