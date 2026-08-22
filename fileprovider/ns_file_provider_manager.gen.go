// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSFileProviderManager] class.
var (
	_NSFileProviderManagerClass     NSFileProviderManagerClass
	_NSFileProviderManagerClassOnce sync.Once
)

func getNSFileProviderManagerClass() NSFileProviderManagerClass {
	_NSFileProviderManagerClassOnce.Do(func() {
		_NSFileProviderManagerClass = NSFileProviderManagerClass{class: objc.GetClass("NSFileProviderManager")}
	})
	return _NSFileProviderManagerClass
}

// GetNSFileProviderManagerClass returns the class object for NSFileProviderManager.
func GetNSFileProviderManagerClass() NSFileProviderManagerClass {
	return getNSFileProviderManagerClass()
}

type NSFileProviderManagerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSFileProviderManagerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSFileProviderManagerClass) Alloc() NSFileProviderManager {
	rv := objc.Send[NSFileProviderManager](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A manager object that you use to communicate with the file provider from
// either your app or your File Provider extension.
//
// # Translating user-visible URLs
//
//   - [NSFileProviderManager.GetUserVisibleURLForItemIdentifierCompletionHandler]: Returns the user-visible URL for an item.
//
// # Working with items
//
//   - [NSFileProviderManager.ReimportItemsBelowItemWithIdentifierCompletionHandler]: Tells the system to reimport the item and its content recursively.
//   - [NSFileProviderManager.EvictItemWithIdentifierCompletionHandler]: Asks the system to remove an item from its cache.
//   - [NSFileProviderManager.RequestModificationOfFieldsForItemWithIdentifierOptionsCompletionHandler]
//   - [NSFileProviderManager.EnumeratorForMaterializedItems]: Returns an enumerator for all the items the system currently stores on disk.
//   - [NSFileProviderManager.EnumeratorForPendingItems]: Returns an enumerator for the set of pending items.
//
// # Performing actions
//
//   - [NSFileProviderManager.RegisterURLSessionTaskForItemWithIdentifierCompletionHandler]: Registers the URL session task responsible for the specified item.
//   - [NSFileProviderManager.SignalEnumeratorForContainerItemIdentifierCompletionHandler]: Alerts the system to changes in the specified folder’s content.
//   - [NSFileProviderManager.WaitForChangesOnItemsBelowItemWithIdentifierCompletionHandler]: Requests a notification after the system completes all the specified changes.
//   - [NSFileProviderManager.GlobalProgressForKind]: Returns a progress object that tracks either the uploading or downloading of items from the File Provider extension’s remote storage.
//
// # Working with domains
//
//   - [NSFileProviderManager.DisconnectWithReasonOptionsCompletionHandler]: Disconnects the domain from the extension.
//   - [NSFileProviderManager.ReconnectWithCompletionHandler]: Reconnects the domain with the extension.
//   - [NSFileProviderManager.WaitForStabilizationWithCompletionHandler]: Requests a notification after the domain stabilizes.
//   - [NSFileProviderManager.TemporaryDirectoryURLWithError]: Returns the URL of a directory that the File Provider extension can use to temporarily store files before passing them to the system.
//
// # Syncing Desktop and Documents folders
//
//   - [NSFileProviderManager.ClaimKnownFoldersLocalizedReasonCompletionHandler]: Asks the domain to sync the specified known folders.
//   - [NSFileProviderManager.ReleaseKnownFoldersLocalizedReasonCompletionHandler]: Asks the system to stop replicating the specified known folders in the domain.
//
// # Working with external volumes
//
//   - [NSFileProviderManager.StateDirectoryURLWithError]: Returns a URL for a directory for storing state information for the domain.
//
// # Using services
//
//   - [NSFileProviderManager.GetServiceWithNameItemIdentifierCompletionHandler]
//
// # Testing
//
//   - [NSFileProviderManager.ListAvailableTestingOperationsWithError]: Lists all the operations that are ready for scheduling.
//   - [NSFileProviderManager.RunTestingOperationsError]: Asks the system to schedule and execute the specified operations.
//
// # Handling errors
//
//   - [NSFileProviderManager.SignalErrorResolvedCompletionHandler]: Indicates a resolved error.
//
// # Collecting diagnostic reports
//
//   - [NSFileProviderManager.RequestDiagnosticCollectionForItemWithIdentifierErrorReasonCompletionHandler]: Requests a diagnostics collection for use when working directly with Apple to improve sync behavior.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager
type NSFileProviderManager struct {
	objectivec.Object
}

// NSFileProviderManagerFromID constructs a [NSFileProviderManager] from an objc.ID.
//
// A manager object that you use to communicate with the file provider from
// either your app or your File Provider extension.
func NSFileProviderManagerFromID(id objc.ID) NSFileProviderManager {
	return NSFileProviderManager{objectivec.Object{ID: id}}
}

// NOTE: NSFileProviderManager adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSFileProviderManager] class.
//
// # Translating user-visible URLs
//
//   - [INSFileProviderManager.GetUserVisibleURLForItemIdentifierCompletionHandler]: Returns the user-visible URL for an item.
//
// # Working with items
//
//   - [INSFileProviderManager.ReimportItemsBelowItemWithIdentifierCompletionHandler]: Tells the system to reimport the item and its content recursively.
//   - [INSFileProviderManager.EvictItemWithIdentifierCompletionHandler]: Asks the system to remove an item from its cache.
//   - [INSFileProviderManager.RequestModificationOfFieldsForItemWithIdentifierOptionsCompletionHandler]
//   - [INSFileProviderManager.EnumeratorForMaterializedItems]: Returns an enumerator for all the items the system currently stores on disk.
//   - [INSFileProviderManager.EnumeratorForPendingItems]: Returns an enumerator for the set of pending items.
//
// # Performing actions
//
//   - [INSFileProviderManager.RegisterURLSessionTaskForItemWithIdentifierCompletionHandler]: Registers the URL session task responsible for the specified item.
//   - [INSFileProviderManager.SignalEnumeratorForContainerItemIdentifierCompletionHandler]: Alerts the system to changes in the specified folder’s content.
//   - [INSFileProviderManager.WaitForChangesOnItemsBelowItemWithIdentifierCompletionHandler]: Requests a notification after the system completes all the specified changes.
//   - [INSFileProviderManager.GlobalProgressForKind]: Returns a progress object that tracks either the uploading or downloading of items from the File Provider extension’s remote storage.
//
// # Working with domains
//
//   - [INSFileProviderManager.DisconnectWithReasonOptionsCompletionHandler]: Disconnects the domain from the extension.
//   - [INSFileProviderManager.ReconnectWithCompletionHandler]: Reconnects the domain with the extension.
//   - [INSFileProviderManager.WaitForStabilizationWithCompletionHandler]: Requests a notification after the domain stabilizes.
//   - [INSFileProviderManager.TemporaryDirectoryURLWithError]: Returns the URL of a directory that the File Provider extension can use to temporarily store files before passing them to the system.
//
// # Syncing Desktop and Documents folders
//
//   - [INSFileProviderManager.ClaimKnownFoldersLocalizedReasonCompletionHandler]: Asks the domain to sync the specified known folders.
//   - [INSFileProviderManager.ReleaseKnownFoldersLocalizedReasonCompletionHandler]: Asks the system to stop replicating the specified known folders in the domain.
//
// # Working with external volumes
//
//   - [INSFileProviderManager.StateDirectoryURLWithError]: Returns a URL for a directory for storing state information for the domain.
//
// # Using services
//
//   - [INSFileProviderManager.GetServiceWithNameItemIdentifierCompletionHandler]
//
// # Testing
//
//   - [INSFileProviderManager.ListAvailableTestingOperationsWithError]: Lists all the operations that are ready for scheduling.
//   - [INSFileProviderManager.RunTestingOperationsError]: Asks the system to schedule and execute the specified operations.
//
// # Handling errors
//
//   - [INSFileProviderManager.SignalErrorResolvedCompletionHandler]: Indicates a resolved error.
//
// # Collecting diagnostic reports
//
//   - [INSFileProviderManager.RequestDiagnosticCollectionForItemWithIdentifierErrorReasonCompletionHandler]: Requests a diagnostics collection for use when working directly with Apple to improve sync behavior.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager
type INSFileProviderManager interface {
	objectivec.IObject

	// Topic: Translating user-visible URLs

	// Returns the user-visible URL for an item.
	GetUserVisibleURLForItemIdentifierCompletionHandler(itemIdentifier NSFileProviderItemIdentifier, completionHandler URLErrorHandler)

	// Topic: Working with items

	// Tells the system to reimport the item and its content recursively.
	ReimportItemsBelowItemWithIdentifierCompletionHandler(itemIdentifier NSFileProviderItemIdentifier, completionHandler ErrorHandler)
	// Asks the system to remove an item from its cache.
	EvictItemWithIdentifierCompletionHandler(itemIdentifier NSFileProviderItemIdentifier, completionHandler ErrorHandler)
	RequestModificationOfFieldsForItemWithIdentifierOptionsCompletionHandler(fields NSFileProviderItemFields, itemIdentifier NSFileProviderItemIdentifier, options NSFileProviderModifyItemOptions, completionHandler ErrorHandler)
	// Returns an enumerator for all the items the system currently stores on disk.
	EnumeratorForMaterializedItems() NSFileProviderEnumerator
	// Returns an enumerator for the set of pending items.
	EnumeratorForPendingItems() NSFileProviderPendingSetEnumerator

	// Topic: Performing actions

	// Registers the URL session task responsible for the specified item.
	RegisterURLSessionTaskForItemWithIdentifierCompletionHandler(task foundation.NSURLSessionTask, identifier NSFileProviderItemIdentifier, completion ErrorHandler)
	// Alerts the system to changes in the specified folder’s content.
	SignalEnumeratorForContainerItemIdentifierCompletionHandler(containerItemIdentifier NSFileProviderItemIdentifier, completion ErrorHandler)
	// Requests a notification after the system completes all the specified changes.
	WaitForChangesOnItemsBelowItemWithIdentifierCompletionHandler(itemIdentifier NSFileProviderItemIdentifier, completionHandler ErrorHandler)
	// Returns a progress object that tracks either the uploading or downloading of items from the File Provider extension’s remote storage.
	GlobalProgressForKind(kind foundation.NSProgressFileOperationKind) foundation.Progress

	// Topic: Working with domains

	// Disconnects the domain from the extension.
	DisconnectWithReasonOptionsCompletionHandler(localizedReason string, options NSFileProviderManagerDisconnectionOptions, completionHandler ErrorHandler)
	// Reconnects the domain with the extension.
	ReconnectWithCompletionHandler(completionHandler ErrorHandler)
	// Requests a notification after the domain stabilizes.
	WaitForStabilizationWithCompletionHandler(completionHandler ErrorHandler)
	// Returns the URL of a directory that the File Provider extension can use to temporarily store files before passing them to the system.
	TemporaryDirectoryURLWithError() (foundation.NSURL, error)

	// Topic: Syncing Desktop and Documents folders

	// Asks the domain to sync the specified known folders.
	ClaimKnownFoldersLocalizedReasonCompletionHandler(knownFolders INSFileProviderKnownFolderLocations, localizedReason string, completionHandler ErrorHandler)
	// Asks the system to stop replicating the specified known folders in the domain.
	ReleaseKnownFoldersLocalizedReasonCompletionHandler(knownFolders NSFileProviderKnownFolders, localizedReason string, completionHandler ErrorHandler)

	// Topic: Working with external volumes

	// Returns a URL for a directory for storing state information for the domain.
	StateDirectoryURLWithError() (foundation.NSURL, error)

	// Topic: Using services

	GetServiceWithNameItemIdentifierCompletionHandler(serviceName foundation.NSFileProviderServiceName, itemIdentifier NSFileProviderItemIdentifier, completionHandler FileProviderServiceErrorHandler)

	// Topic: Testing

	// Lists all the operations that are ready for scheduling.
	ListAvailableTestingOperationsWithError() ([]objectivec.IObject, error)
	// Asks the system to schedule and execute the specified operations.
	RunTestingOperationsError(operations []objectivec.IObject) (foundation.INSDictionary, error)

	// Topic: Handling errors

	// Indicates a resolved error.
	SignalErrorResolvedCompletionHandler(error_ foundation.NSError, completionHandler ErrorHandler)

	// Topic: Collecting diagnostic reports

	// Requests a diagnostics collection for use when working directly with Apple to improve sync behavior.
	RequestDiagnosticCollectionForItemWithIdentifierErrorReasonCompletionHandler(itemIdentifier NSFileProviderItemIdentifier, errorReason foundation.NSError, completionHandler ErrorHandler)
}

// Init initializes the instance.
func (f NSFileProviderManager) Init() NSFileProviderManager {
	rv := objc.Send[NSFileProviderManager](f.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f NSFileProviderManager) Autorelease() NSFileProviderManager {
	rv := objc.Send[NSFileProviderManager](f.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSFileProviderManager creates a new NSFileProviderManager instance.
func NewNSFileProviderManager() NSFileProviderManager {
	class := getNSFileProviderManagerClass()
	rv := objc.Send[NSFileProviderManager](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a newly created file provider manager for the specified domain.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/init(for:)
func NewFileProviderManagerForDomain(domain INSFileProviderDomain) NSFileProviderManager {
	rv := objc.Send[objc.ID](objc.ID(getNSFileProviderManagerClass().class), objc.Sel("managerForDomain:"), domain)
	return NSFileProviderManagerFromID(rv)
}

// Returns the user-visible URL for an item.
//
// itemIdentifier: The item’s identifier.
//
// completionHandler: A block that the system calls after determining the item’s URL. The
// system passes the following parameters:
//
// `userVisibleFile`: The URL of the user visible file, or `nil` if an error
// occurs. `error`: If an error occurs, this object contains information about
// the error; otherwise, it’s `nil`.
//
// # Discussion
//
// Calling this method marks the process so that accessing the URL won’t
// materialize the item. Instead, any attempt to read or write to an
// unmaterialized item fails with a [EDEADLK] POSIX error.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/getUserVisibleURL(for:completionHandler:)
//
// [EDEADLK]: https://developer.apple.com/documentation/Foundation/POSIXError/EDEADLK
func (f NSFileProviderManager) GetUserVisibleURLForItemIdentifierCompletionHandler(itemIdentifier NSFileProviderItemIdentifier, completionHandler URLErrorHandler) {
	_block1, _ := NewURLErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("getUserVisibleURLForItemIdentifier:completionHandler:"), itemIdentifier, _block1)
}

// Tells the system to reimport the item and its content recursively.
//
// itemIdentifier: The identifier of the item to reimport. The system reimports the item and
// all of its children.
//
// completionHandler: A block called by the system immediately after receiving the request. The
// completion handler takes the following parameters:
//
// error: If an error occurs, this object contains information about the
// error; otherwise, it’s `nil`.
//
// # Discussion
//
// This method tells the system that the specified item identifiers are no
// longer valid. Your file provider extension should call this method if it
// has lost track of its synchronization state and can’t guarantee the
// stability of the item identifiers anymore.
//
// The system calls
// [CreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler]
// and passes the [NSFileProviderCreateItemMayAlreadyExist] option for each
// affected identifier in its working set. Your File Provider extension can
// then specify a new identifier for each item.
//
// The system then calls [ImportDidFinishWithCompletionHandler] when the
// import is complete. If successful, the system always reimports the
// specified subtree, but it may reimport other items as well. If the item
// specified by the `itemIdentifier` parameter has no on-disk representation,
// the method fails with an [NSFileProviderError.Code.noSuchItem] error.
//
// If your file provider loses synchronization but is still able to guarantee
// the stability of the identifiers, you don’t need to reimport the items.
// Instead, if the system queries the working set with an anchor that predates
// the synchronization loss, your File Provider extension can fail with an
// [NSFileProviderError.Code.syncAnchorExpired] error.
//
// If your file provider loses synchronization, but you aren’t interested in
// preserving the local data, you can resolve the issue by removing and then
// adding the domain back.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/reimportItems(below:completionHandler:)
//
// [NSFileProviderError.Code.noSuchItem]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/noSuchItem
// [NSFileProviderError.Code.syncAnchorExpired]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/syncAnchorExpired
func (f NSFileProviderManager) ReimportItemsBelowItemWithIdentifierCompletionHandler(itemIdentifier NSFileProviderItemIdentifier, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("reimportItemsBelowItemWithIdentifier:completionHandler:"), itemIdentifier, _block1)
}

// Asks the system to remove an item from its cache.
//
// itemIdentifier: The item’s identifier.
//
// completionHandler: A block that the system calls after removing the item from disk. The system
// passes the following parameter:
//
// `error`: If an error occurs, this object contains information about the
// error; otherwise, it’s `nil`.
//
// # Discussion
//
// Calling this method turns a materialized item into a dataless item to free
// up disk space. For more information on materialized and dataless items, see
// [Synchronizing the File Provider Extension].
//
// If the item is a document without local changes, this method deletes the
// local copy of the item’s content. If the item has local changes, it fails
// with an [NSFileWriteNoPermissionError] error.
//
// When called on a directory, the system recursively evicts the directory’s
// content. It deletes the content of any materialized files, and recursively
// evicts any subdirectories. After it has successfully evicted all the
// content, it deletes its list of the directory’s content, making the
// directory dataless. The next time the system accesses the directory, it
// requests a list of the contents using the [NSFileProviderEnumerating]
// protocol.
//
// If the system encounters a nonevictable child, eviction stops immediately,
// and the system calls the completion handler with a
// [NSFileProviderError.Code.nonEvictableChildren] error. The error includes
// information about the nonevictable child in its [underlyingErrors]
// property. The system may have evicted other materialized items, based on
// the traversal order.
//
// The system calls the completion handler after it successfully evicts all
// items, or immediately when an error occurs. Eviction might fail with the
// following errors:
//
// - [NSFileProviderError.Code.unsyncedEdits] if the item had nonuploaded
// changes. - [NSFileProviderError.Code.nonEvictable] if the user has marked
// the item as nonevictable. - [EBUSY] if the item has open file descriptors
// on it. - [EMLINK] if the item has too many hardlinks. - Other
// [NSPOSIXErrorDomain] error codes if the system can’t access or manipulate
// the corresponding file.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/evictItem(identifier:completionHandler:)
//
// [NSFileProviderError.Code.nonEvictableChildren]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/nonEvictableChildren
// [NSFileProviderError.Code.nonEvictable]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/nonEvictable
// [NSFileProviderError.Code.unsyncedEdits]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/unsyncedEdits
// [NSFileWriteNoPermissionError]: https://developer.apple.com/documentation/Foundation/NSFileWriteNoPermissionError-swift.var
// [NSPOSIXErrorDomain]: https://developer.apple.com/documentation/Foundation/NSPOSIXErrorDomain
// [Synchronizing the File Provider Extension]: https://developer.apple.com/documentation/FileProvider/synchronizing-the-file-provider-extension
// [underlyingErrors]: https://developer.apple.com/documentation/Foundation/NSError/underlyingErrors
func (f NSFileProviderManager) EvictItemWithIdentifierCompletionHandler(itemIdentifier NSFileProviderItemIdentifier, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("evictItemWithIdentifier:completionHandler:"), itemIdentifier, _block1)
}

// # Discussion
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/requestModification(of:forItemWithIdentifier:options:completionHandler:)
func (f NSFileProviderManager) RequestModificationOfFieldsForItemWithIdentifierOptionsCompletionHandler(fields NSFileProviderItemFields, itemIdentifier NSFileProviderItemIdentifier, options NSFileProviderModifyItemOptions, completionHandler ErrorHandler) {
	_block3, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("requestModificationOfFields:forItemWithIdentifier:options:completionHandler:"), fields, itemIdentifier, options, _block3)
}

// Returns an enumerator for all the items the system currently stores on
// disk.
//
// # Discussion
//
// In most cases, the system requests an enumerator from the File Provider.
// The file provider creates the enumerator, and uses it to pass information
// back to the system. In this case, however, the roles are reversed. The File
// Provider extension calls this method to request an enumerator from the
// system. The system then creates the enumerator and uses it to pass the list
// of items currently stored on disk back to the File Provider extension.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/enumeratorForMaterializedItems()
func (f NSFileProviderManager) EnumeratorForMaterializedItems() NSFileProviderEnumerator {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("enumeratorForMaterializedItems"))
	return NSFileProviderEnumeratorObjectFromID(rv)
}

// Returns an enumerator for the set of pending items.
//
// # Discussion
//
// When the set of pending items changes, the system calls
// [PendingItemsDidChangeWithCompletionHandler].
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/enumeratorForPendingItems()
func (f NSFileProviderManager) EnumeratorForPendingItems() NSFileProviderPendingSetEnumerator {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("enumeratorForPendingItems"))
	return NSFileProviderPendingSetEnumeratorObjectFromID(rv)
}

// Registers the URL session task responsible for the specified item.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/register(_:forItemWithIdentifier:completionHandler:)
func (f NSFileProviderManager) RegisterURLSessionTaskForItemWithIdentifierCompletionHandler(task foundation.NSURLSessionTask, identifier NSFileProviderItemIdentifier, completion ErrorHandler) {
	_block2, _ := NewErrorBlock(completion)
	objc.Send[objc.ID](f.ID, objc.Sel("registerURLSessionTask:forItemWithIdentifier:completionHandler:"), task, identifier, _block2)
}

// Alerts the system to changes in the specified folder’s content.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/signalEnumerator(for:completionHandler:)
func (f NSFileProviderManager) SignalEnumeratorForContainerItemIdentifierCompletionHandler(containerItemIdentifier NSFileProviderItemIdentifier, completion ErrorHandler) {
	_block1, _ := NewErrorBlock(completion)
	objc.Send[objc.ID](f.ID, objc.Sel("signalEnumeratorForContainerItemIdentifier:completionHandler:"), containerItemIdentifier, _block1)
}

// Requests a notification after the system completes all the specified
// changes.
//
// itemIdentifier: The item’s identifier.
//
// completionHandler: A block that the system calls after all the changes are complete. The block
// takes the following parameters:
//
// `error`: If an error occurs, this object contains information about the
// error; otherwise, it’s `nil`.
//
// # Discussion
//
// This method waits for all the changes to the item’s descendants to
// complete before calling the completion handler. If an error occurs during
// this process, the system immediately passes the error to the completion
// handler, and you can’t assume all the changes have completed.
//
// If the `itemIdentifier` property doesn’t refer to a directory, this
// method immediately calls the completion handler.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/waitForChanges(below:completionHandler:)
func (f NSFileProviderManager) WaitForChangesOnItemsBelowItemWithIdentifierCompletionHandler(itemIdentifier NSFileProviderItemIdentifier, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("waitForChangesOnItemsBelowItemWithIdentifier:completionHandler:"), itemIdentifier, _block1)
}

