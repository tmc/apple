//go:build darwin

package lldb_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
	"unsafe"

	"github.com/ebitengine/purego"

	"github.com/tmc/apple/x/lldb"
)

// TestCreate exercises the core call machinery: creating a debugger and target
// (both return SB values through the indirect-result buffer) and reading a
// string-returning method. It needs no debug privileges.
func TestCreate(t *testing.T) {
	d, err := lldb.Create()
	if err != nil {
		t.Skipf("lldb unavailable: %v", err)
	}
	defer d.Close()
	d.SetAsync(false)
	tgt := d.CreateTarget("")
	defer tgt.Close()
	_ = tgt.Triple() // empty target: value unspecified, must not crash
}

// TestLaunchAndInspect launches a child under the debugger, stops it at a
// breakpoint, and reads a register and the frame's function name. Launching a
// child needs no root, so this validates the breakpoint/thread/frame/register
// path end to end.
func TestLaunchAndInspect(t *testing.T) {
	d, err := lldb.Create()
	if err != nil {
		t.Skipf("lldb unavailable: %v", err)
	}
	defer d.Close()
	d.SetAsync(false)

	child := buildChild(t)
	tgt := d.CreateTargetWithArch(child, "")
	defer tgt.Close()
	bp := tgt.BreakpointCreateByName("breakpoint_target")
	defer bp.Close()
	if !bp.IsValid() {
		t.Skip("could not set breakpoint on breakpoint_target")
	}

	proc := tgt.LaunchSimple(nil, nil, filepath.Dir(child))
	defer proc.Close()
	if !proc.IsValid() {
		t.Skip("could not launch child")
	}
	defer func() { proc.Kill().Close() }()

	if got := proc.State(); got != lldb.StateStopped {
		t.Fatalf("process state = %v, want stopped", got)
	}

	var hit bool
	for i := uint32(0); i < proc.NumThreads(); i++ {
		th := proc.ThreadAtIndex(i)
		defer th.Close()
		if th.StopReason() != lldb.StopReasonBreakpoint {
			continue
		}
		hit = true
		fr := th.FrameAtIndex(0)
		defer fr.Close()
		if name := fr.FunctionName(); !strings.Contains(name, "breakpoint_target") {
			t.Errorf("function name = %q, want to contain breakpoint_target", name)
		}
		reg := fr.FindRegister(selfRegister(tgt.Triple()))
		defer reg.Close()
		if v := reg.Value(); v == "" {
			t.Errorf("register value is empty")
		}
	}
	if !hit {
		t.Fatal("no thread stopped at the breakpoint")
	}
}

func selfRegister(triple string) string {
	if strings.Contains(triple, "x86_64") {
		return "rdi"
	}
	return "x0"
}

// TestStopAndDetach checks shutdown without attaching to a system process.
func TestStopAndDetach(t *testing.T) {
	d, err := lldb.Create()
	if err != nil {
		t.Fatal(err)
	}
	defer d.Close()
	d.SetAsync(false)
	child := buildChild(t)
	tgt := d.CreateTargetWithArch(child, "")
	defer tgt.Close()
	bp := tgt.BreakpointCreateByName("breakpoint_target")
	defer bp.Close()
	proc := tgt.LaunchSimple(nil, nil, filepath.Dir(child))
	defer proc.Close()
	if !proc.IsValid() || proc.State() != lldb.StateStopped {
		t.Fatal("child did not stop at breakpoint")
	}
	detached := false
	defer func() {
		if !detached {
			proc.Kill().Close()
		}
	}()
	done := make(chan lldb.Error, 1)
	go func() { done <- proc.Continue() }()
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()
	timeout := time.After(2 * time.Second)
wait:
	for {
		select {
		case e := <-done:
			defer e.Close()
			if !e.Success() {
				t.Fatal(e.String())
			}
			break wait
		case <-ticker.C:
			proc.SendAsyncInterrupt()
		case <-timeout:
			t.Fatal("continue did not return after interrupt")
		}
	}
	if got := proc.State(); got != lldb.StateStopped {
		t.Fatalf("state after interrupt = %v, want stopped", got)
	}
	e := proc.Detach()
	defer e.Close()
	if !e.Success() {
		t.Fatal(e.String())
	}
	detached = true
	if got := proc.State(); got != lldb.StateDetached && got != lldb.StateExited {
		t.Fatalf("state = %v, want detached or exited", got)
	}
	// The detached child exits after its one-second sleep.
	time.Sleep(1100 * time.Millisecond)
}

func buildChild(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	src := filepath.Join(dir, "child.c")
	if err := os.WriteFile(src, []byte("#include <unistd.h>\n__attribute__((noinline)) void breakpoint_target(void) { sleep(1); }\nint main(void) { breakpoint_target(); return 0; }\n"), 0600); err != nil {
		t.Fatal(err)
	}
	child := filepath.Join(dir, "child")
	if out, err := exec.Command("clang", "-g", "-O0", "-o", child, src).CombinedOutput(); err != nil {
		t.Fatalf("build child: %v\n%s", err, out)
	}
	return child
}

