// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

import (
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [JSValue] class.
var (
	_JSValueClass     JSValueClass
	_JSValueClassOnce sync.Once
)

func getJSValueClass() JSValueClass {
	_JSValueClassOnce.Do(func() {
		_JSValueClass = JSValueClass{class: objc.GetClass("JSValue")}
	})
	return _JSValueClass
}

// GetJSValueClass returns the class object for JSValue.
func GetJSValueClass() JSValueClass {
	return getJSValueClass()
}

type JSValueClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (jc JSValueClass) Class() objc.Class {
	return jc.class
}

// Alloc allocates memory for a new instance of the class.
func (jc JSValueClass) Alloc() JSValue {
	rv := objc.Send[JSValue](objc.ID(jc.class), objc.Sel("alloc"))
	return rv
}

// A JavaScript value.
//
// # Overview
//
// You use the [JSValue] class to convert basic values, such as numbers and
// strings, between JavaScript and Objective-C or Swift representations to
// pass data between native code and JavaScript code. You can also use this
// class to create JavaScript objects that wrap native objects of custom
// classes or JavaScript functions with implementations that native methods or
// blocks provide.
//
// Each [JSValue] instance originates from a [JSContext] object that
// represents the JavaScript execution environment containing that value. The
// value holds a strong reference to its [JSValue.Context] object — as long
// as it retains any value for a particular [JSContext] instance, that context
// remains alive. When you invoke an instance method on a [JSValue] object,
// and that method returns another [JSValue] object, the returned value
// belongs to the same context as the original value.
//
// Each JavaScript value also has an association (indirectly via the
// [JSValue.Context] property) with a specific [JSVirtualMachine] object that
// represents the underlying set of execution resources for its context. You
// can pass [JSValue] instances only to methods on [JSValue] and [JSContext]
// instances on the same virtual machine — attempting to pass a value to a
// different virtual machine raises an Objective-C exception.
//
// # Convert Between JavaScript and Native Types
//
// When you use the [JSValue] methods for creating, reading, and converting
// JavaScript values, JavaScriptCore automatically converts native values to
// JavaScript values and vice versa, using the rules below.
//
// - [NSDictionary] objects or Swift dictionaries and the keys they contain
// become JavaScript objects with matching named properties and vice versa.
// JavaScriptCore recursively copies and converts the values for keys. -
// [NSArray] objects or Swift arrays become JavaScript arrays and vice versa,
// with elements that JavaScriptCore recursively copies and converts. -
// Objective-C blocks (or Swift closures with the `@convention(block)`
// attribute) become JavaScript [Function] objects, with parameter and return
// types that JavaScriptCore converts using the same rules as values.
// Converting a JavaScript function with a backing from a native block or
// method returns that block or method; all other JavaScript functions convert
// as empty dictionaries. - For all other native object types (and class types
// or metatypes), JavaScriptCore creates a JavaScript wrapper object with a
// constructor prototype chain that reflects the native class hierarchy. By
// default, the JavaScript wrapper for a native object doesn’t make that
// object’s properties and methods available in JavaScript. To choose
// properties and methods for export to JavaScript, see [JSExport].
//
// When you convert an object, method, or block, JavaScriptCore implicitly
// converts the types and values of object properties and method parameters
// using the rules below:
//
// [Table data omitted]
//
// # Reading and Converting JavaScript Values
//
//   - [JSValue.ToObject]: Converts the JavaScript value to a native object.
//   - [JSValue.ToObjectOfClass]: Converts the JavaScript value to a native object of the specified class.
//   - [JSValue.ToBool]: Converts the JavaScript value to a native Boolean value.
//   - [JSValue.ToDouble]: Converts the JavaScript value to a native floating-point value.
//   - [JSValue.ToInt32]: Converts the JavaScript value to a native signed integer value.
//   - [JSValue.ToUInt32]: Converts the JavaScript value to a native unsigned integer value.
//   - [JSValue.ToNumber]: Converts the JavaScript value to a [NSNumber](<https://developer.apple.com/documentation/Foundation/NSNumber>) object.
//   - [JSValue.ToString]: Converts the JavaScript value to a native string.
//   - [JSValue.ToDate]: Converts the JavaScript value to a date object.
//   - [JSValue.ToArray]: Converts the JavaScript value to an array.
//   - [JSValue.ToDictionary]: Converts the JavaScript value to a dictionary.
//   - [JSValue.ToPoint]: Converts the value to a point structure.
//   - [JSValue.ToRange]: Converts the value to a range.
//   - [JSValue.ToRect]: Converts the value to a rectangle structure.
//   - [JSValue.ToSize]: Converts the value to a size.
//
// # Determining the Type of a JavaScript Value
//
//   - [JSValue.IsUndefined]: A Boolean value that indicates whether the instance corresponds to the JavaScript `undefined` value.
//   - [JSValue.IsNull]: A Boolean value that indicates whether the instance corresponds to the JavaScript `null` value.
//   - [JSValue.IsBoolean]: A Boolean value that indicates whether the instance is a JavaScript Boolean value.
//   - [JSValue.IsNumber]: A Boolean value that indicates whether the instance is a JavaScript numeric value.
//   - [JSValue.IsString]: A Boolean value that indicates whether the instance is a JavaScript [String] object.
//   - [JSValue.IsObject]: A Boolean value that indicates whether the instance is a JavaScript object.
//   - [JSValue.IsArray]: A Boolean value that indicates whether the instance is a JavaScript array value.
//   - [JSValue.IsDate]: A Boolean value that indicates whether the instance is a JavaScript [Date] object.
//   - [JSValue.IsSymbol]: A Boolean value that indicates whether the instance is a symbol.
//
// # Comparing JavaScript Values
//
//   - [JSValue.IsEqualToObject]: Compares the value to another for strict equality.
//   - [JSValue.IsEqualWithTypeCoercionToObject]: Compares the value to another for equivalence, allowing type conversion.
//   - [JSValue.IsInstanceOf]: Returns a Boolean value indicating whether the value is an instance of another JavaScript object value.
//
// # Working with Function and Constructor Values
//
//   - [JSValue.CallWithArguments]: Invokes the value as a JavaScript function.
//   - [JSValue.ConstructWithArguments]: Invokes the value as a JavaScript constructor.
//   - [JSValue.InvokeMethodWithArguments]: Calls the named JavaScript method on the value.
//
// # Working with Container Values
//
//   - [JSValue.DefinePropertyDescriptor]: Defines a property on the JavaScript object value or modifies a property’s definition.
//   - [JSValue.HasProperty]: Returns a Boolean value indicating whether the JavaScript value has a defined property with the specified name.
//   - [JSValue.DeleteProperty]: Deletes the named property from the JavaScript object value.
//   - [JSValue.ValueAtIndex]: Returns the value at the specified numeric index in the JavaScript object value.
//   - [JSValue.SetValueAtIndex]: Sets the value at the specified numeric index in the JavaScript object value.
//   - [JSValue.ValueForProperty]: Returns the value of the named property in the JavaScript object value.
//   - [JSValue.SetValueForProperty]: Sets the value of the named property in the JavaScript object value.
//
// # Accessing a Value’s JavaScript Context
//
//   - [JSValue.Context]: The JavaScript context hosting this value.
//
// # Accessing Values with Subscript Syntax
//
//   - [JSValue.ObjectAtIndexedSubscript]: Returns the value’s JavaScript property at the specified index, allowing subscript syntax.
//   - [JSValue.SetObjectAtIndexedSubscript]: Sets the value’s JavaScript property at the specified index, allowing subscript syntax.
//   - [JSValue.ObjectForKeyedSubscript]: Returns the value’s JavaScript property named with the specified key, allowing subscript syntax.
//   - [JSValue.SetObjectForKeyedSubscript]: Sets the value’s JavaScript property named with the specified key, allowing subscript syntax.
//
// # Working with the C JavaScriptCore API
//
//   - [JSValue.JSValueRef]: Returns the C representation of the JavaScript value.
//
// # Instance Properties
//
//   - [JSValue.IsBigInt]
//
// # Instance Methods
//
//   - [JSValue.CompareDouble]
//   - [JSValue.CompareJSValue]
//   - [JSValue.CompareUInt64]
//   - [JSValue.CompareInt64]
//   - [JSValue.ToInt64]
//   - [JSValue.ToUInt64]
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue
//
// [NSArray]: https://developer.apple.com/documentation/Foundation/NSArray
// [NSDictionary]: https://developer.apple.com/documentation/Foundation/NSDictionary
type JSValue struct {
	objectivec.Object
}

// JSValueFromID constructs a [JSValue] from an objc.ID.
//
// A JavaScript value.
func JSValueFromID(id objc.ID) JSValue {
	return JSValue{objectivec.Object{ID: id}}
}

// NOTE: JSValue adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [JSValue] class.
//
// # Reading and Converting JavaScript Values
//
//   - [IJSValue.ToObject]: Converts the JavaScript value to a native object.
//   - [IJSValue.ToObjectOfClass]: Converts the JavaScript value to a native object of the specified class.
//   - [IJSValue.ToBool]: Converts the JavaScript value to a native Boolean value.
//   - [IJSValue.ToDouble]: Converts the JavaScript value to a native floating-point value.
//   - [IJSValue.ToInt32]: Converts the JavaScript value to a native signed integer value.
//   - [IJSValue.ToUInt32]: Converts the JavaScript value to a native unsigned integer value.
//   - [IJSValue.ToNumber]: Converts the JavaScript value to a [NSNumber](<https://developer.apple.com/documentation/Foundation/NSNumber>) object.
//   - [IJSValue.ToString]: Converts the JavaScript value to a native string.
//   - [IJSValue.ToDate]: Converts the JavaScript value to a date object.
//   - [IJSValue.ToArray]: Converts the JavaScript value to an array.
//   - [IJSValue.ToDictionary]: Converts the JavaScript value to a dictionary.
//   - [IJSValue.ToPoint]: Converts the value to a point structure.
//   - [IJSValue.ToRange]: Converts the value to a range.
//   - [IJSValue.ToRect]: Converts the value to a rectangle structure.
//   - [IJSValue.ToSize]: Converts the value to a size.
//
// # Determining the Type of a JavaScript Value
//
//   - [IJSValue.IsUndefined]: A Boolean value that indicates whether the instance corresponds to the JavaScript `undefined` value.
//   - [IJSValue.IsNull]: A Boolean value that indicates whether the instance corresponds to the JavaScript `null` value.
//   - [IJSValue.IsBoolean]: A Boolean value that indicates whether the instance is a JavaScript Boolean value.
//   - [IJSValue.IsNumber]: A Boolean value that indicates whether the instance is a JavaScript numeric value.
//   - [IJSValue.IsString]: A Boolean value that indicates whether the instance is a JavaScript [String] object.
//   - [IJSValue.IsObject]: A Boolean value that indicates whether the instance is a JavaScript object.
//   - [IJSValue.IsArray]: A Boolean value that indicates whether the instance is a JavaScript array value.
//   - [IJSValue.IsDate]: A Boolean value that indicates whether the instance is a JavaScript [Date] object.
//   - [IJSValue.IsSymbol]: A Boolean value that indicates whether the instance is a symbol.
//
// # Comparing JavaScript Values
//
//   - [IJSValue.IsEqualToObject]: Compares the value to another for strict equality.
//   - [IJSValue.IsEqualWithTypeCoercionToObject]: Compares the value to another for equivalence, allowing type conversion.
//   - [IJSValue.IsInstanceOf]: Returns a Boolean value indicating whether the value is an instance of another JavaScript object value.
//
// # Working with Function and Constructor Values
//
//   - [IJSValue.CallWithArguments]: Invokes the value as a JavaScript function.
//   - [IJSValue.ConstructWithArguments]: Invokes the value as a JavaScript constructor.
//   - [IJSValue.InvokeMethodWithArguments]: Calls the named JavaScript method on the value.
//
// # Working with Container Values
//
//   - [IJSValue.DefinePropertyDescriptor]: Defines a property on the JavaScript object value or modifies a property’s definition.
//   - [IJSValue.HasProperty]: Returns a Boolean value indicating whether the JavaScript value has a defined property with the specified name.
//   - [IJSValue.DeleteProperty]: Deletes the named property from the JavaScript object value.
//   - [IJSValue.ValueAtIndex]: Returns the value at the specified numeric index in the JavaScript object value.
//   - [IJSValue.SetValueAtIndex]: Sets the value at the specified numeric index in the JavaScript object value.
//   - [IJSValue.ValueForProperty]: Returns the value of the named property in the JavaScript object value.
//   - [IJSValue.SetValueForProperty]: Sets the value of the named property in the JavaScript object value.
//
// # Accessing a Value’s JavaScript Context
//
//   - [IJSValue.Context]: The JavaScript context hosting this value.
//
// # Accessing Values with Subscript Syntax
//
//   - [IJSValue.ObjectAtIndexedSubscript]: Returns the value’s JavaScript property at the specified index, allowing subscript syntax.
//   - [IJSValue.SetObjectAtIndexedSubscript]: Sets the value’s JavaScript property at the specified index, allowing subscript syntax.
//   - [IJSValue.ObjectForKeyedSubscript]: Returns the value’s JavaScript property named with the specified key, allowing subscript syntax.
//   - [IJSValue.SetObjectForKeyedSubscript]: Sets the value’s JavaScript property named with the specified key, allowing subscript syntax.
//
// # Working with the C JavaScriptCore API
//
//   - [IJSValue.JSValueRef]: Returns the C representation of the JavaScript value.
//
// # Instance Properties
//
//   - [IJSValue.IsBigInt]
//
// # Instance Methods
//
//   - [IJSValue.CompareDouble]
//   - [IJSValue.CompareJSValue]
//   - [IJSValue.CompareUInt64]
//   - [IJSValue.CompareInt64]
//   - [IJSValue.ToInt64]
//   - [IJSValue.ToUInt64]
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue
type IJSValue interface {
	objectivec.IObject

	// Topic: Reading and Converting JavaScript Values

	// Converts the JavaScript value to a native object.
	ToObject() objectivec.IObject
	// Converts the JavaScript value to a native object of the specified class.
	ToObjectOfClass(expectedClass objectivec.Class) objectivec.IObject
	// Converts the JavaScript value to a native Boolean value.
	ToBool() bool
	// Converts the JavaScript value to a native floating-point value.
	ToDouble() float64
	// Converts the JavaScript value to a native signed integer value.
	ToInt32() int32
	// Converts the JavaScript value to a native unsigned integer value.
	ToUInt32() uint32
	// Converts the JavaScript value to a [NSNumber](<https://developer.apple.com/documentation/Foundation/NSNumber>) object.
	ToNumber() foundation.NSNumber
	// Converts the JavaScript value to a native string.
	ToString() string
	// Converts the JavaScript value to a date object.
	ToDate() foundation.NSDate
	// Converts the JavaScript value to an array.
	ToArray() foundation.INSArray
	// Converts the JavaScript value to a dictionary.
	ToDictionary() foundation.INSDictionary
	// Converts the value to a point structure.
	ToPoint() corefoundation.CGPoint
	// Converts the value to a range.
	ToRange() foundation.NSRange
	// Converts the value to a rectangle structure.
	ToRect() corefoundation.CGRect
	// Converts the value to a size.
	ToSize() corefoundation.CGSize

	// Topic: Determining the Type of a JavaScript Value

	// A Boolean value that indicates whether the instance corresponds to the JavaScript `undefined` value.
	IsUndefined() bool
	// A Boolean value that indicates whether the instance corresponds to the JavaScript `null` value.
	IsNull() bool
	// A Boolean value that indicates whether the instance is a JavaScript Boolean value.
	IsBoolean() bool
	// A Boolean value that indicates whether the instance is a JavaScript numeric value.
	IsNumber() bool
	// A Boolean value that indicates whether the instance is a JavaScript [String] object.
	IsString() bool
	// A Boolean value that indicates whether the instance is a JavaScript object.
	IsObject() bool
	// A Boolean value that indicates whether the instance is a JavaScript array value.
	IsArray() bool
	// A Boolean value that indicates whether the instance is a JavaScript [Date] object.
	IsDate() bool
	// A Boolean value that indicates whether the instance is a symbol.
	IsSymbol() bool

	// Topic: Comparing JavaScript Values

	// Compares the value to another for strict equality.
	IsEqualToObject(value objectivec.IObject) bool
	// Compares the value to another for equivalence, allowing type conversion.
	IsEqualWithTypeCoercionToObject(value objectivec.IObject) bool
	// Returns a Boolean value indicating whether the value is an instance of another JavaScript object value.
	IsInstanceOf(value objectivec.IObject) bool

	// Topic: Working with Function and Constructor Values

	// Invokes the value as a JavaScript function.
	CallWithArguments(arguments foundation.INSArray) IJSValue
	// Invokes the value as a JavaScript constructor.
	ConstructWithArguments(arguments foundation.INSArray) IJSValue
	// Calls the named JavaScript method on the value.
	InvokeMethodWithArguments(method string, arguments foundation.INSArray) IJSValue

	// Topic: Working with Container Values

	// Defines a property on the JavaScript object value or modifies a property’s definition.
	DefinePropertyDescriptor(property JSValueProperty, descriptor objectivec.IObject)
	// Returns a Boolean value indicating whether the JavaScript value has a defined property with the specified name.
	HasProperty(property JSValueProperty) bool
	// Deletes the named property from the JavaScript object value.
	DeleteProperty(property JSValueProperty) bool
	// Returns the value at the specified numeric index in the JavaScript object value.
	ValueAtIndex(index uint) IJSValue
	// Sets the value at the specified numeric index in the JavaScript object value.
	SetValueAtIndex(value objectivec.IObject, index uint)
	// Returns the value of the named property in the JavaScript object value.
	ValueForProperty(property JSValueProperty) IJSValue
	// Sets the value of the named property in the JavaScript object value.
	SetValueForProperty(value objectivec.IObject, property JSValueProperty)

	// Topic: Accessing a Value’s JavaScript Context

	// The JavaScript context hosting this value.
	Context() IJSContext

	// Topic: Accessing Values with Subscript Syntax

	// Returns the value’s JavaScript property at the specified index, allowing subscript syntax.
	ObjectAtIndexedSubscript(index uint) IJSValue
	// Sets the value’s JavaScript property at the specified index, allowing subscript syntax.
	SetObjectAtIndexedSubscript(object objectivec.IObject, index uint)
	// Returns the value’s JavaScript property named with the specified key, allowing subscript syntax.
	ObjectForKeyedSubscript(key objectivec.IObject) IJSValue
	// Sets the value’s JavaScript property named with the specified key, allowing subscript syntax.
	SetObjectForKeyedSubscript(object objectivec.IObject, key objectivec.IObject)

	// Topic: Working with the C JavaScriptCore API

	// Returns the C representation of the JavaScript value.
	JSValueRef() JSValueRef

	// Topic: Instance Properties

	IsBigInt() bool

	// Topic: Instance Methods

	CompareDouble(other float64) JSRelationCondition
	CompareJSValue(other IJSValue) JSRelationCondition
	CompareUInt64(other uint64) JSRelationCondition
	CompareInt64(other int64) JSRelationCondition
	ToInt64() int64
	ToUInt64() uint64
}

// Init initializes the instance.
func (j JSValue) Init() JSValue {
	rv := objc.Send[JSValue](j.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (j JSValue) Autorelease() JSValue {
	rv := objc.Send[JSValue](j.ID, objc.Sel("autorelease"))
	return rv
}

// NewJSValue creates a new JSValue instance.
func NewJSValue() JSValue {
	class := getJSValueClass()
	rv := objc.Send[JSValue](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a JavaScript representation of the specified Boolean value.
//
// value: A native Boolean value.
//
// context: The JavaScript context in which to create the value.
//
// # Return Value
//
// A JavaScript Boolean value.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(bool:in:)
func NewJSValueWithBoolInContext(value bool, context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithBool:inContext:"), value, context)
	return JSValueFromID(rv)
}

// Creates a JavaScript representation of the specified floating-point value.
//
// value: A native double-precision floating-point value.
//
// context: The JavaScript context in which to create the value.
//
// # Return Value
//
// A JavaScript numeric value.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(double:in:)
func NewJSValueWithDoubleInContext(value float64, context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithDouble:inContext:"), value, context)
	return JSValueFromID(rv)
}

// Creates a JavaScript representation of the specified signed integer value.
//
// value: A native 32-bit signed integer value.
//
// context: The JavaScript context in which to create the value.
//
// # Return Value
//
// A JavaScript numeric value.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(int32:in:)
func NewJSValueWithInt32InContext(value int32, context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithInt32:inContext:"), value, context)
	return JSValueFromID(rv)
}

// Creates a JavaScript value object from the equivalent C representation.
//
// value: A C JavaScript value reference.
//
// context: The JavaScript context in which to create the value.
//
// # Return Value
//
// A JavaScript value object representing the same value.
//
// # Discussion
//
// See [JSValueRef] for the C JavaScriptCore API.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(JSValueRef:inContext:)
func NewJSValueWithJSValueRefInContext(value JSValueRef, context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithJSValueRef:inContext:"), value, context)
	return JSValueFromID(rv)
}

// Creates a new, empty JavaScript array value.
//
// context: The JavaScript context in which to create the value.
//
// # Return Value
//
// An empty JavaScript array value.
//
// # Discussion
//
// Calling this method is equivalent to declaring an empty array literal `[]`
// or using the `new Array()` syntax in JavaScript.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newArrayIn:)
func NewJSValueWithNewArrayInContext(context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewArrayInContext:"), context)
	return JSValueFromID(rv)
}

// value: The value of the BigInt JavaScript value being created.
//
// context: The JSContext to which the resulting JSValue belongs.
//
// # Return Value
//
// The JSValue representing a JavaScript value with type BigInt.
//
// # Discussion
//
// Create a new BigInt value from a double.
//
// If the value is not an integer, an exception is thrown.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newBigIntFrom:in:)-r38z
func NewJSValueWithNewBigIntFromDoubleInContext(value float64, context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewBigIntFromDouble:inContext:"), value, context)
	return JSValueFromID(rv)
}

// int64: The signed 64-bit integer of the BigInt JavaScript value being created.
//
// context: The JSContext to which the resulting JSValue belongs.
//
// # Return Value
//
// The JSValue representing a JavaScript value with type BigInt.
//
// # Discussion
//
// Create a new BigInt value from a int64_t.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newBigIntFrom:in:)-8l9iv
func NewJSValueWithNewBigIntFromInt64InContext(int64_ int64, context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewBigIntFromInt64:inContext:"), int64_, context)
	return JSValueFromID(rv)
}

// string: The string representation of the BigInt JavaScript value being created.
//
// context: The JSContext to which the resulting JSValue belongs.
//
// # Return Value
//
// The JSValue representing a JavaScript value with type BigInt.
//
// # Discussion
//
// Create a new BigInt value from a numeric string.
//
// This is equivalent to calling the BigInt constructor from JavaScript with a
// string argument.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newBigIntFrom:in:)-1f0xs
func NewJSValueWithNewBigIntFromStringInContext(string_ string, context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewBigIntFromString:inContext:"), objc.String(string_), context)
	return JSValueFromID(rv)
}

// uint64: The unsigned 64-bit integer of the BigInt JavaScript value being created.
//
// context: The JSContext to which the resulting JSValue belongs.
//
// # Return Value
//
// The JSValue representing a JavaScript value with type BigInt.
//
// # Discussion
//
// Create a new BigInt value from a uint64_t.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newBigIntFrom:in:)-7worq
func NewJSValueWithNewBigIntFromUInt64InContext(uint64_ uint64, context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewBigIntFromUInt64:inContext:"), uint64_, context)
	return JSValueFromID(rv)
}

// Creates a JavaScript error value with the specified error message.
//
// message: The error message for the error object.
//
// context: The JavaScript context in which to create the value.
//
// # Return Value
//
// A new JavaScript error value.
//
// # Discussion
//
// Calling this method creates a JavaScript [Error] object, and is equivalent
// to calling the [Error] constructor (for example, `new Error("message")`) in
// JavaScript.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newErrorFromMessage:in:)
func NewJSValueWithNewErrorFromMessageInContext(message string, context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewErrorFromMessage:inContext:"), objc.String(message), context)
	return JSValueFromID(rv)
}

// Creates a new, empty JavaScript object value.
//
// context: The JavaScript context in which to create the value.
//
// # Return Value
//
// An empty JavaScript object value.
//
// # Discussion
//
// Calling this method is equivalent to declaring an empty object literal `{}`
// or using the `new Object()` syntax in JavaScript.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newObjectIn:)
func NewJSValueWithNewObjectInContext(context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewObjectInContext:"), context)
	return JSValueFromID(rv)
}

// Creates a rejected promise object with the specified value.
//
// reason: The result value to pass to any reactions.
//
// context: The [JSContext] the resulting [JSValue] belongs to.
//
// # Return Value
//
// A [JSValue] that represents a new promise JavaScript object.
//
// # Discussion
//
// This method is equivalent to calling the following:
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newPromiseRejectedWithReason:in:)
func NewJSValueWithNewPromiseRejectedWithReasonInContext(reason objectivec.IObject, context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewPromiseRejectedWithReason:inContext:"), reason, context)
	return JSValueFromID(rv)
}

// Creates a resolved promise object with the specified value.
//
// result: The result value to pass to any reactions.
//
// context: The [JSContext] the resulting [JSValue] belongs to.
//
// # Return Value
//
// A [JSValue] that represents a new promise JavaScript object.
//
// # Discussion
//
// This method is equivalent to calling the following:
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newPromiseResolvedWithResult:in:)
func NewJSValueWithNewPromiseResolvedWithResultInContext(result objectivec.IObject, context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewPromiseResolvedWithResult:inContext:"), result, context)
	return JSValueFromID(rv)
}

// Creates a JavaScript regular expression value from the specified pattern.
//
// pattern: A string to be interpreted as a regular expression pattern.
//
// flags: A combination of zero or more single-letter flags specifying search
// options.
//
// context: The JavaScript context in which to create the value.
//
// # Return Value
//
// A new JavaScript regular expression object.
//
// # Discussion
//
// Calling this method creates a JavaScript [RegExp] object, and is equivalent
// to declaring a regular expression literal (such as `/ab+c/i`) or calling
// the [RegExp] constructor (for example, `new RegExp("ab+c", "i")`) in
// JavaScript.
//
// The `flags` parameter can include any of the following options:
//
// - `g` (global match): match all occurrences of the pattern in a string, not
// just the first. - `i` (ignore case): perform case-insensitive search. - `m`
// (multiline): treat the `^` and `$` regular expression tokens as matching
// the start or end of any line in a string (as delimited by newline or return
// characters), not just the start or end of the entire string.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newRegularExpressionFromPattern:flags:in:)
func NewJSValueWithNewRegularExpressionFromPatternFlagsInContext(pattern string, flags string, context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewRegularExpressionFromPattern:flags:inContext:"), objc.String(pattern), objc.String(flags), context)
	return JSValueFromID(rv)
}

// Creates a unique symbol object.
//
// description: The description of the symbol object to create.
//
// context: The [JSContext] the resulting [JSValue] belongs to.
//
// # Return Value
//
// A [JSValue] that represents a new, unique JavaScript symbol object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newSymbolFromDescription:in:)
func NewJSValueWithNewSymbolFromDescriptionInContext(description string, context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithNewSymbolFromDescription:inContext:"), objc.String(description), context)
	return JSValueFromID(rv)
}

// Creates a JavaScript `null` value.
//
// context: The JavaScript context in which to create the value.
//
// # Return Value
//
// The JavaScript `null` value.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(nullIn:)
func NewJSValueWithNullInContext(context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithNullInContext:"), context)
	return JSValueFromID(rv)
}

// Creates a JavaScript value by converting the specified native object.
//
// value: The Objective-C or Swift object to be made available to JavaScript.
//
// context: The JavaScript context in which to create the value.
//
// # Return Value
//
// A new JavaScript value representing the object.
//
// # Discussion
//
// Converting a native object creates a JavaScript object, including a
// constructor and prototype chain that reflects the object’s inheritance in
// the Objective-C or Swift type hierarchy. By default, properties and methods
// on the converted object are not exposed to JavaScript: to choose which
// properties and methods should be visible to JavaScript, see [JSExport].
//
// Creating a [JSValue] instance that wraps a native object retains the
// underlying Objective-C or Swift object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(object:in:)
func NewJSValueWithObjectInContext(value objectivec.IObject, context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithObject:inContext:"), value, context)
	return JSValueFromID(rv)
}

// Creates a JavaScript representation of the specified point.
//
// point: A CoreGraphics point structure.
//
// context: The JavaScript context in which to create the value.
//
// # Return Value
//
// A JavaScript object representing the specified point.
//
// # Discussion
//
// Converting a point creates a JavaScript object value with fields named `x`
// and `y`.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(point:inContext:)
func NewJSValueWithPointInContext(point corefoundation.CGPoint, context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithPoint:inContext:"), point, context)
	return JSValueFromID(rv)
}

// Creates a JavaScript representation of the specified range.
//
// range: A range.
//
// context: The JavaScript context in which to create the value.
//
// # Return Value
//
// A JavaScript object representing the specified range.
//
// # Discussion
//
// Converting a range creates a JavaScript object value with fields named
// `location` and `length`.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(range:inContext:)
func NewJSValueWithRangeInContext(range_ foundation.NSRange, context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithRange:inContext:"), range_, context)
	return JSValueFromID(rv)
}

// Creates a JavaScript representation of the specified rectangle.
//
// rect: A CoreGraphics rectangle structure.
//
// context: The JavaScript context in which to create the value.
//
// # Return Value
//
// A JavaScript object representing the specified rectangle.
//
// # Discussion
//
// Converting a rectangle creates a JavaScript object value with fields named
// `x`, `y`, `width`, and `height`.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(rect:inContext:)
func NewJSValueWithRectInContext(rect corefoundation.CGRect, context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithRect:inContext:"), rect, context)
	return JSValueFromID(rv)
}

// Creates a JavaScript representation of the specified width and height.
//
// size: A CoreGraphics size structure.
//
// context: The JavaScript context in which to create the value.
//
// # Return Value
//
// A JavaScript object representing the specified size.
//
// # Discussion
//
// Converting a rectangle creates a JavaScript object value with fields named
// `width` and `height`.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(size:inContext:)
func NewJSValueWithSizeInContext(size corefoundation.CGSize, context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithSize:inContext:"), size, context)
	return JSValueFromID(rv)
}

// Creates a JavaScript representation of the specified unsigned integer
// value.
//
// value: A native 32-bit unsigned integer value.
//
// context: The JavaScript context in which to create the value.
//
// # Return Value
//
// A JavaScript numeric value.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(uInt32:in:)
func NewJSValueWithUInt32InContext(value uint32, context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithUInt32:inContext:"), value, context)
	return JSValueFromID(rv)
}

// Creates a JavaScript `undefined` value.
//
// context: The JavaScript context in which to create the value.
//
// # Return Value
//
// The JavaScript `undefined` value.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(undefinedIn:)
func NewJSValueWithUndefinedInContext(context IJSContext) JSValue {
	rv := objc.Send[objc.ID](objc.ID(getJSValueClass().class), objc.Sel("valueWithUndefinedInContext:"), context)
	return JSValueFromID(rv)
}

// Converts the JavaScript value to a native object.
//
// # Return Value
//
// An Objective-C or Swift object representing the JavaScript value.
//
// # Discussion
//
// The type of the resulting object depends on the contents of the JavaScript
// value. For conversion rules, see [JSValue].
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toObject()
func (j JSValue) ToObject() objectivec.IObject {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("toObject"))
	return objectivec.Object{ID: rv}
}

// Converts the JavaScript value to a native object of the specified class.
//
// expectedClass: The Objective-C or Swift class type to convert the value to.
//
// # Return Value
//
// An Objective-C or Swift object representing the JavaScript value, or `nil`
// if the value cannot be converted to the expected class.
//
// # Discussion
//
// Use this method to enforce a specific type conversion from JavaScript, or
// to retrieve Objective-C or Swift objects of custom classes that were
// bridged into JavaScript using the [JSExport] protocol.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toObjectOf(_:)
func (j JSValue) ToObjectOfClass(expectedClass objectivec.Class) objectivec.IObject {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("toObjectOfClass:"), expectedClass)
	return objectivec.Object{ID: rv}
}

// Converts the JavaScript value to a native Boolean value.
//
// # Return Value
//
// The native Boolean value.
//
// # Discussion
//
// This method uses JavaScript type coercion to convert the value to a
// JavaScript Boolean value, then returns the native representation of the
// result. Thus, this method can return true even when the [JSValue.IsBoolean]
// property does not.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toBool()
func (j JSValue) ToBool() bool {
	rv := objc.Send[bool](j.ID, objc.Sel("toBool"))
	return rv
}

// Converts the JavaScript value to a native floating-point value.
//
// # Return Value
//
// The native double-precision floating-point value.
//
// # Discussion
//
// This method uses JavaScript type coercion to convert the value to a
// JavaScript numeric value, then returns a native representation of the
// result. In JavaScript, all numeric values are treated as double-precision
// floating-point numbers except for certain operations such as bit shifts.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toDouble()
func (j JSValue) ToDouble() float64 {
	rv := objc.Send[float64](j.ID, objc.Sel("toDouble"))
	return rv
}

// Converts the JavaScript value to a native signed integer value.
//
// # Return Value
//
// The native signed 32-bit integer value.
//
// # Discussion
//
// This method uses JavaScript type coercion to convert the value to a
// JavaScript integer value, then returns a native representation of the
// result. In JavaScript, all numeric values are treated as double-precision
// floating-point numbers except for certain operations such as bit shifts.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toInt32()
func (j JSValue) ToInt32() int32 {
	rv := objc.Send[int32](j.ID, objc.Sel("toInt32"))
	return rv
}

// Converts the JavaScript value to a native unsigned integer value.
//
// # Return Value
//
// The native unsigned 32-bit integer value.
//
// # Discussion
//
// This method uses JavaScript type coercion to convert the value to a
// JavaScript integer value, then returns a native representation of the
// result. In JavaScript, all numeric values are treated as double-precision
// floating-point numbers except for certain operations such as bit shifts.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toUInt32()
func (j JSValue) ToUInt32() uint32 {
	rv := objc.Send[uint32](j.ID, objc.Sel("toUInt32"))
	return rv
}

// Converts the JavaScript value to a [NSNumber] object.
//
// # Return Value
//
// A [NSNumber] object encapsulating the native representation of the value.
//
// # Discussion
//
// If the value represents a Boolean value, the resulting [NSNumber] object is
// created as with the
// https://developer.apple.com/documentation/foundation/nsnumber/1551475-numberwithbool
// method. Otherwise, this method uses JavaScript type coercion to convert the
// value to a JavaScript numeric value and creates a [NSNumber] object
// wrapping the result.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toNumber()
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
//
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
// [NSNumber]: https://developer.apple.com/documentation/Foundation/NSNumber
func (j JSValue) ToNumber() foundation.NSNumber {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("toNumber"))
	return foundation.NSNumberFromID(rv)
}

