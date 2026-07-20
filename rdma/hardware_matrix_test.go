package rdma

import (
	"fmt"
	"os"
	"runtime"
	"syscall"
	"testing"
	"time"
	"unsafe"
)

func TestRDMAHardwareProviderMatrix(t *testing.T) {
	if os.Getenv("APPLE_RDMA_HARDWARE_MATRIX") != "1" {
		t.Skip("set APPLE_RDMA_HARDWARE_MATRIX=1 and APPLE_RDMA_DEVICE=<name> for one serialized provider matrix; run the test process under an external timeout")
	}
	wantDevice := os.Getenv("APPLE_RDMA_DEVICE")
	if wantDevice == "" {
		t.Fatal("set APPLE_RDMA_DEVICE=<name> for APPLE_RDMA_HARDWARE_MATRIX=1")
	}
	hardwareLog("env", "device=%q probe=%q oneshot=%q", wantDevice, os.Getenv("APPLE_RDMA_PROBE"), os.Getenv("APPLE_RDMA_ONESHOT"))

	list, err := OpenDeviceList()
	if err != nil {
		t.Fatalf("OpenDeviceList: %v", err)
	}
	defer list.Close()
	devs := list.Devices()
	if len(devs) == 0 {
		t.Fatal("no RDMA devices")
	}
	var dev Device
	for _, candidate := range devs {
		if candidate.Name == wantDevice {
			dev = candidate
			break
		}
	}
	if dev.Handle == 0 {
		t.Fatalf("APPLE_RDMA_DEVICE=%q not found; available devices: %v", wantDevice, deviceNames(devs))
	}

	var ctx RDMAContext
	var pd RDMAPD
	var mr RDMAMR
	var cq RDMACQ
	var qp RDMAQP
	success := false
	defer func() {
		if qp != 0 {
			if rc, err := IbvDestroyQp(qp); err != nil || rc != 0 {
				t.Errorf("cleanup ibv_destroy_qp: rc=%d err=%v", rc, err)
			}
		}
		if mr != 0 {
			if rc, err := IbvDeregMr(mr); err != nil || rc != 0 {
				t.Errorf("cleanup ibv_dereg_mr: rc=%d err=%v", rc, err)
			}
		}
		if cq != 0 {
			if rc, err := IbvDestroyCq(cq); err != nil || rc != 0 {
				t.Errorf("cleanup ibv_destroy_cq: rc=%d err=%v", rc, err)
			}
		}
		if pd != 0 {
			if rc, err := IbvDeallocPd(pd); err != nil || rc != 0 {
				t.Errorf("cleanup ibv_dealloc_pd: rc=%d err=%v", rc, err)
			}
		}
		if ctx != 0 {
			if rc, err := IbvCloseDevice(ctx); err != nil || rc != 0 {
				t.Errorf("cleanup ibv_close_device: rc=%d err=%v", rc, err)
			}
		}
		if !success {
			hardwareLog("stop", "device=%s", wantDevice)
		}
	}()

	ctx, err = hardwareStepHandle("ibv_open_device", func() (RDMAContext, error) {
		return dev.Open()
	})
	if err != nil {
		t.Fatalf("ibv_open_device: %v", err)
	}

	deviceAttr := make([]byte, 512)
	if err := hardwareStepRC("ibv_query_device", func() (int, error) {
		return IbvQueryDeviceBytes(ctx, deviceAttr)
	}); err != nil {
		t.Fatalf("ibv_query_device: %v", err)
	}

	var port IbvPortAttr
	if err := hardwareStepRC("ibv_query_port", func() (int, error) {
		return IbvQueryPortAttr(ctx, 1, &port)
	}); err != nil {
		t.Fatalf("ibv_query_port: %v", err)
	}

	pd, err = hardwareStepHandle("ibv_alloc_pd", func() (RDMAPD, error) {
		return IbvAllocPd(ctx)
	})
	if err != nil {
		t.Fatalf("ibv_alloc_pd: %v", err)
	}

	cq, err = hardwareStepHandle("ibv_create_cq", func() (RDMACQ, error) {
		return IbvCreateCq(ctx, 16, 0, 0, 0)
	})
	if err != nil {
		t.Fatalf("ibv_create_cq: %v", err)
	}

	buf, mapBuf, err := hardwareBuffer(4096)
	if err != nil {
		t.Fatalf("hardware buffer: %v", err)
	}
	defer syscall.Munmap(mapBuf)
	mr, err = hardwareStepHandle("ibv_reg_mr", func() (RDMAMR, error) {
		mr, err := IbvRegMr(pd, uintptr(unsafe.Pointer(unsafe.SliceData(buf))), uintptr(len(buf)), IBV_ACCESS_LOCAL_WRITE)
		runtime.KeepAlive(buf)
		return mr, err
	})
	if err != nil {
		t.Fatalf("ibv_reg_mr: %v", err)
	}

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
	qp, err = hardwareStepHandle("ibv_create_qp", func() (RDMAQP, error) {
		return IbvCreateQpAttr(pd, &init)
	})
	if err != nil {
		t.Fatalf("ibv_create_qp: %v", err)
	}
	success = true
}

func deviceNames(devs []Device) []string {
	names := make([]string, 0, len(devs))
	for _, dev := range devs {
		names = append(names, dev.Name)
	}
	return names
}

func hardwareStepHandle[T ~uintptr](name string, fn func() (T, error)) (T, error) {
	hardwareLog(name+".start", "")
	v, err := fn()
	if err != nil {
		hardwareLog(name+".error", "handle=%#x err=%v", uintptr(v), err)
		return v, err
	}
	if v == 0 {
		err = rdmaNilProviderResultError(name, "handle", 0, 0, false, 0, true)
		hardwareLog(name+".error", "handle=0 err=%v", err)
		return v, err
	}
	hardwareLog(name+".ok", "handle=%#x", uintptr(v))
	return v, nil
}

func hardwareStepRC(name string, fn func() (int, error)) error {
	hardwareLog(name+".start", "")
	rc, err := fn()
	if err != nil {
		hardwareLog(name+".error", "rc=%d err=%v", rc, err)
		return err
	}
	if rc != 0 {
		err = fmt.Errorf("provider return rc=%d", rc)
		hardwareLog(name+".error", "rc=%d err=%v", rc, err)
		return err
	}
	hardwareLog(name+".ok", "rc=0")
	return nil
}

func hardwareBuffer(n int) ([]byte, []byte, error) {
	buf, err := syscall.Mmap(-1, 0, n, syscall.PROT_READ|syscall.PROT_WRITE, syscall.MAP_ANON|syscall.MAP_PRIVATE)
	if err != nil {
		return nil, nil, fmt.Errorf("mmap: %w", err)
	}
	for i := range buf {
		buf[i] = byte(i)
	}
	return buf, buf, nil
}

func hardwareLog(event, format string, args ...any) {
	msg := ""
	if format != "" {
		msg = fmt.Sprintf(format, args...)
	}
	fmt.Fprintf(os.Stderr, "rdma_hardware_matrix ts=%s event=%s %s\n", time.Now().UTC().Format(time.RFC3339Nano), event, msg)
}
