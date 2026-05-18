// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GCKeyboard] class.
var (
	_GCKeyboardClass     GCKeyboardClass
	_GCKeyboardClassOnce sync.Once
)

func getGCKeyboardClass() GCKeyboardClass {
	_GCKeyboardClassOnce.Do(func() {
		_GCKeyboardClass = GCKeyboardClass{class: objc.GetClass("GCKeyboard")}
	})
	return _GCKeyboardClass
}

// GetGCKeyboardClass returns the class object for GCKeyboard.
func GetGCKeyboardClass() GCKeyboardClass {
	return getGCKeyboardClass()
}

type GCKeyboardClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCKeyboardClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCKeyboardClass) Alloc() GCKeyboard {
	rv := objc.Send[GCKeyboard](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a physical keyboard connected to a device.
//
// # Overview
//
// To get the keyboard object and its input values, register for the
// [GCKeyboard.GCKeyboardDidConnect] (Swift) or [GCKeyboardDidConnectNotification]
// (Objective-C) notification for when a keyboard connects to the device, or
// use the [GCKeyboard.CoalescedKeyboard] class property. Then get the input values from
// the keyboard object’s [GCKeyboard.KeyboardInput] controller profile.
//
// # Discovering keyboards
//
//   - [GCKeyboard.GCKeyboardDidConnect]: A notification that posts after a keyboard connects to the device.
//   - [GCKeyboard.GCKeyboardDidDisconnect]: A notification that posts after a single keyboard, or the last of multiple keyboards, disconnects from the device.
//
// # Getting input values
//
//   - [GCKeyboard.KeyboardInput]: The controller profile for the keyboard.
//
// See: https://developer.apple.com/documentation/GameController/GCKeyboard
//
// [GCKeyboardDidConnectNotification]: https://developer.apple.com/documentation/GameController/GCKeyboardDidConnectNotification
type GCKeyboard struct {
	objectivec.Object
}

// GCKeyboardFromID constructs a [GCKeyboard] from an objc.ID.
//
// An object that represents a physical keyboard connected to a device.
func GCKeyboardFromID(id objc.ID) GCKeyboard {
	return GCKeyboard{objectivec.Object{ID: id}}
}

// NOTE: GCKeyboard adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCKeyboard] class.
//
// # Discovering keyboards
//
//   - [IGCKeyboard.GCKeyboardDidConnect]: A notification that posts after a keyboard connects to the device.
//   - [IGCKeyboard.GCKeyboardDidDisconnect]: A notification that posts after a single keyboard, or the last of multiple keyboards, disconnects from the device.
//
// # Getting input values
//
//   - [IGCKeyboard.KeyboardInput]: The controller profile for the keyboard.
//
// See: https://developer.apple.com/documentation/GameController/GCKeyboard
type IGCKeyboard interface {
	objectivec.IObject

	// Topic: Discovering keyboards

	// A notification that posts after a keyboard connects to the device.
	GCKeyboardDidConnect() foundation.NSString
	// A notification that posts after a single keyboard, or the last of multiple keyboards, disconnects from the device.
	GCKeyboardDidDisconnect() foundation.NSString

	// Topic: Getting input values

	// The controller profile for the keyboard.
	KeyboardInput() IGCKeyboardInput
}

// Init initializes the instance.
func (g GCKeyboard) Init() GCKeyboard {
	rv := objc.Send[GCKeyboard](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCKeyboard) Autorelease() GCKeyboard {
	rv := objc.Send[GCKeyboard](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCKeyboard creates a new GCKeyboard instance.
func NewGCKeyboard() GCKeyboard {
	class := getGCKeyboardClass()
	rv := objc.Send[GCKeyboard](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// The dispatch queue that the framework uses to call element value change
// handlers.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice/handlerQueue
func (g GCKeyboard) HandlerQueue() dispatch.Queue {
	rv := objc.Send[uintptr](g.ID, objc.Sel("handlerQueue"))
	return dispatch.QueueFromHandle(rv)
}

// The device’s physical input profile, such as a controller’s extended
// gamepad.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice/physicalInputProfile
func (g GCKeyboard) PhysicalInputProfile() IGCPhysicalInputProfile {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("physicalInputProfile"))
	return GCPhysicalInputProfileFromID(rv)
}

// The product category that identifies the type of controller.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice/productCategory
func (g GCKeyboard) ProductCategory() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("productCategory"))
	return foundation.NSStringFromID(rv).String()
}

// The manufacturer-provided name for the device, or the user’s name for the
// device.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice/vendorName
func (g GCKeyboard) VendorName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("vendorName"))
	return foundation.NSStringFromID(rv).String()
}

// A notification that posts after a keyboard connects to the device.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/GCKeyboardDidConnect
func (g GCKeyboard) GCKeyboardDidConnect() foundation.NSString {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("GCKeyboardDidConnect"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// A notification that posts after a single keyboard, or the last of multiple
// keyboards, disconnects from the device.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/GCKeyboardDidDisconnect
func (g GCKeyboard) GCKeyboardDidDisconnect() foundation.NSString {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("GCKeyboardDidDisconnect"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// The controller profile for the keyboard.
//
// # Discussion
//
// Use this object to get the keyboard’s buttons and button states.
//
// See: https://developer.apple.com/documentation/GameController/GCKeyboard/keyboardInput
func (g GCKeyboard) KeyboardInput() IGCKeyboardInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("keyboardInput"))
	return GCKeyboardInputFromID(objc.ID(rv))
}

// The keyboard currently connected to the device.
//
// # Discussion
//
// Get the keyboard input values from the keyboard’s [KeyboardInput]
// controller profile. If the user connects more than one keyboard, the
// framework represents the combined keyboards with one coalesced keyboard
// object.
//
// See: https://developer.apple.com/documentation/GameController/GCKeyboard/coalesced
func (_GCKeyboardClass GCKeyboardClass) CoalescedKeyboard() GCKeyboard {
	rv := objc.Send[objc.ID](objc.ID(_GCKeyboardClass.class), objc.Sel("coalescedKeyboard"))
	return GCKeyboardFromID(objc.ID(rv))
}

// Protocol methods for GCDevice

// The dispatch queue that the framework uses to call element value change
// handlers.
//
// # Discussion
//
// The default queue is the main queue. Set this property to another queue to
// asynchronously call value change handlers (see [GCControllerAxisInput],
// [GCControllerButtonInput], [GCControllerDirectionPad], and [GCMotion]). For
// example, if you handle input on another queue, set this property when you
// first access the input device.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice/handlerQueue
func (o GCKeyboard) SetHandlerQueue(value dispatch.Queue) {
	objc.Send[struct{}](o.ID, objc.Sel("setHandlerQueue:"), value)
}
