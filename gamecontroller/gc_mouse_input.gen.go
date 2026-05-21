// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [GCMouseInput] class.
var (
	_GCMouseInputClass     GCMouseInputClass
	_GCMouseInputClassOnce sync.Once
)

func getGCMouseInputClass() GCMouseInputClass {
	_GCMouseInputClassOnce.Do(func() {
		_GCMouseInputClass = GCMouseInputClass{class: objc.GetClass("GCMouseInput")}
	})
	return _GCMouseInputClass
}

// GetGCMouseInputClass returns the class object for GCMouseInput.
func GetGCMouseInputClass() GCMouseInputClass {
	return getGCMouseInputClass()
}

type GCMouseInputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCMouseInputClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCMouseInputClass) Alloc() GCMouseInput {
	rv := objc.Send[GCMouseInput](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// A controller profile that tracks input from a mouse.
//
// # Overview
//
// This profile supports a mouse with the following features:
//
// - A two-axis cursor and scroll - A left button - An optional right button -
// An optional middle button - An optional set of auxiliary buttons
//
// This profile provides only raw mouse movement delta values. For the cursor
// position at a specific time, use the [UIHoverGestureRecognizer] class and
// the [NSEvent] [mouseLocation] method.
//
// # Getting Change Information
//
//   - [GCMouseInput.MouseMovedHandler]: The block that the profile calls when the mouse moves.
//   - [GCMouseInput.SetMouseMovedHandler]
//
// # Accessing Buttons
//
//   - [GCMouseInput.LeftButton]: The left button on the mouse.
//   - [GCMouseInput.RightButton]: The optional right button on the mouse.
//   - [GCMouseInput.MiddleButton]: The optional middle button on the mouse.
//   - [GCMouseInput.AuxiliaryButtons]: The optional additional buttons on the mouse.
//
// # Scrolling
//
//   - [GCMouseInput.Scroll]: The location of the directional pad cursor with an undefined range.
//
// See: https://developer.apple.com/documentation/GameController/GCMouseInput
//
// [UIHoverGestureRecognizer]: https://developer.apple.com/documentation/UIKit/UIHoverGestureRecognizer
// [mouseLocation]: https://developer.apple.com/documentation/AppKit/NSEvent/mouseLocation
type GCMouseInput struct {
	GCPhysicalInputProfile
}

// GCMouseInputFromID constructs a [GCMouseInput] from an objc.ID.
//
// A controller profile that tracks input from a mouse.
func GCMouseInputFromID(id objc.ID) GCMouseInput {
	return GCMouseInput{GCPhysicalInputProfile: GCPhysicalInputProfileFromID(id)}
}

// NOTE: GCMouseInput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCMouseInput] class.
//
// # Getting Change Information
//
//   - [IGCMouseInput.MouseMovedHandler]: The block that the profile calls when the mouse moves.
//   - [IGCMouseInput.SetMouseMovedHandler]
//
// # Accessing Buttons
//
//   - [IGCMouseInput.LeftButton]: The left button on the mouse.
//   - [IGCMouseInput.RightButton]: The optional right button on the mouse.
//   - [IGCMouseInput.MiddleButton]: The optional middle button on the mouse.
//   - [IGCMouseInput.AuxiliaryButtons]: The optional additional buttons on the mouse.
//
// # Scrolling
//
//   - [IGCMouseInput.Scroll]: The location of the directional pad cursor with an undefined range.
//
// See: https://developer.apple.com/documentation/GameController/GCMouseInput
type IGCMouseInput interface {
	IGCPhysicalInputProfile

	// Topic: Getting Change Information

	// The block that the profile calls when the mouse moves.
	MouseMovedHandler() GCMouseMoved
	SetMouseMovedHandler(value GCMouseMoved)

	// Topic: Accessing Buttons

	// The left button on the mouse.
	LeftButton() IGCControllerButtonInput
	// The optional right button on the mouse.
	RightButton() IGCControllerButtonInput
	// The optional middle button on the mouse.
	MiddleButton() IGCControllerButtonInput
	// The optional additional buttons on the mouse.
	AuxiliaryButtons() []GCControllerButtonInput

	// Topic: Scrolling

	// The location of the directional pad cursor with an undefined range.
	Scroll() IGCDeviceCursor
}

// Init initializes the instance.
func (g GCMouseInput) Init() GCMouseInput {
	rv := objc.Send[GCMouseInput](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCMouseInput) Autorelease() GCMouseInput {
	rv := objc.Send[GCMouseInput](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCMouseInput creates a new GCMouseInput instance.
func NewGCMouseInput() GCMouseInput {
	class := getGCMouseInputClass()
	rv := objc.Send[GCMouseInput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The block that the profile calls when the mouse moves.
//
// See: https://developer.apple.com/documentation/GameController/GCMouseInput/mouseMovedHandler
func (g GCMouseInput) MouseMovedHandler() GCMouseMoved {
	rv := objc.Send[GCMouseMoved](g.ID, objc.Sel("mouseMovedHandler"))
	return GCMouseMoved(rv)
}
func (g GCMouseInput) SetMouseMovedHandler(value GCMouseMoved) {
	objc.Send[struct{}](g.ID, objc.Sel("setMouseMovedHandler:"), value)
}

// The left button on the mouse.
//
// See: https://developer.apple.com/documentation/GameController/GCMouseInput/leftButton
func (g GCMouseInput) LeftButton() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("leftButton"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The optional right button on the mouse.
//
// # Discussion
//
// If the mouse doesn’t have a right button, this property is `nil`.
//
// See: https://developer.apple.com/documentation/GameController/GCMouseInput/rightButton
func (g GCMouseInput) RightButton() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("rightButton"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The optional middle button on the mouse.
//
// # Discussion
//
// If the mouse doesn’t have a middle button, this property is `nil`.
//
// See: https://developer.apple.com/documentation/GameController/GCMouseInput/middleButton
func (g GCMouseInput) MiddleButton() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("middleButton"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The optional additional buttons on the mouse.
//
// # Discussion
//
// If the mouse doesn’t have additional buttons, this property is `nil`.
//
// See: https://developer.apple.com/documentation/GameController/GCMouseInput/auxiliaryButtons
func (g GCMouseInput) AuxiliaryButtons() []GCControllerButtonInput {
	rv := objc.Send[[]objc.ID](g.ID, objc.Sel("auxiliaryButtons"))
	return objc.ConvertSlice(rv, func(id objc.ID) GCControllerButtonInput {
		return GCControllerButtonInputFromID(id)
	})
}

// The location of the directional pad cursor with an undefined range.
//
// See: https://developer.apple.com/documentation/GameController/GCMouseInput/scroll
func (g GCMouseInput) Scroll() IGCDeviceCursor {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("scroll"))
	return GCDeviceCursorFromID(objc.ID(rv))
}
