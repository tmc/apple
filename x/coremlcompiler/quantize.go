package coremlcompiler

import (
	"fmt"
	"math"
)

// An AffineDequantization describes a weight tensor stored quantized to 8 bits
// and the parameters that recover it, following the equation
// constexpr_affine_dequantize is defined by (ops/defs/iOS16/constexpr_ops.py):
//
//	unquantized = scale * (quantized_data - zero_point)
//
// Scale and ZeroPoint are either one value each, quantizing the whole tensor,
// or one per slice along Axis. Per-output-channel quantization of a projection
// matrix shaped [out, in] is the latter with Axis 0.
type AffineDequantization struct {
	// Shape is the shape of the quantized data, which is also the shape of
	// the dequantized output.
	Shape []int64

	// QuantizedType is the stored element type: [DataTypeInt8] or
	// [DataTypeUInt8].
	QuantizedType DataType

	// OutputType is the type the data dequantizes to: [DataTypeFloat16] or
	// [DataTypeFloat32].
	OutputType DataType

	// Axis is the dimension Scale and ZeroPoint vary along. It is ignored
	// when they hold a single value.
	Axis int32

	// Scale and ZeroPoint hold either one value, or Shape[Axis] values.
	// They must be the same length as each other.
	Scale     []float32
	ZeroPoint []int64
}

// Validate reports whether q satisfies constexpr_affine_dequantize's
// constraints on its parameters.
func (q AffineDequantization) Validate() error {
	if len(q.Shape) == 0 {
		return fmt.Errorf("coremlcompiler: quantized data has no shape")
	}
	if _, err := quantizedElementCount(q.Shape); err != nil {
		return err
	}
	switch q.QuantizedType {
	case DataTypeInt8, DataTypeUInt8:
	default:
		return fmt.Errorf("coremlcompiler: quantized type %s is not int8 or uint8", q.QuantizedType)
	}
	switch q.OutputType {
	case DataTypeFloat16, DataTypeFloat32:
	default:
		return fmt.Errorf("coremlcompiler: output type %s is not fp16 or fp32", q.OutputType)
	}
	if len(q.Scale) == 0 {
		return fmt.Errorf("coremlcompiler: no scale")
	}
	if len(q.ZeroPoint) != len(q.Scale) {
		return fmt.Errorf("coremlcompiler: %d scales but %d zero points", len(q.Scale), len(q.ZeroPoint))
	}
	for i, s := range q.Scale {
		if s == 0 || math.IsNaN(float64(s)) || math.IsInf(float64(s), 0) {
			return fmt.Errorf("coremlcompiler: scale %d is %v, which does not describe a recoverable tensor", i, s)
		}
		if q.OutputType == DataTypeFloat16 {
			h := Float16bits(s)
			if h&0x7fff == 0 || h&0x7c00 == 0x7c00 {
				return fmt.Errorf("coremlcompiler: scale %d is %v, which fp16 cannot hold as a finite nonzero value", i, s)
			}
		}
	}
	min, max := quantizedRange(q.QuantizedType)
	for i, z := range q.ZeroPoint {
		if z < min || z > max {
			return fmt.Errorf("coremlcompiler: zero point %d is %d, outside %s's range [%d, %d]", i, z, q.QuantizedType, min, max)
		}
	}
	if len(q.Scale) == 1 {
		return nil // scalar scale broadcasts over the whole tensor
	}
	if q.Axis < 0 || int(q.Axis) >= len(q.Shape) {
		return fmt.Errorf("coremlcompiler: axis %d is outside the rank-%d quantized data", q.Axis, len(q.Shape))
	}
	// "size(scale-vector) == quantized_data.shape[axis]".
	if want := q.Shape[q.Axis]; int64(len(q.Scale)) != want {
		return fmt.Errorf("coremlcompiler: %d scales for axis %d of length %d", len(q.Scale), q.Axis, want)
	}
	return nil
}

func quantizedRange(dt DataType) (min, max int64) {
	if dt == DataTypeUInt8 {
		return 0, 255
	}
	return -128, 127
}

// NumElements returns the number of elements the quantized tensor holds.
// It returns zero if the shape is invalid or its element count overflows int.
func (q AffineDequantization) NumElements() int {
	if len(q.Shape) == 0 {
		return 0
	}
	n, _ := quantizedElementCount(q.Shape)
	return n
}

func quantizedElementCount(shape []int64) (int, error) {
	n := 1
	for _, d := range shape {
		if d < 1 {
			return 0, fmt.Errorf("coremlcompiler: quantized data shape %v has a non-positive dimension", shape)
		}
		if uint64(d) > uint64(int(^uint(0)>>1)/n) {
			return 0, fmt.Errorf("coremlcompiler: quantized data shape %v is too large", shape)
		}
		n *= int(d)
	}
	return n, nil
}

