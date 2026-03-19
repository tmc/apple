// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSDocumentController] class.
var (
	_NSDocumentControllerClass     NSDocumentControllerClass
	_NSDocumentControllerClassOnce sync.Once
)

func getNSDocumentControllerClass() NSDocumentControllerClass {
	_NSDocumentControllerClassOnce.Do(func() {
		_NSDocumentControllerClass = NSDocumentControllerClass{class: objc.GetClass("NSDocumentController")}
	})
	return _NSDocumentControllerClass
}

// GetNSDocumentControllerClass returns the class object for NSDocumentController.
func GetNSDocumentControllerClass() NSDocumentControllerClass {
	return getNSDocumentControllerClass()
}

type NSDocumentControllerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDocumentControllerClass) Alloc() NSDocumentController {
	rv := objc.Send[NSDocumentController](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that manages an app’s documents.
//
// # Overview
// 
// As the first-responder target of New and Open menu commands,
// [NSDocumentController] creates and opens documents and tracks them
// throughout a session of the app. When opening documents, a document
// controller runs and manages the modal Open panel. [NSDocumentController]
// objects also maintain and manage the mappings of document types,
// extensions, and [NSDocument] subclasses as specified in the
// [CFBundleDocumentTypes] property loaded from the information property list
// (`Info.Plist()`).
// 
// You can use various [NSDocumentController] methods to get a list of the
// current documents, get the current document (which is the document whose
// window is currently key), get documents based on a given filename or
// window, and find out about a document’s extension, type, display name,
// and document class.
// 
// In some situations, it’s worthwhile to subclass [NSDocumentController] in
// non-[NSDocument]-based apps to get some of its features. For example, the
// [NSDocumentController] management of the Open Recent menu is useful in apps
// that don’t use subclasses of [NSDocument].
//
// [CFBundleDocumentTypes]: https://developer.apple.com/library/archive/documentation/General/Reference/InfoPlistKeyReference/Articles/CoreFoundationKeys.html#//apple_ref/doc/uid/20001431-101685
//
// # Initializing a New NSDocumentController
//
//   - [NSDocumentController.InitWithCoder]: This method initializes a new NSDocumentController from the coder.
//
// # Creating and Opening Documents
//
//   - [NSDocumentController.DocumentForURL]: Returns, for a given URL, the open document whose file or file package is located by the URL, or `nil` if there is no such open document.
//   - [NSDocumentController.DuplicateDocumentWithContentsOfURLCopyingDisplayNameError]: Creates a new document by reading the contents for the document from another URL, presents its user interface, and returns the document if successful.
//   - [NSDocumentController.OpenDocumentWithContentsOfURLDisplayCompletionHandler]: Opens a document located by a URL, optionally presents its user interface, and calls the passed-in completion handler.
//   - [NSDocumentController.OpenUntitledDocumentAndDisplayError]: Creates a new untitled document, presents its user interface if `displayDocument` is `true`, and returns the document if successful.
//   - [NSDocumentController.MakeDocumentForURLWithContentsOfURLOfTypeError]: Instantiates a document located by a URL, of a specified type, but by reading the contents for the document from another URL, and returns it if successful.
//   - [NSDocumentController.MakeDocumentWithContentsOfURLOfTypeError]: Instantiates a document located by a URL, of a specified type, and returns it if successful.
//   - [NSDocumentController.MakeUntitledDocumentOfTypeError]: Instantiates a new untitled document of the specified type and returns it if successful.
//   - [NSDocumentController.ReopenDocumentForURLWithContentsOfURLDisplayCompletionHandler]: Reopens a document, optionally located by a URL, by reading the contents for the document from another URL, optionally presents its user interface, and calls the passed-in completion handler.
//
// # Managing Documents
//
//   - [NSDocumentController.Documents]: The document objects managed by the receiver.
//   - [NSDocumentController.AddDocument]: Adds the given document to the list of open documents.
//   - [NSDocumentController.CurrentDocument]: The document object associated with the main window.
//   - [NSDocumentController.DocumentForWindow]: Returns the document object whose window controller owns a specified window.
//   - [NSDocumentController.HasEditedDocuments]: A Boolean value indicating whether the receiver has any documents with unsaved changes.
//   - [NSDocumentController.RemoveDocument]: Removes the given document from the list of open documents.
//
// # Managing Document Types
//
//   - [NSDocumentController.DocumentClassNames]: An array of strings representing the custom document classes supported by this app.
//   - [NSDocumentController.DefaultType]: Returns the name of the document type that should be used when creating new documents.
//   - [NSDocumentController.DocumentClassForType]: Returns the [NSDocument] subclass associated with a given document type.
//   - [NSDocumentController.DisplayNameForType]: Returns the descriptive name for the specified document type, which is used in the File Format pop-up menu of the Save As dialog.
//   - [NSDocumentController.TypeForContentsOfURLError]: Returns, for a specified URL, the document type identifier to use when opening the document at that location, if successful.
//
// # Autosaving
//
//   - [NSDocumentController.AutosavingDelay]: The time interval (in seconds) for periodic autosaving.
//   - [NSDocumentController.SetAutosavingDelay]
//
// # Closing Documents
//
//   - [NSDocumentController.CloseAllDocumentsWithDelegateDidCloseAllSelectorContextInfo]: Iterates through all the open documents and tries to close them one by one using the specified delegate.
//   - [NSDocumentController.ReviewUnsavedDocumentsWithAlertTitleCancellableDelegateDidReviewAllSelectorContextInfo]: Displays an alert asking if the user wants to review unsaved documents, quit regardless of unsaved documents, or cancel the save operation.
//
// # Responding to Action Messages
//
//   - [NSDocumentController.NewDocument]: An action method called by the New menu command, this method creates a new [NSDocument] object and adds it to the list of such objects managed by the document controller.
//   - [NSDocumentController.OpenDocument]: An action method called by the Open menu command, it runs the modal Open panel and, based on the selected filenames, creates one or more [NSDocument] objects from the contents of the files.
//   - [NSDocumentController.SaveAllDocuments]: As the action method called by the Save All command, saves all open documents of the application that need to be saved.
//
// # Managing the Open Dialog
//
//   - [NSDocumentController.BeginOpenPanelForTypesCompletionHandler]: Presents a nonmodal Open dialog that displays files you can open from a list of UTIs.
//   - [NSDocumentController.RunModalOpenPanelForTypes]: Presents a modal Open dialog and limits selection to specific file types.
//   - [NSDocumentController.CurrentDirectory]: The directory path to use as the starting point in the Open dialog.
//   - [NSDocumentController.URLsFromRunningOpenPanel]: An array of URLs that correspond to the selected files in a running Open dialog.
//
// # Managing the Open Recent Menu
//
//   - [NSDocumentController.MaximumRecentDocumentCount]: The maximum number of items that may be presented in the standard Open Recent menu.
//   - [NSDocumentController.ClearRecentDocuments]: Empties the recent documents list for the application.
//   - [NSDocumentController.NoteNewRecentDocumentURL]: Adds or replaces an Open Recent menu item corresponding to the data located by the URL.
//   - [NSDocumentController.NoteNewRecentDocument]: Adds or replaces an Open Recent menu item corresponding to the document.
//   - [NSDocumentController.RecentDocumentURLs]: The list of recent-document URLs.
//
// # Sharing
//
//   - [NSDocumentController.AllowsAutomaticShareMenu]: A Boolean value that the system uses to insert a Share menu in the File menu.
//   - [NSDocumentController.StandardShareMenuItem]: Returns a menu item that your app uses for sharing the current document.
//
// # Handling Errors
//
//   - [NSDocumentController.PresentError]: Presents an error alert to the user as a modal panel.
//   - [NSDocumentController.PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo]: Presents an error alert to the user as a modal panel.
//   - [NSDocumentController.WillPresentError]: Indicates an error condition and provides the opportunity to return the same or a different error.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController
type NSDocumentController struct {
	objectivec.Object
}

// NSDocumentControllerFromID constructs a [NSDocumentController] from an objc.ID.
//
// An object that manages an app’s documents.
func NSDocumentControllerFromID(id objc.ID) NSDocumentController {
	return NSDocumentController{objectivec.Object{ID: id}}
}
// NOTE: NSDocumentController adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSDocumentController] class.
//
// # Initializing a New NSDocumentController
//
//   - [INSDocumentController.InitWithCoder]: This method initializes a new NSDocumentController from the coder.
//
// # Creating and Opening Documents
//
//   - [INSDocumentController.DocumentForURL]: Returns, for a given URL, the open document whose file or file package is located by the URL, or `nil` if there is no such open document.
//   - [INSDocumentController.DuplicateDocumentWithContentsOfURLCopyingDisplayNameError]: Creates a new document by reading the contents for the document from another URL, presents its user interface, and returns the document if successful.
//   - [INSDocumentController.OpenDocumentWithContentsOfURLDisplayCompletionHandler]: Opens a document located by a URL, optionally presents its user interface, and calls the passed-in completion handler.
//   - [INSDocumentController.OpenUntitledDocumentAndDisplayError]: Creates a new untitled document, presents its user interface if `displayDocument` is `true`, and returns the document if successful.
//   - [INSDocumentController.MakeDocumentForURLWithContentsOfURLOfTypeError]: Instantiates a document located by a URL, of a specified type, but by reading the contents for the document from another URL, and returns it if successful.
//   - [INSDocumentController.MakeDocumentWithContentsOfURLOfTypeError]: Instantiates a document located by a URL, of a specified type, and returns it if successful.
//   - [INSDocumentController.MakeUntitledDocumentOfTypeError]: Instantiates a new untitled document of the specified type and returns it if successful.
//   - [INSDocumentController.ReopenDocumentForURLWithContentsOfURLDisplayCompletionHandler]: Reopens a document, optionally located by a URL, by reading the contents for the document from another URL, optionally presents its user interface, and calls the passed-in completion handler.
//
// # Managing Documents
//
//   - [INSDocumentController.Documents]: The document objects managed by the receiver.
//   - [INSDocumentController.AddDocument]: Adds the given document to the list of open documents.
//   - [INSDocumentController.CurrentDocument]: The document object associated with the main window.
//   - [INSDocumentController.DocumentForWindow]: Returns the document object whose window controller owns a specified window.
//   - [INSDocumentController.HasEditedDocuments]: A Boolean value indicating whether the receiver has any documents with unsaved changes.
//   - [INSDocumentController.RemoveDocument]: Removes the given document from the list of open documents.
//
// # Managing Document Types
//
//   - [INSDocumentController.DocumentClassNames]: An array of strings representing the custom document classes supported by this app.
//   - [INSDocumentController.DefaultType]: Returns the name of the document type that should be used when creating new documents.
//   - [INSDocumentController.DocumentClassForType]: Returns the [NSDocument] subclass associated with a given document type.
//   - [INSDocumentController.DisplayNameForType]: Returns the descriptive name for the specified document type, which is used in the File Format pop-up menu of the Save As dialog.
//   - [INSDocumentController.TypeForContentsOfURLError]: Returns, for a specified URL, the document type identifier to use when opening the document at that location, if successful.
//
// # Autosaving
//
//   - [INSDocumentController.AutosavingDelay]: The time interval (in seconds) for periodic autosaving.
//   - [INSDocumentController.SetAutosavingDelay]
//
// # Closing Documents
//
//   - [INSDocumentController.CloseAllDocumentsWithDelegateDidCloseAllSelectorContextInfo]: Iterates through all the open documents and tries to close them one by one using the specified delegate.
//   - [INSDocumentController.ReviewUnsavedDocumentsWithAlertTitleCancellableDelegateDidReviewAllSelectorContextInfo]: Displays an alert asking if the user wants to review unsaved documents, quit regardless of unsaved documents, or cancel the save operation.
//
// # Responding to Action Messages
//
//   - [INSDocumentController.NewDocument]: An action method called by the New menu command, this method creates a new [NSDocument] object and adds it to the list of such objects managed by the document controller.
//   - [INSDocumentController.OpenDocument]: An action method called by the Open menu command, it runs the modal Open panel and, based on the selected filenames, creates one or more [NSDocument] objects from the contents of the files.
//   - [INSDocumentController.SaveAllDocuments]: As the action method called by the Save All command, saves all open documents of the application that need to be saved.
//
// # Managing the Open Dialog
//
//   - [INSDocumentController.BeginOpenPanelForTypesCompletionHandler]: Presents a nonmodal Open dialog that displays files you can open from a list of UTIs.
//   - [INSDocumentController.RunModalOpenPanelForTypes]: Presents a modal Open dialog and limits selection to specific file types.
//   - [INSDocumentController.CurrentDirectory]: The directory path to use as the starting point in the Open dialog.
//   - [INSDocumentController.URLsFromRunningOpenPanel]: An array of URLs that correspond to the selected files in a running Open dialog.
//
// # Managing the Open Recent Menu
//
//   - [INSDocumentController.MaximumRecentDocumentCount]: The maximum number of items that may be presented in the standard Open Recent menu.
//   - [INSDocumentController.ClearRecentDocuments]: Empties the recent documents list for the application.
//   - [INSDocumentController.NoteNewRecentDocumentURL]: Adds or replaces an Open Recent menu item corresponding to the data located by the URL.
//   - [INSDocumentController.NoteNewRecentDocument]: Adds or replaces an Open Recent menu item corresponding to the document.
//   - [INSDocumentController.RecentDocumentURLs]: The list of recent-document URLs.
//
// # Sharing
//
//   - [INSDocumentController.AllowsAutomaticShareMenu]: A Boolean value that the system uses to insert a Share menu in the File menu.
//   - [INSDocumentController.StandardShareMenuItem]: Returns a menu item that your app uses for sharing the current document.
//
// # Handling Errors
//
//   - [INSDocumentController.PresentError]: Presents an error alert to the user as a modal panel.
//   - [INSDocumentController.PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo]: Presents an error alert to the user as a modal panel.
//   - [INSDocumentController.WillPresentError]: Indicates an error condition and provides the opportunity to return the same or a different error.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController
type INSDocumentController interface {
	objectivec.IObject
	NSMenuItemValidation
	NSUserInterfaceValidations
	NSWindowRestoration

	// Topic: Initializing a New NSDocumentController

	// This method initializes a new NSDocumentController from the coder.
	InitWithCoder(coder foundation.INSCoder) NSDocumentController

	// Topic: Creating and Opening Documents

	// Returns, for a given URL, the open document whose file or file package is located by the URL, or `nil` if there is no such open document.
	DocumentForURL(url foundation.INSURL) INSDocument
	// Creates a new document by reading the contents for the document from another URL, presents its user interface, and returns the document if successful.
	DuplicateDocumentWithContentsOfURLCopyingDisplayNameError(url foundation.INSURL, duplicateByCopying bool, displayNameOrNil string) (INSDocument, error)
	// Opens a document located by a URL, optionally presents its user interface, and calls the passed-in completion handler.
	OpenDocumentWithContentsOfURLDisplayCompletionHandler(url foundation.INSURL, displayDocument bool, completionHandler DocumentErrorHandler)
	// Creates a new untitled document, presents its user interface if `displayDocument` is `true`, and returns the document if successful.
	OpenUntitledDocumentAndDisplayError(displayDocument bool) (INSDocument, error)
	// Instantiates a document located by a URL, of a specified type, but by reading the contents for the document from another URL, and returns it if successful.
	MakeDocumentForURLWithContentsOfURLOfTypeError(urlOrNil foundation.INSURL, contentsURL foundation.INSURL, typeName string) (INSDocument, error)
	// Instantiates a document located by a URL, of a specified type, and returns it if successful.
	MakeDocumentWithContentsOfURLOfTypeError(url foundation.INSURL, typeName string) (INSDocument, error)
	// Instantiates a new untitled document of the specified type and returns it if successful.
	MakeUntitledDocumentOfTypeError(typeName string) (INSDocument, error)
	// Reopens a document, optionally located by a URL, by reading the contents for the document from another URL, optionally presents its user interface, and calls the passed-in completion handler.
	ReopenDocumentForURLWithContentsOfURLDisplayCompletionHandler(urlOrNil foundation.INSURL, contentsURL foundation.INSURL, displayDocument bool, completionHandler DocumentErrorHandler)

	// Topic: Managing Documents

	// The document objects managed by the receiver.
	Documents() []NSDocument
	// Adds the given document to the list of open documents.
	AddDocument(document INSDocument)
	// The document object associated with the main window.
	CurrentDocument() INSDocument
	// Returns the document object whose window controller owns a specified window.
	DocumentForWindow(window INSWindow) INSDocument
	// A Boolean value indicating whether the receiver has any documents with unsaved changes.
	HasEditedDocuments() bool
	// Removes the given document from the list of open documents.
	RemoveDocument(document INSDocument)

	// Topic: Managing Document Types

	// An array of strings representing the custom document classes supported by this app.
	DocumentClassNames() []string
	// Returns the name of the document type that should be used when creating new documents.
	DefaultType() string
	// Returns the [NSDocument] subclass associated with a given document type.
	DocumentClassForType(typeName string) objc.Class
	// Returns the descriptive name for the specified document type, which is used in the File Format pop-up menu of the Save As dialog.
	DisplayNameForType(typeName string) string
	// Returns, for a specified URL, the document type identifier to use when opening the document at that location, if successful.
	TypeForContentsOfURLError(url foundation.INSURL) (string, error)

	// Topic: Autosaving

	// The time interval (in seconds) for periodic autosaving.
	AutosavingDelay() float64
	SetAutosavingDelay(value float64)

	// Topic: Closing Documents

	// Iterates through all the open documents and tries to close them one by one using the specified delegate.
	CloseAllDocumentsWithDelegateDidCloseAllSelectorContextInfo(delegate objectivec.IObject, didCloseAllSelector objc.SEL, contextInfo uintptr)
	// Displays an alert asking if the user wants to review unsaved documents, quit regardless of unsaved documents, or cancel the save operation.
	ReviewUnsavedDocumentsWithAlertTitleCancellableDelegateDidReviewAllSelectorContextInfo(title string, cancellable bool, delegate objectivec.IObject, didReviewAllSelector objc.SEL, contextInfo uintptr)

	// Topic: Responding to Action Messages

	// An action method called by the New menu command, this method creates a new [NSDocument] object and adds it to the list of such objects managed by the document controller.
	NewDocument(sender objectivec.IObject)
	// An action method called by the Open menu command, it runs the modal Open panel and, based on the selected filenames, creates one or more [NSDocument] objects from the contents of the files.
	OpenDocument(sender objectivec.IObject)
	// As the action method called by the Save All command, saves all open documents of the application that need to be saved.
	SaveAllDocuments(sender objectivec.IObject)

	// Topic: Managing the Open Dialog

	// Presents a nonmodal Open dialog that displays files you can open from a list of UTIs.
	BeginOpenPanelForTypesCompletionHandler(openPanel INSOpenPanel, inTypes []string, completionHandler IntHandler)
	// Presents a modal Open dialog and limits selection to specific file types.
	RunModalOpenPanelForTypes(openPanel INSOpenPanel, types []string) int
	// The directory path to use as the starting point in the Open dialog.
	CurrentDirectory() string
	// An array of URLs that correspond to the selected files in a running Open dialog.
	URLsFromRunningOpenPanel() []foundation.NSURL

	// Topic: Managing the Open Recent Menu

	// The maximum number of items that may be presented in the standard Open Recent menu.
	MaximumRecentDocumentCount() uint
	// Empties the recent documents list for the application.
	ClearRecentDocuments(sender objectivec.IObject)
	// Adds or replaces an Open Recent menu item corresponding to the data located by the URL.
	NoteNewRecentDocumentURL(url foundation.INSURL)
	// Adds or replaces an Open Recent menu item corresponding to the document.
	NoteNewRecentDocument(document INSDocument)
	// The list of recent-document URLs.
	RecentDocumentURLs() []foundation.NSURL

	// Topic: Sharing

	// A Boolean value that the system uses to insert a Share menu in the File menu.
	AllowsAutomaticShareMenu() bool
	// Returns a menu item that your app uses for sharing the current document.
	StandardShareMenuItem() INSMenuItem

	// Topic: Handling Errors

	// Presents an error alert to the user as a modal panel.
	PresentError(error_ foundation.INSError) bool
	// Presents an error alert to the user as a modal panel.
	PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error_ foundation.INSError, window INSWindow, delegate objectivec.IObject, didPresentSelector objc.SEL, contextInfo uintptr)
	// Indicates an error condition and provides the opportunity to return the same or a different error.
	WillPresentError(error_ foundation.INSError) foundation.INSError

	EncodeWithCoder(coder foundation.INSCoder)
}

// Init initializes the instance.
func (d NSDocumentController) Init() NSDocumentController {
	rv := objc.Send[NSDocumentController](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDocumentController) Autorelease() NSDocumentController {
	rv := objc.Send[NSDocumentController](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDocumentController creates a new NSDocumentController instance.
func NewNSDocumentController() NSDocumentController {
	class := getNSDocumentControllerClass()
	rv := objc.Send[NSDocumentController](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// This method initializes a new NSDocumentController from the coder.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/init(coder:)
func NewDocumentControllerWithCoder(coder foundation.INSCoder) NSDocumentController {
	instance := getNSDocumentControllerClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSDocumentControllerFromID(rv)
}

// This method initializes a new NSDocumentController from the coder.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/init(coder:)
func (d NSDocumentController) InitWithCoder(coder foundation.INSCoder) NSDocumentController {
	rv := objc.Send[NSDocumentController](d.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}
// Returns, for a given URL, the open document whose file or file package is
// located by the URL, or `nil` if there is no such open document.
//
// url: The URL of the document you wish to look up.
//
// # Return Value
// 
// The newly created [NSDocument] object, or `nil` if the document could not
// be created.
//
// # Discussion
// 
// The default implementation of this method queries each open document to
// find one whose URL matches, and returns the first one whose URL does match.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/document(for:)-i5zi
func (d NSDocumentController) DocumentForURL(url foundation.INSURL) INSDocument {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("documentForURL:"), url)
	return NSDocumentFromID(rv)
}
// Creates a new document by reading the contents for the document from
// another URL, presents its user interface, and returns the document if
// successful.
//
// url: The URL locating the document from which contents of the new document are
// copied.
//
// duplicateByCopying: If [true], the contents located at the passed-in URL are copied into a file
// located in the directory used for the autosaved contents of untitled
// documents.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// displayNameOrNil: If not `nil` then this value is used to derive a display name for the new
// document that does not match one that is already in use by an open
// document.
//
// # Return Value
// 
// The newly created [NSDocument] object, or `nil` if the document could not
// be created.
//
// # Discussion
// 
// The default implementation of this method copies the file if specified,
// determines the type of the document, calls
// [DocumentForURLWithContentsOfURLOfTypeError] to instantiate it, sends the
// document `` to name it if `displayNameOrNil` is not `nil`, calls
// [AddDocument] to record its opening, and sends the document
// [WindowControllers] and [ShowWindows] messages.
// 
// The default implementation of this method uses the file coordination
// mechanism introduced in OS X v10.7. It passes the document to the
// [NSFileCoordinator] method [addFilePresenter(_:)] immediately after calling
// the [AddDocument] method. (The balancing invocation of the
// [NSFileCoordinator] method [removeFilePresenter(_:)] is in the [NSDocument]
// method [Close].)
// 
// You can override this method to customize how documents are duplicated. It
// is called by the [NSDocument] method [DuplicateAndReturnError]. It may also
// be called from other places in AppKit.
// 
// In most cases, an app does not need to call this method directly.
//
// [NSFileCoordinator]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator
// [addFilePresenter(_:)]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/addFilePresenter(_:)
// [removeFilePresenter(_:)]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/removeFilePresenter(_:)
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/duplicateDocument(withContentsOf:copying:displayName:)
func (d NSDocumentController) DuplicateDocumentWithContentsOfURLCopyingDisplayNameError(url foundation.INSURL, duplicateByCopying bool, displayNameOrNil string) (INSDocument, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("duplicateDocumentWithContentsOfURL:copying:displayName:error:"), url, duplicateByCopying, objc.String(displayNameOrNil), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSDocument{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSDocumentFromID(rv), nil

}
// Opens a document located by a URL, optionally presents its user interface,
// and calls the passed-in completion handler.
//
// url: The URL locating the document to open.
//
// displayDocument: If [true], displays the document’s user interface.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// completionHandler: The completion handler block object passed in to be called at some point in
// the future, perhaps after the method invocation has returned. The
// completion handler must be called on the main thread.
// 
// The block takes three arguments:
// 
// `document`: The document that was opened, if successful. Otherwise, `nil`.
// `documentWasAlreadyOpen`: Whether the document was already open or being
// opened when this method was called. `error`: If not successful, an
// [NSError] object that encapsulates the reason why the document could not be
// opened.
//
// # Discussion
// 
// The default implementation of this method checks to see if the document is
// already open or being opened, and if it is not determines the type of the
// document, calls [DocumentWithContentsOfURLOfTypeError] to instantiate it,
// and calls [AddDocument] to record its opening. If `displayDocument` is
// [true] and the document is not already open, the default implementation
// calls [WindowControllers] and [ShowWindows]. If the document is already
// open, the implementation just calls [ShowWindows] if `displayDocument` is
// [true]. If the relevant document class returns [true] when sent
// [CanConcurrentlyReadDocumentsOfType] then the invocation of
// [DocumentWithContentsOfURLOfTypeError] is done on a thread other than the
// main one, and when that has returned, the rest of the operation is done on
// the main thread.
// 
// The default implementation of this method uses the file coordination
// mechanism that was added to the Foundation framework in OS X v10.7. All of
// the work it does is one big coordinated read, and it passes the document to
// the [NSFileCoordinator] method [addFilePresenter(_:)] right after calling
// [AddDocument]. (The balancing invocation of the [NSFileCoordinator] method
// [removeFilePresenter(_:)] is in the [NSDocument] method [Close].)
// 
// You can override this method to customize how documents are opened. Its
// implementation, however, is somewhat complex, so you should generally
// investigate overriding one of the methods that it calls instead. However,
// you can override this method to do additional work before calling the
// underlying method on `super`. You can also call the underlying method on
// `super` with a custom completion handler that performs additional work
// before calling the original completion handler. If you do override this
// method you should investigate whether you should also override
// [ReopenDocumentForURLWithContentsOfURLDisplayCompletionHandler] to apply
// the same customization. In either case, take care to always call the
// completion handler on the main thread.
// 
// You can call this method to open a document.
// 
// # Special Considerations
// 
// For backward binary compatibility with OS X v10.6 and earlier, the default
// implementation of this method calls `[self url displayDocument &anError]`
// if that method or the even older [openDocumentWithContentsOfFile:display:]
// method is overridden and this one is not, instead of calling
// [DocumentWithContentsOfURLOfTypeError] and all the rest.
//
// [addFilePresenter(_:)]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/addFilePresenter(_:)
// [openDocumentWithContentsOfFile:display:]: https://developer.apple.com/documentation/AppKit/NSDocumentController/openDocumentWithContentsOfFile:display:
// [removeFilePresenter(_:)]: https://developer.apple.com/documentation/Foundation/NSFileCoordinator/removeFilePresenter(_:)
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/openDocument(withContentsOf:display:completionHandler:)
func (d NSDocumentController) OpenDocumentWithContentsOfURLDisplayCompletionHandler(url foundation.INSURL, displayDocument bool, completionHandler DocumentErrorHandler) {
_block2, _cleanup2 := NewDocumentErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](d.ID, objc.Sel("openDocumentWithContentsOfURL:display:completionHandler:"), url, displayDocument, _block2)
}
// Creates a new untitled document, presents its user interface if
// `displayDocument` is `true`, and returns the document if successful.
//
// displayDocument: [true] if the user interface for the document should be shown, otherwise
// [false].
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Return Value
// 
// Returns the new [NSDocument] object, or `nil` if a new untitled document
// could not be created. If this method returns `nil`, it also sets the
// address referenced by `outError` to an [NSError] object that tell why the
// document could not be created.
//
// # Discussion
// 
// The default implementation of this method calls [DefaultType] to determine
// the type of new document to create, calls [UntitledDocumentOfTypeError] to
// create it, then calls [AddDocument] to record its opening.
// 
// When `displayDocument` is [true], this method sends the new document
// [WindowControllers] and [ShowWindows] messages. In this scenario,
// [ShowWindows] shows only the window controllers that have been assigned to
// the document.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/openUntitledDocumentAndDisplay(_:)
func (d NSDocumentController) OpenUntitledDocumentAndDisplayError(displayDocument bool) (INSDocument, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("openUntitledDocumentAndDisplay:error:"), displayDocument, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSDocument{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSDocumentFromID(rv), nil

}
// Instantiates a document located by a URL, of a specified type, but by
// reading the contents for the document from another URL, and returns it if
// successful.
//
// urlOrNil: The location of the new document object.
//
// contentsURL: The URL of the file that provides the contents for the new document.
//
// typeName: The type of the document.
//
// # Return Value
// 
// The newly created [NSDocument] object, or `nil` if the document could not
// be created.
//
// # Discussion
// 
// The URL is specified by `absoluteDocumentURL`, the type by `typeName`, and
// the other URL providing the contents by `absoluteDocumentContentsURL`. If
// not successful, the method returns `nil` after setting `outError` to point
// to an [NSError] object that encapsulates the reason why the document could
// not be instantiated. The default implementation of this method calls
// [DocumentClassForType] to find out the class of document to instantiate,
// allocates a document object, and initializes it by sending it an
// [InitForURLWithContentsOfURLOfTypeError] message.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/makeDocument(for:withContentsOf:ofType:)
func (d NSDocumentController) MakeDocumentForURLWithContentsOfURLOfTypeError(urlOrNil foundation.INSURL, contentsURL foundation.INSURL, typeName string) (INSDocument, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("makeDocumentForURL:withContentsOfURL:ofType:error:"), urlOrNil, contentsURL, objc.String(typeName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSDocument{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSDocumentFromID(rv), nil

}
// Instantiates a document located by a URL, of a specified type, and returns
// it if successful.
//
// url: The location of the new document object.
//
// typeName: The type of the document.
//
// # Return Value
// 
// The newly created [NSDocument] object, or `nil` if the document could not
// be created.
//
// # Discussion
// 
// The URL is specified by `absoluteURL` and the document type by `typeName`.
// If not successful, the method returns `nil` after setting `outError` to
// point to an NSError that encapsulates the reason why the document could not
// be instantiated. The default implementation of this method calls
// [DocumentClassForType] to find out the class of document to instantiate,
// allocates a document object, and initializes it by sending it an
// [InitWithContentsOfURLOfTypeError] message.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/makeDocument(withContentsOf:ofType:)
func (d NSDocumentController) MakeDocumentWithContentsOfURLOfTypeError(url foundation.INSURL, typeName string) (INSDocument, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("makeDocumentWithContentsOfURL:ofType:error:"), url, objc.String(typeName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSDocument{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSDocumentFromID(rv), nil

}
// Instantiates a new untitled document of the specified type and returns it
// if successful.
//
// typeName: The type of the document.
//
// # Discussion
// 
// The document type is specified by `typeName`. If not successful, the method
// returns `nil` after setting `outError` to point to an [NSError] object that
// encapsulates the reason why a new untitled document could not be
// instantiated. The default implementation of this method calls
// [DocumentClassForType] to find out the class of document to instantiate,
// then allocates and initializes a document by sending it
// [InitWithTypeError].
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/makeUntitledDocument(ofType:)
func (d NSDocumentController) MakeUntitledDocumentOfTypeError(typeName string) (INSDocument, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("makeUntitledDocumentOfType:error:"), objc.String(typeName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSDocument{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSDocumentFromID(rv), nil

}
// Reopens a document, optionally located by a URL, by reading the contents
// for the document from another URL, optionally presents its user interface,
// and calls the passed-in completion handler.
//
// urlOrNil: The URL locating the reopened document, unless `nil`. A `nil` parameter
// value indicates that the reopened document is to have no [FileURL], like an
// untitled document.
//
// contentsURL: The URL (which may or may not be different from the URL of the reopened
// document) of the document from which the contents are read.
//
// displayDocument: If [true], displays the document’s user interface.
// //
// [true]: https://developer.apple.com/documentation/Swift/true
//
// completionHandler: The completion handler block object passed in to be called at some point in
// the future, perhaps after the method invocation has returned. The
// completion handler must be called on the main thread.
// 
// The block takes three arguments:
// 
// `document`: The document that was opened, if successful. Otherwise, `nil`.
// `documentWasAlreadyOpen`: Whether the document was already open or being
// opened when this method was called. `error`: If not successful, an
// [NSError] object that encapsulates the reason why the document could not be
// opened.
//
// # Discussion
// 
// The default implementation of this method is very similar to
// [OpenDocumentWithContentsOfURLDisplayCompletionHandler], the primary
// difference being that it calls [DocumentForURLWithContentsOfURLOfTypeError]
// instead of [DocumentWithContentsOfURLOfTypeError].
// 
// You can override this method to customize how documents are reopened during
// application launching by the restorable state mechanism introduced in OS X
// v10.7. Its implementation, however, is somewhat complex, so you should
// generally investigate overriding one of the methods that it calls instead.
// However, you can override this method to do additional work before calling
// the underlying method on `super`. You can also call the underlying method
// on `super` with a custom completion handler that performs additional work
// before calling the original completion handler.
// 
// Applications probably do not need to call this method directly.
// 
// For backward binary compatibility with OS X v10.6 and earlier, the default
// implementation of this method calls `[self url contentsURL &anError]` if
// that method is overridden and this one is not, instead of calling
// [DocumentForURLWithContentsOfURLOfTypeError] and all the rest.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/reopenDocument(for:withContentsOf:display:completionHandler:)
func (d NSDocumentController) ReopenDocumentForURLWithContentsOfURLDisplayCompletionHandler(urlOrNil foundation.INSURL, contentsURL foundation.INSURL, displayDocument bool, completionHandler DocumentErrorHandler) {
_block3, _cleanup3 := NewDocumentErrorBlock(completionHandler)
	defer _cleanup3()
	objc.Send[objc.ID](d.ID, objc.Sel("reopenDocumentForURL:withContentsOfURL:display:completionHandler:"), urlOrNil, contentsURL, displayDocument, _block3)
}
// Adds the given document to the list of open documents.
//
// document: The document to add.
//
// # Discussion
// 
// The `open...` methods automatically call [AddDocument]. This method is
// mostly provided for subclasses that want to know when documents arrive.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/addDocument(_:)
func (d NSDocumentController) AddDocument(document INSDocument) {
	objc.Send[objc.ID](d.ID, objc.Sel("addDocument:"), document)
}
// Returns the document object whose window controller owns a specified
// window.
//
// window: The window owned by the window controller.
//
// # Return Value
// 
// The document object whose window controller owns `window`. Returns `nil` if
// `window` is `nil`, if `window` has no window controller, or if the window
// controller does not have an association with an instance of [NSDocument].
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/document(for:)-a5yd
func (d NSDocumentController) DocumentForWindow(window INSWindow) INSDocument {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("documentForWindow:"), window)
	return NSDocumentFromID(rv)
}
// Removes the given document from the list of open documents.
//
// document: The document to remove.
//
// # Discussion
// 
// A document will automatically call [RemoveDocument] when it closes. This
// method is mostly provided for subclasses that want to know when documents
// close.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/removeDocument(_:)
func (d NSDocumentController) RemoveDocument(document INSDocument) {
	objc.Send[objc.ID](d.ID, objc.Sel("removeDocument:"), document)
}
// Returns the [NSDocument] subclass associated with a given document type.
//
// typeName: The name of a document type, specified by [CFBundleTypeName] in the
// application’s `Info.Plist()` file.
// 
// The document type must be one the receiver can read.
//
// # Return Value
// 
// Returns the [NSDocument] subclass associated with `documentTypeName`. If
// the class cannot be found, returns `nil`.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/documentClass(forType:)
func (d NSDocumentController) DocumentClassForType(typeName string) objc.Class {
	rv := objc.Send[objc.Class](d.ID, objc.Sel("documentClassForType:"), objc.String(typeName))
	return rv
}
// Returns the descriptive name for the specified document type, which is used
// in the File Format pop-up menu of the Save As dialog.
//
// typeName: The name of a document type, specified by [CFBundleTypeName] in the
// application’s `Info.Plist()` file.
//
// # Return Value
// 
// The descriptive name for the document type specified by `documentTypeName`.
// If there is no descriptive name, returns `documentTypeName`.
//
// # Discussion
// 
// For a document-based application, supported document types are specified in
// the `Info.Plist()` file by the [CFBundleDocumentTypes] array. Each document
// type is specified by a dictionary in this array, and is named by the
// [CFBundleTypeName] attribute. You can provide a descriptive, localized,
// representation of this name by providing a corresponding entry in the
// `InfoPlist.Strings()` file(s). For example, given an `Info.Plist()` file
// that contains the following fragment:
// 
// you could provide a descriptive name by adding an entry in the
// `InfoPlist.Strings()` file:
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/displayName(forType:)
func (d NSDocumentController) DisplayNameForType(typeName string) string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("displayNameForType:"), objc.String(typeName))
	return foundation.NSStringFromID(rv).String()
}
// Returns, for a specified URL, the document type identifier to use when
// opening the document at that location, if successful.
//
// url: The URL to use for locating the type identifier.
//
// # Discussion
// 
// The URL is represented by `url`. If not successful, the method returns
// `nil` after setting `outError` to point to an [NSError] object that
// encapsulates the reason why the document type could not be determined, or
// the fact that the document type is unrecognized.
// 
// You can override this method to customize type determination for documents
// being opened.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/typeForContents(of:)
func (d NSDocumentController) TypeForContentsOfURLError(url foundation.INSURL) (string, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("typeForContentsOfURL:error:"), url, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return "", foundation.NSErrorFrom(errorPtr)
	}
	return objc.IDToString(rv), nil

}
// Iterates through all the open documents and tries to close them one by one
// using the specified delegate.
//
// delegate: The object responsible for closing the document.
//
// didCloseAllSelector: The selector to call after all documents have been closed.
//
// contextInfo: A pointer to user-supplied data.
//
// # Discussion
// 
// Each [NSDocument] object is sent
// [CanCloseDocumentWithDelegateShouldCloseSelectorContextInfo], which, if the
// document is dirty, gives it a chance to refuse to close or to save itself
// first. This method may ask whether to save or to perform a save.
// 
// The `didCloseAllSelector` callback method is called with [true] if all
// documents are closed, and [false] otherwise. Pass the `contextInfo` object
// with the callback. The `didCloseAllSelector` callback method should have
// the following signature:
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/closeAllDocuments(withDelegate:didCloseAllSelector:contextInfo:)
func (d NSDocumentController) CloseAllDocumentsWithDelegateDidCloseAllSelectorContextInfo(delegate objectivec.IObject, didCloseAllSelector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](d.ID, objc.Sel("closeAllDocumentsWithDelegate:didCloseAllSelector:contextInfo:"), delegate, didCloseAllSelector, contextInfo)
}
// Displays an alert asking if the user wants to review unsaved documents,
// quit regardless of unsaved documents, or cancel the save operation.
//
// title: The title of the alert.
//
// cancellable: A Boolean indicating whether the operation can be canceled.
//
// delegate: The object that calls the selector.
//
// didReviewAllSelector: The selector to call when all documents have been reviewed.
//
// contextInfo: A pointer to user-supplied data.
//
// # Discussion
// 
// Assigns `delegate` to the panel. Calls `didReviewAllSelector` with [true]
// if quit without saving is chosen or if there are no dirty documents, and
// [false] otherwise. If the user selects the “Review Unsaved” option,
// [CloseAllDocumentsWithDelegateDidCloseAllSelectorContextInfo] is called.
// This method is called when the user chooses the Quit menu command, and also
// when the computer power is being turned off. Note that `title` is ignored.
// Pass the `contextInfo` object with the callback.
// 
// The `didReviewAllSelector` callback method should have the following
// signature:
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/reviewUnsavedDocuments(withAlertTitle:cancellable:delegate:didReviewAllSelector:contextInfo:)
func (d NSDocumentController) ReviewUnsavedDocumentsWithAlertTitleCancellableDelegateDidReviewAllSelectorContextInfo(title string, cancellable bool, delegate objectivec.IObject, didReviewAllSelector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](d.ID, objc.Sel("reviewUnsavedDocumentsWithAlertTitle:cancellable:delegate:didReviewAllSelector:contextInfo:"), objc.String(title), cancellable, delegate, didReviewAllSelector, contextInfo)
}
// An action method called by the New menu command, this method creates a new
// [NSDocument] object and adds it to the list of such objects managed by the
// document controller.
//
// # Discussion
// 
// This method calls [OpenUntitledDocumentAndDisplayError].
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/newDocument(_:)
func (d NSDocumentController) NewDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("newDocument:"), sender)
}
// An action method called by the Open menu command, it runs the modal Open
// panel and, based on the selected filenames, creates one or more
// [NSDocument] objects from the contents of the files.
//
// # Discussion
// 
// The method adds the newly created objects to the list of [NSDocument]
// objects managed by the document controller. This method calls
// [OpenDocumentWithContentsOfURLDisplayCompletionHandler], which actually
// creates the [NSDocument] objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/openDocument(_:)
func (d NSDocumentController) OpenDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("openDocument:"), sender)
}
// As the action method called by the Save All command, saves all open
// documents of the application that need to be saved.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/saveAllDocuments(_:)
func (d NSDocumentController) SaveAllDocuments(sender objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("saveAllDocuments:"), sender)
}
// Presents a nonmodal Open dialog that displays files you can open from a
// list of UTIs.
//
// openPanel: The Open dialog to present.
//
// inTypes: A list of file types that the user can choose from in the Open dialog.
//
// completionHandler: The completion handler that runs when the user clicks the OK or Cancel
// button in the Open dialog.
// 
// The block takes the following parameter:
// 
// `result`: Either [OKButton] or [CancelButton], depending on which button
// the user clicks to dismiss the dialog.
//
// # Discussion
// 
// [OpenDocument] and [BeginOpenPanelWithCompletionHandler] call this method
// to do the actual work. You typically don’t call this method directly.
// Override this method as necessary to customize the Open dialog or to alter
// the list of UTIs in the `inTypes` parameter.
// 
// You can also override this method if you want to perform additional cleanup
// (for example, if you customize the Open dialog and need to tear down an
// accessory view). Your overridden implementation needs to call the
// underlying method on `super`, passing a custom completion handler. That
// handler does the additional cleanup work, and then calls the completion
// handler block that the caller provides.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/beginOpenPanel(_:forTypes:completionHandler:)
func (d NSDocumentController) BeginOpenPanelForTypesCompletionHandler(openPanel INSOpenPanel, inTypes []string, completionHandler IntHandler) {
_block2, _cleanup2 := NewIntBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](d.ID, objc.Sel("beginOpenPanel:forTypes:completionHandler:"), openPanel, inTypes, _block2)
}
// Presents a modal Open dialog and limits selection to specific file types.
//
// openPanel: The open panel to display.
//
// types: An array of allowable types to open.
//
// # Discussion
// 
// This method is called by the [URLsFromRunningOpenPanel] method. It calls
// the [NSOpenPanel] [runModalForTypes:] method, passing the `openPanel`
// object and the file extensions associated with a document type. The
// `extensions` parameter may also contain encoded HFS file types as well as
// filename extensions.
//
// [runModalForTypes:]: https://developer.apple.com/documentation/AppKit/NSOpenPanel/runModalForTypes:
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/runModalOpenPanel(_:forTypes:)
func (d NSDocumentController) RunModalOpenPanelForTypes(openPanel INSOpenPanel, types []string) int {
	rv := objc.Send[int](d.ID, objc.Sel("runModalOpenPanel:forTypes:"), openPanel, objectivec.StringSliceToNSArray(types))
	return rv
}
// An array of URLs that correspond to the selected files in a running Open
// dialog.
//
// # Discussion
// 
// Accessing this property creates an [NSOpenPanel] object and runs it using
// the [RunModalOpenPanelForTypes] method. When the user dismisses the panel,
// the returned value is an array of URLs corresponding to the files chosen by
// the user. The value is `nil` if the user cancels the Open panel or makes no
// selection.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/urlsFromRunningOpenPanel()
func (d NSDocumentController) URLsFromRunningOpenPanel() []foundation.NSURL {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("URLsFromRunningOpenPanel"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSURL {
		return foundation.NSURLFromID(id)
	})
}
// Empties the recent documents list for the application.
//
// # Discussion
// 
// This is the action for the Clear menu command, but it can be called
// directly if necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/clearRecentDocuments(_:)
func (d NSDocumentController) ClearRecentDocuments(sender objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("clearRecentDocuments:"), sender)
}
// Adds or replaces an Open Recent menu item corresponding to the data located
// by the URL.
//
// url: The URL to evaluate.
//
// # Discussion
// 
// [NSDocument] automatically calls this method when appropriate for
// [NSDocument]-based applications. Applications not based on [NSDocument]
// must also implement the [ApplicationOpenFile] method in the application
// delegate to handle requests from the Open Recent menu command. You can
// override this method in an [NSDocument]-based application to prevent
// certain kinds of documents from getting into the list (but you have to
// identify them by URL).
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/noteNewRecentDocumentURL(_:)
func (d NSDocumentController) NoteNewRecentDocumentURL(url foundation.INSURL) {
	objc.Send[objc.ID](d.ID, objc.Sel("noteNewRecentDocumentURL:"), url)
}
// Adds or replaces an Open Recent menu item corresponding to the document.
//
// document: The document to evaluate.
//
// # Discussion
// 
// This method constructs a URL and calls [NoteNewRecentDocumentURL].
// Subclasses might override this method to prevent certain documents or kinds
// of documents from getting into the list.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/noteNewRecentDocument(_:)
func (d NSDocumentController) NoteNewRecentDocument(document INSDocument) {
	objc.Send[objc.ID](d.ID, objc.Sel("noteNewRecentDocument:"), document)
}
// Returns a Boolean value that indicates whether a given user interface item
// should be enabled.
//
// item: The user interface item to validate. You can send `anItem` the [Action] and
// [Tag] messages.
//
// # Return Value
// 
// [true] if `item` should be enabled, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Subclasses can override this method to perform additional validations.
// Subclasses should call the underlying method on `super` in their
// implementation for items they don’t handle themselves.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/validateUserInterfaceItem(_:)
func (d NSDocumentController) ValidateUserInterfaceItem(item NSValidatedUserInterfaceItem) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("validateUserInterfaceItem:"), item)
	return rv
}
// Returns a menu item that your app uses for sharing the current document.
//
// # Return Value
// 
// An [NSMenuItem] for the Share menu.
//
// # Discussion
// 
// Use this method to perform custom placement of the Share menu if your
// [NSDocument] subclass returns `false` for [AllowsAutomaticShareMenu].
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/standardShareMenuItem()
func (d NSDocumentController) StandardShareMenuItem() INSMenuItem {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("standardShareMenuItem"))
	return NSMenuItemFromID(rv)
}
// Presents an error alert to the user as a modal panel.
//
// error: An object containing the error to present.
//
// # Return Value
// 
// Returns [true] if error recovery was done; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method does not return until the user dismisses the alert and, if the
// error has recovery options and a recovery delegate, the error’s recovery
// delegate is sent an `` message.
// 
// The default [NSDocumentController] implementation of this method is
// equivalent to that of [NSResponder] while treating the application object
// as the next responder and forwarding error presentation messages to it.
// (The default [NSDocument] implementation of this method treats the shared
// [NSDocumentController] instance as the next responder and forwards these
// messages to it.) The default implementations of several
// [NSDocumentController] methods call this method.
// 
// The default implementation of this method calls [WillPresentError] to give
// subclasses an opportunity to customize error presentation. You should not
// override this method but should instead override [WillPresentError].
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/presentError(_:)
func (d NSDocumentController) PresentError(error_ foundation.INSError) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("presentError:"), error_)
	return rv
}
// Presents an error alert to the user as a modal panel.
//
// error: The error to present.
//
// window: The window to present the error.
//
// delegate: The object to receive the selector.
//
// didPresentSelector: The selector to send to the delegate.
//
// contextInfo: A pointer to user-supplied data.
//
// # Discussion
// 
// When the user dismisses the alert and any recovery possible for the error
// and chosen by the user has been attempted, sends the message
// `didPresentSelector` to the specified `delegate`. The method selected by
// `didPresentSelector` must have the same signature as:
// 
// The default [NSDocumentController] implementation of this method is
// equivalent to that of [NSResponder] while treating the application object
// as the next responder and forwarding error presentation messages to it.
// (The default [NSDocument] implementation of this method treats the shared
// [NSDocumentController] instance as the next responder and forwards these
// messages to it.)
// 
// The default implementation of this method calls [WillPresentError] to give
// subclasses an opportunity to customize error presentation. You should not
// override this method but should instead override [WillPresentError].
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/presentError(_:modalFor:delegate:didPresent:contextInfo:)
func (d NSDocumentController) PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error_ foundation.INSError, window INSWindow, delegate objectivec.IObject, didPresentSelector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](d.ID, objc.Sel("presentError:modalForWindow:delegate:didPresentSelector:contextInfo:"), error_, window, delegate, didPresentSelector, contextInfo)
}
// Indicates an error condition and provides the opportunity to return the
// same or a different error.
//
// error: The error to present.
//
// # Discussion
// 
// The default implementation of this method merely returns the passed-in
// error. The returned error may simply be forwarded to the application
// object.
// 
// You can override this method to customize the presentation of errors by
// examining the passed-in error and, for example, returning more specific
// information. When you override this method always check the [NSError]
// object’s domain and code to discriminate between errors whose
// presentation you want to customize and those you don’t. For errors you
// don’t want to customize, call the superclass implementation, passing the
// original error.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/willPresentError(_:)
func (d NSDocumentController) WillPresentError(error_ foundation.INSError) foundation.INSError {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("willPresentError:"), error_)
	return foundation.NSErrorFromID(rv)
}
// Implemented to override the default action of enabling or disabling a
// specific menu item.
//
// menuItem: An [NSMenuItem] object that represents the menu item.
//
// # Return Value
// 
// [true] to enable `menuItem`, [false] to disable it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The object implementing this method must be the target of `menuItem`. You
// can determine which menu item `menuItem` is by querying it for its tag or
// action.
// 
// The following example disables the menu item associated with the
// `nextRecord` action method when the selected line in a table view is the
// last one; conversely, it disables the menu item with `priorRecord` as its
// action method when the selected row is the first one in the table view.
// (The `countryOrRegionKeys` array contains names that appear in the table
// view.)
//
// See: https://developer.apple.com/documentation/AppKit/NSMenuItemValidation/validateMenuItem(_:)
func (d NSDocumentController) ValidateMenuItem(menuItem INSMenuItem) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("validateMenuItem:"), menuItem)
	return rv
}
func (d NSDocumentController) EncodeWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeWithCoder:"), coder)
}

