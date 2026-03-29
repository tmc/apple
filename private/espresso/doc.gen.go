
// Code generated from Apple documentation for espresso. DO NOT EDIT.

// Package espresso provides Go bindings for the espresso framework.
//
// # Key Types
//
//   - [EspressoFDOverfeatNetwork]
//   - [EspressoImage2Image]
//   - [ETTaskDefinition]
//   - [ETImageDescriptorExtractor]
//   - [ETDataTensor]
//   - [ETTask]
//   - [ETModelDef]
//   - [ETVariable]
//   - [EspressoANEIOSurface]
//   - [ETDataSourceFromFolderData]
package espresso

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/PrivateFrameworks/Espresso.framework/Espresso"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: espresso: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

