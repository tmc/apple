// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [AVCaptureDevice] class.
var (
	_AVCaptureDeviceClass     AVCaptureDeviceClass
	_AVCaptureDeviceClassOnce sync.Once
)

func getAVCaptureDeviceClass() AVCaptureDeviceClass {
	_AVCaptureDeviceClassOnce.Do(func() {
		_AVCaptureDeviceClass = AVCaptureDeviceClass{class: objc.GetClass("AVCaptureDevice")}
	})
	return _AVCaptureDeviceClass
}

// GetAVCaptureDeviceClass returns the class object for AVCaptureDevice.
func GetAVCaptureDeviceClass() AVCaptureDeviceClass {
	return getAVCaptureDeviceClass()
}

type AVCaptureDeviceClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureDeviceClass) Alloc() AVCaptureDevice {
	rv := objc.Send[AVCaptureDevice](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a hardware or virtual capture device like a
// camera or microphone.
//
// # Overview
// 
// Capture devices provide media data to capture session inputs that you
// connect to an [AVCaptureSession]. An individual device can provide one or
// more streams of media of a particular type.
// 
// You don’t create capture device instances directly. Instead, retrieve
// them using an instance of [AVCaptureDeviceDiscoverySession], or by calling
// the [AVCaptureDevice.DefaultDeviceWithDeviceTypeMediaTypePosition] method.
// 
// A capture device provides several configuration options. Before attempting
// to configure device properties, such as its focus mode, exposure mode, and
// so on, you must first acquire a lock on the device by calling the
// [AVCaptureDevice.LockForConfiguration] method. You should also query the device’s
// capabilities to ensure that the new modes you intend to set are valid for
// the device. You can then set the properties and release the lock using the
// [AVCaptureDevice.UnlockForConfiguration] method. You may hold the lock if you want all
// settable device properties to remain unchanged. However, holding the device
// lock unnecessarily may degrade capture quality in other apps sharing the
// device and isn’t recommended.
//
// # Identifying a device
//
//   - [AVCaptureDevice.UniqueID]: An identifier that uniquely identifies the device.
//   - [AVCaptureDevice.ModelID]: A model identifier for the device.
//   - [AVCaptureDevice.LocalizedName]: A localized device name for display in the user interface.
//   - [AVCaptureDevice.Manufacturer]: A human-readable string for the manufacturer of the device.
//   - [AVCaptureDevice.DeviceType]: The type of device, such as a built-in microphone or wide-angle camera.
//   - [AVCaptureDevice.Position]: The physical position of the capture device hardware.
//
// # Accessing device state
//
//   - [AVCaptureDevice.Connected]: A Boolean value that indicates whether a device is currently connected to the system and available for use.
//   - [AVCaptureDevice.Suspended]: A Boolean value that indicates whether the device is in a suspended state.
//   - [AVCaptureDevice.InUseByAnotherApplication]: A Boolean value that indicates whether another app is using the device.
//
// # Inspecting device characteristics
//
//   - [AVCaptureDevice.HasMediaType]: Returns a Boolean value that indicates whether the device captures media of a particular type.
//   - [AVCaptureDevice.TransportType]: The transport type of the device.
//   - [AVCaptureDevice.SupportsAVCaptureSessionPreset]: Returns a Boolean value that indicates whether you can use the device with capture session configured with the specified preset.
//
// # Configuring camera hardware
//
//   - [AVCaptureDevice.LockForConfiguration]: Requests exclusive access to configure device hardware properties.
//   - [AVCaptureDevice.UnlockForConfiguration]: Releases exclusive control over device hardware properties.
//
// # Configuring Cinematic video
//
//   - [AVCaptureDevice.SetCinematicVideoFixedFocusAtPointFocusMode]: Fix focus at a distance.
//   - [AVCaptureDevice.SetCinematicVideoTrackingFocusAtPointFocusMode]: Focus on and start tracking an object if it can be detected at the region specified by the point.
//   - [AVCaptureDevice.SetCinematicVideoTrackingFocusWithDetectedObjectIDFocusMode]: Focus on and start tracking a detected object.
//   - [AVCaptureDevice.NotEnoughLight]: The light level of the current scene is insufficient for the current set of features to function optimally.
//   - [AVCaptureDevice.CinematicVideoCaptureSceneMonitoringStatuses]: The current scene monitoring statuses related to Cinematic Video capture.
//
// # Enabling automatic frame rate
//
//   - [AVCaptureDevice.AutoVideoFrameRateEnabled]: A Boolean value that indicates whether the capture device performs automatic video frame rate adjustments.
//   - [AVCaptureDevice.SetAutoVideoFrameRateEnabled]
//
// # Supporting spatial capture
//
//   - [AVCaptureDevice.SpatialCaptureDiscomfortReasons]: Reasons why current environmental conditions aren’t suitable to capturing spatial videos that are comfortable to view.
//
// # Supporting Continuity Camera
//
//   - [AVCaptureDevice.ContinuityCamera]: A Boolean value that indicates whether the device is a Continuity Camera.
//   - [AVCaptureDevice.CompanionDeskViewCamera]: A Desk View camera associated with a device.
//
// # Monitoring system pressure
//
//   - [AVCaptureDevice.AVCaptureSessionInterruptionSystemPressureStateKey]: A key to retrieve a state value that indicates the system pressure level and contributing factors that caused the interruption.
//
// # Restricting camera switching
//
//   - [AVCaptureDevice.SetPrimaryConstituentDeviceSwitchingBehaviorRestrictedSwitchingBehaviorConditions]: Sets the switching behavior of the primary constituent device.
//   - [AVCaptureDevice.PrimaryConstituentDeviceSwitchingBehavior]: The switching behavior for the primary constituent device.
//   - [AVCaptureDevice.PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions]: The conditions that restrict the primary constituent device’s switching behavior.
//   - [AVCaptureDevice.ActivePrimaryConstituentDeviceSwitchingBehavior]: The switching behavior of the active constituent device.
//   - [AVCaptureDevice.ActivePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions]: The conditions that restrict camera switching behavior for the active primary constituent device.
//   - [AVCaptureDevice.ActivePrimaryConstituentDevice]: A virtual device’s active primary constituent device.
//   - [AVCaptureDevice.SupportedFallbackPrimaryConstituentDevices]: The constituent devices available to select as a fallback for a longer focal length primary constituent device.
//   - [AVCaptureDevice.FallbackPrimaryConstituentDevices]: The fallback devices to use when a constituent device with a longer focal length becomes limited by its light sensitivity or minimum focus distance.
//   - [AVCaptureDevice.SetFallbackPrimaryConstituentDevices]
//
// # Configuring lens smudge detection
//
//   - [AVCaptureDevice.CameraLensSmudgeDetectionEnabled]: Whether camera lens smudge detection is enabled.
//   - [AVCaptureDevice.SetCameraLensSmudgeDetectionEnabledDetectionInterval]: Specify whether to enable camera lens smudge detection, and the interval time between each run of detections.
//   - [AVCaptureDevice.CameraLensSmudgeDetectionInterval]: The camera lens smudge detection interval.
//   - [AVCaptureDevice.CameraLensSmudgeDetectionStatus]: A value specifying the status of camera lens smudge detection.
//
// # Synchronizing with external devices
//
//   - [AVCaptureDevice.FollowingExternalSyncDevice]: Whether the device is following an external sync device.
//   - [AVCaptureDevice.MinSupportedExternalSyncFrameDuration]: The minimum frame duration that can be passed as the `videoFrameDuration` when directing your device input to follow an external sync device.
//   - [AVCaptureDevice.VideoFrameDurationLocked]: Whether the device’s video frame rate (expressed as a duration) is currently locked.
//   - [AVCaptureDevice.MinSupportedLockedVideoFrameDuration]: The maximum frame rate (expressed as a minimum duration) that can be set on an input associated with this device.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice
type AVCaptureDevice struct {
	objectivec.Object
}

// AVCaptureDeviceFromID constructs a [AVCaptureDevice] from an objc.ID.
//
// An object that represents a hardware or virtual capture device like a
// camera or microphone.
func AVCaptureDeviceFromID(id objc.ID) AVCaptureDevice {
	return AVCaptureDevice{objectivec.Object{ID: id}}
}
// NOTE: AVCaptureDevice adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureDevice] class.
//
// # Identifying a device
//
//   - [IAVCaptureDevice.UniqueID]: An identifier that uniquely identifies the device.
//   - [IAVCaptureDevice.ModelID]: A model identifier for the device.
//   - [IAVCaptureDevice.LocalizedName]: A localized device name for display in the user interface.
//   - [IAVCaptureDevice.Manufacturer]: A human-readable string for the manufacturer of the device.
//   - [IAVCaptureDevice.DeviceType]: The type of device, such as a built-in microphone or wide-angle camera.
//   - [IAVCaptureDevice.Position]: The physical position of the capture device hardware.
//
// # Accessing device state
//
//   - [IAVCaptureDevice.Connected]: A Boolean value that indicates whether a device is currently connected to the system and available for use.
//   - [IAVCaptureDevice.Suspended]: A Boolean value that indicates whether the device is in a suspended state.
//   - [IAVCaptureDevice.InUseByAnotherApplication]: A Boolean value that indicates whether another app is using the device.
//
// # Inspecting device characteristics
//
//   - [IAVCaptureDevice.HasMediaType]: Returns a Boolean value that indicates whether the device captures media of a particular type.
//   - [IAVCaptureDevice.TransportType]: The transport type of the device.
//   - [IAVCaptureDevice.SupportsAVCaptureSessionPreset]: Returns a Boolean value that indicates whether you can use the device with capture session configured with the specified preset.
//
// # Configuring camera hardware
//
//   - [IAVCaptureDevice.LockForConfiguration]: Requests exclusive access to configure device hardware properties.
//   - [IAVCaptureDevice.UnlockForConfiguration]: Releases exclusive control over device hardware properties.
//
// # Configuring Cinematic video
//
//   - [IAVCaptureDevice.SetCinematicVideoFixedFocusAtPointFocusMode]: Fix focus at a distance.
//   - [IAVCaptureDevice.SetCinematicVideoTrackingFocusAtPointFocusMode]: Focus on and start tracking an object if it can be detected at the region specified by the point.
//   - [IAVCaptureDevice.SetCinematicVideoTrackingFocusWithDetectedObjectIDFocusMode]: Focus on and start tracking a detected object.
//   - [IAVCaptureDevice.NotEnoughLight]: The light level of the current scene is insufficient for the current set of features to function optimally.
//   - [IAVCaptureDevice.CinematicVideoCaptureSceneMonitoringStatuses]: The current scene monitoring statuses related to Cinematic Video capture.
//
// # Enabling automatic frame rate
//
//   - [IAVCaptureDevice.AutoVideoFrameRateEnabled]: A Boolean value that indicates whether the capture device performs automatic video frame rate adjustments.
//   - [IAVCaptureDevice.SetAutoVideoFrameRateEnabled]
//
// # Supporting spatial capture
//
//   - [IAVCaptureDevice.SpatialCaptureDiscomfortReasons]: Reasons why current environmental conditions aren’t suitable to capturing spatial videos that are comfortable to view.
//
// # Supporting Continuity Camera
//
//   - [IAVCaptureDevice.ContinuityCamera]: A Boolean value that indicates whether the device is a Continuity Camera.
//   - [IAVCaptureDevice.CompanionDeskViewCamera]: A Desk View camera associated with a device.
//
// # Monitoring system pressure
//
//   - [IAVCaptureDevice.AVCaptureSessionInterruptionSystemPressureStateKey]: A key to retrieve a state value that indicates the system pressure level and contributing factors that caused the interruption.
//
// # Restricting camera switching
//
//   - [IAVCaptureDevice.SetPrimaryConstituentDeviceSwitchingBehaviorRestrictedSwitchingBehaviorConditions]: Sets the switching behavior of the primary constituent device.
//   - [IAVCaptureDevice.PrimaryConstituentDeviceSwitchingBehavior]: The switching behavior for the primary constituent device.
//   - [IAVCaptureDevice.PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions]: The conditions that restrict the primary constituent device’s switching behavior.
//   - [IAVCaptureDevice.ActivePrimaryConstituentDeviceSwitchingBehavior]: The switching behavior of the active constituent device.
//   - [IAVCaptureDevice.ActivePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions]: The conditions that restrict camera switching behavior for the active primary constituent device.
//   - [IAVCaptureDevice.ActivePrimaryConstituentDevice]: A virtual device’s active primary constituent device.
//   - [IAVCaptureDevice.SupportedFallbackPrimaryConstituentDevices]: The constituent devices available to select as a fallback for a longer focal length primary constituent device.
//   - [IAVCaptureDevice.FallbackPrimaryConstituentDevices]: The fallback devices to use when a constituent device with a longer focal length becomes limited by its light sensitivity or minimum focus distance.
//   - [IAVCaptureDevice.SetFallbackPrimaryConstituentDevices]
//
// # Configuring lens smudge detection
//
//   - [IAVCaptureDevice.CameraLensSmudgeDetectionEnabled]: Whether camera lens smudge detection is enabled.
//   - [IAVCaptureDevice.SetCameraLensSmudgeDetectionEnabledDetectionInterval]: Specify whether to enable camera lens smudge detection, and the interval time between each run of detections.
//   - [IAVCaptureDevice.CameraLensSmudgeDetectionInterval]: The camera lens smudge detection interval.
//   - [IAVCaptureDevice.CameraLensSmudgeDetectionStatus]: A value specifying the status of camera lens smudge detection.
//
// # Synchronizing with external devices
//
//   - [IAVCaptureDevice.FollowingExternalSyncDevice]: Whether the device is following an external sync device.
//   - [IAVCaptureDevice.MinSupportedExternalSyncFrameDuration]: The minimum frame duration that can be passed as the `videoFrameDuration` when directing your device input to follow an external sync device.
//   - [IAVCaptureDevice.VideoFrameDurationLocked]: Whether the device’s video frame rate (expressed as a duration) is currently locked.
//   - [IAVCaptureDevice.MinSupportedLockedVideoFrameDuration]: The maximum frame rate (expressed as a minimum duration) that can be set on an input associated with this device.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice
type IAVCaptureDevice interface {
	objectivec.IObject

	// Topic: Identifying a device

	// An identifier that uniquely identifies the device.
	UniqueID() string
	// A model identifier for the device.
	ModelID() string
	// A localized device name for display in the user interface.
	LocalizedName() string
	// A human-readable string for the manufacturer of the device.
	Manufacturer() string
	// The type of device, such as a built-in microphone or wide-angle camera.
	DeviceType() AVCaptureDeviceType
	// The physical position of the capture device hardware.
	Position() AVCaptureDevicePosition

	// Topic: Accessing device state

	// A Boolean value that indicates whether a device is currently connected to the system and available for use.
	Connected() bool
	// A Boolean value that indicates whether the device is in a suspended state.
	Suspended() bool
	// A Boolean value that indicates whether another app is using the device.
	InUseByAnotherApplication() bool

	// Topic: Inspecting device characteristics

	// Returns a Boolean value that indicates whether the device captures media of a particular type.
	HasMediaType(mediaType AVMediaType) bool
	// The transport type of the device.
	TransportType() int32
	// Returns a Boolean value that indicates whether you can use the device with capture session configured with the specified preset.
	SupportsAVCaptureSessionPreset(preset AVCaptureSessionPreset) bool

	// Topic: Configuring camera hardware

	// Requests exclusive access to configure device hardware properties.
	LockForConfiguration() (bool, error)
	// Releases exclusive control over device hardware properties.
	UnlockForConfiguration()

	// Topic: Configuring Cinematic video

	// Fix focus at a distance.
	SetCinematicVideoFixedFocusAtPointFocusMode(point corefoundation.CGPoint, focusMode AVCaptureCinematicVideoFocusMode)
	// Focus on and start tracking an object if it can be detected at the region specified by the point.
	SetCinematicVideoTrackingFocusAtPointFocusMode(point corefoundation.CGPoint, focusMode AVCaptureCinematicVideoFocusMode)
	// Focus on and start tracking a detected object.
	SetCinematicVideoTrackingFocusWithDetectedObjectIDFocusMode(detectedObjectID int, focusMode AVCaptureCinematicVideoFocusMode)
	// The light level of the current scene is insufficient for the current set of features to function optimally.
	NotEnoughLight() AVCaptureSceneMonitoringStatus
	// The current scene monitoring statuses related to Cinematic Video capture.
	CinematicVideoCaptureSceneMonitoringStatuses() foundation.INSSet

	// Topic: Enabling automatic frame rate

	// A Boolean value that indicates whether the capture device performs automatic video frame rate adjustments.
	AutoVideoFrameRateEnabled() bool
	SetAutoVideoFrameRateEnabled(value bool)

	// Topic: Supporting spatial capture

	// Reasons why current environmental conditions aren’t suitable to capturing spatial videos that are comfortable to view.
	SpatialCaptureDiscomfortReasons() foundation.INSSet

	// Topic: Supporting Continuity Camera

	// A Boolean value that indicates whether the device is a Continuity Camera.
	ContinuityCamera() bool
	// A Desk View camera associated with a device.
	CompanionDeskViewCamera() IAVCaptureDevice

	// Topic: Monitoring system pressure

	// A key to retrieve a state value that indicates the system pressure level and contributing factors that caused the interruption.
	AVCaptureSessionInterruptionSystemPressureStateKey() string

	// Topic: Restricting camera switching

	// Sets the switching behavior of the primary constituent device.
	SetPrimaryConstituentDeviceSwitchingBehaviorRestrictedSwitchingBehaviorConditions(switchingBehavior AVCapturePrimaryConstituentDeviceSwitchingBehavior, restrictedSwitchingBehaviorConditions AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions)
	// The switching behavior for the primary constituent device.
	PrimaryConstituentDeviceSwitchingBehavior() AVCapturePrimaryConstituentDeviceSwitchingBehavior
	// The conditions that restrict the primary constituent device’s switching behavior.
	PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions() AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions
	// The switching behavior of the active constituent device.
	ActivePrimaryConstituentDeviceSwitchingBehavior() AVCapturePrimaryConstituentDeviceSwitchingBehavior
	// The conditions that restrict camera switching behavior for the active primary constituent device.
	ActivePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions() AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions
	// A virtual device’s active primary constituent device.
	ActivePrimaryConstituentDevice() IAVCaptureDevice
	// The constituent devices available to select as a fallback for a longer focal length primary constituent device.
	SupportedFallbackPrimaryConstituentDevices() []AVCaptureDevice
	// The fallback devices to use when a constituent device with a longer focal length becomes limited by its light sensitivity or minimum focus distance.
	FallbackPrimaryConstituentDevices() []AVCaptureDevice
	SetFallbackPrimaryConstituentDevices(value []AVCaptureDevice)

	// Topic: Configuring lens smudge detection

	// Whether camera lens smudge detection is enabled.
	CameraLensSmudgeDetectionEnabled() bool
	// Specify whether to enable camera lens smudge detection, and the interval time between each run of detections.
	SetCameraLensSmudgeDetectionEnabledDetectionInterval(cameraLensSmudgeDetectionEnabled bool, detectionInterval objectivec.IObject)
	// The camera lens smudge detection interval.
	CameraLensSmudgeDetectionInterval() objectivec.IObject
	// A value specifying the status of camera lens smudge detection.
	CameraLensSmudgeDetectionStatus() AVCaptureCameraLensSmudgeDetectionStatus

	// Topic: Synchronizing with external devices

	// Whether the device is following an external sync device.
	FollowingExternalSyncDevice() bool
	// The minimum frame duration that can be passed as the `videoFrameDuration` when directing your device input to follow an external sync device.
	MinSupportedExternalSyncFrameDuration() objectivec.IObject
	// Whether the device’s video frame rate (expressed as a duration) is currently locked.
	VideoFrameDurationLocked() bool
	// The maximum frame rate (expressed as a minimum duration) that can be set on an input associated with this device.
	MinSupportedLockedVideoFrameDuration() objectivec.IObject

	// The currently active color space for capture.
	ActiveColorSpace() AVCaptureColorSpace
	SetActiveColorSpace(value AVCaptureColorSpace)
	// The capture format in use by the device.
	ActiveFormat() IAVCaptureDeviceFormat
	SetActiveFormat(value IAVCaptureDeviceFormat)
	// The currently active input source of the device.
	ActiveInputSource() IAVCaptureDeviceInputSource
	SetActiveInputSource(value IAVCaptureDeviceInputSource)
	// The currently active maximum frame duration.
	ActiveVideoMaxFrameDuration() objectivec.IObject
	SetActiveVideoMaxFrameDuration(value objectivec.IObject)
	// The currently active minimum frame duration.
	ActiveVideoMinFrameDuration() objectivec.IObject
	SetActiveVideoMinFrameDuration(value objectivec.IObject)
	// A Boolean value that indicates whether the device is currently adjusting its exposure setting.
	AdjustingExposure() bool
	// A Boolean value that indicates whether the device is currently adjusting its focus setting.
	AdjustingFocus() bool
	// A Boolean value that indicates whether the device is currently adjusting the white balance.
	AdjustingWhiteBalance() bool
	// A set of reactions types that a device supports performing.
	AvailableReactionTypes() foundation.INSSet
	// A Boolean value that indicates whether Background Replacement is currently active on a capture device.
	BackgroundReplacementActive() bool
	// A Boolean value that indicates whether you can perform reaction effects on a capture device.
	CanPerformReactionEffects() bool
	// A Boolean value that indicates whether Center Stage is active on a device.
	CenterStageActive() bool
	// The effective region within the output pixel buffer to perform Center Stage framing.
	CenterStageRectOfInterest() corefoundation.CGRect
	SetCenterStageRectOfInterest(value corefoundation.CGRect)
	CenterStageRectOfInterestSupported() bool
	// A video zoom factor multiplier to use when displaying zoom information in a user interface.
	DisplayVideoZoomFactorMultiplier() float64
	// The exposure mode for the device.
	ExposureMode() AVCaptureExposureMode
	SetExposureMode(value AVCaptureExposureMode)
	// The point of interest for exposure.
	ExposurePointOfInterest() corefoundation.CGPoint
	SetExposurePointOfInterest(value corefoundation.CGPoint)
	// A Boolean value that indicates whether the device supports a point of interest for exposure.
	ExposurePointOfInterestSupported() bool
	// The device’s current exposure rectangle of interest, if it has one.
	ExposureRectOfInterest() corefoundation.CGRect
	SetExposureRectOfInterest(value corefoundation.CGRect)
	// Whether the device supports exposure rectangles of interest.
	ExposureRectOfInterestSupported() bool
	// A Boolean value that indicates whether the flash is currently available for use.
	FlashAvailable() bool
	// The device’s current flash mode.
	FlashMode() AVCaptureFlashMode
	SetFlashMode(value AVCaptureFlashMode)
	// The capture device’s focus mode.
	FocusMode() AVCaptureFocusMode
	SetFocusMode(value AVCaptureFocusMode)
	// The point of interest for focusing.
	FocusPointOfInterest() corefoundation.CGPoint
	SetFocusPointOfInterest(value corefoundation.CGPoint)
	// A Boolean value that indicates whether the device supports a point of interest for focus.
	FocusPointOfInterestSupported() bool
	// The device’s current focus rectangle of interest, if it has one.
	FocusRectOfInterest() corefoundation.CGRect
	SetFocusRectOfInterest(value corefoundation.CGRect)
	// Whether the receiver supports focus rectangles of interest.
	FocusRectOfInterestSupported() bool
	// The capture formats a device supports.
	Formats() []AVCaptureDeviceFormat
	// A Boolean value that indicates whether the capture device has a flash.
	HasFlash() bool
	// A Boolean value that specifies whether the capture device has a torch.
	HasTorch() bool
	// An array of input sources that the device supports.
	InputSources() []AVCaptureDeviceInputSource
	// An array of capture devices that are physically linked to a device.
	LinkedDevices() []AVCaptureDevice
	// The minimum size you may use when specifying a rectangle of interest.
	MinExposureRectOfInterestSize() corefoundation.CGSize
	// The minimum size you may use when specifying a rectangle of interest.
	MinFocusRectOfInterestSize() corefoundation.CGSize
	// The capture device’s minimum focus distance in millimeters.
	MinimumFocusDistance() int
	// A Boolean value that indicates whether the Portrait video effect is active on a device.
	PortraitEffectActive() bool
	// An array of reaction effects that the device is currently performing, sorted by timestamp.
	ReactionEffectsInProgress() []AVCaptureReactionEffectState
	// A Boolean value that indicates whether Studio Light is active on a device.
	StudioLightActive() bool
	// A Boolean value that indicates whether the device’s torch is currently active.
	TorchActive() bool
	// A Boolean value that indicates whether the torch is currently available for use.
	TorchAvailable() bool
	// The current torch brightness level.
	TorchLevel() float32
	// The current torch mode.
	TorchMode() AVCaptureTorchMode
	SetTorchMode(value AVCaptureTorchMode)
	// The current playback mode.
	TransportControlsPlaybackMode() AVCaptureDeviceTransportControlsPlaybackMode
	// The current playback speed.
	TransportControlsSpeed() AVCaptureDeviceTransportControlsSpeed
	// A Boolean value that indicates whether the device supports transport control commands.
	TransportControlsSupported() bool
	// The current white balance mode.
	WhiteBalanceMode() AVCaptureWhiteBalanceMode
	SetWhiteBalanceMode(value AVCaptureWhiteBalanceMode)
	// The default rectangle of interest used for a given exposure point of interest.
	DefaultRectForExposurePointOfInterest(pointOfInterest corefoundation.CGPoint) corefoundation.CGRect
	// The default rectangle of interest used for a given focus point of interest.
	DefaultRectForFocusPointOfInterest(pointOfInterest corefoundation.CGPoint) corefoundation.CGRect
	// Returns a Boolean value that indicates whether a device supports the specified exposure mode.
	IsExposureModeSupported(exposureMode AVCaptureExposureMode) bool
	// Returns a Boolean value that indicates whether the device supports the specified focus mode.
	IsFocusModeSupported(focusMode AVCaptureFocusMode) bool
	// Returns a Boolean value that indicates whether the device supports the specified torch mode.
	IsTorchModeSupported(torchMode AVCaptureTorchMode) bool
	// Returns a Boolean value that indicates whether the device supports the specified white balance mode.
	IsWhiteBalanceModeSupported(whiteBalanceMode AVCaptureWhiteBalanceMode) bool
	// Performs the specified reaction type on the video stream.
	PerformEffectForReaction(reactionType AVCaptureReactionType)
	// Sets the illumination level when in torch mode.
	SetTorchModeOnWithLevelError(torchLevel float32) (bool, error)
	// Sets the transport control’s playback mode and speed.
	SetTransportControlsPlaybackModeSpeed(mode AVCaptureDeviceTransportControlsPlaybackMode, speed AVCaptureDeviceTransportControlsSpeed)
}

// Init initializes the instance.
func (c AVCaptureDevice) Init() AVCaptureDevice {
	rv := objc.Send[AVCaptureDevice](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureDevice) Autorelease() AVCaptureDevice {
	rv := objc.Send[AVCaptureDevice](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureDevice creates a new AVCaptureDevice instance.
func NewAVCaptureDevice() AVCaptureDevice {
	class := getAVCaptureDeviceClass()
	rv := objc.Send[AVCaptureDevice](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an object that represents a device with the specified identifier.
//
// deviceUniqueID: An identifier that uniquely identifies the device.
//
// # Return Value
// 
// A capture device, or `nil` if no device with the specified identifier
// exists.
//
// # Discussion
// 
// Every capture device has a unique identifier that persists on a system
// across device connections, app restarts, and reboots of the system itself.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/init(uniqueID:)
func NewCaptureDeviceWithUniqueID(deviceUniqueID string) AVCaptureDevice {
	rv := objc.Send[objc.ID](objc.ID(getAVCaptureDeviceClass().class), objc.Sel("deviceWithUniqueID:"), objc.String(deviceUniqueID))
	return AVCaptureDeviceFromID(rv)
}

// Returns a Boolean value that indicates whether the device captures media of
// a particular type.
//
// mediaType: A media type, such as [video], [audio], or [muxed].
// //
// [audio]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/audio
// [muxed]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/muxed
// [video]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/video
//
// # Return Value
// 
// [true] if the device captures media of the specified type; otherwise,
// [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/hasMediaType(_:)
func (c AVCaptureDevice) HasMediaType(mediaType AVMediaType) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasMediaType:"), objc.String(string(mediaType)))
	return rv
}
// Returns a Boolean value that indicates whether you can use the device with
// capture session configured with the specified preset.
//
// preset: A capture session preset.
//
// # Return Value
// 
// [true] if you can use the device; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/supportsSessionPreset(_:)
func (c AVCaptureDevice) SupportsAVCaptureSessionPreset(preset AVCaptureSessionPreset) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("supportsAVCaptureSessionPreset:"), objc.String(string(preset)))
	return rv
}
// Requests exclusive access to configure device hardware properties.
//
// # Discussion
// 
// To set hardware properties on a capture device, such as the [FocusMode] and
// [ExposureMode], your app must first acquire a lock on the device. Only hold
// the device lock if your app requires settable device properties to remain
// unchanged. Holding the device lock unnecessarily may degrade capture
// quality in other apps sharing the device.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/lockForConfiguration()
func (c AVCaptureDevice) LockForConfiguration() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("lockForConfiguration:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("lockForConfiguration: returned NO with nil NSError")
	}
	return rv, nil

}
// Releases exclusive control over device hardware properties.
//
// # Discussion
// 
// If you’ve previously locked a device by calling [LockForConfiguration],
// call this method when your app no longer requires preventing device
// properties from changing automatically.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/unlockForConfiguration()
func (c AVCaptureDevice) UnlockForConfiguration() {
	objc.Send[objc.ID](c.ID, objc.Sel("unlockForConfiguration"))
}
// Fix focus at a distance.
//
// point: A normalized point of interest (i.e., [0,1]) in the coordinate space of the
// device.
//
// focusMode: Specify whether to focus strongly or weakly.
//
// # Discussion
// 
// The distance at which focus is set is determined internally using signals
// such as depth data.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setCinematicVideoFixedFocus(at:focusMode:)
func (c AVCaptureDevice) SetCinematicVideoFixedFocusAtPointFocusMode(point corefoundation.CGPoint, focusMode AVCaptureCinematicVideoFocusMode) {
	objc.Send[objc.ID](c.ID, objc.Sel("setCinematicVideoFixedFocusAtPoint:focusMode:"), point, focusMode)
}
// Focus on and start tracking an object if it can be detected at the region
// specified by the point.
//
// point: A normalized point of interest (i.e., [0,1]) in the coordinate space of the
// device.
//
// focusMode: Specify whether to focus strongly or weakly.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setCinematicVideoTrackingFocus(at:focusMode:)
func (c AVCaptureDevice) SetCinematicVideoTrackingFocusAtPointFocusMode(point corefoundation.CGPoint, focusMode AVCaptureCinematicVideoFocusMode) {
	objc.Send[objc.ID](c.ID, objc.Sel("setCinematicVideoTrackingFocusAtPoint:focusMode:"), point, focusMode)
}
// Focus on and start tracking a detected object.
//
// detectedObjectID: The ID of the detected object.
//
// focusMode: Specify whether to focus strongly or weakly.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setCinematicVideoTrackingFocus(detectedObjectID:focusMode:)
func (c AVCaptureDevice) SetCinematicVideoTrackingFocusWithDetectedObjectIDFocusMode(detectedObjectID int, focusMode AVCaptureCinematicVideoFocusMode) {
	objc.Send[objc.ID](c.ID, objc.Sel("setCinematicVideoTrackingFocusWithDetectedObjectID:focusMode:"), detectedObjectID, focusMode)
}
// Sets the switching behavior of the primary constituent device.
//
// switchingBehavior: The switching behavior to set on the device.
//
// restrictedSwitchingBehaviorConditions: Sets the conditions during which the system restricts switching cameras.
// 
// Setting the switching behavior to a value other than
// [CapturePrimaryConstituentDeviceSwitchingBehaviorRestricted] requires that
// you set this argument to an empty option set.
//
// # Discussion
// 
// Use this method to configure the camera switching behavior of a capture
// device. Before calling it, determine if a device supports configuring its
// device switching behavior by querying the device’s
// [ActivePrimaryConstituentDeviceSwitchingBehavior] property. If the value
// equals `XCUIElementTypeUnsupported`, attempting to configure its switching
// behavior results in an error.
// 
// When recording using an instance of [AVCaptureMovieFileOutput], you may
// override the switching behavior by calling the movie file output’s
// [SetPrimaryConstituentDeviceSwitchingBehaviorForRecordingRestrictedSwitchingBehaviorConditions]
// method.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setPrimaryConstituentDeviceSwitchingBehavior(_:restrictedSwitchingBehaviorConditions:)
func (c AVCaptureDevice) SetPrimaryConstituentDeviceSwitchingBehaviorRestrictedSwitchingBehaviorConditions(switchingBehavior AVCapturePrimaryConstituentDeviceSwitchingBehavior, restrictedSwitchingBehaviorConditions AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions) {
	objc.Send[objc.ID](c.ID, objc.Sel("setPrimaryConstituentDeviceSwitchingBehavior:restrictedSwitchingBehaviorConditions:"), switchingBehavior, restrictedSwitchingBehaviorConditions)
}
// Specify whether to enable camera lens smudge detection, and the interval
// time between each run of detections.
//
// cameraLensSmudgeDetectionEnabled: Specify whether camera lens smudge detection should be enabled.
//
// detectionInterval: The detection running interval if detection is enabled.
//
// detectionInterval is a [coremedia.CMTime].
//
// # Discussion
// 
// Each run of detection processes frames over a short period, and produces
// one detection result. Use `detectionInterval` to specify the interval time
// between each run of detections. For example, when
// [CameraLensSmudgeDetectionEnabled] is set to `true` and `detectionInterval`
// is set to 1 minute, detection runs once per minute, and updates
// [AVCaptureCameraLensSmudgeDetectionStatus]. If `detectionInterval` is set
// to `kCMTimeInvalid`, detection runs only once after the session starts. If
// `detectionInterval` is set to `kCMTimeZero`, detection runs continuously.
// 
// [AVCaptureDevice] throws an [NSInvalidArgumentException] if the
// [CameraLensSmudgeDetectionSupported] property on the current active format
// returns `false`. Enabling detection requires a lengthy reconfiguration of
// the capture render pipeline, so you should enable detection before calling
// [StartRunning] or within [BeginConfiguration] and [CommitConfiguration]
// while running.
//
// [AVCaptureCameraLensSmudgeDetectionStatus]: https://developer.apple.com/documentation/AVFoundation/AVCaptureCameraLensSmudgeDetectionStatus
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setCameraLensSmudgeDetectionEnabled(_:detectionInterval:)
func (c AVCaptureDevice) SetCameraLensSmudgeDetectionEnabledDetectionInterval(cameraLensSmudgeDetectionEnabled bool, detectionInterval objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("setCameraLensSmudgeDetectionEnabled:detectionInterval:"), cameraLensSmudgeDetectionEnabled, detectionInterval)
}
// The default rectangle of interest used for a given exposure point of
// interest.
//
// pointOfInterest: The point of interest for which you want the default rectangle of interest.
//
// # Discussion
// 
// For example, pass `(0.5, 0.5)` to get the exposure rectangle of interest
// used for the default exposure point of interest at `(0.5, 0.5)`.
// 
// This method returns [CGRectNull] if [ExposureRectOfInterestSupported]
// returns `false`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/defaultRectForExposurePoint(ofInterest:)
func (c AVCaptureDevice) DefaultRectForExposurePointOfInterest(pointOfInterest corefoundation.CGPoint) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("defaultRectForExposurePointOfInterest:"), pointOfInterest)
	return corefoundation.CGRect(rv)
}
// The default rectangle of interest used for a given focus point of interest.
//
// pointOfInterest: The point of interest for which you want the default rectangle of interest.
//
// # Discussion
// 
// For example, pass `(0.5, 0.5)` to get the focus rectangle of interest used
// for the default focus point of interest at `(0.5, 0.5)`.
// 
// This method returns [CGRectNull] if [FocusRectOfInterestSupported] returns
// `false`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/defaultRectForFocusPoint(ofInterest:)
func (c AVCaptureDevice) DefaultRectForFocusPointOfInterest(pointOfInterest corefoundation.CGPoint) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("defaultRectForFocusPointOfInterest:"), pointOfInterest)
	return corefoundation.CGRect(rv)
}
// Returns a Boolean value that indicates whether a device supports the
// specified exposure mode.
//
// exposureMode: An exposure mode to query.
//
// # Return Value
// 
// [true] if `exposureMode` is supported; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isExposureModeSupported(_:)
func (c AVCaptureDevice) IsExposureModeSupported(exposureMode AVCaptureExposureMode) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isExposureModeSupported:"), exposureMode)
	return rv
}
// Returns a Boolean value that indicates whether the device supports the
// specified focus mode.
//
// focusMode: A focus mode to query.
//
// # Return Value
// 
// [true] if the device supports the focus mode; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isFocusModeSupported(_:)
func (c AVCaptureDevice) IsFocusModeSupported(focusMode AVCaptureFocusMode) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isFocusModeSupported:"), focusMode)
	return rv
}
// Returns a Boolean value that indicates whether the device supports the
// specified torch mode.
//
// torchMode: The desired torch mode.
//
// # Return Value
// 
// [true] if the device supports the torch mode; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isTorchModeSupported(_:)
func (c AVCaptureDevice) IsTorchModeSupported(torchMode AVCaptureTorchMode) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isTorchModeSupported:"), torchMode)
	return rv
}
// Returns a Boolean value that indicates whether the device supports the
// specified white balance mode.
//
// whiteBalanceMode: A white balance mode to use.
//
// # Return Value
// 
// [true] if the device supports the white balance mode; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isWhiteBalanceModeSupported(_:)
func (c AVCaptureDevice) IsWhiteBalanceModeSupported(whiteBalanceMode AVCaptureWhiteBalanceMode) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isWhiteBalanceModeSupported:"), whiteBalanceMode)
	return rv
}
// Performs the specified reaction type on the video stream.
//
// reactionType: A reaction type to perform. Specifying a type that doesn’t exists within
// the set of [AvailableReactionTypes] for the device results in an exception.
//
// # Discussion
// 
// The entries in the [ReactionEffectsInProgress] property may not reflect
// one-to-one with calls to this method. Depending on reaction style or
// resource limits, the system may coalesce overlapping reactions of the same
// type by extending an existing reaction rather than overlaying a new one.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/performEffect(for:)
func (c AVCaptureDevice) PerformEffectForReaction(reactionType AVCaptureReactionType) {
	objc.Send[objc.ID](c.ID, objc.Sel("performEffectForReaction:"), objc.String(string(reactionType)))
}
// Sets the illumination level when in torch mode.
//
// torchLevel: The new torch mode level. This value must be a floating-point number
// between `0.0` and `1.0`. To set the torch mode level to the currently
// available maximum, specify the constant [maxAvailableTorchLevel] for this
// parameter.
// //
// [maxAvailableTorchLevel]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/maxAvailableTorchLevel
//
// # Discussion
// 
// This method sets the torch mode to [CaptureTorchModeOn] and sets the level
// to the specified value. If the device doesn’t support this mode or if you
// specify a value for `torchLevel` that’s outside the accepted range, this
// method raises an exception. If the torch value is within the accepted range
// but greater than the currently supported maximum—perhaps because the
// device is overheating—this method returns [false].
// 
// Before changing the value of this property, you must call
// [LockForConfiguration] to acquire exclusive access to the device’s
// configuration properties. Otherwise, calling this method raises an
// exception. When you finish configuring the device, call
// [UnlockForConfiguration] to release the lock and allow other devices to
// configure the settings.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setTorchModeOn(level:)
func (c AVCaptureDevice) SetTorchModeOnWithLevelError(torchLevel float32) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](c.ID, objc.Sel("setTorchModeOnWithLevel:error:"), torchLevel, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("setTorchModeOnWithLevel:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Sets the transport control’s playback mode and speed.
//
// mode: A playback mode constant that indicates whether to put the deck should into
// play mode.
//
// speed: The speed at which to wind or play the tape.
//
// # Discussion
// 
// Before changing the value of this property, you must call
// [LockForConfiguration] to acquire exclusive access to the device’s
// configuration properties. Otherwise, calling this method raises an
// exception. When you’re finished configuring the device, call
// [UnlockForConfiguration] to release the lock and allow other devices to
// configure the settings.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/setTransportControlsPlaybackMode(_:speed:)
func (c AVCaptureDevice) SetTransportControlsPlaybackModeSpeed(mode AVCaptureDeviceTransportControlsPlaybackMode, speed AVCaptureDeviceTransportControlsSpeed) {
	objc.Send[objc.ID](c.ID, objc.Sel("setTransportControlsPlaybackMode:speed:"), mode, speed)
}

// Returns the default device for the specified device type, media type, and
// position.
//
// deviceType: The type of capture device to request, such as [builtInWideAngleCamera].
// //
// [builtInWideAngleCamera]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DeviceType-swift.struct/builtInWideAngleCamera
//
// mediaType: The type of media to request capture of, such as [video] or [audio].
// //
// [audio]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/audio
// [video]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/video
//
// position: The position of capture device to request relative to system hardware
// (front- or back-facing). Pass [CaptureDevicePositionUnspecified] to search
// for devices regardless of position.
//
// # Return Value
// 
// The default system device, or `nil` if no device currently exists that
// satisfies the specified criteria.
//
// # Discussion
// 
// Use this method to select the system default capture device for a given
// scenario. For example, to obtain the dual camera on supported hardware and
// fall back to the standard wide-angle camera otherwise, call this method
// twice, as shown below.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/default(_:for:position:)
func (_AVCaptureDeviceClass AVCaptureDeviceClass) DefaultDeviceWithDeviceTypeMediaTypePosition(deviceType AVCaptureDeviceType, mediaType AVMediaType, position AVCaptureDevicePosition) AVCaptureDevice {
	rv := objc.Send[objc.ID](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("defaultDeviceWithDeviceType:mediaType:position:"), objc.String(string(deviceType)), objc.String(string(mediaType)), position)
	return AVCaptureDeviceFromID(rv)
}
// Returns the default device that captures the specified media type.
//
// mediaType: A media type for the device.
//
// # Return Value
// 
// The default device, or `nil` if no device with that media type exists.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/default(for:)
func (_AVCaptureDeviceClass AVCaptureDeviceClass) DefaultDeviceWithMediaType(mediaType AVMediaType) AVCaptureDevice {
	rv := objc.Send[objc.ID](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("defaultDeviceWithMediaType:"), objc.String(string(mediaType)))
	return AVCaptureDeviceFromID(rv)
}
// Requests the user’s permission to allow the app to capture media of a
// particular type.
//
// mediaType: A media type for which to check the authorization status. The supported
// media types are [video] and [audio].
// //
// [audio]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/audio
// [video]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/video
//
// handler: A callback the system invokes with a Boolean value that indicates whether
// the user granted or denied access to your app.
// 
// Return control to the main queue or [MainActor] before performing user
// interface updates.
// //
// [MainActor]: https://developer.apple.com/documentation/Swift/MainActor
//
// # Discussion
// 
// Capturing media requires explicit permission from the user. An app’s
// default authorization status is [AuthorizationStatusNotDetermined], which
// means the user hasn’t yet granted it permission to capture media. The
// first time you create an [AVCaptureDeviceInput] object for a media type
// that requires permission, the system automatically displays an alert to
// request recording permission. Alternatively, call this method to prompt the
// user at a time of your choosing. The system saves the user’s selection so
// that it doesn’t have to prompt the user again. A user can change their
// authorization status in the Settings app.
// 
// Calling this method doesn’t block the thread while the system is
// prompting the user for access. However, until the grants permission, the
// system only vends black video frames and silent audio samples.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/requestAccess(for:completionHandler:)
func (_AVCaptureDeviceClass AVCaptureDeviceClass) RequestAccessForMediaTypeCompletionHandler(mediaType AVMediaType, handler BoolHandler) {
_block1, _cleanup1 := NewBoolBlock(handler)
	defer _cleanup1()
	objc.Send[objc.ID](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("requestAccessForMediaType:completionHandler:"), mediaType, _block1)
}
// Returns an authorization status that indicates whether the user grants the
// app permission to capture media of a particular type.
//
// mediaType: A media type for which to check the authorization status. The supported
// media types are [video] and [audio].
// //
// [audio]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/audio
// [video]: https://developer.apple.com/documentation/AVFoundation/AVMediaType/video
//
// # Return Value
// 
// An authorization status value.
//
// # Discussion
// 
// A user must explicitly grant your app access to record audio or video. Call
// this method to determine your app’s current authorization status. If it
// returns a value of [AuthorizationStatusNotDetermined], call
// [RequestAccessForMediaTypeCompletionHandler] to prompt the user for capture
// permission.
// 
// After the user grants permission, the system remembers their choice and
// doesn’t prompt them again. However, a user can change their choice at any
// time in the Settings app.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/authorizationStatus(for:)
func (_AVCaptureDeviceClass AVCaptureDeviceClass) AuthorizationStatusForMediaType(mediaType AVMediaType) AVAuthorizationStatus {
	rv := objc.Send[AVAuthorizationStatus](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("authorizationStatusForMediaType:"), objc.String(string(mediaType)))
	return AVAuthorizationStatus(rv)
}
// Displays the system’s user interface to configure video effects or
// microphone modes.
//
// systemUserInterface: The system user interface to present.
//
// # Discussion
// 
// Use this method to prompt the user to make changes to Video Effects (such
// as Center Stage or Portrait Effect) or Microphone Modes. It presents the
// system user interface and deep links to the appropriate module.
// 
// Calling this method isn’t a blocking operation. After the system presents
// the indicated user interface, control returns immediately to the app.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/showSystemUserInterface(_:)
func (_AVCaptureDeviceClass AVCaptureDeviceClass) ShowSystemUserInterface(systemUserInterface AVCaptureSystemUserInterface) {
	objc.Send[objc.ID](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("showSystemUserInterface:"), systemUserInterface)
}

// An identifier that uniquely identifies the device.
//
// # Discussion
// 
// Capture devices have a unique identifier that persists on one system across
// device connections and disconnections, application restarts, and reboots of
// the system itself. You can store the value returned by this property to
// recall or track the status of a specific device in the future.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/uniqueID
func (c AVCaptureDevice) UniqueID() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("uniqueID"))
	return foundation.NSStringFromID(rv).String()
}
// A model identifier for the device.
//
// # Discussion
// 
// The value of this property is an identifier unique to all devices of the
// same model. The value is persistent across device connections and
// disconnections, and across different systems. For example, the model
// identifier of a built-in camera on two identical iPhone models is the same
// even though they’re different physical devices.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/modelID
func (c AVCaptureDevice) ModelID() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("modelID"))
	return foundation.NSStringFromID(rv).String()
}
// A localized device name for display in the user interface.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/localizedName
func (c AVCaptureDevice) LocalizedName() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("localizedName"))
	return foundation.NSStringFromID(rv).String()
}
// A human-readable string for the manufacturer of the device.
//
// # Discussion
// 
// You can use this property to identify capture devices by manufacturer. For
// all Apple devices, the value of this property is `Apple Inc.`
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/manufacturer
func (c AVCaptureDevice) Manufacturer() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("manufacturer"))
	return foundation.NSStringFromID(rv).String()
}
// The type of device, such as a built-in microphone or wide-angle camera.
//
// # Discussion
// 
// Use the [DefaultDeviceWithDeviceTypeMediaTypePosition] method or the
// [AVCaptureDeviceDiscoverySession] class to find capture devices by device
// type.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/deviceType-swift.property
func (c AVCaptureDevice) DeviceType() AVCaptureDeviceType {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("deviceType"))
	return AVCaptureDeviceType(foundation.NSStringFromID(rv).String())
}
// The physical position of the capture device hardware.
//
// # Discussion
// 
// This property value is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/position-swift.property
func (c AVCaptureDevice) Position() AVCaptureDevicePosition {
	rv := objc.Send[AVCaptureDevicePosition](c.ID, objc.Sel("position"))
	return AVCaptureDevicePosition(rv)
}
// A Boolean value that indicates whether a device is currently connected to
// the system and available for use.
//
// # Discussion
// 
// When the value of this property is [false] for a particular capture device
// instance, it doesn’t become [true] again. If the same physical device
// reconnects, the system represents it as a new capture device instance.
// 
// You can key-value observe this property value to monitor when a device is
// no longer available.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isConnected
func (c AVCaptureDevice) Connected() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isConnected"))
	return rv
}
// A Boolean value that indicates whether the device is in a suspended state.
//
// # Discussion
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isSuspended
func (c AVCaptureDevice) Suspended() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isSuspended"))
	return rv
}
// A Boolean value that indicates whether another app is using the device.
//
// # Discussion
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isInUseByAnotherApplication
func (c AVCaptureDevice) InUseByAnotherApplication() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isInUseByAnotherApplication"))
	return rv
}
// The transport type of the device.
//
// # Discussion
// 
// The value of this property represents a capture device’s transport type,
// such as USB or PCI. The value is an IOKit framework transport type constant
// (`kIOAudioDeviceTransportType`).
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/transportType
func (c AVCaptureDevice) TransportType() int32 {
	rv := objc.Send[int32](c.ID, objc.Sel("transportType"))
	return rv
}
// The light level of the current scene is insufficient for the current set of
// features to function optimally.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturescenemonitoringstatus/notenoughlight
func (c AVCaptureDevice) NotEnoughLight() AVCaptureSceneMonitoringStatus {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("AVCaptureSceneMonitoringStatusNotEnoughLight"))
	return AVCaptureSceneMonitoringStatus(foundation.NSStringFromID(rv).String())
}
// The current scene monitoring statuses related to Cinematic Video capture.
//
// # Discussion
// 
// Monitor this property via key-value observation to present a UI informing
// the user that they should reframe their scene for a better Cinematic Video
// experience (“scene is too dark”).
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/cinematicVideoCaptureSceneMonitoringStatuses
func (c AVCaptureDevice) CinematicVideoCaptureSceneMonitoringStatuses() foundation.INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("cinematicVideoCaptureSceneMonitoringStatuses"))
	return foundation.NSSetFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the capture device performs
