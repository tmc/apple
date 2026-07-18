package main

import (
	"errors"
	"io"
	"os"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/tmc/apple/rdma"
)

func TestPrintRKeyCapability(t *testing.T) {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w
	printResult(result{Mode: "rdma-rkey-capability", RKeyCapability: &rkeyCapabilityResult{
		Device: "rdma_en3", Outcome: "zero", Attempts: 1, NoQP: true, NoRTR: true, NoData: true,
		Addr: "0x1000", LKey: "0x101", RKey: "0x0",
	}})
	w.Close()
	os.Stdout = old
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"outcome=zero", "addr=0x1000", "lkey=0x101", "rkey=0x0", "no_qp=true"} {
		if !strings.Contains(string(out), want) {
			t.Fatalf("printResult output %q missing %q", out, want)
		}
	}
}

func TestPrintPDLifecycle(t *testing.T) {
	old := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w
	printResult(result{Mode: "rdma-pd-lifecycle", PDLifecycle: &pdLifecycleResult{
		Device: "rdma_en2", Mode: "reclaim", Outcome: "reclaimed", Cycles: 1, AllocPerCycle: 11, RoundsDone: 1, Allocations: 11, Deallocations: 11,
		NoMR: true, NoQP: true, NoRTR: true, NoData: true,
	}})
	w.Close()
	os.Stdout = old
	out, err := io.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	for _, want := range []string{"mode=reclaim", "outcome=reclaimed", "allocations=11", "deallocations=11", "no_mr=true"} {
		if !strings.Contains(string(out), want) {
			t.Fatalf("printResult output %q missing %q", out, want)
		}
	}
}

func TestRequireRCCapabilityProbeAllowed(t *testing.T) {
	t.Setenv(rcCapabilityConfirmEnv, "")
	if err := requireRCCapabilityProbeAllowed(false); err == nil {
		t.Fatal("requireRCCapabilityProbeAllowed(false) succeeded")
	}
	if err := requireRCCapabilityProbeAllowed(true); err == nil {
		t.Fatal("requireRCCapabilityProbeAllowed(true) succeeded without confirmation")
	}
	t.Setenv(rcCapabilityConfirmEnv, rcCapabilityConfirmValue)
	if err := requireRCCapabilityProbeAllowed(true); err != nil {
		t.Fatalf("requireRCCapabilityProbeAllowed(true) = %v", err)
	}
}

func TestValidateRCCapabilityTimeout(t *testing.T) {
	for _, timeout := range []time.Duration{-time.Second, 0} {
		if err := validateRCCapabilityTimeout(timeout); err == nil {
			t.Fatalf("validateRCCapabilityTimeout(%s) succeeded", timeout)
		}
	}
	if err := validateRCCapabilityTimeout(time.Second); err != nil {
		t.Fatalf("validateRCCapabilityTimeout(time.Second) = %v", err)
	}
}

func TestRequireRKeyCapabilityProbeAllowed(t *testing.T) {
	t.Setenv(rkeyCapabilityConfirmEnv, "")
	if err := requireRKeyCapabilityProbeAllowed(false); err == nil {
		t.Fatal("requireRKeyCapabilityProbeAllowed(false) succeeded")
	}
	if err := requireRKeyCapabilityProbeAllowed(true); err == nil {
		t.Fatal("requireRKeyCapabilityProbeAllowed(true) succeeded without confirmation")
	}
	t.Setenv(rkeyCapabilityConfirmEnv, rkeyCapabilityConfirmValue)
	if err := requireRKeyCapabilityProbeAllowed(true); err != nil {
		t.Fatalf("requireRKeyCapabilityProbeAllowed(true) = %v", err)
	}
}

