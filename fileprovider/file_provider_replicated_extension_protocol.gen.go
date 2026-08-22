// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A File Provider extension in which the system replicates the contents on disk.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderReplicatedExtension
type NSFileProviderReplicatedExtension interface {
	objectivec.IObject
	NSFileProviderEnumerating

	// Tells the file provider to perform any necessary cleanup so that the system can deallocate it.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderReplicatedExtension/invalidate()
	Invalidate()

	// Asks the file provider for the metadata of the provided item.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderReplicatedExtension/item(for:request:completionHandler:)
	ItemForIdentifierRequestCompletionHandler(identifier NSFileProviderItemIdentifier, request INSFileProviderRequest, completionHandler FileProviderItemErrorHandler) foundation.Progress

	// Tells the file provider to download the requested item from remote storage.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderReplicatedExtension/fetchContents(for:version:request:completionHandler:)
	FetchContentsForItemWithIdentifierVersionRequestCompletionHandler(itemIdentifier NSFileProviderItemIdentifier, requestedVersion INSFileProviderItemVersion, request INSFileProviderRequest, completionHandler URLFileProviderItemErrorHandler) foundation.Progress

	// Tells the file provider to create or import an item based on a template.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderReplicatedExtension/createItem(basedOn:fields:contents:options:request:completionHandler:)
	CreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler(itemTemplate NSFileProviderItem, fields NSFileProviderItemFields, url foundation.NSURL, options NSFileProviderCreateItemOptions, request INSFileProviderRequest, completionHandler FileProviderItemNSFileProviderItemFieldsBoolErrorHandler) foundation.Progress

	// Tells the file provider that an item’s content or metadata changed.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderReplicatedExtension/modifyItem(_:baseVersion:changedFields:contents:options:request:completionHandler:)
	ModifyItemBaseVersionChangedFieldsContentsOptionsRequestCompletionHandler(item NSFileProviderItem, version INSFileProviderItemVersion, changedFields NSFileProviderItemFields, newContents foundation.NSURL, options NSFileProviderModifyItemOptions, request INSFileProviderRequest, completionHandler FileProviderItemNSFileProviderItemFieldsBoolErrorHandler) foundation.Progress

	// Tells the file provider to delete an item forever.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderReplicatedExtension/deleteItem(identifier:baseVersion:options:request:completionHandler:)
	DeleteItemWithIdentifierBaseVersionOptionsRequestCompletionHandler(identifier NSFileProviderItemIdentifier, version INSFileProviderItemVersion, options NSFileProviderDeleteItemOptions, request INSFileProviderRequest, completionHandler ErrorHandler) foundation.Progress
}

// NSFileProviderReplicatedExtensionObject wraps an existing Objective-C object that conforms to the NSFileProviderReplicatedExtension protocol.
type NSFileProviderReplicatedExtensionObject struct {
	objectivec.Object
}

func (o NSFileProviderReplicatedExtensionObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderReplicatedExtensionObjectFromID constructs a [NSFileProviderReplicatedExtensionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderReplicatedExtensionObjectFromID(id objc.ID) NSFileProviderReplicatedExtensionObject {
	return NSFileProviderReplicatedExtensionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Tells the file provider to perform any necessary cleanup so that the system
// can deallocate it.
//
// # Discussion
//
// Your implementation should perform any necessary cleanup so that the system
// can dismiss and deallocate the file provider.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderReplicatedExtension/invalidate()
func (o NSFileProviderReplicatedExtensionObject) Invalidate() {
	objc.Send[struct{}](o.ID, objc.Sel("invalidate"))
}

// Asks the file provider for the metadata of the provided item.
//
// identifier: The item’s identifier.
//
// request: An object that identifies the context of that request, such as the
// requesting app.
//
// completionHandler: A block that you call after downloading the item’s metadata. The block
// takes the following parameters:
//
// `item`: A new file provider item, containing all the item’s metadata.
// `error`: If an error occurs, this object contains information about the
// error; otherwise, it’s `nil`.
//
// # Return Value
//
// An item that tracks your extension’s progress. The system automatically
// calls [cancel()] on the progress object when an error occurs.
//
// # Discussion
//
// If your extension doesn’t recognize the item, pass
// [NSFileProviderError.Code.noSuchItem] to the handler. The system assumes
// the item is no longer in the domain, and attempts to delete the local copy.
// If the delete attempt fails because the item has local changes, the system
// reimports the item by calling
// [CreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler].
//
// If you pass [NSFileProviderError.Code.notAuthenticated] or
// [NSFileProviderError.Code.serverUnreachable] to the handler, the system
// presents an appropriate alert to the user, but doesn’t try to access the
// metadata until triggered again by the user.
//
// The system considers any other errors to be transient, and automatically
// retries the method call.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderReplicatedExtension/item(for:request:completionHandler:)
//
// [cancel()]: https://developer.apple.com/documentation/Foundation/Progress/cancel()
// [NSFileProviderError.Code.noSuchItem]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/noSuchItem
// [NSFileProviderError.Code.notAuthenticated]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/notAuthenticated
// [NSFileProviderError.Code.serverUnreachable]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/serverUnreachable
func (o NSFileProviderReplicatedExtensionObject) ItemForIdentifierRequestCompletionHandler(identifier NSFileProviderItemIdentifier, request INSFileProviderRequest, completionHandler FileProviderItemErrorHandler) foundation.Progress {
	_block2, _ := NewFileProviderItemErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](o.ID, objc.Sel("itemForIdentifier:request:completionHandler:"), identifier, request, _block2)
	return foundation.NSProgressFromID(rv)
}

// Tells the file provider to download the requested item from remote storage.
//
// itemIdentifier: The item to fetch.
//
// requestedVersion: The version of the item. If this is `nil`, download the latest version.
//
// request: An object that identifies the context of that request, such as the
// requesting app.
//
// completionHandler: A block that you call after downloading the item from your remote storage.
// You pass the following parameters:
//
// `fileContents`: A URL to the item’s contents. `item`: The downloaded
// item. `error`: If an error occurs, this object contains information about
// the error; otherwise, it’s `nil`.
//
// # Return Value
//
// An item that tracks your extension’s progress. The system automatically
// calls [cancel()] on the progress object when an error occurs, or when the
// system or user cancels the download.
//
// # Discussion
//
// The system initially learns about available items through enumerations;
// however, the enumeration only provides the item’s metadata. When the user
// accesses the item, the system needs to download the full contents from your
// remote store. After you call the completion handler, the system takes
// complete control over the local copy.
//
// If the `requestedVersion` parameter is not `nil`, you must return the
// specified version of the item, or return an error. If the parameter is
// `nil`, return a version that is the same or newer than the most recent
// version enumerated to the system. In either case, the
// [NSFileProviderItemVersion.ContentVersion] of the [NSFileProviderItem]
// passed to the completion handler must match the version you return.
//
// If your extension doesn’t recognize the item, pass
// [NSFileProviderError.Code.noSuchItem] to the handler. The system assumes
// the item is no longer in the domain, and attempts to delete the local copy.
// If the delete attempt fails because the item has local changes, the system
// reimports the item by calling
// [CreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler].
//
// If you pass [NSFileProviderError.Code.notAuthenticated] or
// [NSFileProviderError.Code.serverUnreachable] to the handler, the system
// presents an appropriate alert to the user, but doesn’t try to access the
// metadata until triggered again by the user.
//
// If the user deletes the item before the download completes, the system
// calls the progress object’s [cancel()] method. Your file provider should
// stop fetching the item, and pass [NSUserCancelledError] to the handler.
//
// The system considers any other errors to be transient, and automatically
// retries the method call.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderReplicatedExtension/fetchContents(for:version:request:completionHandler:)
//
// [cancel()]: https://developer.apple.com/documentation/Foundation/Progress/cancel()
// [NSFileProviderError.Code.noSuchItem]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/noSuchItem
// [NSFileProviderError.Code.notAuthenticated]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/notAuthenticated
// [NSFileProviderError.Code.serverUnreachable]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/serverUnreachable
// [NSFileProviderItem]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItem-swift.typealias
// [NSUserCancelledError]: https://developer.apple.com/documentation/Foundation/NSUserCancelledError-swift.var
//
// [cancel()]: https://developer.apple.com/documentation/Foundation/Progress/cancel()
func (o NSFileProviderReplicatedExtensionObject) FetchContentsForItemWithIdentifierVersionRequestCompletionHandler(itemIdentifier NSFileProviderItemIdentifier, requestedVersion INSFileProviderItemVersion, request INSFileProviderRequest, completionHandler URLFileProviderItemErrorHandler) foundation.Progress {
	_block3, _ := NewURLFileProviderItemErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](o.ID, objc.Sel("fetchContentsForItemWithIdentifier:version:request:completionHandler:"), itemIdentifier, requestedVersion, request, _block3)
	return foundation.NSProgressFromID(rv)
}

// Tells the file provider to create or import an item based on a template.
//
// itemTemplate: An object that defines the state of the new or imported item.
//
// fields: The fields that you should apply to the new or imported item.
//
// url: If the item is a file with the [NSFileProviderItemContents] field set, this
// is the URL to the item’s content. Otherwise, it’s `nil`.
//
// options: The item creation options.
//
// request: An object that identifies the context of that request, such as the
// requesting app.
//
// completionHandler: A block that you call after uploading the item to your remote storage. You
// pass the following parameters:
//
// `createdItem`: The newly created item. `stillPendingFields`: Any fields
// that you haven’t yet applied. If you can apply all the fields at once,
// pass an empty [NSFileProviderItemField] instance. `shouldFetchContent`: A
// Boolean value that indicates whether the system should fetch the item’s
// content from your remote storage. `error`: If an error occurs, this object
// contains information about the error; otherwise, it’s `nil`.
//
// # Return Value
//
// A progress that tracks creating the item in your remote storage and
// uploading its content. The system automatically calls [cancel()] on the
// progress object when an error occurs.
//
// # Discussion
//
// The system calls this method when the user creates a new item or imports an
// item into the file provider. The system manages the local copy of the item.
// You’re responsible for syncing and saving the item to your remote
// storage.
//
// Implement this method to create an item in your remote storage that matches
// the template, and then call the callback handler.
//
// The `itemTemplate` parameter describes the item’s intended state,
// including:
//
// [Filename]: The item’s name. [ContentType]: The item’s type. The item
// can be a file, directory, symlink, or alias. [UTTypeFolder],
// [UTTypeSymbolicLink], and [UTTypeAliasFile] types typically need special
// handling. [ParentItemIdentifier]: The item’s location.
//
// The system sets the template’s [ItemIdentifier] to a unique value and
// guarantees that it remains the same for the specified item. For example,
// the system can reuse the identifier to replay this method after a crash.
//
// In general, set the properties in your `createdItem` to match the
// `itemTemplate`. One exception is the [ItemIdentifier] property; always
// provide your own identifier for the item. If you reuse an existing
// identifier, the system replaces the local copy of the old item with the new
// one.
//
// If the item is a document, fetch its contents from the `url` parameter.
// Otherwise, the `url` is `nil`. For symlinks, you can access the content
// using the template’s [SymlinkTargetPath] parameter. For both symlinks and
// aliases, make sure to return the correct UTI for the item, because the UTI
// can’t be inferred from the item’s filename.
//
// If you are reimporting an item and the system finds a local copy without
// any content, it sets the [NSFileProviderCreateItemMayAlreadyExist] option,
// and sets the `url` to nil. In this case, if you can’t match the item with
// an existing item from remote storage, pass `nil` as the completion
// handler’s `createdItem` parameter. The system then deletes the local copy
// of the item.
//
// If the attempt to create an item fails because the parent directory
// doesn’t exist, pass [NSFileProviderError.Code.noSuchItem] to the handler.
// The system attempts to create the parent directory, and then tries to
// create the item again.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderReplicatedExtension/createItem(basedOn:fields:contents:options:request:completionHandler:)
//
// [cancel()]: https://developer.apple.com/documentation/Foundation/Progress/cancel()
// [NSFileProviderError.Code.noSuchItem]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/noSuchItem
// [UTTypeAliasFile]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeAliasFile
// [UTTypeFolder]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeFolder
// [UTTypeSymbolicLink]: https://developer.apple.com/documentation/UniformTypeIdentifiers/UTTypeSymbolicLink
func (o NSFileProviderReplicatedExtensionObject) CreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler(itemTemplate NSFileProviderItem, fields NSFileProviderItemFields, url foundation.NSURL, options NSFileProviderCreateItemOptions, request INSFileProviderRequest, completionHandler FileProviderItemNSFileProviderItemFieldsBoolErrorHandler) foundation.Progress {
	_block5, _ := NewFileProviderItemNSFileProviderItemFieldsBoolErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](o.ID, objc.Sel("createItemBasedOnTemplate:fields:contents:options:request:completionHandler:"), itemTemplate, fields, url, options, request, _block5)
	return foundation.NSProgressFromID(rv)
}

