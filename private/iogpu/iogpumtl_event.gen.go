// Code generated from Apple documentation for iogpu. DO NOT EDIT.

package iogpu

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [IOGPUMTLEvent] class.
var (
	_IOGPUMTLEventClass     IOGPUMTLEventClass
	_IOGPUMTLEventClassOnce sync.Once
)

func getIOGPUMTLEventClass() IOGPUMTLEventClass {
	_IOGPUMTLEventClassOnce.Do(func() {
		_IOGPUMTLEventClass = IOGPUMTLEventClass{class: objc.GetClass("IOGPUMTLEvent")}
	})
	return _IOGPUMTLEventClass
}

// GetIOGPUMTLEventClass returns the class object for IOGPUMTLEvent.
func GetIOGPUMTLEventClass() IOGPUMTLEventClass {
	return getIOGPUMTLEventClass()
}

type IOGPUMTLEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ic IOGPUMTLEventClass) Class() objc.Class {
	return ic.class
}

// Alloc allocates memory for a new instance of the class.
func (ic IOGPUMTLEventClass) Alloc() IOGPUMTLEvent {
	rv := objc.SendIfResponds[IOGPUMTLEvent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [IOGPUMTLEvent._encodeIOGPUKernelConditionalEventAbortCommandArgs]
//   - [IOGPUMTLEvent._encodeIOGPUKernelSignalEventAgentCommandArgsValueAgentMask]
//   - [IOGPUMTLEvent._encodeIOGPUKernelSignalEventCommandArgsValue]
//   - [IOGPUMTLEvent._encodeIOGPUKernelSignalEventScheduledCommandArgs]
//   - [IOGPUMTLEvent._encodeIOGPUKernelWaitEventCommandArgsValueTimeout]
//   - [IOGPUMTLEvent._isSharedEvent]
//   - [IOGPUMTLEvent.EnableBarrier]
//   - [IOGPUMTLEvent.SetEnableBarrier]
//   - [IOGPUMTLEvent.EventName]
//   - [IOGPUMTLEvent.SupportsRollback]
//   - [IOGPUMTLEvent.InitWithDeviceRef]
//   - [IOGPUMTLEvent.InitWithDeviceRefOptions]
type IOGPUMTLEvent struct {
	objectivec.Object
}

// IOGPUMTLEventFromID constructs a [IOGPUMTLEvent] from an objc.ID.
func IOGPUMTLEventFromID(id objc.ID) IOGPUMTLEvent {
	return IOGPUMTLEvent{objectivec.Object{ID: id}}
}

// Ensure IOGPUMTLEvent implements IIOGPUMTLEvent.
var _ IIOGPUMTLEvent = IOGPUMTLEvent{}

// An interface definition for the [IOGPUMTLEvent] class.
//
// # Methods
//
//   - [IIOGPUMTLEvent._encodeIOGPUKernelConditionalEventAbortCommandArgs]
//   - [IIOGPUMTLEvent._encodeIOGPUKernelSignalEventAgentCommandArgsValueAgentMask]
//   - [IIOGPUMTLEvent._encodeIOGPUKernelSignalEventCommandArgsValue]
//   - [IIOGPUMTLEvent._encodeIOGPUKernelSignalEventScheduledCommandArgs]
//   - [IIOGPUMTLEvent._encodeIOGPUKernelWaitEventCommandArgsValueTimeout]
//   - [IIOGPUMTLEvent._isSharedEvent]
//   - [IIOGPUMTLEvent.EnableBarrier]
//   - [IIOGPUMTLEvent.SetEnableBarrier]
//   - [IIOGPUMTLEvent.EventName]
//   - [IIOGPUMTLEvent.SupportsRollback]
//   - [IIOGPUMTLEvent.InitWithDeviceRef]
//   - [IIOGPUMTLEvent.InitWithDeviceRefOptions]
type IIOGPUMTLEvent interface {
	objectivec.IObject

	// Topic: Methods

	_encodeIOGPUKernelConditionalEventAbortCommandArgs(args *IOGPUKernelCommandSignalOrWaitEventArgs) uint32
	_encodeIOGPUKernelSignalEventAgentCommandArgsValueAgentMask(args *IOGPUKernelCommandSignalEventAgentArgs, value uint64, mask uint64) uint32
	_encodeIOGPUKernelSignalEventCommandArgsValue(args *IOGPUKernelCommandSignalOrWaitEventArgs, value uint64) uint32
	_encodeIOGPUKernelSignalEventScheduledCommandArgs(args *IOGPUKernelCommandSignalOrWaitEventArgs) uint32
	_encodeIOGPUKernelWaitEventCommandArgsValueTimeout(args *IOGPUKernelCommandSignalOrWaitEventArgs, value uint64, timeout uint32) uint32
	_isSharedEvent() bool
	EnableBarrier() bool
	SetEnableBarrier(value bool)
	EventName() uint32
	SupportsRollback() bool
	InitWithDeviceRef(ref *uintptr) IOGPUMTLEvent
	InitWithDeviceRefOptions(ref *uintptr, options uint64) IOGPUMTLEvent
}

// Init initializes the instance.
func (i IOGPUMTLEvent) Init() IOGPUMTLEvent {
	rv := objc.SendIfResponds[IOGPUMTLEvent](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i IOGPUMTLEvent) Autorelease() IOGPUMTLEvent {
	rv := objc.SendIfResponds[IOGPUMTLEvent](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewIOGPUMTLEvent creates a new IOGPUMTLEvent instance.
func NewIOGPUMTLEvent() IOGPUMTLEvent {
	class := getIOGPUMTLEventClass()
	rv := objc.SendIfResponds[IOGPUMTLEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func NewGPUMTLEventWithDeviceRef(ref *uintptr) IOGPUMTLEvent {
	instance := getIOGPUMTLEventClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDeviceRef:"), ref)
	return IOGPUMTLEventFromID(rv)
}

func NewGPUMTLEventWithDeviceRefOptions(ref *uintptr, options uint64) IOGPUMTLEvent {
	instance := getIOGPUMTLEventClass().Alloc()
	rv := objc.SendIfResponds[objc.ID](instance.ID, objc.Sel("initWithDeviceRef:options:"), ref, options)
	return IOGPUMTLEventFromID(rv)
}

func (i IOGPUMTLEvent) _encodeIOGPUKernelConditionalEventAbortCommandArgs(args *IOGPUKernelCommandSignalOrWaitEventArgs) uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("_encodeIOGPUKernelConditionalEventAbortCommandArgs:"), unsafe.Pointer(args))
	return rv
}

// EncodeIOGPUKernelConditionalEventAbortCommandArgs is an exported wrapper for the private method _encodeIOGPUKernelConditionalEventAbortCommandArgs.
func (i IOGPUMTLEvent) EncodeIOGPUKernelConditionalEventAbortCommandArgs(args *IOGPUKernelCommandSignalOrWaitEventArgs) (uint32, error) {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_encodeIOGPUKernelConditionalEventAbortCommandArgs:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_encodeIOGPUKernelConditionalEventAbortCommandArgs:"}
		return 0, err
	}
	return i._encodeIOGPUKernelConditionalEventAbortCommandArgs(args), nil
}

// CanEncodeIOGPUKernelConditionalEventAbortCommandArgs reports whether the receiver responds to the private selector _encodeIOGPUKernelConditionalEventAbortCommandArgs:.
func (i IOGPUMTLEvent) CanEncodeIOGPUKernelConditionalEventAbortCommandArgs() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_encodeIOGPUKernelConditionalEventAbortCommandArgs:"))
}
func (i IOGPUMTLEvent) _encodeIOGPUKernelSignalEventAgentCommandArgsValueAgentMask(args *IOGPUKernelCommandSignalEventAgentArgs, value uint64, mask uint64) uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("_encodeIOGPUKernelSignalEventAgentCommandArgs:value:agentMask:"), unsafe.Pointer(args), value, mask)
	return rv
}

// EncodeIOGPUKernelSignalEventAgentCommandArgsValueAgentMask is an exported wrapper for the private method _encodeIOGPUKernelSignalEventAgentCommandArgsValueAgentMask.
func (i IOGPUMTLEvent) EncodeIOGPUKernelSignalEventAgentCommandArgsValueAgentMask(args *IOGPUKernelCommandSignalEventAgentArgs, value uint64, mask uint64) (uint32, error) {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_encodeIOGPUKernelSignalEventAgentCommandArgs:value:agentMask:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_encodeIOGPUKernelSignalEventAgentCommandArgs:value:agentMask:"}
		return 0, err
	}
	return i._encodeIOGPUKernelSignalEventAgentCommandArgsValueAgentMask(args, value, mask), nil
}

