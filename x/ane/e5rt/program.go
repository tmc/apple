package e5rt

import (
	"encoding/binary"
	"errors"
	"fmt"
	"os"
	"sync"
	"unsafe"

	"github.com/tmc/apple/x/ane"
)

// A Port describes one externally bound program port.
//
// Size is the number of bytes the program reads or writes. The caller is
// responsible for using the representation the MIL program declares.
type Port struct {
	Name string
	Size int
}

// ProgramOptions describes a single-function E5RT program.
//
// ModelPath names the MIL program to compile. CacheDir receives the compiler's
// bundles. Inputs and Outputs name the function's externally bound ports.
// FunctionName defaults to "main", DeviceMask defaults to [ComputeDeviceANE],
// and Segmenter defaults to "graph".
//
// A Program has one encoded operation. It deliberately does not expose E5RT's
// reset, prepare, events, or multi-operation interfaces: their useful behavior
// has not been established by this package.
type ProgramOptions struct {
	ModelPath string
	CacheDir  string

	FunctionName string
	DeviceMask   uint64
	Segmenter    string

	ForceRecompilation bool
	Inputs             []Port
	Outputs            []Port
}

// BundleOptions describes an already-compiled E5RT program bundle.
//
// BundlePath names the bundle returned by the compiler. FunctionName defaults
// to "main". Inputs and Outputs name the function's externally bound ports.
//
// A bundle is not a portable artifact: E5RT bundle reuse has been observed
// across processes that share a code-signing identity and a parent on macOS
// 26.x. Reuse after an aned restart, a reboot, or from an unrelated process is
// unmeasured.
type BundleOptions struct {
	BundlePath string

	FunctionName string
	Inputs       []Port
	Outputs      []Port
}

// A Buffer is the host-visible storage bound to one [Program] port.
//
// Bytes remains valid until the Program is closed. Execute serializes with
// Close, but callers must not read or write Bytes while another goroutine calls
// Execute.
type Buffer struct {
	data []byte
}

// Bytes returns the storage bound to the port.
func (b *Buffer) Bytes() []byte {
	if b == nil {
		return nil
	}
	return b.data
}

// WriteFP16 stores data as contiguous little-endian IEEE 754 binary16 values.
func (b *Buffer) WriteFP16(data []float32) error {
	if b == nil {
		return errors.New("e5rt: nil buffer")
	}
	if len(data) != len(b.data)/2 || len(b.data)%2 != 0 {
		return fmt.Errorf("e5rt: write %d fp16 values to %d-byte buffer", len(data), len(b.data))
	}
	for i, value := range data {
		binary.LittleEndian.PutUint16(b.data[2*i:], ane.Float32ToFP16(value))
	}
	return nil
}

// ReadFP16 decodes contiguous little-endian IEEE 754 binary16 values into dst.
func (b *Buffer) ReadFP16(dst []float32) error {
	if b == nil {
		return errors.New("e5rt: nil buffer")
	}
	if len(dst) != len(b.data)/2 || len(b.data)%2 != 0 {
		return fmt.Errorf("e5rt: read %d-byte buffer into %d fp16 values", len(b.data), len(dst))
	}
	for i := range dst {
		dst[i] = ane.FP16ToFloat32(binary.LittleEndian.Uint16(b.data[2*i:]))
	}
	return nil
}

// A Program is a compiled, bound, and encoded E5RT function.
//
// The zero value is not usable. Programs are not safe for concurrent use.
// Call [Program.Close] when finished.
type Program struct {
	mu     sync.Mutex
	lib    *Lib
	closed bool

	config    uintptr
	compiler  uintptr
	options   uintptr
	library   uintptr
	function  uintptr
	opOptions uintptr
	op        uintptr
	stream    uintptr

	inputs  map[string]*programPort
	outputs map[string]*programPort
}

type programPort struct {
	port   uintptr
	buffer uintptr
	data   Buffer
}

