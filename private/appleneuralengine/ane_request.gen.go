// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
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

// Class returns the underlying Objective-C class pointer.
func (ac ANERequestClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANERequestClass) Alloc() ANERequest {
	rv := objc.Send[ANERequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

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
type ANERequest struct {
	objectivec.Object
}

// ANERequestFromID constructs a [ANERequest] from an objc.ID.
func ANERequestFromID(id objc.ID) ANERequest {
	return ANERequest{objectivec.Object{ID: id}}
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
type IANERequest interface {
	objectivec.IObject

	// Topic: Methods

	InputArray() foundation.INSArray
	InputIndexArray() foundation.INSArray
	IoSurfacesCount() uint64
	OutputArray() foundation.INSArray
	OutputIndexArray() foundation.INSArray
	PerfStats() IANEPerformanceStats
	SetPerfStats(value IANEPerformanceStats)
	PerfStatsArray() foundation.INSArray
	ProcedureIndex() foundation.NSNumber
	SetCompletionHandler(handler ErrorHandler)
	SharedEvents() IANESharedEvents
	SetSharedEvents(value IANESharedEvents)
	TransactionHandle() foundation.NSNumber
	SetTransactionHandle(value foundation.NSNumber)
	Validate() bool
	WeightsBuffer() IANEIOSurfaceObject
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

func NewANERequestWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndexSharedEventsTransactionHandle(inputs objectivec.IObject, indices objectivec.IObject, outputs objectivec.IObject, indices2 objectivec.IObject, buffer objectivec.IObject, stats objectivec.IObject, index objectivec.IObject, events objectivec.IObject, handle objectivec.IObject) ANERequest {
	instance := getANERequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInputs:inputIndices:outputs:outputIndices:weightsBuffer:perfStats:procedureIndex:sharedEvents:transactionHandle:"), inputs, indices, outputs, indices2, buffer, stats, index, events, handle)
	return ANERequestFromID(rv)
}

func NewANERequestWithVirtualModel(model unsafe.Pointer) ANERequest {
	instance := getANERequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVirtualModel:"), model)
	return ANERequestFromID(rv)
}

func (a ANERequest) IoSurfacesCount() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("ioSurfacesCount"))
	return rv
}
func (a ANERequest) SetCompletionHandler(handler ErrorHandler) {
	_block0, _ := NewErrorBlock(handler)
	objc.Send[objc.ID](a.ID, objc.Sel("setCompletionHandler:"), _block0)
}
func (a ANERequest) Validate() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("validate"))
	return rv
}
func (a ANERequest) InitWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndexSharedEventsTransactionHandle(inputs objectivec.IObject, indices objectivec.IObject, outputs objectivec.IObject, indices2 objectivec.IObject, buffer objectivec.IObject, stats objectivec.IObject, index objectivec.IObject, events objectivec.IObject, handle objectivec.IObject) ANERequest {
	rv := objc.Send[ANERequest](a.ID, objc.Sel("initWithInputs:inputIndices:outputs:outputIndices:weightsBuffer:perfStats:procedureIndex:sharedEvents:transactionHandle:"), inputs, indices, outputs, indices2, buffer, stats, index, events, handle)
	return rv
}
func (a ANERequest) InitWithVirtualModel(model unsafe.Pointer) ANERequest {
	rv := objc.Send[ANERequest](a.ID, objc.Sel("initWithVirtualModel:"), model)
	return rv
}