// Returns a progress object that tracks either the uploading or downloading
// of items from the File Provider extension’s remote storage.
//
// kind: The kind of operation. This method only accepts two values: [uploading] and
// [downloading].
//
// # Discussion
//
// The returned progress instance tracks ongoing operations. This method
// supports two kinds of operations:
//
// [uploading]: Uploading items from the local storage to the remote storage.
// [downloading]: Downloading items from the remote storage to the local
// storage.
//
// The progress instance has its [fileOperationKind] property set. It also
// provides the number of items to upload or download, the number of bytes
// already transferred, and the total number of bytes to transfer. The grand
// total is reset to `0` when there are no operations left.
//
// If new matching operations begin while the progress instance is running, it
// adds the new operations to the existing data. By default, when there are no
// matching operations, the progress has its values set to `1` and its state
// set to finished.
//
// The system updates the progress instance on the main queue. You must retain
// the progress item, and observe its changes through key-value observing. For
// more information, see `Using Key-Value Observing in Swift`.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/globalProgress(for:)
//
// [downloading]: https://developer.apple.com/documentation/Foundation/Progress/FileOperationKind-swift.struct/downloading
// [uploading]: https://developer.apple.com/documentation/Foundation/Progress/FileOperationKind-swift.struct/uploading
// [fileOperationKind]: https://developer.apple.com/documentation/Foundation/Progress/fileOperationKind-swift.property
//
// [downloading]: https://developer.apple.com/documentation/Foundation/Progress/FileOperationKind-swift.struct/downloading
// [uploading]: https://developer.apple.com/documentation/Foundation/Progress/FileOperationKind-swift.struct/uploading
func (f NSFileProviderManager) GlobalProgressForKind(kind foundation.NSProgressFileOperationKind) foundation.Progress {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("globalProgressForKind:"), objc.String(string(kind)))
	return foundation.NSProgressFromID(rv)
}

