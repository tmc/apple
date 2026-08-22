// Command threeaccel runs the same matrix multiplication on all three of
// Apple silicon's matrix engines from one Go process: the AMX blocks (via
// Accelerate's cblas_sgemm), the GPU (three ways: a naive Metal kernel, a
// simdgroup-matrix tiled kernel, and MPSMatrixMultiplication — Apple's
// tuned GEMM), and the Apple Neural Engine (via x/ane/dynamicmatmul).
//
// Every arm computes C = A×B for row-major float32 square matrices, is
// checksum-validated against a float64 CPU reference, and reports GFLOPS.
// Timing is calibrated to a 200ms floor per measurement rather than
// best-of-N single calls. GPU columns are kernel-only time from the
// command buffer's GPUStartTime/GPUEndTime; wall-clock (including
// submission latency) is reported in the notes. The ANE gets two columns:
// the dynamic path that re-sends weights every call, and the resident path
// that primes weights once and streams only activations. The MPS arm also
// runs in fp16 (reported in the notes) — all three engines are fp16-first
// and fp32 understates them.
//
// After the per-engine table, all three engines run the largest size
// concurrently from three goroutines — each on its fastest path —
// answering whether the engines contend when one process drives them all
// at once.
//
// The ANE path depends on private frameworks and is presence-checked: when
// unavailable the arm reports itself unavailable and the demo continues.
//
//	go run ./examples/compute/threeaccel
//	go run ./examples/compute/threeaccel -sizes 256,512,1024 -window 2s
package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
	"unsafe"

	"github.com/tmc/apple/accelerate"
	"github.com/tmc/apple/metal"
	mps "github.com/tmc/apple/metalperformanceshaders"
	"github.com/tmc/apple/x/ane/dynamicmatmul"
	"github.com/tmc/apple/x/powersample"
)

func init() {
	runtime.LockOSThread()
}

// kernelSource holds two hand-written matmuls: a deliberately naive one
// (one thread per output element, straight inner loop) and a
// simdgroup-matrix tiled one (each simdgroup accumulates an 8×8 tile of C
// through the hardware's matrix units). The gap between them — and between
// both and MPS — is the point: same silicon, three effort levels.
const kernelSource = `
#include <metal_stdlib>
using namespace metal;

kernel void matmul(device const float *a [[buffer(0)]],
                   device const float *b [[buffer(1)]],
                   device float *c       [[buffer(2)]],
                   constant uint &n      [[buffer(3)]],
                   uint2 gid             [[thread_position_in_grid]]) {
    if (gid.x >= n || gid.y >= n) {
        return;
    }
    float acc = 0.0f;
    for (uint k = 0; k < n; k++) {
        acc += a[gid.y * n + k] * b[k * n + gid.x];
    }
    c[gid.y * n + gid.x] = acc;
}

// One simdgroup (32 threads) per threadgroup; each simdgroup owns one
// 8x8 output tile and marches down the k dimension 8 columns at a time.
// Requires n % 8 == 0.
kernel void matmul_simd(device const float *a [[buffer(0)]],
                        device const float *b [[buffer(1)]],
                        device float *c       [[buffer(2)]],
                        constant uint &n      [[buffer(3)]],
                        uint2 tgid            [[threadgroup_position_in_grid]]) {
    const uint row = tgid.y * 8;
    const uint col = tgid.x * 8;
    simdgroup_float8x8 acc(0.0f);
    for (uint k = 0; k < n; k += 8) {
        simdgroup_float8x8 ma, mb;
        simdgroup_load(ma, a + row * n + k, n);
        simdgroup_load(mb, b + k * n + col, n);
        simdgroup_multiply_accumulate(acc, ma, mb, acc);
    }
    simdgroup_store(acc, c + row * n + col, n);
}
`

type result struct {
	gflops float64
	extra  string // arm-specific note (ANE hardware time, GPU wall-clock)
	err    error
}

