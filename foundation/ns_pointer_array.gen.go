// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPointerArray] class.
var (
	_NSPointerArrayClass     NSPointerArrayClass
	_NSPointerArrayClassOnce sync.Once
)

func getNSPointerArrayClass() NSPointerArrayClass {
	_NSPointerArrayClassOnce.Do(func() {
		_NSPointerArrayClass = NSPointerArrayClass{class: objc.GetClass("NSPointerArray")}
	})
	return _NSPointerArrayClass
}

// GetNSPointerArrayClass returns the class object for NSPointerArray.
func GetNSPointerArrayClass() NSPointerArrayClass {
	return getNSPointerArrayClass()
}

type NSPointerArrayClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPointerArrayClass) Alloc() NSPointerArray {
	rv := objc.Send[NSPointerArray](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A collection similar to an array, but with a broader range of available
// memory semantics.
//
// # Overview
// 
// The pointer array class is modeled after [NSArray], but can also hold `nil`
// values. You can insert or remove `nil` values which contribute to the
// array’s [NSPointerArray.Count].
// 
// A pointer array can be initialized to maintain strong or weak references to
// objects, or according to any of the memory or personality options defined
// by [NSPointerFunctions.Options].
// 
// The [NSCopying] and [NSCoding] protocols are applicable only when a pointer
// array is initialized to maintain strong or weak references to objects.
// 
// When enumerating a pointer array with [NSFastEnumeration] using
// `for..XCUIElementTypeIn`, the loop will yield any `nil` values present in
// the array. See [Fast Enumeration Makes It Easy to Enumerate a Collection]
// in [Programming with Objective-C] for more information.
// 
// # Subclassing Notes
// 
// [NSPointerArray] is not suitable for subclassing.
//
// [Fast Enumeration Makes It Easy to Enumerate a Collection]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ProgrammingWithObjectiveC/FoundationTypesandCollections/FoundationTypesandCollections.html#//apple_ref/doc/uid/TP40011210-CH7-SW30
// [NSPointerFunctions.Options]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/Options
// [Programming with Objective-C]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ProgrammingWithObjectiveC/Introduction/Introduction.html#//apple_ref/doc/uid/TP40011210
//
// # Creating and Initializing a New Pointer Array
//
//   - [NSPointerArray.InitWithOptions]: Initializes the receiver to use the given options.
//   - [NSPointerArray.InitWithPointerFunctions]: Initializes the receiver to use the given functions.
//
// # Managing the Collection
//
//   - [NSPointerArray.Count]: The number of elements in the receiver.
//   - [NSPointerArray.SetCount]
//   - [NSPointerArray.AllObjects]: All the objects in the receiver.
//   - [NSPointerArray.PointerAtIndex]: Returns the pointer at a given index.
//   - [NSPointerArray.AddPointer]: Adds a given pointer to the receiver.
//   - [NSPointerArray.RemovePointerAtIndex]: Removes the pointer at a given index.
//   - [NSPointerArray.InsertPointerAtIndex]: Inserts a pointer at a given index.
//   - [NSPointerArray.ReplacePointerAtIndexWithPointer]: Replaces the pointer at a given index.
//   - [NSPointerArray.Compact]: Removes [NULL] values from the receiver.
//
// # Getting the Pointer Functions
//
//   - [NSPointerArray.PointerFunctions]: The functions in use by the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerArray
type NSPointerArray struct {
	objectivec.Object
}

// NSPointerArrayFromID constructs a [NSPointerArray] from an objc.ID.
//
// A collection similar to an array, but with a broader range of available
// memory semantics.
func NSPointerArrayFromID(id objc.ID) NSPointerArray {
	return NSPointerArray{objectivec.Object{id}}
}
// NOTE: NSPointerArray adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSPointerArray] class.
//
// # Creating and Initializing a New Pointer Array
//
//   - [INSPointerArray.InitWithOptions]: Initializes the receiver to use the given options.
//   - [INSPointerArray.InitWithPointerFunctions]: Initializes the receiver to use the given functions.
//
// # Managing the Collection
//
//   - [INSPointerArray.Count]: The number of elements in the receiver.
//   - [INSPointerArray.SetCount]
//   - [INSPointerArray.AllObjects]: All the objects in the receiver.
//   - [INSPointerArray.PointerAtIndex]: Returns the pointer at a given index.
//   - [INSPointerArray.AddPointer]: Adds a given pointer to the receiver.
//   - [INSPointerArray.RemovePointerAtIndex]: Removes the pointer at a given index.
//   - [INSPointerArray.InsertPointerAtIndex]: Inserts a pointer at a given index.
//   - [INSPointerArray.ReplacePointerAtIndexWithPointer]: Replaces the pointer at a given index.
//   - [INSPointerArray.Compact]: Removes [NULL] values from the receiver.
//
// # Getting the Pointer Functions
//
//   - [INSPointerArray.PointerFunctions]: The functions in use by the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerArray
type INSPointerArray interface {
	objectivec.IObject
	NSCopying

	// Topic: Creating and Initializing a New Pointer Array

	// Initializes the receiver to use the given options.
	InitWithOptions(options NSPointerFunctionsOptions) NSPointerArray
	// Initializes the receiver to use the given functions.
	InitWithPointerFunctions(functions INSPointerFunctions) NSPointerArray

	// Topic: Managing the Collection

	// The number of elements in the receiver.
	Count() uint
	SetCount(value uint)
	// All the objects in the receiver.
	AllObjects() INSArray
	// Returns the pointer at a given index.
	PointerAtIndex(index uint)
	// Adds a given pointer to the receiver.
	AddPointer(pointer unsafe.Pointer)
	// Removes the pointer at a given index.
	RemovePointerAtIndex(index uint)
	// Inserts a pointer at a given index.
	InsertPointerAtIndex(item unsafe.Pointer, index uint)
	// Replaces the pointer at a given index.
	ReplacePointerAtIndexWithPointer(index uint, item unsafe.Pointer)
	// Removes [NULL] values from the receiver.
	Compact()

	// Topic: Getting the Pointer Functions

	// The functions in use by the receiver.
	PointerFunctions() INSPointerFunctions

	// Encodes the receiver using a given archiver.
	EncodeWithCoder(coder INSCoder)
	InitWithCoder(coder INSCoder) NSPointerArray
}





// Init initializes the instance.
func (p NSPointerArray) Init() NSPointerArray {
	rv := objc.Send[NSPointerArray](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPointerArray) Autorelease() NSPointerArray {
	rv := objc.Send[NSPointerArray](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPointerArray creates a new NSPointerArray instance.
func NewNSPointerArray() NSPointerArray {
	class := getNSPointerArrayClass()
	rv := objc.Send[NSPointerArray](objc.ID(class.class), objc.Sel("new"))
	return rv
}






//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewPointerArrayWithCoder(coder INSCoder) NSPointerArray {
	instance := getNSPointerArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSPointerArrayFromID(rv)
}


// Initializes the receiver to use the given options.
//
// options: The pointer functions options for the new instance.
//
// # Return Value
// 
// The receiver, initialized to use the given options.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerArray/init(options:)
func NewPointerArrayWithOptions(options NSPointerFunctionsOptions) NSPointerArray {
	instance := getNSPointerArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOptions:"), options)
	return NSPointerArrayFromID(rv)
}


// Initializes the receiver to use the given functions.
//
// functions: The pointer functions for the new instance.
//
// # Return Value
// 
// The receiver, initialized to use the given functions.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerArray/init(pointerFunctions:)
func NewPointerArrayWithPointerFunctions(functions INSPointerFunctions) NSPointerArray {
	instance := getNSPointerArrayClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithPointerFunctions:"), functions)
	return NSPointerArrayFromID(rv)
}







// Initializes the receiver to use the given options.
//
// options: The pointer functions options for the new instance.
//
// # Return Value
// 
// The receiver, initialized to use the given options.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerArray/init(options:)
func (p NSPointerArray) InitWithOptions(options NSPointerFunctionsOptions) NSPointerArray {
	rv := objc.Send[NSPointerArray](p.ID, objc.Sel("initWithOptions:"), options)
	return rv
}

// Initializes the receiver to use the given functions.
//
// functions: The pointer functions for the new instance.
//
// # Return Value
// 
// The receiver, initialized to use the given functions.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerArray/init(pointerFunctions:)
func (p NSPointerArray) InitWithPointerFunctions(functions INSPointerFunctions) NSPointerArray {
	rv := objc.Send[NSPointerArray](p.ID, objc.Sel("initWithPointerFunctions:"), functions)
	return rv
}

// Returns the pointer at a given index.
//
// index: The index of an element in the receiver. This value must be less than the
// [Count] of the receiver.
//
// # Return Value
// 
// The pointer at `index`.
//
// # Discussion
// 
// The returned value may be [NULL].
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerArray/pointer(at:)
func (p NSPointerArray) PointerAtIndex(index uint) {
	objc.Send[objc.ID](p.ID, objc.Sel("pointerAtIndex:"), index)
}

// Adds a given pointer to the receiver.
//
// pointer: The pointer to add. This value may be [NULL].
//
// # Discussion
// 
// `pointer` is added at index [Count].
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerArray/addPointer(_:)
func (p NSPointerArray) AddPointer(pointer unsafe.Pointer) {
	objc.Send[objc.ID](p.ID, objc.Sel("addPointer:"), pointer)
}

// Removes the pointer at a given index.
//
// index: The index of an element in the receiver. This value must be less than the
// [Count] of the receiver.
//
// # Discussion
// 
// Elements above `index`, including [NULL] values, slide lower.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerArray/removePointer(at:)
func (p NSPointerArray) RemovePointerAtIndex(index uint) {
	objc.Send[objc.ID](p.ID, objc.Sel("removePointerAtIndex:"), index)
}

// Inserts a pointer at a given index.
//
// item: The pointer to add.
//
// index: The index of an element in the receiver. This value must be less than the
// [Count] of the receiver.
//
// # Discussion
// 
// Elements at and above `index`, including [NULL] values, slide higher.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerArray/insertPointer(_:at:)
func (p NSPointerArray) InsertPointerAtIndex(item unsafe.Pointer, index uint) {
	objc.Send[objc.ID](p.ID, objc.Sel("insertPointer:atIndex:"), item, index)
}

// Replaces the pointer at a given index.
//
// index: The index of an element in the receiver. This value must be less than the
// [Count] of the receiver.
//
// item: The item with which to replace the element at `index`. This value may be
// [NULL].
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerArray/replacePointer(at:withPointer:)
func (p NSPointerArray) ReplacePointerAtIndexWithPointer(index uint, item unsafe.Pointer) {
	objc.Send[objc.ID](p.ID, objc.Sel("replacePointerAtIndex:withPointer:"), index, item)
}

// Removes [NULL] values from the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerArray/compact()
func (p NSPointerArray) Compact() {
	objc.Send[objc.ID](p.ID, objc.Sel("compact"))
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
func (p NSPointerArray) CountByEnumeratingWithStateObjectsCount(state NSFastEnumerationState, buffer []objectivec.IObject, len_ uint) uint {
	rv := objc.Send[uint](p.ID, objc.Sel("countByEnumeratingWithState:objects:count:"), state, objc.CArray(buffer), len_)
	return rv
}

// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (p NSPointerArray) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](p.ID, objc.Sel("encodeWithCoder:"), coder)
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (p NSPointerArray) InitWithCoder(coder INSCoder) NSPointerArray {
	rv := objc.Send[NSPointerArray](p.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}





// Returns a new pointer array that maintains strong references to its
// elements.
//
// # Return Value
// 
// A new pointer array that maintains strong references to its elements.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerArray/strongObjects()
func (_NSPointerArrayClass NSPointerArrayClass) StrongObjectsPointerArray() NSPointerArray {
	rv := objc.Send[objc.ID](objc.ID(_NSPointerArrayClass.class), objc.Sel("strongObjectsPointerArray"))
	return NSPointerArrayFromID(rv)
}

// Returns a new pointer array that maintains weak references to its elements.
//
// # Return Value
// 
// A new pointer array that maintains weak references to its elements.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerArray/weakObjects()
func (_NSPointerArrayClass NSPointerArrayClass) WeakObjectsPointerArray() NSPointerArray {
	rv := objc.Send[objc.ID](objc.ID(_NSPointerArrayClass.class), objc.Sel("weakObjectsPointerArray"))
	return NSPointerArrayFromID(rv)
}

// Returns a new pointer array initialized to use the given options.
//
// options: The pointer functions options for the new instance.
//
// # Return Value
// 
// A new pointer array initialized to use the given options.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerArray/pointerArrayWithOptions:
func (_NSPointerArrayClass NSPointerArrayClass) PointerArrayWithOptions(options NSPointerFunctionsOptions) NSPointerArray {
	rv := objc.Send[objc.ID](objc.ID(_NSPointerArrayClass.class), objc.Sel("pointerArrayWithOptions:"), options)
	return NSPointerArrayFromID(rv)
}

// A new pointer array initialized to use the given functions.
//
// functions: The pointer functions for the new instance.
//
// # Return Value
// 
// A new pointer array initialized to use the given pointer functions.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerArray/pointerArrayWithPointerFunctions:
func (_NSPointerArrayClass NSPointerArrayClass) PointerArrayWithPointerFunctions(functions INSPointerFunctions) NSPointerArray {
	rv := objc.Send[objc.ID](objc.ID(_NSPointerArrayClass.class), objc.Sel("pointerArrayWithPointerFunctions:"), functions)
	return NSPointerArrayFromID(rv)
}








// The number of elements in the receiver.
//
// # Discussion
// 
// If you increase the `count`, [NULL] values are added. If you decrease the
// `count`, elements at indexes `count` and greater are removed.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerArray/count
func (p NSPointerArray) Count() uint {
	rv := objc.Send[uint](p.ID, objc.Sel("count"))
	return rv
}
func (p NSPointerArray) SetCount(value uint) {
	objc.Send[struct{}](p.ID, objc.Sel("setCount:"), value)
}



// All the objects in the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerArray/allObjects
func (p NSPointerArray) AllObjects() INSArray {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("allObjects"))
	return NSArrayFromID(objc.ID(rv))
}



// The functions in use by the receiver.
//
// # Discussion
// 
// The returned object is a new [NSPointerFunctions] object that you can
// modify and/or use directly to create other pointer collections.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerArray/pointerFunctions
func (p NSPointerArray) PointerFunctions() INSPointerFunctions {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("pointerFunctions"))
	return NSPointerFunctionsFromID(objc.ID(rv))
}














			// Protocol methods for NSCopying
			
















