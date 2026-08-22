// Command e5rtblock runs a transformer block on the Neural Engine through the
// private e5rt direct-dispatch route, and checks what comes back against a
// float64 CPU reference of the whole block.
//
// The other e5rt example, e5rtdispatch, proves the route runs by dispatching a
// 1x1 convolution. This one runs work with the shape of a real model: RMSNorm,
// a QKV projection, multi-head attention with a causal mask, an output
// projection and a residual, then a second RMSNorm and a ReLU-squared feed
// forward network. The MIL comes from x/ane/mil's GenSDPAForward and
// GenFFNForwardRMSReLU2, which until now had only ever been checked for
// compiling — nothing verified that either computes what its name says.
//
// # The two programs share a buffer
//
// Attention and the feed forward network are separate compiled programs, and
// they are chained without a host round trip: the buffer object bound to the
// attention program's output port is bound again to the feed forward program's
// input port, so the second reads the first's result in place. Both operations
// are encoded into one execution stream and dispatched by one execute call.
// Nothing copies the intermediate activation, and the host never sees it.
//
// # The residual the generator does not add
//
// GenFFNForwardRMSReLU2 computes W2(relu(W1(rms_norm(x)))^2) and stops there.
// It has no residual add, so the block output is the attention result plus the
// feed forward result, and this program does that last add on the host. That is
// a gap in the generator rather than a choice made here: a caller who wires the
// two together and takes the second program's output as the block output gets a
// transformer block missing one of its two residual connections, silently and
// with plausible-looking numbers.
//
// # What is checked
//
// The reference recomputes the entire block in float64 from the same weights.
// Everything on the engine runs in fp16, so the comparison is to an fp16-scale
// tolerance and the program prints the worst difference it saw next to the
// magnitude of the reference, rather than claiming agreement.
//
// Two controls, because a matching number is weak evidence on its own. The
// compiled bundle is read for which backend the compiler chose, since the
// device mask is a permission and not a placement, and a block that fell back
// to the CPU would produce the same correct answer. And one input element is
// perturbed and the block re-dispatched, which must move the output and must
// still match a reference recomputed from the perturbed input; a program bound
// to a stale buffer, or returning a constant, passes the first check and fails
// this one.
//
//	go run ./examples/ane/e5rtblock
//	go run ./examples/ane/e5rtblock -dim 128 -heads 8 -seq 32 -hidden 256
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/tmc/apple/x/ane"
	"github.com/tmc/apple/x/ane/e5rt"
	"github.com/tmc/apple/x/ane/mil"
)

func main() {
	log.SetFlags(0)
	dim := flag.Int("dim", 64, "model width")
	heads := flag.Int("heads", 4, "attention heads")
	seq := flag.Int("seq", 16, "sequence length")
	hidden := flag.Int("hidden", 128, "feed forward width")
	flag.Parse()

	if *dim%*heads != 0 {
		log.Fatalf("dim %d is not a multiple of heads %d", *dim, *heads)
	}
	if err := run(*dim, *heads, *seq, *hidden); err != nil {
		log.Fatal(err)
	}
}

// blockWeights are the parameters of one transformer block.
type blockWeights struct {
	rms1 []float32 // [dim]
	wq   []float32 // [dim][dim] OIHW
	wk   []float32
	wv   []float32
	wo   []float32
	rms2 []float32 // [dim]
	w1   []float32 // [hidden][dim]
	w2   []float32 // [dim][hidden]
}

