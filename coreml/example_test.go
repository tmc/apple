//go:build darwin

package coreml_test

import (
	"fmt"

	"github.com/tmc/apple/coreml"
)

func ExampleMLFeatureValue() {
	strVal := coreml.NewFeatureValueWithString("CoreML Example")
	intVal := coreml.NewFeatureValueWithInt64(42)
	dblVal := coreml.NewFeatureValueWithDouble(3.14159)

	fmt.Println("String:", strVal.StringValue())
	fmt.Println("Int64:", intVal.Int64Value())
	fmt.Printf("Double: %.5f\n", dblVal.DoubleValue())
	// Output:
	// String: CoreML Example
	// Int64: 42
	// Double: 3.14159
}

func ExampleMLModelConfiguration() {
	config := coreml.NewMLModelConfiguration()
	config.SetComputeUnits(coreml.MLComputeUnitsAll)

	fmt.Println("Compute units:", config.ComputeUnits())
	// Output:
	// Compute units: MLComputeUnitsAll
}

func ExampleMLPredictionOptions() {
	options := coreml.NewMLPredictionOptions()

	fmt.Println("Output backings count:", options.OutputBackings().Count())
	// Output:
	// Output backings count: 0
}
