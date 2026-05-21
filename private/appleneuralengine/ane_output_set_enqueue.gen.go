// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANEOutputSetEnqueue] class.
var (
	_ANEOutputSetEnqueueClass     ANEOutputSetEnqueueClass
	_ANEOutputSetEnqueueClassOnce sync.Once
)

func getANEOutputSetEnqueueClass() ANEOutputSetEnqueueClass {
	_ANEOutputSetEnqueueClassOnce.Do(func() {
		_ANEOutputSetEnqueueClass = ANEOutputSetEnqueueClass{class: objc.GetClass("_ANEOutputSetEnqueue")}
	})
	return _ANEOutputSetEnqueueClass
}

// GetANEOutputSetEnqueueClass returns the class object for _ANEOutputSetEnqueue.
func GetANEOutputSetEnqueueClass() ANEOutputSetEnqueueClass {
	return getANEOutputSetEnqueueClass()
}

type ANEOutputSetEnqueueClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANEOutputSetEnqueueClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANEOutputSetEnqueueClass) Alloc() ANEOutputSetEnqueue {
	rv := objc.Send[ANEOutputSetEnqueue](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ANEOutputSetEnqueue.IsOpenLoop]
//   - [ANEOutputSetEnqueue.ProcedureIndex]
//   - [ANEOutputSetEnqueue.SetIndex]
//   - [ANEOutputSetEnqueue.SignalNotRequired]
//   - [ANEOutputSetEnqueue.SignalValue]
//   - [ANEOutputSetEnqueue.InitOutputSetWithProcedureIndexSetIndexSignalValueSignalNotRequiredIsOpenLoop]
type ANEOutputSetEnqueue struct {
	objectivec.Object
}

// ANEOutputSetEnqueueFromID constructs a [ANEOutputSetEnqueue] from an objc.ID.
func ANEOutputSetEnqueueFromID(id objc.ID) ANEOutputSetEnqueue {
	return ANEOutputSetEnqueue{objectivec.Object{ID: id}}
}

// Ensure ANEOutputSetEnqueue implements IANEOutputSetEnqueue.
var _ IANEOutputSetEnqueue = ANEOutputSetEnqueue{}

// An interface definition for the [ANEOutputSetEnqueue] class.
//
// # Methods
//
//   - [IANEOutputSetEnqueue.IsOpenLoop]
//   - [IANEOutputSetEnqueue.ProcedureIndex]
//   - [IANEOutputSetEnqueue.SetIndex]
//   - [IANEOutputSetEnqueue.SignalNotRequired]
//   - [IANEOutputSetEnqueue.SignalValue]
//   - [IANEOutputSetEnqueue.InitOutputSetWithProcedureIndexSetIndexSignalValueSignalNotRequiredIsOpenLoop]
type IANEOutputSetEnqueue interface {
	objectivec.IObject

	// Topic: Methods

	IsOpenLoop() bool
	ProcedureIndex() uint32
	SetIndex() uint32
	SignalNotRequired() bool
	SignalValue() uint64
	InitOutputSetWithProcedureIndexSetIndexSignalValueSignalNotRequiredIsOpenLoop(index uint32, index2 uint32, value uint64, required bool, loop bool) ANEOutputSetEnqueue
}

// Init initializes the instance.
func (a ANEOutputSetEnqueue) Init() ANEOutputSetEnqueue {
	rv := objc.Send[ANEOutputSetEnqueue](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANEOutputSetEnqueue) Autorelease() ANEOutputSetEnqueue {
	rv := objc.Send[ANEOutputSetEnqueue](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANEOutputSetEnqueue creates a new ANEOutputSetEnqueue instance.
func NewANEOutputSetEnqueue() ANEOutputSetEnqueue {
	class := getANEOutputSetEnqueueClass()
	rv := objc.Send[ANEOutputSetEnqueue](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewANEOutputSetEnqueueOutputSetWithProcedureIndexSetIndexSignalValueSignalNotRequiredIsOpenLoop(index uint32, index2 uint32, value uint64, required bool, loop bool) ANEOutputSetEnqueue {
	instance := getANEOutputSetEnqueueClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initOutputSetWithProcedureIndex:setIndex:signalValue:signalNotRequired:isOpenLoop:"), index, index2, value, required, loop)
	return ANEOutputSetEnqueueFromID(rv)
}

func (a ANEOutputSetEnqueue) InitOutputSetWithProcedureIndexSetIndexSignalValueSignalNotRequiredIsOpenLoop(index uint32, index2 uint32, value uint64, required bool, loop bool) ANEOutputSetEnqueue {
	rv := objc.Send[ANEOutputSetEnqueue](a.ID, objc.Sel("initOutputSetWithProcedureIndex:setIndex:signalValue:signalNotRequired:isOpenLoop:"), index, index2, value, required, loop)
	return rv
}

func (_ANEOutputSetEnqueueClass ANEOutputSetEnqueueClass) OutputSetWithProcedureIndexSetIndexSignalValueSignalNotRequiredIsOpenLoop(index uint32, index2 uint32, value uint64, required bool, loop bool) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANEOutputSetEnqueueClass.class), objc.Sel("outputSetWithProcedureIndex:setIndex:signalValue:signalNotRequired:isOpenLoop:"), index, index2, value, required, loop)
	return objectivec.Object{ID: rv}
}

func (a ANEOutputSetEnqueue) IsOpenLoop() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("isOpenLoop"))
	return rv
}
func (a ANEOutputSetEnqueue) ProcedureIndex() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("procedureIndex"))
	return rv
}
func (a ANEOutputSetEnqueue) SetIndex() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("setIndex"))
	return rv
}
func (a ANEOutputSetEnqueue) SignalNotRequired() bool {
	rv := objc.Send[bool](a.ID, objc.Sel("signalNotRequired"))
	return rv
}
func (a ANEOutputSetEnqueue) SignalValue() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("signalValue"))
	return rv
}
