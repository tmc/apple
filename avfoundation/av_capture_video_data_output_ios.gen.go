// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
// +build ios

package avfoundation

import (
"github.com/tmc/apple/objc"
"github.com/tmc/apple/coremedia"
)

// A Boolean value that indicates whether the output automatically configures
// the size of output buffers.
//
// # Discussion
// 
// In most configurations, [AVCaptureVideoDataOutput] delivers full-resolution
// buffers that match the video dimensions of the capture device’s
// [ActiveFormat] property. When this property is [true], the output is free
// to scale the buffers delivered to
// [CaptureOutputDidOutputSampleBufferFromConnection] to a size suitable for
// preview (approximately the size of the screen).
// 
// You can query this property to find out whether the automatic configuration
// of output buffer dimensions is downscaling buffers to a preview size. You
// can also query the output’s [VideoSettings] dictionary to find the
// buffer’s exact dimensions.
// 
// The default value of this property is [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/automaticallyConfiguresOutputBufferDimensions
func (c AVCaptureVideoDataOutput) AutomaticallyConfiguresOutputBufferDimensions() bool {
rv := objc.Send[bool](c.ID, objc.Sel("automaticallyConfiguresOutputBufferDimensions"))
		return rv
}
func (c AVCaptureVideoDataOutput) SetAutomaticallyConfiguresOutputBufferDimensions(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setAutomaticallyConfiguresOutputBufferDimensions:"), value)
}
// A Boolean value that indicates whether the output is configured to deliver
// preview-sized buffers.
//
// # Discussion
// 
// [AVCaptureVideoDataOutput] produces preview-sized buffers, which are
// approximately the size of the screen, when its
// [AutomaticallyConfiguresOutputBufferDimensions] property is [true]. If you
// wish to manually set this property, you must first set
// [AutomaticallyConfiguresOutputBufferDimensions] to [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/deliversPreviewSizedOutputBuffers
func (c AVCaptureVideoDataOutput) DeliversPreviewSizedOutputBuffers() bool {
rv := objc.Send[bool](c.ID, objc.Sel("deliversPreviewSizedOutputBuffers"))
		return rv
}
func (c AVCaptureVideoDataOutput) SetDeliversPreviewSizedOutputBuffers(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setDeliversPreviewSizedOutputBuffers:"), value)
}
// Indicates whether the receiver should prepare the cellular radio for
// imminent network activity.
//
// # Discussion
// 
// Apps that scan video data output buffers for information that will result
// in network activity (such as detecting a QRCode containing a URL) should
// set this property `true` to allow the cellular radio to prepare for an
// imminent network request. Enabling this property requires a lengthy
// reconfiguration of the capture render pipeline, so you should set this
// property to `true` before calling [StartRunning].
// 
// Using this API requires your app to adopt the entitlement
// `com.AppleXCUIElementTypeDeveloperXCUIElementTypeAvfoundationXCUIElementTypeVideo()-data-output-prepares-cellular-radio-for-machine-readable-code-scanning`.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/preparesCellularRadioForNetworkConnection
func (c AVCaptureVideoDataOutput) PreparesCellularRadioForNetworkConnection() bool {
rv := objc.Send[bool](c.ID, objc.Sel("preparesCellularRadioForNetworkConnection"))
		return rv
}
func (c AVCaptureVideoDataOutput) SetPreparesCellularRadioForNetworkConnection(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setPreparesCellularRadioForNetworkConnection:"), value)
}
// The minimum frame duration.
//
// # Discussion
// 
// This property specifies the minimum duration of each video frame produced
// by the video data output, placing a lower bound on the amount of time that
// should separate consecutive frames. This is equivalent to the inverse of
// the maximum frame rate. A value of [zero] or [invalid] indicates an
// unlimited maximum frame rate.
// 
// The default value is [invalid].
//
// [invalid]: https://developer.apple.com/documentation/CoreMedia/CMTime/invalid
// [zero]: https://developer.apple.com/documentation/CoreMedia/CMTime/zero
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/minFrameDuration
func (c AVCaptureVideoDataOutput) MinFrameDuration() coremedia.CMTime {
rv := objc.Send[coremedia.CMTime](c.ID, objc.Sel("minFrameDuration"))
		return coremedia.CMTime(rv)
}
func (c AVCaptureVideoDataOutput) SetMinFrameDuration(value coremedia.CMTime) {
objc.Send[struct{}](c.ID, objc.Sel("setMinFrameDuration:"), value)
}

