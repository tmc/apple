// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A type that synchronizes memory operations to one or more resources across multiple CPUs, GPUs, and processes.
//
// See: https://developer.apple.com/documentation/Metal/MTLSharedEvent
type MTLSharedEvent interface {
	objectivec.IObject
	MTLEvent

	// The current signal value for the shareable event.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLSharedEvent/signaledValue
	SignaledValue() uint64

	// Schedules a notification handler to be called after the shareable event’s signal value equals or exceeds a given value.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLSharedEvent/notify(_:atValue:block:)
	NotifyListenerAtValueBlock(listener IMTLSharedEventListener, value uint64, block MTLSharedEventNotificationBlock)

	// Creates a new shareable event handle.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLSharedEvent/makeSharedEventHandle()
	NewSharedEventHandle() IMTLSharedEventHandle

	// WaitUntilSignaledValueTimeoutMS protocol.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLSharedEvent/wait(untilSignaledValue:timeoutMS:)
	WaitUntilSignaledValueTimeoutMS(value uint64, milliseconds uint64) bool

	// The current signal value for the shareable event.
	//
	// See: https://developer.apple.com/documentation/Metal/MTLSharedEvent/signaledValue
	SetSignaledValue(value uint64)
}

// MTLSharedEventObject wraps an existing Objective-C object that conforms to the MTLSharedEvent protocol.
type MTLSharedEventObject struct {
	objectivec.Object
}

func (o MTLSharedEventObject) BaseObject() objectivec.Object {
	return o.Object
}

// MTLSharedEventObjectFromID constructs a [MTLSharedEventObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func MTLSharedEventObjectFromID(id objc.ID) MTLSharedEventObject {
	return MTLSharedEventObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The current signal value for the shareable event.
//
// See: https://developer.apple.com/documentation/Metal/MTLSharedEvent/signaledValue
func (o MTLSharedEventObject) SignaledValue() uint64 {
	rv := objc.Send[uint64](o.ID, objc.Sel("signaledValue"))
	return rv
}

// Schedules a notification handler to be called after the shareable event’s
// signal value equals or exceeds a given value.
//
// listener: The listener object used to dispatch the notification.
//
// value: The minimum value that needs to be signaled before the notification handler
// is called.
//
// block: The notification handler to call.
//
// See: https://developer.apple.com/documentation/Metal/MTLSharedEvent/notify(_:atValue:block:)
func (o MTLSharedEventObject) NotifyListenerAtValueBlock(listener IMTLSharedEventListener, value uint64, block MTLSharedEventNotificationBlock) {
	objc.Send[struct{}](o.ID, objc.Sel("notifyListener:atValue:block:"), listener, value, block)
}

// Creates a new shareable event handle.
//
// See: https://developer.apple.com/documentation/Metal/MTLSharedEvent/makeSharedEventHandle()
func (o MTLSharedEventObject) NewSharedEventHandle() IMTLSharedEventHandle {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("newSharedEventHandle"))
	return MTLSharedEventHandleFromID(rv)
}

// See: https://developer.apple.com/documentation/Metal/MTLSharedEvent/wait(untilSignaledValue:timeoutMS:)
func (o MTLSharedEventObject) WaitUntilSignaledValueTimeoutMS(value uint64, milliseconds uint64) bool {
	rv := objc.Send[bool](o.ID, objc.Sel("waitUntilSignaledValue:timeoutMS:"), value, milliseconds)
	return rv
}

// The device object that created the event.
//
// See: https://developer.apple.com/documentation/Metal/MTLEvent/device
func (o MTLSharedEventObject) Device() MTLDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return MTLDeviceObjectFromID(rv)
}

// A string that identifies the event.
//
// See: https://developer.apple.com/documentation/Metal/MTLEvent/label
func (o MTLSharedEventObject) Label() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("label"))
	return foundation.NSStringFromID(rv).String()
}

// The current signal value for the shareable event.
//
// # Discussion
//
// When you set the value of a shared event, its value is changed only if you
// provide a larger value than the value currently stored in the event.
// Setting this property signals the event. Commands waiting on the event are
// allowed to run if the new value is equal to or greater than the value for
// which they are waiting. Similarly, setting the event’s value triggers
// notifications if the value is equal to or greater than the value for which
// they are waiting.
//
// See: https://developer.apple.com/documentation/Metal/MTLSharedEvent/signaledValue
func (o MTLSharedEventObject) SetSignaledValue(value uint64) {
	objc.Send[struct{}](o.ID, objc.Sel("setSignaledValue:"), value)
}

// A string that identifies the event.
//
// # Discussion
//
// Object and command labels are useful identifiers at runtime or when
// profiling and debugging your app using any Metal tool. See [Naming
// resources and commands].
//
// See: https://developer.apple.com/documentation/Metal/MTLEvent/label
//
// [Naming resources and commands]: https://developer.apple.com/documentation/Xcode/Naming-resources-and-commands
func (o MTLSharedEventObject) SetLabel(value string) {
	objc.Send[struct{}](o.ID, objc.Sel("setLabel:"), objc.String(value))
}
