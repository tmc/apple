// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

// Package objectivec provides Go bindings for the ObjectiveC framework.
//
// Gain low-level access to the Objective-C runtime and the Objective-C root
// types.
//
// The Objective-C Runtime module APIs define the base of the Objective-C
// language. These APIs include:
//
// # Classes
//
//   - [NSObject]: The root class of most Objective-C class hierarchies, from which subclasses inherit a basic interface to the runtime system and the ability to behave as Objective-C objects.
//
// # Protocols
//
//   - [NSObject]: The group of methods that are fundamental to all Objective-C objects.
//
// # Key Types
//
//   - [Protocol]
package objectivec

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the ObjectiveC library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{"/usr/lib/libobjc.dylib"}

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
		fmt.Fprintf(os.Stderr, "warning: ObjectiveC: failed to load framework from any known path\n")
	}
}
