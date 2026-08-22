package coremlcompiler

// Program invariants that coremltools enforces before it will serialize a MIL
// program. Each check below mirrors a rule that raises in coremltools; a
// program that violates one is rejected by Apple's compiler with a message
// that does not name the offending var, op or function.

import (
	"fmt"
	"math"
	"strings"

	"github.com/tmc/apple/x/coremlcompiler/internal/opschema"
)

// maxTensorRank is the highest tensor rank the Core ML runtime accepts.
const maxTensorRank = 5

// maxListElementRank is the highest rank a list element type may have.
const maxListElementRank = 4

// validateFunction checks the per-function invariants: variable-name
// uniqueness, definition before use, block-output visibility, tensor rank and
// the per-op structural rules.
func validateFunction(fname string, fn *Function) error {
	// Variable names share one namespace per function, nested blocks
	// included: a duplicate makes a MIL reference ambiguous.
	// coremltools renames duplicates (dedup_op_and_var_names.py:59-81) and
	// raises outright when the collision involves a function input or output
	// (:83-94), which cannot be renamed. We have no renaming stage on the
	// emit path, so every collision is an error.
	seen := make(map[string]bool)
	declare := func(what, name string) error {
		if seen[name] {
			return fmt.Errorf("function %q: duplicate variable name %q", fname, name)
		}
		seen[name] = true
		return nil
	}
	for _, in := range fn.Inputs {
		if err := declare("input", in.Name); err != nil {
			return err
		}
		if err := validateFunctionInput(fname, in); err != nil {
			return err
		}
	}

	// Rank-6 tensors are legal for a const that only feeds constexpr_ ops,
	// so the consumers must be known before the rank check runs.
	consumers := make(map[string][]string)
	for _, bname := range sortedKeys(fn.BlockSpecializations) {
		collectConsumers(fn.BlockSpecializations[bname], consumers)
	}

	// Block specializations are alternative bodies for the same function, so
	// each starts from the same namespace and the same visible set.
	for _, bname := range sortedKeys(fn.BlockSpecializations) {
		blk := fn.BlockSpecializations[bname]
		names := make(map[string]bool, len(seen))
		inputs := make(map[string]bool, len(fn.Inputs))
		for n := range seen {
			names[n] = true
			inputs[n] = true
		}
		if err := validateBlockInvariants(blk, names, inputs, consumers, milOpSet(fn.OpSet)); err != nil {
			return fmt.Errorf("function %q block %q: %w", fname, bname, err)
		}
	}
	return nil
}

// collectConsumers records, for every var name, the types of the ops that read
// it, including from nested blocks.
func collectConsumers(blk *Block, consumers map[string][]string) {
	if blk == nil {
		return
	}
	for _, op := range blk.Operations {
		if op == nil {
			continue
		}
		for _, arg := range op.Inputs {
			if arg == nil {
				continue
			}
			for _, bind := range arg.Bindings {
				if bind.Name != "" {
					consumers[bind.Name] = append(consumers[bind.Name], op.Type)
				}
			}
		}
		for _, nested := range op.Blocks {
			collectConsumers(nested, consumers)
		}
	}
}

// validateBlockInvariants walks blk in order. names is the function-wide
// variable namespace, mutated as new definitions are seen; visible is the set
// of vars defined in enclosing blocks. opset is the MIL text opset name the
// enclosing function is declared with, used to pick the per-op schema.
func validateBlockInvariants(blk *Block, names, outer map[string]bool, consumers map[string][]string, opset string) error {
	if blk == nil {
		return nil
	}
	visible := make(map[string]bool, len(outer))
	for n := range outer {
		visible[n] = true
	}
	for _, in := range blk.Inputs {
		if names[in.Name] {
			return fmt.Errorf("duplicate variable name %q", in.Name)
		}
		names[in.Name] = true
		visible[in.Name] = true
	}

	for i, op := range blk.Operations {
		if op == nil {
			return fmt.Errorf("operation %d is nil", i)
		}
		if err := validateOpStructure(op); err != nil {
			return err
		}
		if err := validateOpSchema(op, opset); err != nil {
			return err
		}
		// Nested blocks see everything visible up to this op.
		for _, nested := range op.Blocks {
			if err := validateBlockInvariants(nested, names, visible, consumers, opset); err != nil {
				return err
			}
		}
		// Def-before-use: an operand must already be visible here.
		for _, iname := range sortedKeys(op.Inputs) {
			arg := op.Inputs[iname]
			if arg == nil {
				continue
			}
			for _, bind := range arg.Bindings {
				if bind.Name == "" {
					continue
				}
				if !visible[bind.Name] {
					return fmt.Errorf("op %s input %s references %q, which is not visible at that point", op.Type, iname, bind.Name)
				}
			}
		}
		for _, out := range op.Outputs {
			if names[out.Name] {
				return fmt.Errorf("duplicate variable name %q", out.Name)
			}
			names[out.Name] = true
			visible[out.Name] = true
			if err := validateOutputRank(op, out, consumers); err != nil {
				return err
			}
		}
	}

	for _, out := range blk.Outputs {
		if !visible[out] {
			return fmt.Errorf("block output %q is not visible in the block", out)
		}
	}
	return nil
}

