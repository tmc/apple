// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GCRacingWheelInputState] class.
var (
	_GCRacingWheelInputStateClass     GCRacingWheelInputStateClass
	_GCRacingWheelInputStateClassOnce sync.Once
)

func getGCRacingWheelInputStateClass() GCRacingWheelInputStateClass {
	_GCRacingWheelInputStateClassOnce.Do(func() {
		_GCRacingWheelInputStateClass = GCRacingWheelInputStateClass{class: objc.GetClass("GCRacingWheelInputState")}
	})
	return _GCRacingWheelInputStateClass
}

// GetGCRacingWheelInputStateClass returns the class object for GCRacingWheelInputState.
func GetGCRacingWheelInputStateClass() GCRacingWheelInputStateClass {
	return getGCRacingWheelInputStateClass()
}

type GCRacingWheelInputStateClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCRacingWheelInputStateClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCRacingWheelInputStateClass) Alloc() GCRacingWheelInputState {
	rv := objc.Send[GCRacingWheelInputState](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// The input for the wheel of a racing wheel controller.
//
// # Getting input elements
//
//   - [GCRacingWheelInputState.Wheel]: The controller’s wheel element.
//   - [GCRacingWheelInputState.AcceleratorPedal]: The controller’s accelerator pedal element.
//   - [GCRacingWheelInputState.BrakePedal]: The controller’s brake pedal element.
//   - [GCRacingWheelInputState.ClutchPedal]: The controller’s clutch element.
//   - [GCRacingWheelInputState.Shifter]: The controller’s gear shift element.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheelInputState
type GCRacingWheelInputState struct {
	objectivec.Object
}

// GCRacingWheelInputStateFromID constructs a [GCRacingWheelInputState] from an objc.ID.
//
// The input for the wheel of a racing wheel controller.
func GCRacingWheelInputStateFromID(id objc.ID) GCRacingWheelInputState {
	return GCRacingWheelInputState{objectivec.Object{ID: id}}
}

// NOTE: GCRacingWheelInputState adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCRacingWheelInputState] class.
//
// # Getting input elements
//
//   - [IGCRacingWheelInputState.Wheel]: The controller’s wheel element.
//   - [IGCRacingWheelInputState.AcceleratorPedal]: The controller’s accelerator pedal element.
//   - [IGCRacingWheelInputState.BrakePedal]: The controller’s brake pedal element.
//   - [IGCRacingWheelInputState.ClutchPedal]: The controller’s clutch element.
//   - [IGCRacingWheelInputState.Shifter]: The controller’s gear shift element.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheelInputState
type IGCRacingWheelInputState interface {
	objectivec.IObject

	// Topic: Getting input elements

	// The controller’s wheel element.
	Wheel() IGCSteeringWheelElement
	// The controller’s accelerator pedal element.
	AcceleratorPedal() GCButtonElement
	// The controller’s brake pedal element.
	BrakePedal() GCButtonElement
	// The controller’s clutch element.
	ClutchPedal() GCButtonElement
	// The controller’s gear shift element.
	Shifter() IGCGearShifterElement
}

// Init initializes the instance.
func (g GCRacingWheelInputState) Init() GCRacingWheelInputState {
	rv := objc.Send[GCRacingWheelInputState](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCRacingWheelInputState) Autorelease() GCRacingWheelInputState {
	rv := objc.Send[GCRacingWheelInputState](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCRacingWheelInputState creates a new GCRacingWheelInputState instance.
func NewGCRacingWheelInputState() GCRacingWheelInputState {
	class := getGCRacingWheelInputStateClass()
	rv := objc.Send[GCRacingWheelInputState](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The physical device that this profile represents.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/device
func (g GCRacingWheelInputState) Device() GCDevice {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("device"))
	return GCDeviceObjectFromID(rv)
}

// The time in seconds between the last event and the current time.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/lastEventLatency
func (g GCRacingWheelInputState) LastEventLatency() foundation.NSTimeInterval {
	rv := objc.Send[foundation.NSTimeInterval](g.ID, objc.Sel("lastEventLatency"))
	return foundation.NSTimeInterval(rv)
}

// The time of the most recent event.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/lastEventTimestamp
func (g GCRacingWheelInputState) LastEventTimestamp() foundation.NSTimeInterval {
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
func (g GCRacingWheelInputState) ObjectForKeyedSubscript(key string) GCPhysicalInputElement {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("objectForKeyedSubscript:"), objc.String(key))
	return GCPhysicalInputElementObjectFromID(rv)
}

// The controller’s wheel element.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheelInputState/wheel
func (g GCRacingWheelInputState) Wheel() IGCSteeringWheelElement {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("wheel"))
	return GCSteeringWheelElementFromID(objc.ID(rv))
}

// The controller’s accelerator pedal element.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheelInputState/acceleratorPedal
func (g GCRacingWheelInputState) AcceleratorPedal() GCButtonElement {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("acceleratorPedal"))
	return GCButtonElementObjectFromID(rv)
}

// The controller’s brake pedal element.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheelInputState/brakePedal
func (g GCRacingWheelInputState) BrakePedal() GCButtonElement {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("brakePedal"))
	return GCButtonElementObjectFromID(rv)
}

// The controller’s clutch element.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheelInputState/clutchPedal
func (g GCRacingWheelInputState) ClutchPedal() GCButtonElement {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("clutchPedal"))
	return GCButtonElementObjectFromID(rv)
}

// The controller’s gear shift element.
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheelInputState/shifter
func (g GCRacingWheelInputState) Shifter() IGCGearShifterElement {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("shifter"))
	return GCGearShifterElementFromID(objc.ID(rv))
}

// Protocol methods for GCDevicePhysicalInputState

// The device’s elements as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/elements-1shp2
func (o GCRacingWheelInputState) Elements() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("elements"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s axes as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/axes-80rx
func (o GCRacingWheelInputState) Axes() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("axes"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s buttons as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/buttons-3257g
func (o GCRacingWheelInputState) Buttons() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("buttons"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s directional pads as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/dpads-5yr9x
func (o GCRacingWheelInputState) Dpads() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("dpads"))
	return GCPhysicalInputElementCollectionFromID(rv)
}

// The device’s switches as key-value pairs for lookup by name.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/switches-6bws2
func (o GCRacingWheelInputState) Switches() IGCPhysicalInputElementCollection {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("switches"))
	return GCPhysicalInputElementCollectionFromID(rv)
}
