//go:build !darwin

package coremlcompiler

import "fmt"

// CompileMLModelAtURL uses [MLModel compileModelAtURL:error:] to compile a
// .mlpackage or .mlmodel to a .mlmodelc directory. Only available on darwin.
func CompileMLModelAtURL(sourcePath string) (string, error) {
	return "", fmt.Errorf("coreml compile: only supported on darwin")
}

// CoreMLModel wraps a loaded MLModel for inference. Only functional on darwin.
type CoreMLModel struct{}

// LoadCoreMLModel loads a compiled .mlmodelc from disk. Only available on darwin.
func LoadCoreMLModel(modelcPath string) (*CoreMLModel, error) {
	return nil, fmt.Errorf("coreml: only supported on darwin")
}

// Predict runs inference. Only available on darwin.
func (m *CoreMLModel) Predict(inputs []PredictInput, outputName string) (*PredictOutput, error) {
	return nil, fmt.Errorf("coreml: only supported on darwin")
}

// Close releases the CoreML model reference.
func (m *CoreMLModel) Close() {}
