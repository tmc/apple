// Code generated from Apple documentation for XPC. DO NOT EDIT.

// Package xpc provides Go bindings for the XPC framework.
//
// High-level Go bindings for the XPC framework.
//
// The public API is centered on Listener, Session, and ReceivedMessage.
//
// # Notes
//
//   - Dictionary-first messaging is always available: CallDictionary and NotifyDictionary send a Dictionary, and ReceivedMessage.Dictionary returns one.
//   - Typed payloads are supported through Call, Notify, Decode, Marshaler, and Unmarshaler.
//   - Swift members without a safe public C path are omitted and recorded in xpc.omissions.gen.go.
package xpc

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the XPC library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/usr/lib/system/libxpc.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: XPC: failed to load framework from any known path\n")
	}
}
