// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [GCExtendedGamepad] class.
var (
	_GCExtendedGamepadClass     GCExtendedGamepadClass
	_GCExtendedGamepadClassOnce sync.Once
)

func getGCExtendedGamepadClass() GCExtendedGamepadClass {
	_GCExtendedGamepadClassOnce.Do(func() {
		_GCExtendedGamepadClass = GCExtendedGamepadClass{class: objc.GetClass("GCExtendedGamepad")}
	})
	return _GCExtendedGamepadClass
}

// GetGCExtendedGamepadClass returns the class object for GCExtendedGamepad.
func GetGCExtendedGamepadClass() GCExtendedGamepadClass {
	return getGCExtendedGamepadClass()
}

type GCExtendedGamepadClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCExtendedGamepadClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCExtendedGamepadClass) Alloc() GCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// A controller profile that supports the extended set of gamepad controls.
//
// # Overview
//
// The extended gamepad controller profile represents a physical or virtual
// controller with the following input elements:
//
// - Two shoulder buttons - Two trigger buttons - Four face buttons in a
// diamond pattern - One directional pad - Two thumbsticks with optional
// thumbstick buttons - Optional Home and Options buttons - A Menu button
//
// [media-3850406]
//
// If a [GCController] object supports this type of profile, get the input
// values of the elements from the controller’s
// [GCController.ExtendedGamepad] property or use the profile’s
// [GCExtendedGamepad.ValueChangedHandler] method to receive a callback when
// the input values change. Alternatively, use the [saveSnapshot()] method to
// capture the input values at a moment in time.
//
// If the controller’s [GCController.ExtendedGamepad] property is `nil`, the
// controller doesn’t support this type of profile. See [GCController] for
// other profiles you can use.
//
// # Getting the controller
//
//   - [GCExtendedGamepad.Controller]: The controller for the profile.
//
// # Getting change information
//
//   - [GCExtendedGamepad.ValueChangedHandler]: The block that the profile calls when an element’s value changes.
//   - [GCExtendedGamepad.SetValueChangedHandler]
//
// # Getting shoulder button inputs
//
//   - [GCExtendedGamepad.LeftShoulder]: The controller’s left shoulder button element.
//   - [GCExtendedGamepad.RightShoulder]: The controller’s right shoulder button element.
//
// # Getting trigger inputs
//
//   - [GCExtendedGamepad.LeftTrigger]: The controller’s left trigger element.
//   - [GCExtendedGamepad.RightTrigger]: The controller’s right trigger element.
//
// # Getting face button inputs
//
//   - [GCExtendedGamepad.ButtonMenu]: The primary menu button element that players use to enter the main menu and pause the game.
//   - [GCExtendedGamepad.ButtonOptions]: The controller’s secondary menu button element.
//   - [GCExtendedGamepad.ButtonHome]: The main menu button element that players use to enter the secondary menu and pause the game.
//   - [GCExtendedGamepad.ButtonA]: The bottom face button that uses A or another indicator as its label.
//   - [GCExtendedGamepad.ButtonB]: The right face button that uses B or another indicator as its label.
//   - [GCExtendedGamepad.ButtonX]: The left face button that uses X or another indicator as its label.
//   - [GCExtendedGamepad.ButtonY]: The top face button that uses Y or another indicator as its label.
//
// # Getting directional pad inputs
//
//   - [GCExtendedGamepad.Dpad]: The controller’s directional pad element.
//
// # Getting thumbstick and thumbstick button inputs
//
//   - [GCExtendedGamepad.LeftThumbstick]: The controller’s left thumbstick element.
//   - [GCExtendedGamepad.RightThumbstick]: The controller’s right thumbstick element.
//   - [GCExtendedGamepad.LeftThumbstickButton]: The button on the left thumbstick of the controller.
//   - [GCExtendedGamepad.RightThumbstickButton]: The button on the right thumbstick of the controller.
//
// # Setting snapshot values
//
//   - [GCExtendedGamepad.SetStateFromExtendedGamepad]: Copies the input values from a specified extended gamepad to a snapshot of an extended gamepad.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad
//
// [saveSnapshot()]: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/saveSnapshot()
type GCExtendedGamepad struct {
	GCPhysicalInputProfile
}

