// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSEnumerator] class.
var (
	_NSEnumeratorClass     NSEnumeratorClass
	_NSEnumeratorClassOnce sync.Once
)

func getNSEnumeratorClass() NSEnumeratorClass {
	_NSEnumeratorClassOnce.Do(func() {
		_NSEnumeratorClass = NSEnumeratorClass{class: objc.GetClass("NSEnumerator")}
	})
	return _NSEnumeratorClass
}

// GetNSEnumeratorClass returns the class object for NSEnumerator.
func GetNSEnumeratorClass() NSEnumeratorClass {
	return getNSEnumeratorClass()
}

type NSEnumeratorClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSEnumeratorClass) Alloc() NSEnumerator {
	rv := objc.Send[NSEnumerator](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An abstract class whose subclasses enumerate collections of objects, such
// as arrays and dictionaries.
//
// # Overview
// 
// All creation methods are defined in the collection classes—such as
// [NSArray], [NSSet], and [NSDictionary]—which provide special
// [NSEnumerator] objects with which to enumerate their contents. For example,
// [NSArray] has two methods that return an [NSEnumerator] object:
// [ObjectEnumerator] and [ReverseObjectEnumerator]. [NSDictionary] also has
// two methods that return an [NSEnumerator] object: [KeyEnumerator] and
// [ObjectEnumerator]. These methods let you enumerate the contents of a
// dictionary by key or by value, respectively.
// 
// You send [NSEnumerator.NextObject] repeatedly to a newly created [NSEnumerator] object
// to have it return the next object in the original collection. When the
// collection is exhausted, `nil` is returned. You cannot “reset” an
// enumerator after it has exhausted its collection. To enumerate a collection
// again, you need a new enumerator.
// 
// The enumerator subclasses used by [NSArray], [NSDictionary], and [NSSet]
// retain the collection during enumeration. When the enumeration is
// exhausted, the collection is released.
//
// # Getting the Enumerated Objects
//
//   - [NSEnumerator.AllObjects]: The array of unenumerated objects.
//   - [NSEnumerator.NextObject]: Returns the next object from the collection being enumerated.
//
// See: https://developer.apple.com/documentation/Foundation/NSEnumerator
type NSEnumerator struct {
	objectivec.Object
}

// NSEnumeratorFromID constructs a [NSEnumerator] from an objc.ID.
//
// An abstract class whose subclasses enumerate collections of objects, such
// as arrays and dictionaries.
func NSEnumeratorFromID(id objc.ID) NSEnumerator {
	return NSEnumerator{objectivec.Object{ID: id}}
}
// NOTE: NSEnumerator adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSEnumerator] class.
//
// # Getting the Enumerated Objects
//
//   - [INSEnumerator.AllObjects]: The array of unenumerated objects.
//   - [INSEnumerator.NextObject]: Returns the next object from the collection being enumerated.
//
// See: https://developer.apple.com/documentation/Foundation/NSEnumerator
type INSEnumerator interface {
	objectivec.IObject

	// Topic: Getting the Enumerated Objects

	// The array of unenumerated objects.
	AllObjects() []objectivec.IObject
	// Returns the next object from the collection being enumerated.
	NextObject() objectivec.IObject
}

// Init initializes the instance.
func (e NSEnumerator) Init() NSEnumerator {
	rv := objc.Send[NSEnumerator](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e NSEnumerator) Autorelease() NSEnumerator {
	rv := objc.Send[NSEnumerator](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSEnumerator creates a new NSEnumerator instance.
func NewNSEnumerator() NSEnumerator {
	class := getNSEnumeratorClass()
	rv := objc.Send[NSEnumerator](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the next object from the collection being enumerated.
//
// # Return Value
// 
// The next object from the collection being enumerated, or `nil` when all
// objects have been enumerated.
//
// # Discussion
// 
// The following code illustrates how this method works using an array:
//
// See: https://developer.apple.com/documentation/Foundation/NSEnumerator/nextObject()
func (e NSEnumerator) NextObject() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("nextObject"))
	return objectivec.Object{ID: rv}
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
func (e NSEnumerator) CountByEnumeratingWithStateObjectsCount(state NSFastEnumerationState, buffer []objectivec.IObject, len_ uint) uint {
	rv := objc.Send[uint](e.ID, objc.Sel("countByEnumeratingWithState:objects:count:"), state, objc.CArray(buffer), len_)
	return rv
}

// The array of unenumerated objects.
//
// # Discussion
// 
// This array contains all the remaining objects of the enumerator in
// enumerated order. It does not contain objects that have already been
// enumerated with previous [NextObject] messages.
// 
// Accessing this property exhausts the enumerator’s collection so that
// subsequent invocations of [NextObject] return `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSEnumerator/allObjects
func (e NSEnumerator) AllObjects() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("allObjects"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

