//go:build darwin

package lldb_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

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
	d.SetAsync(false)
	tgt := d.CreateTarget("")
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
	d.SetAsync(false)

	child := buildChild(t)
	tgt := d.CreateTargetWithArch(child, "")
	if bp := tgt.BreakpointCreateByName("breakpoint_target"); !bp.IsValid() {
		t.Skip("could not set breakpoint on breakpoint_target")
	}

	proc := tgt.LaunchSimple(nil, nil, filepath.Dir(child))
	if !proc.IsValid() {
		t.Skip("could not launch child")
	}
	defer proc.Kill()

	if got := proc.State(); got != lldb.StateStopped {
		t.Fatalf("process state = %v, want stopped", got)
	}

	var hit bool
	for i := uint32(0); i < proc.NumThreads(); i++ {
		th := proc.ThreadAtIndex(i)
		if th.StopReason() != lldb.StopReasonBreakpoint {
			continue
		}
		hit = true
		fr := th.FrameAtIndex(0)
		if name := fr.FunctionName(); !strings.Contains(name, "breakpoint_target") {
			t.Errorf("function name = %q, want to contain breakpoint_target", name)
		}
		reg := fr.FindRegister(selfRegister(tgt.Triple()))
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
	d.SetAsync(false)
	child := buildChild(t)
	tgt := d.CreateTargetWithArch(child, "")
	tgt.BreakpointCreateByName("breakpoint_target")
	proc := tgt.LaunchSimple(nil, nil, filepath.Dir(child))
	if !proc.IsValid() || proc.State() != lldb.StateStopped {
		t.Fatal("child did not stop at breakpoint")
	}
	detached := false
	defer func() {
		if !detached {
			proc.Kill()
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
	if e := proc.Detach(); !e.Success() {
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
