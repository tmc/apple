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

// frameworkPaths lists paths to try when loading the espresso library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{"/System/Library/PrivateFrameworks/Espresso.framework/Espresso"}

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
	fmt.Fprintf(os.Stderr, "warning: espresso: failed to load framework from any known path\n")
}
