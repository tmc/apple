package rdma

import (
	"fmt"
	"sync"
	"unsafe"

	"github.com/tmc/apple/objc"
)

// DeviceList owns an ibv_get_device_list result.
//
// Device handles returned by [DeviceList.Devices] remain valid until Close is
// called. Open devices before closing the list.
type DeviceList struct {
	mu      sync.Mutex
	list    RDMADeviceList
	devices []Device
}

// OpenDeviceList returns the RDMA device list currently reported by librdma.
func OpenDeviceList() (*DeviceList, error) {
	var n int32
	list, err := IbvGetDeviceList(uintptr(unsafe.Pointer(&n)))
	rdmaKeepAlive(&n)
	if err != nil {
		return nil, err
	}

	dl := &DeviceList{list: list}
	rdmaSetFinalizer(dl, (*DeviceList).finalize)
	if list == 0 {
		return dl, nil
	}
	if n < 0 {
		_ = dl.Close()
		return nil, fmt.Errorf("rdma: ibv_get_device_list returned negative device count %d", n)
	}

	raw, err := rdmaDeviceListEntries(list, n)
	if err != nil {
		_ = dl.Close()
		return nil, err
	}
	dl.devices = make([]Device, 0, len(raw))
	for _, dev := range raw {
		if dev == 0 {
			continue
		}
		namePtr, err := IbvGetDeviceName(dev)
		if err != nil {
			_ = dl.Close()
			return nil, err
		}
		dl.devices = append(dl.devices, Device{
			Handle: dev,
			Name:   objc.GoString((*byte)(unsafe.Pointer(namePtr))),
			list:   dl,
		})
	}
	return dl, nil
}

func rdmaDeviceListEntries(list RDMADeviceList, n int32) ([]RDMADevice, error) {
	if list == 0 {
		return nil, nil
	}
	if n < 0 {
		return nil, fmt.Errorf("rdma: ibv_get_device_list returned negative device count %d", n)
	}
	if n > 0 {
		return unsafe.Slice((*RDMADevice)(unsafe.Pointer(list)), int(n)), nil
	}
	const maxDeviceListScan = 1024
	ptrs := (*[maxDeviceListScan]RDMADevice)(unsafe.Pointer(list))
	for i, dev := range ptrs {
		if dev == 0 {
			return ptrs[:i:i], nil
		}
	}
	return nil, fmt.Errorf("rdma: ibv_get_device_list returned unterminated device list")
}

// Devices returns a copy of the devices in the list.
func (l *DeviceList) Devices() []Device {
	if l == nil {
		return nil
	}
	l.mu.Lock()
	defer l.mu.Unlock()
	out := make([]Device, len(l.devices))
	copy(out, l.devices)
	return out
}

// Close releases the underlying device list. It is safe to call Close more than
// once.
func (l *DeviceList) Close() error {
	if l == nil {
		return nil
	}
	l.mu.Lock()
	list := l.list
	l.list = 0
	l.devices = nil
	l.mu.Unlock()
	if list == 0 {
		return nil
	}
	rdmaSetFinalizer(l, nil)
	return IbvFreeDeviceList(list)
}

func (l *DeviceList) finalize() {
	_ = l.Close()
}

// Open opens d with ibv_open_device and keeps the owning list alive for the
// duration of the provider call.
func (d Device) Open() (RDMAContext, error) {
	context, err := IbvOpenDevice(d.Handle)
	rdmaKeepAlive(d.list)
	return context, err
}