// GCExtendedGamepadFromID constructs a [GCExtendedGamepad] from an objc.ID.
//
// A controller profile that supports the extended set of gamepad controls.
func GCExtendedGamepadFromID(id objc.ID) GCExtendedGamepad {
	return GCExtendedGamepad{GCPhysicalInputProfile: GCPhysicalInputProfileFromID(id)}
}

// NOTE: GCExtendedGamepad adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCExtendedGamepad] class.
//
// # Getting the controller
//
//   - [IGCExtendedGamepad.Controller]: The controller for the profile.
//
// # Getting change information
//
//   - [IGCExtendedGamepad.ValueChangedHandler]: The block that the profile calls when an element’s value changes.
//   - [IGCExtendedGamepad.SetValueChangedHandler]
//
// # Getting shoulder button inputs
//
//   - [IGCExtendedGamepad.LeftShoulder]: The controller’s left shoulder button element.
//   - [IGCExtendedGamepad.RightShoulder]: The controller’s right shoulder button element.
//
// # Getting trigger inputs
//
//   - [IGCExtendedGamepad.LeftTrigger]: The controller’s left trigger element.
//   - [IGCExtendedGamepad.RightTrigger]: The controller’s right trigger element.
//
// # Getting face button inputs
//
//   - [IGCExtendedGamepad.ButtonMenu]: The primary menu button element that players use to enter the main menu and pause the game.
//   - [IGCExtendedGamepad.ButtonOptions]: The controller’s secondary menu button element.
//   - [IGCExtendedGamepad.ButtonHome]: The main menu button element that players use to enter the secondary menu and pause the game.
//   - [IGCExtendedGamepad.ButtonA]: The bottom face button that uses A or another indicator as its label.
//   - [IGCExtendedGamepad.ButtonB]: The right face button that uses B or another indicator as its label.
//   - [IGCExtendedGamepad.ButtonX]: The left face button that uses X or another indicator as its label.
//   - [IGCExtendedGamepad.ButtonY]: The top face button that uses Y or another indicator as its label.
//
// # Getting directional pad inputs
//
//   - [IGCExtendedGamepad.Dpad]: The controller’s directional pad element.
//
// # Getting thumbstick and thumbstick button inputs
//
//   - [IGCExtendedGamepad.LeftThumbstick]: The controller’s left thumbstick element.
//   - [IGCExtendedGamepad.RightThumbstick]: The controller’s right thumbstick element.
//   - [IGCExtendedGamepad.LeftThumbstickButton]: The button on the left thumbstick of the controller.
//   - [IGCExtendedGamepad.RightThumbstickButton]: The button on the right thumbstick of the controller.
//
// # Setting snapshot values
//
//   - [IGCExtendedGamepad.SetStateFromExtendedGamepad]: Copies the input values from a specified extended gamepad to a snapshot of an extended gamepad.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad
type IGCExtendedGamepad interface {
	IGCPhysicalInputProfile

	// Topic: Getting the controller

	// The controller for the profile.
	Controller() IGCController

	// Topic: Getting change information

	// The block that the profile calls when an element’s value changes.
	ValueChangedHandler() GCExtendedGamepadGCControllerElementHandler
	SetValueChangedHandler(value GCExtendedGamepadGCControllerElementHandler)

	// Topic: Getting shoulder button inputs

	// The controller’s left shoulder button element.
	LeftShoulder() IGCControllerButtonInput
	// The controller’s right shoulder button element.
	RightShoulder() IGCControllerButtonInput

	// Topic: Getting trigger inputs

	// The controller’s left trigger element.
	LeftTrigger() IGCControllerButtonInput
	// The controller’s right trigger element.
	RightTrigger() IGCControllerButtonInput

	// Topic: Getting face button inputs

	// The primary menu button element that players use to enter the main menu and pause the game.
	ButtonMenu() IGCControllerButtonInput
	// The controller’s secondary menu button element.
	ButtonOptions() IGCControllerButtonInput
	// The main menu button element that players use to enter the secondary menu and pause the game.
	ButtonHome() IGCControllerButtonInput
	// The bottom face button that uses A or another indicator as its label.
	ButtonA() IGCControllerButtonInput
	// The right face button that uses B or another indicator as its label.
	ButtonB() IGCControllerButtonInput
	// The left face button that uses X or another indicator as its label.
	ButtonX() IGCControllerButtonInput
	// The top face button that uses Y or another indicator as its label.
	ButtonY() IGCControllerButtonInput

	// Topic: Getting directional pad inputs

	// The controller’s directional pad element.
	Dpad() IGCControllerDirectionPad

	// Topic: Getting thumbstick and thumbstick button inputs

	// The controller’s left thumbstick element.
	LeftThumbstick() IGCControllerDirectionPad
	// The controller’s right thumbstick element.
	RightThumbstick() IGCControllerDirectionPad
	// The button on the left thumbstick of the controller.
	LeftThumbstickButton() IGCControllerButtonInput
	// The button on the right thumbstick of the controller.
	RightThumbstickButton() IGCControllerButtonInput

	// Topic: Setting snapshot values

	// Copies the input values from a specified extended gamepad to a snapshot of an extended gamepad.
	SetStateFromExtendedGamepad(extendedGamepad IGCExtendedGamepad)
}

