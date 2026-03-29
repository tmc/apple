// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSException] class.
var (
	_NSExceptionClass     NSExceptionClass
	_NSExceptionClassOnce sync.Once
)

func getNSExceptionClass() NSExceptionClass {
	_NSExceptionClassOnce.Do(func() {
		_NSExceptionClass = NSExceptionClass{class: objc.GetClass("NSException")}
	})
	return _NSExceptionClass
}

// GetNSExceptionClass returns the class object for NSException.
func GetNSExceptionClass() NSExceptionClass {
	return getNSExceptionClass()
}

type NSExceptionClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSExceptionClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSExceptionClass) Alloc() NSException {
	rv := objc.Send[NSException](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that represents a special condition that interrupts the normal
// flow of program execution.
//
// # Overview
// 
// Use [NSException] to implement exception handling. An exception is a
// special condition that interrupts the normal flow of program execution.
// Each application can interrupt the program for different reasons. For
// example, one application might interpret saving a file in a directory that
// is write-protected as an exception. In this sense, the exception is
// equivalent to an error. Another application might interpret the user’s
// key-press (for example, Control-C) as an exception: an indication that a
// long-running process should abort.
//
// # Creating and Raising an NSException Object
//
//   - [NSException.InitWithNameReasonUserInfo]: Initializes and returns a newly allocated exception object.
//   - [NSException.Raise]: Raises the receiver, causing program flow to jump to the local exception handler.
//
// # Querying an NSException Object
//
//   - [NSException.Name]: A string used to uniquely identify the receiver.
//   - [NSException.Reason]: A string containing a “human-readable” reason for the receiver.
//   - [NSException.UserInfo]: A dictionary containing application-specific data pertaining to the receiver.
//
// # Getting Exception Stack Frames
//
//   - [NSException.CallStackReturnAddresses]: The call return addresses related to a raised exception.
//   - [NSException.CallStackSymbols]: An array containing the current call stack symbols.
//
// See: https://developer.apple.com/documentation/Foundation/NSException
type NSException struct {
	objectivec.Object
}

// NSExceptionFromID constructs a [NSException] from an objc.ID.
//
// An object that represents a special condition that interrupts the normal
// flow of program execution.
func NSExceptionFromID(id objc.ID) NSException {
	return NSException{objectivec.Object{ID: id}}
}
// NOTE: NSException adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSException] class.
//
// # Creating and Raising an NSException Object
//
//   - [INSException.InitWithNameReasonUserInfo]: Initializes and returns a newly allocated exception object.
//   - [INSException.Raise]: Raises the receiver, causing program flow to jump to the local exception handler.
//
// # Querying an NSException Object
//
//   - [INSException.Name]: A string used to uniquely identify the receiver.
//   - [INSException.Reason]: A string containing a “human-readable” reason for the receiver.
//   - [INSException.UserInfo]: A dictionary containing application-specific data pertaining to the receiver.
//
// # Getting Exception Stack Frames
//
//   - [INSException.CallStackReturnAddresses]: The call return addresses related to a raised exception.
//   - [INSException.CallStackSymbols]: An array containing the current call stack symbols.
//
// See: https://developer.apple.com/documentation/Foundation/NSException
type INSException interface {
	objectivec.IObject
	NSCoding
	NSCopying
	NSSecureCoding

	// Topic: Creating and Raising an NSException Object

	// Initializes and returns a newly allocated exception object.
	InitWithNameReasonUserInfo(aName NSExceptionName, aReason string, aUserInfo INSDictionary) NSException
	// Raises the receiver, causing program flow to jump to the local exception handler.
	Raise()

	// Topic: Querying an NSException Object

	// A string used to uniquely identify the receiver.
	Name() NSExceptionName
	// A string containing a “human-readable” reason for the receiver.
	Reason() string
	// A dictionary containing application-specific data pertaining to the receiver.
	UserInfo() INSDictionary

	// Topic: Getting Exception Stack Frames

	// The call return addresses related to a raised exception.
	CallStackReturnAddresses() []NSNumber
	// An array containing the current call stack symbols.
	CallStackSymbols() []string
}

// Init initializes the instance.
func (e NSException) Init() NSException {
	rv := objc.Send[NSException](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e NSException) Autorelease() NSException {
	rv := objc.Send[NSException](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSException creates a new NSException instance.
func NewNSException() NSException {
	class := getNSExceptionClass()
	rv := objc.Send[NSException](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func NewExceptionWithCoder(coder INSCoder) NSException {
	instance := getNSExceptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithCoder:"), coder)
	return NSExceptionFromID(rv)
}

// Initializes and returns a newly allocated exception object.
//
// aName: The name of the exception.
//
// aReason: A human-readable message string summarizing the reason for the exception.
//
// aUserInfo: A dictionary containing user-defined information relating to the exception
//
// # Return Value
// 
// The created [NSException] object or `nil` if the object couldn’t be
// created.
//
// # Discussion
// 
// This is the designated initializer.
//
// See: https://developer.apple.com/documentation/Foundation/NSException/init(name:reason:userInfo:)
func NewExceptionWithNameReasonUserInfo(aName NSExceptionName, aReason string, aUserInfo INSDictionary) NSException {
	instance := getNSExceptionClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithName:reason:userInfo:"), objc.String(string(aName)), objc.String(aReason), aUserInfo)
	return NSExceptionFromID(rv)
}

// Initializes and returns a newly allocated exception object.
//
// aName: The name of the exception.
//
// aReason: A human-readable message string summarizing the reason for the exception.
//
// aUserInfo: A dictionary containing user-defined information relating to the exception
//
// # Return Value
// 
// The created [NSException] object or `nil` if the object couldn’t be
// created.
//
// # Discussion
// 
// This is the designated initializer.
//
// See: https://developer.apple.com/documentation/Foundation/NSException/init(name:reason:userInfo:)
func (e NSException) InitWithNameReasonUserInfo(aName NSExceptionName, aReason string, aUserInfo INSDictionary) NSException {
	rv := objc.Send[NSException](e.ID, objc.Sel("initWithName:reason:userInfo:"), objc.String(string(aName)), objc.String(aReason), aUserInfo)
	return rv
}
// Raises the receiver, causing program flow to jump to the local exception
// handler.
//
// # Discussion
// 
// When there are no exception handlers in the exception handler stack, unless
// the exception is raised during the posting of a notification, this method
// calls the uncaught exception handler, in which last-minute logging can be
// performed. The program then terminates, regardless of the actions taken by
// the uncaught exception handler.
//
// See: https://developer.apple.com/documentation/Foundation/NSException/raise()
func (e NSException) Raise() {
	objc.Send[objc.ID](e.ID, objc.Sel("raise"))
}
// Encodes the receiver using a given archiver.
//
// coder: An archiver object.
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/encode(with:)
func (e NSException) EncodeWithCoder(coder INSCoder) {
	objc.Send[objc.ID](e.ID, objc.Sel("encodeWithCoder:"), coder)
}
//
// See: https://developer.apple.com/documentation/Foundation/NSCoding/init(coder:)
func (e NSException) InitWithCoder(coder INSCoder) NSException {
	rv := objc.Send[NSException](e.ID, objc.Sel("initWithCoder:"), coder)
	return rv
}

// Creates and raises an exception with the specified name, reason, and
// arguments.
//
// name: The name of the exception.
//
// format: A human-readable message string (that is, the exception reason) with
// conversion specifications for the variable arguments in `argList`.
//
// argList: Variable information to be inserted into the formatted exception reason (in
// the manner of `vprintf`).
//
// # Discussion
// 
// The user-defined dictionary of the generated object is `nil`.
//
// See: https://developer.apple.com/documentation/Foundation/NSException/raise(_:format:arguments:)
func (_NSExceptionClass NSExceptionClass) RaiseFormatArguments(name NSExceptionName, format string, argList unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(_NSExceptionClass.class), objc.Sel("raise:format:arguments:"), objc.String(string(name)), objc.String(format), argList)
}
// Creates and returns an exception object .
//
// name: The name of the exception.
//
// reason: A human-readable message string summarizing the reason for the exception.
//
// userInfo: A dictionary containing user-defined information relating to the exception
//
// # Return Value
// 
// The created [NSException] object or `nil` if the object couldn’t be
// created.
//
// See: https://developer.apple.com/documentation/Foundation/NSException/exceptionWithName:reason:userInfo:
func (_NSExceptionClass NSExceptionClass) ExceptionWithNameReasonUserInfo(name NSExceptionName, reason string, userInfo INSDictionary) NSException {
	rv := objc.Send[objc.ID](objc.ID(_NSExceptionClass.class), objc.Sel("exceptionWithName:reason:userInfo:"), objc.String(string(name)), objc.String(reason), userInfo)
	return NSExceptionFromID(rv)
}
// A convenience method that creates and raises an exception.
//
// name: The name of the exception.
//
// format: A human-readable message string (that is, the exception reason) with
// conversion specifications for the variable arguments that follow.
//
// # Discussion
// 
// The user-defined information is `nil` for the generated exception object.
// 
// Pass variable information to be inserted into the formatted exception
// reason (in the manner of `printf`) as variadic arguments.
//
// See: https://developer.apple.com/documentation/Foundation/NSException/raise:format:
func (_NSExceptionClass NSExceptionClass) RaiseFormat(name NSExceptionName, format string) {
	objc.Send[objc.ID](objc.ID(_NSExceptionClass.class), objc.Sel("raise:format:"), objc.String(string(name)), objc.String(format))
}

// A string used to uniquely identify the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSException/name-swift.property
func (e NSException) Name() NSExceptionName {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("name"))
	return NSExceptionName(NSStringFromID(rv).String())
}
// A string containing a “human-readable” reason for the receiver.
//
// See: https://developer.apple.com/documentation/Foundation/NSException/reason-swift.property
func (e NSException) Reason() string {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("reason"))
	return NSStringFromID(rv).String()
}
// A dictionary containing application-specific data pertaining to the
// receiver.
//
// # Discussion
// 
// `nil` if no application-specific data exists. As an example, if a
// method’s return value caused the exception to be raised, the return value
// might be available to the exception handler through this method.
//
// See: https://developer.apple.com/documentation/Foundation/NSException/userInfo-swift.property
func (e NSException) UserInfo() INSDictionary {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("userInfo"))
	return NSDictionaryFromID(objc.ID(rv))
}
// The call return addresses related to a raised exception.
//
// # Discussion
// 
// An array of [NSNumber] objects encapsulating [NSUInteger] values. Each
// value is a call frame return address. The array of stack frames starts at
// the point at which the exception was first raised, with the first items
// being the most recent stack frames.
// 
// [NSException] subclasses posing as the [NSException] class or subclasses or
// other API elements that interfere with the exception-raising mechanism may
// not get this information.
//
// [NSUInteger]: https://developer.apple.com/documentation/ObjectiveC/NSUInteger
//
// See: https://developer.apple.com/documentation/Foundation/NSException/callStackReturnAddresses
func (e NSException) CallStackReturnAddresses() []NSNumber {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("callStackReturnAddresses"))
	return objc.ConvertSlice(rv, func(id objc.ID) NSNumber {
		return NSNumberFromID(id)
	})
}
// An array containing the current call stack symbols.
//
// # Discussion
// 
// An array of strings describing the call stack backtrace at the moment the
// exception was first raised. The format of each string is determined by the
// `backtrace_symbols()` API
//
// See: https://developer.apple.com/documentation/Foundation/NSException/callStackSymbols
func (e NSException) CallStackSymbols() []string {
	rv := objc.Send[[]objc.ID](e.ID, objc.Sel("callStackSymbols"))
	return objc.ConvertSliceToStrings(rv)
}

			// Protocol methods for NSCopying
			

			// Protocol methods for NSSecureCoding
			

