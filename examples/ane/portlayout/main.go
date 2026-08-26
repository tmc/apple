// Command portlayout sweeps ordinary input and output E5RT ports across
// tensor shapes with the same self-indexing technique that mapped the state
// buffer layout, and checks the claim that only the state carries padding.
//
// The padding finding came from examples/ane/statecache: each channel of a
// bound state starts on a 64-byte boundary, so the packed element count lies
// about the footprint. Ordinary ports were asserted packed from a narrow set
// of shapes exercised by passing programs. This command widens the observation:
// for every swept shape it binds oversized buffers to an identity program's
// input and output ports, fills the input with slot-index marks, executes, and
// reads back which slot produced each output value.
//
// # What each check establishes, per shape
//
//   - Packed round trip. Output value i equals input mark i exactly, in fp16,
//     which represents integers exactly up to 2048 — hence the sweep guard. A
//     mismatch names the first element and the stride error it implies.
//
//   - No over-read. Every returned mark was written inside the input's logical
//     size. An engine that read past the declared port size would fetch tail
//     marks whose indices exceed the element count, and the round trip above
//     would catch them as wrong values; the tail is filled with continuing
//     indices precisely so that such a fetch cannot pass silently.
//
//   - No over-write. Everything past the output buffer's logical size carries
//     a position-dependent poison byte pattern that must survive execution.
//     This is the same instrument x/ane/e5rt applies to state ports
//     (TestStatePoisonTail): E5RT exposes no way to place a guard page after
//     its own allocation, so detection happens at check time, not fault time.
//
//   - A refusal control. Retaining a port name the compiled function does not
//     declare must fail, so that a successful bind means resolution rather
//     than indifference.
//
// # Scope
//
// The sweep covers channel and spatial extents including odd and non-multiple
// values, plus one two-batch shape. Element type is fp16 throughout, because
// that is what every ANE-bound program here uses. Ranks other than 4 are not
// emitted by any MIL generator in this module and stay unmeasured, as do fp32
// ports; a shape this command refuses to build is reported as such rather than
// quietly dropped from the claim.
package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unsafe"

	"github.com/tmc/apple/x/ane"
	"github.com/tmc/apple/x/ane/e5rt"
	"github.com/tmc/apple/x/ane/mil"
)

var (
	keepGoing = flag.Bool("keep-going", false, "report every failing shape instead of stopping at the first")
)

// A sweptShape names one shape to measure. Batch is the leading extent; the
// MIL text below emits [Batch, Channels, 1, Spatial].
type sweptShape struct {
	batch, channels, spatial int
}

// sweepShapes deliberately mixes even, odd, prime, and 64-multiple extents,
// and includes shapes where channels*spatial lands on either side of a
// 64-byte row boundary.
var sweepShapes = []sweptShape{
	{1, 1, 1},
	{1, 1, 3},
	{1, 3, 5},
	{1, 4, 4},
	{1, 4, 33},
	{1, 8, 2},
	{1, 7, 64},
	{1, 2, 127},
	{1, 16, 17},
	{1, 5, 129},
	{2, 3, 7},
}

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "FAIL:", err)
		os.Exit(1)
	}
	fmt.Println("\nOK")
}

func run() error {
	lib, err := e5rt.Open()
	if err != nil {
		return fmt.Errorf("open e5rt: %w", err)
	}
	root, err := os.MkdirTemp("", "portlayout")
	if err != nil {
		return err
	}
	defer os.RemoveAll(root)

	var failures []error
	measured := 0
	for _, s := range sweepShapes {
		n := s.batch * s.channels * s.spatial
		if n > 2048 {
			return fmt.Errorf("shape %v has %d elements, past the %d fp16 indexes exactly", s, n, 2048)
		}
		err := measure(lib, filepath.Join(root, fmt.Sprintf("b%dc%dd%d", s.batch, s.channels, s.spatial)), s)
		if err == nil {
			measured++
			continue
		}
		if *keepGoing {
			failures = append(failures, fmt.Errorf("shape [1..%d, %d, 1, %d]: %w", s.batch, s.channels, s.spatial, err))
			fmt.Printf("  FAIL\n    %v\n", err)
			continue
		}
		return fmt.Errorf("shape [1..%d, %d, 1, %d]: %w", s.batch, s.channels, s.spatial, err)
	}
	if len(failures) > 0 {
		return errors.Join(append([]error{fmt.Errorf("%d of %d shapes failed", len(failures), len(sweepShapes))}, failures...)...)
	}
	fmt.Printf("\nmeasured %d shapes: ordinary ports are packed at every one, with no over-read and no past-the-end write\n", measured)
	return nil
}

