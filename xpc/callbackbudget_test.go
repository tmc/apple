// Copyright 2026 The apple Authors.

package xpc

// The callback budget gate.
//
// purego's callback table is a fixed array of 2000 entries with no unregister
// API. Exhausting it panics from a native callback thread, which aborts the
// process, so a binding that registers a callback per message dies after a few
// hundred round trips. That is exactly how the decode-path leak was found: a
// human ran `go test -count=300 -tags xpclive -run TestLiveRoundTrip`, and the
// service died at some N in (200, 300]. Nothing in the suite noticed, because
// every test ran a handful of messages and every assertion was about a value,
// not about a resource.
//
// This file adds the missing gate in two shapes:
//
//   - TestCallbackBudgetIsConstantInMessages measures the resource across two
//     workload sizes and fails unless consumption is O(1) in the number of
//     messages. It is a budget, not a snapshot: it never compares against a
//     recorded constant, so it cannot be silenced by updating one.
//   - TestNewXPCBlockCallSitesAreLedgered is the static half. A dynamic budget
//     can only see paths a non-live test can drive; the async reply path needs
//     a real peer. So every newXPCBlock call site in the package is enumerated
//     from source and must appear in a reviewed ledger, classified as
//     session-lifetime or per-message.
//
// Instrument choice: blockCount (defined in xpc.highlevel.gen_test.go), which
// reads len(xpcBlockKeepalive)/3. purego exposes no counter of registered
// callbacks — the table and its index are unexported and there is no accessor —
// so the count has to be taken on our side of the boundary. The keepalive slice
// is the right place to take it: newXPCBlock appends to it in the same critical
// section in which it calls purego.NewCallback, and nothing else appends to it,
// so its length is an exact tally of registrations rather than a proxy. Its
// unit (three entries per registration: the Go func, the block literal, the
// descriptor) is checked by the sensitivity step in
// TestApplierBlockSharesOneCallback and again by the positive control below.

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"testing"
	"unsafe"
)

// budgetWorkload runs n units of a per-message workload.
type budgetWorkload func(t *testing.T, n int)

// checkConstantInN measures work at two sizes and reports whether callback
// consumption is O(1) in n.
//
// The two-point comparison is what makes this a budget rather than a snapshot.
// It asserts a shape — that the cost of the larger workload is no greater than
// the cost of the smaller one — so it holds no number that a future maintainer
// can "update" to make a leak pass. A leak of even one slot per message shows
// up as small > 0 and large >= 4*small.
//
// Warm-up is charged separately and deliberately: shared trampolines and any
// one-off symbol resolution register on first use, and attributing those to the
// loop would make a correct implementation look like a leak at small n.
func checkConstantInN(t *testing.T, work budgetWorkload, small, large int) error {
	t.Helper()

	// Warm-up: pay every one-off registration before either measurement.
	work(t, 2)

	before := blockCount()
	work(t, small)
	costSmall := blockCount() - before

	before = blockCount()
	work(t, large)
	costLarge := blockCount() - before

	t.Logf("callback slots: %d messages cost %d, %d messages cost %d", small, costSmall, large, costLarge)

	if costLarge > costSmall {
		return fmt.Errorf("callback consumption grows with the workload: %d messages cost %d slots, %d messages cost %d "+
			"(%.3f slots per additional message); a per-message registration exhausts purego's 2000-entry table",
			small, costSmall, large, costLarge, float64(costLarge-costSmall)/float64(large-small))
	}
	if costSmall != 0 {
		return fmt.Errorf("%d messages cost %d callback slots after warm-up, want 0: the per-message path registers callbacks",
			small, costSmall)
	}
	return nil
}

// decodeWorkload is the path that leaked: a message is decoded once per unit,
// as a service handling one request per message does.
func decodeWorkload(t *testing.T, n int) {
	t.Helper()
	msg := Dictionary{
		"op":    "add",
		"inner": Dictionary{"a": int64(1), "b": "two"},
		"list":  []any{int64(1), "x", true},
	}
	for i := 0; i < n; i++ {
		raw, err := dictionaryToRawObject(msg)
		if err != nil {
			t.Fatalf("unit %d: dictionaryToRawObject: %v", i, err)
		}
		got := ReceivedMessage{raw: raw}.Dictionary()
		// The work must really happen. A decode that silently produced
		// nothing would also consume no callbacks, and the budget would
		// pass by never doing anything.
		if got["op"] != "add" {
			t.Fatalf("unit %d: op = %#v, want \"add\"", i, got["op"])
		}
		inner, ok := got["inner"].(Dictionary)
		if !ok || inner["b"] != "two" {
			t.Fatalf("unit %d: inner = %#v", i, got["inner"])
		}
		list, ok := got["list"].([]any)
		if !ok || len(list) != 3 || list[1] != "x" {
			t.Fatalf("unit %d: list = %#v", i, got["list"])
		}
		releaseRaw(raw)
	}
}

