// Code generated from Apple documentation for FinderSync. DO NOT EDIT.

// Package findersync provides Go bindings for the FinderSync framework.
//
// Modify the Finder’s user interface to express file synchronization and
// control.
//
// Use Finder Sync to cleanly and safely modify the Finder’s user interface
// to express file synchronization status and control. Unlike most extension
// points, Finder Sync doesn’t add features to a host app. Instead, it lets
// you modify the behavior of the Finder itself.
//
// # Classes
//
//   - [FIFinderSync]: A type to subclass to add badges, custom shortcut menus, and toolbar buttons to the Finder.
//   - [FIFinderSyncController]: A controller that acts as a bridge between your Finder Sync extension and the Finder itself.
//
// # Protocols
//
//   - [FIFinderSync]: The group of methods to implement for modifying the Finder user interface to express file synchronization status and control. ([FIMenuKind])
//
// # Key Types
//
//   - [FIFinderSync] - A type to subclass to add badges, custom shortcut menus, and toolbar buttons to the Finder.
//   - [FIFinderSyncController] - A controller that acts as a bridge between your Finder Sync extension and the Finder itself.
package findersync

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the FinderSync library,
// in order. Frameworks whose symbols live in a known dylib resolve to that
// dylib alone; the rest try the framework bundle first and then a /usr/lib
// dylib fallback, which covers C-API frameworks that are not in the dyld
// shared cache as bundles.
var frameworkPaths = []string{
	"/System/Library/Frameworks/FinderSync.framework/FinderSync",
	"/usr/lib/libFinderSync.dylib",
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
		fmt.Fprintf(os.Stderr, "warning: FinderSync: failed to load framework from any known path\n")
	}
}
