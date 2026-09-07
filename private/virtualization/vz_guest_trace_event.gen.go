// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZGuestTraceEvent] class.
var (
	_VZGuestTraceEventClass     VZGuestTraceEventClass
	_VZGuestTraceEventClassOnce sync.Once
)

func getVZGuestTraceEventClass() VZGuestTraceEventClass {
	_VZGuestTraceEventClassOnce.Do(func() {
		_VZGuestTraceEventClass = VZGuestTraceEventClass{class: objc.GetClass("_VZGuestTraceEvent")}
	})
	return _VZGuestTraceEventClass
}

// GetVZGuestTraceEventClass returns the class object for _VZGuestTraceEvent.
func GetVZGuestTraceEventClass() VZGuestTraceEventClass {
	return getVZGuestTraceEventClass()
}

type VZGuestTraceEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZGuestTraceEventClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZGuestTraceEventClass) Alloc() VZGuestTraceEvent {
	rv := objc.SendIfResponds[VZGuestTraceEvent](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZGuestTraceEvent._init]
type VZGuestTraceEvent struct {
	objectivec.Object
}

// VZGuestTraceEventFromID constructs a [VZGuestTraceEvent] from an objc.ID.
func VZGuestTraceEventFromID(id objc.ID) VZGuestTraceEvent {
	return VZGuestTraceEvent{objectivec.Object{ID: id}}
}

// Ensure VZGuestTraceEvent implements IVZGuestTraceEvent.
var _ IVZGuestTraceEvent = VZGuestTraceEvent{}

// An interface definition for the [VZGuestTraceEvent] class.
//
// # Methods
//
//   - [IVZGuestTraceEvent._init]
type IVZGuestTraceEvent interface {
	objectivec.IObject

	// Topic: Methods

	_init() objectivec.IObject
}

// Init initializes the instance.
func (v VZGuestTraceEvent) Init() VZGuestTraceEvent {
	rv := objc.SendIfResponds[VZGuestTraceEvent](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZGuestTraceEvent) Autorelease() VZGuestTraceEvent {
	rv := objc.SendIfResponds[VZGuestTraceEvent](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZGuestTraceEvent creates a new VZGuestTraceEvent instance.
func NewVZGuestTraceEvent() VZGuestTraceEvent {
	class := getVZGuestTraceEventClass()
	rv := objc.SendIfResponds[VZGuestTraceEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZGuestTraceEvent) _init() objectivec.IObject {
	rv := objc.SendIfResponds[objc.ID](v.ID, objc.Sel("_init"))
	return objectivec.Object{ID: rv}
}
