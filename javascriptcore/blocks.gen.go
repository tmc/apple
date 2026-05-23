// Code generated from Apple documentation. DO NOT EDIT.

package javascriptcore

import (
	"github.com/tmc/apple/objc"
)

// JSContextJSValueHandler is the signature for a completion handler block.
type JSContextJSValueHandler = func(*JSContext, *JSValue)

// NewJSContextJSValueBlock wraps a Go [JSContextJSValueHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
func NewJSContextJSValueBlock(handler JSContextJSValueHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID) {
		var result *JSContext
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := JSContextFromID(resultID)
			result = &v
		}
		var extra0 *JSValue
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := JSValueFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}

// JSValueJSValueHandler handles A callback block to invoke during initialization of the promise object.
//
// Used by:
//   - [JSValue.ValueWithNewPromiseInContextFromExecutor]
type JSValueJSValueHandler = func(*JSValue, *JSValue)

// NewJSValueJSValueBlock wraps a Go [JSValueJSValueHandler] as an Objective-C block.
// The caller must defer the returned cleanup function.
//
// Used by:
//   - [JSValue.ValueWithNewPromiseInContextFromExecutor]
func NewJSValueJSValueBlock(handler JSValueJSValueHandler) (objc.ID, func()) {
	if handler == nil {
		return 0, func() {}
	}
	block := objc.NewBlock(func(b objc.Block, resultID objc.ID, extra0ID objc.ID) {
		var result *JSValue
		if resultID != 0 {
			objc.Send[objc.ID](resultID, objc.Sel("retain"))
			v := JSValueFromID(resultID)
			result = &v
		}
		var extra0 *JSValue
		if extra0ID != 0 {
			objc.Send[objc.ID](extra0ID, objc.Sel("retain"))
			v := JSValueFromID(extra0ID)
			extra0 = &v
		}
		handler(result, extra0)
	})
	return objc.ID(block), func() { block.Release() }
}
