//go:build darwin

package e5rt_test

import (
	"os"
	"path/filepath"
	"testing"
	"unsafe"

	"github.com/tmc/apple/x/ane/e5rt"
	"github.com/tmc/apple/x/ane/mil"
)

// A poisonTail is an E5RT buffer object deliberately larger than the binding
// it serves, whose storage past the logical size carries a known pattern.
//
// The only constructor E5RT exposes for buffer objects is BufferObjectAlloc;
// the caller never controls the mapping, so a guard page after the logical
// end is not possible. Filling the over-allocation instead converts any
// past-the-end engine write into a detectable change at check time rather
// than a fault at write time. It cannot see a write that lands past the
// poison region entirely, and it observes only when someone looks, which
// makes it weaker than a guard page — but it needs no entry point the ABI
// does not already offer.
type poisonTail struct {
	lib   *e5rt.Lib
	obj   uintptr
	data  []byte // whole allocation, including the poison region
	start int    // first poisoned byte
}

// tailBytes is how far past the logical size the poison region runs. It only
// has to exceed any plausible overshoot; 256 bytes covers four padded rows at
// the measured shapes.
const tailBytes = 256

// allocPoisoned allocates nbytes plus the poison tail, zeroes the logical
// region so a program's arithmetic is deterministic, poisons the rest, and
// arranges for the buffer to be released with the test.
func allocPoisoned(t *testing.T, lib *e5rt.Lib, nbytes int) *poisonTail {
	t.Helper()
	if nbytes <= 0 || nbytes > 1<<20 {
		t.Fatalf("logical size %d out of range", nbytes)
	}
	obj, err := lib.BufferObjectAlloc(uintptr(nbytes+tailBytes), 0)
	if err != nil {
		t.Fatalf("allocate %d-byte buffer: %v", nbytes+tailBytes, err)
	}
	ptr, err := lib.BufferObjectGetDataPtr(obj)
	if err != nil {
		lib.BufferObjectRelease(obj)
		t.Fatalf("get buffer address: %v", err)
	}
	data := unsafe.Slice((*byte)(pointerAt(ptr)), nbytes+tailBytes)
	clear(data[:nbytes])
	p := &poisonTail{lib: lib, obj: obj, data: data, start: nbytes}
	t.Cleanup(func() { lib.BufferObjectRelease(p.obj) })
	p.refill()
	return p
}

// refill restores the poison pattern over everything past the logical size.
// The pattern is position-dependent, so a write that shifts the tail rather
// than stomping it still reads as damage.
func (p *poisonTail) refill() {
	for i := p.start; i < len(p.data); i++ {
		p.data[i] = byte(i*31 + 7)
	}
}

// intact reports whether every poisoned byte survived an execution, and on
// failure names the first offset that changed relative to the logical size.
func (p *poisonTail) intact() (bool, int) {
	for i := p.start; i < len(p.data); i++ {
		if want := byte(i*31 + 7); p.data[i] != want {
			return false, i - p.start
		}
	}
	return true, 0
}

// runStateOnce compiles an accumulate-state program over shape [1, C, 1, D],
// binds its state port to buf through the checked path using lay, binds the
// ordinary ports, encodes one operation, executes it with an all-ones update,
// and returns the values the program reports.
func runStateOnce(t *testing.T, lib *e5rt.Lib, dir string, lay e5rt.StateLayout, buf *poisonTail, claimed e5rt.StateLayout) []float32 {
	t.Helper()
	text := mil.GenAccumulateState("kv", [4]int{1, lay.Channels, 1, lay.Dim})
	if err := os.WriteFile(filepath.Join(dir, "model.mil"), []byte(text), 0o644); err != nil {
		t.Fatal(err)
	}
	op := openStateOperation(t, lib, dir)
	t.Cleanup(func() { lib.OperationRelease(op) })

	port, err := lib.BindStatePort(op, "kv", buf.obj, claimed)
	if err != nil {
		t.Fatalf("bind state port: %v", err)
	}
	t.Cleanup(func() { lib.IOPortRelease(port) })
	_, inPtr := bindExamplePort(lib, op, "value", lay.Values()*2, true)
	_, outPtr := bindExamplePort(lib, op, "y", lay.Values()*2, false)

	stream, err := lib.ExecutionStreamCreate()
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { lib.ExecutionStreamRelease(stream) })
	if err := lib.EncodeOperation(stream, op); err != nil {
		t.Fatal(err)
	}
	ones := make([]float32, lay.Values())
	for i := range ones {
		ones[i] = 1
	}
	writeExampleFP16(inPtr, ones)
	if err := lib.ExecuteSync(stream); err != nil {
		t.Fatalf("execute: %v", err)
	}
	return readExampleFP16(outPtr, lay.Values())
}

// TestStatePoisonTail runs a state-writing program against an oversized,
// poisoned buffer and requires the poison to survive.
//
// The layout work found that the engine addresses padded channel rows
// regardless of what the caller allocated, and that a reader sharing the same
// wrong offsets hides the corruption behind a self-consistent round trip.
// This is the standing instrument for that bug class: whatever the engine
// writes behind a correctly sized bind must stop inside StateLayout.Size.
func TestStatePoisonTail(t *testing.T) {
	lib, err := e5rt.Open()
	if err != nil {
		t.Skipf("e5rt unavailable: %v", err)
	}
	lay := e5rt.StateLayout{Channels: 4, Dim: 4}
	if lay.Packed() {
		t.Fatal("this shape must be unpacked for the instrument to mean anything")
	}

	buf := allocPoisoned(t, lib, lay.Size())
	got := runStateOnce(t, lib, t.TempDir(), lay, buf, lay)

	// The run must have done real work before the poison verdict counts: an
	// execution that wrote nothing would pass vacuously. The state started
	// zeroed and the update was all ones, so every reported value is exactly 1.
	for i, v := range got {
		if v != 1 {
			t.Fatalf("reported value %d = %v, want exactly 1; the run did not do its job", i, v)
		}
	}

	if ok, at := buf.intact(); !ok {
		gotBytes := append([]byte(nil), buf.data[buf.start+at:buf.start+at+16]...)
		t.Fatalf("the engine wrote past StateLayout.Size (%d): first changed byte at +%d, got % x",
			lay.Size(), at, gotBytes)
	}
}

// TestStatePoisonTailFires proves the instrument can detect the failure mode
// it exists for. The bind claims a two-channel layout while the compiled
// program has four channels, so the engine writes channels 2 and 3 into the
// region this test declared to be padding. The over-allocation keeps every
// write inside the real allocation, which makes this safe to run: the
// deliberate mismatch is in the claim, never in the memory.
func TestStatePoisonTailFires(t *testing.T) {
	lib, err := e5rt.Open()
	if err != nil {
		t.Skipf("e5rt unavailable: %v", err)
	}
	real := e5rt.StateLayout{Channels: 4, Dim: 4}
	claimed := e5rt.StateLayout{Channels: 2, Dim: 4}
	if claimed.Size() >= real.Size() {
		t.Fatalf("claimed size %d does not undercut the real footprint %d", claimed.Size(), real.Size())
	}

	buf := allocPoisoned(t, lib, claimed.Size())
	runStateOnce(t, lib, t.TempDir(), real, buf, claimed)

	if ok, _ := buf.intact(); ok {
		t.Fatal("the engine stayed inside the under-claimed layout, so the poison instrument cannot fire; " +
			"either the runtime clamps writes to the claimed bind or the footprint model is wrong — " +
			"re-derive before trusting negative results from this instrument")
	}
}