// Converts the JavaScript value to a native string.
//
// # Return Value
//
// The string representation of the value.
//
// # Discussion
//
// This method uses JavaScript type coercion rules to convert the value to a
// JavaScript string, then creates a native string from the result. Thus, this
// method can return a string even when the [JSValue.IsString] property is
// false; for example, an empty object becomes the string `"[object Object]"`.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toString()
func (j JSValue) ToString() string {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("toString"))
	return foundation.NSStringFromID(rv).String()
}

// Converts the JavaScript value to a date object.
//
// # Return Value
//
// The date representation of the value.
//
// # Discussion
//
// If the value contains a JavaScript [Date] object, this method returns an
// equivalent [NSDate] representation. Otherwise, this method uses JavaScript
// type coercion to interpret the value as a number of seconds and creates an
// [NSDate] object with the
// https://developer.apple.com/documentation/foundation/nsdate/1591576-datewithtimeintervalsince1970
// method.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toDate()
//
// [NSDate]: https://developer.apple.com/documentation/Foundation/NSDate
func (j JSValue) ToDate() foundation.NSDate {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("toDate"))
	return foundation.NSDateFromID(rv)
}

// Converts the JavaScript value to an array.
//
// # Return Value
//
// The array representation of the value.
//
// # Discussion
//
// If the value is a JavaScript object, this method reads the object’s
// `length` property as an unsigned integer, creates an [NSArray] object of
// the corresponding size, and recursively copies and converts any properties
// corresponding to indices within the array bounds. JavaScript converts each
// element to a native object using the rules listed in [JSValue].
//
// This method returns `nil` if the JavaScript value is `null` or `undefined`,
// and throws a JavaScript [TypeError] if the value is not a JavaScript
// object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toArray()
//
// [NSArray]: https://developer.apple.com/documentation/Foundation/NSArray
func (j JSValue) ToArray() foundation.INSArray {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("toArray"))
	return foundation.NSArrayFromID(rv)
}

