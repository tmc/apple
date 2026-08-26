package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/tmc/apple/x/experiment"
)

// The tests exercise the built binary end to end, including child-process
// arms, because the arena's trust boundary is a process boundary.
var binPath string

func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "mlxperf-bin")
	if err != nil {
		panic(err)
	}
	defer os.RemoveAll(dir)
	binPath = filepath.Join(dir, "mlxperf")
	out, err := exec.Command("go", "build", "-o", binPath, ".").CombinedOutput()
	if err != nil {
		panic(string(out) + err.Error())
	}
	os.Exit(m.Run())
}

func mlxperf(t *testing.T, args ...string) (int, string) {
	t.Helper()
	var stdout, stderr bytes.Buffer
	cmd := exec.Command(binPath, args...)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	code := 0
	if exit, ok := err.(*exec.ExitError); ok {
		code = exit.ExitCode()
	} else if err != nil {
		t.Fatal(err)
	}
	return code, stdout.String() + stderr.String()
}

func echoDigest(t testing.TB, s string) string {
	t.Helper()
	sum := sha256.Sum256([]byte(s))
	return hex.EncodeToString(sum[:])
}

// Slice 0's done gate: baseline and candidate workers execute, an output
// digest is validated, one canonical receipt is written, and compare shows it.
func TestArenaRunWritesReceiptAndCompareShowsIt(t *testing.T) {
	d := t.TempDir()
	receipt := filepath.Join(d, "receipt.json")
	digest := echoDigest(t, "ok\n")
	code, out := mlxperf(t, "arena", "run",
		"-workload", "matmul/test",
		"-work-units", "64",
		"-rounds", "2",
		"-baseline-cmd", "echo ok",
		"-candidate-cmd", "echo ok",
		"-expect", digest,
		"-out", receipt,
	)
	if code != 0 {
		t.Fatalf("run exited %d: %s", code, out)
	}
	b, err := os.ReadFile(receipt)
	if err != nil {
		t.Fatal(err)
	}
	rcpt, err := experiment.ParseReceipt(b)
	if err != nil {
		t.Fatalf("written receipt does not verify: %v", err)
	}
	if rcpt.Verdict != experiment.VerdictNoWinner {
		t.Fatalf("identical workers verdict %q", rcpt.Verdict)
	}
	if !strings.Contains(out, "verdict "+experiment.VerdictNoWinner) {
		t.Fatalf("summary missing from output: %s", out)
	}

	// A second identical receipt compares cleanly.
	code, out = mlxperf(t, "arena", "compare", receipt, receipt)
	if code != 0 || strings.Contains(out, "refusal") {
		t.Fatalf("compare refused: %d %s", code, out)
	}
}

// A planted wrong result must refuse instead of producing a timing verdict.
func TestArenaRunRefusesWrongOutputDigest(t *testing.T) {
	d := t.TempDir()
	receipt := filepath.Join(d, "receipt.json")
	expect := echoDigest(t, "ok\n")
	code, out := mlxperf(t, "arena", "run",
		"-workload", "matmul/test",
		"-work-units", "8",
		"-rounds", "1",
		"-baseline-cmd", "echo ok",
		"-candidate-cmd", "echo wrong",
		"-expect", expect,
		"-out", receipt,
	)
	if code == 0 {
		t.Fatalf("wrong output accepted: %s", out)
	}
	b, _ := os.ReadFile(receipt)
	rcpt, err := experiment.ParseReceipt(b)
	if err != nil {
		t.Fatal(err)
	}
	if !rcpt.HasRefusal(experiment.RefusalCorrectnessFailed) {
		t.Fatalf("correctness refusal missing: %+v", rcpt.Refusals)
	}
}

// A changed workload identity refuses comparison.
func TestArenaCompareRefusesIdentityDrift(t *testing.T) {
	d := t.TempDir()
	var receipts []string
	for _, id := range []string{"matmul/a", "matmul/b"} {
		p := filepath.Join(d, strings.ReplaceAll(id, "/", "_")+".json")
		code, out := mlxperf(t, "arena", "run",
			"-workload", id,
			"-rounds", "1",
			"-baseline-cmd", "true",
			"-candidate-cmd", "true",
			"-out", p,
		)
		if code != 0 && !strings.Contains(out, "verdict") {
			t.Fatalf("run %s failed: %s", id, out)
		}
		receipts = append(receipts, p)
	}
	code, out := mlxperf(t, "arena", "compare", receipts[0], receipts[1])
	if code == 0 {
		t.Fatal("cross-workload compare accepted")
	}
	if !strings.Contains(out, experiment.RefusalIdentityDrift) {
		t.Fatalf("identity drift not reported: %s", out)
	}
}

func TestArenaRunUsageErrors(t *testing.T) {
	if code, _ := mlxperf(t, "arena", "run"); code != 2 {
		t.Fatalf("missing flags exited %d", code)
	}
	if code, _ := mlxperf(t, "arena", "bogus"); code != 2 {
		t.Fatalf("unknown subcommand exited %d", code)
	}
	if code, _ := mlxperf(t); code != 2 {
		t.Fatalf("no subcommand exited %d", code)
	}
}
