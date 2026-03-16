// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANESharedWaitEvent] class.
var (
	_ANESharedWaitEventClass     ANESharedWaitEventClass
	_ANESharedWaitEventClassOnce sync.Once
)

func getANESharedWaitEventClass() ANESharedWaitEventClass {
	_ANESharedWaitEventClassOnce.Do(func() {
		_ANESharedWaitEventClass = ANESharedWaitEventClass{class: objc.GetClass("_ANESharedWaitEvent")}
	})
	return _ANESharedWaitEventClass
}

// GetANESharedWaitEventClass returns the class object for _ANESharedWaitEvent.
func GetANESharedWaitEventClass() ANESharedWaitEventClass {
	return getANESharedWaitEventClass()
}

type ANESharedWaitEventClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANESharedWaitEventClass) Alloc() ANESharedWaitEvent {
	rv := objc.Send[ANESharedWaitEvent](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [ANESharedWaitEvent.EventType]
//   - [ANESharedWaitEvent.SharedEvent]
//   - [ANESharedWaitEvent.Value]
//   - [ANESharedWaitEvent.SetValue]
//   - [ANESharedWaitEvent.InitWithValueSharedEventEventType]
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedWaitEvent
type ANESharedWaitEvent struct {
	objectivec.Object
}

// ANESharedWaitEventFromID constructs a [ANESharedWaitEvent] from an objc.ID.
func ANESharedWaitEventFromID(id objc.ID) ANESharedWaitEvent {
	return ANESharedWaitEvent{objectivec.Object{id}}
}
// Ensure ANESharedWaitEvent implements IANESharedWaitEvent.
var _ IANESharedWaitEvent = ANESharedWaitEvent{}

// An interface definition for the [ANESharedWaitEvent] class.
//
// # Methods
//
//   - [IANESharedWaitEvent.EventType]
//   - [IANESharedWaitEvent.SharedEvent]
//   - [IANESharedWaitEvent.Value]
//   - [IANESharedWaitEvent.SetValue]
//   - [IANESharedWaitEvent.InitWithValueSharedEventEventType]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedWaitEvent
type IANESharedWaitEvent interface {
	objectivec.IObject

	// Topic: Methods

	EventType() uint64
	SharedEvent() objectivec.IObject
	Value() uint64
	SetValue(value uint64)
	InitWithValueSharedEventEventType(value uint64, event objectivec.IObject, type_ uint64) ANESharedWaitEvent
}

// Init initializes the instance.
func (a ANESharedWaitEvent) Init() ANESharedWaitEvent {
	rv := objc.Send[ANESharedWaitEvent](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANESharedWaitEvent) Autorelease() ANESharedWaitEvent {
	rv := objc.Send[ANESharedWaitEvent](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANESharedWaitEvent creates a new ANESharedWaitEvent instance.
func NewANESharedWaitEvent() ANESharedWaitEvent {
	class := getANESharedWaitEventClass()
	rv := objc.Send[ANESharedWaitEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedWaitEvent/initWithValue:sharedEvent:eventType:
func NewANESharedWaitEventWithValueSharedEventEventType(value uint64, event objectivec.IObject, type_ uint64) ANESharedWaitEvent {
	instance := getANESharedWaitEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValue:sharedEvent:eventType:"), value, event, type_)
	return ANESharedWaitEventFromID(rv)
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedWaitEvent/initWithValue:sharedEvent:eventType:
func (a ANESharedWaitEvent) InitWithValueSharedEventEventType(value uint64, event objectivec.IObject, type_ uint64) ANESharedWaitEvent {
	rv := objc.Send[ANESharedWaitEvent](a.ID, objc.Sel("initWithValue:sharedEvent:eventType:"), value, event, type_)
	return rv
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedWaitEvent/waitEventWithValue:sharedEvent:
func (_ANESharedWaitEventClass ANESharedWaitEventClass) WaitEventWithValueSharedEvent(value uint64, event objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANESharedWaitEventClass.class), objc.Sel("waitEventWithValue:sharedEvent:"), value, event)
	return objectivec.Object{ID: rv}
}

//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedWaitEvent/waitEventWithValue:sharedEvent:eventType:
func (_ANESharedWaitEventClass ANESharedWaitEventClass) WaitEventWithValueSharedEventEventType(value uint64, event objectivec.IObject, type_ uint64) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANESharedWaitEventClass.class), objc.Sel("waitEventWithValue:sharedEvent:eventType:"), value, event, type_)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedWaitEvent/eventType
func (a ANESharedWaitEvent) EventType() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("eventType"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedWaitEvent/sharedEvent
func (a ANESharedWaitEvent) SharedEvent() objectivec.IObject {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("sharedEvent"))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedWaitEvent/value
func (a ANESharedWaitEvent) Value() uint64 {
	rv := objc.Send[uint64](a.ID, objc.Sel("value"))
	return rv
}
func (a ANESharedWaitEvent) SetValue(value uint64) {
	objc.Send[struct{}](a.ID, objc.Sel("setValue:"), value)
}