// Disconnects the domain from the extension.
//
// localizedReason: A localized string that describes the reason for disconnecting the domain.
//
// options: Options for the disconnection. For a complete list of valid options, see
// [NSFileProviderManager.DisconnectionOptions].
//
// completionHandler: A block that the system calls after disconnecting the domain. The block
// takes the following parameter:
//
// `error`: If an error occurs, this object contains information about the
// error; otherwise, it’s `nil`.
//
// # Discussion
//
// Call this method to disconnect the domain from the extension. While the
// domain is disconnected, the user can continue to browse its content, but
// the extension no longer receives updates about changes.
//
// Call the [NSFileProviderManager.ReconnectWithCompletionHandler] method to
// reconnect the domain.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/disconnect(reason:options:completionHandler:)
//
// [NSFileProviderManager.DisconnectionOptions]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/DisconnectionOptions
func (f NSFileProviderManager) DisconnectWithReasonOptionsCompletionHandler(localizedReason string, options NSFileProviderManagerDisconnectionOptions, completionHandler ErrorHandler) {
	_block2, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("disconnectWithReason:options:completionHandler:"), objc.String(localizedReason), options, _block2)
}

// Reconnects the domain with the extension.
//
// completionHandler: A block that the system calls after reconnecting the domain. The block
// takes the following parameter:
//
// `error`: If an error occurs, this object contains information about the
// error; otherwise, it’s `nil`.
//
// # Discussion
//
// Call this method to reconnect the domain after a call to the
// [NSFileProviderManager.DisconnectWithReasonOptionsCompletionHandler]
// method.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/reconnect(completionHandler:)
func (f NSFileProviderManager) ReconnectWithCompletionHandler(completionHandler ErrorHandler) {
	_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("reconnectWithCompletionHandler:"), _block0)
}

