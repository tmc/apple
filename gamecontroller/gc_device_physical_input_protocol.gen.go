// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The common properties and methods for objects that represent the input profile of a device.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput
type GCDevicePhysicalInput interface {
	objectivec.IObject
	GCDevicePhysicalInputState

	// The device that the physical input represents.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/device
	Device() GCDevice

	// Returns the next input state from the queue.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/nextInputState()
	NextInputState() objectivec.IObject

	// The block that the profile calls when Game Controller adds an input state to the queue.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/inputStateAvailableHandler
	InputStateAvailableHandler() func(objc.ID)

	// The maximum number of input values that the queue stores.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/inputStateQueueDepth
	InputStateQueueDepth() int

	// Returns a snapshot of the physical device inputs.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/capture()
	Capture() GCDevicePhysicalInputState

	// A block that the profile calls when an element’s value changes.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/elementValueDidChangeHandler
	ElementValueDidChangeHandler() func(objc.ID)

	// The dispatch queue that the system uses for callbacks.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/queue
	Queue() dispatch.Queue

	// The maximum number of input values that the queue stores.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/inputStateQueueDepth
	SetInputStateQueueDepth(value int)

	// The dispatch queue that the system uses for callbacks.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/queue
	SetQueue(value dispatch.Queue)
}

// GCDevicePhysicalInputObject wraps an existing Objective-C object that conforms to the GCDevicePhysicalInput protocol.
type GCDevicePhysicalInputObject struct {
	objectivec.Object
}

func (o GCDevicePhysicalInputObject) BaseObject() objectivec.Object {
	return o.Object
}

// GCDevicePhysicalInputObjectFromID constructs a [GCDevicePhysicalInputObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func GCDevicePhysicalInputObjectFromID(id objc.ID) GCDevicePhysicalInputObject {
	return GCDevicePhysicalInputObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The device that the physical input represents.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/device
func (o GCDevicePhysicalInputObject) Device() GCDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return GCDeviceObjectFromID(rv)
}

// Returns the next input state from the queue.
//
// # Return Value
//
// The next input state in the queue or `nil` if the queue is empty.
//
// # Discussion
//
// This method removes the next input state from the queue.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/nextInputState()
func (o GCDevicePhysicalInputObject) NextInputState() objectivec.IObject {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("nextInputState"))
	return objectivec.Object{ID: rv}
}

// The block that the profile calls when Game Controller adds an input state
// to the queue.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/inputStateAvailableHandler
func (o GCDevicePhysicalInputObject) InputStateAvailableHandler() func(objc.ID) {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("inputStateAvailableHandler"))
	// Block/function return - cannot convert from objc.ID to Go func
	_ = rv
	return nil
}

// The maximum number of input values that the queue stores.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/inputStateQueueDepth
func (o GCDevicePhysicalInputObject) InputStateQueueDepth() int {
	rv := objc.Send[int](o.ID, objc.Sel("inputStateQueueDepth"))
	return rv
}

// Returns a snapshot of the physical device inputs.
//
// # Return Value
//
// A new instance containing the current state of the physical device input.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/capture()
func (o GCDevicePhysicalInputObject) Capture() GCDevicePhysicalInputState {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("capture"))
	return GCDevicePhysicalInputStateObjectFromID(rv)
}

// A block that the profile calls when an element’s value changes.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/elementValueDidChangeHandler
func (o GCDevicePhysicalInputObject) ElementValueDidChangeHandler() func(objc.ID) {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("elementValueDidChangeHandler"))
	// Block/function return - cannot convert from objc.ID to Go func
	_ = rv
	return nil
}

// The dispatch queue that the system uses for callbacks.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/queue
func (o GCDevicePhysicalInputObject) Queue() dispatch.Queue {
	rv := objc.Send[uintptr](o.ID, objc.Sel("queue"))
	return dispatch.QueueFromHandle(rv)
}

// The time of the most recent event.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/lastEventTimestamp
func (o GCDevicePhysicalInputObject) LastEventTimestamp() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("lastEventTimestamp"))
	return rv
}

// The time in seconds between the last event and the current time.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/lastEventLatency
func (o GCDevicePhysicalInputObject) LastEventLatency() float64 {
	rv := objc.Send[float64](o.ID, objc.Sel("lastEventLatency"))
	return rv
}

// Returns the element that the key specifies.
//
// key: A key that identifies an element.
//
// # Return Value
//
// The element that matches the key.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/subscript(_:)
func (o GCDevicePhysicalInputObject) ObjectForKeyedSubscript(key string) GCPhysicalInputElement {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("objectForKeyedSubscript:"), objc.String(key))
	return GCPhysicalInputElementObjectFromID(rv)
}

// The maximum number of input values that the queue stores.
//
// # Discussion
//
// When the queue reaches this limit, Game Controller starts removing the
// oldest input states from the queue. The default value for this property is
// `1` which indicates no buffering.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/inputStateQueueDepth
func (o GCDevicePhysicalInputObject) SetInputStateQueueDepth(value int) {
	objc.Send[struct{}](o.ID, objc.Sel("setInputStateQueueDepth:"), value)
}

// The dispatch queue that the system uses for callbacks.
//
// # Discussion
//
// Objects that conform to the [GCDevicePhysicalInput] protocol dispatch
// callbacks on the device’s [HandlerQueue] property by default. If you want
// to use a different dispatch queue, set this property to the preferred queue
// before you set callbacks.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/queue
func (o GCDevicePhysicalInputObject) SetQueue(value dispatch.Queue) {
	objc.Send[struct{}](o.ID, objc.Sel("setQueue:"), value)
}