func TestRequirePDLifecycleProbeAllowed(t *testing.T) {
	t.Setenv(pdLifecycleConfirmEnv, "")
	if err := requirePDLifecycleProbeAllowed(false); err == nil {
		t.Fatal("requirePDLifecycleProbeAllowed(false) succeeded")
	}
	if err := requirePDLifecycleProbeAllowed(true); err == nil {
		t.Fatal("requirePDLifecycleProbeAllowed(true) succeeded without confirmation")
	}
	t.Setenv(pdLifecycleConfirmEnv, pdLifecycleConfirmValue)
	if err := requirePDLifecycleProbeAllowed(true); err != nil {
		t.Fatalf("requirePDLifecycleProbeAllowed(true) = %v", err)
	}
}

func TestValidatePDLifecycleProbe(t *testing.T) {
	for _, test := range []struct {
		name                     string
		mode                     string
		cycles, allocs, maxAlloc int
		timeout, opTimeout       time.Duration
	}{
		{"bad mode", "bad", 1, 1, 1, time.Second, time.Second},
		{"bad cycles", "reclaim", 0, 1, 1, time.Second, time.Second},
		{"bad allocs", "reclaim", 1, 0, 1, time.Second, time.Second},
		{"bad max", "exhaust", 1, 1, 0, time.Second, time.Second},
		{"op exceeds total", "reclaim", 1, 1, 1, time.Second, 2 * time.Second},
	} {
		t.Run(test.name, func(t *testing.T) {
			if err := validatePDLifecycleProbe(test.mode, test.cycles, test.allocs, test.maxAlloc, test.timeout, test.opTimeout); err == nil {
				t.Fatal("validatePDLifecycleProbe succeeded")
			}
		})
	}
	if err := validatePDLifecycleProbe("reclaim", 1, 1, 1, time.Second, time.Second); err != nil {
		t.Fatalf("validatePDLifecycleProbe() = %v", err)
	}
}

func TestRCCapabilityProbeDoesNotDriveDataPath(t *testing.T) {
	data, err := os.ReadFile("main.go")
	if err != nil {
		t.Fatal(err)
	}
	src := string(data)
	start := strings.Index(src, "func probeRCCapability(")
	if start < 0 {
		t.Fatal("could not locate probeRCCapability source")
	}
	end := strings.Index(src[start:], "\nfunc probeRKeyCapability(")
	if end < 0 {
		t.Fatal("could not locate end of probeRCCapability source")
	}
	body := src[start : start+end]
	if n := strings.Count(body, "IbvCreateQpAttr("); n != 1 {
		t.Fatalf("RC capability probe has %d create-QP calls, want 1", n)
	}
	for _, forbidden := range []string{"IbvModifyQp", "IbvRegMr", "PostSend", "PostRecv", "IbvQueryPort", "IbvQueryGid"} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("RC capability probe must not call %s", forbidden)
		}
	}
}

func TestRKeyCapabilityProbeDoesNotCreateQueuePairOrDriveDataPath(t *testing.T) {
	data, err := os.ReadFile("main.go")
	if err != nil {
		t.Fatal(err)
	}
	src := string(data)
	start := strings.Index(src, "func probeRKeyCapability(")
	if start < 0 {
		t.Fatal("could not locate probeRKeyCapability source")
	}
	end := strings.Index(src[start:], "\nfunc classifyRKeyCapabilityRegistration(")
	if end < 0 {
		t.Fatal("could not locate end of probeRKeyCapability source")
	}
	body := src[start : start+end]
	if n := strings.Count(body, "IbvRegMr("); n != 1 {
		t.Fatalf("rkey capability probe has %d register-MR calls, want 1", n)
	}
	if n := strings.Count(body, "IbvDeregMr("); n != 1 {
		t.Fatalf("rkey capability probe has %d deregister-MR calls, want 1", n)
	}
	for _, forbidden := range []string{"IbvCreateQp", "IbvModifyQp", "PostSend", "PostRecv", "IbvQueryPort", "IbvQueryGid"} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("rkey capability probe must not call %s", forbidden)
		}
	}
}