// randomWeights returns n deterministic pseudo-random values in [-spread,
// spread), from a xorshift generator seeded by seed.
//
// The generator matters. An earlier version of this program filled the weights
// from a smooth sine, and the attention scores came out with a spread of 0.004
// across the whole sequence — a softmax that uniform makes attention an average
// over positions, and the block reduces to something a wrong head split, a
// wrong transpose and a missing causal mask all reproduce exactly. Correlated
// weights cancel under a dot product over hundreds of terms; decorrelated ones
// accumulate like a random walk, which is what puts a real spread on the scores.
func randomWeights(n int, spread float64, seed uint64) []float32 {
	state := seed*2862933555777941757 + 3037000493
	next := func() float64 {
		state ^= state << 13
		state ^= state >> 7
		state ^= state << 17
		// The top 53 bits, mapped to [-1, 1).
		return float64(state>>11)/(1<<52) - 1
	}
	out := make([]float32, n)
	for i := range out {
		out[i] = float32(next() * spread)
	}
	return out
}

// newBlockWeights builds deterministic weights whose scales are chosen so that
// every stage of the block contributes.
//
// A projection sums over its input width, so a weight of scale s produces an
// output of roughly s*sqrt(width) for a unit-scale input. The scales below
// invert that for width 64: the query and key projections land near unit
// magnitude, which puts the scaled dot products over a range of a few and makes
// the softmax select rather than average; the output and second feed forward
// projections land near the residual's own magnitude, so neither the attention
// nor the feed forward result is lost against it.
//
// Everything on the engine is fp16, so the values stay well inside its range
// rather than being made as large as possible.
func newBlockWeights(dim, hidden int) *blockWeights {
	// RMSNorm weights vary around one: a constant would hide a channel misorder.
	rms := func(n int, seed uint64) []float32 {
		out := randomWeights(n, 0.2, seed)
		for i := range out {
			out[i] += 1
		}
		return out
	}
	inv := func(width int) float64 { return 1 / math.Sqrt(float64(width)) }
	return &blockWeights{
		rms1: rms(dim, 11),
		wq:   randomWeights(dim*dim, inv(dim), 22),
		wk:   randomWeights(dim*dim, inv(dim), 33),
		wv:   randomWeights(dim*dim, inv(dim), 44),
		wo:   randomWeights(dim*dim, 0.5*inv(dim), 55),
		rms2: rms(dim, 66),
		w1:   randomWeights(hidden*dim, inv(dim), 77),
		w2:   randomWeights(dim*hidden, 0.5*inv(hidden), 88),
	}
}

