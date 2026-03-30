// Code generated from Apple documentation for DiskArbitration. DO NOT EDIT.

// Package diskarbitration provides Go bindings for the DiskArbitration framework.
//
// Provides mechanisms to register and block disk mount or unmount events.
//
// For related documentation, see [Mac Technology Overview](<https://developer.apple.com/library/archive/documentation/MacOSX/Conceptual/OSX_Technology_Overview/About/About.html#//apple_ref/doc/uid/TP40001067>).
//
// # Variables
//
//   - kDADiskDescriptionFSKitPrefix
//   - kDADiskDescriptionRepairRunningKey
//   - kDADiskMountOptionNoFollow
//
// [DiskArbitration Documentation]: https://developer.apple.com/documentation/DiskArbitration
package diskarbitration

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the DiskArbitration library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
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
	fmt.Fprintf(os.Stderr, "warning: DiskArbitration: failed to load framework from any known path\n")
}
