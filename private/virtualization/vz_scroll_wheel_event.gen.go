// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZScrollWheelEvent] class.
var (
	_VZScrollWheelEventClass     VZScrollWheelEventClass
	_VZScrollWheelEventClassOnce sync.Once
)

func getVZScrollWheelEventClass() VZScrollWheelEventClass {
	_VZScrollWheelEventClassOnce.Do(func() {
		_VZScrollWheelEventClass = VZScrollWheelEventClass{class: objc.GetClass("_VZScrollWheelEvent")}
	})
	return _VZScrollWheelEventClass
}

// GetVZScrollWheelEventClass returns the class object for _VZScrollWheelEvent.
func GetVZScrollWheelEventClass() VZScrollWheelEventClass {
	return getVZScrollWheelEventClass()
}

type VZScrollWheelEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZScrollWheelEventClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZScrollWheelEventClass) Alloc() VZScrollWheelEvent {
	rv := objc.Send[VZScrollWheelEvent](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZScrollWheelEvent.AcceleratedScrollingDeltaX]
//   - [VZScrollWheelEvent.AcceleratedScrollingDeltaY]
//   - [VZScrollWheelEvent.MomentumPhase]
//   - [VZScrollWheelEvent.ScrollPhase]
//   - [VZScrollWheelEvent.ScrollingDeltaX]
//   - [VZScrollWheelEvent.ScrollingDeltaY]
//   - [VZScrollWheelEvent.InitWithEvent]
//   - [VZScrollWheelEvent.InitWithScrollingDeltaXScrollingDeltaYAcceleratedScrollingDeltaXAcceleratedScrollingDeltaYScrollPhaseMomentumPhase]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZScrollWheelEvent
type VZScrollWheelEvent struct {
	objectivec.Object
}

// VZScrollWheelEventFromID constructs a [VZScrollWheelEvent] from an objc.ID.
func VZScrollWheelEventFromID(id objc.ID) VZScrollWheelEvent {
	return VZScrollWheelEvent{objectivec.Object{ID: id}}
}

// Ensure VZScrollWheelEvent implements IVZScrollWheelEvent.
var _ IVZScrollWheelEvent = VZScrollWheelEvent{}

// An interface definition for the [VZScrollWheelEvent] class.
//
// # Methods
//
//   - [IVZScrollWheelEvent.AcceleratedScrollingDeltaX]
//   - [IVZScrollWheelEvent.AcceleratedScrollingDeltaY]
//   - [IVZScrollWheelEvent.MomentumPhase]
//   - [IVZScrollWheelEvent.ScrollPhase]
//   - [IVZScrollWheelEvent.ScrollingDeltaX]
//   - [IVZScrollWheelEvent.ScrollingDeltaY]
//   - [IVZScrollWheelEvent.InitWithEvent]
//   - [IVZScrollWheelEvent.InitWithScrollingDeltaXScrollingDeltaYAcceleratedScrollingDeltaXAcceleratedScrollingDeltaYScrollPhaseMomentumPhase]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZScrollWheelEvent
type IVZScrollWheelEvent interface {
	objectivec.IObject

	// Topic: Methods

	AcceleratedScrollingDeltaX() float64
	AcceleratedScrollingDeltaY() float64
	MomentumPhase() uint64
	ScrollPhase() uint64
	ScrollingDeltaX() float64
	ScrollingDeltaY() float64
	InitWithEvent(event objectivec.IObject) VZScrollWheelEvent
	InitWithScrollingDeltaXScrollingDeltaYAcceleratedScrollingDeltaXAcceleratedScrollingDeltaYScrollPhaseMomentumPhase(x float64, y float64, x2 float64, y2 float64, phase uint64, phase2 uint64) VZScrollWheelEvent
}

// Init initializes the instance.
func (v VZScrollWheelEvent) Init() VZScrollWheelEvent {
	rv := objc.Send[VZScrollWheelEvent](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZScrollWheelEvent) Autorelease() VZScrollWheelEvent {
	rv := objc.Send[VZScrollWheelEvent](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZScrollWheelEvent creates a new VZScrollWheelEvent instance.
func NewVZScrollWheelEvent() VZScrollWheelEvent {
	class := getVZScrollWheelEventClass()
	rv := objc.Send[VZScrollWheelEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZScrollWheelEvent/initWithEvent:
func NewVZScrollWheelEventWithEvent(event objectivec.IObject) VZScrollWheelEvent {
	instance := getVZScrollWheelEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEvent:"), event)
	return VZScrollWheelEventFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZScrollWheelEvent/initWithScrollingDeltaX:scrollingDeltaY:acceleratedScrollingDeltaX:acceleratedScrollingDeltaY:scrollPhase:momentumPhase:
func NewVZScrollWheelEventWithScrollingDeltaXScrollingDeltaYAcceleratedScrollingDeltaXAcceleratedScrollingDeltaYScrollPhaseMomentumPhase(x float64, y float64, x2 float64, y2 float64, phase uint64, phase2 uint64) VZScrollWheelEvent {
	instance := getVZScrollWheelEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithScrollingDeltaX:scrollingDeltaY:acceleratedScrollingDeltaX:acceleratedScrollingDeltaY:scrollPhase:momentumPhase:"), x, y, x2, y2, phase, phase2)
	return VZScrollWheelEventFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZScrollWheelEvent/initWithEvent:
func (v VZScrollWheelEvent) InitWithEvent(event objectivec.IObject) VZScrollWheelEvent {
	rv := objc.Send[VZScrollWheelEvent](v.ID, objc.Sel("initWithEvent:"), event)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZScrollWheelEvent/initWithScrollingDeltaX:scrollingDeltaY:acceleratedScrollingDeltaX:acceleratedScrollingDeltaY:scrollPhase:momentumPhase:
func (v VZScrollWheelEvent) InitWithScrollingDeltaXScrollingDeltaYAcceleratedScrollingDeltaXAcceleratedScrollingDeltaYScrollPhaseMomentumPhase(x float64, y float64, x2 float64, y2 float64, phase uint64, phase2 uint64) VZScrollWheelEvent {
	rv := objc.Send[VZScrollWheelEvent](v.ID, objc.Sel("initWithScrollingDeltaX:scrollingDeltaY:acceleratedScrollingDeltaX:acceleratedScrollingDeltaY:scrollPhase:momentumPhase:"), x, y, x2, y2, phase, phase2)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZScrollWheelEvent/acceleratedScrollingDeltaX
func (v VZScrollWheelEvent) AcceleratedScrollingDeltaX() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("acceleratedScrollingDeltaX"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZScrollWheelEvent/acceleratedScrollingDeltaY
func (v VZScrollWheelEvent) AcceleratedScrollingDeltaY() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("acceleratedScrollingDeltaY"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZScrollWheelEvent/momentumPhase
func (v VZScrollWheelEvent) MomentumPhase() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("momentumPhase"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZScrollWheelEvent/scrollPhase
func (v VZScrollWheelEvent) ScrollPhase() uint64 {
	rv := objc.Send[uint64](v.ID, objc.Sel("scrollPhase"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZScrollWheelEvent/scrollingDeltaX
func (v VZScrollWheelEvent) ScrollingDeltaX() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("scrollingDeltaX"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZScrollWheelEvent/scrollingDeltaY
func (v VZScrollWheelEvent) ScrollingDeltaY() float64 {
	rv := objc.Send[float64](v.ID, objc.Sel("scrollingDeltaY"))
	return rv
}
