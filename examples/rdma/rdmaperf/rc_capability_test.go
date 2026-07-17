package main

import (
	"errors"
	"os"
	"strings"
	"syscall"
	"testing"
	"time"

	"github.com/tmc/apple/rdma"
)

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
	end := strings.Index(src[start:], "\nfunc selectRCCapabilityDevice(")
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
