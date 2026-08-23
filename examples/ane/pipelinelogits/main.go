// Command pipelinelogits runs a language model's logits tail — final RMSNorm,
// classifier projection, softmax — as one linked [e5rt.Pipeline] on the Neural
// Engine, and checks every probability against a float64 CPU reference.
//
// The other e5rt examples each encode a bespoke stream by hand. This one uses
// the package's own multi-stage interface: three independently compiled MIL
// programs, two links, one Execute. The point of the interface is that a
// linked pair of ports shares one buffer object, so an intermediate tensor
// never travels to the host between stages. That claim is what this program
// measures.
//
// # Why a logits tail
//
// It is the part of a decoder that a host round trip hurts most and that no
// other example covers. It is also the smallest interesting chain whose stages
// do not commute: RMSNorm before the projection is not the projection before
// RMSNorm, so a demo that silently ran the stages in the wrong order, or that
// dropped one, could not match the reference by luck.
//
// The three programs come from x/ane/mil's GenFinalRMSNorm,
// GenClassifierForward and GenSoftmaxVocab. Until this program was written
// none of the three had a caller anywhere in the module — they had never been
// compiled, let alone run. Their first execution is here.
//
// # What is checked, and what each check would catch
//
//   - Every output probability is compared against a float64 CPU evaluation of
//     the whole three-stage composition. A stage that ran the wrong arithmetic
//     fails here.
//
//   - The same three programs are run again as three separate [e5rt.Program]s
//     with the host copying between them, and the two arms must agree. This is
//     the control on the links specifically: the staged arm reaches the same
//     answer by a route that uses no links at all, so if the linked arm agrees
//     with it, the links carried what the host copies carried.
//
//   - The linked ports are checked to be one buffer and not two, by writing
//     through stage 0's output and reading the bytes back through stage 1's
//     input. Two buffers that happen to hold equal values would pass a value
//     comparison and fail this.
//
//   - A mutation control perturbs one input element and requires the output to
//     move by far more than the fp16 agreement tolerance. Without it, a program
//     that returned a stale or constant buffer would pass every check above.
//
//   - Two malformed pipelines — a link between ports of different sizes, and a
//     link that runs backwards — are required to be refused. A validator that
//     accepted everything would let the first three checks pass while proving
//     nothing about the shape of the interface.
//
// The compiler's chosen backend is printed for each stage. The device mask is
// a permission and not a placement, so the emitted main_<backend> directory is
// the only local evidence that this ran on the engine at all.
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/tmc/apple/x/ane/e5rt"
	"github.com/tmc/apple/x/ane/mil"
)

var (
	dim   = flag.Int("dim", 64, "model width")
	vocab = flag.Int("vocab", 32, "vocabulary size")
	seq   = flag.Int("seq", 8, "sequence length")
	tol   = flag.Float64("tol", 5e-3, "maximum allowed difference from the CPU reference")
)

