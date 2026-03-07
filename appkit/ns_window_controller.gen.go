// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSWindowController] class.
var (
	_NSWindowControllerClass     NSWindowControllerClass
	_NSWindowControllerClassOnce sync.Once
)

func getNSWindowControllerClass() NSWindowControllerClass {
	_NSWindowControllerClassOnce.Do(func() {
		_NSWindowControllerClass = NSWindowControllerClass{class: objc.GetClass("NSWindowController")}
	})
	return _NSWindowControllerClass
}

// GetNSWindowControllerClass returns the class object for NSWindowController.
func GetNSWindowControllerClass() NSWindowControllerClass {
	return getNSWindowControllerClass()
}

type NSWindowControllerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSWindowControllerClass) Alloc() NSWindowController {
	rv := objc.Send[NSWindowController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A controller that manages a window, usually a window stored in a nib file.
//
// # Overview
// 
// Managing a window entails:
// 
// - Loading and displaying the window - Closing the window when appropriate -
// Customizing the window’s title - Storing the window’s frame (size and
// location) in the defaults database - Cascading the window in relation to
// other document windows of the app
// 
// A window controller can manage a window by itself or as a role player in
// AppKit’s document-based architecture, which also includes [NSDocument]
// and [NSDocumentController] objects. In this architecture, a window
// controller is created and managed by a “document” (an instance of an
// [NSDocument] subclass) and, in turn, keeps a reference to the document.
// 
// The relationship between a window controller and a nib file is important.
// Although a window controller can manage a programmatically created window,
// it usually manages a window in a nib file. The nib file can contain other
// top-level objects, including other windows, but the window controller’s
// responsibility is this primary window. The window controller is usually the
// owner of the nib file, even when it is part of a document-based app.
// Regardless of who is the file’s owner, the window controller is
// responsible for freeing all top-level objects in the nib file it loads.
// 
// For simple documents—that is, documents with only one nib file containing
// a window—you need to do little directly with [NSWindowController]; AppKit
// creates one for you. However, if the default window controller is not
// sufficient, you can create a custom subclass of [NSWindowController]. For
// documents with multiple windows or panels, your document must create
// separate instances of [NSWindowController] (or of custom subclasses of
// [NSWindowController]), one for each window or panel. An example is a CAD
// app that has different windows for side, top, and front views of drawn
// objects. What you do in your [NSDocument] subclass determines whether the
// default [NSWindowController] or separately created and configured
// [NSWindowController] objects are used.
// 
// # Subclassing NSWindowController
// 
// You should create a subclass of [NSWindowController] when you want to
// augment the default behavior, such as to give the window a custom title or
// to perform some setup tasks before the window is loaded. In your class’s
// initialization method, be sure to invoke on `super` either one of the `...`
// initializers or the [NSWindowController.InitWithWindow] initializer. The initializer you
// choose depends on whether the window object originates in a nib file or is
// programmatically created.
// 
// You can also implement an [NSWindowController] subclass to avoid requiring
// client code to get the corresponding nib’s filename and pass it to
// [NSWindowController.InitWithWindowNibName] or [NSWindowController.InitWithWindowNibNameOwner] when instantiating
// the window controller. The best way to do this is to override
// [NSWindowController.WindowNibName] to return the nib’s filename and instantiate the window
// controller by passing `nil` to [NSWindowController.InitWithWindow]. Using the [NSWindowController.InitWithWindow]
// designated initializer simplifies compliance with Swift initializer
// requirements.
// 
// Typically, you override one of the methods listed below.
// 
// [Table data omitted]
// 
// You can also override [NSWindowController.LoadWindow] to get different nib-searching or
// nib-loading behavior, although there is usually no need to do this.
//
// # Initializing Window Controllers
//
//   - [NSWindowController.InitWithWindow]: Returns a window controller initialized with a given window.
//   - [NSWindowController.InitWithWindowNibName]: Returns a window controller initialized with a nib file.
//   - [NSWindowController.InitWithWindowNibNameOwner]: Returns a window controller initialized with a nib file and a specified owner for that nib file.
//   - [NSWindowController.InitWithWindowNibPathOwner]: Returns a window controller initialized with a nib file at an absolute path and a specified owner.
//
// # Loading and Displaying the Window
//
//   - [NSWindowController.LoadWindow]: Loads the receiver’s window from the nib file.
//   - [NSWindowController.ShowWindow]: Displays the window associated with the receiver.
//   - [NSWindowController.WindowLoaded]: A Boolean value that indicates whether the nib file containing the receiver’s window has been loaded.
//   - [NSWindowController.Window]: The window owned by the receiver.
//   - [NSWindowController.SetWindow]
//   - [NSWindowController.WindowDidLoad]: Sent after the window owned by the receiver has been loaded.
//   - [NSWindowController.WindowWillLoad]: Sent before the window owned by the receiver is loaded.
//
// # Accessing the Document
//
//   - [NSWindowController.Document]: The document associated with the window controller.
//   - [NSWindowController.SetDocument]
//   - [NSWindowController.SetDocumentEdited]: Sets the document edited flag for the window controller.
//
// # Closing the Window
//
//   - [NSWindowController.Close]: Closes the window if it was loaded.
//   - [NSWindowController.ShouldCloseDocument]: A Boolean value that indicates whether the receiver necessarily closes the associated document when the window it manages is closed.
//   - [NSWindowController.SetShouldCloseDocument]
//
// # Getting Nib and Storyboard Information
//
//   - [NSWindowController.Owner]: The owner of the nib file containing the window managed by the receiver.
//   - [NSWindowController.Storyboard]: The storyboard file from which the window controller was loaded.
//   - [NSWindowController.WindowNibName]: The name of the nib file that stores the window associated with the receiver.
//   - [NSWindowController.WindowNibPath]: The full path of the nib file that stores the window associated with the receiver.
//
// # Accessing Window Attributes and Content
//
//   - [NSWindowController.ShouldCascadeWindows]: A Boolean value that indicates whether the window will cascade in relation to other document windows when it is displayed.
//   - [NSWindowController.SetShouldCascadeWindows]
//   - [NSWindowController.WindowFrameAutosaveName]: The name under which the frame rectangle of the window owned by the receiver is stored in the defaults database.
//   - [NSWindowController.SetWindowFrameAutosaveName]
//   - [NSWindowController.SynchronizeWindowTitleWithDocumentName]: Synchronizes the displayed window title and the represented filename with the information in the associated document.
//   - [NSWindowController.WindowTitleForDocumentDisplayName]: Returns the window title to be used for a given document display name.
//   - [NSWindowController.ContentViewController]: The view controller for the window’s content view.
//   - [NSWindowController.SetContentViewController]
//   - [NSWindowController.DismissController]: Dismisses the window controller.
//
// # Instance Properties
//
//   - [NSWindowController.PreviewRepresentableActivityItems]
//   - [NSWindowController.SetPreviewRepresentableActivityItems]
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController
type NSWindowController struct {
	NSResponder
}

// NSWindowControllerFromID constructs a [NSWindowController] from an objc.ID.
//
// A controller that manages a window, usually a window stored in a nib file.
func NSWindowControllerFromID(id objc.ID) NSWindowController {
	return NSWindowController{NSResponder: NSResponderFromID(id)}
}
// NOTE: NSWindowController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSWindowController] class.
//
// # Initializing Window Controllers
//
//   - [INSWindowController.InitWithWindow]: Returns a window controller initialized with a given window.
//   - [INSWindowController.InitWithWindowNibName]: Returns a window controller initialized with a nib file.
//   - [INSWindowController.InitWithWindowNibNameOwner]: Returns a window controller initialized with a nib file and a specified owner for that nib file.
//   - [INSWindowController.InitWithWindowNibPathOwner]: Returns a window controller initialized with a nib file at an absolute path and a specified owner.
//
// # Loading and Displaying the Window
//
//   - [INSWindowController.LoadWindow]: Loads the receiver’s window from the nib file.
//   - [INSWindowController.ShowWindow]: Displays the window associated with the receiver.
//   - [INSWindowController.WindowLoaded]: A Boolean value that indicates whether the nib file containing the receiver’s window has been loaded.
//   - [INSWindowController.Window]: The window owned by the receiver.
//   - [INSWindowController.SetWindow]
//   - [INSWindowController.WindowDidLoad]: Sent after the window owned by the receiver has been loaded.
//   - [INSWindowController.WindowWillLoad]: Sent before the window owned by the receiver is loaded.
//
// # Accessing the Document
//
//   - [INSWindowController.Document]: The document associated with the window controller.
//   - [INSWindowController.SetDocument]
//   - [INSWindowController.SetDocumentEdited]: Sets the document edited flag for the window controller.
//
// # Closing the Window
//
//   - [INSWindowController.Close]: Closes the window if it was loaded.
//   - [INSWindowController.ShouldCloseDocument]: A Boolean value that indicates whether the receiver necessarily closes the associated document when the window it manages is closed.
//   - [INSWindowController.SetShouldCloseDocument]
//
// # Getting Nib and Storyboard Information
//
//   - [INSWindowController.Owner]: The owner of the nib file containing the window managed by the receiver.
//   - [INSWindowController.Storyboard]: The storyboard file from which the window controller was loaded.
//   - [INSWindowController.WindowNibName]: The name of the nib file that stores the window associated with the receiver.
//   - [INSWindowController.WindowNibPath]: The full path of the nib file that stores the window associated with the receiver.
//
// # Accessing Window Attributes and Content
//
//   - [INSWindowController.ShouldCascadeWindows]: A Boolean value that indicates whether the window will cascade in relation to other document windows when it is displayed.
//   - [INSWindowController.SetShouldCascadeWindows]
//   - [INSWindowController.WindowFrameAutosaveName]: The name under which the frame rectangle of the window owned by the receiver is stored in the defaults database.
//   - [INSWindowController.SetWindowFrameAutosaveName]
//   - [INSWindowController.SynchronizeWindowTitleWithDocumentName]: Synchronizes the displayed window title and the represented filename with the information in the associated document.
//   - [INSWindowController.WindowTitleForDocumentDisplayName]: Returns the window title to be used for a given document display name.
//   - [INSWindowController.ContentViewController]: The view controller for the window’s content view.
//   - [INSWindowController.SetContentViewController]
//   - [INSWindowController.DismissController]: Dismisses the window controller.
//
// # Instance Properties
//
//   - [INSWindowController.PreviewRepresentableActivityItems]
//   - [INSWindowController.SetPreviewRepresentableActivityItems]
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController
type INSWindowController interface {
	INSResponder
	NSSeguePerforming

	// Topic: Initializing Window Controllers

	// Returns a window controller initialized with a given window.
	InitWithWindow(window INSWindow) NSWindowController
	// Returns a window controller initialized with a nib file.
	InitWithWindowNibName(windowNibName NSNibName) NSWindowController
	// Returns a window controller initialized with a nib file and a specified owner for that nib file.
	InitWithWindowNibNameOwner(windowNibName NSNibName, owner objectivec.IObject) NSWindowController
	// Returns a window controller initialized with a nib file at an absolute path and a specified owner.
	InitWithWindowNibPathOwner(windowNibPath string, owner objectivec.IObject) NSWindowController

	// Topic: Loading and Displaying the Window

	// Loads the receiver’s window from the nib file.
	LoadWindow()
	// Displays the window associated with the receiver.
	ShowWindow(sender objectivec.IObject)
	// A Boolean value that indicates whether the nib file containing the receiver’s window has been loaded.
	WindowLoaded() bool
	// The window owned by the receiver.
	Window() INSWindow
	SetWindow(value INSWindow)
	// Sent after the window owned by the receiver has been loaded.
	WindowDidLoad()
	// Sent before the window owned by the receiver is loaded.
	WindowWillLoad()

	// Topic: Accessing the Document

	// The document associated with the window controller.
	Document() objectivec.IObject
	SetDocument(value objectivec.IObject)
	// Sets the document edited flag for the window controller.
	SetDocumentEdited(dirtyFlag bool)

	// Topic: Closing the Window

	// Closes the window if it was loaded.
	Close()
	// A Boolean value that indicates whether the receiver necessarily closes the associated document when the window it manages is closed.
	ShouldCloseDocument() bool
	SetShouldCloseDocument(value bool)

	// Topic: Getting Nib and Storyboard Information

	// The owner of the nib file containing the window managed by the receiver.
	Owner() objectivec.IObject
	// The storyboard file from which the window controller was loaded.
	Storyboard() INSStoryboard
	// The name of the nib file that stores the window associated with the receiver.
	WindowNibName() NSNibName
	// The full path of the nib file that stores the window associated with the receiver.
	WindowNibPath() string

	// Topic: Accessing Window Attributes and Content

	// A Boolean value that indicates whether the window will cascade in relation to other document windows when it is displayed.
	ShouldCascadeWindows() bool
	SetShouldCascadeWindows(value bool)
	// The name under which the frame rectangle of the window owned by the receiver is stored in the defaults database.
	WindowFrameAutosaveName() NSWindowFrameAutosaveName
	SetWindowFrameAutosaveName(value NSWindowFrameAutosaveName)
	// Synchronizes the displayed window title and the represented filename with the information in the associated document.
	SynchronizeWindowTitleWithDocumentName()
	// Returns the window title to be used for a given document display name.
	WindowTitleForDocumentDisplayName(displayName string) string
	// The view controller for the window’s content view.
	ContentViewController() INSViewController
	SetContentViewController(value INSViewController)
	// Dismisses the window controller.
	DismissController(sender objectivec.IObject)

	// Topic: Instance Properties

	PreviewRepresentableActivityItems() []objectivec.IObject
	SetPreviewRepresentableActivityItems(value []objectivec.IObject)

	EncodeWithCoder(coder foundation.INSCoder)
}





// Init initializes the instance.
func (w NSWindowController) Init() NSWindowController {
	rv := objc.Send[NSWindowController](w.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (w NSWindowController) Autorelease() NSWindowController {
	rv := objc.Send[NSWindowController](w.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSWindowController creates a new NSWindowController instance.
func NewNSWindowController() NSWindowController {
	class := getNSWindowControllerClass()
	rv := objc.Send[NSWindowController](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/init(coder:)
func NewWindowControllerWithCoder(coder foundation.INSCoder) NSWindowController {
	instance := getNSWindowControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSWindowControllerFromID(rv)
}


// Returns a window controller initialized with a given window.
//
// window: The window object to manage; can be `nil`.
//
// # Return Value
// 
// A newly initialized window controller.
//
// # Discussion
// 
// This method is the designated initializer for [NSWindowController].
// 
// This initializer is useful when a window has been loaded but no window
// controller is assigned. The default initialization turns on cascading, sets
// the [ShouldCloseDocument] property to [false], and sets the window frame
// autosave name to an empty string. As a side effect, the created window
// controller is added as an observer of the [willCloseNotification]s posted
// by that window object (which is handled by a private method). If you make
// the window controller a delegate of the window, you can implement
// NSWindow’s windowShouldClose: delegate method.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [willCloseNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/willCloseNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/init(window:)
func NewWindowControllerWithWindow(window INSWindow) NSWindowController {
	instance := getNSWindowControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWindow:"), window)
	return NSWindowControllerFromID(rv)
}


// Returns a window controller initialized with a nib file.
//
// windowNibName: The name of the nib file (minus the “`XCUIElementTypeNib`” extension)
// that archives the receiver’s window; cannot be `nil`.
//
// # Discussion
// 
// Sets the owner of the nib file to the receiver. The default initialization
// turns on cascading, sets the [ShouldCloseDocument] property to [false], and
// sets the autosave name for the window’s frame to an empty string.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/init(windowNibName:)
func NewWindowControllerWithWindowNibName(windowNibName NSNibName) NSWindowController {
	instance := getNSWindowControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWindowNibName:"), objc.String(string(windowNibName)))
	return NSWindowControllerFromID(rv)
}


// Returns a window controller initialized with a nib file and a specified
// owner for that nib file.
//
// windowNibName: The name of the nib file (minus the “`XCUIElementTypeNib`” extension)
// that archives the receiver’s window; cannot be `nil`.
//
// owner: The nib file’s owner; cannot be `nil`.
//
// # Discussion
// 
// The default initialization turns on cascading, sets the
// [ShouldCloseDocument] property to [false], and sets the autosave name for
// the window’s frame to an empty string.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/init(windowNibName:owner:)
func NewWindowControllerWithWindowNibNameOwner(windowNibName NSNibName, owner objectivec.IObject) NSWindowController {
	instance := getNSWindowControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWindowNibName:owner:"), objc.String(string(windowNibName)), owner)
	return NSWindowControllerFromID(rv)
}


// Returns a window controller initialized with a nib file at an absolute path
// and a specified owner.
//
// windowNibPath: The full path to the nib file that archives the receiver’s window; cannot
// be `nil`.
//
// owner: The nib file’s owner; cannot be `nil`.
//
// # Discussion
// 
// Use this method if your nib file is at a fixed location (which is not
// inside either the file’s owner’s class’s bundle or in the
// application’s main bundle). The default initialization turns on
// cascading, sets the [ShouldCloseDocument] property to [false], and sets the
// autosave name for the window’s frame to an empty string.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/init(windowNibPath:owner:)
func NewWindowControllerWithWindowNibPathOwner(windowNibPath string, owner objectivec.IObject) NSWindowController {
	instance := getNSWindowControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithWindowNibPath:owner:"), objc.String(windowNibPath), owner)
	return NSWindowControllerFromID(rv)
}







// Returns a window controller initialized with a given window.
//
// window: The window object to manage; can be `nil`.
//
// # Return Value
// 
// A newly initialized window controller.
//
// # Discussion
// 
// This method is the designated initializer for [NSWindowController].
// 
// This initializer is useful when a window has been loaded but no window
// controller is assigned. The default initialization turns on cascading, sets
// the [ShouldCloseDocument] property to [false], and sets the window frame
// autosave name to an empty string. As a side effect, the created window
// controller is added as an observer of the [willCloseNotification]s posted
// by that window object (which is handled by a private method). If you make
// the window controller a delegate of the window, you can implement
// NSWindow’s windowShouldClose: delegate method.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [willCloseNotification]: https://developer.apple.com/documentation/AppKit/NSWindow/willCloseNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/init(window:)
func (w NSWindowController) InitWithWindow(window INSWindow) NSWindowController {
	rv := objc.Send[NSWindowController](w.ID, objc.Sel("initWithWindow:"), window)
	return rv
}

// Returns a window controller initialized with a nib file.
//
// windowNibName: The name of the nib file (minus the “`XCUIElementTypeNib`” extension)
// that archives the receiver’s window; cannot be `nil`.
//
// # Discussion
// 
// Sets the owner of the nib file to the receiver. The default initialization
// turns on cascading, sets the [ShouldCloseDocument] property to [false], and
// sets the autosave name for the window’s frame to an empty string.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/init(windowNibName:)
func (w NSWindowController) InitWithWindowNibName(windowNibName NSNibName) NSWindowController {
	rv := objc.Send[NSWindowController](w.ID, objc.Sel("initWithWindowNibName:"), objc.String(string(windowNibName)))
	return rv
}

// Returns a window controller initialized with a nib file and a specified
// owner for that nib file.
//
// windowNibName: The name of the nib file (minus the “`XCUIElementTypeNib`” extension)
// that archives the receiver’s window; cannot be `nil`.
//
// owner: The nib file’s owner; cannot be `nil`.
//
// # Discussion
// 
// The default initialization turns on cascading, sets the
// [ShouldCloseDocument] property to [false], and sets the autosave name for
// the window’s frame to an empty string.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/init(windowNibName:owner:)
func (w NSWindowController) InitWithWindowNibNameOwner(windowNibName NSNibName, owner objectivec.IObject) NSWindowController {
	rv := objc.Send[NSWindowController](w.ID, objc.Sel("initWithWindowNibName:owner:"), objc.String(string(windowNibName)), owner)
	return rv
}

// Returns a window controller initialized with a nib file at an absolute path
// and a specified owner.
//
// windowNibPath: The full path to the nib file that archives the receiver’s window; cannot
// be `nil`.
//
// owner: The nib file’s owner; cannot be `nil`.
//
// # Discussion
// 
// Use this method if your nib file is at a fixed location (which is not
// inside either the file’s owner’s class’s bundle or in the
// application’s main bundle). The default initialization turns on
// cascading, sets the [ShouldCloseDocument] property to [false], and sets the
// autosave name for the window’s frame to an empty string.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/init(windowNibPath:owner:)
func (w NSWindowController) InitWithWindowNibPathOwner(windowNibPath string, owner objectivec.IObject) NSWindowController {
	rv := objc.Send[NSWindowController](w.ID, objc.Sel("initWithWindowNibPath:owner:"), objc.String(windowNibPath), owner)
	return rv
}

// Loads the receiver’s window from the nib file.
//
// # Discussion
// 
// You should never directly invoke this method. Instead, access the [Window]
// property so the [WindowDidLoad] and [WindowWillLoad] methods are invoked.
// Subclasses can override this method if the way it finds and loads the
// window is not adequate. It uses the [Bundle] class’s [init(for:)] method
// to get the bundle, using the class of the nib file owner as argument. It
// then locates the nib file within the bundle and, if successful, loads it;
// if unsuccessful, it tries to find the nib file in the main bundle.
//
// [Bundle]: https://developer.apple.com/documentation/Foundation/Bundle
// [init(for:)]: https://developer.apple.com/documentation/Foundation/Bundle/init(for:)
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/loadWindow()
func (w NSWindowController) LoadWindow() {
	objc.Send[objc.ID](w.ID, objc.Sel("loadWindow"))
}

// Displays the window associated with the receiver.
//
// sender: The control sending the message; can be `nil`.
//
// # Discussion
// 
// If the window is an [NSPanel] object and has its [BecomesKeyOnlyIfNeeded]
// flag set to [true], the window is displayed in front of all other windows
// but is not made key; otherwise it is displayed in front and is made key.
// This method is useful for menu actions.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/showWindow(_:)
func (w NSWindowController) ShowWindow(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("showWindow:"), sender)
}

// Sent after the window owned by the receiver has been loaded.
//
// # Discussion
// 
// The default implementation does nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/windowDidLoad()
func (w NSWindowController) WindowDidLoad() {
	objc.Send[objc.ID](w.ID, objc.Sel("windowDidLoad"))
}

// Sent before the window owned by the receiver is loaded.
//
// # Discussion
// 
// The default implementation does nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/windowWillLoad()
func (w NSWindowController) WindowWillLoad() {
	objc.Send[objc.ID](w.ID, objc.Sel("windowWillLoad"))
}

// Sets the document edited flag for the window controller.
//
// dirtyFlag: [true] if the document has been edited since its last save, [false] if it
// hasn’t.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The window controller uses this flag to control whether its associated
// window shows up as dirty. You should not call this method directly for
// window controllers with an associated document; the document calls this
// method on its window controllers as needed.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/setDocumentEdited(_:)
func (w NSWindowController) SetDocumentEdited(dirtyFlag bool) {
	objc.Send[objc.ID](w.ID, objc.Sel("setDocumentEdited:"), dirtyFlag)
}

// Closes the window if it was loaded.
//
// # Discussion
// 
// Because this method closes the window without asking the user for
// confirmation, you usually do not invoke it when the Close menu command is
// chosen. Instead invoke NSWindow’s [PerformClose] on the receiver’s
// window.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/close()
func (w NSWindowController) Close() {
	objc.Send[objc.ID](w.ID, objc.Sel("close"))
}

// Synchronizes the displayed window title and the represented filename with
// the information in the associated document.
//
// # Discussion
// 
// Does nothing if the window controller has no associated document or loaded
// window. This method queries the window controller’s document to get the
// document’s display name and full filename path, then calls
// [WindowTitleForDocumentDisplayName] to get the display name to show in the
// window title.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/synchronizeWindowTitleWithDocumentName()
func (w NSWindowController) SynchronizeWindowTitleWithDocumentName() {
	objc.Send[objc.ID](w.ID, objc.Sel("synchronizeWindowTitleWithDocumentName"))
}

// Returns the window title to be used for a given document display name.
//
// displayName: The display name for the document. This is the last path component under
// which the document file is saved.
//
// # Discussion
// 
// The default implementation returns `displayName`. Subclasses can override
// this method to customize the window title. For example, a CAD application
// could append “-Top” or “-Side,” depending on the view displayed by
// the window.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/windowTitle(forDocumentDisplayName:)
func (w NSWindowController) WindowTitleForDocumentDisplayName(displayName string) string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("windowTitleForDocumentDisplayName:"), objc.String(displayName))
	return foundation.NSStringFromID(rv).String()
}

// Dismisses the window controller.
//
// sender: The sender of the message.
//
// # Discussion
// 
// This method does nothing if the receiver is not currently presented.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/dismissController(_:)
func (w NSWindowController) DismissController(sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("dismissController:"), sender)
}

// Performs the specified segue.
//
// identifier: The string that uniquely identifies the segue in the storyboard file.
// 
// In Interface Builder, you can provide an identifier string to a segue using
// the inspector. Pass this string to this parameter.
//
// sender: The object that you want to use to initiate the segue. This parameter makes
// the object available to your implementation during the segue.
//
// # Discussion
// 
// Apps typically do not need to trigger segues programmatically. If needed,
// you can call this method to trigger a segue for an action that cannot be
// expressed in a storyboard file, such as a transition between scenes in
// different storyboards.
// 
// Typically, a segue is triggered by a user action, such as clicking a
// button. In Interface Builder, configure an object, such as a control
// embedded in the view controller’s view hierarchy, to trigger the segue.
//
// See: https://developer.apple.com/documentation/AppKit/NSSeguePerforming/performSegue(withIdentifier:sender:)
func (w NSWindowController) PerformSegueWithIdentifierSender(identifier NSStoryboardSegueIdentifier, sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("performSegueWithIdentifier:sender:"), objc.String(string(identifier)), sender)
}

// Called when a segue is about to be performed.
//
// segue: The segue object containing information about the view controllers involved
// in the segue.
//
// sender: The object that initiated the segue. You might use this parameter to
// perform different actions based on which control (or other object)
// initiated the segue.
//
// # Discussion
// 
// The default implementation of this method does nothing; you can override it
// to pass relevant data to the new view controller or window controller,
// based on the context of the segue. The `segue` object describes the
// transition and includes references to both controllers involved in the
// segue.
// 
// Segues can be triggered from multiple sources, so use the information in
// the `segue` and `sender` parameters to disambiguate between different
// logical paths in your app. For example, if the segue originated from a
// table view, the `sender` parameter would identify the cell that the user
// clicked. You could use that information to set the data on the destination
// view controller.
//
// See: https://developer.apple.com/documentation/AppKit/NSSeguePerforming/prepare(for:sender:)
func (w NSWindowController) PrepareForSegueSender(segue INSStoryboardSegue, sender objectivec.IObject) {
	objc.Send[objc.ID](w.ID, objc.Sel("prepareForSegue:sender:"), segue, sender)
}

// Called immediately prior to the performance of a storyboard segue.
//
// identifier: The string that identifies the segue to be performed.
// 
// Using the Interface Builder inspector, provide a unique identifier string
// for each segue in a storyboard. The system provides a segue’s identifier
// to this parameter when it calls this method. The identifier string is used
// to locate the segue inside the storyboard file that contains the view
// controller.
//
// sender: The object that initiated the segue. This object is made available for
// informational purposes during the segue.
//
// # Return Value
// 
// [true] to allow a segue to proceed or [false] to stop it from proceeding.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Override this method to return [false] for cases where you want to prevent
// the performance of a segue. By default, invocation of a segue results in
// the segue being performed.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSSeguePerforming/shouldPerformSegue(withIdentifier:sender:)
func (w NSWindowController) ShouldPerformSegueWithIdentifierSender(identifier NSStoryboardSegueIdentifier, sender objectivec.IObject) bool {
	rv := objc.Send[bool](w.ID, objc.Sel("shouldPerformSegueWithIdentifier:sender:"), objc.String(string(identifier)), sender)
	return rv
}
func (w NSWindowController) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](w.ID, objc.Sel("encodeWithCoder:"), coder)
}












// A Boolean value that indicates whether the nib file containing the
// receiver’s window has been loaded.
//
// # Discussion
// 
// The value of this property is [true] if the nib file containing the
// receiver’s window has been loaded, [false] otherwise.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/isWindowLoaded
func (w NSWindowController) WindowLoaded() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("isWindowLoaded"))
	return rv
}



