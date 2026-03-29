// Code generated from Apple documentation for ScreenCaptureKit. DO NOT EDIT.

package screencapturekit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SCContentFilter] class.
var (
	_SCContentFilterClass     SCContentFilterClass
	_SCContentFilterClassOnce sync.Once
)

func getSCContentFilterClass() SCContentFilterClass {
	_SCContentFilterClassOnce.Do(func() {
		_SCContentFilterClass = SCContentFilterClass{class: objc.GetClass("SCContentFilter")}
	})
	return _SCContentFilterClass
}

// GetSCContentFilterClass returns the class object for SCContentFilter.
func GetSCContentFilterClass() SCContentFilterClass {
	return getSCContentFilterClass()
}

type SCContentFilterClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SCContentFilterClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SCContentFilterClass) Alloc() SCContentFilter {
	rv := objc.Send[SCContentFilter](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// An instance that filters the content a stream captures.
//
// # Overview
// 
// Use a content filter to limit an [SCStream] object’s output to only that
// matching your filter criteria. Retrieve the displays, apps, and windows
// that your app can capture from an instance of [SCShareableContent].
//
// # Creating a filter
//
//   - [SCContentFilter.InitWithDesktopIndependentWindow]: Creates a filter that captures only the specified window.
//   - [SCContentFilter.InitWithDisplayIncludingWindows]: Creates a filter that captures only specific windows from a display.
//   - [SCContentFilter.InitWithDisplayExcludingWindows]: Creates a filter that captures the contents of a display, excluding the specified windows.
//   - [SCContentFilter.InitWithDisplayIncludingApplicationsExceptingWindows]: Creates a filter that captures a display, including only windows of the specified apps.
//   - [SCContentFilter.InitWithDisplayExcludingApplicationsExceptingWindows]: Creates a filter that captures a display, excluding windows of the specified apps.
//
// # Filter properties
//
//   - [SCContentFilter.ContentRect]: The size and location of the content to filter, in screen points.
//   - [SCContentFilter.PointPixelScale]: The scaling factor used to translate screen points into pixels.
//   - [SCContentFilter.StreamType]: The type of the streaming content.
//   - [SCContentFilter.Style]: The display style of the sharable content.
//
// # Instance Properties
//
//   - [SCContentFilter.IncludeMenuBar]
//   - [SCContentFilter.SetIncludeMenuBar]
//   - [SCContentFilter.IncludedApplications]
//   - [SCContentFilter.IncludedDisplays]
//   - [SCContentFilter.IncludedWindows]
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter
type SCContentFilter struct {
	objectivec.Object
}

// SCContentFilterFromID constructs a [SCContentFilter] from an objc.ID.
//
// An instance that filters the content a stream captures.
func SCContentFilterFromID(id objc.ID) SCContentFilter {
	return SCContentFilter{objectivec.Object{ID: id}}
}
// NOTE: SCContentFilter adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SCContentFilter] class.
//
// # Creating a filter
//
//   - [ISCContentFilter.InitWithDesktopIndependentWindow]: Creates a filter that captures only the specified window.
//   - [ISCContentFilter.InitWithDisplayIncludingWindows]: Creates a filter that captures only specific windows from a display.
//   - [ISCContentFilter.InitWithDisplayExcludingWindows]: Creates a filter that captures the contents of a display, excluding the specified windows.
//   - [ISCContentFilter.InitWithDisplayIncludingApplicationsExceptingWindows]: Creates a filter that captures a display, including only windows of the specified apps.
//   - [ISCContentFilter.InitWithDisplayExcludingApplicationsExceptingWindows]: Creates a filter that captures a display, excluding windows of the specified apps.
//
// # Filter properties
//
//   - [ISCContentFilter.ContentRect]: The size and location of the content to filter, in screen points.
//   - [ISCContentFilter.PointPixelScale]: The scaling factor used to translate screen points into pixels.
//   - [ISCContentFilter.StreamType]: The type of the streaming content.
//   - [ISCContentFilter.Style]: The display style of the sharable content.
//
// # Instance Properties
//
//   - [ISCContentFilter.IncludeMenuBar]
//   - [ISCContentFilter.SetIncludeMenuBar]
//   - [ISCContentFilter.IncludedApplications]
//   - [ISCContentFilter.IncludedDisplays]
//   - [ISCContentFilter.IncludedWindows]
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter
type ISCContentFilter interface {
	objectivec.IObject

	// Topic: Creating a filter

	// Creates a filter that captures only the specified window.
	InitWithDesktopIndependentWindow(window ISCWindow) SCContentFilter
	// Creates a filter that captures only specific windows from a display.
	InitWithDisplayIncludingWindows(display ISCDisplay, includedWindows []SCWindow) SCContentFilter
	// Creates a filter that captures the contents of a display, excluding the specified windows.
	InitWithDisplayExcludingWindows(display ISCDisplay, excluded []SCWindow) SCContentFilter
	// Creates a filter that captures a display, including only windows of the specified apps.
	InitWithDisplayIncludingApplicationsExceptingWindows(display ISCDisplay, applications []SCRunningApplication, exceptingWindows []SCWindow) SCContentFilter
	// Creates a filter that captures a display, excluding windows of the specified apps.
	InitWithDisplayExcludingApplicationsExceptingWindows(display ISCDisplay, applications []SCRunningApplication, exceptingWindows []SCWindow) SCContentFilter

	// Topic: Filter properties

	// The size and location of the content to filter, in screen points.
	ContentRect() corefoundation.CGRect
	// The scaling factor used to translate screen points into pixels.
	PointPixelScale() float32
	// The type of the streaming content.
	StreamType() SCStreamType
	// The display style of the sharable content.
	Style() SCShareableContentStyle

	// Topic: Instance Properties

	IncludeMenuBar() bool
	SetIncludeMenuBar(value bool)
	IncludedApplications() []SCRunningApplication
	IncludedDisplays() []SCDisplay
	IncludedWindows() []SCWindow
}

// Init initializes the instance.
func (c SCContentFilter) Init() SCContentFilter {
	rv := objc.Send[SCContentFilter](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c SCContentFilter) Autorelease() SCContentFilter {
	rv := objc.Send[SCContentFilter](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewSCContentFilter creates a new SCContentFilter instance.
func NewSCContentFilter() SCContentFilter {
	class := getSCContentFilterClass()
	rv := objc.Send[SCContentFilter](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a filter that captures only the specified window.
//
// window: A window to capture.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/init(desktopIndependentWindow:)
func NewContentFilterWithDesktopIndependentWindow(window ISCWindow) SCContentFilter {
	instance := getSCContentFilterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDesktopIndependentWindow:"), window)
	return SCContentFilterFromID(rv)
}

// Creates a filter that captures a display, excluding windows of the
// specified apps.
//
// display: A display to capture.
//
// applications: An array of apps to exclude from capture.
//
// exceptingWindows: An array of windows that are exceptions to the previous rules.
//
// # Discussion
// 
// The initializer arguments provide a three-stage filter that gives you
// fine-grained control over the output:
// 
// - Specify a display to capture. If you don’t specify additional filter
// criteria, the stream includes all content for a display. - Specify one or
// more apps with windows to exclude from the output. - Specify one or more
// windows that are exceptions to the previous rules. If the previous rules
// include a window, specifying it as an exception excludes it from the
// output. Likewise, if the previous rules exclude a window, specifying it as
// an exception includes it in the output.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/init(display:excludingApplications:exceptingWindows:)
func NewContentFilterWithDisplayExcludingApplicationsExceptingWindows(display ISCDisplay, applications []SCRunningApplication, exceptingWindows []SCWindow) SCContentFilter {
	instance := getSCContentFilterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisplay:excludingApplications:exceptingWindows:"), display, objectivec.IObjectSliceToNSArray(applications), objectivec.IObjectSliceToNSArray(exceptingWindows))
	return SCContentFilterFromID(rv)
}

// Creates a filter that captures the contents of a display, excluding the
// specified windows.
//
// display: A display to capture.
//
// excluded: An array of windows to exclude from the output.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/init(display:excludingWindows:)
func NewContentFilterWithDisplayExcludingWindows(display ISCDisplay, excluded []SCWindow) SCContentFilter {
	instance := getSCContentFilterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisplay:excludingWindows:"), display, objectivec.IObjectSliceToNSArray(excluded))
	return SCContentFilterFromID(rv)
}

// Creates a filter that captures a display, including only windows of the
// specified apps.
//
// display: A display to capture.
//
// applications: An array of apps with windows to capture.
//
// exceptingWindows: An array of windows that are exceptions to the previous rules.
//
// # Discussion
// 
// The initializer arguments provide a three-stage filter that gives you
// fine-grained control over the output:
// 
// - Specify a display to capture. If you don’t specify additional filter
// criteria, the stream includes all content for a display. - Specify one or
// more apps to capture only the windows that they own. - Specify one or more
// windows that are exceptions to the previous rules. If the previous rules
// include a window, specifying it as an exception excludes it from the
// output. Likewise, if the previous rules exclude a window, specifying it as
// an exception includes it in the output.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/init(display:including:exceptingWindows:)
func NewContentFilterWithDisplayIncludingApplicationsExceptingWindows(display ISCDisplay, applications []SCRunningApplication, exceptingWindows []SCWindow) SCContentFilter {
	instance := getSCContentFilterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisplay:includingApplications:exceptingWindows:"), display, objectivec.IObjectSliceToNSArray(applications), objectivec.IObjectSliceToNSArray(exceptingWindows))
	return SCContentFilterFromID(rv)
}

// Creates a filter that captures only specific windows from a display.
//
// display: A display to capture.
//
// includedWindows: An array of windows to include in the output.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/init(display:including:)
func NewContentFilterWithDisplayIncludingWindows(display ISCDisplay, includedWindows []SCWindow) SCContentFilter {
	instance := getSCContentFilterClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithDisplay:includingWindows:"), display, objectivec.IObjectSliceToNSArray(includedWindows))
	return SCContentFilterFromID(rv)
}

// Creates a filter that captures only the specified window.
//
// window: A window to capture.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/init(desktopIndependentWindow:)
func (c SCContentFilter) InitWithDesktopIndependentWindow(window ISCWindow) SCContentFilter {
	rv := objc.Send[SCContentFilter](c.ID, objc.Sel("initWithDesktopIndependentWindow:"), window)
	return rv
}
// Creates a filter that captures only specific windows from a display.
//
// display: A display to capture.
//
// includedWindows: An array of windows to include in the output.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/init(display:including:)
func (c SCContentFilter) InitWithDisplayIncludingWindows(display ISCDisplay, includedWindows []SCWindow) SCContentFilter {
	rv := objc.Send[SCContentFilter](c.ID, objc.Sel("initWithDisplay:includingWindows:"), display, objectivec.IObjectSliceToNSArray(includedWindows))
	return rv
}
// Creates a filter that captures the contents of a display, excluding the
// specified windows.
//
// display: A display to capture.
//
// excluded: An array of windows to exclude from the output.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/init(display:excludingWindows:)
func (c SCContentFilter) InitWithDisplayExcludingWindows(display ISCDisplay, excluded []SCWindow) SCContentFilter {
	rv := objc.Send[SCContentFilter](c.ID, objc.Sel("initWithDisplay:excludingWindows:"), display, objectivec.IObjectSliceToNSArray(excluded))
	return rv
}
// Creates a filter that captures a display, including only windows of the
// specified apps.
//
// display: A display to capture.
//
// applications: An array of apps with windows to capture.
//
// exceptingWindows: An array of windows that are exceptions to the previous rules.
//
// # Discussion
// 
// The initializer arguments provide a three-stage filter that gives you
// fine-grained control over the output:
// 
// - Specify a display to capture. If you don’t specify additional filter
// criteria, the stream includes all content for a display. - Specify one or
// more apps to capture only the windows that they own. - Specify one or more
// windows that are exceptions to the previous rules. If the previous rules
// include a window, specifying it as an exception excludes it from the
// output. Likewise, if the previous rules exclude a window, specifying it as
// an exception includes it in the output.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/init(display:including:exceptingWindows:)
func (c SCContentFilter) InitWithDisplayIncludingApplicationsExceptingWindows(display ISCDisplay, applications []SCRunningApplication, exceptingWindows []SCWindow) SCContentFilter {
	rv := objc.Send[SCContentFilter](c.ID, objc.Sel("initWithDisplay:includingApplications:exceptingWindows:"), display, objectivec.IObjectSliceToNSArray(applications), objectivec.IObjectSliceToNSArray(exceptingWindows))
	return rv
}
// Creates a filter that captures a display, excluding windows of the
// specified apps.
//
// display: A display to capture.
//
// applications: An array of apps to exclude from capture.
//
// exceptingWindows: An array of windows that are exceptions to the previous rules.
//
// # Discussion
// 
// The initializer arguments provide a three-stage filter that gives you
// fine-grained control over the output:
// 
// - Specify a display to capture. If you don’t specify additional filter
// criteria, the stream includes all content for a display. - Specify one or
// more apps with windows to exclude from the output. - Specify one or more
// windows that are exceptions to the previous rules. If the previous rules
// include a window, specifying it as an exception excludes it from the
// output. Likewise, if the previous rules exclude a window, specifying it as
// an exception includes it in the output.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/init(display:excludingApplications:exceptingWindows:)
func (c SCContentFilter) InitWithDisplayExcludingApplicationsExceptingWindows(display ISCDisplay, applications []SCRunningApplication, exceptingWindows []SCWindow) SCContentFilter {
	rv := objc.Send[SCContentFilter](c.ID, objc.Sel("initWithDisplay:excludingApplications:exceptingWindows:"), display, objectivec.IObjectSliceToNSArray(applications), objectivec.IObjectSliceToNSArray(exceptingWindows))
	return rv
}

// The size and location of the content to filter, in screen points.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/contentRect
func (c SCContentFilter) ContentRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](c.ID, objc.Sel("contentRect"))
	return corefoundation.CGRect(rv)
}
// The scaling factor used to translate screen points into pixels.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/pointPixelScale
func (c SCContentFilter) PointPixelScale() float32 {
	rv := objc.Send[float32](c.ID, objc.Sel("pointPixelScale"))
	return rv
}
// The type of the streaming content.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/streamType
func (c SCContentFilter) StreamType() SCStreamType {
	rv := objc.Send[SCStreamType](c.ID, objc.Sel("streamType"))
	return SCStreamType(rv)
}
// The display style of the sharable content.
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/style
func (c SCContentFilter) Style() SCShareableContentStyle {
	rv := objc.Send[SCShareableContentStyle](c.ID, objc.Sel("style"))
	return SCShareableContentStyle(rv)
}
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/includeMenuBar
func (c SCContentFilter) IncludeMenuBar() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("includeMenuBar"))
	return rv
}
func (c SCContentFilter) SetIncludeMenuBar(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setIncludeMenuBar:"), value)
}
//
// # Discussion
// 
// Applications that are included in the content filter
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/includedApplications
func (c SCContentFilter) IncludedApplications() []SCRunningApplication {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("includedApplications"))
	return objc.ConvertSlice(rv, func(id objc.ID) SCRunningApplication {
		return SCRunningApplicationFromID(id)
	})
}
//
// # Discussion
// 
// SCDisplays that are included in the content filter
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/includedDisplays
func (c SCContentFilter) IncludedDisplays() []SCDisplay {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("includedDisplays"))
	return objc.ConvertSlice(rv, func(id objc.ID) SCDisplay {
		return SCDisplayFromID(id)
	})
}
//
// # Discussion
// 
// Windows that are included in the content filter
//
// See: https://developer.apple.com/documentation/ScreenCaptureKit/SCContentFilter/includedWindows
func (c SCContentFilter) IncludedWindows() []SCWindow {
	rv := objc.Send[[]objc.ID](c.ID, objc.Sel("includedWindows"))
	return objc.ConvertSlice(rv, func(id objc.ID) SCWindow {
		return SCWindowFromID(id)
	})
}

