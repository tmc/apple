// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZIOHIDEvent] class.
var (
	_VZIOHIDEventClass     VZIOHIDEventClass
	_VZIOHIDEventClassOnce sync.Once
)

func getVZIOHIDEventClass() VZIOHIDEventClass {
	_VZIOHIDEventClassOnce.Do(func() {
		_VZIOHIDEventClass = VZIOHIDEventClass{class: objc.GetClass("_VZIOHIDEvent")}
	})
	return _VZIOHIDEventClass
}

// GetVZIOHIDEventClass returns the class object for _VZIOHIDEvent.
func GetVZIOHIDEventClass() VZIOHIDEventClass {
	return getVZIOHIDEventClass()
}

type VZIOHIDEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZIOHIDEventClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZIOHIDEventClass) Alloc() VZIOHIDEvent {
	rv := objc.Send[VZIOHIDEvent](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZIOHIDEvent.Event]
//   - [VZIOHIDEvent.InitWithIOHIDEvent]
// See: https://developer.apple.com/documentation/Virtualization/_VZIOHIDEvent
type VZIOHIDEvent struct {
	objectivec.Object
}

// VZIOHIDEventFromID constructs a [VZIOHIDEvent] from an objc.ID.
func VZIOHIDEventFromID(id objc.ID) VZIOHIDEvent {
	return VZIOHIDEvent{objectivec.Object{ID: id}}
}
// Ensure VZIOHIDEvent implements IVZIOHIDEvent.
var _ IVZIOHIDEvent = VZIOHIDEvent{}

// An interface definition for the [VZIOHIDEvent] class.
//
// # Methods
//
//   - [IVZIOHIDEvent.Event]
//   - [IVZIOHIDEvent.InitWithIOHIDEvent]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZIOHIDEvent
type IVZIOHIDEvent interface {
	objectivec.IObject

	// Topic: Methods

	Event() objectivec.IObject
	InitWithIOHIDEvent(iOHIDEvent objectivec.IObject) VZIOHIDEvent
}

// Init initializes the instance.
func (v VZIOHIDEvent) Init() VZIOHIDEvent {
	rv := objc.Send[VZIOHIDEvent](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZIOHIDEvent) Autorelease() VZIOHIDEvent {
	rv := objc.Send[VZIOHIDEvent](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZIOHIDEvent creates a new VZIOHIDEvent instance.
func NewVZIOHIDEvent() VZIOHIDEvent {
	class := getVZIOHIDEventClass()
	rv := objc.Send[VZIOHIDEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZIOHIDEvent/initWithIOHIDEvent:
func NewVZIOHIDEventWithIOHIDEvent(iOHIDEvent objectivec.IObject) VZIOHIDEvent {
	instance := getVZIOHIDEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithIOHIDEvent:"), iOHIDEvent)
	return VZIOHIDEventFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZIOHIDEvent/initWithIOHIDEvent:
func (v VZIOHIDEvent) InitWithIOHIDEvent(iOHIDEvent objectivec.IObject) VZIOHIDEvent {
	rv := objc.Send[VZIOHIDEvent](v.ID, objc.Sel("initWithIOHIDEvent:"), iOHIDEvent)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZIOHIDEvent/event
func (v VZIOHIDEvent) Event() objectivec.IObject {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("event"))
	return objectivec.Object{ID: rv}
}

