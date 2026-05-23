// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

import (
	"fmt"
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/corefoundation"
)

type unavailableSymbolError struct {
	symbol     string
	introduced string
	cause      error
}

func (e *unavailableSymbolError) Error() string {
	if e == nil {
		return ""
	}
	if e.introduced != "" {
		return fmt.Sprintf("JavaScriptCore: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("JavaScriptCore: symbol %s unavailable on this system", e.symbol)
}

func (e *unavailableSymbolError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

func missingSymbolError(name, introduced string, cause error) error {
	return &unavailableSymbolError{
		symbol:     name,
		introduced: introduced,
		cause:      cause,
	}
}

func symbolCallError(name, introduced string, err error) error {
	if err != nil {
		return err
	}
	if frameworkHandle == 0 {
		return fmt.Errorf("JavaScriptCore: symbol %s unavailable because the framework could not be loaded", name)
	}
	return missingSymbolError(name, introduced, nil)
}

// registerFunc resolves a framework symbol and registers it as a Go function.
func registerFunc(fptr any, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			*errDst = fmt.Errorf("JavaScriptCore: register symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	*errDst = nil
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	*dst = sym
	*errDst = nil
}

var _jSBigIntCreateWithDouble func(ctx JSContextRef, value float64, exception *JSValueRef) JSValueRef
var _jSBigIntCreateWithDoubleErr error

func tryJSBigIntCreateWithDouble(ctx JSContextRef, value float64, exception *JSValueRef) (JSValueRef, error) {
	if _jSBigIntCreateWithDouble == nil {
		return *new(JSValueRef), symbolCallError("JSBigIntCreateWithDouble", "15.0", _jSBigIntCreateWithDoubleErr)
	}
	return _jSBigIntCreateWithDouble(ctx, value, exception), nil
}

// JSBigIntCreateWithDouble.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSBigIntCreateWithDouble(_:_:_:)
func JSBigIntCreateWithDouble(ctx JSContextRef, value float64, exception *JSValueRef) JSValueRef {
	result, callErr := tryJSBigIntCreateWithDouble(ctx, value, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSBigIntCreateWithInt64 func(ctx JSContextRef, integer int64, exception *JSValueRef) JSValueRef
var _jSBigIntCreateWithInt64Err error

func tryJSBigIntCreateWithInt64(ctx JSContextRef, integer int64, exception *JSValueRef) (JSValueRef, error) {
	if _jSBigIntCreateWithInt64 == nil {
		return *new(JSValueRef), symbolCallError("JSBigIntCreateWithInt64", "15.0", _jSBigIntCreateWithInt64Err)
	}
	return _jSBigIntCreateWithInt64(ctx, integer, exception), nil
}

// JSBigIntCreateWithInt64.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSBigIntCreateWithInt64(_:_:_:)
func JSBigIntCreateWithInt64(ctx JSContextRef, integer int64, exception *JSValueRef) JSValueRef {
	result, callErr := tryJSBigIntCreateWithInt64(ctx, integer, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSBigIntCreateWithString func(ctx JSContextRef, string_ JSStringRef, exception *JSValueRef) JSValueRef
var _jSBigIntCreateWithStringErr error

func tryJSBigIntCreateWithString(ctx JSContextRef, string_ JSStringRef, exception *JSValueRef) (JSValueRef, error) {
	if _jSBigIntCreateWithString == nil {
		return *new(JSValueRef), symbolCallError("JSBigIntCreateWithString", "15.0", _jSBigIntCreateWithStringErr)
	}
	return _jSBigIntCreateWithString(ctx, string_, exception), nil
}

// JSBigIntCreateWithString.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSBigIntCreateWithString(_:_:_:)
func JSBigIntCreateWithString(ctx JSContextRef, string_ JSStringRef, exception *JSValueRef) JSValueRef {
	result, callErr := tryJSBigIntCreateWithString(ctx, string_, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSBigIntCreateWithUInt64 func(ctx JSContextRef, integer uint64, exception *JSValueRef) JSValueRef
var _jSBigIntCreateWithUInt64Err error

func tryJSBigIntCreateWithUInt64(ctx JSContextRef, integer uint64, exception *JSValueRef) (JSValueRef, error) {
	if _jSBigIntCreateWithUInt64 == nil {
		return *new(JSValueRef), symbolCallError("JSBigIntCreateWithUInt64", "15.0", _jSBigIntCreateWithUInt64Err)
	}
	return _jSBigIntCreateWithUInt64(ctx, integer, exception), nil
}

// JSBigIntCreateWithUInt64.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSBigIntCreateWithUInt64(_:_:_:)
func JSBigIntCreateWithUInt64(ctx JSContextRef, integer uint64, exception *JSValueRef) JSValueRef {
	result, callErr := tryJSBigIntCreateWithUInt64(ctx, integer, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSCheckScriptSyntax func(ctx JSContextRef, script JSStringRef, sourceURL JSStringRef, startingLineNumber int, exception *JSValueRef) bool
var _jSCheckScriptSyntaxErr error

func tryJSCheckScriptSyntax(ctx JSContextRef, script JSStringRef, sourceURL JSStringRef, startingLineNumber int, exception *JSValueRef) (bool, error) {
	if _jSCheckScriptSyntax == nil {
		return false, symbolCallError("JSCheckScriptSyntax", "10.5", _jSCheckScriptSyntaxErr)
	}
	return _jSCheckScriptSyntax(ctx, script, sourceURL, startingLineNumber, exception), nil
}

// JSCheckScriptSyntax checks for syntax errors in a string of JavaScript.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSCheckScriptSyntax(_:_:_:_:_:)
func JSCheckScriptSyntax(ctx JSContextRef, script JSStringRef, sourceURL JSStringRef, startingLineNumber int, exception *JSValueRef) bool {
	result, callErr := tryJSCheckScriptSyntax(ctx, script, sourceURL, startingLineNumber, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSClassCreate func(definition *JSClassDefinition) JSClassRef
var _jSClassCreateErr error

func tryJSClassCreate(definition *JSClassDefinition) (JSClassRef, error) {
	if _jSClassCreate == nil {
		return *new(JSClassRef), symbolCallError("JSClassCreate", "10.5", _jSClassCreateErr)
	}
	return _jSClassCreate(definition), nil
}

// JSClassCreate creates a JavaScript class.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSClassCreate(_:)
func JSClassCreate(definition *JSClassDefinition) JSClassRef {
	result, callErr := tryJSClassCreate(definition)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSClassRelease func(jsClass JSClassRef)
var _jSClassReleaseErr error

func tryJSClassRelease(jsClass JSClassRef) error {
	if _jSClassRelease == nil {
		return symbolCallError("JSClassRelease", "10.5", _jSClassReleaseErr)
	}
	_jSClassRelease(jsClass)
	return nil
}

// JSClassRelease releases a JavaScript class.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSClassRelease(_:)
func JSClassRelease(jsClass JSClassRef) {
	if callErr := tryJSClassRelease(jsClass); callErr != nil {
		panic(callErr)
	}
}

var _jSClassRetain func(jsClass JSClassRef) JSClassRef
var _jSClassRetainErr error

func tryJSClassRetain(jsClass JSClassRef) (JSClassRef, error) {
	if _jSClassRetain == nil {
		return *new(JSClassRef), symbolCallError("JSClassRetain", "10.5", _jSClassRetainErr)
	}
	return _jSClassRetain(jsClass), nil
}

// JSClassRetain retains a JavaScript class.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSClassRetain(_:)
func JSClassRetain(jsClass JSClassRef) JSClassRef {
	result, callErr := tryJSClassRetain(jsClass)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSContextGetGlobalContext func(ctx JSContextRef) JSGlobalContextRef
var _jSContextGetGlobalContextErr error

func tryJSContextGetGlobalContext(ctx JSContextRef) (JSGlobalContextRef, error) {
	if _jSContextGetGlobalContext == nil {
		return *new(JSGlobalContextRef), symbolCallError("JSContextGetGlobalContext", "10.7", _jSContextGetGlobalContextErr)
	}
	return _jSContextGetGlobalContext(ctx), nil
}

// JSContextGetGlobalContext gets the global context of a JavaScript execution context.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContextGetGlobalContext(_:)
func JSContextGetGlobalContext(ctx JSContextRef) JSGlobalContextRef {
	result, callErr := tryJSContextGetGlobalContext(ctx)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSContextGetGlobalObject func(ctx JSContextRef) JSObjectRef
var _jSContextGetGlobalObjectErr error

func tryJSContextGetGlobalObject(ctx JSContextRef) (JSObjectRef, error) {
	if _jSContextGetGlobalObject == nil {
		return *new(JSObjectRef), symbolCallError("JSContextGetGlobalObject", "10.5", _jSContextGetGlobalObjectErr)
	}
	return _jSContextGetGlobalObject(ctx), nil
}

// JSContextGetGlobalObject gets the global object of a JavaScript execution context.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContextGetGlobalObject(_:)
func JSContextGetGlobalObject(ctx JSContextRef) JSObjectRef {
	result, callErr := tryJSContextGetGlobalObject(ctx)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSContextGetGroup func(ctx JSContextRef) JSContextGroupRef
var _jSContextGetGroupErr error

func tryJSContextGetGroup(ctx JSContextRef) (JSContextGroupRef, error) {
	if _jSContextGetGroup == nil {
		return *new(JSContextGroupRef), symbolCallError("JSContextGetGroup", "10.6", _jSContextGetGroupErr)
	}
	return _jSContextGetGroup(ctx), nil
}

// JSContextGetGroup gets the context group that a JavaScript execution context belongs to.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContextGetGroup(_:)
func JSContextGetGroup(ctx JSContextRef) JSContextGroupRef {
	result, callErr := tryJSContextGetGroup(ctx)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSContextGroupCreate func() JSContextGroupRef
var _jSContextGroupCreateErr error

func tryJSContextGroupCreate() (JSContextGroupRef, error) {
	if _jSContextGroupCreate == nil {
		return *new(JSContextGroupRef), symbolCallError("JSContextGroupCreate", "10.6", _jSContextGroupCreateErr)
	}
	return _jSContextGroupCreate(), nil
}

// JSContextGroupCreate creates a JavaScript context group.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContextGroupCreate()
func JSContextGroupCreate() JSContextGroupRef {
	result, callErr := tryJSContextGroupCreate()
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSContextGroupRelease func(group JSContextGroupRef)
var _jSContextGroupReleaseErr error

func tryJSContextGroupRelease(group JSContextGroupRef) error {
	if _jSContextGroupRelease == nil {
		return symbolCallError("JSContextGroupRelease", "10.6", _jSContextGroupReleaseErr)
	}
	_jSContextGroupRelease(group)
	return nil
}

// JSContextGroupRelease releases a JavaScript context group.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContextGroupRelease(_:)
func JSContextGroupRelease(group JSContextGroupRef) {
	if callErr := tryJSContextGroupRelease(group); callErr != nil {
		panic(callErr)
	}
}

var _jSContextGroupRetain func(group JSContextGroupRef) JSContextGroupRef
var _jSContextGroupRetainErr error

func tryJSContextGroupRetain(group JSContextGroupRef) (JSContextGroupRef, error) {
	if _jSContextGroupRetain == nil {
		return *new(JSContextGroupRef), symbolCallError("JSContextGroupRetain", "10.6", _jSContextGroupRetainErr)
	}
	return _jSContextGroupRetain(group), nil
}

// JSContextGroupRetain retains a JavaScript context group.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContextGroupRetain(_:)
func JSContextGroupRetain(group JSContextGroupRef) JSContextGroupRef {
	result, callErr := tryJSContextGroupRetain(group)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSEvaluateScript func(ctx JSContextRef, script JSStringRef, thisObject JSObjectRef, sourceURL JSStringRef, startingLineNumber int, exception *JSValueRef) JSValueRef
var _jSEvaluateScriptErr error

func tryJSEvaluateScript(ctx JSContextRef, script JSStringRef, thisObject JSObjectRef, sourceURL JSStringRef, startingLineNumber int, exception *JSValueRef) (JSValueRef, error) {
	if _jSEvaluateScript == nil {
		return *new(JSValueRef), symbolCallError("JSEvaluateScript", "10.5", _jSEvaluateScriptErr)
	}
	return _jSEvaluateScript(ctx, script, thisObject, sourceURL, startingLineNumber, exception), nil
}

// JSEvaluateScript evaluates a string of JavaScript.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSEvaluateScript(_:_:_:_:_:_:)
func JSEvaluateScript(ctx JSContextRef, script JSStringRef, thisObject JSObjectRef, sourceURL JSStringRef, startingLineNumber int, exception *JSValueRef) JSValueRef {
	result, callErr := tryJSEvaluateScript(ctx, script, thisObject, sourceURL, startingLineNumber, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSGarbageCollect func(ctx JSContextRef)
var _jSGarbageCollectErr error

func tryJSGarbageCollect(ctx JSContextRef) error {
	if _jSGarbageCollect == nil {
		return symbolCallError("JSGarbageCollect", "10.5", _jSGarbageCollectErr)
	}
	_jSGarbageCollect(ctx)
	return nil
}

// JSGarbageCollect performs a JavaScript garbage collection.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSGarbageCollect(_:)
func JSGarbageCollect(ctx JSContextRef) {
	if callErr := tryJSGarbageCollect(ctx); callErr != nil {
		panic(callErr)
	}
}

var _jSGlobalContextCopyName func(ctx JSGlobalContextRef) JSStringRef
var _jSGlobalContextCopyNameErr error

func tryJSGlobalContextCopyName(ctx JSGlobalContextRef) (JSStringRef, error) {
	if _jSGlobalContextCopyName == nil {
		return *new(JSStringRef), symbolCallError("JSGlobalContextCopyName", "10.10", _jSGlobalContextCopyNameErr)
	}
	return _jSGlobalContextCopyName(ctx), nil
}

// JSGlobalContextCopyName gets a copy of the name of a context.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextCopyName(_:)
func JSGlobalContextCopyName(ctx JSGlobalContextRef) JSStringRef {
	result, callErr := tryJSGlobalContextCopyName(ctx)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSGlobalContextCreate func(globalObjectClass JSClassRef) JSGlobalContextRef
var _jSGlobalContextCreateErr error

func tryJSGlobalContextCreate(globalObjectClass JSClassRef) (JSGlobalContextRef, error) {
	if _jSGlobalContextCreate == nil {
		return *new(JSGlobalContextRef), symbolCallError("JSGlobalContextCreate", "10.5", _jSGlobalContextCreateErr)
	}
	return _jSGlobalContextCreate(globalObjectClass), nil
}

// JSGlobalContextCreate creates a global JavaScript execution context.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextCreate(_:)
func JSGlobalContextCreate(globalObjectClass JSClassRef) JSGlobalContextRef {
	result, callErr := tryJSGlobalContextCreate(globalObjectClass)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSGlobalContextCreateInGroup func(group JSContextGroupRef, globalObjectClass JSClassRef) JSGlobalContextRef
var _jSGlobalContextCreateInGroupErr error

func tryJSGlobalContextCreateInGroup(group JSContextGroupRef, globalObjectClass JSClassRef) (JSGlobalContextRef, error) {
	if _jSGlobalContextCreateInGroup == nil {
		return *new(JSGlobalContextRef), symbolCallError("JSGlobalContextCreateInGroup", "10.6", _jSGlobalContextCreateInGroupErr)
	}
	return _jSGlobalContextCreateInGroup(group, globalObjectClass), nil
}

// JSGlobalContextCreateInGroup creates a global JavaScript execution context in the provided context group.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextCreateInGroup(_:_:)
func JSGlobalContextCreateInGroup(group JSContextGroupRef, globalObjectClass JSClassRef) JSGlobalContextRef {
	result, callErr := tryJSGlobalContextCreateInGroup(group, globalObjectClass)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSGlobalContextIsInspectable func(ctx JSGlobalContextRef) bool
var _jSGlobalContextIsInspectableErr error

func tryJSGlobalContextIsInspectable(ctx JSGlobalContextRef) (bool, error) {
	if _jSGlobalContextIsInspectable == nil {
		return false, symbolCallError("JSGlobalContextIsInspectable", "13.3", _jSGlobalContextIsInspectableErr)
	}
	return _jSGlobalContextIsInspectable(ctx), nil
}

// JSGlobalContextIsInspectable returns a Boolean value that indicates whether the JavaScript context is inspectable.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextIsInspectable(_:)
func JSGlobalContextIsInspectable(ctx JSGlobalContextRef) bool {
	result, callErr := tryJSGlobalContextIsInspectable(ctx)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSGlobalContextRelease func(ctx JSGlobalContextRef)
var _jSGlobalContextReleaseErr error

func tryJSGlobalContextRelease(ctx JSGlobalContextRef) error {
	if _jSGlobalContextRelease == nil {
		return symbolCallError("JSGlobalContextRelease", "10.5", _jSGlobalContextReleaseErr)
	}
	_jSGlobalContextRelease(ctx)
	return nil
}

// JSGlobalContextRelease releases a global JavaScript execution context.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextRelease(_:)
func JSGlobalContextRelease(ctx JSGlobalContextRef) {
	if callErr := tryJSGlobalContextRelease(ctx); callErr != nil {
		panic(callErr)
	}
}

var _jSGlobalContextRetain func(ctx JSGlobalContextRef) JSGlobalContextRef
var _jSGlobalContextRetainErr error

func tryJSGlobalContextRetain(ctx JSGlobalContextRef) (JSGlobalContextRef, error) {
	if _jSGlobalContextRetain == nil {
		return *new(JSGlobalContextRef), symbolCallError("JSGlobalContextRetain", "10.5", _jSGlobalContextRetainErr)
	}
	return _jSGlobalContextRetain(ctx), nil
}

// JSGlobalContextRetain retains a global JavaScript execution context.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextRetain(_:)
func JSGlobalContextRetain(ctx JSGlobalContextRef) JSGlobalContextRef {
	result, callErr := tryJSGlobalContextRetain(ctx)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSGlobalContextSetInspectable func(ctx JSGlobalContextRef, inspectable bool)
var _jSGlobalContextSetInspectableErr error

func tryJSGlobalContextSetInspectable(ctx JSGlobalContextRef, inspectable bool) error {
	if _jSGlobalContextSetInspectable == nil {
		return symbolCallError("JSGlobalContextSetInspectable", "13.3", _jSGlobalContextSetInspectableErr)
	}
	_jSGlobalContextSetInspectable(ctx, inspectable)
	return nil
}

// JSGlobalContextSetInspectable sets a JavaScript context to be either inspectable or not inspectable.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextSetInspectable(_:_:)
func JSGlobalContextSetInspectable(ctx JSGlobalContextRef, inspectable bool) {
	if callErr := tryJSGlobalContextSetInspectable(ctx, inspectable); callErr != nil {
		panic(callErr)
	}
}

var _jSGlobalContextSetName func(ctx JSGlobalContextRef, name JSStringRef)
var _jSGlobalContextSetNameErr error

func tryJSGlobalContextSetName(ctx JSGlobalContextRef, name JSStringRef) error {
	if _jSGlobalContextSetName == nil {
		return symbolCallError("JSGlobalContextSetName", "10.10", _jSGlobalContextSetNameErr)
	}
	_jSGlobalContextSetName(ctx, name)
	return nil
}

// JSGlobalContextSetName sets the remote debugging name for a context.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSGlobalContextSetName(_:_:)
func JSGlobalContextSetName(ctx JSGlobalContextRef, name JSStringRef) {
	if callErr := tryJSGlobalContextSetName(ctx, name); callErr != nil {
		panic(callErr)
	}
}

var _jSObjectCallAsConstructor func(ctx JSContextRef, object JSObjectRef, argumentCount uintptr, arguments JSValueRef, exception *JSValueRef) JSObjectRef
var _jSObjectCallAsConstructorErr error

func tryJSObjectCallAsConstructor(ctx JSContextRef, object JSObjectRef, argumentCount uintptr, arguments JSValueRef, exception *JSValueRef) (JSObjectRef, error) {
	if _jSObjectCallAsConstructor == nil {
		return *new(JSObjectRef), symbolCallError("JSObjectCallAsConstructor", "10.5", _jSObjectCallAsConstructorErr)
	}
	return _jSObjectCallAsConstructor(ctx, object, argumentCount, arguments, exception), nil
}

// JSObjectCallAsConstructor calls an object as a constructor.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectCallAsConstructor(_:_:_:_:_:)
func JSObjectCallAsConstructor(ctx JSContextRef, object JSObjectRef, argumentCount uintptr, arguments JSValueRef, exception *JSValueRef) JSObjectRef {
	result, callErr := tryJSObjectCallAsConstructor(ctx, object, argumentCount, arguments, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectCallAsFunction func(ctx JSContextRef, object JSObjectRef, thisObject JSObjectRef, argumentCount uintptr, arguments JSValueRef, exception *JSValueRef) JSValueRef
var _jSObjectCallAsFunctionErr error

func tryJSObjectCallAsFunction(ctx JSContextRef, object JSObjectRef, thisObject JSObjectRef, argumentCount uintptr, arguments JSValueRef, exception *JSValueRef) (JSValueRef, error) {
	if _jSObjectCallAsFunction == nil {
		return *new(JSValueRef), symbolCallError("JSObjectCallAsFunction", "10.5", _jSObjectCallAsFunctionErr)
	}
	return _jSObjectCallAsFunction(ctx, object, thisObject, argumentCount, arguments, exception), nil
}

// JSObjectCallAsFunction calls an object as a function.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectCallAsFunction(_:_:_:_:_:_:)
func JSObjectCallAsFunction(ctx JSContextRef, object JSObjectRef, thisObject JSObjectRef, argumentCount uintptr, arguments JSValueRef, exception *JSValueRef) JSValueRef {
	result, callErr := tryJSObjectCallAsFunction(ctx, object, thisObject, argumentCount, arguments, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectCopyPropertyNames func(ctx JSContextRef, object JSObjectRef) JSPropertyNameArrayRef
var _jSObjectCopyPropertyNamesErr error

func tryJSObjectCopyPropertyNames(ctx JSContextRef, object JSObjectRef) (JSPropertyNameArrayRef, error) {
	if _jSObjectCopyPropertyNames == nil {
		return *new(JSPropertyNameArrayRef), symbolCallError("JSObjectCopyPropertyNames", "10.5", _jSObjectCopyPropertyNamesErr)
	}
	return _jSObjectCopyPropertyNames(ctx, object), nil
}

// JSObjectCopyPropertyNames gets the names of an object’s enumerable properties.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectCopyPropertyNames(_:_:)
func JSObjectCopyPropertyNames(ctx JSContextRef, object JSObjectRef) JSPropertyNameArrayRef {
	result, callErr := tryJSObjectCopyPropertyNames(ctx, object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectDeleteProperty func(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef, exception *JSValueRef) bool
var _jSObjectDeletePropertyErr error

func tryJSObjectDeleteProperty(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef, exception *JSValueRef) (bool, error) {
	if _jSObjectDeleteProperty == nil {
		return false, symbolCallError("JSObjectDeleteProperty", "10.5", _jSObjectDeletePropertyErr)
	}
	return _jSObjectDeleteProperty(ctx, object, propertyName, exception), nil
}

// JSObjectDeleteProperty deletes a property from an object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectDeleteProperty(_:_:_:_:)
func JSObjectDeleteProperty(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef, exception *JSValueRef) bool {
	result, callErr := tryJSObjectDeleteProperty(ctx, object, propertyName, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectDeletePropertyForKey func(ctx JSContextRef, object JSObjectRef, propertyKey JSValueRef, exception *JSValueRef) bool
var _jSObjectDeletePropertyForKeyErr error

func tryJSObjectDeletePropertyForKey(ctx JSContextRef, object JSObjectRef, propertyKey JSValueRef, exception *JSValueRef) (bool, error) {
	if _jSObjectDeletePropertyForKey == nil {
		return false, symbolCallError("JSObjectDeletePropertyForKey", "10.15", _jSObjectDeletePropertyForKeyErr)
	}
	return _jSObjectDeletePropertyForKey(ctx, object, propertyKey, exception), nil
}

// JSObjectDeletePropertyForKey deletes a property from an object using a JavaScript value as the property key.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectDeletePropertyForKey(_:_:_:_:)
func JSObjectDeletePropertyForKey(ctx JSContextRef, object JSObjectRef, propertyKey JSValueRef, exception *JSValueRef) bool {
	result, callErr := tryJSObjectDeletePropertyForKey(ctx, object, propertyKey, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectGetArrayBufferByteLength func(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) uintptr
var _jSObjectGetArrayBufferByteLengthErr error

func tryJSObjectGetArrayBufferByteLength(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) (uintptr, error) {
	if _jSObjectGetArrayBufferByteLength == nil {
		return 0, symbolCallError("JSObjectGetArrayBufferByteLength", "10.12", _jSObjectGetArrayBufferByteLengthErr)
	}
	return _jSObjectGetArrayBufferByteLength(ctx, object, exception), nil
}

// JSObjectGetArrayBufferByteLength returns the number of bytes in a JavaScript data object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetArrayBufferByteLength(_:_:_:)
func JSObjectGetArrayBufferByteLength(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) uintptr {
	result, callErr := tryJSObjectGetArrayBufferByteLength(ctx, object, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectGetArrayBufferBytesPtr func(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) unsafe.Pointer
var _jSObjectGetArrayBufferBytesPtrErr error

func tryJSObjectGetArrayBufferBytesPtr(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) (unsafe.Pointer, error) {
	if _jSObjectGetArrayBufferBytesPtr == nil {
		return nil, symbolCallError("JSObjectGetArrayBufferBytesPtr", "10.12", _jSObjectGetArrayBufferBytesPtrErr)
	}
	return _jSObjectGetArrayBufferBytesPtr(ctx, object, exception), nil
}

// JSObjectGetArrayBufferBytesPtr returns a pointer to the data buffer that serves as the backing store for a JavaScript typed array object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetArrayBufferBytesPtr(_:_:_:)
func JSObjectGetArrayBufferBytesPtr(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) unsafe.Pointer {
	result, callErr := tryJSObjectGetArrayBufferBytesPtr(ctx, object, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectGetPrivate func(object JSObjectRef) unsafe.Pointer
var _jSObjectGetPrivateErr error

func tryJSObjectGetPrivate(object JSObjectRef) (unsafe.Pointer, error) {
	if _jSObjectGetPrivate == nil {
		return nil, symbolCallError("JSObjectGetPrivate", "10.5", _jSObjectGetPrivateErr)
	}
	return _jSObjectGetPrivate(object), nil
}

// JSObjectGetPrivate gets an object’s private data.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetPrivate(_:)
func JSObjectGetPrivate(object JSObjectRef) unsafe.Pointer {
	result, callErr := tryJSObjectGetPrivate(object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectGetProperty func(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef, exception *JSValueRef) JSValueRef
var _jSObjectGetPropertyErr error

func tryJSObjectGetProperty(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef, exception *JSValueRef) (JSValueRef, error) {
	if _jSObjectGetProperty == nil {
		return *new(JSValueRef), symbolCallError("JSObjectGetProperty", "10.5", _jSObjectGetPropertyErr)
	}
	return _jSObjectGetProperty(ctx, object, propertyName, exception), nil
}

// JSObjectGetProperty gets a property from an object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetProperty(_:_:_:_:)
func JSObjectGetProperty(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef, exception *JSValueRef) JSValueRef {
	result, callErr := tryJSObjectGetProperty(ctx, object, propertyName, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectGetPropertyAtIndex func(ctx JSContextRef, object JSObjectRef, propertyIndex uint, exception *JSValueRef) JSValueRef
var _jSObjectGetPropertyAtIndexErr error

func tryJSObjectGetPropertyAtIndex(ctx JSContextRef, object JSObjectRef, propertyIndex uint, exception *JSValueRef) (JSValueRef, error) {
	if _jSObjectGetPropertyAtIndex == nil {
		return *new(JSValueRef), symbolCallError("JSObjectGetPropertyAtIndex", "10.5", _jSObjectGetPropertyAtIndexErr)
	}
	return _jSObjectGetPropertyAtIndex(ctx, object, propertyIndex, exception), nil
}

// JSObjectGetPropertyAtIndex gets a property from an object by numeric index.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetPropertyAtIndex(_:_:_:_:)
func JSObjectGetPropertyAtIndex(ctx JSContextRef, object JSObjectRef, propertyIndex uint, exception *JSValueRef) JSValueRef {
	result, callErr := tryJSObjectGetPropertyAtIndex(ctx, object, propertyIndex, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectGetPropertyForKey func(ctx JSContextRef, object JSObjectRef, propertyKey JSValueRef, exception *JSValueRef) JSValueRef
var _jSObjectGetPropertyForKeyErr error

func tryJSObjectGetPropertyForKey(ctx JSContextRef, object JSObjectRef, propertyKey JSValueRef, exception *JSValueRef) (JSValueRef, error) {
	if _jSObjectGetPropertyForKey == nil {
		return *new(JSValueRef), symbolCallError("JSObjectGetPropertyForKey", "10.15", _jSObjectGetPropertyForKeyErr)
	}
	return _jSObjectGetPropertyForKey(ctx, object, propertyKey, exception), nil
}

// JSObjectGetPropertyForKey gets a property from an object using a JavaScript value as the property key.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetPropertyForKey(_:_:_:_:)
func JSObjectGetPropertyForKey(ctx JSContextRef, object JSObjectRef, propertyKey JSValueRef, exception *JSValueRef) JSValueRef {
	result, callErr := tryJSObjectGetPropertyForKey(ctx, object, propertyKey, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectGetPrototype func(ctx JSContextRef, object JSObjectRef) JSValueRef
var _jSObjectGetPrototypeErr error

func tryJSObjectGetPrototype(ctx JSContextRef, object JSObjectRef) (JSValueRef, error) {
	if _jSObjectGetPrototype == nil {
		return *new(JSValueRef), symbolCallError("JSObjectGetPrototype", "10.5", _jSObjectGetPrototypeErr)
	}
	return _jSObjectGetPrototype(ctx, object), nil
}

// JSObjectGetPrototype gets an object’s prototype.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetPrototype(_:_:)
func JSObjectGetPrototype(ctx JSContextRef, object JSObjectRef) JSValueRef {
	result, callErr := tryJSObjectGetPrototype(ctx, object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectGetTypedArrayBuffer func(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) JSObjectRef
var _jSObjectGetTypedArrayBufferErr error

func tryJSObjectGetTypedArrayBuffer(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) (JSObjectRef, error) {
	if _jSObjectGetTypedArrayBuffer == nil {
		return *new(JSObjectRef), symbolCallError("JSObjectGetTypedArrayBuffer", "10.12", _jSObjectGetTypedArrayBufferErr)
	}
	return _jSObjectGetTypedArrayBuffer(ctx, object, exception), nil
}

// JSObjectGetTypedArrayBuffer returns the JavaScript array buffer object to use as the backing of a JavaScript typed array object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetTypedArrayBuffer(_:_:_:)
func JSObjectGetTypedArrayBuffer(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) JSObjectRef {
	result, callErr := tryJSObjectGetTypedArrayBuffer(ctx, object, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectGetTypedArrayByteLength func(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) uintptr
var _jSObjectGetTypedArrayByteLengthErr error

func tryJSObjectGetTypedArrayByteLength(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) (uintptr, error) {
	if _jSObjectGetTypedArrayByteLength == nil {
		return 0, symbolCallError("JSObjectGetTypedArrayByteLength", "10.12", _jSObjectGetTypedArrayByteLengthErr)
	}
	return _jSObjectGetTypedArrayByteLength(ctx, object, exception), nil
}

// JSObjectGetTypedArrayByteLength returns the byte length of a JavaScript typed array object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetTypedArrayByteLength(_:_:_:)
func JSObjectGetTypedArrayByteLength(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) uintptr {
	result, callErr := tryJSObjectGetTypedArrayByteLength(ctx, object, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectGetTypedArrayByteOffset func(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) uintptr
var _jSObjectGetTypedArrayByteOffsetErr error

func tryJSObjectGetTypedArrayByteOffset(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) (uintptr, error) {
	if _jSObjectGetTypedArrayByteOffset == nil {
		return 0, symbolCallError("JSObjectGetTypedArrayByteOffset", "10.12", _jSObjectGetTypedArrayByteOffsetErr)
	}
	return _jSObjectGetTypedArrayByteOffset(ctx, object, exception), nil
}

// JSObjectGetTypedArrayByteOffset returns the byte offset of a JavaScript typed array object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetTypedArrayByteOffset(_:_:_:)
func JSObjectGetTypedArrayByteOffset(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) uintptr {
	result, callErr := tryJSObjectGetTypedArrayByteOffset(ctx, object, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectGetTypedArrayBytesPtr func(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) unsafe.Pointer
var _jSObjectGetTypedArrayBytesPtrErr error

func tryJSObjectGetTypedArrayBytesPtr(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) (unsafe.Pointer, error) {
	if _jSObjectGetTypedArrayBytesPtr == nil {
		return nil, symbolCallError("JSObjectGetTypedArrayBytesPtr", "10.12", _jSObjectGetTypedArrayBytesPtrErr)
	}
	return _jSObjectGetTypedArrayBytesPtr(ctx, object, exception), nil
}

// JSObjectGetTypedArrayBytesPtr returns a temporary pointer to the backing store of a JavaScript typed array object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetTypedArrayBytesPtr(_:_:_:)
func JSObjectGetTypedArrayBytesPtr(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) unsafe.Pointer {
	result, callErr := tryJSObjectGetTypedArrayBytesPtr(ctx, object, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectGetTypedArrayLength func(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) uintptr
var _jSObjectGetTypedArrayLengthErr error

func tryJSObjectGetTypedArrayLength(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) (uintptr, error) {
	if _jSObjectGetTypedArrayLength == nil {
		return 0, symbolCallError("JSObjectGetTypedArrayLength", "10.12", _jSObjectGetTypedArrayLengthErr)
	}
	return _jSObjectGetTypedArrayLength(ctx, object, exception), nil
}

// JSObjectGetTypedArrayLength returns the length of a JavaScript typed array object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectGetTypedArrayLength(_:_:_:)
func JSObjectGetTypedArrayLength(ctx JSContextRef, object JSObjectRef, exception *JSValueRef) uintptr {
	result, callErr := tryJSObjectGetTypedArrayLength(ctx, object, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectHasProperty func(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef) bool
var _jSObjectHasPropertyErr error

func tryJSObjectHasProperty(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef) (bool, error) {
	if _jSObjectHasProperty == nil {
		return false, symbolCallError("JSObjectHasProperty", "10.5", _jSObjectHasPropertyErr)
	}
	return _jSObjectHasProperty(ctx, object, propertyName), nil
}

// JSObjectHasProperty tests whether an object has a specified property.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectHasProperty(_:_:_:)
func JSObjectHasProperty(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef) bool {
	result, callErr := tryJSObjectHasProperty(ctx, object, propertyName)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectHasPropertyForKey func(ctx JSContextRef, object JSObjectRef, propertyKey JSValueRef, exception *JSValueRef) bool
var _jSObjectHasPropertyForKeyErr error

func tryJSObjectHasPropertyForKey(ctx JSContextRef, object JSObjectRef, propertyKey JSValueRef, exception *JSValueRef) (bool, error) {
	if _jSObjectHasPropertyForKey == nil {
		return false, symbolCallError("JSObjectHasPropertyForKey", "10.15", _jSObjectHasPropertyForKeyErr)
	}
	return _jSObjectHasPropertyForKey(ctx, object, propertyKey, exception), nil
}

// JSObjectHasPropertyForKey tests whether an object has the specified property using a JavaScript value as the property key.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectHasPropertyForKey(_:_:_:_:)
func JSObjectHasPropertyForKey(ctx JSContextRef, object JSObjectRef, propertyKey JSValueRef, exception *JSValueRef) bool {
	result, callErr := tryJSObjectHasPropertyForKey(ctx, object, propertyKey, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectIsConstructor func(ctx JSContextRef, object JSObjectRef) bool
var _jSObjectIsConstructorErr error

func tryJSObjectIsConstructor(ctx JSContextRef, object JSObjectRef) (bool, error) {
	if _jSObjectIsConstructor == nil {
		return false, symbolCallError("JSObjectIsConstructor", "10.5", _jSObjectIsConstructorErr)
	}
	return _jSObjectIsConstructor(ctx, object), nil
}

// JSObjectIsConstructor tests whether you can call an object as a constructor.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectIsConstructor(_:_:)
func JSObjectIsConstructor(ctx JSContextRef, object JSObjectRef) bool {
	result, callErr := tryJSObjectIsConstructor(ctx, object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectIsFunction func(ctx JSContextRef, object JSObjectRef) bool
var _jSObjectIsFunctionErr error

func tryJSObjectIsFunction(ctx JSContextRef, object JSObjectRef) (bool, error) {
	if _jSObjectIsFunction == nil {
		return false, symbolCallError("JSObjectIsFunction", "10.5", _jSObjectIsFunctionErr)
	}
	return _jSObjectIsFunction(ctx, object), nil
}

// JSObjectIsFunction tests whether you can call an object as a function.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectIsFunction(_:_:)
func JSObjectIsFunction(ctx JSContextRef, object JSObjectRef) bool {
	result, callErr := tryJSObjectIsFunction(ctx, object)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectMake func(ctx JSContextRef, jsClass JSClassRef, data unsafe.Pointer) JSObjectRef
var _jSObjectMakeErr error

func tryJSObjectMake(ctx JSContextRef, jsClass JSClassRef, data unsafe.Pointer) (JSObjectRef, error) {
	if _jSObjectMake == nil {
		return *new(JSObjectRef), symbolCallError("JSObjectMake", "10.5", _jSObjectMakeErr)
	}
	return _jSObjectMake(ctx, jsClass, data), nil
}

// JSObjectMake creates a JavaScript object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMake(_:_:_:)
func JSObjectMake(ctx JSContextRef, jsClass JSClassRef, data unsafe.Pointer) JSObjectRef {
	result, callErr := tryJSObjectMake(ctx, jsClass, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectMakeArray func(ctx JSContextRef, argumentCount uintptr, arguments JSValueRef, exception *JSValueRef) JSObjectRef
var _jSObjectMakeArrayErr error

func tryJSObjectMakeArray(ctx JSContextRef, argumentCount uintptr, arguments JSValueRef, exception *JSValueRef) (JSObjectRef, error) {
	if _jSObjectMakeArray == nil {
		return *new(JSObjectRef), symbolCallError("JSObjectMakeArray", "10.6", _jSObjectMakeArrayErr)
	}
	return _jSObjectMakeArray(ctx, argumentCount, arguments, exception), nil
}

// JSObjectMakeArray creates a JavaScript array object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeArray(_:_:_:_:)
func JSObjectMakeArray(ctx JSContextRef, argumentCount uintptr, arguments JSValueRef, exception *JSValueRef) JSObjectRef {
	result, callErr := tryJSObjectMakeArray(ctx, argumentCount, arguments, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectMakeArrayBufferWithBytesNoCopy func(ctx JSContextRef, bytes unsafe.Pointer, byteLength uintptr, bytesDeallocator JSTypedArrayBytesDeallocator, deallocatorContext unsafe.Pointer, exception *JSValueRef) JSObjectRef
var _jSObjectMakeArrayBufferWithBytesNoCopyErr error

func tryJSObjectMakeArrayBufferWithBytesNoCopy(ctx JSContextRef, bytes unsafe.Pointer, byteLength uintptr, bytesDeallocator JSTypedArrayBytesDeallocator, deallocatorContext unsafe.Pointer, exception *JSValueRef) (JSObjectRef, error) {
	if _jSObjectMakeArrayBufferWithBytesNoCopy == nil {
		return *new(JSObjectRef), symbolCallError("JSObjectMakeArrayBufferWithBytesNoCopy", "10.12", _jSObjectMakeArrayBufferWithBytesNoCopyErr)
	}
	return _jSObjectMakeArrayBufferWithBytesNoCopy(ctx, bytes, byteLength, bytesDeallocator, deallocatorContext, exception), nil
}

// JSObjectMakeArrayBufferWithBytesNoCopy creates a JavaScript array buffer object from an existing pointer.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeArrayBufferWithBytesNoCopy(_:_:_:_:_:_:)
func JSObjectMakeArrayBufferWithBytesNoCopy(ctx JSContextRef, bytes unsafe.Pointer, byteLength uintptr, bytesDeallocator JSTypedArrayBytesDeallocator, deallocatorContext unsafe.Pointer, exception *JSValueRef) JSObjectRef {
	result, callErr := tryJSObjectMakeArrayBufferWithBytesNoCopy(ctx, bytes, byteLength, bytesDeallocator, deallocatorContext, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectMakeConstructor func(ctx JSContextRef, jsClass JSClassRef, callAsConstructor JSObjectCallAsConstructorCallback) JSObjectRef
var _jSObjectMakeConstructorErr error

func tryJSObjectMakeConstructor(ctx JSContextRef, jsClass JSClassRef, callAsConstructor JSObjectCallAsConstructorCallback) (JSObjectRef, error) {
	if _jSObjectMakeConstructor == nil {
		return *new(JSObjectRef), symbolCallError("JSObjectMakeConstructor", "10.5", _jSObjectMakeConstructorErr)
	}
	return _jSObjectMakeConstructor(ctx, jsClass, callAsConstructor), nil
}

// JSObjectMakeConstructor creates a JavaScript constructor.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeConstructor(_:_:_:)
func JSObjectMakeConstructor(ctx JSContextRef, jsClass JSClassRef, callAsConstructor JSObjectCallAsConstructorCallback) JSObjectRef {
	result, callErr := tryJSObjectMakeConstructor(ctx, jsClass, callAsConstructor)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectMakeDate func(ctx JSContextRef, argumentCount uintptr, arguments JSValueRef, exception *JSValueRef) JSObjectRef
var _jSObjectMakeDateErr error

func tryJSObjectMakeDate(ctx JSContextRef, argumentCount uintptr, arguments JSValueRef, exception *JSValueRef) (JSObjectRef, error) {
	if _jSObjectMakeDate == nil {
		return *new(JSObjectRef), symbolCallError("JSObjectMakeDate", "10.6", _jSObjectMakeDateErr)
	}
	return _jSObjectMakeDate(ctx, argumentCount, arguments, exception), nil
}

// JSObjectMakeDate creates a JavaScript date object as though invoking the built-in date constructor.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeDate(_:_:_:_:)
func JSObjectMakeDate(ctx JSContextRef, argumentCount uintptr, arguments JSValueRef, exception *JSValueRef) JSObjectRef {
	result, callErr := tryJSObjectMakeDate(ctx, argumentCount, arguments, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectMakeDeferredPromise func(ctx JSContextRef, resolve *JSObjectRef, reject *JSObjectRef, exception *JSValueRef) JSObjectRef
var _jSObjectMakeDeferredPromiseErr error

func tryJSObjectMakeDeferredPromise(ctx JSContextRef, resolve *JSObjectRef, reject *JSObjectRef, exception *JSValueRef) (JSObjectRef, error) {
	if _jSObjectMakeDeferredPromise == nil {
		return *new(JSObjectRef), symbolCallError("JSObjectMakeDeferredPromise", "10.15", _jSObjectMakeDeferredPromiseErr)
	}
	return _jSObjectMakeDeferredPromise(ctx, resolve, reject, exception), nil
}

// JSObjectMakeDeferredPromise creates a JavaScript promise object by invoking the provided executor.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeDeferredPromise(_:_:_:_:)
func JSObjectMakeDeferredPromise(ctx JSContextRef, resolve *JSObjectRef, reject *JSObjectRef, exception *JSValueRef) JSObjectRef {
	result, callErr := tryJSObjectMakeDeferredPromise(ctx, resolve, reject, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectMakeError func(ctx JSContextRef, argumentCount uintptr, arguments JSValueRef, exception *JSValueRef) JSObjectRef
var _jSObjectMakeErrorErr error

func tryJSObjectMakeError(ctx JSContextRef, argumentCount uintptr, arguments JSValueRef, exception *JSValueRef) (JSObjectRef, error) {
	if _jSObjectMakeError == nil {
		return *new(JSObjectRef), symbolCallError("JSObjectMakeError", "10.6", _jSObjectMakeErrorErr)
	}
	return _jSObjectMakeError(ctx, argumentCount, arguments, exception), nil
}

// JSObjectMakeError creates a JavaScript error object as though invoking the built-in error constructor.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeError(_:_:_:_:)
func JSObjectMakeError(ctx JSContextRef, argumentCount uintptr, arguments JSValueRef, exception *JSValueRef) JSObjectRef {
	result, callErr := tryJSObjectMakeError(ctx, argumentCount, arguments, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectMakeFunction func(ctx JSContextRef, name JSStringRef, parameterCount uint, parameterNames JSStringRef, body JSStringRef, sourceURL JSStringRef, startingLineNumber int, exception *JSValueRef) JSObjectRef
var _jSObjectMakeFunctionErr error

func tryJSObjectMakeFunction(ctx JSContextRef, name JSStringRef, parameterCount uint, parameterNames JSStringRef, body JSStringRef, sourceURL JSStringRef, startingLineNumber int, exception *JSValueRef) (JSObjectRef, error) {
	if _jSObjectMakeFunction == nil {
		return *new(JSObjectRef), symbolCallError("JSObjectMakeFunction", "10.5", _jSObjectMakeFunctionErr)
	}
	return _jSObjectMakeFunction(ctx, name, parameterCount, parameterNames, body, sourceURL, startingLineNumber, exception), nil
}

// JSObjectMakeFunction creates a function with a specified script as its body.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeFunction(_:_:_:_:_:_:_:_:)
func JSObjectMakeFunction(ctx JSContextRef, name JSStringRef, parameterCount uint, parameterNames JSStringRef, body JSStringRef, sourceURL JSStringRef, startingLineNumber int, exception *JSValueRef) JSObjectRef {
	result, callErr := tryJSObjectMakeFunction(ctx, name, parameterCount, parameterNames, body, sourceURL, startingLineNumber, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectMakeFunctionWithCallback func(ctx JSContextRef, name JSStringRef, callAsFunction JSObjectCallAsFunctionCallback) JSObjectRef
var _jSObjectMakeFunctionWithCallbackErr error

func tryJSObjectMakeFunctionWithCallback(ctx JSContextRef, name JSStringRef, callAsFunction JSObjectCallAsFunctionCallback) (JSObjectRef, error) {
	if _jSObjectMakeFunctionWithCallback == nil {
		return *new(JSObjectRef), symbolCallError("JSObjectMakeFunctionWithCallback", "10.5", _jSObjectMakeFunctionWithCallbackErr)
	}
	return _jSObjectMakeFunctionWithCallback(ctx, name, callAsFunction), nil
}

// JSObjectMakeFunctionWithCallback creates a JavaScript function with a specified callback as its implementation.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeFunctionWithCallback(_:_:_:)
func JSObjectMakeFunctionWithCallback(ctx JSContextRef, name JSStringRef, callAsFunction JSObjectCallAsFunctionCallback) JSObjectRef {
	result, callErr := tryJSObjectMakeFunctionWithCallback(ctx, name, callAsFunction)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectMakeRegExp func(ctx JSContextRef, argumentCount uintptr, arguments JSValueRef, exception *JSValueRef) JSObjectRef
var _jSObjectMakeRegExpErr error

func tryJSObjectMakeRegExp(ctx JSContextRef, argumentCount uintptr, arguments JSValueRef, exception *JSValueRef) (JSObjectRef, error) {
	if _jSObjectMakeRegExp == nil {
		return *new(JSObjectRef), symbolCallError("JSObjectMakeRegExp", "10.6", _jSObjectMakeRegExpErr)
	}
	return _jSObjectMakeRegExp(ctx, argumentCount, arguments, exception), nil
}

// JSObjectMakeRegExp creates a JavaScript regular expression object as though invoking the built-in regular expression constructor.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeRegExp(_:_:_:_:)
func JSObjectMakeRegExp(ctx JSContextRef, argumentCount uintptr, arguments JSValueRef, exception *JSValueRef) JSObjectRef {
	result, callErr := tryJSObjectMakeRegExp(ctx, argumentCount, arguments, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectMakeTypedArray func(ctx JSContextRef, arrayType JSTypedArrayType, length uintptr, exception *JSValueRef) JSObjectRef
var _jSObjectMakeTypedArrayErr error

func tryJSObjectMakeTypedArray(ctx JSContextRef, arrayType JSTypedArrayType, length uintptr, exception *JSValueRef) (JSObjectRef, error) {
	if _jSObjectMakeTypedArray == nil {
		return *new(JSObjectRef), symbolCallError("JSObjectMakeTypedArray", "10.12", _jSObjectMakeTypedArrayErr)
	}
	return _jSObjectMakeTypedArray(ctx, arrayType, length, exception), nil
}

// JSObjectMakeTypedArray creates a JavaScript typed array object with the specified number of elements.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeTypedArray(_:_:_:_:)
func JSObjectMakeTypedArray(ctx JSContextRef, arrayType JSTypedArrayType, length uintptr, exception *JSValueRef) JSObjectRef {
	result, callErr := tryJSObjectMakeTypedArray(ctx, arrayType, length, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectMakeTypedArrayWithArrayBuffer func(ctx JSContextRef, arrayType JSTypedArrayType, buffer JSObjectRef, exception *JSValueRef) JSObjectRef
var _jSObjectMakeTypedArrayWithArrayBufferErr error

func tryJSObjectMakeTypedArrayWithArrayBuffer(ctx JSContextRef, arrayType JSTypedArrayType, buffer JSObjectRef, exception *JSValueRef) (JSObjectRef, error) {
	if _jSObjectMakeTypedArrayWithArrayBuffer == nil {
		return *new(JSObjectRef), symbolCallError("JSObjectMakeTypedArrayWithArrayBuffer", "10.12", _jSObjectMakeTypedArrayWithArrayBufferErr)
	}
	return _jSObjectMakeTypedArrayWithArrayBuffer(ctx, arrayType, buffer, exception), nil
}

// JSObjectMakeTypedArrayWithArrayBuffer creates a JavaScript typed array object from an existing JavaScript array buffer object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeTypedArrayWithArrayBuffer(_:_:_:_:)
func JSObjectMakeTypedArrayWithArrayBuffer(ctx JSContextRef, arrayType JSTypedArrayType, buffer JSObjectRef, exception *JSValueRef) JSObjectRef {
	result, callErr := tryJSObjectMakeTypedArrayWithArrayBuffer(ctx, arrayType, buffer, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectMakeTypedArrayWithArrayBufferAndOffset func(ctx JSContextRef, arrayType JSTypedArrayType, buffer JSObjectRef, byteOffset uintptr, length uintptr, exception *JSValueRef) JSObjectRef
var _jSObjectMakeTypedArrayWithArrayBufferAndOffsetErr error

func tryJSObjectMakeTypedArrayWithArrayBufferAndOffset(ctx JSContextRef, arrayType JSTypedArrayType, buffer JSObjectRef, byteOffset uintptr, length uintptr, exception *JSValueRef) (JSObjectRef, error) {
	if _jSObjectMakeTypedArrayWithArrayBufferAndOffset == nil {
		return *new(JSObjectRef), symbolCallError("JSObjectMakeTypedArrayWithArrayBufferAndOffset", "10.12", _jSObjectMakeTypedArrayWithArrayBufferAndOffsetErr)
	}
	return _jSObjectMakeTypedArrayWithArrayBufferAndOffset(ctx, arrayType, buffer, byteOffset, length, exception), nil
}

// JSObjectMakeTypedArrayWithArrayBufferAndOffset creates a JavaScript typed array object from an existing JavaScript array buffer object with the specified offset and length.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeTypedArrayWithArrayBufferAndOffset(_:_:_:_:_:_:)
func JSObjectMakeTypedArrayWithArrayBufferAndOffset(ctx JSContextRef, arrayType JSTypedArrayType, buffer JSObjectRef, byteOffset uintptr, length uintptr, exception *JSValueRef) JSObjectRef {
	result, callErr := tryJSObjectMakeTypedArrayWithArrayBufferAndOffset(ctx, arrayType, buffer, byteOffset, length, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectMakeTypedArrayWithBytesNoCopy func(ctx JSContextRef, arrayType JSTypedArrayType, bytes unsafe.Pointer, byteLength uintptr, bytesDeallocator JSTypedArrayBytesDeallocator, deallocatorContext unsafe.Pointer, exception *JSValueRef) JSObjectRef
var _jSObjectMakeTypedArrayWithBytesNoCopyErr error

func tryJSObjectMakeTypedArrayWithBytesNoCopy(ctx JSContextRef, arrayType JSTypedArrayType, bytes unsafe.Pointer, byteLength uintptr, bytesDeallocator JSTypedArrayBytesDeallocator, deallocatorContext unsafe.Pointer, exception *JSValueRef) (JSObjectRef, error) {
	if _jSObjectMakeTypedArrayWithBytesNoCopy == nil {
		return *new(JSObjectRef), symbolCallError("JSObjectMakeTypedArrayWithBytesNoCopy", "10.12", _jSObjectMakeTypedArrayWithBytesNoCopyErr)
	}
	return _jSObjectMakeTypedArrayWithBytesNoCopy(ctx, arrayType, bytes, byteLength, bytesDeallocator, deallocatorContext, exception), nil
}

// JSObjectMakeTypedArrayWithBytesNoCopy creates a JavaScript typed array object from an existing pointer.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectMakeTypedArrayWithBytesNoCopy(_:_:_:_:_:_:_:)
func JSObjectMakeTypedArrayWithBytesNoCopy(ctx JSContextRef, arrayType JSTypedArrayType, bytes unsafe.Pointer, byteLength uintptr, bytesDeallocator JSTypedArrayBytesDeallocator, deallocatorContext unsafe.Pointer, exception *JSValueRef) JSObjectRef {
	result, callErr := tryJSObjectMakeTypedArrayWithBytesNoCopy(ctx, arrayType, bytes, byteLength, bytesDeallocator, deallocatorContext, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectSetPrivate func(object JSObjectRef, data unsafe.Pointer) bool
var _jSObjectSetPrivateErr error

func tryJSObjectSetPrivate(object JSObjectRef, data unsafe.Pointer) (bool, error) {
	if _jSObjectSetPrivate == nil {
		return false, symbolCallError("JSObjectSetPrivate", "10.5", _jSObjectSetPrivateErr)
	}
	return _jSObjectSetPrivate(object, data), nil
}

// JSObjectSetPrivate sets a pointer to private data on an object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectSetPrivate(_:_:)
func JSObjectSetPrivate(object JSObjectRef, data unsafe.Pointer) bool {
	result, callErr := tryJSObjectSetPrivate(object, data)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSObjectSetProperty func(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef, value JSValueRef, attributes JSPropertyAttributes, exception *JSValueRef)
var _jSObjectSetPropertyErr error

func tryJSObjectSetProperty(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef, value JSValueRef, attributes JSPropertyAttributes, exception *JSValueRef) error {
	if _jSObjectSetProperty == nil {
		return symbolCallError("JSObjectSetProperty", "10.5", _jSObjectSetPropertyErr)
	}
	_jSObjectSetProperty(ctx, object, propertyName, value, attributes, exception)
	return nil
}

// JSObjectSetProperty sets a property on an object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectSetProperty(_:_:_:_:_:_:)
func JSObjectSetProperty(ctx JSContextRef, object JSObjectRef, propertyName JSStringRef, value JSValueRef, attributes JSPropertyAttributes, exception *JSValueRef) {
	if callErr := tryJSObjectSetProperty(ctx, object, propertyName, value, attributes, exception); callErr != nil {
		panic(callErr)
	}
}

var _jSObjectSetPropertyAtIndex func(ctx JSContextRef, object JSObjectRef, propertyIndex uint, value JSValueRef, exception *JSValueRef)
var _jSObjectSetPropertyAtIndexErr error

func tryJSObjectSetPropertyAtIndex(ctx JSContextRef, object JSObjectRef, propertyIndex uint, value JSValueRef, exception *JSValueRef) error {
	if _jSObjectSetPropertyAtIndex == nil {
		return symbolCallError("JSObjectSetPropertyAtIndex", "10.5", _jSObjectSetPropertyAtIndexErr)
	}
	_jSObjectSetPropertyAtIndex(ctx, object, propertyIndex, value, exception)
	return nil
}

// JSObjectSetPropertyAtIndex sets a property on an object by numeric index.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectSetPropertyAtIndex(_:_:_:_:_:)
func JSObjectSetPropertyAtIndex(ctx JSContextRef, object JSObjectRef, propertyIndex uint, value JSValueRef, exception *JSValueRef) {
	if callErr := tryJSObjectSetPropertyAtIndex(ctx, object, propertyIndex, value, exception); callErr != nil {
		panic(callErr)
	}
}

var _jSObjectSetPropertyForKey func(ctx JSContextRef, object JSObjectRef, propertyKey JSValueRef, value JSValueRef, attributes JSPropertyAttributes, exception *JSValueRef)
var _jSObjectSetPropertyForKeyErr error

func tryJSObjectSetPropertyForKey(ctx JSContextRef, object JSObjectRef, propertyKey JSValueRef, value JSValueRef, attributes JSPropertyAttributes, exception *JSValueRef) error {
	if _jSObjectSetPropertyForKey == nil {
		return symbolCallError("JSObjectSetPropertyForKey", "10.15", _jSObjectSetPropertyForKeyErr)
	}
	_jSObjectSetPropertyForKey(ctx, object, propertyKey, value, attributes, exception)
	return nil
}

// JSObjectSetPropertyForKey sets a property on an object using a JavaScript value as the property key.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectSetPropertyForKey(_:_:_:_:_:_:)
func JSObjectSetPropertyForKey(ctx JSContextRef, object JSObjectRef, propertyKey JSValueRef, value JSValueRef, attributes JSPropertyAttributes, exception *JSValueRef) {
	if callErr := tryJSObjectSetPropertyForKey(ctx, object, propertyKey, value, attributes, exception); callErr != nil {
		panic(callErr)
	}
}

var _jSObjectSetPrototype func(ctx JSContextRef, object JSObjectRef, value JSValueRef)
var _jSObjectSetPrototypeErr error

func tryJSObjectSetPrototype(ctx JSContextRef, object JSObjectRef, value JSValueRef) error {
	if _jSObjectSetPrototype == nil {
		return symbolCallError("JSObjectSetPrototype", "10.5", _jSObjectSetPrototypeErr)
	}
	_jSObjectSetPrototype(ctx, object, value)
	return nil
}

// JSObjectSetPrototype sets an object’s prototype.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSObjectSetPrototype(_:_:_:)
func JSObjectSetPrototype(ctx JSContextRef, object JSObjectRef, value JSValueRef) {
	if callErr := tryJSObjectSetPrototype(ctx, object, value); callErr != nil {
		panic(callErr)
	}
}

var _jSPropertyNameAccumulatorAddName func(accumulator JSPropertyNameAccumulatorRef, propertyName JSStringRef)
var _jSPropertyNameAccumulatorAddNameErr error

func tryJSPropertyNameAccumulatorAddName(accumulator JSPropertyNameAccumulatorRef, propertyName JSStringRef) error {
	if _jSPropertyNameAccumulatorAddName == nil {
		return symbolCallError("JSPropertyNameAccumulatorAddName", "10.5", _jSPropertyNameAccumulatorAddNameErr)
	}
	_jSPropertyNameAccumulatorAddName(accumulator, propertyName)
	return nil
}

// JSPropertyNameAccumulatorAddName adds a property name to a JavaScript property name accumulator.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyNameAccumulatorAddName(_:_:)
func JSPropertyNameAccumulatorAddName(accumulator JSPropertyNameAccumulatorRef, propertyName JSStringRef) {
	if callErr := tryJSPropertyNameAccumulatorAddName(accumulator, propertyName); callErr != nil {
		panic(callErr)
	}
}

var _jSPropertyNameArrayGetCount func(array JSPropertyNameArrayRef) uintptr
var _jSPropertyNameArrayGetCountErr error

func tryJSPropertyNameArrayGetCount(array JSPropertyNameArrayRef) (uintptr, error) {
	if _jSPropertyNameArrayGetCount == nil {
		return 0, symbolCallError("JSPropertyNameArrayGetCount", "10.5", _jSPropertyNameArrayGetCountErr)
	}
	return _jSPropertyNameArrayGetCount(array), nil
}

// JSPropertyNameArrayGetCount gets a count of the number of items in a JavaScript property name array.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyNameArrayGetCount(_:)
func JSPropertyNameArrayGetCount(array JSPropertyNameArrayRef) uintptr {
	result, callErr := tryJSPropertyNameArrayGetCount(array)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSPropertyNameArrayGetNameAtIndex func(array JSPropertyNameArrayRef, index uintptr) JSStringRef
var _jSPropertyNameArrayGetNameAtIndexErr error

func tryJSPropertyNameArrayGetNameAtIndex(array JSPropertyNameArrayRef, index uintptr) (JSStringRef, error) {
	if _jSPropertyNameArrayGetNameAtIndex == nil {
		return *new(JSStringRef), symbolCallError("JSPropertyNameArrayGetNameAtIndex", "10.5", _jSPropertyNameArrayGetNameAtIndexErr)
	}
	return _jSPropertyNameArrayGetNameAtIndex(array, index), nil
}

// JSPropertyNameArrayGetNameAtIndex gets a property name at a specified index in a JavaScript property name array.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyNameArrayGetNameAtIndex(_:_:)
func JSPropertyNameArrayGetNameAtIndex(array JSPropertyNameArrayRef, index uintptr) JSStringRef {
	result, callErr := tryJSPropertyNameArrayGetNameAtIndex(array, index)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSPropertyNameArrayRelease func(array JSPropertyNameArrayRef)
var _jSPropertyNameArrayReleaseErr error

func tryJSPropertyNameArrayRelease(array JSPropertyNameArrayRef) error {
	if _jSPropertyNameArrayRelease == nil {
		return symbolCallError("JSPropertyNameArrayRelease", "10.5", _jSPropertyNameArrayReleaseErr)
	}
	_jSPropertyNameArrayRelease(array)
	return nil
}

// JSPropertyNameArrayRelease releases a JavaScript property name array.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyNameArrayRelease(_:)
func JSPropertyNameArrayRelease(array JSPropertyNameArrayRef) {
	if callErr := tryJSPropertyNameArrayRelease(array); callErr != nil {
		panic(callErr)
	}
}

var _jSPropertyNameArrayRetain func(array JSPropertyNameArrayRef) JSPropertyNameArrayRef
var _jSPropertyNameArrayRetainErr error

func tryJSPropertyNameArrayRetain(array JSPropertyNameArrayRef) (JSPropertyNameArrayRef, error) {
	if _jSPropertyNameArrayRetain == nil {
		return *new(JSPropertyNameArrayRef), symbolCallError("JSPropertyNameArrayRetain", "10.5", _jSPropertyNameArrayRetainErr)
	}
	return _jSPropertyNameArrayRetain(array), nil
}

// JSPropertyNameArrayRetain retains a JavaScript property name array.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyNameArrayRetain(_:)
func JSPropertyNameArrayRetain(array JSPropertyNameArrayRef) JSPropertyNameArrayRef {
	result, callErr := tryJSPropertyNameArrayRetain(array)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSStringCopyCFString func(alloc corefoundation.CFAllocatorRef, string_ JSStringRef) corefoundation.CFStringRef
var _jSStringCopyCFStringErr error

func tryJSStringCopyCFString(alloc corefoundation.CFAllocatorRef, string_ JSStringRef) (corefoundation.CFStringRef, error) {
	if _jSStringCopyCFString == nil {
		return *new(corefoundation.CFStringRef), symbolCallError("JSStringCopyCFString", "10.5", _jSStringCopyCFStringErr)
	}
	return _jSStringCopyCFString(alloc, string_), nil
}

// JSStringCopyCFString creates a Core Foundation string from a JavaScript string.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSStringCopyCFString(_:_:)
func JSStringCopyCFString(alloc corefoundation.CFAllocatorRef, string_ JSStringRef) corefoundation.CFStringRef {
	result, callErr := tryJSStringCopyCFString(alloc, string_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSStringCreateWithCFString func(string_ corefoundation.CFStringRef) JSStringRef
var _jSStringCreateWithCFStringErr error

func tryJSStringCreateWithCFString(string_ corefoundation.CFStringRef) (JSStringRef, error) {
	if _jSStringCreateWithCFString == nil {
		return *new(JSStringRef), symbolCallError("JSStringCreateWithCFString", "10.5", _jSStringCreateWithCFStringErr)
	}
	return _jSStringCreateWithCFString(string_), nil
}

// JSStringCreateWithCFString creates a JavaScript string from a Core Foundation string.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSStringCreateWithCFString(_:)
func JSStringCreateWithCFString(string_ corefoundation.CFStringRef) JSStringRef {
	result, callErr := tryJSStringCreateWithCFString(string_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSStringCreateWithCharacters func(chars *JSChar, numChars uintptr) JSStringRef
var _jSStringCreateWithCharactersErr error

func tryJSStringCreateWithCharacters(chars *JSChar, numChars uintptr) (JSStringRef, error) {
	if _jSStringCreateWithCharacters == nil {
		return *new(JSStringRef), symbolCallError("JSStringCreateWithCharacters", "10.5", _jSStringCreateWithCharactersErr)
	}
	return _jSStringCreateWithCharacters(chars, numChars), nil
}

// JSStringCreateWithCharacters creates a JavaScript string from a buffer of Unicode characters.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSStringCreateWithCharacters(_:_:)
func JSStringCreateWithCharacters(chars *JSChar, numChars uintptr) JSStringRef {
	result, callErr := tryJSStringCreateWithCharacters(chars, numChars)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSStringCreateWithUTF8CString func(string_ string) JSStringRef
var _jSStringCreateWithUTF8CStringErr error

func tryJSStringCreateWithUTF8CString(string_ string) (JSStringRef, error) {
	if _jSStringCreateWithUTF8CString == nil {
		return *new(JSStringRef), symbolCallError("JSStringCreateWithUTF8CString", "10.5", _jSStringCreateWithUTF8CStringErr)
	}
	return _jSStringCreateWithUTF8CString(string_), nil
}

// JSStringCreateWithUTF8CString creates a JavaScript string from a null-terminated UTF-8 string.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSStringCreateWithUTF8CString(_:)
func JSStringCreateWithUTF8CString(string_ string) JSStringRef {
	result, callErr := tryJSStringCreateWithUTF8CString(string_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSStringGetCharactersPtr func(string_ JSStringRef) *JSChar
var _jSStringGetCharactersPtrErr error

func tryJSStringGetCharactersPtr(string_ JSStringRef) (*JSChar, error) {
	if _jSStringGetCharactersPtr == nil {
		return nil, symbolCallError("JSStringGetCharactersPtr", "10.5", _jSStringGetCharactersPtrErr)
	}
	return _jSStringGetCharactersPtr(string_), nil
}

// JSStringGetCharactersPtr returns a pointer to the Unicode character buffer that serves as the backing store for a JavaScript string.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSStringGetCharactersPtr(_:)
func JSStringGetCharactersPtr(string_ JSStringRef) *JSChar {
	result, callErr := tryJSStringGetCharactersPtr(string_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSStringGetLength func(string_ JSStringRef) uintptr
var _jSStringGetLengthErr error

func tryJSStringGetLength(string_ JSStringRef) (uintptr, error) {
	if _jSStringGetLength == nil {
		return 0, symbolCallError("JSStringGetLength", "10.5", _jSStringGetLengthErr)
	}
	return _jSStringGetLength(string_), nil
}

// JSStringGetLength returns the number of Unicode characters in a JavaScript string.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSStringGetLength(_:)
func JSStringGetLength(string_ JSStringRef) uintptr {
	result, callErr := tryJSStringGetLength(string_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSStringGetMaximumUTF8CStringSize func(string_ JSStringRef) uintptr
var _jSStringGetMaximumUTF8CStringSizeErr error

func tryJSStringGetMaximumUTF8CStringSize(string_ JSStringRef) (uintptr, error) {
	if _jSStringGetMaximumUTF8CStringSize == nil {
		return 0, symbolCallError("JSStringGetMaximumUTF8CStringSize", "10.5", _jSStringGetMaximumUTF8CStringSizeErr)
	}
	return _jSStringGetMaximumUTF8CStringSize(string_), nil
}

// JSStringGetMaximumUTF8CStringSize returns the maximum number of bytes a JavaScript string uses when you convert it into a null-terminated UTF-8 string.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSStringGetMaximumUTF8CStringSize(_:)
func JSStringGetMaximumUTF8CStringSize(string_ JSStringRef) uintptr {
	result, callErr := tryJSStringGetMaximumUTF8CStringSize(string_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSStringGetUTF8CString func(string_ JSStringRef, buffer *byte, bufferSize uintptr) uintptr
var _jSStringGetUTF8CStringErr error

func tryJSStringGetUTF8CString(string_ JSStringRef, buffer *byte, bufferSize uintptr) (uintptr, error) {
	if _jSStringGetUTF8CString == nil {
		return 0, symbolCallError("JSStringGetUTF8CString", "10.5", _jSStringGetUTF8CStringErr)
	}
	return _jSStringGetUTF8CString(string_, buffer, bufferSize), nil
}

// JSStringGetUTF8CString converts a JavaScript string into a null-terminated UTF-8 string, and copies the result into an external byte buffer.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSStringGetUTF8CString(_:_:_:)
func JSStringGetUTF8CString(string_ JSStringRef, buffer *byte, bufferSize uintptr) uintptr {
	result, callErr := tryJSStringGetUTF8CString(string_, buffer, bufferSize)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSStringIsEqual func(a JSStringRef, b JSStringRef) bool
var _jSStringIsEqualErr error

func tryJSStringIsEqual(a JSStringRef, b JSStringRef) (bool, error) {
	if _jSStringIsEqual == nil {
		return false, symbolCallError("JSStringIsEqual", "10.5", _jSStringIsEqualErr)
	}
	return _jSStringIsEqual(a, b), nil
}

// JSStringIsEqual tests whether two JavaScript strings match.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSStringIsEqual(_:_:)
func JSStringIsEqual(a JSStringRef, b JSStringRef) bool {
	result, callErr := tryJSStringIsEqual(a, b)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSStringIsEqualToUTF8CString func(a JSStringRef, b string) bool
var _jSStringIsEqualToUTF8CStringErr error

func tryJSStringIsEqualToUTF8CString(a JSStringRef, b string) (bool, error) {
	if _jSStringIsEqualToUTF8CString == nil {
		return false, symbolCallError("JSStringIsEqualToUTF8CString", "10.5", _jSStringIsEqualToUTF8CStringErr)
	}
	return _jSStringIsEqualToUTF8CString(a, b), nil
}

// JSStringIsEqualToUTF8CString tests whether a JavaScript string matches a null-terminated UTF-8 string.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSStringIsEqualToUTF8CString(_:_:)
func JSStringIsEqualToUTF8CString(a JSStringRef, b string) bool {
	result, callErr := tryJSStringIsEqualToUTF8CString(a, b)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSStringRelease func(string_ JSStringRef)
var _jSStringReleaseErr error

func tryJSStringRelease(string_ JSStringRef) error {
	if _jSStringRelease == nil {
		return symbolCallError("JSStringRelease", "10.5", _jSStringReleaseErr)
	}
	_jSStringRelease(string_)
	return nil
}

// JSStringRelease releases a JavaScript string.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSStringRelease(_:)
func JSStringRelease(string_ JSStringRef) {
	if callErr := tryJSStringRelease(string_); callErr != nil {
		panic(callErr)
	}
}

var _jSStringRetain func(string_ JSStringRef) JSStringRef
var _jSStringRetainErr error

func tryJSStringRetain(string_ JSStringRef) (JSStringRef, error) {
	if _jSStringRetain == nil {
		return *new(JSStringRef), symbolCallError("JSStringRetain", "10.5", _jSStringRetainErr)
	}
	return _jSStringRetain(string_), nil
}

// JSStringRetain retains a JavaScript string.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSStringRetain(_:)
func JSStringRetain(string_ JSStringRef) JSStringRef {
	result, callErr := tryJSStringRetain(string_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueCompare func(ctx JSContextRef, left JSValueRef, right JSValueRef, exception *JSValueRef) JSRelationCondition
var _jSValueCompareErr error

func tryJSValueCompare(ctx JSContextRef, left JSValueRef, right JSValueRef, exception *JSValueRef) (JSRelationCondition, error) {
	if _jSValueCompare == nil {
		return *new(JSRelationCondition), symbolCallError("JSValueCompare", "15.0", _jSValueCompareErr)
	}
	return _jSValueCompare(ctx, left, right, exception), nil
}

// JSValueCompare.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueCompare(_:_:_:_:)
func JSValueCompare(ctx JSContextRef, left JSValueRef, right JSValueRef, exception *JSValueRef) JSRelationCondition {
	result, callErr := tryJSValueCompare(ctx, left, right, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueCompareDouble func(ctx JSContextRef, left JSValueRef, right float64, exception *JSValueRef) JSRelationCondition
var _jSValueCompareDoubleErr error

func tryJSValueCompareDouble(ctx JSContextRef, left JSValueRef, right float64, exception *JSValueRef) (JSRelationCondition, error) {
	if _jSValueCompareDouble == nil {
		return *new(JSRelationCondition), symbolCallError("JSValueCompareDouble", "15.0", _jSValueCompareDoubleErr)
	}
	return _jSValueCompareDouble(ctx, left, right, exception), nil
}

// JSValueCompareDouble.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueCompareDouble(_:_:_:_:)
func JSValueCompareDouble(ctx JSContextRef, left JSValueRef, right float64, exception *JSValueRef) JSRelationCondition {
	result, callErr := tryJSValueCompareDouble(ctx, left, right, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueCompareInt64 func(ctx JSContextRef, left JSValueRef, right int64, exception *JSValueRef) JSRelationCondition
var _jSValueCompareInt64Err error

func tryJSValueCompareInt64(ctx JSContextRef, left JSValueRef, right int64, exception *JSValueRef) (JSRelationCondition, error) {
	if _jSValueCompareInt64 == nil {
		return *new(JSRelationCondition), symbolCallError("JSValueCompareInt64", "15.0", _jSValueCompareInt64Err)
	}
	return _jSValueCompareInt64(ctx, left, right, exception), nil
}

// JSValueCompareInt64.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueCompareInt64(_:_:_:_:)
func JSValueCompareInt64(ctx JSContextRef, left JSValueRef, right int64, exception *JSValueRef) JSRelationCondition {
	result, callErr := tryJSValueCompareInt64(ctx, left, right, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueCompareUInt64 func(ctx JSContextRef, left JSValueRef, right uint64, exception *JSValueRef) JSRelationCondition
var _jSValueCompareUInt64Err error

func tryJSValueCompareUInt64(ctx JSContextRef, left JSValueRef, right uint64, exception *JSValueRef) (JSRelationCondition, error) {
	if _jSValueCompareUInt64 == nil {
		return *new(JSRelationCondition), symbolCallError("JSValueCompareUInt64", "15.0", _jSValueCompareUInt64Err)
	}
	return _jSValueCompareUInt64(ctx, left, right, exception), nil
}

// JSValueCompareUInt64.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueCompareUInt64(_:_:_:_:)
func JSValueCompareUInt64(ctx JSContextRef, left JSValueRef, right uint64, exception *JSValueRef) JSRelationCondition {
	result, callErr := tryJSValueCompareUInt64(ctx, left, right, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueCreateJSONString func(ctx JSContextRef, value JSValueRef, indent uint, exception *JSValueRef) JSStringRef
var _jSValueCreateJSONStringErr error

func tryJSValueCreateJSONString(ctx JSContextRef, value JSValueRef, indent uint, exception *JSValueRef) (JSStringRef, error) {
	if _jSValueCreateJSONString == nil {
		return *new(JSStringRef), symbolCallError("JSValueCreateJSONString", "10.7", _jSValueCreateJSONStringErr)
	}
	return _jSValueCreateJSONString(ctx, value, indent, exception), nil
}

// JSValueCreateJSONString creates a JavaScript string that contains the JSON-serialized representation of a JavaScript value.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueCreateJSONString(_:_:_:_:)
func JSValueCreateJSONString(ctx JSContextRef, value JSValueRef, indent uint, exception *JSValueRef) JSStringRef {
	result, callErr := tryJSValueCreateJSONString(ctx, value, indent, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueGetType func(ctx JSContextRef, value JSValueRef) JSType
var _jSValueGetTypeErr error

func tryJSValueGetType(ctx JSContextRef, value JSValueRef) (JSType, error) {
	if _jSValueGetType == nil {
		return *new(JSType), symbolCallError("JSValueGetType", "10.5", _jSValueGetTypeErr)
	}
	return _jSValueGetType(ctx, value), nil
}

// JSValueGetType returns a JavaScript value’s type.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueGetType(_:_:)
func JSValueGetType(ctx JSContextRef, value JSValueRef) JSType {
	result, callErr := tryJSValueGetType(ctx, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueGetTypedArrayType func(ctx JSContextRef, value JSValueRef, exception *JSValueRef) JSTypedArrayType
var _jSValueGetTypedArrayTypeErr error

func tryJSValueGetTypedArrayType(ctx JSContextRef, value JSValueRef, exception *JSValueRef) (JSTypedArrayType, error) {
	if _jSValueGetTypedArrayType == nil {
		return *new(JSTypedArrayType), symbolCallError("JSValueGetTypedArrayType", "10.12", _jSValueGetTypedArrayTypeErr)
	}
	return _jSValueGetTypedArrayType(ctx, value, exception), nil
}

// JSValueGetTypedArrayType returns a JavaScript value’s typed array type.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueGetTypedArrayType(_:_:_:)
func JSValueGetTypedArrayType(ctx JSContextRef, value JSValueRef, exception *JSValueRef) JSTypedArrayType {
	result, callErr := tryJSValueGetTypedArrayType(ctx, value, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueIsArray func(ctx JSContextRef, value JSValueRef) bool
var _jSValueIsArrayErr error

func tryJSValueIsArray(ctx JSContextRef, value JSValueRef) (bool, error) {
	if _jSValueIsArray == nil {
		return false, symbolCallError("JSValueIsArray", "10.11", _jSValueIsArrayErr)
	}
	return _jSValueIsArray(ctx, value), nil
}

// JSValueIsArray tests whether a JavaScript value is an array.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsArray(_:_:)
func JSValueIsArray(ctx JSContextRef, value JSValueRef) bool {
	result, callErr := tryJSValueIsArray(ctx, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueIsBigInt func(ctx JSContextRef, value JSValueRef) bool
var _jSValueIsBigIntErr error

func tryJSValueIsBigInt(ctx JSContextRef, value JSValueRef) (bool, error) {
	if _jSValueIsBigInt == nil {
		return false, symbolCallError("JSValueIsBigInt", "15.0", _jSValueIsBigIntErr)
	}
	return _jSValueIsBigInt(ctx, value), nil
}

// JSValueIsBigInt.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsBigInt(_:_:)
func JSValueIsBigInt(ctx JSContextRef, value JSValueRef) bool {
	result, callErr := tryJSValueIsBigInt(ctx, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueIsBoolean func(ctx JSContextRef, value JSValueRef) bool
var _jSValueIsBooleanErr error

func tryJSValueIsBoolean(ctx JSContextRef, value JSValueRef) (bool, error) {
	if _jSValueIsBoolean == nil {
		return false, symbolCallError("JSValueIsBoolean", "10.5", _jSValueIsBooleanErr)
	}
	return _jSValueIsBoolean(ctx, value), nil
}

// JSValueIsBoolean tests whether a JavaScript value is Boolean.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsBoolean(_:_:)
func JSValueIsBoolean(ctx JSContextRef, value JSValueRef) bool {
	result, callErr := tryJSValueIsBoolean(ctx, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueIsDate func(ctx JSContextRef, value JSValueRef) bool
var _jSValueIsDateErr error

func tryJSValueIsDate(ctx JSContextRef, value JSValueRef) (bool, error) {
	if _jSValueIsDate == nil {
		return false, symbolCallError("JSValueIsDate", "10.11", _jSValueIsDateErr)
	}
	return _jSValueIsDate(ctx, value), nil
}

// JSValueIsDate tests whether a JavaScript value is a date.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsDate(_:_:)
func JSValueIsDate(ctx JSContextRef, value JSValueRef) bool {
	result, callErr := tryJSValueIsDate(ctx, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueIsEqual func(ctx JSContextRef, a JSValueRef, b JSValueRef, exception *JSValueRef) bool
var _jSValueIsEqualErr error

func tryJSValueIsEqual(ctx JSContextRef, a JSValueRef, b JSValueRef, exception *JSValueRef) (bool, error) {
	if _jSValueIsEqual == nil {
		return false, symbolCallError("JSValueIsEqual", "10.5", _jSValueIsEqualErr)
	}
	return _jSValueIsEqual(ctx, a, b, exception), nil
}

// JSValueIsEqual tests whether two JavaScript values are equal.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsEqual(_:_:_:_:)
func JSValueIsEqual(ctx JSContextRef, a JSValueRef, b JSValueRef, exception *JSValueRef) bool {
	result, callErr := tryJSValueIsEqual(ctx, a, b, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueIsInstanceOfConstructor func(ctx JSContextRef, value JSValueRef, constructor JSObjectRef, exception *JSValueRef) bool
var _jSValueIsInstanceOfConstructorErr error

func tryJSValueIsInstanceOfConstructor(ctx JSContextRef, value JSValueRef, constructor JSObjectRef, exception *JSValueRef) (bool, error) {
	if _jSValueIsInstanceOfConstructor == nil {
		return false, symbolCallError("JSValueIsInstanceOfConstructor", "10.5", _jSValueIsInstanceOfConstructorErr)
	}
	return _jSValueIsInstanceOfConstructor(ctx, value, constructor, exception), nil
}

// JSValueIsInstanceOfConstructor tests whether a JavaScript value is an object that the specified constructor creates.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsInstanceOfConstructor(_:_:_:_:)
func JSValueIsInstanceOfConstructor(ctx JSContextRef, value JSValueRef, constructor JSObjectRef, exception *JSValueRef) bool {
	result, callErr := tryJSValueIsInstanceOfConstructor(ctx, value, constructor, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueIsNull func(ctx JSContextRef, value JSValueRef) bool
var _jSValueIsNullErr error

func tryJSValueIsNull(ctx JSContextRef, value JSValueRef) (bool, error) {
	if _jSValueIsNull == nil {
		return false, symbolCallError("JSValueIsNull", "10.5", _jSValueIsNullErr)
	}
	return _jSValueIsNull(ctx, value), nil
}

// JSValueIsNull tests whether a JavaScript value’s type is the null type.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsNull(_:_:)
func JSValueIsNull(ctx JSContextRef, value JSValueRef) bool {
	result, callErr := tryJSValueIsNull(ctx, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueIsNumber func(ctx JSContextRef, value JSValueRef) bool
var _jSValueIsNumberErr error

func tryJSValueIsNumber(ctx JSContextRef, value JSValueRef) (bool, error) {
	if _jSValueIsNumber == nil {
		return false, symbolCallError("JSValueIsNumber", "10.5", _jSValueIsNumberErr)
	}
	return _jSValueIsNumber(ctx, value), nil
}

// JSValueIsNumber tests whether a JavaScript value’s type is the number type.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsNumber(_:_:)
func JSValueIsNumber(ctx JSContextRef, value JSValueRef) bool {
	result, callErr := tryJSValueIsNumber(ctx, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueIsObject func(ctx JSContextRef, value JSValueRef) bool
var _jSValueIsObjectErr error

func tryJSValueIsObject(ctx JSContextRef, value JSValueRef) (bool, error) {
	if _jSValueIsObject == nil {
		return false, symbolCallError("JSValueIsObject", "10.5", _jSValueIsObjectErr)
	}
	return _jSValueIsObject(ctx, value), nil
}

// JSValueIsObject tests whether a JavaScript value’s type is the object type.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsObject(_:_:)
func JSValueIsObject(ctx JSContextRef, value JSValueRef) bool {
	result, callErr := tryJSValueIsObject(ctx, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueIsObjectOfClass func(ctx JSContextRef, value JSValueRef, jsClass JSClassRef) bool
var _jSValueIsObjectOfClassErr error

func tryJSValueIsObjectOfClass(ctx JSContextRef, value JSValueRef, jsClass JSClassRef) (bool, error) {
	if _jSValueIsObjectOfClass == nil {
		return false, symbolCallError("JSValueIsObjectOfClass", "10.5", _jSValueIsObjectOfClassErr)
	}
	return _jSValueIsObjectOfClass(ctx, value, jsClass), nil
}

// JSValueIsObjectOfClass tests whether a JavaScript value is an object with a specified class in its class chain.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsObjectOfClass(_:_:_:)
func JSValueIsObjectOfClass(ctx JSContextRef, value JSValueRef, jsClass JSClassRef) bool {
	result, callErr := tryJSValueIsObjectOfClass(ctx, value, jsClass)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueIsStrictEqual func(ctx JSContextRef, a JSValueRef, b JSValueRef) bool
var _jSValueIsStrictEqualErr error

func tryJSValueIsStrictEqual(ctx JSContextRef, a JSValueRef, b JSValueRef) (bool, error) {
	if _jSValueIsStrictEqual == nil {
		return false, symbolCallError("JSValueIsStrictEqual", "10.5", _jSValueIsStrictEqualErr)
	}
	return _jSValueIsStrictEqual(ctx, a, b), nil
}

// JSValueIsStrictEqual tests whether two JavaScript values are strict equal.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsStrictEqual(_:_:_:)
func JSValueIsStrictEqual(ctx JSContextRef, a JSValueRef, b JSValueRef) bool {
	result, callErr := tryJSValueIsStrictEqual(ctx, a, b)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueIsString func(ctx JSContextRef, value JSValueRef) bool
var _jSValueIsStringErr error

func tryJSValueIsString(ctx JSContextRef, value JSValueRef) (bool, error) {
	if _jSValueIsString == nil {
		return false, symbolCallError("JSValueIsString", "10.5", _jSValueIsStringErr)
	}
	return _jSValueIsString(ctx, value), nil
}

// JSValueIsString tests whether a JavaScript value’s type is the string type.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsString(_:_:)
func JSValueIsString(ctx JSContextRef, value JSValueRef) bool {
	result, callErr := tryJSValueIsString(ctx, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueIsSymbol func(ctx JSContextRef, value JSValueRef) bool
var _jSValueIsSymbolErr error

func tryJSValueIsSymbol(ctx JSContextRef, value JSValueRef) (bool, error) {
	if _jSValueIsSymbol == nil {
		return false, symbolCallError("JSValueIsSymbol", "10.15", _jSValueIsSymbolErr)
	}
	return _jSValueIsSymbol(ctx, value), nil
}

// JSValueIsSymbol tests whether a JavaScript value’s type is the symbol type.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsSymbol(_:_:)
func JSValueIsSymbol(ctx JSContextRef, value JSValueRef) bool {
	result, callErr := tryJSValueIsSymbol(ctx, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueIsUndefined func(ctx JSContextRef, value JSValueRef) bool
var _jSValueIsUndefinedErr error

func tryJSValueIsUndefined(ctx JSContextRef, value JSValueRef) (bool, error) {
	if _jSValueIsUndefined == nil {
		return false, symbolCallError("JSValueIsUndefined", "10.5", _jSValueIsUndefinedErr)
	}
	return _jSValueIsUndefined(ctx, value), nil
}

// JSValueIsUndefined tests whether a JavaScript value’s type is the undefined type.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueIsUndefined(_:_:)
func JSValueIsUndefined(ctx JSContextRef, value JSValueRef) bool {
	result, callErr := tryJSValueIsUndefined(ctx, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueMakeBoolean func(ctx JSContextRef, boolean bool) JSValueRef
var _jSValueMakeBooleanErr error

func tryJSValueMakeBoolean(ctx JSContextRef, boolean bool) (JSValueRef, error) {
	if _jSValueMakeBoolean == nil {
		return *new(JSValueRef), symbolCallError("JSValueMakeBoolean", "10.5", _jSValueMakeBooleanErr)
	}
	return _jSValueMakeBoolean(ctx, boolean), nil
}

// JSValueMakeBoolean creates a JavaScript Boolean value.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeBoolean(_:_:)
func JSValueMakeBoolean(ctx JSContextRef, boolean bool) JSValueRef {
	result, callErr := tryJSValueMakeBoolean(ctx, boolean)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueMakeFromJSONString func(ctx JSContextRef, string_ JSStringRef) JSValueRef
var _jSValueMakeFromJSONStringErr error

func tryJSValueMakeFromJSONString(ctx JSContextRef, string_ JSStringRef) (JSValueRef, error) {
	if _jSValueMakeFromJSONString == nil {
		return *new(JSValueRef), symbolCallError("JSValueMakeFromJSONString", "10.7", _jSValueMakeFromJSONStringErr)
	}
	return _jSValueMakeFromJSONString(ctx, string_), nil
}

// JSValueMakeFromJSONString creates a JavaScript value from a JSON-formatted string.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeFromJSONString(_:_:)
func JSValueMakeFromJSONString(ctx JSContextRef, string_ JSStringRef) JSValueRef {
	result, callErr := tryJSValueMakeFromJSONString(ctx, string_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueMakeNull func(ctx JSContextRef) JSValueRef
var _jSValueMakeNullErr error

func tryJSValueMakeNull(ctx JSContextRef) (JSValueRef, error) {
	if _jSValueMakeNull == nil {
		return *new(JSValueRef), symbolCallError("JSValueMakeNull", "10.5", _jSValueMakeNullErr)
	}
	return _jSValueMakeNull(ctx), nil
}

// JSValueMakeNull creates a JavaScript value of the null type.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeNull(_:)
func JSValueMakeNull(ctx JSContextRef) JSValueRef {
	result, callErr := tryJSValueMakeNull(ctx)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueMakeNumber func(ctx JSContextRef, number float64) JSValueRef
var _jSValueMakeNumberErr error

func tryJSValueMakeNumber(ctx JSContextRef, number float64) (JSValueRef, error) {
	if _jSValueMakeNumber == nil {
		return *new(JSValueRef), symbolCallError("JSValueMakeNumber", "10.5", _jSValueMakeNumberErr)
	}
	return _jSValueMakeNumber(ctx, number), nil
}

// JSValueMakeNumber creates a JavaScript value of the number type.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeNumber(_:_:)
func JSValueMakeNumber(ctx JSContextRef, number float64) JSValueRef {
	result, callErr := tryJSValueMakeNumber(ctx, number)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueMakeString func(ctx JSContextRef, string_ JSStringRef) JSValueRef
var _jSValueMakeStringErr error

func tryJSValueMakeString(ctx JSContextRef, string_ JSStringRef) (JSValueRef, error) {
	if _jSValueMakeString == nil {
		return *new(JSValueRef), symbolCallError("JSValueMakeString", "10.5", _jSValueMakeStringErr)
	}
	return _jSValueMakeString(ctx, string_), nil
}

// JSValueMakeString creates a JavaScript value of the string type.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeString(_:_:)
func JSValueMakeString(ctx JSContextRef, string_ JSStringRef) JSValueRef {
	result, callErr := tryJSValueMakeString(ctx, string_)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueMakeSymbol func(ctx JSContextRef, description JSStringRef) JSValueRef
var _jSValueMakeSymbolErr error

func tryJSValueMakeSymbol(ctx JSContextRef, description JSStringRef) (JSValueRef, error) {
	if _jSValueMakeSymbol == nil {
		return *new(JSValueRef), symbolCallError("JSValueMakeSymbol", "10.15", _jSValueMakeSymbolErr)
	}
	return _jSValueMakeSymbol(ctx, description), nil
}

// JSValueMakeSymbol creates a JavaScript value of the symbol type.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeSymbol(_:_:)
func JSValueMakeSymbol(ctx JSContextRef, description JSStringRef) JSValueRef {
	result, callErr := tryJSValueMakeSymbol(ctx, description)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueMakeUndefined func(ctx JSContextRef) JSValueRef
var _jSValueMakeUndefinedErr error

func tryJSValueMakeUndefined(ctx JSContextRef) (JSValueRef, error) {
	if _jSValueMakeUndefined == nil {
		return *new(JSValueRef), symbolCallError("JSValueMakeUndefined", "10.5", _jSValueMakeUndefinedErr)
	}
	return _jSValueMakeUndefined(ctx), nil
}

// JSValueMakeUndefined creates a JavaScript value of the undefined type.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueMakeUndefined(_:)
func JSValueMakeUndefined(ctx JSContextRef) JSValueRef {
	result, callErr := tryJSValueMakeUndefined(ctx)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueProtect func(ctx JSContextRef, value JSValueRef)
var _jSValueProtectErr error

func tryJSValueProtect(ctx JSContextRef, value JSValueRef) error {
	if _jSValueProtect == nil {
		return symbolCallError("JSValueProtect", "10.5", _jSValueProtectErr)
	}
	_jSValueProtect(ctx, value)
	return nil
}

// JSValueProtect protects a JavaScript value from garbage collection.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueProtect(_:_:)
func JSValueProtect(ctx JSContextRef, value JSValueRef) {
	if callErr := tryJSValueProtect(ctx, value); callErr != nil {
		panic(callErr)
	}
}

var _jSValueToBoolean func(ctx JSContextRef, value JSValueRef) bool
var _jSValueToBooleanErr error

func tryJSValueToBoolean(ctx JSContextRef, value JSValueRef) (bool, error) {
	if _jSValueToBoolean == nil {
		return false, symbolCallError("JSValueToBoolean", "10.5", _jSValueToBooleanErr)
	}
	return _jSValueToBoolean(ctx, value), nil
}

// JSValueToBoolean converts a JavaScript value to a Boolean and returns the resulting Boolean.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueToBoolean(_:_:)
func JSValueToBoolean(ctx JSContextRef, value JSValueRef) bool {
	result, callErr := tryJSValueToBoolean(ctx, value)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueToInt32 func(ctx JSContextRef, value JSValueRef, exception *JSValueRef) int32
var _jSValueToInt32Err error

func tryJSValueToInt32(ctx JSContextRef, value JSValueRef, exception *JSValueRef) (int32, error) {
	if _jSValueToInt32 == nil {
		return 0, symbolCallError("JSValueToInt32", "15.0", _jSValueToInt32Err)
	}
	return _jSValueToInt32(ctx, value, exception), nil
}

// JSValueToInt32.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueToInt32(_:_:_:)
func JSValueToInt32(ctx JSContextRef, value JSValueRef, exception *JSValueRef) int32 {
	result, callErr := tryJSValueToInt32(ctx, value, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueToInt64 func(ctx JSContextRef, value JSValueRef, exception *JSValueRef) int64
var _jSValueToInt64Err error

func tryJSValueToInt64(ctx JSContextRef, value JSValueRef, exception *JSValueRef) (int64, error) {
	if _jSValueToInt64 == nil {
		return 0, symbolCallError("JSValueToInt64", "15.0", _jSValueToInt64Err)
	}
	return _jSValueToInt64(ctx, value, exception), nil
}

// JSValueToInt64.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueToInt64(_:_:_:)
func JSValueToInt64(ctx JSContextRef, value JSValueRef, exception *JSValueRef) int64 {
	result, callErr := tryJSValueToInt64(ctx, value, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueToNumber func(ctx JSContextRef, value JSValueRef, exception *JSValueRef) float64
var _jSValueToNumberErr error

func tryJSValueToNumber(ctx JSContextRef, value JSValueRef, exception *JSValueRef) (float64, error) {
	if _jSValueToNumber == nil {
		return 0.0, symbolCallError("JSValueToNumber", "10.5", _jSValueToNumberErr)
	}
	return _jSValueToNumber(ctx, value, exception), nil
}

// JSValueToNumber converts a JavaScript value to a number and returns the resulting number.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueToNumber(_:_:_:)
func JSValueToNumber(ctx JSContextRef, value JSValueRef, exception *JSValueRef) float64 {
	result, callErr := tryJSValueToNumber(ctx, value, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueToObject func(ctx JSContextRef, value JSValueRef, exception *JSValueRef) JSObjectRef
var _jSValueToObjectErr error

func tryJSValueToObject(ctx JSContextRef, value JSValueRef, exception *JSValueRef) (JSObjectRef, error) {
	if _jSValueToObject == nil {
		return *new(JSObjectRef), symbolCallError("JSValueToObject", "10.5", _jSValueToObjectErr)
	}
	return _jSValueToObject(ctx, value, exception), nil
}

// JSValueToObject converts a JavaScript value to an object and returns the resulting object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueToObject(_:_:_:)
func JSValueToObject(ctx JSContextRef, value JSValueRef, exception *JSValueRef) JSObjectRef {
	result, callErr := tryJSValueToObject(ctx, value, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueToStringCopy func(ctx JSContextRef, value JSValueRef, exception *JSValueRef) JSStringRef
var _jSValueToStringCopyErr error

func tryJSValueToStringCopy(ctx JSContextRef, value JSValueRef, exception *JSValueRef) (JSStringRef, error) {
	if _jSValueToStringCopy == nil {
		return *new(JSStringRef), symbolCallError("JSValueToStringCopy", "10.5", _jSValueToStringCopyErr)
	}
	return _jSValueToStringCopy(ctx, value, exception), nil
}

// JSValueToStringCopy converts a JavaScript value to a string and copies the result into a JavaScript string.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueToStringCopy(_:_:_:)
func JSValueToStringCopy(ctx JSContextRef, value JSValueRef, exception *JSValueRef) JSStringRef {
	result, callErr := tryJSValueToStringCopy(ctx, value, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueToUInt32 func(ctx JSContextRef, value JSValueRef, exception *JSValueRef) uint32
var _jSValueToUInt32Err error

func tryJSValueToUInt32(ctx JSContextRef, value JSValueRef, exception *JSValueRef) (uint32, error) {
	if _jSValueToUInt32 == nil {
		return 0, symbolCallError("JSValueToUInt32", "15.0", _jSValueToUInt32Err)
	}
	return _jSValueToUInt32(ctx, value, exception), nil
}

// JSValueToUInt32.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueToUInt32(_:_:_:)
func JSValueToUInt32(ctx JSContextRef, value JSValueRef, exception *JSValueRef) uint32 {
	result, callErr := tryJSValueToUInt32(ctx, value, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueToUInt64 func(ctx JSContextRef, value JSValueRef, exception *JSValueRef) uint64
var _jSValueToUInt64Err error

func tryJSValueToUInt64(ctx JSContextRef, value JSValueRef, exception *JSValueRef) (uint64, error) {
	if _jSValueToUInt64 == nil {
		return 0, symbolCallError("JSValueToUInt64", "15.0", _jSValueToUInt64Err)
	}
	return _jSValueToUInt64(ctx, value, exception), nil
}

// JSValueToUInt64.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueToUInt64(_:_:_:)
func JSValueToUInt64(ctx JSContextRef, value JSValueRef, exception *JSValueRef) uint64 {
	result, callErr := tryJSValueToUInt64(ctx, value, exception)
	if callErr != nil {
		panic(callErr)
	}
	return result
}

var _jSValueUnprotect func(ctx JSContextRef, value JSValueRef)
var _jSValueUnprotectErr error

func tryJSValueUnprotect(ctx JSContextRef, value JSValueRef) error {
	if _jSValueUnprotect == nil {
		return symbolCallError("JSValueUnprotect", "10.5", _jSValueUnprotectErr)
	}
	_jSValueUnprotect(ctx, value)
	return nil
}

// JSValueUnprotect unprotects a JavaScript value from garbage collection.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSValueUnprotect(_:_:)
func JSValueUnprotect(ctx JSContextRef, value JSValueRef) {
	if callErr := tryJSValueUnprotect(ctx, value); callErr != nil {
		panic(callErr)
	}
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_jSBigIntCreateWithDouble, &_jSBigIntCreateWithDoubleErr, frameworkHandle, "JSBigIntCreateWithDouble", "15.0")
	registerFunc(&_jSBigIntCreateWithInt64, &_jSBigIntCreateWithInt64Err, frameworkHandle, "JSBigIntCreateWithInt64", "15.0")
	registerFunc(&_jSBigIntCreateWithString, &_jSBigIntCreateWithStringErr, frameworkHandle, "JSBigIntCreateWithString", "15.0")
	registerFunc(&_jSBigIntCreateWithUInt64, &_jSBigIntCreateWithUInt64Err, frameworkHandle, "JSBigIntCreateWithUInt64", "15.0")
	registerFunc(&_jSCheckScriptSyntax, &_jSCheckScriptSyntaxErr, frameworkHandle, "JSCheckScriptSyntax", "10.5")
	registerFunc(&_jSClassCreate, &_jSClassCreateErr, frameworkHandle, "JSClassCreate", "10.5")
	registerFunc(&_jSClassRelease, &_jSClassReleaseErr, frameworkHandle, "JSClassRelease", "10.5")
	registerFunc(&_jSClassRetain, &_jSClassRetainErr, frameworkHandle, "JSClassRetain", "10.5")
	registerFunc(&_jSContextGetGlobalContext, &_jSContextGetGlobalContextErr, frameworkHandle, "JSContextGetGlobalContext", "10.7")
	registerFunc(&_jSContextGetGlobalObject, &_jSContextGetGlobalObjectErr, frameworkHandle, "JSContextGetGlobalObject", "10.5")
	registerFunc(&_jSContextGetGroup, &_jSContextGetGroupErr, frameworkHandle, "JSContextGetGroup", "10.6")
	registerFunc(&_jSContextGroupCreate, &_jSContextGroupCreateErr, frameworkHandle, "JSContextGroupCreate", "10.6")
	registerFunc(&_jSContextGroupRelease, &_jSContextGroupReleaseErr, frameworkHandle, "JSContextGroupRelease", "10.6")
	registerFunc(&_jSContextGroupRetain, &_jSContextGroupRetainErr, frameworkHandle, "JSContextGroupRetain", "10.6")
	registerFunc(&_jSEvaluateScript, &_jSEvaluateScriptErr, frameworkHandle, "JSEvaluateScript", "10.5")
	registerFunc(&_jSGarbageCollect, &_jSGarbageCollectErr, frameworkHandle, "JSGarbageCollect", "10.5")
	registerFunc(&_jSGlobalContextCopyName, &_jSGlobalContextCopyNameErr, frameworkHandle, "JSGlobalContextCopyName", "10.10")
	registerFunc(&_jSGlobalContextCreate, &_jSGlobalContextCreateErr, frameworkHandle, "JSGlobalContextCreate", "10.5")
	registerFunc(&_jSGlobalContextCreateInGroup, &_jSGlobalContextCreateInGroupErr, frameworkHandle, "JSGlobalContextCreateInGroup", "10.6")
	registerFunc(&_jSGlobalContextIsInspectable, &_jSGlobalContextIsInspectableErr, frameworkHandle, "JSGlobalContextIsInspectable", "13.3")
	registerFunc(&_jSGlobalContextRelease, &_jSGlobalContextReleaseErr, frameworkHandle, "JSGlobalContextRelease", "10.5")
	registerFunc(&_jSGlobalContextRetain, &_jSGlobalContextRetainErr, frameworkHandle, "JSGlobalContextRetain", "10.5")
	registerFunc(&_jSGlobalContextSetInspectable, &_jSGlobalContextSetInspectableErr, frameworkHandle, "JSGlobalContextSetInspectable", "13.3")
	registerFunc(&_jSGlobalContextSetName, &_jSGlobalContextSetNameErr, frameworkHandle, "JSGlobalContextSetName", "10.10")
	registerFunc(&_jSObjectCallAsConstructor, &_jSObjectCallAsConstructorErr, frameworkHandle, "JSObjectCallAsConstructor", "10.5")
	registerFunc(&_jSObjectCallAsFunction, &_jSObjectCallAsFunctionErr, frameworkHandle, "JSObjectCallAsFunction", "10.5")
	registerFunc(&_jSObjectCopyPropertyNames, &_jSObjectCopyPropertyNamesErr, frameworkHandle, "JSObjectCopyPropertyNames", "10.5")
	registerFunc(&_jSObjectDeleteProperty, &_jSObjectDeletePropertyErr, frameworkHandle, "JSObjectDeleteProperty", "10.5")
	registerFunc(&_jSObjectDeletePropertyForKey, &_jSObjectDeletePropertyForKeyErr, frameworkHandle, "JSObjectDeletePropertyForKey", "10.15")
	registerFunc(&_jSObjectGetArrayBufferByteLength, &_jSObjectGetArrayBufferByteLengthErr, frameworkHandle, "JSObjectGetArrayBufferByteLength", "10.12")
	registerFunc(&_jSObjectGetArrayBufferBytesPtr, &_jSObjectGetArrayBufferBytesPtrErr, frameworkHandle, "JSObjectGetArrayBufferBytesPtr", "10.12")
	registerFunc(&_jSObjectGetPrivate, &_jSObjectGetPrivateErr, frameworkHandle, "JSObjectGetPrivate", "10.5")
	registerFunc(&_jSObjectGetProperty, &_jSObjectGetPropertyErr, frameworkHandle, "JSObjectGetProperty", "10.5")
	registerFunc(&_jSObjectGetPropertyAtIndex, &_jSObjectGetPropertyAtIndexErr, frameworkHandle, "JSObjectGetPropertyAtIndex", "10.5")
	registerFunc(&_jSObjectGetPropertyForKey, &_jSObjectGetPropertyForKeyErr, frameworkHandle, "JSObjectGetPropertyForKey", "10.15")
	registerFunc(&_jSObjectGetPrototype, &_jSObjectGetPrototypeErr, frameworkHandle, "JSObjectGetPrototype", "10.5")
	registerFunc(&_jSObjectGetTypedArrayBuffer, &_jSObjectGetTypedArrayBufferErr, frameworkHandle, "JSObjectGetTypedArrayBuffer", "10.12")
	registerFunc(&_jSObjectGetTypedArrayByteLength, &_jSObjectGetTypedArrayByteLengthErr, frameworkHandle, "JSObjectGetTypedArrayByteLength", "10.12")
	registerFunc(&_jSObjectGetTypedArrayByteOffset, &_jSObjectGetTypedArrayByteOffsetErr, frameworkHandle, "JSObjectGetTypedArrayByteOffset", "10.12")
	registerFunc(&_jSObjectGetTypedArrayBytesPtr, &_jSObjectGetTypedArrayBytesPtrErr, frameworkHandle, "JSObjectGetTypedArrayBytesPtr", "10.12")
	registerFunc(&_jSObjectGetTypedArrayLength, &_jSObjectGetTypedArrayLengthErr, frameworkHandle, "JSObjectGetTypedArrayLength", "10.12")
	registerFunc(&_jSObjectHasProperty, &_jSObjectHasPropertyErr, frameworkHandle, "JSObjectHasProperty", "10.5")
	registerFunc(&_jSObjectHasPropertyForKey, &_jSObjectHasPropertyForKeyErr, frameworkHandle, "JSObjectHasPropertyForKey", "10.15")
	registerFunc(&_jSObjectIsConstructor, &_jSObjectIsConstructorErr, frameworkHandle, "JSObjectIsConstructor", "10.5")
	registerFunc(&_jSObjectIsFunction, &_jSObjectIsFunctionErr, frameworkHandle, "JSObjectIsFunction", "10.5")
	registerFunc(&_jSObjectMake, &_jSObjectMakeErr, frameworkHandle, "JSObjectMake", "10.5")
	registerFunc(&_jSObjectMakeArray, &_jSObjectMakeArrayErr, frameworkHandle, "JSObjectMakeArray", "10.6")
	registerFunc(&_jSObjectMakeArrayBufferWithBytesNoCopy, &_jSObjectMakeArrayBufferWithBytesNoCopyErr, frameworkHandle, "JSObjectMakeArrayBufferWithBytesNoCopy", "10.12")
	registerFunc(&_jSObjectMakeConstructor, &_jSObjectMakeConstructorErr, frameworkHandle, "JSObjectMakeConstructor", "10.5")
	registerFunc(&_jSObjectMakeDate, &_jSObjectMakeDateErr, frameworkHandle, "JSObjectMakeDate", "10.6")
	registerFunc(&_jSObjectMakeDeferredPromise, &_jSObjectMakeDeferredPromiseErr, frameworkHandle, "JSObjectMakeDeferredPromise", "10.15")
	registerFunc(&_jSObjectMakeError, &_jSObjectMakeErrorErr, frameworkHandle, "JSObjectMakeError", "10.6")
	registerFunc(&_jSObjectMakeFunction, &_jSObjectMakeFunctionErr, frameworkHandle, "JSObjectMakeFunction", "10.5")
	registerFunc(&_jSObjectMakeFunctionWithCallback, &_jSObjectMakeFunctionWithCallbackErr, frameworkHandle, "JSObjectMakeFunctionWithCallback", "10.5")
	registerFunc(&_jSObjectMakeRegExp, &_jSObjectMakeRegExpErr, frameworkHandle, "JSObjectMakeRegExp", "10.6")
	registerFunc(&_jSObjectMakeTypedArray, &_jSObjectMakeTypedArrayErr, frameworkHandle, "JSObjectMakeTypedArray", "10.12")
	registerFunc(&_jSObjectMakeTypedArrayWithArrayBuffer, &_jSObjectMakeTypedArrayWithArrayBufferErr, frameworkHandle, "JSObjectMakeTypedArrayWithArrayBuffer", "10.12")
	registerFunc(&_jSObjectMakeTypedArrayWithArrayBufferAndOffset, &_jSObjectMakeTypedArrayWithArrayBufferAndOffsetErr, frameworkHandle, "JSObjectMakeTypedArrayWithArrayBufferAndOffset", "10.12")
	registerFunc(&_jSObjectMakeTypedArrayWithBytesNoCopy, &_jSObjectMakeTypedArrayWithBytesNoCopyErr, frameworkHandle, "JSObjectMakeTypedArrayWithBytesNoCopy", "10.12")
	registerFunc(&_jSObjectSetPrivate, &_jSObjectSetPrivateErr, frameworkHandle, "JSObjectSetPrivate", "10.5")
	registerFunc(&_jSObjectSetProperty, &_jSObjectSetPropertyErr, frameworkHandle, "JSObjectSetProperty", "10.5")
	registerFunc(&_jSObjectSetPropertyAtIndex, &_jSObjectSetPropertyAtIndexErr, frameworkHandle, "JSObjectSetPropertyAtIndex", "10.5")
	registerFunc(&_jSObjectSetPropertyForKey, &_jSObjectSetPropertyForKeyErr, frameworkHandle, "JSObjectSetPropertyForKey", "10.15")
	registerFunc(&_jSObjectSetPrototype, &_jSObjectSetPrototypeErr, frameworkHandle, "JSObjectSetPrototype", "10.5")
	registerFunc(&_jSPropertyNameAccumulatorAddName, &_jSPropertyNameAccumulatorAddNameErr, frameworkHandle, "JSPropertyNameAccumulatorAddName", "10.5")
	registerFunc(&_jSPropertyNameArrayGetCount, &_jSPropertyNameArrayGetCountErr, frameworkHandle, "JSPropertyNameArrayGetCount", "10.5")
	registerFunc(&_jSPropertyNameArrayGetNameAtIndex, &_jSPropertyNameArrayGetNameAtIndexErr, frameworkHandle, "JSPropertyNameArrayGetNameAtIndex", "10.5")
	registerFunc(&_jSPropertyNameArrayRelease, &_jSPropertyNameArrayReleaseErr, frameworkHandle, "JSPropertyNameArrayRelease", "10.5")
	registerFunc(&_jSPropertyNameArrayRetain, &_jSPropertyNameArrayRetainErr, frameworkHandle, "JSPropertyNameArrayRetain", "10.5")
	registerFunc(&_jSStringCopyCFString, &_jSStringCopyCFStringErr, frameworkHandle, "JSStringCopyCFString", "10.5")
	registerFunc(&_jSStringCreateWithCFString, &_jSStringCreateWithCFStringErr, frameworkHandle, "JSStringCreateWithCFString", "10.5")
	registerFunc(&_jSStringCreateWithCharacters, &_jSStringCreateWithCharactersErr, frameworkHandle, "JSStringCreateWithCharacters", "10.5")
	registerFunc(&_jSStringCreateWithUTF8CString, &_jSStringCreateWithUTF8CStringErr, frameworkHandle, "JSStringCreateWithUTF8CString", "10.5")
	registerFunc(&_jSStringGetCharactersPtr, &_jSStringGetCharactersPtrErr, frameworkHandle, "JSStringGetCharactersPtr", "10.5")
	registerFunc(&_jSStringGetLength, &_jSStringGetLengthErr, frameworkHandle, "JSStringGetLength", "10.5")
	registerFunc(&_jSStringGetMaximumUTF8CStringSize, &_jSStringGetMaximumUTF8CStringSizeErr, frameworkHandle, "JSStringGetMaximumUTF8CStringSize", "10.5")
	registerFunc(&_jSStringGetUTF8CString, &_jSStringGetUTF8CStringErr, frameworkHandle, "JSStringGetUTF8CString", "10.5")
	registerFunc(&_jSStringIsEqual, &_jSStringIsEqualErr, frameworkHandle, "JSStringIsEqual", "10.5")
	registerFunc(&_jSStringIsEqualToUTF8CString, &_jSStringIsEqualToUTF8CStringErr, frameworkHandle, "JSStringIsEqualToUTF8CString", "10.5")
	registerFunc(&_jSStringRelease, &_jSStringReleaseErr, frameworkHandle, "JSStringRelease", "10.5")
	registerFunc(&_jSStringRetain, &_jSStringRetainErr, frameworkHandle, "JSStringRetain", "10.5")
	registerFunc(&_jSValueCompare, &_jSValueCompareErr, frameworkHandle, "JSValueCompare", "15.0")
	registerFunc(&_jSValueCompareDouble, &_jSValueCompareDoubleErr, frameworkHandle, "JSValueCompareDouble", "15.0")
	registerFunc(&_jSValueCompareInt64, &_jSValueCompareInt64Err, frameworkHandle, "JSValueCompareInt64", "15.0")
	registerFunc(&_jSValueCompareUInt64, &_jSValueCompareUInt64Err, frameworkHandle, "JSValueCompareUInt64", "15.0")
	registerFunc(&_jSValueCreateJSONString, &_jSValueCreateJSONStringErr, frameworkHandle, "JSValueCreateJSONString", "10.7")
	registerFunc(&_jSValueGetType, &_jSValueGetTypeErr, frameworkHandle, "JSValueGetType", "10.5")
	registerFunc(&_jSValueGetTypedArrayType, &_jSValueGetTypedArrayTypeErr, frameworkHandle, "JSValueGetTypedArrayType", "10.12")
	registerFunc(&_jSValueIsArray, &_jSValueIsArrayErr, frameworkHandle, "JSValueIsArray", "10.11")
	registerFunc(&_jSValueIsBigInt, &_jSValueIsBigIntErr, frameworkHandle, "JSValueIsBigInt", "15.0")
	registerFunc(&_jSValueIsBoolean, &_jSValueIsBooleanErr, frameworkHandle, "JSValueIsBoolean", "10.5")
	registerFunc(&_jSValueIsDate, &_jSValueIsDateErr, frameworkHandle, "JSValueIsDate", "10.11")
	registerFunc(&_jSValueIsEqual, &_jSValueIsEqualErr, frameworkHandle, "JSValueIsEqual", "10.5")
	registerFunc(&_jSValueIsInstanceOfConstructor, &_jSValueIsInstanceOfConstructorErr, frameworkHandle, "JSValueIsInstanceOfConstructor", "10.5")
	registerFunc(&_jSValueIsNull, &_jSValueIsNullErr, frameworkHandle, "JSValueIsNull", "10.5")
	registerFunc(&_jSValueIsNumber, &_jSValueIsNumberErr, frameworkHandle, "JSValueIsNumber", "10.5")
	registerFunc(&_jSValueIsObject, &_jSValueIsObjectErr, frameworkHandle, "JSValueIsObject", "10.5")
	registerFunc(&_jSValueIsObjectOfClass, &_jSValueIsObjectOfClassErr, frameworkHandle, "JSValueIsObjectOfClass", "10.5")
	registerFunc(&_jSValueIsStrictEqual, &_jSValueIsStrictEqualErr, frameworkHandle, "JSValueIsStrictEqual", "10.5")
	registerFunc(&_jSValueIsString, &_jSValueIsStringErr, frameworkHandle, "JSValueIsString", "10.5")
	registerFunc(&_jSValueIsSymbol, &_jSValueIsSymbolErr, frameworkHandle, "JSValueIsSymbol", "10.15")
	registerFunc(&_jSValueIsUndefined, &_jSValueIsUndefinedErr, frameworkHandle, "JSValueIsUndefined", "10.5")
	registerFunc(&_jSValueMakeBoolean, &_jSValueMakeBooleanErr, frameworkHandle, "JSValueMakeBoolean", "10.5")
	registerFunc(&_jSValueMakeFromJSONString, &_jSValueMakeFromJSONStringErr, frameworkHandle, "JSValueMakeFromJSONString", "10.7")
	registerFunc(&_jSValueMakeNull, &_jSValueMakeNullErr, frameworkHandle, "JSValueMakeNull", "10.5")
	registerFunc(&_jSValueMakeNumber, &_jSValueMakeNumberErr, frameworkHandle, "JSValueMakeNumber", "10.5")
	registerFunc(&_jSValueMakeString, &_jSValueMakeStringErr, frameworkHandle, "JSValueMakeString", "10.5")
	registerFunc(&_jSValueMakeSymbol, &_jSValueMakeSymbolErr, frameworkHandle, "JSValueMakeSymbol", "10.15")
	registerFunc(&_jSValueMakeUndefined, &_jSValueMakeUndefinedErr, frameworkHandle, "JSValueMakeUndefined", "10.5")
	registerFunc(&_jSValueProtect, &_jSValueProtectErr, frameworkHandle, "JSValueProtect", "10.5")
	registerFunc(&_jSValueToBoolean, &_jSValueToBooleanErr, frameworkHandle, "JSValueToBoolean", "10.5")
	registerFunc(&_jSValueToInt32, &_jSValueToInt32Err, frameworkHandle, "JSValueToInt32", "15.0")
	registerFunc(&_jSValueToInt64, &_jSValueToInt64Err, frameworkHandle, "JSValueToInt64", "15.0")
	registerFunc(&_jSValueToNumber, &_jSValueToNumberErr, frameworkHandle, "JSValueToNumber", "10.5")
	registerFunc(&_jSValueToObject, &_jSValueToObjectErr, frameworkHandle, "JSValueToObject", "10.5")
	registerFunc(&_jSValueToStringCopy, &_jSValueToStringCopyErr, frameworkHandle, "JSValueToStringCopy", "10.5")
	registerFunc(&_jSValueToUInt32, &_jSValueToUInt32Err, frameworkHandle, "JSValueToUInt32", "15.0")
	registerFunc(&_jSValueToUInt64, &_jSValueToUInt64Err, frameworkHandle, "JSValueToUInt64", "15.0")
	registerFunc(&_jSValueUnprotect, &_jSValueUnprotectErr, frameworkHandle, "JSValueUnprotect", "10.5")
}