// automatic video frame rate adjustments.
//
// # Discussion
// 
// You can enable this property on a device when its active format’s
// [AutoVideoFrameRateSupported] property is [true]. When enabled, a capture
// device automatically adjusts the active frame rate based on light level.
// Under low light conditions, it decreases the frame rate to properly expose
// the scene. For formats with a maximum frame rate of 30 fps, the frame rate
// switches between 30-24. For formats with a maximum frame rate of 60 fps,
// the frame rate switches between 60-30-24.
// 
// Changing the device’s active format resets this property to its default
// value of [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isAutoVideoFrameRateEnabled
func (c AVCaptureDevice) AutoVideoFrameRateEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isAutoVideoFrameRateEnabled"))
	return rv
}
func (c AVCaptureDevice) SetAutoVideoFrameRateEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAutoVideoFrameRateEnabled:"), value)
}
// Reasons why current environmental conditions aren’t suitable to capturing
// spatial videos that are comfortable to view.
//
// # Discussion
// 
// You can monitor this property to determine whether to present UI that
// recommends a person reframe their scene for more pleasing spatial capture.
// For example, you could show a message that indicates the subject is too
// close or the scene is too dark.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/spatialCaptureDiscomfortReasons
func (c AVCaptureDevice) SpatialCaptureDiscomfortReasons() foundation.INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("spatialCaptureDiscomfortReasons"))
	return foundation.NSSetFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the device is a Continuity Camera.
