// Code generated from Apple documentation for virtualization. DO NOT EDIT.

package virtualization

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [VZSealedRegisterHistoryEvent] class.
var (
	_VZSealedRegisterHistoryEventClass     VZSealedRegisterHistoryEventClass
	_VZSealedRegisterHistoryEventClassOnce sync.Once
)

func getVZSealedRegisterHistoryEventClass() VZSealedRegisterHistoryEventClass {
	_VZSealedRegisterHistoryEventClassOnce.Do(func() {
		_VZSealedRegisterHistoryEventClass = VZSealedRegisterHistoryEventClass{class: objc.GetClass("_VZSealedRegisterHistoryEvent")}
	})
	return _VZSealedRegisterHistoryEventClass
}

// GetVZSealedRegisterHistoryEventClass returns the class object for _VZSealedRegisterHistoryEvent.
func GetVZSealedRegisterHistoryEventClass() VZSealedRegisterHistoryEventClass {
	return getVZSealedRegisterHistoryEventClass()
}

type VZSealedRegisterHistoryEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (vc VZSealedRegisterHistoryEventClass) Class() objc.Class {
	return vc.class
}

// Alloc allocates memory for a new instance of the class.
func (vc VZSealedRegisterHistoryEventClass) Alloc() VZSealedRegisterHistoryEvent {
	rv := objc.SendIfResponds[VZSealedRegisterHistoryEvent](objc.ID(vc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [VZSealedRegisterHistoryEvent.Data]
//   - [VZSealedRegisterHistoryEvent.Operation]
//   - [VZSealedRegisterHistoryEvent.Type]
//   - [VZSealedRegisterHistoryEvent.Uuid]
type VZSealedRegisterHistoryEvent struct {
	objectivec.Object
}

// VZSealedRegisterHistoryEventFromID constructs a [VZSealedRegisterHistoryEvent] from an objc.ID.
func VZSealedRegisterHistoryEventFromID(id objc.ID) VZSealedRegisterHistoryEvent {
	return VZSealedRegisterHistoryEvent{objectivec.Object{ID: id}}
}

// Ensure VZSealedRegisterHistoryEvent implements IVZSealedRegisterHistoryEvent.
var _ IVZSealedRegisterHistoryEvent = VZSealedRegisterHistoryEvent{}

// An interface definition for the [VZSealedRegisterHistoryEvent] class.
//
// # Methods
//
//   - [IVZSealedRegisterHistoryEvent.Data]
//   - [IVZSealedRegisterHistoryEvent.Operation]
//   - [IVZSealedRegisterHistoryEvent.Type]
//   - [IVZSealedRegisterHistoryEvent.Uuid]
type IVZSealedRegisterHistoryEvent interface {
	objectivec.IObject

	// Topic: Methods

	Data() foundation.NSData
	Operation() int64
	Type() int64
	Uuid() foundation.NSUUID
}

// Init initializes the instance.
func (v VZSealedRegisterHistoryEvent) Init() VZSealedRegisterHistoryEvent {
	rv := objc.SendIfResponds[VZSealedRegisterHistoryEvent](v.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (v VZSealedRegisterHistoryEvent) Autorelease() VZSealedRegisterHistoryEvent {
	rv := objc.SendIfResponds[VZSealedRegisterHistoryEvent](v.ID, objc.Sel("autorelease"))
	return rv
}

// NewVZSealedRegisterHistoryEvent creates a new VZSealedRegisterHistoryEvent instance.
func NewVZSealedRegisterHistoryEvent() VZSealedRegisterHistoryEvent {
	class := getVZSealedRegisterHistoryEventClass()
	rv := objc.SendIfResponds[VZSealedRegisterHistoryEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

func (v VZSealedRegisterHistoryEvent) Data() foundation.NSData {
	rv := objc.SendIfResponds[foundation.NSData](v.ID, objc.Sel("data"))
	return foundation.NSData(rv)
}
func (v VZSealedRegisterHistoryEvent) Operation() int64 {
	rv := objc.SendIfResponds[int64](v.ID, objc.Sel("operation"))
	return rv
}
func (v VZSealedRegisterHistoryEvent) Type() int64 {
	rv := objc.SendIfResponds[int64](v.ID, objc.Sel("type"))
	return rv
}
func (v VZSealedRegisterHistoryEvent) Uuid() foundation.NSUUID {
	rv := objc.SendIfResponds[foundation.NSUUID](v.ID, objc.Sel("uuid"))
	return foundation.NSUUID(rv)
}
