// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMTLLateEvalEvent] class.
var (
	_IOGPUMTLLateEvalEventClass     IOGPUMTLLateEvalEventClass
	_IOGPUMTLLateEvalEventClassOnce sync.Once
)

func getIOGPUMTLLateEvalEventClass() IOGPUMTLLateEvalEventClass {
	_IOGPUMTLLateEvalEventClassOnce.Do(func() {
		_IOGPUMTLLateEvalEventClass = IOGPUMTLLateEvalEventClass{class: objc.GetClass("IOGPUMTLLateEvalEvent")}
	})
	return _IOGPUMTLLateEvalEventClass
}

// GetIOGPUMTLLateEvalEventClass returns the class object for IOGPUMTLLateEvalEvent.
func GetIOGPUMTLLateEvalEventClass() IOGPUMTLLateEvalEventClass {
	return getIOGPUMTLLateEvalEventClass()
}

type IOGPUMTLLateEvalEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMTLLateEvalEventClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMTLLateEvalEventClass) Alloc() IOGPUMTLLateEvalEvent {
	rv := objc.SendIfResponds[IOGPUMTLLateEvalEvent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMTLLateEvalEvent.SignaledValue]
//   - [IOGPUMTLLateEvalEvent.SetSignaledValue]
//   - [IOGPUMTLLateEvalEvent.InitWithDevice]
type IOGPUMTLLateEvalEvent struct {
	objectivec.Object
}

// IOGPUMTLLateEvalEventFromID constructs a [IOGPUMTLLateEvalEvent] from an objc.ID.
func IOGPUMTLLateEvalEventFromID(id objc.ID) IOGPUMTLLateEvalEvent {
	return IOGPUMTLLateEvalEvent{objectivec.Object{ID: id}}
}

// Ensure IOGPUMTLLateEvalEvent implements IIOGPUMTLLateEvalEvent.
var _ IIOGPUMTLLateEvalEvent = IOGPUMTLLateEvalEvent{}

// An interface definition for the [IOGPUMTLLateEvalEvent] class.
//
// # Methods
//
//   - [IIOGPUMTLLateEvalEvent.SignaledValue]
//   - [IIOGPUMTLLateEvalEvent.SetSignaledValue]
//   - [IIOGPUMTLLateEvalEvent.InitWithDevice]
type IIOGPUMTLLateEvalEvent interface {
	objectivec.IObject

	// Topic: Methods

	SignaledValue() uint64
	SetSignaledValue(value uint64)
	InitWithDevice(device *uintptr) IOGPUMTLLateEvalEvent
}

// Init initializes the instance.
func (i IOGPUMTLLateEvalEvent) Init() IOGPUMTLLateEvalEvent {
	rv := objc.SendIfResponds[IOGPUMTLLateEvalEvent](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMTLLateEvalEvent) Autorelease() IOGPUMTLLateEvalEvent {
	rv := objc.SendIfResponds[IOGPUMTLLateEvalEvent](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMTLLateEvalEvent creates a new IOGPUMTLLateEvalEvent instance.
func NewIOGPUMTLLateEvalEvent() IOGPUMTLLateEvalEvent {
	class := getIOGPUMTLLateEvalEventClass()
	rv := objc.SendIfResponds[IOGPUMTLLateEvalEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMTLLateEvalEventWithDevice(device *uintptr) IOGPUMTLLateEvalEvent {
	instance := getIOGPUMTLLateEvalEventClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return IOGPUMTLLateEvalEventFromID(rv)
}

func (i IOGPUMTLLateEvalEvent) InitWithDevice(device *uintptr) IOGPUMTLLateEvalEvent {
	rv := objc.SendIfResponds[IOGPUMTLLateEvalEvent](i.ID, objc.Sel("initWithDevice:"), device)
	return rv
}

func (i IOGPUMTLLateEvalEvent) SignaledValue() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("signaledValue"))
	return rv
}
func (i IOGPUMTLLateEvalEvent) SetSignaledValue(value uint64) {
	objc.SendIfResponds[struct{}](i.ID, objc.Sel("setSignaledValue:"), value)
}
