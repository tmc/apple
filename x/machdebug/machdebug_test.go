package machdebug_test

import (
	"encoding/binary"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/tmc/apple/x/machdebug"
)

// TestGetTaskAllowGate is the assumption the rest of the suite rests on:
// task_for_pid must succeed, without root, against a child this test signed
// ad hoc with com.apple.security.get-task-allow. If this fails, nothing below
// it is meaningful.
func TestGetTaskAllowGate(t *testing.T) {
	c := startChild(t)
	if os.Geteuid() == 0 {
		t.Fatalf("this test must run unprivileged; euid is 0")
	}
	p, err := machdebug.Attach(c.pid)
	if err != nil {
		t.Fatalf("attach to get-task-allow child (euid %d): %v", os.Geteuid(), err)
	}
	if p.Pid() != c.pid {
		t.Errorf("Pid() = %d, want %d", p.Pid(), c.pid)
	}
	if err := p.Detach(); err != nil {
		t.Errorf("detach: %v", err)
	}
}

// TestAttachUnsignedIsRefused is the mutation control for the gate: the same
// code path against a child without the entitlement must fail. Without it,
// TestGetTaskAllowGate could be passing because task_for_pid works on
// anything, which would say nothing about the entitlement.
func TestAttachUnsignedIsRefused(t *testing.T) {
	c := startStrippedChild(t)
	p, err := machdebug.Attach(c.pid)
	if err == nil {
		p.Detach()
		t.Fatalf("attach to a child signed without get-task-allow succeeded; the gate proves nothing")
	}
	t.Logf("attach refused as expected: %v", err)
}

func TestReadWriteMemory(t *testing.T) {
	c := startChild(t)
	p := attach(t, c.pid)

	const want uint64 = 0x1111222233334444
	var b [8]byte
	if _, err := p.ReadAt(b[:], int64(c.counter)); err != nil {
		t.Fatalf("read counter: %v", err)
	}
	if got := binary.LittleEndian.Uint64(b[:]); got != want {
		t.Fatalf("counter in target = %#x, want %#x", got, want)
	}
	if got := c.value(t, "get"); got != want {
		t.Fatalf("counter per the child = %#x, want %#x", got, want)
	}

	const set uint64 = 0xdeadbeefcafef00d
	binary.LittleEndian.PutUint64(b[:], set)
	if _, err := p.WriteAt(b[:], int64(c.counter)); err != nil {
		t.Fatalf("write counter: %v", err)
	}
	if got := c.value(t, "get"); got != set {
		t.Fatalf("after write the child reports %#x, want %#x", got, set)
	}
}

func TestReadAtUnmapped(t *testing.T) {
	c := startChild(t)
	p := attach(t, c.pid)

	var b [8]byte
	if _, err := p.ReadAt(b[:], 0x10); err == nil {
		t.Fatal("read of the target's null page succeeded")
	}
	// Mutation control: the same call at a mapped address must succeed, so
	// the failure above is about the address and not about ReadAt.
	if _, err := p.ReadAt(b[:], int64(c.counter)); err != nil {
		t.Fatalf("read of a mapped address failed too: %v", err)
	}
}

func TestThreads(t *testing.T) {
	c := startChild(t)
	p := attach(t, c.pid)

	ts, err := p.Threads()
	if err != nil {
		t.Fatalf("threads: %v", err)
	}
	if len(ts) == 0 {
		t.Fatal("no threads")
	}
	var ok int
	for _, th := range ts {
		r, err := th.Regs()
		if err != nil {
			t.Errorf("regs for thread %d: %v", th.Port(), err)
			continue
		}
		if r.PC == 0 && r.SP == 0 {
			t.Errorf("thread %d has an all-zero state", th.Port())
			continue
		}
		ok++
	}
	if ok == 0 {
		t.Fatal("no thread produced usable registers")
	}
}