// Requests a notification after the domain stabilizes.
//
// completionHandler: A block that the system calls after pending changes to both the file system
// and the provider have completed. The system passes the following
// parameters:
//
// `error`: If an error occurs, this object contains information about the
// error; otherwise, it’s `nil`.
//
// # Discussion
//
// Use this method to enforce a consistent state for testing. The system calls
// the completion handler after all the pending changes to the local cache and
// the remote storage have completed. The system waits on any changes that
// requested before the call to
// [NSFileProviderManager.WaitForStabilizationWithCompletionHandler], but
// which haven’t completed yet.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/waitForStabilization(completionHandler:)
func (f NSFileProviderManager) WaitForStabilizationWithCompletionHandler(completionHandler ErrorHandler) {
	_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("waitForStabilizationWithCompletionHandler:"), _block0)
}

// Returns the URL of a directory that the File Provider extension can use to
// temporarily store files before passing them to the system.
//
// # Discussion
//
// The system guarantees that the temporary URL refers to a directory on the
// same volume as the user-visible URL so that the system can automatically
// clone or move files between the temporary URL and the user-visible URL. For
// example, the File Provider extension can use the temporary directory to
// store content passed to the
// [CreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler] or
// [ModifyItemBaseVersionChangedFieldsContentsOptionsRequestCompletionHandler]
// methods.
//
// When you implement your File Provider extension’s
// [FetchContentsForItemWithIdentifierVersionRequestCompletionHandler] method,
// the URL you pass to the completion handler must be on the same volume as
// the temporary directory, so the system can clone it to provide the content
// for the dataless item.
//
// This method fails if the system can’t find a suitable directory, for
// example, if the domain doesn’t exist. However, it can’t fail if the
// file provider has an active instance for the specified domain.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/temporaryDirectoryURL()
func (f NSFileProviderManager) TemporaryDirectoryURLWithError() (foundation.NSURL, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("temporaryDirectoryURLWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSURL{}, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSURLFromID(rv), nil

}

// Asks the domain to sync the specified known folders.
//
// # Discussion
//
// Use this method to claim a set of known folders according to the
// information in the `knownFolders` parameter. The system only enables sync
// for these folders in the domain if the set of locations is valid and if the
// user agrees.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/claimKnownFolders(_:localizedReason:completionHandler:)
func (f NSFileProviderManager) ClaimKnownFoldersLocalizedReasonCompletionHandler(knownFolders INSFileProviderKnownFolderLocations, localizedReason string, completionHandler ErrorHandler) {
	_block2, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("claimKnownFolders:localizedReason:completionHandler:"), knownFolders, objc.String(localizedReason), _block2)
}

