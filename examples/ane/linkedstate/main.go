// Command linkedstate measures what happens when retained E5RT state flows
// through a pipeline-style link: two separately compiled functions encoded in
// one stream whose state inout ports share one buffer object.
//
// e5rt.Pipeline links an earlier stage's output buffer to a later stage's input
// by binding both ports to the same buffer object. Nothing exercises the same
// sharing for state ports, and the state is exactly where a silent corruption
// hides: examples/ane/statecache found that each channel of a bound state
// starts on a 64-byte boundary, so two functions declaring different shapes can
// share a buffer while addressing different offsets for the same element. The
// ordinary link path shares buffers by construction, which means neither the
// Go side nor a size check sees the disagreement — only the engine's addresses
// matter.
//
// This command binds the shared buffer by hand, the way the pipeline would,
// and measures three things.
//
// # Arms
//
//   - Baseline. An accumulate-state function and a separately compiled reader
//     of the same shape share the state buffer inside one stream. Each
//     synchronous execution advances the state by one update and the reader
//     reports it; the sequence must follow a float64 recurrence exactly, in
//     fp16-exact arithmetic, checked after every step. This establishes that
//     the composition itself behaves before anything leans on it.
//
//   - Layout oracle. With the shared buffer filled with self-indexing marks, a
//     reader reports which host slot each element came from. The result must
//     equal the [e5rt.StateLayout] prediction — the documented layout has to
//     hold for a link-shared binding, not only for a singly-bound one.
//
//   - Mismatch control. An accumulator of shape [1, 2, 1, 8] and a reader of
//     shape [1, 4, 1, 4] share the buffer. Both layouts predict the same
//     physical addresses for every element position where their views overlap:
//     channel c starts at byte c*64 in either shape, so the reader should see
//     the writer's first two channels laid out flat across four of its own, and
//     zeros wherever the writer never wrote. If the runtime instead normalized,
//     padded, or reordered the shared state, the observed values would diverge
//     from that prediction and the command fails — silently corrupting is the
//     failure mode this whole program exists to rule out.
//
// # What is claimed and what is not
//
// Measured here: fp16 states on macOS on this machine, rank-4 shapes with a
// leading extent of 1, and sharing implemented as two operations in one stream.
// Element types other than fp16, other ranks, and any behavior across separate
// streams are outside the claim. The poison tail on every shared buffer keeps
// any past-the-end engine write observable even where a check above does not
// name it.
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
	steps = flag.Int("steps", 4, "number of updates the baseline applies (2..8, to keep the recurrence exact in fp16)")
)

const stateName = "kv"

func main() {
	flag.Parse()
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, "FAIL:", err)
		os.Exit(1)
	}
	fmt.Println("\nOK")
}

func run() error {
	if *steps < 2 || *steps > 8 {
		return errors.New("-steps must be between 2 and 8 to keep the recurrence exact in fp16")
	}
	lib, err := e5rt.Open()
	if err != nil {
		return fmt.Errorf("open e5rt: %w", err)
	}
	root, err := os.MkdirTemp("", "linkedstate")
	if err != nil {
		return err
	}
	defer os.RemoveAll(root)

	fmt.Println("arm 1: baseline — accumulate and read share one state buffer in one stream")
	if err := baselineArm(lib, filepath.Join(root, "base")); err != nil {
		return fmt.Errorf("baseline: %w", err)
	}

	fmt.Println("\narm 2: layout oracle — which slots does a reader reach through the shared binding")
	if err := oracleArm(lib, filepath.Join(root, "oracle")); err != nil {
		return fmt.Errorf("layout oracle: %w", err)
	}

	fmt.Println("\narm 3: mismatch control — writer [1,2,1,8] against reader [1,4,1,4]")
	if err := mismatchArm(lib, filepath.Join(root, "mismatch")); err != nil {
		return fmt.Errorf("mismatch control: %w", err)
	}
	return nil
}

// linkedStages holds two compiled functions whose state ports are bound to one
// shared buffer, plus the stream they are encoded in.
type linkedStages struct {
	lib    *e5rt.Lib
	stream uintptr
	writer *stage
	reader *stage
}

type stage struct {
	op      uintptr
	port    uintptr // retained inout port for the shared state
	valueIn uintptr // buffer object for "value", or zero
	valueP  []byte
	outP    []byte
	closes  []func() error
}

