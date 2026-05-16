// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [GCControllerButtonInput] class.
var (
	_GCControllerButtonInputClass     GCControllerButtonInputClass
	_GCControllerButtonInputClassOnce sync.Once
)

func getGCControllerButtonInputClass() GCControllerButtonInputClass {
	_GCControllerButtonInputClassOnce.Do(func() {
		_GCControllerButtonInputClass = GCControllerButtonInputClass{class: objc.GetClass("GCControllerButtonInput")}
	})
	return _GCControllerButtonInputClass
}

// GetGCControllerButtonInputClass returns the class object for GCControllerButtonInput.
func GetGCControllerButtonInputClass() GCControllerButtonInputClass {
	return getGCControllerButtonInputClass()
}

type GCControllerButtonInputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCControllerButtonInputClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCControllerButtonInputClass) Alloc() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// A control element that represents a button touch or press.
//
// # Overview
//
// A [GCControllerButtonInput] object represents a button on a controller that
// can report either analog or digital values.
//
// # Accessing input values
//
//   - [GCControllerButtonInput.Touched]: A Boolean value that indicates whether the user is touching the button.
//   - [GCControllerButtonInput.Pressed]: A Boolean value that indicates whether the user is pressing the button.
//   - [GCControllerButtonInput.Value]: The level of pressure the user is applying to the button.
//
// # Getting change information
//
//   - [GCControllerButtonInput.TouchedChangedHandler]: The block that the element calls when the user touches the button.
//   - [GCControllerButtonInput.SetTouchedChangedHandler]
//   - [GCControllerButtonInput.PressedChangedHandler]: The block that the element calls when the user presses or releases the button.
//   - [GCControllerButtonInput.SetPressedChangedHandler]
//   - [GCControllerButtonInput.ValueChangedHandler]: The block that the element calls when the user changes the level of pressure on the button.
//   - [GCControllerButtonInput.SetValueChangedHandler]
//
// See: https://developer.apple.com/documentation/GameController/GCControllerButtonInput
type GCControllerButtonInput struct {
	GCControllerElement
}

// GCControllerButtonInputFromID constructs a [GCControllerButtonInput] from an objc.ID.
//
// A control element that represents a button touch or press.
func GCControllerButtonInputFromID(id objc.ID) GCControllerButtonInput {
	return GCControllerButtonInput{GCControllerElement: GCControllerElementFromID(id)}
}

// NOTE: GCControllerButtonInput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCControllerButtonInput] class.
//
// # Accessing input values
//
//   - [IGCControllerButtonInput.Touched]: A Boolean value that indicates whether the user is touching the button.
//   - [IGCControllerButtonInput.Pressed]: A Boolean value that indicates whether the user is pressing the button.
//   - [IGCControllerButtonInput.Value]: The level of pressure the user is applying to the button.
//
// # Getting change information
//
//   - [IGCControllerButtonInput.TouchedChangedHandler]: The block that the element calls when the user touches the button.
//   - [IGCControllerButtonInput.SetTouchedChangedHandler]
//   - [IGCControllerButtonInput.PressedChangedHandler]: The block that the element calls when the user presses or releases the button.
//   - [IGCControllerButtonInput.SetPressedChangedHandler]
//   - [IGCControllerButtonInput.ValueChangedHandler]: The block that the element calls when the user changes the level of pressure on the button.
//   - [IGCControllerButtonInput.SetValueChangedHandler]
//
// See: https://developer.apple.com/documentation/GameController/GCControllerButtonInput
type IGCControllerButtonInput interface {
	IGCControllerElement

	// Topic: Accessing input values

	// A Boolean value that indicates whether the user is touching the button.
	Touched() bool
	// A Boolean value that indicates whether the user is pressing the button.
	Pressed() bool
	// The level of pressure the user is applying to the button.
	Value() float32

	// Topic: Getting change information

	// The block that the element calls when the user touches the button.
	TouchedChangedHandler() GCControllerButtonTouchedChangedHandler
	SetTouchedChangedHandler(value GCControllerButtonTouchedChangedHandler)
	// The block that the element calls when the user presses or releases the button.
	PressedChangedHandler() GCControllerButtonValueChangedHandler
	SetPressedChangedHandler(value GCControllerButtonValueChangedHandler)
	// The block that the element calls when the user changes the level of pressure on the button.
	ValueChangedHandler() GCControllerButtonValueChangedHandler
	SetValueChangedHandler(value GCControllerButtonValueChangedHandler)
}

