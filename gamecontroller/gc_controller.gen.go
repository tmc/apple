// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"context"
	"sync"

	"github.com/tmc/apple/dispatch"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/iokit"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GCController] class.
var (
	_GCControllerClass     GCControllerClass
	_GCControllerClassOnce sync.Once
)

func getGCControllerClass() GCControllerClass {
	_GCControllerClassOnce.Do(func() {
		_GCControllerClass = GCControllerClass{class: objc.GetClass("GCController")}
	})
	return _GCControllerClass
}

// GetGCControllerClass returns the class object for GCController.
func GetGCControllerClass() GCControllerClass {
	return getGCControllerClass()
}

type GCControllerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCControllerClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCControllerClass) Alloc() GCController {
	rv := objc.Send[GCController](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// A representation of a real game controller, a virtual controller, or a
// snapshot of a controller.
//
// # Overview
//
// This class represents a real or virtual controller that a user interacts
// with during a game. A is a physical controller that connects directly or
// wirelessly to the device. A real controller can be formfitting or can
// attach closely to a device so players can use controls on both
// simultaneously. A is a software emulation of a real controller.
//
// You discover controllers, and then you process the input from those
// controllers during gameplay. Use the [GCController.Controllers] method to get the
// currently connected controllers. If necessary, use the
// [GCController.StartWirelessControllerDiscoveryWithCompletionHandler] method to connect
// with wireless controllers.
//
// This framework supports multiple connected game controllers. To identify
// which player is using a controller in a multiplayer game, check the
// [GCController.PlayerIndex] property and set it, if necessary. For single-player games,
// use the [GCController.Current] property to get the controller that the player is
// actively using.
//
// A controller’s profile encapsulates the details about a controller’s
// buttons, pads, axis, and other input elements. Get the controller’s
// profile using one of the profile properties, such as [GCController.ExtendedGamepad], and
// then process the input from its elements.
//
// You can either get the values of input elements on each iteration of your
// game loop, or set handlers to receive callbacks when those values change.
// For example, use the [GCController.LeftThumbstick] property of the [GCExtendedGamepad]
// profile to get the thumbstick state. Use the [GCController.ValueChangedHandler] property
// to set a handler that you implement to process any input values that change
// in the profile.
//
// Alternatively, you can create a snapshot of a real or virtual controller
// using the [GCController.Capture] method. A is a copy of a controller at a moment in time
// with its current element values. Creating a snapshot may impact
// performance, and over time a snapshot doesn’t stay current. Unlike other
// types of controllers, you can set the values of elements in a snapshot.
//
// # Discovering controllers
//
//   - [GCController.GCControllerDidConnect]: A notification that posts after a controller connects to the device.
//   - [GCController.GCControllerDidDisconnect]: A notification that posts after a controller disconnects from the device.
//
// # Handling multiple controllers
//
//   - [GCController.GCControllerDidBecomeCurrent]: A notification that posts when a controller becomes the current controller.
//   - [GCController.GCControllerDidStopBeingCurrent]: A notification that posts when a controller stops being the current controller.
//
// # Inspecting a controller
//
//   - [GCController.IsAttachedToDevice]: A Boolean value that indicates whether the controller closely integrates with the device.
//
// # Accessing controller input
//
//   - [GCController.Input]: The input profile for the controller.
//
// # Accessing controller profiles
//
//   - [GCController.ExtendedGamepad]: The extended gamepad profile.
//   - [GCController.MicroGamepad]: The micro gamepad profile.
//   - [GCController.Motion]: The motion input profile.
//
// # Identifying controllers and displaying a player index
//
//   - [GCController.PlayerIndex]: The player index for the controller.
//   - [GCController.SetPlayerIndex]
//
// # Accessing battery, haptics, and light objects
//
//   - [GCController.Battery]: The controller’s battery information.
//   - [GCController.Haptics]: The controller’s haptics information.
//   - [GCController.Light]: The controller’s light settings.
//
// # Creating snapshots
//
//   - [GCController.Capture]: Returns a snapshot of the controller with its current element values.
//   - [GCController.IsSnapshot]: A Boolean value that indicates whether the controller is a snapshot of a controller.
//
// See: https://developer.apple.com/documentation/GameController/GCController
type GCController struct {
	objectivec.Object
}

// GCControllerFromID constructs a [GCController] from an objc.ID.
//
// A representation of a real game controller, a virtual controller, or a
// snapshot of a controller.
func GCControllerFromID(id objc.ID) GCController {
	return GCController{objectivec.Object{ID: id}}
}

// NOTE: GCController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCController] class.
//
// # Discovering controllers
//
//   - [IGCController.GCControllerDidConnect]: A notification that posts after a controller connects to the device.
//   - [IGCController.GCControllerDidDisconnect]: A notification that posts after a controller disconnects from the device.
//
// # Handling multiple controllers
//
//   - [IGCController.GCControllerDidBecomeCurrent]: A notification that posts when a controller becomes the current controller.
//   - [IGCController.GCControllerDidStopBeingCurrent]: A notification that posts when a controller stops being the current controller.
//
// # Inspecting a controller
//
//   - [IGCController.IsAttachedToDevice]: A Boolean value that indicates whether the controller closely integrates with the device.
//
// # Accessing controller input
//
//   - [IGCController.Input]: The input profile for the controller.
//
// # Accessing controller profiles
//
//   - [IGCController.ExtendedGamepad]: The extended gamepad profile.
//   - [IGCController.MicroGamepad]: The micro gamepad profile.
//   - [IGCController.Motion]: The motion input profile.
//
// # Identifying controllers and displaying a player index
//
//   - [IGCController.PlayerIndex]: The player index for the controller.
//   - [IGCController.SetPlayerIndex]
//
// # Accessing battery, haptics, and light objects
//
//   - [IGCController.Battery]: The controller’s battery information.
//   - [IGCController.Haptics]: The controller’s haptics information.
//   - [IGCController.Light]: The controller’s light settings.
//
// # Creating snapshots
//
//   - [IGCController.Capture]: Returns a snapshot of the controller with its current element values.
//   - [IGCController.IsSnapshot]: A Boolean value that indicates whether the controller is a snapshot of a controller.
//
// See: https://developer.apple.com/documentation/GameController/GCController
type IGCController interface {
	objectivec.IObject

	// Topic: Discovering controllers

	// A notification that posts after a controller connects to the device.
	GCControllerDidConnect() foundation.NSString
	// A notification that posts after a controller disconnects from the device.
	GCControllerDidDisconnect() foundation.NSString

	// Topic: Handling multiple controllers

	// A notification that posts when a controller becomes the current controller.
	GCControllerDidBecomeCurrent() foundation.NSString
	// A notification that posts when a controller stops being the current controller.
	GCControllerDidStopBeingCurrent() foundation.NSString

	// Topic: Inspecting a controller

	// A Boolean value that indicates whether the controller closely integrates with the device.
	IsAttachedToDevice() bool

	// Topic: Accessing controller input

	// The input profile for the controller.
	Input() IGCControllerLiveInput

	// Topic: Accessing controller profiles

	// The extended gamepad profile.
	ExtendedGamepad() IGCExtendedGamepad
	// The micro gamepad profile.
	MicroGamepad() IGCMicroGamepad
	// The motion input profile.
	Motion() IGCMotion

	// Topic: Identifying controllers and displaying a player index

	// The player index for the controller.
	PlayerIndex() GCControllerPlayerIndex
	SetPlayerIndex(value GCControllerPlayerIndex)

	// Topic: Accessing battery, haptics, and light objects

	// The controller’s battery information.
	Battery() IGCDeviceBattery
	// The controller’s haptics information.
	Haptics() IGCDeviceHaptics
	// The controller’s light settings.
	Light() IGCDeviceLight

	// Topic: Creating snapshots

	// Returns a snapshot of the controller with its current element values.
	Capture() IGCController
	// A Boolean value that indicates whether the controller is a snapshot of a controller.
	IsSnapshot() bool

	// The controller’s left thumbstick element.
	LeftThumbstick() IGCControllerDirectionPad
	SetLeftThumbstick(value IGCControllerDirectionPad)
	// The block that the profile calls when an element’s value changes.
	ValueChangedHandler() GCExtendedGamepadValueChangedHandler
	SetValueChangedHandler(value GCExtendedGamepadValueChangedHandler)
}

// Init initializes the instance.
func (g GCController) Init() GCController {
	rv := objc.Send[GCController](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCController) Autorelease() GCController {
	rv := objc.Send[GCController](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCController creates a new GCController instance.
func NewGCController() GCController {
	class := getGCControllerClass()
	rv := objc.Send[GCController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a snapshot of the controller with its current element values.
//
// # Return Value
//
// A snapshot of the controller.
//
// # Discussion
//
// A snapshot is a copy of a real or virtual controller at a moment in time
// with its current element values. Unlike other controllers, you can set the
// values of a snapshot’s [GCControllerElement] objects.
//
// See: https://developer.apple.com/documentation/GameController/GCController/capture()
func (g GCController) Capture() IGCController {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("capture"))
	return GCControllerFromID(rv)
}

// The dispatch queue that the framework uses to call element value change
// handlers.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice/handlerQueue
func (g GCController) HandlerQueue() dispatch.Queue {
	rv := objc.Send[uintptr](g.ID, objc.Sel("handlerQueue"))
	return dispatch.QueueFromHandle(rv)
}

// The product category that identifies the type of controller.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice/productCategory
func (g GCController) ProductCategory() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("productCategory"))
	return foundation.NSStringFromID(rv).String()
}

// The manufacturer-provided name for the device, or the user’s name for the
// device.
//
// See: https://developer.apple.com/documentation/GameController/GCDevice/vendorName
func (g GCController) VendorName() string {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("vendorName"))
	return foundation.NSStringFromID(rv).String()
}

// Returns the connected controllers for the device.
//
// # Return Value
//
// The currently connected controllers.
//
// # Discussion
//
// To track the connection status of controllers, observe the framework
// notifications. The framework posts the [GCControllerDidConnect] (Swift) and
// [GCControllerDidBecomeCurrent] (Swift) notifications when a controller
// connects to a device. For Objective-C, it posts the
// [GCControllerDidConnectNotification] and
// [GCControllerDidBecomeCurrentNotification] notifications. When a controller
// disconnects from a device, it posts the [GCControllerDidDisconnect] (Swift)
// and [GCControllerDidStopBeingCurrent] (Swift) notifications. For
// Objective-C, it posts the [GCControllerDidDisconnectNotification] and
// [GCControllerDidStopBeingCurrentNotification] notifications.
//
// See: https://developer.apple.com/documentation/GameController/GCController/controllers()
//
// [GCControllerDidBecomeCurrentNotification]: https://developer.apple.com/documentation/GameController/GCControllerDidBecomeCurrentNotification
// [GCControllerDidConnectNotification]: https://developer.apple.com/documentation/GameController/GCControllerDidConnectNotification
// [GCControllerDidDisconnectNotification]: https://developer.apple.com/documentation/GameController/GCControllerDidDisconnectNotification
// [GCControllerDidStopBeingCurrentNotification]: https://developer.apple.com/documentation/GameController/GCControllerDidStopBeingCurrentNotification
func (_GCControllerClass GCControllerClass) Controllers() []GCController {
	rv := objc.Send[[]objc.ID](objc.ID(_GCControllerClass.class), objc.Sel("controllers"))
	return objc.ConvertSlice(rv, func(id objc.ID) GCController {
		return GCControllerFromID(id)
	})
}

// Starts searching for nearby wireless controllers.
//
// completionHandler: The block that the framework calls when it completes the request.
//
// # Discussion
//
// Call this method when the user chooses to discover wireless controllers
// from your interface. The framework searches asynchronously for discoverable
// wireless controllers. The framework posts the [GCControllerDidConnect]
// (Swift) or [GCControllerDidConnectNotification] (Objective-C) notification
// when it discovers new controllers. Implement the completion handler you
// pass to this method to handle when the framework finishes discovering
// controllers or when it times out.
//
// If you call the [StartWirelessControllerDiscoveryWithCompletionHandler]
// method multiple times during discovery, the framework only calls the last
// completion handler you pass to this method.
//
// See: https://developer.apple.com/documentation/GameController/GCController/startWirelessControllerDiscovery(completionHandler:)
//
// [GCControllerDidConnectNotification]: https://developer.apple.com/documentation/GameController/GCControllerDidConnectNotification
func (_GCControllerClass GCControllerClass) StartWirelessControllerDiscoveryWithCompletionHandler(completionHandler VoidHandler) {
	_block0, _ := NewVoidBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_GCControllerClass.class), objc.Sel("startWirelessControllerDiscoveryWithCompletionHandler:"), _block0)
}

// Stops searching for nearby wireless controllers.
//
// # Discussion
//
// If you call this method while the framework searches for wireless
// controllers, the framework stops searching and invokes the completion
// handler you pass to the
// [StartWirelessControllerDiscoveryWithCompletionHandler] method.
//
// See: https://developer.apple.com/documentation/GameController/GCController/stopWirelessControllerDiscovery()
func (_GCControllerClass GCControllerClass) StopWirelessControllerDiscovery() {
	objc.Send[objc.ID](objc.ID(_GCControllerClass.class), objc.Sel("stopWirelessControllerDiscovery"))
}

// Returns a Boolean value that indicates whether the framework supports the
// specified human interface device.
//
// device: A human interface input device.
//
// # Return Value
//
// true if the framework supports the device; otherwise, false.
//
// # Discussion
//
// If the Game Controller framework supports the input device, you can use the
// Game Controller APIs to interact with the device instead of the IOKit APIs.
//
// See: https://developer.apple.com/documentation/GameController/GCController/supportsHIDDevice(_:)
func (_GCControllerClass GCControllerClass) SupportsHIDDevice(device iokit.IOHIDDeviceRef) bool {
	rv := objc.Send[bool](objc.ID(_GCControllerClass.class), objc.Sel("supportsHIDDevice:"), device)
	return rv
}

// Returns a snapshot of a newly created controller with an extended gamepad
// profile.
//
// # Return Value
//
// A snapshot with an extended gamepad profile.
//
// # Discussion
//
// A snapshot is a copy of a real or virtual controller at a moment in time
// with its current element values. Unlike other controllers, you can set the
// values of a snapshot’s [GCControllerElement] objects.
//
// See: https://developer.apple.com/documentation/GameController/GCController/withExtendedGamepad()
func (_GCControllerClass GCControllerClass) ControllerWithExtendedGamepad() GCController {
	rv := objc.Send[objc.ID](objc.ID(_GCControllerClass.class), objc.Sel("controllerWithExtendedGamepad"))
	return GCControllerFromID(rv)
}

// Returns a snapshot of a newly created controller with a micro gamepad
// profile.
//
// # Return Value
//
// A snapshot with a micro gamepad profile.
//
// # Discussion
//
// A snapshot is a copy of a real or virtual controller at a moment in time
// with its current element values. Unlike other controllers, you can set the
// values of a snapshot’s [GCControllerElement] objects.
//
// See: https://developer.apple.com/documentation/GameController/GCController/withMicroGamepad()
func (_GCControllerClass GCControllerClass) ControllerWithMicroGamepad() GCController {
	rv := objc.Send[objc.ID](objc.ID(_GCControllerClass.class), objc.Sel("controllerWithMicroGamepad"))
	return GCControllerFromID(rv)
}

// A notification that posts after a controller connects to the device.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/GCControllerDidConnect
func (g GCController) GCControllerDidConnect() foundation.NSString {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("GCControllerDidConnect"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// A notification that posts after a controller disconnects from the device.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/GCControllerDidDisconnect
func (g GCController) GCControllerDidDisconnect() foundation.NSString {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("GCControllerDidDisconnect"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// A notification that posts when a controller becomes the current controller.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/GCControllerDidBecomeCurrent
func (g GCController) GCControllerDidBecomeCurrent() foundation.NSString {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("GCControllerDidBecomeCurrent"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// A notification that posts when a controller stops being the current
// controller.
//
// See: https://developer.apple.com/documentation/Foundation/NSNotification/Name-swift.struct/GCControllerDidStopBeingCurrent
func (g GCController) GCControllerDidStopBeingCurrent() foundation.NSString {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("GCControllerDidStopBeingCurrent"))
	return foundation.NSStringFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the controller closely integrates
// with the device.
//
// # Discussion
//
// If true, the controller may be formfitting or otherwise closely attach to
// the device so that the player can interact simultaneously with the
// controller and the device. If false, the controller doesn’t have an
// attachment to the device.
//
// See: https://developer.apple.com/documentation/GameController/GCController/isAttachedToDevice
func (g GCController) IsAttachedToDevice() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isAttachedToDevice"))
	return rv
}

// The input profile for the controller.
//
// See: https://developer.apple.com/documentation/GameController/GCController/input
func (g GCController) Input() IGCControllerLiveInput {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("input"))
	return GCControllerLiveInputFromID(objc.ID(rv))
}

// The extended gamepad profile.
//
// # Discussion
//
// If the controller supports the extended gamepad profile, this property is a
// [GCExtendedGamepad] object that you use to access the input elements of the
// controller. If the controller doesn’t support the extended gamepad
// profile, this property is `nil`.
//
// See: https://developer.apple.com/documentation/GameController/GCController/extendedGamepad
func (g GCController) ExtendedGamepad() IGCExtendedGamepad {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("extendedGamepad"))
	return GCExtendedGamepadFromID(objc.ID(rv))
}

// The micro gamepad profile.
//
// # Discussion
//
// If the controller supports the micro gamepad profile, this property is a
// [GCMicroGamepad] object that you use to access the input elements of the
// controller. If the controller doesn’t support the micro gamepad profile,
// this property is `nil`.
//
// See: https://developer.apple.com/documentation/GameController/GCController/microGamepad
func (g GCController) MicroGamepad() IGCMicroGamepad {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("microGamepad"))
	return GCMicroGamepadFromID(objc.ID(rv))
}

// The motion input profile.
//
// # Discussion
//
// If the controller supports the motion profile, this property is a
// [GCMotion] object that you use to access the controller’s motion data. If
// the controller doesn’t support the motion input profile, this property is
// `nil`.
//
// See: https://developer.apple.com/documentation/GameController/GCController/motion
func (g GCController) Motion() IGCMotion {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("motion"))
	return GCMotionFromID(objc.ID(rv))
}

// The physical input profile for the controller.
//
// # Discussion
//
// This is a convenience property that returns the [ExtendedGamepad] or
// [MicroGamepad] properties.
//
// See: https://developer.apple.com/documentation/GameController/GCController/physicalInputProfile
func (g GCController) PhysicalInputProfile() IGCPhysicalInputProfile {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("physicalInputProfile"))
	return GCPhysicalInputProfileFromID(objc.ID(rv))
}

// The player index for the controller.
//
// # Discussion
//
// Use the player index to identify which player is using the controller. Set
// the player index when the controller first connects to the device and you
// configure your game.
//
// When you set the player index, the matching LED on the controller for that
// player lights up. You don’t need to provide a unique player index for
// each active game controller. For example, players on the same team can
// share a common player index. If your game no longer uses a controller, set
// the controller’s index value to [GCControllerPlayerIndexUnset].
//
// The default value for this property is [GCControllerPlayerIndexUnset].
//
// See: https://developer.apple.com/documentation/GameController/GCController/playerIndex
func (g GCController) PlayerIndex() GCControllerPlayerIndex {
	rv := objc.Send[GCControllerPlayerIndex](g.ID, objc.Sel("playerIndex"))
	return GCControllerPlayerIndex(rv)
}
func (g GCController) SetPlayerIndex(value GCControllerPlayerIndex) {
	objc.Send[struct{}](g.ID, objc.Sel("setPlayerIndex:"), value)
}

// The controller’s battery information.
//
// # Discussion
//
// Use this property to display the battery life or to warn the user when the
// controller’s battery level is low. If the controller doesn’t provide
// battery information, this property is `nil`.
//
// See: https://developer.apple.com/documentation/GameController/GCController/battery
func (g GCController) Battery() IGCDeviceBattery {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("battery"))
	return GCDeviceBatteryFromID(objc.ID(rv))
}

