// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/uniformtypeidentifiers"
)

// The class instance for the [NSSavePanel] class.
var (
	_NSSavePanelClass     NSSavePanelClass
	_NSSavePanelClassOnce sync.Once
)

func getNSSavePanelClass() NSSavePanelClass {
	_NSSavePanelClassOnce.Do(func() {
		_NSSavePanelClass = NSSavePanelClass{class: objc.GetClass("NSSavePanel")}
	})
	return _NSSavePanelClass
}

// GetNSSavePanelClass returns the class object for NSSavePanel.
func GetNSSavePanelClass() NSSavePanelClass {
	return getNSSavePanelClass()
}

type NSSavePanelClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSSavePanelClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSSavePanelClass) Alloc() NSSavePanel {
	rv := objc.Send[NSSavePanel](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A panel that prompts the user for information about where to save a file.
//
// # Overview
//
// The Save panel provides an interface for specifying the location to save a
// file and the name of that file. You present this panel when the user
// attempts to save a new document, or when the user saves a copy of an
// existing document to a new location. The panel includes UI for browsing the
// file system, selecting a directory, and specifying the new name for the
// file. You can also add custom UI for your app using an accessory view.
//
// An [NSSavePanel] object reports user interactions to its associated
// [NSSavePanel.Delegate] object, which must adopt the [NSOpenSavePanelDelegate] protocol.
// Use your delegate object to validate the user’s selection and respond to
// user interactions with the panel.
//
// In macOS 10.15, the system always displays the Save dialog in a separate
// process, regardless of whether the app is sandboxed. When the user saves
// the document, macOS adds the saved file to the app’s sandbox (if
// necessary) so that the app can write to the file. Prior to macOS 10.15, the
// system used a separate process only for sandboxed apps.
//
// # Showing the Panel
//
//   - [NSSavePanel.BeginSheetModalForWindowCompletionHandler]: Presents the panel as a sheet modal to the specified window.
//   - [NSSavePanel.BeginWithCompletionHandler]: Presents the panel as a modeless window.
//   - [NSSavePanel.RunModal]: Displays the panel and begins its event loop with the current working (or last-selected) directory as the default starting point.
//   - [NSSavePanel.ValidateVisibleColumns]: Validates and reloads the browser columns visible in the panel.
//
// # Getting the Selected Item
//
//   - [NSSavePanel.URL]: A URL that contains the fully specified location of the targeted file.
//
// # Configuring the Panel’s Appearance
//
//   - [NSSavePanel.Prompt]: The text to display in the default button.
//   - [NSSavePanel.SetPrompt]
//   - [NSSavePanel.Message]: The message text displayed in the panel.
//   - [NSSavePanel.SetMessage]
//   - [NSSavePanel.NameFieldLabel]: The label text displayed in front of the filename text field.
//   - [NSSavePanel.SetNameFieldLabel]
//   - [NSSavePanel.NameFieldStringValue]: The user-editable filename currently shown in the name field.
//   - [NSSavePanel.SetNameFieldStringValue]
//   - [NSSavePanel.DirectoryURL]: The current directory shown in the panel.
//   - [NSSavePanel.SetDirectoryURL]
//   - [NSSavePanel.AccessoryView]: The custom accessory view for the current app.
//   - [NSSavePanel.SetAccessoryView]
//   - [NSSavePanel.ShowsTagField]: A Boolean value that indicates whether the panel displays the Tags field.
//   - [NSSavePanel.SetShowsTagField]
//   - [NSSavePanel.TagNames]: The tag names that you want to include on a saved file.
//   - [NSSavePanel.SetTagNames]
//
// # Configuring the Panel’s Behavior
//
//   - [NSSavePanel.CanCreateDirectories]: A Boolean value that indicates whether the panel displays UI for creating directories.
//   - [NSSavePanel.SetCanCreateDirectories]
//   - [NSSavePanel.CanSelectHiddenExtension]: A Boolean value that indicates whether the panel displays UI for hiding or showing filename extensions.
//   - [NSSavePanel.SetCanSelectHiddenExtension]
//   - [NSSavePanel.ShowsHiddenFiles]: A Boolean value that indicates whether the panel displays files that are normally hidden from the user.
//   - [NSSavePanel.SetShowsHiddenFiles]
//   - [NSSavePanel.ExtensionHidden]: A Boolean value that indicates whether to display filename extensions.
//   - [NSSavePanel.SetExtensionHidden]
//   - [NSSavePanel.Expanded]: A Boolean value that indicates whether whether the panel is expanded.
//
// # Configuring the File Types
//
//   - [NSSavePanel.AllowedContentTypes]: An array of types that specify the files types to which you can save.
//   - [NSSavePanel.SetAllowedContentTypes]
//   - [NSSavePanel.AllowsOtherFileTypes]: A Boolean value that indicates whether the panel allows the user to save files with a filename extension that’s not in the list of allowed types.
//   - [NSSavePanel.SetAllowsOtherFileTypes]
//   - [NSSavePanel.TreatsFilePackagesAsDirectories]: A Boolean value that indicates whether the panel displays file packages as directories.
//   - [NSSavePanel.SetTreatsFilePackagesAsDirectories]
//
// # Handling Actions
//
//   - [NSSavePanel.Ok]: The action method that the panel calls when the user clicks the OK button.
//   - [NSSavePanel.Cancel]: The action method that the panel calls when the user clicks the Cancel button.
//
// # Instance Properties
//
//   - [NSSavePanel.CurrentContentType]: [NSSavePanel]:The current type. If set to `nil`, resets to the first allowed content type. Returns `nil` if `allowedContentTypes` is empty. [NSOpenPanel]: Not used.
//   - [NSSavePanel.SetCurrentContentType]
//   - [NSSavePanel.ShowsContentTypes]: [NSSavePanel]: Whether or not to show a control for selecting the type of the saved file. The control shows the types in `allowedContentTypes`. Default is [NO]. [NSOpenPanel]: Not used.
//   - [NSSavePanel.SetShowsContentTypes]
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel
type NSSavePanel struct {
	NSPanel
}

// NSSavePanelFromID constructs a [NSSavePanel] from an objc.ID.
//
// A panel that prompts the user for information about where to save a file.
func NSSavePanelFromID(id objc.ID) NSSavePanel {
	return NSSavePanel{NSPanel: NSPanelFromID(id)}
}

// NOTE: NSSavePanel adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSSavePanel] class.
//
// # Showing the Panel
//
//   - [INSSavePanel.BeginSheetModalForWindowCompletionHandler]: Presents the panel as a sheet modal to the specified window.
//   - [INSSavePanel.BeginWithCompletionHandler]: Presents the panel as a modeless window.
//   - [INSSavePanel.RunModal]: Displays the panel and begins its event loop with the current working (or last-selected) directory as the default starting point.
//   - [INSSavePanel.ValidateVisibleColumns]: Validates and reloads the browser columns visible in the panel.
//
// # Getting the Selected Item
//
//   - [INSSavePanel.URL]: A URL that contains the fully specified location of the targeted file.
//
// # Configuring the Panel’s Appearance
//
//   - [INSSavePanel.Prompt]: The text to display in the default button.
//   - [INSSavePanel.SetPrompt]
//   - [INSSavePanel.Message]: The message text displayed in the panel.
//   - [INSSavePanel.SetMessage]
//   - [INSSavePanel.NameFieldLabel]: The label text displayed in front of the filename text field.
//   - [INSSavePanel.SetNameFieldLabel]
//   - [INSSavePanel.NameFieldStringValue]: The user-editable filename currently shown in the name field.
//   - [INSSavePanel.SetNameFieldStringValue]
//   - [INSSavePanel.DirectoryURL]: The current directory shown in the panel.
//   - [INSSavePanel.SetDirectoryURL]
//   - [INSSavePanel.AccessoryView]: The custom accessory view for the current app.
//   - [INSSavePanel.SetAccessoryView]
//   - [INSSavePanel.ShowsTagField]: A Boolean value that indicates whether the panel displays the Tags field.
//   - [INSSavePanel.SetShowsTagField]
//   - [INSSavePanel.TagNames]: The tag names that you want to include on a saved file.
//   - [INSSavePanel.SetTagNames]
//
// # Configuring the Panel’s Behavior
//
//   - [INSSavePanel.CanCreateDirectories]: A Boolean value that indicates whether the panel displays UI for creating directories.
//   - [INSSavePanel.SetCanCreateDirectories]
//   - [INSSavePanel.CanSelectHiddenExtension]: A Boolean value that indicates whether the panel displays UI for hiding or showing filename extensions.
//   - [INSSavePanel.SetCanSelectHiddenExtension]
//   - [INSSavePanel.ShowsHiddenFiles]: A Boolean value that indicates whether the panel displays files that are normally hidden from the user.
//   - [INSSavePanel.SetShowsHiddenFiles]
//   - [INSSavePanel.ExtensionHidden]: A Boolean value that indicates whether to display filename extensions.
//   - [INSSavePanel.SetExtensionHidden]
//   - [INSSavePanel.Expanded]: A Boolean value that indicates whether whether the panel is expanded.
//
// # Configuring the File Types
//
//   - [INSSavePanel.AllowedContentTypes]: An array of types that specify the files types to which you can save.
//   - [INSSavePanel.SetAllowedContentTypes]
//   - [INSSavePanel.AllowsOtherFileTypes]: A Boolean value that indicates whether the panel allows the user to save files with a filename extension that’s not in the list of allowed types.
//   - [INSSavePanel.SetAllowsOtherFileTypes]
//   - [INSSavePanel.TreatsFilePackagesAsDirectories]: A Boolean value that indicates whether the panel displays file packages as directories.
//   - [INSSavePanel.SetTreatsFilePackagesAsDirectories]
//
// # Handling Actions
//
//   - [INSSavePanel.Ok]: The action method that the panel calls when the user clicks the OK button.
//   - [INSSavePanel.Cancel]: The action method that the panel calls when the user clicks the Cancel button.
//
// # Instance Properties
//
//   - [INSSavePanel.CurrentContentType]: [NSSavePanel]:The current type. If set to `nil`, resets to the first allowed content type. Returns `nil` if `allowedContentTypes` is empty. [NSOpenPanel]: Not used.
//   - [INSSavePanel.SetCurrentContentType]
//   - [INSSavePanel.ShowsContentTypes]: [NSSavePanel]: Whether or not to show a control for selecting the type of the saved file. The control shows the types in `allowedContentTypes`. Default is [NO]. [NSOpenPanel]: Not used.
//   - [INSSavePanel.SetShowsContentTypes]
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel
type INSSavePanel interface {
	INSPanel

	// Topic: Showing the Panel

	// Presents the panel as a sheet modal to the specified window.
	BeginSheetModalForWindowCompletionHandler(window INSWindow, handler ModalResponseHandler)
	// Presents the panel as a modeless window.
	BeginWithCompletionHandler(handler ModalResponseHandler)
	// Displays the panel and begins its event loop with the current working (or last-selected) directory as the default starting point.
	RunModal() NSModalResponse
	// Validates and reloads the browser columns visible in the panel.
	ValidateVisibleColumns()

	// Topic: Getting the Selected Item

	// A URL that contains the fully specified location of the targeted file.
	URL() foundation.INSURL

	// Topic: Configuring the Panel’s Appearance

	// The text to display in the default button.
	Prompt() string
	SetPrompt(value string)
	// The message text displayed in the panel.
	Message() string
	SetMessage(value string)
	// The label text displayed in front of the filename text field.
	NameFieldLabel() string
	SetNameFieldLabel(value string)
	// The user-editable filename currently shown in the name field.
	NameFieldStringValue() string
	SetNameFieldStringValue(value string)
	// The current directory shown in the panel.
	DirectoryURL() foundation.INSURL
	SetDirectoryURL(value foundation.INSURL)
	// The custom accessory view for the current app.
	AccessoryView() INSView
	SetAccessoryView(value INSView)
	// A Boolean value that indicates whether the panel displays the Tags field.
	ShowsTagField() bool
	SetShowsTagField(value bool)
	// The tag names that you want to include on a saved file.
	TagNames() []string
	SetTagNames(value []string)

	// Topic: Configuring the Panel’s Behavior

	// A Boolean value that indicates whether the panel displays UI for creating directories.
	CanCreateDirectories() bool
	SetCanCreateDirectories(value bool)
	// A Boolean value that indicates whether the panel displays UI for hiding or showing filename extensions.
	CanSelectHiddenExtension() bool
	SetCanSelectHiddenExtension(value bool)
	// A Boolean value that indicates whether the panel displays files that are normally hidden from the user.
	ShowsHiddenFiles() bool
	SetShowsHiddenFiles(value bool)
	// A Boolean value that indicates whether to display filename extensions.
	ExtensionHidden() bool
	SetExtensionHidden(value bool)
	// A Boolean value that indicates whether whether the panel is expanded.
	Expanded() bool

	// Topic: Configuring the File Types

	// An array of types that specify the files types to which you can save.
	AllowedContentTypes() []uniformtypeidentifiers.UTType
	SetAllowedContentTypes(value []uniformtypeidentifiers.UTType)
	// A Boolean value that indicates whether the panel allows the user to save files with a filename extension that’s not in the list of allowed types.
	AllowsOtherFileTypes() bool
	SetAllowsOtherFileTypes(value bool)
	// A Boolean value that indicates whether the panel displays file packages as directories.
	TreatsFilePackagesAsDirectories() bool
	SetTreatsFilePackagesAsDirectories(value bool)

	// Topic: Handling Actions

	// The action method that the panel calls when the user clicks the OK button.
	Ok(sender objectivec.IObject)
	// The action method that the panel calls when the user clicks the Cancel button.
	Cancel(sender objectivec.IObject)

	// Topic: Instance Properties

	// [NSSavePanel]:The current type. If set to `nil`, resets to the first allowed content type. Returns `nil` if `allowedContentTypes` is empty. [NSOpenPanel]: Not used.
	CurrentContentType() uniformtypeidentifiers.UTType
	SetCurrentContentType(value uniformtypeidentifiers.UTType)
	// [NSSavePanel]: Whether or not to show a control for selecting the type of the saved file. The control shows the types in `allowedContentTypes`. Default is [NO]. [NSOpenPanel]: Not used.
	ShowsContentTypes() bool
	SetShowsContentTypes(value bool)
}

// Init initializes the instance.
func (s NSSavePanel) Init() NSSavePanel {
	rv := objc.Send[NSSavePanel](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSSavePanel) Autorelease() NSSavePanel {
	rv := objc.Send[NSSavePanel](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSSavePanel creates a new NSSavePanel instance.
func NewNSSavePanel() NSSavePanel {
	class := getNSSavePanelClass()
	rv := objc.Send[NSSavePanel](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a titled window that contains the specified content view
// controller.
//
// contentViewController: The view controller that provides the main content view for the window. The
// window’s [ContentView] property is set to
// `contentViewController“XCUIElementTypeView`.
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
// `contentViewController“XCUIElementTypeView`). The newly created window has
// [ReleasedWhenClosed] set to false, and it must be explicitly retained to
// keep the window instance alive.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindow/init(contentViewController:)
func NewSavePanelWindowWithContentViewController(contentViewController INSViewController) NSSavePanel {
	rv := objc.Send[objc.ID](objc.ID(getNSSavePanelClass().class), objc.Sel("windowWithContentViewController:"), contentViewController)
	return NSSavePanelFromID(rv)
}

// Creates a new responder object with data in an unarchiver.
//
// coder: An unarchiver object.
//
// See: https://developer.apple.com/documentation/AppKit/NSResponder/init(coder:)
func NewSavePanelWithCoder(coder foundation.INSCoder) NSSavePanel {
	instance := getNSSavePanelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSSavePanelFromID(rv)
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
//
// backingStoreType: Specifies how the drawing done in the window is buffered by the window
// device, and possible values are described in [NSWindow.BackingStoreType].
//
// flag: Specifies whether the window server creates a window device for the window
// immediately. When true, the window server defers creating the window device
// until the window is moved onscreen. All display messages sent to the window
// or its views are postponed until the window is created, just before it’s
// moved onscreen.
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
//
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
// [NSWindow.BackingStoreType]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType
func NewSavePanelWithContentRectStyleMaskBackingDefer(contentRect corefoundation.CGRect, style NSWindowStyleMask, backingStoreType NSBackingStoreType, flag bool) NSSavePanel {
	instance := getNSSavePanelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentRect:styleMask:backing:defer:"), contentRect, style, backingStoreType, flag)
	return NSSavePanelFromID(rv)
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
//
// backingStoreType: Specifies how the drawing done in the window is buffered by the window
// device; possible values are described in [NSWindow.BackingStoreType].
//
// flag: Specifies whether the window server creates a window device for the window
// immediately. When true, the window server defers creating the window device
// until the window is moved onscreen. All display messages sent to the window
// or its views are postponed until the window is created, just before it’s
// moved onscreen.
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
//
// [NSWindow.StyleMask]: https://developer.apple.com/documentation/AppKit/NSWindow/StyleMask-swift.struct
// [NSWindow.BackingStoreType]: https://developer.apple.com/documentation/AppKit/NSWindow/BackingStoreType
func NewSavePanelWithContentRectStyleMaskBackingDeferScreen(contentRect corefoundation.CGRect, style NSWindowStyleMask, backingStoreType NSBackingStoreType, flag bool, screen INSScreen) NSSavePanel {
	instance := getNSSavePanelClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentRect:styleMask:backing:defer:screen:"), contentRect, style, backingStoreType, flag, screen)
	return NSSavePanelFromID(rv)
}

// Presents the panel as a sheet modal to the specified window.
//
// window: The window in which the panel will be presented.
//
// handler: The block called after the user dismisses the panel. The argument passed in
// will be [NSFileHandlingPanelOKButton] if the user chose the OK button or
// [NSFileHandlingPanelCancelButton] if the user chose the Cancel button.
//
// # Discussion
//
// Configure all the relevant properties of the panel before you call this
// method. The completion handler block runs after the user dismisses the
// panel, but while the panel may still be onscreen. If you need to dismiss
// the panel from the screen—for example, if the completion block displays
// an alert—close the panel by calling its [OrderOut] method with the value
// `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/beginSheetModal(for:completionHandler:)
func (s NSSavePanel) BeginSheetModalForWindowCompletionHandler(window INSWindow, handler ModalResponseHandler) {
	_block1, _ := NewModalResponseBlock(handler)
	objc.Send[objc.ID](s.ID, objc.Sel("beginSheetModalForWindow:completionHandler:"), window, _block1)
}

// Presents the panel as a modeless window.
//
// handler: The block to call after the user closes the panel. This block has no return
// value and takes a single parameter:
//
// result: The action taken by the user. The value of this parameter is
// [NSFileHandlingPanelOKButton] if the user chose the OK button or
// [NSFileHandlingPanelCancelButton] if the user chose the Cancel button.
//
// # Discussion
//
// Configure all of the relevant properties of the panel before you call this
// method.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/begin(completionHandler:)
func (s NSSavePanel) BeginWithCompletionHandler(handler ModalResponseHandler) {
	_block0, _ := NewModalResponseBlock(handler)
	objc.Send[objc.ID](s.ID, objc.Sel("beginWithCompletionHandler:"), _block0)
}

// Displays the panel and begins its event loop with the current working (or
// last-selected) directory as the default starting point.
//
// # Return Value
//
// [NSFileHandlingPanelOKButton] (if the user clicks the OK button) or
// [NSFileHandlingPanelCancelButton] (if the user clicks the Cancel button).
//
// # Discussion
//
// This method invokes [NSApplication]’s [RunModalForWindow] method with
// `self` as the argument.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/runModal()
func (s NSSavePanel) RunModal() NSModalResponse {
	rv := objc.Send[NSModalResponse](s.ID, objc.Sel("runModal"))
	return NSModalResponse(rv)
}

// Validates and reloads the browser columns visible in the panel.
//
// # Discussion
//
// Call this method to validate the contents of the panel. For example, you
// might call it to allow the selection of files with certain extensions based
// on the selection made in an accessory-view pop-up list. When the user
// changes the selection, invoke this method to revalidate the visible
// columns.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/validateVisibleColumns()
func (s NSSavePanel) ValidateVisibleColumns() {
	objc.Send[objc.ID](s.ID, objc.Sel("validateVisibleColumns"))
}

// The action method that the panel calls when the user clicks the OK button.
//
// sender: The [NSSavePanel] object that contains the OK button.
//
// # Discussion
//
// In macOS 10.15 and later, you cannot call this method programmatically to
// trigger the OK action. Prior to macOS 10.15, AppKit prevented only
// sandboxed apps from calling this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/ok(_:)
func (s NSSavePanel) Ok(sender objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("ok:"), sender)
}

// The action method that the panel calls when the user clicks the Cancel
// button.
//
// sender: The [NSSavePanel] object that contains the Cancel button.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/cancel(_:)
func (s NSSavePanel) Cancel(sender objectivec.IObject) {
	objc.Send[objc.ID](s.ID, objc.Sel("cancel:"), sender)
}

// Creates a new Save panel and initializes it with default information.
//
// # Return Value
//
// The initialized Save panel.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/savePanel
func (_NSSavePanelClass NSSavePanelClass) SavePanel() NSSavePanel {
	rv := objc.Send[objc.ID](objc.ID(_NSSavePanelClass.class), objc.Sel("savePanel"))
	return NSSavePanelFromID(rv)
}

// A URL that contains the fully specified location of the targeted file.
//
// # Discussion
//
// The [NSOpenPanel] subclass sets this property to `nil` when the selection
// contains multiple items.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/url
func (s NSSavePanel) URL() foundation.INSURL {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("URL"))
	return foundation.NSURLFromID(objc.ID(rv))
}

// The text to display in the default button.
//
// # Discussion
//
// The prompt appears on all [NSSavePanel] objects (or all [NSOpenPanel]
// objects if this property is on an [NSOpenPanel] instance) in your
// application. By default, the text in the default button is “Open” for
// an Open panel and “Save” for a Save panel.
//
// Use short words or phrases, such as “Open,” “Save,” “Set,” or
// “Choose,” on the button. The button is not resized to accommodate long
// prompts.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/prompt
func (s NSSavePanel) Prompt() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("prompt"))
	return foundation.NSStringFromID(rv).String()
}
func (s NSSavePanel) SetPrompt(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setPrompt:"), objc.String(value))
}

// The message text displayed in the panel.
//
// # Discussion
//
// This prompt appears on all [NSSavePanel] objects (or all [NSOpenPanel]
// objects if this property is on an [NSOpenPanel] instance) in your
// application. The default message text is an empty string.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/message
func (s NSSavePanel) Message() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("message"))
	return foundation.NSStringFromID(rv).String()
}
func (s NSSavePanel) SetMessage(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setMessage:"), objc.String(value))
}

// The label text displayed in front of the filename text field.
//
// # Discussion
//
// The default value of this property is “Save As:”.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/nameFieldLabel
func (s NSSavePanel) NameFieldLabel() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("nameFieldLabel"))
	return foundation.NSStringFromID(rv).String()
}
func (s NSSavePanel) SetNameFieldLabel(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setNameFieldLabel:"), objc.String(value))
}

// The user-editable filename currently shown in the name field.
//
// # Discussion
//
// The value of this property must not be `nil`. Note that the filename may
// not display an extension if the value of [ExtensionHidden] is true.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/nameFieldStringValue
func (s NSSavePanel) NameFieldStringValue() string {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("nameFieldStringValue"))
	return foundation.NSStringFromID(rv).String()
}
func (s NSSavePanel) SetNameFieldStringValue(value string) {
	objc.Send[struct{}](s.ID, objc.Sel("setNameFieldStringValue:"), objc.String(value))
}

