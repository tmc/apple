package coremlcompiler

import (
	"testing"
)

// TestDataTypeWireNumbers pins every MILSpec.DataType enum number against
// mlmodel/format/MIL.proto. A wrong number here is silent element-type
// corruption: the runtime reads the varint and believes it.
func TestDataTypeWireNumbers(t *testing.T) {
	tests := []struct {
		name string
		dt   DataType
		want int32
	}{
		{"bool", DataTypeBool, 1},
		{"string", DataTypeString, 2},
		{"fp16", DataTypeFloat16, 10},
		{"fp32", DataTypeFloat32, 11},
		{"fp64", DataTypeFloat64, 12},
		{"bf16", DataTypeBFloat16, 13},
		{"int8", DataTypeInt8, 21},
		{"int16", DataTypeInt16, 22},
		{"int32", DataTypeInt32, 23},
		{"int64", DataTypeInt64, 24},
		{"int4", DataTypeInt4, 25},
		{"uint8", DataTypeUInt8, 31},
		{"uint16", DataTypeUInt16, 32},
		{"uint32", DataTypeUInt32, 33},
		{"uint64", DataTypeUInt64, 34},
		{"uint4", DataTypeUInt4, 35},
		{"uint2", DataTypeUInt2, 36},
		{"uint1", DataTypeUInt1, 37},
		{"uint6", DataTypeUInt6, 38},
		{"uint3", DataTypeUInt3, 39},
		{"fp8_e4m3fn", DataTypeFloat8E4M3FN, 40},
		{"fp8_e5m2", DataTypeFloat8E5M2, 41},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if int32(tt.dt) != tt.want {
				t.Errorf("%s = %d, want %d", tt.name, int32(tt.dt), tt.want)
			}
			if got := tt.dt.String(); got != tt.name {
				t.Errorf("String() = %q, want %q", got, tt.name)
			}
		})
	}
}

// TestArrayDataTypeWireNumbers pins FeatureTypes.proto ArrayDataType values.
func TestArrayDataTypeWireNumbers(t *testing.T) {
	tests := []struct {
		name string
		dt   ArrayDataType
		want int32
	}{
		{"invalid", ArrayDataTypeInvalid, 0},
		{"float16", ArrayDataTypeFloat16, 65552},
		{"float32", ArrayDataTypeFloat32, 65568},
		{"double", ArrayDataTypeDouble, 65600},
		{"int32", ArrayDataTypeInt32, 131104},
		{"int8", ArrayDataTypeInt8, 131080},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if int32(tt.dt) != tt.want {
				t.Errorf("%s = %d, want %d", tt.name, int32(tt.dt), tt.want)
			}
		})
	}
}

// TestFieldForDataType mirrors coremltools' _tensor_field_by_type.
func TestFieldForDataType(t *testing.T) {
	tests := []struct {
		name    string
		dt      DataType
		want    TensorValueField
		wantErr bool
	}{
		{"bool", DataTypeBool, TensorValueBools, false},
		{"string", DataTypeString, TensorValueStrings, false},
		{"int64", DataTypeInt64, TensorValueLongs, false},
		{"uint64", DataTypeUInt64, TensorValueLongs, false},
		{"fp64", DataTypeFloat64, TensorValueDoubles, false},
		{"fp32", DataTypeFloat32, TensorValueFloats, false},
		{"fp16", DataTypeFloat16, TensorValueBytes, false},
		{"int8", DataTypeInt8, TensorValueBytes, false},
		{"int4", DataTypeInt4, TensorValueBytes, false},
		{"uint8", DataTypeUInt8, TensorValueBytes, false},
		{"uint32", DataTypeUInt32, TensorValueBytes, false},
		{"uint4", DataTypeUInt4, TensorValueBytes, false},
		{"int16", DataTypeInt16, TensorValueInts, false},
		{"uint16", DataTypeUInt16, TensorValueInts, false},
		{"int32", DataTypeInt32, TensorValueInts, false},
		{"bf16 has no immediate encoding", DataTypeBFloat16, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.dt.FieldForDataType()
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("field = %d, want %d", got, tt.want)
			}
		})
	}
}

func TestValidateTensorValue(t *testing.T) {
	tests := []struct {
		name    string
		dt      DataType
		tv      TensorValue
		wantErr bool
	}{
		{"fp32 in floats", DataTypeFloat32, TensorValue{Floats: []float32{1}}, false},
		{"fp16 in bytes", DataTypeFloat16, TensorValue{Bytes: []byte{0, 0}}, false},
		{"fp16 in floats", DataTypeFloat16, TensorValue{Floats: []float32{1}}, true},
		{"uint8 in ints", DataTypeUInt8, TensorValue{Ints: []int32{1}}, true},
		{"int32 in ints", DataTypeInt32, TensorValue{Ints: []int32{1}}, false},
		{"uint32 in ints", DataTypeUInt32, TensorValue{Ints: []int32{1}}, true},
		{"nothing set", DataTypeFloat32, TensorValue{}, true},
		{"empty float tensor is set", DataTypeFloat32, TensorValue{Floats: []float32{}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateTensorValue(tt.dt, &tt.tv)
			if (err != nil) != tt.wantErr {
				t.Errorf("err = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
