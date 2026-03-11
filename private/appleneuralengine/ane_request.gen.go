// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANERequest] class.
var (
	_ANERequestClass     ANERequestClass
	_ANERequestClassOnce sync.Once
)

func getANERequestClass() ANERequestClass {
	_ANERequestClassOnce.Do(func() {
		_ANERequestClass = ANERequestClass{class: objc.GetClass("_ANERequest")}
	})
	return _ANERequestClass
}

// GetANERequestClass returns the class object for _ANERequest.
func GetANERequestClass() ANERequestClass {
	return getANERequestClass()
}

type ANERequestClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANERequestClass) Alloc() ANERequest {
	rv := objc.Send[ANERequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}







//
// # Methods
//
//   - [ANERequest.InputArray]
//   - [ANERequest.InputIndexArray]
//   - [ANERequest.IoSurfacesCount]
//   - [ANERequest.OutputArray]
//   - [ANERequest.OutputIndexArray]
//   - [ANERequest.PerfStats]
//   - [ANERequest.SetPerfStats]
//   - [ANERequest.PerfStatsArray]
//   - [ANERequest.ProcedureIndex]
//   - [ANERequest.SetCompletionHandler]
//   - [ANERequest.SharedEvents]
//   - [ANERequest.SetSharedEvents]
//   - [ANERequest.TransactionHandle]
//   - [ANERequest.SetTransactionHandle]
//   - [ANERequest.Validate]
//   - [ANERequest.WeightsBuffer]
//   - [ANERequest.InitWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndexSharedEventsTransactionHandle]
//   - [ANERequest.InitWithVirtualModel]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest
type ANERequest struct {
	objectivec.Object
}

// ANERequestFromID constructs a [ANERequest] from an objc.ID.
func ANERequestFromID(id objc.ID) ANERequest {
	return ANERequest{objectivec.Object{id}}
}
// Ensure ANERequest implements IANERequest.
var _ IANERequest = ANERequest{}





// An interface definition for the [ANERequest] class.
//
// # Methods
//
//   - [IANERequest.InputArray]
//   - [IANERequest.InputIndexArray]
//   - [IANERequest.IoSurfacesCount]
//   - [IANERequest.OutputArray]
//   - [IANERequest.OutputIndexArray]
//   - [IANERequest.PerfStats]
//   - [IANERequest.SetPerfStats]
//   - [IANERequest.PerfStatsArray]
//   - [IANERequest.ProcedureIndex]
//   - [IANERequest.SetCompletionHandler]
//   - [IANERequest.SharedEvents]
//   - [IANERequest.SetSharedEvents]
//   - [IANERequest.TransactionHandle]
//   - [IANERequest.SetTransactionHandle]
//   - [IANERequest.Validate]
//   - [IANERequest.WeightsBuffer]
//   - [IANERequest.InitWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndexSharedEventsTransactionHandle]
//   - [IANERequest.InitWithVirtualModel]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest
type IANERequest interface {
	objectivec.IObject

	// Topic: Methods

	InputArray() foundation.INSArray
	InputIndexArray() foundation.INSArray
	IoSurfacesCount() uint64
	OutputArray() foundation.INSArray
	OutputIndexArray() foundation.INSArray
	PerfStats() *ANEPerformanceStats
	SetPerfStats(value *ANEPerformanceStats)
	PerfStatsArray() foundation.INSArray
	ProcedureIndex() foundation.NSNumber
	SetCompletionHandler(handler BoolErrorHandler)
	SharedEvents() *ANESharedEvents
	SetSharedEvents(value *ANESharedEvents)
	TransactionHandle() foundation.NSNumber
	SetTransactionHandle(value foundation.NSNumber)
	Validate() bool
	WeightsBuffer() *ANEIOSurfaceObject
	InitWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndexSharedEventsTransactionHandle(inputs objectivec.IObject, indices objectivec.IObject, outputs objectivec.IObject, indices2 objectivec.IObject, buffer objectivec.IObject, stats objectivec.IObject, index objectivec.IObject, events objectivec.IObject, handle objectivec.IObject) ANERequest
	InitWithVirtualModel(model unsafe.Pointer) ANERequest
}





// Init initializes the instance.
func (a ANERequest) Init() ANERequest {
	rv := objc.Send[ANERequest](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANERequest) Autorelease() ANERequest {
	rv := objc.Send[ANERequest](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANERequest creates a new ANERequest instance.
func NewANERequest() ANERequest {
	class := getANERequestClass()
	rv := objc.Send[ANERequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/initWithInputs:inputIndices:outputs:outputIndices:weightsBuffer:perfStats:procedureIndex:sharedEvents:transactionHandle:
func NewANERequestWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndexSharedEventsTransactionHandle(inputs objectivec.IObject, indices objectivec.IObject, outputs objectivec.IObject, indices2 objectivec.IObject, buffer objectivec.IObject, stats objectivec.IObject, index objectivec.IObject, events objectivec.IObject, handle objectivec.IObject) ANERequest {
	instance := getANERequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInputs:inputIndices:outputs:outputIndices:weightsBuffer:perfStats:procedureIndex:sharedEvents:transactionHandle:"), inputs, indices, outputs, indices2, buffer, stats, index, events, handle)
	return ANERequestFromID(rv)
}


//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/initWithVirtualModel:
func NewANERequestWithVirtualModel(model unsafe.Pointer) ANERequest {
	instance := getANERequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVirtualModel:"), model)
	return ANERequestFromID(rv)
}







// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/ioSurfacesCount
func (a ANERequest) IoSurfacesCount() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("ioSurfacesCount"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/setCompletionHandler:
func (a ANERequest) SetCompletionHandler(handler BoolErrorHandler) {
		_block0, _ := NewBoolErrorBlock(handler)
		objc.Send[objc.ID](a.ID, objc.Sel("setCompletionHandler:"), _block0)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/validate
func (a ANERequest) Validate() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("validate"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/initWithInputs:inputIndices:outputs:outputIndices:weightsBuffer:perfStats:procedureIndex:sharedEvents:transactionHandle:
func (a ANERequest) InitWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndexSharedEventsTransactionHandle(inputs objectivec.IObject, indices objectivec.IObject, outputs objectivec.IObject, indices2 objectivec.IObject, buffer objectivec.IObject, stats objectivec.IObject, index objectivec.IObject, events objectivec.IObject, handle objectivec.IObject) ANERequest {
	rv := objc.Send[ANERequest](a.ID, objc.Sel("initWithInputs:inputIndices:outputs:outputIndices:weightsBuffer:perfStats:procedureIndex:sharedEvents:transactionHandle:"), inputs, indices, outputs, indices2, buffer, stats, index, events, handle)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/initWithVirtualModel:
func (a ANERequest) InitWithVirtualModel(model unsafe.Pointer) ANERequest {
	rv := objc.Send[ANERequest](a.ID, objc.Sel("initWithVirtualModel:"), model)
	return rv
}





//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/requestWithInputs:inputIndices:outputs:outputIndices:perfStats:procedureIndex:
func (_ANERequestClass ANERequestClass) RequestWithInputsInputIndicesOutputsOutputIndicesPerfStatsProcedureIndex(inputs objectivec.IObject, indices objectivec.IObject, outputs objectivec.IObject, indices2 objectivec.IObject, stats objectivec.IObject, index objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANERequestClass.class), objc.Sel("requestWithInputs:inputIndices:outputs:outputIndices:perfStats:procedureIndex:"), inputs, indices, outputs, indices2, stats, index)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/requestWithInputs:inputIndices:outputs:outputIndices:procedureIndex:
func (_ANERequestClass ANERequestClass) RequestWithInputsInputIndicesOutputsOutputIndicesProcedureIndex(inputs objectivec.IObject, indices objectivec.IObject, outputs objectivec.IObject, indices2 objectivec.IObject, index objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANERequestClass.class), objc.Sel("requestWithInputs:inputIndices:outputs:outputIndices:procedureIndex:"), inputs, indices, outputs, indices2, index)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/requestWithInputs:inputIndices:outputs:outputIndices:weightsBuffer:perfStats:procedureIndex:
func (_ANERequestClass ANERequestClass) RequestWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndex(inputs objectivec.IObject, indices objectivec.IObject, outputs objectivec.IObject, indices2 objectivec.IObject, buffer objectivec.IObject, stats objectivec.IObject, index objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANERequestClass.class), objc.Sel("requestWithInputs:inputIndices:outputs:outputIndices:weightsBuffer:perfStats:procedureIndex:"), inputs, indices, outputs, indices2, buffer, stats, index)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/requestWithInputs:inputIndices:outputs:outputIndices:weightsBuffer:perfStats:procedureIndex:sharedEvents:
func (_ANERequestClass ANERequestClass) RequestWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndexSharedEvents(inputs objectivec.IObject, indices objectivec.IObject, outputs objectivec.IObject, indices2 objectivec.IObject, buffer objectivec.IObject, stats objectivec.IObject, index objectivec.IObject, events objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANERequestClass.class), objc.Sel("requestWithInputs:inputIndices:outputs:outputIndices:weightsBuffer:perfStats:procedureIndex:sharedEvents:"), inputs, indices, outputs, indices2, buffer, stats, index, events)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/requestWithInputs:inputIndices:outputs:outputIndices:weightsBuffer:perfStats:procedureIndex:sharedEvents:transactionHandle:
func (_ANERequestClass ANERequestClass) RequestWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndexSharedEventsTransactionHandle(inputs objectivec.IObject, indices objectivec.IObject, outputs objectivec.IObject, indices2 objectivec.IObject, buffer objectivec.IObject, stats objectivec.IObject, index objectivec.IObject, events objectivec.IObject, handle objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANERequestClass.class), objc.Sel("requestWithInputs:inputIndices:outputs:outputIndices:weightsBuffer:perfStats:procedureIndex:sharedEvents:transactionHandle:"), inputs, indices, outputs, indices2, buffer, stats, index, events, handle)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/requestWithInputs:inputIndices:outputs:outputIndices:weightsBuffer:procedureIndex:
func (_ANERequestClass ANERequestClass) RequestWithInputsInputIndicesOutputsOutputIndicesWeightsBufferProcedureIndex(inputs objectivec.IObject, indices objectivec.IObject, outputs objectivec.IObject, indices2 objectivec.IObject, buffer objectivec.IObject, index objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANERequestClass.class), objc.Sel("requestWithInputs:inputIndices:outputs:outputIndices:weightsBuffer:procedureIndex:"), inputs, indices, outputs, indices2, buffer, index)
	return objectivec.Object{ID: rv}
}








// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/inputArray
func (a ANERequest) InputArray() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputArray"))
	return foundation.NSArrayFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/inputIndexArray
func (a ANERequest) InputIndexArray() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputIndexArray"))
	return foundation.NSArrayFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/outputArray
func (a ANERequest) OutputArray() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputArray"))
	return foundation.NSArrayFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/outputIndexArray
func (a ANERequest) OutputIndexArray() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputIndexArray"))
	return foundation.NSArrayFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/perfStats
func (a ANERequest) PerfStats() *ANEPerformanceStats {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("perfStats"))
	if rv == 0 {
		return nil
	}
	val := ANEPerformanceStatsFromID(objc.ID(rv))
	return &val
}
func (a ANERequest) SetPerfStats(value *ANEPerformanceStats) {
	objc.Send[struct{}](a.ID, objc.Sel("setPerfStats:"), value)
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/perfStatsArray
func (a ANERequest) PerfStatsArray() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("perfStatsArray"))
	return foundation.NSArrayFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/procedureIndex
func (a ANERequest) ProcedureIndex() foundation.NSNumber {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("procedureIndex"))
	return foundation.NSNumberFromID(objc.ID(rv))
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/sharedEvents
func (a ANERequest) SharedEvents() *ANESharedEvents {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sharedEvents"))
	if rv == 0 {
		return nil
	}
	val := ANESharedEventsFromID(objc.ID(rv))
	return &val
}
func (a ANERequest) SetSharedEvents(value *ANESharedEvents) {
	objc.Send[struct{}](a.ID, objc.Sel("setSharedEvents:"), value)
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/transactionHandle
func (a ANERequest) TransactionHandle() foundation.NSNumber {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("transactionHandle"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (a ANERequest) SetTransactionHandle(value foundation.NSNumber) {
	objc.Send[struct{}](a.ID, objc.Sel("setTransactionHandle:"), value)
}



// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANERequest/weightsBuffer
func (a ANERequest) WeightsBuffer() *ANEIOSurfaceObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("weightsBuffer"))
	if rv == 0 {
		return nil
	}
	val := ANEIOSurfaceObjectFromID(objc.ID(rv))
	return &val
}












// Set is a synchronous wrapper around [ANERequest.SetCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a ANERequest) Set(ctx context.Context) (bool, error) {
	type result struct {
		val bool
		err error
	}
	done := make(chan result, 1)
	a.SetCompletionHandler(func(val bool, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return false, ctx.Err()
	}
}






