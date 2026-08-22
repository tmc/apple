// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"errors"
	"sync"
	"unsafe"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iosurface"
	"github.com/tmc/apple/objc"
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

// Class returns the underlying Objective-C class pointer.
func (ac ANEVirtualClientClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEVirtualClientClass) Alloc() ANEVirtualClient {
	rv := objc.SendIfResponds[ANEVirtualClient](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ANEVirtualClient.AneArchitectureTypeStr]
//   - [ANEVirtualClient.AneBoardtype]
//   - [ANEVirtualClient.AneSubTypeAndVariant]
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
type ANEVirtualClient struct {
	objectivec.Object
}

// ANEVirtualClientFromID constructs a [ANEVirtualClient] from an objc.ID.
func ANEVirtualClientFromID(id objc.ID) ANEVirtualClient {
	return ANEVirtualClient{objectivec.Object{ID: id}}
}

// Ensure ANEVirtualClient implements IANEVirtualClient.
var _ IANEVirtualClient = ANEVirtualClient{}

// An interface definition for the [ANEVirtualClient] class.
//
// # Methods
//
//   - [IANEVirtualClient.AneArchitectureTypeStr]
//   - [IANEVirtualClient.AneBoardtype]
//   - [IANEVirtualClient.AneSubTypeAndVariant]
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
type IANEVirtualClient interface {
	objectivec.IObject

	// Topic: Methods

	AneArchitectureTypeStr() objectivec.IObject
	AneBoardtype() int64
	AneSubTypeAndVariant() objectivec.IObject
	BeginRealTimeTask() bool
	CallIOUserClientInParamsOutParams(client uint32, params *VirtANEModel, params2 *VirtANEModel) bool
	CallIOUserClientWithDictionaryInDictionaryError(dictionary uint32, dictionary2 corefoundation.CFDictionaryRef) (corefoundation.CFDictionaryRef, error)
	CheckKernReturnValueSelectorOutParams(value int, selector uint32, params *VirtANEModel)
	CompileModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error)
	CompiledModelExistsFor(for_ objectivec.IObject) bool
	CompiledModelExistsMatchingHash(hash objectivec.IObject) bool
	Connect() uint32
	CopyAllModelFilesDictionaryIoSurfaceRefs(files objectivec.IObject, dictionary corefoundation.CFDictionaryRef, refs corefoundation.CFArrayRef) bool
	CopyDictionaryToIOSurfaceCopiedDataSizeCreatedIOSID(iOSurface objectivec.IObject, size *uint64, iosid *uint32) iosurface.IOSurfaceRef
	CopyErrorValue(value *VMData)
	CopyErrorValueVmData(value corefoundation.CFDictionaryRef, data *VMData)
	CopyFilesInDirectoryToIOSurfacesIoSurfaceRefsIoSurfaceSizesFileNames(iOSurfaces objectivec.IObject, refs corefoundation.CFArrayRef, sizes objectivec.IObject, names objectivec.IObject) bool
	CopyModelOptionsVmData(model objectivec.IObject, options objectivec.IObject, data *VMData)
	CopyModelMetaDataOptionsDictionaryVmData(data objectivec.IObject, options objectivec.IObject, dictionary corefoundation.CFDictionaryRef, data2 *VMData)
	CopyModelOptionFilesOptionsDictionaryVmData(files objectivec.IObject, options objectivec.IObject, dictionary corefoundation.CFDictionaryRef, data *VMData)
	CopyModelOptionFilesOptionsVmData(files objectivec.IObject, options objectivec.IObject, data *VMData)
	CopyOptionsDictionaryVmData(options objectivec.IObject, dictionary corefoundation.CFDictionaryRef, data *VMData)
	CopyOptionsVmData(options objectivec.IObject, data *VMData)
	CopyToIOSurfaceLengthIoSID(iOSurface objectivec.IObject, length uint64, sid *uint32) iosurface.IOSurfaceRef
	CopyToIOSurfaceSizeIoSID(iOSurface string, size uint64, sid *uint32) iosurface.IOSurfaceRef
	DoEvaluateWithModelOptionsRequestQosCompletionEventError(model objectivec.IObject, options objectivec.IObject, request objectivec.IObject, qos uint32, event objectivec.IObject) (bool, error)
	DoEvaluateWithModelLegacyOptionsRequestQosCompletionEventError(legacy objectivec.IObject, options objectivec.IObject, request objectivec.IObject, qos uint32, event objectivec.IObject) (bool, error)
	DoJsonParsingMatchWeightName(name objectivec.IObject) objectivec.IObject
	DoMapIOSurfacesWithModelRequestCacheInferenceError(model objectivec.IObject, request objectivec.IObject, inference bool) (bool, error)
	Echo(echo objectivec.IObject) bool
	EndRealTimeTask() bool
	EvaluateWithModelOptionsRequestQosError(model objectivec.IObject, options objectivec.IObject, request objectivec.IObject, qos uint32) (bool, error)
	ExchangeBuildVersionInfo() BuildVersionInfo
	GetDeviceInfo() DeviceExtendedInfo
	GetModelAttribute(attribute *VMData) objectivec.IObject
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
	PrintDictionary(dictionary corefoundation.CFDictionaryRef)
	PurgeCompiledModel(model objectivec.IObject)
	PurgeCompiledModelMatchingHash(hash objectivec.IObject)
	Queue() objectivec.Object
	ReadWeightFilename(filename objectivec.IObject) objectivec.IObject
	ReleaseIOSurfaces(iOSurfaces *VMData)
	SendGuestBuildVersion()
	SessionHintWithModelHintOptionsReportError(model objectivec.IObject, hint objectivec.IObject, options objectivec.IObject, report objectivec.IObject) (bool, error)
	TransferFileToHostWithPathWithChunkSizeWithUUIDWithModelInputPathOverWriteFileNameWith(path objectivec.IObject, size uint32, uuid objectivec.IObject, path2 objectivec.IObject, with objectivec.IObject) bool
	UnloadModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error)
	UpdateErrorError(error_ *VMData) (bool, error)
	UpdatePerformanceStats(stats *VMData) objectivec.IObject
	ValidateEnvironmentForPrecompiledBinarySupport() bool
	ValidateNetworkCreateUuidFunctionDirectoryPathScratchPadPathMilTextData(create uint64, uuid objectivec.IObject, function objectivec.IObject, path objectivec.IObject, path2 objectivec.IObject, data objectivec.IObject) corefoundation.CFDictionaryRef
	ValidateNetworkCreateMLIRValidation_params(mlir uint64, validation_params corefoundation.CFDictionaryRef) corefoundation.CFDictionaryRef
	InitWithSingletonAccess() ANEVirtualClient
}

