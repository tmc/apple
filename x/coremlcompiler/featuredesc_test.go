package coremlcompiler

import (
	"strings"
	"testing"
)

func TestParseDataTypeRoundTrip(t *testing.T) {
	// Every type DataType.String names must parse back to itself, so the two
	// directions cannot drift as MIL gains types.
	for dt := DataType(0); dt <= 64; dt++ {
		name := dt.String()
		if strings.HasPrefix(name, "unknown(") {
			continue
		}
		got, err := ParseDataType(name)
		if err != nil {
			t.Errorf("ParseDataType(%q) = %v, want %v", name, err, dt)
			continue
		}
		if got != dt {
			t.Errorf("ParseDataType(%q) = %v, want %v", name, got, dt)
		}
	}
}

func TestParseDataTypeUnknown(t *testing.T) {
	// float32 is the spelling every other toolchain uses; MIL says fp32. A
	// caller that guesses must get an error, not a type.
	for _, name := range []string{"", "float32", "fp32 ", "FP32", "half", "unknown(11)"} {
		if dt, err := ParseDataType(name); err == nil {
			t.Errorf("ParseDataType(%q) = %v, want error", name, dt)
		}
	}
}

func TestArrayDataType(t *testing.T) {
	for _, tt := range []struct {
		dt   DataType
		want ArrayDataType
	}{
		{DataTypeFloat16, ArrayDataTypeFloat16},
		{DataTypeFloat32, ArrayDataTypeFloat32},
		{DataTypeFloat64, ArrayDataTypeDouble},
		{DataTypeInt32, ArrayDataTypeInt32},
		{DataTypeInt8, ArrayDataTypeInt8},
	} {
		got, err := tt.dt.ArrayDataType()
		if err != nil || got != tt.want {
			t.Errorf("%v.ArrayDataType() = %v, %v; want %v, nil", tt.dt, got, err, tt.want)
		}
	}
}

// TestArrayDataTypeUnrepresentable is the check the adapter this replaces did
// not have: it mapped every type it did not know to Float32 and compiled a
// model whose inputs were the wrong width. Each type below reached that
// fallback.
func TestArrayDataTypeUnrepresentable(t *testing.T) {
	for _, dt := range []DataType{
		DataTypeBFloat16, DataTypeInt16, DataTypeInt64, DataTypeInt4,
		DataTypeUInt8, DataTypeUInt16, DataTypeUInt32, DataTypeUInt64,
		DataTypeUInt1, DataTypeUInt2, DataTypeUInt3, DataTypeUInt4, DataTypeUInt6,
		DataTypeFloat8E4M3FN, DataTypeFloat8E5M2, DataTypeBool, DataTypeString,
	} {
		got, err := dt.ArrayDataType()
		if err == nil {
			t.Errorf("%v.ArrayDataType() = %v, want error", dt, got)
			continue
		}
		if got == ArrayDataTypeFloat32 {
			t.Errorf("%v.ArrayDataType() returned Float32 alongside its error", dt)
		}
	}
}

// TestModelIRDTypes covers every element type github.com/tmc/modelir defines
// (modelir/program.go), the vocabulary this seam exists to accept. Four of the
// eight have no Core ML multi-array spelling and must be refused: "object" is
// not a MIL type at all, and bool, uint8 and uint32 are MIL types a multi-array
// cannot hold. None may reach a default case.
func TestModelIRDTypes(t *testing.T) {
	for _, tt := range []struct {
		dtype string
		want  ArrayDataType // ArrayDataTypeInvalid means "must error"
	}{
		{"fp16", ArrayDataTypeFloat16},
		{"fp32", ArrayDataTypeFloat32},
		{"int32", ArrayDataTypeInt32},
		{"int8", ArrayDataTypeInt8},
		{"object", ArrayDataTypeInvalid},
		{"bool", ArrayDataTypeInvalid},
		{"uint8", ArrayDataTypeInvalid},
		{"uint32", ArrayDataTypeInvalid},
	} {
		t.Run(tt.dtype, func(t *testing.T) {
			fd, err := TensorFeatureDescription("x", tt.dtype, []int64{1, 4})
			if tt.want == ArrayDataTypeInvalid {
				if err == nil {
					t.Fatalf("TensorFeatureDescription(x, %q) = %v, want error", tt.dtype, fd.Type.MultiArrayType.DataType)
				}
				return
			}
			if err != nil {
				t.Fatalf("TensorFeatureDescription(x, %q): %v", tt.dtype, err)
			}
			if got := fd.Type.MultiArrayType.DataType; got != tt.want {
				t.Fatalf("TensorFeatureDescription(x, %q) = %v, want %v", tt.dtype, got, tt.want)
			}
		})
	}
}

