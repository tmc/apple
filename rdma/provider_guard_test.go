package rdma

import (
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

type savedRDMAFuncs struct {
	allocPD        func(RDMAContext) RDMAPD
	closeDevice    func(RDMAContext) int
	createCQ       func(RDMAContext, int, uintptr, uintptr, int) RDMACQ
	createQP       func(RDMAPD, uintptr) RDMAQP
	deallocPD      func(RDMAPD) int
	deregMR        func(RDMAMR) int
	destroyCQ      func(RDMACQ) int
	destroyQP      func(RDMAQP) int
	freeDeviceList func(RDMADeviceList)
	modifyQP       func(RDMAQP, uintptr, int) int
	openDevice     func(RDMADevice) RDMAContext
	queryDevice    func(RDMAContext, uintptr) int
	queryGID       func(RDMAContext, uint8, int, uintptr) int
	queryPort      func(RDMAContext, uint8, uintptr) int
	regMR          func(RDMAPD, uintptr, uintptr, int) RDMAMR
}

func saveRDMAFuncs(t *testing.T) {
	t.Helper()
	saved := savedRDMAFuncs{
		allocPD:        _ibvAllocPd,
		closeDevice:    _ibvCloseDevice,
		createCQ:       _ibvCreateCq,
		createQP:       _ibvCreateQp,
		deallocPD:      _ibvDeallocPd,
		deregMR:        _ibvDeregMr,
		destroyCQ:      _ibvDestroyCq,
		destroyQP:      _ibvDestroyQp,
		freeDeviceList: _ibvFreeDeviceList,
		modifyQP:       _ibvModifyQp,
		openDevice:     _ibvOpenDevice,
		queryDevice:    _ibvQueryDevice,
		queryGID:       _ibvQueryGid,
		queryPort:      _ibvQueryPort,
		regMR:          _ibvRegMr,
	}
	t.Cleanup(func() {
		_ibvAllocPd = saved.allocPD
		_ibvCloseDevice = saved.closeDevice
		_ibvCreateCq = saved.createCQ
		_ibvCreateQp = saved.createQP
		_ibvDeallocPd = saved.deallocPD
		_ibvDeregMr = saved.deregMR
		_ibvDestroyCq = saved.destroyCQ
		_ibvDestroyQp = saved.destroyQP
		_ibvFreeDeviceList = saved.freeDeviceList
		_ibvModifyQp = saved.modifyQP
		_ibvOpenDevice = saved.openDevice
		_ibvQueryDevice = saved.queryDevice
		_ibvQueryGid = saved.queryGID
		_ibvQueryPort = saved.queryPort
		_ibvRegMr = saved.regMR
	})
}

func TestZeroHandleTeardownRefusesProviderCall(t *testing.T) {
	saveRDMAFuncs(t)
	var calls int
	_ibvCloseDevice = func(RDMAContext) int { calls++; return 0 }
	_ibvDeallocPd = func(RDMAPD) int { calls++; return 0 }
	_ibvDeregMr = func(RDMAMR) int { calls++; return 0 }
	_ibvDestroyCq = func(RDMACQ) int { calls++; return 0 }
	_ibvDestroyQp = func(RDMAQP) int { calls++; return 0 }
	_ibvFreeDeviceList = func(RDMADeviceList) { calls++ }

	tests := []struct {
		name string
		call func() error
	}{
		{"close device", func() error { _, err := IbvCloseDevice(0); return err }},
		{"dealloc pd", func() error { _, err := IbvDeallocPd(0); return err }},
		{"dereg mr", func() error { _, err := IbvDeregMr(0); return err }},
		{"destroy cq", func() error { _, err := IbvDestroyCq(0); return err }},
		{"destroy qp", func() error { _, err := IbvDestroyQp(0); return err }},
	}
	for _, tt := range tests {
		if err := tt.call(); err == nil {
			t.Fatalf("%s returned nil error", tt.name)
		}
	}
	if err := IbvFreeDeviceList(0); err != nil {
		t.Fatalf("free nil device list: %v", err)
	}
	if calls != 0 {
		t.Fatalf("provider called %d times for nil teardown handles", calls)
	}
}

func TestResourcesCloseOrderAndIdempotence(t *testing.T) {
	saveRDMAFuncs(t)
	var order []string
	record := func(name string) int {
		order = append(order, name)
		return 0
	}
	_ibvDestroyQp = func(RDMAQP) int { return record("qp") }
	_ibvDeregMr = func(RDMAMR) int { return record("mr") }
	_ibvDestroyCq = func(RDMACQ) int { return record("cq") }
	_ibvDeallocPd = func(RDMAPD) int { return record("pd") }
	_ibvCloseDevice = func(RDMAContext) int { return record("context") }

	res := &Resources{Context: 1, PD: 2, MR: 3, CQ: 4, QP: 5}
	if err := res.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}
	if got, want := strings.Join(order, ","), "qp,mr,cq,pd,context"; got != want {
		t.Fatalf("close order = %s, want %s", got, want)
	}
	if res.Context != 0 || res.PD != 0 || res.MR != 0 || res.CQ != 0 || res.QP != 0 {
		t.Fatalf("resources not cleared after Close: %#v", res)
	}
	if err := res.Close(); err != nil {
		t.Fatalf("second Close: %v", err)
	}
	if got, want := strings.Join(order, ","), "qp,mr,cq,pd,context"; got != want {
		t.Fatalf("second Close changed order to %s, want %s", got, want)
	}
}

