// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/coremedia"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SCStreamConfiguration] class.
var (
	_SCStreamConfigurationClass     SCStreamConfigurationClass
	_SCStreamConfigurationClassOnce sync.Once
)

func getSCStreamConfigurationClass() SCStreamConfigurationClass {
	_SCStreamConfigurationClassOnce.Do(func() {
		_SCStreamConfigurationClass = SCStreamConfigurationClass{class: objc.GetClass("SCStreamConfiguration")}
	})
	return _SCStreamConfigurationClass
}

// GetSCStreamConfigurationClass returns the class object for SCStreamConfiguration.
func GetSCStreamConfigurationClass() SCStreamConfigurationClass {
	return getSCStreamConfigurationClass()
}

type SCStreamConfigurationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SCStreamConfigurationClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SCStreamConfigurationClass) Alloc() SCStreamConfiguration {
	rv := objc.Send[SCStreamConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An instance that provides the output configuration for a stream.
//
// # Overview
//
// Creating an instance of this class provides a default configuration for a
// stream. Only configure its properties if you need to customize the output.
//
// # Specifying dimensions
//
//   - [SCStreamConfiguration.Width]: The width of the output.
//   - [SCStreamConfiguration.SetWidth]
//   - [SCStreamConfiguration.Height]: The height of the output.
//   - [SCStreamConfiguration.SetHeight]
//   - [SCStreamConfiguration.ScalesToFit]: A Boolean value that indicates whether to scale the output to fit the configured width and height.
//   - [SCStreamConfiguration.SetScalesToFit]
//   - [SCStreamConfiguration.SourceRect]: A rectangle that specifies the source area to capture.
//   - [SCStreamConfiguration.SetSourceRect]
//   - [SCStreamConfiguration.DestinationRect]: A rectangle that specifies a destination into which to write the output.
//   - [SCStreamConfiguration.SetDestinationRect]
//   - [SCStreamConfiguration.PreservesAspectRatio]: A Boolean value that determines if the stream preserves aspect ratio.
//   - [SCStreamConfiguration.SetPreservesAspectRatio]
//
// # Configuring colors
//
//   - [SCStreamConfiguration.PixelFormat]: A pixel format for sample buffers that a stream outputs.
//   - [SCStreamConfiguration.SetPixelFormat]
//   - [SCStreamConfiguration.ColorMatrix]: A color matrix to apply to the output surface.
//   - [SCStreamConfiguration.SetColorMatrix]
//   - [SCStreamConfiguration.ColorSpaceName]: A color space to use for the output buffer.
//   - [SCStreamConfiguration.SetColorSpaceName]
//   - [SCStreamConfiguration.BackgroundColor]: A background color for the output.
//   - [SCStreamConfiguration.SetBackgroundColor]
//
// # Configuring captured elements
//
//   - [SCStreamConfiguration.ShowsCursor]: A Boolean value that determines whether the cursor is visible in the stream.
//   - [SCStreamConfiguration.SetShowsCursor]
//   - [SCStreamConfiguration.ShouldBeOpaque]: A Boolean value that indicates if semitransparent content presents as opaque.
//   - [SCStreamConfiguration.SetShouldBeOpaque]
//   - [SCStreamConfiguration.CapturesShadowsOnly]: A Boolean value that indicates if the stream only captures shadows.
//   - [SCStreamConfiguration.SetCapturesShadowsOnly]
//   - [SCStreamConfiguration.IgnoreShadowsDisplay]: A Boolean value that indicates if the stream ignores the capturing of window shadows when streaming in display style.
//   - [SCStreamConfiguration.SetIgnoreShadowsDisplay]
//   - [SCStreamConfiguration.IgnoreShadowsSingleWindow]: A Boolean value that indicates if the stream ignores the capturing of window shadows when streaming in window style.
//   - [SCStreamConfiguration.SetIgnoreShadowsSingleWindow]
//   - [SCStreamConfiguration.IgnoreGlobalClipDisplay]: A Boolean value that indicates if the stream ignores content clipped past the edge of a display, when streaming in display style.
//   - [SCStreamConfiguration.SetIgnoreGlobalClipDisplay]
//   - [SCStreamConfiguration.IgnoreGlobalClipSingleWindow]: A Boolean value that indicates if the stream ignores content clipped past the edge of a display, when streaming in window style.
//   - [SCStreamConfiguration.SetIgnoreGlobalClipSingleWindow]
//
// # Configuring captured frames
//
//   - [SCStreamConfiguration.QueueDepth]: The maximum number of frames for the queue to store.
//   - [SCStreamConfiguration.SetQueueDepth]
//   - [SCStreamConfiguration.MinimumFrameInterval]: The desired minimum time between frame updates, in seconds.
//   - [SCStreamConfiguration.SetMinimumFrameInterval]
//   - [SCStreamConfiguration.CaptureResolution]: The resolution at which to capture source content.
//   - [SCStreamConfiguration.SetCaptureResolution]
//
// # Configuring audio
//
//   - [SCStreamConfiguration.CapturesAudio]: A Boolean value that indicates whether to capture audio.
//   - [SCStreamConfiguration.SetCapturesAudio]
//   - [SCStreamConfiguration.SampleRate]: The sample rate for audio capture.
//   - [SCStreamConfiguration.SetSampleRate]
//   - [SCStreamConfiguration.ChannelCount]: The number of audio channels to capture.
//   - [SCStreamConfiguration.SetChannelCount]
//   - [SCStreamConfiguration.ExcludesCurrentProcessAudio]: A Boolean value that indicates whether to exclude audio from your app during capture.
//   - [SCStreamConfiguration.SetExcludesCurrentProcessAudio]
//
// # Identifying a stream
//
//   - [SCStreamConfiguration.StreamName]: A name that you provide for identifying the stream.
//   - [SCStreamConfiguration.SetStreamName]
//
// # Notifying presenters
//
//   - [SCStreamConfiguration.PresenterOverlayPrivacyAlertSetting]: A value indicating if alerts appear to presenters while using Presenter Overlay.
//   - [SCStreamConfiguration.SetPresenterOverlayPrivacyAlertSetting]
//
// # Instance Properties
//
//   - [SCStreamConfiguration.CaptureDynamicRange]
//   - [SCStreamConfiguration.SetCaptureDynamicRange]
//   - [SCStreamConfiguration.CaptureMicrophone]
//   - [SCStreamConfiguration.SetCaptureMicrophone]
//   - [SCStreamConfiguration.IncludeChildWindows]
//   - [SCStreamConfiguration.SetIncludeChildWindows]
//   - [SCStreamConfiguration.MicrophoneCaptureDeviceID]
//   - [SCStreamConfiguration.SetMicrophoneCaptureDeviceID]
//   - [SCStreamConfiguration.ShowMouseClicks]
//   - [SCStreamConfiguration.SetShowMouseClicks]
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration
type SCStreamConfiguration struct {
	objectivec.Object
}

// SCStreamConfigurationFromID constructs a [SCStreamConfiguration] from an objc.ID.
//
// An instance that provides the output configuration for a stream.
func SCStreamConfigurationFromID(id objc.ID) SCStreamConfiguration {
	return SCStreamConfiguration{objectivec.Object{ID: id}}
}

// NOTE: SCStreamConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SCStreamConfiguration] class.
//
// # Specifying dimensions
//
//   - [ISCStreamConfiguration.Width]: The width of the output.
//   - [ISCStreamConfiguration.SetWidth]
//   - [ISCStreamConfiguration.Height]: The height of the output.
//   - [ISCStreamConfiguration.SetHeight]
//   - [ISCStreamConfiguration.ScalesToFit]: A Boolean value that indicates whether to scale the output to fit the configured width and height.
//   - [ISCStreamConfiguration.SetScalesToFit]
//   - [ISCStreamConfiguration.SourceRect]: A rectangle that specifies the source area to capture.
//   - [ISCStreamConfiguration.SetSourceRect]
//   - [ISCStreamConfiguration.DestinationRect]: A rectangle that specifies a destination into which to write the output.
//   - [ISCStreamConfiguration.SetDestinationRect]
//   - [ISCStreamConfiguration.PreservesAspectRatio]: A Boolean value that determines if the stream preserves aspect ratio.
//   - [ISCStreamConfiguration.SetPreservesAspectRatio]
//
// # Configuring colors
//
//   - [ISCStreamConfiguration.PixelFormat]: A pixel format for sample buffers that a stream outputs.
//   - [ISCStreamConfiguration.SetPixelFormat]
//   - [ISCStreamConfiguration.ColorMatrix]: A color matrix to apply to the output surface.
//   - [ISCStreamConfiguration.SetColorMatrix]
//   - [ISCStreamConfiguration.ColorSpaceName]: A color space to use for the output buffer.
//   - [ISCStreamConfiguration.SetColorSpaceName]
//   - [ISCStreamConfiguration.BackgroundColor]: A background color for the output.
//   - [ISCStreamConfiguration.SetBackgroundColor]
//
// # Configuring captured elements
//
//   - [ISCStreamConfiguration.ShowsCursor]: A Boolean value that determines whether the cursor is visible in the stream.
//   - [ISCStreamConfiguration.SetShowsCursor]
//   - [ISCStreamConfiguration.ShouldBeOpaque]: A Boolean value that indicates if semitransparent content presents as opaque.
//   - [ISCStreamConfiguration.SetShouldBeOpaque]
//   - [ISCStreamConfiguration.CapturesShadowsOnly]: A Boolean value that indicates if the stream only captures shadows.
//   - [ISCStreamConfiguration.SetCapturesShadowsOnly]
//   - [ISCStreamConfiguration.IgnoreShadowsDisplay]: A Boolean value that indicates if the stream ignores the capturing of window shadows when streaming in display style.
//   - [ISCStreamConfiguration.SetIgnoreShadowsDisplay]
//   - [ISCStreamConfiguration.IgnoreShadowsSingleWindow]: A Boolean value that indicates if the stream ignores the capturing of window shadows when streaming in window style.
//   - [ISCStreamConfiguration.SetIgnoreShadowsSingleWindow]
//   - [ISCStreamConfiguration.IgnoreGlobalClipDisplay]: A Boolean value that indicates if the stream ignores content clipped past the edge of a display, when streaming in display style.
//   - [ISCStreamConfiguration.SetIgnoreGlobalClipDisplay]
//   - [ISCStreamConfiguration.IgnoreGlobalClipSingleWindow]: A Boolean value that indicates if the stream ignores content clipped past the edge of a display, when streaming in window style.
//   - [ISCStreamConfiguration.SetIgnoreGlobalClipSingleWindow]
//
// # Configuring captured frames
//
//   - [ISCStreamConfiguration.QueueDepth]: The maximum number of frames for the queue to store.
//   - [ISCStreamConfiguration.SetQueueDepth]
//   - [ISCStreamConfiguration.MinimumFrameInterval]: The desired minimum time between frame updates, in seconds.
//   - [ISCStreamConfiguration.SetMinimumFrameInterval]
//   - [ISCStreamConfiguration.CaptureResolution]: The resolution at which to capture source content.
//   - [ISCStreamConfiguration.SetCaptureResolution]
//
// # Configuring audio
//
//   - [ISCStreamConfiguration.CapturesAudio]: A Boolean value that indicates whether to capture audio.
//   - [ISCStreamConfiguration.SetCapturesAudio]
//   - [ISCStreamConfiguration.SampleRate]: The sample rate for audio capture.
//   - [ISCStreamConfiguration.SetSampleRate]
//   - [ISCStreamConfiguration.ChannelCount]: The number of audio channels to capture.
//   - [ISCStreamConfiguration.SetChannelCount]
//   - [ISCStreamConfiguration.ExcludesCurrentProcessAudio]: A Boolean value that indicates whether to exclude audio from your app during capture.
//   - [ISCStreamConfiguration.SetExcludesCurrentProcessAudio]
//
// # Identifying a stream
//
//   - [ISCStreamConfiguration.StreamName]: A name that you provide for identifying the stream.
//   - [ISCStreamConfiguration.SetStreamName]
//
// # Notifying presenters
//
//   - [ISCStreamConfiguration.PresenterOverlayPrivacyAlertSetting]: A value indicating if alerts appear to presenters while using Presenter Overlay.
//   - [ISCStreamConfiguration.SetPresenterOverlayPrivacyAlertSetting]
//
// # Instance Properties
//
//   - [ISCStreamConfiguration.CaptureDynamicRange]
//   - [ISCStreamConfiguration.SetCaptureDynamicRange]
//   - [ISCStreamConfiguration.CaptureMicrophone]
//   - [ISCStreamConfiguration.SetCaptureMicrophone]
//   - [ISCStreamConfiguration.IncludeChildWindows]
//   - [ISCStreamConfiguration.SetIncludeChildWindows]
//   - [ISCStreamConfiguration.MicrophoneCaptureDeviceID]
//   - [ISCStreamConfiguration.SetMicrophoneCaptureDeviceID]
//   - [ISCStreamConfiguration.ShowMouseClicks]
//   - [ISCStreamConfiguration.SetShowMouseClicks]
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration
type ISCStreamConfiguration interface {
	objectivec.IObject

	// Topic: Specifying dimensions

	// The width of the output.
	Width() uintptr
	SetWidth(value uintptr)
	// The height of the output.
	Height() uintptr
	SetHeight(value uintptr)
	// A Boolean value that indicates whether to scale the output to fit the configured width and height.
	ScalesToFit() bool
	SetScalesToFit(value bool)
	// A rectangle that specifies the source area to capture.
	SourceRect() corefoundation.CGRect
	SetSourceRect(value corefoundation.CGRect)
	// A rectangle that specifies a destination into which to write the output.
	DestinationRect() corefoundation.CGRect
	SetDestinationRect(value corefoundation.CGRect)
	// A Boolean value that determines if the stream preserves aspect ratio.
	PreservesAspectRatio() bool
	SetPreservesAspectRatio(value bool)

	// Topic: Configuring colors

	// A pixel format for sample buffers that a stream outputs.
	PixelFormat() uint32
	SetPixelFormat(value uint32)
	// A color matrix to apply to the output surface.
	ColorMatrix() corefoundation.CFStringRef
	SetColorMatrix(value corefoundation.CFStringRef)
	// A color space to use for the output buffer.
	ColorSpaceName() corefoundation.CFStringRef
	SetColorSpaceName(value corefoundation.CFStringRef)
	// A background color for the output.
	BackgroundColor() coregraphics.CGColorRef
	SetBackgroundColor(value coregraphics.CGColorRef)

	// Topic: Configuring captured elements

	// A Boolean value that determines whether the cursor is visible in the stream.
	ShowsCursor() bool
	SetShowsCursor(value bool)
	// A Boolean value that indicates if semitransparent content presents as opaque.
	ShouldBeOpaque() bool
	SetShouldBeOpaque(value bool)
	// A Boolean value that indicates if the stream only captures shadows.
	CapturesShadowsOnly() bool
	SetCapturesShadowsOnly(value bool)
	// A Boolean value that indicates if the stream ignores the capturing of window shadows when streaming in display style.
	IgnoreShadowsDisplay() bool
	SetIgnoreShadowsDisplay(value bool)
	// A Boolean value that indicates if the stream ignores the capturing of window shadows when streaming in window style.
	IgnoreShadowsSingleWindow() bool
	SetIgnoreShadowsSingleWindow(value bool)
	// A Boolean value that indicates if the stream ignores content clipped past the edge of a display, when streaming in display style.
	IgnoreGlobalClipDisplay() bool
	SetIgnoreGlobalClipDisplay(value bool)
	// A Boolean value that indicates if the stream ignores content clipped past the edge of a display, when streaming in window style.
	IgnoreGlobalClipSingleWindow() bool
	SetIgnoreGlobalClipSingleWindow(value bool)

	// Topic: Configuring captured frames

	// The maximum number of frames for the queue to store.
	QueueDepth() int
	SetQueueDepth(value int)
	// The desired minimum time between frame updates, in seconds.
	MinimumFrameInterval() coremedia.CMTime
	SetMinimumFrameInterval(value coremedia.CMTime)
	// The resolution at which to capture source content.
	CaptureResolution() SCCaptureResolutionType
	SetCaptureResolution(value SCCaptureResolutionType)

	// Topic: Configuring audio

	// A Boolean value that indicates whether to capture audio.
	CapturesAudio() bool
	SetCapturesAudio(value bool)
	// The sample rate for audio capture.
	SampleRate() int
	SetSampleRate(value int)
	// The number of audio channels to capture.
	ChannelCount() int
	SetChannelCount(value int)
	// A Boolean value that indicates whether to exclude audio from your app during capture.
	ExcludesCurrentProcessAudio() bool
	SetExcludesCurrentProcessAudio(value bool)

	// Topic: Identifying a stream

	// A name that you provide for identifying the stream.
	StreamName() string
	SetStreamName(value string)

	// Topic: Notifying presenters

	// A value indicating if alerts appear to presenters while using Presenter Overlay.
	PresenterOverlayPrivacyAlertSetting() SCPresenterOverlayAlertSetting
	SetPresenterOverlayPrivacyAlertSetting(value SCPresenterOverlayAlertSetting)

	// Topic: Instance Properties

	CaptureDynamicRange() SCCaptureDynamicRange
	SetCaptureDynamicRange(value SCCaptureDynamicRange)
	CaptureMicrophone() bool
	SetCaptureMicrophone(value bool)
	IncludeChildWindows() bool
	SetIncludeChildWindows(value bool)
	MicrophoneCaptureDeviceID() string
	SetMicrophoneCaptureDeviceID(value string)
	ShowMouseClicks() bool
	SetShowMouseClicks(value bool)
}

// Init initializes the instance.
func (s SCStreamConfiguration) Init() SCStreamConfiguration {
	rv := objc.Send[SCStreamConfiguration](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SCStreamConfiguration) Autorelease() SCStreamConfiguration {
	rv := objc.Send[SCStreamConfiguration](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCStreamConfiguration creates a new SCStreamConfiguration instance.
func NewSCStreamConfiguration() SCStreamConfiguration {
	class := getSCStreamConfigurationClass()
	rv := objc.Send[SCStreamConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/init(preset:)
func NewStreamConfigurationWithPreset(preset SCStreamConfigurationPreset) SCStreamConfiguration {
	rv := objc.Send[objc.ID](objc.ID(getSCStreamConfigurationClass().class), objc.Sel("streamConfigurationWithPreset:"), preset)
	return SCStreamConfigurationFromID(rv)
}

// The width of the output.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/width
func (s SCStreamConfiguration) Width() uintptr {
	rv := objc.Send[uintptr](s.ID, objc.Sel("width"))
	return rv
}
func (s SCStreamConfiguration) SetWidth(value uintptr) {
	objc.Send[struct{}](s.ID, objc.Sel("setWidth:"), value)
}

// The height of the output.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/height
func (s SCStreamConfiguration) Height() uintptr {
	rv := objc.Send[uintptr](s.ID, objc.Sel("height"))
	return rv
}
func (s SCStreamConfiguration) SetHeight(value uintptr) {
	objc.Send[struct{}](s.ID, objc.Sel("setHeight:"), value)
}

// A Boolean value that indicates whether to scale the output to fit the
// configured width and height.
//
// # Discussion
//
// The system uses this value during independent window capture.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/scalesToFit
func (s SCStreamConfiguration) ScalesToFit() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("scalesToFit"))
	return rv
}
func (s SCStreamConfiguration) SetScalesToFit(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setScalesToFit:"), value)
}

// A rectangle that specifies the source area to capture.
//
// # Discussion
//
// If you don’t specify a source rectangle to capture, the system captures
// the entire display.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/sourceRect
func (s SCStreamConfiguration) SourceRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("sourceRect"))
	return corefoundation.CGRect(rv)
}
func (s SCStreamConfiguration) SetSourceRect(value corefoundation.CGRect) {
	objc.Send[struct{}](s.ID, objc.Sel("setSourceRect:"), value)
}

// A rectangle that specifies a destination into which to write the output.
//
// # Discussion
//
// If you don’t specify a destination rectangle, the system uses the full
// dimensions of the output surface.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/destinationRect
func (s SCStreamConfiguration) DestinationRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("destinationRect"))
	return corefoundation.CGRect(rv)
}
func (s SCStreamConfiguration) SetDestinationRect(value corefoundation.CGRect) {
	objc.Send[struct{}](s.ID, objc.Sel("setDestinationRect:"), value)
}

// A Boolean value that determines if the stream preserves aspect ratio.
//
// # Discussion
//
// The default value is `true`.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/preservesAspectRatio
func (s SCStreamConfiguration) PreservesAspectRatio() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("preservesAspectRatio"))
	return rv
}
func (s SCStreamConfiguration) SetPreservesAspectRatio(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setPreservesAspectRatio:"), value)
}

