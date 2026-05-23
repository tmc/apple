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

func (e *unavailableSymbolError) Is(target error) bool {
	return target == ErrSymbolUnavailable
}

func missingSymbolError(name, introduced string, cause error) error {
	return &unavailableSymbolError{
		symbol:     name,
		introduced: introduced,
		cause:      cause,
	}
}

func symbolCallError(name, introduced string, err error) error {
	var cause error
	if err != nil {
		cause = err
	} else if frameworkHandle == 0 {
		cause = fmt.Errorf("rdma: framework unavailable")
	} else {
		cause = missingSymbolError(name, introduced, nil)
	}
	return &ProviderError{
		Operation: name,
		Failure:   FailureSymbolUnavailable,
		Cause:     cause,
	}
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
		return 0, rdmaNilHandleError("ibv_alloc_pd", "context")
	}
	pd, errno, errnoSet := rdmaProviderCallWithErrno(func() RDMAPD {
		return _ibvAllocPd(context)
	})
	rdmaKeepAlive(context)
	if pd == 0 {
		return 0, rdmaNilProviderResultError("ibv_alloc_pd", "protection domain", 0, errno, errnoSet, context, true)
	}
	return pd, nil
}

// IbvAllocPd.
func IbvAllocPd(context RDMAContext) (RDMAPD, error) {
	return tryIbvAllocPd(context)
}

var _ibvCloseDevice func(context RDMAContext) int
var _ibvCloseDeviceErr error

func tryIbvCloseDevice(context RDMAContext) (int, error) {
	if _ibvCloseDevice == nil {
		return 0, symbolCallError("ibv_close_device", "", _ibvCloseDeviceErr)
	}
	if context == 0 {
		return 0, rdmaNilHandleError("ibv_close_device", "context")
	}
	rc, errno, errnoSet := rdmaProviderCallWithErrno(func() int {
		return _ibvCloseDevice(context)
	})
	rdmaKeepAlive(context)
	if rc < 0 {
		return rc, rdmaNegativeProviderReturnError("ibv_close_device", rc, errno, errnoSet, context, true)
	}
	if rc == 0 {
		rdmaForgetContext(context)
	}
	return rc, nil
}

// IbvCloseDevice.
func IbvCloseDevice(context RDMAContext) (int, error) {
	return tryIbvCloseDevice(context)
}

var _ibvCreateCq func(context RDMAContext, cqe int, cq_context uintptr, channel uintptr, comp_vector int) RDMACQ
var _ibvCreateCqErr error

func tryIbvCreateCq(context RDMAContext, cqe int, cq_context uintptr, channel uintptr, comp_vector int) (RDMACQ, error) {
	if _ibvCreateCq == nil {
		return *new(RDMACQ), symbolCallError("ibv_create_cq", "", _ibvCreateCqErr)
	}
	if context == 0 {
		return 0, rdmaNilHandleError("ibv_create_cq", "context")
	}
	cq, errno, errnoSet := rdmaProviderCallWithErrno(func() RDMACQ {
		return _ibvCreateCq(context, cqe, cq_context, channel, comp_vector)
	})
	rdmaKeepAlive(context)
	if cq == 0 {
		return 0, rdmaNilProviderResultError("ibv_create_cq", "completion queue", 0, errno, errnoSet, context, true)
	}
	return cq, nil
}

// IbvCreateCq.
func IbvCreateCq(context RDMAContext, cqe int, cq_context uintptr, channel uintptr, comp_vector int) (RDMACQ, error) {
	return tryIbvCreateCq(context, cqe, cq_context, channel, comp_vector)
}

var _ibvCreateQp func(pd RDMAPD, qp_init_attr uintptr) RDMAQP
var _ibvCreateQpErr error

