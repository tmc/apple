// Copyright 2026 The tmc/apple Authors. All rights reserved.

package signaturegate

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"os/exec"
	"sort"
	"strings"
	"testing"
)

// A regeneration can replace a type with a wider or emptier one and still
// compile, still pass every test, and still look like noise in git diff --stat.
// The regen that produced kernel/types.gen.go with Timespec.Tv_sec as
// unsafe.Pointer changed 198 files; the build noticed only because one example
// assigns an integer to that one field.
//
// This gate compares the generated tree against a baseline revision and reports
// three degradations, none of which a build catches on its own:
//
//   - an integer struct field that is now unsafe.Pointer, which is an ABI
//     change and, for a 4-byte member, a size change as well;
//   - a named type that is now a builtin or a byte array, which keeps the
//     size and loses the identity;
//   - an enum constant whose value is now zero, which is what an SDK that
//     does not declare the symbol produces — two constants collapse onto one
//     value and the String method silently loses a case.
//
// The baseline is a fixed revision, and it has to be. A relative baseline —
// HEAD~1 — measures what the last commit introduced rather than what the tree
// carries, so every commit launders the previous commit's degradations into
// the baseline and the gate's answer changes meaning when unrelated work
// lands. Written against HEAD~1 this gate reported 3 findings one hour and 4
// the next, and Backtrace_control.Btc_flags disappeared from the list while
// still being degraded in the tree.
//
// Moving baselineRev is therefore a deliberate act: it declares that
// everything the tree carries up to that revision is accepted.
//
// Set APPLE_GATE_BASE to compare against something else for a one-off.

// baselineRev is the v0.7.0 tip. It answers, going forward, "what has degraded
// since v0.7.0 shipped".
//
// It was 77eeb3409, a mid-branch commit, and that is a failure mode worth
// recording rather than quietly correcting. A revision inside a branch that
// may be rewritten is a pin that dangles: the history reconstruction of
// 2026-08-22 collapsed 233 commits to 19 and left 77eeb3409 unreachable. The
// object survives locally as an unreachable commit, so rev-parse keeps
// resolving it and this gate keeps passing on the machine that did the
// rewrite -- and hard-fails on a fresh clone, or here after a gc. A gate that
// works only where it was written is worse than one that is merely absent.
//
// It also asked a question nobody wanted. Pinned mid-branch, the gate reported
// what changed since an arbitrary Thursday in August. Pinned at a release, it
// reports what changed since a release.
//
// Two things this pin still does NOT fix, both tracked for v0.7.1:
//
//   - Moving it laundered the applicationservices/QD.framework damage into the
//     baseline: CMDeviceInfo, CMMultiFunctLutType, CMDeviceProfileArray and
//     CMDeviceScope render every field as unsafe.Pointer at this revision, and
//     comparing against it can no longer see that. They are held by
//     TestNoAllOpaqueRecords instead, which takes no baseline at all.
//   - A revision pin is still a pin. The replacement is an in-tree golden
//     manifest: regenerating a file shows the damage as a reviewable diff,
//     where moving a hash is a one-line change that reveals nothing. Today the
//     most consequential act is the least visible one.
const baselineRev = "62547ebc14a6217ff3144ff94f1fefaee8cf342d"