func run(dim, heads, seq, hidden int) error {
	lib, err := e5rt.Open()
	if err != nil {
		return err
	}
	fmt.Printf("resolved %d of %d listed symbols\n\n", len(lib.Resolved()), len(e5rt.Symbols))

	w := newBlockWeights(dim, hidden)
	attnDir, err := writeAttentionModel(dim, heads, seq, w)
	if err != nil {
		return err
	}
	ffnDir, err := writeFFNModel(dim, hidden, seq, w)
	if err != nil {
		return err
	}
	defer os.RemoveAll(attnDir)
	defer os.RemoveAll(ffnDir)

	fmt.Printf("block: dim %d, heads %d, seq %d, hidden %d\n", dim, heads, seq, hidden)

	start := time.Now()
	attnFn, err := compile(lib, attnDir)
	if err != nil {
		return fmt.Errorf("compile attention: %w", err)
	}
	ffnFn, err := compile(lib, ffnDir)
	if err != nil {
		return fmt.Errorf("compile feed forward: %w", err)
	}
	fmt.Printf("  compiled both programs in %v\n", time.Since(start).Round(time.Millisecond))
	reportBackend(attnDir, "attention")
	reportBackend(ffnDir, "feed forward")

	// Build both operations, then alias the attention output onto the feed
	// forward input.
	attnOp, err := newOperation(lib, attnFn)
	if err != nil {
		return err
	}
	ffnOp, err := newOperation(lib, ffnFn)
	if err != nil {
		return err
	}
	acts := dim * seq

	xBuf, xPtr, err := bindNewBuffer(lib, attnOp, "x", acts*2, true)
	if err != nil {
		return fmt.Errorf("bind attention input: %w", err)
	}
	_ = xBuf
	x2Buf, x2Ptr, err := bindNewBuffer(lib, attnOp, "out", acts*2, false)
	if err != nil {
		return fmt.Errorf("bind attention output: %w", err)
	}
	// The chain. The feed forward program's input port takes the buffer object
	// the attention program writes, rather than a buffer of its own.
	if err := bindExisting(lib, ffnOp, "x", x2Buf, true); err != nil {
		return fmt.Errorf("alias attention output onto feed forward input: %w", err)
	}
	_, fPtr, err := bindNewBuffer(lib, ffnOp, "out", acts*2, false)
	if err != nil {
		return fmt.Errorf("bind feed forward output: %w", err)
	}

	stream, err := lib.ExecutionStreamCreate()
	if err != nil {
		return err
	}
	if err := lib.EncodeOperation(stream, attnOp); err != nil {
		return fmt.Errorf("encode attention: %w", err)
	}
	if err := lib.EncodeOperation(stream, ffnOp); err != nil {
		return fmt.Errorf("encode feed forward: %w", err)
	}
	fmt.Printf("  encoded both operations into one stream, sharing the intermediate buffer\n")

	input := make([]float32, acts)
	for i := range input {
		input[i] = float32(0.5 * math.Sin(float64(i)*0.21))
	}

	if err := checkDiscriminating(input, w, dim, heads, seq, hidden); err != nil {
		return err
	}

	got, elapsed, err := dispatch(lib, stream, xPtr, x2Ptr, fPtr, input, acts)
	if err != nil {
		return err
	}
	fmt.Printf("  dispatch %v\n", elapsed.Round(time.Microsecond))
	want := blockReference(input, w, dim, heads, seq, hidden)
	if err := compare("block output", got, want); err != nil {
		return err
	}

	// Control: move one input element and require the output to follow.
	perturbed := append([]float32(nil), input...)
	perturbed[0] += 1
	got2, _, err := dispatch(lib, stream, xPtr, x2Ptr, fPtr, perturbed, acts)
	if err != nil {
		return err
	}
	moved := 0
	for i := range got {
		if got[i] != got2[i] {
			moved++
		}
	}
	if moved == 0 {
		return fmt.Errorf("perturbing an input changed nothing; the engine is not reading the bound buffer")
	}
	want2 := blockReference(perturbed, w, dim, heads, seq, hidden)
	if err := compare("perturbed block output", got2, want2); err != nil {
		return err
	}
	fmt.Printf("  control: perturbing one input element moved %d of %d outputs, and the new output still matches its reference\n",
		moved, len(got))
	return nil
}

// dispatch writes the input, runs the stream, and returns the block output.
// The stream is re-executed without re-encoding, which is what makes the second
// call a test of the bound buffers rather than of a fresh setup.
func dispatch(lib *e5rt.Lib, stream, xPtr, x2Ptr, fPtr uintptr, input []float32, acts int) ([]float32, time.Duration, error) {
	writeFP16(xPtr, input)
	start := time.Now()
	if err := lib.ExecuteSync(stream); err != nil {
		return nil, 0, fmt.Errorf("execute: %w", err)
	}
	elapsed := time.Since(start)

	// The block's second residual. See the package comment: the feed forward
	// generator does not add it.
	x2 := readFP16(x2Ptr, acts)
	f := readFP16(fPtr, acts)
	out := make([]float32, acts)
	for i := range out {
		out[i] = x2[i] + f[i]
	}
	return out, elapsed, nil
}

