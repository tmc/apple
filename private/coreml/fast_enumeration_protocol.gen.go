// Code generated from Apple documentation for CoreML. DO NOT EDIT.

package coreml

import (
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// NSFastEnumeration protocol.
type NSFastEnumeration interface {
	objectivec.IObject

	// CountByEnumeratingWithStateObjectsCount protocol.
	CountByEnumeratingWithStateObjectsCount(state unsafe.Pointer, objects []objectivec.IObject, count uint64) uint64
}

// NSFastEnumerationObject wraps an existing Objective-C object that conforms to the NSFastEnumeration protocol.
type NSFastEnumerationObject struct {
	objectivec.Object
}

func (o NSFastEnumerationObject) BaseObject() objectivec.Object {
	return o.Object
}

// NSFastEnumerationObjectFromID constructs a [NSFastEnumerationObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSFastEnumerationObjectFromID(id objc.ID) NSFastEnumerationObject {
	return NSFastEnumerationObject{
		Object: objectivec.ObjectFromID(id),
	}
}

func (o NSFastEnumerationObject) CountByEnumeratingWithStateObjectsCount(state unsafe.Pointer, objects []objectivec.IObject, count uint64) uint64 {
	rv := objc.SendIfResponds[uint64](o.ID, objc.Sel("countByEnumeratingWithState:objects:count:"), objc.CArray(state), objc.CArray(objects), count)
	return rv
}
