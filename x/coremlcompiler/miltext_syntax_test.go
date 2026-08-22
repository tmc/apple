package coremlcompiler

import (
	"math"
	"strings"
	"testing"
)

func strScalar(s string) *Value {
	return &Value{
		Type:      &ValueType{TensorType: &TensorType{DataType: DataTypeString}},
		Immediate: &ImmediateValue{Tensor: &TensorValue{Strings: []string{s}}},
	}
}

func strDictType() *ValueType {
	str := &ValueType{TensorType: &TensorType{DataType: DataTypeString}}
	return &ValueType{DictionaryType: &DictionaryType{KeyType: str, ValueType: str}}
}

// TestBuildInfoLiteral pins the dictionary literal spelling against the one
// coremlc-produced ground truth in the coremltools tree
// (coremltools/modelrunner/ModelRunner/add_model.mlmodelc/model.mil:2).
func TestBuildInfoLiteral(t *testing.T) {
	v := &Value{
		Type: strDictType(),
		Immediate: &ImmediateValue{Dictionary: &DictionaryValue{Entries: []DictionaryMapEntry{
			{Key: strScalar("coremlc-version"), Value: strScalar("3402.4.1")},
			{Key: strScalar("coremltools-version"), Value: strScalar("8.1")},
		}}},
	}
	want := `dict<tensor<string, []>, tensor<string, []>>({{"coremlc-version", "3402.4.1"}, {"coremltools-version", "8.1"}})`
	if got := formatValue(v); got != want {
		t.Errorf("formatValue:\n got %s\nwant %s", got, want)
	}
}

func TestMILStringEscaping(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"plain", "abc", `"abc"`},
		{"quote", `a"b`, `"a\"b"`},
		{"backslash", `a\b`, `"a\\b"`},
		{"newline", "a\nb", `"a\nb"`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatStringSlice([]string{tt.in}); got != tt.want {
				t.Errorf("formatStringSlice(%q) = %s, want %s", tt.in, got, tt.want)
			}
		})
	}

	bf := &Value{
		Type:     &ValueType{TensorType: &TensorType{DataType: DataTypeFloat16, Dimensions: []Dimension{{Constant: 2}}}},
		BlobFile: &BlobFileValue{FileName: `@model_path/weights/we"ird.bin`, Offset: 64},
	}
	if got := formatValue(bf); !strings.Contains(got, `we\"ird.bin`) {
		t.Errorf("blob path not escaped: %s", got)
	}
}

func TestValidateProgramAttributes(t *testing.T) {
	tests := []struct {
		name    string
		attrs   map[string]*Value
		wantErr bool
	}{
		{"buildInfo", map[string]*Value{"buildInfo": strScalar("x")}, false},
		{"other", map[string]*Value{"myKey": strScalar("x")}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateProgramAttributes(&Program{Attributes: tt.attrs})
			if (err != nil) != tt.wantErr {
				t.Errorf("validateProgramAttributes = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateNonFiniteFloats(t *testing.T) {
	tests := []struct {
		name    string
		vals    []float32
		wantErr bool
	}{
		{"finite", []float32{1, -2.5}, false},
		{"inf", []float32{float32(math.Inf(-1))}, true},
		{"nan", []float32{float32(math.NaN())}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			op := &Operation{
				Type:    "const",
				Outputs: []NamedValueType{{Name: "c", Type: &ValueType{TensorType: &TensorType{DataType: DataTypeFloat32, Dimensions: []Dimension{{Constant: uint64(len(tt.vals))}}}}}},
				Attributes: map[string]*Value{"val": {
					Type:      &ValueType{TensorType: &TensorType{DataType: DataTypeFloat32, Dimensions: []Dimension{{Constant: uint64(len(tt.vals))}}}},
					Immediate: &ImmediateValue{Tensor: &TensorValue{Floats: tt.vals}},
				}},
			}
			err := validateOpStructure(op)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateOpStructure = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestWriteStateNoOutputs guards the one op that is legally output-less
// (backend/mil/load.py:299-320 builds write_state with no outputs= kwarg).
func TestWriteStateNoOutputs(t *testing.T) {
	op := &Operation{Type: "write_state"}
	if err := validateOpStructure(op); err != nil {
		t.Errorf("write_state rejected: %v", err)
	}
	if err := validateOpStructure(&Operation{Type: "mul"}); err == nil {
		t.Error("output-less mul accepted, want error")
	}
}
