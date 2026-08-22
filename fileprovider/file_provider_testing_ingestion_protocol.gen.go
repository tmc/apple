// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An operation that alerts the system to either local or remote storage changes.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingIngestion
type NSFileProviderTestingIngestion interface {
	objectivec.IObject
	NSFileProviderTestingOperation

	// The unique identifier for the item that changed.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingIngestion/itemIdentifier
	ItemIdentifier() NSFileProviderItemIdentifier

	// A description of the item that changed.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingIngestion/item
	Item() NSFileProviderItem

	// The location where the change occurred.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingIngestion/side
	Side() NSFileProviderTestingOperationSide
}

// NSFileProviderTestingIngestionObject wraps an existing Objective-C object that conforms to the NSFileProviderTestingIngestion protocol.
type NSFileProviderTestingIngestionObject struct {
	objectivec.Object
}

func (o NSFileProviderTestingIngestionObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderTestingIngestionObjectFromID constructs a [NSFileProviderTestingIngestionObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderTestingIngestionObjectFromID(id objc.ID) NSFileProviderTestingIngestionObject {
	return NSFileProviderTestingIngestionObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The operation’s type.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperation/type
func (o NSFileProviderTestingIngestionObject) Type() NSFileProviderTestingOperationType {
	rv := objc.Send[NSFileProviderTestingOperationType](o.ID, objc.Sel("type"))
	return rv
}

// The unique identifier for the item that changed.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingIngestion/itemIdentifier
func (o NSFileProviderTestingIngestionObject) ItemIdentifier() NSFileProviderItemIdentifier {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("itemIdentifier"))
	return NSFileProviderItemIdentifier(foundation.NSStringFromID(rv).String())
}

// A description of the item that changed.
//
// # Discussion
//
// This property is `nil` for deletion events.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingIngestion/item
func (o NSFileProviderTestingIngestionObject) Item() NSFileProviderItem {
	rv := objc.Send[objc.ID](o.ID, objc.Sel("item"))
	return NSFileProviderItemObjectFromID(rv)
}

// The location where the change occurred.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingIngestion/side
func (o NSFileProviderTestingIngestionObject) Side() NSFileProviderTestingOperationSide {
	rv := objc.Send[NSFileProviderTestingOperationSide](o.ID, objc.Sel("side"))
	return NSFileProviderTestingOperationSide(rv)
}
