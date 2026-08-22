package coremlcompiler

import (
	"fmt"
	"strconv"
	"strings"
)

// milReservedNames are MIL keywords and type names that cannot be used as
// identifiers. The list mirrors coremltools' NameSanitizer.
var milReservedNames = map[string]bool{
	"any": true, "bool": true, "program": true, "func": true, "tensor": true,
	"list": true, "dict": true, "tuple": true, "true": true, "false": true,
	"string": true, "bf16": true, "fp16": true, "fp32": true, "fp64": true,
	"int8": true, "int16": true, "int32": true, "int64": true,
	"uint8": true, "uint16": true, "uint32": true, "uint64": true,
	"state": true,
}

// validMILName reports whether name is a legal MIL identifier:
// [A-Za-z_][A-Za-z0-9_]* and not a reserved word.
func validMILName(name string) bool {
	if name == "" || milReservedNames[name] {
		return false
	}
	for i := 0; i < len(name); i++ {
		c := name[i]
		switch {
		case c == '_':
		case c >= 'A' && c <= 'Z':
		case c >= 'a' && c <= 'z':
		case c >= '0' && c <= '9':
			if i == 0 {
				return false
			}
		default:
			return false
		}
	}
	return true
}

// nameSanitizer rewrites arbitrary source-framework names into legal, unique
// MIL identifiers. It mirrors coremltools' NameSanitizer so that names that
// survive a coremltools conversion survive ours identically.
type nameSanitizer struct {
	seen map[string]bool
}

func newNameSanitizer() *nameSanitizer {
	return &nameSanitizer{seen: make(map[string]bool)}
}

func (s *nameSanitizer) sanitize(name string) string {
	var b strings.Builder
	for _, r := range name {
		switch {
		case r == '_',
			r >= 'A' && r <= 'Z',
			r >= 'a' && r <= 'z',
			r >= '0' && r <= '9':
			b.WriteRune(r)
		default:
			b.WriteByte('_')
		}
	}
	newName := b.String()

	if newName == "" || !(newName[0] == '_' || (newName[0] >= 'A' && newName[0] <= 'Z') || (newName[0] >= 'a' && newName[0] <= 'z')) {
		newName = "_" + newName
	}
	if milReservedNames[newName] {
		newName += "_workaround"
	}

	if newName == name {
		s.seen[name] = true
		return name
	}
	// The name changed, so it may now collide with one already issued.
	if s.seen[newName] {
		base := newName
		for idx := 0; ; idx++ {
			newName = base + "_" + strconv.Itoa(idx)
			if !s.seen[newName] {
				break
			}
			base = newName
		}
	}
	s.seen[newName] = true
	return newName
}

// ValidateNames reports the first identifier in prog that is not a legal MIL
// identifier. Invalid names are emitted verbatim by the MIL text writer and
// produce a syntactically invalid program that only fails, opaquely, inside
// Apple's compiler. Use SanitizeProgram to repair them.
func ValidateNames(prog *Program) error {
	if prog == nil {
		return nil
	}
	for _, fname := range sortedKeys(prog.Functions) {
		if !validMILName(fname) {
			return fmt.Errorf("function name %q is not a valid MIL identifier", fname)
		}
		fn := prog.Functions[fname]
		if fn == nil {
			continue
		}
		for _, in := range fn.Inputs {
			if !validMILName(in.Name) {
				return fmt.Errorf("function %q input name %q is not a valid MIL identifier", fname, in.Name)
			}
		}
		for _, bname := range sortedKeys(fn.BlockSpecializations) {
			if err := validateBlockNames(fn.BlockSpecializations[bname]); err != nil {
				return fmt.Errorf("function %q block %q: %w", fname, bname, err)
			}
		}
	}
	return nil
}

