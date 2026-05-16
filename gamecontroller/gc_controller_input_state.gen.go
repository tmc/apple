// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GCControllerInputState] class.
var (
	_GCControllerInputStateClass     GCControllerInputStateClass
	_GCControllerInputStateClassOnce sync.Once
)

func getGCControllerInputStateClass() GCControllerInputStateClass {
	_GCControllerInputStateClassOnce.Do(func() {
		_GCControllerInputStateClass = GCControllerInputStateClass{class: objc.GetClass("GCControllerInputState")}
	})
	return _GCControllerInputStateClass
}

// GetGCControllerInputStateClass returns the class object for GCControllerInputState.
func GetGCControllerInputStateClass() GCControllerInputStateClass {
	return getGCControllerInputStateClass()
}

type GCControllerInputStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCControllerInputStateClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCControllerInputStateClass) Alloc() GCControllerInputState {
	rv := objc.Send[GCControllerInputState](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// A class that represents an input state for gamepads and arcade sticks.
//
// # Overview
//
// This class implements the [GCDevicePhysicalInputState] protocol for
// gamepads and arcade sticks. Instances of this class represent the state of
// the controller’s inputs at a moment in time, which can be the current
// time.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerInputState
type GCControllerInputState struct {
	objectivec.Object
}

// GCControllerInputStateFromID constructs a [GCControllerInputState] from an objc.ID.
//
// A class that represents an input state for gamepads and arcade sticks.
func GCControllerInputStateFromID(id objc.ID) GCControllerInputState {
	return GCControllerInputState{objectivec.Object{ID: id}}
}

// NOTE: GCControllerInputState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCControllerInputState] class.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerInputState
type IGCControllerInputState interface {
	objectivec.IObject
	GCDevicePhysicalInputState

	// The input profile for the controller.
	Input() IGCControllerLiveInput
	SetInput(value IGCControllerLiveInput)
}

// Init initializes the instance.
func (g GCControllerInputState) Init() GCControllerInputState {
	rv := objc.Send[GCControllerInputState](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCControllerInputState) Autorelease() GCControllerInputState {
	rv := objc.Send[GCControllerInputState](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCControllerInputState creates a new GCControllerInputState instance.
func NewGCControllerInputState() GCControllerInputState {
	class := getGCControllerInputStateClass()
	rv := objc.Send[GCControllerInputState](objc.ID(class.class), objc.Sel("new"))
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
func (g GCControllerInputState) ObjectForKeyedSubscript(key string) GCPhysicalInputElement {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("objectForKeyedSubscript:"), objc.String(key))
	return GCPhysicalInputElementObjectFromID(rv)
}

// The physical device that this profile represents.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/device
func (g GCControllerInputState) Device() GCDevice {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("device"))
	return GCDeviceObjectFromID(rv)
}

// The input profile for the controller.
//
// See: https://developer.apple.com/documentation/gamecontroller/gccontroller/input
func (g GCControllerInputState) Input() IGCControllerLiveInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("input"))
	return GCControllerLiveInputFromID(objc.ID(rv))
}
func (g GCControllerInputState) SetInput(value IGCControllerLiveInput) {
	objc.Send[struct{}](g.ID, objc.Sel("setInput:"), value)
}

// The time in seconds between the last event and the current time.
//
// # Discussion
//
// Use this property as a minimum latency value that may not include latency
// that accrues on the device or when it transmits the event. If the host goes
// to sleep between when the event occurs and when you get this property, the
// value may not be accurate.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/lastEventLatency
func (g GCControllerInputState) LastEventLatency() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("lastEventLatency"))
	return rv
}

// The time of the most recent event.
//
// # Discussion
//
// This property isn’t relative to any specific date and time. To determine
// the time between events, subtract a previous value of this property from
// the current value. You can also compare [LastEventTimestamp] properties of
// two different devices to determine which event occurs first.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/lastEventTimestamp
func (g GCControllerInputState) LastEventTimestamp() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("lastEventTimestamp"))
	return rv
}

// Protocol methods for GCDevicePhysicalInputState
