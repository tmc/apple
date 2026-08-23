//go:build darwin

package e5rt_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tmc/apple/x/ane/e5rt"
	"github.com/tmc/apple/x/ane/mil"
)

// TestStateLayout checks the arithmetic against the shapes the layout was
// measured at. The values are the measured ones, not a restatement of the
// formula: a test that recomputed round_up(dim*2, 64) would agree with a
// RowBytes that had the rule wrong.
func TestStateLayout(t *testing.T) {
	tests := []struct {
		name           string
		lay            e5rt.StateLayout
		rowBytes, size int
		packed         bool
	}{
		{"4x4", e5rt.StateLayout{Channels: 4, Dim: 4}, 64, 256, false},
		{"4x32", e5rt.StateLayout{Channels: 4, Dim: 32}, 64, 256, true},
		{"4x40", e5rt.StateLayout{Channels: 4, Dim: 40}, 128, 512, false},
		{"2x64", e5rt.StateLayout{Channels: 2, Dim: 64}, 128, 256, true},
		{"8x2", e5rt.StateLayout{Channels: 8, Dim: 2}, 64, 512, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.lay.RowBytes(); got != tt.rowBytes {
				t.Errorf("RowBytes() = %d, want %d", got, tt.rowBytes)
			}
			if got := tt.lay.Size(); got != tt.size {
				t.Errorf("Size() = %d, want %d", got, tt.size)
			}
			if got := tt.lay.Packed(); got != tt.packed {
				t.Errorf("Packed() = %v, want %v", got, tt.packed)
			}
			// The packed size is the mistake this type exists to prevent, so
			// check the two agree only where they are supposed to.
			if packed := tt.lay.Values() * 2; (packed == tt.lay.Size()) != tt.packed {
				t.Errorf("packed size %d vs Size() %d disagrees with Packed() = %v", packed, tt.lay.Size(), tt.packed)
			}
		})
	}
}

func TestStateLayoutOffset(t *testing.T) {
	lay := e5rt.StateLayout{Channels: 4, Dim: 4}
	// Measured: channel c element i lives at slot c*32+i, that is byte c*64+i*2.
	for c := range 4 {
		for i := range 4 {
			if got, want := lay.Offset(c, i), c*64+i*2; got != want {
				t.Errorf("Offset(%d, %d) = %d, want %d", c, i, got, want)
			}
		}
	}
	// Every element must fall inside the allocation the same layout demands.
	if got := lay.Offset(3, 3) + 2; got > lay.Size() {
		t.Errorf("last element ends at %d, past Size() = %d", got, lay.Size())
	}
	for _, bad := range [][2]int{{-1, 0}, {4, 0}, {0, -1}, {0, 4}} {
		func() {
			defer func() {
				if recover() == nil {
					t.Errorf("Offset(%d, %d) did not panic", bad[0], bad[1])
				}
			}()
			lay.Offset(bad[0], bad[1])
		}()
	}
}

// TestBufferObjectGetSize records what the newly bound entry point returns.
// It reports the requested size unrounded, which is what makes it usable for
// checking a caller's intent rather than the allocation granule.
func TestBufferObjectGetSize(t *testing.T) {
	lib, err := e5rt.Open()
	if err != nil {
		t.Skipf("e5rt unavailable: %v", err)
	}
	for _, want := range []uintptr{1, 32, 63, 64, 65, 256, 320, 4096} {
		buf, err := lib.BufferObjectAlloc(want, 0)
		if err != nil {
			t.Fatalf("BufferObjectAlloc(%d): %v", want, err)
		}
		got, err := lib.BufferObjectGetSize(buf)
		if err != nil {
			t.Fatalf("BufferObjectGetSize after alloc(%d): %v", want, err)
		}
		if got != want {
			t.Errorf("BufferObjectGetSize after alloc(%d) = %d", want, got)
		}
		if err := lib.BufferObjectRelease(buf); err != nil {
			t.Fatal(err)
		}
	}
}

// TestBindStatePort exercises both polarities on one real operation: a buffer
// sized for the packed element count must be refused, and a buffer sized by the
// layout must bind. Without the second arm the refusal would be consistent with
// BindStatePort refusing everything.
func TestBindStatePort(t *testing.T) {
	lib, err := e5rt.Open()
	if err != nil {
		t.Skipf("e5rt unavailable: %v", err)
	}
	lay := e5rt.StateLayout{Channels: 4, Dim: 4}
	if lay.Packed() {
		t.Fatal("this shape must be unpacked for the refusal arm to mean anything")
	}

	dir := t.TempDir()
	text := mil.GenAccumulateState("kv", [4]int{1, lay.Channels, 1, lay.Dim})
	if err := os.WriteFile(filepath.Join(dir, "model.mil"), []byte(text), 0o644); err != nil {
		t.Fatal(err)
	}
	op := openStateOperation(t, lib, dir)
	defer lib.OperationRelease(op)

	// Negative: the packed size, which is what a caller reading the MIL text
	// would compute, and which binds without complaint through the raw calls.
	small, err := lib.BufferObjectAlloc(uintptr(lay.Values()*2), 0)
	if err != nil {
		t.Fatal(err)
	}
	defer lib.BufferObjectRelease(small)
	_, err = lib.BindStatePort(op, "kv", small, lay)
	if err == nil {
		t.Fatalf("binding a %d-byte buffer to a state needing %d bytes was accepted", lay.Values()*2, lay.Size())
	}
	for _, want := range []string{"256", "32"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("refusal %q does not mention %s", err, want)
		}
	}

	// The refusal is only worth having if the raw calls accept the same
	// buffer, so bind it through them and require success. This is the
	// mutation control: if IOPortBindBufferObject had started rejecting
	// undersized buffers on its own, BindStatePort would be redundant and
	// the arm above would prove nothing about it.
	//
	// Binding does not execute, so nothing writes to the undersized buffer
	// here. Executing with it bound is what corrupts memory, which is the
	// whole point.
	raw, err := lib.OperationRetainInoutPort(op, "kv")
	if err != nil {
		t.Fatal(err)
	}
	if err := lib.IOPortBindBufferObject(raw, small); err != nil {
		t.Errorf("the raw bind refused a %d-byte buffer: %v; BindStatePort's check may now be redundant",
			lay.Values()*2, err)
	}
	if err := lib.IOPortRelease(raw); err != nil {
		t.Fatal(err)
	}

	// Positive: the size the layout asks for, on the same operation, so the
	// refusal above is about the size and not about the call being broken.
	big, err := lib.BufferObjectAlloc(uintptr(lay.Size()), 0)
	if err != nil {
		t.Fatal(err)
	}
	defer lib.BufferObjectRelease(big)
	port, err := lib.BindStatePort(op, "kv", big, lay)
	if err != nil {
		t.Fatalf("binding a correctly sized buffer failed: %v", err)
	}
	if err := lib.IOPortRelease(port); err != nil {
		t.Fatal(err)
	}

	// A name the program does not declare must still be refused, and must not
	// be refused by the size check reached first.
	if _, err := lib.BindStatePort(op, "no_such_state", big, lay); err == nil {
		t.Fatal("binding an undeclared state port was accepted")
	} else if !strings.Contains(err.Error(), "retain state port") {
		t.Errorf("undeclared port refused by the wrong check: %v", err)
	}
}