// Asks the class to provide a new window for the specified identifier.
//
// identifier: The unique interface item identifier string that was previously associated
// with the window. Use this string to determine which window to create.
//
// state: A coder object containing the window state information. This coder object
// contains the combined restorable state of the window, which can include the
// state of the window, its delegate, window controller, and document object.
// You can use this state to determine which window to create.
//
// completionHandler: A Block object to execute with the results of creating the window. You must
// execute this block at some point but may do so after the method returns if
// needed. This block takes the following parameters:
// 
// - The window that was created or `nil` if the window could not be created.
// - An error object if the window was not recognized or could not be created
// for whatever reason; otherwise, specify `nil`.
//
// # Discussion
// 
// At launch time, the application object uses this method to recreate a
// window that was preserved during a previous running of your application.
// Your implementation of this method can use the `identifier` string and the
// data in the `state` coder object to determine the type of window and
// associated objects to create. You only need to create the window. You do
// not need to restore the configuration of the window using the provided
// state. AppKit does that for you when you execute the completion handler
// with a valid window.
// 
// For windows that your restoration class recognizes, create (or obtain) the
// window object and execute the completion handler, passing the window to it.
// If your restoration class does not recognize the window, or recognizes it
// but cannot recreate it for whatever reason, pass an error object to the
// completion handler instead.
// 
// If you plan to invoke `completionHandler` after this method returns, you
// must copy the Block object yourself and release it when you are done with
// it. For example, you can copy the block and submit it to a queue for
// execution, after which you can release the block.
//
// See: https://developer.apple.com/documentation/AppKit/NSWindowRestoration/restoreWindow(withIdentifier:state:completionHandler:)
func (_NSDocumentControllerClass NSDocumentControllerClass) RestoreWindowWithIdentifierStateCompletionHandler(identifier NSUserInterfaceItemIdentifier, state foundation.INSCoder, completionHandler WindowErrorHandler) {
_block2, _cleanup2 := NewWindowErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](objc.ID(_NSDocumentControllerClass.class), objc.Sel("restoreWindowWithIdentifier:state:completionHandler:"), identifier, state, _block2)
}

