package coremlcompiler

import (
	"fmt"
	"sort"
	"strings"
)

// emitMILText converts a parsed MIL Program to text format using spec version 8
// (MIL 1.3 dialect) as default. Use emitMILTextWithSpec for explicit control.
//
// The output matches the format consumed by CoreML's runtime and the ANE
// compiler. Key format details:
//   - Function parameters use bare names: tensor<fp32, [1, 64]> x
//   - Variable references use bare names in op outputs: tensor<fp16, ...> y = ...
//   - Block specializations are flattened into the Function body
//   - Const ops put value in attributes: const()[name = ..., val = ...]
func emitMILText(prog *Program) string {
	return emitMILTextWithSpec(prog, 8)
}

func emitFunction(b *strings.Builder, name string, fn *Function) {
	opsetSuffix := ""
	if fn.OpSet != "" {
		opsetSuffix = "<" + fn.OpSet + ">"
	}

	// Function signature — params use bare names (no % prefix).
	fmt.Fprintf(b, "    func %s%s(", name, opsetSuffix)
	for i, input := range fn.Inputs {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(b, "%s %s", formatType(input.Type), input.Name)
	}
	b.WriteString(") {\n")

	// Flatten Block specializations into the Function body.
	// Use the first Block (sorted by name for determinism).
	blockNames := sortedKeys(fn.BlockSpecializations)
	for _, bname := range blockNames {
		Block := fn.BlockSpecializations[bname]
		emitBlockBody(b, Block, 2)
		// Only emit the first Block's content; the Block name
		// is a proto-level concept not present in MIL text.
		_ = bname
		break
	}

	b.WriteString("    }\n")
}

// emitBlockBody emits the operations and return statement of a Block.
func emitBlockBody(b *strings.Builder, Block *Block, indent int) {
	for _, op := range Block.Operations {
		emitOperation(b, op, indent)
	}

	// Block return.
	prefix := strings.Repeat("    ", indent-1)
	fmt.Fprintf(b, "%s} -> (%s);\n", prefix, strings.Join(Block.Outputs, ", "))
}

func emitOperation(b *strings.Builder, op *Operation, indent int) {
	prefix := strings.Repeat("    ", indent)

	// Output declarations: type %name = op(...)
	b.WriteString(prefix)
	for i, out := range op.Outputs {
		if i > 0 {
			b.WriteString(", ")
		}
		fmt.Fprintf(b, "%s %s", formatType(out.Type), out.Name)
	}
	if len(op.Outputs) > 0 {
		b.WriteString(" = ")
	}

	// Operation name and inputs.
	fmt.Fprintf(b, "%s(", op.Type)

	// Emit inputs in sorted order.
	inputNames := sortedKeys(op.Inputs)
	first := true
	for _, iname := range inputNames {
		arg := op.Inputs[iname]
		for _, Binding := range arg.Bindings {
			if !first {
				b.WriteString(", ")
			}
			first = false
			fmt.Fprintf(b, "%s = %s", iname, formatBinding(&Binding))
		}
	}
	b.WriteString(")")

	// Attributes.
	if len(op.Attributes) > 0 {
		b.WriteString("[")
		attrNames := sortedKeys(op.Attributes)
		for i, aname := range attrNames {
			if i > 0 {
				b.WriteString(", ")
			}
			fmt.Fprintf(b, "%s = %s", aname, formatValue(op.Attributes[aname]))
		}
		b.WriteString("]")
	}

	b.WriteString(";\n")

	// Nested blocks (for control flow ops like cond, while_loop).
	for _, blk := range op.Blocks {
		nestedPrefix := strings.Repeat("    ", indent+1)
		// Nested Block header.
		fmt.Fprintf(b, "%sblock(", nestedPrefix)
		for i, input := range blk.Inputs {
			if i > 0 {
				b.WriteString(", ")
			}
			fmt.Fprintf(b, "%s %s", formatType(input.Type), input.Name)
		}
		b.WriteString(") {\n")
		emitBlockBody(b, blk, indent+2)
	}
}

// formatType converts a ValueType to MIL text representation.
func formatType(vt *ValueType) string {
	if vt == nil {
		return "<<unknown>>"
	}
	if vt.TensorType != nil {
		return formatTensorType(vt.TensorType)
	}
	if vt.StateType != nil {
		return fmt.Sprintf("state<%s>", formatType(vt.StateType.WrappedType))
	}
	return "<<unknown>>"
}

// formatTensorType formats a tensor type like "tensor<fp16, [1, 128, 1, 64]>".
func formatTensorType(tt *TensorType) string {
	dtype := tt.DataType.String()

	// Scalar (rank 0 or no dimensions).
	if len(tt.Dimensions) == 0 {
		return dtype
	}

	var dims []string
	for _, d := range tt.Dimensions {
		if d.Unknown {
			dims = append(dims, "?")
		} else {
			dims = append(dims, fmt.Sprintf("%d", d.Constant))
		}
	}
	return fmt.Sprintf("tensor<%s, [%s]>", dtype, strings.Join(dims, ", "))
}