// Converts the JavaScript value to a dictionary.
//
// # Return Value
//
// The dictionary representation of the value.
//
// # Discussion
//
// If the value is a JavaScript object, this method creates an [NSDictionary]
// object of the corresponding size, and recursively copies and converts all
// enumerable properties of the object into the dictionary with
// correspondingly named keys. JavaScript converts each element to a native
// object using the rules listed in [JSValue].
//
// This method returns `nil` if the JavaScript value is `null` or `undefined`,
// and throws a JavaScript [TypeError] if the value is not a JavaScript
// object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toDictionary()
//
// [NSDictionary]: https://developer.apple.com/documentation/Foundation/NSDictionary
func (j JSValue) ToDictionary() foundation.INSDictionary {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("toDictionary"))
	return foundation.NSDictionaryFromID(rv)
}

// Converts the value to a point structure.
//
// # Return Value
//
// A CoreGraphics point representation of the value.
//
// # Discussion
//
// This method treats the value as a JavaScript object, reading the values of
// its `x` and `y` properties using the [JSValue.ToDouble] method and creating
// a [CGPoint] structure from the result. If the value is not a JavaScript
// object or does not have the appropriate properties, each of the resulting
// point’s coordinates is not a number (NaN).
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toPoint()
//
// [CGPoint]: https://developer.apple.com/documentation/CoreFoundation/CGPoint
func (j JSValue) ToPoint() corefoundation.CGPoint {
	rv := objc.Send[corefoundation.CGPoint](j.ID, objc.Sel("toPoint"))
	return corefoundation.CGPoint(rv)
}

