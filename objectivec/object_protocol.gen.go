// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import (
	"unsafe"
	"github.com/tmc/apple/objc"
)

// The group of methods that are fundamental to all Objective-C objects.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol
type NSObject interface {

	// Returns the class object for the receiver’s superclass.
	//
	// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/superclass
	Superclass() objc.Class

	// Returns an integer that can be used as a table address in a hash table structure.
	//
	// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/hash
	Hash() uint

	// Returns the receiver.
	//
	// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/self()
	Self() IObject

	// Returns a Boolean value that indicates whether the receiver is an instance of given class or an instance of any class that inherits from that class.
	//
	// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isKind(of:)
	IsKindOfClass(aClass objc.Class) bool

	// Returns a Boolean value that indicates whether the receiver is an instance of a given class.
	//
	// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isMember(of:)
	IsMemberOfClass(aClass objc.Class) bool

	// Returns a Boolean value that indicates whether the receiver implements or inherits a method that can respond to a specified message.
	//
	// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/responds(to:)
	RespondsToSelector(aSelector objc.SEL) bool

	// Returns a Boolean value that indicates whether the receiver conforms to a given protocol.
	//
	// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/conforms(to:)
	ConformsToProtocol(aProtocol unsafe.Pointer) bool

	// A textual representation of the receiver.
	//
	// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/description
	Description() IObject

	// A textual representation of the receiver to use with a debugger.
	//
	// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/debugDescription
	DebugDescription() IObject

	// Returns a Boolean value that indicates whether the receiver does not descend from [NSObject](<doc://com.apple.objectivec/documentation/ObjectiveC/NSObject-swift.class>).
	//
	// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isProxy()
	IsProxy() bool

	// Decrements the receiver’s reference count.
	//
	// See: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/release
	Release()

	// Increments the receiver’s reference count.
	//
	// See: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/retain
	Retain() IObject
}

// NSObjectObject wraps an existing Objective-C object that conforms to the NSObject protocol.
type NSObjectObject struct {
	Object
}
func (o NSObjectObject) BaseObject() Object {
	return o.Object
}

// NSObjectObjectFromID constructs a [NSObjectObject] from an objc.ID.
// The object is determined to conform to the protocol at runtime.
func NSObjectObjectFromID(id objc.ID) NSObjectObject {
	return NSObjectObject{
		Object: ObjectFromID(id),
	}
}

// Returns the class object for the receiver’s superclass.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/superclass

func (o NSObjectObject) Superclass() objc.Class {
	
	rv := objc.Send[objc.Class](o.ID, objc.Sel("superclass"))
	return rv
	}

// Returns a Boolean value that indicates whether the receiver and a given
// object are equal.
//
// object: The object to be compared to the receiver. May be `nil`, in which case this
// method returns [NO].
// //
// [NO]: https://developer.apple.com/documentation/ObjectiveC/NO
//
// # Return Value
// 
// [YES] if the receiver and `anObject` are equal, otherwise [NO].
//
// [NO]: https://developer.apple.com/documentation/ObjectiveC/NO
// [YES]: https://developer.apple.com/documentation/ObjectiveC/YES
//
// # Discussion
// 
// This method defines what it means for instances to be equal. For example, a
// container object might define two containers as equal if their
// corresponding objects all respond [YES] to an [IsEqual] request. See the
// [NSData], [NSDictionary], [NSArray], and [NSString] class specifications
// for examples of the use of this method.
// 
// If two objects are equal, they must have the same hash value. This last
// point is particularly important if you define [IsEqual] in a subclass and
// intend to put instances of that subclass into a collection. Make sure you
// also define [Hash] in your subclass.
//
// [NSArray]: https://developer.apple.com/documentation/Foundation/NSArray
// [NSData]: https://developer.apple.com/documentation/Foundation/NSData
// [NSDictionary]: https://developer.apple.com/documentation/Foundation/NSDictionary
// [NSString]: https://developer.apple.com/documentation/Foundation/NSString
// [YES]: https://developer.apple.com/documentation/ObjectiveC/YES
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isEqual(_:)

func (o NSObjectObject) IsEqual(object IObject) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isEqual:"), object)
	return rv
	}

// Returns an integer that can be used as a table address in a hash table
// structure.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/hash

func (o NSObjectObject) Hash() uint {
	
	rv := objc.Send[uint](o.ID, objc.Sel("hash"))
	return rv
	}

// Returns the receiver.
//
// # Return Value
// 
// The receiver.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/self()

func (o NSObjectObject) Self() IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("self"))
	return Object{ID: rv}
	}