// Init initializes the instance.
func (g GCExtendedGamepad) Init() GCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCExtendedGamepad) Autorelease() GCExtendedGamepad {
	rv := objc.Send[GCExtendedGamepad](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCExtendedGamepad creates a new GCExtendedGamepad instance.
func NewGCExtendedGamepad() GCExtendedGamepad {
	class := getGCExtendedGamepadClass()
	rv := objc.Send[GCExtendedGamepad](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Copies the input values from a specified extended gamepad to a snapshot of
// an extended gamepad.
//
// extendedGamepad: The extended gamepad to copy the input values from.
//
// # Discussion
//
// If this extended gamepad isn’t a snapshot, this method does nothing. A
// snapshot is a copy of a controller at a moment in time that has element
// values you can set.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/setStateFrom(_:)
func (g GCExtendedGamepad) SetStateFromExtendedGamepad(extendedGamepad IGCExtendedGamepad) {
	objc.Send[objc.ID](g.ID, objc.Sel("setStateFromExtendedGamepad:"), extendedGamepad)
}

// The controller for the profile.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/controller
func (g GCExtendedGamepad) Controller() IGCController {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("controller"))
	return GCControllerFromID(objc.ID(rv))
}

// The block that the profile calls when an element’s value changes.
//
// # Discussion
//
// If multiple elements change values at the same time, the profile calls this
// block once for each element that changes. If the value of a subelement
// changes, the profile only calls the block for the containing element.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/valueChangedHandler
func (g GCExtendedGamepad) ValueChangedHandler() GCExtendedGamepadGCControllerElementHandler {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("valueChangedHandler"))
	_ = rv
	return nil
}
func (g GCExtendedGamepad) SetValueChangedHandler(value GCExtendedGamepadGCControllerElementHandler) {
	block, cleanup := NewGCExtendedGamepadGCControllerElementBlock(value)
	defer cleanup()
	objc.Send[struct{}](g.ID, objc.Sel("setValueChangedHandler:"), block)
}

// The controller’s left shoulder button element.
//
// # Discussion
//
// The shoulder buttons in the extended gamepad profile are analog buttons.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/leftShoulder
func (g GCExtendedGamepad) LeftShoulder() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("leftShoulder"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The controller’s right shoulder button element.
//
// # Discussion
//
// The shoulder buttons in the extended gamepad profile are analog buttons.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/rightShoulder
func (g GCExtendedGamepad) RightShoulder() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("rightShoulder"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The controller’s left trigger element.
//
// # Discussion
//
// The triggers in the extended gamepad profile are analog buttons.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/leftTrigger
func (g GCExtendedGamepad) LeftTrigger() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("leftTrigger"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The controller’s right trigger element.
//
// # Discussion
//
// The triggers in the extended gamepad profile are analog buttons.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/rightTrigger
func (g GCExtendedGamepad) RightTrigger() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("rightTrigger"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The primary menu button element that players use to enter the main menu and
// pause the game.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/buttonMenu
func (g GCExtendedGamepad) ButtonMenu() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("buttonMenu"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The controller’s secondary menu button element.
//
// # Discussion
//
// You can use the secondary menu to configure graphics and sound, and to
// pause the game.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/buttonOptions
func (g GCExtendedGamepad) ButtonOptions() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("buttonOptions"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The main menu button element that players use to enter the secondary menu
// and pause the game.
//
// # Discussion
//
// If the system doesn’t process the main menu events, it passes the events
// to your app.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/buttonHome
func (g GCExtendedGamepad) ButtonHome() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("buttonHome"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The bottom face button that uses A or another indicator as its label.
//
// # Discussion
//
// The face buttons in the extended gamepad profile may be either analog or
// digital buttons.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/buttonA
func (g GCExtendedGamepad) ButtonA() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("buttonA"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The right face button that uses B or another indicator as its label.
//
// # Discussion
//
// The face buttons in the extended gamepad profile may be either analog or
// digital buttons.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/buttonB
func (g GCExtendedGamepad) ButtonB() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("buttonB"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The left face button that uses X or another indicator as its label.
//
// # Discussion
//
// The face buttons in the extended gamepad profile may be either analog or
// digital buttons.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/buttonX
func (g GCExtendedGamepad) ButtonX() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("buttonX"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The top face button that uses Y or another indicator as its label.
//
// # Discussion
//
// The face buttons in the extended gamepad profile may be either analog or
// digital buttons.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/buttonY
func (g GCExtendedGamepad) ButtonY() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("buttonY"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The controller’s directional pad element.
//
// # Discussion
//
// The directional pad in the extended gamepad profile reports analog
// directional information.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/dpad
func (g GCExtendedGamepad) Dpad() IGCControllerDirectionPad {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("dpad"))
	return GCControllerDirectionPadFromID(objc.ID(rv))
}

// The controller’s left thumbstick element.
//
// # Discussion
//
// The thumbsticks in the extended gamepad profile are analog buttons.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/leftThumbstick
func (g GCExtendedGamepad) LeftThumbstick() IGCControllerDirectionPad {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("leftThumbstick"))
	return GCControllerDirectionPadFromID(objc.ID(rv))
}

// The controller’s right thumbstick element.
//
// # Discussion
//
// The thumbsticks in the extended gamepad profile are analog buttons.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/rightThumbstick
func (g GCExtendedGamepad) RightThumbstick() IGCControllerDirectionPad {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("rightThumbstick"))
	return GCControllerDirectionPadFromID(objc.ID(rv))
}

// The button on the left thumbstick of the controller.
//
// # Discussion
//
// If the thumbstick has a clickable component, this is a digital button.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/leftThumbstickButton
func (g GCExtendedGamepad) LeftThumbstickButton() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("leftThumbstickButton"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The button on the right thumbstick of the controller.
//
// # Discussion
//
// If the thumbstick has a clickable component, this is a digital button.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepad/rightThumbstickButton
func (g GCExtendedGamepad) RightThumbstickButton() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("rightThumbstickButton"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}