// formatBinding formats an Argument Binding (name ref or inline value).
func formatBinding(b *Binding) string {
	if b.Name != "" {
		return b.Name
	}
	if b.Value != nil {
		return formatValue(b.Value)
	}
	return "<<nil>>"
}

// formatValue formats a Value for MIL text output.
func formatValue(v *Value) string {
	if v.BlobFile != nil {
		return formatBlobFileRef(v)
	}
	if v.Immediate != nil && v.Immediate.Tensor != nil {
		return formatImmediateTensor(v)
	}
	return "<<empty>>"
}

// formatBlobFileRef formats a BLOBFILE reference.
// Real Apple format: tensor<fp16, [...]>(BLOBFILE(path = tensor<string, []>("@model_path/..."), offset = tensor<uint64, []>(64)))
func formatBlobFileRef(v *Value) string {
	bf := v.BlobFile
	typStr := formatType(v.Type)
	return fmt.Sprintf("%s(BLOBFILE(path = tensor<string, []>(\"%s\"), offset = tensor<uint64, []>(%d)))",
		typStr, bf.FileName, bf.Offset)
}

// formatImmediateTensor formats an inline tensor value.
func formatImmediateTensor(v *Value) string {
	tv := v.Immediate.Tensor
	typeStr := formatType(v.Type)

	// Check if this is a scalar type (no tensor wrapper).
	isScalar := v.Type != nil && v.Type.TensorType != nil && len(v.Type.TensorType.Dimensions) == 0

	switch {
	case tv.Floats != nil:
		return formatTypedValues(typeStr, isScalar, formatFloat32Slice(tv.Floats))
	case tv.Doubles != nil:
		return formatTypedValues(typeStr, isScalar, formatFloat64Slice(tv.Doubles))
	case tv.Ints != nil:
		return formatTypedValues(typeStr, isScalar, formatInt32Slice(tv.Ints))
	case tv.Longs != nil:
		return formatTypedValues(typeStr, isScalar, formatInt64Slice(tv.Longs))
	case tv.Bools != nil:
		return formatTypedValues(typeStr, isScalar, formatBoolSlice(tv.Bools))
	case tv.Strings != nil:
		return formatTypedValues(typeStr, isScalar, formatStringSlice(tv.Strings))
	case tv.Bytes != nil:
		return fmt.Sprintf("%s(<<bytes:%d>>)", typeStr, len(tv.Bytes))
	default:
		return typeStr + "()"
	}
}

func formatTypedValues(typeStr string, isScalar bool, vals string) string {
	if isScalar {
		return fmt.Sprintf("%s(%s)", typeStr, vals)
	}
	return fmt.Sprintf("%s([%s])", typeStr, vals)
}

func formatFloat32Slice(vals []float32) string {
	parts := make([]string, len(vals))
	for i, v := range vals {
		parts[i] = formatFloat(float64(v))
	}
	return strings.Join(parts, ", ")
}

func formatFloat64Slice(vals []float64) string {
	parts := make([]string, len(vals))
	for i, v := range vals {
		parts[i] = formatFloat(v)
	}
	return strings.Join(parts, ", ")
}

func formatFloat(v float64) string {
	return fmt.Sprintf("%g", v)
}

func formatInt32Slice(vals []int32) string {
	parts := make([]string, len(vals))
	for i, v := range vals {
		parts[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(parts, ", ")
}

func formatInt64Slice(vals []int64) string {
	parts := make([]string, len(vals))
	for i, v := range vals {
		parts[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(parts, ", ")
}

func formatBoolSlice(vals []bool) string {
	parts := make([]string, len(vals))
	for i, v := range vals {
		if v {
			parts[i] = "true"
		} else {
			parts[i] = "false"
		}
	}
	return strings.Join(parts, ", ")
}

func formatStringSlice(vals []string) string {
	parts := make([]string, len(vals))
	for i, v := range vals {
		parts[i] = fmt.Sprintf("\"%s\"", v)
	}
	return strings.Join(parts, ", ")
}

func sortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// milVersionForSpec maps the CoreML spec version to the MIL dialect string.
//
//	spec 7  → "1.0" (iOS 17)
//	spec 8+ → "1.3" (iOS 18)
func milVersionForSpec(specVersion int32) string {
	if specVersion >= 8 {
		return "1.3"
	}
	return "1.0"
}

// emitMILTextWithSpec is the primary entry point, using spec version for
// correct MIL dialect selection.
func emitMILTextWithSpec(prog *Program, specVersion int32) string {
	var b strings.Builder
	fmt.Fprintf(&b, "Program(%s)\n", milVersionForSpec(specVersion))

	for key, val := range prog.Attributes {
		fmt.Fprintf(&b, "[%s = %s]\n", key, formatValue(val))
	}

	b.WriteString("{\n")

	names := sortedKeys(prog.Functions)
	for _, name := range names {
		fn := prog.Functions[name]
		emitFunction(&b, name, fn)
	}

	b.WriteString("}\n")
	return b.String()
}