// Operation returns the constexpr_affine_dequantize operation producing a
// tensor named name from quantized data held at data.
//
// The parameters go in the operation's attributes, not its inputs, whatever
// opset the enclosing function declares. coremltools chooses between the two
// on op.opset_version -- the opset the op class itself was declared at, not
// the program's (mil/operation.py:571-577, backend/mil/load.py
// translate_constexpr). constexpr_affine_dequantize is declared only in
// ops/defs/iOS16 and is not among the constexpr ops iOS18 redefines, so it is
// always an iOS16 op and always takes attributes. The Core ML runtime agrees
// and names it that way: a program that puts these in inputs is rejected with
// "Attribute quantized_data is undefined for ios16.constexpr_affine_dequantize".
//
// The constexpr ops introduced at iOS18 -- constexpr_blockwise_shift_scale and
// the rest of ops/defs/iOS18/compression.py -- are the ones that take inputs.
func (q AffineDequantization) Operation(name string, data BlobRef) (*Operation, error) {
	if name == "" {
		return nil, fmt.Errorf("coremlcompiler: constexpr_affine_dequantize has no output name")
	}
	if err := q.Validate(); err != nil {
		return nil, err
	}
	if data.Path == "" {
		return nil, fmt.Errorf("coremlcompiler: %s: quantized data has no blob path", name)
	}

	quantized := &Value{
		Type:     tensorValueType(q.QuantizedType, q.Shape),
		BlobFile: &BlobFileValue{FileName: data.Path, Offset: data.Offset},
	}
	scale, err := q.scaleValue()
	if err != nil {
		return nil, fmt.Errorf("coremlcompiler: %s scale: %w", name, err)
	}
	zeroPoint, err := immediateTensor(q.QuantizedType, q.zeroPointBytes(), q.paramShape())
	if err != nil {
		return nil, fmt.Errorf("coremlcompiler: %s zero_point: %w", name, err)
	}
	axis := &Value{
		Type:      tensorValueType(DataTypeInt32, nil),
		Immediate: &ImmediateValue{Tensor: &TensorValue{Ints: []int32{q.Axis}}},
	}

	params := map[string]*Value{
		"quantized_data": quantized,
		"zero_point":     zeroPoint,
		"scale":          scale,
		"axis":           axis,
	}
	return &Operation{
		Type:       "constexpr_affine_dequantize",
		Attributes: params,
		Outputs: []NamedValueType{{
			Name: name,
			Type: tensorValueType(q.OutputType, q.Shape),
		}},
	}, nil
}

// paramShape is the shape of the scale and zero_point tensors: a scalar when
// one value covers the tensor, otherwise a vector.
func (q AffineDequantization) paramShape() []int64 {
	if len(q.Scale) == 1 {
		return nil
	}
	return []int64{int64(len(q.Scale))}
}

// scaleValue builds the inline scale tensor. Which TensorValue field carries
// it is fixed by the element type, not free choice: fp16 goes in Bytes as
// packed halves and fp32 goes in Floats, and putting either in the other field
// is read as an empty value rather than rejected.
func (q AffineDequantization) scaleValue() (*Value, error) {
	switch q.OutputType {
	case DataTypeFloat16:
		return immediateTensor(DataTypeFloat16, Float16Bytes(q.Scale), q.paramShape())
	case DataTypeFloat32:
		return &Value{
			Type: tensorValueType(DataTypeFloat32, q.paramShape()),
			Immediate: &ImmediateValue{Tensor: &TensorValue{
				Floats: append([]float32(nil), q.Scale...),
			}},
		}, nil
	}
	return nil, fmt.Errorf("element type %s is not fp16 or fp32", q.OutputType)
}

func (q AffineDequantization) zeroPointBytes() []byte {
	out := make([]byte, len(q.ZeroPoint))
	for i, z := range q.ZeroPoint {
		out[i] = byte(int8(z))
		if q.QuantizedType == DataTypeUInt8 {
			out[i] = byte(z)
		}
	}
	return out
}

func tensorValueType(dt DataType, shape []int64) *ValueType {
	dims := make([]Dimension, len(shape))
	for i, d := range shape {
		dims[i] = Dimension{Constant: uint64(d)}
	}
	return &ValueType{TensorType: &TensorType{DataType: dt, Rank: int64(len(shape)), Dimensions: dims}}
}

// immediateTensor builds an inline tensor constant, putting the payload in the
// one TensorValue field the MIL reader consults for dt.
func immediateTensor(dt DataType, data []byte, shape []int64) (*Value, error) {
	field, err := dt.FieldForDataType()
	if err != nil {
		return nil, err
	}
	if field != TensorValueBytes {
		return nil, fmt.Errorf("element type %s is not carried as raw bytes", dt)
	}
	return &Value{
		Type:      tensorValueType(dt, shape),
		Immediate: &ImmediateValue{Tensor: &TensorValue{Bytes: data}},
	}, nil
}

