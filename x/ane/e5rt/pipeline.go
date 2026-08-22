package e5rt

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"unsafe"
)

// MaxPipelineStages is the largest pipeline one stream may contain.
//
// A macOS 26.x stream-retention control measured 15 distinct compiled
// functions in this process: releasing an encoded operation did not free the
// next function slot, while releasing its stream did. The pool is shared, so
// contention can make a smaller pipeline fail. A pipeline with more stages
// must be planned into separate streams by its caller.
const MaxPipelineStages = 15

// A PipelinePort identifies a named port on a pipeline stage.
type PipelinePort struct {
	Stage int
	Name  string
}

// A PipelineLink binds an output buffer to a later stage's input port.
//
// The linked stages execute in order in one stream. The two ports use the same
// buffer object; no host copy is made between them.
type PipelineLink struct {
	From PipelinePort
	To   PipelinePort
}

// A PipelineStage describes one function in a [Pipeline].
type PipelineStage struct {
	ModelPath    string
	FunctionName string
	Inputs       []Port
	Outputs      []Port
}

// PipelineOptions describes one encoded multi-stage E5RT stream.
//
// Each stage is compiled independently under CacheDir and then encoded in the
// order listed. Links must flow from an earlier stage to a later one. Unlinked
// inputs and outputs remain host-visible through [Pipeline.Input] and
// [Pipeline.Output].
type PipelineOptions struct {
	CacheDir string

	DeviceMask         uint64
	Segmenter          string
	ForceRecompilation bool

	Stages []PipelineStage
	Links  []PipelineLink
}

// A Pipeline is a compiled, bound, and encoded sequence of E5RT functions.
//
// The zero value is not usable. A Pipeline is not safe for concurrent use.
// Call [Pipeline.Close] when finished.
type Pipeline struct {
	mu     sync.Mutex
	lib    *Lib
	closed bool

	stages []*pipelineStage
	stream uintptr
}

type pipelineStage struct {
	config    uintptr
	compiler  uintptr
	options   uintptr
	library   uintptr
	function  uintptr
	opOptions uintptr
	op        uintptr

	inputs  map[string]*pipelinePort
	outputs map[string]*pipelinePort

	inputSpecs  []Port
	outputSpecs []Port
}

type pipelinePort struct {
	port   uintptr
	buffer uintptr // zero when this port borrows a linked output buffer
	data   *Buffer
}

// CompilePipeline compiles, links, and encodes stages into one E5RT stream.
func CompilePipeline(opts PipelineOptions) (_ *Pipeline, err error) {
	if err := validatePipelineOptions(&opts); err != nil {
		return nil, err
	}
	if err := os.MkdirAll(opts.CacheDir, 0o755); err != nil {
		return nil, fmt.Errorf("create cache directory: %w", err)
	}
	lib, err := Open()
	if err != nil {
		return nil, fmt.Errorf("open e5rt: %w", err)
	}
	p := &Pipeline{lib: lib, stages: make([]*pipelineStage, len(opts.Stages))}
	defer func() {
		if err != nil {
			err = errors.Join(err, p.Close())
		}
	}()
	for i, spec := range opts.Stages {
		stage, err := compilePipelineStage(lib, filepath.Join(opts.CacheDir, fmt.Sprintf("stage-%02d", i)), spec, opts)
		if err != nil {
			return nil, fmt.Errorf("compile stage %d: %w", i, err)
		}
		p.stages[i] = stage
	}
	if err := p.bind(opts.Links); err != nil {
		return nil, err
	}
	if p.stream, err = lib.ExecutionStreamCreate(); err != nil {
		return nil, fmt.Errorf("create execution stream: %w", err)
	}
	for i, stage := range p.stages {
		if err := lib.EncodeOperation(p.stream, stage.op); err != nil {
			return nil, fmt.Errorf("encode stage %d: %w", i, err)
		}
	}
	return p, nil
}

