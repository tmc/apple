// Command espressoe5rtdispatch runs a MIL program on the Apple Neural Engine through
// Espresso's private e5rt_* direct dispatch route, and checks the answer.
// It calls the generated private/espresso binding through the example-local
// espressoe5rt adapter.
//
// The route skips Core ML entirely. A MIL program and its weights are written
// to a directory, Espresso's compiler lowers them to a signed hardware program
// inside the aned daemon, and the result is bound to buffer objects and
// submitted on an execution stream:
//
//	config -> compiler -> options -> compile -> library -> function
//	       -> operation -> ports -> buffers -> stream -> execute
//
// Compile runs once and is expensive; bind and dispatch are the hot loop.
//
// Every stage prints what it observed. The demo's claim to have run on the
// Neural Engine rests on the last stage: the output buffer is compared against
// a float64 CPU reference computed in Go, so a route that silently returned
// zeros, or ran the wrong program, fails the check rather than printing a
// success it did not earn.
//
// Two programs run. The first has a single weight tensor in its own file, the
// shape every MIL template in this repository emits. The second packs two
// weight tensors into one file at two different offsets, which is how the
// reference implementation lays out real models; it exercises
// [github.com/tmc/apple/x/ane/mil.BlobWriter] multi-entry output, a path this
// repository has never previously submitted to the compiler.
//
//	go run ./examples/ane/espressoe5rtdispatch
//	go run ./examples/ane/espressoe5rtdispatch -in 64 -out 64 -spatial 32
//	go run ./examples/ane/espressoe5rtdispatch -device cpu
//
// The e5rt_* argument lists are not Apple-documented. They come from ANEForge,
// the source paper author's own working implementation. If a signature is
// wrong the process faults in C, where the fault is fatal and recover cannot
// see it, so a crash here is a real result and is reported as one.
package main

import (
	"flag"
	"fmt"
	"log"
	"maps"
	"math"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
	"unsafe"

	e5rt "github.com/tmc/apple/examples/ane/internal/espressoe5rt"
	"github.com/tmc/apple/x/ane"
	"github.com/tmc/apple/x/ane/mil"
)

func main() {
	log.SetFlags(0)
	inCh := flag.Int("in", 16, "input channels")
	outCh := flag.Int("out", 16, "output channels")
	spatial := flag.Int("spatial", 8, "spatial width")
	device := flag.String("device", "ane", "compute device: ane, cpu, gpu, or auto")
	iters := flag.Int("iters", 10, "dispatches to time after the first")
	keep := flag.String("cache", "", "keep compiler cache bundles under this directory instead of a temporary one")
	multiCh := flag.Int("multi", 12, "channels for the two-weight program; the default is chosen so the first tensor is not a multiple of 64 bytes")
	flag.Parse()

	mask, err := deviceMask(*device)
	if err != nil {
		log.Fatal(err)
	}

	lib, err := e5rt.Open()
	if err != nil {
		fmt.Println("Espresso e5rt route unavailable:", err)
		return
	}
	fmt.Printf("opened %s\n", e5rt.FrameworkPath)
	fmt.Printf("resolved %d of %d listed symbols\n\n", len(lib.Resolved()), len(e5rt.Symbols))

	cacheRoot = *keep
	if err := runSingleWeight(lib, mask, *inCh, *outCh, *spatial, *iters); err != nil {
		log.Fatalf("single-weight program: %v", err)
	}
	fmt.Println()
	if err := runTwoWeightsOneFile(lib, mask, *multiCh, *spatial); err != nil {
		log.Fatalf("two-weight program: %v", err)
	}
}

func deviceMask(name string) (uint64, error) {
	switch name {
	case "ane":
		return e5rt.ComputeDeviceANE, nil
	case "cpu":
		return e5rt.ComputeDeviceCPU, nil
	case "gpu":
		return e5rt.ComputeDeviceGPU, nil
	case "auto":
		return e5rt.ComputeDeviceCPU | e5rt.ComputeDeviceGPU | e5rt.ComputeDeviceANE, nil
	}
	return 0, fmt.Errorf("unknown device %q: want ane, cpu, gpu, or auto", name)
}

