// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that defines the interface for sharing known-folder locations with the system.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderSupporting
type NSFileProviderKnownFolderSupporting interface {
	objectivec.IObject

	// Requests suitable locations for known folders.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderSupporting/getKnownFolderLocations(_:completionHandler:)
	GetKnownFolderLocationsCompletionHandler(knownFolders NSFileProviderKnownFolders, completionHandler FileProviderKnownFolderLocationsErrorHandler)
}

// NSFileProviderKnownFolderSupportingObject wraps an existing Objective-C object that conforms to the NSFileProviderKnownFolderSupporting protocol.
type NSFileProviderKnownFolderSupportingObject struct {
	objectivec.Object
}

func (o NSFileProviderKnownFolderSupportingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderKnownFolderSupportingObjectFromID constructs a [NSFileProviderKnownFolderSupportingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderKnownFolderSupportingObjectFromID(id objc.ID) NSFileProviderKnownFolderSupportingObject {
	return NSFileProviderKnownFolderSupportingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Requests suitable locations for known folders.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderKnownFolderSupporting/getKnownFolderLocations(_:completionHandler:)
func (o NSFileProviderKnownFolderSupportingObject) GetKnownFolderLocationsCompletionHandler(knownFolders NSFileProviderKnownFolders, completionHandler FileProviderKnownFolderLocationsErrorHandler) {
	_block1, _ := NewFileProviderKnownFolderLocationsErrorBlock(completionHandler)
	objc.Send[struct{}](o.ID, objc.Sel("getKnownFolderLocations:completionHandler:"), knownFolders, _block1)
}
