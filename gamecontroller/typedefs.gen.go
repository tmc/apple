// Code generated from Apple documentation. DO NOT EDIT.

package gamecontroller

// GCControllerAxisValueChangedHandler is the signature for the block that executes when the user changes the axis value.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerAxisValueChangedHandler
type GCControllerAxisValueChangedHandler = func(axis GCControllerAxisInput, value float32)

// GCControllerButtonTouchedChangedHandler is the signature for the block that executes when the user touches the button if the controller supports that feature.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerButtonTouchedChangedHandler
type GCControllerButtonTouchedChangedHandler = func(button GCControllerButtonInput, value float32, pressed bool, touched bool)

// GCControllerButtonValueChangedHandler is the signature for the block that executes when a button’s state changes.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerButtonValueChangedHandler
type GCControllerButtonValueChangedHandler = func(button GCControllerButtonInput, value float32, pressed bool)

// GCControllerDirectionPadValueChangedHandler is the signature for the block that executes when either axis changes values.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerDirectionPadValueChangedHandler
type GCControllerDirectionPadValueChangedHandler = func(dpad GCControllerDirectionPad, xValue float32, yValue float32)

// GCControllerTouchpadHandler is the signature for the block that executes when the user interacts with the touchpad.
//
// See: https://developer.apple.com/documentation/GameController/GCControllerTouchpadHandler
type GCControllerTouchpadHandler = func(touchpad GCControllerTouchpad, xValue float32, yValue float32, buttonValue float32, buttonPressed bool)

// GCExtendedGamepadValueChangedHandler is the signature for the block that the profile calls when an element’s value changes.
//
// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepadValueChangedHandler
type GCExtendedGamepadValueChangedHandler = func(gamepad GCExtendedGamepad, element GCControllerElement)

// GCGamepadValueChangedHandler is signature for the block executed if any element in the gamepad profile changes value.
//
// See: https://developer.apple.com/documentation/GameController/GCGamepadValueChangedHandler
type GCGamepadValueChangedHandler = func(*uintptr, GCControllerElement)

// GCHapticsLocality is the location of one or more haptics actuators on a game controller.
//
// See: https://developer.apple.com/documentation/GameController/GCHapticsLocality
type GCHapticsLocality = string

// GCInputAxisName is the Objective-C type for an input axis name.
//
// See: https://developer.apple.com/documentation/GameController/GCInputAxisName
type GCInputAxisName = string

// GCInputButtonName is the Objective-C type for an input button name.
//
// See: https://developer.apple.com/documentation/GameController/GCInputButtonName
type GCInputButtonName = string

// GCInputDirectionPadName is the Objective-C type for the name of a directional pad.
//
// See: https://developer.apple.com/documentation/GameController/GCInputDirectionPadName
type GCInputDirectionPadName = string

// GCInputElementName is the Objective-C type for an input element name.
//
// See: https://developer.apple.com/documentation/GameController/GCInputElementName
type GCInputElementName = string

// GCInputSwitchName is the Objective-C type for an input switch name.
//
// See: https://developer.apple.com/documentation/GameController/GCInputSwitchName
type GCInputSwitchName = string

// GCKeyCode is the key codes for keys on a keyboard.
//
// See: https://developer.apple.com/documentation/GameController/GCKeyCode
type GCKeyCode = int

// GCKeyboardValueChangedHandler is the signature for the block that the keyboard input profile calls when a key value changes.
//
// See: https://developer.apple.com/documentation/GameController/GCKeyboardValueChangedHandler
type GCKeyboardValueChangedHandler = func(keyboard GCKeyboardInput, key GCControllerButtonInput, keyCode int, pressed bool)

// GCMicroGamepadValueChangedHandler is signature for the block that this profile calls when an element’s value changes.
//
// See: https://developer.apple.com/documentation/GameController/GCMicroGamepadValueChangedHandler
type GCMicroGamepadValueChangedHandler = func(gamepad GCMicroGamepad, element GCControllerElement)

// GCMotionValueChangedHandler is the signature for the block that the profile calls when an element’s value changes.
//
// See: https://developer.apple.com/documentation/GameController/GCMotionValueChangedHandler
type GCMotionValueChangedHandler = func(motion GCMotion)

// GCMouseMoved is the signature for the block that the mouse input profile calls when the mouse moves.
//
// See: https://developer.apple.com/documentation/GameController/GCMouseMoved
type GCMouseMoved = func(mouse GCMouseInput, deltaX float32, deltaY float32)