// Returns a Boolean value that indicates whether the receiver is an instance
// of given class or an instance of any class that inherits from that class.
//
// aClass: A class object representing the Objective-C class to be tested.
//
// # Return Value
// 
// [YES] if the receiver is an instance of `aClass` or an instance of any
// class that inherits from `aClass`, otherwise [NO].
//
// [NO]: https://developer.apple.com/documentation/ObjectiveC/NO
// [YES]: https://developer.apple.com/documentation/ObjectiveC/YES
//
// # Discussion
// 
// For example, in this code, [IsKindOfClass] would return [YES] because, in
// Foundation, the [NSArchiver] class inherits from [NSCoder]:
// 
// Be careful when using this method on objects represented by a class
// cluster. Because of the nature of class clusters, the object you get back
// may not always be the type you expected. If you call a method that returns
// a class cluster, the exact type returned by the method is the best
// indicator of what you can do with that object. For example, if a method
// returns a pointer to an [NSArray] object, you should not use this method to
// see if the array is mutable, as shown in the following code:
// 
// If you use such constructs in your code, you might think it is alright to
// modify an object that in reality should not be modified. Doing so might
// then create problems for other code that expected the object to remain
// unchanged.
// 
// If the receiver is a class object, this method returns [YES] if `aClass` is
// a Class object of the same type, [NO] otherwise.
//
// [NO]: https://developer.apple.com/documentation/ObjectiveC/NO
// [NSArchiver]: https://developer.apple.com/documentation/Foundation/NSArchiver
// [NSArray]: https://developer.apple.com/documentation/Foundation/NSArray
// [NSCoder]: https://developer.apple.com/documentation/Foundation/NSCoder
// [YES]: https://developer.apple.com/documentation/ObjectiveC/YES
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isKind(of:)

func (o NSObjectObject) IsKindOfClass(aClass objc.Class) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isKindOfClass:"), aClass)
	return rv
	}

// Returns a Boolean value that indicates whether the receiver is an instance
// of a given class.
//
// aClass: A class object representing the Objective-C class to be tested.
//
// # Return Value
// 
// [YES] if the receiver is an instance of `aClass`, otherwise [NO].
//
// [NO]: https://developer.apple.com/documentation/ObjectiveC/NO
// [YES]: https://developer.apple.com/documentation/ObjectiveC/YES
//
// # Discussion
// 
// For example, in this code, [IsMemberOfClass] would return [NO]:
// 
// Class objects may be compiler-created objects but they still support the
// concept of membership. Thus, you can use this method to verify that the
// receiver is a specific Class object.
//
// [NO]: https://developer.apple.com/documentation/ObjectiveC/NO
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isMember(of:)

func (o NSObjectObject) IsMemberOfClass(aClass objc.Class) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isMemberOfClass:"), aClass)
	return rv
	}

// Returns a Boolean value that indicates whether the receiver implements or
// inherits a method that can respond to a specified message.
//
// aSelector: A selector that identifies a message.
//
// # Return Value
// 
// [YES] if the receiver implements or inherits a method that can respond to
// `aSelector`, otherwise [NO].
//
// [NO]: https://developer.apple.com/documentation/ObjectiveC/NO
// [YES]: https://developer.apple.com/documentation/ObjectiveC/YES
//
// # Discussion
// 
// The application is responsible for determining whether a [NO] response
// should be considered an error.
// 
// You cannot test whether an object inherits a method from its superclass by
// sending [RespondsToSelector] to the object using the `super` keyword. This
// method will still be testing the object as a whole, not just the
// superclass’s implementation. Therefore, sending [RespondsToSelector] to
// `super` is equivalent to sending it to `self`. Instead, you must invoke the
// [NSObject] class method [instancesRespond(to:)] directly on the object’s
// superclass, as illustrated in the following code fragment.
// 
// You cannot simply use `[[self superclass] @selector(aMethod)]` since this
// may cause the method to fail if it is invoked by a subclass.
// 
// Note that if the receiver is able to forward `aSelector` messages to
// another object, it will be able to respond to the message, albeit
// indirectly, even though this method returns [NO].
//
// [NO]: https://developer.apple.com/documentation/ObjectiveC/NO
// [instancesRespond(to:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/instancesRespond(to:)
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/responds(to:)

func (o NSObjectObject) RespondsToSelector(aSelector objc.SEL) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("respondsToSelector:"), aSelector)
	return rv
	}

// Returns a Boolean value that indicates whether the receiver conforms to a
// given protocol.
//
// aProtocol: A protocol object that represents a particular protocol.
//
// # Return Value
// 
// [YES] if the receiver conforms to `aProtocol`, otherwise [NO].
//
// [NO]: https://developer.apple.com/documentation/ObjectiveC/NO
// [YES]: https://developer.apple.com/documentation/ObjectiveC/YES
//
// # Discussion
// 
// This method works identically to the [conforms(to:)] class method declared
// in [NSObject]. It’s provided as a convenience so that you don’t need to
// get the class object to find out whether an instance can respond to a given
// set of messages.
//
// [NSObject]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class
// [conforms(to:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/conforms(to:)
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/conforms(to:)