// Asks the system to stop replicating the specified known folders in the
// domain.
//
// # Discussion
//
// Use this method to immediately disable replication of the specified known
// folders.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/releaseKnownFolders(_:localizedReason:completionHandler:)
func (f NSFileProviderManager) ReleaseKnownFoldersLocalizedReasonCompletionHandler(knownFolders NSFileProviderKnownFolders, localizedReason string, completionHandler ErrorHandler) {
	_block2, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("releaseKnownFolders:localizedReason:completionHandler:"), knownFolders, objc.String(localizedReason), _block2)
}

// Returns a URL for a directory for storing state information for the domain.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/stateDirectoryURL()
func (f NSFileProviderManager) StateDirectoryURLWithError() (foundation.NSURL, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("stateDirectoryURLWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return foundation.NSURL{}, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSURLFromID(rv), nil

}

// # Discussion
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/getService(named:for:completionHandler:)
func (f NSFileProviderManager) GetServiceWithNameItemIdentifierCompletionHandler(serviceName foundation.NSFileProviderServiceName, itemIdentifier NSFileProviderItemIdentifier, completionHandler FileProviderServiceErrorHandler) {
	_block2, _ := NewFileProviderServiceErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("getServiceWithName:itemIdentifier:completionHandler:"), serviceName, itemIdentifier, _block2)
}

// Lists all the operations that are ready for scheduling.
//
// # Discussion
//
// The system waits for all the pending disk and working set updates before
// returning the list of available operations. The operations that it returns
// may become invalid if the system receives new events, or when you schedule
// and execute operations using the
// [NSFileProviderManager.RunTestingOperationsError] method.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/listAvailableTestingOperations()
func (f NSFileProviderManager) ListAvailableTestingOperationsWithError() ([]objectivec.IObject, error) {
	var errorPtr objc.ID
	rv := objc.Send[[]objc.ID](f.ID, objc.Sel("listAvailableTestingOperationsWithError:"), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	}), nil

}

// Asks the system to schedule and execute the specified operations.
//
// operations: An array of operations. Populate this array with one or more operations
// returned by the
// [NSFileProviderManager.ListAvailableTestingOperationsWithError] method.
//
// # Discussion
//
// The system waits until all of the specified operations complete and reports
// an error for any operations that fail.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/run(_:)
func (f NSFileProviderManager) RunTestingOperationsError(operations []objectivec.IObject) (foundation.INSDictionary, error) {
	var errorPtr objc.ID
	rv := objc.Send[objc.ID](f.ID, objc.Sel("runTestingOperations:error:"), objectivec.IObjectSliceToNSArray(operations), unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return nil, foundation.NSErrorFrom(errorPtr)
	}
	return foundation.NSDictionaryFromID(rv), nil

}

