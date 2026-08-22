// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GCRacingWheelInput] class.
var (
	_GCRacingWheelInputClass     GCRacingWheelInputClass
	_GCRacingWheelInputClassOnce sync.Once
)

func getGCRacingWheelInputClass() GCRacingWheelInputClass {
	_GCRacingWheelInputClassOnce.Do(func() {
		_GCRacingWheelInputClass = GCRacingWheelInputClass{class: objc.GetClass("GCRacingWheelInput")}
	})
	return _GCRacingWheelInputClass
}

// GetGCRacingWheelInputClass returns the class object for GCRacingWheelInput.
func GetGCRacingWheelInputClass() GCRacingWheelInputClass {
	return getGCRacingWheelInputClass()
}

type GCRacingWheelInputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCRacingWheelInputClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCRacingWheelInputClass) Alloc() GCRacingWheelInput {
	rv := objc.Send[GCRacingWheelInput](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// A controller profile that supports a racing wheel.
//
// # Creating snapshots
//
//   - [GCRacingWheelInput.Capture]: Returns a snapshot of the racing wheel inputs.
//
// # Polling for input
//
//   - [GCRacingWheelInput.NextInputState]: Returns the next input state of the racing wheel from the queue.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheelInput
type GCRacingWheelInput struct {
	GCRacingWheelInputState
}

// GCRacingWheelInputFromID constructs a [GCRacingWheelInput] from an objc.ID.
//
// A controller profile that supports a racing wheel.
func GCRacingWheelInputFromID(id objc.ID) GCRacingWheelInput {
	return GCRacingWheelInput{GCRacingWheelInputState: GCRacingWheelInputStateFromID(id)}
}

// NOTE: GCRacingWheelInput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCRacingWheelInput] class.
//
// # Creating snapshots
//
//   - [IGCRacingWheelInput.Capture]: Returns a snapshot of the racing wheel inputs.
//
// # Polling for input
//
//   - [IGCRacingWheelInput.NextInputState]: Returns the next input state of the racing wheel from the queue.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheelInput
type IGCRacingWheelInput interface {
	IGCRacingWheelInputState

	// Topic: Creating snapshots

	// Returns a snapshot of the racing wheel inputs.
	Capture() GCDevicePhysicalInputState

	// Topic: Polling for input

	// Returns the next input state of the racing wheel from the queue.
	NextInputState() interface {
		GCDevicePhysicalInputState
		GCDevicePhysicalInputStateDiff
	}

	// A block that the profile calls when an element’s value changes.
	ElementValueDidChangeHandler() func(objectivec.IObject)
	// The block that the profile calls when Game Controller adds an input state to the queue.
	InputStateAvailableHandler() func(objectivec.IObject)
	// The maximum number of input values that the queue stores.
	InputStateQueueDepth() int
	// The dispatch queue that the system uses for callbacks.
	Queue() dispatch.Queue
}

// Init initializes the instance.
func (g GCRacingWheelInput) Init() GCRacingWheelInput {
	rv := objc.Send[GCRacingWheelInput](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCRacingWheelInput) Autorelease() GCRacingWheelInput {
	rv := objc.Send[GCRacingWheelInput](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCRacingWheelInput creates a new GCRacingWheelInput instance.
func NewGCRacingWheelInput() GCRacingWheelInput {
	class := getGCRacingWheelInputClass()
	rv := objc.Send[GCRacingWheelInput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a snapshot of the racing wheel inputs.
//
// SDK narrowing: GCRacingWheelInput declares this as - (GCRacingWheelInputState *)capture.
// Go has no covariance, so the base declaration on GCDevicePhysicalInput is used.
// Convert with GCRacingWheelInputStateFromID(v.GetID()) to recover the narrowed type.
//
// # Return Value
//
// A new instance containing the current state vector of the racing wheel
// input.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheelInput/capture()
func (g GCRacingWheelInput) Capture() GCDevicePhysicalInputState {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("capture"))
	return GCDevicePhysicalInputStateObjectFromID(rv)
}

// Returns the next input state of the racing wheel from the queue.
//
// SDK narrowing: GCRacingWheelInput declares this as - (GCRacingWheelInputState<GCDevicePhysicalInputStateDiff> *)nextInputState.
// Go has no covariance, so the base declaration on GCDevicePhysicalInput is used.
// Convert with GCRacingWheelInputStateFromID(v.GetID()) to recover the narrowed type.
//
// # Return Value
//
// The next input state in the queue or `nil` if the queue is empty.
//
// # Discussion
//
// This method removes the next input state from the queue.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheelInput/nextInputState()
func (g GCRacingWheelInput) NextInputState() interface {
	GCDevicePhysicalInputState
	GCDevicePhysicalInputStateDiff
} {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("nextInputState"))
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

// A block that the profile calls when an element’s value changes.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/elementValueDidChangeHandler
func (g GCRacingWheelInput) ElementValueDidChangeHandler() func(objectivec.IObject) {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("elementValueDidChangeHandler"))
	_ = rv
	return nil
}

// The block that the profile calls when Game Controller adds an input state
// to the queue.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/inputStateAvailableHandler
func (g GCRacingWheelInput) InputStateAvailableHandler() func(objectivec.IObject) {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("inputStateAvailableHandler"))
	_ = rv
	return nil
}

// The maximum number of input values that the queue stores.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/inputStateQueueDepth
func (g GCRacingWheelInput) InputStateQueueDepth() int {
	rv := objc.Send[int](g.ID, objc.Sel("inputStateQueueDepth"))
	return rv
}

// The dispatch queue that the system uses for callbacks.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/queue
func (g GCRacingWheelInput) Queue() dispatch.Queue {
	rv := objc.Send[uintptr](g.ID, objc.Sel("queue"))
	return dispatch.QueueFromHandle(rv)
}

// Protocol methods for GCDevicePhysicalInput

// The maximum number of input values that the queue stores.
//
// # Discussion
//
// When the queue reaches this limit, Game Controller starts removing the
// oldest input states from the queue. The default value for this property is
// `1` which indicates no buffering.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/inputStateQueueDepth
func (o GCRacingWheelInput) SetInputStateQueueDepth(value int) {
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
func (o GCRacingWheelInput) SetQueue(value dispatch.Queue) {
	objc.Send[struct{}](o.ID, objc.Sel("setQueue:"), value)
}

// The device’s elements as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/elements-1shp2
func (o GCRacingWheelInput) Elements() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("elements"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s axes as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/axes-80rx
func (o GCRacingWheelInput) Axes() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("axes"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s buttons as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/buttons-3257g
func (o GCRacingWheelInput) Buttons() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("buttons"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s directional pads as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/dpads-5yr9x
func (o GCRacingWheelInput) Dpads() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("dpads"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s switches as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/switches-6bws2
func (o GCRacingWheelInput) Switches() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("switches"))
	return GCPhysicalInputElementCollectionFromID(rv)
}
