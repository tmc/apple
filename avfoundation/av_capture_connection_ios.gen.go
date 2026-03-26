// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
// +build ios

package avfoundation

import (
"github.com/tmc/apple/objc"
)

// A Boolean value that indicates whether this connection supports video
// stabilization.
//
// # Discussion
// 
// The connection only supports video stabilization for video connection
// types, but may not be available for all resolutions.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoStabilizationSupported
func (c AVCaptureConnection) SupportsVideoStabilization() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isVideoStabilizationSupported"))
		return rv
}
// The connection’s current stabilization mode.
//
// # Discussion
// 
// The property only applies to a video connection, and it explicitly
// indicates whether it’s using stabilization, which means the value is
// never [CaptureVideoStabilizationModeAuto].
// 
// You can monitor this property to detect when the connection applies video
// stabilization to its video data with key-value observation. See
// [NSKeyValueObserving] and [Using Key-Value Observing in Swift] for more
// information.
//
// [NSKeyValueObserving]: https://developer.apple.com/documentation/ObjectiveC/nskeyvalueobserving
// [Using Key-Value Observing in Swift]: https://developer.apple.com/documentation/Swift/using-key-value-observing-in-swift
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/activeVideoStabilizationMode
func (c AVCaptureConnection) ActiveVideoStabilizationMode() AVCaptureVideoStabilizationMode {
rv := objc.Send[AVCaptureVideoStabilizationMode](c.ID, objc.Sel("activeVideoStabilizationMode"))
		return AVCaptureVideoStabilizationMode(rv)
}
// The stabilization mode that’s the most appropriate for a video
// connection.
//
// # Discussion
// 
// The property only applies to a video connection, and defaults to
// [CaptureVideoStabilizationModeOff].
// 
// You can enable video stabilization by setting it to an available
// stabilization mode (other than [CaptureVideoStabilizationModeOff]). Video
// stabilization introduces additional latency into the video capture pipeline
// and may consume more system memory, depending on the stabilization mode and
// format. If a stabilization mode isn’t available, the connection sets its
// [ActiveVideoStabilizationMode] property to
// [CaptureVideoStabilizationModeOff]. You can make the connection use an
// appropriate capture format and frame rate by setting the property to
// [CaptureVideoStabilizationModeAuto].
// 
// Use key-value observing with the [ActiveVideoStabilizationMode] property to
// determine which stabilization mode is in use.
// 
// You can monitor the [ActiveVideoStabilizationMode] property to detect which
// stabilization mode the connection’s using. See [NSKeyValueObserving] and
// [Using Key-Value Observing in Swift] for more information.
//
// [NSKeyValueObserving]: https://developer.apple.com/documentation/ObjectiveC/nskeyvalueobserving
// [Using Key-Value Observing in Swift]: https://developer.apple.com/documentation/Swift/using-key-value-observing-in-swift
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/preferredVideoStabilizationMode
func (c AVCaptureConnection) PreferredVideoStabilizationMode() AVCaptureVideoStabilizationMode {
rv := objc.Send[AVCaptureVideoStabilizationMode](c.ID, objc.Sel("preferredVideoStabilizationMode"))
		return AVCaptureVideoStabilizationMode(rv)
}
func (c AVCaptureConnection) SetPreferredVideoStabilizationMode(value AVCaptureVideoStabilizationMode) {
objc.Send[struct{}](c.ID, objc.Sel("setPreferredVideoStabilizationMode:"), value)
}
// A Boolean value that indicates whether the capture connection currently
// supports delivering camera intrinsics information.
//
// # Discussion
// 
// A value of [true] means you can set [CameraIntrinsicMatrixDeliveryEnabled]
// to [true]. The property is only [true] if both the connection’s input
// device format and output type support delivering camera intrinsics. In iOS
// 11, the [AVCaptureVideoDataOutput] class is the only output type that
// supports camera intrinsics.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isCameraIntrinsicMatrixDeliverySupported
func (c AVCaptureConnection) CameraIntrinsicMatrixDeliverySupported() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isCameraIntrinsicMatrixDeliverySupported"))
		return rv
}
// A Boolean value that indicates whether the connection can configure the
// capture pipeline to deliver camera intrinsics information.
//
// # Discussion
// 
// You can set this property to [true] for a video connection if
// [CameraIntrinsicMatrixDeliverySupported] is [true], and only before calling
// the [AVCaptureSession] [StartRunning] method. The default value is [false].
// 
// Camera intrinsics describe the current imaging parameters of a capture
// device in ways that you can use to render overlays or perform computer
// vision tasks. If [true], any [AVCaptureVideoDataOutput] instance in this
// connection can include the
// [kCMSampleBufferAttachmentKey_CameraIntrinsicMatrix] attachment for each
// sample buffer it vends.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [kCMSampleBufferAttachmentKey_CameraIntrinsicMatrix]: https://developer.apple.com/documentation/CoreMedia/kCMSampleBufferAttachmentKey_CameraIntrinsicMatrix
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isCameraIntrinsicMatrixDeliveryEnabled
func (c AVCaptureConnection) CameraIntrinsicMatrixDeliveryEnabled() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isCameraIntrinsicMatrixDeliveryEnabled"))
		return rv
}
func (c AVCaptureConnection) SetCameraIntrinsicMatrixDeliveryEnabled(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setCameraIntrinsicMatrixDeliveryEnabled:"), value)
}
// The connection’s maximum video scale and crop factor.
//
// # Discussion
// 
// The value defines the largest value you can set the
// [VideoScaleAndCropFactor] property to, which only applies to a video
// connection.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoMaxScaleAndCropFactor
func (c AVCaptureConnection) VideoMaxScaleAndCropFactor() float64 {
rv := objc.Send[float64](c.ID, objc.Sel("videoMaxScaleAndCropFactor"))
		return rv
}
// The current scale and crop factor the video output uses.
//
// # Discussion
// 
// The property only applies to a video connection. You can set this property
// to a value in the range `[1.0,`
// ``AVCaptureConnection/videoMaxScaleAndCropFactor```]`. A factor of `1.0`
// keeps the image at its original. Factors greater than `1.0` scale the image
// up and center-crop the image to its original dimensions.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/videoScaleAndCropFactor
func (c AVCaptureConnection) VideoScaleAndCropFactor() float64 {
rv := objc.Send[float64](c.ID, objc.Sel("videoScaleAndCropFactor"))
		return rv
}
func (c AVCaptureConnection) SetVideoScaleAndCropFactor(value float64) {
objc.Send[struct{}](c.ID, objc.Sel("setVideoScaleAndCropFactor:"), value)
}
// A Boolean value that indicates whether video stabilization is active for
// the connection.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/isVideoStabilizationEnabled
func (c AVCaptureConnection) VideoStabilizationEnabled() bool {
rv := objc.Send[bool](c.ID, objc.Sel("isVideoStabilizationEnabled"))
		return rv
}
// A Boolean value that indicates whether the system enables video
// stabilization when it’s available.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection/enablesVideoStabilizationWhenAvailable
func (c AVCaptureConnection) EnablesVideoStabilizationWhenAvailable() bool {
rv := objc.Send[bool](c.ID, objc.Sel("enablesVideoStabilizationWhenAvailable"))
		return rv
}
func (c AVCaptureConnection) SetEnablesVideoStabilizationWhenAvailable(value bool) {
objc.Send[struct{}](c.ID, objc.Sel("setEnablesVideoStabilizationWhenAvailable:"), value)
}

