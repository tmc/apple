//go:build darwin

package ane

import (
	"fmt"
	"math"
	"runtime"
	"sync"

	"github.com/tmc/apple/corefoundation"
	"github.com/tmc/apple/coregraphics"
	"github.com/tmc/apple/foundation"
	"github.com/tmc/apple/private/appleneuralengine"
)

// Model represents a compiled and loaded model ready for evaluation.
type Model struct {
	client    *Client
	modelType ModelType

	// For MIL models.
	inMemModel appleneuralengine.ANEInMemoryModel

	// For package models.
	aneModel  appleneuralengine.ANEModel
	aneClient appleneuralengine.ANEClient

	request       appleneuralengine.ANERequest
	inputs        []coregraphics.IOSurfaceRef
	outputs       []coregraphics.IOSurfaceRef
	perfStatsMask uint32 // driver-converted mask, non-zero when stats are enabled

	// Model-driven tensor layouts parsed from compiled model attributes.
	// The source of truth for strides and channel offsets.
	inputLayouts  []TensorLayout
	outputLayouts []TensorLayout

	mu              sync.Mutex
	mapped          bool // IOSurfaces are mapped
	closed          bool
	sharedEventUsed bool
}

// InputSurface returns the i-th input IOSurfaceRef.
func (m *Model) InputSurface(i int) coregraphics.IOSurfaceRef { return m.inputs[i] }

// OutputSurface returns the i-th output IOSurfaceRef.
func (m *Model) OutputSurface(i int) coregraphics.IOSurfaceRef { return m.outputs[i] }

// InputSurfaces returns all input IOSurfaceRefs.
func (m *Model) InputSurfaces() []coregraphics.IOSurfaceRef { return m.inputs }

// OutputSurfaces returns all output IOSurfaceRefs.
func (m *Model) OutputSurfaces() []coregraphics.IOSurfaceRef { return m.outputs }

// InputAllocSize returns the IOSurface allocation size in bytes for the i-th input.
// This includes stride padding and is larger than the logical data size.
func (m *Model) InputAllocSize(i int) int { return surfaceSize(m.inputs[i]) }

// OutputAllocSize returns the IOSurface allocation size in bytes for the i-th output.
// This includes stride padding and is larger than the logical data size.
func (m *Model) OutputAllocSize(i int) int { return surfaceSize(m.outputs[i]) }

// NumInputs returns the number of input surfaces.
func (m *Model) NumInputs() int { return len(m.inputs) }

// NumOutputs returns the number of output surfaces.
func (m *Model) NumOutputs() int { return len(m.outputs) }

// InputChannels returns the channel count for the i-th input tensor.
// Returns 0 if i is out of range.
func (m *Model) InputChannels(i int) int {
	if i < len(m.inputLayouts) {
		return m.inputLayouts[i].Channels
	}
	return 0
}

// OutputChannels returns the channel count for the i-th output tensor.
// Returns 0 if i is out of range.
func (m *Model) OutputChannels(i int) int {
	if i < len(m.outputLayouts) {
		return m.outputLayouts[i].Channels
	}
	return 0
}

// Spatial returns the spatial (width) dimension for the i-th input tensor.
// Returns 0 if i is out of range.
func (m *Model) Spatial(i int) int {
	if i < len(m.inputLayouts) {
		return m.inputLayouts[i].Width
	}
	return 0
}

// InputLayout returns the tensor layout for the i-th input.
// Returns a zero TensorLayout if i is out of range.
func (m *Model) InputLayout(i int) TensorLayout {
	if i < len(m.inputLayouts) {
		return m.inputLayouts[i]
	}
	return TensorLayout{}
}

// OutputLayout returns the tensor layout for the i-th output.
// Returns a zero TensorLayout if i is out of range.
func (m *Model) OutputLayout(i int) TensorLayout {
	if i < len(m.outputLayouts) {
		return m.outputLayouts[i]
	}
	return TensorLayout{}
}

// WriteInput copies raw bytes into the i-th input IOSurface.
// len(data) must equal the surface allocation size (InputAllocSize).
// For logical (non-padded) I/O, use WriteInputF32 or WriteInputFP16 instead.
func (m *Model) WriteInput(i int, data []byte) error {
	if i >= len(m.inputs) {
		return fmt.Errorf("ane: input index %d out of range [0,%d)", i, len(m.inputs))
	}
	alloc := surfaceSize(m.inputs[i])
	if len(data) != alloc {
		return fmt.Errorf("ane: input[%d] raw data length %d, want alloc size %d", i, len(data), alloc)
	}
	return writeSurface(m.inputs[i], data)
}

