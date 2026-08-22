// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/Accessibility/AXChartDescriptor/ContentDirection-swift.enum
type AXChartDescriptorContentDirection int

const (
	// AXChartContentDirectionBottomToTop: A content direction with an x-axis that increases from bottom to top.
	AXChartContentDirectionBottomToTop AXChartDescriptorContentDirection = 3
	// AXChartContentDirectionLeftToRight: A content direction with an x-axis that increases from left to right.
	AXChartContentDirectionLeftToRight AXChartDescriptorContentDirection = 0
	// AXChartContentDirectionRadialClockwise: A content direction with a radial x-axis that increases clockwise.
	AXChartContentDirectionRadialClockwise AXChartDescriptorContentDirection = 4
	// AXChartContentDirectionRadialCounterClockwise: A content direction with a radial x-axis that increases counterclockwise.
	AXChartContentDirectionRadialCounterClockwise AXChartDescriptorContentDirection = 5
	// AXChartContentDirectionRightToLeft: A content direction with an x-axis that increases from right to left.
	AXChartContentDirectionRightToLeft AXChartDescriptorContentDirection = 1
	// AXChartContentDirectionTopToBottom: A content direction with an x-axis that increases from top to bottom.
	AXChartContentDirectionTopToBottom AXChartDescriptorContentDirection = 2
)

func (e AXChartDescriptorContentDirection) String() string {
	switch e {
	case AXChartContentDirectionBottomToTop:
		return "AXChartContentDirectionBottomToTop"
	case AXChartContentDirectionLeftToRight:
		return "AXChartContentDirectionLeftToRight"
	case AXChartContentDirectionRadialClockwise:
		return "AXChartContentDirectionRadialClockwise"
	case AXChartContentDirectionRadialCounterClockwise:
		return "AXChartContentDirectionRadialCounterClockwise"
	case AXChartContentDirectionRightToLeft:
		return "AXChartContentDirectionRightToLeft"
	case AXChartContentDirectionTopToBottom:
		return "AXChartContentDirectionTopToBottom"
	default:
		return fmt.Sprintf("AXChartDescriptorContentDirection(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accessibility/AXCustomContent/Importance-swift.enum
type AXCustomContentImportance uint

const (
	// AXCustomContentImportanceDefault: Output the content to the user on demand.
	AXCustomContentImportanceDefault AXCustomContentImportance = 0
	// AXCustomContentImportanceHigh: Output the content to the user immediately.
	AXCustomContentImportanceHigh AXCustomContentImportance = 1
)

func (e AXCustomContentImportance) String() string {
	switch e {
	case AXCustomContentImportanceDefault:
		return "AXCustomContentImportanceDefault"
	case AXCustomContentImportanceHigh:
		return "AXCustomContentImportanceHigh"
	default:
		return fmt.Sprintf("AXCustomContentImportance(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSessionError-swift.struct/Code
type AXFeatureOverrideSessionError int

const ()

// See: https://developer.apple.com/documentation/Accessibility/AXFeatureOverrideSession/Options
type AXFeatureOverrideSessionOptions uint

const ()

// See: https://developer.apple.com/documentation/Accessibility/AXMFiHearingDevice/Ear
type AXHearingDeviceEar uint

const (
	// AXHearingDeviceEarBoth: A constant that represents both ears.
	AXHearingDeviceEarBoth AXHearingDeviceEar = 6
	// AXHearingDeviceEarLeft: A constant that represents the left ear.
	AXHearingDeviceEarLeft AXHearingDeviceEar = 2
	// AXHearingDeviceEarNone: A constant that represents neither ear.
	AXHearingDeviceEarNone AXHearingDeviceEar = 0
	// AXHearingDeviceEarRight: A constant that represents the right ear.
	AXHearingDeviceEarRight AXHearingDeviceEar = 4
)

func (e AXHearingDeviceEar) String() string {
	switch e {
	case AXHearingDeviceEarBoth:
		return "AXHearingDeviceEarBoth"
	case AXHearingDeviceEarLeft:
		return "AXHearingDeviceEarLeft"
	case AXHearingDeviceEarNone:
		return "AXHearingDeviceEarNone"
	case AXHearingDeviceEarRight:
		return "AXHearingDeviceEarRight"
	default:
		return fmt.Sprintf("AXHearingDeviceEar(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accessibility/AXNumericDataAxisDescriptor/ScaleType-swift.enum
type AXNumericDataAxisDescriptorScale int

const (
	// AXScaleTypeLinear: A linear scale.
	AXScaleTypeLinear AXNumericDataAxisDescriptorScale = 0
	// AXScaleTypeLn: A natural log scale.
	AXScaleTypeLn AXNumericDataAxisDescriptorScale = 2
	// AXScaleTypeLog10: A log scale.
	AXScaleTypeLog10 AXNumericDataAxisDescriptorScale = 1
)

func (e AXNumericDataAxisDescriptorScale) String() string {
	switch e {
	case AXScaleTypeLinear:
		return "AXScaleTypeLinear"
	case AXScaleTypeLn:
		return "AXScaleTypeLn"
	case AXScaleTypeLog10:
		return "AXScaleTypeLog10"
	default:
		return fmt.Sprintf("AXNumericDataAxisDescriptorScale(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/Accessibility/AccessibilitySettings/Feature
type AXSettingsFeature int

const (
	AXSettingsFeatureAllowAppsToAddAudioToCalls AXSettingsFeature = 2
	AXSettingsFeatureAssistiveTouch             AXSettingsFeature = 3
	AXSettingsFeatureAssistiveTouchDevices      AXSettingsFeature = 4
	AXSettingsFeatureCaptionStyles              AXSettingsFeature = 6
	AXSettingsFeatureDwellControl               AXSettingsFeature = 5
	// AXSettingsFeaturePersonalVoiceAllowAppsToRequestToUse: A constant for opening the Settings app to the setting for Personal Voice > Allow Apps to Request to Use.
	AXSettingsFeaturePersonalVoiceAllowAppsToRequestToUse AXSettingsFeature = 1
)

func (e AXSettingsFeature) String() string {
	switch e {
	case AXSettingsFeatureAllowAppsToAddAudioToCalls:
		return "AXSettingsFeatureAllowAppsToAddAudioToCalls"
	case AXSettingsFeatureAssistiveTouch:
		return "AXSettingsFeatureAssistiveTouch"
	case AXSettingsFeatureAssistiveTouchDevices:
		return "AXSettingsFeatureAssistiveTouchDevices"
	case AXSettingsFeatureCaptionStyles:
		return "AXSettingsFeatureCaptionStyles"
	case AXSettingsFeatureDwellControl:
		return "AXSettingsFeatureDwellControl"
	case AXSettingsFeaturePersonalVoiceAllowAppsToRequestToUse:
		return "AXSettingsFeaturePersonalVoiceAllowAppsToRequestToUse"
	default:
		return fmt.Sprintf("AXSettingsFeature(%d)", e)
	}
}