func TestSetRegsRejectsPCRedirection(t *testing.T) {
	c := startChild(t)
	p := attach(t, c.pid)

	ts, err := p.Threads()
	if err != nil {
		t.Fatalf("threads: %v", err)
	}
	r, err := ts[0].Regs()
	if err != nil {
		t.Fatalf("regs: %v", err)
	}
	r.PC += 4
	if err := ts[0].SetRegs(r); err == nil {
		t.Fatal("SetRegs accepted a redirected PC")
	}
}

// TestOneShotBreakpoint is M2: arm a breakpoint on the child's bump
// function, catch the stop, rewrite x0, resume, and assert the child's answer
// changed. The register write is the whole point — a test that only checked
// that the breakpoint fired would pass with a broken resume path.
func TestOneShotBreakpoint(t *testing.T) {
	c := startChild(t)
	p := attach(t, c.pid)

	bp, err := p.SetOneShot(c.bump)
	if err != nil {
		t.Fatalf("set one-shot at %#x: %v", c.bump, err)
	}
	if bp.Addr != c.bump {
		t.Errorf("bp.Addr = %#x, want %#x", bp.Addr, c.bump)
	}

	// The child blocks on our reply inside bump, so its answer cannot be read
	// until the event is released.
	c.send(t, "bump 5")

	ev, err := p.WaitTimeout(10 * time.Second)
	if err != nil {
		t.Fatalf("wait: %v", err)
	}
	if ev.Kind != machdebug.EventBreakpoint {
		t.Errorf("event kind = %v, want breakpoint", ev.Kind)
	}
	if ev.Addr != c.bump {
		t.Errorf("event addr = %#x, want %#x", ev.Addr, c.bump)
	}
	r, err := ev.Thread.Regs()
	if err != nil {
		t.Fatalf("regs at breakpoint: %v", err)
	}
	if r.PC != c.bump {
		t.Fatalf("pc at breakpoint = %#x, want %#x", r.PC, c.bump)
	}
	if r.X[0] != 5 {
		t.Fatalf("x0 at bump(5) = %#x, want 5; the Go arm64 register ABI is not passing the argument where this test assumes", r.X[0])
	}
	r.X[0] = 0x40
	if err := ev.Thread.SetRegs(r); err != nil {
		t.Fatalf("setregs: %v", err)
	}

	if err := p.Detach(); err != nil {
		t.Fatalf("detach: %v", err)
	}
	if got := c.result(t); got != 0x41 {
		t.Fatalf("child computed bump = %#x, want 0x41 (x0 rewritten to 0x40)", got)
	}

	// The one-shot disarmed itself, so the next call runs unmodified.
	if got := c.value(t, "bump 5"); got != 6 {
		t.Fatalf("after the one-shot fired, bump(5) = %#x, want 6", got)
	}
}

// TestBreakpointPatchIsPrivateToTheTask is the VM_PROT_COPY assertion. Two
// children run the same binary and share its text pages; patching one must
// not change the other.
func TestBreakpointPatchIsPrivateToTheTask(t *testing.T) {
	a := startChild(t)
	b := startChild(t)
	// The two run the same binary, so the page holding bump is the same
	// file-backed page in both, whatever slide each landed at.
	pa := attach(t, a.pid)
	pb := attach(t, b.pid)

	var before [4]byte
	if _, err := pb.ReadAt(before[:], int64(b.bump)); err != nil {
		t.Fatalf("read the other child's instruction: %v", err)
	}

	if _, err := pa.SetOneShot(a.bump); err != nil {
		t.Fatalf("set breakpoint: %v", err)
	}

	// Mutation control: the patched task really did change, so the unchanged
	// reading below is about privacy and not about a write that never landed.
	var patched [4]byte
	if _, err := pa.ReadAt(patched[:], int64(a.bump)); err != nil {
		t.Fatalf("read the patched instruction: %v", err)
	}
	const brk uint32 = 0xD4200000
	if got := binary.LittleEndian.Uint32(patched[:]); got != brk {
		t.Fatalf("patched task holds %#08x, want %#08x", got, brk)
	}

	var after [4]byte
	if _, err := pb.ReadAt(after[:], int64(b.bump)); err != nil {
		t.Fatalf("re-read the other child's instruction: %v", err)
	}
	if after != before {
		t.Fatalf("patching one task changed the other: %#08x became %#08x — VM_PROT_COPY is not in effect",
			binary.LittleEndian.Uint32(before[:]), binary.LittleEndian.Uint32(after[:]))
	}
}