// knownFindings are the degradations the tree carried when the baseline was
// pinned, each recorded exactly as the gate reports it — never a pattern, so
// nothing else can hide behind an entry.
//
// An entry that no longer reproduces is a failure, not a pass. An allowlist
// that only ever suppresses becomes permanent; one that goes red when a
// finding is fixed forces its own cleanup.
// It is empty at v0.7.0, and that is a consequence of moving the pin rather
// than a claim that the tree is clean. Every entry it held was a constant that
// collapsed to zero BEFORE this revision, so against this baseline the value
// reads 0 on both sides and the comparison can no longer see it. Emptying the
// map is forced -- the stale-entry rule below would fail on all three.
//
// The three findings are not thereby resolved, and are recorded here because
// deleting the entries would otherwise delete the knowledge:
//
//   - appkit NSViewExclusiveGestureBehaviorExclusive (1 -> 0) and
//     NSViewExclusiveGestureBehaviorNotExclusive (2 -> 0), 2026-08-14.
//   - metal MTLFloatingPointConversionRoundingModeTowardZero (1 -> 0),
//     2026-08-17. This one has a runtime consequence, not merely a cosmetic
//     one: the collision leaves TowardZero equal to ToNearestEven, so a caller
//     asking for truncation silently rounds. Measured, not assumed --
//     MTLFloatingPointConversionRoundingMode appears nowhere under the 25F70
//     SDK's System/Library/Frameworks, and the generated String lost its
//     TowardZero case because a duplicate case would not compile.
//
// All three are SDK skew: the release SDK does not declare the symbol, so it
// takes 0 rather than being omitted. The fix belongs with SDK-skew handling.
//
// This class needs an absolute check for the same reason TestNoAllOpaqueRecords
// exists -- a comparison cannot see a constant that was already wrong on both
// sides. 278 enum types in the tree currently have more than one zero-valued
// constant, and that population is UNMEASURED: it is not known how many are
// benign (a legitimate zero plus a genuinely absent symbol) and how many are
// collisions like the metal one. Sizing it is v0.7.1 work.
var knownFindings = map[string]string{}

func TestNoTypeDegradationAgainstBaseline(t *testing.T) {
	// This gate reads every generated file at two revisions and runs about
	// 170s, against a 4m default that the rest of the package eats into. A
	// package that panics on `go test ./...` reports no verdict at all, and a
	// gate that intermittently reports nothing is worse than one that is
	// declared expensive: the first teaches readers to rerun until green.
	if testing.Short() {
		t.Skip("compares every generated file at two revisions; run without -short")
	}
	root := repoRoot(t)
	base := os.Getenv("APPLE_GATE_BASE")
	if base == "" {
		base = baselineRev
	}
	// A baseline that does not resolve is a broken gate, not an empty one.
	if _, err := git(root, "rev-parse", "--verify", base+"^{commit}"); err != nil {
		t.Fatalf("baseline %q does not resolve: %v", base, err)
	}

	files := generatedFiles(t, root)
	if len(files) == 0 {
		t.Fatal("no generated files found; the gate would pass without comparing anything")
	}

	var (
		compared int
		added    int
		removed  int
		unparsed int
		findings []string
	)
	for _, path := range files {
		name := rel(root, path)
		oldSrc, err := git(root, "show", base+":"+name)
		if err != nil {
			added++
			continue
		}
		newSrc, err := os.ReadFile(path)
		if os.IsNotExist(err) {
			// The tree is shared and a regeneration can remove a file between
			// the listing above and this read. A file that is gone carries no
			// types to compare; it is not a degradation and not a broken gate.
			removed++
			continue
		}
		if err != nil {
			t.Fatalf("read %s: %v", name, err)
		}
		oldFile, err1 := parseGo(name, oldSrc)
		newFile, err2 := parseGo(name, newSrc)
		if err1 != nil || err2 != nil {
			unparsed++
			continue
		}
		compared++
		findings = append(findings, degradations(name, oldFile, newFile)...)
	}

	t.Logf("compared %d file(s) against %s; %d not in baseline, %d removed since listing, %d unparsable", compared, base, added, removed, unparsed)
	if compared == 0 {
		t.Fatalf("every generated file was skipped; nothing was compared")
	}
	sort.Strings(findings)
	seen := map[string]bool{}
	var unexpected []string
	for _, f := range findings {
		if _, known := knownFindings[f]; known {
			seen[f] = true
			continue
		}
		unexpected = append(unexpected, f)
	}
	// One line that accounts for every finding, because two unrelated counts
	// printed by one run get reconciled by whoever is reading under pressure,
	// and 9-new beside 3-listed has already been read as a total of 3.
	t.Logf("%d degradation(s) found: %d listed in knownFindings, %d new", len(findings), len(seen), len(unexpected))

	if len(unexpected) > 0 {
		t.Errorf("%d of %d degradation(s) against %s are not in knownFindings:\n%s",
			len(unexpected), len(findings), base, strings.Join(unexpected, "\n"))
	}

	// A listed finding that no longer reproduces means the list is stale, and
	// a stale entry is what lets the next real one hide behind it.
	var stale []string
	for f := range knownFindings {
		if !seen[f] {
			stale = append(stale, f)
		}
	}
	if len(stale) > 0 {
		sort.Strings(stale)
		t.Errorf("%d of %d entry(s) in knownFindings no longer reproduce; delete them:\n%s",
			len(stale), len(knownFindings), strings.Join(stale, "\n"))
	}
}

