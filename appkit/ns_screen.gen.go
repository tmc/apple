// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/quartzcore"
)

// The class instance for the [NSScreen] class.
var (
	_NSScreenClass     NSScreenClass
	_NSScreenClassOnce sync.Once
)

func getNSScreenClass() NSScreenClass {
	_NSScreenClassOnce.Do(func() {
		_NSScreenClass = NSScreenClass{class: objc.GetClass("NSScreen")}
	})
	return _NSScreenClass
}

// GetNSScreenClass returns the class object for NSScreen.
func GetNSScreenClass() NSScreenClass {
	return getNSScreenClass()
}

type NSScreenClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSScreenClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScreenClass) Alloc() NSScreen {
	rv := objc.Send[NSScreen](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that describes the attributes of a computer’s monitor or
// screen.
//
// # Overview
// 
// An app may use an [NSScreen] object to retrieve information about a screen
// and use this information to decide what to display on that screen. For
// example, an app may use the [NSScreen.DeepestScreen] method to find out which of the
// available screens can best represent color and then might choose to display
// all of its windows on that screen.
// 
// Create the application object before you use the methods in this class, so
// that the application object can make the necessary connection to the window
// system. You can make sure the application object exists by invoking the
// [SharedApplication] method of [NSApplication]. If you created your app with
// Xcode, the application object is automatically created for you during
// initialization.
//
// # Getting Screen Information
//
//   - [NSScreen.Depth]: The current bit depth and colorspace information of the screen.
//   - [NSScreen.Frame]: The dimensions and location of the screen.
//   - [NSScreen.SupportedWindowDepths]: A zero-terminated array of the window depths supported by the screen.
//   - [NSScreen.DeviceDescription]: The device dictionary for the screen.
//   - [NSScreen.ColorSpace]: The color space of the screen.
//   - [NSScreen.LocalizedName]: The localized name of the display.
//   - [NSScreen.CanRepresentDisplayGamut]: A Boolean value indicating whether the color space of the screen is capable of representing the specified display gamut.
//
// # Converting Between Screen and Backing Coordinates
//
//   - [NSScreen.BackingAlignedRectOptions]: Converts a rectangle in global screen coordinates to a pixel aligned rectangle.
//   - [NSScreen.BackingScaleFactor]: The backing store pixel scale factor for the screen.
//   - [NSScreen.ConvertRectFromBacking]: Converts the rectangle from the device pixel aligned coordinates system of a screen.
//   - [NSScreen.ConvertRectToBacking]: Converts the rectangle to the device pixel aligned coordinates system of a screen.
//
// # Getting the Visible Portion of the Screen
//
//   - [NSScreen.VisibleFrame]: The current location and dimensions of the visible screen.
//   - [NSScreen.SafeAreaInsets]: The distances from the screen’s edges at which content isn’t obscured.
//   - [NSScreen.AuxiliaryTopLeftArea]: The unobscured portion of the top-left corner of the screen.
//   - [NSScreen.SetAuxiliaryTopLeftArea]
//   - [NSScreen.AuxiliaryTopRightArea]: The unobscured portion of the top-right corner of the screen.
//   - [NSScreen.SetAuxiliaryTopRightArea]
//
// # Getting Extended Dynamic Range Details
//
//   - [NSScreen.MaximumPotentialExtendedDynamicRangeColorComponentValue]: The maximum possible color component value for the screen when it’s in extended dynamic range (EDR) mode.
//   - [NSScreen.MaximumExtendedDynamicRangeColorComponentValue]: The current maximum color component value for the screen.
//   - [NSScreen.MaximumReferenceExtendedDynamicRangeColorComponentValue]: The current maximum color component value for reference rendering to the screen.
//
// # Getting Variable Refresh Rate Details
//
//   - [NSScreen.MaximumFramesPerSecond]: The maximum number of frames per second that the screen supports.
//   - [NSScreen.MinimumRefreshInterval]: The shortest refresh interval that the screen supports.
//   - [NSScreen.MaximumRefreshInterval]: The largest refresh interval that the screen supports.
//   - [NSScreen.DisplayUpdateGranularity]: The number of seconds between the screen’s supported update rates, for screens that support fixed update rates.
//   - [NSScreen.LastDisplayUpdateTimestamp]: The time of the last framebuffer update, expressed as the number of seconds since system startup.
//
// # Synchronizing with the display’s refresh rate
//
//   - [NSScreen.DisplayLinkWithTargetSelector]
//
// # Instance Properties
//
//   - [NSScreen.CgDirectDisplayID]: The CGDirectDisplayID for this screen. This will return nil if there isn’t one and will never return kCGNullDirectDisplay.
//   - [NSScreen.SetCgDirectDisplayID]
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen
type NSScreen struct {
	objectivec.Object
}

// NSScreenFromID constructs a [NSScreen] from an objc.ID.
//
// An object that describes the attributes of a computer’s monitor or
// screen.
func NSScreenFromID(id objc.ID) NSScreen {
	return NSScreen{objectivec.Object{ID: id}}
}
// NOTE: NSScreen adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSScreen] class.
//
// # Getting Screen Information
//
//   - [INSScreen.Depth]: The current bit depth and colorspace information of the screen.
//   - [INSScreen.Frame]: The dimensions and location of the screen.
//   - [INSScreen.SupportedWindowDepths]: A zero-terminated array of the window depths supported by the screen.
//   - [INSScreen.DeviceDescription]: The device dictionary for the screen.
//   - [INSScreen.ColorSpace]: The color space of the screen.
//   - [INSScreen.LocalizedName]: The localized name of the display.
//   - [INSScreen.CanRepresentDisplayGamut]: A Boolean value indicating whether the color space of the screen is capable of representing the specified display gamut.
//
// # Converting Between Screen and Backing Coordinates
//
//   - [INSScreen.BackingAlignedRectOptions]: Converts a rectangle in global screen coordinates to a pixel aligned rectangle.
//   - [INSScreen.BackingScaleFactor]: The backing store pixel scale factor for the screen.
//   - [INSScreen.ConvertRectFromBacking]: Converts the rectangle from the device pixel aligned coordinates system of a screen.
//   - [INSScreen.ConvertRectToBacking]: Converts the rectangle to the device pixel aligned coordinates system of a screen.
//
// # Getting the Visible Portion of the Screen
//
//   - [INSScreen.VisibleFrame]: The current location and dimensions of the visible screen.
//   - [INSScreen.SafeAreaInsets]: The distances from the screen’s edges at which content isn’t obscured.
//   - [INSScreen.AuxiliaryTopLeftArea]: The unobscured portion of the top-left corner of the screen.
//   - [INSScreen.SetAuxiliaryTopLeftArea]
//   - [INSScreen.AuxiliaryTopRightArea]: The unobscured portion of the top-right corner of the screen.
//   - [INSScreen.SetAuxiliaryTopRightArea]
//
// # Getting Extended Dynamic Range Details
//
//   - [INSScreen.MaximumPotentialExtendedDynamicRangeColorComponentValue]: The maximum possible color component value for the screen when it’s in extended dynamic range (EDR) mode.
//   - [INSScreen.MaximumExtendedDynamicRangeColorComponentValue]: The current maximum color component value for the screen.
//   - [INSScreen.MaximumReferenceExtendedDynamicRangeColorComponentValue]: The current maximum color component value for reference rendering to the screen.
//
// # Getting Variable Refresh Rate Details
//
//   - [INSScreen.MaximumFramesPerSecond]: The maximum number of frames per second that the screen supports.
//   - [INSScreen.MinimumRefreshInterval]: The shortest refresh interval that the screen supports.
//   - [INSScreen.MaximumRefreshInterval]: The largest refresh interval that the screen supports.
//   - [INSScreen.DisplayUpdateGranularity]: The number of seconds between the screen’s supported update rates, for screens that support fixed update rates.
//   - [INSScreen.LastDisplayUpdateTimestamp]: The time of the last framebuffer update, expressed as the number of seconds since system startup.
//
// # Synchronizing with the display’s refresh rate
//
//   - [INSScreen.DisplayLinkWithTargetSelector]
//
// # Instance Properties
//
//   - [INSScreen.CgDirectDisplayID]: The CGDirectDisplayID for this screen. This will return nil if there isn’t one and will never return kCGNullDirectDisplay.
//   - [INSScreen.SetCgDirectDisplayID]
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen
type INSScreen interface {
	objectivec.IObject

	// Topic: Getting Screen Information

	// The current bit depth and colorspace information of the screen.
	Depth() NSWindowDepth
	// The dimensions and location of the screen.
	Frame() corefoundation.CGRect
	// A zero-terminated array of the window depths supported by the screen.
	SupportedWindowDepths() NSWindowDepth
	// The device dictionary for the screen.
	DeviceDescription() foundation.INSDictionary
	// The color space of the screen.
	ColorSpace() INSColorSpace
	// The localized name of the display.
	LocalizedName() string
	// A Boolean value indicating whether the color space of the screen is capable of representing the specified display gamut.
	CanRepresentDisplayGamut(displayGamut NSDisplayGamut) bool

	// Topic: Converting Between Screen and Backing Coordinates

	// Converts a rectangle in global screen coordinates to a pixel aligned rectangle.
	BackingAlignedRectOptions(rect corefoundation.CGRect, options foundation.AlignmentOptions) corefoundation.CGRect
	// The backing store pixel scale factor for the screen.
	BackingScaleFactor() float64
	// Converts the rectangle from the device pixel aligned coordinates system of a screen.
	ConvertRectFromBacking(rect corefoundation.CGRect) corefoundation.CGRect
	// Converts the rectangle to the device pixel aligned coordinates system of a screen.
	ConvertRectToBacking(rect corefoundation.CGRect) corefoundation.CGRect

	// Topic: Getting the Visible Portion of the Screen

	// The current location and dimensions of the visible screen.
	VisibleFrame() corefoundation.CGRect
	// The distances from the screen’s edges at which content isn’t obscured.
	SafeAreaInsets() foundation.NSEdgeInsets
	// The unobscured portion of the top-left corner of the screen.
	AuxiliaryTopLeftArea() corefoundation.CGRect
	SetAuxiliaryTopLeftArea(value corefoundation.CGRect)
	// The unobscured portion of the top-right corner of the screen.
	AuxiliaryTopRightArea() corefoundation.CGRect
	SetAuxiliaryTopRightArea(value corefoundation.CGRect)

	// Topic: Getting Extended Dynamic Range Details

	// The maximum possible color component value for the screen when it’s in extended dynamic range (EDR) mode.
	MaximumPotentialExtendedDynamicRangeColorComponentValue() float64
	// The current maximum color component value for the screen.
	MaximumExtendedDynamicRangeColorComponentValue() float64
	// The current maximum color component value for reference rendering to the screen.
	MaximumReferenceExtendedDynamicRangeColorComponentValue() float64

	// Topic: Getting Variable Refresh Rate Details

	// The maximum number of frames per second that the screen supports.
	MaximumFramesPerSecond() int
	// The shortest refresh interval that the screen supports.
	MinimumRefreshInterval() float64
	// The largest refresh interval that the screen supports.
	MaximumRefreshInterval() float64
	// The number of seconds between the screen’s supported update rates, for screens that support fixed update rates.
	DisplayUpdateGranularity() float64
	// The time of the last framebuffer update, expressed as the number of seconds since system startup.
	LastDisplayUpdateTimestamp() float64

	// Topic: Synchronizing with the display’s refresh rate

	DisplayLinkWithTargetSelector(target objectivec.IObject, selector objc.SEL) quartzcore.CADisplayLink

	// Topic: Instance Properties

	// The CGDirectDisplayID for this screen. This will return nil if there isn’t one and will never return kCGNullDirectDisplay.
	CgDirectDisplayID() uint32
	SetCgDirectDisplayID(value uint32)
}

// Init initializes the instance.
func (s NSScreen) Init() NSScreen {
	rv := objc.Send[NSScreen](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScreen) Autorelease() NSScreen {
	rv := objc.Send[NSScreen](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScreen creates a new NSScreen instance.
func NewNSScreen() NSScreen {
	class := getNSScreenClass()
	rv := objc.Send[NSScreen](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// A Boolean value indicating whether the color space of the screen is capable
// of representing the specified display gamut.
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/canRepresent(_:)
func (s NSScreen) CanRepresentDisplayGamut(displayGamut NSDisplayGamut) bool {
	rv := objc.Send[bool](s.ID, objc.Sel("canRepresentDisplayGamut:"), displayGamut)
	return rv
}
// Converts a rectangle in global screen coordinates to a pixel aligned
// rectangle.
//
// rect: The input rectangle in global screen coordinates.
//
// options: Specifies the alignment options. See [AlignmentOptions] for possible
// values.
// //
// [AlignmentOptions]: https://developer.apple.com/documentation/Foundation/AlignmentOptions
//
// # Return Value
// 
// Returns a a pixel aligned rectangle on the target screen from the given
// input rectangle in global screen coordinates.
//
// # Discussion
// 
// This method uses [NSIntegralRectWithOptions(_:_:)] to produce the pixel
// aligned rectangle.
//
// [NSIntegralRectWithOptions(_:_:)]: https://developer.apple.com/documentation/Foundation/NSIntegralRectWithOptions(_:_:)
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/backingAlignedRect(_:options:)
func (s NSScreen) BackingAlignedRectOptions(rect corefoundation.CGRect, options foundation.AlignmentOptions) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("backingAlignedRect:options:"), rect, options)
	return corefoundation.CGRect(rv)
}
// Converts the rectangle from the device pixel aligned coordinates system of
// a screen.
//
// rect: The rectangle.
//
// # Return Value
// 
// The rectangle converted from the device pixel aligned coordinates system of
// the screen.
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/convertRectFromBacking(_:)
func (s NSScreen) ConvertRectFromBacking(rect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("convertRectFromBacking:"), rect)
	return corefoundation.CGRect(rv)
}
// Converts the rectangle to the device pixel aligned coordinates system of a
// screen.
//
// rect: The rectangle.
//
// # Return Value
// 
// The rectangle converted to the device pixel aligned coordinates system of
// the screen.
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/convertRectToBacking(_:)
func (s NSScreen) ConvertRectToBacking(rect corefoundation.CGRect) corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("convertRectToBacking:"), rect)
	return corefoundation.CGRect(rv)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/displayLink(target:selector:)
func (s NSScreen) DisplayLinkWithTargetSelector(target objectivec.IObject, selector objc.SEL) quartzcore.CADisplayLink {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("displayLinkWithTarget:selector:"), target, selector)
	return quartzcore.CADisplayLinkFromID(rv)
}

// The current bit depth and colorspace information of the screen.
//
// # Discussion
// 
// This value cannot be used directly. You must pass it to a function such as
// [NSBitsPerPixelFromDepth] or [NSColorSpaceFromDepth] to obtain a concrete
// value for the desired information.
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/depth
func (s NSScreen) Depth() NSWindowDepth {
	rv := objc.Send[NSWindowDepth](s.ID, objc.Sel("depth"))
	return NSWindowDepth(rv)
}
// The dimensions and location of the screen.
//
// # Discussion
// 
// This is the full screen rectangle at the current resolution. This rectangle
// includes any space currently occupied by the menu bar and dock.
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/frame
func (s NSScreen) Frame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("frame"))
	return corefoundation.CGRect(rv)
}
// A zero-terminated array of the window depths supported by the screen.
//
// # Discussion
// 
// This is a C-style array of window depths. The returned values cannot be
// used directly. You must pass each one to a function such as
// [NSBitsPerPixelFromDepth] or [NSColorSpaceFromDepth] to obtain a concrete
// value for the desired screen.
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/supportedWindowDepths
func (s NSScreen) SupportedWindowDepths() NSWindowDepth {
	rv := objc.Send[NSWindowDepth](s.ID, objc.Sel("supportedWindowDepths"))
	return NSWindowDepth(rv)
}
// The device dictionary for the screen.
//
// # Discussion
// 
// This is a dictionary containing the attributes of the receiver’s screen.
// For the list of keys you can use to retrieve values from the returned
// dictionary, see `Display Device—Descriptions`.
// 
// In addition to the display device constants described in [NSWindow], you
// can also retrieve the [CGDirectDisplayID] value associated with the screen
// from this dictionary. To access this value, specify the Objective-C string
// `@"NSScreenNumber"` as the key when requesting the item from the
// dictionary. The value associated with this key is an [NSNumber] object
// containing the display ID value. This string is only valid when used as a
// key for the dictionary returned by this method.
//
// [CGDirectDisplayID]: https://developer.apple.com/documentation/CoreGraphics/CGDirectDisplayID
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/deviceDescription
func (s NSScreen) DeviceDescription() foundation.INSDictionary {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("deviceDescription"))
	return foundation.NSDictionaryFromID(objc.ID(rv))
}
// The color space of the screen.
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/colorSpace
func (s NSScreen) ColorSpace() INSColorSpace {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("colorSpace"))
	return NSColorSpaceFromID(objc.ID(rv))
}
// The localized name of the display.
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/localizedName
func (s NSScreen) LocalizedName() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("localizedName"))
	return foundation.NSStringFromID(rv).String()
}
// The backing store pixel scale factor for the screen.
//
// # Discussion
// 
// This is the scale factor representing the number of backing store pixels
// corresponding to each linear unit in screen space on this screen.
// 
// This method is provided for rare cases when the explicit scale factor is
// needed. As often as possible, you should use the [NSView] class’s convert
// backing methods.
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/backingScaleFactor
func (s NSScreen) BackingScaleFactor() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("backingScaleFactor"))
	return rv
}
// The current location and dimensions of the visible screen.
//
// # Discussion
// 
// This rectangle defines the portion of the screen in which it is currently
// safe to draw your app’s content.
// 
// The returned rectangle is always based on the current user-interface
// settings and does not include the area currently occupied by the dock and
// menu bar. On Macs that include a camera housing in the bezel, this
// rectangle does not include the bezel or the visible areas on either side of
// the bezel. Get those areas using the [auxiliaryTopLeftArea] and
// [auxiliaryTopRightArea] properties. Don’t cache the rectangle in this
// property; it is based on the current user-interface settings, which can
// change between calls.
//
// [auxiliaryTopLeftArea]: https://developer.apple.com/documentation/AppKit/NSScreen/auxiliaryTopLeftArea-uglc
// [auxiliaryTopRightArea]: https://developer.apple.com/documentation/AppKit/NSScreen/auxiliaryTopRightArea-gr2n
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/visibleFrame
func (s NSScreen) VisibleFrame() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("visibleFrame"))
	return corefoundation.CGRect(rv)
}
// The distances from the screen’s edges at which content isn’t obscured.
//
// # Discussion
// 
// The safe area reflects the unobscured portion of the screen. On some Macs,
// the insets reflect the portion of the screen covered by the camera housing.
// If your app offers a custom full-screen experience, apply the specified
// insets to the screen’s frame rectangle to obtain the area within which it
// is safe to display your content. Content in the safe area is guaranteed to
// be unobscured.
// 
// If your app uses the system’s full-screen experience, you don’t need to
// account for the safe area in your window. When you call your window’s
// [ToggleFullScreen] method to enter full-screen mode, the system
// automatically positions the window’s contents within the safe area.
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/safeAreaInsets
func (s NSScreen) SafeAreaInsets() foundation.NSEdgeInsets {
	rv := objc.Send[foundation.NSEdgeInsets](s.ID, objc.Sel("safeAreaInsets"))
	return foundation.NSEdgeInsets(rv)
}
// The unobscured portion of the top-left corner of the screen.
//
// See: https://developer.apple.com/documentation/appkit/nsscreen/auxiliarytopleftarea-uglc
func (s NSScreen) AuxiliaryTopLeftArea() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("auxiliaryTopLeftArea"))
	return corefoundation.CGRect(rv)
}
func (s NSScreen) SetAuxiliaryTopLeftArea(value corefoundation.CGRect) {
	objc.Send[struct{}](s.ID, objc.Sel("setAuxiliaryTopLeftArea:"), value)
}
// The unobscured portion of the top-right corner of the screen.
//
// See: https://developer.apple.com/documentation/appkit/nsscreen/auxiliarytoprightarea-gr2n
func (s NSScreen) AuxiliaryTopRightArea() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](s.ID, objc.Sel("auxiliaryTopRightArea"))
	return corefoundation.CGRect(rv)
}
func (s NSScreen) SetAuxiliaryTopRightArea(value corefoundation.CGRect) {
	objc.Send[struct{}](s.ID, objc.Sel("setAuxiliaryTopRightArea:"), value)
}
// The maximum possible color component value for the screen when it’s in
// extended dynamic range (EDR) mode.
//
// # Discussion
// 
// The value of this property is determined when you create the [NSScreen]
// object, and doesn’t change afterwards. If this property is greater than
// `1.0`, the screen supports EDR values. If the screen doesn’t support EDR
// values, the value is `1.0`.
// 
// The actual maximum value might be lower than this property’s value, and
// can change dynamically, depending on the capabilities of the display
// hardware and other conditions. See
// [MaximumExtendedDynamicRangeColorComponentValue].
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/maximumPotentialExtendedDynamicRangeColorComponentValue
func (s NSScreen) MaximumPotentialExtendedDynamicRangeColorComponentValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("maximumPotentialExtendedDynamicRangeColorComponentValue"))
	return rv
}
// The current maximum color component value for the screen.
//
// # Discussion
// 
// If no content on screen provides extended dynamic range (EDR) values, the
// value of this property is `1.0`. If any content onscreen has requested EDR,
// the value may be greater than `1.0`, depending on the capabilities of the
// display hardware and other conditions. Only rendering contexts that support
// EDR can use values greater than `1.0`.
// 
// When the value changes, [didChangeScreenParametersNotification] is posted.
//
// [didChangeScreenParametersNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/didChangeScreenParametersNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/maximumExtendedDynamicRangeColorComponentValue
func (s NSScreen) MaximumExtendedDynamicRangeColorComponentValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("maximumExtendedDynamicRangeColorComponentValue"))
	return rv
}
// The current maximum color component value for reference rendering to the
// screen.
//
// # Discussion
// 
// Reference displays are calibrated to provide accurate color and lighting
// information that helps to optimize video content. Not all displays support
// reference rendering. If the display hardware doesn’t support reference
// rendering, the value of this property is `0`.
// 
// On reference displays, if you constrain pixel component values to values
// between `0` and [MaximumReferenceExtendedDynamicRangeColorComponentValue],
// the display hardware doesn’t apply any additional tone mapping to the
// pixels before rendering them. If you use values above this range, display
// hardware may adjust content to fit into its dynamic range.
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/maximumReferenceExtendedDynamicRangeColorComponentValue
func (s NSScreen) MaximumReferenceExtendedDynamicRangeColorComponentValue() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("maximumReferenceExtendedDynamicRangeColorComponentValue"))
	return rv
}
// The maximum number of frames per second that the screen supports.
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/maximumFramesPerSecond
func (s NSScreen) MaximumFramesPerSecond() int {
	rv := objc.Send[int](s.ID, objc.Sel("maximumFramesPerSecond"))
	return rv
}
// The shortest refresh interval that the screen supports.
//
// # Discussion
// 
// This interval represents the minimum amount of time, in seconds, your app
// has to generate new frames. It corresponds to the highest refresh rate of
// the display.
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/minimumRefreshInterval
func (s NSScreen) MinimumRefreshInterval() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("minimumRefreshInterval"))
	return rv
}
// The largest refresh interval that the screen supports.
//
// # Discussion
// 
// This interval represents the maximum amount of time (in seconds) your app
// has to generate new frames. It corresponds to the lowest refresh rate of
// the display.
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/maximumRefreshInterval
func (s NSScreen) MaximumRefreshInterval() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("maximumRefreshInterval"))
	return rv
}
// The number of seconds between the screen’s supported update rates, for
// screens that support fixed update rates.
//
// # Discussion
// 
// All screen refresh rates fall between the values in the
// [MinimumRefreshInterval] and [MaximumRefreshInterval] properties. For
// screens that support fixed update rates, this property contains the amount
// of time between two successive rates. For example, if a screen supports
// update rates between 30Hz and 120Hz with an update granularity of 5ms, the
// screen supports additional refresh rates of approximately 35Hz, 43Hz, 55Hz,
// and 75Hz.
// 
// If the value of this property is `0`, the screen supports any update rate
// between the minimum and maximum refresh intervals.
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/displayUpdateGranularity
func (s NSScreen) DisplayUpdateGranularity() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("displayUpdateGranularity"))
	return rv
}
// The time of the last framebuffer update, expressed as the number of seconds
// since system startup.
//
// # Discussion
// 
// Use this property to determine how much time elapsed since the last frame
// update.
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/lastDisplayUpdateTimestamp
func (s NSScreen) LastDisplayUpdateTimestamp() float64 {
	rv := objc.Send[float64](s.ID, objc.Sel("lastDisplayUpdateTimestamp"))
	return rv
}
// The CGDirectDisplayID for this screen. This will return nil if there
// isn’t one and will never return kCGNullDirectDisplay.
//
// See: https://developer.apple.com/documentation/appkit/nsscreen/cgdirectdisplayid-8ph5i
func (s NSScreen) CgDirectDisplayID() uint32 {
	rv := objc.Send[uint32](s.ID, objc.Sel("cgDirectDisplayID"))
	return rv
}
func (s NSScreen) SetCgDirectDisplayID(value uint32) {
	objc.Send[struct{}](s.ID, objc.Sel("setCgDirectDisplayID:"), value)
}

