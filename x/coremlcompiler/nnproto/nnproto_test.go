// Test macOS environment note: Tested on macOS 26.6.1 (Darwin 25.6.0) arm64.
package nnproto

import (
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"runtime"
	"testing"

	anemodel "github.com/tmc/apple/x/ane/model"
	"github.com/tmc/apple/x/coremlcompiler"
)

func refFFN(w1, w2, x []float32, d, f, seq int) []float64 {
	x64 := make([]float64, len(x))
	for i, v := range x {
		x64[i] = float64(v)
	}
	h := make([]float64, f*seq)
	for i := range f {
		for c := range d {
			wv := float64(w1[i*d+c])
			for j := range seq {
				h[i*seq+j] += wv * x64[c*seq+j]
			}
		}
	}
	for i, v := range h {
		if v < 0 {
			h[i] = 0
		} else {
			h[i] = v * v
		}
	}
	y := make([]float64, d*seq)
	for i := range d {
		for c := range f {
			wv := float64(w2[i*f+c])
			for j := range seq {
				y[i*seq+j] += wv * h[c*seq+j]
			}
		}
	}
	return y
}

func TestEncode(t *testing.T) {
	d, f, seq := 256, 1024, 32
	w1 := make([]float32, f*d)
	w2 := make([]float32, d*f)

	data, err := FFNReLU2(d, f, seq, w1, w2)
	if err != nil {
		t.Fatalf("FFNReLU2 failed: %v", err)
	}
	if len(data) == 0 {
		t.Fatalf("expected non-empty protobuf byte output")
	}

	// Negative dimension validation check
	_, errNeg := ppackedInt64(nil, 1, []int64{-1, 10})
	if errNeg == nil {
		t.Fatalf("expected error for negative dimension shape")
	}
}

func TestCompileAndEvalOnANE(t *testing.T) {
	if runtime.GOOS != "darwin" {
		t.Skip("skipping ANE compilation test on non-darwin host")
	}

	d, f, seq := 256, 1024, 32
	rng := rand.New(rand.NewSource(1))

	scaleW1 := 1.0 / math.Sqrt(float64(d))
	scaleW2 := 1.0 / math.Sqrt(float64(f))

	w1 := make([]float32, f*d)
	for i := range w1 {
		w1[i] = float32(rng.NormFloat64() * scaleW1)
	}

	w2 := make([]float32, d*f)
	for i := range w2 {
		w2[i] = float32(rng.NormFloat64() * scaleW2)
	}

	modelProto, err := FFNReLU2(d, f, seq, w1, w2)
	if err != nil {
		t.Fatalf("FFNReLU2 failed: %v", err)
	}

	tmpDir, err := os.MkdirTemp("", "nnproto-test-")
	if err != nil {
		t.Fatalf("failed to create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	protoPath := filepath.Join(tmpDir, "model.mlmodel")
	if err := os.WriteFile(protoPath, modelProto, 0644); err != nil {
		t.Fatalf("failed to write mlmodel file: %v", err)
	}

	modelcPath, err := coremlcompiler.CompileMLModelAtURL(protoPath)
	if err != nil {
		t.Fatalf("CoreML model compilation rejected by compiler: %v", err)
	}

	kernel, err := anemodel.Compile(anemodel.CompileOptions{PackagePath: modelcPath})
	if err != nil {
		t.Skipf("skipping test: ANE kernel hardware compilation unavailable: %v", err)
	}
	defer kernel.Close()

	// 1. Evaluate initial input data
	inputData1 := make([]float32, d*seq)
	for i := range inputData1 {
		inputData1[i] = float32(rng.NormFloat64())
	}

	if err := kernel.WriteInputFP16(0, inputData1); err != nil {
		t.Fatalf("WriteInputFP16 failed: %v", err)
	}
	if err := kernel.Eval(); err != nil {
		t.Fatalf("Eval failed: %v", err)
	}

	outputData1 := make([]float32, d*seq)
	if err := kernel.ReadOutputFP16(0, outputData1); err != nil {
		t.Fatalf("ReadOutputFP16 failed: %v", err)
	}

	// Parity check vs CPU reference (Step 4: relative error <= 2e-2)
	expectedOut1 := refFFN(w1, w2, inputData1, d, f, seq)
	maxAbs, worst := 0.0, 0.0
	for i, v := range expectedOut1 {
		if a := math.Abs(v); a > maxAbs {
			maxAbs = a
		}
		if dd := math.Abs(float64(outputData1[i]) - v); dd > worst {
			worst = dd
		}
	}
	if maxAbs == 0 {
		maxAbs = 1
	}
	rel := worst / maxAbs
	t.Logf("parity vs float64 CPU reference: max relative error %.2e", rel)
	if rel > 2e-2 {
		t.Fatalf("parity vs CPU reference exceeded tolerance 2e-2 (got %.2e)", rel)
	}

	// 2. Step 5: Change input and verify output moves (differs from outputData1)
	inputData2 := make([]float32, d*seq)
	for i := range inputData2 {
		inputData2[i] = inputData1[i] + 2.0
	}
	if err := kernel.WriteInputFP16(0, inputData2); err != nil {
		t.Fatalf("WriteInputFP16 for input2 failed: %v", err)
	}
	if err := kernel.Eval(); err != nil {
		t.Fatalf("Eval for input2 failed: %v", err)
	}
	outputData2 := make([]float32, d*seq)
	if err := kernel.ReadOutputFP16(0, outputData2); err != nil {
		t.Fatalf("ReadOutputFP16 for input2 failed: %v", err)
	}

	diffCount := 0
	for i := range outputData1 {
		if math.Abs(float64(outputData1[i]-outputData2[i])) > 1e-3 {
			diffCount++
		}
	}
	if diffCount == 0 {
		t.Fatalf("expected output to move when input changes, but outputs were identical")
	}
}