// A pixel format for sample buffers that a stream outputs.
//
// # Discussion
//
// A stream supports the following pixel formats:
//
// [BGRA]: Packed little endian ARGB8888. `l10r`: Packed little endian
// ARGB2101010. `420v`: Two-plane “video” range YCbCr 4:2:0. `420f`:
// Two-plane “full” range YCbCr 4:2:0.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/pixelFormat
func (s SCStreamConfiguration) PixelFormat() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("pixelFormat"))
	return rv
}
func (s SCStreamConfiguration) SetPixelFormat(value uint32) {
	objc.Send[struct{}](s.ID, objc.Sel("setPixelFormat:"), value)
}

// A color matrix to apply to the output surface.
//
// # Discussion
//
// You can specify a value for this property if your pixel format is `420v` or
// `420f`. The value must be one of the strings specified in [Display Stream
// YCbCr to RGB conversion Matrix Options].
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/colorMatrix
//
// [Display Stream YCbCr to RGB conversion Matrix Options]: https://developer.apple.com/documentation/CoreGraphics/display-stream-ycbcr-to-rgb-conversion-matrix-options
func (s SCStreamConfiguration) ColorMatrix() corefoundation.CFStringRef {
	rv := objc.Send[corefoundation.CFStringRef](s.ID, objc.Sel("colorMatrix"))
	return corefoundation.CFStringRef(rv)
}
func (s SCStreamConfiguration) SetColorMatrix(value corefoundation.CFStringRef) {
	objc.Send[struct{}](s.ID, objc.Sel("setColorMatrix:"), value)
}

