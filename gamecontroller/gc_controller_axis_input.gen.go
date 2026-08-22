// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [GCControllerAxisInput] class.
var (
	_GCControllerAxisInputClass     GCControllerAxisInputClass
	_GCControllerAxisInputClassOnce sync.Once
)

func getGCControllerAxisInputClass() GCControllerAxisInputClass {
	_GCControllerAxisInputClassOnce.Do(func() {
		_GCControllerAxisInputClass = GCControllerAxisInputClass{class: objc.GetClass("GCControllerAxisInput")}
	})
	return _GCControllerAxisInputClass
}

// GetGCControllerAxisInputClass returns the class object for GCControllerAxisInput.
func GetGCControllerAxisInputClass() GCControllerAxisInputClass {
	return getGCControllerAxisInputClass()
}

type GCControllerAxisInputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCControllerAxisInputClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCControllerAxisInputClass) Alloc() GCControllerAxisInput {
	rv := objc.Send[GCControllerAxisInput](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// A control element that tracks movement along an axis.
//
// # Overview
//
// A [GCControllerAxisInput] object represents the value of a physical
// controller’s axis. For example, a [GCControllerDirectionPad] has x-axis
// and y-axis subelements.
//
// # Accessing the input values
//
//   - [GCControllerAxisInput.Value]: The current value of the axis.
//
// # Getting change information
//
//   - [GCControllerAxisInput.ValueChangedHandler]: The block that the element calls when the user changes the axis value.
//   - [GCControllerAxisInput.SetValueChangedHandler]
//
// See: https://developer.apple.com/documentation/GameController/GCControllerAxisInput
type GCControllerAxisInput struct {
	GCControllerElement
}

// GCControllerAxisInputFromID constructs a [GCControllerAxisInput] from an objc.ID.
//
// A control element that tracks movement along an axis.
func GCControllerAxisInputFromID(id objc.ID) GCControllerAxisInput {
	return GCControllerAxisInput{GCControllerElement: GCControllerElementFromID(id)}
}

// NOTE: GCControllerAxisInput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCControllerAxisInput] class.
//
// # Accessing the input values
//
//   - [IGCControllerAxisInput.Value]: The current value of the axis.
//
// # Getting change information
//
//   - [IGCControllerAxisInput.ValueChangedHandler]: The block that the element calls when the user changes the axis value.
//   - [IGCControllerAxisInput.SetValueChangedHandler]
//
// See: https://developer.apple.com/documentation/GameController/GCControllerAxisInput
type IGCControllerAxisInput interface {
	IGCControllerElement

	// Topic: Accessing the input values

	// The current value of the axis.
	Value() float32

	// Topic: Getting change information

	// The block that the element calls when the user changes the axis value.
	ValueChangedHandler() GCControllerAxisInputFloat32Handler
	SetValueChangedHandler(value GCControllerAxisInputFloat32Handler)
}

// Init initializes the instance.
func (g GCControllerAxisInput) Init() GCControllerAxisInput {
	rv := objc.Send[GCControllerAxisInput](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCControllerAxisInput) Autorelease() GCControllerAxisInput {
	rv := objc.Send[GCControllerAxisInput](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCControllerAxisInput creates a new GCControllerAxisInput instance.
func NewGCControllerAxisInput() GCControllerAxisInput {
	class := getGCControllerAxisInputClass()
	rv := objc.Send[GCControllerAxisInput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The current value of the axis.
//
// # Discussion
//
// Often a physical controller ignores values near the neutral position called
// the dead zone. The [GCControllerAxisInput] element handles this dead zone,
// and other physical constraints of a hardware control, by computing a
// normalized value.
//
// The normalized value ranges from `-1` to `1`. If the value is `0`, the
// movement is in the dead zone. A nonzero value indicates the moment is
// outside of the dead zone.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerAxisInput/value
func (g GCControllerAxisInput) Value() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("value"))
	return rv
}

// The block that the element calls when the user changes the axis value.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerAxisInput/valueChangedHandler
func (g GCControllerAxisInput) ValueChangedHandler() GCControllerAxisInputFloat32Handler {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("valueChangedHandler"))
	_ = rv
	return nil
}
func (g GCControllerAxisInput) SetValueChangedHandler(value GCControllerAxisInputFloat32Handler) {
	block, cleanup := NewGCControllerAxisInputFloat32Block(value)
	defer cleanup()
	objc.Send[struct{}](g.ID, objc.Sel("setValueChangedHandler:"), block)
}