func (s *stage) close() error {
	var errs []error
	for i := len(s.closes) - 1; i >= 0; i-- {
		errs = append(errs, s.closes[i]())
	}
	return errors.Join(errs...)
}

func (l *linkedStages) close() error {
	return errors.Join(
		releaseStream(l.lib, l.stream),
		l.writer.close(),
		l.reader.close(),
	)
}

func releaseStream(lib *e5rt.Lib, stream uintptr) error {
	if stream == 0 {
		return nil
	}
	return lib.ExecutionStreamRelease(stream)
}

// openStage compiles modelText into dir, creates the operation, and binds its
// state inout port to shared through the checked path using lay.
func openStage(lib *e5rt.Lib, dir, modelText string, shared *sharedBuffer, lay e5rt.StateLayout, withValue bool) (_ *stage, err error) {
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	modelPath := filepath.Join(dir, "model.mil")
	if err := os.WriteFile(modelPath, []byte(modelText), 0o644); err != nil {
		return nil, err
	}
	s := &stage{}
	defer func() {
		if err != nil {
			err = errors.Join(err, s.close())
		}
	}()

	config, err := lib.CompilerConfigOptionsCreate()
	if err != nil {
		return nil, err
	}
	s.closes = append(s.closes, func() error { return lib.CompilerConfigOptionsRelease(config) })
	if err := lib.CompilerConfigOptionsSetCacheBundleLocation(config, dir); err != nil {
		return nil, err
	}
	compiler, err := lib.CompilerCreateWithConfig(config)
	if err != nil {
		return nil, err
	}
	s.closes = append(s.closes, func() error { return lib.CompilerRelease(compiler) })
	options, err := lib.CompilerOptionsCreate()
	if err != nil {
		return nil, err
	}
	s.closes = append(s.closes, func() error { return lib.CompilerOptionsRelease(options) })
	if err := lib.CompilerOptionsSetComputeDeviceTypesMask(options, e5rt.ComputeDeviceANE); err != nil {
		return nil, err
	}
	library, err := lib.CompilerCompile(compiler, modelPath, options)
	if err != nil {
		return nil, fmt.Errorf("compile: %w", err)
	}
	if err := requireANE(dir); err != nil {
		return nil, err
	}
	s.closes = append(s.closes, func() error { return lib.ProgramLibraryRelease(library) })
	function, err := lib.ProgramLibraryRetainProgramFunction(library, "main")
	if err != nil {
		return nil, fmt.Errorf("retain function: %w", err)
	}
	s.closes = append(s.closes, func() error { return lib.ProgramFunctionRelease(function) })
	opOptions, err := lib.PrecompiledComputeOpOptionsCreate(function)
	if err != nil {
		return nil, err
	}
	s.closes = append(s.closes, func() error { return lib.PrecompiledComputeOpOptionsRelease(opOptions) })
	if err := lib.PrecompiledComputeOpOptionsSetOperationName(opOptions, "main"); err != nil {
		return nil, err
	}
	if err := lib.PrecompiledComputeOpOptionsSetAllocateIntermediateBuffers(opOptions, true); err != nil {
		return nil, err
	}
	if s.op, err = lib.OperationCreatePrecompiled(opOptions); err != nil {
		return nil, fmt.Errorf("create operation: %w", err)
	}
	s.closes = append(s.closes, func() error { return lib.OperationRelease(s.op) })

	if s.port, err = lib.BindStatePort(s.op, stateName, shared.obj, lay); err != nil {
		return nil, fmt.Errorf("bind state %q: %w", stateName, err)
	}
	s.closes = append(s.closes, func() error { return lib.IOPortRelease(s.port) })

	if withValue {
		inPort, err := lib.OperationRetainInputPort(s.op, "value")
		if err != nil {
			return nil, fmt.Errorf("retain value input: %w", err)
		}
		s.closes = append(s.closes, func() error { return lib.IOPortRelease(inPort) })
		obj, data, err := allocOversized(lib, lay.Values()*2, tailPad)
		if err != nil {
			return nil, fmt.Errorf("allocate value buffer: %w", err)
		}
		s.closes = append(s.closes, func() error { return lib.BufferObjectRelease(obj) })
		clear(data[:lay.Values()*2])
		poison(data[lay.Values()*2:])
		if err := lib.IOPortBindBufferObject(inPort, obj); err != nil {
			return nil, fmt.Errorf("bind value buffer: %w", err)
		}
		s.valueIn, s.valueP = obj, data
	}

	outPort, err := lib.OperationRetainOutputPort(s.op, "y")
	if err != nil {
		return nil, fmt.Errorf("retain y output: %w", err)
	}
	s.closes = append(s.closes, func() error { return lib.IOPortRelease(outPort) })
	obj, data, err := allocOversized(lib, lay.Values()*2, tailPad)
	if err != nil {
		return nil, fmt.Errorf("allocate y buffer: %w", err)
	}
	s.closes = append(s.closes, func() error { return lib.BufferObjectRelease(obj) })
	clear(data[:lay.Values()*2])
	poison(data[lay.Values()*2:])
	if err := lib.IOPortBindBufferObject(outPort, obj); err != nil {
		return nil, fmt.Errorf("bind y buffer: %w", err)
	}
	s.outP = data
	return s, nil
}

