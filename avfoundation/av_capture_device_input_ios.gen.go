// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
// +build ios

package avfoundation

import (
"github.com/tmc/apple/objc"
"github.com/tmc/apple/coremedia"
)

// Retrieves a virtual device’s constituent device ports for use in a
// multi-camera session.
//
// mediaType: The media type of the port you’re searching for, or `nil` to consider all
// media types.
//
// sourceDeviceType: The device type of the port you’re searching for, or `nil` if to consider
// all device types.
//
// sourceDevicePosition: The device position of the port you’re searching for.
// 
// When you’re searching for a camera device, a position of
// [CaptureDevicePositionUnspecified] indicates to search for all positions.
// 
// When you’re searching for an audio device,
// [CaptureDevicePositionUnspecified] indicates to search omnidirectional
// audio.
//
// # Return Value
// 
// An array of ports that satisfy the search criteria, or an empty array if
// there are none.
//
// # Discussion
// 
// [AVCaptureMultiCamSession] lets you simultaneously capture from multiple
// devices. It also lets you capture simultaneous streams from a virtual
// device, such as the dual camera. You use this method to find the ports
// associated with a virtual device’s underlying physical devices. A virtual
// device input’s [Ports] property doesn’t include constituent device
// ports.
// 
// Using the dual camera as an example, the [Ports] property exposes only
// those ports supported by the virtual device (it switches automatically
// between wide-angle and telephoto cameras, depending on the zoom factor).
// You may use this method to find the video ports for the constituent
// devices.
// 
// You can use these ports to create connections to two instances of
// [AVCaptureVideoDataOutput], allowing for synchronized, full-frame-rate
// delivery of both wide-angle and telephoto streams.
// 
// When used in conjunction with an audio device, this method allows you to
// discover microphones in different positions. You can use the microphone
// ports to make output connections to simultaneously capture both
// front-facing and back-facing audio. The audio device port whose
// [SourceDevicePosition] is [CaptureDevicePositionUnspecified] produces
// omnidirectional sound.
//
// [AVCaptureMultiCamSession]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultiCamSession
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/ports(for:sourceDeviceType:sourceDevicePosition:)
func (c AVCaptureDeviceInput) PortsWithMediaTypeSourceDeviceTypeSourceDevicePosition(mediaType AVMediaType, sourceDeviceType AVCaptureDeviceType, sourceDevicePosition AVCaptureDevicePosition) []AVCaptureInputPort {
rv := objc.Send[[]objc.ID](c.ID, objc.Sel("portsWithMediaType:sourceDeviceType:sourceDevicePosition:"), objc.String(string(mediaType)), objc.String(string(sourceDeviceType)), sourceDevicePosition)
return objc.ConvertSlice(rv, func(id objc.ID) AVCaptureInputPort {
	return AVCaptureInputPortFromID(id)
})
}

// A Boolean value that indicates whether the input enables unified
// auto-exposure defaults.
//
// # Discussion
// 
// You may set the value of a capture device’s [ActiveFormat] in two ways:
// 
// - Set it directly using one of the formats in the device’s [Formats]
// property. - The capture session sets it on your behalf when you set its
// [SessionPreset] property.
// 
// Depending on the device and format, you may configure the default auto
// exposure behavior differently when you use one method or the other,
// resulting in non-uniform auto exposure behavior. Auto exposure defaults
// include [MinFrameRate], [MaxFrameRate], and [MaxExposureDuration]. You can
// set this property to [true] to ensure that the system applies consistent
// default behaviors to the device regardless of the way you set the active
// format.
// 
// The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/unifiedAutoExposureDefaultsEnabled
func (c AVCaptureDeviceInput) UnifiedAutoExposureDefaultsEnabled() bool {
rv := objc.Send[bool](c.ID, objc.Sel("unifiedAutoExposureDefaultsEnabled"))
		return rv
}
func (c AVCaptureDeviceInput) SetUnifiedAutoExposureDefaultsEnabled(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setUnifiedAutoExposureDefaultsEnabled:"), value)
}
// A time value that acts as a modifier to a capture device’s active video
// minimum frame duration.
//
// # Discussion
// 
// A capture device’s [ActiveVideoMinFrameDuration] property is the
// reciprocal of its active maximum frame rate. To limit the maximum frame
// rate of the capture device, you may set [ActiveVideoMinFrameDuration] to a
// value supported by the device’s [ActiveFormat]. Changes you make to the
// device’s [ActiveVideoMinFrameDuration] value take effect immediately
// without disrupting preview. Therefore, [AVCaptureSession] must always
// allocate sufficient resources to allow the device to run at its active
// format’s maximum frame rate.
// 
// If you wish to use a particular device format but run it only at lower
// frame rates (for instance, only run a 1080p240 fps format at a maximum
// frame rate of 60), set the input’s [VideoMinFrameDurationOverride]
// property to the reciprocal of the maximum frame rate you intend to use
// before starting the session (or within a [BeginConfiguration] /
// [CommitConfiguration] block while running the session).
// 
// This property’s default value is [invalid]. When you remove a device
// input from a session and then add it back, this property reverts to its
// default value.
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput/videoMinFrameDurationOverride
func (c AVCaptureDeviceInput) VideoMinFrameDurationOverride() coremedia.CMTime {
rv := objc.Send[coremedia.CMTime](c.ID, objc.Sel("videoMinFrameDurationOverride"))
		return coremedia.CMTime(rv)
}
func (c AVCaptureDeviceInput) SetVideoMinFrameDurationOverride(value coremedia.CMTime) {
objc.Send[struct{}](c.ID, objc.Sel("setVideoMinFrameDurationOverride:"), value)
}