//
// # Discussion
// 
// Continuity Camera enables you to use the rear camera system of iPhone as an
// external webcam in macOS.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isContinuityCamera
func (c AVCaptureDevice) ContinuityCamera() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isContinuityCamera"))
	return rv
}
// A Desk View camera associated with a device.
//
// # Discussion
// 
// The value provides an Desk View camera for a device, if one exists, that
// derives its framing from the device’s ultra wide camera. When multiple
// Continuity Camera devices are available on the system, use this property to
// a relate a particular instance with its associated Desk View device.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/companionDeskViewCamera
func (c AVCaptureDevice) CompanionDeskViewCamera() IAVCaptureDevice {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("companionDeskViewCamera"))
	return AVCaptureDeviceFromID(objc.ID(rv))
}
// A key to retrieve a state value that indicates the system pressure level
// and contributing factors that caused the interruption.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturesessioninterruptionsystempressurestatekey
func (c AVCaptureDevice) AVCaptureSessionInterruptionSystemPressureStateKey() string {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("AVCaptureSessionInterruptionSystemPressureStateKey"))
	return foundation.NSStringFromID(rv).String()
}
// The switching behavior for the primary constituent device.
//
// # Discussion
// 
// The default value of this property is
// [CapturePrimaryConstituentDeviceSwitchingBehaviorAuto] for devices that
// support camera switching, and
// [CapturePrimaryConstituentDeviceSwitchingBehaviorUnsupported] for those
// that don’t.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/primaryConstituentDeviceSwitchingBehavior-swift.property
func (c AVCaptureDevice) PrimaryConstituentDeviceSwitchingBehavior() AVCapturePrimaryConstituentDeviceSwitchingBehavior {
	rv := objc.Send[AVCapturePrimaryConstituentDeviceSwitchingBehavior](c.ID, objc.Sel("primaryConstituentDeviceSwitchingBehavior"))
	return AVCapturePrimaryConstituentDeviceSwitchingBehavior(rv)
}
// The conditions that restrict the primary constituent device’s switching
// behavior.
//
// # Discussion
// 
// The default value of this property is
// [CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionNone].
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/primaryConstituentDeviceRestrictedSwitchingBehaviorConditions-swift.property
func (c AVCaptureDevice) PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions() AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions {
	rv := objc.Send[AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions](c.ID, objc.Sel("primaryConstituentDeviceRestrictedSwitchingBehaviorConditions"))
	return AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions(rv)
}
// The switching behavior of the active constituent device.
//
// # Discussion
// 
// For virtual devices with multiple constituent devices, this property
// provides the active switching behavior. It equals the value of the
// [PrimaryConstituentDeviceSwitchingBehavior] property except when you record
// with an [AVCaptureMovieFileOutput] that you configure to use different
// behavior.
// 
// The value of this property is
// [CapturePrimaryConstituentDeviceSwitchingBehaviorUnsupported] for devices
// that don’t support constituent device switching.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activePrimaryConstituentDeviceSwitchingBehavior
func (c AVCaptureDevice) ActivePrimaryConstituentDeviceSwitchingBehavior() AVCapturePrimaryConstituentDeviceSwitchingBehavior {
	rv := objc.Send[AVCapturePrimaryConstituentDeviceSwitchingBehavior](c.ID, objc.Sel("activePrimaryConstituentDeviceSwitchingBehavior"))
	return AVCapturePrimaryConstituentDeviceSwitchingBehavior(rv)
}
// The conditions that restrict camera switching behavior for the active
// primary constituent device.
//
// # Discussion
// 
// For virtual devices with multiple constituent devices, this property
// returns the active restricted switching behavior conditions. This is equal
// to [PrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions] except
// while recording using an [AVCaptureMovieFileOutput] that you configure with
// different restricted switching behavior conditions.
// 
// Devices that don’t support constituent device switching return
// [CapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditionNone].
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions
func (c AVCaptureDevice) ActivePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions() AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions {
	rv := objc.Send[AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions](c.ID, objc.Sel("activePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions"))
	return AVCapturePrimaryConstituentDeviceRestrictedSwitchingBehaviorConditions(rv)
}
// A virtual device’s active primary constituent device.
//
// # Discussion
// 
// The primary constituent device may change when zoom, exposure, or focus
// changes. The value is `nil` for nonvirtual devices.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activePrimaryConstituent
func (c AVCaptureDevice) ActivePrimaryConstituentDevice() IAVCaptureDevice {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("activePrimaryConstituentDevice"))
	return AVCaptureDeviceFromID(objc.ID(rv))
}
// The constituent devices available to select as a fallback for a longer
// focal length primary constituent device.
//
// # Discussion
// 
// The value of this property doesn’t change for a particular virtual
// device, and is `nil` for nonvirtual devices.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/supportedFallbackPrimaryConstituentDevices
func (c AVCaptureDevice) SupportedFallbackPrimaryConstituentDevices() []AVCaptureDevice {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("supportedFallbackPrimaryConstituentDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaptureDevice {
		return AVCaptureDeviceFromID(id)
	})
}
// The fallback devices to use when a constituent device with a longer focal
// length becomes limited by its light sensitivity or minimum focus distance.
//
// # Discussion
// 
// By default, this value contains the array of devices that the
// [SupportedFallbackPrimaryConstituentDevices] property provides. The system
// throws an exception if you attempt to specify a device other than the ones
// found in the device array.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/fallbackPrimaryConstituentDevices
func (c AVCaptureDevice) FallbackPrimaryConstituentDevices() []AVCaptureDevice {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("fallbackPrimaryConstituentDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaptureDevice {
		return AVCaptureDeviceFromID(id)
	})
}
func (c AVCaptureDevice) SetFallbackPrimaryConstituentDevices(value []AVCaptureDevice) {
	objc.Send[struct{}](c.ID, objc.Sel("setFallbackPrimaryConstituentDevices:"), objectivec.IObjectSliceToNSArray(value))
}
// Whether camera lens smudge detection is enabled.
//
// # Discussion
// 
// You enable lens smudge detection by calling
// [SetCameraLensSmudgeDetectionEnabledDetectionInterval]. By default, this
// property is returns `false`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isCameraLensSmudgeDetectionEnabled
func (c AVCaptureDevice) CameraLensSmudgeDetectionEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isCameraLensSmudgeDetectionEnabled"))
	return rv
}
// The camera lens smudge detection interval.
//
// # Discussion
// 
// [CameraLensSmudgeDetectionInterval] is set by calling
// [SetCameraLensSmudgeDetectionEnabledDetectionInterval]. By default, this
// property returns `kCMTimeInvalid`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/cameraLensSmudgeDetectionInterval
func (c AVCaptureDevice) CameraLensSmudgeDetectionInterval() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("cameraLensSmudgeDetectionInterval"))
	return objectivec.Object{ID: rv}
}
// A value specifying the status of camera lens smudge detection.
//
// # Discussion
// 
// During initial detection execution, [CameraLensSmudgeDetectionStatus]
// returns [AVCaptureCameraLensSmudgeDetectionStatusUnknown] until the
// detection result settles. Once a detection result is produced,
// [CameraLensSmudgeDetectionStatus] returns the most recent detection result.
// This property can be key-value observed.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/cameraLensSmudgeDetectionStatus
func (c AVCaptureDevice) CameraLensSmudgeDetectionStatus() AVCaptureCameraLensSmudgeDetectionStatus {
	rv := objc.Send[AVCaptureCameraLensSmudgeDetectionStatus](c.ID, objc.Sel("cameraLensSmudgeDetectionStatus"))
	return AVCaptureCameraLensSmudgeDetectionStatus(rv)
}
// Whether the device is following an external sync device.
//
// # Discussion
// 
// See [FollowExternalSyncDeviceVideoFrameDurationDelegate] for more
// information on external sync.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isFollowingExternalSyncDevice
func (c AVCaptureDevice) FollowingExternalSyncDevice() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isFollowingExternalSyncDevice"))
	return rv
}
// The minimum frame duration that can be passed as the `videoFrameDuration`
// when directing your device input to follow an external sync device.
//
// # Discussion
// 
// Use this property as the minimum allowable frame duration to pass to
// `AVCaptureDeviceInput/` when you want to follow an external sync device.
// This property returns `kCMTimeInvalid` when the device’s’ current
// configuration does not support external sync device following.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/minSupportedExternalSyncFrameDuration
func (c AVCaptureDevice) MinSupportedExternalSyncFrameDuration() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("minSupportedExternalSyncFrameDuration"))
	return objectivec.Object{ID: rv}
}
// Whether the device’s video frame rate (expressed as a duration) is
// currently locked.
//
// # Discussion
// 
// Returns `true` when an [AVCaptureDeviceInput] associated with the device
// has its [ActiveLockedVideoFrameDuration] property set to something other
// than `kCMTimeInvalid`. See [ActiveLockedVideoFrameDuration] for more
// information on video frame duration locking.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isVideoFrameDurationLocked
func (c AVCaptureDevice) VideoFrameDurationLocked() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isVideoFrameDurationLocked"))
	return rv
}
// The maximum frame rate (expressed as a minimum duration) that can be set on
// an input associated with this device.
//
// # Discussion
// 
// `kCMTimeInvalid` is returned when the device or its current configuration
// does not support locked frame rate. Use [ActiveLockedVideoFrameDuration] to
// set the locked frame rate on the input.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/minSupportedLockedVideoFrameDuration
func (c AVCaptureDevice) MinSupportedLockedVideoFrameDuration() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("minSupportedLockedVideoFrameDuration"))
	return objectivec.Object{ID: rv}
}
// The currently active color space for capture.
//
// # Discussion
// 
// All devices and formats support capture in the sRGB color space. Some
// devices and formats can also capture in the P3 color space, which includes
// a much wider gamut of colors than the sRGB color space. By default, a
// capture session automatically enables wide-gamut capture for supported
// devices and capture workflows—for details, see the
// [AutomaticallyConfiguresCaptureDeviceForWideColor] of your capture session.
// To instead set the color space manually, disable that [AVCaptureSession]
// property before setting the active color space.
// 
// For best results, choose a color space before calling [StartRunning] on
// your capture session. Changing this property while a capture session is
// running requires a disruptive reconfiguration of the capture render
// pipeline—movie captures in progress ends immediately, unfulfilled photo
// requests abort, and video preview temporarily freeze.
// 
// Before changing this property, you must call the [LockForConfiguration]
// method to obtain exclusive access to the capture device. Attempting to
// change this property without locking the device raises an exception
// ([genericException]).
//
// [genericException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/genericException
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeColorSpace
func (c AVCaptureDevice) ActiveColorSpace() AVCaptureColorSpace {
	rv := objc.Send[AVCaptureColorSpace](c.ID, objc.Sel("activeColorSpace"))
	return AVCaptureColorSpace(rv)
}
func (c AVCaptureDevice) SetActiveColorSpace(value AVCaptureColorSpace) {
	objc.Send[struct{}](c.ID, objc.Sel("setActiveColorSpace:"), value)
}
// The capture format in use by the device.
//
// # Discussion
// 
// In iOS, a device’s active format and a capture session’s
// [SessionPreset] are mutually exclusive. If you set a device’s active
// format, the session to which it’s attached changes its preset to
// [inputPriority]. Likewise if you set a preset on a capture session, the
// session assumes control of its input devices, and configures their active
// format appropriately.
// 
// Set the [ActiveFormat], [ActiveVideoMinFrameDuration], and
// [ActiveVideoMaxFrameDuration] properties simultaneously by performing the
// configuration between calls to the session’s [BeginConfiguration] and
// [CommitConfiguration] methods.
// 
// If you configure a session to use an active format intended for high
// resolution still photography, and you apply zoom, orientation, or format
// changes to an [AVCaptureVideoDataOutput], the system may not meet the
// target framerate.
// 
// This property is key-value observable.
//
// [inputPriority]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/Preset/inputPriority
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeFormat
func (c AVCaptureDevice) ActiveFormat() IAVCaptureDeviceFormat {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("activeFormat"))
	return AVCaptureDeviceFormatFromID(objc.ID(rv))
}
func (c AVCaptureDevice) SetActiveFormat(value IAVCaptureDeviceFormat) {
	objc.Send[struct{}](c.ID, objc.Sel("setActiveFormat:"), value)
}
// The currently active input source of the device.
//
// # Discussion
// 
// You must call [LockForConfiguration] before attempting to set a format.
// Setting a format that isn’t present in the [InputSources] array results
// in an exception.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeInputSource
func (c AVCaptureDevice) ActiveInputSource() IAVCaptureDeviceInputSource {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("activeInputSource"))
	return AVCaptureDeviceInputSourceFromID(objc.ID(rv))
}
func (c AVCaptureDevice) SetActiveInputSource(value IAVCaptureDeviceInputSource) {
	objc.Send[struct{}](c.ID, objc.Sel("setActiveInputSource:"), value)
}
// The currently active maximum frame duration.
//
// # Discussion
// 
// A device’s maximum frame duration is the reciprocal of its minimum frame
// rate. You can set the value of this property to limit the minimum frame
// rate during a capture session. The capture device automatically chooses a
// default maximum frame duration based on its active format. After changing
// the value of this property, you can return to the default maximum frame
// duration by setting this property’s value to [invalid]. Choosing a new
// preset for the capture session also resets this property to its default
// value.
// 
// Attempting to set this property to a value not found in the active
// format’s [VideoSupportedFrameRateRanges] array raises an exception
// ([InvalidArgumentException]).
// 
// This property value is key-value observable.
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeVideoMaxFrameDuration
func (c AVCaptureDevice) ActiveVideoMaxFrameDuration() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("activeVideoMaxFrameDuration"))
	return objectivec.Object{ID: rv}
}
func (c AVCaptureDevice) SetActiveVideoMaxFrameDuration(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setActiveVideoMaxFrameDuration:"), value)
}
// The currently active minimum frame duration.
//
// # Discussion
// 
// A device’s minimum frame duration is the reciprocal of its maximum frame
// rate. You can set the value of this property to limit the maximum frame
// rate during a capture session. The capture device automatically chooses a
// default minimum frame duration based on its active format. After changing
// the value of this property, you can return to the default minimum frame
// duration by setting this property’s value to [invalid]. Choosing a new
// preset for the capture session also resets this property to its default
// value.
// 
// Attempting to set this property to a value not found in the active
// format’s [VideoSupportedFrameRateRanges] array raises an exception
// ([InvalidArgumentException]).
// 
// This property value is key-value observable.
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeVideoMinFrameDuration
func (c AVCaptureDevice) ActiveVideoMinFrameDuration() objectivec.IObject {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("activeVideoMinFrameDuration"))
	return objectivec.Object{ID: rv}
}
func (c AVCaptureDevice) SetActiveVideoMinFrameDuration(value objectivec.IObject) {
	objc.Send[struct{}](c.ID, objc.Sel("setActiveVideoMinFrameDuration:"), value)
}
// A Boolean value that indicates whether the device is currently adjusting
// its exposure setting.
//
// # Discussion
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isAdjustingExposure
func (c AVCaptureDevice) AdjustingExposure() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isAdjustingExposure"))
	return rv
}
// A Boolean value that indicates whether the device is currently adjusting
// its focus setting.
//
// # Discussion
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isAdjustingFocus
func (c AVCaptureDevice) AdjustingFocus() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isAdjustingFocus"))
	return rv
}
// A Boolean value that indicates whether the device is currently adjusting
// the white balance.
//
// # Discussion
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isAdjustingWhiteBalance
func (c AVCaptureDevice) AdjustingWhiteBalance() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isAdjustingWhiteBalance"))
	return rv
}
// A set of reactions types that a device supports performing.
//
// # Discussion
// 
// The list may differ between devices, and may change for a specific device
// when it’s active format changes.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/availableReactionTypes
func (c AVCaptureDevice) AvailableReactionTypes() foundation.INSSet {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("availableReactionTypes"))
	return foundation.NSSetFromID(objc.ID(rv))
}
// A Boolean value that indicates whether Background Replacement is currently
// active on a capture device.
//
// # Discussion
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isBackgroundReplacementActive
func (c AVCaptureDevice) BackgroundReplacementActive() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isBackgroundReplacementActive"))
	return rv
}
// A Boolean value that indicates whether you can perform reaction effects on
// a capture device.
//
// # Discussion
// 
// This value is [true] when a device’s [ReactionEffectsEnabled] and its
// active format’s [ReactionEffectsSupported] property values are [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/canPerformReactionEffects
func (c AVCaptureDevice) CanPerformReactionEffects() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("canPerformReactionEffects"))
	return rv
}
// A Boolean value that indicates whether Center Stage is active on a device.
//
// # Discussion
// 
// When Center Stage is active, the camera automatically pans, tightens, or
// widens the field of view as it requires to keep people optimally framed. If
// an app or a user enables Center Stage, this property value is [true] if the
// device supports the feature in its current configuration.
// 
// The system imposes the following restrictions on a device when Center Stage
// is active:
// 
// - It limits the range of values the device supports for its
// [MinAvailableVideoZoomFactor] and [MaxAvailableVideoZoomFactor] properties
// to those of the active capture format’s
// [VideoMinZoomFactorForCenterStage] and [VideoMaxZoomFactorForCenterStage],
// respectively. - It limits the [ActiveVideoMinFrameDuration] and
// [ActiveVideoMaxFrameDuration] to the value set by the active capture
// format’s [VideoFrameRateRangeForCenterStage] property.
// 
// The system deactivates Center Stage in the following cases:
// 
// - You enable depth data delivery on a capture output, such as
// [AVCaptureDepthDataOutput] or [AVCapturePhotoOutput]. - The device supports
// geometric distortion correction, but you haven’t enabled it by setting
// the value of [GeometricDistortionCorrectionEnabled] to [true].
// 
// This property is key-value observable.
//
// [AVCaptureDepthDataOutput]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDepthDataOutput
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isCenterStageActive
func (c AVCaptureDevice) CenterStageActive() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isCenterStageActive"))
	return rv
}
// The effective region within the output pixel buffer to perform Center Stage
// framing.
//
// # Discussion
// 
// Apps that require cropping the output from Center Stage can use this
// property to guide how it performs its framing. The rectangle’s origin is
// top-left and is relative to the coordinate space of the output pixel
// buffer. The default value of this property is a rectangle with a normalized
// (`0`-`1`) coordinate space, where (`0`,`0`) represents the top left of the
// picture area, and (`1`,`1`) represents the bottom-right of the unrotated
// picture. The system applies the rectangle of interest prior to rotation,
// mirroring, or scaling.
// 
// Attempting to set a rectangle of interest in the following cases results in
// the system throwing an illegal argument exception:
// 
// - If none of the capture device’s supported formats support Center Stage.
// - If the capture device’s [CenterStageEnabled] property value is [false].
// - If you specify a value that’s outside the normalized (`0`-`1`)
// coordinate space.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/centerStageRectOfInterest
func (c AVCaptureDevice) CenterStageRectOfInterest() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("centerStageRectOfInterest"))
	return corefoundation.CGRect(rv)
}
func (c AVCaptureDevice) SetCenterStageRectOfInterest(value corefoundation.CGRect) {
	objc.Send[struct{}](c.ID, objc.Sel("setCenterStageRectOfInterest:"), value)
}
//
// # Discussion
// 
// Indicates whether the device supports the Center Stage Rect of Interest
// feature.
// 
// This property returns YES if the device supports Center Stage Rect of
// Interest.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/centerStageRectOfInterestSupported
func (c AVCaptureDevice) CenterStageRectOfInterestSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isCenterStageRectOfInterestSupported"))
	return rv
}
// A video zoom factor multiplier to use when displaying zoom information in a
// user interface.
//
// # Discussion
// 
// Some system user interfaces, like the macOS Video Effects Menu, display a
// video zoom factor value in a way most appropriate for visual presentation,
// which might differ from the [VideoZoomFactor] property value.
// 
// Your app can key-value observe this property to update the display video
// zoom factor values in its user interface to stay consistent with Apple’s
// system UIs.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/displayVideoZoomFactorMultiplier
func (c AVCaptureDevice) DisplayVideoZoomFactorMultiplier() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("displayVideoZoomFactorMultiplier"))
	return rv
}
// The exposure mode for the device.
//
// # Discussion
// 
// Before changing the value of this property, you must call
// [LockForConfiguration] to acquire exclusive access to the device’s
// configuration properties. Otherwise, setting the value of this property
// raises an exception. When you’re done configuring the device, call
// [UnlockForConfiguration] to release the lock and allow other devices to
// configure the settings.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/exposureMode-swift.property
func (c AVCaptureDevice) ExposureMode() AVCaptureExposureMode {
	rv := objc.Send[AVCaptureExposureMode](c.ID, objc.Sel("exposureMode"))
	return AVCaptureExposureMode(rv)
}
func (c AVCaptureDevice) SetExposureMode(value AVCaptureExposureMode) {
	objc.Send[struct{}](c.ID, objc.Sel("setExposureMode:"), value)
}
// The point of interest for exposure.
//
// # Discussion
// 
// Setting a value for this property doesn’t initiate an exposure
// rebalancing operation. To set exposure using a point of interest, first set
// this property’s value, then set the [ExposureMode] property to
// [CaptureExposureModeAutoExpose] or
// [CaptureExposureModeContinuousAutoExposure].
// 
// This property’s [CGPoint] value uses a coordinate system where `{0,0}` is
// the top-left of the picture area and `{1,1}` is the bottom-right. This
// coordinate system is always relative to a landscape device orientation with
// the home button on the right, regardless of the actual device orientation.
// You can convert between this coordinate system and view coordinates using
// [AVCaptureVideoPreviewLayer] methods.
// 
// Before changing the value of this property, you must call
// [LockForConfiguration] to acquire exclusive access to the device’s
// configuration properties. Otherwise, setting the value of this property
// raises an exception. When you’re done configuring the device, call
// [UnlockForConfiguration] to release the lock and allow other devices to
// configure the settings.
// 
// This property is key-value observable.
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/exposurePointOfInterest
func (c AVCaptureDevice) ExposurePointOfInterest() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c.ID, objc.Sel("exposurePointOfInterest"))
	return corefoundation.CGPoint(rv)
}
func (c AVCaptureDevice) SetExposurePointOfInterest(value corefoundation.CGPoint) {
	objc.Send[struct{}](c.ID, objc.Sel("setExposurePointOfInterest:"), value)
}
// A Boolean value that indicates whether the device supports a point of
// interest for exposure.
//
// # Discussion
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isExposurePointOfInterestSupported
func (c AVCaptureDevice) ExposurePointOfInterestSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isExposurePointOfInterestSupported"))
	return rv
}
// The device’s current exposure rectangle of interest, if it has one.
//
// # Discussion
// 
// The value of this property is a [CGRect] determining the device’s
// exposure rectangle of interest. Use this as an alternative to setting
// [ExposurePointOfInterest], as it allows you to specify both a location and
// size. For example, a value of `CGRectMake(0, 0, 1, 1)` tells the device to
// use the entire field of view when determining the exposure, while
// `CGRectMake(0, 0, 0.25, 0.25)` indicates the top left sixteenth, and
// `CGRectMake(0.75, 0.75, 0.25, 0.25)` indicates the bottom right sixteenth.
// Setting [ExposureRectOfInterest] throws an [NSInvalidArgumentException] if
// [ExposureRectOfInterestSupported] returns `false`. Setting
// [ExposureRectOfInterest] throws an [NSInvalidArgumentException] if your
// provided rectangle’s size is smaller than the
// [MinExposureRectOfInterestSize]. Setting [ExposureRectOfInterest] throws an
// [NSGenericException] if you call it without first obtaining exclusive
// access to the device using [LockForConfiguration]. Setting
// [ExposureRectOfInterest] updates the device’s [ExposurePointOfInterest]
// to the center of your provided rectangle of interest. If you later set the
// device’s [ExposurePointOfInterest], the [ExposureRectOfInterest] resets
// to the default sized rectangle of interest for the new exposure point of
// interest. If you change your [ActiveFormat], the point of interest and
// rectangle of interest both revert to their default values. You can observe
// automatic changes to the device’s [ExposureRectOfInterest] by key-value
// observing this property.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/exposureRectOfInterest
func (c AVCaptureDevice) ExposureRectOfInterest() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("exposureRectOfInterest"))
	return corefoundation.CGRect(rv)
}
func (c AVCaptureDevice) SetExposureRectOfInterest(value corefoundation.CGRect) {
	objc.Send[struct{}](c.ID, objc.Sel("setExposureRectOfInterest:"), value)
}
// Whether the device supports exposure rectangles of interest.
//
// # Discussion
// 
// You may only set the device’s [ExposureRectOfInterest] property if this
// property returns `true`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isExposureRectOfInterestSupported
func (c AVCaptureDevice) ExposureRectOfInterestSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isExposureRectOfInterestSupported"))
	return rv
}
// A Boolean value that indicates whether the flash is currently available for
// use.
//
// # Discussion
// 
// The flash may become unavailable if, for example, the device overheats and
// needs to cool off.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isFlashAvailable
func (c AVCaptureDevice) FlashAvailable() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isFlashAvailable"))
	return rv
}
// The device’s current flash mode.
//
// # Discussion
// 
// Before changing the value of this property, you must call
// [LockForConfiguration] to acquire exclusive access to the device’s
// configuration properties. Otherwise, setting the value of this property
// raises an exception. When you finish configuring the device, call
// [UnlockForConfiguration] to release the lock and allow other devices to
// configure the settings.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/flashMode-swift.property
func (c AVCaptureDevice) FlashMode() AVCaptureFlashMode {
	rv := objc.Send[AVCaptureFlashMode](c.ID, objc.Sel("flashMode"))
	return AVCaptureFlashMode(rv)
}
func (c AVCaptureDevice) SetFlashMode(value AVCaptureFlashMode) {
	objc.Send[struct{}](c.ID, objc.Sel("setFlashMode:"), value)
}
// The capture device’s focus mode.
//
// # Discussion
// 
// Before changing the value of this property, you must call
// [LockForConfiguration] to acquire exclusive access to the device’s
// configuration properties. Otherwise, setting the value of this property
// raises an exception. When you finish configuring the device, call
// [UnlockForConfiguration] to release the lock and allow other devices to
// configure the settings.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/focusMode-swift.property
func (c AVCaptureDevice) FocusMode() AVCaptureFocusMode {
	rv := objc.Send[AVCaptureFocusMode](c.ID, objc.Sel("focusMode"))
	return AVCaptureFocusMode(rv)
}
func (c AVCaptureDevice) SetFocusMode(value AVCaptureFocusMode) {
	objc.Send[struct{}](c.ID, objc.Sel("setFocusMode:"), value)
}
// The point of interest for focusing.
//
// # Discussion
// 
// Setting a value for this property doesn’t initiate a focusing operation.
// To focus the camera on a point of interest, first set this property’s
// value, then set the [FocusMode] property to [CaptureFocusModeAutoFocus] or
// [CaptureFocusModeContinuousAutoFocus].
// 
// This property’s [CGPoint] value uses a coordinate system where `{0,0}` is
// the top-left of the picture area and `{1,1}` is the bottom-right. This
// coordinate system is always relative to a landscape device orientation with
// the home button on the right, regardless of the actual device orientation.
// You can convert between this coordinate system and view coordinates using
// [AVCaptureVideoPreviewLayer] methods.
// 
// Before changing the value of this property, you must call
// [LockForConfiguration] to acquire exclusive access to the device’s
// configuration properties. Otherwise, setting the value of this property
// raises an exception. When you finish configuring the device, call
// [UnlockForConfiguration] to release the lock and allow other devices to
// configure the settings.
// 
// This property is key-value observable.
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/focusPointOfInterest
func (c AVCaptureDevice) FocusPointOfInterest() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](c.ID, objc.Sel("focusPointOfInterest"))
	return corefoundation.CGPoint(rv)
}
func (c AVCaptureDevice) SetFocusPointOfInterest(value corefoundation.CGPoint) {
	objc.Send[struct{}](c.ID, objc.Sel("setFocusPointOfInterest:"), value)
}
// A Boolean value that indicates whether the device supports a point of
// interest for focus.
//
// # Discussion
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isFocusPointOfInterestSupported
func (c AVCaptureDevice) FocusPointOfInterestSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isFocusPointOfInterestSupported"))
	return rv
}
// The device’s current focus rectangle of interest, if it has one.
//
// # Discussion
// 
// The value of this property is a [CGRect] determining the device’s focus
// rectangle of interest. Use this as an alternative to setting
// [FocusPointOfInterest], as it allows you to specify both a location and
// size. For example, a value of `CGRectMake(0, 0, 1, 1)` tells the device to
// use the entire field of view when determining the focus, while
// `CGRectMake(0, 0, 0.25, 0.25)` indicates the top left sixteenth, and
// `CGRectMake(0.75, 0.75, 0.25, 0.25)` indicates the bottom right sixteenth.
// Setting [FocusRectOfInterest] throws an [NSInvalidArgumentException] if
// [FocusRectOfInterestSupported] returns `false`. Setting
// [FocusRectOfInterest] throws an [NSInvalidArgumentException] if your
// provided rectangle’s size is smaller than the
// [MinFocusRectOfInterestSize]. Setting [FocusRectOfInterest] throws an
// [NSGenericException] if you call it without first obtaining exclusive
// access to the device using [LockForConfiguration]. Setting
// [FocusRectOfInterest] updates the device’s [FocusPointOfInterest] to the
// center of your provided rectangle of interest. If you later set the
// device’s [FocusPointOfInterest], the [FocusRectOfInterest] resets to the
// default sized rectangle of interest for the new focus point of interest. If
// you change your [ActiveFormat], the point of interest and rectangle of
// interest both revert to their default values. You can observe automatic
// changes to the device’s [FocusRectOfInterest] by key-value observing this
// property.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/focusRectOfInterest
func (c AVCaptureDevice) FocusRectOfInterest() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("focusRectOfInterest"))
	return corefoundation.CGRect(rv)
}
func (c AVCaptureDevice) SetFocusRectOfInterest(value corefoundation.CGRect) {
	objc.Send[struct{}](c.ID, objc.Sel("setFocusRectOfInterest:"), value)
}
// Whether the receiver supports focus rectangles of interest.
//
// # Discussion
// 
// You may only set the device’s [FocusRectOfInterest] property if this
// property returns `true`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isFocusRectOfInterestSupported
func (c AVCaptureDevice) FocusRectOfInterestSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isFocusRectOfInterestSupported"))
	return rv
}
// The capture formats a device supports.
//
// # Discussion
// 
// A capture device format describes the details of the video, image, or audio
// parameters of a specific mode of capture. If you require specifying capture
// settings not covered by a capture session preset, you can set the
// [ActiveFormat] property to any of the formats in this array.
// 
// This property value is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/formats
func (c AVCaptureDevice) Formats() []AVCaptureDeviceFormat {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("formats"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaptureDeviceFormat {
		return AVCaptureDeviceFormatFromID(id)
	})
}
// A Boolean value that indicates whether the capture device has a flash.
//
// # Discussion
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/hasFlash
func (c AVCaptureDevice) HasFlash() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasFlash"))
	return rv
}
// A Boolean value that specifies whether the capture device has a torch.
//
// # Discussion
// 
// A torch is a light source, such as an LED flash, that’s available on the
// device and used for illuminating captured content or providing general
// illumination. This property reflects whether the current device has such
// illumination hardware built-in.
// 
// Even if the device has a torch, that torch might not be available for use,
// so check the value of the [TorchAvailable] property before using it.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/hasTorch
func (c AVCaptureDevice) HasTorch() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("hasTorch"))
	return rv
}
// An array of input sources that the device supports.
//
// # Discussion
// 
// Some devices can capture data from one of multiple data sources (different
// input jacks on the same audio device, for example). For devices with
// multiple possible data sources, you can use this property to enumerate the
// possible choices.
// 
// This value is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/inputSources
func (c AVCaptureDevice) InputSources() []AVCaptureDeviceInputSource {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("inputSources"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaptureDeviceInputSource {
		return AVCaptureDeviceInputSourceFromID(id)
	})
}
// An array of capture devices that are physically linked to a device.
//
// # Discussion
// 
// For an external iSight camera, the array contains an [AVCaptureDevice]
// instance that represents the external iSight microphone.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/linkedDevices
func (c AVCaptureDevice) LinkedDevices() []AVCaptureDevice {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("linkedDevices"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaptureDevice {
		return AVCaptureDeviceFromID(id)
	})
}
// The minimum size you may use when specifying a rectangle of interest.
//
// # Discussion
// 
// The size returned is in normalized coordinates, and depends on the current
// [ActiveFormat]. If [ExposureRectOfInterestSupported] returns `false`, this
// property returns { 0, 0 }.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/minExposureRectOfInterestSize
func (c AVCaptureDevice) MinExposureRectOfInterestSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("minExposureRectOfInterestSize"))
	return corefoundation.CGSize(rv)
}
// The minimum size you may use when specifying a rectangle of interest.
//
// # Discussion
// 
// The size returned is in normalized coordinates, and depends on the current
// [ActiveFormat]. If [FocusRectOfInterestSupported] returns `false`, this
// property returns { 0, 0 }.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/minFocusRectOfInterestSize
func (c AVCaptureDevice) MinFocusRectOfInterestSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](c.ID, objc.Sel("minFocusRectOfInterestSize"))
	return corefoundation.CGSize(rv)
}
// The capture device’s minimum focus distance in millimeters.
//
// # Discussion
// 
// For virtual cameras, like [builtInDualCamera] or [builtInTripleCamera],
// this value represents the smallest minimum focus distance of the
// autofocus-capable cameras that it sources.
// 
// This value is `-1` if the distance is unknown.
//
// [builtInDualCamera]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DeviceType-swift.struct/builtInDualCamera
// [builtInTripleCamera]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DeviceType-swift.struct/builtInTripleCamera
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/minimumFocusDistance
func (c AVCaptureDevice) MinimumFocusDistance() int {
	rv := objc.Send[int](c.ID, objc.Sel("minimumFocusDistance"))
	return rv
}
// A Boolean value that indicates whether the Portrait video effect is active
// on a device.
//
// # Discussion
// 
// When active, the device blurs the background, simulating a shallow depth of
// field effect. The device also limits the values of its
// [ActiveVideoMinFrameDuration] and [ActiveVideoMaxFrameDuration] to the
// value that the device format’s [VideoFrameRateRangeForPortraitEffect]
// specifies.
// 
// When a capture device’s [PortraitEffectEnabled] property value is [true],
// it may also return [true] for this property, depending on whether it
// supports the feature in its current configuration.
// 
// This property is key-value observable.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isPortraitEffectActive
func (c AVCaptureDevice) PortraitEffectActive() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isPortraitEffectActive"))
	return rv
}
// An array of reaction effects that the device is currently performing,
// sorted by timestamp.
//
// # Discussion
// 
// Key-value observe this property to determine when reaction effects begin
// and end. If your key-value observing callback provides old and new values,
// any in-progress reaction effects in the new array have a value of [invalid]
// for their [EndTime] property value. Completed reaction effects are only in
// the old array, and have their [EndTime] property value set to the
// presentation time of the first frame where the reaction effect was no
// longer present.
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/reactionEffectsInProgress
func (c AVCaptureDevice) ReactionEffectsInProgress() []AVCaptureReactionEffectState {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("reactionEffectsInProgress"))
	return objc.ConvertSlice(rv, func(id objc.ID) AVCaptureReactionEffectState {
		return AVCaptureReactionEffectStateFromID(id)
	})
}
// A Boolean value that indicates whether Studio Light is active on a device.
//
// # Discussion
// 
// When the value is [true], the system artificially lights the subject’s
// face to simulate the presence of a studio light near the camera.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isStudioLightActive
func (c AVCaptureDevice) StudioLightActive() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isStudioLightActive"))
	return rv
}
// A Boolean value that indicates whether the device’s torch is currently
// active.
//
// # Discussion
// 
// A torch must be present on the device and currently available before it can
// be active.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isTorchActive
func (c AVCaptureDevice) TorchActive() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isTorchActive"))
	return rv
}
// A Boolean value that indicates whether the torch is currently available for
// use.
//
// # Discussion
// 
// The torch may become unavailable if, for example, the device overheats and
// needs to cool off.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isTorchAvailable
func (c AVCaptureDevice) TorchAvailable() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isTorchAvailable"))
	return rv
}
// The current torch brightness level.
//
// # Discussion
// 
// The value of this property is a floating-point number whose value is in the
// range 0.0 to 1.0. A torch level of 0.0 indicates that the torch is off. A
// torch level of 1.0 represents the theoretical maximum value, although the
// actual maximum value may be lower if the device is currently overheated.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/torchLevel
func (c AVCaptureDevice) TorchLevel() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("torchLevel"))
	return rv
}
// The current torch mode.
//
// # Discussion
// 
// Setting the value of this property also sets the torch level to its maximum
// current value.
// 
// Before setting the value of this property, call the [IsTorchModeSupported]
// method to make sure the device supports the desired mode. Setting the
// device to an unsupported torch mode results in the raising of an exception.
// For a list of possible values for this property, see
// [AVCaptureDevice.TorchMode].
// 
// Before changing the value of this property, you must call
// [LockForConfiguration] to acquire exclusive access to the device’s
// configuration properties. Otherwise, setting the value of this property
// raises an exception. When you finish configuring the device, call
// [UnlockForConfiguration] to release the lock and allow other devices to
// configure the settings.
// 
// This property is key-value observable.
//
// [AVCaptureDevice.TorchMode]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/TorchMode-swift.enum
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/torchMode-swift.property
func (c AVCaptureDevice) TorchMode() AVCaptureTorchMode {
	rv := objc.Send[AVCaptureTorchMode](c.ID, objc.Sel("torchMode"))
	return AVCaptureTorchMode(rv)
}
func (c AVCaptureDevice) SetTorchMode(value AVCaptureTorchMode) {
	objc.Send[struct{}](c.ID, objc.Sel("setTorchMode:"), value)
}
// The current playback mode.
//
// # Discussion
// 
// For devices that support transport control, query this property to discover
// the current playback mode.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/transportControlsPlaybackMode-swift.property
func (c AVCaptureDevice) TransportControlsPlaybackMode() AVCaptureDeviceTransportControlsPlaybackMode {
	rv := objc.Send[AVCaptureDeviceTransportControlsPlaybackMode](c.ID, objc.Sel("transportControlsPlaybackMode"))
	return AVCaptureDeviceTransportControlsPlaybackMode(rv)
}
// The current playback speed.
//
// # Discussion
// 
// For devices that support transport control, the value of this property
// indicates the current playback speed of the deck. The following table gives
// examples of the meaning of values:
// 
// [Table data omitted]
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/transportControlsSpeed-swift.property
func (c AVCaptureDevice) TransportControlsSpeed() AVCaptureDeviceTransportControlsSpeed {
	rv := objc.Send[AVCaptureDeviceTransportControlsSpeed](c.ID, objc.Sel("transportControlsSpeed"))
	return AVCaptureDeviceTransportControlsSpeed(rv)
}
// A Boolean value that indicates whether the device supports transport
// control commands.
//
// # Discussion
// 
// For devices with transport controls, such as AVC tape-based camcorders or
// pro capture devices with RS422 deck control, the value of this property is
// [true]. If transport controls aren’t supported, none of the associated
// transport control methods and properties are available to the device.
// 
// This property is key-value observable.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/transportControlsSupported
func (c AVCaptureDevice) TransportControlsSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("transportControlsSupported"))
	return rv
}
// The current white balance mode.
//
// # Discussion
// 
// Before changing the value of this property, you must call
// [LockForConfiguration] to acquire exclusive access to the device’s
// configuration properties. Otherwise, setting the value of this property
// raises an exception. When you’re done configuring the device, call
// [UnlockForConfiguration] to release the lock and allow other devices to
// configure the settings.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/whiteBalanceMode-swift.property
func (c AVCaptureDevice) WhiteBalanceMode() AVCaptureWhiteBalanceMode {
	rv := objc.Send[AVCaptureWhiteBalanceMode](c.ID, objc.Sel("whiteBalanceMode"))
	return AVCaptureWhiteBalanceMode(rv)
}
func (c AVCaptureDevice) SetWhiteBalanceMode(value AVCaptureWhiteBalanceMode) {
	objc.Send[struct{}](c.ID, objc.Sel("setWhiteBalanceMode:"), value)
}

