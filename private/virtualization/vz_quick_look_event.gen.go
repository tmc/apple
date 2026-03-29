// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZQuickLookEvent] class.
var (
	_VZQuickLookEventClass     VZQuickLookEventClass
	_VZQuickLookEventClassOnce sync.Once
)

func getVZQuickLookEventClass() VZQuickLookEventClass {
	_VZQuickLookEventClassOnce.Do(func() {
		_VZQuickLookEventClass = VZQuickLookEventClass{class: objc.GetClass("_VZQuickLookEvent")}
	})
	return _VZQuickLookEventClass
}

// GetVZQuickLookEventClass returns the class object for _VZQuickLookEvent.
func GetVZQuickLookEventClass() VZQuickLookEventClass {
	return getVZQuickLookEventClass()
}

type VZQuickLookEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZQuickLookEventClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZQuickLookEventClass) Alloc() VZQuickLookEvent {
	rv := objc.Send[VZQuickLookEvent](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

//
// # Methods
//
//   - [VZQuickLookEvent.InitWithEvent]
// See: https://developer.apple.com/documentation/Virtualization/_VZQuickLookEvent
type VZQuickLookEvent struct {
	objectivec.Object
}

// VZQuickLookEventFromID constructs a [VZQuickLookEvent] from an objc.ID.
func VZQuickLookEventFromID(id objc.ID) VZQuickLookEvent {
	return VZQuickLookEvent{objectivec.Object{ID: id}}
}
// Ensure VZQuickLookEvent implements IVZQuickLookEvent.
var _ IVZQuickLookEvent = VZQuickLookEvent{}

// An interface definition for the [VZQuickLookEvent] class.
//
// # Methods
//
//   - [IVZQuickLookEvent.InitWithEvent]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZQuickLookEvent
type IVZQuickLookEvent interface {
	objectivec.IObject

	// Topic: Methods

	InitWithEvent(event objectivec.IObject) VZQuickLookEvent
}

// Init initializes the instance.
func (v VZQuickLookEvent) Init() VZQuickLookEvent {
	rv := objc.Send[VZQuickLookEvent](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZQuickLookEvent) Autorelease() VZQuickLookEvent {
	rv := objc.Send[VZQuickLookEvent](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZQuickLookEvent creates a new VZQuickLookEvent instance.
func NewVZQuickLookEvent() VZQuickLookEvent {
	class := getVZQuickLookEventClass()
	rv := objc.Send[VZQuickLookEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZQuickLookEvent/initWithEvent:
func NewVZQuickLookEventWithEvent(event objectivec.IObject) VZQuickLookEvent {
	instance := getVZQuickLookEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithEvent:"), event)
	return VZQuickLookEventFromID(rv)
}

//
// See: https://developer.apple.com/documentation/Virtualization/_VZQuickLookEvent/initWithEvent:
func (v VZQuickLookEvent) InitWithEvent(event objectivec.IObject) VZQuickLookEvent {
	rv := objc.Send[VZQuickLookEvent](v.ID, objc.Sel("initWithEvent:"), event)
	return rv
}

