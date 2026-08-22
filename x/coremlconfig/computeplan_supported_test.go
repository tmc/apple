//go:build darwin

package coremlconfig_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/tmc/apple/x/ane/mil"
	"github.com/tmc/apple/x/coremlcompiler"
	"github.com/tmc/apple/x/coremlconfig"
)

// TestPlanSupportedIsWiderThanPreferred pins the distinction that a zero
// ANE cost fraction does not mean the Neural Engine cannot run the model.
// On an FFN of fp16 convolutions CoreML reports every operation as
// supported on the ANE while preferring the CPU for all of them, so
// ANEFraction is 0 and SupportsANE is true at the same time.
func TestPlanSupportedIsWiderThanPreferred(t *testing.T) {
	const dim, hidden, seq = 256, 1024, 32
	dir := t.TempDir()
	weights := filepath.Join(dir, "weights")
	if err := os.MkdirAll(weights, 0o755); err != nil {
		t.Fatal(err)
	}
	vals := make([]float32, hidden*dim)
	for i := range vals {
		vals[i] = float32(i%7) / 1000
	}
	w1, err := mil.BuildWeightBlob(vals, hidden, dim)
	if err != nil {
		t.Fatal(err)
	}
	w2, err := mil.BuildWeightBlob(vals[:dim*hidden], dim, hidden)
	if err != nil {
		t.Fatal(err)
	}
	for name, blob := range map[string][]byte{"w1.bin": w1, "w2.bin": w2} {
		if err := os.WriteFile(filepath.Join(weights, name), blob, 0o644); err != nil {
			t.Fatal(err)
		}
	}
	out := filepath.Join(dir, "ffn.mlmodelc")
	desc := coremlcompiler.ModelDescription{
		Inputs:  []coremlcompiler.FeatureDescription{{Name: "x"}},
		Outputs: []coremlcompiler.FeatureDescription{{Name: "out"}},
	}
	if err := coremlcompiler.CompileMILText(mil.GenFFNForwardReLU2(dim, hidden, seq), 8, desc, dir, out); err != nil {
		t.Fatal(err)
	}
	p, err := coremlconfig.LoadPlan(out, coremlconfig.PlanOptions{})
	if err != nil {
		t.Fatalf("LoadPlan: %v", err)
	}
	if !p.Analyzable() {
		t.Fatalf("plan unreadable: %s", p.Unreadable())
	}
	var withSupported int
	for _, op := range p.Operations() {
		if len(op.Supported) > 0 {
			withSupported++
		}
	}
	if withSupported == 0 {
		t.Fatal("no operation reported any supported device; Supported is not being read")
	}
	if !p.SupportsANE() {
		t.Errorf("SupportsANE = false; ANE absent from every operation's supported set")
	}
	t.Logf("ANEFraction=%.3f SupportsANE=%v (%d/%d ops carry a supported set)",
		p.ANEFraction(), p.SupportsANE(), withSupported, len(p.Operations()))
}
