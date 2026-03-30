// Code generated from Apple documentation for gtshaderprofiler. DO NOT EDIT.

package gtshaderprofiler

import (
	"context"
	"sync"
	"unsafe"

	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
	"github.com/tmc/apple/objectivec"
)

// The class instance for the [GTJSScriptingContext] class.
var (
	_GTJSScriptingContextClass     GTJSScriptingContextClass
	_GTJSScriptingContextClassOnce sync.Once
)

func getGTJSScriptingContextClass() GTJSScriptingContextClass {
	_GTJSScriptingContextClassOnce.Do(func() {
		_GTJSScriptingContextClass = GTJSScriptingContextClass{class: objc.GetClass("GTJSScriptingContext")}
	})
	return _GTJSScriptingContextClass
}

// GetGTJSScriptingContextClass returns the class object for GTJSScriptingContext.
func GetGTJSScriptingContextClass() GTJSScriptingContextClass {
	return getGTJSScriptingContextClass()
}

type GTJSScriptingContextClass struct {
	class objc.Class
}

// Class returns the underlying Objective-C class pointer.
func (gc GTJSScriptingContextClass) Class() objc.Class {
	return gc.class
}

// Alloc allocates memory for a new instance of the class.
func (gc GTJSScriptingContextClass) Alloc() GTJSScriptingContext {
	rv := objc.Send[GTJSScriptingContext](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// # Methods
//
//   - [GTJSScriptingContext._cachedStringFromString]
//   - [GTJSScriptingContext._clearCache]
//   - [GTJSScriptingContext._jsStringToString]
//   - [GTJSScriptingContext._jsValueToString]
//   - [GTJSScriptingContext.AllocNewContext]
//   - [GTJSScriptingContext.CallFunctionWithArguments]
//   - [GTJSScriptingContext.CallGlobalFunction]
//   - [GTJSScriptingContext.Context]
//   - [GTJSScriptingContext.CreateArrayRef]
//   - [GTJSScriptingContext.EvaluteScriptScriptURL]
//   - [GTJSScriptingContext.GetGlobalDouble]
//   - [GTJSScriptingContext.GetGlobalJSONObject]
//   - [GTJSScriptingContext.GetValue]
//   - [GTJSScriptingContext.SetExceptionHandler]
//   - [GTJSScriptingContext.SetGlobalDoubleValue]
//   - [GTJSScriptingContext.SetGlobalJSONObjectValue]
//   - [GTJSScriptingContext.SetRawArrayValuesWithDoubleValuesAndNumCounters]
//   - [GTJSScriptingContext.SetRawArrayValuesWithUint32ValuesAndNumCounters]
//   - [GTJSScriptingContext.SetRawArrayValuesWithUint64ValuesAndNumCounters]
//   - [GTJSScriptingContext.SetValueValue]
//   - [GTJSScriptingContext.SetValues]
//   - [GTJSScriptingContext.VirtualMachine]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext
type GTJSScriptingContext struct {
	objectivec.Object
}

// GTJSScriptingContextFromID constructs a [GTJSScriptingContext] from an objc.ID.
func GTJSScriptingContextFromID(id objc.ID) GTJSScriptingContext {
	return GTJSScriptingContext{objectivec.Object{ID: id}}
}

// Ensure GTJSScriptingContext implements IGTJSScriptingContext.
var _ IGTJSScriptingContext = GTJSScriptingContext{}

// An interface definition for the [GTJSScriptingContext] class.
//
// # Methods
//
//   - [IGTJSScriptingContext._cachedStringFromString]
//   - [IGTJSScriptingContext._clearCache]
//   - [IGTJSScriptingContext._jsStringToString]
//   - [IGTJSScriptingContext._jsValueToString]
//   - [IGTJSScriptingContext.AllocNewContext]
//   - [IGTJSScriptingContext.CallFunctionWithArguments]
//   - [IGTJSScriptingContext.CallGlobalFunction]
//   - [IGTJSScriptingContext.Context]
//   - [IGTJSScriptingContext.CreateArrayRef]
//   - [IGTJSScriptingContext.EvaluteScriptScriptURL]
//   - [IGTJSScriptingContext.GetGlobalDouble]
//   - [IGTJSScriptingContext.GetGlobalJSONObject]
//   - [IGTJSScriptingContext.GetValue]
//   - [IGTJSScriptingContext.SetExceptionHandler]
//   - [IGTJSScriptingContext.SetGlobalDoubleValue]
//   - [IGTJSScriptingContext.SetGlobalJSONObjectValue]
//   - [IGTJSScriptingContext.SetRawArrayValuesWithDoubleValuesAndNumCounters]
//   - [IGTJSScriptingContext.SetRawArrayValuesWithUint32ValuesAndNumCounters]
//   - [IGTJSScriptingContext.SetRawArrayValuesWithUint64ValuesAndNumCounters]
//   - [IGTJSScriptingContext.SetValueValue]
//   - [IGTJSScriptingContext.SetValues]
//   - [IGTJSScriptingContext.VirtualMachine]
//
// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext
type IGTJSScriptingContext interface {
	objectivec.IObject

	// Topic: Methods

	_cachedStringFromString(string_ string) unsafe.Pointer
	_clearCache()
	_jsStringToString(string_ unsafe.Pointer) objectivec.IObject
	_jsValueToString(string_ unsafe.Pointer) objectivec.IObject
	AllocNewContext()
	CallFunctionWithArguments(function objectivec.IObject, arguments objectivec.IObject) float64
	CallGlobalFunction(function string) float64
	Context() unsafe.Pointer
	CreateArrayRef(ref objectivec.IObject) unsafe.Pointer
	EvaluteScriptScriptURL(script objectivec.IObject, url foundation.INSURL) bool
	GetGlobalDouble(double string) float64
	GetGlobalJSONObject(jSONObject string) objectivec.IObject
	GetValue(value objectivec.IObject) objectivec.IObject
	SetExceptionHandler(handler VoidHandler)
	SetGlobalDoubleValue(double string, value float64)
	SetGlobalJSONObjectValue(jSONObject string, value objectivec.IObject) bool
	SetRawArrayValuesWithDoubleValuesAndNumCounters(values objectivec.IObject, values2 []float64, counters uint64)
	SetRawArrayValuesWithUint32ValuesAndNumCounters(values objectivec.IObject, uint32Values unsafe.Pointer, counters uint64)
	SetRawArrayValuesWithUint64ValuesAndNumCounters(values objectivec.IObject, uint64Values unsafe.Pointer, counters uint64)
	SetValueValue(value objectivec.IObject, value2 objectivec.IObject) objectivec.IObject
	SetValues(values objectivec.IObject)
	VirtualMachine() unsafe.Pointer
}

// Init initializes the instance.
func (g GTJSScriptingContext) Init() GTJSScriptingContext {
	rv := objc.Send[GTJSScriptingContext](g.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g GTJSScriptingContext) Autorelease() GTJSScriptingContext {
	rv := objc.Send[GTJSScriptingContext](g.ID, objc.Sel("autorelease"))
	return rv
}

// NewGTJSScriptingContext creates a new GTJSScriptingContext instance.
func NewGTJSScriptingContext() GTJSScriptingContext {
	class := getGTJSScriptingContextClass()
	rv := objc.Send[GTJSScriptingContext](objc.ID(class.class), objc.Sel("new"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/_cachedStringFromString:
func (g GTJSScriptingContext) _cachedStringFromString(string_ string) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("_cachedStringFromString:"), unsafe.Pointer(unsafe.StringData(string_+"\x00")))
	return rv
}

// CachedStringFromString is an exported wrapper for the private method _cachedStringFromString.
func (g GTJSScriptingContext) CachedStringFromString(string_ string) unsafe.Pointer {
	return g._cachedStringFromString(string_)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/_clearCache
func (g GTJSScriptingContext) _clearCache() {
	objc.Send[objc.ID](g.ID, objc.Sel("_clearCache"))
}

// ClearCache is an exported wrapper for the private method _clearCache.
func (g GTJSScriptingContext) ClearCache() {
	g._clearCache()
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/_jsStringToString:
func (g GTJSScriptingContext) _jsStringToString(string_ unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_jsStringToString:"), string_)
	return objectivec.Object{ID: rv}
}

// JsStringToString is an exported wrapper for the private method _jsStringToString.
func (g GTJSScriptingContext) JsStringToString(string_ unsafe.Pointer) objectivec.IObject {
	return g._jsStringToString(string_)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/_jsValueToString:
func (g GTJSScriptingContext) _jsValueToString(string_ unsafe.Pointer) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("_jsValueToString:"), string_)
	return objectivec.Object{ID: rv}
}

// JsValueToString is an exported wrapper for the private method _jsValueToString.
func (g GTJSScriptingContext) JsValueToString(string_ unsafe.Pointer) objectivec.IObject {
	return g._jsValueToString(string_)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/allocNewContext
func (g GTJSScriptingContext) AllocNewContext() {
	objc.Send[objc.ID](g.ID, objc.Sel("allocNewContext"))
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/callFunction:withArguments:
func (g GTJSScriptingContext) CallFunctionWithArguments(function objectivec.IObject, arguments objectivec.IObject) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("callFunction:withArguments:"), function, arguments)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/callGlobalFunction:
func (g GTJSScriptingContext) CallGlobalFunction(function string) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("callGlobalFunction:"), unsafe.Pointer(unsafe.StringData(function+"\x00")))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/createArrayRef:
func (g GTJSScriptingContext) CreateArrayRef(ref objectivec.IObject) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("createArrayRef:"), ref)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/evaluteScript:scriptURL:
func (g GTJSScriptingContext) EvaluteScriptScriptURL(script objectivec.IObject, url foundation.INSURL) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("evaluteScript:scriptURL:"), script, url)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/getGlobalDouble:
func (g GTJSScriptingContext) GetGlobalDouble(double string) float64 {
	rv := objc.Send[float64](g.ID, objc.Sel("getGlobalDouble:"), unsafe.Pointer(unsafe.StringData(double+"\x00")))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/getGlobalJSONObject:
func (g GTJSScriptingContext) GetGlobalJSONObject(jSONObject string) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("getGlobalJSONObject:"), unsafe.Pointer(unsafe.StringData(jSONObject+"\x00")))
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/getValue:
func (g GTJSScriptingContext) GetValue(value objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("getValue:"), value)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/setExceptionHandler:
func (g GTJSScriptingContext) SetExceptionHandler(handler VoidHandler) {
	_block0, _ := NewVoidBlock(handler)
	objc.Send[objc.ID](g.ID, objc.Sel("setExceptionHandler:"), _block0)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/setGlobalDouble:value:
func (g GTJSScriptingContext) SetGlobalDoubleValue(double string, value float64) {
	objc.Send[objc.ID](g.ID, objc.Sel("setGlobalDouble:value:"), unsafe.Pointer(unsafe.StringData(double+"\x00")), value)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/setGlobalJSONObject:value:
func (g GTJSScriptingContext) SetGlobalJSONObjectValue(jSONObject string, value objectivec.IObject) bool {
	rv := objc.Send[bool](g.ID, objc.Sel("setGlobalJSONObject:value:"), unsafe.Pointer(unsafe.StringData(jSONObject+"\x00")), value)
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/setRawArrayValues:withDoubleValues:andNumCounters:
func (g GTJSScriptingContext) SetRawArrayValuesWithDoubleValuesAndNumCounters(values objectivec.IObject, values2 []float64, counters uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("setRawArrayValues:withDoubleValues:andNumCounters:"), values, values2, counters)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/setRawArrayValues:withUint32Values:andNumCounters:
func (g GTJSScriptingContext) SetRawArrayValuesWithUint32ValuesAndNumCounters(values objectivec.IObject, uint32Values unsafe.Pointer, counters uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("setRawArrayValues:withUint32Values:andNumCounters:"), values, uint32Values, counters)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/setRawArrayValues:withUint64Values:andNumCounters:
func (g GTJSScriptingContext) SetRawArrayValuesWithUint64ValuesAndNumCounters(values objectivec.IObject, uint64Values unsafe.Pointer, counters uint64) {
	objc.Send[objc.ID](g.ID, objc.Sel("setRawArrayValues:withUint64Values:andNumCounters:"), values, uint64Values, counters)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/setValue:value:
func (g GTJSScriptingContext) SetValueValue(value objectivec.IObject, value2 objectivec.IObject) objectivec.IObject {
	rv := objc.Send[objc.ID](g.ID, objc.Sel("setValue:value:"), value, value2)
	return objectivec.Object{ID: rv}
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/setValues:
func (g GTJSScriptingContext) SetValues(values objectivec.IObject) {
	objc.Send[objc.ID](g.ID, objc.Sel("setValues:"), values)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/sharedContext
func (_GTJSScriptingContextClass GTJSScriptingContextClass) SharedContext() GTJSScriptingContext {
	rv := objc.Send[objc.ID](objc.ID(_GTJSScriptingContextClass.class), objc.Sel("sharedContext"))
	return GTJSScriptingContextFromID(rv)
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/context
func (g GTJSScriptingContext) Context() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("context"))
	return rv
}

// See: https://developer.apple.com/documentation/GTShaderProfiler/GTJSScriptingContext/virtualMachine
func (g GTJSScriptingContext) VirtualMachine() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g.ID, objc.Sel("virtualMachine"))
	return rv
}

// SetExceptionHandlerSync is a synchronous wrapper around [GTJSScriptingContext.SetExceptionHandler].
// It blocks until the completion handler fires or the context is cancelled.
func (g GTJSScriptingContext) SetExceptionHandlerSync(ctx context.Context) error {
	done := make(chan struct{}, 1)
	g.SetExceptionHandler(func() {
		done <- struct{}{}
	})
	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}
