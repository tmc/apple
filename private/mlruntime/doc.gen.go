
// Code generated from Apple documentation for mlruntime. DO NOT EDIT.

// Package mlruntime provides Go bindings for the mlruntime framework.
//
// # Key Types
//
//   - [MLRTaskParameters]
//   - [MLROnDemandConnectionHandler]
//   - [MLRTrialDediscoTaskResult]
//   - [MLRTaskAttachments]
//   - [MLRTask]
//   - [MLRExtensionRemoteContext]
//   - [MLRServiceClient]
//   - [MLRTaskResult]
//   - [MLRTrialDediscoRecipe]
//   - [MLRDonationManager]
package mlruntime

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/PrivateFrameworks/MLRuntime.framework/MLRuntime"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: mlruntime: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

