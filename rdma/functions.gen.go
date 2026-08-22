// Code generated from Apple documentation for rdma. DO NOT EDIT.

package rdma

import (
	"fmt"

	"github.com/ebitengine/purego"
)

type unavailableSymbolError struct {
	symbol     string
	introduced string
	cause      error
}

func (e *unavailableSymbolError) Error() string {
	if e == nil {
		return ""
	}
	if e.introduced != "" {
		return fmt.Sprintf("rdma: symbol %s unavailable on this system (introduced in macOS %s)", e.symbol, e.introduced)
	}
	return fmt.Sprintf("rdma: symbol %s unavailable on this system", e.symbol)
}

func (e *unavailableSymbolError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.cause
}

func missingSymbolError(name, introduced string, cause error) error {
	return &unavailableSymbolError{
		symbol:     name,
		introduced: introduced,
		cause:      cause,
	}
}

func symbolCallError(name, introduced string, err error) error {
	if err != nil {
		return err
	}
	if frameworkHandle == 0 {
		return fmt.Errorf("rdma: symbol %s unavailable because the framework could not be loaded", name)
	}
	return missingSymbolError(name, introduced, nil)
}

// registerFunc resolves a framework symbol and registers it as a Go function.
func registerFunc(fptr any, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	defer func() {
		if r := recover(); r != nil {
			*errDst = fmt.Errorf("rdma: register symbol %s: %v", name, r)
		}
	}()
	purego.RegisterFunc(fptr, sym)
	*errDst = nil
}

// registerSymbol resolves a framework symbol and stores its raw address.
func registerSymbol(dst *uintptr, errDst *error, handle uintptr, name, introduced string) {
	sym, err := purego.Dlsym(handle, name)
	if err != nil || sym == 0 {
		*errDst = missingSymbolError(name, introduced, err)
		return
	}
	*dst = sym
	*errDst = nil
}

var _ibvAllocPd func(context RDMAContext) RDMAPD
var _ibvAllocPdErr error

func tryIbvAllocPd(context RDMAContext) (RDMAPD, error) {
	if _ibvAllocPd == nil {
		return *new(RDMAPD), symbolCallError("ibv_alloc_pd", "", _ibvAllocPdErr)
	}
	if context == 0 {
		return *new(RDMAPD), rdmaNilHandleError("ibv_alloc_pd", "context")
	}
	rv, errno, errnoSet := rdmaProviderCallWithErrno(func() RDMAPD {
		return _ibvAllocPd(context)
	})
	if rv == 0 {
		return rv, rdmaNilProviderResultError("ibv_alloc_pd", "protection domain", int64(rv), errno, errnoSet, context, context != 0)
	}
	return rv, nil
}

// IbvAllocPd.
func IbvAllocPd(context RDMAContext) (RDMAPD, error) {
	return tryIbvAllocPd(context)
}

var _ibvCloseDevice func(context RDMAContext) int32
var _ibvCloseDeviceErr error

func tryIbvCloseDevice(context RDMAContext) (int32, error) {
	if _ibvCloseDevice == nil {
		return 0, symbolCallError("ibv_close_device", "", _ibvCloseDeviceErr)
	}
	if context == 0 {
		return 0, rdmaNilHandleError("ibv_close_device", "context")
	}
	return _ibvCloseDevice(context), nil
}

// IbvCloseDevice.
func IbvCloseDevice(context RDMAContext) (int32, error) {
	return tryIbvCloseDevice(context)
}

var _ibvCreateCq func(context RDMAContext, cqe int32, cq_context uintptr, channel uintptr, comp_vector int32) RDMACQ
var _ibvCreateCqErr error

func tryIbvCreateCq(context RDMAContext, cqe int32, cq_context uintptr, channel uintptr, comp_vector int32) (RDMACQ, error) {
	if _ibvCreateCq == nil {
		return *new(RDMACQ), symbolCallError("ibv_create_cq", "", _ibvCreateCqErr)
	}
	if context == 0 {
		return *new(RDMACQ), rdmaNilHandleError("ibv_create_cq", "context")
	}
	rv, errno, errnoSet := rdmaProviderCallWithErrno(func() RDMACQ {
		return _ibvCreateCq(context, cqe, cq_context, channel, comp_vector)
	})
	if rv == 0 {
		return rv, rdmaNilProviderResultError("ibv_create_cq", "completion queue", int64(rv), errno, errnoSet, context, context != 0)
	}
	return rv, nil
}

// IbvCreateCq.
func IbvCreateCq(context RDMAContext, cqe int32, cq_context uintptr, channel uintptr, comp_vector int32) (RDMACQ, error) {
	return tryIbvCreateCq(context, cqe, cq_context, channel, comp_vector)
}

