//go:build darwin

package powersample

import (
	"os"
	"os/exec"
	"testing"
	"time"
)

func burnCPU(d time.Duration) {
	deadline := time.Now().Add(d)
	x := 1.0
	for time.Now().Before(deadline) {
		x = x*1.0000001 + 1
	}
	_ = x
}

// TestLiveMeter runs the default (IOReport) backend for ~1s. It needs no
// privileges on Apple silicon.
func TestLiveMeter(t *testing.T) {
	m, err := Start(250 * time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	burnCPU(1 * time.Second)
	r, err := m.Stop()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("samples=%d duration=%v energy: cpu=%.3fJ gpu=%.3fJ ane=%.3fJ combined=%.3fJ",
		r.Samples, r.Duration, r.Energy.CPU, r.Energy.GPU, r.Energy.ANE, r.Energy.Combined)
	if r.Samples == 0 {
		t.Fatal("no samples — IOReport Energy Model channels missing on this machine?")
	}
	if r.Energy.CPU <= 0 {
		t.Errorf("CPU energy = %v J, want > 0 under a CPU burn", r.Energy.CPU)
	}
}

// TestLivePowermetricsAgreement runs both backends over the same CPU burn
// and demands the CPU rails agree within a factor of two — a cross-check
// that the unprivileged IOReport numbers are the same physics powermetrics
// reports, not merely nonzero. Needs root or passwordless sudo for the
// powermetrics side; skips honestly otherwise.
func TestLivePowermetricsAgreement(t *testing.T) {
	if os.Geteuid() != 0 {
		if err := exec.Command("sudo", "-n", "-v").Run(); err != nil {
			t.Skip("needs root or passwordless sudo for the powermetrics arm; run: sudo go test -run TestLivePowermetricsAgreement ./x/powersample")
		}
	}
	ior, err := startIOReport(500 * time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	pm, err := startPowermetrics(500 * time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	burnCPU(3 * time.Second)
	ri, err := ior.stop()
	if err != nil {
		t.Fatal(err)
	}
	rp, err := pm.stop()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("ioreport:     cpu=%.3fJ over %v (%d samples)", ri.Energy.CPU, ri.Duration, ri.Samples)
	t.Logf("powermetrics: cpu=%.3fJ over %v (%d samples)", rp.Energy.CPU, rp.Duration, rp.Samples)
	if rp.Samples == 0 {
		t.Fatal("powermetrics arm parsed 0 samples")
	}
	// Compare average watts, not joules: the two windows differ slightly.
	wi, wp := ri.Average.CPU, rp.Average.CPU
	if wi <= 0 || wp <= 0 {
		t.Fatalf("nonpositive CPU watts: ioreport %.3f, powermetrics %.3f", wi, wp)
	}
	if ratio := wi / wp; ratio < 0.5 || ratio > 2.0 {
		t.Errorf("CPU watts disagree: ioreport %.3f W vs powermetrics %.3f W (ratio %.2f, want within 2x)", wi, wp, ratio)
	}
}