// Returns the screen object containing the window with the keyboard focus.
//
// # Return Value
// 
// The main screen object.
// 
// # Discussion
// 
// The main screen is not necessarily the same screen that contains the menu
// bar or has its origin at `(0, 0)`. The main screen refers to the screen
// containing the window that is currently receiving keyboard events. It is
// the main screen because it is the one with which the user is most likely
// interacting.
// 
// The screen containing the menu bar is always the first object (index `0`)
// in the array returned by the [Screens] method.
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/main
func (_NSScreenClass NSScreenClass) MainScreen() NSScreen {
	rv := objc.Send[objc.ID](objc.ID(_NSScreenClass.class), objc.Sel("mainScreen"))
	return NSScreenFromID(objc.ID(rv))
}
// Returns a screen object representing the screen that can best represent
// color.
//
// # Return Value
// 
// The screen with the highest bit depth.
// 
// # Discussion
// 
// This method always returns an object, even if there is only one screen and
// it is not a color screen.
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/deepest
func (_NSScreenClass NSScreenClass) DeepestScreen() NSScreen {
	rv := objc.Send[objc.ID](objc.ID(_NSScreenClass.class), objc.Sel("deepestScreen"))
	return NSScreenFromID(objc.ID(rv))
}
// Returns an array of screen objects representing all of the screens
// available on the system.
//
// # Return Value
// 
// An array of the [NSScreen] objects available on the current system.
// 
// # Discussion
// 
// The screen at index `0` in the returned array corresponds to the primary
// screen of the user’s system. This is the screen that contains the menu
// bar and whose origin is at the point `(0, 0)`. In the case of mirroring,
// the first screen is the largest drawable display; if all screens are the
// same size, it is the screen with the highest pixel depth. This primary
// screen may not be the same as the one returned by the [MainScreen] method,
// which returns the screen with the active window.
// 
// The array should not be cached. Screens can be added, removed, or
// dynamically reconfigured at any time. When the display configuration is
// changed, the default notification center sends a
// [didChangeScreenParametersNotification] notification.
//
// [didChangeScreenParametersNotification]: https://developer.apple.com/documentation/AppKit/NSApplication/didChangeScreenParametersNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/screens
func (_NSScreenClass NSScreenClass) Screens() []NSScreen {
	rv := objc.Send[[]objc.ID](objc.ID(_NSScreenClass.class), objc.Sel("screens"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSScreen {
		return NSScreenFromID(id)
	})
}
// Returns a Boolean value indicating whether each screen can have its own set
// of spaces.
//
// # Discussion
// 
// This method reflects whether the “Displays have separate Spaces” option
// is enabled in Mission Control system preference. You might use the return
// value to determine how to present your app when in fullscreen mode.
//
// See: https://developer.apple.com/documentation/AppKit/NSScreen/screensHaveSeparateSpaces
func (_NSScreenClass NSScreenClass) ScreensHaveSeparateSpaces() bool {
	rv := objc.Send[bool](objc.ID(_NSScreenClass.class), objc.Sel("screensHaveSeparateSpaces"))
	return rv
}
// Returns the application instance, creating it if it doesn’t exist yet.
//
// See: https://developer.apple.com/documentation/appkit/nsapplication/shared
func (_NSScreenClass NSScreenClass) Shared() NSApplication {
	rv := objc.Send[objc.ID](objc.ID(_NSScreenClass.class), objc.Sel("sharedApplication"))
	return NSApplicationFromID(objc.ID(rv))
}
func (_NSScreenClass NSScreenClass) SetShared(value NSApplication) {
	objc.Send[struct{}](objc.ID(_NSScreenClass.class), objc.Sel("setSharedApplication:"), value)
}