// measure compiles an identity program for s and checks the three properties
// the package comment lists.
func measure(lib *e5rt.Lib, dir string, s sweptShape) error {
	n := s.batch * s.channels * s.spatial
	fmt.Printf("[batch=%d chan=%d spatial=%d] (%d values) ", s.batch, s.channels, s.spatial, n)

	if err := os.MkdirAll(filepath.Join(dir, "weights"), 0o755); err != nil {
		return err
	}
	weights := make([]float32, s.channels*s.channels)
	for i := range s.channels {
		weights[i*s.channels+i] = 1
	}
	blob, err := mil.BuildWeightBlob(weights, s.channels, s.channels)
	if err != nil {
		return err
	}
	if err := os.WriteFile(filepath.Join(dir, "weights", "weight.bin"), blob, 0o644); err != nil {
		return err
	}
	modelPath := filepath.Join(dir, "model.mil")
	if err := os.WriteFile(modelPath, []byte(genIdentity(s.batch, s.channels, s.spatial)), 0o644); err != nil {
		return err
	}

	config, err := lib.CompilerConfigOptionsCreate()
	if err != nil {
		return err
	}
	defer lib.CompilerConfigOptionsRelease(config)
	if err := lib.CompilerConfigOptionsSetCacheBundleLocation(config, dir); err != nil {
		return err
	}
	compiler, err := lib.CompilerCreateWithConfig(config)
	if err != nil {
		return err
	}
	defer lib.CompilerRelease(compiler)
	options, err := lib.CompilerOptionsCreate()
	if err != nil {
		return err
	}
	defer lib.CompilerOptionsRelease(options)
	if err := lib.CompilerOptionsSetComputeDeviceTypesMask(options, e5rt.ComputeDeviceANE); err != nil {
		return err
	}
	library, err := lib.CompilerCompile(compiler, modelPath, options)
	if err != nil {
		return fmt.Errorf("compile: %w", err)
	}
	defer lib.ProgramLibraryRelease(library)
	if err := requireANE(dir); err != nil {
		return err
	}
	function, err := lib.ProgramLibraryRetainProgramFunction(library, "main")
	if err != nil {
		return fmt.Errorf("retain function: %w", err)
	}
	defer lib.ProgramFunctionRelease(function)
	opOptions, err := lib.PrecompiledComputeOpOptionsCreate(function)
	if err != nil {
		return err
	}
	defer lib.PrecompiledComputeOpOptionsRelease(opOptions)
	if err := lib.PrecompiledComputeOpOptionsSetOperationName(opOptions, "main"); err != nil {
		return err
	}
	if err := lib.PrecompiledComputeOpOptionsSetAllocateIntermediateBuffers(opOptions, true); err != nil {
		return err
	}
	op, err := lib.OperationCreatePrecompiled(opOptions)
	if err != nil {
		return fmt.Errorf("create operation: %w", err)
	}
	defer lib.OperationRelease(op)

	logicalBytes := n * 2
	tailBytes := 256

	// Input: logical region holds mark i at slot i; the tail continues the
	// index sequence so a read past the declared size returns an index no
	// correct output would carry.
	inBuf, inData, err := allocOversized(lib, logicalBytes, tailBytes)
	if err != nil {
		return fmt.Errorf("allocate input: %w", err)
	}
	defer lib.BufferObjectRelease(inBuf)
	marks := make([]float32, (logicalBytes+tailBytes)/2)
	for i := range marks {
		marks[i] = float32(i)
	}
	writeSlots(inData, marks)

	// Output: logical region zeroed; the tail carries a position-dependent
	// poison pattern that must survive the run.
	outBuf, outData, err := allocOversized(lib, logicalBytes, tailBytes)
	if err != nil {
		return fmt.Errorf("allocate output: %w", err)
	}
	defer lib.BufferObjectRelease(outBuf)
	clear(outData[:logicalBytes])
	poison(outData[logicalBytes:])
	wantTail := append([]byte(nil), outData[logicalBytes:]...)

	inPort, err := lib.OperationRetainInputPort(op, "x")
	if err != nil {
		return fmt.Errorf("retain input: %w", err)
	}
	defer lib.IOPortRelease(inPort)
	if err := lib.IOPortBindBufferObject(inPort, inBuf); err != nil {
		return fmt.Errorf("bind input: %w", err)
	}
	outPort, err := lib.OperationRetainOutputPort(op, "y")
	if err != nil {
		return fmt.Errorf("retain output: %w", err)
	}
	defer lib.IOPortRelease(outPort)
	if err := lib.IOPortBindBufferObject(outPort, outBuf); err != nil {
		return fmt.Errorf("bind output: %w", err)
	}

	// The refusal control: a name the function does not declare must not
	// resolve, so the successful binds above mean something.
	if _, err := lib.OperationRetainInputPort(op, "no_such_port"); err == nil {
		return errors.New(`retaining an input port named "no_such_port" succeeded`)
	}

	stream, err := lib.ExecutionStreamCreate()
	if err != nil {
		return err
	}
	defer lib.ExecutionStreamRelease(stream)
	if err := lib.EncodeOperation(stream, op); err != nil {
		return fmt.Errorf("encode: %w", err)
	}
	if err := lib.ExecuteSync(stream); err != nil {
		return fmt.Errorf("execute: %w", err)
	}

	got := readSlots(outData[:logicalBytes], n)
	for i, v := range got {
		if v != float32(i) {
			elem := i % (s.channels * s.spatial)
			strideErr := v - float32(i)
			return fmt.Errorf("output element %d (channel %d, position %d) carried mark %.0f instead of %.0f (stride offset %.0f slots)",
				i, elem/s.spatial, elem%s.spatial, v, float32(i), strideErr)
		}
	}
	if !equalBytes(outData[logicalBytes:], wantTail) {
		for i := range outData[logicalBytes:] {
			if outData[logicalBytes+i] != wantTail[i] {
				return fmt.Errorf("the engine wrote past the output port's logical size: first changed tail byte at +%d",
					i)
			}
		}
		return errors.New("the output buffer's poison tail changed without naming an offset")
	}
	fmt.Println("packed round trip, no over-read, tail intact")
	return nil
}