func main() {
	log.SetFlags(0)
	sizesFlag := flag.String("sizes", "256,512,1024", "comma-separated square matrix sizes (multiples of 8)")
	window := flag.Duration("window", 2*time.Second, "measurement window for the concurrent phase")
	seed := flag.Int64("seed", 1, "RNG seed for matrix contents")
	power := flag.Bool("power", false, "meter CPU/GPU/ANE energy (unprivileged via IOReport; powermetrics fallback)")
	flag.Parse()

	sizes, err := parseSizes(*sizesFlag)
	if err != nil {
		log.Fatalf("threeaccel: %v", err)
	}
	rng := rand.New(rand.NewSource(*seed))

	g, err := newGPU()
	if err != nil {
		log.Fatalf("threeaccel: %v", err)
	}

	fmt.Println("One matmul, three engines: AMX (Accelerate BLAS), GPU (naive/simdgroup/MPS), ANE (x/ane)")
	fmt.Println("All arms checksum-validated against a float64 CPU reference. GFLOPS, higher is better.")
	fmt.Println("GPU columns are kernel-only time; ANE-res primes weights once and streams activations.")
	fmt.Println()
	fmt.Printf("%-6s  %9s  %9s  %9s  %9s  %9s  %9s\n", "Size", "AMX", "GPU-naive", "GPU-simd", "GPU-MPS", "ANE", "ANE-res")
	fmt.Println("------  ---------  ---------  ---------  ---------  ---------  ---------")

	var notes []string
	var solo map[string]result
	for _, n := range sizes {
		a := randMatrix(rng, n*n)
		b := randMatrix(rng, n*n)
		want := reference(a, b, n)

		amx := benchAMX(a, b, n, want)
		naive, simd, mpsRes := g.bench(a, b, n, want)
		ane, aneRes := benchANE(a, b, n, want)

		fmt.Printf("%-6d  %9s  %9s  %9s  %9s  %9s  %9s\n",
			n, cell(amx), cell(naive), cell(simd), cell(mpsRes), cell(ane), cell(aneRes))
		for _, r := range []struct {
			name string
			res  result
		}{{"AMX", amx}, {"GPU-naive", naive}, {"GPU-simd", simd}, {"GPU-MPS", mpsRes}, {"ANE", ane}, {"ANE-res", aneRes}} {
			if r.res.err != nil {
				notes = append(notes, fmt.Sprintf("n=%d %s: %v", n, r.name, r.res.err))
			} else if r.res.extra != "" {
				notes = append(notes, fmt.Sprintf("n=%d %s: %s", n, r.name, r.res.extra))
			}
		}
		solo = map[string]result{"AMX": amx, "GPU": mpsRes, "ANE": aneRes}
	}

	if len(notes) > 0 {
		fmt.Println()
		for _, s := range notes {
			fmt.Println("  " + s)
		}
	}

	nMax := sizes[len(sizes)-1]
	concurrent(g, rng, nMax, *window, solo)
	if *power {
		powerPhase(g, rng, nMax, *window)
	}
}

func cell(r result) string {
	if r.err != nil {
		return "unavail"
	}
	return fmt.Sprintf("%.1f", r.gflops)
}

// calibrateIters finds an iteration count whose total runtime is at least
// 200ms, so a measurement is never a single (possibly flaky) call.
func calibrateIters(fn func()) int {
	iters := 1
	for {
		start := time.Now()
		for range iters {
			fn()
		}
		if time.Since(start) >= 200*time.Millisecond {
			return iters
		}
		iters *= 2
	}
}

// perOp times fn calibrated to the 200ms floor and returns the average
// duration of one call.
func perOp(fn func()) time.Duration {
	iters := calibrateIters(fn)
	start := time.Now()
	for range iters {
		fn()
	}
	return time.Since(start) / time.Duration(iters)
}