// Tells the file provider that an item’s content or metadata changed.
//
// item: The item to modify.
//
// version: The item’s version.
//
// changedFields: The fields that have changed.
//
// newContents: A URL for the local copy of the item’s new contents.
//
// options: The modification options.
//
// request: An object that identifies the context of that request, such as the
// requesting app.
//
// completionHandler: A block that you call after uploading the changes to your remote storage.
// You pass the following parameters:
//
// `item`: The newly modified item. `stillPendingFields`: Any fields that you
// haven’t yet applied. If you can apply all the fields at once, pass an
// empty [NSFileProviderItemField] instance. `shouldFetchContent`: A Boolean
// value that indicates whether the system should fetch the item’s content
// from your remote storage. `error`: If an error occurs, this object contains
// information about the error; otherwise, it’s `nil`.
//
// # Return Value
//
// An item that tracks your extension’s progress.
//
// # Discussion
//
// The system calls this method when the user modifies an item—for example,
// moving it, renaming it, or updating its content. The `changedFields`
// parameter may contain multiple items, indicating that multiple changes have
// occurred. Update the version of the item in your remote storage to match,
// and then call the callback handler.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderReplicatedExtension/modifyItem(_:baseVersion:changedFields:contents:options:request:completionHandler:)
func (o NSFileProviderReplicatedExtensionObject) ModifyItemBaseVersionChangedFieldsContentsOptionsRequestCompletionHandler(item NSFileProviderItem, version INSFileProviderItemVersion, changedFields NSFileProviderItemFields, newContents foundation.NSURL, options NSFileProviderModifyItemOptions, request INSFileProviderRequest, completionHandler FileProviderItemNSFileProviderItemFieldsBoolErrorHandler) foundation.Progress {
	_block6, _ := NewFileProviderItemNSFileProviderItemFieldsBoolErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](o.ID, objc.Sel("modifyItem:baseVersion:changedFields:contents:options:request:completionHandler:"), item, version, changedFields, newContents, options, request, _block6)
	return foundation.NSProgressFromID(rv)
}