// A color space to use for the output buffer.
//
// # Discussion
//
// If you don’t specify a value, the output buffer uses the same color space
// as the display. If you specify a value, if must be one of the strings
// specified in [CGColorSpace].
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/colorSpaceName
//
// [CGColorSpace]: https://developer.apple.com/documentation/CoreGraphics/CGColorSpace
func (s SCStreamConfiguration) ColorSpaceName() corefoundation.CFStringRef {
	rv := objc.Send[corefoundation.CFStringRef](s.ID, objc.Sel("colorSpaceName"))
	return corefoundation.CFStringRef(rv)
}
func (s SCStreamConfiguration) SetColorSpaceName(value corefoundation.CFStringRef) {
	objc.Send[struct{}](s.ID, objc.Sel("setColorSpaceName:"), value)
}

// A background color for the output.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/backgroundColor
func (s SCStreamConfiguration) BackgroundColor() coregraphics.CGColorRef {
	rv := objc.Send[coregraphics.CGColorRef](s.ID, objc.Sel("backgroundColor"))
	return coregraphics.CGColorRef(rv)
}
func (s SCStreamConfiguration) SetBackgroundColor(value coregraphics.CGColorRef) {
	objc.Send[struct{}](s.ID, objc.Sel("setBackgroundColor:"), value)
}