// benchAMX times Accelerate's cblas_sgemm, which Apple routes to the AMX
// coprocessor blocks for this shape.
func benchAMX(a, b []float32, n int, want []float64) result {
	c := make([]float32, n*n)
	run := func() {
		accelerate.Cblas_sgemm(
			accelerate.CblasRowMajor,
			accelerate.CblasNoTrans, accelerate.CblasNoTrans,
			n, n, n,
			1.0, a, n, b, n, 0.0, c, n)
	}
	run()
	if err := verify(c, want, 1e-4); err != nil {
		return result{err: fmt.Errorf("verification: %w", err)}
	}
	return result{gflops: gflops(n, perOp(run))}
}

// benchANE times x/ane/dynamicmatmul twice: the dynamic path (weights
// re-sent with every eval) and the resident path (weights primed once via
// PrimeWeightsIO, only activations move per eval). The executor compiles
// once per shape; compile cost is reported separately and excluded from
// the timing, matching how the other arms exclude pipeline setup.
func benchANE(a, b []float32, n int, want []float64) (dynamic, resident result) {
	compileStart := time.Now()
	ex, err := dynamicmatmul.New(n, n, n, dynamicmatmul.Options{})
	if err != nil {
		e := result{err: fmt.Errorf("ANE unavailable: %w", err)}
		return e, e
	}
	defer ex.Close()
	compile := time.Since(compileStart)

	out, err := ex.Eval(a, b)
	if err != nil {
		e := result{err: err}
		return e, e
	}
	// The ANE computes in reduced precision internally; float32 in and out,
	// but the accumulation error is fp16-scale. The tolerance says so
	// honestly rather than pretending fp32 accuracy.
	if err := verify(out, want, 2e-2); err != nil {
		e := result{err: fmt.Errorf("verification: %w", err)}
		return e, e
	}

	dst := make([]float32, n*n)
	var hwNS uint64
	dynDur := perOp(func() {
		st, evalErr := ex.EvalInto(dst, a, b)
		if evalErr != nil {
			err = evalErr
			return
		}
		if st.HWExecutionNS > 0 {
			hwNS = st.HWExecutionNS
		}
	})
	if err != nil {
		return result{err: err}, result{err: err}
	}
	extra := fmt.Sprintf("compile %v", compile.Round(time.Millisecond))
	if hwNS > 0 {
		extra += fmt.Sprintf(", hardware-reported eval %v (%.1f GFLOPS on-device)",
			time.Duration(hwNS).Round(time.Microsecond), gflops(n, time.Duration(hwNS)))
	}
	dynamic = result{gflops: gflops(n, dynDur), extra: extra}

	// Resident path: weights live on the ANE side; each eval moves only
	// the [inDim, batch] channel-first activation tensor.
	if err := ex.PrimeWeightsIO(b); err != nil {
		return dynamic, result{err: fmt.Errorf("prime weights: %w", err)}
	}
	xCF := transpose(a, n)
	dstCF := make([]float32, n*n)
	if _, err := ex.EvalCFIOInto(dstCF, xCF); err != nil {
		return dynamic, result{err: err}
	}
	if err := verify(transpose(dstCF, n), want, 2e-2); err != nil {
		return dynamic, result{err: fmt.Errorf("verification: %w", err)}
	}
	var resHWNS uint64
	resDur := perOp(func() {
		st, evalErr := ex.EvalCFIOInto(dstCF, xCF)
		if evalErr != nil {
			err = evalErr
			return
		}
		if st.HWExecutionNS > 0 {
			resHWNS = st.HWExecutionNS
		}
	})
	if err != nil {
		return dynamic, result{err: err}
	}
	var resExtra string
	if resHWNS > 0 {
		resExtra = fmt.Sprintf("hardware-reported eval %v (%.1f GFLOPS on-device)",
			time.Duration(resHWNS).Round(time.Microsecond), gflops(n, time.Duration(resHWNS)))
	}
	resident = result{gflops: gflops(n, resDur), extra: resExtra}
	return dynamic, resident
}

