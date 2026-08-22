// Code generated from Apple documentation for os. DO NOT EDIT.

// Package os provides Go bindings for the os framework.
//
// Coordinate the scheduling and synchronization of your app’s tasks, and
// log information to the console to diagnose issues.
//
// # Essentials
//
//   - [Reading UNIX Manual Pages]: Use the Terminal app to read the documentation for low-level UNIX tools and APIs.
//
// # Logs
//
//   - Logging: Capture telemetry from your app for debugging and performance analysis using the unified logging system. ([OSLogType], [OSSignpostType])
//
// # Memory
//
//   - [OSBlock]: A block that takes no arguments and returns no value.
//   - [OSFunction]: A pointer to a function.
//   - [OSRelease]
//   - [OSRetain]
//
// # Functions
//
//   - [OSSecurityConfigGet]
//   - [OSSecurityConfigGetForProc]
//   - [OSSecurityConfigGetForTask]
//
// # Enumerations
//
//   - [OSSecurityConfig]//
//
// # Key Types
//
//   - [OSObject]
//   - [OSOSWorkgroup]
//   - [OSOSWorkgroupInterval]
//   - [OSOSWorkgroupParallel]
//   - [OSLog] - A container of related log messages.
//
// [Reading UNIX Manual Pages]: https://developer.apple.com/documentation/os/reading-unix-manual-pages
package os

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the os library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/usr/lib/system/libsystem_trace.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: os: failed to load framework from any known path\n")
	}
}
