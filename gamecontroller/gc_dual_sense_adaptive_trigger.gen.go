// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"sync"

	"github.com/tmc/apple/objc"
)

// The class instance for the [GCDualSenseAdaptiveTrigger] class.
var (
	_GCDualSenseAdaptiveTriggerClass     GCDualSenseAdaptiveTriggerClass
	_GCDualSenseAdaptiveTriggerClassOnce sync.Once
)

func getGCDualSenseAdaptiveTriggerClass() GCDualSenseAdaptiveTriggerClass {
	_GCDualSenseAdaptiveTriggerClassOnce.Do(func() {
		_GCDualSenseAdaptiveTriggerClass = GCDualSenseAdaptiveTriggerClass{class: objc.GetClass("GCDualSenseAdaptiveTrigger")}
	})
	return _GCDualSenseAdaptiveTriggerClass
}

// GetGCDualSenseAdaptiveTriggerClass returns the class object for GCDualSenseAdaptiveTrigger.
func GetGCDualSenseAdaptiveTriggerClass() GCDualSenseAdaptiveTriggerClass {
	return getGCDualSenseAdaptiveTriggerClass()
}

type GCDualSenseAdaptiveTriggerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GCDualSenseAdaptiveTriggerClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GCDualSenseAdaptiveTriggerClass) Alloc() GCDualSenseAdaptiveTrigger {
	rv := objc.Send[GCDualSenseAdaptiveTrigger](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// A class that encapsulates the features of a DualSense adaptive trigger.
//
// # Overview
//
// A [GCDualSenseAdaptiveTrigger] object allows you to specify a dynamic
// resistance force that the DualSense controller applies when the user pulls
// the trigger. For example, set the resistance to give the user the feeling
// of pulling back on a bow string, firing a weapon, or pulling a lever.
//
// # Getting the mode
//
//   - [GCDualSenseAdaptiveTrigger.Mode]: The current configuration of the adaptive trigger.
//   - [GCDualSenseAdaptiveTrigger.SetModeOff]: Sets the mode to off and stops any trigger effect.
//
// # Configuring the trigger
//
//   - [GCDualSenseAdaptiveTrigger.SetModeFeedbackWithStartPositionResistiveStrength]: Sets the mode to provide feedback when the user depresses the trigger at the start position or at a greater value.
//   - [GCDualSenseAdaptiveTrigger.SetModeFeedbackWithResistiveStrengths]: Sets the mode to provide feedback with the specified strengths for each possible trigger position.
//   - [GCDualSenseAdaptiveTrigger.SetModeWeaponWithStartPositionEndPositionResistiveStrength]: Sets the mode to provide feedback when the user depresses the trigger between the start and the end positions.
//   - [GCDualSenseAdaptiveTrigger.SetModeVibrationWithStartPositionAmplitudeFrequency]: Sets the mode to vibrate when the user depresses the trigger at the start position or at a greater value.
//   - [GCDualSenseAdaptiveTrigger.SetModeVibrationWithAmplitudesFrequency]: Sets the mode to vibrate with the specified amplitudes for each possible trigger position.
//   - [GCDualSenseAdaptiveTrigger.SetModeSlopeFeedbackWithStartPositionEndPositionStartStrengthEndStrength]: Sets the mode to provide feedback when the user tilts the trigger between the start and the end positions.
//
// # Getting the arm position
//
//   - [GCDualSenseAdaptiveTrigger.ArmPosition]: The position of the trigger’s arm.
//
// # Checking the status
//
//   - [GCDualSenseAdaptiveTrigger.Status]: The current status of the adaptive trigger and whether it’s applying effects.
//
// See: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger
type GCDualSenseAdaptiveTrigger struct {
	GCControllerButtonInput
}

// GCDualSenseAdaptiveTriggerFromID constructs a [GCDualSenseAdaptiveTrigger] from an objc.ID.
//
// A class that encapsulates the features of a DualSense adaptive trigger.
func GCDualSenseAdaptiveTriggerFromID(id objc.ID) GCDualSenseAdaptiveTrigger {
	return GCDualSenseAdaptiveTrigger{GCControllerButtonInput: GCControllerButtonInputFromID(id)}
}

// NOTE: GCDualSenseAdaptiveTrigger adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [GCDualSenseAdaptiveTrigger] class.
//
// # Getting the mode
//
//   - [IGCDualSenseAdaptiveTrigger.Mode]: The current configuration of the adaptive trigger.
//   - [IGCDualSenseAdaptiveTrigger.SetModeOff]: Sets the mode to off and stops any trigger effect.
//
// # Configuring the trigger
//
//   - [IGCDualSenseAdaptiveTrigger.SetModeFeedbackWithStartPositionResistiveStrength]: Sets the mode to provide feedback when the user depresses the trigger at the start position or at a greater value.
//   - [IGCDualSenseAdaptiveTrigger.SetModeFeedbackWithResistiveStrengths]: Sets the mode to provide feedback with the specified strengths for each possible trigger position.
//   - [IGCDualSenseAdaptiveTrigger.SetModeWeaponWithStartPositionEndPositionResistiveStrength]: Sets the mode to provide feedback when the user depresses the trigger between the start and the end positions.
//   - [IGCDualSenseAdaptiveTrigger.SetModeVibrationWithStartPositionAmplitudeFrequency]: Sets the mode to vibrate when the user depresses the trigger at the start position or at a greater value.
//   - [IGCDualSenseAdaptiveTrigger.SetModeVibrationWithAmplitudesFrequency]: Sets the mode to vibrate with the specified amplitudes for each possible trigger position.
//   - [IGCDualSenseAdaptiveTrigger.SetModeSlopeFeedbackWithStartPositionEndPositionStartStrengthEndStrength]: Sets the mode to provide feedback when the user tilts the trigger between the start and the end positions.
//
// # Getting the arm position
//
//   - [IGCDualSenseAdaptiveTrigger.ArmPosition]: The position of the trigger’s arm.
//
// # Checking the status
//
//   - [IGCDualSenseAdaptiveTrigger.Status]: The current status of the adaptive trigger and whether it’s applying effects.
//
// See: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger
type IGCDualSenseAdaptiveTrigger interface {
	IGCControllerButtonInput

	// Topic: Getting the mode

	// The current configuration of the adaptive trigger.
	Mode() GCDualSenseAdaptiveTriggerMode
	// Sets the mode to off and stops any trigger effect.
	SetModeOff()

	// Topic: Configuring the trigger

	// Sets the mode to provide feedback when the user depresses the trigger at the start position or at a greater value.
	SetModeFeedbackWithStartPositionResistiveStrength(startPosition float32, resistiveStrength float32)
	// Sets the mode to provide feedback with the specified strengths for each possible trigger position.
	SetModeFeedbackWithResistiveStrengths(positionalResistiveStrengths GCDualSenseAdaptiveTriggerPositionalResistiveStrengths)
	// Sets the mode to provide feedback when the user depresses the trigger between the start and the end positions.
	SetModeWeaponWithStartPositionEndPositionResistiveStrength(startPosition float32, endPosition float32, resistiveStrength float32)
	// Sets the mode to vibrate when the user depresses the trigger at the start position or at a greater value.
	SetModeVibrationWithStartPositionAmplitudeFrequency(startPosition float32, amplitude float32, frequency float32)
	// Sets the mode to vibrate with the specified amplitudes for each possible trigger position.
	SetModeVibrationWithAmplitudesFrequency(positionalAmplitudes GCDualSenseAdaptiveTriggerPositionalAmplitudes, frequency float32)
	// Sets the mode to provide feedback when the user tilts the trigger between the start and the end positions.
	SetModeSlopeFeedbackWithStartPositionEndPositionStartStrengthEndStrength(startPosition float32, endPosition float32, startStrength float32, endStrength float32)

	// Topic: Getting the arm position

	// The position of the trigger’s arm.
	ArmPosition() float32

	// Topic: Checking the status

	// The current status of the adaptive trigger and whether it’s applying effects.
	Status() GCDualSenseAdaptiveTriggerStatus
}

// Init initializes the instance.
func (g GCDualSenseAdaptiveTrigger) Init() GCDualSenseAdaptiveTrigger {
	rv := objc.Send[GCDualSenseAdaptiveTrigger](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GCDualSenseAdaptiveTrigger) Autorelease() GCDualSenseAdaptiveTrigger {
	rv := objc.Send[GCDualSenseAdaptiveTrigger](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGCDualSenseAdaptiveTrigger creates a new GCDualSenseAdaptiveTrigger instance.
func NewGCDualSenseAdaptiveTrigger() GCDualSenseAdaptiveTrigger {
	class := getGCDualSenseAdaptiveTriggerClass()
	rv := objc.Send[GCDualSenseAdaptiveTrigger](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Sets the mode to off and stops any trigger effect.
//
// See: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/setModeOff()
func (g GCDualSenseAdaptiveTrigger) SetModeOff() {
	objc.Send[objc.ID](g.ID, objc.Sel("setModeOff"))
}

// Sets the mode to provide feedback when the user depresses the trigger at
// the start position or at a greater value.
//
// startPosition: The effect’s start position. A value between `0` and `1` , where `0` is
// the minimum and `1` is the maximum trigger depression.
//
// resistiveStrength: The strength of the feedback. A value between `0` and `1`, where `0` is the
// minimum and `1` is the maximum strength.
//
// See: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/setModeFeedbackWithStartPosition(_:resistiveStrength:)
func (g GCDualSenseAdaptiveTrigger) SetModeFeedbackWithStartPositionResistiveStrength(startPosition float32, resistiveStrength float32) {
	objc.Send[objc.ID](g.ID, objc.Sel("setModeFeedbackWithStartPosition:resistiveStrength:"), startPosition, resistiveStrength)
}

// Sets the mode to provide feedback with the specified strengths for each
// possible trigger position.
//
// positionalResistiveStrengths: The resistance values for each possible trigger position.
//
// See: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/setModeFeedback(resistiveStrengths:)
func (g GCDualSenseAdaptiveTrigger) SetModeFeedbackWithResistiveStrengths(positionalResistiveStrengths GCDualSenseAdaptiveTriggerPositionalResistiveStrengths) {
	objc.Send[objc.ID](g.ID, objc.Sel("setModeFeedbackWithResistiveStrengths:"), positionalResistiveStrengths)
}

// Sets the mode to provide feedback when the user depresses the trigger
// between the start and the end positions.
//
// startPosition: The effect’s start position. A value between `0` and `1` , where `0` is
// the minimum and `1` is the maximum trigger depression.
//
// endPosition: The effect’s end position. A value between `0` and `1` , where `0` is the
// minimum and `1` is the maximum trigger depression. This value must be
// greater than `startPosition`.
//
// resistiveStrength: The strength of the effect. A value between `0` and `1`, where `0` is the
// minimum or off value, and `1` is the maximum strength.
//
// # Discussion
//
// When the user depresses the trigger beyond the value of the end position,
// it stops providing feedback, giving the user a sense of release, similar to
// pulling a weapon trigger.
//
// See: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/setModeWeaponWithStartPosition(_:endPosition:resistiveStrength:)
func (g GCDualSenseAdaptiveTrigger) SetModeWeaponWithStartPositionEndPositionResistiveStrength(startPosition float32, endPosition float32, resistiveStrength float32) {
	objc.Send[objc.ID](g.ID, objc.Sel("setModeWeaponWithStartPosition:endPosition:resistiveStrength:"), startPosition, endPosition, resistiveStrength)
}

// Sets the mode to vibrate when the user depresses the trigger at the start
// position or at a greater value.
//
// startPosition: The effect’s start position. A value between `0` and `1` , where `0` is
// the minimum and `1` is the maximum trigger depression.
//
// amplitude: The amplitude of the vibration effect. A value between `0` and `1`, where
// `0` is the minimum and `1` is the maximum amplitude.
//
// frequency: The frequency of the vibration effect, which is a value between `0` and
// `1`, where `0` is the minimum and `1` is the maximum frequency.
//
// See: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/setModeVibrationWithStartPosition(_:amplitude:frequency:)
func (g GCDualSenseAdaptiveTrigger) SetModeVibrationWithStartPositionAmplitudeFrequency(startPosition float32, amplitude float32, frequency float32) {
	objc.Send[objc.ID](g.ID, objc.Sel("setModeVibrationWithStartPosition:amplitude:frequency:"), startPosition, amplitude, frequency)
}

// Sets the mode to vibrate with the specified amplitudes for each possible
// trigger position.
//
// positionalAmplitudes: The amplitudes for each possible trigger position.
//
// frequency: The frequency of the vibration effect, which is a value between `0` and
// `1`, where `0` is the minimum and `1` is the maximum frequency.
//
// See: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/setModeVibration(amplitudes:frequency:)
func (g GCDualSenseAdaptiveTrigger) SetModeVibrationWithAmplitudesFrequency(positionalAmplitudes GCDualSenseAdaptiveTriggerPositionalAmplitudes, frequency float32) {
	objc.Send[objc.ID](g.ID, objc.Sel("setModeVibrationWithAmplitudes:frequency:"), positionalAmplitudes, frequency)
}

// Sets the mode to provide feedback when the user tilts the trigger between
// the start and the end positions.
//
// startPosition: The effect’s start position, which is a value between `0` and `1` , where
// `0` is the minimum and `1` is the maximum trigger tilt.
//
// endPosition: The effect’s end position, which is a value between `0` and `1` , where
// `0` is the minimum and `1` is the maximum trigger tilt. This value must be
// greater than `startPosition`.
//
// startStrength: The effect’s start strength, which is a value between `0` and `1`, where
// `0` is the minimum or off value, and `1` is the maximum strength.
//
// endStrength: The effect’s end strength, which is a value between `0` and `1`, where
// `0` is the minimum or off value, and `1` is the maximum strength.
//
// See: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/setModeSlopeFeedback(startPosition:endPosition:startStrength:endStrength:)
func (g GCDualSenseAdaptiveTrigger) SetModeSlopeFeedbackWithStartPositionEndPositionStartStrengthEndStrength(startPosition float32, endPosition float32, startStrength float32, endStrength float32) {
	objc.Send[objc.ID](g.ID, objc.Sel("setModeSlopeFeedbackWithStartPosition:endPosition:startStrength:endStrength:"), startPosition, endPosition, startStrength, endStrength)
}

// The current configuration of the adaptive trigger.
//
// # Discussion
//
// There may be a delay updating this property value after you set the mode
// using one of the set mode methods because setting the mode requires a
// response from the controller.
//
// See: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/mode-swift.property
func (g GCDualSenseAdaptiveTrigger) Mode() GCDualSenseAdaptiveTriggerMode {
	rv := objc.Send[GCDualSenseAdaptiveTriggerMode](g.ID, objc.Sel("mode"))
	return GCDualSenseAdaptiveTriggerMode(rv)
}

// The position of the trigger’s arm.
//
// # Discussion
//
// This property represents the value of the stepped mechanical arm inside the
// trigger and isn’t the same as the trigger’s inherited `value` property.
// This property ranges between `0` and `1`, where `0` represents the minimum
// and `1` represents the maximum position.
//
// See: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/armPosition
func (g GCDualSenseAdaptiveTrigger) ArmPosition() float32 {
	rv := objc.Send[float32](g.ID, objc.Sel("armPosition"))
	return rv
}

// The current status of the adaptive trigger and whether it’s applying
// effects.
//
// See: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/status-swift.property
func (g GCDualSenseAdaptiveTrigger) Status() GCDualSenseAdaptiveTriggerStatus {
	rv := objc.Send[GCDualSenseAdaptiveTriggerStatus](g.ID, objc.Sel("status"))
	return GCDualSenseAdaptiveTriggerStatus(rv)
}
