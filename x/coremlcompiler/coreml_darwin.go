//go:build darwin

package coremlcompiler

import (
	"fmt"
	"unsafe"

	"github.com/tmc/apple/coreml"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/objc"
)

// CompileMLModelAtURL uses [MLModel compileModelAtURL:error:] to compile a
// .mlpackage or .mlmodel to a .mlmodelc directory. Returns the path to the
// compiled model in a temporary location chosen by CoreML.
func CompileMLModelAtURL(sourcePath string) (string, error) {
	sourceURL := foundation.NewURLFileURLWithPath(sourcePath)
	resultURL, err := coreml.GetMLModelClass().CompileModelAtURLError(sourceURL)
	if err != nil {
		return "", fmt.Errorf("coreml compile: %w", err)
	}
	return resultURL.Path(), nil
}

// CoreMLModel wraps a loaded MLModel for inference.
type CoreMLModel struct {
	model coreml.MLModel
}

// LoadCoreMLModel loads a compiled .mlmodelc from disk using
// [MLModel modelWithContentsOfURL:error:].
func LoadCoreMLModel(modelcPath string) (*CoreMLModel, error) {
	url := foundation.NewURLFileURLWithPath(modelcPath)
	model, err := coreml.NewModelWithContentsOfURLError(url)
	if err != nil {
		return nil, fmt.Errorf("coreml: load model: %w", err)
	}
	return &CoreMLModel{model: model}, nil
}

// Predict runs inference with the given named inputs and returns the
// named output feature's MLMultiArray data pointer and shape.
func (m *CoreMLModel) Predict(inputs []PredictInput, outputName string) (*PredictOutput, error) {
	dict := foundation.NewNSMutableDictionary()

	for _, inp := range inputs {
		shape := intsToNSNumbers(inp.Shape)
		strides := intsToNSNumbers(inp.Strides)

		multiArr, err := coreml.GetMLMultiArrayClass().Alloc().InitWithDataPointerShapeDataTypeStridesDeallocatorError(
			inp.Data, shape, inp.DType, strides, nil,
		)
		if err != nil {
			return nil, fmt.Errorf("coreml: create MLMultiArray for %s: %w", inp.Name, err)
		}

		fv := coreml.NewFeatureValueWithMultiArray(multiArr)
		key := foundation.NewStringWithString(inp.Name)
		dict.SetObjectForKey(fv, key)
	}

	provider, err := coreml.NewDictionaryFeatureProviderWithDictionaryError(dict)
	if err != nil {
		return nil, fmt.Errorf("coreml: create feature provider: %w", err)
	}

	result, err := m.model.PredictionFromFeaturesError(provider)
	if err != nil {
		return nil, fmt.Errorf("coreml: prediction failed: %w", err)
	}

	outFV := result.FeatureValueForName(outputName)
	if outFV == nil {
		return nil, fmt.Errorf("coreml: output %q not found", outputName)
	}

	// Convert IMLFeatureValue → concrete MLFeatureValue to access MultiArrayValue.
	fv := coreml.MLFeatureValueFromID(outFV.GetID())
	outMultiArr := coreml.MLMultiArrayFromID(fv.MultiArrayValue().GetID())

	// TODO: use outMultiArr.DataPointer() once the generator emits the accessor.
	dataPtr := objc.Send[unsafe.Pointer](outMultiArr.ID, objc.Sel("dataPointer"))

	outShape := nsNumbersToInts(outMultiArr.Shape())

	return &PredictOutput{
		Data:  dataPtr,
		Shape: outShape,
		DType: outMultiArr.DataType(),
	}, nil
}

// Close releases the CoreML model reference.
func (m *CoreMLModel) Close() {
	m.model = coreml.MLModel{}
}

func intsToNSNumbers(vals []int) []foundation.NSNumber {
	nums := make([]foundation.NSNumber, len(vals))
	for i, v := range vals {
		nums[i] = foundation.NewNumberWithInteger(v)
	}
	return nums
}

func nsNumbersToInts(nums []foundation.NSNumber) []int {
	vals := make([]int, len(nums))
	for i, n := range nums {
		vals[i] = n.IntegerValue()
	}
	return vals
}
