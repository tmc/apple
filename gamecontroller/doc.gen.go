// Code generated from Apple documentation for GameController. DO NOT EDIT.

// Package gamecontroller provides Go bindings for the GameController framework.
//
// Support hardware game controllers in your game.
//
// Use Game Controller to support users interacting with your app using a
// physical or virtual game controller. Game controllers include third-party
// products, such as the DualShock 4, DualSense, and Xbox, as well as the
// mouse, keyboard, Siri Remote, and racing wheels.
//
// # Essentials
//
//   - [Game Controller updates]: Learn about important changes to Game Controller.
//   - [Discovering game controllers]: Implement connection and input handling to provide seamless physical controller support for players.
//   - [Handling input events]: Receive controller input using either polling or callbacks.
//
// # View controller
//
//   - [GCEventViewController]: A view controller that delivers input either from the responder chain to views, or from game controllers to profiles.
//
// # Game controllers
//
//   - [Letting players use their second-generation Siri Remote as a game controller]: Support the second-generation Siri Remote as a game controller in your Apple TV game.
//   - [Discovering and tracking spatial game controllers and styli]: Receive controller and stylus input to interact with content in your augmented reality app.
//   - [GCDevice]: A protocol that defines a common interface for game input devices.
//   - [GCController]: A representation of a real game controller, a virtual controller, or a snapshot of a controller. ([GCControllerLiveInput], [GCControllerInputState], [GCPhysicalInputProfile], [GCKeyboardInput], [GCMouseInput])
//   - [GCRacingWheel]: An object that represents a physical racing wheel controller connected to a device.
//   - [GCKeyboard]: An object that represents a physical keyboard connected to a device.
//   - [GCMouse]: An object that represents a physical mouse connected to a device.
//
// # Game controller profiles
//
//   - Input: Receive controller input in the way that best integrates with the flow of your game or game engine. ([GCDevicePhysicalInput], [GCDevicePhysicalInputState], [GCDevicePhysicalInputStateDiff], [GCPhysicalInputElementCollection], [GCPhysicalInputElement])
//   - [GCMotion]: A controller profile that supports orientation and motion. ([GCMotionValueChangedHandler], [GCQuaternion], [GCRotationRate], [GCEulerAngles], [GCAcceleration])
//   - [GCDeviceBattery]: The charge level and state of a device’s battery.
//   - [GCDeviceHaptics]: The locations of haptic actuators on a game controller. ([GCHapticsLocality])
//   - [GCDeviceLight]: The colored light on a device. ([GCColor])
//
// # Virtual controller
//
//   - [Adding virtual controls to games that support game controllers in iOS]: Use touch input and virtual controllers to make your game available to players without controllers.
//
// # Button elements and names
//
//   - [GCTouchedStateInput]: The common properties for an element that has touch state input.
//   - [GCPressedStateInput]: The common properties for an element that has press state input, such as input from a button.
//   - [GCInputButtonName]: The Objective-C type for an input button name.
//
// # Racing wheels
//
//   - [Racing wheel device support]: Add support for racing wheel devices in macOS. ([GCRacingWheel], [GCRacingWheelInput], [GCRacingWheelInputState], [GCAxisInput], [GCGearShifterElement])
//
// # Game Controller framework migration from IOKit
//
//   - [Understanding game controller backward compatibility]: Learn how macOS brings support for the latest game controllers to software that predates the introduction of the Game Controller framework.
//
// # Protocols
//
//   - [GCPhysicalInputExtents]: Physical extents scale the normalized value reported by into physical units.
//
// # Variables
//
//   - [GCInputGripButton]
//   - [GCInputLeftSideButton]
//   - [GCInputRightSideButton]
//   - [GCInputThumbstick]
//   - [GCInputThumbstickButton]
//   - [GCInputTrigger]//
//
// # Key Types
//
//   - [GCController] - A representation of a real game controller, a virtual controller, or a snapshot of a controller.
//   - [GCControllerLiveInput] - The input profile for a controller.
//   - [GCRacingWheelInput] - A controller profile that supports a racing wheel.
//   - [GCExtendedGamepad] - A controller profile that supports the extended set of gamepad controls.
//   - [GCPhysicalInputProfile] - The base class for controller profiles that support physical buttons, thumbsticks, and directional pads.
//   - [GCMotion] - A controller profile that supports orientation and motion.
//   - [GCRacingWheelInputState] - The input for the wheel of a racing wheel controller.
//   - [GCRacingWheel] - An object that represents a physical racing wheel controller connected to a device.
//   - [GCControllerInputState] - A class that represents an input state for gamepads and arcade sticks.
//   - [GCSteeringWheelElement] - The element that represents the wheel of a racing wheel controller.
//
// [Adding virtual controls to games that support game controllers in iOS]: https://developer.apple.com/documentation/gamecontroller/adding-virtual-controls-to-games-that-support-game-controllers-in-ios
// [Discovering and tracking spatial game controllers and styli]: https://developer.apple.com/documentation/gamecontroller/discovering-and-tracking-spatial-game-controllers-and-styli
// [Discovering game controllers]: https://developer.apple.com/documentation/gamecontroller/discovering-game-controllers
// [Game Controller updates]: https://developer.apple.com/documentation/Updates/GameController
// [Handling input events]: https://developer.apple.com/documentation/gamecontroller/handling-input-events
// [Letting players use their second-generation Siri Remote as a game controller]: https://developer.apple.com/documentation/gamecontroller/letting-players-use-their-second-generation-siri-remote-as-a-game-controller
// [Racing wheel device support]: https://developer.apple.com/documentation/gamecontroller/racing-wheel-device-support
// [Understanding game controller backward compatibility]: https://developer.apple.com/documentation/gamecontroller/understanding-game-controller-backward-compatibility
package gamecontroller

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the GameController library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/GameController.framework/GameController",
	"/usr/lib/libGameController.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: GameController: failed to load framework from any known path\n")
	}
}
