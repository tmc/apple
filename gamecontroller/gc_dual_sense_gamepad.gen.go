// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [GCDualSenseGamepad] class.
var (
	_GCDualSenseGamepadClass     GCDualSenseGamepadClass
	_GCDualSenseGamepadClassOnce sync.Once
)

func getGCDualSenseGamepadClass() GCDualSenseGamepadClass {
	_GCDualSenseGamepadClassOnce.Do(func() {
		_GCDualSenseGamepadClass = GCDualSenseGamepadClass{class: objc.GetClass("GCDualSenseGamepad")}
	})
	return _GCDualSenseGamepadClass
}

// GetGCDualSenseGamepadClass returns the class object for GCDualSenseGamepad.
func GetGCDualSenseGamepadClass() GCDualSenseGamepadClass {
	return getGCDualSenseGamepadClass()
}

type GCDualSenseGamepadClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCDualSenseGamepadClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCDualSenseGamepadClass) Alloc() GCDualSenseGamepad {
	rv := objc.Send[GCDualSenseGamepad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// A controller profile that supported the DualSense controller.
//
// # Overview
//
// The DualSense controller profile is similar to a DualShock profile
// ([GCDualShockGamepad]), but has adaptive triggers that allow you to specify
// a dynamic resistance force when the user pulls the trigger. For example,
// you can emulate the feeling of pulling back a bow string, firing a weapon,
// or pulling a lever.
//
// This profile also supports motion — that is, the controller’s [GCDualSenseGamepad.Motion]
// property is non-nil. If you hold the controller in front of you, the
// direction of the axes are:
//
// - The positive x-axis points to your right. - The positive y-axis points up
// out of the USB-C port. - The positive z-axis starts at the touchpad and
// points to you.
//
// # Getting button input
//
//   - [GCDualSenseGamepad.TouchpadButton]: The button element on the touchpad of the controller.
//
// # Tracking finger locations
//
//   - [GCDualSenseGamepad.TouchpadPrimary]: The location of the player’s primary finger on the touchpad.
//   - [GCDualSenseGamepad.TouchpadSecondary]: The location of the player’s secondary finger on the touchpad.
//
// See: https://developer.apple.com/documentation/GameController/GCDualSenseGamepad
type GCDualSenseGamepad struct {
	GCExtendedGamepad
}

// GCDualSenseGamepadFromID constructs a [GCDualSenseGamepad] from an objc.ID.
//
// A controller profile that supported the DualSense controller.
func GCDualSenseGamepadFromID(id objc.ID) GCDualSenseGamepad {
	return GCDualSenseGamepad{GCExtendedGamepad: GCExtendedGamepadFromID(id)}
}

// NOTE: GCDualSenseGamepad adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCDualSenseGamepad] class.
//
// # Getting button input
//
//   - [IGCDualSenseGamepad.TouchpadButton]: The button element on the touchpad of the controller.
//
// # Tracking finger locations
//
//   - [IGCDualSenseGamepad.TouchpadPrimary]: The location of the player’s primary finger on the touchpad.
//   - [IGCDualSenseGamepad.TouchpadSecondary]: The location of the player’s secondary finger on the touchpad.
//
// See: https://developer.apple.com/documentation/GameController/GCDualSenseGamepad
type IGCDualSenseGamepad interface {
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
func (g GCDualSenseGamepad) Init() GCDualSenseGamepad {
	rv := objc.Send[GCDualSenseGamepad](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCDualSenseGamepad) Autorelease() GCDualSenseGamepad {
	rv := objc.Send[GCDualSenseGamepad](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCDualSenseGamepad creates a new GCDualSenseGamepad instance.
func NewGCDualSenseGamepad() GCDualSenseGamepad {
	class := getGCDualSenseGamepadClass()
	rv := objc.Send[GCDualSenseGamepad](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The button element on the touchpad of the controller.
//
// See: https://developer.apple.com/documentation/GameController/GCDualSenseGamepad/touchpadButton
func (g GCDualSenseGamepad) TouchpadButton() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("touchpadButton"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The location of the player’s primary finger on the touchpad.
//
// See: https://developer.apple.com/documentation/GameController/GCDualSenseGamepad/touchpadPrimary
func (g GCDualSenseGamepad) TouchpadPrimary() IGCControllerDirectionPad {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("touchpadPrimary"))
	return GCControllerDirectionPadFromID(objc.ID(rv))
}

// The location of the player’s secondary finger on the touchpad.
//
// See: https://developer.apple.com/documentation/GameController/GCDualSenseGamepad/touchpadSecondary
func (g GCDualSenseGamepad) TouchpadSecondary() IGCControllerDirectionPad {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("touchpadSecondary"))
	return GCControllerDirectionPadFromID(objc.ID(rv))
}
