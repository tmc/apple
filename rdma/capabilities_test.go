package rdma

import (
	"errors"
	"testing"
)

func TestAppleThunderboltCapabilities(t *testing.T) {
	got := AppleThunderboltCapabilities()
	if got.QueuePairUC != CapabilityConfirmed || got.SendRecv != CapabilityConfirmed {
		t.Fatalf("capabilities = %+v, want confirmed UC/SEND/RECV", got)
	}
	if got.QueuePairRC != CapabilityUnknown || got.RDMARead != CapabilityUnknown || got.RDMAWrite != CapabilityUnknown {
		t.Fatalf("capabilities = %+v, want unknown RC/READ/WRITE", got)
	}
	if !errors.Is(ErrRDMAReadUnavailable, ErrUnsupportedOperation) {
		t.Fatalf("ErrRDMAReadUnavailable does not wrap ErrUnsupportedOperation")
	}
}
