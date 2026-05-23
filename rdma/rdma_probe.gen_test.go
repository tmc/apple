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
	if os.Getenv("APPLE_RDMA_PROBE") != "1" {
		t.Skip("set APPLE_RDMA_PROBE=1 to call ibv_get_device_list")
	}
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

func TestRDMAOneShotProviderLifecycle(t *testing.T) {
	if os.Getenv("APPLE_RDMA_ONESHOT") != "1" {
		t.Skip("set APPLE_RDMA_ONESHOT=1 for one serialized hardware lifecycle probe; optional APPLE_RDMA_DEVICE selects a device")
	}
	if !Available() {
		t.Skip("librdma probe symbols are unavailable")
	}
	list, err := OpenDeviceList()
	if err != nil {
		t.Fatalf("OpenDeviceList: %v", err)
	}
	defer list.Close()
	devs := list.Devices()
	if len(devs) == 0 {
		t.Fatal("no RDMA devices")
	}
	dev := devs[0]
	if name := os.Getenv("APPLE_RDMA_DEVICE"); name != "" {
		dev = Device{}
		for _, candidate := range devs {
			if candidate.Name == name {
				dev = candidate
				break
			}
		}
		if dev.Handle == 0 {
			t.Fatalf("APPLE_RDMA_DEVICE=%q not found; available devices: %v", name, deviceNames(devs))
		}
	}
	t.Logf("APPLE_RDMA_ONESHOT=1 APPLE_RDMA_DEVICE=%q selected=%s", os.Getenv("APPLE_RDMA_DEVICE"), dev.Name)
	ctx, err := dev.Open()
	if err != nil {
		t.Fatalf("ibv_open_device: %v", err)
	}
	if ctx == 0 {
		t.Fatal("ibv_open_device returned nil context")
	}
	defer IbvCloseDevice(ctx)

	var port IbvPortAttr
	if rc, err := IbvQueryPortAttr(ctx, 1, &port); err != nil || rc != 0 {
		t.Fatalf("ibv_query_port: rc=%d err=%v", rc, err)
	}
	pd, err := IbvAllocPd(ctx)
	if err != nil {
		t.Fatalf("ibv_alloc_pd: %v", err)
	}
	if pd == 0 {
		t.Fatal("ibv_alloc_pd returned nil protection domain")
	}
	defer IbvDeallocPd(pd)
	cq, err := IbvCreateCq(ctx, 16, 0, 0, 0)
	if err != nil {
		t.Fatalf("ibv_create_cq: %v", err)
	}
	if cq == 0 {
		t.Fatal("ibv_create_cq returned nil completion queue")
	}
	defer IbvDestroyCq(cq)

	init := IbvQPInitAttr{
		SendCQ: cq,
		RecvCQ: cq,
		Cap: IbvQPCap{
			MaxSendWR:  1,
			MaxRecvWR:  1,
			MaxSendSGE: 1,
			MaxRecvSGE: 1,
		},
		QPType: IBV_QPT_UC,
	}
	qp, err := IbvCreateQpAttr(pd, &init)
	if err != nil {
		t.Fatalf("ibv_create_qp: %v", err)
	}
	if qp == 0 {
		t.Fatal("ibv_create_qp returned nil queue pair")
	}
	if rc, err := IbvDestroyQp(qp); err != nil || rc != 0 {
		t.Fatalf("ibv_destroy_qp: rc=%d err=%v", rc, err)
	}
}

func deviceNames(devs []Device) []string {
	names := make([]string, 0, len(devs))
	for _, dev := range devs {
		names = append(names, dev.Name)
	}
	return names
}