// QuantizeAffinePerChannel quantizes values, laid out row-major in shape, to
// 8 bits with one scale and zero point per slice along axis. It returns the
// descriptor and the packed quantized payload, ready for a [WeightTensor].
//
// Quantizing a projection matrix per output channel means axis 0 of an
// [out, in] matrix: each output row gets the scale and zero point that spread
// its own range across the full 8-bit range.
func QuantizeAffinePerChannel(values []float32, shape []int64, axis int32, quantizedType DataType) (AffineDequantization, []byte, error) {
	q := AffineDequantization{
		Shape:         append([]int64(nil), shape...),
		QuantizedType: quantizedType,
		OutputType:    DataTypeFloat16,
		Axis:          axis,
	}
	if len(shape) == 0 {
		return q, nil, fmt.Errorf("coremlcompiler: quantize: no shape")
	}
	if axis < 0 || int(axis) >= len(shape) {
		return q, nil, fmt.Errorf("coremlcompiler: quantize: axis %d is outside the rank-%d tensor", axis, len(shape))
	}
	n, err := quantizedElementCount(shape)
	if err != nil {
		return q, nil, err
	}
	if len(values) != n {
		return q, nil, fmt.Errorf("coremlcompiler: quantize: %d values for shape %v, want %d", len(values), shape, n)
	}

	channels := int(shape[axis])
	// Elements of one slice along axis are strided: inner is the size of the
	// dimensions after axis, and the axis index repeats every inner elements.
	inner := 1
	for _, d := range shape[axis+1:] {
		inner *= int(d)
	}

	qMin, qMax := quantizedRange(quantizedType)
	if quantizedType != DataTypeInt8 && quantizedType != DataTypeUInt8 {
		return q, nil, fmt.Errorf("coremlcompiler: quantize: %s is not int8 or uint8", quantizedType)
	}

	mins := make([]float32, channels)
	maxs := make([]float32, channels)
	for i := range mins {
		mins[i], maxs[i] = float32(math.Inf(1)), float32(math.Inf(-1))
	}
	channelOf := func(i int) int { return i / inner % channels }
	for i, v := range values {
		if math.IsNaN(float64(v)) || math.IsInf(float64(v), 0) {
			return q, nil, fmt.Errorf("coremlcompiler: quantize: value %d is %v", i, v)
		}
		c := channelOf(i)
		if v < mins[c] {
			mins[c] = v
		}
		if v > maxs[c] {
			maxs[c] = v
		}
	}

	q.Scale = make([]float32, channels)
	q.ZeroPoint = make([]int64, channels)
	for c := range q.Scale {
		lo, hi := float64(mins[c]), float64(maxs[c])
		// Include zero, so an all-positive or all-negative channel still
		// represents it exactly; a zero-width channel gets unit scale rather
		// than a scale of zero, which would be unrecoverable.
		lo, hi = math.Min(lo, 0), math.Max(hi, 0)
		scale := (hi - lo) / float64(qMax-qMin)
		if scale == 0 {
			scale = 1
		}
		zp := int64(math.RoundToEven(float64(qMin) - lo/scale))
		zp = clampInt64(zp, qMin, qMax)
		s32 := float32(scale)
		if s32 == 0 || math.IsInf(float64(s32), 0) {
			return q, nil, fmt.Errorf("coremlcompiler: quantize: channel %d has scale %v, which fp32 cannot hold", c, scale)
		}
		q.Scale[c] = s32
		q.ZeroPoint[c] = zp
	}

	packed := make([]byte, len(values))
	for i, v := range values {
		c := channelOf(i)
		x := int64(math.RoundToEven(float64(v)/float64(q.Scale[c]))) + q.ZeroPoint[c]
		x = clampInt64(x, qMin, qMax)
		if quantizedType == DataTypeUInt8 {
			packed[i] = byte(x)
		} else {
			packed[i] = byte(int8(x))
		}
	}
	if err := q.Validate(); err != nil {
		return q, nil, err
	}
	return q, packed, nil
}

// Dequantize recovers the values q describes from its packed payload. It is
// the inverse of [QuantizeAffinePerChannel] up to quantization error, and
// exists so a caller can measure that error before shipping a model.
func (q AffineDequantization) Dequantize(packed []byte) ([]float32, error) {
	if err := q.Validate(); err != nil {
		return nil, err
	}
	if len(packed) != q.NumElements() {
		return nil, fmt.Errorf("coremlcompiler: %d quantized bytes for shape %v, want %d", len(packed), q.Shape, q.NumElements())
	}
	inner, channels := 1, 1
	if len(q.Scale) > 1 {
		for _, d := range q.Shape[q.Axis+1:] {
			inner *= int(d)
		}
		channels = int(q.Shape[q.Axis])
	}
	out := make([]float32, len(packed))
	for i, b := range packed {
		c := 0
		if len(q.Scale) > 1 {
			c = i / inner % channels
		}
		var x int64
		if q.QuantizedType == DataTypeUInt8 {
			x = int64(b)
		} else {
			x = int64(int8(b))
		}
		out[i] = q.Scale[c] * float32(x-q.ZeroPoint[c])
	}
	return out, nil
}

func clampInt64(v, lo, hi int64) int64 {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