// TestNoUncommittedTypeDegradation compares the working tree against HEAD.
//
// It exists because the fixed baseline above cannot see one whole class of
// regression, and the class is not rare: a repair landed AFTER baselineRev can
// be silently undone, and the gate stays green because both sides of its
// comparison read the pre-repair type. os.OSLogCreate is the worked example —
// objc.ID at the baseline, kernel.Os_log_t once repaired, objc.ID again after
// the 2026-08-18 whole-tree regeneration. Against a baseline pinned before the
// repair, old and new both read objc.ID and nothing is reported.
//
// This is not an argument for a relative baseline; the reasoning above still
// holds and the fixed baseline stays. The two tests answer different questions.
// Fixed baseline: what damage does the tree carry? HEAD: what damage did the
// regeneration I have not committed yet introduce? Neither subsumes the other,
// and tonight only the second one can see the OSLogCreate revert.
//
// A clean tree compares nothing and passes, which is correct: with nothing
// uncommitted there is nothing for this test to be about. That is also this
// test's failure mode, so it does not get to decide on its own that the tree is
// clean — git is asked independently, and a disagreement is a broken gate
// rather than a quiet pass. An instrument whose "nothing to report" and
// "nothing was read" look identical is answering a question nobody asked.
func TestNoUncommittedTypeDegradation(t *testing.T) {
	root := repoRoot(t)

	files := generatedFiles(t, root)
	if len(files) == 0 {
		t.Fatal("no generated files found; the gate would pass without comparing anything")
	}

	// git's own answer, before ours: every tracked generated file that differs
	// from HEAD. Deletions are excluded — they carry no types to compare.
	out, err := git(root, "diff", "--name-only", "--diff-filter=d", "HEAD", "--", "*.gen.go")
	if err != nil {
		t.Fatalf("listing modified files: %v", err)
	}
	expected := map[string]bool{}
	for line := range strings.SplitSeq(strings.TrimSpace(string(out)), "\n") {
		if line != "" {
			expected[line] = true
		}
	}

	var compared, added int
	var findings []string
	for _, path := range files {
		name := rel(root, path)
		oldSrc, err := git(root, "show", "HEAD:"+name)
		if err != nil {
			continue // not tracked at HEAD; a new file degrades nothing
		}
		newSrc, err := os.ReadFile(path)
		if err != nil {
			continue
		}
		if bytes.Equal(oldSrc, newSrc) {
			continue
		}
		oldFile, err1 := parseGo(name, oldSrc)
		newFile, err2 := parseGo(name, newSrc)
		if err1 != nil || err2 != nil {
			continue
		}
		compared++
		delete(expected, name)
		findings = append(findings, degradations(name, oldFile, newFile)...)
		findings = append(findings, removals(name, oldFile, newFile)...)
		// Additions are the opposite direction and not a degradation, so they
		// are counted rather than reported. They are counted because EVERY rule
		// above compares a key present on both sides, which means a function
		// that only exists on the new side is examined by nothing at all.
		//
		// That is not hypothetical. VTFrameSiloCallFunctionForEachSampleBuffer
		// came back in the 2026-08-18 regeneration with three of its four
		// parameters, dropping the callback the function exists to deliver, and
		// no check here could see it: there was no old signature to compare.
		// A restored function is public API arriving unaudited.
		added += len(removals(name, newFile, oldFile))
	}

	// Whatever git listed and we did not compare was skipped by one of the
	// continues above — untracked, unreadable, or unparsable. Each is silent on
	// its own, and silence is what this check exists to break.
	if len(expected) > 0 {
		var missed []string
		for name := range expected {
			missed = append(missed, name)
		}
		sort.Strings(missed)
		t.Errorf("git reports %d modified generated file(s) this test did not compare:\n%s",
			len(missed), strings.Join(missed, "\n"))
	}

	sort.Strings(findings)
	t.Logf("compared %d modified file(s) against HEAD; %d degradation(s); %d exported function(s) added and therefore compared against nothing", compared, len(findings), added)
	if len(findings) > 0 {
		t.Errorf("%d degradation(s) in the working tree that HEAD does not carry:\n%s",
			len(findings), strings.Join(findings, "\n"))
	}
}

