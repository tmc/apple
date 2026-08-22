// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [GCControllerDirectionPad] class.
var (
	_GCControllerDirectionPadClass     GCControllerDirectionPadClass
	_GCControllerDirectionPadClassOnce sync.Once
)

func getGCControllerDirectionPadClass() GCControllerDirectionPadClass {
	_GCControllerDirectionPadClassOnce.Do(func() {
		_GCControllerDirectionPadClass = GCControllerDirectionPadClass{class: objc.GetClass("GCControllerDirectionPad")}
	})
	return _GCControllerDirectionPadClass
}

// GetGCControllerDirectionPadClass returns the class object for GCControllerDirectionPad.
func GetGCControllerDirectionPadClass() GCControllerDirectionPadClass {
	return getGCControllerDirectionPadClass()
}

type GCControllerDirectionPadClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCControllerDirectionPadClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCControllerDirectionPadClass) Alloc() GCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// A control element associated with a directional pad or a thumbstick.
//
// # Overview
//
// You get the input values for this element from its subelements. You can use
// either the [GCControllerDirectionPad.XAxis] and
// [GCControllerDirectionPad.YAxis] properties to get coordinates, or the
// [GCControllerDirectionPad.Up], [GCControllerDirectionPad.Down],
// [GCControllerDirectionPad.Left], and [GCControllerDirectionPad.Right]
// buttons that simulate directional pad buttons.
//
// # Accessing values using the axes
//
//   - [GCControllerDirectionPad.XAxis]: The x-axis element of the directional pad.
//   - [GCControllerDirectionPad.YAxis]: The y-axis element of the directional pad.
//
// # Accessing values using directional buttons
//
//   - [GCControllerDirectionPad.Right]: The button element that changes the positive x-axis.
//   - [GCControllerDirectionPad.Left]: The button element that changes the negative x-axis.
//   - [GCControllerDirectionPad.Up]: The button element that changes the positive y-axis.
//   - [GCControllerDirectionPad.Down]: The button element used for the negative y-axis direction.
//
// # Getting change information
//
//   - [GCControllerDirectionPad.ValueChangedHandler]: The block that the directional pad calls when the user changes its values.
//   - [GCControllerDirectionPad.SetValueChangedHandler]
//
// # Setting snapshot values
//
//   - [GCControllerDirectionPad.SetValueForXAxisYAxis]: Sets the input values of a snapshot of a directional pad.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad
type GCControllerDirectionPad struct {
	GCControllerElement
}

// GCControllerDirectionPadFromID constructs a [GCControllerDirectionPad] from an objc.ID.
//
// A control element associated with a directional pad or a thumbstick.
func GCControllerDirectionPadFromID(id objc.ID) GCControllerDirectionPad {
	return GCControllerDirectionPad{GCControllerElement: GCControllerElementFromID(id)}
}

// NOTE: GCControllerDirectionPad adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCControllerDirectionPad] class.
//
// # Accessing values using the axes
//
//   - [IGCControllerDirectionPad.XAxis]: The x-axis element of the directional pad.
//   - [IGCControllerDirectionPad.YAxis]: The y-axis element of the directional pad.
//
// # Accessing values using directional buttons
//
//   - [IGCControllerDirectionPad.Right]: The button element that changes the positive x-axis.
//   - [IGCControllerDirectionPad.Left]: The button element that changes the negative x-axis.
//   - [IGCControllerDirectionPad.Up]: The button element that changes the positive y-axis.
//   - [IGCControllerDirectionPad.Down]: The button element used for the negative y-axis direction.
//
// # Getting change information
//
//   - [IGCControllerDirectionPad.ValueChangedHandler]: The block that the directional pad calls when the user changes its values.
//   - [IGCControllerDirectionPad.SetValueChangedHandler]
//
// # Setting snapshot values
//
//   - [IGCControllerDirectionPad.SetValueForXAxisYAxis]: Sets the input values of a snapshot of a directional pad.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad
type IGCControllerDirectionPad interface {
	IGCControllerElement

	// Topic: Accessing values using the axes

	// The x-axis element of the directional pad.
	XAxis() IGCControllerAxisInput
	// The y-axis element of the directional pad.
	YAxis() IGCControllerAxisInput

	// Topic: Accessing values using directional buttons

	// The button element that changes the positive x-axis.
	Right() IGCControllerButtonInput
	// The button element that changes the negative x-axis.
	Left() IGCControllerButtonInput
	// The button element that changes the positive y-axis.
	Up() IGCControllerButtonInput
	// The button element used for the negative y-axis direction.
	Down() IGCControllerButtonInput

	// Topic: Getting change information

	// The block that the directional pad calls when the user changes its values.
	ValueChangedHandler() GCControllerDirectionPadFloat32Float32Handler
	SetValueChangedHandler(value GCControllerDirectionPadFloat32Float32Handler)

	// Topic: Setting snapshot values

	// Sets the input values of a snapshot of a directional pad.
	SetValueForXAxisYAxis(xAxis float32, yAxis float32)
}