// gpu owns the Metal state shared by all GPU arms, reused across sizes.
type gpu struct {
	device metal.MTLDeviceObject
	queue  metal.MTLCommandQueue
	naive  metal.MTLComputePipelineState
	simd   metal.MTLComputePipelineState
}

func newGPU() (*gpu, error) {
	device := metal.MTLCreateSystemDefaultDevice()
	if device.GetID() == 0 {
		return nil, fmt.Errorf("no Metal device")
	}
	lib, err := device.NewLibraryWithSourceOptionsError(kernelSource, nil)
	if err != nil {
		return nil, fmt.Errorf("shader compilation: %w", err)
	}
	g := &gpu{device: device, queue: device.NewCommandQueue()}
	for _, k := range []struct {
		name string
		dst  *metal.MTLComputePipelineState
	}{{"matmul", &g.naive}, {"matmul_simd", &g.simd}} {
		fn := lib.NewFunctionWithName(k.name)
		if fn.GetID() == 0 {
			return nil, fmt.Errorf("kernel %s not found", k.name)
		}
		p, err := device.NewComputePipelineStateWithFunctionError(fn)
		if err != nil {
			return nil, fmt.Errorf("pipeline %s: %w", k.name, err)
		}
		*k.dst = p
	}
	return g, nil
}

// buffers holds one size's uploaded operands and mapped output.
type buffers struct {
	bufA, bufB, bufC metal.MTLBuffer
	c                []float32
}

func (g *gpu) upload(a, b []float32, n int) (*buffers, error) {
	bytes := uint(n) * uint(n) * 4
	bufA := g.device.NewBufferWithBytesLengthOptions(unsafe.Pointer(&a[0]), bytes, metal.MTLResourceStorageModeShared)
	bufB := g.device.NewBufferWithBytesLengthOptions(unsafe.Pointer(&b[0]), bytes, metal.MTLResourceStorageModeShared)
	bufC := g.device.NewBufferWithLengthOptions(bytes, metal.MTLResourceStorageModeShared)
	runtime.KeepAlive(a)
	runtime.KeepAlive(b)
	if bufA.GetID() == 0 || bufB.GetID() == 0 || bufC.GetID() == 0 {
		return nil, fmt.Errorf("buffer allocation failed")
	}
	return &buffers{
		bufA: bufA, bufB: bufB, bufC: bufC,
		c: unsafe.Slice((*float32)(bufC.Contents()), n*n),
	}, nil
}

// kernelRun returns a closure that dispatches one matmul through the given
// hand-written pipeline and reports GPU-side kernel seconds.
func (g *gpu) kernelRun(pipeline metal.MTLComputePipelineState, bufs *buffers, n int, simdTiled bool) func() float64 {
	nDim := uint32(n)
	nBytes := unsafe.Slice((*byte)(unsafe.Pointer(&nDim)), 4)
	return func() float64 {
		cb := g.queue.CommandBuffer()
		enc := cb.ComputeCommandEncoder()
		enc.SetComputePipelineState(pipeline)
		enc.SetBufferWithOffsetAtIndex(bufs.bufA, 0, 0)
		enc.SetBufferWithOffsetAtIndex(bufs.bufB, 0, 1)
		enc.SetBufferWithOffsetAtIndex(bufs.bufC, 0, 2)
		enc.SetBytesLengthAtIndex(nBytes, 3)
		if simdTiled {
			enc.DispatchThreadgroupsThreadsPerThreadgroup(
				metal.MTLSize{Width: uint(n) / 8, Height: uint(n) / 8, Depth: 1},
				metal.MTLSize{Width: 32, Height: 1, Depth: 1})
		} else {
			enc.DispatchThreadsThreadsPerThreadgroup(
				metal.MTLSize{Width: uint(n), Height: uint(n), Depth: 1},
				metal.MTLSize{Width: 16, Height: 16, Depth: 1})
		}
		enc.EndEncoding()
		cb.Commit()
		cb.WaitUntilCompleted()
		return float64(cb.GPUEndTime() - cb.GPUStartTime())
	}
}

