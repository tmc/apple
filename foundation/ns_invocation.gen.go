// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSInvocation] class.
var (
	_NSInvocationClass     NSInvocationClass
	_NSInvocationClassOnce sync.Once
)

func getNSInvocationClass() NSInvocationClass {
	_NSInvocationClassOnce.Do(func() {
		_NSInvocationClass = NSInvocationClass{class: objc.GetClass("NSInvocation")}
	})
	return _NSInvocationClass
}

// GetNSInvocationClass returns the class object for NSInvocation.
func GetNSInvocationClass() NSInvocationClass {
	return getNSInvocationClass()
}

type NSInvocationClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSInvocationClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSInvocationClass) Alloc() NSInvocation {
	rv := objc.Send[NSInvocation](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An Objective-C message rendered as an object.
//
// # Overview
//
// [NSInvocation] objects are used to store and forward messages between
// objects and between applications, primarily by [NSTimer] objects and the
// distributed objects system. An [NSInvocation] object contains all the
// elements of an Objective-C message: a target, a selector, arguments, and
// the return value. Each of these elements can be set directly, and the
// return value is set automatically when the [NSInvocation] object is
// dispatched.
//
// An [NSInvocation] object can be repeatedly dispatched to different targets;
// its arguments can be modified between dispatch for varying results; even
// its selector can be changed to another with the same method signature
// (argument and return types). This flexibility makes [NSInvocation] useful
// for repeating messages with many arguments and variations; rather than
// retyping a slightly different expression for each message, you modify the
// [NSInvocation] object as needed each time before dispatching it to a new
// target.
//
// [NSInvocation] does not support invocations of methods with either variable
// numbers of arguments or `union` arguments. You should use the
// [NSInvocation.InvocationWithMethodSignature] class method to create [NSInvocation]
// objects; you should not create these objects using [alloc] and [init()].
//
// This class does not retain the arguments for the contained invocation by
// default. If those objects might disappear between the time you create your
// instance of [NSInvocation] and the time you use it, you should explicitly
// retain the objects yourself or invoke the [NSInvocation.RetainArguments] method to have
// the invocation object retain them itself.
//
// # Configuring an Invocation Object
//
//   - [NSInvocation.Selector]: The receiver’s selector, or 0 if it hasn’t been set.
//   - [NSInvocation.SetSelector]
//   - [NSInvocation.Target]: The receiver’s target, or `nil` if the receiver has no target.
//   - [NSInvocation.SetTarget]
//   - [NSInvocation.SetArgumentAtIndex]: Sets an argument of the receiver.
//   - [NSInvocation.GetArgumentAtIndex]: Returns by indirection the receiver’s argument at a specified index.
//   - [NSInvocation.ArgumentsRetained]: A Boolean value that indicates if the receiver has retained its arguments.
//   - [NSInvocation.RetainArguments]: If the receiver hasn’t already done so, retains the target and all object arguments of the receiver and copies all of its C-string arguments and blocks. If a returnvalue has been set, this is also retained or copied.
//   - [NSInvocation.SetReturnValue]: Sets the receiver’s return value.
//   - [NSInvocation.GetReturnValue]: Gets the invocation’s return value.
//
// # Dispatching an Invocation
//
//   - [NSInvocation.Invoke]: Sends the receiver’s message (with arguments) to its target and sets the return value.
//   - [NSInvocation.InvokeWithTarget]: Sets the receiver’s target, sends the receiver’s message (with arguments) to that target, and sets the return value.
//
// # Getting the Method Signature
//
//   - [NSInvocation.MethodSignature]: The receiver’s method signature.
//
// # Instance Methods
//
//   - [NSInvocation.InvokeUsingIMP]
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocation
//
// [alloc]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/alloc
// [init()]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/init()
type NSInvocation struct {
	objectivec.Object
}

// NSInvocationFromID constructs a [NSInvocation] from an objc.ID.
//
// An Objective-C message rendered as an object.
func NSInvocationFromID(id objc.ID) NSInvocation {
	return NSInvocation{objectivec.Object{ID: id}}
}

// Ensure NSInvocation implements INSInvocation.
var _ INSInvocation = NSInvocation{}

// An interface definition for the [NSInvocation] class.
//
// # Configuring an Invocation Object
//
//   - [INSInvocation.Selector]: The receiver’s selector, or 0 if it hasn’t been set.
//   - [INSInvocation.SetSelector]
//   - [INSInvocation.Target]: The receiver’s target, or `nil` if the receiver has no target.
//   - [INSInvocation.SetTarget]
//   - [INSInvocation.SetArgumentAtIndex]: Sets an argument of the receiver.
//   - [INSInvocation.GetArgumentAtIndex]: Returns by indirection the receiver’s argument at a specified index.
//   - [INSInvocation.ArgumentsRetained]: A Boolean value that indicates if the receiver has retained its arguments.
//   - [INSInvocation.RetainArguments]: If the receiver hasn’t already done so, retains the target and all object arguments of the receiver and copies all of its C-string arguments and blocks. If a returnvalue has been set, this is also retained or copied.
//   - [INSInvocation.SetReturnValue]: Sets the receiver’s return value.
//   - [INSInvocation.GetReturnValue]: Gets the invocation’s return value.
//
// # Dispatching an Invocation
//
//   - [INSInvocation.Invoke]: Sends the receiver’s message (with arguments) to its target and sets the return value.
//   - [INSInvocation.InvokeWithTarget]: Sets the receiver’s target, sends the receiver’s message (with arguments) to that target, and sets the return value.
//
// # Getting the Method Signature
//
//   - [INSInvocation.MethodSignature]: The receiver’s method signature.
//
// # Instance Methods
//
//   - [INSInvocation.InvokeUsingIMP]
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocation
type INSInvocation interface {
	objectivec.IObject

	// Topic: Configuring an Invocation Object

	// The receiver’s selector, or 0 if it hasn’t been set.
	Selector() objc.SEL
	SetSelector(value objc.SEL)
	// The receiver’s target, or `nil` if the receiver has no target.
	Target() objectivec.IObject
	SetTarget(value objectivec.IObject)
	// Sets an argument of the receiver.
	SetArgumentAtIndex(argumentLocation unsafe.Pointer, idx int)
	// Returns by indirection the receiver’s argument at a specified index.
	GetArgumentAtIndex(argumentLocation unsafe.Pointer, idx int)
	// A Boolean value that indicates if the receiver has retained its arguments.
	ArgumentsRetained() bool
	// If the receiver hasn’t already done so, retains the target and all object arguments of the receiver and copies all of its C-string arguments and blocks. If a returnvalue has been set, this is also retained or copied.
	RetainArguments()
	// Sets the receiver’s return value.
	SetReturnValue(retLoc unsafe.Pointer)
	// Gets the invocation’s return value.
	GetReturnValue(retLoc unsafe.Pointer)

	// Topic: Dispatching an Invocation

	// Sends the receiver’s message (with arguments) to its target and sets the return value.
	Invoke()
	// Sets the receiver’s target, sends the receiver’s message (with arguments) to that target, and sets the return value.
	InvokeWithTarget(target objectivec.IObject)

	// Topic: Getting the Method Signature

	// The receiver’s method signature.
	MethodSignature() INSMethodSignature

	// Topic: Instance Methods

	InvokeUsingIMP(imp objectivec.IMP)
}

// Init initializes the instance.
func (i NSInvocation) Init() NSInvocation {
	rv := objc.Send[NSInvocation](i.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i NSInvocation) Autorelease() NSInvocation {
	rv := objc.Send[NSInvocation](i.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSInvocation creates a new NSInvocation instance.
func NewNSInvocation() NSInvocation {
	class := getNSInvocationClass()
	rv := objc.Send[NSInvocation](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Sets an argument of the receiver.
//
// argumentLocation: An untyped buffer containing an argument to be assigned to the receiver.
// See the discussion below relating to argument values that are objects.
//
// idx: An integer specifying the index of the argument.
//
// Indices 0 and 1 indicate the hidden arguments `self` and `_cmd`,
// respectively; you should set these values directly with the [Target] and
// [Selector] properties. Use indices 2 and greater for the arguments normally
// passed in a message.
//
// # Discussion
//
// This method copies the contents of `buffer` as the argument at `index`. The
// number of bytes copied is determined by the argument size.
//
// When the argument value is an object, pass a pointer to the variable (or
// memory) from which the object should be copied:
//
// This method raises [NSInvalidArgumentException] if the value of `index` is
// greater than the actual number of arguments for the selector.
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocation/setArgument:atIndex:
func (i NSInvocation) SetArgumentAtIndex(argumentLocation unsafe.Pointer, idx int) {
	objc.Send[objc.ID](i.ID, objc.Sel("setArgument:atIndex:"), argumentLocation, idx)
}

// Returns by indirection the receiver’s argument at a specified index.
//
// argumentLocation: An untyped buffer to hold the returned argument. See the discussion below
// relating to argument values that are objects.
//
// idx: An integer specifying the index of the argument to get.
//
// Indices 0 and 1 indicate the hidden arguments `self` and `_cmd`,
// respectively; these values can be retrieved directly with the `target` and
// `selector` methods. Use indices 2 and greater for the arguments normally
// passed in a message.
//
// # Discussion
//
// This method copies the argument stored at `index` into the storage pointed
// to by `buffer`. The size of `buffer` must be large enough to accommodate
// the argument value. When the argument value is an object, pass a pointer to
// the variable (or memory) into which the object should be placed.
//
// In the following example, `myInvocation` represents a call to a
// two-argument method called “ in a class called [MyClass], which takes an
// [NSMutableString] and an `int`. The example performs the invocation with
// [Invoke], then retrieves the first argument with [GetArgumentAtIndex] and
// copies it to a strongly-held property called `myObject`.
//
// This method raises [InvalidArgumentException] if `index` is greater than
// the actual number of arguments for the selector.
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocation/getArgument:atIndex:
func (i NSInvocation) GetArgumentAtIndex(argumentLocation unsafe.Pointer, idx int) {
	objc.Send[objc.ID](i.ID, objc.Sel("getArgument:atIndex:"), argumentLocation, idx)
}

// If the receiver hasn’t already done so, retains the target and all object
// arguments of the receiver and copies all of its C-string arguments and
// blocks. If a returnvalue has been set, this is also retained or copied.
//
// # Discussion
//
// Before this method is invoked, [ArgumentsRetained] returns false; after, it
// returns true.
//
// For efficiency, newly created [NSInvocation] objects don’t retain or copy
// their arguments, nor do they retain their targets, copy C strings, or copy
// any associated blocks. You should instruct an [NSInvocation] object to
// retain its arguments if you intend to cache it, because the arguments may
// otherwise be released before the invocation is invoked. [NSTimer] objects
// always instruct their invocations to retain their arguments, for example,
// because there’s usually a delay before a timer fires.
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocation/retainArguments
func (i NSInvocation) RetainArguments() {
	objc.Send[objc.ID](i.ID, objc.Sel("retainArguments"))
}

// Sets the receiver’s return value.
//
// retLoc: An untyped buffer whose contents are copied as the receiver’s return
// value.
//
// # Discussion
//
// This value is normally set when you send an [Invoke] or [InvokeWithTarget]
// message.
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocation/setReturnValue:
func (i NSInvocation) SetReturnValue(retLoc unsafe.Pointer) {
	objc.Send[objc.ID](i.ID, objc.Sel("setReturnValue:"), retLoc)
}

// Gets the invocation’s return value.
//
// retLoc: An untyped buffer into which the invocation copies its return value. It
// should be large enough to accommodate the value. See the discussion below
// for more information about `buffer`.
//
// # Discussion
//
// Use the [NSMethodSignature] method [MethodReturnLength] to determine the
// size needed for `buffer`:
//
// When the return value is an object, pass a pointer to the variable (or
// memory) into which [NSInvocation] should place the object. In the following
// example, `myInvocation` represents a call to a no-argument method called
// `createReturnValue` in a class called [MyClass], which returns an
// [NSMutableString]. The example performs the invocation with [Invoke], then
// retrieves the object with [GetReturnValue] and copies it to a strongly-held
// property called `myObject`.
//
// If you haven’t invoked the [NSInvocation] object, the result of this
// method is undefined.
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocation/getReturnValue:
func (i NSInvocation) GetReturnValue(retLoc unsafe.Pointer) {
	objc.Send[objc.ID](i.ID, objc.Sel("getReturnValue:"), retLoc)
}

// Sends the receiver’s message (with arguments) to its target and sets the
// return value.
//
// # Discussion
//
// You must set the receiver’s target, selector, and argument values before
// calling this method.
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocation/invoke
func (i NSInvocation) Invoke() {
	objc.Send[objc.ID](i.ID, objc.Sel("invoke"))
}

// Sets the receiver’s target, sends the receiver’s message (with
// arguments) to that target, and sets the return value.
//
// target: The object to set as the receiver’s target.
//
// # Discussion
//
// You must set the receiver’s selector and argument values before calling
// this method.
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocation/invokeWithTarget:
func (i NSInvocation) InvokeWithTarget(target objectivec.IObject) {
	objc.Send[objc.ID](i.ID, objc.Sel("invokeWithTarget:"), target)
}

// See: https://developer.apple.com/documentation/Foundation/NSInvocation/invokeUsingIMP:
func (i NSInvocation) InvokeUsingIMP(imp objectivec.IMP) {
	objc.Send[objc.ID](i.ID, objc.Sel("invokeUsingIMP:"), imp)
}

// Returns an [NSInvocation] object able to construct messages using a given
// method signature.
//
// sig: An object encapsulating a method signature.
//
// # Discussion
//
// The new object must have its selector set with [NSInvocation] and its
// arguments set with [SetArgumentAtIndex] before it can be invoked. Do not
// use the [alloc]/[init()] approach to create [NSInvocation] objects.
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocation/invocationWithMethodSignature:
//
// [alloc]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/alloc
// [init()]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/init()
func (_NSInvocationClass NSInvocationClass) InvocationWithMethodSignature(sig INSMethodSignature) NSInvocation {
	rv := objc.Send[objc.ID](objc.ID(_NSInvocationClass.class), objc.Sel("invocationWithMethodSignature:"), sig)
	return NSInvocationFromID(rv)
}

// The receiver’s selector, or 0 if it hasn’t been set.
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocation/selector
func (i NSInvocation) Selector() objc.SEL {
	rv := objc.Send[objc.SEL](i.ID, objc.Sel("selector"))
	return rv
}
func (i NSInvocation) SetSelector(value objc.SEL) {
	objc.Send[struct{}](i.ID, objc.Sel("setSelector:"), value)
}

// The receiver’s target, or `nil` if the receiver has no target.
//
// # Discussion
//
// The target is the receiver of the message sent by [Invoke].
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocation/target
func (i NSInvocation) Target() objectivec.IObject {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("target"))
	return objectivec.Object{ID: rv}
}
func (i NSInvocation) SetTarget(value objectivec.IObject) {
	objc.Send[struct{}](i.ID, objc.Sel("setTarget:"), value)
}

// A Boolean value that indicates if the receiver has retained its arguments.
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocation/argumentsRetained
func (i NSInvocation) ArgumentsRetained() bool {
	rv := objc.Send[bool](i.ID, objc.Sel("argumentsRetained"))
	return rv
}

// The receiver’s method signature.
//
// See: https://developer.apple.com/documentation/Foundation/NSInvocation/methodSignature
func (i NSInvocation) MethodSignature() INSMethodSignature {
	rv := objc.Send[objc.ID](i.ID, objc.Sel("methodSignature"))
	return NSMethodSignatureFromID(objc.ID(rv))
}
