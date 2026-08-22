package coremlcompiler

import "testing"

// TestFormatImmediateTensorNested verifies that multi-dimensional immediate
// tensor literals are emitted with bracket structure matching the declared
// shape. CoreML's MIL text parser reads a flat list of 12 elements as shape
// [12] and rejects it against a declared [3, 4] ("Size of dimensions of
// declared type and provided tensor literal are not compatible").
func TestFormatImmediateTensorNested(t *testing.T) {
	f32 := func(dims ...uint64) *ValueType {
		ds := make([]Dimension, len(dims))
		for i, d := range dims {
			ds[i] = Dimension{Constant: d}
		}
		return &ValueType{TensorType: &TensorType{
			DataType:   DataTypeFloat32,
			Rank:       int64(len(dims)),
			Dimensions: ds,
		}}
	}

	tests := []struct {
		name string
		v    *Value
		want string
	}{
		{
			name: "rank2 floats nest",
			v: &Value{
				Type:      f32(2, 3),
				Immediate: &ImmediateValue{Tensor: &TensorValue{Floats: []float32{1, 2, 3, 4, 5, 6}}},
			},
			want: "tensor<fp32, [2, 3]>([[1, 2, 3], [4, 5, 6]])",
		},
		{
			name: "rank3 ints nest",
			v: &Value{
				Type: &ValueType{TensorType: &TensorType{
					DataType: DataTypeInt32, Rank: 3,
					Dimensions: []Dimension{{Constant: 2}, {Constant: 2}, {Constant: 2}},
				}},
				Immediate: &ImmediateValue{Tensor: &TensorValue{Ints: []int32{1, 2, 3, 4, 5, 6, 7, 8}}},
			},
			want: "tensor<int32, [2, 2, 2]>([[[1, 2], [3, 4]], [[5, 6], [7, 8]]])",
		},
		{
			name: "rank1 stays flat",
			v: &Value{
				Type:      f32(3),
				Immediate: &ImmediateValue{Tensor: &TensorValue{Floats: []float32{1, 2, 3}}},
			},
			want: "tensor<fp32, [3]>([1, 2, 3])",
		},
		{
			name: "count mismatch stays flat",
			v: &Value{
				Type:      f32(3, 4),
				Immediate: &ImmediateValue{Tensor: &TensorValue{Floats: []float32{1, 2}}},
			},
			want: "tensor<fp32, [3, 4]>([1, 2])",
		},
		{
			name: "unknown dim stays flat",
			v: &Value{
				Type: &ValueType{TensorType: &TensorType{
					DataType: DataTypeFloat32, Rank: 2,
					Dimensions: []Dimension{{Unknown: true}, {Constant: 2}},
				}},
				Immediate: &ImmediateValue{Tensor: &TensorValue{Floats: []float32{1, 2, 3, 4}}},
			},
			want: "tensor<fp32, [?, 2]>([1, 2, 3, 4])",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatValue(tt.v); got != tt.want {
				t.Errorf("formatValue = %s, want %s", got, tt.want)
			}
		})
	}
}