var _ibvCreateQp func(pd RDMAPD, qp_init_attr uintptr) RDMAQP
var _ibvCreateQpErr error

func tryIbvCreateQp(pd RDMAPD, qp_init_attr uintptr) (RDMAQP, error) {
	if _ibvCreateQp == nil {
		return *new(RDMAQP), symbolCallError("ibv_create_qp", "", _ibvCreateQpErr)
	}
	if pd == 0 {
		return *new(RDMAQP), rdmaNilHandleError("ibv_create_qp", "protection domain")
	}
	rv, errno, errnoSet := rdmaProviderCallWithErrno(func() RDMAQP {
		return _ibvCreateQp(pd, qp_init_attr)
	})
	if rv == 0 {
		return rv, rdmaNilProviderResultError("ibv_create_qp", "queue pair", int64(rv), errno, errnoSet, 0, false)
	}
	return rv, nil
}

// IbvCreateQp.
func IbvCreateQp(pd RDMAPD, qp_init_attr uintptr) (RDMAQP, error) {
	return tryIbvCreateQp(pd, qp_init_attr)
}

var _ibvDeallocPd func(pd RDMAPD) int32
var _ibvDeallocPdErr error

func tryIbvDeallocPd(pd RDMAPD) (int32, error) {
	if _ibvDeallocPd == nil {
		return 0, symbolCallError("ibv_dealloc_pd", "", _ibvDeallocPdErr)
	}
	if pd == 0 {
		return 0, rdmaNilHandleError("ibv_dealloc_pd", "protection domain")
	}
	return _ibvDeallocPd(pd), nil
}

// IbvDeallocPd.
func IbvDeallocPd(pd RDMAPD) (int32, error) {
	return tryIbvDeallocPd(pd)
}

var _ibvDeregMr func(mr RDMAMR) int32
var _ibvDeregMrErr error

func tryIbvDeregMr(mr RDMAMR) (int32, error) {
	if _ibvDeregMr == nil {
		return 0, symbolCallError("ibv_dereg_mr", "", _ibvDeregMrErr)
	}
	if mr == 0 {
		return 0, rdmaNilHandleError("ibv_dereg_mr", "memory region")
	}
	return _ibvDeregMr(mr), nil
}

// IbvDeregMr.
func IbvDeregMr(mr RDMAMR) (int32, error) {
	return tryIbvDeregMr(mr)
}

var _ibvDestroyCq func(cq RDMACQ) int32
var _ibvDestroyCqErr error

func tryIbvDestroyCq(cq RDMACQ) (int32, error) {
	if _ibvDestroyCq == nil {
		return 0, symbolCallError("ibv_destroy_cq", "", _ibvDestroyCqErr)
	}
	if cq == 0 {
		return 0, rdmaNilHandleError("ibv_destroy_cq", "completion queue")
	}
	return _ibvDestroyCq(cq), nil
}

// IbvDestroyCq.
func IbvDestroyCq(cq RDMACQ) (int32, error) {
	return tryIbvDestroyCq(cq)
}

var _ibvDestroyQp func(qp RDMAQP) int32
var _ibvDestroyQpErr error

func tryIbvDestroyQp(qp RDMAQP) (int32, error) {
	if _ibvDestroyQp == nil {
		return 0, symbolCallError("ibv_destroy_qp", "", _ibvDestroyQpErr)
	}
	if qp == 0 {
		return 0, rdmaNilHandleError("ibv_destroy_qp", "queue pair")
	}
	return _ibvDestroyQp(qp), nil
}

// IbvDestroyQp.
func IbvDestroyQp(qp RDMAQP) (int32, error) {
	return tryIbvDestroyQp(qp)
}

var _ibvFreeDeviceList func(list RDMADeviceList)
var _ibvFreeDeviceListErr error

func tryIbvFreeDeviceList(list RDMADeviceList) error {
	if _ibvFreeDeviceList == nil {
		return symbolCallError("ibv_free_device_list", "", _ibvFreeDeviceListErr)
	}
	if list == 0 {
		return nil
	}
	_ibvFreeDeviceList(list)
	return nil
}

// IbvFreeDeviceList.
func IbvFreeDeviceList(list RDMADeviceList) error {
	return tryIbvFreeDeviceList(list)
}

var _ibvGetDeviceList func(num_devices uintptr) RDMADeviceList
var _ibvGetDeviceListErr error

func tryIbvGetDeviceList(num_devices uintptr) (RDMADeviceList, error) {
	if _ibvGetDeviceList == nil {
		return *new(RDMADeviceList), symbolCallError("ibv_get_device_list", "", _ibvGetDeviceListErr)
	}
	rv, errno, errnoSet := rdmaProviderCallWithErrno(func() RDMADeviceList {
		return _ibvGetDeviceList(num_devices)
	})
	if rv == 0 {
		return rv, rdmaNilProviderResultError("ibv_get_device_list", "device list", int64(rv), errno, errnoSet, 0, false)
	}
	return rv, nil
}