// removals reports exported functions and methods that oldFile declares and
// newFile does not.
//
// degradations cannot see this class at all: it compares types under shared
// keys, and a declaration that is gone has no key on the new side, so it is
// skipped by the same test that skips a function nobody touched. Unemitted and
// unchanged look identical to it.
//
// A dropped function cannot fail to build either — nothing references it once
// it is gone — so between that and the type comparison there was no instrument
// in the tree that could see public API leave. The 2026-08-18 regeneration
// removed eight exported accessors from three files under private/ and it was
// caught by hand-diffing function counts.
//
// This runs against HEAD, not the fixed baseline: deletions accumulate
// legitimately over a long history, but a regeneration that has not been
// committed yet should not be quietly deleting exported API.
//
// A REMOVAL IS NOT AUTOMATICALLY DAMAGE, and nothing here can tell the
// difference. Of the eight removals this found on 2026-08-18, six were
// regressions — accessors for properties the runtime reports as retained ObjC
// objects — and two were REPAIRS: AVSpeechUtterance.Action/SetAction bound a
// selector that does not exist on this OS, so deleting the phantom pair was
// correct. (Why the six went missing was not established; the split does not
// depend on knowing.)
//
// Only the ObjC runtime could separate them: class_copyPropertyList and
// property_getAttributes against the loaded framework. Neither the archive nor
// this comparison can, because both describe what was generated rather than
// what the class actually has. Treat this list as a population to diagnose, not
// a count of damage — reporting it as damage over-reports, which is the failure
// direction that gets a gate's real findings dismissed along with its false
// ones.
func removals(name string, oldFile, newFile *ast.File) []string {
	declared := func(file *ast.File) map[string]bool {
		out := map[string]bool{}
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || !fn.Name.IsExported() {
				continue
			}
			key := fn.Name.Name
			if fn.Recv != nil && len(fn.Recv.List) == 1 {
				key = strings.TrimPrefix(render(fn.Recv.List[0].Type), "*") + "." + key
			}
			out[key] = true
		}
		return out
	}
	old, new := declared(oldFile), declared(newFile)
	var out []string
	for key := range old {
		if !new[key] {
			out = append(out, fmt.Sprintf("%s: %s: exported function removed", name, key))
		}
	}
	sort.Strings(out)
	return out
}

