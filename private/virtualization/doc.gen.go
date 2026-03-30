// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

// Package virtualization provides Go bindings for the Virtualization framework.
//
// # Key Types
//
//   - [VZVirtualMachineView]
//   - [VZVirtualMachineConfiguration]
//   - [VZVirtualMachine]
//   - [VZCustomVirtioDeviceConfiguration]
//   - [VZVNCServer]
//   - [VZFramebufferView]
//   - [VZMacPlatformConfiguration]
//   - [VZCustomMMIODeviceConfiguration]
//   - [VZIOUSBHostPassthroughDevice]
//   - [VZMacGraphicsDeviceConfiguration]
package virtualization

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the Virtualization library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{"/System/Library/Frameworks/Virtualization.framework/Virtualization"}

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
	fmt.Fprintf(os.Stderr, "warning: Virtualization: failed to load framework from any known path\n")
}
