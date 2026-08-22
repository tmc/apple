// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMetalMTLEvent] class.
var (
	_IOGPUMetalMTLEventClass     IOGPUMetalMTLEventClass
	_IOGPUMetalMTLEventClassOnce sync.Once
)

func getIOGPUMetalMTLEventClass() IOGPUMetalMTLEventClass {
	_IOGPUMetalMTLEventClassOnce.Do(func() {
		_IOGPUMetalMTLEventClass = IOGPUMetalMTLEventClass{class: objc.GetClass("_IOGPUMetalMTLEvent")}
	})
	return _IOGPUMetalMTLEventClass
}

// GetIOGPUMetalMTLEventClass returns the class object for _IOGPUMetalMTLEvent.
func GetIOGPUMetalMTLEventClass() IOGPUMetalMTLEventClass {
	return getIOGPUMetalMTLEventClass()
}

type IOGPUMetalMTLEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMetalMTLEventClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMetalMTLEventClass) Alloc() IOGPUMetalMTLEvent {
	rv := objc.SendIfResponds[IOGPUMetalMTLEvent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMetalMTLEvent.Device]
//   - [IOGPUMetalMTLEvent.Label]
//   - [IOGPUMetalMTLEvent.SetLabel]
//   - [IOGPUMetalMTLEvent.RetainedLabel]
//   - [IOGPUMetalMTLEvent.InitWithDevice]
//   - [IOGPUMetalMTLEvent.InitWithDeviceOptions]
//   - [IOGPUMetalMTLEvent.DebugDescription]
//   - [IOGPUMetalMTLEvent.Description]
//   - [IOGPUMetalMTLEvent.Hash]
//   - [IOGPUMetalMTLEvent.Superclass]
type IOGPUMetalMTLEvent struct {
	IOGPUMTLEvent
}

// IOGPUMetalMTLEventFromID constructs a [IOGPUMetalMTLEvent] from an objc.ID.
func IOGPUMetalMTLEventFromID(id objc.ID) IOGPUMetalMTLEvent {
	return IOGPUMetalMTLEvent{IOGPUMTLEvent: IOGPUMTLEventFromID(id)}
}

// Ensure IOGPUMetalMTLEvent implements IIOGPUMetalMTLEvent.
var _ IIOGPUMetalMTLEvent = IOGPUMetalMTLEvent{}

// An interface definition for the [IOGPUMetalMTLEvent] class.
//
// # Methods
//
//   - [IIOGPUMetalMTLEvent.Device]
//   - [IIOGPUMetalMTLEvent.Label]
//   - [IIOGPUMetalMTLEvent.SetLabel]
//   - [IIOGPUMetalMTLEvent.RetainedLabel]
//   - [IIOGPUMetalMTLEvent.InitWithDevice]
//   - [IIOGPUMetalMTLEvent.InitWithDeviceOptions]
//   - [IIOGPUMetalMTLEvent.DebugDescription]
//   - [IIOGPUMetalMTLEvent.Description]
//   - [IIOGPUMetalMTLEvent.Hash]
//   - [IIOGPUMetalMTLEvent.Superclass]
type IIOGPUMetalMTLEvent interface {
	IIOGPUMTLEvent

	// Topic: Methods

	Device() unsafe.Pointer
	Label() string
	SetLabel(value string)
	RetainedLabel() objectivec.IObject
	InitWithDevice(device objectivec.IObject) IOGPUMetalMTLEvent
	InitWithDeviceOptions(device objectivec.IObject, options uint64) IOGPUMetalMTLEvent
	DebugDescription() string
	Description() string
	Hash() uint64
	Superclass() objectivec.Class
}

// Init initializes the instance.
func (i IOGPUMetalMTLEvent) Init() IOGPUMetalMTLEvent {
	rv := objc.SendIfResponds[IOGPUMetalMTLEvent](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMetalMTLEvent) Autorelease() IOGPUMetalMTLEvent {
	rv := objc.SendIfResponds[IOGPUMetalMTLEvent](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMetalMTLEvent creates a new IOGPUMetalMTLEvent instance.
func NewIOGPUMetalMTLEvent() IOGPUMetalMTLEvent {
	class := getIOGPUMetalMTLEventClass()
	rv := objc.SendIfResponds[IOGPUMetalMTLEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewIOGPUMetalMTLEventWithDevice(device objectivec.IObject) IOGPUMetalMTLEvent {
	instance := getIOGPUMetalMTLEventClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:"), device)
	return IOGPUMetalMTLEventFromID(rv)
}

func NewIOGPUMetalMTLEventWithDeviceOptions(device objectivec.IObject, options uint64) IOGPUMetalMTLEvent {
	instance := getIOGPUMetalMTLEventClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDevice:options:"), device, options)
	return IOGPUMetalMTLEventFromID(rv)
}

func NewIOGPUMetalMTLEventWithDeviceRef(ref *uintptr) IOGPUMetalMTLEvent {
	instance := getIOGPUMetalMTLEventClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDeviceRef:"), ref)
	return IOGPUMetalMTLEventFromID(rv)
}

func NewIOGPUMetalMTLEventWithDeviceRefOptions(ref *uintptr, options uint64) IOGPUMetalMTLEvent {
	instance := getIOGPUMetalMTLEventClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDeviceRef:options:"), ref, options)
	return IOGPUMetalMTLEventFromID(rv)
}

func (i IOGPUMetalMTLEvent) RetainedLabel() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("retainedLabel"))
	return objectivec.Object{ID: rv}
}
func (i IOGPUMetalMTLEvent) InitWithDevice(device objectivec.IObject) IOGPUMetalMTLEvent {
	rv := objc.SendIfResponds[IOGPUMetalMTLEvent](i.ID, objc.Sel("initWithDevice:"), device)
	return rv
}
func (i IOGPUMetalMTLEvent) InitWithDeviceOptions(device objectivec.IObject, options uint64) IOGPUMetalMTLEvent {
	rv := objc.SendIfResponds[IOGPUMetalMTLEvent](i.ID, objc.Sel("initWithDevice:options:"), device, options)
	return rv
}

func (i IOGPUMetalMTLEvent) DebugDescription() string {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("debugDescription"))
	return foundation.NSStringFromID(rv).String()
}
func (i IOGPUMetalMTLEvent) Description() string {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("description"))
	return foundation.NSStringFromID(rv).String()
}
func (i IOGPUMetalMTLEvent) Device() unsafe.Pointer {
	rv := objc.SendIfResponds[unsafe.Pointer](i.ID, objc.Sel("device"))
	return rv
}
func (i IOGPUMetalMTLEvent) Hash() uint64 {
	rv := objc.SendIfResponds[uint64](i.ID, objc.Sel("hash"))
	return rv
}
func (i IOGPUMetalMTLEvent) Label() string {
	rv := objc.SendIfResponds[objc.ID](i.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}
func (i IOGPUMetalMTLEvent) SetLabel(value string) {
	objc.SendIfResponds[struct{}](i.ID, objc.Sel("setLabel:"), objc.String(value))
}
func (i IOGPUMetalMTLEvent) Superclass() objectivec.Class {
	rv := objc.SendIfResponds[objectivec.Class](i.ID, objc.Sel("superclass"))
	return objectivec.Class(rv)
}