// degradations reports every degradation between two versions of one file.
func degradations(name string, oldFile, newFile *ast.File) []string {
	var out []string

	// Struct fields and function signatures degrade the same way and are
	// compared by the same rules. They are separate populations only because
	// they are read out of the AST differently; keeping the rules in one place
	// is what stops the two from drifting apart, which is how method returns
	// went unmeasured while fields were covered.
	for _, pair := range []struct{ old, new map[string]string }{
		{structFieldTypes(oldFile), structFieldTypes(newFile)},
		{funcSignatureTypes(oldFile), funcSignatureTypes(newFile)},
	} {
		// Positions are only comparable when the arity is unchanged. Adding a
		// result shifts every later one, so "(result 0): error -> bool" on a
		// function that went from one result to two is describing a rename of
		// position 0, not a degradation — ANERequest.Set gaining an (ok, error)
		// pair reported exactly that. Where arity moved, report the arity and
		// stop: everything downstream of it is a position artifact.
		shifted := map[string]bool{}
		for key, oldType := range pair.old {
			fn, isArity := strings.CutSuffix(key, "(arity)")
			if !isArity {
				continue
			}
			if newType, ok := pair.new[key]; ok && newType != oldType {
				shifted[fn] = true
			}
		}
		for key, oldType := range pair.old {
			newType, ok := pair.new[key]
			if !ok || newType == oldType {
				continue
			}
			if fn, _, found := strings.Cut(key, "("); found && shifted[fn] && !strings.HasSuffix(key, "(arity)") {
				continue
			}
			if why := degradedTo(oldType, newType); why != "" {
				out = append(out, fmt.Sprintf("%s: %s: %s -> %s (%s)", name, key, oldType, newType, why))
			}
		}
	}

	oldConsts, newConsts := constValues(oldFile), constValues(newFile)
	for key, oldValue := range oldConsts {
		if newValue, ok := newConsts[key]; ok && oldValue != "0" && newValue == "0" {
			out = append(out, fmt.Sprintf("%s: %s: %s -> 0 (constant collapsed to zero)", name, key, oldValue))
		}
	}
	return out
}

// degradedTo reports why newType is a degradation of oldType, or "" if it is
// not one. Only the direction that loses information is a finding: a type that
// gains an identity — unsafe.Pointer becoming xpc.Xpc_object_t — is the repair,
// not the damage.
func degradedTo(oldType, newType string) string {
	switch {
	case strings.HasSuffix(oldType, "results"):
		// The arity entry funcSignatureTypes records. Any change is a finding:
		// a generated binding mirrors a fixed C signature, so both a lost and a
		// gained parameter mean the emitted call no longer matches the symbol.
		return "signature arity changed"
	case isIntegerType(oldType) && newType == "unsafe.Pointer":
		return "integer became a pointer"
	case isNamedType(oldType) && !isNamedType(newType):
		return "named type lost its identity"
	case isNamedType(oldType) && isErasedType(newType):
		// Both sides name a type, so the rule above cannot see this: the new
		// name is one of the generator's fallback handles, which carries an
		// address as a plain word. os.OSLogCreate returning objc.ID instead of
		// kernel.Os_log_t is the same loss as returning unsafe.Pointer, spelled
		// as a named type.
		return "named type lost its identity"
	}
	return ""
}

// isErasedType reports whether goType is one of the handle spellings the
// generator falls back to when it cannot resolve a type. These are named types,
// so isNamedType accepts them; they carry no identity all the same.
func isErasedType(goType string) bool {
	switch goType {
	case "objc.ID", "objectivec.ID":
		return true
	}
	return false
}