func (o NSObjectObject) ConformsToProtocol(aProtocol unsafe.Pointer) bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("conformsToProtocol:"), aProtocol)
	return rv
	}

// A textual representation of the receiver.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/description

func (o NSObjectObject) Description() IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("description"))
	return Object{ID: rv}
	}

// A textual representation of the receiver to use with a debugger.
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/debugDescription

func (o NSObjectObject) DebugDescription() IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("debugDescription"))
	return Object{ID: rv}
	}

// Sends a specified message to the receiver and returns the result of the
// message.
//
// aSelector: A selector identifying the message to send. The message should take no
// arguments. If `aSelector` is [NULL], an [invalidArgumentException] is
// raised.
// //
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
//
// # Return Value
// 
// An object that is the result of the message.
//
// # Discussion
// 
// Calling the [PerformSelector] method is equivalent to sending the
// `aSelector` message directly to the receiver. For example, the following
// both do the same thing if `anObject` is an instance of [MyObject]:
// 
// The [PerformSelector] method allows you to send messages that aren’t
// determined until run-time. This means that you can pass a variable selector
// as the argument:
// 
// Use caution when doing this. This method returns an implicitly unwrapped
// optional unmanaged pointer to an [AnyObject] instance ([Unmanaged]`!`).
// It’s up to you to decide how to bring the instance into Swift’s memory
// management scheme. Different messages require different memory management
// strategies for their returned objects, and it might not be obvious which to
// use.
// 
// Usually, a caller isn’t responsible for the memory of a returned
// instance, in which case you use [takeUnretainedValue()], as shown above.
// However, for any of the creation methods, such as [copy()], the caller is
// responsible, and you use [takeRetainedValue()] instead. See [Memory
// Management Policy] in [Advanced Memory Management Programming Guide] for a
// description of ownership expectations.
// 
// Due to this uncertainty, the compiler generates a warning if you supply a
// variable selector while using ARC to manage memory. Because it can’t
// determine ownership of the returned object at compile-time, ARC makes the
// assumption that the caller does need to take ownership, but this may not be
// true. The compiler warning alerts you to the potential for a memory leak.
// 
// To avoid the warning, if you know that `aSelector` has no return value, you
// might be able to use [performSelector(onMainThread:with:waitUntilDone:)] or
// one of the related methods available in [NSObject].
// 
// For a more general solution, use [NSInvocation] to construct a message that
// you can invoke with an arbitrary argument list and return value.
// 
// Alternatively, consider restructuring your code to use blocks as a means of
// passing chunks of functionality through an API. See [Blocks Programming
// Topics] for details.
//
// [Advanced Memory Management Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/MemoryMgmt/Articles/MemoryMgmt.html#//apple_ref/doc/uid/10000011i
// [AnyObject]: https://developer.apple.com/documentation/Swift/AnyObject
// [Blocks Programming Topics]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Blocks/Articles/00_Introduction.html#//apple_ref/doc/uid/TP40007502
// [Memory Management Policy]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/MemoryMgmt/Articles/mmRules.html#//apple_ref/doc/uid/20000994
// [NSInvocation]: https://developer.apple.com/documentation/Foundation/NSInvocation
// [NSObject]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class
// [Unmanaged]: https://developer.apple.com/documentation/Swift/Unmanaged
// [copy()]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/copy()
// [performSelector(onMainThread:with:waitUntilDone:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/performSelector(onMainThread:with:waitUntilDone:)
// [takeRetainedValue()]: https://developer.apple.com/documentation/Swift/Unmanaged/takeRetainedValue()
// [takeUnretainedValue()]: https://developer.apple.com/documentation/Swift/Unmanaged/takeUnretainedValue()
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/perform(_:)

func (o NSObjectObject) PerformSelector(aSelector objc.SEL) IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("performSelector:"), aSelector)
	return Object{ID: rv}
	}

// Sends a message to the receiver with an object as the argument.
//
// aSelector: A selector identifying the message to send. If `aSelector` is [NULL], an
// [invalidArgumentException] is raised.
// //
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
//
// object: An object that is the sole argument of the message.
//
// # Return Value
// 
// An object that is the result of the message.
//
// # Discussion
// 
// This method is the same as [PerformSelector] except that you can supply an
// argument for `aSelector`. `aSelector` should identify a method that takes a
// single argument of type id. For methods with other argument types and
// return values, use [NSInvocation].
//
// [NSInvocation]: https://developer.apple.com/documentation/Foundation/NSInvocation
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/perform(_:with:)

