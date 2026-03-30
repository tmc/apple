// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZMultiTouchEvent] class.
var (
	_VZMultiTouchEventClass     VZMultiTouchEventClass
	_VZMultiTouchEventClassOnce sync.Once
)

func getVZMultiTouchEventClass() VZMultiTouchEventClass {
	_VZMultiTouchEventClassOnce.Do(func() {
		_VZMultiTouchEventClass = VZMultiTouchEventClass{class: objc.GetClass("_VZMultiTouchEvent")}
	})
	return _VZMultiTouchEventClass
}

// GetVZMultiTouchEventClass returns the class object for _VZMultiTouchEvent.
func GetVZMultiTouchEventClass() VZMultiTouchEventClass {
	return getVZMultiTouchEventClass()
}

type VZMultiTouchEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMultiTouchEventClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMultiTouchEventClass) Alloc() VZMultiTouchEvent {
	rv := objc.Send[VZMultiTouchEvent](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMultiTouchEvent.Touches]
//   - [VZMultiTouchEvent.InitWithEventView]
//   - [VZMultiTouchEvent.InitWithTouches]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMultiTouchEvent
type VZMultiTouchEvent struct {
	objectivec.Object
}

// VZMultiTouchEventFromID constructs a [VZMultiTouchEvent] from an objc.ID.
func VZMultiTouchEventFromID(id objc.ID) VZMultiTouchEvent {
	return VZMultiTouchEvent{objectivec.Object{ID: id}}
}

// Ensure VZMultiTouchEvent implements IVZMultiTouchEvent.
var _ IVZMultiTouchEvent = VZMultiTouchEvent{}

// An interface definition for the [VZMultiTouchEvent] class.
//
// # Methods
//
//   - [IVZMultiTouchEvent.Touches]
//   - [IVZMultiTouchEvent.InitWithEventView]
//   - [IVZMultiTouchEvent.InitWithTouches]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMultiTouchEvent
type IVZMultiTouchEvent interface {
	objectivec.IObject

	// Topic: Methods

	Touches() foundation.INSSet
	InitWithEventView(event objectivec.IObject, view objectivec.IObject) VZMultiTouchEvent
	InitWithTouches(touches objectivec.IObject) VZMultiTouchEvent
}

// Init initializes the instance.
func (v VZMultiTouchEvent) Init() VZMultiTouchEvent {
	rv := objc.Send[VZMultiTouchEvent](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMultiTouchEvent) Autorelease() VZMultiTouchEvent {
	rv := objc.Send[VZMultiTouchEvent](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMultiTouchEvent creates a new VZMultiTouchEvent instance.
func NewVZMultiTouchEvent() VZMultiTouchEvent {
	class := getVZMultiTouchEventClass()
	rv := objc.Send[VZMultiTouchEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMultiTouchEvent/initWithEvent:view:
func NewVZMultiTouchEventWithEventView(event objectivec.IObject, view objectivec.IObject) VZMultiTouchEvent {
	instance := getVZMultiTouchEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEvent:view:"), event, view)
	return VZMultiTouchEventFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMultiTouchEvent/initWithTouches:
func NewVZMultiTouchEventWithTouches(touches objectivec.IObject) VZMultiTouchEvent {
	instance := getVZMultiTouchEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTouches:"), touches)
	return VZMultiTouchEventFromID(rv)
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMultiTouchEvent/initWithEvent:view:
func (v VZMultiTouchEvent) InitWithEventView(event objectivec.IObject, view objectivec.IObject) VZMultiTouchEvent {
	rv := objc.Send[VZMultiTouchEvent](v.ID, objc.Sel("initWithEvent:view:"), event, view)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMultiTouchEvent/initWithTouches:
func (v VZMultiTouchEvent) InitWithTouches(touches objectivec.IObject) VZMultiTouchEvent {
	rv := objc.Send[VZMultiTouchEvent](v.ID, objc.Sel("initWithTouches:"), touches)
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMultiTouchEvent/touches
func (v VZMultiTouchEvent) Touches() foundation.INSSet {
	rv := objc.Send[objc.ID](v.ID, objc.Sel("touches"))
	return foundation.NSSetFromID(objc.ID(rv))
}