// Indicates a resolved error.
//
// error: The original error.
//
// completionHandler: A block that the system calls after resuming the action that triggered the
// original error. The block takes the following parameters:
//
// error: If an error occurs, this object contains information about the
// error; otherwise, it’s `nil`.
//
// # Discussion
//
// Use this method if any of your extension’s actions fail because of an
// [NSFileProviderError.Code.notAuthenticated],
// [NSFileProviderError.Code.insufficientQuota], or
// [NSFileProviderError.Code.serverUnreachable] error. As soon as you resolve
// the underlying error, call this method to tell the system to retry the
// original action.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/signalErrorResolved(_:completionHandler:)
//
// [NSFileProviderError.Code.insufficientQuota]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/insufficientQuota
// [NSFileProviderError.Code.notAuthenticated]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/notAuthenticated
// [NSFileProviderError.Code.serverUnreachable]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/serverUnreachable
func (f NSFileProviderManager) SignalErrorResolvedCompletionHandler(error_ foundation.NSError, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("signalErrorResolved:completionHandler:"), error_, _block1)
}

// Requests a diagnostics collection for use when working directly with Apple
// to improve sync behavior.
//
// itemIdentifier: The item that failed to sync.
//
// errorReason: An error instance that indicates why the sync failed. If you use a Swift
// [Error], the system coerces it to an [NSError] and uses only the domain and
// code.
//
// completionHandler: A block or closure that executes after the system processes the diagnostic
// collection request. The completion handler receives an error parameter,
// which is non-`nil` if an error prevented collection of diagnostics. In
// Swift, you can omit the completion handler and instead use a `do`-`catch`
// block to handle the thrown error.
//
// # Discussion
//
// Calling this method requests that the system prompt the person using the
// app, who may be experiencing an issue with sync in your provider. It asks
// their permission to collect and send diagnostic information to Apple for
// further analysis. Even if the person gives their approval, however, the
// prompt or the diagnostic collection might not occur, depending on system
// state and other throttling parameters.
//
// When calling this method, use the `errorReason` parameter to describe the
// error that prompted the request for diagnostics. The system doesn’t show
// this error to the person using the app. Instead, the error appears in any
// generated reports created from the diagnostics.
//
// The system may or may not allow the request, and the method call returns
// normally in either case. It produces an error — thrown in Swift and sent
// to the completion handler in Objective-C — if any of the following occur:
//
// - The app isn’t running on a pre-release build. - The app is running on a
// pre-release build, but the system can’t find the item indicated by the
// `itemIdentifier` parameter.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/requestDiagnosticCollection(for:errorReason:completionHandler:)
//
// [Error]: https://developer.apple.com/documentation/Swift/Error
// [NSError]: https://developer.apple.com/documentation/Foundation/NSError
func (f NSFileProviderManager) RequestDiagnosticCollectionForItemWithIdentifierErrorReasonCompletionHandler(itemIdentifier NSFileProviderItemIdentifier, errorReason foundation.NSError, completionHandler ErrorHandler) {
	_block2, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](f.ID, objc.Sel("requestDiagnosticCollectionForItemWithIdentifier:errorReason:completionHandler:"), itemIdentifier, errorReason, _block2)
}

// Returns the identifier and domain for a user-visible URL.
//
// url: The URL of the item.
//
// completionHandler: A block that the system calls after it gets the items identifier. It has
// the following parameters:
//
// `itemIdentifier`: The item’s identifier. `domainIdentifier`: The
// identifier for the item’s domain. `error`: If an error occurs, this
// object contains information about the error; otherwise, it’s `nil`.
//
// # Discussion
//
// If the URL doesn’t refer to an item managed by your File Provider
// extension, the system returns a [NSFileNoSuchFileError] error.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/getIdentifierForUserVisibleFile(at:completionHandler:)
//
// [NSFileNoSuchFileError]: https://developer.apple.com/documentation/Foundation/NSFileNoSuchFileError-swift.var
func (_NSFileProviderManagerClass NSFileProviderManagerClass) GetIdentifierForUserVisibleFileAtURLCompletionHandler(url foundation.NSURL, completionHandler StringStringErrorHandler) {
	_block1, _ := NewStringStringErrorBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_NSFileProviderManagerClass.class), objc.Sel("getIdentifierForUserVisibleFileAtURL:completionHandler:"), url, _block1)
}

// Creates a new domain that takes ownership of on-disk data that your app
// previously managed without a file provider.
//
// domain: The domain to import.
//
// url: A URL that points to the directory to import.
//
// completionHandler: A block that the system calls as soon as it creates the new domain. It
// takes the following parameters:
//
// `error`: If an error occurs, this object contains information about the
// error; otherwise, it’s `nil`.
//
// # Discussion
//
// Use this method to migrate an existing file hierarchy on disk to a
// [NSFileProviderExtension] without redownloading the data. After you call
// the method, the provided URL is no longer valid. The system has moved and
// now manages all of its contents.
//
// If a domain with the same name already exists, the method fails with an
// [NSFileWriteFileExistsError] error. The URL remains untouched. If the
// system doesn’t allow the extension to request a migration, the method
// fails with an [NSFeatureUnsupportedError] error.
//
// The system starts by moving the provided directory into its local cache,
// and then calls the completion handler. Then, for each item in the
// directory, it calls your extension’s
// [CreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler]
// with the [NSFileProviderCreateItemMayAlreadyExist] option.
//
// When the import finishes, the system calls your extension’s
// [ImportDidFinishWithCompletionHandler] method. If you call
// [NSFileProviderManager.ReimportItemsBelowItemWithIdentifierCompletionHandler]
// before the import finishes, the system makes a single call to
// [ImportDidFinishWithCompletionHandler] for both imports.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/import(_:fromDirectoryAt:completionHandler:)
//
// [NSFeatureUnsupportedError]: https://developer.apple.com/documentation/Foundation/NSFeatureUnsupportedError-swift.var
// [NSFileProviderExtension]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension
// [NSFileWriteFileExistsError]: https://developer.apple.com/documentation/Foundation/NSFileWriteFileExistsError-swift.var
func (_NSFileProviderManagerClass NSFileProviderManagerClass) ImportDomainFromDirectoryAtURLCompletionHandler(domain INSFileProviderDomain, url foundation.NSURL, completionHandler ErrorHandler) {
	_block2, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_NSFileProviderManagerClass.class), objc.Sel("importDomain:fromDirectoryAtURL:completionHandler:"), domain, url, _block2)
}

