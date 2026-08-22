//go:build darwin

package coremlcompiler_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/tmc/apple/x/coremlcompiler"
)

func Example() {
	// Compile a non-existent package to test error handling.
	err := coremlcompiler.Compile("nonexistent.mlpackage", "model.mlmodelc")
	fmt.Println("compilation failed as expected:", err != nil)

	// Output:
	// compilation failed as expected: true
}

func ExampleCompileMILText() {
	tmpDir, err := os.MkdirTemp("", "coremlcompiler-example-*")
	if err != nil {
		log.Fatal(err)
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
	outputDir := filepath.Join(tmpDir, "model.mlmodelc")
	err = coremlcompiler.CompileMILText(milText, 9, desc, "", outputDir)
	if err != nil {
		log.Fatal(err)
	}

	_, errMil := os.Stat(filepath.Join(outputDir, "model.mil"))
	fmt.Println("model.mil compiled:", errMil == nil)
	_, errData := os.Stat(filepath.Join(outputDir, "coremldata.bin"))
	fmt.Println("coremldata.bin compiled:", errData == nil)

	// Output:
	// model.mil compiled: true
	// coremldata.bin compiled: true
}
