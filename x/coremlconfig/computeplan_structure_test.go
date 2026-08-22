//go:build darwin

package coremlconfig_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/tmc/apple/x/coremlcompiler"
	"github.com/tmc/apple/x/coremlcompiler/nnproto"
	"github.com/tmc/apple/x/coremlconfig"
)

// TestPlanUnreadableStructureIsNotZeroCost loads a plan for a model whose
// structure is not an ML program. CoreML exposes per-operation devices
// only for the mlprogram form, so nothing can be read; the plan must say
// so rather than present the resulting zeros as a placement measurement.
func TestPlanUnreadableStructureIsNotZeroCost(t *testing.T) {
	const dim, hidden, seq = 256, 1024, 32
	weights := make([]float32, hidden*dim)
	for i := range weights {
		weights[i] = float32(i%7) / 1000
	}
	proto, err := nnproto.FFNReLU2(dim, hidden, seq, weights, weights[:dim*hidden])
	if err != nil {
		t.Fatal(err)
	}
	mlmodel := filepath.Join(t.TempDir(), "ffn.mlmodel")
	if err := os.WriteFile(mlmodel, proto, 0o644); err != nil {
		t.Fatal(err)
	}
	modelc, err := coremlcompiler.CompileMLModelAtURL(mlmodel)
	if err != nil {
		t.Skipf("CoreML compiler unavailable: %v", err)
	}
	p, err := coremlconfig.LoadPlan(modelc, coremlconfig.PlanOptions{})
	if err != nil {
		t.Fatalf("LoadPlan: %v", err)
	}
	if p.Analyzable() {
		t.Fatalf("Analyzable = true for a NeuralNetwork model; got %d operations", len(p.Operations()))
	}
	if p.Unreadable() == "" {
		t.Error("Unreadable is empty while Analyzable is false")
	}
	// The zeros are still zero; the point is that they are now labelled.
	if len(p.Operations()) != 0 || p.ANEFraction() != 0 {
		t.Errorf("unreadable plan reported ops=%d ANE=%v", len(p.Operations()), p.ANEFraction())
	}
	t.Logf("unreadable: %s", p.Unreadable())
}