// blockReference recomputes the block in float64.
func blockReference(x []float32, w *blockWeights, dim, heads, seq, hidden int) []float64 {
	headDim := dim / heads
	in := make([]float64, len(x))
	for i, v := range x {
		in[i] = float64(v)
	}

	// Attention: RMSNorm, QKV, scaled dot product with a causal mask, output
	// projection, residual.
	xn := rmsNorm(in, w.rms1, dim, seq)
	q := project(xn, w.wq, dim, dim, seq)
	k := project(xn, w.wk, dim, dim, seq)
	v := project(xn, w.wv, dim, dim, seq)

	// attn[c][s] in the same [dim][seq] layout the engine uses.
	attn := make([]float64, dim*seq)
	scale := 1 / math.Sqrt(float64(headDim))
	for h := range heads {
		for i := range seq {
			// Scores against every position at or before i.
			scores := make([]float64, i+1)
			maxScore := math.Inf(-1)
			for j := 0; j <= i; j++ {
				var dot float64
				for d := range headDim {
					c := h*headDim + d
					dot += q[c*seq+i] * k[c*seq+j]
				}
				scores[j] = dot * scale
				maxScore = math.Max(maxScore, scores[j])
			}
			var sum float64
			for j := range scores {
				scores[j] = math.Exp(scores[j] - maxScore)
				sum += scores[j]
			}
			for d := range headDim {
				c := h*headDim + d
				var acc float64
				for j := range scores {
					acc += scores[j] * v[c*seq+j]
				}
				attn[c*seq+i] = acc / sum
			}
		}
	}
	proj := project(attn, w.wo, dim, dim, seq)
	x2 := make([]float64, dim*seq)
	for i := range x2 {
		x2[i] = in[i] + proj[i]
	}

	// Feed forward: RMSNorm, W1, ReLU squared, W2. The residual is added here
	// because the generator leaves it out.
	fn := rmsNorm(x2, w.rms2, dim, seq)
	h1 := project(fn, w.w1, hidden, dim, seq)
	for i, v := range h1 {
		if v < 0 {
			h1[i] = 0
		} else {
			h1[i] = v * v
		}
	}
	f := project(h1, w.w2, dim, hidden, seq)
	out := make([]float64, dim*seq)
	for i := range out {
		out[i] = x2[i] + f[i]
	}
	return out
}

// checkDiscriminating reports whether the block, with these weights and this
// input, is actually exercising the parts it claims to.
//
// It is not a check on the engine. It is a check on the test: a block whose
// attention scores are all equal computes an average over positions, and then a
// wrong head split, a wrong query-key transpose and a missing causal mask all
// produce the same numbers, so a match against the reference proves nothing
// about any of them. The same goes for a residual so much larger than what the
// block adds to it that the sum is the input again.
func checkDiscriminating(x []float32, w *blockWeights, dim, heads, seq, hidden int) error {
	headDim := dim / heads
	in := make([]float64, len(x))
	for i, v := range x {
		in[i] = float64(v)
	}
	xn := rmsNorm(in, w.rms1, dim, seq)
	q := project(xn, w.wq, dim, dim, seq)
	k := project(xn, w.wk, dim, dim, seq)

	scale := 1 / math.Sqrt(float64(headDim))
	lo, hi := math.Inf(1), math.Inf(-1)
	for h := range heads {
		for i := range seq {
			for j := 0; j <= i; j++ {
				var dot float64
				for d := range headDim {
					c := h*headDim + d
					dot += q[c*seq+i] * k[c*seq+j]
				}
				lo = math.Min(lo, dot*scale)
				hi = math.Max(hi, dot*scale)
			}
		}
	}
	spread := hi - lo
	fmt.Printf("  attention score spread %.3f", spread)
	if spread < 1 {
		fmt.Println()
		return fmt.Errorf("the softmax is nearly uniform, so attention degenerates to an average and this program cannot tell a correct block from several wrong ones")
	}

	// How much of the output the block contributes, against the residual.
	out := blockReference(x, w, dim, heads, seq, hidden)
	var deltaPeak, inPeak float64
	for i := range out {
		deltaPeak = math.Max(deltaPeak, math.Abs(out[i]-in[i]))
		inPeak = math.Max(inPeak, math.Abs(in[i]))
	}
	ratio := deltaPeak / inPeak
	fmt.Printf(", block contributes %.2fx the input magnitude\n", ratio)
	if ratio < 0.25 {
		return fmt.Errorf("the block barely moves its input, so the residual dominates and a broken block would still match")
	}
	return nil
}