func validatePipelineOptions(opts *PipelineOptions) error {
	if opts.CacheDir == "" {
		return errors.New("e5rt: empty cache directory")
	}
	if len(opts.Stages) == 0 {
		return errors.New("e5rt: pipeline has no stages")
	}
	if len(opts.Stages) > MaxPipelineStages {
		return fmt.Errorf("e5rt: pipeline has %d stages, limit is %d", len(opts.Stages), MaxPipelineStages)
	}
	for i := range opts.Stages {
		stage := &opts.Stages[i]
		if stage.ModelPath == "" {
			return fmt.Errorf("e5rt: stage %d has empty model path", i)
		}
		if err := normalizeProgramPorts(&stage.FunctionName, stage.Inputs, stage.Outputs); err != nil {
			return fmt.Errorf("e5rt: stage %d: %w", i, err)
		}
	}
	if opts.DeviceMask == 0 {
		opts.DeviceMask = ComputeDeviceANE
	}
	if opts.Segmenter == "" {
		opts.Segmenter = "graph"
	}
	seenTargets := make(map[PipelinePort]bool, len(opts.Links))
	for _, link := range opts.Links {
		if err := validatePipelinePort(link.From, opts.Stages, "source output"); err != nil {
			return err
		}
		if err := validatePipelinePort(link.To, opts.Stages, "destination input"); err != nil {
			return err
		}
		fromSize, _ := pipelinePortSize(link.From, opts.Stages, false)
		toSize, _ := pipelinePortSize(link.To, opts.Stages, true)
		if fromSize != toSize {
			return fmt.Errorf("e5rt: link from stage %d port %q (%d bytes) to stage %d port %q (%d bytes)",
				link.From.Stage, link.From.Name, fromSize, link.To.Stage, link.To.Name, toSize)
		}
		if link.From.Stage >= link.To.Stage {
			return fmt.Errorf("e5rt: link from stage %d to stage %d is not forward", link.From.Stage, link.To.Stage)
		}
		if seenTargets[link.To] {
			return fmt.Errorf("e5rt: input stage %d port %q has multiple links", link.To.Stage, link.To.Name)
		}
		seenTargets[link.To] = true
	}
	return nil
}

func validatePipelinePort(port PipelinePort, stages []PipelineStage, kind string) error {
	if port.Stage < 0 || port.Stage >= len(stages) {
		return fmt.Errorf("e5rt: %s stage %d is out of range", kind, port.Stage)
	}
	if port.Name == "" {
		return fmt.Errorf("e5rt: %s at stage %d has empty name", kind, port.Stage)
	}
	if _, ok := pipelinePortSize(port, stages, kind != "source output"); ok {
		return nil
	}
	return fmt.Errorf("e5rt: %s stage %d port %q not found", kind, port.Stage, port.Name)
}

func pipelinePortSize(port PipelinePort, stages []PipelineStage, input bool) (int, bool) {
	ports := stages[port.Stage].Outputs
	if input {
		ports = stages[port.Stage].Inputs
	}
	for _, spec := range ports {
		if spec.Name == port.Name {
			return spec.Size, true
		}
	}
	return 0, false
}

func compilePipelineStage(lib *Lib, cacheDir string, spec PipelineStage, opts PipelineOptions) (_ *pipelineStage, err error) {
	if err := os.MkdirAll(cacheDir, 0o755); err != nil {
		return nil, fmt.Errorf("create cache directory: %w", err)
	}
	stage := &pipelineStage{
		inputs:      make(map[string]*pipelinePort, len(spec.Inputs)),
		outputs:     make(map[string]*pipelinePort, len(spec.Outputs)),
		inputSpecs:  spec.Inputs,
		outputSpecs: spec.Outputs,
	}
	defer func() {
		if err != nil {
			err = errors.Join(err, releasePipelineStage(lib, stage))
		}
	}()
	if stage.config, err = lib.CompilerConfigOptionsCreate(); err != nil {
		return nil, err
	}
	if err := lib.CompilerConfigOptionsSetCacheBundleLocation(stage.config, cacheDir); err != nil {
		return nil, err
	}
	if stage.compiler, err = lib.CompilerCreateWithConfig(stage.config); err != nil {
		return nil, err
	}
	if stage.options, err = lib.CompilerOptionsCreate(); err != nil {
		return nil, err
	}
	if err := lib.CompilerOptionsSetComputeDeviceTypesMask(stage.options, opts.DeviceMask); err != nil {
		return nil, err
	}
	if err := lib.CompilerOptionsSetForceRecompilation(stage.options, opts.ForceRecompilation); err != nil {
		return nil, err
	}
	if err := lib.CompilerOptionsSetSegmenter(stage.options, opts.Segmenter); err != nil {
		return nil, err
	}
	if stage.library, err = lib.CompilerCompile(stage.compiler, spec.ModelPath, stage.options); err != nil {
		return nil, err
	}
	if stage.function, err = lib.ProgramLibraryRetainProgramFunction(stage.library, spec.FunctionName); err != nil {
		return nil, err
	}
	if stage.opOptions, err = lib.PrecompiledComputeOpOptionsCreate(stage.function); err != nil {
		return nil, err
	}
	if err := lib.PrecompiledComputeOpOptionsSetOperationName(stage.opOptions, spec.FunctionName); err != nil {
		return nil, err
	}
	if err := lib.PrecompiledComputeOpOptionsSetAllocateIntermediateBuffers(stage.opOptions, true); err != nil {
		return nil, err
	}
	if stage.op, err = lib.OperationCreatePrecompiled(stage.opOptions); err != nil {
		return nil, err
	}
	return stage, nil
}