// Compile compiles, binds, and encodes one E5RT function.
//
// Execute may then be called repeatedly after changing the input buffers. The
// result is a single encoded operation; Compile does not add an operation to an
// existing stream or plan a multi-function graph.
func Compile(opts ProgramOptions) (_ *Program, err error) {
	if err := validateProgramOptions(&opts); err != nil {
		return nil, err
	}
	if err := os.MkdirAll(opts.CacheDir, 0o755); err != nil {
		return nil, fmt.Errorf("create cache directory: %w", err)
	}

	lib, err := Open()
	if err != nil {
		return nil, fmt.Errorf("open e5rt: %w", err)
	}
	p := &Program{
		lib:     lib,
		inputs:  make(map[string]*programPort, len(opts.Inputs)),
		outputs: make(map[string]*programPort, len(opts.Outputs)),
	}
	defer func() {
		if err != nil {
			err = errors.Join(err, p.Close())
		}
	}()

	if p.config, err = lib.CompilerConfigOptionsCreate(); err != nil {
		return nil, fmt.Errorf("create compiler config: %w", err)
	}
	if err := lib.CompilerConfigOptionsSetCacheBundleLocation(p.config, opts.CacheDir); err != nil {
		return nil, fmt.Errorf("set cache location: %w", err)
	}
	if p.compiler, err = lib.CompilerCreateWithConfig(p.config); err != nil {
		return nil, fmt.Errorf("create compiler: %w", err)
	}
	if p.options, err = lib.CompilerOptionsCreate(); err != nil {
		return nil, fmt.Errorf("create compiler options: %w", err)
	}
	if err := lib.CompilerOptionsSetComputeDeviceTypesMask(p.options, opts.DeviceMask); err != nil {
		return nil, fmt.Errorf("set device mask: %w", err)
	}
	if err := lib.CompilerOptionsSetForceRecompilation(p.options, opts.ForceRecompilation); err != nil {
		return nil, fmt.Errorf("set force recompilation: %w", err)
	}
	if err := lib.CompilerOptionsSetSegmenter(p.options, opts.Segmenter); err != nil {
		return nil, fmt.Errorf("set segmenter: %w", err)
	}
	if p.library, err = lib.CompilerCompile(p.compiler, opts.ModelPath, p.options); err != nil {
		return nil, fmt.Errorf("compile model: %w", err)
	}
	if err := p.prepare(opts.FunctionName, opts.Inputs, opts.Outputs); err != nil {
		return nil, err
	}
	return p, nil
}

// OpenBundle opens, binds, and encodes a function from an existing compiled
// E5RT bundle.
func OpenBundle(opts BundleOptions) (_ *Program, err error) {
	if opts.BundlePath == "" {
		return nil, errors.New("e5rt: empty bundle path")
	}
	if err := normalizeProgramPorts(&opts.FunctionName, opts.Inputs, opts.Outputs); err != nil {
		return nil, err
	}
	lib, err := Open()
	if err != nil {
		return nil, fmt.Errorf("open e5rt: %w", err)
	}
	p := &Program{
		lib:     lib,
		inputs:  make(map[string]*programPort, len(opts.Inputs)),
		outputs: make(map[string]*programPort, len(opts.Outputs)),
	}
	defer func() {
		if err != nil {
			err = errors.Join(err, p.Close())
		}
	}()
	if p.library, err = lib.ProgramLibraryCreate(opts.BundlePath); err != nil {
		return nil, fmt.Errorf("open program bundle: %w", err)
	}
	if err := p.prepare(opts.FunctionName, opts.Inputs, opts.Outputs); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *Program) prepare(functionName string, inputs, outputs []Port) error {
	var err error
	if p.function, err = p.lib.ProgramLibraryRetainProgramFunction(p.library, functionName); err != nil {
		return fmt.Errorf("retain function %q: %w", functionName, err)
	}
	if p.opOptions, err = p.lib.PrecompiledComputeOpOptionsCreate(p.function); err != nil {
		return fmt.Errorf("create operation options: %w", err)
	}
	if err := p.lib.PrecompiledComputeOpOptionsSetOperationName(p.opOptions, functionName); err != nil {
		return fmt.Errorf("set operation name: %w", err)
	}
	if err := p.lib.PrecompiledComputeOpOptionsSetAllocateIntermediateBuffers(p.opOptions, true); err != nil {
		return fmt.Errorf("allocate intermediate buffers: %w", err)
	}
	if p.op, err = p.lib.OperationCreatePrecompiled(p.opOptions); err != nil {
		return fmt.Errorf("create operation: %w", err)
	}
	for _, spec := range inputs {
		port, err := p.bind(spec, true)
		if err != nil {
			return err
		}
		p.inputs[spec.Name] = port
	}
	for _, spec := range outputs {
		port, err := p.bind(spec, false)
		if err != nil {
			return err
		}
		p.outputs[spec.Name] = port
	}
	if p.stream, err = p.lib.ExecutionStreamCreate(); err != nil {
		return fmt.Errorf("create execution stream: %w", err)
	}
	if err := p.lib.EncodeOperation(p.stream, p.op); err != nil {
		return fmt.Errorf("encode operation: %w", err)
	}
	return nil
}

func validateProgramOptions(opts *ProgramOptions) error {
	if opts.ModelPath == "" {
		return errors.New("e5rt: empty model path")
	}
	if opts.CacheDir == "" {
		return errors.New("e5rt: empty cache directory")
	}
	if err := normalizeProgramPorts(&opts.FunctionName, opts.Inputs, opts.Outputs); err != nil {
		return err
	}
	if opts.DeviceMask == 0 {
		opts.DeviceMask = ComputeDeviceANE
	}
	if opts.Segmenter == "" {
		opts.Segmenter = "graph"
	}
	return nil
}