// Converts the value to a range.
//
// # Return Value
//
// A range representation of the value.
//
// # Discussion
//
// This method treats the value as a JavaScript object, reading the values of
// its `location` and `length` properties using the [JSValue.ToDouble] method
// and creating a [NSRange] structure from the result. If the value is not a
// JavaScript object or does not have the appropriate properties, the
// resulting range is invalid.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toRange()
//
// [NSRange]: https://developer.apple.com/documentation/Foundation/NSRange-c.struct
func (j JSValue) ToRange() foundation.NSRange {
	rv := objc.Send[foundation.NSRange](j.ID, objc.Sel("toRange"))
	return foundation.NSRange(rv)
}

// Converts the value to a rectangle structure.
//
// # Return Value
//
// A CoreGraphics point representation of the value.
//
// # Discussion
//
// This method treats the value as a JavaScript object, reading the values of
// its `x`, `y`, `width`, and `height` properties using the [JSValue.ToDouble]
// method and creating a [CGRect] structure from the result. If the value is
// not a JavaScript object or does not have the appropriate properties, each
// of the resulting rectangle’s coordinates is not a number (NaN).
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toRect()
//
// [CGRect]: https://developer.apple.com/documentation/CoreFoundation/CGRect
func (j JSValue) ToRect() corefoundation.CGRect {
	rv := objc.Send[corefoundation.CGRect](j.ID, objc.Sel("toRect"))
	return corefoundation.CGRect(rv)
}

