package coremlcompiler

import (
	"math"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// projectionWeights builds an [out, in] matrix whose rows have deliberately
// different ranges, the case per-output-channel quantization exists for: one
// scale for the whole matrix would waste almost all of the small rows' range.
func projectionWeights(out, in int) []float32 {
	vals := make([]float32, out*in)
	for r := 0; r < out; r++ {
		scale := float32(math.Pow(10, float64(r)-2)) // 0.01, 0.1, 1, 10, ...
		for c := 0; c < in; c++ {
			vals[r*in+c] = scale * (float32(c)/float32(in-1)*2 - 1)
		}
	}
	return vals
}

func TestQuantizeAffinePerChannel(t *testing.T) {
	const out, in = 5, 16
	vals := projectionWeights(out, in)
	q, packed, err := QuantizeAffinePerChannel(vals, []int64{out, in}, 0, DataTypeInt8)
	if err != nil {
		t.Fatalf("QuantizeAffinePerChannel: %v", err)
	}
	if len(packed) != out*in {
		t.Fatalf("packed %d bytes, want %d", len(packed), out*in)
	}
	if len(q.Scale) != out || len(q.ZeroPoint) != out {
		t.Fatalf("got %d scales and %d zero points, want %d of each", len(q.Scale), len(q.ZeroPoint), out)
	}
	if q.Axis != 0 {
		t.Errorf("Axis = %d, want 0", q.Axis)
	}
	// Each row must get its own scale, or this is not per-output-channel.
	for r := 1; r < out; r++ {
		if q.Scale[r] == q.Scale[r-1] {
			t.Errorf("rows %d and %d share scale %v", r-1, r, q.Scale[r])
		}
	}

	got, err := q.Dequantize(packed)
	if err != nil {
		t.Fatalf("Dequantize: %v", err)
	}
	// Error must be small relative to each row's own range, which is the
	// property per-channel scales buy.
	for r := 0; r < out; r++ {
		row := vals[r*in : (r+1)*in]
		lo, hi := row[0], row[0]
		for _, v := range row {
			lo, hi = float32(math.Min(float64(lo), float64(v))), float32(math.Max(float64(hi), float64(v)))
		}
		span := float64(hi - lo)
		for c := 0; c < in; c++ {
			i := r*in + c
			if e := math.Abs(float64(got[i] - vals[i])); e > span/255 {
				t.Errorf("row %d col %d: |%v - %v| = %v, more than one step of %v", r, c, got[i], vals[i], e, span/255)
			}
		}
	}
}

func TestQuantizeAffinePerChannelUInt8(t *testing.T) {
	vals := projectionWeights(3, 8)
	q, packed, err := QuantizeAffinePerChannel(vals, []int64{3, 8}, 0, DataTypeUInt8)
	if err != nil {
		t.Fatalf("QuantizeAffinePerChannel: %v", err)
	}
	for i, z := range q.ZeroPoint {
		if z < 0 || z > 255 {
			t.Errorf("zero point %d = %d, outside uint8's range", i, z)
		}
	}
	if _, err := q.Dequantize(packed); err != nil {
		t.Errorf("Dequantize: %v", err)
	}
}

// TestQuantizeConstantChannel covers a row that is entirely one value, where
// the natural scale is zero and an unguarded implementation divides by it.
func TestQuantizeConstantChannel(t *testing.T) {
	vals := []float32{3, 3, 3, 3, -1, 0, 1, 2}
	q, packed, err := QuantizeAffinePerChannel(vals, []int64{2, 4}, 0, DataTypeInt8)
	if err != nil {
		t.Fatalf("QuantizeAffinePerChannel: %v", err)
	}
	got, err := q.Dequantize(packed)
	if err != nil {
		t.Fatalf("Dequantize: %v", err)
	}
	for i := range got {
		if math.IsNaN(float64(got[i])) || math.IsInf(float64(got[i]), 0) {
			t.Fatalf("element %d dequantized to %v", i, got[i])
		}
	}
}

func TestQuantizeAffinePerChannelRejects(t *testing.T) {
	for _, tt := range []struct {
		name    string
		vals    []float32
		shape   []int64
		axis    int32
		dtype   DataType
		wantErr string
	}{
		{"no shape", []float32{1}, nil, 0, DataTypeInt8, "no shape"},
		{"axis out of range", []float32{1, 2}, []int64{2}, 1, DataTypeInt8, "outside the rank-1"},
		{"wrong value count", []float32{1, 2, 3}, []int64{2, 2}, 0, DataTypeInt8, "want 4"},
		{"not 8 bit", []float32{1, 2, 3, 4}, []int64{2, 2}, 0, DataTypeInt32, "not int8 or uint8"},
		{"NaN value", []float32{1, 2, 3, float32(math.NaN())}, []int64{2, 2}, 0, DataTypeInt8, "value 3 is NaN"},
		{"Inf value", []float32{1, 2, 3, float32(math.Inf(1))}, []int64{2, 2}, 0, DataTypeInt8, "value 3 is +Inf"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, _, err := QuantizeAffinePerChannel(tt.vals, tt.shape, tt.axis, tt.dtype)
			if err == nil {
				t.Fatalf("QuantizeAffinePerChannel = nil error, want %q", tt.wantErr)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("QuantizeAffinePerChannel = %v, want %q", err, tt.wantErr)
			}
		})
	}
}

