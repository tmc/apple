# x/ane Design

## Package Layout

```
x/ane/            # core: Open, Probe, Runtime, Kernel, errors
x/ane/mil/        # MIL text generation + weight blob building
x/ane/linear/     # cached linear operator (kernel-per-shape)
```

No separate `x/ane/iosurface`. Surface I/O uses `github.com/tmc/apple/iosurface`
and `github.com/tmc/apple/coregraphics` directly. Surface creation and strided
I/O are model-driven: layouts are parsed from compiled model attributes, not
guessed from allocation sizes.

## Core Concepts

### Model-Driven Surface Layout

After compiling a model, the ANE reports its expected memory layout via
`ANEModel.ModelAttributes()` → `NetworkStatusList` → `LiveInputList` /
`LiveOutputList`. Each tensor entry has:

```
Channels, Width, Height, PlaneStride, RowStride, Type (Float16/Float32)
```

IOSurfaces are created to match these layouts exactly:

```
IOSurfaceWidth      = Width (spatial dimension)
IOSurfaceHeight     = Channels
IOSurfaceBytesPerElement = ElemSize (2 for fp16, 4 for fp32)
IOSurfaceBytesPerRow = RowStride (from model attributes)
IOSurfaceAllocSize  = Channels * PlaneStride
```

For [1, C, 1, S] tensors (Height=1), PlaneStride == RowStride and each
channel occupies one row. The minimum row stride is 64 bytes (the ANE's
alignment granularity). Strides are never inferred from allocation sizes.

**Height > 1 tensors are not supported.** Compile returns `ErrUnsupportedLayout`
if the model attributes report Height > 1 for any input or output tensor.

### Logical vs Allocation Size

The allocation size (`TensorLayout.AllocSize()`) includes stride padding and
is always >= the logical data size. The logical size is the dense data without
padding:

```
AllocSize     = Channels * PlaneStride          (includes padding)
LogicalBytes  = Channels * Width * Height * ElemSize  (no padding)
LogicalElements = Channels * Width * Height
```

`Kernel.InputAllocSize` / `OutputAllocSize` return the IOSurface allocation
size. Typed I/O methods (`WriteInputF32`, `ReadOutputFP16`, etc.) accept
`LogicalElements` values and handle stride padding internally. Raw I/O
(`WriteInput`, `ReadOutput`) requires exactly `AllocSize` bytes.

### Compile Flow

```
1. Create model descriptor from MIL text + weights
2. Compile and load (ANEInMemoryModel or ANEClient)
3. Parse model attributes → TensorLayout per input/output
4. Reject unsupported layouts (Height > 1)
5. Create IOSurfaces matching parsed layouts
6. Map IOSurfaces (with cache fallback)
7. Return Kernel with layouts stored
```

### Strided I/O

Data is in channel-first (NCHW) order: `data[c * width + x]`.
Byte offset in the IOSurface: `c * PlaneStride + x * ElemSize`.

## API

### TensorLayout

```go
type TensorLayout struct {
    Channels    int // number of channels (C in NCHW)
    Width       int // spatial dimension (S in NCHW)
    Height      int // typically 1; Height > 1 not yet supported
    ElemSize    int // bytes per element (2 for fp16, 4 for fp32)
    RowStride   int // byte stride between rows (64-byte aligned)
    PlaneStride int // byte stride between channels
}

func (l TensorLayout) AllocSize() int       // Channels * PlaneStride
func (l TensorLayout) LogicalBytes() int    // Channels * Width * Height * ElemSize
func (l TensorLayout) LogicalElements() int // Channels * Width * Height
```

### Runtime

```go
func Probe() (DeviceInfo, error)
func Open() (*Runtime, error)
func (rt *Runtime) Close() error
func (rt *Runtime) Info() DeviceInfo
func (rt *Runtime) CompileCount() int64
func (rt *Runtime) Compile(opts CompileOptions) (*Kernel, error)
```

### Kernel

```go
// Tensor metadata (from compiled model attributes).
func (k *Kernel) InputChannels(i int) int
func (k *Kernel) OutputChannels(i int) int
func (k *Kernel) Spatial(i int) int
func (k *Kernel) InputLayout(i int) TensorLayout
func (k *Kernel) OutputLayout(i int) TensorLayout

// Surface access.
func (k *Kernel) InputSurface(i int) coregraphics.IOSurfaceRef
func (k *Kernel) OutputSurface(i int) coregraphics.IOSurfaceRef
func (k *Kernel) NumInputs() int
func (k *Kernel) NumOutputs() int
func (k *Kernel) InputAllocSize(i int) int
func (k *Kernel) OutputAllocSize(i int) int

// Typed I/O with shape validation (LogicalElements count).
func (k *Kernel) WriteInputF32(i int, data []float32) error
func (k *Kernel) ReadOutputF32(i int, data []float32) error
func (k *Kernel) WriteInputFP16(i int, data []float32) error
func (k *Kernel) ReadOutputFP16(i int, data []float32) error

// Raw I/O (requires exactly AllocSize bytes).
func (k *Kernel) WriteInput(i int, data []byte) error
func (k *Kernel) ReadOutput(i int, data []byte) error

// Evaluation.
func (k *Kernel) Eval() error
func (k *Kernel) EvalWithStats() (EvalStats, error)

// Shared event evaluation (ModelTypePackage only).
func (k *Kernel) EvalWithSignalEvent(signalPort uint32, signalValue uint64, cfg SharedEventEvalOptions) error
func (k *Kernel) EvalBidirectional(waitPort, waitValue, signalPort, signalValue uint32/uint64, cfg SharedEventEvalOptions) error

func (k *Kernel) Close() error
```

### CompileOptions

```go
type CompileOptions struct {
    ModelType  ModelType
    MILText    []byte // MIL program text
    WeightBlob []byte // weight binary blob
    WeightPath string // weight dict key (default: "@model_path/weights/weight.bin")
    PackagePath string // .mlmodelc path (ModelTypePackage)
    ModelKey    string // model dictionary key (default: "s")
    QoS           uint32 // quality-of-service class (0 = default 21)
    PerfStatsMask uint32 // perf stats bitmask (0 = disabled)
}
```

Surface sizes, channels, and spatial dimensions are no longer specified
in CompileOptions. They are parsed from the compiled model's attributes.

### Errors

```go
var (
    ErrNoANE                    // no ANE hardware
    ErrCompileBudgetExhausted   // ~119 per-process compile limit reached
    ErrMapFailed                // IOSurface mapping failed
    ErrUnsupportedSelector      // selector not available on this OS version
    ErrVirtualClientUnavailable // virtual client not available
    ErrModelLoad                // model load failed
    ErrEval                     // evaluation failed
    ErrUnsupportedLayout        // tensor layout not supported (e.g. Height > 1)
)

type ANEError struct { Op, Class, Code, Err }
```

## Compilation Paths

### ModelTypeMIL (in-memory)

```
ANEInMemoryModelDescriptor → ANEInMemoryModel → CompileWithQoS →
LoadWithQoS → model.Model() → parseModelLayouts → createSurfaces →
MapIOSurfaces → Kernel
```

Eval: `ANEInMemoryModel.EvaluateWithQoS`

### ModelTypePackage (on-disk .mlmodelc)

```
ANEModel.ModelAtURLKey → ANEClient.CompileModel → LoadModel →
parseModelLayouts → createSurfaces → MapIOSurfaces → Kernel
```

Eval: `ANEClient.DoEvaluateDirectWithModel` → fallback `EvaluateWithModel`