// TestDetachRestores checks the safety contract's core promise: after Detach
// nothing we wrote is still in the target.
func TestDetachRestores(t *testing.T) {
	c := startChild(t)
	p := attach(t, c.pid)

	var orig [4]byte
	if _, err := p.ReadAt(orig[:], int64(c.bump)); err != nil {
		t.Fatalf("read original: %v", err)
	}
	if _, err := p.SetBreakpoint(c.bump); err != nil {
		t.Fatalf("set breakpoint: %v", err)
	}
	var armed [4]byte
	if _, err := p.ReadAt(armed[:], int64(c.bump)); err != nil {
		t.Fatalf("read armed: %v", err)
	}
	if armed == orig {
		t.Fatal("arming the breakpoint changed nothing; the restore assertion below would be vacuous")
	}
	if err := p.Detach(); err != nil {
		t.Fatalf("detach: %v", err)
	}

	// Read back through a fresh attachment: the first Process is gone.
	q := attach(t, c.pid)
	var restored [4]byte
	if _, err := q.ReadAt(restored[:], int64(c.bump)); err != nil {
		t.Fatalf("read after detach: %v", err)
	}
	if restored != orig {
		t.Fatalf("after detach the target holds %#08x, want the original %#08x",
			binary.LittleEndian.Uint32(restored[:]), binary.LittleEndian.Uint32(orig[:]))
	}
	// The child must still work.
	if got := c.value(t, "bump 5"); got != 6 {
		t.Fatalf("after detach bump(5) = %#x, want 6", got)
	}
}

// TestDetachIsIdempotentAndSafeAfterPanic covers the deferred-Detach path a
// consumer relies on when its own code panics.
func TestDetachIsIdempotentAndSafeAfterPanic(t *testing.T) {
	c := startChild(t)

	func() {
		p, err := machdebug.Attach(c.pid)
		if err != nil {
			t.Fatalf("attach: %v", err)
		}
		defer func() {
			if err := p.Detach(); err != nil {
				t.Errorf("detach while panicking: %v", err)
			}
			if err := p.Detach(); err != nil {
				t.Errorf("second detach: %v", err)
			}
			if r := recover(); r == nil {
				t.Error("expected to be recovering from a panic")
			}
		}()
		if _, err := p.SetBreakpoint(c.bump); err != nil {
			t.Fatalf("set breakpoint: %v", err)
		}
		panic("consumer bug")
	}()

	if got := c.value(t, "bump 5"); got != 6 {
		t.Fatalf("after a panicking consumer detached, bump(5) = %#x, want 6", got)
	}
}

// TestOperationsAfterDetach checks that a detached Process reports errors
// rather than touching a task port it no longer owns.
func TestOperationsAfterDetach(t *testing.T) {
	c := startChild(t)
	p := attach(t, c.pid)
	if err := p.Detach(); err != nil {
		t.Fatalf("detach: %v", err)
	}

	var b [8]byte
	for _, tc := range []struct {
		name string
		err  error
	}{
		{"ReadAt", errAt(func() error { _, e := p.ReadAt(b[:], int64(c.counter)); return e })},
		{"WriteAt", errAt(func() error { _, e := p.WriteAt(b[:], int64(c.counter)); return e })},
		{"Threads", errAt(func() error { _, e := p.Threads(); return e })},
		{"SetBreakpoint", errAt(func() error { _, e := p.SetBreakpoint(c.bump); return e })},
		{"Wait", errAt(func() error { _, e := p.Wait(); return e })},
	} {
		if !errors.Is(tc.err, machdebug.ErrDetached) {
			t.Errorf("%s after detach: err = %v, want ErrDetached", tc.name, tc.err)
		}
	}
}

