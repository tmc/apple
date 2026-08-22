package coremlcompiler

import (
	"path/filepath"
	"strings"
	"testing"
)

// arrayFeature builds a multi-array feature description.
func arrayFeature(name string, dt ArrayDataType, shape ...int64) FeatureDescription {
	return FeatureDescription{
		Name: name,
		Type: &FeatureType{MultiArrayType: &ArrayFeatureType{Shape: shape, DataType: dt}},
	}
}

// reluModel wraps a single relu from in to out over prog inputs.
func reluModel(in NamedValueType, out NamedValueType, desc ModelDescription) *Model {
	op := &Operation{
		Type:    "relu",
		Inputs:  map[string]*Argument{"x": ref(in.Name)},
		Outputs: []NamedValueType{out},
	}
	return &Model{
		SpecVersion: 8,
		Description: desc,
		MLProgram:   progOf([]NamedValueType{in}, []*Operation{op}, []string{out.Name}),
	}
}

func TestValidateDescriptionSignature(t *testing.T) {
	x := NamedValueType{Name: "x", Type: tt(DataTypeFloat32, 1, 4)}
	y := NamedValueType{Name: "y", Type: tt(DataTypeFloat32, 1, 4)}

	descOK := ModelDescription{
		Inputs:  []FeatureDescription{arrayFeature("x", ArrayDataTypeFloat32, 1, 4)},
		Outputs: []FeatureDescription{arrayFeature("y", ArrayDataTypeFloat32, 1, 4)},
	}

	// A state model: one fp16 input plus an fp16 state, output read from state.
	stateModel := func(desc ModelDescription) *Model {
		state := NamedValueType{Name: "k", Type: &ValueType{StateType: &StateType{WrappedType: tt(DataTypeFloat16, 1, 4)}}}
		in := NamedValueType{Name: "x", Type: tt(DataTypeFloat16, 1, 4)}
		op := &Operation{
			Type:    "read_state",
			Inputs:  map[string]*Argument{"input": ref("k")},
			Outputs: []NamedValueType{{Name: "cached", Type: tt(DataTypeFloat16, 1, 4)}},
		}
		prog := progOf([]NamedValueType{in, state}, []*Operation{op}, []string{"cached"})
		return &Model{SpecVersion: 8, Description: desc, MLProgram: prog}
	}
	stateDescOK := ModelDescription{
		Inputs:  []FeatureDescription{arrayFeature("x", ArrayDataTypeFloat16, 1, 4)},
		Outputs: []FeatureDescription{arrayFeature("cached", ArrayDataTypeFloat16, 1, 4)},
		States:  []FeatureDescription{arrayFeature("k", ArrayDataTypeFloat16, 1, 4)},
	}

	tests := []struct {
		name    string
		model   *Model
		wantErr string // substring; empty means valid
	}{
		{
			name:  "match",
			model: reluModel(x, y, descOK),
		},
		{
			name: "input name mismatch",
			model: reluModel(x, y, ModelDescription{
				Inputs:  []FeatureDescription{arrayFeature("z", ArrayDataTypeFloat32, 1, 4)},
				Outputs: descOK.Outputs,
			}),
			wantErr: `description declares input "z", which the program's signature does not have`,
		},
		{
			name: "program input undeclared",
			model: reluModel(x, y, ModelDescription{
				Outputs: descOK.Outputs,
			}),
			wantErr: `program input "x" has no description feature`,
		},
		{
			name: "output name not in program",
			model: reluModel(x, y, ModelDescription{
				Inputs:  descOK.Inputs,
				Outputs: []FeatureDescription{arrayFeature("nope", ArrayDataTypeFloat32, 1, 4)},
			}),
			wantErr: `description declares output "nope", which the program's signature does not have`,
		},
		{
			name: "data type mismatch",
			model: reluModel(
				NamedValueType{Name: "x", Type: tt(DataTypeFloat16, 1, 4)},
				NamedValueType{Name: "y", Type: tt(DataTypeFloat16, 1, 4)},
				descOK),
			wantErr: `input "x" has description data type 65568, program has fp16`,
		},
		{
			name: "unmappable data type",
			model: reluModel(
				NamedValueType{Name: "x", Type: tt(DataTypeInt64, 1, 4)},
				NamedValueType{Name: "y", Type: tt(DataTypeInt64, 1, 4)},
				descOK),
			wantErr: `data type int64 has no core ml feature data type`,
		},
		{
			name: "static shape mismatch",
			model: reluModel(x, y, ModelDescription{
				Inputs:  []FeatureDescription{arrayFeature("x", ArrayDataTypeFloat32, 1, 7)},
				Outputs: descOK.Outputs,
			}),
			wantErr: `input "x" has description shape [1 7], program has [1 4]`,
		},
		{
			name: "rank mismatch",
			model: reluModel(x, y, ModelDescription{
				Inputs:  []FeatureDescription{arrayFeature("x", ArrayDataTypeFloat32, 1, 4, 4)},
				Outputs: descOK.Outputs,
			}),
			wantErr: `input "x" has description shape [1 4 4], program has [1 4]`,
		},
		{
			name:  "state match",
			model: stateModel(stateDescOK),
		},
		{
			name: "state declared as input",
			model: stateModel(ModelDescription{
				Inputs: []FeatureDescription{
					arrayFeature("x", ArrayDataTypeFloat16, 1, 4),
					arrayFeature("k", ArrayDataTypeFloat16, 1, 4),
				},
				Outputs: stateDescOK.Outputs,
			}),
			wantErr: `description declares input "k", which the program's signature does not have`,
		},
		{
			name: "state undeclared",
			model: stateModel(ModelDescription{
				Inputs:  stateDescOK.Inputs,
				Outputs: stateDescOK.Outputs,
			}),
			wantErr: `program state "k" has no description feature`,
		},
		{
			name: "symbolic dim carries no shape obligation",
			model: reluModel(
				NamedValueType{Name: "x", Type: &ValueType{TensorType: &TensorType{
					DataType:   DataTypeFloat32,
					Rank:       2,
					Dimensions: []Dimension{{Unknown: true}, {Constant: 4}},
				}}},
				y,
				ModelDescription{
					Inputs:  []FeatureDescription{arrayFeature("x", ArrayDataTypeFloat32, 1, 4)},
					Outputs: descOK.Outputs,
				}),
		},
		{
			name: "empty declared shape",
			model: reluModel(x, y, ModelDescription{
				Inputs:  []FeatureDescription{arrayFeature("x", ArrayDataTypeFloat32)},
				Outputs: []FeatureDescription{arrayFeature("y", ArrayDataTypeFloat32)},
			}),
		},
		{
			name: "image input against rank-4 tensor",
			model: reluModel(
				NamedValueType{Name: "x", Type: tt(DataTypeFloat32, 1, 3, 8, 8)},
				NamedValueType{Name: "y", Type: tt(DataTypeFloat32, 1, 3, 8, 8)},
				ModelDescription{
					Inputs: []FeatureDescription{{
						Name: "x",
						Type: &FeatureType{ImageType: &ImageFeatureType{Width: 8, Height: 8, ColorSpace: ColorSpaceRGB}},
					}},
					Outputs: []FeatureDescription{arrayFeature("y", ArrayDataTypeFloat32, 1, 3, 8, 8)},
				}),
		},
		{
			name: "image input against rank-2 tensor",
			model: reluModel(x, y, ModelDescription{
				Inputs: []FeatureDescription{{
					Name: "x",
					Type: &FeatureType{ImageType: &ImageFeatureType{Width: 8, Height: 8, ColorSpace: ColorSpaceRGB}},
				}},
				Outputs: descOK.Outputs,
			}),
			wantErr: `image input "x" must have rank 4, program has rank 2`,
		},
		{
			name: "unknown default function is not our rule",
			model: &Model{
				SpecVersion: 8,
				Description: ModelDescription{Inputs: descOK.Inputs},
				MLProgram:   &Program{Functions: map[string]*Function{"other": {}}},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateDescriptionSignature(tt.model)
			if tt.wantErr == "" {
				if err != nil {
					t.Fatalf("validateDescriptionSignature = %v, want nil", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("validateDescriptionSignature = nil, want error containing %q", tt.wantErr)
			}
			if !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("validateDescriptionSignature = %v, want error containing %q", err, tt.wantErr)
			}
		})
	}
}

func TestValidateDescriptionSignatureMultiFunction(t *testing.T) {
	fnOf := func(in, out string) *Function {
		return &Function{
			Inputs: []NamedValueType{{Name: in, Type: tt(DataTypeFloat32, 1, 4)}},
			OpSet:  "CoreML7",
			BlockSpecializations: map[string]*Block{
				"CoreML7": {
					Operations: []*Operation{{
						Type:    "relu",
						Inputs:  map[string]*Argument{"x": ref(in)},
						Outputs: []NamedValueType{{Name: out, Type: tt(DataTypeFloat32, 1, 4)}},
					}},
					Outputs: []string{out},
				},
			},
		}
	}
	prog := &Program{Functions: map[string]*Function{
		"main":  fnOf("x", "y"),
		"other": fnOf("a", "b"),
	}}
	descOf := func(otherIn string) ModelDescription {
		return ModelDescription{
			DefaultFunctionName: "main",
			Functions: []FunctionDescription{
				{
					Name:    "main",
					Inputs:  []FeatureDescription{arrayFeature("x", ArrayDataTypeFloat32, 1, 4)},
					Outputs: []FeatureDescription{arrayFeature("y", ArrayDataTypeFloat32, 1, 4)},
				},
				{
					Name:    "other",
					Inputs:  []FeatureDescription{arrayFeature(otherIn, ArrayDataTypeFloat32, 1, 4)},
					Outputs: []FeatureDescription{arrayFeature("b", ArrayDataTypeFloat32, 1, 4)},
				},
			},
		}
	}

	if err := validateDescriptionSignature(&Model{SpecVersion: 9, Description: descOf("a"), MLProgram: prog}); err != nil {
		t.Fatalf("valid multi-function model: %v", err)
	}
	err := validateDescriptionSignature(&Model{SpecVersion: 9, Description: descOf("wrong"), MLProgram: prog})
	if err == nil || !strings.Contains(err.Error(), `function "other": description declares input "wrong"`) {
		t.Fatalf("multi-function mismatch = %v, want an error naming function \"other\"", err)
	}
}

// TestCompileRejectsSignatureMismatch covers the wiring: the check must run on
// the production compile path, not only through the helper.
func TestCompileRejectsSignatureMismatch(t *testing.T) {
	model := reluModel(
		NamedValueType{Name: "y", Type: tt(DataTypeFloat32, 1, 4)},
		NamedValueType{Name: "out", Type: tt(DataTypeFloat32, 1, 4)},
		ModelDescription{
			Inputs:  []FeatureDescription{arrayFeature("x", ArrayDataTypeInt32, 7, 7, 7)},
			Outputs: []FeatureDescription{arrayFeature("nope", ArrayDataTypeFloat32, 1, 4)},
		})
	err := compileMLProgram(model, "", filepath.Join(t.TempDir(), "model.mlmodelc"))
	if err == nil {
		t.Fatal("compileMLProgram = nil, want a signature mismatch error")
	}
	if !strings.Contains(err.Error(), `description declares input "x"`) {
		t.Fatalf("compileMLProgram = %v, want an error naming input \"x\"", err)
	}
}