// The window owned by the receiver.
//
// # Discussion
// 
// Accessing this property loads the nib file if there is one and it has not
// yet been loaded. If the window was loaded, the following methods are called
// in order: [WindowWillLoad], [LoadWindow], and [WindowDidLoad]. If the
// window controller has a document, the document’s corresponding methods
// [WindowControllerWillLoadNib] and [WindowControllerDidLoadNib] are also
// called (if implemented). To affect nib loading or do something before or
// after it happens, you should always override these methods.
// 
// Setting this property releases the window controller’s old window along
// with any associated top-level objects in its nib file, and establishes
// ownership of the specified new window. Typically, you should not use this
// property to set the window. Instead, create a new window controller for the
// new window and then release the old window controller.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/window
func (w NSWindowController) Window() INSWindow {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("window"))
	return NSWindowFromID(objc.ID(rv))
}
func (w NSWindowController) SetWindow(value INSWindow) {
	objc.Send[struct{}](w.ID, objc.Sel("setWindow:"), value)
}



// The document associated with the window controller.
//
// # Discussion
// 
// The value of this property is `nil` if no document is associated with the
// window. When a window controller is added to a document’s list of window
// controllers, the document uses this property to set the window
// controller’s document; you should not set this property. AppKit uses this
// property to access the document for relevant next-responder messages.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/document
func (w NSWindowController) Document() objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("document"))
	return objectivec.Object{ID: rv}
}
func (w NSWindowController) SetDocument(value objectivec.IObject) {
	objc.Send[struct{}](w.ID, objc.Sel("setDocument:"), value)
}



