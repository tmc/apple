// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [GCDualShockGamepad] class.
var (
	_GCDualShockGamepadClass     GCDualShockGamepadClass
	_GCDualShockGamepadClassOnce sync.Once
)

func getGCDualShockGamepadClass() GCDualShockGamepadClass {
	_GCDualShockGamepadClassOnce.Do(func() {
		_GCDualShockGamepadClass = GCDualShockGamepadClass{class: objc.GetClass("GCDualShockGamepad")}
	})
	return _GCDualShockGamepadClass
}

// GetGCDualShockGamepadClass returns the class object for GCDualShockGamepad.
func GetGCDualShockGamepadClass() GCDualShockGamepadClass {
	return getGCDualShockGamepadClass()
}

type GCDualShockGamepadClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCDualShockGamepadClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCDualShockGamepadClass) Alloc() GCDualShockGamepad {
	rv := objc.Send[GCDualShockGamepad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// A controller profile that supports the DualShock 4 controller.
//
// # Overview
//
// The DualShock 4 controller profile is similar to an extended gamepad
// ([GCExtendedGamepad]), but has a touchpad with a button and two-finger
// tracking.
//
// [media-3830806]
//
// This profile also supports motion — that is, the controller’s
// [GCController.Motion] property is non-nil. If you hold the controller in
// front of you, the direction of the axes are:
//
// - The positive x-axis points to your right. - The positive y-axis points
// up. - The positive z-axis starts at the touchpad and points to you.
//
// [media-3856422]
//
// # Getting button input
//
//   - [GCDualShockGamepad.TouchpadButton]: The button element on the touchpad of the controller.
//
// # Tracking finger locations
//
//   - [GCDualShockGamepad.TouchpadPrimary]: The location of the player’s primary finger on the touchpad.
//   - [GCDualShockGamepad.TouchpadSecondary]: The location of the player’s secondary finger on the touchpad.
//
// See: https://developer.apple.com/documentation/GameController/GCDualShockGamepad
type GCDualShockGamepad struct {
	GCExtendedGamepad
}

// GCDualShockGamepadFromID constructs a [GCDualShockGamepad] from an objc.ID.
//
// A controller profile that supports the DualShock 4 controller.
func GCDualShockGamepadFromID(id objc.ID) GCDualShockGamepad {
	return GCDualShockGamepad{GCExtendedGamepad: GCExtendedGamepadFromID(id)}
}

// NOTE: GCDualShockGamepad adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCDualShockGamepad] class.
//
// # Getting button input
//
//   - [IGCDualShockGamepad.TouchpadButton]: The button element on the touchpad of the controller.
//
// # Tracking finger locations
//
//   - [IGCDualShockGamepad.TouchpadPrimary]: The location of the player’s primary finger on the touchpad.
//   - [IGCDualShockGamepad.TouchpadSecondary]: The location of the player’s secondary finger on the touchpad.
//
// See: https://developer.apple.com/documentation/GameController/GCDualShockGamepad
type IGCDualShockGamepad interface {
	IGCExtendedGamepad

	// Topic: Getting button input

	// The button element on the touchpad of the controller.
	TouchpadButton() IGCControllerButtonInput

	// Topic: Tracking finger locations

	// The location of the player’s primary finger on the touchpad.
	TouchpadPrimary() IGCControllerDirectionPad
	// The location of the player’s secondary finger on the touchpad.
	TouchpadSecondary() IGCControllerDirectionPad
}

// Init initializes the instance.
func (g GCDualShockGamepad) Init() GCDualShockGamepad {
	rv := objc.Send[GCDualShockGamepad](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCDualShockGamepad) Autorelease() GCDualShockGamepad {
	rv := objc.Send[GCDualShockGamepad](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCDualShockGamepad creates a new GCDualShockGamepad instance.
func NewGCDualShockGamepad() GCDualShockGamepad {
	class := getGCDualShockGamepadClass()
	rv := objc.Send[GCDualShockGamepad](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The button element on the touchpad of the controller.
//
// See: https://developer.apple.com/documentation/GameController/GCDualShockGamepad/touchpadButton
func (g GCDualShockGamepad) TouchpadButton() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("touchpadButton"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The location of the player’s primary finger on the touchpad.
//
// See: https://developer.apple.com/documentation/GameController/GCDualShockGamepad/touchpadPrimary
func (g GCDualShockGamepad) TouchpadPrimary() IGCControllerDirectionPad {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("touchpadPrimary"))
	return GCControllerDirectionPadFromID(objc.ID(rv))
}

// The location of the player’s secondary finger on the touchpad.
//
// See: https://developer.apple.com/documentation/GameController/GCDualShockGamepad/touchpadSecondary
func (g GCDualShockGamepad) TouchpadSecondary() IGCControllerDirectionPad {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("touchpadSecondary"))
	return GCControllerDirectionPadFromID(objc.ID(rv))
}