// A Boolean value that determines whether the cursor is visible in the
// stream.
//
// # Discussion
//
// The cursor is visible by default.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/showsCursor
func (s SCStreamConfiguration) ShowsCursor() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("showsCursor"))
	return rv
}
func (s SCStreamConfiguration) SetShowsCursor(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setShowsCursor:"), value)
}

// A Boolean value that indicates if semitransparent content presents as
// opaque.
//
// # Discussion
//
// When this property is `true`, semitransparent content in the stream
// presents as backed by a solid white background, making the resulting image
// fully opaque. The default value is `false`.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/shouldBeOpaque
func (s SCStreamConfiguration) ShouldBeOpaque() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("shouldBeOpaque"))
	return rv
}
func (s SCStreamConfiguration) SetShouldBeOpaque(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setShouldBeOpaque:"), value)
}

// A Boolean value that indicates if the stream only captures shadows.
//
// # Discussion
//
// The default value is `false`.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/capturesShadowsOnly
func (s SCStreamConfiguration) CapturesShadowsOnly() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("capturesShadowsOnly"))
	return rv
}
func (s SCStreamConfiguration) SetCapturesShadowsOnly(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setCapturesShadowsOnly:"), value)
}

// A Boolean value that indicates if the stream ignores the capturing of
// window shadows when streaming in display style.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/ignoreShadowsDisplay
func (s SCStreamConfiguration) IgnoreShadowsDisplay() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("ignoreShadowsDisplay"))
	return rv
}
func (s SCStreamConfiguration) SetIgnoreShadowsDisplay(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setIgnoreShadowsDisplay:"), value)
}