// Init initializes the instance.
func (a ANEVirtualClient) Init() ANEVirtualClient {
	rv := objc.SendIfResponds[ANEVirtualClient](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEVirtualClient) Autorelease() ANEVirtualClient {
	rv := objc.SendIfResponds[ANEVirtualClient](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEVirtualClient creates a new ANEVirtualClient instance.
func NewANEVirtualClient() ANEVirtualClient {
	class := getANEVirtualClientClass()
	rv := objc.SendIfResponds[ANEVirtualClient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewANEVirtualClientWithSingletonAccess() ANEVirtualClient {
	instance := getANEVirtualClientClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithSingletonAccess"))
	return ANEVirtualClientFromID(rv)
}

func (a ANEVirtualClient) AneArchitectureTypeStr() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("aneArchitectureTypeStr"))
	return objectivec.Object{ID: rv}
}
func (a ANEVirtualClient) AneBoardtype() int64 {
	rv := objc.SendIfResponds[int64](a.ID, objc.Sel("aneBoardtype"))
	return rv
}
func (a ANEVirtualClient) AneSubTypeAndVariant() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("aneSubTypeAndVariant"))
	return objectivec.Object{ID: rv}
}
func (a ANEVirtualClient) BeginRealTimeTask() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("beginRealTimeTask"))
	return rv
}
func (a ANEVirtualClient) CallIOUserClientInParamsOutParams(client uint32, params *VirtANEModel, params2 *VirtANEModel) bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("callIOUserClient:inParams:outParams:"), client, unsafe.Pointer(params), unsafe.Pointer(params2))
	return rv
}
func (a ANEVirtualClient) CallIOUserClientWithDictionaryInDictionaryError(dictionary uint32, dictionary2 corefoundation.CFDictionaryRef) (corefoundation.CFDictionaryRef, error) {
	var errorPtr objc.ID
	rv := objc.Send[corefoundation.CFDictionaryRef](a.ID, objc.Sel("callIOUserClientWithDictionary:inDictionary:error:"), dictionary, dictionary2, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return *new(corefoundation.CFDictionaryRef), foundation.NSErrorFrom(errorPtr)
	}
	return rv, nil

}
func (a ANEVirtualClient) CheckKernReturnValueSelectorOutParams(value int, selector uint32, params *VirtANEModel) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("checkKernReturnValue:selector:outParams:"), value, selector, unsafe.Pointer(params))
}
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
func (a ANEVirtualClient) CompiledModelExistsFor(for_ objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("compiledModelExistsFor:"), for_)
	return rv
}
func (a ANEVirtualClient) CompiledModelExistsMatchingHash(hash objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("compiledModelExistsMatchingHash:"), hash)
	return rv
}
func (a ANEVirtualClient) CopyAllModelFilesDictionaryIoSurfaceRefs(files objectivec.IObject, dictionary corefoundation.CFDictionaryRef, refs corefoundation.CFArrayRef) bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("copyAllModelFiles:dictionary:ioSurfaceRefs:"), files, dictionary, refs)
	return rv
}
func (a ANEVirtualClient) CopyDictionaryToIOSurfaceCopiedDataSizeCreatedIOSID(iOSurface objectivec.IObject, size *uint64, iosid *uint32) iosurface.IOSurfaceRef {
	rv := objc.SendIfResponds[iosurface.IOSurfaceRef](a.ID, objc.Sel("copyDictionaryToIOSurface:copiedDataSize:createdIOSID:"), iOSurface, unsafe.Pointer(size), unsafe.Pointer(iosid))
	return iosurface.IOSurfaceRef(rv)
}
func (a ANEVirtualClient) CopyErrorValue(value *VMData) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("copyErrorValue:"), unsafe.Pointer(value))
}
func (a ANEVirtualClient) CopyErrorValueVmData(value corefoundation.CFDictionaryRef, data *VMData) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("copyErrorValue:vmData:"), value, unsafe.Pointer(data))
}
func (a ANEVirtualClient) CopyFilesInDirectoryToIOSurfacesIoSurfaceRefsIoSurfaceSizesFileNames(iOSurfaces objectivec.IObject, refs corefoundation.CFArrayRef, sizes objectivec.IObject, names objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("copyFilesInDirectoryToIOSurfaces:ioSurfaceRefs:ioSurfaceSizes:fileNames:"), iOSurfaces, refs, sizes, names)
	return rv
}
func (a ANEVirtualClient) CopyModelOptionsVmData(model objectivec.IObject, options objectivec.IObject, data *VMData) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("copyModel:options:vmData:"), model, options, unsafe.Pointer(data))
}
func (a ANEVirtualClient) CopyModelMetaDataOptionsDictionaryVmData(data objectivec.IObject, options objectivec.IObject, dictionary corefoundation.CFDictionaryRef, data2 *VMData) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("copyModelMetaData:options:dictionary:vmData:"), data, options, dictionary, unsafe.Pointer(data2))
}
func (a ANEVirtualClient) CopyModelOptionFilesOptionsDictionaryVmData(files objectivec.IObject, options objectivec.IObject, dictionary corefoundation.CFDictionaryRef, data *VMData) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("copyModelOptionFiles:options:dictionary:vmData:"), files, options, dictionary, unsafe.Pointer(data))
}
func (a ANEVirtualClient) CopyModelOptionFilesOptionsVmData(files objectivec.IObject, options objectivec.IObject, data *VMData) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("copyModelOptionFiles:options:vmData:"), files, options, unsafe.Pointer(data))
}
func (a ANEVirtualClient) CopyOptionsDictionaryVmData(options objectivec.IObject, dictionary corefoundation.CFDictionaryRef, data *VMData) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("copyOptions:dictionary:vmData:"), options, dictionary, unsafe.Pointer(data))
}
func (a ANEVirtualClient) CopyOptionsVmData(options objectivec.IObject, data *VMData) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("copyOptions:vmData:"), options, unsafe.Pointer(data))
}
func (a ANEVirtualClient) CopyToIOSurfaceLengthIoSID(iOSurface objectivec.IObject, length uint64, sid *uint32) iosurface.IOSurfaceRef {
	rv := objc.SendIfResponds[iosurface.IOSurfaceRef](a.ID, objc.Sel("copyToIOSurface:length:ioSID:"), iOSurface, length, unsafe.Pointer(sid))
	return iosurface.IOSurfaceRef(rv)
}
func (a ANEVirtualClient) CopyToIOSurfaceSizeIoSID(iOSurface string, size uint64, sid *uint32) iosurface.IOSurfaceRef {
	rv := objc.SendIfResponds[iosurface.IOSurfaceRef](a.ID, objc.Sel("copyToIOSurface:size:ioSID:"), unsafe.Pointer(unsafe.StringData(iOSurface+"\x00")), size, unsafe.Pointer(sid))
	return iosurface.IOSurfaceRef(rv)
}
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
func (a ANEVirtualClient) DoJsonParsingMatchWeightName(name objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("doJsonParsingMatchWeightName:"), name)
	return objectivec.Object{ID: rv}
}
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
func (a ANEVirtualClient) Echo(echo objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("echo:"), echo)
	return rv
}
func (a ANEVirtualClient) EndRealTimeTask() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("endRealTimeTask"))
	return rv
}
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
func (a ANEVirtualClient) ExchangeBuildVersionInfo() BuildVersionInfo {
	rv := objc.SendIfResponds[BuildVersionInfo](a.ID, objc.Sel("exchangeBuildVersionInfo"))
	return BuildVersionInfo(rv)
}
func (a ANEVirtualClient) GetDeviceInfo() DeviceExtendedInfo {
	rv := objc.SendIfResponds[DeviceExtendedInfo](a.ID, objc.Sel("getDeviceInfo"))
	return DeviceExtendedInfo(rv)
}
func (a ANEVirtualClient) GetModelAttribute(attribute *VMData) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("getModelAttribute:"), unsafe.Pointer(attribute))
	return objectivec.Object{ID: rv}
}
func (a ANEVirtualClient) GetValidateNetworkVersion() uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("getValidateNetworkVersion"))
	return rv
}
func (a ANEVirtualClient) HasANE() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("hasANE"))
	return rv
}
func (a ANEVirtualClient) HostBuildVersionStr() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("hostBuildVersionStr"))
	return objectivec.Object{ID: rv}
}
func (a ANEVirtualClient) IsInternalBuild() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("isInternalBuild"))
	return rv
}
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
func (a ANEVirtualClient) NegotiatedCapabilityMask() uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("negotiatedCapabilityMask"))
	return rv
}
func (a ANEVirtualClient) NegotiatedDataInterfaceVersion() uint32 {
	rv := objc.SendIfResponds[uint32](a.ID, objc.Sel("negotiatedDataInterfaceVersion"))
	return rv
}
func (a ANEVirtualClient) NumANECores() uint32 {
	rv := objc.SendIfResponds[uint32](a.ID, objc.Sel("numANECores"))
	return rv
}
func (a ANEVirtualClient) NumANEs() uint32 {
	rv := objc.SendIfResponds[uint32](a.ID, objc.Sel("numANEs"))
	return rv
}
func (a ANEVirtualClient) OutputDictIOSurfaceSize() uint64 {
	rv := objc.SendIfResponds[uint64](a.ID, objc.Sel("outputDictIOSurfaceSize"))
	return rv
}
func (a ANEVirtualClient) ParallelDecompressedData(data objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("parallelDecompressedData:"), data)
	return objectivec.Object{ID: rv}
}
func (a ANEVirtualClient) PrintDictionary(dictionary corefoundation.CFDictionaryRef) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("printDictionary:"), dictionary)
}
func (a ANEVirtualClient) PurgeCompiledModel(model objectivec.IObject) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("purgeCompiledModel:"), model)
}
func (a ANEVirtualClient) PurgeCompiledModelMatchingHash(hash objectivec.IObject) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("purgeCompiledModelMatchingHash:"), hash)
}
func (a ANEVirtualClient) ReadWeightFilename(filename objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("readWeightFilename:"), filename)
	return objectivec.Object{ID: rv}
}
func (a ANEVirtualClient) ReleaseIOSurfaces(iOSurfaces *VMData) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("releaseIOSurfaces:"), unsafe.Pointer(iOSurfaces))
}
func (a ANEVirtualClient) SendGuestBuildVersion() {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("sendGuestBuildVersion"))
}
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
func (a ANEVirtualClient) TransferFileToHostWithPathWithChunkSizeWithUUIDWithModelInputPathOverWriteFileNameWith(path objectivec.IObject, size uint32, uuid objectivec.IObject, path2 objectivec.IObject, with objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("transferFileToHostWithPath:withChunkSize:withUUID:withModelInputPath:overWriteFileNameWith:"), path, size, uuid, path2, with)
	return rv
}
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
func (a ANEVirtualClient) UpdateErrorError(error_ *VMData) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("updateError:error:"), unsafe.Pointer(error_), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("updateError:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (a ANEVirtualClient) UpdatePerformanceStats(stats *VMData) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("updatePerformanceStats:"), unsafe.Pointer(stats))
	return objectivec.Object{ID: rv}
}
func (a ANEVirtualClient) ValidateEnvironmentForPrecompiledBinarySupport() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("validateEnvironmentForPrecompiledBinarySupport"))
	return rv
}
func (a ANEVirtualClient) ValidateNetworkCreateUuidFunctionDirectoryPathScratchPadPathMilTextData(create uint64, uuid objectivec.IObject, function objectivec.IObject, path objectivec.IObject, path2 objectivec.IObject, data objectivec.IObject) corefoundation.CFDictionaryRef {
	rv := objc.SendIfResponds[corefoundation.CFDictionaryRef](a.ID, objc.Sel("validateNetworkCreate:uuid:function:directoryPath:scratchPadPath:milTextData:"), create, uuid, function, path, path2, data)
	return corefoundation.CFDictionaryRef(rv)
}
func (a ANEVirtualClient) ValidateNetworkCreateMLIRValidation_params(mlir uint64, validation_params corefoundation.CFDictionaryRef) corefoundation.CFDictionaryRef {
	rv := objc.SendIfResponds[corefoundation.CFDictionaryRef](a.ID, objc.Sel("validateNetworkCreateMLIR:validation_params:"), mlir, validation_params)
	return corefoundation.CFDictionaryRef(rv)
}
func (a ANEVirtualClient) InitWithSingletonAccess() ANEVirtualClient {
	rv := objc.SendIfResponds[ANEVirtualClient](a.ID, objc.Sel("initWithSingletonAccess"))
	return rv
}