// leakyDecodeWorkload is the mutation control: the same workload with the
// pre-fix per-call registration restored. rawObjectToDictionary used to call
// newXPCBlock inside its body, once per container, so this arm registers one
// block per unit — the smallest version of the defect the gate must catch.
func leakyDecodeWorkload(t *testing.T, n int) {
	t.Helper()
	for i := 0; i < n; i++ {
		if _, err := newXPCBlock(func(_ uintptr, _ unsafe.Pointer) bool { return true }); err != nil {
			t.Fatalf("unit %d: newXPCBlock: %v", i, err)
		}
	}
}

// TestCallbackBudgetIsConstantInMessages is Gate A.
func TestCallbackBudgetIsConstantInMessages(t *testing.T) {
	if frameworkHandle == 0 {
		t.Skip("XPC framework unavailable")
	}

	// Positive control (sensitivity): the instrument must be able to see one
	// registration, in units of one. Without it, every zero below could mean
	// the instrument is blind rather than the code clean.
	before := blockCount()
	if _, err := newXPCBlock(func(_ uintptr) {}); err != nil {
		t.Fatalf("newXPCBlock: %v", err)
	}
	if spent := blockCount() - before; spent != 1 {
		t.Fatalf("one registration measured as %d slots: blockCount is in the wrong unit, so no verdict below is about the tree", spent)
	}

	const small, large = 25, 100
	if err := checkConstantInN(t, decodeWorkload, small, large); err != nil {
		t.Errorf("decode path is not within budget: %v", err)
	}
}

// TestCallbackBudgetGateIsSensitive is the mutation control for the gate above.
//
// It runs the identical predicate, checkConstantInN, over a workload that
// restores the pre-fix behaviour (one registration per unit) and fails if the
// predicate reports success. Without this, a budget that could never fail —
// because the instrument was blind, the workload empty, or the comparison
// vacuous — would stay green forever.
//
// Population examined: large-small = 75 additional units, and 125 registrations
// in total across the two measured arms plus warm-up. It is not zero: a control
// that examines nothing passes by never looking.
func TestCallbackBudgetGateIsSensitive(t *testing.T) {
	if frameworkHandle == 0 {
		t.Skip("XPC framework unavailable")
	}
	const small, large = 25, 100
	err := checkConstantInN(t, leakyDecodeWorkload, small, large)
	if err == nil {
		t.Fatalf("mutation control failed: the budget reported success on a workload that registers one callback per unit; "+
			"the gate cannot detect the leak it exists to detect (examined %d units)", small+large+2)
	}
	t.Logf("mutation control: the budget can fail; it reported: %v", err)
}

// blockSite is one newXPCBlock call site, as read from the package source.
type blockSite struct {
	fn         string // enclosing function or method, e.g. "(*Session).callAsyncDictionary"
	perMessage bool   // true if the site runs once per message rather than once per session or listener
	why        string
}

// newXPCBlockLedger records every reviewed newXPCBlock call site.
//
// A dynamic budget can only measure paths a non-live test can drive. The async
// reply path cannot be driven without a real peer, so the static half of the
// gate enumerates the call sites from source and requires each to have been
// classified. Adding a registration to a per-message path without touching this
// ledger fails the test; so does removing a ledgered site, which forces the
// ledger to be updated when a leak is fixed rather than left describing a tree
// that no longer exists.
var newXPCBlockLedger = []blockSite{
	{
		fn:         "newListener",
		perMessage: false,
		why:        "one block per listener, registered before the listener is activated; a process creates a bounded number of listeners",
	},
}

// wantPerMessageSites is the number of ledgered per-message registrations. It
// is a ceiling on a known defect, not a target: it must go to zero, and when it
// does this test fails until the ledger says so.
//
// It is now zero. (*Session).callAsyncDictionary held the last entry: it
// registered one callback per CallDictionary on the cancellable-context path,
// and any caller passing a context with a deadline or a cancel aborted the
// process at about two thousand calls. It now builds an xpcReplyBlock, which
// carries a token and dispatches to one shared trampoline, so it no longer
// calls newXPCBlock at all and the census below no longer finds it.
const wantPerMessageSites = 0

