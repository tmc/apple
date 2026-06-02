package rdma

import (
	"errors"
	"testing"

	"github.com/tmc/apple/rdma"
)

func TestRTRAttrGlobalRoute(t *testing.T) {
	gid := rdma.IbvGID{0xfe, 0x80, 0, 0, 0, 0, 0, 0, 0xde, 0xad, 0, 0, 0, 0, 0, 1}
	attr, mask, err := RTRAttr(
		LocalQP{PortNum: 1, GIDIndex: 3, ActiveMTU: 5},
		RemoteQP{LID: 2, QPN: 99, PSN: 7, GID: gid, UseGlobal: true, ActiveMTU: 4},
		RTRPolicy{},
	)
	if err != nil {
		t.Fatalf("RTRAttr: %v", err)
	}
	if mask != RTRAttrMask {
		t.Fatalf("mask = %#x, want %#x", mask, RTRAttrMask)
	}
	if attr.QPState != rdma.IBV_QPS_RTR {
		t.Fatalf("QPState = %d, want RTR", attr.QPState)
	}
	if attr.PathMTU != 4 {
		t.Fatalf("PathMTU = %d, want 4", attr.PathMTU)
	}
	if attr.AHAttr.DLID != 2 {
		t.Fatalf("DLID = %d, want 2", attr.AHAttr.DLID)
	}
	if attr.AHAttr.IsGlobal != 1 {
		t.Fatalf("IsGlobal = %d, want 1", attr.AHAttr.IsGlobal)
	}
	if attr.AHAttr.GRH.DGID != gid {
		t.Fatalf("DGID = %x, want %x", attr.AHAttr.GRH.DGID, gid)
	}
	if attr.AHAttr.GRH.SGIDIndex != 3 {
		t.Fatalf("SGIDIndex = %d, want 3", attr.AHAttr.GRH.SGIDIndex)
	}
	if attr.AHAttr.GRH.HopLimit != 1 {
		t.Fatalf("HopLimit = %d, want 1", attr.AHAttr.GRH.HopLimit)
	}
}

func TestRTRAttrZeroDLIDWhenGlobal(t *testing.T) {
	attr, _, err := RTRAttr(
		LocalQP{GIDIndex: 1, ActiveMTU: 5},
		RemoteQP{LID: 2, QPN: 99, PSN: 7, GID: rdma.IbvGID{15: 1}, UseGlobal: true, ActiveMTU: 5},
		RTRPolicy{ZeroDLIDWhenGlobal: true, HopLimit: 255, TrafficClass: 3, FlowLabel: 4},
	)
	if err != nil {
		t.Fatalf("RTRAttr: %v", err)
	}
	if attr.AHAttr.DLID != 0 {
		t.Fatalf("DLID = %d, want 0", attr.AHAttr.DLID)
	}
	if attr.AHAttr.PortNum != 1 {
		t.Fatalf("PortNum = %d, want default 1", attr.AHAttr.PortNum)
	}
	if attr.AHAttr.GRH.HopLimit != 255 || attr.AHAttr.GRH.TrafficClass != 3 || attr.AHAttr.GRH.FlowLabel != 4 {
		t.Fatalf("GRH = %+v, want policy fields", attr.AHAttr.GRH)
	}
}

func TestRTRAttrRejectsBadGIDIndex(t *testing.T) {
	_, _, err := RTRAttr(
		LocalQP{GIDIndex: 256, ActiveMTU: 5},
		RemoteQP{QPN: 99, PSN: 7, GID: rdma.IbvGID{15: 1}, UseGlobal: true, ActiveMTU: 5},
		RTRPolicy{},
	)
	var rangeErr GIDIndexRangeError
	if !errors.As(err, &rangeErr) {
		t.Fatalf("RTRAttr err = %T %v, want GIDIndexRangeError", err, err)
	}
}
