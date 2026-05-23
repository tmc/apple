// Code generated from Apple documentation. DO NOT EDIT.

package javascriptcore

import (
	"github.com/tmc/apple/kernel"
	"github.com/tmc/apple/objectivec"
)

// JSChar is a Unicode character.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSChar
type JSChar = uint16

// JSClassAttributes is a set of JavaScript class attributes.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSClassAttributes
type JSClassAttributes = uint

// JSClassRef is a JavaScript class.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSClassRef
type JSClassRef uintptr

// JSContextGroupRef is a group that associates JavaScript contexts with one another.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContextGroupRef
type JSContextGroupRef uintptr

// JSContextRef is a JavaScript execution context.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContextRef
type JSContextRef uintptr

// JSGlobalContextRef is a global JavaScript execution context.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextRef
type JSGlobalContextRef uintptr

// JSObjectCallAsConstructorCallback is the callback type for using an object as a constructor.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectCallAsConstructorCallback
type JSObjectCallAsConstructorCallback = func(uintptr, uintptr, uint, uintptr, uintptr) uintptr

// JSObjectCallAsFunctionCallback is the callback type for calling an object as a function.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectCallAsFunctionCallback
type JSObjectCallAsFunctionCallback = func(uintptr, uintptr, uintptr, uint, uintptr, uintptr) uintptr

// JSObjectConvertToTypeCallback is the callback type for converting an object to a particular JavaScript type.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectConvertToTypeCallback
type JSObjectConvertToTypeCallback = func(uintptr, uintptr, JSType, uintptr) uintptr

// JSObjectDeletePropertyCallback is the callback type for deleting a property.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectDeletePropertyCallback
type JSObjectDeletePropertyCallback = func(uintptr, uintptr, uintptr, uintptr) bool

// JSObjectFinalizeCallback is the callback type for finalizing an object (preparing it for garbage collection).
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectFinalizeCallback
type JSObjectFinalizeCallback = func(uintptr)

// JSObjectGetPropertyCallback is the callback type for getting a property’s value.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetPropertyCallback
type JSObjectGetPropertyCallback = func(uintptr, uintptr, uintptr, uintptr) uintptr

// JSObjectGetPropertyNamesCallback is the callback type for collecting the names of an object’s properties.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetPropertyNamesCallback
type JSObjectGetPropertyNamesCallback = func(uintptr, uintptr, uintptr)

// JSObjectHasInstanceCallback is the callback type for checking whether an object is an instance of a particular type.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectHasInstanceCallback
type JSObjectHasInstanceCallback = func(uintptr, uintptr, uintptr, uintptr) bool

// JSObjectHasPropertyCallback is the callback type for determining whether an object has a property.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectHasPropertyCallback
type JSObjectHasPropertyCallback = func(uintptr, uintptr, uintptr) bool

// JSObjectInitializeCallback is the callback type for first creating an object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectInitializeCallback
type JSObjectInitializeCallback = func(uintptr, uintptr)

// JSObjectRef is a JavaScript object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectRef
type JSObjectRef uintptr

// JSObjectSetPropertyCallback is the callback type for setting a property’s value.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectSetPropertyCallback
type JSObjectSetPropertyCallback = func(uintptr, uintptr, uintptr, uintptr, uintptr) bool

// JSPropertyAttributes is a set of JavaScript property attributes.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyAttributes
type JSPropertyAttributes = uint

// JSPropertyNameAccumulatorRef is an ordered set of the names of a JavaScript object’s properties.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyNameAccumulatorRef
type JSPropertyNameAccumulatorRef uintptr

// JSPropertyNameArrayRef is an array of JavaScript property names.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyNameArrayRef
type JSPropertyNameArrayRef uintptr

// JSStringRef is a UTF-16 character buffer.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSStringRef
type JSStringRef uintptr

// JSTypedArrayBytesDeallocator is a function that deallocates bytes that pass to a typed array constructor.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSTypedArrayBytesDeallocator
type JSTypedArrayBytesDeallocator = func(kernel.Pointer, kernel.Pointer)

// JSValueProperty is a type that identifies a property of a JavaScript value.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueProperty
type JSValueProperty = objectivec.Object

// JSValueRef is a JavaScript value.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueRef
type JSValueRef uintptr
