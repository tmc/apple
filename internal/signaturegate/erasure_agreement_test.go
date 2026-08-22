// Copyright 2026 The tmc/apple Authors. All rights reserved.

// Package signaturegate holds correctness gates that read the generated
// source rather than importing it.
//
// A reflect-based test can only see the types its package imports, and it can
// only enumerate a set someone wrote down by hand. Both limits matter here: the
// defects these gates pin are tree-wide, and a hand-listed sample cannot
// distinguish "the generator is correct" from "the sample missed it". Parsing
// the .gen.go files reaches every package and needs no import graph.
package signaturegate

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"sort"
	"strings"
	"testing"
)

// erasedContainerRenderings are the two Go types one erased ObjC container can
// render as. NSArray<id<Proto>> * becomes []objectivec.IObject and bare NSArray
// * becomes foundation.INSArray; both are faithful to their own declaration.
//
// The pair is what makes this gate self-deriving. It asserts no belief about
// which rendering is right for any given selector -- only that ONE selector
// cannot be both, because Go has no erasure and a caller holding the class
// cannot pass what a caller holding the protocol wrapper must pass.
var erasedContainerRenderings = [2]string{"[]objectivec.IObject", "foundation.INSArray"}

// TestErasedContainerRenderingsAgree fails while any method name is rendered
// both ways across class surfaces and protocol wrappers reachable from a
// generated class conformance. A standalone protocol wrapper with no
// conforming class keeps the rendering faithful to its own declaration.
//
// The comparison is scoped to a substitutability group (see
// [substitutabilityGroups]) rather than to a bare method name. The reason is
// the one this gate can state from its own side: a disagreement is only a
// defect where a caller could hold either receiver, which is a class, its own
// interface, and the protocol wrappers it conforms to. Two unrelated classes
// in unrelated frameworks may share a method name while each faithfully renders
// its own declaration. That is agreement with the SDK, not disagreement with
// each other, and no reconciliation could fix it without discarding SDK
// information from whichever side got weakened.
//
// This comment used to describe the generator's collision pre-pass instead --
// which fields it keys on, and which selectors it therefore misses. That
// description was wrong, and the failure is structural rather than careless:
// prose in this package describing another repository's internals is
// unversioned duplication. Nothing breaks when the generator changes, no test
// covers the sentence, and sitting inside a gate lends it the authority of
// something verified. It cost a peer session a full investigation into a
// generator that was behaving correctly. The rule this file now follows is to
// state only what it can check from here; the pre-pass is real and lives in
// appledocs internal/generator/erasure_collisions.go, and that pointer is
// deliberately the whole of what is said about it.
//
// The gate fails loudly with the full reachable population rather than with
// the first example, so a new collision cannot hide behind a known one.
func TestErasedContainerRenderingsAgree(t *testing.T) {
	root := repoRoot(t)

	// method name -> rendering -> one example site, for the failure message.
	seen := map[string]map[string]string{}
	// Protocol wrappers with no generated class conformance are faithful to
	// their own declaration and outside the comparison population. Count both
	// the receiver types and method names removed so a green result cannot hide
	// a population change.
	skippedReceivers := map[string]bool{}
	exemptedMethods := map[string]bool{}

	files := generatedFiles(t, root)
	if len(files) == 0 {
		t.Fatal("no .gen.go files found: this gate would pass over an empty tree")
	}

	fset := token.NewFileSet()
	parsed := make([]parsedGeneratedFile, 0, len(files))
	for _, path := range files {
		f, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
		if err != nil {
			// A parse failure is not an absent disagreement. Folding the two
			// together would let a broken file read as a clean one.
			t.Fatalf("parse %s: %v", rel(root, path), err)
		}
		parsed = append(parsed, parsedGeneratedFile{path: path, file: f})
	}
	reachable, protocols := reachableProtocolObjects(parsed)
	groups := substitutabilityGroups(parsed)
	// The population under the OLD key -- bare method name, tree-wide -- kept
	// so the narrowing reports both numbers. A green run whose population
	// silently shrank is not a fix, so the size of the change has to travel
	// with the verdict rather than be recoverable only by rerunning an older
	// revision.
	byBareName := map[string]bool{}
	for _, parsed := range parsed {
		rel := rel(root, parsed.path)
		for _, decl := range parsed.file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Recv == nil || fn.Type.Params == nil {
				continue
			}
			renderings := erasedContainerParameterRenderings(fn)
			if len(renderings) == 0 {
				continue
			}
			receiver := receiverName(fn)
			if !scanErasedContainerReceiver(receiver, reachable, protocols) {
				skippedReceivers[receiver] = true
				exemptedMethods[fn.Name.Name] = true
				continue
			}
			// Compare within a substitutability group, not across the tree.
			// The bare method name was the old key, and it made every
			// same-named method in every framework one population.
			group := groups[qualifiedReceiver(parsed.path, receiver)]
			for _, got := range renderings {
				name := group + " " + fn.Name.Name
				byBareName[fn.Name.Name] = true
				if seen[name] == nil {
					seen[name] = map[string]string{}
				}
				if _, dup := seen[name][got]; !dup {
					seen[name][got] = rel + ": " + receiver + "." + fn.Name.Name
				}
			}
		}
	}

	var disagreeing []string
	for name, renderings := range seen {
		if len(renderings) > 1 {
			disagreeing = append(disagreeing, name)
		}
	}
	sort.Strings(disagreeing)

	// A gate that examined nothing passes. Both renderings occur in the tree by
	// the hundreds, so finding neither means the scan broke rather than that
	// the tree is clean.
	if len(seen) == 0 {
		t.Fatal("no method renders either erased container type; the scan is not reading the tree")
	}

	if len(disagreeing) > 0 {
		t.Errorf("%d substitutable method(s) rendered as BOTH %s and %s; one selector cannot be both:",
			len(disagreeing), erasedContainerRenderings[0], erasedContainerRenderings[1])
		for _, name := range disagreeing {
			sites := seen[name]
			t.Errorf("  %s:\n      %s\n      %s", name,
				sites[erasedContainerRenderings[0]], sites[erasedContainerRenderings[1]])
		}
	}
	exemptedOnly := 0
	for name := range exemptedMethods {
		if seen[name] == nil {
			exemptedOnly++
		}
	}
	t.Logf("scanned %d file(s); %d unreachable receiver(s) skipped (%d method name(s), %d absent from compared population); %d bare method name(s) render an erased container, which the substitutability key resolves to %d compared method(s); %d disagree",
		len(files), len(skippedReceivers), len(exemptedMethods), exemptedOnly,
		len(byBareName), len(seen), len(disagreeing))
}

