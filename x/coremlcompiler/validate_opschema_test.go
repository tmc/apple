package coremlcompiler

import (
	"strings"
	"testing"
)

func TestValidateOpSchema(t *testing.T) {
	op := func(typ string, inputs map[string]*Argument) *Operation {
		return &Operation{
			Type:    typ,
			Inputs:  inputs,
			Outputs: []NamedValueType{{Name: "y", Type: tt(DataTypeFloat32, 1, 4)}},
		}
	}

	tests := []struct {
		name    string
		op      *Operation
		opset   string
		wantErr string // substring; empty means valid
	}{
		{
			name:  "known op known inputs",
			op:    op("conv", map[string]*Argument{"x": ref("x"), "weight": ref("w"), "pad_type": ref("pt")}),
			opset: "ios18",
		},
		{
			name:    "misspelled input",
			op:      op("conv", map[string]*Argument{"x": ref("x"), "weights": ref("w")}),
			opset:   "ios18",
			wantErr: `op conv has no input "weights" in opset ios18`,
		},
		{
			name:    "input from another op",
			op:      op("softmax", map[string]*Argument{"x": ref("x"), "perm": ref("p")}),
			opset:   "ios17",
			wantErr: `op softmax has no input "perm"`,
		},
		{
			// scaled_dot_product_attention was added in iOS18
			// (ops/defs/iOS18/transformers.py).
			name:    "op too new for opset",
			op:      op("scaled_dot_product_attention", map[string]*Argument{"query": ref("q")}),
			opset:   "ios17",
			wantErr: "is not available in opset ios17",
		},
		{
			name:  "op available in opset",
			op:    op("scaled_dot_product_attention", map[string]*Argument{"query": ref("q")}),
			opset: "ios18",
		},
		{
			// No schema is dumped past iOS18, so checking is skipped
			// rather than reported.
			name:  "unknown opset skips",
			op:    op("conv", map[string]*Argument{"weights": ref("w")}),
			opset: "ios26",
		},
		{
			// linear is a real op outside the dump: absence from the table
			// must not be read as absence from coremltools.
			name:  "undumped op skips",
			op:    op("linear", map[string]*Argument{"nonsense": ref("w")}),
			opset: "ios18",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateOpSchema(tt.op, tt.opset)
			switch {
			case tt.wantErr == "" && err != nil:
				t.Fatalf("validateOpSchema = %v, want nil", err)
			case tt.wantErr != "" && err == nil:
				t.Fatalf("validateOpSchema = nil, want %q", tt.wantErr)
			case tt.wantErr != "" && !strings.Contains(err.Error(), tt.wantErr):
				t.Fatalf("validateOpSchema = %v, want %q", err, tt.wantErr)
			}
		})
	}
}

// TestValidateProgramOpSchema checks the schema rules reach ValidateProgram,
// with the opset taken from the function it is declared on.
func TestValidateProgramOpSchema(t *testing.T) {
	bad := &Operation{
		Type:    "relu",
		Inputs:  map[string]*Argument{"input": ref("x")},
		Outputs: []NamedValueType{{Name: "y", Type: tt(DataTypeFloat32, 1, 4)}},
	}
	prog := progOf([]NamedValueType{{Name: "x", Type: tt(DataTypeFloat32, 1, 4)}}, []*Operation{bad}, []string{"y"})
	err := ValidateProgram(prog)
	if err == nil || !strings.Contains(err.Error(), `op relu has no input "input" in opset ios17`) {
		t.Fatalf("ValidateProgram = %v, want relu input error", err)
	}
}
