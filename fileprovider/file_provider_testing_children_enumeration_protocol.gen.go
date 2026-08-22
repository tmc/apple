// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An operation that lists a directory’s content.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingChildrenEnumeration
type NSFileProviderTestingChildrenEnumeration interface {
	objectivec.IObject
	NSFileProviderTestingOperation

	// The containing identifier’s unique identifier.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingChildrenEnumeration/itemIdentifier
	ItemIdentifier() NSFileProviderItemIdentifier

	// The item’s location.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingChildrenEnumeration/side
	Side() NSFileProviderTestingOperationSide
}

// NSFileProviderTestingChildrenEnumerationObject wraps an existing Objective-C object that conforms to the NSFileProviderTestingChildrenEnumeration protocol.
type NSFileProviderTestingChildrenEnumerationObject struct {
	objectivec.Object
}

func (o NSFileProviderTestingChildrenEnumerationObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderTestingChildrenEnumerationObjectFromID constructs a [NSFileProviderTestingChildrenEnumerationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderTestingChildrenEnumerationObjectFromID(id objc.ID) NSFileProviderTestingChildrenEnumerationObject {
	return NSFileProviderTestingChildrenEnumerationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The operation’s type.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperation/type
func (o NSFileProviderTestingChildrenEnumerationObject) Type() NSFileProviderTestingOperationType {
	rv := objc.Send[NSFileProviderTestingOperationType](o.ID, objc.Sel("type"))
	return rv
}

// The containing identifier’s unique identifier.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingChildrenEnumeration/itemIdentifier
func (o NSFileProviderTestingChildrenEnumerationObject) ItemIdentifier() NSFileProviderItemIdentifier {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("itemIdentifier"))
	return NSFileProviderItemIdentifier(foundation.NSStringFromID(rv).String())
}

// The item’s location.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingChildrenEnumeration/side
func (o NSFileProviderTestingChildrenEnumerationObject) Side() NSFileProviderTestingOperationSide {
	rv := objc.Send[NSFileProviderTestingOperationSide](o.ID, objc.Sel("side"))
	return NSFileProviderTestingOperationSide(rv)
}
