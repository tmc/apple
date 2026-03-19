// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// A protocol that objects adopt to support fast enumeration.
//
// See: https://developer.apple.com/documentation/Foundation/NSFastEnumeration
type NSFastEnumeration interface {
	objectivec.IObject

	// Returns by reference a C array of objects over which the sender should iterate, and as the return value the number of objects in the array.
	//
	// See: https://developer.apple.com/documentation/Foundation/NSFastEnumeration/countByEnumerating(with:objects:count:)
	CountByEnumeratingWithStateObjectsCount(state NSFastEnumerationState, buffer []objectivec.IObject, len_ uint) uint
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

// Returns by reference a C array of objects over which the sender should
// iterate, and as the return value the number of objects in the array.
//
// state: Context information that is used in the enumeration to, in addition to
// other possibilities, ensure that the collection has not been mutated.
//
// buffer: A C array of objects over which the sender is to iterate.
//
// len: The maximum number of objects to return in `stackbuf`.
//
// # Return Value
// 
// The number of objects returned in `stackbuf`. Returns `0` when the
// iteration is finished.
//
// # Discussion
// 
// The state structure is assumed to be of stack local memory, so you can
// recast the passed in state structure to one more suitable for your
// iteration.
//
// See: https://developer.apple.com/documentation/Foundation/NSFastEnumeration/countByEnumerating(with:objects:count:)
func (o NSFastEnumerationObject) CountByEnumeratingWithStateObjectsCount(state NSFastEnumerationState, buffer []objectivec.IObject, len_ uint) uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("countByEnumeratingWithState:objects:count:"), state, objc.CArray(buffer), len_)
	return rv
	}