func errAt(f func() error) error { return f() }

// attach attaches to pid and registers a Detach that runs even if the test
// fails partway through.
func attach(t *testing.T, pid int) *machdebug.Process {
	t.Helper()
	p, err := machdebug.Attach(pid)
	if err != nil {
		t.Fatalf("attach to %d: %v", pid, err)
	}
	t.Cleanup(func() { p.Detach() })
	return p
}

// TestPersistentBreakpoint is M3: a breakpoint inside a loop must fire on
// every iteration and the target must still compute the right answer. The
// re-arm goes through a single step over the restored instruction, so a
// breakpoint that fires only once — or a target that dies on the second
// iteration — is exactly what this catches.
func TestPersistentBreakpoint(t *testing.T) {
	c := startChild(t)
	p := attach(t, c.pid)

	if _, err := p.SetBreakpoint(c.bump); err != nil {
		t.Fatalf("set breakpoint: %v", err)
	}
	const n = 5
	c.send(t, "loop 5")

	for i := range n {
		ev, err := p.WaitTimeout(10 * time.Second)
		if err != nil {
			t.Fatalf("wait for hit %d of %d: %v", i+1, n, err)
		}
		if ev.Kind != machdebug.EventBreakpoint || ev.Addr != c.bump {
			t.Fatalf("hit %d: got %v at %#x, want breakpoint at %#x", i+1, ev.Kind, ev.Addr, c.bump)
		}
	}
	if err := p.Detach(); err != nil {
		t.Fatalf("detach: %v", err)
	}

	// sum of bump(0..4) = 1+2+3+4+5
	if got := c.result(t); got != 15 {
		t.Fatalf("child computed loop sum %#x, want 15", got)
	}
}

// TestWatchpoint is M4: a hardware watchpoint on the counter global must trap
// the child's own store to it.
func TestWatchpoint(t *testing.T) {
	c := startChild(t)
	p := attach(t, c.pid)

	w, err := p.SetWatchpoint(c.counter, 8, machdebug.WatchWrite)
	if err != nil {
		t.Fatalf("set watchpoint at %#x: %v", c.counter, err)
	}
	if w.Addr != c.counter || w.Size != 8 {
		t.Errorf("watchpoint = %#x/%d, want %#x/8", w.Addr, w.Size, c.counter)
	}
	c.send(t, "set 0x99")

	ev, err := p.WaitTimeout(10 * time.Second)
	if err != nil {
		t.Fatalf("wait: %v", err)
	}
	if ev.Kind != machdebug.EventWatchpoint {
		t.Fatalf("event kind = %v, want watchpoint", ev.Kind)
	}
	if ev.Addr != c.counter {
		t.Errorf("event addr = %#x, want %#x", ev.Addr, c.counter)
	}
	if err := p.Detach(); err != nil {
		t.Fatalf("detach: %v", err)
	}
	if got := c.result(t); got != 0x99 {
		t.Fatalf("child set counter to %#x, want 0x99", got)
	}
}

// TestWatchpointArgs is the table of inputs SetWatchpoint must refuse.
func TestWatchpointArgs(t *testing.T) {
	c := startChild(t)
	p := attach(t, c.pid)

	for _, tc := range []struct {
		name string
		addr uint64
		size int
		kind machdebug.Watch
		ok   bool
	}{
		{"aligned 8", c.counter, 8, machdebug.WatchWrite, true},
		{"size 3", c.counter, 3, machdebug.WatchWrite, false},
		{"size 16", c.counter, 16, machdebug.WatchWrite, false},
		{"misaligned", c.counter + 1, 8, machdebug.WatchWrite, false},
		{"unknown kind", c.counter, 8, machdebug.Watch(99), false},
	} {
		w, err := p.SetWatchpoint(tc.addr, tc.size, tc.kind)
		if (err == nil) != tc.ok {
			t.Errorf("%s: err = %v, want ok = %v", tc.name, err, tc.ok)
		}
		if err == nil {
			w.Clear()
		}
	}
}
