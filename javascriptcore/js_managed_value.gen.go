// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

import (
	"sync"

	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [JSManagedValue] class.
var (
	_JSManagedValueClass     JSManagedValueClass
	_JSManagedValueClassOnce sync.Once
)

func getJSManagedValueClass() JSManagedValueClass {
	_JSManagedValueClassOnce.Do(func() {
		_JSManagedValueClass = JSManagedValueClass{class: objc.GetClass("JSManagedValue")}
	})
	return _JSManagedValueClass
}

// GetJSManagedValueClass returns the class object for JSManagedValue.
func GetJSManagedValueClass() JSManagedValueClass {
	return getJSManagedValueClass()
}

type JSManagedValueClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (jc JSManagedValueClass) Class() objc.Class {
	return jc.class
}

// Alloc allocates memory for a new instance of the class.
func (jc JSManagedValueClass) Alloc() JSManagedValue {
	rv := objc.Send[JSManagedValue](objc.ID(jc.class), objc.Sel("alloc"))
	return rv
}

// A JavaScript value with conditional retain behavior to provide automatic
// memory management.
//
// # Overview
//
// The primary use case for a managed value is to store a JavaScript value in
// an Objective-C or Swift object that exports to JavaScript.
//
// A managed value’s behavior ensures retention of its underlying JavaScript
// value as long as either of the following conditions is true:
//
// - The JavaScript value is reachable through the JavaScript object graph
// (that is, not subject to JavaScript garbage collection). - The
// [JSManagedValue] object is reachable through the Objective-C or Swift
// object graph, as you report to the JavaScriptCore virtual machine using the
// [JSVirtualMachine.AddManagedReferenceWithOwner] method.
//
// However, if neither of these conditions is true, the managed value sets its
// [JSManagedValue.Value] property to `nil`, releasing the underlying
// [JSValue] object.
//
// # Creating a Managed Value
//
//   - [JSManagedValue.InitWithValue]: Initializes a managed value with the specified JavaScript value.
//
// # Accessing the Managed Value
//
//   - [JSManagedValue.Value]: The managed value’s underlying JavaScript value.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSManagedValue
type JSManagedValue struct {
	objectivec.Object
}

// JSManagedValueFromID constructs a [JSManagedValue] from an objc.ID.
//
// A JavaScript value with conditional retain behavior to provide automatic
// memory management.
func JSManagedValueFromID(id objc.ID) JSManagedValue {
	return JSManagedValue{objectivec.Object{ID: id}}
}

// NOTE: JSManagedValue adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [JSManagedValue] class.
//
// # Creating a Managed Value
//
//   - [IJSManagedValue.InitWithValue]: Initializes a managed value with the specified JavaScript value.
//
// # Accessing the Managed Value
//
//   - [IJSManagedValue.Value]: The managed value’s underlying JavaScript value.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSManagedValue
type IJSManagedValue interface {
	objectivec.IObject

	// Topic: Creating a Managed Value

	// Initializes a managed value with the specified JavaScript value.
	InitWithValue(value IJSValue) JSManagedValue

	// Topic: Accessing the Managed Value

	// The managed value’s underlying JavaScript value.
	Value() IJSValue
}

// Init initializes the instance.
func (j JSManagedValue) Init() JSManagedValue {
	rv := objc.Send[JSManagedValue](j.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (j JSManagedValue) Autorelease() JSManagedValue {
	rv := objc.Send[JSManagedValue](j.ID, objc.Sel("autorelease"))
	return rv
}

// NewJSManagedValue creates a new JSManagedValue instance.
func NewJSManagedValue() JSManagedValue {
	class := getJSManagedValueClass()
	rv := objc.Send[JSManagedValue](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Initializes a managed value with the specified JavaScript value.
//
// value: A JavaScript value.
//
// # Return Value
//
// A new managed value.
//
// # Discussion
//
// To ensure that the underlying JavaScript value is retained as long as the
// managed value remains in use in the Objective-C or Swift runtime, report
// the managed value’s owner to the JavaScriptCore virtual machine using the
// [JSVirtualMachine.AddManagedReferenceWithOwner] method.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSManagedValue/init(value:)
func NewJSManagedValueWithValue(value IJSValue) JSManagedValue {
	instance := getJSManagedValueClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithValue:"), value)
	return JSManagedValueFromID(rv)
}

// Creates a managed value and associates it with an owner.
//
// value: A JavaScript value.
//
// owner: The Objective-C or Swift object responsible for
//
// # Return Value
//
// A new managed value.
//
// # Discussion
//
// Calling this method is equivalent to creating a managed value and then
// reporting it to the JavaScriptCore virtual machine using the
// [JSVirtualMachine.AddManagedReferenceWithOwner] method.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSManagedValue/init(value:andOwner:)
func NewJSManagedValueWithValueAndOwner(value IJSValue, owner objectivec.IObject) JSManagedValue {
	rv := objc.Send[objc.ID](objc.ID(getJSManagedValueClass().class), objc.Sel("managedValueWithValue:andOwner:"), value, owner)
	return JSManagedValueFromID(rv)
}

// Initializes a managed value with the specified JavaScript value.
//
// value: A JavaScript value.
//
// # Return Value
//
// A new managed value.
//
// # Discussion
//
// To ensure that the underlying JavaScript value is retained as long as the
// managed value remains in use in the Objective-C or Swift runtime, report
// the managed value’s owner to the JavaScriptCore virtual machine using the
// [JSVirtualMachine.AddManagedReferenceWithOwner] method.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSManagedValue/init(value:)
func (j JSManagedValue) InitWithValue(value IJSValue) JSManagedValue {
	rv := objc.Send[JSManagedValue](j.ID, objc.Sel("initWithValue:"), value)
	return rv
}

// The managed value’s underlying JavaScript value.
//
// # Discussion
//
// If the JavaScript garbage collector removes the underlying value, this
// property becomes `nil`.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSManagedValue/value
func (j JSManagedValue) Value() IJSValue {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("value"))
	return JSValueFromID(objc.ID(rv))
}
