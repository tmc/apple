// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSAssertionHandler] class.
var (
	_NSAssertionHandlerClass     NSAssertionHandlerClass
	_NSAssertionHandlerClassOnce sync.Once
)

func getNSAssertionHandlerClass() NSAssertionHandlerClass {
	_NSAssertionHandlerClassOnce.Do(func() {
		_NSAssertionHandlerClass = NSAssertionHandlerClass{class: objc.GetClass("NSAssertionHandler")}
	})
	return _NSAssertionHandlerClass
}

// GetNSAssertionHandlerClass returns the class object for NSAssertionHandler.
func GetNSAssertionHandlerClass() NSAssertionHandlerClass {
	return getNSAssertionHandlerClass()
}

type NSAssertionHandlerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSAssertionHandlerClass) Alloc() NSAssertionHandler {
	rv := objc.Send[NSAssertionHandler](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}

// An object that logs an assertion to the console.
//
// # Overview
// 
// [NSAssertionHandler] objects are automatically created to handle false
// assertions. Assertion macros, such as [NSAssert] and [NSCAssert], are used
// to evaluate a condition, and if the condition evaluates to false, the
// macros pass a string to an [NSAssertionHandler] object describing the
// failure. Each thread has its own [NSAssertionHandler] object. When invoked,
// an assertion handler prints an error message that includes the method and
// class (or function) containing the assertion and raises an
// [NSInternalInconsistencyException].
// 
// You create assertions only using the assertion macros—you rarely need to
// invoke [NSAssertionHandler] methods directly. The macros for use inside
// methods and functions send
// [NSAssertionHandler.HandleFailureInMethodObjectFileLineNumberDescription] and
// [NSAssertionHandler.HandleFailureInFunctionFileLineNumberDescription] messages respectively to
// the current assertion handler. The assertion handler for the current thread
// is obtained using the [NSAssertionHandler.CurrentHandler] class method. See
// doc:nsassertionhandlerkey if you need to customize the behavior of
// [NSAssertionHandler].
//
// See: https://developer.apple.com/documentation/Foundation/NSAssertionHandler
type NSAssertionHandler struct {
	objectivec.Object
}

// NSAssertionHandlerFromID constructs a [NSAssertionHandler] from an objc.ID.
//
// An object that logs an assertion to the console.
func NSAssertionHandlerFromID(id objc.ID) NSAssertionHandler {
	return NSAssertionHandler{objectivec.Object{ID: id}}
}
// NOTE: NSAssertionHandler adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [NSAssertionHandler] class.
//
// See: https://developer.apple.com/documentation/Foundation/NSAssertionHandler
type INSAssertionHandler interface {
	objectivec.IObject

	HandleFailureInFunctionFileLineNumberDescription(functionName string, fileName string, line int, format string)
	HandleFailureInMethodObjectFileLineNumberDescription(selector objc.SEL, object objectivec.IObject, fileName string, line int, format string)
}

// Init initializes the instance.
func (a NSAssertionHandler) Init() NSAssertionHandler {
	rv := objc.Send[NSAssertionHandler](a.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a NSAssertionHandler) Autorelease() NSAssertionHandler {
	rv := objc.Send[NSAssertionHandler](a.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSAssertionHandler creates a new NSAssertionHandler instance.
func NewNSAssertionHandler() NSAssertionHandler {
	class := getNSAssertionHandlerClass()
	rv := objc.Send[NSAssertionHandler](objc.ID(class.class), objc.Sel("new"))
	return rv
}

//
// See: https://developer.apple.com/documentation/Foundation/NSAssertionHandler/handleFailureInFunction:file:lineNumber:description:
func (a NSAssertionHandler) HandleFailureInFunctionFileLineNumberDescription(functionName string, fileName string, line int, format string) {
	objc.Send[objc.ID](a.ID, objc.Sel("handleFailureInFunction:file:lineNumber:description:"), objc.String(functionName), objc.String(fileName), line, objc.String(format))
}

//
// See: https://developer.apple.com/documentation/Foundation/NSAssertionHandler/handleFailureInMethod:object:file:lineNumber:description:
func (a NSAssertionHandler) HandleFailureInMethodObjectFileLineNumberDescription(selector objc.SEL, object objectivec.IObject, fileName string, line int, format string) {
	objc.Send[objc.ID](a.ID, objc.Sel("handleFailureInMethod:object:file:lineNumber:description:"), selector, object, objc.String(fileName), line, objc.String(format))
}

// Returns the [NSAssertionHandler] object associated with the current thread.
//
// # Return Value
// 
// The [NSAssertionHandler] object associated with the current thread.
// 
// # Discussion
// 
// If no assertion handler is associated with the current thread, this method
// creates one and assigns it to the thread.
//
// See: https://developer.apple.com/documentation/Foundation/NSAssertionHandler/current
func (_NSAssertionHandlerClass NSAssertionHandlerClass) CurrentHandler() NSAssertionHandler {
	rv := objc.Send[objc.ID](objc.ID(_NSAssertionHandlerClass.class), objc.Sel("currentHandler"))
	return NSAssertionHandlerFromID(objc.ID(rv))
}

