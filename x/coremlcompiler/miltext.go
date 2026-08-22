package coremlcompiler

import (
	"fmt"
	"math"
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

// milOpSet maps proto opset names to MIL text opset names.
// Apple's coremlcompiler rewrites CoreMLN → iosNN in MIL text. The pairing
// follows coremltools' _OPSET table (coremltools/__init__.py): CoreMLN is the
// opset for spec version N+1, which is the iOS N+10 release.
func milOpSet(opset string) string {
	switch opset {
	case "CoreML3":
		return "ios13"
	case "CoreML4":
		return "ios14"
	case "CoreML5":
		return "ios15"
	case "CoreML6":
		return "ios16"
	case "CoreML7":
		return "ios17"
	case "CoreML8":
		return "ios18"
	case "CoreML9":
		return "ios26"
	default:
		return opset
	}
}

// ValidateProgram validates structural requirements of an MIL Program.
func ValidateProgram(prog *Program) error {
	if prog == nil {
		return fmt.Errorf("nil program")
	}
	for fname, fn := range prog.Functions {
		if fn == nil {
			return fmt.Errorf("function %q is nil", fname)
		}
		if len(fn.BlockSpecializations) == 0 {
			return fmt.Errorf("function %q has no block specializations", fname)
		}
		if fn.OpSet != "" {
			if _, ok := fn.BlockSpecializations[fn.OpSet]; !ok {
				return fmt.Errorf("function %q does not have block specialization for opset %q", fname, fn.OpSet)
			}
		}
		for bname, blk := range fn.BlockSpecializations {
			if blk == nil {
				return fmt.Errorf("function %q block specialization %q is nil", fname, bname)
			}
			if err := validateBlock(blk); err != nil {
				return fmt.Errorf("function %q block %q: %w", fname, bname, err)
			}
		}
		if err := validateFunction(fname, fn); err != nil {
			return err
		}
	}
	if err := validateProgramOpsets(prog); err != nil {
		return err
	}
	if err := validateProgramAttributes(prog); err != nil {
		return err
	}
	if err := ValidateNames(prog); err != nil {
		return fmt.Errorf("invalid name: %w", err)
	}
	return nil
}

func validateBlock(blk *Block) error {
	for i, op := range blk.Operations {
		if op == nil {
			return fmt.Errorf("operation %d is nil", i)
		}
		for _, nvt := range op.Outputs {
			if nvt.Type == nil {
				return fmt.Errorf("op %s output %q has nil type", op.Type, nvt.Name)
			}
		}
		for _, nested := range op.Blocks {
			if nested != nil {
				if err := validateBlock(nested); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func emitFunction(b *strings.Builder, name string, fn *Function) {
	opsetSuffix := ""
	if fn.OpSet != "" {
		opsetSuffix = "<" + milOpSet(fn.OpSet) + ">"
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
	// Prefer the block matching fn.OpSet; fall back to first sorted name.
	blockNames := sortedKeys(fn.BlockSpecializations)
	if len(blockNames) > 0 {
		targetBlock := blockNames[0]
		for _, bname := range blockNames {
			if bname == fn.OpSet {
				targetBlock = bname
				break
			}
		}
		block := fn.BlockSpecializations[targetBlock]
		emitBlockBody(b, block, 2)
	} else {
		b.WriteString("    }\n")
	}
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

	// For const ops, all inputs (e.g. val) are emitted as attributes
	// in MIL text format: const()[name = ..., val = ...]
	isConst := op.Type == "const"

	// Operation name and inputs.
	fmt.Fprintf(b, "%s(", op.Type)
	if !isConst {
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
	}
	b.WriteString(")")

	// Attributes (for const ops, inputs are merged here).
	hasAttrs := len(op.Attributes) > 0 || (isConst && len(op.Inputs) > 0)
	if hasAttrs {
		b.WriteString("[")
		first := true
		attrNames := sortedKeys(op.Attributes)
		for _, aname := range attrNames {
			if !first {
				b.WriteString(", ")
			}
			first = false
			fmt.Fprintf(b, "%s = %s", aname, formatValue(op.Attributes[aname]))
		}
		if isConst {
			inputNames := sortedKeys(op.Inputs)
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
	if vt.ListType != nil {
		if vt.ListType.Length > 0 {
			return fmt.Sprintf("list<%s, %d>", formatType(vt.ListType.ElementType), vt.ListType.Length)
		}
		return fmt.Sprintf("list<%s>", formatType(vt.ListType.ElementType))
	}
	if vt.TupleType != nil {
		parts := make([]string, len(vt.TupleType.Types))
		for i, elem := range vt.TupleType.Types {
			parts[i] = formatType(elem)
		}
		return fmt.Sprintf("tuple<%s>", strings.Join(parts, ", "))
	}
	if vt.DictionaryType != nil {
		return fmt.Sprintf("dict<%s, %s>", formatType(vt.DictionaryType.KeyType), formatType(vt.DictionaryType.ValueType))
	}
	if vt.StateType != nil {
		return fmt.Sprintf("state<%s>", formatType(vt.StateType.WrappedType))
	}
	return "<<unknown>>"
}

// formatTensorType formats a tensor type like "tensor<fp16, [1, 128, 1, 64]>".
// Scalars are formatted as "tensor<dtype, []>" to match Apple's coremlcompiler.
func formatTensorType(tt *TensorType) string {
	dtype := tt.DataType.String()

	// Scalar (rank 0 or no dimensions).
	if len(tt.Dimensions) == 0 {
		return fmt.Sprintf("tensor<%s, []>", dtype)
	}

	var dims []string
	for _, d := range tt.Dimensions {
		if d.Unknown {
			if d.Variadic {
				dims = append(dims, "*?")
			} else {
				dims = append(dims, "?")
			}
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
	if v.Immediate != nil {
		if v.Immediate.Tensor != nil {
			return formatImmediateTensor(v)
		}
		if v.Immediate.Tuple != nil {
			parts := make([]string, len(v.Immediate.Tuple.Values))
			for i, elem := range v.Immediate.Tuple.Values {
				parts[i] = formatValue(elem)
			}
			return fmt.Sprintf("%s((%s))", formatType(v.Type), strings.Join(parts, ", "))
		}
		if v.Immediate.List != nil {
			parts := make([]string, len(v.Immediate.List.Values))
			for i, elem := range v.Immediate.List.Values {
				parts[i] = formatValue(elem)
			}
			return fmt.Sprintf("%s([%s])", formatType(v.Type), strings.Join(parts, ", "))
		}
		if v.Immediate.Dictionary != nil {
			// Dictionary entries are brace-wrapped pairs inside one outer
			// brace pair: dict<K, V>({{k1, v1}, {k2, v2}}). Apple's coremlc
			// prints buildInfo this way (see the ground-truth model.mil in
			// coremltools/modelrunner); there is no "key: value" spelling.
			parts := make([]string, len(v.Immediate.Dictionary.Entries))
			for i, entry := range v.Immediate.Dictionary.Entries {
				parts[i] = fmt.Sprintf("{%s, %s}", formatValueBare(entry.Key), formatValueBare(entry.Value))
			}
			return fmt.Sprintf("%s({%s})", formatType(v.Type), strings.Join(parts, ", "))
		}
	}
	return "<<empty>>"
}

// formatValueBare formats a value without its type prefix, the spelling used
// for elements inside a dictionary literal: coremlc prints
// {"coremlc-version", "3402.4.1"}, not {tensor<string, []>("coremlc-version"), ...}.
func formatValueBare(v *Value) string {
	if v == nil {
		return "<<nil>>"
	}
	s := formatValue(v)
	if v.Type != nil {
		if t := formatType(v.Type); strings.HasPrefix(s, t+"(") && strings.HasSuffix(s, ")") {
			return s[len(t)+1 : len(s)-1]
		}
	}
	return s
}

// quoteMILString writes a MIL string literal. MIL string values are arbitrary
// UTF-8 (helper.py encodes them raw), so a quote or backslash in a metadata
// value or a weight path would otherwise terminate the literal early and make
// the whole program unparseable.
func quoteMILString(s string) string {
	var b strings.Builder
	b.WriteByte('"')
	for i := 0; i < len(s); i++ {
		switch c := s[i]; c {
		case '"', '\\':
			b.WriteByte('\\')
			b.WriteByte(c)
		case '\n':
			b.WriteString(`\n`)
		case '\r':
			b.WriteString(`\r`)
		case '\t':
			b.WriteString(`\t`)
		default:
			b.WriteByte(c)
		}
	}
	b.WriteByte('"')
	return b.String()
}

// formatBlobFileRef formats a BLOBFILE reference.
// Real Apple format: tensor<fp16, [...]>(BLOBFILE(path = tensor<string, []>("@model_path/..."), offset = tensor<uint64, []>(64)))
func formatBlobFileRef(v *Value) string {
	bf := v.BlobFile
	typStr := formatType(v.Type)
	return fmt.Sprintf("%s(BLOBFILE(path = tensor<string, []>(%s), offset = tensor<uint64, []>(%d)))",
		typStr, quoteMILString(bf.FileName), bf.Offset)
}

// formatImmediateTensor formats an inline tensor value.
func formatImmediateTensor(v *Value) string {
	tv := v.Immediate.Tensor
	typeStr := formatType(v.Type)

	// Check if this is a scalar type (no tensor wrapper).
	isScalar := v.Type != nil && v.Type.TensorType != nil && len(v.Type.TensorType.Dimensions) == 0

	format := func(vals string) string {
		return formatTypedValues(typeStr, isScalar, vals)
	}
	formatMulti := func(parts []string) string {
		if dims := immediateDims(v, len(parts)); dims != nil {
			return formatNestedValues(typeStr, parts, dims)
		}
		return format(strings.Join(parts, ", "))
	}

	switch {
	case tv.Floats != nil:
		return formatMulti(float32Parts(tv.Floats))
	case tv.Doubles != nil:
		return formatMulti(float64Parts(tv.Doubles))
	case tv.Ints != nil:
		return formatMulti(int32Parts(tv.Ints))
	case tv.Longs != nil:
		return formatMulti(int64Parts(tv.Longs))
	case tv.Bools != nil:
		return formatMulti(boolParts(tv.Bools))
	case tv.Strings != nil:
		// Strings may contain ", " so they cannot be re-split; multi-dim
		// string immediates stay flat (none exist in practice).
		return format(formatStringSlice(tv.Strings))
	case tv.Bytes != nil:
		if v.Type != nil && v.Type.TensorType != nil {
			switch v.Type.TensorType.DataType {
			case DataTypeFloat16:
				if len(tv.Bytes)%2 == 0 {
					parts := make([]string, len(tv.Bytes)/2)
					for i := 0; i < len(tv.Bytes); i += 2 {
						bits := uint16(tv.Bytes[i]) | uint16(tv.Bytes[i+1])<<8
						parts[i/2] = fmt.Sprintf("%g", float16ToFloat32(bits))
					}
					return formatMulti(parts)
				}
			case DataTypeInt8:
				parts := make([]string, len(tv.Bytes))
				for i, b := range tv.Bytes {
					parts[i] = fmt.Sprintf("%d", int8(b))
				}
				return formatMulti(parts)
			case DataTypeUInt8:
				parts := make([]string, len(tv.Bytes))
				for i, b := range tv.Bytes {
					parts[i] = fmt.Sprintf("%d", b)
				}
				return formatMulti(parts)
			case DataTypeInt16:
				if len(tv.Bytes)%2 == 0 {
					parts := make([]string, len(tv.Bytes)/2)
					for i := 0; i < len(tv.Bytes); i += 2 {
						val := int16(uint16(tv.Bytes[i]) | uint16(tv.Bytes[i+1])<<8)
						parts[i/2] = fmt.Sprintf("%d", val)
					}
					return formatMulti(parts)
				}
			case DataTypeUInt16:
				if len(tv.Bytes)%2 == 0 {
					parts := make([]string, len(tv.Bytes)/2)
					for i := 0; i < len(tv.Bytes); i += 2 {
						val := uint16(tv.Bytes[i]) | uint16(tv.Bytes[i+1])<<8
						parts[i/2] = fmt.Sprintf("%d", val)
					}
					return formatMulti(parts)
				}
			}
		}
		return fmt.Sprintf("%s(<<bytes:%d>>)", typeStr, len(tv.Bytes))
	default:
		return typeStr + "()"
	}
}

func float16ToFloat32(h uint16) float32 {
	sign := uint32(h&0x8000) << 16
	exp := uint32(h&0x7C00) >> 10
	mant := uint32(h & 0x03FF)
	if exp == 0 {
		if mant == 0 {
			return math.Float32frombits(sign)
		}
		for (mant & 0x0400) == 0 {
			mant <<= 1
			exp--
		}
		exp++
		mant &= 0x03FF
	} else if exp == 0x1F {
		if mant == 0 {
			return math.Float32frombits(sign | 0x7F800000)
		}
		return math.Float32frombits(sign | 0x7F800000 | (mant << 13))
	}
	exp = exp + (127 - 15)
	return math.Float32frombits(sign | (exp << 23) | (mant << 13))
}

func formatTypedValues(typeStr string, isScalar bool, vals string) string {
	if isScalar {
		return fmt.Sprintf("%s(%s)", typeStr, vals)
	}
	return fmt.Sprintf("%s([%s])", typeStr, vals)
}

// formatNestedValues formats already-formatted element strings as a tensor
// literal nested to match dims. MIL text requires the literal's bracket
// structure to match the declared shape: a flat list of 12 elements parses
// as shape [12], not [3, 4].
func formatNestedValues(typeStr string, parts []string, dims []uint64) string {
	return fmt.Sprintf("%s(%s)", typeStr, nestValues(parts, dims))
}

func nestValues(parts []string, dims []uint64) string {
	if len(dims) <= 1 {
		return "[" + strings.Join(parts, ", ") + "]"
	}
	stride := len(parts) / int(dims[0])
	groups := make([]string, dims[0])
	for i := range groups {
		groups[i] = nestValues(parts[i*stride:(i+1)*stride], dims[1:])
	}
	return "[" + strings.Join(groups, ", ") + "]"
}

// immediateDims returns the constant dimensions of v's tensor type when the
// literal should be nested: rank >= 2, all dimensions constant, and the
// element count matching their product. Any other shape returns nil and the
// literal stays flat.
func immediateDims(v *Value, count int) []uint64 {
	if v.Type == nil || v.Type.TensorType == nil {
		return nil
	}
	tt := v.Type.TensorType
	if len(tt.Dimensions) < 2 {
		return nil
	}
	dims := make([]uint64, len(tt.Dimensions))
	product := uint64(1)
	for i, d := range tt.Dimensions {
		if d.Unknown || d.Constant == 0 {
			return nil
		}
		dims[i] = d.Constant
		product *= d.Constant
	}
	if product != uint64(count) {
		return nil
	}
	return dims
}

func float32Parts(vals []float32) []string {
	parts := make([]string, len(vals))
	for i, v := range vals {
		parts[i] = formatFloat(float64(v))
	}
	return parts
}

func float64Parts(vals []float64) []string {
	parts := make([]string, len(vals))
	for i, v := range vals {
		parts[i] = formatFloat(v)
	}
	return parts
}

func formatFloat(v float64) string {
	return fmt.Sprintf("%g", v)
}

func int32Parts(vals []int32) []string {
	parts := make([]string, len(vals))
	for i, v := range vals {
		parts[i] = fmt.Sprintf("%d", v)
	}
	return parts
}

func int64Parts(vals []int64) []string {
	parts := make([]string, len(vals))
	for i, v := range vals {
		parts[i] = fmt.Sprintf("%d", v)
	}
	return parts
}

func boolParts(vals []bool) []string {
	parts := make([]string, len(vals))
	for i, v := range vals {
		if v {
			parts[i] = "true"
		} else {
			parts[i] = "false"
		}
	}
	return parts
}

func formatStringSlice(vals []string) string {
	parts := make([]string, len(vals))
	for i, v := range vals {
		parts[i] = quoteMILString(v)
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
// Apple's coremlcompiler emits "1.0" for all current spec versions.
func milVersionForSpec(specVersion int32) string {
	return "1.0"
}

// emitMILTextWithSpec is the primary entry point, using spec version for
// correct MIL dialect selection.
func emitMILTextWithSpec(prog *Program, specVersion int32) string {
	var b strings.Builder
	fmt.Fprintf(&b, "program(%s)\n", milVersionForSpec(specVersion))

	for _, key := range sortedKeys(prog.Attributes) {
		val := prog.Attributes[key]
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
