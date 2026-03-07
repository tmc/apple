
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

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/usr/lib/system/libdispatch.dylib"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: Dispatch: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