// A camera the system prefers to use for video and photo capture.
//
// # Discussion
// 
// The system chooses the value of this property. It considers the value of
// [UserPreferredCamera], as well as other factors like camera suspension and
// the appearance of Continuity Cameras that apps should choose automatically.
// The property may change spontaneously, such as when the preferred camera
// goes away.
// 
// Apps that adopt this API should always key-value observe this property and
// update their capture session’s input device to reflect changes to this
// value. An app can still offer users the ability to pick a camera by setting
// a [UserPreferredCamera] value. Doing so puts the user’s choice first
// until either another system-preferred device becomes available or the user
// reboots the machine (after which it reverts to its original behavior of
// returning the internally-determined best camera to use).
// 
// If you want to offer users a fully manual camera selection mode in addition
// to automatic camera selection, it’s recommended to set the
// [UserPreferredCamera] value each time the user makes a camera selection,
// but ignore key-value observer updates to this property value while in
// manual selection mode.
// 
// This property always returns a device that’s present. If no camera is
// available, this value is `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/systemPreferredCamera
func (_AVCaptureDeviceClass AVCaptureDeviceClass) SystemPreferredCamera() AVCaptureDevice {
	rv := objc.Send[objc.ID](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("systemPreferredCamera"))
	return AVCaptureDeviceFromID(objc.ID(rv))
}
// A camera the user prefers to use for video and photo capture.
//
// # Discussion
// 
// In addition to being a [SystemPreferredCamera], you can designate a device
// as a user-preferred camera. Setting a value for this property allows an app
// to persist its preference across app launches and system reboots. The
// system internally maintains a short history of devices, so if a user’s
// most recently preferred camera isn’t currently connected, it still
// reports the next best choice.
// 
// This property always returns a device that’s present. If no camera is
// available, this value is `nil`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/userPreferredCamera
func (_AVCaptureDeviceClass AVCaptureDeviceClass) UserPreferredCamera() AVCaptureDevice {
	rv := objc.Send[objc.ID](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("userPreferredCamera"))
	return AVCaptureDeviceFromID(objc.ID(rv))
}
func (_AVCaptureDeviceClass AVCaptureDeviceClass) SetUserPreferredCamera(value AVCaptureDevice) {
	objc.Send[struct{}](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("setUserPreferredCamera:"), value)
}
// A class property indicating whether the edge light UI is actively being
// shown on a screen.
//
// # Discussion
// 
// This readonly property reflects whether the edge light UI is actively being
// shown on a screen. It is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isEdgeLightActive
func (_AVCaptureDeviceClass AVCaptureDeviceClass) EdgeLightActive() bool {
	rv := objc.Send[bool](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("isEdgeLightActive"))
	return rv
}
// A class property indicating whether the Edge Light feature is currently
// enabled in Control Center.
//
// # Discussion
// 
// This readonly property changes to reflect the Edge Light state in Control
// Center. It is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isEdgeLightEnabled
func (_AVCaptureDeviceClass AVCaptureDeviceClass) EdgeLightEnabled() bool {
	rv := objc.Send[bool](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("isEdgeLightEnabled"))
	return rv
}
// The device’s active microphone mode.
//
// # Discussion
// 
// The value may differ from the value of the [PreferredMicrophoneMode]
// property if the app’s active audio route doesn’t support the mode.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/activeMicrophoneMode
func (_AVCaptureDeviceClass AVCaptureDeviceClass) ActiveMicrophoneMode() AVCaptureMicrophoneMode {
	rv := objc.Send[AVCaptureMicrophoneMode](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("activeMicrophoneMode"))
	return AVCaptureMicrophoneMode(rv)
}
// A class property that indicates whether a person enables the Background
// Replacement feature for this app.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isBackgroundReplacementEnabled
func (_AVCaptureDeviceClass AVCaptureDeviceClass) BackgroundReplacementEnabled() bool {
	rv := objc.Send[bool](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("isBackgroundReplacementEnabled"))
	return rv
}
// A value that indicates the current mode of Center Stage control.
//
// # Discussion
// 
// See Control Modes for details on choosing an appropriate control mode.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/centerStageControlMode-swift.type.property
func (_AVCaptureDeviceClass AVCaptureDeviceClass) CenterStageControlMode() AVCaptureCenterStageControlMode {
	rv := objc.Send[AVCaptureCenterStageControlMode](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("centerStageControlMode"))
	return AVCaptureCenterStageControlMode(rv)
}
func (_AVCaptureDeviceClass AVCaptureDeviceClass) SetCenterStageControlMode(value AVCaptureCenterStageControlMode) {
	objc.Send[struct{}](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("setCenterStageControlMode:"), value)
}
// A Boolean value that indicates whether a user or an app enabled Center
// Stage on a device.
//
// # Discussion
// 
// You can only set this value when Center Stage is under app or cooperative
// control. Attempting to change the enabled state when the control mode is
// [CaptureCenterStageControlModeUser], results in the system throwing an
// exception.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isCenterStageEnabled
func (_AVCaptureDeviceClass AVCaptureDeviceClass) CenterStageEnabled() bool {
	rv := objc.Send[bool](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("isCenterStageEnabled"))
	return rv
}
func (_AVCaptureDeviceClass AVCaptureDeviceClass) SetCenterStageEnabled(value bool) {
	objc.Send[struct{}](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("setCenterStageEnabled:"), value)
}
// A special constant that represents the current exposure bias value.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/currentexposuretargetbias
func (_AVCaptureDeviceClass AVCaptureDeviceClass) CurrentExposureTargetBias() float32 {
	rv := objc.Send[float32](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("AVCaptureExposureTargetBiasCurrent"))
	return rv
}
// A constant that represents the current lens position.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/currentlensposition
func (_AVCaptureDeviceClass AVCaptureDeviceClass) CurrentLensPosition() float32 {
	rv := objc.Send[float32](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("AVCaptureLensPositionCurrent"))
	return rv
}
// A constant that indicates to set the torch to its maximum level.
//
// See: https://developer.apple.com/documentation/avfoundation/avcapturedevice/maxavailabletorchlevel
func (_AVCaptureDeviceClass AVCaptureDeviceClass) MaxAvailableTorchLevel() float32 {
	rv := objc.Send[float32](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("AVCaptureMaxAvailableTorchLevel"))
	return rv
}
// A Boolean value that indicates whether the user enabled the Portrait video
// effect in Control Center.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isPortraitEffectEnabled
func (_AVCaptureDeviceClass AVCaptureDeviceClass) PortraitEffectEnabled() bool {
	rv := objc.Send[bool](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("isPortraitEffectEnabled"))
	return rv
}
// The microphone mode that the user selects in Control Center.
//
// # Discussion
// 
// Use key-value observing to monitor the user’s microphone mode selection.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/preferredMicrophoneMode
func (_AVCaptureDeviceClass AVCaptureDeviceClass) PreferredMicrophoneMode() AVCaptureMicrophoneMode {
	rv := objc.Send[AVCaptureMicrophoneMode](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("preferredMicrophoneMode"))
	return AVCaptureMicrophoneMode(rv)
}
// A Boolean value that indicates whether gesture detection triggers reaction
// effects on the video stream.
//
// # Discussion
// 
// This property reflects the enabled state of Gestures in Control Center.
// 
// Gesture detection runs only when the device’s active format supports
// reaction effects, which you determine by querying the value of the
// format’s [ReactionEffectsSupported] property.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/reactionEffectGesturesEnabled
func (_AVCaptureDeviceClass AVCaptureDeviceClass) ReactionEffectGesturesEnabled() bool {
	rv := objc.Send[bool](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("reactionEffectGesturesEnabled"))
	return rv
}
// A Boolean value that indicates whether the app supports performing reaction
// effects.
//
// # Discussion
// 
// The system only renders reaction effects when the device’s active format
// supports the feature, which you determine by querying the value of its
// [ReactionEffectsSupported] property.
// 
// In macOS, the system enables reaction effects for all apps by default. In
// iOS, it enables them by default only for video conferencing apps (apps that
// enable the Voice over IP option in their [UIBackgroundModes]
// configuration). Apps that don’t use this background mode may opt in to
// this feature by adding the following key to the `Info.Plist()` file.
//
// [UIBackgroundModes]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/UIBackgroundModes
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/reactionEffectsEnabled
func (_AVCaptureDeviceClass AVCaptureDeviceClass) ReactionEffectsEnabled() bool {
	rv := objc.Send[bool](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("reactionEffectsEnabled"))
	return rv
}
// A Boolean value that indicates whether a user enabled Studio Light on a
// device.
//
// # Discussion
// 
// When the value is [true], the system artificially lights the subject’s
// face to simulate the presence of a studio light near the camera.
// 
// This property is key-value observable.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/isStudioLightEnabled
func (_AVCaptureDeviceClass AVCaptureDeviceClass) StudioLightEnabled() bool {
	rv := objc.Send[bool](objc.ID(_AVCaptureDeviceClass.class), objc.Sel("isStudioLightEnabled"))
	return rv
}

// RequestAccessForMediaType is a synchronous wrapper around [AVCaptureDevice.RequestAccessForMediaTypeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (cc AVCaptureDeviceClass) RequestAccessForMediaType(ctx context.Context, mediaType AVMediaType) (bool, error) {
	done := make(chan bool, 1)
	cc.RequestAccessForMediaTypeCompletionHandler(mediaType, func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