func TestTensorFeatureDescription(t *testing.T) {
	fd, err := TensorFeatureDescription("x", "fp16", []int64{1, 1, 8})
	if err != nil {
		t.Fatalf("TensorFeatureDescription: %v", err)
	}
	if fd.Name != "x" {
		t.Errorf("Name = %q, want %q", fd.Name, "x")
	}
	arr := fd.Type.MultiArrayType
	if arr == nil {
		t.Fatal("MultiArrayType is nil")
	}
	if arr.DataType != ArrayDataTypeFloat16 {
		t.Errorf("DataType = %v, want Float16", arr.DataType)
	}
	if len(arr.Shape) != 3 || arr.Shape[2] != 8 {
		t.Errorf("Shape = %v, want [1 1 8]", arr.Shape)
	}
	if fd.Type.StateArrayType != nil {
		t.Error("StateArrayType is set on a tensor feature")
	}
}

// TestTensorFeatureDescriptionCopiesShape guards against a caller reusing and
// then mutating the slice it passed in.
func TestTensorFeatureDescriptionCopiesShape(t *testing.T) {
	shape := []int64{1, 4}
	fd, err := TensorFeatureDescription("x", "fp32", shape)
	if err != nil {
		t.Fatal(err)
	}
	shape[1] = 99
	if got := fd.Type.MultiArrayType.Shape[1]; got != 4 {
		t.Errorf("shape aliased the caller's slice: got %d, want 4", got)
	}
}

func TestStateFeatureDescription(t *testing.T) {
	fd, err := StateFeatureDescription("kv", "fp16", []int64{1, 2, 4, 4})
	if err != nil {
		t.Fatalf("StateFeatureDescription: %v", err)
	}
	if fd.Type.StateArrayType == nil {
		t.Fatal("StateArrayType is nil")
	}
	if fd.Type.MultiArrayType != nil {
		t.Error("MultiArrayType is set on a state feature")
	}
	// The resulting description must satisfy the package's own wire rules.
	if err := validateFeatureDescription(fd, true); err != nil {
		t.Errorf("validateFeatureDescription: %v", err)
	}
}

func TestStateFeatureDescriptionRejectsType(t *testing.T) {
	// Core ML states are fp16 or int8 only; fp32 is the type a caller is most
	// likely to reach for, and it fails at load time rather than obviously.
	for _, dtype := range []string{"fp32", "int32", "fp64"} {
		if _, err := StateFeatureDescription("kv", dtype, []int64{1, 4}); err == nil {
			t.Errorf("StateFeatureDescription(kv, %q) = nil error, want error", dtype)
		}
	}
}

func TestFeatureDescriptionRejects(t *testing.T) {
	for _, tt := range []struct {
		name    string
		fname   string
		dtype   string
		shape   []int64
		wantErr string
	}{
		{"no name", "", "fp16", []int64{1}, "no name"},
		{"unknown dtype", "x", "float16", []int64{1}, "not a MIL data type"},
		{"unrepresentable dtype", "x", "uint16", []int64{1}, "no Core ML multi-array"},
		{"no shape", "x", "fp16", nil, "no shape"},
		{"zero dimension", "x", "fp16", []int64{1, 0}, "shapes must be static"},
		{"negative dimension", "x", "fp16", []int64{1, -1}, "shapes must be static"},
	} {
		t.Run(tt.name, func(t *testing.T) {
			_, err := TensorFeatureDescription(tt.fname, tt.dtype, tt.shape)
			if err == nil {
				t.Fatalf("TensorFeatureDescription = nil error, want %q", tt.wantErr)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("TensorFeatureDescription = %v, want %q", err, tt.wantErr)
			}
		})
	}
}