// A Boolean value that indicates if the stream ignores the capturing of
// window shadows when streaming in window style.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/ignoreShadowsSingleWindow
func (s SCStreamConfiguration) IgnoreShadowsSingleWindow() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("ignoreShadowsSingleWindow"))
	return rv
}
func (s SCStreamConfiguration) SetIgnoreShadowsSingleWindow(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setIgnoreShadowsSingleWindow:"), value)
}

// A Boolean value that indicates if the stream ignores content clipped past
// the edge of a display, when streaming in display style.
//
// # Discussion
//
// The display that originates the stream determines clipping bounds. When
// this value is `true`, the stream contains content moved past the clipping
// bounds. The default value is `false`.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/ignoreGlobalClipDisplay
func (s SCStreamConfiguration) IgnoreGlobalClipDisplay() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("ignoreGlobalClipDisplay"))
	return rv
}
func (s SCStreamConfiguration) SetIgnoreGlobalClipDisplay(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setIgnoreGlobalClipDisplay:"), value)
}

// A Boolean value that indicates if the stream ignores content clipped past
// the edge of a display, when streaming in window style.
//
// # Discussion
//
// The display that originates the stream determines clipping bounds. When
// this value is `true`, the stream contains content moved past the clipping
// bounds. The default value is `false`.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/ignoreGlobalClipSingleWindow
func (s SCStreamConfiguration) IgnoreGlobalClipSingleWindow() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("ignoreGlobalClipSingleWindow"))
	return rv
}
func (s SCStreamConfiguration) SetIgnoreGlobalClipSingleWindow(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setIgnoreGlobalClipSingleWindow:"), value)
}

