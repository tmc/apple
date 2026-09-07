package coremlcompiler

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// ios18Prog is progOf's iOS18 counterpart: the state ops the graphs this path
// exists for are built from arrived with that opset.
func ios18Prog(inputs []NamedValueType, ops []*Operation, outs []string) *Program {
	return &Program{
		Functions: map[string]*Function{
			"main": {
				Inputs: inputs,
				OpSet:  "CoreML8",
				BlockSpecializations: map[string]*Block{
					"CoreML8": {Operations: ops, Outputs: outs},
				},
			},
		},
	}
}

func fp16In(name string, dims ...uint64) NamedValueType {
	return NamedValueType{Name: name, Type: tt(DataTypeFloat16, dims...)}
}

// reluProgram is the smallest program that compiles, plus the description
// matching it.
func reluProgram() (*Program, ModelDescription) {
	prog := ios18Prog(
		[]NamedValueType{fp16In("x", 1, 4)},
		[]*Operation{{
			Type:    "relu",
			Inputs:  map[string]*Argument{"x": ref("x")},
			Outputs: []NamedValueType{fp16In("y", 1, 4)},
		}},
		[]string{"y"},
	)
	in, err := TensorFeatureDescription("x", "fp16", []int64{1, 4})
	if err != nil {
		panic(err)
	}
	out, err := TensorFeatureDescription("y", "fp16", []int64{1, 4})
	if err != nil {
		panic(err)
	}
	return prog, ModelDescription{Inputs: []FeatureDescription{in}, Outputs: []FeatureDescription{out}}
}

func TestCompileProgram(t *testing.T) {
	prog, desc := reluProgram()
	out := filepath.Join(t.TempDir(), "model.mlmodelc")
	if err := CompileProgram(prog, 9, desc, "", out); err != nil {
		t.Fatalf("CompileProgram: %v", err)
	}
	for _, name := range []string{"model.mil", "coremldata.bin", "metadata.json"} {
		if _, err := os.Stat(filepath.Join(out, name)); err != nil {
			t.Errorf("bundle is missing %s: %v", name, err)
		}
	}
	// The MIL text is emitted here rather than supplied, so it cannot disagree
	// with the program that was validated.
	milText, err := os.ReadFile(filepath.Join(out, "model.mil"))
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"func main<ios18>", "relu("} {
		if !strings.Contains(string(milText), want) {
			t.Errorf("emitted model.mil does not contain %q:\n%s", want, milText)
		}
	}
}

func TestCompileProgramRejectsNil(t *testing.T) {
	out := filepath.Join(t.TempDir(), "m.mlmodelc")
	if err := CompileProgram(nil, 9, ModelDescription{}, "", out); err == nil {
		t.Error("CompileProgram(nil) = nil error, want error")
	}
	if err := CompileProgram(&Program{}, 9, ModelDescription{}, "", out); err == nil {
		t.Error("CompileProgram(empty) = nil error, want error")
	}
}

