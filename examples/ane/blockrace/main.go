// Command blockrace runs the same feed forward network on the Neural Engine and
// on the GPU, first one at a time and then both at once, to measure whether the
// two engines contend when one process drives them together.
//
// The network is the one x/ane/mil generates for a transformer block's feed
// forward half: out = W2 * relu(W1 * x)^2. The Neural Engine runs it through the
// private e5rt route from a compiled MIL program; the GPU runs the same
// arithmetic as an MPSGraph. Both are checked against the same float64 CPU
// reference before any timing is reported, so an arm cannot win by being fast
// and wrong.
//
// # What is being compared, and what is not
//
// The engines run at different precisions. Everything on the Neural Engine is
// fp16, and the MPSGraph arm here is fp32, so their absolute rates are not
// comparable and this program does not present them as though they were. What
// is comparable is what each arm does to itself: the number reported for each
// engine is its own concurrent rate against its own solo rate, which is a ratio
// of like to like and is what answers the contention question.
//
// The aggregate line is the median over rounds of the two concurrent rates
// summed within each round, against the faster solo rate. Above 1.0 means the
// two engines together did more work per second than the better of them alone —
// that they add capacity rather than divide it. It is a throughput result for
// this machine, this block shape, and these two precisions; it is not a claim
// that the engines do not contend for anything, and the per-engine lines above
// it, which do show a cost, are the evidence on that question.
//
// # Placement
//
// The device mask handed to the compiler is a permission and not a placement, so
// the compiled bundle is read for which backend the compiler actually chose. A
// Neural Engine arm that quietly fell back to the CPU would return the same
// correct answer and would contend with everything.
//
//	go run ./examples/ane/blockrace
//	go run ./examples/ane/blockrace -dim 256 -hidden 1024 -seq 64 -window 3s
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/tmc/apple/x/ane"
	"github.com/tmc/apple/x/ane/e5rt"
	"github.com/tmc/apple/x/ane/mil"
)

func init() {
	// Metal and the ObjC runtime want a stable thread for the main goroutine.
	runtime.LockOSThread()
}

func main() {
	log.SetFlags(0)
	dim := flag.Int("dim", 128, "model width")
	hidden := flag.Int("hidden", 512, "feed forward width")
	seq := flag.Int("seq", 32, "sequence length")
	window := flag.Duration("window", 1500*time.Millisecond, "measurement window per arm; shorter windows on a busy machine fail the stability check")
	rounds := flag.Int("rounds", 5, "alternating solo/concurrent rounds to take the median of")
	flag.Parse()

	if err := run(*dim, *hidden, *seq, *window, *rounds); err != nil {
		log.Fatal(err)
	}
}

// arm is one engine's runnable, verified network.
type arm struct {
	name string
	// once runs the network once over the already-written input.
	once func() error
	// read returns the most recent output in [dim][seq] layout.
	read func() []float32
	// write stores an input in [dim][seq] layout.
	write func([]float32)
}

