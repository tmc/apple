// Copyright 2026 The tmc/apple Authors. All rights reserved.

package xpc_test

// The xpclive suite (xpc_live_test.go, xpc_cancel_live_test.go and
// xpc_codec_live_test.go; 21 TestLive* functions) is behind
// //go:build xpclive. Nothing in the tree passes that tag, so the suite was
// neither run nor compiled: a rename in xpc.highlevel.gen.go rotted it
// silently and no gate reported it.
//
// This file is the invoker. It follows the buildgate precedent: a test that
// imports nothing from the module and shells out to the go tool, so the switch
// that runs it is the one every "go test ./..." already flips.
//
// Two distinct failures are gated, and they stay distinct:
//
//   - compiling the suite is UNCONDITIONAL. That is the part that rots.
//   - running the suite needs launchd and is gated on XPC_LIVE=1. When it is
//     not run, the reason is reported, and SKIPPED is counted apart from clean.

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// Packages are named by import path, not by relative directory: the test's
// working directory is this package, so "./xpc/" resolves to nothing and every
// verdict would be about a path mistake rather than about the tree. See
// internal/buildgate/buildgate_test.go:24-27.
const (
	xpcPath       = "github.com/tmc/apple/xpc"
	buildgatePath = "github.com/tmc/apple/internal/buildgate"
)

func TestXPCLiveCompiles(t *testing.T) {
	// Canary: buildgate has no dependencies and must always vet. If it does
	// not, an UNBUILDABLE verdict below would be measuring the toolchain.
	if out, err := exec.Command("go", "vet", buildgatePath).CombinedOutput(); err != nil {
		t.Fatalf("harness is broken, not the tree: vetting %s failed, so no verdict below can be trusted: %v\n%s",
			buildgatePath, err, out)
	}
	out, err := exec.Command("go", "vet", "-tags", "xpclive", xpcPath).CombinedOutput()
	if err != nil {
		t.Errorf("xpclive suite UNBUILDABLE (not measured, not clean): %v\n%s", err, out)
		return
	}
	t.Logf("xpclive: compiled")
}

