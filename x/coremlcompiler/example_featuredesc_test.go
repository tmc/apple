package coremlcompiler_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/tmc/apple/x/coremlcompiler"
)

// This example builds a compiled bundle from artifacts a MIL generator
// produced: the program text, the weight blobs it references, and the type of
// each feature named the way MIL names it.
//
// A generator such as github.com/tmc/modelir already carries all three, so the
// whole adapter is the two loops below; nothing in this package needs to know
// the generator's types.
func ExampleTensorFeatureDescription() {
	// What the generator emitted.
	milText := `program(1.3) {
    func main<ios18>(tensor<fp16, [1, 1, 8]> x) {
        tensor<fp16, [8, 8]> w = const()[name=string("w"), val=tensor<fp16, [8, 8]>(BLOBFILE(path=string("@model_path/weights/w.bin"), offset=uint64(64)))];
        tensor<fp16, [1, 1, 8]> y = linear(bias=None, weight=w, x=x)[name=string("y")];
    } -> (y);
}`
	weights := []coremlcompiler.WeightFile{
		{Path: "@model_path/weights/w.bin", Blob: make([]byte, 192)},
	}
	type tensor struct {
		name  string
		dtype string // as MIL spells it
		shape []int64
	}
	inputs := []tensor{{"x", "fp16", []int64{1, 1, 8}}}
	outputs := []tensor{{"y", "fp16", []int64{1, 1, 8}}}

	// The adapter.
	var desc coremlcompiler.ModelDescription
	for _, t := range inputs {
		fd, err := coremlcompiler.TensorFeatureDescription(t.name, t.dtype, t.shape)
		if err != nil {
			log.Fatal(err)
		}
		desc.Inputs = append(desc.Inputs, fd)
	}
	for _, t := range outputs {
		fd, err := coremlcompiler.TensorFeatureDescription(t.name, t.dtype, t.shape)
		if err != nil {
			log.Fatal(err)
		}
		desc.Outputs = append(desc.Outputs, fd)
	}

	tmpDir, err := os.MkdirTemp("", "coremlcompiler-example-*")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tmpDir)

	weightRoot := filepath.Join(tmpDir, "weights")
	if err := coremlcompiler.WriteWeightRoot(weightRoot, weights); err != nil {
		log.Fatal(err)
	}
	out := filepath.Join(tmpDir, "model.mlmodelc")
	if err := coremlcompiler.CompileMILText(milText, 9, desc, weightRoot, out); err != nil {
		log.Fatal(err)
	}

	_, err = os.Stat(filepath.Join(out, "weights", "w.bin"))
	fmt.Println("weight in bundle:", err == nil)

	// An element type Core ML multi-arrays cannot express is an error, not a
	// silently substituted fp32.
	_, err = coremlcompiler.TensorFeatureDescription("x", "bf16", []int64{1, 8})
	fmt.Println("bf16 input:", err)

	// Output:
	// weight in bundle: true
	// bf16 input: coremlcompiler: feature "x": MIL type bf16 has no Core ML multi-array element type
}