// runSingleWeight compiles and dispatches a 1x1 convolution whose one weight
// tensor lives in its own file, then checks the output against a CPU reference.
func runSingleWeight(lib *e5rt.Lib, mask uint64, inCh, outCh, spatial, iters int) error {
	fmt.Printf("program 1: 1x1 conv, %d->%d channels, %d spatial, one weight file\n", inCh, outCh, spatial)

	weights := make([]float32, outCh*inCh)
	for i := range weights {
		weights[i] = float32(i%7-3) * 0.25
	}
	blob, err := mil.BuildWeightBlob(weights, outCh, inCh)
	if err != nil {
		return err
	}

	dir, err := os.MkdirTemp("", "e5rtdispatch-single-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	if err := writeModel(dir, mil.GenConvFP16IO(inCh, outCh, spatial), map[string][]byte{
		"weights/weight.bin": blob,
	}); err != nil {
		return err
	}

	input := make([]float32, inCh*spatial)
	for i := range input {
		input[i] = float32(i%5) * 0.5
	}
	// The control: a second input that differs in one element. If the engine
	// were returning a constant, a cached result, or zeros, the two dispatches
	// would agree with each other and disagree with the references.
	perturbed := append([]float32(nil), input...)
	perturbed[0] += 4

	want := convReference(weights, input, outCh, inCh, spatial)
	wantPerturbed := convReference(weights, perturbed, outCh, inCh, spatial)

	got, stats, err := dispatch(lib, dir, mask, [][]float32{input, perturbed}, outCh*spatial, iters)
	if err != nil {
		return err
	}
	stats.report()
	if err := compare(got[0], want); err != nil {
		return err
	}
	return checkResponds(got[0], got[1], wantPerturbed)
}

// runTwoWeightsOneFile compiles a program whose two weight tensors are packed
// into a single blob file at two offsets, the layout real models use.
func runTwoWeightsOneFile(lib *e5rt.Lib, mask uint64, ch, spatial int) error {
	fmt.Printf("program 2: two chained 1x1 convs, %d channels, %d spatial, ONE weight file at two offsets\n", ch, spatial)

	w1 := make([]float32, ch*ch)
	w2 := make([]float32, ch*ch)
	for i := range w1 {
		w1[i] = float32(i%3-1) * 0.5
		w2[i] = float32(i%4-2) * 0.25
	}

	w := mil.NewBlobWriter()
	i1 := w.AddFloat16(w1)
	i2 := w.AddFloat16(w2)
	off1, off2 := w.Offset(i1), w.Offset(i2)
	blob, err := w.Build()
	if err != nil {
		return err
	}
	fmt.Printf("  packed %d tensors into %d bytes; BLOBFILE offsets %d and %d\n", w.Count(), len(blob), off1, off2)

	// Blob Storage v2 rounds each entry up to 64 bytes. A naive packed layout
	// would put the next descriptor immediately after the previous payload, and
	// the two agree whenever every payload length is a multiple of 64 — which is
	// most tensors, so a run using one of those exercises nothing. The default
	// -multi is chosen so they disagree here.
	//
	// What that does and does not show. It shows the compiler accepting a
	// multi-entry file whose second entry sits at a rounded-up offset, so the
	// alignment padding is not rejected and the tensor read back at that offset
	// is the right one — the CPU reference check would fail otherwise. It does
	// NOT show that a packed file would be rejected: the MIL text takes its
	// offsets from the same writer that lays out the file, so a packed writer
	// would be self-consistent and the compiler would simply follow it. This
	// says which tensor placement was accepted, not which one is required.
	const header, descriptor = 64, 64
	packedOff2 := uint64(header + descriptor + len(w1)*2)
	if packedOff2 == off2 {
		fmt.Printf("  layout: entry 1 at %d, a size where the 64-aligned and packed layouts coincide\n", off2)
	} else {
		fmt.Printf("  layout: entry 1 at %d under the 64-aligned rule, where a packed layout would have put it at %d\n", off2, packedOff2)
	}

	dir, err := os.MkdirTemp("", "e5rtdispatch-multi-")
	if err != nil {
		return err
	}
	defer os.RemoveAll(dir)
	if err := writeModel(dir, twoConvMIL(ch, spatial, off1, off2), map[string][]byte{
		"weights.bin": blob,
	}); err != nil {
		return err
	}

	input := make([]float32, ch*spatial)
	for i := range input {
		input[i] = float32(i%3) * 0.25
	}
	perturbed := append([]float32(nil), input...)
	perturbed[0] += 4

	want := convReference(w2, convReference(w1, input, ch, ch, spatial), ch, ch, spatial)
	wantPerturbed := convReference(w2, convReference(w1, perturbed, ch, ch, spatial), ch, ch, spatial)

	got, stats, err := dispatch(lib, dir, mask, [][]float32{input, perturbed}, ch*spatial, 1)
	if err != nil {
		return err
	}
	stats.report()
	if err := compare(got[0], want); err != nil {
		return err
	}
	return checkResponds(got[0], got[1], wantPerturbed)
}

