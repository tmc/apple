//go:build !darwin

package espresso

import (
	"errors"
	"time"
)

// ErrUnsupported is returned on platforms that do not support Espresso.
var ErrUnsupported = errors.New("espresso: unsupported platform")

// Context manages execution state for Espresso networks.
type Context struct{}

// Open creates a new Espresso execution context.
func Open(_ ...ContextOption) (*Context, error) {
	return nil, ErrUnsupported
}

// Close releases context resources.
func (c *Context) Close() error { return nil }

func (c *Context) isClosed() bool { return true }

// Platform returns the platform this context was created with.
func (c *Context) Platform() int { return 0 }

// Network is a loaded Espresso neural network ready for execution.
type Network struct {
	closed bool
}

// LoadNetwork loads a network from Espresso IR.
func (c *Context) LoadNetwork(_ []byte, _ ...NetworkOption) (*Network, error) {
	return nil, ErrUnsupported
}

// LoadNetworkFromFile loads a network from an Espresso net.json file.
func (c *Context) LoadNetworkFromFile(_ string, _ ...NetworkOption) (*Network, error) {
	return nil, ErrUnsupported
}

// LayerCount returns the number of layers.
func (n *Network) LayerCount() int { return 0 }

// LayerNames returns the names of all layers.
func (n *Network) LayerNames() []string { return nil }

// HasLayer reports whether the network contains a layer with the given name.
func (n *Network) HasLayer(_ string) bool { return false }

// Close releases network resources.
func (n *Network) Close() error { return nil }

// Eval executes the network on the given frame.
func (n *Network) Eval(_ *Frame) error { return ErrUnsupported }

// EvalAsync executes the network asynchronously.
func (n *Network) EvalAsync(_ *Frame) <-chan error {
	ch := make(chan error, 1)
	ch <- ErrUnsupported
	return ch
}

// EvalAsyncWithCallback executes the network asynchronously with a callback.
func (n *Network) EvalAsyncWithCallback(_ *Frame, fn func(error)) {
	fn(ErrUnsupported)
}

// EvalWithOptions executes with custom options.
func (n *Network) EvalWithOptions(_ *Frame, _ EvalOptions) error { return ErrUnsupported }

// EvalWithStorage executes the network using pre-allocated storage.
func (n *Network) EvalWithStorage(_ *FrameStorage, _ EvalOptions) error { return ErrUnsupported }

// FrameStorage holds reusable multi-frame storage.
type FrameStorage struct{}

// NewFrameStorage creates reusable storage from the given frames.
func NewFrameStorage(_ ...*Frame) *FrameStorage { return &FrameStorage{} }

// Frame holds named input and output tensors for network execution.
type Frame struct{}

// NewFrame creates an empty data frame.
func NewFrame() *Frame { return &Frame{} }

// SetInput attaches a named float32 input tensor.
func (f *Frame) SetInput(_ string, _ []float32, _ Shape) error { return ErrUnsupported }

// SetInputBytes attaches a named raw byte input.
func (f *Frame) SetInputBytes(_ string, _ []byte) error { return ErrUnsupported }

// SetInputIOSurface attaches a named input backed by an IOSurface.
func (f *Frame) SetInputIOSurface(_ string, _ uintptr) error { return ErrUnsupported }

// SetInputCVPixelBuffer attaches a named input from a CVPixelBuffer.
func (f *Frame) SetInputCVPixelBuffer(_ string, _ uintptr) error { return ErrUnsupported }

// SetInputImage attaches a named image input with raw pixel data and channel count.
func (f *Frame) SetInputImage(_ string, _ []byte, _ int) error { return ErrUnsupported }

// SetGroundTruth attaches a named ground truth tensor.
func (f *Frame) SetGroundTruth(_ string, _ []float32, _ Shape) error { return ErrUnsupported }

// Output returns the named output tensor data.
func (f *Frame) Output(_ string) ([]float32, error) { return nil, ErrUnsupported }

