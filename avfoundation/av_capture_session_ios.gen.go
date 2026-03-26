// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
// +build ios

package avfoundation

import (
"github.com/tmc/apple/objc"
)

// A Boolean value that indicates whether the capture session is in an
// interrupted state.
//
// # Discussion
// 
// This property is key-value observable.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/isInterrupted
func (c AVCaptureSession) Interrupted() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isInterrupted"))
		return rv
}
// A Boolean value that indicates whether the capture session supports using
// the camera while multitasking.
//
// # Discussion
// 
// Query this property to determine whether you can use the camera while
// multitasking by setting the state of the [MultitaskingCameraAccessEnabled]
// property to `true`.
// 
// In iOS and iPadOS, this property is `true` for any of the following cases:
// 
// - The app runs on an iPad that supports Stage Manager with an extended
// display. - The app links against iOS 18 or later and uses `voip` as one of
// its [UIBackgroundModes]. - The app has the
// [com.apple.developer.avfoundation.multitasking-camera-access] entitlement.
// 
// In tvOS, this property is always `true`.
// 
// To learn about best practices for using the camera while multitasking, see
// [Accessing the camera while multitasking on iPad].
//
// [Accessing the camera while multitasking on iPad]: https://developer.apple.com/documentation/AVKit/accessing-the-camera-while-multitasking-on-ipad
// [UIBackgroundModes]: https://developer.apple.com/documentation/BundleResources/Information-Property-List/UIBackgroundModes
// [com.apple.developer.avfoundation.multitasking-camera-access]: https://developer.apple.com/documentation/BundleResources/Entitlements/com.apple.developer.avfoundation.multitasking-camera-access
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/isMultitaskingCameraAccessSupported
func (c AVCaptureSession) MultitaskingCameraAccessSupported() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isMultitaskingCameraAccessSupported"))
		return rv
}
// A Boolean value that indicates whether the capture session enables access
// to the camera while multitasking.
//
// # Discussion
// 
// The default value is [false].
// 
// If the value of the [MultitaskingCameraAccessSupported] property is [true],
// you can enable multitasking camera access by setting this value to [true]
// prior to starting the capture session.
// 
// This property is key-value observable.
// 
// To learn about best practices for using the camera while multitasking, see
// [Accessing the camera while multitasking on iPad].
//
// [Accessing the camera while multitasking on iPad]: https://developer.apple.com/documentation/AVKit/accessing-the-camera-while-multitasking-on-ipad
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/isMultitaskingCameraAccessEnabled
func (c AVCaptureSession) MultitaskingCameraAccessEnabled() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isMultitaskingCameraAccessEnabled"))
		return rv
}
func (c AVCaptureSession) SetMultitaskingCameraAccessEnabled(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setMultitaskingCameraAccessEnabled:"), value)
}
// A value that indicates the percentage of the session’s available hardware
// budget in use.
//
// # Discussion
// 
// This property provides a floating-point value from `0.0` to `1.0` that
// indicates the percentage of a capture session’s hardware that it
// currently uses. When the value is greater than `1.0`, the capture session
// can’t run in its current configuration due to hardware constraints.
// Attempting to start the session while it’s in this state results in a
// runtime error.
// 
// Factors that contribute to the hardware cost include:
// 
// - The active formats of the source devices. Some formats use the full
// sensor (4:3) and others a crop (16:9). Cropped formats require lower
// hardware bandwidth, and therefore lower the cost. - The maximum frame rate
// of the source devices’ active formats. The hardware cost increases when
// using high frame rates. - Whether the source devices uses binned formats.
// Binned formats require substantially less hardware bandwidth, and therefore
// result in a lower cost. - The number of sources configured to deliver
// streaming disparity and depth using [AVCaptureDepthDataOutput]. The higher
// the number of cameras configured to produce depth, the higher the cost.
// 
// To reduce the hardware cost, consider picking a sensor-cropped or binned
// [ActiveFormat]. You may also use a device input’s
// [VideoMinFrameDurationOverride] property to artificially limit the maximum
// frame rate of a source device to a lower value. By doing so, you only pay
// the hardware cost for the maximum frame rate you intend to use.
//
// [AVCaptureDepthDataOutput]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDepthDataOutput
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/hardwareCost
func (c AVCaptureSession) HardwareCost() float32 {
rv := objc.Send[float32](c.ID, objc.Sel("hardwareCost"))
		return rv
}
// A Boolean value that indicates whether the capture session uses the app’s
// shared audio session.
//
// # Discussion
// 
// The default value is [true]. If you set the value to [false], the capture
// session uses a private [AVAudioSession] instance for audio recording, which
// may cause interruptions if your app uses its own audio session for
// playback.
//
// [AVAudioSession]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/usesApplicationAudioSession
func (c AVCaptureSession) UsesApplicationAudioSession() bool {
rv := objc.Send[bool](c.ID, objc.Sel("usesApplicationAudioSession"))
		return rv
}
func (c AVCaptureSession) SetUsesApplicationAudioSession(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setUsesApplicationAudioSession:"), value)
}
// A Boolean value that indicates whether the capture session automatically
// changes settings in the app’s shared audio session.
//
// # Discussion
// 
// This property only takes effect if the value of the
// [UsesApplicationAudioSession] property is [true].
// 
// The value of this property defaults to [true], causing the capture session
// to automatically configure the app’s shared [AVAudioSession] instance for
// optimal recording. For example, if the capture session uses a device’s
// rear-facing camera, the system sets the audio session’s microphone and
// polar pattern for optimal recording of sound from that direction. The audio
// session’s original state isn’t restored after capture finishes.
// 
// If you set value to [false], your app is responsible for selecting
// appropriate audio session settings. Recording may fail if the audio
// session’s settings are incompatible with the capture session.
//
// [AVAudioSession]: https://developer.apple.com/documentation/AVFAudio/AVAudioSession
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/automaticallyConfiguresApplicationAudioSession
func (c AVCaptureSession) AutomaticallyConfiguresApplicationAudioSession() bool {
rv := objc.Send[bool](c.ID, objc.Sel("automaticallyConfiguresApplicationAudioSession"))
		return rv
}
func (c AVCaptureSession) SetAutomaticallyConfiguresApplicationAudioSession(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setAutomaticallyConfiguresApplicationAudioSession:"), value)
}
// A Boolean value that Indicates whether the capture session configures the
// app’s audio session to mix with others.
//
// # Discussion
// 
// By default, a capture session’s audio session interrupts the audio of
// other apps. To enable background audio from other apps to continue while
// capturing video, set this value to [true].
// 
// The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/configuresApplicationAudioSessionToMixWithOthers
func (c AVCaptureSession) ConfiguresApplicationAudioSessionToMixWithOthers() bool {
rv := objc.Send[bool](c.ID, objc.Sel("configuresApplicationAudioSessionToMixWithOthers"))
		return rv
}
func (c AVCaptureSession) SetConfiguresApplicationAudioSessionToMixWithOthers(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setConfiguresApplicationAudioSessionToMixWithOthers:"), value)
}
// A Boolean value that indicates whether the capture session configures the
// app’s audio session for bluetooth high-quality recording.
//
// # Discussion
// 
// Use this property to enable using AirPods as a high-quality microphone. Set
// this value to `true` to tell a capture session to opt-in to high-quality
// bluetooth recording, which enables a person to select AirPods as the active
// mic source for capture. This property has no effect when the value of
// [UsesApplicationAudioSession] is `false`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/configuresApplicationAudioSessionForBluetoothHighQualityRecording
func (c AVCaptureSession) ConfiguresApplicationAudioSessionForBluetoothHighQualityRecording() bool {
rv := objc.Send[bool](c.ID, objc.Sel("configuresApplicationAudioSessionForBluetoothHighQualityRecording"))
		return rv
}
func (c AVCaptureSession) SetConfiguresApplicationAudioSessionForBluetoothHighQualityRecording(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setConfiguresApplicationAudioSessionForBluetoothHighQualityRecording:"), value)
}
// A Boolean value that specifies whether the session should automatically use
// wide-gamut color where available.
//
// # Discussion
// 
// All devices and formats support capture in the sRGB color space. Some
// devices and formats can also capture in the P3 color space, which includes
// a much wider gamut of colors. Wide-gamut capture is appropriate for only
// certain capture workflows, so this property controls automatic
// configuration for those workflows.
// 
// When this property is [true] (the default), and your session configuration
// is appropriate for wide-gamut capture:
// 
// - If you use a session preset other than [inputPriority], the session
// automatically sets the device’s [ActiveFormat] property to one that
// supports wide-gamut capture, and sets the device’s [ActiveColorSpace]
// property to a wide-gamut color space. - If you manually choose a capture
// format (thereby setting the session to input priority), the session
// automatically sets the device’s [ActiveColorSpace] property to a
// wide-gamut color space only if your chosen format supports wide-gamut
// capture.
// 
// For information about which devices and session configurations
// automatically enable wide-gamut capture, see Wide-Gamut Capture in [iOS
// Device Compatibility Reference].
// 
// Set this property to [false] if you want to directly change the value of
// the capture device’s [ActiveColorSpace] property (regardless of whether
// you configure your device with a session preset or by directly setting a
// capture format).
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [iOS Device Compatibility Reference]: https://developer.apple.com/library/archive/documentation/DeviceInformation/Reference/iOSDeviceCompatibility/Introduction/Introduction.html#//apple_ref/doc/uid/TP40013599
// [inputPriority]: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/Preset/inputPriority
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureSession/automaticallyConfiguresCaptureDeviceForWideColor
func (c AVCaptureSession) AutomaticallyConfiguresCaptureDeviceForWideColor() bool {
rv := objc.Send[bool](c.ID, objc.Sel("automaticallyConfiguresCaptureDeviceForWideColor"))
		return rv
}
func (c AVCaptureSession) SetAutomaticallyConfiguresCaptureDeviceForWideColor(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setAutomaticallyConfiguresCaptureDeviceForWideColor:"), value)
}

