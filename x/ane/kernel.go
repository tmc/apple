//go:build darwin

package ane

import (
	"fmt"
	"math"

	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/private/appleneuralengine"
)

// Kernel represents a compiled and loaded model ready for evaluation.
type Kernel struct {
	rt        *Runtime
	modelType ModelType

	// For MIL models.
	inMemModel appleneuralengine.ANEInMemoryModel

	// For package models.
	model  appleneuralengine.ANEModel
	client appleneuralengine.ANEClient

	request appleneuralengine.ANERequest
	inputs  []coregraphics.IOSurfaceRef
	outputs []coregraphics.IOSurfaceRef

	// Model-driven tensor layouts parsed from compiled model attributes.
	// The source of truth for strides and channel offsets.
	inputLayouts  []TensorLayout
	outputLayouts []TensorLayout

	mapped bool // IOSurfaces are mapped
	closed bool
}

// InputSurface returns the i-th input IOSurfaceRef.
func (k *Kernel) InputSurface(i int) coregraphics.IOSurfaceRef { return k.inputs[i] }

// OutputSurface returns the i-th output IOSurfaceRef.
func (k *Kernel) OutputSurface(i int) coregraphics.IOSurfaceRef { return k.outputs[i] }

// InputSurfaces returns all input IOSurfaceRefs.
func (k *Kernel) InputSurfaces() []coregraphics.IOSurfaceRef { return k.inputs }

// OutputSurfaces returns all output IOSurfaceRefs.
func (k *Kernel) OutputSurfaces() []coregraphics.IOSurfaceRef { return k.outputs }

// InputAllocSize returns the IOSurface allocation size in bytes for the i-th input.
// This includes stride padding and is larger than the logical data size.
func (k *Kernel) InputAllocSize(i int) int { return surfaceSize(k.inputs[i]) }

// OutputAllocSize returns the IOSurface allocation size in bytes for the i-th output.
// This includes stride padding and is larger than the logical data size.
func (k *Kernel) OutputAllocSize(i int) int { return surfaceSize(k.outputs[i]) }

// NumInputs returns the number of input surfaces.
func (k *Kernel) NumInputs() int { return len(k.inputs) }

// NumOutputs returns the number of output surfaces.
func (k *Kernel) NumOutputs() int { return len(k.outputs) }

// InputChannels returns the channel count for the i-th input tensor.
// Returns 0 if i is out of range.
func (k *Kernel) InputChannels(i int) int {
	if i < len(k.inputLayouts) {
		return k.inputLayouts[i].Channels
	}
	return 0
}

// OutputChannels returns the channel count for the i-th output tensor.
// Returns 0 if i is out of range.
func (k *Kernel) OutputChannels(i int) int {
	if i < len(k.outputLayouts) {
		return k.outputLayouts[i].Channels
	}
	return 0
}

// Spatial returns the spatial (width) dimension for the i-th input tensor.
// Returns 0 if i is out of range.
func (k *Kernel) Spatial(i int) int {
	if i < len(k.inputLayouts) {
		return k.inputLayouts[i].Width
	}
	return 0
}

// InputLayout returns the tensor layout for the i-th input.
// Returns a zero TensorLayout if i is out of range.
func (k *Kernel) InputLayout(i int) TensorLayout {
	if i < len(k.inputLayouts) {
		return k.inputLayouts[i]
	}
	return TensorLayout{}
}

// OutputLayout returns the tensor layout for the i-th output.
// Returns a zero TensorLayout if i is out of range.
func (k *Kernel) OutputLayout(i int) TensorLayout {
	if i < len(k.outputLayouts) {
		return k.outputLayouts[i]
	}
	return TensorLayout{}
}

// WriteInput copies raw bytes into the i-th input IOSurface.
// len(data) must equal the surface allocation size (InputAllocSize).
// For logical (non-padded) I/O, use WriteInputF32 or WriteInputFP16 instead.
func (k *Kernel) WriteInput(i int, data []byte) error {
	if i >= len(k.inputs) {
		return fmt.Errorf("ane: input index %d out of range [0,%d)", i, len(k.inputs))
	}
	alloc := surfaceSize(k.inputs[i])
	if len(data) != alloc {
		return fmt.Errorf("ane: input[%d] raw data length %d, want alloc size %d", i, len(data), alloc)
	}
	return writeSurface(k.inputs[i], data)
}