// The current directory shown in the panel.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/directoryURL
func (s NSSavePanel) DirectoryURL() foundation.INSURL {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("directoryURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (s NSSavePanel) SetDirectoryURL(value foundation.INSURL) {
	objc.Send[struct{}](s.ID, objc.Sel("setDirectoryURL:"), value)
}

// The custom accessory view for the current app.
//
// # Discussion
//
// You can customize the panel by adding a custom view. The custom object you
// add appears just above the OK and Cancel buttons at the bottom of the
// panel. The [NSSavePanel] object automatically resizes itself to accommodate
// `accessoryView`. Use this property to change the accessory view as needed.
// If `accessoryView` is `nil`, the Save panel removes the current accessory
// view.
//
// The panel relinquishes ownership of the accessory view after the panel is
// closed. If you want to reuse the accessory view, don’t rely on the panel
// to hold onto the accessory view until the next time you use it; instead,
// maintain your own strong reference to the view.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/accessoryView
func (s NSSavePanel) AccessoryView() INSView {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("accessoryView"))
	return NSViewFromID(objc.ID(rv))
}
func (s NSSavePanel) SetAccessoryView(value INSView) {
	objc.Send[struct{}](s.ID, objc.Sel("setAccessoryView:"), value)
}

// A Boolean value that indicates whether the panel displays the Tags field.
//
// # Discussion
//
// When the value of this property is true, the panel displays the Tags field;
// if false, the panel doesn’t display the Tags field. The default value is
// true. (Note that the Tags field is appropriate only in a Save panel.)
//
// If you set this property to true, you are responsible for setting tag names
// on the resulting file after saving is complete. If you don’t set this
// property, macOS will automatically show the tag field and attempt to apply
// the tags to the file. To set tags on files, use the [tagNamesKey].
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/showsTagField
//
// [tagNamesKey]: https://developer.apple.com/documentation/Foundation/URLResourceKey/tagNamesKey
func (s NSSavePanel) ShowsTagField() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("showsTagField"))
	return rv
}
func (s NSSavePanel) SetShowsTagField(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setShowsTagField:"), value)
}