// qualifiedReceiver names a concrete receiver uniquely across the tree. Two
// frameworks may declare the same Go type name, and the old bare-name key
// merged them.
func qualifiedReceiver(path, receiver string) string {
	return filepath.Dir(path) + "." + receiver
}

// substitutabilityGroups maps each qualified concrete receiver to an identifier
// it shares with every receiver a caller could hold in its place: the class,
// and the protocol wrapper object of each protocol the class conforms to. The
// generator marks conformances with a "Protocol methods for P" heading in the
// class file, which is the same marker [reachableProtocolObjects] reads.
//
// A receiver in no group is its own group. That is the conservative direction:
// it compares a receiver only against itself, so an unrelated pair is never
// reported, and a genuinely related pair this function fails to link is a
// missed detection rather than a false one. The gate says which way it errs
// because the two are not equally recoverable -- a false report costs an
// investigation, as this gate has already demonstrated once.
func substitutabilityGroups(files []parsedGeneratedFile) map[string]string {
	// Where each protocol wrapper object type is declared, by bare name. A
	// class may conform to a protocol declared in another framework, and the
	// heading carries only the bare name, so the link is resolved by search.
	declared := map[string][]string{}
	for _, parsed := range files {
		dir := filepath.Dir(parsed.path)
		for _, decl := range parsed.file.Decls {
			gen, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}
			for _, spec := range gen.Specs {
				if typeSpec, ok := spec.(*ast.TypeSpec); ok {
					declared[typeSpec.Name.Name] = append(declared[typeSpec.Name.Name], dir+"."+typeSpec.Name.Name)
				}
			}
		}
	}

	u := newUnionFind()
	for _, parsed := range files {
		dir := filepath.Dir(parsed.path)
		var receivers []string
		for _, decl := range parsed.file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Recv == nil {
				continue
			}
			receivers = append(receivers, dir+"."+receiverName(fn))
		}
		if len(receivers) == 0 {
			continue
		}
		for _, group := range parsed.file.Comments {
			for _, line := range strings.Split(group.Text(), "\n") {
				const prefix = "Protocol methods for "
				if !strings.HasPrefix(line, prefix) {
					continue
				}
				object := strings.TrimSpace(strings.TrimPrefix(line, prefix)) + "Object"
				for _, wrapper := range declared[object] {
					for _, receiver := range receivers {
						u.union(receiver, wrapper)
					}
				}
			}
		}
	}

	groups := make(map[string]string, len(declared))
	for _, parsed := range files {
		dir := filepath.Dir(parsed.path)
		for _, decl := range parsed.file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Recv == nil {
				continue
			}
			key := dir + "." + receiverName(fn)
			groups[key] = u.find(key)
		}
	}
	return groups
}

