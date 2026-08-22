//go:build darwin

package ane

import (
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