func validateBlockNames(blk *Block) error {
	if blk == nil {
		return nil
	}
	for _, in := range blk.Inputs {
		if !validMILName(in.Name) {
			return fmt.Errorf("block input name %q is not a valid MIL identifier", in.Name)
		}
	}
	for _, out := range blk.Outputs {
		if !validMILName(out) {
			return fmt.Errorf("block output name %q is not a valid MIL identifier", out)
		}
	}
	for _, op := range blk.Operations {
		if op == nil {
			continue
		}
		for _, out := range op.Outputs {
			if !validMILName(out.Name) {
				return fmt.Errorf("op %s output name %q is not a valid MIL identifier", op.Type, out.Name)
			}
		}
		for _, iname := range sortedKeys(op.Inputs) {
			for _, bind := range op.Inputs[iname].Bindings {
				if bind.Name != "" && !validMILName(bind.Name) {
					return fmt.Errorf("op %s input %s references %q, not a valid MIL identifier", op.Type, iname, bind.Name)
				}
			}
		}
		for _, nested := range op.Blocks {
			if err := validateBlockNames(nested); err != nil {
				return err
			}
		}
	}
	return nil
}

// SanitizeProgram rewrites every function, variable and block name in m's MIL
// program into a legal MIL identifier, and applies the same rewrite to the
// matching ModelDescription feature names so the description and the MIL stay
// in sync. Names that are already valid are left untouched.
func SanitizeProgram(m *Model) {
	if m == nil || m.MLProgram == nil {
		return
	}
	prog := m.MLProgram

	fnNames := newNameSanitizer()
	renamedFuncs := make(map[string]*Function, len(prog.Functions))
	var mainVars map[string]string
	for _, fname := range sortedKeys(prog.Functions) {
		fn := prog.Functions[fname]
		vars := sanitizeFunctionVars(fn)
		if fname == "main" {
			mainVars = vars
		}
		renamedFuncs[fnNames.sanitize(fname)] = fn
	}
	prog.Functions = renamedFuncs

	// The description names the main function's boundary vars; rename them
	// the same way or the compiled model loses its inputs and outputs.
	if mainVars != nil {
		renameFeatures(m.Description.Inputs, mainVars)
		renameFeatures(m.Description.Outputs, mainVars)
		renameFeatures(m.Description.States, mainVars)
	}
}

func renameFeatures(feats []FeatureDescription, renames map[string]string) {
	for i := range feats {
		if newName, ok := renames[feats[i].Name]; ok {
			feats[i].Name = newName
		}
	}
}

// sanitizeFunctionVars renames every variable in fn and returns the old-to-new
// mapping for the names visible at the function boundary.
func sanitizeFunctionVars(fn *Function) map[string]string {
	if fn == nil {
		return nil
	}
	s := newNameSanitizer()
	renames := make(map[string]string)
	for i := range fn.Inputs {
		old := fn.Inputs[i].Name
		fn.Inputs[i].Name = s.sanitize(old)
		renames[old] = fn.Inputs[i].Name
	}
	for _, bname := range sortedKeys(fn.BlockSpecializations) {
		blk := fn.BlockSpecializations[bname]
		sanitizeBlockDefs(blk, s, renames)
		// References are resolved only after every definition is renamed so
		// that a forward reference inside a nested block still resolves.
		rewriteBlockRefs(blk, renames)
	}
	return renames
}

func sanitizeBlockDefs(blk *Block, s *nameSanitizer, renames map[string]string) {
	if blk == nil {
		return
	}
	for i := range blk.Inputs {
		old := blk.Inputs[i].Name
		blk.Inputs[i].Name = s.sanitize(old)
		renames[old] = blk.Inputs[i].Name
	}
	for _, op := range blk.Operations {
		if op == nil {
			continue
		}
		for i := range op.Outputs {
			old := op.Outputs[i].Name
			op.Outputs[i].Name = s.sanitize(old)
			renames[old] = op.Outputs[i].Name
		}
		for _, nested := range op.Blocks {
			sanitizeBlockDefs(nested, s, renames)
		}
	}
}

func rewriteBlockRefs(blk *Block, renames map[string]string) {
	for i, out := range blk.Outputs {
		if newName, ok := renames[out]; ok {
			blk.Outputs[i] = newName
		}
	}
	for _, op := range blk.Operations {
		if op == nil {
			continue
		}
		for _, iname := range sortedKeys(op.Inputs) {
			arg := op.Inputs[iname]
			for i := range arg.Bindings {
				if newName, ok := renames[arg.Bindings[i].Name]; ok {
					arg.Bindings[i].Name = newName
				}
			}
		}
		for _, nested := range op.Blocks {
			if nested != nil {
				rewriteBlockRefs(nested, renames)
			}
		}
	}
}
