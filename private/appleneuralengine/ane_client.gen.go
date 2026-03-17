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

// Alloc allocates memory for a new instance of the class.
func (ac ANEClientClass) Alloc() ANEClient {
	rv := objc.Send[ANEClient](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
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
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient
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
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient
type IANEClient interface {
	objectivec.IObject

	// Topic: Methods

	AllowRestrictedAccess() bool
	BeginRealTimeTask() bool
	BuffersReadyWithModelInputBuffersOptionsQosError(model objectivec.IObject, buffers objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error)
	CompileModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error)
	CompiledModelExistsFor(for_ objectivec.IObject) bool
	CompiledModelExistsMatchingHash(hash objectivec.IObject) bool
	Conn() *ANEDaemonConnection
	ConnectionForLoadingModelOptions(model objectivec.IObject, options objectivec.IObject) objectivec.IObject
	ConnectionUsedForLoadingModel(model objectivec.IObject) objectivec.IObject
	Connections() foundation.INSDictionary
	ConnectionsUsedForLoadingModels() foundation.INSDictionary
	DoBuffersReadyWithModelInputBuffersOptionsQosError(model objectivec.IObject, buffers objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error)
	DoEnqueueSetsWithModelOutputSetOptionsQosError(model objectivec.IObject, set *ANEOutputSetEnqueue, options objectivec.IObject, qos uint32) (bool, error)
	DoEvaluateDirectWithModelOptionsRequestQosError(model objectivec.IObject, options objectivec.IObject, request objectivec.IObject, qos uint32) (bool, error)
	DoLoadModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error)
	DoLoadModelNewInstanceOptionsModelInstParamsQosError(instance objectivec.IObject, options objectivec.IObject, params objectivec.IObject, qos uint32) (bool, error)
	DoPrepareChainingWithModelOptionsChainingReqQosError(model objectivec.IObject, options objectivec.IObject, req objectivec.IObject, qos uint32) (bool, error)
	DoUnloadModelOptionsQosError(model objectivec.IObject, options objectivec.IObject, qos uint32) (bool, error)
	Echo(echo objectivec.IObject) bool
	EndRealTimeTask() bool
	EnqueueSetsWithModelOutputSetOptionsQosError(model objectivec.IObject, set *ANEOutputSetEnqueue, options objectivec.IObject, qos uint32) (bool, error)
	EvaluateRealTimeWithModelOptionsRequestError(model objectivec.IObject, options objectivec.IObject, request objectivec.IObject) (bool, error)
	EvaluateWithModelOptionsRequestQosError(model objectivec.IObject, options objectivec.IObject, request objectivec.IObject, qos uint32) (bool, error)
	FastConn() *ANEDaemonConnection
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
	VirtualClient() *ANEVirtualClient
	InitWithRestrictedAccessAllowed(allowed bool) ANEClient
}

