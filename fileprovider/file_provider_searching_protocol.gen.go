// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol you implement to support searching in your file provider.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearching
type NSFileProviderSearching interface {
	objectivec.IObject

	// Provides an object that enumerates over search results, in response to a call from the system.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearching/searchEnumerator(for:)
	SearchEnumeratorForStringSearchRequest(request INSFileProviderStringSearchRequest) NSFileProviderSearchEnumerator
}

// NSFileProviderSearchingObject wraps an existing Objective-C object that conforms to the NSFileProviderSearching protocol.
type NSFileProviderSearchingObject struct {
	objectivec.Object
}

func (o NSFileProviderSearchingObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderSearchingObjectFromID constructs a [NSFileProviderSearchingObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderSearchingObjectFromID(id objc.ID) NSFileProviderSearchingObject {
	return NSFileProviderSearchingObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// Provides an object that enumerates over search results, in response to a
// call from the system.
//
// request: An [NSFileProviderStringSearchRequest] that contains the search query.
//
// # Return Value
//
// An [NSFileProviderSearchEnumerator] that you implement to provide search
// results to the system.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearching/searchEnumerator(for:)
func (o NSFileProviderSearchingObject) SearchEnumeratorForStringSearchRequest(request INSFileProviderStringSearchRequest) NSFileProviderSearchEnumerator {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("searchEnumeratorForStringSearchRequest:"), request)
	return NSFileProviderSearchEnumeratorObjectFromID(rv)
}
