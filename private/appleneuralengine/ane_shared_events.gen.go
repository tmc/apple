// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

package appleneuralengine

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [ANESharedEvents] class.
var (
	_ANESharedEventsClass     ANESharedEventsClass
	_ANESharedEventsClassOnce sync.Once
)

func getANESharedEventsClass() ANESharedEventsClass {
	_ANESharedEventsClassOnce.Do(func() {
		_ANESharedEventsClass = ANESharedEventsClass{class: objc.GetClass("_ANESharedEvents")}
	})
	return _ANESharedEventsClass
}

// GetANESharedEventsClass returns the class object for _ANESharedEvents.
func GetANESharedEventsClass() ANESharedEventsClass {
	return getANESharedEventsClass()
}

type ANESharedEventsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac ANESharedEventsClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac ANESharedEventsClass) Alloc() ANESharedEvents {
	rv := objc.Send[ANESharedEvents](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [ANESharedEvents.SignalEvents]
//   - [ANESharedEvents.SetSignalEvents]
//   - [ANESharedEvents.WaitEvents]
//   - [ANESharedEvents.SetWaitEvents]
//   - [ANESharedEvents.InitWithSignalEventsWaitEvents]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedEvents
type ANESharedEvents struct {
	objectivec.Object
}

// ANESharedEventsFromID constructs a [ANESharedEvents] from an objc.ID.
func ANESharedEventsFromID(id objc.ID) ANESharedEvents {
	return ANESharedEvents{objectivec.Object{ID: id}}
}

// Ensure ANESharedEvents implements IANESharedEvents.
var _ IANESharedEvents = ANESharedEvents{}

// An interface definition for the [ANESharedEvents] class.
//
// # Methods
//
//   - [IANESharedEvents.SignalEvents]
//   - [IANESharedEvents.SetSignalEvents]
//   - [IANESharedEvents.WaitEvents]
//   - [IANESharedEvents.SetWaitEvents]
//   - [IANESharedEvents.InitWithSignalEventsWaitEvents]
//
// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedEvents
type IANESharedEvents interface {
	objectivec.IObject

	// Topic: Methods

	SignalEvents() foundation.INSArray
	SetSignalEvents(value foundation.INSArray)
	WaitEvents() foundation.INSArray
	SetWaitEvents(value foundation.INSArray)
	InitWithSignalEventsWaitEvents(events objectivec.IObject, events2 objectivec.IObject) ANESharedEvents
}

// Init initializes the instance.
func (a ANESharedEvents) Init() ANESharedEvents {
	rv := objc.Send[ANESharedEvents](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a ANESharedEvents) Autorelease() ANESharedEvents {
	rv := objc.Send[ANESharedEvents](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewANESharedEvents creates a new ANESharedEvents instance.
func NewANESharedEvents() ANESharedEvents {
	class := getANESharedEventsClass()
	rv := objc.Send[ANESharedEvents](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedEvents/initWithSignalEvents:waitEvents:
func NewANESharedEventsWithSignalEventsWaitEvents(events objectivec.IObject, events2 objectivec.IObject) ANESharedEvents {
	instance := getANESharedEventsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithSignalEvents:waitEvents:"), events, events2)
	return ANESharedEventsFromID(rv)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedEvents/initWithSignalEvents:waitEvents:
func (a ANESharedEvents) InitWithSignalEventsWaitEvents(events objectivec.IObject, events2 objectivec.IObject) ANESharedEvents {
	rv := objc.Send[ANESharedEvents](a.ID, objc.Sel("initWithSignalEvents:waitEvents:"), events, events2)
	return rv
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedEvents/sharedEventsWithSignalEvents:waitEvents:
func (_ANESharedEventsClass ANESharedEventsClass) SharedEventsWithSignalEventsWaitEvents(events objectivec.IObject, events2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](objc.ID(_ANESharedEventsClass.class), objc.Sel("sharedEventsWithSignalEvents:waitEvents:"), events, events2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedEvents/signalEvents
func (a ANESharedEvents) SignalEvents() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("signalEvents"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (a ANESharedEvents) SetSignalEvents(value foundation.INSArray) {
	objc.Send[struct{}](a.ID, objc.Sel("setSignalEvents:"), value)
}

// See: https://developer.apple.com/documentation/AppleNeuralEngine/_ANESharedEvents/waitEvents
func (a ANESharedEvents) WaitEvents() foundation.INSArray {
	rv := objc.Send[objc.ID](a.ID, objc.Sel("waitEvents"))
	return foundation.NSArrayFromID(objc.ID(rv))
}
func (a ANESharedEvents) SetWaitEvents(value foundation.INSArray) {
	objc.Send[struct{}](a.ID, objc.Sel("setWaitEvents:"), value)
}
