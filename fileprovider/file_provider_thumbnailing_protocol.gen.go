// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// Support for item thumbnails.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderThumbnailing
type NSFileProviderThumbnailing interface {
	objectivec.IObject

	// Asks the file provider for a thumbnail of the specified items.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderThumbnailing/fetchThumbnails(for:requestedSize:perThumbnailCompletionHandler:completionHandler:)
	FetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler(itemIdentifiers []string, size corefoundation.CGSize, perThumbnailCompletionHandler StringDataErrorHandler, completionHandler ErrorHandler) foundation.Progress
}

// NSFileProviderThumbnailingObject wraps an existing Objective-C object that conforms to the NSFileProviderThumbnailing protocol.
type NSFileProviderThumbnailingObject struct {
	objectivec.Object
}

func (o NSFileProviderThumbnailingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderThumbnailingObjectFromID constructs a [NSFileProviderThumbnailingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderThumbnailingObjectFromID(id objc.ID) NSFileProviderThumbnailingObject {
	return NSFileProviderThumbnailingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Asks the file provider for a thumbnail of the specified items.
//
// itemIdentifiers: The identifiers of the specified items.
//
// size: The size for the thumbnail image.
//
// perThumbnailCompletionHandler: A block that you call once for each item in the `itemIdentifiers` array.
// Pass the following parameters:
//
// `identifier`: The identifier of the item. `imageData`: A data object
// containing the thumbnail, or `nil` if an error occurred. This data object
// must be in an image format supported by [Image I/O]. `error`: If an error
// occurs, this object contains information about the error; otherwise, it’s
// `nil`.
//
// completionHandler: A block that you call after returning a thumbnail for each item. Pass the
// following parameters:
//
// `error`: If an error occurs, this object contains information about the
// error; otherwise, it’s `nil`.
//
// # Return Value
//
// An object that reports the progress of this request. If the thumbnails are
// no longer needed, the system can cancel this request by calling the
// progress object’s [cancel()] method.
//
// # Discussion
//
// For local files, the system automatically provides thumbnails for supported
// content types, and calls a Quick Look Preview extension to get thumbnails
// for custom types.
//
// However, the system can’t generate thumbnails for remote items. Instead,
// it calls this method to request thumbnails for items stored on a remote
// sever. The system caches these thumbnails, and only requests thumbnails for
// new items as needed. The system caches thumbnails based on the item’s
// [ItemVersion] property. To update an item’s thumbnail, update the
// item’s version identifier.
//
// In your implementation, call the `perThumbnailCompletionHandler` block once
// for each item in the `itemIdentifiers` array. Call the `completionHandler`
// block only after calling `perThumbnailCompletionHandler` for each item.
//
// For the best performance, use PNG images for text and vector graphics, JPEG
// for nontransparent photographs, and JPEG2000 for photographs with
// transparent backgrounds; however, you can use any image formats supported
// by [NSImage] and [UIImage]. You can also select a different image type for
// different thumbnails.
//
// If a global error occurs, you don’t have to call
// `perThumbnailCompletionHandler` on each item. Instead, call the
// `completionHandler` block, and pass in the global error. The system applies
// this error to all outstanding items.
//
// If a given item doesn’t have a thumbnail, call the
// `perThumbnailCompletionHandler` block and pass `nil` for both the
// `imageData` and `error` parameters.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderThumbnailing/fetchThumbnails(for:requestedSize:perThumbnailCompletionHandler:completionHandler:)
//
// [Image I/O]: https://developer.apple.com/documentation/ImageIO
// [cancel()]: https://developer.apple.com/documentation/Foundation/Progress/cancel()
// [NSImage]: https://developer.apple.com/documentation/AppKit/NSImage
// [UIImage]: https://developer.apple.com/documentation/UIKit/UIImage
func (o NSFileProviderThumbnailingObject) FetchThumbnailsForItemIdentifiersRequestedSizePerThumbnailCompletionHandlerCompletionHandler(itemIdentifiers []string, size corefoundation.CGSize, perThumbnailCompletionHandler StringDataErrorHandler, completionHandler ErrorHandler) foundation.Progress {
	_block2, _ := NewStringDataErrorBlock(perThumbnailCompletionHandler)
	_block3, _ := NewErrorBlock(completionHandler)
	rv := objc.Send[objc.ID](o.ID, objc.Sel("fetchThumbnailsForItemIdentifiers:requestedSize:perThumbnailCompletionHandler:completionHandler:"), itemIdentifiers, size, _block2, _block3)
	return foundation.NSProgressFromID(rv)
}
