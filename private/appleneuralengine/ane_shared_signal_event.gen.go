// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANESharedSignalEvent] class.
var (
	_ANESharedSignalEventClass     ANESharedSignalEventClass
	_ANESharedSignalEventClassOnce sync.Once
)

func getANESharedSignalEventClass() ANESharedSignalEventClass {
	_ANESharedSignalEventClassOnce.Do(func() {
		_ANESharedSignalEventClass = ANESharedSignalEventClass{class: objc.GetClass("_ANESharedSignalEvent")}
	})
	return _ANESharedSignalEventClass
}

// GetANESharedSignalEventClass returns the class object for _ANESharedSignalEvent.
func GetANESharedSignalEventClass() ANESharedSignalEventClass {
	return getANESharedSignalEventClass()
}

type ANESharedSignalEventClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANESharedSignalEventClass) Alloc() ANESharedSignalEvent {
	rv := objc.Send[ANESharedSignalEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ANESharedSignalEvent.AgentMask]
//   - [ANESharedSignalEvent.SetAgentMask]
//   - [ANESharedSignalEvent.EncodeWithCoder]
//   - [ANESharedSignalEvent.EventType]
//   - [ANESharedSignalEvent.SharedEvent]
//   - [ANESharedSignalEvent.SymbolIndex]
//   - [ANESharedSignalEvent.Value]
//   - [ANESharedSignalEvent.SetValue]
//   - [ANESharedSignalEvent.WaitEvent]
//   - [ANESharedSignalEvent.InitWithCoder]
//   - [ANESharedSignalEvent.InitWithValueSymbolIndexEventTypeSharedEventAgentMask]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedSignalEvent
type ANESharedSignalEvent struct {
	objectivec.Object
}

// ANESharedSignalEventFromID constructs a [ANESharedSignalEvent] from an objc.ID.
func ANESharedSignalEventFromID(id objc.ID) ANESharedSignalEvent {
	return ANESharedSignalEvent{objectivec.Object{ID: id}}
}
// Ensure ANESharedSignalEvent implements IANESharedSignalEvent.
var _ IANESharedSignalEvent = ANESharedSignalEvent{}

// An interface definition for the [ANESharedSignalEvent] class.
//
// # Methods
//
//   - [IANESharedSignalEvent.AgentMask]
//   - [IANESharedSignalEvent.SetAgentMask]
//   - [IANESharedSignalEvent.EncodeWithCoder]
//   - [IANESharedSignalEvent.EventType]
//   - [IANESharedSignalEvent.SharedEvent]
//   - [IANESharedSignalEvent.SymbolIndex]
//   - [IANESharedSignalEvent.Value]
//   - [IANESharedSignalEvent.SetValue]
//   - [IANESharedSignalEvent.WaitEvent]
//   - [IANESharedSignalEvent.InitWithCoder]
//   - [IANESharedSignalEvent.InitWithValueSymbolIndexEventTypeSharedEventAgentMask]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedSignalEvent
type IANESharedSignalEvent interface {
	objectivec.IObject

	// Topic: Methods

	AgentMask() uint64
	SetAgentMask(value uint64)
	EncodeWithCoder(coder foundation.INSCoder)
	EventType() int64
	SharedEvent() objectivec.IObject
	SymbolIndex() uint32
	Value() uint64
	SetValue(value uint64)
	WaitEvent() objectivec.IObject
	InitWithCoder(coder foundation.INSCoder) ANESharedSignalEvent
	InitWithValueSymbolIndexEventTypeSharedEventAgentMask(value uint64, index uint32, type_ int64, event objectivec.IObject, mask uint64) ANESharedSignalEvent
}

// Init initializes the instance.
func (a ANESharedSignalEvent) Init() ANESharedSignalEvent {
	rv := objc.Send[ANESharedSignalEvent](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANESharedSignalEvent) Autorelease() ANESharedSignalEvent {
	rv := objc.Send[ANESharedSignalEvent](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANESharedSignalEvent creates a new ANESharedSignalEvent instance.
func NewANESharedSignalEvent() ANESharedSignalEvent {
	class := getANESharedSignalEventClass()
	rv := objc.Send[ANESharedSignalEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedSignalEvent/initWithCoder:
func NewANESharedSignalEventWithCoder(coder objectivec.IObject) ANESharedSignalEvent {
	instance := getANESharedSignalEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return ANESharedSignalEventFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedSignalEvent/initWithValue:symbolIndex:eventType:sharedEvent:agentMask:
func NewANESharedSignalEventWithValueSymbolIndexEventTypeSharedEventAgentMask(value uint64, index uint32, type_ int64, event objectivec.IObject, mask uint64) ANESharedSignalEvent {
	instance := getANESharedSignalEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValue:symbolIndex:eventType:sharedEvent:agentMask:"), value, index, type_, event, mask)
	return ANESharedSignalEventFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedSignalEvent/encodeWithCoder:
func (a ANESharedSignalEvent) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](a.ID, objc.Sel("encodeWithCoder:"), coder)
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedSignalEvent/waitEvent
func (a ANESharedSignalEvent) WaitEvent() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("waitEvent"))
	return objectivec.Object{ID: rv}
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedSignalEvent/initWithCoder:
func (a ANESharedSignalEvent) InitWithCoder(coder foundation.INSCoder) ANESharedSignalEvent {
	rv := objc.Send[ANESharedSignalEvent](a.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedSignalEvent/initWithValue:symbolIndex:eventType:sharedEvent:agentMask:
func (a ANESharedSignalEvent) InitWithValueSymbolIndexEventTypeSharedEventAgentMask(value uint64, index uint32, type_ int64, event objectivec.IObject, mask uint64) ANESharedSignalEvent {
	rv := objc.Send[ANESharedSignalEvent](a.ID, objc.Sel("initWithValue:symbolIndex:eventType:sharedEvent:agentMask:"), value, index, type_, event, mask)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedSignalEvent/signalEventWithValue:symbolIndex:eventType:sharedEvent:
func (_ANESharedSignalEventClass ANESharedSignalEventClass) SignalEventWithValueSymbolIndexEventTypeSharedEvent(value uint64, index uint32, type_ int64, event objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANESharedSignalEventClass.class), objc.Sel("signalEventWithValue:symbolIndex:eventType:sharedEvent:"), value, index, type_, event)
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedSignalEvent/supportsSecureCoding
func (_ANESharedSignalEventClass ANESharedSignalEventClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(_ANESharedSignalEventClass.class), objc.Sel("supportsSecureCoding"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedSignalEvent/agentMask
func (a ANESharedSignalEvent) AgentMask() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("agentMask"))
	return rv
}
func (a ANESharedSignalEvent) SetAgentMask(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setAgentMask:"), value)
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedSignalEvent/eventType
func (a ANESharedSignalEvent) EventType() int64 {
	rv := objc.Send[int64](a.ID, objc.Sel("eventType"))
	return rv
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedSignalEvent/sharedEvent
func (a ANESharedSignalEvent) SharedEvent() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sharedEvent"))
	return objectivec.Object{ID: rv}
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedSignalEvent/symbolIndex
func (a ANESharedSignalEvent) SymbolIndex() uint32 {
	rv := objc.Send[uint32](a.ID, objc.Sel("symbolIndex"))
	return rv
}
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedSignalEvent/value
func (a ANESharedSignalEvent) Value() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("value"))
	return rv
}
func (a ANESharedSignalEvent) SetValue(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setValue:"), value)
}

