// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"context"
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"errors"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSDocument] class.
var (
	_NSDocumentClass     NSDocumentClass
	_NSDocumentClassOnce sync.Once
)

func getNSDocumentClass() NSDocumentClass {
	_NSDocumentClassOnce.Do(func() {
		_NSDocumentClass = NSDocumentClass{class: objc.GetClass("NSDocument")}
	})
	return _NSDocumentClass
}

// GetNSDocumentClass returns the class object for NSDocument.
func GetNSDocumentClass() NSDocumentClass {
	return getNSDocumentClass()
}

type NSDocumentClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSDocumentClass) Alloc() NSDocument {
	rv := objc.Send[NSDocument](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract class that defines the interface for macOS documents.
//
// # Overview
// 
// A document is an object that can internally represent data displayed in a
// window and that can read data from and write data to a file or file
// package. Documents create and manage one or more window controllers and are
// in turn managed by a document controller. Documents respond to
// first-responder action messages to save, revert, and print their data.
// 
// Conceptually, a document is a container for a body of information
// identified by a name under which it is stored in a disk file. In this
// sense, however, the document is not the same as the file but is an object
// in memory that owns and manages the document data. In the context of
// AppKit, a document is an instance of a custom [NSDocument] subclass that
// knows how to represent internally, in one or more formats, persistent data
// that is displayed in windows.
// 
// A document can read that data from a file and write it to a file. It is
// also the first-responder target for many menu commands related to
// documents, such as Save, Revert, and Print. A document manages its
// window’s edited status and is set up to perform undo and redo operations.
// When a window is closing, the document is asked before the window delegate
// to approve the closing.
// 
// [NSDocument] is one of the triad of AppKit classes that establish an
// architectural basis for document-based apps (the others being
// [NSDocumentController] and [NSWindowController]).
// 
// For more information about using [NSDocument] in a document-based app, see
// [Developing a Document-Based App].
// 
// # Subclassing NSDocument
// 
// The [NSDocument] class is designed to be subclassed. That is, the
// [NSDocument] class is abstract, and your app must create at least one
// [NSDocument] subclass in order to use the document architecture. To create
// a useful [NSDocument] subclass, you must override some methods, and you can
// optionally override others.
// 
// The [NSDocument] class itself knows how to handle document data as
// undifferentiated lumps; although it understands that these lumps are typed,
// it knows nothing about particular types. In their overrides of the
// data-based reading and writing methods, subclasses must add the knowledge
// of particular types and how data of the document’s native type is
// structured internally. Subclasses are also responsible for the creation of
// the window controllers that manage document windows and for the
// implementation of undo and redo. The [NSDocument] class takes care of much
// of the rest, including generally managing the state of the document.
// 
// See [Document-Based App Programming Guide for Mac] for more information
// about creating subclasses of [NSDocument], particularly the list of
// primitive methods that subclasses must override and those that you can
// optionally override.
// 
// # Document Saving Behavior
// 
// The [NSDocument] class implements document saving in a way that preserves,
// when possible, various attributes of each document, including:
// 
// - Creation date - Permissions/privileges - Location of the document’s
// icon in its parent folder’s Icon View Finder window - Value of the
// document’s Show Extension setting
// 
// Care is also taken to save documents in a way that does not break any
// user-created aliases that may point to documents. As a result, some methods
// in any class of [NSDocument] may be invoked with parameters that do not
// have the same meaning as they did in early releases of macOS. It is
// important that overrides of [NSDocument.WriteToURLOfTypeError] and
// [NSDocument.WriteToURLOfTypeForSaveOperationOriginalContentsURLError] make no
// assumptions about the file paths passed as parameters, including:
// 
// - The location to which the file is being written. This location might be a
// hidden temporary directory. - The name of the file being written. It is
// possible that this file has no obvious relation to the document name. - The
// relation of any file being passed, including the original file, to the
// value in [NSDocument.FileURL].
// 
// When updating your app to link against OS X v10.5, keep in mind that it is
// usually more appropriate to invoke in your app code one of the [NSDocument]
// `save...` methods than one of the `write...` methods. The `write...`
// methods are there primarily for you to override. The
// [saveToURL:ofType:forSaveOperation:error:] method that is meant always to
// be invoked during document saving, sets the [NSDocument.FileModificationDate] property
// with the file’s new modification date after it has been written (for
// [NSDocument.SaveOperationType.saveOperation] and
// [NSDocument.SaveOperationType.saveAsOperation] only).
// 
// Likewise, it’s usually more appropriate to invoke in your app code one of
// the [NSDocument] `revert...` methods than one of the `read...` methods. The
// `read...` methods are there primarily for you to override. The
// [NSDocument.RevertToContentsOfURLOfTypeError] method that is meant always to be
// invoked during rereading of an open document, sets the
// [NSDocument.FileModificationDate] property with the file’s modification date after
// it has been read.
// 
// # iCloud Support
// 
// The [NSDocument] class implements the file coordination support that is
// required for an iCloud-enabled, document-based Mac app (see [How iCloud
// Document Storage Works] in [iCloud Design Guide]). In addition, this
// class’s methods for moving and renaming documents, new in OS X v10.8,
// ensure that these operations are performed in a safe manner for
// iCloud-enabled apps.
// 
// # Multicore Considerations
// 
// In macOS 10.6 and later, [NSDocument] supports the ability to open multiple
// documents concurrently. However, this support requires the cooperation of
// the document object. If your document subclass is able to read specific
// document types independently of other similar documents, you should
// override the [NSDocument.CanConcurrentlyReadDocumentsOfType] class method and return
// [true] for the appropriate document types. If specific document types rely
// on shared state information, however, you should return [false] for those
// types.
//
// [Developing a Document-Based App]: https://developer.apple.com/documentation/AppKit/developing-a-document-based-app
// [Document-Based App Programming Guide for Mac]: https://developer.apple.com/library/archive/documentation/DataManagement/Conceptual/DocBasedAppProgrammingGuideForOSX/Introduction/Introduction.html#//apple_ref/doc/uid/TP40011179
// [How iCloud Document Storage Works]: https://developer.apple.com/library/archive/documentation/General/Conceptual/iCloudDesignGuide/Chapters/DesigningForDocumentsIniCloud.html#//apple_ref/doc/uid/TP40012094-CH2-SW10
// [NSDocument.SaveOperationType.saveAsOperation]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/saveAsOperation
// [NSDocument.SaveOperationType.saveOperation]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/saveOperation
// [false]: https://developer.apple.com/documentation/Swift/false
// [iCloud Design Guide]: https://developer.apple.com/library/archive/documentation/General/Conceptual/iCloudDesignGuide/Chapters/Introduction.html#//apple_ref/doc/uid/TP40012094
// [saveToURL:ofType:forSaveOperation:error:]: https://developer.apple.com/documentation/AppKit/NSDocument/saveToURL:ofType:forSaveOperation:error:
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Creating a Document Object
//
//   - [NSDocument.InitWithContentsOfURLOfTypeError]: Initializes a document located by a URL of a specified type.
//   - [NSDocument.InitForURLWithContentsOfURLOfTypeError]: Initializes a document with the specified contents, and places the resulting document’s file at the designated location.
//   - [NSDocument.InitWithTypeError]: Initializes a document of a specified type.
//
// # Reading the Document’s Content
//
//   - [NSDocument.ReadFromURLOfTypeError]: Sets the contents of this document by reading from a file or file package, of a specified type, located by a URL.
//   - [NSDocument.ReadFromFileWrapperOfTypeError]: Sets the contents of this document by reading from a file wrapper of a specified type.
//   - [NSDocument.ReadFromDataOfTypeError]: Sets the contents of this document by reading from data of a specified type.
//
// # Writing the Document’s Content
//
//   - [NSDocument.CanAsynchronouslyWriteToURLOfTypeForSaveOperation]: Returns whether the receiver can concurrently write to a file or file package located by a URL, that is formatted for a specific type, for a specific kind of save operation.
//   - [NSDocument.UnblockUserInteraction]: Unblocks the main thread during asynchronous saving.
//   - [NSDocument.WriteToURLOfTypeError]: Writes the contents of the document to a file or file package located by a URL, that is formatted to a specified type.
//   - [NSDocument.WriteSafelyToURLOfTypeForSaveOperationError]: Writes the contents of the document to a file or file package located by a URL.
//   - [NSDocument.FileWrapperOfTypeError]: Creates and returns a file wrapper that contains the contents of the document, formatted to the specified type.
//   - [NSDocument.DataOfTypeError]: Creates and returns a data object that contains the contents of the document, formatted to a specified type.
//   - [NSDocument.WriteToURLOfTypeForSaveOperationOriginalContentsURLError]: Writes the contents of the document to a file or file package located by a URL.
//   - [NSDocument.SaveToURLOfTypeForSaveOperationDelegateDidSaveSelectorContextInfo]: Saves the contents of the document to a file or file package located by a URL, that is formatted to a specified type, for a particular kind of save operation.
//   - [NSDocument.SaveToURLOfTypeForSaveOperationCompletionHandler]: Saves the contents of the document to a file or file package located by a URL, that is formatted to a specified type, for a particular kind of save operation, and invokes the passed-in completion handler.
//   - [NSDocument.FileAttributesToWriteToURLOfTypeForSaveOperationOriginalContentsURLError]: Returns the attributes to write to the file or file package at the specified URL, and targeting the specified type of save operation.
//
// # Getting Document Metadata
//
//   - [NSDocument.FileURL]: The location of the document’s on-disk representation.
//   - [NSDocument.SetFileURL]
//   - [NSDocument.EntireFileLoaded]: A Boolean value that indicates whether the document’s file is completely loaded into memory.
//   - [NSDocument.FileModificationDate]: The last-known modification date of the document’s on-disk representation.
//   - [NSDocument.SetFileModificationDate]
//   - [NSDocument.KeepBackupFile]: A Boolean value that indicates whether the document archives previously saved versions of the document.
//   - [NSDocument.Draft]: A Boolean value that indicates whether the document is a draft that the user has not yet saved.
//   - [NSDocument.SetDraft]
//   - [NSDocument.FileType]: The name of the document type, as specified in the app’s information property-list file.
//   - [NSDocument.SetFileType]
//   - [NSDocument.DocumentEdited]: A Boolean value that indicates whether the document has unsaved changes.
//   - [NSDocument.InViewingMode]: A Boolean value that indicates whether the document is in read-only mode.
//
// # Managing File Type Information
//
//   - [NSDocument.WritableTypesForSaveOperation]: Returns the names of the types to which this document can be saved for a specified kind of save operation.
//   - [NSDocument.FileNameExtensionForTypeSaveOperation]: Returns a filename extension that can be appended to a base filename, for a specified file type and kind of save operation.
//
// # Creating and Managing Window Controllers
//
//   - [NSDocument.MakeWindowControllers]: Creates the window controller objects that the document uses to display its content.
//   - [NSDocument.AddWindowController]: Adds the specified window controller to the current document.
//   - [NSDocument.RemoveWindowController]: Removes the specified window controller from the receiver’s array of window controllers.
//   - [NSDocument.WindowControllers]: The document’s current window controllers.
//   - [NSDocument.WindowNibName]: The name of the document’s sole nib file.
//   - [NSDocument.WindowControllerDidLoadNib]: Called after one of the document’s window controllers loads its nib file.
//   - [NSDocument.WindowControllerWillLoadNib]: Called before one of the document’s window controllers loads its nib file.
//   - [NSDocument.ShouldCloseWindowControllerDelegateShouldCloseSelectorContextInfo]: Determines whether the system should close the document and its associated window.
//
// # Managing Document Windows
//
//   - [NSDocument.ShowWindows]: Displays all of the document’s windows, bringing them to the front and making them main or key as necessary.
//   - [NSDocument.SetWindow]: Sets the window outlet of this document to the specified value.
//   - [NSDocument.WindowForSheet]: Returns the document window to use as the parent of a document-modal sheet.
//   - [NSDocument.DisplayName]: The name of the document as displayed in the title bars of the document’s windows and in alert dialogs related to the document.
//   - [NSDocument.SetDisplayName]
//   - [NSDocument.DefaultDraftName]: Returns the default draft name for the document subclass.
//   - [NSDocument.EncodeRestorableStateWithCoderBackgroundQueue]: Saves the interface-related state of the document.
//
// # Configuring the Autosave Behavior
//
//   - [NSDocument.AutosavedContentsFileURL]: The location of the most recently autosaved document contents.
//   - [NSDocument.SetAutosavedContentsFileURL]
//   - [NSDocument.AutosavingFileType]: The document type to use for an autosave operation.
//   - [NSDocument.AutosavingIsImplicitlyCancellable]: A Boolean value that indicates whether you can cancel an in-progress autosave operation.
//
// # Autosaving the Document
//
//   - [NSDocument.CheckAutosavingSafetyAndReturnError]: Returns a Boolean value that indicates whether it is safe to autosave document changes.
//   - [NSDocument.HasUnautosavedChanges]: A Boolean value that indicates whether the document has changes that have not been autosaved.
//   - [NSDocument.ScheduleAutosaving]: Schedules periodic autosaving for the purpose of crash protection.
//   - [NSDocument.AutosaveDocumentWithDelegateDidAutosaveSelectorContextInfo]: Autosaves the document’s contents to an appropriate location in the file system.
//   - [NSDocument.AutosaveWithImplicitCancellabilityCompletionHandler]: Autosaves the document’s contents to an appropriate file-system location, as needed.
//   - [NSDocument.BackupFileURL]: The URL for the document’s backup file that was created during an autosave operation.
//
// # Browsing Document Versions
//
//   - [NSDocument.BrowseDocumentVersions]: Opens the Versions browser in the document’s main window.
//   - [NSDocument.BrowsingVersions]: A Boolean value that indicates whether the document is currently displaying the Versions browser.
//   - [NSDocument.StopBrowsingVersionsWithCompletionHandler]: Dismiss the Versions browser for the current document.
//
// # Storing Documents in iCloud
//
//   - [NSDocument.MoveDocumentToUbiquityContainer]: Moves the document to the user’s iCloud storage.
//
// # Managing Undo and Redo Actions
//
//   - [NSDocument.UndoManager]: The object that the document uses to support undo/redo operations.
//   - [NSDocument.SetUndoManager]
//   - [NSDocument.HasUndoManager]: A Boolean value that indicates whether the document owns an undo manager object.
//   - [NSDocument.SetHasUndoManager]
//
// # Updating the Document Change Count
//
//   - [NSDocument.UpdateChangeCountWithTokenForSaveOperation]: Updates the document’s change count settings after a successful save operation.
//   - [NSDocument.UpdateChangeCount]: Updates the receiver’s change count according to the given change type.
//   - [NSDocument.ChangeCountTokenForSaveOperation]: Returns an object that encapsulates the current record of document changes at the beginning of a save operation.
//
// # Handling Window Restoration
//
//   - [NSDocument.EncodeRestorableStateWithCoder]: Saves the interface-related state of the document.
//   - [NSDocument.RestoreStateWithCoder]: Restores the interface-related state of the document.
//   - [NSDocument.InvalidateRestorableState]: Marks the document’s interface-related state as dirty.
//   - [NSDocument.RestoreDocumentWindowWithIdentifierStateCompletionHandler]: Restores a window that was associated with a document, after that document is reopened.
//
// # Presenting a Save Panel
//
//   - [NSDocument.RunModalSavePanelForSaveOperationDelegateDidSaveSelectorContextInfo]: Presents a modal Save panel to the user, then tries to save the document if the user approves the operation.
//   - [NSDocument.PrepareSavePanel]: Tells the document to customize the specified Save panel.
//   - [NSDocument.ShouldRunSavePanelWithAccessoryView]: A Boolean value that indicates whether the document’s Save panel displays a list of supported writable document types.
//   - [NSDocument.FileTypeFromLastRunSavePanel]: The file type that was last selected in the Save panel.
//   - [NSDocument.FileNameExtensionWasHiddenInLastRunSavePanel]: A Boolean value that indicates whether the user chose to hide the document’s filename extension.
//
// # Supporting User Activities
//
//   - [NSDocument.UserActivity]: An object that encapsulates a user activity the document supports.
//   - [NSDocument.SetUserActivity]
//   - [NSDocument.UpdateUserActivityState]: Updates the state of the given user activity.
//   - [NSDocument.NSUserActivityDocumentURLKey]: The key that identifies the document associated with a user activity.
//
// # Performing Tasks Serially
//
//   - [NSDocument.PerformSynchronousFileAccessUsingBlock]: Waits for any scheduled file access to complete, then invokes the passed-in block.
//   - [NSDocument.PerformAsynchronousFileAccessUsingBlock]: Waits for any scheduled file access to complete but without blocking the main thread, then invokes the passed-in block.
//   - [NSDocument.PerformActivityWithSynchronousWaitingUsingBlock]: Waits for any work scheduled by previous invocations of this method to complete, then invokes the passed-in block.
//   - [NSDocument.ContinueActivityUsingBlock]: Continues to perform the task for a user activity object using a different block.
//   - [NSDocument.ContinueAsynchronousWorkOnMainThreadUsingBlock]: Invokes the passed-in block on the main thread.
//
// # Handling User Actions
//
//   - [NSDocument.PrintDocument]: Prints the receiver in response to the user choosing the Print menu command.
//   - [NSDocument.RunPageLayout]: The action method invoked in the receiver as first responder when the user chooses the Page Setup menu command.
//   - [NSDocument.RevertDocumentToSaved]: The action of the File menu item Revert in a document-based app.
//   - [NSDocument.SaveDocument]: The action method invoked in the receiver as first responder when the user chooses the Save menu command.
//   - [NSDocument.SaveDocumentAs]: The action method invoked in the receiver as first responder when the user chooses the Save As menu command.
//   - [NSDocument.SaveDocumentTo]: The action method invoked in the receiver as first responder when the user chooses the Save To menu command.
//   - [NSDocument.SaveDocumentWithDelegateDidSaveSelectorContextInfo]: Saves the document and delivers the results to the provided delegate object.
//
// # Closing the Document
//
//   - [NSDocument.CanCloseDocumentWithDelegateShouldCloseSelectorContextInfo]: Determines whether to close the document, prompting the user as needed to choose a course of action.
//   - [NSDocument.Close]: Closes all of the document’s windows and removes the document from its document controller.
//
// # Reverting the Document Contents
//
//   - [NSDocument.RevertToContentsOfURLOfTypeError]: Discards all unsaved document modifications and replaces the document’s contents by reading a file or file package located by a URL of a specified type.
//
// # Duplicating the Document
//
//   - [NSDocument.DuplicateAndReturnError]: Creates a new document whose contents are the same as the receiver and returns an error object if unsuccessful.
//   - [NSDocument.DuplicateDocument]: Creates a copy of the receiving document in response to the user choosing Duplicate from the File menu.
//   - [NSDocument.DuplicateDocumentWithDelegateDidDuplicateSelectorContextInfo]: Creates a new document whose contents are the same as the current document.
//
// # Renaming the Document
//
//   - [NSDocument.RenameDocument]: Renames the current document in response to the user choosing the Rename menu item.
//
// # Moving the Document
//
//   - [NSDocument.MoveDocument]: Moves the document to a new location in response to the user choosing the Move To… menu item.
//   - [NSDocument.MoveDocumentWithCompletionHandler]: Moves the document to a user-selected location.
//   - [NSDocument.MoveToURLCompletionHandler]: Moves the document’s file to the given URL.
//
// # Locking the Document
//
//   - [NSDocument.LockDocument]: Locks the document in response to the user choosing the Lock menu item.
//   - [NSDocument.UnlockDocument]: Unlocks the document in response to the user choosing the Unlock menu item.
//   - [NSDocument.LockDocumentWithCompletionHandler]: Prevents the user from making further changes to the document.
//   - [NSDocument.LockWithCompletionHandler]: Prevents the user from making changes to the document’s file.
//   - [NSDocument.UnlockDocumentWithCompletionHandler]: Allows the user to make modifications to the document.
//   - [NSDocument.UnlockWithCompletionHandler]: Allows the user to make modifications to the document’s file.
//   - [NSDocument.Locked]: A Boolean value that indicates whether or not the file can be written to.
//
// # Printing the Document
//
//   - [NSDocument.PrintInfo]: The printing information associated with the document.
//   - [NSDocument.SetPrintInfo]
//   - [NSDocument.PreparePageLayout]: Adds document-specific content to the Page Layout panel.
//   - [NSDocument.RunModalPageLayoutWithPrintInfoDelegateDidRunSelectorContextInfo]: Runs the modal page layout panel with the receiver’s printing information object.
//   - [NSDocument.RunModalPrintOperationDelegateDidRunSelectorContextInfo]: Runs the specified print operation modally.
//   - [NSDocument.ShouldChangePrintInfo]: Returns a Boolean value that indicates whether the document allows changes to the default printing information.
//   - [NSDocument.PrintDocumentWithSettingsShowPrintPanelDelegateDidPrintSelectorContextInfo]: Prints the document’s contents, optionally displaying a print panel to the user.
//   - [NSDocument.PrintOperationWithSettingsError]: Creates and returns a print operation for the document’s contents.
//   - [NSDocument.PDFPrintOperation]: A print operation you can use to create a PDF representation of the document’s current contents.
//   - [NSDocument.SaveDocumentToPDF]: Exports a PDF representation of the document’s current contents.
//
// # Sharing the Document
//
//   - [NSDocument.AllowsDocumentSharing]: A Boolean value that indicates whether the document is shareable from the standard Share menu.
//   - [NSDocument.PrepareSharingServicePicker]: Perform any custom setup associated with a sharing service picker.
//   - [NSDocument.ShareDocumentWithSharingServiceCompletionHandler]: Share the document’s file using the specified sharing service.
//
// # Handling Script Commands
//
//   - [NSDocument.HandleCloseScriptCommand]: Handles the Close AppleScript command by attempting to close the document.
//   - [NSDocument.HandlePrintScriptCommand]: Handles the Print AppleScript command by attempting to print the document.
//   - [NSDocument.HandleSaveScriptCommand]: Handles the Save AppleScript command by attempting to save the document.
//   - [NSDocument.ObjectSpecifier]: Returns the object specifier that represents the document.
//   - [NSDocument.LastComponentOfFileName]: The name of the document seen by the user in AppleScript.
//   - [NSDocument.SetLastComponentOfFileName]
//
// # Displaying Errors to the User
//
//   - [NSDocument.PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo]: Presents an error alert to the user as a modal panel.
//   - [NSDocument.PresentError]: Presents an error alert to the user as a modal panel.
//   - [NSDocument.WillPresentError]: Called when the receiver is about to present an error.
//   - [NSDocument.WillNotPresentError]: Confirms that the error object is not to be presented to the user and the error cannot be recovered from, so cleanup can be done.
//
// # Instance Properties
//
//   - [NSDocument.ObservedPresentedItemUbiquityAttributes]
//   - [NSDocument.PresentedItemURL]
//   - [NSDocument.PreviewRepresentableActivityItems]
//   - [NSDocument.SetPreviewRepresentableActivityItems]
//   - [NSDocument.SavePanelShowsFileFormatsControl]
//
// # Instance Methods
//
//   - [NSDocument.AccommodatePresentedItemDeletionWithCompletionHandler]
//   - [NSDocument.PresentedItemDidChange]
//   - [NSDocument.PresentedItemDidChangeUbiquityAttributes]
//   - [NSDocument.PresentedItemDidGainVersion]
//   - [NSDocument.PresentedItemDidLoseVersion]
//   - [NSDocument.PresentedItemDidMoveToURL]
//   - [NSDocument.PresentedItemDidResolveConflictVersion]
//   - [NSDocument.RelinquishPresentedItemToReader]
//   - [NSDocument.RelinquishPresentedItemToWriter]
//   - [NSDocument.SavePresentedItemChangesWithCompletionHandler]
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument
type NSDocument struct {
	objectivec.Object
}

// NSDocumentFromID constructs a [NSDocument] from an objc.ID.
//
// An abstract class that defines the interface for macOS documents.
func NSDocumentFromID(id objc.ID) NSDocument {
	return NSDocument{objectivec.Object{ID: id}}
}
// NOTE: NSDocument adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSDocument] class.
//
// # Creating a Document Object
//
//   - [INSDocument.InitWithContentsOfURLOfTypeError]: Initializes a document located by a URL of a specified type.
//   - [INSDocument.InitForURLWithContentsOfURLOfTypeError]: Initializes a document with the specified contents, and places the resulting document’s file at the designated location.
//   - [INSDocument.InitWithTypeError]: Initializes a document of a specified type.
//
// # Reading the Document’s Content
//
//   - [INSDocument.ReadFromURLOfTypeError]: Sets the contents of this document by reading from a file or file package, of a specified type, located by a URL.
//   - [INSDocument.ReadFromFileWrapperOfTypeError]: Sets the contents of this document by reading from a file wrapper of a specified type.
//   - [INSDocument.ReadFromDataOfTypeError]: Sets the contents of this document by reading from data of a specified type.
//
// # Writing the Document’s Content
//
//   - [INSDocument.CanAsynchronouslyWriteToURLOfTypeForSaveOperation]: Returns whether the receiver can concurrently write to a file or file package located by a URL, that is formatted for a specific type, for a specific kind of save operation.
//   - [INSDocument.UnblockUserInteraction]: Unblocks the main thread during asynchronous saving.
//   - [INSDocument.WriteToURLOfTypeError]: Writes the contents of the document to a file or file package located by a URL, that is formatted to a specified type.
//   - [INSDocument.WriteSafelyToURLOfTypeForSaveOperationError]: Writes the contents of the document to a file or file package located by a URL.
//   - [INSDocument.FileWrapperOfTypeError]: Creates and returns a file wrapper that contains the contents of the document, formatted to the specified type.
//   - [INSDocument.DataOfTypeError]: Creates and returns a data object that contains the contents of the document, formatted to a specified type.
//   - [INSDocument.WriteToURLOfTypeForSaveOperationOriginalContentsURLError]: Writes the contents of the document to a file or file package located by a URL.
//   - [INSDocument.SaveToURLOfTypeForSaveOperationDelegateDidSaveSelectorContextInfo]: Saves the contents of the document to a file or file package located by a URL, that is formatted to a specified type, for a particular kind of save operation.
//   - [INSDocument.SaveToURLOfTypeForSaveOperationCompletionHandler]: Saves the contents of the document to a file or file package located by a URL, that is formatted to a specified type, for a particular kind of save operation, and invokes the passed-in completion handler.
//   - [INSDocument.FileAttributesToWriteToURLOfTypeForSaveOperationOriginalContentsURLError]: Returns the attributes to write to the file or file package at the specified URL, and targeting the specified type of save operation.
//
// # Getting Document Metadata
//
//   - [INSDocument.FileURL]: The location of the document’s on-disk representation.
//   - [INSDocument.SetFileURL]
//   - [INSDocument.EntireFileLoaded]: A Boolean value that indicates whether the document’s file is completely loaded into memory.
//   - [INSDocument.FileModificationDate]: The last-known modification date of the document’s on-disk representation.
//   - [INSDocument.SetFileModificationDate]
//   - [INSDocument.KeepBackupFile]: A Boolean value that indicates whether the document archives previously saved versions of the document.
//   - [INSDocument.Draft]: A Boolean value that indicates whether the document is a draft that the user has not yet saved.
//   - [INSDocument.SetDraft]
//   - [INSDocument.FileType]: The name of the document type, as specified in the app’s information property-list file.
//   - [INSDocument.SetFileType]
//   - [INSDocument.DocumentEdited]: A Boolean value that indicates whether the document has unsaved changes.
//   - [INSDocument.InViewingMode]: A Boolean value that indicates whether the document is in read-only mode.
//
// # Managing File Type Information
//
//   - [INSDocument.WritableTypesForSaveOperation]: Returns the names of the types to which this document can be saved for a specified kind of save operation.
//   - [INSDocument.FileNameExtensionForTypeSaveOperation]: Returns a filename extension that can be appended to a base filename, for a specified file type and kind of save operation.
//
// # Creating and Managing Window Controllers
//
//   - [INSDocument.MakeWindowControllers]: Creates the window controller objects that the document uses to display its content.
//   - [INSDocument.AddWindowController]: Adds the specified window controller to the current document.
//   - [INSDocument.RemoveWindowController]: Removes the specified window controller from the receiver’s array of window controllers.
//   - [INSDocument.WindowControllers]: The document’s current window controllers.
//   - [INSDocument.WindowNibName]: The name of the document’s sole nib file.
//   - [INSDocument.WindowControllerDidLoadNib]: Called after one of the document’s window controllers loads its nib file.
//   - [INSDocument.WindowControllerWillLoadNib]: Called before one of the document’s window controllers loads its nib file.
//   - [INSDocument.ShouldCloseWindowControllerDelegateShouldCloseSelectorContextInfo]: Determines whether the system should close the document and its associated window.
//
// # Managing Document Windows
//
//   - [INSDocument.ShowWindows]: Displays all of the document’s windows, bringing them to the front and making them main or key as necessary.
//   - [INSDocument.SetWindow]: Sets the window outlet of this document to the specified value.
//   - [INSDocument.WindowForSheet]: Returns the document window to use as the parent of a document-modal sheet.
//   - [INSDocument.DisplayName]: The name of the document as displayed in the title bars of the document’s windows and in alert dialogs related to the document.
//   - [INSDocument.SetDisplayName]
//   - [INSDocument.DefaultDraftName]: Returns the default draft name for the document subclass.
//   - [INSDocument.EncodeRestorableStateWithCoderBackgroundQueue]: Saves the interface-related state of the document.
//
// # Configuring the Autosave Behavior
//
//   - [INSDocument.AutosavedContentsFileURL]: The location of the most recently autosaved document contents.
//   - [INSDocument.SetAutosavedContentsFileURL]
//   - [INSDocument.AutosavingFileType]: The document type to use for an autosave operation.
//   - [INSDocument.AutosavingIsImplicitlyCancellable]: A Boolean value that indicates whether you can cancel an in-progress autosave operation.
//
// # Autosaving the Document
//
//   - [INSDocument.CheckAutosavingSafetyAndReturnError]: Returns a Boolean value that indicates whether it is safe to autosave document changes.
//   - [INSDocument.HasUnautosavedChanges]: A Boolean value that indicates whether the document has changes that have not been autosaved.
//   - [INSDocument.ScheduleAutosaving]: Schedules periodic autosaving for the purpose of crash protection.
//   - [INSDocument.AutosaveDocumentWithDelegateDidAutosaveSelectorContextInfo]: Autosaves the document’s contents to an appropriate location in the file system.
//   - [INSDocument.AutosaveWithImplicitCancellabilityCompletionHandler]: Autosaves the document’s contents to an appropriate file-system location, as needed.
//   - [INSDocument.BackupFileURL]: The URL for the document’s backup file that was created during an autosave operation.
//
// # Browsing Document Versions
//
//   - [INSDocument.BrowseDocumentVersions]: Opens the Versions browser in the document’s main window.
//   - [INSDocument.BrowsingVersions]: A Boolean value that indicates whether the document is currently displaying the Versions browser.
//   - [INSDocument.StopBrowsingVersionsWithCompletionHandler]: Dismiss the Versions browser for the current document.
//
// # Storing Documents in iCloud
//
//   - [INSDocument.MoveDocumentToUbiquityContainer]: Moves the document to the user’s iCloud storage.
//
// # Managing Undo and Redo Actions
//
//   - [INSDocument.UndoManager]: The object that the document uses to support undo/redo operations.
//   - [INSDocument.SetUndoManager]
//   - [INSDocument.HasUndoManager]: A Boolean value that indicates whether the document owns an undo manager object.
//   - [INSDocument.SetHasUndoManager]
//
// # Updating the Document Change Count
//
//   - [INSDocument.UpdateChangeCountWithTokenForSaveOperation]: Updates the document’s change count settings after a successful save operation.
//   - [INSDocument.UpdateChangeCount]: Updates the receiver’s change count according to the given change type.
//   - [INSDocument.ChangeCountTokenForSaveOperation]: Returns an object that encapsulates the current record of document changes at the beginning of a save operation.
//
// # Handling Window Restoration
//
//   - [INSDocument.EncodeRestorableStateWithCoder]: Saves the interface-related state of the document.
//   - [INSDocument.RestoreStateWithCoder]: Restores the interface-related state of the document.
//   - [INSDocument.InvalidateRestorableState]: Marks the document’s interface-related state as dirty.
//   - [INSDocument.RestoreDocumentWindowWithIdentifierStateCompletionHandler]: Restores a window that was associated with a document, after that document is reopened.
//
// # Presenting a Save Panel
//
//   - [INSDocument.RunModalSavePanelForSaveOperationDelegateDidSaveSelectorContextInfo]: Presents a modal Save panel to the user, then tries to save the document if the user approves the operation.
//   - [INSDocument.PrepareSavePanel]: Tells the document to customize the specified Save panel.
//   - [INSDocument.ShouldRunSavePanelWithAccessoryView]: A Boolean value that indicates whether the document’s Save panel displays a list of supported writable document types.
//   - [INSDocument.FileTypeFromLastRunSavePanel]: The file type that was last selected in the Save panel.
//   - [INSDocument.FileNameExtensionWasHiddenInLastRunSavePanel]: A Boolean value that indicates whether the user chose to hide the document’s filename extension.
//
// # Supporting User Activities
//
//   - [INSDocument.UserActivity]: An object that encapsulates a user activity the document supports.
//   - [INSDocument.SetUserActivity]
//   - [INSDocument.UpdateUserActivityState]: Updates the state of the given user activity.
//   - [INSDocument.NSUserActivityDocumentURLKey]: The key that identifies the document associated with a user activity.
//
// # Performing Tasks Serially
//
//   - [INSDocument.PerformSynchronousFileAccessUsingBlock]: Waits for any scheduled file access to complete, then invokes the passed-in block.
//   - [INSDocument.PerformAsynchronousFileAccessUsingBlock]: Waits for any scheduled file access to complete but without blocking the main thread, then invokes the passed-in block.
//   - [INSDocument.PerformActivityWithSynchronousWaitingUsingBlock]: Waits for any work scheduled by previous invocations of this method to complete, then invokes the passed-in block.
//   - [INSDocument.ContinueActivityUsingBlock]: Continues to perform the task for a user activity object using a different block.
//   - [INSDocument.ContinueAsynchronousWorkOnMainThreadUsingBlock]: Invokes the passed-in block on the main thread.
//
// # Handling User Actions
//
//   - [INSDocument.PrintDocument]: Prints the receiver in response to the user choosing the Print menu command.
//   - [INSDocument.RunPageLayout]: The action method invoked in the receiver as first responder when the user chooses the Page Setup menu command.
//   - [INSDocument.RevertDocumentToSaved]: The action of the File menu item Revert in a document-based app.
//   - [INSDocument.SaveDocument]: The action method invoked in the receiver as first responder when the user chooses the Save menu command.
//   - [INSDocument.SaveDocumentAs]: The action method invoked in the receiver as first responder when the user chooses the Save As menu command.
//   - [INSDocument.SaveDocumentTo]: The action method invoked in the receiver as first responder when the user chooses the Save To menu command.
//   - [INSDocument.SaveDocumentWithDelegateDidSaveSelectorContextInfo]: Saves the document and delivers the results to the provided delegate object.
//
// # Closing the Document
//
//   - [INSDocument.CanCloseDocumentWithDelegateShouldCloseSelectorContextInfo]: Determines whether to close the document, prompting the user as needed to choose a course of action.
//   - [INSDocument.Close]: Closes all of the document’s windows and removes the document from its document controller.
//
// # Reverting the Document Contents
//
//   - [INSDocument.RevertToContentsOfURLOfTypeError]: Discards all unsaved document modifications and replaces the document’s contents by reading a file or file package located by a URL of a specified type.
//
// # Duplicating the Document
//
//   - [INSDocument.DuplicateAndReturnError]: Creates a new document whose contents are the same as the receiver and returns an error object if unsuccessful.
//   - [INSDocument.DuplicateDocument]: Creates a copy of the receiving document in response to the user choosing Duplicate from the File menu.
//   - [INSDocument.DuplicateDocumentWithDelegateDidDuplicateSelectorContextInfo]: Creates a new document whose contents are the same as the current document.
//
// # Renaming the Document
//
//   - [INSDocument.RenameDocument]: Renames the current document in response to the user choosing the Rename menu item.
//
// # Moving the Document
//
//   - [INSDocument.MoveDocument]: Moves the document to a new location in response to the user choosing the Move To… menu item.
//   - [INSDocument.MoveDocumentWithCompletionHandler]: Moves the document to a user-selected location.
//   - [INSDocument.MoveToURLCompletionHandler]: Moves the document’s file to the given URL.
//
// # Locking the Document
//
//   - [INSDocument.LockDocument]: Locks the document in response to the user choosing the Lock menu item.
//   - [INSDocument.UnlockDocument]: Unlocks the document in response to the user choosing the Unlock menu item.
//   - [INSDocument.LockDocumentWithCompletionHandler]: Prevents the user from making further changes to the document.
//   - [INSDocument.LockWithCompletionHandler]: Prevents the user from making changes to the document’s file.
//   - [INSDocument.UnlockDocumentWithCompletionHandler]: Allows the user to make modifications to the document.
//   - [INSDocument.UnlockWithCompletionHandler]: Allows the user to make modifications to the document’s file.
//   - [INSDocument.Locked]: A Boolean value that indicates whether or not the file can be written to.
//
// # Printing the Document
//
//   - [INSDocument.PrintInfo]: The printing information associated with the document.
//   - [INSDocument.SetPrintInfo]
//   - [INSDocument.PreparePageLayout]: Adds document-specific content to the Page Layout panel.
//   - [INSDocument.RunModalPageLayoutWithPrintInfoDelegateDidRunSelectorContextInfo]: Runs the modal page layout panel with the receiver’s printing information object.
//   - [INSDocument.RunModalPrintOperationDelegateDidRunSelectorContextInfo]: Runs the specified print operation modally.
//   - [INSDocument.ShouldChangePrintInfo]: Returns a Boolean value that indicates whether the document allows changes to the default printing information.
//   - [INSDocument.PrintDocumentWithSettingsShowPrintPanelDelegateDidPrintSelectorContextInfo]: Prints the document’s contents, optionally displaying a print panel to the user.
//   - [INSDocument.PrintOperationWithSettingsError]: Creates and returns a print operation for the document’s contents.
//   - [INSDocument.PDFPrintOperation]: A print operation you can use to create a PDF representation of the document’s current contents.
//   - [INSDocument.SaveDocumentToPDF]: Exports a PDF representation of the document’s current contents.
//
// # Sharing the Document
//
//   - [INSDocument.AllowsDocumentSharing]: A Boolean value that indicates whether the document is shareable from the standard Share menu.
//   - [INSDocument.PrepareSharingServicePicker]: Perform any custom setup associated with a sharing service picker.
//   - [INSDocument.ShareDocumentWithSharingServiceCompletionHandler]: Share the document’s file using the specified sharing service.
//
// # Handling Script Commands
//
//   - [INSDocument.HandleCloseScriptCommand]: Handles the Close AppleScript command by attempting to close the document.
//   - [INSDocument.HandlePrintScriptCommand]: Handles the Print AppleScript command by attempting to print the document.
//   - [INSDocument.HandleSaveScriptCommand]: Handles the Save AppleScript command by attempting to save the document.
//   - [INSDocument.ObjectSpecifier]: Returns the object specifier that represents the document.
//   - [INSDocument.LastComponentOfFileName]: The name of the document seen by the user in AppleScript.
//   - [INSDocument.SetLastComponentOfFileName]
//
// # Displaying Errors to the User
//
//   - [INSDocument.PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo]: Presents an error alert to the user as a modal panel.
//   - [INSDocument.PresentError]: Presents an error alert to the user as a modal panel.
//   - [INSDocument.WillPresentError]: Called when the receiver is about to present an error.
//   - [INSDocument.WillNotPresentError]: Confirms that the error object is not to be presented to the user and the error cannot be recovered from, so cleanup can be done.
//
// # Instance Properties
//
//   - [INSDocument.ObservedPresentedItemUbiquityAttributes]
//   - [INSDocument.PresentedItemURL]
//   - [INSDocument.PreviewRepresentableActivityItems]
//   - [INSDocument.SetPreviewRepresentableActivityItems]
//   - [INSDocument.SavePanelShowsFileFormatsControl]
//
// # Instance Methods
//
//   - [INSDocument.AccommodatePresentedItemDeletionWithCompletionHandler]
//   - [INSDocument.PresentedItemDidChange]
//   - [INSDocument.PresentedItemDidChangeUbiquityAttributes]
//   - [INSDocument.PresentedItemDidGainVersion]
//   - [INSDocument.PresentedItemDidLoseVersion]
//   - [INSDocument.PresentedItemDidMoveToURL]
//   - [INSDocument.PresentedItemDidResolveConflictVersion]
//   - [INSDocument.RelinquishPresentedItemToReader]
//   - [INSDocument.RelinquishPresentedItemToWriter]
//   - [INSDocument.SavePresentedItemChangesWithCompletionHandler]
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument
type INSDocument interface {
	objectivec.IObject
	NSEditorRegistration
	NSMenuItemValidation
	NSUserActivityRestoring
	NSUserInterfaceValidations

	// Topic: Creating a Document Object

	// Initializes a document located by a URL of a specified type.
	InitWithContentsOfURLOfTypeError(url foundation.INSURL, typeName string) (NSDocument, error)
	// Initializes a document with the specified contents, and places the resulting document’s file at the designated location.
	InitForURLWithContentsOfURLOfTypeError(urlOrNil foundation.INSURL, contentsURL foundation.INSURL, typeName string) (NSDocument, error)
	// Initializes a document of a specified type.
	InitWithTypeError(typeName string) (NSDocument, error)

	// Topic: Reading the Document’s Content

	// Sets the contents of this document by reading from a file or file package, of a specified type, located by a URL.
	ReadFromURLOfTypeError(url foundation.INSURL, typeName string) (bool, error)
	// Sets the contents of this document by reading from a file wrapper of a specified type.
	ReadFromFileWrapperOfTypeError(fileWrapper foundation.NSFileWrapper, typeName string) (bool, error)
	// Sets the contents of this document by reading from data of a specified type.
	ReadFromDataOfTypeError(data foundation.INSData, typeName string) (bool, error)

	// Topic: Writing the Document’s Content

	// Returns whether the receiver can concurrently write to a file or file package located by a URL, that is formatted for a specific type, for a specific kind of save operation.
	CanAsynchronouslyWriteToURLOfTypeForSaveOperation(url foundation.INSURL, typeName string, saveOperation NSSaveOperationType) bool
	// Unblocks the main thread during asynchronous saving.
	UnblockUserInteraction()
	// Writes the contents of the document to a file or file package located by a URL, that is formatted to a specified type.
	WriteToURLOfTypeError(url foundation.INSURL, typeName string) (bool, error)
	// Writes the contents of the document to a file or file package located by a URL.
	WriteSafelyToURLOfTypeForSaveOperationError(url foundation.INSURL, typeName string, saveOperation NSSaveOperationType) (bool, error)
	// Creates and returns a file wrapper that contains the contents of the document, formatted to the specified type.
	FileWrapperOfTypeError(typeName string) (foundation.NSFileWrapper, error)
	// Creates and returns a data object that contains the contents of the document, formatted to a specified type.
	DataOfTypeError(typeName string) (foundation.INSData, error)
	// Writes the contents of the document to a file or file package located by a URL.
	WriteToURLOfTypeForSaveOperationOriginalContentsURLError(url foundation.INSURL, typeName string, saveOperation NSSaveOperationType, absoluteOriginalContentsURL foundation.INSURL) (bool, error)
	// Saves the contents of the document to a file or file package located by a URL, that is formatted to a specified type, for a particular kind of save operation.
	SaveToURLOfTypeForSaveOperationDelegateDidSaveSelectorContextInfo(url foundation.INSURL, typeName string, saveOperation NSSaveOperationType, delegate objectivec.IObject, didSaveSelector objc.SEL, contextInfo uintptr)
	// Saves the contents of the document to a file or file package located by a URL, that is formatted to a specified type, for a particular kind of save operation, and invokes the passed-in completion handler.
	SaveToURLOfTypeForSaveOperationCompletionHandler(url foundation.INSURL, typeName string, saveOperation NSSaveOperationType, completionHandler ErrorHandler)
	// Returns the attributes to write to the file or file package at the specified URL, and targeting the specified type of save operation.
	FileAttributesToWriteToURLOfTypeForSaveOperationOriginalContentsURLError(url foundation.INSURL, typeName string, saveOperation NSSaveOperationType, absoluteOriginalContentsURL foundation.INSURL) (foundation.INSDictionary, error)

	// Topic: Getting Document Metadata

	// The location of the document’s on-disk representation.
	FileURL() foundation.INSURL
	SetFileURL(value foundation.INSURL)
	// A Boolean value that indicates whether the document’s file is completely loaded into memory.
	EntireFileLoaded() bool
	// The last-known modification date of the document’s on-disk representation.
	FileModificationDate() foundation.INSDate
	SetFileModificationDate(value foundation.INSDate)
	// A Boolean value that indicates whether the document archives previously saved versions of the document.
	KeepBackupFile() bool
	// A Boolean value that indicates whether the document is a draft that the user has not yet saved.
	Draft() bool
	SetDraft(value bool)
	// The name of the document type, as specified in the app’s information property-list file.
	FileType() string
	SetFileType(value string)
	// A Boolean value that indicates whether the document has unsaved changes.
	DocumentEdited() bool
	// A Boolean value that indicates whether the document is in read-only mode.
	InViewingMode() bool

	// Topic: Managing File Type Information

	// Returns the names of the types to which this document can be saved for a specified kind of save operation.
	WritableTypesForSaveOperation(saveOperation NSSaveOperationType) []string
	// Returns a filename extension that can be appended to a base filename, for a specified file type and kind of save operation.
	FileNameExtensionForTypeSaveOperation(typeName string, saveOperation NSSaveOperationType) string

	// Topic: Creating and Managing Window Controllers

	// Creates the window controller objects that the document uses to display its content.
	MakeWindowControllers()
	// Adds the specified window controller to the current document.
	AddWindowController(windowController INSWindowController)
	// Removes the specified window controller from the receiver’s array of window controllers.
	RemoveWindowController(windowController INSWindowController)
	// The document’s current window controllers.
	WindowControllers() []NSWindowController
	// The name of the document’s sole nib file.
	WindowNibName() NSNibName
	// Called after one of the document’s window controllers loads its nib file.
	WindowControllerDidLoadNib(windowController INSWindowController)
	// Called before one of the document’s window controllers loads its nib file.
	WindowControllerWillLoadNib(windowController INSWindowController)
	// Determines whether the system should close the document and its associated window.
	ShouldCloseWindowControllerDelegateShouldCloseSelectorContextInfo(windowController INSWindowController, delegate objectivec.IObject, shouldCloseSelector objc.SEL, contextInfo uintptr)

	// Topic: Managing Document Windows

	// Displays all of the document’s windows, bringing them to the front and making them main or key as necessary.
	ShowWindows()
	// Sets the window outlet of this document to the specified value.
	SetWindow(window INSWindow)
	// Returns the document window to use as the parent of a document-modal sheet.
	WindowForSheet() INSWindow
	// The name of the document as displayed in the title bars of the document’s windows and in alert dialogs related to the document.
	DisplayName() string
	SetDisplayName(value string)
	// Returns the default draft name for the document subclass.
	DefaultDraftName() string
	// Saves the interface-related state of the document.
	EncodeRestorableStateWithCoderBackgroundQueue(coder foundation.INSCoder, queue foundation.NSOperationQueue)

	// Topic: Configuring the Autosave Behavior

	// The location of the most recently autosaved document contents.
	AutosavedContentsFileURL() foundation.INSURL
	SetAutosavedContentsFileURL(value foundation.INSURL)
	// The document type to use for an autosave operation.
	AutosavingFileType() string
	// A Boolean value that indicates whether you can cancel an in-progress autosave operation.
	AutosavingIsImplicitlyCancellable() bool

	// Topic: Autosaving the Document

	// Returns a Boolean value that indicates whether it is safe to autosave document changes.
	CheckAutosavingSafetyAndReturnError() (bool, error)
	// A Boolean value that indicates whether the document has changes that have not been autosaved.
	HasUnautosavedChanges() bool
	// Schedules periodic autosaving for the purpose of crash protection.
	ScheduleAutosaving()
	// Autosaves the document’s contents to an appropriate location in the file system.
	AutosaveDocumentWithDelegateDidAutosaveSelectorContextInfo(delegate objectivec.IObject, didAutosaveSelector objc.SEL, contextInfo uintptr)
	// Autosaves the document’s contents to an appropriate file-system location, as needed.
	AutosaveWithImplicitCancellabilityCompletionHandler(autosavingIsImplicitlyCancellable bool, completionHandler ErrorHandler)
	// The URL for the document’s backup file that was created during an autosave operation.
	BackupFileURL() foundation.INSURL

	// Topic: Browsing Document Versions

	// Opens the Versions browser in the document’s main window.
	BrowseDocumentVersions(sender objectivec.IObject)
	// A Boolean value that indicates whether the document is currently displaying the Versions browser.
	BrowsingVersions() bool
	// Dismiss the Versions browser for the current document.
	StopBrowsingVersionsWithCompletionHandler(completionHandler VoidHandler)

	// Topic: Storing Documents in iCloud

	// Moves the document to the user’s iCloud storage.
	MoveDocumentToUbiquityContainer(sender objectivec.IObject)

	// Topic: Managing Undo and Redo Actions

	// The object that the document uses to support undo/redo operations.
	UndoManager() foundation.NSUndoManager
	SetUndoManager(value foundation.NSUndoManager)
	// A Boolean value that indicates whether the document owns an undo manager object.
	HasUndoManager() bool
	SetHasUndoManager(value bool)

	// Topic: Updating the Document Change Count

	// Updates the document’s change count settings after a successful save operation.
	UpdateChangeCountWithTokenForSaveOperation(changeCountToken objectivec.IObject, saveOperation NSSaveOperationType)
	// Updates the receiver’s change count according to the given change type.
	UpdateChangeCount(change NSDocumentChangeType)
	// Returns an object that encapsulates the current record of document changes at the beginning of a save operation.
	ChangeCountTokenForSaveOperation(saveOperation NSSaveOperationType) objectivec.IObject

	// Topic: Handling Window Restoration

	// Saves the interface-related state of the document.
	EncodeRestorableStateWithCoder(coder foundation.INSCoder)
	// Restores the interface-related state of the document.
	RestoreStateWithCoder(coder foundation.INSCoder)
	// Marks the document’s interface-related state as dirty.
	InvalidateRestorableState()
	// Restores a window that was associated with a document, after that document is reopened.
	RestoreDocumentWindowWithIdentifierStateCompletionHandler(identifier NSUserInterfaceItemIdentifier, state foundation.INSCoder, completionHandler WindowErrorHandler)

	// Topic: Presenting a Save Panel

	// Presents a modal Save panel to the user, then tries to save the document if the user approves the operation.
	RunModalSavePanelForSaveOperationDelegateDidSaveSelectorContextInfo(saveOperation NSSaveOperationType, delegate objectivec.IObject, didSaveSelector objc.SEL, contextInfo uintptr)
	// Tells the document to customize the specified Save panel.
	PrepareSavePanel(savePanel INSSavePanel) bool
	// A Boolean value that indicates whether the document’s Save panel displays a list of supported writable document types.
	ShouldRunSavePanelWithAccessoryView() bool
	// The file type that was last selected in the Save panel.
	FileTypeFromLastRunSavePanel() string
	// A Boolean value that indicates whether the user chose to hide the document’s filename extension.
	FileNameExtensionWasHiddenInLastRunSavePanel() bool

	// Topic: Supporting User Activities

	// An object that encapsulates a user activity the document supports.
	UserActivity() foundation.NSUserActivity
	SetUserActivity(value foundation.NSUserActivity)
	// Updates the state of the given user activity.
	UpdateUserActivityState(activity foundation.NSUserActivity)
	// The key that identifies the document associated with a user activity.
	NSUserActivityDocumentURLKey() string

	// Topic: Performing Tasks Serially

	// Waits for any scheduled file access to complete, then invokes the passed-in block.
	PerformSynchronousFileAccessUsingBlock(block VoidHandler)
	// Waits for any scheduled file access to complete but without blocking the main thread, then invokes the passed-in block.
	PerformAsynchronousFileAccessUsingBlock(block VoidHandler)
	// Waits for any work scheduled by previous invocations of this method to complete, then invokes the passed-in block.
	PerformActivityWithSynchronousWaitingUsingBlock(waitSynchronously bool, block VoidHandler)
	// Continues to perform the task for a user activity object using a different block.
	ContinueActivityUsingBlock(block VoidHandler)
	// Invokes the passed-in block on the main thread.
	ContinueAsynchronousWorkOnMainThreadUsingBlock(block VoidHandler)

	// Topic: Handling User Actions

	// Prints the receiver in response to the user choosing the Print menu command.
	PrintDocument(sender objectivec.IObject)
	// The action method invoked in the receiver as first responder when the user chooses the Page Setup menu command.
	RunPageLayout(sender objectivec.IObject)
	// The action of the File menu item Revert in a document-based app.
	RevertDocumentToSaved(sender objectivec.IObject)
	// The action method invoked in the receiver as first responder when the user chooses the Save menu command.
	SaveDocument(sender objectivec.IObject)
	// The action method invoked in the receiver as first responder when the user chooses the Save As menu command.
	SaveDocumentAs(sender objectivec.IObject)
	// The action method invoked in the receiver as first responder when the user chooses the Save To menu command.
	SaveDocumentTo(sender objectivec.IObject)
	// Saves the document and delivers the results to the provided delegate object.
	SaveDocumentWithDelegateDidSaveSelectorContextInfo(delegate objectivec.IObject, didSaveSelector objc.SEL, contextInfo uintptr)

	// Topic: Closing the Document

	// Determines whether to close the document, prompting the user as needed to choose a course of action.
	CanCloseDocumentWithDelegateShouldCloseSelectorContextInfo(delegate objectivec.IObject, shouldCloseSelector objc.SEL, contextInfo uintptr)
	// Closes all of the document’s windows and removes the document from its document controller.
	Close()

	// Topic: Reverting the Document Contents

	// Discards all unsaved document modifications and replaces the document’s contents by reading a file or file package located by a URL of a specified type.
	RevertToContentsOfURLOfTypeError(url foundation.INSURL, typeName string) (bool, error)

	// Topic: Duplicating the Document

	// Creates a new document whose contents are the same as the receiver and returns an error object if unsuccessful.
	DuplicateAndReturnError() (INSDocument, error)
	// Creates a copy of the receiving document in response to the user choosing Duplicate from the File menu.
	DuplicateDocument(sender objectivec.IObject)
	// Creates a new document whose contents are the same as the current document.
	DuplicateDocumentWithDelegateDidDuplicateSelectorContextInfo(delegate objectivec.IObject, didDuplicateSelector objc.SEL, contextInfo uintptr)

	// Topic: Renaming the Document

	// Renames the current document in response to the user choosing the Rename menu item.
	RenameDocument(sender objectivec.IObject)

	// Topic: Moving the Document

	// Moves the document to a new location in response to the user choosing the Move To… menu item.
	MoveDocument(sender objectivec.IObject)
	// Moves the document to a user-selected location.
	MoveDocumentWithCompletionHandler(completionHandler BoolHandler)
	// Moves the document’s file to the given URL.
	MoveToURLCompletionHandler(url foundation.INSURL, completionHandler ErrorHandler)

	// Topic: Locking the Document

	// Locks the document in response to the user choosing the Lock menu item.
	LockDocument(sender objectivec.IObject)
	// Unlocks the document in response to the user choosing the Unlock menu item.
	UnlockDocument(sender objectivec.IObject)
	// Prevents the user from making further changes to the document.
	LockDocumentWithCompletionHandler(completionHandler BoolHandler)
	// Prevents the user from making changes to the document’s file.
	LockWithCompletionHandler(completionHandler ErrorHandler)
	// Allows the user to make modifications to the document.
	UnlockDocumentWithCompletionHandler(completionHandler BoolHandler)
	// Allows the user to make modifications to the document’s file.
	UnlockWithCompletionHandler(completionHandler ErrorHandler)
	// A Boolean value that indicates whether or not the file can be written to.
	Locked() bool

	// Topic: Printing the Document

	// The printing information associated with the document.
	PrintInfo() INSPrintInfo
	SetPrintInfo(value INSPrintInfo)
	// Adds document-specific content to the Page Layout panel.
	PreparePageLayout(pageLayout INSPageLayout) bool
	// Runs the modal page layout panel with the receiver’s printing information object.
	RunModalPageLayoutWithPrintInfoDelegateDidRunSelectorContextInfo(printInfo INSPrintInfo, delegate objectivec.IObject, didRunSelector objc.SEL, contextInfo uintptr)
	// Runs the specified print operation modally.
	RunModalPrintOperationDelegateDidRunSelectorContextInfo(printOperation INSPrintOperation, delegate objectivec.IObject, didRunSelector objc.SEL, contextInfo uintptr)
	// Returns a Boolean value that indicates whether the document allows changes to the default printing information.
	ShouldChangePrintInfo(newPrintInfo INSPrintInfo) bool
	// Prints the document’s contents, optionally displaying a print panel to the user.
	PrintDocumentWithSettingsShowPrintPanelDelegateDidPrintSelectorContextInfo(printSettings foundation.INSDictionary, showPrintPanel bool, delegate objectivec.IObject, didPrintSelector objc.SEL, contextInfo uintptr)
	// Creates and returns a print operation for the document’s contents.
	PrintOperationWithSettingsError(printSettings foundation.INSDictionary) (INSPrintOperation, error)
	// A print operation you can use to create a PDF representation of the document’s current contents.
	PDFPrintOperation() INSPrintOperation
	// Exports a PDF representation of the document’s current contents.
	SaveDocumentToPDF(sender objectivec.IObject)

	// Topic: Sharing the Document

	// A Boolean value that indicates whether the document is shareable from the standard Share menu.
	AllowsDocumentSharing() bool
	// Perform any custom setup associated with a sharing service picker.
	PrepareSharingServicePicker(sharingServicePicker INSSharingServicePicker)
	// Share the document’s file using the specified sharing service.
	ShareDocumentWithSharingServiceCompletionHandler(sharingService INSSharingService, completionHandler BoolHandler)

	// Topic: Handling Script Commands

	// Handles the Close AppleScript command by attempting to close the document.
	HandleCloseScriptCommand(command foundation.NSCloseCommand) objectivec.IObject
	// Handles the Print AppleScript command by attempting to print the document.
	HandlePrintScriptCommand(command foundation.NSScriptCommand) objectivec.IObject
	// Handles the Save AppleScript command by attempting to save the document.
	HandleSaveScriptCommand(command foundation.NSScriptCommand) objectivec.IObject
	// Returns the object specifier that represents the document.
	ObjectSpecifier() foundation.NSScriptObjectSpecifier
	// The name of the document seen by the user in AppleScript.
	LastComponentOfFileName() string
	SetLastComponentOfFileName(value string)

	// Topic: Displaying Errors to the User

	// Presents an error alert to the user as a modal panel.
	PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error_ foundation.INSError, window INSWindow, delegate objectivec.IObject, didPresentSelector objc.SEL, contextInfo uintptr)
	// Presents an error alert to the user as a modal panel.
	PresentError(error_ foundation.INSError) bool
	// Called when the receiver is about to present an error.
	WillPresentError(error_ foundation.INSError) foundation.INSError
	// Confirms that the error object is not to be presented to the user and the error cannot be recovered from, so cleanup can be done.
	WillNotPresentError(error_ foundation.INSError)

	// Topic: Instance Properties

	ObservedPresentedItemUbiquityAttributes() foundation.INSSet
	PresentedItemURL() foundation.INSURL
	PreviewRepresentableActivityItems() []objectivec.IObject
	SetPreviewRepresentableActivityItems(value []objectivec.IObject)
	SavePanelShowsFileFormatsControl() bool

	// Topic: Instance Methods

	AccommodatePresentedItemDeletionWithCompletionHandler(completionHandler ErrorHandler)
	PresentedItemDidChange()
	PresentedItemDidChangeUbiquityAttributes(attributes foundation.INSSet)
	PresentedItemDidGainVersion(version foundation.NSFileVersion)
	PresentedItemDidLoseVersion(version foundation.NSFileVersion)
	PresentedItemDidMoveToURL(newURL foundation.INSURL)
	PresentedItemDidResolveConflictVersion(version foundation.NSFileVersion)
	RelinquishPresentedItemToReader(reader VoidHandler)
	RelinquishPresentedItemToWriter(writer VoidHandler)
	SavePresentedItemChangesWithCompletionHandler(completionHandler ErrorHandler)
}

// Init initializes the instance.
func (d NSDocument) Init() NSDocument {
	rv := objc.Send[NSDocument](d.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (d NSDocument) Autorelease() NSDocument {
	rv := objc.Send[NSDocument](d.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSDocument creates a new NSDocument instance.
func NewNSDocument() NSDocument {
	class := getNSDocumentClass()
	rv := objc.Send[NSDocument](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a document with the specified contents, and places the
// resulting document’s file at the designated location.
//
// urlOrNil: The location for the document’s file. This value is `nil` for an
// autosaved document that the user never explicitly saved.
//
// contentsURL: The URL of the file that contains the document’s contents. When loading
// an autosaved document, this URL contains the location of the autosave file.
// The contents of this file replace the contents of the file in
// `absoluteDocumentURL`.
//
// typeName: The string that identifies the document type.
//
// # Return Value
// 
// The initialized document object, or `nil` if the document could not be
// created.
//
// # Discussion
// 
// The system calls this method to open a document that has an associated
// autosave file . You can override this method to handle any document
// initialization specific to autosave contents.
// 
// After reading the contents from the specified autosave file, this method
// updates the document’s change count using the [NSChangeReadOtherContents]
// change type.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/init(for:withContentsOf:ofType:)
func NewDocumentForURLWithContentsOfURLOfTypeError(urlOrNil foundation.INSURL, contentsURL foundation.INSURL, typeName string) (NSDocument, error) {
	var errorPtr objc.ID
	instance := getNSDocumentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initForURL:withContentsOfURL:ofType:error:"), urlOrNil, contentsURL, objc.String(typeName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSDocument{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSDocumentFromID(rv), nil
}

// Initializes a document located by a URL of a specified type.
//
// url: The URL from which the contents of the document are obtained.
//
// typeName: The string that identifies the document type.
//
// # Return Value
// 
// The initialized [NSDocument] object, or, if the document could not be
// created, `nil`.
//
// # Discussion
// 
// You can override this method to customize the reopening of autosaved
// documents.
// 
// This method is invoked by the [NSDocumentController] method
// [DocumentWithContentsOfURLOfTypeError]. The default implementation of this
// method calls the [Init] and [ReadFromURLOfTypeError] methods and sets
// values for the [FileURL], [FileType], and [FileModificationDate]
// properties.
// 
// For backward binary compatibility with OS X v10.3 and earlier, the default
// implementation of this method instead invokes
// [initWithContentsOfFile:ofType:] if it is overridden and the URL uses the
// `` scheme. It still updates the [FileModificationDate] property in this
// situation.
//
// [initWithContentsOfFile:ofType:]: https://developer.apple.com/documentation/AppKit/NSDocument/initWithContentsOfFile:ofType:
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/init(contentsOf:ofType:)
func NewDocumentWithContentsOfURLOfTypeError(url foundation.INSURL, typeName string) (NSDocument, error) {
	var errorPtr objc.ID
	instance := getNSDocumentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithContentsOfURL:ofType:error:"), url, objc.String(typeName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSDocument{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSDocumentFromID(rv), nil
}

// Initializes a document of a specified type.
//
// typeName: The string that identifies the document type.
//
// # Return Value
// 
// The initialized [NSDocument] object, or, if the document could not be
// created, `nil`.
//
// # Discussion
// 
// The default implementation of this method just invokes `[self init]` and
// `[self typeName]`.
// 
// You can override this method to perform initialization that must be done
// when creating new documents but should not be done when opening existing
// documents. Your override should typically invoke `super`, or at least it
// must invoke [Init], the [NSDocument] designated initializer, to initialize
// the [NSDocument] private instance variables.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/init(type:)
func NewDocumentWithTypeError(typeName string) (NSDocument, error) {
	var errorPtr objc.ID
	instance := getNSDocumentClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithType:error:"), objc.String(typeName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSDocument{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSDocumentFromID(rv), nil
}

// Initializes a document located by a URL of a specified type.
//
// url: The URL from which the contents of the document are obtained.
//
// typeName: The string that identifies the document type.
//
// # Return Value
// 
// The initialized [NSDocument] object, or, if the document could not be
// created, `nil`.
//
// # Discussion
// 
// You can override this method to customize the reopening of autosaved
// documents.
// 
// This method is invoked by the [NSDocumentController] method
// [DocumentWithContentsOfURLOfTypeError]. The default implementation of this
// method calls the [Init] and [ReadFromURLOfTypeError] methods and sets
// values for the [FileURL], [FileType], and [FileModificationDate]
// properties.
// 
// For backward binary compatibility with OS X v10.3 and earlier, the default
// implementation of this method instead invokes
// [initWithContentsOfFile:ofType:] if it is overridden and the URL uses the
// `` scheme. It still updates the [FileModificationDate] property in this
// situation.
//
// [initWithContentsOfFile:ofType:]: https://developer.apple.com/documentation/AppKit/NSDocument/initWithContentsOfFile:ofType:
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/init(contentsOf:ofType:)
func (d NSDocument) InitWithContentsOfURLOfTypeError(url foundation.INSURL, typeName string) (NSDocument, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithContentsOfURL:ofType:error:"), url, objc.String(typeName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSDocument{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSDocumentFromID(rv), nil

}
// Initializes a document with the specified contents, and places the
// resulting document’s file at the designated location.
//
// urlOrNil: The location for the document’s file. This value is `nil` for an
// autosaved document that the user never explicitly saved.
//
// contentsURL: The URL of the file that contains the document’s contents. When loading
// an autosaved document, this URL contains the location of the autosave file.
// The contents of this file replace the contents of the file in
// `absoluteDocumentURL`.
//
// typeName: The string that identifies the document type.
//
// # Return Value
// 
// The initialized document object, or `nil` if the document could not be
// created.
//
// # Discussion
// 
// The system calls this method to open a document that has an associated
// autosave file . You can override this method to handle any document
// initialization specific to autosave contents.
// 
// After reading the contents from the specified autosave file, this method
// updates the document’s change count using the [NSChangeReadOtherContents]
// change type.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/init(for:withContentsOf:ofType:)
func (d NSDocument) InitForURLWithContentsOfURLOfTypeError(urlOrNil foundation.INSURL, contentsURL foundation.INSURL, typeName string) (NSDocument, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initForURL:withContentsOfURL:ofType:error:"), urlOrNil, contentsURL, objc.String(typeName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSDocument{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSDocumentFromID(rv), nil

}
// Initializes a document of a specified type.
//
// typeName: The string that identifies the document type.
//
// # Return Value
// 
// The initialized [NSDocument] object, or, if the document could not be
// created, `nil`.
//
// # Discussion
// 
// The default implementation of this method just invokes `[self init]` and
// `[self typeName]`.
// 
// You can override this method to perform initialization that must be done
// when creating new documents but should not be done when opening existing
// documents. Your override should typically invoke `super`, or at least it
// must invoke [Init], the [NSDocument] designated initializer, to initialize
// the [NSDocument] private instance variables.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/init(type:)
func (d NSDocument) InitWithTypeError(typeName string) (NSDocument, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("initWithType:error:"), objc.String(typeName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSDocument{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSDocumentFromID(rv), nil

}
// Sets the contents of this document by reading from a file or file package,
// of a specified type, located by a URL.
//
// url: The location from which the document contents are read.
//
// typeName: The string that identifies the document type.
//
// # Discussion
// 
// The default implementation of this method just creates an NSFileWrapper and
// invokes `[self theFileWrapper typeName outError]`.
// 
// For backward binary compatibility with OS X v10.3 and earlier, the default
// implementation of this method instead invokes `[self [absoluteURL path]
// typeName]` if [readFromFile:ofType:] is overridden and the URL uses the ``
// scheme.
//
// [readFromFile:ofType:]: https://developer.apple.com/documentation/AppKit/NSDocument/readFromFile:ofType:
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/read(from:ofType:)-1vttv
func (d NSDocument) ReadFromURLOfTypeError(url foundation.INSURL, typeName string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("readFromURL:ofType:error:"), url, objc.String(typeName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("readFromURL:ofType:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Sets the contents of this document by reading from a file wrapper of a
// specified type.
//
// fileWrapper: The file wrapper from which the document contents are read.
//
// typeName: The string that identifies the document type.
//
// # Discussion
// 
// The default implementation of this method invokes `[self [fileWrapper
// regularFileContents] typeName outError]`.
// 
// For backward binary compatibility with OS X v10.3 and earlier, the default
// implementation of this method instead invokes `[self fileWrapper typeName]`
// if [loadFileWrapperRepresentation:ofType:] is overridden.
//
// [loadFileWrapperRepresentation:ofType:]: https://developer.apple.com/documentation/AppKit/NSDocument/loadFileWrapperRepresentation:ofType:
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/read(from:ofType:)-3rzsi
func (d NSDocument) ReadFromFileWrapperOfTypeError(fileWrapper foundation.NSFileWrapper, typeName string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("readFromFileWrapper:ofType:error:"), fileWrapper, objc.String(typeName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("readFromFileWrapper:ofType:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Sets the contents of this document by reading from data of a specified
// type.
//
// data: The data object from which the document contents are read.
//
// typeName: The string that identifies the document type.
//
// # Discussion
// 
// The default implementation of this method throws an exception because at
// least one of the three reading methods (this method,
// [ReadFromURLOfTypeError], [ReadFromFileWrapperOfTypeError]), or every
// method that may invoke [ReadFromURLOfTypeError], must be overridden.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/read(from:ofType:)-6g6ai
func (d NSDocument) ReadFromDataOfTypeError(data foundation.INSData, typeName string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("readFromData:ofType:error:"), data, objc.String(typeName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("readFromData:ofType:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Returns whether the receiver can concurrently write to a file or file
// package located by a URL, that is formatted for a specific type, for a
// specific kind of save operation.
//
// url: The location of the file or package to which the document is written.
//
// typeName: The string that identifies the document type.
//
// saveOperation: The type of save operation.
//
// # Return Value
// 
// [false] by default; subclasses can override to return [true], thereby
// enabling asynchronous writing.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The default implementation of this method returns [false]. You are strongly
// encouraged to override it and make it return [true], after making sure your
// overrides of document writing methods can be safely invoked on a non-main
// thread, and making sure that the [UnblockUserInteraction] method is invoked
// at some appropriate time during writing.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/canAsynchronouslyWrite(to:ofType:for:)
func (d NSDocument) CanAsynchronouslyWriteToURLOfTypeForSaveOperation(url foundation.INSURL, typeName string, saveOperation NSSaveOperationType) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("canAsynchronouslyWriteToURL:ofType:forSaveOperation:"), url, objc.String(typeName), saveOperation)
	return rv
}
// Unblocks the main thread during asynchronous saving.
//
// # Discussion
// 
// If [SaveToURLOfTypeForSaveOperationCompletionHandler] is writing on a
// non-main thread because [CanAsynchronouslyWriteToURLOfTypeForSaveOperation]
// has returned [true], but it is still blocking the main thread, this method
// unblocks the main thread. Otherwise, it does nothing. For example, the
// default implementation of [FileWrapperOfTypeError] invokes this when it has
// created the [FileWrapper] object to return. Assuming that the
// [NSFileWrapper] is not mutated by subsequent user actions, it is
// effectively a “snapshot” of the document’s contents, and once it is
// created it is safe to resume handling user events on the main thread, even
// though some of those user events might change the document’s contents
// before the [NSFileWrapper] object has been safely written. You can invoke
// this method to make asynchronous saving actually asynchronous if you’ve
// overridden [WriteSafelyToURLOfTypeForSaveOperationError],
// [WriteToURLOfTypeForSaveOperationOriginalContentsURLError], or
// [WriteToURLOfTypeError] in such a way that the invocation of this method
// done by the [WriteToURLOfTypeError] default implementation won’t happen
// during writing.
//
// [FileWrapper]: https://developer.apple.com/documentation/Foundation/FileWrapper
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/unblockUserInteraction()
func (d NSDocument) UnblockUserInteraction() {
	objc.Send[objc.ID](d.ID, objc.Sel("unblockUserInteraction"))
}
// Writes the contents of the document to a file or file package located by a
// URL, that is formatted to a specified type.
//
// url: The location to which the document contents are written.
//
// typeName: The string that identifies the document type.
//
// # Discussion
// 
// The default implementation of this method just invokes `[self typeName
// outError]` and writes the returned file wrapper to disk.
// 
// For backward binary compatibility with OS X v10.3 and earlier, the default
// implementation of this method instead invokes `[self [absoluteURL path]
// typeName]` if [writeToFile:ofType:] is overridden and the URL uses the ``
// scheme.
//
// [writeToFile:ofType:]: https://developer.apple.com/documentation/AppKit/NSDocument/writeToFile:ofType:
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/write(to:ofType:)
func (d NSDocument) WriteToURLOfTypeError(url foundation.INSURL, typeName string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("writeToURL:ofType:error:"), url, objc.String(typeName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeToURL:ofType:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Writes the contents of the document to a file or file package located by a
// URL.
//
// url: The location to which the document contents are written.
//
// typeName: The string that identifies the document type.
//
// saveOperation: The type of save operation.
//
// # Discussion
// 
// The default implementation of this method invokes
// [WriteToURLOfTypeForSaveOperationOriginalContentsURLError]. It also invokes
// [FileAttributesToWriteToURLOfTypeForSaveOperationOriginalContentsURLError]
// and writes the returned attributes, if any, to the file. It may copy some
// attributes from the old on-disk revision of the document at the same time,
// if applicable.
// 
// This method is responsible for doing document writing in a way that
// minimizes the danger of leaving the disk to which writing is being done in
// an inconsistent state in the event of an app crash, system crash, hardware
// failure, power outage, and so on. If you override this method, be sure to
// invoke the superclass implementation.
// 
// For [NSSaveOperation], the default implementation of this method uses the
// value in the [KeepBackupFile] property to determine whether or not the old
// on-disk revision of the document, if there was one, should be preserved
// after being renamed.
// 
// For backward binary compatibility with OS X v10.3 and earlier, the default
// implementation of this method instead invokes
// [writeWithBackupToFile:ofType:saveOperation:] if that method is is
// overridden and the URL uses the `` scheme. The save operation in this case
// is never [NSAutosaveOperation]; [NSSaveToOperation] is used instead.
//
// [writeWithBackupToFile:ofType:saveOperation:]: https://developer.apple.com/documentation/AppKit/NSDocument/writeWithBackupToFile:ofType:saveOperation:
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/writeSafely(to:ofType:for:)
func (d NSDocument) WriteSafelyToURLOfTypeForSaveOperationError(url foundation.INSURL, typeName string, saveOperation NSSaveOperationType) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("writeSafelyToURL:ofType:forSaveOperation:error:"), url, objc.String(typeName), saveOperation, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeSafelyToURL:ofType:forSaveOperation:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Creates and returns a file wrapper that contains the contents of the
// document, formatted to the specified type.
//
// typeName: The string that identifies the document type.
//
// # Return Value
// 
// A file wrapper containing the document contents, or, if the file wrapper
// could not be created, `nil`.
//
// # Discussion
// 
// For backward binary compatibility with OS X v10.3 and earlier, if
// [fileWrapperRepresentationOfType:] is overridden, the default
// implementation of this method instead invokes `[self typeName]`.
//
// [fileWrapperRepresentationOfType:]: https://developer.apple.com/documentation/AppKit/NSDocument/fileWrapperRepresentationOfType:
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/fileWrapper(ofType:)
func (d NSDocument) FileWrapperOfTypeError(typeName string) (foundation.NSFileWrapper, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("fileWrapperOfType:error:"), objc.String(typeName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSFileWrapper{}, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSFileWrapperFromID(rv), nil

}
// Creates and returns a data object that contains the contents of the
// document, formatted to a specified type.
//
// typeName: The string that identifies the document type.
//
// # Return Value
// 
// A data object containing the document contents, or, if the data object
// could not be created, `nil`.
//
// # Discussion
// 
// The default implementation of this method throws an exception because at
// least one of the writing methods (this method, [WriteToURLOfTypeError],
// [FileWrapperOfTypeError], or
// [WriteToURLOfTypeForSaveOperationOriginalContentsURLError]) must be
// overridden.
// 
// For backward binary compatibility with OS X v10.3 and earlier, the default
// implementation of this method instead invokes ```typeName` on `self` if ``
// is overridden.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/data(ofType:)
func (d NSDocument) DataOfTypeError(typeName string) (foundation.INSData, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("dataOfType:error:"), objc.String(typeName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSData{}, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDataFromID(rv), nil

}
// Writes the contents of the document to a file or file package located by a
// URL.
//
// url: The location to which the document contents are written.
//
// typeName: The string that identifies the document type.
//
// saveOperation: The type of save operation.
//
// absoluteOriginalContentsURL: The location of the previously saved copy of the document (if not `nil`).
//
// # Discussion
// 
// The default implementation of this method merely invokes `[self absoluteURL
// typeName outError]`. You can override this method instead of one of the
// three simple writing methods
// ([WriteToURLOfTypeError],[FileWrapperOfTypeError], and [DataOfTypeError])
// if your document writing machinery needs access to the on-disk
// representation of the document revision that is about to be overwritten.
// The value of `absoluteURL` is often not the same as `[self fileURL]`. Other
// times it is not the same as the URL for the final save destination.
// Likewise, `absoluteOriginalContentsURL` is often not the same value as
// `[self fileURL]`. If `absoluteOriginalContentsURL` is `nil`, either the
// document has never been saved or the user deleted the document file since
// it was opened.
// 
// For backward binary compatibility with OS X v10.3 and earlier, if
// [writeToFile:ofType:originalFile:saveOperation:] is overridden and both
// URLs use the `` scheme, the default implementation of this method instead
// invokes:
// 
// The save operation used in this case is never [NSAutosaveOperation];
// [NSSaveToOperation] is used instead.
//
// [writeToFile:ofType:originalFile:saveOperation:]: https://developer.apple.com/documentation/AppKit/NSDocument/writeToFile:ofType:originalFile:saveOperation:
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/write(to:ofType:for:originalContentsURL:)
func (d NSDocument) WriteToURLOfTypeForSaveOperationOriginalContentsURLError(url foundation.INSURL, typeName string, saveOperation NSSaveOperationType, absoluteOriginalContentsURL foundation.INSURL) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("writeToURL:ofType:forSaveOperation:originalContentsURL:error:"), url, objc.String(typeName), saveOperation, absoluteOriginalContentsURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writeToURL:ofType:forSaveOperation:originalContentsURL:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Saves the contents of the document to a file or file package located by a
// URL, that is formatted to a specified type, for a particular kind of save
// operation.
//
// url: The location of the file or file package to which the document contents are
// saved.
//
// typeName: The string that identifies the document type.
//
// saveOperation: The type of save operation.
//
// delegate: The delegate to which the selector message is sent.
//
// didSaveSelector: The selector of the message sent to the delegate.
//
// contextInfo: Object passed with the callback to provide any additional context
// information.
//
// # Discussion
// 
// When saving is completed, regardless of success or failure, the method
// sends the message selected by `didSaveSelector` to the `delegate`, with the
// `contextInfo` as the last argument. The method selected by
// `didSaveSelector` must have the same signature as:
// 
// The default implementation of this method invokes `[self absoluteURL
// typeName saveOperation &anError]` and, if [false] is returned, presents the
// error to the user in a document-modal panel before messaging the delegate.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/save(to:ofType:for:delegate:didSave:contextInfo:)
func (d NSDocument) SaveToURLOfTypeForSaveOperationDelegateDidSaveSelectorContextInfo(url foundation.INSURL, typeName string, saveOperation NSSaveOperationType, delegate objectivec.IObject, didSaveSelector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](d.ID, objc.Sel("saveToURL:ofType:forSaveOperation:delegate:didSaveSelector:contextInfo:"), url, objc.String(typeName), saveOperation, delegate, didSaveSelector, contextInfo)
}
// Saves the contents of the document to a file or file package located by a
// URL, that is formatted to a specified type, for a particular kind of save
// operation, and invokes the passed-in completion handler.
//
// url: The location to which the document contents are written.
//
// typeName: The document type.
//
// saveOperation: The type of save operation.
//
// completionHandler: The completion handler block object passed in to be invoked at some point
// in the future, perhaps after the method invocation has returned. The
// completion handler must be invoked on the main thread.
// 
// The block takes one argument:
// 
// `errorOrNil`: If successful, pass a `nil` error. If not successful, pass an
// [NSError] object that encapsulates the reason why the document could not be
// saved.
//
// # Discussion
// 
// The default implementation of this method invokes
// [CanAsynchronouslyWriteToURLOfTypeForSaveOperation]. If writing can’t be
// done concurrently, it invokes [WriteSafelyToURLOfTypeForSaveOperationError]
// on the main thread thread. If writing can be done concurrently, it invokes
// that method on a different thread, but blocks the main thread until
// something invokes [UnblockUserInteraction]. Either way, if
// [WriteSafelyToURLOfTypeForSaveOperationError] returns [true], it updates
// the values in the [FileModificationDate], [FileType], [FileURL], and
// [AutosavedContentsFileURL] properties and calls the [UpdateChangeCount]
// method as appropriate on the main thread. It also updates information that
// [SaveDocumentWithDelegateDidSaveSelectorContextInfo] uses to check for
// modification, renaming, moving, deleting, and trashing of open documents,
// and deletes autosaved contents files when they have become obsolete. You
// can override this method to do things that need to be done before or after
// any save operation but be sure to invoke `super`.
// 
// For backward binary compatibility with OS X v10.6 and earlier, the default
// implementation of this method instead invokes
// [saveToURL:ofType:forSaveOperation:error:] if that method is overridden and
// this one is not, and it passes any error to the completion handler.
//
// [saveToURL:ofType:forSaveOperation:error:]: https://developer.apple.com/documentation/AppKit/NSDocument/saveToURL:ofType:forSaveOperation:error:
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/save(to:ofType:for:completionHandler:)
func (d NSDocument) SaveToURLOfTypeForSaveOperationCompletionHandler(url foundation.INSURL, typeName string, saveOperation NSSaveOperationType, completionHandler ErrorHandler) {
_block3, _cleanup3 := NewErrorBlock(completionHandler)
	defer _cleanup3()
	objc.Send[objc.ID](d.ID, objc.Sel("saveToURL:ofType:forSaveOperation:completionHandler:"), url, objc.String(typeName), saveOperation, _block3)
}
// Returns the attributes to write to the file or file package at the
// specified URL, and targeting the specified type of save operation.
//
// url: The location to which the document is being written.
//
// typeName: The string that identifies the document type.
//
// saveOperation: The type of save operation.
//
// absoluteOriginalContentsURL: The location of the previously saved copy of the document (if not `nil`).
//
// # Return Value
// 
// A dictionary containing the attributes to be written, or `nil` if
// unsuccessful.
//
// # Discussion
// 
// Your subclass of [NSDocument] can override this method to control the
// attributes that are set during a save operation. An override of this method
// should return a copy of the dictionary returned by its superclass’s
// version of this method, with appropriate alterations.
// 
// The set of valid file attributes is a subset of those understood by the
// [NSFileManager] class. The default implementation of this method returns an
// empty dictionary for an [NSSaveOperation] or [NSAutosaveInPlaceOperation],
// or a dictionary with an appropriate [NSFileExtensionHidden] entry for any
// other kind of save operation. You can override this method to customize the
// attributes that are written to document files.
// 
// For backward binary compatibility with OS X v10.5 and earlier, the default
// implementation of this method returns a dictionary with [NSFileHFSCode] and
// [NSFileHFSTypeCode] entries that have a value of 0 for [NSSaveOperation],
// in apps linked against OS X v10.5 or earlier.
// 
// For backward binary compatibility with OS X v10.3 and earlier, the default
// implementation of this method instead invokes `[self [url path] typeName
// aSaveOperation]` if [fileAttributesToWriteToFile:ofType:saveOperation:] is
// overridden and the URL uses the `` scheme. The save operation used in this
// case is never one of the autosaving ones: [NSSaveToOperation] is used
// instead.
// 
// The default implementation of [WriteSafelyToURLOfTypeForSaveOperationError]
// automatically copies important attributes like file permissions, creation
// date, and Finder information from the old on-disk version of a document to
// the new one during an [NSSaveOperation] or [NSAutosaveInPlaceOperation].
// This method is meant to be used just for attributes that need to be written
// for the first time, for [NSSaveAsOperation] and [NSSaveToOperation]. The
// `url` and `absoluteOriginalContentsURL` parameters are passed in for
// completeness; NSDocument’s default implementation doesn’t need to use
// them.
//
// [fileAttributesToWriteToFile:ofType:saveOperation:]: https://developer.apple.com/documentation/AppKit/NSDocument/fileAttributesToWriteToFile:ofType:saveOperation:
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/fileAttributesToWrite(to:ofType:for:originalContentsURL:)
func (d NSDocument) FileAttributesToWriteToURLOfTypeForSaveOperationOriginalContentsURLError(url foundation.INSURL, typeName string, saveOperation NSSaveOperationType, absoluteOriginalContentsURL foundation.INSURL) (foundation.INSDictionary, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("fileAttributesToWriteToURL:ofType:forSaveOperation:originalContentsURL:error:"), url, objc.String(typeName), saveOperation, absoluteOriginalContentsURL, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDictionaryFromID(rv), nil

}
// Returns the names of the types to which this document can be saved for a
// specified kind of save operation.
//
// saveOperation: The kind of save operation.
//
// # Return Value
// 
// An array of [NSString] objects representing the writable document types.
//
// # Discussion
// 
// The save operation type is represented by `saveOperation`. For every kind
// of save operation except [NSSaveToOperation], the returned array must only
// include types for which the app can play the Editor role. For
// [NSSaveToOperation] the returned array may include types for which the app
// can only play the Viewer role, and other types that the app can merely
// export. The default implementation of this method returns `[[self class]
// writableTypes]` with, except during [NSSaveToOperation], types for which
// [IsNativeType] returns [false] filtered out.
// 
// You can override this method to limit the set of writable types when the
// document currently contains data that is not representable in all types.
// For example, you can disallow saving to RTF files when the document
// contains an attachment and can only be saved properly to RTFD files.
// 
// You can invoke this method when creating a custom save panel accessory view
// to present easily the same set of types as [NSDocument] does in its
// standard file format popup menu.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/writableTypes(for:)
func (d NSDocument) WritableTypesForSaveOperation(saveOperation NSSaveOperationType) []string {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("writableTypesForSaveOperation:"), saveOperation)
	return objc.ConvertSliceToStrings(rv)
}
// Returns a filename extension that can be appended to a base filename, for a
// specified file type and kind of save operation.
//
// typeName: The file type.
//
// saveOperation: The kind of save operation.
//
// # Return Value
// 
// The filename extension.
//
// # Discussion
// 
// The default implementation of this method invokes
// [preferredFilenameExtension(forType:)] on the shared workspace object if
// the type is a UTI or, if it is not, for backward binary compatibility with
// OS X v10.4 and earlier, invokes [fileExtensionsFromType:] on the shared
// document controller and chooses the first filename extension in the
// returned array.
// 
// You can override this method to customize the appending of extensions to
// filenames by [NSDocument]. Starting in OS X v10.5, it is invoked from only
// two places in AppKit:
// 
// - The [AutosaveDocumentWithDelegateDidAutosaveSelectorContextInfo] method
// uses this method when creating a new filename for the autosaved contents. -
// The [HandleSaveScriptCommand] method uses this method when adding an
// extension to the filename specified by a script.
// 
// In all other cases, the name of any file being saved will have been fully
// specified by the user with the Save panel (whether they know it or not).
//
// [fileExtensionsFromType:]: https://developer.apple.com/documentation/AppKit/NSDocumentController/fileExtensionsFromType:
// [preferredFilenameExtension(forType:)]: https://developer.apple.com/documentation/AppKit/NSWorkspace/preferredFilenameExtension(forType:)
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/fileNameExtension(forType:saveOperation:)
func (d NSDocument) FileNameExtensionForTypeSaveOperation(typeName string, saveOperation NSSaveOperationType) string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("fileNameExtensionForType:saveOperation:"), objc.String(typeName), saveOperation)
	return foundation.NSStringFromID(rv).String()
}
// Creates the window controller objects that the document uses to display its
// content.
//
// # Discussion
// 
// Subclasses may override this method to create the initial window
// controller(s) for the document.
// 
// The base class implementation creates an [NSWindowController] object with
// [WindowNibName] and with the document as the file’s owner if
// [WindowNibName] returns a name. If you override this method to create your
// own window controllers, be sure to use [AddWindowController] to add them to
// the document after creating them.
// 
// This method is called by the [NSDocumentController] `open...` methods, but
// you might want to call it directly in some circumstances.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/makeWindowControllers()
func (d NSDocument) MakeWindowControllers() {
	objc.Send[objc.ID](d.ID, objc.Sel("makeWindowControllers"))
}
// Adds the specified window controller to the current document.
//
// windowController: The window controller that is added.
//
// # Discussion
// 
// An [NSDocument] object uses its list of window controllers when it displays
// all document windows, sets window edited status upon an undo or redo
// operation, and modifies window titles. If you create window controllers by
// overriding [WindowNibName], this method is invoked automatically. If you
// create window controllers in [WindowControllers] or in any other context,
// such as in apps that present multiple windows per document, you should
// invoke this method for each window controller created.
// 
// You cannot attach a window controller to more than one document at a time.
// The default implementation of this method removes the passed-in window
// controller from the document to which it is attached, if it is already
// attached to one, then sends it a [Document] message with `self` as the
// argument. It also ignores redundant invocations.
// 
// You would not typically override this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/addWindowController(_:)
func (d NSDocument) AddWindowController(windowController INSWindowController) {
	objc.Send[objc.ID](d.ID, objc.Sel("addWindowController:"), windowController)
}
// Removes the specified window controller from the receiver’s array of
// window controllers.
//
// windowController: The window controller that is removed.
//
// # Discussion
// 
// A document with no window controllers is not necessarily closed. However, a
// window controller can be set to close its associated document when the
// window is closed or the window controller is deallocated.
// 
// The default implementation of this method sends a [Document] message to the
// passed-in window controller with a `nil` argument. You would not typically
// override this method.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/removeWindowController(_:)
func (d NSDocument) RemoveWindowController(windowController INSWindowController) {
	objc.Send[objc.ID](d.ID, objc.Sel("removeWindowController:"), windowController)
}
// Called after one of the document’s window controllers loads its nib file.
//
// windowController: The window controller that loaded the nib file.
//
// # Discussion
// 
// See the class description for [NSWindowController] for additional
// information about nib files and the file’s owner object.
// 
// Typically an [NSDocument] subclass overrides [WindowNibName] or
// [WindowControllers], but not both. If [WindowNibName] is overridden, the
// default implementation of [WindowControllers] will load the named nib file,
// making the [NSDocument] object the nib file’s owner. In that case, you
// can override [WindowControllerDidLoadNib] and do custom processing after
// the nib file is loaded.
// 
// The default implementation of this method does nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/windowControllerDidLoadNib(_:)
func (d NSDocument) WindowControllerDidLoadNib(windowController INSWindowController) {
	objc.Send[objc.ID](d.ID, objc.Sel("windowControllerDidLoadNib:"), windowController)
}
// Called before one of the document’s window controllers loads its nib
// file.
//
// windowController: The window controller that loads the nib file.
//
// # Discussion
// 
// See the class description for [NSWindowController] for additional
// information about nib files and the file’s owner object.
// 
// Typically an [NSDocument] subclass overrides [WindowNibName] or
// [WindowControllers], but not both. If [WindowNibName] is overridden, the
// default implementation of [WindowControllers] will load the named nib file,
// making the NSDocument the nib file’s owner. In that case, you can
// override [WindowControllerWillLoadNib] and do custom processing before the
// nib file is loaded.
// 
// The default implementation of this method does nothing.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/windowControllerWillLoadNib(_:)
func (d NSDocument) WindowControllerWillLoadNib(windowController INSWindowController) {
	objc.Send[objc.ID](d.ID, objc.Sel("windowControllerWillLoadNib:"), windowController)
}
// Determines whether the system should close the document and its associated
// window.
//
// windowController: The window controller that is closed.
//
// delegate: The delegate to which the selector message is sent.
//
// shouldCloseSelector: The selector of the message sent to the delegate.
//
// contextInfo: Object passed with the callback to provide any additional context
// information.
//
// # Discussion
// 
// If the window controller is the document’s last one, or is marked as
// causing the document to close, this method calls the method in the
// `shouldCloseSelector` parameter with the result of
// [CanCloseDocumentWithDelegateShouldCloseSelectorContextInfo]. In all other
// cases, this method calls `shouldCloseSelector` with [true]. This method is
// called automatically by [NSWindow] for any window that has a window
// controller and a document associated with it. [NSWindow] calls this method
// prior to sending its `delegate` the [WindowShouldClose] message. Pass the
// `contextInfo` object with the callback.
// 
// The `shouldCloseSelector` callback method should have the following
// signature:
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/shouldCloseWindowController(_:delegate:shouldClose:contextInfo:)
func (d NSDocument) ShouldCloseWindowControllerDelegateShouldCloseSelectorContextInfo(windowController INSWindowController, delegate objectivec.IObject, shouldCloseSelector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](d.ID, objc.Sel("shouldCloseWindowController:delegate:shouldCloseSelector:contextInfo:"), windowController, delegate, shouldCloseSelector, contextInfo)
}
// Displays all of the document’s windows, bringing them to the front and
// making them main or key as necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/showWindows()
func (d NSDocument) ShowWindows() {
	objc.Send[objc.ID](d.ID, objc.Sel("showWindows"))
}
// Sets the window outlet of this document to the specified value.
//
// window: The window to which the receiver’s `window` outlet points.
//
// # Discussion
// 
// This method is invoked automatically during the loading of any nib for
// which this document is the file’s owner, if the file’s owner `window`
// outlet is connected in the nib. You should not invoke this method directly,
// and typically you would not override it either.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/setWindow(_:)
func (d NSDocument) SetWindow(window INSWindow) {
	objc.Send[objc.ID](d.ID, objc.Sel("setWindow:"), window)
}
// Returns the default draft name for the document subclass.
//
// # Discussion
// 
// The default implementation of this method returns the string
// “Untitled”, as adjusted according to the user’s specified locale.
// Your app should typically return a name that describes the kind of
// document. For example, a spreadsheet app could return “Spreadsheet”. A
// document created from a template could return the name of the template, for
// example, “Résumé”.
// 
// When a document has not yet been assigned a name, and has not yet been
// autosaved with the [NSDocument.SaveOperationType.autosaveAsOperation] save
// operation type, the document bases the default name on the value in the
// [DisplayName] property.
// 
// If there is a already another document or file in the same place and with
// the same name as would be returned by this method, [NSDocument] appends a
// number to the [DefaultDraftName] string.
//
// [NSDocument.SaveOperationType.autosaveAsOperation]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/autosaveAsOperation
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/defaultDraftName()
func (d NSDocument) DefaultDraftName() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("defaultDraftName"))
	return foundation.NSStringFromID(rv).String()
}
// Saves the interface-related state of the document.
//
// coder: The coder object in which to save the responder’s interface-related
// state.
//
// queue: A serial background operation queue on which to encode additional state
// asynchronously.
//
// # Discussion
// 
// This method is part of the window restoration system and is called at
// appropriate times to save the visual state of your responder to the
// specified archive. The default implementation of this method does nothing
// but specific subclasses (such as [NSView] and [NSWindow]) override it to
// save important state information. Therefore, if you override this method,
// always call `super` at some point in your implementation.
// 
// Subclasses can override this method and use it to restore any information
// that would be needed to restore the responder to its current state. For
// example, the [NSTabView] class uses this method to save information about
// the currently selected tab. You must store enough data to reconfigure the
// responder and return it to its current state during a subsequent launch of
// the application.
// 
// AppKit calls this method on your app’s main thread. If you want to encode
// any state information asynchronously on a background thread, submit one or
// more operation objects to the provided `queue`. Performing long-running
// operations asynchronously lets you free up the main thread for other
// operations more quickly. The encoding operation is not considered final
// until all operations submitted to `queue` finish.
// 
// For information about using a coder object to write data to an archive, see
// [Archives and Serializations Programming Guide].
//
// [Archives and Serializations Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Archiving/Archiving.html#//apple_ref/doc/uid/10000047i
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/encodeRestorableState(with:backgroundQueue:)
func (d NSDocument) EncodeRestorableStateWithCoderBackgroundQueue(coder foundation.INSCoder, queue foundation.NSOperationQueue) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeRestorableStateWithCoder:backgroundQueue:"), coder, queue)
}
// Returns a Boolean value that indicates whether it is safe to autosave
// document changes.
//
// # Discussion
// 
// The default implementation of this method checks for documents that have
// not been changed in a while (“a while” is subject to change) or are
// saved in folders where the user typically does not edit documents (the
// `~/Downloads` folder, for example). When it senses one of those cases it
// returns [false] with an [NSError] object that has recovery options like
// Duplicate, Cancel, and Unlock.
// 
// In an app that has adopted autosaving in place by overriding
// [AutosavesInPlace] to return [true], you can override this method to
// customize the autosaving safety checking that [NSDocument] does by default.
// You can remove the [NSDocument] default checking by overriding this method
// and not invoking super. You can add to the [NSDocument] default checking by
// invoking super and then doing your own checking if `[super ]` did not
// signal an error. For example, TextEdit overrides this method to ask the
// user what to do when opening a document file has been lossy and overwriting
// that file might therefore be lossy.
// 
// When autosaving in place is turned on an [NSDocument] object may invoke
// this method when it receives notification from its [NSUndoManager] object
// that the user changed the document, or undid or redid a change. If an error
// is returned, [NSDocument] presents the error to the user, allowing the user
// to choose a recovery option. If the user chooses a recovery option, then
// [NSDocument] invokes this method again until no error is signaled, to make
// sure that all checks have been done. This means that when you signal an
// error and the user’s choice of recovery option indicates that they have
// seen and disregarded a safety concern, you must record that fact and not do
// that particular safety check again. Once all errors are handled,
// [NSDocument] continues by invoking [UpdateChangeCount]. If the user does
// not recover from an error, then [NSDocument] invokes one of the
// [NSUndoManager] methods [undo()] or [redo()] to roll back the change. So,
// some of the [NSError] recovery options the user can choose, like the
// [NSDocument] options Duplicate and Cancel, should indicate failed recovery
// and cause the document to remain unchanged afterward.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [redo()]: https://developer.apple.com/documentation/Foundation/UndoManager/redo()
// [true]: https://developer.apple.com/documentation/Swift/true
// [undo()]: https://developer.apple.com/documentation/Foundation/UndoManager/undo()
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/checkAutosavingSafety()
func (d NSDocument) CheckAutosavingSafetyAndReturnError() (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("checkAutosavingSafetyAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("checkAutosavingSafetyAndReturnError: returned NO with nil NSError")
	}
	return rv, nil

}
// Schedules periodic autosaving for the purpose of crash protection.
//
// # Discussion
// 
// The default implementation of this method checks to see if autosaving is
// turned on and, if so and if `[self hasUnautosavedChanges]` returns [true],
// schedules an [NSTimer] to invoke
// [AutosaveDocumentWithDelegateDidAutosaveSelectorContextInfo] in the future.
// If `[self hasUnautosavedChanges]` returns [false] it unschedules any
// previously scheduled timer. It takes care not to cause
// [AutosaveDocumentWithDelegateDidAutosaveSelectorContextInfo] to be invoked
// before a previous invocation caused by it has finished. The exact timings
// it uses are complicated and subject to change in future releases of macOS.
// You can override this method to control when exactly periodic autosaving
// happens. It is invoked by [UpdateChangeCount] and
// [UpdateChangeCountWithTokenForSaveOperation].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/scheduleAutosaving()
func (d NSDocument) ScheduleAutosaving() {
	objc.Send[objc.ID](d.ID, objc.Sel("scheduleAutosaving"))
}
// Autosaves the document’s contents to an appropriate location in the file
// system.
//
// delegate: The delegate to which the selector message is sent.
//
// didAutosaveSelector: The selector of the message sent to the delegate.
//
// contextInfo: Object passed with the callback to provide any additional context
// information.
//
// # Discussion
// 
// After autosaving, sends the message selected by `didAutosaveSelector` to
// the delegate, with `contextInfo` as the last argument. The method selected
// by `didAutosaveSelector` must have the same signature as:
// 
// If an error occurs while autosaving, the method reports it to the user
// before sending the delegate a `NO` message.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/autosave(withDelegate:didAutosave:contextInfo:)
func (d NSDocument) AutosaveDocumentWithDelegateDidAutosaveSelectorContextInfo(delegate objectivec.IObject, didAutosaveSelector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](d.ID, objc.Sel("autosaveDocumentWithDelegate:didAutosaveSelector:contextInfo:"), delegate, didAutosaveSelector, contextInfo)
}
// Autosaves the document’s contents to an appropriate file-system location,
// as needed.
//
// autosavingIsImplicitlyCancellable: The value in the [AutosavingIsImplicitlyCancellable] property while
// autosaving is happening.
//
// completionHandler: The completion handler block object passed in to be invoked at some point
// in the future, perhaps after the method invocation has returned. The
// completion handler must be invoked on the main thread.
// 
// The block takes one argument:
// 
// `errorOrNil`: If successful, pass a `nil` error. If not successful, pass an
// [NSError] object that encapsulates the reason why the document could not be
// autosaved.
//
// # Discussion
// 
// The default implementation of this method does the following:
// 
// - Checks the value of the [HasUnautosavedChanges] property. - If the value
// of that property is [false], the method runs the completion handler with a
// `nil` error and returns immediately.
// 
// If the value is [true], calls [AutosavesInPlace] on the class to determine
// where the autosaved document contents should go.
// 
// The method also gets the value in [FileURL] to ensure that the file has an
// actual URL, because it is not possible to autosave in place if the document
// does not yet have a permanent location. 3. Checks the value in the
// [AutosavingFileType] property to determine the file type for the autosaved
// file. 4. Calls [SaveToURLOfTypeForSaveOperationCompletionHandler].
// 
// The value of the `saveToURL` parameter is the location where the file
// should be saved. If the file has a URL and the class specifies that
// autosave should occur in place, this is the URL of the file. Otherwise,
// this is the location of a nonexistent file in the specified autosave
// location.
// 
// The value for the `ofType` parameter is determined by a call to
// [AutosavingFileType].
// 
// The value of the `forSaveOperation` parameter is
// [NSAutosaveInPlaceOperation] if the class is configured to autosave in
// place and the file has a URL. Otherwise, the value is
// [NSAutosaveElsewhereOperation].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/autosave(withImplicitCancellability:completionHandler:)
func (d NSDocument) AutosaveWithImplicitCancellabilityCompletionHandler(autosavingIsImplicitlyCancellable bool, completionHandler ErrorHandler) {
_block1, _cleanup1 := NewErrorBlock(completionHandler)
	defer _cleanup1()
	objc.Send[objc.ID](d.ID, objc.Sel("autosaveWithImplicitCancellability:completionHandler:"), autosavingIsImplicitlyCancellable, _block1)
}
// Opens the Versions browser in the document’s main window.
//
// sender: The control sending the message.
//
// # Discussion
// 
// This is the action of the Browse Saved Versions menu item in a
// document-based app.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/browseVersions(_:)
func (d NSDocument) BrowseDocumentVersions(sender objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("browseDocumentVersions:"), sender)
}
// Dismiss the Versions browser for the current document.
//
// completionHandler: The completion handler block to call when the Versions browser is fully
// dismissed. AppKit calls this block on your app’s main thread, waiting for
// any dismissal animations to complete before calling it. The block block has
// no return value and no parameters.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/stopBrowsingVersions(completionHandler:)
func (d NSDocument) StopBrowsingVersionsWithCompletionHandler(completionHandler VoidHandler) {
_block0, _cleanup0 := NewVoidBlock(completionHandler)
	defer _cleanup0()
	objc.Send[objc.ID](d.ID, objc.Sel("stopBrowsingVersionsWithCompletionHandler:"), _block0)
}
// Moves the document to the user’s iCloud storage.
//
// sender: The control sending the message.
//
// # Discussion
// 
// AppKit calls this method automatically in response to the user selecting
// the Move to iCloud… menu item in a document-based app. The default
// implementation presents the user with an alert asking to confirm the move
// before invoking the [MoveToURLCompletionHandler] method with a URL in the
// app’s default ubiquity container.
// 
// See Moving the Document for descriptions of methods for moving a document
// to a local path.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/moveToUbiquityContainer(_:)
func (d NSDocument) MoveDocumentToUbiquityContainer(sender objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("moveDocumentToUbiquityContainer:"), sender)
}
// Updates the document’s change count settings after a successful save
// operation.
//
// changeCountToken: An object encapsulating the document changes, returned from
// [ChangeCountTokenForSaveOperation].
//
// saveOperation: The type of save operation.
//
// # Discussion
// 
// This method updates the values in the [DocumentEdited] and
// [HasUnautosavedChanges] properties. For example,
// [SaveToURLOfTypeForSaveOperationCompletionHandler] invokes this method, on
// the main thread, when it is done saving. The default implementation of this
// method also sends all of the document’s window controllers
// [SetDocumentEdited] messages when appropriate.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/updateChangeCount(withToken:for:)
func (d NSDocument) UpdateChangeCountWithTokenForSaveOperation(changeCountToken objectivec.IObject, saveOperation NSSaveOperationType) {
	objc.Send[objc.ID](d.ID, objc.Sel("updateChangeCountWithToken:forSaveOperation:"), changeCountToken, saveOperation)
}
// Updates the receiver’s change count according to the given change type.
//
// change: The type of change made to the document. For a list of values, see
// [NSDocument.ChangeType].
// //
// [NSDocument.ChangeType]: https://developer.apple.com/documentation/AppKit/NSDocument/ChangeType
//
// # Discussion
// 
// The change count indicates the document’s edited status; if the change
// count is 0, the document has no changes to save, and if the change count is
// greater than 0, the document has been edited and is unsaved. If you are
// implementing undo and redo in an app, you should increment the change count
// every time you create an undo group and decrement the change count when an
// undo or redo operation is performed.
// 
// Note that if you are using the [NSDocument] default undo/redo features,
// setting the document’s edited status by updating the change count happens
// automatically. You only need to invoke this method when you are not using
// these features.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/updateChangeCount(_:)
func (d NSDocument) UpdateChangeCount(change NSDocumentChangeType) {
	objc.Send[objc.ID](d.ID, objc.Sel("updateChangeCount:"), change)
}
// Returns an object that encapsulates the current record of document changes
// at the beginning of a save operation.
//
// saveOperation: The type of save operation.
//
// # Return Value
// 
// An object encapsulating the document changes.
//
// # Discussion
// 
// The returned object is meant to be passed to
// [UpdateChangeCountWithTokenForSaveOperation] at the end of the save
// operation. For example, [SaveToURLOfTypeForSaveOperationCompletionHandler]
// invokes this method, on the main thread, before it does any actual saving.
// This method facilitates asynchronous saving, during which a user can change
// a document while it is being saved.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/changeCountToken(for:)
func (d NSDocument) ChangeCountTokenForSaveOperation(saveOperation NSSaveOperationType) objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("changeCountTokenForSaveOperation:"), saveOperation)
	return objectivec.Object{ID: rv}
}
// Saves the interface-related state of the document.
//
// coder: The coder object in which to save the document’s interface-related state.
//
// # Discussion
// 
// This method is part of the window restoration system and is called at
// appropriate times to save your document’s window-related state
// information to the specified archive. The default implementation of this
// method records some basic information about the document. If you override
// this method, you must call `super` at some point in your implementation.
// 
// Subclasses can override this method and use it to restore any information
// that would be needed to restore the document’s window to its current
// state. For example, you could use this method to record references to the
// data currently managed by the document and displayed by the window. (Do not
// store the actual data itself. Store only references to the data so that you
// can load it later from disk.) You must store enough data to reconfigure the
// document and its window to their current state during a subsequent launch
// of the app.
// 
// For information about using a coder object to write data to an archive, see
// [Archives and Serializations Programming Guide].
//
// [Archives and Serializations Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Archiving/Archiving.html#//apple_ref/doc/uid/10000047i
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/encodeRestorableState(with:)
func (d NSDocument) EncodeRestorableStateWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("encodeRestorableStateWithCoder:"), coder)
}
// Restores the interface-related state of the document.
//
// coder: The coder object to use to restore the document’s interface-related
// state.
//
// # Discussion
// 
// This method is part of the window restoration system and is called at
// launch time to restore the window-related state of your document object.
// The default implementation restores some basic information about the
// document. If you override this method, you must call `super` at some point
// in your implementation.
// 
// Subclasses can override this method and use it to restore the
// document-related information that was saved in the
// [EncodeRestorableStateWithCoder] method. You can also use this method to
// reconfigure the document (or its associated window controller and window)
// to their previous appearance.
// 
// For information about using a coder object to read data from an archive,
// see [Encoding and Decoding Custom Types].
//
// [Encoding and Decoding Custom Types]: https://developer.apple.com/documentation/Foundation/encoding-and-decoding-custom-types
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/restoreState(with:)
func (d NSDocument) RestoreStateWithCoder(coder foundation.INSCoder) {
	objc.Send[objc.ID](d.ID, objc.Sel("restoreStateWithCoder:"), coder)
}
// Marks the document’s interface-related state as dirty.
//
// # Discussion
// 
// Call this method whenever the restorable state of your document changes.
// This method marks the document’s state as dirty, which causes that state
// to be written to disk at some point in the future. Do not override this
// method.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/invalidateRestorableState()
func (d NSDocument) InvalidateRestorableState() {
	objc.Send[objc.ID](d.ID, objc.Sel("invalidateRestorableState"))
}
// Restores a window that was associated with a document, after that document
// is reopened.
//
// identifier: The unique interface item identifier string that was previously associated
// with the window. Use this string to determine which window to create.
//
// state: A coder object containing the window state information. This coder object
// contains the combined restorable state of the window, which can include the
// state of the window, its delegate, window controller, and document object.
// You can use this state to determine which window to create.
//
// completionHandler: A block object to execute with the results of creating the window. You must
// execute this block at some point but may do so after the method returns if
// needed. This block takes the following parameters:
// 
// - The window that was created or `nil` if the window could not be created.
// - An error object if the window was not recognized or could not be created
// for whatever reason; otherwise, specify `nil`.
//
// # Discussion
// 
// This method is called by the default [NSDocumentController] implementation
// of [RestoreWindowWithIdentifierStateCompletionHandler].
// 
// The default implementation of this method first checks if the document has
// window controllers, and if not, it calls [WindowControllers]. If there is
// then exactly one window controller, it invokes the completion handler with
// its window. If there is more than one, it searches the receiver’s window
// controllers for a window that matches the given identifier, and then calls
// the completion handler with it. If no window could be found, it invokes the
// completion handler with a `nil` window.
// 
// If your document has variable or optional windows, you may override this to
// create the requested window, and then call the completion handler with it.
// This allows you to use the default document reopening behavior, but
// intervene at the point of creating the windows. The parameters are the same
// as in the class method [RestoreWindowWithIdentifierStateCompletionHandler].
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/restoreWindow(withIdentifier:state:completionHandler:)
func (d NSDocument) RestoreDocumentWindowWithIdentifierStateCompletionHandler(identifier NSUserInterfaceItemIdentifier, state foundation.INSCoder, completionHandler WindowErrorHandler) {
_block2, _cleanup2 := NewWindowErrorBlock(completionHandler)
	defer _cleanup2()
	objc.Send[objc.ID](d.ID, objc.Sel("restoreDocumentWindowWithIdentifier:state:completionHandler:"), identifier, state, _block2)
}
// Presents a modal Save panel to the user, then tries to save the document if
// the user approves the operation.
//
// saveOperation: The type of save operation.
//
// delegate: The delegate to which the selector message is sent.
//
// didSaveSelector: The selector of the message sent to the delegate.
//
// contextInfo: Object passed with the callback to provide any additional context
// information.
//
// # Discussion
// 
// When saving is completed, regardless of success or failure, or has been
// canceled, sends the message selected by `didSaveSelector` to the
// `delegate`, with `contextInfo` as the last argument. The method selected by
// `didSaveSelector` must have the same signature as:
// 
// Invoked from [SaveDocumentWithDelegateDidSaveSelectorContextInfo], and from
// the [SaveDocumentAs] and [SaveDocumentTo] action methods. The default
// implementation of this method first makes sure that any editor registered
// using the Cocoa Bindings [NSEditorRegistration] informal protocol has
// committed its changes, then creates a Save panel, adds a standard file
// format accessory view (if there is more than one file type for the user to
// choose from and [ShouldRunSavePanelWithAccessoryView] returns [true]), sets
// various attributes of the panel, invokes [PrepareSavePanel] to provide an
// opportunity for customization, then presents the panel. If the user
// approves the panel, the default implementation sends the message
// [SaveToURLOfTypeForSaveOperationDelegateDidSaveSelectorContextInfo].
// 
// For backward binary compatibility with Mac OS v10.3 and earlier, the
// default implementation of this method instead invokes the deprecated
// [saveToFile:saveOperation:delegate:didSaveSelector:contextInfo:] if it is
// overridden, even if the user cancels the panel.
//
// [saveToFile:saveOperation:delegate:didSaveSelector:contextInfo:]: https://developer.apple.com/documentation/AppKit/NSDocument/saveToFile:saveOperation:delegate:didSaveSelector:contextInfo:
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/runModalSavePanel(for:delegate:didSave:contextInfo:)
func (d NSDocument) RunModalSavePanelForSaveOperationDelegateDidSaveSelectorContextInfo(saveOperation NSSaveOperationType, delegate objectivec.IObject, didSaveSelector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](d.ID, objc.Sel("runModalSavePanelForSaveOperation:delegate:didSaveSelector:contextInfo:"), saveOperation, delegate, didSaveSelector, contextInfo)
}
// Tells the document to customize the specified Save panel.
//
// savePanel: The Save panel.
//
// # Return Value
// 
// [true] if successfully prepared; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The default implementation is empty and returns [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/prepareSavePanel(_:)
func (d NSDocument) PrepareSavePanel(savePanel INSSavePanel) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("prepareSavePanel:"), savePanel)
	return rv
}
// Updates the state of the given user activity.
//
// activity: The user activity to be updated.
//
// # Discussion
// 
// The default implementation of this method puts the document’s [FileURL]
// into the [NSUserActivity] object’s [userInfo] dictionary with the
// [NSUserActivityDocumentURLKey]. [NSDocument] automatically sets the
// [needsSave] property of the [NSUserActivity] to [true] when the [FileURL]
// changes.
//
// [NSUserActivityDocumentURLKey]: https://developer.apple.com/documentation/AppKit/NSUserActivityDocumentURLKey
// [NSUserActivity]: https://developer.apple.com/documentation/Foundation/NSUserActivity
// [needsSave]: https://developer.apple.com/documentation/Foundation/NSUserActivity/needsSave
// [true]: https://developer.apple.com/documentation/Swift/true
// [userInfo]: https://developer.apple.com/documentation/Foundation/NSUserActivity/userInfo
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/updateUserActivityState(_:)
func (d NSDocument) UpdateUserActivityState(activity foundation.NSUserActivity) {
	objc.Send[objc.ID](d.ID, objc.Sel("updateUserActivityState:"), activity)
}
// Validates the specified user interface item that the receiver manages.
//
// item: The user interface item to validate.
//
// # Return Value
// 
// [true] if the item is valid; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// These items currently include only Revert and Save. You can override this
// method to add more selectors validated by your document subclass.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/validateUserInterfaceItem(_:)
func (d NSDocument) ValidateUserInterfaceItem(item NSValidatedUserInterfaceItem) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("validateUserInterfaceItem:"), item)
	return rv
}
// Waits for any scheduled file access to complete, then invokes the passed-in
// block.
//
// block: A block that performs file access.
//
// # Discussion
// 
// Given a block that will perform file access, this method waits for any file
// access scheduled by previous invocations of this method or
// [PerformAsynchronousFileAccessUsingBlock] to complete, then invokes the
// passed-in block. When the block invocation returns, the method allows the
// next scheduled file access to to be performed, if any.
// 
// Like [PerformActivityWithSynchronousWaitingUsingBlock], this method’s
// primary use is to wait for asynchronous saving, but in contrast with that
// method it is only for the part of an asynchronous saving operation that
// actually touches the document’s file or values in memory that are
// relative to the document’s file.
// 
// In general, you should use this method or
// [PerformAsynchronousFileAccessUsingBlock] around code that gets or sets
// values in memory that only make sense in the context of the document
// file’s current state. For example, [NSDocument] itself consistently uses
// this mechanism when using the following methods and properties:
// 
// - The [FileType], [FileURL], [FileModificationDate], and
// [AutosavedContentsFileURL] properties, because you can’t reliably make
// decisions based on a file’s location, type, or modification date when it
// is being asynchronously moved, renamed, or changed at that moment. - The
// [DocumentEdited] and [HasUnautosavedChanges] properties, because you
// can’t reliably make decisions based on whether the document’s contents
// in memory have been saved to a file when it is being asynchronously saved
// at that moment. - [UpdateChangeCountWithTokenForSaveOperation] and,
// sometimes, [UpdateChangeCount], to make using this mechanism when invoking
// [DocumentEdited] and [HasUnautosavedChanges] meaningful.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/performSynchronousFileAccess(_:)
func (d NSDocument) PerformSynchronousFileAccessUsingBlock(block VoidHandler) {
_block0, _cleanup0 := NewVoidBlock(block)
	defer _cleanup0()
	objc.Send[objc.ID](d.ID, objc.Sel("performSynchronousFileAccessUsingBlock:"), _block0)
}
// Waits for any scheduled file access to complete but without blocking the
// main thread, then invokes the passed-in block.
//
// block: A block that performs file access.
//
// # Discussion
// 
// This method does the same sort of work as
// [PerformSynchronousFileAccessUsingBlock], but without ever blocking the
// main thread, and may not invoke the block until after the method invocation
// has returned, though still always on the same thread as the method
// invocation. The block is passed another block, the file access completion
// handler, which must be invoked when the file access is complete, though it
// can be invoked from any thread. This method is for use with file access
// that might begin on one thread but continue on another before it is
// complete. For example, [SaveToURLOfTypeForSaveOperationCompletionHandler]
// uses this method instead of [PerformSynchronousFileAccessUsingBlock]
// because if it does asynchronous saving then there is no way for it to
// complete all of its file access before returning from the file access
// block.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/performAsynchronousFileAccess(_:)
func (d NSDocument) PerformAsynchronousFileAccessUsingBlock(block VoidHandler) {
_block0, _cleanup0 := NewVoidBlock(block)
	defer _cleanup0()
	objc.Send[objc.ID](d.ID, objc.Sel("performAsynchronousFileAccessUsingBlock:"), _block0)
}
// Waits for any work scheduled by previous invocations of this method to
// complete, then invokes the passed-in block.
//
// waitSynchronously: If [true], the method does not return until previous activities are
// complete and the passed-in block has been invoked. If [false], the method
// might return before passed-in block is invoked. It might instead be invoked
// later, on the main thread, after previous activities are complete.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// block: A block that performs work that might result in the presentation of a modal
// dialog.
//
// # Discussion
// 
// The block is passed another block, the activity completion handler, which
// must be invoked when the activity is complete.
// 
// This method’s primary use is to wait for asynchronous saving. With
// asynchronous saving it is possible for the user to instigate a user
// interface action that might present modal dialog, a sheet for example, when
// asynchronous saving is about to fail and present an error alert sheet of
// its own, which would not work. This method solves that problem. If your
// [NSDocument] subclass supports asynchronous saving you should invoke this
// method around the performance of any work that might cause the presentation
// of a modal dialog, regardless of whether that work is performed
// synchronously or asynchronously. Here is a list of [NSDocument] methods
// whose default implementations invoke this method because they might present
// sheets, either to ask the user what to do as they begin their work or
// because they may fail and present errors to user:
// 
// - [RevertDocumentToSaved] -
// [SaveDocumentWithDelegateDidSaveSelectorContextInfo] -
// [RunModalSavePanelForSaveOperationDelegateDidSaveSelectorContextInfo] -
// [SaveToURLOfTypeForSaveOperationDelegateDidSaveSelectorContextInfo] -
// [CanCloseDocumentWithDelegateShouldCloseSelectorContextInfo] -
// [DuplicateDocumentWithDelegateDidDuplicateSelectorContextInfo] -
// [RunModalPageLayoutWithPrintInfoDelegateDidRunSelectorContextInfo] -
// [PrintDocumentWithSettingsShowPrintPanelDelegateDidPrintSelectorContextInfo]
// - [RunModalPrintOperationDelegateDidRunSelectorContextInfo]
// 
// More uses of this method may be added to [NSDocument] in the future.
// 
// This method must be invoked on the main thread. If [true] is passed for the
// `waitSynchronously` parameter, the method waits on the main thread,
// blocking further user interaction with the document. The purpose of
// blocking the main thread is so that the user cannot continue to change the
// document while an activity is pending. This prevents, for example, the
// situation in which the user chooses to revert the document, but reverting
// does not happen immediately because asynchronous saving is still in
// progress, yet the user is able to continue changing the document, and then
// those changes are immediately discarded when the asynchronous saving is
// complete and the document is reverted. All of the [NSDocument] methods
// listed above pass [true] for `waitSynchronously`.
// 
// You pass [false] for `waitSynchronously` when the work to be done is
// instigated by the user so indirectly that the work might begin when a modal
// dialog is already being presented. For example, another method whose
// default implementation invokes this method, this time passing [false] for
// `waitSynchronously`, is:
// 
// [AutosaveDocumentWithDelegateDidAutosaveSelectorContextInfo]
// 
// This method might present an error alert, but it is typically invoked by a
// timer. If it passed [true] for `waitSynchronously`, and the timer fired
// while the user was looking at a sheet presented by a previous activity,
// blocking of the main thread would prevent the handling of the user
// interface events necessary to dismiss that sheet and complete that previous
// activity. Deadlock would result.
// 
// Whether you make this method wait synchronously or asynchronously to do
// your work is separate from whether your work is done synchronously or
// asynchronously. For example, as mentioned above,
// [SaveToURLOfTypeForSaveOperationDelegateDidSaveSelectorContextInfo] passes
// [true] for `waitSynchronously` when it uses this method, even though the
// majority of the work it does may be done asynchronously.
// 
// You should not invoke this method during the invocation of the block passed
// to [PerformSynchronousFileAccessUsingBlock] or in between the time
// [PerformAsynchronousFileAccessUsingBlock] invokes the block passed to it
// and the time at which the corresponding file access completion handler is
// invoked. If you do, deadlock can result. In other words, you cannot begin a
// new activity as part of file access. You can, on the other hand, invoke
// [PerformSynchronousFileAccessUsingBlock] or
// [PerformAsynchronousFileAccessUsingBlock] as part of an activity.
// 
// Some asynchronous activities, such as saving, need to do work on the main
// thread as they are completing. A deadlock would be inevitable if there were
// no way to interrupt this method’s blocking of the main thread. See
// [ContinueAsynchronousWorkOnMainThreadUsingBlock] to find out how to
// interrupt this method’s blocking of the main thread.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/performActivity(withSynchronousWaiting:using:)
func (d NSDocument) PerformActivityWithSynchronousWaitingUsingBlock(waitSynchronously bool, block VoidHandler) {
_block1, _cleanup1 := NewVoidBlock(block)
	defer _cleanup1()
	objc.Send[objc.ID](d.ID, objc.Sel("performActivityWithSynchronousWaiting:usingBlock:"), waitSynchronously, _block1)
}
// Continues to perform the task for a user activity object using a different
// block.
//
// block: The block to be invoked.
//
// # Discussion
// 
// When AppKit calls [PerformActivityWithSynchronousWaitingUsingBlock]
// recursively, it may execute this method with the specified block to avoid a
// deadlock.
// 
// If a block that was passed to
// [PerformActivityWithSynchronousWaitingUsingBlock] is being invoked, this
// method invokes the passed-in block, having recorded state that makes inner
// invocations of [PerformActivityWithSynchronousWaitingUsingBlock] not wait.
// If this method is invoked outside of an invocation of a block passed to
// [PerformActivityWithSynchronousWaitingUsingBlock], this method simply
// invokes the passed-in block.
// 
// This method is useful when code executed in a block passed to
// [PerformActivityWithSynchronousWaitingUsingBlock] may also invoke that
// method. For example, [SaveDocumentWithDelegateDidSaveSelectorContextInfo],
// which uses [PerformActivityWithSynchronousWaitingUsingBlock], uses this
// around its invocation of
// [RunModalSavePanelForSaveOperationDelegateDidSaveSelectorContextInfo] or
// [SaveToURLOfTypeForSaveOperationDelegateDidSaveSelectorContextInfo] because
// both of those methods also use
// [PerformActivityWithSynchronousWaitingUsingBlock]. Without the use of this
// method the inner invocation of
// [PerformActivityWithSynchronousWaitingUsingBlock] would wait forever.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/continueActivity(_:)
func (d NSDocument) ContinueActivityUsingBlock(block VoidHandler) {
_block0, _cleanup0 := NewVoidBlock(block)
	defer _cleanup0()
	objc.Send[objc.ID](d.ID, objc.Sel("continueActivityUsingBlock:"), _block0)
}
// Invokes the passed-in block on the main thread.
//
// block: The block to be invoked.
//
// # Discussion
// 
// If the main thread is blocked by an invocation of
// [PerformActivityWithSynchronousWaitingUsingBlock] or
// [PerformSynchronousFileAccessUsingBlock], this method interrupts that
// blocking activity, performs the specified `block`, and then resumes the
// blocking activity after `block` returns. Invocations of this method always
// return before the passed-in block is invoked.
// 
// You can invoke this method when work is being done on a non-main thread and
// part of the work must be continued on the main thread. For example,
// [SaveToURLOfTypeForSaveOperationCompletionHandler] uses this method when it
// has just completed the actual writing of the file during asynchronous
// saving and, to finish the saving operation, must invoke
// [UpdateChangeCountWithTokenForSaveOperation] and other methods on the main
// thread.
// 
// This method can be invoked on any thread.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/continueAsynchronousWorkOnMainThread(_:)
func (d NSDocument) ContinueAsynchronousWorkOnMainThreadUsingBlock(block VoidHandler) {
_block0, _cleanup0 := NewVoidBlock(block)
	defer _cleanup0()
	objc.Send[objc.ID](d.ID, objc.Sel("continueAsynchronousWorkOnMainThreadUsingBlock:"), _block0)
}
// Prints the receiver in response to the user choosing the Print menu
// command.
//
// sender: The control sending the message.
//
// # Discussion
// 
// An [NSDocument] object receives this action message as it travels up the
// responder chain. The default implementation invokes
// [PrintDocumentWithSettingsShowPrintPanelDelegateDidPrintSelectorContextInfo].
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/printDocument(_:)
func (d NSDocument) PrintDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("printDocument:"), sender)
}
// The action method invoked in the receiver as first responder when the user
// chooses the Page Setup menu command.
//
// sender: The control sending the message.
//
// # Discussion
// 
// The default implementation invokes
// [RunModalPageLayoutWithPrintInfoDelegateDidRunSelectorContextInfo] with the
// document’s current NSPrintInfo object as argument; if the user clicks the
// OK button and the document authorizes changes to its printing information
// ([ShouldChangePrintInfo]), the method sets the document’s new
// [NSPrintInfo] object and increments the document’s change count.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/runPageLayout(_:)
func (d NSDocument) RunPageLayout(sender objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("runPageLayout:"), sender)
}
// The action of the File menu item Revert in a document-based app.
//
// sender: The control sending the message.
//
// # Discussion
// 
// The default implementation of this method presents an alert dialog giving
// the user the opportunity to cancel the operation. If the user chooses to
// continue, the method ensures that any editor registered using the Cocoa
// Bindings [NSEditorRegistration] informal protocol has discarded its changes
// and then invokes [RevertToContentsOfURLOfTypeError]. If that returns
// [false], the method presents the error to the user in an document-modal
// alert dialog.
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/revertToSaved(_:)
func (d NSDocument) RevertDocumentToSaved(sender objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("revertDocumentToSaved:"), sender)
}
// The action method invoked in the receiver as first responder when the user
// chooses the Save menu command.
//
// sender: The control sending the message.
//
// # Discussion
// 
// The default implementation saves the document in two different ways,
// depending on whether the document has a file path and a document type
// assigned. If path and type are assigned, it simply writes the document
// under its current file path and type after making a backup copy of the
// previous file. If the document is new (no file path and type), it runs the
// modal Save panel to get the file location under which to save the document.
// It writes the document to this file, sets the document’s file location
// and document type (if a native type), and clears the document’s edited
// status.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/save(_:)
func (d NSDocument) SaveDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("saveDocument:"), sender)
}
// The action method invoked in the receiver as first responder when the user
// chooses the Save As menu command.
//
// sender: The control sending the message.
//
// # Discussion
// 
// The default implementation runs the modal Save panel to get the file
// location under which to save the document. It writes the document to this
// file, sets the document’s file location and document type (if a native
// type), and clears the document’s edited status.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/saveAs(_:)
func (d NSDocument) SaveDocumentAs(sender objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("saveDocumentAs:"), sender)
}
// The action method invoked in the receiver as first responder when the user
// chooses the Save To menu command.
//
// sender: The control sending the message.
//
// # Discussion
// 
// The default implementation is identical to [SaveDocumentAs] except that
// this method doesn’t clear the document’s edited status and doesn’t
// reset file location and document type if the document is a native type.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/saveTo(_:)
func (d NSDocument) SaveDocumentTo(sender objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("saveDocumentTo:"), sender)
}
// Saves the document and delivers the results to the provided delegate
// object.
//
// delegate: The delegate to which the selector message is sent.
//
// didSaveSelector: The selector of the message sent to the delegate.
//
// contextInfo: Object passed with the callback to provide any additional context
// information.
//
// # Discussion
// 
// If an [NSDocument.SaveOperationType.saveOperation] can be performed without
// further user intervention (at the very least, neither [FileURL] nor
// [FileType] are `nil`), then the method immediately saves the document.
// Otherwise, it presents a Save panel to the user and saves the document if
// the user approves the panel. When saving has been completed or canceled,
// the method sends the message selected by `didSaveSelector` to the
// `delegate`, with the `contextInfo` as the last argument.
// 
// As of OS X v10.5, this method checks to see if the document’s file has
// been modified since the document was opened or most recently saved or
// reverted, in addition to the checking for file moving, renaming, and
// trashing that it has done since OS X v10.1. When it senses file
// modification it presents an alert telling the user “This document’s
// file has been changed by another app since you opened or saved it,”
// giving them the choice of saving or not saving. For backward binary
// compatibility this is only done in apps linked against macOS 10.5 or later.
// 
// The `didSaveSelector` callback method should have the following signature:
//
// [NSDocument.SaveOperationType.saveOperation]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/saveOperation
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/save(withDelegate:didSave:contextInfo:)
func (d NSDocument) SaveDocumentWithDelegateDidSaveSelectorContextInfo(delegate objectivec.IObject, didSaveSelector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](d.ID, objc.Sel("saveDocumentWithDelegate:didSaveSelector:contextInfo:"), delegate, didSaveSelector, contextInfo)
}
// Determines whether to close the document, prompting the user as needed to
// choose a course of action.
//
// delegate: The delegate to which the selector message is sent.
//
// shouldCloseSelector: The selector of the message sent to the delegate.
//
// contextInfo: Object passed with the callback to provide any additional context
// information.
//
// # Discussion
// 
// If the document is not dirty, this method immediately calls the
// `shouldCloseSelector` callback on the specified delegate with [true].
// 
// If the document is dirty, an alert is presented giving the user a chance to
// save, not save, or cancel. If the user chooses to save, this method saves
// the document. If the save completes successfully, this method calls the
// callback with [true]. If the save is canceled or otherwise unsuccessful,
// this method calls the callback with [false]. This method may be called by
// [ShouldCloseWindowControllerDelegateShouldCloseSelectorContextInfo]. It is
// also called by the [NSDocumentController] method
// [CloseAllDocumentsWithDelegateDidCloseAllSelectorContextInfo]. You should
// call it before you call [Close] if you are closing the document and want to
// give the user a chance to save any edits. Pass the `contextInfo` object
// with the callback.
// 
// The `shouldCloseSelector` callback method should have the following
// signature:
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/canClose(withDelegate:shouldClose:contextInfo:)
func (d NSDocument) CanCloseDocumentWithDelegateShouldCloseSelectorContextInfo(delegate objectivec.IObject, shouldCloseSelector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](d.ID, objc.Sel("canCloseDocumentWithDelegate:shouldCloseSelector:contextInfo:"), delegate, shouldCloseSelector, contextInfo)
}
// Closes all of the document’s windows and removes the document from its
// document controller.
//
// # Discussion
// 
// This method closes the document immediately, without asking users if they
// want to save the document. This method may not always be called.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/close()
func (d NSDocument) Close() {
	objc.Send[objc.ID](d.ID, objc.Sel("close"))
}
// Discards all unsaved document modifications and replaces the document’s
// contents by reading a file or file package located by a URL of a specified
// type.
//
// url: The location from which the document contents are read.
//
// typeName: The string that identifies the document type.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/revert(toContentsOf:ofType:)
func (d NSDocument) RevertToContentsOfURLOfTypeError(url foundation.INSURL, typeName string) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](d.ID, objc.Sel("revertToContentsOfURL:ofType:error:"), url, objc.String(typeName), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("revertToContentsOfURL:ofType:error: returned NO with nil NSError")
	}
	return rv, nil

}
// Creates a new document whose contents are the same as the receiver and
// returns an error object if unsuccessful.
//
// # Return Value
// 
// The new document if duplication is successful; otherwise `nil`.
//
// # Discussion
// 
// The new document returned doesn’t yet have a value to return from
// [FileURL].
// 
// The default implementation of this method first uses
// [WriteSafelyToURLOfTypeForSaveOperationError] to write the document’s
// current contents to a file located in the same directory that is used for
// the autosaved contents of untitled documents and with the same sort of
// name, then invokes `[[NSDocumentController sharedDocumentController]
// newContentsURL NO aDisplayName outError]`.
// 
// You can override this method to customize what is done during document
// duplication, but if your override does not invoke `[NSDocumentController ]`
// you must take care to do things that that method does, especially invoking
// `[NSDocumentController ]` and `[NSFileCoordinator ]`.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/duplicate()
func (d NSDocument) DuplicateAndReturnError() (INSDocument, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("duplicateAndReturnError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSDocument{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSDocumentFromID(rv), nil

}
// Creates a copy of the receiving document in response to the user choosing
// Duplicate from the File menu.
//
// sender: The control sending the action message.
//
// # Discussion
// 
// The default implementation of this method merely invokes `[self nil NULL
// NULL]`.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/duplicate(_:)
func (d NSDocument) DuplicateDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("duplicateDocument:"), sender)
}
// Creates a new document whose contents are the same as the current document.
//
// delegate: The delegate to which the selector message is sent.
//
// didDuplicateSelector: The selector of the message sent to the delegate.
//
// contextInfo: Object passed with the callback to provide any additional context
// information.
//
// # Discussion
// 
// The new document that is created doesn’t yet have a value to return from
// [FileURL]. When duplicating is completed, regardless of success or failure,
// or whether the operation is rejected by the user, this method sends the
// message indicated by `didDuplicateSelector` to the delegate, with
// `contextInfo` as the last argument. The method selected by
// `didDuplicateSelector` must have the same signature as:
// 
// The default implementation of this method first makes sure that any editor
// registered using the Cocoa Bindings [NSEditorRegistration] informal
// protocol has committed its changes, then checks to see if there are recent
// changes that might have been inadvertent and, if so, presents a panel
// giving the user the choice of canceling, duplicating, or duplicating then
// discarding recent changes. Unless the user cancels duplicating, or if no
// panel was presented, it then invokes [DuplicateAndReturnError]. If the user
// chose duplicating and discarding, it also discards recent changes after
// duplicating.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/duplicate(withDelegate:didDuplicate:contextInfo:)
func (d NSDocument) DuplicateDocumentWithDelegateDidDuplicateSelectorContextInfo(delegate objectivec.IObject, didDuplicateSelector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](d.ID, objc.Sel("duplicateDocumentWithDelegate:didDuplicateSelector:contextInfo:"), delegate, didDuplicateSelector, contextInfo)
}
// Renames the current document in response to the user choosing the Rename
// menu item.
//
// sender: The control sending the message.
//
// # Discussion
// 
// This is the action method of the Rename menu item in a document-based app.
// The default implementation of this method initiates a renaming session in a
// window created by the `[self windowForSheet]` message.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/rename(_:)
func (d NSDocument) RenameDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("renameDocument:"), sender)
}
// Moves the document to a new location in response to the user choosing the
// Move To… menu item.
//
// sender: The control sending the message.
//
// # Discussion
// 
// This is the action method of the Move To… menu item in a document-based
// app. By default, this method invokes the
// [MoveDocumentWithCompletionHandler] method, passing `nil` as a parameter
// value.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/move(_:)
func (d NSDocument) MoveDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("moveDocument:"), sender)
}
// Moves the document to a user-selected location.
//
// completionHandler: The completion handler block object passed in to be invoked after moving is
// completed, regardless of success, failure, or cancellation of moving
// action.
//
// # Discussion
// 
// This method presents the user with a move panel if `[self fileURL]` is
// non-nil and then tries to save the document to the new location by invoking
// the [MoveToURLCompletionHandler] method if the user accepts the location
// presented by the panel. If a file with the same name already exists at that
// location, the user will be asked to choose between replacing the
// pre-existing file, renaming the current document, or canceling the move
// process. If `[self fileURL]` is `nil`, then the `[self NSSaveAsOperation ]`
// message is sent instead.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/move(completionHandler:)
func (d NSDocument) MoveDocumentWithCompletionHandler(completionHandler BoolHandler) {
_block0, _cleanup0 := NewBoolBlock(completionHandler)
	defer _cleanup0()
	objc.Send[objc.ID](d.ID, objc.Sel("moveDocumentWithCompletionHandler:"), _block0)
}
// Moves the document’s file to the given URL.
//
// url: The location where the file will ultimately end up, if the move is
// successful.
//
// completionHandler: The completion handler block object passed in to be invoked at some point
// in the future, perhaps after the method invocation has returned. The
// completion handler must be invoked on the main thread. On output, a `nil`
// error is passed if the move is successful; otherwise an [NSError] object is
// passed that encapsulates the reason for failure.
// //
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// # Discussion
// 
// The default implementation of this method replaces any file that may
// currently exist at the given URL with the one being moved, as necessary.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/move(to:completionHandler:)
func (d NSDocument) MoveToURLCompletionHandler(url foundation.INSURL, completionHandler ErrorHandler) {
_block1, _cleanup1 := NewErrorBlock(completionHandler)
	defer _cleanup1()
	objc.Send[objc.ID](d.ID, objc.Sel("moveToURL:completionHandler:"), url, _block1)
}
// Locks the document in response to the user choosing the Lock menu item.
//
// sender: The control sending the message.
//
// # Discussion
// 
// This is the action of the Lock menu item in a document-based app. This
// action method invokes the [LockDocumentWithCompletionHandler] method by
// default.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/lock(_:)
func (d NSDocument) LockDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("lockDocument:"), sender)
}
// Unlocks the document in response to the user choosing the Unlock menu item.
//
// sender: The control sending the message.
//
// # Discussion
// 
// This is the action of the Unlock menu item in a document-based app. This
// action method invokes the [UnlockDocumentWithCompletionHandler] method by
// default.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/unlock(_:)
func (d NSDocument) UnlockDocument(sender objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("unlockDocument:"), sender)
}
// Prevents the user from making further changes to the document.
//
// completionHandler: The completion handler block object passed in to be invoked after locking
// is completed, regardless of success or failure of locking.
//
// # Discussion
// 
// By default, this method first ensures that any editor who has registered
// using Cocoa Binding’s NSEditorRegistration informal protocol has
// committed all changes and then autosaves the document, if necessary, before
// attempting to lock it using the [LockWithCompletionHandler] method. Upon
// successful locking, the [Locked] property is set to `[YES]`. Documents
// whose [FileURL] property is set to `nil` cannot be locked.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/lock(completionHandler:)-6zuhh
func (d NSDocument) LockDocumentWithCompletionHandler(completionHandler BoolHandler) {
_block0, _cleanup0 := NewBoolBlock(completionHandler)
	defer _cleanup0()
	objc.Send[objc.ID](d.ID, objc.Sel("lockDocumentWithCompletionHandler:"), _block0)
}
// Prevents the user from making changes to the document’s file.
//
// completionHandler: The completion handler block object passed in to be invoked after locking
// is completed, regardless of success or failure of locking.
//
// # Discussion
// 
// This method first locks the file at `[self fileURL]` and then invokes the
// given block. The default locking implementation is to enable the “user
// immutable” flag on the file.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/lock(completionHandler:)-161qv
func (d NSDocument) LockWithCompletionHandler(completionHandler ErrorHandler) {
_block0, _cleanup0 := NewErrorBlock(completionHandler)
	defer _cleanup0()
	objc.Send[objc.ID](d.ID, objc.Sel("lockWithCompletionHandler:"), _block0)
}
// Allows the user to make modifications to the document.
//
// completionHandler: The completion handler block object passed in to be invoked after unlocking
// is completed, regardless of success or failure.
//
// # Discussion
// 
// By default, this method invokes the [UnlockWithCompletionHandler] method to
// unlock the document. This method disables autosaving safety checking,
// meaning that [CheckAutosavingSafetyAndReturnError] will no longer be
// invoked on this document. When unlocking succeeds, the [Locked] method will
// begin returning [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/unlock(completionHandler:)-8p8zd
func (d NSDocument) UnlockDocumentWithCompletionHandler(completionHandler BoolHandler) {
_block0, _cleanup0 := NewBoolBlock(completionHandler)
	defer _cleanup0()
	objc.Send[objc.ID](d.ID, objc.Sel("unlockDocumentWithCompletionHandler:"), _block0)
}
// Allows the user to make modifications to the document’s file.
//
// completionHandler: The completion handler block object passed in to be invoked after unlocking
// is completed, regardless of success or failure.
//
// # Discussion
// 
// By default, this method tries to clear the “user immutable” flag and,
// if necessary, add write permissions to the file itself.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/unlock(completionHandler:)-6m7rh
func (d NSDocument) UnlockWithCompletionHandler(completionHandler ErrorHandler) {
_block0, _cleanup0 := NewErrorBlock(completionHandler)
	defer _cleanup0()
	objc.Send[objc.ID](d.ID, objc.Sel("unlockWithCompletionHandler:"), _block0)
}
// Adds document-specific content to the Page Layout panel.
//
// pageLayout: The page layout panel to prepare.
//
// # Return Value
// 
// [true] if successfully prepared; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// The [runModalPageLayoutWithPrintInfo:] and
// [RunModalPageLayoutWithPrintInfoDelegateDidRunSelectorContextInfo] methods
// call this method to allow the document to customize the Page Layout panel
// `pageLayout`. You might use this method to add a document-related accessory
// view.
// 
// The default implementation returns [true].
//
// [runModalPageLayoutWithPrintInfo:]: https://developer.apple.com/documentation/AppKit/NSDocument/runModalPageLayoutWithPrintInfo:
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/preparePageLayout(_:)
func (d NSDocument) PreparePageLayout(pageLayout INSPageLayout) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("preparePageLayout:"), pageLayout)
	return rv
}
// Runs the modal page layout panel with the receiver’s printing information
// object.
//
// printInfo: The [NSPrintInfo] object for the page layout panel to use.
//
// delegate: The delegate to which the selector message is sent.
//
// didRunSelector: The selector of the message sent to the delegate.
//
// contextInfo: Object passed with the callback to provide any additional context
// information.
//
// # Discussion
// 
// Invoked from the action method [RunPageLayout]. Presents the page layout
// panel app modally if there is no document window to which it can be
// presented document modally.
// 
// When the panel is dismissed, `delegate` is sent a `didRunSelector` message.
// The `didRunSelector` callback method should have the following signature:
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/runModalPageLayout(with:delegate:didRun:contextInfo:)
func (d NSDocument) RunModalPageLayoutWithPrintInfoDelegateDidRunSelectorContextInfo(printInfo INSPrintInfo, delegate objectivec.IObject, didRunSelector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](d.ID, objc.Sel("runModalPageLayoutWithPrintInfo:delegate:didRunSelector:contextInfo:"), printInfo, delegate, didRunSelector, contextInfo)
}
// Runs the specified print operation modally.
//
// printOperation: The print operation to run.
//
// delegate: The delegate to which the selector message is sent.
//
// didRunSelector: The selector of the message sent to the delegate.
//
// contextInfo: Object passed with the callback to provide any additional context
// information.
//
// # Discussion
// 
// Overrides of [printShowingPrintPanel:] can invoke this method.
// 
// When the panel is dismissed, `delegate` is sent a `didRunSelector` message.
// Pass the `contextInfo` object with the callback. The `didRunSelector`
// callback method should have the following signature:
//
// [printShowingPrintPanel:]: https://developer.apple.com/documentation/AppKit/NSDocument/printShowingPrintPanel:
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/runModalPrintOperation(_:delegate:didRun:contextInfo:)
func (d NSDocument) RunModalPrintOperationDelegateDidRunSelectorContextInfo(printOperation INSPrintOperation, delegate objectivec.IObject, didRunSelector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](d.ID, objc.Sel("runModalPrintOperation:delegate:didRunSelector:contextInfo:"), printOperation, delegate, didRunSelector, contextInfo)
}
// Returns a Boolean value that indicates whether the document allows changes
// to the default printing information.
//
// newPrintInfo: The [NSPrintInfo] object that is the result of the user approving the page
// layout panel presented by [RunPageLayout].
//
// # Return Value
// 
// [true] by default; subclasses can override this method to return [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method is invoked by the [RunPageLayout] method, which sets a new
// [NSPrintInfo]object for the document only if this method returns [true].
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/shouldChangePrintInfo(_:)
func (d NSDocument) ShouldChangePrintInfo(newPrintInfo INSPrintInfo) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("shouldChangePrintInfo:"), newPrintInfo)
	return rv
}
// Prints the document’s contents, optionally displaying a print panel to
// the user.
//
// printSettings: The print settings dictionary to use.
//
// showPrintPanel: A Boolean value indicating whether the print panel is shown.
//
// delegate: The delegate to which the selector message is sent.
//
// didPrintSelector: The selector of the message sent to the delegate.
//
// contextInfo: Object passed with the callback to provide any additional context
// information.
//
// # Discussion
// 
// If showing of the print panel is specified by `showPrintPanel`, the method
// presents it first and prints only if the user approves the panel. The
// [NSPrintInfo] attributes in the passed-in `printSettings` dictionary are
// added to a copy of the document’s print info, and the resulting print
// info settings are used for the operation. When printing is complete or
// canceled, the method sends the message selected by `didPrintSelector` to
// the `delegate`, with the `contextInfo` as the last argument. The method
// selected by `didPrintSelector` must have the same signature as:
// 
// The default implementation of this method invokes
// [PrintOperationWithSettingsError]. If `nil` is returned it presents the
// error to the user in a document-modal panel before messaging the delegate.
// Otherwise it invokes `[thePrintOperation showPrintPanel]` then `[self
// thePrintOperation delegate didPrintSelector contextInfo]`.
// 
// For backward binary compatibility with OS X v10.3 and earlier, the default
// implementation of this method invokes [printShowingPrintPanel:] if it is
// overridden. When doing this it uses private functionality to arrange for
// the print settings to take effect (despite the fact that the override of
// [printShowingPrintPanel:] can’t possibly know about them) and to get
// notified when the print operation has been completed, so it can message the
// delegate at the correct time. Correct messaging of the delegate is
// necessary for correct handling of the Print Apple event.
//
// [printShowingPrintPanel:]: https://developer.apple.com/documentation/AppKit/NSDocument/printShowingPrintPanel:
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/print(withSettings:showPrintPanel:delegate:didPrint:contextInfo:)
func (d NSDocument) PrintDocumentWithSettingsShowPrintPanelDelegateDidPrintSelectorContextInfo(printSettings foundation.INSDictionary, showPrintPanel bool, delegate objectivec.IObject, didPrintSelector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](d.ID, objc.Sel("printDocumentWithSettings:showPrintPanel:delegate:didPrintSelector:contextInfo:"), printSettings, showPrintPanel, delegate, didPrintSelector, contextInfo)
}
// Creates and returns a print operation for the document’s contents.
//
// printSettings: The print settings dictionary to use.
//
// # Return Value
// 
// The print operation, or `nil` if unsuccessful.
//
// # Discussion
// 
// The print operation can be run to print the document’s current contents.
// The [NSPrintInfo] attributes in the passed-in `printSettings` dictionary
// are added to a copy of the document’s print info, and the resulting print
// info is used for the operation.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/printOperation(withSettings:)
func (d NSDocument) PrintOperationWithSettingsError(printSettings foundation.INSDictionary) (INSPrintOperation, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](d.ID, objc.Sel("printOperationWithSettings:error:"), printSettings, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return NSPrintOperation{}, foundation.NSErrorFrom(errorPtr)
	}
	return NSPrintOperationFromID(rv), nil

}
// Exports a PDF representation of the document’s current contents.
//
// sender: The control sending the message.
//
// # Discussion
// 
// This action method handles the File menu’s “Export As PDF…” item in
// a document-based application.
// 
// The default implementation of this method calls the
// [PrintDocumentWithSettingsShowPrintPanelDelegateDidPrintSelectorContextInfo]
// method, passing a print settings object that contains only the disposition
// ([save]), with user interaction disabled and [NULL] or `nil` for all other
// parameters.
// 
// If your document subclass supports creating PDF representations, you can
// override this method, if needed.
//
// [save]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/JobDisposition-swift.struct/save
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/saveToPDF(_:)
func (d NSDocument) SaveDocumentToPDF(sender objectivec.IObject) {
	objc.Send[objc.ID](d.ID, objc.Sel("saveDocumentToPDF:"), sender)
}
// Perform any custom setup associated with a sharing service picker.
//
// sharingServicePicker: The sharing service picker the system is about to display.
//
// # Discussion
// 
// Override this method, as needed, and use it to configure the sharing
// service picker before AppKit displays it. The default implementation of
// this method does nothing. You might customize the contents of the share
// menu or provide a custom delegate for the chosen sharing service. You can
// get the default sharing menu item by calling [StandardShareMenuItem] on the
// current document controller.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/prepare(_:)
func (d NSDocument) PrepareSharingServicePicker(sharingServicePicker INSSharingServicePicker) {
	objc.Send[objc.ID](d.ID, objc.Sel("prepareSharingServicePicker:"), sharingServicePicker)
}
// Share the document’s file using the specified sharing service.
//
// sharingService: The sharing service that is sharing the document.
//
// completionHandler: The completion handler block to call with the results. AppKit calls this
// method after attempting to share the document. The handler block has no
// return value and takes the following parameter:
// 
// success: A Boolean value indicating whether the sharing attempt was
// successful. The value is [true] if sharing succeeded, or [false] if it
// didn’t.
// //
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method shares the document’s file, located at [FileURL], with the
// specified sharing service. If the document has never been saved, this
// method prompts the user to save it before proceeding; otherwise, the method
// initiates an autosave operation to save any outstanding changes. As needed,
// the method moves the file to an appropriate location for sharing. For
// example, when saving to iCloud, the method moves the file to the user’s
// iCloud Drive folder. After the sharing service finishes, the method calls
// `completionHandler` with the results.
// 
// The document temporarily replaces the current delegate of `sharingService`
// with itself. However, the document retains a reference to the old delegate
// object and forwards all delegate method calls to it.
// 
// If an [NSDocument] object is the only item associated with an
// [NSSharingServicePicker] or [NSSharingServicePickerTouchBarItem], the
// picker automatically calls this method directly, instead of sharing the
// items using the [PerformWithItems] method.
// 
// If you override this method, you must ensure that the document is in the
// proper state before attempting to share it. For example, you are
// responsible for ensuring the file is in an appropriate location. When your
// sharing service finishes, you must execute the provided `completionHandler`
// block with the results.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/share(with:completionHandler:)
func (d NSDocument) ShareDocumentWithSharingServiceCompletionHandler(sharingService INSSharingService, completionHandler BoolHandler) {
_block1, _cleanup1 := NewBoolBlock(completionHandler)
	defer _cleanup1()
	objc.Send[objc.ID](d.ID, objc.Sel("shareDocumentWithSharingService:completionHandler:"), sharingService, _block1)
}
// Handles the Close AppleScript command by attempting to close the document.
//
// command: A Close AppleScript command object.
//
// # Discussion
// 
// Extracts Close command arguments from the `command` object and uses them to
// determine how to close the document—specifically, whether to ignore
// unsaved changes, save changes automatically, or ask the user and to
// identify the file in which to save the document (by default, the file that
// was opened or previously saved to). A Close AppleScript command may specify
// more than one document to close. If so, a message is sent to each document
// object.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/handleClose(_:)
func (d NSDocument) HandleCloseScriptCommand(command foundation.NSCloseCommand) objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("handleCloseScriptCommand:"), command)
	return objectivec.Object{ID: rv}
}
// Handles the Print AppleScript command by attempting to print the document.
//
// command: An AppleScript command object.
//
// # Discussion
// 
// Extracts Print command arguments from the `command` object and uses them to
// determine how to print the document—specifically, any print settings and
// whether to show the Print dialog. A Print AppleScript command may specify
// more than one document to print. If so, a message is sent to each document.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/handlePrint(_:)
func (d NSDocument) HandlePrintScriptCommand(command foundation.NSScriptCommand) objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("handlePrintScriptCommand:"), command)
	return objectivec.Object{ID: rv}
}
// Handles the Save AppleScript command by attempting to save the document.
//
// command: An AppleScript command object.
//
// # Discussion
// 
// Extracts Save command arguments from the `command` object and uses them to
// determine the file in which to save the document and the file type.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/handleSave(_:)
func (d NSDocument) HandleSaveScriptCommand(command foundation.NSScriptCommand) objectivec.IObject {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("handleSaveScriptCommand:"), command)
	return objectivec.Object{ID: rv}
}
// Presents an error alert to the user as a modal panel.
//
// error: The error object encapsulating the information to present to the user.
//
// window: The window to which the modal alert belongs.
//
// delegate: The delegate to which the selector message is sent.
//
// didPresentSelector: The selector of the message sent to the delegate.
//
// contextInfo: Object passed with the callback to provide any additional context
// information.
//
// # Discussion
// 
// When the user dismisses the alert and any recovery possible for the error
// and chosen by the user has been attempted, sends the message
// `didPresentSelector` to the specified `delegate`.
// 
// The [NSDocument] default implementation of this method is equivalent to
// that of [NSResponder] and treats the shared [NSDocumentController] object
// as the next responder and forwards these messages to it. The default
// implementations of several [NSDocument] methods invoke this method.
// 
// The default implementation of this method invokes [WillPresentError] to
// give subclasses an opportunity to customize error presentation. You should
// not override this method but should instead override [WillPresentError].
// 
// The method selected by `didPresentSelector` must have the same signature
// as:
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/presentError(_:modalFor:delegate:didPresent:contextInfo:)
func (d NSDocument) PresentErrorModalForWindowDelegateDidPresentSelectorContextInfo(error_ foundation.INSError, window INSWindow, delegate objectivec.IObject, didPresentSelector objc.SEL, contextInfo uintptr) {
	objc.Send[objc.ID](d.ID, objc.Sel("presentError:modalForWindow:delegate:didPresentSelector:contextInfo:"), error_, window, delegate, didPresentSelector, contextInfo)
}
// Presents an error alert to the user as a modal panel.
//
// error: The error object encapsulating the information to present to the user.
//
// # Return Value
// 
// [true] if error recovery was done; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// This method does not return until the user dismisses the alert and, if the
// error has recovery options and a recovery delegate, the error’s recovery
// delegate is sent an [attemptRecovery(fromError:optionIndex:)] message.
// 
// The [NSDocument] default implementation of this method is equivalent to
// that of [NSResponder] and treats the shared [NSDocumentController] as the
// next responder and forwards these messages to it.
// 
// The default implementation of this method invokes [WillPresentError] to
// give subclasses an opportunity to customize error presentation. You should
// not override this method but should instead override [WillPresentError].
//
// [attemptRecovery(fromError:optionIndex:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/attemptRecovery(fromError:optionIndex:)
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/presentError(_:)
func (d NSDocument) PresentError(error_ foundation.INSError) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("presentError:"), error_)
	return rv
}
// Called when the receiver is about to present an error.
//
// error: The error object that is about to be presented to the user.
//
// # Return Value
// 
// The error that should actually be presented.
//
// # Discussion
// 
// The default implementation of this method merely returns the passed-in
// error. The returned error may simply be forwarded to the document
// controller.
// 
// You can override this method to customize the presentation of errors by
// examining the passed-in error and, for example, returning more specific
// information. When you override this method always check the [NSError]
// object’s domain and code to discriminate between errors whose
// presentation you want to customize and those you don’t. For errors you
// don’t want to customize, call the superclass implementation, passing the
// original error.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/willPresentError(_:)
func (d NSDocument) WillPresentError(error_ foundation.INSError) foundation.INSError {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("willPresentError:"), error_)
	return foundation.NSErrorFromID(rv)
}
// Confirms that the error object is not to be presented to the user and the
// error cannot be recovered from, so cleanup can be done.
//
// error: An [NSError] object returned by another [NSDocument] method.
// //
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
//
// # Discussion
// 
// Some [NSDocument] methods, like those involved in writing, may not
// immediately delete temporary files if there is a chance that the error can
// be recovered from and the operation can continue. To make sure that cleanup
// is always done, you should invoke this method with [NSDocument] errors that
// are not going to be passed to one of the `...` methods. For example, the
// [NSDocument] implementation of the [NSFilePresenter] method
// [savePresentedItemChanges(completionHandler:)] invokes this method when it
// invokes [AutosaveWithImplicitCancellabilityCompletionHandler] and the
// completion handler is passed an [NSError] object, because it does not
// present the error to the user.
//
// [savePresentedItemChanges(completionHandler:)]: https://developer.apple.com/documentation/Foundation/NSFilePresenter/savePresentedItemChanges(completionHandler:)
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/willNotPresentError(_:)
func (d NSDocument) WillNotPresentError(error_ foundation.INSError) {
	objc.Send[objc.ID](d.ID, objc.Sel("willNotPresentError:"), error_)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/accommodatePresentedItemDeletion(completionHandler:)
func (d NSDocument) AccommodatePresentedItemDeletionWithCompletionHandler(completionHandler ErrorHandler) {
_block0, _cleanup0 := NewErrorBlock(completionHandler)
	defer _cleanup0()
	objc.Send[objc.ID](d.ID, objc.Sel("accommodatePresentedItemDeletionWithCompletionHandler:"), _block0)
}
// See: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemDidChange()
func (d NSDocument) PresentedItemDidChange() {
	objc.Send[objc.ID](d.ID, objc.Sel("presentedItemDidChange"))
}
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemDidChangeUbiquityAttributes(_:)
func (d NSDocument) PresentedItemDidChangeUbiquityAttributes(attributes foundation.INSSet) {
	objc.Send[objc.ID](d.ID, objc.Sel("presentedItemDidChangeUbiquityAttributes:"), attributes)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemDidGain(_:)
func (d NSDocument) PresentedItemDidGainVersion(version foundation.NSFileVersion) {
	objc.Send[objc.ID](d.ID, objc.Sel("presentedItemDidGainVersion:"), version)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemDidLose(_:)
func (d NSDocument) PresentedItemDidLoseVersion(version foundation.NSFileVersion) {
	objc.Send[objc.ID](d.ID, objc.Sel("presentedItemDidLoseVersion:"), version)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemDidMove(to:)
func (d NSDocument) PresentedItemDidMoveToURL(newURL foundation.INSURL) {
	objc.Send[objc.ID](d.ID, objc.Sel("presentedItemDidMoveToURL:"), newURL)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemDidResolveConflict(_:)
func (d NSDocument) PresentedItemDidResolveConflictVersion(version foundation.NSFileVersion) {
	objc.Send[objc.ID](d.ID, objc.Sel("presentedItemDidResolveConflictVersion:"), version)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/relinquishPresentedItem(toReader:)
func (d NSDocument) RelinquishPresentedItemToReader(reader VoidHandler) {
_block0, _cleanup0 := NewVoidBlock(reader)
	defer _cleanup0()
	objc.Send[objc.ID](d.ID, objc.Sel("relinquishPresentedItemToReader:"), _block0)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/relinquishPresentedItem(toWriter:)
func (d NSDocument) RelinquishPresentedItemToWriter(writer VoidHandler) {
_block0, _cleanup0 := NewVoidBlock(writer)
	defer _cleanup0()
	objc.Send[objc.ID](d.ID, objc.Sel("relinquishPresentedItemToWriter:"), _block0)
}
//
// # Discussion
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/savePresentedItemChanges(completionHandler:)
func (d NSDocument) SavePresentedItemChangesWithCompletionHandler(completionHandler ErrorHandler) {
_block0, _cleanup0 := NewErrorBlock(completionHandler)
	defer _cleanup0()
	objc.Send[objc.ID](d.ID, objc.Sel("savePresentedItemChangesWithCompletionHandler:"), _block0)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSEditorRegistration/objectDidBeginEditing(_:)
func (d NSDocument) ObjectDidBeginEditing(editor NSEditor) {
	objc.Send[objc.ID](d.ID, objc.Sel("objectDidBeginEditing:"), editor)
}
//
// See: https://developer.apple.com/documentation/AppKit/NSEditorRegistration/objectDidEndEditing(_:)
func (d NSDocument) ObjectDidEndEditing(editor NSEditor) {
	objc.Send[objc.ID](d.ID, objc.Sel("objectDidEndEditing:"), editor)
}
// Restores the state necessary to continue the specified user activity.
//
// userActivity: The user activity to continue.
//
// # Discussion
// 
// Implement this method to restore an object’s state using the specified
// user activity. The system calls this method on any responders or documents
// passed to the `restorationHandler` in
// [ApplicationContinueUserActivityRestorationHandler]. The system calls this
// method on the main thread. Your implementation should use the state data
// contained in the specified user activity’s [userInfo] dictionary to
// restore the object.
// 
// On macOS, the system can automatically restore activities managed by
// [NSDocument] if you don’t implement
// [ApplicationContinueUserActivityRestorationHandler], or if you return
// [false]. When this occurs, the system opens the document using
// [OpenDocumentWithContentsOfURLDisplayCompletionHandler], and calls
// `restoreUserActivityState` on it.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [userInfo]: https://developer.apple.com/documentation/Foundation/NSUserActivity/userInfo
//
// See: https://developer.apple.com/documentation/AppKit/NSUserActivityRestoring/restoreUserActivityState(_:)
func (d NSDocument) RestoreUserActivityState(userActivity foundation.NSUserActivity) {
	objc.Send[objc.ID](d.ID, objc.Sel("restoreUserActivityState:"), userActivity)
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
func (d NSDocument) ValidateMenuItem(menuItem INSMenuItem) bool {
	rv := objc.Send[bool](d.ID, objc.Sel("validateMenuItem:"), menuItem)
	return rv
}

// Returns a Boolean value that indicates whether the receiver reads multiple
// documents of the given type concurrently.
//
// typeName: The string that identifies the document type.
//
// # Return Value
// 
// [false] by default; subclasses can override to return [true], thereby
// causing documents of the specified type to be read concurrently.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// Your [NSDocument] subclass can implement this method to return [true] to
// enable loading of documents concurrently, using background threads. When
// this facility is enabled in this way, [InitWithContentsOfURLOfTypeError]
// executes on a background thread when opening files via the Open panel or
// from the Finder. This allows concurrent reading of multiple documents and
// also allows the app to be responsive while reading a large document.
// 
// The default implementation of this method returns [false]. A subclass
// override should return [true] only for document types whose reading is
// thread-safe, as described in [NSDocument]. You should disable undo
// registration during document reading, which is a good idea even in the
// absence of concurrency.
// 
// If you are checking the current Apple Event for a search string, you should
// not enable concurrent document opening, because code handling a document
// opening triggered by an Apple Event cannot get the current Apple Event.
// This happens because the event is suspended until all documents are read to
// enable correct reporting of success or error.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/canConcurrentlyReadDocuments(ofType:)
func (_NSDocumentClass NSDocumentClass) CanConcurrentlyReadDocumentsOfType(typeName string) bool {
	rv := objc.Send[bool](objc.ID(_NSDocumentClass.class), objc.Sel("canConcurrentlyReadDocumentsOfType:"), objc.String(typeName))
	return rv
}
// Returns a Boolean value that indicates whether the document can read and
// write the data natively.
//
// type: The string that identifies the document type to test.
//
// # Return Value
// 
// [true] if the document type is a native type; otherwise, [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/isNativeType(_:)
func (_NSDocumentClass NSDocumentClass) IsNativeType(type_ string) bool {
	rv := objc.Send[bool](objc.ID(_NSDocumentClass.class), objc.Sel("isNativeType:"), objc.String(type_))
	return rv
}
// Returns the classes that support secure coding.
//
// keyPath: The key path of the restorable object.
//
// # Return Value
// 
// An array of classes that support secure coding.
//
// # Discussion
// 
// The system calls the function during a secure state restoration and
// restores values only for the allowed classes your app returns in the array.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/allowedClasses(forRestorableStateKeyPath:)
func (_NSDocumentClass NSDocumentClass) AllowedClassesForRestorableStateKeyPath(keyPath string) []objc.Class {
	rv := objc.Send[[]objc.ID](objc.ID(_NSDocumentClass.class), objc.Sel("allowedClassesForRestorableStateKeyPath:"), objc.String(keyPath))
	return objc.ConvertSlice(rv, func(id objc.ID) objc.Class {
		return objc.Class(id)
	})
}

// The location of the document’s on-disk representation.
//
// # Return Value
// 
// The document’s location.
// 
// # Discussion
// 
// The default implementation of this property returns the URL of the file
// that was opened. Changing the value of this property does not actually
// change the document’s name or location; it is only for recording the
// document’s location during its initial opening or saving.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/fileURL
func (d NSDocument) FileURL() foundation.INSURL {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("fileURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (d NSDocument) SetFileURL(value foundation.INSURL) {
	objc.Send[struct{}](d.ID, objc.Sel("setFileURL:"), value)
}
// A Boolean value that indicates whether the document’s file is completely
// loaded into memory.
//
// # Return Value
// 
// [true] if the document’s entire file is loaded into memory; otherwise
// [false].
// 
// # Discussion
// 
// The default value of this property is [true], which signifies that the
// entire file is loaded into memory. You can override this property to return
// [false] if additional data needs to be read from the file. [NSDocument] may
// use this value to do things like prevent volume ejection or warn the user
// when a partially loaded file disappears from the file system.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/isEntireFileLoaded
func (d NSDocument) EntireFileLoaded() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isEntireFileLoaded"))
	return rv
}
// The last-known modification date of the document’s on-disk
// representation.
//
// # Discussion
// 
// The [NSDocument] default file saving machinery uses this information to
// warn the user when the on-disk representation of an open document has been
// modified by something other than the current app.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/fileModificationDate
func (d NSDocument) FileModificationDate() foundation.INSDate {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("fileModificationDate"))
	return foundation.NSDateFromID(objc.ID(rv))
}
func (d NSDocument) SetFileModificationDate(value foundation.INSDate) {
	objc.Send[struct{}](d.ID, objc.Sel("setFileModificationDate:"), value)
}
// A Boolean value that indicates whether the document archives previously
// saved versions of the document.
//
// # Discussion
// 
// The default value of this property is [false], which causes each new save
// operation to replace the document’s on-disk content. If you override this
// method and return [true], a save operation saves the document’s previous
// contents in a backup file before saving the current contents.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/keepBackupFile
func (d NSDocument) KeepBackupFile() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("keepBackupFile"))
	return rv
}
// A Boolean value that indicates whether the document is a draft that the
// user has not yet saved.
//
// # Discussion
// 
// The system presents a Save dialog when the user closes a draft document.
// Only documents with non-`nil` values for the [FileURL] property should be
// considered drafts.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/isDraft
func (d NSDocument) Draft() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isDraft"))
	return rv
}
func (d NSDocument) SetDraft(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setDraft:"), value)
}
// The name of the document type, as specified in the app’s information
// property-list file.
//
// # Discussion
// 
// The document type affects how the data is filtered when it is written to or
// read from a file. When a document is saved, the type is determined by the
// entries in the app’s information property list (specified in
// `Info.Plist())`.
// 
// You cannot use this property to change the document’s format after it has
// already been opened or saved. This property records only the initial
// document format used when first opening or saving the file.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/fileType
func (d NSDocument) FileType() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("fileType"))
	return foundation.NSStringFromID(rv).String()
}
func (d NSDocument) SetFileType(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setFileType:"), objc.String(value))
}
// A Boolean value that indicates whether the document has unsaved changes.
//
// # Discussion
// 
// The value of this property is [true] if the document has been edited. The
// edited status of each document window reflects the document’s edited
// status.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/isDocumentEdited
func (d NSDocument) DocumentEdited() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isDocumentEdited"))
	return rv
}
// A Boolean value that indicates whether the document is in read-only mode.
//
// # Discussion
// 
// The value of this property is [true] if the document is in read-only
// “viewing mode,” that is, if the document is locked. You can use this
// information to prevent certain kinds of user actions or changes when the
// user is viewing an old document revision.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/isInViewingMode
func (d NSDocument) InViewingMode() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isInViewingMode"))
	return rv
}
// The document’s current window controllers.
//
// # Discussion
// 
// The value of this property is an array of [NSWindowController] objects
// belonging to the current document. If there are no window controllers, the
// value is an empty array object.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/windowControllers
func (d NSDocument) WindowControllers() []NSWindowController {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("windowControllers"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSWindowController {
		return NSWindowControllerFromID(id)
	})
}
// The name of the document’s sole nib file.
//
// # Discussion
// 
// Using this name, [NSDocument] creates and instantiates a default instance
// of [NSWindowController] to manage the window. If your document has multiple
// nib files, each with its own single window, or if the default
// [NSWindowController] instance is not adequate for your purposes, you should
// override [WindowControllers].
// 
// The default value of this property is `nil`. Subclasses must override it to
// specify a nib file name.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/windowNibName
func (d NSDocument) WindowNibName() NSNibName {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("windowNibName"))
	return NSNibName(foundation.NSStringFromID(rv).String())
}
// Returns the document window to use as the parent of a document-modal sheet.
//
// # Discussion
// 
// This method searches the document’s window controllers for the most
// suitable window to use when displaying the sheet.
// 
// The value of this property may be `nil`, in which case the sender should
// present an app-modal panel. The [NSDocument] implementation of this
// property sets the value to the window of the first window controller, or
// `[NSApp mainWindow]` if there are no window controllers or if the first
// window controller has no window.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/windowForSheet
func (d NSDocument) WindowForSheet() INSWindow {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("windowForSheet"))
	return NSWindowFromID(objc.ID(rv))
}
// The name of the document as displayed in the title bars of the document’s
// windows and in alert dialogs related to the document.
//
// # Discussion
// 
// If the document has been saved, the display name is the last component of
// the directory location of the saved file (for example, “[MyDocument]”
// if the path is “`/tmp/MyDocument.Rtf()`”). If the document is new,
// [NSDocument] makes the display name “Untitled n,” where n is a number
// in a sequence of new and unsaved documents. The displayable name also takes
// into account whether the document’s filename extension should be hidden.
// Subclasses of [NSWindowController] can override
// [WindowTitleForDocumentDisplayName] to modify the display name as it
// appears in window titles.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/displayName
func (d NSDocument) DisplayName() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("displayName"))
	return foundation.NSStringFromID(rv).String()
}
func (d NSDocument) SetDisplayName(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setDisplayName:"), objc.String(value))
}
// The location of the most recently autosaved document contents.
//
// # Discussion
// 
// Use this property to specify the location where you want the document to
// store autosave files. The URL you specify should specify an absolute path,
// not a relative path.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/autosavedContentsFileURL
func (d NSDocument) AutosavedContentsFileURL() foundation.INSURL {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("autosavedContentsFileURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
func (d NSDocument) SetAutosavedContentsFileURL(value foundation.INSURL) {
	objc.Send[struct{}](d.ID, objc.Sel("setAutosavedContentsFileURL:"), value)
}
// The document type to use for an autosave operation.
//
// # Discussion
// 
// This properties contains a string that identifies the document type for
// autosave files. The default implementation just returns the value provided
// by the [FileType] property. You can override this property and return `nil`
// to completely disable autosaving of individual documents (because the
// [NSDocumentController] object does not call the
// [AutosaveDocumentWithDelegateDidAutosaveSelectorContextInfo] method of a
// document that has no autosaving file type). You can also override it if
// your app defines a document type that is specifically designed for
// autosaving, for example, one that efficiently represents document content
// changes instead of complete document contents.
// 
// Overriding this property can result in incorrect behavior during reopening
// of autosaved documents. The [NSDocument] method
// [InitForURLWithContentsOfURLOfTypeError], which is invoked during reopening
// of autosaved documents after a crash, takes two URLs, but only the type
// name of the autosaved contents file. The default implementation updates the
// [FileType] property with that type name, but that may not be the right
// thing to do if this property contains something other than [FileType]
// during document autosaving. If you override `autosavingFileType`, you
// probably need to override [InitForURLWithContentsOfURLOfTypeError] too, and
// make the override update [FileType] with the type of the actual document
// file, after invoking `super`. See TextEdit’s [Document] class for an
// example of how to do this.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/autosavingFileType
func (d NSDocument) AutosavingFileType() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("autosavingFileType"))
	return foundation.NSStringFromID(rv).String()
}
// A Boolean value that indicates whether you can cancel an in-progress
// autosave operation.
//
// # Discussion
// 
// The value of this property is [true] if autosaving is in progress but
// nothing bad would happen if it were cancelled. For example, when periodic
// autosaving is being done only for crash protection, which doesn’t need to
// be done all of the time, this property is set to [true]. When autosaving is
// being done because the document is being closed, the property is set to
// [false].
// 
// When the value is [true], your document-writing code can invoke
// [UnblockUserInteraction] after recording the fact that changes to the
// document model made by the user should first cancel the rest of the
// writing. Your code that makes changes to the document model then must
// always do that cancellation first. If your writing code is implicitly
// cancelled in this way, it should set the [NSError] object passed by
// reference to the writing method to [NSUserCancelledError] in
// [NSCocoaErrorDomain].
//
// [NSCocoaErrorDomain]: https://developer.apple.com/documentation/Foundation/NSCocoaErrorDomain
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
// [NSUserCancelledError]: https://developer.apple.com/documentation/Foundation/NSUserCancelledError-swift.var
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/autosavingIsImplicitlyCancellable
func (d NSDocument) AutosavingIsImplicitlyCancellable() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("autosavingIsImplicitlyCancellable"))
	return rv
}
// A Boolean value that indicates whether the document has changes that have
// not been autosaved.
//
// # Discussion
// 
// The value of this property is [true] if the document has changes that have
// not been autosaved; otherwise, the value is [false]. A document has unsaved
// changes when the [UpdateChangeCount] method has been called since the last
// save.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/hasUnautosavedChanges
func (d NSDocument) HasUnautosavedChanges() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("hasUnautosavedChanges"))
	return rv
}
// The URL for the document’s backup file that was created during an
// autosave operation.
//
// # Discussion
// 
// This property specifies the location of the backup file, if any. If a
// backup file cannot be created or is not needed, the value of this property
// is `nil`.
// 
// Starting in OS X v10.8, document versions can be preserved using a backup
// file created during an autosave operation, which supports document
// versioning. This property gives you access to the backup file’s URL.
// 
// Using an autosave backup file for preserving versions is efficient. This is
// because an [NSDocument] instance is able to use the [byMoving] option when
// it calls the [replaceItem(at:options:)] method. The document gets the value
// of this property twice during saving:
// 
// - Before calling the [WriteSafelyToURLOfTypeForSaveOperationError] method:
// This is to check whether using the replace-by-moving option is possible
// and, if not, to allow the system to preserve data by instead using copying.
// - Within the [WriteSafelyToURLOfTypeForSaveOperationError] method: This is
// to discover where to put the backup file.
// 
// When you implement the [WriteSafelyToURLOfTypeForSaveOperationError] method
// with the [NSDocument.SaveOperationType.saveOperation] or
// [NSDocument.SaveOperationType.autosaveInPlaceOperation] operation type, you
// must check this property’s value. If it is not `nil`, move the previous
// contents of the file (that would be overwritten) to the URL’s location.
// The default implementation of [WriteSafelyToURLOfTypeForSaveOperationError]
// does this.
// 
// To create a backup file from within your custom implementation of the
// [WriteSafelyToURLOfTypeForSaveOperationError] method, call the
// [FileManager] method
// [replaceItem(at:withItemAt:backupItemName:options:resultingItemURL:)],
// using a backup item name of `[[self backupFileURL] lastPathComponent]` and
// an option of [withoutDeletingBackupItem] option. If your custom
// implementation is unable to keep the backup file, you must override this
// property and return `nil` to ensure that the document’s file gets
// correctly preserved before it gets overwritten.
// 
// The default implementation of the
// [WriteSafelyToURLOfTypeForSaveOperationError] method returns a non-`nil`
// value based on the value of `[self fileURL]`, but only if the document’s
// file needs to be preserved prior to saving or if the [PreservesVersions]
// method returns [false]. Otherwise, it returns `nil`.
//
// [FileManager]: https://developer.apple.com/documentation/Foundation/FileManager
// [NSDocument.SaveOperationType.autosaveInPlaceOperation]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/autosaveInPlaceOperation
// [NSDocument.SaveOperationType.saveOperation]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/saveOperation
// [byMoving]: https://developer.apple.com/documentation/Foundation/NSFileVersion/ReplacingOptions/byMoving
// [false]: https://developer.apple.com/documentation/Swift/false
// [replaceItem(at:options:)]: https://developer.apple.com/documentation/Foundation/NSFileVersion/replaceItem(at:options:)
// [replaceItem(at:withItemAt:backupItemName:options:resultingItemURL:)]: https://developer.apple.com/documentation/Foundation/FileManager/replaceItem(at:withItemAt:backupItemName:options:resultingItemURL:)
// [withoutDeletingBackupItem]: https://developer.apple.com/documentation/Foundation/FileManager/ItemReplacementOptions/withoutDeletingBackupItem
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/backupFileURL
func (d NSDocument) BackupFileURL() foundation.INSURL {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("backupFileURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the document is currently displaying
// the Versions browser.
//
// # Discussion
// 
// The value of this property is [true] when the versions browser is visible.
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/isBrowsingVersions
func (d NSDocument) BrowsingVersions() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isBrowsingVersions"))
	return rv
}
// The object that the document uses to support undo/redo operations.
//
// # Discussion
// 
// When the [HasUndoManager] property is [true], accessing this property
// creates an [UndoManager] before returning it. If the [HasUndoManager]
// property is [false], the value of this property is `nil` by default.
// Assigning an undo manager to this property stores a reference to the object
// and automatically changes the [HasUndoManager] property to [true].
// 
// Whether you assign an undo manager or let the document create one, the
// document registers itself as an observer of various [UndoManager]
// notifications so that appropriate document actions are stored on the undo
// stack.
//
// [UndoManager]: https://developer.apple.com/documentation/Foundation/UndoManager
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/undoManager
func (d NSDocument) UndoManager() foundation.NSUndoManager {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("undoManager"))
	return foundation.NSUndoManagerFromID(objc.ID(rv))
}
func (d NSDocument) SetUndoManager(value foundation.NSUndoManager) {
	objc.Send[struct{}](d.ID, objc.Sel("setUndoManager:"), value)
}
// A Boolean value that indicates whether the document owns an undo manager
// object.
//
// # Discussion
// 
// If you change the value of this property to NO and the document already
// owns an [UndoManager] object, the document removes the undo manager as an
// observer of undo-related notifications and then removes its reference to
// the object.
//
// [UndoManager]: https://developer.apple.com/documentation/Foundation/UndoManager
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/hasUndoManager
func (d NSDocument) HasUndoManager() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("hasUndoManager"))
	return rv
}
func (d NSDocument) SetHasUndoManager(value bool) {
	objc.Send[struct{}](d.ID, objc.Sel("setHasUndoManager:"), value)
}
// A Boolean value that indicates whether the document’s Save panel displays
// a list of supported writable document types.
//
// # Discussion
// 
// When the value of this property is [true], the document includes a pop-up
// menu of supported writable document types when it displays the Save panel.
// The default value of this property is [true]. Subclasses can override this
// property to provide a different value. For example, you might provide the
// following implementation:
//
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/shouldRunSavePanelWithAccessoryView
func (d NSDocument) ShouldRunSavePanelWithAccessoryView() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("shouldRunSavePanelWithAccessoryView"))
	return rv
}
// The file type that was last selected in the Save panel.
//
// # Discussion
// 
// This type is primarily used by the [SaveDocument], [SaveDocumentAs], and
// [SaveDocumentTo] methods to determine the type the user chose after the
// Save panel has been run. The string corresponds to the name of the document
// type as it is specified in the app’s `Info.Plist()` file.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/fileTypeFromLastRunSavePanel
func (d NSDocument) FileTypeFromLastRunSavePanel() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("fileTypeFromLastRunSavePanel"))
	return foundation.NSStringFromID(rv).String()
}
// A Boolean value that indicates whether the user chose to hide the
// document’s filename extension.
//
// # Return Value
// 
// [true] if a Save panel was presented and the user chose to hide the
// extension; otherwise, [false].
// 
// # Discussion
// 
// The Save panel includes an option to hide a document’s filename
// extension. If the user selects this option, AppKit sets the value of this
// property to [true]; otherwise, the value is [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/fileNameExtensionWasHiddenInLastRunSavePanel
func (d NSDocument) FileNameExtensionWasHiddenInLastRunSavePanel() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("fileNameExtensionWasHiddenInLastRunSavePanel"))
	return rv
}
// An object that encapsulates a user activity the document supports.
//
// # Discussion
// 
// [NSDocument] automatically creates [NSUserActivity] objects. The system
// makes user activities eligible for Handoff if the document is iCloud-based
// and the app’s `Info.Plist()` property list file includes a
// [CFBundleDocumentTypes] key of [NSUbiquitousDocumentUserActivityType]. The
// value of [NSUbiquitousDocumentUserActivityType] is a string that represents
// the [NSUserActivity] object’s activity type. The document’s URL is in
// the [NSUserActivity] object’s [userInfo] dictionary with the
// [NSUserActivityDocumentURLKey].
// 
// In macOS, [NSUserActivity] objects that [NSDocument] manage automatically
// become current when any of the document window controller’s windows
// become main. Otherwise, you need to invoke [becomeCurrent()] at appropriate
// times.
// 
// AppKit automatically invalidates any [NSUserActivity] objects that have no
// associated documents or responders.
// 
// You can use this property from any thread. It’s KVO-observable in case
// you share the [NSDocument] object with other objects that need to be in
// sync as the document moves into and out of iCloud.
//
// [NSUserActivityDocumentURLKey]: https://developer.apple.com/documentation/AppKit/NSUserActivityDocumentURLKey
// [NSUserActivity]: https://developer.apple.com/documentation/Foundation/NSUserActivity
// [becomeCurrent()]: https://developer.apple.com/documentation/Foundation/NSUserActivity/becomeCurrent()
// [userInfo]: https://developer.apple.com/documentation/Foundation/NSUserActivity/userInfo
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/userActivity
func (d NSDocument) UserActivity() foundation.NSUserActivity {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("userActivity"))
	return foundation.NSUserActivityFromID(objc.ID(rv))
}
func (d NSDocument) SetUserActivity(value foundation.NSUserActivity) {
	objc.Send[struct{}](d.ID, objc.Sel("setUserActivity:"), value)
}
// The key that identifies the document associated with a user activity.
//
// See: https://developer.apple.com/documentation/appkit/nsuseractivitydocumenturlkey
func (d NSDocument) NSUserActivityDocumentURLKey() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("NSUserActivityDocumentURLKey"))
	return foundation.NSStringFromID(rv).String()
}
// A Boolean value that indicates whether or not the file can be written to.
//
// # Discussion
// 
// This property may contain the value [true] because the user lacks the
// appropriate write permissions, the “user immutable” flag was raised,
// the parent directory or volume is read only, or the
// [CheckAutosavingSafetyAndReturnError] method returned [false]. Do not
// override this property.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/isLocked
func (d NSDocument) Locked() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("isLocked"))
	return rv
}
// The printing information associated with the document.
//
// # Return Value
// 
// The receiver’s [NSPrintInfo] object.
// 
// # Discussion
// 
// The default value of this property is the default [NSPrintInfo] object. To
// customize the printing information, assign a new value to this property.
// The Page Layout panel may also update the object in this property to
// reflect the options selected by the user.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/printInfo
func (d NSDocument) PrintInfo() INSPrintInfo {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("printInfo"))
	return NSPrintInfoFromID(objc.ID(rv))
}
func (d NSDocument) SetPrintInfo(value INSPrintInfo) {
	objc.Send[struct{}](d.ID, objc.Sel("setPrintInfo:"), value)
}
// A print operation you can use to create a PDF representation of the
// document’s current contents.
//
// # Discussion
// 
// The object in this property can be run to print the document’s current
// contents to a PDF file.
// 
// The default print operation stored by this property is obtained by calling
// the [PrintOperationWithSettingsError] method and passing a print settings
// object that contains only the disposition ([save]) and a [NULL] error
// object reference. If your document subclass supports creating PDF
// representations, you can override this property as needed to customize the
// options.
//
// [save]: https://developer.apple.com/documentation/AppKit/NSPrintInfo/JobDisposition-swift.struct/save
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/pdfPrintOperation
func (d NSDocument) PDFPrintOperation() INSPrintOperation {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("PDFPrintOperation"))
	return NSPrintOperationFromID(objc.ID(rv))
}
// A Boolean value that indicates whether the document is shareable from the
// standard Share menu.
//
// # Discussion
// 
// When the value of this property is [true], the owning document controller
// enables share options for this document. When the value is [false], the
// document controller disables the Share menu when this document is selected.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/allowsDocumentSharing
func (d NSDocument) AllowsDocumentSharing() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("allowsDocumentSharing"))
	return rv
}
// Returns the object specifier that represents the document.
//
// # Return Value
// 
// The document object specifier.
// 
// # Discussion
// 
// An object specifier represents an AppleScript reference form, which is a
// natural-language expression such as `words 10 through 20` or `front
// document`. During script processing, an object contained by a document
// (such as the `second paragraph` or the `third rectangle`) may need to
// specify its container (the document).
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/objectSpecifier
func (d NSDocument) ObjectSpecifier() foundation.NSScriptObjectSpecifier {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("objectSpecifier"))
	return foundation.NSScriptObjectSpecifierFromID(objc.ID(rv))
}
// The name of the document seen by the user in AppleScript.
//
// # Discussion
// 
// This property contains the document name used during scripting. Note that
// this name may be different than the name used in [FileURL].
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/lastComponentOfFileName
func (d NSDocument) LastComponentOfFileName() string {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("lastComponentOfFileName"))
	return foundation.NSStringFromID(rv).String()
}
func (d NSDocument) SetLastComponentOfFileName(value string) {
	objc.Send[struct{}](d.ID, objc.Sel("setLastComponentOfFileName:"), objc.String(value))
}
// See: https://developer.apple.com/documentation/AppKit/NSDocument/observedPresentedItemUbiquityAttributes
func (d NSDocument) ObservedPresentedItemUbiquityAttributes() foundation.INSSet {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("observedPresentedItemUbiquityAttributes"))
	return foundation.NSSetFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AppKit/NSDocument/presentedItemURL