// twoConvMIL emits a program with two 1x1 convolutions whose weights are two
// entries in a single blob file, referenced by their metadata offsets.
func twoConvMIL(ch, spatial int, off1, off2 uint64) string {
	return fmt.Sprintf(`program(1.3)
{
    func main<ios18>(tensor<fp16, [1, %[1]d, 1, %[2]d]> x) {
        string c_pad_type = const()[name = string("c_pad_type"), val = string("valid")];
        tensor<int32, [2]> c_strides = const()[name = string("c_strides"), val = tensor<int32, [2]>([1, 1])];
        tensor<int32, [4]> c_pad = const()[name = string("c_pad"), val = tensor<int32, [4]>([0, 0, 0, 0])];
        tensor<int32, [2]> c_dilations = const()[name = string("c_dilations"), val = tensor<int32, [2]>([1, 1])];
        int32 c_groups = const()[name = string("c_groups"), val = int32(1)];
        tensor<fp16, [%[1]d, %[1]d, 1, 1]> W1 = const()[name = string("W1"), val = tensor<fp16, [%[1]d, %[1]d, 1, 1]>(BLOBFILE(path = string("@model_path/weights.bin"), offset = uint64(%[3]d)))];
        tensor<fp16, [%[1]d, %[1]d, 1, 1]> W2 = const()[name = string("W2"), val = tensor<fp16, [%[1]d, %[1]d, 1, 1]>(BLOBFILE(path = string("@model_path/weights.bin"), offset = uint64(%[4]d)))];
        tensor<fp16, [1, %[1]d, 1, %[2]d]> h = conv(dilations = c_dilations, groups = c_groups, pad = c_pad, pad_type = c_pad_type, strides = c_strides, weight = W1, x = x)[name = string("conv1")];
        tensor<fp16, [1, %[1]d, 1, %[2]d]> y = conv(dilations = c_dilations, groups = c_groups, pad = c_pad, pad_type = c_pad_type, strides = c_strides, weight = W2, x = h)[name = string("conv2")];
    } -> (y);
}
`, ch, spatial, off1, off2)
}

// writeModel lays out the directory the compiler is pointed at: model.mil at
// the top, and each weight file at the path its BLOBFILE reference names
// relative to @model_path.
func writeModel(dir, milText string, weights map[string][]byte) error {
	if err := os.WriteFile(filepath.Join(dir, "model.mil"), []byte(milText), 0o644); err != nil {
		return err
	}
	for rel, data := range weights {
		path := filepath.Join(dir, filepath.FromSlash(rel))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return err
		}
		if err := os.WriteFile(path, data, 0o644); err != nil {
			return err
		}
	}
	return nil
}

// timings records what each phase of the route cost.
type timings struct {
	compile time.Duration
	bind    time.Duration
	first   time.Duration
	steady  time.Duration
	iters   int
}

