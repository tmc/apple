//go:build darwin

package engineattest

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/tmc/apple/x/ane"
	"github.com/tmc/apple/x/ane/mil"
	"github.com/tmc/apple/x/ane/model"
	"github.com/tmc/apple/x/coremlcompiler"
	"github.com/tmc/apple/x/coremlcompiler/nnproto"
)

// Model shape for the package-backed tests: a small FFN, the same
// NeuralNetwork-protobuf + compileModelAtURL route the anepingpong
// probes proved end-to-end. (The pure-Go compiler's .mlmodelc output is
// currently rejected by the ANE espresso translator — see
// design/bugreport-ane-package-fp32-identity.md — so Apple's compiler
// does the packaging here.)
const testDim, testHidden, testSeq = 256, 1024, 32

// compilePackageOrSkip builds an FFN .mlmodel, compiles it with CoreML's
// compileModelAtURL, and loads it package-backed, which is the route
// that reports real ANE hardware execution time. In-memory MIL models
// report none — see TestANEInMemoryMILUnattestable.
func compilePackageOrSkip(t *testing.T) *model.Kernel {
	t.Helper()
	weights := make([]float32, testHidden*testDim)
	for i := range weights {
		weights[i] = float32(i%7) / 1000
	}
	proto, err := nnproto.FFNReLU2(testDim, testHidden, testSeq, weights, weights[:testDim*testHidden])
	if err != nil {
		t.Fatal(err)
	}
	mlmodel := filepath.Join(t.TempDir(), "ffn.mlmodel")
	if err := os.WriteFile(mlmodel, proto, 0o644); err != nil {
		t.Fatal(err)
	}
	modelc, err := coremlcompiler.CompileMLModelAtURL(mlmodel)
	if err != nil {
		t.Fatalf("CompileMLModelAtURL: %v", err)
	}
	k, err := model.Compile(model.CompileOptions{PackagePath: modelc, PerfStatsMask: 0xF})
	if err != nil {
		t.Skipf("package-backed ANE compile unavailable: %v", err)
	}
	t.Cleanup(k.Close)
	input := make([]float32, testDim*testSeq)
	for i := range input {
		input[i] = float32(i%5) / 10
	}
	if err := k.WriteInputFP16(0, input); err != nil {
		t.Fatal(err)
	}
	return k
}

// attestableOrSkip skips when the driver reports zero hardware time for
// a real evaluation of a package-backed model — ErrUnattestable from the
// sensitivity canary. Observed on macOS 26.6: every perf counter stays
// zero for package-backed models too (the 2026-08-12 264µs measurement
// no longer reproduces), so the counter-based tests below cannot judge
// anything here. The skip is environmental and loud; compile failures
// stay fatal so a broken compile route can never look green. Release
// gates must run with -v (or assert no SKIPs) — a skipped ANE arm still
// exits 0.
func attestableOrSkip(t *testing.T, k *model.Kernel) {
	t.Helper()
	err := ANE(k.Model(), func() error { return nil })
	if errors.Is(err, ErrUnattestable) {
		t.Skipf("ANE perf counters do not populate on this firmware: %v", err)
	}
}

func TestANERealEvalPasses(t *testing.T) {
	k := compilePackageOrSkip(t)
	attestableOrSkip(t, k)
	if err := ANE(k.Model(), func() error { return k.Eval() }); err != nil {
		t.Errorf("real eval failed attestation: %v", err)
	}
}

// TestANEImpostorFails is the reason this package exists: a function
// that computes the same result on the CPU and claims the ANE ran must
// fail the assertion.
func TestANEImpostorFails(t *testing.T) {
	k := compilePackageOrSkip(t)
	attestableOrSkip(t, k)
	err := ANE(k.Model(), func() error {
		out := make([]float32, testDim*testSeq)
		_ = out // "the FFN", computed on the CPU
		return nil
	})
	if !errors.Is(err, ErrDidNotRun) {
		t.Errorf("impostor: got %v, want ErrDidNotRun", err)
	}
}

// TestANEErrorPropagates verifies a failing fn is returned unchanged,
// not converted into a verdict about the hardware.
func TestANEErrorPropagates(t *testing.T) {
	k := compilePackageOrSkip(t)
	attestableOrSkip(t, k)
	want := errors.New("workload failed")
	err := ANE(k.Model(), func() error { return want })
	if !errors.Is(err, want) {
		t.Errorf("got %v, want the workload's own error", err)
	}
	if errors.Is(err, ErrDidNotRun) || errors.Is(err, ErrUnattestable) {
		t.Errorf("workload error was converted into a verdict: %v", err)
	}
}

// TestANEInMemoryMILUnattestable: in-memory MIL models report no
// hardware execution time (measured 2026-08-12: HwExecutionTime and all
// 24 named perf counters read zero after a successful eval), so the
// sensitivity canary must refuse to judge rather than report a false
// "did not run".
func TestANEInMemoryMILUnattestable(t *testing.T) {
	c, err := ane.Open()
	if errors.Is(err, ane.ErrNoANE) {
		t.Skip("no ANE available")
	}
	if err != nil {
		t.Fatal(err)
	}
	defer c.Close()
	blob, err := mil.BuildIdentityWeightBlob(4)
	if err != nil {
		t.Fatal(err)
	}
	m, err := c.Compile(ane.CompileOptions{
		ModelType:  ane.ModelTypeMIL,
		MILText:    []byte(mil.GenIdentity(4, 1)),
		WeightBlob: blob,
	})
	if err != nil {
		t.Fatal(err)
	}
	defer m.Close()
	if err := m.WriteInputF32(0, []float32{1, 2, 3, 4}); err != nil {
		t.Fatal(err)
	}
	err = ANE(m, func() error { return m.Eval() })
	if !errors.Is(err, ErrUnattestable) {
		t.Errorf("in-memory MIL: got %v, want ErrUnattestable", err)
	}
}

func TestANENilModel(t *testing.T) {
	if err := ANE(nil, func() error { return nil }); err == nil {
		t.Error("nil model: got nil error")
	}
}
