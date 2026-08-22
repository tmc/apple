// Copyright 2026 The tmc/apple Authors. All rights reserved.

package signaturegate

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"testing"
)

// The width of a generated enum type is public API and it is checkable against
// the SDK, so it is checked against the SDK. Comparing it to any other artifact
// of the same pipeline compares two things that can be wrong the same way.
//
// The join key is the CONSTANT name, not the type name. A generated enum type's
// name is frequently synthesized from the longest common prefix of its
// constants -- mach_vm_range_tag_t is emitted as MachVmRange -- so matching on
// type names would silently drop exactly the enums most likely to be wrong.
// Constants keep their C spelling on both sides.

// __enum_decl(name, backing_type, { ... }) is the SDK spelling that carries an
// explicit width. Enums without it are outside this gate's reach and are
// counted as such rather than assumed correct.
var enumDeclRe = regexp.MustCompile(`__enum_decl\(\s*([A-Za-z_][A-Za-z0-9_]*)\s*,\s*([A-Za-z_][A-Za-z0-9_]*)\s*,\s*\{([^}]*)\}`)

// cWidthToGo maps an SDK backing type to the Go type a correct rendering uses.
var cWidthToGo = map[string]string{
	"uint8_t": "uint8", "uint16_t": "uint16", "uint32_t": "uint32", "uint64_t": "uint64",
	"int8_t": "int8", "int16_t": "int16", "int32_t": "int32", "int64_t": "int64",
}

// TestEnumWidthsMatchTheSDK fails while any generated enum whose constants can
// be matched to an __enum_decl renders a width the SDK contradicts.
//
// The Mach range enums are the fail-closed canary. Their flags and tags share a
// case prefix but have different SDK widths; merging them into one synthetic Go
// enum makes at least one set of constants fail this check.
func TestEnumWidthsMatchTheSDK(t *testing.T) {
	sdk := sdkPath(t)
	root := repoRoot(t)

	// constant name -> declared Go width, from C.
	wantWidth := map[string]string{}
	wantOwner := map[string]string{}
	var declsSeen, declsUnmappableWidth int
	for _, hdr := range sdkHeaders(t, sdk) {
		data, err := os.ReadFile(hdr)
		if err != nil {
			continue
		}
		for _, m := range enumDeclRe.FindAllStringSubmatch(string(data), -1) {
			declsSeen++
			goWidth, ok := cWidthToGo[m[2]]
			if !ok {
				// A backing type this gate cannot translate is NOT a pass.
				declsUnmappableWidth++
				continue
			}
			for _, c := range strings.Split(m[3], ",") {
				c = strings.TrimSpace(c)
				if i := strings.IndexAny(c, " =\t\n"); i >= 0 {
					c = c[:i]
				}
				if c == "" || strings.HasPrefix(c, "/") {
					continue
				}
				wantWidth[c] = goWidth
				wantOwner[c] = m[1]
			}
		}
	}
	if declsSeen == 0 {
		t.Fatal("no __enum_decl found in the SDK: the gate has no C side and would pass over nothing")
	}

	// Go side: constant name -> (type name, underlying width).
	underlying := map[string]string{}
	constType := map[string]string{}
	constFile := map[string]string{}
	fset := token.NewFileSet()
	for _, path := range generatedFiles(t, root) {
		f, err := parser.ParseFile(fset, path, nil, 0)
		if err != nil {
			t.Fatalf("parse %s: %v", rel(root, path), err)
		}
		for _, decl := range f.Decls {
			gd, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}
			for _, spec := range gd.Specs {
				switch s := spec.(type) {
				case *ast.TypeSpec:
					if id, ok := s.Type.(*ast.Ident); ok {
						underlying[s.Name.Name] = id.Name
					}
				case *ast.ValueSpec:
					id, ok := s.Type.(*ast.Ident)
					if !ok {
						continue
					}
					for _, n := range s.Names {
						constType[n.Name] = id.Name
						constFile[n.Name] = rel(root, path)
					}
				}
			}
		}
	}
	if len(constType) == 0 {
		t.Fatal("no typed constants found in the tree: the Go side is empty")
	}

	type mismatch struct{ constName, goType, got, want, owner, file string }
	var bad []mismatch
	var matched, unmatched int
	for c, want := range wantWidth {
		goType, ok := constType[c]
		if !ok {
			// Present in C, absent from the tree. Counted, never folded into
			// the pass: an unchecked enum is not a correct one.
			unmatched++
			continue
		}
		matched++
		got := underlying[goType]
		if got != want {
			bad = append(bad, mismatch{c, goType, got, want, wantOwner[c], constFile[c]})
		}
	}

	if matched == 0 {
		t.Fatal("no C enum constant matched the generated tree; the join is broken, not the tree clean")
	}
	sort.Slice(bad, func(i, j int) bool { return bad[i].constName < bad[j].constName })

	for _, b := range bad {
		t.Errorf("%s: Go %s is %s, SDK declares %s as %s (%s)",
			b.constName, b.goType, b.got, b.owner, b.want, b.file)
	}
	t.Logf("SDK __enum_decl seen=%d (width untranslatable=%d); constants matched=%d unmatched=%d; mismatched=%d",
		declsSeen, declsUnmappableWidth, matched, unmatched, len(bad))
}

func sdkPath(t *testing.T) string {
	t.Helper()
	out, err := exec.Command("xcrun", "--show-sdk-path").Output()
	if err != nil {
		t.Skipf("no SDK available: %v", err)
	}
	return strings.TrimSpace(string(out))
}

// sdkHeaders returns the mach and kernel headers that carry __enum_decl. Scoped
// rather than whole-SDK so the gate stays fast; widening it can only find more.
func sdkHeaders(t *testing.T, sdk string) []string {
	t.Helper()
	var out []string
	for _, dir := range []string{"usr/include/mach", "usr/include/sys", "usr/include"} {
		matches, err := filepath.Glob(filepath.Join(sdk, dir, "*.h"))
		if err != nil {
			t.Fatal(err)
		}
		out = append(out, matches...)
	}
	return out
}
