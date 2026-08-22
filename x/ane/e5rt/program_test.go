//go:build darwin

package e5rt_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/tmc/apple/x/ane/e5rt"
	"github.com/tmc/apple/x/ane/mil"
)

func TestProgramExecuteReusesEncodedOperation(t *testing.T) {
	dir := t.TempDir()
	if err := os.Mkdir(filepath.Join(dir, "weights"), 0o755); err != nil {
		t.Fatal(err)
	}
	modelPath := filepath.Join(dir, "model.mil")
	if err := os.WriteFile(modelPath, []byte(mil.GenConvFP16IO(1, 1, 1)), 0o644); err != nil {
		t.Fatal(err)
	}
	blob, err := mil.BuildWeightBlob([]float32{1}, 1, 1)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "weights", "weight.bin"), blob, 0o644); err != nil {
		t.Fatal(err)
	}

	p, err := e5rt.Compile(e5rt.ProgramOptions{
		ModelPath: modelPath,
		CacheDir:  filepath.Join(dir, "cache"),
		Inputs:    []e5rt.Port{{Name: "x", Size: 2}},
		Outputs:   []e5rt.Port{{Name: "y", Size: 2}},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer p.Close()
	in, err := p.Input("x")
	if err != nil {
		t.Fatal(err)
	}
	out, err := p.Output("y")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []float32{1.5, 2.25} {
		if err := in.WriteFP16([]float32{want}); err != nil {
			t.Fatal(err)
		}
		if err := p.Execute(); err != nil {
			t.Fatal(err)
		}
		var gotValues [1]float32
		if err := out.ReadFP16(gotValues[:]); err != nil {
			t.Fatal(err)
		}
		got := gotValues[0]
		if got != want {
			t.Errorf("output = %v, want %v", got, want)
		}
	}
}

func TestProgramRejectsBadOptions(t *testing.T) {
	for _, test := range []struct {
		name string
		opts e5rt.ProgramOptions
	}{
		{"model", e5rt.ProgramOptions{CacheDir: "cache"}},
		{"cache", e5rt.ProgramOptions{ModelPath: "model"}},
		{"empty port", e5rt.ProgramOptions{ModelPath: "model", CacheDir: "cache", Inputs: []e5rt.Port{{Size: 1}}}},
		{"zero size", e5rt.ProgramOptions{ModelPath: "model", CacheDir: "cache", Inputs: []e5rt.Port{{Name: "x"}}}},
		{"duplicate", e5rt.ProgramOptions{ModelPath: "model", CacheDir: "cache", Inputs: []e5rt.Port{{Name: "x", Size: 1}}, Outputs: []e5rt.Port{{Name: "x", Size: 1}}}},
	} {
		t.Run(test.name, func(t *testing.T) {
			if p, err := e5rt.Compile(test.opts); err == nil {
				if p != nil {
					p.Close()
				}
				t.Fatal("Compile succeeded, want an error")
			}
		})
	}
}
