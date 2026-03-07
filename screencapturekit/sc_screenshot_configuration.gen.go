// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/uniformtypeidentifiers"
)

// The class instance for the [SCScreenshotConfiguration] class.
var (
	_SCScreenshotConfigurationClass     SCScreenshotConfigurationClass
	_SCScreenshotConfigurationClassOnce sync.Once
)

func getSCScreenshotConfigurationClass() SCScreenshotConfigurationClass {
	_SCScreenshotConfigurationClassOnce.Do(func() {
		_SCScreenshotConfigurationClass = SCScreenshotConfigurationClass{class: objc.GetClass("SCScreenshotConfiguration")}
	})
	return _SCScreenshotConfigurationClass
}

// GetSCScreenshotConfigurationClass returns the class object for SCScreenshotConfiguration.
func GetSCScreenshotConfigurationClass() SCScreenshotConfigurationClass {
	return getSCScreenshotConfigurationClass()
}

type SCScreenshotConfigurationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (sc SCScreenshotConfigurationClass) Alloc() SCScreenshotConfiguration {
	rv := objc.Send[SCScreenshotConfiguration](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}







// An object that contains screenshot properties such as output width, height,
// and image quality specifications.
//
// # Overview
// 
// [SCScreenshotConfiguration] provides a default image capture configuration
// for [SCScreenshotManager]. Only configure its properties if you need to
// customize the output. Additional options for customization include dynamic
// range settings, image reproduction optimizations, and ignoring user
// interface elements.
//
// # Instance Properties
//
//   - [SCScreenshotConfiguration.ContentType]: A uniform type identifier that specifies the screenshot’s file format; HEIC, JPEG, or PNG.
//   - [SCScreenshotConfiguration.SetContentType]
//   - [SCScreenshotConfiguration.DestinationRect]: A rectangle that specifies whether to output screenshots in a subset of the output image.
//   - [SCScreenshotConfiguration.SetDestinationRect]
//   - [SCScreenshotConfiguration.DisplayIntent]: Specifies whether the screen capture uses attributes of the local or canonical display.
//   - [SCScreenshotConfiguration.SetDisplayIntent]
//   - [SCScreenshotConfiguration.DynamicRange]: Specifies the type of image returned to the client; standard dynamic range, high dynamic range, or both.
//   - [SCScreenshotConfiguration.SetDynamicRange]
//   - [SCScreenshotConfiguration.FileURL]: Specifies the URL where the screenshot process saves the output.
//   - [SCScreenshotConfiguration.SetFileURL]
//   - [SCScreenshotConfiguration.Height]: An integer value that specifies the output height, measured in pixels.
//   - [SCScreenshotConfiguration.SetHeight]
//   - [SCScreenshotConfiguration.IgnoreClipping]: A Boolean value that specifies whether to ignore framing on windows when using content filters.
//   - [SCScreenshotConfiguration.SetIgnoreClipping]
//   - [SCScreenshotConfiguration.IgnoreShadows]: A Boolean value that specifies whether to ignore framing on windows.
//   - [SCScreenshotConfiguration.SetIgnoreShadows]
//   - [SCScreenshotConfiguration.IncludeChildWindows]: A Boolean that specifies whether the screenshot captures subwindows of the included apps and windows.
//   - [SCScreenshotConfiguration.SetIncludeChildWindows]
//   - [SCScreenshotConfiguration.ShowsCursor]: A Boolean value that specifies whether the pointer appears in the screenshot.
//   - [SCScreenshotConfiguration.SetShowsCursor]
//   - [SCScreenshotConfiguration.SourceRect]: A rectangle that specifies that the screenshot only samples a subset of the frame input.
//   - [SCScreenshotConfiguration.SetSourceRect]
//   - [SCScreenshotConfiguration.Width]: An integer value that specifies the output width in pixels.
//   - [SCScreenshotConfiguration.SetWidth]
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration
type SCScreenshotConfiguration struct {
	objectivec.Object
}

// SCScreenshotConfigurationFromID constructs a [SCScreenshotConfiguration] from an objc.ID.
//
// An object that contains screenshot properties such as output width, height,
// and image quality specifications.
func SCScreenshotConfigurationFromID(id objc.ID) SCScreenshotConfiguration {
	return SCScreenshotConfiguration{objectivec.Object{id}}
}
// NOTE: SCScreenshotConfiguration adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [SCScreenshotConfiguration] class.
//
// # Instance Properties
//
//   - [ISCScreenshotConfiguration.ContentType]: A uniform type identifier that specifies the screenshot’s file format; HEIC, JPEG, or PNG.
//   - [ISCScreenshotConfiguration.SetContentType]
//   - [ISCScreenshotConfiguration.DestinationRect]: A rectangle that specifies whether to output screenshots in a subset of the output image.
//   - [ISCScreenshotConfiguration.SetDestinationRect]
//   - [ISCScreenshotConfiguration.DisplayIntent]: Specifies whether the screen capture uses attributes of the local or canonical display.
//   - [ISCScreenshotConfiguration.SetDisplayIntent]
//   - [ISCScreenshotConfiguration.DynamicRange]: Specifies the type of image returned to the client; standard dynamic range, high dynamic range, or both.
//   - [ISCScreenshotConfiguration.SetDynamicRange]
//   - [ISCScreenshotConfiguration.FileURL]: Specifies the URL where the screenshot process saves the output.
//   - [ISCScreenshotConfiguration.SetFileURL]
//   - [ISCScreenshotConfiguration.Height]: An integer value that specifies the output height, measured in pixels.
//   - [ISCScreenshotConfiguration.SetHeight]
//   - [ISCScreenshotConfiguration.IgnoreClipping]: A Boolean value that specifies whether to ignore framing on windows when using content filters.
//   - [ISCScreenshotConfiguration.SetIgnoreClipping]
//   - [ISCScreenshotConfiguration.IgnoreShadows]: A Boolean value that specifies whether to ignore framing on windows.
//   - [ISCScreenshotConfiguration.SetIgnoreShadows]
//   - [ISCScreenshotConfiguration.IncludeChildWindows]: A Boolean that specifies whether the screenshot captures subwindows of the included apps and windows.
//   - [ISCScreenshotConfiguration.SetIncludeChildWindows]
//   - [ISCScreenshotConfiguration.ShowsCursor]: A Boolean value that specifies whether the pointer appears in the screenshot.
//   - [ISCScreenshotConfiguration.SetShowsCursor]
//   - [ISCScreenshotConfiguration.SourceRect]: A rectangle that specifies that the screenshot only samples a subset of the frame input.
//   - [ISCScreenshotConfiguration.SetSourceRect]
//   - [ISCScreenshotConfiguration.Width]: An integer value that specifies the output width in pixels.
//   - [ISCScreenshotConfiguration.SetWidth]
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration
type ISCScreenshotConfiguration interface {
	objectivec.IObject

	// Topic: Instance Properties

	// A uniform type identifier that specifies the screenshot’s file format; HEIC, JPEG, or PNG.
	ContentType() uniformtypeidentifiers.UTType
	SetContentType(value uniformtypeidentifiers.UTType)
	// A rectangle that specifies whether to output screenshots in a subset of the output image.
	DestinationRect() corefoundation.CGRect
	SetDestinationRect(value corefoundation.CGRect)
	// Specifies whether the screen capture uses attributes of the local or canonical display.
	DisplayIntent() SCScreenshotDisplayIntent
	SetDisplayIntent(value SCScreenshotDisplayIntent)
	// Specifies the type of image returned to the client; standard dynamic range, high dynamic range, or both.
	DynamicRange() SCScreenshotDynamicRange
	SetDynamicRange(value SCScreenshotDynamicRange)
	// Specifies the URL where the screenshot process saves the output.
	FileURL() foundation.INSURL
	SetFileURL(value foundation.INSURL)
	// An integer value that specifies the output height, measured in pixels.
	Height() int
	SetHeight(value int)
	// A Boolean value that specifies whether to ignore framing on windows when using content filters.
	IgnoreClipping() bool
	SetIgnoreClipping(value bool)
	// A Boolean value that specifies whether to ignore framing on windows.
	IgnoreShadows() bool
	SetIgnoreShadows(value bool)
	// A Boolean that specifies whether the screenshot captures subwindows of the included apps and windows.
	IncludeChildWindows() bool
	SetIncludeChildWindows(value bool)
	// A Boolean value that specifies whether the pointer appears in the screenshot.
	ShowsCursor() bool
	SetShowsCursor(value bool)
	// A rectangle that specifies that the screenshot only samples a subset of the frame input.
	SourceRect() corefoundation.CGRect
	SetSourceRect(value corefoundation.CGRect)
	// An integer value that specifies the output width in pixels.
	Width() int
	SetWidth(value int)
}





// Init initializes the instance.
func (s SCScreenshotConfiguration) Init() SCScreenshotConfiguration {
	rv := objc.Send[SCScreenshotConfiguration](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SCScreenshotConfiguration) Autorelease() SCScreenshotConfiguration {
	rv := objc.Send[SCScreenshotConfiguration](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCScreenshotConfiguration creates a new SCScreenshotConfiguration instance.
func NewSCScreenshotConfiguration() SCScreenshotConfiguration {
	class := getSCScreenshotConfigurationClass()
	rv := objc.Send[SCScreenshotConfiguration](objc.ID(class.class), objc.Sel("new"))
	return rv
}





















// A uniform type identifier that specifies the screenshot’s file format;
// HEIC, JPEG, or PNG.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/contentType
func (s SCScreenshotConfiguration) ContentType() uniformtypeidentifiers.UTType {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("contentType"))
	return uniformtypeidentifiers.UTTypeFromID(objc.ID(rv))
}
func (s SCScreenshotConfiguration) SetContentType(value uniformtypeidentifiers.UTType) {
	objc.Send[struct{}](s.ID, objc.Sel("setContentType:"), value)
}



// A rectangle that specifies whether to output screenshots in a subset of the
// output image.
//
// # Discussion
// 
// If you don’t specify a destination rectangle, the system uses the full
// dimensions of the output surface. The rectangle is specified in pixels in
// the display’s coordinate system.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/destinationRect
func (s SCScreenshotConfiguration) DestinationRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("destinationRect"))
	return corefoundation.CGRect(rv)
}
func (s SCScreenshotConfiguration) SetDestinationRect(value corefoundation.CGRect) {
	objc.Send[struct{}](s.ID, objc.Sel("setDestinationRect:"), value)
}



// Specifies whether the screen capture uses attributes of the local or
// canonical display.
//
// # Discussion
// 
// Performing a screenshot with either the local or canonical display
// attributes optimizes output for presentation on either the capture display
// or any high dynamic range display respectively.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/displayIntent-swift.property
func (s SCScreenshotConfiguration) DisplayIntent() SCScreenshotDisplayIntent {
	rv := objc.Send[SCScreenshotDisplayIntent](s.ID, objc.Sel("displayIntent"))
	return SCScreenshotDisplayIntent(rv)
}
func (s SCScreenshotConfiguration) SetDisplayIntent(value SCScreenshotDisplayIntent) {
	objc.Send[struct{}](s.ID, objc.Sel("setDisplayIntent:"), value)
}



// Specifies the type of image returned to the client; standard dynamic range,
// high dynamic range, or both.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/dynamicRange-swift.property
func (s SCScreenshotConfiguration) DynamicRange() SCScreenshotDynamicRange {
	rv := objc.Send[SCScreenshotDynamicRange](s.ID, objc.Sel("dynamicRange"))
	return SCScreenshotDynamicRange(rv)
}
func (s SCScreenshotConfiguration) SetDynamicRange(value SCScreenshotDynamicRange) {
	objc.Send[struct{}](s.ID, objc.Sel("setDynamicRange:"), value)
}



// Specifies the URL where the screenshot process saves the output.
//
// # Discussion
// 
// If `imageOutputURL` is `nil`, then the file isn’t saved.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/fileURL
func (s SCScreenshotConfiguration) FileURL() foundation.INSURL {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("fileURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (s SCScreenshotConfiguration) SetFileURL(value foundation.INSURL) {
	objc.Send[struct{}](s.ID, objc.Sel("setFileURL:"), value)
}



// An integer value that specifies the output height, measured in pixels.
//
// # Discussion
// 
// The default value is the height of the captured content.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/height
func (s SCScreenshotConfiguration) Height() int {
	rv := objc.Send[int](s.ID, objc.Sel("height"))
	return rv
}
func (s SCScreenshotConfiguration) SetHeight(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setHeight:"), value)
}



// A Boolean value that specifies whether to ignore framing on windows when
// using content filters.
//
// # Discussion
// 
// Use [SCNPropertyFilter] in conjunction with this property to ignore window
// framing on specified apps and windows. Setting this value to `true` ignores
// shadows.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/ignoreClipping
func (s SCScreenshotConfiguration) IgnoreClipping() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("ignoreClipping"))
	return rv
}
func (s SCScreenshotConfiguration) SetIgnoreClipping(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setIgnoreClipping:"), value)
}



