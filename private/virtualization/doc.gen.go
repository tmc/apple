
// Code generated from Apple documentation for Virtualization. DO NOT EDIT.

// Package virtualization provides Go bindings for the Virtualization framework.
//
// # Key Types
//
//   - [VZVirtualMachineView]
//   - [VZVirtualMachineConfiguration]
//   - [VZVirtualMachine]
//   - [VZMacPlatformConfiguration]
//   - [VZMacGraphicsDeviceConfiguration]
//   - [VZUSBMassStorageDevice]
//   - [VZMacMachineIdentifier]
//   - [VZGraphicsDisplay]
//   - [VZMacAuxiliaryStorage]
//   - [VZMacGraphicsDisplay]
package virtualization

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/Virtualization.framework/Virtualization"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: Virtualization: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

