package machdebug_test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"testing"
	"time"
)

// getTaskAllow is the entitlement that makes a process debuggable by its
// owner without root. It is the ordinary debugger path: no SIP change, no
// privilege. Every test target in this package is signed with it, and if it
// stops working the whole test suite loses its footing rather than quietly
// skipping.
const getTaskAllow = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0"><dict>
<key>com.apple.security.get-task-allow</key><true/>
</dict></plist>
`

// buildChild builds testdata/child and signs it ad-hoc with
// com.apple.security.get-task-allow. The binary is cached for the whole test
// binary's run, since building and signing it costs a second.
var buildChild = sync.OnceValues(func() (string, error) {
	dir, err := os.MkdirTemp(childTempRoot(), "machdebug-child-")
	if err != nil {
		return "", err
	}
	bin := filepath.Join(dir, "child")
	build := exec.Command("go", "build", "-o", bin, "./testdata/child")
	if out, err := build.CombinedOutput(); err != nil {
		return "", fmt.Errorf("go build child: %v: %s", err, out)
	}
	ents := filepath.Join(dir, "child.entitlements")
	if err := os.WriteFile(ents, []byte(getTaskAllow), 0o644); err != nil {
		return "", err
	}
	sign := exec.Command("codesign", "-s", "-", "-f", "--entitlements", ents, bin)
	if out, err := sign.CombinedOutput(); err != nil {
		return "", fmt.Errorf("codesign child: %v: %s", err, out)
	}
	return bin, nil
})

// childTempRoot keeps scratch builds out of /tmp, which macOS sweeps.
func childTempRoot() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	dir := filepath.Join(home, "tmp")
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return ""
	}
	return dir
}

// buildStrippedChild builds the same target signed ad hoc with no
// entitlements at all. It exists to prove the entitlement is what makes
// TestGetTaskAllowGate pass.
var buildStrippedChild = sync.OnceValues(func() (string, error) {
	dir, err := os.MkdirTemp(childTempRoot(), "machdebug-stripped-")
	if err != nil {
		return "", err
	}
	bin := filepath.Join(dir, "child")
	build := exec.Command("go", "build", "-o", bin, "./testdata/child")
	if out, err := build.CombinedOutput(); err != nil {
		return "", fmt.Errorf("go build child: %v: %s", err, out)
	}
	sign := exec.Command("codesign", "-s", "-", "-f", bin)
	if out, err := sign.CombinedOutput(); err != nil {
		return "", fmt.Errorf("codesign child: %v: %s", err, out)
	}
	return bin, nil
})

// child is a running test target: its pid, the addresses it reported, and a
// line channel to ask it what it observes.
type child struct {
	pid     int
	counter uint64 // address of the counter global
	bump    uint64 // address of the bump function

	cmd *exec.Cmd
	in  io.WriteCloser
	out *bufio.Scanner
}

// startChild builds, signs, spawns and handshakes with a test target, and
// registers its cleanup.
func startChild(t *testing.T) *child {
	t.Helper()
	bin, err := buildChild()
	if err != nil {
		t.Fatalf("build child: %v", err)
	}
	return spawn(t, bin)
}

// startStrippedChild spawns a target signed without get-task-allow.
func startStrippedChild(t *testing.T) *child {
	t.Helper()
	bin, err := buildStrippedChild()
	if err != nil {
		t.Fatalf("build stripped child: %v", err)
	}
	return spawn(t, bin)
}

func spawn(t *testing.T, bin string) *child {
	t.Helper()
	cmd := exec.Command(bin)
	in, err := cmd.StdinPipe()
	if err != nil {
		t.Fatalf("stdin pipe: %v", err)
	}
	out, err := cmd.StdoutPipe()
	if err != nil {
		t.Fatalf("stdout pipe: %v", err)
	}
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		t.Fatalf("start child: %v", err)
	}
	c := &child{pid: cmd.Process.Pid, cmd: cmd, in: in, out: bufio.NewScanner(out)}
	t.Cleanup(func() {
		in.Close()
		cmd.Process.Kill()
		cmd.Wait()
	})

	for {
		line := c.line(t)
		f := strings.Fields(line)
		switch {
		case len(f) == 2 && f[0] == "counter":
			c.counter = c.hex(t, f[1])
		case len(f) == 2 && f[0] == "bump":
			c.bump = c.hex(t, f[1])
		case line == "ready":
			if c.counter == 0 || c.bump == 0 {
				t.Fatalf("child handshake incomplete: counter=%#x bump=%#x", c.counter, c.bump)
			}
			return c
		default:
			t.Fatalf("unexpected child handshake line %q", line)
		}
	}
}

func (c *child) hex(t *testing.T, s string) uint64 {
	t.Helper()
	v, err := strconv.ParseUint(strings.TrimPrefix(s, "0x"), 16, 64)
	if err != nil {
		t.Fatalf("parse %q: %v", s, err)
	}
	return v
}

// line reads one line of the child's output, failing the test if the child
// went away or went quiet.
func (c *child) line(t *testing.T) string {
	t.Helper()
	type res struct {
		s  string
		ok bool
	}
	ch := make(chan res, 1)
	go func() {
		ok := c.out.Scan()
		ch <- res{c.out.Text(), ok}
	}()
	select {
	case r := <-ch:
		if !r.ok {
			t.Fatalf("child output ended: %v", c.out.Err())
		}
		return r.s
	case <-time.After(10 * time.Second):
		t.Fatalf("timed out waiting for child output")
		return ""
	}
}

// send writes a command without waiting for the answer. It is what a test
// uses when the child is about to stop inside the command it is being given.
func (c *child) send(t *testing.T, cmd string) {
	t.Helper()
	if _, err := io.WriteString(c.in, cmd+"\n"); err != nil {
		t.Fatalf("write to child: %v", err)
	}
}

// ask sends a command and returns the child's one-line answer.
func (c *child) ask(t *testing.T, cmd string) string {
	t.Helper()
	if _, err := io.WriteString(c.in, cmd+"\n"); err != nil {
		t.Fatalf("write to child: %v", err)
	}
	return c.line(t)
}

// result parses an answer the child already sent, for a command issued with
// send.
func (c *child) result(t *testing.T) uint64 {
	t.Helper()
	f := strings.Fields(c.line(t))
	if len(f) != 2 {
		t.Fatalf("child answered %q, want two fields", strings.Join(f, " "))
	}
	return c.hex(t, f[1])
}

// value asks the child for a "<word> <hex>" answer and returns the number.
func (c *child) value(t *testing.T, cmd string) uint64 {
	t.Helper()
	f := strings.Fields(c.ask(t, cmd))
	if len(f) != 2 {
		t.Fatalf("child answered %q, want two fields", strings.Join(f, " "))
	}
	return c.hex(t, f[1])
}
