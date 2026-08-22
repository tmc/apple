// Code generated from Apple documentation for ScriptingBridge. DO NOT EDIT.

package scriptingbridge

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [SBElementArray] class.
var (
	_SBElementArrayClass     SBElementArrayClass
	_SBElementArrayClassOnce sync.Once
)

func getSBElementArrayClass() SBElementArrayClass {
	_SBElementArrayClassOnce.Do(func() {
		_SBElementArrayClass = SBElementArrayClass{class: objc.GetClass("SBElementArray")}
	})
	return _SBElementArrayClass
}

// GetSBElementArrayClass returns the class object for SBElementArray.
func GetSBElementArrayClass() SBElementArrayClass {
	return getSBElementArrayClass()
}

type SBElementArrayClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (sc SBElementArrayClass) Class() objc.Class {
	return sc.class
}

// Alloc allocates memory for a new instance of the class.
func (sc SBElementArrayClass) Alloc() SBElementArray {
	rv := objc.Send[SBElementArray](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// [SBElementArray] is subclass of [NSMutableArray] that manages collections
// of related [SBObject] objects. For example, when you ask the Finder for a
// list of disks, or ask iTunes for a list of playlists, you get the result
// back as an [SBElementArray] containing Scripting Bridge objects
// representing those items.
//
// # Overview
//
// [SBElementArray] defines methods beyond those of [NSArray] for obtaining
// individual objects. In addition to [object(at:)], [SBElementArray] also
// defines [SBElementArray.ObjectWithName], [SBElementArray.ObjectWithID], and
// [SBElementArray.ObjectAtLocation].
//
// # Subclassing Notes
//
// The [SBElementArray] class is not designed for subclassing.
//
// # Getting Objects in the Array
//
//   - [SBElementArray.ObjectWithName]: Returns the object in the array with the given name.
//   - [SBElementArray.ObjectWithID]: Returns the object in the array with the given identifier.
//   - [SBElementArray.ObjectAtLocation]: Returns the object at the given location in the receiver.
//
// # Getting the Referenced Array
//
//   - [SBElementArray.Get]: Forces evaluation of the receiver, causing the real object to be returned immediately.
//
// # Filtering an Element Array
//
//   - [SBElementArray.ArrayByApplyingSelector]: Returns an array containing the results of sending the specified message to each object in the receiver.
//   - [SBElementArray.ArrayByApplyingSelectorWithObject]: Returns an array containing the results of sending the specified message to each object in the receiver.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBElementArray
//
// [NSArray]: https://developer.apple.com/documentation/Foundation/NSArray
// [object(at:)]: https://developer.apple.com/documentation/Foundation/NSArray/object(at:)
type SBElementArray struct {
	foundation.NSMutableArray
}

// SBElementArrayFromID constructs a [SBElementArray] from an objc.ID.
//
// [SBElementArray] is subclass of [NSMutableArray] that manages collections
// of related [SBObject] objects. For example, when you ask the Finder for a
// list of disks, or ask iTunes for a list of playlists, you get the result
// back as an [SBElementArray] containing Scripting Bridge objects
// representing those items.
func SBElementArrayFromID(id objc.ID) SBElementArray {
	return SBElementArray{NSMutableArray: foundation.NSMutableArrayFromID(id)}
}

// NOTE: SBElementArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [SBElementArray] class.
//
// # Getting Objects in the Array
//
//   - [ISBElementArray.ObjectWithName]: Returns the object in the array with the given name.
//   - [ISBElementArray.ObjectWithID]: Returns the object in the array with the given identifier.
//   - [ISBElementArray.ObjectAtLocation]: Returns the object at the given location in the receiver.
//
// # Getting the Referenced Array
//
//   - [ISBElementArray.Get]: Forces evaluation of the receiver, causing the real object to be returned immediately.
//
// # Filtering an Element Array
//
//   - [ISBElementArray.ArrayByApplyingSelector]: Returns an array containing the results of sending the specified message to each object in the receiver.
//   - [ISBElementArray.ArrayByApplyingSelectorWithObject]: Returns an array containing the results of sending the specified message to each object in the receiver.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBElementArray
type ISBElementArray interface {
	foundation.INSMutableArray

	// Topic: Getting Objects in the Array

	// Returns the object in the array with the given name.
	ObjectWithName(name string) objectivec.IObject
	// Returns the object in the array with the given identifier.
	ObjectWithID(identifier objectivec.IObject) objectivec.IObject
	// Returns the object at the given location in the receiver.
	ObjectAtLocation(location objectivec.IObject) objectivec.IObject

	// Topic: Getting the Referenced Array

	// Forces evaluation of the receiver, causing the real object to be returned immediately.
	Get() []objectivec.IObject

	// Topic: Filtering an Element Array

	// Returns an array containing the results of sending the specified message to each object in the receiver.
	ArrayByApplyingSelector(selector objc.SEL) []objectivec.IObject
	// Returns an array containing the results of sending the specified message to each object in the receiver.
	ArrayByApplyingSelectorWithObject(aSelector objc.SEL, argument objectivec.IObject) []objectivec.IObject
}

// Init initializes the instance.
func (s SBElementArray) Init() SBElementArray {
	rv := objc.Send[SBElementArray](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s SBElementArray) Autorelease() SBElementArray {
	rv := objc.Send[SBElementArray](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewSBElementArray creates a new SBElementArray instance.
func NewSBElementArray() SBElementArray {
	class := getSBElementArrayClass()
	rv := objc.Send[SBElementArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the object in the array with the given name.
//
// name: The name of one of the receiver’s objects.
//
// # Return Value
//
// A reference to the designated object or `nil` if the object couldn’t be
// found.
//
// # Discussion
//
// This method is provided as an alternative to [object(at:)] for applications
// where a name is available instead of (or in addition to) an index. A name
// is generally more stable than an index. For example, it is typically more
// useful to identify a mailbox in Mail by its name than by its index in the
// list of mailboxes.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBElementArray/object(withName:)
//
// [object(at:)]: https://developer.apple.com/documentation/Foundation/NSArray/object(at:)
func (s SBElementArray) ObjectWithName(name string) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("objectWithName:"), objc.String(name))
	return objectivec.Object{ID: rv}
}

// Returns the object in the array with the given identifier.
//
// identifier: The identifier of one of the receiver’s objects.
//
// # Return Value
//
// A reference to the identified object or `nil` if could not be found.
//
// # Discussion
//
// This method is provided as an alternative to [object(at:)] for applications
// where an identifier is available instead of (or in addition to) an index. A
// unique ID is generally more stable than an index. For example, it may be
// more useful to identify a contact in Address Book by its identifier (which
// doesn’t change over time) than by its index in the list of contacts
// (which can change as contacts are added or removed).
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBElementArray/object(withID:)
//
// [object(at:)]: https://developer.apple.com/documentation/Foundation/NSArray/object(at:)
func (s SBElementArray) ObjectWithID(identifier objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("objectWithID:"), identifier)
	return objectivec.Object{ID: rv}
}

// Returns the object at the given location in the receiver.
//
// # Return Value
//
// A reference to the [SBObject] object identified by `loc` or `nil` if the
// object couldn’t be located.
//
// # Discussion
//
// This method is a generalization of [object(at:)] for applications where the
// “index” is not simply an integer. For example, Finder can specify
// objects using a [NSURL] object as a location. In OSA this is known as
// “absolute position,” a generalization of the notion of “index” in
// Foundation—it could be an integer, but it doesn’t have to be. A single
// object may even have a number of different “absolute position” values
// depending on the container.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBElementArray/object(atLocation:)
//
// [NSURL]: https://developer.apple.com/documentation/Foundation/NSURL
// [object(at:)]: https://developer.apple.com/documentation/Foundation/NSArray/object(at:)
func (s SBElementArray) ObjectAtLocation(location objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("objectAtLocation:"), location)
	return objectivec.Object{ID: rv}
}

// Forces evaluation of the receiver, causing the real object to be returned
// immediately.
//
// # Return Value
//
// The object referenced by the receiver.
//
// # Discussion
//
// This method forces the evaluation of the current object reference (the
// receiver), resulting in the return of the referenced object. By default,
// Scripting Bridge deals with references to objects until you actually
// request some concrete data from them or until you call the `get` method.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBElementArray/get()
func (s SBElementArray) Get() []objectivec.IObject {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("get"))
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns an array containing the results of sending the specified message to
// each object in the receiver.
//
// selector: A selector identifying the message to be sent to each object in the array.
//
// # Return Value
//
// A new array containing the results of sending the `selector` message to
// each object in the receiver, starting with the first object and continuing
// through the element array to the last object.
//
// # Discussion
//
// The method identified by `selector` must not take any arguments and must
// return an Objective-C object. It should not have the side effect of
// modifying the receiving array. The order of the items in the result array
// corresponds to the order of the items in the original array.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBElementArray/array(byApplying:)
func (s SBElementArray) ArrayByApplyingSelector(selector objc.SEL) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("arrayByApplyingSelector:"), selector)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}

// Returns an array containing the results of sending the specified message to
// each object in the receiver.
//
// argument: The value for the parameter of the message identified by `selector`.
//
// # Return Value
//
// A new array containing the results of sending the `selector` message to
// each object in the receiver, starting with the first object and continuing
// through the element array to the last object.
//
// # Discussion
//
// The method identified by `selector` must take a single argument—whose
// value is provided in `argument`—and must return an object. It should not
// have the side effect of modifying the receiving array. The order of the
// items in the result array corresponds to the order of the items in the
// original array.
//
// See: https://developer.apple.com/documentation/ScriptingBridge/SBElementArray/array(byApplying:with:)
func (s SBElementArray) ArrayByApplyingSelectorWithObject(aSelector objc.SEL, argument objectivec.IObject) []objectivec.IObject {
	rv := objc.Send[[]objc.ID](s.ID, objc.Sel("arrayByApplyingSelector:withObject:"), aSelector, argument)
	return objc.ConvertSlice(rv, func(id objc.ID) objectivec.IObject {
		return objectivec.Object{ID: id}
	})
}