// genIdentity emits an identity map as a 1x1 convolution with identity weights,
// parameterized by batch — the generators in x/ane/mil fix the leading extent
// at 1, and this sweep needs a two-batch shape.
func genIdentity(batch, channels, spatial int) string {
	return fmt.Sprintf(`program(1.3)
%[1]s
{
    func main<ios18>(tensor<fp16, [%[2]d, %[3]d, 1, %[4]d]> x) {
        string c_pad_type = const()[name = string("c_pad_type"), val = string("valid")];
        tensor<int32, [2]> c_strides = const()[name = string("c_strides"), val = tensor<int32, [2]>([1, 1])];
        tensor<int32, [4]> c_pad = const()[name = string("c_pad"), val = tensor<int32, [4]>([0, 0, 0, 0])];
        tensor<int32, [2]> c_dilations = const()[name = string("c_dilations"), val = tensor<int32, [2]>([1, 1])];
        int32 c_groups = const()[name = string("c_groups"), val = int32(1)];
        tensor<fp16, [%[3]d, %[3]d, 1, 1]> W = const()[name = string("W"), val = tensor<fp16, [%[3]d, %[3]d, 1, 1]>(BLOBFILE(path = string("@model_path/weights/weight.bin"), offset = uint64(64)))];
        tensor<fp16, [%[2]d, %[3]d, 1, %[4]d]> y = conv(dilations = c_dilations, groups = c_groups, pad = c_pad, pad_type = c_pad_type, strides = c_strides, weight = W, x = x)[name = string("conv")];
    } -> (y);
}
`, mil.BuildInfo,
		batch, channels, spatial,
	)
}

// allocOversized returns a buffer object of nbytes plus a tail, together with
// the whole host-visible storage.
func allocOversized(lib *e5rt.Lib, nbytes, tail int) (uintptr, []byte, error) {
	obj, err := lib.BufferObjectAlloc(uintptr(nbytes+tail), 0)
	if err != nil {
		return 0, nil, err
	}
	ptr, err := lib.BufferObjectGetDataPtr(obj)
	if err != nil {
		lib.BufferObjectRelease(obj)
		return 0, nil, err
	}
	data := unsafe.Slice((*byte)(pointerAt(ptr)), nbytes+tail)
	return obj, data, nil
}

func pointerAt(addr uintptr) unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&addr))
}

func writeSlots(data []byte, v []float32) {
	for i, x := range v {
		bits := ane.Float32ToFP16(x)
		data[2*i] = byte(bits)
		data[2*i+1] = byte(bits >> 8)
	}
}

func readSlots(data []byte, n int) []float32 {
	out := make([]float32, n)
	for i := range out {
		out[i] = ane.FP16ToFloat32(uint16(data[2*i]) | uint16(data[2*i+1])<<8)
	}
	return out
}

func poison(tail []byte) {
	for i := range tail {
		tail[i] = byte(i*31 + 7)
	}
}

func equalBytes(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// requireANE fails unless the compiler left an ANE artifact under dir.
func requireANE(cacheDir string) error {
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
	if !seen["ane"] {
		names := make([]string, 0, len(seen))
		for name := range seen {
			names = append(names, name)
		}
		return fmt.Errorf("the compiler emitted %v, not an ANE artifact", names)
	}
	return nil
}
