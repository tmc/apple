package rdma

import (
	"fmt"

	"github.com/tmc/apple/rdma"
)

// RemoteQP describes the remote side of a UC queue-pair connection.
type RemoteQP struct {
	Name      string
	LID       uint16
	QPN       uint32
	PSN       uint32
	GIDIndex  int
	GID       rdma.IbvGID
	UseGlobal bool
	ActiveMTU int32
}

// LocalQP describes the local side of a UC queue-pair connection.
type LocalQP struct {
	PortNum   uint8
	GIDIndex  int
	ActiveMTU int32
	LinkLayer uint8
}

// RTRPolicy controls provider-specific address-vector choices.
type RTRPolicy struct {
	ZeroDLIDWhenGlobal bool
	HopLimit           uint8
	TrafficClass       uint8
	FlowLabel          uint32
}

// RTRAttr returns the ibv_modify_qp attributes and mask for INIT->RTR.
func RTRAttr(local LocalQP, remote RemoteQP, policy RTRPolicy) (rdma.IbvQPAttr, int, error) {
	port := local.PortNum
	if port == 0 {
		port = 1
	}
	attr := rdma.IbvQPAttr{
		QPState:   rdma.IBV_QPS_RTR,
		PathMTU:   NegotiatedPathMTU(local.ActiveMTU, remote.ActiveMTU),
		RQPSN:     remote.PSN,
		DestQPNum: remote.QPN,
		AHAttr: rdma.IbvAHAttr{
			DLID:     remote.LID,
			PortNum:  port,
			IsGlobal: boolByte(remote.UseGlobal),
		},
	}
	if !remote.UseGlobal {
		return attr, RTRAttrMask, nil
	}
	if local.GIDIndex < 0 || local.GIDIndex > 255 {
		return attr, RTRAttrMask, GIDIndexRangeError(local.GIDIndex)
	}
	if policy.ZeroDLIDWhenGlobal {
		attr.AHAttr.DLID = 0
	}
	attr.AHAttr.GRH.DGID = remote.GID
	attr.AHAttr.GRH.SGIDIndex = uint8(local.GIDIndex)
	attr.AHAttr.GRH.HopLimit = policy.HopLimit
	if attr.AHAttr.GRH.HopLimit == 0 {
		attr.AHAttr.GRH.HopLimit = 1
	}
	attr.AHAttr.GRH.TrafficClass = policy.TrafficClass
	attr.AHAttr.GRH.FlowLabel = policy.FlowLabel
	return attr, RTRAttrMask, nil
}

// RTRAttrMask is the ibv_modify_qp mask for INIT->RTR.
const RTRAttrMask = rdma.IBV_QP_STATE | rdma.IBV_QP_AV | rdma.IBV_QP_PATH_MTU | rdma.IBV_QP_DEST_QPN | rdma.IBV_QP_RQ_PSN

// NegotiatedPathMTU returns the path MTU both peers can use, i.e. the smaller
// of the two active MTUs.
func NegotiatedPathMTU(local, remote int32) int32 {
	if MTUBytes(local) == 0 {
		local = rdma.IBV_MTU_1024
	}
	if MTUBytes(remote) == 0 {
		remote = rdma.IBV_MTU_1024
	}
	if local < remote {
		return local
	}
	return remote
}

// MTUBytes returns the byte size for an ibv_mtu value.
func MTUBytes(mtu int32) int {
	switch mtu {
	case 1:
		return 256
	case 2:
		return 512
	case rdma.IBV_MTU_1024:
		return 1024
	case 4:
		return 2048
	case 5:
		return 4096
	default:
		return 0
	}
}

// GIDIndexRangeError reports a source GID index that cannot fit in ibv_global_route.
type GIDIndexRangeError int

func (e GIDIndexRangeError) Error() string {
	return fmt.Sprintf("gid index %d out of uint8 range", e)
}

func boolByte(v bool) uint8 {
	if v {
		return 1
	}
	return 0
}