// The document objects managed by the receiver.
//
// # Discussion
// 
// The array contains zero or more [NSDocument] objects.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/documents
func (d NSDocumentController) Documents() []NSDocument {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("documents"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSDocument {
		return NSDocumentFromID(id)
	})
}
// The document object associated with the main window.
//
// # Discussion
// 
// The value of this property is `nil` if it is called when the app is not
// active. This can occur during processing of a drag-and-drop operation, for
// example, in an implementation of ``. In such a case, send the following
// message instead from an [NSView] subclass associated with the document:
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/currentDocument
func (d NSDocumentController) CurrentDocument() INSDocument {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("currentDocument"))
	return NSDocumentFromID(objc.ID(rv))
}
// A Boolean value indicating whether the receiver has any documents with
// unsaved changes.
//
// # Discussion
// 
// The value of this property is [true] if the document controller contains
// documents with unsaved changes; otherwise, the value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/hasEditedDocuments
func (d NSDocumentController) HasEditedDocuments() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("hasEditedDocuments"))
	return rv
}
// An array of strings representing the custom document classes supported by
// this app.
//
// # Discussion
// 
// The items in the array are [NSString] objects, each of which represents the
// name of a document subclasses supported by the app. The document class
// names are derived from the app’s `Info.Plist()`. You can override this
// property and use it to return the names of document classes that are
// dynamically loaded from plugins.
//
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/documentClassNames
func (d NSDocumentController) DocumentClassNames() []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("documentClassNames"))
	return objc.ConvertSliceToStrings(rv)
}
// Returns the name of the document type that should be used when creating new
// documents.
//
// # Discussion
// 
// The default implementation of this method returns the first Editor type
// declared by the [CFBundleDocumentTypes] array in the application’s
// `Info.Plist()`, or returns `nil` if no Editor type is declared. You can
// override it to customize the type of document that is created when, for
// instance, the user chooses New in the File menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/defaultType
func (d NSDocumentController) DefaultType() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("defaultType"))
	return foundation.NSStringFromID(rv).String()
}
// The time interval (in seconds) for periodic autosaving.
//
// # Discussion
// 
// A value of 0 indicates that periodic autosaving should not be done at all.
// The [NSDocumentController] object uses this number as the amount of time to
// wait between detecting that a document has unautosaved changes and sending
// the document an
// [AutosaveDocumentWithDelegateDidAutosaveSelectorContextInfo] message. The
// default value is `0`.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/autosavingDelay
func (d NSDocumentController) AutosavingDelay() float64 {
	rv := objc.Send[float64](d.ID, objc.Sel("autosavingDelay"))
	return rv
}
func (d NSDocumentController) SetAutosavingDelay(value float64) {
	objc.Send[struct{}](d.ID, objc.Sel("setAutosavingDelay:"), value)
}
// The directory path to use as the starting point in the Open dialog.
//
// # Discussion
// 
// The first valid directory from the following list is returned:
// 
// - The directory location where the current document was last saved - The
// last directory visited in the Open panel - The user’s home directory
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/currentDirectory
func (d NSDocumentController) CurrentDirectory() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("currentDirectory"))
	return foundation.NSStringFromID(rv).String()
}
// The maximum number of items that may be presented in the standard Open
// Recent menu.
//
// # Discussion
// 
// A value of 0 indicates that [NSDocumentController] will not attempt to add
// an Open Recent menu to your application’s File menu, although
// [NSDocumentController] will not attempt to remove any preexisting Open
// Recent menu item. The default implementation returns a value that is
// subject to change and may or may not be derived from a setting made by the
// user in System Preferences.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/maximumRecentDocumentCount
func (d NSDocumentController) MaximumRecentDocumentCount() uint {
	rv := objc.Send[uint](d.ID, objc.Sel("maximumRecentDocumentCount"))
	return rv
}
// The list of recent-document URLs.
//
// # Discussion
// 
// This property contains an array of [NSURL] objects corresponding to the
// recently opened documents. Do not override this property.
//
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/recentDocumentURLs
func (d NSDocumentController) RecentDocumentURLs() []foundation.NSURL {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("recentDocumentURLs"))
	return objc.ConvertSlice(rv, func(id objc.ID) foundation.NSURL {
		return foundation.NSURLFromID(id)
	})
}
// A Boolean value that the system uses to insert a Share menu in the File
// menu.
//
// # Discussion
// 
// If your application has any [NSDocument] subclasses with [AutosavesInPlace]
// set to `true`, the system defaults `allowsAutomaticShareMenu` to `true`. To
// disable the Share menu entirely, or to enable custom placement or
// construction of the share menu, override this property to return `false` in
// your app.
// 
// The system may not insert a Share menu if `allowsAutomaticShareMenu` is
// `true` and [NSDocumentController] detects that the app has a Share menu.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/allowsAutomaticShareMenu
func (d NSDocumentController) AllowsAutomaticShareMenu() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("allowsAutomaticShareMenu"))
	return rv
}

