
// Code generated from Apple documentation for CoreML. DO NOT EDIT.

// Package coreml provides Go bindings for the CoreML framework.
//
// # Key Types
//
//   - [MLNeuralNetworkEngine]
//   - [MLKNearestNeighborsClassifier]
//   - [MLModel]
//   - [MLNeuralNetworkUpdateEngine]
//   - [MLE5Engine]
//   - [MLModelConfiguration]
//   - [MLE5ExecutionStreamOperation]
//   - [MLNeuralNetworkContainer]
//   - [MLMultiArray]
//   - [MLNeuralNetworkMLComputeUpdateEngine]
package coreml

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/Frameworks/CoreML.framework/CoreML"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: CoreML: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

