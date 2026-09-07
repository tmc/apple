//go:build darwin

package coremlcompiler

import (
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"unsafe"

	"github.com/tmc/apple/coreml"
)

// Where a constexpr op's parameters belong -- the operation's attributes or
// its inputs -- is decided by the opset the op CLASS was declared at, not by
// the program's opset. constexpr_affine_dequantize is declared only in
// coremltools' ops/defs/iOS16 and is not among the constexpr ops iOS18
// redefines, so it takes attributes even inside an iOS18 program.
//
// That is a claim about the runtime, so this test asks the runtime. It builds
// one program each way and records which loads and which computes correctly,
// with a positive and a negative control: without them "accepted" and
// "rejected" cannot be told apart from a rig that always says one or the other.

const (
	constexprOut = 4
	constexprIn  = 8
)

func constexprTestWeights() []float32 {
	w := make([]float32, constexprOut*constexprIn)
	for r := 0; r < constexprOut; r++ {
		for c := 0; c < constexprIn; c++ {
			w[r*constexprIn+c] = float32(r+1) * (float32(c)/float32(constexprIn-1)*2 - 1)
		}
	}
	return w
}

func constexprTestInput() []float32 {
	x := make([]float32, constexprIn)
	for i := range x {
		x[i] = float32(i%3) - 1
	}
	return x
}

// constexprReference computes y = x * Wᵀ in float64 from the weights the model
// actually stores, so an arm is compared against the quantized values rather
// than the pre-quantization ideal.
func constexprReference(w, x []float32) []float64 {
	y := make([]float64, constexprOut)
	for r := 0; r < constexprOut; r++ {
		var s float64
		for c := 0; c < constexprIn; c++ {
			s += float64(x[c]) * float64(w[r*constexprIn+c])
		}
		y[r] = s
	}
	return y
}

func f32vt(dims ...uint64) *ValueType { return tt(DataTypeFloat32, dims...) }

func constexprProg(ops []*Operation) *Program {
	return &Program{
		Functions: map[string]*Function{
			"main": {
				Inputs: []NamedValueType{{Name: "x", Type: f32vt(1, constexprIn)}},
				OpSet:  "CoreML8", // iOS18
				BlockSpecializations: map[string]*Block{
					"CoreML8": {Operations: ops, Outputs: []string{"y"}},
				},
			},
		},
	}
}

func constexprLinear() *Operation {
	return &Operation{
		Type:    "linear",
		Inputs:  map[string]*Argument{"x": ref("x"), "weight": ref("w")},
		Outputs: []NamedValueType{{Name: "y", Type: f32vt(1, constexprOut)}},
	}
}

// constexprArm compiles, loads and predicts, reporting how far it got.
func constexprArm(t *testing.T, label string, prog *Program, files []WeightFile, stored []float32) (loaded bool, maxRelErr float64, detail string) {
	t.Helper()
	dir := t.TempDir()
	weightRoot := ""
	if len(files) > 0 {
		weightRoot = filepath.Join(dir, "weightroot")
		if err := WriteWeightRoot(weightRoot, files); err != nil {
			return false, 0, fmt.Sprintf("WriteWeightRoot: %v", err)
		}
	}
	in, err := TensorFeatureDescription("x", "fp32", []int64{1, constexprIn})
	if err != nil {
		t.Fatal(err)
	}
	out, err := TensorFeatureDescription("y", "fp32", []int64{1, constexprOut})
	if err != nil {
		t.Fatal(err)
	}
	desc := ModelDescription{Inputs: []FeatureDescription{in}, Outputs: []FeatureDescription{out}}

	bundle := filepath.Join(dir, "m.mlmodelc")
	if err := CompileProgram(prog, 9, desc, weightRoot, bundle); err != nil {
		return false, 0, fmt.Sprintf("CompileProgram rejected: %v", err)
	}
	if mil, rerr := os.ReadFile(filepath.Join(bundle, "model.mil")); rerr == nil {
		for _, line := range strings.Split(string(mil), "\n") {
			if strings.Contains(line, "constexpr_affine_dequantize") {
				t.Logf("  [%s] emitted: %s", label, strings.TrimSpace(line))
			}
		}
	}
	model, err := LoadCoreMLModel(bundle)
	if err != nil {
		return false, 0, fmt.Sprintf("runtime REJECTED at load: %v", scrubTemp(err.Error(), dir))
	}
	defer model.Close()

	x := constexprTestInput()
	got, err := model.Predict([]PredictInput{{
		Name:    "x",
		Data:    unsafe.Pointer(&x[0]),
		Shape:   []int{1, constexprIn},
		Strides: []int{constexprIn, 1},
		DType:   coreml.MLMultiArrayDataTypeFloat32,
	}}, "y")
	if err != nil {
		return true, 0, fmt.Sprintf("loaded but prediction failed: %v", scrubTemp(err.Error(), dir))
	}
	vals := make([]float64, constexprOut)
	for i := range vals {
		switch got.DType {
		case coreml.MLMultiArrayDataTypeFloat32:
			vals[i] = float64(math.Float32frombits(binary.LittleEndian.Uint32(got.Bytes[i*4:])))
		case coreml.MLMultiArrayDataTypeFloat16:
			vals[i] = float64(Float16frombits(binary.LittleEndian.Uint16(got.Bytes[i*2:])))
		default:
			return true, 0, fmt.Sprintf("unexpected output dtype %v", got.DType)
		}
	}
	want := constexprReference(stored, x)
	scale := 0.0
	for _, v := range want {
		scale = math.Max(scale, math.Abs(v))
	}
	for i := range want {
		if math.IsNaN(vals[i]) || math.IsInf(vals[i], 0) {
			return true, math.Inf(1), "prediction contains a non-finite value"
		}
		if e := math.Abs(vals[i]-want[i]) / math.Max(scale, 1e-9); e > maxRelErr {
			maxRelErr = e
		}
	}
	return true, maxRelErr, "loaded and predicted"
}