func run(dim, hidden, seq int, window time.Duration, rounds int) error {
	// Scales chosen so the result lands near unit magnitude. A projection sums
	// over its input width, so a weight of scale s gives roughly s*sqrt(width);
	// the squaring in the middle shrinks things further, and weights that leave
	// the output near zero would put the comparison below entirely inside its
	// own absolute tolerance floor, where a badly wrong answer still passes.
	w1 := randomWeights(hidden*dim, 2/math.Sqrt(float64(dim)), 77)
	w2 := randomWeights(dim*hidden, 2/math.Sqrt(float64(hidden)), 88)

	input := make([]float32, dim*seq)
	for i := range input {
		input[i] = float32(0.5 * math.Sin(float64(i)*0.21))
	}
	want := ffnReference(input, w1, w2, dim, hidden, seq)

	// The comparison has an absolute floor, so a network whose output sits near
	// zero would be checked against a tolerance larger than its own signal.
	var peak float64
	for _, v := range want {
		peak = math.Max(peak, math.Abs(v))
	}
	if peak < 0.25 {
		return fmt.Errorf("the reference peaks at %.4f, which is too small for the comparison tolerance to mean anything", peak)
	}

	fmt.Printf("feed forward: dim %d, hidden %d, seq %d — out = W2*relu(W1*x)^2\n", dim, hidden, seq)
	flops := 2 * float64(seq) * float64(dim) * float64(hidden) * 2 // two matmuls

	aneArm, cleanup, err := newANEArm(dim, hidden, seq, w1, w2)
	if err != nil {
		return fmt.Errorf("neural engine arm: %w", err)
	}
	defer cleanup()

	gpuArm, err := newGPUArm(dim, hidden, seq, w1, w2)
	if err != nil {
		return fmt.Errorf("gpu arm: %w", err)
	}

	// Verify both arms before timing anything.
	for _, a := range []*arm{aneArm, gpuArm} {
		a.write(input)
		if err := a.once(); err != nil {
			return fmt.Errorf("%s: %w", a.name, err)
		}
		if err := compare(a.name, a.read(), want); err != nil {
			return err
		}
	}

	// Warm up before measuring anything. Both engines ramp their clocks under
	// sustained load, and a cold measurement reads slow.
	fmt.Printf("\nwarming up for %v per arm\n", window)
	measure(aneArm, window)
	measure(gpuArm, window)

	// Alternate solo and concurrent rounds rather than measuring all of one and
	// then all of the other.
	//
	// This machine drifts. An earlier version measured both solo rates, then
	// both concurrent rates, and reported each engine running *faster* while
	// contending — which is not a contention result, it is the clocks still
	// coming up during the solo phase. Ordering artifacts of that kind survive
	// a warm-up: what removes them is interleaving, so that any drift lands on
	// both phases alike, and taking the median of several rounds.
	fmt.Printf("measuring %d rounds of solo and concurrent, %v per arm per round\n", rounds, window)
	var soloA, soloG, concA, concG []float64
	for range rounds {
		soloA = append(soloA, measure(aneArm, window))
		soloG = append(soloG, measure(gpuArm, window))
		a, g := measureConcurrent(aneArm, gpuArm, window)
		concA = append(concA, a)
		concG = append(concG, g)
	}
	soloANE, spreadA := median(soloA), spread(soloA)
	soloGPU, spreadG := median(soloG), spread(soloG)
	concANE, concGPU := median(concA), median(concG)

	fmt.Printf("\nsolo, median of %d rounds\n", rounds)
	fmt.Printf("  %-14s %8.1f runs/s  %8.1f GFLOPS  (spread %.0f%% across rounds)\n",
		aneArm.name, soloANE, soloANE*flops/1e9, 100*spreadA)
	fmt.Printf("  %-14s %8.1f runs/s  %8.1f GFLOPS  (spread %.0f%% across rounds)\n",
		gpuArm.name, soloGPU, soloGPU*flops/1e9, 100*spreadG)
	if spreadA > 0.25 || spreadG > 0.25 {
		return fmt.Errorf("the solo rates varied by more than a quarter across rounds (%.0f%% and %.0f%%), so nothing can be concluded from comparing them against the concurrent rates; try a longer -window on a quieter machine",
			100*spreadA, 100*spreadG)
	}

	fmt.Printf("\nconcurrent, both engines at once, median of %d rounds\n", rounds)
	fmt.Printf("  %-14s %8.1f runs/s  %+6.1f%% against its own solo rate\n",
		aneArm.name, concANE, 100*(concANE/soloANE-1))
	fmt.Printf("  %-14s %8.1f runs/s  %+6.1f%% against its own solo rate\n",
		gpuArm.name, concGPU, 100*(concGPU/soloGPU-1))

	// Both arms must still be correct after the concurrent run, or the rates
	// above are rates of producing garbage.
	for _, a := range []*arm{aneArm, gpuArm} {
		if err := compare(a.name+" after the concurrent run", a.read(), want); err != nil {
			return err
		}
	}

	// Aggregate from the paired rounds, not from the two medians.
	//
	// concANE and concGPU are each a median over rounds, and the round that
	// supplies one is not generally the round that supplies the other, so
	// adding them totals two rates that were never observed together. Summing
	// within each round first and taking the median of those sums keeps the
	// aggregate a number the machine actually produced.
	combined := make([]float64, len(concA))
	for i := range concA {
		combined[i] = concA[i] + concG[i]
	}
	best := math.Max(soloANE, soloGPU)
	total := median(combined)
	fmt.Printf("\n  together %.1f runs/s against %.1f for the faster engine alone: %.2fx\n",
		total, best, total/best)
	if total > best {
		fmt.Printf("  the engines add capacity rather than divide it\n")
	} else {
		fmt.Printf("  running both at once did not beat the faster engine alone\n")
	}
	return nil
}