// mpsRun returns a closure that dispatches one MPSMatrixMultiplication —
// Apple's tuned GEMM — over the same buffers and reports kernel seconds.
func (g *gpu) mpsRun(bufs *buffers, n int) func() float64 {
	desc := mps.NewMatrixDescriptorWithRowsColumnsRowBytesDataType(uint(n), uint(n), uint(n)*4, mps.MPSDataTypeFloat32)
	left := mps.NewMatrixWithBufferDescriptor(bufs.bufA, desc)
	right := mps.NewMatrixWithBufferDescriptor(bufs.bufB, desc)
	product := mps.NewMatrixWithBufferDescriptor(bufs.bufC, desc)
	mm := mps.NewMatrixMultiplicationWithDeviceTransposeLeftTransposeRightResultRowsResultColumnsInteriorColumnsAlphaBeta(
		g.device, false, false, uint(n), uint(n), uint(n), 1, 0)
	return func() float64 {
		cb := g.queue.CommandBuffer()
		mm.EncodeToCommandBufferLeftMatrixRightMatrixResultMatrix(cb, left, right, product)
		cb.Commit()
		cb.WaitUntilCompleted()
		return float64(cb.GPUEndTime() - cb.GPUStartTime())
	}
}

// mpsRunFP16 is mpsRun with float16 operands: the numbers all three
// engines really want. Inputs are converted on the CPU; the product comes
// back float16 and is widened for verification.
func (g *gpu) mpsRunFP16(a, b []float32, n int) (run func() float64, readC func() []float32, err error) {
	bytes := uint(n) * uint(n) * 2
	a16 := toFloat16(a)
	b16 := toFloat16(b)
	bufA := g.device.NewBufferWithBytesLengthOptions(unsafe.Pointer(&a16[0]), bytes, metal.MTLResourceStorageModeShared)
	bufB := g.device.NewBufferWithBytesLengthOptions(unsafe.Pointer(&b16[0]), bytes, metal.MTLResourceStorageModeShared)
	bufC := g.device.NewBufferWithLengthOptions(bytes, metal.MTLResourceStorageModeShared)
	runtime.KeepAlive(a16)
	runtime.KeepAlive(b16)
	if bufA.GetID() == 0 || bufB.GetID() == 0 || bufC.GetID() == 0 {
		return nil, nil, fmt.Errorf("fp16 buffer allocation failed")
	}
	desc := mps.NewMatrixDescriptorWithRowsColumnsRowBytesDataType(uint(n), uint(n), uint(n)*2, mps.MPSDataTypeFloat16)
	left := mps.NewMatrixWithBufferDescriptor(bufA, desc)
	right := mps.NewMatrixWithBufferDescriptor(bufB, desc)
	product := mps.NewMatrixWithBufferDescriptor(bufC, desc)
	mm := mps.NewMatrixMultiplicationWithDeviceTransposeLeftTransposeRightResultRowsResultColumnsInteriorColumnsAlphaBeta(
		g.device, false, false, uint(n), uint(n), uint(n), 1, 0)
	c16 := unsafe.Slice((*uint16)(bufC.Contents()), n*n)
	run = func() float64 {
		cb := g.queue.CommandBuffer()
		mm.EncodeToCommandBufferLeftMatrixRightMatrixResultMatrix(cb, left, right, product)
		cb.Commit()
		cb.WaitUntilCompleted()
		return float64(cb.GPUEndTime() - cb.GPUStartTime())
	}
	readC = func() []float32 {
		out := make([]float32, len(c16))
		for i, h := range c16 {
			out[i] = float16to32(h)
		}
		return out
	}
	return run, readC, nil
}

