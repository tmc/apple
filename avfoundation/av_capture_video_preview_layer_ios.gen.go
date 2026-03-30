// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.
//go:build ios
// +build ios

package avfoundation

import (
	"github.com/tmc/apple/objc"
)

// A Boolean value that indicates whether the layer is rendering video frames
// from its source.
//
// # Discussion
//
// A preview layer begins displaying content when you call the capture
// session’s [StartRunning] method. If you associate the layer with an
// instance of [AVCaptureMultiCamSession], the system guarantees that all
// video preview layers display content by the time the blocking call to
// [StartRunning] or [CommitConfiguration] returns.
//
// While a session is running, you may enable or disable a video preview
// layer’s connection to start or stop the flow of video to the layer. You
// may key-value observe the connection’s [Enabled] property to observe this
// property changing, and synchronize any user interface changes to take place
// precisely when the video resumes rendering to the video preview layer.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/isPreviewing
//
// [AVCaptureMultiCamSession]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultiCamSession
func (c AVCaptureVideoPreviewLayer) Previewing() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isPreviewing"))
	return rv
}

// Indicates whether the layer display automatically adjusts mirroring.
//
// # Discussion
//
// For some session configurations, preview will be mirrored by default.
//
// When the value of this property is true, the value of [Mirrored] may change
// depending on the configuration of the session, for example after switching
// to a different capture device input.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/automaticallyAdjustsMirroring
func (c AVCaptureVideoPreviewLayer) AutomaticallyAdjustsMirroring() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("automaticallyAdjustsMirroring"))
	return rv
}
func (c AVCaptureVideoPreviewLayer) SetAutomaticallyAdjustsMirroring(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setAutomaticallyAdjustsMirroring:"), value)
}

// Indicates whether the layer display is mirrored.
//
// # Discussion
//
// To change the value of this property, the value of
// [AutomaticallyAdjustsMirroring] must be false.
//
// Mirroring is not supported on all hardware configurations. You should check
// the value of [SupportsVideoMirroring] ([AVCaptureConnection]) before
// attempting to change this value.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/mirrored
func (c AVCaptureVideoPreviewLayer) Mirrored() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isMirrored"))
	return rv
}
func (c AVCaptureVideoPreviewLayer) SetMirrored(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setMirrored:"), value)
}

// Indicates whether the layer display supports mirroring.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/mirroringSupported
func (c AVCaptureVideoPreviewLayer) MirroringSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isMirroringSupported"))
	return rv
}

// The layer’s orientation.
//
// # Discussion
//
// Changes in orientation are not supported on all hardware configurations.
// You should check the value of [SupportsVideoOrientation]
// ([AVCaptureConnection]) before attempting to change the orientation of the
// receiver. An exception is raised if this requirement is ignored.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/orientation
func (c AVCaptureVideoPreviewLayer) Orientation() AVCaptureVideoOrientation {
	rv := objc.Send[AVCaptureVideoOrientation](c.ID, objc.Sel("orientation"))
	return AVCaptureVideoOrientation(rv)
}
func (c AVCaptureVideoPreviewLayer) SetOrientation(value AVCaptureVideoOrientation) {
	objc.Send[struct{}](c.ID, objc.Sel("setOrientation:"), value)
}

// Indicates whether the layer display supports changing the orientation.
//
// See: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoPreviewLayer/orientationSupported
func (c AVCaptureVideoPreviewLayer) OrientationSupported() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isOrientationSupported"))
	return rv
}