// funcSignatureTypes maps each function's parameters and results to their
// rendered types, for every function and method declared in the file.
//
// Parameters and results are keyed by POSITION, not by name: the generator
// renames parameters freely and a rename is not a degradation. The receiver
// type is part of the key so two methods with one name on different types stay
// separate, and its own type is not compared — a method moving between types is
// a different change than a signature losing one.
func funcSignatureTypes(file *ast.File) map[string]string {
	out := map[string]string{}
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Type == nil {
			continue
		}
		key := fn.Name.Name
		if fn.Recv != nil && len(fn.Recv.List) == 1 {
			// Rendered without the pointer, so a receiver changing between T
			// and *T does not split the key and hide every signature under it.
			key = strings.TrimPrefix(render(fn.Recv.List[0].Type), "*") + "." + key
		}
		record := func(kind string, fields *ast.FieldList) {
			if fields == nil {
				return
			}
			i := 0
			for _, field := range fields.List {
				rendered := render(field.Type)
				// A field with n names declares n entries of one type; one with
				// no names declares exactly one.
				n := len(field.Names)
				if n == 0 {
					n = 1
				}
				for j := 0; j < n; j++ {
					out[fmt.Sprintf("%s(%s %d)", key, kind, i)] = rendered
					i++
				}
			}
		}
		record("param", fn.Type.Params)
		record("result", fn.Type.Results)
		// Arity is recorded as its own comparable value. Without it a parameter
		// that vanishes has no key on the new side and is skipped by the same
		// branch that skips an untouched one — unemitted and unchanged sharing a
		// code path, one level below whole functions.
		//
		// VTFrameSiloCallFunctionForEachSampleBuffer is why: the C symbol takes
		// four parameters, the generator emitted three, dropping the callback
		// the function exists to deliver. That builds, links, and calls the real
		// symbol with three arguments where it wants four.
		out[key+"(arity)"] = fmt.Sprintf("%d params, %d results", count(fn.Type.Params), count(fn.Type.Results))
	}
	return out
}

// count returns the number of entries a field list declares, which is not the
// number of ast.Fields in it: "a, b int" is one field and two parameters.
func count(fields *ast.FieldList) int {
	if fields == nil {
		return 0
	}
	n := 0
	for _, field := range fields.List {
		if len(field.Names) == 0 {
			n++
			continue
		}
		n += len(field.Names)
	}
	return n
}

// structFieldTypes maps "Type.Field" to the field's rendered type, for every
// struct type declared in the file. Embedded fields are keyed by their type.
func structFieldTypes(file *ast.File) map[string]string {
	out := map[string]string{}
	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.TYPE {
			continue
		}
		for _, spec := range gen.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			st, ok := ts.Type.(*ast.StructType)
			if !ok || st.Fields == nil {
				continue
			}
			for _, field := range st.Fields.List {
				rendered := render(field.Type)
				if len(field.Names) == 0 {
					out[ts.Name.Name+"."+rendered] = rendered
					continue
				}
				for _, id := range field.Names {
					out[ts.Name.Name+"."+id.Name] = rendered
				}
			}
		}
	}
	return out
}

// constValues maps each constant's name to its literal value. A constant whose
// value is an expression rather than a literal is left out: this gate is about
// values the generator writes down.
func constValues(file *ast.File) map[string]string {
	out := map[string]string{}
	for _, decl := range file.Decls {
		gen, ok := decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.CONST {
			continue
		}
		for _, spec := range gen.Specs {
			vs, ok := spec.(*ast.ValueSpec)
			if !ok {
				continue
			}
			for i, id := range vs.Names {
				if i >= len(vs.Values) {
					continue
				}
				if lit, ok := vs.Values[i].(*ast.BasicLit); ok && lit.Kind == token.INT {
					out[id.Name] = lit.Value
				}
			}
		}
	}
	return out
}

func isIntegerType(goType string) bool {
	switch goType {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64", "uintptr",
		"byte", "rune":
		return true
	}
	return false
}

// isNamedType reports whether goType names a declared type rather than a Go
// builtin, an array, a pointer, or a func. A qualified name counts: coremedia
// .CMTime is as much an identity as CMTime.
func isNamedType(goType string) bool {
	if goType == "" || isIntegerType(goType) {
		return false
	}
	switch goType {
	case "bool", "string", "float32", "float64", "complex64", "complex128", "any", "unsafe.Pointer":
		return false
	}
	if strings.ContainsAny(goType, "[]*(){} \t") {
		return false
	}
	return true
}

func render(node ast.Expr) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, token.NewFileSet(), node); err != nil {
		return ""
	}
	return buf.String()
}

