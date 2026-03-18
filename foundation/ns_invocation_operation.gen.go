// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSInvocationOperation] class.
var (
	_NSInvocationOperationClass     NSInvocationOperationClass
	_NSInvocationOperationClassOnce sync.Once
)

func getNSInvocationOperationClass() NSInvocationOperationClass {
	_NSInvocationOperationClassOnce.Do(func() {
		_NSInvocationOperationClass = NSInvocationOperationClass{class: objc.GetClass("NSInvocationOperation")}
	})
	return _NSInvocationOperationClass
}

// GetNSInvocationOperationClass returns the class object for NSInvocationOperation.
func GetNSInvocationOperationClass() NSInvocationOperationClass {
	return getNSInvocationOperationClass()
}

type NSInvocationOperationClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSInvocationOperationClass) Alloc() NSInvocationOperation {
	rv := objc.Send[NSInvocationOperation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An operation that manages the execution of a single encapsulated task
// specified as an invocation.
//
// # Overview
// 
// The [NSInvocationOperation] class is a concrete subclass of [NSOperation]
// that you use to initiate an operation that consists of invoking a selector
// on a specified object. This class implements a non-concurrent operation.
// 
// For more information on concurrent versus non-concurrent operations, see
// [NSOperation].
//
// # Initialization
//
//   - [NSInvocationOperation.InitWithTargetSelectorObject]: Returns an [NSInvocationOperation] object initialized with the specified target and selector.
//   - [NSInvocationOperation.InitWithInvocation]: Returns an [NSInvocationOperation] object initialized with the specified invocation object.
//
// # Getting Attributes
//
//   - [NSInvocationOperation.Invocation]: The receiver’s invocation object.
//   - [NSInvocationOperation.Result]: The result of the invocation or method.
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocationOperation
type NSInvocationOperation struct {
	NSOperation
}

// NSInvocationOperationFromID constructs a [NSInvocationOperation] from an objc.ID.
//
// An operation that manages the execution of a single encapsulated task
// specified as an invocation.
func NSInvocationOperationFromID(id objc.ID) NSInvocationOperation {
	return NSInvocationOperation{NSOperation: NSOperationFromID(id)}
}
// Ensure NSInvocationOperation implements INSInvocationOperation.
var _ INSInvocationOperation = NSInvocationOperation{}

// An interface definition for the [NSInvocationOperation] class.
//
// # Initialization
//
//   - [INSInvocationOperation.InitWithTargetSelectorObject]: Returns an [NSInvocationOperation] object initialized with the specified target and selector.
//   - [INSInvocationOperation.InitWithInvocation]: Returns an [NSInvocationOperation] object initialized with the specified invocation object.
//
// # Getting Attributes
//
//   - [INSInvocationOperation.Invocation]: The receiver’s invocation object.
//   - [INSInvocationOperation.Result]: The result of the invocation or method.
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocationOperation
type INSInvocationOperation interface {
	INSOperation

	// Topic: Initialization

	// Returns an [NSInvocationOperation] object initialized with the specified target and selector.
	InitWithTargetSelectorObject(target objectivec.IObject, sel objc.SEL, arg objectivec.IObject) NSInvocationOperation
	// Returns an [NSInvocationOperation] object initialized with the specified invocation object.
	InitWithInvocation(inv INSInvocation) NSInvocationOperation

	// Topic: Getting Attributes

	// The receiver’s invocation object.
	Invocation() INSInvocation
	// The result of the invocation or method.
	Result() objectivec.IObject
}

// Init initializes the instance.
func (i NSInvocationOperation) Init() NSInvocationOperation {
	rv := objc.Send[NSInvocationOperation](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i NSInvocationOperation) Autorelease() NSInvocationOperation {
	rv := objc.Send[NSInvocationOperation](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSInvocationOperation creates a new NSInvocationOperation instance.
func NewNSInvocationOperation() NSInvocationOperation {
	class := getNSInvocationOperationClass()
	rv := objc.Send[NSInvocationOperation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns an [NSInvocationOperation] object initialized with the specified
// invocation object.
//
// inv: The invocation object identifying the target object, selector, and
// parameter objects.
//
// # Return Value
// 
// An initialized [NSInvocationOperation] object or `nil` if the object could
// not be initialized.
//
// # Discussion
// 
// This method is the designated initializer. The receiver tells the
// invocation object to retain its arguments.
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocationOperation/initWithInvocation:
func NewInvocationOperationWithInvocation(inv INSInvocation) NSInvocationOperation {
	instance := getNSInvocationOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithInvocation:"), inv)
	return NSInvocationOperationFromID(rv)
}

// Returns an [NSInvocationOperation] object initialized with the specified
// target and selector.
//
// target: The object defining the specified selector.
//
// sel: The selector to invoke when running the operation. The selector may take 0
// or 1 parameters; if it accepts a parameter, the type of that parameter must
// be `id`. The return type of the method may be `void`, a scalar value, or an
// object that can be returned as an `id` type.
//
// arg: The parameter object to pass to the selector. If the selector does not take
// an argument, specify `nil`.
//
// # Return Value
// 
// An initialized [NSInvocationOperation] object or `nil` if the target object
// does not implement the specified selector.
//
// # Discussion
// 
// If you specify a selector with a non-void return type, you can get the
// return value by calling the [Result] method after the operation finishes
// executing. The receiver tells the invocation object to retain its
// arguments.
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocationOperation/initWithTarget:selector:object:
func NewInvocationOperationWithTargetSelectorObject(target objectivec.IObject, sel objc.SEL, arg objectivec.IObject) NSInvocationOperation {
	instance := getNSInvocationOperationClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithTarget:selector:object:"), target, sel, arg)
	return NSInvocationOperationFromID(rv)
}

// Returns an [NSInvocationOperation] object initialized with the specified
// target and selector.
//
// target: The object defining the specified selector.
//
// sel: The selector to invoke when running the operation. The selector may take 0
// or 1 parameters; if it accepts a parameter, the type of that parameter must
// be `id`. The return type of the method may be `void`, a scalar value, or an
// object that can be returned as an `id` type.
//
// arg: The parameter object to pass to the selector. If the selector does not take
// an argument, specify `nil`.
//
// # Return Value
// 
// An initialized [NSInvocationOperation] object or `nil` if the target object
// does not implement the specified selector.
//
// # Discussion
// 
// If you specify a selector with a non-void return type, you can get the
// return value by calling the [Result] method after the operation finishes
// executing. The receiver tells the invocation object to retain its
// arguments.
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocationOperation/initWithTarget:selector:object:
func (i NSInvocationOperation) InitWithTargetSelectorObject(target objectivec.IObject, sel objc.SEL, arg objectivec.IObject) NSInvocationOperation {
	rv := objc.Send[NSInvocationOperation](i.ID, objc.Sel("initWithTarget:selector:object:"), target, sel, arg)
	return rv
}

// Returns an [NSInvocationOperation] object initialized with the specified
// invocation object.
//
// inv: The invocation object identifying the target object, selector, and
// parameter objects.
//
// # Return Value
// 
// An initialized [NSInvocationOperation] object or `nil` if the object could
// not be initialized.
//
// # Discussion
// 
// This method is the designated initializer. The receiver tells the
// invocation object to retain its arguments.
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocationOperation/initWithInvocation:
func (i NSInvocationOperation) InitWithInvocation(inv INSInvocation) NSInvocationOperation {
	rv := objc.Send[NSInvocationOperation](i.ID, objc.Sel("initWithInvocation:"), inv)
	return rv
}

// The receiver’s invocation object.
//
// # Discussion
// 
// The invocation object identifying the target object, selector, and
// parameters to use to execute the operation’s task.
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocationOperation/invocation
func (i NSInvocationOperation) Invocation() INSInvocation {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("invocation"))
	return NSInvocationFromID(objc.ID(rv))
}

// The result of the invocation or method.
//
// # Discussion
// 
// The object returned by the method or an [NSValue] object containing the
// return value if it is not an object. `nil` if the method or invocation is
// not finished executing.
// 
// If an exception was raised during the execution of the method or
// invocation, accessing this property raises that exception again. If the
// operation was cancelled or the invocation or method has a `void` return
// type, accessing this property raises an exception; see [Result Exceptions].
//
// [Result Exceptions]: https://developer.apple.com/documentation/Foundation/result-exceptions
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocationOperation/result
func (i NSInvocationOperation) Result() objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("result"))
	return objectivec.Object{ID: rv}
}