// The tag names that you want to include on a saved file.
//
// # Discussion
//
// When the value of [ShowsTagField] is true, use this property to provide an
// array of strings that represent the initial tag names to display in the
// panel. If you set the property to `nil` or an empty array, the panel
// displays no initial tag names.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/tagNames
func (s NSSavePanel) TagNames() []string {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("tagNames"))
	return objc.ConvertSliceToStrings(rv)
}
func (s NSSavePanel) SetTagNames(value []string) {
	objc.Send[struct{}](s.ID, objc.Sel("setTagNames:"), objectivec.StringSliceToNSArray(value))
}

// A Boolean value that indicates whether the panel displays UI for creating
// directories.
//
// # Discussion
//
// When the value of this property is true, the panel includes UI to create
// new directories. When the value is false, the panel does not expose that
// UI.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/canCreateDirectories
func (s NSSavePanel) CanCreateDirectories() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("canCreateDirectories"))
	return rv
}
func (s NSSavePanel) SetCanCreateDirectories(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setCanCreateDirectories:"), value)
}

// A Boolean value that indicates whether the panel displays UI for hiding or
// showing filename extensions.
//
// # Discussion
//
// Set the value of this property before displaying the panel. When the value
// of this property is true, and the Finder preference “Show all
// extensions” is false, the panel displays the Hide Extension menu item.
// The default value of this property is false.
//
// Use the [ExtensionHidden] property to hide or shows extensions.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/canSelectHiddenExtension
func (s NSSavePanel) CanSelectHiddenExtension() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("canSelectHiddenExtension"))
	return rv
}
func (s NSSavePanel) SetCanSelectHiddenExtension(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setCanSelectHiddenExtension:"), value)
}

