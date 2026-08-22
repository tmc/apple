// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [GCKeyboardInput] class.
var (
	_GCKeyboardInputClass     GCKeyboardInputClass
	_GCKeyboardInputClassOnce sync.Once
)

func getGCKeyboardInputClass() GCKeyboardInputClass {
	_GCKeyboardInputClassOnce.Do(func() {
		_GCKeyboardInputClass = GCKeyboardInputClass{class: objc.GetClass("GCKeyboardInput")}
	})
	return _GCKeyboardInputClass
}

// GetGCKeyboardInputClass returns the class object for GCKeyboardInput.
func GetGCKeyboardInputClass() GCKeyboardInputClass {
	return getGCKeyboardInputClass()
}

type GCKeyboardInputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCKeyboardInputClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCKeyboardInputClass) Alloc() GCKeyboardInput {
	rv := objc.Send[GCKeyboardInput](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// A controller profile that uses the keyboard as the input device.
//
// # Overview
//
// Use this profile to get the state of the keyboard buttons that the
// [GCKeyCode] structure defines.
//
// # Getting Change Information
//
//   - [GCKeyboardInput.KeyChangedHandler]: The block that the profile calls when the user presses a key.
//   - [GCKeyboardInput.SetKeyChangedHandler]
//
// # Accessing Buttons
//
//   - [GCKeyboardInput.IsAnyKeyPressed]: A Boolean value that indicates whether the user is pressing any of the keys.
//   - [GCKeyboardInput.ButtonForKeyCode]: Returns the button element for the specified key code.
//
// See: https://developer.apple.com/documentation/GameController/GCKeyboardInput
type GCKeyboardInput struct {
	GCPhysicalInputProfile
}

// GCKeyboardInputFromID constructs a [GCKeyboardInput] from an objc.ID.
//
// A controller profile that uses the keyboard as the input device.
func GCKeyboardInputFromID(id objc.ID) GCKeyboardInput {
	return GCKeyboardInput{GCPhysicalInputProfile: GCPhysicalInputProfileFromID(id)}
}

// NOTE: GCKeyboardInput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCKeyboardInput] class.
//
// # Getting Change Information
//
//   - [IGCKeyboardInput.KeyChangedHandler]: The block that the profile calls when the user presses a key.
//   - [IGCKeyboardInput.SetKeyChangedHandler]
//
// # Accessing Buttons
//
//   - [IGCKeyboardInput.IsAnyKeyPressed]: A Boolean value that indicates whether the user is pressing any of the keys.
//   - [IGCKeyboardInput.ButtonForKeyCode]: Returns the button element for the specified key code.
//
// See: https://developer.apple.com/documentation/GameController/GCKeyboardInput
type IGCKeyboardInput interface {
	IGCPhysicalInputProfile

	// Topic: Getting Change Information

	// The block that the profile calls when the user presses a key.
	KeyChangedHandler() GCKeyboardInputGCControllerButtonInputInt64BoolHandler
	SetKeyChangedHandler(value GCKeyboardInputGCControllerButtonInputInt64BoolHandler)

	// Topic: Accessing Buttons

	// A Boolean value that indicates whether the user is pressing any of the keys.
	IsAnyKeyPressed() bool
	// Returns the button element for the specified key code.
	ButtonForKeyCode(code GCKeyCode) IGCControllerButtonInput
}

// Init initializes the instance.
func (g GCKeyboardInput) Init() GCKeyboardInput {
	rv := objc.Send[GCKeyboardInput](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCKeyboardInput) Autorelease() GCKeyboardInput {
	rv := objc.Send[GCKeyboardInput](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCKeyboardInput creates a new GCKeyboardInput instance.
func NewGCKeyboardInput() GCKeyboardInput {
	class := getGCKeyboardInputClass()
	rv := objc.Send[GCKeyboardInput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the button element for the specified key code.
//
// code: The code for the keyboard button element.
//
// # Return Value
//
// The keyboard button element that this profile defines for the specified key
// code.
//
// # Discussion
//
// Alternatively, you can get a button element for a key using the
// [GCPhysicalInputProfile.ObjectForKeyedSubscript] notation that you inherit
// from [GCPhysicalInputProfile], as in `keyboard[GCKeyUpArrow]`.
//
// See: https://developer.apple.com/documentation/GameController/GCKeyboardInput/button(forKeyCode:)
func (g GCKeyboardInput) ButtonForKeyCode(code GCKeyCode) IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("buttonForKeyCode:"), code)
	return GCControllerButtonInputFromID(rv)
}

// The block that the profile calls when the user presses a key.
//
// # Discussion
//
// If multiple keys change values at the same time, the profile calls this
// block once for each key that changes.
//
// See: https://developer.apple.com/documentation/GameController/GCKeyboardInput/keyChangedHandler
func (g GCKeyboardInput) KeyChangedHandler() GCKeyboardInputGCControllerButtonInputInt64BoolHandler {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("keyChangedHandler"))
	_ = rv
	return nil
}
func (g GCKeyboardInput) SetKeyChangedHandler(value GCKeyboardInputGCControllerButtonInputInt64BoolHandler) {
	block, cleanup := NewGCKeyboardInputGCControllerButtonInputInt64BoolBlock(value)
	defer cleanup()
	objc.Send[struct{}](g.ID, objc.Sel("setKeyChangedHandler:"), block)
}

// A Boolean value that indicates whether the user is pressing any of the
// keys.
//
// # Discussion
//
// If true, the user is pressing a key; otherwise, the user isn’t. You can
// use this property to check whether the user presses any key before getting
// the state of specific keys.
//
// See: https://developer.apple.com/documentation/GameController/GCKeyboardInput/isAnyKeyPressed
func (g GCKeyboardInput) IsAnyKeyPressed() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isAnyKeyPressed"))
	return rv
}