func tryIbvCreateQp(pd RDMAPD, qp_init_attr uintptr) (RDMAQP, error) {
	if _ibvCreateQp == nil {
		return *new(RDMAQP), symbolCallError("ibv_create_qp", "", _ibvCreateQpErr)
	}
	if pd == 0 {
		return 0, rdmaNilHandleError("ibv_create_qp", "protection domain")
	}
	if qp_init_attr == 0 {
		return 0, rdmaNilPointerError("ibv_create_qp", "qp init attr")
	}
	qp, errno, errnoSet := rdmaProviderCallWithErrno(func() RDMAQP {
		return _ibvCreateQp(pd, qp_init_attr)
	})
	rdmaKeepAlive(pd)
	if qp == 0 {
		return 0, rdmaNilProviderResultError("ibv_create_qp", "queue pair", 0, errno, errnoSet, 0, true)
	}
	return qp, nil
}

// IbvCreateQp.
func IbvCreateQp(pd RDMAPD, qp_init_attr uintptr) (RDMAQP, error) {
	return tryIbvCreateQp(pd, qp_init_attr)
}

var _ibvDeallocPd func(pd RDMAPD) int
var _ibvDeallocPdErr error

func tryIbvDeallocPd(pd RDMAPD) (int, error) {
	if _ibvDeallocPd == nil {
		return 0, symbolCallError("ibv_dealloc_pd", "", _ibvDeallocPdErr)
	}
	if pd == 0 {
		return 0, rdmaNilHandleError("ibv_dealloc_pd", "protection domain")
	}
	rc, errno, errnoSet := rdmaProviderCallWithErrno(func() int {
		return _ibvDeallocPd(pd)
	})
	rdmaKeepAlive(pd)
	if rc < 0 {
		return rc, rdmaNegativeProviderReturnError("ibv_dealloc_pd", rc, errno, errnoSet, 0, true)
	}
	return rc, nil
}

// IbvDeallocPd.
func IbvDeallocPd(pd RDMAPD) (int, error) {
	return tryIbvDeallocPd(pd)
}

var _ibvDeregMr func(mr RDMAMR) int
var _ibvDeregMrErr error

func tryIbvDeregMr(mr RDMAMR) (int, error) {
	if _ibvDeregMr == nil {
		return 0, symbolCallError("ibv_dereg_mr", "", _ibvDeregMrErr)
	}
	if mr == 0 {
		return 0, rdmaNilHandleError("ibv_dereg_mr", "memory region")
	}
	rc, errno, errnoSet := rdmaProviderCallWithErrno(func() int {
		return _ibvDeregMr(mr)
	})
	rdmaKeepAlive(mr)
	if rc < 0 {
		return rc, rdmaNegativeProviderReturnError("ibv_dereg_mr", rc, errno, errnoSet, 0, true)
	}
	return rc, nil
}

// IbvDeregMr.
func IbvDeregMr(mr RDMAMR) (int, error) {
	return tryIbvDeregMr(mr)
}

var _ibvDestroyCq func(cq RDMACQ) int
var _ibvDestroyCqErr error

func tryIbvDestroyCq(cq RDMACQ) (int, error) {
	if _ibvDestroyCq == nil {
		return 0, symbolCallError("ibv_destroy_cq", "", _ibvDestroyCqErr)
	}
	if cq == 0 {
		return 0, rdmaNilHandleError("ibv_destroy_cq", "completion queue")
	}
	rc, errno, errnoSet := rdmaProviderCallWithErrno(func() int {
		return _ibvDestroyCq(cq)
	})
	rdmaKeepAlive(cq)
	if rc < 0 {
		return rc, rdmaNegativeProviderReturnError("ibv_destroy_cq", rc, errno, errnoSet, 0, true)
	}
	return rc, nil
}

// IbvDestroyCq.
func IbvDestroyCq(cq RDMACQ) (int, error) {
	return tryIbvDestroyCq(cq)
}

var _ibvDestroyQp func(qp RDMAQP) int
var _ibvDestroyQpErr error

