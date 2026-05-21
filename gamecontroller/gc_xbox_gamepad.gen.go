// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [GCXboxGamepad] class.
var (
	_GCXboxGamepadClass     GCXboxGamepadClass
	_GCXboxGamepadClassOnce sync.Once
)

func getGCXboxGamepadClass() GCXboxGamepadClass {
	_GCXboxGamepadClassOnce.Do(func() {
		_GCXboxGamepadClass = GCXboxGamepadClass{class: objc.GetClass("GCXboxGamepad")}
	})
	return _GCXboxGamepadClass
}

// GetGCXboxGamepadClass returns the class object for GCXboxGamepad.
func GetGCXboxGamepadClass() GCXboxGamepadClass {
	return getGCXboxGamepadClass()
}

type GCXboxGamepadClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCXboxGamepadClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCXboxGamepadClass) Alloc() GCXboxGamepad {
	rv := objc.Send[GCXboxGamepad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// A controller profile that supports the Xbox controller.
//
// # Overview
//
// The Xbox controller profile is similar to an extended game pad
// ([GCExtendedGamepad]), but has four paddle button elements.
//
// [media-3830808]
//
// # Getting button inputs
//
//   - [GCXboxGamepad.PaddleButton1]: The controller’s paddle 1 button element, which has a P1 label on the back of the controller.
//   - [GCXboxGamepad.PaddleButton2]: The paddle 2 button element, which has a P2 label on the back of the controller.
//   - [GCXboxGamepad.PaddleButton3]: The paddle 3 button element, which has a P3 label on the back of the controller.
//   - [GCXboxGamepad.PaddleButton4]: The paddle 4 button element, which has a P4 label on the back of the controller.
//   - [GCXboxGamepad.ButtonShare]: The share button on an Xbox Series X|S controller or later.
//
// See: https://developer.apple.com/documentation/GameController/GCXboxGamepad
type GCXboxGamepad struct {
	GCExtendedGamepad
}

// GCXboxGamepadFromID constructs a [GCXboxGamepad] from an objc.ID.
//
// A controller profile that supports the Xbox controller.
func GCXboxGamepadFromID(id objc.ID) GCXboxGamepad {
	return GCXboxGamepad{GCExtendedGamepad: GCExtendedGamepadFromID(id)}
}

// NOTE: GCXboxGamepad adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCXboxGamepad] class.
//
// # Getting button inputs
//
//   - [IGCXboxGamepad.PaddleButton1]: The controller’s paddle 1 button element, which has a P1 label on the back of the controller.
//   - [IGCXboxGamepad.PaddleButton2]: The paddle 2 button element, which has a P2 label on the back of the controller.
//   - [IGCXboxGamepad.PaddleButton3]: The paddle 3 button element, which has a P3 label on the back of the controller.
//   - [IGCXboxGamepad.PaddleButton4]: The paddle 4 button element, which has a P4 label on the back of the controller.
//   - [IGCXboxGamepad.ButtonShare]: The share button on an Xbox Series X|S controller or later.
//
// See: https://developer.apple.com/documentation/GameController/GCXboxGamepad
type IGCXboxGamepad interface {
	IGCExtendedGamepad

	// Topic: Getting button inputs

	// The controller’s paddle 1 button element, which has a P1 label on the back of the controller.
	PaddleButton1() IGCControllerButtonInput
	// The paddle 2 button element, which has a P2 label on the back of the controller.
	PaddleButton2() IGCControllerButtonInput
	// The paddle 3 button element, which has a P3 label on the back of the controller.
	PaddleButton3() IGCControllerButtonInput
	// The paddle 4 button element, which has a P4 label on the back of the controller.
	PaddleButton4() IGCControllerButtonInput
	// The share button on an Xbox Series X|S controller or later.
	ButtonShare() IGCControllerButtonInput
}

// Init initializes the instance.
func (g GCXboxGamepad) Init() GCXboxGamepad {
	rv := objc.Send[GCXboxGamepad](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCXboxGamepad) Autorelease() GCXboxGamepad {
	rv := objc.Send[GCXboxGamepad](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCXboxGamepad creates a new GCXboxGamepad instance.
func NewGCXboxGamepad() GCXboxGamepad {
	class := getGCXboxGamepadClass()
	rv := objc.Send[GCXboxGamepad](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The controller’s paddle 1 button element, which has a P1 label on the
// back of the controller.
//
// See: https://developer.apple.com/documentation/GameController/GCXboxGamepad/paddleButton1
func (g GCXboxGamepad) PaddleButton1() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("paddleButton1"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The paddle 2 button element, which has a P2 label on the back of the
// controller.
//
// See: https://developer.apple.com/documentation/GameController/GCXboxGamepad/paddleButton2
func (g GCXboxGamepad) PaddleButton2() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("paddleButton2"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The paddle 3 button element, which has a P3 label on the back of the
// controller.
//
// See: https://developer.apple.com/documentation/GameController/GCXboxGamepad/paddleButton3
func (g GCXboxGamepad) PaddleButton3() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("paddleButton3"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The paddle 4 button element, which has a P4 label on the back of the
// controller.
//
// See: https://developer.apple.com/documentation/GameController/GCXboxGamepad/paddleButton4
func (g GCXboxGamepad) PaddleButton4() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("paddleButton4"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The share button on an Xbox Series X|S controller or later.
//
// # Discussion
//
// The system reserves the Share button for screenshot and video recording
// gestures. If you want to disable these gestures in your app, set the
// button’s [GCControllerElement.PreferredSystemGestureState] to
// [GCSystemGestureStateDisabled].
//
// See: https://developer.apple.com/documentation/GameController/GCXboxGamepad/buttonShare
func (g GCXboxGamepad) ButtonShare() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("buttonShare"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}