// PrecomputedComputeOpOptionsReleaseSafe exists to keep release ordering next
// to the option creation site without capturing err shadows in openStage.
type noop struct{}

// A sharedBuffer is one buffer object serving as the linked state for several
// compiled functions, with a poisoned tail past its logical size.
type sharedBuffer struct {
	obj     uintptr
	data    []byte
	logical int
}

const tailPad = 256

func allocShared(lib *e5rt.Lib, nbytes int) (*sharedBuffer, error) {
	obj, err := lib.BufferObjectAlloc(uintptr(nbytes+tailPad), 0)
	if err != nil {
		return nil, err
	}
	ptr, err := lib.BufferObjectGetDataPtr(obj)
	if err != nil {
		lib.BufferObjectRelease(obj)
		return nil, err
	}
	data := unsafe.Slice((*byte)(pointerAt(ptr)), nbytes+tailPad)
	return &sharedBuffer{obj: obj, data: data, logical: nbytes}, nil
}

func (b *sharedBuffer) refillTail() { poison(b.data[b.logical:]) }

func (b *sharedBuffer) tailIntact() (bool, int) {
	for i := b.logical; i < len(b.data); i++ {
		if want := byte(i*31 + 7); b.data[i] != want {
			return false, i - b.logical
		}
	}
	return true, 0
}

func (b *sharedBuffer) close(lib *e5rt.Lib) error {
	if b == nil || b.obj == 0 {
		return nil
	}
	err := lib.BufferObjectRelease(b.obj)
	b.obj, b.data = 0, nil
	return err
}

// baselineArm composes an accumulator and a same-shape reader in one stream
// over one shared state buffer, and requires an exact fp16 recurrence visible
// through the reader after every execution.
func baselineArm(lib *e5rt.Lib, root string) error {
	shape := [4]int{1, 4, 1, 4}
	lay := e5rt.StateLayout{Channels: 4, Dim: 4}
	n := lay.Values()

	shared, err := allocShared(lib, lay.Size())
	if err != nil {
		return err
	}
	defer shared.close(lib)
	clear(shared.data[:lay.Size()])
	shared.refillTail()

	l, err := compose(lib, root, shape, shape, shared,
		mil.GenAccumulateState(stateName, shape),
		mil.GenReadState(stateName, shape))
	if err != nil {
		return err
	}
	defer l.close()

	sum := make([]float64, n)
	for t := range *steps {
		v := make([]float32, n)
		for i := range v {
			v[i] = float32((t*3+i*5)%7-2) * 0.5
		}
		writeFP16(l.writer.valueP[:n*2], v)
		if err := lib.ExecuteSync(l.stream); err != nil {
			return fmt.Errorf("execute step %d: %w", t, err)
		}
		for i := range sum {
			sum[i] += float64(v[i])
		}
		got := readFP16(l.reader.outP[:n*2], n)
		want := toFloat32(sum)
		for i := range got {
			if got[i] != want[i] {
				return fmt.Errorf("step %d element %d: reader reported %v, the exact recurrence says %v", t, i, got[i], want[i])
			}
		}
	}
	if ok, at := shared.tailIntact(); !ok {
		return fmt.Errorf("the engine wrote past the shared state's logical size (tail byte +%d changed)", at)
	}
	// The reverse direction: the engine's writes must be in the shared buffer
	// itself, not in some private copy the reader happens to agree with.
	direct := readTensorPadded(shared.data, lay)
	want := toFloat32(sum)
	for i := range direct {
		if direct[i] != want[i] {
			return fmt.Errorf("reading the shared buffer directly: element %d is %v, want %v", i, direct[i], want[i])
		}
	}
	fmt.Printf("  %d executions advanced the state through the link, each matching the float64 recurrence exactly\n", *steps)
	fmt.Println("  the engine's writes landed in the shared buffer, and its poison tail survived")
	return nil
}

