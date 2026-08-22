//go:build darwin

package e5rt_test

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tmc/apple/x/ane/e5rt"
	"github.com/tmc/apple/x/ane/mil"
)

func TestProgramExecuteReusesEncodedOperation(t *testing.T) {
	dir := t.TempDir()
	modelPath, err := writeIdentityModel(dir)
	if err != nil {
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

func ExampleCompile() {
	dir, err := os.MkdirTemp("", "e5rt-program-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)
	modelPath, err := writeIdentityModel(dir)
	if err != nil {
		log.Fatal(err)
	}
	p, err := e5rt.Compile(e5rt.ProgramOptions{
		ModelPath: modelPath,
		CacheDir:  filepath.Join(dir, "cache"),
		Inputs:    []e5rt.Port{{Name: "x", Size: 2}},
		Outputs:   []e5rt.Port{{Name: "y", Size: 2}},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()
	in, err := p.Input("x")
	if err != nil {
		log.Fatal(err)
	}
	out, err := p.Output("y")
	if err != nil {
		log.Fatal(err)
	}
	if err := in.WriteFP16([]float32{1.5}); err != nil {
		log.Fatal(err)
	}
	if err := p.Execute(); err != nil {
		log.Fatal(err)
	}
	values := make([]float32, 1)
	if err := out.ReadFP16(values); err != nil {
		log.Fatal(err)
	}
	fmt.Println(values)

	// Output:
	// [1.5]
}

func TestOpenBundle(t *testing.T) {
	dir := t.TempDir()
	modelPath, err := writeIdentityModel(dir)
	if err != nil {
		t.Fatal(err)
	}
	cacheDir := filepath.Join(dir, "cache")
	p, err := e5rt.Compile(e5rt.ProgramOptions{
		ModelPath:          modelPath,
		CacheDir:           cacheDir,
		ForceRecompilation: true,
		Inputs:             []e5rt.Port{{Name: "x", Size: 2}},
		Outputs:            []e5rt.Port{{Name: "y", Size: 2}},
	})
	if err != nil {
		t.Fatal(err)
	}
	if err := p.Close(); err != nil {
		t.Fatal(err)
	}
	bundlePath, err := compiledBundle(cacheDir)
	if err != nil {
		t.Fatal(err)
	}
	p, err = e5rt.OpenBundle(e5rt.BundleOptions{
		BundlePath: bundlePath,
		Inputs:     []e5rt.Port{{Name: "x", Size: 2}},
		Outputs:    []e5rt.Port{{Name: "y", Size: 2}},
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
	if err := in.WriteFP16([]float32{3.5}); err != nil {
		t.Fatal(err)
	}
	if err := p.Execute(); err != nil {
		t.Fatal(err)
	}
	got := make([]float32, 1)
	if err := out.ReadFP16(got); err != nil {
		t.Fatal(err)
	}
	if got[0] != 3.5 {
		t.Errorf("output = %v, want 3.5", got[0])
	}
}

func TestPipelineLinksStages(t *testing.T) {
	dir := t.TempDir()
	first, err := writeIdentityModel(filepath.Join(dir, "first"))
	if err != nil {
		t.Fatal(err)
	}
	second, err := writeIdentityModel(filepath.Join(dir, "second"))
	if err != nil {
		t.Fatal(err)
	}
	p, err := e5rt.CompilePipeline(e5rt.PipelineOptions{
		CacheDir: filepath.Join(dir, "cache"),
		Stages: []e5rt.PipelineStage{
			{ModelPath: first, Inputs: []e5rt.Port{{Name: "x", Size: 2}}, Outputs: []e5rt.Port{{Name: "y", Size: 2}}},
			{ModelPath: second, Inputs: []e5rt.Port{{Name: "x", Size: 2}}, Outputs: []e5rt.Port{{Name: "y", Size: 2}}},
		},
		Links: []e5rt.PipelineLink{{
			From: e5rt.PipelinePort{Stage: 0, Name: "y"},
			To:   e5rt.PipelinePort{Stage: 1, Name: "x"},
		}},
	})
	if err != nil {
		t.Fatal(err)
	}
	defer p.Close()
	in, err := p.Input(0, "x")
	if err != nil {
		t.Fatal(err)
	}
	intermediate, err := p.Output(0, "y")
	if err != nil {
		t.Fatal(err)
	}
	linked, err := p.Input(1, "x")
	if err != nil {
		t.Fatal(err)
	}
	if &intermediate.Bytes()[0] != &linked.Bytes()[0] {
		t.Fatal("linked ports do not share storage")
	}
	out, err := p.Output(1, "y")
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []float32{1.5, 2.5} {
		if err := in.WriteFP16([]float32{want}); err != nil {
			t.Fatal(err)
		}
		if err := p.Execute(); err != nil {
			t.Fatal(err)
		}
		got := make([]float32, 1)
		if err := out.ReadFP16(got); err != nil {
			t.Fatal(err)
		}
		if got[0] != want {
			t.Errorf("output = %v, want %v", got[0], want)
		}
	}
}

func ExampleCompilePipeline() {
	dir, err := os.MkdirTemp("", "e5rt-pipeline-")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(dir)
	first, err := writeIdentityModel(filepath.Join(dir, "first"))
	if err != nil {
		log.Fatal(err)
	}
	second, err := writeIdentityModel(filepath.Join(dir, "second"))
	if err != nil {
		log.Fatal(err)
	}
	p, err := e5rt.CompilePipeline(e5rt.PipelineOptions{
		CacheDir: filepath.Join(dir, "cache"),
		Stages: []e5rt.PipelineStage{
			{ModelPath: first, Inputs: []e5rt.Port{{Name: "x", Size: 2}}, Outputs: []e5rt.Port{{Name: "y", Size: 2}}},
			{ModelPath: second, Inputs: []e5rt.Port{{Name: "x", Size: 2}}, Outputs: []e5rt.Port{{Name: "y", Size: 2}}},
		},
		Links: []e5rt.PipelineLink{{
			From: e5rt.PipelinePort{Stage: 0, Name: "y"},
			To:   e5rt.PipelinePort{Stage: 1, Name: "x"},
		}},
	})
	if err != nil {
		log.Fatal(err)
	}
	defer p.Close()
	in, err := p.Input(0, "x")
	if err != nil {
		log.Fatal(err)
	}
	out, err := p.Output(1, "y")
	if err != nil {
		log.Fatal(err)
	}
	if err := in.WriteFP16([]float32{1.5}); err != nil {
		log.Fatal(err)
	}
	if err := p.Execute(); err != nil {
		log.Fatal(err)
	}
	values := make([]float32, 1)
	if err := out.ReadFP16(values); err != nil {
		log.Fatal(err)
	}
	fmt.Println(values)

	// Output:
	// [1.5]
}

func TestPipelineRejectsBadOptions(t *testing.T) {
	stage := func(inputSize, outputSize int) e5rt.PipelineStage {
		return e5rt.PipelineStage{
			ModelPath: "model",
			Inputs:    []e5rt.Port{{Name: "x", Size: inputSize}},
			Outputs:   []e5rt.Port{{Name: "y", Size: outputSize}},
		}
	}
	tooMany := make([]e5rt.PipelineStage, e5rt.MaxPipelineStages+1)
	for i := range tooMany {
		tooMany[i] = stage(2, 2)
	}
	for _, test := range []struct {
		name string
		opts e5rt.PipelineOptions
	}{
		{"empty", e5rt.PipelineOptions{CacheDir: "cache"}},
		{"too many", e5rt.PipelineOptions{CacheDir: "cache", Stages: tooMany}},
		{"backward", e5rt.PipelineOptions{CacheDir: "cache", Stages: []e5rt.PipelineStage{stage(2, 2), stage(2, 2)}, Links: []e5rt.PipelineLink{{From: e5rt.PipelinePort{Stage: 1, Name: "y"}, To: e5rt.PipelinePort{Stage: 0, Name: "x"}}}}},
		{"mismatched", e5rt.PipelineOptions{CacheDir: "cache", Stages: []e5rt.PipelineStage{stage(2, 2), stage(4, 2)}, Links: []e5rt.PipelineLink{{From: e5rt.PipelinePort{Stage: 0, Name: "y"}, To: e5rt.PipelinePort{Stage: 1, Name: "x"}}}}},
	} {
		t.Run(test.name, func(t *testing.T) {
			if p, err := e5rt.CompilePipeline(test.opts); err == nil {
				if p != nil {
					p.Close()
				}
				t.Fatal("CompilePipeline succeeded, want an error")
			}
		})
	}
}

func writeIdentityModel(dir string) (string, error) {
	if err := os.MkdirAll(filepath.Join(dir, "weights"), 0o755); err != nil {
		return "", err
	}
	modelPath := filepath.Join(dir, "model.mil")
	if err := os.WriteFile(modelPath, []byte(mil.GenConvFP16IO(1, 1, 1)), 0o644); err != nil {
		return "", err
	}
	blob, err := mil.BuildWeightBlob([]float32{1}, 1, 1)
	if err != nil {
		return "", err
	}
	if err := os.WriteFile(filepath.Join(dir, "weights", "weight.bin"), blob, 0o644); err != nil {
		return "", err
	}
	return modelPath, nil
}

func compiledBundle(cacheDir string) (string, error) {
	var bundle string
	err := filepath.WalkDir(cacheDir, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if entry.IsDir() && strings.HasSuffix(entry.Name(), ".bundle") {
			bundle = path
			return filepath.SkipDir
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	if bundle == "" {
		return "", fmt.Errorf("no program bundle under %s", cacheDir)
	}
	return bundle, nil
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
