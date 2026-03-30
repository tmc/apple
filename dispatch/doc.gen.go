// Code generated from Apple documentation for Dispatch. DO NOT EDIT.

// Package dispatch provides Go bindings for the Dispatch framework.
//
// Execute code concurrently on multicore hardware by submitting work to dispatch queues managed by the system.
//
// Dispatch, also known as Grand Central Dispatch (GCD), contains language features, runtime libraries, and system enhancements that provide systemic, comprehensive improvements to the support for concurrent code execution on multicore hardware in macOS, iOS, watchOS, and tvOS.
//
// # Notes
//
//   - High-level wrappers (Queue, Group, Semaphore) are provided alongside generated interop types.
//
// [Dispatch Documentation]: https://developer.apple.com/documentation/Dispatch
package dispatch

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the Dispatch library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{"/usr/lib/system/libdispatch.dylib"}

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
	fmt.Fprintf(os.Stderr, "warning: Dispatch: failed to load framework from any known path\n")
}