// ReadOutput copies raw bytes from the i-th output IOSurface.
// len(data) must equal the surface allocation size (OutputAllocSize).
// For logical (non-padded) I/O, use ReadOutputF32 or ReadOutputFP16 instead.
func (k *Kernel) ReadOutput(i int, data []byte) error {
	if i >= len(k.outputs) {
		return fmt.Errorf("ane: output index %d out of range [0,%d)", i, len(k.outputs))
	}
	alloc := surfaceSize(k.outputs[i])
	if len(data) != alloc {
		return fmt.Errorf("ane: output[%d] raw data length %d, want alloc size %d", i, len(data), alloc)
	}
	return readSurface(k.outputs[i], data)
}

// WriteInputF32 writes float32 data into the i-th input IOSurface.
// Data should be in channel-first (NCHW) order: [channels * width] elements.
func (k *Kernel) WriteInputF32(i int, data []float32) error {
	if i >= len(k.inputs) {
		return fmt.Errorf("ane: input index %d out of range [0,%d)", i, len(k.inputs))
	}
	l := k.inputLayouts[i]
	want := l.Channels * l.Width
	if len(data) != want {
		return fmt.Errorf("ane: input[%d] f32 data length %d, want %d (%d channels x %d width)", i, len(data), want, l.Channels, l.Width)
	}
	return writeStridedF32WithLayout(k.inputs[i], data, l)
}

// ReadOutputF32 reads float32 data from the i-th output IOSurface.
// Data is returned in channel-first (NCHW) order: [channels * width] elements.
func (k *Kernel) ReadOutputF32(i int, data []float32) error {
	if i >= len(k.outputs) {
		return fmt.Errorf("ane: output index %d out of range [0,%d)", i, len(k.outputs))
	}
	l := k.outputLayouts[i]
	want := l.Channels * l.Width
	if len(data) != want {
		return fmt.Errorf("ane: output[%d] f32 data length %d, want %d (%d channels x %d width)", i, len(data), want, l.Channels, l.Width)
	}
	return readStridedF32WithLayout(k.outputs[i], data, l)
}

// WriteInputFP16 writes float32 data as float16 into the i-th input IOSurface.
// Data should be in channel-first (NCHW) order: [channels * width] elements.
func (k *Kernel) WriteInputFP16(i int, data []float32) error {
	if i >= len(k.inputs) {
		return fmt.Errorf("ane: input index %d out of range [0,%d)", i, len(k.inputs))
	}
	l := k.inputLayouts[i]
	want := l.Channels * l.Width
	if len(data) != want {
		return fmt.Errorf("ane: input[%d] fp16 data length %d, want %d (%d channels x %d width)", i, len(data), want, l.Channels, l.Width)
	}
	return writeStridedFP16WithLayout(k.inputs[i], data, l)
}

// WriteInputFP16Channels writes float32 data as float16 into a subset of the
// i-th input IOSurface starting at the given channel offset.
func (k *Kernel) WriteInputFP16Channels(i, channel int, data []float32) error {
	if i >= len(k.inputs) {
		return fmt.Errorf("ane: input index %d out of range [0,%d)", i, len(k.inputs))
	}
	l := k.inputLayouts[i]
	channelElems := l.Height * l.Width
	if channelElems <= 0 {
		return fmt.Errorf("ane: input[%d] has invalid channel size %d", i, channelElems)
	}
	if len(data)%channelElems != 0 {
		return fmt.Errorf("ane: input[%d] fp16 channel data length %d is not a multiple of %d", i, len(data), channelElems)
	}
	channels := len(data) / channelElems
	if channel < 0 || channel+channels > l.Channels {
		return fmt.Errorf("ane: input[%d] channel range [%d,%d) out of range [0,%d)", i, channel, channel+channels, l.Channels)
	}
	return writeStridedFP16ChannelsWithLayout(k.inputs[i], data, l, channel)
}

// ReadOutputFP16 reads float16 data from the i-th output IOSurface into float32s.
// Data is returned in channel-first (NCHW) order: [channels * width] elements.
func (k *Kernel) ReadOutputFP16(i int, data []float32) error {
	if i >= len(k.outputs) {
		return fmt.Errorf("ane: output index %d out of range [0,%d)", i, len(k.outputs))
	}
	l := k.outputLayouts[i]
	want := l.Channels * l.Width
	if len(data) != want {
		return fmt.Errorf("ane: output[%d] fp16 data length %d, want %d (%d channels x %d width)", i, len(data), want, l.Channels, l.Width)
	}
	return readStridedFP16WithLayout(k.outputs[i], data, l)
}