// Init initializes the instance.
func (g GCControllerButtonInput) Init() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCControllerButtonInput) Autorelease() GCControllerButtonInput {
	rv := objc.Send[GCControllerButtonInput](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCControllerButtonInput creates a new GCControllerButtonInput instance.
func NewGCControllerButtonInput() GCControllerButtonInput {
	class := getGCControllerButtonInputClass()
	rv := objc.Send[GCControllerButtonInput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value that indicates whether the user is touching the button.
//
// # Discussion
//
// If this property is true, the user is touching the button; otherwise, the
// user isn’t. For controllers that support capacitive touch, the user can
// start touching the button without pressure when the value property is `0`.
// For controllers that don’t support capacitive touch, the user starts
// touching the button when the value property is greater than `0`.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerButtonInput/isTouched
func (g GCControllerButtonInput) Touched() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isTouched"))
	return rv
}

// A Boolean value that indicates whether the user is pressing the button.
//
// # Discussion
//
// If this property is true, the user is putting pressure on the button;
// otherwise, the user isn’t.
//
// For the DualSense, DualShock 4, and Siri Remote controllers, the framework
// simulates whether the user presses the button and the level of pressure for
// its touch surfaces.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerButtonInput/isPressed
func (g GCControllerButtonInput) Pressed() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isPressed"))
	return rv
}

// The level of pressure the user is applying to the button.
//
// # Discussion
//
// If the user applies pressure to the button, the [Pressed] property is true
// and this property indicates the amount of pressure. The framework
// normalizes the value to a number between `0.0` (minimum) and `1.0`
// (maximum). If the user isn’t pressing the button, the [Pressed] property
// is false and this property is `0.0`.
//
// For axis buttons, such as thumbsticks and touchpads, the location on the
// positive or negative axis of the element simulates the pressure.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerButtonInput/value
func (g GCControllerButtonInput) Value() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("value"))
	return rv
}

// The block that the element calls when the user touches the button.
//
// # Discussion
//
// Set this handler when you want to know when the user touches the button
// before pressing the button.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerButtonInput/touchedChangedHandler
func (g GCControllerButtonInput) TouchedChangedHandler() GCControllerButtonTouchedChangedHandler {
	rv := objc.Send[GCControllerButtonTouchedChangedHandler](g.ID, objc.Sel("touchedChangedHandler"))
	return GCControllerButtonTouchedChangedHandler(rv)
}
func (g GCControllerButtonInput) SetTouchedChangedHandler(value GCControllerButtonTouchedChangedHandler) {
	objc.Send[struct{}](g.ID, objc.Sel("setTouchedChangedHandler:"), value)
}

// The block that the element calls when the user presses or releases the
// button.
//
// # Discussion
//
// Set this handler when you only want to know when the user presses or
// releases the button — that is, when the [Pressed] property changes.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerButtonInput/pressedChangedHandler
func (g GCControllerButtonInput) PressedChangedHandler() GCControllerButtonValueChangedHandler {
	rv := objc.Send[GCControllerButtonValueChangedHandler](g.ID, objc.Sel("pressedChangedHandler"))
	return GCControllerButtonValueChangedHandler(rv)
}
func (g GCControllerButtonInput) SetPressedChangedHandler(value GCControllerButtonValueChangedHandler) {
	objc.Send[struct{}](g.ID, objc.Sel("setPressedChangedHandler:"), value)
}

// The block that the element calls when the user changes the level of
// pressure on the button.
//
// # Discussion
//
// Set this handler when you want to know when the pressure level changes.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerButtonInput/valueChangedHandler
func (g GCControllerButtonInput) ValueChangedHandler() GCControllerButtonValueChangedHandler {
	rv := objc.Send[GCControllerButtonValueChangedHandler](g.ID, objc.Sel("valueChangedHandler"))
	return GCControllerButtonValueChangedHandler(rv)
}
func (g GCControllerButtonInput) SetValueChangedHandler(value GCControllerButtonValueChangedHandler) {
	objc.Send[struct{}](g.ID, objc.Sel("setValueChangedHandler:"), value)
}