// oracleArm fills the shared buffer with self-indexing marks and requires a
// reader to fetch exactly the slots StateLayout predicts for its shape.
func oracleArm(lib *e5rt.Lib, root string) error {
	shape := [4]int{1, 4, 1, 4}
	lay := e5rt.StateLayout{Channels: 4, Dim: 4}
	n := lay.Values()

	// Generous buffer: four times any plausible row, so a surprise stride is
	// visible rather than running off the end.
	buf, err := allocShared(lib, lay.Channels*4*max((lay.Dim*2+63)&^63, 64))
	if err != nil {
		return err
	}
	defer buf.close(lib)
	slots := len(buf.data) / 2
	if slots > 2048 {
		return fmt.Errorf("%d slots exceed what fp16 indexes exactly", slots)
	}
	marks := make([]float32, slots)
	for i := range marks {
		marks[i] = float32(i)
	}
	writeFP16(buf.data, marks)

	r, err := openStage(lib, filepath.Join(root, "reader"), mil.GenReadState(stateName, shape), buf, lay, false)
	if err != nil {
		return err
	}
	defer r.close()
	stream, err := lib.ExecutionStreamCreate()
	if err != nil {
		return err
	}
	defer lib.ExecutionStreamRelease(stream)
	if err := lib.EncodeOperation(stream, r.op); err != nil {
		return err
	}
	if err := lib.ExecuteSync(stream); err != nil {
		return fmt.Errorf("execute: %w", err)
	}
	got := readFP16(r.outP[:n*2], n)
	bad := false
	for c := range lay.Channels {
		row := got[c*lay.Dim : (c+1)*lay.Dim]
		fmt.Printf("  channel %d read from slots %v\n", c, row)
		for i, v := range row {
			if v != float32(lay.Offset(c, i)/2) {
				bad = true
			}
		}
	}
	if bad {
		return errors.New("the reader did not follow the StateLayout offsets through the shared binding; " +
			"a link changes how state addresses resolve, and every caller needs to know that")
	}
	fmt.Printf("  every element came from the slot StateLayout predicts (%d-byte rows)\n", lay.RowBytes())
	return nil
}

// mismatchArm shares one buffer between an accumulator of shape [1, 2, 1, 8]
// and a reader of shape [1, 4, 1, 4]. Both layouts put channel starts at the
// same bytes, so the overlap must read through exactly and the untouched
// remainder must stay zero.
func mismatchArm(lib *e5rt.Lib, root string) error {
	wShape := [4]int{1, 2, 1, 8}
	rShape := [4]int{1, 4, 1, 4}
	wLay := e5rt.StateLayout{Channels: 2, Dim: 8}
	rLay := e5rt.StateLayout{Channels: 4, Dim: 4}
	size := max(wLay.Size(), rLay.Size())

	shared, err := allocShared(lib, size)
	if err != nil {
		return err
	}
	defer shared.close(lib)
	clear(shared.data[:size])
	shared.refillTail()

	l, err := composeShapes(lib, root, wShape, rShape, wLay, rLay, shared)
	if err != nil {
		return err
	}
	defer l.close()

	// Distinct, exactly-representable values: element (c, d) gets 1 + c*8 + d.
	wv := make([]float32, wLay.Values())
	for c := range wLay.Channels {
		for d := range wLay.Dim {
			wv[c*wLay.Dim+d] = float32(1 + c*wLay.Dim + d)
		}
	}
	writeFP16(l.writer.valueP[:wLay.Values()*2], wv)
	if err := lib.ExecuteSync(l.stream); err != nil {
		return fmt.Errorf("execute: %w", err)
	}

	got := readFP16(l.reader.outP[:rLay.Values()*2], rLay.Values())
	for c := range rLay.Channels {
		for d := range rLay.Dim {
			i := c*rLay.Dim + d
			var want float32
			if c < wLay.Channels && d < wLay.Dim {
				// The accumulator adds its update onto a zeroed state, so the
				// stored value equals the update itself.
				want = wv[c*wLay.Dim+d]
			}
			if got[i] != want {
				return fmt.Errorf("reader element (channel %d, dim %d) reported %v, want %v: "+
					"the runtime resolves shared-state addresses differently than the two layouts predict",
					c, d, got[i], want)
			}
		}
	}
	if ok, at := shared.tailIntact(); !ok {
		return fmt.Errorf("the engine wrote past the shared buffer's logical size (tail byte +%d changed)", at)
	}
	fmt.Println("  the reader saw the writer's two channels laid flat across four of its own, zeros beyond,")
	println("  exactly as the two StateLayout maps predict — no normalization, no silent re-padding")
	return nil
}