// IbvGetDeviceList.
func IbvGetDeviceList(num_devices uintptr) (RDMADeviceList, error) {
	return tryIbvGetDeviceList(num_devices)
}

var _ibvGetDeviceName func(device RDMADevice) uintptr
var _ibvGetDeviceNameErr error

func tryIbvGetDeviceName(device RDMADevice) (uintptr, error) {
	if _ibvGetDeviceName == nil {
		return 0, symbolCallError("ibv_get_device_name", "", _ibvGetDeviceNameErr)
	}
	if device == 0 {
		return 0, rdmaNilHandleError("ibv_get_device_name", "device")
	}
	return _ibvGetDeviceName(device), nil
}

// IbvGetDeviceName.
func IbvGetDeviceName(device RDMADevice) (uintptr, error) {
	return tryIbvGetDeviceName(device)
}

var _ibvModifyQp func(qp RDMAQP, attr uintptr, attr_mask int32) int32
var _ibvModifyQpErr error

func tryIbvModifyQp(qp RDMAQP, attr uintptr, attr_mask int32) (int32, error) {
	if _ibvModifyQp == nil {
		return 0, symbolCallError("ibv_modify_qp", "", _ibvModifyQpErr)
	}
	if qp == 0 {
		return 0, rdmaNilHandleError("ibv_modify_qp", "queue pair")
	}
	rv, errno, errnoSet := rdmaProviderCallWithErrno(func() int32 {
		return _ibvModifyQp(qp, attr, attr_mask)
	})
	if rv < 0 {
		return rv, rdmaNegativeProviderReturnError("ibv_modify_qp", int(rv), errno, errnoSet, 0, false)
	}
	return rv, nil
}

// IbvModifyQp.
func IbvModifyQp(qp RDMAQP, attr uintptr, attr_mask int32) (int32, error) {
	return tryIbvModifyQp(qp, attr, attr_mask)
}

var _ibvOpenDevice func(device RDMADevice) RDMAContext
var _ibvOpenDeviceErr error

func tryIbvOpenDevice(device RDMADevice) (RDMAContext, error) {
	if _ibvOpenDevice == nil {
		return *new(RDMAContext), symbolCallError("ibv_open_device", "", _ibvOpenDeviceErr)
	}
	if device == 0 {
		return *new(RDMAContext), rdmaNilHandleError("ibv_open_device", "device")
	}
	rv, errno, errnoSet := rdmaProviderCallWithErrno(func() RDMAContext {
		return _ibvOpenDevice(device)
	})
	if rv == 0 {
		return rv, rdmaNilProviderResultError("ibv_open_device", "context", int64(rv), errno, errnoSet, 0, false)
	}
	return rv, nil
}

// IbvOpenDevice.
func IbvOpenDevice(device RDMADevice) (RDMAContext, error) {
	return tryIbvOpenDevice(device)
}

var _ibvQueryDevice func(context RDMAContext, device_attr uintptr) int32
var _ibvQueryDeviceErr error

func tryIbvQueryDevice(context RDMAContext, device_attr uintptr) (int32, error) {
	if _ibvQueryDevice == nil {
		return 0, symbolCallError("ibv_query_device", "", _ibvQueryDeviceErr)
	}
	if context == 0 {
		return 0, rdmaNilHandleError("ibv_query_device", "context")
	}
	rv, errno, errnoSet := rdmaProviderCallWithErrno(func() int32 {
		return _ibvQueryDevice(context, device_attr)
	})
	if rv < 0 {
		return rv, rdmaNegativeProviderReturnError("ibv_query_device", int(rv), errno, errnoSet, context, context != 0)
	}
	return rv, nil
}

// IbvQueryDevice.
func IbvQueryDevice(context RDMAContext, device_attr uintptr) (int32, error) {
	return tryIbvQueryDevice(context, device_attr)
}

var _ibvQueryGid func(context RDMAContext, port_num uint8, index int32, gid uintptr) int32
var _ibvQueryGidErr error

func tryIbvQueryGid(context RDMAContext, port_num uint8, index int32, gid uintptr) (int32, error) {
	if _ibvQueryGid == nil {
		return 0, symbolCallError("ibv_query_gid", "", _ibvQueryGidErr)
	}
	if context == 0 {
		return 0, rdmaNilHandleError("ibv_query_gid", "context")
	}
	rv, errno, errnoSet := rdmaProviderCallWithErrno(func() int32 {
		return _ibvQueryGid(context, port_num, index, gid)
	})
	if rv < 0 {
		return rv, rdmaNegativeProviderReturnError("ibv_query_gid", int(rv), errno, errnoSet, context, context != 0)
	}
	return rv, nil
}

