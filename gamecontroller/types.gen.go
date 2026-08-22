// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"encoding/binary"
	"math"
)

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
	Values [10]float32 // The amplitude values for possible trigger positions.

}

// GCDualSenseAdaptiveTriggerPositionalResistiveStrengths - The resistive strengths for multiple positions on a trigger.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/PositionalResistiveStrengths
type GCDualSenseAdaptiveTriggerPositionalResistiveStrengths struct {
	Values [10]float32 // The resistive strength values for possible trigger positions.

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
	Version          uint16  // A value that indicates the version number of the data structure.
	Size             uint16  // The size of the recorded structure, in bytes.
	DpadX            float32 // The value of the horizontal axis of the dpad.
	DpadY            float32 // The value of the vertical axis of the dpad.
	ButtonA          float32 // The value of the A button.
	ButtonB          float32 // The value of the B button.
	ButtonX          float32 // The value of the X button.
	ButtonY          float32 // The value of the Y button.
	LeftShoulder     float32 // The value of the left shoulder button.
	RightShoulder    float32 // The value of the right shoulder button.
	LeftThumbstickX  float32 // The value of the horizontal axis of the left thumbstick.
	LeftThumbstickY  float32 // The value of the vertical axis of the left thumbstick.
	RightThumbstickX float32 // The value of the horizontal axis of the right thumbstick.
	RightThumbstickY float32 // The value of the vertical axis of the right thumbstick.
	LeftTrigger      float32 // The value of the left trigger.
	RightTrigger     float32 // The value of the right trigger.

}

// GCExtendedGamepadSnapshotData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCExtendedGamepadSnapshotData
type GCExtendedGamepadSnapshotData struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [63]byte
}

// Version returns the Version field from the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) Version() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) SetVersion(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// Size returns the Size field from the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) Size() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetSize updates the Size field in the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) SetSize(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// DpadX returns the DpadX field from the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) DpadX() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetDpadX updates the DpadX field in the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) SetDpadX(v float32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], math.Float32bits(v))
}

// DpadY returns the DpadY field from the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) DpadY() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetDpadY updates the DpadY field in the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) SetDpadY(v float32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], math.Float32bits(v))
}

// ButtonA returns the ButtonA field from the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) ButtonA() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetButtonA updates the ButtonA field in the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) SetButtonA(v float32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], math.Float32bits(v))
}

// ButtonB returns the ButtonB field from the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) ButtonB() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetButtonB updates the ButtonB field in the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) SetButtonB(v float32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], math.Float32bits(v))
}

// ButtonX returns the ButtonX field from the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) ButtonX() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetButtonX updates the ButtonX field in the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) SetButtonX(v float32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], math.Float32bits(v))
}

// ButtonY returns the ButtonY field from the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) ButtonY() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetButtonY updates the ButtonY field in the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) SetButtonY(v float32) {
	binary.NativeEndian.PutUint32(s.storage[24:28], math.Float32bits(v))
}

// LeftShoulder returns the LeftShoulder field from the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) LeftShoulder() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[28:32]))
}

// SetLeftShoulder updates the LeftShoulder field in the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) SetLeftShoulder(v float32) {
	binary.NativeEndian.PutUint32(s.storage[28:32], math.Float32bits(v))
}

// RightShoulder returns the RightShoulder field from the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) RightShoulder() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[32:36]))
}

// SetRightShoulder updates the RightShoulder field in the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) SetRightShoulder(v float32) {
	binary.NativeEndian.PutUint32(s.storage[32:36], math.Float32bits(v))
}

// LeftThumbstickX returns the LeftThumbstickX field from the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) LeftThumbstickX() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[36:40]))
}

// SetLeftThumbstickX updates the LeftThumbstickX field in the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) SetLeftThumbstickX(v float32) {
	binary.NativeEndian.PutUint32(s.storage[36:40], math.Float32bits(v))
}

// LeftThumbstickY returns the LeftThumbstickY field from the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) LeftThumbstickY() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[40:44]))
}

// SetLeftThumbstickY updates the LeftThumbstickY field in the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) SetLeftThumbstickY(v float32) {
	binary.NativeEndian.PutUint32(s.storage[40:44], math.Float32bits(v))
}

// RightThumbstickX returns the RightThumbstickX field from the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) RightThumbstickX() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[44:48]))
}

// SetRightThumbstickX updates the RightThumbstickX field in the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) SetRightThumbstickX(v float32) {
	binary.NativeEndian.PutUint32(s.storage[44:48], math.Float32bits(v))
}

// RightThumbstickY returns the RightThumbstickY field from the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) RightThumbstickY() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[48:52]))
}

// SetRightThumbstickY updates the RightThumbstickY field in the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) SetRightThumbstickY(v float32) {
	binary.NativeEndian.PutUint32(s.storage[48:52], math.Float32bits(v))
}

// LeftTrigger returns the LeftTrigger field from the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) LeftTrigger() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[52:56]))
}

