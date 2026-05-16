// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [GCMicroGamepad] class.
var (
	_GCMicroGamepadClass     GCMicroGamepadClass
	_GCMicroGamepadClassOnce sync.Once
)

func getGCMicroGamepadClass() GCMicroGamepadClass {
	_GCMicroGamepadClassOnce.Do(func() {
		_GCMicroGamepadClass = GCMicroGamepadClass{class: objc.GetClass("GCMicroGamepad")}
	})
	return _GCMicroGamepadClass
}

// GetGCMicroGamepadClass returns the class object for GCMicroGamepad.
func GetGCMicroGamepadClass() GCMicroGamepadClass {
	return getGCMicroGamepadClass()
}

type GCMicroGamepadClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCMicroGamepadClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCMicroGamepadClass) Alloc() GCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// A controller profile that supports the Siri Remote.
//
// # Overview
//
// The micro gamepad controller profile supports the following input elements:
//
// - Two digital face buttons (A and X). - One analog directional pad (D-pad)
// that functions as a touchpad.
//
// Users can rotate game controllers that support the micro gamepad profile,
// switching them between landscape and portrait orientation. If you want to
// get directional values according to the orientation, set the
// [GCMicroGamepad.AllowsRotation] property to true.
//
// [media-3830807]
//
// # Getting the controller
//
//   - [GCMicroGamepad.Controller]: The controller associated with this profile.
//
// # Receiving a callback when input values change
//
//   - [GCMicroGamepad.ValueChangedHandler]: The block that this profile calls when an element’s value changes.
//   - [GCMicroGamepad.SetValueChangedHandler]
//
// # Getting face button inputs
//
//   - [GCMicroGamepad.ButtonMenu]: The menu face button that players use to enter the main menu and pause the game.
//   - [GCMicroGamepad.ButtonA]: The button that the user activates by pressing harder on the touchpad.
//   - [GCMicroGamepad.ButtonX]: The second face button element.
//
// # Getting directional pad inputs
//
//   - [GCMicroGamepad.Dpad]: The controller’s directional pad element.
//   - [GCMicroGamepad.ReportsAbsoluteDpadValues]: A Boolean value that indicates whether the directional pad reports absolute or relative values.
//   - [GCMicroGamepad.SetReportsAbsoluteDpadValues]
//   - [GCMicroGamepad.AllowsRotation]: A Boolean value that indicates whether the profile reports the directional pad values relative to its current orientation.
//   - [GCMicroGamepad.SetAllowsRotation]
//
// # Setting snapshot avlues
//
//   - [GCMicroGamepad.SetStateFromMicroGamepad]: Copies the input values from a specified micro gamepad to a snapshot of a micro gamepad.
//
// See: https://developer.apple.com/documentation/GameController/GCMicroGamepad
type GCMicroGamepad struct {
	GCPhysicalInputProfile
}

// GCMicroGamepadFromID constructs a [GCMicroGamepad] from an objc.ID.
//
// A controller profile that supports the Siri Remote.
func GCMicroGamepadFromID(id objc.ID) GCMicroGamepad {
	return GCMicroGamepad{GCPhysicalInputProfile: GCPhysicalInputProfileFromID(id)}
}

// NOTE: GCMicroGamepad adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCMicroGamepad] class.
//
// # Getting the controller
//
//   - [IGCMicroGamepad.Controller]: The controller associated with this profile.
//
// # Receiving a callback when input values change
//
//   - [IGCMicroGamepad.ValueChangedHandler]: The block that this profile calls when an element’s value changes.
//   - [IGCMicroGamepad.SetValueChangedHandler]
//
// # Getting face button inputs
//
//   - [IGCMicroGamepad.ButtonMenu]: The menu face button that players use to enter the main menu and pause the game.
//   - [IGCMicroGamepad.ButtonA]: The button that the user activates by pressing harder on the touchpad.
//   - [IGCMicroGamepad.ButtonX]: The second face button element.
//
// # Getting directional pad inputs
//
//   - [IGCMicroGamepad.Dpad]: The controller’s directional pad element.
//   - [IGCMicroGamepad.ReportsAbsoluteDpadValues]: A Boolean value that indicates whether the directional pad reports absolute or relative values.
//   - [IGCMicroGamepad.SetReportsAbsoluteDpadValues]
//   - [IGCMicroGamepad.AllowsRotation]: A Boolean value that indicates whether the profile reports the directional pad values relative to its current orientation.
//   - [IGCMicroGamepad.SetAllowsRotation]
//
// # Setting snapshot avlues
//
//   - [IGCMicroGamepad.SetStateFromMicroGamepad]: Copies the input values from a specified micro gamepad to a snapshot of a micro gamepad.
//
// See: https://developer.apple.com/documentation/GameController/GCMicroGamepad
type IGCMicroGamepad interface {
	IGCPhysicalInputProfile

	// Topic: Getting the controller

	// The controller associated with this profile.
	Controller() IGCController

	// Topic: Receiving a callback when input values change

	// The block that this profile calls when an element’s value changes.
	ValueChangedHandler() GCMicroGamepadValueChangedHandler
	SetValueChangedHandler(value GCMicroGamepadValueChangedHandler)

	// Topic: Getting face button inputs

	// The menu face button that players use to enter the main menu and pause the game.
	ButtonMenu() IGCControllerButtonInput
	// The button that the user activates by pressing harder on the touchpad.
	ButtonA() IGCControllerButtonInput
	// The second face button element.
	ButtonX() IGCControllerButtonInput

	// Topic: Getting directional pad inputs

	// The controller’s directional pad element.
	Dpad() IGCControllerDirectionPad
	// A Boolean value that indicates whether the directional pad reports absolute or relative values.
	ReportsAbsoluteDpadValues() bool
	SetReportsAbsoluteDpadValues(value bool)
	// A Boolean value that indicates whether the profile reports the directional pad values relative to its current orientation.
	AllowsRotation() bool
	SetAllowsRotation(value bool)

	// Topic: Setting snapshot avlues

	// Copies the input values from a specified micro gamepad to a snapshot of a micro gamepad.
	SetStateFromMicroGamepad(microGamepad IGCMicroGamepad)
}

