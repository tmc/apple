package signpost

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"testing"
	"time"
)

// TestLoadNamesRoundTrip builds a names dylib with signpostnames -dylib,
// loads it, and verifies through log show that both the name and a message
// argument decode. This is the path for internally linked (CGO_ENABLED=0)
// builds, where a .syso's __oslogstring section is dropped, so the test runs
// identically in both build modes.
func TestLoadNamesRoundTrip(t *testing.T) {
	if _, err := exec.LookPath("clang"); err != nil {
		t.Skip("clang not available")
	}
	dylib := filepath.Join(t.TempDir(), "names.dylib")
	out, err := exec.Command("go", "run", "./cmd/signpostnames",
		"-dylib", "-o", dylib, "-extra", "LoadNamesRT", t.TempDir()).CombinedOutput()
	if err != nil {
		t.Fatalf("signpostnames -dylib: %v\n%s", err, out)
	}

	if err := LoadNames(dylib); err != nil {
		t.Fatalf("LoadNames: %v", err)
	}
	if np, fp, dso, ok := lookup("LoadNamesRT", messageFormat); !ok || np == nil || fp == nil || dso == nil {
		t.Fatalf("lookup(LoadNamesRT) not resolved from loaded dylib")
	}

	log := New("com.tmc.apple.signposttest.dylib", PointsOfInterest)
	if !log.Enabled() {
		t.Skip("signposts disabled for this log handle")
	}
	msg := fmt.Sprintf("dylib-roundtrip-%d", time.Now().UnixNano())
	id := log.NewID()
	log.IntervalBeginMessage(id, "LoadNamesRT", msg)
	log.IntervalEndMessage(id, "LoadNamesRT", msg)

	deadline := time.Now().Add(15 * time.Second)
	for {
		out, err := exec.Command("log", "show", "--last", "1m", "--signpost",
			"--predicate", `subsystem == "com.tmc.apple.signposttest.dylib"`).Output()
		if err != nil {
			t.Skipf("log show unavailable: %v", err)
		}
		if bytes.Contains(out, []byte("LoadNamesRT: "+msg)) {
			return
		}
		if time.Now().After(deadline) {
			t.Fatalf("decoded name+message %q not found in log show output:\n%s", "LoadNamesRT: "+msg, out)
		}
		time.Sleep(time.Second)
	}
}
