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

// The class instance for the [ANEClient] class.
var (
	_ANEClientClass     ANEClientClass
	_ANEClientClassOnce sync.Once
)

func getANEClientClass() ANEClientClass {
	_ANEClientClassOnce.Do(func() {
		_ANEClientClass = ANEClientClass{class: objc.GetClass("_ANEClient")}
	})
	return _ANEClientClass
}

// GetANEClientClass returns the class object for _ANEClient.
func GetANEClientClass() ANEClientClass {
	return getANEClientClass()
}

type ANEClientClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEClientClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEClientClass) Alloc() ANEClient {
	rv := objc.SendIfResponds[ANEClient](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ANEClient.AllowRestrictedAccess]
//   - [ANEClient.BeginRealTimeTask]
//   - [ANEClient.BuffersReadyWithModelInputBuffersOptionsQosError]
//   - [ANEClient.CompileModelOptionsQosError]
//   - [ANEClient.CompiledModelExistsFor]
//   - [ANEClient.CompiledModelExistsMatchingHash]
//   - [ANEClient.Conn]
//   - [ANEClient.ConnectionForLoadingModelOptions]
//   - [ANEClient.ConnectionUsedForLoadingModel]
//   - [ANEClient.Connections]
//   - [ANEClient.ConnectionsUsedForLoadingModels]
//   - [ANEClient.DoBuffersReadyWithModelInputBuffersOptionsQosError]
//   - [ANEClient.DoEnqueueSetsWithModelOutputSetOptionsQosError]
//   - [ANEClient.DoEvaluateDirectWithModelOptionsRequestQosError]
//   - [ANEClient.DoLoadModelOptionsQosError]
//   - [ANEClient.DoLoadModelNewInstanceOptionsModelInstParamsQosError]
//   - [ANEClient.DoPrepareChainingWithModelOptionsChainingReqQosError]
//   - [ANEClient.DoUnloadModelOptionsQosError]
//   - [ANEClient.Echo]
//   - [ANEClient.EndRealTimeTask]
//   - [ANEClient.EnqueueSetsWithModelOutputSetOptionsQosError]
//   - [ANEClient.EvaluateRealTimeWithModelOptionsRequestError]
//   - [ANEClient.EvaluateWithModelOptionsRequestQosError]
//   - [ANEClient.FastConn]
//   - [ANEClient.FastConnWithoutLock]
//   - [ANEClient.IsAnetoolRootDaemonConnection]
//   - [ANEClient.IsRootDaemon]
//   - [ANEClient.IsVirtualClient]
//   - [ANEClient.LoadModelOptionsQosError]
//   - [ANEClient.LoadModelNewInstanceOptionsModelInstParamsQosError]
//   - [ANEClient.LoadRealTimeModelOptionsQosError]
//   - [ANEClient.MapIOSurfacesWithModelRequestCacheInferenceError]
//   - [ANEClient.PrepareChainingWithModelOptionsChainingReqQosError]
//   - [ANEClient.PriorityQ]
//   - [ANEClient.PurgeCompiledModel]
//   - [ANEClient.PurgeCompiledModelMatchingHash]
//   - [ANEClient.ReportEvaluateFailureFailureReasonQIdx]
//   - [ANEClient.SessionHintWithModelHintOptionsReportError]
//   - [ANEClient.UnloadModelOptionsQosError]
//   - [ANEClient.UnloadRealTimeModelOptionsQosError]
//   - [ANEClient.UnmapIOSurfacesWithModelRequest]
//   - [ANEClient.VirtualClient]
//   - [ANEClient.InitWithRestrictedAccessAllowed]
type ANEClient struct {
	objectivec.Object
}

// ANEClientFromID constructs a [ANEClient] from an objc.ID.
func ANEClientFromID(id objc.ID) ANEClient {
	return ANEClient{objectivec.Object{ID: id}}
}

// Ensure ANEClient implements IANEClient.
var _ IANEClient = ANEClient{}

// An interface definition for the [ANEClient] class.
//
// # Methods
//
//   - [IANEClient.AllowRestrictedAccess]
//   - [IANEClient.BeginRealTimeTask]
//   - [IANEClient.BuffersReadyWithModelInputBuffersOptionsQosError]
//   - [IANEClient.CompileModelOptionsQosError]
//   - [IANEClient.CompiledModelExistsFor]
//   - [IANEClient.CompiledModelExistsMatchingHash]
//   - [IANEClient.Conn]
//   - [IANEClient.ConnectionForLoadingModelOptions]
//   - [IANEClient.ConnectionUsedForLoadingModel]
//   - [IANEClient.Connections]
//   - [IANEClient.ConnectionsUsedForLoadingModels]
//   - [IANEClient.DoBuffersReadyWithModelInputBuffersOptionsQosError]
//   - [IANEClient.DoEnqueueSetsWithModelOutputSetOptionsQosError]
//   - [IANEClient.DoEvaluateDirectWithModelOptionsRequestQosError]
//   - [IANEClient.DoLoadModelOptionsQosError]
//   - [IANEClient.DoLoadModelNewInstanceOptionsModelInstParamsQosError]
//   - [IANEClient.DoPrepareChainingWithModelOptionsChainingReqQosError]
//   - [IANEClient.DoUnloadModelOptionsQosError]
//   - [IANEClient.Echo]
//   - [IANEClient.EndRealTimeTask]
//   - [IANEClient.EnqueueSetsWithModelOutputSetOptionsQosError]
//   - [IANEClient.EvaluateRealTimeWithModelOptionsRequestError]
//   - [IANEClient.EvaluateWithModelOptionsRequestQosError]
//   - [IANEClient.FastConn]
//   - [IANEClient.FastConnWithoutLock]
//   - [IANEClient.IsAnetoolRootDaemonConnection]
//   - [IANEClient.IsRootDaemon]
//   - [IANEClient.IsVirtualClient]
//   - [IANEClient.LoadModelOptionsQosError]
//   - [IANEClient.LoadModelNewInstanceOptionsModelInstParamsQosError]
//   - [IANEClient.LoadRealTimeModelOptionsQosError]
//   - [IANEClient.MapIOSurfacesWithModelRequestCacheInferenceError]
//   - [IANEClient.PrepareChainingWithModelOptionsChainingReqQosError]
//   - [IANEClient.PriorityQ]
//   - [IANEClient.PurgeCompiledModel]
//   - [IANEClient.PurgeCompiledModelMatchingHash]
//   - [IANEClient.ReportEvaluateFailureFailureReasonQIdx]
//   - [IANEClient.SessionHintWithModelHintOptionsReportError]
//   - [IANEClient.UnloadModelOptionsQosError]
//   - [IANEClient.UnloadRealTimeModelOptionsQosError]
//   - [IANEClient.UnmapIOSurfacesWithModelRequest]
//   - [IANEClient.VirtualClient]
//   - [IANEClient.InitWithRestrictedAccessAllowed]
type IANEClient interface {
	objectivec.IObject

	// Topic: Methods

	AllowRestrictedAccess() bool
	BeginRealTimeTask() bool
	BuffersReadyWithModelInputBuffersOptionsQosError(model objectivec.IObject, buffers objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error)
	CompileModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error)
	CompiledModelExistsFor(for_ objectivec.IObject) bool
	CompiledModelExistsMatchingHash(hash objectivec.IObject) bool
	Conn() IANEDaemonConnection
	ConnectionForLoadingModelOptions(model objectivec.IObject, options objectivec.IObject) objectivec.IObject
	ConnectionUsedForLoadingModel(model objectivec.IObject) objectivec.IObject
	Connections() foundation.INSDictionary
	ConnectionsUsedForLoadingModels() foundation.INSDictionary
	DoBuffersReadyWithModelInputBuffersOptionsQosError(model objectivec.IObject, buffers objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error)
	DoEnqueueSetsWithModelOutputSetOptionsQosError(model objectivec.IObject, set IANEOutputSetEnqueue, options objectivec.IObject, qos uint32) (bool, error)
	DoEvaluateDirectWithModelOptionsRequestQosError(model objectivec.IObject, options objectivec.IObject, request objectivec.IObject, qos uint32) (bool, error)
	DoLoadModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error)
	DoLoadModelNewInstanceOptionsModelInstParamsQosError(instance objectivec.IObject, options objectivec.IObject, params objectivec.IObject, qos uint32) (bool, error)
	DoPrepareChainingWithModelOptionsChainingReqQosError(model objectivec.IObject, options objectivec.IObject, req objectivec.IObject, qos uint32) (bool, error)
	DoUnloadModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error)
	Echo(echo objectivec.IObject) bool
	EndRealTimeTask() bool
	EnqueueSetsWithModelOutputSetOptionsQosError(model objectivec.IObject, set IANEOutputSetEnqueue, options objectivec.IObject, qos uint32) (bool, error)
	EvaluateRealTimeWithModelOptionsRequestError(model objectivec.IObject, options objectivec.IObject, request objectivec.IObject) (bool, error)
	EvaluateWithModelOptionsRequestQosError(model objectivec.IObject, options objectivec.IObject, request objectivec.IObject, qos uint32) (bool, error)
	FastConn() IANEDaemonConnection
	FastConnWithoutLock() objectivec.IObject
	IsAnetoolRootDaemonConnection() bool
	IsRootDaemon() bool
	IsVirtualClient() bool
	LoadModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error)
	LoadModelNewInstanceOptionsModelInstParamsQosError(instance objectivec.IObject, options objectivec.IObject, params objectivec.IObject, qos uint32) (bool, error)
	LoadRealTimeModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error)
	MapIOSurfacesWithModelRequestCacheInferenceError(model objectivec.IObject, request objectivec.IObject, inference bool) (bool, error)
	PrepareChainingWithModelOptionsChainingReqQosError(model objectivec.IObject, options objectivec.IObject, req objectivec.IObject, qos uint32) (bool, error)
	PriorityQ() foundation.INSArray
	PurgeCompiledModel(model objectivec.IObject)
	PurgeCompiledModelMatchingHash(hash objectivec.IObject)
	ReportEvaluateFailureFailureReasonQIdx(failure objectivec.IObject, reason uint32, idx uint64)
	SessionHintWithModelHintOptionsReportError(model objectivec.IObject, hint objectivec.IObject, options objectivec.IObject, report objectivec.IObject) (bool, error)
	UnloadModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error)
	UnloadRealTimeModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error)
	UnmapIOSurfacesWithModelRequest(model objectivec.IObject, request objectivec.IObject)
	VirtualClient() IANEVirtualClient
	InitWithRestrictedAccessAllowed(allowed bool) ANEClient
}

