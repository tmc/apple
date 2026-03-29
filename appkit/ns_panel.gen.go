// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSPanel] class.
var (
	_NSPanelClass     NSPanelClass
	_NSPanelClassOnce sync.Once
)

func getNSPanelClass() NSPanelClass {
	_NSPanelClassOnce.Do(func() {
		_NSPanelClass = NSPanelClass{class: objc.GetClass("NSPanel")}
	})
	return _NSPanelClass
}

// GetNSPanelClass returns the class object for NSPanel.
func GetNSPanelClass() NSPanelClass {
	return getNSPanelClass()
}

type NSPanelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPanelClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPanelClass) Alloc() NSPanel {
	rv := objc.Send[NSPanel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A special kind of window that typically performs a function that is
// auxiliary to the main window.
//
// # Overview
// 
// For details about how panels work (especially to find out how their
// behavior differs from window behavior), see [How Panels Work].
//
// [How Panels Work]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/WinPanel/Concepts/UsingPanels.html#//apple_ref/doc/uid/20000224
//
// # Configuring Panels
//
//   - [NSPanel.BecomesKeyOnlyIfNeeded]: A Boolean value that indicates whether the receiver becomes the key window only when needed.
//   - [NSPanel.SetBecomesKeyOnlyIfNeeded]
//
// See: https://developer.apple.com/documentation/AppKit/NSPanel
type NSPanel struct {
	NSWindow
}

// NSPanelFromID constructs a [NSPanel] from an objc.ID.
//
// A special kind of window that typically performs a function that is
// auxiliary to the main window.
func NSPanelFromID(id objc.ID) NSPanel {
	return NSPanel{NSWindow: NSWindowFromID(id)}
}
// NOTE: NSPanel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPanel] class.
//
// # Configuring Panels
//
//   - [INSPanel.BecomesKeyOnlyIfNeeded]: A Boolean value that indicates whether the receiver becomes the key window only when needed.
//   - [INSPanel.SetBecomesKeyOnlyIfNeeded]
//
// See: https://developer.apple.com/documentation/AppKit/NSPanel
type INSPanel interface {
	INSWindow

	// Topic: Configuring Panels

	// A Boolean value that indicates whether the receiver becomes the key window only when needed.
	BecomesKeyOnlyIfNeeded() bool
	SetBecomesKeyOnlyIfNeeded(value bool)
}

// Init initializes the instance.
func (p NSPanel) Init() NSPanel {
	rv := objc.Send[NSPanel](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPanel) Autorelease() NSPanel {
	rv := objc.Send[NSPanel](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPanel creates a new NSPanel instance.
func NewNSPanel() NSPanel {
	class := getNSPanelClass()
	rv := objc.Send[NSPanel](objc.ID(class.class), objc.Sel("new"))
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
func NewPanelWindowWithContentViewController(contentViewController INSViewController) NSPanel {
	rv := objc.Send[objc.ID](objc.ID(getNSPanelClass().class), objc.Sel("windowWithContentViewController:"), contentViewController)
	return NSPanelFromID(rv)
}

// Creates a new responder object with data in an unarchiver.
//
// coder: An unarchiver object.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/init(coder:)
func NewPanelWithCoder(coder foundation.INSCoder) NSPanel {
	instance := getNSPanelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSPanelFromID(rv)
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
func NewPanelWithContentRectStyleMaskBackingDefer(contentRect corefoundation.CGRect, style NSWindowStyleMask, backingStoreType NSBackingStoreType, flag bool) NSPanel {
	instance := getNSPanelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return NSPanelFromID(rv)
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
func NewPanelWithContentRectStyleMaskBackingDeferScreen(contentRect corefoundation.CGRect, style NSWindowStyleMask, backingStoreType NSBackingStoreType, flag bool, screen INSScreen) NSPanel {
	instance := getNSPanelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, screen)
	return NSPanelFromID(rv)
}

// A Boolean value that indicates whether the receiver becomes the key window
// only when needed.
//
// # Discussion
// 
// The value of this property is [true] when the panel becomes the key window
// only when keyboard input is required; the value is [false] when the panel
// becomes key when it’s clicked. The default value is [false].
// 
// This behavior is not set by default. You should consider setting it only if
// most user interface elements in the panel aren’t text fields, and if the
// choices that can be made by entering text can also be made in another way
// (such as by clicking an item in a list).
// 
// If the panel is a non-activating panel, then it becomes key only if the hit
// view returns [true] from [needsPanelToBecomeKey]. This way, a
// non-activating panel can control whether it takes keyboard focus.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [needsPanelToBecomeKey]: https://developer.apple.com/documentation/AppKit/NSView/needsPanelToBecomeKey
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSPanel/becomesKeyOnlyIfNeeded
func (p NSPanel) BecomesKeyOnlyIfNeeded() bool {
	rv := objc.Send[bool](p.ID, objc.Sel("becomesKeyOnlyIfNeeded"))
	return rv
}
func (p NSPanel) SetBecomesKeyOnlyIfNeeded(value bool) {
	objc.Send[struct{}](p.ID, objc.Sel("setBecomesKeyOnlyIfNeeded:"), value)
}

