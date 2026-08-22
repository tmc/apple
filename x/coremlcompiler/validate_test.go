package coremlcompiler

import (
	"strings"
	"testing"
)

func tt(dt DataType, dims ...uint64) *ValueType {
	ds := make([]Dimension, len(dims))
	for i, d := range dims {
		ds[i] = Dimension{Constant: d}
	}
	return &ValueType{TensorType: &TensorType{DataType: dt, Rank: int64(len(dims)), Dimensions: ds}}
}

func ref(name string) *Argument {
	return &Argument{Bindings: []Binding{{Name: name}}}
}

// progOf wraps a single main function whose body is ops, taking inputs and
// returning outs.
func progOf(inputs []NamedValueType, ops []*Operation, outs []string) *Program {
	return &Program{
		Functions: map[string]*Function{
			"main": {
				Inputs: inputs,
				OpSet:  "CoreML7",
				BlockSpecializations: map[string]*Block{
					"CoreML7": {Operations: ops, Outputs: outs},
				},
			},
		},
	}
}

func TestValidateProgramInvariants(t *testing.T) {
	x := NamedValueType{Name: "x", Type: tt(DataTypeFloat32, 1, 4)}
	relu := func(out, in string) *Operation {
		return &Operation{
			Type:    "relu",
			Inputs:  map[string]*Argument{"x": ref(in)},
			Outputs: []NamedValueType{{Name: out, Type: tt(DataTypeFloat32, 1, 4)}},
		}
	}

	tests := []struct {
		name    string
		prog    *Program
		wantErr string // substring; empty means valid
	}{
		{
			name: "valid",
			prog: progOf([]NamedValueType{x}, []*Operation{relu("y", "x")}, []string{"y"}),
		},
		{
			name: "use before def",
			prog: progOf([]NamedValueType{x},
				[]*Operation{relu("y", "z"), relu("z", "x")}, []string{"y"}),
			wantErr: `references "z", which is not visible`,
		},
		{
			name: "undefined operand",
			prog: progOf([]NamedValueType{x},
				[]*Operation{relu("y", "nope")}, []string{"y"}),
			wantErr: `references "nope", which is not visible`,
		},
		{
			name: "duplicate var name",
			prog: progOf([]NamedValueType{x},
				[]*Operation{relu("y", "x"), relu("y", "y")}, []string{"y"}),
			wantErr: `duplicate variable name "y"`,
		},
		{
			name: "op output shadows function input",
			prog: progOf([]NamedValueType{x},
				[]*Operation{relu("x", "x")}, []string{"x"}),
			wantErr: `duplicate variable name "x"`,
		},
		{
			name: "block output not visible",
			prog: progOf([]NamedValueType{x},
				[]*Operation{relu("y", "x")}, []string{"ghost"}),
			wantErr: `block output "ghost" is not visible`,
		},
		{
			name: "op with no outputs",
			prog: progOf([]NamedValueType{x},
				[]*Operation{{Type: "relu", Inputs: map[string]*Argument{"x": ref("x")}}}, nil),
			wantErr: "op relu has no outputs",
		},
		{
			name: "const with two outputs",
			prog: progOf(nil, []*Operation{{
				Type: "const",
				Outputs: []NamedValueType{
					{Name: "a", Type: tt(DataTypeFloat32, 1)},
					{Name: "b", Type: tt(DataTypeFloat32, 1)},
				},
			}}, []string{"a"}),
			wantErr: "const op must have exactly 1 output",
		},
		{
			name: "rank 6 output",
			prog: progOf([]NamedValueType{x}, []*Operation{{
				Type:    "reshape",
				Inputs:  map[string]*Argument{"x": ref("x")},
				Outputs: []NamedValueType{{Name: "y", Type: tt(DataTypeFloat32, 1, 1, 1, 1, 1, 1)}},
			}}, []string{"y"}),
			wantErr: "Core ML supports rank <= 5",
		},
		{
			name: "rank 6 const feeding constexpr is allowed",
			prog: progOf(nil, []*Operation{
				{
					Type:    "const",
					Outputs: []NamedValueType{{Name: "w", Type: tt(DataTypeInt8, 1, 1, 1, 1, 1, 1)}},
				},
				{
					Type:    "constexpr_affine_dequantize",
					Inputs:  map[string]*Argument{"quantized_data": ref("w")},
					Outputs: []NamedValueType{{Name: "y", Type: tt(DataTypeFloat16, 1, 4)}},
				},
			}, []string{"y"}),
		},
		{
			name: "list element rank 5",
			prog: progOf([]NamedValueType{x}, []*Operation{{
				Type:   "list_write",
				Inputs: map[string]*Argument{"x": ref("x")},
				Outputs: []NamedValueType{{Name: "y", Type: &ValueType{
					ListType: &ListType{ElementType: tt(DataTypeFloat32, 1, 1, 1, 1, 1)},
				}}},
			}}, []string{"y"}),
			wantErr: "Core ML supports rank <= 4",
		},
		{
			name: "non coreml dialect",
			prog: progOf([]NamedValueType{x}, []*Operation{{
				Type:    "torch::foo",
				Inputs:  map[string]*Argument{"x": ref("x")},
				Outputs: []NamedValueType{{Name: "y", Type: tt(DataTypeFloat32, 1, 4)}},
			}}, []string{"y"}),
			wantErr: `dialect namespace "torch"`,
		},
		{
			name: "rank 0 function input",
			prog: progOf([]NamedValueType{{Name: "s", Type: tt(DataTypeFloat32)}},
				[]*Operation{relu("y", "s")}, []string{"y"}),
			wantErr: "is rank 0",
		},
		{
			name: "fp32 state input",
			prog: progOf([]NamedValueType{{Name: "s", Type: &ValueType{
				StateType: &StateType{WrappedType: tt(DataTypeFloat32, 1, 4)},
			}}}, []*Operation{relu("y", "s")}, []string{"y"}),
			wantErr: "states must be fp16",
		},
		{
			name: "flexible state input",
			prog: progOf([]NamedValueType{{Name: "s", Type: &ValueType{
				StateType: &StateType{WrappedType: &ValueType{TensorType: &TensorType{
					DataType:   DataTypeFloat16,
					Rank:       2,
					Dimensions: []Dimension{{Constant: 1}, {Unknown: true}},
				}}},
			}}}, []*Operation{relu("y", "s")}, []string{"y"}),
			wantErr: "flexible shape",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateProgram(tc.prog)
			if tc.wantErr == "" {
				if err != nil {
					t.Fatalf("ValidateProgram = %v, want nil", err)
				}
				return
			}
			if err == nil {
				t.Fatalf("ValidateProgram = nil, want error containing %q", tc.wantErr)
			}
			if !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("ValidateProgram = %v, want error containing %q", err, tc.wantErr)
			}
		})
	}
}

