// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [AVCaptureDeviceInput] class.
var (
	_AVCaptureDeviceInputClass     AVCaptureDeviceInputClass
	_AVCaptureDeviceInputClassOnce sync.Once
)

func getAVCaptureDeviceInputClass() AVCaptureDeviceInputClass {
	_AVCaptureDeviceInputClassOnce.Do(func() {
		_AVCaptureDeviceInputClass = AVCaptureDeviceInputClass{class: objc.GetClass("AVCaptureDeviceInput")}
	})
	return _AVCaptureDeviceInputClass
}

// GetAVCaptureDeviceInputClass returns the class object for AVCaptureDeviceInput.
func GetAVCaptureDeviceInputClass() AVCaptureDeviceInputClass {
	return getAVCaptureDeviceInputClass()
}

type AVCaptureDeviceInputClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (ac AVCaptureDeviceInputClass) Class() objc.Class {
	return ac.class
}

// Alloc allocates memory for a new instance of the class.
func (ac AVCaptureDeviceInputClass) Alloc() AVCaptureDeviceInput {
	rv := objc.Send[AVCaptureDeviceInput](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// An object that provides media input from a capture device to a capture
// session.
//
// # Overview
// 
// This class is a concrete subclass of [AVCaptureInput] that you use to
// connect a capture device to a capture session.
//
// # Creating an input
//
//   - [AVCaptureDeviceInput.InitWithDeviceError]: Creates an input for the specified capture device.
//
// # Configuring audio properties
//
//   - [AVCaptureDeviceInput.IsMultichannelAudioModeSupported]: A Boolean value that indicates whether the input supports the specified multichannel audio mode.
//   - [AVCaptureDeviceInput.MultichannelAudioMode]: The multichannel audio mode to apply when recording audio.
//   - [AVCaptureDeviceInput.SetMultichannelAudioMode]
//   - [AVCaptureDeviceInput.WindNoiseRemovalSupported]
//   - [AVCaptureDeviceInput.WindNoiseRemovalEnabled]
//   - [AVCaptureDeviceInput.SetWindNoiseRemovalEnabled]
//
// # Configuring Cinematic video capture
//
//   - [AVCaptureDeviceInput.CinematicVideoCaptureSupported]: A BOOL value specifying whether Cinematic Video capture is supported.
//   - [AVCaptureDeviceInput.CinematicVideoCaptureEnabled]: A BOOL value specifying whether the Cinematic Video effect is being applied to any movie file output, video data output, metadata output, or video preview layer added to the capture session.
//   - [AVCaptureDeviceInput.SetCinematicVideoCaptureEnabled]
//   - [AVCaptureDeviceInput.SimulatedAperture]: Shallow depth of field simulated aperture.
//   - [AVCaptureDeviceInput.SetSimulatedAperture]
//
// # Locking frame duration
//
//   - [AVCaptureDeviceInput.ActiveLockedVideoFrameDuration]: The receiver’s locked frame duration (the reciprocal of its frame rate). Setting this property guarantees the intra-frame duration delivered by the device input is precisely the frame duration you request.
//   - [AVCaptureDeviceInput.SetActiveLockedVideoFrameDuration]
//   - [AVCaptureDeviceInput.LockedVideoFrameDurationSupported]: Indicates whether the device input supports locked frame durations.
//
// # Synchronizing with external devices
//
//   - [AVCaptureDeviceInput.ExternalSyncSupported]: Indicates whether the device input supports being configured to follow an external sync device.
//   - [AVCaptureDeviceInput.FollowExternalSyncDeviceVideoFrameDurationDelegate]: Configures the the device input to follow an external sync device at the given frame duration.
//   - [AVCaptureDeviceInput.UnfollowExternalSyncDevice]: Discontinues external sync.
//   - [AVCaptureDeviceInput.ActiveExternalSyncVideoFrameDuration]: The receiver’s external sync frame duration (the reciprocal of its frame rate) when being driven by an external sync device.
//   - [AVCaptureDeviceInput.ExternalSyncDevice]: The external sync device currently being followed by this input.
//
// # Accessing the device
//
//   - [AVCaptureDeviceInput.Device]: A capture device associated with this input.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput
type AVCaptureDeviceInput struct {
	AVCaptureInput
}

// AVCaptureDeviceInputFromID constructs a [AVCaptureDeviceInput] from an objc.ID.
//
// An object that provides media input from a capture device to a capture
// session.
func AVCaptureDeviceInputFromID(id objc.ID) AVCaptureDeviceInput {
	return AVCaptureDeviceInput{AVCaptureInput: AVCaptureInputFromID(id)}
}
// NOTE: AVCaptureDeviceInput adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [AVCaptureDeviceInput] class.
//
// # Creating an input
//
//   - [IAVCaptureDeviceInput.InitWithDeviceError]: Creates an input for the specified capture device.
//
// # Configuring audio properties
//
//   - [IAVCaptureDeviceInput.IsMultichannelAudioModeSupported]: A Boolean value that indicates whether the input supports the specified multichannel audio mode.
//   - [IAVCaptureDeviceInput.MultichannelAudioMode]: The multichannel audio mode to apply when recording audio.
//   - [IAVCaptureDeviceInput.SetMultichannelAudioMode]
//   - [IAVCaptureDeviceInput.WindNoiseRemovalSupported]
//   - [IAVCaptureDeviceInput.WindNoiseRemovalEnabled]
//   - [IAVCaptureDeviceInput.SetWindNoiseRemovalEnabled]
//
// # Configuring Cinematic video capture
//
//   - [IAVCaptureDeviceInput.CinematicVideoCaptureSupported]: A BOOL value specifying whether Cinematic Video capture is supported.
//   - [IAVCaptureDeviceInput.CinematicVideoCaptureEnabled]: A BOOL value specifying whether the Cinematic Video effect is being applied to any movie file output, video data output, metadata output, or video preview layer added to the capture session.
//   - [IAVCaptureDeviceInput.SetCinematicVideoCaptureEnabled]
//   - [IAVCaptureDeviceInput.SimulatedAperture]: Shallow depth of field simulated aperture.
//   - [IAVCaptureDeviceInput.SetSimulatedAperture]
//
// # Locking frame duration
//
//   - [IAVCaptureDeviceInput.ActiveLockedVideoFrameDuration]: The receiver’s locked frame duration (the reciprocal of its frame rate). Setting this property guarantees the intra-frame duration delivered by the device input is precisely the frame duration you request.
//   - [IAVCaptureDeviceInput.SetActiveLockedVideoFrameDuration]
//   - [IAVCaptureDeviceInput.LockedVideoFrameDurationSupported]: Indicates whether the device input supports locked frame durations.
//
// # Synchronizing with external devices
//
//   - [IAVCaptureDeviceInput.ExternalSyncSupported]: Indicates whether the device input supports being configured to follow an external sync device.
//   - [IAVCaptureDeviceInput.FollowExternalSyncDeviceVideoFrameDurationDelegate]: Configures the the device input to follow an external sync device at the given frame duration.
//   - [IAVCaptureDeviceInput.UnfollowExternalSyncDevice]: Discontinues external sync.
//   - [IAVCaptureDeviceInput.ActiveExternalSyncVideoFrameDuration]: The receiver’s external sync frame duration (the reciprocal of its frame rate) when being driven by an external sync device.
//   - [IAVCaptureDeviceInput.ExternalSyncDevice]: The external sync device currently being followed by this input.
//
// # Accessing the device
//
//   - [IAVCaptureDeviceInput.Device]: A capture device associated with this input.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput
type IAVCaptureDeviceInput interface {
	IAVCaptureInput

	// Topic: Creating an input

	// Creates an input for the specified capture device.
	InitWithDeviceError(device IAVCaptureDevice) (AVCaptureDeviceInput, error)

	// Topic: Configuring audio properties

	// A Boolean value that indicates whether the input supports the specified multichannel audio mode.
	IsMultichannelAudioModeSupported(multichannelAudioMode AVCaptureMultichannelAudioMode) bool
	// The multichannel audio mode to apply when recording audio.
	MultichannelAudioMode() AVCaptureMultichannelAudioMode
	SetMultichannelAudioMode(value AVCaptureMultichannelAudioMode)
	WindNoiseRemovalSupported() bool
	WindNoiseRemovalEnabled() bool
	SetWindNoiseRemovalEnabled(value bool)

	// Topic: Configuring Cinematic video capture

	// A BOOL value specifying whether Cinematic Video capture is supported.
	CinematicVideoCaptureSupported() bool
	// A BOOL value specifying whether the Cinematic Video effect is being applied to any movie file output, video data output, metadata output, or video preview layer added to the capture session.
	CinematicVideoCaptureEnabled() bool
	SetCinematicVideoCaptureEnabled(value bool)
	// Shallow depth of field simulated aperture.
	SimulatedAperture() float32
	SetSimulatedAperture(value float32)

	// Topic: Locking frame duration

	// The receiver’s locked frame duration (the reciprocal of its frame rate). Setting this property guarantees the intra-frame duration delivered by the device input is precisely the frame duration you request.
	ActiveLockedVideoFrameDuration() coremedia.CMTime
	SetActiveLockedVideoFrameDuration(value coremedia.CMTime)
	// Indicates whether the device input supports locked frame durations.
	LockedVideoFrameDurationSupported() bool

	// Topic: Synchronizing with external devices

	// Indicates whether the device input supports being configured to follow an external sync device.
	ExternalSyncSupported() bool
	// Configures the the device input to follow an external sync device at the given frame duration.
	FollowExternalSyncDeviceVideoFrameDurationDelegate(externalSyncDevice IAVExternalSyncDevice, frameDuration coremedia.CMTime, delegate AVExternalSyncDeviceDelegate)
	// Discontinues external sync.
	UnfollowExternalSyncDevice()
	// The receiver’s external sync frame duration (the reciprocal of its frame rate) when being driven by an external sync device.
	ActiveExternalSyncVideoFrameDuration() coremedia.CMTime
	// The external sync device currently being followed by this input.
	ExternalSyncDevice() IAVExternalSyncDevice

	// Topic: Accessing the device

	// A capture device associated with this input.
	Device() IAVCaptureDevice
}

// Init initializes the instance.
func (c AVCaptureDeviceInput) Init() AVCaptureDeviceInput {
	rv := objc.Send[AVCaptureDeviceInput](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c AVCaptureDeviceInput) Autorelease() AVCaptureDeviceInput {
	rv := objc.Send[AVCaptureDeviceInput](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureDeviceInput creates a new AVCaptureDeviceInput instance.
func NewAVCaptureDeviceInput() AVCaptureDeviceInput {
	class := getAVCaptureDeviceInputClass()
	rv := objc.Send[AVCaptureDeviceInput](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates an input for the specified capture device.
//
// device: A device from which to capture media.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/init(device:)
func NewCaptureDeviceInputWithDeviceError(device IAVCaptureDevice) (AVCaptureDeviceInput, error) {
	var errorPtr objc.ID
	instance := getAVCaptureDeviceInputClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDevice:error:"), device, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVCaptureDeviceInput{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVCaptureDeviceInputFromID(rv), nil
}

// Creates an input for the specified capture device.
//
// device: A device from which to capture media.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/init(device:)
func (c AVCaptureDeviceInput) InitWithDeviceError(device IAVCaptureDevice) (AVCaptureDeviceInput, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](c.ID, objc.Sel("initWithDevice:error:"), device, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVCaptureDeviceInput{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVCaptureDeviceInputFromID(rv), nil

}
// A Boolean value that indicates whether the input supports the specified
// multichannel audio mode.
//
// multichannelAudioMode: The multichannel audio mode to test.
//
// # Return Value
// 
// [true] if the input supports the mode; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// You can only set the [MultichannelAudioMode] property if the input supports
// the value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isMultichannelAudioModeSupported(_:)
func (c AVCaptureDeviceInput) IsMultichannelAudioModeSupported(multichannelAudioMode AVCaptureMultichannelAudioMode) bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isMultichannelAudioModeSupported:"), multichannelAudioMode)
	return rv
}
// Configures the the device input to follow an external sync device at the
// given frame duration.
//
// externalSyncDevice: The [AVExternalSyncDevice] hardware to follow.
//
// delegate: The delegate to notify when the connection status changes, or an error
// occurs.
//
// # Discussion
// 
// Call this method to direct your [AVCaptureDeviceInput] to follow the
// external sync pulse from a sync device at the given frame duration.
// 
// Your provided `videoFrameDuration` value must match the sync pulse duration
// of the external sync device. If it does not, the request times out, the
// external sync device’s status returns to
// [AVExternalSyncDeviceStatusReady], and your session stops running, posting
// a [runtimeErrorNotification] with
// [AVErrorFollowExternalSyncDeviceTimedOut].
// 
// The ability to follow an external sync device may change depending on the
// device configuration. For example,
// [FollowExternalSyncDeviceVideoFrameDurationDelegate] cannot be used when
// [AutoVideoFrameRateEnabled] is `true`.
// 
// To stop following an external pulse, call [UnfollowExternalSyncDevice].
// External sync device following is also disabled when your device’s
// [AVCaptureDeviceFormat] changes.
// 
// Your provided delegate’s [ExternalSyncDeviceStatusDidChange] method is
// called with a status of [AVExternalSyncDeviceStatusReady] if the external
// pulse signal is not close enough to the provided `videoFrameDuration` for
// successful calibration.
// 
// Once your [Status] changes to [AVExternalSyncDeviceStatusActiveSync], your
// input’s `AVCaptureInput/activeExternalSyncVideoFrameDuration` property
// reports the up-to-date frame duration.
// `AVCaptureInput/activeExternalSyncVideoFrameDuration` is also reflected in
// the [ActiveVideoMinFrameDuration] and [ActiveVideoMaxFrameDuration] of your
// input’s associated device.
//
// [runtimeErrorNotification]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/runtimeErrorNotification
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/follow(_:videoFrameDuration:delegate:)
func (c AVCaptureDeviceInput) FollowExternalSyncDeviceVideoFrameDurationDelegate(externalSyncDevice IAVExternalSyncDevice, frameDuration coremedia.CMTime, delegate AVExternalSyncDeviceDelegate) {
	objc.Send[objc.ID](c.ID, objc.Sel("followExternalSyncDevice:videoFrameDuration:delegate:"), externalSyncDevice, frameDuration, delegate)
}
// Discontinues external sync.
//
// # Discussion
// 
// This method stops your input from syncing to the external sync device you
// specified in [FollowExternalSyncDeviceVideoFrameDurationDelegate].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/unfollowExternalSyncDevice()
func (c AVCaptureDeviceInput) UnfollowExternalSyncDevice() {
	objc.Send[objc.ID](c.ID, objc.Sel("unfollowExternalSyncDevice"))
}

// Returns a new input for the specified capture device.
//
// device: The device from which to capture input.
//
// outError: If an error occurs during initialization, upon return contains an [NSError]
// object describing the problem.
//
// # Return Value
// 
// A new capture input.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/deviceInputWithDevice:error:
func (_AVCaptureDeviceInputClass AVCaptureDeviceInputClass) DeviceInputWithDeviceError(device IAVCaptureDevice) (AVCaptureDeviceInput, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](objc.ID(_AVCaptureDeviceInputClass.class), objc.Sel("deviceInputWithDevice:error:"), device, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return AVCaptureDeviceInput{}, foundation.NSErrorFrom(errorPtr)
	}
	return AVCaptureDeviceInputFromID(rv), nil

}

// The multichannel audio mode to apply when recording audio.
//
// # Discussion
// 
// This property only takes effect when the system routes audio through the
// built-in microphone. The system ignores the value when an external
// microphone is in use.
// 
// The default value is [CaptureMultichannelAudioModeNone], which indicates to
// use single channel audio recording.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/multichannelAudioMode
func (c AVCaptureDeviceInput) MultichannelAudioMode() AVCaptureMultichannelAudioMode {
	rv := objc.Send[AVCaptureMultichannelAudioMode](c.ID, objc.Sel("multichannelAudioMode"))
	return AVCaptureMultichannelAudioMode(rv)
}
func (c AVCaptureDeviceInput) SetMultichannelAudioMode(value AVCaptureMultichannelAudioMode) {
	objc.Send[struct{}](c.ID, objc.Sel("setMultichannelAudioMode:"), value)
}
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isWindNoiseRemovalSupported
func (c AVCaptureDeviceInput) WindNoiseRemovalSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isWindNoiseRemovalSupported"))
	return rv
}
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isWindNoiseRemovalEnabled
func (c AVCaptureDeviceInput) WindNoiseRemovalEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isWindNoiseRemovalEnabled"))
	return rv
}
func (c AVCaptureDeviceInput) SetWindNoiseRemovalEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setWindNoiseRemovalEnabled:"), value)
}
// A BOOL value specifying whether Cinematic Video capture is supported.
//
// # Discussion
// 
// With Cinematic Video capture, you get a simulated depth-of-field effect
// that keeps your subjects (people, pets, and more) in sharp focus while
// applying a pleasing blur to the background (or foreground). Depending on
// the focus mode (see [AVCaptureDevice.CinematicVideoFocusMode] for detail),
// the camera either uses machine learning to automatically detect and focus
// on subjects in the scene, or it fixes focus on a subject until it exits the
// scene. Cinematic Videos can be played back and edited using the Cinematic
// framework.
// 
// You can adjust the video’s simulated aperture before starting a recording
// using the [SimulatedAperture] property. With Cinematic Video specific focus
// methods on [AVCaptureDevice], you can dynamically control focus
// transitions.
// 
// Movie files captured with Cinematic Video enabled can be played back and
// edited with the [Cinematic framework]
// (https://developer.apple.com/documentation/cinematic/playing-and-editing-cinematic-mode-video?language=objc).
// 
// This property returns `true` if the session’s current configuration
// allows Cinematic Video capture. When switching cameras or formats, this
// property may change. When this property changes from `true` to `false`,
// [CinematicVideoCaptureEnabled] also reverts to `false`. If you’ve
// previously opted in for Cinematic Video capture and then change
// configuration, you may need to set [CinematicVideoCaptureEnabled] to `true`
// again. This property is key-value observable.
//
// [AVCaptureDevice.CinematicVideoFocusMode]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/CinematicVideoFocusMode
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isCinematicVideoCaptureSupported
func (c AVCaptureDeviceInput) CinematicVideoCaptureSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isCinematicVideoCaptureSupported"))
	return rv
}
// A BOOL value specifying whether the Cinematic Video effect is being applied
// to any movie file output, video data output, metadata output, or video
// preview layer added to the capture session.
//
// # Discussion
// 
// Default is `false`. Set to `true` to enable support for Cinematic Video
// capture.
// 
// When you set this property to `true`, your input’s associated [FocusMode]
// changes to [AVCaptureFocusModeContinuousAutoFocus]. While Cinematic Video
// capture is enabled, you are not permitted to change your device’s focus
// mode, and any attempt to do so results in an [NSInvalidArgumentException].
// You may only set this property to `true` if
// [CinematicVideoCaptureSupported] is `true`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isCinematicVideoCaptureEnabled
func (c AVCaptureDeviceInput) CinematicVideoCaptureEnabled() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isCinematicVideoCaptureEnabled"))
	return rv
}
func (c AVCaptureDeviceInput) SetCinematicVideoCaptureEnabled(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setCinematicVideoCaptureEnabled:"), value)
}
// Shallow depth of field simulated aperture.
//
// # Discussion
// 
// When capturing a Cinematic Video, use this property to control the amount
// of blur in the simulated depth of field effect.
// 
// This property only takes effect when [CinematicVideoCaptureEnabled] is set
// to `true`.
// 
// This property is initialized to the associated
// `AVCaptureDevice/activeFormat/defaultSimulatedAperture`.
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/simulatedAperture
func (c AVCaptureDeviceInput) SimulatedAperture() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("simulatedAperture"))
	return rv
}
func (c AVCaptureDeviceInput) SetSimulatedAperture(value float32) {
	objc.Send[struct{}](c.ID, objc.Sel("setSimulatedAperture:"), value)
}
// The receiver’s locked frame duration (the reciprocal of its frame rate).
// Setting this property guarantees the intra-frame duration delivered by the
// device input is precisely the frame duration you request.
//
// # Discussion
// 
// Set this property to run the receiver’s associated [AVCaptureDevice] at
// precisely your provided frame rate (expressed as a duration). Query
// [MinSupportedLockedVideoFrameDuration] to find the minimum value supported
// by this [AVCaptureDeviceInput]. In order to disable locked video frame
// duration, set this property to `kCMTimeInvalid`. This property resets
// itself to `kCMTimeInvalid` when the receiver’s attached [ActiveFormat]
// changes. When you set this property, its value is also reflected in the
// receiver’s [ActiveVideoMinFrameDuration] and
// [ActiveVideoMaxFrameDuration].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/activeLockedVideoFrameDuration
func (c AVCaptureDeviceInput) ActiveLockedVideoFrameDuration() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](c.ID, objc.Sel("activeLockedVideoFrameDuration"))
	return coremedia.CMTime(rv)
}
func (c AVCaptureDeviceInput) SetActiveLockedVideoFrameDuration(value coremedia.CMTime) {
	objc.Send[struct{}](c.ID, objc.Sel("setActiveLockedVideoFrameDuration:"), value)
}
// Indicates whether the device input supports locked frame durations.
//
// # Discussion
// 
// See [ActiveLockedVideoFrameDuration] for more information on video frame
// duration locking.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isLockedVideoFrameDurationSupported
func (c AVCaptureDeviceInput) LockedVideoFrameDurationSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isLockedVideoFrameDurationSupported"))
	return rv
}
// Indicates whether the device input supports being configured to follow an
// external sync device.
//
// # Discussion
// 
// See [FollowExternalSyncDeviceVideoFrameDurationDelegate] for more
// information on external sync.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/isExternalSyncSupported
func (c AVCaptureDeviceInput) ExternalSyncSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isExternalSyncSupported"))
	return rv
}
// The receiver’s external sync frame duration (the reciprocal of its frame
// rate) when being driven by an external sync device.
//
// # Discussion
// 
// Set up your input to follow an external sync device by calling
// [FollowExternalSyncDeviceVideoFrameDurationDelegate].
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/activeExternalSyncVideoFrameDuration
func (c AVCaptureDeviceInput) ActiveExternalSyncVideoFrameDuration() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](c.ID, objc.Sel("activeExternalSyncVideoFrameDuration"))
	return coremedia.CMTime(rv)
}
// The external sync device currently being followed by this input.
//
// # Discussion
// 
// This readonly property returns the [AVExternalSyncDevice] instance you
// provided in [FollowExternalSyncDeviceVideoFrameDurationDelegate]. This
// property returns `nil` when an external sync device is disconnected or
// fails to calibrate.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/externalSyncDevice
func (c AVCaptureDeviceInput) ExternalSyncDevice() IAVExternalSyncDevice {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("externalSyncDevice"))
	return AVExternalSyncDeviceFromID(objc.ID(rv))
}
// A capture device associated with this input.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/device
func (c AVCaptureDeviceInput) Device() IAVCaptureDevice {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("device"))
	return AVCaptureDeviceFromID(objc.ID(rv))
}