// Converts the value to a size.
//
// # Return Value
//
// A CoreGraphics size representation of the value.
//
// # Discussion
//
// This method treats the value as a JavaScript object, reading the values of
// its `width` and `height` properties using the [JSValue.ToDouble] method and
// creating a [CGSize] structure from the result. If the value is not a
// JavaScript object or does not have the appropriate properties, the size’s
// width and height are each not a number (NaN).
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toSize()
//
// [CGSize]: https://developer.apple.com/documentation/CoreFoundation/CGSize
func (j JSValue) ToSize() corefoundation.CGSize {
	rv := objc.Send[corefoundation.CGSize](j.ID, objc.Sel("toSize"))
	return corefoundation.CGSize(rv)
}

// Compares the value to another for strict equality.
//
// value: The value to be compared against.
//
// # Return Value
//
// true if the values are strictly equal; otherwise, false.
//
// # Discussion
//
// This method is analogous to the identity or strict equality operator `===`
// in JavaScript.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isEqual(to:)
func (j JSValue) IsEqualToObject(value objectivec.IObject) bool {
	rv := objc.Send[bool](j.ID, objc.Sel("isEqualToObject:"), value)
	return rv
}

// Compares the value to another for equivalence, allowing type conversion.
//
// value: The value to be compared against.
//
// # Return Value
//
// true if the values are equivalent; otherwise, false.
//
// # Discussion
//
// This method is analogous to the equality operator `==` in JavaScript: it
// first converts its operands to the same type (if they are not already of
// the same type), then applies a strict equality comparison to the result.
// JavaScript object values are equal if and only if they refer to the same
// object instance.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isEqualWithTypeCoercion(to:)
func (j JSValue) IsEqualWithTypeCoercionToObject(value objectivec.IObject) bool {
	rv := objc.Send[bool](j.ID, objc.Sel("isEqualWithTypeCoercionToObject:"), value)
	return rv
}