func TestAffineDequantizationValidate(t *testing.T) {
	ok := AffineDequantization{
		Shape: []int64{4, 8}, QuantizedType: DataTypeInt8, OutputType: DataTypeFloat16,
		Axis: 0, Scale: []float32{1, 1, 1, 1}, ZeroPoint: []int64{0, 0, 0, 0},
	}
	if err := ok.Validate(); err != nil {
		t.Fatalf("Validate: %v", err)
	}
	// A scalar scale is legal for any shape and ignores Axis.
	scalar := ok
	scalar.Scale, scalar.ZeroPoint, scalar.Axis = []float32{0.5}, []int64{0}, 7
	if err := scalar.Validate(); err != nil {
		t.Errorf("scalar scale: %v", err)
	}

	for _, tt := range []struct {
		name    string
		mutate  func(*AffineDequantization)
		wantErr string
	}{
		{"no shape", func(q *AffineDequantization) { q.Shape = nil }, "no shape"},
		{"bad quantized type", func(q *AffineDequantization) { q.QuantizedType = DataTypeInt32 }, "not int8 or uint8"},
		{"bad output type", func(q *AffineDequantization) { q.OutputType = DataTypeInt8 }, "not fp16 or fp32"},
		{"no scale", func(q *AffineDequantization) { q.Scale = nil }, "no scale"},
		{
			"length mismatch",
			func(q *AffineDequantization) { q.ZeroPoint = []int64{0, 0} },
			"4 scales but 2 zero points",
		},
		{
			// scale * (q - zp) with scale 0 recovers zeros, whatever was
			// stored. It is the one scale value that loses the tensor.
			"zero scale",
			func(q *AffineDequantization) { q.Scale = []float32{1, 0, 1, 1} },
			"does not describe a recoverable tensor",
		},
		{
			"zero point out of range",
			func(q *AffineDequantization) { q.ZeroPoint = []int64{0, 200, 0, 0} },
			"outside int8's range",
		},
		{
			// "size(scale-vector) == quantized_data.shape[axis]"
			"scale count does not match the axis",
			func(q *AffineDequantization) { q.Axis = 1 },
			"4 scales for axis 1 of length 8",
		},
		{"axis out of range", func(q *AffineDequantization) { q.Axis = 2 }, "outside the rank-2"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			q := ok
			q.Scale = append([]float32(nil), ok.Scale...)
			q.ZeroPoint = append([]int64(nil), ok.ZeroPoint...)
			tt.mutate(&q)
			err := q.Validate()
			if err == nil {
				t.Fatalf("Validate = nil error, want %q", tt.wantErr)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("Validate = %v, want %q", err, tt.wantErr)
			}
		})
	}
}

