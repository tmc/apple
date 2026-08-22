// Code generated from Apple documentation for ScriptingBridge. DO NOT EDIT.

// Package scriptingbridge provides Go bindings for the ScriptingBridge framework.
//
// Automate scriptable apps by sending and receiving Apple events.
//
// Scripting Bridge is a technology that lets you control scriptable Apple and
// third-party applications using standard Objective-C syntax. Introduced in
// OS X version 10.5 (Leopard), the Scripting Bridge framework dynamically
// implements an Objective-C bridge to OSA-compliant applications—that is,
// applications having a scripting interface (usually defined in a `sdef`
// file). As part of this implementation, it generates Objective-C class
// implementations of the classes it finds in the scripting interface,
// including objects and methods representing properties, elements, commands,
// and so on. The objects are derived from classes defined in the Scripting
// Bridge framework.
//
// # Classes
//
//   - [SBApplication]: The class provides a mechanism enabling an Objective-C program to send Apple events to a scriptable application and receive Apple events in response.
//   - [SBElementArray]: is subclass of that manages collections of related objects.
//   - [SBObject]: The class declares methods that can be invoked on any object in a scriptable application.
//
// # Protocols
//
//   - [SBApplicationDelegate]: This informal protocol defines a delegation method for handling Apple event errors that are sent from a target application to an object.
//
// # Key Types
//
//   - [SBApplication] - The [SBApplication] class provides a mechanism enabling an Objective-C program to send Apple events to a scriptable application and receive Apple events in response.
//   - [SBObject] - The [SBObject] class declares methods that can be invoked on any object in a scriptable application.
//   - [SBElementArray] - [SBElementArray] is subclass of [NSMutableArray] that manages collections of related SBObject objects.
package scriptingbridge

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the ScriptingBridge library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/ScriptingBridge.framework/ScriptingBridge",
	"/usr/lib/libScriptingBridge.dylib",
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
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: ScriptingBridge: failed to load framework from any known path\n")
	}
}