type unionFind struct{ parent map[string]string }

func newUnionFind() *unionFind { return &unionFind{parent: map[string]string{}} }

func (u *unionFind) find(x string) string {
	root, ok := u.parent[x]
	if !ok {
		u.parent[x] = x
		return x
	}
	for root != u.parent[root] {
		u.parent[root] = u.parent[u.parent[root]]
		root = u.parent[root]
	}
	u.parent[x] = root
	return root
}

func (u *unionFind) union(a, b string) {
	ra, rb := u.find(a), u.find(b)
	if ra == rb {
		return
	}
	// Order by name so the group identifier does not depend on file walk
	// order; an identifier that moved between runs would make the failure
	// message unstable for no benefit.
	if ra > rb {
		ra, rb = rb, ra
	}
	u.parent[rb] = ra
}

type parsedGeneratedFile struct {
	path string
	file *ast.File
}

// reachableProtocolObjects reports protocol wrappers backed by at least one
// generated class conformance. A protocol inherited by a reachable protocol is
// reachable too. The generator marks class conformances with a stable
// "Protocol methods for" heading in the generated class file.
func reachableProtocolObjects(files []parsedGeneratedFile) (map[string]bool, map[string]bool) {
	parents := make(map[string][]string)
	reachable := make(map[string]bool)
	for _, parsed := range files {
		for _, decl := range parsed.file.Decls {
			gen, ok := decl.(*ast.GenDecl)
			if !ok {
				continue
			}
			for _, spec := range gen.Specs {
				typeSpec, ok := spec.(*ast.TypeSpec)
				if !ok {
					continue
				}
				iface, ok := typeSpec.Type.(*ast.InterfaceType)
				if !ok || iface.Methods == nil {
					continue
				}
				parents[typeSpec.Name.Name] = nil
				for _, field := range iface.Methods.List {
					if len(field.Names) != 0 {
						continue
					}
					if name, ok := field.Type.(*ast.Ident); ok {
						parents[typeSpec.Name.Name] = append(parents[typeSpec.Name.Name], name.Name)
					}
				}
			}
		}
		for _, group := range parsed.file.Comments {
			for _, line := range strings.Split(group.Text(), "\n") {
				const prefix = "Protocol methods for "
				if strings.HasPrefix(line, prefix) {
					reachable[strings.TrimSpace(strings.TrimPrefix(line, prefix))] = true
				}
			}
		}
	}

	for changed := true; changed; {
		changed = false
		for protocol := range reachable {
			for _, parent := range parents[protocol] {
				if _, ok := parents[parent]; ok && !reachable[parent] {
					reachable[parent] = true
					changed = true
				}
			}
		}
	}
	protocols := make(map[string]bool, len(parents))
	for protocol := range parents {
		protocols[protocol] = true
	}
	return reachable, protocols
}

func scanErasedContainerReceiver(receiver string, reachable, protocols map[string]bool) bool {
	if !strings.HasSuffix(receiver, "Object") {
		return true
	}
	protocol := strings.TrimSuffix(receiver, "Object")
	return !protocols[protocol] || reachable[protocol]
}

func erasedContainerParameterRenderings(fn *ast.FuncDecl) []string {
	var renderings []string
	for _, param := range fn.Type.Params.List {
		// Open Directory's authentication methods use []objectivec.IObject
		// for the outContext out-parameter. That is a Go slice used as an
		// Objective-C out buffer, not an erased NSArray parameter, so it must
		// not be compared with foundation.INSArray.
		if len(param.Names) > 0 && param.Names[0].Name == "outContext" {
			continue
		}
		got := exprString(param.Type)
		if got == erasedContainerRenderings[0] || got == erasedContainerRenderings[1] {
			renderings = append(renderings, got)
		}
	}
	return renderings
}