// CanEncodeIOGPUKernelSignalEventAgentCommandArgsValueAgentMask reports whether the receiver responds to the private selector _encodeIOGPUKernelSignalEventAgentCommandArgs:value:agentMask:.
func (i IOGPUMTLEvent) CanEncodeIOGPUKernelSignalEventAgentCommandArgsValueAgentMask() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_encodeIOGPUKernelSignalEventAgentCommandArgs:value:agentMask:"))
}
func (i IOGPUMTLEvent) _encodeIOGPUKernelSignalEventCommandArgsValue(args *IOGPUKernelCommandSignalOrWaitEventArgs, value uint64) uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("_encodeIOGPUKernelSignalEventCommandArgs:value:"), unsafe.Pointer(args), value)
	return rv
}

// EncodeIOGPUKernelSignalEventCommandArgsValue is an exported wrapper for the private method _encodeIOGPUKernelSignalEventCommandArgsValue.
func (i IOGPUMTLEvent) EncodeIOGPUKernelSignalEventCommandArgsValue(args *IOGPUKernelCommandSignalOrWaitEventArgs, value uint64) (uint32, error) {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_encodeIOGPUKernelSignalEventCommandArgs:value:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_encodeIOGPUKernelSignalEventCommandArgs:value:"}
		return 0, err
	}
	return i._encodeIOGPUKernelSignalEventCommandArgsValue(args, value), nil
}