// Tells the file provider to delete an item forever.
//
// identifier: The identifier of the object to delete.
//
// version: The version of the item to delete.
//
// options: The options for deleting the item.
//
// request: An object that identifies the context of that request, such as the
// requesting app.
//
// completionHandler: A block that you call after deleting the item from your remote storage. You
// pass the following parameter:
//
// `error`: If an error occurs, this object contains information about the
// error; otherwise, it’s `nil`.
//
// # Return Value
//
// An item that tracks your extension’s progress.
//
// # Discussion
//
// The system calls this method when the user deletes an item that was already
// in the trash. Users can only delete items that have the
// [NSFileProviderItemCapabilitiesAllowsDeleting] capability.
//
// Remove the item from the trash and delete it from your remote storage. If
// the item is in the working set, notify the system about the change by
// calling
// [NSFileProviderManager.SignalEnumeratorForContainerItemIdentifierCompletionHandler]
// and passing [workingSet] for the `containerItemIdentifier` parameter. If
// the deletion is recursive, be sure to check all the deleted items, and
// notify the system to any changes in the working set.
//
// If your extension doesn’t recognize the item, you can just report
// success. The system then removes the local copy of the item.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderReplicatedExtension/deleteItem(identifier:baseVersion:options:request:completionHandler:)
//
// [workingSet]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemIdentifier/workingSet
func (o NSFileProviderReplicatedExtensionObject) DeleteItemWithIdentifierBaseVersionOptionsRequestCompletionHandler(identifier NSFileProviderItemIdentifier, version INSFileProviderItemVersion, options NSFileProviderDeleteItemOptions, request INSFileProviderRequest, completionHandler ErrorHandler) foundation.Progress {
	_block4, _ := NewErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](o.ID, objc.Sel("deleteItemWithIdentifier:baseVersion:options:request:completionHandler:"), identifier, version, options, request, _block4)
	return foundation.NSProgressFromID(rv)
}

