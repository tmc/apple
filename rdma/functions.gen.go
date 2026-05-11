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

var _ibv_alloc_pd func(context RDMAContext) RDMAPD
var _ibv_alloc_pdErr error

func tryIbv_alloc_pd(context RDMAContext) (RDMAPD, error) {
	if _ibv_alloc_pd == nil {
		return *new(RDMAPD), symbolCallError("ibv_alloc_pd", "", _ibv_alloc_pdErr)
	}
	return _ibv_alloc_pd(context), nil
}

// Ibv_alloc_pd.
//
// See: https://developer.apple.com/documentation/RDMA/ibv_alloc_pd
func Ibv_alloc_pd(context RDMAContext) (RDMAPD, error) {
	return tryIbv_alloc_pd(context)
}

var _ibv_close_device func(context RDMAContext) int
var _ibv_close_deviceErr error

func tryIbv_close_device(context RDMAContext) (int, error) {
	if _ibv_close_device == nil {
		return 0, symbolCallError("ibv_close_device", "", _ibv_close_deviceErr)
	}
	return _ibv_close_device(context), nil
}

// Ibv_close_device.
//
// See: https://developer.apple.com/documentation/RDMA/ibv_close_device
func Ibv_close_device(context RDMAContext) (int, error) {
	return tryIbv_close_device(context)
}

var _ibv_create_cq func(context RDMAContext, cqe int, cq_context uintptr, channel uintptr, comp_vector int) RDMACQ
var _ibv_create_cqErr error

func tryIbv_create_cq(context RDMAContext, cqe int, cq_context uintptr, channel uintptr, comp_vector int) (RDMACQ, error) {
	if _ibv_create_cq == nil {
		return *new(RDMACQ), symbolCallError("ibv_create_cq", "", _ibv_create_cqErr)
	}
	return _ibv_create_cq(context, cqe, cq_context, channel, comp_vector), nil
}

// Ibv_create_cq.
//
// See: https://developer.apple.com/documentation/RDMA/ibv_create_cq
func Ibv_create_cq(context RDMAContext, cqe int, cq_context uintptr, channel uintptr, comp_vector int) (RDMACQ, error) {
	return tryIbv_create_cq(context, cqe, cq_context, channel, comp_vector)
}

var _ibv_create_qp func(pd RDMAPD, qp_init_attr uintptr) RDMAQP
var _ibv_create_qpErr error

func tryIbv_create_qp(pd RDMAPD, qp_init_attr uintptr) (RDMAQP, error) {
	if _ibv_create_qp == nil {
		return *new(RDMAQP), symbolCallError("ibv_create_qp", "", _ibv_create_qpErr)
	}
	return _ibv_create_qp(pd, qp_init_attr), nil
}

// Ibv_create_qp.
//
// See: https://developer.apple.com/documentation/RDMA/ibv_create_qp
func Ibv_create_qp(pd RDMAPD, qp_init_attr uintptr) (RDMAQP, error) {
	return tryIbv_create_qp(pd, qp_init_attr)
}

var _ibv_dealloc_pd func(pd RDMAPD) int
var _ibv_dealloc_pdErr error

func tryIbv_dealloc_pd(pd RDMAPD) (int, error) {
	if _ibv_dealloc_pd == nil {
		return 0, symbolCallError("ibv_dealloc_pd", "", _ibv_dealloc_pdErr)
	}
	return _ibv_dealloc_pd(pd), nil
}

// Ibv_dealloc_pd.
//
// See: https://developer.apple.com/documentation/RDMA/ibv_dealloc_pd
func Ibv_dealloc_pd(pd RDMAPD) (int, error) {
	return tryIbv_dealloc_pd(pd)
}

var _ibv_dereg_mr func(mr RDMAMR) int
var _ibv_dereg_mrErr error

func tryIbv_dereg_mr(mr RDMAMR) (int, error) {
	if _ibv_dereg_mr == nil {
		return 0, symbolCallError("ibv_dereg_mr", "", _ibv_dereg_mrErr)
	}
	return _ibv_dereg_mr(mr), nil
}

// Ibv_dereg_mr.
//
// See: https://developer.apple.com/documentation/RDMA/ibv_dereg_mr
func Ibv_dereg_mr(mr RDMAMR) (int, error) {
	return tryIbv_dereg_mr(mr)
}

var _ibv_destroy_cq func(cq RDMACQ) int
var _ibv_destroy_cqErr error

