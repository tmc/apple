//go:build !darwin

package ane

import "time"

// Runtime manages a connection to the ANE hardware.
type Runtime struct{}

// Probe returns device information about the ANE.
// On non-Darwin platforms, it always reports no ANE.
func Probe() (DeviceInfo, error) { return DeviceInfo{}, ErrNoANE }

// Open creates a new Runtime for ANE inference.
// On non-Darwin platforms, it always returns ErrNoANE.
func Open() (*Runtime, error) { return nil, ErrNoANE }

// Close releases the runtime resources.
func (rt *Runtime) Close() error { return nil }

// Info returns the device information for this runtime.
func (rt *Runtime) Info() DeviceInfo { return DeviceInfo{} }

// CompileCount returns the number of compilations performed.
func (rt *Runtime) CompileCount() int64 { return 0 }

// Compile compiles a model and returns a ready-to-evaluate Kernel.
func (rt *Runtime) Compile(opts CompileOptions) (*Kernel, error) { return nil, ErrNoANE }

// CompileWithStats compiles a model and returns a Kernel along with compilation timing.
func (rt *Runtime) CompileWithStats(opts CompileOptions) (*Kernel, CompileStats, error) {
	return nil, CompileStats{}, ErrNoANE
}

// TensorLayout describes the compiled model's memory layout for a single tensor.
type TensorLayout struct {
	Channels    int
	Width       int
	Height      int
	ElemSize    int
	RowStride   int
	PlaneStride int
}

func (l TensorLayout) AllocSize() int       { return l.Channels * l.PlaneStride }
func (l TensorLayout) LogicalBytes() int    { return l.Channels * l.Width * l.Height * l.ElemSize }
func (l TensorLayout) LogicalElements() int { return l.Channels * l.Width * l.Height }

// RowStrideFor computes the minimum 64-byte-aligned row stride.
func RowStrideFor(width, elemSize int) int {
	raw := width * elemSize
	const align = 64
	return (raw + align - 1) &^ (align - 1)
}

// Kernel represents a compiled and loaded model ready for evaluation.
type Kernel struct{}

func (k *Kernel) InputSurface(i int) uintptr      { return 0 }
func (k *Kernel) OutputSurface(i int) uintptr     { return 0 }
func (k *Kernel) InputSurfaces() []uintptr        { return nil }
func (k *Kernel) OutputSurfaces() []uintptr       { return nil }
func (k *Kernel) InputAllocSize(i int) int        { return 0 }
func (k *Kernel) OutputAllocSize(i int) int       { return 0 }
func (k *Kernel) NumInputs() int                  { return 0 }
func (k *Kernel) NumOutputs() int                 { return 0 }
func (k *Kernel) InputChannels(i int) int         { return 0 }
func (k *Kernel) OutputChannels(i int) int        { return 0 }
func (k *Kernel) Spatial(i int) int               { return 0 }
func (k *Kernel) InputLayout(i int) TensorLayout  { return TensorLayout{} }
func (k *Kernel) OutputLayout(i int) TensorLayout { return TensorLayout{} }

func (k *Kernel) WriteInput(i int, data []byte) error        { return ErrNoANE }
func (k *Kernel) ReadOutput(i int, data []byte) error        { return ErrNoANE }
func (k *Kernel) WriteInputF32(i int, data []float32) error  { return ErrNoANE }
func (k *Kernel) ReadOutputF32(i int, data []float32) error  { return ErrNoANE }
func (k *Kernel) WriteInputFP16(i int, data []float32) error { return ErrNoANE }
func (k *Kernel) WriteInputFP16Channels(i, channel int, data []float32) error {
	return ErrNoANE
}
func (k *Kernel) ReadOutputFP16(i int, data []float32) error { return ErrNoANE }
func (k *Kernel) ReadOutputFP16Channels(i, channel int, data []float32) error {
	return ErrNoANE
}

func (k *Kernel) Eval() error { return ErrNoANE }

func (k *Kernel) ModelObjcID() uintptr        { return 0 }
func (k *Kernel) InMemModelObjcID() uintptr   { return 0 }
func (k *Kernel) CompileModelType() ModelType { return 0 }
func (k *Kernel) RawRequest() uintptr         { return 0 }
func (k *Kernel) RawPerfStatsMask() uint32    { return 0 }

func (rt *Runtime) ClientObjcID() uintptr { return 0 }
func (k *Kernel) Close() error            { return nil }