// Tells the file provider that the set of materialized items changed.
//
// completionHandler: A block that you call after you finish processing the changes.
//
// # Discussion
//
// The system calls this method when the set of materialized items changes,
// such as when the system downloads the content of a dataless item, or
// deletes the contents of a materialized item.
//
// Keep track of the items in the working set—the set of items that the
// system has downloaded and is managing on disk—to optimize updates. Your
// file provider must let the system know about any changes to the working
// set. If you don’t track the working set, then you need to let the system
// know about any changes to any items in the remote storage.
//
// For more information, see [Synchronizing the File Provider Extension].
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderReplicatedExtension/materializedItemsDidChange(completionHandler:)
//
// [Synchronizing the File Provider Extension]: https://developer.apple.com/documentation/FileProvider/synchronizing-the-file-provider-extension
func (o NSFileProviderReplicatedExtensionObject) MaterializedItemsDidChangeWithCompletionHandler(completionHandler VoidHandler) {
	_block0, _ := NewVoidBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("materializedItemsDidChangeWithCompletionHandler:"), _block0)
}

// Tells the file provider extension that the set of pending items has
// changed.
//
// completionHandler: A block that you call after you finish processing the changes.
//
// # Discussion
//
// The system calls this method whenever the set of pending items changes. It
// updates the pending set regularly, but only when there are meaningful
// changes, such as:
//
// - New items are now pending. - The system has successfully synced one or
// more pending items. - The domain version changed when the pending item set
// wasn’t empty.
//
// To enumerate the pending set, create an object that adopts the
// [NSFileProviderEnumerationObserver] and [NSFileProviderChangeObserver]
// protocols. Then pass this item to the
// [NSFileProviderManager.EnumeratorForPendingItems] method on a
// [NSFileProviderManager] instance for your domain. The system then calls
// your observer object’s methods when the pending set changes.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderReplicatedExtension/pendingItemsDidChange(completionHandler:)
func (o NSFileProviderReplicatedExtensionObject) PendingItemsDidChangeWithCompletionHandler(completionHandler VoidHandler) {
	_block0, _ := NewVoidBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("pendingItemsDidChangeWithCompletionHandler:"), _block0)
}

// Tells the File Provider extension that the system finished importing items.
//
// completionHandler: A block that your implementation must call.
//
// # Discussion
//
// The system calls this method after importing on-disk items. You can trigger
// an import by calling either
// [NSFileProviderManager.ReimportItemsBelowItemWithIdentifierCompletionHandler]
// or
// [NSFileProviderManagerClass.ImportDomainFromDirectoryAtURLCompletionHandler].
// The system can also initiate its own imports as needed.
//
// During the import, the system calls your File Provider extension’s
// [CreateItemBasedOnTemplateFieldsContentsOptionsRequestCompletionHandler]
// method and passes the [NSFileProviderCreateItemMayAlreadyExist] option.
// Check to see if the item already exists in your remote storage—uploading
// it if necessary.
//
// After importing all the items, the system calls your
// [ImportDidFinishWithCompletionHandler] method. Handle any necessary cleanup
// operations, and then call the completion handler.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderReplicatedExtension/importDidFinish(completionHandler:)
func (o NSFileProviderReplicatedExtensionObject) ImportDidFinishWithCompletionHandler(completionHandler VoidHandler) {
	_block0, _ := NewVoidBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("importDidFinishWithCompletionHandler:"), _block0)
}

