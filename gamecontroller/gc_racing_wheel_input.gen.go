// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/objc"
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
	Capture() IGCRacingWheelInputState

	// Topic: Polling for input

	// Returns the next input state of the racing wheel from the queue.
	NextInputState() IGCRacingWheelInputState
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
// # Return Value
//
// A new instance containing the current state vector of the racing wheel
// input.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheelInput/capture()
func (g GCRacingWheelInput) Capture() IGCRacingWheelInputState {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("capture"))
	return GCRacingWheelInputStateFromID(rv)
}

// Returns the next input state of the racing wheel from the queue.
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
func (g GCRacingWheelInput) NextInputState() IGCRacingWheelInputState {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("nextInputState"))
	return GCRacingWheelInputStateFromID(rv)
}

// A block that the profile calls when an element’s value changes.
//
// # Discussion
//
// Use this property to get the latest state of the element. If multiple
// elements change, Game Controller invokes this block for each element that
// changes. The block’s parameters are:
//
// element: The element whose value changes.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/elementValueDidChangeHandler
func (g GCRacingWheelInput) ElementValueDidChangeHandler() VoidHandler {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("elementValueDidChangeHandler"))
	_ = rv
	return nil
}
func (g GCRacingWheelInput) SetElementValueDidChangeHandler(value VoidHandler) {
	block, cleanup := NewVoidBlock(value)
	defer cleanup()
	objc.Send[struct{}](g.ID, objc.Sel("setElementValueDidChangeHandler:"), block)
}

// The block that the profile calls when Game Controller adds an input state
// to the queue.
//
// # Discussion
//
// Set this property to track every element value change, not just the current
// value. When Game Controller invokes the handler, invoke the
// [NextInputState] method repeatedly to get all the buffered changes until
// the queue is empty.
//
// To get just the current element value, use the
// [ElementValueDidChangeHandler] property instead.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInput/inputStateAvailableHandler
func (g GCRacingWheelInput) InputStateAvailableHandler() GCDevicePhysicalInputHandler {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("inputStateAvailableHandler"))
	_ = rv
	return nil
}
func (g GCRacingWheelInput) SetInputStateAvailableHandler(value GCDevicePhysicalInputHandler) {
	block, cleanup := NewGCDevicePhysicalInputBlock(value)
	defer cleanup()
	objc.Send[struct{}](g.ID, objc.Sel("setInputStateAvailableHandler:"), block)
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
func (g GCRacingWheelInput) InputStateQueueDepth() int {
	rv := objc.Send[int](g.ID, objc.Sel("inputStateQueueDepth"))
	return rv
}
func (g GCRacingWheelInput) SetInputStateQueueDepth(value int) {
	objc.Send[struct{}](g.ID, objc.Sel("setInputStateQueueDepth:"), value)
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
func (g GCRacingWheelInput) Queue() dispatch.Queue {
	rv := objc.Send[uintptr](g.ID, objc.Sel("queue"))
	return dispatch.QueueFromHandle(rv)
}
func (g GCRacingWheelInput) SetQueue(value dispatch.Queue) {
	objc.Send[struct{}](g.ID, objc.Sel("setQueue:"), uintptr(value.Handle()))
}