func (k *Kernel) EvalWithSignalEvent(signalPort uint32, signalValue uint64, cfg SharedEventEvalOptions) error {
	return ErrNoANE
}
func (k *Kernel) EvalBidirectional(waitPort uint32, waitValue uint64, signalPort uint32, signalValue uint64, cfg SharedEventEvalOptions) error {
	return ErrNoANE
}

// SharedEvent wraps an IOSurfaceSharedEvent for ANE↔GPU/CPU synchronization.
type SharedEvent struct{}

func NewSharedEvent() (*SharedEvent, error)                          { return nil, ErrNoANE }
func SharedEventFromPort(port uint32) (*SharedEvent, error)          { return nil, ErrNoANE }
func (e *SharedEvent) Port() uint32                                  { return 0 }
func (e *SharedEvent) SignaledValue() uint64                         { return 0 }
func (e *SharedEvent) Signal(value uint64)                           {}
func (e *SharedEvent) Wait(value uint64, timeout time.Duration) bool             { return false }
func (e *SharedEvent) TimeWait(value uint64, timeout time.Duration) (bool, time.Duration) { return false, 0 }
func (e *SharedEvent) Close() error                                              { return nil }

// Float32ToFP16 converts a float32 to IEEE 754 half-precision.
func Float32ToFP16(f float32) uint16 { return 0 }

// FP16ToFloat32 converts an IEEE 754 half-precision value to float32.
func FP16ToFloat32(h uint16) float32 { return 0 }

// MetalDevice wraps a Metal GPU device for zero-copy interop with ANE.
type MetalDevice struct{}

func OpenMetal() (*MetalDevice, error) { return nil, ErrNoANE }
func (d *MetalDevice) Close() error    { return nil }

func (k *Kernel) MetalInputBuffer(d *MetalDevice, i int) (any, error)  { return nil, ErrNoANE }
func (k *Kernel) MetalOutputBuffer(d *MetalDevice, i int) (any, error) { return nil, ErrNoANE }

func (d *MetalDevice) MetalSharedEvent(ev *SharedEvent) (any, error) { return nil, ErrNoANE }
func (d *MetalDevice) NewMetalSharedEvent() (any, *SharedEvent, error) {
	return nil, nil, ErrNoANE
}

func (k *Kernel) EvalAsync() <-chan error {
	ch := make(chan error, 1)
	ch <- ErrNoANE
	return ch
}
func (k *Kernel) EvalAsyncWithCallback(fn func(error)) { fn(ErrNoANE) }

// RequestPool pre-allocates a ring of ANE requests for pipelined evaluation.
type RequestPool struct{}

// PooledRequest is a request checked out from a pool.
type PooledRequest struct{}

func NewRequestPool(k *Kernel, depth int) (*RequestPool, error) { return nil, ErrNoANE }
func (p *RequestPool) Acquire() *PooledRequest                  { return nil }
func (p *RequestPool) Close() error                             { return nil }
func (pr *PooledRequest) Eval() error                           { return ErrNoANE }
func (pr *PooledRequest) Release()                              {}

// StateHandle manages KV cache state for stateful MIL models.
type StateHandle struct{}

func NewStateHandle(_ *Kernel, _ int) *StateHandle { return &StateHandle{} }
func (s *StateHandle) Kernel() *Kernel             { return nil }
func (s *StateHandle) Position() int               { return 0 }
func (s *StateHandle) MaxSeq() int                 { return 0 }
func (s *StateHandle) Reset()                      {}
func (s *StateHandle) Advance(_ int)               {}
func (s *StateHandle) Remaining() int              { return 0 }
func (s *StateHandle) Close() error                { return nil }

// Pipeline manages SharedEvent synchronization between ANE and Metal.
type Pipeline struct{}

func NewPipeline(_ *MetalDevice) (*Pipeline, error) { return nil, ErrNoANE }
func (p *Pipeline) Counter() uint64                 { return 0 }
func (p *Pipeline) ANEToMetal(_ *Kernel) error      { return ErrNoANE }
func (p *Pipeline) WaitOnANE(_ *Kernel) error       { return ErrNoANE }
func (p *Pipeline) Bidirectional(_ *Kernel) error   { return ErrNoANE }
func (p *Pipeline) Metal() *MetalDevice             { return nil }
func (p *Pipeline) MetalEvent() any                 { return nil }
func (p *Pipeline) ANEEvent() *SharedEvent          { return nil }
func (p *Pipeline) Close() error                    { return nil }
