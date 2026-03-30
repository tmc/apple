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

// frameworkPaths lists paths to try when loading the appleneuralengine library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{"/System/Library/PrivateFrameworks/AppleNeuralEngine.framework/AppleNeuralEngine"}

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
	fmt.Fprintf(os.Stderr, "warning: appleneuralengine: failed to load framework from any known path\n")
}
