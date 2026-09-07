package coremlcompiler

import (
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"

	"github.com/tmc/apple/x/coremlcompiler/internal/opschema"
)

// testdata/modelir_transformer.mil is MIL text emitted by
// github.com/tmc/modelir; see testdata/gen_modelir_transformer.go. These are
// the features it declares, transcribed from its func signature and return.
// The types are spelled as MIL spells them, which is exactly the form a
// generator hands to [TensorFeatureDescription] and [StateFeatureDescription].
var (
	modelIRInputs = []struct {
		name  string
		dtype string
		shape []int64
	}{
		{"x_in", "fp32", []int64{1, 1, 8}},
	}
	modelIROutputs = []struct {
		name  string
		dtype string
		shape []int64
	}{
		{"y", "fp32", []int64{1, 1, 12}},
	}
	modelIRStates = []struct {
		name  string
		dtype string
		shape []int64
	}{
		{"l0_k_cache_state", "fp16", []int64{1, 2, 4, 4}},
		{"l0_v_cache_state", "fp16", []int64{1, 2, 4, 4}},
	}
)

var blobFilePath = regexp.MustCompile(`@model_path/[^"]+`)

// weightFilesFromMILText returns one WeightFile per distinct BLOBFILE path the
// MIL text references. The blobs are placeholders: CompileMILText copies the
// weight root verbatim, so what is being exercised is that every path the
// program reads resolves inside the bundle.
func weightFilesFromMILText(t *testing.T, milText string) []WeightFile {
	t.Helper()
	seen := map[string]bool{}
	var files []WeightFile
	for _, p := range blobFilePath.FindAllString(milText, -1) {
		if seen[p] {
			continue
		}
		seen[p] = true
		files = append(files, WeightFile{Path: p, Blob: []byte(p)})
	}
	if len(files) == 0 {
		t.Fatal("MIL text references no BLOBFILE paths; testdata is not the expected program")
	}
	return files
}

// TestCompileModelIRProgram compiles real modelir output end to end through
// the exported bridge: no adapter code lives in this test that a caller would
// not also be able to reach.
func TestCompileModelIRProgram(t *testing.T) {
	milBytes, err := os.ReadFile(filepath.Join("testdata", "modelir_transformer.mil"))
	if err != nil {
		t.Fatal(err)
	}
	milText := string(milBytes)

	var desc ModelDescription
	for _, v := range modelIRInputs {
		fd, err := TensorFeatureDescription(v.name, v.dtype, v.shape)
		if err != nil {
			t.Fatalf("input %s: %v", v.name, err)
		}
		desc.Inputs = append(desc.Inputs, fd)
	}
	for _, v := range modelIROutputs {
		fd, err := TensorFeatureDescription(v.name, v.dtype, v.shape)
		if err != nil {
			t.Fatalf("output %s: %v", v.name, err)
		}
		desc.Outputs = append(desc.Outputs, fd)
	}
	for _, v := range modelIRStates {
		fd, err := StateFeatureDescription(v.name, v.dtype, v.shape)
		if err != nil {
			t.Fatalf("state %s: %v", v.name, err)
		}
		desc.States = append(desc.States, fd)
	}

	files := weightFilesFromMILText(t, milText)
	tmp := t.TempDir()
	weightRoot := filepath.Join(tmp, "weightroot")
	if err := WriteWeightRoot(weightRoot, files); err != nil {
		t.Fatalf("WriteWeightRoot: %v", err)
	}

	out := filepath.Join(tmp, "model.mlmodelc")
	if err := CompileMILText(milText, 9, desc, weightRoot, out); err != nil {
		t.Fatalf("CompileMILText: %v", err)
	}

	for _, name := range []string{"model.mil", "coremldata.bin", "metadata.json"} {
		if _, err := os.Stat(filepath.Join(out, name)); err != nil {
			t.Errorf("bundle is missing %s: %v", name, err)
		}
	}
	// Every BLOBFILE path the program reads must resolve inside the bundle.
	for _, f := range files {
		rel := strings.TrimPrefix(f.Path, ModelPathPrefix)
		got, err := os.ReadFile(filepath.Join(out, filepath.FromSlash(rel)))
		if err != nil {
			t.Errorf("bundle is missing weight %s: %v", rel, err)
			continue
		}
		if string(got) != f.Path {
			t.Errorf("weight %s has the wrong contents", rel)
		}
	}
	// Everything the bridge built must satisfy the package's own wire rules.
	for _, fd := range desc.Inputs {
		if err := validateFeatureDescription(fd, false); err != nil {
			t.Errorf("input %s: %v", fd.Name, err)
		}
	}
	for _, fd := range desc.Outputs {
		if err := validateFeatureDescription(fd, false); err != nil {
			t.Errorf("output %s: %v", fd.Name, err)
		}
	}
	for _, fd := range desc.States {
		if err := validateFeatureDescription(fd, true); err != nil {
			t.Errorf("state %s: %v", fd.Name, err)
		}
	}
}

// TestModelIRProgramOpsAreCovered records how much of a real modelir program
// the op-schema table describes. The table is advisory: validateOpSchema skips
// an op it has no schema for, so an op missing here is one nothing checks.
func TestModelIRProgramOpsAreCovered(t *testing.T) {
	milBytes, err := os.ReadFile(filepath.Join("testdata", "modelir_transformer.mil"))
	if err != nil {
		t.Fatal(err)
	}
	counts := milTextOpCounts(string(milBytes))
	if len(counts) < 10 {
		t.Fatalf("found %d op types in testdata, expected a whole transformer graph", len(counts))
	}
	var uncovered []string
	var covered, skipped int
	for _, op := range sortedKeys(counts) {
		if _, ok := opschema.Lookup("ios18", op); ok {
			covered += counts[op]
			continue
		}
		skipped += counts[op]
		uncovered = append(uncovered, op)
	}
	sort.Strings(uncovered)
	t.Logf("ios18: %d/%d op types, %d/%d invocations described by the table",
		len(counts)-len(uncovered), len(counts), covered, covered+skipped)
	if len(uncovered) > 0 {
		t.Errorf("no schema for %v; add them to internal/opschema/gen/dumpops.py, "+
			"or validateOpSchema silently waves them through", uncovered)
	}
}

// An op call is either a bare statement, "write_state(...)", or an assignment,
// "tensor<fp16, [1,8]> y = add(...)". The prefix must not cross an "=", or the
// scan would also match the string(...) inside an op's attribute list.
var milOpCall = regexp.MustCompile(`^\s*(?:[^=]*=\s*)?([a-z_][a-z0-9_]*)\(`)

// milTextOpCounts counts op invocations per op type in MIL text. It is a
// lexical scan, not a parser: enough to census which ops a program uses.
func milTextOpCounts(milText string) map[string]int {
	counts := map[string]int{}
	for _, line := range strings.Split(milText, "\n") {
		m := milOpCall.FindStringSubmatch(line)
		if m == nil {
			continue
		}
		// "program(1.3)" is the file header, not an op.
		if m[1] == "program" {
			continue
		}
		counts[m[1]]++
	}
	return counts
}
