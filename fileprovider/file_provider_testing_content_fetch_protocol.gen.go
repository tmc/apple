// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An operation that fetches an item’s content.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingContentFetch
type NSFileProviderTestingContentFetch interface {
	objectivec.IObject
	NSFileProviderTestingOperation

	// The containing item’s unique identifier.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingContentFetch/itemIdentifier
	ItemIdentifier() NSFileProviderItemIdentifier

	// The item’s location.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingContentFetch/side
	Side() NSFileProviderTestingOperationSide
}

// NSFileProviderTestingContentFetchObject wraps an existing Objective-C object that conforms to the NSFileProviderTestingContentFetch protocol.
type NSFileProviderTestingContentFetchObject struct {
	objectivec.Object
}

func (o NSFileProviderTestingContentFetchObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderTestingContentFetchObjectFromID constructs a [NSFileProviderTestingContentFetchObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderTestingContentFetchObjectFromID(id objc.ID) NSFileProviderTestingContentFetchObject {
	return NSFileProviderTestingContentFetchObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The operation’s type.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperation/type
func (o NSFileProviderTestingContentFetchObject) Type() NSFileProviderTestingOperationType {
	rv := objc.Send[NSFileProviderTestingOperationType](o.ID, objc.Sel("type"))
	return rv
}

// The containing item’s unique identifier.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingContentFetch/itemIdentifier
func (o NSFileProviderTestingContentFetchObject) ItemIdentifier() NSFileProviderItemIdentifier {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("itemIdentifier"))
	return NSFileProviderItemIdentifier(foundation.NSStringFromID(rv).String())
}

// The item’s location.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingContentFetch/side
func (o NSFileProviderTestingContentFetchObject) Side() NSFileProviderTestingOperationSide {
	rv := objc.Send[NSFileProviderTestingOperationSide](o.ID, objc.Sel("side"))
	return NSFileProviderTestingOperationSide(rv)
}