func tryIbv_destroy_cq(cq RDMACQ) (int, error) {
	if _ibv_destroy_cq == nil {
		return 0, symbolCallError("ibv_destroy_cq", "", _ibv_destroy_cqErr)
	}
	return _ibv_destroy_cq(cq), nil
}

// Ibv_destroy_cq.
//
// See: https://developer.apple.com/documentation/RDMA/ibv_destroy_cq
func Ibv_destroy_cq(cq RDMACQ) (int, error) {
	return tryIbv_destroy_cq(cq)
}

var _ibv_destroy_qp func(qp RDMAQP) int
var _ibv_destroy_qpErr error

func tryIbv_destroy_qp(qp RDMAQP) (int, error) {
	if _ibv_destroy_qp == nil {
		return 0, symbolCallError("ibv_destroy_qp", "", _ibv_destroy_qpErr)
	}
	return _ibv_destroy_qp(qp), nil
}

// Ibv_destroy_qp.
//
// See: https://developer.apple.com/documentation/RDMA/ibv_destroy_qp
func Ibv_destroy_qp(qp RDMAQP) (int, error) {
	return tryIbv_destroy_qp(qp)
}

var _ibv_free_device_list func(list RDMADeviceList)
var _ibv_free_device_listErr error

func tryIbv_free_device_list(list RDMADeviceList) error {
	if _ibv_free_device_list == nil {
		return symbolCallError("ibv_free_device_list", "", _ibv_free_device_listErr)
	}
	_ibv_free_device_list(list)
	return nil
}

// Ibv_free_device_list.
//
// See: https://developer.apple.com/documentation/RDMA/ibv_free_device_list
func Ibv_free_device_list(list RDMADeviceList) error {
	return tryIbv_free_device_list(list)
}

var _ibv_get_device_list func(num_devices uintptr) RDMADeviceList
var _ibv_get_device_listErr error

func tryIbv_get_device_list(num_devices uintptr) (RDMADeviceList, error) {
	if _ibv_get_device_list == nil {
		return *new(RDMADeviceList), symbolCallError("ibv_get_device_list", "", _ibv_get_device_listErr)
	}
	return _ibv_get_device_list(num_devices), nil
}

// Ibv_get_device_list.
//
// See: https://developer.apple.com/documentation/RDMA/ibv_get_device_list
func Ibv_get_device_list(num_devices uintptr) (RDMADeviceList, error) {
	return tryIbv_get_device_list(num_devices)
}

var _ibv_get_device_name func(device RDMADevice) uintptr
var _ibv_get_device_nameErr error

func tryIbv_get_device_name(device RDMADevice) (uintptr, error) {
	if _ibv_get_device_name == nil {
		return 0, symbolCallError("ibv_get_device_name", "", _ibv_get_device_nameErr)
	}
	return _ibv_get_device_name(device), nil
}

// Ibv_get_device_name.
//
// See: https://developer.apple.com/documentation/RDMA/ibv_get_device_name
func Ibv_get_device_name(device RDMADevice) (uintptr, error) {
	return tryIbv_get_device_name(device)
}

var _ibv_modify_qp func(qp RDMAQP, attr uintptr, attr_mask int) int
var _ibv_modify_qpErr error

func tryIbv_modify_qp(qp RDMAQP, attr uintptr, attr_mask int) (int, error) {
	if _ibv_modify_qp == nil {
		return 0, symbolCallError("ibv_modify_qp", "", _ibv_modify_qpErr)
	}
	return _ibv_modify_qp(qp, attr, attr_mask), nil
}

// Ibv_modify_qp.
//
// See: https://developer.apple.com/documentation/RDMA/ibv_modify_qp
func Ibv_modify_qp(qp RDMAQP, attr uintptr, attr_mask int) (int, error) {
	return tryIbv_modify_qp(qp, attr, attr_mask)
}

var _ibv_open_device func(device RDMADevice) RDMAContext
var _ibv_open_deviceErr error

func tryIbv_open_device(device RDMADevice) (RDMAContext, error) {
	if _ibv_open_device == nil {
		return *new(RDMAContext), symbolCallError("ibv_open_device", "", _ibv_open_deviceErr)
	}
	return _ibv_open_device(device), nil
}

// Ibv_open_device.
//
// See: https://developer.apple.com/documentation/RDMA/ibv_open_device
func Ibv_open_device(device RDMADevice) (RDMAContext, error) {
	return tryIbv_open_device(device)
}