func (p *Pipeline) bind(links []PipelineLink) error {
	linkedInputs := make(map[PipelinePort]PipelinePort, len(links))
	for _, link := range links {
		linkedInputs[link.To] = link.From
	}
	for i, stage := range p.stages {
		for _, spec := range stageSpecs(stage, false) {
			port, err := p.bindNew(stage, spec, false)
			if err != nil {
				return fmt.Errorf("bind stage %d output %q: %w", i, spec.Name, err)
			}
			stage.outputs[spec.Name] = port
		}
		for _, spec := range stageSpecs(stage, true) {
			key := PipelinePort{Stage: i, Name: spec.Name}
			if source, ok := linkedInputs[key]; ok {
				sourcePort := p.stages[source.Stage].outputs[source.Name]
				port, err := p.bindExisting(stage, spec.Name, sourcePort, true)
				if err != nil {
					return fmt.Errorf("bind stage %d input %q: %w", i, spec.Name, err)
				}
				stage.inputs[spec.Name] = port
				continue
			}
			port, err := p.bindNew(stage, spec, true)
			if err != nil {
				return fmt.Errorf("bind stage %d input %q: %w", i, spec.Name, err)
			}
			stage.inputs[spec.Name] = port
		}
	}
	return nil
}

func stageSpecs(stage *pipelineStage, input bool) []Port {
	if input {
		return stage.inputSpecs
	}
	return stage.outputSpecs
}

func (p *Pipeline) bindNew(stage *pipelineStage, spec Port, input bool) (*pipelinePort, error) {
	size, err := bufferAllocationSize(spec.Size)
	if err != nil {
		return nil, err
	}
	port, err := retainPipelinePort(p.lib, stage.op, spec.Name, input)
	if err != nil {
		return nil, err
	}
	result := &pipelinePort{port: port}
	if result.buffer, err = p.lib.BufferObjectAlloc(uintptr(size), 0); err != nil {
		_ = p.lib.IOPortRelease(result.port)
		return nil, err
	}
	ptr, err := p.lib.BufferObjectGetDataPtr(result.buffer)
	if err != nil {
		_ = p.lib.BufferObjectRelease(result.buffer)
		_ = p.lib.IOPortRelease(result.port)
		return nil, err
	}
	if err := p.lib.IOPortBindBufferObject(result.port, result.buffer); err != nil {
		_ = p.lib.BufferObjectRelease(result.buffer)
		_ = p.lib.IOPortRelease(result.port)
		return nil, err
	}
	result.data = &Buffer{data: unsafeBytes(ptr, spec.Size)}
	return result, nil
}

func (p *Pipeline) bindExisting(stage *pipelineStage, name string, source *pipelinePort, input bool) (*pipelinePort, error) {
	port, err := retainPipelinePort(p.lib, stage.op, name, input)
	if err != nil {
		return nil, err
	}
	if err := p.lib.IOPortBindBufferObject(port, source.buffer); err != nil {
		_ = p.lib.IOPortRelease(port)
		return nil, err
	}
	return &pipelinePort{port: port, data: source.data}, nil
}

func retainPipelinePort(lib *Lib, op uintptr, name string, input bool) (uintptr, error) {
	if input {
		return lib.OperationRetainInputPort(op, name)
	}
	return lib.OperationRetainOutputPort(op, name)
}

func unsafeBytes(addr uintptr, n int) []byte {
	return unsafe.Slice((*byte)(pointerAt(addr)), n)
}

// Input returns the buffer bound to name on stage.
func (p *Pipeline) Input(stage int, name string) (*Buffer, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed {
		return nil, errors.New("e5rt: pipeline is closed")
	}
	if stage < 0 || stage >= len(p.stages) {
		return nil, fmt.Errorf("e5rt: stage %d is out of range", stage)
	}
	port, ok := p.stages[stage].inputs[name]
	if !ok {
		return nil, fmt.Errorf("e5rt: stage %d input port %q not found", stage, name)
	}
	return port.data, nil
}