// A Boolean value that indicates whether the panel displays files that are
// normally hidden from the user.
//
// # Discussion
//
// When the value of this property is true, the panel displays hidden files;
// if false, it does not. The default value is false.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/showsHiddenFiles
func (s NSSavePanel) ShowsHiddenFiles() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("showsHiddenFiles"))
	return rv
}
func (s NSSavePanel) SetShowsHiddenFiles(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setShowsHiddenFiles:"), value)
}

// A Boolean value that indicates whether to display filename extensions.
//
// # Discussion
//
// When the value of this property is false, [NSSavePanel] shows the filename
// extension in places where you refer to the file by name. The user can
// override this value by checking the hide extension menu item, which
// reflects this value. The default value is true.
//
// If a user adds or removes a filename extension in the panel’s name field,
// the panel updates this property to reflect that choice.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/isExtensionHidden
func (s NSSavePanel) ExtensionHidden() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isExtensionHidden"))
	return rv
}
func (s NSSavePanel) SetExtensionHidden(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setExtensionHidden:"), value)
}

// A Boolean value that indicates whether whether the panel is expanded.
//
// # Discussion
//
// The value of this property is true if the panel is expanded; otherwise,
// false.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/isExpanded
func (s NSSavePanel) Expanded() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("isExpanded"))
	return rv
}

