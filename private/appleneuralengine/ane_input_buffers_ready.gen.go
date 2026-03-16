// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEInputBuffersReady] class.
var (
	_ANEInputBuffersReadyClass     ANEInputBuffersReadyClass
	_ANEInputBuffersReadyClassOnce sync.Once
)

func getANEInputBuffersReadyClass() ANEInputBuffersReadyClass {
	_ANEInputBuffersReadyClassOnce.Do(func() {
		_ANEInputBuffersReadyClass = ANEInputBuffersReadyClass{class: objc.GetClass("_ANEInputBuffersReady")}
	})
	return _ANEInputBuffersReadyClass
}

// GetANEInputBuffersReadyClass returns the class object for _ANEInputBuffersReady.
func GetANEInputBuffersReadyClass() ANEInputBuffersReadyClass {
	return getANEInputBuffersReadyClass()
}

type ANEInputBuffersReadyClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEInputBuffersReadyClass) Alloc() ANEInputBuffersReady {
	rv := objc.Send[ANEInputBuffersReady](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ANEInputBuffersReady.ExecutionDelay]
//   - [ANEInputBuffersReady.InputBufferInfoIndex]
//   - [ANEInputBuffersReady.InputFreeValue]
//   - [ANEInputBuffersReady.ProcedureIndex]
//   - [ANEInputBuffersReady.Validate]
//   - [ANEInputBuffersReady.InitInputsProcedureIndexInputBufferInfoIndexInputFreeValueExecutionDelay]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInputBuffersReady
type ANEInputBuffersReady struct {
	objectivec.Object
}

// ANEInputBuffersReadyFromID constructs a [ANEInputBuffersReady] from an objc.ID.
func ANEInputBuffersReadyFromID(id objc.ID) ANEInputBuffersReady {
	return ANEInputBuffersReady{objectivec.Object{id}}
}
// Ensure ANEInputBuffersReady implements IANEInputBuffersReady.
var _ IANEInputBuffersReady = ANEInputBuffersReady{}

// An interface definition for the [ANEInputBuffersReady] class.
//
// # Methods
//
//   - [IANEInputBuffersReady.ExecutionDelay]
//   - [IANEInputBuffersReady.InputBufferInfoIndex]
//   - [IANEInputBuffersReady.InputFreeValue]
//   - [IANEInputBuffersReady.ProcedureIndex]
//   - [IANEInputBuffersReady.Validate]
//   - [IANEInputBuffersReady.InitInputsProcedureIndexInputBufferInfoIndexInputFreeValueExecutionDelay]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInputBuffersReady
type IANEInputBuffersReady interface {
	objectivec.IObject

	// Topic: Methods

	ExecutionDelay() uint64
	InputBufferInfoIndex() foundation.INSArray
	InputFreeValue() foundation.INSArray
	ProcedureIndex() uint32
	Validate() bool
	InitInputsProcedureIndexInputBufferInfoIndexInputFreeValueExecutionDelay(index uint32, index2 objectivec.IObject, value objectivec.IObject, delay uint64) ANEInputBuffersReady
}

// Init initializes the instance.
func (a ANEInputBuffersReady) Init() ANEInputBuffersReady {
	rv := objc.Send[ANEInputBuffersReady](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEInputBuffersReady) Autorelease() ANEInputBuffersReady {
	rv := objc.Send[ANEInputBuffersReady](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEInputBuffersReady creates a new ANEInputBuffersReady instance.
func NewANEInputBuffersReady() ANEInputBuffersReady {
	class := getANEInputBuffersReadyClass()
	rv := objc.Send[ANEInputBuffersReady](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInputBuffersReady/initInputsProcedureIndex:inputBufferInfoIndex:inputFreeValue:executionDelay:
func NewANEInputBuffersReadyInputsProcedureIndexInputBufferInfoIndexInputFreeValueExecutionDelay(index uint32, index2 objectivec.IObject, value objectivec.IObject, delay uint64) ANEInputBuffersReady {
	instance := getANEInputBuffersReadyClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initInputsProcedureIndex:inputBufferInfoIndex:inputFreeValue:executionDelay:"), index, index2, value, delay)
	return ANEInputBuffersReadyFromID(rv)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInputBuffersReady/validate
func (a ANEInputBuffersReady) Validate() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("validate"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInputBuffersReady/initInputsProcedureIndex:inputBufferInfoIndex:inputFreeValue:executionDelay:
func (a ANEInputBuffersReady) InitInputsProcedureIndexInputBufferInfoIndexInputFreeValueExecutionDelay(index uint32, index2 objectivec.IObject, value objectivec.IObject, delay uint64) ANEInputBuffersReady {
	rv := objc.Send[ANEInputBuffersReady](a.ID, objc.Sel("initInputsProcedureIndex:inputBufferInfoIndex:inputFreeValue:executionDelay:"), index, index2, value, delay)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInputBuffersReady/inputBuffersWithProcedureIndex:inputBufferInfoIndex:inputFreeValue:executionDelay:
func (_ANEInputBuffersReadyClass ANEInputBuffersReadyClass) InputBuffersWithProcedureIndexInputBufferInfoIndexInputFreeValueExecutionDelay(index uint32, index2 objectivec.IObject, value objectivec.IObject, delay uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEInputBuffersReadyClass.class), objc.Sel("inputBuffersWithProcedureIndex:inputBufferInfoIndex:inputFreeValue:executionDelay:"), index, index2, value, delay)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInputBuffersReady/executionDelay
func (a ANEInputBuffersReady) ExecutionDelay() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("executionDelay"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInputBuffersReady/inputBufferInfoIndex
func (a ANEInputBuffersReady) InputBufferInfoIndex() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputBufferInfoIndex"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInputBuffersReady/inputFreeValue
func (a ANEInputBuffersReady) InputFreeValue() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("inputFreeValue"))
	return foundation.NSArrayFromID(objc.ID(rv))
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANEInputBuffersReady/procedureIndex
func (a ANEInputBuffersReady) ProcedureIndex() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("procedureIndex"))
	return rv
}