func tryIbvDestroyQp(qp RDMAQP) (int, error) {
	if _ibvDestroyQp == nil {
		return 0, symbolCallError("ibv_destroy_qp", "", _ibvDestroyQpErr)
	}
	if qp == 0 {
		return 0, rdmaNilHandleError("ibv_destroy_qp", "queue pair")
	}
	rc, errno, errnoSet := rdmaProviderCallWithErrno(func() int {
		return _ibvDestroyQp(qp)
	})
	rdmaKeepAlive(qp)
	if rc < 0 {
		return rc, rdmaNegativeProviderReturnError("ibv_destroy_qp", rc, errno, errnoSet, 0, true)
	}
	return rc, nil
}

// IbvDestroyQp.
func IbvDestroyQp(qp RDMAQP) (int, error) {
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
	rdmaProviderCall0(func() {
		_ibvFreeDeviceList(list)
	})
	rdmaKeepAlive(list)
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
	list := rdmaProviderCall(func() RDMADeviceList {
		return _ibvGetDeviceList(num_devices)
	})
	return list, nil
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
	name, errno, errnoSet := rdmaProviderCallWithErrno(func() uintptr {
		return _ibvGetDeviceName(device)
	})
	rdmaKeepAlive(device)
	if name == 0 {
		return 0, rdmaNilProviderResultError("ibv_get_device_name", "device name", 0, errno, errnoSet, 0, false)
	}
	return name, nil
}

// IbvGetDeviceName.
func IbvGetDeviceName(device RDMADevice) (uintptr, error) {
	return tryIbvGetDeviceName(device)
}

var _ibvModifyQp func(qp RDMAQP, attr uintptr, attr_mask int) int
var _ibvModifyQpErr error

func tryIbvModifyQp(qp RDMAQP, attr uintptr, attr_mask int) (int, error) {
	if _ibvModifyQp == nil {
		return 0, symbolCallError("ibv_modify_qp", "", _ibvModifyQpErr)
	}
	if qp == 0 {
		return 0, rdmaNilHandleError("ibv_modify_qp", "queue pair")
	}
	if attr == 0 {
		return 0, rdmaNilPointerError("ibv_modify_qp", "qp attr")
	}
	rc, errno, errnoSet := rdmaProviderCallWithErrno(func() int {
		return _ibvModifyQp(qp, attr, attr_mask)
	})
	rdmaKeepAlive(qp)
	if rc < 0 {
		return rc, rdmaNegativeProviderReturnError("ibv_modify_qp", rc, errno, errnoSet, 0, true)
	}
	return rc, nil
}

// IbvModifyQp.
func IbvModifyQp(qp RDMAQP, attr uintptr, attr_mask int) (int, error) {
	return tryIbvModifyQp(qp, attr, attr_mask)
}

var _ibvOpenDevice func(device RDMADevice) RDMAContext
var _ibvOpenDeviceErr error

func tryIbvOpenDevice(device RDMADevice) (RDMAContext, error) {
	if _ibvOpenDevice == nil {
		return *new(RDMAContext), symbolCallError("ibv_open_device", "", _ibvOpenDeviceErr)
	}
	if device == 0 {
		return 0, rdmaNilHandleError("ibv_open_device", "device")
	}
	context, errno, errnoSet := rdmaProviderCallWithErrno(func() RDMAContext {
		return _ibvOpenDevice(device)
	})
	rdmaKeepAlive(device)
	if context == 0 {
		return 0, rdmaNilProviderResultError("ibv_open_device", "context", 0, errno, errnoSet, 0, false)
	}
	return context, nil
}

// IbvOpenDevice.
func IbvOpenDevice(device RDMADevice) (RDMAContext, error) {
	return tryIbvOpenDevice(device)
}

var _ibvQueryDevice func(context RDMAContext, device_attr uintptr) int
var _ibvQueryDeviceErr error

func tryIbvQueryDevice(context RDMAContext, device_attr uintptr) (int, error) {
	if _ibvQueryDevice == nil {
		return 0, symbolCallError("ibv_query_device", "", _ibvQueryDeviceErr)
	}
	if context == 0 {
		return 0, rdmaNilHandleError("ibv_query_device", "context")
	}
	if device_attr == 0 {
		return 0, rdmaNilPointerError("ibv_query_device", "device attr")
	}
	rc, errno, errnoSet := rdmaProviderCallWithErrno(func() int {
		return _ibvQueryDevice(context, device_attr)
	})
	rdmaKeepAlive(context)
	if rc < 0 {
		return rc, rdmaNegativeProviderReturnError("ibv_query_device", rc, errno, errnoSet, context, true)
	}
	return rc, nil
}

