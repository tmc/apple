
// Code generated from Apple documentation for diskimages2. DO NOT EDIT.

// Package diskimages2 provides Go bindings for the diskimages2 framework.
//
// # Key Types
//
//   - [DIConvertParams]
//   - [DiskImageGraph]
//   - [DIAttachParams]
//   - [BaseDiskImageCreator]
//   - [DIEncryptionFrontend]
//   - [DiskImageParamsXPC]
//   - [DICreateParams]
//   - [DIBaseParams]
//   - [DiskImageGraphNode]
//   - [DIError]
package diskimages2

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/PrivateFrameworks/DiskImages2.framework/DiskImages2"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: diskimages2: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