func parseGo(name string, src []byte) (*ast.File, error) {
	return parser.ParseFile(token.NewFileSet(), name, src, parser.SkipObjectResolution)
}

func git(dir string, args ...string) ([]byte, error) {
	cmd := exec.Command("git", args...)
	cmd.Dir = dir
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("git %s: %v: %s", strings.Join(args, " "), err, stderr.String())
	}
	return stdout.Bytes(), nil
}

func TestRemovalsDetectsDroppedExports(t *testing.T) {
	tests := []struct {
		name string
		old  string
		new  string
		want []string
	}{
		{
			name: "exported accessor pair dropped",
			old:  "package p\n\nfunc (m MLEngine) Snapshot() T { return zero }\nfunc (m MLEngine) SetSnapshot(v T) {}\n",
			new:  "package p\n",
			// Sorted, so SetSnapshot precedes Snapshot.
			want: []string{"MLEngine.SetSnapshot: exported function removed", "MLEngine.Snapshot: exported function removed"},
		},
		{
			name: "unexported drop is not reported",
			old:  "package p\n\nfunc trySnapshot() {}\n",
			new:  "package p\n",
			want: nil,
		},
		{
			name: "a method moving between types is reported against the type it left",
			old:  "package p\n\nfunc (a A) Do() {}\n",
			new:  "package p\n\nfunc (b B) Do() {}\n",
			want: []string{"A.Do: exported function removed"},
		},
		{
			name: "a pointer receiver is the same key as a value receiver",
			old:  "package p\n\nfunc (a A) Do() {}\n",
			new:  "package p\n\nfunc (a *A) Do() {}\n",
			want: nil,
		},
		{
			name: "adding functions removes nothing",
			old:  "package p\n\nfunc Keep() {}\n",
			new:  "package p\n\nfunc Keep() {}\nfunc Added() {}\n",
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldFile, err := parseGo("old.go", []byte(tt.old))
			if err != nil {
				t.Fatal(err)
			}
			newFile, err := parseGo("new.go", []byte(tt.new))
			if err != nil {
				t.Fatal(err)
			}
			got := removals("f.gen.go", oldFile, newFile)
			if len(got) != len(tt.want) {
				t.Fatalf("removals() = %v, want %v", got, tt.want)
			}
			for i, w := range tt.want {
				if !strings.Contains(got[i], w) {
					t.Fatalf("removals()[%d] = %q, want it to contain %q", i, got[i], w)
				}
			}
		})
	}
}