// compose builds one stream holding an accumulate stage and a same-shape read
// stage sharing shared.
func compose(lib *e5rt.Lib, root string, wShape, rShape [4]int, shared *sharedBuffer, writerText, readerText string) (*linkedStages, error) {
	wLay := e5rt.StateLayout{Channels: wShape[1], Dim: wShape[3]}
	rLay := e5rt.StateLayout{Channels: rShape[1], Dim: rShape[3]}
	return composeShapes(lib, root, wShape, rShape, wLay, rLay, shared, writerText, readerText)
}

func composeShapes(lib *e5rt.Lib, root string, wShape, rShape [4]int, wLay, rLay e5rt.StateLayout, shared *sharedBuffer, texts ...string) (*linkedStages, error) {
	writerText := mil.GenAccumulateState(stateName, wShape)
	readerText := mil.GenReadState(stateName, rShape)
	if len(texts) >= 2 {
		writerText, readerText = texts[0], texts[1]
	}
	l := &linkedStages{lib: lib}
	var err error
	defer func() {
		if err != nil {
			err = errors.Join(err, l.close())
		}
	}()
	if l.writer, err = openStage(lib, filepath.Join(root, "writer"), writerText, shared, wLay, true); err != nil {
		return nil, fmt.Errorf("writer stage: %w", err)
	}
	if l.reader, err = openStage(lib, filepath.Join(root, "reader"), readerText, shared, rLay, false); err != nil {
		return nil, fmt.Errorf("reader stage: %w", err)
	}
	if l.stream, err = lib.ExecutionStreamCreate(); err != nil {
		return nil, err
	}
	if err := lib.EncodeOperation(l.stream, l.writer.op); err != nil {
		return nil, fmt.Errorf("encode writer: %w", err)
	}
	if err := lib.EncodeOperation(l.stream, l.reader.op); err != nil {
		return nil, fmt.Errorf("encode reader: %w", err)
	}
	return l, nil
}

func pointerAt(addr uintptr) unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&addr))
}

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
	return obj, unsafe.Slice((*byte)(pointerAt(ptr)), nbytes+tail), nil
}

func writeFP16(data []byte, v []float32) {
	for i, x := range v {
		bits := ane.Float32ToFP16(x)
		data[2*i] = byte(bits)
		data[2*i+1] = byte(bits >> 8)
	}
}

func readFP16(data []byte, n int) []float32 {
	out := make([]float32, n)
	for i := range out {
		out[i] = ane.FP16ToFloat32(uint16(data[2*i]) | uint16(data[2*i+1])<<8)
	}
	return out
}

// readTensorPadded reads a [1, C, 1, D] fp16 state out of data following the
// StateLayout offsets: each channel starts on its own RowBytes boundary, so
// element (c, d) lives at c*RowBytes + d*2 and not at the packed position.
func readTensorPadded(data []byte, lay e5rt.StateLayout) []float32 {
	out := make([]float32, lay.Values())
	for c := range lay.Channels {
		row := data[c*lay.RowBytes():]
		for d := range lay.Dim {
			out[c*lay.Dim+d] = ane.FP16ToFloat32(uint16(row[2*d]) | uint16(row[2*d+1])<<8)
		}
	}
	return out
}

func poison(tail []byte) {
	for i := range tail {
		tail[i] = byte(i*31 + 7)
	}
}

func toFloat32(v []float64) []float32 {
	out := make([]float32, len(v))
	for i, x := range v {
		out[i] = float32(x)
	}
	return out
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
