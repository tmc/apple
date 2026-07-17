package rdma

import (
	"errors"
	"testing"
)

func TestAppleThunderboltCapabilities(t *testing.T) {
	got := AppleThunderboltCapabilities()
	if !got.QueuePairUC || !got.SendRecv || !got.RDMAWrite {
		t.Fatalf("capabilities = %+v, want UC SEND/RECV and WRITE", got)
	}
	if got.RDMARead {
		t.Fatalf("capabilities = %+v, RDMARead must be false", got)
	}
	if !errors.Is(ErrRDMAReadUnsupported, ErrUnsupportedOperation) {
		t.Fatalf("ErrRDMAReadUnsupported does not wrap ErrUnsupportedOperation")
	}
}