// ReadOutput copies raw bytes from the i-th output IOSurface.
// len(data) must equal the surface allocation size (OutputAllocSize).
// For logical (non-padded) I/O, use ReadOutputF32 or ReadOutputFP16 instead.
func (m *Model) ReadOutput(i int, data []byte) error {
	if i >= len(m.outputs) {
		return fmt.Errorf("ane: output index %d out of range [0,%d)", i, len(m.outputs))
	}
	alloc := surfaceSize(m.outputs[i])
	if len(data) != alloc {
		return fmt.Errorf("ane: output[%d] raw data length %d, want alloc size %d", i, len(data), alloc)
	}
	return readSurface(m.outputs[i], data)
}

// WriteInputF32 writes float32 data into the i-th input IOSurface.
// Data should be in channel-first (NCHW) order: [channels * width] elements.
func (m *Model) WriteInputF32(i int, data []float32) error {
	if i >= len(m.inputs) {
		return fmt.Errorf("ane: input index %d out of range [0,%d)", i, len(m.inputs))
	}
	l := m.inputLayouts[i]
	want := l.Channels * l.Width
	if len(data) != want {
		return fmt.Errorf("ane: input[%d] f32 data length %d, want %d (%d channels x %d width)", i, len(data), want, l.Channels, l.Width)
	}
	return writeStridedF32WithLayout(m.inputs[i], data, l)
}

// ReadOutputF32 reads float32 data from the i-th output IOSurface.
// Data is returned in channel-first (NCHW) order: [channels * width] elements.
func (m *Model) ReadOutputF32(i int, data []float32) error {
	if i >= len(m.outputs) {
		return fmt.Errorf("ane: output index %d out of range [0,%d)", i, len(m.outputs))
	}
	l := m.outputLayouts[i]
	want := l.Channels * l.Width
	if len(data) != want {
		return fmt.Errorf("ane: output[%d] f32 data length %d, want %d (%d channels x %d width)", i, len(data), want, l.Channels, l.Width)
	}
	return readStridedF32WithLayout(m.outputs[i], data, l)
}

// WriteInputFP16 writes float32 data as float16 into the i-th input IOSurface.
// Data should be in channel-first (NCHW) order: [channels * width] elements.
func (m *Model) WriteInputFP16(i int, data []float32) error {
	if i >= len(m.inputs) {
		return fmt.Errorf("ane: input index %d out of range [0,%d)", i, len(m.inputs))
	}
	l := m.inputLayouts[i]
	want := l.Channels * l.Width
	if len(data) != want {
		return fmt.Errorf("ane: input[%d] fp16 data length %d, want %d (%d channels x %d width)", i, len(data), want, l.Channels, l.Width)
	}
	return writeStridedFP16WithLayout(m.inputs[i], data, l)
}

// WriteInputFP16Channels writes float32 data as float16 into a subset of the
// i-th input IOSurface starting at the given channel offset.
func (m *Model) WriteInputFP16Channels(i, channel int, data []float32) error {
	if i >= len(m.inputs) {
		return fmt.Errorf("ane: input index %d out of range [0,%d)", i, len(m.inputs))
	}
	l := m.inputLayouts[i]
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
	return writeStridedFP16ChannelsWithLayout(m.inputs[i], data, l, channel)
}

// ReadOutputFP16 reads float16 data from the i-th output IOSurface into float32s.
// Data is returned in channel-first (NCHW) order: [channels * width] elements.
func (m *Model) ReadOutputFP16(i int, data []float32) error {
	if i >= len(m.outputs) {
		return fmt.Errorf("ane: output index %d out of range [0,%d)", i, len(m.outputs))
	}
	l := m.outputLayouts[i]
	want := l.Channels * l.Width
	if len(data) != want {
		return fmt.Errorf("ane: output[%d] fp16 data length %d, want %d (%d channels x %d width)", i, len(data), want, l.Channels, l.Width)
	}
	return readStridedFP16WithLayout(m.outputs[i], data, l)
}

