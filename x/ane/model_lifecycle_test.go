//go:build darwin

package ane

import (
	"os"
	"strings"
	"testing"

	"github.com/tmc/apple/x/ane/mil"
)

// compileIdentity compiles a small identity model with the given channel count.
func compileIdentity(t *testing.T, c *Client, ch, spatial int) (*Model, error) {
	t.Helper()
	blob, err := mil.BuildIdentityWeightBlob(ch)
	if err != nil {
		t.Fatal(err)
	}
	return c.Compile(CompileOptions{
		ModelType:  ModelTypeMIL,
		MILText:    []byte(mil.GenIdentity(ch, spatial)),
		WeightBlob: blob,
	})
}

// TestCompileCloseDoesNotAccumulate compiles and closes far more models than the
// device holds live at once. Before compile ran under an autorelease pool, the
// ANE objects survived Close and held their program instances, so the sixteenth
// compile in a process failed with 0x50004 no matter how promptly callers closed.
func TestCompileCloseDoesNotAccumulate(t *testing.T) {
	c := openOrSkip(t)
	defer c.Close()

	const n = 40 // comfortably past the device's ~15 live program instances
	for i := range n {
		m, err := compileIdentity(t, c, 1+i%8, 1+i)
		if err != nil {
			t.Fatalf("compile %d of %d: %v", i+1, n, err)
		}
		if err := m.Close(); err != nil {
			t.Fatalf("close %d of %d: %v", i+1, n, err)
		}
	}
}

// TestCloseRemovesStagingDir checks that the directory compile writes for the
// Espresso IR translator lives exactly as long as the model does.
func TestCloseRemovesStagingDir(t *testing.T) {
	c := openOrSkip(t)
	defer c.Close()

	m, err := compileIdentity(t, c, 4, 1)
	if err != nil {
		t.Fatal(err)
	}
	if m.tmpDir == "" {
		t.Fatal("model has no staging directory")
	}
	if _, err := os.Stat(m.tmpDir); err != nil {
		t.Errorf("staging dir absent while the model is live: %v", err)
	}
	dir := m.tmpDir
	if err := m.Close(); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		t.Errorf("staging dir survived Close: %v", err)
	}
}

// TestFailedCompileRemovesStagingDir checks the path that orphaned directories:
// compile writes the staging directory before the load that fails, and a failed
// load returns no Model, so nothing downstream would ever remove it.
func TestFailedCompileRemovesStagingDir(t *testing.T) {
	c := openOrSkip(t)
	defer c.Close()

	// Hold models open until a load fails, then count what compile left behind.
	before := countStagingDirs(t)
	var open []*Model
	defer func() {
		for _, m := range open {
			m.Close()
		}
	}()
	for i := range 64 {
		m, err := compileIdentity(t, c, 1+i%8, 1+i)
		if err != nil {
			if got := countStagingDirs(t) - before; got != len(open) {
				t.Errorf("after a failed compile: %d staging dirs, want %d (one per live model)", got, len(open))
			}
			return
		}
		open = append(open, m)
	}
	t.Skip("device accepted 64 live models; no failed load to observe")
}

// countStagingDirs counts compile staging directories in the temp directory.
// They are named for the model's hex identifier: three underscore-separated
// SHA-256 digests.
func countStagingDirs(t *testing.T) int {
	t.Helper()
	entries, err := os.ReadDir(os.TempDir())
	if err != nil {
		t.Fatal(err)
	}
	n := 0
	for _, e := range entries {
		if e.IsDir() && isStagingDirName(e.Name()) {
			n++
		}
	}
	return n
}

func isStagingDirName(name string) bool {
	parts := strings.Split(name, "_")
	if len(parts) != 3 {
		return false
	}
	for _, p := range parts {
		if len(p) != 64 || strings.TrimLeft(p, "0123456789ABCDEF") != "" {
			return false
		}
	}
	return true
}