// TestNewXPCBlockCallSitesAreLedgered is the static half of Gate A.
func TestNewXPCBlockCallSitesAreLedgered(t *testing.T) {
	found := newXPCBlockCallSites(t)
	if len(found) == 0 {
		// A census that finds nothing is an instrument result, not a clean
		// tree: newXPCBlock exists and is called.
		t.Fatalf("census found no newXPCBlock call sites; the scanner is broken, so no verdict below is about the tree")
	}

	ledger := map[string]blockSite{}
	for _, s := range newXPCBlockLedger {
		ledger[s.fn] = s
	}

	perMessage := 0
	for _, fn := range found {
		s, ok := ledger[fn]
		if !ok {
			t.Errorf("newXPCBlock call site %s is not in newXPCBlockLedger: classify it as session-lifetime or "+
				"per-message. A per-message registration exhausts purego's 2000-entry callback table and aborts the process.", fn)
			continue
		}
		if s.perMessage {
			perMessage++
			t.Logf("ledgered per-message registration: %s — %s", fn, s.why)
		}
	}
	for fn := range ledger {
		if !contains(found, fn) {
			t.Errorf("newXPCBlockLedger records %s, which no longer calls newXPCBlock: update the ledger "+
				"(and wantPerMessageSites) so it describes the tree it gates", fn)
		}
	}
	if perMessage != wantPerMessageSites {
		t.Errorf("found %d per-message newXPCBlock registrations, ledger expects %d: "+
			"if a leak was fixed, lower wantPerMessageSites; if one was added, it is a process-aborting defect",
			perMessage, wantPerMessageSites)
	}
}

// TestNewXPCBlockCensusIsSensitive is the mutation control for the census: it
// runs the same scanner over a synthetic file containing one call inside a
// method and one inside a plain function, and fails if either is missed.
//
// This is a positive control, and it proves sensitivity only. The negative
// control is in the same run: the synthetic file also calls a similarly named
// function that must NOT be reported, so a scanner that reported every call
// expression would fail here rather than pass by over-reporting.
func TestNewXPCBlockCensusIsSensitive(t *testing.T) {
	dir := t.TempDir()
	src := `package fake

func newXPCBlock(any) (int, error) { return 0, nil }
func newXPCApplierBlock(any) int   { return 0 }

type S struct{}

func (s *S) method() { newXPCBlock(nil) }
func plain()         { newXPCBlock(nil) }
func decoy()         { newXPCApplierBlock(nil) }
`
	if err := os.WriteFile(filepath.Join(dir, "fake.go"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	got := scanNewXPCBlockCallSites(t, dir)
	want := []string{"(*S).method", "plain"}
	if strings.Join(got, ",") != strings.Join(want, ",") {
		t.Fatalf("scanner reported %v, want %v: the census is either blind (positive control) or over-reporting (decoy)", got, want)
	}
	t.Logf("mutation control: the census sees both shapes and rejects the decoy")
}

// newXPCBlockCallSites returns the sorted enclosing-function names of every
// newXPCBlock call in this package's non-test source.
func newXPCBlockCallSites(t *testing.T) []string {
	t.Helper()
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	return scanNewXPCBlockCallSites(t, wd)
}

func scanNewXPCBlockCallSites(t *testing.T, dir string) []string {
	t.Helper()
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, func(fi os.FileInfo) bool {
		return !strings.HasSuffix(fi.Name(), "_test.go")
	}, 0)
	if err != nil {
		t.Fatalf("parsing %s: %v", dir, err)
	}
	var out []string
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			for _, decl := range file.Decls {
				fn, ok := decl.(*ast.FuncDecl)
				if !ok || fn.Body == nil {
					continue
				}
				name := funcDeclName(fn)
				// Skip the definition itself; it is not a call site.
				if name == "newXPCBlock" {
					continue
				}
				calls := false
				ast.Inspect(fn.Body, func(n ast.Node) bool {
					call, ok := n.(*ast.CallExpr)
					if !ok {
						return true
					}
					if id, ok := call.Fun.(*ast.Ident); ok && id.Name == "newXPCBlock" {
						calls = true
					}
					return true
				})
				if calls {
					out = append(out, name)
				}
			}
		}
	}
	sort.Strings(out)
	return out
}

// funcDeclName renders a function or method name the way a stack trace does.
func funcDeclName(fn *ast.FuncDecl) string {
	if fn.Recv == nil || len(fn.Recv.List) == 0 {
		return fn.Name.Name
	}
	recv := fn.Recv.List[0].Type
	star := ""
	if s, ok := recv.(*ast.StarExpr); ok {
		star = "*"
		recv = s.X
	}
	id, ok := recv.(*ast.Ident)
	if !ok {
		return fn.Name.Name
	}
	return fmt.Sprintf("(%s%s).%s", star, id.Name, fn.Name.Name)
}

func contains(list []string, s string) bool {
	for _, v := range list {
		if v == s {
			return true
		}
	}
	return false
}
