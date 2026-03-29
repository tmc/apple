// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEChainingRequest] class.
var (
	_ANEChainingRequestClass     ANEChainingRequestClass
	_ANEChainingRequestClassOnce sync.Once
)

func getANEChainingRequestClass() ANEChainingRequestClass {
	_ANEChainingRequestClassOnce.Do(func() {
		_ANEChainingRequestClass = ANEChainingRequestClass{class: objc.GetClass("_ANEChainingRequest")}
	})
	return _ANEChainingRequestClass
}

// GetANEChainingRequestClass returns the class object for _ANEChainingRequest.
func GetANEChainingRequestClass() ANEChainingRequestClass {
	return getANEChainingRequestClass()
}

type ANEChainingRequestClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEChainingRequestClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEChainingRequestClass) Alloc() ANEChainingRequest {
	rv := objc.Send[ANEChainingRequest](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ANEChainingRequest.EncodeWithCoder]
//   - [ANEChainingRequest.FwEnqueueDelay]
//   - [ANEChainingRequest.InputBuffer]
//   - [ANEChainingRequest.LoopbackInputSymbolIndex]
//   - [ANEChainingRequest.LoopbackOutputSymbolIndex]
//   - [ANEChainingRequest.MemoryPoolId]
//   - [ANEChainingRequest.OutputSets]
//   - [ANEChainingRequest.ProcedureIndex]
//   - [ANEChainingRequest.SignalEvents]
//   - [ANEChainingRequest.TransactionHandle]
//   - [ANEChainingRequest.Validate]
//   - [ANEChainingRequest.InitWithCoder]
//   - [ANEChainingRequest.InitWithInputsOutputsLbInputSymbolIdLbOutputSymbolIdProcedureIndexSignalEventsTransactionHandleFwEnqueueDelayMemoryPoolId]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEChainingRequest
type ANEChainingRequest struct {
	objectivec.Object
}

// ANEChainingRequestFromID constructs a [ANEChainingRequest] from an objc.ID.
func ANEChainingRequestFromID(id objc.ID) ANEChainingRequest {
	return ANEChainingRequest{objectivec.Object{ID: id}}
}
// Ensure ANEChainingRequest implements IANEChainingRequest.
var _ IANEChainingRequest = ANEChainingRequest{}

// An interface definition for the [ANEChainingRequest] class.
//
// # Methods
//
//   - [IANEChainingRequest.EncodeWithCoder]
//   - [IANEChainingRequest.FwEnqueueDelay]
//   - [IANEChainingRequest.InputBuffer]
//   - [IANEChainingRequest.LoopbackInputSymbolIndex]
//   - [IANEChainingRequest.LoopbackOutputSymbolIndex]
//   - [IANEChainingRequest.MemoryPoolId]
//   - [IANEChainingRequest.OutputSets]
//   - [IANEChainingRequest.ProcedureIndex]
//   - [IANEChainingRequest.SignalEvents]
//   - [IANEChainingRequest.TransactionHandle]
//   - [IANEChainingRequest.Validate]
//   - [IANEChainingRequest.InitWithCoder]
//   - [IANEChainingRequest.InitWithInputsOutputsLbInputSymbolIdLbOutputSymbolIdProcedureIndexSignalEventsTransactionHandleFwEnqueueDelayMemoryPoolId]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEChainingRequest
type IANEChainingRequest interface {
	objectivec.IObject

	// Topic: Methods

	EncodeWithCoder(coder foundation.INSCoder)
	FwEnqueueDelay() foundation.NSNumber
	InputBuffer() foundation.INSArray
	LoopbackInputSymbolIndex() foundation.INSArray
	LoopbackOutputSymbolIndex() foundation.INSArray
	MemoryPoolId() foundation.NSNumber
	OutputSets() foundation.INSArray
	ProcedureIndex() foundation.NSNumber
	SignalEvents() foundation.INSArray
	TransactionHandle() foundation.NSNumber
	Validate() bool
	InitWithCoder(coder foundation.INSCoder) ANEChainingRequest
	InitWithInputsOutputsLbInputSymbolIdLbOutputSymbolIdProcedureIndexSignalEventsTransactionHandleFwEnqueueDelayMemoryPoolId(inputs objectivec.IObject, outputs objectivec.IObject, id objectivec.IObject, id2 objectivec.IObject, index objectivec.IObject, events objectivec.IObject, handle objectivec.IObject, delay objectivec.IObject, id3 objectivec.IObject) ANEChainingRequest
}

// Init initializes the instance.
func (a ANEChainingRequest) Init() ANEChainingRequest {
	rv := objc.Send[ANEChainingRequest](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEChainingRequest) Autorelease() ANEChainingRequest {
	rv := objc.Send[ANEChainingRequest](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEChainingRequest creates a new ANEChainingRequest instance.
func NewANEChainingRequest() ANEChainingRequest {
	class := getANEChainingRequestClass()
	rv := objc.Send[ANEChainingRequest](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEChainingRequest/initWithCoder:
func NewANEChainingRequestWithCoder(coder objectivec.IObject) ANEChainingRequest {
	instance := getANEChainingRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return ANEChainingRequestFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEChainingRequest/initWithInputs:outputs:lbInputSymbolId:lbOutputSymbolId:procedureIndex:signalEvents:transactionHandle:fwEnqueueDelay:memoryPoolId:
func NewANEChainingRequestWithInputsOutputsLbInputSymbolIdLbOutputSymbolIdProcedureIndexSignalEventsTransactionHandleFwEnqueueDelayMemoryPoolId(inputs objectivec.IObject, outputs objectivec.IObject, id objectivec.IObject, id2 objectivec.IObject, index objectivec.IObject, events objectivec.IObject, handle objectivec.IObject, delay objectivec.IObject, id3 objectivec.IObject) ANEChainingRequest {
	instance := getANEChainingRequestClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInputs:outputs:lbInputSymbolId:lbOutputSymbolId:procedureIndex:signalEvents:transactionHandle:fwEnqueueDelay:memoryPoolId:"), inputs, outputs, id, id2, index, events, handle, delay, id3)
	return ANEChainingRequestFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEChainingRequest/encodeWithCoder:
func (a ANEChainingRequest) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEChainingRequest/validate
func (a ANEChainingRequest) Validate() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("validate"))
	return rv
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEChainingRequest/initWithCoder:
func (a ANEChainingRequest) InitWithCoder(coder foundation.INSCoder) ANEChainingRequest {
	rv := objc.Send[ANEChainingRequest](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEChainingRequest/initWithInputs:outputs:lbInputSymbolId:lbOutputSymbolId:procedureIndex:signalEvents:transactionHandle:fwEnqueueDelay:memoryPoolId:
func (a ANEChainingRequest) InitWithInputsOutputsLbInputSymbolIdLbOutputSymbolIdProcedureIndexSignalEventsTransactionHandleFwEnqueueDelayMemoryPoolId(inputs objectivec.IObject, outputs objectivec.IObject, id objectivec.IObject, id2 objectivec.IObject, index objectivec.IObject, events objectivec.IObject, handle objectivec.IObject, delay objectivec.IObject, id3 objectivec.IObject) ANEChainingRequest {
	rv := objc.Send[ANEChainingRequest](a.ID, objc.Sel("initWithInputs:outputs:lbInputSymbolId:lbOutputSymbolId:procedureIndex:signalEvents:transactionHandle:fwEnqueueDelay:memoryPoolId:"), inputs, outputs, id, id2, index, events, handle, delay, id3)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEChainingRequest/chainingRequestWithInputs:outputSets:lbInputSymbolId:lbOutputSymbolId:procedureIndex:signalEvents:transactionHandle:fwEnqueueDelay:memoryPoolId:
func (_ANEChainingRequestClass ANEChainingRequestClass) ChainingRequestWithInputsOutputSetsLbInputSymbolIdLbOutputSymbolIdProcedureIndexSignalEventsTransactionHandleFwEnqueueDelayMemoryPoolId(inputs objectivec.IObject, sets objectivec.IObject, id objectivec.IObject, id2 objectivec.IObject, index objectivec.IObject, events objectivec.IObject, handle objectivec.IObject, delay objectivec.IObject, id3 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEChainingRequestClass.class), objc.Sel("chainingRequestWithInputs:outputSets:lbInputSymbolId:lbOutputSymbolId:procedureIndex:signalEvents:transactionHandle:fwEnqueueDelay:memoryPoolId:"), inputs, sets, id, id2, index, events, handle, delay, id3)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEChainingRequest/supportsSecureCoding
func (_ANEChainingRequestClass ANEChainingRequestClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_ANEChainingRequestClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEChainingRequest/fwEnqueueDelay
func (a ANEChainingRequest) FwEnqueueDelay() foundation.NSNumber {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("fwEnqueueDelay"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEChainingRequest/inputBuffer
func (a ANEChainingRequest) InputBuffer() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputBuffer"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEChainingRequest/loopbackInputSymbolIndex
func (a ANEChainingRequest) LoopbackInputSymbolIndex() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("loopbackInputSymbolIndex"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEChainingRequest/loopbackOutputSymbolIndex
func (a ANEChainingRequest) LoopbackOutputSymbolIndex() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("loopbackOutputSymbolIndex"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEChainingRequest/memoryPoolId
func (a ANEChainingRequest) MemoryPoolId() foundation.NSNumber {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("memoryPoolId"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEChainingRequest/outputSets
func (a ANEChainingRequest) OutputSets() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("outputSets"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEChainingRequest/procedureIndex
func (a ANEChainingRequest) ProcedureIndex() foundation.NSNumber {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("procedureIndex"))
	return foundation.NSNumberFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEChainingRequest/signalEvents
func (a ANEChainingRequest) SignalEvents() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("signalEvents"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEChainingRequest/transactionHandle
func (a ANEChainingRequest) TransactionHandle() foundation.NSNumber {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("transactionHandle"))
	return foundation.NSNumberFromID(objc.ID(rv))
}