// bench measures the three GPU arms over shared operand buffers. Columns
// are kernel-only GFLOPS; wall-clock and the fp16 MPS run go in the notes.
func (g *gpu) bench(a, b []float32, n int, want []float64) (naive, simd, mpsResult result) {
	bufs, err := g.upload(a, b, n)
	if err != nil {
		e := result{err: err}
		return e, e, e
	}

	measure := func(run func() float64, tol float64) result {
		run()
		if err := verify(bufs.c, want, tol); err != nil {
			return result{err: fmt.Errorf("verification: %w", err)}
		}
		iters := calibrateIters(func() { run() })
		var kernelSec float64
		start := time.Now()
		for range iters {
			kernelSec += run()
		}
		wall := time.Since(start) / time.Duration(iters)
		kernel := time.Duration(kernelSec / float64(iters) * float64(time.Second))
		return result{
			gflops: gflops(n, kernel),
			extra:  fmt.Sprintf("wall-clock %.1f GFLOPS (kernel %v, wall %v per op)", gflops(n, wall), kernel.Round(time.Microsecond), wall.Round(time.Microsecond)),
		}
	}

	naive = measure(g.kernelRun(g.naive, bufs, n, false), 1e-4)
	if n%8 == 0 {
		simd = measure(g.kernelRun(g.simd, bufs, n, true), 1e-4)
	} else {
		simd = result{err: fmt.Errorf("requires n %% 8 == 0")}
	}
	mpsResult = measure(g.mpsRun(bufs, n), 1e-4)

	// fp16 MPS: same tuned GEMM at the precision the hardware prefers.
	// Tolerance is fp16-scale, like the ANE's.
	if run16, readC, err := g.mpsRunFP16(a, b, n); err == nil {
		kern16 := run16()
		if verr := verify(readC(), want, 2e-2); verr != nil {
			mpsResult.extra += fmt.Sprintf("; fp16 verification: %v", verr)
		} else {
			iters := calibrateIters(func() { run16() })
			var kernelSec float64
			for range iters {
				kernelSec += run16()
			}
			kern16 = kernelSec / float64(iters)
			mpsResult.extra += fmt.Sprintf("; fp16 %.1f GFLOPS kernel-only", gflops(n, time.Duration(kern16*float64(time.Second))))
		}
	}
	return naive, simd, mpsResult
}

// concurrent drives all available engines at once for the given window,
// one goroutine per engine — each on its fastest path (AMX sgemm, GPU via
// MPS, ANE resident-weights) — and reports per-engine and aggregate
// throughput next to the solo numbers: the contention question no
// sequential table can answer.
// engineArm is one accelerator running one n×n×n matmul, plus the energy
// rail that work lands on.
type engineArm struct {
	name string
	run  func() // one operation
	rail func(powersample.Power) float64
}

// buildArms constructs the one-op-per-call arms for the engines available
// on this machine (AMX sgemm, GPU MPS, ANE resident-weights). The caller
// must call cleanup when done with the arms.
func buildArms(g *gpu, rng *rand.Rand, n int) (arms []engineArm, cleanup func()) {
	a := randMatrix(rng, n*n)
	b := randMatrix(rng, n*n)
	cleanup = func() {}

	amxC := make([]float32, n*n)
	arms = append(arms, engineArm{"AMX", func() {
		accelerate.Cblas_sgemm(
			accelerate.CblasRowMajor,
			accelerate.CblasNoTrans, accelerate.CblasNoTrans,
			n, n, n,
			1.0, a, n, b, n, 0.0, amxC, n)
	}, func(p powersample.Power) float64 { return p.CPU }})

	bufs, err := g.upload(a, b, n)
	if err != nil {
		fmt.Printf("\nGPU arm unavailable: %v\n", err)
	} else {
		gpuRun := g.mpsRun(bufs, n)
		arms = append(arms, engineArm{"GPU", func() { gpuRun() },
			func(p powersample.Power) float64 { return p.GPU }})
	}

	ex, err := dynamicmatmul.New(n, n, n, dynamicmatmul.Options{})
	if err == nil {
		err = ex.PrimeWeightsIO(b)
	}
	if err == nil {
		cleanup = func() { ex.Close() }
		xCF := transpose(a, n)
		dstCF := make([]float32, n*n)
		arms = append(arms, engineArm{"ANE", func() {
			if _, evalErr := ex.EvalCFIOInto(dstCF, xCF); evalErr != nil {
				log.Fatalf("threeaccel: ANE eval: %v", evalErr)
			}
		}, func(p powersample.Power) float64 { return p.ANE }})
	}
	return arms, cleanup
}

