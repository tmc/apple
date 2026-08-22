// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [BNNSDeviceSharedEvent] class.
var (
	_BNNSDeviceSharedEventClass     BNNSDeviceSharedEventClass
	_BNNSDeviceSharedEventClassOnce sync.Once
)

func getBNNSDeviceSharedEventClass() BNNSDeviceSharedEventClass {
	_BNNSDeviceSharedEventClassOnce.Do(func() {
		_BNNSDeviceSharedEventClass = BNNSDeviceSharedEventClass{class: objc.GetClass("_TtCC6CoreML10BNNSDevice11SharedEvent")}
	})
	return _BNNSDeviceSharedEventClass
}

// GetBNNSDeviceSharedEventClass returns the class object for _TtCC6CoreML10BNNSDevice11SharedEvent.
func GetBNNSDeviceSharedEventClass() BNNSDeviceSharedEventClass {
	return getBNNSDeviceSharedEventClass()
}

type BNNSDeviceSharedEventClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (bc BNNSDeviceSharedEventClass) Class() objc.Class {
	return bc.class
}

// Alloc allocates memory for a new instance of the class.
func (bc BNNSDeviceSharedEventClass) Alloc() BNNSDeviceSharedEvent {
	rv := objc.SendIfResponds[BNNSDeviceSharedEvent](objc.ID(bc.class), objc.Sel("alloc"))
	return rv
}

type BNNSDeviceSharedEvent struct {
	objectivec.Object
}

// BNNSDeviceSharedEventFromID constructs a [BNNSDeviceSharedEvent] from an objc.ID.
func BNNSDeviceSharedEventFromID(id objc.ID) BNNSDeviceSharedEvent {
	return BNNSDeviceSharedEvent{objectivec.Object{ID: id}}
}

// Ensure BNNSDeviceSharedEvent implements IBNNSDeviceSharedEvent.
var _ IBNNSDeviceSharedEvent = BNNSDeviceSharedEvent{}

// An interface definition for the [BNNSDeviceSharedEvent] class.
type IBNNSDeviceSharedEvent interface {
	objectivec.IObject
}

// Init initializes the instance.
func (b BNNSDeviceSharedEvent) Init() BNNSDeviceSharedEvent {
	rv := objc.SendIfResponds[BNNSDeviceSharedEvent](b.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (b BNNSDeviceSharedEvent) Autorelease() BNNSDeviceSharedEvent {
	rv := objc.SendIfResponds[BNNSDeviceSharedEvent](b.ID, objc.Sel("autorelease"))
	return rv
}

// NewBNNSDeviceSharedEvent creates a new BNNSDeviceSharedEvent instance.
func NewBNNSDeviceSharedEvent() BNNSDeviceSharedEvent {
	class := getBNNSDeviceSharedEventClass()
	rv := objc.SendIfResponds[BNNSDeviceSharedEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}
