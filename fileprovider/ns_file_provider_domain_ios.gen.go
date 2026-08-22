// Code generated from Apple documentation for FileProvider. DO NOT EDIT.
//go:build ios
// +build ios

package fileprovider

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// Returns a newly instantiated domain.
//
// identifier: A string that identifies the domain. The file provider extension can select
// any string to uniquely identify the domain, as long as it doesn’t contain
// the colon (:) or slash (/) symbols.
//
// displayName: The name for the domain that the system shows to the user.
//
// pathRelativeToDocumentStorage: A path relative to the file provider extension’s [documentStorageURL]
// that the system uses to store the domain’s content.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/init(identifier:displayName:pathRelativeToDocumentStorage:)
//
// [documentStorageURL]: https://developer.apple.com/documentation/FileProvider/NSFileProviderExtension/documentStorageURL
func (f NSFileProviderDomain) InitWithIdentifierDisplayNamePathRelativeToDocumentStorage(identifier NSFileProviderDomainIdentifier, displayName string, pathRelativeToDocumentStorage string) NSFileProviderDomain {
	rv := objc.Send[NSFileProviderDomain](f.ID, objc.Sel("initWithIdentifier:displayName:pathRelativeToDocumentStorage:"), objc.String(string(identifier)), objc.String(displayName), objc.String(pathRelativeToDocumentStorage))
	return rv
}

// The path of the domain’s subdirectory relative to the file provider’s
// shared container.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderDomain/pathRelativeToDocumentStorage
func (f NSFileProviderDomain) PathRelativeToDocumentStorage() string {
	rv := objc.Send[objc.ID](f.ID, objc.Sel("pathRelativeToDocumentStorage"))
	return foundation.NSStringFromID(rv).String()
}
