// Code generated from Apple documentation for JavaScriptCore. DO NOT EDIT.

// Package javascriptcore provides Go bindings for the JavaScriptCore framework.
//
// Evaluate JavaScript programs from within an app, and support JavaScript
// scripting of your app.
//
// The JavaScriptCore framework provides the ability to evaluate JavaScript
// programs from within Swift, Objective-C, and C-based apps. You can use also
// use JavaScriptCore to insert custom objects into the JavaScript
// environment.
//
// # Execution Environment
//
//   - [JSVirtualMachine]: A self-contained environment for JavaScript execution.
//   - [JSContext]: A JavaScript execution environment.
//
// # JavaScript Code
//
//   - [JSValue]: A JavaScript value. ([JSValueProperty])
//   - [JSManagedValue]: A JavaScript value with conditional retain behavior to provide automatic memory management.
//
// # Native Code
//
//   - [JSExport]: The protocol for exporting Objective-C objects to JavaScript.
//
// # C API
//
//   - [C JavaScriptCore API]: Browse the alternative C-based APIs for JavaScriptCore. ([JSContextGroupRef], [JSContextRef], [JSGlobalContextRef], [JSStringRef], [JSClassRef])
//
// # Functions
//
//   - [JSBigIntCreateWithDouble]
//   - [JSBigIntCreateWithInt64]
//   - [JSBigIntCreateWithString]
//   - [JSBigIntCreateWithUInt64]
//   - [JSValueCompare]
//   - [JSValueCompareDouble]
//   - [JSValueCompareInt64]
//   - [JSValueCompareUInt64]
//   - [JSValueIsBigInt]
//   - [JSValueToInt32]
//   - [JSValueToInt64]
//   - [JSValueToUInt32]
//   - [JSValueToUInt64]
//
// # Macros
//
//   - JSC_ASSUME_NONNULL_BEGIN
//   - JSC_ASSUME_NONNULL_END
//   - JSC_CF_ENUM
//   - JSC_NONNULL
//   - JSC_NULLABLE
//   - JSC_NULL_UNSPECIFIED
//
// # Enumerations
//
//   - [JSRelationCondition]//
//
// # Key Types
//
//   - [JSValue] - A JavaScript value.
//   - [JSContext] - A JavaScript execution environment.
//   - [JSManagedValue] - A JavaScript value with conditional retain behavior to provide automatic memory management.
//   - [JSVirtualMachine] - A self-contained environment for JavaScript execution.
//
// [C JavaScriptCore API]: https://developer.apple.com/documentation/javascriptcore/c-javascriptcore-api
package javascriptcore

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the JavaScriptCore library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/JavaScriptCore.framework/JavaScriptCore",
	"/usr/lib/libJavaScriptCore.dylib",
}

// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	for _, path := range frameworkPaths {
		h, err := purego.Dlopen(path, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
		if err == nil {
			frameworkHandle = h
			return
		}
	}
	fmt.Fprintf(os.Stderr, "warning: JavaScriptCore: failed to load framework from any known path\n")
}