// Init initializes the instance.
func (a ANEClient) Init() ANEClient {
	rv := objc.SendIfResponds[ANEClient](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEClient) Autorelease() ANEClient {
	rv := objc.SendIfResponds[ANEClient](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEClient creates a new ANEClient instance.
func NewANEClient() ANEClient {
	class := getANEClientClass()
	rv := objc.SendIfResponds[ANEClient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewANEClientWithRestrictedAccessAllowed(allowed bool) ANEClient {
	instance := getANEClientClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithRestrictedAccessAllowed:"), allowed)
	return ANEClientFromID(rv)
}

func (a ANEClient) BeginRealTimeTask() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("beginRealTimeTask"))
	return rv
}
func (a ANEClient) BuffersReadyWithModelInputBuffersOptionsQosError(model objectivec.IObject, buffers objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("buffersReadyWithModel:inputBuffers:options:qos:error:"), model, buffers, options, qos, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("buffersReadyWithModel:inputBuffers:options:qos:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (a ANEClient) CompileModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error) {
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
func (a ANEClient) CompiledModelExistsFor(for_ objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("compiledModelExistsFor:"), for_)
	return rv
}
func (a ANEClient) CompiledModelExistsMatchingHash(hash objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("compiledModelExistsMatchingHash:"), hash)
	return rv
}
func (a ANEClient) ConnectionForLoadingModelOptions(model objectivec.IObject, options objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("connectionForLoadingModel:options:"), model, options)
	return objectivec.Object{ID: rv}
}
func (a ANEClient) ConnectionUsedForLoadingModel(model objectivec.IObject) objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("connectionUsedForLoadingModel:"), model)
	return objectivec.Object{ID: rv}
}
func (a ANEClient) DoBuffersReadyWithModelInputBuffersOptionsQosError(model objectivec.IObject, buffers objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("doBuffersReadyWithModel:inputBuffers:options:qos:error:"), model, buffers, options, qos, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("doBuffersReadyWithModel:inputBuffers:options:qos:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (a ANEClient) DoEnqueueSetsWithModelOutputSetOptionsQosError(model objectivec.IObject, set IANEOutputSetEnqueue, options objectivec.IObject, qos uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("doEnqueueSetsWithModel:outputSet:options:qos:error:"), model, set, options, qos, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("doEnqueueSetsWithModel:outputSet:options:qos:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (a ANEClient) DoEvaluateDirectWithModelOptionsRequestQosError(model objectivec.IObject, options objectivec.IObject, request objectivec.IObject, qos uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("doEvaluateDirectWithModel:options:request:qos:error:"), model, options, request, qos, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("doEvaluateDirectWithModel:options:request:qos:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (a ANEClient) DoLoadModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("doLoadModel:options:qos:error:"), model, options, qos, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("doLoadModel:options:qos:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (a ANEClient) DoLoadModelNewInstanceOptionsModelInstParamsQosError(instance objectivec.IObject, options objectivec.IObject, params objectivec.IObject, qos uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("doLoadModelNewInstance:options:modelInstParams:qos:error:"), instance, options, params, qos, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("doLoadModelNewInstance:options:modelInstParams:qos:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (a ANEClient) DoPrepareChainingWithModelOptionsChainingReqQosError(model objectivec.IObject, options objectivec.IObject, req objectivec.IObject, qos uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("doPrepareChainingWithModel:options:chainingReq:qos:error:"), model, options, req, qos, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("doPrepareChainingWithModel:options:chainingReq:qos:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (a ANEClient) DoUnloadModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("doUnloadModel:options:qos:error:"), model, options, qos, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("doUnloadModel:options:qos:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (a ANEClient) Echo(echo objectivec.IObject) bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("echo:"), echo)
	return rv
}
func (a ANEClient) EndRealTimeTask() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("endRealTimeTask"))
	return rv
}
func (a ANEClient) EnqueueSetsWithModelOutputSetOptionsQosError(model objectivec.IObject, set IANEOutputSetEnqueue, options objectivec.IObject, qos uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("enqueueSetsWithModel:outputSet:options:qos:error:"), model, set, options, qos, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("enqueueSetsWithModel:outputSet:options:qos:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (a ANEClient) EvaluateRealTimeWithModelOptionsRequestError(model objectivec.IObject, options objectivec.IObject, request objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("evaluateRealTimeWithModel:options:request:error:"), model, options, request, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("evaluateRealTimeWithModel:options:request:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (a ANEClient) EvaluateWithModelOptionsRequestQosError(model objectivec.IObject, options objectivec.IObject, request objectivec.IObject, qos uint32) (bool, error) {
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
func (a ANEClient) FastConnWithoutLock() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("fastConnWithoutLock"))
	return objectivec.Object{ID: rv}
}
func (a ANEClient) IsAnetoolRootDaemonConnection() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("isAnetoolRootDaemonConnection"))
	return rv
}
func (a ANEClient) IsVirtualClient() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("isVirtualClient"))
	return rv
}
func (a ANEClient) LoadModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error) {
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
func (a ANEClient) LoadModelNewInstanceOptionsModelInstParamsQosError(instance objectivec.IObject, options objectivec.IObject, params objectivec.IObject, qos uint32) (bool, error) {
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
func (a ANEClient) LoadRealTimeModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("loadRealTimeModel:options:qos:error:"), model, options, qos, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("loadRealTimeModel:options:qos:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (a ANEClient) MapIOSurfacesWithModelRequestCacheInferenceError(model objectivec.IObject, request objectivec.IObject, inference bool) (bool, error) {
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
func (a ANEClient) PrepareChainingWithModelOptionsChainingReqQosError(model objectivec.IObject, options objectivec.IObject, req objectivec.IObject, qos uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("prepareChainingWithModel:options:chainingReq:qos:error:"), model, options, req, qos, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("prepareChainingWithModel:options:chainingReq:qos:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (a ANEClient) PurgeCompiledModel(model objectivec.IObject) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("purgeCompiledModel:"), model)
}
func (a ANEClient) PurgeCompiledModelMatchingHash(hash objectivec.IObject) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("purgeCompiledModelMatchingHash:"), hash)
}
func (a ANEClient) ReportEvaluateFailureFailureReasonQIdx(failure objectivec.IObject, reason uint32, idx uint64) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("reportEvaluateFailure:failureReason:qIdx:"), failure, reason, idx)
}
func (a ANEClient) SessionHintWithModelHintOptionsReportError(model objectivec.IObject, hint objectivec.IObject, options objectivec.IObject, report objectivec.IObject) (bool, error) {
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
func (a ANEClient) UnloadModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error) {
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
func (a ANEClient) UnloadRealTimeModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("unloadRealTimeModel:options:qos:error:"), model, options, qos, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("unloadRealTimeModel:options:qos:error: returned NO with nil NSError")
	}
	return rv, nil

}
func (a ANEClient) UnmapIOSurfacesWithModelRequest(model objectivec.IObject, request objectivec.IObject) {
	objc.SendIfResponds[objc.ID](a.ID, objc.Sel("unmapIOSurfacesWithModel:request:"), model, request)
}
func (a ANEClient) InitWithRestrictedAccessAllowed(allowed bool) ANEClient {
	rv := objc.SendIfResponds[ANEClient](a.ID, objc.Sel("initWithRestrictedAccessAllowed:"), allowed)
	return rv
}

