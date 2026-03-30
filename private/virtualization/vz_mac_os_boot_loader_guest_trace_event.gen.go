// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [VZMacOSBootLoaderGuestTraceEvent] class.
var (
	_VZMacOSBootLoaderGuestTraceEventClass     VZMacOSBootLoaderGuestTraceEventClass
	_VZMacOSBootLoaderGuestTraceEventClassOnce sync.Once
)

func getVZMacOSBootLoaderGuestTraceEventClass() VZMacOSBootLoaderGuestTraceEventClass {
	_VZMacOSBootLoaderGuestTraceEventClassOnce.Do(func() {
		_VZMacOSBootLoaderGuestTraceEventClass = VZMacOSBootLoaderGuestTraceEventClass{class: objc.GetClass("_VZMacOSBootLoaderGuestTraceEvent")}
	})
	return _VZMacOSBootLoaderGuestTraceEventClass
}

// GetVZMacOSBootLoaderGuestTraceEventClass returns the class object for _VZMacOSBootLoaderGuestTraceEvent.
func GetVZMacOSBootLoaderGuestTraceEventClass() VZMacOSBootLoaderGuestTraceEventClass {
	return getVZMacOSBootLoaderGuestTraceEventClass()
}

type VZMacOSBootLoaderGuestTraceEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZMacOSBootLoaderGuestTraceEventClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZMacOSBootLoaderGuestTraceEventClass) Alloc() VZMacOSBootLoaderGuestTraceEvent {
	rv := objc.Send[VZMacOSBootLoaderGuestTraceEvent](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZMacOSBootLoaderGuestTraceEvent.Data0]
//   - [VZMacOSBootLoaderGuestTraceEvent.Data1]
//   - [VZMacOSBootLoaderGuestTraceEvent.Data2]
//   - [VZMacOSBootLoaderGuestTraceEvent.Status]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacOSBootLoaderGuestTraceEvent
type VZMacOSBootLoaderGuestTraceEvent struct {
	VZGuestTraceEvent
}

// VZMacOSBootLoaderGuestTraceEventFromID constructs a [VZMacOSBootLoaderGuestTraceEvent] from an objc.ID.
func VZMacOSBootLoaderGuestTraceEventFromID(id objc.ID) VZMacOSBootLoaderGuestTraceEvent {
	return VZMacOSBootLoaderGuestTraceEvent{VZGuestTraceEvent: VZGuestTraceEventFromID(id)}
}

// Ensure VZMacOSBootLoaderGuestTraceEvent implements IVZMacOSBootLoaderGuestTraceEvent.
var _ IVZMacOSBootLoaderGuestTraceEvent = VZMacOSBootLoaderGuestTraceEvent{}

// An interface definition for the [VZMacOSBootLoaderGuestTraceEvent] class.
//
// # Methods
//
//   - [IVZMacOSBootLoaderGuestTraceEvent.Data0]
//   - [IVZMacOSBootLoaderGuestTraceEvent.Data1]
//   - [IVZMacOSBootLoaderGuestTraceEvent.Data2]
//   - [IVZMacOSBootLoaderGuestTraceEvent.Status]
//
// See: https://developer.apple.com/documentation/Virtualization/_VZMacOSBootLoaderGuestTraceEvent
type IVZMacOSBootLoaderGuestTraceEvent interface {
	IVZGuestTraceEvent

	// Topic: Methods

	Data0() uint32
	Data1() uint32
	Data2() uint32
	Status() uint32
}

// Init initializes the instance.
func (v VZMacOSBootLoaderGuestTraceEvent) Init() VZMacOSBootLoaderGuestTraceEvent {
	rv := objc.Send[VZMacOSBootLoaderGuestTraceEvent](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZMacOSBootLoaderGuestTraceEvent) Autorelease() VZMacOSBootLoaderGuestTraceEvent {
	rv := objc.Send[VZMacOSBootLoaderGuestTraceEvent](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZMacOSBootLoaderGuestTraceEvent creates a new VZMacOSBootLoaderGuestTraceEvent instance.
func NewVZMacOSBootLoaderGuestTraceEvent() VZMacOSBootLoaderGuestTraceEvent {
	class := getVZMacOSBootLoaderGuestTraceEventClass()
	rv := objc.Send[VZMacOSBootLoaderGuestTraceEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacOSBootLoaderGuestTraceEvent/data0
func (v VZMacOSBootLoaderGuestTraceEvent) Data0() uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("data0"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacOSBootLoaderGuestTraceEvent/data1
func (v VZMacOSBootLoaderGuestTraceEvent) Data1() uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("data1"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacOSBootLoaderGuestTraceEvent/data2
func (v VZMacOSBootLoaderGuestTraceEvent) Data2() uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("data2"))
	return rv
}

// See: https://developer.apple.com/documentation/Virtualization/_VZMacOSBootLoaderGuestTraceEvent/status
func (v VZMacOSBootLoaderGuestTraceEvent) Status() uint32 {
	rv := objc.Send[uint32](v.ID, objc.Sel("status"))
	return rv
}