// TestValidateProgramOpsetAgreement covers the cross-function rule: every
// function in a program must declare the same opset.
func TestValidateProgramOpsetAgreement(t *testing.T) {
	fn := func(opset string) *Function {
		return &Function{
			Inputs: []NamedValueType{{Name: "x", Type: tt(DataTypeFloat32, 1, 4)}},
			OpSet:  opset,
			BlockSpecializations: map[string]*Block{
				opset: {Outputs: []string{"x"}},
			},
		}
	}
	prog := &Program{Functions: map[string]*Function{
		"main":  fn("CoreML7"),
		"other": fn("CoreML8"),
	}}
	err := ValidateProgram(prog)
	if err == nil || !strings.Contains(err.Error(), "same opset") {
		t.Fatalf("ValidateProgram = %v, want same-opset error", err)
	}

	prog.Functions["other"] = fn("CoreML7")
	if err := ValidateProgram(prog); err != nil {
		t.Fatalf("ValidateProgram = %v, want nil", err)
	}
}

// TestValidateModelProgram covers the description-to-program rules.
func TestValidateModelProgram(t *testing.T) {
	base := func() *Model {
		return &Model{
			SpecVersion: 9,
			MLProgram:   progOf([]NamedValueType{{Name: "x", Type: tt(DataTypeFloat32, 1, 4)}}, nil, []string{"x"}),
		}
	}

	tests := []struct {
		name    string
		mutate  func(*Model)
		wantErr string
	}{
		{name: "valid", mutate: func(m *Model) { m.Description.DefaultFunctionName = "main" }},
		{
			name:    "default function missing",
			mutate:  func(m *Model) { m.Description.DefaultFunctionName = "nope" },
			wantErr: `default function "nope" not found`,
		},
		{
			name: "multifunction below ios18",
			mutate: func(m *Model) {
				m.SpecVersion = 8
				m.Description.Functions = []FunctionDescription{{Name: "main"}, {Name: "other"}}
			},
			wantErr: "requires specification version >= 9",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			m := base()
			tc.mutate(m)
			err := validateModelProgram(m)
			if tc.wantErr == "" {
				if err != nil {
					t.Fatalf("validateModelProgram = %v, want nil", err)
				}
				return
			}
			if err == nil || !strings.Contains(err.Error(), tc.wantErr) {
				t.Fatalf("validateModelProgram = %v, want error containing %q", err, tc.wantErr)
			}
		})
	}
}