// ReadOutputFP16Channels reads float16 data from a subset of the i-th output
// IOSurface starting at the given channel offset into float32s.
func (m *Model) ReadOutputFP16Channels(i, channel int, data []float32) error {
	if i >= len(m.outputs) {
		return fmt.Errorf("ane: output index %d out of range [0,%d)", i, len(m.outputs))
	}
	l := m.outputLayouts[i]
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
	return readStridedFP16ChannelsWithLayout(m.outputs[i], data, l, channel)
}

// Eval executes the model on the ANE hardware.
func (m *Model) Eval() error {
	if !m.request.Validate() {
		return &ANEError{Op: "eval", Err: fmt.Errorf("request validation failed")}
	}

	switch m.modelType {
	case ModelTypeMIL:
		return m.evalMIL()
	case ModelTypePackage:
		return m.evalPackage()
	default:
		return &ANEError{Op: "eval", Err: fmt.Errorf("unknown model type %d", m.modelType)}
	}
}

func (m *Model) evalMIL() error {
	const qos = 21

	emptyOpts := foundation.NewNSMutableDictionary()
	ok, err := m.inMemModel.EvaluateWithQoSOptionsRequestError(qos, emptyOpts, m.request)
	if err == nil && ok {
		return nil
	}
	return wrapErr("eval", err)
}

func (m *Model) evalPackage() error {
	const qos = 21

	ok, err := m.aneClient.DoEvaluateDirectWithModelOptionsRequestQosError(m.aneModel, nil, m.request, qos)
	if err == nil && ok {
		return nil
	}

	ok, err = m.aneClient.EvaluateWithModelOptionsRequestQosError(m.aneModel, nil, m.request, qos)
	if err == nil && ok {
		return nil
	}
	return wrapErr("eval", err)
}

// Close releases the model's resources.
func (m *Model) Close() error {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.closed {
		return nil
	}
	m.closed = true
	runtime.SetFinalizer(m, nil)

	const qos = 21
	switch m.modelType {
	case ModelTypeMIL:
		if m.mapped {
			m.inMemModel.UnmapIOSurfacesWithRequest(m.request)
		}
		m.inMemModel.UnloadWithQoSError(qos)
	case ModelTypePackage:
		if m.mapped && !m.sharedEventUsed {
			m.aneClient.UnmapIOSurfacesWithModelRequest(m.aneModel, m.request)
		}
		if m.sharedEventUsed {
			// Package-model unload after a shared-event eval still crashes inside
			// AppleNeuralEngine on this host. Skip unload but still release IOSurfaces below.
		} else {
			m.aneClient.UnloadModelOptionsQosError(m.aneModel, nil, qos)
		}
	}

	// Release all IOSurface refs. IOSurfaceCreate returns +1; balance here.
	for _, ref := range m.inputs {
		if ref != 0 {
			corefoundation.CFRelease(corefoundation.CFTypeRef(ref))
		}
	}
	for _, ref := range m.outputs {
		if ref != 0 {
			corefoundation.CFRelease(corefoundation.CFTypeRef(ref))
		}
	}
	m.inputs = nil
	m.outputs = nil

	return nil
}

// ModelObjcID returns the ObjC object pointer for the underlying ANEModel.
// For MIL models this returns the inner model; for package models the direct model.
func (m *Model) ModelObjcID() uintptr {
	switch m.modelType {
	case ModelTypeMIL:
		if inner := m.inMemModel.Model(); inner != nil {
			return uintptr(inner.GetID())
		}
		return uintptr(m.inMemModel.ID)
	case ModelTypePackage:
		return uintptr(m.aneModel.ID)
	}
	return 0
}

// InMemModelObjcID returns the ObjC object pointer for the ANEInMemoryModel wrapper.
// Returns 0 for non-MIL models.
func (m *Model) InMemModelObjcID() uintptr { return uintptr(m.inMemModel.ID) }

// CompileModelType returns the model type used during compilation.
func (m *Model) CompileModelType() ModelType { return m.modelType }

// RawRequest returns the ObjC object pointer for the underlying ANERequest.
func (m *Model) RawRequest() uintptr { return uintptr(m.request.ID) }

// RawPerfStatsMask returns the driver-converted performance stats mask.
func (m *Model) RawPerfStatsMask() uint32 { return m.perfStatsMask }

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
