package rdma

import (
	"strings"
	"testing"
	"unsafe"
)

func TestCompletionChannelHelpersRejectNilInputs(t *testing.T) {
	saveRDMAFuncs(t)
	var calls int
	_ibvCreateCompChannel = func(RDMAContext) RDMACompChannel { calls++; return 1 }
	_ibvDestroyCompChannel = func(RDMACompChannel) int { calls++; return 0 }
	_ibvGetCqEvent = func(RDMACompChannel, uintptr, uintptr) int { calls++; return 0 }
	_ibvAckCqEvents = func(RDMACQ, uint32) { calls++ }

	tests := []struct {
		name string
		call func() error
	}{
		{"create channel", func() error { _, err := IbvCreateCompChannel(0); return err }},
		{"destroy channel", func() error { _, err := IbvDestroyCompChannel(0); return err }},
		{"get event channel", func() error {
			var cq RDMACQ
			var ctx uintptr
			_, err := IbvGetCqEvent(0, &cq, &ctx)
			return err
		}},
		{"get event cq", func() error {
			var ctx uintptr
			_, err := IbvGetCqEvent(1, nil, &ctx)
			return err
		}},
		{"get event context", func() error {
			var cq RDMACQ
			_, err := IbvGetCqEvent(1, &cq, nil)
			return err
		}},
		{"ack events", func() error { return IbvAckCqEvents(0, 1) }},
		{"notify cq", func() error { _, err := IbvReqNotifyCq(0, 0); return err }},
	}
	for _, tt := range tests {
		if err := tt.call(); err == nil {
			t.Fatalf("%s returned nil error", tt.name)
		}
	}
	if calls != 0 {
		t.Fatalf("provider called %d times for nil completion inputs", calls)
	}
}

func TestCompletionChannelProviderWrappers(t *testing.T) {
	saveRDMAFuncs(t)
	var gotContext RDMAContext
	_ibvCreateCompChannel = func(context RDMAContext) RDMACompChannel {
		gotContext = context
		return 2
	}
	channel, err := IbvCreateCompChannel(1)
	if err != nil {
		t.Fatalf("IbvCreateCompChannel: %v", err)
	}
	if channel != 2 || gotContext != 1 {
		t.Fatalf("create channel = (%d, context %d), want (2, 1)", channel, gotContext)
	}

	var gotChannel RDMACompChannel
	_ibvDestroyCompChannel = func(channel RDMACompChannel) int {
		gotChannel = channel
		return 0
	}
	rc, err := IbvDestroyCompChannel(2)
	if err != nil || rc != 0 || gotChannel != 2 {
		t.Fatalf("IbvDestroyCompChannel = rc %d err %v channel %d, want rc 0 channel 2", rc, err, gotChannel)
	}

	_ibvGetCqEvent = func(channel RDMACompChannel, cq uintptr, cqContext uintptr) int {
		*(*RDMACQ)(unsafe.Pointer(cq)) = 3
		*(*uintptr)(unsafe.Pointer(cqContext)) = 4
		return 0
	}
	var cq RDMACQ
	var cqContext uintptr
	rc, err = IbvGetCqEvent(2, &cq, &cqContext)
	if err != nil || rc != 0 || cq != 3 || cqContext != 4 {
		t.Fatalf("IbvGetCqEvent = rc %d err %v cq %d ctx %d, want rc 0 cq 3 ctx 4", rc, err, cq, cqContext)
	}

	var ackCQ RDMACQ
	var ackEvents uint32
	_ibvAckCqEvents = func(cq RDMACQ, nevents uint32) {
		ackCQ = cq
		ackEvents = nevents
	}
	if err := IbvAckCqEvents(3, 5); err != nil {
		t.Fatalf("IbvAckCqEvents: %v", err)
	}
	if ackCQ != 3 || ackEvents != 5 {
		t.Fatalf("ack = cq %d events %d, want cq 3 events 5", ackCQ, ackEvents)
	}
}

func TestCompletionChannelProviderErrors(t *testing.T) {
	saveRDMAFuncs(t)
	_ibvCreateCompChannel = func(RDMAContext) RDMACompChannel { return 0 }
	if _, err := IbvCreateCompChannel(1); err == nil || !strings.Contains(err.Error(), "provider returned nil completion channel") {
		t.Fatalf("IbvCreateCompChannel nil err = %v, want provider returned nil completion channel", err)
	}
	type channel struct {
		context RDMAContext
		fd      int32
		refcnt  int32
	}
	ch := RDMACompChannel(uintptr(unsafe.Pointer(&channel{context: 1})))
	_ibvDestroyCompChannel = func(RDMACompChannel) int { return -1 }
	if _, err := IbvDestroyCompChannel(ch); err == nil || !strings.Contains(err.Error(), "provider returned negative status") {
		t.Fatalf("IbvDestroyCompChannel negative err = %v, want provider returned negative status", err)
	}
	_ibvGetCqEvent = func(RDMACompChannel, uintptr, uintptr) int { return -1 }
	var cq RDMACQ
	var cqContext uintptr
	if _, err := IbvGetCqEvent(ch, &cq, &cqContext); err == nil || !strings.Contains(err.Error(), "provider returned negative status") {
		t.Fatalf("IbvGetCqEvent negative err = %v, want provider returned negative status", err)
	}
}

func TestCompletionChannelStructAccessors(t *testing.T) {
	type channel struct {
		context RDMAContext
		fd      int32
		refcnt  int32
	}
	type cq struct {
		context RDMAContext
		channel RDMACompChannel
	}
	ch := channel{context: 11, fd: 42, refcnt: 1}
	rawCQ := cq{context: 11, channel: RDMACompChannel(uintptr(unsafe.Pointer(&ch)))}

	if got := IbvCompChannelFD(RDMACompChannel(uintptr(unsafe.Pointer(&ch)))); got != 42 {
		t.Fatalf("IbvCompChannelFD = %d, want 42", got)
	}
	if got := IbvCQChannel(RDMACQ(uintptr(unsafe.Pointer(&rawCQ)))); got != RDMACompChannel(uintptr(unsafe.Pointer(&ch))) {
		t.Fatalf("IbvCQChannel = %d, want %d", got, uintptr(unsafe.Pointer(&ch)))
	}
	if got := rdmaContextFromCompChannel(RDMACompChannel(uintptr(unsafe.Pointer(&ch)))); got != 11 {
		t.Fatalf("rdmaContextFromCompChannel = %d, want 11", got)
	}
}