// The maximum number of frames for the queue to store.
//
// # Discussion
//
// By default, the system sets the queue depth to its minimum value of three
// frames. Specifying more frames uses more memory, but may allow you to
// process frame data without stalling the display stream.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/queueDepth
func (s SCStreamConfiguration) QueueDepth() int {
	rv := objc.Send[int](s.ID, objc.Sel("queueDepth"))
	return rv
}
func (s SCStreamConfiguration) SetQueueDepth(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setQueueDepth:"), value)
}

// The desired minimum time between frame updates, in seconds.
//
// # Discussion
//
// Use this value to throttle the rate at which you receive updates. The
// default value is `0`, which indicates that the system uses the maximum
// supported frame rate.
//
// You specify the minimum frame interval as the reciprocal of the maximum
// frame rate. For example, to configure the stream to capture at 60 fps,
// specify a minimum frame interval equal to `1/60`.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/minimumFrameInterval
func (s SCStreamConfiguration) MinimumFrameInterval() coremedia.CMTime {
	rv := objc.Send[coremedia.CMTime](s.ID, objc.Sel("minimumFrameInterval"))
	return coremedia.CMTime(rv)
}
func (s SCStreamConfiguration) SetMinimumFrameInterval(value coremedia.CMTime) {
	objc.Send[struct{}](s.ID, objc.Sel("setMinimumFrameInterval:"), value)
}

