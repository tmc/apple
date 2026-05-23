// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

package javascriptcore

import (
	"sync"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [JSContext] class.
var (
	_JSContextClass     JSContextClass
	_JSContextClassOnce sync.Once
)

func getJSContextClass() JSContextClass {
	_JSContextClassOnce.Do(func() {
		_JSContextClass = JSContextClass{class: objc.GetClass("JSContext")}
	})
	return _JSContextClass
}

// GetJSContextClass returns the class object for JSContext.
func GetJSContextClass() JSContextClass {
	return getJSContextClass()
}

type JSContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (jc JSContextClass) Class() objc.Class {
	return jc.class
}

// Alloc allocates memory for a new instance of the class.
func (jc JSContextClass) Alloc() JSContext {
	rv := objc.Send[JSContext](objc.ID(jc.class), objc.Sel("alloc"))
	return rv
}

// A JavaScript execution environment.
//
// # Overview
//
// You create and use JavaScript contexts to evaluate JavaScript scripts from
// Objective-C or Swift code; to access values that JavaScript defines or
// calculates; and to make native objects, methods, or functions accessible to
// JavaScript.
//
// # Creating JavaScript contexts
//
//   - [JSContext.InitWithVirtualMachine]: Creates a new JavaScript context associated with a specific virtual machine.
//
// # Making JavaScript context inspectable
//
//   - [JSContext.IsInspectable]: A Boolean value that indicates whether you can inspect the JavaScript context with Safari Web Inspector.
//   - [JSContext.SetInspectable]
//
// # Evaluating scripts
//
//   - [JSContext.EvaluateScript]: Executes the specified JavaScript code.
//   - [JSContext.EvaluateScriptWithSourceURL]: Executes the specified JavaScript code, treating the specified URL as its source location.
//
// # Working with JavaScript global state
//
//   - [JSContext.GlobalObject]: The JavaScript global object associated with the context.
//   - [JSContext.Exception]: A JavaScript exception to be thrown in evaluation of the script.
//   - [JSContext.SetException]
//   - [JSContext.ExceptionHandler]: A block to be invoked should evaluating a script result in a JavaScript exception being thrown.
//   - [JSContext.SetExceptionHandler]
//   - [JSContext.VirtualMachine]: The JavaScript virtual machine to which the context belongs.
//   - [JSContext.Name]: A descriptive name for the context.
//   - [JSContext.SetName]
//
// # Accessing JavaScript global state with subscripts
//
//   - [JSContext.ObjectForKeyedSubscript]: Returns the value of the specified JavaScript property in the context’s global object, allowing subscript getter syntax.
//   - [JSContext.SetObjectForKeyedSubscript]: Sets the specified JavaScript property of the context’s global object, allowing subscript setter syntax.
//
// # Working with the C JavaScriptCore API
//
//   - [JSContext.JSGlobalContextRef]: Returns the C representation of the JavaScript context.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContext
type JSContext struct {
	objectivec.Object
}

// JSContextFromID constructs a [JSContext] from an objc.ID.
//
// A JavaScript execution environment.
func JSContextFromID(id objc.ID) JSContext {
	return JSContext{objectivec.Object{ID: id}}
}

// NOTE: JSContext adopts protocols; skip strict compile-time interface assertion.
// Protocol method surfaces are generated separately and may include optional methods.

// An interface definition for the [JSContext] class.
//
// # Creating JavaScript contexts
//
//   - [IJSContext.InitWithVirtualMachine]: Creates a new JavaScript context associated with a specific virtual machine.
//
// # Making JavaScript context inspectable
//
//   - [IJSContext.IsInspectable]: A Boolean value that indicates whether you can inspect the JavaScript context with Safari Web Inspector.
//   - [IJSContext.SetInspectable]
//
// # Evaluating scripts
//
//   - [IJSContext.EvaluateScript]: Executes the specified JavaScript code.
//   - [IJSContext.EvaluateScriptWithSourceURL]: Executes the specified JavaScript code, treating the specified URL as its source location.
//
// # Working with JavaScript global state
//
//   - [IJSContext.GlobalObject]: The JavaScript global object associated with the context.
//   - [IJSContext.Exception]: A JavaScript exception to be thrown in evaluation of the script.
//   - [IJSContext.SetException]
//   - [IJSContext.ExceptionHandler]: A block to be invoked should evaluating a script result in a JavaScript exception being thrown.
//   - [IJSContext.SetExceptionHandler]
//   - [IJSContext.VirtualMachine]: The JavaScript virtual machine to which the context belongs.
//   - [IJSContext.Name]: A descriptive name for the context.
//   - [IJSContext.SetName]
//
// # Accessing JavaScript global state with subscripts
//
//   - [IJSContext.ObjectForKeyedSubscript]: Returns the value of the specified JavaScript property in the context’s global object, allowing subscript getter syntax.
//   - [IJSContext.SetObjectForKeyedSubscript]: Sets the specified JavaScript property of the context’s global object, allowing subscript setter syntax.
//
// # Working with the C JavaScriptCore API
//
//   - [IJSContext.JSGlobalContextRef]: Returns the C representation of the JavaScript context.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContext
type IJSContext interface {
	objectivec.IObject

	// Topic: Creating JavaScript contexts

	// Creates a new JavaScript context associated with a specific virtual machine.
	InitWithVirtualMachine(virtualMachine IJSVirtualMachine) JSContext

	// Topic: Making JavaScript context inspectable

	// A Boolean value that indicates whether you can inspect the JavaScript context with Safari Web Inspector.
	IsInspectable() bool
	SetInspectable(value bool)

	// Topic: Evaluating scripts

	// Executes the specified JavaScript code.
	EvaluateScript(script string) IJSValue
	// Executes the specified JavaScript code, treating the specified URL as its source location.
	EvaluateScriptWithSourceURL(script string, sourceURL foundation.NSURL) IJSValue

	// Topic: Working with JavaScript global state

	// The JavaScript global object associated with the context.
	GlobalObject() IJSValue
	// A JavaScript exception to be thrown in evaluation of the script.
	Exception() IJSValue
	SetException(value IJSValue)
	// A block to be invoked should evaluating a script result in a JavaScript exception being thrown.
	ExceptionHandler() JSContextJSValueHandler
	SetExceptionHandler(value JSContextJSValueHandler)
	// The JavaScript virtual machine to which the context belongs.
	VirtualMachine() IJSVirtualMachine
	// A descriptive name for the context.
	Name() string
	SetName(value string)

	// Topic: Accessing JavaScript global state with subscripts

	// Returns the value of the specified JavaScript property in the context’s global object, allowing subscript getter syntax.
	ObjectForKeyedSubscript(key objectivec.IObject) IJSValue
	// Sets the specified JavaScript property of the context’s global object, allowing subscript setter syntax.
	SetObjectForKeyedSubscript(object objectivec.IObject, key objectivec.NSObject)

	// Topic: Working with the C JavaScriptCore API

	// Returns the C representation of the JavaScript context.
	JSGlobalContextRef() JSGlobalContextRef
}

// Init initializes the instance.
func (j JSContext) Init() JSContext {
	rv := objc.Send[JSContext](j.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (j JSContext) Autorelease() JSContext {
	rv := objc.Send[JSContext](j.ID, objc.Sel("autorelease"))
	return rv
}

// NewJSContext creates a new JSContext instance.
func NewJSContext() JSContext {
	class := getJSContextClass()
	rv := objc.Send[JSContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// Creates a JavaScript context object from the equivalent C representation.
//
// jsGlobalContextRef: A C JavaScript context reference.
//
// # Return Value
//
// A JavaScript context object representing the same context.
//
// # Discussion
//
// See [JSContextRef] for the C JavaScriptCore API.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContext/init(JSGlobalContextRef:)-9m51
func NewJSContextWithJSGlobalContextRef(jsGlobalContextRef JSGlobalContextRef) JSContext {
	rv := objc.Send[objc.ID](objc.ID(getJSContextClass().class), objc.Sel("contextWithJSGlobalContextRef:"), jsGlobalContextRef)
	return JSContextFromID(rv)
}

// Creates a new JavaScript context associated with a specific virtual
// machine.
//
// virtualMachine: The virtual machine with which to associate the new context.
//
// # Return Value
//
// A new JavaScript context.
//
// # Discussion
//
// By default, each context has an independent virtual machine (a
// [JSVirtualMachine] object). You cannot pass JavaScript values between
// contexts in different virtual machines. Use this initializer to create a
// context that shares its virtual machine with other JavaScript contexts to
// allow passing [JSValue] objects between those contexts.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContext/init(virtualMachine:)
func NewJSContextWithVirtualMachine(virtualMachine IJSVirtualMachine) JSContext {
	instance := getJSContextClass().Alloc()
	rv := objc.Send[objc.ID](instance.ID, objc.Sel("initWithVirtualMachine:"), virtualMachine)
	return JSContextFromID(rv)
}

// Creates a new JavaScript context associated with a specific virtual
// machine.
//
// virtualMachine: The virtual machine with which to associate the new context.
//
// # Return Value
//
// A new JavaScript context.
//
// # Discussion
//
// By default, each context has an independent virtual machine (a
// [JSVirtualMachine] object). You cannot pass JavaScript values between
// contexts in different virtual machines. Use this initializer to create a
// context that shares its virtual machine with other JavaScript contexts to
// allow passing [JSValue] objects between those contexts.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContext/init(virtualMachine:)
func (j JSContext) InitWithVirtualMachine(virtualMachine IJSVirtualMachine) JSContext {
	rv := objc.Send[JSContext](j.ID, objc.Sel("initWithVirtualMachine:"), virtualMachine)
	return rv
}

// Executes the specified JavaScript code.
//
// script: The JavaScript source code to evaluate.
//
// # Return Value
//
// The last value generated by the script. Note that a script can result in
// the JavaScript value `undefined`.
//
// # Discussion
//
// Evaluating a script runs any top-level code and adds function and object
// definitions to the context’s global object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContext/evaluateScript(_:)
func (j JSContext) EvaluateScript(script string) IJSValue {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("evaluateScript:"), objc.String(script))
	return JSValueFromID(rv)
}

// Executes the specified JavaScript code, treating the specified URL as its
// source location.
//
// script: The JavaScript source code to evaluate.
//
// sourceURL: A URL to be considered as the script’s origin.
//
// # Return Value
//
// The last value generated by the script. Note that a script can result in
// the JavaScript value `undefined`.
//
// # Discussion
//
// Evaluating a script runs any top-level code and adds function or object
// definitions to the context’s global object.
//
// The `sourceURL` parameter is informative only; debuggers may use this URL
// when reporting exceptions.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContext/evaluateScript(_:withSourceURL:)
func (j JSContext) EvaluateScriptWithSourceURL(script string, sourceURL foundation.NSURL) IJSValue {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("evaluateScript:withSourceURL:"), objc.String(script), sourceURL)
	return JSValueFromID(rv)
}

// Returns the value of the specified JavaScript property in the context’s
// global object, allowing subscript getter syntax.
//
// key: The name of a JavaScript property in the context’s global JavaScript
// object.
//
// # Return Value
//
// The JavaScript property named by `key`, or `nil` if no such field or
// function exists.
//
// # Discussion
//
// This method first constructs a [JSValue] object from the `key` parameter,
// then uses that value in JavaScript to look up the name of a property in the
// context’s global object.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContext/objectForKeyedSubscript(_:)
func (j JSContext) ObjectForKeyedSubscript(key objectivec.IObject) IJSValue {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("objectForKeyedSubscript:"), key)
	return JSValueFromID(rv)
}

// Sets the specified JavaScript property of the context’s global object,
// allowing subscript setter syntax.
//
// object: The value to set for the JavaScript property.
//
// key: The JavaScript property name to use in the context’s global JavaScript
// object.
//
// # Discussion
//
// This method first constructs a [JSValue] object from the `key` parameter,
// then uses that value in JavaScript to set the property in the context’s
// global object.
//
// Use this method (or Objective-C subscript syntax) to bridge native objects
// or functions for use in JavaScript. For example, the following code creates
// a JavaScript function whose implementation is an Objective-C block:
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContext/setObject(_:forKeyedSubscript:)
func (j JSContext) SetObjectForKeyedSubscript(object objectivec.IObject, key objectivec.NSObject) {
	objc.Send[objc.ID](j.ID, objc.Sel("setObject:forKeyedSubscript:"), object, key)
}

// Returns the context currently executing JavaScript code.
//
// # Return Value
//
// The currently executing context, or `nil` if not within native code called
// from JavaScript.
//
// # Discussion
//
// Call this method within an Objective-C or Swift block or method invoked
// from within JavaScript to obtain the [JSContext] object responsible for
// executing that Javascript code.
//
// If not currently in code invoked as a callback from JavaScript, this method
// returns `nil`.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContext/current()
func (_JSContextClass JSContextClass) CurrentContext() JSContext {
	rv := objc.Send[objc.ID](objc.ID(_JSContextClass.class), objc.Sel("currentContext"))
	return JSContextFromID(rv)
}

// Returns the currently executing JavaScript function.
//
// # Return Value
//
// The currently executing JavaScript function, or `nil` if not within native
// code called from JavaScript.
//
// # Discussion
//
// Call this method within an Objective-C or Swift block or method invoked
// from within JavaScript to obtain a [JSValue] object representing the
// JavaScript function responsible for executing that code.
//
// If not currently in code invoked as a callback from JavaScript, this method
// returns `nil`.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContext/currentCallee()
func (_JSContextClass JSContextClass) CurrentCallee() JSValue {
	rv := objc.Send[objc.ID](objc.ID(_JSContextClass.class), objc.Sel("currentCallee"))
	return JSValueFromID(rv)
}

// Returns the value of the `this` keyword in currently executing JavaScript
// code.
//
// # Return Value
//
// The current value of the JavaScript `this` keyword, or `nil` if not within
// native code called from JavaScript.
//
// # Discussion
//
// Call this method within an Objective-C or Swift block or method invoked
// from within JavaScript to obtain a [JSValue] object representing the
// current value of the `this` keyword in that JavaScript code.
//
// If not currently in code invoked as a callback from JavaScript, this method
// returns `nil`.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContext/currentThis()
func (_JSContextClass JSContextClass) CurrentThis() JSValue {
	rv := objc.Send[objc.ID](objc.ID(_JSContextClass.class), objc.Sel("currentThis"))
	return JSValueFromID(rv)
}

// Returns the arguments to the current native callback from JavaScript code.
//
// # Return Value
//
// The current callback arguments, or `nil` if not within native code called
// from JavaScript.
//
// # Discussion
//
// Call this method within an Objective-C or Swift block or method invoked
// from within JavaScript to obtain an array of [JSValue] objects representing
// the arguments to the JavaScript function responsible for that callback.
//
// If not currently in code invoked as a callback from JavaScript, this method
// returns `nil`.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContext/currentArguments()
func (_JSContextClass JSContextClass) CurrentArguments() foundation.INSArray {
	rv := objc.Send[objc.ID](objc.ID(_JSContextClass.class), objc.Sel("currentArguments"))
	return foundation.NSArrayFromID(rv)
}

// A Boolean value that indicates whether you can inspect the JavaScript
// context with Safari Web Inspector.
//
// # Discussion
//
// Defaults to `false`.
//
// Set to `true` at any point in the context’s lifetime to allow Safari Web
// Inspector access to inspect the context. Then, select your context in
// Safari’s Develop menu for either your computer or an attached device to
// inspect it.
//
// If you set this value to `false` during inspection, the system immediately
// closes Safari Web Inspector and does not provide any further information
// about the context.
//
// For more information, see [Enabling the Inspection of Web Content in Apps].
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContext/isInspectable
//
// [Enabling the Inspection of Web Content in Apps]: https://webkit.org/blog/13936/enabling-the-inspection-of-web-content-in-apps/
func (j JSContext) IsInspectable() bool {
	rv := objc.Send[bool](j.ID, objc.Sel("isInspectable"))
	return rv
}
func (j JSContext) SetInspectable(value bool) {
	objc.Send[struct{}](j.ID, objc.Sel("setInspectable:"), value)
}

// The JavaScript global object associated with the context.
//
// # Discussion
//
// In a web browser, the global object of a JavaScript context is the browser
// window (the `window` object in JavaScript). Outside of web-browser use, a
// context’s global object serves a similar role, separating the JavaScript
// namespaces of different contexts. Global variables within a script appear
// as fields or subscripts in the global object—you can access them either
// through this [JSValue] object or through the methods listed in the
// Accessing JavaScript global state with subscripts section in [JSContext].
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContext/globalObject
func (j JSContext) GlobalObject() IJSValue {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("globalObject"))
	return JSValueFromID(objc.ID(rv))
}

// A JavaScript exception to be thrown in evaluation of the script.
//
// # Discussion
//
// Before performing a callback from JavaScript to an Objective-C or Swift
// block or method, the context preserves the prior value of this property and
// then sets its value to `nil`. After the callback has completed, the context
// reads the new value of the [JSContext.Exception] property—if this value
// is not nil, the context treats the value as an exception to be thrown in
// JavaScript as a result of the callback. After reading the property (and
// possibly throwing a JavaScript exception), the context restores the prior
// value of this property.
//
// By default, JavaScriptCore assigns any uncaught exception to this property,
// so you can check this property’s value to find uncaught exceptions
// arising from JavaScript function calls. To change the exception handling
// behavior, use the [JSContext.ExceptionHandler] property.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContext/exception
func (j JSContext) Exception() IJSValue {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("exception"))
	return JSValueFromID(objc.ID(rv))
}
func (j JSContext) SetException(value IJSValue) {
	objc.Send[struct{}](j.ID, objc.Sel("setException:"), value)
}

