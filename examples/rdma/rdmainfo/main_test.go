package main

import (
	"testing"
	"unsafe"
)

func TestParsePorts(t *testing.T) {
	got := parsePorts("1, 2,255")
	want := []uint8{1, 2, 255}
	if len(got) != len(want) {
		t.Fatalf("parsePorts length = %d, want %d", len(got), len(want))
	}
	for i := range want {
		if got[i] != want[i] {
			t.Fatalf("parsePorts[%d] = %d, want %d", i, got[i], want[i])
		}
	}
}

func TestQueryBuffer(t *testing.T) {
	min := int(unsafe.Sizeof(ibvDeviceAttr{}))
	buf := queryBuffer(min, min, "device attr")
	if len(buf) != min {
		t.Fatalf("queryBuffer length = %d, want %d", len(buf), min)
	}
	if got := queryBuffer(0, min, "device attr"); got != nil {
		t.Fatalf("queryBuffer zero = %#v, want nil", got)
	}
}

func TestQueryStepPreview(t *testing.T) {
	step := queryStep("ibv_query_device", 0, nil, []byte{0, 1, 2, 3}, 2)
	if !step.OK {
		t.Fatal("queryStep OK = false, want true")
	}
	if step.Preview != "0001" {
		t.Fatalf("queryStep preview = %q, want 0001", step.Preview)
	}
}

func TestRDMAReadinessNames(t *testing.T) {
	if got := rdmaNetInterface("rdma_en3"); got != "en3" {
		t.Fatalf("rdmaNetInterface = %q, want en3", got)
	}
	if got := portStateName(4); got != "PORT_ACTIVE" {
		t.Fatalf("portStateName = %q, want PORT_ACTIVE", got)
	}
	if got := linkLayerName(100); got != "Thunderbolt" {
		t.Fatalf("linkLayerName = %q, want Thunderbolt", got)
	}
	if got := mtuBytes(5); got != 4096 {
		t.Fatalf("mtuBytes = %d, want 4096", got)
	}
}
