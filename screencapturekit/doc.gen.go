
// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

// Package screencapturekit provides Go bindings for the ScreenCaptureKit framework.
//
// Filter and select screen content and stream it to your app.
//
// Use the ScreenCaptureKit framework to add support for high-performance frame capture of screen and audio content to your Mac app. The framework gives you fine-grained control to select and stream only the content that you want to capture. As a stream captures new video frames and audio samples, it passes them to your app as [CMSampleBuffer](<doc://com.apple.documentation/documentation/CoreMedia/CMSampleBuffer>) objects that contain the media data and its related metadata. ScreenCaptureKit also provides a macOS-integrated picker for streaming selection and management, [SCContentSharingPicker](<doc://com.apple.screencapturekit/documentation/ScreenCaptureKit/SCContentSharingPicker>).
//
// # Essentials
//
//   - ScreenCaptureKit updates: Learn about important changes to ScreenCaptureKit.
//   - Persistent Content Capture: A Boolean value that indicates whether a Virtual Network Computing (VNC) app needs persistent access to screen capture.
//   - Capturing screen content in macOS: Stream desktop content like displays, apps, and windows by adopting screen capture in your app.
//
// # Shareable content
//
//   - SCShareableContent: An instance that represents a set of displays, apps, and windows that your app can capture.
//   - SCShareableContentInfo: An instance that provides information for the content in a given stream.
//   - SCShareableContentStyle: The style of content presented in a stream.
//   - SCDisplay: An instance that represents a display device.
//   - SCRunningApplication: An instance that represents an app running on a device.
//   - SCWindow: An instance that represents an onscreen window.
//
// # Content capture
//
//   - SCStream: An instance that represents a stream of shareable content. ([SCRecordingOutput])
//   - SCStreamConfiguration: An instance that provides the output configuration for a stream. ([SCCaptureResolutionType], [SCPresenterOverlayAlertSetting], [SCCaptureDynamicRange])
//   - SCContentFilter: An instance that filters the content a stream captures. ([SCStreamType])
//   - SCStreamDelegate: A delegate protocol your app implements to respond to stream events.
//   - SCScreenshotManager: An instance for the capture of single frames from a stream.
//   - SCScreenshotConfiguration: An object that contains screenshot properties such as output width, height, and image quality specifications.
//   - SCScreenshotOutput: An object that contains all images requested by the client.
//
// # Output processing
//
//   - SCStreamOutput: A delegate protocol your app implements to receive capture stream output events.
//   - SCStreamOutputType: Constants that represent output types for a stream frame.
//   - SCStreamFrameInfo: An instance that defines metadata keys for a stream frame.
//   - SCFrameStatus: Status values for a frame from a stream.
//
// # System content-sharing picker
//
//   - SCContentSharingPicker: An instance of a picker presented by the operating system for managing frame-capture streams.
//   - SCContentSharingPickerConfiguration: An instance for configuring the system content-sharing picker.
//   - SCContentSharingPickerMode: Available modes for selecting streaming content from a picker presented by the operating system.
//   - SCContentSharingPickerObserver: An observer protocol your app implements to receive messages from the operating system’s content picker.
//
// # Stream errors
//
//   - SCStreamErrorDomain: A string representation of the error domain.
//   - SCStreamError: An instance representing a ScreenCaptureKit framework error.
//
// # Key Types
//
//   - [SCStreamConfiguration] - An instance that provides the output configuration for a stream.
//   - [SCContentFilter] - An instance that filters the content a stream captures.
//   - [SCScreenshotConfiguration] - An object that contains screenshot properties such as output width, height, and image quality specifications.
//   - [SCContentSharingPicker] - An instance of a picker presented by the operating system for managing frame-capture streams.
//   - [SCStream] - An instance that represents a stream of shareable content.
//   - [SCShareableContent] - An instance that represents a set of displays, apps, and windows that your app can capture.
//   - [SCWindow] - An instance that represents an onscreen window.
//   - [SCRecordingOutputConfiguration]
//   - [SCScreenshotManager] - An instance for the capture of single frames from a stream.
//   - [SCContentSharingPickerConfiguration] - An instance for configuring the system content-sharing picker.
//
// [ScreenCaptureKit Documentation]: https://developer.apple.com/documentation/ScreenCaptureKit
package screencapturekit

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/ScreenCaptureKit.framework/ScreenCaptureKit"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: ScreenCaptureKit: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

