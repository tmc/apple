// Code generated from Apple documentation for FileProvider. DO NOT EDIT.
//go:build ios
// +build ios

package fileprovider

import (
	"errors"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// Returns a placeholder URL for a given document URL.
//
// url: The document URL to be converted.
//
// # Return Value
//
// A placeholder URL for the given document.
//
// # Discussion
//
// This method maps file URLs into their corresponding placeholder URLs. You
// typically call this method to generate the placeholder URL before calling
// [writePlaceholder(at:withMetadata:)].
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/placeholderURL(for:)
//
// [writePlaceholder(at:withMetadata:)]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/writePlaceholder(at:withMetadata:)
func (_NSFileProviderManagerClass NSFileProviderManagerClass) PlaceholderURLForURL(url foundation.NSURL) foundation.NSURL {
	rv := objc.Send[objc.ID](objc.ID(_NSFileProviderManagerClass.class), objc.Sel("placeholderURLForURL:"), url)
	return foundation.NSURLFromID(rv)
}

// Writes a document placeholder with the provided metadata.
//
// placeholderURL: The placeholder URL for the document. You can generate a placeholder URL
// from a document URL by calling [placeholderURL(for:)].
//
// metadata: The metadata for this document.
//
// # Discussion
//
// Call this method whenever you need to create a placeholder for a document.
// The metadata that you provide sets the data provided to the user in the
// browser interface.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/writePlaceholder(at:withMetadata:)
//
// [placeholderURL(for:)]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/placeholderURL(for:)
func (_NSFileProviderManagerClass NSFileProviderManagerClass) WritePlaceholderAtURLWithMetadataError(placeholderURL foundation.NSURL, metadata NSFileProviderItem) (bool, error) {
	var errorPtr objc.ID
	rv := objc.Send[bool](objc.ID(_NSFileProviderManagerClass.class), objc.Sel("writePlaceholderAtURL:withMetadata:error:"), placeholderURL, metadata, unsafe.Pointer(&errorPtr))
	if errorPtr != 0 {
		objc.Send[objc.ID](errorPtr, objc.Sel("retain"))
		return false, foundation.NSErrorFrom(errorPtr)
	}
	if !rv {
		return false, errors.New("writePlaceholderAtURL:withMetadata:error: returned NO with nil NSError")
	}
	return rv, nil

}

// The root URL for all shared documents.
//
// # Discussion
//
// This property contains the URL `/File Provider Storage`, where container
// URL is the value returned by the
// [containerURL(forSecurityApplicationGroupIdentifier:)] method.
//
// The container URL refers to an app group container directory used by the
// [NSFileProviderExtension] extension. You can specify this shared container
// using the [NSExtensionFileProviderDocumentGroup] key in the File Provider
// extension’s `info.Plist()` file.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/documentStorageURL
//
// [containerURL(forSecurityApplicationGroupIdentifier:)]: https://developer.apple.com/documentation/Foundation/FileManager/containerURL(forSecurityApplicationGroupIdentifier:)
func (f NSFileProviderManager) DocumentStorageURL() foundation.NSURL {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("documentStorageURL"))
	return foundation.NSURLFromID(rv)
}

// A purpose identifier for coordinated reads and writes.
//
// # Discussion
//
// This property contains a unique string that can be used as a purpose
// identifier for file coordination. The File Provider extension should use
// this identifier when performing coordinated reads and writes, to help
// prevent deadlocks.
//
// Pass this identifier to the file coordinator’s “ method before
// performing a coordinated read or write.
//
// This method returns the containing app’s bundle identifier.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/providerIdentifier
func (f NSFileProviderManager) ProviderIdentifier() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("providerIdentifier"))
	return foundation.NSStringFromID(rv).String()
}

// A property that returns the shared file provider manager object.
//
// # Discussion
//
// This property returns a manager for the default domain on iOS. You can
// access the default domain in both the containing app and the File Provider
// extension. On macOS, use an explicit domain by calling
// [NSFileProviderManagerClass.ManagerForDomain] instead.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderManager/default
func (_NSFileProviderManagerClass NSFileProviderManagerClass) DefaultManager() NSFileProviderManager {
	rv := objc.Send[objc.ID](objc.ID(_NSFileProviderManagerClass.class), objc.Sel("defaultManager"))
	return NSFileProviderManagerFromID(rv)
}
