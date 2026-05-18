// Code generated from Apple documentation for rdma. DO NOT EDIT.

package rdma

import (
	"unsafe"

	"github.com/tmc/apple/objc"
)

// Device describes an RDMA device returned by [Devices].
type Device struct {
	Handle RDMADevice
	Name   string
}

// Available reports whether librdma and the generated probe symbols are available.
func Available() bool {
	return frameworkHandle != 0 &&
		_ibvGetDeviceList != nil &&
		_ibvFreeDeviceList != nil &&
		_ibvGetDeviceName != nil
}

// Devices returns the RDMA devices currently reported by librdma.
//
// It only calls exported lifecycle/probe symbols. It does not enable RDMA and
// may return an empty slice on systems where RDMA over Thunderbolt is disabled.
func Devices() ([]Device, error) {
	var n int32
	list, err := IbvGetDeviceList(uintptr(unsafe.Pointer(&n)))
	if err != nil {
		return nil, err
	}
	if list == 0 {
		return nil, nil
	}
	defer IbvFreeDeviceList(list)

	raw := unsafe.Slice((*RDMADevice)(unsafe.Pointer(list)), int(n))
	out := make([]Device, 0, len(raw))
	for _, dev := range raw {
		if dev == 0 {
			continue
		}
		namePtr, err := IbvGetDeviceName(dev)
		if err != nil {
			return nil, err
		}
		out = append(out, Device{
			Handle: dev,
			Name:   objc.GoString((*byte)(unsafe.Pointer(namePtr))),
		})
	}
	return out, nil
}