// A Boolean value that indicates whether the receiver necessarily closes the
// associated document when the window it manages is closed.
//
// # Discussion
// 
// The value of this property is [true] if the receiver necessarily closes the
// associated document when the window it manages is closed, [false]
// otherwise. The default value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/shouldCloseDocument
func (w NSWindowController) ShouldCloseDocument() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("shouldCloseDocument"))
	return rv
}
func (w NSWindowController) SetShouldCloseDocument(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setShouldCloseDocument:"), value)
}



// The owner of the nib file containing the window managed by the receiver.
//
// # Discussion
// 
// The owner of the nib file containing the window managed by the receiver is
// usually `self`, but it can be the receiver’s document or some other
// object.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/owner
func (w NSWindowController) Owner() objectivec.IObject {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("owner"))
	return objectivec.Object{ID: rv}
}



// The storyboard file from which the window controller was loaded.
//
// # Discussion
// 
// The value of this property is `nil` if the window controller was not loaded
// from a storyboard.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/storyboard
func (w NSWindowController) Storyboard() INSStoryboard {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("storyboard"))
	return NSStoryboardFromID(objc.ID(rv))
}



// The name of the nib file that stores the window associated with the
// receiver.
//
// # Discussion
// 
// If [InitWithWindowNibPathOwner] was used to initialize the instance,
// [WindowNibName] contains the last path component with the
// “`XCUIElementTypeNib`” extension stripped off. If
// [InitWithWindowNibName] or [InitWithWindowNibNameOwner] was used,
// [NSWindowController] contains the name without the
// “`XCUIElementTypeNib`” extension.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/windowNibName
func (w NSWindowController) WindowNibName() NSNibName {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("windowNibName"))
	return NSNibName(foundation.NSStringFromID(rv).String())
}