// Init initializes the instance.
func (g GCControllerDirectionPad) Init() GCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCControllerDirectionPad) Autorelease() GCControllerDirectionPad {
	rv := objc.Send[GCControllerDirectionPad](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCControllerDirectionPad creates a new GCControllerDirectionPad instance.
func NewGCControllerDirectionPad() GCControllerDirectionPad {
	class := getGCControllerDirectionPadClass()
	rv := objc.Send[GCControllerDirectionPad](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Sets the input values of a snapshot of a directional pad.
//
// xAxis: A normalized value of the x-axis ranging from `-1` to `1`.
//
// yAxis: A normalized value for the y-axis ranging from `-1` to `1`.
//
// # Discussion
//
// This method does nothing if the associated controller isn’t a snapshot
// (its [GCController.Snapshot] property is false`)`. Otherwise, this method
// sets the value of the direction pad’s buttons as well.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad/setValueForXAxis(_:yAxis:)
func (g GCControllerDirectionPad) SetValueForXAxisYAxis(xAxis float32, yAxis float32) {
	objc.Send[objc.ID](g.ID, objc.Sel("setValueForXAxis:yAxis:"), xAxis, yAxis)
}

// The x-axis element of the directional pad.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad/xAxis
func (g GCControllerDirectionPad) XAxis() IGCControllerAxisInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("xAxis"))
	return GCControllerAxisInputFromID(objc.ID(rv))
}

// The y-axis element of the directional pad.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad/yAxis
func (g GCControllerDirectionPad) YAxis() IGCControllerAxisInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("yAxis"))
	return GCControllerAxisInputFromID(objc.ID(rv))
}

// The button element that changes the positive x-axis.
//
// # Discussion
//
// The value of the `right` and `left` buttons are mutually exclusive because
// the user can only press one of these buttons at a time. Therefore, when the
// `right` button is nonzero, the `left` button is `0`.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad/right
func (g GCControllerDirectionPad) Right() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("right"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The button element that changes the negative x-axis.
//
// # Discussion
//
// The value of the `right` and `left` buttons are mutually exclusive because
// the user can only press one of these buttons at a time. Therefore, when the
// `left` button is nonzero, the `right` button is `0`.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad/left
func (g GCControllerDirectionPad) Left() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("left"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The button element that changes the positive y-axis.
//
// # Discussion
//
// The value of the `up` and `down` buttons are mutually exclusive because the
// user can only press one of these buttons at a time. Therefore, when the
// `up` button is nonzero, the `down` button is `0`.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad/up
func (g GCControllerDirectionPad) Up() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("up"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The button element used for the negative y-axis direction.
//
// # Discussion
//
// The value of the `up` and `down` buttons are mutually exclusive because the
// user can only press one of these buttons at a time. Therefore, when the
// `down` button is nonzero, the `up` button is `0`.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad/down
func (g GCControllerDirectionPad) Down() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("down"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The block that the directional pad calls when the user changes its values.
//
// # Discussion
//
// Set this handler to receive notifications when the user changes a direction
// value.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerDirectionPad/valueChangedHandler
func (g GCControllerDirectionPad) ValueChangedHandler() GCControllerDirectionPadFloat32Float32Handler {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("valueChangedHandler"))
	_ = rv
	return nil
}
func (g GCControllerDirectionPad) SetValueChangedHandler(value GCControllerDirectionPadFloat32Float32Handler) {
	block, cleanup := NewGCControllerDirectionPadFloat32Float32Block(value)
	defer cleanup()
	objc.Send[struct{}](g.ID, objc.Sel("setValueChangedHandler:"), block)
}
