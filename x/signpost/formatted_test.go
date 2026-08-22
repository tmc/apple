package signpost

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

// TestFormattedRoundTrip drives the full formatted-emission chain: the
// signpostnames scanner finds the name and format literals of an f-variant
// call in Go source, pools them into a dylib, and after LoadNames an Eventf
// with typed arguments decodes in log show with the values rendered into the
// format. Runs identically under CGO_ENABLED=0 and 1.
func TestFormattedRoundTrip(t *testing.T) {
	if _, err := exec.LookPath("clang"); err != nil {
		t.Skip("clang not available")
	}
	const name = "FmtRT"
	const format = "bytes=%ld dur=%f note=%{public}s"

	// A scratch source file carrying the literals, so the test exercises the
	// scanner rather than -extra.
	scratch := t.TempDir()
	src := fmt.Sprintf("package p\n\nfunc f() {\n\tl.Eventf(id, %q, %q, n, d, s)\n}\n", name, format)
	if err := os.WriteFile(filepath.Join(scratch, "calls.go"), []byte(src), 0o644); err != nil {
		t.Fatal(err)
	}
	dylib := filepath.Join(t.TempDir(), "names.dylib")
	out, err := exec.Command("go", "run", "./cmd/signpostnames", "-dylib", "-o", dylib, scratch).CombinedOutput()
	if err != nil {
		t.Fatalf("signpostnames -dylib: %v\n%s", err, out)
	}
	if err := LoadNames(dylib); err != nil {
		t.Fatalf("LoadNames: %v", err)
	}
	if _, fp, _, ok := lookup(name, format); !ok || fp == nil {
		t.Fatalf("scanner did not pool name %q and format %q", name, format)
	}

	log := New("com.tmc.apple.signposttest.fmt", PointsOfInterest)
	if !log.Enabled() {
		t.Skip("signposts disabled for this log handle")
	}
	marker := fmt.Sprintf("fmt-rt-%d", time.Now().UnixNano())
	log.Eventf(log.NewID(), name, format, int64(4096), 0.25, marker)

	// log show renders the arguments into the format string; %f prints with
	// C's default six decimal places.
	want := "FmtRT: bytes=4096 dur=0.250000 note=" + marker
	deadline := time.Now().Add(15 * time.Second)
	for {
		out, err := exec.Command("log", "show", "--last", "1m", "--signpost",
			"--predicate", `subsystem == "com.tmc.apple.signposttest.fmt"`).Output()
		if err != nil {
			t.Skipf("log show unavailable: %v", err)
		}
		if bytes.Contains(out, []byte(want)) {
			return
		}
		if time.Now().After(deadline) {
			t.Fatalf("rendered message %q not found in log show output:\n%s", want, out)
		}
		time.Sleep(time.Second)
	}
}
