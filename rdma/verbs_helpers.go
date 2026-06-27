package rdma

import "unsafe"

// IbvModifyQpToErr moves a queue pair to the ERR state. Outstanding send and
// receive work requests are flushed to the completion queue with a non-success
// status, which lets the caller poll them out and reclaim the queue pair and
// its memory regions without racing in-flight transfers.
func IbvModifyQpToErr(qp RDMAQP) error {
	attr := IbvQPAttr{QPState: IBV_QPS_ERR}
	rc, err := IbvModifyQpAttr(qp, &attr, IBV_QP_STATE)
	if err != nil || rc != 0 {
		return NewModifyQPError(qp, &attr, IBV_QP_STATE, rc, err)
	}
	return nil
}

// IbvCreateQpAttr calls ibv_create_qp with a typed init-attr pointer.
func IbvCreateQpAttr(pd RDMAPD, attr *IbvQPInitAttr) (RDMAQP, error) {
	if attr == nil {
		return 0, rdmaNilPointerError("ibv_create_qp", "qp init attr")
	}
	qp, err := IbvCreateQp(pd, uintptr(unsafe.Pointer(attr)))
	rdmaKeepAlive(attr)
	return qp, err
}

// IbvModifyQpAttr calls ibv_modify_qp with a typed attr pointer.
func IbvModifyQpAttr(qp RDMAQP, attr *IbvQPAttr, attrMask int) (int, error) {
	if attr == nil {
		return 0, rdmaNilPointerError("ibv_modify_qp", "qp attr")
	}
	rc, err := IbvModifyQp(qp, uintptr(unsafe.Pointer(attr)), attrMask)
	rdmaKeepAlive(attr)
	return rc, err
}

// IbvQueryPortAttr calls ibv_query_port with a typed port-attr pointer.
func IbvQueryPortAttr(context RDMAContext, portNum uint8, attr *IbvPortAttr) (int, error) {
	if attr == nil {
		return 0, rdmaNilPointerError("ibv_query_port", "port attr")
	}
	rc, err := IbvQueryPort(context, portNum, uintptr(unsafe.Pointer(attr)))
	rdmaKeepAlive(attr)
	return rc, err
}

// IbvQueryGidInto calls ibv_query_gid with a typed GID pointer.
func IbvQueryGidInto(context RDMAContext, portNum uint8, index int, gid *IbvGID) (int, error) {
	if gid == nil {
		return 0, rdmaNilPointerError("ibv_query_gid", "gid")
	}
	rc, err := IbvQueryGid(context, portNum, index, uintptr(unsafe.Pointer(gid)))
	rdmaKeepAlive(gid)
	return rc, err
}

// IbvQueryDeviceBytes calls ibv_query_device with a byte buffer.
func IbvQueryDeviceBytes(context RDMAContext, buf []byte) (int, error) {
	if len(buf) == 0 {
		return 0, rdmaNilPointerError("ibv_query_device", "device attr")
	}
	rc, err := IbvQueryDevice(context, uintptr(unsafe.Pointer(unsafe.SliceData(buf))))
	rdmaKeepAlive(buf)
	return rc, err
}
