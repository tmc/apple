// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An operation that the system can schedule.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperation
type NSFileProviderTestingOperation interface {
	objectivec.IObject

	// The operation’s type.
	//
	// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperation/type
	Type() NSFileProviderTestingOperationType
}

// NSFileProviderTestingOperationObject wraps an existing Objective-C object that conforms to the NSFileProviderTestingOperation protocol.
type NSFileProviderTestingOperationObject struct {
	objectivec.Object
}

func (o NSFileProviderTestingOperationObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFileProviderTestingOperationObjectFromID constructs a [NSFileProviderTestingOperationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFileProviderTestingOperationObjectFromID(id objc.ID) NSFileProviderTestingOperationObject {
	return NSFileProviderTestingOperationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

// The operation’s type.
//
// See: https://developer.apple.com/documentation/FileProvider/NSFileProviderTestingOperation/type
func (o NSFileProviderTestingOperationObject) Type() NSFileProviderTestingOperationType {
	rv := objc.Send[NSFileProviderTestingOperationType](o.ID, objc.Sel("type"))
	return NSFileProviderTestingOperationType(rv)
}