// measure runs an arm in a loop for the window and returns runs per second.
//
// A fixed window rather than a fixed count, so that both arms are under load
// for the same span during the concurrent phase; a count-based loop would let
// the faster arm finish early and leave the slower one running alone.
func measure(a *arm, window time.Duration) float64 {
	start := time.Now()
	deadline := start.Add(window)
	runs := 0
	for time.Now().Before(deadline) {
		if err := a.once(); err != nil {
			log.Fatalf("%s: %v", a.name, err)
		}
		runs++
	}
	return float64(runs) / time.Since(start).Seconds()
}

// ffnReference recomputes out = W2*relu(W1*x)^2 in float64, in the [dim][seq]
// layout the engine uses.
func ffnReference(x []float32, w1, w2 []float32, dim, hidden, seq int) []float64 {
	in := make([]float64, len(x))
	for i, v := range x {
		in[i] = float64(v)
	}
	h := project(in, w1, hidden, dim, seq)
	for i, v := range h {
		if v < 0 {
			h[i] = 0
		} else {
			h[i] = v * v
		}
	}
	return project(h, w2, dim, hidden, seq)
}

// project applies an OIHW 1x1 convolution: out[o][s] = sum_i w[o][i] * x[i][s].
func project(x []float64, weight []float32, outCh, inCh, seq int) []float64 {
	out := make([]float64, outCh*seq)
	for o := range outCh {
		row := weight[o*inCh : (o+1)*inCh]
		for s := range seq {
			var acc float64
			for i := range inCh {
				acc += float64(row[i]) * x[i*seq+s]
			}
			out[o*seq+s] = acc
		}
	}
	return out
}

// randomWeights returns n deterministic pseudo-random values in [-spread,
// spread), from a xorshift generator seeded by seed. Decorrelated values matter
// here for the same reason they do in the other examples: correlated weights
// cancel under a long dot product instead of accumulating.
func randomWeights(n int, spread float64, seed uint64) []float32 {
	state := seed*2862933555777941757 + 3037000493
	next := func() float64 {
		state ^= state << 13
		state ^= state >> 7
		state ^= state << 17
		return float64(state>>11)/(1<<52) - 1
	}
	out := make([]float32, n)
	for i := range out {
		out[i] = float32(next() * spread)
	}
	return out
}

