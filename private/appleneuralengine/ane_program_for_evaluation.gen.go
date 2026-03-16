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

// The class instance for the [ANEProgramForEvaluation] class.
var (
	_ANEProgramForEvaluationClass     ANEProgramForEvaluationClass
	_ANEProgramForEvaluationClassOnce sync.Once
)

func getANEProgramForEvaluationClass() ANEProgramForEvaluationClass {
	_ANEProgramForEvaluationClassOnce.Do(func() {
		_ANEProgramForEvaluationClass = ANEProgramForEvaluationClass{class: objc.GetClass("_ANEProgramForEvaluation")}
	})
	return _ANEProgramForEvaluationClass
}

// GetANEProgramForEvaluationClass returns the class object for _ANEProgramForEvaluation.
func GetANEProgramForEvaluationClass() ANEProgramForEvaluationClass {
	return getANEProgramForEvaluationClass()
}

type ANEProgramForEvaluationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEProgramForEvaluationClass) Alloc() ANEProgramForEvaluation {
	rv := objc.Send[ANEProgramForEvaluation](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ANEProgramForEvaluation.Controller]
//   - [ANEProgramForEvaluation.CurrentAsyncRequestsInFlight]
//   - [ANEProgramForEvaluation.SetCurrentAsyncRequestsInFlight]
//   - [ANEProgramForEvaluation.IntermediateBufferHandle]
//   - [ANEProgramForEvaluation.SetIntermediateBufferHandle]
//   - [ANEProgramForEvaluation.ProcessInputBuffersModelOptionsError]
//   - [ANEProgramForEvaluation.ProcessOutputSetModelOptionsError]
//   - [ANEProgramForEvaluation.ProcessRequestModelQosQIndexModelStringIDOptionsReturnValueError]
//   - [ANEProgramForEvaluation.ProcessSessionHintOptionsReportError]
//   - [ANEProgramForEvaluation.ProgramHandle]
//   - [ANEProgramForEvaluation.SetProgramHandle]
//   - [ANEProgramForEvaluation.ProgramInferenceOtherErrorForMessageModelMethodName]
//   - [ANEProgramForEvaluation.QueueDepth]
//   - [ANEProgramForEvaluation.RequestsInFlight]
//   - [ANEProgramForEvaluation.InitWithControllerIntermediateBufferHandleQueueDepth]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramForEvaluation
type ANEProgramForEvaluation struct {
	objectivec.Object
}

// ANEProgramForEvaluationFromID constructs a [ANEProgramForEvaluation] from an objc.ID.
func ANEProgramForEvaluationFromID(id objc.ID) ANEProgramForEvaluation {
	return ANEProgramForEvaluation{objectivec.Object{id}}
}
// Ensure ANEProgramForEvaluation implements IANEProgramForEvaluation.
var _ IANEProgramForEvaluation = ANEProgramForEvaluation{}

// An interface definition for the [ANEProgramForEvaluation] class.
//
// # Methods
//
//   - [IANEProgramForEvaluation.Controller]
//   - [IANEProgramForEvaluation.CurrentAsyncRequestsInFlight]
//   - [IANEProgramForEvaluation.SetCurrentAsyncRequestsInFlight]
//   - [IANEProgramForEvaluation.IntermediateBufferHandle]
//   - [IANEProgramForEvaluation.SetIntermediateBufferHandle]
//   - [IANEProgramForEvaluation.ProcessInputBuffersModelOptionsError]
//   - [IANEProgramForEvaluation.ProcessOutputSetModelOptionsError]
//   - [IANEProgramForEvaluation.ProcessRequestModelQosQIndexModelStringIDOptionsReturnValueError]
//   - [IANEProgramForEvaluation.ProcessSessionHintOptionsReportError]
//   - [IANEProgramForEvaluation.ProgramHandle]
//   - [IANEProgramForEvaluation.SetProgramHandle]
//   - [IANEProgramForEvaluation.ProgramInferenceOtherErrorForMessageModelMethodName]
//   - [IANEProgramForEvaluation.QueueDepth]
//   - [IANEProgramForEvaluation.RequestsInFlight]
//   - [IANEProgramForEvaluation.InitWithControllerIntermediateBufferHandleQueueDepth]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramForEvaluation
type IANEProgramForEvaluation interface {
	objectivec.IObject

	// Topic: Methods

	Controller() *ANEDeviceController
	CurrentAsyncRequestsInFlight() int64
	SetCurrentAsyncRequestsInFlight(value int64)
	IntermediateBufferHandle() uint64
	SetIntermediateBufferHandle(value uint64)
	ProcessInputBuffersModelOptionsError(buffers objectivec.IObject, model objectivec.IObject, options objectivec.IObject) (bool, error)
	ProcessOutputSetModelOptionsError(set objectivec.IObject, model objectivec.IObject, options objectivec.IObject) (bool, error)
	ProcessRequestModelQosQIndexModelStringIDOptionsReturnValueError(request objectivec.IObject, model objectivec.IObject, qos uint32, index uint64, id uint64, options objectivec.IObject, value unsafe.Pointer) (bool, error)
	ProcessSessionHintOptionsReportError(hint objectivec.IObject, options objectivec.IObject, report objectivec.IObject) (bool, error)
	ProgramHandle() uint64
	SetProgramHandle(value uint64)
	ProgramInferenceOtherErrorForMessageModelMethodName(message unsafe.Pointer, model objectivec.IObject, name objectivec.IObject) objectivec.IObject
	QueueDepth() int8
	RequestsInFlight() objectivec.Object
	InitWithControllerIntermediateBufferHandleQueueDepth(controller objectivec.IObject, handle uint64, depth int8) ANEProgramForEvaluation
}

// Init initializes the instance.
func (a ANEProgramForEvaluation) Init() ANEProgramForEvaluation {
	rv := objc.Send[ANEProgramForEvaluation](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEProgramForEvaluation) Autorelease() ANEProgramForEvaluation {
	rv := objc.Send[ANEProgramForEvaluation](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEProgramForEvaluation creates a new ANEProgramForEvaluation instance.
func NewANEProgramForEvaluation() ANEProgramForEvaluation {
	class := getANEProgramForEvaluationClass()
	rv := objc.Send[ANEProgramForEvaluation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramForEvaluation/initWithController:intermediateBufferHandle:queueDepth:
func NewANEProgramForEvaluationWithControllerIntermediateBufferHandleQueueDepth(controller objectivec.IObject, handle uint64, depth int8) ANEProgramForEvaluation {
	instance := getANEProgramForEvaluationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithController:intermediateBufferHandle:queueDepth:"), controller, handle, depth)
	return ANEProgramForEvaluationFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramForEvaluation/processInputBuffers:model:options:error:
func (a ANEProgramForEvaluation) ProcessInputBuffersModelOptionsError(buffers objectivec.IObject, model objectivec.IObject, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("processInputBuffers:model:options:error:"), buffers, model, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("processInputBuffers:model:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramForEvaluation/processOutputSet:model:options:error:
func (a ANEProgramForEvaluation) ProcessOutputSetModelOptionsError(set objectivec.IObject, model objectivec.IObject, options objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("processOutputSet:model:options:error:"), set, model, options, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("processOutputSet:model:options:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramForEvaluation/processRequest:model:qos:qIndex:modelStringID:options:returnValue:error:
func (a ANEProgramForEvaluation) ProcessRequestModelQosQIndexModelStringIDOptionsReturnValueError(request objectivec.IObject, model objectivec.IObject, qos uint32, index uint64, id uint64, options objectivec.IObject, value unsafe.Pointer) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("processRequest:model:qos:qIndex:modelStringID:options:returnValue:error:"), request, model, qos, index, id, options, value, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("processRequest:model:qos:qIndex:modelStringID:options:returnValue:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramForEvaluation/processSessionHint:options:report:error:
func (a ANEProgramForEvaluation) ProcessSessionHintOptionsReportError(hint objectivec.IObject, options objectivec.IObject, report objectivec.IObject) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](a.ID, objc.Sel("processSessionHint:options:report:error:"), hint, options, report, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("processSessionHint:options:report:error: returned NO with nil NSError")
	}
	return rv, nil

}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramForEvaluation/programInferenceOtherErrorForMessage:model:methodName:
func (a ANEProgramForEvaluation) ProgramInferenceOtherErrorForMessageModelMethodName(message unsafe.Pointer, model objectivec.IObject, name objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("programInferenceOtherErrorForMessage:model:methodName:"), message, model, name)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramForEvaluation/initWithController:intermediateBufferHandle:queueDepth:
func (a ANEProgramForEvaluation) InitWithControllerIntermediateBufferHandleQueueDepth(controller objectivec.IObject, handle uint64, depth int8) ANEProgramForEvaluation {
	rv := objc.Send[ANEProgramForEvaluation](a.ID, objc.Sel("initWithController:intermediateBufferHandle:queueDepth:"), controller, handle, depth)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramForEvaluation/programWithController:intermediateBufferHandle:queueDepth:
func (_ANEProgramForEvaluationClass ANEProgramForEvaluationClass) ProgramWithControllerIntermediateBufferHandleQueueDepth(controller objectivec.IObject, handle uint64, depth int8) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEProgramForEvaluationClass.class), objc.Sel("programWithController:intermediateBufferHandle:queueDepth:"), controller, handle, depth)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramForEvaluation/programWithHandle:intermediateBufferHandle:queueDepth:
func (_ANEProgramForEvaluationClass ANEProgramForEvaluationClass) ProgramWithHandleIntermediateBufferHandleQueueDepth(handle uint64, handle2 uint64, depth int8) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEProgramForEvaluationClass.class), objc.Sel("programWithHandle:intermediateBufferHandle:queueDepth:"), handle, handle2, depth)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramForEvaluation/controller
func (a ANEProgramForEvaluation) Controller() *ANEDeviceController {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("controller"))
	if rv == 0 {
		return nil
	}
	val := ANEDeviceControllerFromID(objc.ID(rv))
	return &val
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramForEvaluation/currentAsyncRequestsInFlight
func (a ANEProgramForEvaluation) CurrentAsyncRequestsInFlight() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("currentAsyncRequestsInFlight"))
	return rv
}
func (a ANEProgramForEvaluation) SetCurrentAsyncRequestsInFlight(value int64) {
	objc.Send[struct{}](a.ID, objc.Sel("setCurrentAsyncRequestsInFlight:"), value)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramForEvaluation/intermediateBufferHandle
func (a ANEProgramForEvaluation) IntermediateBufferHandle() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("intermediateBufferHandle"))
	return rv
}
func (a ANEProgramForEvaluation) SetIntermediateBufferHandle(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setIntermediateBufferHandle:"), value)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramForEvaluation/programHandle
func (a ANEProgramForEvaluation) ProgramHandle() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("programHandle"))
	return rv
}
func (a ANEProgramForEvaluation) SetProgramHandle(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setProgramHandle:"), value)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramForEvaluation/queueDepth
func (a ANEProgramForEvaluation) QueueDepth() int8 {
	rv := objc.Send[int8](a.ID, objc.Sel("queueDepth"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEProgramForEvaluation/requestsInFlight
func (a ANEProgramForEvaluation) RequestsInFlight() objectivec.Object {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("requestsInFlight"))
	return objectivec.ObjectFromID(objc.ID(rv))
}