// The resolution at which to capture source content.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/captureResolution
func (s SCStreamConfiguration) CaptureResolution() SCCaptureResolutionType {
	rv := objc.Send[SCCaptureResolutionType](s.ID, objc.Sel("captureResolution"))
	return SCCaptureResolutionType(rv)
}
func (s SCStreamConfiguration) SetCaptureResolution(value SCCaptureResolutionType) {
	objc.Send[struct{}](s.ID, objc.Sel("setCaptureResolution:"), value)
}

// A Boolean value that indicates whether to capture audio.
//
// # Discussion
//
// A stream doesn’t capture audio by default. Set this value to true if you
// require audio capture.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/capturesAudio
func (s SCStreamConfiguration) CapturesAudio() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("capturesAudio"))
	return rv
}
func (s SCStreamConfiguration) SetCapturesAudio(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setCapturesAudio:"), value)
}

// The sample rate for audio capture.
//
// # Discussion
//
// The framework supports sample rates of `8000`, `16000`, `24000`, and
// `48000`. If you don’t specify a sample rate, or specify an unsupported
// value, the system uses a default sample rate of 48 kHz.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/sampleRate
func (s SCStreamConfiguration) SampleRate() int {
	rv := objc.Send[int](s.ID, objc.Sel("sampleRate"))
	return rv
}
func (s SCStreamConfiguration) SetSampleRate(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setSampleRate:"), value)
}

