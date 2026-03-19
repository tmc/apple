// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
)

// The class instance for the [NSOpenPanel] class.
var (
	_NSOpenPanelClass     NSOpenPanelClass
	_NSOpenPanelClassOnce sync.Once
)

func getNSOpenPanelClass() NSOpenPanelClass {
	_NSOpenPanelClassOnce.Do(func() {
		_NSOpenPanelClass = NSOpenPanelClass{class: objc.GetClass("NSOpenPanel")}
	})
	return _NSOpenPanelClass
}

// GetNSOpenPanelClass returns the class object for NSOpenPanel.
func GetNSOpenPanelClass() NSOpenPanelClass {
	return getNSOpenPanelClass()
}

type NSOpenPanelClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSOpenPanelClass) Alloc() NSOpenPanel {
	rv := objc.Send[NSOpenPanel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A panel that prompts the user to select a file to open.
//
// # Overview
// 
// Apps use the Open panel as a convenient way to query the user for the name
// of a file to open. In macOS 10.15 and later, the system always draws Open
// panels in a separate process, regardless of whether the app is sandboxed.
// When the user chooses a file to open, macOS adds that file to the app’s
// sandbox. Prior to macOS 10.15, the system drew the panels in a separate
// process only for sandboxed apps.
//
// # Configuring the Open Panel
//
//   - [NSOpenPanel.CanChooseFiles]: A Boolean that indicates whether the user can choose files in the panel.
//   - [NSOpenPanel.SetCanChooseFiles]
//   - [NSOpenPanel.CanChooseDirectories]: A Boolean that indicates whether the user can choose directories in the panel.
//   - [NSOpenPanel.SetCanChooseDirectories]
//   - [NSOpenPanel.ResolvesAliases]: A Boolean that indicates whether the panel resolves aliases.
//   - [NSOpenPanel.SetResolvesAliases]
//   - [NSOpenPanel.AllowsMultipleSelection]: A Boolean that indicates whether the user may select multiple files and directories.
//   - [NSOpenPanel.SetAllowsMultipleSelection]
//   - [NSOpenPanel.AccessoryViewDisclosed]: A Boolean value that indicates whether the panel’s accessory view is visible.
//   - [NSOpenPanel.SetAccessoryViewDisclosed]
//
// # Accessing User Selection
//
//   - [NSOpenPanel.URLs]: An array of URLs, each of which contains the fully specified location of a selected file or directory.
//
// # Supporting iCloud Documents
//
//   - [NSOpenPanel.CanDownloadUbiquitousContents]: A Boolean value that indicates how the panel responds to iCloud documents that aren’t fully downloaded locally.
//   - [NSOpenPanel.SetCanDownloadUbiquitousContents]
//   - [NSOpenPanel.CanResolveUbiquitousConflicts]: A Boolean value that indicates how the panel responds to iCloud documents that have conflicting versions.
//   - [NSOpenPanel.SetCanResolveUbiquitousConflicts]
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenPanel
type NSOpenPanel struct {
	NSSavePanel
}

// NSOpenPanelFromID constructs a [NSOpenPanel] from an objc.ID.
//
// A panel that prompts the user to select a file to open.
func NSOpenPanelFromID(id objc.ID) NSOpenPanel {
	return NSOpenPanel{NSSavePanel: NSSavePanelFromID(id)}
}
// NOTE: NSOpenPanel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSOpenPanel] class.
//
// # Configuring the Open Panel
//
//   - [INSOpenPanel.CanChooseFiles]: A Boolean that indicates whether the user can choose files in the panel.
//   - [INSOpenPanel.SetCanChooseFiles]
//   - [INSOpenPanel.CanChooseDirectories]: A Boolean that indicates whether the user can choose directories in the panel.
//   - [INSOpenPanel.SetCanChooseDirectories]
//   - [INSOpenPanel.ResolvesAliases]: A Boolean that indicates whether the panel resolves aliases.
//   - [INSOpenPanel.SetResolvesAliases]
//   - [INSOpenPanel.AllowsMultipleSelection]: A Boolean that indicates whether the user may select multiple files and directories.
//   - [INSOpenPanel.SetAllowsMultipleSelection]
//   - [INSOpenPanel.AccessoryViewDisclosed]: A Boolean value that indicates whether the panel’s accessory view is visible.
//   - [INSOpenPanel.SetAccessoryViewDisclosed]
//
// # Accessing User Selection
//
//   - [INSOpenPanel.URLs]: An array of URLs, each of which contains the fully specified location of a selected file or directory.
//
// # Supporting iCloud Documents
//
//   - [INSOpenPanel.CanDownloadUbiquitousContents]: A Boolean value that indicates how the panel responds to iCloud documents that aren’t fully downloaded locally.
//   - [INSOpenPanel.SetCanDownloadUbiquitousContents]
//   - [INSOpenPanel.CanResolveUbiquitousConflicts]: A Boolean value that indicates how the panel responds to iCloud documents that have conflicting versions.
//   - [INSOpenPanel.SetCanResolveUbiquitousConflicts]
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenPanel
type INSOpenPanel interface {
	INSSavePanel

	// Topic: Configuring the Open Panel

	// A Boolean that indicates whether the user can choose files in the panel.
	CanChooseFiles() bool
	SetCanChooseFiles(value bool)
	// A Boolean that indicates whether the user can choose directories in the panel.
	CanChooseDirectories() bool
	SetCanChooseDirectories(value bool)
	// A Boolean that indicates whether the panel resolves aliases.
	ResolvesAliases() bool
	SetResolvesAliases(value bool)
	// A Boolean that indicates whether the user may select multiple files and directories.
	AllowsMultipleSelection() bool
	SetAllowsMultipleSelection(value bool)
	// A Boolean value that indicates whether the panel’s accessory view is visible.
	AccessoryViewDisclosed() bool
	SetAccessoryViewDisclosed(value bool)

	// Topic: Accessing User Selection

	// An array of URLs, each of which contains the fully specified location of a selected file or directory.
	URLs() []foundation.NSURL

	// Topic: Supporting iCloud Documents

	// A Boolean value that indicates how the panel responds to iCloud documents that aren’t fully downloaded locally.
	CanDownloadUbiquitousContents() bool
	SetCanDownloadUbiquitousContents(value bool)
	// A Boolean value that indicates how the panel responds to iCloud documents that have conflicting versions.
	CanResolveUbiquitousConflicts() bool
	SetCanResolveUbiquitousConflicts(value bool)
}

// Init initializes the instance.
func (o NSOpenPanel) Init() NSOpenPanel {
	rv := objc.Send[NSOpenPanel](o.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o NSOpenPanel) Autorelease() NSOpenPanel {
	rv := objc.Send[NSOpenPanel](o.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSOpenPanel creates a new NSOpenPanel instance.
func NewNSOpenPanel() NSOpenPanel {
	class := getNSOpenPanelClass()
	rv := objc.Send[NSOpenPanel](objc.ID(class.class), objc.Sel("new"))
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
func NewOpenPanelWindowWithContentViewController(contentViewController INSViewController) NSOpenPanel {
	rv := objc.Send[objc.ID](objc.ID(getNSOpenPanelClass().class), objc.Sel("windowWithContentViewController:"), contentViewController)
	return NSOpenPanelFromID(rv)
}

// Creates a new responder object with data in an unarchiver.
//
// coder: An unarchiver object.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/init(coder:)
func NewOpenPanelWithCoder(coder foundation.INSCoder) NSOpenPanel {
	instance := getNSOpenPanelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSOpenPanelFromID(rv)
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
func NewOpenPanelWithContentRectStyleMaskBackingDefer(contentRect corefoundation.CGRect, style NSWindowStyleMask, backingStoreType NSBackingStoreType, flag bool) NSOpenPanel {
	instance := getNSOpenPanelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return NSOpenPanelFromID(rv)
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
func NewOpenPanelWithContentRectStyleMaskBackingDeferScreen(contentRect corefoundation.CGRect, style NSWindowStyleMask, backingStoreType NSBackingStoreType, flag bool, screen INSScreen) NSOpenPanel {
	instance := getNSOpenPanelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, screen)
	return NSOpenPanelFromID(rv)
}

// A Boolean that indicates whether the user can choose files in the panel.
//
// # Discussion
// 
// When the value of this property is [true], users can choose files in the
// panel.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenPanel/canChooseFiles
func (o NSOpenPanel) CanChooseFiles() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("canChooseFiles"))
	return rv
}
func (o NSOpenPanel) SetCanChooseFiles(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setCanChooseFiles:"), value)
}
// A Boolean that indicates whether the user can choose directories in the
// panel.
//
// # Discussion
// 
// When the value of this property is [true], users can choose directories in
// the panel.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenPanel/canChooseDirectories
func (o NSOpenPanel) CanChooseDirectories() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("canChooseDirectories"))
	return rv
}
func (o NSOpenPanel) SetCanChooseDirectories(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setCanChooseDirectories:"), value)
}
// A Boolean that indicates whether the panel resolves aliases.
//
// # Discussion
// 
// When the value of this property is [true], dropping an alias on the panel
// or asking for filenames or URLs returns the resolved aliases. The default
// value of this property is [true]. When this value is [false], selecting an
// alias returns the alias instead of the file or directory it represents.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenPanel/resolvesAliases
func (o NSOpenPanel) ResolvesAliases() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("resolvesAliases"))
	return rv
}
func (o NSOpenPanel) SetResolvesAliases(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setResolvesAliases:"), value)
}
// A Boolean that indicates whether the user may select multiple files and
// directories.
//
// # Discussion
// 
// When the value of this property is [true], the user may select multiple
// items from the browser. When the selection contains multiple items, use the
// [URLs] property to retrieve those items instead of the inherited [URL]
// property.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenPanel/allowsMultipleSelection
func (o NSOpenPanel) AllowsMultipleSelection() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("allowsMultipleSelection"))
	return rv
}
func (o NSOpenPanel) SetAllowsMultipleSelection(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAllowsMultipleSelection:"), value)
}
// A Boolean value that indicates whether the panel’s accessory view is
// visible.
//
// # Discussion
// 
// The value of this property is [true] when the accessory view is visible,
// and [false] when it isn’t. Setting the value of this property
// programmatically changes the visibility of the accessory panel. If no
// accessory panel is present, setting this property does nothing.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenPanel/isAccessoryViewDisclosed
func (o NSOpenPanel) AccessoryViewDisclosed() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("isAccessoryViewDisclosed"))
	return rv
}
func (o NSOpenPanel) SetAccessoryViewDisclosed(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setAccessoryViewDisclosed:"), value)
}
// An array of URLs, each of which contains the fully specified location of a
// selected file or directory.
//
// # Discussion
// 
// This property contains a valid value when the user selects one item or
// multiple items.
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenPanel/urls
func (o NSOpenPanel) URLs() []foundation.NSURL {
	rv := objc.Send[[]objc.ID](o.ID, objc.Sel("URLs"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSURL {
		return foundation.NSURLFromID(id)
	})
}
// A Boolean value that indicates how the panel responds to iCloud documents
// that aren’t fully downloaded locally.
//
// # Discussion
// 
// When the value of this property is [true], the panel disallows opening
// non-local iCloud files. If the user selects a non-local file, the panel
// attempts to download that file. When the value of this property is [false],
// the user may select and open non-local files. Your app is responsible for
// downloading the files and reporting progress or any issues.
// 
// The default value of this property is [true], except for applications
// linked against the OS X v10.9 SDK or earlier that have adopted iCloud by
// specifying a ubiquitous container identifier entitlement.
// 
// For a better user experience, set this property to [false] and download the
// file’s contents with an [NSFileCoordinator] object. Show the dlownload
// progress using a [Progress] or [NSMetadataQuery] object.
//
// [NSFileCoordinator]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator
// [NSMetadataQuery]: https://developer.apple.com/documentation/Foundation/NSMetadataQuery
// [Progress]: https://developer.apple.com/documentation/Foundation/Progress
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenPanel/canDownloadUbiquitousContents
func (o NSOpenPanel) CanDownloadUbiquitousContents() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("canDownloadUbiquitousContents"))
	return rv
}
func (o NSOpenPanel) SetCanDownloadUbiquitousContents(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setCanDownloadUbiquitousContents:"), value)
}
// A Boolean value that indicates how the panel responds to iCloud documents
// that have conflicting versions.
//
// # Discussion
// 
// When the value of this property is [true], and the user attempts to open
// one or more documents with conflicts, the panel displays the conflict
// resolution UI. The user must resolve any conflicts before opening the
// documents. When the value of this property is [false], the your application
// is responsible for handling any conflicts.
// 
// The default value of this property is [true], except for applications
// linked against the OS X v10.9 SDK or earlier that have adopted iCloud by
// specifying a ubiquitous container identifier entitlement.
// 
// For a better user experience, set this property to [false] and check the
// [ubiquitousItemHasUnresolvedConflictsKey] key of each item. When a conflict
// exists, retrieve a [NSFileVersion] object for each version and present your
// own UI to resolve that conflict.
//
// [NSFileVersion]: https://developer.apple.com/documentation/Foundation/NSFileVersion
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
// [ubiquitousItemHasUnresolvedConflictsKey]: https://developer.apple.com/documentation/Foundation/URLResourceKey/ubiquitousItemHasUnresolvedConflictsKey
//
// See: https://developer.apple.com/documentation/AppKit/NSOpenPanel/canResolveUbiquitousConflicts
func (o NSOpenPanel) CanResolveUbiquitousConflicts() bool {
	rv := objc.Send[bool](o.ID, objc.Sel("canResolveUbiquitousConflicts"))
	return rv
}
func (o NSOpenPanel) SetCanResolveUbiquitousConflicts(value bool) {
	objc.Send[struct{}](o.ID, objc.Sel("setCanResolveUbiquitousConflicts:"), value)
}

