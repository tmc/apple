// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalMTLLateEvalEvent] class.
var (
	_IOGPUMetalMTLLateEvalEventClass     IOGPUMetalMTLLateEvalEventClass
	_IOGPUMetalMTLLateEvalEventClassOnce sync.Once
)

func getIOGPUMetalMTLLateEvalEventClass() IOGPUMetalMTLLateEvalEventClass {
	_IOGPUMetalMTLLateEvalEventClassOnce.Do(func() {
		_IOGPUMetalMTLLateEvalEventClass = IOGPUMetalMTLLateEvalEventClass{class: objc.GetClass("_IOGPUMetalMTLLateEvalEvent")}
	})
	return _IOGPUMetalMTLLateEvalEventClass
}

// GetIOGPUMetalMTLLateEvalEventClass returns the class object for _IOGPUMetalMTLLateEvalEvent.
func GetIOGPUMetalMTLLateEvalEventClass() IOGPUMetalMTLLateEvalEventClass {
	return getIOGPUMetalMTLLateEvalEventClass()
}

type IOGPUMetalMTLLateEvalEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalMTLLateEvalEventClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalMTLLateEvalEventClass) Alloc() IOGPUMetalMTLLateEvalEvent {
	rv := objc.SendIfResponds[IOGPUMetalMTLLateEvalEvent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalMTLLateEvalEvent.Device]
//   - [IOGPUMetalMTLLateEvalEvent.Label]
//   - [IOGPUMetalMTLLateEvalEvent.SetLabel]
//   - [IOGPUMetalMTLLateEvalEvent.NewSharedEventHandle]
//   - [IOGPUMetalMTLLateEvalEvent.NotifyListenerAtValueBlock]
//   - [IOGPUMetalMTLLateEvalEvent.RetainedLabel]
//   - [IOGPUMetalMTLLateEvalEvent.WaitUntilSignaledValueTimeoutMS]
//   - [IOGPUMetalMTLLateEvalEvent.DebugDescription]
//   - [IOGPUMetalMTLLateEvalEvent.Description]
//   - [IOGPUMetalMTLLateEvalEvent.Hash]
//   - [IOGPUMetalMTLLateEvalEvent.Superclass]
type IOGPUMetalMTLLateEvalEvent struct {
	IOGPUMTLLateEvalEvent
}

// IOGPUMetalMTLLateEvalEventFromID constructs a [IOGPUMetalMTLLateEvalEvent] from an objc.ID.
func IOGPUMetalMTLLateEvalEventFromID(id objc.ID) IOGPUMetalMTLLateEvalEvent {
	return IOGPUMetalMTLLateEvalEvent{IOGPUMTLLateEvalEvent: IOGPUMTLLateEvalEventFromID(id)}
}

// Ensure IOGPUMetalMTLLateEvalEvent implements IIOGPUMetalMTLLateEvalEvent.
var _ IIOGPUMetalMTLLateEvalEvent = IOGPUMetalMTLLateEvalEvent{}

// An interface definition for the [IOGPUMetalMTLLateEvalEvent] class.
//
// # Methods
//
//   - [IIOGPUMetalMTLLateEvalEvent.Device]
//   - [IIOGPUMetalMTLLateEvalEvent.Label]
//   - [IIOGPUMetalMTLLateEvalEvent.SetLabel]
//   - [IIOGPUMetalMTLLateEvalEvent.NewSharedEventHandle]
//   - [IIOGPUMetalMTLLateEvalEvent.NotifyListenerAtValueBlock]
//   - [IIOGPUMetalMTLLateEvalEvent.RetainedLabel]
//   - [IIOGPUMetalMTLLateEvalEvent.WaitUntilSignaledValueTimeoutMS]
//   - [IIOGPUMetalMTLLateEvalEvent.DebugDescription]
//   - [IIOGPUMetalMTLLateEvalEvent.Description]
//   - [IIOGPUMetalMTLLateEvalEvent.Hash]
//   - [IIOGPUMetalMTLLateEvalEvent.Superclass]
type IIOGPUMetalMTLLateEvalEvent interface {
	IIOGPUMTLLateEvalEvent

	// Topic: Methods

	Device() unsafe.Pointer
	Label() string
	SetLabel(value string)
	NewSharedEventHandle() objectivec.IObject
	NotifyListenerAtValueBlock(listener objectivec.IObject, value uint64, block VoidHandler)
	RetainedLabel() objectivec.IObject
	WaitUntilSignaledValueTimeoutMS(value uint64, ms uint64) bool
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (i IOGPUMetalMTLLateEvalEvent) Init() IOGPUMetalMTLLateEvalEvent {
	rv := objc.SendIfResponds[IOGPUMetalMTLLateEvalEvent](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalMTLLateEvalEvent) Autorelease() IOGPUMetalMTLLateEvalEvent {
	rv := objc.SendIfResponds[IOGPUMetalMTLLateEvalEvent](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalMTLLateEvalEvent creates a new IOGPUMetalMTLLateEvalEvent instance.
func NewIOGPUMetalMTLLateEvalEvent() IOGPUMetalMTLLateEvalEvent {
	class := getIOGPUMetalMTLLateEvalEventClass()
	rv := objc.SendIfResponds[IOGPUMetalMTLLateEvalEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewIOGPUMetalMTLLateEvalEventWithDevice(device objectivec.IObject) IOGPUMetalMTLLateEvalEvent {
	instance := getIOGPUMetalMTLLateEvalEventClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return IOGPUMetalMTLLateEvalEventFromID(rv)
}

func (i IOGPUMetalMTLLateEvalEvent) NewSharedEventHandle() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("newSharedEventHandle"))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalMTLLateEvalEvent) NotifyListenerAtValueBlock(listener objectivec.IObject, value uint64, block VoidHandler) {
	_block2, _ := NewVoidBlock(block)
	objc.SendIfResponds[objc.ID](i.ID, objc.Sel("notifyListener:atValue:block:"), listener, value, _block2)
}
func (i IOGPUMetalMTLLateEvalEvent) RetainedLabel() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("retainedLabel"))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalMTLLateEvalEvent) WaitUntilSignaledValueTimeoutMS(value uint64, ms uint64) bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("waitUntilSignaledValue:timeoutMS:"), value, ms)
	return rv
}

func (i IOGPUMetalMTLLateEvalEvent) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (i IOGPUMetalMTLLateEvalEvent) Description() string {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (i IOGPUMetalMTLLateEvalEvent) Device() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("device"))
	return rv
}
func (i IOGPUMetalMTLLateEvalEvent) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("hash"))
	return rv
}
func (i IOGPUMetalMTLLateEvalEvent) Label() string {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (i IOGPUMetalMTLLateEvalEvent) SetLabel(value string) {
	objc.SendIfResponds[struct{}](i.ID, objc.Sel("setLabel:"), objc.String(value))
}
func (i IOGPUMetalMTLLateEvalEvent) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](i.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}

// NotifyListenerAtValueBlockSync is a synchronous wrapper around [IOGPUMetalMTLLateEvalEvent.NotifyListenerAtValueBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (i IOGPUMetalMTLLateEvalEvent) NotifyListenerAtValueBlockSync(ctx context.Context, listener objectivec.IObject, value uint64) error {
	done := make(chan struct{}, 1)
	i.NotifyListenerAtValueBlock(listener, value, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
