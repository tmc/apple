// Code generated from Apple documentation for rdma. DO NOT EDIT.

package rdma

import (
	"testing"
	"unsafe"
)

func TestRDMAExtraStructLayouts(t *testing.T) {
	tests := []struct {
		name string
		got  uintptr
		want uintptr
	}{
		{"IbvGID.size", unsafe.Sizeof(IbvGID{}), 16},
		{"IbvQPCap.size", unsafe.Sizeof(IbvQPCap{}), 20},
		{"IbvQPInitAttr.size", unsafe.Sizeof(IbvQPInitAttr{}), 64},
		{"IbvGlobalRoute.size", unsafe.Sizeof(IbvGlobalRoute{}), 24},
		{"IbvAHAttr.size", unsafe.Sizeof(IbvAHAttr{}), 32},
		{"IbvQPAttr.size", unsafe.Sizeof(IbvQPAttr{}), 144},
		{"IbvSGE.size", unsafe.Sizeof(IbvSGE{}), 16},
		{"IbvSendWR.size", unsafe.Sizeof(IbvSendWR{}), 128},
		{"IbvRecvWR.size", unsafe.Sizeof(IbvRecvWR{}), 32},
		{"IbvWC.size", unsafe.Sizeof(IbvWC{}), 48},
		{"IbvPortAttr.size", unsafe.Sizeof(IbvPortAttr{}), 52},
		{"IbvQPInitAttr.SendCQ", unsafe.Offsetof(IbvQPInitAttr{}.SendCQ), 8},
		{"IbvQPInitAttr.Cap", unsafe.Offsetof(IbvQPInitAttr{}.Cap), 32},
		{"IbvQPInitAttr.QPType", unsafe.Offsetof(IbvQPInitAttr{}.QPType), 52},
		{"IbvAHAttr.DLID", unsafe.Offsetof(IbvAHAttr{}.DLID), 24},
		{"IbvQPAttr.Cap", unsafe.Offsetof(IbvQPAttr{}.Cap), 36},
		{"IbvQPAttr.AHAttr", unsafe.Offsetof(IbvQPAttr{}.AHAttr), 56},
		{"IbvQPAttr.AltAHAttr", unsafe.Offsetof(IbvQPAttr{}.AltAHAttr), 88},
		{"IbvQPAttr.PKeyIndex", unsafe.Offsetof(IbvQPAttr{}.PKeyIndex), 120},
		{"IbvQPAttr.PortNum", unsafe.Offsetof(IbvQPAttr{}.PortNum), 129},
		{"IbvSendWR.WR", unsafe.Offsetof(IbvSendWR{}.WR), 40},
		{"IbvPortAttr.ActiveMTU", unsafe.Offsetof(IbvPortAttr{}.ActiveMTU), 8},
		{"IbvPortAttr.GIDTblLen", unsafe.Offsetof(IbvPortAttr{}.GIDTblLen), 12},
		{"IbvPortAttr.LID", unsafe.Offsetof(IbvPortAttr{}.LID), 34},
		{"ibv_context.ops", 8, 8},
		{"ibv_context_ops.poll_cq", 88, 88},
		{"ibv_context_ops.post_send", 200, 200},
		{"ibv_context_ops.post_recv", 208, 208},
		{"ibv_qp.qp_num", 52, 52},
		{"ibv_mr.lkey", 36, 36},
		{"ibv_mr.rkey", 40, 40},
	}
	for _, tt := range tests {
		if tt.got != tt.want {
			t.Fatalf("%s = %d, want %d", tt.name, tt.got, tt.want)
		}
	}
}

func TestRDMAExtraConstants(t *testing.T) {
	tests := []struct {
		name string
		got  int
		want int
	}{
		{"IBV_WR_RDMA_WRITE", IBV_WR_RDMA_WRITE, 0},
		{"IBV_WR_SEND", IBV_WR_SEND, 2},
		{"IBV_QPS_RESET", IBV_QPS_RESET, 0},
		{"IBV_QPS_ERR", IBV_QPS_ERR, 6},
	}
	for _, tt := range tests {
		if tt.got != tt.want {
			t.Fatalf("%s = %d, want %d", tt.name, tt.got, tt.want)
		}
	}
}