// TestXPCLiveCompileGateIsSensitive is the mutation control for the gate
// above: it vets a deliberately broken throwaway package and fails if go vet
// reports success. Without it, a gate that can never fail would stay green
// forever.
func TestXPCLiveCompileGateIsSensitive(t *testing.T) {
	dir := t.TempDir()
	src := "package broken\n\nfunc F() int { return \"not an int\" }\n"
	if err := os.WriteFile(filepath.Join(dir, "broken.go"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(dir, "go.mod"), []byte("module broken\n\ngo 1.24\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	cmd := exec.Command("go", "vet", ".")
	cmd.Dir = dir
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatalf("mutation control failed: go vet reported success on a package that does not compile\n%s", out)
	}
	t.Logf("mutation control: the compile gate can fail")
}

func TestXPCLiveRuns(t *testing.T) {
	// Recursion guard: the child inherits this environment, so without it the
	// child's own TestXPCLiveRuns would re-invoke the suite.
	if os.Getenv("XPC_LIVE_CHILD") == "1" {
		return
	}
	reason := ""
	switch {
	case runtime.GOOS != "darwin":
		reason = "XPC is only available on darwin"
	case os.Getenv("XPC_LIVE") != "1":
		reason = "XPC_LIVE unset"
	}
	if reason != "" {
		// This counts the child subprocess this gate would launch, not the
		// TestLive functions in this binary: when the binary itself was built
		// with -tags xpclive they have already run in-process, and reporting
		// "0 run" for them would contradict the PASS lines above.
		t.Logf("xpclive: child suite invocations: 1 compiled, 0 run (SKIPPED: %s)", reason)
		t.Skip(reason)
	}
	cmd := exec.Command("go", "test", "-tags", "xpclive", "-v", "-count=1",
		"-run", "TestLive", xpcPath)
	cmd.Env = append(os.Environ(), "XPC_LIVE_CHILD=1")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("xpclive suite failed: %v\n%s", err, out)
		return
	}
	// The count is PASS lines, so it includes subtests (t.Run cases inside
	// TestLive* functions); it is not the number of TestLive functions.
	t.Logf("xpclive: child suite invocations: 1 compiled, 1 run; %d TestLive PASS lines (functions and subtests)",
		strings.Count(string(out), "--- PASS: TestLive"))
}

// wantLiveTestInventory is the xpclive suite as this gate believes it to be:
// every file carrying the xpclive build tag, and the number of TestLive*
// functions each declares.
//
// It exists because the prose above ("16 TestLive* functions") was written
// when the suite was one file and was still saying 16 after it had grown to
// three files and 20 functions. That is exactly the drift these gates are for,
// and no gate caught it: the compile gate vets whatever is there, and the run
// gate counts PASS lines only on the runs it actually performs, which on a
// machine without XPC_LIVE=1 is none. Both are blind to a test that was
// deleted, renamed out of the TestLive prefix, or added in a file nobody
// wired up.
//
// The counts are pinned deliberately. Changing the suite is meant to require
// editing this table, which is the moment the header prose gets re-read.
var wantLiveTestInventory = map[string]int{
	"xpc_live_test.go":        16,
	"xpc_cancel_live_test.go": 3,
	"xpc_codec_live_test.go":  2,
}

// countLiveTests reports how many TestLive* functions src declares, and
// whether src is part of the xpclive suite at all. Both answers come from the
// same pass so that a file which loses its build tag cannot keep contributing
// its functions to the total.
func countLiveTests(src string) (n int, tagged bool) {
	for _, line := range strings.Split(src, "\n") {
		switch {
		case strings.HasPrefix(line, "//go:build xpclive"):
			tagged = true
		case strings.HasPrefix(line, "func TestLive"):
			n++
		}
	}
	return n, tagged
}

// TestXPCLiveInventoryIsSensitive is the mutation control for the counter:
// synthetic sources with known answers, including the two shapes that would
// silently deflate the real total (an untagged file, and a live test renamed
// out of the prefix).
func TestXPCLiveInventoryIsSensitive(t *testing.T) {
	cases := []struct {
		name       string
		src        string
		wantN      int
		wantTagged bool
	}{
		{"tagged two", "//go:build xpclive\n\nfunc TestLiveA(t *testing.T) {}\nfunc TestLiveB(t *testing.T) {}\n", 2, true},
		{"untagged", "func TestLiveA(t *testing.T) {}\n", 1, false},
		{"renamed out of the prefix", "//go:build xpclive\n\nfunc TestSomethingElse(t *testing.T) {}\n", 0, true},
		{"not at line start", "//go:build xpclive\n\n// func TestLiveA is documented here\n", 0, true},
	}
	for _, c := range cases {
		n, tagged := countLiveTests(c.src)
		if n != c.wantN || tagged != c.wantTagged {
			t.Errorf("%s: countLiveTests = (%d, %v), want (%d, %v)", c.name, n, tagged, c.wantN, c.wantTagged)
		}
	}
}

// TestXPCLiveInventory fails when the suite on disk stops matching
// wantLiveTestInventory, and when the header comment in this file stops
// stating the resulting total.
func TestXPCLiveInventory(t *testing.T) {
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatalf("read package directory: %v", err)
	}
	got := map[string]int{}
	scanned := 0
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), "_test.go") {
			continue
		}
		b, err := os.ReadFile(e.Name())
		if err != nil {
			t.Fatalf("read %s: %v", e.Name(), err)
		}
		scanned++
		n, tagged := countLiveTests(string(b))
		if !tagged {
			continue
		}
		got[e.Name()] = n
	}
	// A gate that examined nothing passes by never looking.
	if scanned == 0 {
		t.Fatalf("scanned 0 test files: the inventory verdict below would be about the working directory, not the suite")
	}
	total := 0
	for name, n := range got {
		total += n
		want, ok := wantLiveTestInventory[name]
		if !ok {
			t.Errorf("%s carries //go:build xpclive with %d TestLive* functions but is not in wantLiveTestInventory: "+
				"add it (and update the header comment's total)", name, n)
			continue
		}
		if n != want {
			t.Errorf("%s declares %d TestLive* functions, want %d: update wantLiveTestInventory "+
				"and the header comment's total", name, n, want)
		}
	}
	for name := range wantLiveTestInventory {
		if _, ok := got[name]; !ok {
			t.Errorf("wantLiveTestInventory records %s, which is not an xpclive-tagged file in this package: "+
				"it was deleted, renamed, or lost its build tag", name)
		}
	}
	self, err := os.ReadFile("xpclive_gate_test.go")
	if err != nil {
		t.Fatalf("read this file: %v", err)
	}
	claim := fmt.Sprintf("%d TestLive* functions", total)
	if !strings.Contains(string(self), claim) {
		t.Errorf("the header comment in xpclive_gate_test.go does not say %q; the suite has %d "+
			"TestLive* functions across %d files", claim, total, len(got))
	}
	t.Logf("xpclive inventory: %d TestLive* functions across %d tagged files (%d test files scanned)",
		total, len(got), scanned)
}