// TestManyBreakpoints measures live native allocations, excluding Go's heap and
// malloc's reserved pages. After warmup, releasing each hit's handles should
// prevent live allocations from growing with the number of stops.
func TestManyBreakpoints(t *testing.T) {
	d, err := lldb.Create()
	if err != nil {
		t.Fatal(err)
	}
	defer d.Close()
	d.SetAsync(false)
	dir := t.TempDir()
	src := filepath.Join(dir, "loop.c")
	if err := os.WriteFile(src, []byte(`
__attribute__((noinline)) int breakpoint_target(int i) { return i + 1; }
int main(void) {
    volatile int sum = 0;
    for (int i = 0; i < 1000; i++) sum += breakpoint_target(i);
    return 0;
}
`), 0600); err != nil {
		t.Fatal(err)
	}
	child := filepath.Join(dir, "loop")
	if out, err := exec.Command("clang", "-g", "-O0", "-o", child, src).CombinedOutput(); err != nil {
		t.Fatalf("build child: %v\n%s", err, out)
	}
	target := d.CreateTargetWithArch(child, "")
	defer target.Close()
	bp := target.BreakpointCreateByName("breakpoint_target")
	defer bp.Close()
	if !bp.IsValid() {
		t.Fatal("invalid breakpoint")
	}
	proc := target.LaunchSimple(nil, nil, dir)
	defer proc.Close()
	if !proc.IsValid() {
		t.Fatal("invalid process")
	}
	defer func() { proc.Kill().Close() }()
	sample := nativeMemorySampler(t)
	regName := selfRegister(target.Triple())
	var baseline uint64
	for hit := 0; hit < 1000; hit++ {
		inspectHit(t, proc, regName)
		e := proc.Continue()
		ok, message := e.Success(), e.String()
		e.Close()
		if !ok {
			t.Fatalf("continue after hit %d: %s", hit, message)
		}
		if hit == 199 {
			baseline = sample()
		}
		if hit == 599 || hit == 998 {
			live := sample()
			growth := int64(live) - int64(baseline)
			t.Logf("hit %d: native live bytes %d, growth %+d", hit+1, live, growth)
			// Allow allocator and debugger cache noise, but not retained frames
			// and register objects from hundreds of earlier stops.
			if growth > 1<<20 {
				t.Fatalf("native allocations grew %d bytes after warmup", growth)
			}
		}
	}
	if got := proc.State(); got != lldb.StateExited {
		t.Fatalf("state = %v, want exited", got)
	}
}

func inspectHit(t *testing.T, proc lldb.Process, regName string) {
	t.Helper()
	if got := proc.State(); got != lldb.StateStopped {
		t.Fatalf("state = %v, want stopped", got)
	}
	thread := proc.ThreadAtIndex(0)
	defer thread.Close()
	if thread.StopReason() != lldb.StopReasonBreakpoint {
		t.Fatal("child did not hit breakpoint")
	}
	frame := thread.FrameAtIndex(0)
	defer frame.Close()
	if !strings.Contains(frame.FunctionName(), "breakpoint_target") {
		t.Fatal("unexpected stopped frame")
	}
	for i := 0; i < 16; i++ {
		reg := frame.FindRegister(regName)
		if !reg.IsValid() || reg.Value() == "" {
			reg.Close()
			t.Fatal("invalid register")
		}
		alias := reg
		reg.Close()
		alias.Close()
	}
	// Reacquiring handles after closing other aliases must remain safe.
	alias := frame
	frame.Close()
	alias.Close()
}

func nativeMemorySampler(t *testing.T) func() uint64 {
	t.Helper()
	h, err := purego.Dlopen("/usr/lib/libSystem.B.dylib", purego.RTLD_NOW)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { purego.Dlclose(h) })
	// malloc_statistics_t from <malloc/malloc.h>. A nil zone sums all zones.
	type stats struct {
		blocks                     uint32
		inUse, maxInUse, allocated uint64
	}
	var statistics func(unsafe.Pointer, *stats)
	purego.RegisterLibFunc(&statistics, h, "malloc_zone_statistics")
	return func() uint64 {
		var s stats
		statistics(nil, &s)
		return s.inUse
	}
}

func TestExpressionClose(t *testing.T) {
	d, err := lldb.Create()
	if err != nil {
		t.Fatal(err)
	}
	defer d.Close()
	target := d.CreateTargetWithArch("/bin/sleep", "")
	defer target.Close()
	for i := 0; i < 10; i++ {
		value := target.EvaluateExpression("(int)42")
		valid, got := value.IsValid(), value.ValueAsUnsigned(0)
		alias := value
		value.Close()
		alias.Close()
		if !valid || got != 42 {
			t.Fatalf("expression = %d (valid %t), want 42", got, valid)
		}
	}
}