func (t timings) report() {
	fmt.Printf("  compile %v, bind %v, first dispatch %v", t.compile.Round(time.Millisecond), t.bind.Round(time.Microsecond), t.first.Round(time.Microsecond))
	if t.iters > 0 {
		fmt.Printf(", %d further dispatches %v each", t.iters, (t.steady / time.Duration(t.iters)).Round(time.Microsecond))
	}
	fmt.Println()
}

// dispatch drives the whole route for the model in dir. It compiles once and
// then evaluates every input in turn, which is the shape real use takes: the
// compile is the expensive step and the bound buffers are reused across
// dispatches. It returns one output tensor per input.
func dispatch(lib *e5rt.Lib, dir string, mask uint64, inputs [][]float32, outLen, iters int) ([][]float32, timings, error) {
	var t timings
	cacheDir := filepath.Join(dir, "cache")
	if cacheRoot != "" {
		cacheDir = filepath.Join(cacheRoot, filepath.Base(dir))
	}
	if err := os.MkdirAll(cacheDir, 0o755); err != nil {
		return nil, t, err
	}

	start := time.Now()
	config, err := lib.CompilerConfigOptionsCreate()
	if err != nil {
		return nil, t, fmt.Errorf("config create: %w", err)
	}
	defer lib.CompilerConfigOptionsRelease(config)
	if err := lib.CompilerConfigOptionsSetCacheBundleLocation(config, cacheDir); err != nil {
		return nil, t, fmt.Errorf("set cache location: %w", err)
	}

	compiler, err := lib.CompilerCreateWithConfig(config)
	if err != nil {
		return nil, t, fmt.Errorf("compiler create: %w", err)
	}
	defer lib.CompilerRelease(compiler)

	options, err := lib.CompilerOptionsCreate()
	if err != nil {
		return nil, t, fmt.Errorf("options create: %w", err)
	}
	defer lib.CompilerOptionsRelease(options)
	if err := lib.CompilerOptionsSetComputeDeviceTypesMask(options, mask); err != nil {
		return nil, t, fmt.Errorf("set device mask: %w", err)
	}
	if err := lib.CompilerOptionsSetForceRecompilation(options, true); err != nil {
		return nil, t, fmt.Errorf("set force recompilation: %w", err)
	}
	if err := lib.CompilerOptionsSetSegmenter(options, "graph"); err != nil {
		return nil, t, fmt.Errorf("set segmenter: %w", err)
	}

	library, err := lib.CompilerCompile(compiler, filepath.Join(dir, "model.mil"), options)
	if err != nil {
		return nil, t, fmt.Errorf("compile: %w", err)
	}
	defer lib.ProgramLibraryRelease(library)
	reportBackend(cacheDir, "main")

	function, err := lib.ProgramLibraryRetainProgramFunction(library, "main")
	if err != nil {
		return nil, t, fmt.Errorf("retain function main: %w", err)
	}
	defer lib.ProgramFunctionRelease(function)

	opOptions, err := lib.PrecompiledComputeOpOptionsCreate(function)
	if err != nil {
		return nil, t, fmt.Errorf("op options create: %w", err)
	}
	defer lib.PrecompiledComputeOpOptionsRelease(opOptions)
	if err := lib.PrecompiledComputeOpOptionsSetOperationName(opOptions, "main"); err != nil {
		return nil, t, fmt.Errorf("set operation name: %w", err)
	}
	if err := lib.PrecompiledComputeOpOptionsSetAllocateIntermediateBuffers(opOptions, true); err != nil {
		return nil, t, fmt.Errorf("set allocate intermediate buffers: %w", err)
	}

	op, err := lib.OperationCreatePrecompiled(opOptions)
	if err != nil {
		return nil, t, fmt.Errorf("operation create: %w", err)
	}
	defer lib.OperationRelease(op)
	t.compile = time.Since(start)

	start = time.Now()
	inBuf, inPtr, err := bindPort(lib, op, "x", len(inputs[0])*2, true)
	if err != nil {
		return nil, t, err
	}
	defer lib.BufferObjectRelease(inBuf)
	outBuf, outPtr, err := bindPort(lib, op, "y", outLen*2, false)
	if err != nil {
		return nil, t, err
	}
	defer lib.BufferObjectRelease(outBuf)

	stream, err := lib.ExecutionStreamCreate()
	if err != nil {
		return nil, t, fmt.Errorf("stream create: %w", err)
	}
	defer lib.ExecutionStreamRelease(stream)
	if err := lib.EncodeOperation(stream, op); err != nil {
		return nil, t, fmt.Errorf("encode: %w", err)
	}
	t.bind = time.Since(start)

	outputs := make([][]float32, 0, len(inputs))
	for n, input := range inputs {
		writeFP16(inPtr, input)
		start = time.Now()
		if err := lib.ExecuteSync(stream); err != nil {
			return nil, t, fmt.Errorf("execute (input %d): %w", n, err)
		}
		if n == 0 {
			t.first = time.Since(start)
		}
		outputs = append(outputs, readFP16(outPtr, outLen))
	}

	if iters > 0 {
		start = time.Now()
		for range iters {
			if err := lib.ExecuteSync(stream); err != nil {
				return nil, t, fmt.Errorf("execute (steady state): %w", err)
			}
		}
		t.steady = time.Since(start)
		t.iters = iters
	}

	return outputs, t, nil
}