// ReadOutputFP16Channels reads float16 data from a subset of the i-th output
// IOSurface starting at the given channel offset into float32s.
func (k *Kernel) ReadOutputFP16Channels(i, channel int, data []float32) error {
	if i >= len(k.outputs) {
		return fmt.Errorf("ane: output index %d out of range [0,%d)", i, len(k.outputs))
	}
	l := k.outputLayouts[i]
	channelElems := l.Height * l.Width
	if channelElems <= 0 {
		return fmt.Errorf("ane: output[%d] has invalid channel size %d", i, channelElems)
	}
	if len(data)%channelElems != 0 {
		return fmt.Errorf("ane: output[%d] fp16 channel data length %d is not a multiple of %d", i, len(data), channelElems)
	}
	channels := len(data) / channelElems
	if channel < 0 || channel+channels > l.Channels {
		return fmt.Errorf("ane: output[%d] channel range [%d,%d) out of range [0,%d)", i, channel, channel+channels, l.Channels)
	}
	return readStridedFP16ChannelsWithLayout(k.outputs[i], data, l, channel)
}

// Eval executes the kernel on the ANE hardware.
func (k *Kernel) Eval() error {
	if !k.request.Validate() {
		return &ANEError{Op: "eval", Err: fmt.Errorf("request validation failed")}
	}

	switch k.modelType {
	case ModelTypeMIL:
		return k.evalMIL()
	case ModelTypePackage:
		return k.evalPackage()
	default:
		return &ANEError{Op: "eval", Err: fmt.Errorf("unknown model type %d", k.modelType)}
	}
}

// EvalWithStats executes the kernel and returns hardware performance stats.
// If perf stats were not enabled at compile time (PerfStatsMask=0), the
// returned stats will be zero-valued.
func (k *Kernel) EvalWithStats() (EvalStats, error) {
	perfStats := appleneuralengine.NewANEPerformanceStats()
	k.request.SetPerfStats(&perfStats)

	if err := k.Eval(); err != nil {
		return EvalStats{}, err
	}

	var hwNS uint64
	func() {
		defer func() { recover() }()
		hwNS = perfStats.HwExecutionTime()
	}()

	return EvalStats{HWExecutionNS: hwNS}, nil
}

func (k *Kernel) evalMIL() error {
	const qos = 21

	emptyOpts := foundation.NewNSMutableDictionary()
	ok, err := k.inMemModel.EvaluateWithQoSOptionsRequestError(qos, emptyOpts, k.request)
	if err == nil && ok {
		return nil
	}
	return wrapErr("eval", err)
}

func (k *Kernel) evalPackage() error {
	const qos = 21

	ok, err := k.client.DoEvaluateDirectWithModelOptionsRequestQosError(k.model, nil, k.request, qos)
	if err == nil && ok {
		return nil
	}

	ok, err = k.client.EvaluateWithModelOptionsRequestQosError(k.model, nil, k.request, qos)
	if err == nil && ok {
		return nil
	}
	return wrapErr("eval", err)
}

// Close releases the kernel's resources.
func (k *Kernel) Close() error {
	if k.closed {
		return nil
	}
	k.closed = true

	const qos = 21
	switch k.modelType {
	case ModelTypeMIL:
		if k.mapped {
			k.inMemModel.UnmapIOSurfacesWithRequest(k.request)
		}
		k.inMemModel.UnloadWithQoSError(qos)
	case ModelTypePackage:
		if k.mapped {
			k.client.UnmapIOSurfacesWithModelRequest(k.model, k.request)
		}
		k.client.UnloadModelOptionsQosError(k.model, nil, qos)
	}

	return nil
}

// Float32ToFP16 converts a float32 to IEEE 754 half-precision.
func Float32ToFP16(f float32) uint16 {
	b := math.Float32bits(f)
	sign := (b >> 16) & 0x8000
	exp := int((b>>23)&0xFF) - 127 + 15
	frac := b & 0x7FFFFF

	switch {
	case exp <= 0:
		return uint16(sign)
	case exp >= 31:
		return uint16(sign | 0x7C00)
	default:
		return uint16(sign | uint32(exp)<<10 | (frac >> 13))
	}
}

// FP16ToFloat32 converts an IEEE 754 half-precision value to float32.
func FP16ToFloat32(h uint16) float32 {
	sign := uint32(h>>15) & 1
	exp := uint32(h>>10) & 0x1F
	frac := uint32(h) & 0x3FF

	switch {
	case exp == 0:
		if frac == 0 {
			return math.Float32frombits(sign << 31)
		}
		for frac&0x400 == 0 {
			frac <<= 1
			exp--
		}
		exp++
		frac &= 0x3FF
		fallthrough
	case exp < 31:
		return math.Float32frombits(sign<<31 | (exp+127-15)<<23 | frac<<13)
	default:
		return math.Float32frombits(sign<<31 | 0x7F800000 | frac<<13)
	}
}
