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
//   - Logging: Capture telemetry from your app for debugging and performance analysis using the unified logging system. ([Os_signpost_id_t])
//
// # Task Management
//
//   - Workgroups: Schedule one or more threads to run at regular intervals and before specific deadlines.
//   - Synchronization: Access low-level synchronization mechanisms to control state across threads.
//
// # Memory
//
//   - os_proc_available_memory: Determines the amount of memory available to the current app.
//   - [Os_block_t]: A block that takes no arguments and returns no value.
//   - [Os_function_t]: A pointer to a function.
//   - [Os_release]
//   - [Os_retain]
//
// # Functions
//
//   - os_lockdown_mode_enabled
//   - [Os_security_config_get]
//   - [Os_security_config_get_for_proc]
//   - [Os_security_config_get_for_task]
//
// # Macros
//
//   - API_OBSOLETED
//   - API_OBSOLETED_BEGIN
//   - API_OBSOLETED_END
//   - API_OBSOLETED_WITH_REPLACEMENT
//   - API_OBSOLETED_WITH_REPLACEMENT_BEGIN
//   - API_OBSOLETED_WITH_REPLACEMENT_END
//   - LOG_SWIFT_NAME
//   - LOG_SWIFT_NEWTYPE
//   - OS_LOG_STRING_SECTION
//   - OS_WORKGROUP_ENUM_API_DEPRECATED_WITH_REPLACEMENT
//
// # Enumerations
//
//   - [Os_security_config_t]//
//
// # Key Types
//
//   - [OS_object]
//   - [OS_os_workgroup]
//   - [OS_os_workgroup_interval]
//   - [OS_os_workgroup_parallel]
//   - [OSLog] - A container of related log messages.
//
// [Reading UNIX Manual Pages]: https://developer.apple.com/documentation/os/reading-unix-manual-pages
package os

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the os library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{"/usr/lib/system/libsystem_trace.dylib"}

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