// OutputInto copies the named output tensor data into the caller's buffer.
func (f *Frame) OutputInto(_ string, _ []float32) error { return ErrUnsupported }

// OutputFloat64 returns the named output as float64 values.
func (f *Frame) OutputFloat64(_ string) ([]float64, error) { return nil, ErrUnsupported }

// OutputRaw returns the raw bytes of a named output.
func (f *Frame) OutputRaw(_ string) ([]byte, error) { return nil, ErrUnsupported }

// InputNames returns the names of all input attachments.
func (f *Frame) InputNames() []string { return nil }

// OutputNames returns the names of all output attachments.
func (f *Frame) OutputNames() []string { return nil }

// FunctionName returns the function name set on this frame.
func (f *Frame) FunctionName() string { return "" }

// SetFunctionName sets the function name.
func (f *Frame) SetFunctionName(_ string) {}

// Optimize applies a standard set of optimization passes.
func Optimize(_ *Network) error { return ErrUnsupported }

// OptimizeWith applies the given passes in order.
func OptimizeWith(_ *Network, _ ...Pass) error { return ErrUnsupported }

type contextConfig struct{}

// ContextOption configures a Context.
type ContextOption func(*contextConfig)

// WithPlatform selects the execution backend.
func WithPlatform(_ Platform) ContextOption { return func(*contextConfig) {} }

// WithGPUPriority sets the GPU command buffer priority.
func WithGPUPriority(_ uint32) ContextOption { return func(*contextConfig) {} }

// WithLowPriority sets the max ms per command buffer for low priority.
func WithLowPriority(_ float32) ContextOption { return func(*contextConfig) {} }

type networkConfig struct{}

// NetworkOption configures network loading.
type NetworkOption func(*networkConfig)

// WithComputePath sets the compute path.
func WithComputePath(_ int) NetworkOption { return func(*networkConfig) {} }

// WithBinSerializerID sets the binary serializer ID.
func WithBinSerializerID(_ []byte) NetworkOption { return func(*networkConfig) {} }

// EvalOptions configures execution behavior.
type EvalOptions struct {
	UseCVPixelBuffers bool
	Timeout           time.Duration
}

// MetalDevice wraps a Metal GPU device for zero-copy interop with Espresso.
type MetalDevice struct{}

// OpenMetal opens the system default Metal device.
func OpenMetal() (*MetalDevice, error) { return nil, ErrUnsupported }

// Close releases the Metal device.
func (d *MetalDevice) Close() error { return nil }

// IsMemoryTight returns whether Espresso's shared Metal singleton reports memory pressure.
func IsMemoryTight() bool { return false }

// ANESurface provides access to Espresso's ANE IOSurface for zero-copy
// Metal buffer creation and multi-buffer async execution.
type ANESurface struct{}

// NewANESurface creates a new ANE IOSurface wrapper.
func NewANESurface() *ANESurface { return nil }

// NewANESurfaceWithProperties creates an ANE IOSurface with specific properties and pixel formats.
func NewANESurfaceWithProperties(_, _ any) *ANESurface { return nil }

// SurfaceLayout describes the IOSurface dimensions required by a compiled model.
type SurfaceLayout struct {
	Width       int
	Height      int
	BytesPerRow int
	AllocSize   int
	ElemSize    int
}

// NewANESurfaceFromLayout creates an ANE IOSurface with explicit model-driven dimensions.
func NewANESurfaceFromLayout(_ SurfaceLayout) *ANESurface { return nil }

// NFrames returns the number of allocated multi-buffer frames.
func (a *ANESurface) NFrames() uint64 { return 0 }

// ResizeForAsync resizes the surface for multiple async buffer frames.
func (a *ANESurface) ResizeForAsync(_ uint64) {}

// PixelFormat returns the surface pixel format.
func (a *ANESurface) PixelFormat() uint32 { return 0 }

// IOSurfaceForFrame returns the IOSurface backing the given multi-buffer frame.
func (a *ANESurface) IOSurfaceForFrame(_ uint64) (uintptr, error) { return 0, ErrUnsupported }

