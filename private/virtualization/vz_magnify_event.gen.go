// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMagnifyEvent] class.
var (
	_VZMagnifyEventClass     VZMagnifyEventClass
	_VZMagnifyEventClassOnce sync.Once
)

func getVZMagnifyEventClass() VZMagnifyEventClass {
	_VZMagnifyEventClassOnce.Do(func() {
		_VZMagnifyEventClass = VZMagnifyEventClass{class: objc.GetClass("_VZMagnifyEvent")}
	})
	return _VZMagnifyEventClass
}

// GetVZMagnifyEventClass returns the class object for _VZMagnifyEvent.
func GetVZMagnifyEventClass() VZMagnifyEventClass {
	return getVZMagnifyEventClass()
}

type VZMagnifyEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMagnifyEventClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMagnifyEventClass) Alloc() VZMagnifyEvent {
	rv := objc.Send[VZMagnifyEvent](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMagnifyEvent.Magnification]
//   - [VZMagnifyEvent.Phase]
//   - [VZMagnifyEvent.InitWithEvent]
//   - [VZMagnifyEvent.InitWithMagnificationPhase]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMagnifyEvent
type VZMagnifyEvent struct {
	objectivec.Object
}

// VZMagnifyEventFromID constructs a [VZMagnifyEvent] from an objc.ID.
func VZMagnifyEventFromID(id objc.ID) VZMagnifyEvent {
	return VZMagnifyEvent{objectivec.Object{ID: id}}
}

// Ensure VZMagnifyEvent implements IVZMagnifyEvent.
var _ IVZMagnifyEvent = VZMagnifyEvent{}

// An interface definition for the [VZMagnifyEvent] class.
//
// # Methods
//
//   - [IVZMagnifyEvent.Magnification]
//   - [IVZMagnifyEvent.Phase]
//   - [IVZMagnifyEvent.InitWithEvent]
//   - [IVZMagnifyEvent.InitWithMagnificationPhase]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMagnifyEvent
type IVZMagnifyEvent interface {
	objectivec.IObject

	// Topic: Methods

	Magnification() float64
	Phase() uint64
	InitWithEvent(event objectivec.IObject) VZMagnifyEvent
	InitWithMagnificationPhase(magnification float64, phase uint64) VZMagnifyEvent
}

// Init initializes the instance.
func (v VZMagnifyEvent) Init() VZMagnifyEvent {
	rv := objc.Send[VZMagnifyEvent](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMagnifyEvent) Autorelease() VZMagnifyEvent {
	rv := objc.Send[VZMagnifyEvent](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMagnifyEvent creates a new VZMagnifyEvent instance.
func NewVZMagnifyEvent() VZMagnifyEvent {
	class := getVZMagnifyEventClass()
	rv := objc.Send[VZMagnifyEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMagnifyEvent/initWithEvent:
func NewVZMagnifyEventWithEvent(event objectivec.IObject) VZMagnifyEvent {
	instance := getVZMagnifyEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEvent:"), event)
	return VZMagnifyEventFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMagnifyEvent/initWithMagnification:phase:
func NewVZMagnifyEventWithMagnificationPhase(magnification float64, phase uint64) VZMagnifyEvent {
	instance := getVZMagnifyEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMagnification:phase:"), magnification, phase)
	return VZMagnifyEventFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMagnifyEvent/initWithEvent:
func (v VZMagnifyEvent) InitWithEvent(event objectivec.IObject) VZMagnifyEvent {
	rv := objc.Send[VZMagnifyEvent](v.ID, objc.Sel("initWithEvent:"), event)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMagnifyEvent/initWithMagnification:phase:
func (v VZMagnifyEvent) InitWithMagnificationPhase(magnification float64, phase uint64) VZMagnifyEvent {
	rv := objc.Send[VZMagnifyEvent](v.ID, objc.Sel("initWithMagnification:phase:"), magnification, phase)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMagnifyEvent/magnification
func (v VZMagnifyEvent) Magnification() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("magnification"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMagnifyEvent/phase
func (v VZMagnifyEvent) Phase() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("phase"))
	return rv
}
