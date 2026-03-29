// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZScreenCoordinatePointerEvent] class.
var (
	_VZScreenCoordinatePointerEventClass     VZScreenCoordinatePointerEventClass
	_VZScreenCoordinatePointerEventClassOnce sync.Once
)

func getVZScreenCoordinatePointerEventClass() VZScreenCoordinatePointerEventClass {
	_VZScreenCoordinatePointerEventClassOnce.Do(func() {
		_VZScreenCoordinatePointerEventClass = VZScreenCoordinatePointerEventClass{class: objc.GetClass("_VZScreenCoordinatePointerEvent")}
	})
	return _VZScreenCoordinatePointerEventClass
}

// GetVZScreenCoordinatePointerEventClass returns the class object for _VZScreenCoordinatePointerEvent.
func GetVZScreenCoordinatePointerEventClass() VZScreenCoordinatePointerEventClass {
	return getVZScreenCoordinatePointerEventClass()
}

type VZScreenCoordinatePointerEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZScreenCoordinatePointerEventClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZScreenCoordinatePointerEventClass) Alloc() VZScreenCoordinatePointerEvent {
	rv := objc.Send[VZScreenCoordinatePointerEvent](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZScreenCoordinatePointerEvent.Location]
//   - [VZScreenCoordinatePointerEvent.PressedButtons]
//   - [VZScreenCoordinatePointerEvent.InitWithEventView]
//   - [VZScreenCoordinatePointerEvent.InitWithLocationPressedButtons]
// See: https://developer.apple.com/documentation/Virtualization/_VZScreenCoordinatePointerEvent
type VZScreenCoordinatePointerEvent struct {
	objectivec.Object
}

// VZScreenCoordinatePointerEventFromID constructs a [VZScreenCoordinatePointerEvent] from an objc.ID.
func VZScreenCoordinatePointerEventFromID(id objc.ID) VZScreenCoordinatePointerEvent {
	return VZScreenCoordinatePointerEvent{objectivec.Object{ID: id}}
}
// Ensure VZScreenCoordinatePointerEvent implements IVZScreenCoordinatePointerEvent.
var _ IVZScreenCoordinatePointerEvent = VZScreenCoordinatePointerEvent{}

// An interface definition for the [VZScreenCoordinatePointerEvent] class.
//
// # Methods
//
//   - [IVZScreenCoordinatePointerEvent.Location]
//   - [IVZScreenCoordinatePointerEvent.PressedButtons]
//   - [IVZScreenCoordinatePointerEvent.InitWithEventView]
//   - [IVZScreenCoordinatePointerEvent.InitWithLocationPressedButtons]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZScreenCoordinatePointerEvent
type IVZScreenCoordinatePointerEvent interface {
	objectivec.IObject

	// Topic: Methods

	Location() corefoundation.CGPoint
	PressedButtons() int64
	InitWithEventView(event objectivec.IObject, view objectivec.IObject) VZScreenCoordinatePointerEvent
	InitWithLocationPressedButtons(location corefoundation.CGPoint, buttons int64) VZScreenCoordinatePointerEvent
}

// Init initializes the instance.
func (v VZScreenCoordinatePointerEvent) Init() VZScreenCoordinatePointerEvent {
	rv := objc.Send[VZScreenCoordinatePointerEvent](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZScreenCoordinatePointerEvent) Autorelease() VZScreenCoordinatePointerEvent {
	rv := objc.Send[VZScreenCoordinatePointerEvent](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZScreenCoordinatePointerEvent creates a new VZScreenCoordinatePointerEvent instance.
func NewVZScreenCoordinatePointerEvent() VZScreenCoordinatePointerEvent {
	class := getVZScreenCoordinatePointerEventClass()
	rv := objc.Send[VZScreenCoordinatePointerEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZScreenCoordinatePointerEvent/initWithEvent:view:
func NewVZScreenCoordinatePointerEventWithEventView(event objectivec.IObject, view objectivec.IObject) VZScreenCoordinatePointerEvent {
	instance := getVZScreenCoordinatePointerEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEvent:view:"), event, view)
	return VZScreenCoordinatePointerEventFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZScreenCoordinatePointerEvent/initWithLocation:pressedButtons:
func NewVZScreenCoordinatePointerEventWithLocationPressedButtons(location corefoundation.CGPoint, buttons int64) VZScreenCoordinatePointerEvent {
	instance := getVZScreenCoordinatePointerEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithLocation:pressedButtons:"), location, buttons)
	return VZScreenCoordinatePointerEventFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZScreenCoordinatePointerEvent/initWithEvent:view:
func (v VZScreenCoordinatePointerEvent) InitWithEventView(event objectivec.IObject, view objectivec.IObject) VZScreenCoordinatePointerEvent {
	rv := objc.Send[VZScreenCoordinatePointerEvent](v.ID, objc.Sel("initWithEvent:view:"), event, view)
	return rv
}
//
// See: https://developer.apple.com/documentation/Virtualization/_VZScreenCoordinatePointerEvent/initWithLocation:pressedButtons:
func (v VZScreenCoordinatePointerEvent) InitWithLocationPressedButtons(location corefoundation.CGPoint, buttons int64) VZScreenCoordinatePointerEvent {
	rv := objc.Send[VZScreenCoordinatePointerEvent](v.ID, objc.Sel("initWithLocation:pressedButtons:"), location, buttons)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZScreenCoordinatePointerEvent/location
func (v VZScreenCoordinatePointerEvent) Location() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](v.ID, objc.Sel("location"))
	return corefoundation.CGPoint(rv)
}
// See: https://developer.apple.com/documentation/Virtualization/_VZScreenCoordinatePointerEvent/pressedButtons
func (v VZScreenCoordinatePointerEvent) PressedButtons() int64 {
	rv := objc.Send[int64](v.ID, objc.Sel("pressedButtons"))
	return rv
}