// TestErasureAgreementGateCanFail proves the comparison distinguishes, so that
// a green run above is evidence rather than an artifact of a scan that matches
// nothing. Without this, deleting the body of the gate would also make it pass.
func TestErasureAgreementGateCanFail(t *testing.T) {
	const src = `package p
func (a A) SetRows(rows []objectivec.IObject) {}
func (b B) SetRows(rows foundation.INSArray) {}
func (c C) SetCols(cols foundation.INSArray) {}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", src, 0)
	if err != nil {
		t.Fatal(err)
	}
	seen := map[string]map[string]string{}
	for _, decl := range f.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Recv == nil || fn.Type.Params == nil {
			continue
		}
		for _, got := range erasedContainerParameterRenderings(fn) {
			if seen[fn.Name.Name] == nil {
				seen[fn.Name.Name] = map[string]string{}
			}
			seen[fn.Name.Name][got] = "x"
		}
	}
	if len(seen["SetRows"]) != 2 {
		t.Errorf("SetRows: got %d rendering(s), want 2; the gate would not detect a real disagreement", len(seen["SetRows"]))
	}
	if len(seen["SetCols"]) != 1 {
		t.Errorf("SetCols: got %d rendering(s), want 1; the gate reports agreement as disagreement", len(seen["SetCols"]))
	}
}

// TestSubstitutabilityGroupsSeparateUnrelatedReceivers proves the narrowing in
// BOTH directions. A control that only showed the gate still fires on a related
// pair would not distinguish this key from the old bare-name key, and a control
// that only showed it stays quiet on an unrelated pair is satisfied by a gate
// that examines nothing.
//
// The unrelated arm is the SetOutputChannels shape: avfaudio's
// AVSpeechSynthesizer takes NSArray<AVAudioSessionChannelDescription *> * and
// private/texttospeech's TTSSpeechManager takes bare NSArray *. Both render
// their own declaration correctly, and the old key called that a defect.
func TestSubstitutabilityGroupsSeparateUnrelatedReceivers(t *testing.T) {
	const conformer = `package p
type FooProtoObject struct{}
func (o FooProtoObject) SetRows(rows foundation.INSArray) {}
// Protocol methods for FooProto
func (c Class) SetRows(rows []objectivec.IObject) {}
`
	const unrelated = `package q
func (d Distant) SetRows(rows foundation.INSArray) {}
`
	fset := token.NewFileSet()
	var parsed []parsedGeneratedFile
	for _, f := range []struct{ path, src string }{
		{"avfaudio/a.gen.go", conformer},
		{"texttospeech/b.gen.go", unrelated},
	} {
		file, err := parser.ParseFile(fset, f.path, f.src, parser.ParseComments)
		if err != nil {
			t.Fatal(err)
		}
		parsed = append(parsed, parsedGeneratedFile{path: f.path, file: file})
	}
	groups := substitutabilityGroups(parsed)

	class := groups[qualifiedReceiver("avfaudio/a.gen.go", "Class")]
	wrapper := groups[qualifiedReceiver("avfaudio/a.gen.go", "FooProtoObject")]
	distant := groups[qualifiedReceiver("texttospeech/b.gen.go", "Distant")]

	if class == "" || wrapper == "" || distant == "" {
		t.Fatalf("a receiver got no group at all (class=%q wrapper=%q distant=%q); "+
			"the grouping did not read the tree", class, wrapper, distant)
	}
	if class != wrapper {
		t.Errorf("Class and the wrapper of a protocol it conforms to are in different groups "+
			"(%q vs %q); a real erasure disagreement would go undetected", class, wrapper)
	}
	if class == distant {
		t.Errorf("an unrelated receiver in another framework shares Class's group (%q); "+
			"this is the SetOutputChannels false positive the narrowing removes", class)
	}
}

func TestErasureAgreementGateSkipsUnreachableProtocolObjects(t *testing.T) {
	const src = `package p
type Parent interface{}
type Child interface{ Parent }
type Orphan interface{}
// Protocol methods for Child
type Class struct{}
`
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", src, parser.ParseComments)
	if err != nil {
		t.Fatal(err)
	}
	reachable, protocols := reachableProtocolObjects([]parsedGeneratedFile{{path: "x.go", file: f}})
	for _, receiver := range []string{"ChildObject", "ParentObject", "Class"} {
		if !scanErasedContainerReceiver(receiver, reachable, protocols) {
			t.Errorf("%s was skipped, want reachable", receiver)
		}
	}
	if scanErasedContainerReceiver("OrphanObject", reachable, protocols) {
		t.Error("OrphanObject was scanned, want unreachable protocol wrapper skipped")
	}
}

func receiverName(fn *ast.FuncDecl) string {
	if fn.Recv == nil || len(fn.Recv.List) == 0 {
		return ""
	}
	return exprString(fn.Recv.List[0].Type)
}

// exprString renders the type expressions this gate compares. It handles only
// the forms the two erased-container renderings take; anything else returns a
// spelling that matches neither, which is the correct outcome for a type this
// gate says nothing about.
func exprString(e ast.Expr) string {
	switch t := e.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.SelectorExpr:
		return exprString(t.X) + "." + t.Sel.Name
	case *ast.ArrayType:
		if t.Len == nil {
			return "[]" + exprString(t.Elt)
		}
		return "[N]" + exprString(t.Elt)
	case *ast.StarExpr:
		return "*" + exprString(t.X)
	}
	return "?"
}