// Output returns the buffer bound to name on stage.
func (p *Pipeline) Output(stage int, name string) (*Buffer, error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed {
		return nil, errors.New("e5rt: pipeline is closed")
	}
	if stage < 0 || stage >= len(p.stages) {
		return nil, fmt.Errorf("e5rt: stage %d is out of range", stage)
	}
	port, ok := p.stages[stage].outputs[name]
	if !ok {
		return nil, fmt.Errorf("e5rt: stage %d output port %q not found", stage, name)
	}
	return port.data, nil
}

// Execute synchronously runs every encoded pipeline stage in order.
func (p *Pipeline) Execute() error {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.closed {
		return errors.New("e5rt: pipeline is closed")
	}
	if err := p.lib.ExecuteSync(p.stream); err != nil {
		return fmt.Errorf("execute pipeline: %w", err)
	}
	return nil
}

// Close releases the E5RT objects owned by p. It is safe to call Close more
// than once. Any buffers returned by Input or Output become invalid on return.
func (p *Pipeline) Close() error {
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
	if p.stream != 0 {
		if err := p.lib.ExecutionStreamRelease(p.stream); err != nil {
			errs = append(errs, fmt.Errorf("release stream: %w", err))
		}
		p.stream = 0
	}
	for _, stage := range p.stages {
		errs = append(errs, releasePipelinePorts(p.lib, stage))
	}
	for _, stage := range p.stages {
		errs = append(errs, releasePipelineBuffers(p.lib, stage))
	}
	for _, stage := range p.stages {
		errs = append(errs, releasePipelineStageHandles(p.lib, stage))
	}
	return errors.Join(errs...)
}

func releasePipelineStage(lib *Lib, stage *pipelineStage) error {
	return errors.Join(
		releasePipelinePorts(lib, stage),
		releasePipelineBuffers(lib, stage),
		releasePipelineStageHandles(lib, stage),
	)
}

func releasePipelinePorts(lib *Lib, stage *pipelineStage) error {
	if stage == nil {
		return nil
	}
	var errs []error
	for _, port := range stage.inputs {
		if err := releasePipelineHandle("input port", &port.port, lib.IOPortRelease); err != nil {
			errs = append(errs, err)
		}
	}
	for _, port := range stage.outputs {
		if err := releasePipelineHandle("output port", &port.port, lib.IOPortRelease); err != nil {
			errs = append(errs, err)
		}
	}
	return errors.Join(errs...)
}

func releasePipelineBuffers(lib *Lib, stage *pipelineStage) error {
	if stage == nil {
		return nil
	}
	var errs []error
	for _, port := range stage.inputs {
		if port.buffer == 0 {
			continue
		}
		if err := releasePipelineHandle("input buffer", &port.buffer, lib.BufferObjectRelease); err != nil {
			errs = append(errs, err)
		}
		port.data.data = nil
	}
	for _, port := range stage.outputs {
		if port.buffer == 0 {
			continue
		}
		if err := releasePipelineHandle("output buffer", &port.buffer, lib.BufferObjectRelease); err != nil {
			errs = append(errs, err)
		}
		port.data.data = nil
	}
	return errors.Join(errs...)
}

func releasePipelineStageHandles(lib *Lib, stage *pipelineStage) error {
	if stage == nil {
		return nil
	}
	return errors.Join(
		releasePipelineHandle("operation", &stage.op, lib.OperationRelease),
		releasePipelineHandle("operation options", &stage.opOptions, lib.PrecompiledComputeOpOptionsRelease),
		releasePipelineHandle("function", &stage.function, lib.ProgramFunctionRelease),
		releasePipelineHandle("library", &stage.library, lib.ProgramLibraryRelease),
		releasePipelineHandle("compiler options", &stage.options, lib.CompilerOptionsRelease),
		releasePipelineHandle("compiler", &stage.compiler, lib.CompilerRelease),
		releasePipelineHandle("compiler config", &stage.config, lib.CompilerConfigOptionsRelease),
	)
}

func releasePipelineHandle(name string, handle *uintptr, release func(uintptr) error) error {
	if *handle == 0 {
		return nil
	}
	err := release(*handle)
	*handle = 0
	if err != nil {
		return fmt.Errorf("release %s: %w", name, err)
	}
	return nil
}