func (d NSDocument) PresentedItemURL() foundation.INSURL {
	rv := objc.Send[objc.ID](d.ID, objc.Sel("presentedItemURL"))
	return foundation.NSURLFromID(objc.ID(rv))
}
// See: https://developer.apple.com/documentation/AppKit/NSDocument/previewRepresentableActivityItems
func (d NSDocument) PreviewRepresentableActivityItems() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](d.ID, objc.Sel("previewRepresentableActivityItems"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
func (d NSDocument) SetPreviewRepresentableActivityItems(value []objectivec.IObject) {
	objc.Send[struct{}](d.ID, objc.Sel("setPreviewRepresentableActivityItems:"), objectivec.IObjectSliceToNSArray(value))
}
// See: https://developer.apple.com/documentation/AppKit/NSDocument/savePanelShowsFileFormatsControl
func (d NSDocument) SavePanelShowsFileFormatsControl() bool {
	rv := objc.Send[bool](d.ID, objc.Sel("savePanelShowsFileFormatsControl"))
	return rv
}

// Returns the types of data the receiver can read natively and any types
// filterable to that native type.
//
// # Return Value
// 
// An array of [NSString] objects representing the readable document types.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/readableTypes
func (_NSDocumentClass NSDocumentClass) ReadableTypes() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSDocumentClass.class), objc.Sel("readableTypes"))
	return objc.ConvertSliceToStrings(rv)
}
// Returns the types of data the receiver can write natively and any types
// filterable to that native type.
//
// # Return Value
// 
// An array of [NSString] objects representing the writable document types.
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/writableTypes
func (_NSDocumentClass NSDocumentClass) WritableTypes() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSDocumentClass.class), objc.Sel("writableTypes"))
	return objc.ConvertSliceToStrings(rv)
}
// A Boolean value that indicates whether the document subclass supports
// autosaving in place.
//
// # Return Value
// 
// [true] if the receiver supports autosaving in place; [false] otherwise.
// 
// # Discussion
// 
// The default implementation of this method returns [false]. You can override
// it and return [true] to declare that your subclass of [NSDocument] can do
// autosaving in place. You should not invoke this method to find out whether
// autosaving in place is actually being done at any particular moment. You
// should instead use the [NSDocument.SaveOperationType] parameter that the
// system passes to your overrides of save and write methods.
// 
// AppKit invokes this method at a variety of times, and not always on the
// main thread. For example,
// [AutosaveWithImplicitCancellabilityCompletionHandler] invokes this method
// as part of determining whether the autosaving will be performed in place
// ([NSDocument.SaveOperationType.autosaveInPlaceOperation]) or in a separate
// autosave directory
// ([NSDocument.SaveOperationType.autosaveElsewhereOperation]). As another
// example, the [CanCloseDocumentWithDelegateShouldCloseSelectorContextInfo]
// method and the [NSDocumentController] machinery for handling unsaved
// changes at app termination time both invoke this method as part of
// determining whether alerts about unsaved changes should be presented to the
// user.
//
// [NSDocument.SaveOperationType.autosaveElsewhereOperation]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/autosaveElsewhereOperation
// [NSDocument.SaveOperationType.autosaveInPlaceOperation]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/autosaveInPlaceOperation
// [NSDocument.SaveOperationType]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/autosavesInPlace
func (_NSDocumentClass NSDocumentClass) AutosavesInPlace() bool {
	rv := objc.Send[bool](objc.ID(_NSDocumentClass.class), objc.Sel("autosavesInPlace"))
	return rv
}
// A Boolean value that indicates whether the document subclass supports
// autosaving of drafts.
//
// # Return Value
// 
// [true] if the receiving subclass of [NSDocument] supports autosaving of
// drafts; otherwise, [false].
// 
// # Discussion
// 
// The system expects that an [NSDocument] subclass that returns [true] from
// this property can properly handle save operations that use the
// [NSDocument.SaveOperationType.autosaveAsOperation] save operation type.
// 
// The default implementation of this property returns [true]. To opt out of
// autosaving in your [NSDocument] subclass, override this property to return
// [false].
// 
// AppKit invokes this property at various times. For example, when calling
// the [UpdateChangeCount] method with [NSDocument.ChangeType.changeDone], but
// without the [NSDocument.ChangeType.changeDiscardable] change type,
// [NSDocument] uses [NSDocument.SaveOperationType.autosaveAsOperation] on the
// next autosave. The operation writes the document’s contents to a new file
// or file package, then changes the document’s current location to point to
// the new file or file package.
// 
// Don’t invoke this property to find out whether autosaving of a draft
// might occur.
//
// [NSDocument.ChangeType.changeDiscardable]: https://developer.apple.com/documentation/AppKit/NSDocument/ChangeType/changeDiscardable
// [NSDocument.ChangeType.changeDone]: https://developer.apple.com/documentation/AppKit/NSDocument/ChangeType/changeDone
// [NSDocument.SaveOperationType.autosaveAsOperation]: https://developer.apple.com/documentation/AppKit/NSDocument/SaveOperationType/autosaveAsOperation
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/autosavesDrafts
func (_NSDocumentClass NSDocumentClass) AutosavesDrafts() bool {
	rv := objc.Send[bool](objc.ID(_NSDocumentClass.class), objc.Sel("autosavesDrafts"))
	return rv
}
// A Boolean value that indicates whether the document subclass supports
// version management.
//
// # Return Value
// 
// [true] if the receiving subclass of [NSDocument] supports Versions;
// otherwise [false].
// 
// # Discussion
// 
// The default implementation of this method returns `[self
// autosavesInPlace]`. You can override it and return [false] to declare that
// [NSDocument] should not preserve old document versions.
// 
// Returning [false] from this method disables version browsing and
// [RevertDocumentToSaved], which rely on version preservation when autosaving
// in place. Returning [true] from this method when [AutosavesInPlace] returns
// [false] will result in undefined behavior.
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/preservesVersions
func (_NSDocumentClass NSDocumentClass) PreservesVersions() bool {
	rv := objc.Send[bool](objc.ID(_NSDocumentClass.class), objc.Sel("preservesVersions"))
	return rv
}
// Returns whether the document object stores its contents in the user’s
// iCloud document storage.
//
// # Discussion
// 
// This method’s default implementation returns [true] if your app has a
// valid iCloud document storage entitlement
// (`com.AppleXCUIElementTypeDeveloperXCUIElementTypeUbiquity()-container-identifiers`
// or
// `com.AppleXCUIElementTypeDeveloperXCUIElementTypeIcloud()-container-identifiers`,
// as described in [Entitlement Key Reference]). When this method returns
// [true], the system adds new menu items and other UI for iCloud documents,
// as appropriate, and allows documents to be saved or moved into the primary
// iCloud container. (The primary iCloud container is the one identified by
// the first container identifier string in the iCloud Containers list in the
// Xcode target editor.)
// 
// To indicate that your [NSDocument] subclass does not use iCloud storage,
// override this method to return [false].
//
// [Entitlement Key Reference]: https://developer.apple.com/library/archive/documentation/Miscellaneous/Reference/EntitlementKeyReference/Chapters/AboutEntitlements.html#//apple_ref/doc/uid/TP40011195
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/usesUbiquitousStorage
func (_NSDocumentClass NSDocumentClass) UsesUbiquitousStorage() bool {
	rv := objc.Send[bool](objc.ID(_NSDocumentClass.class), objc.Sel("usesUbiquitousStorage"))
	return rv
}
// Returns an array of key paths that represent the restorable attributes of
// the document.
//
// # Return Value
// 
// An array of [NSString] objects, each of which contains a key path to one of
// the document’s attributes.
// 
// # Discussion
// 
// You can use this method instead of, or in addition to, the
// [EncodeRestorableStateWithCoder] and [RestoreStateWithCoder] methods to
// save and restore the state of your document. The key paths must refer to
// attributes that are [Key-value coding] and [Key-value observing] compliant.
// 
// When changes are detected, the specified attributes are automatically
// written to disk with the rest of the app’s interface-related state. At
// launch time, the attributes are automatically restored to their previous
// values.
//
// [Key-value coding]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/KeyValueCoding.html#//apple_ref/doc/uid/TP40008195-CH25
// [Key-value observing]: https://developer.apple.com/library/archive/documentation/General/Conceptual/DevPedia-CocoaCore/KVO.html#//apple_ref/doc/uid/TP40008195-CH16
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
//
// See: https://developer.apple.com/documentation/AppKit/NSDocument/restorableStateKeyPaths
func (_NSDocumentClass NSDocumentClass) RestorableStateKeyPaths() []string {
	rv := objc.Send[[]objc.ID](objc.ID(_NSDocumentClass.class), objc.Sel("restorableStateKeyPaths"))
	return objc.ConvertSliceToStrings(rv)
}

			// Protocol methods for NSEditorRegistration
			

			// Protocol methods for NSMenuItemValidation
			

			// Protocol methods for NSUserActivityRestoring
			

			// Protocol methods for NSUserInterfaceValidations
			