func main() {
	log.SetFlags(0)
	flag.Parse()
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	if *dim <= 0 || *vocab <= 0 || *seq <= 0 {
		return fmt.Errorf("dim, vocab and seq must all be positive")
	}
	fmt.Printf("logits tail: dim=%d vocab=%d seq=%d\n", *dim, *vocab, *seq)

	dir, err := os.MkdirTemp("", "pipelinelogits-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)

	m, err := buildModels(dir)
	if err != nil {
		return fmt.Errorf("build models: %w", err)
	}

	x := sampleInput(*dim * *seq)
	want := cpuReference(x, m.rmsW, m.embed)

	// The linked arm: one Execute for the whole tail.
	got, shared, err := runPipeline(m, x)
	if err != nil {
		return fmt.Errorf("pipeline: %w", err)
	}
	fmt.Println()
	for i, d := range []string{m.normDir, m.clsDir, m.smaxDir} {
		reportBackend(filepath.Join(m.cacheDir, fmt.Sprintf("stage-%02d", i)), d)
	}

	fmt.Println()
	maxDiff := compare(got, want)
	fmt.Printf("  linked pipeline vs float64 CPU reference: max diff %.3g (tolerance %.3g)\n", maxDiff, *tol)
	if maxDiff > *tol {
		return fmt.Errorf("linked pipeline disagrees with the CPU reference by %.3g", maxDiff)
	}

	// The links share storage rather than holding equal copies.
	if !shared {
		return fmt.Errorf("stage 0 output and stage 1 input are separate buffers; the link copied instead of aliasing")
	}
	fmt.Println("  linked ports: stage 0 output and stage 1 input are one buffer, not two equal ones")

	// The staged arm reaches the same answer with host copies and no links.
	staged, err := runStaged(m, x)
	if err != nil {
		return fmt.Errorf("host-staged arm: %w", err)
	}
	stagedDiff := compare(staged, got)
	fmt.Printf("  host-staged arm vs linked pipeline: max diff %.3g\n", stagedDiff)
	if stagedDiff > *tol {
		return fmt.Errorf("the linked and host-staged arms disagree by %.3g; the links did not carry what the host copies carried", stagedDiff)
	}

	// Mutation control: without it every check above would pass on a stale
	// buffer that never saw this input.
	//
	// The bar is the noise floor this run actually measured, not the agreement
	// tolerance. Those are different quantities: the tolerance is how far the
	// engine may sit from the reference before the demo calls it broken, while
	// what makes a mutation legible is being much larger than the disagreement
	// already present. Perturbing one element moves one channel at one position,
	// which a softmax over the whole vocabulary then dilutes, so it lands well
	// under a tolerance sized for the fp16 error of the whole chain — and a
	// control keyed to the tolerance would report that dilution as a defect.
	floor := 10 * maxDiff
	if floor < 1e-4 {
		floor = 1e-4
	}
	moved, err := mutationControl(m, x, got)
	if err != nil {
		return fmt.Errorf("mutation control: %w", err)
	}
	fmt.Printf("  mutation control: perturbing one input moved the output by %.3g (must exceed %.3g, ten times the measured disagreement)\n", moved, floor)
	if moved <= floor {
		return fmt.Errorf("perturbing the input moved the output by only %.3g, not clear of the %.3g noise floor; the pipeline may not be reading its input", moved, floor)
	}

	// Refusal controls, so the checks above are known to be discriminating.
	if err := refusalControls(m); err != nil {
		return err
	}

	fmt.Println("\nOK")
	return nil
}

// models holds the three compiled-from-MIL stage directories and the weights
// the CPU reference must use to agree with them.
type models struct {
	cacheDir                 string
	normDir, clsDir, smaxDir string
	norm, cls, smax          string
	rmsW, embed              []float32
}

// buildModels writes the three MIL programs and their weight blobs.
func buildModels(dir string) (*models, error) {
	m := &models{
		cacheDir: filepath.Join(dir, "cache"),
		normDir:  filepath.Join(dir, "norm"),
		clsDir:   filepath.Join(dir, "cls"),
		smaxDir:  filepath.Join(dir, "smax"),
	}
	m.rmsW = sampleWeights(*dim, 1)
	m.embed = sampleWeights(*vocab**dim, 2)

	rmsBlob, err := mil.BuildFP16Blob(m.rmsW)
	if err != nil {
		return nil, err
	}
	embedBlob, err := mil.BuildFP16Blob(m.embed)
	if err != nil {
		return nil, err
	}

	m.norm, err = writeModel(m.normDir, mil.GenFinalRMSNorm(*dim, *seq),
		map[string][]byte{"rms_w.bin": rmsBlob})
	if err != nil {
		return nil, err
	}
	m.cls, err = writeModel(m.clsDir, mil.GenClassifierForward(*dim, *vocab, *seq),
		map[string][]byte{"embed.bin": embedBlob})
	if err != nil {
		return nil, err
	}
	m.smax, err = writeModel(m.smaxDir, mil.GenSoftmaxVocab(*vocab, *seq), nil)
	if err != nil {
		return nil, err
	}
	return m, nil
}

// writeModel lays out one model directory: the MIL text beside a weights
// subdirectory holding the blobs its BLOBFILE references name.
func writeModel(dir, text string, weights map[string][]byte) (string, error) {
	if err := os.MkdirAll(filepath.Join(dir, "weights"), 0o755); err != nil {
		return "", err
	}
	path := filepath.Join(dir, "model.mil")
	if err := os.WriteFile(path, []byte(text), 0o644); err != nil {
		return "", err
	}
	for name, blob := range weights {
		if err := os.WriteFile(filepath.Join(dir, "weights", name), blob, 0o644); err != nil {
			return "", err
		}
	}
	return path, nil
}

// stages describes the three-stage tail. The ports are the names the MIL
// programs declare: every generator in x/ane/mil takes x and returns out.
func (m *models) stages() []e5rt.PipelineStage {
	actBytes := *dim * *seq * 2
	logitBytes := *vocab * *seq * 2
	return []e5rt.PipelineStage{
		{
			ModelPath: m.norm,
			Inputs:    []e5rt.Port{{Name: "x", Size: actBytes}},
			Outputs:   []e5rt.Port{{Name: "out", Size: actBytes}},
		},
		{
			ModelPath: m.cls,
			Inputs:    []e5rt.Port{{Name: "x", Size: actBytes}},
			Outputs:   []e5rt.Port{{Name: "out", Size: logitBytes}},
		},
		{
			ModelPath: m.smax,
			Inputs:    []e5rt.Port{{Name: "x", Size: logitBytes}},
			Outputs:   []e5rt.Port{{Name: "out", Size: logitBytes}},
		},
	}
}

func links() []e5rt.PipelineLink {
	return []e5rt.PipelineLink{
		{From: e5rt.PipelinePort{Stage: 0, Name: "out"}, To: e5rt.PipelinePort{Stage: 1, Name: "x"}},
		{From: e5rt.PipelinePort{Stage: 1, Name: "out"}, To: e5rt.PipelinePort{Stage: 2, Name: "x"}},
	}
}

// runPipeline executes the linked tail once and reports the probabilities and
// whether the first link's two ports are backed by one buffer.
func runPipeline(m *models, x []float32) ([]float32, bool, error) {
	p, err := e5rt.CompilePipeline(e5rt.PipelineOptions{
		CacheDir: m.cacheDir,
		Stages:   m.stages(),
		Links:    links(),
	})
	if err != nil {
		return nil, false, err
	}
	defer p.Close()

	in, err := p.Input(0, "x")
	if err != nil {
		return nil, false, err
	}
	out, err := p.Output(2, "out")
	if err != nil {
		return nil, false, err
	}

	// Aliasing check, before execution so it cannot be confused with a result.
	// Two distinct buffers holding equal values pass a value comparison; only
	// a write seen through the other handle distinguishes them.
	shared, err := portsAlias(p)
	if err != nil {
		return nil, false, err
	}

	if err := in.WriteFP16(x); err != nil {
		return nil, false, err
	}
	if err := p.Execute(); err != nil {
		return nil, false, err
	}
	got := make([]float32, *vocab**seq)
	if err := out.ReadFP16(got); err != nil {
		return nil, false, err
	}
	return got, shared, nil
}

// portsAlias reports whether stage 0's output and stage 1's input are the same
// storage, by writing a byte through one handle and reading it through the
// other. The byte is overwritten by execution, so this must run first.
func portsAlias(p *e5rt.Pipeline) (bool, error) {
	producer, err := p.Output(0, "out")
	if err != nil {
		return false, err
	}
	consumer, err := p.Input(1, "x")
	if err != nil {
		return false, err
	}
	a, b := producer.Bytes(), consumer.Bytes()
	if len(a) == 0 || len(a) != len(b) {
		return false, nil
	}
	const marker = 0xA5
	saved := a[0]
	a[0] = marker
	aliased := b[0] == marker
	a[0] = saved
	return aliased, nil
}

// runStaged runs the same three programs as separate compiled programs with
// the host copying each intermediate, which is the route the pipeline exists
// to avoid. Agreement between the two arms is what makes the links credible.
func runStaged(m *models, x []float32) ([]float32, error) {
	act := make([]float32, *dim**seq)
	copy(act, x)

	norm, err := stageProgram(m, m.norm, "staged-norm", *dim**seq, *dim**seq)
	if err != nil {
		return nil, err
	}
	defer norm.Close()
	if act, err = evalProgram(norm, act, *dim**seq); err != nil {
		return nil, err
	}

	cls, err := stageProgram(m, m.cls, "staged-cls", *dim**seq, *vocab**seq)
	if err != nil {
		return nil, err
	}
	defer cls.Close()
	logits, err := evalProgram(cls, act, *vocab**seq)
	if err != nil {
		return nil, err
	}

	smax, err := stageProgram(m, m.smax, "staged-smax", *vocab**seq, *vocab**seq)
	if err != nil {
		return nil, err
	}
	defer smax.Close()
	return evalProgram(smax, logits, *vocab**seq)
}

func stageProgram(m *models, modelPath, cacheName string, inElems, outElems int) (*e5rt.Program, error) {
	return e5rt.Compile(e5rt.ProgramOptions{
		ModelPath: modelPath,
		CacheDir:  filepath.Join(m.cacheDir, cacheName),
		Inputs:    []e5rt.Port{{Name: "x", Size: inElems * 2}},
		Outputs:   []e5rt.Port{{Name: "out", Size: outElems * 2}},
	})
}

func evalProgram(p *e5rt.Program, in []float32, outElems int) ([]float32, error) {
	inBuf, err := p.Input("x")
	if err != nil {
		return nil, err
	}
	outBuf, err := p.Output("out")
	if err != nil {
		return nil, err
	}
	if err := inBuf.WriteFP16(in); err != nil {
		return nil, err
	}
	if err := p.Execute(); err != nil {
		return nil, err
	}
	out := make([]float32, outElems)
	if err := outBuf.ReadFP16(out); err != nil {
		return nil, err
	}
	return out, nil
}

// mutationControl reruns the pipeline with one input element changed and
// reports how far the output moved. A pipeline that ignored its input, or
// returned a buffer left over from the previous run, would report zero.
func mutationControl(m *models, x, baseline []float32) (float64, error) {
	perturbed := make([]float32, len(x))
	copy(perturbed, x)
	perturbed[0] += 4
	got, _, err := runPipeline(m, perturbed)
	if err != nil {
		return 0, err
	}
	return compare(got, baseline), nil
}

// refusalControls requires the validator to reject two malformed pipelines.
// Without them the successful runs above would say nothing about whether the
// interface checks anything at all.
func refusalControls(m *models) error {
	cases := []struct {
		name  string
		links []e5rt.PipelineLink
	}{
		{
			// Stage 0 emits dim*seq values; stage 2 wants vocab*seq.
			name: "link between ports of different sizes",
			links: []e5rt.PipelineLink{
				{From: e5rt.PipelinePort{Stage: 0, Name: "out"}, To: e5rt.PipelinePort{Stage: 2, Name: "x"}},
			},
		},
		{
			name: "link that runs backwards",
			links: []e5rt.PipelineLink{
				{From: e5rt.PipelinePort{Stage: 2, Name: "out"}, To: e5rt.PipelinePort{Stage: 1, Name: "x"}},
			},
		},
	}
	fmt.Println()
	for _, c := range cases {
		p, err := e5rt.CompilePipeline(e5rt.PipelineOptions{
			CacheDir: m.cacheDir,
			Stages:   m.stages(),
			Links:    c.links,
		})
		if err == nil {
			p.Close()
			return fmt.Errorf("refusal control: a pipeline with a %s was accepted", c.name)
		}
		fmt.Printf("  refused, as it must be: %s\n    %v\n", c.name, err)
	}
	return nil
}

// cpuReference evaluates the whole tail in float64: RMSNorm with baked
// weights, then the classifier projection, then a softmax over the vocabulary
// axis. Tensors are [1, C, 1, S] row-major, so element (c, s) is at c*seq+s.
func cpuReference(x, rmsW, embed []float32) []float32 {
	d, v, s := *dim, *vocab, *seq

	normed := make([]float64, d*s)
	for j := 0; j < s; j++ {
		var sum float64
		for c := 0; c < d; c++ {
			val := float64(x[c*s+j])
			sum += val * val
		}
		rrms := 1 / math.Sqrt(sum/float64(d)+1e-5)
		for c := 0; c < d; c++ {
			normed[c*s+j] = float64(x[c*s+j]) * rrms * float64(rmsW[c])
		}
	}

	logits := make([]float64, v*s)
	for o := 0; o < v; o++ {
		for j := 0; j < s; j++ {
			var acc float64
			for c := 0; c < d; c++ {
				acc += float64(embed[o*d+c]) * normed[c*s+j]
			}
			logits[o*s+j] = acc
		}
	}

	out := make([]float32, v*s)
	for j := 0; j < s; j++ {
		max := math.Inf(-1)
		for o := 0; o < v; o++ {
			if logits[o*s+j] > max {
				max = logits[o*s+j]
			}
		}
		var sum float64
		for o := 0; o < v; o++ {
			sum += math.Exp(logits[o*s+j] - max)
		}
		for o := 0; o < v; o++ {
			out[o*s+j] = float32(math.Exp(logits[o*s+j]-max) / sum)
		}
	}
	return out
}

func compare(got, want []float32) float64 {
	var max float64
	for i := range got {
		if d := math.Abs(float64(got[i]) - float64(want[i])); d > max {
			max = d
		}
	}
	return max
}

// sampleInput and sampleWeights produce deterministic values, so the CPU
// reference and the engine see exactly the same numbers on every run. The
// scales keep every intermediate well inside fp16 range: an overflow to
// infinity would show up as a reference disagreement and be misread as a
// binding defect.
func sampleInput(n int) []float32 {
	out := make([]float32, n)
	for i := range out {
		out[i] = float32(math.Sin(float64(i)*0.7+0.3)) * 2
	}
	return out
}

func sampleWeights(n, salt int) []float32 {
	out := make([]float32, n)
	for i := range out {
		out[i] = float32(math.Cos(float64(i)*0.37+float64(salt))) * 0.25
	}
	return out
}

// reportBackend prints which backend the compiler chose for one stage, read
// from the main_<backend> directory it leaves in the cache location. The
// device mask is a permission and not a placement, so this is the only local
// evidence that the work was put on the engine.
func reportBackend(cacheDir, label string) {
	seen := map[string]bool{}
	filepath.WalkDir(cacheDir, func(path string, d os.DirEntry, err error) error {
		if err != nil || !d.IsDir() || filepath.Base(filepath.Dir(path)) != "main" {
			return nil
		}
		if suffix, ok := strings.CutPrefix(d.Name(), "main_"); ok {
			seen[suffix] = true
		}
		return nil
	})
	name := filepath.Base(label)
	if len(seen) == 0 {
		fmt.Printf("  backend for %s: UNKNOWN (no main_* directory in the compiled bundle)\n", name)
		return
	}
	names := make([]string, 0, len(seen))
	for n := range seen {
		names = append(names, n)
	}
	fmt.Printf("  backend for %s: the compiler emitted %v\n", name, names)
}
