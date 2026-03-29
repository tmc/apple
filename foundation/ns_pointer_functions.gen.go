// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSPointerFunctions] class.
var (
	_NSPointerFunctionsClass     NSPointerFunctionsClass
	_NSPointerFunctionsClassOnce sync.Once
)

func getNSPointerFunctionsClass() NSPointerFunctionsClass {
	_NSPointerFunctionsClassOnce.Do(func() {
		_NSPointerFunctionsClass = NSPointerFunctionsClass{class: objc.GetClass("NSPointerFunctions")}
	})
	return _NSPointerFunctionsClass
}

// GetNSPointerFunctionsClass returns the class object for NSPointerFunctions.
func GetNSPointerFunctionsClass() NSPointerFunctionsClass {
	return getNSPointerFunctionsClass()
}

type NSPointerFunctionsClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSPointerFunctionsClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSPointerFunctionsClass) Alloc() NSPointerFunctions {
	rv := objc.Send[NSPointerFunctions](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An instance of [NSPointerFunctions] defines callout functions appropriate
// for managing a pointer reference held somewhere else.
//
// # Overview
// 
// The functions specified by an instance of [NSPointerFunctions] are
// separated into two clusters—those that define “personality” such as
// “object” or “C-string”, and those that describe memory management
// issues such as a memory deallocation function. There are constants for
// common personalities and memory manager selections (see `Memory and
// Personality Options`).
// 
// [NSHashTable], [NSMapTable], and [NSPointerArray] use an
// [NSPointerFunctions] object to define the acquisition and retention
// behavior for the pointers they manage. Note, however, that not all
// combinations of personality and memory management behavior are valid for
// these collections. The pointer collection objects copy the
// [NSPointerFunctions] object on input and output, so you cannot usefully
// subclass [NSPointerFunctions].
// 
// # Subclassing Notes
// 
// [NSPointerFunctions] is not suitable for subclassing.
//
// # Creating and Initializing an NSPointerFunctions Object
//
//   - [NSPointerFunctions.InitWithOptions]: Returns an [NSPointerFunctions] object initialized with the given options.
//
// # Personality Functions
//
//   - [NSPointerFunctions.HashFunction]: The hash function.
//   - [NSPointerFunctions.SetHashFunction]
//   - [NSPointerFunctions.IsEqualFunction]: The function used to compare pointers.
//   - [NSPointerFunctions.SetIsEqualFunction]
//   - [NSPointerFunctions.SizeFunction]: The function used to determine the size of pointers.
//   - [NSPointerFunctions.SetSizeFunction]
//   - [NSPointerFunctions.DescriptionFunction]: The function used to describe elements.
//   - [NSPointerFunctions.SetDescriptionFunction]
//
// # Memory Configuration
//
//   - [NSPointerFunctions.AcquireFunction]: The function used to acquire memory.
//   - [NSPointerFunctions.SetAcquireFunction]
//   - [NSPointerFunctions.RelinquishFunction]: The function used to relinquish memory.
//   - [NSPointerFunctions.SetRelinquishFunction]
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerFunctions
type NSPointerFunctions struct {
	objectivec.Object
}

// NSPointerFunctionsFromID constructs a [NSPointerFunctions] from an objc.ID.
//
// An instance of [NSPointerFunctions] defines callout functions appropriate
// for managing a pointer reference held somewhere else.
func NSPointerFunctionsFromID(id objc.ID) NSPointerFunctions {
	return NSPointerFunctions{objectivec.Object{ID: id}}
}
// NOTE: NSPointerFunctions adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSPointerFunctions] class.
//
// # Creating and Initializing an NSPointerFunctions Object
//
//   - [INSPointerFunctions.InitWithOptions]: Returns an [NSPointerFunctions] object initialized with the given options.
//
// # Personality Functions
//
//   - [INSPointerFunctions.HashFunction]: The hash function.
//   - [INSPointerFunctions.SetHashFunction]
//   - [INSPointerFunctions.IsEqualFunction]: The function used to compare pointers.
//   - [INSPointerFunctions.SetIsEqualFunction]
//   - [INSPointerFunctions.SizeFunction]: The function used to determine the size of pointers.
//   - [INSPointerFunctions.SetSizeFunction]
//   - [INSPointerFunctions.DescriptionFunction]: The function used to describe elements.
//   - [INSPointerFunctions.SetDescriptionFunction]
//
// # Memory Configuration
//
//   - [INSPointerFunctions.AcquireFunction]: The function used to acquire memory.
//   - [INSPointerFunctions.SetAcquireFunction]
//   - [INSPointerFunctions.RelinquishFunction]: The function used to relinquish memory.
//   - [INSPointerFunctions.SetRelinquishFunction]
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerFunctions
type INSPointerFunctions interface {
	objectivec.IObject
	NSCopying

	// Topic: Creating and Initializing an NSPointerFunctions Object

	// Returns an [NSPointerFunctions] object initialized with the given options.
	InitWithOptions(options NSPointerFunctionsOptions) NSPointerFunctions

	// Topic: Personality Functions

	// The hash function.
	HashFunction() int
	SetHashFunction(value int)
	// The function used to compare pointers.
	IsEqualFunction() objectivec.IObject
	SetIsEqualFunction(value objectivec.IObject)
	// The function used to determine the size of pointers.
	SizeFunction() int
	SetSizeFunction(value int)
	// The function used to describe elements.
	DescriptionFunction() string
	SetDescriptionFunction(value string)

	// Topic: Memory Configuration

	// The function used to acquire memory.
	AcquireFunction() unsafe.Pointer
	SetAcquireFunction(value unsafe.Pointer)
	// The function used to relinquish memory.
	RelinquishFunction() objectivec.IObject
	SetRelinquishFunction(value objectivec.IObject)

	// The pointer functions for the hash table.
	PointerFunctions() INSPointerFunctions
	SetPointerFunctions(value INSPointerFunctions)
}

// Init initializes the instance.
func (p NSPointerFunctions) Init() NSPointerFunctions {
	rv := objc.Send[NSPointerFunctions](p.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p NSPointerFunctions) Autorelease() NSPointerFunctions {
	rv := objc.Send[NSPointerFunctions](p.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSPointerFunctions creates a new NSPointerFunctions instance.
func NewNSPointerFunctions() NSPointerFunctions {
	class := getNSPointerFunctionsClass()
	rv := objc.Send[NSPointerFunctions](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an [NSPointerFunctions] object initialized with the given options.
//
// options: The options for the new [NSPointerFunctions] object.
//
// # Return Value
// 
// The [NSPointerFunctions] object, initialized with the given options.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/init(options:)
func NewPointerFunctionsWithOptions(options NSPointerFunctionsOptions) NSPointerFunctions {
	instance := getNSPointerFunctionsClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithOptions:"), options)
	return NSPointerFunctionsFromID(rv)
}

// Returns an [NSPointerFunctions] object initialized with the given options.
//
// options: The options for the new [NSPointerFunctions] object.
//
// # Return Value
// 
// The [NSPointerFunctions] object, initialized with the given options.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/init(options:)
func (p NSPointerFunctions) InitWithOptions(options NSPointerFunctionsOptions) NSPointerFunctions {
	rv := objc.Send[NSPointerFunctions](p.ID, objc.Sel("initWithOptions:"), options)
	return rv
}

// Returns a new [NSPointerFunctions] object initialized with the given
// options.
//
// options: The options for the new [NSPointerFunctions] object.
//
// # Return Value
// 
// A new [NSPointerFunctions] object initialized with the given options.
//
// See: https://developer.apple.com/documentation/Foundation/NSPointerFunctions/pointerFunctionsWithOptions:
func (_NSPointerFunctionsClass NSPointerFunctionsClass) PointerFunctionsWithOptions(options NSPointerFunctionsOptions) NSPointerFunctions {
	rv := objc.Send[objc.ID](objc.ID(_NSPointerFunctionsClass.class), objc.Sel("pointerFunctionsWithOptions:"), options)
	return NSPointerFunctionsFromID(rv)
}

// The hash function.
//
// See: https://developer.apple.com/documentation/foundation/nspointerfunctions/hashfunction
func (p NSPointerFunctions) HashFunction() int {
	rv := objc.Send[int](p.ID, objc.Sel("hashFunction"))
	return rv
}
func (p NSPointerFunctions) SetHashFunction(value int) {
	objc.Send[struct{}](p.ID, objc.Sel("setHashFunction:"), value)
}
// The function used to compare pointers.
//
// See: https://developer.apple.com/documentation/foundation/nspointerfunctions/isequalfunction
func (p NSPointerFunctions) IsEqualFunction() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("isEqualFunction"))
	return objectivec.Object{ID: rv}
}
func (p NSPointerFunctions) SetIsEqualFunction(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setIsEqualFunction:"), value)
}
// The function used to determine the size of pointers.
//
// See: https://developer.apple.com/documentation/foundation/nspointerfunctions/sizefunction
func (p NSPointerFunctions) SizeFunction() int {
	rv := objc.Send[int](p.ID, objc.Sel("sizeFunction"))
	return rv
}
func (p NSPointerFunctions) SetSizeFunction(value int) {
	objc.Send[struct{}](p.ID, objc.Sel("setSizeFunction:"), value)
}
// The function used to describe elements.
//
// See: https://developer.apple.com/documentation/foundation/nspointerfunctions/descriptionfunction
func (p NSPointerFunctions) DescriptionFunction() string {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("descriptionFunction"))
	return NSStringFromID(rv).String()
}
func (p NSPointerFunctions) SetDescriptionFunction(value string) {
	objc.Send[struct{}](p.ID, objc.Sel("setDescriptionFunction:"), objc.String(value))
}
// The function used to acquire memory.
//
// See: https://developer.apple.com/documentation/foundation/nspointerfunctions/acquirefunction
func (p NSPointerFunctions) AcquireFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p.ID, objc.Sel("acquireFunction"))
	return rv
}
func (p NSPointerFunctions) SetAcquireFunction(value unsafe.Pointer) {
	objc.Send[struct{}](p.ID, objc.Sel("setAcquireFunction:"), value)
}
// The function used to relinquish memory.
//
// See: https://developer.apple.com/documentation/foundation/nspointerfunctions/relinquishfunction
func (p NSPointerFunctions) RelinquishFunction() objectivec.IObject {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("relinquishFunction"))
	return objectivec.Object{ID: rv}
}
func (p NSPointerFunctions) SetRelinquishFunction(value objectivec.IObject) {
	objc.Send[struct{}](p.ID, objc.Sel("setRelinquishFunction:"), value)
}
// The pointer functions for the hash table.
//
// See: https://developer.apple.com/documentation/foundation/nshashtable/pointerfunctions
func (p NSPointerFunctions) PointerFunctions() INSPointerFunctions {
	rv := objc.Send[objc.ID](p.ID, objc.Sel("pointerFunctions"))
	return NSPointerFunctionsFromID(objc.ID(rv))
}
func (p NSPointerFunctions) SetPointerFunctions(value INSPointerFunctions) {
	objc.Send[struct{}](p.ID, objc.Sel("setPointerFunctions:"), value)
}

			// Protocol methods for NSCopying
			

