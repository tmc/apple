// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSMethodSignature] class.
var (
	_NSMethodSignatureClass     NSMethodSignatureClass
	_NSMethodSignatureClassOnce sync.Once
)

func getNSMethodSignatureClass() NSMethodSignatureClass {
	_NSMethodSignatureClassOnce.Do(func() {
		_NSMethodSignatureClass = NSMethodSignatureClass{class: objc.GetClass("NSMethodSignature")}
	})
	return _NSMethodSignatureClass
}

// GetNSMethodSignatureClass returns the class object for NSMethodSignature.
func GetNSMethodSignatureClass() NSMethodSignatureClass {
	return getNSMethodSignatureClass()
}

type NSMethodSignatureClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSMethodSignatureClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSMethodSignatureClass) Alloc() NSMethodSignature {
	rv := objc.Send[NSMethodSignature](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// A record of the type information for the return value and parameters of a
// method.
//
// # Overview
// 
// Use an [NSMethodSignature] object to forward messages that the receiving
// object does not respond to—most notably in the case of distributed
// objects. You typically create an [NSMethodSignature] object using the
// [NSObject] [methodSignatureForSelector:] instance method (in macOS 10.5 and
// later you can also use [NSMethodSignature.SignatureWithObjCTypes]). It is then used to create
// an [NSInvocation] object, which is passed as the argument to a
// [forwardInvocation:] message to send the invocation on to whatever other
// object can handle the message. In the default case, [NSObject] invokes
// [doesNotRecognizeSelector(_:)], which raises an exception. For distributed
// objects, the [NSInvocation] object is encoded using the information in the
// [NSMethodSignature] object and sent to the real object represented by the
// receiver of the message.
// 
// # Type Encodings
// 
// An [NSMethodSignature] object is initialized with an array of characters
// representing the string encoding of return and argument types for a method.
// You can get the string encoding of a particular type using the `@encode()`
// compiler directive. Because string encodings are implementation-specific,
// you should not hard-code these values.
// 
// A method signature consists of one or more characters for the method return
// type, followed by the string encodings of the implicit arguments `self` and
// `_cmd`, followed by zero or more explicit arguments. You can determine the
// string encoding and the length of a return type using [NSMethodSignature.MethodReturnType]
// and [NSMethodSignature.MethodReturnLength] properties. You can access arguments individually
// using the [NSMethodSignature.GetArgumentTypeAtIndex] method and [NSMethodSignature.NumberOfArguments] property.
// 
// For example, the [NSString] instance method [ContainsString] has a method
// signature with the following arguments:
// 
// - `@encode(BOOL)` (`c`) for the return type - `@encode(id)` (`@`) for the
// receiver (`self`) - `@encode(SEL)` (`:`) for the selector (`_cmd`) -
// `@encode(NSString *)` (`@`) for the first explicit argument
// 
// See [Type Encodings] in [Objective-C Runtime Programming Guide] for more
// information.
//
// [NSObject]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class
// [Objective-C Runtime Programming Guide]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ObjCRuntimeGuide/Introduction/Introduction.html#//apple_ref/doc/uid/TP40008048
// [Type Encodings]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/ObjCRuntimeGuide/Articles/ocrtTypeEncodings.html#//apple_ref/doc/uid/TP40008048-CH100
// [doesNotRecognizeSelector(_:)]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/doesNotRecognizeSelector(_:)
// [forwardInvocation:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/forwardInvocation:
// [methodSignatureForSelector:]: https://developer.apple.com/documentation/ObjectiveC/NSObject-swift.class/methodSignatureForSelector:
//
// # Getting Information on Argument Types
//
//   - [NSMethodSignature.GetArgumentTypeAtIndex]: Returns the type encoding for the argument at a given index.
//   - [NSMethodSignature.NumberOfArguments]: The number of arguments recorded in the receiver.
//   - [NSMethodSignature.FrameLength]: The number of bytes that the arguments, taken together, occupy on the stack.
//
// # Getting Information on Return Types
//
//   - [NSMethodSignature.MethodReturnType]: A C string encoding the return type of the method in Objective-C type encoding.
//   - [NSMethodSignature.MethodReturnLength]: The number of bytes required for the return value.
//
// # Determining Synchronous Status
//
//   - [NSMethodSignature.IsOneway]: Whether the receiver is asynchronous when invoked through distributed objects.
//
// See: https://developer.apple.com/documentation/Foundation/NSMethodSignature
type NSMethodSignature struct {
	objectivec.Object
}

// NSMethodSignatureFromID constructs a [NSMethodSignature] from an objc.ID.
//
// A record of the type information for the return value and parameters of a
// method.
func NSMethodSignatureFromID(id objc.ID) NSMethodSignature {
	return NSMethodSignature{objectivec.Object{ID: id}}
}
// Ensure NSMethodSignature implements INSMethodSignature.
var _ INSMethodSignature = NSMethodSignature{}

// An interface definition for the [NSMethodSignature] class.
//
// # Getting Information on Argument Types
//
//   - [INSMethodSignature.GetArgumentTypeAtIndex]: Returns the type encoding for the argument at a given index.
//   - [INSMethodSignature.NumberOfArguments]: The number of arguments recorded in the receiver.
//   - [INSMethodSignature.FrameLength]: The number of bytes that the arguments, taken together, occupy on the stack.
//
// # Getting Information on Return Types
//
//   - [INSMethodSignature.MethodReturnType]: A C string encoding the return type of the method in Objective-C type encoding.
//   - [INSMethodSignature.MethodReturnLength]: The number of bytes required for the return value.
//
// # Determining Synchronous Status
//
//   - [INSMethodSignature.IsOneway]: Whether the receiver is asynchronous when invoked through distributed objects.
//
// See: https://developer.apple.com/documentation/Foundation/NSMethodSignature
type INSMethodSignature interface {
	objectivec.IObject

	// Topic: Getting Information on Argument Types

	// Returns the type encoding for the argument at a given index.
	GetArgumentTypeAtIndex(idx uint) string
	// The number of arguments recorded in the receiver.
	NumberOfArguments() uint
	// The number of bytes that the arguments, taken together, occupy on the stack.
	FrameLength() uint

	// Topic: Getting Information on Return Types

	// A C string encoding the return type of the method in Objective-C type encoding.
	MethodReturnType() string
	// The number of bytes required for the return value.
	MethodReturnLength() uint

	// Topic: Determining Synchronous Status

	// Whether the receiver is asynchronous when invoked through distributed objects.
	IsOneway() bool
}

// Init initializes the instance.
func (m NSMethodSignature) Init() NSMethodSignature {
	rv := objc.Send[NSMethodSignature](m.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m NSMethodSignature) Autorelease() NSMethodSignature {
	rv := objc.Send[NSMethodSignature](m.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSMethodSignature creates a new NSMethodSignature instance.
func NewNSMethodSignature() NSMethodSignature {
	class := getNSMethodSignatureClass()
	rv := objc.Send[NSMethodSignature](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns the type encoding for the argument at a given index.
//
// idx: The index of the argument to get.
//
// # Return Value
// 
// The type encoding for the argument at `idx`.
//
// # Discussion
// 
// Indexes begin with 0. The implicit arguments `self` (of type `id`) and
// `_cmd` (of type [SEL]) are at indexes 0 and 1; explicit arguments begin at
// index 2.
//
// See: https://developer.apple.com/documentation/Foundation/NSMethodSignature/getArgumentTypeAtIndex:
func (m NSMethodSignature) GetArgumentTypeAtIndex(idx uint) string {
	rv := objc.Send[*byte](m.ID, objc.Sel("getArgumentTypeAtIndex:"), idx)
	return objc.GoString(rv)
}
// Whether the receiver is asynchronous when invoked through distributed
// objects.
//
// # Return Value
// 
// [true] if the receiver is asynchronous when invoked through distributed
// objects, otherwise [false].
//
// [false]: https://developer.apple.com/documentation/Swift/false
// [true]: https://developer.apple.com/documentation/Swift/true
//
// # Discussion
// 
// If the method is `oneway`, the sender of the remote message doesn’t block
// awaiting a reply.
//
// See: https://developer.apple.com/documentation/Foundation/NSMethodSignature/isOneway
func (m NSMethodSignature) IsOneway() bool {
	rv := objc.Send[bool](m.ID, objc.Sel("isOneway"))
	return rv
}

// Returns an [NSMethodSignature] object for the given Objective-C method type
// string.
//
// types: An array of characters containing the type encodings for the method
// arguments.
//
// # Return Value
// 
// An [NSMethodSignature] object for the given Objective-C method type string
// in `types`.
//
// See: https://developer.apple.com/documentation/Foundation/NSMethodSignature/signatureWithObjCTypes:
func (_NSMethodSignatureClass NSMethodSignatureClass) SignatureWithObjCTypes(types string) NSMethodSignature {
	rv := objc.Send[objc.ID](objc.ID(_NSMethodSignatureClass.class), objc.Sel("signatureWithObjCTypes:"), unsafe.Pointer(unsafe.StringData(types + "\x00")))
	return NSMethodSignatureFromID(rv)
}

// The number of arguments recorded in the receiver.
//
// # Discussion
// 
// There are always at least two arguments, because an [NSMethodSignature]
// object includes the implicit arguments `self` and `_cmd`, which are the
// first two arguments passed to every method implementation.
//
// See: https://developer.apple.com/documentation/Foundation/NSMethodSignature/numberOfArguments
func (m NSMethodSignature) NumberOfArguments() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("numberOfArguments"))
	return rv
}
// The number of bytes that the arguments, taken together, occupy on the
// stack.
//
// # Discussion
// 
// This number varies with the hardware architecture the application runs on.
//
// See: https://developer.apple.com/documentation/Foundation/NSMethodSignature/frameLength
func (m NSMethodSignature) FrameLength() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("frameLength"))
	return rv
}
// A C string encoding the return type of the method in Objective-C type
// encoding.
//
// See: https://developer.apple.com/documentation/Foundation/NSMethodSignature/methodReturnType
func (m NSMethodSignature) MethodReturnType() string {
	rv := objc.Send[*byte](m.ID, objc.Sel("methodReturnType"))
	return objc.GoString(rv)
}
// The number of bytes required for the return value.
//
// See: https://developer.apple.com/documentation/Foundation/NSMethodSignature/methodReturnLength
func (m NSMethodSignature) MethodReturnLength() uint {
	rv := objc.Send[uint](m.ID, objc.Sel("methodReturnLength"))
	return rv
}

