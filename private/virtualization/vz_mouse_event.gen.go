// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMouseEvent] class.
var (
	_VZMouseEventClass     VZMouseEventClass
	_VZMouseEventClassOnce sync.Once
)

func getVZMouseEventClass() VZMouseEventClass {
	_VZMouseEventClassOnce.Do(func() {
		_VZMouseEventClass = VZMouseEventClass{class: objc.GetClass("_VZMouseEvent")}
	})
	return _VZMouseEventClass
}

// GetVZMouseEventClass returns the class object for _VZMouseEvent.
func GetVZMouseEventClass() VZMouseEventClass {
	return getVZMouseEventClass()
}

type VZMouseEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMouseEventClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMouseEventClass) Alloc() VZMouseEvent {
	rv := objc.Send[VZMouseEvent](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZMouseEvent.InitWithPressedButtonsDeltaXDeltaY]
// See: https://developer.apple.com/documentation/Virtualization/_VZMouseEvent
type VZMouseEvent struct {
	objectivec.Object
}

// VZMouseEventFromID constructs a [VZMouseEvent] from an objc.ID.
func VZMouseEventFromID(id objc.ID) VZMouseEvent {
	return VZMouseEvent{objectivec.Object{ID: id}}
}
// Ensure VZMouseEvent implements IVZMouseEvent.
var _ IVZMouseEvent = VZMouseEvent{}

// An interface definition for the [VZMouseEvent] class.
//
// # Methods
//
//   - [IVZMouseEvent.InitWithPressedButtonsDeltaXDeltaY]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMouseEvent
type IVZMouseEvent interface {
	objectivec.IObject

	// Topic: Methods

	InitWithPressedButtonsDeltaXDeltaY(buttons int64, x float64, y float64) VZMouseEvent
}

// Init initializes the instance.
func (v VZMouseEvent) Init() VZMouseEvent {
	rv := objc.Send[VZMouseEvent](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMouseEvent) Autorelease() VZMouseEvent {
	rv := objc.Send[VZMouseEvent](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMouseEvent creates a new VZMouseEvent instance.
func NewVZMouseEvent() VZMouseEvent {
	class := getVZMouseEventClass()
	rv := objc.Send[VZMouseEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZMouseEvent/initWithPressedButtons:deltaX:deltaY:
func NewVZMouseEventWithPressedButtonsDeltaXDeltaY(buttons int64, x float64, y float64) VZMouseEvent {
	instance := getVZMouseEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPressedButtons:deltaX:deltaY:"), buttons, x, y)
	return VZMouseEventFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZMouseEvent/initWithPressedButtons:deltaX:deltaY:
func (v VZMouseEvent) InitWithPressedButtonsDeltaXDeltaY(buttons int64, x float64, y float64) VZMouseEvent {
	rv := objc.Send[VZMouseEvent](v.ID, objc.Sel("initWithPressedButtons:deltaX:deltaY:"), buttons, x, y)
	return rv
}