// compare reports the worst difference against a tolerance scaled to the
// reference's own magnitude, and prints the peak alongside so a small
// difference against a small reference is not read as agreement.
//
// The tolerance is generous enough for the fp16 arm; the fp32 arm lands far
// inside it, and both print what they actually achieved.
func compare(label string, got []float32, want []float64) error {
	if len(got) != len(want) {
		return fmt.Errorf("%s: %d values, want %d", label, len(got), len(want))
	}
	var worst, peak float64
	for i := range got {
		g := float64(got[i])
		if math.IsNaN(g) || math.IsInf(g, 0) {
			return fmt.Errorf("%s: element %d is %v", label, i, got[i])
		}
		worst = math.Max(worst, math.Abs(g-want[i]))
		peak = math.Max(peak, math.Abs(want[i]))
	}
	tol := 0.02*peak + 0.01
	if worst > tol {
		return fmt.Errorf("%s: worst |diff| %.4f exceeds tolerance %.4f (reference peak |%.4f|)", label, worst, tol, peak)
	}
	fmt.Printf("  %-14s matches the float64 reference: worst |diff| %.5f, tolerance %.4f, reference peak |%.4f|\n",
		label, worst, tol, peak)
	return nil
}

// newANEArm compiles the feed forward network for the Neural Engine and returns
// an arm that dispatches it through e5rt.
func newANEArm(dim, hidden, seq int, w1, w2 []float32) (*arm, func(), error) {
	lib, err := e5rt.Open()
	if err != nil {
		return nil, nil, err
	}
	w1Blob, err := mil.BuildWeightBlob(w1, hidden, dim)
	if err != nil {
		return nil, nil, err
	}
	w2Blob, err := mil.BuildWeightBlob(w2, dim, hidden)
	if err != nil {
		return nil, nil, err
	}
	dir, err := os.MkdirTemp("", "blockrace")
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() { os.RemoveAll(dir) }
	if err := os.MkdirAll(filepath.Join(dir, "weights"), 0o755); err != nil {
		return nil, cleanup, err
	}
	if err := os.WriteFile(filepath.Join(dir, "model.mil"), []byte(mil.GenFFNForwardReLU2(dim, hidden, seq)), 0o644); err != nil {
		return nil, cleanup, err
	}
	for name, blob := range map[string][]byte{"w1.bin": w1Blob, "w2.bin": w2Blob} {
		if err := os.WriteFile(filepath.Join(dir, "weights", name), blob, 0o644); err != nil {
			return nil, cleanup, err
		}
	}

	config, err := lib.CompilerConfigOptionsCreate()
	if err != nil {
		return nil, cleanup, err
	}
	defer lib.CompilerConfigOptionsRelease(config)
	if err := lib.CompilerConfigOptionsSetCacheBundleLocation(config, dir); err != nil {
		return nil, cleanup, err
	}
	compiler, err := lib.CompilerCreateWithConfig(config)
	if err != nil {
		return nil, cleanup, err
	}
	defer lib.CompilerRelease(compiler)
	options, err := lib.CompilerOptionsCreate()
	if err != nil {
		return nil, cleanup, err
	}
	defer lib.CompilerOptionsRelease(options)
	if err := lib.CompilerOptionsSetComputeDeviceTypesMask(options, e5rt.ComputeDeviceANE); err != nil {
		return nil, cleanup, err
	}
	library, err := lib.CompilerCompile(compiler, filepath.Join(dir, "model.mil"), options)
	if err != nil {
		return nil, cleanup, err
	}
	function, err := lib.ProgramLibraryRetainProgramFunction(library, "main")
	if err != nil {
		return nil, cleanup, err
	}
	reportBackend(dir)

	opOptions, err := lib.PrecompiledComputeOpOptionsCreate(function)
	if err != nil {
		return nil, cleanup, err
	}
	defer lib.PrecompiledComputeOpOptionsRelease(opOptions)
	if err := lib.PrecompiledComputeOpOptionsSetOperationName(opOptions, "main"); err != nil {
		return nil, cleanup, err
	}
	op, err := lib.OperationCreatePrecompiled(opOptions)
	if err != nil {
		return nil, cleanup, err
	}

	acts := dim * seq
	inBuf, inPtr, err := allocBuffer(lib, acts*2)
	if err != nil {
		return nil, cleanup, err
	}
	outBuf, outPtr, err := allocBuffer(lib, acts*2)
	if err != nil {
		return nil, cleanup, err
	}
	if err := bind(lib, op, "x", inBuf, true); err != nil {
		return nil, cleanup, err
	}
	if err := bind(lib, op, "out", outBuf, false); err != nil {
		return nil, cleanup, err
	}
	stream, err := lib.ExecutionStreamCreate()
	if err != nil {
		return nil, cleanup, err
	}
	if err := lib.EncodeOperation(stream, op); err != nil {
		return nil, cleanup, err
	}

	return &arm{
		name: "neural engine",
		once: func() error { return lib.ExecuteSync(stream) },
		read: func() []float32 {
			in := fp16SliceAt(outPtr, acts)
			out := make([]float32, acts)
			for i, v := range in {
				out[i] = ane.FP16ToFloat32(v)
			}
			return out
		},
		write: func(data []float32) {
			dst := fp16SliceAt(inPtr, len(data))
			for i, v := range data {
				dst[i] = ane.Float32ToFP16(v)
			}
		},
	}, cleanup, nil
}