// IbvQueryGid.
func IbvQueryGid(context RDMAContext, port_num uint8, index int32, gid uintptr) (int32, error) {
	return tryIbvQueryGid(context, port_num, index, gid)
}

var _ibvQueryPort func(context RDMAContext, port_num uint8, port_attr uintptr) int32
var _ibvQueryPortErr error

func tryIbvQueryPort(context RDMAContext, port_num uint8, port_attr uintptr) (int32, error) {
	if _ibvQueryPort == nil {
		return 0, symbolCallError("ibv_query_port", "", _ibvQueryPortErr)
	}
	if context == 0 {
		return 0, rdmaNilHandleError("ibv_query_port", "context")
	}
	rv, errno, errnoSet := rdmaProviderCallWithErrno(func() int32 {
		return _ibvQueryPort(context, port_num, port_attr)
	})
	if rv < 0 {
		return rv, rdmaNegativeProviderReturnError("ibv_query_port", int(rv), errno, errnoSet, context, context != 0)
	}
	return rv, nil
}

// IbvQueryPort.
func IbvQueryPort(context RDMAContext, port_num uint8, port_attr uintptr) (int32, error) {
	return tryIbvQueryPort(context, port_num, port_attr)
}

var _ibvRegMr func(pd RDMAPD, addr uintptr, length uintptr, access int32) RDMAMR
var _ibvRegMrErr error

func tryIbvRegMr(pd RDMAPD, addr uintptr, length uintptr, access int32) (RDMAMR, error) {
	if _ibvRegMr == nil {
		return *new(RDMAMR), symbolCallError("ibv_reg_mr", "", _ibvRegMrErr)
	}
	if pd == 0 {
		return *new(RDMAMR), rdmaNilHandleError("ibv_reg_mr", "protection domain")
	}
	rv, errno, errnoSet := rdmaProviderCallWithErrno(func() RDMAMR {
		return _ibvRegMr(pd, addr, length, access)
	})
	if rv == 0 {
		return rv, rdmaNilProviderResultError("ibv_reg_mr", "memory region", int64(rv), errno, errnoSet, 0, false)
	}
	return rv, nil
}

// IbvRegMr.
func IbvRegMr(pd RDMAPD, addr uintptr, length uintptr, access int32) (RDMAMR, error) {
	return tryIbvRegMr(pd, addr, length, access)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_ibvAllocPd, &_ibvAllocPdErr, frameworkHandle, "ibv_alloc_pd", "")
	registerFunc(&_ibvCloseDevice, &_ibvCloseDeviceErr, frameworkHandle, "ibv_close_device", "")
	registerFunc(&_ibvCreateCq, &_ibvCreateCqErr, frameworkHandle, "ibv_create_cq", "")
	registerFunc(&_ibvCreateQp, &_ibvCreateQpErr, frameworkHandle, "ibv_create_qp", "")
	registerFunc(&_ibvDeallocPd, &_ibvDeallocPdErr, frameworkHandle, "ibv_dealloc_pd", "")
	registerFunc(&_ibvDeregMr, &_ibvDeregMrErr, frameworkHandle, "ibv_dereg_mr", "")
	registerFunc(&_ibvDestroyCq, &_ibvDestroyCqErr, frameworkHandle, "ibv_destroy_cq", "")
	registerFunc(&_ibvDestroyQp, &_ibvDestroyQpErr, frameworkHandle, "ibv_destroy_qp", "")
	registerFunc(&_ibvFreeDeviceList, &_ibvFreeDeviceListErr, frameworkHandle, "ibv_free_device_list", "")
	registerFunc(&_ibvGetDeviceList, &_ibvGetDeviceListErr, frameworkHandle, "ibv_get_device_list", "")
	registerFunc(&_ibvGetDeviceName, &_ibvGetDeviceNameErr, frameworkHandle, "ibv_get_device_name", "")
	registerFunc(&_ibvModifyQp, &_ibvModifyQpErr, frameworkHandle, "ibv_modify_qp", "")
	registerFunc(&_ibvOpenDevice, &_ibvOpenDeviceErr, frameworkHandle, "ibv_open_device", "")
	registerFunc(&_ibvQueryDevice, &_ibvQueryDeviceErr, frameworkHandle, "ibv_query_device", "")
	registerFunc(&_ibvQueryGid, &_ibvQueryGidErr, frameworkHandle, "ibv_query_gid", "")
	registerFunc(&_ibvQueryPort, &_ibvQueryPortErr, frameworkHandle, "ibv_query_port", "")
	registerFunc(&_ibvRegMr, &_ibvRegMrErr, frameworkHandle, "ibv_reg_mr", "")
}