// Returns a Boolean value indicating whether the value is an instance of
// another JavaScript object value.
//
// value: The value to be compared against.
//
// # Return Value
//
// true if this value inherits from `value`; otherwise, false.
//
// # Discussion
//
// This method is analogous to the `instanceof` operator in JavaScript: it
// tests for the presence of the specified value’s constructor prototype in
// this value’s prototype chain.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isInstance(of:)
func (j JSValue) IsInstanceOf(value objectivec.IObject) bool {
	rv := objc.Send[bool](j.ID, objc.Sel("isInstanceOf:"), value)
	return rv
}

// Invokes the value as a JavaScript function.
//
// arguments: The parameters to pass to the function. The objects in this array must be
// other [JSValue] objects or objects that can be converted to JavaScript
// values using the methods listed in the Creating JavaScript Values section
// in [JSValue].
//
// # Return Value
//
// The result of calling the value as a function, or `nil` if the value cannot
// be treated as a JavaScript function.
//
// # Discussion
//
// In JavaScript, if a function does not explicitly return a value, it
// implicitly returns the value `undefined`—use the [JSValue.IsUndefined]
// property to test for this result.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/call(withArguments:)
func (j JSValue) CallWithArguments(arguments foundation.INSArray) IJSValue {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("callWithArguments:"), arguments)
	return JSValueFromID(rv)
}

// Invokes the value as a JavaScript constructor.
//
// arguments: The parameters to pass to the constructor. The objects in this array must
// be other [JSValue] objects or objects that can be converted to JavaScript
// values using the methods listed in Creating JavaScript Values.
//
// # Return Value
//
// The result of calling the value as a constructor, or `nil` if the value
// cannot be treated as a JavaScript constructor.
//
// # Discussion
//
// Calling a constructor is equivalent to using the `new` keyword in
// JavaScript.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/construct(withArguments:)
func (j JSValue) ConstructWithArguments(arguments foundation.INSArray) IJSValue {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("constructWithArguments:"), arguments)
	return JSValueFromID(rv)
}

// Calls the named JavaScript method on the value.
//
// method: The name of a method on the value; that is, of a field whose contents are a
// function value.
//
// arguments: The parameters to pass to the method. The objects in this array must be
// other [JSValue] objects or objects that can be converted to JavaScript
// values using the methods listed in the Creating JavaScript Values section
// in [JSValue].
//
// # Return Value
//
// The result of calling the value as a constructor, or `nil` if the value
// cannot be treated as a JavaScript constructor.
//
// # Discussion
//
// Calling this Objective-C method first uses the [JSValue.ValueForProperty]
// method to look up the named field of the JavaScript value. Then,
// JavaScriptCore treats that field’s contents as a JavaScript function and
// sets the JavaScript `this` keyword to refer to this [JSValue] instance.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/invokeMethod(_:withArguments:)
func (j JSValue) InvokeMethodWithArguments(method string, arguments foundation.INSArray) IJSValue {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("invokeMethod:withArguments:"), objc.String(method), arguments)
	return JSValueFromID(rv)
}

