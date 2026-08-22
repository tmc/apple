// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GCControllerLiveInput] class.
var (
	_GCControllerLiveInputClass     GCControllerLiveInputClass
	_GCControllerLiveInputClassOnce sync.Once
)

func getGCControllerLiveInputClass() GCControllerLiveInputClass {
	_GCControllerLiveInputClassOnce.Do(func() {
		_GCControllerLiveInputClass = GCControllerLiveInputClass{class: objc.GetClass("GCControllerLiveInput")}
	})
	return _GCControllerLiveInputClass
}

// GetGCControllerLiveInputClass returns the class object for GCControllerLiveInput.
func GetGCControllerLiveInputClass() GCControllerLiveInputClass {
	return getGCControllerLiveInputClass()
}

type GCControllerLiveInputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCControllerLiveInputClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCControllerLiveInputClass) Alloc() GCControllerLiveInput {
	rv := objc.Send[GCControllerLiveInput](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// The input profile for a controller.
//
// # Overview
//
// Instances of [GCControllerLiveInput] represent the current input state of a
// controller. You can save snapshots of the input state and receive callbacks
// when the input state changes. You can also get the elements of the
// controller and their current input values from [GCControllerLiveInput]
// instances.
//
// Use the [GCControllerLiveInput.Capture] method to save a copy of the
// current input state. If you want Game Controller to buffer snapshots of the
// input states for you, use the [GCControllerLiveInput.InputStateQueueDepth] property to set the
// buffer’s queue depth to a value other than `0`. Then use the
// [GCControllerLiveInput.NextInputState] method to get the snapshots when
// you’re ready to process input.
//
// # Handling device input
//
//   - [GCControllerLiveInput.NextInputState]: Returns the next device input state from the queue.
//   - [GCControllerLiveInput.Capture]: Returns a snapshot of the physical device inputs.
//
// # Remapping controls
//
//   - [GCControllerLiveInput.UnmappedInput]: The live input of a controller without any system-level remapping of the controls.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerLiveInput
type GCControllerLiveInput struct {
	GCControllerInputState
}

// GCControllerLiveInputFromID constructs a [GCControllerLiveInput] from an objc.ID.
//
// The input profile for a controller.
func GCControllerLiveInputFromID(id objc.ID) GCControllerLiveInput {
	return GCControllerLiveInput{GCControllerInputState: GCControllerInputStateFromID(id)}
}

// NOTE: GCControllerLiveInput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCControllerLiveInput] class.
//
// # Handling device input
//
//   - [IGCControllerLiveInput.NextInputState]: Returns the next device input state from the queue.
//   - [IGCControllerLiveInput.Capture]: Returns a snapshot of the physical device inputs.
//
// # Remapping controls
//
//   - [IGCControllerLiveInput.UnmappedInput]: The live input of a controller without any system-level remapping of the controls.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerLiveInput
type IGCControllerLiveInput interface {
	IGCControllerInputState

	// Topic: Handling device input

	// Returns the next device input state from the queue.
	NextInputState() interface {
		GCDevicePhysicalInputState
		GCDevicePhysicalInputStateDiff
	}
	// Returns a snapshot of the physical device inputs.
	Capture() GCDevicePhysicalInputState

	// Topic: Remapping controls

	// The live input of a controller without any system-level remapping of the controls.
	UnmappedInput() IGCControllerLiveInput

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
func (g GCControllerLiveInput) Init() GCControllerLiveInput {
	rv := objc.Send[GCControllerLiveInput](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCControllerLiveInput) Autorelease() GCControllerLiveInput {
	rv := objc.Send[GCControllerLiveInput](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCControllerLiveInput creates a new GCControllerLiveInput instance.
func NewGCControllerLiveInput() GCControllerLiveInput {
	class := getGCControllerLiveInputClass()
	rv := objc.Send[GCControllerLiveInput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the next device input state from the queue.
//
// SDK narrowing: GCControllerLiveInput declares this as - (GCControllerInputState<GCDevicePhysicalInputStateDiff> *)nextInputState.
// Go has no covariance, so the base declaration on GCDevicePhysicalInput is used.
// Convert with GCControllerInputStateFromID(v.GetID()) to recover the narrowed type.
//
// # Return Value
//
// The next input state in the queue or `nil` if the queue is empty.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerLiveInput/nextInputState()
func (g GCControllerLiveInput) NextInputState() interface {
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

// Returns a snapshot of the physical device inputs.
//
// SDK narrowing: GCControllerLiveInput declares this as - (GCControllerInputState *)capture.
// Go has no covariance, so the base declaration on GCDevicePhysicalInput is used.
// Convert with GCControllerInputStateFromID(v.GetID()) to recover the narrowed type.
//
// # Return Value
//
// A new instance containing the current state of the physical device input.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerLiveInput/capture()
func (g GCControllerLiveInput) Capture() GCDevicePhysicalInputState {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("capture"))
	return GCDevicePhysicalInputStateObjectFromID(rv)
}

// A block that the profile calls when an element’s value changes.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/elementValueDidChangeHandler
func (g GCControllerLiveInput) ElementValueDidChangeHandler() func(objectivec.IObject) {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("elementValueDidChangeHandler"))
	_ = rv
	return nil
}

// The block that the profile calls when Game Controller adds an input state
// to the queue.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/inputStateAvailableHandler
func (g GCControllerLiveInput) InputStateAvailableHandler() func(objectivec.IObject) {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("inputStateAvailableHandler"))
	_ = rv
	return nil
}

// The maximum number of input values that the queue stores.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/inputStateQueueDepth
func (g GCControllerLiveInput) InputStateQueueDepth() int {
	rv := objc.Send[int](g.ID, objc.Sel("inputStateQueueDepth"))
	return rv
}

// The dispatch queue that the system uses for callbacks.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/queue
func (g GCControllerLiveInput) Queue() dispatch.Queue {
	rv := objc.Send[uintptr](g.ID, objc.Sel("queue"))
	return dispatch.QueueFromHandle(rv)
}

// The live input of a controller without any system-level remapping of the
// controls.
//
// # Discussion
//
// Players should use the system game controller settings to remap controls.
// If you implement your own controller remapping feature, use this property
// to access the controller’s physical input without remapping applied.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerLiveInput/unmapped
func (g GCControllerLiveInput) UnmappedInput() IGCControllerLiveInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("unmappedInput"))
	return GCControllerLiveInputFromID(objc.ID(rv))
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
func (o GCControllerLiveInput) SetInputStateQueueDepth(value int) {
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
func (o GCControllerLiveInput) SetQueue(value dispatch.Queue) {
	objc.Send[struct{}](o.ID, objc.Sel("setQueue:"), value)
}

// The device’s elements as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/elements-1shp2
func (o GCControllerLiveInput) Elements() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("elements"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s axes as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/axes-80rx
func (o GCControllerLiveInput) Axes() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("axes"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s buttons as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/buttons-3257g
func (o GCControllerLiveInput) Buttons() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("buttons"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s directional pads as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/dpads-5yr9x
func (o GCControllerLiveInput) Dpads() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("dpads"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s switches as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/switches-6bws2
func (o GCControllerLiveInput) Switches() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("switches"))
	return GCPhysicalInputElementCollectionFromID(rv)
}