// IbvQueryDevice.
func IbvQueryDevice(context RDMAContext, device_attr uintptr) (int, error) {
	return tryIbvQueryDevice(context, device_attr)
}

var _ibvQueryGid func(context RDMAContext, port_num uint8, index int, gid uintptr) int
var _ibvQueryGidErr error

func tryIbvQueryGid(context RDMAContext, port_num uint8, index int, gid uintptr) (int, error) {
	if _ibvQueryGid == nil {
		return 0, symbolCallError("ibv_query_gid", "", _ibvQueryGidErr)
	}
	if context == 0 {
		return 0, rdmaNilHandleError("ibv_query_gid", "context")
	}
	if gid == 0 {
		return 0, rdmaNilPointerError("ibv_query_gid", "gid")
	}
	rc, errno, errnoSet := rdmaProviderCallWithErrno(func() int {
		return _ibvQueryGid(context, port_num, index, gid)
	})
	rdmaKeepAlive(context)
	if rc < 0 {
		return rc, rdmaNegativeProviderReturnError("ibv_query_gid", rc, errno, errnoSet, context, true)
	}
	return rc, nil
}

// IbvQueryGid.
func IbvQueryGid(context RDMAContext, port_num uint8, index int, gid uintptr) (int, error) {
	return tryIbvQueryGid(context, port_num, index, gid)
}

var _ibvQueryPort func(context RDMAContext, port_num uint8, port_attr uintptr) int
var _ibvQueryPortErr error

func tryIbvQueryPort(context RDMAContext, port_num uint8, port_attr uintptr) (int, error) {
	if _ibvQueryPort == nil {
		return 0, symbolCallError("ibv_query_port", "", _ibvQueryPortErr)
	}
	if context == 0 {
		return 0, rdmaNilHandleError("ibv_query_port", "context")
	}
	if port_attr == 0 {
		return 0, rdmaNilPointerError("ibv_query_port", "port attr")
	}
	rc, errno, errnoSet := rdmaProviderCallWithErrno(func() int {
		return _ibvQueryPort(context, port_num, port_attr)
	})
	rdmaKeepAlive(context)
	if rc < 0 {
		return rc, rdmaNegativeProviderReturnError("ibv_query_port", rc, errno, errnoSet, context, true)
	}
	return rc, nil
}

// IbvQueryPort.
func IbvQueryPort(context RDMAContext, port_num uint8, port_attr uintptr) (int, error) {
	return tryIbvQueryPort(context, port_num, port_attr)
}

var _ibvRegMr func(pd RDMAPD, addr uintptr, length uintptr, access int) RDMAMR
var _ibvRegMrErr error

func tryIbvRegMr(pd RDMAPD, addr uintptr, length uintptr, access int) (RDMAMR, error) {
	if _ibvRegMr == nil {
		return *new(RDMAMR), symbolCallError("ibv_reg_mr", "", _ibvRegMrErr)
	}
	if pd == 0 {
		return 0, rdmaNilHandleError("ibv_reg_mr", "protection domain")
	}
	if addr == 0 && length != 0 {
		return 0, rdmaNilPointerError("ibv_reg_mr", "memory region")
	}
	mr, errno, errnoSet := rdmaProviderCallWithErrno(func() RDMAMR {
		return _ibvRegMr(pd, addr, length, access)
	})
	rdmaKeepAlive(pd)
	if mr == 0 {
		return 0, rdmaNilProviderResultError("ibv_reg_mr", "memory region", 0, errno, errnoSet, 0, true)
	}
	return mr, nil
}

// IbvRegMr.
func IbvRegMr(pd RDMAPD, addr uintptr, length uintptr, access int) (RDMAMR, error) {
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
