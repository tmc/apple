// Code generated from Apple documentation for DiskArbitration. DO NOT EDIT.

// Package diskarbitration provides Go bindings for the DiskArbitration framework.
//
// Provides mechanisms to register and block disk mount or unmount events.
//
// For related documentation, see Mac Technology Overview.
//
// # Variables
//
//   - [KDADiskDescriptionFSKitPrefix]
//   - [KDADiskDescriptionRepairRunningKey]
package diskarbitration

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the DiskArbitration library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/DiskArbitration.framework/DiskArbitration",
	"/usr/lib/libDiskArbitration.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: DiskArbitration: failed to load framework from any known path\n")
	}
}