// cacheRoot, when set by -cache, keeps compiled bundles for inspection instead
// of discarding them with the model directory.
var cacheRoot string

// reportBackend prints which backend the compiler actually chose for fn.
//
// The device mask is a permission, not a placement: with the Neural Engine bit
// set the planner may still fall back to another backend, so the mask alone is
// no evidence the program ran on the engine. The compiled bundle names the
// decision. Inside it each function gets one directory per selected backend,
// named <function>_<backend>, and these were observed on macOS 26.x:
//
//	main_ane        model.anehash                    Neural Engine
//	main_bnns       bnns_program.bnnsir              CPU
//	main_mps_graph  main_mps_graph.mpsgraphpackage   GPU
//
// Those three were each produced by compiling this same program under -device
// ane, cpu, and gpu, so the reading is anchored at both ends rather than
// assumed. Nothing is inferred when no bundle is found: it says UNKNOWN.
//
// ANEForge documents a different oracle for this, a SelectedBackend entry in
// the bundle's analytics.mil. No analytics.mil appears in the bundles produced
// here, so that form is either older or written only under a compiler option
// this program does not set.
func reportBackend(cacheDir, fn string) {
	seen := map[string]bool{}
	filepath.WalkDir(cacheDir, func(path string, d os.DirEntry, err error) error {
		if err != nil || !d.IsDir() {
			return nil
		}
		// Only the function directory's direct children name a backend. A
		// deeper match would also catch payloads like
		// main_mps_graph/main_mps_graph.mpsgraphpackage, which is itself a
		// directory.
		if filepath.Base(filepath.Dir(path)) != fn {
			return nil
		}
		if suffix, ok := strings.CutPrefix(d.Name(), fn+"_"); ok {
			seen[suffix] = true
		}
		return nil
	})
	if len(seen) == 0 {
		fmt.Printf("  backend: UNKNOWN (no %s_* directory in the compiled bundle)\n", fn)
		return
	}
	names := slices.Sorted(maps.Keys(seen))
	fmt.Printf("  backend: the compiler emitted %v for %q\n", names, fn)
}

