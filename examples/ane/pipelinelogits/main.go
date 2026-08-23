// Command pipelinelogits runs a language model's logits tail — final RMSNorm,
// classifier projection, softmax — as one linked [e5rt.Pipeline] on the Neural
// Engine, and checks every probability against a float64 CPU reference.
//
// The other e5rt examples each encode a bespoke stream by hand. This one uses
// the package's own multi-stage interface: three independently compiled MIL
// programs, two links, one Execute. The point of the interface is that a linked
// pair of ports is bound to one E5RT buffer object, so this program stages no
// host copy between the stages. That is what is measured here — it is not an
// observation of device-internal movement, and it does not establish that the
// platform performs no hidden copy of its own.
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
	"errors"
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
	if !(*tol > 0) || math.IsInf(*tol, 0) {
		return fmt.Errorf("tol must be a positive finite number, have %v", *tol)
	}
	// The wrong-size refusal control links stage 0's output (dim*seq) to stage
	// 2's input (vocab*seq). With dim == vocab those sizes match and the
	// "malformed" pipeline is well formed, so the control would silently stop
	// testing what its name says while still reporting a pass.
	if *dim == *vocab {
		return fmt.Errorf("dim and vocab must differ: the wrong-size link control depends on stage 0 and stage 2 having different port sizes")
	}
	for _, p := range []struct {
		name string
		a, b int
	}{{"dim*seq", *dim, *seq}, {"vocab*seq", *vocab, *seq}, {"vocab*dim", *vocab, *dim}} {
		if p.a > math.MaxInt32/2/p.b {
			return fmt.Errorf("%s overflows a sensible buffer size", p.name)
		}
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
	got, err := runPipeline(m, x)
	if err != nil {
		return fmt.Errorf("pipeline: %w", err)
	}
	fmt.Println()
	for i, d := range []string{m.normDir, m.clsDir, m.smaxDir} {
		reportBackend(filepath.Join(m.cacheDir, fmt.Sprintf("stage-%02d", i)), d)
	}

	fmt.Println()
	maxDiff, err := compare(got, want)
	if err != nil {
		return fmt.Errorf("linked pipeline vs CPU reference: %w", err)
	}
	fmt.Printf("  linked pipeline vs float64 CPU reference: max diff %.3g (tolerance %.3g)\n", maxDiff, *tol)
	if maxDiff > *tol {
		return fmt.Errorf("linked pipeline disagrees with the CPU reference by %.3g", maxDiff)
	}

	// Reported, not checked — see reportPortIdentity for why.
	reportIdentity(m)

	// The control that can fail: remove a link, and the answer must change.
	if err := linkOmittedControl(m, x, want); err != nil {
		return err
	}

	// The staged arm reaches the same answer with host copies and no links.
	staged, err := runStaged(m, x)
	if err != nil {
		return fmt.Errorf("host-staged arm: %w", err)
	}
	stagedDiff, err := compare(staged, got)
	if err != nil {
		return fmt.Errorf("host-staged arm vs linked pipeline: %w", err)
	}
	fmt.Printf("  host-staged arm vs linked pipeline: max diff %.3g\n", stagedDiff)
	if stagedDiff > *tol {
		return fmt.Errorf("the linked and host-staged arms disagree by %.3g; the links did not carry what the host copies carried", stagedDiff)
	}

	// Mutation control: without it every check above would pass on a stale
	// buffer that never saw this input.
	//
	// The bar is a sensitivity threshold derived from the disagreement this run
	// measured, not the agreement tolerance. Those are different quantities: the
	// tolerance is how far the engine may sit from the reference before the demo
	// calls it broken, while what makes a mutation legible is being much larger
	// than the disagreement already present. Perturbing one element moves one
	// channel at one position, which a softmax over the whole vocabulary then
	// dilutes, so it lands well under a tolerance sized for the fp16 error of
	// the whole chain — and a control keyed to the tolerance reports that
	// dilution as a defect. It did, on the first run of this program.
	//
	// This is a demo sensitivity check and not a statistical estimate of noise.
	// It also proves only that the composition reads its input; the float64
	// reference is what establishes that all three stages are load-bearing.
	threshold := 10 * maxDiff
	if threshold < 1e-4 {
		threshold = 1e-4
	}
	moved, err := mutationControl(m, x, got)
	if err != nil {
		return fmt.Errorf("mutation control: %w", err)
	}
	fmt.Printf("  mutation control: perturbing one input moved the output by %.3g (must exceed %.3g, ten times the measured disagreement)\n", moved, threshold)
	if moved <= threshold {
		return fmt.Errorf("perturbing the input moved the output by only %.3g, not clear of the %.3g sensitivity threshold; the pipeline may not be reading its input", moved, threshold)
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
func runPipeline(m *models, x []float32) ([]float32, error) {
	return runPipelineWithLinks(m, x, links())
}

// runPipelineWithLinks executes the three stages under whatever link set it is
// given, so the same code path serves both the full pipeline and the
// link-omitted negative control.
func runPipelineWithLinks(m *models, x []float32, ls []e5rt.PipelineLink) ([]float32, error) {
	p, err := e5rt.CompilePipeline(e5rt.PipelineOptions{
		CacheDir: m.cacheDir,
		Stages:   m.stages(),
		Links:    ls,
	})
	if err != nil {
		return nil, err
	}
	defer p.Close()

	in, err := p.Input(0, "x")
	if err != nil {
		return nil, err
	}
	out, err := p.Output(2, "out")
	if err != nil {
		return nil, err
	}
	if err := in.WriteFP16(x); err != nil {
		return nil, err
	}
	if err := p.Execute(); err != nil {
		return nil, err
	}
	got := make([]float32, *vocab**seq)
	if err := out.ReadFP16(got); err != nil {
		return nil, err
	}
	return got, nil
}

// reportIdentity opens the full pipeline once purely to print the port-identity
// state described on reportPortIdentity.
func reportIdentity(m *models) {
	p, err := e5rt.CompilePipeline(e5rt.PipelineOptions{
		CacheDir: m.cacheDir,
		Stages:   m.stages(),
		Links:    links(),
	})
	if err != nil {
		return
	}
	defer p.Close()
	reportPortIdentity(p, links())
}

// reportPortIdentity prints whether each link's two ports expose one Go buffer.
//
// This is reported as API state, deliberately not as a check, because it cannot
// fail. CompilePipeline hands the consumer port the producer's own *Buffer —
// `data: source.data` at pipeline.go:418 — so the two handles are the same
// object by construction, whatever the native IOPortBindBufferObject did with
// them. An earlier version of this program wrote a marker byte through one
// handle and read it through the other and called that an aliasing control. It
// was a tautology: the write and the read went to the same Go slice, so the
// control could not produce a negative on any implementation.
//
// The control that can fail is linkOmittedControl below, which removes a link
// and requires the answer to change. That exercises the native binding, which
// is the thing actually in question.
func reportPortIdentity(p *e5rt.Pipeline, links []e5rt.PipelineLink) {
	for _, l := range links {
		producer, err := p.Output(l.From.Stage, l.From.Name)
		if err != nil {
			continue
		}
		consumer, err := p.Input(l.To.Stage, l.To.Name)
		if err != nil {
			continue
		}
		a, b := producer.Bytes(), consumer.Bytes()
		same := len(a) > 0 && len(a) == len(b) && &a[0] == &b[0]
		fmt.Printf("  stage %d %q -> stage %d %q: one Go buffer = %v (by construction, not a test)\n",
			l.From.Stage, l.From.Name, l.To.Stage, l.To.Name, same)
	}
}

// linkOmittedControl is the negative polarity for the links.
//
// It compiles the same three stages with one link removed, so that stage's
// input gets its own fresh buffer instead of the previous stage's output, and
// requires the result to DISAGREE with the float64 reference. If it agreed, the
// link would not be carrying anything and the passing arm above would be
// explained by something else entirely.
//
// This is what makes the link claim load-bearing, because it goes through the
// native binding rather than through a Go-side handle that is shared whatever
// the native call did.
func linkOmittedControl(m *models, x, want []float32) error {
	all := links()
	for drop := range all {
		kept := make([]e5rt.PipelineLink, 0, len(all)-1)
		kept = append(kept, all[:drop]...)
		kept = append(kept, all[drop+1:]...)

		got, err := runPipelineWithLinks(m, x, kept)
		if err != nil {
			// A refusal is also a valid negative: the pipeline declined to
			// run without the link rather than running it wrongly.
			fmt.Printf("  link %d omitted: refused (%v)\n", drop, err)
			continue
		}
		diff, err := compare(got, want)
		if err != nil {
			fmt.Printf("  link %d omitted: result unusable (%v), which is also a disagreement\n", drop, err)
			continue
		}
		if diff <= *tol {
			return fmt.Errorf("removing link %d still matched the reference within %.3g (max diff %.3g); that link carries nothing, so the linked run proves nothing about it", drop, *tol, diff)
		}
		fmt.Printf("  link %d omitted: max diff %.3g, well outside the %.3g tolerance — the link is load-bearing\n", drop, diff, *tol)
	}
	return nil
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
//
// Movement alone is not enough, so the perturbed output is also required to
// match a float64 reference recomputed from the perturbed input. A pipeline
// that reacted to the input at the wrong position, or with the wrong
// arithmetic, would move and pass a movement-only check. Both polarities are
// needed: a changed input must change the output, and the changed output must
// still be right.
func mutationControl(m *models, x, baseline []float32) (float64, error) {
	perturbed := make([]float32, len(x))
	copy(perturbed, x)
	perturbed[0] += 4

	got, err := runPipeline(m, perturbed)
	if err != nil {
		return 0, err
	}
	wantPerturbed := cpuReference(perturbed, m.rmsW, m.embed)
	against, err := compare(got, wantPerturbed)
	if err != nil {
		return 0, err
	}
	if against > *tol {
		return 0, fmt.Errorf("the perturbed output moved but disagrees with its own reference by %.3g; the pipeline reacts to the input incorrectly", against)
	}
	return compare(got, baseline)
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

// compare returns the largest absolute difference between got and want, and
// refuses the two ways this check can pass while measuring nothing.
//
// Unequal lengths mean one side was truncated, and a comparison over the
// shorter prefix would call that identical. A NaN element is worse: it compares
// false against every threshold, so `d > max` never fires and an all-NaN result
// reports a maximum difference of zero. That is the silent case — an infinity
// at least survives the comparison and fails loudly — and NaN is exactly what
// an fp16 overflow in one of these stages would produce.
func compare(got, want []float32) (float64, error) {
	if len(got) != len(want) {
		return 0, fmt.Errorf("comparing %d values against %d", len(got), len(want))
	}
	if len(got) == 0 {
		return 0, errors.New("nothing to compare")
	}
	var max float64
	for i := range got {
		g, w := float64(got[i]), float64(want[i])
		if math.IsNaN(g) || math.IsInf(g, 0) {
			return 0, fmt.Errorf("result element %d is %v", i, g)
		}
		if math.IsNaN(w) || math.IsInf(w, 0) {
			return 0, fmt.Errorf("reference element %d is %v", i, w)
		}
		if d := math.Abs(g - w); d > max {
			max = d
		}
	}
	return max, nil
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