func scrubTemp(s, dir string) string { return strings.ReplaceAll(s, dir, "<tmp>") }

func TestConstexprPlacementRuntime(t *testing.T) {
	for _, c := range [][]string{{"sw_vers", "-productVersion"}, {"sw_vers", "-buildVersion"}, {"uname", "-m"}} {
		out, err := exec.Command(c[0], c[1:]...).Output()
		if err == nil {
			t.Logf("host: %s %v = %s", c[0], c[1:], strings.TrimSpace(string(out)))
		}
	}

	q, packed, err := QuantizeAffinePerChannel(constexprTestWeights(), []int64{constexprOut, constexprIn}, 0, DataTypeInt8)
	if err != nil {
		t.Fatal(err)
	}
	q.OutputType = DataTypeFloat32 // keep the whole graph fp32
	stored, err := q.Dequantize(packed)
	if err != nil {
		t.Fatal(err)
	}
	files, refs, err := BuildWeights([]WeightTensor{{
		Name: "w", DType: BlobDataTypeInt8, Data: packed,
	}}, BlobLayoutSingleFile)
	if err != nil {
		t.Fatal(err)
	}

	// Positive control: no constexpr at all. If this fails, nothing below is
	// interpretable -- it would mean the rig, not the spelling, is the problem.
	plainConst := &Operation{
		Type:    "const",
		Outputs: []NamedValueType{{Name: "w", Type: f32vt(constexprOut, constexprIn)}},
		Attributes: map[string]*Value{"val": {
			Type:      f32vt(constexprOut, constexprIn),
			Immediate: &ImmediateValue{Tensor: &TensorValue{Floats: stored}},
		}},
	}
	okC, errC, detailC := constexprArm(t, "control-plain-const",
		constexprProg([]*Operation{plainConst, constexprLinear()}), nil, stored)
	t.Logf("CONTROL plain fp32 const    : loaded=%v maxRelErr=%.3g  %s", okC, errC, detailC)
	if !okC {
		t.Fatalf("positive control failed; no arm below is interpretable: %s", detailC)
	}
	if errC > 1e-5 {
		t.Fatalf("positive control is numerically wrong (%.3g); the rig is not trustworthy", errC)
	}

	// Negative control: the accepted spelling with a BLOBFILE offset pointing
	// at the file header instead of a metadata block. It must use the spelling
	// that otherwise loads, or it would fail for the reason under test and
	// prove nothing.
	badRef := refs["w"]
	badRef.Offset = 0
	badOp, err := q.Operation("w", badRef)
	if err != nil {
		t.Fatal(err)
	}
	okN, _, detailN := constexprArm(t, "control-bad-offset",
		constexprProg([]*Operation{badOp, constexprLinear()}), files, stored)
	t.Logf("CONTROL bad BLOBFILE offset : loaded=%v  %s", okN, detailN)
	if okN {
		t.Errorf("NEGATIVE CONTROL FAILED: a BLOBFILE offset pointing at the file header was "+
			"accepted; 'accepted' cannot be trusted below: %s", detailN)
	}

	// Arm: parameters as attributes, what Operation produces.
	attrsOp, err := q.Operation("w", refs["w"])
	if err != nil {
		t.Fatal(err)
	}
	okA, errA, detailA := constexprArm(t, "attributes",
		constexprProg([]*Operation{attrsOp, constexprLinear()}), files, stored)
	t.Logf("ARM attributes (ours)       : loaded=%v maxRelErr=%.3g  %s", okA, errA, detailA)

	// Arm: the same parameters moved to the operation's inputs.
	inputsOp := &Operation{
		Type:    attrsOp.Type,
		Outputs: attrsOp.Outputs,
		Inputs:  map[string]*Argument{},
	}
	for k, v := range attrsOp.Attributes {
		inputsOp.Inputs[k] = &Argument{Bindings: []Binding{{Value: v}}}
	}
	okB, errB, detailB := constexprArm(t, "inputs",
		constexprProg([]*Operation{inputsOp, constexprLinear()}), files, stored)
	t.Logf("ARM inputs                  : loaded=%v maxRelErr=%.3g  %s", okB, errB, detailB)

	if !okA {
		t.Errorf("the attributes spelling was rejected by the runtime; Operation emits a program that will not load: %s", detailA)
	}
	if errA > 1e-5 {
		t.Errorf("the attributes spelling loaded but computed wrong values (maxRelErr %.3g)", errA)
	}
	if okB {
		t.Errorf("the inputs spelling was ALSO accepted; both spellings work and this test's premise needs revisiting: %s", detailB)
	}
}
