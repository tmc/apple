package main

import (
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
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

func TestLifecycleProbeIsGatedAndDoesNotPostData(t *testing.T) {
	data, err := os.ReadFile("main.go")
	if err != nil {
		t.Fatal(err)
	}
	src := string(data)
	start := strings.Index(src, "func rdmaLifecycleProbe(")
	end := strings.Index(src[start:], "\nfunc runLifecycleServer(")
	if start < 0 || end < 0 {
		t.Fatal("could not locate lifecycle probe")
	}
	body := src[start : start+end]
	for _, want := range []string{"lifecycleProbeConfirmEnv", "allow-lifecycle-probe", "RequireRTRAttemptAllowed", "-mrs in [1,4]"} {
		if !strings.Contains(body, want) {
			t.Fatalf("lifecycle probe missing %q", want)
		}
	}
	for _, forbidden := range []string{"PostSend", "PostRecv", "runRDMAPingpongClientLoop", "runRDMAPingpongServerLoop"} {
		if strings.Contains(body, forbidden) {
			t.Fatalf("lifecycle probe must not call %s", forbidden)
		}
	}
}

func TestValidateLifecycleStress(t *testing.T) {
	tests := []struct {
		name                  string
		level, listen, addr   string
		rounds, mrs, qps      int
		data                  bool
		size, iters           int
		timeout, setupTimeout time.Duration
		wantErr               bool
	}{
		{"l1 minimum", lifecycleStressCountScale, ":1234", "", 1, 1, 1, false, 64, 0, time.Minute, time.Second, false},
		{"l1 scale", lifecycleStressCountScale, ":1234", "", 3, 101, 11, false, 64, 0, time.Minute, time.Second, false},
		{"l1 data", lifecycleStressCountScale, ":1234", "", 1, 1, 1, true, 512 * 1024, 1, time.Minute, time.Second, false},
		{"l1 too many mrs", lifecycleStressCountScale, ":1234", "", 1, 102, 1, false, 64, 0, time.Minute, time.Second, true},
		{"l1 too few mrs", lifecycleStressCountScale, ":1234", "", 1, 2, 3, false, 64, 0, time.Minute, time.Second, true},
		{"l1 too many qps", lifecycleStressCountScale, ":1234", "", 1, 12, 12, false, 64, 0, time.Minute, time.Second, true},
		{"l2 minimum", lifecycleStressRoundDepth, "", "127.0.0.1:1234", 1, 1, 1, false, 64, 0, time.Minute, time.Second, false},
		{"l2 maximum", lifecycleStressRoundDepth, "", "127.0.0.1:1234", 1000, 4, 1, false, 64, 0, time.Hour, time.Second, false},
		{"l2 multiple qps", lifecycleStressRoundDepth, "", "127.0.0.1:1234", 50, 2, 2, false, 64, 0, time.Minute, time.Second, true},
		{"l2 too many mrs", lifecycleStressRoundDepth, "", "127.0.0.1:1234", 50, 5, 1, false, 64, 0, time.Minute, time.Second, true},
		{"l4 concurrent", lifecycleStressConcurrency, "", "127.0.0.1:1234", 1, 2, 2, true, 64, 1, time.Minute, time.Second, false},
		{"l4 requires data", lifecycleStressConcurrency, "", "127.0.0.1:1234", 1, 2, 2, false, 64, 0, time.Minute, time.Second, true},
		{"l4 boundary", lifecycleStressConcurrency, "", "127.0.0.1:1234", 1, 11, 11, true, 64, 1, time.Minute, time.Second, false},
		{"l4 too many qps", lifecycleStressConcurrency, "", "127.0.0.1:1234", 1, 12, 12, true, 64, 1, time.Minute, time.Second, true},
		{"data needs iters", lifecycleStressRoundDepth, "", "127.0.0.1:1234", 1, 1, 1, true, 64, 0, time.Minute, time.Second, true},
		{"iters needs data", lifecycleStressRoundDepth, "", "127.0.0.1:1234", 1, 1, 1, false, 64, 1, time.Minute, time.Second, true},
		{"bad level", "l3-data-path", ":1234", "", 1, 1, 1, false, 64, 0, time.Minute, time.Second, true},
		{"both roles", lifecycleStressCountScale, ":1234", "127.0.0.1:1234", 1, 1, 1, false, 64, 0, time.Minute, time.Second, true},
		{"setup exceeds whole", lifecycleStressCountScale, ":1234", "", 1, 1, 1, false, 64, 0, time.Second, time.Minute, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateLifecycleStress(test.level, test.listen, test.addr, test.rounds, test.mrs, test.qps, test.data, test.size, test.iters, test.timeout, test.setupTimeout)
			if (err != nil) != test.wantErr {
				t.Fatalf("validateLifecycleStress() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func TestValidateLifecycleStressIdle(t *testing.T) {
	tests := []struct {
		name      string
		dwell     time.Duration
		timeout   time.Duration
		wantError bool
	}{
		{"valid", 5 * time.Minute, 10 * time.Minute, false},
		{"zero dwell", 0, 10 * time.Minute, true},
		{"dwell exceeds timeout", 5 * time.Minute, 5 * time.Minute, true},
		{"dwell too long", 3 * time.Hour, 4 * time.Hour, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateLifecycleStressIdle(lifecycleStressIdleDegradation, ":1234", "", 1, 2, 1, true, 64, 1, test.dwell, test.timeout, time.Second)
			if (err != nil) != test.wantError {
				t.Fatalf("validateLifecycleStressIdle() error = %v, wantError %v", err, test.wantError)
			}
		})
	}
}

func TestLifecycleStressRequiresDataGate(t *testing.T) {
	data, err := os.ReadFile("main.go")
	if err != nil {
		t.Fatal(err)
	}
	src := string(data)
	start := strings.Index(src, "func rdmaLifecycleStress(")
	end := strings.Index(src[start:], "\nfunc validateLifecycleStress(")
	if start < 0 || end < 0 {
		t.Fatal("could not locate lifecycle stress probe")
	}
	body := src[start : start+end]
	for _, want := range []string{"lifecycleStressConfirmEnv", "allow-lifecycle-stress", "CONFIRM_RDMA_LIFECYCLE_DATA", "allow-lifecycle-data", "RequireRTRAttemptAllowed", "startRDMAWatchdog"} {
		if !strings.Contains(body, want) {
			t.Fatalf("lifecycle stress probe missing %q", want)
		}
	}
}

func TestRDMAPayload(t *testing.T) {
	buf := make([]byte, 64)
	fillRDMAPayload(buf)
	if err := checkRDMAPayload(buf); err != nil {
		t.Fatalf("checkRDMAPayload() = %v", err)
	}
	buf[17]++
	if err := checkRDMAPayload(buf); err == nil || !strings.Contains(err.Error(), "data mismatch") {
		t.Fatalf("checkRDMAPayload() = %v, want data mismatch", err)
	}
}

func TestWireBenchSampleJSON(t *testing.T) {
	samples := []wireBenchSample{
		{
			Total: 200, StagingCopy: 10, Post: 20, CompletionWait: 150, ReceiveCopy: 15,
			Reduction: 0, AllocBytes: 30, Allocs: 2, GCPause: 3, QueueDepth: 2, Completions: 2, Polls: 4,
		},
		{
			Total: 210, StagingCopy: 11, Post: 21, CompletionWait: 157, ReceiveCopy: 16,
			Reduction: 0, AllocBytes: 30, Allocs: 2, GCPause: 3, QueueDepth: 2, Completions: 2, Polls: 5,
		},
	}
	path := filepath.Join(t.TempDir(), "samples.json")
	collector := wireSampleCollector{samples: samples}
	if err := collector.write(path); err != nil {
		t.Fatalf("write() = %v", err)
	}
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var raw []map[string]json.RawMessage
	if err := json.Unmarshal(data, &raw); err != nil {
		t.Fatal(err)
	}
	wantFields := []string{"Total", "StagingCopy", "Post", "CompletionWait", "ReceiveCopy", "Reduction", "AllocBytes", "Allocs", "GCPause", "QueueDepth", "Completions", "Polls"}
	for _, sample := range raw {
		for _, field := range wantFields {
			if _, ok := sample[field]; !ok {
				t.Fatalf("JSON sample missing %q: %s", field, data)
			}
		}
	}
	var got []wireBenchSample
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	if len(got) != len(samples) || got[1].CompletionWait != samples[1].CompletionWait || got[1].Polls != samples[1].Polls {
		t.Fatalf("JSON round trip = %#v, want %#v", got, samples)
	}
	if err := validateWireBenchSamples(got); err != nil {
		t.Fatalf("validateWireBenchSamples() = %v", err)
	}
	if err := validateWireBenchSamples([]wireBenchSample{{Total: -1}}); err == nil {
		t.Fatal("validateWireBenchSamples accepted negative duration")
	}
	t.Logf("wire sample JSON: %s", data)
}

func TestSGCapabilityWireSampleJSON(t *testing.T) {
	samples := []wireBenchSample{{
		Total: 200, StagingCopy: 10, Post: 20, CompletionWait: 150, ReceiveCopy: 15,
		QueueDepth: 2, Completions: 2, Polls: 4,
	}}
	data, err := json.Marshal(samples)
	if err != nil {
		t.Fatal(err)
	}
	var got []wireBenchSample
	if err := json.Unmarshal(data, &got); err != nil {
		t.Fatal(err)
	}
	if err := validateWireBenchSamples(got); err != nil {
		t.Fatalf("validateWireBenchSamples() = %v", err)
	}
	if len(got) != 1 || got[0].Total != 200 || got[0].QueueDepth != 2 || got[0].Completions != 2 || got[0].Polls != 4 {
		t.Fatalf("scatter/gather wire sample round trip = %#v", got)
	}
}

func TestValidateSGCapability(t *testing.T) {
	if err := validateSGCapability(":1234", "", 4, 1<<20, "", time.Minute, time.Second); err != nil {
		t.Fatalf("validateSGCapability() = %v", err)
	}
	for _, test := range []struct {
		name                  string
		listen, addr, samples string
		segments, segmentSize int
		timeout, setupTimeout time.Duration
	}{
		{"both roles", ":1234", "127.0.0.1:1234", "", 4, 1 << 20, time.Minute, time.Second},
		{"bad segments", ":1234", "", "", 3, 1 << 20, time.Minute, time.Second},
		{"not four mib", ":1234", "", "", 4, 512 << 10, time.Minute, time.Second},
		{"server samples", ":1234", "", "samples.json", 4, 1 << 20, time.Minute, time.Second},
		{"setup exceeds total", ":1234", "", "", 4, 1 << 20, time.Second, time.Minute},
	} {
		t.Run(test.name, func(t *testing.T) {
			if err := validateSGCapability(test.listen, test.addr, test.segments, test.segmentSize, test.samples, test.timeout, test.setupTimeout); err == nil {
				t.Fatal("validateSGCapability succeeded")
			}
		})
	}
}

func TestSGCeilingError(t *testing.T) {
	for _, test := range []struct {
		input string
		want  bool
	}{
		{"ibv_post_recv: errno 4294967284", true},
		{"ibv_post_recv: errno 12 (ENOMEM)", true},
		{"data mismatch", false},
	} {
		if got := isSGCeilingError(test.input); got != test.want {
			t.Errorf("isSGCeilingError(%q) = %v, want %v", test.input, got, test.want)
		}
	}
}

func TestValidateWireSamples(t *testing.T) {
	tests := []struct {
		name        string
		path, addr  string
		rounds, qps int
		data        bool
		wantErr     bool
	}{
		{"disabled", "", "", 1, 1, false, false},
		{"client", "samples.json", "127.0.0.1:1234", 1, 1, true, false},
		{"needs data", "samples.json", "127.0.0.1:1234", 1, 1, false, true},
		{"needs client", "samples.json", "", 1, 1, true, true},
		{"one round", "samples.json", "127.0.0.1:1234", 2, 1, true, true},
		{"multiple qps", "samples.json", "127.0.0.1:1234", 1, 2, true, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateWireSamples(test.path, test.addr, test.rounds, test.qps, test.data)
			if (err != nil) != test.wantErr {
				t.Fatalf("validateWireSamples() error = %v, wantErr %v", err, test.wantErr)
			}
		})
	}
}

func TestLifecycleStressGateEnv(t *testing.T) {
	if got, want := lifecycleStressGateEnv(lifecycleStressCountScale, false), "CONFIRM_RDMA_LIFECYCLE_STRESS=l1-count-scale"; got != want {
		t.Fatalf("lifecycleStressGateEnv(false) = %q, want %q", got, want)
	}
	if got, want := lifecycleStressGateEnv(lifecycleStressRoundDepth, true), "CONFIRM_RDMA_LIFECYCLE_STRESS=l2-round-depth,CONFIRM_RDMA_LIFECYCLE_DATA=uc-send-recv"; got != want {
		t.Fatalf("lifecycleStressGateEnv(true) = %q, want %q", got, want)
	}
}

func TestClassifyRDMABenchFailure(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"ibv_create_qp: provider returned nil queue pair (return=0, errno 16 (EBUSY))", "resource_exhausted"},
		{"ibv_reg_mr: provider returned nil memory region (return=0, errno 0)", "resource_exhausted"},
		{"data mismatch at byte 1", "data_mismatch"},
		{"rdma setup timeout", string(rdma.FailureProviderTimeout)},
	}
	for _, test := range tests {
		if got := classifyRDMABenchFailure(test.input); got != test.want {
			t.Errorf("classifyRDMABenchFailure(%q) = %q, want %q", test.input, got, test.want)
		}
	}
}

func TestLifecycleResourceCounts(t *testing.T) {
	resources := []*rdmaResources{
		{qp: 1, pd: 1, mr: 1, extraMRs: []rdma.RDMAMR{1, 2}},
		{qp: 2, pd: 2, mr: 2},
	}
	qps, pds, mrs := lifecycleResourceCounts(resources)
	if qps != 2 || pds != 2 || mrs != 4 {
		t.Fatalf("lifecycleResourceCounts() = (%d, %d, %d), want (2, 2, 4)", qps, pds, mrs)
	}
	var result rdmaBenchResult
	setLifecycleResourceCounts(&result, resources)
	if result.QPsOpened != 2 || result.PDsOpened != 2 || result.MRsOpened != 4 {
		t.Fatalf("setLifecycleResourceCounts() = (%d, %d, %d), want (2, 2, 4)", result.QPsOpened, result.PDsOpened, result.MRsOpened)
	}
}

func TestFinishRDMADatapath(t *testing.T) {
	var res rdmaBenchResult
	finishRDMADatapath(&res, 2*time.Second, 2048, 5)
	if res.DatapathElapsed != "2s" || res.DatapathBytesPerSec != 1024 || res.DatapathMessagesPerSec != 2.5 {
		t.Fatalf("finishRDMADatapath() = (%q, %v, %v), want (2s, 1024, 2.5)", res.DatapathElapsed, res.DatapathBytesPerSec, res.DatapathMessagesPerSec)
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