// SetLeftTrigger updates the LeftTrigger field in the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) SetLeftTrigger(v float32) {
	binary.NativeEndian.PutUint32(s.storage[52:56], math.Float32bits(v))
}

// RightTrigger returns the RightTrigger field from the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) RightTrigger() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[56:60]))
}

// SetRightTrigger updates the RightTrigger field in the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) SetRightTrigger(v float32) {
	binary.NativeEndian.PutUint32(s.storage[56:60], math.Float32bits(v))
}

// SupportsClickableThumbsticks returns the SupportsClickableThumbsticks field from the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) SupportsClickableThumbsticks() bool {
	return s.storage[60] != 0
}

// SetSupportsClickableThumbsticks updates the SupportsClickableThumbsticks field in the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) SetSupportsClickableThumbsticks(v bool) {
	if v {
		s.storage[60] = 1
	} else {
		s.storage[60] = 0
	}
}

// LeftThumbstickButton returns the LeftThumbstickButton field from the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) LeftThumbstickButton() bool {
	return s.storage[61] != 0
}

// SetLeftThumbstickButton updates the LeftThumbstickButton field in the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) SetLeftThumbstickButton(v bool) {
	if v {
		s.storage[61] = 1
	} else {
		s.storage[61] = 0
	}
}

// RightThumbstickButton returns the RightThumbstickButton field from the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) RightThumbstickButton() bool {
	return s.storage[62] != 0
}

// SetRightThumbstickButton updates the RightThumbstickButton field in the record's packed storage.
func (s *GCExtendedGamepadSnapshotData) SetRightThumbstickButton(v bool) {
	if v {
		s.storage[62] = 1
	} else {
		s.storage[62] = 0
	}
}

// GCGamepadSnapShotDataV100 - A structure that holds a snapshot of a gamepad controller’s input data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCGamepadSnapShotDataV100
type GCGamepadSnapShotDataV100 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [36]byte
}

// Version returns the Version field from the record's packed storage.
func (s *GCGamepadSnapShotDataV100) Version() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *GCGamepadSnapShotDataV100) SetVersion(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// Size returns the Size field from the record's packed storage.
func (s *GCGamepadSnapShotDataV100) Size() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetSize updates the Size field in the record's packed storage.
func (s *GCGamepadSnapShotDataV100) SetSize(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// DpadX returns the DpadX field from the record's packed storage.
func (s *GCGamepadSnapShotDataV100) DpadX() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetDpadX updates the DpadX field in the record's packed storage.
func (s *GCGamepadSnapShotDataV100) SetDpadX(v float32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], math.Float32bits(v))
}

// DpadY returns the DpadY field from the record's packed storage.
func (s *GCGamepadSnapShotDataV100) DpadY() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetDpadY updates the DpadY field in the record's packed storage.
func (s *GCGamepadSnapShotDataV100) SetDpadY(v float32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], math.Float32bits(v))
}

// ButtonA returns the ButtonA field from the record's packed storage.
func (s *GCGamepadSnapShotDataV100) ButtonA() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetButtonA updates the ButtonA field in the record's packed storage.
func (s *GCGamepadSnapShotDataV100) SetButtonA(v float32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], math.Float32bits(v))
}

// ButtonB returns the ButtonB field from the record's packed storage.
func (s *GCGamepadSnapShotDataV100) ButtonB() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetButtonB updates the ButtonB field in the record's packed storage.
func (s *GCGamepadSnapShotDataV100) SetButtonB(v float32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], math.Float32bits(v))
}

// ButtonX returns the ButtonX field from the record's packed storage.
func (s *GCGamepadSnapShotDataV100) ButtonX() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[20:24]))
}

// SetButtonX updates the ButtonX field in the record's packed storage.
func (s *GCGamepadSnapShotDataV100) SetButtonX(v float32) {
	binary.NativeEndian.PutUint32(s.storage[20:24], math.Float32bits(v))
}

// ButtonY returns the ButtonY field from the record's packed storage.
func (s *GCGamepadSnapShotDataV100) ButtonY() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[24:28]))
}

// SetButtonY updates the ButtonY field in the record's packed storage.
func (s *GCGamepadSnapShotDataV100) SetButtonY(v float32) {
	binary.NativeEndian.PutUint32(s.storage[24:28], math.Float32bits(v))
}

// LeftShoulder returns the LeftShoulder field from the record's packed storage.
func (s *GCGamepadSnapShotDataV100) LeftShoulder() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[28:32]))
}

// SetLeftShoulder updates the LeftShoulder field in the record's packed storage.
func (s *GCGamepadSnapShotDataV100) SetLeftShoulder(v float32) {
	binary.NativeEndian.PutUint32(s.storage[28:32], math.Float32bits(v))
}

// RightShoulder returns the RightShoulder field from the record's packed storage.
func (s *GCGamepadSnapShotDataV100) RightShoulder() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[32:36]))
}