func (_ANEClientClass ANEClientClass) SharedConnection() ANEClient {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEClientClass.class), objc.Sel("sharedConnection"))
	return ANEClientFromID(rv)
}
func (_ANEClientClass ANEClientClass) SharedPrivateConnection() ANEClient {
	rv := objc.SendIfResponds[objc.ID](objc.ID(_ANEClientClass.class), objc.Sel("sharedPrivateConnection"))
	return ANEClientFromID(rv)
}

func (a ANEClient) AllowRestrictedAccess() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("allowRestrictedAccess"))
	return rv
}
func (a ANEClient) Conn() IANEDaemonConnection {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("conn"))
	return ANEDaemonConnectionFromID(objc.ID(rv))
}
func (a ANEClient) Connections() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("connections"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (a ANEClient) ConnectionsUsedForLoadingModels() foundation.INSDictionary {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("connectionsUsedForLoadingModels"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
func (a ANEClient) FastConn() IANEDaemonConnection {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("fastConn"))
	return ANEDaemonConnectionFromID(objc.ID(rv))
}
func (a ANEClient) IsRootDaemon() bool {
	rv := objc.SendIfResponds[bool](a.ID, objc.Sel("isRootDaemon"))
	return rv
}
func (a ANEClient) PriorityQ() foundation.INSArray {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("priorityQ"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (a ANEClient) VirtualClient() IANEVirtualClient {
	rv := objc.SendIfResponds[objc.ID](a.ID, objc.Sel("virtualClient"))
	return ANEVirtualClientFromID(objc.ID(rv))
}
