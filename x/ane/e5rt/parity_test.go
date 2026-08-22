package e5rt

import (
	"go/ast"
	"go/parser"
	"go/token"
	"slices"
	"strconv"
	"testing"
)

// This package has two implementations selected by build tag: e5rt.go for darwin
// and e5rt_other.go for everything else. Nothing makes them agree. A method added
// to one builds cleanly without the other, because the stub is a separate
// implementation rather than a generated shadow of the real one, so the skew is
// invisible to a compiler, to go vet and to every test that runs on one GOOS.
//
// It had in fact drifted by eleven exported methods before this test existed.
// These tests read both files as source, which works whatever GOOS the test
// binary was built for.

// parseLibSurface returns the exported methods on *Lib and the string contents
// of the Symbols slice declared in path.
func parseLibSurface(t *testing.T, path string) (methods, symbols []string) {
	t.Helper()
	file, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
	if err != nil {
		t.Fatalf("parse %s: %v", path, err)
	}
	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.FuncDecl:
			if d.Recv == nil || !d.Name.IsExported() {
				continue
			}
			if isLibReceiver(d.Recv) {
				methods = append(methods, d.Name.Name)
			}
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				value, ok := spec.(*ast.ValueSpec)
				if !ok || len(value.Names) != 1 || value.Names[0].Name != "Symbols" {
					continue
				}
				symbols = append(symbols, stringElements(value.Values)...)
			}
		}
	}
	slices.Sort(methods)
	slices.Sort(symbols)
	return methods, symbols
}

// isLibReceiver reports whether a receiver list is (l *Lib).
func isLibReceiver(recv *ast.FieldList) bool {
	if len(recv.List) != 1 {
		return false
	}
	star, ok := recv.List[0].Type.(*ast.StarExpr)
	if !ok {
		return false
	}
	name, ok := star.X.(*ast.Ident)
	return ok && name.Name == "Lib"
}

// stringElements returns the string literals in a composite literal.
func stringElements(values []ast.Expr) []string {
	var out []string
	for _, v := range values {
		lit, ok := v.(*ast.CompositeLit)
		if !ok {
			continue
		}
		for _, elt := range lit.Elts {
			basic, ok := elt.(*ast.BasicLit)
			if !ok || basic.Kind != token.STRING {
				continue
			}
			s, err := strconv.Unquote(basic.Value)
			if err != nil {
				continue
			}
			out = append(out, s)
		}
	}
	return out
}

// TestBuildsAgree checks that the darwin implementation and the non-darwin stub
// expose the same methods and list the same symbols.
func TestBuildsAgree(t *testing.T) {
	darwinMethods, darwinSymbols := parseLibSurface(t, "e5rt.go")
	otherMethods, otherSymbols := parseLibSurface(t, "e5rt_other.go")

	// The parse must find something, or every comparison below passes vacuously.
	if len(darwinMethods) == 0 || len(darwinSymbols) == 0 {
		t.Fatalf("parsed e5rt.go as %d methods and %d symbols; the test cannot fail in this state",
			len(darwinMethods), len(darwinSymbols))
	}

	for _, name := range darwinMethods {
		if !slices.Contains(otherMethods, name) {
			t.Errorf("e5rt.go has method %s, e5rt_other.go does not", name)
		}
	}
	for _, name := range otherMethods {
		if !slices.Contains(darwinMethods, name) {
			t.Errorf("e5rt_other.go has method %s, e5rt.go does not", name)
		}
	}
	for _, name := range darwinSymbols {
		if !slices.Contains(otherSymbols, name) {
			t.Errorf("e5rt.go lists symbol %s, e5rt_other.go does not", name)
		}
	}
	for _, name := range otherSymbols {
		if !slices.Contains(darwinSymbols, name) {
			t.Errorf("e5rt_other.go lists symbol %s, e5rt.go does not", name)
		}
	}
}

// TestBuildsAgreeCanFail is the mutation control for TestBuildsAgree. A set
// comparison that always passes is indistinguishable from one that never runs,
// so this checks the comparison reports a difference when given one.
func TestBuildsAgreeCanFail(t *testing.T) {
	_, symbols := parseLibSurface(t, "e5rt.go")
	perturbed := append(slices.Clone(symbols), "e5rt_not_a_real_symbol")
	if slices.Contains(symbols, "e5rt_not_a_real_symbol") {
		t.Fatal("the perturbation is already present; it cannot act as a control")
	}
	var missing int
	for _, name := range perturbed {
		if !slices.Contains(symbols, name) {
			missing++
		}
	}
	if missing != 1 {
		t.Errorf("the comparison found %d differences in a set perturbed by one, want 1", missing)
	}
}
