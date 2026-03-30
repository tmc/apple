// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZRotationEvent] class.
var (
	_VZRotationEventClass     VZRotationEventClass
	_VZRotationEventClassOnce sync.Once
)

func getVZRotationEventClass() VZRotationEventClass {
	_VZRotationEventClassOnce.Do(func() {
		_VZRotationEventClass = VZRotationEventClass{class: objc.GetClass("_VZRotationEvent")}
	})
	return _VZRotationEventClass
}

// GetVZRotationEventClass returns the class object for _VZRotationEvent.
func GetVZRotationEventClass() VZRotationEventClass {
	return getVZRotationEventClass()
}

type VZRotationEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZRotationEventClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZRotationEventClass) Alloc() VZRotationEvent {
	rv := objc.Send[VZRotationEvent](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZRotationEvent.Phase]
//   - [VZRotationEvent.Rotation]
//   - [VZRotationEvent.InitWithEvent]
//   - [VZRotationEvent.InitWithRotationPhase]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZRotationEvent
type VZRotationEvent struct {
	objectivec.Object
}

// VZRotationEventFromID constructs a [VZRotationEvent] from an objc.ID.
func VZRotationEventFromID(id objc.ID) VZRotationEvent {
	return VZRotationEvent{objectivec.Object{ID: id}}
}

// Ensure VZRotationEvent implements IVZRotationEvent.
var _ IVZRotationEvent = VZRotationEvent{}

// An interface definition for the [VZRotationEvent] class.
//
// # Methods
//
//   - [IVZRotationEvent.Phase]
//   - [IVZRotationEvent.Rotation]
//   - [IVZRotationEvent.InitWithEvent]
//   - [IVZRotationEvent.InitWithRotationPhase]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZRotationEvent
type IVZRotationEvent interface {
	objectivec.IObject

	// Topic: Methods

	Phase() uint64
	Rotation() float64
	InitWithEvent(event objectivec.IObject) VZRotationEvent
	InitWithRotationPhase(rotation float64, phase uint64) VZRotationEvent
}

// Init initializes the instance.
func (v VZRotationEvent) Init() VZRotationEvent {
	rv := objc.Send[VZRotationEvent](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZRotationEvent) Autorelease() VZRotationEvent {
	rv := objc.Send[VZRotationEvent](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZRotationEvent creates a new VZRotationEvent instance.
func NewVZRotationEvent() VZRotationEvent {
	class := getVZRotationEventClass()
	rv := objc.Send[VZRotationEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZRotationEvent/initWithEvent:
func NewVZRotationEventWithEvent(event objectivec.IObject) VZRotationEvent {
	instance := getVZRotationEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEvent:"), event)
	return VZRotationEventFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZRotationEvent/initWithRotation:phase:
func NewVZRotationEventWithRotationPhase(rotation float64, phase uint64) VZRotationEvent {
	instance := getVZRotationEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithRotation:phase:"), rotation, phase)
	return VZRotationEventFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZRotationEvent/initWithEvent:
func (v VZRotationEvent) InitWithEvent(event objectivec.IObject) VZRotationEvent {
	rv := objc.Send[VZRotationEvent](v.ID, objc.Sel("initWithEvent:"), event)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZRotationEvent/initWithRotation:phase:
func (v VZRotationEvent) InitWithRotationPhase(rotation float64, phase uint64) VZRotationEvent {
	rv := objc.Send[VZRotationEvent](v.ID, objc.Sel("initWithRotation:phase:"), rotation, phase)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZRotationEvent/phase
func (v VZRotationEvent) Phase() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("phase"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZRotationEvent/rotation
func (v VZRotationEvent) Rotation() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("rotation"))
	return rv
}