// The controller’s haptics information.
//
// # Discussion
//
// Use this property to create [CHHapticEngine] instances as necessary in your
// app. If the controller doesn’t provide haptics information, this property
// is `nil`.
//
// See: https://developer.apple.com/documentation/GameController/GCController/haptics
//
// [CHHapticEngine]: https://developer.apple.com/documentation/CoreHaptics/CHHapticEngine
func (g GCController) Haptics() IGCDeviceHaptics {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("haptics"))
	return GCDeviceHapticsFromID(objc.ID(rv))
}

// The controller’s light settings.
//
// # Discussion
//
// Use the light settings to signal the user or to create a more immersive
// experience. If the controller doesn’t provide light settings, this
// property is `nil`.
//
// See: https://developer.apple.com/documentation/GameController/GCController/light
func (g GCController) Light() IGCDeviceLight {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("light"))
	return GCDeviceLightFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the controller is a snapshot of a
// controller.
//
// # Discussion
//
// If true, the controller is a snapshot of a controller. A snapshot is a copy
// of a real or virtual controller at a moment in time with its current
// element values. If false, the controller is a real or virtual controller.
//
// See: https://developer.apple.com/documentation/GameController/GCController/isSnapshot
func (g GCController) IsSnapshot() bool {
	rv := objc.Send[bool](g.ID, objc.Sel("isSnapshot"))
	return rv
}