// Defines a property on the JavaScript object value or modifies a
// property’s definition.
//
// property: The name of the property to define or modify.
//
// descriptor: A JavaScript object whose keys and values define the property’s behavior.
//
// # Discussion
//
// Calling this method is equivalent to using the `Object.DefineProperty()`
// method in JavaScript. The `descriptor` parameter has the same format
// required by that JavaScript method; for convenience when calling from
// Objective-C or Swift, you can also construct it as a dictionary with the
// keys listed in [Property Descriptor Keys].
//
// The descriptor determines the behavior of the JavaScript property, and must
// fit one of three cases:
//
// - Data Descriptor: Contains one or both of the keys `value` and `writable`,
// and optionally also contains the keys `enumerable` or `configurable`.
// Cannot contain the keys `get` or `set`. Use a data descriptor to create or
// modify the attributes of a data property on an object (replacing any
// existing accessor property). - Accessor Descriptor: Contains one or both of
// the keys `get` or `set`, and optionally also contains the keys `enumerable`
// or `configurable`. Cannot contain the keys `value` and `writable`. Use an
// accessor descriptor to create or modify the attributes of an accessor
// property on an object (replacing any existing data property). - Generic
// Descriptor: Contains one or both of the keys `enumerable` or
// `configurable`, and cannot contain any other keys. Use a genetic descriptor
// to modify the attributes of an existing data or accessor property, or to
// create a new data property.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/defineProperty(_:descriptor:)
//
// [Property Descriptor Keys]: https://developer.apple.com/documentation/JavaScriptCore/property-descriptor-keys
func (j JSValue) DefinePropertyDescriptor(property JSValueProperty, descriptor objectivec.IObject) {
	objc.Send[objc.ID](j.ID, objc.Sel("defineProperty:descriptor:"), property, descriptor)
}

// Returns a Boolean value indicating whether the JavaScript value has a
// defined property with the specified name.
//
// property: The name of a property to query for in the JavaScript object value.
//
// # Return Value
//
// true if the JavaScript object has a defined property by that name;
// otherwise, false.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/hasProperty(_:)
func (j JSValue) HasProperty(property JSValueProperty) bool {
	rv := objc.Send[bool](j.ID, objc.Sel("hasProperty:"), property)
	return rv
}

// Deletes the named property from the JavaScript object value.
//
// property: The name of a property in the JavaScript object value.
//
// # Return Value
//
// true if property deletion was successful; otherwise, false.
//
// # Discussion
//
// Calling this method is equivalent to using the JavaScript `delete` operator
// on an object (for example, `delete object.Property()`). After deletion,
// attempting to retrieve the property’s value results in the undefined
// value, and any descriptor information that defines the property’s
// behavior (see the [JSValue.DefinePropertyDescriptor] method or the
// JavaScript `defineProperty` function) is lost.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/deleteProperty(_:)
func (j JSValue) DeleteProperty(property JSValueProperty) bool {
	rv := objc.Send[bool](j.ID, objc.Sel("deleteProperty:"), property)
	return rv
}

// Returns the value at the specified numeric index in the JavaScript object
// value.
//
// index: An index in the JavaScript object.
//
// # Return Value
//
// The value at the specified index, or the JavaScript `undefined` value if no
// property exists at that index.
//
// # Discussion
//
// Calling this method is equivalent to using the subscript operator with a
// numeric subscript in JavaScript. Use it to access elements of JavaScript
// arrays or of objects with numerically-indexed properties.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/atIndex(_:)
func (j JSValue) ValueAtIndex(index uint) IJSValue {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("valueAtIndex:"), index)
	return JSValueFromID(rv)
}

// Sets the value at the specified numeric index in the JavaScript object
// value.
//
// value: The value to set at the specified index.
//
// index: An index in the JavaScript object.
//
// # Discussion
//
// Calling this method is equivalent to using the subscript operator with a
// numeric subscript in JavaScript. Use it to access elements of JavaScript
// arrays or of objects with numerically-indexed properties.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/setValue(_:at:)
func (j JSValue) SetValueAtIndex(value objectivec.IObject, index uint) {
	objc.Send[objc.ID](j.ID, objc.Sel("setValue:atIndex:"), value, index)
}

// Returns the value of the named property in the JavaScript object value.
//
// property: The name of a property in the JavaScript object.
//
// # Return Value
//
// The value of the named property, or the JavaScript `undefined` value if no
// property exists by that name.
//
// # Discussion
//
// Calling this method is equivalent to using the subscript operator with a
// string subscript in JavaScript. Use it to access fields or properties in
// JavaScript objects.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/forProperty(_:)
func (j JSValue) ValueForProperty(property JSValueProperty) IJSValue {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("valueForProperty:"), property)
	return JSValueFromID(rv)
}

// Sets the value of the named property in the JavaScript object value.
//
// value: The value to set for the named property.
//
// property: The name of a property in the JavaScript object.
//
// # Discussion
//
// Calling this method is equivalent to using the subscript operator with a
// string subscript in JavaScript. Use it to set or create fields or
// properties in JavaScript objects.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/setValue(_:forProperty:)
func (j JSValue) SetValueForProperty(value objectivec.IObject, property JSValueProperty) {
	objc.Send[objc.ID](j.ID, objc.Sel("setValue:forProperty:"), value, property)
}

// Returns the value’s JavaScript property at the specified index, allowing
// subscript syntax.
//
// index: An index in the JavaScript object.
//
// # Return Value
//
// The value at the specified index, or the JavaScript `undefined` value if no
// property exists at that index.
//
// # Discussion
//
// This method is equivalent to the [JSValue.ValueAtIndex] method, but
// provides Objective-C subscripting support.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/objectAtIndexedSubscript(_:)
func (j JSValue) ObjectAtIndexedSubscript(index uint) IJSValue {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("objectAtIndexedSubscript:"), index)
	return JSValueFromID(rv)
}

// Sets the value’s JavaScript property at the specified index, allowing
// subscript syntax.
//
// object: The value to set at the specified index.
//
// index: An index in the JavaScript object.
//
// # Discussion
//
// This method is equivalent to the [JSValue.SetValueAtIndex] method, but
// provides Objective-C subscripting support.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/setObject(_:atIndexedSubscript:)
func (j JSValue) SetObjectAtIndexedSubscript(object objectivec.IObject, index uint) {
	objc.Send[objc.ID](j.ID, objc.Sel("setObject:atIndexedSubscript:"), object, index)
}

// Returns the value’s JavaScript property named with the specified key,
// allowing subscript syntax.
//
// key: The name of a property in the JavaScript object.
//
// # Return Value
//
// The value of the named property, or the JavaScript `undefined` value if no
// property exists by that name.
//
// # Discussion
//
// This method is equivalent to the [JSValue.ValueForProperty] method, but
// provides Objective-C subscripting support.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/objectForKeyedSubscript(_:)
func (j JSValue) ObjectForKeyedSubscript(key objectivec.IObject) IJSValue {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("objectForKeyedSubscript:"), key)
	return JSValueFromID(rv)
}

// Sets the value’s JavaScript property named with the specified key,
// allowing subscript syntax.
//
// object: The value to set for the named JavaScript property.
//
// key: The name of a property in the JavaScript object.
//
// # Discussion
//
// This method is equivalent to the [JSValue.SetValueForProperty] method, but
// provides Objective-C subscripting support.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/setObject(_:forKeyedSubscript:)
func (j JSValue) SetObjectForKeyedSubscript(object objectivec.IObject, key objectivec.IObject) {
	objc.Send[objc.ID](j.ID, objc.Sel("setObject:forKeyedSubscript:"), object, key)
}

// # Return Value
//
// A value of JSRelationCondition, a kJSRelationConditionUndefined is returned
// if an exception is thrown.
//
// # Discussion
//
// Compare a JSValue with a double.
//
// The JSValue is converted to a double according to the rules specified by
// the JavaScript language then compared with other.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/compare(_:)-35b2t
func (j JSValue) CompareDouble(other float64) JSRelationCondition {
	rv := objc.Send[JSRelationCondition](j.ID, objc.Sel("compareDouble:"), other)
	return JSRelationCondition(rv)
}