func concurrent(g *gpu, rng *rand.Rand, n int, window time.Duration, solo map[string]result) {
	arms, cleanup := buildArms(g, rng, n)
	defer cleanup()

	fmt.Printf("\nConcurrent: %d engines at once, n=%d, %v window (AMX sgemm + GPU MPS + ANE resident)\n", len(arms), n, window)

	iters := make([]int, len(arms))
	elapsed := make([]time.Duration, len(arms))
	start := make(chan struct{})
	var wg sync.WaitGroup
	for i, am := range arms {
		wg.Go(func() {
			<-start
			t0 := time.Now()
			deadline := t0.Add(window)
			for time.Now().Before(deadline) {
				am.run()
				iters[i]++
			}
			elapsed[i] = time.Since(t0)
		})
	}
	close(start)
	wg.Wait()

	var aggregate float64
	for i, am := range arms {
		gf := 2 * float64(n) * float64(n) * float64(n) * float64(iters[i]) / elapsed[i].Seconds() / 1e9
		aggregate += gf
		line := fmt.Sprintf("  %-4s  %8.1f GFLOPS", am.name, gf)
		if s, ok := solo[am.name]; ok && s.err == nil && s.gflops > 0 {
			line += fmt.Sprintf("  (solo %.1f, %+.0f%%)", s.gflops, (gf/s.gflops-1)*100)
		}
		fmt.Println(line)
	}
	fmt.Printf("  aggregate %.1f GFLOPS\n", aggregate)
	if len(arms) < 3 {
		fmt.Printf("  (only %d of 3 engines available)\n", len(arms))
	}
}

// powerPhase runs each engine alone for the window under an energy meter
// and prints GFLOPS per watt, the number that decides which engine a
// battery-powered caller should pick. Each arm is charged only for its
// own rail: the GPU and ANE rails draw ~0 at idle so their numbers need
// no baseline, but AMX runs on the CPU rail, which also carries every
// other process — its row is only meaningful on a quiet machine.
func powerPhase(g *gpu, rng *rand.Rand, n int, window time.Duration) {
	arms, cleanup := buildArms(g, rng, n)
	defer cleanup()

	fmt.Printf("\nPerf per watt: each engine alone, n=%d, %v window, own rail only\n", n, window)
	for _, am := range arms {
		am.run() // warm the path before the meter starts
		m, err := powersample.Start(500 * time.Millisecond)
		if err != nil {
			log.Fatalf("threeaccel: start power meter: %v", err)
		}
		iters := 0
		t0 := time.Now()
		deadline := t0.Add(window)
		for time.Now().Before(deadline) {
			am.run()
			iters++
		}
		elapsed := time.Since(t0)
		r, err := m.Stop()
		if err != nil {
			log.Fatalf("threeaccel: stop power meter: %v", err)
		}
		gf := 2 * float64(n) * float64(n) * float64(n) * float64(iters) / elapsed.Seconds() / 1e9
		joules := am.rail(r.Energy)
		if joules <= 0 {
			fmt.Printf("  %-4s  %8.1f GFLOPS   rail reported 0 J — not attributable\n", am.name, gf)
			continue
		}
		// The meter window and the timed loop overlap but are not equal;
		// charge the rail's average watts against the loop's GFLOPS.
		watts := joules / r.Duration.Seconds()
		line := fmt.Sprintf("  %-4s  %8.1f GFLOPS  %6.2f W on its rail  %8.1f GFLOPS/W", am.name, gf, watts, gf/watts)
		if am.name == "AMX" {
			line += "  (CPU rail: includes every other process — quiet machine only)"
		}
		fmt.Println(line)
	}
}

