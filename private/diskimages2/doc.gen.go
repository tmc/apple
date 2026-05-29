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

// frameworkPaths lists paths to try when loading the diskimages2 library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{"/System/Library/PrivateFrameworks/DiskImages2.framework/DiskImages2"}

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
	// Loading is best-effort: the warning is silent by default because a missing
	// framework is harmless unless one of its symbols is actually called. Set
	// APPLE_FRAMEWORK_LOAD_DEBUG to surface load failures while diagnosing.
	if os.Getenv("APPLE_FRAMEWORK_LOAD_DEBUG") != "" {
		fmt.Fprintf(os.Stderr, "warning: diskimages2: failed to load framework from any known path\n")
	}
}