// Init initializes the instance.
func (g GCMicroGamepad) Init() GCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCMicroGamepad) Autorelease() GCMicroGamepad {
	rv := objc.Send[GCMicroGamepad](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCMicroGamepad creates a new GCMicroGamepad instance.
func NewGCMicroGamepad() GCMicroGamepad {
	class := getGCMicroGamepadClass()
	rv := objc.Send[GCMicroGamepad](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Copies the input values from a specified micro gamepad to a snapshot of a
// micro gamepad.
//
// microGamepad: The micro gamepad to copy the input values from.
//
// See: https://developer.apple.com/documentation/GameController/GCMicroGamepad/setStateFrom(_:)
func (g GCMicroGamepad) SetStateFromMicroGamepad(microGamepad IGCMicroGamepad) {
	objc.Send[objc.ID](g.ID, objc.Sel("setStateFromMicroGamepad:"), microGamepad)
}

// The controller associated with this profile.
//
// See: https://developer.apple.com/documentation/GameController/GCMicroGamepad/controller
func (g GCMicroGamepad) Controller() IGCController {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("controller"))
	return GCControllerFromID(objc.ID(rv))
}

// The block that this profile calls when an element’s value changes.
//
// # Discussion
//
// If multiple elements change values at the same time, the profile calls this
// block once for each element that changed. If the value of a child element
// changes, the profile only calls the block for the containing element.
//
// See: https://developer.apple.com/documentation/GameController/GCMicroGamepad/valueChangedHandler
func (g GCMicroGamepad) ValueChangedHandler() GCMicroGamepadValueChangedHandler {
	rv := objc.Send[GCMicroGamepadValueChangedHandler](g.ID, objc.Sel("valueChangedHandler"))
	return GCMicroGamepadValueChangedHandler(rv)
}
func (g GCMicroGamepad) SetValueChangedHandler(value GCMicroGamepadValueChangedHandler) {
	objc.Send[struct{}](g.ID, objc.Sel("setValueChangedHandler:"), value)
}

// The menu face button that players use to enter the main menu and pause the
// game.
//
// See: https://developer.apple.com/documentation/GameController/GCMicroGamepad/buttonMenu
func (g GCMicroGamepad) ButtonMenu() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("buttonMenu"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The button that the user activates by pressing harder on the touchpad.
//
// # Discussion
//
// This button is digital.
//
// See: https://developer.apple.com/documentation/GameController/GCMicroGamepad/buttonA
func (g GCMicroGamepad) ButtonA() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("buttonA"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The second face button element.
//
// # Discussion
//
// This button is digital.
//
// See: https://developer.apple.com/documentation/GameController/GCMicroGamepad/buttonX
func (g GCMicroGamepad) ButtonX() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("buttonX"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The controller’s directional pad element.
//
// # Discussion
//
// The directional pad in the micro gamepad profile reports analog directional
// information.
//
// See: https://developer.apple.com/documentation/GameController/GCMicroGamepad/dpad
func (g GCMicroGamepad) Dpad() IGCControllerDirectionPad {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("dpad"))
	return GCControllerDirectionPadFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the directional pad reports absolute
// or relative values.
//
// # Discussion
//
// If this property is false, the profile assumes the location where the user
// first touches the pad is the origin value (`0.0,0.0`) for the pad. The
// profile calculates all subsequent values relative to this position until
// the user lifts their finger. The next time the user touches the pad, the
// profile uses that location as the new origin. If this property is true, the
// profile calculates values relative to the physical center of the touchpad.
// The default value for this property is false.
//
// See: https://developer.apple.com/documentation/GameController/GCMicroGamepad/reportsAbsoluteDpadValues
func (g GCMicroGamepad) ReportsAbsoluteDpadValues() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("reportsAbsoluteDpadValues"))
	return rv
}
func (g GCMicroGamepad) SetReportsAbsoluteDpadValues(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setReportsAbsoluteDpadValues:"), value)
}

// A Boolean value that indicates whether the profile reports the directional
// pad values relative to its current orientation.
//
// # Discussion
//
// If this property is false, the profile reports the value of the directional
// pad only in portrait orientation even when the user rotates the controller.
// If this property is true, the profile reports the values using the current
// orientation. The default value for this property is false.
//
// See: https://developer.apple.com/documentation/GameController/GCMicroGamepad/allowsRotation
func (g GCMicroGamepad) AllowsRotation() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("allowsRotation"))
	return rv
}
func (g GCMicroGamepad) SetAllowsRotation(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setAllowsRotation:"), value)
}
