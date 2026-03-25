
// Code generated from Apple documentation for appleneuralengine. DO NOT EDIT.

// Package appleneuralengine provides Go bindings for the appleneuralengine framework.
//
// # Key Types
//
//   - [ANEStrings]
//   - [ANEVirtualClient]
//   - [ANEModel]
//   - [ANEClient]
//   - [ANEInMemoryModel]
//   - [ANEErrors]
//   - [ANERequest]
//   - [ANEDaemonConnection]
//   - [ANEDeviceInfo]
//   - [ANEProgramForEvaluation]
package appleneuralengine

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/PrivateFrameworks/AppleNeuralEngine.framework/AppleNeuralEngine"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: appleneuralengine: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