// The full path of the nib file that stores the window associated with the
// receiver.
//
// # Discussion
// 
// If [InitWithWindowNibPathOwner] was used to initialize the instance, this
// property contains the path. If [InitWithWindowNibName] or
// [InitWithWindowNibNameOwner] was used, [WindowNibPath] locates the nib in
// the file’s owner’s class’ bundle or in the application’s main
// bundle and returns the full path (or `nil` if it cannot be located).
// Subclasses can override this behavior to augment the search behavior, but
// probably ought to call `super` first.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/windowNibPath
func (w NSWindowController) WindowNibPath() string {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("windowNibPath"))
	return foundation.NSStringFromID(rv).String()
}



// A Boolean value that indicates whether the window will cascade in relation
// to other document windows when it is displayed.
//
// # Discussion
// 
// Cascading in relation to other document windows means having a slightly
// offset location so that the title bars of previously displayed windows are
// still visible. The value of this property is [true] if the window will
// cascade in relation to other document windows, [false] otherwise. The
// default value is [true].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/shouldCascadeWindows
func (w NSWindowController) ShouldCascadeWindows() bool {
	rv := objc.Send[bool](w.ID, objc.Sel("shouldCascadeWindows"))
	return rv
}
func (w NSWindowController) SetShouldCascadeWindows(value bool) {
	objc.Send[struct{}](w.ID, objc.Sel("setShouldCascadeWindows:"), value)
}



