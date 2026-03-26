// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"fmt"
)

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCCaptureDynamicRange
type SCCaptureDynamicRange int

const (
	// SCCaptureDynamicRangeHDRCanonicalDisplay: Specifies that the system captures the screen in high dynamic range with attributes of the canonical display.
	SCCaptureDynamicRangeHDRCanonicalDisplay SCCaptureDynamicRange = 2
	// SCCaptureDynamicRangeHDRLocalDisplay: Specifies that the system captures the screen in high dynamic range with attributes of the local display.
	SCCaptureDynamicRangeHDRLocalDisplay SCCaptureDynamicRange = 1
	// SCCaptureDynamicRangeSDR: Specifies that the system captures the screen in standard dynamic range.
	SCCaptureDynamicRangeSDR SCCaptureDynamicRange = 0
)

func (e SCCaptureDynamicRange) String() string {
	switch e {
	case SCCaptureDynamicRangeHDRCanonicalDisplay:
		return "SCCaptureDynamicRangeHDRCanonicalDisplay"
	case SCCaptureDynamicRangeHDRLocalDisplay:
		return "SCCaptureDynamicRangeHDRLocalDisplay"
	case SCCaptureDynamicRangeSDR:
		return "SCCaptureDynamicRangeSDR"
	default:
		return fmt.Sprintf("SCCaptureDynamicRange(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCCaptureResolutionType
type SCCaptureResolutionType int

const (
	// SCCaptureResolutionAutomatic: Allow ScreenCaptureKit to automatically select the quality of content depending on factors such as network connection.
	SCCaptureResolutionAutomatic SCCaptureResolutionType = 0
	// SCCaptureResolutionBest: Capture streaming content at the best available resolution.
	SCCaptureResolutionBest SCCaptureResolutionType = 1
	// SCCaptureResolutionNominal: Capture streaming content with a one point to one pixel conversion factor.
	SCCaptureResolutionNominal SCCaptureResolutionType = 2
)

func (e SCCaptureResolutionType) String() string {
	switch e {
	case SCCaptureResolutionAutomatic:
		return "SCCaptureResolutionAutomatic"
	case SCCaptureResolutionBest:
		return "SCCaptureResolutionBest"
	case SCCaptureResolutionNominal:
		return "SCCaptureResolutionNominal"
	default:
		return fmt.Sprintf("SCCaptureResolutionType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentSharingPickerMode
type SCContentSharingPickerMode uint

const (
	// SCContentSharingPickerModeMultipleApplications: The mode allowing the selection of multiple applications through the presented picker.
	SCContentSharingPickerModeMultipleApplications SCContentSharingPickerMode = 8
	// SCContentSharingPickerModeMultipleWindows: The mode allowing the selection of multiple windows through the presented picker.
	SCContentSharingPickerModeMultipleWindows SCContentSharingPickerMode = 2
	// SCContentSharingPickerModeSingleApplication: The mode allowing the selection of a single application through the presented picker.
	SCContentSharingPickerModeSingleApplication SCContentSharingPickerMode = 4
	// SCContentSharingPickerModeSingleDisplay: The mode allowing the selection of a single display through the presented picker.
	SCContentSharingPickerModeSingleDisplay SCContentSharingPickerMode = 16
	// SCContentSharingPickerModeSingleWindow: The mode allowing the selection of a single window through the presented picker.
	SCContentSharingPickerModeSingleWindow SCContentSharingPickerMode = 1
)

func (e SCContentSharingPickerMode) String() string {
	switch e {
	case SCContentSharingPickerModeMultipleApplications:
		return "SCContentSharingPickerModeMultipleApplications"
	case SCContentSharingPickerModeMultipleWindows:
		return "SCContentSharingPickerModeMultipleWindows"
	case SCContentSharingPickerModeSingleApplication:
		return "SCContentSharingPickerModeSingleApplication"
	case SCContentSharingPickerModeSingleDisplay:
		return "SCContentSharingPickerModeSingleDisplay"
	case SCContentSharingPickerModeSingleWindow:
		return "SCContentSharingPickerModeSingleWindow"
	default:
		return fmt.Sprintf("SCContentSharingPickerMode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCFrameStatus
type SCFrameStatus int

const (
	// SCFrameStatusBlank: A status that indicates the system didn’t generate a new frame because the display is blank.
	SCFrameStatusBlank SCFrameStatus = 2
	// SCFrameStatusComplete: A status that indicates the system successfully generated a new frame.
	SCFrameStatusComplete SCFrameStatus = 0
	// SCFrameStatusIdle: A status that indicates the system didn’t generate a new frame because the display didn’t change.
	SCFrameStatusIdle SCFrameStatus = 1
	// SCFrameStatusStarted: A status that indicates the frame is the first one sent after the stream starts.
	SCFrameStatusStarted SCFrameStatus = 4
	// SCFrameStatusStopped: A status that indicates the frame is in a stopped state.
	SCFrameStatusStopped SCFrameStatus = 5
	// SCFrameStatusSuspended: A status that indicates the system didn’t generate a new frame because you suspended updates.
	SCFrameStatusSuspended SCFrameStatus = 3
)

func (e SCFrameStatus) String() string {
	switch e {
	case SCFrameStatusBlank:
		return "SCFrameStatusBlank"
	case SCFrameStatusComplete:
		return "SCFrameStatusComplete"
	case SCFrameStatusIdle:
		return "SCFrameStatusIdle"
	case SCFrameStatusStarted:
		return "SCFrameStatusStarted"
	case SCFrameStatusStopped:
		return "SCFrameStatusStopped"
	case SCFrameStatusSuspended:
		return "SCFrameStatusSuspended"
	default:
		return fmt.Sprintf("SCFrameStatus(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCPresenterOverlayAlertSetting
type SCPresenterOverlayAlertSetting int

const (
	// SCPresenterOverlayAlertSettingAlways: Always display an alert when using Presenter Overlay.
	SCPresenterOverlayAlertSettingAlways SCPresenterOverlayAlertSetting = 2
	// SCPresenterOverlayAlertSettingNever: Never display an alert when using Presenter Overlay.
	SCPresenterOverlayAlertSettingNever SCPresenterOverlayAlertSetting = 1
	// SCPresenterOverlayAlertSettingSystem: Displays an alert when using Presenter Overlay based on the System Settings.
	SCPresenterOverlayAlertSettingSystem SCPresenterOverlayAlertSetting = 0
)

func (e SCPresenterOverlayAlertSetting) String() string {
	switch e {
	case SCPresenterOverlayAlertSettingAlways:
		return "SCPresenterOverlayAlertSettingAlways"
	case SCPresenterOverlayAlertSettingNever:
		return "SCPresenterOverlayAlertSettingNever"
	case SCPresenterOverlayAlertSettingSystem:
		return "SCPresenterOverlayAlertSettingSystem"
	default:
		return fmt.Sprintf("SCPresenterOverlayAlertSetting(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/DisplayIntent-swift.enum
type SCScreenshotDisplayIntent int

const (
	// SCScreenshotDisplayIntentCanonical: Specifies that the screenshot renders with canonical display attributes optimizing output for presentation on a high dynamic range display.
	SCScreenshotDisplayIntentCanonical SCScreenshotDisplayIntent = 0
	// SCScreenshotDisplayIntentLocal: Specifies that the screenshot renders with local display attributes optimizing output for presentation on the capture display.
	SCScreenshotDisplayIntentLocal SCScreenshotDisplayIntent = 1
)

func (e SCScreenshotDisplayIntent) String() string {
	switch e {
	case SCScreenshotDisplayIntentCanonical:
		return "SCScreenshotDisplayIntentCanonical"
	case SCScreenshotDisplayIntentLocal:
		return "SCScreenshotDisplayIntentLocal"
	default:
		return fmt.Sprintf("SCScreenshotDisplayIntent(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/DynamicRange-swift.enum
type SCScreenshotDynamicRange int

const (
	// SCScreenshotDynamicRangeHDR: Returns a high dynamic range image to the client.
	SCScreenshotDynamicRangeHDR SCScreenshotDynamicRange = 1
	// SCScreenshotDynamicRangeSDR: Returns a standard dynamic range image to the client.
	SCScreenshotDynamicRangeSDR SCScreenshotDynamicRange = 0
	// SCScreenshotDynamicRangeSDRAndHDR: Returns both standard dynamic range and high dynamic range image versions to the client.
	SCScreenshotDynamicRangeSDRAndHDR SCScreenshotDynamicRange = 2
)

func (e SCScreenshotDynamicRange) String() string {
	switch e {
	case SCScreenshotDynamicRangeHDR:
		return "SCScreenshotDynamicRangeHDR"
	case SCScreenshotDynamicRangeSDR:
		return "SCScreenshotDynamicRangeSDR"
	case SCScreenshotDynamicRangeSDRAndHDR:
		return "SCScreenshotDynamicRangeSDRAndHDR"
	default:
		return fmt.Sprintf("SCScreenshotDynamicRange(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCShareableContentStyle
type SCShareableContentStyle int

const (
	// SCShareableContentStyleApplication: The stream is currently presenting one or more applications.
	SCShareableContentStyleApplication SCShareableContentStyle = 3
	// SCShareableContentStyleDisplay: The stream is currently presenting a complete display.
	SCShareableContentStyleDisplay SCShareableContentStyle = 2
	// SCShareableContentStyleNone: The stream isn’t currently presenting any content.
	SCShareableContentStyleNone SCShareableContentStyle = 0
	// SCShareableContentStyleWindow: The stream is currently presenting one or more windows.
	SCShareableContentStyleWindow SCShareableContentStyle = 1
)

func (e SCShareableContentStyle) String() string {
	switch e {
	case SCShareableContentStyleApplication:
		return "SCShareableContentStyleApplication"
	case SCShareableContentStyleDisplay:
		return "SCShareableContentStyleDisplay"
	case SCShareableContentStyleNone:
		return "SCShareableContentStyleNone"
	case SCShareableContentStyleWindow:
		return "SCShareableContentStyleWindow"
	default:
		return fmt.Sprintf("SCShareableContentStyle(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/Preset
type SCStreamConfigurationPreset int

const (
	SCStreamConfigurationPresetCaptureHDRRecordingPreservedSDRHDR10 SCStreamConfigurationPreset = 4
	SCStreamConfigurationPresetCaptureHDRScreenshotCanonicalDisplay SCStreamConfigurationPreset = 3
	SCStreamConfigurationPresetCaptureHDRScreenshotLocalDisplay SCStreamConfigurationPreset = 2
	SCStreamConfigurationPresetCaptureHDRStreamCanonicalDisplay SCStreamConfigurationPreset = 1
	SCStreamConfigurationPresetCaptureHDRStreamLocalDisplay SCStreamConfigurationPreset = 0
)

func (e SCStreamConfigurationPreset) String() string {
	switch e {
	case SCStreamConfigurationPresetCaptureHDRRecordingPreservedSDRHDR10:
		return "SCStreamConfigurationPresetCaptureHDRRecordingPreservedSDRHDR10"
	case SCStreamConfigurationPresetCaptureHDRScreenshotCanonicalDisplay:
		return "SCStreamConfigurationPresetCaptureHDRScreenshotCanonicalDisplay"
	case SCStreamConfigurationPresetCaptureHDRScreenshotLocalDisplay:
		return "SCStreamConfigurationPresetCaptureHDRScreenshotLocalDisplay"
	case SCStreamConfigurationPresetCaptureHDRStreamCanonicalDisplay:
		return "SCStreamConfigurationPresetCaptureHDRStreamCanonicalDisplay"
	case SCStreamConfigurationPresetCaptureHDRStreamLocalDisplay:
		return "SCStreamConfigurationPresetCaptureHDRStreamLocalDisplay"
	default:
		return fmt.Sprintf("SCStreamConfigurationPreset(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamError/Code
type SCStreamErrorCode int

const (
	// SCStreamErrorAttemptToConfigState: An error message that indicates a stream couldn’t update its configuration.
	SCStreamErrorAttemptToConfigState SCStreamErrorCode = -3810
	// SCStreamErrorAttemptToStartStreamState: An error message that indicates a stream is already running or doesn’t exist when trying to start a stream.
	SCStreamErrorAttemptToStartStreamState SCStreamErrorCode = -3807
	// SCStreamErrorAttemptToStopStreamState: An error message that indicates a stream is already stopped or doesn’t exist when trying to stop a stream.
	SCStreamErrorAttemptToStopStreamState SCStreamErrorCode = -3808
	// SCStreamErrorAttemptToUpdateFilterState: An error message that indicates a stream couldn’t update its content filter.
	SCStreamErrorAttemptToUpdateFilterState SCStreamErrorCode = -3809
	// SCStreamErrorFailedApplicationConnectionInterrupted: An error message that indicates there was an interruption in a connection to an app.
	SCStreamErrorFailedApplicationConnectionInterrupted SCStreamErrorCode = -3805
	// SCStreamErrorFailedApplicationConnectionInvalid: An error message that indicates the stream lost its connection to an app.
	SCStreamErrorFailedApplicationConnectionInvalid SCStreamErrorCode = -3804
	// SCStreamErrorFailedNoMatchingApplicationContext: An error message that indicates there isn’t a matching app context for streaming.
	SCStreamErrorFailedNoMatchingApplicationContext SCStreamErrorCode = -3806
	// SCStreamErrorFailedToStart: An error message that indicates a stream failed to start.
	SCStreamErrorFailedToStart SCStreamErrorCode = -3802
	// SCStreamErrorFailedToStartAudioCapture: An error message that indicates an audio stream failed to start.
	SCStreamErrorFailedToStartAudioCapture SCStreamErrorCode = -3818
	// SCStreamErrorFailedToStartMicrophoneCapture: An error message that indicates microphone capture failed to start.
	SCStreamErrorFailedToStartMicrophoneCapture SCStreamErrorCode = -3820
	// SCStreamErrorFailedToStopAudioCapture: An error message that indicates an audio stream failed to stop.
	SCStreamErrorFailedToStopAudioCapture SCStreamErrorCode = -3819
	// SCStreamErrorInternalError: An error message that indicates a stream can’t start due to a failure in ScreenCaptureKit’s internals.
	SCStreamErrorInternalError SCStreamErrorCode = -3811
	// SCStreamErrorInvalidParameter: An error message that indicates an operation failed because of an invalid parameter value.
	SCStreamErrorInvalidParameter SCStreamErrorCode = -3812
	// SCStreamErrorMissingEntitlements: An error message that indicates missing entitlements in your app.
	SCStreamErrorMissingEntitlements SCStreamErrorCode = -3803
	// SCStreamErrorNoCaptureSource: An error message that indicates a stream doesn’t have a source to capture.
	SCStreamErrorNoCaptureSource SCStreamErrorCode = -3815
	// SCStreamErrorNoDisplayList: An error message that indicates a stream doesn’t have displays available.
	SCStreamErrorNoDisplayList SCStreamErrorCode = -3814
	// SCStreamErrorNoWindowList: An error message that indicates a stream doesn’t have windows available.
	SCStreamErrorNoWindowList SCStreamErrorCode = -3813
	// SCStreamErrorRemovingStream: An error message that indicates a stream wasn’t removed.
	SCStreamErrorRemovingStream SCStreamErrorCode = -3816
	// SCStreamErrorSystemStoppedStream: An error message that indicates the system stopped the stream.
	SCStreamErrorSystemStoppedStream SCStreamErrorCode = -3821
	// SCStreamErrorUserDeclined: An error message that indicates the user didn’t grant Screen Recording permission to your app.
	SCStreamErrorUserDeclined SCStreamErrorCode = -3801
	// SCStreamErrorUserStopped: An error message that indicates the user stopped the stream.
	SCStreamErrorUserStopped SCStreamErrorCode = -3817
)

func (e SCStreamErrorCode) String() string {
	switch e {
	case SCStreamErrorAttemptToConfigState:
		return "SCStreamErrorAttemptToConfigState"
	case SCStreamErrorAttemptToStartStreamState:
		return "SCStreamErrorAttemptToStartStreamState"
	case SCStreamErrorAttemptToStopStreamState:
		return "SCStreamErrorAttemptToStopStreamState"
	case SCStreamErrorAttemptToUpdateFilterState:
		return "SCStreamErrorAttemptToUpdateFilterState"
	case SCStreamErrorFailedApplicationConnectionInterrupted:
		return "SCStreamErrorFailedApplicationConnectionInterrupted"
	case SCStreamErrorFailedApplicationConnectionInvalid:
		return "SCStreamErrorFailedApplicationConnectionInvalid"
	case SCStreamErrorFailedNoMatchingApplicationContext:
		return "SCStreamErrorFailedNoMatchingApplicationContext"
	case SCStreamErrorFailedToStart:
		return "SCStreamErrorFailedToStart"
	case SCStreamErrorFailedToStartAudioCapture:
		return "SCStreamErrorFailedToStartAudioCapture"
	case SCStreamErrorFailedToStartMicrophoneCapture:
		return "SCStreamErrorFailedToStartMicrophoneCapture"
	case SCStreamErrorFailedToStopAudioCapture:
		return "SCStreamErrorFailedToStopAudioCapture"
	case SCStreamErrorInternalError:
		return "SCStreamErrorInternalError"
	case SCStreamErrorInvalidParameter:
		return "SCStreamErrorInvalidParameter"
	case SCStreamErrorMissingEntitlements:
		return "SCStreamErrorMissingEntitlements"
	case SCStreamErrorNoCaptureSource:
		return "SCStreamErrorNoCaptureSource"
	case SCStreamErrorNoDisplayList:
		return "SCStreamErrorNoDisplayList"
	case SCStreamErrorNoWindowList:
		return "SCStreamErrorNoWindowList"
	case SCStreamErrorRemovingStream:
		return "SCStreamErrorRemovingStream"
	case SCStreamErrorSystemStoppedStream:
		return "SCStreamErrorSystemStoppedStream"
	case SCStreamErrorUserDeclined:
		return "SCStreamErrorUserDeclined"
	case SCStreamErrorUserStopped:
		return "SCStreamErrorUserStopped"
	default:
		return fmt.Sprintf("SCStreamErrorCode(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamOutputType
type SCStreamOutputType int

const (
	// SCStreamOutputTypeAudio: An output type that represents an audio capture sample buffer.
	SCStreamOutputTypeAudio SCStreamOutputType = 1
	// SCStreamOutputTypeScreen: An output type that represents a screen capture sample buffer.
	SCStreamOutputTypeScreen SCStreamOutputType = 0
)

func (e SCStreamOutputType) String() string {
	switch e {
	case SCStreamOutputTypeAudio:
		return "SCStreamOutputTypeAudio"
	case SCStreamOutputTypeScreen:
		return "SCStreamOutputTypeScreen"
	default:
		return fmt.Sprintf("SCStreamOutputType(%d)", e)
	}
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamType
type SCStreamType int

const (
	// SCStreamTypeDisplay: The stream is currently on a complete display.
	SCStreamTypeDisplay SCStreamType = 1
	// SCStreamTypeWindow: The stream is currently presented as a window.
	SCStreamTypeWindow SCStreamType = 0
)

func (e SCStreamType) String() string {
	switch e {
	case SCStreamTypeDisplay:
		return "SCStreamTypeDisplay"
	case SCStreamTypeWindow:
		return "SCStreamTypeWindow"
	default:
		return fmt.Sprintf("SCStreamType(%d)", e)
	}
}

