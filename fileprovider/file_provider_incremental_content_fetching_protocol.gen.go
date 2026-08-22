// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Support for fetching changes to the item’s content.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderIncrementalContentFetching
type NSFileProviderIncrementalContentFetching interface {
	objectivec.IObject

	// Asks the file provider for an update of the specified item.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderIncrementalContentFetching/fetchContents(for:version:usingExistingContentsAt:existingVersion:request:completionHandler:)
	FetchContentsForItemWithIdentifierVersionUsingExistingContentsAtURLExistingVersionRequestCompletionHandler(itemIdentifier NSFileProviderItemIdentifier, requestedVersion INSFileProviderItemVersion, existingContents foundation.NSURL, existingVersion INSFileProviderItemVersion, request INSFileProviderRequest, completionHandler URLFileProviderItemErrorHandler) foundation.Progress
}

// NSFileProviderIncrementalContentFetchingObject wraps an existing Objective-C object that conforms to the NSFileProviderIncrementalContentFetching protocol.
type NSFileProviderIncrementalContentFetchingObject struct {
	objectivec.Object
}

func (o NSFileProviderIncrementalContentFetchingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderIncrementalContentFetchingObjectFromID constructs a [NSFileProviderIncrementalContentFetchingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderIncrementalContentFetchingObjectFromID(id objc.ID) NSFileProviderIncrementalContentFetchingObject {
	return NSFileProviderIncrementalContentFetchingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the file provider for an update of the specified item.
//
// itemIdentifier: The identifier of the item this method should fetch from your remote
// server.
//
// requestedVersion: The new version. If this value is set, you must fetch the specified
// version, or fail with an error. If this value isn’t set, fetch the most
// recent version from your server.
//
// existingContents: A URL that points to the content of the currently cached item.
//
// existingVersion: The version of the currently cached item.
//
// request: An object that identifies the context of that request, such as the
// requesting app.
//
// completionHandler: A block that you call after downloading the update. Pass the following
// parameters:
//
// `fileContents`: A URL that points to the new contents, or `nil` if an error
// occurs. `item`: The item’s identifier, or `nil` if an error occurs. The
// item’s version must match the version of the content passed to the
// `fileContents` parameter. `error`: If an error occurs, this object contains
// information about the error; otherwise, it’s `nil`.
//
// # Return Value
//
// An item that tracks the extension’s progress.
//
// # Discussion
//
// Implement this method to optimize downloading changes. Rather than
// downloading an entire new copy of the item, you can download just the
// changes from your remote storage.
//
// The system can call this method when it has already cached an item and
// learns that a new version is available. It passes a copy of the locally
// stored item to the method. In your implementation of the method, download
// any updates, apply them to the existing content, and then pass the updated
// version to the completion handler.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderIncrementalContentFetching/fetchContents(for:version:usingExistingContentsAt:existingVersion:request:completionHandler:)
func (o NSFileProviderIncrementalContentFetchingObject) FetchContentsForItemWithIdentifierVersionUsingExistingContentsAtURLExistingVersionRequestCompletionHandler(itemIdentifier NSFileProviderItemIdentifier, requestedVersion INSFileProviderItemVersion, existingContents foundation.NSURL, existingVersion INSFileProviderItemVersion, request INSFileProviderRequest, completionHandler URLFileProviderItemErrorHandler) foundation.Progress {
	_block5, _ := NewURLFileProviderItemErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](o.ID, objc.Sel("fetchContentsForItemWithIdentifier:version:usingExistingContentsAtURL:existingVersion:request:completionHandler:"), itemIdentifier, requestedVersion, existingContents, existingVersion, request, _block5)
	return foundation.NSProgressFromID(rv)
}
