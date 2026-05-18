// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objc"
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
// Use the [GCControllerLiveInput.Capture] method to save a copy of the current input state. If you
// want Game Controller to buffer snapshots of the input states for you, use
// the [GCControllerLiveInput.InputStateQueueDepth] property to set the buffer’s queue depth to a
// value other than `0`. Then use the [GCControllerLiveInput.NextInputState] method to get the
// snapshots when you’re ready to process input.
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
	NextInputState() IGCControllerInputState
	// Returns a snapshot of the physical device inputs.
	Capture() IGCControllerInputState

	// Topic: Remapping controls

	// The live input of a controller without any system-level remapping of the controls.
	UnmappedInput() IGCControllerLiveInput

	// A block that the profile calls when an element’s value changes.
	ElementValueDidChangeHandler() func(objc.ID)
	// The block that the profile calls when Game Controller adds an input state to the queue.
	InputStateAvailableHandler() func(objc.ID)
	// The dispatch queue that the system uses for callbacks.
	Queue() dispatch.Queue
	Elements() IGCPhysicalInputElementCollection
	Axes() IGCPhysicalInputElementCollection
	Buttons() IGCPhysicalInputElementCollection
	Dpads() IGCPhysicalInputElementCollection
	Switches() IGCPhysicalInputElementCollection
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
// # Return Value
//
// The next input state in the queue or `nil` if the queue is empty.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerLiveInput/nextInputState()
func (g GCControllerLiveInput) NextInputState() IGCControllerInputState {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("nextInputState"))
	return GCControllerInputStateFromID(rv)
}

// Returns a snapshot of the physical device inputs.
//
// # Return Value
//
// A new instance containing the current state of the physical device input.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerLiveInput/capture()
func (g GCControllerLiveInput) Capture() IGCControllerInputState {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("capture"))
	return GCControllerInputStateFromID(rv)
}

// A block that the profile calls when an element’s value changes.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/elementValueDidChangeHandler
func (g GCControllerLiveInput) ElementValueDidChangeHandler() func(objc.ID) {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("elementValueDidChangeHandler"))
	_ = rv
	return nil
}

// The block that the profile calls when Game Controller adds an input state
// to the queue.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/inputStateAvailableHandler
func (g GCControllerLiveInput) InputStateAvailableHandler() func(objc.ID) {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("inputStateAvailableHandler"))
	_ = rv
	return nil
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

// The maximum number of input values that the queue stores.
//
// See: https://developer.apple.com/documentation/gamecontroller/gcdevicephysicalinput/inputstatequeuedepth
func (g GCControllerLiveInput) InputStateQueueDepth() int {
	rv := objc.Send[int](g.ID, objc.Sel("inputStateQueueDepth"))
	return rv
}
func (g GCControllerLiveInput) SetInputStateQueueDepth(value int) {
	objc.Send[struct{}](g.ID, objc.Sel("setInputStateQueueDepth:"), value)
}
