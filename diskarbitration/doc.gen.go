
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

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/DiskArbitration.framework/DiskArbitration"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: DiskArbitration: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