func normalizeProgramPorts(functionName *string, inputs, outputs []Port) error {
	if *functionName == "" {
		*functionName = "main"
	}
	seen := make(map[string]bool, len(inputs)+len(outputs))
	for _, ports := range [][]Port{inputs, outputs} {
		for _, port := range ports {
			if port.Name == "" {
				return errors.New("e5rt: empty port name")
			}
			if port.Size <= 0 {
				return fmt.Errorf("e5rt: port %q has non-positive size", port.Name)
			}
			if seen[port.Name] {
				return fmt.Errorf("e5rt: duplicate port %q", port.Name)
			}
			seen[port.Name] = true
		}
	}
	return nil
}

func (p *Program) bind(spec Port, input bool) (*programPort, error) {
	size, err := bufferAllocationSize(spec.Size)
	if err != nil {
		return nil, fmt.Errorf("allocate buffer for port %q: %w", spec.Name, err)
	}
	var port uintptr
	if input {
		port, err = p.lib.OperationRetainInputPort(p.op, spec.Name)
	} else {
		port, err = p.lib.OperationRetainOutputPort(p.op, spec.Name)
	}
	if err != nil {
		return nil, fmt.Errorf("retain port %q: %w", spec.Name, err)
	}
	result := &programPort{port: port}
	if result.buffer, err = p.lib.BufferObjectAlloc(uintptr(size), 0); err != nil {
		_ = p.lib.IOPortRelease(result.port)
		return nil, fmt.Errorf("allocate buffer for port %q: %w", spec.Name, err)
	}
	ptr, err := p.lib.BufferObjectGetDataPtr(result.buffer)
	if err != nil {
		_ = p.lib.BufferObjectRelease(result.buffer)
		_ = p.lib.IOPortRelease(result.port)
		return nil, fmt.Errorf("get buffer address for port %q: %w", spec.Name, err)
	}
	if err := p.lib.IOPortBindBufferObject(result.port, result.buffer); err != nil {
		_ = p.lib.BufferObjectRelease(result.buffer)
		_ = p.lib.IOPortRelease(result.port)
		return nil, fmt.Errorf("bind buffer for port %q: %w", spec.Name, err)
	}
	result.data.data = unsafe.Slice((*byte)(pointerAt(ptr)), spec.Size)
	return result, nil
}

func bufferAllocationSize(n int) (int, error) {
	maxInt := int(^uint(0) >> 1)
	if n > maxInt-63 {
		return 0, fmt.Errorf("buffer size %d overflows 64-byte alignment", n)
	}
	return max((n+63)&^63, 64), nil
}

// Input returns the buffer bound to the named input port.
func (p *Program) Input(name string) (*Buffer, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed {
		return nil, errors.New("e5rt: program is closed")
	}
	port, ok := p.inputs[name]
	if !ok {
		return nil, fmt.Errorf("e5rt: input port %q not found", name)
	}
	return &port.data, nil
}

// Output returns the buffer bound to the named output port.
func (p *Program) Output(name string) (*Buffer, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed {
		return nil, errors.New("e5rt: program is closed")
	}
	port, ok := p.outputs[name]
	if !ok {
		return nil, fmt.Errorf("e5rt: output port %q not found", name)
	}
	return &port.data, nil
}

// Execute synchronously runs the encoded operation.
func (p *Program) Execute() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed {
		return errors.New("e5rt: program is closed")
	}
	if err := p.lib.ExecuteSync(p.stream); err != nil {
		return fmt.Errorf("execute program: %w", err)
	}
	return nil
}

// Close releases the E5RT objects owned by p. It is safe to call Close more
// than once. Any buffers returned by Input or Output become invalid on return.
func (p *Program) Close() error {
	if p == nil {
		return nil
	}
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed {
		return nil
	}
	p.closed = true
	var errs []error
	release := func(name string, handle *uintptr, f func(uintptr) error) {
		if *handle == 0 {
			return
		}
		if err := f(*handle); err != nil {
			errs = append(errs, fmt.Errorf("release %s: %w", name, err))
		}
		*handle = 0
	}
	release("stream", &p.stream, p.lib.ExecutionStreamRelease)
	for _, port := range p.inputs {
		release("input port", &port.port, p.lib.IOPortRelease)
		release("input buffer", &port.buffer, p.lib.BufferObjectRelease)
		port.data.data = nil
	}
	for _, port := range p.outputs {
		release("output port", &port.port, p.lib.IOPortRelease)
		release("output buffer", &port.buffer, p.lib.BufferObjectRelease)
		port.data.data = nil
	}
	release("operation", &p.op, p.lib.OperationRelease)
	release("operation options", &p.opOptions, p.lib.PrecompiledComputeOpOptionsRelease)
	release("function", &p.function, p.lib.ProgramFunctionRelease)
	release("library", &p.library, p.lib.ProgramLibraryRelease)
	release("compiler options", &p.options, p.lib.CompilerOptionsRelease)
	release("compiler", &p.compiler, p.lib.CompilerRelease)
	release("compiler config", &p.config, p.lib.CompilerConfigOptionsRelease)
	return errors.Join(errs...)
}

func pointerAt(addr uintptr) unsafe.Pointer {
	return *(*unsafe.Pointer)(unsafe.Pointer(&addr))
}