// Adds a domain to the File Provider extension.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/add(_:completionHandler:)
func (_NSFileProviderManagerClass NSFileProviderManagerClass) AddDomainCompletionHandler(domain INSFileProviderDomain, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_NSFileProviderManagerClass.class), objc.Sel("addDomain:completionHandler:"), domain, _block1)
}

// Returns all of the File Provider extension’s domains.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/getDomainsWithCompletionHandler(_:)
func (_NSFileProviderManagerClass NSFileProviderManagerClass) GetDomainsWithCompletionHandler(completionHandler NSFileProviderDomainArrayErrorHandler) {
	_block0, _ := NewNSFileProviderDomainArrayErrorBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_NSFileProviderManagerClass.class), objc.Sel("getDomainsWithCompletionHandler:"), _block0)
}

// Removes a domain from the File Provider extension.
//
// domain: The domain to remove.
//
// completionHandler: A block that the system calls after removing the domain. It takes the
// following parameters:
//
// `error`: If an error occurs, this object contains information about the
// error; otherwise, it’s `nil`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/remove(_:completionHandler:)
func (_NSFileProviderManagerClass NSFileProviderManagerClass) RemoveDomainCompletionHandler(domain INSFileProviderDomain, completionHandler ErrorHandler) {
	_block1, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_NSFileProviderManagerClass.class), objc.Sel("removeDomain:completionHandler:"), domain, _block1)
}

// Removes a domain from the File Provider extension using the specified
// options.
//
// domain: The domain to remove.
//
// mode: An option that determines how the system handles user data. For a complete
// list of options, see [NSFileProviderManager.DomainRemovalMode].
//
// completionHandler: A block that the system calls after removing the domain. It takes the
// following parameter:
//
// `error`: If an error occurs, this object contains information about the
// error; otherwise, it’s `nil`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/remove(_:mode:completionHandler:)
//
// [NSFileProviderManager.DomainRemovalMode]: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/DomainRemovalMode
func (_NSFileProviderManagerClass NSFileProviderManagerClass) RemoveDomainModeCompletionHandler(domain INSFileProviderDomain, mode NSFileProviderDomainRemovalMode, completionHandler URLErrorHandler) {
	_block2, _ := NewURLErrorBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_NSFileProviderManagerClass.class), objc.Sel("removeDomain:mode:completionHandler:"), domain, mode, _block2)
}

// Removes all domains from the File Provider extension.
//
// completionHandler: A block that the system calls after removing the domains. It takes the
// following parameters:
//
// `error`: If an error occurs, this object contains information about the
// error; otherwise, it’s `nil`.
//
// # Discussion
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/removeAllDomains(completionHandler:)
func (_NSFileProviderManagerClass NSFileProviderManagerClass) RemoveAllDomainsWithCompletionHandler(completionHandler ErrorHandler) {
	_block0, _ := NewErrorBlock(completionHandler)
	objc.Send[objc.ID](objc.ID(_NSFileProviderManagerClass.class), objc.Sel("removeAllDomainsWithCompletionHandler:"), _block0)
}

