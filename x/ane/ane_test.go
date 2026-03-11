//go:build darwin

package ane

import (
	"errors"
	"fmt"
	"math"
	"testing"

	"github.com/tmc/apple/x/ane/mil"
)

func openOrSkip(t *testing.T) *Runtime {
	t.Helper()
	rt, err := Open()
	if errors.Is(err, ErrNoANE) {
		t.Skip("no ANE available")
	}
	if err != nil {
		t.Fatal(err)
	}
	return rt
}

func TestProbe(t *testing.T) {
	info, err := Probe()
	if errors.Is(err, ErrNoANE) {
		t.Skip("no ANE available")
	}
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("HasANE=%v NumCores=%d Arch=%s Product=%s Build=%s IsVM=%v",
		info.HasANE, info.NumCores, info.Architecture, info.Product, info.BuildVersion, info.IsVM)

	if !info.HasANE {
		t.Error("HasANE should be true when Probe succeeds without ErrNoANE")
	}
	if info.NumCores == 0 {
		t.Error("NumCores should be > 0")
	}
}

func TestOpenClose(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	info := rt.Info()
	if !info.HasANE {
		t.Error("runtime info should report HasANE=true")
	}
	if rt.CompileCount() != 0 {
		t.Error("compile count should be 0 after open")
	}
}