// The controller’s left thumbstick element.
//
// See: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/leftthumbstick
func (g GCController) LeftThumbstick() IGCControllerDirectionPad {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("leftThumbstick"))
	return GCControllerDirectionPadFromID(objc.ID(rv))
}
func (g GCController) SetLeftThumbstick(value IGCControllerDirectionPad) {
	objc.Send[struct{}](g.ID, objc.Sel("setLeftThumbstick:"), value)
}

// The block that the profile calls when an element’s value changes.
//
// See: https://developer.apple.com/documentation/gamecontroller/gcextendedgamepad/valuechangedhandler
func (g GCController) ValueChangedHandler() GCExtendedGamepadValueChangedHandler {
	rv := objc.Send[GCExtendedGamepadValueChangedHandler](g.ID, objc.Sel("valueChangedHandler"))
	return GCExtendedGamepadValueChangedHandler(rv)
}
func (g GCController) SetValueChangedHandler(value GCExtendedGamepadValueChangedHandler) {
	objc.Send[struct{}](g.ID, objc.Sel("setValueChangedHandler:"), value)
}

// The most recently used game controller.
//
// # Discussion
//
// Use this property for a single-player game when you don’t need to
// distinguish the input from multiple controllers simultaneously.
//
// See: https://developer.apple.com/documentation/GameController/GCController/current
func (_GCControllerClass GCControllerClass) Current() GCController {
	rv := objc.Send[objc.ID](objc.ID(_GCControllerClass.class), objc.Sel("current"))
	return GCControllerFromID(objc.ID(rv))
}

