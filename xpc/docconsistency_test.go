// Copyright 2026 The apple Authors.

package xpc

// The package-doc consistency gate.
//
// doc.gen.go documented Codec, DefaultCodec, SendDictionary and
// ReplyDictionary long after all four were deleted, because the doc is rendered
// from a different source (appledocs internal/generator/frameworks/docgen/
// policy.go) than the API it describes, and four editing passes over the API
// never re-rendered it. Nothing in the tree compared the two. A doc naming a
// symbol that does not exist is not cosmetic: it is the first thing a caller
// reads, and `go doc` prints it beside a surface that contradicts it.
//
// This test reads the package doc comment out of the source, extracts every
// identifier-shaped word, and requires each to be either an exported identifier
// of this package or a listed prose word. Prose words are listed rather than
// guessed so that new prose forces a deliberate decision, and the list itself is
// checked: a word may not be excused as prose if it is also an exported
// identifier, which would let a real symbol slip past unchecked.

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

// docProseWords are capitalized words in the package doc that are English prose
// or external names, not identifiers of this package.
var docProseWords = map[string]bool{
	"API":     true,
	"Apple":   true,
	"C":       true, // "no safe public C path"
	"Code":    true, // "Code generated ... DO NOT EDIT."
	"DO":      true,
	"EDIT":    true,
	"Go":      true,
	"High":    true, // "High-level"
	"NOT":     true,
	"Notes":   true,
	"Package": true,
	"Swift":   true,
	"The":     true,
	"Typed":   true,
	"XPC":     true,
}

// docIdentRE matches an identifier-shaped word: a capitalized run of letters and
// digits. Trailing punctuation and the "xpc.omissions.gen.go" style of filename
// are stripped before matching.
var docIdentRE = regexp.MustCompile(`\b[A-Z][A-Za-z0-9]*\b`)

// TestPackageDocNamesOnlyRealIdentifiers is Gate B.
func TestPackageDocNamesOnlyRealIdentifiers(t *testing.T) {
	doc, exported := packageDocAndSurface(t)
	if doc == "" {
		t.Fatalf("no package doc comment found; the gate would pass by having nothing to check")
	}
	if len(exported) == 0 {
		t.Fatalf("no exported identifiers found; the scanner is broken, so no verdict below is about the tree")
	}

	// The prose list may not shadow a real identifier, and every entry must
	// actually appear in the doc — a stale exemption is an unexamined claim.
	for word := range docProseWords {
		if exported[word] {
			t.Errorf("docProseWords excuses %q, which is an exported identifier: the exemption would hide a real check", word)
		}
	}

	findings, checked := docFindings(doc, exported)
	for _, word := range findings {
		t.Errorf("the package doc names %q, which is not in the exported surface of package xpc "+
			"(nor a listed prose word): the doc and the API disagree. The doc is rendered from "+
			"appledocs internal/generator/frameworks/docgen/policy.go, not from doc.gen.go.", word)
	}
	if checked == 0 {
		t.Fatalf("the doc contained no identifier-shaped words to check; docProseWords is swallowing everything")
	}
	t.Logf("checked %d identifier-shaped words in the package doc against %d exported names", checked, len(exported))
}

// TestPackageDocGateIsSensitive is the mutation control: it runs the same
// predicate over the real doc with one bogus identifier appended, and fails if
// the check reports success. Without it a gate whose candidate extraction had
// stopped matching anything would stay green forever.
func TestPackageDocGateIsSensitive(t *testing.T) {
	doc, exported := packageDocAndSurface(t)
	base, baseTotal := docFindings(doc, exported)
	mutated, mutTotal := docFindings(doc+"\n\nTyped payloads are supported through NoSuchIdentifierXYZ.\n", exported)
	if mutTotal != baseTotal+1 {
		t.Fatalf("mutation control failed: injecting one identifier changed the examined population from %d to %d, want %d: "+
			"the candidate extraction does not see the injected word at all", baseTotal, mutTotal, baseTotal+1)
	}
	if len(mutated) != len(base)+1 {
		t.Fatalf("mutation control failed: injecting one bogus identifier changed the finding count from %d to %d, want %d; "+
			"the gate cannot see the defect it exists to see", len(base), len(mutated), len(base)+1)
	}
	t.Logf("mutation control: the gate flags exactly the injected identifier (examined %d words, %d pre-existing findings)",
		mutTotal, len(base))
}