// An array of types that specify the files types to which you can save.
//
// # Discussion
//
// Defaults to an empty array that indicates that you can use any file type.
// If you don’t provide an extension, the system uses the first preferred
// extension in the array for the save panel. If you specify a type that
// isn’t in the array and [AllowsOtherFileTypes] is [YES], the system
// presents another dialog when prompting you to save.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/allowedContentTypes
func (s NSSavePanel) AllowedContentTypes() []uniformtypeidentifiers.UTType {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("allowedContentTypes"))
	return objc.ConvertSlice(rv, func(id objc.ID) uniformtypeidentifiers.UTType {
		return uniformtypeidentifiers.UTTypeFromID(id)
	})
}
func (s NSSavePanel) SetAllowedContentTypes(value []uniformtypeidentifiers.UTType) {
	objc.Send[struct{}](s.ID, objc.Sel("setAllowedContentTypes:"), objectivec.IObjectSliceToNSArray(value))
}

// A Boolean value that indicates whether the panel allows the user to save
// files with a filename extension that’s not in the list of allowed types.
//
// # Discussion
//
// When the value of this property is true, the panel allows the user to save
// files with an extension that’s not in the list of allowed types. The
// default value is false.
//
// If the user tries to save a filename with a recognized extension that’s
// not in the list of allowed types, they are presented with a dialog. If the
// value of this property is true, then the dialog presents the option of
// using the extension the user specified.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/allowsOtherFileTypes
func (s NSSavePanel) AllowsOtherFileTypes() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("allowsOtherFileTypes"))
	return rv
}
func (s NSSavePanel) SetAllowsOtherFileTypes(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setAllowsOtherFileTypes:"), value)
}