// A Boolean value that specifies whether to ignore framing on windows.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/ignoreShadows
func (s SCScreenshotConfiguration) IgnoreShadows() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("ignoreShadows"))
	return rv
}
func (s SCScreenshotConfiguration) SetIgnoreShadows(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setIgnoreShadows:"), value)
}



// A Boolean that specifies whether the screenshot captures subwindows of the
// included apps and windows.
//
// # Discussion
// 
// By default taking a screenshot captures subwindows. For example, alerts,
// popovers, and sheets are captured by default.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/includeChildWindows
func (s SCScreenshotConfiguration) IncludeChildWindows() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("includeChildWindows"))
	return rv
}
func (s SCScreenshotConfiguration) SetIncludeChildWindows(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setIncludeChildWindows:"), value)
}



// A Boolean value that specifies whether the pointer appears in the
// screenshot.
//
// # Discussion
// 
// By default the pointer is visible in screenshots.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/showsCursor
func (s SCScreenshotConfiguration) ShowsCursor() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("showsCursor"))
	return rv
}
func (s SCScreenshotConfiguration) SetShowsCursor(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setShowsCursor:"), value)
}



// A rectangle that specifies that the screenshot only samples a subset of the
// frame input.
//
// # Discussion
// 
// If not set, the screenshot captures the entire frame. Specify the rectangle
// in points in the display’s logical coordinate system.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/sourceRect
func (s SCScreenshotConfiguration) SourceRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("sourceRect"))
	return corefoundation.CGRect(rv)
}
func (s SCScreenshotConfiguration) SetSourceRect(value corefoundation.CGRect) {
	objc.Send[struct{}](s.ID, objc.Sel("setSourceRect:"), value)
}



// An integer value that specifies the output width in pixels.
//
// # Discussion
// 
// The default value is the width of the captured content.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/width
func (s SCScreenshotConfiguration) Width() int {
	rv := objc.Send[int](s.ID, objc.Sel("width"))
	return rv
}
func (s SCScreenshotConfiguration) SetWidth(value int) {
	objc.Send[struct{}](s.ID, objc.Sel("setWidth:"), value)
}







// An array of uniform type identifiers that correspond to file formats the
// output image supports.
//
// # Discussion
// 
// You can save the output [CGImage] into HEIC, JPEG, and PNG formats.
//
// [CGImage]: https://developer.apple.com/documentation/CoreGraphics/CGImage
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCScreenshotConfiguration/supportedContentTypes
func (_SCScreenshotConfigurationClass SCScreenshotConfigurationClass) SupportedContentTypes() []uniformtypeidentifiers.UTType {
	rv := objc.Send[[]objc.ID](objc.ID(_SCScreenshotConfigurationClass.class), objc.Sel("supportedContentTypes"))
	return objc.ConvertSlice(rv, func(id objc.ID) uniformtypeidentifiers.UTType {
		return uniformtypeidentifiers.UTTypeFromID(id)
	})
}



