// reference computes C = A×B accumulating in float64.
func reference(a, b []float32, n int) []float64 {
	c := make([]float64, n*n)
	for i := range n {
		for k := range n {
			av := float64(a[i*n+k])
			row := b[k*n:]
			out := c[i*n:]
			for j := range n {
				out[j] += av * float64(row[j])
			}
		}
	}
	return c
}

// transpose returns the transpose of an n×n row-major matrix, which is the
// same bytes reinterpreted channel-first — the layout the ANE's CF entry
// points take.
func transpose(m []float32, n int) []float32 {
	t := make([]float32, n*n)
	for i := range n {
		for j := range n {
			t[j*n+i] = m[i*n+j]
		}
	}
	return t
}

// verify checks got against the float64 reference under a relative
// tolerance scaled by the magnitude of the result.
func verify(got []float32, want []float64, tol float64) error {
	var maxAbs float64
	for _, w := range want {
		if a := math.Abs(w); a > maxAbs {
			maxAbs = a
		}
	}
	if maxAbs == 0 {
		maxAbs = 1
	}
	var worst float64
	worstAt := -1
	for i, w := range want {
		if d := math.Abs(float64(got[i])-w) / maxAbs; d > worst {
			worst, worstAt = d, i
		}
	}
	if worst > tol {
		return fmt.Errorf("max relative error %.2e at index %d exceeds %.0e", worst, worstAt, tol)
	}
	return nil
}

func gflops(n int, d time.Duration) float64 {
	return 2 * float64(n) * float64(n) * float64(n) / d.Seconds() / 1e9
}

func parseSizes(s string) ([]int, error) {
	var sizes []int
	for f := range strings.SplitSeq(s, ",") {
		n, err := strconv.Atoi(strings.TrimSpace(f))
		if err != nil || n <= 0 {
			return nil, fmt.Errorf("bad size %q", f)
		}
		sizes = append(sizes, n)
	}
	return sizes, nil
}

func randMatrix(rng *rand.Rand, size int) []float32 {
	m := make([]float32, size)
	for i := range m {
		m[i] = rng.Float32()*2 - 1
	}
	return m
}

// toFloat16 converts float32 values to IEEE 754 half-precision bit
// patterns (round toward zero; inputs here are within fp16 normal range).
func toFloat16(f []float32) []uint16 {
	h := make([]uint16, len(f))
	for i, v := range f {
		b := math.Float32bits(v)
		sign := uint16(b>>16) & 0x8000
		exp := int32(b>>23&0xff) - 127 + 15
		man := b & 0x7fffff
		switch {
		case exp <= 0:
			h[i] = sign // flush denormals to zero; inputs are O(1)
		case exp >= 31:
			h[i] = sign | 0x7c00
		default:
			h[i] = sign | uint16(exp)<<10 | uint16(man>>13)
		}
	}
	return h
}

// float16to32 widens an IEEE 754 half-precision bit pattern to float32.
func float16to32(h uint16) float32 {
	sign := uint32(h&0x8000) << 16
	exp := uint32(h >> 10 & 0x1f)
	man := uint32(h & 0x3ff)
	switch exp {
	case 0:
		if man == 0 {
			return math.Float32frombits(sign)
		}
		// Denormal: normalize.
		e := uint32(112)
		for man&0x400 == 0 {
			man <<= 1
			e--
		}
		return math.Float32frombits(sign | (e+1)<<23 | (man&0x3ff)<<13)
	case 31:
		return math.Float32frombits(sign | 0x7f800000 | man<<13)
	default:
		return math.Float32frombits(sign | (exp+112)<<23 | man<<13)
	}
}
