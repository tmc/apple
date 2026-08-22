// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The common properties and methods for objects that represent the input profile of a device.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput
type GCDevicePhysicalInput interface {
	objectivec.IObject
	GCDevicePhysicalInputState

	// Returns the next input state from the queue.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/nextInputState()
	NextInputState() interface {
		GCDevicePhysicalInputState
		GCDevicePhysicalInputStateDiff
	}

	// Returns a snapshot of the physical device inputs.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/capture()
	Capture() GCDevicePhysicalInputState

	// The device that the physical input represents.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/device
	Device() GCDevice

	// The maximum number of input values that the queue stores.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/inputStateQueueDepth
	InputStateQueueDepth() int
	SetInputStateQueueDepth(value int)

	// The dispatch queue that the system uses for callbacks.
	//
	// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/queue
	Queue() dispatch.Queue
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
func (o GCDevicePhysicalInputObject) NextInputState() interface {
	GCDevicePhysicalInputState
	GCDevicePhysicalInputStateDiff
} {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("nextInputState"))
	return struct {
		objectivec.Object
		GCDevicePhysicalInputStateObject
		GCDevicePhysicalInputStateDiffObject
	}{
		Object:                               objectivec.ObjectFromID(rv),
		GCDevicePhysicalInputStateObject:     GCDevicePhysicalInputStateObjectFromID(rv),
		GCDevicePhysicalInputStateDiffObject: GCDevicePhysicalInputStateDiffObjectFromID(rv),
	}
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

// The time of the most recent event.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/lastEventTimestamp
func (o GCDevicePhysicalInputObject) LastEventTimestamp() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](o.ID, objc.Sel("lastEventTimestamp"))
	return rv
}

// The time in seconds between the last event and the current time.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/lastEventLatency
func (o GCDevicePhysicalInputObject) LastEventLatency() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](o.ID, objc.Sel("lastEventLatency"))
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

// The device that the physical input represents.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/device
func (o GCDevicePhysicalInputObject) Device() GCDevice {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("device"))
	return GCDeviceObjectFromID(rv)
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
func (o GCDevicePhysicalInputObject) InputStateQueueDepth() int {
	rv := objc.Send[int](o.ID, objc.Sel("inputStateQueueDepth"))
	return int(rv)
}

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
func (o GCDevicePhysicalInputObject) Queue() dispatch.Queue {
	rv := objc.Send[uintptr](o.ID, objc.Sel("queue"))
	return dispatch.QueueFromHandle(rv)
}

func (o GCDevicePhysicalInputObject) SetQueue(value dispatch.Queue) {
	objc.Send[struct{}](o.ID, objc.Sel("setQueue:"), value)
}

// The device’s elements as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/elements-1shp2
func (o GCDevicePhysicalInputObject) Elements() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("elements"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s axes as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/axes-80rx
func (o GCDevicePhysicalInputObject) Axes() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("axes"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s buttons as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/buttons-3257g
func (o GCDevicePhysicalInputObject) Buttons() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("buttons"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s directional pads as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/dpads-5yr9x
func (o GCDevicePhysicalInputObject) Dpads() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("dpads"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s switches as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/switches-6bws2
func (o GCDevicePhysicalInputObject) Switches() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("switches"))
	return GCPhysicalInputElementCollectionFromID(rv)
}
