// Code generated from Apple documentation. DO NOT EDIT.

package javascriptcore

import (
	"unsafe"

	"github.com/ebitengine/purego"
	"github.com/tmc/apple/objc"
)

var (
	// JSPropertyDescriptorConfigurableKey is the Boolean value for this key determines whether the property deleted from its JavaScript object or its descriptor changed.
	//
	// See: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyDescriptorConfigurableKey
	JSPropertyDescriptorConfigurableKey string
	// JSPropertyDescriptorEnumerableKey is the Boolean value for this key determines whether the property appears when enumerating the JavaScript object’s properties.
	//
	// See: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyDescriptorEnumerableKey
	JSPropertyDescriptorEnumerableKey string
	// JSPropertyDescriptorGetKey is the JavaScript function to be invoked when reading the property.
	//
	// See: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyDescriptorGetKey
	JSPropertyDescriptorGetKey string
	// JSPropertyDescriptorSetKey is the JavaScript function to be invoked when writing to the property.
	//
	// See: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyDescriptorSetKey
	JSPropertyDescriptorSetKey string
	// JSPropertyDescriptorValueKey is the value for the property on the JavaScript object.
	//
	// See: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyDescriptorValueKey
	JSPropertyDescriptorValueKey string
	// JSPropertyDescriptorWritableKey is the Boolean value for this key determines whether the property permits assignment with the JavaScript `=` operator.
	//
	// See: https://developer.apple.com/documentation/JavaScriptCore/JSPropertyDescriptorWritableKey
	JSPropertyDescriptorWritableKey string
)

var (
	// KJSClassDefinitionEmpty is a class definition structure of the current version that contains null pointers and has no attributes.
	//
	// See: https://developer.apple.com/documentation/JavaScriptCore/kJSClassDefinitionEmpty
	KJSClassDefinitionEmpty JSClassDefinition
)

func init() {
	if frameworkHandle == 0 {
		return
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "JSPropertyDescriptorConfigurableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				JSPropertyDescriptorConfigurableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "JSPropertyDescriptorEnumerableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				JSPropertyDescriptorEnumerableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "JSPropertyDescriptorGetKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				JSPropertyDescriptorGetKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "JSPropertyDescriptorSetKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				JSPropertyDescriptorSetKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "JSPropertyDescriptorValueKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				JSPropertyDescriptorValueKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "JSPropertyDescriptorWritableKey"); err == nil && ptr != 0 {
		nsStringID := objc.IDValueAt(ptr)
		if nsStringID != 0 {
			cstr := objc.Send[*byte](nsStringID, objc.Sel("UTF8String"))
			if cstr != nil {
				JSPropertyDescriptorWritableKey = objc.GoString(cstr)
			}
		}
	}

	if ptr, err := purego.Dlsym(frameworkHandle, "kJSClassDefinitionEmpty"); err == nil && ptr != 0 {
		KJSClassDefinitionEmpty = *(*JSClassDefinition)(unsafe.Pointer(ptr))
	}

}