// validateOpStructure checks the rules that depend only on a single op.
func validateOpStructure(op *Operation) error {
	// coremltools/converters/mil/mil/block.py:226-227 rejects an op with no
	// outputs; backend/mil/load.py:196-201 additionally requires const to
	// have exactly one.
	// write_state is the one op serialized with no outputs at all: the CoreML
	// backend builds it without an outputs= kwarg
	// (backend/mil/load.py:299-320), and its text statement has no "TYPE name ="
	// prefix.
	if len(op.Outputs) == 0 && op.Type != "write_state" {
		return fmt.Errorf("op %s has no outputs", op.Type)
	}
	if op.Type == "const" && len(op.Outputs) != 1 {
		return fmt.Errorf("const op must have exactly 1 output, got %d", len(op.Outputs))
	}
	for _, aname := range sortedKeys(op.Attributes) {
		if err := validateValue(op.Attributes[aname]); err != nil {
			return fmt.Errorf("op %s attribute %s: %w", op.Type, aname, err)
		}
	}
	for _, iname := range sortedKeys(op.Inputs) {
		arg := op.Inputs[iname]
		if arg == nil {
			continue
		}
		for i := range arg.Bindings {
			if err := validateValue(arg.Bindings[i].Value); err != nil {
				return fmt.Errorf("op %s input %s: %w", op.Type, iname, err)
			}
		}
	}
	// Only the coreml dialect survives serialization
	// (mil/program.py:129-140 _check_invalid_opset).
	if ns, _, ok := strings.Cut(op.Type, "::"); ok && ns != "coreml" {
		return fmt.Errorf("op %s is from dialect namespace %q, only coreml is supported", op.Type, ns)
	}
	return nil
}

// validateOpSchema checks op against the input specification its op class
// declares in coremltools, as dumped into the opschema table. opset is a MIL
// text opset name.
//
// The table covers only the ops we emit, so an op it does not record is
// skipped: absence there means "not dumped", not "does not exist". An op the
// table does record but not at this opset is an error, which is how
// coremltools reports it (ops/helper.py:31-34 "No available version for {} in
// the coremltools.target.{} opset").
func validateOpSchema(op *Operation, opset string) error {
	if !opschema.Known(opset) || !opschema.Registered(op.Type) {
		return nil
	}
	spec, ok := opschema.Lookup(opset, op.Type)
	if !ok {
		return fmt.Errorf("op %s is not available in opset %s", op.Type, opset)
	}
	// input_type.py:116-118 raises "Unrecognized input {}" for any input name
	// the spec does not declare.
	for _, iname := range sortedKeys(op.Inputs) {
		if _, ok := spec.Param(iname); !ok {
			return fmt.Errorf("op %s has no input %q in opset %s", op.Type, iname, opset)
		}
	}
	return nil
}

// validateOutputRank enforces the rank ceilings the Core ML runtime imposes.
func validateOutputRank(op *Operation, out NamedValueType, consumers map[string][]string) error {
	if out.Type == nil {
		return nil
	}
	if t := out.Type.TensorType; t != nil {
		r := tensorRank(t)
		if r > maxTensorRank {
			// A const or constexpr_ output read only by constexpr_ ops is
			// exempt: it is folded before it reaches the runtime.
			// mil/program.py:152-158.
			if op.Type == "const" || strings.HasPrefix(op.Type, "constexpr_") {
				if allConstexpr(consumers[out.Name]) {
					return nil
				}
			}
			return fmt.Errorf("op %s output %q has rank %d, Core ML supports rank <= %d", op.Type, out.Name, r, maxTensorRank)
		}
	}
	if l := out.Type.ListType; l != nil && l.ElementType != nil && l.ElementType.TensorType != nil {
		if r := tensorRank(l.ElementType.TensorType); r > maxListElementRank {
			return fmt.Errorf("op %s output %q is a list of rank-%d elements, Core ML supports rank <= %d", op.Type, out.Name, r, maxListElementRank)
		}
	}
	return nil
}

func allConstexpr(opTypes []string) bool {
	if len(opTypes) == 0 {
		return false
	}
	for _, t := range opTypes {
		if !strings.HasPrefix(t, "constexpr_") {
			return false
		}
	}
	return true
}

func tensorRank(t *TensorType) int {
	if len(t.Dimensions) > int(t.Rank) {
		return len(t.Dimensions)
	}
	return int(t.Rank)
}