// # Return Value
//
// A value of JSRelationCondition, a kJSRelationConditionUndefined is returned
// if an exception is thrown.
//
// # Discussion
//
// Compare two JSValues.
//
// The result is computed by comparing the results of JavaScript’s ==,
// operators. If either self or other is (or would coerce to) NaN in
// JavaScript, then the result is kJSRelationConditionUndefined.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/compare(_:)-5w184
func (j JSValue) CompareJSValue(other IJSValue) JSRelationCondition {
	rv := objc.Send[JSRelationCondition](j.ID, objc.Sel("compareJSValue:"), other)
	return JSRelationCondition(rv)
}

// # Return Value
//
// A value of JSRelationCondition, a kJSRelationConditionUndefined is returned
// if an exception is thrown.
//
// # Discussion
//
// Compare a JSValue with a uint64_t.
//
// The JSValue is converted to an integer according to the rules specified by
// the JavaScript language then compared with other.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/compare(_:)-64n3k
func (j JSValue) CompareUInt64(other uint64) JSRelationCondition {
	rv := objc.Send[JSRelationCondition](j.ID, objc.Sel("compareUInt64:"), other)
	return JSRelationCondition(rv)
}

// # Return Value
//
// A value of JSRelationCondition, a kJSRelationConditionUndefined is returned
// if an exception is thrown.
//
// # Discussion
//
// Compare a JSValue with a int64_t.
//
// The JSValue is converted to an integer according to the rules specified by
// the JavaScript language then compared with other.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/compare(_:)-9d4zq
func (j JSValue) CompareInt64(other int64) JSRelationCondition {
	rv := objc.Send[JSRelationCondition](j.ID, objc.Sel("compareInt64:"), other)
	return JSRelationCondition(rv)
}

// # Discussion
//
// Convert a JSValue to a int64_t.
//
// The JSValue is converted to an integer according to the rules specified by
// the JavaScript language. If the value is a BigInt, then the value is
// truncated to an int64_t.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toInt64()
func (j JSValue) ToInt64() int64 {
	rv := objc.Send[int64](j.ID, objc.Sel("toInt64"))
	return rv
}

// # Discussion
//
// Convert a JSValue to a uint64_t.
//
// The JSValue is converted to an integer according to the rules specified by
// the JavaScript language. If the value is a BigInt, then the value is
// truncated to a uint64_t.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/toUInt64()
func (j JSValue) ToUInt64() uint64 {
	rv := objc.Send[uint64](j.ID, objc.Sel("toUInt64"))
	return rv
}

// Creates a promise object using the specified executor callback.
//
// context: The [JSContext] the resulting [JSValue] belongs to.
//
// callback: A callback block to invoke during initialization of the promise object. The
// `resolve` and `reject` parameters are functions that you can call to notify
// any pending reactions about the state of the new promise object.
//
// # Return Value
//
// A [JSValue] that represents a new promise JavaScript object.
//
// # Discussion
//
// This method is equivalent to calling the `Promise()` constructor in
// JavaScript.
//
// The `resolve` and `reject` callbacks each typically take a single value,
// which they forward to all relevant pending reactions. While inside the
// executor callback, `context` acts as if it is in any other callback, except
// `calleeFunction` is `nil`. This also means you can access the new promise
// object using `[context thisValue]`.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/init(newPromiseIn:fromExecutor:)
func (_JSValueClass JSValueClass) ValueWithNewPromiseInContextFromExecutor(context IJSContext, callback JSValueJSValueHandler) JSValue {
	_block1, _ := NewJSValueJSValueBlock(callback)
	rv := objc.Send[objc.ID](objc.ID(_JSValueClass.class), objc.Sel("valueWithNewPromiseInContext:fromExecutor:"), context, _block1)
	return JSValueFromID(rv)
}

// A Boolean value that indicates whether the instance corresponds to the
// JavaScript `undefined` value.
//
// # Discussion
//
// The JavaScript `undefined` value is used for variables that have not yet
// been assigned a value, for formal parameters in functions for which no
// actual parameter has been passed, and as the result of expressions or
// function calls that do not explicitly return a value. Note that `undefined`
// is not the same as `null`.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isUndefined
func (j JSValue) IsUndefined() bool {
	rv := objc.Send[bool](j.ID, objc.Sel("isUndefined"))
	return rv
}

// A Boolean value that indicates whether the instance corresponds to the
// JavaScript `null` value.
//
// # Discussion
//
// The JavaScript `null` value is used only in cases where an actual value is
// expected but none is applicable. Note that `null` is not the same as
// `undefined`.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isNull
func (j JSValue) IsNull() bool {
	rv := objc.Send[bool](j.ID, objc.Sel("isNull"))
	return rv
}

// A Boolean value that indicates whether the instance is a JavaScript Boolean
// value.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isBoolean
func (j JSValue) IsBoolean() bool {
	rv := objc.Send[bool](j.ID, objc.Sel("isBoolean"))
	return rv
}

// A Boolean value that indicates whether the instance is a JavaScript numeric
// value.
//
// # Discussion
//
// In JavaScript, there is no differentiation between types of numbers.
// Semantically, all numbers behave as double-precision floating-point types,
// except in special cases like bit operations.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isNumber
func (j JSValue) IsNumber() bool {
	rv := objc.Send[bool](j.ID, objc.Sel("isNumber"))
	return rv
}

// A Boolean value that indicates whether the instance is a JavaScript
// [String] object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isString
func (j JSValue) IsString() bool {
	rv := objc.Send[bool](j.ID, objc.Sel("isString"))
	return rv
}

// A Boolean value that indicates whether the instance is a JavaScript object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isObject
func (j JSValue) IsObject() bool {
	rv := objc.Send[bool](j.ID, objc.Sel("isObject"))
	return rv
}

// A Boolean value that indicates whether the instance is a JavaScript array
// value.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isArray
func (j JSValue) IsArray() bool {
	rv := objc.Send[bool](j.ID, objc.Sel("isArray"))
	return rv
}

// A Boolean value that indicates whether the instance is a JavaScript [Date]
// object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isDate
func (j JSValue) IsDate() bool {
	rv := objc.Send[bool](j.ID, objc.Sel("isDate"))
	return rv
}

// A Boolean value that indicates whether the instance is a symbol.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isSymbol
func (j JSValue) IsSymbol() bool {
	rv := objc.Send[bool](j.ID, objc.Sel("isSymbol"))
	return rv
}

// The JavaScript context hosting this value.
//
// # Discussion
//
// A value maintains a strong reference to its enclosing JavaScript
// environment (a [JSContext] object). As such, you should not store
// JavaScript values inside objects that are owned by the same [JSContext]
// object, as this action creates a retain cycle. To properly manage memory
// when storing [JSValue] instances, use the [JSManagedValue] class.
//
// You can pass a value to other JavaScript contexts with the same virtual
// machine, but not to contexts with other virtual machines. Use the
// [JSContext.VirtualMachine] property of a value’s context to determine
// which other contexts can use the value.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/context
func (j JSValue) Context() IJSContext {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("context"))
	return JSContextFromID(objc.ID(rv))
}

// Returns the C representation of the JavaScript value.
//
// # Discussion
//
// See [JSValueRef] for the C JavaScriptCore API.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/jsValueRef
func (j JSValue) JSValueRef() JSValueRef {
	rv := objc.Send[JSValueRef](j.ID, objc.Sel("JSValueRef"))
	return JSValueRef(rv)
}

// # Discussion
//
// Check if a JSValue is a BigInt.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValue/isBigInt
func (j JSValue) IsBigInt() bool {
	rv := objc.Send[bool](j.ID, objc.Sel("isBigInt"))
	return rv
}
