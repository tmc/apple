package coremlcompiler_test

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tmc/apple/x/coremlcompiler"
)

func Example() {
	// Compile a CoreML model package to a compiled bundle.
	err := coremlcompiler.Compile("model.mlpackage", "model.mlmodelc")
	if err != nil {
		fmt.Println(err)
	}

	// The compiled bundle can be loaded by CoreML or x/ane:
	//
	//   rt, _ := ane.Open()
	//   k, _ := rt.Compile(ane.CompileOptions{
	//       ModelType:   ane.ModelTypePackage,
	//       PackagePath: "model.mlmodelc",
	//   })
}

func ExampleCompileMILText() {
	tmpDir, err := os.MkdirTemp("", "coremlcompiler-example-*")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer os.RemoveAll(tmpDir)

	milText := `program(1.3) {
    func main<ios18>(tensor<fp16, [1, 8]> x) {
    } -> (x);
}`
	desc := coremlcompiler.ModelDescription{
		Inputs:  []coremlcompiler.FeatureDescription{{Name: "x"}},
		Outputs: []coremlcompiler.FeatureDescription{{Name: "x"}},
	}
	err = coremlcompiler.CompileMILText(milText, 8, desc, "", filepath.Join(tmpDir, "model.mlmodelc"))
	if err != nil {
		fmt.Println(err)
	}
}