// bindPort retains a named port, allocates a CPU-visible buffer object for it,
// and binds the two together. It returns the buffer and its data pointer.
func bindPort(lib *e5rt.Lib, op uintptr, name string, nbytes int, input bool) (uintptr, uintptr, error) {
	var port uintptr
	var err error
	if input {
		port, err = lib.OperationRetainInputPort(op, name)
	} else {
		port, err = lib.OperationRetainOutputPort(op, name)
	}
	if err != nil {
		return 0, 0, fmt.Errorf("retain port %q: %w", name, err)
	}
	defer lib.IOPortRelease(port)

	// The reference implementation rounds every allocation up to a multiple of
	// 64 bytes and never allocates less than that.
	size := (nbytes + 63) &^ 63
	if size < 64 {
		size = 64
	}
	buf, err := lib.BufferObjectAlloc(uintptr(size), 0)
	if err != nil {
		return 0, 0, fmt.Errorf("alloc buffer for %q: %w", name, err)
	}
	ptr, err := lib.BufferObjectGetDataPtr(buf)
	if err != nil {
		lib.BufferObjectRelease(buf)
		return 0, 0, fmt.Errorf("data pointer for %q: %w", name, err)
	}
	if err := lib.IOPortBindBufferObject(port, buf); err != nil {
		lib.BufferObjectRelease(buf)
		return 0, 0, fmt.Errorf("bind buffer to %q: %w", name, err)
	}
	return buf, ptr, nil
}

// pointerAt converts a foreign address to an unsafe.Pointer. Converting a
// uintptr directly is what go vet flags as a possible misuse; taking the
// address of the uintptr yields a pointer the checker accepts, and the address
// is one the Go collector does not own in either spelling.
func pointerAt(addr uintptr) unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&addr))
}

func writeFP16(dst uintptr, src []float32) {
	out := unsafe.Slice((*uint16)(pointerAt(dst)), len(src))
	for i, v := range src {
		out[i] = ane.Float32ToFP16(v)
	}
}

func readFP16(src uintptr, n int) []float32 {
	in := unsafe.Slice((*uint16)(pointerAt(src)), n)
	out := make([]float32, n)
	for i, v := range in {
		out[i] = ane.FP16ToFloat32(v)
	}
	return out
}

// checkResponds is the negative control for compare. A run that returned zeros,
// a constant, or a stale buffer would satisfy compare only by coincidence, but
// it cannot also track a changed input. So the second dispatch must differ from
// the first, and must match its own reference.
func checkResponds(base, perturbed, wantPerturbed []float32) error {
	var moved float64
	for i := range base {
		if d := math.Abs(float64(base[i]) - float64(perturbed[i])); d > moved {
			moved = d
		}
	}
	if moved == 0 {
		return fmt.Errorf("control failed: perturbing the input changed no output element, so the engine is not reading the bound input buffer")
	}
	if err := compare(perturbed, wantPerturbed); err != nil {
		return fmt.Errorf("control failed: %w", err)
	}
	fmt.Printf("  control: perturbing one input element moved the output by %g and the new output still matches its reference\n", moved)
	return nil
}

// convReference computes the 1x1 convolution in float64 on the CPU:
// y[o][s] = sum over i of w[o][i] * x[i][s].
func convReference(w, x []float32, outCh, inCh, spatial int) []float32 {
	y := make([]float32, outCh*spatial)
	for o := range outCh {
		for s := range spatial {
			var acc float64
			for i := range inCh {
				acc += float64(w[o*inCh+i]) * float64(x[i*spatial+s])
			}
			y[o*spatial+s] = float32(acc)
		}
	}
	return y
}

// compare reports the largest disagreement between the engine's output and the
// CPU reference. The tolerance is fp16 scale: the engine stores and accumulates
// in half precision, so an exact match is not the right expectation and
// demanding one would only produce a false failure.
func compare(got, want []float32) error {
	if len(got) != len(want) {
		return fmt.Errorf("output length %d, want %d", len(got), len(want))
	}
	var worst float64
	var at int
	var scale float64
	for i := range got {
		d := math.Abs(float64(got[i]) - float64(want[i]))
		if d > worst {
			worst, at = d, i
		}
		if m := math.Abs(float64(want[i])); m > scale {
			scale = m
		}
	}
	tol := 0.01 * math.Max(scale, 1)
	if worst > tol {
		return fmt.Errorf("output disagrees with the CPU reference: worst |diff| %g at element %d (got %g, want %g), tolerance %g",
			worst, at, got[at], want[at], tol)
	}
	fmt.Printf("  output matches the CPU reference over %d elements: worst |diff| %g, tolerance %g, reference peak |%g|\n", len(got), worst, tol, scale)
	return nil
}
