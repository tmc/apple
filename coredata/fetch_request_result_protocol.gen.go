// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// An abstract protocol used with parameterized fetch requests.
//
// See: https://developer.apple.com/documentation/CoreData/NSFetchRequestResult
type NSFetchRequestResult interface {
	objectivec.IObject
}

// NSFetchRequestResultObject wraps an existing Objective-C object that conforms to the NSFetchRequestResult protocol.
type NSFetchRequestResultObject struct {
	objectivec.Object
}

func (o NSFetchRequestResultObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFetchRequestResultObjectFromID constructs a [NSFetchRequestResultObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFetchRequestResultObjectFromID(id objc.ID) NSFetchRequestResultObject {
	return NSFetchRequestResultObject{
		Object: objectivec.ObjectFromID(id),
	}
}
