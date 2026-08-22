package coremlcompiler

import (
	"strings"
	"testing"
)

func TestSanitizeName(t *testing.T) {
	tests := []struct {
		name string
		in   []string
		want []string
	}{
		{"valid unchanged", []string{"conv1"}, []string{"conv1"}},
		{"invalid chars", []string{"conv1/BiasAdd:0"}, []string{"conv1_BiasAdd_0"}},
		{"leading digit", []string{"0x"}, []string{"_0x"}},
		{"reserved word", []string{"tensor"}, []string{"tensor_workaround"}},
		{"collision", []string{"a_b", "a/b"}, []string{"a_b", "a_b_0"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := newNameSanitizer()
			for i, in := range tt.in {
				if got := s.sanitize(in); got != tt.want[i] {
					t.Errorf("sanitize(%q) = %q, want %q", in, got, tt.want[i])
				}
			}
		})
	}
}

// badNameModel builds a model whose names are legal in the source frameworks
// but illegal as MIL identifiers.
func badNameModel() *Model {
	f32 := &ValueType{TensorType: &TensorType{DataType: DataTypeFloat32, Dimensions: []Dimension{{Constant: 1}}}}
	op := &Operation{
		Type: "identity",
		Inputs: map[string]*Argument{
			"x": {Bindings: []Binding{{Name: "input:0"}}},
		},
		Outputs: []NamedValueType{{Name: "conv1/BiasAdd:0", Type: f32}},
	}
	blk := &Block{
		Outputs:    []string{"conv1/BiasAdd:0"},
		Operations: []*Operation{op},
	}
	return &Model{
		Description: ModelDescription{
			Inputs:  []FeatureDescription{{Name: "input:0"}},
			Outputs: []FeatureDescription{{Name: "conv1/BiasAdd:0"}},
		},
		MLProgram: &Program{
			Functions: map[string]*Function{
				"main": {
					Inputs:               []NamedValueType{{Name: "input:0", Type: f32}},
					OpSet:                "CoreML7",
					BlockSpecializations: map[string]*Block{"CoreML7": blk},
				},
			},
		},
	}
}

func TestValidateProgramRejectsInvalidNames(t *testing.T) {
	m := badNameModel()
	if err := ValidateProgram(m.MLProgram); err == nil {
		t.Fatal("ValidateProgram accepted a program with invalid MIL identifiers")
	}
}

func TestSanitizeProgramRewritesNamesAndDescription(t *testing.T) {
	m := badNameModel()
	SanitizeProgram(m)

	if err := ValidateProgram(m.MLProgram); err != nil {
		t.Fatalf("ValidateProgram after SanitizeProgram: %v", err)
	}
	if got := m.Description.Inputs[0].Name; got != "input_0" {
		t.Errorf("description input = %q, want %q", got, "input_0")
	}
	if got := m.Description.Outputs[0].Name; got != "conv1_BiasAdd_0" {
		t.Errorf("description output = %q, want %q", got, "conv1_BiasAdd_0")
	}

	text := emitMILText(m.MLProgram)
	if strings.ContainsAny(text, ":/") {
		t.Errorf("emitted MIL still contains invalid identifier characters:\n%s", text)
	}
	// The reference must have followed the definition.
	if !strings.Contains(text, "x = input_0") {
		t.Errorf("op input reference not renamed:\n%s", text)
	}
}