// SaveToURLOfTypeForSaveOperation is a synchronous wrapper around [NSDocument.SaveToURLOfTypeForSaveOperationCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NSDocument) SaveToURLOfTypeForSaveOperation(ctx context.Context, url foundation.INSURL, typeName string, saveOperation NSSaveOperationType) error {
	done := make(chan error, 1)
	d.SaveToURLOfTypeForSaveOperationCompletionHandler(url, typeName, saveOperation, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// AutosaveWithImplicitCancellability is a synchronous wrapper around [NSDocument.AutosaveWithImplicitCancellabilityCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NSDocument) AutosaveWithImplicitCancellability(ctx context.Context, autosavingIsImplicitlyCancellable bool) error {
	done := make(chan error, 1)
	d.AutosaveWithImplicitCancellabilityCompletionHandler(autosavingIsImplicitlyCancellable, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// StopBrowsingVersions is a synchronous wrapper around [NSDocument.StopBrowsingVersionsWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NSDocument) StopBrowsingVersions(ctx context.Context) error {
	done := make(chan struct{}, 1)
	d.StopBrowsingVersionsWithCompletionHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RestoreDocumentWindowWithIdentifierState is a synchronous wrapper around [NSDocument.RestoreDocumentWindowWithIdentifierStateCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NSDocument) RestoreDocumentWindowWithIdentifierState(ctx context.Context, identifier NSUserInterfaceItemIdentifier, state foundation.INSCoder) (*NSWindow, error) {
	type result struct {
		val *NSWindow
		err error
	}
	done := make(chan result, 1)
	d.RestoreDocumentWindowWithIdentifierStateCompletionHandler(identifier, state, func(val *NSWindow, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// PerformSynchronousFileAccessUsingBlockSync is a synchronous wrapper around [NSDocument.PerformSynchronousFileAccessUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (d NSDocument) PerformSynchronousFileAccessUsingBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	d.PerformSynchronousFileAccessUsingBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PerformAsynchronousFileAccessUsingBlockSync is a synchronous wrapper around [NSDocument.PerformAsynchronousFileAccessUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (d NSDocument) PerformAsynchronousFileAccessUsingBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	d.PerformAsynchronousFileAccessUsingBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// PerformActivityWithSynchronousWaitingUsingBlockSync is a synchronous wrapper around [NSDocument.PerformActivityWithSynchronousWaitingUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (d NSDocument) PerformActivityWithSynchronousWaitingUsingBlockSync(ctx context.Context, waitSynchronously bool) error {
	done := make(chan struct{}, 1)
	d.PerformActivityWithSynchronousWaitingUsingBlock(waitSynchronously, func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ContinueActivityUsingBlockSync is a synchronous wrapper around [NSDocument.ContinueActivityUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (d NSDocument) ContinueActivityUsingBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	d.ContinueActivityUsingBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ContinueAsynchronousWorkOnMainThreadUsingBlockSync is a synchronous wrapper around [NSDocument.ContinueAsynchronousWorkOnMainThreadUsingBlock].
// It blocks until the completion handler fires or the context is cancelled.
func (d NSDocument) ContinueAsynchronousWorkOnMainThreadUsingBlockSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	d.ContinueAsynchronousWorkOnMainThreadUsingBlock(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// MoveDocumentSync is a synchronous wrapper around [NSDocument.MoveDocumentWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NSDocument) MoveDocumentSync(ctx context.Context) (bool, error) {
	done := make(chan bool, 1)
	d.MoveDocumentWithCompletionHandler(func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// MoveToURL is a synchronous wrapper around [NSDocument.MoveToURLCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NSDocument) MoveToURL(ctx context.Context, url foundation.INSURL) error {
	done := make(chan error, 1)
	d.MoveToURLCompletionHandler(url, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// LockDocumentSync is a synchronous wrapper around [NSDocument.LockDocumentWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NSDocument) LockDocumentSync(ctx context.Context) (bool, error) {
	done := make(chan bool, 1)
	d.LockDocumentWithCompletionHandler(func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// Lock is a synchronous wrapper around [NSDocument.LockWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NSDocument) Lock(ctx context.Context) error {
	done := make(chan error, 1)
	d.LockWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// UnlockDocumentSync is a synchronous wrapper around [NSDocument.UnlockDocumentWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NSDocument) UnlockDocumentSync(ctx context.Context) (bool, error) {
	done := make(chan bool, 1)
	d.UnlockDocumentWithCompletionHandler(func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// Unlock is a synchronous wrapper around [NSDocument.UnlockWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NSDocument) Unlock(ctx context.Context) error {
	done := make(chan error, 1)
	d.UnlockWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ShareDocumentWithSharingService is a synchronous wrapper around [NSDocument.ShareDocumentWithSharingServiceCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NSDocument) ShareDocumentWithSharingService(ctx context.Context, sharingService INSSharingService) (bool, error) {
	done := make(chan bool, 1)
	d.ShareDocumentWithSharingServiceCompletionHandler(sharingService, func(val bool) {
		done <- val
	})
	select {
	case r := <-done:
		return r, nil
	case <-ctx.Done():
		return false, ctx.Err()
	}
}

// AccommodatePresentedItemDeletion is a synchronous wrapper around [NSDocument.AccommodatePresentedItemDeletionWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NSDocument) AccommodatePresentedItemDeletion(ctx context.Context) error {
	done := make(chan error, 1)
	d.AccommodatePresentedItemDeletionWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RelinquishPresentedItemToReaderSync is a synchronous wrapper around [NSDocument.RelinquishPresentedItemToReader].
// It blocks until the completion handler fires or the context is cancelled.
func (d NSDocument) RelinquishPresentedItemToReaderSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	d.RelinquishPresentedItemToReader(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RelinquishPresentedItemToWriterSync is a synchronous wrapper around [NSDocument.RelinquishPresentedItemToWriter].
// It blocks until the completion handler fires or the context is cancelled.
func (d NSDocument) RelinquishPresentedItemToWriterSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	d.RelinquishPresentedItemToWriter(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SavePresentedItemChanges is a synchronous wrapper around [NSDocument.SavePresentedItemChangesWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (d NSDocument) SavePresentedItemChanges(ctx context.Context) error {
	done := make(chan error, 1)
	d.SavePresentedItemChangesWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

