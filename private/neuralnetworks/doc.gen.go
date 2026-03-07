
// Code generated from Apple documentation for neuralnetworks. DO NOT EDIT.

// Package neuralnetworks provides Go bindings for the neuralnetworks framework.
//
// # Key Types
//
//   - [MPSNNBinaryArithmeticNode]
//   - [MPSNNGraph]
//   - [MPSNNArithmeticGradientNode]
//   - [MPSNNForwardLoss]
//   - [MPSNNLossGradient]
//   - [MPSNNForwardLossNode]
//   - [MPSNNLossGradientNode]
//   - [MPSNNNeuronDescriptor]
//   - [MPSNNOptimizerAdam]
//   - [MPSNNOptimizerRMSProp]
package neuralnetworks

import (
	"fmt"
	"os"
	"github.com/ebitengine/purego"
)

// frameworkPath is the system path to the framework binary.
const frameworkPath = "/System/Library/PrivateFrameworks/NeuralNetworks.framework/NeuralNetworks"
// frameworkHandle is the handle to the loaded framework.
var frameworkHandle uintptr

func init() {
	var err error
	frameworkHandle, err = purego.Dlopen(frameworkPath, purego.RTLD_LAZY|purego.RTLD_GLOBAL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "warning: neuralnetworks: failed to load %s: %v\n", frameworkPath, err)
		return 
	}
}