// TestCompileProgramValidatesWhatMILTextDoesNot is the point of the entry
// point. Each defect below is one CompileProgram rejects and CompileMILText
// compiles without complaint, because CompileMILText has no program to check.
func TestCompileProgramValidatesWhatMILTextDoesNot(t *testing.T) {
	tests := []struct {
		name    string
		mutate  func(*Program, *ModelDescription)
		wantErr string
	}{
		{
			// The op-schema table: an input name the op does not declare.
			// This is the check the table extension made reach real graphs.
			name: "undeclared op input",
			mutate: func(p *Program, _ *ModelDescription) {
				op := p.Functions["main"].BlockSpecializations["CoreML8"].Operations[0]
				op.Inputs = map[string]*Argument{"input": ref("x")}
			},
			wantErr: `op relu has no input "input"`,
		},
		{
			// write_state's inputs are named for the MIL backend's lowering,
			// not for coreml_update_state's own inputs.
			name: "write_state with its source op's input names",
			mutate: func(p *Program, _ *ModelDescription) {
				blk := p.Functions["main"].BlockSpecializations["CoreML8"]
				blk.Operations = append(blk.Operations, &Operation{
					Type:   "write_state",
					Inputs: map[string]*Argument{"state": ref("s"), "value": ref("y")},
				})
			},
			wantErr: `op write_state has no input "state"`,
		},
		{
			// The description must match the program's own signature.
			name: "description names an input the program does not have",
			mutate: func(_ *Program, d *ModelDescription) {
				fd, err := TensorFeatureDescription("not_an_input", "fp16", []int64{1, 4})
				if err != nil {
					t.Fatal(err)
				}
				d.Inputs = []FeatureDescription{fd}
			},
			wantErr: "",
		},
		{
			name: "output is not produced by the program",
			mutate: func(p *Program, _ *ModelDescription) {
				p.Functions["main"].BlockSpecializations["CoreML8"].Outputs = []string{"nope"}
			},
			wantErr: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			prog, desc := reluProgram()
			tt.mutate(prog, &desc)

			dir := t.TempDir()
			err := CompileProgram(prog, 9, desc, "", filepath.Join(dir, "prog.mlmodelc"))
			if err == nil {
				t.Fatalf("CompileProgram = nil error, want the defect rejected")
			}
			if tt.wantErr != "" && !strings.Contains(err.Error(), tt.wantErr) {
				t.Fatalf("CompileProgram = %v, want %q", err, tt.wantErr)
			}
			t.Logf("CompileProgram rejected it: %v", err)

			// The contrast: hand the same defect to CompileMILText as text and
			// it compiles. If this ever starts failing, CompileMILText grew a
			// validator and its doc comment needs updating.
			milText := emitMILTextWithSpec(prog, 9)
			textOut := filepath.Join(dir, "text.mlmodelc")
			if err := CompileMILText(milText, 9, desc, "", textOut); err != nil {
				t.Fatalf("CompileMILText rejected it too (%v); CompileMILText's doc comment says it validates nothing", err)
			}
			if _, err := os.Stat(filepath.Join(textOut, "model.mil")); err != nil {
				t.Fatalf("CompileMILText reported success without writing model.mil: %v", err)
			}
		})
	}
}

// TestCompileProgramEmitsStateOps compiles a program shaped like the state
// half of testdata/modelir_transformer.mil: read_state, then write_state named
// the way the MIL backend names its inputs. It is the closest this package can
// come to compiling that fixture through CompileProgram, since there is no MIL
// text parser to turn the fixture back into a Program.
func TestCompileProgramEmitsStateOps(t *testing.T) {
	stateType := &ValueType{StateType: &StateType{WrappedType: tt(DataTypeFloat16, 1, 2, 4, 4)}}
	prog := ios18Prog(
		[]NamedValueType{fp16In("x", 1, 2, 4, 4), {Name: "s", Type: stateType}},
		[]*Operation{
			{
				Type:    "read_state",
				Inputs:  map[string]*Argument{"input": ref("s")},
				Outputs: []NamedValueType{fp16In("r", 1, 2, 4, 4)},
			},
			{
				Type:   "write_state",
				Inputs: map[string]*Argument{"input": ref("s"), "data": ref("x")},
			},
			{
				Type:    "relu",
				Inputs:  map[string]*Argument{"x": ref("r")},
				Outputs: []NamedValueType{fp16In("y", 1, 2, 4, 4)},
			},
		},
		[]string{"y"},
	)
	in, err := TensorFeatureDescription("x", "fp16", []int64{1, 2, 4, 4})
	if err != nil {
		t.Fatal(err)
	}
	out, err := TensorFeatureDescription("y", "fp16", []int64{1, 2, 4, 4})
	if err != nil {
		t.Fatal(err)
	}
	state, err := StateFeatureDescription("s", "fp16", []int64{1, 2, 4, 4})
	if err != nil {
		t.Fatal(err)
	}
	desc := ModelDescription{
		Inputs:  []FeatureDescription{in},
		Outputs: []FeatureDescription{out},
		States:  []FeatureDescription{state},
	}

	dir := filepath.Join(t.TempDir(), "state.mlmodelc")
	if err := CompileProgram(prog, 9, desc, "", dir); err != nil {
		t.Fatalf("CompileProgram: %v", err)
	}
	milText, err := os.ReadFile(filepath.Join(dir, "model.mil"))
	if err != nil {
		t.Fatal(err)
	}
	// The fixture spells these the same way.
	for _, want := range []string{"read_state(", "write_state("} {
		if !strings.Contains(string(milText), want) {
			t.Errorf("emitted model.mil does not contain %q:\n%s", want, milText)
		}
	}
}