func TestPDLifecycleProbeOnlyUsesPDCalls(t *testing.T) {
	data, err := os.ReadFile("main.go")
	if err != nil {
		t.Fatal(err)
	}
	src := string(data)
	for _, function := range []string{"probePDLifecycle", "probePDExhaustion", "allocatePDs", "deallocatePDs"} {
		functionStart := strings.Index(src, "func "+function+"(")
		if functionStart < 0 {
			t.Fatalf("could not locate %s", function)
		}
		functionBody := src[functionStart:]
		if next := strings.Index(functionBody[1:], "\nfunc "); next >= 0 {
			functionBody = functionBody[:next+1]
		}
		for _, forbidden := range []string{"IbvRegMr", "IbvCreateQp", "IbvCreateCq", "IbvModifyQp", "PostSend", "PostRecv", "IbvQueryPort", "IbvQueryGid"} {
			if strings.Contains(functionBody, forbidden) {
				t.Fatalf("%s must not call %s", function, forbidden)
			}
		}
	}
	if !strings.Contains(src, "startRDMAWatchdog(\"ibv_alloc_pd\"") || !strings.Contains(src, "startRDMAWatchdog(\"ibv_dealloc_pd\"") {
		t.Fatal("PD lifecycle probe has no per-operation watchdog")
	}
}

func TestClassifyRCCapabilityCreate(t *testing.T) {
	tests := []struct {
		name       string
		qp         rdma.RDMAQP
		err        error
		wantStatus string
		wantErrno  int
	}{
		{"created", 1, nil, "supported", 0},
		{"enotsup", 0, &rdma.ProviderError{Errno: int(syscall.ENOTSUP), ErrnoSet: true}, "rejected", int(syscall.ENOTSUP)},
		{"eopnotsupp", 0, &rdma.ProviderError{Errno: 102, ErrnoSet: true}, "rejected", 102},
		{"other provider errno", 0, &rdma.ProviderError{Errno: 22, ErrnoSet: true}, "inconclusive", 22},
		{"ordinary error", 0, errors.New("create failed"), "inconclusive", 0},
		{"nil queue pair", 0, nil, "inconclusive", 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotStatus, gotErrno, _ := classifyRCCapabilityCreate(test.qp, test.err)
			if gotStatus != test.wantStatus || gotErrno != test.wantErrno {
				t.Fatalf("classifyRCCapabilityCreate(%#x, %v) = (%q, %d), want (%q, %d)", test.qp, test.err, gotStatus, gotErrno, test.wantStatus, test.wantErrno)
			}
		})
	}
}

func TestClassifyRKeyCapabilityRegistration(t *testing.T) {
	tests := []struct {
		name       string
		mr         rdma.RDMAMR
		err        error
		wantStatus string
		wantErrno  int
	}{
		{"registered", 1, nil, "registered", 0},
		{"provider errno", 0, &rdma.ProviderError{Errno: 22, ErrnoSet: true}, "inconclusive", 22},
		{"ordinary error", 0, errors.New("register failed"), "inconclusive", 0},
		{"nil memory region", 0, nil, "inconclusive", 0},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotStatus, gotErrno, _ := classifyRKeyCapabilityRegistration(test.mr, test.err)
			if gotStatus != test.wantStatus || gotErrno != test.wantErrno {
				t.Fatalf("classifyRKeyCapabilityRegistration(%#x, %v) = (%q, %d), want (%q, %d)", test.mr, test.err, gotStatus, gotErrno, test.wantStatus, test.wantErrno)
			}
		})
	}
}

func TestSelectRCCapabilityDevice(t *testing.T) {
	devs := []rdma.Device{{Name: "rdma_en2"}, {Name: "rdma_en3"}}
	if got, err := selectRCCapabilityDevice(devs, "en3", -1); err != nil || got.Name != "rdma_en3" {
		t.Fatalf("select name = (%+v, %v), want rdma_en3", got, err)
	}
	if got, err := selectRCCapabilityDevice(devs, "", 0); err != nil || got.Name != "rdma_en2" {
		t.Fatalf("select index = (%+v, %v), want rdma_en2", got, err)
	}
	if _, err := selectRCCapabilityDevice(nil, "", -1); err == nil {
		t.Fatal("empty device list succeeded")
	}
}