// TestAffineDequantizationOperationPlacement pins where the parameters go.
// They are attributes, not inputs, at every opset: coremltools picks on the
// op class's own declaring opset, and constexpr_affine_dequantize is an iOS16
// op even inside an iOS18 program. TestConstexprPlacementRuntime confirms it
// against the real runtime, which rejects the inputs spelling outright.
func TestAffineDequantizationOperationPlacement(t *testing.T) {
	q := AffineDequantization{
		Shape: []int64{2, 4}, QuantizedType: DataTypeInt8, OutputType: DataTypeFloat16,
		Axis: 0, Scale: []float32{1, 2}, ZeroPoint: []int64{0, -1},
	}
	op, err := q.Operation("w", BlobRef{Path: "@model_path/weights/weight.bin", Offset: 64})
	if err != nil {
		t.Fatalf("Operation: %v", err)
	}
	if len(op.Attributes) != 4 {
		t.Errorf("got %d attributes, want the 4 parameters", len(op.Attributes))
	}
	if len(op.Inputs) != 0 {
		t.Errorf("got %d inputs, want none: the runtime rejects the inputs spelling", len(op.Inputs))
	}
	for _, name := range []string{"quantized_data", "zero_point", "scale", "axis"} {
		if op.Attributes[name] == nil {
			t.Errorf("no %q attribute", name)
		}
	}
}

func TestAffineDequantizationOperation(t *testing.T) {
	vals := projectionWeights(4, 8)
	q, packed, err := QuantizeAffinePerChannel(vals, []int64{4, 8}, 0, DataTypeInt8)
	if err != nil {
		t.Fatal(err)
	}
	files, refs, err := BuildWeights([]WeightTensor{{
		Name: "w", DType: BlobDataTypeInt8, Data: packed,
	}}, BlobLayoutSingleFile)
	if err != nil {
		t.Fatal(err)
	}
	op, err := q.Operation("w", refs["w"])
	if err != nil {
		t.Fatalf("Operation: %v", err)
	}

	// quantized_data must reference the blob, not inline the payload.
	qd := op.Attributes["quantized_data"]
	if qd.BlobFile == nil {
		t.Fatal("quantized_data does not reference a blob file")
	}
	if qd.BlobFile.FileName != files[0].Path || qd.BlobFile.Offset != refs["w"].Offset {
		t.Errorf("quantized_data references %s+%d, want %s+%d",
			qd.BlobFile.FileName, qd.BlobFile.Offset, files[0].Path, refs["w"].Offset)
	}
	if qd.Type.TensorType.DataType != DataTypeInt8 {
		t.Errorf("quantized_data type = %v, want int8", qd.Type.TensorType.DataType)
	}
	// scale and zero_point are small enough to inline, and must be.
	for _, name := range []string{"scale", "zero_point"} {
		v := op.Attributes[name]
		if v.Immediate == nil || v.Immediate.Tensor == nil {
			t.Errorf("%s is not an immediate tensor", name)
			continue
		}
		if err := ValidateTensorValue(v.Type.TensorType.DataType, v.Immediate.Tensor); err != nil {
			t.Errorf("%s: %v", name, err)
		}
		if got := v.Type.TensorType.Dimensions[0].Constant; got != 4 {
			t.Errorf("%s has %d entries, want 4", name, got)
		}
	}
	if got := op.Attributes["axis"].Immediate.Tensor.Ints; len(got) != 1 || got[0] != 0 {
		t.Errorf("axis = %v, want [0]", got)
	}
	if op.Outputs[0].Type.TensorType.DataType != DataTypeFloat16 {
		t.Errorf("output type = %v, want fp16", op.Outputs[0].Type.TensorType.DataType)
	}
}

func TestAffineDequantizationOperationRejects(t *testing.T) {
	q := AffineDequantization{
		Shape: []int64{2, 4}, QuantizedType: DataTypeInt8, OutputType: DataTypeFloat16,
		Scale: []float32{1, 2}, ZeroPoint: []int64{0, 0},
	}
	if _, err := q.Operation("", BlobRef{Path: "@model_path/w.bin"}); err == nil {
		t.Error("Operation with no name = nil error, want error")
	}
	if _, err := q.Operation("w", BlobRef{}); err == nil {
		t.Error("Operation with no blob path = nil error, want error")
	}
	scalar := q
	scalar.Scale, scalar.ZeroPoint = []float32{1}, []int64{0}
	if _, err := scalar.Operation("w", BlobRef{Path: "@model_path/w.bin"}); err != nil {
		t.Errorf("scalar scale should be legal: %v", err)
	}
}