// The number of audio channels to capture.
//
// # Discussion
//
// The framework supports channel counts of `1` (mono) or `2` (stereo). If you
// don’t specify a channel count, or specify an unsupported value, the
// system defaults to stereo audio capture.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/channelCount
func (s SCStreamConfiguration) ChannelCount() int {
	rv := objc.Send[int](s.ID, objc.Sel("channelCount"))
	return rv
}
func (s SCStreamConfiguration) SetChannelCount(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setChannelCount:"), value)
}

// A Boolean value that indicates whether to exclude audio from your app
// during capture.
//
// # Discussion
//
// The default value is false. If you include your app process in the stream
// output, you can set this value to true to exclude its audio.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/excludesCurrentProcessAudio
func (s SCStreamConfiguration) ExcludesCurrentProcessAudio() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("excludesCurrentProcessAudio"))
	return rv
}
func (s SCStreamConfiguration) SetExcludesCurrentProcessAudio(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setExcludesCurrentProcessAudio:"), value)
}

// A name that you provide for identifying the stream.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/streamName
func (s SCStreamConfiguration) StreamName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("streamName"))
	return foundation.NSStringFromID(rv).String()
}
func (s SCStreamConfiguration) SetStreamName(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setStreamName:"), objc.String(value))
}

// A value indicating if alerts appear to presenters while using Presenter
// Overlay.
//
// # Discussion
//
// The default value is [SCPresenterOverlayAlertSettingSystem].
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/presenterOverlayPrivacyAlertSetting
func (s SCStreamConfiguration) PresenterOverlayPrivacyAlertSetting() SCPresenterOverlayAlertSetting {
	rv := objc.Send[SCPresenterOverlayAlertSetting](s.ID, objc.Sel("presenterOverlayPrivacyAlertSetting"))
	return SCPresenterOverlayAlertSetting(rv)
}
func (s SCStreamConfiguration) SetPresenterOverlayPrivacyAlertSetting(value SCPresenterOverlayAlertSetting) {
	objc.Send[struct{}](s.ID, objc.Sel("setPresenterOverlayPrivacyAlertSetting:"), value)
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/captureDynamicRange
func (s SCStreamConfiguration) CaptureDynamicRange() SCCaptureDynamicRange {
	rv := objc.Send[SCCaptureDynamicRange](s.ID, objc.Sel("captureDynamicRange"))
	return SCCaptureDynamicRange(rv)
}
func (s SCStreamConfiguration) SetCaptureDynamicRange(value SCCaptureDynamicRange) {
	objc.Send[struct{}](s.ID, objc.Sel("setCaptureDynamicRange:"), value)
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/captureMicrophone
func (s SCStreamConfiguration) CaptureMicrophone() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("captureMicrophone"))
	return rv
}
func (s SCStreamConfiguration) SetCaptureMicrophone(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setCaptureMicrophone:"), value)
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/includeChildWindows
func (s SCStreamConfiguration) IncludeChildWindows() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("includeChildWindows"))
	return rv
}
func (s SCStreamConfiguration) SetIncludeChildWindows(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setIncludeChildWindows:"), value)
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/microphoneCaptureDeviceID
func (s SCStreamConfiguration) MicrophoneCaptureDeviceID() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("microphoneCaptureDeviceID"))
	return foundation.NSStringFromID(rv).String()
}
func (s SCStreamConfiguration) SetMicrophoneCaptureDeviceID(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setMicrophoneCaptureDeviceID:"), objc.String(value))
}

// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCStreamConfiguration/showMouseClicks
func (s SCStreamConfiguration) ShowMouseClicks() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("showMouseClicks"))
	return rv
}
func (s SCStreamConfiguration) SetShowMouseClicks(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setShowMouseClicks:"), value)
}
