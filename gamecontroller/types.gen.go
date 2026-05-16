// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

// C struct types

// GCAcceleration - A three-dimensional acceleration vector.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCAcceleration
type GCAcceleration struct {
	X float64 // The acceleration measurement along the x-axis, in multiples of earth’s gravity.
	Y float64 // The acceleration measurement along the y-axis, in multiples of earth’s gravity.
	Z float64 // The acceleration measurement along the z-axis, in multiples of earth’s gravity.

}

// GCDualSenseAdaptiveTriggerPositionalAmplitudes - The amplitudes for multiple positions on a trigger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/PositionalAmplitudes
type GCDualSenseAdaptiveTriggerPositionalAmplitudes struct {
	Values float32 // The amplitude values for possible trigger positions.

}

// GCDualSenseAdaptiveTriggerPositionalResistiveStrengths - The resistive strengths for multiple positions on a trigger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/PositionalResistiveStrengths
type GCDualSenseAdaptiveTriggerPositionalResistiveStrengths struct {
	Values float32 // The resistive strength values for possible trigger positions.

}

// GCEulerAngles - A structure that specifies the controller’s attitude as a series of rotations around the x, y, and z axes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCEulerAngles
type GCEulerAngles struct {
	Pitch float64 // The pitch of the controller in radians.
	Yaw   float64 // The yaw of the device in radians.
	Roll  float64 // The roll of the controller in radians.

}

// GCExtendedGamepadSnapShotDataV100 - A structure that holds a snapshot of an extended gamepad controller’s input data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepadSnapShotDataV100
type GCExtendedGamepadSnapShotDataV100 struct {
	ButtonA          float32 // The value of the A button.
	ButtonB          float32 // The value of the B button.
	ButtonX          float32 // The value of the X button.
	ButtonY          float32 // The value of the Y button.
	DpadX            float32 // The value of the horizontal axis of the dpad.
	DpadY            float32 // The value of the vertical axis of the dpad.
	LeftShoulder     float32 // The value of the left shoulder button.
	LeftThumbstickX  float32 // The value of the horizontal axis of the left thumbstick.
	LeftThumbstickY  float32 // The value of the vertical axis of the left thumbstick.
	LeftTrigger      float32 // The value of the left trigger.
	RightShoulder    float32 // The value of the right shoulder button.
	RightThumbstickX float32 // The value of the horizontal axis of the right thumbstick.
	RightThumbstickY float32 // The value of the vertical axis of the right thumbstick.
	RightTrigger     float32 // The value of the right trigger.
	Size             uint16  // The size of the recorded structure, in bytes.
	Version          uint16  // A value that indicates the version number of the data structure.

}

// GCExtendedGamepadSnapshotData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepadSnapshotData
type GCExtendedGamepadSnapshotData struct {
	ButtonA                      float32
	ButtonB                      float32
	ButtonX                      float32
	ButtonY                      float32
	DpadX                        float32
	DpadY                        float32
	LeftShoulder                 float32
	LeftThumbstickButton         bool
	LeftThumbstickX              float32
	LeftThumbstickY              float32
	LeftTrigger                  float32
	RightShoulder                float32
	RightThumbstickButton        bool
	RightThumbstickX             float32
	RightThumbstickY             float32
	RightTrigger                 float32
	Size                         uint16
	SupportsClickableThumbsticks bool
	Version                      uint16
}

// GCGamepadSnapShotDataV100 - A structure that holds a snapshot of a gamepad controller’s input data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGamepadSnapShotDataV100
type GCGamepadSnapShotDataV100 struct {
	ButtonA       float32 // The value of the A button.
	ButtonB       float32 // The value of the B button.
	ButtonX       float32 // The value of the X button.
	ButtonY       float32 // The value of the Y button.
	DpadX         float32 // The value of the horizontal axis of the dpad.
	DpadY         float32 // The value of the vertical axis of the dpad.
	LeftShoulder  float32 // The value of the left shoulder button.
	RightShoulder float32 // The value of the right shoulder button.
	Size          uint16  // The size of the recorded structure, in bytes.
	Version       uint16  // A value that indicates the version number of the data structure.

}

// GCMicroGamepadSnapShotDataV100 - A structure that holds a snapshot of a micro gamepad controller’s input data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepadSnapShotDataV100
type GCMicroGamepadSnapShotDataV100 struct {
	ButtonA float32 // The value of the A button.
	ButtonX float32
	DpadX   float32 // The value of the horizontal axis of the dpad.
	DpadY   float32 // The value of the vertical axis of the dpad.
	Size    uint16  // The size of the recorded structure, in bytes.
	Version uint16  // A value that indicates the version number of the data structure.

}

// GCMicroGamepadSnapshotData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepadSnapshotData
type GCMicroGamepadSnapshotData struct {
	ButtonA float32
	ButtonX float32
	DpadX   float32
	DpadY   float32
	Size    uint16
	Version uint16
}

// GCPoint2 - A structure that represents a normalized point in a two-dimensional coordinate system.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCPoint2
type GCPoint2 struct {
	X float32 // The x-coordinate for the point.
	Y float32 // The y-coordinate for the point.

}

// GCQuaternion - A quaternion that represents a controller’s measurement of attitude.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCQuaternion
type GCQuaternion struct {
	X float64 // The value for the x-axis of the quaternion.
	Y float64 // The value for the y-axis of the quaternion.
	Z float64 // The value for the z-axis of the quaternion.
	W float64 // The value for the w-axis of the quaternion.

}

// GCRotationRate - A structure that represents rotation rates around the x, y, and z axes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCRotationRate
type GCRotationRate struct {
	X float64 // The rotation rate around the x-axis in radians per second.
	Y float64 // The rotation rate around the y-axis in radians per second.
	Z float64 // The rotation rate around the z-axis in radians per second.

}
