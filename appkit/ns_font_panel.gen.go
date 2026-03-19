// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSFontPanel] class.
var (
	_NSFontPanelClass     NSFontPanelClass
	_NSFontPanelClassOnce sync.Once
)

func getNSFontPanelClass() NSFontPanelClass {
	_NSFontPanelClassOnce.Do(func() {
		_NSFontPanelClass = NSFontPanelClass{class: objc.GetClass("NSFontPanel")}
	})
	return _NSFontPanelClass
}

// GetNSFontPanelClass returns the class object for NSFontPanel.
func GetNSFontPanelClass() NSFontPanelClass {
	return getNSFontPanelClass()
}

type NSFontPanelClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFontPanelClass) Alloc() NSFontPanel {
	rv := objc.Send[NSFontPanel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The Font panel—a user interface object that displays a list of available
// fonts, letting the user preview them and change the font used to display
// text.
//
// # Overview
// 
// Actual changes to the font panel are made through conversion messages sent
// to the shared [NSFontManager] instance. There’s only one Font panel for
// each app.
//
// # Enabling Font Changes
//
//   - [NSFontPanel.Enabled]: A Boolean that shows whether the receiver’s Set button is enabled.
//   - [NSFontPanel.SetEnabled]
//   - [NSFontPanel.ReloadDefaultFontFamilies]: Triggers a reload to the default state, so that the delegate is called.
//
// # Updating the Font Panel
//
//   - [NSFontPanel.SetPanelFontIsMultiple]: Sets the selected font in the receiver to the specified font.
//
// # Converting Fonts
//
//   - [NSFontPanel.PanelConvertFont]: Converts the specified font using the settings in the receiver, with the aid of the shared [NSFontManager] if necessary.
//
// # Setting an Accessory View
//
//   - [NSFontPanel.AccessoryView]: The specified view as the receiver’s accessory view, allowing you to add custom controls to your application’s Font panel without having to create a subclass.
//   - [NSFontPanel.SetAccessoryView]
//
// See: https://developer.apple.com/documentation/AppKit/NSFontPanel
type NSFontPanel struct {
	NSPanel
}

// NSFontPanelFromID constructs a [NSFontPanel] from an objc.ID.
//
// The Font panel—a user interface object that displays a list of available
// fonts, letting the user preview them and change the font used to display
// text.
func NSFontPanelFromID(id objc.ID) NSFontPanel {
	return NSFontPanel{NSPanel: NSPanelFromID(id)}
}
// NOTE: NSFontPanel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFontPanel] class.
//
// # Enabling Font Changes
//
//   - [INSFontPanel.Enabled]: A Boolean that shows whether the receiver’s Set button is enabled.
//   - [INSFontPanel.SetEnabled]
//   - [INSFontPanel.ReloadDefaultFontFamilies]: Triggers a reload to the default state, so that the delegate is called.
//
// # Updating the Font Panel
//
//   - [INSFontPanel.SetPanelFontIsMultiple]: Sets the selected font in the receiver to the specified font.
//
// # Converting Fonts
//
//   - [INSFontPanel.PanelConvertFont]: Converts the specified font using the settings in the receiver, with the aid of the shared [NSFontManager] if necessary.
//
// # Setting an Accessory View
//
//   - [INSFontPanel.AccessoryView]: The specified view as the receiver’s accessory view, allowing you to add custom controls to your application’s Font panel without having to create a subclass.
//   - [INSFontPanel.SetAccessoryView]
//
// See: https://developer.apple.com/documentation/AppKit/NSFontPanel
type INSFontPanel interface {
	INSPanel

	// Topic: Enabling Font Changes

	// A Boolean that shows whether the receiver’s Set button is enabled.
	Enabled() bool
	SetEnabled(value bool)
	// Triggers a reload to the default state, so that the delegate is called.
	ReloadDefaultFontFamilies()

	// Topic: Updating the Font Panel

	// Sets the selected font in the receiver to the specified font.
	SetPanelFontIsMultiple(fontObj NSFont, flag bool)

	// Topic: Converting Fonts

	// Converts the specified font using the settings in the receiver, with the aid of the shared [NSFontManager] if necessary.
	PanelConvertFont(fontObj NSFont) NSFont

	// Topic: Setting an Accessory View

	// The specified view as the receiver’s accessory view, allowing you to add custom controls to your application’s Font panel without having to create a subclass.
	AccessoryView() INSView
	SetAccessoryView(value INSView)
}

// Init initializes the instance.
func (f NSFontPanel) Init() NSFontPanel {
	rv := objc.Send[NSFontPanel](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFontPanel) Autorelease() NSFontPanel {
	rv := objc.Send[NSFontPanel](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFontPanel creates a new NSFontPanel instance.
func NewNSFontPanel() NSFontPanel {
	class := getNSFontPanelClass()
	rv := objc.Send[NSFontPanel](objc.ID(class.class), objc.Sel("new"))
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
func NewFontPanelWindowWithContentViewController(contentViewController INSViewController) NSFontPanel {
	rv := objc.Send[objc.ID](objc.ID(getNSFontPanelClass().class), objc.Sel("windowWithContentViewController:"), contentViewController)
	return NSFontPanelFromID(rv)
}

// Creates a new responder object with data in an unarchiver.
//
// coder: An unarchiver object.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/init(coder:)
func NewFontPanelWithCoder(coder foundation.INSCoder) NSFontPanel {
	instance := getNSFontPanelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSFontPanelFromID(rv)
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
func NewFontPanelWithContentRectStyleMaskBackingDefer(contentRect corefoundation.CGRect, style NSWindowStyleMask, backingStoreType NSBackingStoreType, flag bool) NSFontPanel {
	instance := getNSFontPanelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return NSFontPanelFromID(rv)
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
func NewFontPanelWithContentRectStyleMaskBackingDeferScreen(contentRect corefoundation.CGRect, style NSWindowStyleMask, backingStoreType NSBackingStoreType, flag bool, screen INSScreen) NSFontPanel {
	instance := getNSFontPanelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, screen)
	return NSFontPanelFromID(rv)
}

// Triggers a reload to the default state, so that the delegate is called.
//
// # Discussion
// 
// This reloading provides the delegate opportunity to scrutinize the default
// list of fonts to be displayed in the panel.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontPanel/reloadDefaultFontFamilies()
func (f NSFontPanel) ReloadDefaultFontFamilies() {
	objc.Send[objc.ID](f.ID, objc.Sel("reloadDefaultFontFamilies"))
}
// Sets the selected font in the receiver to the specified font.
//
// fontObj: The font to be selected.
//
// flag: If [false], selects the specified font; otherwise selects no font and
// displays a message in the preview area indicating that multiple fonts are
// selected.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
//
// # Discussion
// 
// You normally don’t use this method directly; instead, you send
// [SetSelectedFontIsMultiple] to the shared [NSFontManager], which in turn
// invokes this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontPanel/setPanelFont(_:isMultiple:)
func (f NSFontPanel) SetPanelFontIsMultiple(fontObj NSFont, flag bool) {
	objc.Send[objc.ID](f.ID, objc.Sel("setPanelFont:isMultiple:"), fontObj, flag)
}
// Converts the specified font using the settings in the receiver, with the
// aid of the shared [NSFontManager] if necessary.
//
// fontObj: The font to be converted.
//
// # Return Value
// 
// The converted font, or `aFont` itself if it can’t be converted.
//
// # Discussion
// 
// For example, if `aFont` is Helvetica Oblique 12.0 point and the user has
// selected the Times font family (and nothing else) in the Font panel, the
// font returned is Times Italic 12.0 point.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontPanel/convert(_:)
func (f NSFontPanel) PanelConvertFont(fontObj NSFont) NSFont {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("panelConvertFont:"), fontObj)
	return NSFontFromID(rv)
}

// A Boolean that shows whether the receiver’s Set button is enabled.
//
// # Discussion
// 
// The receiver continues to reflect the font of the selection for cooperating
// text objects regardless of this setting.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontPanel/isEnabled
func (f NSFontPanel) Enabled() bool {
	rv := objc.Send[bool](f.ID, objc.Sel("isEnabled"))
	return rv
}
func (f NSFontPanel) SetEnabled(value bool) {
	objc.Send[struct{}](f.ID, objc.Sel("setEnabled:"), value)
}
// The specified view as the receiver’s accessory view, allowing you to add
// custom controls to your application’s Font panel without having to create
// a subclass.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontPanel/accessoryView
func (f NSFontPanel) AccessoryView() INSView {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("accessoryView"))
	return NSViewFromID(objc.ID(rv))
}
func (f NSFontPanel) SetAccessoryView(value INSView) {
	objc.Send[struct{}](f.ID, objc.Sel("setAccessoryView:"), value)
}

// Returns the single [NSFontPanel] instance for the application, creating it
// if necessary.
//
// # Return Value
// 
// The [NSFontPanel] instance for the application.
//
// See: https://developer.apple.com/documentation/AppKit/NSFontPanel/shared
func (_NSFontPanelClass NSFontPanelClass) SharedFontPanel() NSFontPanel {
	rv := objc.Send[objc.ID](objc.ID(_NSFontPanelClass.class), objc.Sel("sharedFontPanel"))
	return NSFontPanelFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the shared Font panel has been
// created.
//
// # Discussion
// 
// The value is [true] if the shared Font panel has been created, and [false]
// if it hasn’t.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSFontPanel/sharedFontPanelExists
func (_NSFontPanelClass NSFontPanelClass) SharedFontPanelExists() bool {
	rv := objc.Send[bool](objc.ID(_NSFontPanelClass.class), objc.Sel("sharedFontPanelExists"))
	return rv
}