func TestCompileMIL1Channel(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	const channels = 1
	milText := mil.GenIdentity(channels, 1)
	blob, err := mil.BuildIdentityWeightBlob(channels)
	if err != nil {
		t.Fatal(err)
	}

	k, err := rt.Compile(CompileOptions{
		ModelType:  ModelTypeMIL,
		MILText:    []byte(milText),
		WeightBlob: blob,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer k.Close()

	input := []float32{42}
	if err := k.WriteInputF32(0, input); err != nil {
		t.Fatal(err)
	}
	if err := k.Eval(); err != nil {
		t.Fatal(err)
	}

	output := make([]float32, channels)
	if err := k.ReadOutputF32(0, output); err != nil {
		t.Fatal(err)
	}
	t.Logf("input=%v output=%v", input, output)

	if diff := input[0] - output[0]; diff < -0.1 || diff > 0.1 {
		t.Errorf("output = %v, want %v", output[0], input[0])
	}
}

func TestCompileMILIdentity(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	const channels = 4
	milText := mil.GenIdentity(channels, 1)
	blob, err := mil.BuildIdentityWeightBlob(channels)
	if err != nil {
		t.Fatal(err)
	}

	k, err := rt.Compile(CompileOptions{
		ModelType:  ModelTypeMIL,
		MILText:    []byte(milText),
		WeightBlob: blob,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer k.Close()

	if k.NumInputs() == 0 || k.NumOutputs() == 0 {
		t.Fatal("kernel has no inputs or outputs")
	}

	input := []float32{1, 2, 3, 4}
	if err := k.WriteInputF32(0, input); err != nil {
		t.Fatal(err)
	}
	if err := k.Eval(); err != nil {
		t.Fatal(err)
	}

	// Try fp32 first, fall back to fp16.
	output := make([]float32, channels)
	if err := k.ReadOutputF32(0, output); err != nil {
		t.Fatal(err)
	}

	allMatch := true
	for i := range input {
		diff := input[i] - output[i]
		if diff < -0.1 || diff > 0.1 {
			allMatch = false
			break
		}
	}

	if !allMatch {
		// ANE outputs fp16 internally; try fp16 readback.
		if err := k.ReadOutputFP16(0, output); err != nil {
			t.Fatal(err)
		}
		t.Logf("fp32 mismatch, using fp16 readback: %v", output)
	}

	for i := range input {
		diff := input[i] - output[i]
		if diff < -0.1 || diff > 0.1 {
			t.Errorf("output[%d] = %v, want %v", i, output[i], input[i])
		}
	}
}

func TestCompileMILSpatial(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	const spatial = 4
	milText := mil.GenScaleFP16IO(spatial)
	weights := []float32{2.0}
	blob, err := mil.BuildWeightBlob(weights, 1, 1)
	if err != nil {
		t.Fatal(err)
	}

	k, err := rt.Compile(CompileOptions{
		ModelType:  ModelTypeMIL,
		MILText:    []byte(milText),
		WeightBlob: blob,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer k.Close()

	input := []float32{1, 2, 3, 4}
	if err := k.WriteInputFP16(0, input); err != nil {
		t.Fatal(err)
	}
	if err := k.Eval(); err != nil {
		t.Fatal(err)
	}

	output := make([]float32, spatial)
	if err := k.ReadOutputFP16(0, output); err != nil {
		t.Fatal(err)
	}
	t.Logf("output=%v (want [2 4 6 8])", output)

	want := []float32{2, 4, 6, 8}
	for i := range want {
		diff := want[i] - output[i]
		if diff < -0.1 || diff > 0.1 {
			t.Errorf("output[%d] = %v, want %v", i, output[i], want[i])
		}
	}
}

func TestCompileMILSoftmaxWeightless(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	const channels = 4
	milText := softmaxMIL(channels, 1)

	k, err := rt.Compile(CompileOptions{
		ModelType: ModelTypeMIL,
		MILText:   []byte(milText),
	})
	if err != nil {
		t.Fatal(err)
	}
	defer k.Close()

	input := []float32{1, 2, 3, 4}
	if err := k.WriteInputFP16(0, input); err != nil {
		t.Fatal(err)
	}
	if err := k.Eval(); err != nil {
		t.Fatal(err)
	}

	output := make([]float32, channels)
	if err := k.ReadOutputFP16(0, output); err != nil {
		t.Fatal(err)
	}

	var sum float32
	for i, v := range output {
		if v <= 0 {
			t.Fatalf("output[%d] = %v, want > 0", i, v)
		}
		sum += v
	}
	diff := sum - 1
	if diff < 0 {
		diff = -diff
	}
	if diff > 0.02 {
		t.Fatalf("softmax sum = %v, want 1", sum)
	}
}

func TestCompileMILMultiWeightFFN(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	const (
		dim    = 4
		hidden = 8
		seq    = 2
	)
	w1 := repeatedWeights(hidden*dim, 0.25)
	w3 := repeatedWeights(hidden*dim, 0.5)
	w2 := repeatedWeights(dim*hidden, 0.125)

	w1Blob, err := mil.BuildWeightBlob(w1, hidden, dim)
	if err != nil {
		t.Fatal(err)
	}
	w2Blob, err := mil.BuildWeightBlob(w2, dim, hidden)
	if err != nil {
		t.Fatal(err)
	}
	w3Blob, err := mil.BuildWeightBlob(w3, hidden, dim)
	if err != nil {
		t.Fatal(err)
	}

	k, err := rt.Compile(CompileOptions{
		ModelType: ModelTypeMIL,
		MILText:   []byte(ffnForwardTapsMIL(dim, hidden, seq)),
		WeightFiles: []WeightFile{
			{Path: "@model_path/weights/w1.bin", Blob: w1Blob},
			{Path: "@model_path/weights/w2.bin", Blob: w2Blob},
			{Path: "@model_path/weights/w3.bin", Blob: w3Blob},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer k.Close()

	input := repeatedWeights(dim*seq, 1)
	if err := k.WriteInputFP16(0, input); err != nil {
		t.Fatal(err)
	}
	if err := k.Eval(); err != nil {
		t.Fatal(err)
	}

	output := make([]float32, (dim+3*hidden)*seq)
	if err := k.ReadOutputFP16(0, output); err != nil {
		t.Fatal(err)
	}
	nonZero := false
	for i, v := range output {
		if math.IsNaN(float64(v)) || math.IsInf(float64(v), 0) {
			t.Fatalf("output[%d] is not finite: %v", i, v)
		}
		if math.Abs(float64(v)) > 1e-4 {
			nonZero = true
		}
	}
	if !nonZero {
		t.Fatal("multi-weight FFN output is all zero")
	}
}

func TestEvalRepeated(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	const channels = 1
	milText := mil.GenIdentity(channels, 1)
	blob, err := mil.BuildIdentityWeightBlob(channels)
	if err != nil {
		t.Fatal(err)
	}

	k, err := rt.Compile(CompileOptions{
		ModelType:  ModelTypeMIL,
		MILText:    []byte(milText),
		WeightBlob: blob,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer k.Close()

	for iter := 0; iter < 10; iter++ {
		input := []float32{float32(iter + 1)}
		if err := k.WriteInputF32(0, input); err != nil {
			t.Fatal(err)
		}
		if err := k.Eval(); err != nil {
			t.Fatal(err)
		}
		output := make([]float32, channels)
		if err := k.ReadOutputF32(0, output); err != nil {
			t.Fatal(err)
		}
		diff := input[0] - output[0]
		if diff < -0.1 || diff > 0.1 {
			t.Errorf("iter %d: output = %v, want %v", iter, output[0], input[0])
		}
	}
}

func TestEvalWithStats(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	const channels = 1
	milText := mil.GenIdentity(channels, 1)
	blob, err := mil.BuildIdentityWeightBlob(channels)
	if err != nil {
		t.Fatal(err)
	}

	k, err := rt.Compile(CompileOptions{
		ModelType:  ModelTypeMIL,
		MILText:    []byte(milText),
		WeightBlob: blob,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer k.Close()

	input := []float32{7}
	if err := k.WriteInputF32(0, input); err != nil {
		t.Fatal(err)
	}

	stats, err := k.EvalWithStats()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("HWExecutionNS=%d", stats.HWExecutionNS)

	output := make([]float32, channels)
	if err := k.ReadOutputF32(0, output); err != nil {
		t.Fatal(err)
	}
	diff := input[0] - output[0]
	if diff < -0.1 || diff > 0.1 {
		t.Errorf("output = %v, want %v", output[0], input[0])
	}
}

func TestCompileCount(t *testing.T) {
	rt := openOrSkip(t)
	defer rt.Close()

	if rt.CompileCount() != 0 {
		t.Fatal("compile count should start at 0")
	}

	const channels = 1
	milText := mil.GenIdentity(channels, 1)
	blob, err := mil.BuildIdentityWeightBlob(channels)
	if err != nil {
		t.Fatal(err)
	}

	k, err := rt.Compile(CompileOptions{
		ModelType:  ModelTypeMIL,
		MILText:    []byte(milText),
		WeightBlob: blob,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer k.Close()

	if rt.CompileCount() != 1 {
		t.Errorf("compile count = %d, want 1", rt.CompileCount())
	}
}

func TestFP16Conversion(t *testing.T) {
	tests := []struct {
		in   float32
		want float32
	}{
		{0, 0},
		{1, 1},
		{-1, -1},
		{0.5, 0.5},
		{65504, 65504},
	}
	for _, tt := range tests {
		h := Float32ToFP16(tt.in)
		got := FP16ToFloat32(h)
		diff := tt.want - got
		if diff < 0 {
			diff = -diff
		}
		if diff > 0.01 {
			t.Errorf("FP16 round-trip(%v): got %v, want %v", tt.in, got, tt.want)
		}
	}
}

func TestErrorClassification(t *testing.T) {
	e := &ANEError{Op: "compile", Class: "ANE", Code: 42, Err: ErrCompileBudgetExhausted}
	if !errors.Is(e, ErrCompileBudgetExhausted) {
		t.Error("ANEError should unwrap to ErrCompileBudgetExhausted")
	}

	target := &ANEError{Op: "compile", Class: "ANE", Code: 42}
	if !errors.Is(e, target) {
		t.Error("ANEError should match target with same Op/Class/Code")
	}

	other := &ANEError{Op: "eval", Class: "ANE", Code: 42}
	if errors.Is(e, other) {
		t.Error("ANEError should not match target with different Op")
	}
}

func softmaxMIL(channels, spatial int) string {
	return fmt.Sprintf(`program(1.3)
[buildInfo = dict<string, string>({{"coremlc-component-MIL", "3510.2.1"}, {"coremlc-version", "3505.4.1"}, {"coremltools-component-milinternal", ""}, {"coremltools-version", "9.0"}})]
{
    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x) {
        int32 ax = const()[name = string("ax"), val = int32(1)];
        tensor<fp16, [1, %d, 1, %d]> y = softmax(axis = ax, x = x)[name = string("sm")];
    } -> (y);
}
`, channels, spatial, channels, spatial)
}

func ffnForwardTapsMIL(dim, hidden, seq int) string {
	return fmt.Sprintf(`program(1.3)
[buildInfo = dict<string, string>({{"coremlc-component-MIL", "3510.2.1"}, {"coremlc-version", "3505.4.1"}, {"coremltools-component-milinternal", ""}, {"coremltools-version", "9.0"}})]
{
    func main<ios18>(tensor<fp16, [1, %d, 1, %d]> x) {
        string pt = const()[name=string("pt"), val=string("valid")];
        tensor<int32, [2]> st = const()[name=string("st"), val=tensor<int32, [2]>([1,1])];
        tensor<int32, [4]> pd = const()[name=string("pd"), val=tensor<int32, [4]>([0,0,0,0])];
        tensor<int32, [2]> dl = const()[name=string("dl"), val=tensor<int32, [2]>([1,1])];
        int32 gr = const()[name=string("gr"), val=int32(1)];
        tensor<fp16, [%d,%d,1,1]> W1 = const()[name=string("W1"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string("@model_path/weights/w1.bin"), offset=uint64(64)))];
        tensor<fp16, [%d,%d,1,1]> W3 = const()[name=string("W3"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string("@model_path/weights/w3.bin"), offset=uint64(64)))];
        tensor<fp16, [%d,%d,1,1]> W2 = const()[name=string("W2"), val=tensor<fp16, [%d,%d,1,1]>(BLOBFILE(path=string("@model_path/weights/w2.bin"), offset=uint64(64)))];
        tensor<fp16, [1,%d,1,%d]> h1 = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=W1,x=x)[name=string("c1")];
        tensor<fp16, [1,%d,1,%d]> h3 = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=W3,x=x)[name=string("c3")];
        tensor<fp16, [1,%d,1,%d]> sig = sigmoid(x=h1)[name=string("sg")];
        tensor<fp16, [1,%d,1,%d]> silu = mul(x=h1,y=sig)[name=string("si")];
        tensor<fp16, [1,%d,1,%d]> gate = mul(x=silu,y=h3)[name=string("gt")];
        tensor<fp16, [1,%d,1,%d]> out = conv(dilations=dl,groups=gr,pad=pd,pad_type=pt,strides=st,weight=W2,x=gate)[name=string("c2")];
        int32 cax = const()[name=string("cax"), val=int32(1)];
        bool cid = const()[name=string("cid"), val=bool(false)];
        tensor<fp16, [1,%d,1,%d]> taps = concat(axis=cax,interleave=cid,values=(out,h1,h3,gate))[name=string("cat")];
    } -> (taps);
}
`, dim, seq,
		hidden, dim, hidden, dim,
		hidden, dim, hidden, dim,
		dim, hidden, dim, hidden,
		hidden, seq,
		hidden, seq,
		hidden, seq,
		hidden, seq,
		hidden, seq,
		dim, seq,
		dim+3*hidden, seq,
	)
}

func repeatedWeights(n int, v float32) []float32 {
	out := make([]float32, n)
	for i := range out {
		out[i] = v
	}
	return out
}
