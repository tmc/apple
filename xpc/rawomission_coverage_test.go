package xpc

import (
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"regexp"
	"sort"
	"strings"
	"testing"
)

var rawRegistration = regexp.MustCompile(`registerRawFunc\(&rawfn_(xpc_[a-z0-9_]+), frameworkHandle, "xpc_[a-z0-9_]+"\)`)

// boundRawSymbols reports the symbols registered by the generated raw layer.
// Registration, rather than declaration, is the relevant boundary: these are
// the functions the package can call at run time.
func boundRawSymbols(t *testing.T) []string {
	t.Helper()
	b, err := os.ReadFile("xpc.raw.gen.go")
	if err != nil {
		t.Fatalf("read raw bindings: %v", err)
	}
	matches := rawRegistration.FindAllSubmatch(b, -1)
	if len(matches) == 0 {
		t.Fatal("found zero raw registrations: the accounting gate is not reading generated bindings")
	}
	syms := make([]string, 0, len(matches))
	seen := make(map[string]bool, len(matches))
	for _, match := range matches {
		sym := string(match[1])
		if seen[sym] {
			t.Fatalf("duplicate raw registration for %s", sym)
		}
		seen[sym] = true
		syms = append(syms, sym)
	}
	sort.Strings(syms)
	return syms
}

// usedRawSymbols parses production source, rather than searching text, so a
// comment or test cannot accidentally account for an otherwise raw-only API.
func usedRawSymbols(t *testing.T) map[string]bool {
	t.Helper()
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatalf("read package directory: %v", err)
	}
	used := make(map[string]bool)
	fset := token.NewFileSet()
	files := 0
	for _, entry := range entries {
		name := entry.Name()
		if entry.IsDir() || !strings.HasSuffix(name, ".go") || strings.HasSuffix(name, "_test.go") || name == "xpc.raw.gen.go" {
			continue
		}
		f, err := parser.ParseFile(fset, name, nil, 0)
		if err != nil {
			t.Fatalf("parse %s: %v", name, err)
		}
		files++
		ast.Inspect(f, func(n ast.Node) bool {
			id, ok := n.(*ast.Ident)
			if ok && strings.HasPrefix(id.Name, "raw_xpc_") {
				used[strings.TrimPrefix(id.Name, "raw_")] = true
			}
			return true
		})
	}
	if files == 0 {
		t.Fatal("scanned zero production source files")
	}
	return used
}

func rawSymbolIsOmitted(t *testing.T, sym string, omissions []rawSymbolOmission) bool {
	t.Helper()
	for _, omission := range omissions {
		for _, pattern := range omission.Symbols {
			re, err := regexp.Compile("^(?:" + pattern + ")$")
			if err != nil {
				t.Fatalf("compile omission pattern %q: %v", pattern, err)
			}
			if re.MatchString(sym) {
				return true
			}
		}
	}
	return false
}

func unaccountedRawSymbols(t *testing.T, bound []string, used map[string]bool, omissions []rawSymbolOmission) []string {
	t.Helper()
	var missing []string
	for _, sym := range bound {
		if !used[sym] && !rawSymbolIsOmitted(t, sym, omissions) {
			missing = append(missing, sym)
		}
	}
	return missing
}

// TestBoundRawSymbolsAreAccountedFor prevents regeneration from silently
// adding a raw-only entry point. Every bound symbol must have production use
// or an explicit, reasoned omission in the generated ledger.
func TestBoundRawSymbolsAreAccountedFor(t *testing.T) {
	missing := unaccountedRawSymbols(t, boundRawSymbols(t), usedRawSymbols(t), xpcRawOmissions)
	if len(missing) != 0 {
		t.Fatalf("bound raw symbols are neither used nor ledgered: %s", strings.Join(missing, ", "))
	}
}

// TestBoundRawSymbolsLedgerSensitivity proves the gate is fail-closed for a
// deleted ledger entry. xpc_main has no production use by design, so removing
// its omission from an in-memory copy must make it unaccounted.
func TestBoundRawSymbolsLedgerSensitivity(t *testing.T) {
	withoutMain := make([]rawSymbolOmission, 0, len(xpcRawOmissions))
	removed := false
	for _, omission := range xpcRawOmissions {
		if len(omission.Symbols) == 1 && omission.Symbols[0] == "xpc_main" {
			removed = true
			continue
		}
		withoutMain = append(withoutMain, omission)
	}
	if !removed {
		t.Fatal("xpc_main omission is absent: sensitivity control has no mutation")
	}
	missing := unaccountedRawSymbols(t, boundRawSymbols(t), usedRawSymbols(t), withoutMain)
	for _, sym := range missing {
		if sym == "xpc_main" {
			return
		}
	}
	t.Fatalf("removing xpc_main from the omission ledger did not make it unaccounted: %v", missing)
}