// A block to be invoked should evaluating a script result in a JavaScript
// exception being thrown.
//
// # Discussion
//
// The block takes the following parameters:
//
// context: The context in which the exception originates. exception: The
// JavaScript exception thrown.
//
// The default value exception handler block stores its `exception` parameter
// value into the context’s [JSContext.Exception] property. As a
// consequence, the default behavior is that unhandled exceptions occurring
// within a callback from JavaScript to native code are thrown again upon
// return. Setting this value to `nil` results in all uncaught exceptions
// being silently consumed.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContext/exceptionHandler
func (j JSContext) ExceptionHandler() JSContextJSValueHandler {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("exceptionHandler"))
	_ = rv
	return nil
}
func (j JSContext) SetExceptionHandler(value JSContextJSValueHandler) {
	block, cleanup := NewJSContextJSValueBlock(value)
	defer cleanup()
	objc.Send[struct{}](j.ID, objc.Sel("setExceptionHandler:"), block)
}

// The JavaScript virtual machine to which the context belongs.
//
// # Discussion
//
// To create a context associated with a specific virtual machine, allowing
// JavaScript values to be passed between contexts that share the same virtual
// machine, use the [JSContext.InitWithVirtualMachine] initializer.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContext/virtualMachine
func (j JSContext) VirtualMachine() IJSVirtualMachine {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("virtualMachine"))
	return JSVirtualMachineFromID(objc.ID(rv))
}

// A descriptive name for the context.
//
// # Discussion
//
// This name appears when using remote debugging to examine the context.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContext/name
func (j JSContext) Name() string {
	rv := objc.Send[objc.ID](j.ID, objc.Sel("name"))
	return foundation.NSStringFromID(rv).String()
}
func (j JSContext) SetName(value string) {
	objc.Send[struct{}](j.ID, objc.Sel("setName:"), objc.String(value))
}

// Returns the C representation of the JavaScript context.
//
// # Discussion
//
// See [JSContextRef] for the C JavaScriptCore API.
//
// See: https://developer.apple.com/documentation/JavaScriptCore/JSContext/jsGlobalContextRef
func (j JSContext) JSGlobalContextRef() JSGlobalContextRef {
	rv := objc.Send[JSGlobalContextRef](j.ID, objc.Sel("JSGlobalContextRef"))
	return JSGlobalContextRef(rv)
}
