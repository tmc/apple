// Code generated from Apple documentation for CoreML. DO NOT EDIT.

// Package coreml provides Go bindings for the CoreML framework.
//
// # Key Types
//
//   - [MLNeuralNetworkEngine]
//   - [MLKNearestNeighborsClassifier]
//   - [MLNeuralNetworkUpdateEngine]
//   - [MLE5Engine]
//   - [MLModelConfiguration]
//   - [MLE5ExecutionStreamOperation]
//   - [MLModel]
//   - [MLNeuralNetworkContainer]
//   - [MLNeuralNetworkMLComputeUpdateEngine]
//   - [MLLoaderEvent]
package coreml

import (
	"fmt"
	"os"

	"github.com/ebitengine/purego"
)

// frameworkPaths lists paths to try when loading the CoreML library.
// The framework bundle path is tried first; a /usr/lib dylib fallback covers
// C-API frameworks that are not in the dyld shared cache as bundles.
var frameworkPaths = []string{"/System/Library/Frameworks/CoreML.framework/CoreML"}

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
	fmt.Fprintf(os.Stderr, "warning: CoreML: failed to load framework from any known path\n")
}
