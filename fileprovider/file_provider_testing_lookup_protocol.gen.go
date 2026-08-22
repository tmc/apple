// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An operation that looks up an item.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingLookup
type NSFileProviderTestingLookup interface {
	objectivec.IObject
	NSFileProviderTestingOperation

	// The unique identifier for the item.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingLookup/itemIdentifier
	ItemIdentifier() NSFileProviderItemIdentifier

	// The location where the lookup occurs.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingLookup/side
	Side() NSFileProviderTestingOperationSide
}

// NSFileProviderTestingLookupObject wraps an existing Objective-C object that conforms to the NSFileProviderTestingLookup protocol.
type NSFileProviderTestingLookupObject struct {
	objectivec.Object
}

func (o NSFileProviderTestingLookupObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderTestingLookupObjectFromID constructs a [NSFileProviderTestingLookupObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderTestingLookupObjectFromID(id objc.ID) NSFileProviderTestingLookupObject {
	return NSFileProviderTestingLookupObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The operation’s type.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperation/type
func (o NSFileProviderTestingLookupObject) Type() NSFileProviderTestingOperationType {
	rv := objc.Send[NSFileProviderTestingOperationType](o.ID, objc.Sel("type"))
	return rv
}

// The unique identifier for the item.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingLookup/itemIdentifier
func (o NSFileProviderTestingLookupObject) ItemIdentifier() NSFileProviderItemIdentifier {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("itemIdentifier"))
	return NSFileProviderItemIdentifier(foundation.NSStringFromID(rv).String())
}

// The location where the lookup occurs.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingLookup/side
func (o NSFileProviderTestingLookupObject) Side() NSFileProviderTestingOperationSide {
	rv := objc.Send[NSFileProviderTestingOperationSide](o.ID, objc.Sel("side"))
	return NSFileProviderTestingOperationSide(rv)
}