func (_ANERequestClass ANERequestClass) RequestWithInputsInputIndicesOutputsOutputIndicesPerfStatsProcedureIndex(inputs objectivec.IObject, indices objectivec.IObject, outputs objectivec.IObject, indices2 objectivec.IObject, stats objectivec.IObject, index objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANERequestClass.class), objc.Sel("requestWithInputs:inputIndices:outputs:outputIndices:perfStats:procedureIndex:"), inputs, indices, outputs, indices2, stats, index)
	return objectivec.Object{ID: rv}
}
func (_ANERequestClass ANERequestClass) RequestWithInputsInputIndicesOutputsOutputIndicesProcedureIndex(inputs objectivec.IObject, indices objectivec.IObject, outputs objectivec.IObject, indices2 objectivec.IObject, index objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANERequestClass.class), objc.Sel("requestWithInputs:inputIndices:outputs:outputIndices:procedureIndex:"), inputs, indices, outputs, indices2, index)
	return objectivec.Object{ID: rv}
}
func (_ANERequestClass ANERequestClass) RequestWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndex(inputs objectivec.IObject, indices objectivec.IObject, outputs objectivec.IObject, indices2 objectivec.IObject, buffer objectivec.IObject, stats objectivec.IObject, index objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANERequestClass.class), objc.Sel("requestWithInputs:inputIndices:outputs:outputIndices:weightsBuffer:perfStats:procedureIndex:"), inputs, indices, outputs, indices2, buffer, stats, index)
	return objectivec.Object{ID: rv}
}
func (_ANERequestClass ANERequestClass) RequestWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndexSharedEvents(inputs objectivec.IObject, indices objectivec.IObject, outputs objectivec.IObject, indices2 objectivec.IObject, buffer objectivec.IObject, stats objectivec.IObject, index objectivec.IObject, events objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANERequestClass.class), objc.Sel("requestWithInputs:inputIndices:outputs:outputIndices:weightsBuffer:perfStats:procedureIndex:sharedEvents:"), inputs, indices, outputs, indices2, buffer, stats, index, events)
	return objectivec.Object{ID: rv}
}
func (_ANERequestClass ANERequestClass) RequestWithInputsInputIndicesOutputsOutputIndicesWeightsBufferPerfStatsProcedureIndexSharedEventsTransactionHandle(inputs objectivec.IObject, indices objectivec.IObject, outputs objectivec.IObject, indices2 objectivec.IObject, buffer objectivec.IObject, stats objectivec.IObject, index objectivec.IObject, events objectivec.IObject, handle objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANERequestClass.class), objc.Sel("requestWithInputs:inputIndices:outputs:outputIndices:weightsBuffer:perfStats:procedureIndex:sharedEvents:transactionHandle:"), inputs, indices, outputs, indices2, buffer, stats, index, events, handle)
	return objectivec.Object{ID: rv}
}
func (_ANERequestClass ANERequestClass) RequestWithInputsInputIndicesOutputsOutputIndicesWeightsBufferProcedureIndex(inputs objectivec.IObject, indices objectivec.IObject, outputs objectivec.IObject, indices2 objectivec.IObject, buffer objectivec.IObject, index objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANERequestClass.class), objc.Sel("requestWithInputs:inputIndices:outputs:outputIndices:weightsBuffer:procedureIndex:"), inputs, indices, outputs, indices2, buffer, index)
	return objectivec.Object{ID: rv}
}

func (a ANERequest) InputArray() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputArray"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (a ANERequest) InputIndexArray() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputIndexArray"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (a ANERequest) OutputArray() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputArray"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (a ANERequest) OutputIndexArray() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputIndexArray"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (a ANERequest) PerfStats() IANEPerformanceStats {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("perfStats"))
	return ANEPerformanceStatsFromID(objc.ID(rv))
}
func (a ANERequest) SetPerfStats(value IANEPerformanceStats) {
	objc.Send[struct{}](a.ID, objc.Sel("setPerfStats:"), value)
}
func (a ANERequest) PerfStatsArray() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("perfStatsArray"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (a ANERequest) ProcedureIndex() foundation.NSNumber {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("procedureIndex"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (a ANERequest) SharedEvents() IANESharedEvents {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sharedEvents"))
	return ANESharedEventsFromID(objc.ID(rv))
}
func (a ANERequest) SetSharedEvents(value IANESharedEvents) {
	objc.Send[struct{}](a.ID, objc.Sel("setSharedEvents:"), value)
}
func (a ANERequest) TransactionHandle() foundation.NSNumber {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("transactionHandle"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
func (a ANERequest) SetTransactionHandle(value foundation.NSNumber) {
	objc.Send[struct{}](a.ID, objc.Sel("setTransactionHandle:"), value)
}
func (a ANERequest) WeightsBuffer() IANEIOSurfaceObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("weightsBuffer"))
	return ANEIOSurfaceObjectFromID(objc.ID(rv))
}

// Set is a synchronous wrapper around [ANERequest.SetCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (a ANERequest) Set(ctx context.Context) error {
	done := make(chan error, 1)
	a.SetCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
