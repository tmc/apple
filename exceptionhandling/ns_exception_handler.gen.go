// Code generated from Apple documentation for ExceptionHandling. DO NOT EDIT.

package exceptionhandling

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSExceptionHandler] class.
var (
	_NSExceptionHandlerClass     NSExceptionHandlerClass
	_NSExceptionHandlerClassOnce sync.Once
)

func getNSExceptionHandlerClass() NSExceptionHandlerClass {
	_NSExceptionHandlerClassOnce.Do(func() {
		_NSExceptionHandlerClass = NSExceptionHandlerClass{class: objc.GetClass("NSExceptionHandler")}
	})
	return _NSExceptionHandlerClass
}

// GetNSExceptionHandlerClass returns the class object for NSExceptionHandler.
func GetNSExceptionHandlerClass() NSExceptionHandlerClass {
	return getNSExceptionHandlerClass()
}

type NSExceptionHandlerClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (nc NSExceptionHandlerClass) Class() objc.Class {
	return nc.class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSExceptionHandlerClass) Alloc() NSExceptionHandler {
	rv := objc.Send[NSExceptionHandler](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// The [NSExceptionHandler] class provides facilities for monitoring and
// debugging exceptional conditions in Objective-C programs. It works by
// installing a special uncaught exception handler via the
// [NSSetUncaughtExceptionHandler(_:)] function. Consequently, to use the
// services of [NSExceptionHandler], you must not install your own custom
// uncaught exception handler.
//
// # Overview
//
// To use these services, you set a bit mask in the singleton
// [NSExceptionHandler] instance and, optionally, a delegate. The constants
// comprising the bit mask indicate the type of exception to be monitored and
// the behavior of the [NSExceptionHandler] object (or, simply, the exception
// handler). The delegate is asked to approve the logging and handling of each
// monitored [NSException] object using the NSExceptionHandlerDelegate
// protocol.
//
// The constants for configuring exception handler behavior can be categorized
// in several ways:
//
// - Uncaught exceptions versus caught exceptions—or, more accurately,
// exceptions that would be caught (for example, by the top-level handler) -
// Exception type or cause: system exceptions (such as invalid memory
// accesses), Objective-C runtime errors (such as messages sent to freed
// objects), and other exceptions - Exception handler behavior: logging the
// exception (including a stack trace) to the console, handling the exception,
// and suspending program execution so the debugger can be attached
//
// The way the exception handler handles an exception depends on the type of
// exception; the exception handler converts system exceptions and runtime
// errors into [NSException] objects with a stack trace embedded in their
// `userInfo` dictionary; for all other uncaught exceptions, it terminates the
// thread on which they occur . The constants used to configure an
// [NSExceptionHandler] object are described in [Logging and Handling
// Constants] and [System Hang Constants].
//
// The `defaults` command-line system also allows you to set values
// corresponding to the `enum` constants used to configure the exception
// handler; see [Controlling a Program’s Response to Exceptions] for
// details.
//
// # Getting and setting exception masks
//
//   - [NSExceptionHandler.ExceptionHandlingMask]: Returns a bit mask representing the types of exceptions monitored by the receiver and its handling and logging behavior.
//   - [NSExceptionHandler.ExceptionHangingMask]: Returns a bit mask representing the types of exceptions that will halt execution for debugging.
//   - [NSExceptionHandler.SetExceptionHandlingMask]: Sets the bit mask of constants specifying the types of exceptions monitored by the receiver and its handling and logging behavior.
//   - [NSExceptionHandler.SetExceptionHangingMask]: Sets the bit mask of constants specifying the types of exceptions that will halt execution for debugging.
//
// # Getting and setting the delegate
//
//   - [NSExceptionHandler.Delegate]: Returns the delegate of the [NSExceptionHandler] object.
//   - [NSExceptionHandler.SetDelegate]: Sets the delegate of the [NSExceptionHandler] object.
//
// See: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler
//
// [NSSetUncaughtExceptionHandler(_:)]: https://developer.apple.com/documentation/Foundation/NSSetUncaughtExceptionHandler(_:)
// [Controlling a Program’s Response to Exceptions]: https://developer.apple.com/library/archive/documentation/Cocoa/Conceptual/Exceptions/Tasks/ControllingAppResponse.html#//apple_ref/doc/uid/20000473
// [Logging and Handling Constants]: https://developer.apple.com/documentation/ExceptionHandling/logging-and-handling-constants
// [NSException]: https://developer.apple.com/documentation/Foundation/NSException
// [System Hang Constants]: https://developer.apple.com/documentation/ExceptionHandling/system-hang-constants
type NSExceptionHandler struct {
	objectivec.Object
}

// NSExceptionHandlerFromID constructs a [NSExceptionHandler] from an objc.ID.
//
// The [NSExceptionHandler] class provides facilities for monitoring and
// debugging exceptional conditions in Objective-C programs. It works by
// installing a special uncaught exception handler via the
// [NSSetUncaughtExceptionHandler(_:)] function. Consequently, to use the
// services of [NSExceptionHandler], you must not install your own custom
// uncaught exception handler.
//
// [NSSetUncaughtExceptionHandler(_:)]: https://developer.apple.com/documentation/Foundation/NSSetUncaughtExceptionHandler(_:)
func NSExceptionHandlerFromID(id objc.ID) NSExceptionHandler {
	return NSExceptionHandler{objectivec.Object{ID: id}}
}

// NOTE: NSExceptionHandler adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSExceptionHandler] class.
//
// # Getting and setting exception masks
//
//   - [INSExceptionHandler.ExceptionHandlingMask]: Returns a bit mask representing the types of exceptions monitored by the receiver and its handling and logging behavior.
//   - [INSExceptionHandler.ExceptionHangingMask]: Returns a bit mask representing the types of exceptions that will halt execution for debugging.
//   - [INSExceptionHandler.SetExceptionHandlingMask]: Sets the bit mask of constants specifying the types of exceptions monitored by the receiver and its handling and logging behavior.
//   - [INSExceptionHandler.SetExceptionHangingMask]: Sets the bit mask of constants specifying the types of exceptions that will halt execution for debugging.
//
// # Getting and setting the delegate
//
//   - [INSExceptionHandler.Delegate]: Returns the delegate of the [NSExceptionHandler] object.
//   - [INSExceptionHandler.SetDelegate]: Sets the delegate of the [NSExceptionHandler] object.
//
// See: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler
type INSExceptionHandler interface {
	objectivec.IObject

	// Topic: Getting and setting exception masks

	// Returns a bit mask representing the types of exceptions monitored by the receiver and its handling and logging behavior.
	ExceptionHandlingMask() uint
	// Returns a bit mask representing the types of exceptions that will halt execution for debugging.
	ExceptionHangingMask() uint
	// Sets the bit mask of constants specifying the types of exceptions monitored by the receiver and its handling and logging behavior.
	SetExceptionHandlingMask(aMask uint)
	// Sets the bit mask of constants specifying the types of exceptions that will halt execution for debugging.
	SetExceptionHangingMask(aMask uint)

	// Topic: Getting and setting the delegate

	// Returns the delegate of the [NSExceptionHandler] object.
	Delegate() objectivec.IObject
	// Sets the delegate of the [NSExceptionHandler] object.
	SetDelegate(anObject objectivec.IObject)
}

// Init initializes the instance.
func (e NSExceptionHandler) Init() NSExceptionHandler {
	rv := objc.Send[NSExceptionHandler](e.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (e NSExceptionHandler) Autorelease() NSExceptionHandler {
	rv := objc.Send[NSExceptionHandler](e.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSExceptionHandler creates a new NSExceptionHandler instance.
func NewNSExceptionHandler() NSExceptionHandler {
	class := getNSExceptionHandlerClass()
	rv := objc.Send[NSExceptionHandler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Returns a bit mask representing the types of exceptions monitored by the
// receiver and its handling and logging behavior.
//
// # Return Value
//
// A bit mask composed of one or more constants specifying the types of
// exceptions monitored and whether they are handled or logged (or both). See
// [Logging and Handling Constants] for information about the constants.
//
// See: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/exceptionHandlingMask()
//
// [Logging and Handling Constants]: https://developer.apple.com/documentation/ExceptionHandling/logging-and-handling-constants
func (e NSExceptionHandler) ExceptionHandlingMask() uint {
	rv := objc.Send[uint](e.ID, objc.Sel("exceptionHandlingMask"))
	return rv
}

// Returns a bit mask representing the types of exceptions that will halt
// execution for debugging.
//
// # Return Value
//
// A bit mask composed of one or more constants specifying the types of
// exceptions that will halt execution for debugging. See [System Hang
// Constants] for information about the constants.
//
// See: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/exceptionHangingMask()
//
// [System Hang Constants]: https://developer.apple.com/documentation/ExceptionHandling/system-hang-constants
func (e NSExceptionHandler) ExceptionHangingMask() uint {
	rv := objc.Send[uint](e.ID, objc.Sel("exceptionHangingMask"))
	return rv
}

// Sets the bit mask of constants specifying the types of exceptions monitored
// by the receiver and its handling and logging behavior.
//
// aMask: A bit mask composed of one or more constants specifying the types of
// exceptions monitored and whether they are handled or logged (or both). You
// specify multiple constants by performing a bitwise-OR operation. See
// [Logging and Handling Constants] for information about the constants.
//
// See: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/setExceptionHandlingMask(_:)
//
// [Logging and Handling Constants]: https://developer.apple.com/documentation/ExceptionHandling/logging-and-handling-constants
func (e NSExceptionHandler) SetExceptionHandlingMask(aMask uint) {
	objc.Send[objc.ID](e.ID, objc.Sel("setExceptionHandlingMask:"), aMask)
}

// Sets the bit mask of constants specifying the types of exceptions that will
// halt execution for debugging.
//
// aMask: A bit mask composed of one or more constants specifying the types of
// exceptions that will halt execution for debugging. You specify multiple
// constants by performing a bitwise-OR operation. See [System Hang Constants]
// for information about the constants.
//
// See: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/setExceptionHangingMask(_:)
//
// [System Hang Constants]: https://developer.apple.com/documentation/ExceptionHandling/system-hang-constants
func (e NSExceptionHandler) SetExceptionHangingMask(aMask uint) {
	objc.Send[objc.ID](e.ID, objc.Sel("setExceptionHangingMask:"), aMask)
}

// Returns the delegate of the [NSExceptionHandler] object.
//
// See: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/delegate()
func (e NSExceptionHandler) Delegate() objectivec.IObject {
	rv := objc.Send[objc.ID](e.ID, objc.Sel("delegate"))
	return objectivec.Object{ID: rv}
}

// Sets the delegate of the [NSExceptionHandler] object.
//
// anObject: The object to receive the delegation messages described in
// [NSExceptionHandlerDelegate]
//
// See: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/setDelegate(_:)
//
// [NSExceptionHandlerDelegate]: https://developer.apple.com/documentation/ExceptionHandling/nsexceptionhandlerdelegate#Logging-and-handling-exceptions
func (e NSExceptionHandler) SetDelegate(anObject objectivec.IObject) {
	objc.Send[objc.ID](e.ID, objc.Sel("setDelegate:"), anObject)
}

// Returns the singleton [NSExceptionHandler] instance.
//
// See: https://developer.apple.com/documentation/ExceptionHandling/NSExceptionHandler/default()
func (_NSExceptionHandlerClass NSExceptionHandlerClass) DefaultExceptionHandler() NSExceptionHandler {
	rv := objc.Send[objc.ID](objc.ID(_NSExceptionHandlerClass.class), objc.Sel("defaultExceptionHandler"))
	return NSExceptionHandlerFromID(rv)
}
