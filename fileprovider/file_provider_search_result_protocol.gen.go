// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
	"github.com/tmc/apple/uniformtypeidentifiers"
)

// A protocol that defines properties of a search result.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchResult
type NSFileProviderSearchResult interface {
	objectivec.IObject

	// The identifier for this search result.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchResult/itemIdentifier
	ItemIdentifier() NSFileProviderItemIdentifier

	// The result’s file name.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchResult/filename
	Filename() string

	// The result file’s creation date.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchResult/creationDate
	CreationDate() foundation.NSDate

	// The result file’s content modification date.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchResult/contentModificationDate
	ContentModificationDate() foundation.NSDate

	// The result file’s last-used date.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchResult/lastUsedDate
	LastUsedDate() foundation.NSDate

	// The result file’s content type.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchResult/contentType
	ContentType() uniformtypeidentifiers.UTType

	// The result file’s size.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchResult/documentSize
	DocumentSize() foundation.NSNumber
}

// NSFileProviderSearchResultObject wraps an existing Objective-C object that conforms to the NSFileProviderSearchResult protocol.
type NSFileProviderSearchResultObject struct {
	objectivec.Object
}

func (o NSFileProviderSearchResultObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderSearchResultObjectFromID constructs a [NSFileProviderSearchResultObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderSearchResultObjectFromID(id objc.ID) NSFileProviderSearchResultObject {
	return NSFileProviderSearchResultObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The identifier for this search result.
//
// # Discussion
//
// Choose an identifier that’s usable with API calls from the
// [NSFileProviderReplicatedExtension] protocol.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchResult/itemIdentifier
func (o NSFileProviderSearchResultObject) ItemIdentifier() NSFileProviderItemIdentifier {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("itemIdentifier"))
	return NSFileProviderItemIdentifier(foundation.NSStringFromID(rv).String())
}

// The result’s file name.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchResult/filename
func (o NSFileProviderSearchResultObject) Filename() string {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("filename"))
	return foundation.NSStringFromID(rv).String()
}

// The result file’s creation date.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchResult/creationDate
func (o NSFileProviderSearchResultObject) CreationDate() foundation.NSDate {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("creationDate"))
	return foundation.NSDateFromID(rv)
}

// The result file’s content modification date.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchResult/contentModificationDate
func (o NSFileProviderSearchResultObject) ContentModificationDate() foundation.NSDate {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("contentModificationDate"))
	return foundation.NSDateFromID(rv)
}

// The result file’s last-used date.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchResult/lastUsedDate
func (o NSFileProviderSearchResultObject) LastUsedDate() foundation.NSDate {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("lastUsedDate"))
	return foundation.NSDateFromID(rv)
}

// The result file’s content type.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchResult/contentType
func (o NSFileProviderSearchResultObject) ContentType() uniformtypeidentifiers.UTType {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("contentType"))
	return uniformtypeidentifiers.UTTypeFromID(rv)
}

// The result file’s size.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderSearchResult/documentSize
func (o NSFileProviderSearchResultObject) DocumentSize() foundation.NSNumber {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("documentSize"))
	return foundation.NSNumberFromID(rv)
}