// The name under which the frame rectangle of the window owned by the
// receiver is stored in the defaults database.
//
// # Discussion
// 
// By default, `name` is an empty string, which means that no information is
// stored in the defaults database.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/windowFrameAutosaveName
func (w NSWindowController) WindowFrameAutosaveName() NSWindowFrameAutosaveName {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("windowFrameAutosaveName"))
	return NSWindowFrameAutosaveName(foundation.NSStringFromID(rv).String())
}
func (w NSWindowController) SetWindowFrameAutosaveName(value NSWindowFrameAutosaveName) {
	objc.Send[struct{}](w.ID, objc.Sel("setWindowFrameAutosaveName:"), objc.String(string(value)))
}



// The view controller for the window’s content view.
//
// # Discussion
// 
// The value of this property tracks the window’s [ContentView] property.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowController/contentViewController
func (w NSWindowController) ContentViewController() INSViewController {
	rv := objc.Send[objc.ID](w.ID, objc.Sel("contentViewController"))
	return NSViewControllerFromID(objc.ID(rv))
}
func (w NSWindowController) SetContentViewController(value INSViewController) {
	objc.Send[struct{}](w.ID, objc.Sel("setContentViewController:"), value)
}



// See: https://developer.apple.com/documentation/AppKit/NSWindowController/previewRepresentableActivityItems
func (w NSWindowController) PreviewRepresentableActivityItems() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](w.ID, objc.Sel("previewRepresentableActivityItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (w NSWindowController) SetPreviewRepresentableActivityItems(value []objectivec.IObject) {
	objc.Send[struct{}](w.ID, objc.Sel("setPreviewRepresentableActivityItems:"), objectivec.IObjectSliceToNSArray(value))
}















			// Protocol methods for NSSeguePerforming
			


