// SetExternalStorage replaces the IOSurface backing a storage slot.
func (a *ANESurface) SetExternalStorage(_ uint64, _ uintptr) {}

// RestoreInternalStorage restores internal storage for a single slot.
func (a *ANESurface) RestoreInternalStorage(_ uint64) {}

// RestoreAllInternalStorage restores internal storage for all multi-buffer frames.
func (a *ANESurface) RestoreAllInternalStorage() {}

// CreateIOSurface creates a new IOSurface with optional extra properties.
func (a *ANESurface) CreateIOSurface(_ any) uintptr { return 0 }

// LazilyAllocFrame lazily allocates an IOSurface for the given frame.
func (a *ANESurface) LazilyAllocFrame(_ uint64) {}

// ForceAlloc forces non-lazy allocation.
func (a *ANESurface) ForceAlloc(_ any) {}

// MatchesCVImageBuffer returns whether the surface matches a CVImageBuffer.
func (a *ANESurface) MatchesCVImageBuffer(_ uintptr) bool { return false }

// MatchesIOSurface returns whether the surface matches an IOSurfaceRef.
func (a *ANESurface) MatchesIOSurface(_ uintptr) bool { return false }

// AliasingMem returns the external storage blob used for aliasing memory.
func (a *ANESurface) AliasingMem() any { return nil }

// SetAliasingMem sets the external storage blob for aliasing memory.
func (a *ANESurface) SetAliasingMem(_ any) {}

// WriteFrame writes raw bytes to the IOSurface backing the given frame.
func (a *ANESurface) WriteFrame(_ uint64, _ []byte) error { return ErrUnsupported }

// ReadFrame reads raw bytes from the IOSurface backing the given frame.
func (a *ANESurface) ReadFrame(_ uint64, _ []byte) error { return ErrUnsupported }

// WriteFrameF32 writes float32 data to the IOSurface backing the given frame.
func (a *ANESurface) WriteFrameF32(_ uint64, _ []float32) error { return ErrUnsupported }

// ReadFrameF32 reads float32 data from the IOSurface backing the given frame.
func (a *ANESurface) ReadFrameF32(_ uint64, _ []float32) error { return ErrUnsupported }

// Cleanup releases surface resources.
func (a *ANESurface) Cleanup() {}

// ANEProfile contains ANE-specific profiling data.
type ANEProfile struct {
	ANETimePerEvalNS  uint64
	TotalANETimeNS    uint64
	CompilerFileNames []string
}

// Profile contains execution profiling data for a network.
type Profile struct {
	Path   string
	Layers []LayerProfile
	ANE    *ANEProfile
}

// LayerProfile contains profiling data for a single layer.
type LayerProfile struct {
	Name           string
	DebugName      string
	AvgRuntimeMS   float64
	RuntimeSamples []float64
	Engine         Engine
	Supported      bool
	HasPerfWarning bool
	LayerType      string
}

// TotalRuntimeMS returns the sum of average runtimes across all layers.
func (p *Profile) TotalRuntimeMS() float64 { return 0 }

// LayersByEngine returns layers grouped by their execution engine.
func (p *Profile) LayersByEngine() map[Engine][]LayerProfile { return nil }

// UnsupportedLayers returns layers that are not supported by the main engine.
func (p *Profile) UnsupportedLayers() []LayerProfile { return nil }

// IsANEResident reports whether this layer runs on the ANE.
func (l *LayerProfile) IsANEResident() bool { return false }

// ANEResidentLayers returns layers that are resident on the ANE.
func (p *Profile) ANEResidentLayers() []LayerProfile { return nil }

// IsFullyANEResident reports whether all supported layers run on the ANE.
func (p *Profile) IsFullyANEResident() bool { return false }

// CPUFallbackLayers returns supported layers that fell back to CPU instead of ANE.
func (p *Profile) CPUFallbackLayers() []LayerProfile { return nil }

// ProfileNetwork extracts profiling data from a loaded network.
func ProfileNetwork(_ *Network) (*Profile, error) { return nil, ErrUnsupported }
