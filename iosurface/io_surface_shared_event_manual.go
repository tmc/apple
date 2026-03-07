package iosurface

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// IOSurfaceSharedEvent is a private IOSurface event object used for
// cross-process signaling via Mach ports.
type IOSurfaceSharedEvent struct {
	objectivec.Object
}

// IOSurfaceSharedEventFromID wraps an Objective-C object ID.
func IOSurfaceSharedEventFromID(id objc.ID) IOSurfaceSharedEvent {
	return IOSurfaceSharedEvent{objectivec.Object{ID: id}}
}

// IIOSurfaceSharedEvent is the minimal interface surface needed by callers.
type IIOSurfaceSharedEvent interface {
	objectivec.IObject
	InitWithOptions(options uint64) IOSurfaceSharedEvent
	InitWithMachPort(port uint32) IOSurfaceSharedEvent
	EventPort() uint32
}

var (
	_IOSurfaceSharedEventClass     IOSurfaceSharedEventClass
	_IOSurfaceSharedEventClassOnce sync.Once
)

func getIOSurfaceSharedEventClass() IOSurfaceSharedEventClass {
	_IOSurfaceSharedEventClassOnce.Do(func() {
		_IOSurfaceSharedEventClass = IOSurfaceSharedEventClass{
			class: objc.GetClass("IOSurfaceSharedEvent"),
		}
	})
	return _IOSurfaceSharedEventClass
}

// GetIOSurfaceSharedEventClass returns the Objective-C class object.
func GetIOSurfaceSharedEventClass() IOSurfaceSharedEventClass {
	return getIOSurfaceSharedEventClass()
}

// IOSurfaceSharedEventClass is the class object wrapper.
type IOSurfaceSharedEventClass struct {
	class objc.Class
}

// Alloc allocates an IOSurfaceSharedEvent instance.
func (ic IOSurfaceSharedEventClass) Alloc() IOSurfaceSharedEvent {
	rv := objc.Send[IOSurfaceSharedEvent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// SupportsSecureCoding reports class-level NSSecureCoding support.
func (ic IOSurfaceSharedEventClass) SupportsSecureCoding() bool {
	rv := objc.Send[bool](objc.ID(ic.class), objc.Sel("supportsSecureCoding"))
	return bool(rv)
}

// NewIOSurfaceSharedEvent allocates and initializes with -new.
func NewIOSurfaceSharedEvent() IOSurfaceSharedEvent {
	class := getIOSurfaceSharedEventClass()
	rv := objc.Send[IOSurfaceSharedEvent](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// NewIOSurfaceSharedEventWithOptions allocates and initializes with options.
func NewIOSurfaceSharedEventWithOptions(options uint64) IOSurfaceSharedEvent {
	instance := getIOSurfaceSharedEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOptions:"), options)
	return IOSurfaceSharedEventFromID(rv)
}

// NewIOSurfaceSharedEventWithMachPort allocates and initializes from a Mach port.
func NewIOSurfaceSharedEventWithMachPort(port uint32) IOSurfaceSharedEvent {
	instance := getIOSurfaceSharedEventClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithMachPort:"), port)
	return IOSurfaceSharedEventFromID(rv)
}

// Init initializes the instance.
func (s IOSurfaceSharedEvent) Init() IOSurfaceSharedEvent {
	rv := objc.Send[IOSurfaceSharedEvent](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s IOSurfaceSharedEvent) Autorelease() IOSurfaceSharedEvent {
	rv := objc.Send[IOSurfaceSharedEvent](s.ID, objc.Sel("autorelease"))
	return rv
}

// InitWithOptions initializes an event and allocates an underlying event port.
func (s IOSurfaceSharedEvent) InitWithOptions(options uint64) IOSurfaceSharedEvent {
	rv := objc.Send[IOSurfaceSharedEvent](s.ID, objc.Sel("initWithOptions:"), options)
	return rv
}

// InitWithMachPort initializes an event from an existing Mach port.
func (s IOSurfaceSharedEvent) InitWithMachPort(port uint32) IOSurfaceSharedEvent {
	rv := objc.Send[IOSurfaceSharedEvent](s.ID, objc.Sel("initWithMachPort:"), port)
	return rv
}

// EncodeWithCoder encodes this object for keyed/XPC coding.
func (s IOSurfaceSharedEvent) EncodeWithCoder(coder objectivec.IObject) {
	objc.Send[struct{}](s.ID, objc.Sel("encodeWithCoder:"), coder)
}

// EventPort returns the underlying Mach port name.
func (s IOSurfaceSharedEvent) EventPort() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("eventPort"))
	return uint32(rv)
}

// GlobalTraceObjectID returns the trace object identifier.
func (s IOSurfaceSharedEvent) GlobalTraceObjectID() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("globalTraceObjectID"))
	return uint64(rv)
}

// SignaledValue returns the current signaled value.
func (s IOSurfaceSharedEvent) SignaledValue() uint64 {
	rv := objc.Send[uint64](s.ID, objc.Sel("signaledValue"))
	return uint64(rv)
}

// SetSignaledValue updates the signaled value.
func (s IOSurfaceSharedEvent) SetSignaledValue(value uint64) {
	objc.Send[struct{}](s.ID, objc.Sel("setSignaledValue:"), value)
}

// SupportsRollback reports whether rollback is supported.
func (s IOSurfaceSharedEvent) SupportsRollback() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("supportsRollback"))
	return bool(rv)
}

// WaitUntilSignaledValueTimeoutMS blocks until the event reaches value or timeout.
func (s IOSurfaceSharedEvent) WaitUntilSignaledValueTimeoutMS(value uint64, timeoutMS uint64) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("waitUntilSignaledValue:timeoutMS:"), value, timeoutMS)
	return bool(rv)
}