// Tells the file provider to return an enumerator for the provided directory.
//
// containerItemIdentifier: The item identifier for the directory.
//
// request: An object that identifies the context of that request, such as the
// requesting app.
//
// # Return Value
//
// An enumerator for the specified directory.
//
// # Discussion
//
// The system calls this method to request an enumerator for the specified
// item.
//
// Possible item identifiers include:
//
// [rootContainer]: The system passes this identifier when the user begins
// browsing your file provider’s content. A directory’s [ItemIdentifier]:
// The system requests a new enumerator each time the user opens a new
// directory. [workingSet]: The system can request an enumerator so that it
// can sync the working set in the background. A document’s
// [ItemIdentifier]: The system subscribes to live updates by requesting an
// enumerator for a document. The [trashContainer] directory.: The system
// passes this identifier when the user browses the contents of the trash. If
// your File Provider extension doesn’t support moving items to the trash,
// your implementation should throw or return an error.
//
// Your implementation should create and return an [NSFileProviderEnumerator]
// object that provides the requested content.
//
// # Handle Errors
//
// If you can’t return the requested enumerator, you must throw an error in
// Swift, or if you return nil in Objective-C, you must set the `error` out
// parameter.
//
// If the `containerItemIdentifier` parameter is [trashContainer] and your
// extension doesn’t support trashing items, then it should fail with the
// [NSFeatureUnsupportedError] error code from the [NSCocoaErrorDomain]
// domain. Additionally, make sure the items managed by your File Provider
// extension don’t have the [NSFileProviderItemCapabilitiesAllowsTrashing]
// capability enabled.
//
// If the `containerItemIdentifier` parameter doesn’t exist in your remote
// storage, you should fail with an [NSFileProviderError.Code.noSuchItem]
// error. The system then attempts to delete the item from disk.
//
// If you pass [NSFileProviderError.Code.notAuthenticated] or
// [NSFileProviderError.Code.serverUnreachable] to the handler, the system
// presents an appropriate alert to the user, but doesn’t try to access the
// metadata until triggered again by the user.
//
// The system considers any other errors to be transient, and automatically
// retries the method call.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderEnumerating/enumerator(for:request:)
//
// [NSCocoaErrorDomain]: https://developer.apple.com/documentation/Foundation/NSCocoaErrorDomain
// [NSFeatureUnsupportedError]: https://developer.apple.com/documentation/Foundation/NSFeatureUnsupportedError-swift.var
// [NSFileProviderError.Code.noSuchItem]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/noSuchItem
// [NSFileProviderError.Code.notAuthenticated]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/notAuthenticated
// [NSFileProviderError.Code.serverUnreachable]: https://developer.apple.com/documentation/FileProvider/NSFileProviderError/Code/serverUnreachable
// [rootContainer]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemIdentifier/rootContainer
// [trashContainer]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemIdentifier/trashContainer
// [workingSet]: https://developer.apple.com/documentation/FileProvider/NSFileProviderItemIdentifier/workingSet
func (o NSFileProviderReplicatedExtensionObject) EnumeratorForContainerItemIdentifierRequestError(containerItemIdentifier NSFileProviderItemIdentifier, request INSFileProviderRequest) (NSFileProviderEnumerator, error) {
	rv, err := objc.SendWithError[objc.ID](o.ID, objc.Sel("enumeratorForContainerItemIdentifier:request:error:"), objc.String(string(containerItemIdentifier)), request)
	if err != nil {
		return nil, err
	}
	return NSFileProviderEnumeratorObjectFromID(rv), nil
}
