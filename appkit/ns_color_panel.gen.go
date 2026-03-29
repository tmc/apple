// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSColorPanel] class.
var (
	_NSColorPanelClass     NSColorPanelClass
	_NSColorPanelClassOnce sync.Once
)

func getNSColorPanelClass() NSColorPanelClass {
	_NSColorPanelClassOnce.Do(func() {
		_NSColorPanelClass = NSColorPanelClass{class: objc.GetClass("NSColorPanel")}
	})
	return _NSColorPanelClass
}

// GetNSColorPanelClass returns the class object for NSColorPanel.
func GetNSColorPanelClass() NSColorPanelClass {
	return getNSColorPanelClass()
}

type NSColorPanelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSColorPanelClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSColorPanelClass) Alloc() NSColorPanel {
	rv := objc.Send[NSColorPanel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A standard user interface for selecting color in an app.
//
// # Overview
// 
// [NSColorPanel] provides a number of standard color selection modes and,
// with the [NSColorPickingDefault] and [NSColorPickingCustom] protocols,
// allows an app to add its own color selection modes. It also allows the user
// to save swatches containing frequently used colors.
//
// # Setting color picker modes
//
//   - [NSColorPanel.Mode]: The mode of the receiver the mode is one of the modes allowed by the color mask.
//   - [NSColorPanel.SetMode]
//
// # Configuring the color panel
//
//   - [NSColorPanel.AccessoryView]: The accessory view.
//   - [NSColorPanel.SetAccessoryView]
//   - [NSColorPanel.Continuous]: A Boolean value indicating whether the receiver continuously sends the action message to the target.
//   - [NSColorPanel.SetContinuous]
//   - [NSColorPanel.SetAction]: Sets the color panel’s action message.
//   - [NSColorPanel.SetTarget]: Sets the target of the receiver.
//   - [NSColorPanel.ShowsAlpha]: A Boolean value that indicates whether the receiver shows alpha values and an opacity slider.
//   - [NSColorPanel.SetShowsAlpha]
//
// # Managing color lists
//
//   - [NSColorPanel.AttachColorList]: Adds the list of [NSColor] objects specified to all the color pickers in the receiver that display color lists by invoking [attachColorList(_:)](<doc://com.apple.appkit/documentation/AppKit/NSColorPanel/attachColorList(_:)>) on all color pickers in the application.
//   - [NSColorPanel.DetachColorList]: Removes the list of colors from all the color pickers in the receiver that display color lists by invoking [detachColorList(_:)](<doc://com.apple.appkit/documentation/AppKit/NSColorPanel/detachColorList(_:)>) on all color pickers in the application.
//
// # Setting color
//
//   - [NSColorPanel.Color]: The color of the receiver.
//   - [NSColorPanel.SetColor]
//
// # Supporting high dynamic range (HDR) colors
//
//   - [NSColorPanel.MaximumLinearExposure]: The maximum linear exposure that can be set on a color picked in the color panel. Defaults to 1 and ignores any value less than 1. If set to a value >= 2, the color picked by the panel may have a linear exposure applied to it.
//   - [NSColorPanel.SetMaximumLinearExposure]
//
// # Getting transparency information
//
//   - [NSColorPanel.Alpha]: The receiver’s current alpha value based on its opacity slider.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPanel
type NSColorPanel struct {
	NSPanel
}

// NSColorPanelFromID constructs a [NSColorPanel] from an objc.ID.
//
// A standard user interface for selecting color in an app.
func NSColorPanelFromID(id objc.ID) NSColorPanel {
	return NSColorPanel{NSPanel: NSPanelFromID(id)}
}
// NOTE: NSColorPanel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSColorPanel] class.
//
// # Setting color picker modes
//
//   - [INSColorPanel.Mode]: The mode of the receiver the mode is one of the modes allowed by the color mask.
//   - [INSColorPanel.SetMode]
//
// # Configuring the color panel
//
//   - [INSColorPanel.AccessoryView]: The accessory view.
//   - [INSColorPanel.SetAccessoryView]
//   - [INSColorPanel.Continuous]: A Boolean value indicating whether the receiver continuously sends the action message to the target.
//   - [INSColorPanel.SetContinuous]
//   - [INSColorPanel.SetAction]: Sets the color panel’s action message.
//   - [INSColorPanel.SetTarget]: Sets the target of the receiver.
//   - [INSColorPanel.ShowsAlpha]: A Boolean value that indicates whether the receiver shows alpha values and an opacity slider.
//   - [INSColorPanel.SetShowsAlpha]
//
// # Managing color lists
//
//   - [INSColorPanel.AttachColorList]: Adds the list of [NSColor] objects specified to all the color pickers in the receiver that display color lists by invoking [attachColorList(_:)](<doc://com.apple.appkit/documentation/AppKit/NSColorPanel/attachColorList(_:)>) on all color pickers in the application.
//   - [INSColorPanel.DetachColorList]: Removes the list of colors from all the color pickers in the receiver that display color lists by invoking [detachColorList(_:)](<doc://com.apple.appkit/documentation/AppKit/NSColorPanel/detachColorList(_:)>) on all color pickers in the application.
//
// # Setting color
//
//   - [INSColorPanel.Color]: The color of the receiver.
//   - [INSColorPanel.SetColor]
//
// # Supporting high dynamic range (HDR) colors
//
//   - [INSColorPanel.MaximumLinearExposure]: The maximum linear exposure that can be set on a color picked in the color panel. Defaults to 1 and ignores any value less than 1. If set to a value >= 2, the color picked by the panel may have a linear exposure applied to it.
//   - [INSColorPanel.SetMaximumLinearExposure]
//
// # Getting transparency information
//
//   - [INSColorPanel.Alpha]: The receiver’s current alpha value based on its opacity slider.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPanel
type INSColorPanel interface {
	INSPanel

	// Topic: Setting color picker modes

	// The mode of the receiver the mode is one of the modes allowed by the color mask.
	Mode() NSColorPanelMode
	SetMode(value NSColorPanelMode)

	// Topic: Configuring the color panel

	// The accessory view.
	AccessoryView() INSView
	SetAccessoryView(value INSView)
	// A Boolean value indicating whether the receiver continuously sends the action message to the target.
	Continuous() bool
	SetContinuous(value bool)
	// Sets the color panel’s action message.
	SetAction(selector objc.SEL)
	// Sets the target of the receiver.
	SetTarget(target objectivec.IObject)
	// A Boolean value that indicates whether the receiver shows alpha values and an opacity slider.
	ShowsAlpha() bool
	SetShowsAlpha(value bool)

	// Topic: Managing color lists

	// Adds the list of [NSColor] objects specified to all the color pickers in the receiver that display color lists by invoking [attachColorList(_:)](<doc://com.apple.appkit/documentation/AppKit/NSColorPanel/attachColorList(_:)>) on all color pickers in the application.
	AttachColorList(colorList INSColorList)
	// Removes the list of colors from all the color pickers in the receiver that display color lists by invoking [detachColorList(_:)](<doc://com.apple.appkit/documentation/AppKit/NSColorPanel/detachColorList(_:)>) on all color pickers in the application.
	DetachColorList(colorList INSColorList)

	// Topic: Setting color

	// The color of the receiver.
	Color() INSColor
	SetColor(value INSColor)

	// Topic: Supporting high dynamic range (HDR) colors

	// The maximum linear exposure that can be set on a color picked in the color panel. Defaults to 1 and ignores any value less than 1. If set to a value >= 2, the color picked by the panel may have a linear exposure applied to it.
	MaximumLinearExposure() float64
	SetMaximumLinearExposure(value float64)

	// Topic: Getting transparency information

	// The receiver’s current alpha value based on its opacity slider.
	Alpha() float64
}

// Init initializes the instance.
func (c NSColorPanel) Init() NSColorPanel {
	rv := objc.Send[NSColorPanel](c.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c NSColorPanel) Autorelease() NSColorPanel {
	rv := objc.Send[NSColorPanel](c.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSColorPanel creates a new NSColorPanel instance.
func NewNSColorPanel() NSColorPanel {
	class := getNSColorPanelClass()
	rv := objc.Send[NSColorPanel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a titled window that contains the specified content view
// controller.
//
// contentViewController: The view controller that provides the main content view for the window. The
// window’s [ContentView] property is set to
// `contentViewController``XCUIElementTypeView`.
//
// # Return Value
// 
// A window with the content view controller set to the passed-in view
// controller object.
//
// # Discussion
// 
// This method creates a basic window object that is titled, closable,
// resizable, and miniaturizable. By default, the window’s title is
// automatically bound to the title of `contentViewController`. You can
// control the size of the window by using Auto Layout and applying size
// constraints to the view or its subviews. The initial size of the window is
// set to the initial size of [ContentView] (that is, the size of
// `contentViewController``XCUIElementTypeView`). The newly created window has
// [ReleasedWhenClosed] set to [false], and it must be explicitly retained to
// keep the window instance alive.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/init(contentViewController:)
func NewColorPanelWindowWithContentViewController(contentViewController INSViewController) NSColorPanel {
	rv := objc.Send[objc.ID](objc.ID(getNSColorPanelClass().class), objc.Sel("windowWithContentViewController:"), contentViewController)
	return NSColorPanelFromID(rv)
}

// Creates a new responder object with data in an unarchiver.
//
// coder: An unarchiver object.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/init(coder:)
func NewColorPanelWithCoder(coder foundation.INSCoder) NSColorPanel {
	instance := getNSColorPanelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSColorPanelFromID(rv)
}

// Initializes the window with the specified values.
//
// contentRect: Origin and size of the window’s content area in screen coordinates. Note
// that the window server limits window position coordinates to ±16,000 and
// sizes to 10,000.
//
// style: The window’s style. It can be [NSBorderlessWindowMask], or it can contain
// any of the options described in [NSWindow.StyleMask], combined using the C
// bitwise OR operator. Borderless windows display none of the usual
// peripheral elements and are generally useful only for display or caching
// purposes; you should normally not need to create them. Also, note that a
// window’s style mask should include [NSTitledWindowMask] if it includes
// any of the others.
// //
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
//
// backingStoreType: Specifies how the drawing done in the window is buffered by the window
// device, and possible values are described in [NSWindow.BackingStoreType].
// //
// [NSWindow.BackingStoreType]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType
//
// flag: Specifies whether the window server creates a window device for the window
// immediately. When [true], the window server defers creating the window
// device until the window is moved onscreen. All display messages sent to the
// window or its views are postponed until the window is created, just before
// it’s moved onscreen.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// The initialized window.
//
// # Discussion
// 
// This method is the designated initializer for the [NSWindow] class.
// 
// Deferring the creation of the window improves launch time and minimizes the
// virtual memory load on the window server.
// 
// The new window creates a view to be its default content view. You can
// replace it with your own object by setting the [ContentView] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/init(contentRect:styleMask:backing:defer:)
func NewColorPanelWithContentRectStyleMaskBackingDefer(contentRect corefoundation.CGRect, style NSWindowStyleMask, backingStoreType NSBackingStoreType, flag bool) NSColorPanel {
	instance := getNSColorPanelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return NSColorPanelFromID(rv)
}

// Initializes an allocated window with the specified values.
//
// contentRect: Origin and size of the window’s content area in screen coordinates. The
// origin is relative to the origin of the provided screen. Note that the
// window server limits window position coordinates to ±16,000 and sizes to
// 10,000.
//
// style: The window’s style. It can be [NSBorderlessWindowMask], or it can contain
// any of the options described in [NSWindow.StyleMask], combined using the C
// bitwise OR operator. Borderless windows display none of the usual
// peripheral elements and are generally useful only for display or caching
// purposes; you should not usually need to create them. Also, note that a
// window’s style mask should include [NSTitledWindowMask] if it includes
// any of the others.
// //
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
//
// backingStoreType: Specifies how the drawing done in the window is buffered by the window
// device; possible values are described in [NSWindow.BackingStoreType].
// //
// [NSWindow.BackingStoreType]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType
//
// flag: Specifies whether the window server creates a window device for the window
// immediately. When [true], the window server defers creating the window
// device until the window is moved onscreen. All display messages sent to the
// window or its views are postponed until the window is created, just before
// it’s moved onscreen.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// screen: Specifies the screen on which the window is positioned. The content
// rectangle is positioned relative to the bottom-left corner of `screen`.
// When `nil`, the content rectangle is positioned relative to (0, 0), which
// is the origin of the primary screen.
//
// # Return Value
// 
// The initialized window.
//
// # Discussion
// 
// The primary screen is the one that contains the current key window or, if
// there is no key window, the one that contains the main menu. If there’s
// neither a key window nor a main menu (if there’s no active application),
// the primary screen is the one where the origin of the screen coordinate
// system is located.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/init(contentRect:styleMask:backing:defer:screen:)
func NewColorPanelWithContentRectStyleMaskBackingDeferScreen(contentRect corefoundation.CGRect, style NSWindowStyleMask, backingStoreType NSBackingStoreType, flag bool, screen INSScreen) NSColorPanel {
	instance := getNSColorPanelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, screen)
	return NSColorPanelFromID(rv)
}

// Sets the color panel’s action message.
//
// selector: The action message.
//
// # Discussion
// 
// When you select a color in the color panel [NSColorPanel] sends its action
// to its target, provided that neither the action nor the target is `nil`.
// The action is [NULL] by default.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPanel/setAction(_:)
func (c NSColorPanel) SetAction(selector objc.SEL) {
	objc.Send[objc.ID](c.ID, objc.Sel("setAction:"), selector)
}
// Sets the target of the receiver.
//
// target: The target of the receiver. When you select a color in the color panel
// [NSColorPanel] sends its action to its target, provided that neither the
// action nor the target is `nil`. The target is `nil` by default.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPanel/setTarget(_:)
func (c NSColorPanel) SetTarget(target objectivec.IObject) {
	objc.Send[objc.ID](c.ID, objc.Sel("setTarget:"), target)
}
// Adds the list of [NSColor] objects specified to all the color pickers in
// the receiver that display color lists by invoking [AttachColorList] on all
// color pickers in the application.
//
// colorList: The list of colors to add to the color pickers in the receiver.
//
// # Discussion
// 
// An application should use this method to add an [NSColorList] saved with a
// document in its file package or in a directory other than [NSColorList]’s
// standard search directories.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPanel/attachColorList(_:)
func (c NSColorPanel) AttachColorList(colorList INSColorList) {
	objc.Send[objc.ID](c.ID, objc.Sel("attachColorList:"), colorList)
}
// Removes the list of colors from all the color pickers in the receiver that
// display color lists by invoking [DetachColorList] on all color pickers in
// the application.
//
// colorList: The list of [NSColor] objects to remove from the color pickers in the color
// panel.
//
// # Discussion
// 
// Your application should use this method to remove an [NSColorList] saved
// with a document in its file package or in a directory other than
// [NSColorList]’s standard search directories.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPanel/detachColorList(_:)
func (c NSColorPanel) DetachColorList(colorList INSColorList) {
	objc.Send[objc.ID](c.ID, objc.Sel("detachColorList:"), colorList)
}

// Specifies the color panel’s initial picker.
//
// mode: A constant specifying which color picker mode is initially visible. This is
// one of the symbolic constants described in `Color Panel Modes`.
//
// # Discussion
// 
// This method may be called at any time, whether or not an application’s
// [NSColorPanel] has been instantiated.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPanel/setPickerMode(_:)
func (_NSColorPanelClass NSColorPanelClass) SetPickerMode(mode NSColorPanelMode) {
	objc.Send[objc.ID](objc.ID(_NSColorPanelClass.class), objc.Sel("setPickerMode:"), mode)
}
// Determines which color selection modes are available in an application’s
// [NSColorPanel].
//
// mask: One or more logically ORed color mode masks described in `Color Picker Mode
// Masks`.
//
// # Discussion
// 
// This method has an effect only before an [NSColorPanel] object is
// instantiated.
// 
// If you create a class that implements the color-picking protocols
// ([NSColorPickingDefault] and [NSColorPickingCustom]), you may want to give
// it a unique mask—one different from those defined for the standard color
// pickers. To display your color picker, your application will need to
// logically OR that unique mask with the standard color mask constants when
// invoking this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPanel/setPickerMask(_:)
func (_NSColorPanelClass NSColorPanelClass) SetPickerMask(mask NSColorPanelOptions) {
	objc.Send[objc.ID](objc.ID(_NSColorPanelClass.class), objc.Sel("setPickerMask:"), mask)
}
// Drags a color into a destination view from the specified source view.
//
// color: The color to drag.
//
// event: The drag event.
//
// sourceView: The view from which the color was dragged.
//
// # Return Value
// 
// [true]
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is usually invoked by the `` method of `sourceView`. The
// dragging mechanism handles all subsequent events.
// 
// Because it is a class method, [DragColorWithEventFromView] can be invoked
// whether or not the instance of [NSColorPanel] exists.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPanel/dragColor(_:with:from:)
func (_NSColorPanelClass NSColorPanelClass) DragColorWithEventFromView(color INSColor, event INSEvent, sourceView INSView) bool {
	rv := objc.Send[bool](objc.ID(_NSColorPanelClass.class), objc.Sel("dragColor:withEvent:fromView:"), color, event, sourceView)
	return rv
}

// The mode of the receiver the mode is one of the modes allowed by the color
// mask.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPanel/mode-swift.property
func (c NSColorPanel) Mode() NSColorPanelMode {
	rv := objc.Send[NSColorPanelMode](c.ID, objc.Sel("mode"))
	return NSColorPanelMode(rv)
}
func (c NSColorPanel) SetMode(value NSColorPanelMode) {
	objc.Send[struct{}](c.ID, objc.Sel("setMode:"), value)
}
// The accessory view.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPanel/accessoryView
func (c NSColorPanel) AccessoryView() INSView {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("accessoryView"))
	return NSViewFromID(objc.ID(rv))
}
func (c NSColorPanel) SetAccessoryView(value INSView) {
	objc.Send[struct{}](c.ID, objc.Sel("setAccessoryView:"), value)
}
// A Boolean value indicating whether the receiver continuously sends the
// action message to the target.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPanel/isContinuous
func (c NSColorPanel) Continuous() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("isContinuous"))
	return rv
}
func (c NSColorPanel) SetContinuous(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setContinuous:"), value)
}
// A Boolean value that indicates whether the receiver shows alpha values and
// an opacity slider.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPanel/showsAlpha
func (c NSColorPanel) ShowsAlpha() bool {
	rv := objc.Send[bool](c.ID, objc.Sel("showsAlpha"))
	return rv
}
func (c NSColorPanel) SetShowsAlpha(value bool) {
	objc.Send[struct{}](c.ID, objc.Sel("setShowsAlpha:"), value)
}
// The color of the receiver.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPanel/color
func (c NSColorPanel) Color() INSColor {
	rv := objc.Send[objc.ID](c.ID, objc.Sel("color"))
	return NSColorFromID(objc.ID(rv))
}
func (c NSColorPanel) SetColor(value INSColor) {
	objc.Send[struct{}](c.ID, objc.Sel("setColor:"), value)
}
// The maximum linear exposure that can be set on a color picked in the color
// panel. Defaults to 1 and ignores any value less than 1. If set to a value
// >= 2, the color picked by the panel may have a linear exposure applied to
// it.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPanel/maximumLinearExposure
func (c NSColorPanel) MaximumLinearExposure() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("maximumLinearExposure"))
	return rv
}
func (c NSColorPanel) SetMaximumLinearExposure(value float64) {
	objc.Send[struct{}](c.ID, objc.Sel("setMaximumLinearExposure:"), value)
}
// The receiver’s current alpha value based on its opacity slider.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPanel/alpha
func (c NSColorPanel) Alpha() float64 {
	rv := objc.Send[float64](c.ID, objc.Sel("alpha"))
	return rv
}

// Returns the shared [NSColorPanel] instance, creating it if necessary.
//
// # Return Value
// 
// The shared [NSColorPanel] instance.
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPanel/shared
func (_NSColorPanelClass NSColorPanelClass) SharedColorPanel() NSColorPanel {
	rv := objc.Send[objc.ID](objc.ID(_NSColorPanelClass.class), objc.Sel("sharedColorPanel"))
	return NSColorPanelFromID(objc.ID(rv))
}
// Returns a Boolean value indicating whether the [NSColorPanel] has been
// created already.
//
// # Return Value
// 
// [true] if the [NSColorPanel] has been created already; otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSColorPanel/sharedColorPanelExists
func (_NSColorPanelClass NSColorPanelClass) SharedColorPanelExists() bool {
	rv := objc.Send[bool](objc.ID(_NSColorPanelClass.class), objc.Sel("sharedColorPanelExists"))
	return rv
}