// allocBuffer allocates a CPU-visible buffer object and returns it with its
// host address.
func allocBuffer(lib *e5rt.Lib, nbytes int) (uintptr, uintptr, error) {
	buf, err := lib.BufferObjectAlloc(uintptr(max((nbytes+63)&^63, 64)), 0)
	if err != nil {
		return 0, 0, err
	}
	ptr, err := lib.BufferObjectGetDataPtr(buf)
	if err != nil {
		return 0, 0, err
	}
	return buf, ptr, nil
}

// bind binds a buffer object to a named port.
func bind(lib *e5rt.Lib, op uintptr, name string, buf uintptr, input bool) error {
	var port uintptr
	var err error
	if input {
		port, err = lib.OperationRetainInputPort(op, name)
	} else {
		port, err = lib.OperationRetainOutputPort(op, name)
	}
	if err != nil {
		return err
	}
	defer lib.IOPortRelease(port)
	return lib.IOPortBindBufferObject(port, buf)
}

// reportBackend prints which backend the compiler chose, read from the
// <function>_<backend> directories it leaves in the cache location.
func reportBackend(cacheDir string) {
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
	if len(seen) == 0 {
		fmt.Printf("  backend: UNKNOWN (no main_* directory in the compiled bundle)\n")
		return
	}
	names := make([]string, 0, len(seen))
	for name := range seen {
		names = append(names, name)
	}
	fmt.Printf("  backend: the compiler emitted %v for the neural engine arm\n", names)
}

// measureConcurrent runs both arms at once for the window and returns each
// one's rate.
func measureConcurrent(a, b *arm, window time.Duration) (float64, float64) {
	var wg sync.WaitGroup
	var ra, rb float64
	wg.Add(2)
	go func() {
		defer wg.Done()
		runtime.LockOSThread()
		ra = measure(a, window)
	}()
	go func() {
		defer wg.Done()
		runtime.LockOSThread()
		rb = measure(b, window)
	}()
	wg.Wait()
	return ra, rb
}

// median returns the middle value of vals.
func median(vals []float64) float64 {
	s := append([]float64(nil), vals...)
	slices.Sort(s)
	return s[len(s)/2]
}

// spread returns the interquartile range of vals as a fraction of the median,
// which is how this program decides whether the machine held still enough for a
// comparison of rates to mean anything.
//
// Interquartile rather than the full range, because a single round disturbed by
// something else on the machine otherwise decides the answer: with a handful of
// rounds, min-to-max is the outlier. This discards the extreme rounds and asks
// whether the rest agree. It is not a way of passing a check that the full
// range would fail — a genuinely unsettled machine spreads every round, not
// one.
func spread(vals []float64) float64 {
	s := append([]float64(nil), vals...)
	slices.Sort(s)
	m := s[len(s)/2]
	if m == 0 {
		return math.Inf(1)
	}
	lo, hi := len(s)/4, len(s)-1-len(s)/4
	return (s[hi] - s[lo]) / m
}