// The gate's three rules, stated on inputs that cannot drift. Each pair is one
// degradation the tree actually shipped.
func TestDegradationsDetectsEachClass(t *testing.T) {
	tests := []struct {
		name string
		old  string
		new  string
		want string
	}{
		{
			name: "integer field becomes a pointer",
			old:  "package p\n\ntype Timespec struct{ Tv_sec int }\n",
			new:  "package p\n\ntype Timespec struct{ Tv_sec unsafe.Pointer }\n",
			want: "integer became a pointer",
		},
		{
			name: "named field type collapses to its underlying",
			old:  "package p\n\ntype Backtrace struct{ Btc_flags Backtrace_flags_t }\n",
			new:  "package p\n\ntype Backtrace struct{ Btc_flags uint64 }\n",
			want: "named type lost its identity",
		},
		{
			name: "constant collapses to zero",
			old:  "package p\n\nconst AVMetricPlaybackModeAirPlayVideo AVMetricPlaybackMode = 1\n",
			new:  "package p\n\nconst AVMetricPlaybackModeAirPlayVideo AVMetricPlaybackMode = 0\n",
			want: "constant collapsed to zero",
		},
		{
			name: "function return loses a named type to a raw handle",
			old:  "package p\n\nfunc OSLogCreate(subsystem string, category string) kernel.Os_log_t { return zero }\n",
			new:  "package p\n\nfunc OSLogCreate(subsystem string, category string) objc.ID { return zero }\n",
			want: "named type lost its identity",
		},
		{
			name: "accessor return becomes unsafe.Pointer",
			old:  "package p\n\nfunc (a AUAudioUnit) OsWorkgroup() objectivec.IObject { return zero }\n",
			new:  "package p\n\nfunc (a AUAudioUnit) OsWorkgroup() unsafe.Pointer { return zero }\n",
			want: "named type lost its identity",
		},
		{
			name: "accessor return becomes a byte array",
			old:  "package p\n\nfunc (s *Cap) Rate_trend_suggestion() Ifnet_rate_trend { return zero }\n",
			new:  "package p\n\nfunc (s *Cap) Rate_trend_suggestion() [4]byte { return zero }\n",
			want: "named type lost its identity",
		},
		{
			name: "setter parameter becomes unsafe.Pointer",
			old:  "package p\n\nfunc (m MLEngine) SetSnapshot(value MLSnapshot) {}\n",
			new:  "package p\n\nfunc (m MLEngine) SetSnapshot(value unsafe.Pointer) {}\n",
			want: "named type lost its identity",
		},
		{
			name: "integer parameter becomes a pointer",
			old:  "package p\n\nfunc Set(v uint64) {}\n",
			new:  "package p\n\nfunc Set(v unsafe.Pointer) {}\n",
			want: "integer became a pointer",
		},
		{
			name: "a parameter silently disappearing",
			old:  "package p\n\nfunc Call(silo S, tr T, refcon uintptr, cb func()) int32 { return 0 }\n",
			new:  "package p\n\nfunc Call(silo S, tr T, refcon uintptr) int32 { return 0 }\n",
			want: "signature arity changed",
		},
		{
			name: "grouped parameters count individually",
			old:  "package p\n\nfunc Call(a, b int) {}\n",
			new:  "package p\n\nfunc Call(a int) {}\n",
			want: "signature arity changed",
		},
		{
			name: "gaining a result reports arity, not a position artifact",
			old:  "package p\n\nfunc (a A) Set(v T) error { return nil }\n",
			new:  "package p\n\nfunc (a A) Set(v T) (bool, error) { return false, nil }\n",
			want: "signature arity changed",
		},
		{
			name: "a signature that only gains identity is not a degradation",
			old:  "package p\n\nfunc Get() unsafe.Pointer { return nil }\n",
			new:  "package p\n\nfunc Get() xpc.Xpc_object_t { return zero }\n",
			want: "",
		},
		{
			name: "renaming a parameter is not a degradation",
			old:  "package p\n\nfunc Set(old Flags_t) {}\n",
			new:  "package p\n\nfunc Set(renamed Flags_t) {}\n",
			want: "",
		},
		{
			name: "unchanged file reports nothing",
			old:  "package p\n\ntype Timespec struct{ Tv_sec int }\n",
			new:  "package p\n\ntype Timespec struct{ Tv_sec int }\n",
			want: "",
		},
		{
			name: "widening within the integers is not a degradation here",
			old:  "package p\n\ntype S struct{ F int32 }\n",
			new:  "package p\n\ntype S struct{ F int64 }\n",
			want: "",
		},
		{
			name: "a field gaining a name is not a degradation",
			old:  "package p\n\ntype S struct{ F uint64 }\n",
			new:  "package p\n\ntype S struct{ F Flags_t }\n",
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldFile, err := parseGo("old.go", []byte(tt.old))
			if err != nil {
				t.Fatal(err)
			}
			newFile, err := parseGo("new.go", []byte(tt.new))
			if err != nil {
				t.Fatal(err)
			}
			got := degradations("f.gen.go", oldFile, newFile)
			if tt.want == "" {
				if len(got) != 0 {
					t.Fatalf("degradations() = %v, want none", got)
				}
				return
			}
			if len(got) != 1 || !strings.Contains(got[0], tt.want) {
				t.Fatalf("degradations() = %v, want one finding containing %q", got, tt.want)
			}
		})
	}
}