func (o NSObjectObject) PerformSelectorWithObject(aSelector objc.SEL, object IObject) IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("performSelector:withObject:"), aSelector, object)
	return Object{ID: rv}
	}

// Sends a message to the receiver with two objects as arguments.
//
// aSelector: A selector identifying the message to send. If `aSelector` is [NULL], an
// [invalidArgumentException] is raised.
// //
// [invalidArgumentException]: https://developer.apple.com/documentation/Foundation/NSExceptionName/invalidArgumentException
//
// object1: An object that is the first argument of the message.
//
// object2: An object that is the second argument of the message
//
// # Return Value
// 
// An object that is the result of the message.
//
// # Discussion
// 
// This method is the same as [PerformSelector] except that you can supply two
// arguments for `aSelector`. `aSelector` should identify a method that can
// take two arguments of type id. For methods with other argument types and
// return values, use [NSInvocation].
//
// [NSInvocation]: https://developer.apple.com/documentation/Foundation/NSInvocation
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/perform(_:with:with:)

func (o NSObjectObject) PerformSelectorWithObjectWithObject(aSelector objc.SEL, object1 IObject, object2 IObject) IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("performSelector:withObject:withObject:"), aSelector, object1, object2)
	return Object{ID: rv}
	}

// Returns a Boolean value that indicates whether the receiver does not
// descend from [NSObject].
//
// [NSObject]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class
//
// # Return Value
// 
// [NO] if the receiver really descends from [NSObject], otherwise [YES].
//
// [NO]: https://developer.apple.com/documentation/ObjectiveC/NO
// [NSObject]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class
// [YES]: https://developer.apple.com/documentation/ObjectiveC/YES
//
// # Discussion
// 
// This method is necessary because sending [IsKindOfClass] or
// [IsMemberOfClass] to an [NSProxy] object will test the object the proxy
// stands in for, not the proxy itself. Use this method to test if the
// receiver is a proxy (or a member of some other root class).
//
// [NSProxy]: https://developer.apple.com/documentation/Foundation/NSProxy
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSObjectProtocol/isProxy()

func (o NSObjectObject) IsProxy() bool {
	
	rv := objc.Send[bool](o.ID, objc.Sel("isProxy"))
	return rv
	}

// Decrements the receiver’s reference count.
//
// # Discussion
// 
// The receiver is sent a [dealloc] message when its reference count reaches
// `0`.
// 
// You would only implement this method to define your own reference-counting
// scheme. Such implementations should not invoke the inherited method; that
// is, they should not include a release message to `super`.
// 
// For more information on the reference counting mechanism, see [Advanced
// Memory Management Programming Guide].
// 
// # Special Considerations
// 
// Instead of using manual reference counting, you should adopt ARC—see
// [Transitioning to ARC Release Notes].
//
// [Advanced Memory Management Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/MemoryMgmt/Articles/MemoryMgmt.html#//apple_ref/doc/uid/10000011i
// [Transitioning to ARC Release Notes]: https://developer.apple.com/library/archive/releasenotes/ObjectiveC/RN-TransitioningToARC/Introduction/Introduction.html#//apple_ref/doc/uid/TP40011226
// [dealloc]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/dealloc
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/release

func (o NSObjectObject) Release() {
	
	objc.Send[struct{}](o.ID, objc.Sel("release"))
	}

// Increments the receiver’s reference count.
//
// # Return Value
// 
// `self`.
//
// # Discussion
// 
// You send an object a [Retain] message when you want to prevent it from
// being deallocated until you have finished using it.
// 
// An object is deallocated automatically when its reference count reaches
// `0`. [Retain] messages increment the reference count, and [Release]
// messages decrement it. For more information on this mechanism, see
// [Advanced Memory Management Programming Guide].
// 
// As a convenience, [Retain] returns `self` because it may be used in nested
// expressions.
// 
// You would implement this method only if you were defining your own
// reference-counting scheme. Such implementations must return `self` and
// should not invoke the inherited method by sending a [Retain] message to
// `super`.
// 
// # Special Considerations
// 
// Instead of using manual reference counting, you should adopt ARC—see
// [Transitioning to ARC Release Notes].
//
// [Advanced Memory Management Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/MemoryMgmt/Articles/MemoryMgmt.html#//apple_ref/doc/uid/10000011i
// [Transitioning to ARC Release Notes]: https://developer.apple.com/library/archive/releasenotes/ObjectiveC/RN-TransitioningToARC/Introduction/Introduction.html#//apple_ref/doc/uid/TP40011226
//
// See: https://developer.apple.com/documentation/ObjectiveC/NSObject-c.protocol/retain

func (o NSObjectObject) Retain() IObject {
	
	rv := objc.Send[objc.ID](o.ID, objc.Sel("retain"))
	return Object{ID: rv}
	}

