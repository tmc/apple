// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [MetalDeviceSharedEvent] class.
var (
	_MetalDeviceSharedEventClass     MetalDeviceSharedEventClass
	_MetalDeviceSharedEventClassOnce sync.Once
)

func getMetalDeviceSharedEventClass() MetalDeviceSharedEventClass {
	_MetalDeviceSharedEventClassOnce.Do(func() {
		_MetalDeviceSharedEventClass = MetalDeviceSharedEventClass{class: objc.GetClass("_TtCC6CoreML11MetalDevice11SharedEvent")}
	})
	return _MetalDeviceSharedEventClass
}

// GetMetalDeviceSharedEventClass returns the class object for _TtCC6CoreML11MetalDevice11SharedEvent.
func GetMetalDeviceSharedEventClass() MetalDeviceSharedEventClass {
	return getMetalDeviceSharedEventClass()
}

type MetalDeviceSharedEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (mc MetalDeviceSharedEventClass) Class() objc.Class {
	return mc.class
}

// Alloc allocates memory for a new instance of the class.
func (mc MetalDeviceSharedEventClass) Alloc() MetalDeviceSharedEvent {
	rv := objc.SendIfResponds[MetalDeviceSharedEvent](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

type MetalDeviceSharedEvent struct {
	objectivec.Object
}

// MetalDeviceSharedEventFromID constructs a [MetalDeviceSharedEvent] from an objc.ID.
func MetalDeviceSharedEventFromID(id objc.ID) MetalDeviceSharedEvent {
	return MetalDeviceSharedEvent{objectivec.Object{ID: id}}
}

// Ensure MetalDeviceSharedEvent implements IMetalDeviceSharedEvent.
var _ IMetalDeviceSharedEvent = MetalDeviceSharedEvent{}

// An interface definition for the [MetalDeviceSharedEvent] class.
type IMetalDeviceSharedEvent interface {
	objectivec.IObject
}

// Init initializes the instance.
func (m MetalDeviceSharedEvent) Init() MetalDeviceSharedEvent {
	rv := objc.SendIfResponds[MetalDeviceSharedEvent](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m MetalDeviceSharedEvent) Autorelease() MetalDeviceSharedEvent {
	rv := objc.SendIfResponds[MetalDeviceSharedEvent](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewMetalDeviceSharedEvent creates a new MetalDeviceSharedEvent instance.
func NewMetalDeviceSharedEvent() MetalDeviceSharedEvent {
	class := getMetalDeviceSharedEventClass()
	rv := objc.SendIfResponds[MetalDeviceSharedEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}