func (_ANEVirtualClientClass ANEVirtualClientClass) CopyDataToExistingIOSurfaceRef(data objectivec.IObject, ref iosurface.IOSurfaceRef) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_ANEVirtualClientClass.class), objc.Sel("copyData:toExistingIOSurfaceRef:"), data, ref)
	return rv
}
func (_ANEVirtualClientClass ANEVirtualClientClass) CopyDictionaryDataToStructDictionary(struct_ *VirtANEModel, dictionary corefoundation.CFDictionaryRef) {
	objc.SendIfResponds[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("copyDictionaryDataToStruct:dictionary:"), unsafe.Pointer(struct_), dictionary)
}
func (_ANEVirtualClientClass ANEVirtualClientClass) CopyLLIRBundleToIOSurfaceWrittenDataSize(iOSurface objectivec.IObject, size *uint64) iosurface.IOSurfaceRef {
	rv := objc.SendIfResponds[iosurface.IOSurfaceRef](objc.ID(_ANEVirtualClientClass.class), objc.Sel("copyLLIRBundleToIOSurface:writtenDataSize:"), iOSurface, unsafe.Pointer(size))
	return iosurface.IOSurfaceRef(rv)
}
func (_ANEVirtualClientClass ANEVirtualClientClass) CreateIOSurfaceIoSID(iOSurface uint64, sid *uint32) iosurface.IOSurfaceRef {
	rv := objc.SendIfResponds[iosurface.IOSurfaceRef](objc.ID(_ANEVirtualClientClass.class), objc.Sel("createIOSurface:ioSID:"), iOSurface, unsafe.Pointer(sid))
	return iosurface.IOSurfaceRef(rv)
}
func (_ANEVirtualClientClass ANEVirtualClientClass) DictionaryGetInt64ForKeyKey(key corefoundation.CFDictionaryRef, key2 corefoundation.CFStringRef) int64 {
	rv := objc.SendIfResponds[int64](objc.ID(_ANEVirtualClientClass.class), objc.Sel("dictionaryGetInt64ForKey:key:"), key, key2)
	return rv
}
func (_ANEVirtualClientClass ANEVirtualClientClass) DictionaryGetInt8ForKeyKey(key corefoundation.CFDictionaryRef, key2 corefoundation.CFStringRef) int8 {
	rv := objc.SendIfResponds[int8](objc.ID(_ANEVirtualClientClass.class), objc.Sel("dictionaryGetInt8ForKey:key:"), key, key2)
	return rv
}
func (_ANEVirtualClientClass ANEVirtualClientClass) DictionaryGetNSStringForKeyKey(key corefoundation.CFDictionaryRef, key2 corefoundation.CFStringRef) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("dictionaryGetNSStringForKey:key:"), key, key2)
	return objectivec.Object{ID: rv}
}
func (_ANEVirtualClientClass ANEVirtualClientClass) DictionaryGetUInt32ForKeyKey(key corefoundation.CFDictionaryRef, key2 corefoundation.CFStringRef) uint32 {
	rv := objc.SendIfResponds[uint32](objc.ID(_ANEVirtualClientClass.class), objc.Sel("dictionaryGetUInt32ForKey:key:"), key, key2)
	return rv
}
func (_ANEVirtualClientClass ANEVirtualClientClass) DictionaryGetUInt64ForKeyKey(key corefoundation.CFDictionaryRef, key2 corefoundation.CFStringRef) uint64 {
	rv := objc.SendIfResponds[uint64](objc.ID(_ANEVirtualClientClass.class), objc.Sel("dictionaryGetUInt64ForKey:key:"), key, key2)
	return rv
}
func (_ANEVirtualClientClass ANEVirtualClientClass) FreeModelFileIOSurfaces(iOSurfaces corefoundation.CFArrayRef) {
	objc.SendIfResponds[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("freeModelFileIOSurfaces:"), iOSurfaces)
}
func (_ANEVirtualClientClass ANEVirtualClientClass) GetCFDictionaryFromIOSurfaceDataLength(iOSurface iosurface.IOSurfaceRef, length uint64) corefoundation.CFDictionaryRef {
	rv := objc.SendIfResponds[corefoundation.CFDictionaryRef](objc.ID(_ANEVirtualClientClass.class), objc.Sel("getCFDictionaryFromIOSurface:dataLength:"), iOSurface, length)
	return corefoundation.CFDictionaryRef(rv)
}
func (_ANEVirtualClientClass ANEVirtualClientClass) GetCodeSigningIdentity() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("getCodeSigningIdentity"))
	return objectivec.Object{ID: rv}
}
func (_ANEVirtualClientClass ANEVirtualClientClass) GetDictionaryWithJSONEncodingFromIOSurfaceWithArchivedDataSize(iOSurface iosurface.IOSurfaceRef, size uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("getDictionaryWithJSONEncodingFromIOSurface:withArchivedDataSize:"), iOSurface, size)
	return objectivec.Object{ID: rv}
}
func (_ANEVirtualClientClass ANEVirtualClientClass) GetObjectFromIOSurfaceClassTypeLength(iOSurface iosurface.IOSurfaceRef, type_ objectivec.Class, length uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("getObjectFromIOSurface:classType:length:"), iOSurface, type_, length)
	return objectivec.Object{ID: rv}
}
func (_ANEVirtualClientClass ANEVirtualClientClass) PrintIOSurfaceDataInBytes(bytes iosurface.IOSurfaceRef) {
	objc.SendIfResponds[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("printIOSurfaceDataInBytes:"), bytes)
}
func (_ANEVirtualClientClass ANEVirtualClientClass) PrintStruct(struct_ *VirtANEModel) {
	objc.SendIfResponds[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("printStruct:"), unsafe.Pointer(struct_))
}
func (_ANEVirtualClientClass ANEVirtualClientClass) SetCodeSigningIdentity(identity corefoundation.CFDictionaryRef) bool {
	rv := objc.SendIfResponds[bool](objc.ID(_ANEVirtualClientClass.class), objc.Sel("setCodeSigningIdentity:"), identity)
	return rv
}
func (_ANEVirtualClientClass ANEVirtualClientClass) SharedConnection() ANEVirtualClient {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("sharedConnection"))
	return ANEVirtualClientFromID(rv)
}
func (_ANEVirtualClientClass ANEVirtualClientClass) ShouldUsePrecompiledPathOptionsShouldUseChunkingChunkingThreshold(path objectivec.IObject, options objectivec.IObject, threshold uint64) (bool, bool) {
	var chunking bool
	rv := objc.Send[bool](objc.ID(_ANEVirtualClientClass.class), objc.Sel("shouldUsePrecompiledPath:options:shouldUseChunking:chunkingThreshold:"), path, options, unsafe.Pointer(&chunking), threshold)
	return chunking, rv
}
func (_ANEVirtualClientClass ANEVirtualClientClass) UpdateErrorErrorLengthError(error_ iosurface.IOSurfaceRef, length uint64) (bool, error) {
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
func (_ANEVirtualClientClass ANEVirtualClientClass) UpdateErrorErrorLengthErrorCodeError(error_ iosurface.IOSurfaceRef, length uint64, code int64) (bool, error) {
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
func (_ANEVirtualClientClass ANEVirtualClientClass) UpdatePerformanceStatsPerformanceStatsLengthPerfStatsRawIOSurfaceRefPerformanceStatsRawLengthHwExecutionTime(stats iosurface.IOSurfaceRef, length uint64, ref iosurface.IOSurfaceRef, length2 uint64, time uint64) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEVirtualClientClass.class), objc.Sel("updatePerformanceStats:performanceStatsLength:perfStatsRawIOSurfaceRef:performanceStatsRawLength:hwExecutionTime:"), stats, length, ref, length2, time)
	return objectivec.Object{ID: rv}
}

func (a ANEVirtualClient) Connect() uint32 {
	rv := objc.SendIfResponds[uint32](a.ID, objc.Sel("connect"))
	return rv
}
func (a ANEVirtualClient) Queue() objectivec.Object {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("queue"))
	return objectivec.ObjectFromID(objc.ID(rv))
}