// Init initializes the instance.
func (a ANEClient) Init() ANEClient {
	rv := objc.Send[ANEClient](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEClient) Autorelease() ANEClient {
	rv := objc.Send[ANEClient](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEClient creates a new ANEClient instance.
func NewANEClient() ANEClient {
	class := getANEClientClass()
	rv := objc.Send[ANEClient](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/initWithRestrictedAccessAllowed:
func NewANEClientWithRestrictedAccessAllowed(allowed bool) ANEClient {
	instance := getANEClientClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRestrictedAccessAllowed:"), allowed)
	return ANEClientFromID(rv)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/beginRealTimeTask
func (a ANEClient) BeginRealTimeTask() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("beginRealTimeTask"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/buffersReadyWithModel:inputBuffers:options:qos:error:
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

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/compileModel:options:qos:error:
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

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/compiledModelExistsFor:
func (a ANEClient) CompiledModelExistsFor(for_ objectivec.IObject) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("compiledModelExistsFor:"), for_)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/compiledModelExistsMatchingHash:
func (a ANEClient) CompiledModelExistsMatchingHash(hash objectivec.IObject) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("compiledModelExistsMatchingHash:"), hash)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/connectionForLoadingModel:options:
func (a ANEClient) ConnectionForLoadingModelOptions(model objectivec.IObject, options objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("connectionForLoadingModel:options:"), model, options)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/connectionUsedForLoadingModel:
func (a ANEClient) ConnectionUsedForLoadingModel(model objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("connectionUsedForLoadingModel:"), model)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/doBuffersReadyWithModel:inputBuffers:options:qos:error:
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

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/doEnqueueSetsWithModel:outputSet:options:qos:error:
func (a ANEClient) DoEnqueueSetsWithModelOutputSetOptionsQosError(model objectivec.IObject, set *ANEOutputSetEnqueue, options objectivec.IObject, qos uint32) (bool, error) {
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

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/doEvaluateDirectWithModel:options:request:qos:error:
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

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/doLoadModel:options:qos:error:
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

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/doLoadModelNewInstance:options:modelInstParams:qos:error:
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

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/doPrepareChainingWithModel:options:chainingReq:qos:error:
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

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/doUnloadModel:options:qos:error:
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

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/echo:
func (a ANEClient) Echo(echo objectivec.IObject) bool {
	rv := objc.Send[bool](a.ID, objc.Sel("echo:"), echo)
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/endRealTimeTask
func (a ANEClient) EndRealTimeTask() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("endRealTimeTask"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/enqueueSetsWithModel:outputSet:options:qos:error:
func (a ANEClient) EnqueueSetsWithModelOutputSetOptionsQosError(model objectivec.IObject, set *ANEOutputSetEnqueue, options objectivec.IObject, qos uint32) (bool, error) {
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

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/evaluateRealTimeWithModel:options:request:error:
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

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/evaluateWithModel:options:request:qos:error:
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

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/fastConnWithoutLock
func (a ANEClient) FastConnWithoutLock() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("fastConnWithoutLock"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/isAnetoolRootDaemonConnection
func (a ANEClient) IsAnetoolRootDaemonConnection() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isAnetoolRootDaemonConnection"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/isVirtualClient
func (a ANEClient) IsVirtualClient() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isVirtualClient"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/loadModel:options:qos:error:
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

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/loadModelNewInstance:options:modelInstParams:qos:error:
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

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/loadRealTimeModel:options:qos:error:
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

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/mapIOSurfacesWithModel:request:cacheInference:error:
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

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/prepareChainingWithModel:options:chainingReq:qos:error:
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

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/purgeCompiledModel:
func (a ANEClient) PurgeCompiledModel(model objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("purgeCompiledModel:"), model)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/purgeCompiledModelMatchingHash:
func (a ANEClient) PurgeCompiledModelMatchingHash(hash objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("purgeCompiledModelMatchingHash:"), hash)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/reportEvaluateFailure:failureReason:qIdx:
func (a ANEClient) ReportEvaluateFailureFailureReasonQIdx(failure objectivec.IObject, reason uint32, idx uint64) {
	objc.Send[objc.ID](a.ID, objc.Sel("reportEvaluateFailure:failureReason:qIdx:"), failure, reason, idx)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/sessionHintWithModel:hint:options:report:error:
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

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/unloadModel:options:qos:error:
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

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/unloadRealTimeModel:options:qos:error:
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

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/unmapIOSurfacesWithModel:request:
func (a ANEClient) UnmapIOSurfacesWithModelRequest(model objectivec.IObject, request objectivec.IObject) {
	objc.Send[objc.ID](a.ID, objc.Sel("unmapIOSurfacesWithModel:request:"), model, request)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/initWithRestrictedAccessAllowed:
func (a ANEClient) InitWithRestrictedAccessAllowed(allowed bool) ANEClient {
	rv := objc.Send[ANEClient](a.ID, objc.Sel("initWithRestrictedAccessAllowed:"), allowed)
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/sharedConnection
func (_ANEClientClass ANEClientClass) SharedConnection() *ANEClient {
	rv := objc.Send[objc.ID](objc.ID(_ANEClientClass.class), objc.Sel("sharedConnection"))
	if rv == 0 {
		return nil
	}
	val := ANEClientFromID(rv)
	return &val
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/sharedPrivateConnection
func (_ANEClientClass ANEClientClass) SharedPrivateConnection() *ANEClient {
	rv := objc.Send[objc.ID](objc.ID(_ANEClientClass.class), objc.Sel("sharedPrivateConnection"))
	if rv == 0 {
		return nil
	}
	val := ANEClientFromID(rv)
	return &val
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/allowRestrictedAccess
func (a ANEClient) AllowRestrictedAccess() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("allowRestrictedAccess"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/conn
func (a ANEClient) Conn() *ANEDaemonConnection {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("conn"))
	if rv == 0 {
		return nil
	}
	val := ANEDaemonConnectionFromID(objc.ID(rv))
	return &val
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/connections
func (a ANEClient) Connections() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("connections"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/connectionsUsedForLoadingModels
func (a ANEClient) ConnectionsUsedForLoadingModels() foundation.INSDictionary {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("connectionsUsedForLoadingModels"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/fastConn
func (a ANEClient) FastConn() *ANEDaemonConnection {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("fastConn"))
	if rv == 0 {
		return nil
	}
	val := ANEDaemonConnectionFromID(objc.ID(rv))
	return &val
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/isRootDaemon
func (a ANEClient) IsRootDaemon() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isRootDaemon"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/priorityQ
func (a ANEClient) PriorityQ() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("priorityQ"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEClient/virtualClient
func (a ANEClient) VirtualClient() *ANEVirtualClient {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("virtualClient"))
	if rv == 0 {
		return nil
	}
	val := ANEVirtualClientFromID(objc.ID(rv))
	return &val
}

