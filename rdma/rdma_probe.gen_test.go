// Code generated from Apple documentation for rdma. DO NOT EDIT.

package rdma

import (
	"os"
	"testing"
)

func TestAvailableDoesNotPanic(t *testing.T) {
	_ = Available()
}

func TestDevicesProbe(t *testing.T) {
	if !Available() {
		t.Skip("librdma probe symbols are unavailable")
	}
	if _, err := Devices(); err != nil {
		t.Fatalf("Devices: %v", err)
	}
}

func TestRDMAIntegration(t *testing.T) {
	if os.Getenv("APPLE_RDMA_INTEGRATION") != "1" {
		t.Skip("set APPLE_RDMA_INTEGRATION=1 to require an RDMA device")
	}
	devs, err := Devices()
	if err != nil {
		t.Fatalf("Devices: %v", err)
	}
	if len(devs) == 0 {
		t.Fatal("no RDMA devices")
	}
}
