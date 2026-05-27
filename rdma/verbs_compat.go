package rdma

// Ibv_alloc_pd is kept for callers generated against the earlier RDMA names.
func Ibv_alloc_pd(context RDMAContext) (RDMAPD, error) { return IbvAllocPd(context) }

// Ibv_close_device is kept for callers generated against the earlier RDMA names.
func Ibv_close_device(context RDMAContext) (int, error) { return IbvCloseDevice(context) }

// Ibv_create_cq is kept for callers generated against the earlier RDMA names.
func Ibv_create_cq(context RDMAContext, cqe int, cqContext uintptr, channel uintptr, compVector int) (RDMACQ, error) {
	return IbvCreateCq(context, cqe, cqContext, channel, compVector)
}

// Ibv_create_qp is kept for callers generated against the earlier RDMA names.
func Ibv_create_qp(pd RDMAPD, qpInitAttr uintptr) (RDMAQP, error) {
	return IbvCreateQp(pd, qpInitAttr)
}

// Ibv_dealloc_pd is kept for callers generated against the earlier RDMA names.
func Ibv_dealloc_pd(pd RDMAPD) (int, error) { return IbvDeallocPd(pd) }

// Ibv_dereg_mr is kept for callers generated against the earlier RDMA names.
func Ibv_dereg_mr(mr RDMAMR) (int, error) { return IbvDeregMr(mr) }

// Ibv_destroy_cq is kept for callers generated against the earlier RDMA names.
func Ibv_destroy_cq(cq RDMACQ) (int, error) { return IbvDestroyCq(cq) }

// Ibv_destroy_qp is kept for callers generated against the earlier RDMA names.
func Ibv_destroy_qp(qp RDMAQP) (int, error) { return IbvDestroyQp(qp) }

// Ibv_free_device_list is kept for callers generated against the earlier RDMA names.
func Ibv_free_device_list(list RDMADeviceList) error { return IbvFreeDeviceList(list) }

// Ibv_get_device_list is kept for callers generated against the earlier RDMA names.
func Ibv_get_device_list(numDevices uintptr) (RDMADeviceList, error) {
	return IbvGetDeviceList(numDevices)
}

// Ibv_get_device_name is kept for callers generated against the earlier RDMA names.
func Ibv_get_device_name(device RDMADevice) (uintptr, error) {
	return IbvGetDeviceName(device)
}

// Ibv_modify_qp is kept for callers generated against the earlier RDMA names.
func Ibv_modify_qp(qp RDMAQP, attr uintptr, attrMask int) (int, error) {
	return IbvModifyQp(qp, attr, attrMask)
}

// Ibv_open_device is kept for callers generated against the earlier RDMA names.
func Ibv_open_device(device RDMADevice) (RDMAContext, error) {
	return IbvOpenDevice(device)
}

// Ibv_query_device is kept for callers generated against the earlier RDMA names.
func Ibv_query_device(context RDMAContext, deviceAttr uintptr) (int, error) {
	return IbvQueryDevice(context, deviceAttr)
}

// Ibv_query_gid is kept for callers generated against the earlier RDMA names.
func Ibv_query_gid(context RDMAContext, portNum uint8, index int, gid uintptr) (int, error) {
	return IbvQueryGid(context, portNum, index, gid)
}

// Ibv_query_port is kept for callers generated against the earlier RDMA names.
func Ibv_query_port(context RDMAContext, portNum uint8, portAttr uintptr) (int, error) {
	return IbvQueryPort(context, portNum, portAttr)
}

// Ibv_reg_mr is kept for callers generated against the earlier RDMA names.
func Ibv_reg_mr(pd RDMAPD, addr uintptr, length uintptr, access int) (RDMAMR, error) {
	return IbvRegMr(pd, addr, length, access)
}