// Returns the shared [NSDocumentController] instance.
//
// # Return Value
// 
// The shared [NSDocumentController] instance.
// 
// # Discussion
// 
// If an [NSDocumentController] instance doesn’t exist yet, it is created.
// 
// Initialization reads in the document types from the [CFBundleDocumentTypes]
// property list (in `Info.Plist()`), registers the instance for
// [willPowerOffNotification]s, and turns on the flag indicating that document
// user interfaces should be visible. You should always obtain your
// application’s [NSDocumentController] using this method.
//
// [willPowerOffNotification]: https://developer.apple.com/documentation/AppKit/NSWorkspace/willPowerOffNotification
//
// See: https://developer.apple.com/documentation/AppKit/NSDocumentController/shared
func (_NSDocumentControllerClass NSDocumentControllerClass) SharedDocumentController() NSDocumentController {
	rv := objc.Send[objc.ID](objc.ID(_NSDocumentControllerClass.class), objc.Sel("sharedDocumentController"))
	return NSDocumentControllerFromID(objc.ID(rv))
}

			// Protocol methods for NSMenuItemValidation
			

			// Protocol methods for NSUserInterfaceValidations
			

			// Protocol methods for NSWindowRestoration
			

// RestoreWindowWithIdentifierState is a synchronous wrapper around [NSDocumentController.RestoreWindowWithIdentifierStateCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (dc NSDocumentControllerClass) RestoreWindowWithIdentifierState(ctx context.Context, identifier NSUserInterfaceItemIdentifier, state foundation.INSCoder) (*NSWindow, error) {
	type result struct {
		val *NSWindow
		err error
	}
	done := make(chan result, 1)
	dc.RestoreWindowWithIdentifierStateCompletionHandler(identifier, state, func(val *NSWindow, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