// SetRightShoulder updates the RightShoulder field in the record's packed storage.
func (s *GCGamepadSnapShotDataV100) SetRightShoulder(v float32) {
	binary.NativeEndian.PutUint32(s.storage[32:36], math.Float32bits(v))
}

// GCMicroGamepadSnapShotDataV100 - A structure that holds a snapshot of a micro gamepad controller’s input data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepadSnapShotDataV100
type GCMicroGamepadSnapShotDataV100 struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [20]byte
}

// Version returns the Version field from the record's packed storage.
func (s *GCMicroGamepadSnapShotDataV100) Version() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *GCMicroGamepadSnapShotDataV100) SetVersion(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// Size returns the Size field from the record's packed storage.
func (s *GCMicroGamepadSnapShotDataV100) Size() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetSize updates the Size field in the record's packed storage.
func (s *GCMicroGamepadSnapShotDataV100) SetSize(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// DpadX returns the DpadX field from the record's packed storage.
func (s *GCMicroGamepadSnapShotDataV100) DpadX() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetDpadX updates the DpadX field in the record's packed storage.
func (s *GCMicroGamepadSnapShotDataV100) SetDpadX(v float32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], math.Float32bits(v))
}

// DpadY returns the DpadY field from the record's packed storage.
func (s *GCMicroGamepadSnapShotDataV100) DpadY() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetDpadY updates the DpadY field in the record's packed storage.
func (s *GCMicroGamepadSnapShotDataV100) SetDpadY(v float32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], math.Float32bits(v))
}

// ButtonA returns the ButtonA field from the record's packed storage.
func (s *GCMicroGamepadSnapShotDataV100) ButtonA() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetButtonA updates the ButtonA field in the record's packed storage.
func (s *GCMicroGamepadSnapShotDataV100) SetButtonA(v float32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], math.Float32bits(v))
}

// ButtonX returns the ButtonX field from the record's packed storage.
func (s *GCMicroGamepadSnapShotDataV100) ButtonX() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetButtonX updates the ButtonX field in the record's packed storage.
func (s *GCMicroGamepadSnapShotDataV100) SetButtonX(v float32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], math.Float32bits(v))
}

// GCMicroGamepadSnapshotData
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameController/GCMicroGamepadSnapshotData
type GCMicroGamepadSnapshotData struct {
	// storage holds the record exactly as the C compiler lays it out. Its
	// members are reached through the accessors below, which slice it at
	// their measured offsets.
	//
	// The members cannot be ordinary Go fields: Go would place at least one
	// of them somewhere other than where C measured it, and every member
	// after that one would move with it.
	storage [20]byte
}

// Version returns the Version field from the record's packed storage.
func (s *GCMicroGamepadSnapshotData) Version() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[0:2]))
}

// SetVersion updates the Version field in the record's packed storage.
func (s *GCMicroGamepadSnapshotData) SetVersion(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[0:2], uint16(v))
}

// Size returns the Size field from the record's packed storage.
func (s *GCMicroGamepadSnapshotData) Size() uint16 {
	return uint16(binary.NativeEndian.Uint16(s.storage[2:4]))
}

// SetSize updates the Size field in the record's packed storage.
func (s *GCMicroGamepadSnapshotData) SetSize(v uint16) {
	binary.NativeEndian.PutUint16(s.storage[2:4], uint16(v))
}

// DpadX returns the DpadX field from the record's packed storage.
func (s *GCMicroGamepadSnapshotData) DpadX() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[4:8]))
}

// SetDpadX updates the DpadX field in the record's packed storage.
func (s *GCMicroGamepadSnapshotData) SetDpadX(v float32) {
	binary.NativeEndian.PutUint32(s.storage[4:8], math.Float32bits(v))
}

// DpadY returns the DpadY field from the record's packed storage.
func (s *GCMicroGamepadSnapshotData) DpadY() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[8:12]))
}

// SetDpadY updates the DpadY field in the record's packed storage.
func (s *GCMicroGamepadSnapshotData) SetDpadY(v float32) {
	binary.NativeEndian.PutUint32(s.storage[8:12], math.Float32bits(v))
}

// ButtonA returns the ButtonA field from the record's packed storage.
func (s *GCMicroGamepadSnapshotData) ButtonA() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[12:16]))
}

// SetButtonA updates the ButtonA field in the record's packed storage.
func (s *GCMicroGamepadSnapshotData) SetButtonA(v float32) {
	binary.NativeEndian.PutUint32(s.storage[12:16], math.Float32bits(v))
}

// ButtonX returns the ButtonX field from the record's packed storage.
func (s *GCMicroGamepadSnapshotData) ButtonX() float32 {
	return math.Float32frombits(binary.NativeEndian.Uint32(s.storage[16:20]))
}

// SetButtonX updates the ButtonX field in the record's packed storage.
func (s *GCMicroGamepadSnapshotData) SetButtonX(v float32) {
	binary.NativeEndian.PutUint32(s.storage[16:20], math.Float32bits(v))
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
