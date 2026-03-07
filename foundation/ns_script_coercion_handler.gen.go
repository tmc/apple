// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [NSScriptCoercionHandler] class.
var (
	_NSScriptCoercionHandlerClass     NSScriptCoercionHandlerClass
	_NSScriptCoercionHandlerClassOnce sync.Once
)

func getNSScriptCoercionHandlerClass() NSScriptCoercionHandlerClass {
	_NSScriptCoercionHandlerClassOnce.Do(func() {
		_NSScriptCoercionHandlerClass = NSScriptCoercionHandlerClass{class: objc.GetClass("NSScriptCoercionHandler")}
	})
	return _NSScriptCoercionHandlerClass
}

// GetNSScriptCoercionHandlerClass returns the class object for NSScriptCoercionHandler.
func GetNSScriptCoercionHandlerClass() NSScriptCoercionHandlerClass {
	return getNSScriptCoercionHandlerClass()
}

type NSScriptCoercionHandlerClass struct {
	class objc.Class
}

// Alloc allocates memory for a new instance of the class.
func (nc NSScriptCoercionHandlerClass) Alloc() NSScriptCoercionHandler {
	rv := objc.Send[NSScriptCoercionHandler](objc.ID(nc.class), objc.Sel("alloc"))
	return rv
}







// A mechanism for converting one kind of scripting data to another.
//
// # Overview
// 
// A shared instance of this class coerces (converts) object values to objects
// of another class using information supplied by classes that register with
// it. Coercions frequently are required during key-value coding.
//
// # Working with handlers
//
//   - [NSScriptCoercionHandler.CoerceValueToClass]: Returns an object of a given class representing a given value.
//   - [NSScriptCoercionHandler.RegisterCoercerSelectorToConvertFromClassToClass]: Registers a given object (typically a class) to handle coercions (conversions) from one given class to another.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCoercionHandler
type NSScriptCoercionHandler struct {
	objectivec.Object
}

// NSScriptCoercionHandlerFromID constructs a [NSScriptCoercionHandler] from an objc.ID.
//
// A mechanism for converting one kind of scripting data to another.
func NSScriptCoercionHandlerFromID(id objc.ID) NSScriptCoercionHandler {
	return NSScriptCoercionHandler{objectivec.Object{id}}
}
// NOTE: NSScriptCoercionHandler adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.





// An interface definition for the [NSScriptCoercionHandler] class.
//
// # Working with handlers
//
//   - [INSScriptCoercionHandler.CoerceValueToClass]: Returns an object of a given class representing a given value.
//   - [INSScriptCoercionHandler.RegisterCoercerSelectorToConvertFromClassToClass]: Registers a given object (typically a class) to handle coercions (conversions) from one given class to another.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCoercionHandler
type INSScriptCoercionHandler interface {
	objectivec.IObject

	// Topic: Working with handlers

	// Returns an object of a given class representing a given value.
	CoerceValueToClass(value objectivec.IObject, toClass objc.Class) objectivec.IObject
	// Registers a given object (typically a class) to handle coercions (conversions) from one given class to another.
	RegisterCoercerSelectorToConvertFromClassToClass(coercer objectivec.IObject, selector objc.SEL, fromClass objc.Class, toClass objc.Class)
}





// Init initializes the instance.
func (s NSScriptCoercionHandler) Init() NSScriptCoercionHandler {
	rv := objc.Send[NSScriptCoercionHandler](s.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s NSScriptCoercionHandler) Autorelease() NSScriptCoercionHandler {
	rv := objc.Send[NSScriptCoercionHandler](s.ID, objc.Sel("autorelease"))
	return rv
}

// NewNSScriptCoercionHandler creates a new NSScriptCoercionHandler instance.
func NewNSScriptCoercionHandler() NSScriptCoercionHandler {
	class := getNSScriptCoercionHandlerClass()
	rv := objc.Send[NSScriptCoercionHandler](objc.ID(class.class), objc.Sel("new"))
	return rv
}










// Returns an object of a given class representing a given value.
//
// value: The value to coerce.
//
// toClass: The class with which to represent `value`.
//
// # Return Value
// 
// An object of the class `toClass` representing the value specified by
// `value`. Returns `nil` if an error occurs.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCoercionHandler/coerceValue(_:to:)
func (s NSScriptCoercionHandler) CoerceValueToClass(value objectivec.IObject, toClass objc.Class) objectivec.IObject {
	rv := objc.Send[objc.ID](s.ID, objc.Sel("coerceValue:toClass:"), value, toClass)
	return objectivec.Object{ID: rv}
}

// Registers a given object (typically a class) to handle coercions
// (conversions) from one given class to another.
//
// coercer: The object that performs the coercion. `coercer` should typically be a
// class object.
//
// selector: A selector that specifies the method to perform the coercion. `selector`
// should typically be a factory method, and must take two arguments. The
// first is the value to be converted. The second is the class to convert it
// to.
//
// fromClass: The class for which instances are coerced.
//
// toClass: The class to which instances of `fromClass` are coerced.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCoercionHandler/registerCoercer(_:selector:toConvertFrom:to:)
func (s NSScriptCoercionHandler) RegisterCoercerSelectorToConvertFromClassToClass(coercer objectivec.IObject, selector objc.SEL, fromClass objc.Class, toClass objc.Class) {
	objc.Send[objc.ID](s.ID, objc.Sel("registerCoercer:selector:toConvertFromClass:toClass:"), coercer, selector, fromClass, toClass)
}





// Returns the shared [NSScriptCoercionHandler] for the application.
//
// # Return Value
// 
// The shared [NSScriptCoercionHandler] for the application.
//
// See: https://developer.apple.com/documentation/Foundation/NSScriptCoercionHandler/shared()
func (_NSScriptCoercionHandlerClass NSScriptCoercionHandlerClass) SharedCoercionHandler() NSScriptCoercionHandler {
	rv := objc.Send[objc.ID](objc.ID(_NSScriptCoercionHandlerClass.class), objc.Sel("sharedCoercionHandler"))
	return NSScriptCoercionHandlerFromID(rv)
}




























