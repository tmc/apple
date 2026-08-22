//go:build darwin

package coremlcompiler

import (
	"path/filepath"
	"testing"
)

// TestCompileRank2ImmediateLoadable builds a model whose linear weights are
// rank-2 immediate values (no weight.bin), compiles it with the pure-Go
// path, and verifies Apple's MLModel runtime loads it. Before nested
// literal emission this failed at load with "Size of dimensions of declared
// type and provided tensor literal are not compatible".
func TestCompileRank2ImmediateLoadable(t *testing.T) {
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
	strVal := func(s string) *Value {
		return &Value{
			Type:      &ValueType{TensorType: &TensorType{DataType: DataTypeString}},
			Immediate: &ImmediateValue{Tensor: &TensorValue{Strings: []string{s}}},
		}
	}
	constOp := func(name string, vt *ValueType, floats []float32) *Operation {
		return &Operation{
			Type:    "const",
			Outputs: []NamedValueType{{Name: name, Type: vt}},
			Attributes: map[string]*Value{
				"name": strVal("const_" + name),
				"val": {
					Type:      vt,
					Immediate: &ImmediateValue{Tensor: &TensorValue{Floats: floats}},
				},
			},
		}
	}
	ref := func(name string) *Argument {
		return &Argument{Bindings: []Binding{{Name: name}}}
	}

	model := &Model{
		SpecVersion: 9,
		Description: ModelDescription{
			Inputs: []FeatureDescription{{
				Name: "x",
				Type: &FeatureType{MultiArrayType: &ArrayFeatureType{
					Shape: []int64{1, 2}, DataType: ArrayDataTypeFloat32,
				}},
			}},
			Outputs: []FeatureDescription{{
				Name: "y",
				Type: &FeatureType{MultiArrayType: &ArrayFeatureType{
					Shape: []int64{1, 2}, DataType: ArrayDataTypeFloat32,
				}},
			}},
		},
		MLProgram: &Program{
			Version: 1,
			Functions: map[string]*Function{
				"main": {
					OpSet:  "CoreML7",
					Inputs: []NamedValueType{{Name: "x", Type: f32(1, 2)}},
					BlockSpecializations: map[string]*Block{
						"CoreML7": {
							Outputs: []string{"y"},
							Operations: []*Operation{
								constOp("w", f32(2, 2), []float32{1, 2, 3, 4}),
								constOp("b", f32(2), []float32{0.5, -0.5}),
								{
									Type: "linear",
									Inputs: map[string]*Argument{
										"x": ref("x"), "weight": ref("w"), "bias": ref("b"),
									},
									Outputs:    []NamedValueType{{Name: "y", Type: f32(1, 2)}},
									Attributes: map[string]*Value{"name": strVal("linear_0")},
								},
							},
						},
					},
				},
			},
		},
	}

	outputDir := filepath.Join(t.TempDir(), "model.mlmodelc")
	if err := CompileMLProgram(EncodeModel(model), "", outputDir); err != nil {
		t.Fatalf("CompileMLProgram: %v", err)
	}

	loaded, err := LoadCoreMLModel(outputDir)
	if err != nil {
		t.Fatalf("LoadCoreMLModel: %v", err)
	}
	loaded.Close()
}