// validateFunctionInput checks the rules that apply to a model boundary input.
func validateFunctionInput(fname string, in NamedValueType) error {
	if in.Type == nil {
		return fmt.Errorf("function %q input %q has nil type", fname, in.Name)
	}
	if t := in.Type.TensorType; t != nil {
		// mil/program.py:340-341: rank-0 model inputs are unsupported.
		if tensorRank(t) == 0 {
			return fmt.Errorf("function %q input %q is rank 0, which Core ML does not support", fname, in.Name)
		}
		// backend/mil/load.py:680-682: variadic rank is unsupported.
		for _, d := range t.Dimensions {
			if d.Variadic {
				return fmt.Errorf("function %q input %q has variadic rank", fname, in.Name)
			}
		}
	}
	// backend/mil/load.py:730-741: state inputs must be fp16 and statically
	// shaped.
	if st := in.Type.StateType; st != nil {
		if st.WrappedType == nil || st.WrappedType.TensorType == nil {
			return fmt.Errorf("function %q state input %q does not wrap a tensor type", fname, in.Name)
		}
		t := st.WrappedType.TensorType
		if t.DataType != DataTypeFloat16 {
			return fmt.Errorf("function %q state input %q has dtype %d, states must be fp16", fname, in.Name, t.DataType)
		}
		for _, d := range t.Dimensions {
			if d.Unknown || d.Variadic {
				return fmt.Errorf("function %q state input %q has a flexible shape, which states do not support", fname, in.Name)
			}
		}
	}
	return nil
}

// validateProgramOpsets checks that every function agrees on an opset.
// coremltools/converters/mil/mil/program.py:95-113.
func validateProgramOpsets(prog *Program) error {
	var first, firstName string
	for _, fname := range sortedKeys(prog.Functions) {
		fn := prog.Functions[fname]
		if fn == nil || fn.OpSet == "" {
			continue
		}
		if first == "" {
			first, firstName = fn.OpSet, fname
			continue
		}
		if fn.OpSet != first {
			return fmt.Errorf("all functions must have the same opset: function %q has %q, function %q has %q", firstName, first, fname, fn.OpSet)
		}
	}
	return nil
}

// validateModelProgram checks the invariants that relate the ModelDescription
// to the MIL program. It is separate from ValidateProgram because a bare
// Program carries no description.
func validateModelProgram(m *Model) error {
	if m == nil || m.MLProgram == nil {
		return nil
	}
	// backend/mil/load.py:1058-1061.
	if name := m.Description.DefaultFunctionName; name != "" {
		if _, ok := m.MLProgram.Functions[name]; !ok {
			return fmt.Errorf("default function %q not found in program", name)
		}
	}
	// backend/mil/load.py:963-968: multi-function export needs iOS18+.
	if len(m.Description.Functions) > 1 && m.SpecVersion < specVersionIOS18 {
		return fmt.Errorf("multi-function model requires specification version >= %d, got %d", specVersionIOS18, m.SpecVersion)
	}
	return nil
}

// validateValue rejects immediate values with no MIL text spelling.
// Float literals must be finite numerals: Go renders a non-finite float as
// "+Inf"/"NaN", tokens the MIL lexer does not accept, so an inf clamp bound or
// a NaN weight would emit a program that fails to compile with a syntax error
// far from the cause. coremltools never hits this because protobuf carries
// native float fields (backend/mil/helper.py:158-163).
func validateValue(v *Value) error {
	if v == nil {
		return nil
	}
	if v.Immediate == nil {
		return nil
	}
	if t := v.Immediate.Tensor; t != nil {
		// The MIL reader picks the TensorValue field from the declared
		// element type, so a value stored in any other field reads as empty.
		if v.Type != nil && v.Type.TensorType != nil {
			if err := ValidateTensorValue(v.Type.TensorType.DataType, t); err != nil {
				return err
			}
		}
		for _, f := range t.Floats {
			if err := checkFinite(float64(f)); err != nil {
				return err
			}
		}
		for _, f := range t.Doubles {
			if err := checkFinite(f); err != nil {
				return err
			}
		}
	}
	if l := v.Immediate.List; l != nil {
		for _, elem := range l.Values {
			if err := validateValue(elem); err != nil {
				return err
			}
		}
	}
	if tp := v.Immediate.Tuple; tp != nil {
		for _, elem := range tp.Values {
			if err := validateValue(elem); err != nil {
				return err
			}
		}
	}
	if d := v.Immediate.Dictionary; d != nil {
		for _, entry := range d.Entries {
			if err := validateValue(entry.Key); err != nil {
				return err
			}
			if err := validateValue(entry.Value); err != nil {
				return err
			}
		}
	}
	return nil
}

func checkFinite(f float64) error {
	if math.IsInf(f, 0) || math.IsNaN(f) {
		return fmt.Errorf("non-finite float value %v has no MIL text spelling", f)
	}
	return nil
}

// validateProgramAttributes restricts program-level attributes to buildInfo.
// coremltools/converters/mil/frontend/milproto/load.py:561-571 raises
// "Invalid attribute ... for program" for anything else, so a program carrying
// another key cannot be read back by Apple's own tooling.
func validateProgramAttributes(prog *Program) error {
	for _, key := range sortedKeys(prog.Attributes) {
		if key != "buildInfo" {
			return fmt.Errorf("invalid program attribute %q, only buildInfo is allowed", key)
		}
		if err := validateValue(prog.Attributes[key]); err != nil {
			return fmt.Errorf("program attribute %q: %w", key, err)
		}
	}
	return nil
}
