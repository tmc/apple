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

// Class returns the underlying Objective-C class pointer.
func (nc NSEnumeratorClass) Class() objc.Class {
	return nc.class
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
// [NSSet.ObjectEnumerator] and [NSArray.ReverseObjectEnumerator].
// [NSDictionary] also has two methods that return an [NSEnumerator] object:
// [NSDictionary.KeyEnumerator] and [NSDictionary.ObjectEnumerator]. These
// methods let you enumerate the contents of a dictionary by key or by value,
// respectively.
//
// You send [NSEnumerator.NextObject] repeatedly to a newly created
// [NSEnumerator] object to have it return the next object in the original
// collection. When the collection is exhausted, `nil` is returned. You cannot
// “reset” an enumerator after it has exhausted its collection. To
// enumerate a collection again, you need a new enumerator.
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

// The array of unenumerated objects.
//
// # Discussion
//
// This array contains all the remaining objects of the enumerator in
// enumerated order. It does not contain objects that have already been
// enumerated with previous [NSEnumerator.NextObject] messages.
//
// Accessing this property exhausts the enumerator’s collection so that
// subsequent invocations of [NSEnumerator.NextObject] return `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSEnumerator/allObjects
func (e NSEnumerator) AllObjects() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("allObjects"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