var _ibv_query_device func(context RDMAContext, device_attr uintptr) int
var _ibv_query_deviceErr error

func tryIbv_query_device(context RDMAContext, device_attr uintptr) (int, error) {
	if _ibv_query_device == nil {
		return 0, symbolCallError("ibv_query_device", "", _ibv_query_deviceErr)
	}
	return _ibv_query_device(context, device_attr), nil
}

// Ibv_query_device.
//
// See: https://developer.apple.com/documentation/RDMA/ibv_query_device
func Ibv_query_device(context RDMAContext, device_attr uintptr) (int, error) {
	return tryIbv_query_device(context, device_attr)
}

var _ibv_query_port func(context RDMAContext, port_num uint8, port_attr uintptr) int
var _ibv_query_portErr error

func tryIbv_query_port(context RDMAContext, port_num uint8, port_attr uintptr) (int, error) {
	if _ibv_query_port == nil {
		return 0, symbolCallError("ibv_query_port", "", _ibv_query_portErr)
	}
	return _ibv_query_port(context, port_num, port_attr), nil
}

// Ibv_query_port.
//
// See: https://developer.apple.com/documentation/RDMA/ibv_query_port
func Ibv_query_port(context RDMAContext, port_num uint8, port_attr uintptr) (int, error) {
	return tryIbv_query_port(context, port_num, port_attr)
}

var _ibv_reg_mr func(pd RDMAPD, addr uintptr, length uintptr, access int) RDMAMR
var _ibv_reg_mrErr error

func tryIbv_reg_mr(pd RDMAPD, addr uintptr, length uintptr, access int) (RDMAMR, error) {
	if _ibv_reg_mr == nil {
		return *new(RDMAMR), symbolCallError("ibv_reg_mr", "", _ibv_reg_mrErr)
	}
	return _ibv_reg_mr(pd, addr, length, access), nil
}

// Ibv_reg_mr.
//
// See: https://developer.apple.com/documentation/RDMA/ibv_reg_mr
func Ibv_reg_mr(pd RDMAPD, addr uintptr, length uintptr, access int) (RDMAMR, error) {
	return tryIbv_reg_mr(pd, addr, length, access)
}

func init() {
	if frameworkHandle == 0 {
		return
	}
	registerFunc(&_ibv_alloc_pd, &_ibv_alloc_pdErr, frameworkHandle, "ibv_alloc_pd", "")
	registerFunc(&_ibv_close_device, &_ibv_close_deviceErr, frameworkHandle, "ibv_close_device", "")
	registerFunc(&_ibv_create_cq, &_ibv_create_cqErr, frameworkHandle, "ibv_create_cq", "")
	registerFunc(&_ibv_create_qp, &_ibv_create_qpErr, frameworkHandle, "ibv_create_qp", "")
	registerFunc(&_ibv_dealloc_pd, &_ibv_dealloc_pdErr, frameworkHandle, "ibv_dealloc_pd", "")
	registerFunc(&_ibv_dereg_mr, &_ibv_dereg_mrErr, frameworkHandle, "ibv_dereg_mr", "")
	registerFunc(&_ibv_destroy_cq, &_ibv_destroy_cqErr, frameworkHandle, "ibv_destroy_cq", "")
	registerFunc(&_ibv_destroy_qp, &_ibv_destroy_qpErr, frameworkHandle, "ibv_destroy_qp", "")
	registerFunc(&_ibv_free_device_list, &_ibv_free_device_listErr, frameworkHandle, "ibv_free_device_list", "")
	registerFunc(&_ibv_get_device_list, &_ibv_get_device_listErr, frameworkHandle, "ibv_get_device_list", "")
	registerFunc(&_ibv_get_device_name, &_ibv_get_device_nameErr, frameworkHandle, "ibv_get_device_name", "")
	registerFunc(&_ibv_modify_qp, &_ibv_modify_qpErr, frameworkHandle, "ibv_modify_qp", "")
	registerFunc(&_ibv_open_device, &_ibv_open_deviceErr, frameworkHandle, "ibv_open_device", "")
	registerFunc(&_ibv_query_device, &_ibv_query_deviceErr, frameworkHandle, "ibv_query_device", "")
	registerFunc(&_ibv_query_port, &_ibv_query_portErr, frameworkHandle, "ibv_query_port", "")
	registerFunc(&_ibv_reg_mr, &_ibv_reg_mrErr, frameworkHandle, "ibv_reg_mr", "")
}