// TestCompileQuantizedProgram takes an int8 per-output-channel weight all the
// way through: quantize, pack, place, describe, validate, compile. The op the
// descriptor builds has to satisfy the same op-schema table the rest of the
// package is held to, which is why constexpr_affine_dequantize had to be in it.
func TestCompileQuantizedProgram(t *testing.T) {
	const out, in = 4, 8
	vals := projectionWeights(out, in)
	q, packed, err := QuantizeAffinePerChannel(vals, []int64{out, in}, 0, DataTypeInt8)
	if err != nil {
		t.Fatal(err)
	}
	files, refs, err := BuildWeights([]WeightTensor{{
		Name: "w", DType: BlobDataTypeInt8, Data: packed,
	}}, BlobLayoutSingleFile)
	if err != nil {
		t.Fatal(err)
	}
	dequant, err := q.Operation("w", refs["w"])
	if err != nil {
		t.Fatal(err)
	}

	prog := ios18Prog(
		[]NamedValueType{fp16In("x", 1, in)},
		[]*Operation{
			dequant,
			{
				Type:    "linear",
				Inputs:  map[string]*Argument{"x": ref("x"), "weight": ref("w")},
				Outputs: []NamedValueType{fp16In("y", 1, out)},
			},
		},
		[]string{"y"},
	)
	inDesc, err := TensorFeatureDescription("x", "fp16", []int64{1, in})
	if err != nil {
		t.Fatal(err)
	}
	outDesc, err := TensorFeatureDescription("y", "fp16", []int64{1, out})
	if err != nil {
		t.Fatal(err)
	}
	desc := ModelDescription{
		Inputs:  []FeatureDescription{inDesc},
		Outputs: []FeatureDescription{outDesc},
	}

	tmp := t.TempDir()
	weightRoot := filepath.Join(tmp, "weightroot")
	if err := WriteWeightRoot(weightRoot, files); err != nil {
		t.Fatalf("WriteWeightRoot: %v", err)
	}
	bundle := filepath.Join(tmp, "model.mlmodelc")
	if err := CompileProgram(prog, 9, desc, weightRoot, bundle); err != nil {
		t.Fatalf("CompileProgram: %v", err)
	}

	milText, err := os.ReadFile(filepath.Join(bundle, "model.mil"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(milText), "constexpr_affine_dequantize(") {
		t.Errorf("emitted model.mil has no constexpr_affine_dequantize:\n%s", milText)
	}
	// The int8 blob the op points at must have been copied into the bundle.
	if _, err := os.Stat(filepath.Join(bundle, "weights", "weight.bin")); err != nil {
		t.Errorf("weight blob is not in the bundle: %v", err)
	}
	// Half the bytes of the fp16 the model would otherwise carry.
	if len(packed) != out*in {
		t.Errorf("packed %d bytes, want %d", len(packed), out*in)
	}
}

func TestDequantizeScalarIgnoresAxis(t *testing.T) {
	for _, axis := range []int32{-1, 0, 2} {
		q := AffineDequantization{Shape: []int64{2}, QuantizedType: DataTypeInt8, OutputType: DataTypeFloat32, Axis: axis, Scale: []float32{2}, ZeroPoint: []int64{0}}
		got, err := q.Dequantize([]byte{1, 2})
		if err != nil {
			t.Fatal(err)
		}
		if len(got) != 2 || got[0] != 2 || got[1] != 4 {
			t.Fatalf("axis %d: got %v, want [2 4]", axis, got)
		}
	}
}

func TestQuantizedShapeOverflow(t *testing.T) {
	shape := []int64{1 << 32, 1 << 32}
	q := AffineDequantization{Shape: shape, QuantizedType: DataTypeInt8, OutputType: DataTypeFloat32, Scale: []float32{1}, ZeroPoint: []int64{0}}
	if err := q.Validate(); err == nil {
		t.Fatal("Validate accepted overflowing shape")
	}
	if _, _, err := QuantizeAffinePerChannel(nil, shape, 0, DataTypeInt8); err == nil {
		t.Fatal("QuantizeAffinePerChannel accepted overflowing shape")
	}
}

func TestQuantizedScaleMustFitOutput(t *testing.T) {
	for _, scale := range []float32{1e-10, 1e10} {
		q := AffineDequantization{Shape: []int64{1}, QuantizedType: DataTypeInt8, OutputType: DataTypeFloat16, Scale: []float32{scale}, ZeroPoint: []int64{0}}
		if err := q.Validate(); err == nil {
			t.Fatalf("accepted fp16 scale %g", scale)
		}
		q.OutputType = DataTypeFloat32
		if err := q.Validate(); err != nil {
			t.Fatalf("rejected fp32 scale %g: %v", scale, err)
		}
	}
}
