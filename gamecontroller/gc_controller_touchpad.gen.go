// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [GCControllerTouchpad] class.
var (
	_GCControllerTouchpadClass     GCControllerTouchpadClass
	_GCControllerTouchpadClassOnce sync.Once
)

func getGCControllerTouchpadClass() GCControllerTouchpadClass {
	_GCControllerTouchpadClassOnce.Do(func() {
		_GCControllerTouchpadClass = GCControllerTouchpadClass{class: objc.GetClass("GCControllerTouchpad")}
	})
	return _GCControllerTouchpadClass
}

// GetGCControllerTouchpadClass returns the class object for GCControllerTouchpad.
func GetGCControllerTouchpadClass() GCControllerTouchpadClass {
	return getGCControllerTouchpadClass()
}

type GCControllerTouchpadClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCControllerTouchpadClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCControllerTouchpadClass) Alloc() GCControllerTouchpad {
	rv := objc.Send[GCControllerTouchpad](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// A control element that represents a touch event on a touchpad.
//
// # Overview
//
// A [GCControllerTouchpad] object provides the state of the touches and
// presses on a touchpad. This is a compound element with button and
// directional pad subelements.
//
// # Getting the subelements
//
//   - [GCControllerTouchpad.TouchSurface]: The element that represents the state of the user’s touch on the surface of the touchpad.
//   - [GCControllerTouchpad.Button]: The element that represents the button component on the touchpad.
//
// # Accessing the input values
//
//   - [GCControllerTouchpad.TouchState]: The state of the user’s touch on the surface of the touchpad.
//   - [GCControllerTouchpad.ReportsAbsoluteTouchSurfaceValues]: A Boolean value that determines whether the touch values are absolute or relative.
//   - [GCControllerTouchpad.SetReportsAbsoluteTouchSurfaceValues]
//
// # Getting change information
//
//   - [GCControllerTouchpad.TouchDown]: The block that the element calls when the user begins touching the touchpad.
//   - [GCControllerTouchpad.SetTouchDown]
//   - [GCControllerTouchpad.TouchMoved]: The block that the element calls when the user continues touching the touchpad, not when the user begins or ends touching the touchpad.
//   - [GCControllerTouchpad.SetTouchMoved]
//   - [GCControllerTouchpad.TouchUp]: The block that the element calls when the user finishes touching the touchpad.
//   - [GCControllerTouchpad.SetTouchUp]
//
// # Setting snapshot values
//
//   - [GCControllerTouchpad.SetValueForXAxisYAxisTouchDownButtonValue]: Sets the input values of a snapshot of a touchpad.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerTouchpad
type GCControllerTouchpad struct {
	GCControllerElement
}

// GCControllerTouchpadFromID constructs a [GCControllerTouchpad] from an objc.ID.
//
// A control element that represents a touch event on a touchpad.
func GCControllerTouchpadFromID(id objc.ID) GCControllerTouchpad {
	return GCControllerTouchpad{GCControllerElement: GCControllerElementFromID(id)}
}

// NOTE: GCControllerTouchpad adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCControllerTouchpad] class.
//
// # Getting the subelements
//
//   - [IGCControllerTouchpad.TouchSurface]: The element that represents the state of the user’s touch on the surface of the touchpad.
//   - [IGCControllerTouchpad.Button]: The element that represents the button component on the touchpad.
//
// # Accessing the input values
//
//   - [IGCControllerTouchpad.TouchState]: The state of the user’s touch on the surface of the touchpad.
//   - [IGCControllerTouchpad.ReportsAbsoluteTouchSurfaceValues]: A Boolean value that determines whether the touch values are absolute or relative.
//   - [IGCControllerTouchpad.SetReportsAbsoluteTouchSurfaceValues]
//
// # Getting change information
//
//   - [IGCControllerTouchpad.TouchDown]: The block that the element calls when the user begins touching the touchpad.
//   - [IGCControllerTouchpad.SetTouchDown]
//   - [IGCControllerTouchpad.TouchMoved]: The block that the element calls when the user continues touching the touchpad, not when the user begins or ends touching the touchpad.
//   - [IGCControllerTouchpad.SetTouchMoved]
//   - [IGCControllerTouchpad.TouchUp]: The block that the element calls when the user finishes touching the touchpad.
//   - [IGCControllerTouchpad.SetTouchUp]
//
// # Setting snapshot values
//
//   - [IGCControllerTouchpad.SetValueForXAxisYAxisTouchDownButtonValue]: Sets the input values of a snapshot of a touchpad.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerTouchpad
type IGCControllerTouchpad interface {
	IGCControllerElement

	// Topic: Getting the subelements

	// The element that represents the state of the user’s touch on the surface of the touchpad.
	TouchSurface() IGCControllerDirectionPad
	// The element that represents the button component on the touchpad.
	Button() IGCControllerButtonInput

	// Topic: Accessing the input values

	// The state of the user’s touch on the surface of the touchpad.
	TouchState() GCTouchState
	// A Boolean value that determines whether the touch values are absolute or relative.
	ReportsAbsoluteTouchSurfaceValues() bool
	SetReportsAbsoluteTouchSurfaceValues(value bool)

	// Topic: Getting change information

	// The block that the element calls when the user begins touching the touchpad.
	TouchDown() GCControllerTouchpadHandler
	SetTouchDown(value GCControllerTouchpadHandler)
	// The block that the element calls when the user continues touching the touchpad, not when the user begins or ends touching the touchpad.
	TouchMoved() GCControllerTouchpadHandler
	SetTouchMoved(value GCControllerTouchpadHandler)
	// The block that the element calls when the user finishes touching the touchpad.
	TouchUp() GCControllerTouchpadHandler
	SetTouchUp(value GCControllerTouchpadHandler)

	// Topic: Setting snapshot values

	// Sets the input values of a snapshot of a touchpad.
	SetValueForXAxisYAxisTouchDownButtonValue(xAxis float32, yAxis float32, touchDown bool, buttonValue float32)
}

// Init initializes the instance.
func (g GCControllerTouchpad) Init() GCControllerTouchpad {
	rv := objc.Send[GCControllerTouchpad](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCControllerTouchpad) Autorelease() GCControllerTouchpad {
	rv := objc.Send[GCControllerTouchpad](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCControllerTouchpad creates a new GCControllerTouchpad instance.
func NewGCControllerTouchpad() GCControllerTouchpad {
	class := getGCControllerTouchpadClass()
	rv := objc.Send[GCControllerTouchpad](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Sets the input values of a snapshot of a touchpad.
//
// xAxis: A normalized value of the x-axis ranging from `-1` to `1`.
//
// yAxis: A normalized value of the y-axis ranging from `-1` to `1`.
//
// touchDown: A Boolean value that indicates whether the user starts touching the
// surface. If true, the user is touching the surface; otherwise, the user
// isn’t.
//
// buttonValue: A normalized number between `0.0` (minimum) and `1.0` (maximum) that
// represents the level of pressure the user applies to the button.
//
// # Discussion
//
// This method does nothing if the associated controller isn’t a snapshot
// (its [GCController.Snapshot] property is false`)`.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/setValueForXAxis(_:yAxis:touchDown:buttonValue:)
func (g GCControllerTouchpad) SetValueForXAxisYAxisTouchDownButtonValue(xAxis float32, yAxis float32, touchDown bool, buttonValue float32) {
	objc.Send[objc.ID](g.ID, objc.Sel("setValueForXAxis:yAxis:touchDown:buttonValue:"), xAxis, yAxis, touchDown, buttonValue)
}

// The element that represents the state of the user’s touch on the surface
// of the touchpad.
//
// # Discussion
//
// This element provides the recent or last touch positions on the two axes.
// Use the [GCControllerTouchpad.TouchState] property to determine whether the
// user is currently touching the surface.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/touchSurface
func (g GCControllerTouchpad) TouchSurface() IGCControllerDirectionPad {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("touchSurface"))
	return GCControllerDirectionPadFromID(objc.ID(rv))
}

// The element that represents the button component on the touchpad.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/button
func (g GCControllerTouchpad) Button() IGCControllerButtonInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("button"))
	return GCControllerButtonInputFromID(objc.ID(rv))
}

// The state of the user’s touch on the surface of the touchpad.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/touchState-swift.property
func (g GCControllerTouchpad) TouchState() GCTouchState {
	rv := objc.Send[GCTouchState](g.ID, objc.Sel("touchState"))
	return GCTouchState(rv)
}

// A Boolean value that determines whether the touch values are absolute or
// relative.
//
// # Discussion
//
// If this property is true, the touch values are absolute on the surface of
// the touchpad. If this property is false, the touch values are relative to
// the first touch on a virtual directional pad. The default value for this
// property is true.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/reportsAbsoluteTouchSurfaceValues
func (g GCControllerTouchpad) ReportsAbsoluteTouchSurfaceValues() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("reportsAbsoluteTouchSurfaceValues"))
	return rv
}
func (g GCControllerTouchpad) SetReportsAbsoluteTouchSurfaceValues(value bool) {
	objc.Send[struct{}](g.ID, objc.Sel("setReportsAbsoluteTouchSurfaceValues:"), value)
}

// The block that the element calls when the user begins touching the
// touchpad.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/touchDown
func (g GCControllerTouchpad) TouchDown() GCControllerTouchpadHandler {
	rv := objc.Send[GCControllerTouchpadHandler](g.ID, objc.Sel("touchDown"))
	return GCControllerTouchpadHandler(rv)
}
func (g GCControllerTouchpad) SetTouchDown(value GCControllerTouchpadHandler) {
	objc.Send[struct{}](g.ID, objc.Sel("setTouchDown:"), value)
}

// The block that the element calls when the user continues touching the
// touchpad, not when the user begins or ends touching the touchpad.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/touchMoved
func (g GCControllerTouchpad) TouchMoved() GCControllerTouchpadHandler {
	rv := objc.Send[GCControllerTouchpadHandler](g.ID, objc.Sel("touchMoved"))
	return GCControllerTouchpadHandler(rv)
}
func (g GCControllerTouchpad) SetTouchMoved(value GCControllerTouchpadHandler) {
	objc.Send[struct{}](g.ID, objc.Sel("setTouchMoved:"), value)
}

// The block that the element calls when the user finishes touching the
// touchpad.
//
// # Discussion
//
// The element invokes this handler when the user removes their fingers from
// the touchpad.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/touchUp
func (g GCControllerTouchpad) TouchUp() GCControllerTouchpadHandler {
	rv := objc.Send[GCControllerTouchpadHandler](g.ID, objc.Sel("touchUp"))
	return GCControllerTouchpadHandler(rv)
}
func (g GCControllerTouchpad) SetTouchUp(value GCControllerTouchpadHandler) {
	objc.Send[struct{}](g.ID, objc.Sel("setTouchUp:"), value)
}