// docFindings returns the identifier-shaped words in doc that are neither
// exported nor listed prose, and the number of words examined. Both gates run
// through it so the mutation control measures the production predicate.
func docFindings(doc string, exported map[string]bool) (findings []string, examined int) {
	for _, word := range docIdentCandidates(doc) {
		if docProseWords[word] {
			continue
		}
		examined++
		if !exported[word] {
			findings = append(findings, word)
		}
	}
	return findings, examined
}

// docIdentCandidates extracts identifier-shaped words from a doc comment,
// skipping code-span filenames like xpc.omissions.gen.go and lowercase words.
func docIdentCandidates(doc string) []string {
	seen := map[string]bool{}
	var out []string
	for _, line := range strings.Split(doc, "\n") {
		for _, m := range docIdentRE.FindAllString(line, -1) {
			if seen[m] {
				continue
			}
			seen[m] = true
			out = append(out, m)
		}
	}
	sort.Strings(out)
	return out
}

// packageDocAndSurface returns the package doc comment text and the set of
// exported names of package xpc, including method names both bare and qualified
// by receiver (Decode and ReceivedMessage.Decode both name a real thing).
func packageDocAndSurface(t *testing.T) (string, map[string]bool) {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, wd, func(fi os.FileInfo) bool {
		return !strings.HasSuffix(fi.Name(), "_test.go")
	}, parser.ParseComments)
	if err != nil {
		t.Fatalf("parsing %s: %v", wd, err)
	}
	pkg := pkgs["xpc"]
	if pkg == nil {
		t.Fatalf("package xpc not found in %s", wd)
	}

	doc := ""
	exported := map[string]bool{}
	for _, file := range pkg.Files {
		// Every file's package comment is checked, not just doc.gen.go's:
		// go/doc concatenates them, so all of them are part of what a reader
		// sees, and a stale name in any of them is the same defect.
		if file.Doc != nil && file.Doc.Text() != "" {
			doc += "\n" + file.Doc.Text()
		}
		for _, decl := range file.Decls {
			switch d := decl.(type) {
			case *ast.FuncDecl:
				if !d.Name.IsExported() {
					continue
				}
				exported[d.Name.Name] = true
				if d.Recv != nil && len(d.Recv.List) > 0 {
					if recv := receiverTypeName(d.Recv.List[0].Type); recv != "" {
						exported[recv+"."+d.Name.Name] = true
					}
				}
			case *ast.GenDecl:
				for _, spec := range d.Specs {
					switch s := spec.(type) {
					case *ast.TypeSpec:
						if s.Name.IsExported() {
							exported[s.Name.Name] = true
							collectExportedFields(s, exported)
						}
					case *ast.ValueSpec:
						for _, id := range s.Names {
							if id.IsExported() {
								exported[id.Name] = true
							}
						}
					}
				}
			}
		}
	}
	return doc, exported
}

// collectExportedFields records exported struct fields and interface methods, so
// a doc may name Listener.Options-style members and interface methods.
func collectExportedFields(s *ast.TypeSpec, exported map[string]bool) {
	var fields *ast.FieldList
	switch t := s.Type.(type) {
	case *ast.StructType:
		fields = t.Fields
	case *ast.InterfaceType:
		fields = t.Methods
	default:
		return
	}
	if fields == nil {
		return
	}
	for _, f := range fields.List {
		for _, id := range f.Names {
			if id.IsExported() {
				exported[id.Name] = true
				exported[s.Name.Name+"."+id.Name] = true
			}
		}
	}
}

func receiverTypeName(expr ast.Expr) string {
	if star, ok := expr.(*ast.StarExpr); ok {
		expr = star.X
	}
	if id, ok := expr.(*ast.Ident); ok {
		return id.Name
	}
	return ""
}