func TestDeviceListCloseIdempotent(t *testing.T) {
	saveRDMAFuncs(t)
	var calls int
	_ibvFreeDeviceList = func(RDMADeviceList) { calls++ }

	list := &DeviceList{
		list:    99,
		devices: []Device{{Handle: 1, Name: "rdma0"}},
	}
	if err := list.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}
	if err := list.Close(); err != nil {
		t.Fatalf("second Close: %v", err)
	}
	if calls != 1 {
		t.Fatalf("free device list calls = %d, want 1", calls)
	}
	if got := list.Devices(); len(got) != 0 {
		t.Fatalf("Devices after Close length = %d, want 0", len(got))
	}
}

func TestDeviceListEntriesScansZeroCountList(t *testing.T) {
	raw := [...]RDMADevice{11, 22, 33, 0}
	got, err := rdmaDeviceListEntries(RDMADeviceList(uintptr(unsafe.Pointer(&raw[0]))), 0)
	if err != nil {
		t.Fatalf("rdmaDeviceListEntries: %v", err)
	}
	if len(got) != 3 {
		t.Fatalf("len = %d, want 3", len(got))
	}
	for i, want := range raw[:3] {
		if got[i] != want {
			t.Fatalf("got[%d] = %d, want %d", i, got[i], want)
		}
	}
}

func TestTypedHelpersRejectNilPointers(t *testing.T) {
	saveRDMAFuncs(t)
	var calls int
	_ibvCreateQp = func(RDMAPD, uintptr) RDMAQP { calls++; return 0 }
	_ibvModifyQp = func(RDMAQP, uintptr, int) int { calls++; return 0 }
	_ibvQueryPort = func(RDMAContext, uint8, uintptr) int { calls++; return 0 }
	_ibvQueryGid = func(RDMAContext, uint8, int, uintptr) int { calls++; return 0 }
	_ibvQueryDevice = func(RDMAContext, uintptr) int { calls++; return 0 }

	tests := []struct {
		name string
		call func() error
	}{
		{"create qp", func() error { _, err := IbvCreateQpAttr(1, nil); return err }},
		{"modify qp", func() error { _, err := IbvModifyQpAttr(1, nil, 0); return err }},
		{"query port", func() error { _, err := IbvQueryPortAttr(1, 1, nil); return err }},
		{"query gid", func() error { _, err := IbvQueryGidInto(1, 1, 0, nil); return err }},
		{"query device", func() error { _, err := IbvQueryDeviceBytes(1, nil); return err }},
	}
	for _, tt := range tests {
		if err := tt.call(); err == nil {
			t.Fatalf("%s returned nil error", tt.name)
		}
	}
	if calls != 0 {
		t.Fatalf("provider called %d times for nil typed helper pointers", calls)
	}
}

func TestZeroPollerPosterRefuseNilCalls(t *testing.T) {
	if got := (IbvCQPoller{}).Poll(1, &IbvWC{}); got != -1 {
		t.Fatalf("zero poller Poll = %d, want -1", got)
	}
	if got := (IbvCQPoller{}).Poll(1, nil); got != -1 {
		t.Fatalf("zero poller nil Poll = %d, want -1", got)
	}
	if got := (IbvQPPoster{}).PostSend(&IbvSendWR{}, new(*IbvSendWR)); got != -1 {
		t.Fatalf("zero poster PostSend = %d, want -1", got)
	}
	if got := (IbvQPPoster{}).PostRecv(&IbvRecvWR{}, new(*IbvRecvWR)); got != -1 {
		t.Fatalf("zero poster PostRecv = %d, want -1", got)
	}
}

func TestProviderCallsAreSerialized(t *testing.T) {
	saveRDMAFuncs(t)
	var active atomic.Int32
	var maxActive atomic.Int32
	_ibvOpenDevice = func(device RDMADevice) RDMAContext {
		now := active.Add(1)
		for {
			max := maxActive.Load()
			if now <= max || maxActive.CompareAndSwap(max, now) {
				break
			}
		}
		time.Sleep(5 * time.Millisecond)
		active.Add(-1)
		return RDMAContext(device + 100)
	}

	var wg sync.WaitGroup
	for i := 0; i < 8; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ctx, err := IbvOpenDevice(RDMADevice(i + 1))
			if err != nil {
				t.Errorf("IbvOpenDevice: %v", err)
			}
			if ctx == 0 {
				t.Errorf("IbvOpenDevice returned nil context")
			}
		}(i)
	}
	wg.Wait()
	if got := maxActive.Load(); got != 1 {
		t.Fatalf("max concurrent provider calls = %d, want 1", got)
	}
}
