package coremlcompiler

import (
	"os"
	"path/filepath"
	"testing"
)

// TestCompileMILTextOutputInsideWeightRoot writes the bundle into the same
// directory that holds the weights. The weight copy walks that directory,
// so without a guard it descends into the output it is creating and
// recurses until the path exceeds the filesystem limit.
func TestCompileMILTextOutputInsideWeightRoot(t *testing.T) {
	dir := t.TempDir()
	if err := os.WriteFile(filepath.Join(dir, "w.bin"), []byte("weight"), 0o644); err != nil {
		t.Fatal(err)
	}
	out := filepath.Join(dir, "model.mlmodelc")
	milText := "program(1.3) {\n    func main<ios18>(tensor<fp16, [1, 8]> x) {\n    } -> (x);\n}"
	desc := ModelDescription{
		Inputs:  []FeatureDescription{{Name: "x"}},
		Outputs: []FeatureDescription{{Name: "x"}},
	}
	if err := CompileMILText(milText, 9, desc, dir, out); err != nil {
		t.Fatalf("CompileMILText: %v", err)
	}
	if _, err := os.Stat(filepath.Join(out, "w.bin")); err != nil {
		t.Errorf("weight not copied: %v", err)
	}
	if _, err := os.Stat(filepath.Join(out, "model.mlmodelc")); err == nil {
		t.Error("output directory copied into itself")
	}
}