// GetUserVisibleURLForItemIdentifier is a synchronous wrapper around [NSFileProviderManager.GetUserVisibleURLForItemIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NSFileProviderManager) GetUserVisibleURLForItemIdentifier(ctx context.Context, itemIdentifier NSFileProviderItemIdentifier) (*foundation.NSURL, error) {
	type result struct {
		val *foundation.NSURL
		err error
	}
	done := make(chan result, 1)
	f.GetUserVisibleURLForItemIdentifierCompletionHandler(itemIdentifier, func(val *foundation.NSURL, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// ReimportItemsBelowItemWithIdentifier is a synchronous wrapper around [NSFileProviderManager.ReimportItemsBelowItemWithIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NSFileProviderManager) ReimportItemsBelowItemWithIdentifier(ctx context.Context, itemIdentifier NSFileProviderItemIdentifier) error {
	done := make(chan error, 1)
	f.ReimportItemsBelowItemWithIdentifierCompletionHandler(itemIdentifier, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// EvictItemWithIdentifier is a synchronous wrapper around [NSFileProviderManager.EvictItemWithIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NSFileProviderManager) EvictItemWithIdentifier(ctx context.Context, itemIdentifier NSFileProviderItemIdentifier) error {
	done := make(chan error, 1)
	f.EvictItemWithIdentifierCompletionHandler(itemIdentifier, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RequestModificationOfFieldsForItemWithIdentifierOptions is a synchronous wrapper around [NSFileProviderManager.RequestModificationOfFieldsForItemWithIdentifierOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NSFileProviderManager) RequestModificationOfFieldsForItemWithIdentifierOptions(ctx context.Context, fields NSFileProviderItemFields, itemIdentifier NSFileProviderItemIdentifier, options NSFileProviderModifyItemOptions) error {
	done := make(chan error, 1)
	f.RequestModificationOfFieldsForItemWithIdentifierOptionsCompletionHandler(fields, itemIdentifier, options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RegisterURLSessionTaskForItemWithIdentifier is a synchronous wrapper around [NSFileProviderManager.RegisterURLSessionTaskForItemWithIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NSFileProviderManager) RegisterURLSessionTaskForItemWithIdentifier(ctx context.Context, task foundation.NSURLSessionTask, identifier NSFileProviderItemIdentifier) error {
	done := make(chan error, 1)
	f.RegisterURLSessionTaskForItemWithIdentifierCompletionHandler(task, identifier, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// SignalEnumeratorForContainerItemIdentifier is a synchronous wrapper around [NSFileProviderManager.SignalEnumeratorForContainerItemIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NSFileProviderManager) SignalEnumeratorForContainerItemIdentifier(ctx context.Context, containerItemIdentifier NSFileProviderItemIdentifier) error {
	done := make(chan error, 1)
	f.SignalEnumeratorForContainerItemIdentifierCompletionHandler(containerItemIdentifier, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// WaitForChangesOnItemsBelowItemWithIdentifier is a synchronous wrapper around [NSFileProviderManager.WaitForChangesOnItemsBelowItemWithIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NSFileProviderManager) WaitForChangesOnItemsBelowItemWithIdentifier(ctx context.Context, itemIdentifier NSFileProviderItemIdentifier) error {
	done := make(chan error, 1)
	f.WaitForChangesOnItemsBelowItemWithIdentifierCompletionHandler(itemIdentifier, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ImportDomainFromDirectoryAtURL is a synchronous wrapper around [NSFileProviderManager.ImportDomainFromDirectoryAtURLCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (fc NSFileProviderManagerClass) ImportDomainFromDirectoryAtURL(ctx context.Context, domain INSFileProviderDomain, url foundation.NSURL) error {
	done := make(chan error, 1)
	fc.ImportDomainFromDirectoryAtURLCompletionHandler(domain, url, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// AddDomain is a synchronous wrapper around [NSFileProviderManager.AddDomainCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (fc NSFileProviderManagerClass) AddDomain(ctx context.Context, domain INSFileProviderDomain) error {
	done := make(chan error, 1)
	fc.AddDomainCompletionHandler(domain, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// GetDomains is a synchronous wrapper around [NSFileProviderManager.GetDomainsWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (fc NSFileProviderManagerClass) GetDomains(ctx context.Context) ([]NSFileProviderDomain, error) {
	type result struct {
		val []NSFileProviderDomain
		err error
	}
	done := make(chan result, 1)
	fc.GetDomainsWithCompletionHandler(func(val *[]NSFileProviderDomain, err error) {
		var out []NSFileProviderDomain
		if val != nil {
			out = append(out, (*val)...)
		}
		done <- result{out, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// RemoveDomain is a synchronous wrapper around [NSFileProviderManager.RemoveDomainCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (fc NSFileProviderManagerClass) RemoveDomain(ctx context.Context, domain INSFileProviderDomain) error {
	done := make(chan error, 1)
	fc.RemoveDomainCompletionHandler(domain, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RemoveDomainMode is a synchronous wrapper around [NSFileProviderManager.RemoveDomainModeCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (fc NSFileProviderManagerClass) RemoveDomainMode(ctx context.Context, domain INSFileProviderDomain, mode NSFileProviderDomainRemovalMode) (*foundation.NSURL, error) {
	type result struct {
		val *foundation.NSURL
		err error
	}
	done := make(chan result, 1)
	fc.RemoveDomainModeCompletionHandler(domain, mode, func(val *foundation.NSURL, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// RemoveAllDomains is a synchronous wrapper around [NSFileProviderManager.RemoveAllDomainsWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (fc NSFileProviderManagerClass) RemoveAllDomains(ctx context.Context) error {
	done := make(chan error, 1)
	fc.RemoveAllDomainsWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// DisconnectWithReasonOptions is a synchronous wrapper around [NSFileProviderManager.DisconnectWithReasonOptionsCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NSFileProviderManager) DisconnectWithReasonOptions(ctx context.Context, localizedReason string, options NSFileProviderManagerDisconnectionOptions) error {
	done := make(chan error, 1)
	f.DisconnectWithReasonOptionsCompletionHandler(localizedReason, options, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Reconnect is a synchronous wrapper around [NSFileProviderManager.ReconnectWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NSFileProviderManager) Reconnect(ctx context.Context) error {
	done := make(chan error, 1)
	f.ReconnectWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// WaitForStabilization is a synchronous wrapper around [NSFileProviderManager.WaitForStabilizationWithCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NSFileProviderManager) WaitForStabilization(ctx context.Context) error {
	done := make(chan error, 1)
	f.WaitForStabilizationWithCompletionHandler(func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ClaimKnownFoldersLocalizedReason is a synchronous wrapper around [NSFileProviderManager.ClaimKnownFoldersLocalizedReasonCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NSFileProviderManager) ClaimKnownFoldersLocalizedReason(ctx context.Context, knownFolders INSFileProviderKnownFolderLocations, localizedReason string) error {
	done := make(chan error, 1)
	f.ClaimKnownFoldersLocalizedReasonCompletionHandler(knownFolders, localizedReason, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// ReleaseKnownFoldersLocalizedReason is a synchronous wrapper around [NSFileProviderManager.ReleaseKnownFoldersLocalizedReasonCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NSFileProviderManager) ReleaseKnownFoldersLocalizedReason(ctx context.Context, knownFolders NSFileProviderKnownFolders, localizedReason string) error {
	done := make(chan error, 1)
	f.ReleaseKnownFoldersLocalizedReasonCompletionHandler(knownFolders, localizedReason, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// GetServiceWithNameItemIdentifier is a synchronous wrapper around [NSFileProviderManager.GetServiceWithNameItemIdentifierCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NSFileProviderManager) GetServiceWithNameItemIdentifier(ctx context.Context, serviceName foundation.NSFileProviderServiceName, itemIdentifier NSFileProviderItemIdentifier) (*foundation.NSFileProviderService, error) {
	type result struct {
		val *foundation.NSFileProviderService
		err error
	}
	done := make(chan result, 1)
	f.GetServiceWithNameItemIdentifierCompletionHandler(serviceName, itemIdentifier, func(val *foundation.NSFileProviderService, err error) {
		done <- result{val, err}
	})
	select {
	case r := <-done:
		return r.val, r.err
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

// SignalErrorResolved is a synchronous wrapper around [NSFileProviderManager.SignalErrorResolvedCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NSFileProviderManager) SignalErrorResolved(ctx context.Context, error_ foundation.NSError) error {
	done := make(chan error, 1)
	f.SignalErrorResolvedCompletionHandler(error_, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}

// RequestDiagnosticCollectionForItemWithIdentifierErrorReason is a synchronous wrapper around [NSFileProviderManager.RequestDiagnosticCollectionForItemWithIdentifierErrorReasonCompletionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (f NSFileProviderManager) RequestDiagnosticCollectionForItemWithIdentifierErrorReason(ctx context.Context, itemIdentifier NSFileProviderItemIdentifier, errorReason foundation.NSError) error {
	done := make(chan error, 1)
	f.RequestDiagnosticCollectionForItemWithIdentifierErrorReasonCompletionHandler(itemIdentifier, errorReason, func(err error) {
		done <- err
	})
	select {
	case err := <-done:
		return err
	case <-ctx.Done():
		return ctx.Err()
	}
}
