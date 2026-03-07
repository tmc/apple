
// Code generated from Apple documentation for remotecoreml. DO NOT EDIT.

// Package remotecoreml provides Go bindings for the remotecoreml framework.
//
// # Key Types
//
//   - [CoreMLModelSecurityServiceToClient]
//   - [CoreMLVersion]
package remotecoreml

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/PrivateFrameworks/RemoteCoreML.framework/RemoteCoreML"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: remotecoreml: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

