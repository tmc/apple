// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZSmartMagnifyEvent] class.
var (
	_VZSmartMagnifyEventClass     VZSmartMagnifyEventClass
	_VZSmartMagnifyEventClassOnce sync.Once
)

func getVZSmartMagnifyEventClass() VZSmartMagnifyEventClass {
	_VZSmartMagnifyEventClassOnce.Do(func() {
		_VZSmartMagnifyEventClass = VZSmartMagnifyEventClass{class: objc.GetClass("_VZSmartMagnifyEvent")}
	})
	return _VZSmartMagnifyEventClass
}

// GetVZSmartMagnifyEventClass returns the class object for _VZSmartMagnifyEvent.
func GetVZSmartMagnifyEventClass() VZSmartMagnifyEventClass {
	return getVZSmartMagnifyEventClass()
}

type VZSmartMagnifyEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSmartMagnifyEventClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSmartMagnifyEventClass) Alloc() VZSmartMagnifyEvent {
	rv := objc.Send[VZSmartMagnifyEvent](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZSmartMagnifyEvent.InitWithEvent]
// See: https://developer.apple.com/documentation/Virtualization/_VZSmartMagnifyEvent
type VZSmartMagnifyEvent struct {
	objectivec.Object
}

// VZSmartMagnifyEventFromID constructs a [VZSmartMagnifyEvent] from an objc.ID.
func VZSmartMagnifyEventFromID(id objc.ID) VZSmartMagnifyEvent {
	return VZSmartMagnifyEvent{objectivec.Object{ID: id}}
}
// Ensure VZSmartMagnifyEvent implements IVZSmartMagnifyEvent.
var _ IVZSmartMagnifyEvent = VZSmartMagnifyEvent{}

// An interface definition for the [VZSmartMagnifyEvent] class.
//
// # Methods
//
//   - [IVZSmartMagnifyEvent.InitWithEvent]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZSmartMagnifyEvent
type IVZSmartMagnifyEvent interface {
	objectivec.IObject

	// Topic: Methods

	InitWithEvent(event objectivec.IObject) VZSmartMagnifyEvent
}

// Init initializes the instance.
func (v VZSmartMagnifyEvent) Init() VZSmartMagnifyEvent {
	rv := objc.Send[VZSmartMagnifyEvent](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZSmartMagnifyEvent) Autorelease() VZSmartMagnifyEvent {
	rv := objc.Send[VZSmartMagnifyEvent](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSmartMagnifyEvent creates a new VZSmartMagnifyEvent instance.
func NewVZSmartMagnifyEvent() VZSmartMagnifyEvent {
	class := getVZSmartMagnifyEventClass()
	rv := objc.Send[VZSmartMagnifyEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZSmartMagnifyEvent/initWithEvent:
func NewVZSmartMagnifyEventWithEvent(event objectivec.IObject) VZSmartMagnifyEvent {
	instance := getVZSmartMagnifyEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEvent:"), event)
	return VZSmartMagnifyEventFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZSmartMagnifyEvent/initWithEvent:
func (v VZSmartMagnifyEvent) InitWithEvent(event objectivec.IObject) VZSmartMagnifyEvent {
	rv := objc.Send[VZSmartMagnifyEvent](v.ID, objc.Sel("initWithEvent:"), event)
	return rv
}