// A Boolean value that indicates whether the app needs to respond to
// controller events when it isn’t the frontmost app.
//
// # Discussion
//
// If false, and the app isn’t in the foreground, the framework doesn’t
// forward any input from the game controller until the app becomes the
// frontmost.
//
// See: https://developer.apple.com/documentation/GameController/GCController/shouldMonitorBackgroundEvents
func (_GCControllerClass GCControllerClass) ShouldMonitorBackgroundEvents() bool {
	rv := objc.Send[bool](objc.ID(_GCControllerClass.class), objc.Sel("shouldMonitorBackgroundEvents"))
	return rv
}
func (_GCControllerClass GCControllerClass) SetShouldMonitorBackgroundEvents(value bool) {
	objc.Send[struct{}](objc.ID(_GCControllerClass.class), objc.Sel("setShouldMonitorBackgroundEvents:"), value)
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
func (o GCController) SetHandlerQueue(value dispatch.Queue) {
	objc.Send[struct{}](o.ID, objc.Sel("setHandlerQueue:"), value)
}

// StartWirelessControllerDiscovery is a synchronous wrapper around [GCController.StartWirelessControllerDiscoveryWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (gc GCControllerClass) StartWirelessControllerDiscovery(ctx context.Context) error {
	done := make(chan struct{}, 1)
	gc.StartWirelessControllerDiscoveryWithCompletionHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