// A Boolean value that indicates whether the panel displays file packages as
// directories.
//
// # Discussion
//
// When the value of this property is true, the panel displays file packages
// as directories; if false, it will not.
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/treatsFilePackagesAsDirectories
func (s NSSavePanel) TreatsFilePackagesAsDirectories() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("treatsFilePackagesAsDirectories"))
	return rv
}
func (s NSSavePanel) SetTreatsFilePackagesAsDirectories(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setTreatsFilePackagesAsDirectories:"), value)
}

// [NSSavePanel]:The current type. If set to `nil`, resets to the first
// allowed content type. Returns `nil` if `allowedContentTypes` is empty.
// [NSOpenPanel]: Not used.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/currentContentType
func (s NSSavePanel) CurrentContentType() uniformtypeidentifiers.UTType {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("currentContentType"))
	return uniformtypeidentifiers.UTTypeFromID(objc.ID(rv))
}
func (s NSSavePanel) SetCurrentContentType(value uniformtypeidentifiers.UTType) {
	objc.Send[struct{}](s.ID, objc.Sel("setCurrentContentType:"), value)
}

// [NSSavePanel]: Whether or not to show a control for selecting the type of
// the saved file. The control shows the types in `allowedContentTypes`.
// Default is [NO]. [NSOpenPanel]: Not used.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AppKit/NSSavePanel/showsContentTypes
func (s NSSavePanel) ShowsContentTypes() bool {
	rv := objc.Send[bool](s.ID, objc.Sel("showsContentTypes"))
	return rv
}
func (s NSSavePanel) SetShowsContentTypes(value bool) {
	objc.Send[struct{}](s.ID, objc.Sel("setShowsContentTypes:"), value)
}

// BeginSheetModalForWindow is a synchronous wrapper around [NSSavePanel.BeginSheetModalForWindowCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s NSSavePanel) BeginSheetModalForWindow(ctx context.Context, window INSWindow) (NSModalResponse, error) {
	done := make(chan NSModalResponse, 1)
	s.BeginSheetModalForWindowCompletionHandler(window, func(val NSModalResponse) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return *new(NSModalResponse), ctx.Err()
	}
}

// Begin is a synchronous wrapper around [NSSavePanel.BeginWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (s NSSavePanel) Begin(ctx context.Context) (NSModalResponse, error) {
	done := make(chan NSModalResponse, 1)
	s.BeginWithCompletionHandler(func(val NSModalResponse) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return *new(NSModalResponse), ctx.Err()
	}
}
