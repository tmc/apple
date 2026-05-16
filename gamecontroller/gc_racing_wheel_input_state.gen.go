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
// # Accessing elements by name
//
//   - [GCRacingWheelInputState.GCInputSteeringWheel]: The name of the steering wheel element.
//   - [GCRacingWheelInputState.SetGCInputSteeringWheel]
//   - [GCRacingWheelInputState.GCInputShifter]: The name of the shifter element.
//   - [GCRacingWheelInputState.SetGCInputShifter]
//   - [GCRacingWheelInputState.GCInputPedalClutch]: The name of the clutch element.
//   - [GCRacingWheelInputState.SetGCInputPedalClutch]
//   - [GCRacingWheelInputState.GCInputPedalAccelerator]: The name of the accelerator element.
//   - [GCRacingWheelInputState.SetGCInputPedalAccelerator]
//   - [GCRacingWheelInputState.GCInputPedalBrake]: The name of the brake element.
//   - [GCRacingWheelInputState.SetGCInputPedalBrake]
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
// # Accessing elements by name
//
//   - [IGCRacingWheelInputState.GCInputSteeringWheel]: The name of the steering wheel element.
//   - [IGCRacingWheelInputState.SetGCInputSteeringWheel]
//   - [IGCRacingWheelInputState.GCInputShifter]: The name of the shifter element.
//   - [IGCRacingWheelInputState.SetGCInputShifter]
//   - [IGCRacingWheelInputState.GCInputPedalClutch]: The name of the clutch element.
//   - [IGCRacingWheelInputState.SetGCInputPedalClutch]
//   - [IGCRacingWheelInputState.GCInputPedalAccelerator]: The name of the accelerator element.
//   - [IGCRacingWheelInputState.SetGCInputPedalAccelerator]
//   - [IGCRacingWheelInputState.GCInputPedalBrake]: The name of the brake element.
//   - [IGCRacingWheelInputState.SetGCInputPedalBrake]
//
// See: https://developer.apple.com/documentation/GameController/GCRacingWheelInputState
type IGCRacingWheelInputState interface {
	objectivec.IObject
	GCDevicePhysicalInputState

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

	// Topic: Accessing elements by name

	// The name of the steering wheel element.
	GCInputSteeringWheel() string
	SetGCInputSteeringWheel(value string)
	// The name of the shifter element.
	GCInputShifter() string
	SetGCInputShifter(value string)
	// The name of the clutch element.
	GCInputPedalClutch() string
	SetGCInputPedalClutch(value string)
	// The name of the accelerator element.
	GCInputPedalAccelerator() string
	SetGCInputPedalAccelerator(value string)
	// The name of the brake element.
	GCInputPedalBrake() string
	SetGCInputPedalBrake(value string)
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

// The name of the steering wheel element.
//
// See: https://developer.apple.com/documentation/gamecontroller/gcinputsteeringwheel-26283
func (g GCRacingWheelInputState) GCInputSteeringWheel() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("GCInputSteeringWheel"))
	return foundation.NSStringFromID(rv).String()
}
func (g GCRacingWheelInputState) SetGCInputSteeringWheel(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setGCInputSteeringWheel:"), objc.String(value))
}

// The name of the shifter element.
//
// See: https://developer.apple.com/documentation/gamecontroller/gcinputshifter-6miga
func (g GCRacingWheelInputState) GCInputShifter() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("GCInputShifter"))
	return foundation.NSStringFromID(rv).String()
}
func (g GCRacingWheelInputState) SetGCInputShifter(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setGCInputShifter:"), objc.String(value))
}

// The name of the clutch element.
//
// See: https://developer.apple.com/documentation/gamecontroller/gcinputpedalclutch-82gwe
func (g GCRacingWheelInputState) GCInputPedalClutch() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("GCInputPedalClutch"))
	return foundation.NSStringFromID(rv).String()
}
func (g GCRacingWheelInputState) SetGCInputPedalClutch(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setGCInputPedalClutch:"), objc.String(value))
}

// The name of the accelerator element.
//
// See: https://developer.apple.com/documentation/gamecontroller/gcinputpedalaccelerator-6kg6u
func (g GCRacingWheelInputState) GCInputPedalAccelerator() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("GCInputPedalAccelerator"))
	return foundation.NSStringFromID(rv).String()
}
func (g GCRacingWheelInputState) SetGCInputPedalAccelerator(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setGCInputPedalAccelerator:"), objc.String(value))
}

// The name of the brake element.
//
// See: https://developer.apple.com/documentation/gamecontroller/gcinputpedalbrake-6wpdc
func (g GCRacingWheelInputState) GCInputPedalBrake() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("GCInputPedalBrake"))
	return foundation.NSStringFromID(rv).String()
}
func (g GCRacingWheelInputState) SetGCInputPedalBrake(value string) {
	objc.Send[struct{}](g.ID, objc.Sel("setGCInputPedalBrake:"), objc.String(value))
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
// # Discussion
//
// Use this property as a minimum latency value that may not include latency
// that accrues on the device or when it transmits the event. If the host goes
// to sleep between when the event occurs and when you get this property, the
// value may not be accurate.
//
// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputState/lastEventLatency
func (g GCRacingWheelInputState) LastEventLatency() float64 {
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
func (g GCRacingWheelInputState) LastEventTimestamp() float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("lastEventTimestamp"))
	return rv
}

// Protocol methods for GCDevicePhysicalInputState