// rmsNorm normalizes each position over the channel axis and applies the
// per-channel weight, matching the generators: mean of squares over channels,
// plus epsilon, raised to -1/2.
func rmsNorm(x []float64, weight []float32, dim, seq int) []float64 {
	out := make([]float64, dim*seq)
	for s := range seq {
		var sum float64
		for c := range dim {
			v := x[c*seq+s]
			sum += v * v
		}
		inv := math.Pow(sum/float64(dim)+1e-5, -0.5)
		for c := range dim {
			out[c*seq+s] = x[c*seq+s] * inv * float64(weight[c])
		}
	}
	return out
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

// compare reports the worst difference against an fp16-scale tolerance. It
// prints the peak magnitude of the reference alongside, so a small difference
// against a small reference is not read as agreement.
func compare(label string, got []float32, want []float64) error {
	if len(got) != len(want) {
		return fmt.Errorf("%s: %d values, want %d", label, len(got), len(want))
	}
	var worst, peak float64
	for i := range got {
		worst = math.Max(worst, math.Abs(float64(got[i])-want[i]))
		peak = math.Max(peak, math.Abs(want[i]))
	}
	// fp16 carries about three decimal digits, and this block accumulates over
	// four matrix products and a softmax, so the tolerance is proportional to
	// the magnitude rather than absolute.
	tol := 0.02*peak + 0.01
	if worst > tol {
		return fmt.Errorf("%s: worst |diff| %.4f exceeds tolerance %.4f (reference peak |%.4f|)", label, worst, tol, peak)
	}
	fmt.Printf("  %s matches the float64 reference over %d values: worst |diff| %.4f, tolerance %.4f, reference peak |%.4f|\n",
		label, len(got), worst, tol, peak)
	return nil
}

// writeAttentionModel writes the attention program and its six weight files.
func writeAttentionModel(dim, heads, seq int, w *blockWeights) (string, error) {
	rms1, err := mil.BuildVectorWeightBlob(w.rms1)
	if err != nil {
		return "", err
	}
	wq, err := mil.BuildWeightBlob(w.wq, dim, dim)
	if err != nil {
		return "", err
	}
	wk, err := mil.BuildWeightBlob(w.wk, dim, dim)
	if err != nil {
		return "", err
	}
	wv, err := mil.BuildWeightBlob(w.wv, dim, dim)
	if err != nil {
		return "", err
	}
	wo, err := mil.BuildWeightBlob(w.wo, dim, dim)
	if err != nil {
		return "", err
	}
	mask, err := mil.BuildCausalMaskBlob(seq)
	if err != nil {
		return "", err
	}
	return writeModel(mil.GenSDPAForward(dim, heads, seq), map[string][]byte{
		"rms1.bin": rms1,
		"wq.bin":   wq,
		"wk.bin":   wk,
		"wv.bin":   wv,
		"wo.bin":   wo,
		"mask.bin": mask,
	})
}

// writeFFNModel writes the feed forward program and its three weight files.
func writeFFNModel(dim, hidden, seq int, w *blockWeights) (string, error) {
	rms2, err := mil.BuildVectorWeightBlob(w.rms2)
	if err != nil {
		return "", err
	}
	w1, err := mil.BuildWeightBlob(w.w1, hidden, dim)
	if err != nil {
		return "", err
	}
	w2, err := mil.BuildWeightBlob(w.w2, dim, hidden)
	if err != nil {
		return "", err
	}
	return writeModel(mil.GenFFNForwardRMSReLU2(dim, hidden, seq), map[string][]byte{
		"rms2.bin": rms2,
		"w1.bin":   w1,
		"w2.bin":   w2,
	})
}

// writeModel writes a model directory holding text and its weight files.
func writeModel(text string, weights map[string][]byte) (string, error) {
	dir, err := os.MkdirTemp("", "e5rtblock")
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(filepath.Join(dir, "weights"), 0o755); err != nil {
		return "", err
	}
	if err := os.WriteFile(filepath.Join(dir, "model.mil"), []byte(text), 0o644); err != nil {
		return "", err
	}
	for name, blob := range weights {
		if err := os.WriteFile(filepath.Join(dir, "weights", name), blob, 0o644); err != nil {
			return "", err
		}
	}
	return dir, nil
}

// compile compiles the model in dir for the Neural Engine and retains "main".
func compile(lib *e5rt.Lib, dir string) (uintptr, error) {
	config, err := lib.CompilerConfigOptionsCreate()
	if err != nil {
		return 0, err
	}
	defer lib.CompilerConfigOptionsRelease(config)
	if err := lib.CompilerConfigOptionsSetCacheBundleLocation(config, dir); err != nil {
		return 0, err
	}
	compiler, err := lib.CompilerCreateWithConfig(config)
	if err != nil {
		return 0, err
	}
	defer lib.CompilerRelease(compiler)
	options, err := lib.CompilerOptionsCreate()
	if err != nil {
		return 0, err
	}
	defer lib.CompilerOptionsRelease(options)
	if err := lib.CompilerOptionsSetComputeDeviceTypesMask(options, e5rt.ComputeDeviceANE); err != nil {
		return 0, err
	}
	library, err := lib.CompilerCompile(compiler, filepath.Join(dir, "model.mil"), options)
	if err != nil {
		return 0, err
	}
	return lib.ProgramLibraryRetainProgramFunction(library, "main")
}

// newOperation builds the executable operation for a compiled function.
func newOperation(lib *e5rt.Lib, function uintptr) (uintptr, error) {
	options, err := lib.PrecompiledComputeOpOptionsCreate(function)
	if err != nil {
		return 0, err
	}
	defer lib.PrecompiledComputeOpOptionsRelease(options)
	if err := lib.PrecompiledComputeOpOptionsSetOperationName(options, "main"); err != nil {
		return 0, err
	}
	return lib.OperationCreatePrecompiled(options)
}

// bindNewBuffer retains a named port, allocates a CPU-visible buffer for it and
// binds the two, returning the buffer object and its host address.
func bindNewBuffer(lib *e5rt.Lib, op uintptr, name string, nbytes int, input bool) (uintptr, uintptr, error) {
	buf, err := lib.BufferObjectAlloc(uintptr(max((nbytes+63)&^63, 64)), 0)
	if err != nil {
		return 0, 0, err
	}
	ptr, err := lib.BufferObjectGetDataPtr(buf)
	if err != nil {
		return 0, 0, err
	}
	if err := bindExisting(lib, op, name, buf, input); err != nil {
		return 0, 0, err
	}
	return buf, ptr, nil
}

// bindExisting binds an already-allocated buffer object to a named port. This
// is what chains the two programs: one buffer, two ports, no copy.
func bindExisting(lib *e5rt.Lib, op uintptr, name string, buf uintptr, input bool) error {
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
// <function>_<backend> directories it leaves in the cache location. The device
// mask is a permission and not a placement, so this is the only local evidence
// that the work was put on the engine.
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
	if len(seen) == 0 {
		fmt.Printf("  backend for %s: UNKNOWN (no main_* directory in the compiled bundle)\n", label)
		return
	}
	names := make([]string, 0, len(seen))
	for name := range seen {
		names = append(names, name)
	}
	fmt.Printf("  backend for %s: the compiler emitted %v\n", label, names)
}

func writeFP16(dst uintptr, src []float32) {
	out := fp16SliceAt(dst, len(src))
	for i, v := range src {
		out[i] = ane.Float32ToFP16(v)
	}
}

func readFP16(src uintptr, n int) []float32 {
	in := fp16SliceAt(src, n)
	out := make([]float32, n)
	for i, v := range in {
		out[i] = ane.FP16ToFloat32(v)
	}
	return out
}
