// Code generated from Apple documentation for GameController. DO NOT EDIT.

package gamecontroller

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/GameController/GCControllerPlayerIndex
type GCControllerPlayerIndex int

const (
	// GCControllerPlayerIndex1: Player one is using the controller.
	GCControllerPlayerIndex1 GCControllerPlayerIndex = 0
	// GCControllerPlayerIndex2: Player two is using the controller.
	GCControllerPlayerIndex2 GCControllerPlayerIndex = 1
	// GCControllerPlayerIndex3: Player three is using the controller.
	GCControllerPlayerIndex3 GCControllerPlayerIndex = 2
	// GCControllerPlayerIndex4: Player four is using the controller.
	GCControllerPlayerIndex4 GCControllerPlayerIndex = 3
	// GCControllerPlayerIndexUnset: The default index for a player on a controller.
	GCControllerPlayerIndexUnset GCControllerPlayerIndex = -1
)

func (e GCControllerPlayerIndex) String() string {
	switch e {
	case GCControllerPlayerIndex1:
		return "GCControllerPlayerIndex1"
	case GCControllerPlayerIndex2:
		return "GCControllerPlayerIndex2"
	case GCControllerPlayerIndex3:
		return "GCControllerPlayerIndex3"
	case GCControllerPlayerIndex4:
		return "GCControllerPlayerIndex4"
	case GCControllerPlayerIndexUnset:
		return "GCControllerPlayerIndexUnset"
	default:
		return fmt.Sprintf("GCControllerPlayerIndex(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GameController/GCDeviceBattery/State
type GCDeviceBatteryState int

const (
	// GCDeviceBatteryStateCharging: The device’s battery has power and is charging, but isn’t fully charged.
	GCDeviceBatteryStateCharging GCDeviceBatteryState = 1
	// GCDeviceBatteryStateDischarging: The device’s battery is discharging.
	GCDeviceBatteryStateDischarging GCDeviceBatteryState = 0
	// GCDeviceBatteryStateFull: The device’s battery has power and is fully charged.
	GCDeviceBatteryStateFull GCDeviceBatteryState = 2
	// GCDeviceBatteryStateUnknown: The state of the device’s battery is unknown.
	GCDeviceBatteryStateUnknown GCDeviceBatteryState = -1
)

func (e GCDeviceBatteryState) String() string {
	switch e {
	case GCDeviceBatteryStateCharging:
		return "GCDeviceBatteryStateCharging"
	case GCDeviceBatteryStateDischarging:
		return "GCDeviceBatteryStateDischarging"
	case GCDeviceBatteryStateFull:
		return "GCDeviceBatteryStateFull"
	case GCDeviceBatteryStateUnknown:
		return "GCDeviceBatteryStateUnknown"
	default:
		return fmt.Sprintf("GCDeviceBatteryState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GameController/GCDevicePhysicalInputElementChange
type GCDevicePhysicalInputElementChange int

const (
	// GCDevicePhysicalInputElementChanged: There’s a change to the input value.
	GCDevicePhysicalInputElementChanged GCDevicePhysicalInputElementChange = 1
	// GCDevicePhysicalInputElementNoChange: There’s no change to the input value.
	GCDevicePhysicalInputElementNoChange GCDevicePhysicalInputElementChange = 0
	// GCDevicePhysicalInputElementUnknownChange: It’s unknown whether there’s a change to the input value.
	GCDevicePhysicalInputElementUnknownChange GCDevicePhysicalInputElementChange = -1
)

func (e GCDevicePhysicalInputElementChange) String() string {
	switch e {
	case GCDevicePhysicalInputElementChanged:
		return "GCDevicePhysicalInputElementChanged"
	case GCDevicePhysicalInputElementNoChange:
		return "GCDevicePhysicalInputElementNoChange"
	case GCDevicePhysicalInputElementUnknownChange:
		return "GCDevicePhysicalInputElementUnknownChange"
	default:
		return fmt.Sprintf("GCDevicePhysicalInputElementChange(%d)", e)
	}
}

type GCDualSenseAdaptiveTriggerDiscretePositionCountConstants uint

const (
	// GCDualSenseAdaptiveTriggerDiscretePositionCount: The number of discrete control positions that the DualSense adaptive triggers support.
	GCDualSenseAdaptiveTriggerDiscretePositionCount GCDualSenseAdaptiveTriggerDiscretePositionCountConstants = 10
)

func (e GCDualSenseAdaptiveTriggerDiscretePositionCountConstants) String() string {
	switch e {
	case GCDualSenseAdaptiveTriggerDiscretePositionCount:
		return "GCDualSenseAdaptiveTriggerDiscretePositionCount"
	default:
		return fmt.Sprintf("GCDualSenseAdaptiveTriggerDiscretePositionCountConstants(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/Mode-swift.enum
type GCDualSenseAdaptiveTriggerMode int

const (
	// GCDualSenseAdaptiveTriggerModeFeedback: Provides feedback when the user depresses the trigger equal to, or greater than, the start position.
	GCDualSenseAdaptiveTriggerModeFeedback GCDualSenseAdaptiveTriggerMode = 1
	// GCDualSenseAdaptiveTriggerModeOff: Provides no adaptive trigger effects.
	GCDualSenseAdaptiveTriggerModeOff GCDualSenseAdaptiveTriggerMode = 0
	// GCDualSenseAdaptiveTriggerModeSlopeFeedback: Provides feedback when the user tilts the trigger between the start and the end positions.
	GCDualSenseAdaptiveTriggerModeSlopeFeedback GCDualSenseAdaptiveTriggerMode = 4
	// GCDualSenseAdaptiveTriggerModeVibration: Vibrates when the user depresses the trigger equal to, or greater than, the start position.
	GCDualSenseAdaptiveTriggerModeVibration GCDualSenseAdaptiveTriggerMode = 3
	// GCDualSenseAdaptiveTriggerModeWeapon: Provides feedback when the user depresses the trigger between the start and the end positions.
	GCDualSenseAdaptiveTriggerModeWeapon GCDualSenseAdaptiveTriggerMode = 2
)

func (e GCDualSenseAdaptiveTriggerMode) String() string {
	switch e {
	case GCDualSenseAdaptiveTriggerModeFeedback:
		return "GCDualSenseAdaptiveTriggerModeFeedback"
	case GCDualSenseAdaptiveTriggerModeOff:
		return "GCDualSenseAdaptiveTriggerModeOff"
	case GCDualSenseAdaptiveTriggerModeSlopeFeedback:
		return "GCDualSenseAdaptiveTriggerModeSlopeFeedback"
	case GCDualSenseAdaptiveTriggerModeVibration:
		return "GCDualSenseAdaptiveTriggerModeVibration"
	case GCDualSenseAdaptiveTriggerModeWeapon:
		return "GCDualSenseAdaptiveTriggerModeWeapon"
	default:
		return fmt.Sprintf("GCDualSenseAdaptiveTriggerMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GameController/GCDualSenseAdaptiveTrigger/Status-swift.enum
type GCDualSenseAdaptiveTriggerStatus int

const (
	// GCDualSenseAdaptiveTriggerStatusFeedbackLoadApplied: The trigger is in feedback mode and is applying the resistive load.
	GCDualSenseAdaptiveTriggerStatusFeedbackLoadApplied GCDualSenseAdaptiveTriggerStatus = 1
	// GCDualSenseAdaptiveTriggerStatusFeedbackNoLoad: The trigger is in feedback mode, but isn’t applying the resistive load.
	GCDualSenseAdaptiveTriggerStatusFeedbackNoLoad GCDualSenseAdaptiveTriggerStatus = 0
	// GCDualSenseAdaptiveTriggerStatusSlopeFeedbackApplyingLoad: The trigger is in slope mode, and is applying the resistive load.
	GCDualSenseAdaptiveTriggerStatusSlopeFeedbackApplyingLoad GCDualSenseAdaptiveTriggerStatus = 8
	// GCDualSenseAdaptiveTriggerStatusSlopeFeedbackFinished: The trigger is in slope mode, and stopped applying the resistive load.
	GCDualSenseAdaptiveTriggerStatusSlopeFeedbackFinished GCDualSenseAdaptiveTriggerStatus = 9
	// GCDualSenseAdaptiveTriggerStatusSlopeFeedbackReady: The trigger is in slope mode, but isn’t applying the resistive load.
	GCDualSenseAdaptiveTriggerStatusSlopeFeedbackReady GCDualSenseAdaptiveTriggerStatus = 7
	// GCDualSenseAdaptiveTriggerStatusUnknown: The trigger status is unknown.
	GCDualSenseAdaptiveTriggerStatusUnknown GCDualSenseAdaptiveTriggerStatus = -1
	// GCDualSenseAdaptiveTriggerStatusVibrationIsVibrating: The trigger is in vibration mode and is vibrating.
	GCDualSenseAdaptiveTriggerStatusVibrationIsVibrating GCDualSenseAdaptiveTriggerStatus = 6
	// GCDualSenseAdaptiveTriggerStatusVibrationNotVibrating: The trigger is in vibration mode, but isn’t vibrating.
	GCDualSenseAdaptiveTriggerStatusVibrationNotVibrating GCDualSenseAdaptiveTriggerStatus = 5
	// GCDualSenseAdaptiveTriggerStatusWeaponFired: The trigger is in weapon mode, has fired, and has stopped applying the resistive load.
	GCDualSenseAdaptiveTriggerStatusWeaponFired GCDualSenseAdaptiveTriggerStatus = 4
	// GCDualSenseAdaptiveTriggerStatusWeaponFiring: The trigger is in weapon mode, firing, and is applying the resistive load.
	GCDualSenseAdaptiveTriggerStatusWeaponFiring GCDualSenseAdaptiveTriggerStatus = 3
	// GCDualSenseAdaptiveTriggerStatusWeaponReady: The trigger is in weapon mode and ready to fire, but isn’t applying the resistive load.
	GCDualSenseAdaptiveTriggerStatusWeaponReady GCDualSenseAdaptiveTriggerStatus = 2
)

func (e GCDualSenseAdaptiveTriggerStatus) String() string {
	switch e {
	case GCDualSenseAdaptiveTriggerStatusFeedbackLoadApplied:
		return "GCDualSenseAdaptiveTriggerStatusFeedbackLoadApplied"
	case GCDualSenseAdaptiveTriggerStatusFeedbackNoLoad:
		return "GCDualSenseAdaptiveTriggerStatusFeedbackNoLoad"
	case GCDualSenseAdaptiveTriggerStatusSlopeFeedbackApplyingLoad:
		return "GCDualSenseAdaptiveTriggerStatusSlopeFeedbackApplyingLoad"
	case GCDualSenseAdaptiveTriggerStatusSlopeFeedbackFinished:
		return "GCDualSenseAdaptiveTriggerStatusSlopeFeedbackFinished"
	case GCDualSenseAdaptiveTriggerStatusSlopeFeedbackReady:
		return "GCDualSenseAdaptiveTriggerStatusSlopeFeedbackReady"
	case GCDualSenseAdaptiveTriggerStatusUnknown:
		return "GCDualSenseAdaptiveTriggerStatusUnknown"
	case GCDualSenseAdaptiveTriggerStatusVibrationIsVibrating:
		return "GCDualSenseAdaptiveTriggerStatusVibrationIsVibrating"
	case GCDualSenseAdaptiveTriggerStatusVibrationNotVibrating:
		return "GCDualSenseAdaptiveTriggerStatusVibrationNotVibrating"
	case GCDualSenseAdaptiveTriggerStatusWeaponFired:
		return "GCDualSenseAdaptiveTriggerStatusWeaponFired"
	case GCDualSenseAdaptiveTriggerStatusWeaponFiring:
		return "GCDualSenseAdaptiveTriggerStatusWeaponFiring"
	case GCDualSenseAdaptiveTriggerStatusWeaponReady:
		return "GCDualSenseAdaptiveTriggerStatusWeaponReady"
	default:
		return fmt.Sprintf("GCDualSenseAdaptiveTriggerStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GameController/GCExtendedGamepadSnapshotDataVersion
type GCExtendedGamepadSnapshotDataVersion int

const (
	GCExtendedGamepadSnapshotDataVersion1 GCExtendedGamepadSnapshotDataVersion = 0x100
	GCExtendedGamepadSnapshotDataVersion2 GCExtendedGamepadSnapshotDataVersion = 0x101
)

func (e GCExtendedGamepadSnapshotDataVersion) String() string {
	switch e {
	case GCExtendedGamepadSnapshotDataVersion1:
		return "GCExtendedGamepadSnapshotDataVersion1"
	case GCExtendedGamepadSnapshotDataVersion2:
		return "GCExtendedGamepadSnapshotDataVersion2"
	default:
		return fmt.Sprintf("GCExtendedGamepadSnapshotDataVersion(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GameController/GCMicroGamepadSnapshotDataVersion
type GCMicroGamepadSnapshotDataVersion int

const (
	GCMicroGamepadSnapshotDataVersion1 GCMicroGamepadSnapshotDataVersion = 0x100
)

func (e GCMicroGamepadSnapshotDataVersion) String() string {
	switch e {
	case GCMicroGamepadSnapshotDataVersion1:
		return "GCMicroGamepadSnapshotDataVersion1"
	default:
		return fmt.Sprintf("GCMicroGamepadSnapshotDataVersion(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GameController/GCPhysicalInputSourceDirection
type GCPhysicalInputSourceDirection uint

const (
	// GCPhysicalInputSourceDirectionDown: The physical input source supports the down direction.
	GCPhysicalInputSourceDirectionDown GCPhysicalInputSourceDirection = 4
	// GCPhysicalInputSourceDirectionLeft: The physical input source supports the left direction.
	GCPhysicalInputSourceDirectionLeft GCPhysicalInputSourceDirection = 8
	// GCPhysicalInputSourceDirectionNotApplicable: The physical input source doesn’t support directions.
	GCPhysicalInputSourceDirectionNotApplicable GCPhysicalInputSourceDirection = 0
	// GCPhysicalInputSourceDirectionRight: The physical input source supports the right direction.
	GCPhysicalInputSourceDirectionRight GCPhysicalInputSourceDirection = 2
	// GCPhysicalInputSourceDirectionUp: The physical input source contains a value for the up direction.
	GCPhysicalInputSourceDirectionUp GCPhysicalInputSourceDirection = 1
)

func (e GCPhysicalInputSourceDirection) String() string {
	switch e {
	case GCPhysicalInputSourceDirectionDown:
		return "GCPhysicalInputSourceDirectionDown"
	case GCPhysicalInputSourceDirectionLeft:
		return "GCPhysicalInputSourceDirectionLeft"
	case GCPhysicalInputSourceDirectionNotApplicable:
		return "GCPhysicalInputSourceDirectionNotApplicable"
	case GCPhysicalInputSourceDirectionRight:
		return "GCPhysicalInputSourceDirectionRight"
	case GCPhysicalInputSourceDirectionUp:
		return "GCPhysicalInputSourceDirectionUp"
	default:
		return fmt.Sprintf("GCPhysicalInputSourceDirection(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GameController/GCControllerElement/SystemGestureState
type GCSystemGestureState int

const (
	// GCSystemGestureStateAlwaysReceive: A state that sends input to your app and a gesture recognizer simultaneously.
	GCSystemGestureStateAlwaysReceive GCSystemGestureState = 1
	// GCSystemGestureStateDisabled: A state that sends input to your app directly and not to a gesture recognizer.
	GCSystemGestureStateDisabled GCSystemGestureState = 2
	// GCSystemGestureStateEnabled: A state that sends input to your app only after a gesture recognizer doesn’t identify a gesture.
	GCSystemGestureStateEnabled GCSystemGestureState = 0
)

func (e GCSystemGestureState) String() string {
	switch e {
	case GCSystemGestureStateAlwaysReceive:
		return "GCSystemGestureStateAlwaysReceive"
	case GCSystemGestureStateDisabled:
		return "GCSystemGestureStateDisabled"
	case GCSystemGestureStateEnabled:
		return "GCSystemGestureStateEnabled"
	default:
		return fmt.Sprintf("GCSystemGestureState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GameController/GCControllerTouchpad/TouchState-swift.enum
type GCTouchState int

const (
	// GCTouchStateDown: The user starts touching the surface.
	GCTouchStateDown GCTouchState = 1
	// GCTouchStateMoving: The user continues touching the surface.
	GCTouchStateMoving GCTouchState = 2
	// GCTouchStateUp: The user stops or isn’t touching the surface.
	GCTouchStateUp GCTouchState = 0
)

func (e GCTouchState) String() string {
	switch e {
	case GCTouchStateDown:
		return "GCTouchStateDown"
	case GCTouchStateMoving:
		return "GCTouchStateMoving"
	case GCTouchStateUp:
		return "GCTouchStateUp"
	default:
		return fmt.Sprintf("GCTouchState(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/GameController/GCUIEventTypes
type GCUIEventTypes uint

const (
	GCUIEventTypeGamepad GCUIEventTypes = 0
	// GCUIEventTypeStylus: A constant that represents events from a stylus.
	GCUIEventTypeStylus GCUIEventTypes = 0
)

func (e GCUIEventTypes) String() string {
	switch e {
	case GCUIEventTypeGamepad:
		return "GCUIEventTypeGamepad"
	default:
		return fmt.Sprintf("GCUIEventTypes(%d)", e)
	}
}
