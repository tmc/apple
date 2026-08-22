// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/foundation"
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

	// The physical device that this profile represents.
	Device() GCDevice
	// The time in seconds between the last event and the current time.
	LastEventLatency() foundation.NSTimeInterval
	// The time of the most recent event.
	LastEventTimestamp() foundation.NSTimeInterval
	// Returns the element that the key specifies.
	ObjectForKeyedSubscript(key string) GCPhysicalInputElement
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

// The physical device that this profile represents.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/device
func (g GCControllerInputState) Device() GCDevice {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("device"))
	return GCDeviceObjectFromID(rv)
}

// The time in seconds between the last event and the current time.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/lastEventLatency
func (g GCControllerInputState) LastEventLatency() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](g.ID, objc.Sel("lastEventLatency"))
	return foundation.NSTimeInterval(rv)
}

// The time of the most recent event.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/lastEventTimestamp
func (g GCControllerInputState) LastEventTimestamp() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](g.ID, objc.Sel("lastEventTimestamp"))
	return foundation.NSTimeInterval(rv)
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

// Protocol methods for GCDevicePhysicalInputState

// The device’s elements as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/elements-1shp2
func (o GCControllerInputState) Elements() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("elements"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s axes as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/axes-80rx
func (o GCControllerInputState) Axes() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("axes"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s buttons as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/buttons-3257g
func (o GCControllerInputState) Buttons() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("buttons"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s directional pads as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/dpads-5yr9x
func (o GCControllerInputState) Dpads() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("dpads"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s switches as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/switches-6bws2
func (o GCControllerInputState) Switches() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("switches"))
	return GCPhysicalInputElementCollectionFromID(rv)
}