// CanEncodeIOGPUKernelSignalEventCommandArgsValue reports whether the receiver responds to the private selector _encodeIOGPUKernelSignalEventCommandArgs:value:.
func (i IOGPUMTLEvent) CanEncodeIOGPUKernelSignalEventCommandArgsValue() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_encodeIOGPUKernelSignalEventCommandArgs:value:"))
}
func (i IOGPUMTLEvent) _encodeIOGPUKernelSignalEventScheduledCommandArgs(args *IOGPUKernelCommandSignalOrWaitEventArgs) uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("_encodeIOGPUKernelSignalEventScheduledCommandArgs:"), unsafe.Pointer(args))
	return rv
}

// EncodeIOGPUKernelSignalEventScheduledCommandArgs is an exported wrapper for the private method _encodeIOGPUKernelSignalEventScheduledCommandArgs.
func (i IOGPUMTLEvent) EncodeIOGPUKernelSignalEventScheduledCommandArgs(args *IOGPUKernelCommandSignalOrWaitEventArgs) (uint32, error) {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_encodeIOGPUKernelSignalEventScheduledCommandArgs:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_encodeIOGPUKernelSignalEventScheduledCommandArgs:"}
		return 0, err
	}
	return i._encodeIOGPUKernelSignalEventScheduledCommandArgs(args), nil
}

// CanEncodeIOGPUKernelSignalEventScheduledCommandArgs reports whether the receiver responds to the private selector _encodeIOGPUKernelSignalEventScheduledCommandArgs:.
func (i IOGPUMTLEvent) CanEncodeIOGPUKernelSignalEventScheduledCommandArgs() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_encodeIOGPUKernelSignalEventScheduledCommandArgs:"))
}
func (i IOGPUMTLEvent) _encodeIOGPUKernelWaitEventCommandArgsValueTimeout(args *IOGPUKernelCommandSignalOrWaitEventArgs, value uint64, timeout uint32) uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("_encodeIOGPUKernelWaitEventCommandArgs:value:timeout:"), unsafe.Pointer(args), value, timeout)
	return rv
}

// EncodeIOGPUKernelWaitEventCommandArgsValueTimeout is an exported wrapper for the private method _encodeIOGPUKernelWaitEventCommandArgsValueTimeout.
func (i IOGPUMTLEvent) EncodeIOGPUKernelWaitEventCommandArgsValueTimeout(args *IOGPUKernelCommandSignalOrWaitEventArgs, value uint64, timeout uint32) (uint32, error) {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_encodeIOGPUKernelWaitEventCommandArgs:value:timeout:")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_encodeIOGPUKernelWaitEventCommandArgs:value:timeout:"}
		return 0, err
	}
	return i._encodeIOGPUKernelWaitEventCommandArgsValueTimeout(args, value, timeout), nil
}

// CanEncodeIOGPUKernelWaitEventCommandArgsValueTimeout reports whether the receiver responds to the private selector _encodeIOGPUKernelWaitEventCommandArgs:value:timeout:.
func (i IOGPUMTLEvent) CanEncodeIOGPUKernelWaitEventCommandArgsValueTimeout() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_encodeIOGPUKernelWaitEventCommandArgs:value:timeout:"))
}
func (i IOGPUMTLEvent) _isSharedEvent() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("_isSharedEvent"))
	return rv
}

// IsSharedEvent is an exported wrapper for the private method _isSharedEvent.
func (i IOGPUMTLEvent) IsSharedEvent() (bool, error) {
	if !objc.RespondsToSelector(i.ID, objc.Sel("_isSharedEvent")) {
		err := &objc.UnrecognizedSelectorError{Selector: "_isSharedEvent"}
		return false, err
	}
	return i._isSharedEvent(), nil
}

// CanIsSharedEvent reports whether the receiver responds to the private selector _isSharedEvent.
func (i IOGPUMTLEvent) CanIsSharedEvent() bool {
	return objc.RespondsToSelector(i.ID, objc.Sel("_isSharedEvent"))
}
func (i IOGPUMTLEvent) InitWithDeviceRef(ref *uintptr) IOGPUMTLEvent {
	rv := objc.SendIfResponds[IOGPUMTLEvent](i.ID, objc.Sel("initWithDeviceRef:"), ref)
	return rv
}
func (i IOGPUMTLEvent) InitWithDeviceRefOptions(ref *uintptr, options uint64) IOGPUMTLEvent {
	rv := objc.SendIfResponds[IOGPUMTLEvent](i.ID, objc.Sel("initWithDeviceRef:options:"), ref, options)
	return rv
}

func (i IOGPUMTLEvent) EnableBarrier() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("enableBarrier"))
	return rv
}
func (i IOGPUMTLEvent) SetEnableBarrier(value bool) {
	objc.SendIfResponds[struct{}](i.ID, objc.Sel("setEnableBarrier:"), value)
}
func (i IOGPUMTLEvent) EventName() uint32 {
	rv := objc.SendIfResponds[uint32](i.ID, objc.Sel("eventName"))
	return rv
}
func (i IOGPUMTLEvent) SupportsRollback() bool {
	rv := objc.SendIfResponds[bool](i.ID, objc.Sel("supportsRollback"))
	return rv
}
